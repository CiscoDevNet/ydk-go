// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-fca package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpa-internal: Management LAN Operational data space
//   mpa: mpa
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package asr9k_lc_fca_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package asr9k_lc_fca_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lc-fca-oper mpa-internal}", reflect.TypeOf(MpaInternal{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal", reflect.TypeOf(MpaInternal{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-asr9k-lc-fca-oper mpa}", reflect.TypeOf(Mpa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-asr9k-lc-fca-oper:mpa", reflect.TypeOf(Mpa{}))
}

// SpaResetReason represents SPA reset reasons
type SpaResetReason string

const (
    // spa reset reason unknown
    SpaResetReason_spa_reset_reason_unknown SpaResetReason = "spa-reset-reason-unknown"

    // spa reset reason manual
    SpaResetReason_spa_reset_reason_manual SpaResetReason = "spa-reset-reason-manual"

    // spa reset reason fpd upgrade
    SpaResetReason_spa_reset_reason_fpd_upgrade SpaResetReason = "spa-reset-reason-fpd-upgrade"

    // spa reset reason audit fail
    SpaResetReason_spa_reset_reason_audit_fail SpaResetReason = "spa-reset-reason-audit-fail"

    // spa reset reason failure
    SpaResetReason_spa_reset_reason_failure SpaResetReason = "spa-reset-reason-failure"
)

// SpaFailureReason represents SPA failure reasons
type SpaFailureReason string

const (
    // spa failure reason unknown
    SpaFailureReason_spa_failure_reason_unknown SpaFailureReason = "spa-failure-reason-unknown"

    // spa failure reason spi failure
    SpaFailureReason_spa_failure_reason_spi_failure SpaFailureReason = "spa-failure-reason-spi-failure"

    // spa failure reason boot
    SpaFailureReason_spa_failure_reason_boot SpaFailureReason = "spa-failure-reason-boot"

    // spa failure reason hw failed
    SpaFailureReason_spa_failure_reason_hw_failed SpaFailureReason = "spa-failure-reason-hw-failed"

    // spa failure reason sw failed
    SpaFailureReason_spa_failure_reason_sw_failed SpaFailureReason = "spa-failure-reason-sw-failed"

    // spa failure reason sw restart
    SpaFailureReason_spa_failure_reason_sw_restart SpaFailureReason = "spa-failure-reason-sw-restart"

    // spa failure reason check fpd
    SpaFailureReason_spa_failure_reason_check_fpd SpaFailureReason = "spa-failure-reason-check-fpd"

    // spa failure reason read type
    SpaFailureReason_spa_failure_reason_read_type SpaFailureReason = "spa-failure-reason-read-type"
)

// SpaOperState represents SPA operational states
type SpaOperState string

const (
    // spa state reset
    SpaOperState_spa_state_reset SpaOperState = "spa-state-reset"

    // spa state failed
    SpaOperState_spa_state_failed SpaOperState = "spa-state-failed"

    // spa state booting
    SpaOperState_spa_state_booting SpaOperState = "spa-state-booting"

    // spa state ready
    SpaOperState_spa_state_ready SpaOperState = "spa-state-ready"
)

// MpaInternal
// Management LAN Operational data space
type MpaInternal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes MpaInternal_Nodes
}

func (mpaInternal *MpaInternal) GetEntityData() *types.CommonEntityData {
    mpaInternal.EntityData.YFilter = mpaInternal.YFilter
    mpaInternal.EntityData.YangName = "mpa-internal"
    mpaInternal.EntityData.BundleName = "cisco_ios_xr"
    mpaInternal.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lc-fca-oper"
    mpaInternal.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal"
    mpaInternal.EntityData.AbsolutePath = mpaInternal.EntityData.SegmentPath
    mpaInternal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpaInternal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpaInternal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpaInternal.EntityData.Children = types.NewOrderedMap()
    mpaInternal.EntityData.Children.Append("nodes", types.YChild{"Nodes", &mpaInternal.Nodes})
    mpaInternal.EntityData.Leafs = types.NewOrderedMap()

    mpaInternal.EntityData.YListKeys = []string {}

    return &(mpaInternal.EntityData)
}

