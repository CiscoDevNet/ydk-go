// This module contains a collection of YANG definitions
// for Cisco IOS-XR dnx-driver package operational data.
// 
// This module contains definitions
// for the following management objects:
//   fia: FIA driver operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIA driver operational data for available nodes.
    Nodes Fia_Nodes
}

func (fia *Fia) GetEntityData() *types.CommonEntityData {
    fia.EntityData.YFilter = fia.YFilter
    fia.EntityData.YangName = "fia"
    fia.EntityData.BundleName = "cisco_ios_xr"
    fia.EntityData.ParentYangName = "Cisco-IOS-XR-dnx-driver-oper"
    fia.EntityData.SegmentPath = "Cisco-IOS-XR-dnx-driver-oper:fia"
    fia.EntityData.AbsolutePath = fia.EntityData.SegmentPath
    fia.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fia.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fia.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fia.EntityData.Children = types.NewOrderedMap()
    fia.EntityData.Children.Append("nodes", types.YChild{"Nodes", &fia.Nodes})
    fia.EntityData.Leafs = types.NewOrderedMap()

    fia.EntityData.YListKeys = []string {}

    return &(fia.EntityData)
}

// Fia_Nodes
// FIA driver operational data for available nodes
type Fia_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FIA operational data for a particular node. The type is slice of
    // Fia_Nodes_Node.
    Node []*Fia_Nodes_Node
}

