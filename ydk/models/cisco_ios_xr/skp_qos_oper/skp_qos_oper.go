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

// PolicyState represents Different Interface states
type PolicyState string

const (
    // active
    PolicyState_active PolicyState = "active"

    // suspended
    PolicyState_suspended PolicyState = "suspended"
)

// PlatformQos
// QoS Skywarp platform operational data 
type PlatformQos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes with platform specific QoS configuration.
    Nodes PlatformQos_Nodes
}

func (platformQos *PlatformQos) GetFilter() yfilter.YFilter { return platformQos.YFilter }

func (platformQos *PlatformQos) SetFilter(yf yfilter.YFilter) { platformQos.YFilter = yf }

func (platformQos *PlatformQos) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (platformQos *PlatformQos) GetSegmentPath() string {
    return "Cisco-IOS-XR-skp-qos-oper:platform-qos"
}

func (platformQos *PlatformQos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &platformQos.Nodes
    }
    return nil
}

func (platformQos *PlatformQos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &platformQos.Nodes
    return children
}

func (platformQos *PlatformQos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformQos *PlatformQos) GetBundleName() string { return "cisco_ios_xr" }

func (platformQos *PlatformQos) GetYangName() string { return "platform-qos" }

func (platformQos *PlatformQos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformQos *PlatformQos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformQos *PlatformQos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformQos *PlatformQos) SetParent(parent types.Entity) { platformQos.parent = parent }

func (platformQos *PlatformQos) GetParent() types.Entity { return platformQos.parent }

func (platformQos *PlatformQos) GetParentYangName() string { return "Cisco-IOS-XR-skp-qos-oper" }

// PlatformQos_Nodes
// List of nodes with platform specific QoS
// configuration
type PlatformQos_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node with platform specific QoS configuration. The type is slice of
    // PlatformQos_Nodes_Node.
    Node []PlatformQos_Nodes_Node
}

func (nodes *PlatformQos_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PlatformQos_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PlatformQos_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PlatformQos_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PlatformQos_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PlatformQos_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PlatformQos_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PlatformQos_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PlatformQos_Nodes) GetYangName() string { return "nodes" }

func (nodes *PlatformQos_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PlatformQos_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PlatformQos_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PlatformQos_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PlatformQos_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PlatformQos_Nodes) GetParentYangName() string { return "platform-qos" }

// PlatformQos_Nodes_Node
// Node with platform specific QoS configuration
type PlatformQos_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // QoS list of bundle interfaces.
    BundleInterfaces PlatformQos_Nodes_Node_BundleInterfaces

    // QoS system capability.
    Capability PlatformQos_Nodes_Node_Capability

    // QoS list of interfaces.
    Interfaces PlatformQos_Nodes_Node_Interfaces
}

func (node *PlatformQos_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PlatformQos_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PlatformQos_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "bundle-interfaces" { return "BundleInterfaces" }
    if yname == "capability" { return "Capability" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *PlatformQos_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *PlatformQos_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-interfaces" {
        return &node.BundleInterfaces
    }
    if childYangName == "capability" {
        return &node.Capability
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *PlatformQos_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle-interfaces"] = &node.BundleInterfaces
    children["capability"] = &node.Capability
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *PlatformQos_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *PlatformQos_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PlatformQos_Nodes_Node) GetYangName() string { return "node" }

func (node *PlatformQos_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PlatformQos_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PlatformQos_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PlatformQos_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PlatformQos_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PlatformQos_Nodes_Node) GetParentYangName() string { return "nodes" }

// PlatformQos_Nodes_Node_BundleInterfaces
// QoS list of bundle interfaces
type PlatformQos_Nodes_Node_BundleInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface.
    BundleInterface []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetFilter() yfilter.YFilter { return bundleInterfaces.YFilter }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) SetFilter(yf yfilter.YFilter) { bundleInterfaces.YFilter = yf }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetGoName(yname string) string {
    if yname == "bundle-interface" { return "BundleInterface" }
    return ""
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetSegmentPath() string {
    return "bundle-interfaces"
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-interface" {
        for _, c := range bundleInterfaces.BundleInterface {
            if bundleInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface{}
        bundleInterfaces.BundleInterface = append(bundleInterfaces.BundleInterface, child)
        return &bundleInterfaces.BundleInterface[len(bundleInterfaces.BundleInterface)-1]
    }
    return nil
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bundleInterfaces.BundleInterface {
        children[bundleInterfaces.BundleInterface[i].GetSegmentPath()] = &bundleInterfaces.BundleInterface[i]
    }
    return children
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetYangName() string { return "bundle-interfaces" }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) SetParent(parent types.Entity) { bundleInterfaces.parent = parent }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetParent() types.Entity { return bundleInterfaces.parent }

func (bundleInterfaces *PlatformQos_Nodes_Node_BundleInterfaces) GetParentYangName() string { return "node" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface
// QoS interface name
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bundle interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetFilter() yfilter.YFilter { return bundleInterface.YFilter }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) SetFilter(yf yfilter.YFilter) { bundleInterface.YFilter = yf }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "member-interfaces" { return "MemberInterfaces" }
    return ""
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetSegmentPath() string {
    return "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member-interfaces" {
        return &bundleInterface.MemberInterfaces
    }
    return nil
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["member-interfaces"] = &bundleInterface.MemberInterfaces
    return children
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bundleInterface.InterfaceName
    return leafs
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetYangName() string { return "bundle-interface" }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) SetParent(parent types.Entity) { bundleInterface.parent = parent }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetParent() types.Entity { return bundleInterface.parent }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetParentYangName() string { return "bundle-interfaces" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
// QoS list of member interfaces
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetFilter() yfilter.YFilter { return memberInterfaces.YFilter }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) SetFilter(yf yfilter.YFilter) { memberInterfaces.YFilter = yf }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    return ""
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetSegmentPath() string {
    return "member-interfaces"
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member-interface" {
        for _, c := range memberInterfaces.MemberInterface {
            if memberInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface{}
        memberInterfaces.MemberInterface = append(memberInterfaces.MemberInterface, child)
        return &memberInterfaces.MemberInterface[len(memberInterfaces.MemberInterface)-1]
    }
    return nil
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range memberInterfaces.MemberInterface {
        children[memberInterfaces.MemberInterface[i].GetSegmentPath()] = &memberInterfaces.MemberInterface[i]
    }
    return children
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetYangName() string { return "member-interfaces" }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) SetParent(parent types.Entity) { memberInterfaces.parent = parent }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetParent() types.Entity { return memberInterfaces.parent }

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetParentYangName() string { return "bundle-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
// QoS interface name
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Memeber interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // QoS policy direction input.
    BundleInput PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput

    // QoS policy direction output.
    BundleOutput PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetFilter() yfilter.YFilter { return memberInterface.YFilter }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) SetFilter(yf yfilter.YFilter) { memberInterface.YFilter = yf }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "bundle-input" { return "BundleInput" }
    if yname == "bundle-output" { return "BundleOutput" }
    return ""
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetSegmentPath() string {
    return "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-input" {
        return &memberInterface.BundleInput
    }
    if childYangName == "bundle-output" {
        return &memberInterface.BundleOutput
    }
    return nil
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle-input"] = &memberInterface.BundleInput
    children["bundle-output"] = &memberInterface.BundleOutput
    return children
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = memberInterface.InterfaceName
    return leafs
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetBundleName() string { return "cisco_ios_xr" }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetYangName() string { return "member-interface" }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) SetParent(parent types.Entity) { memberInterface.parent = parent }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetParent() types.Entity { return memberInterface.parent }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetParentYangName() string { return "member-interfaces" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput
// QoS policy direction input
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetFilter() yfilter.YFilter { return bundleInput.YFilter }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) SetFilter(yf yfilter.YFilter) { bundleInput.YFilter = yf }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetSegmentPath() string {
    return "bundle-input"
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &bundleInput.Header
    }
    if childYangName == "interface-parameters" {
        return &bundleInput.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &bundleInput.SkywarpQosPolicyClass
    }
    return nil
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &bundleInput.Header
    children["interface-parameters"] = &bundleInput.InterfaceParameters
    children["skywarp-qos-policy-class"] = &bundleInput.SkywarpQosPolicyClass
    return children
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetYangName() string { return "bundle-input" }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) SetParent(parent types.Entity) { bundleInput.parent = parent }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetParent() types.Entity { return bundleInput.parent }

