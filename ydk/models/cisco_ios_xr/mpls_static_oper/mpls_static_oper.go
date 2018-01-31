// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-static package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpls-static: MPLS STATIC operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_static_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_static_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-static-oper mpls-static}", reflect.TypeOf(MplsStatic{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-static-oper:mpls-static", reflect.TypeOf(MplsStatic{}))
}

// MgmtMplsStaticPathStatus represents Mgmt mpls static path status
type MgmtMplsStaticPathStatus string

const (
    // Path NextHop No Status
    MgmtMplsStaticPathStatus_path_next_hop_none MgmtMplsStaticPathStatus = "path-next-hop-none"

    // Path NextHop Interface Down 
    MgmtMplsStaticPathStatus_path_next_hop_interface_down MgmtMplsStaticPathStatus = "path-next-hop-interface-down"

    // Path NextHop Valid
    MgmtMplsStaticPathStatus_path_next_hop_valid MgmtMplsStaticPathStatus = "path-next-hop-valid"

    // Path NextHop Resolve Failed
    MgmtMplsStaticPathStatus_resolve_failed MgmtMplsStaticPathStatus = "resolve-failed"

    // FRR Backup
    MgmtMplsStaticPathStatus_frr_backup MgmtMplsStaticPathStatus = "frr-backup"

    // Backup
    MgmtMplsStaticPathStatus_backup MgmtMplsStaticPathStatus = "backup"
)

// MgmtStaticPath represents Mgmt static path
type MgmtStaticPath string

const (
    // Crossconnect Path
    MgmtStaticPath_cross_connect_path MgmtStaticPath = "cross-connect-path"

    // Pop and Lookup Path
    MgmtStaticPath_pop_lookup_path MgmtStaticPath = "pop-lookup-path"
)

// MgmtMplsStaticLabelMode represents Mgmt mpls static label mode
type MgmtMplsStaticLabelMode string

const (
    // No Label Mode
    MgmtMplsStaticLabelMode_none MgmtMplsStaticLabelMode = "none"

    // Per-prefix Label
    MgmtMplsStaticLabelMode_per_prefix MgmtMplsStaticLabelMode = "per-prefix"

    // Per-VRF label
    MgmtMplsStaticLabelMode_per_vrf MgmtMplsStaticLabelMode = "per-vrf"

    // Label with crossconnect
    MgmtMplsStaticLabelMode_cross_connect MgmtMplsStaticLabelMode = "cross-connect"
)

// MgmtMplsStaticLabelStatus represents Mgmt mpls static label status
type MgmtMplsStaticLabelStatus string

const (
    // Label Not Created
    MgmtMplsStaticLabelStatus_not_created MgmtMplsStaticLabelStatus = "not-created"

    // Label without active VRF
    MgmtMplsStaticLabelStatus_vrf_down MgmtMplsStaticLabelStatus = "vrf-down"

    // Rewrite without active VRF
    MgmtMplsStaticLabelStatus_rewrite_vrf_down MgmtMplsStaticLabelStatus = "rewrite-vrf-down"

    // LSD is disconnected
    MgmtMplsStaticLabelStatus_lsd_disconnected MgmtMplsStaticLabelStatus = "lsd-disconnected"

    // LSD operation failed
    MgmtMplsStaticLabelStatus_lsd_failed MgmtMplsStaticLabelStatus = "lsd-failed"

    // Waiting for LSD operation
    MgmtMplsStaticLabelStatus_wait_for_lsd_reply MgmtMplsStaticLabelStatus = "wait-for-lsd-reply"

    // Label Created
    MgmtMplsStaticLabelStatus_label_created MgmtMplsStaticLabelStatus = "label-created"

    // Label Creation Failed
    MgmtMplsStaticLabelStatus_label_create_failed MgmtMplsStaticLabelStatus = "label-create-failed"

    // Rewrite Creation Failed
    MgmtMplsStaticLabelStatus_label_rewrite_failed MgmtMplsStaticLabelStatus = "label-rewrite-failed"

    // Rewrite NextHop Down 
    MgmtMplsStaticLabelStatus_rewrite_next_hop_interface_down MgmtMplsStaticLabelStatus = "rewrite-next-hop-interface-down"

    // Label Discrepancy 
    MgmtMplsStaticLabelStatus_label_discrepancy MgmtMplsStaticLabelStatus = "label-discrepancy"

    // Rewrite Discrepancy 
    MgmtMplsStaticLabelStatus_rewrite_discrepancy MgmtMplsStaticLabelStatus = "rewrite-discrepancy"

    // Label Status Unknown
    MgmtMplsStaticLabelStatus_label_status_unknown MgmtMplsStaticLabelStatus = "label-status-unknown"
)

// MgmtStaticAddr represents Mgmt static addr
type MgmtStaticAddr string

const (
    // Not Applicable
    MgmtStaticAddr_not_applicable MgmtStaticAddr = "not-applicable"

    // IPv4
    MgmtStaticAddr_ipv4 MgmtStaticAddr = "ipv4"

    // IPv6
    MgmtStaticAddr_ipv6 MgmtStaticAddr = "ipv6"
)

// MplsStaticPathRole represents Mpls static path role
type MplsStaticPathRole string

const (
    // Path is only for primary traffic
    MplsStaticPathRole_primary MplsStaticPathRole = "primary"

    // Path is only for backup traffic
    MplsStaticPathRole_backup MplsStaticPathRole = "backup"

    // Path is for primary and backup traffic
    MplsStaticPathRole_primary_and_backup MplsStaticPathRole = "primary-and-backup"
)

// MplsStatic
// MPLS STATIC operational data
type MplsStatic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs MplsStatic_Vrfs

    // MPLS STATIC summary data.
    Summary MplsStatic_Summary

    // data for static local-label table.
    LocalLabels MplsStatic_LocalLabels
}

func (mplsStatic *MplsStatic) GetFilter() yfilter.YFilter { return mplsStatic.YFilter }

func (mplsStatic *MplsStatic) SetFilter(yf yfilter.YFilter) { mplsStatic.YFilter = yf }

func (mplsStatic *MplsStatic) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    if yname == "summary" { return "Summary" }
    if yname == "local-labels" { return "LocalLabels" }
    return ""
}

func (mplsStatic *MplsStatic) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-static-oper:mpls-static"
}

func (mplsStatic *MplsStatic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &mplsStatic.Vrfs
    }
    if childYangName == "summary" {
        return &mplsStatic.Summary
    }
    if childYangName == "local-labels" {
        return &mplsStatic.LocalLabels
    }
    return nil
}

func (mplsStatic *MplsStatic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &mplsStatic.Vrfs
    children["summary"] = &mplsStatic.Summary
    children["local-labels"] = &mplsStatic.LocalLabels
    return children
}

