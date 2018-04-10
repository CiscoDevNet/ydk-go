// This module contains a collection of YANG definitions
// for BFD neighbor monitoring.
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
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

// BfdRemoteStateType represents BFD remote state type
type BfdRemoteStateType string

const (
    BfdRemoteStateType_remote_up BfdRemoteStateType = "remote-up"

    BfdRemoteStateType_remote_down BfdRemoteStateType = "remote-down"

    BfdRemoteStateType_remote_init BfdRemoteStateType = "remote-init"

    BfdRemoteStateType_remote_admindown BfdRemoteStateType = "remote-admindown"

    BfdRemoteStateType_remote_invalid BfdRemoteStateType = "remote-invalid"
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

// BfdLspType represents BFD LSP type
type BfdLspType string

const (
    BfdLspType_working BfdLspType = "working"

    BfdLspType_protect BfdLspType = "protect"

    BfdLspType_unknown BfdLspType = "unknown"
)

// BfdState
// BFD neighbor information
type BfdState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BFD neighbor session information.
    Sessions BfdState_Sessions
}

func (bfdState *BfdState) GetEntityData() *types.CommonEntityData {
    bfdState.EntityData.YFilter = bfdState.YFilter
    bfdState.EntityData.YangName = "bfd-state"
    bfdState.EntityData.BundleName = "cisco_ios_xe"
    bfdState.EntityData.ParentYangName = "Cisco-IOS-XE-bfd-oper"
    bfdState.EntityData.SegmentPath = "Cisco-IOS-XE-bfd-oper:bfd-state"
    bfdState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdState.EntityData.Children = make(map[string]types.YChild)
    bfdState.EntityData.Children["sessions"] = types.YChild{"Sessions", &bfdState.Sessions}
    bfdState.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfdState.EntityData)
}

// BfdState_Sessions
// BFD neighbor session information
type BfdState_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD sessions. The type is slice of BfdState_Sessions_Session.
    Session []BfdState_Sessions_Session
}

func (sessions *BfdState_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xe"
    sessions.EntityData.ParentYangName = "bfd-state"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// BfdState_Sessions_Session
// List of BFD sessions
type BfdState_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session type. The type is BfdOperSessionType.
    Type_ interface{}

    // BFD tunnel path information.
    BfdTunnelPaths BfdState_Sessions_Session_BfdTunnelPaths

    // BFD circuit information.
    BfdCircuits BfdState_Sessions_Session_BfdCircuits

    // BFD neighbor information.
    BfdNbrs BfdState_Sessions_Session_BfdNbrs

    // Multi hop neighbors for multi hop neighbor scenario  for global VRF (no
    // VRF).
    BfdMhopNbrs BfdState_Sessions_Session_BfdMhopNbrs

    // Multi hop neighbors for multi hop neighbor scenario with non-global VRF.
    BfdMhopVrfNbrs BfdState_Sessions_Session_BfdMhopVrfNbrs
}

func (session *BfdState_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xe"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[type='" + fmt.Sprintf("%v", session.Type_) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["bfd-tunnel-paths"] = types.YChild{"BfdTunnelPaths", &session.BfdTunnelPaths}
    session.EntityData.Children["bfd-circuits"] = types.YChild{"BfdCircuits", &session.BfdCircuits}
    session.EntityData.Children["bfd-nbrs"] = types.YChild{"BfdNbrs", &session.BfdNbrs}
    session.EntityData.Children["bfd-mhop-nbrs"] = types.YChild{"BfdMhopNbrs", &session.BfdMhopNbrs}
    session.EntityData.Children["bfd-mhop-vrf-nbrs"] = types.YChild{"BfdMhopVrfNbrs", &session.BfdMhopVrfNbrs}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["type"] = types.YLeaf{"Type_", session.Type_}
    return &(session.EntityData)
}

