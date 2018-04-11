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

// DnxQoseaShowIntfStatus represents Intf Status
type DnxQoseaShowIntfStatus string

const (
    // State is unknown
    DnxQoseaShowIntfStatus_state_unknown DnxQoseaShowIntfStatus = "state-unknown"

    // State is Down
    DnxQoseaShowIntfStatus_state_down DnxQoseaShowIntfStatus = "state-down"
)

// DnxQoseaShowPolicyStatus represents Status
type DnxQoseaShowPolicyStatus string

const (
    // No errors
    DnxQoseaShowPolicyStatus_no_error DnxQoseaShowPolicyStatus = "no-error"

    // QoS policy is reset
    DnxQoseaShowPolicyStatus_policy_in_reset DnxQoseaShowPolicyStatus = "policy-in-reset"
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

// PlatformQos
// DNX QoS EA operational data
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
    platformQos.EntityData.ParentYangName = "Cisco-IOS-XR-ncs5500-qos-oper"
    platformQos.EntityData.SegmentPath = "Cisco-IOS-XR-ncs5500-qos-oper:platform-qos"
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

    // QoS list of interfaces.
    Interfaces PlatformQos_Nodes_Node_Interfaces

    // QoS list of bundle interfaces.
    BundleInterfaceSingles PlatformQos_Nodes_Node_BundleInterfaceSingles

    // QoS list of remote interfaces.
    RemoteInterfaces PlatformQos_Nodes_Node_RemoteInterfaces
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
    node.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &node.Interfaces}
    node.EntityData.Children["bundle-interface-singles"] = types.YChild{"BundleInterfaceSingles", &node.BundleInterfaceSingles}
    node.EntityData.Children["remote-interfaces"] = types.YChild{"RemoteInterfaces", &node.RemoteInterfaces}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces
// QoS list of bundle interfaces
type PlatformQos_Nodes_Node_BundleInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface names. The type is slice of
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
// QoS interface names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bundle interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // NPU ID. The type is interface{} with range: -2147483648..2147483647.
    NpuId interface{}

    // The interface direction on which QoS is applied to. The type is string.
    QosDirection interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes
}

