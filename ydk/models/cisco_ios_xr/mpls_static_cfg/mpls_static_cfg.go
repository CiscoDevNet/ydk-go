// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-static package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-static: MPLS Static Configuration Data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_static_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_static_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-static-cfg mpls-static}", reflect.TypeOf(MplsStatic{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-static-cfg:mpls-static", reflect.TypeOf(MplsStatic{}))
}

// MplsStaticPath represents Mpls static path
type MplsStaticPath string

const (
    // Pop and Lookup
    MplsStaticPath_pop_and_lookup MplsStaticPath = "pop-and-lookup"

    // Crossconnect
    MplsStaticPath_cross_connect MplsStaticPath = "cross-connect"
)

// MplsStaticAddressFamily represents Mpls static address family
type MplsStaticAddressFamily string

const (
    // IPv4 Unicast
    MplsStaticAddressFamily_ipv4_unicast MplsStaticAddressFamily = "ipv4-unicast"
)

// MplsStaticLabelMode represents Mpls static label mode
type MplsStaticLabelMode string

const (
    // Per VRF
    MplsStaticLabelMode_per_vrf MplsStaticLabelMode = "per-vrf"

    // Per Prefix
    MplsStaticLabelMode_per_prefix MplsStaticLabelMode = "per-prefix"

    // Cross connect
    MplsStaticLabelMode_lsp MplsStaticLabelMode = "lsp"
)

// MplsStaticPathRole represents Mpls static path role
type MplsStaticPathRole string

const (
    // Path is only for primary traffic
    MplsStaticPathRole_primary MplsStaticPathRole = "primary"

    // Path is only for backup traffic
    MplsStaticPathRole_backup MplsStaticPathRole = "backup"

    // Path is for primary and backup traffic
    MplsStaticPathRole_primary_backup MplsStaticPathRole = "primary-backup"
)

// MplsStaticNhAddressFamily represents Mpls static nh address family
type MplsStaticNhAddressFamily string

const (
    // IPv4 Next Hop
    MplsStaticNhAddressFamily_ipv4 MplsStaticNhAddressFamily = "ipv4"

    // IPv6 Next Hop
    MplsStaticNhAddressFamily_ipv6 MplsStaticNhAddressFamily = "ipv6"
)

// MplsStaticNhMode represents Mpls static nh mode
type MplsStaticNhMode string

const (
    // Explicitly configured next hop path
    MplsStaticNhMode_configured MplsStaticNhMode = "configured"

    // Resolvable next hop which will result in a path
    // set
    MplsStaticNhMode_resolve MplsStaticNhMode = "resolve"
)

// MplsStaticOutLabelTypes represents Mpls static out label types
type MplsStaticOutLabelTypes string

const (
    // None
    MplsStaticOutLabelTypes_none MplsStaticOutLabelTypes = "none"

    // OutLabel
    MplsStaticOutLabelTypes_out_label MplsStaticOutLabelTypes = "out-label"

    // Pop
    MplsStaticOutLabelTypes_pop MplsStaticOutLabelTypes = "pop"

    // IPv4 explicit-null
    MplsStaticOutLabelTypes_exp_null MplsStaticOutLabelTypes = "exp-null"

    // IPv6 explicit-null
    MplsStaticOutLabelTypes_ipv6_explicit_null MplsStaticOutLabelTypes = "ipv6-explicit-null"
)

// MplsStatic
// MPLS Static Configuration Data
type MplsStatic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // VRF table.
    Vrfs MplsStatic_Vrfs

    // MPLS Static Interface Table.
    Interfaces MplsStatic_Interfaces

    // Default VRF.
    DefaultVrf MplsStatic_DefaultVrf
}

func (mplsStatic *MplsStatic) GetEntityData() *types.CommonEntityData {
    mplsStatic.EntityData.YFilter = mplsStatic.YFilter
    mplsStatic.EntityData.YangName = "mpls-static"
    mplsStatic.EntityData.BundleName = "cisco_ios_xr"
    mplsStatic.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-static-cfg"
    mplsStatic.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-static-cfg:mpls-static"
    mplsStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsStatic.EntityData.Children = make(map[string]types.YChild)
    mplsStatic.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &mplsStatic.Vrfs}
    mplsStatic.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &mplsStatic.Interfaces}
    mplsStatic.EntityData.Children["default-vrf"] = types.YChild{"DefaultVrf", &mplsStatic.DefaultVrf}
    mplsStatic.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsStatic.EntityData.Leafs["enable"] = types.YLeaf{"Enable", mplsStatic.Enable}
    return &(mplsStatic.EntityData)
}

