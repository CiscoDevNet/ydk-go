// This module contains a collection of YANG definitions
// for Cisco IOS-XR asr9k-lc-fca package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpa-internal: Management LAN Operational data space
//   mpa: mpa
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes MpaInternal_Nodes
}

func (mpaInternal *MpaInternal) GetFilter() yfilter.YFilter { return mpaInternal.YFilter }

func (mpaInternal *MpaInternal) SetFilter(yf yfilter.YFilter) { mpaInternal.YFilter = yf }

func (mpaInternal *MpaInternal) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (mpaInternal *MpaInternal) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa-internal"
}

func (mpaInternal *MpaInternal) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &mpaInternal.Nodes
    }
    return nil
}

func (mpaInternal *MpaInternal) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &mpaInternal.Nodes
    return children
}

func (mpaInternal *MpaInternal) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpaInternal *MpaInternal) GetBundleName() string { return "cisco_ios_xr" }

func (mpaInternal *MpaInternal) GetYangName() string { return "mpa-internal" }

func (mpaInternal *MpaInternal) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpaInternal *MpaInternal) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpaInternal *MpaInternal) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpaInternal *MpaInternal) SetParent(parent types.Entity) { mpaInternal.parent = parent }

func (mpaInternal *MpaInternal) GetParent() types.Entity { return mpaInternal.parent }

func (mpaInternal *MpaInternal) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lc-fca-oper" }

// MpaInternal_Nodes
// Table of nodes
type MpaInternal_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of MpaInternal_Nodes_Node.
    Node []MpaInternal_Nodes_Node
}

func (nodes *MpaInternal_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *MpaInternal_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *MpaInternal_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *MpaInternal_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *MpaInternal_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MpaInternal_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *MpaInternal_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *MpaInternal_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *MpaInternal_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *MpaInternal_Nodes) GetYangName() string { return "nodes" }

func (nodes *MpaInternal_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *MpaInternal_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *MpaInternal_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *MpaInternal_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *MpaInternal_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *MpaInternal_Nodes) GetParentYangName() string { return "mpa-internal" }

// MpaInternal_Nodes_Node
// Number
type MpaInternal_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Number. The type is slice of MpaInternal_Nodes_Node_Bay.
    Bay []MpaInternal_Nodes_Node_Bay
}

func (node *MpaInternal_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *MpaInternal_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *MpaInternal_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "bay" { return "Bay" }
    return ""
}

func (node *MpaInternal_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *MpaInternal_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bay" {
        for _, c := range node.Bay {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MpaInternal_Nodes_Node_Bay{}
        node.Bay = append(node.Bay, child)
        return &node.Bay[len(node.Bay)-1]
    }
    return nil
}

func (node *MpaInternal_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.Bay {
        children[node.Bay[i].GetSegmentPath()] = &node.Bay[i]
    }
    return children
}

func (node *MpaInternal_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *MpaInternal_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *MpaInternal_Nodes_Node) GetYangName() string { return "node" }

func (node *MpaInternal_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *MpaInternal_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *MpaInternal_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *MpaInternal_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *MpaInternal_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *MpaInternal_Nodes_Node) GetParentYangName() string { return "nodes" }

// MpaInternal_Nodes_Node_Bay
// Number
type MpaInternal_Nodes_Node_Bay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. bay number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Table of Ifsubsys.
    Ifsubsies MpaInternal_Nodes_Node_Bay_Ifsubsies
}

func (bay *MpaInternal_Nodes_Node_Bay) GetFilter() yfilter.YFilter { return bay.YFilter }

func (bay *MpaInternal_Nodes_Node_Bay) SetFilter(yf yfilter.YFilter) { bay.YFilter = yf }

func (bay *MpaInternal_Nodes_Node_Bay) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "ifsubsies" { return "Ifsubsies" }
    return ""
}

func (bay *MpaInternal_Nodes_Node_Bay) GetSegmentPath() string {
    return "bay" + "[number='" + fmt.Sprintf("%v", bay.Number) + "']"
}

func (bay *MpaInternal_Nodes_Node_Bay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ifsubsies" {
        return &bay.Ifsubsies
    }
    return nil
}

func (bay *MpaInternal_Nodes_Node_Bay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ifsubsies"] = &bay.Ifsubsies
    return children
}

func (bay *MpaInternal_Nodes_Node_Bay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = bay.Number
    return leafs
}

func (bay *MpaInternal_Nodes_Node_Bay) GetBundleName() string { return "cisco_ios_xr" }

func (bay *MpaInternal_Nodes_Node_Bay) GetYangName() string { return "bay" }