func (mplsStatic *MplsStatic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
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

func (mplsStatic *MplsStatic) GetParentYangName() string { return "Cisco-IOS-XR-mpls-static-oper" }

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

    // data for static lsp table.
    Lsps MplsStatic_Vrfs_Vrf_Lsps

    // data for static local-label table.
    LocalLabels MplsStatic_Vrfs_Vrf_LocalLabels
}

func (vrf *MplsStatic_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *MplsStatic_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *MplsStatic_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "lsps" { return "Lsps" }
    if yname == "local-labels" { return "LocalLabels" }
    return ""
}

func (vrf *MplsStatic_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *MplsStatic_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsps" {
        return &vrf.Lsps
    }
    if childYangName == "local-labels" {
        return &vrf.LocalLabels
    }
    return nil
}

func (vrf *MplsStatic_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsps"] = &vrf.Lsps
    children["local-labels"] = &vrf.LocalLabels
    return children
}

func (vrf *MplsStatic_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
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

// MplsStatic_Vrfs_Vrf_Lsps
// data for static lsp table
type MplsStatic_Vrfs_Vrf_Lsps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for static lsp. The type is slice of MplsStatic_Vrfs_Vrf_Lsps_Lsp.
    Lsp []MplsStatic_Vrfs_Vrf_Lsps_Lsp
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetFilter() yfilter.YFilter { return lsps.YFilter }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) SetFilter(yf yfilter.YFilter) { lsps.YFilter = yf }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetGoName(yname string) string {
    if yname == "lsp" { return "Lsp" }
    return ""
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetSegmentPath() string {
    return "lsps"
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp" {
        for _, c := range lsps.Lsp {
            if lsps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Lsps_Lsp{}
        lsps.Lsp = append(lsps.Lsp, child)
        return &lsps.Lsp[len(lsps.Lsp)-1]
    }
    return nil
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range lsps.Lsp {
        children[lsps.Lsp[i].GetSegmentPath()] = &lsps.Lsp[i]
    }
    return children
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetBundleName() string { return "cisco_ios_xr" }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetYangName() string { return "lsps" }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) SetParent(parent types.Entity) { lsps.parent = parent }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetParent() types.Entity { return lsps.parent }

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetParentYangName() string { return "vrf" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp
// Data for static lsp
type MplsStatic_Vrfs_Vrf_Lsps_Lsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. LSP Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LspName interface{}

    // LSP Name. The type is string.
    LspNameXr interface{}

    // Label Information.
    Label MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetFilter() yfilter.YFilter { return lsp.YFilter }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) SetFilter(yf yfilter.YFilter) { lsp.YFilter = yf }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetGoName(yname string) string {
    if yname == "lsp-name" { return "LspName" }
    if yname == "lsp-name-xr" { return "LspNameXr" }
    if yname == "label" { return "Label" }
    return ""
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetSegmentPath() string {
    return "lsp" + "[lsp-name='" + fmt.Sprintf("%v", lsp.LspName) + "']"
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label" {
        return &lsp.Label
    }
    return nil
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["label"] = &lsp.Label
    return children
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsp-name"] = lsp.LspName
    leafs["lsp-name-xr"] = lsp.LspNameXr
    return leafs
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetBundleName() string { return "cisco_ios_xr" }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetYangName() string { return "lsp" }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) SetParent(parent types.Entity) { lsp.parent = parent }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetParent() types.Entity { return lsp.parent }

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetParentYangName() string { return "lsps" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label
// Label Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Label value. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Label Mode. The type is MgmtMplsStaticLabelMode.
    LabelMode interface{}

    // Label Status. The type is MgmtMplsStaticLabelStatus.
    LabelStatus interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Primary Pathset as a result of resolve. The type is bool.
    PathsetViaResolve interface{}

    // Backup Pathset as a result of resolve. The type is bool.
    BackupPathsetViaResolve interface{}

    // Address Family. The type is MgmtStaticAddr.
    AddressFamily interface{}

    // Prefix Information.
    Prefix MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix

    // Primary pathset resolve-nexthop IP Address.
    PathsetResolveNh MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh

    // Backup pathset resolve-nexthop IP Address.
    BackupPathsetResolveNh MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo.
    PathInfo []MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo

    // Backup Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo.
    BackupPathInfo []MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetFilter() yfilter.YFilter { return label.YFilter }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) SetFilter(yf yfilter.YFilter) { label.YFilter = yf }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "label-mode" { return "LabelMode" }
    if yname == "label-status" { return "LabelStatus" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "pathset-via-resolve" { return "PathsetViaResolve" }
    if yname == "backup-pathset-via-resolve" { return "BackupPathsetViaResolve" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "prefix" { return "Prefix" }
    if yname == "pathset-resolve-nh" { return "PathsetResolveNh" }
    if yname == "backup-pathset-resolve-nh" { return "BackupPathsetResolveNh" }
    if yname == "path-info" { return "PathInfo" }
    if yname == "backup-path-info" { return "BackupPathInfo" }
    return ""
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetSegmentPath() string {
    return "label"
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &label.Prefix
    }
    if childYangName == "pathset-resolve-nh" {
        return &label.PathsetResolveNh
    }
    if childYangName == "backup-pathset-resolve-nh" {
        return &label.BackupPathsetResolveNh
    }
    if childYangName == "path-info" {
        for _, c := range label.PathInfo {
            if label.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo{}
        label.PathInfo = append(label.PathInfo, child)
        return &label.PathInfo[len(label.PathInfo)-1]
    }
    if childYangName == "backup-path-info" {
        for _, c := range label.BackupPathInfo {
            if label.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo{}
        label.BackupPathInfo = append(label.BackupPathInfo, child)
        return &label.BackupPathInfo[len(label.BackupPathInfo)-1]
    }
    return nil
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &label.Prefix
    children["pathset-resolve-nh"] = &label.PathsetResolveNh
    children["backup-pathset-resolve-nh"] = &label.BackupPathsetResolveNh
    for i := range label.PathInfo {
        children[label.PathInfo[i].GetSegmentPath()] = &label.PathInfo[i]
    }
    for i := range label.BackupPathInfo {
        children[label.BackupPathInfo[i].GetSegmentPath()] = &label.BackupPathInfo[i]
    }
    return children
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = label.Label
    leafs["label-mode"] = label.LabelMode
    leafs["label-status"] = label.LabelStatus
    leafs["vrf-name"] = label.VrfName
    leafs["pathset-via-resolve"] = label.PathsetViaResolve
    leafs["backup-pathset-via-resolve"] = label.BackupPathsetViaResolve
    leafs["address-family"] = label.AddressFamily
    return leafs
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetBundleName() string { return "cisco_ios_xr" }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetYangName() string { return "label" }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) SetParent(parent types.Entity) { label.parent = parent }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetParent() types.Entity { return label.parent }

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetParentYangName() string { return "lsp" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix
// Prefix Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefix.Prefix
    }
    return nil
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefix.Prefix
    return children
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefix.PrefixLength
    return leafs
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetYangName() string { return "prefix" }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetParentYangName() string { return "label" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix
// Prefix
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = prefix.AfName
    leafs["ipv4-prefix"] = prefix.Ipv4Prefix
    leafs["ipv6-prefix"] = prefix.Ipv6Prefix
    return leafs
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetYangName() string { return "prefix" }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetParentYangName() string { return "prefix" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh
// Primary pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetFilter() yfilter.YFilter { return pathsetResolveNh.YFilter }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) SetFilter(yf yfilter.YFilter) { pathsetResolveNh.YFilter = yf }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetSegmentPath() string {
    return "pathset-resolve-nh"
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = pathsetResolveNh.AfName
    leafs["ipv4-prefix"] = pathsetResolveNh.Ipv4Prefix
    leafs["ipv6-prefix"] = pathsetResolveNh.Ipv6Prefix
    return leafs
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetBundleName() string { return "cisco_ios_xr" }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetYangName() string { return "pathset-resolve-nh" }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) SetParent(parent types.Entity) { pathsetResolveNh.parent = parent }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetParent() types.Entity { return pathsetResolveNh.parent }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetParentYangName() string { return "label" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh
// Backup pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetFilter() yfilter.YFilter { return backupPathsetResolveNh.YFilter }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) SetFilter(yf yfilter.YFilter) { backupPathsetResolveNh.YFilter = yf }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetSegmentPath() string {
    return "backup-pathset-resolve-nh"
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = backupPathsetResolveNh.AfName
    leafs["ipv4-prefix"] = backupPathsetResolveNh.Ipv4Prefix
    leafs["ipv6-prefix"] = backupPathsetResolveNh.Ipv6Prefix
    return leafs
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetBundleName() string { return "cisco_ios_xr" }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetYangName() string { return "backup-pathset-resolve-nh" }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) SetParent(parent types.Entity) { backupPathsetResolveNh.parent = parent }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetParent() types.Entity { return backupPathsetResolveNh.parent }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetParentYangName() string { return "label" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo
// Path Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Number. The type is interface{} with range: 0..4294967295.
    PathNumber interface{}

    // Path Type. The type is MgmtStaticPath.
    Type interface{}

    // Path Role. The type is MplsStaticPathRole.
    PathRole interface{}

    // Path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Path Backup Id. The type is interface{} with range: 0..255.
    BackupId interface{}

    // Path Status. The type is MgmtMplsStaticPathStatus.
    Status interface{}

    // Nexthop information.
    Nexthop MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop
}

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetFilter() yfilter.YFilter { return pathInfo.YFilter }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) SetFilter(yf yfilter.YFilter) { pathInfo.YFilter = yf }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetGoName(yname string) string {
    if yname == "path-number" { return "PathNumber" }
    if yname == "type" { return "Type" }
    if yname == "path-role" { return "PathRole" }
    if yname == "path-id" { return "PathId" }
    if yname == "backup-id" { return "BackupId" }
    if yname == "status" { return "Status" }
    if yname == "nexthop" { return "Nexthop" }
    return ""
}

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetSegmentPath() string {
    return "path-info"
}

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop" {
        return &pathInfo.Nexthop
    }
    return nil
}

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop"] = &pathInfo.Nexthop
    return children
}

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-number"] = pathInfo.PathNumber
    leafs["type"] = pathInfo.Type
    leafs["path-role"] = pathInfo.PathRole
    leafs["path-id"] = pathInfo.PathId
    leafs["backup-id"] = pathInfo.BackupId
    leafs["status"] = pathInfo.Status
    return leafs
}

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetYangName() string { return "path-info" }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) SetParent(parent types.Entity) { pathInfo.parent = parent }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetParent() types.Entity { return pathInfo.parent }

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetParentYangName() string { return "label" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetFilter() yfilter.YFilter { return nexthop.YFilter }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) SetFilter(yf yfilter.YFilter) { nexthop.YFilter = yf }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "address" { return "Address" }
    return ""
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetSegmentPath() string {
    return "nexthop"
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &nexthop.Address
    }
    return nil
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &nexthop.Address
    return children
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = nexthop.Label
    leafs["interface-name"] = nexthop.InterfaceName
    leafs["afi"] = nexthop.Afi
    return leafs
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetBundleName() string { return "cisco_ios_xr" }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetYangName() string { return "nexthop" }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) SetParent(parent types.Entity) { nexthop.parent = parent }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetParent() types.Entity { return nexthop.parent }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetParentYangName() string { return "path-info" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetSegmentPath() string {
    return "address"
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4-prefix"] = address.Ipv4Prefix
    leafs["ipv6-prefix"] = address.Ipv6Prefix
    return leafs
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetYangName() string { return "address" }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetParent() types.Entity { return address.parent }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetParentYangName() string { return "nexthop" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo
// Backup Path Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Number. The type is interface{} with range: 0..4294967295.
    PathNumber interface{}

    // Path Type. The type is MgmtStaticPath.
    Type interface{}

    // Path Role. The type is MplsStaticPathRole.
    PathRole interface{}

    // Path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Path Backup Id. The type is interface{} with range: 0..255.
    BackupId interface{}

    // Path Status. The type is MgmtMplsStaticPathStatus.
    Status interface{}

    // Nexthop information.
    Nexthop MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetFilter() yfilter.YFilter { return backupPathInfo.YFilter }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) SetFilter(yf yfilter.YFilter) { backupPathInfo.YFilter = yf }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetGoName(yname string) string {
    if yname == "path-number" { return "PathNumber" }
    if yname == "type" { return "Type" }
    if yname == "path-role" { return "PathRole" }
    if yname == "path-id" { return "PathId" }
    if yname == "backup-id" { return "BackupId" }
    if yname == "status" { return "Status" }
    if yname == "nexthop" { return "Nexthop" }
    return ""
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetSegmentPath() string {
    return "backup-path-info"
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop" {
        return &backupPathInfo.Nexthop
    }
    return nil
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop"] = &backupPathInfo.Nexthop
    return children
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-number"] = backupPathInfo.PathNumber
    leafs["type"] = backupPathInfo.Type
    leafs["path-role"] = backupPathInfo.PathRole
    leafs["path-id"] = backupPathInfo.PathId
    leafs["backup-id"] = backupPathInfo.BackupId
    leafs["status"] = backupPathInfo.Status
    return leafs
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetBundleName() string { return "cisco_ios_xr" }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetYangName() string { return "backup-path-info" }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) SetParent(parent types.Entity) { backupPathInfo.parent = parent }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetParent() types.Entity { return backupPathInfo.parent }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetParentYangName() string { return "label" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetFilter() yfilter.YFilter { return nexthop.YFilter }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) SetFilter(yf yfilter.YFilter) { nexthop.YFilter = yf }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "address" { return "Address" }
    return ""
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetSegmentPath() string {
    return "nexthop"
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &nexthop.Address
    }
    return nil
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &nexthop.Address
    return children
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = nexthop.Label
    leafs["interface-name"] = nexthop.InterfaceName
    leafs["afi"] = nexthop.Afi
    return leafs
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetBundleName() string { return "cisco_ios_xr" }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetYangName() string { return "nexthop" }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) SetParent(parent types.Entity) { nexthop.parent = parent }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetParent() types.Entity { return nexthop.parent }

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetParentYangName() string { return "backup-path-info" }

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetSegmentPath() string {
    return "address"
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4-prefix"] = address.Ipv4Prefix
    leafs["ipv6-prefix"] = address.Ipv6Prefix
    return leafs
}

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetYangName() string { return "address" }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetParent() types.Entity { return address.parent }

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetParentYangName() string { return "nexthop" }

// MplsStatic_Vrfs_Vrf_LocalLabels
// data for static local-label table
type MplsStatic_Vrfs_Vrf_LocalLabels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for static label. The type is slice of
    // MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetFilter() yfilter.YFilter { return localLabels.YFilter }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) SetFilter(yf yfilter.YFilter) { localLabels.YFilter = yf }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    return ""
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetSegmentPath() string {
    return "local-labels"
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label" {
        for _, c := range localLabels.LocalLabel {
            if localLabels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel{}
        localLabels.LocalLabel = append(localLabels.LocalLabel, child)
        return &localLabels.LocalLabel[len(localLabels.LocalLabel)-1]
    }
    return nil
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabels.LocalLabel {
        children[localLabels.LocalLabel[i].GetSegmentPath()] = &localLabels.LocalLabel[i]
    }
    return children
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetBundleName() string { return "cisco_ios_xr" }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetYangName() string { return "local-labels" }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) SetParent(parent types.Entity) { localLabels.parent = parent }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetParent() types.Entity { return localLabels.parent }

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetParentYangName() string { return "vrf" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel
// Data for static label
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // Label value. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Label Mode. The type is MgmtMplsStaticLabelMode.
    LabelMode interface{}

    // Label Status. The type is MgmtMplsStaticLabelStatus.
    LabelStatus interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Primary Pathset as a result of resolve. The type is bool.
    PathsetViaResolve interface{}

    // Backup Pathset as a result of resolve. The type is bool.
    BackupPathsetViaResolve interface{}

    // Address Family. The type is MgmtStaticAddr.
    AddressFamily interface{}

    // Prefix Information.
    Prefix MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix

    // Primary pathset resolve-nexthop IP Address.
    PathsetResolveNh MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh

    // Backup pathset resolve-nexthop IP Address.
    BackupPathsetResolveNh MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh

    // Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo.
    PathInfo []MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo

    // Backup Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo.
    BackupPathInfo []MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetFilter() yfilter.YFilter { return localLabel.YFilter }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) SetFilter(yf yfilter.YFilter) { localLabel.YFilter = yf }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetGoName(yname string) string {
    if yname == "local-label-id" { return "LocalLabelId" }
    if yname == "label" { return "Label" }
    if yname == "label-mode" { return "LabelMode" }
    if yname == "label-status" { return "LabelStatus" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "pathset-via-resolve" { return "PathsetViaResolve" }
    if yname == "backup-pathset-via-resolve" { return "BackupPathsetViaResolve" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "prefix" { return "Prefix" }
    if yname == "pathset-resolve-nh" { return "PathsetResolveNh" }
    if yname == "backup-pathset-resolve-nh" { return "BackupPathsetResolveNh" }
    if yname == "path-info" { return "PathInfo" }
    if yname == "backup-path-info" { return "BackupPathInfo" }
    return ""
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetSegmentPath() string {
    return "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &localLabel.Prefix
    }
    if childYangName == "pathset-resolve-nh" {
        return &localLabel.PathsetResolveNh
    }
    if childYangName == "backup-pathset-resolve-nh" {
        return &localLabel.BackupPathsetResolveNh
    }
    if childYangName == "path-info" {
        for _, c := range localLabel.PathInfo {
            if localLabel.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo{}
        localLabel.PathInfo = append(localLabel.PathInfo, child)
        return &localLabel.PathInfo[len(localLabel.PathInfo)-1]
    }
    if childYangName == "backup-path-info" {
        for _, c := range localLabel.BackupPathInfo {
            if localLabel.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo{}
        localLabel.BackupPathInfo = append(localLabel.BackupPathInfo, child)
        return &localLabel.BackupPathInfo[len(localLabel.BackupPathInfo)-1]
    }
    return nil
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &localLabel.Prefix
    children["pathset-resolve-nh"] = &localLabel.PathsetResolveNh
    children["backup-pathset-resolve-nh"] = &localLabel.BackupPathsetResolveNh
    for i := range localLabel.PathInfo {
        children[localLabel.PathInfo[i].GetSegmentPath()] = &localLabel.PathInfo[i]
    }
    for i := range localLabel.BackupPathInfo {
        children[localLabel.BackupPathInfo[i].GetSegmentPath()] = &localLabel.BackupPathInfo[i]
    }
    return children
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label-id"] = localLabel.LocalLabelId
    leafs["label"] = localLabel.Label
    leafs["label-mode"] = localLabel.LabelMode
    leafs["label-status"] = localLabel.LabelStatus
    leafs["vrf-name"] = localLabel.VrfName
    leafs["pathset-via-resolve"] = localLabel.PathsetViaResolve
    leafs["backup-pathset-via-resolve"] = localLabel.BackupPathsetViaResolve
    leafs["address-family"] = localLabel.AddressFamily
    return leafs
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetBundleName() string { return "cisco_ios_xr" }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetYangName() string { return "local-label" }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) SetParent(parent types.Entity) { localLabel.parent = parent }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetParent() types.Entity { return localLabel.parent }

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetParentYangName() string { return "local-labels" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix
// Prefix Information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefix.Prefix
    }
    return nil
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefix.Prefix
    return children
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefix.PrefixLength
    return leafs
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetYangName() string { return "prefix" }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix
// Prefix
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = prefix.AfName
    leafs["ipv4-prefix"] = prefix.Ipv4Prefix
    leafs["ipv6-prefix"] = prefix.Ipv6Prefix
    return leafs
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetYangName() string { return "prefix" }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetParentYangName() string { return "prefix" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh
// Primary pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetFilter() yfilter.YFilter { return pathsetResolveNh.YFilter }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) SetFilter(yf yfilter.YFilter) { pathsetResolveNh.YFilter = yf }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetSegmentPath() string {
    return "pathset-resolve-nh"
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = pathsetResolveNh.AfName
    leafs["ipv4-prefix"] = pathsetResolveNh.Ipv4Prefix
    leafs["ipv6-prefix"] = pathsetResolveNh.Ipv6Prefix
    return leafs
}

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetBundleName() string { return "cisco_ios_xr" }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetYangName() string { return "pathset-resolve-nh" }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) SetParent(parent types.Entity) { pathsetResolveNh.parent = parent }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetParent() types.Entity { return pathsetResolveNh.parent }

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh
// Backup pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetFilter() yfilter.YFilter { return backupPathsetResolveNh.YFilter }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) SetFilter(yf yfilter.YFilter) { backupPathsetResolveNh.YFilter = yf }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetSegmentPath() string {
    return "backup-pathset-resolve-nh"
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = backupPathsetResolveNh.AfName
    leafs["ipv4-prefix"] = backupPathsetResolveNh.Ipv4Prefix
    leafs["ipv6-prefix"] = backupPathsetResolveNh.Ipv6Prefix
    return leafs
}

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetBundleName() string { return "cisco_ios_xr" }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetYangName() string { return "backup-pathset-resolve-nh" }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) SetParent(parent types.Entity) { backupPathsetResolveNh.parent = parent }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetParent() types.Entity { return backupPathsetResolveNh.parent }

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo
// Path Information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Number. The type is interface{} with range: 0..4294967295.
    PathNumber interface{}

    // Path Type. The type is MgmtStaticPath.
    Type interface{}

    // Path Role. The type is MplsStaticPathRole.
    PathRole interface{}

    // Path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Path Backup Id. The type is interface{} with range: 0..255.
    BackupId interface{}

    // Path Status. The type is MgmtMplsStaticPathStatus.
    Status interface{}

    // Nexthop information.
    Nexthop MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop
}

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetFilter() yfilter.YFilter { return pathInfo.YFilter }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) SetFilter(yf yfilter.YFilter) { pathInfo.YFilter = yf }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetGoName(yname string) string {
    if yname == "path-number" { return "PathNumber" }
    if yname == "type" { return "Type" }
    if yname == "path-role" { return "PathRole" }
    if yname == "path-id" { return "PathId" }
    if yname == "backup-id" { return "BackupId" }
    if yname == "status" { return "Status" }
    if yname == "nexthop" { return "Nexthop" }
    return ""
}

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetSegmentPath() string {
    return "path-info"
}

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop" {
        return &pathInfo.Nexthop
    }
    return nil
}

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop"] = &pathInfo.Nexthop
    return children
}

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-number"] = pathInfo.PathNumber
    leafs["type"] = pathInfo.Type
    leafs["path-role"] = pathInfo.PathRole
    leafs["path-id"] = pathInfo.PathId
    leafs["backup-id"] = pathInfo.BackupId
    leafs["status"] = pathInfo.Status
    return leafs
}

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetYangName() string { return "path-info" }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) SetParent(parent types.Entity) { pathInfo.parent = parent }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetParent() types.Entity { return pathInfo.parent }

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetFilter() yfilter.YFilter { return nexthop.YFilter }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) SetFilter(yf yfilter.YFilter) { nexthop.YFilter = yf }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "address" { return "Address" }
    return ""
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetSegmentPath() string {
    return "nexthop"
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &nexthop.Address
    }
    return nil
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &nexthop.Address
    return children
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = nexthop.Label
    leafs["interface-name"] = nexthop.InterfaceName
    leafs["afi"] = nexthop.Afi
    return leafs
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetBundleName() string { return "cisco_ios_xr" }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetYangName() string { return "nexthop" }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) SetParent(parent types.Entity) { nexthop.parent = parent }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetParent() types.Entity { return nexthop.parent }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetParentYangName() string { return "path-info" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetSegmentPath() string {
    return "address"
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4-prefix"] = address.Ipv4Prefix
    leafs["ipv6-prefix"] = address.Ipv6Prefix
    return leafs
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetYangName() string { return "address" }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetParent() types.Entity { return address.parent }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetParentYangName() string { return "nexthop" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo
// Backup Path Information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Number. The type is interface{} with range: 0..4294967295.
    PathNumber interface{}

    // Path Type. The type is MgmtStaticPath.
    Type interface{}

    // Path Role. The type is MplsStaticPathRole.
    PathRole interface{}

    // Path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Path Backup Id. The type is interface{} with range: 0..255.
    BackupId interface{}

    // Path Status. The type is MgmtMplsStaticPathStatus.
    Status interface{}

    // Nexthop information.
    Nexthop MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetFilter() yfilter.YFilter { return backupPathInfo.YFilter }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) SetFilter(yf yfilter.YFilter) { backupPathInfo.YFilter = yf }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetGoName(yname string) string {
    if yname == "path-number" { return "PathNumber" }
    if yname == "type" { return "Type" }
    if yname == "path-role" { return "PathRole" }
    if yname == "path-id" { return "PathId" }
    if yname == "backup-id" { return "BackupId" }
    if yname == "status" { return "Status" }
    if yname == "nexthop" { return "Nexthop" }
    return ""
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetSegmentPath() string {
    return "backup-path-info"
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop" {
        return &backupPathInfo.Nexthop
    }
    return nil
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop"] = &backupPathInfo.Nexthop
    return children
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-number"] = backupPathInfo.PathNumber
    leafs["type"] = backupPathInfo.Type
    leafs["path-role"] = backupPathInfo.PathRole
    leafs["path-id"] = backupPathInfo.PathId
    leafs["backup-id"] = backupPathInfo.BackupId
    leafs["status"] = backupPathInfo.Status
    return leafs
}

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetBundleName() string { return "cisco_ios_xr" }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetYangName() string { return "backup-path-info" }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) SetParent(parent types.Entity) { backupPathInfo.parent = parent }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetParent() types.Entity { return backupPathInfo.parent }

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetParentYangName() string { return "local-label" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetFilter() yfilter.YFilter { return nexthop.YFilter }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) SetFilter(yf yfilter.YFilter) { nexthop.YFilter = yf }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "address" { return "Address" }
    return ""
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetSegmentPath() string {
    return "nexthop"
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &nexthop.Address
    }
    return nil
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &nexthop.Address
    return children
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = nexthop.Label
    leafs["interface-name"] = nexthop.InterfaceName
    leafs["afi"] = nexthop.Afi
    return leafs
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetBundleName() string { return "cisco_ios_xr" }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetYangName() string { return "nexthop" }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) SetParent(parent types.Entity) { nexthop.parent = parent }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetParent() types.Entity { return nexthop.parent }

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetParentYangName() string { return "backup-path-info" }

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetSegmentPath() string {
    return "address"
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4-prefix"] = address.Ipv4Prefix
    leafs["ipv6-prefix"] = address.Ipv6Prefix
    return leafs
}

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetYangName() string { return "address" }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetParent() types.Entity { return address.parent }

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetParentYangName() string { return "nexthop" }

