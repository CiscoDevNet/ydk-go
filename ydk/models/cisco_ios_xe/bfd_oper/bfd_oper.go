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
    bfdState.EntityData.AbsolutePath = bfdState.EntityData.SegmentPath
    bfdState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdState.EntityData.Children = types.NewOrderedMap()
    bfdState.EntityData.Children.Append("sessions", types.YChild{"Sessions", &bfdState.Sessions})
    bfdState.EntityData.Leafs = types.NewOrderedMap()

    bfdState.EntityData.YListKeys = []string {}

    return &(bfdState.EntityData)
}

// BfdState_Sessions
// BFD neighbor session information
type BfdState_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD sessions. The type is slice of BfdState_Sessions_Session.
    Session []*BfdState_Sessions_Session
}

func (sessions *BfdState_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xe"
    sessions.EntityData.ParentYangName = "bfd-state"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/" + sessions.EntityData.SegmentPath
    sessions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("session", types.YChild{"Session", nil})
    for i := range sessions.Session {
        sessions.EntityData.Children.Append(types.GetSegmentPath(sessions.Session[i]), types.YChild{"Session", sessions.Session[i]})
    }
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// BfdState_Sessions_Session
// List of BFD sessions
type BfdState_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Session type. The type is BfdOperSessionType.
    Type interface{}

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
    session.EntityData.SegmentPath = "session" + types.AddKeyToken(session.Type, "type")
    session.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("bfd-tunnel-paths", types.YChild{"BfdTunnelPaths", &session.BfdTunnelPaths})
    session.EntityData.Children.Append("bfd-circuits", types.YChild{"BfdCircuits", &session.BfdCircuits})
    session.EntityData.Children.Append("bfd-nbrs", types.YChild{"BfdNbrs", &session.BfdNbrs})
    session.EntityData.Children.Append("bfd-mhop-nbrs", types.YChild{"BfdMhopNbrs", &session.BfdMhopNbrs})
    session.EntityData.Children.Append("bfd-mhop-vrf-nbrs", types.YChild{"BfdMhopVrfNbrs", &session.BfdMhopVrfNbrs})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("type", types.YLeaf{"Type", session.Type})

    session.EntityData.YListKeys = []string {"Type"}

    return &(session.EntityData)
}

// BfdState_Sessions_Session_BfdTunnelPaths
// BFD tunnel path information
type BfdState_Sessions_Session_BfdTunnelPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD tunnel paths. The type is slice of
    // BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath.
    BfdTunnelPath []*BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath
}

func (bfdTunnelPaths *BfdState_Sessions_Session_BfdTunnelPaths) GetEntityData() *types.CommonEntityData {
    bfdTunnelPaths.EntityData.YFilter = bfdTunnelPaths.YFilter
    bfdTunnelPaths.EntityData.YangName = "bfd-tunnel-paths"
    bfdTunnelPaths.EntityData.BundleName = "cisco_ios_xe"
    bfdTunnelPaths.EntityData.ParentYangName = "session"
    bfdTunnelPaths.EntityData.SegmentPath = "bfd-tunnel-paths"
    bfdTunnelPaths.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/" + bfdTunnelPaths.EntityData.SegmentPath
    bfdTunnelPaths.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdTunnelPaths.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdTunnelPaths.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdTunnelPaths.EntityData.Children = types.NewOrderedMap()
    bfdTunnelPaths.EntityData.Children.Append("bfd-tunnel-path", types.YChild{"BfdTunnelPath", nil})
    for i := range bfdTunnelPaths.BfdTunnelPath {
        bfdTunnelPaths.EntityData.Children.Append(types.GetSegmentPath(bfdTunnelPaths.BfdTunnelPath[i]), types.YChild{"BfdTunnelPath", bfdTunnelPaths.BfdTunnelPath[i]})
    }
    bfdTunnelPaths.EntityData.Leafs = types.NewOrderedMap()

    bfdTunnelPaths.EntityData.YListKeys = []string {}

    return &(bfdTunnelPaths.EntityData)
}

// BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath
// List of BFD tunnel paths
type BfdState_Sessions_Session_BfdTunnelPaths_BfdTunnelPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Associated interface. The type is string.
    Interface interface{}

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
    bfdTunnelPath.EntityData.SegmentPath = "bfd-tunnel-path" + types.AddKeyToken(bfdTunnelPath.Interface, "interface") + types.AddKeyToken(bfdTunnelPath.LspType, "lsp-type")
    bfdTunnelPath.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/bfd-tunnel-paths/" + bfdTunnelPath.EntityData.SegmentPath
    bfdTunnelPath.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdTunnelPath.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdTunnelPath.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdTunnelPath.EntityData.Children = types.NewOrderedMap()
    bfdTunnelPath.EntityData.Leafs = types.NewOrderedMap()
    bfdTunnelPath.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", bfdTunnelPath.Interface})
    bfdTunnelPath.EntityData.Leafs.Append("lsp-type", types.YLeaf{"LspType", bfdTunnelPath.LspType})
    bfdTunnelPath.EntityData.Leafs.Append("ld", types.YLeaf{"Ld", bfdTunnelPath.Ld})
    bfdTunnelPath.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", bfdTunnelPath.Rd})
    bfdTunnelPath.EntityData.Leafs.Append("remote-state", types.YLeaf{"RemoteState", bfdTunnelPath.RemoteState})
    bfdTunnelPath.EntityData.Leafs.Append("state", types.YLeaf{"State", bfdTunnelPath.State})

    bfdTunnelPath.EntityData.YListKeys = []string {"Interface", "LspType"}

    return &(bfdTunnelPath.EntityData)
}

// BfdState_Sessions_Session_BfdCircuits
// BFD circuit information
type BfdState_Sessions_Session_BfdCircuits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD circuits. The type is slice of
    // BfdState_Sessions_Session_BfdCircuits_BfdCircuit.
    BfdCircuit []*BfdState_Sessions_Session_BfdCircuits_BfdCircuit
}

func (bfdCircuits *BfdState_Sessions_Session_BfdCircuits) GetEntityData() *types.CommonEntityData {
    bfdCircuits.EntityData.YFilter = bfdCircuits.YFilter
    bfdCircuits.EntityData.YangName = "bfd-circuits"
    bfdCircuits.EntityData.BundleName = "cisco_ios_xe"
    bfdCircuits.EntityData.ParentYangName = "session"
    bfdCircuits.EntityData.SegmentPath = "bfd-circuits"
    bfdCircuits.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/" + bfdCircuits.EntityData.SegmentPath
    bfdCircuits.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdCircuits.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdCircuits.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdCircuits.EntityData.Children = types.NewOrderedMap()
    bfdCircuits.EntityData.Children.Append("bfd-circuit", types.YChild{"BfdCircuit", nil})
    for i := range bfdCircuits.BfdCircuit {
        bfdCircuits.EntityData.Children.Append(types.GetSegmentPath(bfdCircuits.BfdCircuit[i]), types.YChild{"BfdCircuit", bfdCircuits.BfdCircuit[i]})
    }
    bfdCircuits.EntityData.Leafs = types.NewOrderedMap()

    bfdCircuits.EntityData.YListKeys = []string {}

    return &(bfdCircuits.EntityData)
}

// BfdState_Sessions_Session_BfdCircuits_BfdCircuit
// List of BFD circuits
type BfdState_Sessions_Session_BfdCircuits_BfdCircuit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Associated interface. The type is string.
    Interface interface{}

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
    bfdCircuit.EntityData.SegmentPath = "bfd-circuit" + types.AddKeyToken(bfdCircuit.Interface, "interface") + types.AddKeyToken(bfdCircuit.Vcid, "vcid")
    bfdCircuit.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/bfd-circuits/" + bfdCircuit.EntityData.SegmentPath
    bfdCircuit.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdCircuit.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdCircuit.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdCircuit.EntityData.Children = types.NewOrderedMap()
    bfdCircuit.EntityData.Leafs = types.NewOrderedMap()
    bfdCircuit.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", bfdCircuit.Interface})
    bfdCircuit.EntityData.Leafs.Append("vcid", types.YLeaf{"Vcid", bfdCircuit.Vcid})
    bfdCircuit.EntityData.Leafs.Append("ld", types.YLeaf{"Ld", bfdCircuit.Ld})
    bfdCircuit.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", bfdCircuit.Rd})
    bfdCircuit.EntityData.Leafs.Append("remote-state", types.YLeaf{"RemoteState", bfdCircuit.RemoteState})
    bfdCircuit.EntityData.Leafs.Append("state", types.YLeaf{"State", bfdCircuit.State})

    bfdCircuit.EntityData.YListKeys = []string {"Interface", "Vcid"}

    return &(bfdCircuit.EntityData)
}

// BfdState_Sessions_Session_BfdNbrs
// BFD neighbor information
type BfdState_Sessions_Session_BfdNbrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of BFD neighbors. The type is slice of
    // BfdState_Sessions_Session_BfdNbrs_BfdNbr.
    BfdNbr []*BfdState_Sessions_Session_BfdNbrs_BfdNbr
}