func (bay *MpaInternal_Nodes_Node_Bay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bay *MpaInternal_Nodes_Node_Bay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bay *MpaInternal_Nodes_Node_Bay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bay *MpaInternal_Nodes_Node_Bay) SetParent(parent types.Entity) { bay.parent = parent }

func (bay *MpaInternal_Nodes_Node_Bay) GetParent() types.Entity { return bay.parent }

func (bay *MpaInternal_Nodes_Node_Bay) GetParentYangName() string { return "node" }

// MpaInternal_Nodes_Node_Bay_Ifsubsies
// Table of Ifsubsys
type MpaInternal_Nodes_Node_Bay_Ifsubsies struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy.
    Ifsubsy []MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetFilter() yfilter.YFilter { return ifsubsies.YFilter }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) SetFilter(yf yfilter.YFilter) { ifsubsies.YFilter = yf }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetGoName(yname string) string {
    if yname == "ifsubsy" { return "Ifsubsy" }
    return ""
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetSegmentPath() string {
    return "ifsubsies"
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ifsubsy" {
        for _, c := range ifsubsies.Ifsubsy {
            if ifsubsies.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy{}
        ifsubsies.Ifsubsy = append(ifsubsies.Ifsubsy, child)
        return &ifsubsies.Ifsubsy[len(ifsubsies.Ifsubsy)-1]
    }
    return nil
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ifsubsies.Ifsubsy {
        children[ifsubsies.Ifsubsy[i].GetSegmentPath()] = &ifsubsies.Ifsubsy[i]
    }
    return children
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetBundleName() string { return "cisco_ios_xr" }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetYangName() string { return "ifsubsies" }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) SetParent(parent types.Entity) { ifsubsies.parent = parent }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetParent() types.Entity { return ifsubsies.parent }

func (ifsubsies *MpaInternal_Nodes_Node_Bay_Ifsubsies) GetParentYangName() string { return "bay" }

// MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy
// Number
type MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ifsubsys number. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    Number interface{}

    // mpa internal info.
    MpaInternalInfo MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetFilter() yfilter.YFilter { return ifsubsy.YFilter }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) SetFilter(yf yfilter.YFilter) { ifsubsy.YFilter = yf }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "mpa-internal-info" { return "MpaInternalInfo" }
    return ""
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetSegmentPath() string {
    return "ifsubsy" + "[number='" + fmt.Sprintf("%v", ifsubsy.Number) + "']"
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpa-internal-info" {
        return &ifsubsy.MpaInternalInfo
    }
    return nil
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpa-internal-info"] = &ifsubsy.MpaInternalInfo
    return children
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = ifsubsy.Number
    return leafs
}

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetBundleName() string { return "cisco_ios_xr" }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetYangName() string { return "ifsubsy" }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) SetParent(parent types.Entity) { ifsubsy.parent = parent }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetParent() types.Entity { return ifsubsy.parent }

func (ifsubsy *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy) GetParentYangName() string { return "ifsubsies" }

// MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo
// mpa internal info
type MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo struct {
    parent types.Entity
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

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetFilter() yfilter.YFilter { return mpaInternalInfo.YFilter }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) SetFilter(yf yfilter.YFilter) { mpaInternalInfo.YFilter = yf }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetGoName(yname string) string {
    if yname == "bay" { return "Bay" }
    if yname == "ifsubsys" { return "Ifsubsys" }
    if yname == "if-state" { return "IfState" }
    if yname == "if-event" { return "IfEvent" }
    if yname == "ep-type" { return "EpType" }
    if yname == "ep-state" { return "EpState" }
    if yname == "ep-presence" { return "EpPresence" }
    if yname == "ep-idprom-major" { return "EpIdpromMajor" }
    if yname == "ep-idprom-minor" { return "EpIdpromMinor" }
    if yname == "ep-idprom-data" { return "EpIdpromData" }
    return ""
}

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetSegmentPath() string {
    return "mpa-internal-info"
}

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bay"] = mpaInternalInfo.Bay
    leafs["ifsubsys"] = mpaInternalInfo.Ifsubsys
    leafs["if-state"] = mpaInternalInfo.IfState
    leafs["if-event"] = mpaInternalInfo.IfEvent
    leafs["ep-type"] = mpaInternalInfo.EpType
    leafs["ep-state"] = mpaInternalInfo.EpState
    leafs["ep-presence"] = mpaInternalInfo.EpPresence
    leafs["ep-idprom-major"] = mpaInternalInfo.EpIdpromMajor
    leafs["ep-idprom-minor"] = mpaInternalInfo.EpIdpromMinor
    leafs["ep-idprom-data"] = mpaInternalInfo.EpIdpromData
    return leafs
}

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetBundleName() string { return "cisco_ios_xr" }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetYangName() string { return "mpa-internal-info" }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) SetParent(parent types.Entity) { mpaInternalInfo.parent = parent }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetParent() types.Entity { return mpaInternalInfo.parent }