// MpaInternal_Nodes
// Table of nodes
type MpaInternal_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of MpaInternal_Nodes_Node.
    Node []*MpaInternal_Nodes_Node
}

func (nodes *MpaInternal_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "mpa-internal"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal/" + nodes.EntityData.SegmentPath
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

// MpaInternal_Nodes_Node
// Number
type MpaInternal_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Number. The type is slice of MpaInternal_Nodes_Node_Bay.
    Bay []*MpaInternal_Nodes_Node_Bay
}

func (node *MpaInternal_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Node, "node")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("bay", types.YChild{"Bay", nil})
    for i := range node.Bay {
        node.EntityData.Children.Append(types.GetSegmentPath(node.Bay[i]), types.YChild{"Bay", node.Bay[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node", types.YLeaf{"Node", node.Node})

    node.EntityData.YListKeys = []string {"Node"}

    return &(node.EntityData)
}

// MpaInternal_Nodes_Node_Bay
// Number
type MpaInternal_Nodes_Node_Bay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. bay number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // Table of Ifsubsys.
    Ifsubsies MpaInternal_Nodes_Node_Bay_Ifsubsies
}

func (bay *MpaInternal_Nodes_Node_Bay) GetEntityData() *types.CommonEntityData {
    bay.EntityData.YFilter = bay.YFilter
    bay.EntityData.YangName = "bay"
    bay.EntityData.BundleName = "cisco_ios_xr"
    bay.EntityData.ParentYangName = "node"
    bay.EntityData.SegmentPath = "bay" + types.AddKeyToken(bay.Number, "number")
    bay.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal/nodes/node/" + bay.EntityData.SegmentPath
    bay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bay.EntityData.Children = types.NewOrderedMap()
    bay.EntityData.Children.Append("ifsubsies", types.YChild{"Ifsubsies", &bay.Ifsubsies})
    bay.EntityData.Leafs = types.NewOrderedMap()
    bay.EntityData.Leafs.Append("number", types.YLeaf{"Number", bay.Number})

    bay.EntityData.YListKeys = []string {"Number"}

    return &(bay.EntityData)
}

// MpaInternal_Nodes_Node_Bay_Ifsubsies
// Table of Ifsubsys
type MpaInternal_Nodes_Node_Bay_Ifsubsies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy.
    Ifsubsy []*MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetEntityData() *types.CommonEntityData {
    ifsubsies.EntityData.YFilter = ifsubsies.YFilter
    ifsubsies.EntityData.YangName = "ifsubsies"
    ifsubsies.EntityData.BundleName = "cisco_ios_xr"
    ifsubsies.EntityData.ParentYangName = "bay"
    ifsubsies.EntityData.SegmentPath = "ifsubsies"
    ifsubsies.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal/nodes/node/bay/" + ifsubsies.EntityData.SegmentPath
    ifsubsies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifsubsies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifsubsies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifsubsies.EntityData.Children = types.NewOrderedMap()
    ifsubsies.EntityData.Children.Append("ifsubsy", types.YChild{"Ifsubsy", nil})
    for i := range ifsubsies.Ifsubsy {
        ifsubsies.EntityData.Children.Append(types.GetSegmentPath(ifsubsies.Ifsubsy[i]), types.YChild{"Ifsubsy", ifsubsies.Ifsubsy[i]})
    }
    ifsubsies.EntityData.Leafs = types.NewOrderedMap()

    ifsubsies.EntityData.YListKeys = []string {}

    return &(ifsubsies.EntityData)
}

// MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy
// Number
type MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ifsubsys number. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    Number interface{}

    // mpa internal info.
    MpaInternalInfo MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetEntityData() *types.CommonEntityData {
    ifsubsy.EntityData.YFilter = ifsubsy.YFilter
    ifsubsy.EntityData.YangName = "ifsubsy"
    ifsubsy.EntityData.BundleName = "cisco_ios_xr"
    ifsubsy.EntityData.ParentYangName = "ifsubsies"
    ifsubsy.EntityData.SegmentPath = "ifsubsy" + types.AddKeyToken(ifsubsy.Number, "number")
    ifsubsy.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal/nodes/node/bay/ifsubsies/" + ifsubsy.EntityData.SegmentPath
    ifsubsy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ifsubsy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ifsubsy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ifsubsy.EntityData.Children = types.NewOrderedMap()
    ifsubsy.EntityData.Children.Append("mpa-internal-info", types.YChild{"MpaInternalInfo", &ifsubsy.MpaInternalInfo})
    ifsubsy.EntityData.Leafs = types.NewOrderedMap()
    ifsubsy.EntityData.Leafs.Append("number", types.YLeaf{"Number", ifsubsy.Number})

    ifsubsy.EntityData.YListKeys = []string {"Number"}

    return &(ifsubsy.EntityData)
}

// MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo
// mpa internal info
type MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // bay. The type is interface{} with range: 0..4294967295.
    Bay interface{}

    // ifsubsys. The type is interface{} with range: 0..4294967295.
    Ifsubsys interface{}

    // if state. The type is interface{} with range: 0..255.
    IfState interface{}

    // if event. The type is interface{} with range: 0..255.
    IfEvent interface{}

    // ep type. The type is interface{} with range: 0..4294967295.
    EpType interface{}

    // ep state. The type is interface{} with range: 0..255.
    EpState interface{}

    // ep presence. The type is interface{} with range: 0..255.
    EpPresence interface{}

    // ep idprom major. The type is interface{} with range: 0..255.
    EpIdpromMajor interface{}

    // ep idprom minor. The type is interface{} with range: 0..255.
    EpIdpromMinor interface{}

    // ep idprom data. The type is string with length: 0..256.
    EpIdpromData interface{}
}

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetEntityData() *types.CommonEntityData {
    mpaInternalInfo.EntityData.YFilter = mpaInternalInfo.YFilter
    mpaInternalInfo.EntityData.YangName = "mpa-internal-info"
    mpaInternalInfo.EntityData.BundleName = "cisco_ios_xr"
    mpaInternalInfo.EntityData.ParentYangName = "ifsubsy"
    mpaInternalInfo.EntityData.SegmentPath = "mpa-internal-info"
    mpaInternalInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal/nodes/node/bay/ifsubsies/ifsubsy/" + mpaInternalInfo.EntityData.SegmentPath
    mpaInternalInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpaInternalInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpaInternalInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpaInternalInfo.EntityData.Children = types.NewOrderedMap()
    mpaInternalInfo.EntityData.Leafs = types.NewOrderedMap()
    mpaInternalInfo.EntityData.Leafs.Append("bay", types.YLeaf{"Bay", mpaInternalInfo.Bay})
    mpaInternalInfo.EntityData.Leafs.Append("ifsubsys", types.YLeaf{"Ifsubsys", mpaInternalInfo.Ifsubsys})
    mpaInternalInfo.EntityData.Leafs.Append("if-state", types.YLeaf{"IfState", mpaInternalInfo.IfState})
    mpaInternalInfo.EntityData.Leafs.Append("if-event", types.YLeaf{"IfEvent", mpaInternalInfo.IfEvent})
    mpaInternalInfo.EntityData.Leafs.Append("ep-type", types.YLeaf{"EpType", mpaInternalInfo.EpType})
    mpaInternalInfo.EntityData.Leafs.Append("ep-state", types.YLeaf{"EpState", mpaInternalInfo.EpState})
    mpaInternalInfo.EntityData.Leafs.Append("ep-presence", types.YLeaf{"EpPresence", mpaInternalInfo.EpPresence})
    mpaInternalInfo.EntityData.Leafs.Append("ep-idprom-major", types.YLeaf{"EpIdpromMajor", mpaInternalInfo.EpIdpromMajor})
    mpaInternalInfo.EntityData.Leafs.Append("ep-idprom-minor", types.YLeaf{"EpIdpromMinor", mpaInternalInfo.EpIdpromMinor})
    mpaInternalInfo.EntityData.Leafs.Append("ep-idprom-data", types.YLeaf{"EpIdpromData", mpaInternalInfo.EpIdpromData})

    mpaInternalInfo.EntityData.YListKeys = []string {}

    return &(mpaInternalInfo.EntityData)
}

// Mpa
// mpa
type Mpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes Mpa_Nodes
}

