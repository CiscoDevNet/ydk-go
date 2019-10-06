// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-static package operational data.
// 
// This module contains definitions
// for the following management objects:
//   mpls-static: MPLS STATIC operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// MgmtStaticLspAfi represents Mgmt static lsp afi
type MgmtStaticLspAfi string

const (
    // Not Applicable
    MgmtStaticLspAfi_not_applicable MgmtStaticLspAfi = "not-applicable"

    // IPv4
    MgmtStaticLspAfi_ipv4 MgmtStaticLspAfi = "ipv4"

    // IPv6
    MgmtStaticLspAfi_ipv6 MgmtStaticLspAfi = "ipv6"
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

    // Rewrite Nexthop Unresolved
    MgmtMplsStaticLabelStatus_rewrite_nexthop_unresolved MgmtMplsStaticLabelStatus = "rewrite-nexthop-unresolved"

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF table.
    Vrfs MplsStatic_Vrfs

    // MPLS STATIC summary data.
    Summary MplsStatic_Summary

    // data for static local-label table.
    LocalLabels MplsStatic_LocalLabels
}

func (mplsStatic *MplsStatic) GetEntityData() *types.CommonEntityData {
    mplsStatic.EntityData.YFilter = mplsStatic.YFilter
    mplsStatic.EntityData.YangName = "mpls-static"
    mplsStatic.EntityData.BundleName = "cisco_ios_xr"
    mplsStatic.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-static-oper"
    mplsStatic.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-static-oper:mpls-static"
    mplsStatic.EntityData.AbsolutePath = mplsStatic.EntityData.SegmentPath
    mplsStatic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsStatic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsStatic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsStatic.EntityData.Children = types.NewOrderedMap()
    mplsStatic.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &mplsStatic.Vrfs})
    mplsStatic.EntityData.Children.Append("summary", types.YChild{"Summary", &mplsStatic.Summary})
    mplsStatic.EntityData.Children.Append("local-labels", types.YChild{"LocalLabels", &mplsStatic.LocalLabels})
    mplsStatic.EntityData.Leafs = types.NewOrderedMap()

    mplsStatic.EntityData.YListKeys = []string {}

    return &(mplsStatic.EntityData)
}

// MplsStatic_Vrfs
// VRF table
type MplsStatic_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of MplsStatic_Vrfs_Vrf.
    Vrf []*MplsStatic_Vrfs_Vrf
}

func (vrfs *MplsStatic_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "mpls-static"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// MplsStatic_Vrfs_Vrf
// VRF Name
type MplsStatic_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // data for static lsp table.
    Lsps MplsStatic_Vrfs_Vrf_Lsps

    // data for static local-label table.
    LocalLabels MplsStatic_Vrfs_Vrf_LocalLabels
}

func (vrf *MplsStatic_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("lsps", types.YChild{"Lsps", &vrf.Lsps})
    vrf.EntityData.Children.Append("local-labels", types.YChild{"LocalLabels", &vrf.LocalLabels})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps
// data for static lsp table
type MplsStatic_Vrfs_Vrf_Lsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for static lsp. The type is slice of MplsStatic_Vrfs_Vrf_Lsps_Lsp.
    Lsp []*MplsStatic_Vrfs_Vrf_Lsps_Lsp
}