func (bundleInput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetParentYangName() string { return "member-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header struct {
    parent types.Entity
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

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetYangName() string { return "header" }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Header) GetParentYangName() string { return "bundle-input" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetGoName(yname string) string {
    if yname == "interface-config-rate" { return "InterfaceConfigRate" }
    if yname == "interface-program-rate" { return "InterfaceProgramRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    return ""
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-config-rate" {
        return &interfaceParameters.InterfaceConfigRate
    }
    if childYangName == "interface-program-rate" {
        return &interfaceParameters.InterfaceProgramRate
    }
    if childYangName == "port-shaper-rate" {
        return &interfaceParameters.PortShaperRate
    }
    return nil
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-config-rate"] = &interfaceParameters.InterfaceConfigRate
    children["interface-program-rate"] = &interfaceParameters.InterfaceProgramRate
    children["port-shaper-rate"] = &interfaceParameters.PortShaperRate
    return children
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters) GetParentYangName() string { return "bundle-input" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetFilter() yfilter.YFilter { return interfaceConfigRate.YFilter }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) SetFilter(yf yfilter.YFilter) { interfaceConfigRate.YFilter = yf }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetSegmentPath() string {
    return "interface-config-rate"
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceConfigRate.Value
    leafs["unit"] = interfaceConfigRate.Unit
    return leafs
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetYangName() string { return "interface-config-rate" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) SetParent(parent types.Entity) { interfaceConfigRate.parent = parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetParent() types.Entity { return interfaceConfigRate.parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceConfigRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetFilter() yfilter.YFilter { return interfaceProgramRate.YFilter }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) SetFilter(yf yfilter.YFilter) { interfaceProgramRate.YFilter = yf }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetSegmentPath() string {
    return "interface-program-rate"
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceProgramRate.Value
    leafs["unit"] = interfaceProgramRate.Unit
    return leafs
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetYangName() string { return "interface-program-rate" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) SetParent(parent types.Entity) { interfaceProgramRate.parent = parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetParent() types.Entity { return interfaceProgramRate.parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_InterfaceProgramRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetFilter() yfilter.YFilter { return portShaperRate.YFilter }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) SetFilter(yf yfilter.YFilter) { portShaperRate.YFilter = yf }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetSegmentPath() string {
    return "port-shaper-rate"
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = portShaperRate.Value
    leafs["unit"] = portShaperRate.Unit
    return leafs
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetBundleName() string { return "cisco_ios_xr" }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetYangName() string { return "port-shaper-rate" }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) SetParent(parent types.Entity) { portShaperRate.parent = parent }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetParent() types.Entity { return portShaperRate.parent }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_InterfaceParameters_PortShaperRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-pclass-st" { return "QosShowPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt{}
        skywarpQosPolicyClass.QosShowPclassSt = append(skywarpQosPolicyClass.QosShowPclassSt, child)
        return &skywarpQosPolicyClass.QosShowPclassSt[len(skywarpQosPolicyClass.QosShowPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        children[skywarpQosPolicyClass.QosShowPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass) GetParentYangName() string { return "bundle-input" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt struct {
    parent types.Entity
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

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetFilter() yfilter.YFilter { return qosShowPclassSt.YFilter }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) SetFilter(yf yfilter.YFilter) { qosShowPclassSt.YFilter = yf }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetGoName(yname string) string {
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "queue" { return "Queue" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    if yname == "police" { return "Police" }
    if yname == "marking" { return "Marking" }
    return ""
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetSegmentPath() string {
    return "qos-show-pclass-st"
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &qosShowPclassSt.Queue
    }
    if childYangName == "shape" {
        return &qosShowPclassSt.Shape
    }
    if childYangName == "wfq" {
        return &qosShowPclassSt.Wfq
    }
    if childYangName == "police" {
        return &qosShowPclassSt.Police
    }
    if childYangName == "marking" {
        return &qosShowPclassSt.Marking
    }
    return nil
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &qosShowPclassSt.Queue
    children["shape"] = &qosShowPclassSt.Shape
    children["wfq"] = &qosShowPclassSt.Wfq
    children["police"] = &qosShowPclassSt.Police
    children["marking"] = &qosShowPclassSt.Marking
    return children
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-level"] = qosShowPclassSt.ClassLevel
    leafs["class-name"] = qosShowPclassSt.ClassName
    return leafs
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetYangName() string { return "qos-show-pclass-st" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) SetParent(parent types.Entity) { qosShowPclassSt.parent = parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetParent() types.Entity { return qosShowPclassSt.parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "queue-type" { return "QueueType" }
    return ""
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["queue-type"] = queue.QueueType
    return leafs
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetGoName(yname string) string {
    if yname == "excess-weight" { return "ExcessWeight" }
    if yname == "committed-weight" { return "CommittedWeight" }
    if yname == "programmed-wfq" { return "ProgrammedWfq" }
    return ""
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "committed-weight" {
        return &wfq.CommittedWeight
    }
    if childYangName == "programmed-wfq" {
        return &wfq.ProgrammedWfq
    }
    return nil
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["committed-weight"] = &wfq.CommittedWeight
    children["programmed-wfq"] = &wfq.ProgrammedWfq
    return children
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-weight"] = wfq.ExcessWeight
    return leafs
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetFilter() yfilter.YFilter { return committedWeight.YFilter }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetFilter(yf yfilter.YFilter) { committedWeight.YFilter = yf }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetSegmentPath() string {
    return "committed-weight"
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = committedWeight.Value
    leafs["unit"] = committedWeight.Unit
    return leafs
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleName() string { return "cisco_ios_xr" }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetYangName() string { return "committed-weight" }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetParent(parent types.Entity) { committedWeight.parent = parent }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParent() types.Entity { return committedWeight.parent }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetFilter() yfilter.YFilter { return programmedWfq.YFilter }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetFilter(yf yfilter.YFilter) { programmedWfq.YFilter = yf }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetSegmentPath() string {
    return "programmed-wfq"
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &programmedWfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &programmedWfq.SumOfBandwidth
    }
    return nil
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &programmedWfq.Bandwidth
    children["sum-of-bandwidth"] = &programmedWfq.SumOfBandwidth
    return children
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = programmedWfq.ExcessRatio
    return leafs
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleName() string { return "cisco_ios_xr" }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetYangName() string { return "programmed-wfq" }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetParent(parent types.Entity) { programmedWfq.parent = parent }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParent() types.Entity { return programmedWfq.parent }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    parent types.Entity
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

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetGoName(yname string) string {
    if yname == "policer-id" { return "PolicerId" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policer-id"] = police.PolicerId
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetYangName() string { return "police" }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetFilter() yfilter.YFilter { return marking.YFilter }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetFilter(yf yfilter.YFilter) { marking.YFilter = yf }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetGoName(yname string) string {
    if yname == "mark-only" { return "MarkOnly" }
    if yname == "police-conform" { return "PoliceConform" }
    if yname == "police-exceed" { return "PoliceExceed" }
    return ""
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetSegmentPath() string {
    return "marking"
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-only" {
        return &marking.MarkOnly
    }
    if childYangName == "police-conform" {
        return &marking.PoliceConform
    }
    if childYangName == "police-exceed" {
        return &marking.PoliceExceed
    }
    return nil
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mark-only"] = &marking.MarkOnly
    children["police-conform"] = &marking.PoliceConform
    children["police-exceed"] = &marking.PoliceExceed
    return children
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleName() string { return "cisco_ios_xr" }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetYangName() string { return "marking" }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetParent(parent types.Entity) { marking.parent = parent }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParent() types.Entity { return marking.parent }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetFilter() yfilter.YFilter { return markOnly.YFilter }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetFilter(yf yfilter.YFilter) { markOnly.YFilter = yf }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetSegmentPath() string {
    return "mark-only"
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range markOnly.MarkDetail {
            if markOnly.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail{}
        markOnly.MarkDetail = append(markOnly.MarkDetail, child)
        return &markOnly.MarkDetail[len(markOnly.MarkDetail)-1]
    }
    return nil
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range markOnly.MarkDetail {
        children[markOnly.MarkDetail[i].GetSegmentPath()] = &markOnly.MarkDetail[i]
    }
    return children
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = markOnly.ActionType
    return leafs
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleName() string { return "cisco_ios_xr" }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetYangName() string { return "mark-only" }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetParent(parent types.Entity) { markOnly.parent = parent }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParent() types.Entity { return markOnly.parent }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParentYangName() string { return "mark-only" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetFilter() yfilter.YFilter { return policeConform.YFilter }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetFilter(yf yfilter.YFilter) { policeConform.YFilter = yf }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetSegmentPath() string {
    return "police-conform"
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeConform.MarkDetail {
            if policeConform.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail{}
        policeConform.MarkDetail = append(policeConform.MarkDetail, child)
        return &policeConform.MarkDetail[len(policeConform.MarkDetail)-1]
    }
    return nil
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeConform.MarkDetail {
        children[policeConform.MarkDetail[i].GetSegmentPath()] = &policeConform.MarkDetail[i]
    }
    return children
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeConform.ActionType
    return leafs
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleName() string { return "cisco_ios_xr" }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetYangName() string { return "police-conform" }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetParent(parent types.Entity) { policeConform.parent = parent }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParent() types.Entity { return policeConform.parent }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParentYangName() string { return "police-conform" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetFilter() yfilter.YFilter { return policeExceed.YFilter }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetFilter(yf yfilter.YFilter) { policeExceed.YFilter = yf }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetSegmentPath() string {
    return "police-exceed"
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeExceed.MarkDetail {
            if policeExceed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail{}
        policeExceed.MarkDetail = append(policeExceed.MarkDetail, child)
        return &policeExceed.MarkDetail[len(policeExceed.MarkDetail)-1]
    }
    return nil
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeExceed.MarkDetail {
        children[policeExceed.MarkDetail[i].GetSegmentPath()] = &policeExceed.MarkDetail[i]
    }
    return children
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeExceed.ActionType
    return leafs
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleName() string { return "cisco_ios_xr" }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetYangName() string { return "police-exceed" }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetParent(parent types.Entity) { policeExceed.parent = parent }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParent() types.Entity { return policeExceed.parent }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParentYangName() string { return "police-exceed" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput
// QoS policy direction output
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetFilter() yfilter.YFilter { return bundleOutput.YFilter }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) SetFilter(yf yfilter.YFilter) { bundleOutput.YFilter = yf }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetSegmentPath() string {
    return "bundle-output"
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &bundleOutput.Header
    }
    if childYangName == "interface-parameters" {
        return &bundleOutput.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &bundleOutput.SkywarpQosPolicyClass
    }
    return nil
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &bundleOutput.Header
    children["interface-parameters"] = &bundleOutput.InterfaceParameters
    children["skywarp-qos-policy-class"] = &bundleOutput.SkywarpQosPolicyClass
    return children
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetBundleName() string { return "cisco_ios_xr" }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetYangName() string { return "bundle-output" }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) SetParent(parent types.Entity) { bundleOutput.parent = parent }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetParent() types.Entity { return bundleOutput.parent }

func (bundleOutput *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetParentYangName() string { return "member-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header struct {
    parent types.Entity
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

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetYangName() string { return "header" }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Header) GetParentYangName() string { return "bundle-output" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetGoName(yname string) string {
    if yname == "interface-config-rate" { return "InterfaceConfigRate" }
    if yname == "interface-program-rate" { return "InterfaceProgramRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    return ""
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-config-rate" {
        return &interfaceParameters.InterfaceConfigRate
    }
    if childYangName == "interface-program-rate" {
        return &interfaceParameters.InterfaceProgramRate
    }
    if childYangName == "port-shaper-rate" {
        return &interfaceParameters.PortShaperRate
    }
    return nil
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-config-rate"] = &interfaceParameters.InterfaceConfigRate
    children["interface-program-rate"] = &interfaceParameters.InterfaceProgramRate
    children["port-shaper-rate"] = &interfaceParameters.PortShaperRate
    return children
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters) GetParentYangName() string { return "bundle-output" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetFilter() yfilter.YFilter { return interfaceConfigRate.YFilter }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) SetFilter(yf yfilter.YFilter) { interfaceConfigRate.YFilter = yf }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetSegmentPath() string {
    return "interface-config-rate"
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceConfigRate.Value
    leafs["unit"] = interfaceConfigRate.Unit
    return leafs
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetYangName() string { return "interface-config-rate" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) SetParent(parent types.Entity) { interfaceConfigRate.parent = parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetParent() types.Entity { return interfaceConfigRate.parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceConfigRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetFilter() yfilter.YFilter { return interfaceProgramRate.YFilter }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) SetFilter(yf yfilter.YFilter) { interfaceProgramRate.YFilter = yf }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetSegmentPath() string {
    return "interface-program-rate"
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceProgramRate.Value
    leafs["unit"] = interfaceProgramRate.Unit
    return leafs
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetYangName() string { return "interface-program-rate" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) SetParent(parent types.Entity) { interfaceProgramRate.parent = parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetParent() types.Entity { return interfaceProgramRate.parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_InterfaceProgramRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetFilter() yfilter.YFilter { return portShaperRate.YFilter }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) SetFilter(yf yfilter.YFilter) { portShaperRate.YFilter = yf }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetSegmentPath() string {
    return "port-shaper-rate"
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = portShaperRate.Value
    leafs["unit"] = portShaperRate.Unit
    return leafs
}

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetBundleName() string { return "cisco_ios_xr" }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetYangName() string { return "port-shaper-rate" }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) SetParent(parent types.Entity) { portShaperRate.parent = parent }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetParent() types.Entity { return portShaperRate.parent }

func (portShaperRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_InterfaceParameters_PortShaperRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-pclass-st" { return "QosShowPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt{}
        skywarpQosPolicyClass.QosShowPclassSt = append(skywarpQosPolicyClass.QosShowPclassSt, child)
        return &skywarpQosPolicyClass.QosShowPclassSt[len(skywarpQosPolicyClass.QosShowPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        children[skywarpQosPolicyClass.QosShowPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass) GetParentYangName() string { return "bundle-output" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt struct {
    parent types.Entity
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

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetFilter() yfilter.YFilter { return qosShowPclassSt.YFilter }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) SetFilter(yf yfilter.YFilter) { qosShowPclassSt.YFilter = yf }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetGoName(yname string) string {
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "queue" { return "Queue" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    if yname == "police" { return "Police" }
    if yname == "marking" { return "Marking" }
    return ""
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetSegmentPath() string {
    return "qos-show-pclass-st"
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &qosShowPclassSt.Queue
    }
    if childYangName == "shape" {
        return &qosShowPclassSt.Shape
    }
    if childYangName == "wfq" {
        return &qosShowPclassSt.Wfq
    }
    if childYangName == "police" {
        return &qosShowPclassSt.Police
    }
    if childYangName == "marking" {
        return &qosShowPclassSt.Marking
    }
    return nil
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &qosShowPclassSt.Queue
    children["shape"] = &qosShowPclassSt.Shape
    children["wfq"] = &qosShowPclassSt.Wfq
    children["police"] = &qosShowPclassSt.Police
    children["marking"] = &qosShowPclassSt.Marking
    return children
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-level"] = qosShowPclassSt.ClassLevel
    leafs["class-name"] = qosShowPclassSt.ClassName
    return leafs
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetYangName() string { return "qos-show-pclass-st" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) SetParent(parent types.Entity) { qosShowPclassSt.parent = parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetParent() types.Entity { return qosShowPclassSt.parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "queue-type" { return "QueueType" }
    return ""
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["queue-type"] = queue.QueueType
    return leafs
}

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetGoName(yname string) string {
    if yname == "excess-weight" { return "ExcessWeight" }
    if yname == "committed-weight" { return "CommittedWeight" }
    if yname == "programmed-wfq" { return "ProgrammedWfq" }
    return ""
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "committed-weight" {
        return &wfq.CommittedWeight
    }
    if childYangName == "programmed-wfq" {
        return &wfq.ProgrammedWfq
    }
    return nil
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["committed-weight"] = &wfq.CommittedWeight
    children["programmed-wfq"] = &wfq.ProgrammedWfq
    return children
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-weight"] = wfq.ExcessWeight
    return leafs
}

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetFilter() yfilter.YFilter { return committedWeight.YFilter }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetFilter(yf yfilter.YFilter) { committedWeight.YFilter = yf }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetSegmentPath() string {
    return "committed-weight"
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = committedWeight.Value
    leafs["unit"] = committedWeight.Unit
    return leafs
}

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleName() string { return "cisco_ios_xr" }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetYangName() string { return "committed-weight" }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetParent(parent types.Entity) { committedWeight.parent = parent }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParent() types.Entity { return committedWeight.parent }

func (committedWeight *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetFilter() yfilter.YFilter { return programmedWfq.YFilter }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetFilter(yf yfilter.YFilter) { programmedWfq.YFilter = yf }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetSegmentPath() string {
    return "programmed-wfq"
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &programmedWfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &programmedWfq.SumOfBandwidth
    }
    return nil
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &programmedWfq.Bandwidth
    children["sum-of-bandwidth"] = &programmedWfq.SumOfBandwidth
    return children
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = programmedWfq.ExcessRatio
    return leafs
}

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleName() string { return "cisco_ios_xr" }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetYangName() string { return "programmed-wfq" }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetParent(parent types.Entity) { programmedWfq.parent = parent }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParent() types.Entity { return programmedWfq.parent }

func (programmedWfq *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    parent types.Entity
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

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetGoName(yname string) string {
    if yname == "policer-id" { return "PolicerId" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policer-id"] = police.PolicerId
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetYangName() string { return "police" }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetFilter() yfilter.YFilter { return marking.YFilter }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetFilter(yf yfilter.YFilter) { marking.YFilter = yf }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetGoName(yname string) string {
    if yname == "mark-only" { return "MarkOnly" }
    if yname == "police-conform" { return "PoliceConform" }
    if yname == "police-exceed" { return "PoliceExceed" }
    return ""
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetSegmentPath() string {
    return "marking"
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-only" {
        return &marking.MarkOnly
    }
    if childYangName == "police-conform" {
        return &marking.PoliceConform
    }
    if childYangName == "police-exceed" {
        return &marking.PoliceExceed
    }
    return nil
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mark-only"] = &marking.MarkOnly
    children["police-conform"] = &marking.PoliceConform
    children["police-exceed"] = &marking.PoliceExceed
    return children
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleName() string { return "cisco_ios_xr" }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetYangName() string { return "marking" }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetParent(parent types.Entity) { marking.parent = parent }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParent() types.Entity { return marking.parent }

func (marking *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetFilter() yfilter.YFilter { return markOnly.YFilter }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetFilter(yf yfilter.YFilter) { markOnly.YFilter = yf }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetSegmentPath() string {
    return "mark-only"
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range markOnly.MarkDetail {
            if markOnly.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail{}
        markOnly.MarkDetail = append(markOnly.MarkDetail, child)
        return &markOnly.MarkDetail[len(markOnly.MarkDetail)-1]
    }
    return nil
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range markOnly.MarkDetail {
        children[markOnly.MarkDetail[i].GetSegmentPath()] = &markOnly.MarkDetail[i]
    }
    return children
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = markOnly.ActionType
    return leafs
}

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleName() string { return "cisco_ios_xr" }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetYangName() string { return "mark-only" }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetParent(parent types.Entity) { markOnly.parent = parent }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParent() types.Entity { return markOnly.parent }

func (markOnly *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParentYangName() string { return "mark-only" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetFilter() yfilter.YFilter { return policeConform.YFilter }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetFilter(yf yfilter.YFilter) { policeConform.YFilter = yf }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetSegmentPath() string {
    return "police-conform"
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeConform.MarkDetail {
            if policeConform.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail{}
        policeConform.MarkDetail = append(policeConform.MarkDetail, child)
        return &policeConform.MarkDetail[len(policeConform.MarkDetail)-1]
    }
    return nil
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeConform.MarkDetail {
        children[policeConform.MarkDetail[i].GetSegmentPath()] = &policeConform.MarkDetail[i]
    }
    return children
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeConform.ActionType
    return leafs
}

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleName() string { return "cisco_ios_xr" }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetYangName() string { return "police-conform" }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetParent(parent types.Entity) { policeConform.parent = parent }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParent() types.Entity { return policeConform.parent }

func (policeConform *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParentYangName() string { return "police-conform" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetFilter() yfilter.YFilter { return policeExceed.YFilter }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetFilter(yf yfilter.YFilter) { policeExceed.YFilter = yf }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetSegmentPath() string {
    return "police-exceed"
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeExceed.MarkDetail {
            if policeExceed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail{}
        policeExceed.MarkDetail = append(policeExceed.MarkDetail, child)
        return &policeExceed.MarkDetail[len(policeExceed.MarkDetail)-1]
    }
    return nil
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeExceed.MarkDetail {
        children[policeExceed.MarkDetail[i].GetSegmentPath()] = &policeExceed.MarkDetail[i]
    }
    return children
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeExceed.ActionType
    return leafs
}

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleName() string { return "cisco_ios_xr" }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetYangName() string { return "police-exceed" }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetParent(parent types.Entity) { policeExceed.parent = parent }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParent() types.Entity { return policeExceed.parent }

func (policeExceed *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParentYangName() string { return "police-exceed" }

// PlatformQos_Nodes_Node_Capability
// QoS system capability
type PlatformQos_Nodes_Node_Capability struct {
    parent types.Entity
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

func (capability *PlatformQos_Nodes_Node_Capability) GetFilter() yfilter.YFilter { return capability.YFilter }

func (capability *PlatformQos_Nodes_Node_Capability) SetFilter(yf yfilter.YFilter) { capability.YFilter = yf }

func (capability *PlatformQos_Nodes_Node_Capability) GetGoName(yname string) string {
    if yname == "max-policy-maps" { return "MaxPolicyMaps" }
    if yname == "max-policy-hierarchy" { return "MaxPolicyHierarchy" }
    if yname == "max-policy-name-length" { return "MaxPolicyNameLength" }
    if yname == "max-classes-per-policy" { return "MaxClassesPerPolicy" }
    if yname == "max-police-actions-per-class" { return "MaxPoliceActionsPerClass" }
    if yname == "max-marking-actions-per-class" { return "MaxMarkingActionsPerClass" }
    if yname == "max-matches-per-class" { return "MaxMatchesPerClass" }
    if yname == "max-classmap-name-length" { return "MaxClassmapNameLength" }
    if yname == "max-bundle-members" { return "MaxBundleMembers" }
    return ""
}

func (capability *PlatformQos_Nodes_Node_Capability) GetSegmentPath() string {
    return "capability"
}

func (capability *PlatformQos_Nodes_Node_Capability) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capability *PlatformQos_Nodes_Node_Capability) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capability *PlatformQos_Nodes_Node_Capability) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-policy-maps"] = capability.MaxPolicyMaps
    leafs["max-policy-hierarchy"] = capability.MaxPolicyHierarchy
    leafs["max-policy-name-length"] = capability.MaxPolicyNameLength
    leafs["max-classes-per-policy"] = capability.MaxClassesPerPolicy
    leafs["max-police-actions-per-class"] = capability.MaxPoliceActionsPerClass
    leafs["max-marking-actions-per-class"] = capability.MaxMarkingActionsPerClass
    leafs["max-matches-per-class"] = capability.MaxMatchesPerClass
    leafs["max-classmap-name-length"] = capability.MaxClassmapNameLength
    leafs["max-bundle-members"] = capability.MaxBundleMembers
    return leafs
}

func (capability *PlatformQos_Nodes_Node_Capability) GetBundleName() string { return "cisco_ios_xr" }

func (capability *PlatformQos_Nodes_Node_Capability) GetYangName() string { return "capability" }

func (capability *PlatformQos_Nodes_Node_Capability) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (capability *PlatformQos_Nodes_Node_Capability) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (capability *PlatformQos_Nodes_Node_Capability) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (capability *PlatformQos_Nodes_Node_Capability) SetParent(parent types.Entity) { capability.parent = parent }

func (capability *PlatformQos_Nodes_Node_Capability) GetParent() types.Entity { return capability.parent }

func (capability *PlatformQos_Nodes_Node_Capability) GetParentYangName() string { return "node" }

// PlatformQos_Nodes_Node_Interfaces
// QoS list of interfaces
type PlatformQos_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS interface name. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface.
    Interface []PlatformQos_Nodes_Node_Interfaces_Interface
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *PlatformQos_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// PlatformQos_Nodes_Node_Interfaces_Interface
// QoS interface name
type PlatformQos_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // QoS policy direction egress.
    Output PlatformQos_Nodes_Node_Interfaces_Interface_Output

    // QoS policy direction ingress.
    Input PlatformQos_Nodes_Node_Interfaces_Interface_Input
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "output" { return "Output" }
    if yname == "input" { return "Input" }
    return ""
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &self.Output
    }
    if childYangName == "input" {
        return &self.Input
    }
    return nil
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &self.Output
    children["input"] = &self.Input
    return children
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output
// QoS policy direction egress
type PlatformQos_Nodes_Node_Interfaces_Interface_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetSegmentPath() string {
    return "output"
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &output.Header
    }
    if childYangName == "interface-parameters" {
        return &output.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &output.SkywarpQosPolicyClass
    }
    return nil
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &output.Header
    children["interface-parameters"] = &output.InterfaceParameters
    children["skywarp-qos-policy-class"] = &output.SkywarpQosPolicyClass
    return children
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetYangName() string { return "output" }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetParent() types.Entity { return output.parent }

func (output *PlatformQos_Nodes_Node_Interfaces_Interface_Output) GetParentYangName() string { return "interface" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header struct {
    parent types.Entity
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

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetYangName() string { return "header" }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Output_Header) GetParentYangName() string { return "output" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetGoName(yname string) string {
    if yname == "interface-config-rate" { return "InterfaceConfigRate" }
    if yname == "interface-program-rate" { return "InterfaceProgramRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    return ""
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-config-rate" {
        return &interfaceParameters.InterfaceConfigRate
    }
    if childYangName == "interface-program-rate" {
        return &interfaceParameters.InterfaceProgramRate
    }
    if childYangName == "port-shaper-rate" {
        return &interfaceParameters.PortShaperRate
    }
    return nil
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-config-rate"] = &interfaceParameters.InterfaceConfigRate
    children["interface-program-rate"] = &interfaceParameters.InterfaceProgramRate
    children["port-shaper-rate"] = &interfaceParameters.PortShaperRate
    return children
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters) GetParentYangName() string { return "output" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetFilter() yfilter.YFilter { return interfaceConfigRate.YFilter }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) SetFilter(yf yfilter.YFilter) { interfaceConfigRate.YFilter = yf }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetSegmentPath() string {
    return "interface-config-rate"
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceConfigRate.Value
    leafs["unit"] = interfaceConfigRate.Unit
    return leafs
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetYangName() string { return "interface-config-rate" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) SetParent(parent types.Entity) { interfaceConfigRate.parent = parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetParent() types.Entity { return interfaceConfigRate.parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceConfigRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetFilter() yfilter.YFilter { return interfaceProgramRate.YFilter }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) SetFilter(yf yfilter.YFilter) { interfaceProgramRate.YFilter = yf }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetSegmentPath() string {
    return "interface-program-rate"
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceProgramRate.Value
    leafs["unit"] = interfaceProgramRate.Unit
    return leafs
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetYangName() string { return "interface-program-rate" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) SetParent(parent types.Entity) { interfaceProgramRate.parent = parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetParent() types.Entity { return interfaceProgramRate.parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_InterfaceProgramRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetFilter() yfilter.YFilter { return portShaperRate.YFilter }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) SetFilter(yf yfilter.YFilter) { portShaperRate.YFilter = yf }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetSegmentPath() string {
    return "port-shaper-rate"
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = portShaperRate.Value
    leafs["unit"] = portShaperRate.Unit
    return leafs
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetBundleName() string { return "cisco_ios_xr" }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetYangName() string { return "port-shaper-rate" }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) SetParent(parent types.Entity) { portShaperRate.parent = parent }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetParent() types.Entity { return portShaperRate.parent }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Output_InterfaceParameters_PortShaperRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-pclass-st" { return "QosShowPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt{}
        skywarpQosPolicyClass.QosShowPclassSt = append(skywarpQosPolicyClass.QosShowPclassSt, child)
        return &skywarpQosPolicyClass.QosShowPclassSt[len(skywarpQosPolicyClass.QosShowPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        children[skywarpQosPolicyClass.QosShowPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass) GetParentYangName() string { return "output" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt struct {
    parent types.Entity
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

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetFilter() yfilter.YFilter { return qosShowPclassSt.YFilter }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) SetFilter(yf yfilter.YFilter) { qosShowPclassSt.YFilter = yf }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetGoName(yname string) string {
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "queue" { return "Queue" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    if yname == "police" { return "Police" }
    if yname == "marking" { return "Marking" }
    return ""
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetSegmentPath() string {
    return "qos-show-pclass-st"
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &qosShowPclassSt.Queue
    }
    if childYangName == "shape" {
        return &qosShowPclassSt.Shape
    }
    if childYangName == "wfq" {
        return &qosShowPclassSt.Wfq
    }
    if childYangName == "police" {
        return &qosShowPclassSt.Police
    }
    if childYangName == "marking" {
        return &qosShowPclassSt.Marking
    }
    return nil
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &qosShowPclassSt.Queue
    children["shape"] = &qosShowPclassSt.Shape
    children["wfq"] = &qosShowPclassSt.Wfq
    children["police"] = &qosShowPclassSt.Police
    children["marking"] = &qosShowPclassSt.Marking
    return children
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-level"] = qosShowPclassSt.ClassLevel
    leafs["class-name"] = qosShowPclassSt.ClassName
    return leafs
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetYangName() string { return "qos-show-pclass-st" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) SetParent(parent types.Entity) { qosShowPclassSt.parent = parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetParent() types.Entity { return qosShowPclassSt.parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "queue-type" { return "QueueType" }
    return ""
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["queue-type"] = queue.QueueType
    return leafs
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetGoName(yname string) string {
    if yname == "excess-weight" { return "ExcessWeight" }
    if yname == "committed-weight" { return "CommittedWeight" }
    if yname == "programmed-wfq" { return "ProgrammedWfq" }
    return ""
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "committed-weight" {
        return &wfq.CommittedWeight
    }
    if childYangName == "programmed-wfq" {
        return &wfq.ProgrammedWfq
    }
    return nil
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["committed-weight"] = &wfq.CommittedWeight
    children["programmed-wfq"] = &wfq.ProgrammedWfq
    return children
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-weight"] = wfq.ExcessWeight
    return leafs
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetFilter() yfilter.YFilter { return committedWeight.YFilter }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetFilter(yf yfilter.YFilter) { committedWeight.YFilter = yf }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetSegmentPath() string {
    return "committed-weight"
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = committedWeight.Value
    leafs["unit"] = committedWeight.Unit
    return leafs
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleName() string { return "cisco_ios_xr" }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetYangName() string { return "committed-weight" }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetParent(parent types.Entity) { committedWeight.parent = parent }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParent() types.Entity { return committedWeight.parent }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetFilter() yfilter.YFilter { return programmedWfq.YFilter }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetFilter(yf yfilter.YFilter) { programmedWfq.YFilter = yf }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetSegmentPath() string {
    return "programmed-wfq"
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &programmedWfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &programmedWfq.SumOfBandwidth
    }
    return nil
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &programmedWfq.Bandwidth
    children["sum-of-bandwidth"] = &programmedWfq.SumOfBandwidth
    return children
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = programmedWfq.ExcessRatio
    return leafs
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleName() string { return "cisco_ios_xr" }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetYangName() string { return "programmed-wfq" }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetParent(parent types.Entity) { programmedWfq.parent = parent }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParent() types.Entity { return programmedWfq.parent }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    parent types.Entity
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

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetGoName(yname string) string {
    if yname == "policer-id" { return "PolicerId" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policer-id"] = police.PolicerId
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetYangName() string { return "police" }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetFilter() yfilter.YFilter { return marking.YFilter }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetFilter(yf yfilter.YFilter) { marking.YFilter = yf }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetGoName(yname string) string {
    if yname == "mark-only" { return "MarkOnly" }
    if yname == "police-conform" { return "PoliceConform" }
    if yname == "police-exceed" { return "PoliceExceed" }
    return ""
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetSegmentPath() string {
    return "marking"
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-only" {
        return &marking.MarkOnly
    }
    if childYangName == "police-conform" {
        return &marking.PoliceConform
    }
    if childYangName == "police-exceed" {
        return &marking.PoliceExceed
    }
    return nil
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mark-only"] = &marking.MarkOnly
    children["police-conform"] = &marking.PoliceConform
    children["police-exceed"] = &marking.PoliceExceed
    return children
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleName() string { return "cisco_ios_xr" }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetYangName() string { return "marking" }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetParent(parent types.Entity) { marking.parent = parent }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParent() types.Entity { return marking.parent }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetFilter() yfilter.YFilter { return markOnly.YFilter }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetFilter(yf yfilter.YFilter) { markOnly.YFilter = yf }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetSegmentPath() string {
    return "mark-only"
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range markOnly.MarkDetail {
            if markOnly.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail{}
        markOnly.MarkDetail = append(markOnly.MarkDetail, child)
        return &markOnly.MarkDetail[len(markOnly.MarkDetail)-1]
    }
    return nil
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range markOnly.MarkDetail {
        children[markOnly.MarkDetail[i].GetSegmentPath()] = &markOnly.MarkDetail[i]
    }
    return children
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = markOnly.ActionType
    return leafs
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleName() string { return "cisco_ios_xr" }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetYangName() string { return "mark-only" }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetParent(parent types.Entity) { markOnly.parent = parent }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParent() types.Entity { return markOnly.parent }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParentYangName() string { return "mark-only" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetFilter() yfilter.YFilter { return policeConform.YFilter }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetFilter(yf yfilter.YFilter) { policeConform.YFilter = yf }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetSegmentPath() string {
    return "police-conform"
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeConform.MarkDetail {
            if policeConform.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail{}
        policeConform.MarkDetail = append(policeConform.MarkDetail, child)
        return &policeConform.MarkDetail[len(policeConform.MarkDetail)-1]
    }
    return nil
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeConform.MarkDetail {
        children[policeConform.MarkDetail[i].GetSegmentPath()] = &policeConform.MarkDetail[i]
    }
    return children
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeConform.ActionType
    return leafs
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleName() string { return "cisco_ios_xr" }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetYangName() string { return "police-conform" }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetParent(parent types.Entity) { policeConform.parent = parent }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParent() types.Entity { return policeConform.parent }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParentYangName() string { return "police-conform" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetFilter() yfilter.YFilter { return policeExceed.YFilter }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetFilter(yf yfilter.YFilter) { policeExceed.YFilter = yf }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetSegmentPath() string {
    return "police-exceed"
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeExceed.MarkDetail {
            if policeExceed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail{}
        policeExceed.MarkDetail = append(policeExceed.MarkDetail, child)
        return &policeExceed.MarkDetail[len(policeExceed.MarkDetail)-1]
    }
    return nil
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeExceed.MarkDetail {
        children[policeExceed.MarkDetail[i].GetSegmentPath()] = &policeExceed.MarkDetail[i]
    }
    return children
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeExceed.ActionType
    return leafs
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleName() string { return "cisco_ios_xr" }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetYangName() string { return "police-exceed" }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetParent(parent types.Entity) { policeExceed.parent = parent }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParent() types.Entity { return policeExceed.parent }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Output_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParentYangName() string { return "police-exceed" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input
// QoS policy direction ingress
type PlatformQos_Nodes_Node_Interfaces_Interface_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header

    // QoS Interface Parameters.
    InterfaceParameters PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters

    // Skywarp QoS policy class details.
    SkywarpQosPolicyClass PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetSegmentPath() string {
    return "input"
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &input.Header
    }
    if childYangName == "interface-parameters" {
        return &input.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &input.SkywarpQosPolicyClass
    }
    return nil
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &input.Header
    children["interface-parameters"] = &input.InterfaceParameters
    children["skywarp-qos-policy-class"] = &input.SkywarpQosPolicyClass
    return children
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetYangName() string { return "input" }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetParent() types.Entity { return input.parent }

func (input *PlatformQos_Nodes_Node_Interfaces_Interface_Input) GetParentYangName() string { return "interface" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header
// QoS EA policy header
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header struct {
    parent types.Entity
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

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetYangName() string { return "header" }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQos_Nodes_Node_Interfaces_Interface_Input_Header) GetParentYangName() string { return "input" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters
// QoS Interface Parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Configured Rate.
    InterfaceConfigRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate

    // Interface Programmed Rate.
    InterfaceProgramRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate

    // Port Shaper Rate.
    PortShaperRate PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetGoName(yname string) string {
    if yname == "interface-config-rate" { return "InterfaceConfigRate" }
    if yname == "interface-program-rate" { return "InterfaceProgramRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    return ""
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-config-rate" {
        return &interfaceParameters.InterfaceConfigRate
    }
    if childYangName == "interface-program-rate" {
        return &interfaceParameters.InterfaceProgramRate
    }
    if childYangName == "port-shaper-rate" {
        return &interfaceParameters.PortShaperRate
    }
    return nil
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-config-rate"] = &interfaceParameters.InterfaceConfigRate
    children["interface-program-rate"] = &interfaceParameters.InterfaceProgramRate
    children["port-shaper-rate"] = &interfaceParameters.PortShaperRate
    return children
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters) GetParentYangName() string { return "input" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate
// Interface Configured Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetFilter() yfilter.YFilter { return interfaceConfigRate.YFilter }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) SetFilter(yf yfilter.YFilter) { interfaceConfigRate.YFilter = yf }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetSegmentPath() string {
    return "interface-config-rate"
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceConfigRate.Value
    leafs["unit"] = interfaceConfigRate.Unit
    return leafs
}

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetYangName() string { return "interface-config-rate" }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) SetParent(parent types.Entity) { interfaceConfigRate.parent = parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetParent() types.Entity { return interfaceConfigRate.parent }

func (interfaceConfigRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceConfigRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate
// Interface Programmed Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetFilter() yfilter.YFilter { return interfaceProgramRate.YFilter }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) SetFilter(yf yfilter.YFilter) { interfaceProgramRate.YFilter = yf }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetSegmentPath() string {
    return "interface-program-rate"
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = interfaceProgramRate.Value
    leafs["unit"] = interfaceProgramRate.Unit
    return leafs
}

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetYangName() string { return "interface-program-rate" }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) SetParent(parent types.Entity) { interfaceProgramRate.parent = parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetParent() types.Entity { return interfaceProgramRate.parent }

func (interfaceProgramRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_InterfaceProgramRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate
// Port Shaper Rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetFilter() yfilter.YFilter { return portShaperRate.YFilter }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) SetFilter(yf yfilter.YFilter) { portShaperRate.YFilter = yf }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetSegmentPath() string {
    return "port-shaper-rate"
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = portShaperRate.Value
    leafs["unit"] = portShaperRate.Unit
    return leafs
}

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetBundleName() string { return "cisco_ios_xr" }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetYangName() string { return "port-shaper-rate" }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) SetParent(parent types.Entity) { portShaperRate.parent = parent }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetParent() types.Entity { return portShaperRate.parent }

func (portShaperRate *PlatformQos_Nodes_Node_Interfaces_Interface_Input_InterfaceParameters_PortShaperRate) GetParentYangName() string { return "interface-parameters" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass
// Skywarp QoS policy class details
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show pclass st. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt.
    QosShowPclassSt []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-pclass-st" { return "QosShowPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt{}
        skywarpQosPolicyClass.QosShowPclassSt = append(skywarpQosPolicyClass.QosShowPclassSt, child)
        return &skywarpQosPolicyClass.QosShowPclassSt[len(skywarpQosPolicyClass.QosShowPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowPclassSt {
        children[skywarpQosPolicyClass.QosShowPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass) GetParentYangName() string { return "input" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt
// qos show pclass st
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt struct {
    parent types.Entity
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

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetFilter() yfilter.YFilter { return qosShowPclassSt.YFilter }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) SetFilter(yf yfilter.YFilter) { qosShowPclassSt.YFilter = yf }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetGoName(yname string) string {
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "queue" { return "Queue" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    if yname == "police" { return "Police" }
    if yname == "marking" { return "Marking" }
    return ""
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetSegmentPath() string {
    return "qos-show-pclass-st"
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &qosShowPclassSt.Queue
    }
    if childYangName == "shape" {
        return &qosShowPclassSt.Shape
    }
    if childYangName == "wfq" {
        return &qosShowPclassSt.Wfq
    }
    if childYangName == "police" {
        return &qosShowPclassSt.Police
    }
    if childYangName == "marking" {
        return &qosShowPclassSt.Marking
    }
    return nil
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &qosShowPclassSt.Queue
    children["shape"] = &qosShowPclassSt.Shape
    children["wfq"] = &qosShowPclassSt.Wfq
    children["police"] = &qosShowPclassSt.Police
    children["marking"] = &qosShowPclassSt.Marking
    return children
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-level"] = qosShowPclassSt.ClassLevel
    leafs["class-name"] = qosShowPclassSt.ClassName
    return leafs
}

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetYangName() string { return "qos-show-pclass-st" }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) SetParent(parent types.Entity) { qosShowPclassSt.parent = parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetParent() types.Entity { return qosShowPclassSt.parent }

func (qosShowPclassSt *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue
// QoS Queue parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Queue ID. The type is interface{} with range: 0..4294967295.
    QueueId interface{}

    // Queue Type. The type is string with length: 0..101.
    QueueType interface{}
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "queue-type" { return "QueueType" }
    return ""
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["queue-type"] = queue.QueueType
    return leafs
}

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Queue) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape
// QoS EA Shaper parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir
// PIR in kbps
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs
// PBS in bytes
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq
// QoS WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Weight. The type is interface{} with range: 0..65535.
    ExcessWeight interface{}

    // Committed Weight.
    CommittedWeight PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight

    // QoS Programmed WFQ parameters.
    ProgrammedWfq PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetGoName(yname string) string {
    if yname == "excess-weight" { return "ExcessWeight" }
    if yname == "committed-weight" { return "CommittedWeight" }
    if yname == "programmed-wfq" { return "ProgrammedWfq" }
    return ""
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "committed-weight" {
        return &wfq.CommittedWeight
    }
    if childYangName == "programmed-wfq" {
        return &wfq.ProgrammedWfq
    }
    return nil
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["committed-weight"] = &wfq.CommittedWeight
    children["programmed-wfq"] = &wfq.ProgrammedWfq
    return children
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-weight"] = wfq.ExcessWeight
    return leafs
}

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight
// Committed Weight
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetFilter() yfilter.YFilter { return committedWeight.YFilter }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetFilter(yf yfilter.YFilter) { committedWeight.YFilter = yf }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetSegmentPath() string {
    return "committed-weight"
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = committedWeight.Value
    leafs["unit"] = committedWeight.Unit
    return leafs
}

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleName() string { return "cisco_ios_xr" }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetYangName() string { return "committed-weight" }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) SetParent(parent types.Entity) { committedWeight.parent = parent }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParent() types.Entity { return committedWeight.parent }

func (committedWeight *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_CommittedWeight) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq
// QoS Programmed WFQ parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetFilter() yfilter.YFilter { return programmedWfq.YFilter }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetFilter(yf yfilter.YFilter) { programmedWfq.YFilter = yf }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetSegmentPath() string {
    return "programmed-wfq"
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &programmedWfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &programmedWfq.SumOfBandwidth
    }
    return nil
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &programmedWfq.Bandwidth
    children["sum-of-bandwidth"] = &programmedWfq.SumOfBandwidth
    return children
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = programmedWfq.ExcessRatio
    return leafs
}

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleName() string { return "cisco_ios_xr" }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetYangName() string { return "programmed-wfq" }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) SetParent(parent types.Entity) { programmedWfq.parent = parent }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParent() types.Entity { return programmedWfq.parent }

func (programmedWfq *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq) GetParentYangName() string { return "wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth
// Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_Bandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Wfq_ProgrammedWfq_SumOfBandwidth) GetParentYangName() string { return "programmed-wfq" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police
// QoS Policer parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police struct {
    parent types.Entity
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

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetGoName(yname string) string {
    if yname == "policer-id" { return "PolicerId" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policer-id"] = police.PolicerId
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetYangName() string { return "police" }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir
// CIR
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs
// CBS
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking
// QoS Mark parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark Only.
    MarkOnly PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly

    // Police conform mark.
    PoliceConform PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform

    // Police exceed mark.
    PoliceExceed PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetFilter() yfilter.YFilter { return marking.YFilter }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetFilter(yf yfilter.YFilter) { marking.YFilter = yf }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetGoName(yname string) string {
    if yname == "mark-only" { return "MarkOnly" }
    if yname == "police-conform" { return "PoliceConform" }
    if yname == "police-exceed" { return "PoliceExceed" }
    return ""
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetSegmentPath() string {
    return "marking"
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-only" {
        return &marking.MarkOnly
    }
    if childYangName == "police-conform" {
        return &marking.PoliceConform
    }
    if childYangName == "police-exceed" {
        return &marking.PoliceExceed
    }
    return nil
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mark-only"] = &marking.MarkOnly
    children["police-conform"] = &marking.PoliceConform
    children["police-exceed"] = &marking.PoliceExceed
    return children
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleName() string { return "cisco_ios_xr" }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetYangName() string { return "marking" }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) SetParent(parent types.Entity) { marking.parent = parent }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParent() types.Entity { return marking.parent }

func (marking *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking) GetParentYangName() string { return "qos-show-pclass-st" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly
// Mark Only
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetFilter() yfilter.YFilter { return markOnly.YFilter }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetFilter(yf yfilter.YFilter) { markOnly.YFilter = yf }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetSegmentPath() string {
    return "mark-only"
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range markOnly.MarkDetail {
            if markOnly.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail{}
        markOnly.MarkDetail = append(markOnly.MarkDetail, child)
        return &markOnly.MarkDetail[len(markOnly.MarkDetail)-1]
    }
    return nil
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range markOnly.MarkDetail {
        children[markOnly.MarkDetail[i].GetSegmentPath()] = &markOnly.MarkDetail[i]
    }
    return children
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = markOnly.ActionType
    return leafs
}

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleName() string { return "cisco_ios_xr" }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetYangName() string { return "mark-only" }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) SetParent(parent types.Entity) { markOnly.parent = parent }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParent() types.Entity { return markOnly.parent }

func (markOnly *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_MarkOnly_MarkDetail) GetParentYangName() string { return "mark-only" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform
// Police conform mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetFilter() yfilter.YFilter { return policeConform.YFilter }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetFilter(yf yfilter.YFilter) { policeConform.YFilter = yf }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetSegmentPath() string {
    return "police-conform"
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeConform.MarkDetail {
            if policeConform.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail{}
        policeConform.MarkDetail = append(policeConform.MarkDetail, child)
        return &policeConform.MarkDetail[len(policeConform.MarkDetail)-1]
    }
    return nil
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeConform.MarkDetail {
        children[policeConform.MarkDetail[i].GetSegmentPath()] = &policeConform.MarkDetail[i]
    }
    return children
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeConform.ActionType
    return leafs
}

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleName() string { return "cisco_ios_xr" }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetYangName() string { return "police-conform" }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) SetParent(parent types.Entity) { policeConform.parent = parent }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParent() types.Entity { return policeConform.parent }

func (policeConform *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceConform_MarkDetail) GetParentYangName() string { return "police-conform" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed
// Police exceed mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Action type. The type is Action.
    ActionType interface{}

    // Mark value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail.
    MarkDetail []PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetFilter() yfilter.YFilter { return policeExceed.YFilter }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetFilter(yf yfilter.YFilter) { policeExceed.YFilter = yf }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark-detail" { return "MarkDetail" }
    return ""
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetSegmentPath() string {
    return "police-exceed"
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark-detail" {
        for _, c := range policeExceed.MarkDetail {
            if policeExceed.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail{}
        policeExceed.MarkDetail = append(policeExceed.MarkDetail, child)
        return &policeExceed.MarkDetail[len(policeExceed.MarkDetail)-1]
    }
    return nil
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range policeExceed.MarkDetail {
        children[policeExceed.MarkDetail[i].GetSegmentPath()] = &policeExceed.MarkDetail[i]
    }
    return children
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = policeExceed.ActionType
    return leafs
}

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleName() string { return "cisco_ios_xr" }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetYangName() string { return "police-exceed" }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) SetParent(parent types.Entity) { policeExceed.parent = parent }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParent() types.Entity { return policeExceed.parent }

func (policeExceed *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed) GetParentYangName() string { return "marking" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail
// Mark value
type PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark value. The type is interface{} with range: 0..255.
    MarkValue interface{}

    // Action opcode. The type is ActionOpcode.
    ActionOpcode interface{}
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetFilter() yfilter.YFilter { return markDetail.YFilter }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetFilter(yf yfilter.YFilter) { markDetail.YFilter = yf }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetGoName(yname string) string {
    if yname == "mark-value" { return "MarkValue" }
    if yname == "action-opcode" { return "ActionOpcode" }
    return ""
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetSegmentPath() string {
    return "mark-detail"
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-value"] = markDetail.MarkValue
    leafs["action-opcode"] = markDetail.ActionOpcode
    return leafs
}

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleName() string { return "cisco_ios_xr" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetYangName() string { return "mark-detail" }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) SetParent(parent types.Entity) { markDetail.parent = parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParent() types.Entity { return markDetail.parent }

func (markDetail *PlatformQos_Nodes_Node_Interfaces_Interface_Input_SkywarpQosPolicyClass_QosShowPclassSt_Marking_PoliceExceed_MarkDetail) GetParentYangName() string { return "police-exceed" }

// PlatformQosEa
// platform qos ea
type PlatformQosEa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of nodes with platform specific QoS-EA configuration.
    Nodes PlatformQosEa_Nodes
}

func (platformQosEa *PlatformQosEa) GetFilter() yfilter.YFilter { return platformQosEa.YFilter }

func (platformQosEa *PlatformQosEa) SetFilter(yf yfilter.YFilter) { platformQosEa.YFilter = yf }

func (platformQosEa *PlatformQosEa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (platformQosEa *PlatformQosEa) GetSegmentPath() string {
    return "Cisco-IOS-XR-skp-qos-oper:platform-qos-ea"
}

func (platformQosEa *PlatformQosEa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &platformQosEa.Nodes
    }
    return nil
}

func (platformQosEa *PlatformQosEa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &platformQosEa.Nodes
    return children
}

func (platformQosEa *PlatformQosEa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (platformQosEa *PlatformQosEa) GetBundleName() string { return "cisco_ios_xr" }

func (platformQosEa *PlatformQosEa) GetYangName() string { return "platform-qos-ea" }

func (platformQosEa *PlatformQosEa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (platformQosEa *PlatformQosEa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (platformQosEa *PlatformQosEa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (platformQosEa *PlatformQosEa) SetParent(parent types.Entity) { platformQosEa.parent = parent }

func (platformQosEa *PlatformQosEa) GetParent() types.Entity { return platformQosEa.parent }

func (platformQosEa *PlatformQosEa) GetParentYangName() string { return "Cisco-IOS-XR-skp-qos-oper" }

// PlatformQosEa_Nodes
// List of nodes with platform specific QoS-EA
// configuration
type PlatformQosEa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node with platform specific QoS-EA configuration. The type is slice of
    // PlatformQosEa_Nodes_Node.
    Node []PlatformQosEa_Nodes_Node
}

func (nodes *PlatformQosEa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *PlatformQosEa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *PlatformQosEa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *PlatformQosEa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *PlatformQosEa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *PlatformQosEa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *PlatformQosEa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *PlatformQosEa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *PlatformQosEa_Nodes) GetYangName() string { return "nodes" }

func (nodes *PlatformQosEa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *PlatformQosEa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *PlatformQosEa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *PlatformQosEa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *PlatformQosEa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *PlatformQosEa_Nodes) GetParentYangName() string { return "platform-qos-ea" }

// PlatformQosEa_Nodes_Node
// Node with platform specific QoS-EA
// configuration
type PlatformQosEa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // QoS-EA list of bundle interfaces.
    BundleInterfaces PlatformQosEa_Nodes_Node_BundleInterfaces

    // QoS-EA list of interfaces.
    Interfaces PlatformQosEa_Nodes_Node_Interfaces
}

func (node *PlatformQosEa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PlatformQosEa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PlatformQosEa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "bundle-interfaces" { return "BundleInterfaces" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (node *PlatformQosEa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *PlatformQosEa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-interfaces" {
        return &node.BundleInterfaces
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    return nil
}

func (node *PlatformQosEa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle-interfaces"] = &node.BundleInterfaces
    children["interfaces"] = &node.Interfaces
    return children
}

func (node *PlatformQosEa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *PlatformQosEa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *PlatformQosEa_Nodes_Node) GetYangName() string { return "node" }

func (node *PlatformQosEa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *PlatformQosEa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *PlatformQosEa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *PlatformQosEa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *PlatformQosEa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *PlatformQosEa_Nodes_Node) GetParentYangName() string { return "nodes" }

// PlatformQosEa_Nodes_Node_BundleInterfaces
// QoS-EA list of bundle interfaces
type PlatformQosEa_Nodes_Node_BundleInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA interface name. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface.
    BundleInterface []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetFilter() yfilter.YFilter { return bundleInterfaces.YFilter }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) SetFilter(yf yfilter.YFilter) { bundleInterfaces.YFilter = yf }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetGoName(yname string) string {
    if yname == "bundle-interface" { return "BundleInterface" }
    return ""
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetSegmentPath() string {
    return "bundle-interfaces"
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-interface" {
        for _, c := range bundleInterfaces.BundleInterface {
            if bundleInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface{}
        bundleInterfaces.BundleInterface = append(bundleInterfaces.BundleInterface, child)
        return &bundleInterfaces.BundleInterface[len(bundleInterfaces.BundleInterface)-1]
    }
    return nil
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bundleInterfaces.BundleInterface {
        children[bundleInterfaces.BundleInterface[i].GetSegmentPath()] = &bundleInterfaces.BundleInterface[i]
    }
    return children
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetYangName() string { return "bundle-interfaces" }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) SetParent(parent types.Entity) { bundleInterfaces.parent = parent }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetParent() types.Entity { return bundleInterfaces.parent }

func (bundleInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces) GetParentYangName() string { return "node" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface
// QoS-EA interface name
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bundle interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // QoS-EA list of member interfaces.
    MemberInterfaces PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetFilter() yfilter.YFilter { return bundleInterface.YFilter }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) SetFilter(yf yfilter.YFilter) { bundleInterface.YFilter = yf }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "member-interfaces" { return "MemberInterfaces" }
    return ""
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetSegmentPath() string {
    return "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member-interfaces" {
        return &bundleInterface.MemberInterfaces
    }
    return nil
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["member-interfaces"] = &bundleInterface.MemberInterfaces
    return children
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bundleInterface.InterfaceName
    return leafs
}

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetYangName() string { return "bundle-interface" }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) SetParent(parent types.Entity) { bundleInterface.parent = parent }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetParent() types.Entity { return bundleInterface.parent }

func (bundleInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface) GetParentYangName() string { return "bundle-interfaces" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
// QoS-EA list of member interfaces
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA interface name. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetFilter() yfilter.YFilter { return memberInterfaces.YFilter }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) SetFilter(yf yfilter.YFilter) { memberInterfaces.YFilter = yf }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetGoName(yname string) string {
    if yname == "member-interface" { return "MemberInterface" }
    return ""
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetSegmentPath() string {
    return "member-interfaces"
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member-interface" {
        for _, c := range memberInterfaces.MemberInterface {
            if memberInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface{}
        memberInterfaces.MemberInterface = append(memberInterfaces.MemberInterface, child)
        return &memberInterfaces.MemberInterface[len(memberInterfaces.MemberInterface)-1]
    }
    return nil
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range memberInterfaces.MemberInterface {
        children[memberInterfaces.MemberInterface[i].GetSegmentPath()] = &memberInterfaces.MemberInterface[i]
    }
    return children
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetYangName() string { return "member-interfaces" }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) SetParent(parent types.Entity) { memberInterfaces.parent = parent }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetParent() types.Entity { return memberInterfaces.parent }

func (memberInterfaces *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces) GetParentYangName() string { return "bundle-interface" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface
// QoS-EA interface name
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Memeber interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // QoS-EA policy direction output.
    BundleOutput PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput

    // QoS-EA policy direction input.
    BundleInput PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetFilter() yfilter.YFilter { return memberInterface.YFilter }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) SetFilter(yf yfilter.YFilter) { memberInterface.YFilter = yf }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "bundle-output" { return "BundleOutput" }
    if yname == "bundle-input" { return "BundleInput" }
    return ""
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetSegmentPath() string {
    return "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-output" {
        return &memberInterface.BundleOutput
    }
    if childYangName == "bundle-input" {
        return &memberInterface.BundleInput
    }
    return nil
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle-output"] = &memberInterface.BundleOutput
    children["bundle-input"] = &memberInterface.BundleInput
    return children
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = memberInterface.InterfaceName
    return leafs
}

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetBundleName() string { return "cisco_ios_xr" }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetYangName() string { return "member-interface" }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) SetParent(parent types.Entity) { memberInterface.parent = parent }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetParent() types.Entity { return memberInterface.parent }

func (memberInterface *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetParentYangName() string { return "member-interfaces" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput
// QoS-EA policy direction output
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetFilter() yfilter.YFilter { return bundleOutput.YFilter }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) SetFilter(yf yfilter.YFilter) { bundleOutput.YFilter = yf }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetSegmentPath() string {
    return "bundle-output"
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &bundleOutput.Details
    }
    return nil
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &bundleOutput.Details
    return children
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetBundleName() string { return "cisco_ios_xr" }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetYangName() string { return "bundle-output" }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) SetParent(parent types.Entity) { bundleOutput.parent = parent }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetParent() types.Entity { return bundleOutput.parent }

func (bundleOutput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput) GetParentYangName() string { return "member-interface" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetSegmentPath() string {
    return "details"
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &details.Header
    }
    if childYangName == "interface-parameters" {
        return &details.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &details.SkywarpQosPolicyClass
    }
    return nil
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &details.Header
    children["interface-parameters"] = &details.InterfaceParameters
    children["skywarp-qos-policy-class"] = &details.SkywarpQosPolicyClass
    return children
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetYangName() string { return "details" }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetParent() types.Entity { return details.parent }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details) GetParentYangName() string { return "bundle-output" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header struct {
    parent types.Entity
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

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetYangName() string { return "header" }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_Header) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters struct {
    parent types.Entity
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

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "hierarchical-depth" { return "HierarchicalDepth" }
    if yname == "interface-type" { return "InterfaceType" }
    if yname == "interface-rate" { return "InterfaceRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "under-line-interface-handle" { return "UnderLineInterfaceHandle" }
    if yname == "bundle-id" { return "BundleId" }
    if yname == "uidb-index" { return "UidbIndex" }
    if yname == "qos-interface-handle" { return "QosInterfaceHandle" }
    if yname == "port" { return "Port" }
    if yname == "policy-map-id" { return "PolicyMapId" }
    return ""
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = interfaceParameters.PolicyName
    leafs["hierarchical-depth"] = interfaceParameters.HierarchicalDepth
    leafs["interface-type"] = interfaceParameters.InterfaceType
    leafs["interface-rate"] = interfaceParameters.InterfaceRate
    leafs["port-shaper-rate"] = interfaceParameters.PortShaperRate
    leafs["interface-handle"] = interfaceParameters.InterfaceHandle
    leafs["under-line-interface-handle"] = interfaceParameters.UnderLineInterfaceHandle
    leafs["bundle-id"] = interfaceParameters.BundleId
    leafs["uidb-index"] = interfaceParameters.UidbIndex
    leafs["qos-interface-handle"] = interfaceParameters.QosInterfaceHandle
    leafs["port"] = interfaceParameters.Port
    leafs["policy-map-id"] = interfaceParameters.PolicyMapId
    return leafs
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_InterfaceParameters) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-ea-pclass-st" { return "QosShowEaPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-ea-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowEaPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt{}
        skywarpQosPolicyClass.QosShowEaPclassSt = append(skywarpQosPolicyClass.QosShowEaPclassSt, child)
        return &skywarpQosPolicyClass.QosShowEaPclassSt[len(skywarpQosPolicyClass.QosShowEaPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        children[skywarpQosPolicyClass.QosShowEaPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowEaPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    parent types.Entity
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

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetFilter() yfilter.YFilter { return qosShowEaPclassSt.YFilter }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetFilter(yf yfilter.YFilter) { qosShowEaPclassSt.YFilter = yf }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "node-flags" { return "NodeFlags" }
    if yname == "stats-flags" { return "StatsFlags" }
    if yname == "config" { return "Config" }
    if yname == "result" { return "Result" }
    return ""
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetSegmentPath() string {
    return "qos-show-ea-pclass-st"
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &qosShowEaPclassSt.Config
    }
    if childYangName == "result" {
        return &qosShowEaPclassSt.Result
    }
    return nil
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &qosShowEaPclassSt.Config
    children["result"] = &qosShowEaPclassSt.Result
    return children
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = qosShowEaPclassSt.Index
    leafs["class-level"] = qosShowEaPclassSt.ClassLevel
    leafs["class-name"] = qosShowEaPclassSt.ClassName
    leafs["policy-name"] = qosShowEaPclassSt.PolicyName
    leafs["node-flags"] = qosShowEaPclassSt.NodeFlags
    leafs["stats-flags"] = qosShowEaPclassSt.StatsFlags
    return leafs
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetYangName() string { return "qos-show-ea-pclass-st" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetParent(parent types.Entity) { qosShowEaPclassSt.parent = parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParent() types.Entity { return qosShowEaPclassSt.parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    parent types.Entity
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

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetGoName(yname string) string {
    if yname == "node-config" { return "NodeConfig" }
    if yname == "police" { return "Police" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    return ""
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetSegmentPath() string {
    return "config"
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police" {
        return &config.Police
    }
    if childYangName == "shape" {
        return &config.Shape
    }
    if childYangName == "wfq" {
        return &config.Wfq
    }
    return nil
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["police"] = &config.Police
    children["shape"] = &config.Shape
    children["wfq"] = &config.Wfq
    return children
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-config"] = config.NodeConfig
    return leafs
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetYangName() string { return "config" }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParent() types.Entity { return config.parent }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetGoName(yname string) string {
    if yname == "color-aware" { return "ColorAware" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["color-aware"] = police.ColorAware
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &wfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &wfq.SumOfBandwidth
    }
    return nil
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &wfq.Bandwidth
    children["sum-of-bandwidth"] = &wfq.SumOfBandwidth
    return children
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = wfq.ExcessRatio
    return leafs
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetFilter() yfilter.YFilter { return result.YFilter }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetFilter(yf yfilter.YFilter) { result.YFilter = yf }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetGoName(yname string) string {
    if yname == "stats-id" { return "StatsId" }
    if yname == "queue" { return "Queue" }
    if yname == "police" { return "Police" }
    return ""
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetSegmentPath() string {
    return "result"
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &result.Queue
    }
    if childYangName == "police" {
        return &result.Police
    }
    return nil
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &result.Queue
    children["police"] = &result.Police
    return children
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-id"] = result.StatsId
    return leafs
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleName() string { return "cisco_ios_xr" }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetYangName() string { return "result" }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetParent(parent types.Entity) { result.parent = parent }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParent() types.Entity { return result.parent }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    parent types.Entity
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

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "commit-tx" { return "CommitTx" }
    if yname == "excess-tx" { return "ExcessTx" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["commit-tx"] = queue.CommitTx
    leafs["excess-tx"] = queue.ExcessTx
    leafs["drop"] = queue.Drop
    return leafs
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetGoName(yname string) string {
    if yname == "token-bucket-id" { return "TokenBucketId" }
    if yname == "conform" { return "Conform" }
    if yname == "exceed" { return "Exceed" }
    if yname == "violate" { return "Violate" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["token-bucket-id"] = police.TokenBucketId
    leafs["conform"] = police.Conform
    leafs["exceed"] = police.Exceed
    leafs["violate"] = police.Violate
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleOutput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput
// QoS-EA policy direction input
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetFilter() yfilter.YFilter { return bundleInput.YFilter }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) SetFilter(yf yfilter.YFilter) { bundleInput.YFilter = yf }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetSegmentPath() string {
    return "bundle-input"
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &bundleInput.Details
    }
    return nil
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &bundleInput.Details
    return children
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetBundleName() string { return "cisco_ios_xr" }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetYangName() string { return "bundle-input" }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) SetParent(parent types.Entity) { bundleInput.parent = parent }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetParent() types.Entity { return bundleInput.parent }

func (bundleInput *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput) GetParentYangName() string { return "member-interface" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetSegmentPath() string {
    return "details"
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &details.Header
    }
    if childYangName == "interface-parameters" {
        return &details.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &details.SkywarpQosPolicyClass
    }
    return nil
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &details.Header
    children["interface-parameters"] = &details.InterfaceParameters
    children["skywarp-qos-policy-class"] = &details.SkywarpQosPolicyClass
    return children
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetYangName() string { return "details" }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetParent() types.Entity { return details.parent }

func (details *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details) GetParentYangName() string { return "bundle-input" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header struct {
    parent types.Entity
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

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetYangName() string { return "header" }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_Header) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters struct {
    parent types.Entity
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

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "hierarchical-depth" { return "HierarchicalDepth" }
    if yname == "interface-type" { return "InterfaceType" }
    if yname == "interface-rate" { return "InterfaceRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "under-line-interface-handle" { return "UnderLineInterfaceHandle" }
    if yname == "bundle-id" { return "BundleId" }
    if yname == "uidb-index" { return "UidbIndex" }
    if yname == "qos-interface-handle" { return "QosInterfaceHandle" }
    if yname == "port" { return "Port" }
    if yname == "policy-map-id" { return "PolicyMapId" }
    return ""
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = interfaceParameters.PolicyName
    leafs["hierarchical-depth"] = interfaceParameters.HierarchicalDepth
    leafs["interface-type"] = interfaceParameters.InterfaceType
    leafs["interface-rate"] = interfaceParameters.InterfaceRate
    leafs["port-shaper-rate"] = interfaceParameters.PortShaperRate
    leafs["interface-handle"] = interfaceParameters.InterfaceHandle
    leafs["under-line-interface-handle"] = interfaceParameters.UnderLineInterfaceHandle
    leafs["bundle-id"] = interfaceParameters.BundleId
    leafs["uidb-index"] = interfaceParameters.UidbIndex
    leafs["qos-interface-handle"] = interfaceParameters.QosInterfaceHandle
    leafs["port"] = interfaceParameters.Port
    leafs["policy-map-id"] = interfaceParameters.PolicyMapId
    return leafs
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_InterfaceParameters) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-ea-pclass-st" { return "QosShowEaPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-ea-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowEaPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt{}
        skywarpQosPolicyClass.QosShowEaPclassSt = append(skywarpQosPolicyClass.QosShowEaPclassSt, child)
        return &skywarpQosPolicyClass.QosShowEaPclassSt[len(skywarpQosPolicyClass.QosShowEaPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        children[skywarpQosPolicyClass.QosShowEaPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowEaPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    parent types.Entity
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

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetFilter() yfilter.YFilter { return qosShowEaPclassSt.YFilter }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetFilter(yf yfilter.YFilter) { qosShowEaPclassSt.YFilter = yf }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "node-flags" { return "NodeFlags" }
    if yname == "stats-flags" { return "StatsFlags" }
    if yname == "config" { return "Config" }
    if yname == "result" { return "Result" }
    return ""
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetSegmentPath() string {
    return "qos-show-ea-pclass-st"
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &qosShowEaPclassSt.Config
    }
    if childYangName == "result" {
        return &qosShowEaPclassSt.Result
    }
    return nil
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &qosShowEaPclassSt.Config
    children["result"] = &qosShowEaPclassSt.Result
    return children
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = qosShowEaPclassSt.Index
    leafs["class-level"] = qosShowEaPclassSt.ClassLevel
    leafs["class-name"] = qosShowEaPclassSt.ClassName
    leafs["policy-name"] = qosShowEaPclassSt.PolicyName
    leafs["node-flags"] = qosShowEaPclassSt.NodeFlags
    leafs["stats-flags"] = qosShowEaPclassSt.StatsFlags
    return leafs
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetYangName() string { return "qos-show-ea-pclass-st" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetParent(parent types.Entity) { qosShowEaPclassSt.parent = parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParent() types.Entity { return qosShowEaPclassSt.parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    parent types.Entity
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

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetGoName(yname string) string {
    if yname == "node-config" { return "NodeConfig" }
    if yname == "police" { return "Police" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    return ""
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetSegmentPath() string {
    return "config"
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police" {
        return &config.Police
    }
    if childYangName == "shape" {
        return &config.Shape
    }
    if childYangName == "wfq" {
        return &config.Wfq
    }
    return nil
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["police"] = &config.Police
    children["shape"] = &config.Shape
    children["wfq"] = &config.Wfq
    return children
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-config"] = config.NodeConfig
    return leafs
}

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetYangName() string { return "config" }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParent() types.Entity { return config.parent }

func (config *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetGoName(yname string) string {
    if yname == "color-aware" { return "ColorAware" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["color-aware"] = police.ColorAware
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &wfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &wfq.SumOfBandwidth
    }
    return nil
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &wfq.Bandwidth
    children["sum-of-bandwidth"] = &wfq.SumOfBandwidth
    return children
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = wfq.ExcessRatio
    return leafs
}

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetFilter() yfilter.YFilter { return result.YFilter }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetFilter(yf yfilter.YFilter) { result.YFilter = yf }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetGoName(yname string) string {
    if yname == "stats-id" { return "StatsId" }
    if yname == "queue" { return "Queue" }
    if yname == "police" { return "Police" }
    return ""
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetSegmentPath() string {
    return "result"
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &result.Queue
    }
    if childYangName == "police" {
        return &result.Police
    }
    return nil
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &result.Queue
    children["police"] = &result.Police
    return children
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-id"] = result.StatsId
    return leafs
}

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleName() string { return "cisco_ios_xr" }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetYangName() string { return "result" }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetParent(parent types.Entity) { result.parent = parent }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParent() types.Entity { return result.parent }

func (result *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    parent types.Entity
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

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "commit-tx" { return "CommitTx" }
    if yname == "excess-tx" { return "ExcessTx" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["commit-tx"] = queue.CommitTx
    leafs["excess-tx"] = queue.ExcessTx
    leafs["drop"] = queue.Drop
    return leafs
}

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetGoName(yname string) string {
    if yname == "token-bucket-id" { return "TokenBucketId" }
    if yname == "conform" { return "Conform" }
    if yname == "exceed" { return "Exceed" }
    if yname == "violate" { return "Violate" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["token-bucket-id"] = police.TokenBucketId
    leafs["conform"] = police.Conform
    leafs["exceed"] = police.Exceed
    leafs["violate"] = police.Violate
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_BundleInput_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_Interfaces
// QoS-EA list of interfaces
type PlatformQosEa_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA interface name. The type is slice of
    // PlatformQosEa_Nodes_Node_Interfaces_Interface.
    Interface []PlatformQosEa_Nodes_Node_Interfaces_Interface
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *PlatformQosEa_Nodes_Node_Interfaces) GetParentYangName() string { return "node" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface
// QoS-EA interface name
type PlatformQosEa_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // QoS-EA policy direction egress.
    Output PlatformQosEa_Nodes_Node_Interfaces_Interface_Output

    // QoS-EA policy direction ingress.
    Input PlatformQosEa_Nodes_Node_Interfaces_Interface_Input
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "output" { return "Output" }
    if yname == "input" { return "Input" }
    return ""
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        return &self.Output
    }
    if childYangName == "input" {
        return &self.Input
    }
    return nil
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["output"] = &self.Output
    children["input"] = &self.Input
    return children
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *PlatformQosEa_Nodes_Node_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output
// QoS-EA policy direction egress
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetSegmentPath() string {
    return "output"
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &output.Details
    }
    return nil
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &output.Details
    return children
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetYangName() string { return "output" }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetParent() types.Entity { return output.parent }

func (output *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output) GetParentYangName() string { return "interface" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetSegmentPath() string {
    return "details"
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &details.Header
    }
    if childYangName == "interface-parameters" {
        return &details.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &details.SkywarpQosPolicyClass
    }
    return nil
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &details.Header
    children["interface-parameters"] = &details.InterfaceParameters
    children["skywarp-qos-policy-class"] = &details.SkywarpQosPolicyClass
    return children
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetYangName() string { return "details" }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetParent() types.Entity { return details.parent }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details) GetParentYangName() string { return "output" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header struct {
    parent types.Entity
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

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetYangName() string { return "header" }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_Header) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters struct {
    parent types.Entity
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

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "hierarchical-depth" { return "HierarchicalDepth" }
    if yname == "interface-type" { return "InterfaceType" }
    if yname == "interface-rate" { return "InterfaceRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "under-line-interface-handle" { return "UnderLineInterfaceHandle" }
    if yname == "bundle-id" { return "BundleId" }
    if yname == "uidb-index" { return "UidbIndex" }
    if yname == "qos-interface-handle" { return "QosInterfaceHandle" }
    if yname == "port" { return "Port" }
    if yname == "policy-map-id" { return "PolicyMapId" }
    return ""
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = interfaceParameters.PolicyName
    leafs["hierarchical-depth"] = interfaceParameters.HierarchicalDepth
    leafs["interface-type"] = interfaceParameters.InterfaceType
    leafs["interface-rate"] = interfaceParameters.InterfaceRate
    leafs["port-shaper-rate"] = interfaceParameters.PortShaperRate
    leafs["interface-handle"] = interfaceParameters.InterfaceHandle
    leafs["under-line-interface-handle"] = interfaceParameters.UnderLineInterfaceHandle
    leafs["bundle-id"] = interfaceParameters.BundleId
    leafs["uidb-index"] = interfaceParameters.UidbIndex
    leafs["qos-interface-handle"] = interfaceParameters.QosInterfaceHandle
    leafs["port"] = interfaceParameters.Port
    leafs["policy-map-id"] = interfaceParameters.PolicyMapId
    return leafs
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_InterfaceParameters) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-ea-pclass-st" { return "QosShowEaPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-ea-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowEaPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt{}
        skywarpQosPolicyClass.QosShowEaPclassSt = append(skywarpQosPolicyClass.QosShowEaPclassSt, child)
        return &skywarpQosPolicyClass.QosShowEaPclassSt[len(skywarpQosPolicyClass.QosShowEaPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        children[skywarpQosPolicyClass.QosShowEaPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowEaPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    parent types.Entity
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

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetFilter() yfilter.YFilter { return qosShowEaPclassSt.YFilter }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetFilter(yf yfilter.YFilter) { qosShowEaPclassSt.YFilter = yf }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "node-flags" { return "NodeFlags" }
    if yname == "stats-flags" { return "StatsFlags" }
    if yname == "config" { return "Config" }
    if yname == "result" { return "Result" }
    return ""
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetSegmentPath() string {
    return "qos-show-ea-pclass-st"
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &qosShowEaPclassSt.Config
    }
    if childYangName == "result" {
        return &qosShowEaPclassSt.Result
    }
    return nil
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &qosShowEaPclassSt.Config
    children["result"] = &qosShowEaPclassSt.Result
    return children
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = qosShowEaPclassSt.Index
    leafs["class-level"] = qosShowEaPclassSt.ClassLevel
    leafs["class-name"] = qosShowEaPclassSt.ClassName
    leafs["policy-name"] = qosShowEaPclassSt.PolicyName
    leafs["node-flags"] = qosShowEaPclassSt.NodeFlags
    leafs["stats-flags"] = qosShowEaPclassSt.StatsFlags
    return leafs
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetYangName() string { return "qos-show-ea-pclass-st" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetParent(parent types.Entity) { qosShowEaPclassSt.parent = parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParent() types.Entity { return qosShowEaPclassSt.parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    parent types.Entity
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

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetGoName(yname string) string {
    if yname == "node-config" { return "NodeConfig" }
    if yname == "police" { return "Police" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    return ""
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetSegmentPath() string {
    return "config"
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police" {
        return &config.Police
    }
    if childYangName == "shape" {
        return &config.Shape
    }
    if childYangName == "wfq" {
        return &config.Wfq
    }
    return nil
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["police"] = &config.Police
    children["shape"] = &config.Shape
    children["wfq"] = &config.Wfq
    return children
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-config"] = config.NodeConfig
    return leafs
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetYangName() string { return "config" }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParent() types.Entity { return config.parent }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetGoName(yname string) string {
    if yname == "color-aware" { return "ColorAware" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["color-aware"] = police.ColorAware
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &wfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &wfq.SumOfBandwidth
    }
    return nil
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &wfq.Bandwidth
    children["sum-of-bandwidth"] = &wfq.SumOfBandwidth
    return children
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = wfq.ExcessRatio
    return leafs
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetFilter() yfilter.YFilter { return result.YFilter }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetFilter(yf yfilter.YFilter) { result.YFilter = yf }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetGoName(yname string) string {
    if yname == "stats-id" { return "StatsId" }
    if yname == "queue" { return "Queue" }
    if yname == "police" { return "Police" }
    return ""
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetSegmentPath() string {
    return "result"
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &result.Queue
    }
    if childYangName == "police" {
        return &result.Police
    }
    return nil
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &result.Queue
    children["police"] = &result.Police
    return children
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-id"] = result.StatsId
    return leafs
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleName() string { return "cisco_ios_xr" }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetYangName() string { return "result" }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetParent(parent types.Entity) { result.parent = parent }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParent() types.Entity { return result.parent }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    parent types.Entity
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

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "commit-tx" { return "CommitTx" }
    if yname == "excess-tx" { return "ExcessTx" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["commit-tx"] = queue.CommitTx
    leafs["excess-tx"] = queue.ExcessTx
    leafs["drop"] = queue.Drop
    return leafs
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetGoName(yname string) string {
    if yname == "token-bucket-id" { return "TokenBucketId" }
    if yname == "conform" { return "Conform" }
    if yname == "exceed" { return "Exceed" }
    if yname == "violate" { return "Violate" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["token-bucket-id"] = police.TokenBucketId
    leafs["conform"] = police.Conform
    leafs["exceed"] = police.Exceed
    leafs["violate"] = police.Violate
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Output_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input
// QoS-EA policy direction ingress
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS-EA policy details.
    Details PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetGoName(yname string) string {
    if yname == "details" { return "Details" }
    return ""
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetSegmentPath() string {
    return "input"
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &input.Details
    }
    return nil
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &input.Details
    return children
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetBundleName() string { return "cisco_ios_xr" }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetYangName() string { return "input" }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetParent() types.Entity { return input.parent }

func (input *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input) GetParentYangName() string { return "interface" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details
// QoS-EA policy details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS EA policy header.
    Header PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header

    // QoS EA Interface Parameters.
    InterfaceParameters PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters

    // Skywarp QoS EA policy class details.
    SkywarpQosPolicyClass PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetGoName(yname string) string {
    if yname == "header" { return "Header" }
    if yname == "interface-parameters" { return "InterfaceParameters" }
    if yname == "skywarp-qos-policy-class" { return "SkywarpQosPolicyClass" }
    return ""
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetSegmentPath() string {
    return "details"
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "header" {
        return &details.Header
    }
    if childYangName == "interface-parameters" {
        return &details.InterfaceParameters
    }
    if childYangName == "skywarp-qos-policy-class" {
        return &details.SkywarpQosPolicyClass
    }
    return nil
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["header"] = &details.Header
    children["interface-parameters"] = &details.InterfaceParameters
    children["skywarp-qos-policy-class"] = &details.SkywarpQosPolicyClass
    return children
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetBundleName() string { return "cisco_ios_xr" }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetYangName() string { return "details" }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetParent() types.Entity { return details.parent }

func (details *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details) GetParentYangName() string { return "input" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header
// QoS EA policy header
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header struct {
    parent types.Entity
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

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetFilter() yfilter.YFilter { return header.YFilter }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) SetFilter(yf yfilter.YFilter) { header.YFilter = yf }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "direction" { return "Direction" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetSegmentPath() string {
    return "header"
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = header.InterfaceName
    leafs["policy-name"] = header.PolicyName
    leafs["direction"] = header.Direction
    leafs["classes"] = header.Classes
    return leafs
}

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetBundleName() string { return "cisco_ios_xr" }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetYangName() string { return "header" }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) SetParent(parent types.Entity) { header.parent = parent }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetParent() types.Entity { return header.parent }

func (header *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_Header) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters
// QoS EA Interface Parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters struct {
    parent types.Entity
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

    // Interface Handle. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceHandle interface{}

    // UnderLineInterface Handle. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetFilter() yfilter.YFilter { return interfaceParameters.YFilter }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) SetFilter(yf yfilter.YFilter) { interfaceParameters.YFilter = yf }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetGoName(yname string) string {
    if yname == "policy-name" { return "PolicyName" }
    if yname == "hierarchical-depth" { return "HierarchicalDepth" }
    if yname == "interface-type" { return "InterfaceType" }
    if yname == "interface-rate" { return "InterfaceRate" }
    if yname == "port-shaper-rate" { return "PortShaperRate" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "under-line-interface-handle" { return "UnderLineInterfaceHandle" }
    if yname == "bundle-id" { return "BundleId" }
    if yname == "uidb-index" { return "UidbIndex" }
    if yname == "qos-interface-handle" { return "QosInterfaceHandle" }
    if yname == "port" { return "Port" }
    if yname == "policy-map-id" { return "PolicyMapId" }
    return ""
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetSegmentPath() string {
    return "interface-parameters"
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-name"] = interfaceParameters.PolicyName
    leafs["hierarchical-depth"] = interfaceParameters.HierarchicalDepth
    leafs["interface-type"] = interfaceParameters.InterfaceType
    leafs["interface-rate"] = interfaceParameters.InterfaceRate
    leafs["port-shaper-rate"] = interfaceParameters.PortShaperRate
    leafs["interface-handle"] = interfaceParameters.InterfaceHandle
    leafs["under-line-interface-handle"] = interfaceParameters.UnderLineInterfaceHandle
    leafs["bundle-id"] = interfaceParameters.BundleId
    leafs["uidb-index"] = interfaceParameters.UidbIndex
    leafs["qos-interface-handle"] = interfaceParameters.QosInterfaceHandle
    leafs["port"] = interfaceParameters.Port
    leafs["policy-map-id"] = interfaceParameters.PolicyMapId
    return leafs
}

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetYangName() string { return "interface-parameters" }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) SetParent(parent types.Entity) { interfaceParameters.parent = parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetParent() types.Entity { return interfaceParameters.parent }

func (interfaceParameters *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_InterfaceParameters) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass
// Skywarp QoS EA policy class details
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // qos show ea pclass st. The type is slice of
    // PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt.
    QosShowEaPclassSt []PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetFilter() yfilter.YFilter { return skywarpQosPolicyClass.YFilter }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) SetFilter(yf yfilter.YFilter) { skywarpQosPolicyClass.YFilter = yf }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetGoName(yname string) string {
    if yname == "qos-show-ea-pclass-st" { return "QosShowEaPclassSt" }
    return ""
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetSegmentPath() string {
    return "skywarp-qos-policy-class"
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "qos-show-ea-pclass-st" {
        for _, c := range skywarpQosPolicyClass.QosShowEaPclassSt {
            if skywarpQosPolicyClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt{}
        skywarpQosPolicyClass.QosShowEaPclassSt = append(skywarpQosPolicyClass.QosShowEaPclassSt, child)
        return &skywarpQosPolicyClass.QosShowEaPclassSt[len(skywarpQosPolicyClass.QosShowEaPclassSt)-1]
    }
    return nil
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range skywarpQosPolicyClass.QosShowEaPclassSt {
        children[skywarpQosPolicyClass.QosShowEaPclassSt[i].GetSegmentPath()] = &skywarpQosPolicyClass.QosShowEaPclassSt[i]
    }
    return children
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetBundleName() string { return "cisco_ios_xr" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetYangName() string { return "skywarp-qos-policy-class" }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) SetParent(parent types.Entity) { skywarpQosPolicyClass.parent = parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetParent() types.Entity { return skywarpQosPolicyClass.parent }

func (skywarpQosPolicyClass *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass) GetParentYangName() string { return "details" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt
// qos show ea pclass st
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt struct {
    parent types.Entity
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

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetFilter() yfilter.YFilter { return qosShowEaPclassSt.YFilter }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetFilter(yf yfilter.YFilter) { qosShowEaPclassSt.YFilter = yf }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetGoName(yname string) string {
    if yname == "index" { return "Index" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "class-name" { return "ClassName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "node-flags" { return "NodeFlags" }
    if yname == "stats-flags" { return "StatsFlags" }
    if yname == "config" { return "Config" }
    if yname == "result" { return "Result" }
    return ""
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetSegmentPath() string {
    return "qos-show-ea-pclass-st"
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config" {
        return &qosShowEaPclassSt.Config
    }
    if childYangName == "result" {
        return &qosShowEaPclassSt.Result
    }
    return nil
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config"] = &qosShowEaPclassSt.Config
    children["result"] = &qosShowEaPclassSt.Result
    return children
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["index"] = qosShowEaPclassSt.Index
    leafs["class-level"] = qosShowEaPclassSt.ClassLevel
    leafs["class-name"] = qosShowEaPclassSt.ClassName
    leafs["policy-name"] = qosShowEaPclassSt.PolicyName
    leafs["node-flags"] = qosShowEaPclassSt.NodeFlags
    leafs["stats-flags"] = qosShowEaPclassSt.StatsFlags
    return leafs
}

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleName() string { return "cisco_ios_xr" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetYangName() string { return "qos-show-ea-pclass-st" }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) SetParent(parent types.Entity) { qosShowEaPclassSt.parent = parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParent() types.Entity { return qosShowEaPclassSt.parent }

func (qosShowEaPclassSt *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt) GetParentYangName() string { return "skywarp-qos-policy-class" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config
// QoS EA Class Configuration
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config struct {
    parent types.Entity
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

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetFilter() yfilter.YFilter { return config.YFilter }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetFilter(yf yfilter.YFilter) { config.YFilter = yf }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetGoName(yname string) string {
    if yname == "node-config" { return "NodeConfig" }
    if yname == "police" { return "Police" }
    if yname == "shape" { return "Shape" }
    if yname == "wfq" { return "Wfq" }
    return ""
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetSegmentPath() string {
    return "config"
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "police" {
        return &config.Police
    }
    if childYangName == "shape" {
        return &config.Shape
    }
    if childYangName == "wfq" {
        return &config.Wfq
    }
    return nil
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["police"] = &config.Police
    children["shape"] = &config.Shape
    children["wfq"] = &config.Wfq
    return children
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-config"] = config.NodeConfig
    return leafs
}

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleName() string { return "cisco_ios_xr" }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetYangName() string { return "config" }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) SetParent(parent types.Entity) { config.parent = parent }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParent() types.Entity { return config.parent }

func (config *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police
// QoS EA Policer parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetGoName(yname string) string {
    if yname == "color-aware" { return "ColorAware" }
    if yname == "policer-type" { return "PolicerType" }
    if yname == "cir" { return "Cir" }
    if yname == "cbs" { return "Cbs" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cir" {
        return &police.Cir
    }
    if childYangName == "cbs" {
        return &police.Cbs
    }
    return nil
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cir"] = &police.Cir
    children["cbs"] = &police.Cbs
    return children
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["color-aware"] = police.ColorAware
    leafs["policer-type"] = police.PolicerType
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir
// CIR
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetFilter() yfilter.YFilter { return cir.YFilter }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetFilter(yf yfilter.YFilter) { cir.YFilter = yf }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetSegmentPath() string {
    return "cir"
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cir.Value
    leafs["unit"] = cir.Unit
    return leafs
}

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleName() string { return "cisco_ios_xr" }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetYangName() string { return "cir" }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) SetParent(parent types.Entity) { cir.parent = parent }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParent() types.Entity { return cir.parent }

func (cir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cir) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs
// CBS
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetFilter() yfilter.YFilter { return cbs.YFilter }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetFilter(yf yfilter.YFilter) { cbs.YFilter = yf }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetSegmentPath() string {
    return "cbs"
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = cbs.Value
    leafs["unit"] = cbs.Unit
    return leafs
}

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleName() string { return "cisco_ios_xr" }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetYangName() string { return "cbs" }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) SetParent(parent types.Entity) { cbs.parent = parent }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParent() types.Entity { return cbs.parent }

func (cbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Police_Cbs) GetParentYangName() string { return "police" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape
// QoS EA Shaper parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PIR in kbps.
    Pir PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir

    // PBS in bytes.
    Pbs PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetFilter() yfilter.YFilter { return shape.YFilter }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetFilter(yf yfilter.YFilter) { shape.YFilter = yf }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetGoName(yname string) string {
    if yname == "pir" { return "Pir" }
    if yname == "pbs" { return "Pbs" }
    return ""
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetSegmentPath() string {
    return "shape"
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pir" {
        return &shape.Pir
    }
    if childYangName == "pbs" {
        return &shape.Pbs
    }
    return nil
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pir"] = &shape.Pir
    children["pbs"] = &shape.Pbs
    return children
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleName() string { return "cisco_ios_xr" }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetYangName() string { return "shape" }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) SetParent(parent types.Entity) { shape.parent = parent }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParent() types.Entity { return shape.parent }

func (shape *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir
// PIR in kbps
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetFilter() yfilter.YFilter { return pir.YFilter }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetFilter(yf yfilter.YFilter) { pir.YFilter = yf }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetSegmentPath() string {
    return "pir"
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pir.Value
    leafs["unit"] = pir.Unit
    return leafs
}

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleName() string { return "cisco_ios_xr" }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetYangName() string { return "pir" }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) SetParent(parent types.Entity) { pir.parent = parent }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParent() types.Entity { return pir.parent }

func (pir *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pir) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs
// PBS in bytes
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetFilter() yfilter.YFilter { return pbs.YFilter }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetFilter(yf yfilter.YFilter) { pbs.YFilter = yf }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetSegmentPath() string {
    return "pbs"
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = pbs.Value
    leafs["unit"] = pbs.Unit
    return leafs
}

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleName() string { return "cisco_ios_xr" }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetYangName() string { return "pbs" }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) SetParent(parent types.Entity) { pbs.parent = parent }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParent() types.Entity { return pbs.parent }

func (pbs *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Shape_Pbs) GetParentYangName() string { return "shape" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq
// QoS EA WFQ parameters
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Excess Ratio. The type is interface{} with range: 0..65535.
    ExcessRatio interface{}

    // Bandwidth.
    Bandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth

    // Sum of Bandwidth.
    SumOfBandwidth PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetFilter() yfilter.YFilter { return wfq.YFilter }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetFilter(yf yfilter.YFilter) { wfq.YFilter = yf }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetGoName(yname string) string {
    if yname == "excess-ratio" { return "ExcessRatio" }
    if yname == "bandwidth" { return "Bandwidth" }
    if yname == "sum-of-bandwidth" { return "SumOfBandwidth" }
    return ""
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetSegmentPath() string {
    return "wfq"
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bandwidth" {
        return &wfq.Bandwidth
    }
    if childYangName == "sum-of-bandwidth" {
        return &wfq.SumOfBandwidth
    }
    return nil
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bandwidth"] = &wfq.Bandwidth
    children["sum-of-bandwidth"] = &wfq.SumOfBandwidth
    return children
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["excess-ratio"] = wfq.ExcessRatio
    return leafs
}

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleName() string { return "cisco_ios_xr" }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetYangName() string { return "wfq" }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) SetParent(parent types.Entity) { wfq.parent = parent }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParent() types.Entity { return wfq.parent }

func (wfq *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq) GetParentYangName() string { return "config" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth
// Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetFilter() yfilter.YFilter { return bandwidth.YFilter }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetFilter(yf yfilter.YFilter) { bandwidth.YFilter = yf }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetSegmentPath() string {
    return "bandwidth"
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = bandwidth.Value
    leafs["unit"] = bandwidth.Unit
    return leafs
}

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetYangName() string { return "bandwidth" }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) SetParent(parent types.Entity) { bandwidth.parent = parent }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParent() types.Entity { return bandwidth.parent }

func (bandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_Bandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth
// Sum of Bandwidth
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Config value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Config unit. The type is QosUnit.
    Unit interface{}
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetFilter() yfilter.YFilter { return sumOfBandwidth.YFilter }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetFilter(yf yfilter.YFilter) { sumOfBandwidth.YFilter = yf }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "unit" { return "Unit" }
    return ""
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetSegmentPath() string {
    return "sum-of-bandwidth"
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = sumOfBandwidth.Value
    leafs["unit"] = sumOfBandwidth.Unit
    return leafs
}

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleName() string { return "cisco_ios_xr" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetYangName() string { return "sum-of-bandwidth" }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) SetParent(parent types.Entity) { sumOfBandwidth.parent = parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParent() types.Entity { return sumOfBandwidth.parent }

func (sumOfBandwidth *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Config_Wfq_SumOfBandwidth) GetParentYangName() string { return "wfq" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result
// QoS EA Class Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stats ID. The type is interface{} with range: 0..4294967295.
    StatsId interface{}

    // QoS EA Queue Result.
    Queue PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue

    // QoS EA Policer Result.
    Police PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetFilter() yfilter.YFilter { return result.YFilter }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetFilter(yf yfilter.YFilter) { result.YFilter = yf }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetGoName(yname string) string {
    if yname == "stats-id" { return "StatsId" }
    if yname == "queue" { return "Queue" }
    if yname == "police" { return "Police" }
    return ""
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetSegmentPath() string {
    return "result"
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queue" {
        return &result.Queue
    }
    if childYangName == "police" {
        return &result.Police
    }
    return nil
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queue"] = &result.Queue
    children["police"] = &result.Police
    return children
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["stats-id"] = result.StatsId
    return leafs
}

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleName() string { return "cisco_ios_xr" }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetYangName() string { return "result" }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) SetParent(parent types.Entity) { result.parent = parent }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParent() types.Entity { return result.parent }

func (result *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result) GetParentYangName() string { return "qos-show-ea-pclass-st" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue
// QoS EA Queue Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue struct {
    parent types.Entity
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

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetFilter() yfilter.YFilter { return queue.YFilter }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetFilter(yf yfilter.YFilter) { queue.YFilter = yf }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetGoName(yname string) string {
    if yname == "queue-id" { return "QueueId" }
    if yname == "commit-tx" { return "CommitTx" }
    if yname == "excess-tx" { return "ExcessTx" }
    if yname == "drop" { return "Drop" }
    return ""
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetSegmentPath() string {
    return "queue"
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queue-id"] = queue.QueueId
    leafs["commit-tx"] = queue.CommitTx
    leafs["excess-tx"] = queue.ExcessTx
    leafs["drop"] = queue.Drop
    return leafs
}

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleName() string { return "cisco_ios_xr" }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetYangName() string { return "queue" }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) SetParent(parent types.Entity) { queue.parent = parent }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParent() types.Entity { return queue.parent }

func (queue *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Queue) GetParentYangName() string { return "result" }

// PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police
// QoS EA Policer Result
type PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police struct {
    parent types.Entity
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

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetFilter() yfilter.YFilter { return police.YFilter }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetFilter(yf yfilter.YFilter) { police.YFilter = yf }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetGoName(yname string) string {
    if yname == "token-bucket-id" { return "TokenBucketId" }
    if yname == "conform" { return "Conform" }
    if yname == "exceed" { return "Exceed" }
    if yname == "violate" { return "Violate" }
    return ""
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetSegmentPath() string {
    return "police"
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["token-bucket-id"] = police.TokenBucketId
    leafs["conform"] = police.Conform
    leafs["exceed"] = police.Exceed
    leafs["violate"] = police.Violate
    return leafs
}

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleName() string { return "cisco_ios_xr" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetYangName() string { return "police" }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) SetParent(parent types.Entity) { police.parent = parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParent() types.Entity { return police.parent }

func (police *PlatformQosEa_Nodes_Node_Interfaces_Interface_Input_Details_SkywarpQosPolicyClass_QosShowEaPclassSt_Result_Police) GetParentYangName() string { return "result" }