// MplsStatic_Vrfs
// VRF table
type MplsStatic_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of MplsStatic_Vrfs_Vrf.
    Vrf []MplsStatic_Vrfs_Vrf
}

func (vrfs *MplsStatic_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "mpls-static"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// MplsStatic_Vrfs_Vrf
// VRF Name
type MplsStatic_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Table of the Label Switched Paths.
    LabelSwitchedPaths MplsStatic_Vrfs_Vrf_LabelSwitchedPaths

    // Address Family Table.
    Afs MplsStatic_Vrfs_Vrf_Afs
}

func (vrf *MplsStatic_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["label-switched-paths"] = types.YChild{"LabelSwitchedPaths", &vrf.LabelSwitchedPaths}
    vrf.EntityData.Children["afs"] = types.YChild{"Afs", &vrf.Afs}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    vrf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", vrf.Enable}
    return &(vrf.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths
// Table of the Label Switched Paths
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label Switched Path. The type is slice of
    // MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetEntityData() *types.CommonEntityData {
    labelSwitchedPaths.EntityData.YFilter = labelSwitchedPaths.YFilter
    labelSwitchedPaths.EntityData.YangName = "label-switched-paths"
    labelSwitchedPaths.EntityData.BundleName = "cisco_ios_xr"
    labelSwitchedPaths.EntityData.ParentYangName = "vrf"
    labelSwitchedPaths.EntityData.SegmentPath = "label-switched-paths"
    labelSwitchedPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSwitchedPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSwitchedPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSwitchedPaths.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPaths.EntityData.Children["label-switched-path"] = types.YChild{"LabelSwitchedPath", nil}
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        labelSwitchedPaths.EntityData.Children[types.GetSegmentPath(&labelSwitchedPaths.LabelSwitchedPath[i])] = types.YChild{"LabelSwitchedPath", &labelSwitchedPaths.LabelSwitchedPath[i]}
    }
    labelSwitchedPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(labelSwitchedPaths.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath
// Label Switched Path
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSP Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    LspName interface{}

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Backup Path Parameters.
    BackupPaths MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths

    // MPLS Static Local Label Value.
    InLabel MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel

    // Forward Path Parameters.
    Paths MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths
}

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetEntityData() *types.CommonEntityData {
    labelSwitchedPath.EntityData.YFilter = labelSwitchedPath.YFilter
    labelSwitchedPath.EntityData.YangName = "label-switched-path"
    labelSwitchedPath.EntityData.BundleName = "cisco_ios_xr"
    labelSwitchedPath.EntityData.ParentYangName = "label-switched-paths"
    labelSwitchedPath.EntityData.SegmentPath = "label-switched-path" + "[lsp-name='" + fmt.Sprintf("%v", labelSwitchedPath.LspName) + "']"
    labelSwitchedPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSwitchedPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSwitchedPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSwitchedPath.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPath.EntityData.Children["backup-paths"] = types.YChild{"BackupPaths", &labelSwitchedPath.BackupPaths}
    labelSwitchedPath.EntityData.Children["in-label"] = types.YChild{"InLabel", &labelSwitchedPath.InLabel}
    labelSwitchedPath.EntityData.Children["paths"] = types.YChild{"Paths", &labelSwitchedPath.Paths}
    labelSwitchedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    labelSwitchedPath.EntityData.Leafs["lsp-name"] = types.YLeaf{"LspName", labelSwitchedPath.LspName}
    labelSwitchedPath.EntityData.Leafs["enable"] = types.YLeaf{"Enable", labelSwitchedPath.Enable}
    return &(labelSwitchedPath.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths
// Backup Path Parameters
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path.
    Path []MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetEntityData() *types.CommonEntityData {
    backupPaths.EntityData.YFilter = backupPaths.YFilter
    backupPaths.EntityData.YangName = "backup-paths"
    backupPaths.EntityData.BundleName = "cisco_ios_xr"
    backupPaths.EntityData.ParentYangName = "label-switched-path"
    backupPaths.EntityData.SegmentPath = "backup-paths"
    backupPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPaths.EntityData.Children = make(map[string]types.YChild)
    backupPaths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range backupPaths.Path {
        backupPaths.EntityData.Children[types.GetSegmentPath(&backupPaths.Path[i])] = types.YChild{"Path", &backupPaths.Path[i]}
    }
    backupPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backupPaths.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "backup-paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel
// MPLS Static Local Label Value
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Label. The type is interface{} with range: 16..1048575.
    InLabelValue interface{}

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}

    // Top Label Hashing Mode. The type is bool.
    TlhMode interface{}
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetEntityData() *types.CommonEntityData {
    inLabel.EntityData.YFilter = inLabel.YFilter
    inLabel.EntityData.YangName = "in-label"
    inLabel.EntityData.BundleName = "cisco_ios_xr"
    inLabel.EntityData.ParentYangName = "label-switched-path"
    inLabel.EntityData.SegmentPath = "in-label"
    inLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inLabel.EntityData.Children = make(map[string]types.YChild)
    inLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    inLabel.EntityData.Leafs["in-label-value"] = types.YLeaf{"InLabelValue", inLabel.InLabelValue}
    inLabel.EntityData.Leafs["label-mode"] = types.YLeaf{"LabelMode", inLabel.LabelMode}
    inLabel.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", inLabel.Prefix}
    inLabel.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", inLabel.PrefixLength}
    inLabel.EntityData.Leafs["tlh-mode"] = types.YLeaf{"TlhMode", inLabel.TlhMode}
    return &(inLabel.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths
// Forward Path Parameters
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path.
    Path []MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "label-switched-path"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs
// Address Family Table
type MplsStatic_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family. The type is slice of MplsStatic_Vrfs_Vrf_Afs_Af.
    Af []MplsStatic_Vrfs_Vrf_Afs_Af
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = make(map[string]types.YChild)
    afs.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range afs.Af {
        afs.EntityData.Children[types.GetSegmentPath(&afs.Af[i])] = types.YChild{"Af", &afs.Af[i]}
    }
    afs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afs.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af
// Address Family
type MplsStatic_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is
    // MplsStaticAddressFamily.
    Afi interface{}

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Top Label Hash.
    TopLabelHash MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash

    // Local Label.
    LocalLabels MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels
}

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + "[afi='" + fmt.Sprintf("%v", af.Afi) + "']"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["top-label-hash"] = types.YChild{"TopLabelHash", &af.TopLabelHash}
    af.EntityData.Children["local-labels"] = types.YChild{"LocalLabels", &af.LocalLabels}
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["afi"] = types.YLeaf{"Afi", af.Afi}
    af.EntityData.Leafs["enable"] = types.YLeaf{"Enable", af.Enable}
    return &(af.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash
// Top Label Hash
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Label.
    LocalLabels MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetEntityData() *types.CommonEntityData {
    topLabelHash.EntityData.YFilter = topLabelHash.YFilter
    topLabelHash.EntityData.YangName = "top-label-hash"
    topLabelHash.EntityData.BundleName = "cisco_ios_xr"
    topLabelHash.EntityData.ParentYangName = "af"
    topLabelHash.EntityData.SegmentPath = "top-label-hash"
    topLabelHash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topLabelHash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topLabelHash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topLabelHash.EntityData.Children = make(map[string]types.YChild)
    topLabelHash.EntityData.Children["local-labels"] = types.YChild{"LocalLabels", &topLabelHash.LocalLabels}
    topLabelHash.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(topLabelHash.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels
// Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetEntityData() *types.CommonEntityData {
    localLabels.EntityData.YFilter = localLabels.YFilter
    localLabels.EntityData.YangName = "local-labels"
    localLabels.EntityData.BundleName = "cisco_ios_xr"
    localLabels.EntityData.ParentYangName = "top-label-hash"
    localLabels.EntityData.SegmentPath = "local-labels"
    localLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabels.EntityData.Children = make(map[string]types.YChild)
    localLabels.EntityData.Children["local-label"] = types.YChild{"LocalLabel", nil}
    for i := range localLabels.LocalLabel {
        localLabels.EntityData.Children[types.GetSegmentPath(&localLabels.LocalLabel[i])] = types.YChild{"LocalLabel", &localLabels.LocalLabel[i]}
    }
    localLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(localLabels.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetEntityData() *types.CommonEntityData {
    localLabel.EntityData.YFilter = localLabel.YFilter
    localLabel.EntityData.YangName = "local-label"
    localLabel.EntityData.BundleName = "cisco_ios_xr"
    localLabel.EntityData.ParentYangName = "local-labels"
    localLabel.EntityData.SegmentPath = "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
    localLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabel.EntityData.Children = make(map[string]types.YChild)
    localLabel.EntityData.Children["label-type"] = types.YChild{"LabelType", &localLabel.LabelType}
    localLabel.EntityData.Children["paths"] = types.YChild{"Paths", &localLabel.Paths}
    localLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    localLabel.EntityData.Leafs["local-label-id"] = types.YLeaf{"LocalLabelId", localLabel.LocalLabelId}
    return &(localLabel.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetEntityData() *types.CommonEntityData {
    labelType.EntityData.YFilter = labelType.YFilter
    labelType.EntityData.YangName = "label-type"
    labelType.EntityData.BundleName = "cisco_ios_xr"
    labelType.EntityData.ParentYangName = "local-label"
    labelType.EntityData.SegmentPath = "label-type"
    labelType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelType.EntityData.Children = make(map[string]types.YChild)
    labelType.EntityData.Leafs = make(map[string]types.YLeaf)
    labelType.EntityData.Leafs["label-mode"] = types.YLeaf{"LabelMode", labelType.LabelMode}
    labelType.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", labelType.Prefix}
    labelType.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", labelType.PrefixLength}
    return &(labelType.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "local-label"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels
// Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetEntityData() *types.CommonEntityData {
    localLabels.EntityData.YFilter = localLabels.YFilter
    localLabels.EntityData.YangName = "local-labels"
    localLabels.EntityData.BundleName = "cisco_ios_xr"
    localLabels.EntityData.ParentYangName = "af"
    localLabels.EntityData.SegmentPath = "local-labels"
    localLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabels.EntityData.Children = make(map[string]types.YChild)
    localLabels.EntityData.Children["local-label"] = types.YChild{"LocalLabel", nil}
    for i := range localLabels.LocalLabel {
        localLabels.EntityData.Children[types.GetSegmentPath(&localLabels.LocalLabel[i])] = types.YChild{"LocalLabel", &localLabels.LocalLabel[i]}
    }
    localLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(localLabels.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetEntityData() *types.CommonEntityData {
    localLabel.EntityData.YFilter = localLabel.YFilter
    localLabel.EntityData.YangName = "local-label"
    localLabel.EntityData.BundleName = "cisco_ios_xr"
    localLabel.EntityData.ParentYangName = "local-labels"
    localLabel.EntityData.SegmentPath = "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
    localLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabel.EntityData.Children = make(map[string]types.YChild)
    localLabel.EntityData.Children["label-type"] = types.YChild{"LabelType", &localLabel.LabelType}
    localLabel.EntityData.Children["paths"] = types.YChild{"Paths", &localLabel.Paths}
    localLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    localLabel.EntityData.Leafs["local-label-id"] = types.YLeaf{"LocalLabelId", localLabel.LocalLabelId}
    return &(localLabel.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetEntityData() *types.CommonEntityData {
    labelType.EntityData.YFilter = labelType.YFilter
    labelType.EntityData.YangName = "label-type"
    labelType.EntityData.BundleName = "cisco_ios_xr"
    labelType.EntityData.ParentYangName = "local-label"
    labelType.EntityData.SegmentPath = "label-type"
    labelType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelType.EntityData.Children = make(map[string]types.YChild)
    labelType.EntityData.Leafs = make(map[string]types.YLeaf)
    labelType.EntityData.Leafs["label-mode"] = types.YLeaf{"LabelMode", labelType.LabelMode}
    labelType.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", labelType.Prefix}
    labelType.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", labelType.PrefixLength}
    return &(labelType.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "local-label"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_Interfaces
// MPLS Static Interface Table
type MplsStatic_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static Interface Enable. The type is slice of
    // MplsStatic_Interfaces_Interface_.
    Interface_ []MplsStatic_Interfaces_Interface
}

func (interfaces *MplsStatic_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "mpls-static"
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

// MplsStatic_Interfaces_Interface
// MPLS Static Interface Enable
type MplsStatic_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of Interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (self *MplsStatic_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    return &(self.EntityData)
}

// MplsStatic_DefaultVrf
// Default VRF
type MplsStatic_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Table of the Label Switched Paths.
    LabelSwitchedPaths MplsStatic_DefaultVrf_LabelSwitchedPaths

    // Address Family Table.
    Afs MplsStatic_DefaultVrf_Afs
}

func (defaultVrf *MplsStatic_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "mpls-static"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = make(map[string]types.YChild)
    defaultVrf.EntityData.Children["label-switched-paths"] = types.YChild{"LabelSwitchedPaths", &defaultVrf.LabelSwitchedPaths}
    defaultVrf.EntityData.Children["afs"] = types.YChild{"Afs", &defaultVrf.Afs}
    defaultVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultVrf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", defaultVrf.Enable}
    return &(defaultVrf.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths
// Table of the Label Switched Paths
type MplsStatic_DefaultVrf_LabelSwitchedPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label Switched Path. The type is slice of
    // MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetEntityData() *types.CommonEntityData {
    labelSwitchedPaths.EntityData.YFilter = labelSwitchedPaths.YFilter
    labelSwitchedPaths.EntityData.YangName = "label-switched-paths"
    labelSwitchedPaths.EntityData.BundleName = "cisco_ios_xr"
    labelSwitchedPaths.EntityData.ParentYangName = "default-vrf"
    labelSwitchedPaths.EntityData.SegmentPath = "label-switched-paths"
    labelSwitchedPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSwitchedPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSwitchedPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSwitchedPaths.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPaths.EntityData.Children["label-switched-path"] = types.YChild{"LabelSwitchedPath", nil}
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        labelSwitchedPaths.EntityData.Children[types.GetSegmentPath(&labelSwitchedPaths.LabelSwitchedPath[i])] = types.YChild{"LabelSwitchedPath", &labelSwitchedPaths.LabelSwitchedPath[i]}
    }
    labelSwitchedPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(labelSwitchedPaths.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath
// Label Switched Path
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. LSP Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    LspName interface{}

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Backup Path Parameters.
    BackupPaths MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths

    // MPLS Static Local Label Value.
    InLabel MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel

    // Forward Path Parameters.
    Paths MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths
}

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetEntityData() *types.CommonEntityData {
    labelSwitchedPath.EntityData.YFilter = labelSwitchedPath.YFilter
    labelSwitchedPath.EntityData.YangName = "label-switched-path"
    labelSwitchedPath.EntityData.BundleName = "cisco_ios_xr"
    labelSwitchedPath.EntityData.ParentYangName = "label-switched-paths"
    labelSwitchedPath.EntityData.SegmentPath = "label-switched-path" + "[lsp-name='" + fmt.Sprintf("%v", labelSwitchedPath.LspName) + "']"
    labelSwitchedPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSwitchedPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSwitchedPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSwitchedPath.EntityData.Children = make(map[string]types.YChild)
    labelSwitchedPath.EntityData.Children["backup-paths"] = types.YChild{"BackupPaths", &labelSwitchedPath.BackupPaths}
    labelSwitchedPath.EntityData.Children["in-label"] = types.YChild{"InLabel", &labelSwitchedPath.InLabel}
    labelSwitchedPath.EntityData.Children["paths"] = types.YChild{"Paths", &labelSwitchedPath.Paths}
    labelSwitchedPath.EntityData.Leafs = make(map[string]types.YLeaf)
    labelSwitchedPath.EntityData.Leafs["lsp-name"] = types.YLeaf{"LspName", labelSwitchedPath.LspName}
    labelSwitchedPath.EntityData.Leafs["enable"] = types.YLeaf{"Enable", labelSwitchedPath.Enable}
    return &(labelSwitchedPath.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths
// Backup Path Parameters
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path.
    Path []MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetEntityData() *types.CommonEntityData {
    backupPaths.EntityData.YFilter = backupPaths.YFilter
    backupPaths.EntityData.YangName = "backup-paths"
    backupPaths.EntityData.BundleName = "cisco_ios_xr"
    backupPaths.EntityData.ParentYangName = "label-switched-path"
    backupPaths.EntityData.SegmentPath = "backup-paths"
    backupPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPaths.EntityData.Children = make(map[string]types.YChild)
    backupPaths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range backupPaths.Path {
        backupPaths.EntityData.Children[types.GetSegmentPath(&backupPaths.Path[i])] = types.YChild{"Path", &backupPaths.Path[i]}
    }
    backupPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(backupPaths.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
// Path Information
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "backup-paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel
// MPLS Static Local Label Value
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Label. The type is interface{} with range: 16..1048575.
    InLabelValue interface{}

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}

    // Top Label Hashing Mode. The type is bool.
    TlhMode interface{}
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetEntityData() *types.CommonEntityData {
    inLabel.EntityData.YFilter = inLabel.YFilter
    inLabel.EntityData.YangName = "in-label"
    inLabel.EntityData.BundleName = "cisco_ios_xr"
    inLabel.EntityData.ParentYangName = "label-switched-path"
    inLabel.EntityData.SegmentPath = "in-label"
    inLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inLabel.EntityData.Children = make(map[string]types.YChild)
    inLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    inLabel.EntityData.Leafs["in-label-value"] = types.YLeaf{"InLabelValue", inLabel.InLabelValue}
    inLabel.EntityData.Leafs["label-mode"] = types.YLeaf{"LabelMode", inLabel.LabelMode}
    inLabel.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", inLabel.Prefix}
    inLabel.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", inLabel.PrefixLength}
    inLabel.EntityData.Leafs["tlh-mode"] = types.YLeaf{"TlhMode", inLabel.TlhMode}
    return &(inLabel.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths
// Forward Path Parameters
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path.
    Path []MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "label-switched-path"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
// Path Information
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_DefaultVrf_Afs
// Address Family Table
type MplsStatic_DefaultVrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family. The type is slice of MplsStatic_DefaultVrf_Afs_Af.
    Af []MplsStatic_DefaultVrf_Afs_Af
}

func (afs *MplsStatic_DefaultVrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xr"
    afs.EntityData.ParentYangName = "default-vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    afs.EntityData.Children = make(map[string]types.YChild)
    afs.EntityData.Children["af"] = types.YChild{"Af", nil}
    for i := range afs.Af {
        afs.EntityData.Children[types.GetSegmentPath(&afs.Af[i])] = types.YChild{"Af", &afs.Af[i]}
    }
    afs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(afs.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af
// Address Family
type MplsStatic_DefaultVrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is
    // MplsStaticAddressFamily.
    Afi interface{}

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Top Label Hash.
    TopLabelHash MplsStatic_DefaultVrf_Afs_Af_TopLabelHash

    // Local Label.
    LocalLabels MplsStatic_DefaultVrf_Afs_Af_LocalLabels
}

func (af *MplsStatic_DefaultVrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + "[afi='" + fmt.Sprintf("%v", af.Afi) + "']"
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = make(map[string]types.YChild)
    af.EntityData.Children["top-label-hash"] = types.YChild{"TopLabelHash", &af.TopLabelHash}
    af.EntityData.Children["local-labels"] = types.YChild{"LocalLabels", &af.LocalLabels}
    af.EntityData.Leafs = make(map[string]types.YLeaf)
    af.EntityData.Leafs["afi"] = types.YLeaf{"Afi", af.Afi}
    af.EntityData.Leafs["enable"] = types.YLeaf{"Enable", af.Enable}
    return &(af.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash
// Top Label Hash
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Label.
    LocalLabels MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetEntityData() *types.CommonEntityData {
    topLabelHash.EntityData.YFilter = topLabelHash.YFilter
    topLabelHash.EntityData.YangName = "top-label-hash"
    topLabelHash.EntityData.BundleName = "cisco_ios_xr"
    topLabelHash.EntityData.ParentYangName = "af"
    topLabelHash.EntityData.SegmentPath = "top-label-hash"
    topLabelHash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topLabelHash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topLabelHash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topLabelHash.EntityData.Children = make(map[string]types.YChild)
    topLabelHash.EntityData.Children["local-labels"] = types.YChild{"LocalLabels", &topLabelHash.LocalLabels}
    topLabelHash.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(topLabelHash.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels
// Local Label
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetEntityData() *types.CommonEntityData {
    localLabels.EntityData.YFilter = localLabels.YFilter
    localLabels.EntityData.YangName = "local-labels"
    localLabels.EntityData.BundleName = "cisco_ios_xr"
    localLabels.EntityData.ParentYangName = "top-label-hash"
    localLabels.EntityData.SegmentPath = "local-labels"
    localLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabels.EntityData.Children = make(map[string]types.YChild)
    localLabels.EntityData.Children["local-label"] = types.YChild{"LocalLabel", nil}
    for i := range localLabels.LocalLabel {
        localLabels.EntityData.Children[types.GetSegmentPath(&localLabels.LocalLabel[i])] = types.YChild{"LocalLabel", &localLabels.LocalLabel[i]}
    }
    localLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(localLabels.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetEntityData() *types.CommonEntityData {
    localLabel.EntityData.YFilter = localLabel.YFilter
    localLabel.EntityData.YangName = "local-label"
    localLabel.EntityData.BundleName = "cisco_ios_xr"
    localLabel.EntityData.ParentYangName = "local-labels"
    localLabel.EntityData.SegmentPath = "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
    localLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabel.EntityData.Children = make(map[string]types.YChild)
    localLabel.EntityData.Children["label-type"] = types.YChild{"LabelType", &localLabel.LabelType}
    localLabel.EntityData.Children["paths"] = types.YChild{"Paths", &localLabel.Paths}
    localLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    localLabel.EntityData.Leafs["local-label-id"] = types.YLeaf{"LocalLabelId", localLabel.LocalLabelId}
    return &(localLabel.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetEntityData() *types.CommonEntityData {
    labelType.EntityData.YFilter = labelType.YFilter
    labelType.EntityData.YangName = "label-type"
    labelType.EntityData.BundleName = "cisco_ios_xr"
    labelType.EntityData.ParentYangName = "local-label"
    labelType.EntityData.SegmentPath = "label-type"
    labelType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelType.EntityData.Children = make(map[string]types.YChild)
    labelType.EntityData.Leafs = make(map[string]types.YLeaf)
    labelType.EntityData.Leafs["label-mode"] = types.YLeaf{"LabelMode", labelType.LabelMode}
    labelType.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", labelType.Prefix}
    labelType.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", labelType.PrefixLength}
    return &(labelType.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "local-label"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels
// Local Label
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetEntityData() *types.CommonEntityData {
    localLabels.EntityData.YFilter = localLabels.YFilter
    localLabels.EntityData.YangName = "local-labels"
    localLabels.EntityData.BundleName = "cisco_ios_xr"
    localLabels.EntityData.ParentYangName = "af"
    localLabels.EntityData.SegmentPath = "local-labels"
    localLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabels.EntityData.Children = make(map[string]types.YChild)
    localLabels.EntityData.Children["local-label"] = types.YChild{"LocalLabel", nil}
    for i := range localLabels.LocalLabel {
        localLabels.EntityData.Children[types.GetSegmentPath(&localLabels.LocalLabel[i])] = types.YChild{"LocalLabel", &localLabels.LocalLabel[i]}
    }
    localLabels.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(localLabels.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetEntityData() *types.CommonEntityData {
    localLabel.EntityData.YFilter = localLabel.YFilter
    localLabel.EntityData.YangName = "local-label"
    localLabel.EntityData.BundleName = "cisco_ios_xr"
    localLabel.EntityData.ParentYangName = "local-labels"
    localLabel.EntityData.SegmentPath = "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
    localLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabel.EntityData.Children = make(map[string]types.YChild)
    localLabel.EntityData.Children["label-type"] = types.YChild{"LabelType", &localLabel.LabelType}
    localLabel.EntityData.Children["paths"] = types.YChild{"Paths", &localLabel.Paths}
    localLabel.EntityData.Leafs = make(map[string]types.YLeaf)
    localLabel.EntityData.Leafs["local-label-id"] = types.YLeaf{"LocalLabelId", localLabel.LocalLabelId}
    return &(localLabel.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetEntityData() *types.CommonEntityData {
    labelType.EntityData.YFilter = labelType.YFilter
    labelType.EntityData.YangName = "label-type"
    labelType.EntityData.BundleName = "cisco_ios_xr"
    labelType.EntityData.ParentYangName = "local-label"
    labelType.EntityData.SegmentPath = "label-type"
    labelType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelType.EntityData.Children = make(map[string]types.YChild)
    labelType.EntityData.Leafs = make(map[string]types.YLeaf)
    labelType.EntityData.Leafs["label-mode"] = types.YLeaf{"LabelMode", labelType.LabelMode}
    labelType.EntityData.Leafs["prefix"] = types.YLeaf{"Prefix", labelType.Prefix}
    labelType.EntityData.Leafs["prefix-length"] = types.YLeaf{"PrefixLength", labelType.PrefixLength}
    return &(labelType.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "local-label"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = make(map[string]types.YChild)
    paths.EntityData.Children["path"] = types.YChild{"Path", nil}
    for i := range paths.Path {
        paths.EntityData.Children[types.GetSegmentPath(&paths.Path[i])] = types.YChild{"Path", &paths.Path[i]}
    }
    paths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(paths.EntityData)
}

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Number of paths. The type is interface{} with
    // range: 1..16.
    PathId interface{}

    // Type of Path (PopAndLookup, CrossConnect). The type is MplsStaticPath. The
    // default value is cross-connect.
    PathType interface{}

    // Type of label (Outlabel, ExpNull or Pop). The type is
    // MplsStaticOutLabelTypes. The default value is none.
    LabelType interface{}

    // Outgoing/NH Label. The type is interface{} with range: 16..1048575. The
    // default value is 16.
    NextHopLabel interface{}

    // Next Hop IP Address. The type is one of the following types: string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Next hop Address Family. The type is MplsStaticNhAddressFamily. The default
    // value is ipv4.
    Afi interface{}

    // NH Path Metric. The type is interface{} with range: 0..254. The default
    // value is 0.
    Metric interface{}

    // Next hop mode. The type is MplsStaticNhMode. The default value is
    // configured.
    NhMode interface{}

    // Path Role. The type is MplsStaticPathRole. The default value is primary.
    PathRole interface{}

    // Backup ID. The type is interface{} with range: 0..16. The default value is
    // 0.
    BackupId interface{}
}

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "paths"
    path.EntityData.SegmentPath = "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = make(map[string]types.YChild)
    path.EntityData.Leafs = make(map[string]types.YLeaf)
    path.EntityData.Leafs["path-id"] = types.YLeaf{"PathId", path.PathId}
    path.EntityData.Leafs["path-type"] = types.YLeaf{"PathType", path.PathType}
    path.EntityData.Leafs["label-type"] = types.YLeaf{"LabelType", path.LabelType}
    path.EntityData.Leafs["next-hop-label"] = types.YLeaf{"NextHopLabel", path.NextHopLabel}
    path.EntityData.Leafs["next-hop-address"] = types.YLeaf{"NextHopAddress", path.NextHopAddress}
    path.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", path.InterfaceName}
    path.EntityData.Leafs["afi"] = types.YLeaf{"Afi", path.Afi}
    path.EntityData.Leafs["metric"] = types.YLeaf{"Metric", path.Metric}
    path.EntityData.Leafs["nh-mode"] = types.YLeaf{"NhMode", path.NhMode}
    path.EntityData.Leafs["path-role"] = types.YLeaf{"PathRole", path.PathRole}
    path.EntityData.Leafs["backup-id"] = types.YLeaf{"BackupId", path.BackupId}
    return &(path.EntityData)
}

