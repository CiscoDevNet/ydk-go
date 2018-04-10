// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_te_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_te_datatypes"))
}

// MplsTesrlgExclude represents Mpls tesrlg exclude
type MplsTesrlgExclude string

const (
    // SRLG Mandatory Exclude
    MplsTesrlgExclude_mandatory MplsTesrlgExclude = "mandatory"

    // SRLG Preferred Exclude
    MplsTesrlgExclude_preferred MplsTesrlgExclude = "preferred"

    // SRLG Weighted Exclude
    MplsTesrlgExclude_weighted MplsTesrlgExclude = "weighted"
)

// MplsTeAffinityValue represents Mpls te affinity value
type MplsTeAffinityValue string

const (
    // Affinity value in Hex number
    MplsTeAffinityValue_hex_value MplsTeAffinityValue = "hex-value"

    // Affinity value by Bit-Position
    MplsTeAffinityValue_bit_position MplsTeAffinityValue = "bit-position"
)

// SrPrepend represents Sr prepend
type SrPrepend string

const (
    // NoneType
    SrPrepend_none_type SrPrepend = "none-type"

    // Next Label
    SrPrepend_next_label SrPrepend = "next-label"

    // BGP NHOP
    SrPrepend_bgp_n_hop SrPrepend = "bgp-n-hop"
)

// MplsTePathSelectionTiebreaker represents Mpls te path selection tiebreaker
type MplsTePathSelectionTiebreaker string

const (
    // Prefer the path with the least-utilized links
    MplsTePathSelectionTiebreaker_min_fill MplsTePathSelectionTiebreaker = "min-fill"

    // Prefer the path with the most-utilized links
    MplsTePathSelectionTiebreaker_max_fill MplsTePathSelectionTiebreaker = "max-fill"

    // Prefer a path with links utilized randomly
    MplsTePathSelectionTiebreaker_random MplsTePathSelectionTiebreaker = "random"
)

// MplsTeOtnApsProtection represents Mpls te otn aps protection
type MplsTeOtnApsProtection string

const (
    // 1PLUS1 UNIDIR NO APS
    MplsTeOtnApsProtection_Y_1plus1_unidir_no_aps MplsTeOtnApsProtection = "1plus1-unidir-no-aps"

    // 1PLUS1 UNIDIR APS
    MplsTeOtnApsProtection_Y_1plus1_unidir_aps MplsTeOtnApsProtection = "1plus1-unidir-aps"

    // 1PLUS1 BIDIR APS
    MplsTeOtnApsProtection_Y_1plus1_bdir_aps MplsTeOtnApsProtection = "1plus1-bdir-aps"
)

// MplsTePathOptionProperty represents Mpls te path option property
type MplsTePathOptionProperty string

const (
    // No property
    MplsTePathOptionProperty_none MplsTePathOptionProperty = "none"

    // Path is not a canditate forreoptimization
    MplsTePathOptionProperty_lockdown MplsTePathOptionProperty = "lockdown"

    // Explicit path does not require topology
    // database
    MplsTePathOptionProperty_verbatim MplsTePathOptionProperty = "verbatim"

    // Dynamic path found by PCE server
    MplsTePathOptionProperty_pce MplsTePathOptionProperty = "pce"

    // Segment Routing path
    MplsTePathOptionProperty_segment_routing MplsTePathOptionProperty = "segment-routing"
)

// MplsTeTunnelAffinity represents Mpls te tunnel affinity
type MplsTeTunnelAffinity string

const (
    // Include Affinity
    MplsTeTunnelAffinity_include MplsTeTunnelAffinity = "include"

    // Strictly Include Affinity
    MplsTeTunnelAffinity_include_strict MplsTeTunnelAffinity = "include-strict"

    // Exclude Affinity
    MplsTeTunnelAffinity_exclude MplsTeTunnelAffinity = "exclude"

    // Exclude All Affinities
    MplsTeTunnelAffinity_exclude_all MplsTeTunnelAffinity = "exclude-all"

    // Ignore Affinity
    MplsTeTunnelAffinity_ignore MplsTeTunnelAffinity = "ignore"
)