func (mpa *Mpa) GetEntityData() *types.CommonEntityData {
    mpa.EntityData.YFilter = mpa.YFilter
    mpa.EntityData.YangName = "mpa"
    mpa.EntityData.BundleName = "cisco_ios_xr"
    mpa.EntityData.ParentYangName = "Cisco-IOS-XR-asr9k-lc-fca-oper"
    mpa.EntityData.SegmentPath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa"
    mpa.EntityData.AbsolutePath = mpa.EntityData.SegmentPath
    mpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpa.EntityData.Children = types.NewOrderedMap()
    mpa.EntityData.Children.Append("nodes", types.YChild{"Nodes", &mpa.Nodes})
    mpa.EntityData.Leafs = types.NewOrderedMap()

    mpa.EntityData.YListKeys = []string {}

    return &(mpa.EntityData)
}

// Mpa_Nodes
// Table of nodes
type Mpa_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number. The type is slice of Mpa_Nodes_Node.
    Node []*Mpa_Nodes_Node
}

func (nodes *Mpa_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "mpa"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa/" + nodes.EntityData.SegmentPath
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

// Mpa_Nodes_Node
// Number
type Mpa_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Number. The type is slice of Mpa_Nodes_Node_Bay.
    Bay []*Mpa_Nodes_Node_Bay
}

