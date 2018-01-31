// This module contains a collection of YANG definitions for
// monitoring BFD neighbours.Copyright (c) 2016-2017 by Cisco Systems, Inc.All rights reserved.
package bfd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bfd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-bfd-oper bfd-state}", reflect.TypeOf(BfdState{}))
    ydk.RegisterEntity("Cisco-IOS-XE-bfd-oper:bfd-state", reflect.TypeOf(BfdState{}))
}

// BfdRemoteStateType represents BFD remote state type
type BfdRemoteStateType string

const (
    BfdRemoteStateType_up BfdRemoteStateType = "up"

    BfdRemoteStateType_down BfdRemoteStateType = "down"

    BfdRemoteStateType_init BfdRemoteStateType = "init"

    BfdRemoteStateType_admindown BfdRemoteStateType = "admindown"

    BfdRemoteStateType_invalid BfdRemoteStateType = "invalid"
)

// BfdLspType represents BFD lsp type
type BfdLspType string

const (
    BfdLspType_working BfdLspType = "working"

    BfdLspType_protect BfdLspType = "protect"

    BfdLspType_unknown BfdLspType = "unknown"
)

// BfdOperSessionType represents BFD session type
type BfdOperSessionType string

const (
    BfdOperSessionType_ipv4 BfdOperSessionType = "ipv4"

    BfdOperSessionType_ipv6 BfdOperSessionType = "ipv6"

    BfdOperSessionType_vccv BfdOperSessionType = "vccv"

    BfdOperSessionType_mpls_tp BfdOperSessionType = "mpls-tp"

    BfdOperSessionType_ipv4_multihop BfdOperSessionType = "ipv4-multihop"

    BfdOperSessionType_ipv6_multihop BfdOperSessionType = "ipv6-multihop"
)

// BfdStateType represents BFD state type
type BfdStateType string

const (
    BfdStateType_admindown BfdStateType = "admindown"

    BfdStateType_down BfdStateType = "down"

    BfdStateType_fail BfdStateType = "fail"

    BfdStateType_init BfdStateType = "init"

    BfdStateType_up BfdStateType = "up"

    BfdStateType_invalid BfdStateType = "invalid"
)

// BfdState
// Data nodes for BFD neighbors.
type BfdState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Sessions BfdState_Sessions
}

func (bfdState *BfdState) GetFilter() yfilter.YFilter { return bfdState.YFilter }

func (bfdState *BfdState) SetFilter(yf yfilter.YFilter) { bfdState.YFilter = yf }

func (bfdState *BfdState) GetGoName(yname string) string {
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (bfdState *BfdState) GetSegmentPath() string {
    return "Cisco-IOS-XE-bfd-oper:bfd-state"
}

func (bfdState *BfdState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &bfdState.Sessions
    }
    return nil
}

func (bfdState *BfdState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &bfdState.Sessions
    return children
}

func (bfdState *BfdState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdState *BfdState) GetBundleName() string { return "cisco_ios_xe" }

func (bfdState *BfdState) GetYangName() string { return "bfd-state" }

func (bfdState *BfdState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdState *BfdState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdState *BfdState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdState *BfdState) SetParent(parent types.Entity) { bfdState.parent = parent }

func (bfdState *BfdState) GetParent() types.Entity { return bfdState.parent }

func (bfdState *BfdState) GetParentYangName() string { return "Cisco-IOS-XE-bfd-oper" }

// BfdState_Sessions
type BfdState_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The type is slice of BfdState_Sessions_Session.
    Session []BfdState_Sessions_Session
}

func (sessions *BfdState_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *BfdState_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *BfdState_Sessions) GetGoName(yname string) string {
    if yname == "session" { return "Session" }
    return ""
}

func (sessions *BfdState_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *BfdState_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        for _, c := range sessions.Session {
            if sessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BfdState_Sessions_Session{}
        sessions.Session = append(sessions.Session, child)
        return &sessions.Session[len(sessions.Session)-1]
    }
    return nil
}