// MplsTeSwitchingCap represents Mpls te switching cap
type MplsTeSwitchingCap string

const (
    // PSC1
    MplsTeSwitchingCap_psc1 MplsTeSwitchingCap = "psc1"

    // LSC
    MplsTeSwitchingCap_lsc MplsTeSwitchingCap = "lsc"

    // FSC
    MplsTeSwitchingCap_fsc MplsTeSwitchingCap = "fsc"
)

// MplsTeOtnApsProtectionMode represents Mpls te otn aps protection mode
type MplsTeOtnApsProtectionMode string

const (
    // Revertive
    MplsTeOtnApsProtectionMode_revertive MplsTeOtnApsProtectionMode = "revertive"

    // Non Revertive
    MplsTeOtnApsProtectionMode_non_revertive MplsTeOtnApsProtectionMode = "non-revertive"
)

// MplsTeBfdSessionDownAction represents Mpls te bfd session down action
type MplsTeBfdSessionDownAction string

const (
    // Tear down and resetup
    MplsTeBfdSessionDownAction_re_setup MplsTeBfdSessionDownAction = "re-setup"
)

// MplsTeLogFrrProtection represents Mpls te log frr protection
type MplsTeLogFrrProtection string

const (
    // Track only FRR active on primary LSP
    MplsTeLogFrrProtection_frr_active_primary MplsTeLogFrrProtection = "frr-active-primary"

    // backup tunnel
    MplsTeLogFrrProtection_backup MplsTeLogFrrProtection = "backup"

    // Track only FRR ready on primary LSP
    MplsTeLogFrrProtection_frr_ready_primary MplsTeLogFrrProtection = "frr-ready-primary"

    // primary LSP
    MplsTeLogFrrProtection_primary MplsTeLogFrrProtection = "primary"

    // all
    MplsTeLogFrrProtection_all MplsTeLogFrrProtection = "all"
)

// MplsTeAutorouteMetric represents Mpls te autoroute metric
type MplsTeAutorouteMetric string

const (
    // Relative
    MplsTeAutorouteMetric_relative MplsTeAutorouteMetric = "relative"

    // Absolute
    MplsTeAutorouteMetric_absolute MplsTeAutorouteMetric = "absolute"

    // Constant
    MplsTeAutorouteMetric_constant MplsTeAutorouteMetric = "constant"
)

// MplsTeBackupBandwidthPool represents Mpls te backup bandwidth pool
type MplsTeBackupBandwidthPool string

const (
    // Any Pool
    MplsTeBackupBandwidthPool_any_pool MplsTeBackupBandwidthPool = "any-pool"

    // Global Pool
    MplsTeBackupBandwidthPool_global_pool MplsTeBackupBandwidthPool = "global-pool"

    // Sub Pool
    MplsTeBackupBandwidthPool_sub_pool MplsTeBackupBandwidthPool = "sub-pool"
)

// MplsTeOtnSncMode represents Mpls te otn snc mode
type MplsTeOtnSncMode string

const (
    // SNC N
    MplsTeOtnSncMode_snc_n MplsTeOtnSncMode = "snc-n"

    // SNC I
    MplsTeOtnSncMode_snc_i MplsTeOtnSncMode = "snc-i"

    // SNC S
    MplsTeOtnSncMode_snc_s MplsTeOtnSncMode = "snc-s"
)

// BfdReversePath represents Bfd reverse path
type BfdReversePath string

const (
    // BindingLabel
    BfdReversePath_bfd_reverse_path_binding_label BfdReversePath = "bfd-reverse-path-binding-label"
)

// MplsTePathSelectionMetric represents Mpls te path selection metric
type MplsTePathSelectionMetric string

const (
    // IGP Metric
    MplsTePathSelectionMetric_igp MplsTePathSelectionMetric = "igp"

    // TE Metric
    MplsTePathSelectionMetric_te MplsTePathSelectionMetric = "te"

    // DELAY Metric
    MplsTePathSelectionMetric_delay MplsTePathSelectionMetric = "delay"
)

// MplsTePathOption represents Mpls te path option
type MplsTePathOption string