// MplsStatic_Summary
// MPLS STATIC summary data
type MplsStatic_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total Number of Labels. The type is interface{} with range: 0..4294967295.
    LabelCount interface{}

    // Total Number of Labels with Errors. The type is interface{} with range:
    // 0..4294967295.
    LabelErrorCount interface{}

    // Total Number of Labels with Discrepancies. The type is interface{} with
    // range: 0..4294967295.
    LabelDiscrepancyCount interface{}

    // Total Number of VRF configured. The type is interface{} with range:
    // 0..4294967295.
    VrfCount interface{}

    // Total Number of Active VRF Active. The type is interface{} with range:
    // 0..4294967295.
    ActiveVrfCount interface{}

    // Total Number of Interface. The type is interface{} with range:
    // 0..4294967295.
    InterfaceCount interface{}

    // Total Number of Active Interface. The type is interface{} with range:
    // 0..4294967295.
    InterfaceFowardReferenceCount interface{}

    // Total Number of MPLS enabled Interface. The type is interface{} with range:
    // 0..4294967295.
    MplsEnabledInterfaceCount interface{}

    // Total Number of IPv4 Routes. The type is interface{} with range:
    // 0..4294967295.
    Ipv4RouteCount interface{}

    // Total Number of IPv6 Routes. The type is interface{} with range:
    // 0..4294967295.
    Ipv6RouteCount interface{}

    // LSD connection is up. The type is bool.
    LsdConnected interface{}

    // IM is connected. The type is bool.
    ImConnected interface{}

    // RSI is connected. The type is bool.
    RsiConnected interface{}

    // RIBv4 is connected. The type is bool.
    Ribv4Connected interface{}

    // RIBv6 is connected. The type is bool.
    Ribv6Connected interface{}
}