func (bundleInterface *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface) GetEntityData() *types.CommonEntityData {
    bundleInterface.EntityData.YFilter = bundleInterface.YFilter
    bundleInterface.EntityData.YangName = "bundle-interface"
    bundleInterface.EntityData.BundleName = "cisco_ios_xr"
    bundleInterface.EntityData.ParentYangName = "bundle-interfaces"
    bundleInterface.EntityData.SegmentPath = "bundle-interface"
    bundleInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterface.EntityData.Children = make(map[string]types.YChild)
    bundleInterface.EntityData.Children["policy-details"] = types.YChild{"PolicyDetails", &bundleInterface.PolicyDetails}
    bundleInterface.EntityData.Children["member-interfaces"] = types.YChild{"MemberInterfaces", &bundleInterface.MemberInterfaces}
    bundleInterface.EntityData.Children["classes"] = types.YChild{"Classes", &bundleInterface.Classes}
    bundleInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bundleInterface.InterfaceName}
    bundleInterface.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", bundleInterface.NpuId}
    bundleInterface.EntityData.Leafs["qos-direction"] = types.YLeaf{"QosDirection", bundleInterface.QosDirection}
    return &(bundleInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails struct {
    EntityData types.CommonEntityData
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

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_PolicyDetails) GetEntityData() *types.CommonEntityData {
    policyDetails.EntityData.YFilter = policyDetails.YFilter
    policyDetails.EntityData.YangName = "policy-details"
    policyDetails.EntityData.BundleName = "cisco_ios_xr"
    policyDetails.EntityData.ParentYangName = "bundle-interface"
    policyDetails.EntityData.SegmentPath = "policy-details"
    policyDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyDetails.EntityData.Children = make(map[string]types.YChild)
    policyDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    policyDetails.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", policyDetails.NpuId}
    policyDetails.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", policyDetails.InterfaceHandle}
    policyDetails.EntityData.Leafs["interface-bandwidth-kbps"] = types.YLeaf{"InterfaceBandwidthKbps", policyDetails.InterfaceBandwidthKbps}
    policyDetails.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", policyDetails.PolicyName}
    policyDetails.EntityData.Leafs["total-number-of-classes"] = types.YLeaf{"TotalNumberOfClasses", policyDetails.TotalNumberOfClasses}
    policyDetails.EntityData.Leafs["voq-base-address"] = types.YLeaf{"VoqBaseAddress", policyDetails.VoqBaseAddress}
    policyDetails.EntityData.Leafs["voq-stats-handle"] = types.YLeaf{"VoqStatsHandle", policyDetails.VoqStatsHandle}
    policyDetails.EntityData.Leafs["stats-accounting-type"] = types.YLeaf{"StatsAccountingType", policyDetails.StatsAccountingType}
    policyDetails.EntityData.Leafs["policy-status"] = types.YLeaf{"PolicyStatus", policyDetails.PolicyStatus}
    policyDetails.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", policyDetails.InterfaceStatus}
    return &(policyDetails.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces
// QoS list of member interfaces
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface names. The type is slice of
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
// QoS interface names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes
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
    memberInterface.EntityData.Children["policy-details"] = types.YChild{"PolicyDetails", &memberInterface.PolicyDetails}
    memberInterface.EntityData.Children["classes"] = types.YChild{"Classes", &memberInterface.Classes}
    memberInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    memberInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", memberInterface.InterfaceName}
    return &(memberInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails struct {
    EntityData types.CommonEntityData
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

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_PolicyDetails) GetEntityData() *types.CommonEntityData {
    policyDetails.EntityData.YFilter = policyDetails.YFilter
    policyDetails.EntityData.YangName = "policy-details"
    policyDetails.EntityData.BundleName = "cisco_ios_xr"
    policyDetails.EntityData.ParentYangName = "member-interface"
    policyDetails.EntityData.SegmentPath = "policy-details"
    policyDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyDetails.EntityData.Children = make(map[string]types.YChild)
    policyDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    policyDetails.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", policyDetails.NpuId}
    policyDetails.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", policyDetails.InterfaceHandle}
    policyDetails.EntityData.Leafs["interface-bandwidth-kbps"] = types.YLeaf{"InterfaceBandwidthKbps", policyDetails.InterfaceBandwidthKbps}
    policyDetails.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", policyDetails.PolicyName}
    policyDetails.EntityData.Leafs["total-number-of-classes"] = types.YLeaf{"TotalNumberOfClasses", policyDetails.TotalNumberOfClasses}
    policyDetails.EntityData.Leafs["voq-base-address"] = types.YLeaf{"VoqBaseAddress", policyDetails.VoqBaseAddress}
    policyDetails.EntityData.Leafs["voq-stats-handle"] = types.YLeaf{"VoqStatsHandle", policyDetails.VoqStatsHandle}
    policyDetails.EntityData.Leafs["stats-accounting-type"] = types.YLeaf{"StatsAccountingType", policyDetails.StatsAccountingType}
    policyDetails.EntityData.Leafs["policy-status"] = types.YLeaf{"PolicyStatus", policyDetails.PolicyStatus}
    policyDetails.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", policyDetails.InterfaceStatus}
    return &(policyDetails.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class.
    Class []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "member-interface"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class struct {
    EntityData types.CommonEntityData
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

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["config-max-rate"] = types.YChild{"ConfigMaxRate", &class.ConfigMaxRate}
    class.EntityData.Children["config-min-rate"] = types.YChild{"ConfigMinRate", &class.ConfigMinRate}
    class.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &class.ConfigQueueLimit}
    class.EntityData.Children["config-policer-average-rate"] = types.YChild{"ConfigPolicerAverageRate", &class.ConfigPolicerAverageRate}
    class.EntityData.Children["config-policer-peak-rate"] = types.YChild{"ConfigPolicerPeakRate", &class.ConfigPolicerPeakRate}
    class.EntityData.Children["config-policer-conform-burst"] = types.YChild{"ConfigPolicerConformBurst", &class.ConfigPolicerConformBurst}
    class.EntityData.Children["config-policer-excess-burst"] = types.YChild{"ConfigPolicerExcessBurst", &class.ConfigPolicerExcessBurst}
    class.EntityData.Children["conform-action"] = types.YChild{"ConformAction", &class.ConformAction}
    class.EntityData.Children["exceed-action"] = types.YChild{"ExceedAction", &class.ExceedAction}
    class.EntityData.Children["violate-action"] = types.YChild{"ViolateAction", &class.ViolateAction}
    class.EntityData.Children["ip-mark"] = types.YChild{"IpMark", nil}
    for i := range class.IpMark {
        class.EntityData.Children[types.GetSegmentPath(&class.IpMark[i])] = types.YChild{"IpMark", &class.IpMark[i]}
    }
    class.EntityData.Children["common-mark"] = types.YChild{"CommonMark", nil}
    for i := range class.CommonMark {
        class.EntityData.Children[types.GetSegmentPath(&class.CommonMark[i])] = types.YChild{"CommonMark", &class.CommonMark[i]}
    }
    class.EntityData.Children["mpls-mark"] = types.YChild{"MplsMark", nil}
    for i := range class.MplsMark {
        class.EntityData.Children[types.GetSegmentPath(&class.MplsMark[i])] = types.YChild{"MplsMark", &class.MplsMark[i]}
    }
    class.EntityData.Children["wred"] = types.YChild{"Wred", nil}
    for i := range class.Wred {
        class.EntityData.Children[types.GetSegmentPath(&class.Wred[i])] = types.YChild{"Wred", &class.Wred[i]}
    }
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["level-one-class-name"] = types.YLeaf{"LevelOneClassName", class.LevelOneClassName}
    class.EntityData.Leafs["level-two-class-name"] = types.YLeaf{"LevelTwoClassName", class.LevelTwoClassName}
    class.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", class.ClassLevel}
    class.EntityData.Leafs["egress-queue-id"] = types.YLeaf{"EgressQueueId", class.EgressQueueId}
    class.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", class.QueueType}
    class.EntityData.Leafs["priority-level"] = types.YLeaf{"PriorityLevel", class.PriorityLevel}
    class.EntityData.Leafs["hardware-max-rate-kbps"] = types.YLeaf{"HardwareMaxRateKbps", class.HardwareMaxRateKbps}
    class.EntityData.Leafs["hardware-min-rate-kbps"] = types.YLeaf{"HardwareMinRateKbps", class.HardwareMinRateKbps}
    class.EntityData.Leafs["config-excess-bandwidth-percent"] = types.YLeaf{"ConfigExcessBandwidthPercent", class.ConfigExcessBandwidthPercent}
    class.EntityData.Leafs["config-excess-bandwidth-unit"] = types.YLeaf{"ConfigExcessBandwidthUnit", class.ConfigExcessBandwidthUnit}
    class.EntityData.Leafs["hardware-excess-bandwidth-weight"] = types.YLeaf{"HardwareExcessBandwidthWeight", class.HardwareExcessBandwidthWeight}
    class.EntityData.Leafs["network-min-bandwidth-kbps"] = types.YLeaf{"NetworkMinBandwidthKbps", class.NetworkMinBandwidthKbps}
    class.EntityData.Leafs["hardware-queue-limit-bytes"] = types.YLeaf{"HardwareQueueLimitBytes", class.HardwareQueueLimitBytes}
    class.EntityData.Leafs["hardware-queue-limit-microseconds"] = types.YLeaf{"HardwareQueueLimitMicroseconds", class.HardwareQueueLimitMicroseconds}
    class.EntityData.Leafs["policer-bucket-id"] = types.YLeaf{"PolicerBucketId", class.PolicerBucketId}
    class.EntityData.Leafs["policer-stats-handle"] = types.YLeaf{"PolicerStatsHandle", class.PolicerStatsHandle}
    class.EntityData.Leafs["hardware-policer-average-rate-kbps"] = types.YLeaf{"HardwarePolicerAverageRateKbps", class.HardwarePolicerAverageRateKbps}
    class.EntityData.Leafs["hardware-policer-peak-rate-kbps"] = types.YLeaf{"HardwarePolicerPeakRateKbps", class.HardwarePolicerPeakRateKbps}
    class.EntityData.Leafs["hardware-policer-conform-burst-bytes"] = types.YLeaf{"HardwarePolicerConformBurstBytes", class.HardwarePolicerConformBurstBytes}
    class.EntityData.Leafs["hardware-policer-excess-burst-bytes"] = types.YLeaf{"HardwarePolicerExcessBurstBytes", class.HardwarePolicerExcessBurstBytes}
    return &(class.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetEntityData() *types.CommonEntityData {
    configMaxRate.EntityData.YFilter = configMaxRate.YFilter
    configMaxRate.EntityData.YangName = "config-max-rate"
    configMaxRate.EntityData.BundleName = "cisco_ios_xr"
    configMaxRate.EntityData.ParentYangName = "class"
    configMaxRate.EntityData.SegmentPath = "config-max-rate"
    configMaxRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxRate.EntityData.Children = make(map[string]types.YChild)
    configMaxRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxRate.PolicyValue}
    configMaxRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxRate.PolicyUnit}
    return &(configMaxRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetEntityData() *types.CommonEntityData {
    configMinRate.EntityData.YFilter = configMinRate.YFilter
    configMinRate.EntityData.YangName = "config-min-rate"
    configMinRate.EntityData.BundleName = "cisco_ios_xr"
    configMinRate.EntityData.ParentYangName = "class"
    configMinRate.EntityData.SegmentPath = "config-min-rate"
    configMinRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinRate.EntityData.Children = make(map[string]types.YChild)
    configMinRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinRate.PolicyValue}
    configMinRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinRate.PolicyUnit}
    return &(configMinRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "class"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configQueueLimit.PolicyValue}
    configQueueLimit.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configQueueLimit.PolicyUnit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetEntityData() *types.CommonEntityData {
    configPolicerAverageRate.EntityData.YFilter = configPolicerAverageRate.YFilter
    configPolicerAverageRate.EntityData.YangName = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerAverageRate.EntityData.ParentYangName = "class"
    configPolicerAverageRate.EntityData.SegmentPath = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerAverageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerAverageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerAverageRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerAverageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerAverageRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerAverageRate.PolicyValue}
    configPolicerAverageRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerAverageRate.PolicyUnit}
    return &(configPolicerAverageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetEntityData() *types.CommonEntityData {
    configPolicerPeakRate.EntityData.YFilter = configPolicerPeakRate.YFilter
    configPolicerPeakRate.EntityData.YangName = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerPeakRate.EntityData.ParentYangName = "class"
    configPolicerPeakRate.EntityData.SegmentPath = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerPeakRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerPeakRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerPeakRate.PolicyValue}
    configPolicerPeakRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerPeakRate.PolicyUnit}
    return &(configPolicerPeakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetEntityData() *types.CommonEntityData {
    configPolicerConformBurst.EntityData.YFilter = configPolicerConformBurst.YFilter
    configPolicerConformBurst.EntityData.YangName = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerConformBurst.EntityData.ParentYangName = "class"
    configPolicerConformBurst.EntityData.SegmentPath = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerConformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerConformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerConformBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerConformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerConformBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerConformBurst.PolicyValue}
    configPolicerConformBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerConformBurst.PolicyUnit}
    return &(configPolicerConformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetEntityData() *types.CommonEntityData {
    configPolicerExcessBurst.EntityData.YFilter = configPolicerExcessBurst.YFilter
    configPolicerExcessBurst.EntityData.YangName = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerExcessBurst.EntityData.ParentYangName = "class"
    configPolicerExcessBurst.EntityData.SegmentPath = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerExcessBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerExcessBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerExcessBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerExcessBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerExcessBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerExcessBurst.PolicyValue}
    configPolicerExcessBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerExcessBurst.PolicyUnit}
    return &(configPolicerExcessBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetEntityData() *types.CommonEntityData {
    conformAction.EntityData.YFilter = conformAction.YFilter
    conformAction.EntityData.YangName = "conform-action"
    conformAction.EntityData.BundleName = "cisco_ios_xr"
    conformAction.EntityData.ParentYangName = "class"
    conformAction.EntityData.SegmentPath = "conform-action"
    conformAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformAction.EntityData.Children = make(map[string]types.YChild)
    conformAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range conformAction.Mark {
        conformAction.EntityData.Children[types.GetSegmentPath(&conformAction.Mark[i])] = types.YChild{"Mark", &conformAction.Mark[i]}
    }
    conformAction.EntityData.Leafs = make(map[string]types.YLeaf)
    conformAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", conformAction.ActionType}
    return &(conformAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "conform-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetEntityData() *types.CommonEntityData {
    exceedAction.EntityData.YFilter = exceedAction.YFilter
    exceedAction.EntityData.YangName = "exceed-action"
    exceedAction.EntityData.BundleName = "cisco_ios_xr"
    exceedAction.EntityData.ParentYangName = "class"
    exceedAction.EntityData.SegmentPath = "exceed-action"
    exceedAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedAction.EntityData.Children = make(map[string]types.YChild)
    exceedAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range exceedAction.Mark {
        exceedAction.EntityData.Children[types.GetSegmentPath(&exceedAction.Mark[i])] = types.YChild{"Mark", &exceedAction.Mark[i]}
    }
    exceedAction.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", exceedAction.ActionType}
    return &(exceedAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "exceed-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetEntityData() *types.CommonEntityData {
    violateAction.EntityData.YFilter = violateAction.YFilter
    violateAction.EntityData.YangName = "violate-action"
    violateAction.EntityData.BundleName = "cisco_ios_xr"
    violateAction.EntityData.ParentYangName = "class"
    violateAction.EntityData.SegmentPath = "violate-action"
    violateAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    violateAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    violateAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    violateAction.EntityData.Children = make(map[string]types.YChild)
    violateAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range violateAction.Mark {
        violateAction.EntityData.Children[types.GetSegmentPath(&violateAction.Mark[i])] = types.YChild{"Mark", &violateAction.Mark[i]}
    }
    violateAction.EntityData.Leafs = make(map[string]types.YLeaf)
    violateAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", violateAction.ActionType}
    return &(violateAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "violate-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetEntityData() *types.CommonEntityData {
    ipMark.EntityData.YFilter = ipMark.YFilter
    ipMark.EntityData.YangName = "ip-mark"
    ipMark.EntityData.BundleName = "cisco_ios_xr"
    ipMark.EntityData.ParentYangName = "class"
    ipMark.EntityData.SegmentPath = "ip-mark"
    ipMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipMark.EntityData.Children = make(map[string]types.YChild)
    ipMark.EntityData.Leafs = make(map[string]types.YLeaf)
    ipMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", ipMark.MarkType}
    ipMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", ipMark.MarkValue}
    return &(ipMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetEntityData() *types.CommonEntityData {
    commonMark.EntityData.YFilter = commonMark.YFilter
    commonMark.EntityData.YangName = "common-mark"
    commonMark.EntityData.BundleName = "cisco_ios_xr"
    commonMark.EntityData.ParentYangName = "class"
    commonMark.EntityData.SegmentPath = "common-mark"
    commonMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonMark.EntityData.Children = make(map[string]types.YChild)
    commonMark.EntityData.Leafs = make(map[string]types.YLeaf)
    commonMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", commonMark.MarkType}
    commonMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", commonMark.MarkValue}
    return &(commonMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetEntityData() *types.CommonEntityData {
    mplsMark.EntityData.YFilter = mplsMark.YFilter
    mplsMark.EntityData.YangName = "mpls-mark"
    mplsMark.EntityData.BundleName = "cisco_ios_xr"
    mplsMark.EntityData.ParentYangName = "class"
    mplsMark.EntityData.SegmentPath = "mpls-mark"
    mplsMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsMark.EntityData.Children = make(map[string]types.YChild)
    mplsMark.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mplsMark.MarkType}
    mplsMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mplsMark.MarkValue}
    return &(mplsMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred struct {
    EntityData types.CommonEntityData
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

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "class"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["wred-match-value"] = types.YChild{"WredMatchValue", &wred.WredMatchValue}
    wred.EntityData.Children["config-min-threshold"] = types.YChild{"ConfigMinThreshold", &wred.ConfigMinThreshold}
    wred.EntityData.Children["config-max-threshold"] = types.YChild{"ConfigMaxThreshold", &wred.ConfigMaxThreshold}
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["wred-match-type"] = types.YLeaf{"WredMatchType", wred.WredMatchType}
    wred.EntityData.Leafs["hardware-min-threshold-bytes"] = types.YLeaf{"HardwareMinThresholdBytes", wred.HardwareMinThresholdBytes}
    wred.EntityData.Leafs["hardware-max-threshold-bytes"] = types.YLeaf{"HardwareMaxThresholdBytes", wred.HardwareMaxThresholdBytes}
    wred.EntityData.Leafs["first-segment"] = types.YLeaf{"FirstSegment", wred.FirstSegment}
    wred.EntityData.Leafs["segment-size"] = types.YLeaf{"SegmentSize", wred.SegmentSize}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetEntityData() *types.CommonEntityData {
    wredMatchValue.EntityData.YFilter = wredMatchValue.YFilter
    wredMatchValue.EntityData.YangName = "wred-match-value"
    wredMatchValue.EntityData.BundleName = "cisco_ios_xr"
    wredMatchValue.EntityData.ParentYangName = "wred"
    wredMatchValue.EntityData.SegmentPath = "wred-match-value"
    wredMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wredMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wredMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wredMatchValue.EntityData.Children = make(map[string]types.YChild)
    wredMatchValue.EntityData.Children["dnx-qosea-show-red-match-value"] = types.YChild{"DnxQoseaShowRedMatchValue", nil}
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        wredMatchValue.EntityData.Children[types.GetSegmentPath(&wredMatchValue.DnxQoseaShowRedMatchValue[i])] = types.YChild{"DnxQoseaShowRedMatchValue", &wredMatchValue.DnxQoseaShowRedMatchValue[i]}
    }
    wredMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wredMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetEntityData() *types.CommonEntityData {
    dnxQoseaShowRedMatchValue.EntityData.YFilter = dnxQoseaShowRedMatchValue.YFilter
    dnxQoseaShowRedMatchValue.EntityData.YangName = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.BundleName = "cisco_ios_xr"
    dnxQoseaShowRedMatchValue.EntityData.ParentYangName = "wred-match-value"
    dnxQoseaShowRedMatchValue.EntityData.SegmentPath = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnxQoseaShowRedMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnxQoseaShowRedMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnxQoseaShowRedMatchValue.EntityData.Children = make(map[string]types.YChild)
    dnxQoseaShowRedMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-start"] = types.YLeaf{"RangeStart", dnxQoseaShowRedMatchValue.RangeStart}
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-end"] = types.YLeaf{"RangeEnd", dnxQoseaShowRedMatchValue.RangeEnd}
    return &(dnxQoseaShowRedMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetEntityData() *types.CommonEntityData {
    configMinThreshold.EntityData.YFilter = configMinThreshold.YFilter
    configMinThreshold.EntityData.YangName = "config-min-threshold"
    configMinThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMinThreshold.EntityData.ParentYangName = "wred"
    configMinThreshold.EntityData.SegmentPath = "config-min-threshold"
    configMinThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinThreshold.EntityData.Children = make(map[string]types.YChild)
    configMinThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinThreshold.PolicyValue}
    configMinThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinThreshold.PolicyUnit}
    return &(configMinThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetEntityData() *types.CommonEntityData {
    configMaxThreshold.EntityData.YFilter = configMaxThreshold.YFilter
    configMaxThreshold.EntityData.YangName = "config-max-threshold"
    configMaxThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMaxThreshold.EntityData.ParentYangName = "wred"
    configMaxThreshold.EntityData.SegmentPath = "config-max-threshold"
    configMaxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxThreshold.EntityData.Children = make(map[string]types.YChild)
    configMaxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxThreshold.PolicyValue}
    configMaxThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxThreshold.PolicyUnit}
    return &(configMaxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class.
    Class []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "bundle-interface"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class struct {
    EntityData types.CommonEntityData
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

func (class *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["config-max-rate"] = types.YChild{"ConfigMaxRate", &class.ConfigMaxRate}
    class.EntityData.Children["config-min-rate"] = types.YChild{"ConfigMinRate", &class.ConfigMinRate}
    class.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &class.ConfigQueueLimit}
    class.EntityData.Children["config-policer-average-rate"] = types.YChild{"ConfigPolicerAverageRate", &class.ConfigPolicerAverageRate}
    class.EntityData.Children["config-policer-peak-rate"] = types.YChild{"ConfigPolicerPeakRate", &class.ConfigPolicerPeakRate}
    class.EntityData.Children["config-policer-conform-burst"] = types.YChild{"ConfigPolicerConformBurst", &class.ConfigPolicerConformBurst}
    class.EntityData.Children["config-policer-excess-burst"] = types.YChild{"ConfigPolicerExcessBurst", &class.ConfigPolicerExcessBurst}
    class.EntityData.Children["conform-action"] = types.YChild{"ConformAction", &class.ConformAction}
    class.EntityData.Children["exceed-action"] = types.YChild{"ExceedAction", &class.ExceedAction}
    class.EntityData.Children["violate-action"] = types.YChild{"ViolateAction", &class.ViolateAction}
    class.EntityData.Children["ip-mark"] = types.YChild{"IpMark", nil}
    for i := range class.IpMark {
        class.EntityData.Children[types.GetSegmentPath(&class.IpMark[i])] = types.YChild{"IpMark", &class.IpMark[i]}
    }
    class.EntityData.Children["common-mark"] = types.YChild{"CommonMark", nil}
    for i := range class.CommonMark {
        class.EntityData.Children[types.GetSegmentPath(&class.CommonMark[i])] = types.YChild{"CommonMark", &class.CommonMark[i]}
    }
    class.EntityData.Children["mpls-mark"] = types.YChild{"MplsMark", nil}
    for i := range class.MplsMark {
        class.EntityData.Children[types.GetSegmentPath(&class.MplsMark[i])] = types.YChild{"MplsMark", &class.MplsMark[i]}
    }
    class.EntityData.Children["wred"] = types.YChild{"Wred", nil}
    for i := range class.Wred {
        class.EntityData.Children[types.GetSegmentPath(&class.Wred[i])] = types.YChild{"Wred", &class.Wred[i]}
    }
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["level-one-class-name"] = types.YLeaf{"LevelOneClassName", class.LevelOneClassName}
    class.EntityData.Leafs["level-two-class-name"] = types.YLeaf{"LevelTwoClassName", class.LevelTwoClassName}
    class.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", class.ClassLevel}
    class.EntityData.Leafs["egress-queue-id"] = types.YLeaf{"EgressQueueId", class.EgressQueueId}
    class.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", class.QueueType}
    class.EntityData.Leafs["priority-level"] = types.YLeaf{"PriorityLevel", class.PriorityLevel}
    class.EntityData.Leafs["hardware-max-rate-kbps"] = types.YLeaf{"HardwareMaxRateKbps", class.HardwareMaxRateKbps}
    class.EntityData.Leafs["hardware-min-rate-kbps"] = types.YLeaf{"HardwareMinRateKbps", class.HardwareMinRateKbps}
    class.EntityData.Leafs["config-excess-bandwidth-percent"] = types.YLeaf{"ConfigExcessBandwidthPercent", class.ConfigExcessBandwidthPercent}
    class.EntityData.Leafs["config-excess-bandwidth-unit"] = types.YLeaf{"ConfigExcessBandwidthUnit", class.ConfigExcessBandwidthUnit}
    class.EntityData.Leafs["hardware-excess-bandwidth-weight"] = types.YLeaf{"HardwareExcessBandwidthWeight", class.HardwareExcessBandwidthWeight}
    class.EntityData.Leafs["network-min-bandwidth-kbps"] = types.YLeaf{"NetworkMinBandwidthKbps", class.NetworkMinBandwidthKbps}
    class.EntityData.Leafs["hardware-queue-limit-bytes"] = types.YLeaf{"HardwareQueueLimitBytes", class.HardwareQueueLimitBytes}
    class.EntityData.Leafs["hardware-queue-limit-microseconds"] = types.YLeaf{"HardwareQueueLimitMicroseconds", class.HardwareQueueLimitMicroseconds}
    class.EntityData.Leafs["policer-bucket-id"] = types.YLeaf{"PolicerBucketId", class.PolicerBucketId}
    class.EntityData.Leafs["policer-stats-handle"] = types.YLeaf{"PolicerStatsHandle", class.PolicerStatsHandle}
    class.EntityData.Leafs["hardware-policer-average-rate-kbps"] = types.YLeaf{"HardwarePolicerAverageRateKbps", class.HardwarePolicerAverageRateKbps}
    class.EntityData.Leafs["hardware-policer-peak-rate-kbps"] = types.YLeaf{"HardwarePolicerPeakRateKbps", class.HardwarePolicerPeakRateKbps}
    class.EntityData.Leafs["hardware-policer-conform-burst-bytes"] = types.YLeaf{"HardwarePolicerConformBurstBytes", class.HardwarePolicerConformBurstBytes}
    class.EntityData.Leafs["hardware-policer-excess-burst-bytes"] = types.YLeaf{"HardwarePolicerExcessBurstBytes", class.HardwarePolicerExcessBurstBytes}
    return &(class.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMaxRate) GetEntityData() *types.CommonEntityData {
    configMaxRate.EntityData.YFilter = configMaxRate.YFilter
    configMaxRate.EntityData.YangName = "config-max-rate"
    configMaxRate.EntityData.BundleName = "cisco_ios_xr"
    configMaxRate.EntityData.ParentYangName = "class"
    configMaxRate.EntityData.SegmentPath = "config-max-rate"
    configMaxRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxRate.EntityData.Children = make(map[string]types.YChild)
    configMaxRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxRate.PolicyValue}
    configMaxRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxRate.PolicyUnit}
    return &(configMaxRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigMinRate) GetEntityData() *types.CommonEntityData {
    configMinRate.EntityData.YFilter = configMinRate.YFilter
    configMinRate.EntityData.YangName = "config-min-rate"
    configMinRate.EntityData.BundleName = "cisco_ios_xr"
    configMinRate.EntityData.ParentYangName = "class"
    configMinRate.EntityData.SegmentPath = "config-min-rate"
    configMinRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinRate.EntityData.Children = make(map[string]types.YChild)
    configMinRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinRate.PolicyValue}
    configMinRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinRate.PolicyUnit}
    return &(configMinRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "class"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configQueueLimit.PolicyValue}
    configQueueLimit.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configQueueLimit.PolicyUnit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerAverageRate) GetEntityData() *types.CommonEntityData {
    configPolicerAverageRate.EntityData.YFilter = configPolicerAverageRate.YFilter
    configPolicerAverageRate.EntityData.YangName = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerAverageRate.EntityData.ParentYangName = "class"
    configPolicerAverageRate.EntityData.SegmentPath = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerAverageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerAverageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerAverageRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerAverageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerAverageRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerAverageRate.PolicyValue}
    configPolicerAverageRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerAverageRate.PolicyUnit}
    return &(configPolicerAverageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerPeakRate) GetEntityData() *types.CommonEntityData {
    configPolicerPeakRate.EntityData.YFilter = configPolicerPeakRate.YFilter
    configPolicerPeakRate.EntityData.YangName = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerPeakRate.EntityData.ParentYangName = "class"
    configPolicerPeakRate.EntityData.SegmentPath = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerPeakRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerPeakRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerPeakRate.PolicyValue}
    configPolicerPeakRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerPeakRate.PolicyUnit}
    return &(configPolicerPeakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerConformBurst) GetEntityData() *types.CommonEntityData {
    configPolicerConformBurst.EntityData.YFilter = configPolicerConformBurst.YFilter
    configPolicerConformBurst.EntityData.YangName = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerConformBurst.EntityData.ParentYangName = "class"
    configPolicerConformBurst.EntityData.SegmentPath = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerConformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerConformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerConformBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerConformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerConformBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerConformBurst.PolicyValue}
    configPolicerConformBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerConformBurst.PolicyUnit}
    return &(configPolicerConformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConfigPolicerExcessBurst) GetEntityData() *types.CommonEntityData {
    configPolicerExcessBurst.EntityData.YFilter = configPolicerExcessBurst.YFilter
    configPolicerExcessBurst.EntityData.YangName = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerExcessBurst.EntityData.ParentYangName = "class"
    configPolicerExcessBurst.EntityData.SegmentPath = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerExcessBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerExcessBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerExcessBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerExcessBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerExcessBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerExcessBurst.PolicyValue}
    configPolicerExcessBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerExcessBurst.PolicyUnit}
    return &(configPolicerExcessBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction) GetEntityData() *types.CommonEntityData {
    conformAction.EntityData.YFilter = conformAction.YFilter
    conformAction.EntityData.YangName = "conform-action"
    conformAction.EntityData.BundleName = "cisco_ios_xr"
    conformAction.EntityData.ParentYangName = "class"
    conformAction.EntityData.SegmentPath = "conform-action"
    conformAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformAction.EntityData.Children = make(map[string]types.YChild)
    conformAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range conformAction.Mark {
        conformAction.EntityData.Children[types.GetSegmentPath(&conformAction.Mark[i])] = types.YChild{"Mark", &conformAction.Mark[i]}
    }
    conformAction.EntityData.Leafs = make(map[string]types.YLeaf)
    conformAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", conformAction.ActionType}
    return &(conformAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ConformAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "conform-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction) GetEntityData() *types.CommonEntityData {
    exceedAction.EntityData.YFilter = exceedAction.YFilter
    exceedAction.EntityData.YangName = "exceed-action"
    exceedAction.EntityData.BundleName = "cisco_ios_xr"
    exceedAction.EntityData.ParentYangName = "class"
    exceedAction.EntityData.SegmentPath = "exceed-action"
    exceedAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedAction.EntityData.Children = make(map[string]types.YChild)
    exceedAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range exceedAction.Mark {
        exceedAction.EntityData.Children[types.GetSegmentPath(&exceedAction.Mark[i])] = types.YChild{"Mark", &exceedAction.Mark[i]}
    }
    exceedAction.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", exceedAction.ActionType}
    return &(exceedAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ExceedAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "exceed-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction) GetEntityData() *types.CommonEntityData {
    violateAction.EntityData.YFilter = violateAction.YFilter
    violateAction.EntityData.YangName = "violate-action"
    violateAction.EntityData.BundleName = "cisco_ios_xr"
    violateAction.EntityData.ParentYangName = "class"
    violateAction.EntityData.SegmentPath = "violate-action"
    violateAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    violateAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    violateAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    violateAction.EntityData.Children = make(map[string]types.YChild)
    violateAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range violateAction.Mark {
        violateAction.EntityData.Children[types.GetSegmentPath(&violateAction.Mark[i])] = types.YChild{"Mark", &violateAction.Mark[i]}
    }
    violateAction.EntityData.Leafs = make(map[string]types.YLeaf)
    violateAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", violateAction.ActionType}
    return &(violateAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_ViolateAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "violate-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_IpMark) GetEntityData() *types.CommonEntityData {
    ipMark.EntityData.YFilter = ipMark.YFilter
    ipMark.EntityData.YangName = "ip-mark"
    ipMark.EntityData.BundleName = "cisco_ios_xr"
    ipMark.EntityData.ParentYangName = "class"
    ipMark.EntityData.SegmentPath = "ip-mark"
    ipMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipMark.EntityData.Children = make(map[string]types.YChild)
    ipMark.EntityData.Leafs = make(map[string]types.YLeaf)
    ipMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", ipMark.MarkType}
    ipMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", ipMark.MarkValue}
    return &(ipMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_CommonMark) GetEntityData() *types.CommonEntityData {
    commonMark.EntityData.YFilter = commonMark.YFilter
    commonMark.EntityData.YangName = "common-mark"
    commonMark.EntityData.BundleName = "cisco_ios_xr"
    commonMark.EntityData.ParentYangName = "class"
    commonMark.EntityData.SegmentPath = "common-mark"
    commonMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonMark.EntityData.Children = make(map[string]types.YChild)
    commonMark.EntityData.Leafs = make(map[string]types.YLeaf)
    commonMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", commonMark.MarkType}
    commonMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", commonMark.MarkValue}
    return &(commonMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_MplsMark) GetEntityData() *types.CommonEntityData {
    mplsMark.EntityData.YFilter = mplsMark.YFilter
    mplsMark.EntityData.YangName = "mpls-mark"
    mplsMark.EntityData.BundleName = "cisco_ios_xr"
    mplsMark.EntityData.ParentYangName = "class"
    mplsMark.EntityData.SegmentPath = "mpls-mark"
    mplsMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsMark.EntityData.Children = make(map[string]types.YChild)
    mplsMark.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mplsMark.MarkType}
    mplsMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mplsMark.MarkValue}
    return &(mplsMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred struct {
    EntityData types.CommonEntityData
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

func (wred *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "class"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["wred-match-value"] = types.YChild{"WredMatchValue", &wred.WredMatchValue}
    wred.EntityData.Children["config-min-threshold"] = types.YChild{"ConfigMinThreshold", &wred.ConfigMinThreshold}
    wred.EntityData.Children["config-max-threshold"] = types.YChild{"ConfigMaxThreshold", &wred.ConfigMaxThreshold}
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["wred-match-type"] = types.YLeaf{"WredMatchType", wred.WredMatchType}
    wred.EntityData.Leafs["hardware-min-threshold-bytes"] = types.YLeaf{"HardwareMinThresholdBytes", wred.HardwareMinThresholdBytes}
    wred.EntityData.Leafs["hardware-max-threshold-bytes"] = types.YLeaf{"HardwareMaxThresholdBytes", wred.HardwareMaxThresholdBytes}
    wred.EntityData.Leafs["first-segment"] = types.YLeaf{"FirstSegment", wred.FirstSegment}
    wred.EntityData.Leafs["segment-size"] = types.YLeaf{"SegmentSize", wred.SegmentSize}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue) GetEntityData() *types.CommonEntityData {
    wredMatchValue.EntityData.YFilter = wredMatchValue.YFilter
    wredMatchValue.EntityData.YangName = "wred-match-value"
    wredMatchValue.EntityData.BundleName = "cisco_ios_xr"
    wredMatchValue.EntityData.ParentYangName = "wred"
    wredMatchValue.EntityData.SegmentPath = "wred-match-value"
    wredMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wredMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wredMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wredMatchValue.EntityData.Children = make(map[string]types.YChild)
    wredMatchValue.EntityData.Children["dnx-qosea-show-red-match-value"] = types.YChild{"DnxQoseaShowRedMatchValue", nil}
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        wredMatchValue.EntityData.Children[types.GetSegmentPath(&wredMatchValue.DnxQoseaShowRedMatchValue[i])] = types.YChild{"DnxQoseaShowRedMatchValue", &wredMatchValue.DnxQoseaShowRedMatchValue[i]}
    }
    wredMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wredMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetEntityData() *types.CommonEntityData {
    dnxQoseaShowRedMatchValue.EntityData.YFilter = dnxQoseaShowRedMatchValue.YFilter
    dnxQoseaShowRedMatchValue.EntityData.YangName = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.BundleName = "cisco_ios_xr"
    dnxQoseaShowRedMatchValue.EntityData.ParentYangName = "wred-match-value"
    dnxQoseaShowRedMatchValue.EntityData.SegmentPath = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnxQoseaShowRedMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnxQoseaShowRedMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnxQoseaShowRedMatchValue.EntityData.Children = make(map[string]types.YChild)
    dnxQoseaShowRedMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-start"] = types.YLeaf{"RangeStart", dnxQoseaShowRedMatchValue.RangeStart}
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-end"] = types.YLeaf{"RangeEnd", dnxQoseaShowRedMatchValue.RangeEnd}
    return &(dnxQoseaShowRedMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMinThreshold) GetEntityData() *types.CommonEntityData {
    configMinThreshold.EntityData.YFilter = configMinThreshold.YFilter
    configMinThreshold.EntityData.YangName = "config-min-threshold"
    configMinThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMinThreshold.EntityData.ParentYangName = "wred"
    configMinThreshold.EntityData.SegmentPath = "config-min-threshold"
    configMinThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinThreshold.EntityData.Children = make(map[string]types.YChild)
    configMinThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinThreshold.PolicyValue}
    configMinThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinThreshold.PolicyUnit}
    return &(configMinThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaces_BundleInterface_Classes_Class_Wred_ConfigMaxThreshold) GetEntityData() *types.CommonEntityData {
    configMaxThreshold.EntityData.YFilter = configMaxThreshold.YFilter
    configMaxThreshold.EntityData.YangName = "config-max-threshold"
    configMaxThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMaxThreshold.EntityData.ParentYangName = "wred"
    configMaxThreshold.EntityData.SegmentPath = "config-max-threshold"
    configMaxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxThreshold.EntityData.Children = make(map[string]types.YChild)
    configMaxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxThreshold.PolicyValue}
    configMaxThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxThreshold.PolicyUnit}
    return &(configMaxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces
// QoS list of interfaces
type PlatformQos_Nodes_Node_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface names. The type is slice of
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
// QoS interface names
type PlatformQos_Nodes_Node_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // The interface direction on which QoS is applied to. The type is string.
    QosDirection interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_Interfaces_Interface_Classes
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
    self.EntityData.Children["policy-details"] = types.YChild{"PolicyDetails", &self.PolicyDetails}
    self.EntityData.Children["classes"] = types.YChild{"Classes", &self.Classes}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["qos-direction"] = types.YLeaf{"QosDirection", self.QosDirection}
    return &(self.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails struct {
    EntityData types.CommonEntityData
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

func (policyDetails *PlatformQos_Nodes_Node_Interfaces_Interface_PolicyDetails) GetEntityData() *types.CommonEntityData {
    policyDetails.EntityData.YFilter = policyDetails.YFilter
    policyDetails.EntityData.YangName = "policy-details"
    policyDetails.EntityData.BundleName = "cisco_ios_xr"
    policyDetails.EntityData.ParentYangName = "interface"
    policyDetails.EntityData.SegmentPath = "policy-details"
    policyDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyDetails.EntityData.Children = make(map[string]types.YChild)
    policyDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    policyDetails.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", policyDetails.NpuId}
    policyDetails.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", policyDetails.InterfaceHandle}
    policyDetails.EntityData.Leafs["interface-bandwidth-kbps"] = types.YLeaf{"InterfaceBandwidthKbps", policyDetails.InterfaceBandwidthKbps}
    policyDetails.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", policyDetails.PolicyName}
    policyDetails.EntityData.Leafs["total-number-of-classes"] = types.YLeaf{"TotalNumberOfClasses", policyDetails.TotalNumberOfClasses}
    policyDetails.EntityData.Leafs["voq-base-address"] = types.YLeaf{"VoqBaseAddress", policyDetails.VoqBaseAddress}
    policyDetails.EntityData.Leafs["voq-stats-handle"] = types.YLeaf{"VoqStatsHandle", policyDetails.VoqStatsHandle}
    policyDetails.EntityData.Leafs["stats-accounting-type"] = types.YLeaf{"StatsAccountingType", policyDetails.StatsAccountingType}
    policyDetails.EntityData.Leafs["policy-status"] = types.YLeaf{"PolicyStatus", policyDetails.PolicyStatus}
    policyDetails.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", policyDetails.InterfaceStatus}
    return &(policyDetails.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class.
    Class []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_Interfaces_Interface_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "interface"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class struct {
    EntityData types.CommonEntityData
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

func (class *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["config-max-rate"] = types.YChild{"ConfigMaxRate", &class.ConfigMaxRate}
    class.EntityData.Children["config-min-rate"] = types.YChild{"ConfigMinRate", &class.ConfigMinRate}
    class.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &class.ConfigQueueLimit}
    class.EntityData.Children["config-policer-average-rate"] = types.YChild{"ConfigPolicerAverageRate", &class.ConfigPolicerAverageRate}
    class.EntityData.Children["config-policer-peak-rate"] = types.YChild{"ConfigPolicerPeakRate", &class.ConfigPolicerPeakRate}
    class.EntityData.Children["config-policer-conform-burst"] = types.YChild{"ConfigPolicerConformBurst", &class.ConfigPolicerConformBurst}
    class.EntityData.Children["config-policer-excess-burst"] = types.YChild{"ConfigPolicerExcessBurst", &class.ConfigPolicerExcessBurst}
    class.EntityData.Children["conform-action"] = types.YChild{"ConformAction", &class.ConformAction}
    class.EntityData.Children["exceed-action"] = types.YChild{"ExceedAction", &class.ExceedAction}
    class.EntityData.Children["violate-action"] = types.YChild{"ViolateAction", &class.ViolateAction}
    class.EntityData.Children["ip-mark"] = types.YChild{"IpMark", nil}
    for i := range class.IpMark {
        class.EntityData.Children[types.GetSegmentPath(&class.IpMark[i])] = types.YChild{"IpMark", &class.IpMark[i]}
    }
    class.EntityData.Children["common-mark"] = types.YChild{"CommonMark", nil}
    for i := range class.CommonMark {
        class.EntityData.Children[types.GetSegmentPath(&class.CommonMark[i])] = types.YChild{"CommonMark", &class.CommonMark[i]}
    }
    class.EntityData.Children["mpls-mark"] = types.YChild{"MplsMark", nil}
    for i := range class.MplsMark {
        class.EntityData.Children[types.GetSegmentPath(&class.MplsMark[i])] = types.YChild{"MplsMark", &class.MplsMark[i]}
    }
    class.EntityData.Children["wred"] = types.YChild{"Wred", nil}
    for i := range class.Wred {
        class.EntityData.Children[types.GetSegmentPath(&class.Wred[i])] = types.YChild{"Wred", &class.Wred[i]}
    }
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["level-one-class-name"] = types.YLeaf{"LevelOneClassName", class.LevelOneClassName}
    class.EntityData.Leafs["level-two-class-name"] = types.YLeaf{"LevelTwoClassName", class.LevelTwoClassName}
    class.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", class.ClassLevel}
    class.EntityData.Leafs["egress-queue-id"] = types.YLeaf{"EgressQueueId", class.EgressQueueId}
    class.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", class.QueueType}
    class.EntityData.Leafs["priority-level"] = types.YLeaf{"PriorityLevel", class.PriorityLevel}
    class.EntityData.Leafs["hardware-max-rate-kbps"] = types.YLeaf{"HardwareMaxRateKbps", class.HardwareMaxRateKbps}
    class.EntityData.Leafs["hardware-min-rate-kbps"] = types.YLeaf{"HardwareMinRateKbps", class.HardwareMinRateKbps}
    class.EntityData.Leafs["config-excess-bandwidth-percent"] = types.YLeaf{"ConfigExcessBandwidthPercent", class.ConfigExcessBandwidthPercent}
    class.EntityData.Leafs["config-excess-bandwidth-unit"] = types.YLeaf{"ConfigExcessBandwidthUnit", class.ConfigExcessBandwidthUnit}
    class.EntityData.Leafs["hardware-excess-bandwidth-weight"] = types.YLeaf{"HardwareExcessBandwidthWeight", class.HardwareExcessBandwidthWeight}
    class.EntityData.Leafs["network-min-bandwidth-kbps"] = types.YLeaf{"NetworkMinBandwidthKbps", class.NetworkMinBandwidthKbps}
    class.EntityData.Leafs["hardware-queue-limit-bytes"] = types.YLeaf{"HardwareQueueLimitBytes", class.HardwareQueueLimitBytes}
    class.EntityData.Leafs["hardware-queue-limit-microseconds"] = types.YLeaf{"HardwareQueueLimitMicroseconds", class.HardwareQueueLimitMicroseconds}
    class.EntityData.Leafs["policer-bucket-id"] = types.YLeaf{"PolicerBucketId", class.PolicerBucketId}
    class.EntityData.Leafs["policer-stats-handle"] = types.YLeaf{"PolicerStatsHandle", class.PolicerStatsHandle}
    class.EntityData.Leafs["hardware-policer-average-rate-kbps"] = types.YLeaf{"HardwarePolicerAverageRateKbps", class.HardwarePolicerAverageRateKbps}
    class.EntityData.Leafs["hardware-policer-peak-rate-kbps"] = types.YLeaf{"HardwarePolicerPeakRateKbps", class.HardwarePolicerPeakRateKbps}
    class.EntityData.Leafs["hardware-policer-conform-burst-bytes"] = types.YLeaf{"HardwarePolicerConformBurstBytes", class.HardwarePolicerConformBurstBytes}
    class.EntityData.Leafs["hardware-policer-excess-burst-bytes"] = types.YLeaf{"HardwarePolicerExcessBurstBytes", class.HardwarePolicerExcessBurstBytes}
    return &(class.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMaxRate) GetEntityData() *types.CommonEntityData {
    configMaxRate.EntityData.YFilter = configMaxRate.YFilter
    configMaxRate.EntityData.YangName = "config-max-rate"
    configMaxRate.EntityData.BundleName = "cisco_ios_xr"
    configMaxRate.EntityData.ParentYangName = "class"
    configMaxRate.EntityData.SegmentPath = "config-max-rate"
    configMaxRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxRate.EntityData.Children = make(map[string]types.YChild)
    configMaxRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxRate.PolicyValue}
    configMaxRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxRate.PolicyUnit}
    return &(configMaxRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigMinRate) GetEntityData() *types.CommonEntityData {
    configMinRate.EntityData.YFilter = configMinRate.YFilter
    configMinRate.EntityData.YangName = "config-min-rate"
    configMinRate.EntityData.BundleName = "cisco_ios_xr"
    configMinRate.EntityData.ParentYangName = "class"
    configMinRate.EntityData.SegmentPath = "config-min-rate"
    configMinRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinRate.EntityData.Children = make(map[string]types.YChild)
    configMinRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinRate.PolicyValue}
    configMinRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinRate.PolicyUnit}
    return &(configMinRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "class"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configQueueLimit.PolicyValue}
    configQueueLimit.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configQueueLimit.PolicyUnit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerAverageRate) GetEntityData() *types.CommonEntityData {
    configPolicerAverageRate.EntityData.YFilter = configPolicerAverageRate.YFilter
    configPolicerAverageRate.EntityData.YangName = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerAverageRate.EntityData.ParentYangName = "class"
    configPolicerAverageRate.EntityData.SegmentPath = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerAverageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerAverageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerAverageRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerAverageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerAverageRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerAverageRate.PolicyValue}
    configPolicerAverageRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerAverageRate.PolicyUnit}
    return &(configPolicerAverageRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerPeakRate) GetEntityData() *types.CommonEntityData {
    configPolicerPeakRate.EntityData.YFilter = configPolicerPeakRate.YFilter
    configPolicerPeakRate.EntityData.YangName = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerPeakRate.EntityData.ParentYangName = "class"
    configPolicerPeakRate.EntityData.SegmentPath = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerPeakRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerPeakRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerPeakRate.PolicyValue}
    configPolicerPeakRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerPeakRate.PolicyUnit}
    return &(configPolicerPeakRate.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerConformBurst) GetEntityData() *types.CommonEntityData {
    configPolicerConformBurst.EntityData.YFilter = configPolicerConformBurst.YFilter
    configPolicerConformBurst.EntityData.YangName = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerConformBurst.EntityData.ParentYangName = "class"
    configPolicerConformBurst.EntityData.SegmentPath = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerConformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerConformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerConformBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerConformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerConformBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerConformBurst.PolicyValue}
    configPolicerConformBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerConformBurst.PolicyUnit}
    return &(configPolicerConformBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConfigPolicerExcessBurst) GetEntityData() *types.CommonEntityData {
    configPolicerExcessBurst.EntityData.YFilter = configPolicerExcessBurst.YFilter
    configPolicerExcessBurst.EntityData.YangName = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerExcessBurst.EntityData.ParentYangName = "class"
    configPolicerExcessBurst.EntityData.SegmentPath = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerExcessBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerExcessBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerExcessBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerExcessBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerExcessBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerExcessBurst.PolicyValue}
    configPolicerExcessBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerExcessBurst.PolicyUnit}
    return &(configPolicerExcessBurst.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction) GetEntityData() *types.CommonEntityData {
    conformAction.EntityData.YFilter = conformAction.YFilter
    conformAction.EntityData.YangName = "conform-action"
    conformAction.EntityData.BundleName = "cisco_ios_xr"
    conformAction.EntityData.ParentYangName = "class"
    conformAction.EntityData.SegmentPath = "conform-action"
    conformAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformAction.EntityData.Children = make(map[string]types.YChild)
    conformAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range conformAction.Mark {
        conformAction.EntityData.Children[types.GetSegmentPath(&conformAction.Mark[i])] = types.YChild{"Mark", &conformAction.Mark[i]}
    }
    conformAction.EntityData.Leafs = make(map[string]types.YLeaf)
    conformAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", conformAction.ActionType}
    return &(conformAction.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ConformAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "conform-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction) GetEntityData() *types.CommonEntityData {
    exceedAction.EntityData.YFilter = exceedAction.YFilter
    exceedAction.EntityData.YangName = "exceed-action"
    exceedAction.EntityData.BundleName = "cisco_ios_xr"
    exceedAction.EntityData.ParentYangName = "class"
    exceedAction.EntityData.SegmentPath = "exceed-action"
    exceedAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedAction.EntityData.Children = make(map[string]types.YChild)
    exceedAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range exceedAction.Mark {
        exceedAction.EntityData.Children[types.GetSegmentPath(&exceedAction.Mark[i])] = types.YChild{"Mark", &exceedAction.Mark[i]}
    }
    exceedAction.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", exceedAction.ActionType}
    return &(exceedAction.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ExceedAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "exceed-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction) GetEntityData() *types.CommonEntityData {
    violateAction.EntityData.YFilter = violateAction.YFilter
    violateAction.EntityData.YangName = "violate-action"
    violateAction.EntityData.BundleName = "cisco_ios_xr"
    violateAction.EntityData.ParentYangName = "class"
    violateAction.EntityData.SegmentPath = "violate-action"
    violateAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    violateAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    violateAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    violateAction.EntityData.Children = make(map[string]types.YChild)
    violateAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range violateAction.Mark {
        violateAction.EntityData.Children[types.GetSegmentPath(&violateAction.Mark[i])] = types.YChild{"Mark", &violateAction.Mark[i]}
    }
    violateAction.EntityData.Leafs = make(map[string]types.YLeaf)
    violateAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", violateAction.ActionType}
    return &(violateAction.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_ViolateAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "violate-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_IpMark) GetEntityData() *types.CommonEntityData {
    ipMark.EntityData.YFilter = ipMark.YFilter
    ipMark.EntityData.YangName = "ip-mark"
    ipMark.EntityData.BundleName = "cisco_ios_xr"
    ipMark.EntityData.ParentYangName = "class"
    ipMark.EntityData.SegmentPath = "ip-mark"
    ipMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipMark.EntityData.Children = make(map[string]types.YChild)
    ipMark.EntityData.Leafs = make(map[string]types.YLeaf)
    ipMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", ipMark.MarkType}
    ipMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", ipMark.MarkValue}
    return &(ipMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_CommonMark) GetEntityData() *types.CommonEntityData {
    commonMark.EntityData.YFilter = commonMark.YFilter
    commonMark.EntityData.YangName = "common-mark"
    commonMark.EntityData.BundleName = "cisco_ios_xr"
    commonMark.EntityData.ParentYangName = "class"
    commonMark.EntityData.SegmentPath = "common-mark"
    commonMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonMark.EntityData.Children = make(map[string]types.YChild)
    commonMark.EntityData.Leafs = make(map[string]types.YLeaf)
    commonMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", commonMark.MarkType}
    commonMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", commonMark.MarkValue}
    return &(commonMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_MplsMark) GetEntityData() *types.CommonEntityData {
    mplsMark.EntityData.YFilter = mplsMark.YFilter
    mplsMark.EntityData.YangName = "mpls-mark"
    mplsMark.EntityData.BundleName = "cisco_ios_xr"
    mplsMark.EntityData.ParentYangName = "class"
    mplsMark.EntityData.SegmentPath = "mpls-mark"
    mplsMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsMark.EntityData.Children = make(map[string]types.YChild)
    mplsMark.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mplsMark.MarkType}
    mplsMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mplsMark.MarkValue}
    return &(mplsMark.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred struct {
    EntityData types.CommonEntityData
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

func (wred *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "class"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["wred-match-value"] = types.YChild{"WredMatchValue", &wred.WredMatchValue}
    wred.EntityData.Children["config-min-threshold"] = types.YChild{"ConfigMinThreshold", &wred.ConfigMinThreshold}
    wred.EntityData.Children["config-max-threshold"] = types.YChild{"ConfigMaxThreshold", &wred.ConfigMaxThreshold}
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["wred-match-type"] = types.YLeaf{"WredMatchType", wred.WredMatchType}
    wred.EntityData.Leafs["hardware-min-threshold-bytes"] = types.YLeaf{"HardwareMinThresholdBytes", wred.HardwareMinThresholdBytes}
    wred.EntityData.Leafs["hardware-max-threshold-bytes"] = types.YLeaf{"HardwareMaxThresholdBytes", wred.HardwareMaxThresholdBytes}
    wred.EntityData.Leafs["first-segment"] = types.YLeaf{"FirstSegment", wred.FirstSegment}
    wred.EntityData.Leafs["segment-size"] = types.YLeaf{"SegmentSize", wred.SegmentSize}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue) GetEntityData() *types.CommonEntityData {
    wredMatchValue.EntityData.YFilter = wredMatchValue.YFilter
    wredMatchValue.EntityData.YangName = "wred-match-value"
    wredMatchValue.EntityData.BundleName = "cisco_ios_xr"
    wredMatchValue.EntityData.ParentYangName = "wred"
    wredMatchValue.EntityData.SegmentPath = "wred-match-value"
    wredMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wredMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wredMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wredMatchValue.EntityData.Children = make(map[string]types.YChild)
    wredMatchValue.EntityData.Children["dnx-qosea-show-red-match-value"] = types.YChild{"DnxQoseaShowRedMatchValue", nil}
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        wredMatchValue.EntityData.Children[types.GetSegmentPath(&wredMatchValue.DnxQoseaShowRedMatchValue[i])] = types.YChild{"DnxQoseaShowRedMatchValue", &wredMatchValue.DnxQoseaShowRedMatchValue[i]}
    }
    wredMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wredMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetEntityData() *types.CommonEntityData {
    dnxQoseaShowRedMatchValue.EntityData.YFilter = dnxQoseaShowRedMatchValue.YFilter
    dnxQoseaShowRedMatchValue.EntityData.YangName = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.BundleName = "cisco_ios_xr"
    dnxQoseaShowRedMatchValue.EntityData.ParentYangName = "wred-match-value"
    dnxQoseaShowRedMatchValue.EntityData.SegmentPath = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnxQoseaShowRedMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnxQoseaShowRedMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnxQoseaShowRedMatchValue.EntityData.Children = make(map[string]types.YChild)
    dnxQoseaShowRedMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-start"] = types.YLeaf{"RangeStart", dnxQoseaShowRedMatchValue.RangeStart}
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-end"] = types.YLeaf{"RangeEnd", dnxQoseaShowRedMatchValue.RangeEnd}
    return &(dnxQoseaShowRedMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMinThreshold) GetEntityData() *types.CommonEntityData {
    configMinThreshold.EntityData.YFilter = configMinThreshold.YFilter
    configMinThreshold.EntityData.YangName = "config-min-threshold"
    configMinThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMinThreshold.EntityData.ParentYangName = "wred"
    configMinThreshold.EntityData.SegmentPath = "config-min-threshold"
    configMinThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinThreshold.EntityData.Children = make(map[string]types.YChild)
    configMinThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinThreshold.PolicyValue}
    configMinThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinThreshold.PolicyUnit}
    return &(configMinThreshold.EntityData)
}

// PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_Interfaces_Interface_Classes_Class_Wred_ConfigMaxThreshold) GetEntityData() *types.CommonEntityData {
    configMaxThreshold.EntityData.YFilter = configMaxThreshold.YFilter
    configMaxThreshold.EntityData.YangName = "config-max-threshold"
    configMaxThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMaxThreshold.EntityData.ParentYangName = "wred"
    configMaxThreshold.EntityData.SegmentPath = "config-max-threshold"
    configMaxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxThreshold.EntityData.Children = make(map[string]types.YChild)
    configMaxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxThreshold.PolicyValue}
    configMaxThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxThreshold.PolicyUnit}
    return &(configMaxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles
// QoS list of bundle interfaces
type PlatformQos_Nodes_Node_BundleInterfaceSingles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface names. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle.
    BundleInterfaceSingle []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle
}

func (bundleInterfaceSingles *PlatformQos_Nodes_Node_BundleInterfaceSingles) GetEntityData() *types.CommonEntityData {
    bundleInterfaceSingles.EntityData.YFilter = bundleInterfaceSingles.YFilter
    bundleInterfaceSingles.EntityData.YangName = "bundle-interface-singles"
    bundleInterfaceSingles.EntityData.BundleName = "cisco_ios_xr"
    bundleInterfaceSingles.EntityData.ParentYangName = "node"
    bundleInterfaceSingles.EntityData.SegmentPath = "bundle-interface-singles"
    bundleInterfaceSingles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterfaceSingles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterfaceSingles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterfaceSingles.EntityData.Children = make(map[string]types.YChild)
    bundleInterfaceSingles.EntityData.Children["bundle-interface-single"] = types.YChild{"BundleInterfaceSingle", nil}
    for i := range bundleInterfaceSingles.BundleInterfaceSingle {
        bundleInterfaceSingles.EntityData.Children[types.GetSegmentPath(&bundleInterfaceSingles.BundleInterfaceSingle[i])] = types.YChild{"BundleInterfaceSingle", &bundleInterfaceSingles.BundleInterfaceSingle[i]}
    }
    bundleInterfaceSingles.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bundleInterfaceSingles.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle
// QoS interface names
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Bundle interface name. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_PolicyDetails

    // QoS list of member interfaces.
    MemberInterfaces PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes
}

func (bundleInterfaceSingle *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle) GetEntityData() *types.CommonEntityData {
    bundleInterfaceSingle.EntityData.YFilter = bundleInterfaceSingle.YFilter
    bundleInterfaceSingle.EntityData.YangName = "bundle-interface-single"
    bundleInterfaceSingle.EntityData.BundleName = "cisco_ios_xr"
    bundleInterfaceSingle.EntityData.ParentYangName = "bundle-interface-singles"
    bundleInterfaceSingle.EntityData.SegmentPath = "bundle-interface-single" + "[interface-name='" + fmt.Sprintf("%v", bundleInterfaceSingle.InterfaceName) + "']"
    bundleInterfaceSingle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bundleInterfaceSingle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bundleInterfaceSingle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bundleInterfaceSingle.EntityData.Children = make(map[string]types.YChild)
    bundleInterfaceSingle.EntityData.Children["policy-details"] = types.YChild{"PolicyDetails", &bundleInterfaceSingle.PolicyDetails}
    bundleInterfaceSingle.EntityData.Children["member-interfaces"] = types.YChild{"MemberInterfaces", &bundleInterfaceSingle.MemberInterfaces}
    bundleInterfaceSingle.EntityData.Children["classes"] = types.YChild{"Classes", &bundleInterfaceSingle.Classes}
    bundleInterfaceSingle.EntityData.Leafs = make(map[string]types.YLeaf)
    bundleInterfaceSingle.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bundleInterfaceSingle.InterfaceName}
    return &(bundleInterfaceSingle.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_PolicyDetails struct {
    EntityData types.CommonEntityData
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

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_PolicyDetails) GetEntityData() *types.CommonEntityData {
    policyDetails.EntityData.YFilter = policyDetails.YFilter
    policyDetails.EntityData.YangName = "policy-details"
    policyDetails.EntityData.BundleName = "cisco_ios_xr"
    policyDetails.EntityData.ParentYangName = "bundle-interface-single"
    policyDetails.EntityData.SegmentPath = "policy-details"
    policyDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyDetails.EntityData.Children = make(map[string]types.YChild)
    policyDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    policyDetails.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", policyDetails.NpuId}
    policyDetails.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", policyDetails.InterfaceHandle}
    policyDetails.EntityData.Leafs["interface-bandwidth-kbps"] = types.YLeaf{"InterfaceBandwidthKbps", policyDetails.InterfaceBandwidthKbps}
    policyDetails.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", policyDetails.PolicyName}
    policyDetails.EntityData.Leafs["total-number-of-classes"] = types.YLeaf{"TotalNumberOfClasses", policyDetails.TotalNumberOfClasses}
    policyDetails.EntityData.Leafs["voq-base-address"] = types.YLeaf{"VoqBaseAddress", policyDetails.VoqBaseAddress}
    policyDetails.EntityData.Leafs["voq-stats-handle"] = types.YLeaf{"VoqStatsHandle", policyDetails.VoqStatsHandle}
    policyDetails.EntityData.Leafs["stats-accounting-type"] = types.YLeaf{"StatsAccountingType", policyDetails.StatsAccountingType}
    policyDetails.EntityData.Leafs["policy-status"] = types.YLeaf{"PolicyStatus", policyDetails.PolicyStatus}
    policyDetails.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", policyDetails.InterfaceStatus}
    return &(policyDetails.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces
// QoS list of member interfaces
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS interface names. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface.
    MemberInterface []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface
}

func (memberInterfaces *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces) GetEntityData() *types.CommonEntityData {
    memberInterfaces.EntityData.YFilter = memberInterfaces.YFilter
    memberInterfaces.EntityData.YangName = "member-interfaces"
    memberInterfaces.EntityData.BundleName = "cisco_ios_xr"
    memberInterfaces.EntityData.ParentYangName = "bundle-interface-single"
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

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface
// QoS interface names
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Member interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Policy Details.
    PolicyDetails PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_PolicyDetails

    // QoS list of class names.
    Classes PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes
}

func (memberInterface *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface) GetEntityData() *types.CommonEntityData {
    memberInterface.EntityData.YFilter = memberInterface.YFilter
    memberInterface.EntityData.YangName = "member-interface"
    memberInterface.EntityData.BundleName = "cisco_ios_xr"
    memberInterface.EntityData.ParentYangName = "member-interfaces"
    memberInterface.EntityData.SegmentPath = "member-interface" + "[interface-name='" + fmt.Sprintf("%v", memberInterface.InterfaceName) + "']"
    memberInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    memberInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    memberInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    memberInterface.EntityData.Children = make(map[string]types.YChild)
    memberInterface.EntityData.Children["policy-details"] = types.YChild{"PolicyDetails", &memberInterface.PolicyDetails}
    memberInterface.EntityData.Children["classes"] = types.YChild{"Classes", &memberInterface.Classes}
    memberInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    memberInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", memberInterface.InterfaceName}
    return &(memberInterface.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_PolicyDetails
// Policy Details
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_PolicyDetails struct {
    EntityData types.CommonEntityData
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

func (policyDetails *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_PolicyDetails) GetEntityData() *types.CommonEntityData {
    policyDetails.EntityData.YFilter = policyDetails.YFilter
    policyDetails.EntityData.YangName = "policy-details"
    policyDetails.EntityData.BundleName = "cisco_ios_xr"
    policyDetails.EntityData.ParentYangName = "member-interface"
    policyDetails.EntityData.SegmentPath = "policy-details"
    policyDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyDetails.EntityData.Children = make(map[string]types.YChild)
    policyDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    policyDetails.EntityData.Leafs["npu-id"] = types.YLeaf{"NpuId", policyDetails.NpuId}
    policyDetails.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", policyDetails.InterfaceHandle}
    policyDetails.EntityData.Leafs["interface-bandwidth-kbps"] = types.YLeaf{"InterfaceBandwidthKbps", policyDetails.InterfaceBandwidthKbps}
    policyDetails.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", policyDetails.PolicyName}
    policyDetails.EntityData.Leafs["total-number-of-classes"] = types.YLeaf{"TotalNumberOfClasses", policyDetails.TotalNumberOfClasses}
    policyDetails.EntityData.Leafs["voq-base-address"] = types.YLeaf{"VoqBaseAddress", policyDetails.VoqBaseAddress}
    policyDetails.EntityData.Leafs["voq-stats-handle"] = types.YLeaf{"VoqStatsHandle", policyDetails.VoqStatsHandle}
    policyDetails.EntityData.Leafs["stats-accounting-type"] = types.YLeaf{"StatsAccountingType", policyDetails.StatsAccountingType}
    policyDetails.EntityData.Leafs["policy-status"] = types.YLeaf{"PolicyStatus", policyDetails.PolicyStatus}
    policyDetails.EntityData.Leafs["interface-status"] = types.YLeaf{"InterfaceStatus", policyDetails.InterfaceStatus}
    return &(policyDetails.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class.
    Class []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "member-interface"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class struct {
    EntityData types.CommonEntityData
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
    ConfigMaxRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate

    // Configured minimum rate.
    ConfigMinRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate

    // Configured queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit

    // Configured policer average rate.
    ConfigPolicerAverageRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate

    // Config policer peak rate.
    ConfigPolicerPeakRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate

    // Configured policer conform burst.
    ConfigPolicerConformBurst PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst

    // Configured policer excess burst.
    ConfigPolicerExcessBurst PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst

    // Conform action.
    ConformAction PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction

    // Exceed action.
    ExceedAction PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction

    // Violate action.
    ViolateAction PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction

    // IP mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_IpMark.
    IpMark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_IpMark

    // Common mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_CommonMark.
    CommonMark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_CommonMark

    // MPLS mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_MplsMark.
    MplsMark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_MplsMark

    // WRED parameters. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred.
    Wred []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred
}

func (class *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["config-max-rate"] = types.YChild{"ConfigMaxRate", &class.ConfigMaxRate}
    class.EntityData.Children["config-min-rate"] = types.YChild{"ConfigMinRate", &class.ConfigMinRate}
    class.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &class.ConfigQueueLimit}
    class.EntityData.Children["config-policer-average-rate"] = types.YChild{"ConfigPolicerAverageRate", &class.ConfigPolicerAverageRate}
    class.EntityData.Children["config-policer-peak-rate"] = types.YChild{"ConfigPolicerPeakRate", &class.ConfigPolicerPeakRate}
    class.EntityData.Children["config-policer-conform-burst"] = types.YChild{"ConfigPolicerConformBurst", &class.ConfigPolicerConformBurst}
    class.EntityData.Children["config-policer-excess-burst"] = types.YChild{"ConfigPolicerExcessBurst", &class.ConfigPolicerExcessBurst}
    class.EntityData.Children["conform-action"] = types.YChild{"ConformAction", &class.ConformAction}
    class.EntityData.Children["exceed-action"] = types.YChild{"ExceedAction", &class.ExceedAction}
    class.EntityData.Children["violate-action"] = types.YChild{"ViolateAction", &class.ViolateAction}
    class.EntityData.Children["ip-mark"] = types.YChild{"IpMark", nil}
    for i := range class.IpMark {
        class.EntityData.Children[types.GetSegmentPath(&class.IpMark[i])] = types.YChild{"IpMark", &class.IpMark[i]}
    }
    class.EntityData.Children["common-mark"] = types.YChild{"CommonMark", nil}
    for i := range class.CommonMark {
        class.EntityData.Children[types.GetSegmentPath(&class.CommonMark[i])] = types.YChild{"CommonMark", &class.CommonMark[i]}
    }
    class.EntityData.Children["mpls-mark"] = types.YChild{"MplsMark", nil}
    for i := range class.MplsMark {
        class.EntityData.Children[types.GetSegmentPath(&class.MplsMark[i])] = types.YChild{"MplsMark", &class.MplsMark[i]}
    }
    class.EntityData.Children["wred"] = types.YChild{"Wred", nil}
    for i := range class.Wred {
        class.EntityData.Children[types.GetSegmentPath(&class.Wred[i])] = types.YChild{"Wred", &class.Wred[i]}
    }
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["level-one-class-name"] = types.YLeaf{"LevelOneClassName", class.LevelOneClassName}
    class.EntityData.Leafs["level-two-class-name"] = types.YLeaf{"LevelTwoClassName", class.LevelTwoClassName}
    class.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", class.ClassLevel}
    class.EntityData.Leafs["egress-queue-id"] = types.YLeaf{"EgressQueueId", class.EgressQueueId}
    class.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", class.QueueType}
    class.EntityData.Leafs["priority-level"] = types.YLeaf{"PriorityLevel", class.PriorityLevel}
    class.EntityData.Leafs["hardware-max-rate-kbps"] = types.YLeaf{"HardwareMaxRateKbps", class.HardwareMaxRateKbps}
    class.EntityData.Leafs["hardware-min-rate-kbps"] = types.YLeaf{"HardwareMinRateKbps", class.HardwareMinRateKbps}
    class.EntityData.Leafs["config-excess-bandwidth-percent"] = types.YLeaf{"ConfigExcessBandwidthPercent", class.ConfigExcessBandwidthPercent}
    class.EntityData.Leafs["config-excess-bandwidth-unit"] = types.YLeaf{"ConfigExcessBandwidthUnit", class.ConfigExcessBandwidthUnit}
    class.EntityData.Leafs["hardware-excess-bandwidth-weight"] = types.YLeaf{"HardwareExcessBandwidthWeight", class.HardwareExcessBandwidthWeight}
    class.EntityData.Leafs["network-min-bandwidth-kbps"] = types.YLeaf{"NetworkMinBandwidthKbps", class.NetworkMinBandwidthKbps}
    class.EntityData.Leafs["hardware-queue-limit-bytes"] = types.YLeaf{"HardwareQueueLimitBytes", class.HardwareQueueLimitBytes}
    class.EntityData.Leafs["hardware-queue-limit-microseconds"] = types.YLeaf{"HardwareQueueLimitMicroseconds", class.HardwareQueueLimitMicroseconds}
    class.EntityData.Leafs["policer-bucket-id"] = types.YLeaf{"PolicerBucketId", class.PolicerBucketId}
    class.EntityData.Leafs["policer-stats-handle"] = types.YLeaf{"PolicerStatsHandle", class.PolicerStatsHandle}
    class.EntityData.Leafs["hardware-policer-average-rate-kbps"] = types.YLeaf{"HardwarePolicerAverageRateKbps", class.HardwarePolicerAverageRateKbps}
    class.EntityData.Leafs["hardware-policer-peak-rate-kbps"] = types.YLeaf{"HardwarePolicerPeakRateKbps", class.HardwarePolicerPeakRateKbps}
    class.EntityData.Leafs["hardware-policer-conform-burst-bytes"] = types.YLeaf{"HardwarePolicerConformBurstBytes", class.HardwarePolicerConformBurstBytes}
    class.EntityData.Leafs["hardware-policer-excess-burst-bytes"] = types.YLeaf{"HardwarePolicerExcessBurstBytes", class.HardwarePolicerExcessBurstBytes}
    return &(class.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMaxRate) GetEntityData() *types.CommonEntityData {
    configMaxRate.EntityData.YFilter = configMaxRate.YFilter
    configMaxRate.EntityData.YangName = "config-max-rate"
    configMaxRate.EntityData.BundleName = "cisco_ios_xr"
    configMaxRate.EntityData.ParentYangName = "class"
    configMaxRate.EntityData.SegmentPath = "config-max-rate"
    configMaxRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxRate.EntityData.Children = make(map[string]types.YChild)
    configMaxRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxRate.PolicyValue}
    configMaxRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxRate.PolicyUnit}
    return &(configMaxRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigMinRate) GetEntityData() *types.CommonEntityData {
    configMinRate.EntityData.YFilter = configMinRate.YFilter
    configMinRate.EntityData.YangName = "config-min-rate"
    configMinRate.EntityData.BundleName = "cisco_ios_xr"
    configMinRate.EntityData.ParentYangName = "class"
    configMinRate.EntityData.SegmentPath = "config-min-rate"
    configMinRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinRate.EntityData.Children = make(map[string]types.YChild)
    configMinRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinRate.PolicyValue}
    configMinRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinRate.PolicyUnit}
    return &(configMinRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "class"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configQueueLimit.PolicyValue}
    configQueueLimit.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configQueueLimit.PolicyUnit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerAverageRate) GetEntityData() *types.CommonEntityData {
    configPolicerAverageRate.EntityData.YFilter = configPolicerAverageRate.YFilter
    configPolicerAverageRate.EntityData.YangName = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerAverageRate.EntityData.ParentYangName = "class"
    configPolicerAverageRate.EntityData.SegmentPath = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerAverageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerAverageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerAverageRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerAverageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerAverageRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerAverageRate.PolicyValue}
    configPolicerAverageRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerAverageRate.PolicyUnit}
    return &(configPolicerAverageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerPeakRate) GetEntityData() *types.CommonEntityData {
    configPolicerPeakRate.EntityData.YFilter = configPolicerPeakRate.YFilter
    configPolicerPeakRate.EntityData.YangName = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerPeakRate.EntityData.ParentYangName = "class"
    configPolicerPeakRate.EntityData.SegmentPath = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerPeakRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerPeakRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerPeakRate.PolicyValue}
    configPolicerPeakRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerPeakRate.PolicyUnit}
    return &(configPolicerPeakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerConformBurst) GetEntityData() *types.CommonEntityData {
    configPolicerConformBurst.EntityData.YFilter = configPolicerConformBurst.YFilter
    configPolicerConformBurst.EntityData.YangName = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerConformBurst.EntityData.ParentYangName = "class"
    configPolicerConformBurst.EntityData.SegmentPath = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerConformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerConformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerConformBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerConformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerConformBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerConformBurst.PolicyValue}
    configPolicerConformBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerConformBurst.PolicyUnit}
    return &(configPolicerConformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConfigPolicerExcessBurst) GetEntityData() *types.CommonEntityData {
    configPolicerExcessBurst.EntityData.YFilter = configPolicerExcessBurst.YFilter
    configPolicerExcessBurst.EntityData.YangName = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerExcessBurst.EntityData.ParentYangName = "class"
    configPolicerExcessBurst.EntityData.SegmentPath = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerExcessBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerExcessBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerExcessBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerExcessBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerExcessBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerExcessBurst.PolicyValue}
    configPolicerExcessBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerExcessBurst.PolicyUnit}
    return &(configPolicerExcessBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction) GetEntityData() *types.CommonEntityData {
    conformAction.EntityData.YFilter = conformAction.YFilter
    conformAction.EntityData.YangName = "conform-action"
    conformAction.EntityData.BundleName = "cisco_ios_xr"
    conformAction.EntityData.ParentYangName = "class"
    conformAction.EntityData.SegmentPath = "conform-action"
    conformAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformAction.EntityData.Children = make(map[string]types.YChild)
    conformAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range conformAction.Mark {
        conformAction.EntityData.Children[types.GetSegmentPath(&conformAction.Mark[i])] = types.YChild{"Mark", &conformAction.Mark[i]}
    }
    conformAction.EntityData.Leafs = make(map[string]types.YLeaf)
    conformAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", conformAction.ActionType}
    return &(conformAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ConformAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "conform-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction) GetEntityData() *types.CommonEntityData {
    exceedAction.EntityData.YFilter = exceedAction.YFilter
    exceedAction.EntityData.YangName = "exceed-action"
    exceedAction.EntityData.BundleName = "cisco_ios_xr"
    exceedAction.EntityData.ParentYangName = "class"
    exceedAction.EntityData.SegmentPath = "exceed-action"
    exceedAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedAction.EntityData.Children = make(map[string]types.YChild)
    exceedAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range exceedAction.Mark {
        exceedAction.EntityData.Children[types.GetSegmentPath(&exceedAction.Mark[i])] = types.YChild{"Mark", &exceedAction.Mark[i]}
    }
    exceedAction.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", exceedAction.ActionType}
    return &(exceedAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ExceedAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "exceed-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction) GetEntityData() *types.CommonEntityData {
    violateAction.EntityData.YFilter = violateAction.YFilter
    violateAction.EntityData.YangName = "violate-action"
    violateAction.EntityData.BundleName = "cisco_ios_xr"
    violateAction.EntityData.ParentYangName = "class"
    violateAction.EntityData.SegmentPath = "violate-action"
    violateAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    violateAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    violateAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    violateAction.EntityData.Children = make(map[string]types.YChild)
    violateAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range violateAction.Mark {
        violateAction.EntityData.Children[types.GetSegmentPath(&violateAction.Mark[i])] = types.YChild{"Mark", &violateAction.Mark[i]}
    }
    violateAction.EntityData.Leafs = make(map[string]types.YLeaf)
    violateAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", violateAction.ActionType}
    return &(violateAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_ViolateAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "violate-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_IpMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_IpMark) GetEntityData() *types.CommonEntityData {
    ipMark.EntityData.YFilter = ipMark.YFilter
    ipMark.EntityData.YangName = "ip-mark"
    ipMark.EntityData.BundleName = "cisco_ios_xr"
    ipMark.EntityData.ParentYangName = "class"
    ipMark.EntityData.SegmentPath = "ip-mark"
    ipMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipMark.EntityData.Children = make(map[string]types.YChild)
    ipMark.EntityData.Leafs = make(map[string]types.YLeaf)
    ipMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", ipMark.MarkType}
    ipMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", ipMark.MarkValue}
    return &(ipMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_CommonMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_CommonMark) GetEntityData() *types.CommonEntityData {
    commonMark.EntityData.YFilter = commonMark.YFilter
    commonMark.EntityData.YangName = "common-mark"
    commonMark.EntityData.BundleName = "cisco_ios_xr"
    commonMark.EntityData.ParentYangName = "class"
    commonMark.EntityData.SegmentPath = "common-mark"
    commonMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonMark.EntityData.Children = make(map[string]types.YChild)
    commonMark.EntityData.Leafs = make(map[string]types.YLeaf)
    commonMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", commonMark.MarkType}
    commonMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", commonMark.MarkValue}
    return &(commonMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_MplsMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_MplsMark) GetEntityData() *types.CommonEntityData {
    mplsMark.EntityData.YFilter = mplsMark.YFilter
    mplsMark.EntityData.YangName = "mpls-mark"
    mplsMark.EntityData.BundleName = "cisco_ios_xr"
    mplsMark.EntityData.ParentYangName = "class"
    mplsMark.EntityData.SegmentPath = "mpls-mark"
    mplsMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsMark.EntityData.Children = make(map[string]types.YChild)
    mplsMark.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mplsMark.MarkType}
    mplsMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mplsMark.MarkValue}
    return &(mplsMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred struct {
    EntityData types.CommonEntityData
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
    WredMatchValue PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue

    // Configured minimum threshold.
    ConfigMinThreshold PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold

    // Configured maximum threshold.
    ConfigMaxThreshold PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "class"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["wred-match-value"] = types.YChild{"WredMatchValue", &wred.WredMatchValue}
    wred.EntityData.Children["config-min-threshold"] = types.YChild{"ConfigMinThreshold", &wred.ConfigMinThreshold}
    wred.EntityData.Children["config-max-threshold"] = types.YChild{"ConfigMaxThreshold", &wred.ConfigMaxThreshold}
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["wred-match-type"] = types.YLeaf{"WredMatchType", wred.WredMatchType}
    wred.EntityData.Leafs["hardware-min-threshold-bytes"] = types.YLeaf{"HardwareMinThresholdBytes", wred.HardwareMinThresholdBytes}
    wred.EntityData.Leafs["hardware-max-threshold-bytes"] = types.YLeaf{"HardwareMaxThresholdBytes", wred.HardwareMaxThresholdBytes}
    wred.EntityData.Leafs["first-segment"] = types.YLeaf{"FirstSegment", wred.FirstSegment}
    wred.EntityData.Leafs["segment-size"] = types.YLeaf{"SegmentSize", wred.SegmentSize}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue) GetEntityData() *types.CommonEntityData {
    wredMatchValue.EntityData.YFilter = wredMatchValue.YFilter
    wredMatchValue.EntityData.YangName = "wred-match-value"
    wredMatchValue.EntityData.BundleName = "cisco_ios_xr"
    wredMatchValue.EntityData.ParentYangName = "wred"
    wredMatchValue.EntityData.SegmentPath = "wred-match-value"
    wredMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wredMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wredMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wredMatchValue.EntityData.Children = make(map[string]types.YChild)
    wredMatchValue.EntityData.Children["dnx-qosea-show-red-match-value"] = types.YChild{"DnxQoseaShowRedMatchValue", nil}
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        wredMatchValue.EntityData.Children[types.GetSegmentPath(&wredMatchValue.DnxQoseaShowRedMatchValue[i])] = types.YChild{"DnxQoseaShowRedMatchValue", &wredMatchValue.DnxQoseaShowRedMatchValue[i]}
    }
    wredMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wredMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetEntityData() *types.CommonEntityData {
    dnxQoseaShowRedMatchValue.EntityData.YFilter = dnxQoseaShowRedMatchValue.YFilter
    dnxQoseaShowRedMatchValue.EntityData.YangName = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.BundleName = "cisco_ios_xr"
    dnxQoseaShowRedMatchValue.EntityData.ParentYangName = "wred-match-value"
    dnxQoseaShowRedMatchValue.EntityData.SegmentPath = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnxQoseaShowRedMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnxQoseaShowRedMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnxQoseaShowRedMatchValue.EntityData.Children = make(map[string]types.YChild)
    dnxQoseaShowRedMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-start"] = types.YLeaf{"RangeStart", dnxQoseaShowRedMatchValue.RangeStart}
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-end"] = types.YLeaf{"RangeEnd", dnxQoseaShowRedMatchValue.RangeEnd}
    return &(dnxQoseaShowRedMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMinThreshold) GetEntityData() *types.CommonEntityData {
    configMinThreshold.EntityData.YFilter = configMinThreshold.YFilter
    configMinThreshold.EntityData.YangName = "config-min-threshold"
    configMinThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMinThreshold.EntityData.ParentYangName = "wred"
    configMinThreshold.EntityData.SegmentPath = "config-min-threshold"
    configMinThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinThreshold.EntityData.Children = make(map[string]types.YChild)
    configMinThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinThreshold.PolicyValue}
    configMinThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinThreshold.PolicyUnit}
    return &(configMinThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_MemberInterfaces_MemberInterface_Classes_Class_Wred_ConfigMaxThreshold) GetEntityData() *types.CommonEntityData {
    configMaxThreshold.EntityData.YFilter = configMaxThreshold.YFilter
    configMaxThreshold.EntityData.YangName = "config-max-threshold"
    configMaxThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMaxThreshold.EntityData.ParentYangName = "wred"
    configMaxThreshold.EntityData.SegmentPath = "config-max-threshold"
    configMaxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxThreshold.EntityData.Children = make(map[string]types.YChild)
    configMaxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxThreshold.PolicyValue}
    configMaxThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxThreshold.PolicyUnit}
    return &(configMaxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes
// QoS list of class names
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS policy class. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class.
    Class []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class
}

func (classes *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "bundle-interface-single"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class
// QoS policy class
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class struct {
    EntityData types.CommonEntityData
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
    ConfigMaxRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMaxRate

    // Configured minimum rate.
    ConfigMinRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMinRate

    // Configured queue limit.
    ConfigQueueLimit PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigQueueLimit

    // Configured policer average rate.
    ConfigPolicerAverageRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerAverageRate

    // Config policer peak rate.
    ConfigPolicerPeakRate PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerPeakRate

    // Configured policer conform burst.
    ConfigPolicerConformBurst PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerConformBurst

    // Configured policer excess burst.
    ConfigPolicerExcessBurst PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerExcessBurst

    // Conform action.
    ConformAction PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction

    // Exceed action.
    ExceedAction PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction

    // Violate action.
    ViolateAction PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction

    // IP mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_IpMark.
    IpMark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_IpMark

    // Common mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_CommonMark.
    CommonMark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_CommonMark

    // MPLS mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_MplsMark.
    MplsMark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_MplsMark

    // WRED parameters. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred.
    Wred []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred
}

func (class *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[level-one-class-name='" + fmt.Sprintf("%v", class.LevelOneClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["config-max-rate"] = types.YChild{"ConfigMaxRate", &class.ConfigMaxRate}
    class.EntityData.Children["config-min-rate"] = types.YChild{"ConfigMinRate", &class.ConfigMinRate}
    class.EntityData.Children["config-queue-limit"] = types.YChild{"ConfigQueueLimit", &class.ConfigQueueLimit}
    class.EntityData.Children["config-policer-average-rate"] = types.YChild{"ConfigPolicerAverageRate", &class.ConfigPolicerAverageRate}
    class.EntityData.Children["config-policer-peak-rate"] = types.YChild{"ConfigPolicerPeakRate", &class.ConfigPolicerPeakRate}
    class.EntityData.Children["config-policer-conform-burst"] = types.YChild{"ConfigPolicerConformBurst", &class.ConfigPolicerConformBurst}
    class.EntityData.Children["config-policer-excess-burst"] = types.YChild{"ConfigPolicerExcessBurst", &class.ConfigPolicerExcessBurst}
    class.EntityData.Children["conform-action"] = types.YChild{"ConformAction", &class.ConformAction}
    class.EntityData.Children["exceed-action"] = types.YChild{"ExceedAction", &class.ExceedAction}
    class.EntityData.Children["violate-action"] = types.YChild{"ViolateAction", &class.ViolateAction}
    class.EntityData.Children["ip-mark"] = types.YChild{"IpMark", nil}
    for i := range class.IpMark {
        class.EntityData.Children[types.GetSegmentPath(&class.IpMark[i])] = types.YChild{"IpMark", &class.IpMark[i]}
    }
    class.EntityData.Children["common-mark"] = types.YChild{"CommonMark", nil}
    for i := range class.CommonMark {
        class.EntityData.Children[types.GetSegmentPath(&class.CommonMark[i])] = types.YChild{"CommonMark", &class.CommonMark[i]}
    }
    class.EntityData.Children["mpls-mark"] = types.YChild{"MplsMark", nil}
    for i := range class.MplsMark {
        class.EntityData.Children[types.GetSegmentPath(&class.MplsMark[i])] = types.YChild{"MplsMark", &class.MplsMark[i]}
    }
    class.EntityData.Children["wred"] = types.YChild{"Wred", nil}
    for i := range class.Wred {
        class.EntityData.Children[types.GetSegmentPath(&class.Wred[i])] = types.YChild{"Wred", &class.Wred[i]}
    }
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["level-one-class-name"] = types.YLeaf{"LevelOneClassName", class.LevelOneClassName}
    class.EntityData.Leafs["level-two-class-name"] = types.YLeaf{"LevelTwoClassName", class.LevelTwoClassName}
    class.EntityData.Leafs["class-level"] = types.YLeaf{"ClassLevel", class.ClassLevel}
    class.EntityData.Leafs["egress-queue-id"] = types.YLeaf{"EgressQueueId", class.EgressQueueId}
    class.EntityData.Leafs["queue-type"] = types.YLeaf{"QueueType", class.QueueType}
    class.EntityData.Leafs["priority-level"] = types.YLeaf{"PriorityLevel", class.PriorityLevel}
    class.EntityData.Leafs["hardware-max-rate-kbps"] = types.YLeaf{"HardwareMaxRateKbps", class.HardwareMaxRateKbps}
    class.EntityData.Leafs["hardware-min-rate-kbps"] = types.YLeaf{"HardwareMinRateKbps", class.HardwareMinRateKbps}
    class.EntityData.Leafs["config-excess-bandwidth-percent"] = types.YLeaf{"ConfigExcessBandwidthPercent", class.ConfigExcessBandwidthPercent}
    class.EntityData.Leafs["config-excess-bandwidth-unit"] = types.YLeaf{"ConfigExcessBandwidthUnit", class.ConfigExcessBandwidthUnit}
    class.EntityData.Leafs["hardware-excess-bandwidth-weight"] = types.YLeaf{"HardwareExcessBandwidthWeight", class.HardwareExcessBandwidthWeight}
    class.EntityData.Leafs["network-min-bandwidth-kbps"] = types.YLeaf{"NetworkMinBandwidthKbps", class.NetworkMinBandwidthKbps}
    class.EntityData.Leafs["hardware-queue-limit-bytes"] = types.YLeaf{"HardwareQueueLimitBytes", class.HardwareQueueLimitBytes}
    class.EntityData.Leafs["hardware-queue-limit-microseconds"] = types.YLeaf{"HardwareQueueLimitMicroseconds", class.HardwareQueueLimitMicroseconds}
    class.EntityData.Leafs["policer-bucket-id"] = types.YLeaf{"PolicerBucketId", class.PolicerBucketId}
    class.EntityData.Leafs["policer-stats-handle"] = types.YLeaf{"PolicerStatsHandle", class.PolicerStatsHandle}
    class.EntityData.Leafs["hardware-policer-average-rate-kbps"] = types.YLeaf{"HardwarePolicerAverageRateKbps", class.HardwarePolicerAverageRateKbps}
    class.EntityData.Leafs["hardware-policer-peak-rate-kbps"] = types.YLeaf{"HardwarePolicerPeakRateKbps", class.HardwarePolicerPeakRateKbps}
    class.EntityData.Leafs["hardware-policer-conform-burst-bytes"] = types.YLeaf{"HardwarePolicerConformBurstBytes", class.HardwarePolicerConformBurstBytes}
    class.EntityData.Leafs["hardware-policer-excess-burst-bytes"] = types.YLeaf{"HardwarePolicerExcessBurstBytes", class.HardwarePolicerExcessBurstBytes}
    return &(class.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMaxRate
// Configured maximum rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMaxRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMaxRate) GetEntityData() *types.CommonEntityData {
    configMaxRate.EntityData.YFilter = configMaxRate.YFilter
    configMaxRate.EntityData.YangName = "config-max-rate"
    configMaxRate.EntityData.BundleName = "cisco_ios_xr"
    configMaxRate.EntityData.ParentYangName = "class"
    configMaxRate.EntityData.SegmentPath = "config-max-rate"
    configMaxRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxRate.EntityData.Children = make(map[string]types.YChild)
    configMaxRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxRate.PolicyValue}
    configMaxRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxRate.PolicyUnit}
    return &(configMaxRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMinRate
// Configured minimum rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMinRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigMinRate) GetEntityData() *types.CommonEntityData {
    configMinRate.EntityData.YFilter = configMinRate.YFilter
    configMinRate.EntityData.YangName = "config-min-rate"
    configMinRate.EntityData.BundleName = "cisco_ios_xr"
    configMinRate.EntityData.ParentYangName = "class"
    configMinRate.EntityData.SegmentPath = "config-min-rate"
    configMinRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinRate.EntityData.Children = make(map[string]types.YChild)
    configMinRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinRate.PolicyValue}
    configMinRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinRate.PolicyUnit}
    return &(configMinRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigQueueLimit
// Configured queue limit
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigQueueLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configQueueLimit *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigQueueLimit) GetEntityData() *types.CommonEntityData {
    configQueueLimit.EntityData.YFilter = configQueueLimit.YFilter
    configQueueLimit.EntityData.YangName = "config-queue-limit"
    configQueueLimit.EntityData.BundleName = "cisco_ios_xr"
    configQueueLimit.EntityData.ParentYangName = "class"
    configQueueLimit.EntityData.SegmentPath = "config-queue-limit"
    configQueueLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configQueueLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configQueueLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configQueueLimit.EntityData.Children = make(map[string]types.YChild)
    configQueueLimit.EntityData.Leafs = make(map[string]types.YLeaf)
    configQueueLimit.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configQueueLimit.PolicyValue}
    configQueueLimit.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configQueueLimit.PolicyUnit}
    return &(configQueueLimit.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerAverageRate
// Configured policer average rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerAverageRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerAverageRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerAverageRate) GetEntityData() *types.CommonEntityData {
    configPolicerAverageRate.EntityData.YFilter = configPolicerAverageRate.YFilter
    configPolicerAverageRate.EntityData.YangName = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerAverageRate.EntityData.ParentYangName = "class"
    configPolicerAverageRate.EntityData.SegmentPath = "config-policer-average-rate"
    configPolicerAverageRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerAverageRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerAverageRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerAverageRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerAverageRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerAverageRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerAverageRate.PolicyValue}
    configPolicerAverageRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerAverageRate.PolicyUnit}
    return &(configPolicerAverageRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerPeakRate
// Config policer peak rate
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerPeakRate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerPeakRate *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerPeakRate) GetEntityData() *types.CommonEntityData {
    configPolicerPeakRate.EntityData.YFilter = configPolicerPeakRate.YFilter
    configPolicerPeakRate.EntityData.YangName = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.BundleName = "cisco_ios_xr"
    configPolicerPeakRate.EntityData.ParentYangName = "class"
    configPolicerPeakRate.EntityData.SegmentPath = "config-policer-peak-rate"
    configPolicerPeakRate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerPeakRate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerPeakRate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerPeakRate.EntityData.Children = make(map[string]types.YChild)
    configPolicerPeakRate.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerPeakRate.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerPeakRate.PolicyValue}
    configPolicerPeakRate.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerPeakRate.PolicyUnit}
    return &(configPolicerPeakRate.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerConformBurst
// Configured policer conform burst
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerConformBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerConformBurst *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerConformBurst) GetEntityData() *types.CommonEntityData {
    configPolicerConformBurst.EntityData.YFilter = configPolicerConformBurst.YFilter
    configPolicerConformBurst.EntityData.YangName = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerConformBurst.EntityData.ParentYangName = "class"
    configPolicerConformBurst.EntityData.SegmentPath = "config-policer-conform-burst"
    configPolicerConformBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerConformBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerConformBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerConformBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerConformBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerConformBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerConformBurst.PolicyValue}
    configPolicerConformBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerConformBurst.PolicyUnit}
    return &(configPolicerConformBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerExcessBurst
// Configured policer excess burst
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerExcessBurst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configPolicerExcessBurst *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConfigPolicerExcessBurst) GetEntityData() *types.CommonEntityData {
    configPolicerExcessBurst.EntityData.YFilter = configPolicerExcessBurst.YFilter
    configPolicerExcessBurst.EntityData.YangName = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.BundleName = "cisco_ios_xr"
    configPolicerExcessBurst.EntityData.ParentYangName = "class"
    configPolicerExcessBurst.EntityData.SegmentPath = "config-policer-excess-burst"
    configPolicerExcessBurst.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configPolicerExcessBurst.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configPolicerExcessBurst.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configPolicerExcessBurst.EntityData.Children = make(map[string]types.YChild)
    configPolicerExcessBurst.EntityData.Leafs = make(map[string]types.YLeaf)
    configPolicerExcessBurst.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configPolicerExcessBurst.PolicyValue}
    configPolicerExcessBurst.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configPolicerExcessBurst.PolicyUnit}
    return &(configPolicerExcessBurst.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction
// Conform action
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction_Mark
}

func (conformAction *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction) GetEntityData() *types.CommonEntityData {
    conformAction.EntityData.YFilter = conformAction.YFilter
    conformAction.EntityData.YangName = "conform-action"
    conformAction.EntityData.BundleName = "cisco_ios_xr"
    conformAction.EntityData.ParentYangName = "class"
    conformAction.EntityData.SegmentPath = "conform-action"
    conformAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    conformAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    conformAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    conformAction.EntityData.Children = make(map[string]types.YChild)
    conformAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range conformAction.Mark {
        conformAction.EntityData.Children[types.GetSegmentPath(&conformAction.Mark[i])] = types.YChild{"Mark", &conformAction.Mark[i]}
    }
    conformAction.EntityData.Leafs = make(map[string]types.YLeaf)
    conformAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", conformAction.ActionType}
    return &(conformAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ConformAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "conform-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction
// Exceed action
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction_Mark
}

func (exceedAction *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction) GetEntityData() *types.CommonEntityData {
    exceedAction.EntityData.YFilter = exceedAction.YFilter
    exceedAction.EntityData.YangName = "exceed-action"
    exceedAction.EntityData.BundleName = "cisco_ios_xr"
    exceedAction.EntityData.ParentYangName = "class"
    exceedAction.EntityData.SegmentPath = "exceed-action"
    exceedAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exceedAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exceedAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exceedAction.EntityData.Children = make(map[string]types.YChild)
    exceedAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range exceedAction.Mark {
        exceedAction.EntityData.Children[types.GetSegmentPath(&exceedAction.Mark[i])] = types.YChild{"Mark", &exceedAction.Mark[i]}
    }
    exceedAction.EntityData.Leafs = make(map[string]types.YLeaf)
    exceedAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", exceedAction.ActionType}
    return &(exceedAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ExceedAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "exceed-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction
// Violate action
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policer action type. The type is DnxQoseaShowAction.
    ActionType interface{}

    // Action mark. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction_Mark.
    Mark []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction_Mark
}

func (violateAction *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction) GetEntityData() *types.CommonEntityData {
    violateAction.EntityData.YFilter = violateAction.YFilter
    violateAction.EntityData.YangName = "violate-action"
    violateAction.EntityData.BundleName = "cisco_ios_xr"
    violateAction.EntityData.ParentYangName = "class"
    violateAction.EntityData.SegmentPath = "violate-action"
    violateAction.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    violateAction.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    violateAction.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    violateAction.EntityData.Children = make(map[string]types.YChild)
    violateAction.EntityData.Children["mark"] = types.YChild{"Mark", nil}
    for i := range violateAction.Mark {
        violateAction.EntityData.Children[types.GetSegmentPath(&violateAction.Mark[i])] = types.YChild{"Mark", &violateAction.Mark[i]}
    }
    violateAction.EntityData.Leafs = make(map[string]types.YLeaf)
    violateAction.EntityData.Leafs["action-type"] = types.YLeaf{"ActionType", violateAction.ActionType}
    return &(violateAction.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction_Mark
// Action mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction_Mark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_ViolateAction_Mark) GetEntityData() *types.CommonEntityData {
    mark.EntityData.YFilter = mark.YFilter
    mark.EntityData.YangName = "mark"
    mark.EntityData.BundleName = "cisco_ios_xr"
    mark.EntityData.ParentYangName = "violate-action"
    mark.EntityData.SegmentPath = "mark"
    mark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mark.EntityData.Children = make(map[string]types.YChild)
    mark.EntityData.Leafs = make(map[string]types.YLeaf)
    mark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mark.MarkType}
    mark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mark.MarkValue}
    return &(mark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_IpMark
// IP mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_IpMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (ipMark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_IpMark) GetEntityData() *types.CommonEntityData {
    ipMark.EntityData.YFilter = ipMark.YFilter
    ipMark.EntityData.YangName = "ip-mark"
    ipMark.EntityData.BundleName = "cisco_ios_xr"
    ipMark.EntityData.ParentYangName = "class"
    ipMark.EntityData.SegmentPath = "ip-mark"
    ipMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipMark.EntityData.Children = make(map[string]types.YChild)
    ipMark.EntityData.Leafs = make(map[string]types.YLeaf)
    ipMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", ipMark.MarkType}
    ipMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", ipMark.MarkValue}
    return &(ipMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_CommonMark
// Common mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_CommonMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (commonMark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_CommonMark) GetEntityData() *types.CommonEntityData {
    commonMark.EntityData.YFilter = commonMark.YFilter
    commonMark.EntityData.YangName = "common-mark"
    commonMark.EntityData.BundleName = "cisco_ios_xr"
    commonMark.EntityData.ParentYangName = "class"
    commonMark.EntityData.SegmentPath = "common-mark"
    commonMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commonMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commonMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commonMark.EntityData.Children = make(map[string]types.YChild)
    commonMark.EntityData.Leafs = make(map[string]types.YLeaf)
    commonMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", commonMark.MarkType}
    commonMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", commonMark.MarkValue}
    return &(commonMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_MplsMark
// MPLS mark
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_MplsMark struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Mark type. The type is DnxQoseaShowMark.
    MarkType interface{}

    // Mark value. The type is interface{} with range: 0..65535.
    MarkValue interface{}
}

func (mplsMark *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_MplsMark) GetEntityData() *types.CommonEntityData {
    mplsMark.EntityData.YFilter = mplsMark.YFilter
    mplsMark.EntityData.YangName = "mpls-mark"
    mplsMark.EntityData.BundleName = "cisco_ios_xr"
    mplsMark.EntityData.ParentYangName = "class"
    mplsMark.EntityData.SegmentPath = "mpls-mark"
    mplsMark.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsMark.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsMark.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsMark.EntityData.Children = make(map[string]types.YChild)
    mplsMark.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsMark.EntityData.Leafs["mark-type"] = types.YLeaf{"MarkType", mplsMark.MarkType}
    mplsMark.EntityData.Leafs["mark-value"] = types.YLeaf{"MarkValue", mplsMark.MarkValue}
    return &(mplsMark.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred
// WRED parameters
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred struct {
    EntityData types.CommonEntityData
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
    WredMatchValue PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue

    // Configured minimum threshold.
    ConfigMinThreshold PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMinThreshold

    // Configured maximum threshold.
    ConfigMaxThreshold PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMaxThreshold
}

func (wred *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "class"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Children["wred-match-value"] = types.YChild{"WredMatchValue", &wred.WredMatchValue}
    wred.EntityData.Children["config-min-threshold"] = types.YChild{"ConfigMinThreshold", &wred.ConfigMinThreshold}
    wred.EntityData.Children["config-max-threshold"] = types.YChild{"ConfigMaxThreshold", &wred.ConfigMaxThreshold}
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["wred-match-type"] = types.YLeaf{"WredMatchType", wred.WredMatchType}
    wred.EntityData.Leafs["hardware-min-threshold-bytes"] = types.YLeaf{"HardwareMinThresholdBytes", wred.HardwareMinThresholdBytes}
    wred.EntityData.Leafs["hardware-max-threshold-bytes"] = types.YLeaf{"HardwareMaxThresholdBytes", wred.HardwareMaxThresholdBytes}
    wred.EntityData.Leafs["first-segment"] = types.YLeaf{"FirstSegment", wred.FirstSegment}
    wred.EntityData.Leafs["segment-size"] = types.YLeaf{"SegmentSize", wred.SegmentSize}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue
// WRED match values
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dnx qosea show red match value. The type is slice of
    // PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue.
    DnxQoseaShowRedMatchValue []PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
}

func (wredMatchValue *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue) GetEntityData() *types.CommonEntityData {
    wredMatchValue.EntityData.YFilter = wredMatchValue.YFilter
    wredMatchValue.EntityData.YangName = "wred-match-value"
    wredMatchValue.EntityData.BundleName = "cisco_ios_xr"
    wredMatchValue.EntityData.ParentYangName = "wred"
    wredMatchValue.EntityData.SegmentPath = "wred-match-value"
    wredMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wredMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wredMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wredMatchValue.EntityData.Children = make(map[string]types.YChild)
    wredMatchValue.EntityData.Children["dnx-qosea-show-red-match-value"] = types.YChild{"DnxQoseaShowRedMatchValue", nil}
    for i := range wredMatchValue.DnxQoseaShowRedMatchValue {
        wredMatchValue.EntityData.Children[types.GetSegmentPath(&wredMatchValue.DnxQoseaShowRedMatchValue[i])] = types.YChild{"DnxQoseaShowRedMatchValue", &wredMatchValue.DnxQoseaShowRedMatchValue[i]}
    }
    wredMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(wredMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue
// dnx qosea show red match value
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Start value of a range. The type is interface{} with range: 0..255.
    RangeStart interface{}

    // End value of a range. The type is interface{} with range: 0..255.
    RangeEnd interface{}
}

func (dnxQoseaShowRedMatchValue *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_WredMatchValue_DnxQoseaShowRedMatchValue) GetEntityData() *types.CommonEntityData {
    dnxQoseaShowRedMatchValue.EntityData.YFilter = dnxQoseaShowRedMatchValue.YFilter
    dnxQoseaShowRedMatchValue.EntityData.YangName = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.BundleName = "cisco_ios_xr"
    dnxQoseaShowRedMatchValue.EntityData.ParentYangName = "wred-match-value"
    dnxQoseaShowRedMatchValue.EntityData.SegmentPath = "dnx-qosea-show-red-match-value"
    dnxQoseaShowRedMatchValue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnxQoseaShowRedMatchValue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnxQoseaShowRedMatchValue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnxQoseaShowRedMatchValue.EntityData.Children = make(map[string]types.YChild)
    dnxQoseaShowRedMatchValue.EntityData.Leafs = make(map[string]types.YLeaf)
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-start"] = types.YLeaf{"RangeStart", dnxQoseaShowRedMatchValue.RangeStart}
    dnxQoseaShowRedMatchValue.EntityData.Leafs["range-end"] = types.YLeaf{"RangeEnd", dnxQoseaShowRedMatchValue.RangeEnd}
    return &(dnxQoseaShowRedMatchValue.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMinThreshold
// Configured minimum threshold
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMinThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMinThreshold *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMinThreshold) GetEntityData() *types.CommonEntityData {
    configMinThreshold.EntityData.YFilter = configMinThreshold.YFilter
    configMinThreshold.EntityData.YangName = "config-min-threshold"
    configMinThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMinThreshold.EntityData.ParentYangName = "wred"
    configMinThreshold.EntityData.SegmentPath = "config-min-threshold"
    configMinThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMinThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMinThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMinThreshold.EntityData.Children = make(map[string]types.YChild)
    configMinThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMinThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMinThreshold.PolicyValue}
    configMinThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMinThreshold.PolicyUnit}
    return &(configMinThreshold.EntityData)
}

// PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMaxThreshold
// Configured maximum threshold
type PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMaxThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy value. The type is interface{} with range: 0..4294967295.
    PolicyValue interface{}

    // Policy unit. The type is PolicyParamUnit.
    PolicyUnit interface{}
}

func (configMaxThreshold *PlatformQos_Nodes_Node_BundleInterfaceSingles_BundleInterfaceSingle_Classes_Class_Wred_ConfigMaxThreshold) GetEntityData() *types.CommonEntityData {
    configMaxThreshold.EntityData.YFilter = configMaxThreshold.YFilter
    configMaxThreshold.EntityData.YangName = "config-max-threshold"
    configMaxThreshold.EntityData.BundleName = "cisco_ios_xr"
    configMaxThreshold.EntityData.ParentYangName = "wred"
    configMaxThreshold.EntityData.SegmentPath = "config-max-threshold"
    configMaxThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMaxThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMaxThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMaxThreshold.EntityData.Children = make(map[string]types.YChild)
    configMaxThreshold.EntityData.Leafs = make(map[string]types.YLeaf)
    configMaxThreshold.EntityData.Leafs["policy-value"] = types.YLeaf{"PolicyValue", configMaxThreshold.PolicyValue}
    configMaxThreshold.EntityData.Leafs["policy-unit"] = types.YLeaf{"PolicyUnit", configMaxThreshold.PolicyUnit}
    return &(configMaxThreshold.EntityData)
}

// PlatformQos_Nodes_Node_RemoteInterfaces
// QoS list of remote interfaces
type PlatformQos_Nodes_Node_RemoteInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // QoS remote interface names. The type is slice of
    // PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface.
    RemoteInterface []PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface
}

func (remoteInterfaces *PlatformQos_Nodes_Node_RemoteInterfaces) GetEntityData() *types.CommonEntityData {
    remoteInterfaces.EntityData.YFilter = remoteInterfaces.YFilter
    remoteInterfaces.EntityData.YangName = "remote-interfaces"
    remoteInterfaces.EntityData.BundleName = "cisco_ios_xr"
    remoteInterfaces.EntityData.ParentYangName = "node"
    remoteInterfaces.EntityData.SegmentPath = "remote-interfaces"
    remoteInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteInterfaces.EntityData.Children = make(map[string]types.YChild)
    remoteInterfaces.EntityData.Children["remote-interface"] = types.YChild{"RemoteInterface", nil}
    for i := range remoteInterfaces.RemoteInterface {
        remoteInterfaces.EntityData.Children[types.GetSegmentPath(&remoteInterfaces.RemoteInterface[i])] = types.YChild{"RemoteInterface", &remoteInterfaces.RemoteInterface[i]}
    }
    remoteInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(remoteInterfaces.EntityData)
}

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface
// QoS remote interface names
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the remote interface. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
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

func (remoteInterface *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface) GetEntityData() *types.CommonEntityData {
    remoteInterface.EntityData.YFilter = remoteInterface.YFilter
    remoteInterface.EntityData.YangName = "remote-interface"
    remoteInterface.EntityData.BundleName = "cisco_ios_xr"
    remoteInterface.EntityData.ParentYangName = "remote-interfaces"
    remoteInterface.EntityData.SegmentPath = "remote-interface" + "[interface-name='" + fmt.Sprintf("%v", remoteInterface.InterfaceName) + "']"
    remoteInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteInterface.EntityData.Children = make(map[string]types.YChild)
    remoteInterface.EntityData.Children["remote-class"] = types.YChild{"RemoteClass", nil}
    for i := range remoteInterface.RemoteClass {
        remoteInterface.EntityData.Children[types.GetSegmentPath(&remoteInterface.RemoteClass[i])] = types.YChild{"RemoteClass", &remoteInterface.RemoteClass[i]}
    }
    remoteInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", remoteInterface.InterfaceName}
    remoteInterface.EntityData.Leafs["policy-name"] = types.YLeaf{"PolicyName", remoteInterface.PolicyName}
    remoteInterface.EntityData.Leafs["virtual-output-queue-statistics-handle"] = types.YLeaf{"VirtualOutputQueueStatisticsHandle", remoteInterface.VirtualOutputQueueStatisticsHandle}
    remoteInterface.EntityData.Leafs["interface-handle"] = types.YLeaf{"InterfaceHandle", remoteInterface.InterfaceHandle}
    remoteInterface.EntityData.Leafs["number-of-virtual-output-queues"] = types.YLeaf{"NumberOfVirtualOutputQueues", remoteInterface.NumberOfVirtualOutputQueues}
    remoteInterface.EntityData.Leafs["number-of-classes"] = types.YLeaf{"NumberOfClasses", remoteInterface.NumberOfClasses}
    return &(remoteInterface.EntityData)
}

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass
// Remote Class array
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass struct {
    EntityData types.CommonEntityData
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

func (remoteClass *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass) GetEntityData() *types.CommonEntityData {
    remoteClass.EntityData.YFilter = remoteClass.YFilter
    remoteClass.EntityData.YangName = "remote-class"
    remoteClass.EntityData.BundleName = "cisco_ios_xr"
    remoteClass.EntityData.ParentYangName = "remote-interface"
    remoteClass.EntityData.SegmentPath = "remote-class"
    remoteClass.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteClass.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteClass.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteClass.EntityData.Children = make(map[string]types.YChild)
    remoteClass.EntityData.Children["wred"] = types.YChild{"Wred", nil}
    for i := range remoteClass.Wred {
        remoteClass.EntityData.Children[types.GetSegmentPath(&remoteClass.Wred[i])] = types.YChild{"Wred", &remoteClass.Wred[i]}
    }
    remoteClass.EntityData.Children["hw-wred"] = types.YChild{"HwWred", nil}
    for i := range remoteClass.HwWred {
        remoteClass.EntityData.Children[types.GetSegmentPath(&remoteClass.HwWred[i])] = types.YChild{"HwWred", &remoteClass.HwWred[i]}
    }
    remoteClass.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteClass.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", remoteClass.ClassName}
    remoteClass.EntityData.Leafs["class-id"] = types.YLeaf{"ClassId", remoteClass.ClassId}
    remoteClass.EntityData.Leafs["cos-q"] = types.YLeaf{"CosQ", remoteClass.CosQ}
    remoteClass.EntityData.Leafs["queue-limit"] = types.YLeaf{"QueueLimit", remoteClass.QueueLimit}
    remoteClass.EntityData.Leafs["hardware-queue-limit"] = types.YLeaf{"HardwareQueueLimit", remoteClass.HardwareQueueLimit}
    return &(remoteClass.EntityData)
}

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred
// Default/Configured WRED profiles
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum Threshold. The type is interface{} with range: 0..4294967295.
    MinThreshold interface{}

    // Maximum Threshold. The type is interface{} with range: 0..4294967295.
    MaxThreshold interface{}

    // Drop Probability. The type is interface{} with range: 0..4294967295.
    DropProbability interface{}
}

func (wred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_Wred) GetEntityData() *types.CommonEntityData {
    wred.EntityData.YFilter = wred.YFilter
    wred.EntityData.YangName = "wred"
    wred.EntityData.BundleName = "cisco_ios_xr"
    wred.EntityData.ParentYangName = "remote-class"
    wred.EntityData.SegmentPath = "wred"
    wred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    wred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    wred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    wred.EntityData.Children = make(map[string]types.YChild)
    wred.EntityData.Leafs = make(map[string]types.YLeaf)
    wred.EntityData.Leafs["min-threshold"] = types.YLeaf{"MinThreshold", wred.MinThreshold}
    wred.EntityData.Leafs["max-threshold"] = types.YLeaf{"MaxThreshold", wred.MaxThreshold}
    wred.EntityData.Leafs["drop-probability"] = types.YLeaf{"DropProbability", wred.DropProbability}
    return &(wred.EntityData)
}

// PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred
// Hardware WRED profiles
type PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum Threshold. The type is interface{} with range: 0..4294967295.
    MinThreshold interface{}

    // Maximum Threshold. The type is interface{} with range: 0..4294967295.
    MaxThreshold interface{}

    // Drop Probability. The type is interface{} with range: 0..4294967295.
    DropProbability interface{}
}

func (hwWred *PlatformQos_Nodes_Node_RemoteInterfaces_RemoteInterface_RemoteClass_HwWred) GetEntityData() *types.CommonEntityData {
    hwWred.EntityData.YFilter = hwWred.YFilter
    hwWred.EntityData.YangName = "hw-wred"
    hwWred.EntityData.BundleName = "cisco_ios_xr"
    hwWred.EntityData.ParentYangName = "remote-class"
    hwWred.EntityData.SegmentPath = "hw-wred"
    hwWred.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hwWred.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hwWred.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hwWred.EntityData.Children = make(map[string]types.YChild)
    hwWred.EntityData.Leafs = make(map[string]types.YLeaf)
    hwWred.EntityData.Leafs["min-threshold"] = types.YLeaf{"MinThreshold", hwWred.MinThreshold}
    hwWred.EntityData.Leafs["max-threshold"] = types.YLeaf{"MaxThreshold", hwWred.MaxThreshold}
    hwWred.EntityData.Leafs["drop-probability"] = types.YLeaf{"DropProbability", hwWred.DropProbability}
    return &(hwWred.EntityData)
}