func (lsps *MplsStatic_Vrfs_Vrf_Lsps) GetEntityData() *types.CommonEntityData {
    lsps.EntityData.YFilter = lsps.YFilter
    lsps.EntityData.YangName = "lsps"
    lsps.EntityData.BundleName = "cisco_ios_xr"
    lsps.EntityData.ParentYangName = "vrf"
    lsps.EntityData.SegmentPath = "lsps"
    lsps.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/" + lsps.EntityData.SegmentPath
    lsps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsps.EntityData.Children = types.NewOrderedMap()
    lsps.EntityData.Children.Append("lsp", types.YChild{"Lsp", nil})
    for i := range lsps.Lsp {
        lsps.EntityData.Children.Append(types.GetSegmentPath(lsps.Lsp[i]), types.YChild{"Lsp", lsps.Lsp[i]})
    }
    lsps.EntityData.Leafs = types.NewOrderedMap()

    lsps.EntityData.YListKeys = []string {}

    return &(lsps.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp
// Data for static lsp
type MplsStatic_Vrfs_Vrf_Lsps_Lsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LSP Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LspName interface{}

    // LSP Name. The type is string.
    LspNameXr interface{}

    // Label Information.
    Label MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label
}

func (lsp *MplsStatic_Vrfs_Vrf_Lsps_Lsp) GetEntityData() *types.CommonEntityData {
    lsp.EntityData.YFilter = lsp.YFilter
    lsp.EntityData.YangName = "lsp"
    lsp.EntityData.BundleName = "cisco_ios_xr"
    lsp.EntityData.ParentYangName = "lsps"
    lsp.EntityData.SegmentPath = "lsp" + types.AddKeyToken(lsp.LspName, "lsp-name")
    lsp.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/" + lsp.EntityData.SegmentPath
    lsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsp.EntityData.Children = types.NewOrderedMap()
    lsp.EntityData.Children.Append("label", types.YChild{"Label", &lsp.Label})
    lsp.EntityData.Leafs = types.NewOrderedMap()
    lsp.EntityData.Leafs.Append("lsp-name", types.YLeaf{"LspName", lsp.LspName})
    lsp.EntityData.Leafs.Append("lsp-name-xr", types.YLeaf{"LspNameXr", lsp.LspNameXr})

    lsp.EntityData.YListKeys = []string {"LspName"}

    return &(lsp.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label
// Label Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label struct {
    EntityData types.CommonEntityData
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
    PathInfo []*MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo

    // Backup Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo.
    BackupPathInfo []*MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo
}

func (label *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label) GetEntityData() *types.CommonEntityData {
    label.EntityData.YFilter = label.YFilter
    label.EntityData.YangName = "label"
    label.EntityData.BundleName = "cisco_ios_xr"
    label.EntityData.ParentYangName = "lsp"
    label.EntityData.SegmentPath = "label"
    label.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/" + label.EntityData.SegmentPath
    label.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    label.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    label.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    label.EntityData.Children = types.NewOrderedMap()
    label.EntityData.Children.Append("prefix", types.YChild{"Prefix", &label.Prefix})
    label.EntityData.Children.Append("pathset-resolve-nh", types.YChild{"PathsetResolveNh", &label.PathsetResolveNh})
    label.EntityData.Children.Append("backup-pathset-resolve-nh", types.YChild{"BackupPathsetResolveNh", &label.BackupPathsetResolveNh})
    label.EntityData.Children.Append("path-info", types.YChild{"PathInfo", nil})
    for i := range label.PathInfo {
        types.SetYListKey(label.PathInfo[i], i)
        label.EntityData.Children.Append(types.GetSegmentPath(label.PathInfo[i]), types.YChild{"PathInfo", label.PathInfo[i]})
    }
    label.EntityData.Children.Append("backup-path-info", types.YChild{"BackupPathInfo", nil})
    for i := range label.BackupPathInfo {
        types.SetYListKey(label.BackupPathInfo[i], i)
        label.EntityData.Children.Append(types.GetSegmentPath(label.BackupPathInfo[i]), types.YChild{"BackupPathInfo", label.BackupPathInfo[i]})
    }
    label.EntityData.Leafs = types.NewOrderedMap()
    label.EntityData.Leafs.Append("label", types.YLeaf{"Label", label.Label})
    label.EntityData.Leafs.Append("label-mode", types.YLeaf{"LabelMode", label.LabelMode})
    label.EntityData.Leafs.Append("label-status", types.YLeaf{"LabelStatus", label.LabelStatus})
    label.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", label.VrfName})
    label.EntityData.Leafs.Append("pathset-via-resolve", types.YLeaf{"PathsetViaResolve", label.PathsetViaResolve})
    label.EntityData.Leafs.Append("backup-pathset-via-resolve", types.YLeaf{"BackupPathsetViaResolve", label.BackupPathsetViaResolve})
    label.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", label.AddressFamily})

    label.EntityData.YListKeys = []string {}

    return &(label.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix
// Prefix Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix
}

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "label"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefix.Prefix})
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefix.PrefixLength})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix
// Prefix
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_Prefix_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/prefix/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", prefix.AfName})
    prefix.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", prefix.Ipv4Prefix})
    prefix.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", prefix.Ipv6Prefix})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh
// Primary pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh struct {
    EntityData types.CommonEntityData
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

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathsetResolveNh) GetEntityData() *types.CommonEntityData {
    pathsetResolveNh.EntityData.YFilter = pathsetResolveNh.YFilter
    pathsetResolveNh.EntityData.YangName = "pathset-resolve-nh"
    pathsetResolveNh.EntityData.BundleName = "cisco_ios_xr"
    pathsetResolveNh.EntityData.ParentYangName = "label"
    pathsetResolveNh.EntityData.SegmentPath = "pathset-resolve-nh"
    pathsetResolveNh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/" + pathsetResolveNh.EntityData.SegmentPath
    pathsetResolveNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathsetResolveNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathsetResolveNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathsetResolveNh.EntityData.Children = types.NewOrderedMap()
    pathsetResolveNh.EntityData.Leafs = types.NewOrderedMap()
    pathsetResolveNh.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pathsetResolveNh.AfName})
    pathsetResolveNh.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", pathsetResolveNh.Ipv4Prefix})
    pathsetResolveNh.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", pathsetResolveNh.Ipv6Prefix})

    pathsetResolveNh.EntityData.YListKeys = []string {}

    return &(pathsetResolveNh.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh
// Backup pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh struct {
    EntityData types.CommonEntityData
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

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathsetResolveNh) GetEntityData() *types.CommonEntityData {
    backupPathsetResolveNh.EntityData.YFilter = backupPathsetResolveNh.YFilter
    backupPathsetResolveNh.EntityData.YangName = "backup-pathset-resolve-nh"
    backupPathsetResolveNh.EntityData.BundleName = "cisco_ios_xr"
    backupPathsetResolveNh.EntityData.ParentYangName = "label"
    backupPathsetResolveNh.EntityData.SegmentPath = "backup-pathset-resolve-nh"
    backupPathsetResolveNh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/" + backupPathsetResolveNh.EntityData.SegmentPath
    backupPathsetResolveNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPathsetResolveNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPathsetResolveNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPathsetResolveNh.EntityData.Children = types.NewOrderedMap()
    backupPathsetResolveNh.EntityData.Leafs = types.NewOrderedMap()
    backupPathsetResolveNh.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", backupPathsetResolveNh.AfName})
    backupPathsetResolveNh.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", backupPathsetResolveNh.Ipv4Prefix})
    backupPathsetResolveNh.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", backupPathsetResolveNh.Ipv6Prefix})

    backupPathsetResolveNh.EntityData.YListKeys = []string {}

    return &(backupPathsetResolveNh.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo
// Path Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (pathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo) GetEntityData() *types.CommonEntityData {
    pathInfo.EntityData.YFilter = pathInfo.YFilter
    pathInfo.EntityData.YangName = "path-info"
    pathInfo.EntityData.BundleName = "cisco_ios_xr"
    pathInfo.EntityData.ParentYangName = "label"
    pathInfo.EntityData.SegmentPath = "path-info" + types.AddNoKeyToken(pathInfo)
    pathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/" + pathInfo.EntityData.SegmentPath
    pathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfo.EntityData.Children = types.NewOrderedMap()
    pathInfo.EntityData.Children.Append("nexthop", types.YChild{"Nexthop", &pathInfo.Nexthop})
    pathInfo.EntityData.Leafs = types.NewOrderedMap()
    pathInfo.EntityData.Leafs.Append("path-number", types.YLeaf{"PathNumber", pathInfo.PathNumber})
    pathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathInfo.Type})
    pathInfo.EntityData.Leafs.Append("path-role", types.YLeaf{"PathRole", pathInfo.PathRole})
    pathInfo.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", pathInfo.PathId})
    pathInfo.EntityData.Leafs.Append("backup-id", types.YLeaf{"BackupId", pathInfo.BackupId})
    pathInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", pathInfo.Status})

    pathInfo.EntityData.YListKeys = []string {}

    return &(pathInfo.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is MgmtStaticLspAfi.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "path-info"
    nexthop.EntityData.SegmentPath = "nexthop"
    nexthop.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/path-info/" + nexthop.EntityData.SegmentPath
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = types.NewOrderedMap()
    nexthop.EntityData.Children.Append("address", types.YChild{"Address", &nexthop.Address})
    nexthop.EntityData.Leafs = types.NewOrderedMap()
    nexthop.EntityData.Leafs.Append("label", types.YLeaf{"Label", nexthop.Label})
    nexthop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nexthop.InterfaceName})
    nexthop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nexthop.Afi})

    nexthop.EntityData.YListKeys = []string {}

    return &(nexthop.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address struct {
    EntityData types.CommonEntityData
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

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_PathInfo_Nexthop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "nexthop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/path-info/nexthop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", address.AfName})
    address.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", address.Ipv4Prefix})
    address.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", address.Ipv6Prefix})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo
// Backup Path Information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (backupPathInfo *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo) GetEntityData() *types.CommonEntityData {
    backupPathInfo.EntityData.YFilter = backupPathInfo.YFilter
    backupPathInfo.EntityData.YangName = "backup-path-info"
    backupPathInfo.EntityData.BundleName = "cisco_ios_xr"
    backupPathInfo.EntityData.ParentYangName = "label"
    backupPathInfo.EntityData.SegmentPath = "backup-path-info" + types.AddNoKeyToken(backupPathInfo)
    backupPathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/" + backupPathInfo.EntityData.SegmentPath
    backupPathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPathInfo.EntityData.Children = types.NewOrderedMap()
    backupPathInfo.EntityData.Children.Append("nexthop", types.YChild{"Nexthop", &backupPathInfo.Nexthop})
    backupPathInfo.EntityData.Leafs = types.NewOrderedMap()
    backupPathInfo.EntityData.Leafs.Append("path-number", types.YLeaf{"PathNumber", backupPathInfo.PathNumber})
    backupPathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", backupPathInfo.Type})
    backupPathInfo.EntityData.Leafs.Append("path-role", types.YLeaf{"PathRole", backupPathInfo.PathRole})
    backupPathInfo.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", backupPathInfo.PathId})
    backupPathInfo.EntityData.Leafs.Append("backup-id", types.YLeaf{"BackupId", backupPathInfo.BackupId})
    backupPathInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", backupPathInfo.Status})

    backupPathInfo.EntityData.YListKeys = []string {}

    return &(backupPathInfo.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is MgmtStaticLspAfi.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "backup-path-info"
    nexthop.EntityData.SegmentPath = "nexthop"
    nexthop.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/backup-path-info/" + nexthop.EntityData.SegmentPath
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = types.NewOrderedMap()
    nexthop.EntityData.Children.Append("address", types.YChild{"Address", &nexthop.Address})
    nexthop.EntityData.Leafs = types.NewOrderedMap()
    nexthop.EntityData.Leafs.Append("label", types.YLeaf{"Label", nexthop.Label})
    nexthop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nexthop.InterfaceName})
    nexthop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nexthop.Afi})

    nexthop.EntityData.YListKeys = []string {}

    return &(nexthop.EntityData)
}

// MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address struct {
    EntityData types.CommonEntityData
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

func (address *MplsStatic_Vrfs_Vrf_Lsps_Lsp_Label_BackupPathInfo_Nexthop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "nexthop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/lsps/lsp/label/backup-path-info/nexthop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", address.AfName})
    address.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", address.Ipv4Prefix})
    address.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", address.Ipv6Prefix})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels
// data for static local-label table
type MplsStatic_Vrfs_Vrf_LocalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for static label. The type is slice of
    // MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel.
    LocalLabel []*MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_Vrfs_Vrf_LocalLabels) GetEntityData() *types.CommonEntityData {
    localLabels.EntityData.YFilter = localLabels.YFilter
    localLabels.EntityData.YangName = "local-labels"
    localLabels.EntityData.BundleName = "cisco_ios_xr"
    localLabels.EntityData.ParentYangName = "vrf"
    localLabels.EntityData.SegmentPath = "local-labels"
    localLabels.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/" + localLabels.EntityData.SegmentPath
    localLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabels.EntityData.Children = types.NewOrderedMap()
    localLabels.EntityData.Children.Append("local-label", types.YChild{"LocalLabel", nil})
    for i := range localLabels.LocalLabel {
        localLabels.EntityData.Children.Append(types.GetSegmentPath(localLabels.LocalLabel[i]), types.YChild{"LocalLabel", localLabels.LocalLabel[i]})
    }
    localLabels.EntityData.Leafs = types.NewOrderedMap()

    localLabels.EntityData.YListKeys = []string {}

    return &(localLabels.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel
// Data for static label
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    PathInfo []*MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo

    // Backup Path Information. The type is slice of
    // MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo.
    BackupPathInfo []*MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo
}

func (localLabel *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel) GetEntityData() *types.CommonEntityData {
    localLabel.EntityData.YFilter = localLabel.YFilter
    localLabel.EntityData.YangName = "local-label"
    localLabel.EntityData.BundleName = "cisco_ios_xr"
    localLabel.EntityData.ParentYangName = "local-labels"
    localLabel.EntityData.SegmentPath = "local-label" + types.AddKeyToken(localLabel.LocalLabelId, "local-label-id")
    localLabel.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/" + localLabel.EntityData.SegmentPath
    localLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabel.EntityData.Children = types.NewOrderedMap()
    localLabel.EntityData.Children.Append("prefix", types.YChild{"Prefix", &localLabel.Prefix})
    localLabel.EntityData.Children.Append("pathset-resolve-nh", types.YChild{"PathsetResolveNh", &localLabel.PathsetResolveNh})
    localLabel.EntityData.Children.Append("backup-pathset-resolve-nh", types.YChild{"BackupPathsetResolveNh", &localLabel.BackupPathsetResolveNh})
    localLabel.EntityData.Children.Append("path-info", types.YChild{"PathInfo", nil})
    for i := range localLabel.PathInfo {
        types.SetYListKey(localLabel.PathInfo[i], i)
        localLabel.EntityData.Children.Append(types.GetSegmentPath(localLabel.PathInfo[i]), types.YChild{"PathInfo", localLabel.PathInfo[i]})
    }
    localLabel.EntityData.Children.Append("backup-path-info", types.YChild{"BackupPathInfo", nil})
    for i := range localLabel.BackupPathInfo {
        types.SetYListKey(localLabel.BackupPathInfo[i], i)
        localLabel.EntityData.Children.Append(types.GetSegmentPath(localLabel.BackupPathInfo[i]), types.YChild{"BackupPathInfo", localLabel.BackupPathInfo[i]})
    }
    localLabel.EntityData.Leafs = types.NewOrderedMap()
    localLabel.EntityData.Leafs.Append("local-label-id", types.YLeaf{"LocalLabelId", localLabel.LocalLabelId})
    localLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", localLabel.Label})
    localLabel.EntityData.Leafs.Append("label-mode", types.YLeaf{"LabelMode", localLabel.LabelMode})
    localLabel.EntityData.Leafs.Append("label-status", types.YLeaf{"LabelStatus", localLabel.LabelStatus})
    localLabel.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", localLabel.VrfName})
    localLabel.EntityData.Leafs.Append("pathset-via-resolve", types.YLeaf{"PathsetViaResolve", localLabel.PathsetViaResolve})
    localLabel.EntityData.Leafs.Append("backup-pathset-via-resolve", types.YLeaf{"BackupPathsetViaResolve", localLabel.BackupPathsetViaResolve})
    localLabel.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", localLabel.AddressFamily})

    localLabel.EntityData.YListKeys = []string {"LocalLabelId"}

    return &(localLabel.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix
// Prefix Information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix
}

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "local-label"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefix.Prefix})
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefix.PrefixLength})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix
// Prefix
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_Prefix_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/prefix/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", prefix.AfName})
    prefix.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", prefix.Ipv4Prefix})
    prefix.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", prefix.Ipv6Prefix})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh
// Primary pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh struct {
    EntityData types.CommonEntityData
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