const (
    // Not Set
    MplsTePathOption_not_set MplsTePathOption = "not-set"

    // Dynamic
    MplsTePathOption_dynamic MplsTePathOption = "dynamic"

    // Explicit, identified by name
    MplsTePathOption_explicit_name MplsTePathOption = "explicit-name"

    // Explicit, identified by number
    MplsTePathOption_explicit_number MplsTePathOption = "explicit-number"

    // No ERO
    MplsTePathOption_no_ero MplsTePathOption = "no-ero"

    // Segment routing
    MplsTePathOption_sr MplsTePathOption = "sr"
)

// MplsTeBandwidthPool represents Mpls te bandwidth pool
type MplsTeBandwidthPool string

const (
    // Any Pool
    MplsTeBandwidthPool_any_pool MplsTeBandwidthPool = "any-pool"

    // Sub Pool
    MplsTeBandwidthPool_sub_pool MplsTeBandwidthPool = "sub-pool"
)

// MplsTeBandwidthDste represents Mpls te bandwidth dste
type MplsTeBandwidthDste string

const (
    // IETF-Standard DSTE
    MplsTeBandwidthDste_standard_dste MplsTeBandwidthDste = "standard-dste"

    // Pre-Standard DSTE
    MplsTeBandwidthDste_pre_standard_dste MplsTeBandwidthDste = "pre-standard-dste"
)

// MplsTePathSelectionInvalidationTimerExpire represents Mpls te path selection invalidation timer expire
type MplsTePathSelectionInvalidationTimerExpire string

const (
    // Tear down tunnel.
    MplsTePathSelectionInvalidationTimerExpire_tunnel_action_tear MplsTePathSelectionInvalidationTimerExpire = "tunnel-action-tear"

    // Drop tunnel traffic.
    MplsTePathSelectionInvalidationTimerExpire_tunnel_action_drop MplsTePathSelectionInvalidationTimerExpire = "tunnel-action-drop"
)

// Ctype represents Ctype
type Ctype string

const (
    // CTYPE NULL
    Ctype_ctype_null Ctype = "ctype-null"

    // CTYPE IPV4
    Ctype_ctype_ipv4 Ctype = "ctype-ipv4"

    // CTYPE IPV4 P2P TUNNEL
    Ctype_ctype_ipv4_p2p_tunnel Ctype = "ctype-ipv4-p2p-tunnel"

    // CTYPE IPV6 P2P TUNNEL
    Ctype_ctype_ipv6_p2p_tunnel Ctype = "ctype-ipv6-p2p-tunnel"

    // CTYPE IPV4 UNI
    Ctype_ctype_ipv4_uni Ctype = "ctype-ipv4-uni"

    // CTYPE IPV4 P2MP TUNNEL
    Ctype_ctype_ipv4_p2mp_tunnel Ctype = "ctype-ipv4-p2mp-tunnel"

    // CTYPE IPV6 P2MP TUNNEL
    Ctype_ctype_ipv6_p2mp_tunnel Ctype = "ctype-ipv6-p2mp-tunnel"
)

// MplsTePathDiversityConformance represents Mpls te path diversity conformance
type MplsTePathDiversityConformance string

const (
    // Strict
    MplsTePathDiversityConformance_strict MplsTePathDiversityConformance = "strict"

    // Best effort
    MplsTePathDiversityConformance_best_effort MplsTePathDiversityConformance = "best-effort"
)

// MplsTeBackupBandwidthClass represents Mpls te backup bandwidth class
type MplsTeBackupBandwidthClass string

const (
    // Class 0
    MplsTeBackupBandwidthClass_class0 MplsTeBackupBandwidthClass = "class0"

    // Class 1
    MplsTeBackupBandwidthClass_class1 MplsTeBackupBandwidthClass = "class1"

    // Any Class
    MplsTeBackupBandwidthClass_any_class MplsTeBackupBandwidthClass = "any-class"
)

// MplsTeAttrSet represents Mpls te attr set
type MplsTeAttrSet string