// BfdState_Sessions_Session_BfdTunnelPaths
// BFD tunnel path information
type BfdState_Sessions_Session_BfdTunnelPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD tunnel paths. The type is slice of
    // BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath.
    BfdTunnelPath []BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetEntityData() *types.CommonEntityData {
    bfdTunnelPaths.EntityData.YFilter = bfdTunnelPaths.YFilter
    bfdTunnelPaths.EntityData.YangName = "bfd-tunnel-paths"
    bfdTunnelPaths.EntityData.BundleName = "cisco_ios_xe"
    bfdTunnelPaths.EntityData.ParentYangName = "session"
    bfdTunnelPaths.EntityData.SegmentPath = "bfd-tunnel-paths"
    bfdTunnelPaths.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdTunnelPaths.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdTunnelPaths.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdTunnelPaths.EntityData.Children = make(map[string]types.YChild)
    bfdTunnelPaths.EntityData.Children["bfd-tunnel-path"] = types.YChild{"BfdTunnelPath", nil}
    for i := range bfdTunnelPaths.BfdTunnelPath {
        bfdTunnelPaths.EntityData.Children[types.GetSegmentPath(&bfdTunnelPaths.BfdTunnelPath[i])] = types.YChild{"BfdTunnelPath", &bfdTunnelPaths.BfdTunnelPath[i]}
    }
    bfdTunnelPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfdTunnelPaths.EntityData)
}

// BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath
// List of BFD tunnel paths
type BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Associated interface. The type is string.
    Interface_ interface{}

    // This attribute is a key. LSP type. The type is BfdLspType.
    LspType interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'  is also mapped to up/down. The
    // type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdTunnelPath *BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath) GetEntityData() *types.CommonEntityData {
    bfdTunnelPath.EntityData.YFilter = bfdTunnelPath.YFilter
    bfdTunnelPath.EntityData.YangName = "bfd-tunnel-path"
    bfdTunnelPath.EntityData.BundleName = "cisco_ios_xe"
    bfdTunnelPath.EntityData.ParentYangName = "bfd-tunnel-paths"
    bfdTunnelPath.EntityData.SegmentPath = "bfd-tunnel-path" + "[interface='" + fmt.Sprintf("%v", bfdTunnelPath.Interface_) + "']" + "[lsp-type='" + fmt.Sprintf("%v", bfdTunnelPath.LspType) + "']"
    bfdTunnelPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdTunnelPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdTunnelPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdTunnelPath.EntityData.Children = make(map[string]types.YChild)
    bfdTunnelPath.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdTunnelPath.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", bfdTunnelPath.Interface_}
    bfdTunnelPath.EntityData.Leafs["lsp-type"] = types.YLeaf{"LspType", bfdTunnelPath.LspType}
    bfdTunnelPath.EntityData.Leafs["ld"] = types.YLeaf{"Ld", bfdTunnelPath.Ld}
    bfdTunnelPath.EntityData.Leafs["rd"] = types.YLeaf{"Rd", bfdTunnelPath.Rd}
    bfdTunnelPath.EntityData.Leafs["remote-state"] = types.YLeaf{"RemoteState", bfdTunnelPath.RemoteState}
    bfdTunnelPath.EntityData.Leafs["state"] = types.YLeaf{"State", bfdTunnelPath.State}
    return &(bfdTunnelPath.EntityData)
}

// BfdState_Sessions_Session_BfdCircuits
// BFD circuit information
type BfdState_Sessions_Session_BfdCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD circuits. The type is slice of
    // BfdState_Sessions_Session_BfdCircuits_BfdCircuit.
    BfdCircuit []BfdState_Sessions_Session_BfdCircuits_BfdCircuit
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetEntityData() *types.CommonEntityData {
    bfdCircuits.EntityData.YFilter = bfdCircuits.YFilter
    bfdCircuits.EntityData.YangName = "bfd-circuits"
    bfdCircuits.EntityData.BundleName = "cisco_ios_xe"
    bfdCircuits.EntityData.ParentYangName = "session"
    bfdCircuits.EntityData.SegmentPath = "bfd-circuits"
    bfdCircuits.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdCircuits.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdCircuits.EntityData.Children = make(map[string]types.YChild)
    bfdCircuits.EntityData.Children["bfd-circuit"] = types.YChild{"BfdCircuit", nil}
    for i := range bfdCircuits.BfdCircuit {
        bfdCircuits.EntityData.Children[types.GetSegmentPath(&bfdCircuits.BfdCircuit[i])] = types.YChild{"BfdCircuit", &bfdCircuits.BfdCircuit[i]}
    }
    bfdCircuits.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfdCircuits.EntityData)
}