func (nodes *Fia_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "fia"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/" + nodes.EntityData.SegmentPath
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

// Fia_Nodes_Node
// FIA operational data for a particular node
type Fia_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (node *Fia_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("rx-link-information", types.YChild{"RxLinkInformation", &node.RxLinkInformation})
    node.EntityData.Children.Append("driver-information", types.YChild{"DriverInformation", &node.DriverInformation})
    node.EntityData.Children.Append("clear-statistics", types.YChild{"ClearStatistics", &node.ClearStatistics})
    node.EntityData.Children.Append("tx-link-information", types.YChild{"TxLinkInformation", &node.TxLinkInformation})
    node.EntityData.Children.Append("diag-shell", types.YChild{"DiagShell", &node.DiagShell})
    node.EntityData.Children.Append("oir-history", types.YChild{"OirHistory", &node.OirHistory})
    node.EntityData.Children.Append("asic-statistics", types.YChild{"AsicStatistics", &node.AsicStatistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation
// FIA link rx information
type Fia_Nodes_Node_RxLinkInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Option table for link rx information.
    LinkOptions Fia_Nodes_Node_RxLinkInformation_LinkOptions
}

func (rxLinkInformation *Fia_Nodes_Node_RxLinkInformation) GetEntityData() *types.CommonEntityData {
    rxLinkInformation.EntityData.YFilter = rxLinkInformation.YFilter
    rxLinkInformation.EntityData.YangName = "rx-link-information"
    rxLinkInformation.EntityData.BundleName = "cisco_ios_xr"
    rxLinkInformation.EntityData.ParentYangName = "node"
    rxLinkInformation.EntityData.SegmentPath = "rx-link-information"
    rxLinkInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + rxLinkInformation.EntityData.SegmentPath
    rxLinkInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLinkInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLinkInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLinkInformation.EntityData.Children = types.NewOrderedMap()
    rxLinkInformation.EntityData.Children.Append("link-options", types.YChild{"LinkOptions", &rxLinkInformation.LinkOptions})
    rxLinkInformation.EntityData.Leafs = types.NewOrderedMap()

    rxLinkInformation.EntityData.YListKeys = []string {}

    return &(rxLinkInformation.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions
// Option table for link rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Option : topo , flag , stats. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption.
    LinkOption []*Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption
}

func (linkOptions *Fia_Nodes_Node_RxLinkInformation_LinkOptions) GetEntityData() *types.CommonEntityData {
    linkOptions.EntityData.YFilter = linkOptions.YFilter
    linkOptions.EntityData.YangName = "link-options"
    linkOptions.EntityData.BundleName = "cisco_ios_xr"
    linkOptions.EntityData.ParentYangName = "rx-link-information"
    linkOptions.EntityData.SegmentPath = "link-options"
    linkOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/" + linkOptions.EntityData.SegmentPath
    linkOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkOptions.EntityData.Children = types.NewOrderedMap()
    linkOptions.EntityData.Children.Append("link-option", types.YChild{"LinkOption", nil})
    for i := range linkOptions.LinkOption {
        linkOptions.EntityData.Children.Append(types.GetSegmentPath(linkOptions.LinkOption[i]), types.YChild{"LinkOption", linkOptions.LinkOption[i]})
    }
    linkOptions.EntityData.Leafs = types.NewOrderedMap()

    linkOptions.EntityData.YListKeys = []string {}

    return &(linkOptions.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption
// Option : topo , flag , stats
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link option. The type is string with pattern:
    // (flap)|(topo).
    Option interface{}

    // Instance table for rx information.
    RxAsicInstances Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances
}

func (linkOption *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption) GetEntityData() *types.CommonEntityData {
    linkOption.EntityData.YFilter = linkOption.YFilter
    linkOption.EntityData.YangName = "link-option"
    linkOption.EntityData.BundleName = "cisco_ios_xr"
    linkOption.EntityData.ParentYangName = "link-options"
    linkOption.EntityData.SegmentPath = "link-option" + types.AddKeyToken(linkOption.Option, "option")
    linkOption.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/" + linkOption.EntityData.SegmentPath
    linkOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkOption.EntityData.Children = types.NewOrderedMap()
    linkOption.EntityData.Children.Append("rx-asic-instances", types.YChild{"RxAsicInstances", &linkOption.RxAsicInstances})
    linkOption.EntityData.Leafs = types.NewOrderedMap()
    linkOption.EntityData.Leafs.Append("option", types.YLeaf{"Option", linkOption.Option})

    linkOption.EntityData.YListKeys = []string {"Option"}

    return &(linkOption.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances
// Instance table for rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance number for rx link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance.
    RxAsicInstance []*Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance
}

func (rxAsicInstances *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances) GetEntityData() *types.CommonEntityData {
    rxAsicInstances.EntityData.YFilter = rxAsicInstances.YFilter
    rxAsicInstances.EntityData.YangName = "rx-asic-instances"
    rxAsicInstances.EntityData.BundleName = "cisco_ios_xr"
    rxAsicInstances.EntityData.ParentYangName = "link-option"
    rxAsicInstances.EntityData.SegmentPath = "rx-asic-instances"
    rxAsicInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/" + rxAsicInstances.EntityData.SegmentPath
    rxAsicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxAsicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxAsicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxAsicInstances.EntityData.Children = types.NewOrderedMap()
    rxAsicInstances.EntityData.Children.Append("rx-asic-instance", types.YChild{"RxAsicInstance", nil})
    for i := range rxAsicInstances.RxAsicInstance {
        rxAsicInstances.EntityData.Children.Append(types.GetSegmentPath(rxAsicInstances.RxAsicInstance[i]), types.YChild{"RxAsicInstance", rxAsicInstances.RxAsicInstance[i]})
    }
    rxAsicInstances.EntityData.Leafs = types.NewOrderedMap()

    rxAsicInstances.EntityData.YListKeys = []string {}

    return &(rxAsicInstances.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance
// Instance number for rx link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Receive instance. The type is interface{} with
    // range: 0..255.
    Instance interface{}

    // Link table class for rx information.
    RxLinks Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks
}

func (rxAsicInstance *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance) GetEntityData() *types.CommonEntityData {
    rxAsicInstance.EntityData.YFilter = rxAsicInstance.YFilter
    rxAsicInstance.EntityData.YangName = "rx-asic-instance"
    rxAsicInstance.EntityData.BundleName = "cisco_ios_xr"
    rxAsicInstance.EntityData.ParentYangName = "rx-asic-instances"
    rxAsicInstance.EntityData.SegmentPath = "rx-asic-instance" + types.AddKeyToken(rxAsicInstance.Instance, "instance")
    rxAsicInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/" + rxAsicInstance.EntityData.SegmentPath
    rxAsicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxAsicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxAsicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxAsicInstance.EntityData.Children = types.NewOrderedMap()
    rxAsicInstance.EntityData.Children.Append("rx-links", types.YChild{"RxLinks", &rxAsicInstance.RxLinks})
    rxAsicInstance.EntityData.Leafs = types.NewOrderedMap()
    rxAsicInstance.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", rxAsicInstance.Instance})

    rxAsicInstance.EntityData.YListKeys = []string {"Instance"}

    return &(rxAsicInstance.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks
// Link table class for rx information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link number for rx link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink.
    RxLink []*Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink
}

func (rxLinks *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks) GetEntityData() *types.CommonEntityData {
    rxLinks.EntityData.YFilter = rxLinks.YFilter
    rxLinks.EntityData.YangName = "rx-links"
    rxLinks.EntityData.BundleName = "cisco_ios_xr"
    rxLinks.EntityData.ParentYangName = "rx-asic-instance"
    rxLinks.EntityData.SegmentPath = "rx-links"
    rxLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/" + rxLinks.EntityData.SegmentPath
    rxLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLinks.EntityData.Children = types.NewOrderedMap()
    rxLinks.EntityData.Children.Append("rx-link", types.YChild{"RxLink", nil})
    for i := range rxLinks.RxLink {
        types.SetYListKey(rxLinks.RxLink[i], i)
        rxLinks.EntityData.Children.Append(types.GetSegmentPath(rxLinks.RxLink[i]), types.YChild{"RxLink", rxLinks.RxLink[i]})
    }
    rxLinks.EntityData.Leafs = types.NewOrderedMap()

    rxLinks.EntityData.YListKeys = []string {}

    return &(rxLinks.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink
// Link number for rx link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Start number. The type is interface{} with range: 0..47.
    StartNumber interface{}

    // End number. The type is interface{} with range: 0..47.
    EndNumber interface{}

    // RX link status option. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    StatusOption interface{}

    // Single link information. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink.
    RxLink []*Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink
}

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink) GetEntityData() *types.CommonEntityData {
    rxLink.EntityData.YFilter = rxLink.YFilter
    rxLink.EntityData.YangName = "rx-link"
    rxLink.EntityData.BundleName = "cisco_ios_xr"
    rxLink.EntityData.ParentYangName = "rx-links"
    rxLink.EntityData.SegmentPath = "rx-link" + types.AddNoKeyToken(rxLink)
    rxLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/" + rxLink.EntityData.SegmentPath
    rxLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLink.EntityData.Children = types.NewOrderedMap()
    rxLink.EntityData.Children.Append("rx-link", types.YChild{"RxLink", nil})
    for i := range rxLink.RxLink {
        rxLink.EntityData.Children.Append(types.GetSegmentPath(rxLink.RxLink[i]), types.YChild{"RxLink", rxLink.RxLink[i]})
    }
    rxLink.EntityData.Leafs = types.NewOrderedMap()
    rxLink.EntityData.Leafs.Append("start-number", types.YLeaf{"StartNumber", rxLink.StartNumber})
    rxLink.EntityData.Leafs.Append("end-number", types.YLeaf{"EndNumber", rxLink.EndNumber})
    rxLink.EntityData.Leafs.Append("status-option", types.YLeaf{"StatusOption", rxLink.StatusOption})

    rxLink.EntityData.YListKeys = []string {}

    return &(rxLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink
// Single link information
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Single link. The type is interface{} with range:
    // 0..4294967295.
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

func (rxLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink) GetEntityData() *types.CommonEntityData {
    rxLink.EntityData.YFilter = rxLink.YFilter
    rxLink.EntityData.YangName = "rx-link"
    rxLink.EntityData.BundleName = "cisco_ios_xr"
    rxLink.EntityData.ParentYangName = "rx-link"
    rxLink.EntityData.SegmentPath = "rx-link" + types.AddKeyToken(rxLink.Link, "link")
    rxLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/" + rxLink.EntityData.SegmentPath
    rxLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rxLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rxLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rxLink.EntityData.Children = types.NewOrderedMap()
    rxLink.EntityData.Children.Append("this-link", types.YChild{"ThisLink", &rxLink.ThisLink})
    rxLink.EntityData.Children.Append("far-end-link", types.YChild{"FarEndLink", &rxLink.FarEndLink})
    rxLink.EntityData.Children.Append("far-end-link-in-hw", types.YChild{"FarEndLinkInHw", &rxLink.FarEndLinkInHw})
    rxLink.EntityData.Children.Append("history", types.YChild{"History", &rxLink.History})
    rxLink.EntityData.Leafs = types.NewOrderedMap()
    rxLink.EntityData.Leafs.Append("link", types.YLeaf{"Link", rxLink.Link})
    rxLink.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", rxLink.Speed})
    rxLink.EntityData.Leafs.Append("stage", types.YLeaf{"Stage", rxLink.Stage})
    rxLink.EntityData.Leafs.Append("is-link-valid", types.YLeaf{"IsLinkValid", rxLink.IsLinkValid})
    rxLink.EntityData.Leafs.Append("is-conf-pending", types.YLeaf{"IsConfPending", rxLink.IsConfPending})
    rxLink.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", rxLink.AdminState})
    rxLink.EntityData.Leafs.Append("oper-state", types.YLeaf{"OperState", rxLink.OperState})
    rxLink.EntityData.Leafs.Append("error-state", types.YLeaf{"ErrorState", rxLink.ErrorState})
    rxLink.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", rxLink.Flags})
    rxLink.EntityData.Leafs.Append("flap-cnt", types.YLeaf{"FlapCnt", rxLink.FlapCnt})
    rxLink.EntityData.Leafs.Append("num-admin-shuts", types.YLeaf{"NumAdminShuts", rxLink.NumAdminShuts})
    rxLink.EntityData.Leafs.Append("correctable-errors", types.YLeaf{"CorrectableErrors", rxLink.CorrectableErrors})
    rxLink.EntityData.Leafs.Append("uncorrectable-errors", types.YLeaf{"UncorrectableErrors", rxLink.UncorrectableErrors})

    rxLink.EntityData.YListKeys = []string {"Link"}

    return &(rxLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink
// this link
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink struct {
    EntityData types.CommonEntityData
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

func (thisLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink) GetEntityData() *types.CommonEntityData {
    thisLink.EntityData.YFilter = thisLink.YFilter
    thisLink.EntityData.YangName = "this-link"
    thisLink.EntityData.BundleName = "cisco_ios_xr"
    thisLink.EntityData.ParentYangName = "rx-link"
    thisLink.EntityData.SegmentPath = "this-link"
    thisLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/" + thisLink.EntityData.SegmentPath
    thisLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thisLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thisLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thisLink.EntityData.Children = types.NewOrderedMap()
    thisLink.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &thisLink.AsicId})
    thisLink.EntityData.Leafs = types.NewOrderedMap()
    thisLink.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", thisLink.LinkType})
    thisLink.EntityData.Leafs.Append("link-stage", types.YLeaf{"LinkStage", thisLink.LinkStage})
    thisLink.EntityData.Leafs.Append("link-num", types.YLeaf{"LinkNum", thisLink.LinkNum})
    thisLink.EntityData.Leafs.Append("phy-link-num", types.YLeaf{"PhyLinkNum", thisLink.PhyLinkNum})

    thisLink.EntityData.YListKeys = []string {}

    return &(thisLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_ThisLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "this-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/this-link/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink
// far end link
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink struct {
    EntityData types.CommonEntityData
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

func (farEndLink *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink) GetEntityData() *types.CommonEntityData {
    farEndLink.EntityData.YFilter = farEndLink.YFilter
    farEndLink.EntityData.YangName = "far-end-link"
    farEndLink.EntityData.BundleName = "cisco_ios_xr"
    farEndLink.EntityData.ParentYangName = "rx-link"
    farEndLink.EntityData.SegmentPath = "far-end-link"
    farEndLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/" + farEndLink.EntityData.SegmentPath
    farEndLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    farEndLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    farEndLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    farEndLink.EntityData.Children = types.NewOrderedMap()
    farEndLink.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &farEndLink.AsicId})
    farEndLink.EntityData.Leafs = types.NewOrderedMap()
    farEndLink.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", farEndLink.LinkType})
    farEndLink.EntityData.Leafs.Append("link-stage", types.YLeaf{"LinkStage", farEndLink.LinkStage})
    farEndLink.EntityData.Leafs.Append("link-num", types.YLeaf{"LinkNum", farEndLink.LinkNum})
    farEndLink.EntityData.Leafs.Append("phy-link-num", types.YLeaf{"PhyLinkNum", farEndLink.PhyLinkNum})

    farEndLink.EntityData.YListKeys = []string {}

    return &(farEndLink.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "far-end-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/far-end-link/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw
// far end link in hw
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw struct {
    EntityData types.CommonEntityData
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

func (farEndLinkInHw *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw) GetEntityData() *types.CommonEntityData {
    farEndLinkInHw.EntityData.YFilter = farEndLinkInHw.YFilter
    farEndLinkInHw.EntityData.YangName = "far-end-link-in-hw"
    farEndLinkInHw.EntityData.BundleName = "cisco_ios_xr"
    farEndLinkInHw.EntityData.ParentYangName = "rx-link"
    farEndLinkInHw.EntityData.SegmentPath = "far-end-link-in-hw"
    farEndLinkInHw.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/" + farEndLinkInHw.EntityData.SegmentPath
    farEndLinkInHw.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    farEndLinkInHw.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    farEndLinkInHw.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    farEndLinkInHw.EntityData.Children = types.NewOrderedMap()
    farEndLinkInHw.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &farEndLinkInHw.AsicId})
    farEndLinkInHw.EntityData.Leafs = types.NewOrderedMap()
    farEndLinkInHw.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", farEndLinkInHw.LinkType})
    farEndLinkInHw.EntityData.Leafs.Append("link-stage", types.YLeaf{"LinkStage", farEndLinkInHw.LinkStage})
    farEndLinkInHw.EntityData.Leafs.Append("link-num", types.YLeaf{"LinkNum", farEndLinkInHw.LinkNum})
    farEndLinkInHw.EntityData.Leafs.Append("phy-link-num", types.YLeaf{"PhyLinkNum", farEndLinkInHw.PhyLinkNum})

    farEndLinkInHw.EntityData.YListKeys = []string {}

    return &(farEndLinkInHw.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId
// asic id
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_FarEndLinkInHw_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "far-end-link-in-hw"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/far-end-link-in-hw/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History
// history
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // histnum. The type is interface{} with range: 0..255.
    Histnum interface{}

    // start index. The type is interface{} with range: 0..255.
    StartIndex interface{}

    // hist. The type is slice of
    // Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist.
    Hist []*Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist
}

func (history *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "rx-link"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("hist", types.YChild{"Hist", nil})
    for i := range history.Hist {
        types.SetYListKey(history.Hist[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.Hist[i]), types.YChild{"Hist", history.Hist[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("histnum", types.YLeaf{"Histnum", history.Histnum})
    history.EntityData.Leafs.Append("start-index", types.YLeaf{"StartIndex", history.StartIndex})

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist
// hist
type Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hist *Fia_Nodes_Node_RxLinkInformation_LinkOptions_LinkOption_RxAsicInstances_RxAsicInstance_RxLinks_RxLink_RxLink_History_Hist) GetEntityData() *types.CommonEntityData {
    hist.EntityData.YFilter = hist.YFilter
    hist.EntityData.YangName = "hist"
    hist.EntityData.BundleName = "cisco_ios_xr"
    hist.EntityData.ParentYangName = "history"
    hist.EntityData.SegmentPath = "hist" + types.AddNoKeyToken(hist)
    hist.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/rx-link-information/link-options/link-option/rx-asic-instances/rx-asic-instance/rx-links/rx-link/rx-link/history/" + hist.EntityData.SegmentPath
    hist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hist.EntityData.Children = types.NewOrderedMap()
    hist.EntityData.Leafs = types.NewOrderedMap()
    hist.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", hist.AdminState})
    hist.EntityData.Leafs.Append("oper-state", types.YLeaf{"OperState", hist.OperState})
    hist.EntityData.Leafs.Append("error-state", types.YLeaf{"ErrorState", hist.ErrorState})
    hist.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", hist.Timestamp})
    hist.EntityData.Leafs.Append("reasons", types.YLeaf{"Reasons", hist.Reasons})

    hist.EntityData.YListKeys = []string {}

    return &(hist.EntityData)
}

// Fia_Nodes_Node_DriverInformation
// FIA driver information
type Fia_Nodes_Node_DriverInformation struct {
    EntityData types.CommonEntityData
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

    // all wb insync. The type is bool.
    AllWbInsync interface{}

    // all wb insync since. The type is interface{} with range: 0..4294967295.
    AllWbInsyncSince interface{}

    // all startup wb insync. The type is bool.
    AllStartupWbInsync interface{}

    // planeA bitmap. The type is interface{} with range: 0..4294967295.
    PlaneABitmap interface{}

    // planeB bitmap. The type is interface{} with range: 0..4294967295.
    PlaneBBitmap interface{}

    // device info. The type is slice of
    // Fia_Nodes_Node_DriverInformation_DeviceInfo.
    DeviceInfo []*Fia_Nodes_Node_DriverInformation_DeviceInfo

    // card info. The type is slice of Fia_Nodes_Node_DriverInformation_CardInfo.
    CardInfo []*Fia_Nodes_Node_DriverInformation_CardInfo
}

func (driverInformation *Fia_Nodes_Node_DriverInformation) GetEntityData() *types.CommonEntityData {
    driverInformation.EntityData.YFilter = driverInformation.YFilter
    driverInformation.EntityData.YangName = "driver-information"
    driverInformation.EntityData.BundleName = "cisco_ios_xr"
    driverInformation.EntityData.ParentYangName = "node"
    driverInformation.EntityData.SegmentPath = "driver-information"
    driverInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + driverInformation.EntityData.SegmentPath
    driverInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    driverInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    driverInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    driverInformation.EntityData.Children = types.NewOrderedMap()
    driverInformation.EntityData.Children.Append("device-info", types.YChild{"DeviceInfo", nil})
    for i := range driverInformation.DeviceInfo {
        types.SetYListKey(driverInformation.DeviceInfo[i], i)
        driverInformation.EntityData.Children.Append(types.GetSegmentPath(driverInformation.DeviceInfo[i]), types.YChild{"DeviceInfo", driverInformation.DeviceInfo[i]})
    }
    driverInformation.EntityData.Children.Append("card-info", types.YChild{"CardInfo", nil})
    for i := range driverInformation.CardInfo {
        types.SetYListKey(driverInformation.CardInfo[i], i)
        driverInformation.EntityData.Children.Append(types.GetSegmentPath(driverInformation.CardInfo[i]), types.YChild{"CardInfo", driverInformation.CardInfo[i]})
    }
    driverInformation.EntityData.Leafs = types.NewOrderedMap()
    driverInformation.EntityData.Leafs.Append("drv-version", types.YLeaf{"DrvVersion", driverInformation.DrvVersion})
    driverInformation.EntityData.Leafs.Append("coeff-major-rev", types.YLeaf{"CoeffMajorRev", driverInformation.CoeffMajorRev})
    driverInformation.EntityData.Leafs.Append("coeff-minor-rev", types.YLeaf{"CoeffMinorRev", driverInformation.CoeffMinorRev})
    driverInformation.EntityData.Leafs.Append("functional-role", types.YLeaf{"FunctionalRole", driverInformation.FunctionalRole})
    driverInformation.EntityData.Leafs.Append("issu-role", types.YLeaf{"IssuRole", driverInformation.IssuRole})
    driverInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", driverInformation.NodeId})
    driverInformation.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", driverInformation.RackType})
    driverInformation.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", driverInformation.RackNum})
    driverInformation.EntityData.Leafs.Append("is-driver-ready", types.YLeaf{"IsDriverReady", driverInformation.IsDriverReady})
    driverInformation.EntityData.Leafs.Append("card-avail-mask", types.YLeaf{"CardAvailMask", driverInformation.CardAvailMask})
    driverInformation.EntityData.Leafs.Append("asic-avail-mask", types.YLeaf{"AsicAvailMask", driverInformation.AsicAvailMask})
    driverInformation.EntityData.Leafs.Append("exp-asic-avail-mask", types.YLeaf{"ExpAsicAvailMask", driverInformation.ExpAsicAvailMask})
    driverInformation.EntityData.Leafs.Append("ucmc-ratio", types.YLeaf{"UcmcRatio", driverInformation.UcmcRatio})
    driverInformation.EntityData.Leafs.Append("asic-oper-notify-to-fsdb-pending-bmap", types.YLeaf{"AsicOperNotifyToFsdbPendingBmap", driverInformation.AsicOperNotifyToFsdbPendingBmap})
    driverInformation.EntityData.Leafs.Append("is-full-fgid-download-req", types.YLeaf{"IsFullFgidDownloadReq", driverInformation.IsFullFgidDownloadReq})
    driverInformation.EntityData.Leafs.Append("is-fgid-download-in-progress", types.YLeaf{"IsFgidDownloadInProgress", driverInformation.IsFgidDownloadInProgress})
    driverInformation.EntityData.Leafs.Append("is-fgid-download-completed", types.YLeaf{"IsFgidDownloadCompleted", driverInformation.IsFgidDownloadCompleted})
    driverInformation.EntityData.Leafs.Append("fsdb-conn-active", types.YLeaf{"FsdbConnActive", driverInformation.FsdbConnActive})
    driverInformation.EntityData.Leafs.Append("fgid-conn-active", types.YLeaf{"FgidConnActive", driverInformation.FgidConnActive})
    driverInformation.EntityData.Leafs.Append("issu-mgr-conn-active", types.YLeaf{"IssuMgrConnActive", driverInformation.IssuMgrConnActive})
    driverInformation.EntityData.Leafs.Append("fsdb-reg-active", types.YLeaf{"FsdbRegActive", driverInformation.FsdbRegActive})
    driverInformation.EntityData.Leafs.Append("fgid-reg-active", types.YLeaf{"FgidRegActive", driverInformation.FgidRegActive})
    driverInformation.EntityData.Leafs.Append("issu-mgr-reg-active", types.YLeaf{"IssuMgrRegActive", driverInformation.IssuMgrRegActive})
    driverInformation.EntityData.Leafs.Append("num-pm-conn-reqs", types.YLeaf{"NumPmConnReqs", driverInformation.NumPmConnReqs})
    driverInformation.EntityData.Leafs.Append("num-fsdb-conn-reqs", types.YLeaf{"NumFsdbConnReqs", driverInformation.NumFsdbConnReqs})
    driverInformation.EntityData.Leafs.Append("num-fgid-conn-reqs", types.YLeaf{"NumFgidConnReqs", driverInformation.NumFgidConnReqs})
    driverInformation.EntityData.Leafs.Append("num-fstats-conn-reqs", types.YLeaf{"NumFstatsConnReqs", driverInformation.NumFstatsConnReqs})
    driverInformation.EntityData.Leafs.Append("num-cm-conn-reqs", types.YLeaf{"NumCmConnReqs", driverInformation.NumCmConnReqs})
    driverInformation.EntityData.Leafs.Append("num-issu-mgr-conn-reqs", types.YLeaf{"NumIssuMgrConnReqs", driverInformation.NumIssuMgrConnReqs})
    driverInformation.EntityData.Leafs.Append("num-peer-fia-conn-reqs", types.YLeaf{"NumPeerFiaConnReqs", driverInformation.NumPeerFiaConnReqs})
    driverInformation.EntityData.Leafs.Append("is-gaspp-registered", types.YLeaf{"IsGasppRegistered", driverInformation.IsGasppRegistered})
    driverInformation.EntityData.Leafs.Append("is-cih-registered", types.YLeaf{"IsCihRegistered", driverInformation.IsCihRegistered})
    driverInformation.EntityData.Leafs.Append("drvr-initial-startup-timestamp", types.YLeaf{"DrvrInitialStartupTimestamp", driverInformation.DrvrInitialStartupTimestamp})
    driverInformation.EntityData.Leafs.Append("drvr-current-startup-timestamp", types.YLeaf{"DrvrCurrentStartupTimestamp", driverInformation.DrvrCurrentStartupTimestamp})
    driverInformation.EntityData.Leafs.Append("num-intf-ports", types.YLeaf{"NumIntfPorts", driverInformation.NumIntfPorts})
    driverInformation.EntityData.Leafs.Append("uc-weight", types.YLeaf{"UcWeight", driverInformation.UcWeight})
    driverInformation.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", driverInformation.RespawnCount})
    driverInformation.EntityData.Leafs.Append("total-asics", types.YLeaf{"TotalAsics", driverInformation.TotalAsics})
    driverInformation.EntityData.Leafs.Append("issu-ready-ntfy-pending", types.YLeaf{"IssuReadyNtfyPending", driverInformation.IssuReadyNtfyPending})
    driverInformation.EntityData.Leafs.Append("issu-abort-sent", types.YLeaf{"IssuAbortSent", driverInformation.IssuAbortSent})
    driverInformation.EntityData.Leafs.Append("issu-abort-rcvd", types.YLeaf{"IssuAbortRcvd", driverInformation.IssuAbortRcvd})
    driverInformation.EntityData.Leafs.Append("fabric-mode", types.YLeaf{"FabricMode", driverInformation.FabricMode})
    driverInformation.EntityData.Leafs.Append("fc-mode", types.YLeaf{"FcMode", driverInformation.FcMode})
    driverInformation.EntityData.Leafs.Append("board-rev-id", types.YLeaf{"BoardRevId", driverInformation.BoardRevId})
    driverInformation.EntityData.Leafs.Append("all-wb-insync", types.YLeaf{"AllWbInsync", driverInformation.AllWbInsync})
    driverInformation.EntityData.Leafs.Append("all-wb-insync-since", types.YLeaf{"AllWbInsyncSince", driverInformation.AllWbInsyncSince})
    driverInformation.EntityData.Leafs.Append("all-startup-wb-insync", types.YLeaf{"AllStartupWbInsync", driverInformation.AllStartupWbInsync})
    driverInformation.EntityData.Leafs.Append("plane-a-bitmap", types.YLeaf{"PlaneABitmap", driverInformation.PlaneABitmap})
    driverInformation.EntityData.Leafs.Append("plane-b-bitmap", types.YLeaf{"PlaneBBitmap", driverInformation.PlaneBBitmap})

    driverInformation.EntityData.YListKeys = []string {}

    return &(driverInformation.EntityData)
}

// Fia_Nodes_Node_DriverInformation_DeviceInfo
// device info
type Fia_Nodes_Node_DriverInformation_DeviceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // startup wb mtime str. The type is string.
    StartupWbMtimeStr interface{}

    // startup wb outof sync. The type is bool.
    StartupWbOutofSync interface{}

    // local wb sync end str. The type is string.
    LocalWbSyncEndStr interface{}

    // remote wb sync end str. The type is string.
    RemoteWbSyncEndStr interface{}

    // local wb sync pending. The type is bool.
    LocalWbSyncPending interface{}

    // sdk delay msec. The type is interface{} with range: 0..4294967295.
    SdkDelayMsec interface{}

    // asic id.
    AsicId Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId
}

func (deviceInfo *Fia_Nodes_Node_DriverInformation_DeviceInfo) GetEntityData() *types.CommonEntityData {
    deviceInfo.EntityData.YFilter = deviceInfo.YFilter
    deviceInfo.EntityData.YangName = "device-info"
    deviceInfo.EntityData.BundleName = "cisco_ios_xr"
    deviceInfo.EntityData.ParentYangName = "driver-information"
    deviceInfo.EntityData.SegmentPath = "device-info" + types.AddNoKeyToken(deviceInfo)
    deviceInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/driver-information/" + deviceInfo.EntityData.SegmentPath
    deviceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deviceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deviceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deviceInfo.EntityData.Children = types.NewOrderedMap()
    deviceInfo.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &deviceInfo.AsicId})
    deviceInfo.EntityData.Leafs = types.NewOrderedMap()
    deviceInfo.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", deviceInfo.IsValid})
    deviceInfo.EntityData.Leafs.Append("fapid", types.YLeaf{"Fapid", deviceInfo.Fapid})
    deviceInfo.EntityData.Leafs.Append("hotplug-event", types.YLeaf{"HotplugEvent", deviceInfo.HotplugEvent})
    deviceInfo.EntityData.Leafs.Append("slice-state", types.YLeaf{"SliceState", deviceInfo.SliceState})
    deviceInfo.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", deviceInfo.AdminState})
    deviceInfo.EntityData.Leafs.Append("oper-state", types.YLeaf{"OperState", deviceInfo.OperState})
    deviceInfo.EntityData.Leafs.Append("asic-state", types.YLeaf{"AsicState", deviceInfo.AsicState})
    deviceInfo.EntityData.Leafs.Append("last-init-cause", types.YLeaf{"LastInitCause", deviceInfo.LastInitCause})
    deviceInfo.EntityData.Leafs.Append("num-pon-resets", types.YLeaf{"NumPonResets", deviceInfo.NumPonResets})
    deviceInfo.EntityData.Leafs.Append("num-hard-resets", types.YLeaf{"NumHardResets", deviceInfo.NumHardResets})
    deviceInfo.EntityData.Leafs.Append("local-switch-state", types.YLeaf{"LocalSwitchState", deviceInfo.LocalSwitchState})
    deviceInfo.EntityData.Leafs.Append("startup-wb-mtime-str", types.YLeaf{"StartupWbMtimeStr", deviceInfo.StartupWbMtimeStr})
    deviceInfo.EntityData.Leafs.Append("startup-wb-outof-sync", types.YLeaf{"StartupWbOutofSync", deviceInfo.StartupWbOutofSync})
    deviceInfo.EntityData.Leafs.Append("local-wb-sync-end-str", types.YLeaf{"LocalWbSyncEndStr", deviceInfo.LocalWbSyncEndStr})
    deviceInfo.EntityData.Leafs.Append("remote-wb-sync-end-str", types.YLeaf{"RemoteWbSyncEndStr", deviceInfo.RemoteWbSyncEndStr})
    deviceInfo.EntityData.Leafs.Append("local-wb-sync-pending", types.YLeaf{"LocalWbSyncPending", deviceInfo.LocalWbSyncPending})
    deviceInfo.EntityData.Leafs.Append("sdk-delay-msec", types.YLeaf{"SdkDelayMsec", deviceInfo.SdkDelayMsec})

    deviceInfo.EntityData.YListKeys = []string {}

    return &(deviceInfo.EntityData)
}

// Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId
// asic id
type Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_DriverInformation_DeviceInfo_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "device-info"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/driver-information/device-info/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_DriverInformation_CardInfo
// card info
type Fia_Nodes_Node_DriverInformation_CardInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (cardInfo *Fia_Nodes_Node_DriverInformation_CardInfo) GetEntityData() *types.CommonEntityData {
    cardInfo.EntityData.YFilter = cardInfo.YFilter
    cardInfo.EntityData.YangName = "card-info"
    cardInfo.EntityData.BundleName = "cisco_ios_xr"
    cardInfo.EntityData.ParentYangName = "driver-information"
    cardInfo.EntityData.SegmentPath = "card-info" + types.AddNoKeyToken(cardInfo)
    cardInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/driver-information/" + cardInfo.EntityData.SegmentPath
    cardInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInfo.EntityData.Children = types.NewOrderedMap()
    cardInfo.EntityData.Children.Append("oir-circular-buffer", types.YChild{"OirCircularBuffer", &cardInfo.OirCircularBuffer})
    cardInfo.EntityData.Leafs = types.NewOrderedMap()
    cardInfo.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", cardInfo.CardType})
    cardInfo.EntityData.Leafs.Append("card-name", types.YLeaf{"CardName", cardInfo.CardName})
    cardInfo.EntityData.Leafs.Append("slot-no", types.YLeaf{"SlotNo", cardInfo.SlotNo})
    cardInfo.EntityData.Leafs.Append("card-flag", types.YLeaf{"CardFlag", cardInfo.CardFlag})
    cardInfo.EntityData.Leafs.Append("evt-flag", types.YLeaf{"EvtFlag", cardInfo.EvtFlag})
    cardInfo.EntityData.Leafs.Append("reg-flag", types.YLeaf{"RegFlag", cardInfo.RegFlag})
    cardInfo.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", cardInfo.Instance})
    cardInfo.EntityData.Leafs.Append("card-state", types.YLeaf{"CardState", cardInfo.CardState})
    cardInfo.EntityData.Leafs.Append("exp-num-asics", types.YLeaf{"ExpNumAsics", cardInfo.ExpNumAsics})
    cardInfo.EntityData.Leafs.Append("exp-num-asics-per-fsdb", types.YLeaf{"ExpNumAsicsPerFsdb", cardInfo.ExpNumAsicsPerFsdb})
    cardInfo.EntityData.Leafs.Append("is-powered", types.YLeaf{"IsPowered", cardInfo.IsPowered})
    cardInfo.EntityData.Leafs.Append("cxp-avail-bitmap", types.YLeaf{"CxpAvailBitmap", cardInfo.CxpAvailBitmap})
    cardInfo.EntityData.Leafs.Append("num-ilkns-per-asic", types.YLeaf{"NumIlknsPerAsic", cardInfo.NumIlknsPerAsic})
    cardInfo.EntityData.Leafs.Append("num-local-ports-per-ilkn", types.YLeaf{"NumLocalPortsPerIlkn", cardInfo.NumLocalPortsPerIlkn})
    cardInfo.EntityData.Leafs.Append("num-cos-per-port", types.YLeaf{"NumCosPerPort", cardInfo.NumCosPerPort})

    cardInfo.EntityData.YListKeys = []string {}

    return &(cardInfo.EntityData)
}

// Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer
// oir circular buffer
type Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // count. The type is interface{} with range: -2147483648..2147483647.
    Count interface{}

    // start. The type is interface{} with range: -2147483648..2147483647.
    Start interface{}

    // end. The type is interface{} with range: -2147483648..2147483647.
    End interface{}

    // fia oir info. The type is slice of
    // Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo.
    FiaOirInfo []*Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo
}

func (oirCircularBuffer *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer) GetEntityData() *types.CommonEntityData {
    oirCircularBuffer.EntityData.YFilter = oirCircularBuffer.YFilter
    oirCircularBuffer.EntityData.YangName = "oir-circular-buffer"
    oirCircularBuffer.EntityData.BundleName = "cisco_ios_xr"
    oirCircularBuffer.EntityData.ParentYangName = "card-info"
    oirCircularBuffer.EntityData.SegmentPath = "oir-circular-buffer"
    oirCircularBuffer.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/driver-information/card-info/" + oirCircularBuffer.EntityData.SegmentPath
    oirCircularBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirCircularBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirCircularBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirCircularBuffer.EntityData.Children = types.NewOrderedMap()
    oirCircularBuffer.EntityData.Children.Append("fia-oir-info", types.YChild{"FiaOirInfo", nil})
    for i := range oirCircularBuffer.FiaOirInfo {
        types.SetYListKey(oirCircularBuffer.FiaOirInfo[i], i)
        oirCircularBuffer.EntityData.Children.Append(types.GetSegmentPath(oirCircularBuffer.FiaOirInfo[i]), types.YChild{"FiaOirInfo", oirCircularBuffer.FiaOirInfo[i]})
    }
    oirCircularBuffer.EntityData.Leafs = types.NewOrderedMap()
    oirCircularBuffer.EntityData.Leafs.Append("count", types.YLeaf{"Count", oirCircularBuffer.Count})
    oirCircularBuffer.EntityData.Leafs.Append("start", types.YLeaf{"Start", oirCircularBuffer.Start})
    oirCircularBuffer.EntityData.Leafs.Append("end", types.YLeaf{"End", oirCircularBuffer.End})

    oirCircularBuffer.EntityData.YListKeys = []string {}

    return &(oirCircularBuffer.EntityData)
}

// Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo
// fia oir info
type Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (fiaOirInfo *Fia_Nodes_Node_DriverInformation_CardInfo_OirCircularBuffer_FiaOirInfo) GetEntityData() *types.CommonEntityData {
    fiaOirInfo.EntityData.YFilter = fiaOirInfo.YFilter
    fiaOirInfo.EntityData.YangName = "fia-oir-info"
    fiaOirInfo.EntityData.BundleName = "cisco_ios_xr"
    fiaOirInfo.EntityData.ParentYangName = "oir-circular-buffer"
    fiaOirInfo.EntityData.SegmentPath = "fia-oir-info" + types.AddNoKeyToken(fiaOirInfo)
    fiaOirInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/driver-information/card-info/oir-circular-buffer/" + fiaOirInfo.EntityData.SegmentPath
    fiaOirInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fiaOirInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fiaOirInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fiaOirInfo.EntityData.Children = types.NewOrderedMap()
    fiaOirInfo.EntityData.Leafs = types.NewOrderedMap()
    fiaOirInfo.EntityData.Leafs.Append("card-flag", types.YLeaf{"CardFlag", fiaOirInfo.CardFlag})
    fiaOirInfo.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", fiaOirInfo.CardType})
    fiaOirInfo.EntityData.Leafs.Append("reg-flag", types.YLeaf{"RegFlag", fiaOirInfo.RegFlag})
    fiaOirInfo.EntityData.Leafs.Append("evt-flag", types.YLeaf{"EvtFlag", fiaOirInfo.EvtFlag})
    fiaOirInfo.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", fiaOirInfo.RackNum})
    fiaOirInfo.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", fiaOirInfo.Instance})
    fiaOirInfo.EntityData.Leafs.Append("cur-card-state", types.YLeaf{"CurCardState", fiaOirInfo.CurCardState})

    fiaOirInfo.EntityData.YListKeys = []string {}

    return &(fiaOirInfo.EntityData)
}

// Fia_Nodes_Node_ClearStatistics
// Clear statistics information
type Fia_Nodes_Node_ClearStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance table for clear statistics information.
    AsicInstances Fia_Nodes_Node_ClearStatistics_AsicInstances
}

func (clearStatistics *Fia_Nodes_Node_ClearStatistics) GetEntityData() *types.CommonEntityData {
    clearStatistics.EntityData.YFilter = clearStatistics.YFilter
    clearStatistics.EntityData.YangName = "clear-statistics"
    clearStatistics.EntityData.BundleName = "cisco_ios_xr"
    clearStatistics.EntityData.ParentYangName = "node"
    clearStatistics.EntityData.SegmentPath = "clear-statistics"
    clearStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + clearStatistics.EntityData.SegmentPath
    clearStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clearStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clearStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clearStatistics.EntityData.Children = types.NewOrderedMap()
    clearStatistics.EntityData.Children.Append("asic-instances", types.YChild{"AsicInstances", &clearStatistics.AsicInstances})
    clearStatistics.EntityData.Leafs = types.NewOrderedMap()

    clearStatistics.EntityData.YListKeys = []string {}

    return &(clearStatistics.EntityData)
}

// Fia_Nodes_Node_ClearStatistics_AsicInstances
// Instance table for clear statistics
// information
type Fia_Nodes_Node_ClearStatistics_AsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic instance to be cleared. The type is slice of
    // Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance.
    AsicInstance []*Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance
}

func (asicInstances *Fia_Nodes_Node_ClearStatistics_AsicInstances) GetEntityData() *types.CommonEntityData {
    asicInstances.EntityData.YFilter = asicInstances.YFilter
    asicInstances.EntityData.YangName = "asic-instances"
    asicInstances.EntityData.BundleName = "cisco_ios_xr"
    asicInstances.EntityData.ParentYangName = "clear-statistics"
    asicInstances.EntityData.SegmentPath = "asic-instances"
    asicInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/clear-statistics/" + asicInstances.EntityData.SegmentPath
    asicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicInstances.EntityData.Children = types.NewOrderedMap()
    asicInstances.EntityData.Children.Append("asic-instance", types.YChild{"AsicInstance", nil})
    for i := range asicInstances.AsicInstance {
        asicInstances.EntityData.Children.Append(types.GetSegmentPath(asicInstances.AsicInstance[i]), types.YChild{"AsicInstance", asicInstances.AsicInstance[i]})
    }
    asicInstances.EntityData.Leafs = types.NewOrderedMap()

    asicInstances.EntityData.YListKeys = []string {}

    return &(asicInstances.EntityData)
}

// Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance
// Asic instance to be cleared
type Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Asic instance. The type is interface{} with range:
    // 0..255.
    AsicInstance interface{}

    // Clear value. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    Instance interface{}
}

func (asicInstance *Fia_Nodes_Node_ClearStatistics_AsicInstances_AsicInstance) GetEntityData() *types.CommonEntityData {
    asicInstance.EntityData.YFilter = asicInstance.YFilter
    asicInstance.EntityData.YangName = "asic-instance"
    asicInstance.EntityData.BundleName = "cisco_ios_xr"
    asicInstance.EntityData.ParentYangName = "asic-instances"
    asicInstance.EntityData.SegmentPath = "asic-instance" + types.AddKeyToken(asicInstance.AsicInstance, "asic-instance")
    asicInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/clear-statistics/asic-instances/" + asicInstance.EntityData.SegmentPath
    asicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicInstance.EntityData.Children = types.NewOrderedMap()
    asicInstance.EntityData.Leafs = types.NewOrderedMap()
    asicInstance.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicInstance.AsicInstance})
    asicInstance.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", asicInstance.Instance})

    asicInstance.EntityData.YListKeys = []string {"AsicInstance"}

    return &(asicInstance.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation
// FIA link TX information
type Fia_Nodes_Node_TxLinkInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link table for tx information.
    TxStatusOptionTable Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable
}

func (txLinkInformation *Fia_Nodes_Node_TxLinkInformation) GetEntityData() *types.CommonEntityData {
    txLinkInformation.EntityData.YFilter = txLinkInformation.YFilter
    txLinkInformation.EntityData.YangName = "tx-link-information"
    txLinkInformation.EntityData.BundleName = "cisco_ios_xr"
    txLinkInformation.EntityData.ParentYangName = "node"
    txLinkInformation.EntityData.SegmentPath = "tx-link-information"
    txLinkInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + txLinkInformation.EntityData.SegmentPath
    txLinkInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLinkInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLinkInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLinkInformation.EntityData.Children = types.NewOrderedMap()
    txLinkInformation.EntityData.Children.Append("tx-status-option-table", types.YChild{"TxStatusOptionTable", &txLinkInformation.TxStatusOptionTable})
    txLinkInformation.EntityData.Leafs = types.NewOrderedMap()

    txLinkInformation.EntityData.YListKeys = []string {}

    return &(txLinkInformation.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable
// Link table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Option: data, ctrl, all- for now none.
    TxStatusOption Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption
}

func (txStatusOptionTable *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable) GetEntityData() *types.CommonEntityData {
    txStatusOptionTable.EntityData.YFilter = txStatusOptionTable.YFilter
    txStatusOptionTable.EntityData.YangName = "tx-status-option-table"
    txStatusOptionTable.EntityData.BundleName = "cisco_ios_xr"
    txStatusOptionTable.EntityData.ParentYangName = "tx-link-information"
    txStatusOptionTable.EntityData.SegmentPath = "tx-status-option-table"
    txStatusOptionTable.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/" + txStatusOptionTable.EntityData.SegmentPath
    txStatusOptionTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txStatusOptionTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txStatusOptionTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txStatusOptionTable.EntityData.Children = types.NewOrderedMap()
    txStatusOptionTable.EntityData.Children.Append("tx-status-option", types.YChild{"TxStatusOption", &txStatusOptionTable.TxStatusOption})
    txStatusOptionTable.EntityData.Leafs = types.NewOrderedMap()

    txStatusOptionTable.EntityData.YListKeys = []string {}

    return &(txStatusOptionTable.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption
// Option: data, ctrl, all- for now none
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance table for tx information.
    TxAsicInstances Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances
}

func (txStatusOption *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption) GetEntityData() *types.CommonEntityData {
    txStatusOption.EntityData.YFilter = txStatusOption.YFilter
    txStatusOption.EntityData.YangName = "tx-status-option"
    txStatusOption.EntityData.BundleName = "cisco_ios_xr"
    txStatusOption.EntityData.ParentYangName = "tx-status-option-table"
    txStatusOption.EntityData.SegmentPath = "tx-status-option"
    txStatusOption.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/" + txStatusOption.EntityData.SegmentPath
    txStatusOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txStatusOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txStatusOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txStatusOption.EntityData.Children = types.NewOrderedMap()
    txStatusOption.EntityData.Children.Append("tx-asic-instances", types.YChild{"TxAsicInstances", &txStatusOption.TxAsicInstances})
    txStatusOption.EntityData.Leafs = types.NewOrderedMap()

    txStatusOption.EntityData.YListKeys = []string {}

    return &(txStatusOption.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances
// Instance table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance number for tx link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance.
    TxAsicInstance []*Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance
}

func (txAsicInstances *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances) GetEntityData() *types.CommonEntityData {
    txAsicInstances.EntityData.YFilter = txAsicInstances.YFilter
    txAsicInstances.EntityData.YangName = "tx-asic-instances"
    txAsicInstances.EntityData.BundleName = "cisco_ios_xr"
    txAsicInstances.EntityData.ParentYangName = "tx-status-option"
    txAsicInstances.EntityData.SegmentPath = "tx-asic-instances"
    txAsicInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/" + txAsicInstances.EntityData.SegmentPath
    txAsicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txAsicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txAsicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txAsicInstances.EntityData.Children = types.NewOrderedMap()
    txAsicInstances.EntityData.Children.Append("tx-asic-instance", types.YChild{"TxAsicInstance", nil})
    for i := range txAsicInstances.TxAsicInstance {
        txAsicInstances.EntityData.Children.Append(types.GetSegmentPath(txAsicInstances.TxAsicInstance[i]), types.YChild{"TxAsicInstance", txAsicInstances.TxAsicInstance[i]})
    }
    txAsicInstances.EntityData.Leafs = types.NewOrderedMap()

    txAsicInstances.EntityData.YListKeys = []string {}

    return &(txAsicInstances.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance
// Instance number for tx link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Transmit instance. The type is interface{} with
    // range: 0..255.
    Instance interface{}

    // Link table for tx information.
    TxLinks Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks
}

func (txAsicInstance *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance) GetEntityData() *types.CommonEntityData {
    txAsicInstance.EntityData.YFilter = txAsicInstance.YFilter
    txAsicInstance.EntityData.YangName = "tx-asic-instance"
    txAsicInstance.EntityData.BundleName = "cisco_ios_xr"
    txAsicInstance.EntityData.ParentYangName = "tx-asic-instances"
    txAsicInstance.EntityData.SegmentPath = "tx-asic-instance" + types.AddKeyToken(txAsicInstance.Instance, "instance")
    txAsicInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/" + txAsicInstance.EntityData.SegmentPath
    txAsicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txAsicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txAsicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txAsicInstance.EntityData.Children = types.NewOrderedMap()
    txAsicInstance.EntityData.Children.Append("tx-links", types.YChild{"TxLinks", &txAsicInstance.TxLinks})
    txAsicInstance.EntityData.Leafs = types.NewOrderedMap()
    txAsicInstance.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", txAsicInstance.Instance})

    txAsicInstance.EntityData.YListKeys = []string {"Instance"}

    return &(txAsicInstance.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks
// Link table for tx information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link number for tx link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink.
    TxLink []*Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink
}

func (txLinks *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks) GetEntityData() *types.CommonEntityData {
    txLinks.EntityData.YFilter = txLinks.YFilter
    txLinks.EntityData.YangName = "tx-links"
    txLinks.EntityData.BundleName = "cisco_ios_xr"
    txLinks.EntityData.ParentYangName = "tx-asic-instance"
    txLinks.EntityData.SegmentPath = "tx-links"
    txLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/" + txLinks.EntityData.SegmentPath
    txLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLinks.EntityData.Children = types.NewOrderedMap()
    txLinks.EntityData.Children.Append("tx-link", types.YChild{"TxLink", nil})
    for i := range txLinks.TxLink {
        types.SetYListKey(txLinks.TxLink[i], i)
        txLinks.EntityData.Children.Append(types.GetSegmentPath(txLinks.TxLink[i]), types.YChild{"TxLink", txLinks.TxLink[i]})
    }
    txLinks.EntityData.Leafs = types.NewOrderedMap()

    txLinks.EntityData.YListKeys = []string {}

    return &(txLinks.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink
// Link number for tx link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Start number. The type is interface{} with range: 0..47.
    StartNumber interface{}

    // End number. The type is interface{} with range: 0..47.
    EndNumber interface{}

    // Single link information. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink.
    TxLink []*Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink
}

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink) GetEntityData() *types.CommonEntityData {
    txLink.EntityData.YFilter = txLink.YFilter
    txLink.EntityData.YangName = "tx-link"
    txLink.EntityData.BundleName = "cisco_ios_xr"
    txLink.EntityData.ParentYangName = "tx-links"
    txLink.EntityData.SegmentPath = "tx-link" + types.AddNoKeyToken(txLink)
    txLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/" + txLink.EntityData.SegmentPath
    txLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLink.EntityData.Children = types.NewOrderedMap()
    txLink.EntityData.Children.Append("tx-link", types.YChild{"TxLink", nil})
    for i := range txLink.TxLink {
        txLink.EntityData.Children.Append(types.GetSegmentPath(txLink.TxLink[i]), types.YChild{"TxLink", txLink.TxLink[i]})
    }
    txLink.EntityData.Leafs = types.NewOrderedMap()
    txLink.EntityData.Leafs.Append("start-number", types.YLeaf{"StartNumber", txLink.StartNumber})
    txLink.EntityData.Leafs.Append("end-number", types.YLeaf{"EndNumber", txLink.EndNumber})

    txLink.EntityData.YListKeys = []string {}

    return &(txLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink
// Single link information
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Single Link. The type is interface{} with range:
    // 0..4294967295.
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

func (txLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink) GetEntityData() *types.CommonEntityData {
    txLink.EntityData.YFilter = txLink.YFilter
    txLink.EntityData.YangName = "tx-link"
    txLink.EntityData.BundleName = "cisco_ios_xr"
    txLink.EntityData.ParentYangName = "tx-link"
    txLink.EntityData.SegmentPath = "tx-link" + types.AddKeyToken(txLink.Link, "link")
    txLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/" + txLink.EntityData.SegmentPath
    txLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    txLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    txLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    txLink.EntityData.Children = types.NewOrderedMap()
    txLink.EntityData.Children.Append("this-link", types.YChild{"ThisLink", &txLink.ThisLink})
    txLink.EntityData.Children.Append("far-end-link", types.YChild{"FarEndLink", &txLink.FarEndLink})
    txLink.EntityData.Children.Append("stats", types.YChild{"Stats", &txLink.Stats})
    txLink.EntityData.Children.Append("history", types.YChild{"History", &txLink.History})
    txLink.EntityData.Leafs = types.NewOrderedMap()
    txLink.EntityData.Leafs.Append("link", types.YLeaf{"Link", txLink.Link})
    txLink.EntityData.Leafs.Append("speed", types.YLeaf{"Speed", txLink.Speed})
    txLink.EntityData.Leafs.Append("stage", types.YLeaf{"Stage", txLink.Stage})
    txLink.EntityData.Leafs.Append("is-link-valid", types.YLeaf{"IsLinkValid", txLink.IsLinkValid})
    txLink.EntityData.Leafs.Append("is-conf-pending", types.YLeaf{"IsConfPending", txLink.IsConfPending})
    txLink.EntityData.Leafs.Append("is-power-enabled", types.YLeaf{"IsPowerEnabled", txLink.IsPowerEnabled})
    txLink.EntityData.Leafs.Append("coeff1", types.YLeaf{"Coeff1", txLink.Coeff1})
    txLink.EntityData.Leafs.Append("coeff2", types.YLeaf{"Coeff2", txLink.Coeff2})
    txLink.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", txLink.AdminState})
    txLink.EntityData.Leafs.Append("oper-state", types.YLeaf{"OperState", txLink.OperState})
    txLink.EntityData.Leafs.Append("error-state", types.YLeaf{"ErrorState", txLink.ErrorState})
    txLink.EntityData.Leafs.Append("num-admin-shuts", types.YLeaf{"NumAdminShuts", txLink.NumAdminShuts})

    txLink.EntityData.YListKeys = []string {"Link"}

    return &(txLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink
// this link
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink struct {
    EntityData types.CommonEntityData
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

func (thisLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink) GetEntityData() *types.CommonEntityData {
    thisLink.EntityData.YFilter = thisLink.YFilter
    thisLink.EntityData.YangName = "this-link"
    thisLink.EntityData.BundleName = "cisco_ios_xr"
    thisLink.EntityData.ParentYangName = "tx-link"
    thisLink.EntityData.SegmentPath = "this-link"
    thisLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/" + thisLink.EntityData.SegmentPath
    thisLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    thisLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    thisLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    thisLink.EntityData.Children = types.NewOrderedMap()
    thisLink.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &thisLink.AsicId})
    thisLink.EntityData.Leafs = types.NewOrderedMap()
    thisLink.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", thisLink.LinkType})
    thisLink.EntityData.Leafs.Append("link-stage", types.YLeaf{"LinkStage", thisLink.LinkStage})
    thisLink.EntityData.Leafs.Append("link-num", types.YLeaf{"LinkNum", thisLink.LinkNum})
    thisLink.EntityData.Leafs.Append("phy-link-num", types.YLeaf{"PhyLinkNum", thisLink.PhyLinkNum})

    thisLink.EntityData.YListKeys = []string {}

    return &(thisLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId
// asic id
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_ThisLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "this-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/this-link/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink
// far end link
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink struct {
    EntityData types.CommonEntityData
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

func (farEndLink *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink) GetEntityData() *types.CommonEntityData {
    farEndLink.EntityData.YFilter = farEndLink.YFilter
    farEndLink.EntityData.YangName = "far-end-link"
    farEndLink.EntityData.BundleName = "cisco_ios_xr"
    farEndLink.EntityData.ParentYangName = "tx-link"
    farEndLink.EntityData.SegmentPath = "far-end-link"
    farEndLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/" + farEndLink.EntityData.SegmentPath
    farEndLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    farEndLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    farEndLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    farEndLink.EntityData.Children = types.NewOrderedMap()
    farEndLink.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &farEndLink.AsicId})
    farEndLink.EntityData.Leafs = types.NewOrderedMap()
    farEndLink.EntityData.Leafs.Append("link-type", types.YLeaf{"LinkType", farEndLink.LinkType})
    farEndLink.EntityData.Leafs.Append("link-stage", types.YLeaf{"LinkStage", farEndLink.LinkStage})
    farEndLink.EntityData.Leafs.Append("link-num", types.YLeaf{"LinkNum", farEndLink.LinkNum})
    farEndLink.EntityData.Leafs.Append("phy-link-num", types.YLeaf{"PhyLinkNum", farEndLink.PhyLinkNum})

    farEndLink.EntityData.YListKeys = []string {}

    return &(farEndLink.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId
// asic id
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_FarEndLink_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "far-end-link"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/far-end-link/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats
// stats
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // dummy. The type is interface{} with range: 0..4294967295.
    Dummy interface{}
}

func (stats *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "tx-link"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/" + stats.EntityData.SegmentPath
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", stats.Dummy})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History
// history
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // histnum. The type is interface{} with range: 0..255.
    Histnum interface{}

    // start index. The type is interface{} with range: 0..255.
    StartIndex interface{}

    // hist. The type is slice of
    // Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist.
    Hist []*Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist
}

func (history *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History) GetEntityData() *types.CommonEntityData {
    history.EntityData.YFilter = history.YFilter
    history.EntityData.YangName = "history"
    history.EntityData.BundleName = "cisco_ios_xr"
    history.EntityData.ParentYangName = "tx-link"
    history.EntityData.SegmentPath = "history"
    history.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/" + history.EntityData.SegmentPath
    history.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    history.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    history.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    history.EntityData.Children = types.NewOrderedMap()
    history.EntityData.Children.Append("hist", types.YChild{"Hist", nil})
    for i := range history.Hist {
        types.SetYListKey(history.Hist[i], i)
        history.EntityData.Children.Append(types.GetSegmentPath(history.Hist[i]), types.YChild{"Hist", history.Hist[i]})
    }
    history.EntityData.Leafs = types.NewOrderedMap()
    history.EntityData.Leafs.Append("histnum", types.YLeaf{"Histnum", history.Histnum})
    history.EntityData.Leafs.Append("start-index", types.YLeaf{"StartIndex", history.StartIndex})

    history.EntityData.YListKeys = []string {}

    return &(history.EntityData)
}

// Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist
// hist
type Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (hist *Fia_Nodes_Node_TxLinkInformation_TxStatusOptionTable_TxStatusOption_TxAsicInstances_TxAsicInstance_TxLinks_TxLink_TxLink_History_Hist) GetEntityData() *types.CommonEntityData {
    hist.EntityData.YFilter = hist.YFilter
    hist.EntityData.YangName = "hist"
    hist.EntityData.BundleName = "cisco_ios_xr"
    hist.EntityData.ParentYangName = "history"
    hist.EntityData.SegmentPath = "hist" + types.AddNoKeyToken(hist)
    hist.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/tx-link-information/tx-status-option-table/tx-status-option/tx-asic-instances/tx-asic-instance/tx-links/tx-link/tx-link/history/" + hist.EntityData.SegmentPath
    hist.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hist.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hist.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hist.EntityData.Children = types.NewOrderedMap()
    hist.EntityData.Leafs = types.NewOrderedMap()
    hist.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", hist.AdminState})
    hist.EntityData.Leafs.Append("oper-state", types.YLeaf{"OperState", hist.OperState})
    hist.EntityData.Leafs.Append("error-state", types.YLeaf{"ErrorState", hist.ErrorState})
    hist.EntityData.Leafs.Append("timestamp", types.YLeaf{"Timestamp", hist.Timestamp})
    hist.EntityData.Leafs.Append("reasons", types.YLeaf{"Reasons", hist.Reasons})

    hist.EntityData.YListKeys = []string {}

    return &(hist.EntityData)
}

// Fia_Nodes_Node_DiagShell
// FIA diag shell information
type Fia_Nodes_Node_DiagShell struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unit table for diag shell.
    DiagShellUnits Fia_Nodes_Node_DiagShell_DiagShellUnits
}

func (diagShell *Fia_Nodes_Node_DiagShell) GetEntityData() *types.CommonEntityData {
    diagShell.EntityData.YFilter = diagShell.YFilter
    diagShell.EntityData.YangName = "diag-shell"
    diagShell.EntityData.BundleName = "cisco_ios_xr"
    diagShell.EntityData.ParentYangName = "node"
    diagShell.EntityData.SegmentPath = "diag-shell"
    diagShell.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + diagShell.EntityData.SegmentPath
    diagShell.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagShell.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagShell.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagShell.EntityData.Children = types.NewOrderedMap()
    diagShell.EntityData.Children.Append("diag-shell-units", types.YChild{"DiagShellUnits", &diagShell.DiagShellUnits})
    diagShell.EntityData.Leafs = types.NewOrderedMap()

    diagShell.EntityData.YListKeys = []string {}

    return &(diagShell.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits
// Unit table for diag shell
type Fia_Nodes_Node_DiagShell_DiagShellUnits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unit number for diag shell statistics. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit.
    DiagShellUnit []*Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit
}

func (diagShellUnits *Fia_Nodes_Node_DiagShell_DiagShellUnits) GetEntityData() *types.CommonEntityData {
    diagShellUnits.EntityData.YFilter = diagShellUnits.YFilter
    diagShellUnits.EntityData.YangName = "diag-shell-units"
    diagShellUnits.EntityData.BundleName = "cisco_ios_xr"
    diagShellUnits.EntityData.ParentYangName = "diag-shell"
    diagShellUnits.EntityData.SegmentPath = "diag-shell-units"
    diagShellUnits.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/diag-shell/" + diagShellUnits.EntityData.SegmentPath
    diagShellUnits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagShellUnits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagShellUnits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagShellUnits.EntityData.Children = types.NewOrderedMap()
    diagShellUnits.EntityData.Children.Append("diag-shell-unit", types.YChild{"DiagShellUnit", nil})
    for i := range diagShellUnits.DiagShellUnit {
        diagShellUnits.EntityData.Children.Append(types.GetSegmentPath(diagShellUnits.DiagShellUnit[i]), types.YChild{"DiagShellUnit", diagShellUnits.DiagShellUnit[i]})
    }
    diagShellUnits.EntityData.Leafs = types.NewOrderedMap()

    diagShellUnits.EntityData.YListKeys = []string {}

    return &(diagShellUnits.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit
// Unit number for diag shell statistics
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Unit number. The type is interface{} with range:
    // 0..63.
    Unit interface{}

    // Command table for diag shell.
    Commands Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands
}

func (diagShellUnit *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit) GetEntityData() *types.CommonEntityData {
    diagShellUnit.EntityData.YFilter = diagShellUnit.YFilter
    diagShellUnit.EntityData.YangName = "diag-shell-unit"
    diagShellUnit.EntityData.BundleName = "cisco_ios_xr"
    diagShellUnit.EntityData.ParentYangName = "diag-shell-units"
    diagShellUnit.EntityData.SegmentPath = "diag-shell-unit" + types.AddKeyToken(diagShellUnit.Unit, "unit")
    diagShellUnit.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/diag-shell/diag-shell-units/" + diagShellUnit.EntityData.SegmentPath
    diagShellUnit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diagShellUnit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diagShellUnit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diagShellUnit.EntityData.Children = types.NewOrderedMap()
    diagShellUnit.EntityData.Children.Append("commands", types.YChild{"Commands", &diagShellUnit.Commands})
    diagShellUnit.EntityData.Leafs = types.NewOrderedMap()
    diagShellUnit.EntityData.Leafs.Append("unit", types.YLeaf{"Unit", diagShellUnit.Unit})

    diagShellUnit.EntityData.YListKeys = []string {"Unit"}

    return &(diagShellUnit.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands
// Command table for diag shell
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Command for diag shell statistics. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command.
    Command []*Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command
}

func (commands *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands) GetEntityData() *types.CommonEntityData {
    commands.EntityData.YFilter = commands.YFilter
    commands.EntityData.YangName = "commands"
    commands.EntityData.BundleName = "cisco_ios_xr"
    commands.EntityData.ParentYangName = "diag-shell-unit"
    commands.EntityData.SegmentPath = "commands"
    commands.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/diag-shell/diag-shell-units/diag-shell-unit/" + commands.EntityData.SegmentPath
    commands.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    commands.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    commands.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    commands.EntityData.Children = types.NewOrderedMap()
    commands.EntityData.Children.Append("command", types.YChild{"Command", nil})
    for i := range commands.Command {
        commands.EntityData.Children.Append(types.GetSegmentPath(commands.Command[i]), types.YChild{"Command", commands.Command[i]})
    }
    commands.EntityData.Leafs = types.NewOrderedMap()

    commands.EntityData.YListKeys = []string {}

    return &(commands.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command
// Command for diag shell statistics
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Shell command. The type is string.
    Cmd interface{}

    // Added to support datalist. The type is slice of
    // Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output.
    Output []*Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output
}

func (command *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command) GetEntityData() *types.CommonEntityData {
    command.EntityData.YFilter = command.YFilter
    command.EntityData.YangName = "command"
    command.EntityData.BundleName = "cisco_ios_xr"
    command.EntityData.ParentYangName = "commands"
    command.EntityData.SegmentPath = "command" + types.AddKeyToken(command.Cmd, "cmd")
    command.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/diag-shell/diag-shell-units/diag-shell-unit/commands/" + command.EntityData.SegmentPath
    command.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    command.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    command.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    command.EntityData.Children = types.NewOrderedMap()
    command.EntityData.Children.Append("output", types.YChild{"Output", nil})
    for i := range command.Output {
        command.EntityData.Children.Append(types.GetSegmentPath(command.Output[i]), types.YChild{"Output", command.Output[i]})
    }
    command.EntityData.Leafs = types.NewOrderedMap()
    command.EntityData.Leafs.Append("cmd", types.YLeaf{"Cmd", command.Cmd})

    command.EntityData.YListKeys = []string {"Cmd"}

    return &(command.EntityData)
}

// Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output
// Added to support datalist
type Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. First line. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Output interface{}

    // output xr. The type is string.
    OutputXr interface{}
}

func (output *Fia_Nodes_Node_DiagShell_DiagShellUnits_DiagShellUnit_Commands_Command_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xr"
    output.EntityData.ParentYangName = "command"
    output.EntityData.SegmentPath = "output" + types.AddKeyToken(output.Output, "output")
    output.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/diag-shell/diag-shell-units/diag-shell-unit/commands/command/" + output.EntityData.SegmentPath
    output.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("output", types.YLeaf{"Output", output.Output})
    output.EntityData.Leafs.Append("output-xr", types.YLeaf{"OutputXr", output.OutputXr})

    output.EntityData.YListKeys = []string {"Output"}

    return &(output.EntityData)
}

// Fia_Nodes_Node_OirHistory
// FIA operational data of oir history
type Fia_Nodes_Node_OirHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flag table for history.
    Flags Fia_Nodes_Node_OirHistory_Flags
}

func (oirHistory *Fia_Nodes_Node_OirHistory) GetEntityData() *types.CommonEntityData {
    oirHistory.EntityData.YFilter = oirHistory.YFilter
    oirHistory.EntityData.YangName = "oir-history"
    oirHistory.EntityData.BundleName = "cisco_ios_xr"
    oirHistory.EntityData.ParentYangName = "node"
    oirHistory.EntityData.SegmentPath = "oir-history"
    oirHistory.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + oirHistory.EntityData.SegmentPath
    oirHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirHistory.EntityData.Children = types.NewOrderedMap()
    oirHistory.EntityData.Children.Append("flags", types.YChild{"Flags", &oirHistory.Flags})
    oirHistory.EntityData.Leafs = types.NewOrderedMap()

    oirHistory.EntityData.YListKeys = []string {}

    return &(oirHistory.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags
// Flag table for history
type Fia_Nodes_Node_OirHistory_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flag value for physical location. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag.
    Flag []*Fia_Nodes_Node_OirHistory_Flags_Flag
}

func (flags *Fia_Nodes_Node_OirHistory_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xr"
    flags.EntityData.ParentYangName = "oir-history"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/" + flags.EntityData.SegmentPath
    flags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Children.Append("flag", types.YChild{"Flag", nil})
    for i := range flags.Flag {
        flags.EntityData.Children.Append(types.GetSegmentPath(flags.Flag[i]), types.YChild{"Flag", flags.Flag[i]})
    }
    flags.EntityData.Leafs = types.NewOrderedMap()

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag
// Flag value for physical location
type Fia_Nodes_Node_OirHistory_Flags_Flag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Flag value. The type is interface{} with range:
    // 0..4294967295.
    Flag interface{}

    // Slot table for history.
    Slots Fia_Nodes_Node_OirHistory_Flags_Flag_Slots
}

func (flag *Fia_Nodes_Node_OirHistory_Flags_Flag) GetEntityData() *types.CommonEntityData {
    flag.EntityData.YFilter = flag.YFilter
    flag.EntityData.YangName = "flag"
    flag.EntityData.BundleName = "cisco_ios_xr"
    flag.EntityData.ParentYangName = "flags"
    flag.EntityData.SegmentPath = "flag" + types.AddKeyToken(flag.Flag, "flag")
    flag.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/" + flag.EntityData.SegmentPath
    flag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flag.EntityData.Children = types.NewOrderedMap()
    flag.EntityData.Children.Append("slots", types.YChild{"Slots", &flag.Slots})
    flag.EntityData.Leafs = types.NewOrderedMap()
    flag.EntityData.Leafs.Append("flag", types.YLeaf{"Flag", flag.Flag})

    flag.EntityData.YListKeys = []string {"Flag"}

    return &(flag.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots
// Slot table for history
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Slot number for getting history. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot.
    Slot []*Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot
}

func (slots *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots) GetEntityData() *types.CommonEntityData {
    slots.EntityData.YFilter = slots.YFilter
    slots.EntityData.YangName = "slots"
    slots.EntityData.BundleName = "cisco_ios_xr"
    slots.EntityData.ParentYangName = "flag"
    slots.EntityData.SegmentPath = "slots"
    slots.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/" + slots.EntityData.SegmentPath
    slots.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slots.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slots.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slots.EntityData.Children = types.NewOrderedMap()
    slots.EntityData.Children.Append("slot", types.YChild{"Slot", nil})
    for i := range slots.Slot {
        slots.EntityData.Children.Append(types.GetSegmentPath(slots.Slot[i]), types.YChild{"Slot", slots.Slot[i]})
    }
    slots.EntityData.Leafs = types.NewOrderedMap()

    slots.EntityData.YListKeys = []string {}

    return &(slots.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot
// Slot number for getting history
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Slot number. The type is interface{} with range:
    // 0..4294967295.
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

    // all wb insync. The type is bool.
    AllWbInsync interface{}

    // all wb insync since. The type is interface{} with range: 0..4294967295.
    AllWbInsyncSince interface{}

    // all startup wb insync. The type is bool.
    AllStartupWbInsync interface{}

    // planeA bitmap. The type is interface{} with range: 0..4294967295.
    PlaneABitmap interface{}

    // planeB bitmap. The type is interface{} with range: 0..4294967295.
    PlaneBBitmap interface{}

    // device info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo.
    DeviceInfo []*Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo

    // card info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo.
    CardInfo []*Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo
}

func (slot *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot) GetEntityData() *types.CommonEntityData {
    slot.EntityData.YFilter = slot.YFilter
    slot.EntityData.YangName = "slot"
    slot.EntityData.BundleName = "cisco_ios_xr"
    slot.EntityData.ParentYangName = "slots"
    slot.EntityData.SegmentPath = "slot" + types.AddKeyToken(slot.Slot, "slot")
    slot.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/slots/" + slot.EntityData.SegmentPath
    slot.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slot.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slot.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slot.EntityData.Children = types.NewOrderedMap()
    slot.EntityData.Children.Append("device-info", types.YChild{"DeviceInfo", nil})
    for i := range slot.DeviceInfo {
        types.SetYListKey(slot.DeviceInfo[i], i)
        slot.EntityData.Children.Append(types.GetSegmentPath(slot.DeviceInfo[i]), types.YChild{"DeviceInfo", slot.DeviceInfo[i]})
    }
    slot.EntityData.Children.Append("card-info", types.YChild{"CardInfo", nil})
    for i := range slot.CardInfo {
        types.SetYListKey(slot.CardInfo[i], i)
        slot.EntityData.Children.Append(types.GetSegmentPath(slot.CardInfo[i]), types.YChild{"CardInfo", slot.CardInfo[i]})
    }
    slot.EntityData.Leafs = types.NewOrderedMap()
    slot.EntityData.Leafs.Append("slot", types.YLeaf{"Slot", slot.Slot})
    slot.EntityData.Leafs.Append("drv-version", types.YLeaf{"DrvVersion", slot.DrvVersion})
    slot.EntityData.Leafs.Append("coeff-major-rev", types.YLeaf{"CoeffMajorRev", slot.CoeffMajorRev})
    slot.EntityData.Leafs.Append("coeff-minor-rev", types.YLeaf{"CoeffMinorRev", slot.CoeffMinorRev})
    slot.EntityData.Leafs.Append("functional-role", types.YLeaf{"FunctionalRole", slot.FunctionalRole})
    slot.EntityData.Leafs.Append("issu-role", types.YLeaf{"IssuRole", slot.IssuRole})
    slot.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", slot.NodeId})
    slot.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", slot.RackType})
    slot.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", slot.RackNum})
    slot.EntityData.Leafs.Append("is-driver-ready", types.YLeaf{"IsDriverReady", slot.IsDriverReady})
    slot.EntityData.Leafs.Append("card-avail-mask", types.YLeaf{"CardAvailMask", slot.CardAvailMask})
    slot.EntityData.Leafs.Append("asic-avail-mask", types.YLeaf{"AsicAvailMask", slot.AsicAvailMask})
    slot.EntityData.Leafs.Append("exp-asic-avail-mask", types.YLeaf{"ExpAsicAvailMask", slot.ExpAsicAvailMask})
    slot.EntityData.Leafs.Append("ucmc-ratio", types.YLeaf{"UcmcRatio", slot.UcmcRatio})
    slot.EntityData.Leafs.Append("asic-oper-notify-to-fsdb-pending-bmap", types.YLeaf{"AsicOperNotifyToFsdbPendingBmap", slot.AsicOperNotifyToFsdbPendingBmap})
    slot.EntityData.Leafs.Append("is-full-fgid-download-req", types.YLeaf{"IsFullFgidDownloadReq", slot.IsFullFgidDownloadReq})
    slot.EntityData.Leafs.Append("is-fgid-download-in-progress", types.YLeaf{"IsFgidDownloadInProgress", slot.IsFgidDownloadInProgress})
    slot.EntityData.Leafs.Append("is-fgid-download-completed", types.YLeaf{"IsFgidDownloadCompleted", slot.IsFgidDownloadCompleted})
    slot.EntityData.Leafs.Append("fsdb-conn-active", types.YLeaf{"FsdbConnActive", slot.FsdbConnActive})
    slot.EntityData.Leafs.Append("fgid-conn-active", types.YLeaf{"FgidConnActive", slot.FgidConnActive})
    slot.EntityData.Leafs.Append("issu-mgr-conn-active", types.YLeaf{"IssuMgrConnActive", slot.IssuMgrConnActive})
    slot.EntityData.Leafs.Append("fsdb-reg-active", types.YLeaf{"FsdbRegActive", slot.FsdbRegActive})
    slot.EntityData.Leafs.Append("fgid-reg-active", types.YLeaf{"FgidRegActive", slot.FgidRegActive})
    slot.EntityData.Leafs.Append("issu-mgr-reg-active", types.YLeaf{"IssuMgrRegActive", slot.IssuMgrRegActive})
    slot.EntityData.Leafs.Append("num-pm-conn-reqs", types.YLeaf{"NumPmConnReqs", slot.NumPmConnReqs})
    slot.EntityData.Leafs.Append("num-fsdb-conn-reqs", types.YLeaf{"NumFsdbConnReqs", slot.NumFsdbConnReqs})
    slot.EntityData.Leafs.Append("num-fgid-conn-reqs", types.YLeaf{"NumFgidConnReqs", slot.NumFgidConnReqs})
    slot.EntityData.Leafs.Append("num-fstats-conn-reqs", types.YLeaf{"NumFstatsConnReqs", slot.NumFstatsConnReqs})
    slot.EntityData.Leafs.Append("num-cm-conn-reqs", types.YLeaf{"NumCmConnReqs", slot.NumCmConnReqs})
    slot.EntityData.Leafs.Append("num-issu-mgr-conn-reqs", types.YLeaf{"NumIssuMgrConnReqs", slot.NumIssuMgrConnReqs})
    slot.EntityData.Leafs.Append("num-peer-fia-conn-reqs", types.YLeaf{"NumPeerFiaConnReqs", slot.NumPeerFiaConnReqs})
    slot.EntityData.Leafs.Append("is-gaspp-registered", types.YLeaf{"IsGasppRegistered", slot.IsGasppRegistered})
    slot.EntityData.Leafs.Append("is-cih-registered", types.YLeaf{"IsCihRegistered", slot.IsCihRegistered})
    slot.EntityData.Leafs.Append("drvr-initial-startup-timestamp", types.YLeaf{"DrvrInitialStartupTimestamp", slot.DrvrInitialStartupTimestamp})
    slot.EntityData.Leafs.Append("drvr-current-startup-timestamp", types.YLeaf{"DrvrCurrentStartupTimestamp", slot.DrvrCurrentStartupTimestamp})
    slot.EntityData.Leafs.Append("num-intf-ports", types.YLeaf{"NumIntfPorts", slot.NumIntfPorts})
    slot.EntityData.Leafs.Append("uc-weight", types.YLeaf{"UcWeight", slot.UcWeight})
    slot.EntityData.Leafs.Append("respawn-count", types.YLeaf{"RespawnCount", slot.RespawnCount})
    slot.EntityData.Leafs.Append("total-asics", types.YLeaf{"TotalAsics", slot.TotalAsics})
    slot.EntityData.Leafs.Append("issu-ready-ntfy-pending", types.YLeaf{"IssuReadyNtfyPending", slot.IssuReadyNtfyPending})
    slot.EntityData.Leafs.Append("issu-abort-sent", types.YLeaf{"IssuAbortSent", slot.IssuAbortSent})
    slot.EntityData.Leafs.Append("issu-abort-rcvd", types.YLeaf{"IssuAbortRcvd", slot.IssuAbortRcvd})
    slot.EntityData.Leafs.Append("fabric-mode", types.YLeaf{"FabricMode", slot.FabricMode})
    slot.EntityData.Leafs.Append("fc-mode", types.YLeaf{"FcMode", slot.FcMode})
    slot.EntityData.Leafs.Append("board-rev-id", types.YLeaf{"BoardRevId", slot.BoardRevId})
    slot.EntityData.Leafs.Append("all-wb-insync", types.YLeaf{"AllWbInsync", slot.AllWbInsync})
    slot.EntityData.Leafs.Append("all-wb-insync-since", types.YLeaf{"AllWbInsyncSince", slot.AllWbInsyncSince})
    slot.EntityData.Leafs.Append("all-startup-wb-insync", types.YLeaf{"AllStartupWbInsync", slot.AllStartupWbInsync})
    slot.EntityData.Leafs.Append("plane-a-bitmap", types.YLeaf{"PlaneABitmap", slot.PlaneABitmap})
    slot.EntityData.Leafs.Append("plane-b-bitmap", types.YLeaf{"PlaneBBitmap", slot.PlaneBBitmap})

    slot.EntityData.YListKeys = []string {"Slot"}

    return &(slot.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo
// device info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // startup wb mtime str. The type is string.
    StartupWbMtimeStr interface{}

    // startup wb outof sync. The type is bool.
    StartupWbOutofSync interface{}

    // local wb sync end str. The type is string.
    LocalWbSyncEndStr interface{}

    // remote wb sync end str. The type is string.
    RemoteWbSyncEndStr interface{}

    // local wb sync pending. The type is bool.
    LocalWbSyncPending interface{}

    // sdk delay msec. The type is interface{} with range: 0..4294967295.
    SdkDelayMsec interface{}

    // asic id.
    AsicId Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId
}

func (deviceInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo) GetEntityData() *types.CommonEntityData {
    deviceInfo.EntityData.YFilter = deviceInfo.YFilter
    deviceInfo.EntityData.YangName = "device-info"
    deviceInfo.EntityData.BundleName = "cisco_ios_xr"
    deviceInfo.EntityData.ParentYangName = "slot"
    deviceInfo.EntityData.SegmentPath = "device-info" + types.AddNoKeyToken(deviceInfo)
    deviceInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/slots/slot/" + deviceInfo.EntityData.SegmentPath
    deviceInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deviceInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deviceInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deviceInfo.EntityData.Children = types.NewOrderedMap()
    deviceInfo.EntityData.Children.Append("asic-id", types.YChild{"AsicId", &deviceInfo.AsicId})
    deviceInfo.EntityData.Leafs = types.NewOrderedMap()
    deviceInfo.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", deviceInfo.IsValid})
    deviceInfo.EntityData.Leafs.Append("fapid", types.YLeaf{"Fapid", deviceInfo.Fapid})
    deviceInfo.EntityData.Leafs.Append("hotplug-event", types.YLeaf{"HotplugEvent", deviceInfo.HotplugEvent})
    deviceInfo.EntityData.Leafs.Append("slice-state", types.YLeaf{"SliceState", deviceInfo.SliceState})
    deviceInfo.EntityData.Leafs.Append("admin-state", types.YLeaf{"AdminState", deviceInfo.AdminState})
    deviceInfo.EntityData.Leafs.Append("oper-state", types.YLeaf{"OperState", deviceInfo.OperState})
    deviceInfo.EntityData.Leafs.Append("asic-state", types.YLeaf{"AsicState", deviceInfo.AsicState})
    deviceInfo.EntityData.Leafs.Append("last-init-cause", types.YLeaf{"LastInitCause", deviceInfo.LastInitCause})
    deviceInfo.EntityData.Leafs.Append("num-pon-resets", types.YLeaf{"NumPonResets", deviceInfo.NumPonResets})
    deviceInfo.EntityData.Leafs.Append("num-hard-resets", types.YLeaf{"NumHardResets", deviceInfo.NumHardResets})
    deviceInfo.EntityData.Leafs.Append("local-switch-state", types.YLeaf{"LocalSwitchState", deviceInfo.LocalSwitchState})
    deviceInfo.EntityData.Leafs.Append("startup-wb-mtime-str", types.YLeaf{"StartupWbMtimeStr", deviceInfo.StartupWbMtimeStr})
    deviceInfo.EntityData.Leafs.Append("startup-wb-outof-sync", types.YLeaf{"StartupWbOutofSync", deviceInfo.StartupWbOutofSync})
    deviceInfo.EntityData.Leafs.Append("local-wb-sync-end-str", types.YLeaf{"LocalWbSyncEndStr", deviceInfo.LocalWbSyncEndStr})
    deviceInfo.EntityData.Leafs.Append("remote-wb-sync-end-str", types.YLeaf{"RemoteWbSyncEndStr", deviceInfo.RemoteWbSyncEndStr})
    deviceInfo.EntityData.Leafs.Append("local-wb-sync-pending", types.YLeaf{"LocalWbSyncPending", deviceInfo.LocalWbSyncPending})
    deviceInfo.EntityData.Leafs.Append("sdk-delay-msec", types.YLeaf{"SdkDelayMsec", deviceInfo.SdkDelayMsec})

    deviceInfo.EntityData.YListKeys = []string {}

    return &(deviceInfo.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId
// asic id
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId struct {
    EntityData types.CommonEntityData
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

func (asicId *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_DeviceInfo_AsicId) GetEntityData() *types.CommonEntityData {
    asicId.EntityData.YFilter = asicId.YFilter
    asicId.EntityData.YangName = "asic-id"
    asicId.EntityData.BundleName = "cisco_ios_xr"
    asicId.EntityData.ParentYangName = "device-info"
    asicId.EntityData.SegmentPath = "asic-id"
    asicId.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/slots/slot/device-info/" + asicId.EntityData.SegmentPath
    asicId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicId.EntityData.Children = types.NewOrderedMap()
    asicId.EntityData.Leafs = types.NewOrderedMap()
    asicId.EntityData.Leafs.Append("rack-type", types.YLeaf{"RackType", asicId.RackType})
    asicId.EntityData.Leafs.Append("asic-type", types.YLeaf{"AsicType", asicId.AsicType})
    asicId.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", asicId.RackNum})
    asicId.EntityData.Leafs.Append("slot-num", types.YLeaf{"SlotNum", asicId.SlotNum})
    asicId.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", asicId.AsicInstance})

    asicId.EntityData.YListKeys = []string {}

    return &(asicId.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo
// card info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (cardInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo) GetEntityData() *types.CommonEntityData {
    cardInfo.EntityData.YFilter = cardInfo.YFilter
    cardInfo.EntityData.YangName = "card-info"
    cardInfo.EntityData.BundleName = "cisco_ios_xr"
    cardInfo.EntityData.ParentYangName = "slot"
    cardInfo.EntityData.SegmentPath = "card-info" + types.AddNoKeyToken(cardInfo)
    cardInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/slots/slot/" + cardInfo.EntityData.SegmentPath
    cardInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cardInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cardInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cardInfo.EntityData.Children = types.NewOrderedMap()
    cardInfo.EntityData.Children.Append("oir-circular-buffer", types.YChild{"OirCircularBuffer", &cardInfo.OirCircularBuffer})
    cardInfo.EntityData.Leafs = types.NewOrderedMap()
    cardInfo.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", cardInfo.CardType})
    cardInfo.EntityData.Leafs.Append("card-name", types.YLeaf{"CardName", cardInfo.CardName})
    cardInfo.EntityData.Leafs.Append("slot-no", types.YLeaf{"SlotNo", cardInfo.SlotNo})
    cardInfo.EntityData.Leafs.Append("card-flag", types.YLeaf{"CardFlag", cardInfo.CardFlag})
    cardInfo.EntityData.Leafs.Append("evt-flag", types.YLeaf{"EvtFlag", cardInfo.EvtFlag})
    cardInfo.EntityData.Leafs.Append("reg-flag", types.YLeaf{"RegFlag", cardInfo.RegFlag})
    cardInfo.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", cardInfo.Instance})
    cardInfo.EntityData.Leafs.Append("card-state", types.YLeaf{"CardState", cardInfo.CardState})
    cardInfo.EntityData.Leafs.Append("exp-num-asics", types.YLeaf{"ExpNumAsics", cardInfo.ExpNumAsics})
    cardInfo.EntityData.Leafs.Append("exp-num-asics-per-fsdb", types.YLeaf{"ExpNumAsicsPerFsdb", cardInfo.ExpNumAsicsPerFsdb})
    cardInfo.EntityData.Leafs.Append("is-powered", types.YLeaf{"IsPowered", cardInfo.IsPowered})
    cardInfo.EntityData.Leafs.Append("cxp-avail-bitmap", types.YLeaf{"CxpAvailBitmap", cardInfo.CxpAvailBitmap})
    cardInfo.EntityData.Leafs.Append("num-ilkns-per-asic", types.YLeaf{"NumIlknsPerAsic", cardInfo.NumIlknsPerAsic})
    cardInfo.EntityData.Leafs.Append("num-local-ports-per-ilkn", types.YLeaf{"NumLocalPortsPerIlkn", cardInfo.NumLocalPortsPerIlkn})
    cardInfo.EntityData.Leafs.Append("num-cos-per-port", types.YLeaf{"NumCosPerPort", cardInfo.NumCosPerPort})

    cardInfo.EntityData.YListKeys = []string {}

    return &(cardInfo.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer
// oir circular buffer
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // count. The type is interface{} with range: -2147483648..2147483647.
    Count interface{}

    // start. The type is interface{} with range: -2147483648..2147483647.
    Start interface{}

    // end. The type is interface{} with range: -2147483648..2147483647.
    End interface{}

    // fia oir info. The type is slice of
    // Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo.
    FiaOirInfo []*Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo
}

func (oirCircularBuffer *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer) GetEntityData() *types.CommonEntityData {
    oirCircularBuffer.EntityData.YFilter = oirCircularBuffer.YFilter
    oirCircularBuffer.EntityData.YangName = "oir-circular-buffer"
    oirCircularBuffer.EntityData.BundleName = "cisco_ios_xr"
    oirCircularBuffer.EntityData.ParentYangName = "card-info"
    oirCircularBuffer.EntityData.SegmentPath = "oir-circular-buffer"
    oirCircularBuffer.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/slots/slot/card-info/" + oirCircularBuffer.EntityData.SegmentPath
    oirCircularBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oirCircularBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oirCircularBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oirCircularBuffer.EntityData.Children = types.NewOrderedMap()
    oirCircularBuffer.EntityData.Children.Append("fia-oir-info", types.YChild{"FiaOirInfo", nil})
    for i := range oirCircularBuffer.FiaOirInfo {
        types.SetYListKey(oirCircularBuffer.FiaOirInfo[i], i)
        oirCircularBuffer.EntityData.Children.Append(types.GetSegmentPath(oirCircularBuffer.FiaOirInfo[i]), types.YChild{"FiaOirInfo", oirCircularBuffer.FiaOirInfo[i]})
    }
    oirCircularBuffer.EntityData.Leafs = types.NewOrderedMap()
    oirCircularBuffer.EntityData.Leafs.Append("count", types.YLeaf{"Count", oirCircularBuffer.Count})
    oirCircularBuffer.EntityData.Leafs.Append("start", types.YLeaf{"Start", oirCircularBuffer.Start})
    oirCircularBuffer.EntityData.Leafs.Append("end", types.YLeaf{"End", oirCircularBuffer.End})

    oirCircularBuffer.EntityData.YListKeys = []string {}

    return &(oirCircularBuffer.EntityData)
}

// Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo
// fia oir info
type Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (fiaOirInfo *Fia_Nodes_Node_OirHistory_Flags_Flag_Slots_Slot_CardInfo_OirCircularBuffer_FiaOirInfo) GetEntityData() *types.CommonEntityData {
    fiaOirInfo.EntityData.YFilter = fiaOirInfo.YFilter
    fiaOirInfo.EntityData.YangName = "fia-oir-info"
    fiaOirInfo.EntityData.BundleName = "cisco_ios_xr"
    fiaOirInfo.EntityData.ParentYangName = "oir-circular-buffer"
    fiaOirInfo.EntityData.SegmentPath = "fia-oir-info" + types.AddNoKeyToken(fiaOirInfo)
    fiaOirInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/oir-history/flags/flag/slots/slot/card-info/oir-circular-buffer/" + fiaOirInfo.EntityData.SegmentPath
    fiaOirInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fiaOirInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fiaOirInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fiaOirInfo.EntityData.Children = types.NewOrderedMap()
    fiaOirInfo.EntityData.Leafs = types.NewOrderedMap()
    fiaOirInfo.EntityData.Leafs.Append("card-flag", types.YLeaf{"CardFlag", fiaOirInfo.CardFlag})
    fiaOirInfo.EntityData.Leafs.Append("card-type", types.YLeaf{"CardType", fiaOirInfo.CardType})
    fiaOirInfo.EntityData.Leafs.Append("reg-flag", types.YLeaf{"RegFlag", fiaOirInfo.RegFlag})
    fiaOirInfo.EntityData.Leafs.Append("evt-flag", types.YLeaf{"EvtFlag", fiaOirInfo.EvtFlag})
    fiaOirInfo.EntityData.Leafs.Append("rack-num", types.YLeaf{"RackNum", fiaOirInfo.RackNum})
    fiaOirInfo.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", fiaOirInfo.Instance})
    fiaOirInfo.EntityData.Leafs.Append("cur-card-state", types.YLeaf{"CurCardState", fiaOirInfo.CurCardState})

    fiaOirInfo.EntityData.YListKeys = []string {}

    return &(fiaOirInfo.EntityData)
}

// Fia_Nodes_Node_AsicStatistics
// FIA asic statistics information
type Fia_Nodes_Node_AsicStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Instance table for statistics.
    StatisticsAsicInstances Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances
}

func (asicStatistics *Fia_Nodes_Node_AsicStatistics) GetEntityData() *types.CommonEntityData {
    asicStatistics.EntityData.YFilter = asicStatistics.YFilter
    asicStatistics.EntityData.YangName = "asic-statistics"
    asicStatistics.EntityData.BundleName = "cisco_ios_xr"
    asicStatistics.EntityData.ParentYangName = "node"
    asicStatistics.EntityData.SegmentPath = "asic-statistics"
    asicStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/" + asicStatistics.EntityData.SegmentPath
    asicStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asicStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asicStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asicStatistics.EntityData.Children = types.NewOrderedMap()
    asicStatistics.EntityData.Children.Append("statistics-asic-instances", types.YChild{"StatisticsAsicInstances", &asicStatistics.StatisticsAsicInstances})
    asicStatistics.EntityData.Leafs = types.NewOrderedMap()

    asicStatistics.EntityData.YListKeys = []string {}

    return &(asicStatistics.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances
// Instance table for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Asic instance for statistics. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance.
    StatisticsAsicInstance []*Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance
}

func (statisticsAsicInstances *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances) GetEntityData() *types.CommonEntityData {
    statisticsAsicInstances.EntityData.YFilter = statisticsAsicInstances.YFilter
    statisticsAsicInstances.EntityData.YangName = "statistics-asic-instances"
    statisticsAsicInstances.EntityData.BundleName = "cisco_ios_xr"
    statisticsAsicInstances.EntityData.ParentYangName = "asic-statistics"
    statisticsAsicInstances.EntityData.SegmentPath = "statistics-asic-instances"
    statisticsAsicInstances.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/" + statisticsAsicInstances.EntityData.SegmentPath
    statisticsAsicInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsAsicInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsAsicInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsAsicInstances.EntityData.Children = types.NewOrderedMap()
    statisticsAsicInstances.EntityData.Children.Append("statistics-asic-instance", types.YChild{"StatisticsAsicInstance", nil})
    for i := range statisticsAsicInstances.StatisticsAsicInstance {
        statisticsAsicInstances.EntityData.Children.Append(types.GetSegmentPath(statisticsAsicInstances.StatisticsAsicInstance[i]), types.YChild{"StatisticsAsicInstance", statisticsAsicInstances.StatisticsAsicInstance[i]})
    }
    statisticsAsicInstances.EntityData.Leafs = types.NewOrderedMap()

    statisticsAsicInstances.EntityData.YListKeys = []string {}

    return &(statisticsAsicInstances.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance
// Asic instance for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Asic instance. The type is interface{} with range:
    // 0..255.
    Instance interface{}

    // Packet Byte Counter for a Asic.
    PbcStatistics Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics

    // Statistics of FMAC.
    FmacStatistics Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics
}

func (statisticsAsicInstance *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance) GetEntityData() *types.CommonEntityData {
    statisticsAsicInstance.EntityData.YFilter = statisticsAsicInstance.YFilter
    statisticsAsicInstance.EntityData.YangName = "statistics-asic-instance"
    statisticsAsicInstance.EntityData.BundleName = "cisco_ios_xr"
    statisticsAsicInstance.EntityData.ParentYangName = "statistics-asic-instances"
    statisticsAsicInstance.EntityData.SegmentPath = "statistics-asic-instance" + types.AddKeyToken(statisticsAsicInstance.Instance, "instance")
    statisticsAsicInstance.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/" + statisticsAsicInstance.EntityData.SegmentPath
    statisticsAsicInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticsAsicInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticsAsicInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticsAsicInstance.EntityData.Children = types.NewOrderedMap()
    statisticsAsicInstance.EntityData.Children.Append("pbc-statistics", types.YChild{"PbcStatistics", &statisticsAsicInstance.PbcStatistics})
    statisticsAsicInstance.EntityData.Children.Append("fmac-statistics", types.YChild{"FmacStatistics", &statisticsAsicInstance.FmacStatistics})
    statisticsAsicInstance.EntityData.Leafs = types.NewOrderedMap()
    statisticsAsicInstance.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", statisticsAsicInstance.Instance})

    statisticsAsicInstance.EntityData.YListKeys = []string {"Instance"}

    return &(statisticsAsicInstance.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics
// Packet Byte Counter for a Asic
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PBC stats bag.
    PbcStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats
}

func (pbcStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics) GetEntityData() *types.CommonEntityData {
    pbcStatistics.EntityData.YFilter = pbcStatistics.YFilter
    pbcStatistics.EntityData.YangName = "pbc-statistics"
    pbcStatistics.EntityData.BundleName = "cisco_ios_xr"
    pbcStatistics.EntityData.ParentYangName = "statistics-asic-instance"
    pbcStatistics.EntityData.SegmentPath = "pbc-statistics"
    pbcStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/" + pbcStatistics.EntityData.SegmentPath
    pbcStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbcStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbcStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbcStatistics.EntityData.Children = types.NewOrderedMap()
    pbcStatistics.EntityData.Children.Append("pbc-stats", types.YChild{"PbcStats", &pbcStatistics.PbcStats})
    pbcStatistics.EntityData.Leafs = types.NewOrderedMap()

    pbcStatistics.EntityData.YListKeys = []string {}

    return &(pbcStatistics.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats
// PBC stats bag
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats struct {
    EntityData types.CommonEntityData
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

    // aggr stats.
    AggrStats Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_OvfStatus
}

func (pbcStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats) GetEntityData() *types.CommonEntityData {
    pbcStats.EntityData.YFilter = pbcStats.YFilter
    pbcStats.EntityData.YangName = "pbc-stats"
    pbcStats.EntityData.BundleName = "cisco_ios_xr"
    pbcStats.EntityData.ParentYangName = "pbc-statistics"
    pbcStats.EntityData.SegmentPath = "pbc-stats"
    pbcStats.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/pbc-statistics/" + pbcStats.EntityData.SegmentPath
    pbcStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pbcStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pbcStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pbcStats.EntityData.Children = types.NewOrderedMap()
    pbcStats.EntityData.Children.Append("aggr-stats", types.YChild{"AggrStats", &pbcStats.AggrStats})
    pbcStats.EntityData.Children.Append("ovf-status", types.YChild{"OvfStatus", &pbcStats.OvfStatus})
    pbcStats.EntityData.Leafs = types.NewOrderedMap()
    pbcStats.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", pbcStats.Valid})
    pbcStats.EntityData.Leafs.Append("rack-no", types.YLeaf{"RackNo", pbcStats.RackNo})
    pbcStats.EntityData.Leafs.Append("slot-no", types.YLeaf{"SlotNo", pbcStats.SlotNo})
    pbcStats.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", pbcStats.AsicInstance})
    pbcStats.EntityData.Leafs.Append("chip-ver", types.YLeaf{"ChipVer", pbcStats.ChipVer})

    pbcStats.EntityData.YListKeys = []string {}

    return &(pbcStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats
// aggr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RxInternalError. The type is interface{} with range:
    // 0..18446744073709551615.
    RxInternalError interface{}

    // RxInternalDrop. The type is interface{} with range:
    // 0..18446744073709551615.
    RxInternalDrop interface{}

    // TxInternalError. The type is interface{} with range:
    // 0..18446744073709551615.
    TxInternalError interface{}

    // TxInternalDrop. The type is interface{} with range:
    // 0..18446744073709551615.
    TxInternalDrop interface{}

    // CMIC Cmc0PktCountTxPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    CmicCmc0PktCountTxPkt interface{}

    // CMIC Cmc0PktCountRxPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    CmicCmc0PktCountRxPkt interface{}

    // NBI StatRxBurstsErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiStatRxBurstsErrCnt interface{}

    // NBI Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiEcc1bErrCnt interface{}

    // NBI Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiEcc2bErrCnt interface{}

    // NBI ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiParityErrCnt interface{}

    // NBI RxIlknCrc32ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlknCrc32ErrCnt interface{}

    // NBI RxIlkn0Crc24ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0Crc24ErrCnt interface{}

    // NBI RxIlkn0BurstErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0BurstErrCnt interface{}

    // NBI RxIlkn0MissSopErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0MissSopErrCnt interface{}

    // NBI RxIlkn0MissEopErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0MissEopErrCnt interface{}

    // NBI RxIlkn0MisalignedCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0MisalignedCnt interface{}

    // NBI RxIlkn1Crc24ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1Crc24ErrCnt interface{}

    // NBI RxIlkn1BurstErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1BurstErrCnt interface{}

    // NBI RxIlkn1MissSopErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1MissSopErrCnt interface{}

    // NBI RxIlkn1MissEopErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1MissEopErrCnt interface{}

    // NBI RxIlkn1MisalignedCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1MisalignedCnt interface{}

    // NBI TxIlkn1FlushedBurstsCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiTxIlkn1FlushedBurstsCnt interface{}

    // NBI RxIlkn0RetransCRC24ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0RetransCrc24ErrCnt interface{}

    // NBI RxIlkn0RetransRetryErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0RetransRetryErrCnt interface{}

    // NBI RxIlkn0RetransWdogErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0RetransWdogErrCnt interface{}

    // NBI RxIlkn0RetransWrapAfterDiscErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0RetransWrapAfterDiscErrCnt interface{}

    // NBI RxIlkn0RetransWrapB4DiscErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0RetransWrapB4DiscErrCnt interface{}

    // NBI RxIlkn0RetransReachedTimeoutErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn0RetransReachedTimeoutErrCnt interface{}

    // NBI RxIlkn1RetransCRC24ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1RetransCrc24ErrCnt interface{}

    // NBI RxIlkn1RetransRetryErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1RetransRetryErrCnt interface{}

    // NBI RxIlkn1RetransWdogErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1RetransWdogErrCnt interface{}

    // NBI RxIlkn1RetransWrapAfterDiscErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1RetransWrapAfterDiscErrCnt interface{}

    // NBI RxIlkn1RetransWrapB4DiscErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1RetransWrapB4DiscErrCnt interface{}

    // NBI RxIlkn1RetransReachedTimeoutErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxIlkn1RetransReachedTimeoutErrCnt interface{}

    // NBI StatRxFrameErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiStatRxFrameErrCnt interface{}

    // NBI StatTxFrameErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiStatTxFrameErrCnt interface{}

    // NBI RxElkErrBurstsCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxElkErrBurstsCnt interface{}

    // NBI RxNumThrownEops. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxNumThrownEops interface{}

    // NBI RxNumRunts. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiRxNumRunts interface{}

    // NBI BistTxCrcErrBurstsCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistTxCrcErrBurstsCnt interface{}

    // NBI BistRxErrLengthBurstsCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistRxErrLengthBurstsCnt interface{}

    // NBI BistRxErrBurstIndexCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistRxErrBurstIndexCnt interface{}

    // NBI BistRxErrBctCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistRxErrBctCnt interface{}

    // NBI BistRxErrDataCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistRxErrDataCnt interface{}

    // NBI BistRxErrInCrcErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistRxErrInCrcErrCnt interface{}

    // NBI BistRxErrSobCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiBistRxErrSobCnt interface{}

    // NBI StatTxBurstsCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiStatTxBurstsCnt interface{}

    // NBI StatTxTotalLengCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    NbiStatTxTotalLengCnt interface{}

    // RXAUI TotalTxPktCount. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTotalTxPktCount interface{}

    // RXAUI TotalRxPktCount. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTotalRxPktCount interface{}

    // RXAUI RxPktCountBcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiRxPktCountBcastPkt interface{}

    // RXAUI TxPktCountBcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTxPktCountBcastPkt interface{}

    // RXAUI RxPktCountMcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiRxPktCountMcastPkt interface{}

    // RXAUI TxPktCountMcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTxPktCountMcastPkt interface{}

    // RXAUI RxPktCountUcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiRxPktCountUcastPkt interface{}

    // RXAUI TxPktCountUcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTxPktCountUcastPkt interface{}

    // RXAUI RxErrDropPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiRxErrDropPktCnt interface{}

    // RXAUI TxErrDropPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTxErrDropPktCnt interface{}

    // RXAUI ByteCountTxPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiByteCountTxPkt interface{}

    // RXAUI ByteCountRxPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiByteCountRxPkt interface{}

    // RXAUI RxDscrdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiRxDscrdPktCnt interface{}

    // RXAUI TxDscrdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    RxauiTxDscrdPktCnt interface{}

    // IRE NifPacketCounter. The type is interface{} with range:
    // 0..18446744073709551615.
    IreNifPacketCounter interface{}

    // IL TotalRxPktCount. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTotalRxPktCount interface{}

    // IL TotalTxPktCount. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTotalTxPktCount interface{}

    // IL RxErrDropPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlRxErrDropPktCnt interface{}

    // IL TxErrDropPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTxErrDropPktCnt interface{}

    // IL ByteCountTxPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlByteCountTxPkt interface{}

    // IL ByteCountRxPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlByteCountRxPkt interface{}

    // IL RxDscrdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlRxDscrdPktCnt interface{}

    // IL TxDscrdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTxDscrdPktCnt interface{}

    // IL RxPktCountBcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlRxPktCountBcastPkt interface{}

    // IL TxPktCountBcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTxPktCountBcastPkt interface{}

    // IL RxPktCountMcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlRxPktCountMcastPkt interface{}

    // IL TxPktCountMcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTxPktCountMcastPkt interface{}

    // IL RxPktCountUcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlRxPktCountUcastPkt interface{}

    // IL TxPktCountUcastPkt. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTxPktCountUcastPkt interface{}

    // IQM EnqPktCnt. The type is interface{} with range: 0..18446744073709551615.
    IqmEnqPktCnt interface{}

    // IQM EnqByteCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmEnqByteCnt interface{}

    // IQM DeqPktCnt. The type is interface{} with range: 0..18446744073709551615.
    IqmDeqPktCnt interface{}

    // IQM DeqByteCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmDeqByteCnt interface{}

    // IQM TotDscrdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmTotDscrdPktCnt interface{}

    // IQM TotDscrdByteCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmTotDscrdByteCnt interface{}

    // IQM Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmEcc1bErrCnt interface{}

    // IQM Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmEcc2bErrCnt interface{}

    // IQM ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmParityErrCnt interface{}

    // IQM DeqDeletePktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmDeqDeletePktCnt interface{}

    // IQM EcnDscrdMskPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmEcnDscrdMskPktCnt interface{}

    // IQM QTotDscrdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmQTotDscrdPktCnt interface{}

    // IQM QDeqDeletePktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmQDeqDeletePktCnt interface{}

    // IQM RjctDbPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctDbPktCnt interface{}

    // IQM RjctBdbPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctBdbPktCnt interface{}

    // IQM RjctBdbProtctPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctBdbProtctPktCnt interface{}

    // IQM RjctOcBdPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctOcBdPktCnt interface{}

    // IQM RjctSnErrPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctSnErrPktCnt interface{}

    // IQM RjctMcErrPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctMcErrPktCnt interface{}

    // IQM RjctRsrcErrPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctRsrcErrPktCnt interface{}

    // IQM RjctQnvalidErrPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctQnvalidErrPktCnt interface{}

    // IQM RjctCnmPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctCnmPktCnt interface{}

    // IQM RjctDynSpacePktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IqmRjctDynSpacePktCnt interface{}

    // IPT FdtPktCnt. The type is interface{} with range: 0..18446744073709551615.
    IptFdtPktCnt interface{}

    // IPT Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptEcc1bErrCnt interface{}

    // IPT Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptEcc2bErrCnt interface{}

    // IPT ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptParityErrCnt interface{}

    // IPT CrcErrCnt. The type is interface{} with range: 0..18446744073709551615.
    IptCrcErrCnt interface{}

    // IPT CrcErrDelBuffCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptCrcErrDelBuffCnt interface{}

    // IPT CpuDelBuffCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptCpuDelBuffCnt interface{}

    // IPT CpuRelBuffCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptCpuRelBuffCnt interface{}

    // IPT CrcErrBuffFifoFullCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IptCrcErrBuffFifoFullCnt interface{}

    // FDT DataCellCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtDataCellCnt interface{}

    // FDT DataByteCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtDataByteCnt interface{}

    // FDT CrcDroppedPckCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtCrcDroppedPckCnt interface{}

    // FDT invalid destq drop cell cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtInvalidDestqDropCellCnt interface{}

    // FDT IndirectCommandCount. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtIndirectCommandCount interface{}

    // FDT Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtEcc1bErrCnt interface{}

    // FDT Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtEcc2bErrCnt interface{}

    // FDT ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtParityErrCnt interface{}

    // FDT CrcDroppedCellCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdtCrcDroppedCellCnt interface{}

    // FCR ControlCellCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrControlCellCnt interface{}

    // FCR CellDropCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrCellDropCnt interface{}

    // FCR CreditCellDropCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrCreditCellDropCnt interface{}

    // FCR FSCellDropCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrFsCellDropCnt interface{}

    // FCR RTCellDropCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrRtCellDropCnt interface{}

    // FCR Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrEcc1bErrCnt interface{}

    // FCR Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FcrEcc2bErrCnt interface{}

    // FDR DataCellCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrDataCellCnt interface{}

    // FDR DataByteCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrDataByteCnt interface{}

    // FDR CrcDroppedPckCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrCrcDroppedPckCnt interface{}

    // FDR PPktCnt. The type is interface{} with range: 0..18446744073709551615.
    FdrPPktCnt interface{}

    // FDR PrmErrorFilterCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrPrmErrorFilterCnt interface{}

    // FDR SecErrorFilterCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrSecErrorFilterCnt interface{}

    // FDR PrmEcc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrPrmEcc1bErrCnt interface{}

    // FDR PrmEcc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrPrmEcc2bErrCnt interface{}

    // FDR SecEcc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrSecEcc1bErrCnt interface{}

    // FDR SecEcc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FdrSecEcc2bErrCnt interface{}

    // EGQ Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEcc1bErrCnt interface{}

    // EGQ Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEcc2bErrCnt interface{}

    // EGQ ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqParityErrCnt interface{}

    // EGQ DbfEcc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqDbfEcc1bErrCnt interface{}

    // EGQ DbfEcc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqDbfEcc2bErrCnt interface{}

    // EGQ EmptyMcidCounter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEmptyMcidCounter interface{}

    // EGQ RqpDiscardPacketCounter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqRqpDiscardPacketCounter interface{}

    // EGQ EhpDiscardPacketCounter. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqEhpDiscardPacketCounter interface{}

    // EGQ IptPktCnt. The type is interface{} with range: 0..18446744073709551615.
    EgqIptPktCnt interface{}

    // EPNI EpePktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpePktCnt interface{}

    // EPNI EpeByteCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpeByteCnt interface{}

    // EPNI EpeDiscardPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEpeDiscardPktCnt interface{}

    // EPNI Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEcc1bErrCnt interface{}

    // EPNI Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniEcc2bErrCnt interface{}

    // EPNI ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EpniParityErrCnt interface{}

    // EGQ PqpUcastPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcastPktCnt interface{}

    // EGQ PqpUcastHPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcastHPktCnt interface{}

    // EGQ PqpUcastLPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcastLPktCnt interface{}

    // EGQ PqpUcastBytesCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcastBytesCnt interface{}

    // EGQ PqpUcastDiscardPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpUcastDiscardPktCnt interface{}

    // EGQ PqpMcastPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcastPktCnt interface{}

    // EGQ PqpMcastHPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcastHPktCnt interface{}

    // EGQ PqpMcastLPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcastLPktCnt interface{}

    // EGQ PqpMcastBytesCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcastBytesCnt interface{}

    // EGQ PqpMcastDiscardPktCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    EgqPqpMcastDiscardPktCnt interface{}

    // FCT ControlCellCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FctControlCellCnt interface{}

    // FCT UnrchCrdtCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    FctUnrchCrdtCnt interface{}

    // IDR ReassemblyErrors. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrReassemblyErrors interface{}

    // IDR MmuEcc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrMmuEcc1bErrCnt interface{}

    // IDR MmuEcc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrMmuEcc2bErrCnt interface{}

    // IDR DiscardedPackets0Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedPackets0Cnt interface{}

    // IDR DiscardedPackets1Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedPackets1Cnt interface{}

    // IDR DiscardedPackets2Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedPackets2Cnt interface{}

    // IDR DiscardedPackets3Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedPackets3Cnt interface{}

    // IDR DiscardedOctets0Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedOctets0Cnt interface{}

    // IDR DiscardedOctets1Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedOctets1Cnt interface{}

    // IDR DiscardedOctets2Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedOctets2Cnt interface{}

    // IDR DiscardedOctets3Cnt. The type is interface{} with range:
    // 0..18446744073709551615.
    IdrDiscardedOctets3Cnt interface{}

    // MMU Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    MmuEcc1bErrCnt interface{}

    // MMU Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    MmuEcc2bErrCnt interface{}

    // OAMP ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    OampParityErrCnt interface{}

    // OAMP Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    OampEcc1bErrCnt interface{}

    // OAMP Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    OampEcc2bErrCnt interface{}

    // CRPS ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CrpsParityErrCnt interface{}

    // FMAC0 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Kpcs0TstRxErrCnt interface{}

    // FMAC1 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Kpcs0TstRxErrCnt interface{}

    // FMAC2 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Kpcs0TstRxErrCnt interface{}

    // FMAC3 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Kpcs0TstRxErrCnt interface{}

    // FMAC4 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Kpcs0TstRxErrCnt interface{}

    // FMAC5 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Kpcs0TstRxErrCnt interface{}

    // FMAC6 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Kpcs0TstRxErrCnt interface{}

    // FMAC7 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Kpcs0TstRxErrCnt interface{}

    // FMAC8 Kpcs0TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Kpcs0TstRxErrCnt interface{}

    // FMAC0 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Kpcs1TstRxErrCnt interface{}

    // FMAC1 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Kpcs1TstRxErrCnt interface{}

    // FMAC2 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Kpcs1TstRxErrCnt interface{}

    // FMAC3 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Kpcs1TstRxErrCnt interface{}

    // FMAC4 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Kpcs1TstRxErrCnt interface{}

    // FMAC5 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Kpcs1TstRxErrCnt interface{}

    // FMAC6 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Kpcs1TstRxErrCnt interface{}

    // FMAC7 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Kpcs1TstRxErrCnt interface{}

    // FMAC8 Kpcs1TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Kpcs1TstRxErrCnt interface{}

    // FMAC0 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Kpcs2TstRxErrCnt interface{}

    // FMAC1 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Kpcs2TstRxErrCnt interface{}

    // FMAC2 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Kpcs2TstRxErrCnt interface{}

    // FMAC3 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Kpcs2TstRxErrCnt interface{}

    // FMAC4 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Kpcs2TstRxErrCnt interface{}

    // FMAC5 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Kpcs2TstRxErrCnt interface{}

    // FMAC6 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Kpcs2TstRxErrCnt interface{}

    // FMAC7 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Kpcs2TstRxErrCnt interface{}

    // FMAC8 Kpcs2TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Kpcs2TstRxErrCnt interface{}

    // FMAC0 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Kpcs3TstRxErrCnt interface{}

    // FMAC1 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Kpcs3TstRxErrCnt interface{}

    // FMAC2 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Kpcs3TstRxErrCnt interface{}

    // FMAC3 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Kpcs3TstRxErrCnt interface{}

    // FMAC4 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Kpcs3TstRxErrCnt interface{}

    // FMAC5 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Kpcs3TstRxErrCnt interface{}

    // FMAC6 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Kpcs3TstRxErrCnt interface{}

    // FMAC7 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Kpcs3TstRxErrCnt interface{}

    // FMAC8 Kpcs3TstRxErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Kpcs3TstRxErrCnt interface{}

    // FMAC0 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Tst0ErrCnt interface{}

    // FMAC1 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Tst0ErrCnt interface{}

    // FMAC2 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Tst0ErrCnt interface{}

    // FMAC3 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Tst0ErrCnt interface{}

    // FMAC4 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Tst0ErrCnt interface{}

    // FMAC5 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Tst0ErrCnt interface{}

    // FMAC6 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Tst0ErrCnt interface{}

    // FMAC7 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Tst0ErrCnt interface{}

    // FMAC8 Tst0ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Tst0ErrCnt interface{}

    // FMAC0 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Tst1ErrCnt interface{}

    // FMAC1 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Tst1ErrCnt interface{}

    // FMAC2 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Tst1ErrCnt interface{}

    // FMAC3 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Tst1ErrCnt interface{}

    // FMAC4 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Tst1ErrCnt interface{}

    // FMAC5 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Tst1ErrCnt interface{}

    // FMAC6 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Tst1ErrCnt interface{}

    // FMAC7 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Tst1ErrCnt interface{}

    // FMAC8 Tst1ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Tst1ErrCnt interface{}

    // FMAC0 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Tst2ErrCnt interface{}

    // FMAC1 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Tst2ErrCnt interface{}

    // FMAC2 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Tst2ErrCnt interface{}

    // FMAC3 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Tst2ErrCnt interface{}

    // FMAC4 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Tst2ErrCnt interface{}

    // FMAC5 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Tst2ErrCnt interface{}

    // FMAC6 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Tst2ErrCnt interface{}

    // FMAC7 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Tst2ErrCnt interface{}

    // FMAC8 Tst2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Tst2ErrCnt interface{}

    // FMAC0 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Tst3ErrCnt interface{}

    // FMAC1 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Tst3ErrCnt interface{}

    // FMAC2 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Tst3ErrCnt interface{}

    // FMAC3 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Tst3ErrCnt interface{}

    // FMAC4 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Tst3ErrCnt interface{}

    // FMAC5 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Tst3ErrCnt interface{}

    // FMAC6 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Tst3ErrCnt interface{}

    // FMAC7 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Tst3ErrCnt interface{}

    // FMAC8 Tst3ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Tst3ErrCnt interface{}

    // FMAC0 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Ecc1bErrCnt interface{}

    // FMAC1 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Ecc1bErrCnt interface{}

    // FMAC2 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Ecc1bErrCnt interface{}

    // FMAC3 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Ecc1bErrCnt interface{}

    // FMAC4 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Ecc1bErrCnt interface{}

    // FMAC5 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Ecc1bErrCnt interface{}

    // FMAC6 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Ecc1bErrCnt interface{}

    // FMAC7 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Ecc1bErrCnt interface{}

    // FMAC8 Ecc 1bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Ecc1bErrCnt interface{}

    // FMAC0 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac0Ecc2bErrCnt interface{}

    // FMAC1 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac1Ecc2bErrCnt interface{}

    // FMAC2 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac2Ecc2bErrCnt interface{}

    // FMAC3 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac3Ecc2bErrCnt interface{}

    // FMAC4 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac4Ecc2bErrCnt interface{}

    // FMAC5 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac5Ecc2bErrCnt interface{}

    // FMAC6 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac6Ecc2bErrCnt interface{}

    // FMAC7 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac7Ecc2bErrCnt interface{}

    // FMAC8 Ecc 2bErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    Fmac8Ecc2bErrCnt interface{}

    // OLP IncomingBadIdentifierCounter. The type is interface{} with range:
    // 0..18446744073709551615.
    OlpIncomingBadIdentifierCounter interface{}

    // OLP IncomingBadReassemblyCounter. The type is interface{} with range:
    // 0..18446744073709551615.
    OlpIncomingBadReassemblyCounter interface{}

    // CFC ParityErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcParityErrCnt interface{}

    // CFC Ilkn0OobRxCrcErrCntr. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcIlkn0OobRxCrcErrCntr interface{}

    // CFC Ilkn1OobRxCrcErrCntr. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcIlkn1OobRxCrcErrCntr interface{}

    // CFC SpiOobRx0FrmErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcSpiOobRx0FrmErrCnt interface{}

    // CFC SpiOobRx0Dip2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcSpiOobRx0Dip2ErrCnt interface{}

    // CFC SpiOobRx1FrmErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcSpiOobRx1FrmErrCnt interface{}

    // CFC SpiOobRx1Dip2ErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CfcSpiOobRx1Dip2ErrCnt interface{}

    // CGM CgmUcPdDroppedCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CgmCgmUcPdDroppedCnt interface{}

    // CGM CgmMcRepPdDroppedCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CgmCgmMcRepPdDroppedCnt interface{}

    // CGM CgmUcDbDroppedByRqpCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CgmCgmUcDbDroppedByRqpCnt interface{}

    // CGM CgmUcDbDroppedByPqpCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CgmCgmUcDbDroppedByPqpCnt interface{}

    // CGM CgmMcRepDbDroppedCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CgmCgmMcRepDbDroppedCnt interface{}

    // CGM CgmMcDbDroppedCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    CgmCgmMcDbDroppedCnt interface{}

    // DRCA FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcaFullErrCnt interface{}

    // DRCA SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcaSingleErrCnt interface{}

    // DRCA CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcaCalibBistFullErrCnt interface{}

    // DRCA LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcaLoopbackFullErrCnt interface{}

    // DRCB FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbFullErrCnt interface{}

    // DRCB SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbSingleErrCnt interface{}

    // DRCB CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbCalibBistFullErrCnt interface{}

    // DRCB LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbLoopbackFullErrCnt interface{}

    // DRCC FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrccFullErrCnt interface{}

    // DRCC SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrccSingleErrCnt interface{}

    // DRCC CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrccCalibBistFullErrCnt interface{}

    // DRCC LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrccLoopbackFullErrCnt interface{}

    // DRCD FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcdFullErrCnt interface{}

    // DRCD SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcdSingleErrCnt interface{}

    // DRCD CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcdCalibBistFullErrCnt interface{}

    // DRCD LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcdLoopbackFullErrCnt interface{}

    // DRCE FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrceFullErrCnt interface{}

    // DRCE SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrceSingleErrCnt interface{}

    // DRCE CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrceCalibBistFullErrCnt interface{}

    // DRCE LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrceLoopbackFullErrCnt interface{}

    // DRCF FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcfFullErrCnt interface{}

    // DRCF SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcfSingleErrCnt interface{}

    // DRCF CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcfCalibBistFullErrCnt interface{}

    // DRCF LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcfLoopbackFullErrCnt interface{}

    // DRCG FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcgFullErrCnt interface{}

    // DRCG SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcgSingleErrCnt interface{}

    // DRCG CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcgCalibBistFullErrCnt interface{}

    // DRCG LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcgLoopbackFullErrCnt interface{}

    // DRCH FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrchFullErrCnt interface{}

    // DRCH SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrchSingleErrCnt interface{}

    // DRCH CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrchCalibBistFullErrCnt interface{}

    // DRCH LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrchLoopbackFullErrCnt interface{}

    // DRCBROADCAST FullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbroadcastFullErrCnt interface{}

    // DRCBROADCAST SingleErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbroadcastSingleErrCnt interface{}

    // DRCBROADCAST CalibBistFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbroadcastCalibBistFullErrCnt interface{}

    // DRCBROADCAST LoopbackFullErrCnt. The type is interface{} with range:
    // 0..18446744073709551615.
    DrcbroadcastLoopbackFullErrCnt interface{}

    // otn mode. The type is interface{} with range: 0..4294967295.
    OtnMode interface{}

    // num ports. The type is interface{} with range: 0..4294967295.
    NumPorts interface{}

    // aggr stats otn. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats_AggrStatsOtn.
    AggrStatsOtn []*Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats_AggrStatsOtn
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats) GetEntityData() *types.CommonEntityData {
    aggrStats.EntityData.YFilter = aggrStats.YFilter
    aggrStats.EntityData.YangName = "aggr-stats"
    aggrStats.EntityData.BundleName = "cisco_ios_xr"
    aggrStats.EntityData.ParentYangName = "pbc-stats"
    aggrStats.EntityData.SegmentPath = "aggr-stats"
    aggrStats.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/pbc-statistics/pbc-stats/" + aggrStats.EntityData.SegmentPath
    aggrStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggrStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggrStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggrStats.EntityData.Children = types.NewOrderedMap()
    aggrStats.EntityData.Children.Append("aggr-stats-otn", types.YChild{"AggrStatsOtn", nil})
    for i := range aggrStats.AggrStatsOtn {
        types.SetYListKey(aggrStats.AggrStatsOtn[i], i)
        aggrStats.EntityData.Children.Append(types.GetSegmentPath(aggrStats.AggrStatsOtn[i]), types.YChild{"AggrStatsOtn", aggrStats.AggrStatsOtn[i]})
    }
    aggrStats.EntityData.Leafs = types.NewOrderedMap()
    aggrStats.EntityData.Leafs.Append("rx-internal-error", types.YLeaf{"RxInternalError", aggrStats.RxInternalError})
    aggrStats.EntityData.Leafs.Append("rx-internal-drop", types.YLeaf{"RxInternalDrop", aggrStats.RxInternalDrop})
    aggrStats.EntityData.Leafs.Append("tx-internal-error", types.YLeaf{"TxInternalError", aggrStats.TxInternalError})
    aggrStats.EntityData.Leafs.Append("tx-internal-drop", types.YLeaf{"TxInternalDrop", aggrStats.TxInternalDrop})
    aggrStats.EntityData.Leafs.Append("cmic-cmc0-pkt-count-tx-pkt", types.YLeaf{"CmicCmc0PktCountTxPkt", aggrStats.CmicCmc0PktCountTxPkt})
    aggrStats.EntityData.Leafs.Append("cmic-cmc0-pkt-count-rx-pkt", types.YLeaf{"CmicCmc0PktCountRxPkt", aggrStats.CmicCmc0PktCountRxPkt})
    aggrStats.EntityData.Leafs.Append("nbi-stat-rx-bursts-err-cnt", types.YLeaf{"NbiStatRxBurstsErrCnt", aggrStats.NbiStatRxBurstsErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-ecc-1b-err-cnt", types.YLeaf{"NbiEcc1bErrCnt", aggrStats.NbiEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-ecc-2b-err-cnt", types.YLeaf{"NbiEcc2bErrCnt", aggrStats.NbiEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-parity-err-cnt", types.YLeaf{"NbiParityErrCnt", aggrStats.NbiParityErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn-crc32-err-cnt", types.YLeaf{"NbiRxIlknCrc32ErrCnt", aggrStats.NbiRxIlknCrc32ErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-crc24-err-cnt", types.YLeaf{"NbiRxIlkn0Crc24ErrCnt", aggrStats.NbiRxIlkn0Crc24ErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-burst-err-cnt", types.YLeaf{"NbiRxIlkn0BurstErrCnt", aggrStats.NbiRxIlkn0BurstErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-miss-sop-err-cnt", types.YLeaf{"NbiRxIlkn0MissSopErrCnt", aggrStats.NbiRxIlkn0MissSopErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-miss-eop-err-cnt", types.YLeaf{"NbiRxIlkn0MissEopErrCnt", aggrStats.NbiRxIlkn0MissEopErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-misaligned-cnt", types.YLeaf{"NbiRxIlkn0MisalignedCnt", aggrStats.NbiRxIlkn0MisalignedCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-crc24-err-cnt", types.YLeaf{"NbiRxIlkn1Crc24ErrCnt", aggrStats.NbiRxIlkn1Crc24ErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-burst-err-cnt", types.YLeaf{"NbiRxIlkn1BurstErrCnt", aggrStats.NbiRxIlkn1BurstErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-miss-sop-err-cnt", types.YLeaf{"NbiRxIlkn1MissSopErrCnt", aggrStats.NbiRxIlkn1MissSopErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-miss-eop-err-cnt", types.YLeaf{"NbiRxIlkn1MissEopErrCnt", aggrStats.NbiRxIlkn1MissEopErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-misaligned-cnt", types.YLeaf{"NbiRxIlkn1MisalignedCnt", aggrStats.NbiRxIlkn1MisalignedCnt})
    aggrStats.EntityData.Leafs.Append("nbi-tx-ilkn1-flushed-bursts-cnt", types.YLeaf{"NbiTxIlkn1FlushedBurstsCnt", aggrStats.NbiTxIlkn1FlushedBurstsCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-crc24-err-cnt", types.YLeaf{"NbiRxIlkn0RetransCrc24ErrCnt", aggrStats.NbiRxIlkn0RetransCrc24ErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-retry-err-cnt", types.YLeaf{"NbiRxIlkn0RetransRetryErrCnt", aggrStats.NbiRxIlkn0RetransRetryErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-wdog-err-cnt", types.YLeaf{"NbiRxIlkn0RetransWdogErrCnt", aggrStats.NbiRxIlkn0RetransWdogErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-wrap-after-disc-err-cnt", types.YLeaf{"NbiRxIlkn0RetransWrapAfterDiscErrCnt", aggrStats.NbiRxIlkn0RetransWrapAfterDiscErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-wrap-b4-disc-err-cnt", types.YLeaf{"NbiRxIlkn0RetransWrapB4DiscErrCnt", aggrStats.NbiRxIlkn0RetransWrapB4DiscErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-reached-timeout-err-cnt", types.YLeaf{"NbiRxIlkn0RetransReachedTimeoutErrCnt", aggrStats.NbiRxIlkn0RetransReachedTimeoutErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-crc24-err-cnt", types.YLeaf{"NbiRxIlkn1RetransCrc24ErrCnt", aggrStats.NbiRxIlkn1RetransCrc24ErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-retry-err-cnt", types.YLeaf{"NbiRxIlkn1RetransRetryErrCnt", aggrStats.NbiRxIlkn1RetransRetryErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-wdog-err-cnt", types.YLeaf{"NbiRxIlkn1RetransWdogErrCnt", aggrStats.NbiRxIlkn1RetransWdogErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-wrap-after-disc-err-cnt", types.YLeaf{"NbiRxIlkn1RetransWrapAfterDiscErrCnt", aggrStats.NbiRxIlkn1RetransWrapAfterDiscErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-wrap-b4-disc-err-cnt", types.YLeaf{"NbiRxIlkn1RetransWrapB4DiscErrCnt", aggrStats.NbiRxIlkn1RetransWrapB4DiscErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-reached-timeout-err-cnt", types.YLeaf{"NbiRxIlkn1RetransReachedTimeoutErrCnt", aggrStats.NbiRxIlkn1RetransReachedTimeoutErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-stat-rx-frame-err-cnt", types.YLeaf{"NbiStatRxFrameErrCnt", aggrStats.NbiStatRxFrameErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-stat-tx-frame-err-cnt", types.YLeaf{"NbiStatTxFrameErrCnt", aggrStats.NbiStatTxFrameErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-elk-err-bursts-cnt", types.YLeaf{"NbiRxElkErrBurstsCnt", aggrStats.NbiRxElkErrBurstsCnt})
    aggrStats.EntityData.Leafs.Append("nbi-rx-num-thrown-eops", types.YLeaf{"NbiRxNumThrownEops", aggrStats.NbiRxNumThrownEops})
    aggrStats.EntityData.Leafs.Append("nbi-rx-num-runts", types.YLeaf{"NbiRxNumRunts", aggrStats.NbiRxNumRunts})
    aggrStats.EntityData.Leafs.Append("nbi-bist-tx-crc-err-bursts-cnt", types.YLeaf{"NbiBistTxCrcErrBurstsCnt", aggrStats.NbiBistTxCrcErrBurstsCnt})
    aggrStats.EntityData.Leafs.Append("nbi-bist-rx-err-length-bursts-cnt", types.YLeaf{"NbiBistRxErrLengthBurstsCnt", aggrStats.NbiBistRxErrLengthBurstsCnt})
    aggrStats.EntityData.Leafs.Append("nbi-bist-rx-err-burst-index-cnt", types.YLeaf{"NbiBistRxErrBurstIndexCnt", aggrStats.NbiBistRxErrBurstIndexCnt})
    aggrStats.EntityData.Leafs.Append("nbi-bist-rx-err-bct-cnt", types.YLeaf{"NbiBistRxErrBctCnt", aggrStats.NbiBistRxErrBctCnt})
    aggrStats.EntityData.Leafs.Append("nbi-bist-rx-err-data-cnt", types.YLeaf{"NbiBistRxErrDataCnt", aggrStats.NbiBistRxErrDataCnt})
    aggrStats.EntityData.Leafs.Append("nbi-bist-rx-err-in-crc-err-cnt", types.YLeaf{"NbiBistRxErrInCrcErrCnt", aggrStats.NbiBistRxErrInCrcErrCnt})
    aggrStats.EntityData.Leafs.Append("nbi-bist-rx-err-sob-cnt", types.YLeaf{"NbiBistRxErrSobCnt", aggrStats.NbiBistRxErrSobCnt})
    aggrStats.EntityData.Leafs.Append("nbi-stat-tx-bursts-cnt", types.YLeaf{"NbiStatTxBurstsCnt", aggrStats.NbiStatTxBurstsCnt})
    aggrStats.EntityData.Leafs.Append("nbi-stat-tx-total-leng-cnt", types.YLeaf{"NbiStatTxTotalLengCnt", aggrStats.NbiStatTxTotalLengCnt})
    aggrStats.EntityData.Leafs.Append("rxaui-total-tx-pkt-count", types.YLeaf{"RxauiTotalTxPktCount", aggrStats.RxauiTotalTxPktCount})
    aggrStats.EntityData.Leafs.Append("rxaui-total-rx-pkt-count", types.YLeaf{"RxauiTotalRxPktCount", aggrStats.RxauiTotalRxPktCount})
    aggrStats.EntityData.Leafs.Append("rxaui-rx-pkt-count-bcast-pkt", types.YLeaf{"RxauiRxPktCountBcastPkt", aggrStats.RxauiRxPktCountBcastPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-tx-pkt-count-bcast-pkt", types.YLeaf{"RxauiTxPktCountBcastPkt", aggrStats.RxauiTxPktCountBcastPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-rx-pkt-count-mcast-pkt", types.YLeaf{"RxauiRxPktCountMcastPkt", aggrStats.RxauiRxPktCountMcastPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-tx-pkt-count-mcast-pkt", types.YLeaf{"RxauiTxPktCountMcastPkt", aggrStats.RxauiTxPktCountMcastPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-rx-pkt-count-ucast-pkt", types.YLeaf{"RxauiRxPktCountUcastPkt", aggrStats.RxauiRxPktCountUcastPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-tx-pkt-count-ucast-pkt", types.YLeaf{"RxauiTxPktCountUcastPkt", aggrStats.RxauiTxPktCountUcastPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-rx-err-drop-pkt-cnt", types.YLeaf{"RxauiRxErrDropPktCnt", aggrStats.RxauiRxErrDropPktCnt})
    aggrStats.EntityData.Leafs.Append("rxaui-tx-err-drop-pkt-cnt", types.YLeaf{"RxauiTxErrDropPktCnt", aggrStats.RxauiTxErrDropPktCnt})
    aggrStats.EntityData.Leafs.Append("rxaui-byte-count-tx-pkt", types.YLeaf{"RxauiByteCountTxPkt", aggrStats.RxauiByteCountTxPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-byte-count-rx-pkt", types.YLeaf{"RxauiByteCountRxPkt", aggrStats.RxauiByteCountRxPkt})
    aggrStats.EntityData.Leafs.Append("rxaui-rx-dscrd-pkt-cnt", types.YLeaf{"RxauiRxDscrdPktCnt", aggrStats.RxauiRxDscrdPktCnt})
    aggrStats.EntityData.Leafs.Append("rxaui-tx-dscrd-pkt-cnt", types.YLeaf{"RxauiTxDscrdPktCnt", aggrStats.RxauiTxDscrdPktCnt})
    aggrStats.EntityData.Leafs.Append("ire-nif-packet-counter", types.YLeaf{"IreNifPacketCounter", aggrStats.IreNifPacketCounter})
    aggrStats.EntityData.Leafs.Append("il-total-rx-pkt-count", types.YLeaf{"IlTotalRxPktCount", aggrStats.IlTotalRxPktCount})
    aggrStats.EntityData.Leafs.Append("il-total-tx-pkt-count", types.YLeaf{"IlTotalTxPktCount", aggrStats.IlTotalTxPktCount})
    aggrStats.EntityData.Leafs.Append("il-rx-err-drop-pkt-cnt", types.YLeaf{"IlRxErrDropPktCnt", aggrStats.IlRxErrDropPktCnt})
    aggrStats.EntityData.Leafs.Append("il-tx-err-drop-pkt-cnt", types.YLeaf{"IlTxErrDropPktCnt", aggrStats.IlTxErrDropPktCnt})
    aggrStats.EntityData.Leafs.Append("il-byte-count-tx-pkt", types.YLeaf{"IlByteCountTxPkt", aggrStats.IlByteCountTxPkt})
    aggrStats.EntityData.Leafs.Append("il-byte-count-rx-pkt", types.YLeaf{"IlByteCountRxPkt", aggrStats.IlByteCountRxPkt})
    aggrStats.EntityData.Leafs.Append("il-rx-dscrd-pkt-cnt", types.YLeaf{"IlRxDscrdPktCnt", aggrStats.IlRxDscrdPktCnt})
    aggrStats.EntityData.Leafs.Append("il-tx-dscrd-pkt-cnt", types.YLeaf{"IlTxDscrdPktCnt", aggrStats.IlTxDscrdPktCnt})
    aggrStats.EntityData.Leafs.Append("il-rx-pkt-count-bcast-pkt", types.YLeaf{"IlRxPktCountBcastPkt", aggrStats.IlRxPktCountBcastPkt})
    aggrStats.EntityData.Leafs.Append("il-tx-pkt-count-bcast-pkt", types.YLeaf{"IlTxPktCountBcastPkt", aggrStats.IlTxPktCountBcastPkt})
    aggrStats.EntityData.Leafs.Append("il-rx-pkt-count-mcast-pkt", types.YLeaf{"IlRxPktCountMcastPkt", aggrStats.IlRxPktCountMcastPkt})
    aggrStats.EntityData.Leafs.Append("il-tx-pkt-count-mcast-pkt", types.YLeaf{"IlTxPktCountMcastPkt", aggrStats.IlTxPktCountMcastPkt})
    aggrStats.EntityData.Leafs.Append("il-rx-pkt-count-ucast-pkt", types.YLeaf{"IlRxPktCountUcastPkt", aggrStats.IlRxPktCountUcastPkt})
    aggrStats.EntityData.Leafs.Append("il-tx-pkt-count-ucast-pkt", types.YLeaf{"IlTxPktCountUcastPkt", aggrStats.IlTxPktCountUcastPkt})
    aggrStats.EntityData.Leafs.Append("iqm-enq-pkt-cnt", types.YLeaf{"IqmEnqPktCnt", aggrStats.IqmEnqPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-enq-byte-cnt", types.YLeaf{"IqmEnqByteCnt", aggrStats.IqmEnqByteCnt})
    aggrStats.EntityData.Leafs.Append("iqm-deq-pkt-cnt", types.YLeaf{"IqmDeqPktCnt", aggrStats.IqmDeqPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-deq-byte-cnt", types.YLeaf{"IqmDeqByteCnt", aggrStats.IqmDeqByteCnt})
    aggrStats.EntityData.Leafs.Append("iqm-tot-dscrd-pkt-cnt", types.YLeaf{"IqmTotDscrdPktCnt", aggrStats.IqmTotDscrdPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-tot-dscrd-byte-cnt", types.YLeaf{"IqmTotDscrdByteCnt", aggrStats.IqmTotDscrdByteCnt})
    aggrStats.EntityData.Leafs.Append("iqm-ecc-1b-err-cnt", types.YLeaf{"IqmEcc1bErrCnt", aggrStats.IqmEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("iqm-ecc-2b-err-cnt", types.YLeaf{"IqmEcc2bErrCnt", aggrStats.IqmEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("iqm-parity-err-cnt", types.YLeaf{"IqmParityErrCnt", aggrStats.IqmParityErrCnt})
    aggrStats.EntityData.Leafs.Append("iqm-deq-delete-pkt-cnt", types.YLeaf{"IqmDeqDeletePktCnt", aggrStats.IqmDeqDeletePktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-ecn-dscrd-msk-pkt-cnt", types.YLeaf{"IqmEcnDscrdMskPktCnt", aggrStats.IqmEcnDscrdMskPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-q-tot-dscrd-pkt-cnt", types.YLeaf{"IqmQTotDscrdPktCnt", aggrStats.IqmQTotDscrdPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-q-deq-delete-pkt-cnt", types.YLeaf{"IqmQDeqDeletePktCnt", aggrStats.IqmQDeqDeletePktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-db-pkt-cnt", types.YLeaf{"IqmRjctDbPktCnt", aggrStats.IqmRjctDbPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-bdb-pkt-cnt", types.YLeaf{"IqmRjctBdbPktCnt", aggrStats.IqmRjctBdbPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-bdb-protct-pkt-cnt", types.YLeaf{"IqmRjctBdbProtctPktCnt", aggrStats.IqmRjctBdbProtctPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-oc-bd-pkt-cnt", types.YLeaf{"IqmRjctOcBdPktCnt", aggrStats.IqmRjctOcBdPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-sn-err-pkt-cnt", types.YLeaf{"IqmRjctSnErrPktCnt", aggrStats.IqmRjctSnErrPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-mc-err-pkt-cnt", types.YLeaf{"IqmRjctMcErrPktCnt", aggrStats.IqmRjctMcErrPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-rsrc-err-pkt-cnt", types.YLeaf{"IqmRjctRsrcErrPktCnt", aggrStats.IqmRjctRsrcErrPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-qnvalid-err-pkt-cnt", types.YLeaf{"IqmRjctQnvalidErrPktCnt", aggrStats.IqmRjctQnvalidErrPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-cnm-pkt-cnt", types.YLeaf{"IqmRjctCnmPktCnt", aggrStats.IqmRjctCnmPktCnt})
    aggrStats.EntityData.Leafs.Append("iqm-rjct-dyn-space-pkt-cnt", types.YLeaf{"IqmRjctDynSpacePktCnt", aggrStats.IqmRjctDynSpacePktCnt})
    aggrStats.EntityData.Leafs.Append("ipt-fdt-pkt-cnt", types.YLeaf{"IptFdtPktCnt", aggrStats.IptFdtPktCnt})
    aggrStats.EntityData.Leafs.Append("ipt-ecc-1b-err-cnt", types.YLeaf{"IptEcc1bErrCnt", aggrStats.IptEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("ipt-ecc-2b-err-cnt", types.YLeaf{"IptEcc2bErrCnt", aggrStats.IptEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("ipt-parity-err-cnt", types.YLeaf{"IptParityErrCnt", aggrStats.IptParityErrCnt})
    aggrStats.EntityData.Leafs.Append("ipt-crc-err-cnt", types.YLeaf{"IptCrcErrCnt", aggrStats.IptCrcErrCnt})
    aggrStats.EntityData.Leafs.Append("ipt-crc-err-del-buff-cnt", types.YLeaf{"IptCrcErrDelBuffCnt", aggrStats.IptCrcErrDelBuffCnt})
    aggrStats.EntityData.Leafs.Append("ipt-cpu-del-buff-cnt", types.YLeaf{"IptCpuDelBuffCnt", aggrStats.IptCpuDelBuffCnt})
    aggrStats.EntityData.Leafs.Append("ipt-cpu-rel-buff-cnt", types.YLeaf{"IptCpuRelBuffCnt", aggrStats.IptCpuRelBuffCnt})
    aggrStats.EntityData.Leafs.Append("ipt-crc-err-buff-fifo-full-cnt", types.YLeaf{"IptCrcErrBuffFifoFullCnt", aggrStats.IptCrcErrBuffFifoFullCnt})
    aggrStats.EntityData.Leafs.Append("fdt-data-cell-cnt", types.YLeaf{"FdtDataCellCnt", aggrStats.FdtDataCellCnt})
    aggrStats.EntityData.Leafs.Append("fdt-data-byte-cnt", types.YLeaf{"FdtDataByteCnt", aggrStats.FdtDataByteCnt})
    aggrStats.EntityData.Leafs.Append("fdt-crc-dropped-pck-cnt", types.YLeaf{"FdtCrcDroppedPckCnt", aggrStats.FdtCrcDroppedPckCnt})
    aggrStats.EntityData.Leafs.Append("fdt-invalid-destq-drop-cell-cnt", types.YLeaf{"FdtInvalidDestqDropCellCnt", aggrStats.FdtInvalidDestqDropCellCnt})
    aggrStats.EntityData.Leafs.Append("fdt-indirect-command-count", types.YLeaf{"FdtIndirectCommandCount", aggrStats.FdtIndirectCommandCount})
    aggrStats.EntityData.Leafs.Append("fdt-ecc-1b-err-cnt", types.YLeaf{"FdtEcc1bErrCnt", aggrStats.FdtEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fdt-ecc-2b-err-cnt", types.YLeaf{"FdtEcc2bErrCnt", aggrStats.FdtEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fdt-parity-err-cnt", types.YLeaf{"FdtParityErrCnt", aggrStats.FdtParityErrCnt})
    aggrStats.EntityData.Leafs.Append("fdt-crc-dropped-cell-cnt", types.YLeaf{"FdtCrcDroppedCellCnt", aggrStats.FdtCrcDroppedCellCnt})
    aggrStats.EntityData.Leafs.Append("fcr-control-cell-cnt", types.YLeaf{"FcrControlCellCnt", aggrStats.FcrControlCellCnt})
    aggrStats.EntityData.Leafs.Append("fcr-cell-drop-cnt", types.YLeaf{"FcrCellDropCnt", aggrStats.FcrCellDropCnt})
    aggrStats.EntityData.Leafs.Append("fcr-credit-cell-drop-cnt", types.YLeaf{"FcrCreditCellDropCnt", aggrStats.FcrCreditCellDropCnt})
    aggrStats.EntityData.Leafs.Append("fcr-fs-cell-drop-cnt", types.YLeaf{"FcrFsCellDropCnt", aggrStats.FcrFsCellDropCnt})
    aggrStats.EntityData.Leafs.Append("fcr-rt-cell-drop-cnt", types.YLeaf{"FcrRtCellDropCnt", aggrStats.FcrRtCellDropCnt})
    aggrStats.EntityData.Leafs.Append("fcr-ecc-1b-err-cnt", types.YLeaf{"FcrEcc1bErrCnt", aggrStats.FcrEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fcr-ecc-2b-err-cnt", types.YLeaf{"FcrEcc2bErrCnt", aggrStats.FcrEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fdr-data-cell-cnt", types.YLeaf{"FdrDataCellCnt", aggrStats.FdrDataCellCnt})
    aggrStats.EntityData.Leafs.Append("fdr-data-byte-cnt", types.YLeaf{"FdrDataByteCnt", aggrStats.FdrDataByteCnt})
    aggrStats.EntityData.Leafs.Append("fdr-crc-dropped-pck-cnt", types.YLeaf{"FdrCrcDroppedPckCnt", aggrStats.FdrCrcDroppedPckCnt})
    aggrStats.EntityData.Leafs.Append("fdr-p-pkt-cnt", types.YLeaf{"FdrPPktCnt", aggrStats.FdrPPktCnt})
    aggrStats.EntityData.Leafs.Append("fdr-prm-error-filter-cnt", types.YLeaf{"FdrPrmErrorFilterCnt", aggrStats.FdrPrmErrorFilterCnt})
    aggrStats.EntityData.Leafs.Append("fdr-sec-error-filter-cnt", types.YLeaf{"FdrSecErrorFilterCnt", aggrStats.FdrSecErrorFilterCnt})
    aggrStats.EntityData.Leafs.Append("fdr-prm-ecc-1b-err-cnt", types.YLeaf{"FdrPrmEcc1bErrCnt", aggrStats.FdrPrmEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fdr-prm-ecc-2b-err-cnt", types.YLeaf{"FdrPrmEcc2bErrCnt", aggrStats.FdrPrmEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fdr-sec-ecc-1b-err-cnt", types.YLeaf{"FdrSecEcc1bErrCnt", aggrStats.FdrSecEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fdr-sec-ecc-2b-err-cnt", types.YLeaf{"FdrSecEcc2bErrCnt", aggrStats.FdrSecEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-ecc-1b-err-cnt", types.YLeaf{"EgqEcc1bErrCnt", aggrStats.EgqEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-ecc-2b-err-cnt", types.YLeaf{"EgqEcc2bErrCnt", aggrStats.EgqEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-parity-err-cnt", types.YLeaf{"EgqParityErrCnt", aggrStats.EgqParityErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-dbf-ecc-1b-err-cnt", types.YLeaf{"EgqDbfEcc1bErrCnt", aggrStats.EgqDbfEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-dbf-ecc-2b-err-cnt", types.YLeaf{"EgqDbfEcc2bErrCnt", aggrStats.EgqDbfEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-empty-mcid-counter", types.YLeaf{"EgqEmptyMcidCounter", aggrStats.EgqEmptyMcidCounter})
    aggrStats.EntityData.Leafs.Append("egq-rqp-discard-packet-counter", types.YLeaf{"EgqRqpDiscardPacketCounter", aggrStats.EgqRqpDiscardPacketCounter})
    aggrStats.EntityData.Leafs.Append("egq-ehp-discard-packet-counter", types.YLeaf{"EgqEhpDiscardPacketCounter", aggrStats.EgqEhpDiscardPacketCounter})
    aggrStats.EntityData.Leafs.Append("egq-ipt-pkt-cnt", types.YLeaf{"EgqIptPktCnt", aggrStats.EgqIptPktCnt})
    aggrStats.EntityData.Leafs.Append("epni-epe-pkt-cnt", types.YLeaf{"EpniEpePktCnt", aggrStats.EpniEpePktCnt})
    aggrStats.EntityData.Leafs.Append("epni-epe-byte-cnt", types.YLeaf{"EpniEpeByteCnt", aggrStats.EpniEpeByteCnt})
    aggrStats.EntityData.Leafs.Append("epni-epe-discard-pkt-cnt", types.YLeaf{"EpniEpeDiscardPktCnt", aggrStats.EpniEpeDiscardPktCnt})
    aggrStats.EntityData.Leafs.Append("epni-ecc-1b-err-cnt", types.YLeaf{"EpniEcc1bErrCnt", aggrStats.EpniEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("epni-ecc-2b-err-cnt", types.YLeaf{"EpniEcc2bErrCnt", aggrStats.EpniEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("epni-parity-err-cnt", types.YLeaf{"EpniParityErrCnt", aggrStats.EpniParityErrCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-ucast-pkt-cnt", types.YLeaf{"EgqPqpUcastPktCnt", aggrStats.EgqPqpUcastPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-ucast-h-pkt-cnt", types.YLeaf{"EgqPqpUcastHPktCnt", aggrStats.EgqPqpUcastHPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-ucast-l-pkt-cnt", types.YLeaf{"EgqPqpUcastLPktCnt", aggrStats.EgqPqpUcastLPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-ucast-bytes-cnt", types.YLeaf{"EgqPqpUcastBytesCnt", aggrStats.EgqPqpUcastBytesCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-ucast-discard-pkt-cnt", types.YLeaf{"EgqPqpUcastDiscardPktCnt", aggrStats.EgqPqpUcastDiscardPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-mcast-pkt-cnt", types.YLeaf{"EgqPqpMcastPktCnt", aggrStats.EgqPqpMcastPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-mcast-h-pkt-cnt", types.YLeaf{"EgqPqpMcastHPktCnt", aggrStats.EgqPqpMcastHPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-mcast-l-pkt-cnt", types.YLeaf{"EgqPqpMcastLPktCnt", aggrStats.EgqPqpMcastLPktCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-mcast-bytes-cnt", types.YLeaf{"EgqPqpMcastBytesCnt", aggrStats.EgqPqpMcastBytesCnt})
    aggrStats.EntityData.Leafs.Append("egq-pqp-mcast-discard-pkt-cnt", types.YLeaf{"EgqPqpMcastDiscardPktCnt", aggrStats.EgqPqpMcastDiscardPktCnt})
    aggrStats.EntityData.Leafs.Append("fct-control-cell-cnt", types.YLeaf{"FctControlCellCnt", aggrStats.FctControlCellCnt})
    aggrStats.EntityData.Leafs.Append("fct-unrch-crdt-cnt", types.YLeaf{"FctUnrchCrdtCnt", aggrStats.FctUnrchCrdtCnt})
    aggrStats.EntityData.Leafs.Append("idr-reassembly-errors", types.YLeaf{"IdrReassemblyErrors", aggrStats.IdrReassemblyErrors})
    aggrStats.EntityData.Leafs.Append("idr-mmu-ecc-1b-err-cnt", types.YLeaf{"IdrMmuEcc1bErrCnt", aggrStats.IdrMmuEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("idr-mmu-ecc-2b-err-cnt", types.YLeaf{"IdrMmuEcc2bErrCnt", aggrStats.IdrMmuEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-packets0-cnt", types.YLeaf{"IdrDiscardedPackets0Cnt", aggrStats.IdrDiscardedPackets0Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-packets1-cnt", types.YLeaf{"IdrDiscardedPackets1Cnt", aggrStats.IdrDiscardedPackets1Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-packets2-cnt", types.YLeaf{"IdrDiscardedPackets2Cnt", aggrStats.IdrDiscardedPackets2Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-packets3-cnt", types.YLeaf{"IdrDiscardedPackets3Cnt", aggrStats.IdrDiscardedPackets3Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-octets0-cnt", types.YLeaf{"IdrDiscardedOctets0Cnt", aggrStats.IdrDiscardedOctets0Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-octets1-cnt", types.YLeaf{"IdrDiscardedOctets1Cnt", aggrStats.IdrDiscardedOctets1Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-octets2-cnt", types.YLeaf{"IdrDiscardedOctets2Cnt", aggrStats.IdrDiscardedOctets2Cnt})
    aggrStats.EntityData.Leafs.Append("idr-discarded-octets3-cnt", types.YLeaf{"IdrDiscardedOctets3Cnt", aggrStats.IdrDiscardedOctets3Cnt})
    aggrStats.EntityData.Leafs.Append("mmu-ecc-1b-err-cnt", types.YLeaf{"MmuEcc1bErrCnt", aggrStats.MmuEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("mmu-ecc-2b-err-cnt", types.YLeaf{"MmuEcc2bErrCnt", aggrStats.MmuEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("oamp-parity-err-cnt", types.YLeaf{"OampParityErrCnt", aggrStats.OampParityErrCnt})
    aggrStats.EntityData.Leafs.Append("oamp-ecc-1b-err-cnt", types.YLeaf{"OampEcc1bErrCnt", aggrStats.OampEcc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("oamp-ecc-2b-err-cnt", types.YLeaf{"OampEcc2bErrCnt", aggrStats.OampEcc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("crps-parity-err-cnt", types.YLeaf{"CrpsParityErrCnt", aggrStats.CrpsParityErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs0TstRxErrCnt", aggrStats.Fmac0Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs0TstRxErrCnt", aggrStats.Fmac1Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs0TstRxErrCnt", aggrStats.Fmac2Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs0TstRxErrCnt", aggrStats.Fmac3Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs0TstRxErrCnt", aggrStats.Fmac4Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs0TstRxErrCnt", aggrStats.Fmac5Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs0TstRxErrCnt", aggrStats.Fmac6Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs0TstRxErrCnt", aggrStats.Fmac7Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs0TstRxErrCnt", aggrStats.Fmac8Kpcs0TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs1TstRxErrCnt", aggrStats.Fmac0Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs1TstRxErrCnt", aggrStats.Fmac1Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs1TstRxErrCnt", aggrStats.Fmac2Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs1TstRxErrCnt", aggrStats.Fmac3Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs1TstRxErrCnt", aggrStats.Fmac4Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs1TstRxErrCnt", aggrStats.Fmac5Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs1TstRxErrCnt", aggrStats.Fmac6Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs1TstRxErrCnt", aggrStats.Fmac7Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs1TstRxErrCnt", aggrStats.Fmac8Kpcs1TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs2TstRxErrCnt", aggrStats.Fmac0Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs2TstRxErrCnt", aggrStats.Fmac1Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs2TstRxErrCnt", aggrStats.Fmac2Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs2TstRxErrCnt", aggrStats.Fmac3Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs2TstRxErrCnt", aggrStats.Fmac4Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs2TstRxErrCnt", aggrStats.Fmac5Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs2TstRxErrCnt", aggrStats.Fmac6Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs2TstRxErrCnt", aggrStats.Fmac7Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs2TstRxErrCnt", aggrStats.Fmac8Kpcs2TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs3TstRxErrCnt", aggrStats.Fmac0Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs3TstRxErrCnt", aggrStats.Fmac1Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs3TstRxErrCnt", aggrStats.Fmac2Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs3TstRxErrCnt", aggrStats.Fmac3Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs3TstRxErrCnt", aggrStats.Fmac4Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs3TstRxErrCnt", aggrStats.Fmac5Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs3TstRxErrCnt", aggrStats.Fmac6Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs3TstRxErrCnt", aggrStats.Fmac7Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs3TstRxErrCnt", aggrStats.Fmac8Kpcs3TstRxErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-tst0-err-cnt", types.YLeaf{"Fmac0Tst0ErrCnt", aggrStats.Fmac0Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-tst0-err-cnt", types.YLeaf{"Fmac1Tst0ErrCnt", aggrStats.Fmac1Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-tst0-err-cnt", types.YLeaf{"Fmac2Tst0ErrCnt", aggrStats.Fmac2Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-tst0-err-cnt", types.YLeaf{"Fmac3Tst0ErrCnt", aggrStats.Fmac3Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-tst0-err-cnt", types.YLeaf{"Fmac4Tst0ErrCnt", aggrStats.Fmac4Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-tst0-err-cnt", types.YLeaf{"Fmac5Tst0ErrCnt", aggrStats.Fmac5Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-tst0-err-cnt", types.YLeaf{"Fmac6Tst0ErrCnt", aggrStats.Fmac6Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-tst0-err-cnt", types.YLeaf{"Fmac7Tst0ErrCnt", aggrStats.Fmac7Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-tst0-err-cnt", types.YLeaf{"Fmac8Tst0ErrCnt", aggrStats.Fmac8Tst0ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-tst1-err-cnt", types.YLeaf{"Fmac0Tst1ErrCnt", aggrStats.Fmac0Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-tst1-err-cnt", types.YLeaf{"Fmac1Tst1ErrCnt", aggrStats.Fmac1Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-tst1-err-cnt", types.YLeaf{"Fmac2Tst1ErrCnt", aggrStats.Fmac2Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-tst1-err-cnt", types.YLeaf{"Fmac3Tst1ErrCnt", aggrStats.Fmac3Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-tst1-err-cnt", types.YLeaf{"Fmac4Tst1ErrCnt", aggrStats.Fmac4Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-tst1-err-cnt", types.YLeaf{"Fmac5Tst1ErrCnt", aggrStats.Fmac5Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-tst1-err-cnt", types.YLeaf{"Fmac6Tst1ErrCnt", aggrStats.Fmac6Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-tst1-err-cnt", types.YLeaf{"Fmac7Tst1ErrCnt", aggrStats.Fmac7Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-tst1-err-cnt", types.YLeaf{"Fmac8Tst1ErrCnt", aggrStats.Fmac8Tst1ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-tst2-err-cnt", types.YLeaf{"Fmac0Tst2ErrCnt", aggrStats.Fmac0Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-tst2-err-cnt", types.YLeaf{"Fmac1Tst2ErrCnt", aggrStats.Fmac1Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-tst2-err-cnt", types.YLeaf{"Fmac2Tst2ErrCnt", aggrStats.Fmac2Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-tst2-err-cnt", types.YLeaf{"Fmac3Tst2ErrCnt", aggrStats.Fmac3Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-tst2-err-cnt", types.YLeaf{"Fmac4Tst2ErrCnt", aggrStats.Fmac4Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-tst2-err-cnt", types.YLeaf{"Fmac5Tst2ErrCnt", aggrStats.Fmac5Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-tst2-err-cnt", types.YLeaf{"Fmac6Tst2ErrCnt", aggrStats.Fmac6Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-tst2-err-cnt", types.YLeaf{"Fmac7Tst2ErrCnt", aggrStats.Fmac7Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-tst2-err-cnt", types.YLeaf{"Fmac8Tst2ErrCnt", aggrStats.Fmac8Tst2ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-tst3-err-cnt", types.YLeaf{"Fmac0Tst3ErrCnt", aggrStats.Fmac0Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-tst3-err-cnt", types.YLeaf{"Fmac1Tst3ErrCnt", aggrStats.Fmac1Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-tst3-err-cnt", types.YLeaf{"Fmac2Tst3ErrCnt", aggrStats.Fmac2Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-tst3-err-cnt", types.YLeaf{"Fmac3Tst3ErrCnt", aggrStats.Fmac3Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-tst3-err-cnt", types.YLeaf{"Fmac4Tst3ErrCnt", aggrStats.Fmac4Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-tst3-err-cnt", types.YLeaf{"Fmac5Tst3ErrCnt", aggrStats.Fmac5Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-tst3-err-cnt", types.YLeaf{"Fmac6Tst3ErrCnt", aggrStats.Fmac6Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-tst3-err-cnt", types.YLeaf{"Fmac7Tst3ErrCnt", aggrStats.Fmac7Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-tst3-err-cnt", types.YLeaf{"Fmac8Tst3ErrCnt", aggrStats.Fmac8Tst3ErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-ecc-1b-err-cnt", types.YLeaf{"Fmac0Ecc1bErrCnt", aggrStats.Fmac0Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-ecc-1b-err-cnt", types.YLeaf{"Fmac1Ecc1bErrCnt", aggrStats.Fmac1Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-ecc-1b-err-cnt", types.YLeaf{"Fmac2Ecc1bErrCnt", aggrStats.Fmac2Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-ecc-1b-err-cnt", types.YLeaf{"Fmac3Ecc1bErrCnt", aggrStats.Fmac3Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-ecc-1b-err-cnt", types.YLeaf{"Fmac4Ecc1bErrCnt", aggrStats.Fmac4Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-ecc-1b-err-cnt", types.YLeaf{"Fmac5Ecc1bErrCnt", aggrStats.Fmac5Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-ecc-1b-err-cnt", types.YLeaf{"Fmac6Ecc1bErrCnt", aggrStats.Fmac6Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-ecc-1b-err-cnt", types.YLeaf{"Fmac7Ecc1bErrCnt", aggrStats.Fmac7Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-ecc-1b-err-cnt", types.YLeaf{"Fmac8Ecc1bErrCnt", aggrStats.Fmac8Ecc1bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac0-ecc-2b-err-cnt", types.YLeaf{"Fmac0Ecc2bErrCnt", aggrStats.Fmac0Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac1-ecc-2b-err-cnt", types.YLeaf{"Fmac1Ecc2bErrCnt", aggrStats.Fmac1Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac2-ecc-2b-err-cnt", types.YLeaf{"Fmac2Ecc2bErrCnt", aggrStats.Fmac2Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac3-ecc-2b-err-cnt", types.YLeaf{"Fmac3Ecc2bErrCnt", aggrStats.Fmac3Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac4-ecc-2b-err-cnt", types.YLeaf{"Fmac4Ecc2bErrCnt", aggrStats.Fmac4Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac5-ecc-2b-err-cnt", types.YLeaf{"Fmac5Ecc2bErrCnt", aggrStats.Fmac5Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac6-ecc-2b-err-cnt", types.YLeaf{"Fmac6Ecc2bErrCnt", aggrStats.Fmac6Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac7-ecc-2b-err-cnt", types.YLeaf{"Fmac7Ecc2bErrCnt", aggrStats.Fmac7Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("fmac8-ecc-2b-err-cnt", types.YLeaf{"Fmac8Ecc2bErrCnt", aggrStats.Fmac8Ecc2bErrCnt})
    aggrStats.EntityData.Leafs.Append("olp-incoming-bad-identifier-counter", types.YLeaf{"OlpIncomingBadIdentifierCounter", aggrStats.OlpIncomingBadIdentifierCounter})
    aggrStats.EntityData.Leafs.Append("olp-incoming-bad-reassembly-counter", types.YLeaf{"OlpIncomingBadReassemblyCounter", aggrStats.OlpIncomingBadReassemblyCounter})
    aggrStats.EntityData.Leafs.Append("cfc-parity-err-cnt", types.YLeaf{"CfcParityErrCnt", aggrStats.CfcParityErrCnt})
    aggrStats.EntityData.Leafs.Append("cfc-ilkn0-oob-rx-crc-err-cntr", types.YLeaf{"CfcIlkn0OobRxCrcErrCntr", aggrStats.CfcIlkn0OobRxCrcErrCntr})
    aggrStats.EntityData.Leafs.Append("cfc-ilkn1-oob-rx-crc-err-cntr", types.YLeaf{"CfcIlkn1OobRxCrcErrCntr", aggrStats.CfcIlkn1OobRxCrcErrCntr})
    aggrStats.EntityData.Leafs.Append("cfc-spi-oob-rx0-frm-err-cnt", types.YLeaf{"CfcSpiOobRx0FrmErrCnt", aggrStats.CfcSpiOobRx0FrmErrCnt})
    aggrStats.EntityData.Leafs.Append("cfc-spi-oob-rx0-dip2-err-cnt", types.YLeaf{"CfcSpiOobRx0Dip2ErrCnt", aggrStats.CfcSpiOobRx0Dip2ErrCnt})
    aggrStats.EntityData.Leafs.Append("cfc-spi-oob-rx1-frm-err-cnt", types.YLeaf{"CfcSpiOobRx1FrmErrCnt", aggrStats.CfcSpiOobRx1FrmErrCnt})
    aggrStats.EntityData.Leafs.Append("cfc-spi-oob-rx1-dip2-err-cnt", types.YLeaf{"CfcSpiOobRx1Dip2ErrCnt", aggrStats.CfcSpiOobRx1Dip2ErrCnt})
    aggrStats.EntityData.Leafs.Append("cgm-cgm-uc-pd-dropped-cnt", types.YLeaf{"CgmCgmUcPdDroppedCnt", aggrStats.CgmCgmUcPdDroppedCnt})
    aggrStats.EntityData.Leafs.Append("cgm-cgm-mc-rep-pd-dropped-cnt", types.YLeaf{"CgmCgmMcRepPdDroppedCnt", aggrStats.CgmCgmMcRepPdDroppedCnt})
    aggrStats.EntityData.Leafs.Append("cgm-cgm-uc-db-dropped-by-rqp-cnt", types.YLeaf{"CgmCgmUcDbDroppedByRqpCnt", aggrStats.CgmCgmUcDbDroppedByRqpCnt})
    aggrStats.EntityData.Leafs.Append("cgm-cgm-uc-db-dropped-by-pqp-cnt", types.YLeaf{"CgmCgmUcDbDroppedByPqpCnt", aggrStats.CgmCgmUcDbDroppedByPqpCnt})
    aggrStats.EntityData.Leafs.Append("cgm-cgm-mc-rep-db-dropped-cnt", types.YLeaf{"CgmCgmMcRepDbDroppedCnt", aggrStats.CgmCgmMcRepDbDroppedCnt})
    aggrStats.EntityData.Leafs.Append("cgm-cgm-mc-db-dropped-cnt", types.YLeaf{"CgmCgmMcDbDroppedCnt", aggrStats.CgmCgmMcDbDroppedCnt})
    aggrStats.EntityData.Leafs.Append("drca-full-err-cnt", types.YLeaf{"DrcaFullErrCnt", aggrStats.DrcaFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drca-single-err-cnt", types.YLeaf{"DrcaSingleErrCnt", aggrStats.DrcaSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drca-calib-bist-full-err-cnt", types.YLeaf{"DrcaCalibBistFullErrCnt", aggrStats.DrcaCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drca-loopback-full-err-cnt", types.YLeaf{"DrcaLoopbackFullErrCnt", aggrStats.DrcaLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcb-full-err-cnt", types.YLeaf{"DrcbFullErrCnt", aggrStats.DrcbFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcb-single-err-cnt", types.YLeaf{"DrcbSingleErrCnt", aggrStats.DrcbSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drcb-calib-bist-full-err-cnt", types.YLeaf{"DrcbCalibBistFullErrCnt", aggrStats.DrcbCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcb-loopback-full-err-cnt", types.YLeaf{"DrcbLoopbackFullErrCnt", aggrStats.DrcbLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcc-full-err-cnt", types.YLeaf{"DrccFullErrCnt", aggrStats.DrccFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcc-single-err-cnt", types.YLeaf{"DrccSingleErrCnt", aggrStats.DrccSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drcc-calib-bist-full-err-cnt", types.YLeaf{"DrccCalibBistFullErrCnt", aggrStats.DrccCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcc-loopback-full-err-cnt", types.YLeaf{"DrccLoopbackFullErrCnt", aggrStats.DrccLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcd-full-err-cnt", types.YLeaf{"DrcdFullErrCnt", aggrStats.DrcdFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcd-single-err-cnt", types.YLeaf{"DrcdSingleErrCnt", aggrStats.DrcdSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drcd-calib-bist-full-err-cnt", types.YLeaf{"DrcdCalibBistFullErrCnt", aggrStats.DrcdCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcd-loopback-full-err-cnt", types.YLeaf{"DrcdLoopbackFullErrCnt", aggrStats.DrcdLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drce-full-err-cnt", types.YLeaf{"DrceFullErrCnt", aggrStats.DrceFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drce-single-err-cnt", types.YLeaf{"DrceSingleErrCnt", aggrStats.DrceSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drce-calib-bist-full-err-cnt", types.YLeaf{"DrceCalibBistFullErrCnt", aggrStats.DrceCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drce-loopback-full-err-cnt", types.YLeaf{"DrceLoopbackFullErrCnt", aggrStats.DrceLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcf-full-err-cnt", types.YLeaf{"DrcfFullErrCnt", aggrStats.DrcfFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcf-single-err-cnt", types.YLeaf{"DrcfSingleErrCnt", aggrStats.DrcfSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drcf-calib-bist-full-err-cnt", types.YLeaf{"DrcfCalibBistFullErrCnt", aggrStats.DrcfCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcf-loopback-full-err-cnt", types.YLeaf{"DrcfLoopbackFullErrCnt", aggrStats.DrcfLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcg-full-err-cnt", types.YLeaf{"DrcgFullErrCnt", aggrStats.DrcgFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcg-single-err-cnt", types.YLeaf{"DrcgSingleErrCnt", aggrStats.DrcgSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drcg-calib-bist-full-err-cnt", types.YLeaf{"DrcgCalibBistFullErrCnt", aggrStats.DrcgCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcg-loopback-full-err-cnt", types.YLeaf{"DrcgLoopbackFullErrCnt", aggrStats.DrcgLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drch-full-err-cnt", types.YLeaf{"DrchFullErrCnt", aggrStats.DrchFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drch-single-err-cnt", types.YLeaf{"DrchSingleErrCnt", aggrStats.DrchSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drch-calib-bist-full-err-cnt", types.YLeaf{"DrchCalibBistFullErrCnt", aggrStats.DrchCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drch-loopback-full-err-cnt", types.YLeaf{"DrchLoopbackFullErrCnt", aggrStats.DrchLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcbroadcast-full-err-cnt", types.YLeaf{"DrcbroadcastFullErrCnt", aggrStats.DrcbroadcastFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcbroadcast-single-err-cnt", types.YLeaf{"DrcbroadcastSingleErrCnt", aggrStats.DrcbroadcastSingleErrCnt})
    aggrStats.EntityData.Leafs.Append("drcbroadcast-calib-bist-full-err-cnt", types.YLeaf{"DrcbroadcastCalibBistFullErrCnt", aggrStats.DrcbroadcastCalibBistFullErrCnt})
    aggrStats.EntityData.Leafs.Append("drcbroadcast-loopback-full-err-cnt", types.YLeaf{"DrcbroadcastLoopbackFullErrCnt", aggrStats.DrcbroadcastLoopbackFullErrCnt})
    aggrStats.EntityData.Leafs.Append("otn-mode", types.YLeaf{"OtnMode", aggrStats.OtnMode})
    aggrStats.EntityData.Leafs.Append("num-ports", types.YLeaf{"NumPorts", aggrStats.NumPorts})

    aggrStats.EntityData.YListKeys = []string {}

    return &(aggrStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats_AggrStatsOtn
// aggr stats otn
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats_AggrStatsOtn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // IL TotalRxPktCount. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTotalRxPktCount interface{}

    // IL TotalTxPktCount. The type is interface{} with range:
    // 0..18446744073709551615.
    IlTotalTxPktCount interface{}
}

func (aggrStatsOtn *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_AggrStats_AggrStatsOtn) GetEntityData() *types.CommonEntityData {
    aggrStatsOtn.EntityData.YFilter = aggrStatsOtn.YFilter
    aggrStatsOtn.EntityData.YangName = "aggr-stats-otn"
    aggrStatsOtn.EntityData.BundleName = "cisco_ios_xr"
    aggrStatsOtn.EntityData.ParentYangName = "aggr-stats"
    aggrStatsOtn.EntityData.SegmentPath = "aggr-stats-otn" + types.AddNoKeyToken(aggrStatsOtn)
    aggrStatsOtn.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/pbc-statistics/pbc-stats/aggr-stats/" + aggrStatsOtn.EntityData.SegmentPath
    aggrStatsOtn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggrStatsOtn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggrStatsOtn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggrStatsOtn.EntityData.Children = types.NewOrderedMap()
    aggrStatsOtn.EntityData.Leafs = types.NewOrderedMap()
    aggrStatsOtn.EntityData.Leafs.Append("il-total-rx-pkt-count", types.YLeaf{"IlTotalRxPktCount", aggrStatsOtn.IlTotalRxPktCount})
    aggrStatsOtn.EntityData.Leafs.Append("il-total-tx-pkt-count", types.YLeaf{"IlTotalTxPktCount", aggrStatsOtn.IlTotalTxPktCount})

    aggrStatsOtn.EntityData.YListKeys = []string {}

    return &(aggrStatsOtn.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_OvfStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CMIC Cmc0PktCountTxPkt. The type is string with length: 0..6.
    CmicCmc0PktCountTxPkt interface{}

    // CMIC Cmc0PktCountRxPkt. The type is string with length: 0..6.
    CmicCmc0PktCountRxPkt interface{}

    // NBI StatRxBurstsErrCnt. The type is string with length: 0..6.
    NbiStatRxBurstsErrCnt interface{}

    // NBI Ecc 1bErrCnt. The type is string with length: 0..6.
    NbiEcc1bErrCnt interface{}

    // NBI Ecc 2bErrCnt. The type is string with length: 0..6.
    NbiEcc2bErrCnt interface{}

    // NBI ParityErrCnt. The type is string with length: 0..6.
    NbiParityErrCnt interface{}

    // NBI RxIlknCrc32ErrCnt. The type is string with length: 0..6.
    NbiRxIlknCrc32ErrCnt interface{}

    // NBI RxIlkn0Crc24ErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0Crc24ErrCnt interface{}

    // NBI RxIlkn0BurstErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0BurstErrCnt interface{}

    // NBI RxIlkn0MissSopErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0MissSopErrCnt interface{}

    // NBI RxIlkn0MissEopErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0MissEopErrCnt interface{}

    // NBI RxIlkn0MisalignedCnt. The type is string with length: 0..6.
    NbiRxIlkn0MisalignedCnt interface{}

    // NBI RxIlkn1Crc24ErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1Crc24ErrCnt interface{}

    // NBI RxIlkn1BurstErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1BurstErrCnt interface{}

    // NBI RxIlkn1MissSopErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1MissSopErrCnt interface{}

    // NBI RxIlkn1MissEopErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1MissEopErrCnt interface{}

    // NBI RxIlkn1MisalignedCnt. The type is string with length: 0..6.
    NbiRxIlkn1MisalignedCnt interface{}

    // NBI TxIlkn1FlushedBurstsCnt. The type is string with length: 0..6.
    NbiTxIlkn1FlushedBurstsCnt interface{}

    // NBI RxIlkn0RetransCRC24ErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0RetransCrc24ErrCnt interface{}

    // NBI RxIlkn0RetransRetryErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0RetransRetryErrCnt interface{}

    // NBI RxIlkn0RetransWdogErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0RetransWdogErrCnt interface{}

    // NBI RxIlkn0RetransWrapAfterDiscErrCnt. The type is string with length:
    // 0..6.
    NbiRxIlkn0RetransWrapAfterDiscErrCnt interface{}

    // NBI RxIlkn0RetransWrapB4DiscErrCnt. The type is string with length: 0..6.
    NbiRxIlkn0RetransWrapB4DiscErrCnt interface{}

    // NBI RxIlkn0RetransReachedTimeoutErrCnt. The type is string with length:
    // 0..6.
    NbiRxIlkn0RetransReachedTimeoutErrCnt interface{}

    // NBI RxIlkn1RetransCRC24ErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1RetransCrc24ErrCnt interface{}

    // NBI RxIlkn1RetransRetryErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1RetransRetryErrCnt interface{}

    // NBI RxIlkn1RetransWdogErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1RetransWdogErrCnt interface{}

    // NBI RxIlkn1RetransWrapAfterDiscErrCnt. The type is string with length:
    // 0..6.
    NbiRxIlkn1RetransWrapAfterDiscErrCnt interface{}

    // NBI RxIlkn1RetransWrapB4DiscErrCnt. The type is string with length: 0..6.
    NbiRxIlkn1RetransWrapB4DiscErrCnt interface{}

    // NBI RxIlkn1RetransReachedTimeoutErrCnt. The type is string with length:
    // 0..6.
    NbiRxIlkn1RetransReachedTimeoutErrCnt interface{}

    // NBI StatRxFrameErrCnt. The type is string with length: 0..6.
    NbiStatRxFrameErrCnt interface{}

    // NBI StatTxFrameErrCnt. The type is string with length: 0..6.
    NbiStatTxFrameErrCnt interface{}

    // NBI RxElkErrBurstsCnt. The type is string with length: 0..6.
    NbiRxElkErrBurstsCnt interface{}

    // NBI RxNumThrownEops. The type is string with length: 0..6.
    NbiRxNumThrownEops interface{}

    // NBI RxNumRunts. The type is string with length: 0..6.
    NbiRxNumRunts interface{}

    // NBI BistTxCrcErrBurstsCnt. The type is string with length: 0..6.
    NbiBistTxCrcErrBurstsCnt interface{}

    // NBI BistRxErrLengthBurstsCnt. The type is string with length: 0..6.
    NbiBistRxErrLengthBurstsCnt interface{}

    // NBI BistRxErrBurstIndexCnt. The type is string with length: 0..6.
    NbiBistRxErrBurstIndexCnt interface{}

    // NBI BistRxErrBctCnt. The type is string with length: 0..6.
    NbiBistRxErrBctCnt interface{}

    // NBI BistRxErrDataCnt. The type is string with length: 0..6.
    NbiBistRxErrDataCnt interface{}

    // NBI BistRxErrInCrcErrCnt. The type is string with length: 0..6.
    NbiBistRxErrInCrcErrCnt interface{}

    // NBI BistRxErrSobCnt. The type is string with length: 0..6.
    NbiBistRxErrSobCnt interface{}

    // NBI StatTxBurstsCnt. The type is string with length: 0..6.
    NbiStatTxBurstsCnt interface{}

    // NBI StatTxTotalLengCnt. The type is string with length: 0..6.
    NbiStatTxTotalLengCnt interface{}

    // RXAUI TotalTxPktCount. The type is string with length: 0..6.
    RxauiTotalTxPktCount interface{}

    // RXAUI TotalRxPktCount. The type is string with length: 0..6.
    RxauiTotalRxPktCount interface{}

    // RXAUI RxPktCountBcastPkt. The type is string with length: 0..6.
    RxauiRxPktCountBcastPkt interface{}

    // RXAUI TxPktCountBcastPkt. The type is string with length: 0..6.
    RxauiTxPktCountBcastPkt interface{}

    // RXAUI RxPktCountMcastPkt. The type is string with length: 0..6.
    RxauiRxPktCountMcastPkt interface{}

    // RXAUI TxPktCountMcastPkt. The type is string with length: 0..6.
    RxauiTxPktCountMcastPkt interface{}

    // RXAUI RxPktCountUcastPkt. The type is string with length: 0..6.
    RxauiRxPktCountUcastPkt interface{}

    // RXAUI TxPktCountUcastPkt. The type is string with length: 0..6.
    RxauiTxPktCountUcastPkt interface{}

    // RXAUI RxErrDropPktCnt. The type is string with length: 0..6.
    RxauiRxErrDropPktCnt interface{}

    // RXAUI TxErrDropPktCnt. The type is string with length: 0..6.
    RxauiTxErrDropPktCnt interface{}

    // RXAUI ByteCountTxPkt. The type is string with length: 0..6.
    RxauiByteCountTxPkt interface{}

    // RXAUI ByteCountRxPkt. The type is string with length: 0..6.
    RxauiByteCountRxPkt interface{}

    // RXAUI RxDscrdPktCnt. The type is string with length: 0..6.
    RxauiRxDscrdPktCnt interface{}

    // RXAUI TxDscrdPktCnt. The type is string with length: 0..6.
    RxauiTxDscrdPktCnt interface{}

    // IRE NifPacketCounter. The type is string with length: 0..6.
    IreNifPacketCounter interface{}

    // IL TotalRxPktCount. The type is string with length: 0..6.
    IlTotalRxPktCount interface{}

    // IL TotalTxPktCount. The type is string with length: 0..6.
    IlTotalTxPktCount interface{}

    // IL RxErrDropPktCnt. The type is string with length: 0..6.
    IlRxErrDropPktCnt interface{}

    // IL TxErrDropPktCnt. The type is string with length: 0..6.
    IlTxErrDropPktCnt interface{}

    // IL ByteCountTxPkt. The type is string with length: 0..6.
    IlByteCountTxPkt interface{}

    // IL ByteCountRxPkt. The type is string with length: 0..6.
    IlByteCountRxPkt interface{}

    // IL RxDscrdPktCnt. The type is string with length: 0..6.
    IlRxDscrdPktCnt interface{}

    // IL TxDscrdPktCnt. The type is string with length: 0..6.
    IlTxDscrdPktCnt interface{}

    // IL RxPktCountBcastPkt. The type is string with length: 0..6.
    IlRxPktCountBcastPkt interface{}

    // IL TxPktCountBcastPkt. The type is string with length: 0..6.
    IlTxPktCountBcastPkt interface{}

    // IL RxPktCountMcastPkt. The type is string with length: 0..6.
    IlRxPktCountMcastPkt interface{}

    // IL TxPktCountMcastPkt. The type is string with length: 0..6.
    IlTxPktCountMcastPkt interface{}

    // IL RxPktCountUcastPkt. The type is string with length: 0..6.
    IlRxPktCountUcastPkt interface{}

    // IL TxPktCountUcastPkt. The type is string with length: 0..6.
    IlTxPktCountUcastPkt interface{}

    // IQM EnqPktCnt. The type is string with length: 0..6.
    IqmEnqPktCnt interface{}

    // IQM EnqByteCnt. The type is string with length: 0..6.
    IqmEnqByteCnt interface{}

    // IQM DeqPktCnt. The type is string with length: 0..6.
    IqmDeqPktCnt interface{}

    // IQM DeqByteCnt. The type is string with length: 0..6.
    IqmDeqByteCnt interface{}

    // IQM TotDscrdPktCnt. The type is string with length: 0..6.
    IqmTotDscrdPktCnt interface{}

    // IQM TotDscrdByteCnt. The type is string with length: 0..6.
    IqmTotDscrdByteCnt interface{}

    // IQM Ecc 1bErrCnt. The type is string with length: 0..6.
    IqmEcc1bErrCnt interface{}

    // IQM Ecc 2bErrCnt. The type is string with length: 0..6.
    IqmEcc2bErrCnt interface{}

    // IQM ParityErrCnt. The type is string with length: 0..6.
    IqmParityErrCnt interface{}

    // IQM DeqDeletePktCnt. The type is string with length: 0..6.
    IqmDeqDeletePktCnt interface{}

    // IQM EcnDscrdMskPktCnt. The type is string with length: 0..6.
    IqmEcnDscrdMskPktCnt interface{}

    // IQM QTotDscrdPktCnt. The type is string with length: 0..6.
    IqmQTotDscrdPktCnt interface{}

    // IQM QDeqDeletePktCnt. The type is string with length: 0..6.
    IqmQDeqDeletePktCnt interface{}

    // IQM RjctDbPktCnt. The type is string with length: 0..6.
    IqmRjctDbPktCnt interface{}

    // IQM RjctBdbPktCnt. The type is string with length: 0..6.
    IqmRjctBdbPktCnt interface{}

    // IQM RjctBdbProtctPktCnt. The type is string with length: 0..6.
    IqmRjctBdbProtctPktCnt interface{}

    // IQM RjctOcBdPktCnt. The type is string with length: 0..6.
    IqmRjctOcBdPktCnt interface{}

    // IQM RjctSnErrPktCnt. The type is string with length: 0..6.
    IqmRjctSnErrPktCnt interface{}

    // IQM RjctMcErrPktCnt. The type is string with length: 0..6.
    IqmRjctMcErrPktCnt interface{}

    // IQM RjctRsrcErrPktCnt. The type is string with length: 0..6.
    IqmRjctRsrcErrPktCnt interface{}

    // IQM RjctQnvalidErrPktCnt. The type is string with length: 0..6.
    IqmRjctQnvalidErrPktCnt interface{}

    // IQM RjctCnmPktCnt. The type is string with length: 0..6.
    IqmRjctCnmPktCnt interface{}

    // IQM RjctDynSpacePktCnt. The type is string with length: 0..6.
    IqmRjctDynSpacePktCnt interface{}

    // IPT FdtPktCnt. The type is string with length: 0..6.
    IptFdtPktCnt interface{}

    // IPT Ecc 1bErrCnt. The type is string with length: 0..6.
    IptEcc1bErrCnt interface{}

    // IPT Ecc 2bErrCnt. The type is string with length: 0..6.
    IptEcc2bErrCnt interface{}

    // IPT ParityErrCnt. The type is string with length: 0..6.
    IptParityErrCnt interface{}

    // IPT CrcErrCnt. The type is string with length: 0..6.
    IptCrcErrCnt interface{}

    // IPT CrcErrDelBuffCnt. The type is string with length: 0..6.
    IptCrcErrDelBuffCnt interface{}

    // IPT CpuDelBuffCnt. The type is string with length: 0..6.
    IptCpuDelBuffCnt interface{}

    // IPT CpuRelBuffCnt. The type is string with length: 0..6.
    IptCpuRelBuffCnt interface{}

    // IPT CrcErrBuffFifoFullCnt. The type is string with length: 0..6.
    IptCrcErrBuffFifoFullCnt interface{}

    // FDT DataCellCnt. The type is string with length: 0..6.
    FdtDataCellCnt interface{}

    // FDT DataByteCnt. The type is string with length: 0..6.
    FdtDataByteCnt interface{}

    // FDT CrcDroppedPckCnt. The type is string with length: 0..6.
    FdtCrcDroppedPckCnt interface{}

    // FDT invalid destq drop cell cnt. The type is string with length: 0..6.
    FdtInvalidDestqDropCellCnt interface{}

    // FDT IndirectCommandCount. The type is string with length: 0..6.
    FdtIndirectCommandCount interface{}

    // FDT Ecc 1bErrCnt. The type is string with length: 0..6.
    FdtEcc1bErrCnt interface{}

    // FDT Ecc 2bErrCnt. The type is string with length: 0..6.
    FdtEcc2bErrCnt interface{}

    // FDT ParityErrCnt. The type is string with length: 0..6.
    FdtParityErrCnt interface{}

    // FDT CrcDroppedCellCnt. The type is string with length: 0..6.
    FdtCrcDroppedCellCnt interface{}

    // FCR ControlCellCnt. The type is string with length: 0..6.
    FcrControlCellCnt interface{}

    // FCR CellDropCnt. The type is string with length: 0..6.
    FcrCellDropCnt interface{}

    // FCR CreditCellDropCnt. The type is string with length: 0..6.
    FcrCreditCellDropCnt interface{}

    // FCR FSCellDropCnt. The type is string with length: 0..6.
    FcrFsCellDropCnt interface{}

    // FCR RTCellDropCnt. The type is string with length: 0..6.
    FcrRtCellDropCnt interface{}

    // FCR Ecc 1bErrCnt. The type is string with length: 0..6.
    FcrEcc1bErrCnt interface{}

    // FCR Ecc 2bErrCnt. The type is string with length: 0..6.
    FcrEcc2bErrCnt interface{}

    // FDR DataCellCnt. The type is string with length: 0..6.
    FdrDataCellCnt interface{}

    // FDR DataByteCnt. The type is string with length: 0..6.
    FdrDataByteCnt interface{}

    // FDR CrcDroppedPckCnt. The type is string with length: 0..6.
    FdrCrcDroppedPckCnt interface{}

    // FDR PPktCnt. The type is string with length: 0..6.
    FdrPPktCnt interface{}

    // FDR PrmErrorFilterCnt. The type is string with length: 0..6.
    FdrPrmErrorFilterCnt interface{}

    // FDR SecErrorFilterCnt. The type is string with length: 0..6.
    FdrSecErrorFilterCnt interface{}

    // FDR PrmEcc 1bErrCnt. The type is string with length: 0..6.
    FdrPrmEcc1bErrCnt interface{}

    // FDR PrmEcc 2bErrCnt. The type is string with length: 0..6.
    FdrPrmEcc2bErrCnt interface{}

    // FDR SecEcc 1bErrCnt. The type is string with length: 0..6.
    FdrSecEcc1bErrCnt interface{}

    // FDR SecEcc 2bErrCnt. The type is string with length: 0..6.
    FdrSecEcc2bErrCnt interface{}

    // EGQ Ecc 1bErrCnt. The type is string with length: 0..6.
    EgqEcc1bErrCnt interface{}

    // EGQ Ecc 2bErrCnt. The type is string with length: 0..6.
    EgqEcc2bErrCnt interface{}

    // EGQ ParityErrCnt. The type is string with length: 0..6.
    EgqParityErrCnt interface{}

    // EGQ DbfEcc 1bErrCnt. The type is string with length: 0..6.
    EgqDbfEcc1bErrCnt interface{}

    // EGQ DbfEcc 2bErrCnt. The type is string with length: 0..6.
    EgqDbfEcc2bErrCnt interface{}

    // EGQ EmptyMcidCounter. The type is string with length: 0..6.
    EgqEmptyMcidCounter interface{}

    // EGQ RqpDiscardPacketCounter. The type is string with length: 0..6.
    EgqRqpDiscardPacketCounter interface{}

    // EGQ EhpDiscardPacketCounter. The type is string with length: 0..6.
    EgqEhpDiscardPacketCounter interface{}

    // EGQ IptPktCnt. The type is string with length: 0..6.
    EgqIptPktCnt interface{}

    // EPNI EpePktCnt. The type is string with length: 0..6.
    EpniEpePktCnt interface{}

    // EPNI EpeByteCnt. The type is string with length: 0..6.
    EpniEpeByteCnt interface{}

    // EPNI EpeDiscardPktCnt. The type is string with length: 0..6.
    EpniEpeDiscardPktCnt interface{}

    // EPNI Ecc 1bErrCnt. The type is string with length: 0..6.
    EpniEcc1bErrCnt interface{}

    // EPNI Ecc 2bErrCnt. The type is string with length: 0..6.
    EpniEcc2bErrCnt interface{}

    // EPNI ParityErrCnt. The type is string with length: 0..6.
    EpniParityErrCnt interface{}

    // EGQ PqpUcastPktCnt. The type is string with length: 0..6.
    EgqPqpUcastPktCnt interface{}

    // EGQ PqpUcastHPktCnt. The type is string with length: 0..6.
    EgqPqpUcastHPktCnt interface{}

    // EGQ PqpUcastLPktCnt. The type is string with length: 0..6.
    EgqPqpUcastLPktCnt interface{}

    // EGQ PqpUcastBytesCnt. The type is string with length: 0..6.
    EgqPqpUcastBytesCnt interface{}

    // EGQ PqpUcastDiscardPktCnt. The type is string with length: 0..6.
    EgqPqpUcastDiscardPktCnt interface{}

    // EGQ PqpMcastPktCnt. The type is string with length: 0..6.
    EgqPqpMcastPktCnt interface{}

    // EGQ PqpMcastHPktCnt. The type is string with length: 0..6.
    EgqPqpMcastHPktCnt interface{}

    // EGQ PqpMcastLPktCnt. The type is string with length: 0..6.
    EgqPqpMcastLPktCnt interface{}

    // EGQ PqpMcastBytesCnt. The type is string with length: 0..6.
    EgqPqpMcastBytesCnt interface{}

    // EGQ PqpMcastDiscardPktCnt. The type is string with length: 0..6.
    EgqPqpMcastDiscardPktCnt interface{}

    // FCT ControlCellCnt. The type is string with length: 0..6.
    FctControlCellCnt interface{}

    // FCT UnrchCrdtCnt. The type is string with length: 0..6.
    FctUnrchCrdtCnt interface{}

    // IDR ReassemblyErrors. The type is string with length: 0..6.
    IdrReassemblyErrors interface{}

    // IDR MmuEcc 1bErrCnt. The type is string with length: 0..6.
    IdrMmuEcc1bErrCnt interface{}

    // IDR MmuEcc 2bErrCnt. The type is string with length: 0..6.
    IdrMmuEcc2bErrCnt interface{}

    // IDR DiscardedPackets0Cnt. The type is string with length: 0..6.
    IdrDiscardedPackets0Cnt interface{}

    // IDR DiscardedPackets1Cnt. The type is string with length: 0..6.
    IdrDiscardedPackets1Cnt interface{}

    // IDR DiscardedPackets2Cnt. The type is string with length: 0..6.
    IdrDiscardedPackets2Cnt interface{}

    // IDR DiscardedPackets3Cnt. The type is string with length: 0..6.
    IdrDiscardedPackets3Cnt interface{}

    // IDR DiscardedOctets0Cnt. The type is string with length: 0..6.
    IdrDiscardedOctets0Cnt interface{}

    // IDR DiscardedOctets1Cnt. The type is string with length: 0..6.
    IdrDiscardedOctets1Cnt interface{}

    // IDR DiscardedOctets2Cnt. The type is string with length: 0..6.
    IdrDiscardedOctets2Cnt interface{}

    // IDR DiscardedOctets3Cnt. The type is string with length: 0..6.
    IdrDiscardedOctets3Cnt interface{}

    // MMU Ecc 1bErrCnt. The type is string with length: 0..6.
    MmuEcc1bErrCnt interface{}

    // MMU Ecc 2bErrCnt. The type is string with length: 0..6.
    MmuEcc2bErrCnt interface{}

    // OAMP ParityErrCnt. The type is string with length: 0..6.
    OampParityErrCnt interface{}

    // OAMP Ecc 1bErrCnt. The type is string with length: 0..6.
    OampEcc1bErrCnt interface{}

    // OAMP Ecc 2bErrCnt. The type is string with length: 0..6.
    OampEcc2bErrCnt interface{}

    // CRPS ParityErrCnt. The type is string with length: 0..6.
    CrpsParityErrCnt interface{}

    // FMAC0 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac0Kpcs0TstRxErrCnt interface{}

    // FMAC1 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac1Kpcs0TstRxErrCnt interface{}

    // FMAC2 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac2Kpcs0TstRxErrCnt interface{}

    // FMAC3 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac3Kpcs0TstRxErrCnt interface{}

    // FMAC4 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac4Kpcs0TstRxErrCnt interface{}

    // FMAC5 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac5Kpcs0TstRxErrCnt interface{}

    // FMAC6 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac6Kpcs0TstRxErrCnt interface{}

    // FMAC7 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac7Kpcs0TstRxErrCnt interface{}

    // FMAC8 Kpcs0TstRxErrCnt. The type is string with length: 0..6.
    Fmac8Kpcs0TstRxErrCnt interface{}

    // FMAC0 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac0Kpcs1TstRxErrCnt interface{}

    // FMAC1 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac1Kpcs1TstRxErrCnt interface{}

    // FMAC2 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac2Kpcs1TstRxErrCnt interface{}

    // FMAC3 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac3Kpcs1TstRxErrCnt interface{}

    // FMAC4 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac4Kpcs1TstRxErrCnt interface{}

    // FMAC5 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac5Kpcs1TstRxErrCnt interface{}

    // FMAC6 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac6Kpcs1TstRxErrCnt interface{}

    // FMAC7 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac7Kpcs1TstRxErrCnt interface{}

    // FMAC8 Kpcs1TstRxErrCnt. The type is string with length: 0..6.
    Fmac8Kpcs1TstRxErrCnt interface{}

    // FMAC0 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac0Kpcs2TstRxErrCnt interface{}

    // FMAC1 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac1Kpcs2TstRxErrCnt interface{}

    // FMAC2 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac2Kpcs2TstRxErrCnt interface{}

    // FMAC3 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac3Kpcs2TstRxErrCnt interface{}

    // FMAC4 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac4Kpcs2TstRxErrCnt interface{}

    // FMAC5 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac5Kpcs2TstRxErrCnt interface{}

    // FMAC6 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac6Kpcs2TstRxErrCnt interface{}

    // FMAC7 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac7Kpcs2TstRxErrCnt interface{}

    // FMAC8 Kpcs2TstRxErrCnt. The type is string with length: 0..6.
    Fmac8Kpcs2TstRxErrCnt interface{}

    // FMAC0 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac0Kpcs3TstRxErrCnt interface{}

    // FMAC1 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac1Kpcs3TstRxErrCnt interface{}

    // FMAC2 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac2Kpcs3TstRxErrCnt interface{}

    // FMAC3 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac3Kpcs3TstRxErrCnt interface{}

    // FMAC4 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac4Kpcs3TstRxErrCnt interface{}

    // FMAC5 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac5Kpcs3TstRxErrCnt interface{}

    // FMAC6 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac6Kpcs3TstRxErrCnt interface{}

    // FMAC7 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac7Kpcs3TstRxErrCnt interface{}

    // FMAC8 Kpcs3TstRxErrCnt. The type is string with length: 0..6.
    Fmac8Kpcs3TstRxErrCnt interface{}

    // FMAC0 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac0Tst0ErrCnt interface{}

    // FMAC1 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac1Tst0ErrCnt interface{}

    // FMAC2 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac2Tst0ErrCnt interface{}

    // FMAC3 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac3Tst0ErrCnt interface{}

    // FMAC4 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac4Tst0ErrCnt interface{}

    // FMAC5 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac5Tst0ErrCnt interface{}

    // FMAC6 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac6Tst0ErrCnt interface{}

    // FMAC7 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac7Tst0ErrCnt interface{}

    // FMAC8 Tst0ErrCnt. The type is string with length: 0..6.
    Fmac8Tst0ErrCnt interface{}

    // FMAC0 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac0Tst1ErrCnt interface{}

    // FMAC1 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac1Tst1ErrCnt interface{}

    // FMAC2 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac2Tst1ErrCnt interface{}

    // FMAC3 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac3Tst1ErrCnt interface{}

    // FMAC4 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac4Tst1ErrCnt interface{}

    // FMAC5 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac5Tst1ErrCnt interface{}

    // FMAC6 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac6Tst1ErrCnt interface{}

    // FMAC7 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac7Tst1ErrCnt interface{}

    // FMAC8 Tst1ErrCnt. The type is string with length: 0..6.
    Fmac8Tst1ErrCnt interface{}

    // FMAC0 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac0Tst2ErrCnt interface{}

    // FMAC1 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac1Tst2ErrCnt interface{}

    // FMAC2 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac2Tst2ErrCnt interface{}

    // FMAC3 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac3Tst2ErrCnt interface{}

    // FMAC4 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac4Tst2ErrCnt interface{}

    // FMAC5 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac5Tst2ErrCnt interface{}

    // FMAC6 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac6Tst2ErrCnt interface{}

    // FMAC7 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac7Tst2ErrCnt interface{}

    // FMAC8 Tst2ErrCnt. The type is string with length: 0..6.
    Fmac8Tst2ErrCnt interface{}

    // FMAC0 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac0Tst3ErrCnt interface{}

    // FMAC1 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac1Tst3ErrCnt interface{}

    // FMAC2 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac2Tst3ErrCnt interface{}

    // FMAC3 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac3Tst3ErrCnt interface{}

    // FMAC4 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac4Tst3ErrCnt interface{}

    // FMAC5 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac5Tst3ErrCnt interface{}

    // FMAC6 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac6Tst3ErrCnt interface{}

    // FMAC7 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac7Tst3ErrCnt interface{}

    // FMAC8 Tst3ErrCnt. The type is string with length: 0..6.
    Fmac8Tst3ErrCnt interface{}

    // FMAC0 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac0Ecc1bErrCnt interface{}

    // FMAC1 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac1Ecc1bErrCnt interface{}

    // FMAC2 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac2Ecc1bErrCnt interface{}

    // FMAC3 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac3Ecc1bErrCnt interface{}

    // FMAC4 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac4Ecc1bErrCnt interface{}

    // FMAC5 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac5Ecc1bErrCnt interface{}

    // FMAC6 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac6Ecc1bErrCnt interface{}

    // FMAC7 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac7Ecc1bErrCnt interface{}

    // FMAC8 Ecc 1bErrCnt. The type is string with length: 0..6.
    Fmac8Ecc1bErrCnt interface{}

    // FMAC0 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac0Ecc2bErrCnt interface{}

    // FMAC1 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac1Ecc2bErrCnt interface{}

    // FMAC2 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac2Ecc2bErrCnt interface{}

    // FMAC3 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac3Ecc2bErrCnt interface{}

    // FMAC4 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac4Ecc2bErrCnt interface{}

    // FMAC5 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac5Ecc2bErrCnt interface{}

    // FMAC6 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac6Ecc2bErrCnt interface{}

    // FMAC7 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac7Ecc2bErrCnt interface{}

    // FMAC8 Ecc 2bErrCnt. The type is string with length: 0..6.
    Fmac8Ecc2bErrCnt interface{}

    // OLP IncomingBadIdentifierCounter. The type is string with length: 0..6.
    OlpIncomingBadIdentifierCounter interface{}

    // OLP IncomingBadReassemblyCounter. The type is string with length: 0..6.
    OlpIncomingBadReassemblyCounter interface{}

    // CFC ParityErrCnt. The type is string with length: 0..6.
    CfcParityErrCnt interface{}

    // CFC Ilkn0OobRxCrcErrCntr. The type is string with length: 0..6.
    CfcIlkn0OobRxCrcErrCntr interface{}

    // CFC Ilkn1OobRxCrcErrCntr. The type is string with length: 0..6.
    CfcIlkn1OobRxCrcErrCntr interface{}

    // CFC SpiOobRx0FrmErrCnt. The type is string with length: 0..6.
    CfcSpiOobRx0FrmErrCnt interface{}

    // CFC SpiOobRx0Dip2ErrCnt. The type is string with length: 0..6.
    CfcSpiOobRx0Dip2ErrCnt interface{}

    // CFC SpiOobRx1FrmErrCnt. The type is string with length: 0..6.
    CfcSpiOobRx1FrmErrCnt interface{}

    // CFC SpiOobRx1Dip2ErrCnt. The type is string with length: 0..6.
    CfcSpiOobRx1Dip2ErrCnt interface{}

    // CGM CgmUcPdDroppedCnt. The type is string with length: 0..6.
    CgmCgmUcPdDroppedCnt interface{}

    // CGM CgmMcRepPdDroppedCnt. The type is string with length: 0..6.
    CgmCgmMcRepPdDroppedCnt interface{}

    // CGM CgmUcDbDroppedByRqpCnt. The type is string with length: 0..6.
    CgmCgmUcDbDroppedByRqpCnt interface{}

    // CGM CgmUcDbDroppedByPqpCnt. The type is string with length: 0..6.
    CgmCgmUcDbDroppedByPqpCnt interface{}

    // CGM CgmMcRepDbDroppedCnt. The type is string with length: 0..6.
    CgmCgmMcRepDbDroppedCnt interface{}

    // CGM CgmMcDbDroppedCnt. The type is string with length: 0..6.
    CgmCgmMcDbDroppedCnt interface{}

    // DRCA FullErrCnt. The type is string with length: 0..6.
    DrcaFullErrCnt interface{}

    // DRCA SingleErrCnt. The type is string with length: 0..6.
    DrcaSingleErrCnt interface{}

    // DRCA CalibBistFullErrCnt. The type is string with length: 0..6.
    DrcaCalibBistFullErrCnt interface{}

    // DRCA LoopbackFullErrCnt. The type is string with length: 0..6.
    DrcaLoopbackFullErrCnt interface{}

    // DRCB FullErrCnt. The type is string with length: 0..6.
    DrcbFullErrCnt interface{}

    // DRCB SingleErrCnt. The type is string with length: 0..6.
    DrcbSingleErrCnt interface{}

    // DRCB CalibBistFullErrCnt. The type is string with length: 0..6.
    DrcbCalibBistFullErrCnt interface{}

    // DRCB LoopbackFullErrCnt. The type is string with length: 0..6.
    DrcbLoopbackFullErrCnt interface{}

    // DRCC FullErrCnt. The type is string with length: 0..6.
    DrccFullErrCnt interface{}

    // DRCC SingleErrCnt. The type is string with length: 0..6.
    DrccSingleErrCnt interface{}

    // DRCC CalibBistFullErrCnt. The type is string with length: 0..6.
    DrccCalibBistFullErrCnt interface{}

    // DRCC LoopbackFullErrCnt. The type is string with length: 0..6.
    DrccLoopbackFullErrCnt interface{}

    // DRCD FullErrCnt. The type is string with length: 0..6.
    DrcdFullErrCnt interface{}

    // DRCD SingleErrCnt. The type is string with length: 0..6.
    DrcdSingleErrCnt interface{}

    // DRCD CalibBistFullErrCnt. The type is string with length: 0..6.
    DrcdCalibBistFullErrCnt interface{}

    // DRCD LoopbackFullErrCnt. The type is string with length: 0..6.
    DrcdLoopbackFullErrCnt interface{}

    // DRCE FullErrCnt. The type is string with length: 0..6.
    DrceFullErrCnt interface{}

    // DRCE SingleErrCnt. The type is string with length: 0..6.
    DrceSingleErrCnt interface{}

    // DRCE CalibBistFullErrCnt. The type is string with length: 0..6.
    DrceCalibBistFullErrCnt interface{}

    // DRCE LoopbackFullErrCnt. The type is string with length: 0..6.
    DrceLoopbackFullErrCnt interface{}

    // DRCF FullErrCnt. The type is string with length: 0..6.
    DrcfFullErrCnt interface{}

    // DRCF SingleErrCnt. The type is string with length: 0..6.
    DrcfSingleErrCnt interface{}

    // DRCF CalibBistFullErrCnt. The type is string with length: 0..6.
    DrcfCalibBistFullErrCnt interface{}

    // DRCF LoopbackFullErrCnt. The type is string with length: 0..6.
    DrcfLoopbackFullErrCnt interface{}

    // DRCG FullErrCnt. The type is string with length: 0..6.
    DrcgFullErrCnt interface{}

    // DRCG SingleErrCnt. The type is string with length: 0..6.
    DrcgSingleErrCnt interface{}

    // DRCG CalibBistFullErrCnt. The type is string with length: 0..6.
    DrcgCalibBistFullErrCnt interface{}

    // DRCG LoopbackFullErrCnt. The type is string with length: 0..6.
    DrcgLoopbackFullErrCnt interface{}

    // DRCH FullErrCnt. The type is string with length: 0..6.
    DrchFullErrCnt interface{}

    // DRCH SingleErrCnt. The type is string with length: 0..6.
    DrchSingleErrCnt interface{}

    // DRCH CalibBistFullErrCnt. The type is string with length: 0..6.
    DrchCalibBistFullErrCnt interface{}

    // DRCH LoopbackFullErrCnt. The type is string with length: 0..6.
    DrchLoopbackFullErrCnt interface{}

    // DRCBROADCAST FullErrCnt. The type is string with length: 0..6.
    DrcbroadcastFullErrCnt interface{}

    // DRCBROADCAST SingleErrCnt. The type is string with length: 0..6.
    DrcbroadcastSingleErrCnt interface{}

    // DRCBROADCAST CalibBistFullErrCnt. The type is string with length: 0..6.
    DrcbroadcastCalibBistFullErrCnt interface{}

    // DRCBROADCAST LoopbackFullErrCnt. The type is string with length: 0..6.
    DrcbroadcastLoopbackFullErrCnt interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_PbcStatistics_PbcStats_OvfStatus) GetEntityData() *types.CommonEntityData {
    ovfStatus.EntityData.YFilter = ovfStatus.YFilter
    ovfStatus.EntityData.YangName = "ovf-status"
    ovfStatus.EntityData.BundleName = "cisco_ios_xr"
    ovfStatus.EntityData.ParentYangName = "pbc-stats"
    ovfStatus.EntityData.SegmentPath = "ovf-status"
    ovfStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/pbc-statistics/pbc-stats/" + ovfStatus.EntityData.SegmentPath
    ovfStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ovfStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ovfStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ovfStatus.EntityData.Children = types.NewOrderedMap()
    ovfStatus.EntityData.Leafs = types.NewOrderedMap()
    ovfStatus.EntityData.Leafs.Append("cmic-cmc0-pkt-count-tx-pkt", types.YLeaf{"CmicCmc0PktCountTxPkt", ovfStatus.CmicCmc0PktCountTxPkt})
    ovfStatus.EntityData.Leafs.Append("cmic-cmc0-pkt-count-rx-pkt", types.YLeaf{"CmicCmc0PktCountRxPkt", ovfStatus.CmicCmc0PktCountRxPkt})
    ovfStatus.EntityData.Leafs.Append("nbi-stat-rx-bursts-err-cnt", types.YLeaf{"NbiStatRxBurstsErrCnt", ovfStatus.NbiStatRxBurstsErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-ecc-1b-err-cnt", types.YLeaf{"NbiEcc1bErrCnt", ovfStatus.NbiEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-ecc-2b-err-cnt", types.YLeaf{"NbiEcc2bErrCnt", ovfStatus.NbiEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-parity-err-cnt", types.YLeaf{"NbiParityErrCnt", ovfStatus.NbiParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn-crc32-err-cnt", types.YLeaf{"NbiRxIlknCrc32ErrCnt", ovfStatus.NbiRxIlknCrc32ErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-crc24-err-cnt", types.YLeaf{"NbiRxIlkn0Crc24ErrCnt", ovfStatus.NbiRxIlkn0Crc24ErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-burst-err-cnt", types.YLeaf{"NbiRxIlkn0BurstErrCnt", ovfStatus.NbiRxIlkn0BurstErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-miss-sop-err-cnt", types.YLeaf{"NbiRxIlkn0MissSopErrCnt", ovfStatus.NbiRxIlkn0MissSopErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-miss-eop-err-cnt", types.YLeaf{"NbiRxIlkn0MissEopErrCnt", ovfStatus.NbiRxIlkn0MissEopErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-misaligned-cnt", types.YLeaf{"NbiRxIlkn0MisalignedCnt", ovfStatus.NbiRxIlkn0MisalignedCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-crc24-err-cnt", types.YLeaf{"NbiRxIlkn1Crc24ErrCnt", ovfStatus.NbiRxIlkn1Crc24ErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-burst-err-cnt", types.YLeaf{"NbiRxIlkn1BurstErrCnt", ovfStatus.NbiRxIlkn1BurstErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-miss-sop-err-cnt", types.YLeaf{"NbiRxIlkn1MissSopErrCnt", ovfStatus.NbiRxIlkn1MissSopErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-miss-eop-err-cnt", types.YLeaf{"NbiRxIlkn1MissEopErrCnt", ovfStatus.NbiRxIlkn1MissEopErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-misaligned-cnt", types.YLeaf{"NbiRxIlkn1MisalignedCnt", ovfStatus.NbiRxIlkn1MisalignedCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-tx-ilkn1-flushed-bursts-cnt", types.YLeaf{"NbiTxIlkn1FlushedBurstsCnt", ovfStatus.NbiTxIlkn1FlushedBurstsCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-crc24-err-cnt", types.YLeaf{"NbiRxIlkn0RetransCrc24ErrCnt", ovfStatus.NbiRxIlkn0RetransCrc24ErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-retry-err-cnt", types.YLeaf{"NbiRxIlkn0RetransRetryErrCnt", ovfStatus.NbiRxIlkn0RetransRetryErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-wdog-err-cnt", types.YLeaf{"NbiRxIlkn0RetransWdogErrCnt", ovfStatus.NbiRxIlkn0RetransWdogErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-wrap-after-disc-err-cnt", types.YLeaf{"NbiRxIlkn0RetransWrapAfterDiscErrCnt", ovfStatus.NbiRxIlkn0RetransWrapAfterDiscErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-wrap-b4-disc-err-cnt", types.YLeaf{"NbiRxIlkn0RetransWrapB4DiscErrCnt", ovfStatus.NbiRxIlkn0RetransWrapB4DiscErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn0-retrans-reached-timeout-err-cnt", types.YLeaf{"NbiRxIlkn0RetransReachedTimeoutErrCnt", ovfStatus.NbiRxIlkn0RetransReachedTimeoutErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-crc24-err-cnt", types.YLeaf{"NbiRxIlkn1RetransCrc24ErrCnt", ovfStatus.NbiRxIlkn1RetransCrc24ErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-retry-err-cnt", types.YLeaf{"NbiRxIlkn1RetransRetryErrCnt", ovfStatus.NbiRxIlkn1RetransRetryErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-wdog-err-cnt", types.YLeaf{"NbiRxIlkn1RetransWdogErrCnt", ovfStatus.NbiRxIlkn1RetransWdogErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-wrap-after-disc-err-cnt", types.YLeaf{"NbiRxIlkn1RetransWrapAfterDiscErrCnt", ovfStatus.NbiRxIlkn1RetransWrapAfterDiscErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-wrap-b4-disc-err-cnt", types.YLeaf{"NbiRxIlkn1RetransWrapB4DiscErrCnt", ovfStatus.NbiRxIlkn1RetransWrapB4DiscErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-ilkn1-retrans-reached-timeout-err-cnt", types.YLeaf{"NbiRxIlkn1RetransReachedTimeoutErrCnt", ovfStatus.NbiRxIlkn1RetransReachedTimeoutErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-stat-rx-frame-err-cnt", types.YLeaf{"NbiStatRxFrameErrCnt", ovfStatus.NbiStatRxFrameErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-stat-tx-frame-err-cnt", types.YLeaf{"NbiStatTxFrameErrCnt", ovfStatus.NbiStatTxFrameErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-elk-err-bursts-cnt", types.YLeaf{"NbiRxElkErrBurstsCnt", ovfStatus.NbiRxElkErrBurstsCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-num-thrown-eops", types.YLeaf{"NbiRxNumThrownEops", ovfStatus.NbiRxNumThrownEops})
    ovfStatus.EntityData.Leafs.Append("nbi-rx-num-runts", types.YLeaf{"NbiRxNumRunts", ovfStatus.NbiRxNumRunts})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-tx-crc-err-bursts-cnt", types.YLeaf{"NbiBistTxCrcErrBurstsCnt", ovfStatus.NbiBistTxCrcErrBurstsCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-rx-err-length-bursts-cnt", types.YLeaf{"NbiBistRxErrLengthBurstsCnt", ovfStatus.NbiBistRxErrLengthBurstsCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-rx-err-burst-index-cnt", types.YLeaf{"NbiBistRxErrBurstIndexCnt", ovfStatus.NbiBistRxErrBurstIndexCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-rx-err-bct-cnt", types.YLeaf{"NbiBistRxErrBctCnt", ovfStatus.NbiBistRxErrBctCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-rx-err-data-cnt", types.YLeaf{"NbiBistRxErrDataCnt", ovfStatus.NbiBistRxErrDataCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-rx-err-in-crc-err-cnt", types.YLeaf{"NbiBistRxErrInCrcErrCnt", ovfStatus.NbiBistRxErrInCrcErrCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-bist-rx-err-sob-cnt", types.YLeaf{"NbiBistRxErrSobCnt", ovfStatus.NbiBistRxErrSobCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-stat-tx-bursts-cnt", types.YLeaf{"NbiStatTxBurstsCnt", ovfStatus.NbiStatTxBurstsCnt})
    ovfStatus.EntityData.Leafs.Append("nbi-stat-tx-total-leng-cnt", types.YLeaf{"NbiStatTxTotalLengCnt", ovfStatus.NbiStatTxTotalLengCnt})
    ovfStatus.EntityData.Leafs.Append("rxaui-total-tx-pkt-count", types.YLeaf{"RxauiTotalTxPktCount", ovfStatus.RxauiTotalTxPktCount})
    ovfStatus.EntityData.Leafs.Append("rxaui-total-rx-pkt-count", types.YLeaf{"RxauiTotalRxPktCount", ovfStatus.RxauiTotalRxPktCount})
    ovfStatus.EntityData.Leafs.Append("rxaui-rx-pkt-count-bcast-pkt", types.YLeaf{"RxauiRxPktCountBcastPkt", ovfStatus.RxauiRxPktCountBcastPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-tx-pkt-count-bcast-pkt", types.YLeaf{"RxauiTxPktCountBcastPkt", ovfStatus.RxauiTxPktCountBcastPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-rx-pkt-count-mcast-pkt", types.YLeaf{"RxauiRxPktCountMcastPkt", ovfStatus.RxauiRxPktCountMcastPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-tx-pkt-count-mcast-pkt", types.YLeaf{"RxauiTxPktCountMcastPkt", ovfStatus.RxauiTxPktCountMcastPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-rx-pkt-count-ucast-pkt", types.YLeaf{"RxauiRxPktCountUcastPkt", ovfStatus.RxauiRxPktCountUcastPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-tx-pkt-count-ucast-pkt", types.YLeaf{"RxauiTxPktCountUcastPkt", ovfStatus.RxauiTxPktCountUcastPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-rx-err-drop-pkt-cnt", types.YLeaf{"RxauiRxErrDropPktCnt", ovfStatus.RxauiRxErrDropPktCnt})
    ovfStatus.EntityData.Leafs.Append("rxaui-tx-err-drop-pkt-cnt", types.YLeaf{"RxauiTxErrDropPktCnt", ovfStatus.RxauiTxErrDropPktCnt})
    ovfStatus.EntityData.Leafs.Append("rxaui-byte-count-tx-pkt", types.YLeaf{"RxauiByteCountTxPkt", ovfStatus.RxauiByteCountTxPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-byte-count-rx-pkt", types.YLeaf{"RxauiByteCountRxPkt", ovfStatus.RxauiByteCountRxPkt})
    ovfStatus.EntityData.Leafs.Append("rxaui-rx-dscrd-pkt-cnt", types.YLeaf{"RxauiRxDscrdPktCnt", ovfStatus.RxauiRxDscrdPktCnt})
    ovfStatus.EntityData.Leafs.Append("rxaui-tx-dscrd-pkt-cnt", types.YLeaf{"RxauiTxDscrdPktCnt", ovfStatus.RxauiTxDscrdPktCnt})
    ovfStatus.EntityData.Leafs.Append("ire-nif-packet-counter", types.YLeaf{"IreNifPacketCounter", ovfStatus.IreNifPacketCounter})
    ovfStatus.EntityData.Leafs.Append("il-total-rx-pkt-count", types.YLeaf{"IlTotalRxPktCount", ovfStatus.IlTotalRxPktCount})
    ovfStatus.EntityData.Leafs.Append("il-total-tx-pkt-count", types.YLeaf{"IlTotalTxPktCount", ovfStatus.IlTotalTxPktCount})
    ovfStatus.EntityData.Leafs.Append("il-rx-err-drop-pkt-cnt", types.YLeaf{"IlRxErrDropPktCnt", ovfStatus.IlRxErrDropPktCnt})
    ovfStatus.EntityData.Leafs.Append("il-tx-err-drop-pkt-cnt", types.YLeaf{"IlTxErrDropPktCnt", ovfStatus.IlTxErrDropPktCnt})
    ovfStatus.EntityData.Leafs.Append("il-byte-count-tx-pkt", types.YLeaf{"IlByteCountTxPkt", ovfStatus.IlByteCountTxPkt})
    ovfStatus.EntityData.Leafs.Append("il-byte-count-rx-pkt", types.YLeaf{"IlByteCountRxPkt", ovfStatus.IlByteCountRxPkt})
    ovfStatus.EntityData.Leafs.Append("il-rx-dscrd-pkt-cnt", types.YLeaf{"IlRxDscrdPktCnt", ovfStatus.IlRxDscrdPktCnt})
    ovfStatus.EntityData.Leafs.Append("il-tx-dscrd-pkt-cnt", types.YLeaf{"IlTxDscrdPktCnt", ovfStatus.IlTxDscrdPktCnt})
    ovfStatus.EntityData.Leafs.Append("il-rx-pkt-count-bcast-pkt", types.YLeaf{"IlRxPktCountBcastPkt", ovfStatus.IlRxPktCountBcastPkt})
    ovfStatus.EntityData.Leafs.Append("il-tx-pkt-count-bcast-pkt", types.YLeaf{"IlTxPktCountBcastPkt", ovfStatus.IlTxPktCountBcastPkt})
    ovfStatus.EntityData.Leafs.Append("il-rx-pkt-count-mcast-pkt", types.YLeaf{"IlRxPktCountMcastPkt", ovfStatus.IlRxPktCountMcastPkt})
    ovfStatus.EntityData.Leafs.Append("il-tx-pkt-count-mcast-pkt", types.YLeaf{"IlTxPktCountMcastPkt", ovfStatus.IlTxPktCountMcastPkt})
    ovfStatus.EntityData.Leafs.Append("il-rx-pkt-count-ucast-pkt", types.YLeaf{"IlRxPktCountUcastPkt", ovfStatus.IlRxPktCountUcastPkt})
    ovfStatus.EntityData.Leafs.Append("il-tx-pkt-count-ucast-pkt", types.YLeaf{"IlTxPktCountUcastPkt", ovfStatus.IlTxPktCountUcastPkt})
    ovfStatus.EntityData.Leafs.Append("iqm-enq-pkt-cnt", types.YLeaf{"IqmEnqPktCnt", ovfStatus.IqmEnqPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-enq-byte-cnt", types.YLeaf{"IqmEnqByteCnt", ovfStatus.IqmEnqByteCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-deq-pkt-cnt", types.YLeaf{"IqmDeqPktCnt", ovfStatus.IqmDeqPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-deq-byte-cnt", types.YLeaf{"IqmDeqByteCnt", ovfStatus.IqmDeqByteCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-tot-dscrd-pkt-cnt", types.YLeaf{"IqmTotDscrdPktCnt", ovfStatus.IqmTotDscrdPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-tot-dscrd-byte-cnt", types.YLeaf{"IqmTotDscrdByteCnt", ovfStatus.IqmTotDscrdByteCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-ecc-1b-err-cnt", types.YLeaf{"IqmEcc1bErrCnt", ovfStatus.IqmEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-ecc-2b-err-cnt", types.YLeaf{"IqmEcc2bErrCnt", ovfStatus.IqmEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-parity-err-cnt", types.YLeaf{"IqmParityErrCnt", ovfStatus.IqmParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-deq-delete-pkt-cnt", types.YLeaf{"IqmDeqDeletePktCnt", ovfStatus.IqmDeqDeletePktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-ecn-dscrd-msk-pkt-cnt", types.YLeaf{"IqmEcnDscrdMskPktCnt", ovfStatus.IqmEcnDscrdMskPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-q-tot-dscrd-pkt-cnt", types.YLeaf{"IqmQTotDscrdPktCnt", ovfStatus.IqmQTotDscrdPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-q-deq-delete-pkt-cnt", types.YLeaf{"IqmQDeqDeletePktCnt", ovfStatus.IqmQDeqDeletePktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-db-pkt-cnt", types.YLeaf{"IqmRjctDbPktCnt", ovfStatus.IqmRjctDbPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-bdb-pkt-cnt", types.YLeaf{"IqmRjctBdbPktCnt", ovfStatus.IqmRjctBdbPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-bdb-protct-pkt-cnt", types.YLeaf{"IqmRjctBdbProtctPktCnt", ovfStatus.IqmRjctBdbProtctPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-oc-bd-pkt-cnt", types.YLeaf{"IqmRjctOcBdPktCnt", ovfStatus.IqmRjctOcBdPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-sn-err-pkt-cnt", types.YLeaf{"IqmRjctSnErrPktCnt", ovfStatus.IqmRjctSnErrPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-mc-err-pkt-cnt", types.YLeaf{"IqmRjctMcErrPktCnt", ovfStatus.IqmRjctMcErrPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-rsrc-err-pkt-cnt", types.YLeaf{"IqmRjctRsrcErrPktCnt", ovfStatus.IqmRjctRsrcErrPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-qnvalid-err-pkt-cnt", types.YLeaf{"IqmRjctQnvalidErrPktCnt", ovfStatus.IqmRjctQnvalidErrPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-cnm-pkt-cnt", types.YLeaf{"IqmRjctCnmPktCnt", ovfStatus.IqmRjctCnmPktCnt})
    ovfStatus.EntityData.Leafs.Append("iqm-rjct-dyn-space-pkt-cnt", types.YLeaf{"IqmRjctDynSpacePktCnt", ovfStatus.IqmRjctDynSpacePktCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-fdt-pkt-cnt", types.YLeaf{"IptFdtPktCnt", ovfStatus.IptFdtPktCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-ecc-1b-err-cnt", types.YLeaf{"IptEcc1bErrCnt", ovfStatus.IptEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-ecc-2b-err-cnt", types.YLeaf{"IptEcc2bErrCnt", ovfStatus.IptEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-parity-err-cnt", types.YLeaf{"IptParityErrCnt", ovfStatus.IptParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-crc-err-cnt", types.YLeaf{"IptCrcErrCnt", ovfStatus.IptCrcErrCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-crc-err-del-buff-cnt", types.YLeaf{"IptCrcErrDelBuffCnt", ovfStatus.IptCrcErrDelBuffCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-cpu-del-buff-cnt", types.YLeaf{"IptCpuDelBuffCnt", ovfStatus.IptCpuDelBuffCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-cpu-rel-buff-cnt", types.YLeaf{"IptCpuRelBuffCnt", ovfStatus.IptCpuRelBuffCnt})
    ovfStatus.EntityData.Leafs.Append("ipt-crc-err-buff-fifo-full-cnt", types.YLeaf{"IptCrcErrBuffFifoFullCnt", ovfStatus.IptCrcErrBuffFifoFullCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-data-cell-cnt", types.YLeaf{"FdtDataCellCnt", ovfStatus.FdtDataCellCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-data-byte-cnt", types.YLeaf{"FdtDataByteCnt", ovfStatus.FdtDataByteCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-crc-dropped-pck-cnt", types.YLeaf{"FdtCrcDroppedPckCnt", ovfStatus.FdtCrcDroppedPckCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-invalid-destq-drop-cell-cnt", types.YLeaf{"FdtInvalidDestqDropCellCnt", ovfStatus.FdtInvalidDestqDropCellCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-indirect-command-count", types.YLeaf{"FdtIndirectCommandCount", ovfStatus.FdtIndirectCommandCount})
    ovfStatus.EntityData.Leafs.Append("fdt-ecc-1b-err-cnt", types.YLeaf{"FdtEcc1bErrCnt", ovfStatus.FdtEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-ecc-2b-err-cnt", types.YLeaf{"FdtEcc2bErrCnt", ovfStatus.FdtEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-parity-err-cnt", types.YLeaf{"FdtParityErrCnt", ovfStatus.FdtParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdt-crc-dropped-cell-cnt", types.YLeaf{"FdtCrcDroppedCellCnt", ovfStatus.FdtCrcDroppedCellCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-control-cell-cnt", types.YLeaf{"FcrControlCellCnt", ovfStatus.FcrControlCellCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-cell-drop-cnt", types.YLeaf{"FcrCellDropCnt", ovfStatus.FcrCellDropCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-credit-cell-drop-cnt", types.YLeaf{"FcrCreditCellDropCnt", ovfStatus.FcrCreditCellDropCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-fs-cell-drop-cnt", types.YLeaf{"FcrFsCellDropCnt", ovfStatus.FcrFsCellDropCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-rt-cell-drop-cnt", types.YLeaf{"FcrRtCellDropCnt", ovfStatus.FcrRtCellDropCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-ecc-1b-err-cnt", types.YLeaf{"FcrEcc1bErrCnt", ovfStatus.FcrEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fcr-ecc-2b-err-cnt", types.YLeaf{"FcrEcc2bErrCnt", ovfStatus.FcrEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-data-cell-cnt", types.YLeaf{"FdrDataCellCnt", ovfStatus.FdrDataCellCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-data-byte-cnt", types.YLeaf{"FdrDataByteCnt", ovfStatus.FdrDataByteCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-crc-dropped-pck-cnt", types.YLeaf{"FdrCrcDroppedPckCnt", ovfStatus.FdrCrcDroppedPckCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-p-pkt-cnt", types.YLeaf{"FdrPPktCnt", ovfStatus.FdrPPktCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-prm-error-filter-cnt", types.YLeaf{"FdrPrmErrorFilterCnt", ovfStatus.FdrPrmErrorFilterCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-sec-error-filter-cnt", types.YLeaf{"FdrSecErrorFilterCnt", ovfStatus.FdrSecErrorFilterCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-prm-ecc-1b-err-cnt", types.YLeaf{"FdrPrmEcc1bErrCnt", ovfStatus.FdrPrmEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-prm-ecc-2b-err-cnt", types.YLeaf{"FdrPrmEcc2bErrCnt", ovfStatus.FdrPrmEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-sec-ecc-1b-err-cnt", types.YLeaf{"FdrSecEcc1bErrCnt", ovfStatus.FdrSecEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fdr-sec-ecc-2b-err-cnt", types.YLeaf{"FdrSecEcc2bErrCnt", ovfStatus.FdrSecEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-ecc-1b-err-cnt", types.YLeaf{"EgqEcc1bErrCnt", ovfStatus.EgqEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-ecc-2b-err-cnt", types.YLeaf{"EgqEcc2bErrCnt", ovfStatus.EgqEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-parity-err-cnt", types.YLeaf{"EgqParityErrCnt", ovfStatus.EgqParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-dbf-ecc-1b-err-cnt", types.YLeaf{"EgqDbfEcc1bErrCnt", ovfStatus.EgqDbfEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-dbf-ecc-2b-err-cnt", types.YLeaf{"EgqDbfEcc2bErrCnt", ovfStatus.EgqDbfEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-empty-mcid-counter", types.YLeaf{"EgqEmptyMcidCounter", ovfStatus.EgqEmptyMcidCounter})
    ovfStatus.EntityData.Leafs.Append("egq-rqp-discard-packet-counter", types.YLeaf{"EgqRqpDiscardPacketCounter", ovfStatus.EgqRqpDiscardPacketCounter})
    ovfStatus.EntityData.Leafs.Append("egq-ehp-discard-packet-counter", types.YLeaf{"EgqEhpDiscardPacketCounter", ovfStatus.EgqEhpDiscardPacketCounter})
    ovfStatus.EntityData.Leafs.Append("egq-ipt-pkt-cnt", types.YLeaf{"EgqIptPktCnt", ovfStatus.EgqIptPktCnt})
    ovfStatus.EntityData.Leafs.Append("epni-epe-pkt-cnt", types.YLeaf{"EpniEpePktCnt", ovfStatus.EpniEpePktCnt})
    ovfStatus.EntityData.Leafs.Append("epni-epe-byte-cnt", types.YLeaf{"EpniEpeByteCnt", ovfStatus.EpniEpeByteCnt})
    ovfStatus.EntityData.Leafs.Append("epni-epe-discard-pkt-cnt", types.YLeaf{"EpniEpeDiscardPktCnt", ovfStatus.EpniEpeDiscardPktCnt})
    ovfStatus.EntityData.Leafs.Append("epni-ecc-1b-err-cnt", types.YLeaf{"EpniEcc1bErrCnt", ovfStatus.EpniEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("epni-ecc-2b-err-cnt", types.YLeaf{"EpniEcc2bErrCnt", ovfStatus.EpniEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("epni-parity-err-cnt", types.YLeaf{"EpniParityErrCnt", ovfStatus.EpniParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-ucast-pkt-cnt", types.YLeaf{"EgqPqpUcastPktCnt", ovfStatus.EgqPqpUcastPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-ucast-h-pkt-cnt", types.YLeaf{"EgqPqpUcastHPktCnt", ovfStatus.EgqPqpUcastHPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-ucast-l-pkt-cnt", types.YLeaf{"EgqPqpUcastLPktCnt", ovfStatus.EgqPqpUcastLPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-ucast-bytes-cnt", types.YLeaf{"EgqPqpUcastBytesCnt", ovfStatus.EgqPqpUcastBytesCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-ucast-discard-pkt-cnt", types.YLeaf{"EgqPqpUcastDiscardPktCnt", ovfStatus.EgqPqpUcastDiscardPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-mcast-pkt-cnt", types.YLeaf{"EgqPqpMcastPktCnt", ovfStatus.EgqPqpMcastPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-mcast-h-pkt-cnt", types.YLeaf{"EgqPqpMcastHPktCnt", ovfStatus.EgqPqpMcastHPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-mcast-l-pkt-cnt", types.YLeaf{"EgqPqpMcastLPktCnt", ovfStatus.EgqPqpMcastLPktCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-mcast-bytes-cnt", types.YLeaf{"EgqPqpMcastBytesCnt", ovfStatus.EgqPqpMcastBytesCnt})
    ovfStatus.EntityData.Leafs.Append("egq-pqp-mcast-discard-pkt-cnt", types.YLeaf{"EgqPqpMcastDiscardPktCnt", ovfStatus.EgqPqpMcastDiscardPktCnt})
    ovfStatus.EntityData.Leafs.Append("fct-control-cell-cnt", types.YLeaf{"FctControlCellCnt", ovfStatus.FctControlCellCnt})
    ovfStatus.EntityData.Leafs.Append("fct-unrch-crdt-cnt", types.YLeaf{"FctUnrchCrdtCnt", ovfStatus.FctUnrchCrdtCnt})
    ovfStatus.EntityData.Leafs.Append("idr-reassembly-errors", types.YLeaf{"IdrReassemblyErrors", ovfStatus.IdrReassemblyErrors})
    ovfStatus.EntityData.Leafs.Append("idr-mmu-ecc-1b-err-cnt", types.YLeaf{"IdrMmuEcc1bErrCnt", ovfStatus.IdrMmuEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("idr-mmu-ecc-2b-err-cnt", types.YLeaf{"IdrMmuEcc2bErrCnt", ovfStatus.IdrMmuEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-packets0-cnt", types.YLeaf{"IdrDiscardedPackets0Cnt", ovfStatus.IdrDiscardedPackets0Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-packets1-cnt", types.YLeaf{"IdrDiscardedPackets1Cnt", ovfStatus.IdrDiscardedPackets1Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-packets2-cnt", types.YLeaf{"IdrDiscardedPackets2Cnt", ovfStatus.IdrDiscardedPackets2Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-packets3-cnt", types.YLeaf{"IdrDiscardedPackets3Cnt", ovfStatus.IdrDiscardedPackets3Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-octets0-cnt", types.YLeaf{"IdrDiscardedOctets0Cnt", ovfStatus.IdrDiscardedOctets0Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-octets1-cnt", types.YLeaf{"IdrDiscardedOctets1Cnt", ovfStatus.IdrDiscardedOctets1Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-octets2-cnt", types.YLeaf{"IdrDiscardedOctets2Cnt", ovfStatus.IdrDiscardedOctets2Cnt})
    ovfStatus.EntityData.Leafs.Append("idr-discarded-octets3-cnt", types.YLeaf{"IdrDiscardedOctets3Cnt", ovfStatus.IdrDiscardedOctets3Cnt})
    ovfStatus.EntityData.Leafs.Append("mmu-ecc-1b-err-cnt", types.YLeaf{"MmuEcc1bErrCnt", ovfStatus.MmuEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("mmu-ecc-2b-err-cnt", types.YLeaf{"MmuEcc2bErrCnt", ovfStatus.MmuEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("oamp-parity-err-cnt", types.YLeaf{"OampParityErrCnt", ovfStatus.OampParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("oamp-ecc-1b-err-cnt", types.YLeaf{"OampEcc1bErrCnt", ovfStatus.OampEcc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("oamp-ecc-2b-err-cnt", types.YLeaf{"OampEcc2bErrCnt", ovfStatus.OampEcc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("crps-parity-err-cnt", types.YLeaf{"CrpsParityErrCnt", ovfStatus.CrpsParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs0TstRxErrCnt", ovfStatus.Fmac0Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs0TstRxErrCnt", ovfStatus.Fmac1Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs0TstRxErrCnt", ovfStatus.Fmac2Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs0TstRxErrCnt", ovfStatus.Fmac3Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs0TstRxErrCnt", ovfStatus.Fmac4Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs0TstRxErrCnt", ovfStatus.Fmac5Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs0TstRxErrCnt", ovfStatus.Fmac6Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs0TstRxErrCnt", ovfStatus.Fmac7Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-kpcs0-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs0TstRxErrCnt", ovfStatus.Fmac8Kpcs0TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs1TstRxErrCnt", ovfStatus.Fmac0Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs1TstRxErrCnt", ovfStatus.Fmac1Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs1TstRxErrCnt", ovfStatus.Fmac2Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs1TstRxErrCnt", ovfStatus.Fmac3Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs1TstRxErrCnt", ovfStatus.Fmac4Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs1TstRxErrCnt", ovfStatus.Fmac5Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs1TstRxErrCnt", ovfStatus.Fmac6Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs1TstRxErrCnt", ovfStatus.Fmac7Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-kpcs1-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs1TstRxErrCnt", ovfStatus.Fmac8Kpcs1TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs2TstRxErrCnt", ovfStatus.Fmac0Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs2TstRxErrCnt", ovfStatus.Fmac1Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs2TstRxErrCnt", ovfStatus.Fmac2Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs2TstRxErrCnt", ovfStatus.Fmac3Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs2TstRxErrCnt", ovfStatus.Fmac4Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs2TstRxErrCnt", ovfStatus.Fmac5Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs2TstRxErrCnt", ovfStatus.Fmac6Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs2TstRxErrCnt", ovfStatus.Fmac7Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-kpcs2-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs2TstRxErrCnt", ovfStatus.Fmac8Kpcs2TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac0Kpcs3TstRxErrCnt", ovfStatus.Fmac0Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac1Kpcs3TstRxErrCnt", ovfStatus.Fmac1Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac2Kpcs3TstRxErrCnt", ovfStatus.Fmac2Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac3Kpcs3TstRxErrCnt", ovfStatus.Fmac3Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac4Kpcs3TstRxErrCnt", ovfStatus.Fmac4Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac5Kpcs3TstRxErrCnt", ovfStatus.Fmac5Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac6Kpcs3TstRxErrCnt", ovfStatus.Fmac6Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac7Kpcs3TstRxErrCnt", ovfStatus.Fmac7Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-kpcs3-tst-rx-err-cnt", types.YLeaf{"Fmac8Kpcs3TstRxErrCnt", ovfStatus.Fmac8Kpcs3TstRxErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-tst0-err-cnt", types.YLeaf{"Fmac0Tst0ErrCnt", ovfStatus.Fmac0Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-tst0-err-cnt", types.YLeaf{"Fmac1Tst0ErrCnt", ovfStatus.Fmac1Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-tst0-err-cnt", types.YLeaf{"Fmac2Tst0ErrCnt", ovfStatus.Fmac2Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-tst0-err-cnt", types.YLeaf{"Fmac3Tst0ErrCnt", ovfStatus.Fmac3Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-tst0-err-cnt", types.YLeaf{"Fmac4Tst0ErrCnt", ovfStatus.Fmac4Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-tst0-err-cnt", types.YLeaf{"Fmac5Tst0ErrCnt", ovfStatus.Fmac5Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-tst0-err-cnt", types.YLeaf{"Fmac6Tst0ErrCnt", ovfStatus.Fmac6Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-tst0-err-cnt", types.YLeaf{"Fmac7Tst0ErrCnt", ovfStatus.Fmac7Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-tst0-err-cnt", types.YLeaf{"Fmac8Tst0ErrCnt", ovfStatus.Fmac8Tst0ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-tst1-err-cnt", types.YLeaf{"Fmac0Tst1ErrCnt", ovfStatus.Fmac0Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-tst1-err-cnt", types.YLeaf{"Fmac1Tst1ErrCnt", ovfStatus.Fmac1Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-tst1-err-cnt", types.YLeaf{"Fmac2Tst1ErrCnt", ovfStatus.Fmac2Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-tst1-err-cnt", types.YLeaf{"Fmac3Tst1ErrCnt", ovfStatus.Fmac3Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-tst1-err-cnt", types.YLeaf{"Fmac4Tst1ErrCnt", ovfStatus.Fmac4Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-tst1-err-cnt", types.YLeaf{"Fmac5Tst1ErrCnt", ovfStatus.Fmac5Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-tst1-err-cnt", types.YLeaf{"Fmac6Tst1ErrCnt", ovfStatus.Fmac6Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-tst1-err-cnt", types.YLeaf{"Fmac7Tst1ErrCnt", ovfStatus.Fmac7Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-tst1-err-cnt", types.YLeaf{"Fmac8Tst1ErrCnt", ovfStatus.Fmac8Tst1ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-tst2-err-cnt", types.YLeaf{"Fmac0Tst2ErrCnt", ovfStatus.Fmac0Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-tst2-err-cnt", types.YLeaf{"Fmac1Tst2ErrCnt", ovfStatus.Fmac1Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-tst2-err-cnt", types.YLeaf{"Fmac2Tst2ErrCnt", ovfStatus.Fmac2Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-tst2-err-cnt", types.YLeaf{"Fmac3Tst2ErrCnt", ovfStatus.Fmac3Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-tst2-err-cnt", types.YLeaf{"Fmac4Tst2ErrCnt", ovfStatus.Fmac4Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-tst2-err-cnt", types.YLeaf{"Fmac5Tst2ErrCnt", ovfStatus.Fmac5Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-tst2-err-cnt", types.YLeaf{"Fmac6Tst2ErrCnt", ovfStatus.Fmac6Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-tst2-err-cnt", types.YLeaf{"Fmac7Tst2ErrCnt", ovfStatus.Fmac7Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-tst2-err-cnt", types.YLeaf{"Fmac8Tst2ErrCnt", ovfStatus.Fmac8Tst2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-tst3-err-cnt", types.YLeaf{"Fmac0Tst3ErrCnt", ovfStatus.Fmac0Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-tst3-err-cnt", types.YLeaf{"Fmac1Tst3ErrCnt", ovfStatus.Fmac1Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-tst3-err-cnt", types.YLeaf{"Fmac2Tst3ErrCnt", ovfStatus.Fmac2Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-tst3-err-cnt", types.YLeaf{"Fmac3Tst3ErrCnt", ovfStatus.Fmac3Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-tst3-err-cnt", types.YLeaf{"Fmac4Tst3ErrCnt", ovfStatus.Fmac4Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-tst3-err-cnt", types.YLeaf{"Fmac5Tst3ErrCnt", ovfStatus.Fmac5Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-tst3-err-cnt", types.YLeaf{"Fmac6Tst3ErrCnt", ovfStatus.Fmac6Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-tst3-err-cnt", types.YLeaf{"Fmac7Tst3ErrCnt", ovfStatus.Fmac7Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-tst3-err-cnt", types.YLeaf{"Fmac8Tst3ErrCnt", ovfStatus.Fmac8Tst3ErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-ecc-1b-err-cnt", types.YLeaf{"Fmac0Ecc1bErrCnt", ovfStatus.Fmac0Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-ecc-1b-err-cnt", types.YLeaf{"Fmac1Ecc1bErrCnt", ovfStatus.Fmac1Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-ecc-1b-err-cnt", types.YLeaf{"Fmac2Ecc1bErrCnt", ovfStatus.Fmac2Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-ecc-1b-err-cnt", types.YLeaf{"Fmac3Ecc1bErrCnt", ovfStatus.Fmac3Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-ecc-1b-err-cnt", types.YLeaf{"Fmac4Ecc1bErrCnt", ovfStatus.Fmac4Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-ecc-1b-err-cnt", types.YLeaf{"Fmac5Ecc1bErrCnt", ovfStatus.Fmac5Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-ecc-1b-err-cnt", types.YLeaf{"Fmac6Ecc1bErrCnt", ovfStatus.Fmac6Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-ecc-1b-err-cnt", types.YLeaf{"Fmac7Ecc1bErrCnt", ovfStatus.Fmac7Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-ecc-1b-err-cnt", types.YLeaf{"Fmac8Ecc1bErrCnt", ovfStatus.Fmac8Ecc1bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac0-ecc-2b-err-cnt", types.YLeaf{"Fmac0Ecc2bErrCnt", ovfStatus.Fmac0Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac1-ecc-2b-err-cnt", types.YLeaf{"Fmac1Ecc2bErrCnt", ovfStatus.Fmac1Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac2-ecc-2b-err-cnt", types.YLeaf{"Fmac2Ecc2bErrCnt", ovfStatus.Fmac2Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac3-ecc-2b-err-cnt", types.YLeaf{"Fmac3Ecc2bErrCnt", ovfStatus.Fmac3Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac4-ecc-2b-err-cnt", types.YLeaf{"Fmac4Ecc2bErrCnt", ovfStatus.Fmac4Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac5-ecc-2b-err-cnt", types.YLeaf{"Fmac5Ecc2bErrCnt", ovfStatus.Fmac5Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac6-ecc-2b-err-cnt", types.YLeaf{"Fmac6Ecc2bErrCnt", ovfStatus.Fmac6Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac7-ecc-2b-err-cnt", types.YLeaf{"Fmac7Ecc2bErrCnt", ovfStatus.Fmac7Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("fmac8-ecc-2b-err-cnt", types.YLeaf{"Fmac8Ecc2bErrCnt", ovfStatus.Fmac8Ecc2bErrCnt})
    ovfStatus.EntityData.Leafs.Append("olp-incoming-bad-identifier-counter", types.YLeaf{"OlpIncomingBadIdentifierCounter", ovfStatus.OlpIncomingBadIdentifierCounter})
    ovfStatus.EntityData.Leafs.Append("olp-incoming-bad-reassembly-counter", types.YLeaf{"OlpIncomingBadReassemblyCounter", ovfStatus.OlpIncomingBadReassemblyCounter})
    ovfStatus.EntityData.Leafs.Append("cfc-parity-err-cnt", types.YLeaf{"CfcParityErrCnt", ovfStatus.CfcParityErrCnt})
    ovfStatus.EntityData.Leafs.Append("cfc-ilkn0-oob-rx-crc-err-cntr", types.YLeaf{"CfcIlkn0OobRxCrcErrCntr", ovfStatus.CfcIlkn0OobRxCrcErrCntr})
    ovfStatus.EntityData.Leafs.Append("cfc-ilkn1-oob-rx-crc-err-cntr", types.YLeaf{"CfcIlkn1OobRxCrcErrCntr", ovfStatus.CfcIlkn1OobRxCrcErrCntr})
    ovfStatus.EntityData.Leafs.Append("cfc-spi-oob-rx0-frm-err-cnt", types.YLeaf{"CfcSpiOobRx0FrmErrCnt", ovfStatus.CfcSpiOobRx0FrmErrCnt})
    ovfStatus.EntityData.Leafs.Append("cfc-spi-oob-rx0-dip2-err-cnt", types.YLeaf{"CfcSpiOobRx0Dip2ErrCnt", ovfStatus.CfcSpiOobRx0Dip2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("cfc-spi-oob-rx1-frm-err-cnt", types.YLeaf{"CfcSpiOobRx1FrmErrCnt", ovfStatus.CfcSpiOobRx1FrmErrCnt})
    ovfStatus.EntityData.Leafs.Append("cfc-spi-oob-rx1-dip2-err-cnt", types.YLeaf{"CfcSpiOobRx1Dip2ErrCnt", ovfStatus.CfcSpiOobRx1Dip2ErrCnt})
    ovfStatus.EntityData.Leafs.Append("cgm-cgm-uc-pd-dropped-cnt", types.YLeaf{"CgmCgmUcPdDroppedCnt", ovfStatus.CgmCgmUcPdDroppedCnt})
    ovfStatus.EntityData.Leafs.Append("cgm-cgm-mc-rep-pd-dropped-cnt", types.YLeaf{"CgmCgmMcRepPdDroppedCnt", ovfStatus.CgmCgmMcRepPdDroppedCnt})
    ovfStatus.EntityData.Leafs.Append("cgm-cgm-uc-db-dropped-by-rqp-cnt", types.YLeaf{"CgmCgmUcDbDroppedByRqpCnt", ovfStatus.CgmCgmUcDbDroppedByRqpCnt})
    ovfStatus.EntityData.Leafs.Append("cgm-cgm-uc-db-dropped-by-pqp-cnt", types.YLeaf{"CgmCgmUcDbDroppedByPqpCnt", ovfStatus.CgmCgmUcDbDroppedByPqpCnt})
    ovfStatus.EntityData.Leafs.Append("cgm-cgm-mc-rep-db-dropped-cnt", types.YLeaf{"CgmCgmMcRepDbDroppedCnt", ovfStatus.CgmCgmMcRepDbDroppedCnt})
    ovfStatus.EntityData.Leafs.Append("cgm-cgm-mc-db-dropped-cnt", types.YLeaf{"CgmCgmMcDbDroppedCnt", ovfStatus.CgmCgmMcDbDroppedCnt})
    ovfStatus.EntityData.Leafs.Append("drca-full-err-cnt", types.YLeaf{"DrcaFullErrCnt", ovfStatus.DrcaFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drca-single-err-cnt", types.YLeaf{"DrcaSingleErrCnt", ovfStatus.DrcaSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drca-calib-bist-full-err-cnt", types.YLeaf{"DrcaCalibBistFullErrCnt", ovfStatus.DrcaCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drca-loopback-full-err-cnt", types.YLeaf{"DrcaLoopbackFullErrCnt", ovfStatus.DrcaLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcb-full-err-cnt", types.YLeaf{"DrcbFullErrCnt", ovfStatus.DrcbFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcb-single-err-cnt", types.YLeaf{"DrcbSingleErrCnt", ovfStatus.DrcbSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcb-calib-bist-full-err-cnt", types.YLeaf{"DrcbCalibBistFullErrCnt", ovfStatus.DrcbCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcb-loopback-full-err-cnt", types.YLeaf{"DrcbLoopbackFullErrCnt", ovfStatus.DrcbLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcc-full-err-cnt", types.YLeaf{"DrccFullErrCnt", ovfStatus.DrccFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcc-single-err-cnt", types.YLeaf{"DrccSingleErrCnt", ovfStatus.DrccSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcc-calib-bist-full-err-cnt", types.YLeaf{"DrccCalibBistFullErrCnt", ovfStatus.DrccCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcc-loopback-full-err-cnt", types.YLeaf{"DrccLoopbackFullErrCnt", ovfStatus.DrccLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcd-full-err-cnt", types.YLeaf{"DrcdFullErrCnt", ovfStatus.DrcdFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcd-single-err-cnt", types.YLeaf{"DrcdSingleErrCnt", ovfStatus.DrcdSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcd-calib-bist-full-err-cnt", types.YLeaf{"DrcdCalibBistFullErrCnt", ovfStatus.DrcdCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcd-loopback-full-err-cnt", types.YLeaf{"DrcdLoopbackFullErrCnt", ovfStatus.DrcdLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drce-full-err-cnt", types.YLeaf{"DrceFullErrCnt", ovfStatus.DrceFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drce-single-err-cnt", types.YLeaf{"DrceSingleErrCnt", ovfStatus.DrceSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drce-calib-bist-full-err-cnt", types.YLeaf{"DrceCalibBistFullErrCnt", ovfStatus.DrceCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drce-loopback-full-err-cnt", types.YLeaf{"DrceLoopbackFullErrCnt", ovfStatus.DrceLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcf-full-err-cnt", types.YLeaf{"DrcfFullErrCnt", ovfStatus.DrcfFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcf-single-err-cnt", types.YLeaf{"DrcfSingleErrCnt", ovfStatus.DrcfSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcf-calib-bist-full-err-cnt", types.YLeaf{"DrcfCalibBistFullErrCnt", ovfStatus.DrcfCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcf-loopback-full-err-cnt", types.YLeaf{"DrcfLoopbackFullErrCnt", ovfStatus.DrcfLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcg-full-err-cnt", types.YLeaf{"DrcgFullErrCnt", ovfStatus.DrcgFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcg-single-err-cnt", types.YLeaf{"DrcgSingleErrCnt", ovfStatus.DrcgSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcg-calib-bist-full-err-cnt", types.YLeaf{"DrcgCalibBistFullErrCnt", ovfStatus.DrcgCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcg-loopback-full-err-cnt", types.YLeaf{"DrcgLoopbackFullErrCnt", ovfStatus.DrcgLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drch-full-err-cnt", types.YLeaf{"DrchFullErrCnt", ovfStatus.DrchFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drch-single-err-cnt", types.YLeaf{"DrchSingleErrCnt", ovfStatus.DrchSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drch-calib-bist-full-err-cnt", types.YLeaf{"DrchCalibBistFullErrCnt", ovfStatus.DrchCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drch-loopback-full-err-cnt", types.YLeaf{"DrchLoopbackFullErrCnt", ovfStatus.DrchLoopbackFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcbroadcast-full-err-cnt", types.YLeaf{"DrcbroadcastFullErrCnt", ovfStatus.DrcbroadcastFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcbroadcast-single-err-cnt", types.YLeaf{"DrcbroadcastSingleErrCnt", ovfStatus.DrcbroadcastSingleErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcbroadcast-calib-bist-full-err-cnt", types.YLeaf{"DrcbroadcastCalibBistFullErrCnt", ovfStatus.DrcbroadcastCalibBistFullErrCnt})
    ovfStatus.EntityData.Leafs.Append("drcbroadcast-loopback-full-err-cnt", types.YLeaf{"DrcbroadcastLoopbackFullErrCnt", ovfStatus.DrcbroadcastLoopbackFullErrCnt})

    ovfStatus.EntityData.YListKeys = []string {}

    return &(ovfStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics
// Statistics of FMAC
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link table for statistics.
    FmacLinks Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks
}

func (fmacStatistics *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics) GetEntityData() *types.CommonEntityData {
    fmacStatistics.EntityData.YFilter = fmacStatistics.YFilter
    fmacStatistics.EntityData.YangName = "fmac-statistics"
    fmacStatistics.EntityData.BundleName = "cisco_ios_xr"
    fmacStatistics.EntityData.ParentYangName = "statistics-asic-instance"
    fmacStatistics.EntityData.SegmentPath = "fmac-statistics"
    fmacStatistics.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/" + fmacStatistics.EntityData.SegmentPath
    fmacStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacStatistics.EntityData.Children = types.NewOrderedMap()
    fmacStatistics.EntityData.Children.Append("fmac-links", types.YChild{"FmacLinks", &fmacStatistics.FmacLinks})
    fmacStatistics.EntityData.Leafs = types.NewOrderedMap()

    fmacStatistics.EntityData.YListKeys = []string {}

    return &(fmacStatistics.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks
// Link table for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Link number for statistics. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink.
    FmacLink []*Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink
}

func (fmacLinks *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks) GetEntityData() *types.CommonEntityData {
    fmacLinks.EntityData.YFilter = fmacLinks.YFilter
    fmacLinks.EntityData.YangName = "fmac-links"
    fmacLinks.EntityData.BundleName = "cisco_ios_xr"
    fmacLinks.EntityData.ParentYangName = "fmac-statistics"
    fmacLinks.EntityData.SegmentPath = "fmac-links"
    fmacLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/" + fmacLinks.EntityData.SegmentPath
    fmacLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacLinks.EntityData.Children = types.NewOrderedMap()
    fmacLinks.EntityData.Children.Append("fmac-link", types.YChild{"FmacLink", nil})
    for i := range fmacLinks.FmacLink {
        fmacLinks.EntityData.Children.Append(types.GetSegmentPath(fmacLinks.FmacLink[i]), types.YChild{"FmacLink", fmacLinks.FmacLink[i]})
    }
    fmacLinks.EntityData.Leafs = types.NewOrderedMap()

    fmacLinks.EntityData.YListKeys = []string {}

    return &(fmacLinks.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink
// Link number for statistics
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Link number. The type is interface{} with range:
    // 0..4294967295.
    Link interface{}

    // Single aisc information. The type is slice of
    // Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic.
    FmacAsic []*Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic
}

func (fmacLink *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink) GetEntityData() *types.CommonEntityData {
    fmacLink.EntityData.YFilter = fmacLink.YFilter
    fmacLink.EntityData.YangName = "fmac-link"
    fmacLink.EntityData.BundleName = "cisco_ios_xr"
    fmacLink.EntityData.ParentYangName = "fmac-links"
    fmacLink.EntityData.SegmentPath = "fmac-link" + types.AddKeyToken(fmacLink.Link, "link")
    fmacLink.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/" + fmacLink.EntityData.SegmentPath
    fmacLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacLink.EntityData.Children = types.NewOrderedMap()
    fmacLink.EntityData.Children.Append("fmac-asic", types.YChild{"FmacAsic", nil})
    for i := range fmacLink.FmacAsic {
        fmacLink.EntityData.Children.Append(types.GetSegmentPath(fmacLink.FmacAsic[i]), types.YChild{"FmacAsic", fmacLink.FmacAsic[i]})
    }
    fmacLink.EntityData.Leafs = types.NewOrderedMap()
    fmacLink.EntityData.Leafs.Append("link", types.YLeaf{"Link", fmacLink.Link})

    fmacLink.EntityData.YListKeys = []string {"Link"}

    return &(fmacLink.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic
// Single aisc information
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Single asic. The type is interface{} with range:
    // 0..4294967295.
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

func (fmacAsic *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic) GetEntityData() *types.CommonEntityData {
    fmacAsic.EntityData.YFilter = fmacAsic.YFilter
    fmacAsic.EntityData.YangName = "fmac-asic"
    fmacAsic.EntityData.BundleName = "cisco_ios_xr"
    fmacAsic.EntityData.ParentYangName = "fmac-link"
    fmacAsic.EntityData.SegmentPath = "fmac-asic" + types.AddKeyToken(fmacAsic.Asic, "asic")
    fmacAsic.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/" + fmacAsic.EntityData.SegmentPath
    fmacAsic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fmacAsic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fmacAsic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fmacAsic.EntityData.Children = types.NewOrderedMap()
    fmacAsic.EntityData.Children.Append("aggr-stats", types.YChild{"AggrStats", &fmacAsic.AggrStats})
    fmacAsic.EntityData.Children.Append("incr-stats", types.YChild{"IncrStats", &fmacAsic.IncrStats})
    fmacAsic.EntityData.Leafs = types.NewOrderedMap()
    fmacAsic.EntityData.Leafs.Append("asic", types.YLeaf{"Asic", fmacAsic.Asic})
    fmacAsic.EntityData.Leafs.Append("valid", types.YLeaf{"Valid", fmacAsic.Valid})
    fmacAsic.EntityData.Leafs.Append("rack-no", types.YLeaf{"RackNo", fmacAsic.RackNo})
    fmacAsic.EntityData.Leafs.Append("slot-no", types.YLeaf{"SlotNo", fmacAsic.SlotNo})
    fmacAsic.EntityData.Leafs.Append("asic-instance", types.YLeaf{"AsicInstance", fmacAsic.AsicInstance})
    fmacAsic.EntityData.Leafs.Append("link-no", types.YLeaf{"LinkNo", fmacAsic.LinkNo})
    fmacAsic.EntityData.Leafs.Append("link-valid", types.YLeaf{"LinkValid", fmacAsic.LinkValid})

    fmacAsic.EntityData.YListKeys = []string {"Asic"}

    return &(fmacAsic.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats
// aggr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // link error status.
    LinkErrorStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus

    // link counters.
    LinkCounters Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus
}

func (aggrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats) GetEntityData() *types.CommonEntityData {
    aggrStats.EntityData.YFilter = aggrStats.YFilter
    aggrStats.EntityData.YangName = "aggr-stats"
    aggrStats.EntityData.BundleName = "cisco_ios_xr"
    aggrStats.EntityData.ParentYangName = "fmac-asic"
    aggrStats.EntityData.SegmentPath = "aggr-stats"
    aggrStats.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/" + aggrStats.EntityData.SegmentPath
    aggrStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggrStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggrStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggrStats.EntityData.Children = types.NewOrderedMap()
    aggrStats.EntityData.Children.Append("link-error-status", types.YChild{"LinkErrorStatus", &aggrStats.LinkErrorStatus})
    aggrStats.EntityData.Children.Append("link-counters", types.YChild{"LinkCounters", &aggrStats.LinkCounters})
    aggrStats.EntityData.Children.Append("ovf-status", types.YChild{"OvfStatus", &aggrStats.OvfStatus})
    aggrStats.EntityData.Leafs = types.NewOrderedMap()

    aggrStats.EntityData.YListKeys = []string {}

    return &(aggrStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus
// link error status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus struct {
    EntityData types.CommonEntityData
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

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkErrorStatus) GetEntityData() *types.CommonEntityData {
    linkErrorStatus.EntityData.YFilter = linkErrorStatus.YFilter
    linkErrorStatus.EntityData.YangName = "link-error-status"
    linkErrorStatus.EntityData.BundleName = "cisco_ios_xr"
    linkErrorStatus.EntityData.ParentYangName = "aggr-stats"
    linkErrorStatus.EntityData.SegmentPath = "link-error-status"
    linkErrorStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/aggr-stats/" + linkErrorStatus.EntityData.SegmentPath
    linkErrorStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkErrorStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkErrorStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkErrorStatus.EntityData.Children = types.NewOrderedMap()
    linkErrorStatus.EntityData.Leafs = types.NewOrderedMap()
    linkErrorStatus.EntityData.Leafs.Append("link-crc-error", types.YLeaf{"LinkCrcError", linkErrorStatus.LinkCrcError})
    linkErrorStatus.EntityData.Leafs.Append("link-size-error", types.YLeaf{"LinkSizeError", linkErrorStatus.LinkSizeError})
    linkErrorStatus.EntityData.Leafs.Append("link-mis-align-error", types.YLeaf{"LinkMisAlignError", linkErrorStatus.LinkMisAlignError})
    linkErrorStatus.EntityData.Leafs.Append("link-code-group-error", types.YLeaf{"LinkCodeGroupError", linkErrorStatus.LinkCodeGroupError})
    linkErrorStatus.EntityData.Leafs.Append("link-no-sig-lock-error", types.YLeaf{"LinkNoSigLockError", linkErrorStatus.LinkNoSigLockError})
    linkErrorStatus.EntityData.Leafs.Append("link-no-sig-accept-error", types.YLeaf{"LinkNoSigAcceptError", linkErrorStatus.LinkNoSigAcceptError})
    linkErrorStatus.EntityData.Leafs.Append("link-tokens-error", types.YLeaf{"LinkTokensError", linkErrorStatus.LinkTokensError})
    linkErrorStatus.EntityData.Leafs.Append("error-token-count", types.YLeaf{"ErrorTokenCount", linkErrorStatus.ErrorTokenCount})

    linkErrorStatus.EntityData.YListKeys = []string {}

    return &(linkErrorStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters
// link counters
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters struct {
    EntityData types.CommonEntityData
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
    Rx8b10bDisparityErrors interface{}

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
    Rx8b10bCodeErrors interface{}
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_LinkCounters) GetEntityData() *types.CommonEntityData {
    linkCounters.EntityData.YFilter = linkCounters.YFilter
    linkCounters.EntityData.YangName = "link-counters"
    linkCounters.EntityData.BundleName = "cisco_ios_xr"
    linkCounters.EntityData.ParentYangName = "aggr-stats"
    linkCounters.EntityData.SegmentPath = "link-counters"
    linkCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/aggr-stats/" + linkCounters.EntityData.SegmentPath
    linkCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkCounters.EntityData.Children = types.NewOrderedMap()
    linkCounters.EntityData.Leafs = types.NewOrderedMap()
    linkCounters.EntityData.Leafs.Append("tx-control-cells-counter", types.YLeaf{"TxControlCellsCounter", linkCounters.TxControlCellsCounter})
    linkCounters.EntityData.Leafs.Append("tx-data-cell-counter", types.YLeaf{"TxDataCellCounter", linkCounters.TxDataCellCounter})
    linkCounters.EntityData.Leafs.Append("tx-data-byte-counter", types.YLeaf{"TxDataByteCounter", linkCounters.TxDataByteCounter})
    linkCounters.EntityData.Leafs.Append("rx-crc-errors-counter", types.YLeaf{"RxCrcErrorsCounter", linkCounters.RxCrcErrorsCounter})
    linkCounters.EntityData.Leafs.Append("rx-lfec-fec-correctable-error", types.YLeaf{"RxLfecFecCorrectableError", linkCounters.RxLfecFecCorrectableError})
    linkCounters.EntityData.Leafs.Append("rx-8b-10b-disparity-errors", types.YLeaf{"Rx8b10bDisparityErrors", linkCounters.Rx8b10bDisparityErrors})
    linkCounters.EntityData.Leafs.Append("rx-control-cells-counter", types.YLeaf{"RxControlCellsCounter", linkCounters.RxControlCellsCounter})
    linkCounters.EntityData.Leafs.Append("rx-data-cell-counter", types.YLeaf{"RxDataCellCounter", linkCounters.RxDataCellCounter})
    linkCounters.EntityData.Leafs.Append("rx-data-byte-counter", types.YLeaf{"RxDataByteCounter", linkCounters.RxDataByteCounter})
    linkCounters.EntityData.Leafs.Append("rx-dropped-retransmitted-control", types.YLeaf{"RxDroppedRetransmittedControl", linkCounters.RxDroppedRetransmittedControl})
    linkCounters.EntityData.Leafs.Append("tx-asyn-fifo-rate", types.YLeaf{"TxAsynFifoRate", linkCounters.TxAsynFifoRate})
    linkCounters.EntityData.Leafs.Append("rx-asyn-fifo-rate", types.YLeaf{"RxAsynFifoRate", linkCounters.RxAsynFifoRate})
    linkCounters.EntityData.Leafs.Append("rx-lfec-fec-uncorrectable-errors", types.YLeaf{"RxLfecFecUncorrectableErrors", linkCounters.RxLfecFecUncorrectableErrors})
    linkCounters.EntityData.Leafs.Append("rx-8b-10b-code-errors", types.YLeaf{"Rx8b10bCodeErrors", linkCounters.Rx8b10bCodeErrors})

    linkCounters.EntityData.YListKeys = []string {}

    return &(linkCounters.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus struct {
    EntityData types.CommonEntityData
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
    Rx8b10bDisparityErrors interface{}

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
    Rx8b10bCodeErrors interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_AggrStats_OvfStatus) GetEntityData() *types.CommonEntityData {
    ovfStatus.EntityData.YFilter = ovfStatus.YFilter
    ovfStatus.EntityData.YangName = "ovf-status"
    ovfStatus.EntityData.BundleName = "cisco_ios_xr"
    ovfStatus.EntityData.ParentYangName = "aggr-stats"
    ovfStatus.EntityData.SegmentPath = "ovf-status"
    ovfStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/aggr-stats/" + ovfStatus.EntityData.SegmentPath
    ovfStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ovfStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ovfStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ovfStatus.EntityData.Children = types.NewOrderedMap()
    ovfStatus.EntityData.Leafs = types.NewOrderedMap()
    ovfStatus.EntityData.Leafs.Append("tx-control-cells-counter", types.YLeaf{"TxControlCellsCounter", ovfStatus.TxControlCellsCounter})
    ovfStatus.EntityData.Leafs.Append("tx-data-cell-counter", types.YLeaf{"TxDataCellCounter", ovfStatus.TxDataCellCounter})
    ovfStatus.EntityData.Leafs.Append("tx-data-byte-counter", types.YLeaf{"TxDataByteCounter", ovfStatus.TxDataByteCounter})
    ovfStatus.EntityData.Leafs.Append("rx-crc-errors-counter", types.YLeaf{"RxCrcErrorsCounter", ovfStatus.RxCrcErrorsCounter})
    ovfStatus.EntityData.Leafs.Append("rx-lfec-fec-correctable-error", types.YLeaf{"RxLfecFecCorrectableError", ovfStatus.RxLfecFecCorrectableError})
    ovfStatus.EntityData.Leafs.Append("rx-8b-10b-disparity-errors", types.YLeaf{"Rx8b10bDisparityErrors", ovfStatus.Rx8b10bDisparityErrors})
    ovfStatus.EntityData.Leafs.Append("rx-control-cells-counter", types.YLeaf{"RxControlCellsCounter", ovfStatus.RxControlCellsCounter})
    ovfStatus.EntityData.Leafs.Append("rx-data-cell-counter", types.YLeaf{"RxDataCellCounter", ovfStatus.RxDataCellCounter})
    ovfStatus.EntityData.Leafs.Append("rx-data-byte-counter", types.YLeaf{"RxDataByteCounter", ovfStatus.RxDataByteCounter})
    ovfStatus.EntityData.Leafs.Append("rx-dropped-retransmitted-control", types.YLeaf{"RxDroppedRetransmittedControl", ovfStatus.RxDroppedRetransmittedControl})
    ovfStatus.EntityData.Leafs.Append("tx-asyn-fifo-rate", types.YLeaf{"TxAsynFifoRate", ovfStatus.TxAsynFifoRate})
    ovfStatus.EntityData.Leafs.Append("rx-asyn-fifo-rate", types.YLeaf{"RxAsynFifoRate", ovfStatus.RxAsynFifoRate})
    ovfStatus.EntityData.Leafs.Append("rx-lfec-fec-uncorrectable-errors", types.YLeaf{"RxLfecFecUncorrectableErrors", ovfStatus.RxLfecFecUncorrectableErrors})
    ovfStatus.EntityData.Leafs.Append("rx-8b-10b-code-errors", types.YLeaf{"Rx8b10bCodeErrors", ovfStatus.Rx8b10bCodeErrors})

    ovfStatus.EntityData.YListKeys = []string {}

    return &(ovfStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats
// incr stats
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // link error status.
    LinkErrorStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus

    // link counters.
    LinkCounters Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters

    // ovf status.
    OvfStatus Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus
}

func (incrStats *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats) GetEntityData() *types.CommonEntityData {
    incrStats.EntityData.YFilter = incrStats.YFilter
    incrStats.EntityData.YangName = "incr-stats"
    incrStats.EntityData.BundleName = "cisco_ios_xr"
    incrStats.EntityData.ParentYangName = "fmac-asic"
    incrStats.EntityData.SegmentPath = "incr-stats"
    incrStats.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/" + incrStats.EntityData.SegmentPath
    incrStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    incrStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    incrStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    incrStats.EntityData.Children = types.NewOrderedMap()
    incrStats.EntityData.Children.Append("link-error-status", types.YChild{"LinkErrorStatus", &incrStats.LinkErrorStatus})
    incrStats.EntityData.Children.Append("link-counters", types.YChild{"LinkCounters", &incrStats.LinkCounters})
    incrStats.EntityData.Children.Append("ovf-status", types.YChild{"OvfStatus", &incrStats.OvfStatus})
    incrStats.EntityData.Leafs = types.NewOrderedMap()

    incrStats.EntityData.YListKeys = []string {}

    return &(incrStats.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus
// link error status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus struct {
    EntityData types.CommonEntityData
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

func (linkErrorStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkErrorStatus) GetEntityData() *types.CommonEntityData {
    linkErrorStatus.EntityData.YFilter = linkErrorStatus.YFilter
    linkErrorStatus.EntityData.YangName = "link-error-status"
    linkErrorStatus.EntityData.BundleName = "cisco_ios_xr"
    linkErrorStatus.EntityData.ParentYangName = "incr-stats"
    linkErrorStatus.EntityData.SegmentPath = "link-error-status"
    linkErrorStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/incr-stats/" + linkErrorStatus.EntityData.SegmentPath
    linkErrorStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkErrorStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkErrorStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkErrorStatus.EntityData.Children = types.NewOrderedMap()
    linkErrorStatus.EntityData.Leafs = types.NewOrderedMap()
    linkErrorStatus.EntityData.Leafs.Append("link-crc-error", types.YLeaf{"LinkCrcError", linkErrorStatus.LinkCrcError})
    linkErrorStatus.EntityData.Leafs.Append("link-size-error", types.YLeaf{"LinkSizeError", linkErrorStatus.LinkSizeError})
    linkErrorStatus.EntityData.Leafs.Append("link-mis-align-error", types.YLeaf{"LinkMisAlignError", linkErrorStatus.LinkMisAlignError})
    linkErrorStatus.EntityData.Leafs.Append("link-code-group-error", types.YLeaf{"LinkCodeGroupError", linkErrorStatus.LinkCodeGroupError})
    linkErrorStatus.EntityData.Leafs.Append("link-no-sig-lock-error", types.YLeaf{"LinkNoSigLockError", linkErrorStatus.LinkNoSigLockError})
    linkErrorStatus.EntityData.Leafs.Append("link-no-sig-accept-error", types.YLeaf{"LinkNoSigAcceptError", linkErrorStatus.LinkNoSigAcceptError})
    linkErrorStatus.EntityData.Leafs.Append("link-tokens-error", types.YLeaf{"LinkTokensError", linkErrorStatus.LinkTokensError})
    linkErrorStatus.EntityData.Leafs.Append("error-token-count", types.YLeaf{"ErrorTokenCount", linkErrorStatus.ErrorTokenCount})

    linkErrorStatus.EntityData.YListKeys = []string {}

    return &(linkErrorStatus.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters
// link counters
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters struct {
    EntityData types.CommonEntityData
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
    Rx8b10bDisparityErrors interface{}

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
    Rx8b10bCodeErrors interface{}
}

func (linkCounters *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_LinkCounters) GetEntityData() *types.CommonEntityData {
    linkCounters.EntityData.YFilter = linkCounters.YFilter
    linkCounters.EntityData.YangName = "link-counters"
    linkCounters.EntityData.BundleName = "cisco_ios_xr"
    linkCounters.EntityData.ParentYangName = "incr-stats"
    linkCounters.EntityData.SegmentPath = "link-counters"
    linkCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/incr-stats/" + linkCounters.EntityData.SegmentPath
    linkCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkCounters.EntityData.Children = types.NewOrderedMap()
    linkCounters.EntityData.Leafs = types.NewOrderedMap()
    linkCounters.EntityData.Leafs.Append("tx-control-cells-counter", types.YLeaf{"TxControlCellsCounter", linkCounters.TxControlCellsCounter})
    linkCounters.EntityData.Leafs.Append("tx-data-cell-counter", types.YLeaf{"TxDataCellCounter", linkCounters.TxDataCellCounter})
    linkCounters.EntityData.Leafs.Append("tx-data-byte-counter", types.YLeaf{"TxDataByteCounter", linkCounters.TxDataByteCounter})
    linkCounters.EntityData.Leafs.Append("rx-crc-errors-counter", types.YLeaf{"RxCrcErrorsCounter", linkCounters.RxCrcErrorsCounter})
    linkCounters.EntityData.Leafs.Append("rx-lfec-fec-correctable-error", types.YLeaf{"RxLfecFecCorrectableError", linkCounters.RxLfecFecCorrectableError})
    linkCounters.EntityData.Leafs.Append("rx-8b-10b-disparity-errors", types.YLeaf{"Rx8b10bDisparityErrors", linkCounters.Rx8b10bDisparityErrors})
    linkCounters.EntityData.Leafs.Append("rx-control-cells-counter", types.YLeaf{"RxControlCellsCounter", linkCounters.RxControlCellsCounter})
    linkCounters.EntityData.Leafs.Append("rx-data-cell-counter", types.YLeaf{"RxDataCellCounter", linkCounters.RxDataCellCounter})
    linkCounters.EntityData.Leafs.Append("rx-data-byte-counter", types.YLeaf{"RxDataByteCounter", linkCounters.RxDataByteCounter})
    linkCounters.EntityData.Leafs.Append("rx-dropped-retransmitted-control", types.YLeaf{"RxDroppedRetransmittedControl", linkCounters.RxDroppedRetransmittedControl})
    linkCounters.EntityData.Leafs.Append("tx-asyn-fifo-rate", types.YLeaf{"TxAsynFifoRate", linkCounters.TxAsynFifoRate})
    linkCounters.EntityData.Leafs.Append("rx-asyn-fifo-rate", types.YLeaf{"RxAsynFifoRate", linkCounters.RxAsynFifoRate})
    linkCounters.EntityData.Leafs.Append("rx-lfec-fec-uncorrectable-errors", types.YLeaf{"RxLfecFecUncorrectableErrors", linkCounters.RxLfecFecUncorrectableErrors})
    linkCounters.EntityData.Leafs.Append("rx-8b-10b-code-errors", types.YLeaf{"Rx8b10bCodeErrors", linkCounters.Rx8b10bCodeErrors})

    linkCounters.EntityData.YListKeys = []string {}

    return &(linkCounters.EntityData)
}

// Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus
// ovf status
type Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus struct {
    EntityData types.CommonEntityData
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
    Rx8b10bDisparityErrors interface{}

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
    Rx8b10bCodeErrors interface{}
}

func (ovfStatus *Fia_Nodes_Node_AsicStatistics_StatisticsAsicInstances_StatisticsAsicInstance_FmacStatistics_FmacLinks_FmacLink_FmacAsic_IncrStats_OvfStatus) GetEntityData() *types.CommonEntityData {
    ovfStatus.EntityData.YFilter = ovfStatus.YFilter
    ovfStatus.EntityData.YangName = "ovf-status"
    ovfStatus.EntityData.BundleName = "cisco_ios_xr"
    ovfStatus.EntityData.ParentYangName = "incr-stats"
    ovfStatus.EntityData.SegmentPath = "ovf-status"
    ovfStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-dnx-driver-oper:fia/nodes/node/asic-statistics/statistics-asic-instances/statistics-asic-instance/fmac-statistics/fmac-links/fmac-link/fmac-asic/incr-stats/" + ovfStatus.EntityData.SegmentPath
    ovfStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ovfStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ovfStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ovfStatus.EntityData.Children = types.NewOrderedMap()
    ovfStatus.EntityData.Leafs = types.NewOrderedMap()
    ovfStatus.EntityData.Leafs.Append("tx-control-cells-counter", types.YLeaf{"TxControlCellsCounter", ovfStatus.TxControlCellsCounter})
    ovfStatus.EntityData.Leafs.Append("tx-data-cell-counter", types.YLeaf{"TxDataCellCounter", ovfStatus.TxDataCellCounter})
    ovfStatus.EntityData.Leafs.Append("tx-data-byte-counter", types.YLeaf{"TxDataByteCounter", ovfStatus.TxDataByteCounter})
    ovfStatus.EntityData.Leafs.Append("rx-crc-errors-counter", types.YLeaf{"RxCrcErrorsCounter", ovfStatus.RxCrcErrorsCounter})
    ovfStatus.EntityData.Leafs.Append("rx-lfec-fec-correctable-error", types.YLeaf{"RxLfecFecCorrectableError", ovfStatus.RxLfecFecCorrectableError})
    ovfStatus.EntityData.Leafs.Append("rx-8b-10b-disparity-errors", types.YLeaf{"Rx8b10bDisparityErrors", ovfStatus.Rx8b10bDisparityErrors})
    ovfStatus.EntityData.Leafs.Append("rx-control-cells-counter", types.YLeaf{"RxControlCellsCounter", ovfStatus.RxControlCellsCounter})
    ovfStatus.EntityData.Leafs.Append("rx-data-cell-counter", types.YLeaf{"RxDataCellCounter", ovfStatus.RxDataCellCounter})
    ovfStatus.EntityData.Leafs.Append("rx-data-byte-counter", types.YLeaf{"RxDataByteCounter", ovfStatus.RxDataByteCounter})
    ovfStatus.EntityData.Leafs.Append("rx-dropped-retransmitted-control", types.YLeaf{"RxDroppedRetransmittedControl", ovfStatus.RxDroppedRetransmittedControl})
    ovfStatus.EntityData.Leafs.Append("tx-asyn-fifo-rate", types.YLeaf{"TxAsynFifoRate", ovfStatus.TxAsynFifoRate})
    ovfStatus.EntityData.Leafs.Append("rx-asyn-fifo-rate", types.YLeaf{"RxAsynFifoRate", ovfStatus.RxAsynFifoRate})
    ovfStatus.EntityData.Leafs.Append("rx-lfec-fec-uncorrectable-errors", types.YLeaf{"RxLfecFecUncorrectableErrors", ovfStatus.RxLfecFecUncorrectableErrors})
    ovfStatus.EntityData.Leafs.Append("rx-8b-10b-code-errors", types.YLeaf{"Rx8b10bCodeErrors", ovfStatus.Rx8b10bCodeErrors})

    ovfStatus.EntityData.YListKeys = []string {}

    return &(ovfStatus.EntityData)
}

