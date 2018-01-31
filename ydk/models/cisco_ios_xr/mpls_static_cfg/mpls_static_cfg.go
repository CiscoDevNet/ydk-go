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

// MplsStaticNhMode represents Mpls static nh mode
type MplsStaticNhMode string

const (
    // Explicitly configured next hop path
    MplsStaticNhMode_configured MplsStaticNhMode = "configured"

    // Resolvable next hop which will result in a path
    // set
    MplsStaticNhMode_resolve MplsStaticNhMode = "resolve"
)

// MplsStaticNhAddressFamily represents Mpls static nh address family
type MplsStaticNhAddressFamily string

const (
    // IPv4 Next Hop
    MplsStaticNhAddressFamily_ipv4 MplsStaticNhAddressFamily = "ipv4"

    // IPv6 Next Hop
    MplsStaticNhAddressFamily_ipv6 MplsStaticNhAddressFamily = "ipv6"
)

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

// MplsStatic
// MPLS Static Configuration Data
type MplsStatic struct {
    parent types.Entity
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

func (mplsStatic *MplsStatic) GetFilter() yfilter.YFilter { return mplsStatic.YFilter }

func (mplsStatic *MplsStatic) SetFilter(yf yfilter.YFilter) { mplsStatic.YFilter = yf }

func (mplsStatic *MplsStatic) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "default-vrf" { return "DefaultVrf" }
    return ""
}

func (mplsStatic *MplsStatic) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-static-cfg:mpls-static"
}

func (mplsStatic *MplsStatic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &mplsStatic.Vrfs
    }
    if childYangName == "interfaces" {
        return &mplsStatic.Interfaces
    }
    if childYangName == "default-vrf" {
        return &mplsStatic.DefaultVrf
    }
    return nil
}

func (mplsStatic *MplsStatic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &mplsStatic.Vrfs
    children["interfaces"] = &mplsStatic.Interfaces
    children["default-vrf"] = &mplsStatic.DefaultVrf
    return children
}

func (mplsStatic *MplsStatic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = mplsStatic.Enable
    return leafs
}

func (mplsStatic *MplsStatic) GetBundleName() string { return "cisco_ios_xr" }

func (mplsStatic *MplsStatic) GetYangName() string { return "mpls-static" }

func (mplsStatic *MplsStatic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsStatic *MplsStatic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsStatic *MplsStatic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsStatic *MplsStatic) SetParent(parent types.Entity) { mplsStatic.parent = parent }

func (mplsStatic *MplsStatic) GetParent() types.Entity { return mplsStatic.parent }

func (mplsStatic *MplsStatic) GetParentYangName() string { return "Cisco-IOS-XR-mpls-static-cfg" }

// MplsStatic_Vrfs
// VRF table
type MplsStatic_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of MplsStatic_Vrfs_Vrf.
    Vrf []MplsStatic_Vrfs_Vrf
}

func (vrfs *MplsStatic_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *MplsStatic_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *MplsStatic_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *MplsStatic_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *MplsStatic_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *MplsStatic_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *MplsStatic_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *MplsStatic_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *MplsStatic_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *MplsStatic_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *MplsStatic_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *MplsStatic_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *MplsStatic_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *MplsStatic_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *MplsStatic_Vrfs) GetParentYangName() string { return "mpls-static" }

// MplsStatic_Vrfs_Vrf
// VRF Name
type MplsStatic_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Table of the Label Switched Paths.
    LabelSwitchedPaths MplsStatic_Vrfs_Vrf_LabelSwitchedPaths

    // Address Family Table.
    Afs MplsStatic_Vrfs_Vrf_Afs
}