func (node *Mpa_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Node, "node")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("bay", types.YChild{"Bay", nil})
    for i := range node.Bay {
        node.EntityData.Children.Append(types.GetSegmentPath(node.Bay[i]), types.YChild{"Bay", node.Bay[i]})
    }
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node", types.YLeaf{"Node", node.Node})

    node.EntityData.YListKeys = []string {"Node"}

    return &(node.EntityData)
}

// Mpa_Nodes_Node_Bay
// Number
type Mpa_Nodes_Node_Bay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. bay number. The type is interface{} with range:
    // 0..4294967295.
    Number interface{}

    // Table of Mpa Detail Info.
    MpaDetailTable Mpa_Nodes_Node_Bay_MpaDetailTable
}

func (bay *Mpa_Nodes_Node_Bay) GetEntityData() *types.CommonEntityData {
    bay.EntityData.YFilter = bay.YFilter
    bay.EntityData.YangName = "bay"
    bay.EntityData.BundleName = "cisco_ios_xr"
    bay.EntityData.ParentYangName = "node"
    bay.EntityData.SegmentPath = "bay" + types.AddKeyToken(bay.Number, "number")
    bay.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa/nodes/node/" + bay.EntityData.SegmentPath
    bay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bay.EntityData.Children = types.NewOrderedMap()
    bay.EntityData.Children.Append("mpa-detail-table", types.YChild{"MpaDetailTable", &bay.MpaDetailTable})
    bay.EntityData.Leafs = types.NewOrderedMap()
    bay.EntityData.Leafs.Append("number", types.YLeaf{"Number", bay.Number})

    bay.EntityData.YListKeys = []string {"Number"}

    return &(bay.EntityData)
}

// Mpa_Nodes_Node_Bay_MpaDetailTable
// Table of Mpa Detail Info
type Mpa_Nodes_Node_Bay_MpaDetailTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // mpa detail status info.
    MpaDetail Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetEntityData() *types.CommonEntityData {
    mpaDetailTable.EntityData.YFilter = mpaDetailTable.YFilter
    mpaDetailTable.EntityData.YangName = "mpa-detail-table"
    mpaDetailTable.EntityData.BundleName = "cisco_ios_xr"
    mpaDetailTable.EntityData.ParentYangName = "bay"
    mpaDetailTable.EntityData.SegmentPath = "mpa-detail-table"
    mpaDetailTable.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa/nodes/node/bay/" + mpaDetailTable.EntityData.SegmentPath
    mpaDetailTable.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpaDetailTable.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpaDetailTable.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpaDetailTable.EntityData.Children = types.NewOrderedMap()
    mpaDetailTable.EntityData.Children.Append("mpa-detail", types.YChild{"MpaDetail", &mpaDetailTable.MpaDetail})
    mpaDetailTable.EntityData.Leafs = types.NewOrderedMap()

    mpaDetailTable.EntityData.YListKeys = []string {}

    return &(mpaDetailTable.EntityData)
}

// Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail
// mpa detail status info
type Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BAY number. The type is interface{} with range: 0..65535.
    BayNumber interface{}

    // If SPA inserted. The type is bool.
    IsSpaInserted interface{}

    // SPA type. The type is interface{} with range: 0..65535.
    SpaType interface{}

    // If SPA admin state is Up. The type is bool.
    IsSpaAdminUp interface{}

    // SPA operational state. The type is SpaOperState.
    SpaOperState interface{}

    // If SPA power admin state is Up. The type is bool.
    IsSpaPowerAdminUp interface{}

    // If SPA powered. The type is bool.
    IsSpaPowered interface{}

    // If SPA in reset. The type is bool.
    IsSpaInReset interface{}

    // Last reset reason. The type is SpaResetReason.
    LastResetReason interface{}

    // Last Failure Reason. The type is SpaFailureReason.
    LastFailureReason interface{}

    // Time when SPA last insertedin calendar format: seconds since00:00:00 UTC,
    // January 1, 1970. The type is interface{} with range: 0..4294967295. Units
    // are second.
    InsertionTime interface{}

    // Time when SPA last reached Ready statein calendar format: seconds
    // since00:00:00 UTC, January 1, 1970. The type is interface{} with range:
    // 0..4294967295. Units are second.
    LastReadyTime interface{}

    // Uptime in seconds. The type is interface{} with range: 0..4294967295. Units
    // are second.
    UpTime interface{}
}

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetEntityData() *types.CommonEntityData {
    mpaDetail.EntityData.YFilter = mpaDetail.YFilter
    mpaDetail.EntityData.YangName = "mpa-detail"
    mpaDetail.EntityData.BundleName = "cisco_ios_xr"
    mpaDetail.EntityData.ParentYangName = "mpa-detail-table"
    mpaDetail.EntityData.SegmentPath = "mpa-detail"
    mpaDetail.EntityData.AbsolutePath = "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa/nodes/node/bay/mpa-detail-table/" + mpaDetail.EntityData.SegmentPath
    mpaDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpaDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpaDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpaDetail.EntityData.Children = types.NewOrderedMap()
    mpaDetail.EntityData.Leafs = types.NewOrderedMap()
    mpaDetail.EntityData.Leafs.Append("bay-number", types.YLeaf{"BayNumber", mpaDetail.BayNumber})
    mpaDetail.EntityData.Leafs.Append("is-spa-inserted", types.YLeaf{"IsSpaInserted", mpaDetail.IsSpaInserted})
    mpaDetail.EntityData.Leafs.Append("spa-type", types.YLeaf{"SpaType", mpaDetail.SpaType})
    mpaDetail.EntityData.Leafs.Append("is-spa-admin-up", types.YLeaf{"IsSpaAdminUp", mpaDetail.IsSpaAdminUp})
    mpaDetail.EntityData.Leafs.Append("spa-oper-state", types.YLeaf{"SpaOperState", mpaDetail.SpaOperState})
    mpaDetail.EntityData.Leafs.Append("is-spa-power-admin-up", types.YLeaf{"IsSpaPowerAdminUp", mpaDetail.IsSpaPowerAdminUp})
    mpaDetail.EntityData.Leafs.Append("is-spa-powered", types.YLeaf{"IsSpaPowered", mpaDetail.IsSpaPowered})
    mpaDetail.EntityData.Leafs.Append("is-spa-in-reset", types.YLeaf{"IsSpaInReset", mpaDetail.IsSpaInReset})
    mpaDetail.EntityData.Leafs.Append("last-reset-reason", types.YLeaf{"LastResetReason", mpaDetail.LastResetReason})
    mpaDetail.EntityData.Leafs.Append("last-failure-reason", types.YLeaf{"LastFailureReason", mpaDetail.LastFailureReason})
    mpaDetail.EntityData.Leafs.Append("insertion-time", types.YLeaf{"InsertionTime", mpaDetail.InsertionTime})
    mpaDetail.EntityData.Leafs.Append("last-ready-time", types.YLeaf{"LastReadyTime", mpaDetail.LastReadyTime})
    mpaDetail.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", mpaDetail.UpTime})

    mpaDetail.EntityData.YListKeys = []string {}

    return &(mpaDetail.EntityData)
}

