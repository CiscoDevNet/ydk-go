// This module contains a collection of YANG definitions
// for Cisco IOS-XR ncs5500-qos package operational data.
// 
// This module contains definitions
// for the following management objects:
//   platform-qos: DNX QoS EA operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ncs5500_qos_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ncs5500_qos_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ncs5500-qos-oper platform-qos}", reflect.TypeOf(PlatformQos{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ncs5500-qos-oper:platform-qos", reflect.TypeOf(PlatformQos{}))
}

// DnxQoseaShowAction represents Policer action type
type DnxQoseaShowAction string

const (
    // None
    DnxQoseaShowAction_action_none DnxQoseaShowAction = "action-none"

    // Transmit
    DnxQoseaShowAction_action_transmit DnxQoseaShowAction = "action-transmit"

    // Drop
    DnxQoseaShowAction_action_drop DnxQoseaShowAction = "action-drop"

    // Mark
    DnxQoseaShowAction_action_mark DnxQoseaShowAction = "action-mark"
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

// DnxQoseaShowWred represents WRED type
type DnxQoseaShowWred string

const (
    // WRED based on COS
    DnxQoseaShowWred_wred_cos DnxQoseaShowWred = "wred-cos"

    // WRED based on DSCP
    DnxQoseaShowWred_wred_dscp DnxQoseaShowWred = "wred-dscp"

    // WRED based on Precedence
    DnxQoseaShowWred_wred_precedence DnxQoseaShowWred = "wred-precedence"

    // WRED based on discard class
    DnxQoseaShowWred_wred_discard_class DnxQoseaShowWred = "wred-discard-class"

    // WRED based on MPLS EXP
    DnxQoseaShowWred_wred_mpls_exp DnxQoseaShowWred = "wred-mpls-exp"

    // RED with user defined min and max
    DnxQoseaShowWred_red_with_user_min_max DnxQoseaShowWred = "red-with-user-min-max"

    // RED with default min and max
    DnxQoseaShowWred_red_with_default_min_max DnxQoseaShowWred = "red-with-default-min-max"

    // Invalid
    DnxQoseaShowWred_wred_invalid DnxQoseaShowWred = "wred-invalid"
)

// DnxQoseaShowHpLevel represents Priority level
type DnxQoseaShowHpLevel string

const (
    // High priority queue level 1
    DnxQoseaShowHpLevel_high_priority_level1 DnxQoseaShowHpLevel = "high-priority-level1"

    // High priority queue level 2
    DnxQoseaShowHpLevel_high_priority_level2 DnxQoseaShowHpLevel = "high-priority-level2"

    // High priority queue level 3
    DnxQoseaShowHpLevel_high_priority_level3 DnxQoseaShowHpLevel = "high-priority-level3"

    // High priority queue level 4
    DnxQoseaShowHpLevel_high_priority_level4 DnxQoseaShowHpLevel = "high-priority-level4"

    // High priority queue level 5
    DnxQoseaShowHpLevel_high_priority_level5 DnxQoseaShowHpLevel = "high-priority-level5"

    // High priority queue level 6
    DnxQoseaShowHpLevel_high_priority_level6 DnxQoseaShowHpLevel = "high-priority-level6"

    // High priority queue level 7
    DnxQoseaShowHpLevel_high_priority_level7 DnxQoseaShowHpLevel = "high-priority-level7"

    // Unknown
    DnxQoseaShowHpLevel_unknown DnxQoseaShowHpLevel = "unknown"
)

// QosPolicyAccountEnum represents Qos policy account enum
type QosPolicyAccountEnum string

const (
    // qos serv policy no ac count pref
    QosPolicyAccountEnum_qos_serv_policy_no_ac_count_pref QosPolicyAccountEnum = "qos-serv-policy-no-ac-count-pref"

    // qos serv policy ac count l2
    QosPolicyAccountEnum_qos_serv_policy_ac_count_l2 QosPolicyAccountEnum = "qos-serv-policy-ac-count-l2"

    // qos serv policy no ac count l2
    QosPolicyAccountEnum_qos_serv_policy_no_ac_count_l2 QosPolicyAccountEnum = "qos-serv-policy-no-ac-count-l2"

    // qos serv policy ac count user def
    QosPolicyAccountEnum_qos_serv_policy_ac_count_user_def QosPolicyAccountEnum = "qos-serv-policy-ac-count-user-def"

    // qos serv policy ac count l1
    QosPolicyAccountEnum_qos_serv_policy_ac_count_l1 QosPolicyAccountEnum = "qos-serv-policy-ac-count-l1"
)

// DnxQoseaShowMark represents Mark type
type DnxQoseaShowMark string

const (
    // None
    DnxQoseaShowMark_mark_none DnxQoseaShowMark = "mark-none"

    // DSCP
    DnxQoseaShowMark_dscp DnxQoseaShowMark = "dscp"

    // Precedence
    DnxQoseaShowMark_precedence DnxQoseaShowMark = "precedence"

    // MPLS topmost
    DnxQoseaShowMark_mpls_topmost DnxQoseaShowMark = "mpls-topmost"

    // MPLS imposition
    DnxQoseaShowMark_mpls_imposition DnxQoseaShowMark = "mpls-imposition"

    // Qos group
    DnxQoseaShowMark_qos_group DnxQoseaShowMark = "qos-group"

    // Discard class
    DnxQoseaShowMark_discard_class DnxQoseaShowMark = "discard-class"

    // COS
    DnxQoseaShowMark_cos DnxQoseaShowMark = "cos"

    // Inner COS
    DnxQoseaShowMark_inner_cos DnxQoseaShowMark = "inner-cos"

    // Unsupported type 9
    DnxQoseaShowMark_un_supported9 DnxQoseaShowMark = "un-supported9"

    // Unsupported type 10
    DnxQoseaShowMark_un_supported10 DnxQoseaShowMark = "un-supported10"

    // DSCP tunnel
    DnxQoseaShowMark_dscp_tunnel DnxQoseaShowMark = "dscp-tunnel"

    // Precedence tunnel
    DnxQoseaShowMark_precedence_tunnel DnxQoseaShowMark = "precedence-tunnel"

    // DEI
    DnxQoseaShowMark_dei DnxQoseaShowMark = "dei"

    // DEI Imposition
    DnxQoseaShowMark_dei_imposition DnxQoseaShowMark = "dei-imposition"
)

// DnxQoseaShowPolicyStatus represents Status
type DnxQoseaShowPolicyStatus string

const (
    // No errors
    DnxQoseaShowPolicyStatus_no_error DnxQoseaShowPolicyStatus = "no-error"

    // QoS policy is reset
    DnxQoseaShowPolicyStatus_policy_in_reset DnxQoseaShowPolicyStatus = "policy-in-reset"
)

// DnxQoseaShowIntfStatus represents Intf Status
type DnxQoseaShowIntfStatus string

const (
    // State is unknown
    DnxQoseaShowIntfStatus_state_unknown DnxQoseaShowIntfStatus = "state-unknown"

    // State is Down
    DnxQoseaShowIntfStatus_state_down DnxQoseaShowIntfStatus = "state-down"
)

// DnxQoseaShowLevel represents Level type
type DnxQoseaShowLevel string

const (
    // QoS level1 class
    DnxQoseaShowLevel_level1 DnxQoseaShowLevel = "level1"

    // QoS level2 class
    DnxQoseaShowLevel_level2 DnxQoseaShowLevel = "level2"

    // QoS level3 class
    DnxQoseaShowLevel_level3 DnxQoseaShowLevel = "level3"

    // QoS level4 class
    DnxQoseaShowLevel_level4 DnxQoseaShowLevel = "level4"

    // QoS level5 class
    DnxQoseaShowLevel_level5 DnxQoseaShowLevel = "level5"
)

// DnxQoseaShowQueue represents Priority Queue Type
type DnxQoseaShowQueue string

const (
    // Low priority default queue
    DnxQoseaShowQueue_low_priority_default_queue DnxQoseaShowQueue = "low-priority-default-queue"

    // Low priority queue
    DnxQoseaShowQueue_low_priority_queue DnxQoseaShowQueue = "low-priority-queue"

    // High priority queue
    DnxQoseaShowQueue_high_priority_queue DnxQoseaShowQueue = "high-priority-queue"

    // Queue priority unknown
    DnxQoseaShowQueue_unknown_queue_type DnxQoseaShowQueue = "unknown-queue-type"
)

// PlatformQos
// DNX QoS EA operational data
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
    return "Cisco-IOS-XR-ncs5500-qos-oper:platform-qos"
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

func (platformQos *PlatformQos) GetParentYangName() string { return "Cisco-IOS-XR-ncs5500-qos-oper" }

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

    // QoS list of interfaces.
    Interfaces PlatformQos_Nodes_Node_Interfaces

    // QoS list of remote interfaces.
    RemoteInterfaces PlatformQos_Nodes_Node_RemoteInterfaces
}

func (node *PlatformQos_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *PlatformQos_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *PlatformQos_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "bundle-interfaces" { return "BundleInterfaces" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "remote-interfaces" { return "RemoteInterfaces" }
    return ""
}

func (node *PlatformQos_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *PlatformQos_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-interfaces" {
        return &node.BundleInterfaces
    }
    if childYangName == "interfaces" {
        return &node.Interfaces
    }
    if childYangName == "remote-interfaces" {
        return &node.RemoteInterfaces
    }
    return nil
}