func (sessions *BfdState_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessions.Session {
        children[sessions.Session[i].GetSegmentPath()] = &sessions.Session[i]
    }
    return children
}

func (sessions *BfdState_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *BfdState_Sessions) GetBundleName() string { return "cisco_ios_xe" }

func (sessions *BfdState_Sessions) GetYangName() string { return "sessions" }

func (sessions *BfdState_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sessions *BfdState_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sessions *BfdState_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sessions *BfdState_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *BfdState_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *BfdState_Sessions) GetParentYangName() string { return "bfd-state" }

// BfdState_Sessions_Session
type BfdState_Sessions_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session type. The type is BfdOperSessionType.
    Type interface{}

    
    BfdTunnelPaths BfdState_Sessions_Session_BfdTunnelPaths

    
    BfdCircuits BfdState_Sessions_Session_BfdCircuits

    
    BfdNbrs BfdState_Sessions_Session_BfdNbrs

    
    BfdMhopNbrs BfdState_Sessions_Session_BfdMhopNbrs

    
    BfdMhopVrfNbrs BfdState_Sessions_Session_BfdMhopVrfNbrs
}

func (session *BfdState_Sessions_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *BfdState_Sessions_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *BfdState_Sessions_Session) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "bfd-tunnel-paths" { return "BfdTunnelPaths" }
    if yname == "bfd-circuits" { return "BfdCircuits" }
    if yname == "bfd-nbrs" { return "BfdNbrs" }
    if yname == "bfd-mhop-nbrs" { return "BfdMhopNbrs" }
    if yname == "bfd-mhop-vrf-nbrs" { return "BfdMhopVrfNbrs" }
    return ""
}

func (session *BfdState_Sessions_Session) GetSegmentPath() string {
    return "session" + "[type='" + fmt.Sprintf("%v", session.Type) + "']"
}

func (session *BfdState_Sessions_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-tunnel-paths" {
        return &session.BfdTunnelPaths
    }
    if childYangName == "bfd-circuits" {
        return &session.BfdCircuits
    }
    if childYangName == "bfd-nbrs" {
        return &session.BfdNbrs
    }
    if childYangName == "bfd-mhop-nbrs" {
        return &session.BfdMhopNbrs
    }
    if childYangName == "bfd-mhop-vrf-nbrs" {
        return &session.BfdMhopVrfNbrs
    }
    return nil
}

func (session *BfdState_Sessions_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bfd-tunnel-paths"] = &session.BfdTunnelPaths
    children["bfd-circuits"] = &session.BfdCircuits
    children["bfd-nbrs"] = &session.BfdNbrs
    children["bfd-mhop-nbrs"] = &session.BfdMhopNbrs
    children["bfd-mhop-vrf-nbrs"] = &session.BfdMhopVrfNbrs
    return children
}

func (session *BfdState_Sessions_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = session.Type
    return leafs
}

func (session *BfdState_Sessions_Session) GetBundleName() string { return "cisco_ios_xe" }

func (session *BfdState_Sessions_Session) GetYangName() string { return "session" }