func (pathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathsetResolveNh) GetEntityData() *types.CommonEntityData {
    pathsetResolveNh.EntityData.YFilter = pathsetResolveNh.YFilter
    pathsetResolveNh.EntityData.YangName = "pathset-resolve-nh"
    pathsetResolveNh.EntityData.BundleName = "cisco_ios_xr"
    pathsetResolveNh.EntityData.ParentYangName = "local-label"
    pathsetResolveNh.EntityData.SegmentPath = "pathset-resolve-nh"
    pathsetResolveNh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/" + pathsetResolveNh.EntityData.SegmentPath
    pathsetResolveNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathsetResolveNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathsetResolveNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathsetResolveNh.EntityData.Children = types.NewOrderedMap()
    pathsetResolveNh.EntityData.Leafs = types.NewOrderedMap()
    pathsetResolveNh.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pathsetResolveNh.AfName})
    pathsetResolveNh.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", pathsetResolveNh.Ipv4Prefix})
    pathsetResolveNh.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", pathsetResolveNh.Ipv6Prefix})

    pathsetResolveNh.EntityData.YListKeys = []string {}

    return &(pathsetResolveNh.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh
// Backup pathset resolve-nexthop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh struct {
    EntityData types.CommonEntityData
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

func (backupPathsetResolveNh *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetEntityData() *types.CommonEntityData {
    backupPathsetResolveNh.EntityData.YFilter = backupPathsetResolveNh.YFilter
    backupPathsetResolveNh.EntityData.YangName = "backup-pathset-resolve-nh"
    backupPathsetResolveNh.EntityData.BundleName = "cisco_ios_xr"
    backupPathsetResolveNh.EntityData.ParentYangName = "local-label"
    backupPathsetResolveNh.EntityData.SegmentPath = "backup-pathset-resolve-nh"
    backupPathsetResolveNh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/" + backupPathsetResolveNh.EntityData.SegmentPath
    backupPathsetResolveNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPathsetResolveNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPathsetResolveNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPathsetResolveNh.EntityData.Children = types.NewOrderedMap()
    backupPathsetResolveNh.EntityData.Leafs = types.NewOrderedMap()
    backupPathsetResolveNh.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", backupPathsetResolveNh.AfName})
    backupPathsetResolveNh.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", backupPathsetResolveNh.Ipv4Prefix})
    backupPathsetResolveNh.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", backupPathsetResolveNh.Ipv6Prefix})

    backupPathsetResolveNh.EntityData.YListKeys = []string {}

    return &(backupPathsetResolveNh.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo
// Path Information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (pathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo) GetEntityData() *types.CommonEntityData {
    pathInfo.EntityData.YFilter = pathInfo.YFilter
    pathInfo.EntityData.YangName = "path-info"
    pathInfo.EntityData.BundleName = "cisco_ios_xr"
    pathInfo.EntityData.ParentYangName = "local-label"
    pathInfo.EntityData.SegmentPath = "path-info" + types.AddNoKeyToken(pathInfo)
    pathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/" + pathInfo.EntityData.SegmentPath
    pathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfo.EntityData.Children = types.NewOrderedMap()
    pathInfo.EntityData.Children.Append("nexthop", types.YChild{"Nexthop", &pathInfo.Nexthop})
    pathInfo.EntityData.Leafs = types.NewOrderedMap()
    pathInfo.EntityData.Leafs.Append("path-number", types.YLeaf{"PathNumber", pathInfo.PathNumber})
    pathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathInfo.Type})
    pathInfo.EntityData.Leafs.Append("path-role", types.YLeaf{"PathRole", pathInfo.PathRole})
    pathInfo.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", pathInfo.PathId})
    pathInfo.EntityData.Leafs.Append("backup-id", types.YLeaf{"BackupId", pathInfo.BackupId})
    pathInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", pathInfo.Status})

    pathInfo.EntityData.YListKeys = []string {}

    return &(pathInfo.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is MgmtStaticLspAfi.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "path-info"
    nexthop.EntityData.SegmentPath = "nexthop"
    nexthop.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/path-info/" + nexthop.EntityData.SegmentPath
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = types.NewOrderedMap()
    nexthop.EntityData.Children.Append("address", types.YChild{"Address", &nexthop.Address})
    nexthop.EntityData.Leafs = types.NewOrderedMap()
    nexthop.EntityData.Leafs.Append("label", types.YLeaf{"Label", nexthop.Label})
    nexthop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nexthop.InterfaceName})
    nexthop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nexthop.Afi})

    nexthop.EntityData.YListKeys = []string {}

    return &(nexthop.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address struct {
    EntityData types.CommonEntityData
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

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "nexthop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/path-info/nexthop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", address.AfName})
    address.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", address.Ipv4Prefix})
    address.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", address.Ipv6Prefix})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo
// Backup Path Information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (backupPathInfo *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo) GetEntityData() *types.CommonEntityData {
    backupPathInfo.EntityData.YFilter = backupPathInfo.YFilter
    backupPathInfo.EntityData.YangName = "backup-path-info"
    backupPathInfo.EntityData.BundleName = "cisco_ios_xr"
    backupPathInfo.EntityData.ParentYangName = "local-label"
    backupPathInfo.EntityData.SegmentPath = "backup-path-info" + types.AddNoKeyToken(backupPathInfo)
    backupPathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/" + backupPathInfo.EntityData.SegmentPath
    backupPathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPathInfo.EntityData.Children = types.NewOrderedMap()
    backupPathInfo.EntityData.Children.Append("nexthop", types.YChild{"Nexthop", &backupPathInfo.Nexthop})
    backupPathInfo.EntityData.Leafs = types.NewOrderedMap()
    backupPathInfo.EntityData.Leafs.Append("path-number", types.YLeaf{"PathNumber", backupPathInfo.PathNumber})
    backupPathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", backupPathInfo.Type})
    backupPathInfo.EntityData.Leafs.Append("path-role", types.YLeaf{"PathRole", backupPathInfo.PathRole})
    backupPathInfo.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", backupPathInfo.PathId})
    backupPathInfo.EntityData.Leafs.Append("backup-id", types.YLeaf{"BackupId", backupPathInfo.BackupId})
    backupPathInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", backupPathInfo.Status})

    backupPathInfo.EntityData.YListKeys = []string {}

    return &(backupPathInfo.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop
// Nexthop information
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is MgmtStaticLspAfi.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "backup-path-info"
    nexthop.EntityData.SegmentPath = "nexthop"
    nexthop.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/backup-path-info/" + nexthop.EntityData.SegmentPath
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = types.NewOrderedMap()
    nexthop.EntityData.Children.Append("address", types.YChild{"Address", &nexthop.Address})
    nexthop.EntityData.Leafs = types.NewOrderedMap()
    nexthop.EntityData.Leafs.Append("label", types.YLeaf{"Label", nexthop.Label})
    nexthop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nexthop.InterfaceName})
    nexthop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nexthop.Afi})

    nexthop.EntityData.YListKeys = []string {}

    return &(nexthop.EntityData)
}

// MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address struct {
    EntityData types.CommonEntityData
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

func (address *MplsStatic_Vrfs_Vrf_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "nexthop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/vrfs/vrf/local-labels/local-label/backup-path-info/nexthop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", address.AfName})
    address.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", address.Ipv4Prefix})
    address.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", address.Ipv6Prefix})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// MplsStatic_Summary
// MPLS STATIC summary data
type MplsStatic_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of LSPs. The type is interface{} with range: 0..4294967295.
    LspCount interface{}

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

    // Total Number of IPv4 ResolveNextHops. The type is interface{} with range:
    // 0..4294967295.
    Ipv4ResNhCount interface{}

    // Total Number of IPv6 ResoleNextHops. The type is interface{} with range:
    // 0..4294967295.
    Ipv6ResNhCount interface{}

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