func (vrf *MplsStatic_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *MplsStatic_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *MplsStatic_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enable" { return "Enable" }
    if yname == "label-switched-paths" { return "LabelSwitchedPaths" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (vrf *MplsStatic_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *MplsStatic_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-paths" {
        return &vrf.LabelSwitchedPaths
    }
    if childYangName == "afs" {
        return &vrf.Afs
    }
    return nil
}

func (vrf *MplsStatic_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-switched-paths"] = &vrf.LabelSwitchedPaths
    children["afs"] = &vrf.Afs
    return children
}

func (vrf *MplsStatic_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["enable"] = vrf.Enable
    return leafs
}

func (vrf *MplsStatic_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *MplsStatic_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *MplsStatic_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *MplsStatic_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *MplsStatic_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *MplsStatic_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *MplsStatic_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *MplsStatic_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths
// Table of the Label Switched Paths
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label Switched Path. The type is slice of
    // MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetFilter() yfilter.YFilter { return labelSwitchedPaths.YFilter }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) SetFilter(yf yfilter.YFilter) { labelSwitchedPaths.YFilter = yf }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetGoName(yname string) string {
    if yname == "label-switched-path" { return "LabelSwitchedPath" }
    return ""
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetSegmentPath() string {
    return "label-switched-paths"
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-path" {
        for _, c := range labelSwitchedPaths.LabelSwitchedPath {
            if labelSwitchedPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath{}
        labelSwitchedPaths.LabelSwitchedPath = append(labelSwitchedPaths.LabelSwitchedPath, child)
        return &labelSwitchedPaths.LabelSwitchedPath[len(labelSwitchedPaths.LabelSwitchedPath)-1]
    }
    return nil
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        children[labelSwitchedPaths.LabelSwitchedPath[i].GetSegmentPath()] = &labelSwitchedPaths.LabelSwitchedPath[i]
    }
    return children
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetBundleName() string { return "cisco_ios_xr" }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetYangName() string { return "label-switched-paths" }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) SetParent(parent types.Entity) { labelSwitchedPaths.parent = parent }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetParent() types.Entity { return labelSwitchedPaths.parent }

func (labelSwitchedPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths) GetParentYangName() string { return "vrf" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath
// Label Switched Path
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSP Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetFilter() yfilter.YFilter { return labelSwitchedPath.YFilter }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) SetFilter(yf yfilter.YFilter) { labelSwitchedPath.YFilter = yf }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetGoName(yname string) string {
    if yname == "lsp-name" { return "LspName" }
    if yname == "enable" { return "Enable" }
    if yname == "backup-paths" { return "BackupPaths" }
    if yname == "in-label" { return "InLabel" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetSegmentPath() string {
    return "label-switched-path" + "[lsp-name='" + fmt.Sprintf("%v", labelSwitchedPath.LspName) + "']"
}

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "backup-paths" {
        return &labelSwitchedPath.BackupPaths
    }
    if childYangName == "in-label" {
        return &labelSwitchedPath.InLabel
    }
    if childYangName == "paths" {
        return &labelSwitchedPath.Paths
    }
    return nil
}

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["backup-paths"] = &labelSwitchedPath.BackupPaths
    children["in-label"] = &labelSwitchedPath.InLabel
    children["paths"] = &labelSwitchedPath.Paths
    return children
}

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsp-name"] = labelSwitchedPath.LspName
    leafs["enable"] = labelSwitchedPath.Enable
    return leafs
}

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetBundleName() string { return "cisco_ios_xr" }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetYangName() string { return "label-switched-path" }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) SetParent(parent types.Entity) { labelSwitchedPath.parent = parent }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetParent() types.Entity { return labelSwitchedPath.parent }

func (labelSwitchedPath *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath) GetParentYangName() string { return "label-switched-paths" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths
// Backup Path Parameters
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path.
    Path []MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetFilter() yfilter.YFilter { return backupPaths.YFilter }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) SetFilter(yf yfilter.YFilter) { backupPaths.YFilter = yf }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetSegmentPath() string {
    return "backup-paths"
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range backupPaths.Path {
            if backupPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path{}
        backupPaths.Path = append(backupPaths.Path, child)
        return &backupPaths.Path[len(backupPaths.Path)-1]
    }
    return nil
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range backupPaths.Path {
        children[backupPaths.Path[i].GetSegmentPath()] = &backupPaths.Path[i]
    }
    return children
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetBundleName() string { return "cisco_ios_xr" }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetYangName() string { return "backup-paths" }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) SetParent(parent types.Entity) { backupPaths.parent = parent }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetParent() types.Entity { return backupPaths.parent }