func (node *PlatformQos_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle-interfaces"] = &node.BundleInterfaces
    children["interfaces"] = &node.Interfaces
    children["remote-interfaces"] = &node.RemoteInterfaces
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

    // QoS interface names. The type is slice of
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
// QoS interface names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Bundle interface name. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // NPU ID. The type is interface{} with range: -2147483648..2147483647.
    NpuId interface{}

    // The interface direction on which QoS is applied to. The type is string.
    QosDirection interface{}

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetFilter() yfilter.YFilter { return bundleInterface.YFilter }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) SetFilter(yf yfilter.YFilter) { bundleInterface.YFilter = yf }

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "npu-id" { return "NpuId" }
    if yname == "qos-direction" { return "QosDirection" }
    if yname == "member-interfaces" { return "MemberInterfaces" }
    if yname == "policy-details" { return "PolicyDetails" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetSegmentPath() string {
    return "bundle-interface" + "[interface-name='" + fmt.Sprintf("%v", bundleInterface.InterfaceName) + "']"
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "member-interfaces" {
        return &bundleInterface.MemberInterfaces
    }
    if childYangName == "policy-details" {
        return &bundleInterface.PolicyDetails
    }
    if childYangName == "classes" {
        return &bundleInterface.Classes
    }
    return nil
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["member-interfaces"] = &bundleInterface.MemberInterfaces
    children["policy-details"] = &bundleInterface.PolicyDetails
    children["classes"] = &bundleInterface.Classes
    return children
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = bundleInterface.InterfaceName
    leafs["npu-id"] = bundleInterface.NpuId
    leafs["qos-direction"] = bundleInterface.QosDirection
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

    // QoS interface names. The type is slice of
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
// QoS interface names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Member interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetFilter() yfilter.YFilter { return memberInterface.YFilter }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) SetFilter(yf yfilter.YFilter) { memberInterface.YFilter = yf }

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-details" { return "PolicyDetails" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetSegmentPath() string {
    return "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-details" {
        return &memberInterface.PolicyDetails
    }
    if childYangName == "classes" {
        return &memberInterface.Classes
    }
    return nil
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-details"] = &memberInterface.PolicyDetails
    children["classes"] = &memberInterface.Classes
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

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NPU ID. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // InterfaceHandle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface Bandwidth (in kbps). The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    InterfaceBandwidthKbps interface{}

    // Policy name. The type is string with length: 0..64.
    PolicyName interface{}

    // Number of Classes. The type is interface{} with range: 0..65535.
    TotalNumberOfClasses interface{}

    // VOQ base address. The type is interface{} with range: 0..4294967295.
    VoqBaseAddress interface{}

    // VOQ stats handle. The type is interface{} with range:
    // 0..18446744073709551615.
    VoqStatsHandle interface{}

    // QoS Statistics Accounting Type. The type is QosPolicyAccountEnum.
    StatsAccountingType interface{}

    // Policy Status. The type is DnxQoseaShowPolicyStatus.
    PolicyStatus interface{}

    // Interface Status. The type is DnxQoseaShowIntfStatus.
    InterfaceStatus interface{}
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetFilter() yfilter.YFilter { return policyDetails.YFilter }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) SetFilter(yf yfilter.YFilter) { policyDetails.YFilter = yf }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "interface-bandwidth-kbps" { return "InterfaceBandwidthKbps" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "total-number-of-classes" { return "TotalNumberOfClasses" }
    if yname == "voq-base-address" { return "VoqBaseAddress" }
    if yname == "voq-stats-handle" { return "VoqStatsHandle" }
    if yname == "stats-accounting-type" { return "StatsAccountingType" }
    if yname == "policy-status" { return "PolicyStatus" }
    if yname == "interface-status" { return "InterfaceStatus" }
    return ""
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetSegmentPath() string {
    return "policy-details"
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = policyDetails.NpuId
    leafs["interface-handle"] = policyDetails.InterfaceHandle
    leafs["interface-bandwidth-kbps"] = policyDetails.InterfaceBandwidthKbps
    leafs["policy-name"] = policyDetails.PolicyName
    leafs["total-number-of-classes"] = policyDetails.TotalNumberOfClasses
    leafs["voq-base-address"] = policyDetails.VoqBaseAddress
    leafs["voq-stats-handle"] = policyDetails.VoqStatsHandle
    leafs["stats-accounting-type"] = policyDetails.StatsAccountingType
    leafs["policy-status"] = policyDetails.PolicyStatus
    leafs["interface-status"] = policyDetails.InterfaceStatus
    return leafs
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetBundleName() string { return "cisco_ios_xr" }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetYangName() string { return "policy-details" }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) SetParent(parent types.Entity) { policyDetails.parent = parent }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetParent() types.Entity { return policyDetails.parent }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetParentYangName() string { return "member-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class.
    Class []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetYangName() string { return "classes" }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetParent() types.Entity { return classes.parent }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetParentYangName() string { return "member-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. QoS policy class name at level 1. The type is
    // string.
    LevelOneClassName interface{}

    // QoS policy child class name at level 2. The type is string.
    LevelTwoClassName interface{}

    // Class level. The type is DnxQoseaShowLevel.
    ClassLevel interface{}

    // Egress Queue ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EgressQueueId interface{}

    // Queue type. The type is DnxQoseaShowQueue.
    QueueType interface{}

    // Priority level. The type is DnxQoseaShowHpLevel.
    PriorityLevel interface{}

    // Hardware maximum rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwareMaxRateKbps interface{}

    // Hardware minimum rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwareMinRateKbps interface{}

    // Configured excess bandwidth percentage. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    ConfigExcessBandwidthPercent interface{}

    // Configured excess bandwidth unit. The type is interface{} with range:
    // 0..4294967295.
    ConfigExcessBandwidthUnit interface{}

    // Hardware excess bandwidth weight. The type is interface{} with range:
    // 0..4294967295.
    HardwareExcessBandwidthWeight interface{}

    // Network minimum Bandwith. The type is interface{} with range:
    // 0..4294967295.
    NetworkMinBandwidthKbps interface{}

    // Hardware queue limit in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    HardwareQueueLimitBytes interface{}

    // Hardware queue limit in microseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    HardwareQueueLimitMicroseconds interface{}

    // PolicerBucketID. The type is interface{} with range: 0..4294967295.
    PolicerBucketId interface{}

    // PolicerStatsHandle. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicerStatsHandle interface{}

    // Hardware policer average in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwarePolicerAverageRateKbps interface{}

    // Hardware policer peak rate. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerPeakRateKbps interface{}

    // Hardware policer conform burst. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerConformBurstBytes interface{}

    // Hardware policer excess burst. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerExcessBurstBytes interface{}

    // Configured maximum rate.
    ConfigMaxRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate

    // Configured minimum rate.
    ConfigMinRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate

    // Configured queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit

    // Configured policer average rate.
    ConfigPolicerAverageRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate

    // Config policer peak rate.
    ConfigPolicerPeakRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate

    // Configured policer conform burst.
    ConfigPolicerConformBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst

    // Configured policer excess burst.
    ConfigPolicerExcessBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst

    // Conform action.
    ConformAction PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction

    // Exceed action.
    ExceedAction PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction

    // Violate action.
    ViolateAction PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction

    // IP mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark.
    IpMark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark

    // Common mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark.
    CommonMark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark

    // MPLS mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark.
    MplsMark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark

    // WRED parameters. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred.
    Wred []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetGoName(yname string) string {
    if yname == "level-one-class-name" { return "LevelOneClassName" }
    if yname == "level-two-class-name" { return "LevelTwoClassName" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "egress-queue-id" { return "EgressQueueId" }
    if yname == "queue-type" { return "QueueType" }
    if yname == "priority-level" { return "PriorityLevel" }
    if yname == "hardware-max-rate-kbps" { return "HardwareMaxRateKbps" }
    if yname == "hardware-min-rate-kbps" { return "HardwareMinRateKbps" }
    if yname == "config-excess-bandwidth-percent" { return "ConfigExcessBandwidthPercent" }
    if yname == "config-excess-bandwidth-unit" { return "ConfigExcessBandwidthUnit" }
    if yname == "hardware-excess-bandwidth-weight" { return "HardwareExcessBandwidthWeight" }
    if yname == "network-min-bandwidth-kbps" { return "NetworkMinBandwidthKbps" }
    if yname == "hardware-queue-limit-bytes" { return "HardwareQueueLimitBytes" }
    if yname == "hardware-queue-limit-microseconds" { return "HardwareQueueLimitMicroseconds" }
    if yname == "policer-bucket-id" { return "PolicerBucketId" }
    if yname == "policer-stats-handle" { return "PolicerStatsHandle" }
    if yname == "hardware-policer-average-rate-kbps" { return "HardwarePolicerAverageRateKbps" }
    if yname == "hardware-policer-peak-rate-kbps" { return "HardwarePolicerPeakRateKbps" }
    if yname == "hardware-policer-conform-burst-bytes" { return "HardwarePolicerConformBurstBytes" }
    if yname == "hardware-policer-excess-burst-bytes" { return "HardwarePolicerExcessBurstBytes" }
    if yname == "config-max-rate" { return "ConfigMaxRate" }
    if yname == "config-min-rate" { return "ConfigMinRate" }
    if yname == "config-queue-limit" { return "ConfigQueueLimit" }
    if yname == "config-policer-average-rate" { return "ConfigPolicerAverageRate" }
    if yname == "config-policer-peak-rate" { return "ConfigPolicerPeakRate" }
    if yname == "config-policer-conform-burst" { return "ConfigPolicerConformBurst" }
    if yname == "config-policer-excess-burst" { return "ConfigPolicerExcessBurst" }
    if yname == "conform-action" { return "ConformAction" }
    if yname == "exceed-action" { return "ExceedAction" }
    if yname == "violate-action" { return "ViolateAction" }
    if yname == "ip-mark" { return "IpMark" }
    if yname == "common-mark" { return "CommonMark" }
    if yname == "mpls-mark" { return "MplsMark" }
    if yname == "wred" { return "Wred" }
    return ""
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetSegmentPath() string {
    return "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-max-rate" {
        return &class.ConfigMaxRate
    }
    if childYangName == "config-min-rate" {
        return &class.ConfigMinRate
    }
    if childYangName == "config-queue-limit" {
        return &class.ConfigQueueLimit
    }
    if childYangName == "config-policer-average-rate" {
        return &class.ConfigPolicerAverageRate
    }
    if childYangName == "config-policer-peak-rate" {
        return &class.ConfigPolicerPeakRate
    }
    if childYangName == "config-policer-conform-burst" {
        return &class.ConfigPolicerConformBurst
    }
    if childYangName == "config-policer-excess-burst" {
        return &class.ConfigPolicerExcessBurst
    }
    if childYangName == "conform-action" {
        return &class.ConformAction
    }
    if childYangName == "exceed-action" {
        return &class.ExceedAction
    }
    if childYangName == "violate-action" {
        return &class.ViolateAction
    }
    if childYangName == "ip-mark" {
        for _, c := range class.IpMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark{}
        class.IpMark = append(class.IpMark, child)
        return &class.IpMark[len(class.IpMark)-1]
    }
    if childYangName == "common-mark" {
        for _, c := range class.CommonMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark{}
        class.CommonMark = append(class.CommonMark, child)
        return &class.CommonMark[len(class.CommonMark)-1]
    }
    if childYangName == "mpls-mark" {
        for _, c := range class.MplsMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark{}
        class.MplsMark = append(class.MplsMark, child)
        return &class.MplsMark[len(class.MplsMark)-1]
    }
    if childYangName == "wred" {
        for _, c := range class.Wred {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred{}
        class.Wred = append(class.Wred, child)
        return &class.Wred[len(class.Wred)-1]
    }
    return nil
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config-max-rate"] = &class.ConfigMaxRate
    children["config-min-rate"] = &class.ConfigMinRate
    children["config-queue-limit"] = &class.ConfigQueueLimit
    children["config-policer-average-rate"] = &class.ConfigPolicerAverageRate
    children["config-policer-peak-rate"] = &class.ConfigPolicerPeakRate
    children["config-policer-conform-burst"] = &class.ConfigPolicerConformBurst
    children["config-policer-excess-burst"] = &class.ConfigPolicerExcessBurst
    children["conform-action"] = &class.ConformAction
    children["exceed-action"] = &class.ExceedAction
    children["violate-action"] = &class.ViolateAction
    for i := range class.IpMark {
        children[class.IpMark[i].GetSegmentPath()] = &class.IpMark[i]
    }
    for i := range class.CommonMark {
        children[class.CommonMark[i].GetSegmentPath()] = &class.CommonMark[i]
    }
    for i := range class.MplsMark {
        children[class.MplsMark[i].GetSegmentPath()] = &class.MplsMark[i]
    }
    for i := range class.Wred {
        children[class.Wred[i].GetSegmentPath()] = &class.Wred[i]
    }
    return children
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level-one-class-name"] = class.LevelOneClassName
    leafs["level-two-class-name"] = class.LevelTwoClassName
    leafs["class-level"] = class.ClassLevel
    leafs["egress-queue-id"] = class.EgressQueueId
    leafs["queue-type"] = class.QueueType
    leafs["priority-level"] = class.PriorityLevel
    leafs["hardware-max-rate-kbps"] = class.HardwareMaxRateKbps
    leafs["hardware-min-rate-kbps"] = class.HardwareMinRateKbps
    leafs["config-excess-bandwidth-percent"] = class.ConfigExcessBandwidthPercent
    leafs["config-excess-bandwidth-unit"] = class.ConfigExcessBandwidthUnit
    leafs["hardware-excess-bandwidth-weight"] = class.HardwareExcessBandwidthWeight
    leafs["network-min-bandwidth-kbps"] = class.NetworkMinBandwidthKbps
    leafs["hardware-queue-limit-bytes"] = class.HardwareQueueLimitBytes
    leafs["hardware-queue-limit-microseconds"] = class.HardwareQueueLimitMicroseconds
    leafs["policer-bucket-id"] = class.PolicerBucketId
    leafs["policer-stats-handle"] = class.PolicerStatsHandle
    leafs["hardware-policer-average-rate-kbps"] = class.HardwarePolicerAverageRateKbps
    leafs["hardware-policer-peak-rate-kbps"] = class.HardwarePolicerPeakRateKbps
    leafs["hardware-policer-conform-burst-bytes"] = class.HardwarePolicerConformBurstBytes
    leafs["hardware-policer-excess-burst-bytes"] = class.HardwarePolicerExcessBurstBytes
    return leafs
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetYangName() string { return "class" }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetParentYangName() string { return "classes" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetFilter() yfilter.YFilter { return configMaxRate.YFilter }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) SetFilter(yf yfilter.YFilter) { configMaxRate.YFilter = yf }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetSegmentPath() string {
    return "config-max-rate"
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMaxRate.PolicyValue
    leafs["policy-unit"] = configMaxRate.PolicyUnit
    return leafs
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetBundleName() string { return "cisco_ios_xr" }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetYangName() string { return "config-max-rate" }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) SetParent(parent types.Entity) { configMaxRate.parent = parent }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetParent() types.Entity { return configMaxRate.parent }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetFilter() yfilter.YFilter { return configMinRate.YFilter }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) SetFilter(yf yfilter.YFilter) { configMinRate.YFilter = yf }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetSegmentPath() string {
    return "config-min-rate"
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMinRate.PolicyValue
    leafs["policy-unit"] = configMinRate.PolicyUnit
    return leafs
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetBundleName() string { return "cisco_ios_xr" }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetYangName() string { return "config-min-rate" }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) SetParent(parent types.Entity) { configMinRate.parent = parent }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetParent() types.Entity { return configMinRate.parent }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetFilter() yfilter.YFilter { return configQueueLimit.YFilter }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) SetFilter(yf yfilter.YFilter) { configQueueLimit.YFilter = yf }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetSegmentPath() string {
    return "config-queue-limit"
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configQueueLimit.PolicyValue
    leafs["policy-unit"] = configQueueLimit.PolicyUnit
    return leafs
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetBundleName() string { return "cisco_ios_xr" }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetYangName() string { return "config-queue-limit" }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) SetParent(parent types.Entity) { configQueueLimit.parent = parent }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetParent() types.Entity { return configQueueLimit.parent }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetFilter() yfilter.YFilter { return configPolicerAverageRate.YFilter }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) SetFilter(yf yfilter.YFilter) { configPolicerAverageRate.YFilter = yf }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetSegmentPath() string {
    return "config-policer-average-rate"
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerAverageRate.PolicyValue
    leafs["policy-unit"] = configPolicerAverageRate.PolicyUnit
    return leafs
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetYangName() string { return "config-policer-average-rate" }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) SetParent(parent types.Entity) { configPolicerAverageRate.parent = parent }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetParent() types.Entity { return configPolicerAverageRate.parent }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetFilter() yfilter.YFilter { return configPolicerPeakRate.YFilter }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) SetFilter(yf yfilter.YFilter) { configPolicerPeakRate.YFilter = yf }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetSegmentPath() string {
    return "config-policer-peak-rate"
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerPeakRate.PolicyValue
    leafs["policy-unit"] = configPolicerPeakRate.PolicyUnit
    return leafs
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetYangName() string { return "config-policer-peak-rate" }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) SetParent(parent types.Entity) { configPolicerPeakRate.parent = parent }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetParent() types.Entity { return configPolicerPeakRate.parent }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetFilter() yfilter.YFilter { return configPolicerConformBurst.YFilter }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) SetFilter(yf yfilter.YFilter) { configPolicerConformBurst.YFilter = yf }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetSegmentPath() string {
    return "config-policer-conform-burst"
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerConformBurst.PolicyValue
    leafs["policy-unit"] = configPolicerConformBurst.PolicyUnit
    return leafs
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetYangName() string { return "config-policer-conform-burst" }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) SetParent(parent types.Entity) { configPolicerConformBurst.parent = parent }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetParent() types.Entity { return configPolicerConformBurst.parent }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetFilter() yfilter.YFilter { return configPolicerExcessBurst.YFilter }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) SetFilter(yf yfilter.YFilter) { configPolicerExcessBurst.YFilter = yf }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetSegmentPath() string {
    return "config-policer-excess-burst"
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerExcessBurst.PolicyValue
    leafs["policy-unit"] = configPolicerExcessBurst.PolicyUnit
    return leafs
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetYangName() string { return "config-policer-excess-burst" }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) SetParent(parent types.Entity) { configPolicerExcessBurst.parent = parent }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetParent() types.Entity { return configPolicerExcessBurst.parent }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetFilter() yfilter.YFilter { return conformAction.YFilter }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) SetFilter(yf yfilter.YFilter) { conformAction.YFilter = yf }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetSegmentPath() string {
    return "conform-action"
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range conformAction.Mark {
            if conformAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark{}
        conformAction.Mark = append(conformAction.Mark, child)
        return &conformAction.Mark[len(conformAction.Mark)-1]
    }
    return nil
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range conformAction.Mark {
        children[conformAction.Mark[i].GetSegmentPath()] = &conformAction.Mark[i]
    }
    return children
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = conformAction.ActionType
    return leafs
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetBundleName() string { return "cisco_ios_xr" }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetYangName() string { return "conform-action" }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) SetParent(parent types.Entity) { conformAction.parent = parent }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetParent() types.Entity { return conformAction.parent }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetParentYangName() string { return "conform-action" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetFilter() yfilter.YFilter { return exceedAction.YFilter }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) SetFilter(yf yfilter.YFilter) { exceedAction.YFilter = yf }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetSegmentPath() string {
    return "exceed-action"
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range exceedAction.Mark {
            if exceedAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark{}
        exceedAction.Mark = append(exceedAction.Mark, child)
        return &exceedAction.Mark[len(exceedAction.Mark)-1]
    }
    return nil
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range exceedAction.Mark {
        children[exceedAction.Mark[i].GetSegmentPath()] = &exceedAction.Mark[i]
    }
    return children
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = exceedAction.ActionType
    return leafs
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetBundleName() string { return "cisco_ios_xr" }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetYangName() string { return "exceed-action" }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) SetParent(parent types.Entity) { exceedAction.parent = parent }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetParent() types.Entity { return exceedAction.parent }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetParentYangName() string { return "exceed-action" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetFilter() yfilter.YFilter { return violateAction.YFilter }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) SetFilter(yf yfilter.YFilter) { violateAction.YFilter = yf }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetSegmentPath() string {
    return "violate-action"
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range violateAction.Mark {
            if violateAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark{}
        violateAction.Mark = append(violateAction.Mark, child)
        return &violateAction.Mark[len(violateAction.Mark)-1]
    }
    return nil
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range violateAction.Mark {
        children[violateAction.Mark[i].GetSegmentPath()] = &violateAction.Mark[i]
    }
    return children
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = violateAction.ActionType
    return leafs
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetBundleName() string { return "cisco_ios_xr" }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetYangName() string { return "violate-action" }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) SetParent(parent types.Entity) { violateAction.parent = parent }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetParent() types.Entity { return violateAction.parent }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetParentYangName() string { return "violate-action" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetFilter() yfilter.YFilter { return ipMark.YFilter }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) SetFilter(yf yfilter.YFilter) { ipMark.YFilter = yf }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetSegmentPath() string {
    return "ip-mark"
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = ipMark.MarkType
    leafs["mark-value"] = ipMark.MarkValue
    return leafs
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetBundleName() string { return "cisco_ios_xr" }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetYangName() string { return "ip-mark" }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) SetParent(parent types.Entity) { ipMark.parent = parent }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetParent() types.Entity { return ipMark.parent }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetFilter() yfilter.YFilter { return commonMark.YFilter }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) SetFilter(yf yfilter.YFilter) { commonMark.YFilter = yf }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetSegmentPath() string {
    return "common-mark"
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = commonMark.MarkType
    leafs["mark-value"] = commonMark.MarkValue
    return leafs
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetBundleName() string { return "cisco_ios_xr" }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetYangName() string { return "common-mark" }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) SetParent(parent types.Entity) { commonMark.parent = parent }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetParent() types.Entity { return commonMark.parent }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetFilter() yfilter.YFilter { return mplsMark.YFilter }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) SetFilter(yf yfilter.YFilter) { mplsMark.YFilter = yf }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetSegmentPath() string {
    return "mpls-mark"
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mplsMark.MarkType
    leafs["mark-value"] = mplsMark.MarkValue
    return leafs
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetBundleName() string { return "cisco_ios_xr" }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetYangName() string { return "mpls-mark" }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) SetParent(parent types.Entity) { mplsMark.parent = parent }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetParent() types.Entity { return mplsMark.parent }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // WREDMatchType. The type is DnxQoseaShowWred.
    WredMatchType interface{}

    // Hardware minimum threshold. The type is interface{} with range:
    // 0..4294967295.
    HardwareMinThresholdBytes interface{}

    // Hardware maximum threshold. The type is interface{} with range:
    // 0..4294967295.
    HardwareMaxThresholdBytes interface{}

    // First segment. The type is interface{} with range: 0..65535.
    FirstSegment interface{}

    // Segment size. The type is interface{} with range: 0..4294967295.
    SegmentSize interface{}

    // WRED match values.
    WredMatchValue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue

    // Configured minimum threshold.
    ConfigMinThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold

    // Configured maximum threshold.
    ConfigMaxThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetFilter() yfilter.YFilter { return wred.YFilter }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) SetFilter(yf yfilter.YFilter) { wred.YFilter = yf }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetGoName(yname string) string {
    if yname == "wred-match-type" { return "WredMatchType" }
    if yname == "hardware-min-threshold-bytes" { return "HardwareMinThresholdBytes" }
    if yname == "hardware-max-threshold-bytes" { return "HardwareMaxThresholdBytes" }
    if yname == "first-segment" { return "FirstSegment" }
    if yname == "segment-size" { return "SegmentSize" }
    if yname == "wred-match-value" { return "WredMatchValue" }
    if yname == "config-min-threshold" { return "ConfigMinThreshold" }
    if yname == "config-max-threshold" { return "ConfigMaxThreshold" }
    return ""
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetSegmentPath() string {
    return "wred"
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wred-match-value" {
        return &wred.WredMatchValue
    }
    if childYangName == "config-min-threshold" {
        return &wred.ConfigMinThreshold
    }
    if childYangName == "config-max-threshold" {
        return &wred.ConfigMaxThreshold
    }
    return nil
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wred-match-value"] = &wred.WredMatchValue
    children["config-min-threshold"] = &wred.ConfigMinThreshold
    children["config-max-threshold"] = &wred.ConfigMaxThreshold
    return children
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wred-match-type"] = wred.WredMatchType
    leafs["hardware-min-threshold-bytes"] = wred.HardwareMinThresholdBytes
    leafs["hardware-max-threshold-bytes"] = wred.HardwareMaxThresholdBytes
    leafs["first-segment"] = wred.FirstSegment
    leafs["segment-size"] = wred.SegmentSize
    return leafs
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetBundleName() string { return "cisco_ios_xr" }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetYangName() string { return "wred" }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) SetParent(parent types.Entity) { wred.parent = parent }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetParent() types.Entity { return wred.parent }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetFilter() yfilter.YFilter { return wredMatchValue.YFilter }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) SetFilter(yf yfilter.YFilter) { wredMatchValue.YFilter = yf }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetGoName(yname string) string {
    if yname == "dnx-qosea-show-red-match-value" { return "DnxQoseaShowRedMatchValue" }
    return ""
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetSegmentPath() string {
    return "wred-match-value"
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dnx-qosea-show-red-match-value" {
        for _, c := range wredMatchValue.DnxQoseaShowRedMatchValue {
            if wredMatchValue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue{}
        wredMatchValue.DnxQoseaShowRedMatchValue = append(wredMatchValue.DnxQoseaShowRedMatchValue, child)
        return &wredMatchValue.DnxQoseaShowRedMatchValue[len(wredMatchValue.DnxQoseaShowRedMatchValue)-1]
    }
    return nil
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        children[wredMatchValue.DnxQoseaShowRedMatchValue[i].GetSegmentPath()] = &wredMatchValue.DnxQoseaShowRedMatchValue[i]
    }
    return children
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetBundleName() string { return "cisco_ios_xr" }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetYangName() string { return "wred-match-value" }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) SetParent(parent types.Entity) { wredMatchValue.parent = parent }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetParent() types.Entity { return wredMatchValue.parent }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetFilter() yfilter.YFilter { return dnxQoseaShowRedMatchValue.YFilter }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) SetFilter(yf yfilter.YFilter) { dnxQoseaShowRedMatchValue.YFilter = yf }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetGoName(yname string) string {
    if yname == "range-start" { return "RangeStart" }
    if yname == "range-end" { return "RangeEnd" }
    return ""
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetSegmentPath() string {
    return "dnx-qosea-show-red-match-value"
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["range-start"] = dnxQoseaShowRedMatchValue.RangeStart
    leafs["range-end"] = dnxQoseaShowRedMatchValue.RangeEnd
    return leafs
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetBundleName() string { return "cisco_ios_xr" }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetYangName() string { return "dnx-qosea-show-red-match-value" }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) SetParent(parent types.Entity) { dnxQoseaShowRedMatchValue.parent = parent }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetParent() types.Entity { return dnxQoseaShowRedMatchValue.parent }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetParentYangName() string { return "wred-match-value" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetFilter() yfilter.YFilter { return configMinThreshold.YFilter }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) SetFilter(yf yfilter.YFilter) { configMinThreshold.YFilter = yf }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetSegmentPath() string {
    return "config-min-threshold"
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMinThreshold.PolicyValue
    leafs["policy-unit"] = configMinThreshold.PolicyUnit
    return leafs
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetYangName() string { return "config-min-threshold" }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) SetParent(parent types.Entity) { configMinThreshold.parent = parent }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetParent() types.Entity { return configMinThreshold.parent }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetFilter() yfilter.YFilter { return configMaxThreshold.YFilter }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) SetFilter(yf yfilter.YFilter) { configMaxThreshold.YFilter = yf }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetSegmentPath() string {
    return "config-max-threshold"
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMaxThreshold.PolicyValue
    leafs["policy-unit"] = configMaxThreshold.PolicyUnit
    return leafs
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetYangName() string { return "config-max-threshold" }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) SetParent(parent types.Entity) { configMaxThreshold.parent = parent }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetParent() types.Entity { return configMaxThreshold.parent }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NPU ID. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // InterfaceHandle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface Bandwidth (in kbps). The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    InterfaceBandwidthKbps interface{}

    // Policy name. The type is string with length: 0..64.
    PolicyName interface{}

    // Number of Classes. The type is interface{} with range: 0..65535.
    TotalNumberOfClasses interface{}

    // VOQ base address. The type is interface{} with range: 0..4294967295.
    VoqBaseAddress interface{}

    // VOQ stats handle. The type is interface{} with range:
    // 0..18446744073709551615.
    VoqStatsHandle interface{}

    // QoS Statistics Accounting Type. The type is QosPolicyAccountEnum.
    StatsAccountingType interface{}

    // Policy Status. The type is DnxQoseaShowPolicyStatus.
    PolicyStatus interface{}

    // Interface Status. The type is DnxQoseaShowIntfStatus.
    InterfaceStatus interface{}
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetFilter() yfilter.YFilter { return policyDetails.YFilter }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) SetFilter(yf yfilter.YFilter) { policyDetails.YFilter = yf }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "interface-bandwidth-kbps" { return "InterfaceBandwidthKbps" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "total-number-of-classes" { return "TotalNumberOfClasses" }
    if yname == "voq-base-address" { return "VoqBaseAddress" }
    if yname == "voq-stats-handle" { return "VoqStatsHandle" }
    if yname == "stats-accounting-type" { return "StatsAccountingType" }
    if yname == "policy-status" { return "PolicyStatus" }
    if yname == "interface-status" { return "InterfaceStatus" }
    return ""
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetSegmentPath() string {
    return "policy-details"
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = policyDetails.NpuId
    leafs["interface-handle"] = policyDetails.InterfaceHandle
    leafs["interface-bandwidth-kbps"] = policyDetails.InterfaceBandwidthKbps
    leafs["policy-name"] = policyDetails.PolicyName
    leafs["total-number-of-classes"] = policyDetails.TotalNumberOfClasses
    leafs["voq-base-address"] = policyDetails.VoqBaseAddress
    leafs["voq-stats-handle"] = policyDetails.VoqStatsHandle
    leafs["stats-accounting-type"] = policyDetails.StatsAccountingType
    leafs["policy-status"] = policyDetails.PolicyStatus
    leafs["interface-status"] = policyDetails.InterfaceStatus
    return leafs
}

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetBundleName() string { return "cisco_ios_xr" }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetYangName() string { return "policy-details" }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) SetParent(parent types.Entity) { policyDetails.parent = parent }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetParent() types.Entity { return policyDetails.parent }

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetParentYangName() string { return "bundle-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class.
    Class []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetYangName() string { return "classes" }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetParent() types.Entity { return classes.parent }

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetParentYangName() string { return "bundle-interface" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. QoS policy class name at level 1. The type is
    // string.
    LevelOneClassName interface{}

    // QoS policy child class name at level 2. The type is string.
    LevelTwoClassName interface{}

    // Class level. The type is DnxQoseaShowLevel.
    ClassLevel interface{}

    // Egress Queue ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EgressQueueId interface{}

    // Queue type. The type is DnxQoseaShowQueue.
    QueueType interface{}

    // Priority level. The type is DnxQoseaShowHpLevel.
    PriorityLevel interface{}

    // Hardware maximum rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwareMaxRateKbps interface{}

    // Hardware minimum rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwareMinRateKbps interface{}

    // Configured excess bandwidth percentage. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    ConfigExcessBandwidthPercent interface{}

    // Configured excess bandwidth unit. The type is interface{} with range:
    // 0..4294967295.
    ConfigExcessBandwidthUnit interface{}

    // Hardware excess bandwidth weight. The type is interface{} with range:
    // 0..4294967295.
    HardwareExcessBandwidthWeight interface{}

    // Network minimum Bandwith. The type is interface{} with range:
    // 0..4294967295.
    NetworkMinBandwidthKbps interface{}

    // Hardware queue limit in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    HardwareQueueLimitBytes interface{}

    // Hardware queue limit in microseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    HardwareQueueLimitMicroseconds interface{}

    // PolicerBucketID. The type is interface{} with range: 0..4294967295.
    PolicerBucketId interface{}

    // PolicerStatsHandle. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicerStatsHandle interface{}

    // Hardware policer average in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwarePolicerAverageRateKbps interface{}

    // Hardware policer peak rate. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerPeakRateKbps interface{}

    // Hardware policer conform burst. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerConformBurstBytes interface{}

    // Hardware policer excess burst. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerExcessBurstBytes interface{}

    // Configured maximum rate.
    ConfigMaxRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate

    // Configured minimum rate.
    ConfigMinRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate

    // Configured queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit

    // Configured policer average rate.
    ConfigPolicerAverageRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate

    // Config policer peak rate.
    ConfigPolicerPeakRate PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate

    // Configured policer conform burst.
    ConfigPolicerConformBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst

    // Configured policer excess burst.
    ConfigPolicerExcessBurst PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst

    // Conform action.
    ConformAction PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction

    // Exceed action.
    ExceedAction PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction

    // Violate action.
    ViolateAction PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction

    // IP mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark.
    IpMark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark

    // Common mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark.
    CommonMark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark

    // MPLS mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark.
    MplsMark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark

    // WRED parameters. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred.
    Wred []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetGoName(yname string) string {
    if yname == "level-one-class-name" { return "LevelOneClassName" }
    if yname == "level-two-class-name" { return "LevelTwoClassName" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "egress-queue-id" { return "EgressQueueId" }
    if yname == "queue-type" { return "QueueType" }
    if yname == "priority-level" { return "PriorityLevel" }
    if yname == "hardware-max-rate-kbps" { return "HardwareMaxRateKbps" }
    if yname == "hardware-min-rate-kbps" { return "HardwareMinRateKbps" }
    if yname == "config-excess-bandwidth-percent" { return "ConfigExcessBandwidthPercent" }
    if yname == "config-excess-bandwidth-unit" { return "ConfigExcessBandwidthUnit" }
    if yname == "hardware-excess-bandwidth-weight" { return "HardwareExcessBandwidthWeight" }
    if yname == "network-min-bandwidth-kbps" { return "NetworkMinBandwidthKbps" }
    if yname == "hardware-queue-limit-bytes" { return "HardwareQueueLimitBytes" }
    if yname == "hardware-queue-limit-microseconds" { return "HardwareQueueLimitMicroseconds" }
    if yname == "policer-bucket-id" { return "PolicerBucketId" }
    if yname == "policer-stats-handle" { return "PolicerStatsHandle" }
    if yname == "hardware-policer-average-rate-kbps" { return "HardwarePolicerAverageRateKbps" }
    if yname == "hardware-policer-peak-rate-kbps" { return "HardwarePolicerPeakRateKbps" }
    if yname == "hardware-policer-conform-burst-bytes" { return "HardwarePolicerConformBurstBytes" }
    if yname == "hardware-policer-excess-burst-bytes" { return "HardwarePolicerExcessBurstBytes" }
    if yname == "config-max-rate" { return "ConfigMaxRate" }
    if yname == "config-min-rate" { return "ConfigMinRate" }
    if yname == "config-queue-limit" { return "ConfigQueueLimit" }
    if yname == "config-policer-average-rate" { return "ConfigPolicerAverageRate" }
    if yname == "config-policer-peak-rate" { return "ConfigPolicerPeakRate" }
    if yname == "config-policer-conform-burst" { return "ConfigPolicerConformBurst" }
    if yname == "config-policer-excess-burst" { return "ConfigPolicerExcessBurst" }
    if yname == "conform-action" { return "ConformAction" }
    if yname == "exceed-action" { return "ExceedAction" }
    if yname == "violate-action" { return "ViolateAction" }
    if yname == "ip-mark" { return "IpMark" }
    if yname == "common-mark" { return "CommonMark" }
    if yname == "mpls-mark" { return "MplsMark" }
    if yname == "wred" { return "Wred" }
    return ""
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetSegmentPath() string {
    return "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-max-rate" {
        return &class.ConfigMaxRate
    }
    if childYangName == "config-min-rate" {
        return &class.ConfigMinRate
    }
    if childYangName == "config-queue-limit" {
        return &class.ConfigQueueLimit
    }
    if childYangName == "config-policer-average-rate" {
        return &class.ConfigPolicerAverageRate
    }
    if childYangName == "config-policer-peak-rate" {
        return &class.ConfigPolicerPeakRate
    }
    if childYangName == "config-policer-conform-burst" {
        return &class.ConfigPolicerConformBurst
    }
    if childYangName == "config-policer-excess-burst" {
        return &class.ConfigPolicerExcessBurst
    }
    if childYangName == "conform-action" {
        return &class.ConformAction
    }
    if childYangName == "exceed-action" {
        return &class.ExceedAction
    }
    if childYangName == "violate-action" {
        return &class.ViolateAction
    }
    if childYangName == "ip-mark" {
        for _, c := range class.IpMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark{}
        class.IpMark = append(class.IpMark, child)
        return &class.IpMark[len(class.IpMark)-1]
    }
    if childYangName == "common-mark" {
        for _, c := range class.CommonMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark{}
        class.CommonMark = append(class.CommonMark, child)
        return &class.CommonMark[len(class.CommonMark)-1]
    }
    if childYangName == "mpls-mark" {
        for _, c := range class.MplsMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark{}
        class.MplsMark = append(class.MplsMark, child)
        return &class.MplsMark[len(class.MplsMark)-1]
    }
    if childYangName == "wred" {
        for _, c := range class.Wred {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred{}
        class.Wred = append(class.Wred, child)
        return &class.Wred[len(class.Wred)-1]
    }
    return nil
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config-max-rate"] = &class.ConfigMaxRate
    children["config-min-rate"] = &class.ConfigMinRate
    children["config-queue-limit"] = &class.ConfigQueueLimit
    children["config-policer-average-rate"] = &class.ConfigPolicerAverageRate
    children["config-policer-peak-rate"] = &class.ConfigPolicerPeakRate
    children["config-policer-conform-burst"] = &class.ConfigPolicerConformBurst
    children["config-policer-excess-burst"] = &class.ConfigPolicerExcessBurst
    children["conform-action"] = &class.ConformAction
    children["exceed-action"] = &class.ExceedAction
    children["violate-action"] = &class.ViolateAction
    for i := range class.IpMark {
        children[class.IpMark[i].GetSegmentPath()] = &class.IpMark[i]
    }
    for i := range class.CommonMark {
        children[class.CommonMark[i].GetSegmentPath()] = &class.CommonMark[i]
    }
    for i := range class.MplsMark {
        children[class.MplsMark[i].GetSegmentPath()] = &class.MplsMark[i]
    }
    for i := range class.Wred {
        children[class.Wred[i].GetSegmentPath()] = &class.Wred[i]
    }
    return children
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level-one-class-name"] = class.LevelOneClassName
    leafs["level-two-class-name"] = class.LevelTwoClassName
    leafs["class-level"] = class.ClassLevel
    leafs["egress-queue-id"] = class.EgressQueueId
    leafs["queue-type"] = class.QueueType
    leafs["priority-level"] = class.PriorityLevel
    leafs["hardware-max-rate-kbps"] = class.HardwareMaxRateKbps
    leafs["hardware-min-rate-kbps"] = class.HardwareMinRateKbps
    leafs["config-excess-bandwidth-percent"] = class.ConfigExcessBandwidthPercent
    leafs["config-excess-bandwidth-unit"] = class.ConfigExcessBandwidthUnit
    leafs["hardware-excess-bandwidth-weight"] = class.HardwareExcessBandwidthWeight
    leafs["network-min-bandwidth-kbps"] = class.NetworkMinBandwidthKbps
    leafs["hardware-queue-limit-bytes"] = class.HardwareQueueLimitBytes
    leafs["hardware-queue-limit-microseconds"] = class.HardwareQueueLimitMicroseconds
    leafs["policer-bucket-id"] = class.PolicerBucketId
    leafs["policer-stats-handle"] = class.PolicerStatsHandle
    leafs["hardware-policer-average-rate-kbps"] = class.HardwarePolicerAverageRateKbps
    leafs["hardware-policer-peak-rate-kbps"] = class.HardwarePolicerPeakRateKbps
    leafs["hardware-policer-conform-burst-bytes"] = class.HardwarePolicerConformBurstBytes
    leafs["hardware-policer-excess-burst-bytes"] = class.HardwarePolicerExcessBurstBytes
    return leafs
}

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetYangName() string { return "class" }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetParentYangName() string { return "classes" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetFilter() yfilter.YFilter { return configMaxRate.YFilter }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) SetFilter(yf yfilter.YFilter) { configMaxRate.YFilter = yf }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetSegmentPath() string {
    return "config-max-rate"
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMaxRate.PolicyValue
    leafs["policy-unit"] = configMaxRate.PolicyUnit
    return leafs
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetBundleName() string { return "cisco_ios_xr" }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetYangName() string { return "config-max-rate" }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) SetParent(parent types.Entity) { configMaxRate.parent = parent }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetParent() types.Entity { return configMaxRate.parent }

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetFilter() yfilter.YFilter { return configMinRate.YFilter }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) SetFilter(yf yfilter.YFilter) { configMinRate.YFilter = yf }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetSegmentPath() string {
    return "config-min-rate"
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMinRate.PolicyValue
    leafs["policy-unit"] = configMinRate.PolicyUnit
    return leafs
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetBundleName() string { return "cisco_ios_xr" }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetYangName() string { return "config-min-rate" }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) SetParent(parent types.Entity) { configMinRate.parent = parent }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetParent() types.Entity { return configMinRate.parent }

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetFilter() yfilter.YFilter { return configQueueLimit.YFilter }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) SetFilter(yf yfilter.YFilter) { configQueueLimit.YFilter = yf }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetSegmentPath() string {
    return "config-queue-limit"
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configQueueLimit.PolicyValue
    leafs["policy-unit"] = configQueueLimit.PolicyUnit
    return leafs
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetBundleName() string { return "cisco_ios_xr" }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetYangName() string { return "config-queue-limit" }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) SetParent(parent types.Entity) { configQueueLimit.parent = parent }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetParent() types.Entity { return configQueueLimit.parent }

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetFilter() yfilter.YFilter { return configPolicerAverageRate.YFilter }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) SetFilter(yf yfilter.YFilter) { configPolicerAverageRate.YFilter = yf }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetSegmentPath() string {
    return "config-policer-average-rate"
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerAverageRate.PolicyValue
    leafs["policy-unit"] = configPolicerAverageRate.PolicyUnit
    return leafs
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetYangName() string { return "config-policer-average-rate" }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) SetParent(parent types.Entity) { configPolicerAverageRate.parent = parent }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetParent() types.Entity { return configPolicerAverageRate.parent }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetFilter() yfilter.YFilter { return configPolicerPeakRate.YFilter }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) SetFilter(yf yfilter.YFilter) { configPolicerPeakRate.YFilter = yf }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetSegmentPath() string {
    return "config-policer-peak-rate"
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerPeakRate.PolicyValue
    leafs["policy-unit"] = configPolicerPeakRate.PolicyUnit
    return leafs
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetYangName() string { return "config-policer-peak-rate" }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) SetParent(parent types.Entity) { configPolicerPeakRate.parent = parent }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetParent() types.Entity { return configPolicerPeakRate.parent }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetFilter() yfilter.YFilter { return configPolicerConformBurst.YFilter }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) SetFilter(yf yfilter.YFilter) { configPolicerConformBurst.YFilter = yf }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetSegmentPath() string {
    return "config-policer-conform-burst"
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerConformBurst.PolicyValue
    leafs["policy-unit"] = configPolicerConformBurst.PolicyUnit
    return leafs
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetYangName() string { return "config-policer-conform-burst" }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) SetParent(parent types.Entity) { configPolicerConformBurst.parent = parent }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetParent() types.Entity { return configPolicerConformBurst.parent }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetFilter() yfilter.YFilter { return configPolicerExcessBurst.YFilter }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) SetFilter(yf yfilter.YFilter) { configPolicerExcessBurst.YFilter = yf }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetSegmentPath() string {
    return "config-policer-excess-burst"
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerExcessBurst.PolicyValue
    leafs["policy-unit"] = configPolicerExcessBurst.PolicyUnit
    return leafs
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetYangName() string { return "config-policer-excess-burst" }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) SetParent(parent types.Entity) { configPolicerExcessBurst.parent = parent }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetParent() types.Entity { return configPolicerExcessBurst.parent }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetFilter() yfilter.YFilter { return conformAction.YFilter }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) SetFilter(yf yfilter.YFilter) { conformAction.YFilter = yf }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetSegmentPath() string {
    return "conform-action"
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range conformAction.Mark {
            if conformAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark{}
        conformAction.Mark = append(conformAction.Mark, child)
        return &conformAction.Mark[len(conformAction.Mark)-1]
    }
    return nil
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range conformAction.Mark {
        children[conformAction.Mark[i].GetSegmentPath()] = &conformAction.Mark[i]
    }
    return children
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = conformAction.ActionType
    return leafs
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetBundleName() string { return "cisco_ios_xr" }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetYangName() string { return "conform-action" }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) SetParent(parent types.Entity) { conformAction.parent = parent }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetParent() types.Entity { return conformAction.parent }

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetParentYangName() string { return "conform-action" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetFilter() yfilter.YFilter { return exceedAction.YFilter }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) SetFilter(yf yfilter.YFilter) { exceedAction.YFilter = yf }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetSegmentPath() string {
    return "exceed-action"
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range exceedAction.Mark {
            if exceedAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark{}
        exceedAction.Mark = append(exceedAction.Mark, child)
        return &exceedAction.Mark[len(exceedAction.Mark)-1]
    }
    return nil
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range exceedAction.Mark {
        children[exceedAction.Mark[i].GetSegmentPath()] = &exceedAction.Mark[i]
    }
    return children
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = exceedAction.ActionType
    return leafs
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetBundleName() string { return "cisco_ios_xr" }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetYangName() string { return "exceed-action" }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) SetParent(parent types.Entity) { exceedAction.parent = parent }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetParent() types.Entity { return exceedAction.parent }

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetParentYangName() string { return "exceed-action" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetFilter() yfilter.YFilter { return violateAction.YFilter }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) SetFilter(yf yfilter.YFilter) { violateAction.YFilter = yf }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetSegmentPath() string {
    return "violate-action"
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range violateAction.Mark {
            if violateAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark{}
        violateAction.Mark = append(violateAction.Mark, child)
        return &violateAction.Mark[len(violateAction.Mark)-1]
    }
    return nil
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range violateAction.Mark {
        children[violateAction.Mark[i].GetSegmentPath()] = &violateAction.Mark[i]
    }
    return children
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = violateAction.ActionType
    return leafs
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetBundleName() string { return "cisco_ios_xr" }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetYangName() string { return "violate-action" }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) SetParent(parent types.Entity) { violateAction.parent = parent }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetParent() types.Entity { return violateAction.parent }

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetParentYangName() string { return "violate-action" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetFilter() yfilter.YFilter { return ipMark.YFilter }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) SetFilter(yf yfilter.YFilter) { ipMark.YFilter = yf }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetSegmentPath() string {
    return "ip-mark"
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = ipMark.MarkType
    leafs["mark-value"] = ipMark.MarkValue
    return leafs
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetBundleName() string { return "cisco_ios_xr" }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetYangName() string { return "ip-mark" }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) SetParent(parent types.Entity) { ipMark.parent = parent }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetParent() types.Entity { return ipMark.parent }

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetFilter() yfilter.YFilter { return commonMark.YFilter }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) SetFilter(yf yfilter.YFilter) { commonMark.YFilter = yf }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetSegmentPath() string {
    return "common-mark"
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = commonMark.MarkType
    leafs["mark-value"] = commonMark.MarkValue
    return leafs
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetBundleName() string { return "cisco_ios_xr" }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetYangName() string { return "common-mark" }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) SetParent(parent types.Entity) { commonMark.parent = parent }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetParent() types.Entity { return commonMark.parent }

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetFilter() yfilter.YFilter { return mplsMark.YFilter }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) SetFilter(yf yfilter.YFilter) { mplsMark.YFilter = yf }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetSegmentPath() string {
    return "mpls-mark"
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mplsMark.MarkType
    leafs["mark-value"] = mplsMark.MarkValue
    return leafs
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetBundleName() string { return "cisco_ios_xr" }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetYangName() string { return "mpls-mark" }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) SetParent(parent types.Entity) { mplsMark.parent = parent }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetParent() types.Entity { return mplsMark.parent }

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // WREDMatchType. The type is DnxQoseaShowWred.
    WredMatchType interface{}

    // Hardware minimum threshold. The type is interface{} with range:
    // 0..4294967295.
    HardwareMinThresholdBytes interface{}

    // Hardware maximum threshold. The type is interface{} with range:
    // 0..4294967295.
    HardwareMaxThresholdBytes interface{}

    // First segment. The type is interface{} with range: 0..65535.
    FirstSegment interface{}

    // Segment size. The type is interface{} with range: 0..4294967295.
    SegmentSize interface{}

    // WRED match values.
    WredMatchValue PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue

    // Configured minimum threshold.
    ConfigMinThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold

    // Configured maximum threshold.
    ConfigMaxThreshold PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetFilter() yfilter.YFilter { return wred.YFilter }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) SetFilter(yf yfilter.YFilter) { wred.YFilter = yf }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetGoName(yname string) string {
    if yname == "wred-match-type" { return "WredMatchType" }
    if yname == "hardware-min-threshold-bytes" { return "HardwareMinThresholdBytes" }
    if yname == "hardware-max-threshold-bytes" { return "HardwareMaxThresholdBytes" }
    if yname == "first-segment" { return "FirstSegment" }
    if yname == "segment-size" { return "SegmentSize" }
    if yname == "wred-match-value" { return "WredMatchValue" }
    if yname == "config-min-threshold" { return "ConfigMinThreshold" }
    if yname == "config-max-threshold" { return "ConfigMaxThreshold" }
    return ""
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetSegmentPath() string {
    return "wred"
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wred-match-value" {
        return &wred.WredMatchValue
    }
    if childYangName == "config-min-threshold" {
        return &wred.ConfigMinThreshold
    }
    if childYangName == "config-max-threshold" {
        return &wred.ConfigMaxThreshold
    }
    return nil
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wred-match-value"] = &wred.WredMatchValue
    children["config-min-threshold"] = &wred.ConfigMinThreshold
    children["config-max-threshold"] = &wred.ConfigMaxThreshold
    return children
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wred-match-type"] = wred.WredMatchType
    leafs["hardware-min-threshold-bytes"] = wred.HardwareMinThresholdBytes
    leafs["hardware-max-threshold-bytes"] = wred.HardwareMaxThresholdBytes
    leafs["first-segment"] = wred.FirstSegment
    leafs["segment-size"] = wred.SegmentSize
    return leafs
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetBundleName() string { return "cisco_ios_xr" }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetYangName() string { return "wred" }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) SetParent(parent types.Entity) { wred.parent = parent }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetParent() types.Entity { return wred.parent }

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetFilter() yfilter.YFilter { return wredMatchValue.YFilter }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) SetFilter(yf yfilter.YFilter) { wredMatchValue.YFilter = yf }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetGoName(yname string) string {
    if yname == "dnx-qosea-show-red-match-value" { return "DnxQoseaShowRedMatchValue" }
    return ""
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetSegmentPath() string {
    return "wred-match-value"
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dnx-qosea-show-red-match-value" {
        for _, c := range wredMatchValue.DnxQoseaShowRedMatchValue {
            if wredMatchValue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue{}
        wredMatchValue.DnxQoseaShowRedMatchValue = append(wredMatchValue.DnxQoseaShowRedMatchValue, child)
        return &wredMatchValue.DnxQoseaShowRedMatchValue[len(wredMatchValue.DnxQoseaShowRedMatchValue)-1]
    }
    return nil
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        children[wredMatchValue.DnxQoseaShowRedMatchValue[i].GetSegmentPath()] = &wredMatchValue.DnxQoseaShowRedMatchValue[i]
    }
    return children
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetBundleName() string { return "cisco_ios_xr" }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetYangName() string { return "wred-match-value" }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) SetParent(parent types.Entity) { wredMatchValue.parent = parent }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetParent() types.Entity { return wredMatchValue.parent }

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetFilter() yfilter.YFilter { return dnxQoseaShowRedMatchValue.YFilter }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) SetFilter(yf yfilter.YFilter) { dnxQoseaShowRedMatchValue.YFilter = yf }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetGoName(yname string) string {
    if yname == "range-start" { return "RangeStart" }
    if yname == "range-end" { return "RangeEnd" }
    return ""
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetSegmentPath() string {
    return "dnx-qosea-show-red-match-value"
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["range-start"] = dnxQoseaShowRedMatchValue.RangeStart
    leafs["range-end"] = dnxQoseaShowRedMatchValue.RangeEnd
    return leafs
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetBundleName() string { return "cisco_ios_xr" }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetYangName() string { return "dnx-qosea-show-red-match-value" }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) SetParent(parent types.Entity) { dnxQoseaShowRedMatchValue.parent = parent }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetParent() types.Entity { return dnxQoseaShowRedMatchValue.parent }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetParentYangName() string { return "wred-match-value" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetFilter() yfilter.YFilter { return configMinThreshold.YFilter }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) SetFilter(yf yfilter.YFilter) { configMinThreshold.YFilter = yf }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetSegmentPath() string {
    return "config-min-threshold"
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMinThreshold.PolicyValue
    leafs["policy-unit"] = configMinThreshold.PolicyUnit
    return leafs
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetYangName() string { return "config-min-threshold" }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) SetParent(parent types.Entity) { configMinThreshold.parent = parent }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetParent() types.Entity { return configMinThreshold.parent }

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetFilter() yfilter.YFilter { return configMaxThreshold.YFilter }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) SetFilter(yf yfilter.YFilter) { configMaxThreshold.YFilter = yf }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetSegmentPath() string {
    return "config-max-threshold"
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMaxThreshold.PolicyValue
    leafs["policy-unit"] = configMaxThreshold.PolicyUnit
    return leafs
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetYangName() string { return "config-max-threshold" }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) SetParent(parent types.Entity) { configMaxThreshold.parent = parent }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetParent() types.Entity { return configMaxThreshold.parent }

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_Interfaces
// QoS list of interfaces
type PlatformQos_Nodes_Node_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS interface names. The type is slice of
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
// QoS interface names
type PlatformQos_Nodes_Node_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // The interface direction on which QoS is applied to. The type is string.
    QosDirection interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_Interfaces_Interface_Classes
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "qos-direction" { return "QosDirection" }
    if yname == "policy-details" { return "PolicyDetails" }
    if yname == "classes" { return "Classes" }
    return ""
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "policy-details" {
        return &self.PolicyDetails
    }
    if childYangName == "classes" {
        return &self.Classes
    }
    return nil
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["policy-details"] = &self.PolicyDetails
    children["classes"] = &self.Classes
    return children
}