func (summary *MplsStatic_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *MplsStatic_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *MplsStatic_Summary) GetGoName(yname string) string {
    if yname == "label-count" { return "LabelCount" }
    if yname == "label-error-count" { return "LabelErrorCount" }
    if yname == "label-discrepancy-count" { return "LabelDiscrepancyCount" }
    if yname == "vrf-count" { return "VrfCount" }
    if yname == "active-vrf-count" { return "ActiveVrfCount" }
    if yname == "interface-count" { return "InterfaceCount" }
    if yname == "interface-foward-reference-count" { return "InterfaceFowardReferenceCount" }
    if yname == "mpls-enabled-interface-count" { return "MplsEnabledInterfaceCount" }
    if yname == "ipv4-route-count" { return "Ipv4RouteCount" }
    if yname == "ipv6-route-count" { return "Ipv6RouteCount" }
    if yname == "lsd-connected" { return "LsdConnected" }
    if yname == "im-connected" { return "ImConnected" }
    if yname == "rsi-connected" { return "RsiConnected" }
    if yname == "ribv4-connected" { return "Ribv4Connected" }
    if yname == "ribv6-connected" { return "Ribv6Connected" }
    return ""
}

func (summary *MplsStatic_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *MplsStatic_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *MplsStatic_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *MplsStatic_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label-count"] = summary.LabelCount
    leafs["label-error-count"] = summary.LabelErrorCount
    leafs["label-discrepancy-count"] = summary.LabelDiscrepancyCount
    leafs["vrf-count"] = summary.VrfCount
    leafs["active-vrf-count"] = summary.ActiveVrfCount
    leafs["interface-count"] = summary.InterfaceCount
    leafs["interface-foward-reference-count"] = summary.InterfaceFowardReferenceCount
    leafs["mpls-enabled-interface-count"] = summary.MplsEnabledInterfaceCount
    leafs["ipv4-route-count"] = summary.Ipv4RouteCount
    leafs["ipv6-route-count"] = summary.Ipv6RouteCount
    leafs["lsd-connected"] = summary.LsdConnected
    leafs["im-connected"] = summary.ImConnected
    leafs["rsi-connected"] = summary.RsiConnected
    leafs["ribv4-connected"] = summary.Ribv4Connected
    leafs["ribv6-connected"] = summary.Ribv6Connected
    return leafs
}