func (mpaInternalInfo *MpaInternal_Nodes_Node_Bay_Ifsubsies_Ifsubsy_MpaInternalInfo) GetParentYangName() string { return "ifsubsy" }

// Mpa
// mpa
type Mpa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of nodes.
    Nodes Mpa_Nodes
}

func (mpa *Mpa) GetFilter() yfilter.YFilter { return mpa.YFilter }

func (mpa *Mpa) SetFilter(yf yfilter.YFilter) { mpa.YFilter = yf }

func (mpa *Mpa) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (mpa *Mpa) GetSegmentPath() string {
    return "Cisco-IOS-XR-asr9k-lc-fca-oper:mpa"
}

func (mpa *Mpa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &mpa.Nodes
    }
    return nil
}

func (mpa *Mpa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &mpa.Nodes
    return children
}

func (mpa *Mpa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpa *Mpa) GetBundleName() string { return "cisco_ios_xr" }

func (mpa *Mpa) GetYangName() string { return "mpa" }

func (mpa *Mpa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpa *Mpa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpa *Mpa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpa *Mpa) SetParent(parent types.Entity) { mpa.parent = parent }

func (mpa *Mpa) GetParent() types.Entity { return mpa.parent }

func (mpa *Mpa) GetParentYangName() string { return "Cisco-IOS-XR-asr9k-lc-fca-oper" }

// Mpa_Nodes
// Table of nodes
type Mpa_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number. The type is slice of Mpa_Nodes_Node.
    Node []Mpa_Nodes_Node
}

func (nodes *Mpa_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Mpa_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Mpa_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Mpa_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Mpa_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpa_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Mpa_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Mpa_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Mpa_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Mpa_Nodes) GetYangName() string { return "nodes" }

func (nodes *Mpa_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Mpa_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Mpa_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Mpa_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Mpa_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Mpa_Nodes) GetParentYangName() string { return "mpa" }

// Mpa_Nodes_Node
// Number
type Mpa_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. node number. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // Number. The type is slice of Mpa_Nodes_Node_Bay.
    Bay []Mpa_Nodes_Node_Bay
}

func (node *Mpa_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Mpa_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Mpa_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "bay" { return "Bay" }
    return ""
}

func (node *Mpa_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *Mpa_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bay" {
        for _, c := range node.Bay {
            if node.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mpa_Nodes_Node_Bay{}
        node.Bay = append(node.Bay, child)
        return &node.Bay[len(node.Bay)-1]
    }
    return nil
}

func (node *Mpa_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range node.Bay {
        children[node.Bay[i].GetSegmentPath()] = &node.Bay[i]
    }
    return children
}

func (node *Mpa_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *Mpa_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Mpa_Nodes_Node) GetYangName() string { return "node" }

func (node *Mpa_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Mpa_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Mpa_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Mpa_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Mpa_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Mpa_Nodes_Node) GetParentYangName() string { return "nodes" }

// Mpa_Nodes_Node_Bay
// Number
type Mpa_Nodes_Node_Bay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. bay number. The type is interface{} with range:
    // -2147483648..2147483647.
    Number interface{}

    // Table of Mpa Detail Info.
    MpaDetailTable Mpa_Nodes_Node_Bay_MpaDetailTable
}

func (bay *Mpa_Nodes_Node_Bay) GetFilter() yfilter.YFilter { return bay.YFilter }

func (bay *Mpa_Nodes_Node_Bay) SetFilter(yf yfilter.YFilter) { bay.YFilter = yf }

func (bay *Mpa_Nodes_Node_Bay) GetGoName(yname string) string {
    if yname == "number" { return "Number" }
    if yname == "mpa-detail-table" { return "MpaDetailTable" }
    return ""
}

func (bay *Mpa_Nodes_Node_Bay) GetSegmentPath() string {
    return "bay" + "[number='" + fmt.Sprintf("%v", bay.Number) + "']"
}

func (bay *Mpa_Nodes_Node_Bay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpa-detail-table" {
        return &bay.MpaDetailTable
    }
    return nil
}

func (bay *Mpa_Nodes_Node_Bay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpa-detail-table"] = &bay.MpaDetailTable
    return children
}