// BfdState_Sessions_Session_BfdCircuits_BfdCircuit
// List of BFD circuits
type BfdState_Sessions_Session_BfdCircuits_BfdCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Associated interface. The type is string.
    Interface_ interface{}

    // This attribute is a key. Virtual circuit identifier. The type is
    // interface{} with range: 0..4294967295.
    Vcid interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'  is also mapped to up/down. The
    // type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdCircuit *BfdState_Sessions_Session_BfdCircuits_BfdCircuit) GetEntityData() *types.CommonEntityData {
    bfdCircuit.EntityData.YFilter = bfdCircuit.YFilter
    bfdCircuit.EntityData.YangName = "bfd-circuit"
    bfdCircuit.EntityData.BundleName = "cisco_ios_xe"
    bfdCircuit.EntityData.ParentYangName = "bfd-circuits"
    bfdCircuit.EntityData.SegmentPath = "bfd-circuit" + "[interface='" + fmt.Sprintf("%v", bfdCircuit.Interface_) + "']" + "[vcid='" + fmt.Sprintf("%v", bfdCircuit.Vcid) + "']"
    bfdCircuit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdCircuit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdCircuit.EntityData.Children = make(map[string]types.YChild)
    bfdCircuit.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdCircuit.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", bfdCircuit.Interface_}
    bfdCircuit.EntityData.Leafs["vcid"] = types.YLeaf{"Vcid", bfdCircuit.Vcid}
    bfdCircuit.EntityData.Leafs["ld"] = types.YLeaf{"Ld", bfdCircuit.Ld}
    bfdCircuit.EntityData.Leafs["rd"] = types.YLeaf{"Rd", bfdCircuit.Rd}
    bfdCircuit.EntityData.Leafs["remote-state"] = types.YLeaf{"RemoteState", bfdCircuit.RemoteState}
    bfdCircuit.EntityData.Leafs["state"] = types.YLeaf{"State", bfdCircuit.State}
    return &(bfdCircuit.EntityData)
}

// BfdState_Sessions_Session_BfdNbrs
// BFD neighbor information
type BfdState_Sessions_Session_BfdNbrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD neighbors. The type is slice of
    // BfdState_Sessions_Session_BfdNbrs_BfdNbr.
    BfdNbr []BfdState_Sessions_Session_BfdNbrs_BfdNbr
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetEntityData() *types.CommonEntityData {
    bfdNbrs.EntityData.YFilter = bfdNbrs.YFilter
    bfdNbrs.EntityData.YangName = "bfd-nbrs"
    bfdNbrs.EntityData.BundleName = "cisco_ios_xe"
    bfdNbrs.EntityData.ParentYangName = "session"
    bfdNbrs.EntityData.SegmentPath = "bfd-nbrs"
    bfdNbrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdNbrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdNbrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdNbrs.EntityData.Children = make(map[string]types.YChild)
    bfdNbrs.EntityData.Children["bfd-nbr"] = types.YChild{"BfdNbr", nil}
    for i := range bfdNbrs.BfdNbr {
        bfdNbrs.EntityData.Children[types.GetSegmentPath(&bfdNbrs.BfdNbr[i])] = types.YChild{"BfdNbr", &bfdNbrs.BfdNbr[i]}
    }
    bfdNbrs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfdNbrs.EntityData)
}

// BfdState_Sessions_Session_BfdNbrs_BfdNbr
// List of BFD neighbors
type BfdState_Sessions_Session_BfdNbrs_BfdNbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // This attribute is a key. Interface. The type is string.
    Interface_ interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'  is also mapped to up/down. The
    // type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdNbr *BfdState_Sessions_Session_BfdNbrs_BfdNbr) GetEntityData() *types.CommonEntityData {
    bfdNbr.EntityData.YFilter = bfdNbr.YFilter
    bfdNbr.EntityData.YangName = "bfd-nbr"
    bfdNbr.EntityData.BundleName = "cisco_ios_xe"
    bfdNbr.EntityData.ParentYangName = "bfd-nbrs"
    bfdNbr.EntityData.SegmentPath = "bfd-nbr" + "[ip='" + fmt.Sprintf("%v", bfdNbr.Ip) + "']" + "[interface='" + fmt.Sprintf("%v", bfdNbr.Interface_) + "']"
    bfdNbr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdNbr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdNbr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdNbr.EntityData.Children = make(map[string]types.YChild)
    bfdNbr.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdNbr.EntityData.Leafs["ip"] = types.YLeaf{"Ip", bfdNbr.Ip}
    bfdNbr.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", bfdNbr.Interface_}
    bfdNbr.EntityData.Leafs["ld"] = types.YLeaf{"Ld", bfdNbr.Ld}
    bfdNbr.EntityData.Leafs["rd"] = types.YLeaf{"Rd", bfdNbr.Rd}
    bfdNbr.EntityData.Leafs["remote-state"] = types.YLeaf{"RemoteState", bfdNbr.RemoteState}
    bfdNbr.EntityData.Leafs["state"] = types.YLeaf{"State", bfdNbr.State}
    return &(bfdNbr.EntityData)
}