func (self *PlatformQos_Nodes_Node_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["qos-direction"] = self.QosDirection
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

// PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NPU ID. The type is interface{} with range: 0..4294967295.
    NpuId interface{}

    // InterfaceHandle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Interface Bandwidth (in kbps). The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    InterfaceBandwidthKbps interface{}

    // Policy name. The type is string with length: 0..64.
    PolicyName interface{}

    // Number of Classes. The type is interface{} with range: 0..65535.
    TotalNumberOfClasses interface{}

    // VOQ base address. The type is interface{} with range: 0..4294967295.
    VoqBaseAddress interface{}

    // VOQ stats handle. The type is interface{} with range:
    // 0..18446744073709551615.
    VoqStatsHandle interface{}

    // QoS Statistics Accounting Type. The type is QosPolicyAccountEnum.
    StatsAccountingType interface{}

    // Policy Status. The type is DnxQoseaShowPolicyStatus.
    PolicyStatus interface{}

    // Interface Status. The type is DnxQoseaShowIntfStatus.
    InterfaceStatus interface{}
}

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetFilter() yfilter.YFilter { return policyDetails.YFilter }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) SetFilter(yf yfilter.YFilter) { policyDetails.YFilter = yf }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetGoName(yname string) string {
    if yname == "npu-id" { return "NpuId" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "interface-bandwidth-kbps" { return "InterfaceBandwidthKbps" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "total-number-of-classes" { return "TotalNumberOfClasses" }
    if yname == "voq-base-address" { return "VoqBaseAddress" }
    if yname == "voq-stats-handle" { return "VoqStatsHandle" }
    if yname == "stats-accounting-type" { return "StatsAccountingType" }
    if yname == "policy-status" { return "PolicyStatus" }
    if yname == "interface-status" { return "InterfaceStatus" }
    return ""
}

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetSegmentPath() string {
    return "policy-details"
}

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["npu-id"] = policyDetails.NpuId
    leafs["interface-handle"] = policyDetails.InterfaceHandle
    leafs["interface-bandwidth-kbps"] = policyDetails.InterfaceBandwidthKbps
    leafs["policy-name"] = policyDetails.PolicyName
    leafs["total-number-of-classes"] = policyDetails.TotalNumberOfClasses
    leafs["voq-base-address"] = policyDetails.VoqBaseAddress
    leafs["voq-stats-handle"] = policyDetails.VoqStatsHandle
    leafs["stats-accounting-type"] = policyDetails.StatsAccountingType
    leafs["policy-status"] = policyDetails.PolicyStatus
    leafs["interface-status"] = policyDetails.InterfaceStatus
    return leafs
}

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetBundleName() string { return "cisco_ios_xr" }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetYangName() string { return "policy-details" }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) SetParent(parent types.Entity) { policyDetails.parent = parent }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetParent() types.Entity { return policyDetails.parent }

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetParentYangName() string { return "interface" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class.
    Class []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetYangName() string { return "classes" }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetParent() types.Entity { return classes.parent }

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetParentYangName() string { return "interface" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. QoS policy class name at level 1. The type is
    // string.
    LevelOneClassName interface{}

    // QoS policy child class name at level 2. The type is string.
    LevelTwoClassName interface{}

    // Class level. The type is DnxQoseaShowLevel.
    ClassLevel interface{}

    // Egress Queue ID. The type is interface{} with range:
    // -2147483648..2147483647.
    EgressQueueId interface{}

    // Queue type. The type is DnxQoseaShowQueue.
    QueueType interface{}

    // Priority level. The type is DnxQoseaShowHpLevel.
    PriorityLevel interface{}

    // Hardware maximum rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwareMaxRateKbps interface{}

    // Hardware minimum rate in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwareMinRateKbps interface{}

    // Configured excess bandwidth percentage. The type is interface{} with range:
    // 0..4294967295. Units are percentage.
    ConfigExcessBandwidthPercent interface{}

    // Configured excess bandwidth unit. The type is interface{} with range:
    // 0..4294967295.
    ConfigExcessBandwidthUnit interface{}

    // Hardware excess bandwidth weight. The type is interface{} with range:
    // 0..4294967295.
    HardwareExcessBandwidthWeight interface{}

    // Network minimum Bandwith. The type is interface{} with range:
    // 0..4294967295.
    NetworkMinBandwidthKbps interface{}

    // Hardware queue limit in bytes. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    HardwareQueueLimitBytes interface{}

    // Hardware queue limit in microseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are microsecond.
    HardwareQueueLimitMicroseconds interface{}

    // PolicerBucketID. The type is interface{} with range: 0..4294967295.
    PolicerBucketId interface{}

    // PolicerStatsHandle. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicerStatsHandle interface{}

    // Hardware policer average in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    HardwarePolicerAverageRateKbps interface{}

    // Hardware policer peak rate. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerPeakRateKbps interface{}

    // Hardware policer conform burst. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerConformBurstBytes interface{}

    // Hardware policer excess burst. The type is interface{} with range:
    // 0..4294967295.
    HardwarePolicerExcessBurstBytes interface{}

    // Configured maximum rate.
    ConfigMaxRate PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate

    // Configured minimum rate.
    ConfigMinRate PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate

    // Configured queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit

    // Configured policer average rate.
    ConfigPolicerAverageRate PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate

    // Config policer peak rate.
    ConfigPolicerPeakRate PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate

    // Configured policer conform burst.
    ConfigPolicerConformBurst PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst

    // Configured policer excess burst.
    ConfigPolicerExcessBurst PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst

    // Conform action.
    ConformAction PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction

    // Exceed action.
    ExceedAction PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction

    // Violate action.
    ViolateAction PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction

    // IP mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark.
    IpMark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark

    // Common mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark.
    CommonMark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark

    // MPLS mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark.
    MplsMark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark

    // WRED parameters. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred.
    Wred []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred
}

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetGoName(yname string) string {
    if yname == "level-one-class-name" { return "LevelOneClassName" }
    if yname == "level-two-class-name" { return "LevelTwoClassName" }
    if yname == "class-level" { return "ClassLevel" }
    if yname == "egress-queue-id" { return "EgressQueueId" }
    if yname == "queue-type" { return "QueueType" }
    if yname == "priority-level" { return "PriorityLevel" }
    if yname == "hardware-max-rate-kbps" { return "HardwareMaxRateKbps" }
    if yname == "hardware-min-rate-kbps" { return "HardwareMinRateKbps" }
    if yname == "config-excess-bandwidth-percent" { return "ConfigExcessBandwidthPercent" }
    if yname == "config-excess-bandwidth-unit" { return "ConfigExcessBandwidthUnit" }
    if yname == "hardware-excess-bandwidth-weight" { return "HardwareExcessBandwidthWeight" }
    if yname == "network-min-bandwidth-kbps" { return "NetworkMinBandwidthKbps" }
    if yname == "hardware-queue-limit-bytes" { return "HardwareQueueLimitBytes" }
    if yname == "hardware-queue-limit-microseconds" { return "HardwareQueueLimitMicroseconds" }
    if yname == "policer-bucket-id" { return "PolicerBucketId" }
    if yname == "policer-stats-handle" { return "PolicerStatsHandle" }
    if yname == "hardware-policer-average-rate-kbps" { return "HardwarePolicerAverageRateKbps" }
    if yname == "hardware-policer-peak-rate-kbps" { return "HardwarePolicerPeakRateKbps" }
    if yname == "hardware-policer-conform-burst-bytes" { return "HardwarePolicerConformBurstBytes" }
    if yname == "hardware-policer-excess-burst-bytes" { return "HardwarePolicerExcessBurstBytes" }
    if yname == "config-max-rate" { return "ConfigMaxRate" }
    if yname == "config-min-rate" { return "ConfigMinRate" }
    if yname == "config-queue-limit" { return "ConfigQueueLimit" }
    if yname == "config-policer-average-rate" { return "ConfigPolicerAverageRate" }
    if yname == "config-policer-peak-rate" { return "ConfigPolicerPeakRate" }
    if yname == "config-policer-conform-burst" { return "ConfigPolicerConformBurst" }
    if yname == "config-policer-excess-burst" { return "ConfigPolicerExcessBurst" }
    if yname == "conform-action" { return "ConformAction" }
    if yname == "exceed-action" { return "ExceedAction" }
    if yname == "violate-action" { return "ViolateAction" }
    if yname == "ip-mark" { return "IpMark" }
    if yname == "common-mark" { return "CommonMark" }
    if yname == "mpls-mark" { return "MplsMark" }
    if yname == "wred" { return "Wred" }
    return ""
}

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetSegmentPath() string {
    return "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
}

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-max-rate" {
        return &class.ConfigMaxRate
    }
    if childYangName == "config-min-rate" {
        return &class.ConfigMinRate
    }
    if childYangName == "config-queue-limit" {
        return &class.ConfigQueueLimit
    }
    if childYangName == "config-policer-average-rate" {
        return &class.ConfigPolicerAverageRate
    }
    if childYangName == "config-policer-peak-rate" {
        return &class.ConfigPolicerPeakRate
    }
    if childYangName == "config-policer-conform-burst" {
        return &class.ConfigPolicerConformBurst
    }
    if childYangName == "config-policer-excess-burst" {
        return &class.ConfigPolicerExcessBurst
    }
    if childYangName == "conform-action" {
        return &class.ConformAction
    }
    if childYangName == "exceed-action" {
        return &class.ExceedAction
    }
    if childYangName == "violate-action" {
        return &class.ViolateAction
    }
    if childYangName == "ip-mark" {
        for _, c := range class.IpMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark{}
        class.IpMark = append(class.IpMark, child)
        return &class.IpMark[len(class.IpMark)-1]
    }
    if childYangName == "common-mark" {
        for _, c := range class.CommonMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark{}
        class.CommonMark = append(class.CommonMark, child)
        return &class.CommonMark[len(class.CommonMark)-1]
    }
    if childYangName == "mpls-mark" {
        for _, c := range class.MplsMark {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark{}
        class.MplsMark = append(class.MplsMark, child)
        return &class.MplsMark[len(class.MplsMark)-1]
    }
    if childYangName == "wred" {
        for _, c := range class.Wred {
            if class.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred{}
        class.Wred = append(class.Wred, child)
        return &class.Wred[len(class.Wred)-1]
    }
    return nil
}

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config-max-rate"] = &class.ConfigMaxRate
    children["config-min-rate"] = &class.ConfigMinRate
    children["config-queue-limit"] = &class.ConfigQueueLimit
    children["config-policer-average-rate"] = &class.ConfigPolicerAverageRate
    children["config-policer-peak-rate"] = &class.ConfigPolicerPeakRate
    children["config-policer-conform-burst"] = &class.ConfigPolicerConformBurst
    children["config-policer-excess-burst"] = &class.ConfigPolicerExcessBurst
    children["conform-action"] = &class.ConformAction
    children["exceed-action"] = &class.ExceedAction
    children["violate-action"] = &class.ViolateAction
    for i := range class.IpMark {
        children[class.IpMark[i].GetSegmentPath()] = &class.IpMark[i]
    }
    for i := range class.CommonMark {
        children[class.CommonMark[i].GetSegmentPath()] = &class.CommonMark[i]
    }
    for i := range class.MplsMark {
        children[class.MplsMark[i].GetSegmentPath()] = &class.MplsMark[i]
    }
    for i := range class.Wred {
        children[class.Wred[i].GetSegmentPath()] = &class.Wred[i]
    }
    return children
}

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["level-one-class-name"] = class.LevelOneClassName
    leafs["level-two-class-name"] = class.LevelTwoClassName
    leafs["class-level"] = class.ClassLevel
    leafs["egress-queue-id"] = class.EgressQueueId
    leafs["queue-type"] = class.QueueType
    leafs["priority-level"] = class.PriorityLevel
    leafs["hardware-max-rate-kbps"] = class.HardwareMaxRateKbps
    leafs["hardware-min-rate-kbps"] = class.HardwareMinRateKbps
    leafs["config-excess-bandwidth-percent"] = class.ConfigExcessBandwidthPercent
    leafs["config-excess-bandwidth-unit"] = class.ConfigExcessBandwidthUnit
    leafs["hardware-excess-bandwidth-weight"] = class.HardwareExcessBandwidthWeight
    leafs["network-min-bandwidth-kbps"] = class.NetworkMinBandwidthKbps
    leafs["hardware-queue-limit-bytes"] = class.HardwareQueueLimitBytes
    leafs["hardware-queue-limit-microseconds"] = class.HardwareQueueLimitMicroseconds
    leafs["policer-bucket-id"] = class.PolicerBucketId
    leafs["policer-stats-handle"] = class.PolicerStatsHandle
    leafs["hardware-policer-average-rate-kbps"] = class.HardwarePolicerAverageRateKbps
    leafs["hardware-policer-peak-rate-kbps"] = class.HardwarePolicerPeakRateKbps
    leafs["hardware-policer-conform-burst-bytes"] = class.HardwarePolicerConformBurstBytes
    leafs["hardware-policer-excess-burst-bytes"] = class.HardwarePolicerExcessBurstBytes
    return leafs
}

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetYangName() string { return "class" }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetParentYangName() string { return "classes" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetFilter() yfilter.YFilter { return configMaxRate.YFilter }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) SetFilter(yf yfilter.YFilter) { configMaxRate.YFilter = yf }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetSegmentPath() string {
    return "config-max-rate"
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMaxRate.PolicyValue
    leafs["policy-unit"] = configMaxRate.PolicyUnit
    return leafs
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetBundleName() string { return "cisco_ios_xr" }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetYangName() string { return "config-max-rate" }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) SetParent(parent types.Entity) { configMaxRate.parent = parent }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetParent() types.Entity { return configMaxRate.parent }

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetFilter() yfilter.YFilter { return configMinRate.YFilter }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) SetFilter(yf yfilter.YFilter) { configMinRate.YFilter = yf }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetSegmentPath() string {
    return "config-min-rate"
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMinRate.PolicyValue
    leafs["policy-unit"] = configMinRate.PolicyUnit
    return leafs
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetBundleName() string { return "cisco_ios_xr" }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetYangName() string { return "config-min-rate" }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) SetParent(parent types.Entity) { configMinRate.parent = parent }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetParent() types.Entity { return configMinRate.parent }

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetFilter() yfilter.YFilter { return configQueueLimit.YFilter }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) SetFilter(yf yfilter.YFilter) { configQueueLimit.YFilter = yf }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetSegmentPath() string {
    return "config-queue-limit"
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configQueueLimit.PolicyValue
    leafs["policy-unit"] = configQueueLimit.PolicyUnit
    return leafs
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetBundleName() string { return "cisco_ios_xr" }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetYangName() string { return "config-queue-limit" }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) SetParent(parent types.Entity) { configQueueLimit.parent = parent }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetParent() types.Entity { return configQueueLimit.parent }

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetFilter() yfilter.YFilter { return configPolicerAverageRate.YFilter }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) SetFilter(yf yfilter.YFilter) { configPolicerAverageRate.YFilter = yf }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetSegmentPath() string {
    return "config-policer-average-rate"
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerAverageRate.PolicyValue
    leafs["policy-unit"] = configPolicerAverageRate.PolicyUnit
    return leafs
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetYangName() string { return "config-policer-average-rate" }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) SetParent(parent types.Entity) { configPolicerAverageRate.parent = parent }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetParent() types.Entity { return configPolicerAverageRate.parent }

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetFilter() yfilter.YFilter { return configPolicerPeakRate.YFilter }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) SetFilter(yf yfilter.YFilter) { configPolicerPeakRate.YFilter = yf }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetSegmentPath() string {
    return "config-policer-peak-rate"
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerPeakRate.PolicyValue
    leafs["policy-unit"] = configPolicerPeakRate.PolicyUnit
    return leafs
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetYangName() string { return "config-policer-peak-rate" }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) SetParent(parent types.Entity) { configPolicerPeakRate.parent = parent }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetParent() types.Entity { return configPolicerPeakRate.parent }

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetFilter() yfilter.YFilter { return configPolicerConformBurst.YFilter }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) SetFilter(yf yfilter.YFilter) { configPolicerConformBurst.YFilter = yf }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetSegmentPath() string {
    return "config-policer-conform-burst"
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerConformBurst.PolicyValue
    leafs["policy-unit"] = configPolicerConformBurst.PolicyUnit
    return leafs
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetYangName() string { return "config-policer-conform-burst" }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) SetParent(parent types.Entity) { configPolicerConformBurst.parent = parent }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetParent() types.Entity { return configPolicerConformBurst.parent }

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetFilter() yfilter.YFilter { return configPolicerExcessBurst.YFilter }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) SetFilter(yf yfilter.YFilter) { configPolicerExcessBurst.YFilter = yf }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetSegmentPath() string {
    return "config-policer-excess-burst"
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configPolicerExcessBurst.PolicyValue
    leafs["policy-unit"] = configPolicerExcessBurst.PolicyUnit
    return leafs
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetBundleName() string { return "cisco_ios_xr" }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetYangName() string { return "config-policer-excess-burst" }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) SetParent(parent types.Entity) { configPolicerExcessBurst.parent = parent }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetParent() types.Entity { return configPolicerExcessBurst.parent }

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetFilter() yfilter.YFilter { return conformAction.YFilter }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) SetFilter(yf yfilter.YFilter) { conformAction.YFilter = yf }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetSegmentPath() string {
    return "conform-action"
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range conformAction.Mark {
            if conformAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark{}
        conformAction.Mark = append(conformAction.Mark, child)
        return &conformAction.Mark[len(conformAction.Mark)-1]
    }
    return nil
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range conformAction.Mark {
        children[conformAction.Mark[i].GetSegmentPath()] = &conformAction.Mark[i]
    }
    return children
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = conformAction.ActionType
    return leafs
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetBundleName() string { return "cisco_ios_xr" }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetYangName() string { return "conform-action" }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) SetParent(parent types.Entity) { conformAction.parent = parent }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetParent() types.Entity { return conformAction.parent }

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetParentYangName() string { return "conform-action" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetFilter() yfilter.YFilter { return exceedAction.YFilter }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) SetFilter(yf yfilter.YFilter) { exceedAction.YFilter = yf }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetSegmentPath() string {
    return "exceed-action"
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range exceedAction.Mark {
            if exceedAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark{}
        exceedAction.Mark = append(exceedAction.Mark, child)
        return &exceedAction.Mark[len(exceedAction.Mark)-1]
    }
    return nil
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range exceedAction.Mark {
        children[exceedAction.Mark[i].GetSegmentPath()] = &exceedAction.Mark[i]
    }
    return children
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = exceedAction.ActionType
    return leafs
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetBundleName() string { return "cisco_ios_xr" }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetYangName() string { return "exceed-action" }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) SetParent(parent types.Entity) { exceedAction.parent = parent }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetParent() types.Entity { return exceedAction.parent }

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetParentYangName() string { return "exceed-action" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetFilter() yfilter.YFilter { return violateAction.YFilter }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) SetFilter(yf yfilter.YFilter) { violateAction.YFilter = yf }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetGoName(yname string) string {
    if yname == "action-type" { return "ActionType" }
    if yname == "mark" { return "Mark" }
    return ""
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetSegmentPath() string {
    return "violate-action"
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mark" {
        for _, c := range violateAction.Mark {
            if violateAction.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark{}
        violateAction.Mark = append(violateAction.Mark, child)
        return &violateAction.Mark[len(violateAction.Mark)-1]
    }
    return nil
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range violateAction.Mark {
        children[violateAction.Mark[i].GetSegmentPath()] = &violateAction.Mark[i]
    }
    return children
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["action-type"] = violateAction.ActionType
    return leafs
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetBundleName() string { return "cisco_ios_xr" }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetYangName() string { return "violate-action" }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) SetParent(parent types.Entity) { violateAction.parent = parent }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetParent() types.Entity { return violateAction.parent }

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetFilter() yfilter.YFilter { return mark.YFilter }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) SetFilter(yf yfilter.YFilter) { mark.YFilter = yf }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetSegmentPath() string {
    return "mark"
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mark.MarkType
    leafs["mark-value"] = mark.MarkValue
    return leafs
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetBundleName() string { return "cisco_ios_xr" }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetYangName() string { return "mark" }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) SetParent(parent types.Entity) { mark.parent = parent }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetParent() types.Entity { return mark.parent }

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetParentYangName() string { return "violate-action" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetFilter() yfilter.YFilter { return ipMark.YFilter }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) SetFilter(yf yfilter.YFilter) { ipMark.YFilter = yf }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetSegmentPath() string {
    return "ip-mark"
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = ipMark.MarkType
    leafs["mark-value"] = ipMark.MarkValue
    return leafs
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetBundleName() string { return "cisco_ios_xr" }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetYangName() string { return "ip-mark" }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) SetParent(parent types.Entity) { ipMark.parent = parent }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetParent() types.Entity { return ipMark.parent }

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetFilter() yfilter.YFilter { return commonMark.YFilter }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) SetFilter(yf yfilter.YFilter) { commonMark.YFilter = yf }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetSegmentPath() string {
    return "common-mark"
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = commonMark.MarkType
    leafs["mark-value"] = commonMark.MarkValue
    return leafs
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetBundleName() string { return "cisco_ios_xr" }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetYangName() string { return "common-mark" }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) SetParent(parent types.Entity) { commonMark.parent = parent }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetParent() types.Entity { return commonMark.parent }

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetFilter() yfilter.YFilter { return mplsMark.YFilter }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) SetFilter(yf yfilter.YFilter) { mplsMark.YFilter = yf }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetGoName(yname string) string {
    if yname == "mark-type" { return "MarkType" }
    if yname == "mark-value" { return "MarkValue" }
    return ""
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetSegmentPath() string {
    return "mpls-mark"
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-type"] = mplsMark.MarkType
    leafs["mark-value"] = mplsMark.MarkValue
    return leafs
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetBundleName() string { return "cisco_ios_xr" }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetYangName() string { return "mpls-mark" }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) SetParent(parent types.Entity) { mplsMark.parent = parent }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetParent() types.Entity { return mplsMark.parent }

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // WREDMatchType. The type is DnxQoseaShowWred.
    WredMatchType interface{}

    // Hardware minimum threshold. The type is interface{} with range:
    // 0..4294967295.
    HardwareMinThresholdBytes interface{}

    // Hardware maximum threshold. The type is interface{} with range:
    // 0..4294967295.
    HardwareMaxThresholdBytes interface{}

    // First segment. The type is interface{} with range: 0..65535.
    FirstSegment interface{}

    // Segment size. The type is interface{} with range: 0..4294967295.
    SegmentSize interface{}

    // WRED match values.
    WredMatchValue PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue

    // Configured minimum threshold.
    ConfigMinThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold

    // Configured maximum threshold.
    ConfigMaxThreshold PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetFilter() yfilter.YFilter { return wred.YFilter }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) SetFilter(yf yfilter.YFilter) { wred.YFilter = yf }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetGoName(yname string) string {
    if yname == "wred-match-type" { return "WredMatchType" }
    if yname == "hardware-min-threshold-bytes" { return "HardwareMinThresholdBytes" }
    if yname == "hardware-max-threshold-bytes" { return "HardwareMaxThresholdBytes" }
    if yname == "first-segment" { return "FirstSegment" }
    if yname == "segment-size" { return "SegmentSize" }
    if yname == "wred-match-value" { return "WredMatchValue" }
    if yname == "config-min-threshold" { return "ConfigMinThreshold" }
    if yname == "config-max-threshold" { return "ConfigMaxThreshold" }
    return ""
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetSegmentPath() string {
    return "wred"
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wred-match-value" {
        return &wred.WredMatchValue
    }
    if childYangName == "config-min-threshold" {
        return &wred.ConfigMinThreshold
    }
    if childYangName == "config-max-threshold" {
        return &wred.ConfigMaxThreshold
    }
    return nil
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["wred-match-value"] = &wred.WredMatchValue
    children["config-min-threshold"] = &wred.ConfigMinThreshold
    children["config-max-threshold"] = &wred.ConfigMaxThreshold
    return children
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["wred-match-type"] = wred.WredMatchType
    leafs["hardware-min-threshold-bytes"] = wred.HardwareMinThresholdBytes
    leafs["hardware-max-threshold-bytes"] = wred.HardwareMaxThresholdBytes
    leafs["first-segment"] = wred.FirstSegment
    leafs["segment-size"] = wred.SegmentSize
    return leafs
}

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetBundleName() string { return "cisco_ios_xr" }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetYangName() string { return "wred" }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) SetParent(parent types.Entity) { wred.parent = parent }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetParent() types.Entity { return wred.parent }

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetParentYangName() string { return "class" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetFilter() yfilter.YFilter { return wredMatchValue.YFilter }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) SetFilter(yf yfilter.YFilter) { wredMatchValue.YFilter = yf }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetGoName(yname string) string {
    if yname == "dnx-qosea-show-red-match-value" { return "DnxQoseaShowRedMatchValue" }
    return ""
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetSegmentPath() string {
    return "wred-match-value"
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dnx-qosea-show-red-match-value" {
        for _, c := range wredMatchValue.DnxQoseaShowRedMatchValue {
            if wredMatchValue.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue{}
        wredMatchValue.DnxQoseaShowRedMatchValue = append(wredMatchValue.DnxQoseaShowRedMatchValue, child)
        return &wredMatchValue.DnxQoseaShowRedMatchValue[len(wredMatchValue.DnxQoseaShowRedMatchValue)-1]
    }
    return nil
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        children[wredMatchValue.DnxQoseaShowRedMatchValue[i].GetSegmentPath()] = &wredMatchValue.DnxQoseaShowRedMatchValue[i]
    }
    return children
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetBundleName() string { return "cisco_ios_xr" }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetYangName() string { return "wred-match-value" }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) SetParent(parent types.Entity) { wredMatchValue.parent = parent }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetParent() types.Entity { return wredMatchValue.parent }

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetFilter() yfilter.YFilter { return dnxQoseaShowRedMatchValue.YFilter }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) SetFilter(yf yfilter.YFilter) { dnxQoseaShowRedMatchValue.YFilter = yf }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetGoName(yname string) string {
    if yname == "range-start" { return "RangeStart" }
    if yname == "range-end" { return "RangeEnd" }
    return ""
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetSegmentPath() string {
    return "dnx-qosea-show-red-match-value"
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["range-start"] = dnxQoseaShowRedMatchValue.RangeStart
    leafs["range-end"] = dnxQoseaShowRedMatchValue.RangeEnd
    return leafs
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetBundleName() string { return "cisco_ios_xr" }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetYangName() string { return "dnx-qosea-show-red-match-value" }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) SetParent(parent types.Entity) { dnxQoseaShowRedMatchValue.parent = parent }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetParent() types.Entity { return dnxQoseaShowRedMatchValue.parent }

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetParentYangName() string { return "wred-match-value" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetFilter() yfilter.YFilter { return configMinThreshold.YFilter }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) SetFilter(yf yfilter.YFilter) { configMinThreshold.YFilter = yf }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetSegmentPath() string {
    return "config-min-threshold"
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMinThreshold.PolicyValue
    leafs["policy-unit"] = configMinThreshold.PolicyUnit
    return leafs
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetYangName() string { return "config-min-threshold" }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) SetParent(parent types.Entity) { configMinThreshold.parent = parent }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetParent() types.Entity { return configMinThreshold.parent }

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetFilter() yfilter.YFilter { return configMaxThreshold.YFilter }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) SetFilter(yf yfilter.YFilter) { configMaxThreshold.YFilter = yf }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetGoName(yname string) string {
    if yname == "policy-value" { return "PolicyValue" }
    if yname == "policy-unit" { return "PolicyUnit" }
    return ""
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetSegmentPath() string {
    return "config-max-threshold"
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy-value"] = configMaxThreshold.PolicyValue
    leafs["policy-unit"] = configMaxThreshold.PolicyUnit
    return leafs
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetBundleName() string { return "cisco_ios_xr" }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetYangName() string { return "config-max-threshold" }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) SetParent(parent types.Entity) { configMaxThreshold.parent = parent }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetParent() types.Entity { return configMaxThreshold.parent }

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetParentYangName() string { return "wred" }

// PlatformQos_Nodes_Node_RemoteInterfaces
// QoS list of remote interfaces
type PlatformQos_Nodes_Node_RemoteInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // QoS remote interface names. The type is slice of
    // PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface.
    RemoteInterface []PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetFilter() yfilter.YFilter { return remoteInterfaces.YFilter }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) SetFilter(yf yfilter.YFilter) { remoteInterfaces.YFilter = yf }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetGoName(yname string) string {
    if yname == "remote-interface" { return "RemoteInterface" }
    return ""
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetSegmentPath() string {
    return "remote-interfaces"
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-interface" {
        for _, c := range remoteInterfaces.RemoteInterface {
            if remoteInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface{}
        remoteInterfaces.RemoteInterface = append(remoteInterfaces.RemoteInterface, child)
        return &remoteInterfaces.RemoteInterface[len(remoteInterfaces.RemoteInterface)-1]
    }
    return nil
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteInterfaces.RemoteInterface {
        children[remoteInterfaces.RemoteInterface[i].GetSegmentPath()] = &remoteInterfaces.RemoteInterface[i]
    }
    return children
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetYangName() string { return "remote-interfaces" }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) SetParent(parent types.Entity) { remoteInterfaces.parent = parent }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetParent() types.Entity { return remoteInterfaces.parent }

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetParentYangName() string { return "node" }

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface
// QoS remote interface names
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the remote interface. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Policy Name. The type is string with length: 0..64.
    PolicyName interface{}

    // Virtual output queue statistics handle. The type is interface{} with range:
    // 0..18446744073709551615.
    VirtualOutputQueueStatisticsHandle interface{}

    // Interface Handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Number of Virtual Output Queues. The type is interface{} with range:
    // 0..4294967295.
    NumberOfVirtualOutputQueues interface{}

    // Number of Classes. The type is interface{} with range: 0..4294967295.
    NumberOfClasses interface{}

    // Remote Class array. The type is slice of
    // PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass.
    RemoteClass []PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass
}

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetFilter() yfilter.YFilter { return remoteInterface.YFilter }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) SetFilter(yf yfilter.YFilter) { remoteInterface.YFilter = yf }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "policy-name" { return "PolicyName" }
    if yname == "virtual-output-queue-statistics-handle" { return "VirtualOutputQueueStatisticsHandle" }
    if yname == "interface-handle" { return "InterfaceHandle" }
    if yname == "number-of-virtual-output-queues" { return "NumberOfVirtualOutputQueues" }
    if yname == "number-of-classes" { return "NumberOfClasses" }
    if yname == "remote-class" { return "RemoteClass" }
    return ""
}

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetSegmentPath() string {
    return "remote-interface" + "[interface-name='" + fmt.Sprintf("%v", remoteInterface.InterfaceName) + "']"
}

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-class" {
        for _, c := range remoteInterface.RemoteClass {
            if remoteInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass{}
        remoteInterface.RemoteClass = append(remoteInterface.RemoteClass, child)
        return &remoteInterface.RemoteClass[len(remoteInterface.RemoteClass)-1]
    }
    return nil
}

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteInterface.RemoteClass {
        children[remoteInterface.RemoteClass[i].GetSegmentPath()] = &remoteInterface.RemoteClass[i]
    }
    return children
}

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = remoteInterface.InterfaceName
    leafs["policy-name"] = remoteInterface.PolicyName
    leafs["virtual-output-queue-statistics-handle"] = remoteInterface.VirtualOutputQueueStatisticsHandle
    leafs["interface-handle"] = remoteInterface.InterfaceHandle
    leafs["number-of-virtual-output-queues"] = remoteInterface.NumberOfVirtualOutputQueues
    leafs["number-of-classes"] = remoteInterface.NumberOfClasses
    return leafs
}

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetBundleName() string { return "cisco_ios_xr" }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetYangName() string { return "remote-interface" }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) SetParent(parent types.Entity) { remoteInterface.parent = parent }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetParent() types.Entity { return remoteInterface.parent }

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetParentYangName() string { return "remote-interfaces" }

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass
// Remote Class array
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Class Name. The type is string with length: 0..64.
    ClassName interface{}

    // Class ID. The type is interface{} with range: 0..4294967295.
    ClassId interface{}

    // Class of Service Queue. The type is interface{} with range: 0..4294967295.
    CosQ interface{}

    // Default/Configured queue limit in bytes. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    QueueLimit interface{}

    // Hardware queue limit in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    HardwareQueueLimit interface{}

    // Default/Configured WRED profiles. The type is slice of
    // PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred.
    Wred []PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred

    // Hardware WRED profiles. The type is slice of
    // PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred.
    HwWred []PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred
}

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetFilter() yfilter.YFilter { return remoteClass.YFilter }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) SetFilter(yf yfilter.YFilter) { remoteClass.YFilter = yf }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "class-id" { return "ClassId" }
    if yname == "cos-q" { return "CosQ" }
    if yname == "queue-limit" { return "QueueLimit" }
    if yname == "hardware-queue-limit" { return "HardwareQueueLimit" }
    if yname == "wred" { return "Wred" }
    if yname == "hw-wred" { return "HwWred" }
    return ""
}

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetSegmentPath() string {
    return "remote-class"
}

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "wred" {
        for _, c := range remoteClass.Wred {
            if remoteClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred{}
        remoteClass.Wred = append(remoteClass.Wred, child)
        return &remoteClass.Wred[len(remoteClass.Wred)-1]
    }
    if childYangName == "hw-wred" {
        for _, c := range remoteClass.HwWred {
            if remoteClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred{}
        remoteClass.HwWred = append(remoteClass.HwWred, child)
        return &remoteClass.HwWred[len(remoteClass.HwWred)-1]
    }
    return nil
}

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteClass.Wred {
        children[remoteClass.Wred[i].GetSegmentPath()] = &remoteClass.Wred[i]
    }
    for i := range remoteClass.HwWred {
        children[remoteClass.HwWred[i].GetSegmentPath()] = &remoteClass.HwWred[i]
    }
    return children
}

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = remoteClass.ClassName
    leafs["class-id"] = remoteClass.ClassId
    leafs["cos-q"] = remoteClass.CosQ
    leafs["queue-limit"] = remoteClass.QueueLimit
    leafs["hardware-queue-limit"] = remoteClass.HardwareQueueLimit
    return leafs
}

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetBundleName() string { return "cisco_ios_xr" }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetYangName() string { return "remote-class" }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) SetParent(parent types.Entity) { remoteClass.parent = parent }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetParent() types.Entity { return remoteClass.parent }

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetParentYangName() string { return "remote-interface" }

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred
// Default/Configured WRED profiles
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum Threshold. The type is interface{} with range: 0..4294967295.
    MinThreshold interface{}

    // Maximum Threshold. The type is interface{} with range: 0..4294967295.
    MaxThreshold interface{}

    // Drop Probability. The type is interface{} with range: 0..4294967295.
    DropProbability interface{}
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetFilter() yfilter.YFilter { return wred.YFilter }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) SetFilter(yf yfilter.YFilter) { wred.YFilter = yf }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetGoName(yname string) string {
    if yname == "min-threshold" { return "MinThreshold" }
    if yname == "max-threshold" { return "MaxThreshold" }
    if yname == "drop-probability" { return "DropProbability" }
    return ""
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetSegmentPath() string {
    return "wred"
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min-threshold"] = wred.MinThreshold
    leafs["max-threshold"] = wred.MaxThreshold
    leafs["drop-probability"] = wred.DropProbability
    return leafs
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetBundleName() string { return "cisco_ios_xr" }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetYangName() string { return "wred" }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) SetParent(parent types.Entity) { wred.parent = parent }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetParent() types.Entity { return wred.parent }

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetParentYangName() string { return "remote-class" }

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred
// Hardware WRED profiles
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum Threshold. The type is interface{} with range: 0..4294967295.
    MinThreshold interface{}

    // Maximum Threshold. The type is interface{} with range: 0..4294967295.
    MaxThreshold interface{}

    // Drop Probability. The type is interface{} with range: 0..4294967295.
    DropProbability interface{}
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetFilter() yfilter.YFilter { return hwWred.YFilter }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) SetFilter(yf yfilter.YFilter) { hwWred.YFilter = yf }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetGoName(yname string) string {
    if yname == "min-threshold" { return "MinThreshold" }
    if yname == "max-threshold" { return "MaxThreshold" }
    if yname == "drop-probability" { return "DropProbability" }
    return ""
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetSegmentPath() string {
    return "hw-wred"
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min-threshold"] = hwWred.MinThreshold
    leafs["max-threshold"] = hwWred.MaxThreshold
    leafs["drop-probability"] = hwWred.DropProbability
    return leafs
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetBundleName() string { return "cisco_ios_xr" }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetYangName() string { return "hw-wred" }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) SetParent(parent types.Entity) { hwWred.parent = parent }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetParent() types.Entity { return hwWred.parent }

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetParentYangName() string { return "remote-class" }