func (bay *Mpa_Nodes_Node_Bay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number"] = bay.Number
    return leafs
}

func (bay *Mpa_Nodes_Node_Bay) GetBundleName() string { return "cisco_ios_xr" }

func (bay *Mpa_Nodes_Node_Bay) GetYangName() string { return "bay" }

func (bay *Mpa_Nodes_Node_Bay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bay *Mpa_Nodes_Node_Bay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bay *Mpa_Nodes_Node_Bay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bay *Mpa_Nodes_Node_Bay) SetParent(parent types.Entity) { bay.parent = parent }

func (bay *Mpa_Nodes_Node_Bay) GetParent() types.Entity { return bay.parent }

func (bay *Mpa_Nodes_Node_Bay) GetParentYangName() string { return "node" }

// Mpa_Nodes_Node_Bay_MpaDetailTable
// Table of Mpa Detail Info
type Mpa_Nodes_Node_Bay_MpaDetailTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // mpa detail status info.
    MpaDetail Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetFilter() yfilter.YFilter { return mpaDetailTable.YFilter }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) SetFilter(yf yfilter.YFilter) { mpaDetailTable.YFilter = yf }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetGoName(yname string) string {
    if yname == "mpa-detail" { return "MpaDetail" }
    return ""
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetSegmentPath() string {
    return "mpa-detail-table"
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpa-detail" {
        return &mpaDetailTable.MpaDetail
    }
    return nil
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpa-detail"] = &mpaDetailTable.MpaDetail
    return children
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetBundleName() string { return "cisco_ios_xr" }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetYangName() string { return "mpa-detail-table" }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) SetParent(parent types.Entity) { mpaDetailTable.parent = parent }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetParent() types.Entity { return mpaDetailTable.parent }

func (mpaDetailTable *Mpa_Nodes_Node_Bay_MpaDetailTable) GetParentYangName() string { return "bay" }

// Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail
// mpa detail status info
type Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail struct {
    parent types.Entity
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

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetFilter() yfilter.YFilter { return mpaDetail.YFilter }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) SetFilter(yf yfilter.YFilter) { mpaDetail.YFilter = yf }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetGoName(yname string) string {
    if yname == "bay-number" { return "BayNumber" }
    if yname == "is-spa-inserted" { return "IsSpaInserted" }
    if yname == "spa-type" { return "SpaType" }
    if yname == "is-spa-admin-up" { return "IsSpaAdminUp" }
    if yname == "spa-oper-state" { return "SpaOperState" }
    if yname == "is-spa-power-admin-up" { return "IsSpaPowerAdminUp" }
    if yname == "is-spa-powered" { return "IsSpaPowered" }
    if yname == "is-spa-in-reset" { return "IsSpaInReset" }
    if yname == "last-reset-reason" { return "LastResetReason" }
    if yname == "last-failure-reason" { return "LastFailureReason" }
    if yname == "insertion-time" { return "InsertionTime" }
    if yname == "last-ready-time" { return "LastReadyTime" }
    if yname == "up-time" { return "UpTime" }
    return ""
}

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetSegmentPath() string {
    return "mpa-detail"
}

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bay-number"] = mpaDetail.BayNumber
    leafs["is-spa-inserted"] = mpaDetail.IsSpaInserted
    leafs["spa-type"] = mpaDetail.SpaType
    leafs["is-spa-admin-up"] = mpaDetail.IsSpaAdminUp
    leafs["spa-oper-state"] = mpaDetail.SpaOperState
    leafs["is-spa-power-admin-up"] = mpaDetail.IsSpaPowerAdminUp
    leafs["is-spa-powered"] = mpaDetail.IsSpaPowered
    leafs["is-spa-in-reset"] = mpaDetail.IsSpaInReset
    leafs["last-reset-reason"] = mpaDetail.LastResetReason
    leafs["last-failure-reason"] = mpaDetail.LastFailureReason
    leafs["insertion-time"] = mpaDetail.InsertionTime
    leafs["last-ready-time"] = mpaDetail.LastReadyTime
    leafs["up-time"] = mpaDetail.UpTime
    return leafs
}

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetBundleName() string { return "cisco_ios_xr" }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetYangName() string { return "mpa-detail" }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) SetParent(parent types.Entity) { mpaDetail.parent = parent }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetParent() types.Entity { return mpaDetail.parent }

func (mpaDetail *Mpa_Nodes_Node_Bay_MpaDetailTable_MpaDetail) GetParentYangName() string { return "mpa-detail-table" }