// BfdState_Sessions_Session_BfdMhopNbrs
// Multi hop neighbors for multi hop neighbor scenario 
// for global VRF (no VRF)
type BfdState_Sessions_Session_BfdMhopNbrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of MHOP neighbors. The type is slice of
    // BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr.
    BfdMhopNbr []BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetEntityData() *types.CommonEntityData {
    bfdMhopNbrs.EntityData.YFilter = bfdMhopNbrs.YFilter
    bfdMhopNbrs.EntityData.YangName = "bfd-mhop-nbrs"
    bfdMhopNbrs.EntityData.BundleName = "cisco_ios_xe"
    bfdMhopNbrs.EntityData.ParentYangName = "session"
    bfdMhopNbrs.EntityData.SegmentPath = "bfd-mhop-nbrs"
    bfdMhopNbrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopNbrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopNbrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopNbrs.EntityData.Children = make(map[string]types.YChild)
    bfdMhopNbrs.EntityData.Children["bfd-mhop-nbr"] = types.YChild{"BfdMhopNbr", nil}
    for i := range bfdMhopNbrs.BfdMhopNbr {
        bfdMhopNbrs.EntityData.Children[types.GetSegmentPath(&bfdMhopNbrs.BfdMhopNbr[i])] = types.YChild{"BfdMhopNbr", &bfdMhopNbrs.BfdMhopNbr[i]}
    }
    bfdMhopNbrs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfdMhopNbrs.EntityData)
}

// BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr
// List of MHOP neighbors
type BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // This attribute is a key. Source IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SrcIp interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'  is also mapped to up/down. The
    // type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdMhopNbr *BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr) GetEntityData() *types.CommonEntityData {
    bfdMhopNbr.EntityData.YFilter = bfdMhopNbr.YFilter
    bfdMhopNbr.EntityData.YangName = "bfd-mhop-nbr"
    bfdMhopNbr.EntityData.BundleName = "cisco_ios_xe"
    bfdMhopNbr.EntityData.ParentYangName = "bfd-mhop-nbrs"
    bfdMhopNbr.EntityData.SegmentPath = "bfd-mhop-nbr" + "[ip='" + fmt.Sprintf("%v", bfdMhopNbr.Ip) + "']" + "[src-ip='" + fmt.Sprintf("%v", bfdMhopNbr.SrcIp) + "']"
    bfdMhopNbr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopNbr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopNbr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopNbr.EntityData.Children = make(map[string]types.YChild)
    bfdMhopNbr.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdMhopNbr.EntityData.Leafs["ip"] = types.YLeaf{"Ip", bfdMhopNbr.Ip}
    bfdMhopNbr.EntityData.Leafs["src-ip"] = types.YLeaf{"SrcIp", bfdMhopNbr.SrcIp}
    bfdMhopNbr.EntityData.Leafs["ld"] = types.YLeaf{"Ld", bfdMhopNbr.Ld}
    bfdMhopNbr.EntityData.Leafs["rd"] = types.YLeaf{"Rd", bfdMhopNbr.Rd}
    bfdMhopNbr.EntityData.Leafs["remote-state"] = types.YLeaf{"RemoteState", bfdMhopNbr.RemoteState}
    bfdMhopNbr.EntityData.Leafs["state"] = types.YLeaf{"State", bfdMhopNbr.State}
    return &(bfdMhopNbr.EntityData)
}