func (summary *MplsStatic_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *MplsStatic_Summary) GetYangName() string { return "summary" }

func (summary *MplsStatic_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *MplsStatic_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *MplsStatic_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *MplsStatic_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *MplsStatic_Summary) GetParent() types.Entity { return summary.parent }

func (summary *MplsStatic_Summary) GetParentYangName() string { return "mpls-static" }

// MplsStatic_LocalLabels
// data for static local-label table
type MplsStatic_LocalLabels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Data for static label. The type is slice of
    // MplsStatic_LocalLabels_LocalLabel.
    LocalLabel []MplsStatic_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_LocalLabels) GetFilter() yfilter.YFilter { return localLabels.YFilter }

func (localLabels *MplsStatic_LocalLabels) SetFilter(yf yfilter.YFilter) { localLabels.YFilter = yf }

func (localLabels *MplsStatic_LocalLabels) GetGoName(yname string) string {
    if yname == "local-label" { return "LocalLabel" }
    return ""
}

func (localLabels *MplsStatic_LocalLabels) GetSegmentPath() string {
    return "local-labels"
}

func (localLabels *MplsStatic_LocalLabels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-label" {
        for _, c := range localLabels.LocalLabel {
            if localLabels.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_LocalLabels_LocalLabel{}
        localLabels.LocalLabel = append(localLabels.LocalLabel, child)
        return &localLabels.LocalLabel[len(localLabels.LocalLabel)-1]
    }
    return nil
}