func (backupPaths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetParentYangName() string { return "backup-paths" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel
// MPLS Static Local Label Value
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local Label. The type is interface{} with range: 16..1048575.
    InLabelValue interface{}

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}

    // Top Label Hashing Mode. The type is bool.
    TlhMode interface{}
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetFilter() yfilter.YFilter { return inLabel.YFilter }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) SetFilter(yf yfilter.YFilter) { inLabel.YFilter = yf }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetGoName(yname string) string {
    if yname == "in-label-value" { return "InLabelValue" }
    if yname == "label-mode" { return "LabelMode" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "tlh-mode" { return "TlhMode" }
    return ""
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetSegmentPath() string {
    return "in-label"
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-label-value"] = inLabel.InLabelValue
    leafs["label-mode"] = inLabel.LabelMode
    leafs["prefix"] = inLabel.Prefix
    leafs["prefix-length"] = inLabel.PrefixLength
    leafs["tlh-mode"] = inLabel.TlhMode
    return leafs
}

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetBundleName() string { return "cisco_ios_xr" }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetYangName() string { return "in-label" }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) SetParent(parent types.Entity) { inLabel.parent = parent }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetParent() types.Entity { return inLabel.parent }

func (inLabel *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths
// Forward Path Parameters
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path.
    Path []MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetYangName() string { return "paths" }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_Vrfs_Vrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetParentYangName() string { return "paths" }

// MplsStatic_Vrfs_Vrf_Afs
// Address Family Table
type MplsStatic_Vrfs_Vrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family. The type is slice of MplsStatic_Vrfs_Vrf_Afs_Af.
    Af []MplsStatic_Vrfs_Vrf_Afs_Af
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsStatic_Vrfs_Vrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsStatic_Vrfs_Vrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsStatic_Vrfs_Vrf_Afs) GetParentYangName() string { return "vrf" }

// MplsStatic_Vrfs_Vrf_Afs_Af
// Address Family
type MplsStatic_Vrfs_Vrf_Afs_Af struct {
    parent types.Entity
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

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "enable" { return "Enable" }
    if yname == "top-label-hash" { return "TopLabelHash" }
    if yname == "local-labels" { return "LocalLabels" }
    return ""
}

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[afi='" + fmt.Sprintf("%v", af.Afi) + "']"
}

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "top-label-hash" {
        return &af.TopLabelHash
    }
    if childYangName == "local-labels" {
        return &af.LocalLabels
    }
    return nil
}

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["top-label-hash"] = &af.TopLabelHash
    children["local-labels"] = &af.LocalLabels
    return children
}

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = af.Afi
    leafs["enable"] = af.Enable
    return leafs
}

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsStatic_Vrfs_Vrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash
// Top Label Hash
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local Label.
    LocalLabels MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetFilter() yfilter.YFilter { return topLabelHash.YFilter }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) SetFilter(yf yfilter.YFilter) { topLabelHash.YFilter = yf }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetGoName(yname string) string {
    if yname == "local-labels" { return "LocalLabels" }
    return ""
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetSegmentPath() string {
    return "top-label-hash"
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-labels" {
        return &topLabelHash.LocalLabels
    }
    return nil
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-labels"] = &topLabelHash.LocalLabels
    return children
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetBundleName() string { return "cisco_ios_xr" }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetYangName() string { return "top-label-hash" }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) SetParent(parent types.Entity) { topLabelHash.parent = parent }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetParent() types.Entity { return topLabelHash.parent }