// BfdState_Sessions_Session_BfdMhopVrfNbrs
// Multi hop neighbors for multi hop neighbor scenario
// with non-global VRF
type BfdState_Sessions_Session_BfdMhopVrfNbrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of multi hop neighbors. The type is slice of
    // BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr.
    BfdMhopVrfNbr []BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetEntityData() *types.CommonEntityData {
    bfdMhopVrfNbrs.EntityData.YFilter = bfdMhopVrfNbrs.YFilter
    bfdMhopVrfNbrs.EntityData.YangName = "bfd-mhop-vrf-nbrs"
    bfdMhopVrfNbrs.EntityData.BundleName = "cisco_ios_xe"
    bfdMhopVrfNbrs.EntityData.ParentYangName = "session"
    bfdMhopVrfNbrs.EntityData.SegmentPath = "bfd-mhop-vrf-nbrs"
    bfdMhopVrfNbrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopVrfNbrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopVrfNbrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopVrfNbrs.EntityData.Children = make(map[string]types.YChild)
    bfdMhopVrfNbrs.EntityData.Children["bfd-mhop-vrf-nbr"] = types.YChild{"BfdMhopVrfNbr", nil}
    for i := range bfdMhopVrfNbrs.BfdMhopVrfNbr {
        bfdMhopVrfNbrs.EntityData.Children[types.GetSegmentPath(&bfdMhopVrfNbrs.BfdMhopVrfNbr[i])] = types.YChild{"BfdMhopVrfNbr", &bfdMhopVrfNbrs.BfdMhopVrfNbr[i]}
    }
    bfdMhopVrfNbrs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bfdMhopVrfNbrs.EntityData)
}

// BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr
// List of multi hop neighbors
type BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // This attribute is a key. Neighbor VFR. The type is string.
    Vrf interface{}

    // This attribute is a key. Source IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SrcIp interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Ld interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Rd interface{}

    // Remote Heard. RH state of BFD version '0'  is also mapped to up/down. The
    // type is BfdRemoteStateType.
    RemoteState interface{}

    // BFD state. The type is BfdStateType.
    State interface{}
}

func (bfdMhopVrfNbr *BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr) GetEntityData() *types.CommonEntityData {
    bfdMhopVrfNbr.EntityData.YFilter = bfdMhopVrfNbr.YFilter
    bfdMhopVrfNbr.EntityData.YangName = "bfd-mhop-vrf-nbr"
    bfdMhopVrfNbr.EntityData.BundleName = "cisco_ios_xe"
    bfdMhopVrfNbr.EntityData.ParentYangName = "bfd-mhop-vrf-nbrs"
    bfdMhopVrfNbr.EntityData.SegmentPath = "bfd-mhop-vrf-nbr" + "[ip='" + fmt.Sprintf("%v", bfdMhopVrfNbr.Ip) + "']" + "[vrf='" + fmt.Sprintf("%v", bfdMhopVrfNbr.Vrf) + "']" + "[src-ip='" + fmt.Sprintf("%v", bfdMhopVrfNbr.SrcIp) + "']"
    bfdMhopVrfNbr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopVrfNbr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopVrfNbr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopVrfNbr.EntityData.Children = make(map[string]types.YChild)
    bfdMhopVrfNbr.EntityData.Leafs = make(map[string]types.YLeaf)
    bfdMhopVrfNbr.EntityData.Leafs["ip"] = types.YLeaf{"Ip", bfdMhopVrfNbr.Ip}
    bfdMhopVrfNbr.EntityData.Leafs["vrf"] = types.YLeaf{"Vrf", bfdMhopVrfNbr.Vrf}
    bfdMhopVrfNbr.EntityData.Leafs["src-ip"] = types.YLeaf{"SrcIp", bfdMhopVrfNbr.SrcIp}
    bfdMhopVrfNbr.EntityData.Leafs["ld"] = types.YLeaf{"Ld", bfdMhopVrfNbr.Ld}
    bfdMhopVrfNbr.EntityData.Leafs["rd"] = types.YLeaf{"Rd", bfdMhopVrfNbr.Rd}
    bfdMhopVrfNbr.EntityData.Leafs["remote-state"] = types.YLeaf{"RemoteState", bfdMhopVrfNbr.RemoteState}
    bfdMhopVrfNbr.EntityData.Leafs["state"] = types.YLeaf{"State", bfdMhopVrfNbr.State}
    return &(bfdMhopVrfNbr.EntityData)
}