const (
    // Not used
    MplsTeAttrSet_not_used MplsTeAttrSet = "not-used"

    // Static
    MplsTeAttrSet_static MplsTeAttrSet = "static"

    // LSP
    MplsTeAttrSet_lsp MplsTeAttrSet = "lsp"

    // Unassigned
    MplsTeAttrSet_unassigned MplsTeAttrSet = "unassigned"

    // Auto backup
    MplsTeAttrSet_auto_backup MplsTeAttrSet = "auto-backup"

    // Auto mesh
    MplsTeAttrSet_auto_mesh MplsTeAttrSet = "auto-mesh"

    // XRO
    MplsTeAttrSet_xro MplsTeAttrSet = "xro"

    // P2MP TE
    MplsTeAttrSet_p2mp_te MplsTeAttrSet = "p2mp-te"

    // OTN Path Protection
    MplsTeAttrSet_otn_pp MplsTeAttrSet = "otn-pp"

    // P2P TE
    MplsTeAttrSet_p2p_te MplsTeAttrSet = "p2p-te"
)

// MplsTeOtnApsRestorationStyle represents Mpls te otn aps restoration style
type MplsTeOtnApsRestorationStyle string

const (
    // Keep Failed Lsp
    MplsTeOtnApsRestorationStyle_keep_failed_lsp MplsTeOtnApsRestorationStyle = "keep-failed-lsp"

    // Delete Failed Lsp
    MplsTeOtnApsRestorationStyle_delete_failed_lsp MplsTeOtnApsRestorationStyle = "delete-failed-lsp"
)

// MplsTePathSelectionSegmentRoutingAdjacencyProtection represents protection
type MplsTePathSelectionSegmentRoutingAdjacencyProtection string

const (
    // Any segment can be used in a path.
    MplsTePathSelectionSegmentRoutingAdjacencyProtection_not_set MplsTePathSelectionSegmentRoutingAdjacencyProtection = "not-set"

    // Only unprotected adjacency segments can be used
    // in a path.
    MplsTePathSelectionSegmentRoutingAdjacencyProtection_adj_unprotected MplsTePathSelectionSegmentRoutingAdjacencyProtection = "adj-unprotected"

    // Only protected adjacency segments can be used
    // in a path.
    MplsTePathSelectionSegmentRoutingAdjacencyProtection_adj_protected MplsTePathSelectionSegmentRoutingAdjacencyProtection = "adj-protected"
)

// MplsTePathOptionProtection represents Mpls te path option protection
type MplsTePathOptionProtection string

const (
    // Active path
    MplsTePathOptionProtection_active MplsTePathOptionProtection = "active"

    // Protecting Path
    MplsTePathOptionProtection_protecting MplsTePathOptionProtection = "protecting"
)

// MplsTeSigNameOption represents Mpls te sig name option
type MplsTeSigNameOption string

const (
    // None
    MplsTeSigNameOption_none MplsTeSigNameOption = "none"

    // Address
    MplsTeSigNameOption_address MplsTeSigNameOption = "address"

    // Name
    MplsTeSigNameOption_name MplsTeSigNameOption = "name"
)

// MplsTeBandwidthLimit represents Mpls te bandwidth limit
type MplsTeBandwidthLimit string

const (
    // Unlimited
    MplsTeBandwidthLimit_unlimited MplsTeBandwidthLimit = "unlimited"

    // Limited
    MplsTeBandwidthLimit_limited MplsTeBandwidthLimit = "limited"
)

// PathInvalidationAction represents Path invalidation action
type PathInvalidationAction string

const (
    // Tear
    PathInvalidationAction_tear PathInvalidationAction = "tear"

    // Drop
    PathInvalidationAction_drop PathInvalidationAction = "drop"
)

// MplsTeIgpProtocol represents Mpls te igp protocol
type MplsTeIgpProtocol string

const (
    // Not set
    MplsTeIgpProtocol_none MplsTeIgpProtocol = "none"

    // IS IS
    MplsTeIgpProtocol_isis MplsTeIgpProtocol = "isis"

    // OSPF
    MplsTeIgpProtocol_ospf MplsTeIgpProtocol = "ospf"
)