func (topLabelHash *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash) GetParentYangName() string { return "af" }

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels
// Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetFilter() yfilter.YFilter { return localLabels.YFilter }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) SetFilter(yf yfilter.YFilter) { localLabels.YFilter = yf }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    return ""
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetSegmentPath() string {
    return "local-labels"
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label" {
        for _, c := range localLabels.LocalLabel {
            if localLabels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel{}
        localLabels.LocalLabel = append(localLabels.LocalLabel, child)
        return &localLabels.LocalLabel[len(localLabels.LocalLabel)-1]
    }
    return nil
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabels.LocalLabel {
        children[localLabels.LocalLabel[i].GetSegmentPath()] = &localLabels.LocalLabel[i]
    }
    return children
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetBundleName() string { return "cisco_ios_xr" }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetYangName() string { return "local-labels" }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) SetParent(parent types.Entity) { localLabels.parent = parent }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetParent() types.Entity { return localLabels.parent }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels) GetParentYangName() string { return "top-label-hash" }

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetFilter() yfilter.YFilter { return localLabel.YFilter }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) SetFilter(yf yfilter.YFilter) { localLabel.YFilter = yf }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetGoName(yname string) string {
    if yname == "local-label-id" { return "LocalLabelId" }
    if yname == "label-type" { return "LabelType" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetSegmentPath() string {
    return "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-type" {
        return &localLabel.LabelType
    }
    if childYangName == "paths" {
        return &localLabel.Paths
    }
    return nil
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-type"] = &localLabel.LabelType
    children["paths"] = &localLabel.Paths
    return children
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label-id"] = localLabel.LocalLabelId
    return leafs
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetBundleName() string { return "cisco_ios_xr" }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetYangName() string { return "local-label" }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) SetParent(parent types.Entity) { localLabel.parent = parent }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetParent() types.Entity { return localLabel.parent }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetParentYangName() string { return "local-labels" }

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetFilter() yfilter.YFilter { return labelType.YFilter }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) SetFilter(yf yfilter.YFilter) { labelType.YFilter = yf }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetGoName(yname string) string {
    if yname == "label-mode" { return "LabelMode" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetSegmentPath() string {
    return "label-type"
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-mode"] = labelType.LabelMode
    leafs["prefix"] = labelType.Prefix
    leafs["prefix-length"] = labelType.PrefixLength
    return leafs
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetBundleName() string { return "cisco_ios_xr" }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetYangName() string { return "label-type" }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) SetParent(parent types.Entity) { labelType.parent = parent }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetParent() types.Entity { return labelType.parent }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetYangName() string { return "paths" }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetParentYangName() string { return "paths" }

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels
// Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetFilter() yfilter.YFilter { return localLabels.YFilter }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) SetFilter(yf yfilter.YFilter) { localLabels.YFilter = yf }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    return ""
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetSegmentPath() string {
    return "local-labels"
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label" {
        for _, c := range localLabels.LocalLabel {
            if localLabels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel{}
        localLabels.LocalLabel = append(localLabels.LocalLabel, child)
        return &localLabels.LocalLabel[len(localLabels.LocalLabel)-1]
    }
    return nil
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabels.LocalLabel {
        children[localLabels.LocalLabel[i].GetSegmentPath()] = &localLabels.LocalLabel[i]
    }
    return children
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetBundleName() string { return "cisco_ios_xr" }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetYangName() string { return "local-labels" }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) SetParent(parent types.Entity) { localLabels.parent = parent }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetParent() types.Entity { return localLabels.parent }

func (localLabels *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels) GetParentYangName() string { return "af" }

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetFilter() yfilter.YFilter { return localLabel.YFilter }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) SetFilter(yf yfilter.YFilter) { localLabel.YFilter = yf }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetGoName(yname string) string {
    if yname == "local-label-id" { return "LocalLabelId" }
    if yname == "label-type" { return "LabelType" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetSegmentPath() string {
    return "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-type" {
        return &localLabel.LabelType
    }
    if childYangName == "paths" {
        return &localLabel.Paths
    }
    return nil
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-type"] = &localLabel.LabelType
    children["paths"] = &localLabel.Paths
    return children
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label-id"] = localLabel.LocalLabelId
    return leafs
}

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetBundleName() string { return "cisco_ios_xr" }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetYangName() string { return "local-label" }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) SetParent(parent types.Entity) { localLabel.parent = parent }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetParent() types.Entity { return localLabel.parent }

func (localLabel *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel) GetParentYangName() string { return "local-labels" }

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetFilter() yfilter.YFilter { return labelType.YFilter }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) SetFilter(yf yfilter.YFilter) { labelType.YFilter = yf }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetGoName(yname string) string {
    if yname == "label-mode" { return "LabelMode" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetSegmentPath() string {
    return "label-type"
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-mode"] = labelType.LabelMode
    leafs["prefix"] = labelType.Prefix
    leafs["prefix-length"] = labelType.PrefixLength
    return leafs
}

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetBundleName() string { return "cisco_ios_xr" }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetYangName() string { return "label-type" }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) SetParent(parent types.Entity) { labelType.parent = parent }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetParent() types.Entity { return labelType.parent }

func (labelType *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetYangName() string { return "paths" }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_Vrfs_Vrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetParentYangName() string { return "paths" }

// MplsStatic_Interfaces
// MPLS Static Interface Table
type MplsStatic_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static Interface Enable. The type is slice of
    // MplsStatic_Interfaces_Interface.
    Interface []MplsStatic_Interfaces_Interface
}

func (interfaces *MplsStatic_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsStatic_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsStatic_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsStatic_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsStatic_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsStatic_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsStatic_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsStatic_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *MplsStatic_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsStatic_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *MplsStatic_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *MplsStatic_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *MplsStatic_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsStatic_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsStatic_Interfaces) GetParentYangName() string { return "mpls-static" }

// MplsStatic_Interfaces_Interface
// MPLS Static Interface Enable
type MplsStatic_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (self *MplsStatic_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsStatic_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsStatic_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (self *MplsStatic_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *MplsStatic_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *MplsStatic_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *MplsStatic_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *MplsStatic_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *MplsStatic_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsStatic_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *MplsStatic_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *MplsStatic_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *MplsStatic_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsStatic_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsStatic_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsStatic_DefaultVrf
// Default VRF
type MplsStatic_DefaultVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS Static Apply Enable. The type is interface{}.
    Enable interface{}

    // Table of the Label Switched Paths.
    LabelSwitchedPaths MplsStatic_DefaultVrf_LabelSwitchedPaths

    // Address Family Table.
    Afs MplsStatic_DefaultVrf_Afs
}

func (defaultVrf *MplsStatic_DefaultVrf) GetFilter() yfilter.YFilter { return defaultVrf.YFilter }

func (defaultVrf *MplsStatic_DefaultVrf) SetFilter(yf yfilter.YFilter) { defaultVrf.YFilter = yf }

func (defaultVrf *MplsStatic_DefaultVrf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "label-switched-paths" { return "LabelSwitchedPaths" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (defaultVrf *MplsStatic_DefaultVrf) GetSegmentPath() string {
    return "default-vrf"
}

func (defaultVrf *MplsStatic_DefaultVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-paths" {
        return &defaultVrf.LabelSwitchedPaths
    }
    if childYangName == "afs" {
        return &defaultVrf.Afs
    }
    return nil
}

func (defaultVrf *MplsStatic_DefaultVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-switched-paths"] = &defaultVrf.LabelSwitchedPaths
    children["afs"] = &defaultVrf.Afs
    return children
}

func (defaultVrf *MplsStatic_DefaultVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = defaultVrf.Enable
    return leafs
}

func (defaultVrf *MplsStatic_DefaultVrf) GetBundleName() string { return "cisco_ios_xr" }

func (defaultVrf *MplsStatic_DefaultVrf) GetYangName() string { return "default-vrf" }

func (defaultVrf *MplsStatic_DefaultVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultVrf *MplsStatic_DefaultVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultVrf *MplsStatic_DefaultVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultVrf *MplsStatic_DefaultVrf) SetParent(parent types.Entity) { defaultVrf.parent = parent }

func (defaultVrf *MplsStatic_DefaultVrf) GetParent() types.Entity { return defaultVrf.parent }

func (defaultVrf *MplsStatic_DefaultVrf) GetParentYangName() string { return "mpls-static" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths
// Table of the Label Switched Paths
type MplsStatic_DefaultVrf_LabelSwitchedPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label Switched Path. The type is slice of
    // MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath.
    LabelSwitchedPath []MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetFilter() yfilter.YFilter { return labelSwitchedPaths.YFilter }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) SetFilter(yf yfilter.YFilter) { labelSwitchedPaths.YFilter = yf }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetGoName(yname string) string {
    if yname == "label-switched-path" { return "LabelSwitchedPath" }
    return ""
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetSegmentPath() string {
    return "label-switched-paths"
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-switched-path" {
        for _, c := range labelSwitchedPaths.LabelSwitchedPath {
            if labelSwitchedPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath{}
        labelSwitchedPaths.LabelSwitchedPath = append(labelSwitchedPaths.LabelSwitchedPath, child)
        return &labelSwitchedPaths.LabelSwitchedPath[len(labelSwitchedPaths.LabelSwitchedPath)-1]
    }
    return nil
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelSwitchedPaths.LabelSwitchedPath {
        children[labelSwitchedPaths.LabelSwitchedPath[i].GetSegmentPath()] = &labelSwitchedPaths.LabelSwitchedPath[i]
    }
    return children
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetBundleName() string { return "cisco_ios_xr" }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetYangName() string { return "label-switched-paths" }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) SetParent(parent types.Entity) { labelSwitchedPaths.parent = parent }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetParent() types.Entity { return labelSwitchedPaths.parent }

func (labelSwitchedPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths) GetParentYangName() string { return "default-vrf" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath
// Label Switched Path
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSP Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetFilter() yfilter.YFilter { return labelSwitchedPath.YFilter }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) SetFilter(yf yfilter.YFilter) { labelSwitchedPath.YFilter = yf }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetGoName(yname string) string {
    if yname == "lsp-name" { return "LspName" }
    if yname == "enable" { return "Enable" }
    if yname == "backup-paths" { return "BackupPaths" }
    if yname == "in-label" { return "InLabel" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetSegmentPath() string {
    return "label-switched-path" + "[lsp-name='" + fmt.Sprintf("%v", labelSwitchedPath.LspName) + "']"
}

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "backup-paths" {
        return &labelSwitchedPath.BackupPaths
    }
    if childYangName == "in-label" {
        return &labelSwitchedPath.InLabel
    }
    if childYangName == "paths" {
        return &labelSwitchedPath.Paths
    }
    return nil
}

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["backup-paths"] = &labelSwitchedPath.BackupPaths
    children["in-label"] = &labelSwitchedPath.InLabel
    children["paths"] = &labelSwitchedPath.Paths
    return children
}

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsp-name"] = labelSwitchedPath.LspName
    leafs["enable"] = labelSwitchedPath.Enable
    return leafs
}

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetBundleName() string { return "cisco_ios_xr" }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetYangName() string { return "label-switched-path" }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) SetParent(parent types.Entity) { labelSwitchedPath.parent = parent }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetParent() types.Entity { return labelSwitchedPath.parent }

func (labelSwitchedPath *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath) GetParentYangName() string { return "label-switched-paths" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths
// Backup Path Parameters
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path.
    Path []MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetFilter() yfilter.YFilter { return backupPaths.YFilter }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) SetFilter(yf yfilter.YFilter) { backupPaths.YFilter = yf }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetSegmentPath() string {
    return "backup-paths"
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range backupPaths.Path {
            if backupPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path{}
        backupPaths.Path = append(backupPaths.Path, child)
        return &backupPaths.Path[len(backupPaths.Path)-1]
    }
    return nil
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range backupPaths.Path {
        children[backupPaths.Path[i].GetSegmentPath()] = &backupPaths.Path[i]
    }
    return children
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetBundleName() string { return "cisco_ios_xr" }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetYangName() string { return "backup-paths" }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) SetParent(parent types.Entity) { backupPaths.parent = parent }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetParent() types.Entity { return backupPaths.parent }

func (backupPaths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path
// Path Information
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_BackupPaths_Path) GetParentYangName() string { return "backup-paths" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel
// MPLS Static Local Label Value
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local Label. The type is interface{} with range: 16..1048575.
    InLabelValue interface{}

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}

    // Top Label Hashing Mode. The type is bool.
    TlhMode interface{}
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetFilter() yfilter.YFilter { return inLabel.YFilter }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) SetFilter(yf yfilter.YFilter) { inLabel.YFilter = yf }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetGoName(yname string) string {
    if yname == "in-label-value" { return "InLabelValue" }
    if yname == "label-mode" { return "LabelMode" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "tlh-mode" { return "TlhMode" }
    return ""
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetSegmentPath() string {
    return "in-label"
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-label-value"] = inLabel.InLabelValue
    leafs["label-mode"] = inLabel.LabelMode
    leafs["prefix"] = inLabel.Prefix
    leafs["prefix-length"] = inLabel.PrefixLength
    leafs["tlh-mode"] = inLabel.TlhMode
    return leafs
}

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetBundleName() string { return "cisco_ios_xr" }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetYangName() string { return "in-label" }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) SetParent(parent types.Entity) { inLabel.parent = parent }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetParent() types.Entity { return inLabel.parent }

func (inLabel *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_InLabel) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths
// Forward Path Parameters
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path.
    Path []MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetYangName() string { return "paths" }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths) GetParentYangName() string { return "label-switched-path" }

// MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path
// Path Information
type MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_DefaultVrf_LabelSwitchedPaths_LabelSwitchedPath_Paths_Path) GetParentYangName() string { return "paths" }

// MplsStatic_DefaultVrf_Afs
// Address Family Table
type MplsStatic_DefaultVrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family. The type is slice of MplsStatic_DefaultVrf_Afs_Af.
    Af []MplsStatic_DefaultVrf_Afs_Af
}

func (afs *MplsStatic_DefaultVrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsStatic_DefaultVrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsStatic_DefaultVrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsStatic_DefaultVrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsStatic_DefaultVrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsStatic_DefaultVrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsStatic_DefaultVrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsStatic_DefaultVrf_Afs) GetBundleName() string { return "cisco_ios_xr" }

func (afs *MplsStatic_DefaultVrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsStatic_DefaultVrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (afs *MplsStatic_DefaultVrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (afs *MplsStatic_DefaultVrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (afs *MplsStatic_DefaultVrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsStatic_DefaultVrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsStatic_DefaultVrf_Afs) GetParentYangName() string { return "default-vrf" }

// MplsStatic_DefaultVrf_Afs_Af
// Address Family
type MplsStatic_DefaultVrf_Afs_Af struct {
    parent types.Entity
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

func (af *MplsStatic_DefaultVrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsStatic_DefaultVrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "enable" { return "Enable" }
    if yname == "top-label-hash" { return "TopLabelHash" }
    if yname == "local-labels" { return "LocalLabels" }
    return ""
}

func (af *MplsStatic_DefaultVrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[afi='" + fmt.Sprintf("%v", af.Afi) + "']"
}

func (af *MplsStatic_DefaultVrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "top-label-hash" {
        return &af.TopLabelHash
    }
    if childYangName == "local-labels" {
        return &af.LocalLabels
    }
    return nil
}

func (af *MplsStatic_DefaultVrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["top-label-hash"] = &af.TopLabelHash
    children["local-labels"] = &af.LocalLabels
    return children
}

func (af *MplsStatic_DefaultVrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = af.Afi
    leafs["enable"] = af.Enable
    return leafs
}

func (af *MplsStatic_DefaultVrf_Afs_Af) GetBundleName() string { return "cisco_ios_xr" }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (af *MplsStatic_DefaultVrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsStatic_DefaultVrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash
// Top Label Hash
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local Label.
    LocalLabels MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetFilter() yfilter.YFilter { return topLabelHash.YFilter }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) SetFilter(yf yfilter.YFilter) { topLabelHash.YFilter = yf }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetGoName(yname string) string {
    if yname == "local-labels" { return "LocalLabels" }
    return ""
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetSegmentPath() string {
    return "top-label-hash"
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-labels" {
        return &topLabelHash.LocalLabels
    }
    return nil
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-labels"] = &topLabelHash.LocalLabels
    return children
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetBundleName() string { return "cisco_ios_xr" }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetYangName() string { return "top-label-hash" }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) SetParent(parent types.Entity) { topLabelHash.parent = parent }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetParent() types.Entity { return topLabelHash.parent }

func (topLabelHash *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash) GetParentYangName() string { return "af" }

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels
// Local Label
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetFilter() yfilter.YFilter { return localLabels.YFilter }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) SetFilter(yf yfilter.YFilter) { localLabels.YFilter = yf }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    return ""
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetSegmentPath() string {
    return "local-labels"
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label" {
        for _, c := range localLabels.LocalLabel {
            if localLabels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel{}
        localLabels.LocalLabel = append(localLabels.LocalLabel, child)
        return &localLabels.LocalLabel[len(localLabels.LocalLabel)-1]
    }
    return nil
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabels.LocalLabel {
        children[localLabels.LocalLabel[i].GetSegmentPath()] = &localLabels.LocalLabel[i]
    }
    return children
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetBundleName() string { return "cisco_ios_xr" }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetYangName() string { return "local-labels" }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) SetParent(parent types.Entity) { localLabels.parent = parent }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetParent() types.Entity { return localLabels.parent }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels) GetParentYangName() string { return "top-label-hash" }

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetFilter() yfilter.YFilter { return localLabel.YFilter }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) SetFilter(yf yfilter.YFilter) { localLabel.YFilter = yf }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetGoName(yname string) string {
    if yname == "local-label-id" { return "LocalLabelId" }
    if yname == "label-type" { return "LabelType" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetSegmentPath() string {
    return "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-type" {
        return &localLabel.LabelType
    }
    if childYangName == "paths" {
        return &localLabel.Paths
    }
    return nil
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-type"] = &localLabel.LabelType
    children["paths"] = &localLabel.Paths
    return children
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label-id"] = localLabel.LocalLabelId
    return leafs
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetBundleName() string { return "cisco_ios_xr" }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetYangName() string { return "local-label" }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) SetParent(parent types.Entity) { localLabel.parent = parent }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetParent() types.Entity { return localLabel.parent }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel) GetParentYangName() string { return "local-labels" }

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetFilter() yfilter.YFilter { return labelType.YFilter }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) SetFilter(yf yfilter.YFilter) { labelType.YFilter = yf }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetGoName(yname string) string {
    if yname == "label-mode" { return "LabelMode" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetSegmentPath() string {
    return "label-type"
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-mode"] = labelType.LabelMode
    leafs["prefix"] = labelType.Prefix
    leafs["prefix-length"] = labelType.PrefixLength
    return leafs
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetBundleName() string { return "cisco_ios_xr" }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetYangName() string { return "label-type" }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) SetParent(parent types.Entity) { labelType.parent = parent }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetParent() types.Entity { return labelType.parent }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_LabelType) GetParentYangName() string { return "local-label" }

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetYangName() string { return "paths" }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths) GetParentYangName() string { return "local-label" }

// MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_DefaultVrf_Afs_Af_TopLabelHash_LocalLabels_LocalLabel_Paths_Path) GetParentYangName() string { return "paths" }

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels
// Local Label
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify Local Label. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetFilter() yfilter.YFilter { return localLabels.YFilter }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) SetFilter(yf yfilter.YFilter) { localLabels.YFilter = yf }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    return ""
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetSegmentPath() string {
    return "local-labels"
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label" {
        for _, c := range localLabels.LocalLabel {
            if localLabels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel{}
        localLabels.LocalLabel = append(localLabels.LocalLabel, child)
        return &localLabels.LocalLabel[len(localLabels.LocalLabel)-1]
    }
    return nil
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabels.LocalLabel {
        children[localLabels.LocalLabel[i].GetSegmentPath()] = &localLabels.LocalLabel[i]
    }
    return children
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetBundleName() string { return "cisco_ios_xr" }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetYangName() string { return "local-labels" }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) SetParent(parent types.Entity) { localLabels.parent = parent }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetParent() types.Entity { return localLabels.parent }

func (localLabels *MplsStatic_DefaultVrf_Afs_Af_LocalLabels) GetParentYangName() string { return "af" }

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel
// Specify Local Label
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // MPLS Static Local Label Value.
    LabelType MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType

    // Forward Path Parameters.
    Paths MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetFilter() yfilter.YFilter { return localLabel.YFilter }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) SetFilter(yf yfilter.YFilter) { localLabel.YFilter = yf }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetGoName(yname string) string {
    if yname == "local-label-id" { return "LocalLabelId" }
    if yname == "label-type" { return "LabelType" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetSegmentPath() string {
    return "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-type" {
        return &localLabel.LabelType
    }
    if childYangName == "paths" {
        return &localLabel.Paths
    }
    return nil
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label-type"] = &localLabel.LabelType
    children["paths"] = &localLabel.Paths
    return children
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label-id"] = localLabel.LocalLabelId
    return leafs
}

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetBundleName() string { return "cisco_ios_xr" }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetYangName() string { return "local-label" }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) SetParent(parent types.Entity) { localLabel.parent = parent }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetParent() types.Entity { return localLabel.parent }

func (localLabel *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel) GetParentYangName() string { return "local-labels" }

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType
// MPLS Static Local Label Value
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label Mode (PerVRF, PerPrefix or LSP). The type is MplsStaticLabelMode.
    LabelMode interface{}

    // Address (IPv4/6 depending on AFI). The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Prefix interface{}

    // Prefix length. The type is interface{} with range: -2147483648..2147483647.
    PrefixLength interface{}
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetFilter() yfilter.YFilter { return labelType.YFilter }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) SetFilter(yf yfilter.YFilter) { labelType.YFilter = yf }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetGoName(yname string) string {
    if yname == "label-mode" { return "LabelMode" }
    if yname == "prefix" { return "Prefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    return ""
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetSegmentPath() string {
    return "label-type"
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-mode"] = labelType.LabelMode
    leafs["prefix"] = labelType.Prefix
    leafs["prefix-length"] = labelType.PrefixLength
    return leafs
}

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetBundleName() string { return "cisco_ios_xr" }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetYangName() string { return "label-type" }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) SetParent(parent types.Entity) { labelType.parent = parent }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetParent() types.Entity { return labelType.parent }

func (labelType *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_LabelType) GetParentYangName() string { return "local-label" }

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths
// Forward Path Parameters
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Information. The type is slice of
    // MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path.
    Path []MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetGoName(yname string) string {
    if yname == "path" { return "Path" }
    return ""
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "path" {
        for _, c := range paths.Path {
            if paths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path{}
        paths.Path = append(paths.Path, child)
        return &paths.Path[len(paths.Path)-1]
    }
    return nil
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range paths.Path {
        children[paths.Path[i].GetSegmentPath()] = &paths.Path[i]
    }
    return children
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetBundleName() string { return "cisco_ios_xr" }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetYangName() string { return "paths" }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths) GetParentYangName() string { return "local-label" }

// MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path
// Path Information
type MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path struct {
    parent types.Entity
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // The default value is 0.0.0.0..
    NextHopAddress interface{}

    // Next hop Interface with form <Interface>R/S/I/P. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
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

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetFilter() yfilter.YFilter { return path.YFilter }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) SetFilter(yf yfilter.YFilter) { path.YFilter = yf }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetGoName(yname string) string {
    if yname == "path-id" { return "PathId" }
    if yname == "path-type" { return "PathType" }
    if yname == "label-type" { return "LabelType" }
    if yname == "next-hop-label" { return "NextHopLabel" }
    if yname == "next-hop-address" { return "NextHopAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "metric" { return "Metric" }
    if yname == "nh-mode" { return "NhMode" }
    if yname == "path-role" { return "PathRole" }
    if yname == "backup-id" { return "BackupId" }
    return ""
}

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetSegmentPath() string {
    return "path" + "[path-id='" + fmt.Sprintf("%v", path.PathId) + "']"
}

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-id"] = path.PathId
    leafs["path-type"] = path.PathType
    leafs["label-type"] = path.LabelType
    leafs["next-hop-label"] = path.NextHopLabel
    leafs["next-hop-address"] = path.NextHopAddress
    leafs["interface-name"] = path.InterfaceName
    leafs["afi"] = path.Afi
    leafs["metric"] = path.Metric
    leafs["nh-mode"] = path.NhMode
    leafs["path-role"] = path.PathRole
    leafs["backup-id"] = path.BackupId
    return leafs
}

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetBundleName() string { return "cisco_ios_xr" }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetYangName() string { return "path" }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) SetParent(parent types.Entity) { path.parent = parent }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetParent() types.Entity { return path.parent }

func (path *MplsStatic_DefaultVrf_Afs_Af_LocalLabels_LocalLabel_Paths_Path) GetParentYangName() string { return "paths" }