func (localLabels *MplsStatic_LocalLabels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localLabels.LocalLabel {
        children[localLabels.LocalLabel[i].GetSegmentPath()] = &localLabels.LocalLabel[i]
    }
    return children
}

func (localLabels *MplsStatic_LocalLabels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localLabels *MplsStatic_LocalLabels) GetBundleName() string { return "cisco_ios_xr" }

func (localLabels *MplsStatic_LocalLabels) GetYangName() string { return "local-labels" }

func (localLabels *MplsStatic_LocalLabels) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabels *MplsStatic_LocalLabels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabels *MplsStatic_LocalLabels) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabels *MplsStatic_LocalLabels) SetParent(parent types.Entity) { localLabels.parent = parent }

func (localLabels *MplsStatic_LocalLabels) GetParent() types.Entity { return localLabels.parent }

func (localLabels *MplsStatic_LocalLabels) GetParentYangName() string { return "mpls-static" }

// MplsStatic_LocalLabels_LocalLabel
// Data for static label
type MplsStatic_LocalLabels_LocalLabel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local Label. The type is interface{} with range:
    // 16..1048575.
    LocalLabelId interface{}

    // Label value. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Label Mode. The type is MgmtMplsStaticLabelMode.
    LabelMode interface{}

    // Label Status. The type is MgmtMplsStaticLabelStatus.
    LabelStatus interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Primary Pathset as a result of resolve. The type is bool.
    PathsetViaResolve interface{}

    // Backup Pathset as a result of resolve. The type is bool.
    BackupPathsetViaResolve interface{}

    // Address Family. The type is MgmtStaticAddr.
    AddressFamily interface{}

    // Prefix Information.
    Prefix MplsStatic_LocalLabels_LocalLabel_Prefix

    // Primary pathset resolve-nexthop IP Address.
    PathsetResolveNh MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh

    // Backup pathset resolve-nexthop IP Address.
    BackupPathsetResolveNh MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh

    // Path Information. The type is slice of
    // MplsStatic_LocalLabels_LocalLabel_PathInfo.
    PathInfo []MplsStatic_LocalLabels_LocalLabel_PathInfo

    // Backup Path Information. The type is slice of
    // MplsStatic_LocalLabels_LocalLabel_BackupPathInfo.
    BackupPathInfo []MplsStatic_LocalLabels_LocalLabel_BackupPathInfo
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetFilter() yfilter.YFilter { return localLabel.YFilter }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) SetFilter(yf yfilter.YFilter) { localLabel.YFilter = yf }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetGoName(yname string) string {
    if yname == "local-label-id" { return "LocalLabelId" }
    if yname == "label" { return "Label" }
    if yname == "label-mode" { return "LabelMode" }
    if yname == "label-status" { return "LabelStatus" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "pathset-via-resolve" { return "PathsetViaResolve" }
    if yname == "backup-pathset-via-resolve" { return "BackupPathsetViaResolve" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "prefix" { return "Prefix" }
    if yname == "pathset-resolve-nh" { return "PathsetResolveNh" }
    if yname == "backup-pathset-resolve-nh" { return "BackupPathsetResolveNh" }
    if yname == "path-info" { return "PathInfo" }
    if yname == "backup-path-info" { return "BackupPathInfo" }
    return ""
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetSegmentPath() string {
    return "local-label" + "[local-label-id='" + fmt.Sprintf("%v", localLabel.LocalLabelId) + "']"
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &localLabel.Prefix
    }
    if childYangName == "pathset-resolve-nh" {
        return &localLabel.PathsetResolveNh
    }
    if childYangName == "backup-pathset-resolve-nh" {
        return &localLabel.BackupPathsetResolveNh
    }
    if childYangName == "path-info" {
        for _, c := range localLabel.PathInfo {
            if localLabel.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_LocalLabels_LocalLabel_PathInfo{}
        localLabel.PathInfo = append(localLabel.PathInfo, child)
        return &localLabel.PathInfo[len(localLabel.PathInfo)-1]
    }
    if childYangName == "backup-path-info" {
        for _, c := range localLabel.BackupPathInfo {
            if localLabel.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsStatic_LocalLabels_LocalLabel_BackupPathInfo{}
        localLabel.BackupPathInfo = append(localLabel.BackupPathInfo, child)
        return &localLabel.BackupPathInfo[len(localLabel.BackupPathInfo)-1]
    }
    return nil
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &localLabel.Prefix
    children["pathset-resolve-nh"] = &localLabel.PathsetResolveNh
    children["backup-pathset-resolve-nh"] = &localLabel.BackupPathsetResolveNh
    for i := range localLabel.PathInfo {
        children[localLabel.PathInfo[i].GetSegmentPath()] = &localLabel.PathInfo[i]
    }
    for i := range localLabel.BackupPathInfo {
        children[localLabel.BackupPathInfo[i].GetSegmentPath()] = &localLabel.BackupPathInfo[i]
    }
    return children
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-label-id"] = localLabel.LocalLabelId
    leafs["label"] = localLabel.Label
    leafs["label-mode"] = localLabel.LabelMode
    leafs["label-status"] = localLabel.LabelStatus
    leafs["vrf-name"] = localLabel.VrfName
    leafs["pathset-via-resolve"] = localLabel.PathsetViaResolve
    leafs["backup-pathset-via-resolve"] = localLabel.BackupPathsetViaResolve
    leafs["address-family"] = localLabel.AddressFamily
    return leafs
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetBundleName() string { return "cisco_ios_xr" }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetYangName() string { return "local-label" }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) SetParent(parent types.Entity) { localLabel.parent = parent }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetParent() types.Entity { return localLabel.parent }

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetParentYangName() string { return "local-labels" }

// MplsStatic_LocalLabels_LocalLabel_Prefix
// Prefix Information
type MplsStatic_LocalLabels_LocalLabel_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetGoName(yname string) string {
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "prefix" { return "Prefix" }
    return ""
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix" {
        return &prefix.Prefix
    }
    return nil
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["prefix"] = &prefix.Prefix
    return children
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-length"] = prefix.PrefixLength
    return leafs
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetYangName() string { return "prefix" }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetParentYangName() string { return "local-label" }

// MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix
// Prefix
type MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetFilter() yfilter.YFilter { return prefix.YFilter }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) SetFilter(yf yfilter.YFilter) { prefix.YFilter = yf }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetSegmentPath() string {
    return "prefix"
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = prefix.AfName
    leafs["ipv4-prefix"] = prefix.Ipv4Prefix
    leafs["ipv6-prefix"] = prefix.Ipv6Prefix
    return leafs
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetBundleName() string { return "cisco_ios_xr" }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetYangName() string { return "prefix" }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) SetParent(parent types.Entity) { prefix.parent = parent }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetParent() types.Entity { return prefix.parent }

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetParentYangName() string { return "prefix" }

// MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh
// Primary pathset resolve-nexthop IP Address
type MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetFilter() yfilter.YFilter { return pathsetResolveNh.YFilter }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) SetFilter(yf yfilter.YFilter) { pathsetResolveNh.YFilter = yf }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetSegmentPath() string {
    return "pathset-resolve-nh"
}

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = pathsetResolveNh.AfName
    leafs["ipv4-prefix"] = pathsetResolveNh.Ipv4Prefix
    leafs["ipv6-prefix"] = pathsetResolveNh.Ipv6Prefix
    return leafs
}

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetBundleName() string { return "cisco_ios_xr" }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetYangName() string { return "pathset-resolve-nh" }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) SetParent(parent types.Entity) { pathsetResolveNh.parent = parent }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetParent() types.Entity { return pathsetResolveNh.parent }

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetParentYangName() string { return "local-label" }

// MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh
// Backup pathset resolve-nexthop IP Address
type MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetFilter() yfilter.YFilter { return backupPathsetResolveNh.YFilter }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) SetFilter(yf yfilter.YFilter) { backupPathsetResolveNh.YFilter = yf }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetSegmentPath() string {
    return "backup-pathset-resolve-nh"
}

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = backupPathsetResolveNh.AfName
    leafs["ipv4-prefix"] = backupPathsetResolveNh.Ipv4Prefix
    leafs["ipv6-prefix"] = backupPathsetResolveNh.Ipv6Prefix
    return leafs
}

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetBundleName() string { return "cisco_ios_xr" }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetYangName() string { return "backup-pathset-resolve-nh" }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) SetParent(parent types.Entity) { backupPathsetResolveNh.parent = parent }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetParent() types.Entity { return backupPathsetResolveNh.parent }

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetParentYangName() string { return "local-label" }

// MplsStatic_LocalLabels_LocalLabel_PathInfo
// Path Information
type MplsStatic_LocalLabels_LocalLabel_PathInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Number. The type is interface{} with range: 0..4294967295.
    PathNumber interface{}

    // Path Type. The type is MgmtStaticPath.
    Type interface{}

    // Path Role. The type is MplsStaticPathRole.
    PathRole interface{}

    // Path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Path Backup Id. The type is interface{} with range: 0..255.
    BackupId interface{}

    // Path Status. The type is MgmtMplsStaticPathStatus.
    Status interface{}

    // Nexthop information.
    Nexthop MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop
}

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetFilter() yfilter.YFilter { return pathInfo.YFilter }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) SetFilter(yf yfilter.YFilter) { pathInfo.YFilter = yf }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetGoName(yname string) string {
    if yname == "path-number" { return "PathNumber" }
    if yname == "type" { return "Type" }
    if yname == "path-role" { return "PathRole" }
    if yname == "path-id" { return "PathId" }
    if yname == "backup-id" { return "BackupId" }
    if yname == "status" { return "Status" }
    if yname == "nexthop" { return "Nexthop" }
    return ""
}

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetSegmentPath() string {
    return "path-info"
}

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop" {
        return &pathInfo.Nexthop
    }
    return nil
}

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop"] = &pathInfo.Nexthop
    return children
}

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-number"] = pathInfo.PathNumber
    leafs["type"] = pathInfo.Type
    leafs["path-role"] = pathInfo.PathRole
    leafs["path-id"] = pathInfo.PathId
    leafs["backup-id"] = pathInfo.BackupId
    leafs["status"] = pathInfo.Status
    return leafs
}

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetBundleName() string { return "cisco_ios_xr" }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetYangName() string { return "path-info" }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) SetParent(parent types.Entity) { pathInfo.parent = parent }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetParent() types.Entity { return pathInfo.parent }

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetParentYangName() string { return "local-label" }

// MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop
// Nexthop information
type MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetFilter() yfilter.YFilter { return nexthop.YFilter }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) SetFilter(yf yfilter.YFilter) { nexthop.YFilter = yf }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "address" { return "Address" }
    return ""
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetSegmentPath() string {
    return "nexthop"
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &nexthop.Address
    }
    return nil
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &nexthop.Address
    return children
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = nexthop.Label
    leafs["interface-name"] = nexthop.InterfaceName
    leafs["afi"] = nexthop.Afi
    return leafs
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetBundleName() string { return "cisco_ios_xr" }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetYangName() string { return "nexthop" }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) SetParent(parent types.Entity) { nexthop.parent = parent }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetParent() types.Entity { return nexthop.parent }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetParentYangName() string { return "path-info" }

// MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetSegmentPath() string {
    return "address"
}

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4-prefix"] = address.Ipv4Prefix
    leafs["ipv6-prefix"] = address.Ipv6Prefix
    return leafs
}

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetYangName() string { return "address" }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetParent() types.Entity { return address.parent }

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetParentYangName() string { return "nexthop" }

// MplsStatic_LocalLabels_LocalLabel_BackupPathInfo
// Backup Path Information
type MplsStatic_LocalLabels_LocalLabel_BackupPathInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Path Number. The type is interface{} with range: 0..4294967295.
    PathNumber interface{}

    // Path Type. The type is MgmtStaticPath.
    Type interface{}

    // Path Role. The type is MplsStaticPathRole.
    PathRole interface{}

    // Path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Path Backup Id. The type is interface{} with range: 0..255.
    BackupId interface{}

    // Path Status. The type is MgmtMplsStaticPathStatus.
    Status interface{}

    // Nexthop information.
    Nexthop MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop
}

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetFilter() yfilter.YFilter { return backupPathInfo.YFilter }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) SetFilter(yf yfilter.YFilter) { backupPathInfo.YFilter = yf }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetGoName(yname string) string {
    if yname == "path-number" { return "PathNumber" }
    if yname == "type" { return "Type" }
    if yname == "path-role" { return "PathRole" }
    if yname == "path-id" { return "PathId" }
    if yname == "backup-id" { return "BackupId" }
    if yname == "status" { return "Status" }
    if yname == "nexthop" { return "Nexthop" }
    return ""
}

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetSegmentPath() string {
    return "backup-path-info"
}

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop" {
        return &backupPathInfo.Nexthop
    }
    return nil
}

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop"] = &backupPathInfo.Nexthop
    return children
}

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["path-number"] = backupPathInfo.PathNumber
    leafs["type"] = backupPathInfo.Type
    leafs["path-role"] = backupPathInfo.PathRole
    leafs["path-id"] = backupPathInfo.PathId
    leafs["backup-id"] = backupPathInfo.BackupId
    leafs["status"] = backupPathInfo.Status
    return leafs
}

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetBundleName() string { return "cisco_ios_xr" }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetYangName() string { return "backup-path-info" }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) SetParent(parent types.Entity) { backupPathInfo.parent = parent }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetParent() types.Entity { return backupPathInfo.parent }

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetParentYangName() string { return "local-label" }

// MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop
// Nexthop information
type MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is interface{} with range: 0..4294967295.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetFilter() yfilter.YFilter { return nexthop.YFilter }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) SetFilter(yf yfilter.YFilter) { nexthop.YFilter = yf }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetGoName(yname string) string {
    if yname == "label" { return "Label" }
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "afi" { return "Afi" }
    if yname == "address" { return "Address" }
    return ""
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetSegmentPath() string {
    return "nexthop"
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address" {
        return &nexthop.Address
    }
    return nil
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["address"] = &nexthop.Address
    return children
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["label"] = nexthop.Label
    leafs["interface-name"] = nexthop.InterfaceName
    leafs["afi"] = nexthop.Afi
    return leafs
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetBundleName() string { return "cisco_ios_xr" }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetYangName() string { return "nexthop" }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) SetParent(parent types.Entity) { nexthop.parent = parent }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetParent() types.Entity { return nexthop.parent }

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetParentYangName() string { return "backup-path-info" }

// MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is MgmtStaticAddr.
    AfName interface{}

    // IPv4 context. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 context. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetSegmentPath() string {
    return "address"
}

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4-prefix"] = address.Ipv4Prefix
    leafs["ipv6-prefix"] = address.Ipv6Prefix
    return leafs
}

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetYangName() string { return "address" }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetParent() types.Entity { return address.parent }

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetParentYangName() string { return "nexthop" }