func (bfdNbrs *BfdState_Sessions_Session_BfdNbrs) GetEntityData() *types.CommonEntityData {
    bfdNbrs.EntityData.YFilter = bfdNbrs.YFilter
    bfdNbrs.EntityData.YangName = "bfd-nbrs"
    bfdNbrs.EntityData.BundleName = "cisco_ios_xe"
    bfdNbrs.EntityData.ParentYangName = "session"
    bfdNbrs.EntityData.SegmentPath = "bfd-nbrs"
    bfdNbrs.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/" + bfdNbrs.EntityData.SegmentPath
    bfdNbrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdNbrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdNbrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdNbrs.EntityData.Children = types.NewOrderedMap()
    bfdNbrs.EntityData.Children.Append("bfd-nbr", types.YChild{"BfdNbr", nil})
    for i := range bfdNbrs.BfdNbr {
        bfdNbrs.EntityData.Children.Append(types.GetSegmentPath(bfdNbrs.BfdNbr[i]), types.YChild{"BfdNbr", bfdNbrs.BfdNbr[i]})
    }
    bfdNbrs.EntityData.Leafs = types.NewOrderedMap()

    bfdNbrs.EntityData.YListKeys = []string {}

    return &(bfdNbrs.EntityData)
}

// BfdState_Sessions_Session_BfdNbrs_BfdNbr
// List of BFD neighbors
type BfdState_Sessions_Session_BfdNbrs_BfdNbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // This attribute is a key. Interface. The type is string.
    Interface interface{}

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
    bfdNbr.EntityData.SegmentPath = "bfd-nbr" + types.AddKeyToken(bfdNbr.Ip, "ip") + types.AddKeyToken(bfdNbr.Interface, "interface")
    bfdNbr.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/bfd-nbrs/" + bfdNbr.EntityData.SegmentPath
    bfdNbr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdNbr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdNbr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdNbr.EntityData.Children = types.NewOrderedMap()
    bfdNbr.EntityData.Leafs = types.NewOrderedMap()
    bfdNbr.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", bfdNbr.Ip})
    bfdNbr.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", bfdNbr.Interface})
    bfdNbr.EntityData.Leafs.Append("ld", types.YLeaf{"Ld", bfdNbr.Ld})
    bfdNbr.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", bfdNbr.Rd})
    bfdNbr.EntityData.Leafs.Append("remote-state", types.YLeaf{"RemoteState", bfdNbr.RemoteState})
    bfdNbr.EntityData.Leafs.Append("state", types.YLeaf{"State", bfdNbr.State})

    bfdNbr.EntityData.YListKeys = []string {"Ip", "Interface"}

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
    BfdMhopNbr []*BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr
}

func (bfdMhopNbrs *BfdState_Sessions_Session_BfdMhopNbrs) GetEntityData() *types.CommonEntityData {
    bfdMhopNbrs.EntityData.YFilter = bfdMhopNbrs.YFilter
    bfdMhopNbrs.EntityData.YangName = "bfd-mhop-nbrs"
    bfdMhopNbrs.EntityData.BundleName = "cisco_ios_xe"
    bfdMhopNbrs.EntityData.ParentYangName = "session"
    bfdMhopNbrs.EntityData.SegmentPath = "bfd-mhop-nbrs"
    bfdMhopNbrs.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/" + bfdMhopNbrs.EntityData.SegmentPath
    bfdMhopNbrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopNbrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopNbrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopNbrs.EntityData.Children = types.NewOrderedMap()
    bfdMhopNbrs.EntityData.Children.Append("bfd-mhop-nbr", types.YChild{"BfdMhopNbr", nil})
    for i := range bfdMhopNbrs.BfdMhopNbr {
        bfdMhopNbrs.EntityData.Children.Append(types.GetSegmentPath(bfdMhopNbrs.BfdMhopNbr[i]), types.YChild{"BfdMhopNbr", bfdMhopNbrs.BfdMhopNbr[i]})
    }
    bfdMhopNbrs.EntityData.Leafs = types.NewOrderedMap()

    bfdMhopNbrs.EntityData.YListKeys = []string {}

    return &(bfdMhopNbrs.EntityData)
}

// BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr
// List of MHOP neighbors
type BfdState_Sessions_Session_BfdMhopNbrs_BfdMhopNbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    bfdMhopNbr.EntityData.SegmentPath = "bfd-mhop-nbr" + types.AddKeyToken(bfdMhopNbr.Ip, "ip") + types.AddKeyToken(bfdMhopNbr.SrcIp, "src-ip")
    bfdMhopNbr.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/bfd-mhop-nbrs/" + bfdMhopNbr.EntityData.SegmentPath
    bfdMhopNbr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopNbr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopNbr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopNbr.EntityData.Children = types.NewOrderedMap()
    bfdMhopNbr.EntityData.Leafs = types.NewOrderedMap()
    bfdMhopNbr.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", bfdMhopNbr.Ip})
    bfdMhopNbr.EntityData.Leafs.Append("src-ip", types.YLeaf{"SrcIp", bfdMhopNbr.SrcIp})
    bfdMhopNbr.EntityData.Leafs.Append("ld", types.YLeaf{"Ld", bfdMhopNbr.Ld})
    bfdMhopNbr.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", bfdMhopNbr.Rd})
    bfdMhopNbr.EntityData.Leafs.Append("remote-state", types.YLeaf{"RemoteState", bfdMhopNbr.RemoteState})
    bfdMhopNbr.EntityData.Leafs.Append("state", types.YLeaf{"State", bfdMhopNbr.State})

    bfdMhopNbr.EntityData.YListKeys = []string {"Ip", "SrcIp"}

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
    BfdMhopVrfNbr []*BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr
}

func (bfdMhopVrfNbrs *BfdState_Sessions_Session_BfdMhopVrfNbrs) GetEntityData() *types.CommonEntityData {
    bfdMhopVrfNbrs.EntityData.YFilter = bfdMhopVrfNbrs.YFilter
    bfdMhopVrfNbrs.EntityData.YangName = "bfd-mhop-vrf-nbrs"
    bfdMhopVrfNbrs.EntityData.BundleName = "cisco_ios_xe"
    bfdMhopVrfNbrs.EntityData.ParentYangName = "session"
    bfdMhopVrfNbrs.EntityData.SegmentPath = "bfd-mhop-vrf-nbrs"
    bfdMhopVrfNbrs.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/" + bfdMhopVrfNbrs.EntityData.SegmentPath
    bfdMhopVrfNbrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopVrfNbrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopVrfNbrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopVrfNbrs.EntityData.Children = types.NewOrderedMap()
    bfdMhopVrfNbrs.EntityData.Children.Append("bfd-mhop-vrf-nbr", types.YChild{"BfdMhopVrfNbr", nil})
    for i := range bfdMhopVrfNbrs.BfdMhopVrfNbr {
        bfdMhopVrfNbrs.EntityData.Children.Append(types.GetSegmentPath(bfdMhopVrfNbrs.BfdMhopVrfNbr[i]), types.YChild{"BfdMhopVrfNbr", bfdMhopVrfNbrs.BfdMhopVrfNbr[i]})
    }
    bfdMhopVrfNbrs.EntityData.Leafs = types.NewOrderedMap()

    bfdMhopVrfNbrs.EntityData.YListKeys = []string {}

    return &(bfdMhopVrfNbrs.EntityData)
}

// BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr
// List of multi hop neighbors
type BfdState_Sessions_Session_BfdMhopVrfNbrs_BfdMhopVrfNbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    bfdMhopVrfNbr.EntityData.SegmentPath = "bfd-mhop-vrf-nbr" + types.AddKeyToken(bfdMhopVrfNbr.Ip, "ip") + types.AddKeyToken(bfdMhopVrfNbr.Vrf, "vrf") + types.AddKeyToken(bfdMhopVrfNbr.SrcIp, "src-ip")
    bfdMhopVrfNbr.EntityData.AbsolutePath = "Cisco-IOS-XE-bfd-oper:bfd-state/sessions/session/bfd-mhop-vrf-nbrs/" + bfdMhopVrfNbr.EntityData.SegmentPath
    bfdMhopVrfNbr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bfdMhopVrfNbr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bfdMhopVrfNbr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bfdMhopVrfNbr.EntityData.Children = types.NewOrderedMap()
    bfdMhopVrfNbr.EntityData.Leafs = types.NewOrderedMap()
    bfdMhopVrfNbr.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", bfdMhopVrfNbr.Ip})
    bfdMhopVrfNbr.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", bfdMhopVrfNbr.Vrf})
    bfdMhopVrfNbr.EntityData.Leafs.Append("src-ip", types.YLeaf{"SrcIp", bfdMhopVrfNbr.SrcIp})
    bfdMhopVrfNbr.EntityData.Leafs.Append("ld", types.YLeaf{"Ld", bfdMhopVrfNbr.Ld})
    bfdMhopVrfNbr.EntityData.Leafs.Append("rd", types.YLeaf{"Rd", bfdMhopVrfNbr.Rd})
    bfdMhopVrfNbr.EntityData.Leafs.Append("remote-state", types.YLeaf{"RemoteState", bfdMhopVrfNbr.RemoteState})
    bfdMhopVrfNbr.EntityData.Leafs.Append("state", types.YLeaf{"State", bfdMhopVrfNbr.State})

    bfdMhopVrfNbr.EntityData.YListKeys = []string {"Ip", "Vrf", "SrcIp"}

    return &(bfdMhopVrfNbr.EntityData)
}

