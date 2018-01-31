// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-driver package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fia: FIA driver operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package dnx_driver_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package dnx_driver_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-dnx-driver-oper fia}", reflect.TypeOf(Fia{}))
    ydk.RegisterEntity("Cisco-IOS-XR-dnx-driver-oper:fia", reflect.TypeOf(Fia{}))
}

// SliceState represents Slice state
type SliceState string

const (
    // slice oper unset
    SliceState_slice_oper_unset SliceState = "slice-oper-unset"

    // slice oper down
    SliceState_slice_oper_down SliceState = "slice-oper-down"

    // slice oper up
    SliceState_slice_oper_up SliceState = "slice-oper-up"

    // slice oper na
    SliceState_slice_oper_na SliceState = "slice-oper-na"
)

// AsicAccessState represents Asic access state
type AsicAccessState string

const (
    // asic state unset
    AsicAccessState_asic_state_unset AsicAccessState = "asic-state-unset"

    // asic state none
    AsicAccessState_asic_state_none AsicAccessState = "asic-state-none"

    // asic state device off line
    AsicAccessState_asic_state_device_off_line AsicAccessState = "asic-state-device-off-line"

    // asic state device created
    AsicAccessState_asic_state_device_created AsicAccessState = "asic-state-device-created"

    // asic state device online
    AsicAccessState_asic_state_device_online AsicAccessState = "asic-state-device-online"

    // asic state warmboot
    AsicAccessState_asic_state_warmboot AsicAccessState = "asic-state-warmboot"

    // asic state de init start
    AsicAccessState_asic_state_de_init_start AsicAccessState = "asic-state-de-init-start"

    // asic state intr de init
    AsicAccessState_asic_state_intr_de_init AsicAccessState = "asic-state-intr-de-init"

    // asic state bcm detach
    AsicAccessState_asic_state_bcm_detach AsicAccessState = "asic-state-bcm-detach"

    // asic state soc de init
    AsicAccessState_asic_state_soc_de_init AsicAccessState = "asic-state-soc-de-init"

    // asic state de init done
    AsicAccessState_asic_state_de_init_done AsicAccessState = "asic-state-de-init-done"

    // asic state soc init
    AsicAccessState_asic_state_soc_init AsicAccessState = "asic-state-soc-init"

    // asic state bcm init
    AsicAccessState_asic_state_bcm_init AsicAccessState = "asic-state-bcm-init"

    // asic state intr init
    AsicAccessState_asic_state_intr_init AsicAccessState = "asic-state-intr-init"

    // asic state soc init start
    AsicAccessState_asic_state_soc_init_start AsicAccessState = "asic-state-soc-init-start"

    // asic state bcm init start
    AsicAccessState_asic_state_bcm_init_start AsicAccessState = "asic-state-bcm-init-start"

    // asic state intr init start
    AsicAccessState_asic_state_intr_init_start AsicAccessState = "asic-state-intr-init-start"

    // asic state hard reset
    AsicAccessState_asic_state_hard_reset AsicAccessState = "asic-state-hard-reset"

    // asic state normal
    AsicAccessState_asic_state_normal AsicAccessState = "asic-state-normal"

    // asic state exception
    AsicAccessState_asic_state_exception AsicAccessState = "asic-state-exception"

    // asic state hp attached
    AsicAccessState_asic_state_hp_attached AsicAccessState = "asic-state-hp-attached"

    // asic state quiesce
    AsicAccessState_asic_state_quiesce AsicAccessState = "asic-state-quiesce"

    // asic state issu started
    AsicAccessState_asic_state_issu_started AsicAccessState = "asic-state-issu-started"

    // asic state issu started nn
    AsicAccessState_asic_state_issu_started_nn AsicAccessState = "asic-state-issu-started-nn"

    // asic state issu abort
    AsicAccessState_asic_state_issu_abort AsicAccessState = "asic-state-issu-abort"

    // asic state max
    AsicAccessState_asic_state_max AsicAccessState = "asic-state-max"
)

// LinkErrorState represents Link error state
type LinkErrorState string

const (
    // link error unset
    LinkErrorState_link_error_unset LinkErrorState = "link-error-unset"

    // link error none
    LinkErrorState_link_error_none LinkErrorState = "link-error-none"

    // link error shut
    LinkErrorState_link_error_shut LinkErrorState = "link-error-shut"

    // link error max
    LinkErrorState_link_error_max LinkErrorState = "link-error-max"
)

// FcMode represents Fc mode
type FcMode string

const (
    // fc mode unset
    FcMode_fc_mode_unset FcMode = "fc-mode-unset"

    // fc mode unavail
    FcMode_fc_mode_unavail FcMode = "fc-mode-unavail"

    // fc mode inband
    FcMode_fc_mode_inband FcMode = "fc-mode-inband"

    // fc mode oob
    FcMode_fc_mode_oob FcMode = "fc-mode-oob"
)

// Asic represents Asic
type Asic string

const (
    // asic unset
    Asic_asic_unset Asic = "asic-unset"

    // asic unavail
    Asic_asic_unavail Asic = "asic-unavail"

    // asic fia
    Asic_asic_fia Asic = "asic-fia"

    // asic s123
    Asic_asic_s123 Asic = "asic-s123"

    // asic s13
    Asic_asic_s13 Asic = "asic-s13"

    // asic s2
    Asic_asic_s2 Asic = "asic-s2"

    // asic b2b
    Asic_asic_b2b Asic = "asic-b2b"

    // asic type unknown
    Asic_asic_type_unknown Asic = "asic-type-unknown"
)

// AsicOperState represents Asic oper state
type AsicOperState string

const (
    // asic oper unset
    AsicOperState_asic_oper_unset AsicOperState = "asic-oper-unset"

    // asic oper unknown
    AsicOperState_asic_oper_unknown AsicOperState = "asic-oper-unknown"

    // asic oper up
    AsicOperState_asic_oper_up AsicOperState = "asic-oper-up"

    // asic oper down
    AsicOperState_asic_oper_down AsicOperState = "asic-oper-down"

    // asic card down
    AsicOperState_asic_card_down AsicOperState = "asic-card-down"
)

// Link represents Link
type Link string

const (
    // link type unset
    Link_link_type_unset Link = "link-type-unset"

    // link type unavail
    Link_link_type_unavail Link = "link-type-unavail"

    // link type tx
    Link_link_type_tx Link = "link-type-tx"

    // link type rx
    Link_link_type_rx Link = "link-type-rx"
)

// OperState represents Oper state
type OperState string

const (
    // oper unset
    OperState_oper_unset OperState = "oper-unset"

    // oper unknown
    OperState_oper_unknown OperState = "oper-unknown"

    // oper up
    OperState_oper_up OperState = "oper-up"

    // oper down
    OperState_oper_down OperState = "oper-down"

    // card down
    OperState_card_down OperState = "card-down"
)

// AsicInitMethod represents Asic init method
type AsicInitMethod string

const (
    // asic init method unset
    AsicInitMethod_asic_init_method_unset AsicInitMethod = "asic-init-method-unset"

    // asic init method no reset
    AsicInitMethod_asic_init_method_no_reset AsicInitMethod = "asic-init-method-no-reset"

    // asic init method pon reset
    AsicInitMethod_asic_init_method_pon_reset AsicInitMethod = "asic-init-method-pon-reset"

    // asic init method pon reset on intr
    AsicInitMethod_asic_init_method_pon_reset_on_intr AsicInitMethod = "asic-init-method-pon-reset-on-intr"

    // asic init method hard reset
    AsicInitMethod_asic_init_method_hard_reset AsicInitMethod = "asic-init-method-hard-reset"

    // asic init method warmboot
    AsicInitMethod_asic_init_method_warmboot AsicInitMethod = "asic-init-method-warmboot"

    // asic init method issu wb
    AsicInitMethod_asic_init_method_issu_wb AsicInitMethod = "asic-init-method-issu-wb"

    // asic init method pci shutdown
    AsicInitMethod_asic_init_method_pci_shutdown AsicInitMethod = "asic-init-method-pci-shutdown"

    // asic init method quiesce
    AsicInitMethod_asic_init_method_quiesce AsicInitMethod = "asic-init-method-quiesce"

    // asic init method issu started
    AsicInitMethod_asic_init_method_issu_started AsicInitMethod = "asic-init-method-issu-started"

    // asic init method issu rollback
    AsicInitMethod_asic_init_method_issu_rollback AsicInitMethod = "asic-init-method-issu-rollback"

    // asic init method issu abort
    AsicInitMethod_asic_init_method_issu_abort AsicInitMethod = "asic-init-method-issu-abort"

    // asic init method slice cleanup
    AsicInitMethod_asic_init_method_slice_cleanup AsicInitMethod = "asic-init-method-slice-cleanup"

    // asic init method lc remove
    AsicInitMethod_asic_init_method_lc_remove AsicInitMethod = "asic-init-method-lc-remove"

    // asic init method node down
    AsicInitMethod_asic_init_method_node_down AsicInitMethod = "asic-init-method-node-down"

    // asic init method intr
    AsicInitMethod_asic_init_method_intr AsicInitMethod = "asic-init-method-intr"

    // asic init method board reload
    AsicInitMethod_asic_init_method_board_reload AsicInitMethod = "asic-init-method-board-reload"

    // asic init method max
    AsicInitMethod_asic_init_method_max AsicInitMethod = "asic-init-method-max"
)

// AdminState represents Admin state
type AdminState string

const (
    // admin unset
    AdminState_admin_unset AdminState = "admin-unset"

    // admin up
    AdminState_admin_up AdminState = "admin-up"

    // admin down
    AdminState_admin_down AdminState = "admin-down"
)

// LinkStage represents Link stage
type LinkStage string

const (
    // link stage unset
    LinkStage_link_stage_unset LinkStage = "link-stage-unset"

    // link stage unused
    LinkStage_link_stage_unused LinkStage = "link-stage-unused"

    // link stage fia
    LinkStage_link_stage_fia LinkStage = "link-stage-fia"

    // link stage s1
    LinkStage_link_stage_s1 LinkStage = "link-stage-s1"

    // link stage s2
    LinkStage_link_stage_s2 LinkStage = "link-stage-s2"

    // link stage s3
    LinkStage_link_stage_s3 LinkStage = "link-stage-s3"

    // link stage unknown
    LinkStage_link_stage_unknown LinkStage = "link-stage-unknown"
)

// Rack represents Rack
type Rack string

const (
    // rack type unset
    Rack_rack_type_unset Rack = "rack-type-unset"

    // rack type lcc
    Rack_rack_type_lcc Rack = "rack-type-lcc"

    // rack type fcc
    Rack_rack_type_fcc Rack = "rack-type-fcc"
)

// Fia
// FIA driver operational data
type Fia struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FIA driver operational data for available nodes.
    Nodes Fia_Nodes
}

func (fia *Fia) GetFilter() yfilter.YFilter { return fia.YFilter }

func (fia *Fia) SetFilter(yf yfilter.YFilter) { fia.YFilter = yf }

func (fia *Fia) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (fia *Fia) GetSegmentPath() string {
    return "Cisco-IOS-XR-dnx-driver-oper:fia"
}

func (fia *Fia) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &fia.Nodes
    }
    return nil
}

func (fia *Fia) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &fia.Nodes
    return children
}

func (fia *Fia) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fia *Fia) GetBundleName() string { return "cisco_ios_xr" }

func (fia *Fia) GetYangName() string { return "fia" }

func (fia *Fia) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fia *Fia) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fia *Fia) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fia *Fia) SetParent(parent types.Entity) { fia.parent = parent }

func (fia *Fia) GetParent() types.Entity { return fia.parent }

func (fia *Fia) GetParentYangName() string { return "Cisco-IOS-XR-dnx-driver-oper" }

// Fia_Nodes
// FIA driver operational data for available nodes
type Fia_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FIA operational data for a particular node. The type is slice of
    // Fia_Nodes_Node.
    Node []Fia_Nodes_Node
}

func (nodes *Fia_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Fia_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Fia_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Fia_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Fia_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Fia_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Fia_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Fia_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Fia_Nodes) GetYangName() string { return "nodes" }

func (nodes *Fia_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Fia_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Fia_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Fia_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Fia_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Fia_Nodes) GetParentYangName() string { return "fia" }

// Fia_Nodes_Node
// FIA operational data for a particular node
type Fia_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // FIA link rx information.
    RxLinkInformation Fia_Nodes_Node_RxLinkInformation

    // FIA driver information.
    DriverInformation Fia_Nodes_Node_DriverInformation

    // Clear statistics information.
    ClearStatistics Fia_Nodes_Node_ClearStatistics

    // FIA link TX information.
    TxLinkInformation Fia_Nodes_Node_TxLinkInformation

    // FIA diag shell information.
    DiagShell Fia_Nodes_Node_DiagShell

    // FIA operational data of oir history.
    OirHistory Fia_Nodes_Node_OirHistory

    // FIA asic statistics information.
    AsicStatistics Fia_Nodes_Node_AsicStatistics
}

func (node *Fia_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Fia_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Fia_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "rx-link-information" { return "RxLinkInformation" }
    if yname == "driver-information" { return "DriverInformation" }
    if yname == "clear-statistics" { return "ClearStatistics" }
    if yname == "tx-link-information" { return "TxLinkInformation" }
    if yname == "diag-shell" { return "DiagShell" }
    if yname == "oir-history" { return "OirHistory" }
    if yname == "asic-statistics" { return "AsicStatistics" }
    return ""
}

func (node *Fia_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Fia_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-link-information" {
        return &node.RxLinkInformation
    }
    if childYangName == "driver-information" {
        return &node.DriverInformation
    }
    if childYangName == "clear-statistics" {
        return &node.ClearStatistics
    }
    if childYangName == "tx-link-information" {
        return &node.TxLinkInformation
    }
    if childYangName == "diag-shell" {
        return &node.DiagShell
    }
    if childYangName == "oir-history" {
        return &node.OirHistory
    }
    if childYangName == "asic-statistics" {
        return &node.AsicStatistics
    }
    return nil
}

func (node *Fia_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rx-link-information"] = &node.RxLinkInformation
    children["driver-information"] = &node.DriverInformation
    children["clear-statistics"] = &node.ClearStatistics
    children["tx-link-information"] = &node.TxLinkInformation
    children["diag-shell"] = &node.DiagShell
    children["oir-history"] = &node.OirHistory
    children["asic-statistics"] = &node.AsicStatistics
    return children
}

func (node *Fia_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Fia_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Fia_Nodes_Node) GetYangName() string { return "node" }

func (node *Fia_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Fia_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Fia_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Fia_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Fia_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Fia_Nodes_Node) GetParentYangName() string { return "nodes" }