func (summary *MplsStatic_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "mpls-static"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("lsp-count", types.YLeaf{"LspCount", summary.LspCount})
    summary.EntityData.Leafs.Append("label-count", types.YLeaf{"LabelCount", summary.LabelCount})
    summary.EntityData.Leafs.Append("label-error-count", types.YLeaf{"LabelErrorCount", summary.LabelErrorCount})
    summary.EntityData.Leafs.Append("label-discrepancy-count", types.YLeaf{"LabelDiscrepancyCount", summary.LabelDiscrepancyCount})
    summary.EntityData.Leafs.Append("vrf-count", types.YLeaf{"VrfCount", summary.VrfCount})
    summary.EntityData.Leafs.Append("active-vrf-count", types.YLeaf{"ActiveVrfCount", summary.ActiveVrfCount})
    summary.EntityData.Leafs.Append("interface-count", types.YLeaf{"InterfaceCount", summary.InterfaceCount})
    summary.EntityData.Leafs.Append("interface-foward-reference-count", types.YLeaf{"InterfaceFowardReferenceCount", summary.InterfaceFowardReferenceCount})
    summary.EntityData.Leafs.Append("mpls-enabled-interface-count", types.YLeaf{"MplsEnabledInterfaceCount", summary.MplsEnabledInterfaceCount})
    summary.EntityData.Leafs.Append("ipv4-res-nh-count", types.YLeaf{"Ipv4ResNhCount", summary.Ipv4ResNhCount})
    summary.EntityData.Leafs.Append("ipv6-res-nh-count", types.YLeaf{"Ipv6ResNhCount", summary.Ipv6ResNhCount})
    summary.EntityData.Leafs.Append("lsd-connected", types.YLeaf{"LsdConnected", summary.LsdConnected})
    summary.EntityData.Leafs.Append("im-connected", types.YLeaf{"ImConnected", summary.ImConnected})
    summary.EntityData.Leafs.Append("rsi-connected", types.YLeaf{"RsiConnected", summary.RsiConnected})
    summary.EntityData.Leafs.Append("ribv4-connected", types.YLeaf{"Ribv4Connected", summary.Ribv4Connected})
    summary.EntityData.Leafs.Append("ribv6-connected", types.YLeaf{"Ribv6Connected", summary.Ribv6Connected})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// MplsStatic_LocalLabels
// data for static local-label table
type MplsStatic_LocalLabels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data for static label. The type is slice of
    // MplsStatic_LocalLabels_LocalLabel.
    LocalLabel []*MplsStatic_LocalLabels_LocalLabel
}