func (session *BfdState_Sessions_Session) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (session *BfdState_Sessions_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (session *BfdState_Sessions_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (session *BfdState_Sessions_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *BfdState_Sessions_Session) GetParent() types.Entity { return session.parent }

func (session *BfdState_Sessions_Session) GetParentYangName() string { return "sessions" }

// BfdState_Sessions_Session_BfdTunnelPaths
type BfdState_Sessions_Session_BfdTunnelPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel Path. The type is slice of
    // BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath.
    BfdTunnelPath []BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetFilter() yfilter.YFilter { return bfdTunnelPaths.YFilter }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) SetFilter(yf yfilter.YFilter) { bfdTunnelPaths.YFilter = yf }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetGoName(yname string) string {
    if yname == "bfd-tunnel-path" { return "BfdTunnelPath" }
    return ""
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetSegmentPath() string {
    return "bfd-tunnel-paths"
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-tunnel-path" {
        for _, c := range bfdTunnelPaths.BfdTunnelPath {
            if bfdTunnelPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath{}
        bfdTunnelPaths.BfdTunnelPath = append(bfdTunnelPaths.BfdTunnelPath, child)
        return &bfdTunnelPaths.BfdTunnelPath[len(bfdTunnelPaths.BfdTunnelPath)-1]
    }
    return nil
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdTunnelPaths.BfdTunnelPath {
        children[bfdTunnelPaths.BfdTunnelPath[i].GetSegmentPath()] = &bfdTunnelPaths.BfdTunnelPath[i]
    }
    return children
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetBundleName() string { return "cisco_ios_xe" }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetYangName() string { return "bfd-tunnel-paths" }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) SetParent(parent types.Entity) { bfdTunnelPaths.parent = parent }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetParent() types.Entity { return bfdTunnelPaths.parent }

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetParentYangName() string { return "session" }

// BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath
// Tunnel Path
type BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Interface interface{}

    // This attribute is a key. The type is BfdLspType.
    LspType interface{}

    // local-discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // remote-discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'   is also mapped to up/down. .
    // The type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetFilter() yfilter.YFilter { return bfdTunnelPath.YFilter }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) SetFilter(yf yfilter.YFilter) { bfdTunnelPath.YFilter = yf }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "lsp-type" { return "LspType" }
    if yname == "ld" { return "Ld" }
    if yname == "rd" { return "Rd" }
    if yname == "remote-state" { return "RemoteState" }
    if yname == "state" { return "State" }
    return ""
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetSegmentPath() string {
    return "bfd-tunnel-path" + "[interface='" + fmt.Sprintf("%v", bfdTunnelPath.Interface) + "']" + "[lsp-type='" + fmt.Sprintf("%v", bfdTunnelPath.LspType) + "']"
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = bfdTunnelPath.Interface
    leafs["lsp-type"] = bfdTunnelPath.LspType
    leafs["ld"] = bfdTunnelPath.Ld
    leafs["rd"] = bfdTunnelPath.Rd
    leafs["remote-state"] = bfdTunnelPath.RemoteState
    leafs["state"] = bfdTunnelPath.State
    return leafs
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetBundleName() string { return "cisco_ios_xe" }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetYangName() string { return "bfd-tunnel-path" }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) SetParent(parent types.Entity) { bfdTunnelPath.parent = parent }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetParent() types.Entity { return bfdTunnelPath.parent }

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetParentYangName() string { return "bfd-tunnel-paths" }

// BfdState_Sessions_Session_BfdCircuits
type BfdState_Sessions_Session_BfdCircuits struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BFD circuit. The type is slice of
    // BfdState_Sessions_Session_BfdCircuits_BfdCircuit.
    BfdCircuit []BfdState_Sessions_Session_BfdCircuits_BfdCircuit
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetFilter() yfilter.YFilter { return bfdCircuits.YFilter }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) SetFilter(yf yfilter.YFilter) { bfdCircuits.YFilter = yf }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetGoName(yname string) string {
    if yname == "bfd-circuit" { return "BfdCircuit" }
    return ""
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetSegmentPath() string {
    return "bfd-circuits"
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-circuit" {
        for _, c := range bfdCircuits.BfdCircuit {
            if bfdCircuits.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BfdState_Sessions_Session_BfdCircuits_BfdCircuit{}
        bfdCircuits.BfdCircuit = append(bfdCircuits.BfdCircuit, child)
        return &bfdCircuits.BfdCircuit[len(bfdCircuits.BfdCircuit)-1]
    }
    return nil
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdCircuits.BfdCircuit {
        children[bfdCircuits.BfdCircuit[i].GetSegmentPath()] = &bfdCircuits.BfdCircuit[i]
    }
    return children
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetBundleName() string { return "cisco_ios_xe" }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetYangName() string { return "bfd-circuits" }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) SetParent(parent types.Entity) { bfdCircuits.parent = parent }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetParent() types.Entity { return bfdCircuits.parent }

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetParentYangName() string { return "session" }

// BfdState_Sessions_Session_BfdCircuits_BfdCircuit
// BFD circuit
type BfdState_Sessions_Session_BfdCircuits_BfdCircuit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string.
    Interface interface{}

    // This attribute is a key. The type is interface{} with range: 0..4294967295.
    Vcid interface{}

    // local-discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // remote-discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'   is also mapped to up/down. .
    // The type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetFilter() yfilter.YFilter { return bfdCircuit.YFilter }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) SetFilter(yf yfilter.YFilter) { bfdCircuit.YFilter = yf }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "vcid" { return "Vcid" }
    if yname == "ld" { return "Ld" }
    if yname == "rd" { return "Rd" }
    if yname == "remote-state" { return "RemoteState" }
    if yname == "state" { return "State" }
    return ""
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetSegmentPath() string {
    return "bfd-circuit" + "[interface='" + fmt.Sprintf("%v", bfdCircuit.Interface) + "']" + "[vcid='" + fmt.Sprintf("%v", bfdCircuit.Vcid) + "']"
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = bfdCircuit.Interface
    leafs["vcid"] = bfdCircuit.Vcid
    leafs["ld"] = bfdCircuit.Ld
    leafs["rd"] = bfdCircuit.Rd
    leafs["remote-state"] = bfdCircuit.RemoteState
    leafs["state"] = bfdCircuit.State
    return leafs
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetBundleName() string { return "cisco_ios_xe" }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetYangName() string { return "bfd-circuit" }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) SetParent(parent types.Entity) { bfdCircuit.parent = parent }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetParent() types.Entity { return bfdCircuit.parent }

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetParentYangName() string { return "bfd-circuits" }

// BfdState_Sessions_Session_BfdNbrs
type BfdState_Sessions_Session_BfdNbrs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is for directly connected neighbor case. The type is slice of
    // BfdState_Sessions_Session_BfdNbrs_BfdNbr.
    BfdNbr []BfdState_Sessions_Session_BfdNbrs_BfdNbr
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetFilter() yfilter.YFilter { return bfdNbrs.YFilter }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) SetFilter(yf yfilter.YFilter) { bfdNbrs.YFilter = yf }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetGoName(yname string) string {
    if yname == "bfd-nbr" { return "BfdNbr" }
    return ""
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetSegmentPath() string {
    return "bfd-nbrs"
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-nbr" {
        for _, c := range bfdNbrs.BfdNbr {
            if bfdNbrs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BfdState_Sessions_Session_BfdNbrs_BfdNbr{}
        bfdNbrs.BfdNbr = append(bfdNbrs.BfdNbr, child)
        return &bfdNbrs.BfdNbr[len(bfdNbrs.BfdNbr)-1]
    }
    return nil
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdNbrs.BfdNbr {
        children[bfdNbrs.BfdNbr[i].GetSegmentPath()] = &bfdNbrs.BfdNbr[i]
    }
    return children
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetBundleName() string { return "cisco_ios_xe" }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetYangName() string { return "bfd-nbrs" }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) SetParent(parent types.Entity) { bfdNbrs.parent = parent }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetParent() types.Entity { return bfdNbrs.parent }

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetParentYangName() string { return "session" }

// BfdState_Sessions_Session_BfdNbrs_BfdNbr
// This is for directly connected neighbor case
type BfdState_Sessions_Session_BfdNbrs_BfdNbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // This attribute is a key. The type is string.
    Interface interface{}

    // local-discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // remote-discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'   is also mapped to up/down. .
    // The type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetFilter() yfilter.YFilter { return bfdNbr.YFilter }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) SetFilter(yf yfilter.YFilter) { bfdNbr.YFilter = yf }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "interface" { return "Interface" }
    if yname == "ld" { return "Ld" }
    if yname == "rd" { return "Rd" }
    if yname == "remote-state" { return "RemoteState" }
    if yname == "state" { return "State" }
    return ""
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetSegmentPath() string {
    return "bfd-nbr" + "[ip='" + fmt.Sprintf("%v", bfdNbr.Ip) + "']" + "[interface='" + fmt.Sprintf("%v", bfdNbr.Interface) + "']"
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = bfdNbr.Ip
    leafs["interface"] = bfdNbr.Interface
    leafs["ld"] = bfdNbr.Ld
    leafs["rd"] = bfdNbr.Rd
    leafs["remote-state"] = bfdNbr.RemoteState
    leafs["state"] = bfdNbr.State
    return leafs
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetBundleName() string { return "cisco_ios_xe" }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetYangName() string { return "bfd-nbr" }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) SetParent(parent types.Entity) { bfdNbr.parent = parent }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetParent() types.Entity { return bfdNbr.parent }

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetParentYangName() string { return "bfd-nbrs" }

// BfdState_Sessions_Session_BfdMhopNbrs
type BfdState_Sessions_Session_BfdMhopNbrs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is for multi hop neighbor scenario  for global VRF (no VRF). The type
    // is slice of BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr.
    BfdMhopNbr []BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetFilter() yfilter.YFilter { return bfdMhopNbrs.YFilter }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) SetFilter(yf yfilter.YFilter) { bfdMhopNbrs.YFilter = yf }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetGoName(yname string) string {
    if yname == "bfd-mhop-nbr" { return "BfdMhopNbr" }
    return ""
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetSegmentPath() string {
    return "bfd-mhop-nbrs"
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-mhop-nbr" {
        for _, c := range bfdMhopNbrs.BfdMhopNbr {
            if bfdMhopNbrs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr{}
        bfdMhopNbrs.BfdMhopNbr = append(bfdMhopNbrs.BfdMhopNbr, child)
        return &bfdMhopNbrs.BfdMhopNbr[len(bfdMhopNbrs.BfdMhopNbr)-1]
    }
    return nil
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdMhopNbrs.BfdMhopNbr {
        children[bfdMhopNbrs.BfdMhopNbr[i].GetSegmentPath()] = &bfdMhopNbrs.BfdMhopNbr[i]
    }
    return children
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetBundleName() string { return "cisco_ios_xe" }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetYangName() string { return "bfd-mhop-nbrs" }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) SetParent(parent types.Entity) { bfdMhopNbrs.parent = parent }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetParent() types.Entity { return bfdMhopNbrs.parent }

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetParentYangName() string { return "session" }

// BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr
// This is for multi hop neighbor scenario 
// for global VRF (no VRF).
type BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // local-discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // remote-discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'   is also mapped to up/down. .
    // The type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetFilter() yfilter.YFilter { return bfdMhopNbr.YFilter }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) SetFilter(yf yfilter.YFilter) { bfdMhopNbr.YFilter = yf }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "ld" { return "Ld" }
    if yname == "rd" { return "Rd" }
    if yname == "remote-state" { return "RemoteState" }
    if yname == "state" { return "State" }
    return ""
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetSegmentPath() string {
    return "bfd-mhop-nbr" + "[ip='" + fmt.Sprintf("%v", bfdMhopNbr.Ip) + "']"
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = bfdMhopNbr.Ip
    leafs["ld"] = bfdMhopNbr.Ld
    leafs["rd"] = bfdMhopNbr.Rd
    leafs["remote-state"] = bfdMhopNbr.RemoteState
    leafs["state"] = bfdMhopNbr.State
    return leafs
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetBundleName() string { return "cisco_ios_xe" }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetYangName() string { return "bfd-mhop-nbr" }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) SetParent(parent types.Entity) { bfdMhopNbr.parent = parent }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetParent() types.Entity { return bfdMhopNbr.parent }

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetParentYangName() string { return "bfd-mhop-nbrs" }

// BfdState_Sessions_Session_BfdMhopVrfNbrs
type BfdState_Sessions_Session_BfdMhopVrfNbrs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is for multi hop neighbor scenario  for non-global VRF. The type is
    // slice of BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr.
    BfdMhopVrfNbr []BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetFilter() yfilter.YFilter { return bfdMhopVrfNbrs.YFilter }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) SetFilter(yf yfilter.YFilter) { bfdMhopVrfNbrs.YFilter = yf }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetGoName(yname string) string {
    if yname == "bfd-mhop-vrf-nbr" { return "BfdMhopVrfNbr" }
    return ""
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetSegmentPath() string {
    return "bfd-mhop-vrf-nbrs"
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd-mhop-vrf-nbr" {
        for _, c := range bfdMhopVrfNbrs.BfdMhopVrfNbr {
            if bfdMhopVrfNbrs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr{}
        bfdMhopVrfNbrs.BfdMhopVrfNbr = append(bfdMhopVrfNbrs.BfdMhopVrfNbr, child)
        return &bfdMhopVrfNbrs.BfdMhopVrfNbr[len(bfdMhopVrfNbrs.BfdMhopVrfNbr)-1]
    }
    return nil
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bfdMhopVrfNbrs.BfdMhopVrfNbr {
        children[bfdMhopVrfNbrs.BfdMhopVrfNbr[i].GetSegmentPath()] = &bfdMhopVrfNbrs.BfdMhopVrfNbr[i]
    }
    return children
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetBundleName() string { return "cisco_ios_xe" }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetYangName() string { return "bfd-mhop-vrf-nbrs" }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) SetParent(parent types.Entity) { bfdMhopVrfNbrs.parent = parent }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetParent() types.Entity { return bfdMhopVrfNbrs.parent }

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetParentYangName() string { return "session" }

// BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr
// This is for multi hop neighbor scenario 
// for non-global VRF.
type BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ip interface{}

    // This attribute is a key. The type is string.
    Vrf interface{}

    // local-discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // remote-discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'   is also mapped to up/down. .
    // The type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetFilter() yfilter.YFilter { return bfdMhopVrfNbr.YFilter }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) SetFilter(yf yfilter.YFilter) { bfdMhopVrfNbr.YFilter = yf }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetGoName(yname string) string {
    if yname == "ip" { return "Ip" }
    if yname == "vrf" { return "Vrf" }
    if yname == "ld" { return "Ld" }
    if yname == "rd" { return "Rd" }
    if yname == "remote-state" { return "RemoteState" }
    if yname == "state" { return "State" }
    return ""
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetSegmentPath() string {
    return "bfd-mhop-vrf-nbr" + "[ip='" + fmt.Sprintf("%v", bfdMhopVrfNbr.Ip) + "']" + "[vrf='" + fmt.Sprintf("%v", bfdMhopVrfNbr.Vrf) + "']"
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip"] = bfdMhopVrfNbr.Ip
    leafs["vrf"] = bfdMhopVrfNbr.Vrf
    leafs["ld"] = bfdMhopVrfNbr.Ld
    leafs["rd"] = bfdMhopVrfNbr.Rd
    leafs["remote-state"] = bfdMhopVrfNbr.RemoteState
    leafs["state"] = bfdMhopVrfNbr.State
    return leafs
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetBundleName() string { return "cisco_ios_xe" }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetYangName() string { return "bfd-mhop-vrf-nbr" }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) SetParent(parent types.Entity) { bfdMhopVrfNbr.parent = parent }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetParent() types.Entity { return bfdMhopVrfNbr.parent }

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetParentYangName() string { return "bfd-mhop-vrf-nbrs" }