// Fia_Nodes_Node_RxLinkInformation
// FIA link rx information
type Fia_Nodes_Node_RxLinkInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Option table for link rx information.
    LinkOptions Fia_Nodes_Node_RxLinkInformation_LinkOptions
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetFilter() yfilter.YFilter { return rxLinkInformation.YFilter }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) SetFilter(yf yfilter.YFilter) { rxLinkInformation.YFilter = yf }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetGoName(yname string) string {
    if yname == "link-options" { return "LinkOptions" }
    return ""
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetSegmentPath() string {
    return "rx-link-information"
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-options" {
        return &rxLinkInformation.LinkOptions
    }
    return nil
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-options"] = &rxLinkInformation.LinkOptions
    return children
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetBundleName() string { return "cisco_ios_xr" }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetYangName() string { return "rx-link-information" }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) SetParent(parent types.Entity) { rxLinkInformation.parent = parent }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetParent() types.Entity { return rxLinkInformation.parent }

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions
// Option table for link rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Option : topo , flag , stats. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption.
    LinkOption []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetFilter() yfilter.YFilter { return linkOptions.YFilter }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) SetFilter(yf yfilter.YFilter) { linkOptions.YFilter = yf }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetGoName(yname string) string {
    if yname == "link-option" { return "LinkOption" }
    return ""
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetSegmentPath() string {
    return "link-options"
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-option" {
        for _, c := range linkOptions.LinkOption {
            if linkOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption{}
        linkOptions.LinkOption = append(linkOptions.LinkOption, child)
        return &linkOptions.LinkOption[len(linkOptions.LinkOption)-1]
    }
    return nil
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkOptions.LinkOption {
        children[linkOptions.LinkOption[i].GetSegmentPath()] = &linkOptions.LinkOption[i]
    }
    return children
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetBundleName() string { return "cisco_ios_xr" }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetYangName() string { return "link-options" }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) SetParent(parent types.Entity) { linkOptions.parent = parent }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetParent() types.Entity { return linkOptions.parent }

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetParentYangName() string { return "rx-link-information" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption
// Option : topo , flag , stats
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Link option. The type is string with pattern:
    // (flap)|(topo).
    Option interface{}

    // Instance table for rx information.
    RxAsicInstances Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetFilter() yfilter.YFilter { return linkOption.YFilter }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) SetFilter(yf yfilter.YFilter) { linkOption.YFilter = yf }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetGoName(yname string) string {
    if yname == "option" { return "Option" }
    if yname == "rx-asic-instances" { return "RxAsicInstances" }
    return ""
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetSegmentPath() string {
    return "link-option" + "[option='" + fmt.Sprintf("%v", linkOption.Option) + "']"
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-asic-instances" {
        return &linkOption.RxAsicInstances
    }
    return nil
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rx-asic-instances"] = &linkOption.RxAsicInstances
    return children
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["option"] = linkOption.Option
    return leafs
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetBundleName() string { return "cisco_ios_xr" }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetYangName() string { return "link-option" }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) SetParent(parent types.Entity) { linkOption.parent = parent }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetParent() types.Entity { return linkOption.parent }

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetParentYangName() string { return "link-options" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances
// Instance table for rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance number for rx link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance.
    RxAsicInstance []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetFilter() yfilter.YFilter { return rxAsicInstances.YFilter }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) SetFilter(yf yfilter.YFilter) { rxAsicInstances.YFilter = yf }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetGoName(yname string) string {
    if yname == "rx-asic-instance" { return "RxAsicInstance" }
    return ""
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetSegmentPath() string {
    return "rx-asic-instances"
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-asic-instance" {
        for _, c := range rxAsicInstances.RxAsicInstance {
            if rxAsicInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance{}
        rxAsicInstances.RxAsicInstance = append(rxAsicInstances.RxAsicInstance, child)
        return &rxAsicInstances.RxAsicInstance[len(rxAsicInstances.RxAsicInstance)-1]
    }
    return nil
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rxAsicInstances.RxAsicInstance {
        children[rxAsicInstances.RxAsicInstance[i].GetSegmentPath()] = &rxAsicInstances.RxAsicInstance[i]
    }
    return children
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetBundleName() string { return "cisco_ios_xr" }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetYangName() string { return "rx-asic-instances" }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) SetParent(parent types.Entity) { rxAsicInstances.parent = parent }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetParent() types.Entity { return rxAsicInstances.parent }

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetParentYangName() string { return "link-option" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance
// Instance number for rx link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Receive instance. The type is interface{} with
    // range: 0..255.
    Instance interface{}

    // Link table class for rx information.
    RxLinks Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetFilter() yfilter.YFilter { return rxAsicInstance.YFilter }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) SetFilter(yf yfilter.YFilter) { rxAsicInstance.YFilter = yf }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "rx-links" { return "RxLinks" }
    return ""
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetSegmentPath() string {
    return "rx-asic-instance" + "[instance='" + fmt.Sprintf("%v", rxAsicInstance.Instance) + "']"
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-links" {
        return &rxAsicInstance.RxLinks
    }
    return nil
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rx-links"] = &rxAsicInstance.RxLinks
    return children
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance"] = rxAsicInstance.Instance
    return leafs
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetBundleName() string { return "cisco_ios_xr" }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetYangName() string { return "rx-asic-instance" }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) SetParent(parent types.Entity) { rxAsicInstance.parent = parent }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetParent() types.Entity { return rxAsicInstance.parent }

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetParentYangName() string { return "rx-asic-instances" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks
// Link table class for rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link number for rx link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink.
    RxLink []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetFilter() yfilter.YFilter { return rxLinks.YFilter }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) SetFilter(yf yfilter.YFilter) { rxLinks.YFilter = yf }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetGoName(yname string) string {
    if yname == "rx-link" { return "RxLink" }
    return ""
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetSegmentPath() string {
    return "rx-links"
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-link" {
        for _, c := range rxLinks.RxLink {
            if rxLinks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink{}
        rxLinks.RxLink = append(rxLinks.RxLink, child)
        return &rxLinks.RxLink[len(rxLinks.RxLink)-1]
    }
    return nil
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rxLinks.RxLink {
        children[rxLinks.RxLink[i].GetSegmentPath()] = &rxLinks.RxLink[i]
    }
    return children
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetBundleName() string { return "cisco_ios_xr" }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetYangName() string { return "rx-links" }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) SetParent(parent types.Entity) { rxLinks.parent = parent }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetParent() types.Entity { return rxLinks.parent }

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetParentYangName() string { return "rx-asic-instance" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink
// Link number for rx link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start number. The type is interface{} with range: 0..47.
    StartNumber interface{}

    // End number. The type is interface{} with range: 0..47.
    EndNumber interface{}

    // RX link status option. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    StatusOption interface{}

    // Single link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink.
    RxLink []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetFilter() yfilter.YFilter { return rxLink.YFilter }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) SetFilter(yf yfilter.YFilter) { rxLink.YFilter = yf }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetGoName(yname string) string {
    if yname == "start-number" { return "StartNumber" }
    if yname == "end-number" { return "EndNumber" }
    if yname == "status-option" { return "StatusOption" }
    if yname == "rx-link" { return "RxLink" }
    return ""
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetSegmentPath() string {
    return "rx-link"
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rx-link" {
        for _, c := range rxLink.RxLink {
            if rxLink.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink{}
        rxLink.RxLink = append(rxLink.RxLink, child)
        return &rxLink.RxLink[len(rxLink.RxLink)-1]
    }
    return nil
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rxLink.RxLink {
        children[rxLink.RxLink[i].GetSegmentPath()] = &rxLink.RxLink[i]
    }
    return children
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-number"] = rxLink.StartNumber
    leafs["end-number"] = rxLink.EndNumber
    leafs["status-option"] = rxLink.StatusOption
    return leafs
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetBundleName() string { return "cisco_ios_xr" }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetYangName() string { return "rx-link" }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) SetParent(parent types.Entity) { rxLink.parent = parent }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetParent() types.Entity { return rxLink.parent }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetParentYangName() string { return "rx-links" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink
// Single link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single link. The type is interface{} with range:
    // -2147483648..2147483647.
    Link interface{}

    // speed. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // Stage. The type is LinkStage.
    Stage interface{}

    // is link valid. The type is bool.
    IsLinkValid interface{}

    // is conf pending. The type is bool.
    IsConfPending interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // flags. The type is string.
    Flags interface{}

    // flap cnt. The type is interface{} with range: 0..4294967295.
    FlapCnt interface{}

    // num admin shuts. The type is interface{} with range: 0..4294967295.
    NumAdminShuts interface{}

    // correctable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    CorrectableErrors interface{}

    // uncorrectable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    UncorrectableErrors interface{}

    // this link.
    ThisLink Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink

    // far end link.
    FarEndLink Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink

    // far end link in hw.
    FarEndLinkInHw Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw

    // history.
    History Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetFilter() yfilter.YFilter { return rxLink.YFilter }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) SetFilter(yf yfilter.YFilter) { rxLink.YFilter = yf }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetGoName(yname string) string {
    if yname == "link" { return "Link" }
    if yname == "speed" { return "Speed" }
    if yname == "stage" { return "Stage" }
    if yname == "is-link-valid" { return "IsLinkValid" }
    if yname == "is-conf-pending" { return "IsConfPending" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state" { return "OperState" }
    if yname == "error-state" { return "ErrorState" }
    if yname == "flags" { return "Flags" }
    if yname == "flap-cnt" { return "FlapCnt" }
    if yname == "num-admin-shuts" { return "NumAdminShuts" }
    if yname == "correctable-errors" { return "CorrectableErrors" }
    if yname == "uncorrectable-errors" { return "UncorrectableErrors" }
    if yname == "this-link" { return "ThisLink" }
    if yname == "far-end-link" { return "FarEndLink" }
    if yname == "far-end-link-in-hw" { return "FarEndLinkInHw" }
    if yname == "history" { return "History" }
    return ""
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetSegmentPath() string {
    return "rx-link" + "[link='" + fmt.Sprintf("%v", rxLink.Link) + "']"
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "this-link" {
        return &rxLink.ThisLink
    }
    if childYangName == "far-end-link" {
        return &rxLink.FarEndLink
    }
    if childYangName == "far-end-link-in-hw" {
        return &rxLink.FarEndLinkInHw
    }
    if childYangName == "history" {
        return &rxLink.History
    }
    return nil
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["this-link"] = &rxLink.ThisLink
    children["far-end-link"] = &rxLink.FarEndLink
    children["far-end-link-in-hw"] = &rxLink.FarEndLinkInHw
    children["history"] = &rxLink.History
    return children
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link"] = rxLink.Link
    leafs["speed"] = rxLink.Speed
    leafs["stage"] = rxLink.Stage
    leafs["is-link-valid"] = rxLink.IsLinkValid
    leafs["is-conf-pending"] = rxLink.IsConfPending
    leafs["admin-state"] = rxLink.AdminState
    leafs["oper-state"] = rxLink.OperState
    leafs["error-state"] = rxLink.ErrorState
    leafs["flags"] = rxLink.Flags
    leafs["flap-cnt"] = rxLink.FlapCnt
    leafs["num-admin-shuts"] = rxLink.NumAdminShuts
    leafs["correctable-errors"] = rxLink.CorrectableErrors
    leafs["uncorrectable-errors"] = rxLink.UncorrectableErrors
    return leafs
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetBundleName() string { return "cisco_ios_xr" }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetYangName() string { return "rx-link" }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) SetParent(parent types.Entity) { rxLink.parent = parent }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetParent() types.Entity { return rxLink.parent }

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetParentYangName() string { return "rx-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink
// this link
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetFilter() yfilter.YFilter { return thisLink.YFilter }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) SetFilter(yf yfilter.YFilter) { thisLink.YFilter = yf }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-stage" { return "LinkStage" }
    if yname == "link-num" { return "LinkNum" }
    if yname == "phy-link-num" { return "PhyLinkNum" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetSegmentPath() string {
    return "this-link"
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &thisLink.AsicId
    }
    return nil
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &thisLink.AsicId
    return children
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = thisLink.LinkType
    leafs["link-stage"] = thisLink.LinkStage
    leafs["link-num"] = thisLink.LinkNum
    leafs["phy-link-num"] = thisLink.PhyLinkNum
    return leafs
}

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetBundleName() string { return "cisco_ios_xr" }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetYangName() string { return "this-link" }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) SetParent(parent types.Entity) { thisLink.parent = parent }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetParent() types.Entity { return thisLink.parent }

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetParentYangName() string { return "rx-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetParentYangName() string { return "this-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink
// far end link
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetFilter() yfilter.YFilter { return farEndLink.YFilter }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) SetFilter(yf yfilter.YFilter) { farEndLink.YFilter = yf }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-stage" { return "LinkStage" }
    if yname == "link-num" { return "LinkNum" }
    if yname == "phy-link-num" { return "PhyLinkNum" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetSegmentPath() string {
    return "far-end-link"
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &farEndLink.AsicId
    }
    return nil
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &farEndLink.AsicId
    return children
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = farEndLink.LinkType
    leafs["link-stage"] = farEndLink.LinkStage
    leafs["link-num"] = farEndLink.LinkNum
    leafs["phy-link-num"] = farEndLink.PhyLinkNum
    return leafs
}

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetBundleName() string { return "cisco_ios_xr" }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetYangName() string { return "far-end-link" }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) SetParent(parent types.Entity) { farEndLink.parent = parent }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetParent() types.Entity { return farEndLink.parent }

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetParentYangName() string { return "rx-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetParentYangName() string { return "far-end-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw
// far end link in hw
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetFilter() yfilter.YFilter { return farEndLinkInHw.YFilter }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) SetFilter(yf yfilter.YFilter) { farEndLinkInHw.YFilter = yf }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-stage" { return "LinkStage" }
    if yname == "link-num" { return "LinkNum" }
    if yname == "phy-link-num" { return "PhyLinkNum" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetSegmentPath() string {
    return "far-end-link-in-hw"
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &farEndLinkInHw.AsicId
    }
    return nil
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &farEndLinkInHw.AsicId
    return children
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = farEndLinkInHw.LinkType
    leafs["link-stage"] = farEndLinkInHw.LinkStage
    leafs["link-num"] = farEndLinkInHw.LinkNum
    leafs["phy-link-num"] = farEndLinkInHw.PhyLinkNum
    return leafs
}

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetBundleName() string { return "cisco_ios_xr" }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetYangName() string { return "far-end-link-in-hw" }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) SetParent(parent types.Entity) { farEndLinkInHw.parent = parent }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetParent() types.Entity { return farEndLinkInHw.parent }

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetParentYangName() string { return "rx-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetParentYangName() string { return "far-end-link-in-hw" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History
// history
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // histnum. The type is interface{} with range: 0..255.
    Histnum interface{}

    // start index. The type is interface{} with range: 0..255.
    StartIndex interface{}

    // hist. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist.
    Hist []Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetGoName(yname string) string {
    if yname == "histnum" { return "Histnum" }
    if yname == "start-index" { return "StartIndex" }
    if yname == "hist" { return "Hist" }
    return ""
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetSegmentPath() string {
    return "history"
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hist" {
        for _, c := range history.Hist {
            if history.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist{}
        history.Hist = append(history.Hist, child)
        return &history.Hist[len(history.Hist)-1]
    }
    return nil
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range history.Hist {
        children[history.Hist[i].GetSegmentPath()] = &history.Hist[i]
    }
    return children
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["histnum"] = history.Histnum
    leafs["start-index"] = history.StartIndex
    return leafs
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetYangName() string { return "history" }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetParent() types.Entity { return history.parent }

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetParentYangName() string { return "rx-link" }

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist
// hist
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // timestamp. The type is interface{} with range: 0..18446744073709551615.
    Timestamp interface{}

    // reasons. The type is string.
    Reasons interface{}
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetFilter() yfilter.YFilter { return hist.YFilter }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) SetFilter(yf yfilter.YFilter) { hist.YFilter = yf }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetGoName(yname string) string {
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state" { return "OperState" }
    if yname == "error-state" { return "ErrorState" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "reasons" { return "Reasons" }
    return ""
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetSegmentPath() string {
    return "hist"
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-state"] = hist.AdminState
    leafs["oper-state"] = hist.OperState
    leafs["error-state"] = hist.ErrorState
    leafs["timestamp"] = hist.Timestamp
    leafs["reasons"] = hist.Reasons
    return leafs
}

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetBundleName() string { return "cisco_ios_xr" }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetYangName() string { return "hist" }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) SetParent(parent types.Entity) { hist.parent = parent }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetParent() types.Entity { return hist.parent }

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetParentYangName() string { return "history" }

// Fia_Nodes_Node_DriverInformation
// FIA driver information
type Fia_Nodes_Node_DriverInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // drv version. The type is interface{} with range: 0..4294967295.
    DrvVersion interface{}

    // coeff major rev. The type is interface{} with range: 0..4294967295.
    CoeffMajorRev interface{}

    // coeff minor rev. The type is interface{} with range: 0..4294967295.
    CoeffMinorRev interface{}

    // functional role. The type is interface{} with range: 0..255.
    FunctionalRole interface{}

    // issu role. The type is interface{} with range: 0..255.
    IssuRole interface{}

    // node id. The type is string.
    NodeId interface{}

    // rack type. The type is interface{} with range: -2147483648..2147483647.
    RackType interface{}

    // rack num. The type is interface{} with range: 0..255.
    RackNum interface{}

    // is driver ready. The type is bool.
    IsDriverReady interface{}

    // card avail mask. The type is interface{} with range: 0..4294967295.
    CardAvailMask interface{}

    // asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicAvailMask interface{}

    // exp asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    ExpAsicAvailMask interface{}

    // ucmc ratio. The type is interface{} with range: 0..4294967295.
    UcmcRatio interface{}

    // asic oper notify to fsdb pending bmap. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicOperNotifyToFsdbPendingBmap interface{}

    // is full fgid download req. The type is bool.
    IsFullFgidDownloadReq interface{}

    // is fgid download in progress. The type is bool.
    IsFgidDownloadInProgress interface{}

    // is fgid download completed. The type is bool.
    IsFgidDownloadCompleted interface{}

    // fsdb conn active. The type is bool.
    FsdbConnActive interface{}

    // fgid conn active. The type is bool.
    FgidConnActive interface{}

    // issu mgr conn active. The type is bool.
    IssuMgrConnActive interface{}

    // fsdb reg active. The type is bool.
    FsdbRegActive interface{}

    // fgid reg active. The type is bool.
    FgidRegActive interface{}

    // issu mgr reg active. The type is bool.
    IssuMgrRegActive interface{}

    // num pm conn reqs. The type is interface{} with range: 0..255.
    NumPmConnReqs interface{}

    // num fsdb conn reqs. The type is interface{} with range: 0..255.
    NumFsdbConnReqs interface{}

    // num fgid conn reqs. The type is interface{} with range: 0..255.
    NumFgidConnReqs interface{}

    // num fstats conn reqs. The type is interface{} with range: 0..255.
    NumFstatsConnReqs interface{}

    // num cm conn reqs. The type is interface{} with range: 0..255.
    NumCmConnReqs interface{}

    // num issu mgr conn reqs. The type is interface{} with range: 0..255.
    NumIssuMgrConnReqs interface{}

    // num peer fia conn reqs. The type is interface{} with range: 0..255.
    NumPeerFiaConnReqs interface{}

    // is gaspp registered. The type is bool.
    IsGasppRegistered interface{}

    // is cih registered. The type is bool.
    IsCihRegistered interface{}

    // drvr initial startup timestamp. The type is string.
    DrvrInitialStartupTimestamp interface{}

    // drvr current startup timestamp. The type is string.
    DrvrCurrentStartupTimestamp interface{}

    // num intf ports. The type is interface{} with range: 0..4294967295.
    NumIntfPorts interface{}

    // uc weight. The type is interface{} with range: 0..255.
    UcWeight interface{}

    // respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // total asics. The type is interface{} with range: 0..255.
    TotalAsics interface{}

    // issu ready ntfy pending. The type is bool.
    IssuReadyNtfyPending interface{}

    // issu abort sent. The type is bool.
    IssuAbortSent interface{}

    // issu abort rcvd. The type is bool.
    IssuAbortRcvd interface{}

    // fabric mode. The type is interface{} with range: 0..255.
    FabricMode interface{}

    // FC Mode. The type is FcMode.
    FcMode interface{}

    // board rev id. The type is interface{} with range: 0..4294967295.
    BoardRevId interface{}

    // device info. The type is slice of
    // Fia_Nodes_Node_DriverInformation_DeviceInfo.
    DeviceInfo []Fia_Nodes_Node_DriverInformation_DeviceInfo

    // card info. The type is slice of Fia_Nodes_Node_DriverInformation_CardInfo.
    CardInfo []Fia_Nodes_Node_DriverInformation_CardInfo
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetFilter() yfilter.YFilter { return driverInformation.YFilter }

func (driverInformation *Fia_Nodes_Node_DriverInformation) SetFilter(yf yfilter.YFilter) { driverInformation.YFilter = yf }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetGoName(yname string) string {
    if yname == "drv-version" { return "DrvVersion" }
    if yname == "coeff-major-rev" { return "CoeffMajorRev" }
    if yname == "coeff-minor-rev" { return "CoeffMinorRev" }
    if yname == "functional-role" { return "FunctionalRole" }
    if yname == "issu-role" { return "IssuRole" }
    if yname == "node-id" { return "NodeId" }
    if yname == "rack-type" { return "RackType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "is-driver-ready" { return "IsDriverReady" }
    if yname == "card-avail-mask" { return "CardAvailMask" }
    if yname == "asic-avail-mask" { return "AsicAvailMask" }
    if yname == "exp-asic-avail-mask" { return "ExpAsicAvailMask" }
    if yname == "ucmc-ratio" { return "UcmcRatio" }
    if yname == "asic-oper-notify-to-fsdb-pending-bmap" { return "AsicOperNotifyToFsdbPendingBmap" }
    if yname == "is-full-fgid-download-req" { return "IsFullFgidDownloadReq" }
    if yname == "is-fgid-download-in-progress" { return "IsFgidDownloadInProgress" }
    if yname == "is-fgid-download-completed" { return "IsFgidDownloadCompleted" }
    if yname == "fsdb-conn-active" { return "FsdbConnActive" }
    if yname == "fgid-conn-active" { return "FgidConnActive" }
    if yname == "issu-mgr-conn-active" { return "IssuMgrConnActive" }
    if yname == "fsdb-reg-active" { return "FsdbRegActive" }
    if yname == "fgid-reg-active" { return "FgidRegActive" }
    if yname == "issu-mgr-reg-active" { return "IssuMgrRegActive" }
    if yname == "num-pm-conn-reqs" { return "NumPmConnReqs" }
    if yname == "num-fsdb-conn-reqs" { return "NumFsdbConnReqs" }
    if yname == "num-fgid-conn-reqs" { return "NumFgidConnReqs" }
    if yname == "num-fstats-conn-reqs" { return "NumFstatsConnReqs" }
    if yname == "num-cm-conn-reqs" { return "NumCmConnReqs" }
    if yname == "num-issu-mgr-conn-reqs" { return "NumIssuMgrConnReqs" }
    if yname == "num-peer-fia-conn-reqs" { return "NumPeerFiaConnReqs" }
    if yname == "is-gaspp-registered" { return "IsGasppRegistered" }
    if yname == "is-cih-registered" { return "IsCihRegistered" }
    if yname == "drvr-initial-startup-timestamp" { return "DrvrInitialStartupTimestamp" }
    if yname == "drvr-current-startup-timestamp" { return "DrvrCurrentStartupTimestamp" }
    if yname == "num-intf-ports" { return "NumIntfPorts" }
    if yname == "uc-weight" { return "UcWeight" }
    if yname == "respawn-count" { return "RespawnCount" }
    if yname == "total-asics" { return "TotalAsics" }
    if yname == "issu-ready-ntfy-pending" { return "IssuReadyNtfyPending" }
    if yname == "issu-abort-sent" { return "IssuAbortSent" }
    if yname == "issu-abort-rcvd" { return "IssuAbortRcvd" }
    if yname == "fabric-mode" { return "FabricMode" }
    if yname == "fc-mode" { return "FcMode" }
    if yname == "board-rev-id" { return "BoardRevId" }
    if yname == "device-info" { return "DeviceInfo" }
    if yname == "card-info" { return "CardInfo" }
    return ""
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetSegmentPath() string {
    return "driver-information"
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "device-info" {
        for _, c := range driverInformation.DeviceInfo {
            if driverInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_DriverInformation_DeviceInfo{}
        driverInformation.DeviceInfo = append(driverInformation.DeviceInfo, child)
        return &driverInformation.DeviceInfo[len(driverInformation.DeviceInfo)-1]
    }
    if childYangName == "card-info" {
        for _, c := range driverInformation.CardInfo {
            if driverInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_DriverInformation_CardInfo{}
        driverInformation.CardInfo = append(driverInformation.CardInfo, child)
        return &driverInformation.CardInfo[len(driverInformation.CardInfo)-1]
    }
    return nil
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range driverInformation.DeviceInfo {
        children[driverInformation.DeviceInfo[i].GetSegmentPath()] = &driverInformation.DeviceInfo[i]
    }
    for i := range driverInformation.CardInfo {
        children[driverInformation.CardInfo[i].GetSegmentPath()] = &driverInformation.CardInfo[i]
    }
    return children
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["drv-version"] = driverInformation.DrvVersion
    leafs["coeff-major-rev"] = driverInformation.CoeffMajorRev
    leafs["coeff-minor-rev"] = driverInformation.CoeffMinorRev
    leafs["functional-role"] = driverInformation.FunctionalRole
    leafs["issu-role"] = driverInformation.IssuRole
    leafs["node-id"] = driverInformation.NodeId
    leafs["rack-type"] = driverInformation.RackType
    leafs["rack-num"] = driverInformation.RackNum
    leafs["is-driver-ready"] = driverInformation.IsDriverReady
    leafs["card-avail-mask"] = driverInformation.CardAvailMask
    leafs["asic-avail-mask"] = driverInformation.AsicAvailMask
    leafs["exp-asic-avail-mask"] = driverInformation.ExpAsicAvailMask
    leafs["ucmc-ratio"] = driverInformation.UcmcRatio
    leafs["asic-oper-notify-to-fsdb-pending-bmap"] = driverInformation.AsicOperNotifyToFsdbPendingBmap
    leafs["is-full-fgid-download-req"] = driverInformation.IsFullFgidDownloadReq
    leafs["is-fgid-download-in-progress"] = driverInformation.IsFgidDownloadInProgress
    leafs["is-fgid-download-completed"] = driverInformation.IsFgidDownloadCompleted
    leafs["fsdb-conn-active"] = driverInformation.FsdbConnActive
    leafs["fgid-conn-active"] = driverInformation.FgidConnActive
    leafs["issu-mgr-conn-active"] = driverInformation.IssuMgrConnActive
    leafs["fsdb-reg-active"] = driverInformation.FsdbRegActive
    leafs["fgid-reg-active"] = driverInformation.FgidRegActive
    leafs["issu-mgr-reg-active"] = driverInformation.IssuMgrRegActive
    leafs["num-pm-conn-reqs"] = driverInformation.NumPmConnReqs
    leafs["num-fsdb-conn-reqs"] = driverInformation.NumFsdbConnReqs
    leafs["num-fgid-conn-reqs"] = driverInformation.NumFgidConnReqs
    leafs["num-fstats-conn-reqs"] = driverInformation.NumFstatsConnReqs
    leafs["num-cm-conn-reqs"] = driverInformation.NumCmConnReqs
    leafs["num-issu-mgr-conn-reqs"] = driverInformation.NumIssuMgrConnReqs
    leafs["num-peer-fia-conn-reqs"] = driverInformation.NumPeerFiaConnReqs
    leafs["is-gaspp-registered"] = driverInformation.IsGasppRegistered
    leafs["is-cih-registered"] = driverInformation.IsCihRegistered
    leafs["drvr-initial-startup-timestamp"] = driverInformation.DrvrInitialStartupTimestamp
    leafs["drvr-current-startup-timestamp"] = driverInformation.DrvrCurrentStartupTimestamp
    leafs["num-intf-ports"] = driverInformation.NumIntfPorts
    leafs["uc-weight"] = driverInformation.UcWeight
    leafs["respawn-count"] = driverInformation.RespawnCount
    leafs["total-asics"] = driverInformation.TotalAsics
    leafs["issu-ready-ntfy-pending"] = driverInformation.IssuReadyNtfyPending
    leafs["issu-abort-sent"] = driverInformation.IssuAbortSent
    leafs["issu-abort-rcvd"] = driverInformation.IssuAbortRcvd
    leafs["fabric-mode"] = driverInformation.FabricMode
    leafs["fc-mode"] = driverInformation.FcMode
    leafs["board-rev-id"] = driverInformation.BoardRevId
    return leafs
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetBundleName() string { return "cisco_ios_xr" }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetYangName() string { return "driver-information" }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (driverInformation *Fia_Nodes_Node_DriverInformation) SetParent(parent types.Entity) { driverInformation.parent = parent }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetParent() types.Entity { return driverInformation.parent }

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_DriverInformation_DeviceInfo
// device info
type Fia_Nodes_Node_DriverInformation_DeviceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // is valid. The type is bool.
    IsValid interface{}

    // fapid. The type is interface{} with range: 0..4294967295.
    Fapid interface{}

    // hotplug event. The type is interface{} with range: 0..4294967295.
    HotplugEvent interface{}

    // Slice State. The type is SliceState.
    SliceState interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is AsicOperState.
    OperState interface{}

    // Asic State. The type is AsicAccessState.
    AsicState interface{}

    // last init cause. The type is AsicInitMethod.
    LastInitCause interface{}

    // num pon resets. The type is interface{} with range: 0..4294967295.
    NumPonResets interface{}

    // num hard resets. The type is interface{} with range: 0..4294967295.
    NumHardResets interface{}

    // local switch state. The type is bool.
    LocalSwitchState interface{}

    // asic id.
    AsicId Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetFilter() yfilter.YFilter { return deviceInfo.YFilter }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) SetFilter(yf yfilter.YFilter) { deviceInfo.YFilter = yf }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetGoName(yname string) string {
    if yname == "is-valid" { return "IsValid" }
    if yname == "fapid" { return "Fapid" }
    if yname == "hotplug-event" { return "HotplugEvent" }
    if yname == "slice-state" { return "SliceState" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state" { return "OperState" }
    if yname == "asic-state" { return "AsicState" }
    if yname == "last-init-cause" { return "LastInitCause" }
    if yname == "num-pon-resets" { return "NumPonResets" }
    if yname == "num-hard-resets" { return "NumHardResets" }
    if yname == "local-switch-state" { return "LocalSwitchState" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetSegmentPath() string {
    return "device-info"
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &deviceInfo.AsicId
    }
    return nil
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &deviceInfo.AsicId
    return children
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-valid"] = deviceInfo.IsValid
    leafs["fapid"] = deviceInfo.Fapid
    leafs["hotplug-event"] = deviceInfo.HotplugEvent
    leafs["slice-state"] = deviceInfo.SliceState
    leafs["admin-state"] = deviceInfo.AdminState
    leafs["oper-state"] = deviceInfo.OperState
    leafs["asic-state"] = deviceInfo.AsicState
    leafs["last-init-cause"] = deviceInfo.LastInitCause
    leafs["num-pon-resets"] = deviceInfo.NumPonResets
    leafs["num-hard-resets"] = deviceInfo.NumHardResets
    leafs["local-switch-state"] = deviceInfo.LocalSwitchState
    return leafs
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetYangName() string { return "device-info" }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) SetParent(parent types.Entity) { deviceInfo.parent = parent }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetParent() types.Entity { return deviceInfo.parent }

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetParentYangName() string { return "driver-information" }

// Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId
// asic id
type Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetParentYangName() string { return "device-info" }

// Fia_Nodes_Node_DriverInformation_CardInfo
// card info
type Fia_Nodes_Node_DriverInformation_CardInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // card name. The type is string.
    CardName interface{}

    // slot no. The type is interface{} with range: -2147483648..2147483647.
    SlotNo interface{}

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // card state. The type is interface{} with range: 0..255.
    CardState interface{}

    // exp num asics. The type is interface{} with range: 0..4294967295.
    ExpNumAsics interface{}

    // exp num asics per fsdb. The type is interface{} with range: 0..4294967295.
    ExpNumAsicsPerFsdb interface{}

    // is powered. The type is bool.
    IsPowered interface{}

    // cxp avail bitmap. The type is interface{} with range:
    // 0..18446744073709551615.
    CxpAvailBitmap interface{}

    // num ilkns per asic. The type is interface{} with range: 0..4294967295.
    NumIlknsPerAsic interface{}

    // num local ports per ilkn. The type is interface{} with range:
    // 0..4294967295.
    NumLocalPortsPerIlkn interface{}

    // num cos per port. The type is interface{} with range: 0..255.
    NumCosPerPort interface{}

    // oir circular buffer.
    OirCircularBuffer Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetFilter() yfilter.YFilter { return cardInfo.YFilter }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) SetFilter(yf yfilter.YFilter) { cardInfo.YFilter = yf }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetGoName(yname string) string {
    if yname == "card-type" { return "CardType" }
    if yname == "card-name" { return "CardName" }
    if yname == "slot-no" { return "SlotNo" }
    if yname == "card-flag" { return "CardFlag" }
    if yname == "evt-flag" { return "EvtFlag" }
    if yname == "reg-flag" { return "RegFlag" }
    if yname == "instance" { return "Instance" }
    if yname == "card-state" { return "CardState" }
    if yname == "exp-num-asics" { return "ExpNumAsics" }
    if yname == "exp-num-asics-per-fsdb" { return "ExpNumAsicsPerFsdb" }
    if yname == "is-powered" { return "IsPowered" }
    if yname == "cxp-avail-bitmap" { return "CxpAvailBitmap" }
    if yname == "num-ilkns-per-asic" { return "NumIlknsPerAsic" }
    if yname == "num-local-ports-per-ilkn" { return "NumLocalPortsPerIlkn" }
    if yname == "num-cos-per-port" { return "NumCosPerPort" }
    if yname == "oir-circular-buffer" { return "OirCircularBuffer" }
    return ""
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetSegmentPath() string {
    return "card-info"
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oir-circular-buffer" {
        return &cardInfo.OirCircularBuffer
    }
    return nil
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["oir-circular-buffer"] = &cardInfo.OirCircularBuffer
    return children
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type"] = cardInfo.CardType
    leafs["card-name"] = cardInfo.CardName
    leafs["slot-no"] = cardInfo.SlotNo
    leafs["card-flag"] = cardInfo.CardFlag
    leafs["evt-flag"] = cardInfo.EvtFlag
    leafs["reg-flag"] = cardInfo.RegFlag
    leafs["instance"] = cardInfo.Instance
    leafs["card-state"] = cardInfo.CardState
    leafs["exp-num-asics"] = cardInfo.ExpNumAsics
    leafs["exp-num-asics-per-fsdb"] = cardInfo.ExpNumAsicsPerFsdb
    leafs["is-powered"] = cardInfo.IsPowered
    leafs["cxp-avail-bitmap"] = cardInfo.CxpAvailBitmap
    leafs["num-ilkns-per-asic"] = cardInfo.NumIlknsPerAsic
    leafs["num-local-ports-per-ilkn"] = cardInfo.NumLocalPortsPerIlkn
    leafs["num-cos-per-port"] = cardInfo.NumCosPerPort
    return leafs
}

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetBundleName() string { return "cisco_ios_xr" }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetYangName() string { return "card-info" }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) SetParent(parent types.Entity) { cardInfo.parent = parent }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetParent() types.Entity { return cardInfo.parent }

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetParentYangName() string { return "driver-information" }

// Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer
// oir circular buffer
type Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // count. The type is interface{} with range: -2147483648..2147483647.
    Count interface{}

    // start. The type is interface{} with range: -2147483648..2147483647.
    Start interface{}

    // end. The type is interface{} with range: -2147483648..2147483647.
    End interface{}

    // fia oir info. The type is slice of
    // Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo.
    FiaOirInfo []Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetFilter() yfilter.YFilter { return oirCircularBuffer.YFilter }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) SetFilter(yf yfilter.YFilter) { oirCircularBuffer.YFilter = yf }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "fia-oir-info" { return "FiaOirInfo" }
    return ""
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetSegmentPath() string {
    return "oir-circular-buffer"
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fia-oir-info" {
        for _, c := range oirCircularBuffer.FiaOirInfo {
            if oirCircularBuffer.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo{}
        oirCircularBuffer.FiaOirInfo = append(oirCircularBuffer.FiaOirInfo, child)
        return &oirCircularBuffer.FiaOirInfo[len(oirCircularBuffer.FiaOirInfo)-1]
    }
    return nil
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oirCircularBuffer.FiaOirInfo {
        children[oirCircularBuffer.FiaOirInfo[i].GetSegmentPath()] = &oirCircularBuffer.FiaOirInfo[i]
    }
    return children
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = oirCircularBuffer.Count
    leafs["start"] = oirCircularBuffer.Start
    leafs["end"] = oirCircularBuffer.End
    return leafs
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetYangName() string { return "oir-circular-buffer" }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) SetParent(parent types.Entity) { oirCircularBuffer.parent = parent }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetParent() types.Entity { return oirCircularBuffer.parent }

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetParentYangName() string { return "card-info" }

// Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo
// fia oir info
type Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // rack num. The type is interface{} with range: -2147483648..2147483647.
    RackNum interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // cur card state. The type is interface{} with range:
    // -2147483648..2147483647.
    CurCardState interface{}
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetFilter() yfilter.YFilter { return fiaOirInfo.YFilter }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) SetFilter(yf yfilter.YFilter) { fiaOirInfo.YFilter = yf }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetGoName(yname string) string {
    if yname == "card-flag" { return "CardFlag" }
    if yname == "card-type" { return "CardType" }
    if yname == "reg-flag" { return "RegFlag" }
    if yname == "evt-flag" { return "EvtFlag" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "instance" { return "Instance" }
    if yname == "cur-card-state" { return "CurCardState" }
    return ""
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetSegmentPath() string {
    return "fia-oir-info"
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-flag"] = fiaOirInfo.CardFlag
    leafs["card-type"] = fiaOirInfo.CardType
    leafs["reg-flag"] = fiaOirInfo.RegFlag
    leafs["evt-flag"] = fiaOirInfo.EvtFlag
    leafs["rack-num"] = fiaOirInfo.RackNum
    leafs["instance"] = fiaOirInfo.Instance
    leafs["cur-card-state"] = fiaOirInfo.CurCardState
    return leafs
}

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetYangName() string { return "fia-oir-info" }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) SetParent(parent types.Entity) { fiaOirInfo.parent = parent }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetParent() types.Entity { return fiaOirInfo.parent }

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetParentYangName() string { return "oir-circular-buffer" }

// Fia_Nodes_Node_ClearStatistics
// Clear statistics information
type Fia_Nodes_Node_ClearStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance table for clear statistics information.
    AsicInstances Fia_Nodes_Node_ClearStatistics_AsicInstances
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetFilter() yfilter.YFilter { return clearStatistics.YFilter }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) SetFilter(yf yfilter.YFilter) { clearStatistics.YFilter = yf }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetGoName(yname string) string {
    if yname == "asic-instances" { return "AsicInstances" }
    return ""
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetSegmentPath() string {
    return "clear-statistics"
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-instances" {
        return &clearStatistics.AsicInstances
    }
    return nil
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-instances"] = &clearStatistics.AsicInstances
    return children
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetYangName() string { return "clear-statistics" }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) SetParent(parent types.Entity) { clearStatistics.parent = parent }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetParent() types.Entity { return clearStatistics.parent }

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_ClearStatistics_AsicInstances
// Instance table for clear statistics
// information
type Fia_Nodes_Node_ClearStatistics_AsicInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Asic instance to be cleared. The type is slice of
    // Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance.
    AsicInstance []Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetFilter() yfilter.YFilter { return asicInstances.YFilter }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) SetFilter(yf yfilter.YFilter) { asicInstances.YFilter = yf }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetGoName(yname string) string {
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetSegmentPath() string {
    return "asic-instances"
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-instance" {
        for _, c := range asicInstances.AsicInstance {
            if asicInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance{}
        asicInstances.AsicInstance = append(asicInstances.AsicInstance, child)
        return &asicInstances.AsicInstance[len(asicInstances.AsicInstance)-1]
    }
    return nil
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range asicInstances.AsicInstance {
        children[asicInstances.AsicInstance[i].GetSegmentPath()] = &asicInstances.AsicInstance[i]
    }
    return children
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetBundleName() string { return "cisco_ios_xr" }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetYangName() string { return "asic-instances" }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) SetParent(parent types.Entity) { asicInstances.parent = parent }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetParent() types.Entity { return asicInstances.parent }

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetParentYangName() string { return "clear-statistics" }

// Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance
// Asic instance to be cleared
type Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Asic instance. The type is interface{} with range:
    // 0..255.
    AsicInstance interface{}

    // Clear value. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    Instance interface{}
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetFilter() yfilter.YFilter { return asicInstance.YFilter }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) SetFilter(yf yfilter.YFilter) { asicInstance.YFilter = yf }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetGoName(yname string) string {
    if yname == "asic-instance" { return "AsicInstance" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetSegmentPath() string {
    return "asic-instance" + "[asic-instance='" + fmt.Sprintf("%v", asicInstance.AsicInstance) + "']"
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asic-instance"] = asicInstance.AsicInstance
    leafs["instance"] = asicInstance.Instance
    return leafs
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetBundleName() string { return "cisco_ios_xr" }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetYangName() string { return "asic-instance" }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) SetParent(parent types.Entity) { asicInstance.parent = parent }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetParent() types.Entity { return asicInstance.parent }

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetParentYangName() string { return "asic-instances" }

// Fia_Nodes_Node_TxLinkInformation
// FIA link TX information
type Fia_Nodes_Node_TxLinkInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link table for tx information.
    TxStatusOptionTable Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetFilter() yfilter.YFilter { return txLinkInformation.YFilter }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) SetFilter(yf yfilter.YFilter) { txLinkInformation.YFilter = yf }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetGoName(yname string) string {
    if yname == "tx-status-option-table" { return "TxStatusOptionTable" }
    return ""
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetSegmentPath() string {
    return "tx-link-information"
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-status-option-table" {
        return &txLinkInformation.TxStatusOptionTable
    }
    return nil
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-status-option-table"] = &txLinkInformation.TxStatusOptionTable
    return children
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetBundleName() string { return "cisco_ios_xr" }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetYangName() string { return "tx-link-information" }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) SetParent(parent types.Entity) { txLinkInformation.parent = parent }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetParent() types.Entity { return txLinkInformation.parent }

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable
// Link table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Option: data, ctrl, all- for now none.
    TxStatusOption Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetFilter() yfilter.YFilter { return txStatusOptionTable.YFilter }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) SetFilter(yf yfilter.YFilter) { txStatusOptionTable.YFilter = yf }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetGoName(yname string) string {
    if yname == "tx-status-option" { return "TxStatusOption" }
    return ""
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetSegmentPath() string {
    return "tx-status-option-table"
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-status-option" {
        return &txStatusOptionTable.TxStatusOption
    }
    return nil
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-status-option"] = &txStatusOptionTable.TxStatusOption
    return children
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetBundleName() string { return "cisco_ios_xr" }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetYangName() string { return "tx-status-option-table" }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) SetParent(parent types.Entity) { txStatusOptionTable.parent = parent }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetParent() types.Entity { return txStatusOptionTable.parent }

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetParentYangName() string { return "tx-link-information" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption
// Option: data, ctrl, all- for now none
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance table for tx information.
    TxAsicInstances Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetFilter() yfilter.YFilter { return txStatusOption.YFilter }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) SetFilter(yf yfilter.YFilter) { txStatusOption.YFilter = yf }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetGoName(yname string) string {
    if yname == "tx-asic-instances" { return "TxAsicInstances" }
    return ""
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetSegmentPath() string {
    return "tx-status-option"
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-asic-instances" {
        return &txStatusOption.TxAsicInstances
    }
    return nil
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-asic-instances"] = &txStatusOption.TxAsicInstances
    return children
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetBundleName() string { return "cisco_ios_xr" }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetYangName() string { return "tx-status-option" }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) SetParent(parent types.Entity) { txStatusOption.parent = parent }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetParent() types.Entity { return txStatusOption.parent }

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetParentYangName() string { return "tx-status-option-table" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances
// Instance table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance number for tx link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance.
    TxAsicInstance []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetFilter() yfilter.YFilter { return txAsicInstances.YFilter }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) SetFilter(yf yfilter.YFilter) { txAsicInstances.YFilter = yf }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetGoName(yname string) string {
    if yname == "tx-asic-instance" { return "TxAsicInstance" }
    return ""
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetSegmentPath() string {
    return "tx-asic-instances"
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-asic-instance" {
        for _, c := range txAsicInstances.TxAsicInstance {
            if txAsicInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance{}
        txAsicInstances.TxAsicInstance = append(txAsicInstances.TxAsicInstance, child)
        return &txAsicInstances.TxAsicInstance[len(txAsicInstances.TxAsicInstance)-1]
    }
    return nil
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range txAsicInstances.TxAsicInstance {
        children[txAsicInstances.TxAsicInstance[i].GetSegmentPath()] = &txAsicInstances.TxAsicInstance[i]
    }
    return children
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetBundleName() string { return "cisco_ios_xr" }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetYangName() string { return "tx-asic-instances" }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) SetParent(parent types.Entity) { txAsicInstances.parent = parent }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetParent() types.Entity { return txAsicInstances.parent }

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetParentYangName() string { return "tx-status-option" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance
// Instance number for tx link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Transmit instance. The type is interface{} with
    // range: 0..255.
    Instance interface{}

    // Link table for tx information.
    TxLinks Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetFilter() yfilter.YFilter { return txAsicInstance.YFilter }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) SetFilter(yf yfilter.YFilter) { txAsicInstance.YFilter = yf }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "tx-links" { return "TxLinks" }
    return ""
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetSegmentPath() string {
    return "tx-asic-instance" + "[instance='" + fmt.Sprintf("%v", txAsicInstance.Instance) + "']"
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-links" {
        return &txAsicInstance.TxLinks
    }
    return nil
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tx-links"] = &txAsicInstance.TxLinks
    return children
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance"] = txAsicInstance.Instance
    return leafs
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetBundleName() string { return "cisco_ios_xr" }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetYangName() string { return "tx-asic-instance" }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) SetParent(parent types.Entity) { txAsicInstance.parent = parent }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetParent() types.Entity { return txAsicInstance.parent }

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetParentYangName() string { return "tx-asic-instances" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks
// Link table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link number for tx link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink.
    TxLink []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetFilter() yfilter.YFilter { return txLinks.YFilter }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) SetFilter(yf yfilter.YFilter) { txLinks.YFilter = yf }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetGoName(yname string) string {
    if yname == "tx-link" { return "TxLink" }
    return ""
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetSegmentPath() string {
    return "tx-links"
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-link" {
        for _, c := range txLinks.TxLink {
            if txLinks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink{}
        txLinks.TxLink = append(txLinks.TxLink, child)
        return &txLinks.TxLink[len(txLinks.TxLink)-1]
    }
    return nil
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range txLinks.TxLink {
        children[txLinks.TxLink[i].GetSegmentPath()] = &txLinks.TxLink[i]
    }
    return children
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetBundleName() string { return "cisco_ios_xr" }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetYangName() string { return "tx-links" }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) SetParent(parent types.Entity) { txLinks.parent = parent }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetParent() types.Entity { return txLinks.parent }

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetParentYangName() string { return "tx-asic-instance" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink
// Link number for tx link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start number. The type is interface{} with range: 0..47.
    StartNumber interface{}

    // End number. The type is interface{} with range: 0..47.
    EndNumber interface{}

    // Single link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink.
    TxLink []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetFilter() yfilter.YFilter { return txLink.YFilter }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) SetFilter(yf yfilter.YFilter) { txLink.YFilter = yf }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetGoName(yname string) string {
    if yname == "start-number" { return "StartNumber" }
    if yname == "end-number" { return "EndNumber" }
    if yname == "tx-link" { return "TxLink" }
    return ""
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetSegmentPath() string {
    return "tx-link"
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tx-link" {
        for _, c := range txLink.TxLink {
            if txLink.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink{}
        txLink.TxLink = append(txLink.TxLink, child)
        return &txLink.TxLink[len(txLink.TxLink)-1]
    }
    return nil
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range txLink.TxLink {
        children[txLink.TxLink[i].GetSegmentPath()] = &txLink.TxLink[i]
    }
    return children
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start-number"] = txLink.StartNumber
    leafs["end-number"] = txLink.EndNumber
    return leafs
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetBundleName() string { return "cisco_ios_xr" }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetYangName() string { return "tx-link" }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) SetParent(parent types.Entity) { txLink.parent = parent }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetParent() types.Entity { return txLink.parent }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetParentYangName() string { return "tx-links" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink
// Single link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single Link. The type is interface{} with range:
    // -2147483648..2147483647.
    Link interface{}

    // speed. The type is interface{} with range: 0..4294967295.
    Speed interface{}

    // stage. The type is interface{} with range: 0..255.
    Stage interface{}

    // is link valid. The type is bool.
    IsLinkValid interface{}

    // is conf pending. The type is bool.
    IsConfPending interface{}

    // is power enabled. The type is bool.
    IsPowerEnabled interface{}

    // coeff1. The type is interface{} with range: 0..4294967295.
    Coeff1 interface{}

    // coeff2. The type is interface{} with range: 0..4294967295.
    Coeff2 interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // num admin shuts. The type is interface{} with range: 0..4294967295.
    NumAdminShuts interface{}

    // this link.
    ThisLink Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink

    // far end link.
    FarEndLink Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink

    // stats.
    Stats Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats

    // history.
    History Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetFilter() yfilter.YFilter { return txLink.YFilter }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) SetFilter(yf yfilter.YFilter) { txLink.YFilter = yf }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetGoName(yname string) string {
    if yname == "link" { return "Link" }
    if yname == "speed" { return "Speed" }
    if yname == "stage" { return "Stage" }
    if yname == "is-link-valid" { return "IsLinkValid" }
    if yname == "is-conf-pending" { return "IsConfPending" }
    if yname == "is-power-enabled" { return "IsPowerEnabled" }
    if yname == "coeff1" { return "Coeff1" }
    if yname == "coeff2" { return "Coeff2" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state" { return "OperState" }
    if yname == "error-state" { return "ErrorState" }
    if yname == "num-admin-shuts" { return "NumAdminShuts" }
    if yname == "this-link" { return "ThisLink" }
    if yname == "far-end-link" { return "FarEndLink" }
    if yname == "stats" { return "Stats" }
    if yname == "history" { return "History" }
    return ""
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetSegmentPath() string {
    return "tx-link" + "[link='" + fmt.Sprintf("%v", txLink.Link) + "']"
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "this-link" {
        return &txLink.ThisLink
    }
    if childYangName == "far-end-link" {
        return &txLink.FarEndLink
    }
    if childYangName == "stats" {
        return &txLink.Stats
    }
    if childYangName == "history" {
        return &txLink.History
    }
    return nil
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["this-link"] = &txLink.ThisLink
    children["far-end-link"] = &txLink.FarEndLink
    children["stats"] = &txLink.Stats
    children["history"] = &txLink.History
    return children
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link"] = txLink.Link
    leafs["speed"] = txLink.Speed
    leafs["stage"] = txLink.Stage
    leafs["is-link-valid"] = txLink.IsLinkValid
    leafs["is-conf-pending"] = txLink.IsConfPending
    leafs["is-power-enabled"] = txLink.IsPowerEnabled
    leafs["coeff1"] = txLink.Coeff1
    leafs["coeff2"] = txLink.Coeff2
    leafs["admin-state"] = txLink.AdminState
    leafs["oper-state"] = txLink.OperState
    leafs["error-state"] = txLink.ErrorState
    leafs["num-admin-shuts"] = txLink.NumAdminShuts
    return leafs
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetBundleName() string { return "cisco_ios_xr" }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetYangName() string { return "tx-link" }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) SetParent(parent types.Entity) { txLink.parent = parent }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetParent() types.Entity { return txLink.parent }

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetParentYangName() string { return "tx-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink
// this link
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetFilter() yfilter.YFilter { return thisLink.YFilter }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) SetFilter(yf yfilter.YFilter) { thisLink.YFilter = yf }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-stage" { return "LinkStage" }
    if yname == "link-num" { return "LinkNum" }
    if yname == "phy-link-num" { return "PhyLinkNum" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetSegmentPath() string {
    return "this-link"
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &thisLink.AsicId
    }
    return nil
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &thisLink.AsicId
    return children
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = thisLink.LinkType
    leafs["link-stage"] = thisLink.LinkStage
    leafs["link-num"] = thisLink.LinkNum
    leafs["phy-link-num"] = thisLink.PhyLinkNum
    return leafs
}

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetBundleName() string { return "cisco_ios_xr" }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetYangName() string { return "this-link" }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) SetParent(parent types.Entity) { thisLink.parent = parent }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetParent() types.Entity { return thisLink.parent }

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetParentYangName() string { return "tx-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId
// asic id
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetParentYangName() string { return "this-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink
// far end link
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link Type. The type is Link.
    LinkType interface{}

    // Link Stage. The type is LinkStage.
    LinkStage interface{}

    // link num. The type is interface{} with range: 0..4294967295.
    LinkNum interface{}

    // phy link num. The type is interface{} with range: 0..4294967295.
    PhyLinkNum interface{}

    // asic id.
    AsicId Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetFilter() yfilter.YFilter { return farEndLink.YFilter }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) SetFilter(yf yfilter.YFilter) { farEndLink.YFilter = yf }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetGoName(yname string) string {
    if yname == "link-type" { return "LinkType" }
    if yname == "link-stage" { return "LinkStage" }
    if yname == "link-num" { return "LinkNum" }
    if yname == "phy-link-num" { return "PhyLinkNum" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetSegmentPath() string {
    return "far-end-link"
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &farEndLink.AsicId
    }
    return nil
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &farEndLink.AsicId
    return children
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-type"] = farEndLink.LinkType
    leafs["link-stage"] = farEndLink.LinkStage
    leafs["link-num"] = farEndLink.LinkNum
    leafs["phy-link-num"] = farEndLink.PhyLinkNum
    return leafs
}

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetBundleName() string { return "cisco_ios_xr" }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetYangName() string { return "far-end-link" }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) SetParent(parent types.Entity) { farEndLink.parent = parent }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetParent() types.Entity { return farEndLink.parent }

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetParentYangName() string { return "tx-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId
// asic id
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetParentYangName() string { return "far-end-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats
// stats
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // dummy. The type is interface{} with range: 0..4294967295.
    Dummy interface{}
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetFilter() yfilter.YFilter { return stats.YFilter }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) SetFilter(yf yfilter.YFilter) { stats.YFilter = yf }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetGoName(yname string) string {
    if yname == "dummy" { return "Dummy" }
    return ""
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetSegmentPath() string {
    return "stats"
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dummy"] = stats.Dummy
    return leafs
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetBundleName() string { return "cisco_ios_xr" }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetYangName() string { return "stats" }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) SetParent(parent types.Entity) { stats.parent = parent }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetParent() types.Entity { return stats.parent }

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetParentYangName() string { return "tx-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History
// history
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // histnum. The type is interface{} with range: 0..255.
    Histnum interface{}

    // start index. The type is interface{} with range: 0..255.
    StartIndex interface{}

    // hist. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist.
    Hist []Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetFilter() yfilter.YFilter { return history.YFilter }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) SetFilter(yf yfilter.YFilter) { history.YFilter = yf }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetGoName(yname string) string {
    if yname == "histnum" { return "Histnum" }
    if yname == "start-index" { return "StartIndex" }
    if yname == "hist" { return "Hist" }
    return ""
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetSegmentPath() string {
    return "history"
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hist" {
        for _, c := range history.Hist {
            if history.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist{}
        history.Hist = append(history.Hist, child)
        return &history.Hist[len(history.Hist)-1]
    }
    return nil
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range history.Hist {
        children[history.Hist[i].GetSegmentPath()] = &history.Hist[i]
    }
    return children
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["histnum"] = history.Histnum
    leafs["start-index"] = history.StartIndex
    return leafs
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetBundleName() string { return "cisco_ios_xr" }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetYangName() string { return "history" }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) SetParent(parent types.Entity) { history.parent = parent }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetParent() types.Entity { return history.parent }

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetParentYangName() string { return "tx-link" }

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist
// hist
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is OperState.
    OperState interface{}

    // Error State. The type is LinkErrorState.
    ErrorState interface{}

    // timestamp. The type is interface{} with range: 0..18446744073709551615.
    Timestamp interface{}

    // reasons. The type is string.
    Reasons interface{}
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetFilter() yfilter.YFilter { return hist.YFilter }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) SetFilter(yf yfilter.YFilter) { hist.YFilter = yf }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetGoName(yname string) string {
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state" { return "OperState" }
    if yname == "error-state" { return "ErrorState" }
    if yname == "timestamp" { return "Timestamp" }
    if yname == "reasons" { return "Reasons" }
    return ""
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetSegmentPath() string {
    return "hist"
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["admin-state"] = hist.AdminState
    leafs["oper-state"] = hist.OperState
    leafs["error-state"] = hist.ErrorState
    leafs["timestamp"] = hist.Timestamp
    leafs["reasons"] = hist.Reasons
    return leafs
}

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetBundleName() string { return "cisco_ios_xr" }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetYangName() string { return "hist" }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) SetParent(parent types.Entity) { hist.parent = parent }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetParent() types.Entity { return hist.parent }

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetParentYangName() string { return "history" }

// Fia_Nodes_Node_DiagShell
// FIA diag shell information
type Fia_Nodes_Node_DiagShell struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unit table for diag shell.
    DiagShellUnits Fia_Nodes_Node_DiagShell_DiagShellUnits
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetFilter() yfilter.YFilter { return diagShell.YFilter }

func (diagShell *Fia_Nodes_Node_DiagShell) SetFilter(yf yfilter.YFilter) { diagShell.YFilter = yf }

func (diagShell *Fia_Nodes_Node_DiagShell) GetGoName(yname string) string {
    if yname == "diag-shell-units" { return "DiagShellUnits" }
    return ""
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetSegmentPath() string {
    return "diag-shell"
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diag-shell-units" {
        return &diagShell.DiagShellUnits
    }
    return nil
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["diag-shell-units"] = &diagShell.DiagShellUnits
    return children
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetBundleName() string { return "cisco_ios_xr" }

func (diagShell *Fia_Nodes_Node_DiagShell) GetYangName() string { return "diag-shell" }

func (diagShell *Fia_Nodes_Node_DiagShell) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diagShell *Fia_Nodes_Node_DiagShell) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diagShell *Fia_Nodes_Node_DiagShell) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diagShell *Fia_Nodes_Node_DiagShell) SetParent(parent types.Entity) { diagShell.parent = parent }

func (diagShell *Fia_Nodes_Node_DiagShell) GetParent() types.Entity { return diagShell.parent }

func (diagShell *Fia_Nodes_Node_DiagShell) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_DiagShell_DiagShellUnits
// Unit table for diag shell
type Fia_Nodes_Node_DiagShell_DiagShellUnits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unit number for diag shell statistics. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit.
    DiagShellUnit []Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetFilter() yfilter.YFilter { return diagShellUnits.YFilter }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) SetFilter(yf yfilter.YFilter) { diagShellUnits.YFilter = yf }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetGoName(yname string) string {
    if yname == "diag-shell-unit" { return "DiagShellUnit" }
    return ""
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetSegmentPath() string {
    return "diag-shell-units"
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "diag-shell-unit" {
        for _, c := range diagShellUnits.DiagShellUnit {
            if diagShellUnits.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit{}
        diagShellUnits.DiagShellUnit = append(diagShellUnits.DiagShellUnit, child)
        return &diagShellUnits.DiagShellUnit[len(diagShellUnits.DiagShellUnit)-1]
    }
    return nil
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range diagShellUnits.DiagShellUnit {
        children[diagShellUnits.DiagShellUnit[i].GetSegmentPath()] = &diagShellUnits.DiagShellUnit[i]
    }
    return children
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetBundleName() string { return "cisco_ios_xr" }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetYangName() string { return "diag-shell-units" }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) SetParent(parent types.Entity) { diagShellUnits.parent = parent }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetParent() types.Entity { return diagShellUnits.parent }

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetParentYangName() string { return "diag-shell" }

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit
// Unit number for diag shell statistics
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Unit number. The type is interface{} with range:
    // 0..63.
    Unit interface{}

    // Command table for diag shell.
    Commands Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetFilter() yfilter.YFilter { return diagShellUnit.YFilter }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) SetFilter(yf yfilter.YFilter) { diagShellUnit.YFilter = yf }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetGoName(yname string) string {
    if yname == "unit" { return "Unit" }
    if yname == "commands" { return "Commands" }
    return ""
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetSegmentPath() string {
    return "diag-shell-unit" + "[unit='" + fmt.Sprintf("%v", diagShellUnit.Unit) + "']"
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "commands" {
        return &diagShellUnit.Commands
    }
    return nil
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["commands"] = &diagShellUnit.Commands
    return children
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unit"] = diagShellUnit.Unit
    return leafs
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetBundleName() string { return "cisco_ios_xr" }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetYangName() string { return "diag-shell-unit" }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) SetParent(parent types.Entity) { diagShellUnit.parent = parent }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetParent() types.Entity { return diagShellUnit.parent }

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetParentYangName() string { return "diag-shell-units" }

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands
// Command table for diag shell
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Command for diag shell statistics. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command.
    Command []Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetFilter() yfilter.YFilter { return commands.YFilter }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) SetFilter(yf yfilter.YFilter) { commands.YFilter = yf }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetGoName(yname string) string {
    if yname == "command" { return "Command" }
    return ""
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetSegmentPath() string {
    return "commands"
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "command" {
        for _, c := range commands.Command {
            if commands.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command{}
        commands.Command = append(commands.Command, child)
        return &commands.Command[len(commands.Command)-1]
    }
    return nil
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range commands.Command {
        children[commands.Command[i].GetSegmentPath()] = &commands.Command[i]
    }
    return children
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetBundleName() string { return "cisco_ios_xr" }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetYangName() string { return "commands" }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) SetParent(parent types.Entity) { commands.parent = parent }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetParent() types.Entity { return commands.parent }

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetParentYangName() string { return "diag-shell-unit" }

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command
// Command for diag shell statistics
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Shell command. The type is string.
    Cmd interface{}

    // Added to support datalist. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output.
    Output []Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetFilter() yfilter.YFilter { return command.YFilter }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) SetFilter(yf yfilter.YFilter) { command.YFilter = yf }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetGoName(yname string) string {
    if yname == "cmd" { return "Cmd" }
    if yname == "output" { return "Output" }
    return ""
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetSegmentPath() string {
    return "command" + "[cmd='" + fmt.Sprintf("%v", command.Cmd) + "']"
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output" {
        for _, c := range command.Output {
            if command.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output{}
        command.Output = append(command.Output, child)
        return &command.Output[len(command.Output)-1]
    }
    return nil
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range command.Output {
        children[command.Output[i].GetSegmentPath()] = &command.Output[i]
    }
    return children
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmd"] = command.Cmd
    return leafs
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetBundleName() string { return "cisco_ios_xr" }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetYangName() string { return "command" }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) SetParent(parent types.Entity) { command.parent = parent }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetParent() types.Entity { return command.parent }

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetParentYangName() string { return "commands" }

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output
// Added to support datalist
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. First line. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Output interface{}

    // output xr. The type is string.
    OutputXr interface{}
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetGoName(yname string) string {
    if yname == "output" { return "Output" }
    if yname == "output-xr" { return "OutputXr" }
    return ""
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetSegmentPath() string {
    return "output" + "[output='" + fmt.Sprintf("%v", output.Output) + "']"
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["output"] = output.Output
    leafs["output-xr"] = output.OutputXr
    return leafs
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetBundleName() string { return "cisco_ios_xr" }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetYangName() string { return "output" }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetParent() types.Entity { return output.parent }

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetParentYangName() string { return "command" }

// Fia_Nodes_Node_OirHistory
// FIA operational data of oir history
type Fia_Nodes_Node_OirHistory struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flag table for history.
    Flags Fia_Nodes_Node_OirHistory_Flags
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetFilter() yfilter.YFilter { return oirHistory.YFilter }

func (oirHistory *Fia_Nodes_Node_OirHistory) SetFilter(yf yfilter.YFilter) { oirHistory.YFilter = yf }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetGoName(yname string) string {
    if yname == "flags" { return "Flags" }
    return ""
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetSegmentPath() string {
    return "oir-history"
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flags" {
        return &oirHistory.Flags
    }
    return nil
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flags"] = &oirHistory.Flags
    return children
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetBundleName() string { return "cisco_ios_xr" }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetYangName() string { return "oir-history" }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oirHistory *Fia_Nodes_Node_OirHistory) SetParent(parent types.Entity) { oirHistory.parent = parent }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetParent() types.Entity { return oirHistory.parent }

func (oirHistory *Fia_Nodes_Node_OirHistory) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_OirHistory_Flags
// Flag table for history
type Fia_Nodes_Node_OirHistory_Flags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flag value for physical location. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag.
    Flag []Fia_Nodes_Node_OirHistory_Flags_Flag
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetFilter() yfilter.YFilter { return flags.YFilter }

func (flags *Fia_Nodes_Node_OirHistory_Flags) SetFilter(yf yfilter.YFilter) { flags.YFilter = yf }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetGoName(yname string) string {
    if yname == "flag" { return "Flag" }
    return ""
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetSegmentPath() string {
    return "flags"
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flag" {
        for _, c := range flags.Flag {
            if flags.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_OirHistory_Flags_Flag{}
        flags.Flag = append(flags.Flag, child)
        return &flags.Flag[len(flags.Flag)-1]
    }
    return nil
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range flags.Flag {
        children[flags.Flag[i].GetSegmentPath()] = &flags.Flag[i]
    }
    return children
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetBundleName() string { return "cisco_ios_xr" }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetYangName() string { return "flags" }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flags *Fia_Nodes_Node_OirHistory_Flags) SetParent(parent types.Entity) { flags.parent = parent }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetParent() types.Entity { return flags.parent }

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetParentYangName() string { return "oir-history" }

// Fia_Nodes_Node_OirHistory_Flags_Flag
// Flag value for physical location
type Fia_Nodes_Node_OirHistory_Flags_Flag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Flag value. The type is interface{} with range:
    // -2147483648..2147483647.
    Flag interface{}

    // Slot table for history.
    Slots Fia_Nodes_Node_OirHistory_Flags_Flag_Slots
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetFilter() yfilter.YFilter { return flag.YFilter }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) SetFilter(yf yfilter.YFilter) { flag.YFilter = yf }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetGoName(yname string) string {
    if yname == "flag" { return "Flag" }
    if yname == "slots" { return "Slots" }
    return ""
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetSegmentPath() string {
    return "flag" + "[flag='" + fmt.Sprintf("%v", flag.Flag) + "']"
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slots" {
        return &flag.Slots
    }
    return nil
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slots"] = &flag.Slots
    return children
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["flag"] = flag.Flag
    return leafs
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetBundleName() string { return "cisco_ios_xr" }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetYangName() string { return "flag" }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) SetParent(parent types.Entity) { flag.parent = parent }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetParent() types.Entity { return flag.parent }

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetParentYangName() string { return "flags" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots
// Slot table for history
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Slot number for getting history. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot.
    Slot []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetFilter() yfilter.YFilter { return slots.YFilter }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) SetFilter(yf yfilter.YFilter) { slots.YFilter = yf }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    return ""
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetSegmentPath() string {
    return "slots"
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slot" {
        for _, c := range slots.Slot {
            if slots.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot{}
        slots.Slot = append(slots.Slot, child)
        return &slots.Slot[len(slots.Slot)-1]
    }
    return nil
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slots.Slot {
        children[slots.Slot[i].GetSegmentPath()] = &slots.Slot[i]
    }
    return children
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetBundleName() string { return "cisco_ios_xr" }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetYangName() string { return "slots" }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) SetParent(parent types.Entity) { slots.parent = parent }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetParent() types.Entity { return slots.parent }

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetParentYangName() string { return "flag" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot
// Slot number for getting history
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Slot number. The type is interface{} with range:
    // -2147483648..2147483647.
    Slot interface{}

    // drv version. The type is interface{} with range: 0..4294967295.
    DrvVersion interface{}

    // coeff major rev. The type is interface{} with range: 0..4294967295.
    CoeffMajorRev interface{}

    // coeff minor rev. The type is interface{} with range: 0..4294967295.
    CoeffMinorRev interface{}

    // functional role. The type is interface{} with range: 0..255.
    FunctionalRole interface{}

    // issu role. The type is interface{} with range: 0..255.
    IssuRole interface{}

    // node id. The type is string.
    NodeId interface{}

    // rack type. The type is interface{} with range: -2147483648..2147483647.
    RackType interface{}

    // rack num. The type is interface{} with range: 0..255.
    RackNum interface{}

    // is driver ready. The type is bool.
    IsDriverReady interface{}

    // card avail mask. The type is interface{} with range: 0..4294967295.
    CardAvailMask interface{}

    // asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicAvailMask interface{}

    // exp asic avail mask. The type is interface{} with range:
    // 0..18446744073709551615.
    ExpAsicAvailMask interface{}

    // ucmc ratio. The type is interface{} with range: 0..4294967295.
    UcmcRatio interface{}

    // asic oper notify to fsdb pending bmap. The type is interface{} with range:
    // 0..18446744073709551615.
    AsicOperNotifyToFsdbPendingBmap interface{}

    // is full fgid download req. The type is bool.
    IsFullFgidDownloadReq interface{}

    // is fgid download in progress. The type is bool.
    IsFgidDownloadInProgress interface{}

    // is fgid download completed. The type is bool.
    IsFgidDownloadCompleted interface{}

    // fsdb conn active. The type is bool.
    FsdbConnActive interface{}

    // fgid conn active. The type is bool.
    FgidConnActive interface{}

    // issu mgr conn active. The type is bool.
    IssuMgrConnActive interface{}

    // fsdb reg active. The type is bool.
    FsdbRegActive interface{}

    // fgid reg active. The type is bool.
    FgidRegActive interface{}

    // issu mgr reg active. The type is bool.
    IssuMgrRegActive interface{}

    // num pm conn reqs. The type is interface{} with range: 0..255.
    NumPmConnReqs interface{}

    // num fsdb conn reqs. The type is interface{} with range: 0..255.
    NumFsdbConnReqs interface{}

    // num fgid conn reqs. The type is interface{} with range: 0..255.
    NumFgidConnReqs interface{}

    // num fstats conn reqs. The type is interface{} with range: 0..255.
    NumFstatsConnReqs interface{}

    // num cm conn reqs. The type is interface{} with range: 0..255.
    NumCmConnReqs interface{}

    // num issu mgr conn reqs. The type is interface{} with range: 0..255.
    NumIssuMgrConnReqs interface{}

    // num peer fia conn reqs. The type is interface{} with range: 0..255.
    NumPeerFiaConnReqs interface{}

    // is gaspp registered. The type is bool.
    IsGasppRegistered interface{}

    // is cih registered. The type is bool.
    IsCihRegistered interface{}

    // drvr initial startup timestamp. The type is string.
    DrvrInitialStartupTimestamp interface{}

    // drvr current startup timestamp. The type is string.
    DrvrCurrentStartupTimestamp interface{}

    // num intf ports. The type is interface{} with range: 0..4294967295.
    NumIntfPorts interface{}

    // uc weight. The type is interface{} with range: 0..255.
    UcWeight interface{}

    // respawn count. The type is interface{} with range: 0..255.
    RespawnCount interface{}

    // total asics. The type is interface{} with range: 0..255.
    TotalAsics interface{}

    // issu ready ntfy pending. The type is bool.
    IssuReadyNtfyPending interface{}

    // issu abort sent. The type is bool.
    IssuAbortSent interface{}

    // issu abort rcvd. The type is bool.
    IssuAbortRcvd interface{}

    // fabric mode. The type is interface{} with range: 0..255.
    FabricMode interface{}

    // FC Mode. The type is FcMode.
    FcMode interface{}

    // board rev id. The type is interface{} with range: 0..4294967295.
    BoardRevId interface{}

    // device info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo.
    DeviceInfo []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo

    // card info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo.
    CardInfo []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetFilter() yfilter.YFilter { return slot.YFilter }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) SetFilter(yf yfilter.YFilter) { slot.YFilter = yf }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetGoName(yname string) string {
    if yname == "slot" { return "Slot" }
    if yname == "drv-version" { return "DrvVersion" }
    if yname == "coeff-major-rev" { return "CoeffMajorRev" }
    if yname == "coeff-minor-rev" { return "CoeffMinorRev" }
    if yname == "functional-role" { return "FunctionalRole" }
    if yname == "issu-role" { return "IssuRole" }
    if yname == "node-id" { return "NodeId" }
    if yname == "rack-type" { return "RackType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "is-driver-ready" { return "IsDriverReady" }
    if yname == "card-avail-mask" { return "CardAvailMask" }
    if yname == "asic-avail-mask" { return "AsicAvailMask" }
    if yname == "exp-asic-avail-mask" { return "ExpAsicAvailMask" }
    if yname == "ucmc-ratio" { return "UcmcRatio" }
    if yname == "asic-oper-notify-to-fsdb-pending-bmap" { return "AsicOperNotifyToFsdbPendingBmap" }
    if yname == "is-full-fgid-download-req" { return "IsFullFgidDownloadReq" }
    if yname == "is-fgid-download-in-progress" { return "IsFgidDownloadInProgress" }
    if yname == "is-fgid-download-completed" { return "IsFgidDownloadCompleted" }
    if yname == "fsdb-conn-active" { return "FsdbConnActive" }
    if yname == "fgid-conn-active" { return "FgidConnActive" }
    if yname == "issu-mgr-conn-active" { return "IssuMgrConnActive" }
    if yname == "fsdb-reg-active" { return "FsdbRegActive" }
    if yname == "fgid-reg-active" { return "FgidRegActive" }
    if yname == "issu-mgr-reg-active" { return "IssuMgrRegActive" }
    if yname == "num-pm-conn-reqs" { return "NumPmConnReqs" }
    if yname == "num-fsdb-conn-reqs" { return "NumFsdbConnReqs" }
    if yname == "num-fgid-conn-reqs" { return "NumFgidConnReqs" }
    if yname == "num-fstats-conn-reqs" { return "NumFstatsConnReqs" }
    if yname == "num-cm-conn-reqs" { return "NumCmConnReqs" }
    if yname == "num-issu-mgr-conn-reqs" { return "NumIssuMgrConnReqs" }
    if yname == "num-peer-fia-conn-reqs" { return "NumPeerFiaConnReqs" }
    if yname == "is-gaspp-registered" { return "IsGasppRegistered" }
    if yname == "is-cih-registered" { return "IsCihRegistered" }
    if yname == "drvr-initial-startup-timestamp" { return "DrvrInitialStartupTimestamp" }
    if yname == "drvr-current-startup-timestamp" { return "DrvrCurrentStartupTimestamp" }
    if yname == "num-intf-ports" { return "NumIntfPorts" }
    if yname == "uc-weight" { return "UcWeight" }
    if yname == "respawn-count" { return "RespawnCount" }
    if yname == "total-asics" { return "TotalAsics" }
    if yname == "issu-ready-ntfy-pending" { return "IssuReadyNtfyPending" }
    if yname == "issu-abort-sent" { return "IssuAbortSent" }
    if yname == "issu-abort-rcvd" { return "IssuAbortRcvd" }
    if yname == "fabric-mode" { return "FabricMode" }
    if yname == "fc-mode" { return "FcMode" }
    if yname == "board-rev-id" { return "BoardRevId" }
    if yname == "device-info" { return "DeviceInfo" }
    if yname == "card-info" { return "CardInfo" }
    return ""
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetSegmentPath() string {
    return "slot" + "[slot='" + fmt.Sprintf("%v", slot.Slot) + "']"
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "device-info" {
        for _, c := range slot.DeviceInfo {
            if slot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo{}
        slot.DeviceInfo = append(slot.DeviceInfo, child)
        return &slot.DeviceInfo[len(slot.DeviceInfo)-1]
    }
    if childYangName == "card-info" {
        for _, c := range slot.CardInfo {
            if slot.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo{}
        slot.CardInfo = append(slot.CardInfo, child)
        return &slot.CardInfo[len(slot.CardInfo)-1]
    }
    return nil
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slot.DeviceInfo {
        children[slot.DeviceInfo[i].GetSegmentPath()] = &slot.DeviceInfo[i]
    }
    for i := range slot.CardInfo {
        children[slot.CardInfo[i].GetSegmentPath()] = &slot.CardInfo[i]
    }
    return children
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slot"] = slot.Slot
    leafs["drv-version"] = slot.DrvVersion
    leafs["coeff-major-rev"] = slot.CoeffMajorRev
    leafs["coeff-minor-rev"] = slot.CoeffMinorRev
    leafs["functional-role"] = slot.FunctionalRole
    leafs["issu-role"] = slot.IssuRole
    leafs["node-id"] = slot.NodeId
    leafs["rack-type"] = slot.RackType
    leafs["rack-num"] = slot.RackNum
    leafs["is-driver-ready"] = slot.IsDriverReady
    leafs["card-avail-mask"] = slot.CardAvailMask
    leafs["asic-avail-mask"] = slot.AsicAvailMask
    leafs["exp-asic-avail-mask"] = slot.ExpAsicAvailMask
    leafs["ucmc-ratio"] = slot.UcmcRatio
    leafs["asic-oper-notify-to-fsdb-pending-bmap"] = slot.AsicOperNotifyToFsdbPendingBmap
    leafs["is-full-fgid-download-req"] = slot.IsFullFgidDownloadReq
    leafs["is-fgid-download-in-progress"] = slot.IsFgidDownloadInProgress
    leafs["is-fgid-download-completed"] = slot.IsFgidDownloadCompleted
    leafs["fsdb-conn-active"] = slot.FsdbConnActive
    leafs["fgid-conn-active"] = slot.FgidConnActive
    leafs["issu-mgr-conn-active"] = slot.IssuMgrConnActive
    leafs["fsdb-reg-active"] = slot.FsdbRegActive
    leafs["fgid-reg-active"] = slot.FgidRegActive
    leafs["issu-mgr-reg-active"] = slot.IssuMgrRegActive
    leafs["num-pm-conn-reqs"] = slot.NumPmConnReqs
    leafs["num-fsdb-conn-reqs"] = slot.NumFsdbConnReqs
    leafs["num-fgid-conn-reqs"] = slot.NumFgidConnReqs
    leafs["num-fstats-conn-reqs"] = slot.NumFstatsConnReqs
    leafs["num-cm-conn-reqs"] = slot.NumCmConnReqs
    leafs["num-issu-mgr-conn-reqs"] = slot.NumIssuMgrConnReqs
    leafs["num-peer-fia-conn-reqs"] = slot.NumPeerFiaConnReqs
    leafs["is-gaspp-registered"] = slot.IsGasppRegistered
    leafs["is-cih-registered"] = slot.IsCihRegistered
    leafs["drvr-initial-startup-timestamp"] = slot.DrvrInitialStartupTimestamp
    leafs["drvr-current-startup-timestamp"] = slot.DrvrCurrentStartupTimestamp
    leafs["num-intf-ports"] = slot.NumIntfPorts
    leafs["uc-weight"] = slot.UcWeight
    leafs["respawn-count"] = slot.RespawnCount
    leafs["total-asics"] = slot.TotalAsics
    leafs["issu-ready-ntfy-pending"] = slot.IssuReadyNtfyPending
    leafs["issu-abort-sent"] = slot.IssuAbortSent
    leafs["issu-abort-rcvd"] = slot.IssuAbortRcvd
    leafs["fabric-mode"] = slot.FabricMode
    leafs["fc-mode"] = slot.FcMode
    leafs["board-rev-id"] = slot.BoardRevId
    return leafs
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetBundleName() string { return "cisco_ios_xr" }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetYangName() string { return "slot" }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) SetParent(parent types.Entity) { slot.parent = parent }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetParent() types.Entity { return slot.parent }

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetParentYangName() string { return "slots" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo
// device info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // is valid. The type is bool.
    IsValid interface{}

    // fapid. The type is interface{} with range: 0..4294967295.
    Fapid interface{}

    // hotplug event. The type is interface{} with range: 0..4294967295.
    HotplugEvent interface{}

    // Slice State. The type is SliceState.
    SliceState interface{}

    // Admin State. The type is AdminState.
    AdminState interface{}

    // Oper State. The type is AsicOperState.
    OperState interface{}

    // Asic State. The type is AsicAccessState.
    AsicState interface{}

    // last init cause. The type is AsicInitMethod.
    LastInitCause interface{}

    // num pon resets. The type is interface{} with range: 0..4294967295.
    NumPonResets interface{}

    // num hard resets. The type is interface{} with range: 0..4294967295.
    NumHardResets interface{}

    // local switch state. The type is bool.
    LocalSwitchState interface{}

    // asic id.
    AsicId Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetFilter() yfilter.YFilter { return deviceInfo.YFilter }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) SetFilter(yf yfilter.YFilter) { deviceInfo.YFilter = yf }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetGoName(yname string) string {
    if yname == "is-valid" { return "IsValid" }
    if yname == "fapid" { return "Fapid" }
    if yname == "hotplug-event" { return "HotplugEvent" }
    if yname == "slice-state" { return "SliceState" }
    if yname == "admin-state" { return "AdminState" }
    if yname == "oper-state" { return "OperState" }
    if yname == "asic-state" { return "AsicState" }
    if yname == "last-init-cause" { return "LastInitCause" }
    if yname == "num-pon-resets" { return "NumPonResets" }
    if yname == "num-hard-resets" { return "NumHardResets" }
    if yname == "local-switch-state" { return "LocalSwitchState" }
    if yname == "asic-id" { return "AsicId" }
    return ""
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetSegmentPath() string {
    return "device-info"
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "asic-id" {
        return &deviceInfo.AsicId
    }
    return nil
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["asic-id"] = &deviceInfo.AsicId
    return children
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-valid"] = deviceInfo.IsValid
    leafs["fapid"] = deviceInfo.Fapid
    leafs["hotplug-event"] = deviceInfo.HotplugEvent
    leafs["slice-state"] = deviceInfo.SliceState
    leafs["admin-state"] = deviceInfo.AdminState
    leafs["oper-state"] = deviceInfo.OperState
    leafs["asic-state"] = deviceInfo.AsicState
    leafs["last-init-cause"] = deviceInfo.LastInitCause
    leafs["num-pon-resets"] = deviceInfo.NumPonResets
    leafs["num-hard-resets"] = deviceInfo.NumHardResets
    leafs["local-switch-state"] = deviceInfo.LocalSwitchState
    return leafs
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetBundleName() string { return "cisco_ios_xr" }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetYangName() string { return "device-info" }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) SetParent(parent types.Entity) { deviceInfo.parent = parent }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetParent() types.Entity { return deviceInfo.parent }

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetParentYangName() string { return "slot" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId
// asic id
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rack Type. The type is Rack.
    RackType interface{}

    // Asic Type. The type is Asic.
    AsicType interface{}

    // rack num. The type is interface{} with range: 0..4294967295.
    RackNum interface{}

    // slot num. The type is interface{} with range: 0..4294967295.
    SlotNum interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetFilter() yfilter.YFilter { return asicId.YFilter }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) SetFilter(yf yfilter.YFilter) { asicId.YFilter = yf }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetGoName(yname string) string {
    if yname == "rack-type" { return "RackType" }
    if yname == "asic-type" { return "AsicType" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "slot-num" { return "SlotNum" }
    if yname == "asic-instance" { return "AsicInstance" }
    return ""
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetSegmentPath() string {
    return "asic-id"
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rack-type"] = asicId.RackType
    leafs["asic-type"] = asicId.AsicType
    leafs["rack-num"] = asicId.RackNum
    leafs["slot-num"] = asicId.SlotNum
    leafs["asic-instance"] = asicId.AsicInstance
    return leafs
}

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetBundleName() string { return "cisco_ios_xr" }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetYangName() string { return "asic-id" }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) SetParent(parent types.Entity) { asicId.parent = parent }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetParent() types.Entity { return asicId.parent }

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetParentYangName() string { return "device-info" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo
// card info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // card name. The type is string.
    CardName interface{}

    // slot no. The type is interface{} with range: -2147483648..2147483647.
    SlotNo interface{}

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // card state. The type is interface{} with range: 0..255.
    CardState interface{}

    // exp num asics. The type is interface{} with range: 0..4294967295.
    ExpNumAsics interface{}

    // exp num asics per fsdb. The type is interface{} with range: 0..4294967295.
    ExpNumAsicsPerFsdb interface{}

    // is powered. The type is bool.
    IsPowered interface{}

    // cxp avail bitmap. The type is interface{} with range:
    // 0..18446744073709551615.
    CxpAvailBitmap interface{}

    // num ilkns per asic. The type is interface{} with range: 0..4294967295.
    NumIlknsPerAsic interface{}

    // num local ports per ilkn. The type is interface{} with range:
    // 0..4294967295.
    NumLocalPortsPerIlkn interface{}

    // num cos per port. The type is interface{} with range: 0..255.
    NumCosPerPort interface{}

    // oir circular buffer.
    OirCircularBuffer Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetFilter() yfilter.YFilter { return cardInfo.YFilter }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) SetFilter(yf yfilter.YFilter) { cardInfo.YFilter = yf }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetGoName(yname string) string {
    if yname == "card-type" { return "CardType" }
    if yname == "card-name" { return "CardName" }
    if yname == "slot-no" { return "SlotNo" }
    if yname == "card-flag" { return "CardFlag" }
    if yname == "evt-flag" { return "EvtFlag" }
    if yname == "reg-flag" { return "RegFlag" }
    if yname == "instance" { return "Instance" }
    if yname == "card-state" { return "CardState" }
    if yname == "exp-num-asics" { return "ExpNumAsics" }
    if yname == "exp-num-asics-per-fsdb" { return "ExpNumAsicsPerFsdb" }
    if yname == "is-powered" { return "IsPowered" }
    if yname == "cxp-avail-bitmap" { return "CxpAvailBitmap" }
    if yname == "num-ilkns-per-asic" { return "NumIlknsPerAsic" }
    if yname == "num-local-ports-per-ilkn" { return "NumLocalPortsPerIlkn" }
    if yname == "num-cos-per-port" { return "NumCosPerPort" }
    if yname == "oir-circular-buffer" { return "OirCircularBuffer" }
    return ""
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetSegmentPath() string {
    return "card-info"
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oir-circular-buffer" {
        return &cardInfo.OirCircularBuffer
    }
    return nil
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["oir-circular-buffer"] = &cardInfo.OirCircularBuffer
    return children
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-type"] = cardInfo.CardType
    leafs["card-name"] = cardInfo.CardName
    leafs["slot-no"] = cardInfo.SlotNo
    leafs["card-flag"] = cardInfo.CardFlag
    leafs["evt-flag"] = cardInfo.EvtFlag
    leafs["reg-flag"] = cardInfo.RegFlag
    leafs["instance"] = cardInfo.Instance
    leafs["card-state"] = cardInfo.CardState
    leafs["exp-num-asics"] = cardInfo.ExpNumAsics
    leafs["exp-num-asics-per-fsdb"] = cardInfo.ExpNumAsicsPerFsdb
    leafs["is-powered"] = cardInfo.IsPowered
    leafs["cxp-avail-bitmap"] = cardInfo.CxpAvailBitmap
    leafs["num-ilkns-per-asic"] = cardInfo.NumIlknsPerAsic
    leafs["num-local-ports-per-ilkn"] = cardInfo.NumLocalPortsPerIlkn
    leafs["num-cos-per-port"] = cardInfo.NumCosPerPort
    return leafs
}

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetBundleName() string { return "cisco_ios_xr" }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetYangName() string { return "card-info" }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) SetParent(parent types.Entity) { cardInfo.parent = parent }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetParent() types.Entity { return cardInfo.parent }

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetParentYangName() string { return "slot" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer
// oir circular buffer
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // count. The type is interface{} with range: -2147483648..2147483647.
    Count interface{}

    // start. The type is interface{} with range: -2147483648..2147483647.
    Start interface{}

    // end. The type is interface{} with range: -2147483648..2147483647.
    End interface{}

    // fia oir info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo.
    FiaOirInfo []Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetFilter() yfilter.YFilter { return oirCircularBuffer.YFilter }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) SetFilter(yf yfilter.YFilter) { oirCircularBuffer.YFilter = yf }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetGoName(yname string) string {
    if yname == "count" { return "Count" }
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "fia-oir-info" { return "FiaOirInfo" }
    return ""
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetSegmentPath() string {
    return "oir-circular-buffer"
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fia-oir-info" {
        for _, c := range oirCircularBuffer.FiaOirInfo {
            if oirCircularBuffer.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo{}
        oirCircularBuffer.FiaOirInfo = append(oirCircularBuffer.FiaOirInfo, child)
        return &oirCircularBuffer.FiaOirInfo[len(oirCircularBuffer.FiaOirInfo)-1]
    }
    return nil
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range oirCircularBuffer.FiaOirInfo {
        children[oirCircularBuffer.FiaOirInfo[i].GetSegmentPath()] = &oirCircularBuffer.FiaOirInfo[i]
    }
    return children
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["count"] = oirCircularBuffer.Count
    leafs["start"] = oirCircularBuffer.Start
    leafs["end"] = oirCircularBuffer.End
    return leafs
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetYangName() string { return "oir-circular-buffer" }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) SetParent(parent types.Entity) { oirCircularBuffer.parent = parent }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetParent() types.Entity { return oirCircularBuffer.parent }

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetParentYangName() string { return "card-info" }

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo
// fia oir info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // card flag. The type is interface{} with range: -2147483648..2147483647.
    CardFlag interface{}

    // card type. The type is interface{} with range: -2147483648..2147483647.
    CardType interface{}

    // reg flag. The type is interface{} with range: -2147483648..2147483647.
    RegFlag interface{}

    // evt flag. The type is interface{} with range: -2147483648..2147483647.
    EvtFlag interface{}

    // rack num. The type is interface{} with range: -2147483648..2147483647.
    RackNum interface{}

    // instance. The type is interface{} with range: -2147483648..2147483647.
    Instance interface{}

    // cur card state. The type is interface{} with range:
    // -2147483648..2147483647.
    CurCardState interface{}
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetFilter() yfilter.YFilter { return fiaOirInfo.YFilter }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) SetFilter(yf yfilter.YFilter) { fiaOirInfo.YFilter = yf }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetGoName(yname string) string {
    if yname == "card-flag" { return "CardFlag" }
    if yname == "card-type" { return "CardType" }
    if yname == "reg-flag" { return "RegFlag" }
    if yname == "evt-flag" { return "EvtFlag" }
    if yname == "rack-num" { return "RackNum" }
    if yname == "instance" { return "Instance" }
    if yname == "cur-card-state" { return "CurCardState" }
    return ""
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetSegmentPath() string {
    return "fia-oir-info"
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["card-flag"] = fiaOirInfo.CardFlag
    leafs["card-type"] = fiaOirInfo.CardType
    leafs["reg-flag"] = fiaOirInfo.RegFlag
    leafs["evt-flag"] = fiaOirInfo.EvtFlag
    leafs["rack-num"] = fiaOirInfo.RackNum
    leafs["instance"] = fiaOirInfo.Instance
    leafs["cur-card-state"] = fiaOirInfo.CurCardState
    return leafs
}

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetYangName() string { return "fia-oir-info" }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) SetParent(parent types.Entity) { fiaOirInfo.parent = parent }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetParent() types.Entity { return fiaOirInfo.parent }

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetParentYangName() string { return "oir-circular-buffer" }

// Fia_Nodes_Node_AsicStatistics
// FIA asic statistics information
type Fia_Nodes_Node_AsicStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Instance table for statistics.
    StatisticsAsicInstances Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetFilter() yfilter.YFilter { return asicStatistics.YFilter }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) SetFilter(yf yfilter.YFilter) { asicStatistics.YFilter = yf }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetGoName(yname string) string {
    if yname == "statistics-asic-instances" { return "StatisticsAsicInstances" }
    return ""
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetSegmentPath() string {
    return "asic-statistics"
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-asic-instances" {
        return &asicStatistics.StatisticsAsicInstances
    }
    return nil
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics-asic-instances"] = &asicStatistics.StatisticsAsicInstances
    return children
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetYangName() string { return "asic-statistics" }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) SetParent(parent types.Entity) { asicStatistics.parent = parent }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetParent() types.Entity { return asicStatistics.parent }

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetParentYangName() string { return "node" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances
// Instance table for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Asic instance for statistics. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance.
    StatisticsAsicInstance []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetFilter() yfilter.YFilter { return statisticsAsicInstances.YFilter }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) SetFilter(yf yfilter.YFilter) { statisticsAsicInstances.YFilter = yf }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetGoName(yname string) string {
    if yname == "statistics-asic-instance" { return "StatisticsAsicInstance" }
    return ""
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetSegmentPath() string {
    return "statistics-asic-instances"
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics-asic-instance" {
        for _, c := range statisticsAsicInstances.StatisticsAsicInstance {
            if statisticsAsicInstances.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance{}
        statisticsAsicInstances.StatisticsAsicInstance = append(statisticsAsicInstances.StatisticsAsicInstance, child)
        return &statisticsAsicInstances.StatisticsAsicInstance[len(statisticsAsicInstances.StatisticsAsicInstance)-1]
    }
    return nil
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticsAsicInstances.StatisticsAsicInstance {
        children[statisticsAsicInstances.StatisticsAsicInstance[i].GetSegmentPath()] = &statisticsAsicInstances.StatisticsAsicInstance[i]
    }
    return children
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetYangName() string { return "statistics-asic-instances" }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) SetParent(parent types.Entity) { statisticsAsicInstances.parent = parent }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetParent() types.Entity { return statisticsAsicInstances.parent }

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetParentYangName() string { return "asic-statistics" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance
// Asic instance for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Asic instance. The type is interface{} with range:
    // 0..255.
    Instance interface{}

    // Packet Byte Counter for a Asic.
    PbcStatistics Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics

    // Statistics of FMAC.
    FmacStatistics Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetFilter() yfilter.YFilter { return statisticsAsicInstance.YFilter }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) SetFilter(yf yfilter.YFilter) { statisticsAsicInstance.YFilter = yf }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetGoName(yname string) string {
    if yname == "instance" { return "Instance" }
    if yname == "pbc-statistics" { return "PbcStatistics" }
    if yname == "fmac-statistics" { return "FmacStatistics" }
    return ""
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetSegmentPath() string {
    return "statistics-asic-instance" + "[instance='" + fmt.Sprintf("%v", statisticsAsicInstance.Instance) + "']"
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbc-statistics" {
        return &statisticsAsicInstance.PbcStatistics
    }
    if childYangName == "fmac-statistics" {
        return &statisticsAsicInstance.FmacStatistics
    }
    return nil
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbc-statistics"] = &statisticsAsicInstance.PbcStatistics
    children["fmac-statistics"] = &statisticsAsicInstance.FmacStatistics
    return children
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance"] = statisticsAsicInstance.Instance
    return leafs
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetBundleName() string { return "cisco_ios_xr" }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetYangName() string { return "statistics-asic-instance" }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) SetParent(parent types.Entity) { statisticsAsicInstance.parent = parent }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetParent() types.Entity { return statisticsAsicInstance.parent }

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetParentYangName() string { return "statistics-asic-instances" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics
// Packet Byte Counter for a Asic
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PBC stats bag.
    PbcStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetFilter() yfilter.YFilter { return pbcStatistics.YFilter }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) SetFilter(yf yfilter.YFilter) { pbcStatistics.YFilter = yf }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetGoName(yname string) string {
    if yname == "pbc-stats" { return "PbcStats" }
    return ""
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetSegmentPath() string {
    return "pbc-statistics"
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pbc-stats" {
        return &pbcStatistics.PbcStats
    }
    return nil
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pbc-stats"] = &pbcStatistics.PbcStats
    return children
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetYangName() string { return "pbc-statistics" }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) SetParent(parent types.Entity) { pbcStatistics.parent = parent }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetParent() types.Entity { return pbcStatistics.parent }

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetParentYangName() string { return "statistics-asic-instance" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats
// PBC stats bag
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // valid. The type is bool.
    Valid interface{}

    // rack no. The type is interface{} with range: 0..4294967295.
    RackNo interface{}

    // slot no. The type is interface{} with range: 0..4294967295.
    SlotNo interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // chip ver. The type is interface{} with range: 0..65535.
    ChipVer interface{}

    // stats info.
    StatsInfo Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetFilter() yfilter.YFilter { return pbcStats.YFilter }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) SetFilter(yf yfilter.YFilter) { pbcStats.YFilter = yf }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetGoName(yname string) string {
    if yname == "valid" { return "Valid" }
    if yname == "rack-no" { return "RackNo" }
    if yname == "slot-no" { return "SlotNo" }
    if yname == "asic-instance" { return "AsicInstance" }
    if yname == "chip-ver" { return "ChipVer" }
    if yname == "stats-info" { return "StatsInfo" }
    return ""
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetSegmentPath() string {
    return "pbc-stats"
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats-info" {
        return &pbcStats.StatsInfo
    }
    return nil
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats-info"] = &pbcStats.StatsInfo
    return children
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["valid"] = pbcStats.Valid
    leafs["rack-no"] = pbcStats.RackNo
    leafs["slot-no"] = pbcStats.SlotNo
    leafs["asic-instance"] = pbcStats.AsicInstance
    leafs["chip-ver"] = pbcStats.ChipVer
    return leafs
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetBundleName() string { return "cisco_ios_xr" }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetYangName() string { return "pbc-stats" }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) SetParent(parent types.Entity) { pbcStats.parent = parent }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetParent() types.Entity { return pbcStats.parent }

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetParentYangName() string { return "pbc-statistics" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo
// stats info
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Num Blocks. The type is interface{} with range: 0..255.
    NumBlocks interface{}

    // block info. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo.
    BlockInfo []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetFilter() yfilter.YFilter { return statsInfo.YFilter }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) SetFilter(yf yfilter.YFilter) { statsInfo.YFilter = yf }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetGoName(yname string) string {
    if yname == "num-blocks" { return "NumBlocks" }
    if yname == "block-info" { return "BlockInfo" }
    return ""
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetSegmentPath() string {
    return "stats-info"
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "block-info" {
        for _, c := range statsInfo.BlockInfo {
            if statsInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo{}
        statsInfo.BlockInfo = append(statsInfo.BlockInfo, child)
        return &statsInfo.BlockInfo[len(statsInfo.BlockInfo)-1]
    }
    return nil
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statsInfo.BlockInfo {
        children[statsInfo.BlockInfo[i].GetSegmentPath()] = &statsInfo.BlockInfo[i]
    }
    return children
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-blocks"] = statsInfo.NumBlocks
    return leafs
}

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetYangName() string { return "stats-info" }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) SetParent(parent types.Entity) { statsInfo.parent = parent }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetParent() types.Entity { return statsInfo.parent }

func (statsInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo) GetParentYangName() string { return "pbc-stats" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo
// block info
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Block Name. The type is string with length: 0..10.
    BlockName interface{}

    // Num Fields. The type is interface{} with range: 0..255.
    NumFields interface{}

    // field info. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo.
    FieldInfo []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetFilter() yfilter.YFilter { return blockInfo.YFilter }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) SetFilter(yf yfilter.YFilter) { blockInfo.YFilter = yf }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetGoName(yname string) string {
    if yname == "block-name" { return "BlockName" }
    if yname == "num-fields" { return "NumFields" }
    if yname == "field-info" { return "FieldInfo" }
    return ""
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetSegmentPath() string {
    return "block-info"
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "field-info" {
        for _, c := range blockInfo.FieldInfo {
            if blockInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo{}
        blockInfo.FieldInfo = append(blockInfo.FieldInfo, child)
        return &blockInfo.FieldInfo[len(blockInfo.FieldInfo)-1]
    }
    return nil
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range blockInfo.FieldInfo {
        children[blockInfo.FieldInfo[i].GetSegmentPath()] = &blockInfo.FieldInfo[i]
    }
    return children
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["block-name"] = blockInfo.BlockName
    leafs["num-fields"] = blockInfo.NumFields
    return leafs
}

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetBundleName() string { return "cisco_ios_xr" }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetYangName() string { return "block-info" }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) SetParent(parent types.Entity) { blockInfo.parent = parent }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetParent() types.Entity { return blockInfo.parent }

func (blockInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo) GetParentYangName() string { return "stats-info" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo
// field info
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Field Name. The type is string with length: 0..80.
    FieldName interface{}

    // Field Value. The type is interface{} with range: 0..18446744073709551615.
    FieldValue interface{}

    // Is Ovf. The type is bool.
    IsOvf interface{}
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetFilter() yfilter.YFilter { return fieldInfo.YFilter }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) SetFilter(yf yfilter.YFilter) { fieldInfo.YFilter = yf }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetGoName(yname string) string {
    if yname == "field-name" { return "FieldName" }
    if yname == "field-value" { return "FieldValue" }
    if yname == "is-ovf" { return "IsOvf" }
    return ""
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetSegmentPath() string {
    return "field-info"
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["field-name"] = fieldInfo.FieldName
    leafs["field-value"] = fieldInfo.FieldValue
    leafs["is-ovf"] = fieldInfo.IsOvf
    return leafs
}

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetBundleName() string { return "cisco_ios_xr" }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetYangName() string { return "field-info" }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) SetParent(parent types.Entity) { fieldInfo.parent = parent }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetParent() types.Entity { return fieldInfo.parent }

func (fieldInfo *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_StatsInfo_BlockInfo_FieldInfo) GetParentYangName() string { return "block-info" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics
// Statistics of FMAC
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link table for statistics.
    FmacLinks Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetFilter() yfilter.YFilter { return fmacStatistics.YFilter }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) SetFilter(yf yfilter.YFilter) { fmacStatistics.YFilter = yf }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetGoName(yname string) string {
    if yname == "fmac-links" { return "FmacLinks" }
    return ""
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetSegmentPath() string {
    return "fmac-statistics"
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fmac-links" {
        return &fmacStatistics.FmacLinks
    }
    return nil
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["fmac-links"] = &fmacStatistics.FmacLinks
    return children
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetYangName() string { return "fmac-statistics" }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) SetParent(parent types.Entity) { fmacStatistics.parent = parent }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetParent() types.Entity { return fmacStatistics.parent }

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetParentYangName() string { return "statistics-asic-instance" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks
// Link table for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Link number for statistics. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink.
    FmacLink []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetFilter() yfilter.YFilter { return fmacLinks.YFilter }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) SetFilter(yf yfilter.YFilter) { fmacLinks.YFilter = yf }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetGoName(yname string) string {
    if yname == "fmac-link" { return "FmacLink" }
    return ""
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetSegmentPath() string {
    return "fmac-links"
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fmac-link" {
        for _, c := range fmacLinks.FmacLink {
            if fmacLinks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink{}
        fmacLinks.FmacLink = append(fmacLinks.FmacLink, child)
        return &fmacLinks.FmacLink[len(fmacLinks.FmacLink)-1]
    }
    return nil
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fmacLinks.FmacLink {
        children[fmacLinks.FmacLink[i].GetSegmentPath()] = &fmacLinks.FmacLink[i]
    }
    return children
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetBundleName() string { return "cisco_ios_xr" }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetYangName() string { return "fmac-links" }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) SetParent(parent types.Entity) { fmacLinks.parent = parent }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetParent() types.Entity { return fmacLinks.parent }

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetParentYangName() string { return "fmac-statistics" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink
// Link number for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Link number. The type is interface{} with range:
    // -2147483648..2147483647.
    Link interface{}

    // Single aisc information. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic.
    FmacAsic []Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetFilter() yfilter.YFilter { return fmacLink.YFilter }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) SetFilter(yf yfilter.YFilter) { fmacLink.YFilter = yf }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetGoName(yname string) string {
    if yname == "link" { return "Link" }
    if yname == "fmac-asic" { return "FmacAsic" }
    return ""
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetSegmentPath() string {
    return "fmac-link" + "[link='" + fmt.Sprintf("%v", fmacLink.Link) + "']"
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "fmac-asic" {
        for _, c := range fmacLink.FmacAsic {
            if fmacLink.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic{}
        fmacLink.FmacAsic = append(fmacLink.FmacAsic, child)
        return &fmacLink.FmacAsic[len(fmacLink.FmacAsic)-1]
    }
    return nil
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range fmacLink.FmacAsic {
        children[fmacLink.FmacAsic[i].GetSegmentPath()] = &fmacLink.FmacAsic[i]
    }
    return children
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link"] = fmacLink.Link
    return leafs
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetBundleName() string { return "cisco_ios_xr" }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetYangName() string { return "fmac-link" }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) SetParent(parent types.Entity) { fmacLink.parent = parent }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetParent() types.Entity { return fmacLink.parent }

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetParentYangName() string { return "fmac-links" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic
// Single aisc information
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Single asic. The type is interface{} with range:
    // -2147483648..2147483647.
    Asic interface{}

    // valid. The type is bool.
    Valid interface{}

    // rack no. The type is interface{} with range: 0..4294967295.
    RackNo interface{}

    // slot no. The type is interface{} with range: 0..4294967295.
    SlotNo interface{}

    // asic instance. The type is interface{} with range: 0..4294967295.
    AsicInstance interface{}

    // link no. The type is interface{} with range: 0..4294967295.
    LinkNo interface{}

    // link valid. The type is bool.
    LinkValid interface{}

    // aggr stats.
    AggrStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats

    // incr stats.
    IncrStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetFilter() yfilter.YFilter { return fmacAsic.YFilter }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) SetFilter(yf yfilter.YFilter) { fmacAsic.YFilter = yf }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetGoName(yname string) string {
    if yname == "asic" { return "Asic" }
    if yname == "valid" { return "Valid" }
    if yname == "rack-no" { return "RackNo" }
    if yname == "slot-no" { return "SlotNo" }
    if yname == "asic-instance" { return "AsicInstance" }
    if yname == "link-no" { return "LinkNo" }
    if yname == "link-valid" { return "LinkValid" }
    if yname == "aggr-stats" { return "AggrStats" }
    if yname == "incr-stats" { return "IncrStats" }
    return ""
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetSegmentPath() string {
    return "fmac-asic" + "[asic='" + fmt.Sprintf("%v", fmacAsic.Asic) + "']"
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "aggr-stats" {
        return &fmacAsic.AggrStats
    }
    if childYangName == "incr-stats" {
        return &fmacAsic.IncrStats
    }
    return nil
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["aggr-stats"] = &fmacAsic.AggrStats
    children["incr-stats"] = &fmacAsic.IncrStats
    return children
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["asic"] = fmacAsic.Asic
    leafs["valid"] = fmacAsic.Valid
    leafs["rack-no"] = fmacAsic.RackNo
    leafs["slot-no"] = fmacAsic.SlotNo
    leafs["asic-instance"] = fmacAsic.AsicInstance
    leafs["link-no"] = fmacAsic.LinkNo
    leafs["link-valid"] = fmacAsic.LinkValid
    return leafs
}

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetBundleName() string { return "cisco_ios_xr" }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetYangName() string { return "fmac-asic" }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) SetParent(parent types.Entity) { fmacAsic.parent = parent }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetParent() types.Entity { return fmacAsic.parent }

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetParentYangName() string { return "fmac-link" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats
// aggr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // link error status.
    LinkErrorStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus

    // link counters.
    LinkCounters Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetFilter() yfilter.YFilter { return aggrStats.YFilter }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) SetFilter(yf yfilter.YFilter) { aggrStats.YFilter = yf }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetGoName(yname string) string {
    if yname == "link-error-status" { return "LinkErrorStatus" }
    if yname == "link-counters" { return "LinkCounters" }
    if yname == "ovf-status" { return "OvfStatus" }
    return ""
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetSegmentPath() string {
    return "aggr-stats"
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-error-status" {
        return &aggrStats.LinkErrorStatus
    }
    if childYangName == "link-counters" {
        return &aggrStats.LinkCounters
    }
    if childYangName == "ovf-status" {
        return &aggrStats.OvfStatus
    }
    return nil
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-error-status"] = &aggrStats.LinkErrorStatus
    children["link-counters"] = &aggrStats.LinkCounters
    children["ovf-status"] = &aggrStats.OvfStatus
    return children
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetBundleName() string { return "cisco_ios_xr" }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetYangName() string { return "aggr-stats" }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) SetParent(parent types.Entity) { aggrStats.parent = parent }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetParent() types.Entity { return aggrStats.parent }

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetParentYangName() string { return "fmac-asic" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus
// link error status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // link crc error. The type is interface{} with range: 0..4294967295.
    LinkCrcError interface{}

    // link size error. The type is interface{} with range: 0..4294967295.
    LinkSizeError interface{}

    // link mis align error. The type is interface{} with range: 0..4294967295.
    LinkMisAlignError interface{}

    // link code group error. The type is interface{} with range: 0..4294967295.
    LinkCodeGroupError interface{}

    // link no sig lock error. The type is interface{} with range: 0..4294967295.
    LinkNoSigLockError interface{}

    // link no sig accept error. The type is interface{} with range:
    // 0..4294967295.
    LinkNoSigAcceptError interface{}

    // link tokens error. The type is interface{} with range: 0..4294967295.
    LinkTokensError interface{}

    // error token count. The type is interface{} with range: 0..4294967295.
    ErrorTokenCount interface{}
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetFilter() yfilter.YFilter { return linkErrorStatus.YFilter }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) SetFilter(yf yfilter.YFilter) { linkErrorStatus.YFilter = yf }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetGoName(yname string) string {
    if yname == "link-crc-error" { return "LinkCrcError" }
    if yname == "link-size-error" { return "LinkSizeError" }
    if yname == "link-mis-align-error" { return "LinkMisAlignError" }
    if yname == "link-code-group-error" { return "LinkCodeGroupError" }
    if yname == "link-no-sig-lock-error" { return "LinkNoSigLockError" }
    if yname == "link-no-sig-accept-error" { return "LinkNoSigAcceptError" }
    if yname == "link-tokens-error" { return "LinkTokensError" }
    if yname == "error-token-count" { return "ErrorTokenCount" }
    return ""
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetSegmentPath() string {
    return "link-error-status"
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-crc-error"] = linkErrorStatus.LinkCrcError
    leafs["link-size-error"] = linkErrorStatus.LinkSizeError
    leafs["link-mis-align-error"] = linkErrorStatus.LinkMisAlignError
    leafs["link-code-group-error"] = linkErrorStatus.LinkCodeGroupError
    leafs["link-no-sig-lock-error"] = linkErrorStatus.LinkNoSigLockError
    leafs["link-no-sig-accept-error"] = linkErrorStatus.LinkNoSigAcceptError
    leafs["link-tokens-error"] = linkErrorStatus.LinkTokensError
    leafs["error-token-count"] = linkErrorStatus.ErrorTokenCount
    return leafs
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetBundleName() string { return "cisco_ios_xr" }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetYangName() string { return "link-error-status" }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) SetParent(parent types.Entity) { linkErrorStatus.parent = parent }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetParent() types.Entity { return linkErrorStatus.parent }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetParentYangName() string { return "aggr-stats" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters
// link counters
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BCodeErrors interface{}
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetFilter() yfilter.YFilter { return linkCounters.YFilter }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) SetFilter(yf yfilter.YFilter) { linkCounters.YFilter = yf }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetGoName(yname string) string {
    if yname == "tx-control-cells-counter" { return "TxControlCellsCounter" }
    if yname == "tx-data-cell-counter" { return "TxDataCellCounter" }
    if yname == "tx-data-byte-counter" { return "TxDataByteCounter" }
    if yname == "rx-crc-errors-counter" { return "RxCrcErrorsCounter" }
    if yname == "rx-lfec-fec-correctable-error" { return "RxLfecFecCorrectableError" }
    if yname == "rx-8b-10b-disparity-errors" { return "Rx8B10BDisparityErrors" }
    if yname == "rx-control-cells-counter" { return "RxControlCellsCounter" }
    if yname == "rx-data-cell-counter" { return "RxDataCellCounter" }
    if yname == "rx-data-byte-counter" { return "RxDataByteCounter" }
    if yname == "rx-dropped-retransmitted-control" { return "RxDroppedRetransmittedControl" }
    if yname == "tx-asyn-fifo-rate" { return "TxAsynFifoRate" }
    if yname == "rx-asyn-fifo-rate" { return "RxAsynFifoRate" }
    if yname == "rx-lfec-fec-uncorrectable-errors" { return "RxLfecFecUncorrectableErrors" }
    if yname == "rx-8b-10b-code-errors" { return "Rx8B10BCodeErrors" }
    return ""
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetSegmentPath() string {
    return "link-counters"
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-control-cells-counter"] = linkCounters.TxControlCellsCounter
    leafs["tx-data-cell-counter"] = linkCounters.TxDataCellCounter
    leafs["tx-data-byte-counter"] = linkCounters.TxDataByteCounter
    leafs["rx-crc-errors-counter"] = linkCounters.RxCrcErrorsCounter
    leafs["rx-lfec-fec-correctable-error"] = linkCounters.RxLfecFecCorrectableError
    leafs["rx-8b-10b-disparity-errors"] = linkCounters.Rx8B10BDisparityErrors
    leafs["rx-control-cells-counter"] = linkCounters.RxControlCellsCounter
    leafs["rx-data-cell-counter"] = linkCounters.RxDataCellCounter
    leafs["rx-data-byte-counter"] = linkCounters.RxDataByteCounter
    leafs["rx-dropped-retransmitted-control"] = linkCounters.RxDroppedRetransmittedControl
    leafs["tx-asyn-fifo-rate"] = linkCounters.TxAsynFifoRate
    leafs["rx-asyn-fifo-rate"] = linkCounters.RxAsynFifoRate
    leafs["rx-lfec-fec-uncorrectable-errors"] = linkCounters.RxLfecFecUncorrectableErrors
    leafs["rx-8b-10b-code-errors"] = linkCounters.Rx8B10BCodeErrors
    return leafs
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetBundleName() string { return "cisco_ios_xr" }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetYangName() string { return "link-counters" }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) SetParent(parent types.Entity) { linkCounters.parent = parent }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetParent() types.Entity { return linkCounters.parent }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetParentYangName() string { return "aggr-stats" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is string with length: 0..6.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is string with length: 0..6.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is string with length: 0..6.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is string with length: 0..6.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is string with length: 0..6.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is string with length: 0..6.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is string with length: 0..6.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is string with length: 0..6.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is string with length: 0..6.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is string with length: 0..6.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is string with length: 0..6.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is string with length: 0..6.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is string with length: 0..6.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is string with length: 0..6.
    Rx8B10BCodeErrors interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetFilter() yfilter.YFilter { return ovfStatus.YFilter }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) SetFilter(yf yfilter.YFilter) { ovfStatus.YFilter = yf }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetGoName(yname string) string {
    if yname == "tx-control-cells-counter" { return "TxControlCellsCounter" }
    if yname == "tx-data-cell-counter" { return "TxDataCellCounter" }
    if yname == "tx-data-byte-counter" { return "TxDataByteCounter" }
    if yname == "rx-crc-errors-counter" { return "RxCrcErrorsCounter" }
    if yname == "rx-lfec-fec-correctable-error" { return "RxLfecFecCorrectableError" }
    if yname == "rx-8b-10b-disparity-errors" { return "Rx8B10BDisparityErrors" }
    if yname == "rx-control-cells-counter" { return "RxControlCellsCounter" }
    if yname == "rx-data-cell-counter" { return "RxDataCellCounter" }
    if yname == "rx-data-byte-counter" { return "RxDataByteCounter" }
    if yname == "rx-dropped-retransmitted-control" { return "RxDroppedRetransmittedControl" }
    if yname == "tx-asyn-fifo-rate" { return "TxAsynFifoRate" }
    if yname == "rx-asyn-fifo-rate" { return "RxAsynFifoRate" }
    if yname == "rx-lfec-fec-uncorrectable-errors" { return "RxLfecFecUncorrectableErrors" }
    if yname == "rx-8b-10b-code-errors" { return "Rx8B10BCodeErrors" }
    return ""
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetSegmentPath() string {
    return "ovf-status"
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-control-cells-counter"] = ovfStatus.TxControlCellsCounter
    leafs["tx-data-cell-counter"] = ovfStatus.TxDataCellCounter
    leafs["tx-data-byte-counter"] = ovfStatus.TxDataByteCounter
    leafs["rx-crc-errors-counter"] = ovfStatus.RxCrcErrorsCounter
    leafs["rx-lfec-fec-correctable-error"] = ovfStatus.RxLfecFecCorrectableError
    leafs["rx-8b-10b-disparity-errors"] = ovfStatus.Rx8B10BDisparityErrors
    leafs["rx-control-cells-counter"] = ovfStatus.RxControlCellsCounter
    leafs["rx-data-cell-counter"] = ovfStatus.RxDataCellCounter
    leafs["rx-data-byte-counter"] = ovfStatus.RxDataByteCounter
    leafs["rx-dropped-retransmitted-control"] = ovfStatus.RxDroppedRetransmittedControl
    leafs["tx-asyn-fifo-rate"] = ovfStatus.TxAsynFifoRate
    leafs["rx-asyn-fifo-rate"] = ovfStatus.RxAsynFifoRate
    leafs["rx-lfec-fec-uncorrectable-errors"] = ovfStatus.RxLfecFecUncorrectableErrors
    leafs["rx-8b-10b-code-errors"] = ovfStatus.Rx8B10BCodeErrors
    return leafs
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetBundleName() string { return "cisco_ios_xr" }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetYangName() string { return "ovf-status" }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) SetParent(parent types.Entity) { ovfStatus.parent = parent }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetParent() types.Entity { return ovfStatus.parent }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetParentYangName() string { return "aggr-stats" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats
// incr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // link error status.
    LinkErrorStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus

    // link counters.
    LinkCounters Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetFilter() yfilter.YFilter { return incrStats.YFilter }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) SetFilter(yf yfilter.YFilter) { incrStats.YFilter = yf }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetGoName(yname string) string {
    if yname == "link-error-status" { return "LinkErrorStatus" }
    if yname == "link-counters" { return "LinkCounters" }
    if yname == "ovf-status" { return "OvfStatus" }
    return ""
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetSegmentPath() string {
    return "incr-stats"
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-error-status" {
        return &incrStats.LinkErrorStatus
    }
    if childYangName == "link-counters" {
        return &incrStats.LinkCounters
    }
    if childYangName == "ovf-status" {
        return &incrStats.OvfStatus
    }
    return nil
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-error-status"] = &incrStats.LinkErrorStatus
    children["link-counters"] = &incrStats.LinkCounters
    children["ovf-status"] = &incrStats.OvfStatus
    return children
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetBundleName() string { return "cisco_ios_xr" }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetYangName() string { return "incr-stats" }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) SetParent(parent types.Entity) { incrStats.parent = parent }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetParent() types.Entity { return incrStats.parent }

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetParentYangName() string { return "fmac-asic" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus
// link error status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // link crc error. The type is interface{} with range: 0..4294967295.
    LinkCrcError interface{}

    // link size error. The type is interface{} with range: 0..4294967295.
    LinkSizeError interface{}

    // link mis align error. The type is interface{} with range: 0..4294967295.
    LinkMisAlignError interface{}

    // link code group error. The type is interface{} with range: 0..4294967295.
    LinkCodeGroupError interface{}

    // link no sig lock error. The type is interface{} with range: 0..4294967295.
    LinkNoSigLockError interface{}

    // link no sig accept error. The type is interface{} with range:
    // 0..4294967295.
    LinkNoSigAcceptError interface{}

    // link tokens error. The type is interface{} with range: 0..4294967295.
    LinkTokensError interface{}

    // error token count. The type is interface{} with range: 0..4294967295.
    ErrorTokenCount interface{}
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetFilter() yfilter.YFilter { return linkErrorStatus.YFilter }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) SetFilter(yf yfilter.YFilter) { linkErrorStatus.YFilter = yf }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetGoName(yname string) string {
    if yname == "link-crc-error" { return "LinkCrcError" }
    if yname == "link-size-error" { return "LinkSizeError" }
    if yname == "link-mis-align-error" { return "LinkMisAlignError" }
    if yname == "link-code-group-error" { return "LinkCodeGroupError" }
    if yname == "link-no-sig-lock-error" { return "LinkNoSigLockError" }
    if yname == "link-no-sig-accept-error" { return "LinkNoSigAcceptError" }
    if yname == "link-tokens-error" { return "LinkTokensError" }
    if yname == "error-token-count" { return "ErrorTokenCount" }
    return ""
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetSegmentPath() string {
    return "link-error-status"
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-crc-error"] = linkErrorStatus.LinkCrcError
    leafs["link-size-error"] = linkErrorStatus.LinkSizeError
    leafs["link-mis-align-error"] = linkErrorStatus.LinkMisAlignError
    leafs["link-code-group-error"] = linkErrorStatus.LinkCodeGroupError
    leafs["link-no-sig-lock-error"] = linkErrorStatus.LinkNoSigLockError
    leafs["link-no-sig-accept-error"] = linkErrorStatus.LinkNoSigAcceptError
    leafs["link-tokens-error"] = linkErrorStatus.LinkTokensError
    leafs["error-token-count"] = linkErrorStatus.ErrorTokenCount
    return leafs
}

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetBundleName() string { return "cisco_ios_xr" }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetYangName() string { return "link-error-status" }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) SetParent(parent types.Entity) { linkErrorStatus.parent = parent }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetParent() types.Entity { return linkErrorStatus.parent }

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetParentYangName() string { return "incr-stats" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters
// link counters
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is interface{} with range:
    // 0..18446744073709551615.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is interface{} with range:
    // 0..18446744073709551615.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is interface{} with range:
    // 0..18446744073709551615.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is interface{} with range:
    // 0..18446744073709551615.
    Rx8B10BCodeErrors interface{}
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetFilter() yfilter.YFilter { return linkCounters.YFilter }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) SetFilter(yf yfilter.YFilter) { linkCounters.YFilter = yf }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetGoName(yname string) string {
    if yname == "tx-control-cells-counter" { return "TxControlCellsCounter" }
    if yname == "tx-data-cell-counter" { return "TxDataCellCounter" }
    if yname == "tx-data-byte-counter" { return "TxDataByteCounter" }
    if yname == "rx-crc-errors-counter" { return "RxCrcErrorsCounter" }
    if yname == "rx-lfec-fec-correctable-error" { return "RxLfecFecCorrectableError" }
    if yname == "rx-8b-10b-disparity-errors" { return "Rx8B10BDisparityErrors" }
    if yname == "rx-control-cells-counter" { return "RxControlCellsCounter" }
    if yname == "rx-data-cell-counter" { return "RxDataCellCounter" }
    if yname == "rx-data-byte-counter" { return "RxDataByteCounter" }
    if yname == "rx-dropped-retransmitted-control" { return "RxDroppedRetransmittedControl" }
    if yname == "tx-asyn-fifo-rate" { return "TxAsynFifoRate" }
    if yname == "rx-asyn-fifo-rate" { return "RxAsynFifoRate" }
    if yname == "rx-lfec-fec-uncorrectable-errors" { return "RxLfecFecUncorrectableErrors" }
    if yname == "rx-8b-10b-code-errors" { return "Rx8B10BCodeErrors" }
    return ""
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetSegmentPath() string {
    return "link-counters"
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-control-cells-counter"] = linkCounters.TxControlCellsCounter
    leafs["tx-data-cell-counter"] = linkCounters.TxDataCellCounter
    leafs["tx-data-byte-counter"] = linkCounters.TxDataByteCounter
    leafs["rx-crc-errors-counter"] = linkCounters.RxCrcErrorsCounter
    leafs["rx-lfec-fec-correctable-error"] = linkCounters.RxLfecFecCorrectableError
    leafs["rx-8b-10b-disparity-errors"] = linkCounters.Rx8B10BDisparityErrors
    leafs["rx-control-cells-counter"] = linkCounters.RxControlCellsCounter
    leafs["rx-data-cell-counter"] = linkCounters.RxDataCellCounter
    leafs["rx-data-byte-counter"] = linkCounters.RxDataByteCounter
    leafs["rx-dropped-retransmitted-control"] = linkCounters.RxDroppedRetransmittedControl
    leafs["tx-asyn-fifo-rate"] = linkCounters.TxAsynFifoRate
    leafs["rx-asyn-fifo-rate"] = linkCounters.RxAsynFifoRate
    leafs["rx-lfec-fec-uncorrectable-errors"] = linkCounters.RxLfecFecUncorrectableErrors
    leafs["rx-8b-10b-code-errors"] = linkCounters.Rx8B10BCodeErrors
    return leafs
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetBundleName() string { return "cisco_ios_xr" }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetYangName() string { return "link-counters" }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) SetParent(parent types.Entity) { linkCounters.parent = parent }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetParent() types.Entity { return linkCounters.parent }

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetParentYangName() string { return "incr-stats" }

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TX Control cells counter. The type is string with length: 0..6.
    TxControlCellsCounter interface{}

    // TX Data cell counter. The type is string with length: 0..6.
    TxDataCellCounter interface{}

    // TX Data byte counter. The type is string with length: 0..6.
    TxDataByteCounter interface{}

    // RX CRC errors counter. The type is string with length: 0..6.
    RxCrcErrorsCounter interface{}

    // RX LFEC FEC correctable error. The type is string with length: 0..6.
    RxLfecFecCorrectableError interface{}

    // RX 8b 10b disparity errors. The type is string with length: 0..6.
    Rx8B10BDisparityErrors interface{}

    // RX Control cells counter. The type is string with length: 0..6.
    RxControlCellsCounter interface{}

    // RX Data cell counter. The type is string with length: 0..6.
    RxDataCellCounter interface{}

    // RX Data byte counter. The type is string with length: 0..6.
    RxDataByteCounter interface{}

    // RX dropped retransmitted control. The type is string with length: 0..6.
    RxDroppedRetransmittedControl interface{}

    // TX Asyn fifo rate. The type is string with length: 0..6.
    TxAsynFifoRate interface{}

    // RX Asyn fifo rate. The type is string with length: 0..6.
    RxAsynFifoRate interface{}

    // RX LFEC FEC uncorrectable errors. The type is string with length: 0..6.
    RxLfecFecUncorrectableErrors interface{}

    // RX 8b 10b code errors. The type is string with length: 0..6.
    Rx8B10BCodeErrors interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetFilter() yfilter.YFilter { return ovfStatus.YFilter }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) SetFilter(yf yfilter.YFilter) { ovfStatus.YFilter = yf }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetGoName(yname string) string {
    if yname == "tx-control-cells-counter" { return "TxControlCellsCounter" }
    if yname == "tx-data-cell-counter" { return "TxDataCellCounter" }
    if yname == "tx-data-byte-counter" { return "TxDataByteCounter" }
    if yname == "rx-crc-errors-counter" { return "RxCrcErrorsCounter" }
    if yname == "rx-lfec-fec-correctable-error" { return "RxLfecFecCorrectableError" }
    if yname == "rx-8b-10b-disparity-errors" { return "Rx8B10BDisparityErrors" }
    if yname == "rx-control-cells-counter" { return "RxControlCellsCounter" }
    if yname == "rx-data-cell-counter" { return "RxDataCellCounter" }
    if yname == "rx-data-byte-counter" { return "RxDataByteCounter" }
    if yname == "rx-dropped-retransmitted-control" { return "RxDroppedRetransmittedControl" }
    if yname == "tx-asyn-fifo-rate" { return "TxAsynFifoRate" }
    if yname == "rx-asyn-fifo-rate" { return "RxAsynFifoRate" }
    if yname == "rx-lfec-fec-uncorrectable-errors" { return "RxLfecFecUncorrectableErrors" }
    if yname == "rx-8b-10b-code-errors" { return "Rx8B10BCodeErrors" }
    return ""
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetSegmentPath() string {
    return "ovf-status"
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tx-control-cells-counter"] = ovfStatus.TxControlCellsCounter
    leafs["tx-data-cell-counter"] = ovfStatus.TxDataCellCounter
    leafs["tx-data-byte-counter"] = ovfStatus.TxDataByteCounter
    leafs["rx-crc-errors-counter"] = ovfStatus.RxCrcErrorsCounter
    leafs["rx-lfec-fec-correctable-error"] = ovfStatus.RxLfecFecCorrectableError
    leafs["rx-8b-10b-disparity-errors"] = ovfStatus.Rx8B10BDisparityErrors
    leafs["rx-control-cells-counter"] = ovfStatus.RxControlCellsCounter
    leafs["rx-data-cell-counter"] = ovfStatus.RxDataCellCounter
    leafs["rx-data-byte-counter"] = ovfStatus.RxDataByteCounter
    leafs["rx-dropped-retransmitted-control"] = ovfStatus.RxDroppedRetransmittedControl
    leafs["tx-asyn-fifo-rate"] = ovfStatus.TxAsynFifoRate
    leafs["rx-asyn-fifo-rate"] = ovfStatus.RxAsynFifoRate
    leafs["rx-lfec-fec-uncorrectable-errors"] = ovfStatus.RxLfecFecUncorrectableErrors
    leafs["rx-8b-10b-code-errors"] = ovfStatus.Rx8B10BCodeErrors
    return leafs
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetBundleName() string { return "cisco_ios_xr" }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetYangName() string { return "ovf-status" }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) SetParent(parent types.Entity) { ovfStatus.parent = parent }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetParent() types.Entity { return ovfStatus.parent }

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetParentYangName() string { return "incr-stats" }