func (localLabels *MplsStatic_LocalLabels) GetEntityData() *types.CommonEntityData {
    localLabels.EntityData.YFilter = localLabels.YFilter
    localLabels.EntityData.YangName = "local-labels"
    localLabels.EntityData.BundleName = "cisco_ios_xr"
    localLabels.EntityData.ParentYangName = "mpls-static"
    localLabels.EntityData.SegmentPath = "local-labels"
    localLabels.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/" + localLabels.EntityData.SegmentPath
    localLabels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabels.EntityData.Children = types.NewOrderedMap()
    localLabels.EntityData.Children.Append("local-label", types.YChild{"LocalLabel", nil})
    for i := range localLabels.LocalLabel {
        localLabels.EntityData.Children.Append(types.GetSegmentPath(localLabels.LocalLabel[i]), types.YChild{"LocalLabel", localLabels.LocalLabel[i]})
    }
    localLabels.EntityData.Leafs = types.NewOrderedMap()

    localLabels.EntityData.YListKeys = []string {}

    return &(localLabels.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel
// Data for static label
type MplsStatic_LocalLabels_LocalLabel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    PathInfo []*MplsStatic_LocalLabels_LocalLabel_PathInfo

    // Backup Path Information. The type is slice of
    // MplsStatic_LocalLabels_LocalLabel_BackupPathInfo.
    BackupPathInfo []*MplsStatic_LocalLabels_LocalLabel_BackupPathInfo
}

func (localLabel *MplsStatic_LocalLabels_LocalLabel) GetEntityData() *types.CommonEntityData {
    localLabel.EntityData.YFilter = localLabel.YFilter
    localLabel.EntityData.YangName = "local-label"
    localLabel.EntityData.BundleName = "cisco_ios_xr"
    localLabel.EntityData.ParentYangName = "local-labels"
    localLabel.EntityData.SegmentPath = "local-label" + types.AddKeyToken(localLabel.LocalLabelId, "local-label-id")
    localLabel.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/" + localLabel.EntityData.SegmentPath
    localLabel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLabel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLabel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLabel.EntityData.Children = types.NewOrderedMap()
    localLabel.EntityData.Children.Append("prefix", types.YChild{"Prefix", &localLabel.Prefix})
    localLabel.EntityData.Children.Append("pathset-resolve-nh", types.YChild{"PathsetResolveNh", &localLabel.PathsetResolveNh})
    localLabel.EntityData.Children.Append("backup-pathset-resolve-nh", types.YChild{"BackupPathsetResolveNh", &localLabel.BackupPathsetResolveNh})
    localLabel.EntityData.Children.Append("path-info", types.YChild{"PathInfo", nil})
    for i := range localLabel.PathInfo {
        types.SetYListKey(localLabel.PathInfo[i], i)
        localLabel.EntityData.Children.Append(types.GetSegmentPath(localLabel.PathInfo[i]), types.YChild{"PathInfo", localLabel.PathInfo[i]})
    }
    localLabel.EntityData.Children.Append("backup-path-info", types.YChild{"BackupPathInfo", nil})
    for i := range localLabel.BackupPathInfo {
        types.SetYListKey(localLabel.BackupPathInfo[i], i)
        localLabel.EntityData.Children.Append(types.GetSegmentPath(localLabel.BackupPathInfo[i]), types.YChild{"BackupPathInfo", localLabel.BackupPathInfo[i]})
    }
    localLabel.EntityData.Leafs = types.NewOrderedMap()
    localLabel.EntityData.Leafs.Append("local-label-id", types.YLeaf{"LocalLabelId", localLabel.LocalLabelId})
    localLabel.EntityData.Leafs.Append("label", types.YLeaf{"Label", localLabel.Label})
    localLabel.EntityData.Leafs.Append("label-mode", types.YLeaf{"LabelMode", localLabel.LabelMode})
    localLabel.EntityData.Leafs.Append("label-status", types.YLeaf{"LabelStatus", localLabel.LabelStatus})
    localLabel.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", localLabel.VrfName})
    localLabel.EntityData.Leafs.Append("pathset-via-resolve", types.YLeaf{"PathsetViaResolve", localLabel.PathsetViaResolve})
    localLabel.EntityData.Leafs.Append("backup-pathset-via-resolve", types.YLeaf{"BackupPathsetViaResolve", localLabel.BackupPathsetViaResolve})
    localLabel.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", localLabel.AddressFamily})

    localLabel.EntityData.YListKeys = []string {"LocalLabelId"}

    return &(localLabel.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_Prefix
// Prefix Information
type MplsStatic_LocalLabels_LocalLabel_Prefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // Prefix.
    Prefix MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix
}

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "local-label"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Children.Append("prefix", types.YChild{"Prefix", &prefix.Prefix})
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", prefix.PrefixLength})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix
// Prefix
type MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix struct {
    EntityData types.CommonEntityData
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

func (prefix *MplsStatic_LocalLabels_LocalLabel_Prefix_Prefix) GetEntityData() *types.CommonEntityData {
    prefix.EntityData.YFilter = prefix.YFilter
    prefix.EntityData.YangName = "prefix"
    prefix.EntityData.BundleName = "cisco_ios_xr"
    prefix.EntityData.ParentYangName = "prefix"
    prefix.EntityData.SegmentPath = "prefix"
    prefix.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/prefix/" + prefix.EntityData.SegmentPath
    prefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefix.EntityData.Children = types.NewOrderedMap()
    prefix.EntityData.Leafs = types.NewOrderedMap()
    prefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", prefix.AfName})
    prefix.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", prefix.Ipv4Prefix})
    prefix.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", prefix.Ipv6Prefix})

    prefix.EntityData.YListKeys = []string {}

    return &(prefix.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh
// Primary pathset resolve-nexthop IP Address
type MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh struct {
    EntityData types.CommonEntityData
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

func (pathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_PathsetResolveNh) GetEntityData() *types.CommonEntityData {
    pathsetResolveNh.EntityData.YFilter = pathsetResolveNh.YFilter
    pathsetResolveNh.EntityData.YangName = "pathset-resolve-nh"
    pathsetResolveNh.EntityData.BundleName = "cisco_ios_xr"
    pathsetResolveNh.EntityData.ParentYangName = "local-label"
    pathsetResolveNh.EntityData.SegmentPath = "pathset-resolve-nh"
    pathsetResolveNh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/" + pathsetResolveNh.EntityData.SegmentPath
    pathsetResolveNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathsetResolveNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathsetResolveNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathsetResolveNh.EntityData.Children = types.NewOrderedMap()
    pathsetResolveNh.EntityData.Leafs = types.NewOrderedMap()
    pathsetResolveNh.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pathsetResolveNh.AfName})
    pathsetResolveNh.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", pathsetResolveNh.Ipv4Prefix})
    pathsetResolveNh.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", pathsetResolveNh.Ipv6Prefix})

    pathsetResolveNh.EntityData.YListKeys = []string {}

    return &(pathsetResolveNh.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh
// Backup pathset resolve-nexthop IP Address
type MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh struct {
    EntityData types.CommonEntityData
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

func (backupPathsetResolveNh *MplsStatic_LocalLabels_LocalLabel_BackupPathsetResolveNh) GetEntityData() *types.CommonEntityData {
    backupPathsetResolveNh.EntityData.YFilter = backupPathsetResolveNh.YFilter
    backupPathsetResolveNh.EntityData.YangName = "backup-pathset-resolve-nh"
    backupPathsetResolveNh.EntityData.BundleName = "cisco_ios_xr"
    backupPathsetResolveNh.EntityData.ParentYangName = "local-label"
    backupPathsetResolveNh.EntityData.SegmentPath = "backup-pathset-resolve-nh"
    backupPathsetResolveNh.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/" + backupPathsetResolveNh.EntityData.SegmentPath
    backupPathsetResolveNh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPathsetResolveNh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPathsetResolveNh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPathsetResolveNh.EntityData.Children = types.NewOrderedMap()
    backupPathsetResolveNh.EntityData.Leafs = types.NewOrderedMap()
    backupPathsetResolveNh.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", backupPathsetResolveNh.AfName})
    backupPathsetResolveNh.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", backupPathsetResolveNh.Ipv4Prefix})
    backupPathsetResolveNh.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", backupPathsetResolveNh.Ipv6Prefix})

    backupPathsetResolveNh.EntityData.YListKeys = []string {}

    return &(backupPathsetResolveNh.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_PathInfo
// Path Information
type MplsStatic_LocalLabels_LocalLabel_PathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (pathInfo *MplsStatic_LocalLabels_LocalLabel_PathInfo) GetEntityData() *types.CommonEntityData {
    pathInfo.EntityData.YFilter = pathInfo.YFilter
    pathInfo.EntityData.YangName = "path-info"
    pathInfo.EntityData.BundleName = "cisco_ios_xr"
    pathInfo.EntityData.ParentYangName = "local-label"
    pathInfo.EntityData.SegmentPath = "path-info" + types.AddNoKeyToken(pathInfo)
    pathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/" + pathInfo.EntityData.SegmentPath
    pathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathInfo.EntityData.Children = types.NewOrderedMap()
    pathInfo.EntityData.Children.Append("nexthop", types.YChild{"Nexthop", &pathInfo.Nexthop})
    pathInfo.EntityData.Leafs = types.NewOrderedMap()
    pathInfo.EntityData.Leafs.Append("path-number", types.YLeaf{"PathNumber", pathInfo.PathNumber})
    pathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", pathInfo.Type})
    pathInfo.EntityData.Leafs.Append("path-role", types.YLeaf{"PathRole", pathInfo.PathRole})
    pathInfo.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", pathInfo.PathId})
    pathInfo.EntityData.Leafs.Append("backup-id", types.YLeaf{"BackupId", pathInfo.BackupId})
    pathInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", pathInfo.Status})

    pathInfo.EntityData.YListKeys = []string {}

    return &(pathInfo.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop
// Nexthop information
type MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is MgmtStaticLspAfi.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "path-info"
    nexthop.EntityData.SegmentPath = "nexthop"
    nexthop.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/path-info/" + nexthop.EntityData.SegmentPath
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = types.NewOrderedMap()
    nexthop.EntityData.Children.Append("address", types.YChild{"Address", &nexthop.Address})
    nexthop.EntityData.Leafs = types.NewOrderedMap()
    nexthop.EntityData.Leafs.Append("label", types.YLeaf{"Label", nexthop.Label})
    nexthop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nexthop.InterfaceName})
    nexthop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nexthop.Afi})

    nexthop.EntityData.YListKeys = []string {}

    return &(nexthop.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address struct {
    EntityData types.CommonEntityData
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

func (address *MplsStatic_LocalLabels_LocalLabel_PathInfo_Nexthop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "nexthop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/path-info/nexthop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", address.AfName})
    address.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", address.Ipv4Prefix})
    address.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", address.Ipv6Prefix})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_BackupPathInfo
// Backup Path Information
type MplsStatic_LocalLabels_LocalLabel_BackupPathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (backupPathInfo *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo) GetEntityData() *types.CommonEntityData {
    backupPathInfo.EntityData.YFilter = backupPathInfo.YFilter
    backupPathInfo.EntityData.YangName = "backup-path-info"
    backupPathInfo.EntityData.BundleName = "cisco_ios_xr"
    backupPathInfo.EntityData.ParentYangName = "local-label"
    backupPathInfo.EntityData.SegmentPath = "backup-path-info" + types.AddNoKeyToken(backupPathInfo)
    backupPathInfo.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/" + backupPathInfo.EntityData.SegmentPath
    backupPathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPathInfo.EntityData.Children = types.NewOrderedMap()
    backupPathInfo.EntityData.Children.Append("nexthop", types.YChild{"Nexthop", &backupPathInfo.Nexthop})
    backupPathInfo.EntityData.Leafs = types.NewOrderedMap()
    backupPathInfo.EntityData.Leafs.Append("path-number", types.YLeaf{"PathNumber", backupPathInfo.PathNumber})
    backupPathInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", backupPathInfo.Type})
    backupPathInfo.EntityData.Leafs.Append("path-role", types.YLeaf{"PathRole", backupPathInfo.PathRole})
    backupPathInfo.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", backupPathInfo.PathId})
    backupPathInfo.EntityData.Leafs.Append("backup-id", types.YLeaf{"BackupId", backupPathInfo.BackupId})
    backupPathInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", backupPathInfo.Status})

    backupPathInfo.EntityData.YListKeys = []string {}

    return &(backupPathInfo.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop
// Nexthop information
type MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Next-Hop Label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // Next-Hop Interface Name. The type is string.
    InterfaceName interface{}

    // Next-Hop AFI. The type is MgmtStaticLspAfi.
    Afi interface{}

    // Next-Hop IP Address.
    Address MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
}

func (nexthop *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop) GetEntityData() *types.CommonEntityData {
    nexthop.EntityData.YFilter = nexthop.YFilter
    nexthop.EntityData.YangName = "nexthop"
    nexthop.EntityData.BundleName = "cisco_ios_xr"
    nexthop.EntityData.ParentYangName = "backup-path-info"
    nexthop.EntityData.SegmentPath = "nexthop"
    nexthop.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/backup-path-info/" + nexthop.EntityData.SegmentPath
    nexthop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nexthop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nexthop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nexthop.EntityData.Children = types.NewOrderedMap()
    nexthop.EntityData.Children.Append("address", types.YChild{"Address", &nexthop.Address})
    nexthop.EntityData.Leafs = types.NewOrderedMap()
    nexthop.EntityData.Leafs.Append("label", types.YLeaf{"Label", nexthop.Label})
    nexthop.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nexthop.InterfaceName})
    nexthop.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", nexthop.Afi})

    nexthop.EntityData.YListKeys = []string {}

    return &(nexthop.EntityData)
}

// MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address
// Next-Hop IP Address
type MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address struct {
    EntityData types.CommonEntityData
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

func (address *MplsStatic_LocalLabels_LocalLabel_BackupPathInfo_Nexthop_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "nexthop"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-static-oper:mpls-static/local-labels/local-label/backup-path-info/nexthop/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", address.AfName})
    address.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", address.Ipv4Prefix})
    address.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", address.Ipv6Prefix})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

