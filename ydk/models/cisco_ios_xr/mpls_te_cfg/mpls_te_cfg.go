// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-te package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-te: The root of MPLS TE configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg,
//   Cisco-IOS-XR-snmp-agent-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_te_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_te_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-te-cfg mpls-te}", reflect.TypeOf(MplsTe{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-te-cfg:mpls-te", reflect.TypeOf(MplsTe{}))
}

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

// MplsTePathComputationMethod represents Mpls te path computation method
type MplsTePathComputationMethod string

const (
    // NotSet
    MplsTePathComputationMethod_not_set MplsTePathComputationMethod = "not-set"

    // Dynamic
    MplsTePathComputationMethod_dynamic MplsTePathComputationMethod = "dynamic"

    // PCE
    MplsTePathComputationMethod_pce MplsTePathComputationMethod = "pce"

    // Explicit
    MplsTePathComputationMethod_explicit MplsTePathComputationMethod = "explicit"
)

// OtnDestination represents Otn destination
type OtnDestination string

const (
    // Destination numbered
    OtnDestination_number_ed OtnDestination = "number-ed"

    // Destination unnumbered
    OtnDestination_un_number_ed OtnDestination = "un-number-ed"
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

// MplsTeBfdSessionDownAction represents Mpls te bfd session down action
type MplsTeBfdSessionDownAction string

const (
    // Tear down and resetup
    MplsTeBfdSessionDownAction_re_setup MplsTeBfdSessionDownAction = "re-setup"
)

// RoutePriorityRole represents Route priority role
type RoutePriorityRole string

const (
    // TE Route Priority Role Head Backup
    RoutePriorityRole_route_priority_role_head_back_up RoutePriorityRole = "route-priority-role-head-back-up"

    // TE Route Priority Role Head Primary
    RoutePriorityRole_route_priority_role_head_primary RoutePriorityRole = "route-priority-role-head-primary"

    // TE Route Priority Role Middle
    RoutePriorityRole_route_priority_role_middle RoutePriorityRole = "route-priority-role-middle"
)

// MplsTebfdSession represents Mpls tebfd session
type MplsTebfdSession string

const (
    // Regular BFD
    MplsTebfdSession_regular_bfd MplsTebfdSession = "regular-bfd"

    // Seamless BFD
    MplsTebfdSession_sbfd MplsTebfdSession = "sbfd"

    // Redundant SBFD
    MplsTebfdSession_redundant_sbfd MplsTebfdSession = "redundant-sbfd"
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

// IetfMode represents Ietf mode
type IetfMode string

const (
    // IETF Standard
    IetfMode_standard IetfMode = "standard"
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

// GmplsttiMode represents Gmplstti mode
type GmplsttiMode string

const (
    // Section Monitoring
    GmplsttiMode_sm GmplsttiMode = "sm"

    // Path Monitoring
    GmplsttiMode_pm GmplsttiMode = "pm"

    // Tandem Connection
    GmplsttiMode_tcm GmplsttiMode = "tcm"
)

// MplsTePathSelectionInvalidationTimerExpire represents Mpls te path selection invalidation timer expire
type MplsTePathSelectionInvalidationTimerExpire string

const (
    // Tear down tunnel.
    MplsTePathSelectionInvalidationTimerExpire_tunnel_action_tear MplsTePathSelectionInvalidationTimerExpire = "tunnel-action-tear"

    // Drop tunnel traffic.
    MplsTePathSelectionInvalidationTimerExpire_tunnel_action_drop MplsTePathSelectionInvalidationTimerExpire = "tunnel-action-drop"
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

// MplsTeSwitchingEncoding represents Mpls te switching encoding
type MplsTeSwitchingEncoding string

const (
    // Packet
    MplsTeSwitchingEncoding_packet MplsTeSwitchingEncoding = "packet"

    // Ethernet
    MplsTeSwitchingEncoding_ethernet MplsTeSwitchingEncoding = "ethernet"

    // SONET SDH
    MplsTeSwitchingEncoding_sondet_sdh MplsTeSwitchingEncoding = "sondet-sdh"
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

// OtnProtectionSwitchLockout represents Otn protection switch lockout
type OtnProtectionSwitchLockout string

const (
    // No Lockout
    OtnProtectionSwitchLockout_none OtnProtectionSwitchLockout = "none"

    // Lockout Working
    OtnProtectionSwitchLockout_working OtnProtectionSwitchLockout = "working"
)

// MplsTeTunnelId represents Mpls te tunnel id
type MplsTeTunnelId string

const (
    // Auto
    MplsTeTunnelId_auto MplsTeTunnelId = "auto"

    // Explicit
    MplsTeTunnelId_explicit MplsTeTunnelId = "explicit"
)

// MplsTeAffinityValue represents Mpls te affinity value
type MplsTeAffinityValue string

const (
    // Affinity value in Hex number
    MplsTeAffinityValue_hex_value MplsTeAffinityValue = "hex-value"

    // Affinity value by Bit-Position
    MplsTeAffinityValue_bit_position MplsTeAffinityValue = "bit-position"
)

// OtnStaticUni represents Otn static uni
type OtnStaticUni string

const (
    // Uni-Type None
    OtnStaticUni_unknown OtnStaticUni = "unknown"

    // Uni-Type XC
    OtnStaticUni_xc OtnStaticUni = "xc"

    // Uni-Type Termination
    OtnStaticUni_termination OtnStaticUni = "termination"
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

// OtnPayload represents Otn payload
type OtnPayload string

const (
    // Payload unknown
    OtnPayload_unknown OtnPayload = "unknown"

    // Bmp Payload
    OtnPayload_bmp OtnPayload = "bmp"

    // Gfp_F Payload
    OtnPayload_gfp_f OtnPayload = "gfp-f"

    // GMP Payload
    OtnPayload_gmp OtnPayload = "gmp"

    // Gfp_F_EXT Payload
    OtnPayload_gfp_f_ext OtnPayload = "gfp-f-ext"
)

// OspfAreaMode represents Ospf area mode
type OspfAreaMode string

const (
    // OSPF area in integer format
    OspfAreaMode_ospf_int OspfAreaMode = "ospf-int"

    // OSPF area in IP address format
    OspfAreaMode_ospfip_addr OspfAreaMode = "ospfip-addr"
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

// PathInvalidationAction represents Path invalidation action
type PathInvalidationAction string

const (
    // Tear
    PathInvalidationAction_tear PathInvalidationAction = "tear"

    // Drop
    PathInvalidationAction_drop PathInvalidationAction = "drop"
)

// OtnSignaledBandwidth represents Otn signaled bandwidth
type OtnSignaledBandwidth string

const (
    // Signalled BW for ODU1
    OtnSignaledBandwidth_odu1 OtnSignaledBandwidth = "odu1"

    // Signalled BW for ODU2
    OtnSignaledBandwidth_odu2 OtnSignaledBandwidth = "odu2"

    // Signalled BW for ODU3
    OtnSignaledBandwidth_odu3 OtnSignaledBandwidth = "odu3"

    // Signalled BW for ODU4
    OtnSignaledBandwidth_odu4 OtnSignaledBandwidth = "odu4"

    // Signalled BW for ODU0
    OtnSignaledBandwidth_odu0 OtnSignaledBandwidth = "odu0"

    // Signalled BW for ODU2e
    OtnSignaledBandwidth_odu2e OtnSignaledBandwidth = "odu2e"

    // Signalled BW for ODUflex CBR
    OtnSignaledBandwidth_od_uflex_cbr OtnSignaledBandwidth = "od-uflex-cbr"

    // Signalled BW for ODUflex GFP Resizable
    OtnSignaledBandwidth_od_uflex_gfp_resize OtnSignaledBandwidth = "od-uflex-gfp-resize"

    // Signalled BW for ODUflex GFP not Resizable
    OtnSignaledBandwidth_od_uflex_gfp_not_resize OtnSignaledBandwidth = "od-uflex-gfp-not-resize"

    // Signalled BW for ODU1e
    OtnSignaledBandwidth_odu1e OtnSignaledBandwidth = "odu1e"

    // Signalled BW for ODU1f
    OtnSignaledBandwidth_odu1f OtnSignaledBandwidth = "odu1f"

    // Signalled BW for ODU2f
    OtnSignaledBandwidth_odu2f OtnSignaledBandwidth = "odu2f"

    // Signalled BW for ODU3e1
    OtnSignaledBandwidth_odu3e1 OtnSignaledBandwidth = "odu3e1"

    // Signalled BW for ODU3e2
    OtnSignaledBandwidth_odu3e2 OtnSignaledBandwidth = "odu3e2"
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

// BindingSegmentId represents Binding segment id
type BindingSegmentId string

const (
    // AnyLabel
    BindingSegmentId_any_label BindingSegmentId = "any-label"

    // SpecifiedLabel
    BindingSegmentId_specified_label BindingSegmentId = "specified-label"
)

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

// MplsTeSignaledLabel represents Mpls te signaled label
type MplsTeSignaledLabel string

const (
    // Not Set
    MplsTeSignaledLabel_not_set MplsTeSignaledLabel = "not-set"

    // DWDM Label (RFC 6205), 50GHz channel spacing
    MplsTeSignaledLabel_dwdm MplsTeSignaledLabel = "dwdm"
)

// BandwidthConstraint represents Bandwidth constraint
type BandwidthConstraint string

const (
    // Maximum Allocation Bandwidth Constaints Model
    BandwidthConstraint_bandwidth_constraint_maximum_allocation_model BandwidthConstraint = "bandwidth-constraint-maximum-allocation-model"
)

// MplsTeSwitchingEncode represents Mpls te switching encode
type MplsTeSwitchingEncode string

const (
    // None
    MplsTeSwitchingEncode_none MplsTeSwitchingEncode = "none"

    // Packet
    MplsTeSwitchingEncode_packet MplsTeSwitchingEncode = "packet"

    // Ethernet
    MplsTeSwitchingEncode_ethernet MplsTeSwitchingEncode = "ethernet"

    // SONET SDH
    MplsTeSwitchingEncode_sondet_sdh MplsTeSwitchingEncode = "sondet-sdh"
)

// MplsTePathSelectionMetric represents Mpls te path selection metric
type MplsTePathSelectionMetric string

const (
    // IGP Metric
    MplsTePathSelectionMetric_igp MplsTePathSelectionMetric = "igp"

    // TE Metric
    MplsTePathSelectionMetric_te MplsTePathSelectionMetric = "te"
)

// MplsTePathOptionProtection represents Mpls te path option protection
type MplsTePathOptionProtection string

const (
    // Active path
    MplsTePathOptionProtection_active MplsTePathOptionProtection = "active"

    // Protecting Path
    MplsTePathOptionProtection_protecting MplsTePathOptionProtection = "protecting"
)

// OtnSignaledBandwidthFlexFraming represents Otn signaled bandwidth flex framing
type OtnSignaledBandwidthFlexFraming string

const (
    // CBR
    OtnSignaledBandwidthFlexFraming_cbr OtnSignaledBandwidthFlexFraming = "cbr"

    // GFP fixed framing type
    OtnSignaledBandwidthFlexFraming_framed_gfp_fixed OtnSignaledBandwidthFlexFraming = "framed-gfp-fixed"

    // GFP resizeable framing type
    OtnSignaledBandwidthFlexFraming_framed_gfp_resize OtnSignaledBandwidthFlexFraming = "framed-gfp-resize"
)

// MplsTeBandwidthLimit represents Mpls te bandwidth limit
type MplsTeBandwidthLimit string

const (
    // Unlimited
    MplsTeBandwidthLimit_unlimited MplsTeBandwidthLimit = "unlimited"

    // Limited
    MplsTeBandwidthLimit_limited MplsTeBandwidthLimit = "limited"
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

// LinkNextHop represents Link next hop
type LinkNextHop string

const (
    // No next hop
    LinkNextHop_none LinkNextHop = "none"

    // IPv4 next-hop address
    LinkNextHop_ipv4_address LinkNextHop = "ipv4-address"
)

// MplsLcacFloodingIgp represents Mpls lcac flooding igp
type MplsLcacFloodingIgp string

const (
    // OSPF
    MplsLcacFloodingIgp_ospf MplsLcacFloodingIgp = "ospf"
)

// BfdReversePath represents Bfd reverse path
type BfdReversePath string

const (
    // BindingLabel
    BfdReversePath_bfd_reverse_path_binding_label BfdReversePath = "bfd-reverse-path-binding-label"
)

// MplsTeOtnApsRestorationStyle represents Mpls te otn aps restoration style
type MplsTeOtnApsRestorationStyle string

const (
    // Keep Failed Lsp
    MplsTeOtnApsRestorationStyle_keep_failed_lsp MplsTeOtnApsRestorationStyle = "keep-failed-lsp"

    // Delete Failed Lsp
    MplsTeOtnApsRestorationStyle_delete_failed_lsp MplsTeOtnApsRestorationStyle = "delete-failed-lsp"
)

// MplsTeBandwidthDste represents Mpls te bandwidth dste
type MplsTeBandwidthDste string

const (
    // IETF-Standard DSTE
    MplsTeBandwidthDste_standard_dste MplsTeBandwidthDste = "standard-dste"

    // Pre-Standard DSTE
    MplsTeBandwidthDste_pre_standard_dste MplsTeBandwidthDste = "pre-standard-dste"
)

// MplsTePathDiversityConformance represents Mpls te path diversity conformance
type MplsTePathDiversityConformance string

const (
    // Strict
    MplsTePathDiversityConformance_strict MplsTePathDiversityConformance = "strict"

    // Best effort
    MplsTePathDiversityConformance_best_effort MplsTePathDiversityConformance = "best-effort"
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

// MplsTeOtnApsProtectionMode represents Mpls te otn aps protection mode
type MplsTeOtnApsProtectionMode string

const (
    // Revertive
    MplsTeOtnApsProtectionMode_revertive MplsTeOtnApsProtectionMode = "revertive"

    // Non Revertive
    MplsTeOtnApsProtectionMode_non_revertive MplsTeOtnApsProtectionMode = "non-revertive"
)

// MplsTeSwitchingIndex represents Mpls te switching index
type MplsTeSwitchingIndex string

const (
    // Link
    MplsTeSwitchingIndex_link MplsTeSwitchingIndex = "link"
)

// MplsTeConfigTunnel represents Mpls te config tunnel
type MplsTeConfigTunnel string

const (
    // P2P
    MplsTeConfigTunnel_p2p MplsTeConfigTunnel = "p2p"

    // P2MP
    MplsTeConfigTunnel_p2mp MplsTeConfigTunnel = "p2mp"
)

// MplsTe
// The root of MPLS TE configuration
type MplsTe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable MPLS Traffic Engineering. The type is interface{}.
    EnableTrafficEngineering interface{}

    // Configure Diff-Serv Traffic-Engineering.
    DiffServTrafficEngineering MplsTe_DiffServTrafficEngineering

    // Configure MPLS TE tunnel.
    NamedTunnels MplsTe_NamedTunnels

    // GMPLS-UNI configuration.
    GmplsUni MplsTe_GmplsUni

    // Configure MPLS TE global attributes.
    GlobalAttributes MplsTe_GlobalAttributes

    // MPLS transport profile configuration data.
    TransportProfile MplsTe_TransportProfile

    // Configure MPLS TE interfaces.
    Interfaces MplsTe_Interfaces

    // GMPLS-NNI configuration.
    GmplsNni MplsTe_GmplsNni

    // LCAC specific MPLS global configuration.
    Lcac MplsTe_Lcac
}

func (mplsTe *MplsTe) GetEntityData() *types.CommonEntityData {
    mplsTe.EntityData.YFilter = mplsTe.YFilter
    mplsTe.EntityData.YangName = "mpls-te"
    mplsTe.EntityData.BundleName = "cisco_ios_xr"
    mplsTe.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-te-cfg"
    mplsTe.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te"
    mplsTe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTe.EntityData.Children = types.NewOrderedMap()
    mplsTe.EntityData.Children.Append("diff-serv-traffic-engineering", types.YChild{"DiffServTrafficEngineering", &mplsTe.DiffServTrafficEngineering})
    mplsTe.EntityData.Children.Append("named-tunnels", types.YChild{"NamedTunnels", &mplsTe.NamedTunnels})
    mplsTe.EntityData.Children.Append("gmpls-uni", types.YChild{"GmplsUni", &mplsTe.GmplsUni})
    mplsTe.EntityData.Children.Append("global-attributes", types.YChild{"GlobalAttributes", &mplsTe.GlobalAttributes})
    mplsTe.EntityData.Children.Append("transport-profile", types.YChild{"TransportProfile", &mplsTe.TransportProfile})
    mplsTe.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &mplsTe.Interfaces})
    mplsTe.EntityData.Children.Append("gmpls-nni", types.YChild{"GmplsNni", &mplsTe.GmplsNni})
    mplsTe.EntityData.Children.Append("lcac", types.YChild{"Lcac", &mplsTe.Lcac})
    mplsTe.EntityData.Leafs = types.NewOrderedMap()
    mplsTe.EntityData.Leafs.Append("enable-traffic-engineering", types.YLeaf{"EnableTrafficEngineering", mplsTe.EnableTrafficEngineering})

    mplsTe.EntityData.YListKeys = []string {}

    return &(mplsTe.EntityData)
}

// MplsTe_DiffServTrafficEngineering
// Configure Diff-Serv Traffic-Engineering
type MplsTe_DiffServTrafficEngineering struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Diff-Serv Traffic-Engineering Bandwidth Constraint Model. The type is
    // BandwidthConstraint.
    BandwidthConstraintModel interface{}

    // Diff-Serv Traffic-Engineering IETF mode. The type is IetfMode.
    ModeIetf interface{}

    // Configure Diff-Serv Traffic-Engineering Classes.
    Classes MplsTe_DiffServTrafficEngineering_Classes
}

func (diffServTrafficEngineering *MplsTe_DiffServTrafficEngineering) GetEntityData() *types.CommonEntityData {
    diffServTrafficEngineering.EntityData.YFilter = diffServTrafficEngineering.YFilter
    diffServTrafficEngineering.EntityData.YangName = "diff-serv-traffic-engineering"
    diffServTrafficEngineering.EntityData.BundleName = "cisco_ios_xr"
    diffServTrafficEngineering.EntityData.ParentYangName = "mpls-te"
    diffServTrafficEngineering.EntityData.SegmentPath = "diff-serv-traffic-engineering"
    diffServTrafficEngineering.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diffServTrafficEngineering.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diffServTrafficEngineering.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diffServTrafficEngineering.EntityData.Children = types.NewOrderedMap()
    diffServTrafficEngineering.EntityData.Children.Append("classes", types.YChild{"Classes", &diffServTrafficEngineering.Classes})
    diffServTrafficEngineering.EntityData.Leafs = types.NewOrderedMap()
    diffServTrafficEngineering.EntityData.Leafs.Append("bandwidth-constraint-model", types.YLeaf{"BandwidthConstraintModel", diffServTrafficEngineering.BandwidthConstraintModel})
    diffServTrafficEngineering.EntityData.Leafs.Append("mode-ietf", types.YLeaf{"ModeIetf", diffServTrafficEngineering.ModeIetf})

    diffServTrafficEngineering.EntityData.YListKeys = []string {}

    return &(diffServTrafficEngineering.EntityData)
}

// MplsTe_DiffServTrafficEngineering_Classes
// Configure Diff-Serv Traffic-Engineering Classes
type MplsTe_DiffServTrafficEngineering_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DSTE class number. The type is slice of
    // MplsTe_DiffServTrafficEngineering_Classes_Class.
    Class []*MplsTe_DiffServTrafficEngineering_Classes_Class
}

func (classes *MplsTe_DiffServTrafficEngineering_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "diff-serv-traffic-engineering"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

    return &(classes.EntityData)
}

// MplsTe_DiffServTrafficEngineering_Classes_Class
// DSTE class number
type MplsTe_DiffServTrafficEngineering_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. DS-TE class number. The type is interface{} with
    // range: 0..7.
    ClassNumber interface{}

    // Class type number. The type is interface{} with range: 0..1. The default
    // value is 1.
    ClassType interface{}

    // Class-type priority. The type is interface{} with range: 0..7. The default
    // value is 1.
    ClassPriority interface{}

    // TRUE to skip classtype and class priority provisioning FALSE to provision
    // them. The type is bool.
    Unused interface{}
}

func (class *MplsTe_DiffServTrafficEngineering_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassNumber, "class-number")
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-number", types.YLeaf{"ClassNumber", class.ClassNumber})
    class.EntityData.Leafs.Append("class-type", types.YLeaf{"ClassType", class.ClassType})
    class.EntityData.Leafs.Append("class-priority", types.YLeaf{"ClassPriority", class.ClassPriority})
    class.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", class.Unused})

    class.EntityData.YListKeys = []string {"ClassNumber"}

    return &(class.EntityData)
}

// MplsTe_NamedTunnels
// Configure MPLS TE tunnel
type MplsTe_NamedTunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Named Tunnels. The type is interface{}.
    Enable interface{}

    // Configure MPLS TE tunnel.
    Tunnels MplsTe_NamedTunnels_Tunnels
}

func (namedTunnels *MplsTe_NamedTunnels) GetEntityData() *types.CommonEntityData {
    namedTunnels.EntityData.YFilter = namedTunnels.YFilter
    namedTunnels.EntityData.YangName = "named-tunnels"
    namedTunnels.EntityData.BundleName = "cisco_ios_xr"
    namedTunnels.EntityData.ParentYangName = "mpls-te"
    namedTunnels.EntityData.SegmentPath = "named-tunnels"
    namedTunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    namedTunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    namedTunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    namedTunnels.EntityData.Children = types.NewOrderedMap()
    namedTunnels.EntityData.Children.Append("tunnels", types.YChild{"Tunnels", &namedTunnels.Tunnels})
    namedTunnels.EntityData.Leafs = types.NewOrderedMap()
    namedTunnels.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", namedTunnels.Enable})

    namedTunnels.EntityData.YListKeys = []string {}

    return &(namedTunnels.EntityData)
}

// MplsTe_NamedTunnels_Tunnels
// Configure MPLS TE tunnel
type MplsTe_NamedTunnels_Tunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a MPLS TE tunnel. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel.
    Tunnel []*MplsTe_NamedTunnels_Tunnels_Tunnel
}

func (tunnels *MplsTe_NamedTunnels_Tunnels) GetEntityData() *types.CommonEntityData {
    tunnels.EntityData.YFilter = tunnels.YFilter
    tunnels.EntityData.YangName = "tunnels"
    tunnels.EntityData.BundleName = "cisco_ios_xr"
    tunnels.EntityData.ParentYangName = "named-tunnels"
    tunnels.EntityData.SegmentPath = "tunnels"
    tunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnels.EntityData.Children = types.NewOrderedMap()
    tunnels.EntityData.Children.Append("tunnel", types.YChild{"Tunnel", nil})
    for i := range tunnels.Tunnel {
        tunnels.EntityData.Children.Append(types.GetSegmentPath(tunnels.Tunnel[i]), types.YChild{"Tunnel", tunnels.Tunnel[i]})
    }
    tunnels.EntityData.Leafs = types.NewOrderedMap()

    tunnels.EntityData.YListKeys = []string {}

    return &(tunnels.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel
// Configure a MPLS TE tunnel
type MplsTe_NamedTunnels_Tunnels_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Tunnel name. The type is string with length:
    // 1..59.
    TunnelName interface{}

    // This attribute is a key. Tunnel Type. The type is MplsTeConfigTunnel.
    TunnelType interface{}

    // Always set to true. The type is interface{}.
    Enable interface{}

    // MPLS tunnel attributes.
    TunnelAttributes MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes

    // Set the tunnel ID.
    TunnelId MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelId
}

func (tunnel *MplsTe_NamedTunnels_Tunnels_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "tunnels"
    tunnel.EntityData.SegmentPath = "tunnel" + types.AddKeyToken(tunnel.TunnelName, "tunnel-name") + types.AddKeyToken(tunnel.TunnelType, "tunnel-type")
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = types.NewOrderedMap()
    tunnel.EntityData.Children.Append("tunnel-attributes", types.YChild{"TunnelAttributes", &tunnel.TunnelAttributes})
    tunnel.EntityData.Children.Append("tunnel-id", types.YChild{"TunnelId", &tunnel.TunnelId})
    tunnel.EntityData.Leafs = types.NewOrderedMap()
    tunnel.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", tunnel.TunnelName})
    tunnel.EntityData.Leafs.Append("tunnel-type", types.YLeaf{"TunnelType", tunnel.TunnelType})
    tunnel.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", tunnel.Enable})

    tunnel.EntityData.YListKeys = []string {"TunnelName", "TunnelType"}

    return &(tunnel.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes
// MPLS tunnel attributes
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // shutdown the tunnel. The type is interface{}.
    Shutdown interface{}

    // Forward class value. The type is interface{} with range: 0..7.
    ForwardClass interface{}

    // Set the destination of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Destination interface{}

    // Record the route used by the tunnel. The type is interface{}.
    RecordRoute interface{}

    // Path selection metric to use in path calculation. The type is
    // MplsTePathSelectionMetric.
    PathSelectionMetric interface{}

    // Enable the soft-preemption feature on the tunnel. The type is interface{}.
    SoftPreemption interface{}

    // Tunnel loadsharing metric. The type is interface{} with range:
    // 1..4294967295.
    LoadShare interface{}

    // Tunnel path setup table.
    PathSetups MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups

    // Configure path selection properties.
    TunnelPathSelection MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection

    // Tunnel Interface Auto-bandwidth configuration data.
    AutoBandwidth MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth

    // Tunnel Setup and Hold Priorities.
    Priority MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Priority

    // Set the affinity flags and mask.
    AffinityMask MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AffinityMask

    // Log tunnel LSP messages.
    Logging MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Logging

    // Tunnel bandwidth requirement.
    Bandwidth MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Bandwidth

    // Parameters for IGP routing over tunnel.
    Autoroute MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute

    // Tunnel new style affinity attributes table.
    NewStyleAffinityAffinityTypes MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes

    // Specify MPLS tunnel can be fast-rerouted.
    FastReroute MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_FastReroute
}

func (tunnelAttributes *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes) GetEntityData() *types.CommonEntityData {
    tunnelAttributes.EntityData.YFilter = tunnelAttributes.YFilter
    tunnelAttributes.EntityData.YangName = "tunnel-attributes"
    tunnelAttributes.EntityData.BundleName = "cisco_ios_xr"
    tunnelAttributes.EntityData.ParentYangName = "tunnel"
    tunnelAttributes.EntityData.SegmentPath = "tunnel-attributes"
    tunnelAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelAttributes.EntityData.Children = types.NewOrderedMap()
    tunnelAttributes.EntityData.Children.Append("path-setups", types.YChild{"PathSetups", &tunnelAttributes.PathSetups})
    tunnelAttributes.EntityData.Children.Append("tunnel-path-selection", types.YChild{"TunnelPathSelection", &tunnelAttributes.TunnelPathSelection})
    tunnelAttributes.EntityData.Children.Append("auto-bandwidth", types.YChild{"AutoBandwidth", &tunnelAttributes.AutoBandwidth})
    tunnelAttributes.EntityData.Children.Append("priority", types.YChild{"Priority", &tunnelAttributes.Priority})
    tunnelAttributes.EntityData.Children.Append("affinity-mask", types.YChild{"AffinityMask", &tunnelAttributes.AffinityMask})
    tunnelAttributes.EntityData.Children.Append("logging", types.YChild{"Logging", &tunnelAttributes.Logging})
    tunnelAttributes.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &tunnelAttributes.Bandwidth})
    tunnelAttributes.EntityData.Children.Append("autoroute", types.YChild{"Autoroute", &tunnelAttributes.Autoroute})
    tunnelAttributes.EntityData.Children.Append("new-style-affinity-affinity-types", types.YChild{"NewStyleAffinityAffinityTypes", &tunnelAttributes.NewStyleAffinityAffinityTypes})
    tunnelAttributes.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &tunnelAttributes.FastReroute})
    tunnelAttributes.EntityData.Leafs = types.NewOrderedMap()
    tunnelAttributes.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", tunnelAttributes.Shutdown})
    tunnelAttributes.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", tunnelAttributes.ForwardClass})
    tunnelAttributes.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", tunnelAttributes.Destination})
    tunnelAttributes.EntityData.Leafs.Append("record-route", types.YLeaf{"RecordRoute", tunnelAttributes.RecordRoute})
    tunnelAttributes.EntityData.Leafs.Append("path-selection-metric", types.YLeaf{"PathSelectionMetric", tunnelAttributes.PathSelectionMetric})
    tunnelAttributes.EntityData.Leafs.Append("soft-preemption", types.YLeaf{"SoftPreemption", tunnelAttributes.SoftPreemption})
    tunnelAttributes.EntityData.Leafs.Append("load-share", types.YLeaf{"LoadShare", tunnelAttributes.LoadShare})

    tunnelAttributes.EntityData.YListKeys = []string {}

    return &(tunnelAttributes.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups
// Tunnel path setup table
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel path setup. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup.
    PathSetup []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup
}

func (pathSetups *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups) GetEntityData() *types.CommonEntityData {
    pathSetups.EntityData.YFilter = pathSetups.YFilter
    pathSetups.EntityData.YangName = "path-setups"
    pathSetups.EntityData.BundleName = "cisco_ios_xr"
    pathSetups.EntityData.ParentYangName = "tunnel-attributes"
    pathSetups.EntityData.SegmentPath = "path-setups"
    pathSetups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSetups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSetups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSetups.EntityData.Children = types.NewOrderedMap()
    pathSetups.EntityData.Children.Append("path-setup", types.YChild{"PathSetup", nil})
    for i := range pathSetups.PathSetup {
        pathSetups.EntityData.Children.Append(types.GetSegmentPath(pathSetups.PathSetup[i]), types.YChild{"PathSetup", pathSetups.PathSetup[i]})
    }
    pathSetups.EntityData.Leafs = types.NewOrderedMap()

    pathSetups.EntityData.YListKeys = []string {}

    return &(pathSetups.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup
// Tunnel path setup
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    PathSetupName interface{}

    // Path preference level. The type is interface{} with range: 0..4294967295.
    Preference interface{}

    // Always set to true. The type is interface{}.
    Enable interface{}

    // Path computation method.
    PathComputation MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup_PathComputation
}

func (pathSetup *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup) GetEntityData() *types.CommonEntityData {
    pathSetup.EntityData.YFilter = pathSetup.YFilter
    pathSetup.EntityData.YangName = "path-setup"
    pathSetup.EntityData.BundleName = "cisco_ios_xr"
    pathSetup.EntityData.ParentYangName = "path-setups"
    pathSetup.EntityData.SegmentPath = "path-setup" + types.AddKeyToken(pathSetup.PathSetupName, "path-setup-name")
    pathSetup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSetup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSetup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSetup.EntityData.Children = types.NewOrderedMap()
    pathSetup.EntityData.Children.Append("path-computation", types.YChild{"PathComputation", &pathSetup.PathComputation})
    pathSetup.EntityData.Leafs = types.NewOrderedMap()
    pathSetup.EntityData.Leafs.Append("path-setup-name", types.YLeaf{"PathSetupName", pathSetup.PathSetupName})
    pathSetup.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", pathSetup.Preference})
    pathSetup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSetup.Enable})

    pathSetup.EntityData.YListKeys = []string {"PathSetupName"}

    return &(pathSetup.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup_PathComputation
// Path computation method
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup_PathComputation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Path computation method. The type is MplsTePathComputationMethod. This
    // attribute is mandatory.
    PathComputationMethod interface{}

    // Explicit Path Name. The type is string.
    ExplicitPathName interface{}

    // Path Computation Server Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // The default value is 0.0.0.0.
    PathComputationServer interface{}
}

func (pathComputation *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_PathSetups_PathSetup_PathComputation) GetEntityData() *types.CommonEntityData {
    pathComputation.EntityData.YFilter = pathComputation.YFilter
    pathComputation.EntityData.YangName = "path-computation"
    pathComputation.EntityData.BundleName = "cisco_ios_xr"
    pathComputation.EntityData.ParentYangName = "path-setup"
    pathComputation.EntityData.SegmentPath = "path-computation"
    pathComputation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathComputation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathComputation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathComputation.EntityData.Children = types.NewOrderedMap()
    pathComputation.EntityData.Leafs = types.NewOrderedMap()
    pathComputation.EntityData.Leafs.Append("path-computation-method", types.YLeaf{"PathComputationMethod", pathComputation.PathComputationMethod})
    pathComputation.EntityData.Leafs.Append("explicit-path-name", types.YLeaf{"ExplicitPathName", pathComputation.ExplicitPathName})
    pathComputation.EntityData.Leafs.Append("path-computation-server", types.YLeaf{"PathComputationServer", pathComputation.PathComputationServer})

    pathComputation.EntityData.YListKeys = []string {}

    return &(pathComputation.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection
// Configure path selection properties
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // CSPF tiebreaker to use in path calculation. The type is
    // MplsTePathSelectionTiebreaker.
    Tiebreaker interface{}

    // Path selection hop limit configuration for this specific tunnel. The type
    // is interface{} with range: 1..255.
    PathSelectionHopLimit interface{}

    // Path selection cost limit configuration for this specific tunnel. The type
    // is interface{} with range: 1..4294967295.
    PathSelectionCostLimit interface{}

    // Path invalidation configuration for this specific tunnel.
    Invalidation MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection_Invalidation
}

func (tunnelPathSelection *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection) GetEntityData() *types.CommonEntityData {
    tunnelPathSelection.EntityData.YFilter = tunnelPathSelection.YFilter
    tunnelPathSelection.EntityData.YangName = "tunnel-path-selection"
    tunnelPathSelection.EntityData.BundleName = "cisco_ios_xr"
    tunnelPathSelection.EntityData.ParentYangName = "tunnel-attributes"
    tunnelPathSelection.EntityData.SegmentPath = "tunnel-path-selection"
    tunnelPathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelPathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelPathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelPathSelection.EntityData.Children = types.NewOrderedMap()
    tunnelPathSelection.EntityData.Children.Append("invalidation", types.YChild{"Invalidation", &tunnelPathSelection.Invalidation})
    tunnelPathSelection.EntityData.Leafs = types.NewOrderedMap()
    tunnelPathSelection.EntityData.Leafs.Append("tiebreaker", types.YLeaf{"Tiebreaker", tunnelPathSelection.Tiebreaker})
    tunnelPathSelection.EntityData.Leafs.Append("path-selection-hop-limit", types.YLeaf{"PathSelectionHopLimit", tunnelPathSelection.PathSelectionHopLimit})
    tunnelPathSelection.EntityData.Leafs.Append("path-selection-cost-limit", types.YLeaf{"PathSelectionCostLimit", tunnelPathSelection.PathSelectionCostLimit})

    tunnelPathSelection.EntityData.YListKeys = []string {}

    return &(tunnelPathSelection.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection_Invalidation
// Path invalidation configuration for this
// specific tunnel
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection_Invalidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Path Invalidation Timeout. The type is interface{} with range: 0..60000.
    PathInvalidationTimeout interface{}

    // Path Invalidation Action. The type is PathInvalidationAction.
    PathInvalidationAction interface{}
}

func (invalidation *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_TunnelPathSelection_Invalidation) GetEntityData() *types.CommonEntityData {
    invalidation.EntityData.YFilter = invalidation.YFilter
    invalidation.EntityData.YangName = "invalidation"
    invalidation.EntityData.BundleName = "cisco_ios_xr"
    invalidation.EntityData.ParentYangName = "tunnel-path-selection"
    invalidation.EntityData.SegmentPath = "invalidation"
    invalidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidation.EntityData.Children = types.NewOrderedMap()
    invalidation.EntityData.Leafs = types.NewOrderedMap()
    invalidation.EntityData.Leafs.Append("path-invalidation-timeout", types.YLeaf{"PathInvalidationTimeout", invalidation.PathInvalidationTimeout})
    invalidation.EntityData.Leafs.Append("path-invalidation-action", types.YLeaf{"PathInvalidationAction", invalidation.PathInvalidationAction})

    invalidation.EntityData.YListKeys = []string {}

    return &(invalidation.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth
// Tunnel Interface Auto-bandwidth configuration
// data
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable auto bandwidth underflow detection. The type is bool.
    UnderflowEnable interface{}

    // This object is only valid for tunnel interfaces and it controls whether
    // that interface has auto-bw enabled on it or not.The object must be set
    // before any other auto-bw configuration is supplied for the interface, and
    // must be the last auto-bw configuration object to be removed . The type is
    // bool.
    Enabled interface{}

    // Set the tunnel auto-bw application frequency in minutes. The type is
    // interface{} with range: 5..10080. Units are minute.
    ApplicationFrequency interface{}

    // Enable auto bandwidth overflow detection. The type is bool.
    OverflowEnable interface{}

    // Enable bandwidth collection only, no auto-bw adjustment. The type is
    // interface{}.
    CollectionOnly interface{}

    // Configuring the tunnel underflow detection.
    Underflow MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Underflow

    // Configuring the tunnel overflow detection.
    Overflow MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Overflow

    // Set min/max bandwidth auto-bw can apply on a tunnel.
    BandwidthLimits MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_BandwidthLimits

    // Set the bandwidth change threshold to trigger adjustment.
    AdjustmentThreshold MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_AdjustmentThreshold
}

func (autoBandwidth *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth) GetEntityData() *types.CommonEntityData {
    autoBandwidth.EntityData.YFilter = autoBandwidth.YFilter
    autoBandwidth.EntityData.YangName = "auto-bandwidth"
    autoBandwidth.EntityData.BundleName = "cisco_ios_xr"
    autoBandwidth.EntityData.ParentYangName = "tunnel-attributes"
    autoBandwidth.EntityData.SegmentPath = "auto-bandwidth"
    autoBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoBandwidth.EntityData.Children = types.NewOrderedMap()
    autoBandwidth.EntityData.Children.Append("underflow", types.YChild{"Underflow", &autoBandwidth.Underflow})
    autoBandwidth.EntityData.Children.Append("overflow", types.YChild{"Overflow", &autoBandwidth.Overflow})
    autoBandwidth.EntityData.Children.Append("bandwidth-limits", types.YChild{"BandwidthLimits", &autoBandwidth.BandwidthLimits})
    autoBandwidth.EntityData.Children.Append("adjustment-threshold", types.YChild{"AdjustmentThreshold", &autoBandwidth.AdjustmentThreshold})
    autoBandwidth.EntityData.Leafs = types.NewOrderedMap()
    autoBandwidth.EntityData.Leafs.Append("underflow-enable", types.YLeaf{"UnderflowEnable", autoBandwidth.UnderflowEnable})
    autoBandwidth.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", autoBandwidth.Enabled})
    autoBandwidth.EntityData.Leafs.Append("application-frequency", types.YLeaf{"ApplicationFrequency", autoBandwidth.ApplicationFrequency})
    autoBandwidth.EntityData.Leafs.Append("overflow-enable", types.YLeaf{"OverflowEnable", autoBandwidth.OverflowEnable})
    autoBandwidth.EntityData.Leafs.Append("collection-only", types.YLeaf{"CollectionOnly", autoBandwidth.CollectionOnly})

    autoBandwidth.EntityData.YListKeys = []string {}

    return &(autoBandwidth.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Underflow
// Configuring the tunnel underflow detection
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Underflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bandwidth change percent to trigger an underflow. The type is interface{}
    // with range: 1..100. This attribute is mandatory. Units are percentage.
    UnderflowThresholdPercent interface{}

    // Bandwidth change value to trigger an underflow (kbps). The type is
    // interface{} with range: 10..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    UnderflowThresholdValue interface{}

    // Number of consecutive collections exceeding threshold. The type is
    // interface{} with range: 1..10. This attribute is mandatory.
    UnderflowThresholdLimit interface{}
}

func (underflow *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Underflow) GetEntityData() *types.CommonEntityData {
    underflow.EntityData.YFilter = underflow.YFilter
    underflow.EntityData.YangName = "underflow"
    underflow.EntityData.BundleName = "cisco_ios_xr"
    underflow.EntityData.ParentYangName = "auto-bandwidth"
    underflow.EntityData.SegmentPath = "underflow"
    underflow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    underflow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    underflow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    underflow.EntityData.Children = types.NewOrderedMap()
    underflow.EntityData.Leafs = types.NewOrderedMap()
    underflow.EntityData.Leafs.Append("underflow-threshold-percent", types.YLeaf{"UnderflowThresholdPercent", underflow.UnderflowThresholdPercent})
    underflow.EntityData.Leafs.Append("underflow-threshold-value", types.YLeaf{"UnderflowThresholdValue", underflow.UnderflowThresholdValue})
    underflow.EntityData.Leafs.Append("underflow-threshold-limit", types.YLeaf{"UnderflowThresholdLimit", underflow.UnderflowThresholdLimit})

    underflow.EntityData.YListKeys = []string {}

    return &(underflow.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Overflow
// Configuring the tunnel overflow detection
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Overflow struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bandwidth change percent to trigger an overflow. The type is interface{}
    // with range: 1..100. This attribute is mandatory. Units are percentage.
    OverflowThresholdPercent interface{}

    // Bandwidth change value to trigger an overflow (kbps). The type is
    // interface{} with range: 10..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    OverflowThresholdValue interface{}

    // Number of consecutive collections exceeding threshold. The type is
    // interface{} with range: 1..10. This attribute is mandatory.
    OverflowThresholdLimit interface{}
}

func (overflow *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_Overflow) GetEntityData() *types.CommonEntityData {
    overflow.EntityData.YFilter = overflow.YFilter
    overflow.EntityData.YangName = "overflow"
    overflow.EntityData.BundleName = "cisco_ios_xr"
    overflow.EntityData.ParentYangName = "auto-bandwidth"
    overflow.EntityData.SegmentPath = "overflow"
    overflow.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overflow.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overflow.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overflow.EntityData.Children = types.NewOrderedMap()
    overflow.EntityData.Leafs = types.NewOrderedMap()
    overflow.EntityData.Leafs.Append("overflow-threshold-percent", types.YLeaf{"OverflowThresholdPercent", overflow.OverflowThresholdPercent})
    overflow.EntityData.Leafs.Append("overflow-threshold-value", types.YLeaf{"OverflowThresholdValue", overflow.OverflowThresholdValue})
    overflow.EntityData.Leafs.Append("overflow-threshold-limit", types.YLeaf{"OverflowThresholdLimit", overflow.OverflowThresholdLimit})

    overflow.EntityData.YListKeys = []string {}

    return &(overflow.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_BandwidthLimits
// Set min/max bandwidth auto-bw can apply on a
// tunnel
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_BandwidthLimits struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Set minimum bandwidth auto-bw can apply on a tunnel. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    BandwidthMinLimit interface{}

    // Set maximum bandwidth auto-bw can apply on a tunnel. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    BandwidthMaxLimit interface{}
}

func (bandwidthLimits *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_BandwidthLimits) GetEntityData() *types.CommonEntityData {
    bandwidthLimits.EntityData.YFilter = bandwidthLimits.YFilter
    bandwidthLimits.EntityData.YangName = "bandwidth-limits"
    bandwidthLimits.EntityData.BundleName = "cisco_ios_xr"
    bandwidthLimits.EntityData.ParentYangName = "auto-bandwidth"
    bandwidthLimits.EntityData.SegmentPath = "bandwidth-limits"
    bandwidthLimits.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidthLimits.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidthLimits.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidthLimits.EntityData.Children = types.NewOrderedMap()
    bandwidthLimits.EntityData.Leafs = types.NewOrderedMap()
    bandwidthLimits.EntityData.Leafs.Append("bandwidth-min-limit", types.YLeaf{"BandwidthMinLimit", bandwidthLimits.BandwidthMinLimit})
    bandwidthLimits.EntityData.Leafs.Append("bandwidth-max-limit", types.YLeaf{"BandwidthMaxLimit", bandwidthLimits.BandwidthMaxLimit})

    bandwidthLimits.EntityData.YListKeys = []string {}

    return &(bandwidthLimits.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_AdjustmentThreshold
// Set the bandwidth change threshold to trigger
// adjustment
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_AdjustmentThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bandwidth change percent to trigger adjustment. The type is interface{}
    // with range: 1..100. This attribute is mandatory. Units are percentage.
    AdjustmentThresholdPercent interface{}

    // Bandwidth change value to trigger adjustment (kbps). The type is
    // interface{} with range: 10..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    AdjustmentThresholdValue interface{}
}

func (adjustmentThreshold *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AutoBandwidth_AdjustmentThreshold) GetEntityData() *types.CommonEntityData {
    adjustmentThreshold.EntityData.YFilter = adjustmentThreshold.YFilter
    adjustmentThreshold.EntityData.YangName = "adjustment-threshold"
    adjustmentThreshold.EntityData.BundleName = "cisco_ios_xr"
    adjustmentThreshold.EntityData.ParentYangName = "auto-bandwidth"
    adjustmentThreshold.EntityData.SegmentPath = "adjustment-threshold"
    adjustmentThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjustmentThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjustmentThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjustmentThreshold.EntityData.Children = types.NewOrderedMap()
    adjustmentThreshold.EntityData.Leafs = types.NewOrderedMap()
    adjustmentThreshold.EntityData.Leafs.Append("adjustment-threshold-percent", types.YLeaf{"AdjustmentThresholdPercent", adjustmentThreshold.AdjustmentThresholdPercent})
    adjustmentThreshold.EntityData.Leafs.Append("adjustment-threshold-value", types.YLeaf{"AdjustmentThresholdValue", adjustmentThreshold.AdjustmentThresholdValue})

    adjustmentThreshold.EntityData.YListKeys = []string {}

    return &(adjustmentThreshold.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "tunnel-attributes"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", priority.SetupPriority})
    priority.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", priority.HoldPriority})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AffinityMask
// Set the affinity flags and mask
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AffinityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Mask interface{}
}

func (affinityMask *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_AffinityMask) GetEntityData() *types.CommonEntityData {
    affinityMask.EntityData.YFilter = affinityMask.YFilter
    affinityMask.EntityData.YangName = "affinity-mask"
    affinityMask.EntityData.BundleName = "cisco_ios_xr"
    affinityMask.EntityData.ParentYangName = "tunnel-attributes"
    affinityMask.EntityData.SegmentPath = "affinity-mask"
    affinityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMask.EntityData.Children = types.NewOrderedMap()
    affinityMask.EntityData.Leafs = types.NewOrderedMap()
    affinityMask.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", affinityMask.Affinity})
    affinityMask.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", affinityMask.Mask})

    affinityMask.EntityData.YListKeys = []string {}

    return &(affinityMask.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Logging
// Log tunnel LSP messages
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log tunnel messages for bandwidth change. The type is interface{}.
    LspSwitchOverChangeMessage interface{}

    // Log all events for a tunnel. The type is interface{}.
    All interface{}

    // Log tunnel record-route messages. The type is interface{}.
    RecordRouteMesssage interface{}

    // Enable BFD session state change alarm. The type is interface{}.
    BfdStateMessage interface{}

    // Log tunnel messages for bandwidth change. The type is interface{}.
    BandwidthChangeMessage interface{}

    // Log tunnel reoptimization attempts messages. The type is interface{}.
    ReoptimizeAttemptsMessage interface{}

    // Log tunnel rereoute messages. The type is interface{}.
    RerouteMesssage interface{}

    // Log tunnel state messages. The type is interface{}.
    StateMessage interface{}

    // Log tunnel messages for insufficient bandwidth. The type is interface{}.
    InsufficientBwMessage interface{}

    // Log tunnel reoptimized messages. The type is interface{}.
    ReoptimizedMessage interface{}

    // Enable logging for path-calculation failures. The type is interface{}.
    PcalcFailureMessage interface{}
}

func (logging *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "tunnel-attributes"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("lsp-switch-over-change-message", types.YLeaf{"LspSwitchOverChangeMessage", logging.LspSwitchOverChangeMessage})
    logging.EntityData.Leafs.Append("all", types.YLeaf{"All", logging.All})
    logging.EntityData.Leafs.Append("record-route-messsage", types.YLeaf{"RecordRouteMesssage", logging.RecordRouteMesssage})
    logging.EntityData.Leafs.Append("bfd-state-message", types.YLeaf{"BfdStateMessage", logging.BfdStateMessage})
    logging.EntityData.Leafs.Append("bandwidth-change-message", types.YLeaf{"BandwidthChangeMessage", logging.BandwidthChangeMessage})
    logging.EntityData.Leafs.Append("reoptimize-attempts-message", types.YLeaf{"ReoptimizeAttemptsMessage", logging.ReoptimizeAttemptsMessage})
    logging.EntityData.Leafs.Append("reroute-messsage", types.YLeaf{"RerouteMesssage", logging.RerouteMesssage})
    logging.EntityData.Leafs.Append("state-message", types.YLeaf{"StateMessage", logging.StateMessage})
    logging.EntityData.Leafs.Append("insufficient-bw-message", types.YLeaf{"InsufficientBwMessage", logging.InsufficientBwMessage})
    logging.EntityData.Leafs.Append("reoptimized-message", types.YLeaf{"ReoptimizedMessage", logging.ReoptimizedMessage})
    logging.EntityData.Leafs.Append("pcalc-failure-message", types.YLeaf{"PcalcFailureMessage", logging.PcalcFailureMessage})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Bandwidth
// Tunnel bandwidth requirement
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DSTE-standard flag. The type is MplsTeBandwidthDste. This attribute is
    // mandatory.
    DsteType interface{}

    // Class type for the bandwidth allocation. The type is interface{} with
    // range: 0..1. This attribute is mandatory.
    ClassOrPoolType interface{}

    // The value of the bandwidth reserved by this tunnel in kbps. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    Bandwidth interface{}
}

func (bandwidth *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "tunnel-attributes"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Leafs = types.NewOrderedMap()
    bandwidth.EntityData.Leafs.Append("dste-type", types.YLeaf{"DsteType", bandwidth.DsteType})
    bandwidth.EntityData.Leafs.Append("class-or-pool-type", types.YLeaf{"ClassOrPoolType", bandwidth.ClassOrPoolType})
    bandwidth.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", bandwidth.Bandwidth})

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute
// Parameters for IGP routing over tunnel
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Announce tunnel to IGP.
    AutorouteAnnounce MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce

    // Tunnel Autoroute Destination(s).
    Destinations MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations
}

func (autoroute *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute) GetEntityData() *types.CommonEntityData {
    autoroute.EntityData.YFilter = autoroute.YFilter
    autoroute.EntityData.YangName = "autoroute"
    autoroute.EntityData.BundleName = "cisco_ios_xr"
    autoroute.EntityData.ParentYangName = "tunnel-attributes"
    autoroute.EntityData.SegmentPath = "autoroute"
    autoroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoroute.EntityData.Children = types.NewOrderedMap()
    autoroute.EntityData.Children.Append("autoroute-announce", types.YChild{"AutorouteAnnounce", &autoroute.AutorouteAnnounce})
    autoroute.EntityData.Children.Append("destinations", types.YChild{"Destinations", &autoroute.Destinations})
    autoroute.EntityData.Leafs = types.NewOrderedMap()

    autoroute.EntityData.YListKeys = []string {}

    return &(autoroute.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce
// Announce tunnel to IGP
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable autoroute announce. The type is interface{}.
    Enable interface{}

    // Specify that the tunnel should be an IPv6 autoroute announce also. The type
    // is interface{}.
    IncludeIpv6 interface{}

    // Exclude traffic on autorouted tunnel.
    ExcludeTraffic MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_ExcludeTraffic

    // Specify MPLS tunnel metric.
    Metric MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_Metric
}

func (autorouteAnnounce *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce) GetEntityData() *types.CommonEntityData {
    autorouteAnnounce.EntityData.YFilter = autorouteAnnounce.YFilter
    autorouteAnnounce.EntityData.YangName = "autoroute-announce"
    autorouteAnnounce.EntityData.BundleName = "cisco_ios_xr"
    autorouteAnnounce.EntityData.ParentYangName = "autoroute"
    autorouteAnnounce.EntityData.SegmentPath = "autoroute-announce"
    autorouteAnnounce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autorouteAnnounce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autorouteAnnounce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autorouteAnnounce.EntityData.Children = types.NewOrderedMap()
    autorouteAnnounce.EntityData.Children.Append("exclude-traffic", types.YChild{"ExcludeTraffic", &autorouteAnnounce.ExcludeTraffic})
    autorouteAnnounce.EntityData.Children.Append("metric", types.YChild{"Metric", &autorouteAnnounce.Metric})
    autorouteAnnounce.EntityData.Leafs = types.NewOrderedMap()
    autorouteAnnounce.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autorouteAnnounce.Enable})
    autorouteAnnounce.EntityData.Leafs.Append("include-ipv6", types.YLeaf{"IncludeIpv6", autorouteAnnounce.IncludeIpv6})

    autorouteAnnounce.EntityData.YListKeys = []string {}

    return &(autorouteAnnounce.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_ExcludeTraffic
// Exclude traffic on autorouted tunnel
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_ExcludeTraffic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude tunnel in IGP for SR prefixes. The type is interface{}.
    SegmentRouting interface{}
}

func (excludeTraffic *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_ExcludeTraffic) GetEntityData() *types.CommonEntityData {
    excludeTraffic.EntityData.YFilter = excludeTraffic.YFilter
    excludeTraffic.EntityData.YangName = "exclude-traffic"
    excludeTraffic.EntityData.BundleName = "cisco_ios_xr"
    excludeTraffic.EntityData.ParentYangName = "autoroute-announce"
    excludeTraffic.EntityData.SegmentPath = "exclude-traffic"
    excludeTraffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeTraffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeTraffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeTraffic.EntityData.Children = types.NewOrderedMap()
    excludeTraffic.EntityData.Leafs = types.NewOrderedMap()
    excludeTraffic.EntityData.Leafs.Append("segment-routing", types.YLeaf{"SegmentRouting", excludeTraffic.SegmentRouting})

    excludeTraffic.EntityData.YListKeys = []string {}

    return &(excludeTraffic.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_Metric
// Specify MPLS tunnel metric
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_Metric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autoroute tunnel metric type. The type is MplsTeAutorouteMetric.
    MetricType interface{}

    // The absolute metric value. The type is interface{} with range:
    // 1..2147483647.
    AbsoluteMetric interface{}

    // The value of the adjustment. The type is interface{} with range: -10..10.
    RelativeMetric interface{}

    // The constant metric value. The type is interface{} with range:
    // 1..2147483647.
    ConstantMetric interface{}
}

func (metric *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_AutorouteAnnounce_Metric) GetEntityData() *types.CommonEntityData {
    metric.EntityData.YFilter = metric.YFilter
    metric.EntityData.YangName = "metric"
    metric.EntityData.BundleName = "cisco_ios_xr"
    metric.EntityData.ParentYangName = "autoroute-announce"
    metric.EntityData.SegmentPath = "metric"
    metric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    metric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    metric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    metric.EntityData.Children = types.NewOrderedMap()
    metric.EntityData.Leafs = types.NewOrderedMap()
    metric.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", metric.MetricType})
    metric.EntityData.Leafs.Append("absolute-metric", types.YLeaf{"AbsoluteMetric", metric.AbsoluteMetric})
    metric.EntityData.Leafs.Append("relative-metric", types.YLeaf{"RelativeMetric", metric.RelativeMetric})
    metric.EntityData.Leafs.Append("constant-metric", types.YLeaf{"ConstantMetric", metric.ConstantMetric})

    metric.EntityData.YListKeys = []string {}

    return &(metric.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations
// Tunnel Autoroute Destination(s)
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Destination address to add in RIB. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations_Destination.
    Destination []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations_Destination
}

func (destinations *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations) GetEntityData() *types.CommonEntityData {
    destinations.EntityData.YFilter = destinations.YFilter
    destinations.EntityData.YangName = "destinations"
    destinations.EntityData.BundleName = "cisco_ios_xr"
    destinations.EntityData.ParentYangName = "autoroute"
    destinations.EntityData.SegmentPath = "destinations"
    destinations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinations.EntityData.Children = types.NewOrderedMap()
    destinations.EntityData.Children.Append("destination", types.YChild{"Destination", nil})
    for i := range destinations.Destination {
        destinations.EntityData.Children.Append(types.GetSegmentPath(destinations.Destination[i]), types.YChild{"Destination", destinations.Destination[i]})
    }
    destinations.EntityData.Leafs = types.NewOrderedMap()

    destinations.EntityData.YListKeys = []string {}

    return &(destinations.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations_Destination
// Destination address to add in RIB
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of destination. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}
}

func (destination *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_Autoroute_Destinations_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "destinations"
    destination.EntityData.SegmentPath = "destination" + types.AddKeyToken(destination.DestinationAddress, "destination-address")
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", destination.DestinationAddress})

    destination.EntityData.YListKeys = []string {"DestinationAddress"}

    return &(destination.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes
// Tunnel new style affinity attributes table
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType.
    NewStyleAffinityAffinityType []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1.
    NewStyleAffinityAffinityTypeAffinity1 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2.
    NewStyleAffinityAffinityTypeAffinity1Affinity2 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 []*MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
}

func (newStyleAffinityAffinityTypes *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypes.EntityData.YFilter = newStyleAffinityAffinityTypes.YFilter
    newStyleAffinityAffinityTypes.EntityData.YangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypes.EntityData.ParentYangName = "tunnel-attributes"
    newStyleAffinityAffinityTypes.EntityData.SegmentPath = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypes.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type", types.YChild{"NewStyleAffinityAffinityType", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]), types.YChild{"NewStyleAffinityAffinityType", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1", types.YChild{"NewStyleAffinityAffinityTypeAffinity1", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Leafs = types.NewOrderedMap()

    newStyleAffinityAffinityTypes.EntityData.YListKeys = []string {}

    return &(newStyleAffinityAffinityTypes.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}
}

func (newStyleAffinityAffinityType *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityType.EntityData.YFilter = newStyleAffinityAffinityType.YFilter
    newStyleAffinityAffinityType.EntityData.YangName = "new-style-affinity-affinity-type"
    newStyleAffinityAffinityType.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityType.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityType.EntityData.SegmentPath = "new-style-affinity-affinity-type" + types.AddKeyToken(newStyleAffinityAffinityType.AffinityType, "affinity-type")
    newStyleAffinityAffinityType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityType.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityType.AffinityType})

    newStyleAffinityAffinityType.EntityData.YListKeys = []string {"AffinityType"}

    return &(newStyleAffinityAffinityType.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1.YFilter
    newStyleAffinityAffinityTypeAffinity1.EntityData.YangName = "new-style-affinity-affinity-type-affinity1"
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.Affinity1, "affinity1")
    newStyleAffinityAffinityTypeAffinity1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1.AffinityType})
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1.Affinity1})

    newStyleAffinityAffinityTypeAffinity1.EntityData.YListKeys = []string {"AffinityType", "Affinity1"}

    return &(newStyleAffinityAffinityTypeAffinity1.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2, "affinity2")
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2})

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3, "affinity3")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4, "affinity4")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5, "affinity5")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6, "affinity6")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7, "affinity7")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8, "affinity8")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9, "affinity9")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
// Tunnel new style affinity attribute
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}

    // This attribute is a key. The name of the tenth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity10 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9, "affinity9") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10, "affinity10")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity10", types.YLeaf{"Affinity10", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9", "Affinity10"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_FastReroute
// Specify MPLS tunnel can be fast-rerouted
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bandwidth Protection. The type is interface{} with range: 0..1. This
    // attribute is mandatory.
    BandwidthProtection interface{}

    // Node Protection. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    NodeProtection interface{}
}

func (fastReroute *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelAttributes_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "tunnel-attributes"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("bandwidth-protection", types.YLeaf{"BandwidthProtection", fastReroute.BandwidthProtection})
    fastReroute.EntityData.Leafs.Append("node-protection", types.YLeaf{"NodeProtection", fastReroute.NodeProtection})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelId
// Set the tunnel ID
// This type is a presence type.
type MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Tunnel ID Type. The type is MplsTeTunnelId. This attribute is mandatory.
    TunnelIdType interface{}

    // Tunnel ID. The type is interface{} with range: 0..65535.
    TunnelId interface{}
}

func (tunnelId *MplsTe_NamedTunnels_Tunnels_Tunnel_TunnelId) GetEntityData() *types.CommonEntityData {
    tunnelId.EntityData.YFilter = tunnelId.YFilter
    tunnelId.EntityData.YangName = "tunnel-id"
    tunnelId.EntityData.BundleName = "cisco_ios_xr"
    tunnelId.EntityData.ParentYangName = "tunnel"
    tunnelId.EntityData.SegmentPath = "tunnel-id"
    tunnelId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelId.EntityData.Children = types.NewOrderedMap()
    tunnelId.EntityData.Leafs = types.NewOrderedMap()
    tunnelId.EntityData.Leafs.Append("tunnel-id-type", types.YLeaf{"TunnelIdType", tunnelId.TunnelIdType})
    tunnelId.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", tunnelId.TunnelId})

    tunnelId.EntityData.YListKeys = []string {}

    return &(tunnelId.EntityData)
}

// MplsTe_GmplsUni
// GMPLS-UNI configuration
type MplsTe_GmplsUni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GMPLS-UNI timer configuration.
    Timers MplsTe_GmplsUni_Timers

    // GMPLS-UNI controllers.
    Controllers MplsTe_GmplsUni_Controllers
}

func (gmplsUni *MplsTe_GmplsUni) GetEntityData() *types.CommonEntityData {
    gmplsUni.EntityData.YFilter = gmplsUni.YFilter
    gmplsUni.EntityData.YangName = "gmpls-uni"
    gmplsUni.EntityData.BundleName = "cisco_ios_xr"
    gmplsUni.EntityData.ParentYangName = "mpls-te"
    gmplsUni.EntityData.SegmentPath = "gmpls-uni"
    gmplsUni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmplsUni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmplsUni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmplsUni.EntityData.Children = types.NewOrderedMap()
    gmplsUni.EntityData.Children.Append("timers", types.YChild{"Timers", &gmplsUni.Timers})
    gmplsUni.EntityData.Children.Append("controllers", types.YChild{"Controllers", &gmplsUni.Controllers})
    gmplsUni.EntityData.Leafs = types.NewOrderedMap()

    gmplsUni.EntityData.YListKeys = []string {}

    return &(gmplsUni.EntityData)
}

// MplsTe_GmplsUni_Timers
// GMPLS-UNI timer configuration
type MplsTe_GmplsUni_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GMPLS-UNI path-option timer configuration.
    PathOptionTimers MplsTe_GmplsUni_Timers_PathOptionTimers
}

func (timers *MplsTe_GmplsUni_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "gmpls-uni"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("path-option-timers", types.YChild{"PathOptionTimers", &timers.PathOptionTimers})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// MplsTe_GmplsUni_Timers_PathOptionTimers
// GMPLS-UNI path-option timer configuration
type MplsTe_GmplsUni_Timers_PathOptionTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GMPLS-UNI path-option holddown timer configuration.
    Holddown MplsTe_GmplsUni_Timers_PathOptionTimers_Holddown
}

func (pathOptionTimers *MplsTe_GmplsUni_Timers_PathOptionTimers) GetEntityData() *types.CommonEntityData {
    pathOptionTimers.EntityData.YFilter = pathOptionTimers.YFilter
    pathOptionTimers.EntityData.YangName = "path-option-timers"
    pathOptionTimers.EntityData.BundleName = "cisco_ios_xr"
    pathOptionTimers.EntityData.ParentYangName = "timers"
    pathOptionTimers.EntityData.SegmentPath = "path-option-timers"
    pathOptionTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOptionTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOptionTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOptionTimers.EntityData.Children = types.NewOrderedMap()
    pathOptionTimers.EntityData.Children.Append("holddown", types.YChild{"Holddown", &pathOptionTimers.Holddown})
    pathOptionTimers.EntityData.Leafs = types.NewOrderedMap()

    pathOptionTimers.EntityData.YListKeys = []string {}

    return &(pathOptionTimers.EntityData)
}

// MplsTe_GmplsUni_Timers_PathOptionTimers_Holddown
// GMPLS-UNI path-option holddown timer
// configuration
type MplsTe_GmplsUni_Timers_PathOptionTimers_Holddown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum holddown (seconds). The type is interface{} with range: 5..3600.
    // Units are second.
    Minimum interface{}

    // Maximum holddown (seconds). The type is interface{} with range: 5..3600.
    // Units are second.
    Maximum interface{}
}

func (holddown *MplsTe_GmplsUni_Timers_PathOptionTimers_Holddown) GetEntityData() *types.CommonEntityData {
    holddown.EntityData.YFilter = holddown.YFilter
    holddown.EntityData.YangName = "holddown"
    holddown.EntityData.BundleName = "cisco_ios_xr"
    holddown.EntityData.ParentYangName = "path-option-timers"
    holddown.EntityData.SegmentPath = "holddown"
    holddown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    holddown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    holddown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    holddown.EntityData.Children = types.NewOrderedMap()
    holddown.EntityData.Leafs = types.NewOrderedMap()
    holddown.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", holddown.Minimum})
    holddown.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", holddown.Maximum})

    holddown.EntityData.YListKeys = []string {}

    return &(holddown.EntityData)
}

// MplsTe_GmplsUni_Controllers
// GMPLS-UNI controllers
type MplsTe_GmplsUni_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a GMPLS controller. The type is slice of
    // MplsTe_GmplsUni_Controllers_Controller.
    Controller []*MplsTe_GmplsUni_Controllers_Controller
}

func (controllers *MplsTe_GmplsUni_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "gmpls-uni"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller
// Configure a GMPLS controller
type MplsTe_GmplsUni_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // Enable GMPLS-UNI on the link. The type is interface{}.
    Enable interface{}

    // Announce discovered tunnel properties to system.
    Announce MplsTe_GmplsUni_Controllers_Controller_Announce

    // Controller logging.
    ControllerLogging MplsTe_GmplsUni_Controllers_Controller_ControllerLogging

    // GMPLS-UNI tunnel-head properties.
    GmplsUnitunnelHead MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead
}

func (controller *MplsTe_GmplsUni_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("announce", types.YChild{"Announce", &controller.Announce})
    controller.EntityData.Children.Append("controller-logging", types.YChild{"ControllerLogging", &controller.ControllerLogging})
    controller.EntityData.Children.Append("gmpls-unitunnel-head", types.YChild{"GmplsUnitunnelHead", &controller.GmplsUnitunnelHead})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})
    controller.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", controller.Enable})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_Announce
// Announce discovered tunnel properties to
// system
type MplsTe_GmplsUni_Controllers_Controller_Announce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable announcement of discovered SRLGs. The type is interface{}.
    SrlGs interface{}
}

func (announce *MplsTe_GmplsUni_Controllers_Controller_Announce) GetEntityData() *types.CommonEntityData {
    announce.EntityData.YFilter = announce.YFilter
    announce.EntityData.YangName = "announce"
    announce.EntityData.BundleName = "cisco_ios_xr"
    announce.EntityData.ParentYangName = "controller"
    announce.EntityData.SegmentPath = "announce"
    announce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    announce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    announce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    announce.EntityData.Children = types.NewOrderedMap()
    announce.EntityData.Leafs = types.NewOrderedMap()
    announce.EntityData.Leafs.Append("srl-gs", types.YLeaf{"SrlGs", announce.SrlGs})

    announce.EntityData.YListKeys = []string {}

    return &(announce.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_ControllerLogging
// Controller logging
type MplsTe_GmplsUni_Controllers_Controller_ControllerLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable logging of changes to of discovered SRLGs. The type is interface{}.
    DiscoveredSrlgChangeLogging interface{}
}

func (controllerLogging *MplsTe_GmplsUni_Controllers_Controller_ControllerLogging) GetEntityData() *types.CommonEntityData {
    controllerLogging.EntityData.YFilter = controllerLogging.YFilter
    controllerLogging.EntityData.YangName = "controller-logging"
    controllerLogging.EntityData.BundleName = "cisco_ios_xr"
    controllerLogging.EntityData.ParentYangName = "controller"
    controllerLogging.EntityData.SegmentPath = "controller-logging"
    controllerLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllerLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllerLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllerLogging.EntityData.Children = types.NewOrderedMap()
    controllerLogging.EntityData.Leafs = types.NewOrderedMap()
    controllerLogging.EntityData.Leafs.Append("discovered-srlg-change-logging", types.YLeaf{"DiscoveredSrlgChangeLogging", controllerLogging.DiscoveredSrlgChangeLogging})

    controllerLogging.EntityData.YListKeys = []string {}

    return &(controllerLogging.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead
// GMPLS-UNI tunnel-head properties
type MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GMPLS-UNI head tunnel-id. The type is interface{} with range: 0..65535.
    TunnelId interface{}

    // Set link as a GMPLS tunnel head. The type is interface{}.
    Enable interface{}

    // Set the destination of the tunnel. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Destination interface{}

    // Record the route used by the tunnel. The type is interface{}.
    RecordRoute interface{}

    // The name of the tunnel to be included in signalling messages. The type is
    // string with length: 1..254.
    SignalledName interface{}

    // Path-option configuration.
    PathOptions MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions

    // Tunnel property recording.
    Recording MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Recording

    // Tunnel event logging.
    Logging MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Logging

    // Tunnel Setup and Hold Priorities.
    Priority MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Priority
}

func (gmplsUnitunnelHead *MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead) GetEntityData() *types.CommonEntityData {
    gmplsUnitunnelHead.EntityData.YFilter = gmplsUnitunnelHead.YFilter
    gmplsUnitunnelHead.EntityData.YangName = "gmpls-unitunnel-head"
    gmplsUnitunnelHead.EntityData.BundleName = "cisco_ios_xr"
    gmplsUnitunnelHead.EntityData.ParentYangName = "controller"
    gmplsUnitunnelHead.EntityData.SegmentPath = "gmpls-unitunnel-head"
    gmplsUnitunnelHead.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmplsUnitunnelHead.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmplsUnitunnelHead.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmplsUnitunnelHead.EntityData.Children = types.NewOrderedMap()
    gmplsUnitunnelHead.EntityData.Children.Append("path-options", types.YChild{"PathOptions", &gmplsUnitunnelHead.PathOptions})
    gmplsUnitunnelHead.EntityData.Children.Append("recording", types.YChild{"Recording", &gmplsUnitunnelHead.Recording})
    gmplsUnitunnelHead.EntityData.Children.Append("logging", types.YChild{"Logging", &gmplsUnitunnelHead.Logging})
    gmplsUnitunnelHead.EntityData.Children.Append("priority", types.YChild{"Priority", &gmplsUnitunnelHead.Priority})
    gmplsUnitunnelHead.EntityData.Leafs = types.NewOrderedMap()
    gmplsUnitunnelHead.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", gmplsUnitunnelHead.TunnelId})
    gmplsUnitunnelHead.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", gmplsUnitunnelHead.Enable})
    gmplsUnitunnelHead.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", gmplsUnitunnelHead.Destination})
    gmplsUnitunnelHead.EntityData.Leafs.Append("record-route", types.YLeaf{"RecordRoute", gmplsUnitunnelHead.RecordRoute})
    gmplsUnitunnelHead.EntityData.Leafs.Append("signalled-name", types.YLeaf{"SignalledName", gmplsUnitunnelHead.SignalledName})

    gmplsUnitunnelHead.EntityData.YListKeys = []string {}

    return &(gmplsUnitunnelHead.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions
// Path-option configuration
type MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A Path-option. The type is slice of
    // MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions_PathOption.
    PathOption []*MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions_PathOption
}

func (pathOptions *MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions) GetEntityData() *types.CommonEntityData {
    pathOptions.EntityData.YFilter = pathOptions.YFilter
    pathOptions.EntityData.YangName = "path-options"
    pathOptions.EntityData.BundleName = "cisco_ios_xr"
    pathOptions.EntityData.ParentYangName = "gmpls-unitunnel-head"
    pathOptions.EntityData.SegmentPath = "path-options"
    pathOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOptions.EntityData.Children = types.NewOrderedMap()
    pathOptions.EntityData.Children.Append("path-option", types.YChild{"PathOption", nil})
    for i := range pathOptions.PathOption {
        pathOptions.EntityData.Children.Append(types.GetSegmentPath(pathOptions.PathOption[i]), types.YChild{"PathOption", pathOptions.PathOption[i]})
    }
    pathOptions.EntityData.Leafs = types.NewOrderedMap()

    pathOptions.EntityData.YListKeys = []string {}

    return &(pathOptions.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions_PathOption
// A Path-option
type MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions_PathOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Preference level for this path option. The type is
    // interface{} with range: 1..1000.
    PreferenceLevel interface{}

    // The type of the path option. The type is MplsTePathOption. This attribute
    // is mandatory.
    PathType interface{}

    // The ID of the explicit path associated with this option. The type is
    // interface{} with range: 1..65535. The default value is 1.
    PathId interface{}

    // The name of the explicit path associated with this option. The type is
    // string.
    PathName interface{}

    // The route-exclusion type. The type is interface{}. This attribute is
    // mandatory.
    XroType interface{}

    // The name of the XRO attribute set to be used for this path-option. The type
    // is string with length: 1..64.
    XroAttributeSetName interface{}

    // Path option properties: must be Lockdown. The type is
    // MplsTePathOptionProperty. This attribute is mandatory.
    Lockdown interface{}

    // Path option properties: must be verbatim if set. The type is
    // MplsTePathOptionProperty. The default value is none.
    Verbatim interface{}

    // Signaled label type. The type is MplsTeSignaledLabel. The default value is
    // not-set.
    SignaledLabel interface{}

    // DWDM channel number. The type is interface{} with range: 1..89. The default
    // value is 1.
    DwdmChannel interface{}
}

func (pathOption *MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_PathOptions_PathOption) GetEntityData() *types.CommonEntityData {
    pathOption.EntityData.YFilter = pathOption.YFilter
    pathOption.EntityData.YangName = "path-option"
    pathOption.EntityData.BundleName = "cisco_ios_xr"
    pathOption.EntityData.ParentYangName = "path-options"
    pathOption.EntityData.SegmentPath = "path-option" + types.AddKeyToken(pathOption.PreferenceLevel, "preference-level")
    pathOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOption.EntityData.Children = types.NewOrderedMap()
    pathOption.EntityData.Leafs = types.NewOrderedMap()
    pathOption.EntityData.Leafs.Append("preference-level", types.YLeaf{"PreferenceLevel", pathOption.PreferenceLevel})
    pathOption.EntityData.Leafs.Append("path-type", types.YLeaf{"PathType", pathOption.PathType})
    pathOption.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", pathOption.PathId})
    pathOption.EntityData.Leafs.Append("path-name", types.YLeaf{"PathName", pathOption.PathName})
    pathOption.EntityData.Leafs.Append("xro-type", types.YLeaf{"XroType", pathOption.XroType})
    pathOption.EntityData.Leafs.Append("xro-attribute-set-name", types.YLeaf{"XroAttributeSetName", pathOption.XroAttributeSetName})
    pathOption.EntityData.Leafs.Append("lockdown", types.YLeaf{"Lockdown", pathOption.Lockdown})
    pathOption.EntityData.Leafs.Append("verbatim", types.YLeaf{"Verbatim", pathOption.Verbatim})
    pathOption.EntityData.Leafs.Append("signaled-label", types.YLeaf{"SignaledLabel", pathOption.SignaledLabel})
    pathOption.EntityData.Leafs.Append("dwdm-channel", types.YLeaf{"DwdmChannel", pathOption.DwdmChannel})

    pathOption.EntityData.YListKeys = []string {"PreferenceLevel"}

    return &(pathOption.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Recording
// Tunnel property recording
type MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Recording struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SRLG-recording during signaling. The type is interface{}.
    Srlg interface{}
}

func (recording *MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Recording) GetEntityData() *types.CommonEntityData {
    recording.EntityData.YFilter = recording.YFilter
    recording.EntityData.YangName = "recording"
    recording.EntityData.BundleName = "cisco_ios_xr"
    recording.EntityData.ParentYangName = "gmpls-unitunnel-head"
    recording.EntityData.SegmentPath = "recording"
    recording.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    recording.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    recording.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    recording.EntityData.Children = types.NewOrderedMap()
    recording.EntityData.Leafs = types.NewOrderedMap()
    recording.EntityData.Leafs.Append("srlg", types.YLeaf{"Srlg", recording.Srlg})

    recording.EntityData.YListKeys = []string {}

    return &(recording.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Logging
// Tunnel event logging
type MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log tunnel state messages. The type is interface{}.
    StateMessage interface{}
}

func (logging *MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "gmpls-unitunnel-head"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("state-message", types.YLeaf{"StateMessage", logging.StateMessage})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *MplsTe_GmplsUni_Controllers_Controller_GmplsUnitunnelHead_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "gmpls-unitunnel-head"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", priority.SetupPriority})
    priority.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", priority.HoldPriority})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// MplsTe_GlobalAttributes
// Configure MPLS TE global attributes
type MplsTe_GlobalAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log NSR status messages. The type is interface{}.
    LogNsrStatus interface{}

    // Log ISSU status messages. The type is interface{}.
    LogIssuStatus interface{}

    // Enable reoptimization based on link-up events. The type is interface{}.
    ReoptimizeLinkUp interface{}

    // Reoptimization Delay Cleanup Value (seconds). The type is interface{} with
    // range: 0..300. Units are second.
    ReoptimizeDelayCleanupTimer interface{}

    // Disable reoptimization after affinity failures. The type is interface{}.
    DisableReoptimizeAffinityFailure interface{}

    // The maximum number of tunnel heads that will be allowed. The type is
    // interface{} with range: 1..65536. The default value is 4096.
    MaximumTunnels interface{}

    // Holddown time for links which had Path Errors in seconds. The type is
    // interface{} with range: 0..300. Units are second. The default value is 10.
    LinkHolddownTimer interface{}

    // Enable Fault-OAM functionality for bidirectional tunnels. The type is
    // interface{}.
    FaultOam interface{}

    // Enable unequal load-balancing over tunnels to the same destination. The
    // type is interface{}.
    EnableUnequalLoadBalancing interface{}

    // Log all tail tunnel events. The type is interface{}.
    LogTail interface{}

    // Reoptimization Delay After FRR Value (seconds). The type is interface{}
    // with range: 0..120. Units are second.
    ReoptimizeDelayAfterFrrTimer interface{}

    // Auto-bandwidth global collection frequency in minutes. The type is
    // interface{} with range: 1..10080. Units are minute. The default value is 5.
    AutoBandwidthCollectFrequency interface{}

    // Seconds between path protect switchover and tunnel re-optimization. Set to
    // 0 to disable. The type is interface{} with range: 0..604800. Units are
    // second. The default value is 180.
    ReoptDelayPathProtectSwitchoverTimer interface{}

    // Always set to true. The type is interface{}.
    LogAll interface{}

    // Signalling retry for tunnels terminating outside the headend area. The type
    // is interface{} with range: 30..600. The default value is 120.
    LoosePathRetryPeriod interface{}

    // Load balance bandwidth during reoptimization. The type is interface{}.
    ReoptimizeLoadBalancing interface{}

    // Log all head tunnel events. The type is interface{}.
    LogHead interface{}

    // Deprecated - do not use. The type is interface{}.
    PathSelectionIgnoreOverload interface{}

    // Enable graceful preemption when there is a bandwidth reduction. The type is
    // interface{}.
    GracefulPreemptionOnBandwidthReduction interface{}

    // Enable explicit-null advertising to PHOP. The type is interface{}.
    AdvertiseExplicitNulls interface{}

    // Reoptimization Delay Install Value (seconds). The type is interface{} with
    // range: 0..3600. Units are second.
    ReoptimizeDelayInstallTimer interface{}

    // Delay reoptimizing current LSP after affinity failures. The type is
    // interface{} with range: 1..604800. Units are second.
    ReoptimizeDelayAfterAffinityFailureTimer interface{}

    // Log FRR Protection messages. The type is MplsTeLogFrrProtection.
    LogFrrProtection interface{}

    // Reoptimize timers period in seconds. The type is interface{} with range:
    // 0..604800. Units are second. The default value is 3600.
    ReoptimizeTimerFrequency interface{}

    // Log all mid tunnel events. The type is interface{}.
    LogMid interface{}

    // Log tunnel preemption messages. The type is interface{}.
    LogPreemption interface{}

    // Configure auto-tunnels feature.
    AutoTunnel MplsTe_GlobalAttributes_AutoTunnel

    // Configure HW OOR processing in MPLS-TE.
    HardwareOutOfResource MplsTe_GlobalAttributes_HardwareOutOfResource

    // Configure MPLS TE Secondary Router ID.
    SecondaryRouterIds MplsTe_GlobalAttributes_SecondaryRouterIds

    // Configure SRLG values and MPLS-TE properties.
    Srlg MplsTe_GlobalAttributes_Srlg

    // Configure MPLS TE route priority.
    Queues MplsTe_GlobalAttributes_Queues

    // MPLS-TE MIB properties.
    Mib MplsTe_GlobalAttributes_Mib

    // Attribute AttributeSets.
    AttributeSet MplsTe_GlobalAttributes_AttributeSet

    // BFD over MPLS TE Global Configurations.
    BfdOverLsp MplsTe_GlobalAttributes_BfdOverLsp

    // Bandwidth accounting configuration data.
    BandwidthAccounting MplsTe_GlobalAttributes_BandwidthAccounting

    // Configuration MPLS TE PCE attributes.
    PceAttributes MplsTe_GlobalAttributes_PceAttributes

    // Configure LSP OOR attributes in MPLS-TE.
    LspOutOfResource MplsTe_GlobalAttributes_LspOutOfResource

    // Soft preemption configuration data.
    SoftPreemption MplsTe_GlobalAttributes_SoftPreemption

    // Configure fast reroute attributes.
    FastReroute MplsTe_GlobalAttributes_FastReroute

    // Path selection configuration.
    PathSelection MplsTe_GlobalAttributes_PathSelection

    // Affinity Mapping Table configuration.
    AffinityMappings MplsTe_GlobalAttributes_AffinityMappings
}

func (globalAttributes *MplsTe_GlobalAttributes) GetEntityData() *types.CommonEntityData {
    globalAttributes.EntityData.YFilter = globalAttributes.YFilter
    globalAttributes.EntityData.YangName = "global-attributes"
    globalAttributes.EntityData.BundleName = "cisco_ios_xr"
    globalAttributes.EntityData.ParentYangName = "mpls-te"
    globalAttributes.EntityData.SegmentPath = "global-attributes"
    globalAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAttributes.EntityData.Children = types.NewOrderedMap()
    globalAttributes.EntityData.Children.Append("auto-tunnel", types.YChild{"AutoTunnel", &globalAttributes.AutoTunnel})
    globalAttributes.EntityData.Children.Append("hardware-out-of-resource", types.YChild{"HardwareOutOfResource", &globalAttributes.HardwareOutOfResource})
    globalAttributes.EntityData.Children.Append("secondary-router-ids", types.YChild{"SecondaryRouterIds", &globalAttributes.SecondaryRouterIds})
    globalAttributes.EntityData.Children.Append("srlg", types.YChild{"Srlg", &globalAttributes.Srlg})
    globalAttributes.EntityData.Children.Append("queues", types.YChild{"Queues", &globalAttributes.Queues})
    globalAttributes.EntityData.Children.Append("mib", types.YChild{"Mib", &globalAttributes.Mib})
    globalAttributes.EntityData.Children.Append("attribute-set", types.YChild{"AttributeSet", &globalAttributes.AttributeSet})
    globalAttributes.EntityData.Children.Append("bfd-over-lsp", types.YChild{"BfdOverLsp", &globalAttributes.BfdOverLsp})
    globalAttributes.EntityData.Children.Append("bandwidth-accounting", types.YChild{"BandwidthAccounting", &globalAttributes.BandwidthAccounting})
    globalAttributes.EntityData.Children.Append("pce-attributes", types.YChild{"PceAttributes", &globalAttributes.PceAttributes})
    globalAttributes.EntityData.Children.Append("lsp-out-of-resource", types.YChild{"LspOutOfResource", &globalAttributes.LspOutOfResource})
    globalAttributes.EntityData.Children.Append("soft-preemption", types.YChild{"SoftPreemption", &globalAttributes.SoftPreemption})
    globalAttributes.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &globalAttributes.FastReroute})
    globalAttributes.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &globalAttributes.PathSelection})
    globalAttributes.EntityData.Children.Append("affinity-mappings", types.YChild{"AffinityMappings", &globalAttributes.AffinityMappings})
    globalAttributes.EntityData.Leafs = types.NewOrderedMap()
    globalAttributes.EntityData.Leafs.Append("log-nsr-status", types.YLeaf{"LogNsrStatus", globalAttributes.LogNsrStatus})
    globalAttributes.EntityData.Leafs.Append("log-issu-status", types.YLeaf{"LogIssuStatus", globalAttributes.LogIssuStatus})
    globalAttributes.EntityData.Leafs.Append("reoptimize-link-up", types.YLeaf{"ReoptimizeLinkUp", globalAttributes.ReoptimizeLinkUp})
    globalAttributes.EntityData.Leafs.Append("reoptimize-delay-cleanup-timer", types.YLeaf{"ReoptimizeDelayCleanupTimer", globalAttributes.ReoptimizeDelayCleanupTimer})
    globalAttributes.EntityData.Leafs.Append("disable-reoptimize-affinity-failure", types.YLeaf{"DisableReoptimizeAffinityFailure", globalAttributes.DisableReoptimizeAffinityFailure})
    globalAttributes.EntityData.Leafs.Append("maximum-tunnels", types.YLeaf{"MaximumTunnels", globalAttributes.MaximumTunnels})
    globalAttributes.EntityData.Leafs.Append("link-holddown-timer", types.YLeaf{"LinkHolddownTimer", globalAttributes.LinkHolddownTimer})
    globalAttributes.EntityData.Leafs.Append("fault-oam", types.YLeaf{"FaultOam", globalAttributes.FaultOam})
    globalAttributes.EntityData.Leafs.Append("enable-unequal-load-balancing", types.YLeaf{"EnableUnequalLoadBalancing", globalAttributes.EnableUnequalLoadBalancing})
    globalAttributes.EntityData.Leafs.Append("log-tail", types.YLeaf{"LogTail", globalAttributes.LogTail})
    globalAttributes.EntityData.Leafs.Append("reoptimize-delay-after-frr-timer", types.YLeaf{"ReoptimizeDelayAfterFrrTimer", globalAttributes.ReoptimizeDelayAfterFrrTimer})
    globalAttributes.EntityData.Leafs.Append("auto-bandwidth-collect-frequency", types.YLeaf{"AutoBandwidthCollectFrequency", globalAttributes.AutoBandwidthCollectFrequency})
    globalAttributes.EntityData.Leafs.Append("reopt-delay-path-protect-switchover-timer", types.YLeaf{"ReoptDelayPathProtectSwitchoverTimer", globalAttributes.ReoptDelayPathProtectSwitchoverTimer})
    globalAttributes.EntityData.Leafs.Append("log-all", types.YLeaf{"LogAll", globalAttributes.LogAll})
    globalAttributes.EntityData.Leafs.Append("loose-path-retry-period", types.YLeaf{"LoosePathRetryPeriod", globalAttributes.LoosePathRetryPeriod})
    globalAttributes.EntityData.Leafs.Append("reoptimize-load-balancing", types.YLeaf{"ReoptimizeLoadBalancing", globalAttributes.ReoptimizeLoadBalancing})
    globalAttributes.EntityData.Leafs.Append("log-head", types.YLeaf{"LogHead", globalAttributes.LogHead})
    globalAttributes.EntityData.Leafs.Append("path-selection-ignore-overload", types.YLeaf{"PathSelectionIgnoreOverload", globalAttributes.PathSelectionIgnoreOverload})
    globalAttributes.EntityData.Leafs.Append("graceful-preemption-on-bandwidth-reduction", types.YLeaf{"GracefulPreemptionOnBandwidthReduction", globalAttributes.GracefulPreemptionOnBandwidthReduction})
    globalAttributes.EntityData.Leafs.Append("advertise-explicit-nulls", types.YLeaf{"AdvertiseExplicitNulls", globalAttributes.AdvertiseExplicitNulls})
    globalAttributes.EntityData.Leafs.Append("reoptimize-delay-install-timer", types.YLeaf{"ReoptimizeDelayInstallTimer", globalAttributes.ReoptimizeDelayInstallTimer})
    globalAttributes.EntityData.Leafs.Append("reoptimize-delay-after-affinity-failure-timer", types.YLeaf{"ReoptimizeDelayAfterAffinityFailureTimer", globalAttributes.ReoptimizeDelayAfterAffinityFailureTimer})
    globalAttributes.EntityData.Leafs.Append("log-frr-protection", types.YLeaf{"LogFrrProtection", globalAttributes.LogFrrProtection})
    globalAttributes.EntityData.Leafs.Append("reoptimize-timer-frequency", types.YLeaf{"ReoptimizeTimerFrequency", globalAttributes.ReoptimizeTimerFrequency})
    globalAttributes.EntityData.Leafs.Append("log-mid", types.YLeaf{"LogMid", globalAttributes.LogMid})
    globalAttributes.EntityData.Leafs.Append("log-preemption", types.YLeaf{"LogPreemption", globalAttributes.LogPreemption})

    globalAttributes.EntityData.YListKeys = []string {}

    return &(globalAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel
// Configure auto-tunnels feature
type MplsTe_GlobalAttributes_AutoTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure auto-tunnel PCC (Path Computation Client) feature.
    Pcc MplsTe_GlobalAttributes_AutoTunnel_Pcc

    // Configure P2P auto-tunnel feature.
    P2pAutoTunnel MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel

    // Configure auto-tunnel backup feature.
    Backup MplsTe_GlobalAttributes_AutoTunnel_Backup

    // Configure auto-tunnel mesh feature.
    Mesh MplsTe_GlobalAttributes_AutoTunnel_Mesh

    // Configure P2MP auto-tunnel feature.
    P2mpAutoTunnel MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel
}

func (autoTunnel *MplsTe_GlobalAttributes_AutoTunnel) GetEntityData() *types.CommonEntityData {
    autoTunnel.EntityData.YFilter = autoTunnel.YFilter
    autoTunnel.EntityData.YangName = "auto-tunnel"
    autoTunnel.EntityData.BundleName = "cisco_ios_xr"
    autoTunnel.EntityData.ParentYangName = "global-attributes"
    autoTunnel.EntityData.SegmentPath = "auto-tunnel"
    autoTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoTunnel.EntityData.Children = types.NewOrderedMap()
    autoTunnel.EntityData.Children.Append("pcc", types.YChild{"Pcc", &autoTunnel.Pcc})
    autoTunnel.EntityData.Children.Append("p2p-auto-tunnel", types.YChild{"P2pAutoTunnel", &autoTunnel.P2pAutoTunnel})
    autoTunnel.EntityData.Children.Append("backup", types.YChild{"Backup", &autoTunnel.Backup})
    autoTunnel.EntityData.Children.Append("mesh", types.YChild{"Mesh", &autoTunnel.Mesh})
    autoTunnel.EntityData.Children.Append("p2mp-auto-tunnel", types.YChild{"P2mpAutoTunnel", &autoTunnel.P2mpAutoTunnel})
    autoTunnel.EntityData.Leafs = types.NewOrderedMap()

    autoTunnel.EntityData.YListKeys = []string {}

    return &(autoTunnel.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Pcc
// Configure auto-tunnel PCC (Path Computation
// Client) feature
type MplsTe_GlobalAttributes_AutoTunnel_Pcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tunnel ID range for auto-tunnel features.
    TunnelRange MplsTe_GlobalAttributes_AutoTunnel_Pcc_TunnelRange
}

func (pcc *MplsTe_GlobalAttributes_AutoTunnel_Pcc) GetEntityData() *types.CommonEntityData {
    pcc.EntityData.YFilter = pcc.YFilter
    pcc.EntityData.YangName = "pcc"
    pcc.EntityData.BundleName = "cisco_ios_xr"
    pcc.EntityData.ParentYangName = "auto-tunnel"
    pcc.EntityData.SegmentPath = "pcc"
    pcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcc.EntityData.Children = types.NewOrderedMap()
    pcc.EntityData.Children.Append("tunnel-range", types.YChild{"TunnelRange", &pcc.TunnelRange})
    pcc.EntityData.Leafs = types.NewOrderedMap()

    pcc.EntityData.YListKeys = []string {}

    return &(pcc.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Pcc_TunnelRange
// Configure tunnel ID range for auto-tunnel
// features
// This type is a presence type.
type MplsTe_GlobalAttributes_AutoTunnel_Pcc_TunnelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Minimum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MinTunnelId interface{}

    // Maximum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MaxTunnelId interface{}
}

func (tunnelRange *MplsTe_GlobalAttributes_AutoTunnel_Pcc_TunnelRange) GetEntityData() *types.CommonEntityData {
    tunnelRange.EntityData.YFilter = tunnelRange.YFilter
    tunnelRange.EntityData.YangName = "tunnel-range"
    tunnelRange.EntityData.BundleName = "cisco_ios_xr"
    tunnelRange.EntityData.ParentYangName = "pcc"
    tunnelRange.EntityData.SegmentPath = "tunnel-range"
    tunnelRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelRange.EntityData.Children = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs.Append("min-tunnel-id", types.YLeaf{"MinTunnelId", tunnelRange.MinTunnelId})
    tunnelRange.EntityData.Leafs.Append("max-tunnel-id", types.YLeaf{"MaxTunnelId", tunnelRange.MaxTunnelId})

    tunnelRange.EntityData.YListKeys = []string {}

    return &(tunnelRange.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel
// Configure P2P auto-tunnel feature
type MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tunnel ID range for auto-tunnel features.
    TunnelRange MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel_TunnelRange
}

func (p2pAutoTunnel *MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel) GetEntityData() *types.CommonEntityData {
    p2pAutoTunnel.EntityData.YFilter = p2pAutoTunnel.YFilter
    p2pAutoTunnel.EntityData.YangName = "p2p-auto-tunnel"
    p2pAutoTunnel.EntityData.BundleName = "cisco_ios_xr"
    p2pAutoTunnel.EntityData.ParentYangName = "auto-tunnel"
    p2pAutoTunnel.EntityData.SegmentPath = "p2p-auto-tunnel"
    p2pAutoTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2pAutoTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2pAutoTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2pAutoTunnel.EntityData.Children = types.NewOrderedMap()
    p2pAutoTunnel.EntityData.Children.Append("tunnel-range", types.YChild{"TunnelRange", &p2pAutoTunnel.TunnelRange})
    p2pAutoTunnel.EntityData.Leafs = types.NewOrderedMap()

    p2pAutoTunnel.EntityData.YListKeys = []string {}

    return &(p2pAutoTunnel.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel_TunnelRange
// Configure tunnel ID range for auto-tunnel
// features
// This type is a presence type.
type MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel_TunnelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Minimum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MinTunnelId interface{}

    // Maximum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MaxTunnelId interface{}
}

func (tunnelRange *MplsTe_GlobalAttributes_AutoTunnel_P2pAutoTunnel_TunnelRange) GetEntityData() *types.CommonEntityData {
    tunnelRange.EntityData.YFilter = tunnelRange.YFilter
    tunnelRange.EntityData.YangName = "tunnel-range"
    tunnelRange.EntityData.BundleName = "cisco_ios_xr"
    tunnelRange.EntityData.ParentYangName = "p2p-auto-tunnel"
    tunnelRange.EntityData.SegmentPath = "tunnel-range"
    tunnelRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelRange.EntityData.Children = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs.Append("min-tunnel-id", types.YLeaf{"MinTunnelId", tunnelRange.MinTunnelId})
    tunnelRange.EntityData.Leafs.Append("max-tunnel-id", types.YLeaf{"MaxTunnelId", tunnelRange.MaxTunnelId})

    tunnelRange.EntityData.YListKeys = []string {}

    return &(tunnelRange.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Backup
// Configure auto-tunnel backup feature
type MplsTe_GlobalAttributes_AutoTunnel_Backup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ignore affinity during CSPF for auto backup tunnels. The type is
    // interface{}.
    AffinityIgnore interface{}

    // Configure auto-tunnel backup timers value.
    Timers MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers

    // Configure tunnel ID range for auto-tunnel features.
    TunnelRange MplsTe_GlobalAttributes_AutoTunnel_Backup_TunnelRange
}

func (backup *MplsTe_GlobalAttributes_AutoTunnel_Backup) GetEntityData() *types.CommonEntityData {
    backup.EntityData.YFilter = backup.YFilter
    backup.EntityData.YangName = "backup"
    backup.EntityData.BundleName = "cisco_ios_xr"
    backup.EntityData.ParentYangName = "auto-tunnel"
    backup.EntityData.SegmentPath = "backup"
    backup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backup.EntityData.Children = types.NewOrderedMap()
    backup.EntityData.Children.Append("timers", types.YChild{"Timers", &backup.Timers})
    backup.EntityData.Children.Append("tunnel-range", types.YChild{"TunnelRange", &backup.TunnelRange})
    backup.EntityData.Leafs = types.NewOrderedMap()
    backup.EntityData.Leafs.Append("affinity-ignore", types.YLeaf{"AffinityIgnore", backup.AffinityIgnore})

    backup.EntityData.YListKeys = []string {}

    return &(backup.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers
// Configure auto-tunnel backup timers value
type MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure auto-tunnel backup removal timers value.
    Removal MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers_Removal
}

func (timers *MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "backup"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("removal", types.YChild{"Removal", &timers.Removal})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers_Removal
// Configure auto-tunnel backup removal timers
// value
type MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers_Removal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto-tunnel backup unused timeout in minutes (0=never timeout). The type is
    // interface{} with range: 0..10080. Units are minute. The default value is
    // 3600.
    Unused interface{}
}

func (removal *MplsTe_GlobalAttributes_AutoTunnel_Backup_Timers_Removal) GetEntityData() *types.CommonEntityData {
    removal.EntityData.YFilter = removal.YFilter
    removal.EntityData.YangName = "removal"
    removal.EntityData.BundleName = "cisco_ios_xr"
    removal.EntityData.ParentYangName = "timers"
    removal.EntityData.SegmentPath = "removal"
    removal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    removal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    removal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    removal.EntityData.Children = types.NewOrderedMap()
    removal.EntityData.Leafs = types.NewOrderedMap()
    removal.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", removal.Unused})

    removal.EntityData.YListKeys = []string {}

    return &(removal.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Backup_TunnelRange
// Configure tunnel ID range for auto-tunnel
// features
// This type is a presence type.
type MplsTe_GlobalAttributes_AutoTunnel_Backup_TunnelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Minimum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MinTunnelId interface{}

    // Maximum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MaxTunnelId interface{}
}

func (tunnelRange *MplsTe_GlobalAttributes_AutoTunnel_Backup_TunnelRange) GetEntityData() *types.CommonEntityData {
    tunnelRange.EntityData.YFilter = tunnelRange.YFilter
    tunnelRange.EntityData.YangName = "tunnel-range"
    tunnelRange.EntityData.BundleName = "cisco_ios_xr"
    tunnelRange.EntityData.ParentYangName = "backup"
    tunnelRange.EntityData.SegmentPath = "tunnel-range"
    tunnelRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelRange.EntityData.Children = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs.Append("min-tunnel-id", types.YLeaf{"MinTunnelId", tunnelRange.MinTunnelId})
    tunnelRange.EntityData.Leafs.Append("max-tunnel-id", types.YLeaf{"MaxTunnelId", tunnelRange.MaxTunnelId})

    tunnelRange.EntityData.YListKeys = []string {}

    return &(tunnelRange.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Mesh
// Configure auto-tunnel mesh feature
type MplsTe_GlobalAttributes_AutoTunnel_Mesh struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure auto-tunnel mesh group.
    MeshGroups MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups

    // Configure auto-tunnel backup timers value.
    Timers MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers

    // Configure tunnel ID range for auto-tunnel features.
    TunnelRange MplsTe_GlobalAttributes_AutoTunnel_Mesh_TunnelRange
}

func (mesh *MplsTe_GlobalAttributes_AutoTunnel_Mesh) GetEntityData() *types.CommonEntityData {
    mesh.EntityData.YFilter = mesh.YFilter
    mesh.EntityData.YangName = "mesh"
    mesh.EntityData.BundleName = "cisco_ios_xr"
    mesh.EntityData.ParentYangName = "auto-tunnel"
    mesh.EntityData.SegmentPath = "mesh"
    mesh.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mesh.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mesh.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mesh.EntityData.Children = types.NewOrderedMap()
    mesh.EntityData.Children.Append("mesh-groups", types.YChild{"MeshGroups", &mesh.MeshGroups})
    mesh.EntityData.Children.Append("timers", types.YChild{"Timers", &mesh.Timers})
    mesh.EntityData.Children.Append("tunnel-range", types.YChild{"TunnelRange", &mesh.TunnelRange})
    mesh.EntityData.Leafs = types.NewOrderedMap()

    mesh.EntityData.YListKeys = []string {}

    return &(mesh.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups
// Configure auto-tunnel mesh group
type MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto-mesh group identifier. The type is slice of
    // MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups_MeshGroup.
    MeshGroup []*MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups_MeshGroup
}

func (meshGroups *MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups) GetEntityData() *types.CommonEntityData {
    meshGroups.EntityData.YFilter = meshGroups.YFilter
    meshGroups.EntityData.YangName = "mesh-groups"
    meshGroups.EntityData.BundleName = "cisco_ios_xr"
    meshGroups.EntityData.ParentYangName = "mesh"
    meshGroups.EntityData.SegmentPath = "mesh-groups"
    meshGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    meshGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    meshGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    meshGroups.EntityData.Children = types.NewOrderedMap()
    meshGroups.EntityData.Children.Append("mesh-group", types.YChild{"MeshGroup", nil})
    for i := range meshGroups.MeshGroup {
        meshGroups.EntityData.Children.Append(types.GetSegmentPath(meshGroups.MeshGroup[i]), types.YChild{"MeshGroup", meshGroups.MeshGroup[i]})
    }
    meshGroups.EntityData.Leafs = types.NewOrderedMap()

    meshGroups.EntityData.YListKeys = []string {}

    return &(meshGroups.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups_MeshGroup
// Auto-mesh group identifier
type MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups_MeshGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Mesh group ID. The type is interface{} with range:
    // 0..4294967295.
    MeshGroupId interface{}

    // The name of prefix-list to be applied to this destination-list. The type is
    // string with length: 1..32.
    DestinationList interface{}

    // Disables mesh group. The type is interface{}.
    Disable interface{}

    // The name of auto-mesh attribute set to be applied to this group. The type
    // is string with length: 1..64.
    AttributeSet interface{}

    // Auto-mesh group enable object that controls whether this group is
    // configured or not .This object must be set before other configuration
    // supplied for this group. The type is interface{}.
    Create interface{}

    // Automatically create tunnel to all next-hops. The type is interface{}.
    OneHop interface{}
}

func (meshGroup *MplsTe_GlobalAttributes_AutoTunnel_Mesh_MeshGroups_MeshGroup) GetEntityData() *types.CommonEntityData {
    meshGroup.EntityData.YFilter = meshGroup.YFilter
    meshGroup.EntityData.YangName = "mesh-group"
    meshGroup.EntityData.BundleName = "cisco_ios_xr"
    meshGroup.EntityData.ParentYangName = "mesh-groups"
    meshGroup.EntityData.SegmentPath = "mesh-group" + types.AddKeyToken(meshGroup.MeshGroupId, "mesh-group-id")
    meshGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    meshGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    meshGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    meshGroup.EntityData.Children = types.NewOrderedMap()
    meshGroup.EntityData.Leafs = types.NewOrderedMap()
    meshGroup.EntityData.Leafs.Append("mesh-group-id", types.YLeaf{"MeshGroupId", meshGroup.MeshGroupId})
    meshGroup.EntityData.Leafs.Append("destination-list", types.YLeaf{"DestinationList", meshGroup.DestinationList})
    meshGroup.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", meshGroup.Disable})
    meshGroup.EntityData.Leafs.Append("attribute-set", types.YLeaf{"AttributeSet", meshGroup.AttributeSet})
    meshGroup.EntityData.Leafs.Append("create", types.YLeaf{"Create", meshGroup.Create})
    meshGroup.EntityData.Leafs.Append("one-hop", types.YLeaf{"OneHop", meshGroup.OneHop})

    meshGroup.EntityData.YListKeys = []string {"MeshGroupId"}

    return &(meshGroup.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers
// Configure auto-tunnel backup timers value
type MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure auto-tunnel backup removal timers value.
    Removal MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers_Removal
}

func (timers *MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "mesh"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("removal", types.YChild{"Removal", &timers.Removal})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers_Removal
// Configure auto-tunnel backup removal timers
// value
type MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers_Removal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto-tunnel backup unused timeout in minutes (0=never timeout). The type is
    // interface{} with range: 0..10080. Units are minute. The default value is
    // 3600.
    Unused interface{}
}

func (removal *MplsTe_GlobalAttributes_AutoTunnel_Mesh_Timers_Removal) GetEntityData() *types.CommonEntityData {
    removal.EntityData.YFilter = removal.YFilter
    removal.EntityData.YangName = "removal"
    removal.EntityData.BundleName = "cisco_ios_xr"
    removal.EntityData.ParentYangName = "timers"
    removal.EntityData.SegmentPath = "removal"
    removal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    removal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    removal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    removal.EntityData.Children = types.NewOrderedMap()
    removal.EntityData.Leafs = types.NewOrderedMap()
    removal.EntityData.Leafs.Append("unused", types.YLeaf{"Unused", removal.Unused})

    removal.EntityData.YListKeys = []string {}

    return &(removal.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_Mesh_TunnelRange
// Configure tunnel ID range for auto-tunnel
// features
// This type is a presence type.
type MplsTe_GlobalAttributes_AutoTunnel_Mesh_TunnelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Minimum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MinTunnelId interface{}

    // Maximum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MaxTunnelId interface{}
}

func (tunnelRange *MplsTe_GlobalAttributes_AutoTunnel_Mesh_TunnelRange) GetEntityData() *types.CommonEntityData {
    tunnelRange.EntityData.YFilter = tunnelRange.YFilter
    tunnelRange.EntityData.YangName = "tunnel-range"
    tunnelRange.EntityData.BundleName = "cisco_ios_xr"
    tunnelRange.EntityData.ParentYangName = "mesh"
    tunnelRange.EntityData.SegmentPath = "tunnel-range"
    tunnelRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelRange.EntityData.Children = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs.Append("min-tunnel-id", types.YLeaf{"MinTunnelId", tunnelRange.MinTunnelId})
    tunnelRange.EntityData.Leafs.Append("max-tunnel-id", types.YLeaf{"MaxTunnelId", tunnelRange.MaxTunnelId})

    tunnelRange.EntityData.YListKeys = []string {}

    return &(tunnelRange.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel
// Configure P2MP auto-tunnel feature
type MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure tunnel ID range for auto-tunnel features.
    TunnelRange MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel_TunnelRange
}

func (p2mpAutoTunnel *MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel) GetEntityData() *types.CommonEntityData {
    p2mpAutoTunnel.EntityData.YFilter = p2mpAutoTunnel.YFilter
    p2mpAutoTunnel.EntityData.YangName = "p2mp-auto-tunnel"
    p2mpAutoTunnel.EntityData.BundleName = "cisco_ios_xr"
    p2mpAutoTunnel.EntityData.ParentYangName = "auto-tunnel"
    p2mpAutoTunnel.EntityData.SegmentPath = "p2mp-auto-tunnel"
    p2mpAutoTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2mpAutoTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2mpAutoTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2mpAutoTunnel.EntityData.Children = types.NewOrderedMap()
    p2mpAutoTunnel.EntityData.Children.Append("tunnel-range", types.YChild{"TunnelRange", &p2mpAutoTunnel.TunnelRange})
    p2mpAutoTunnel.EntityData.Leafs = types.NewOrderedMap()

    p2mpAutoTunnel.EntityData.YListKeys = []string {}

    return &(p2mpAutoTunnel.EntityData)
}

// MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel_TunnelRange
// Configure tunnel ID range for auto-tunnel
// features
// This type is a presence type.
type MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel_TunnelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Minimum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MinTunnelId interface{}

    // Maximum tunnel ID for auto-tunnels. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    MaxTunnelId interface{}
}

func (tunnelRange *MplsTe_GlobalAttributes_AutoTunnel_P2mpAutoTunnel_TunnelRange) GetEntityData() *types.CommonEntityData {
    tunnelRange.EntityData.YFilter = tunnelRange.YFilter
    tunnelRange.EntityData.YangName = "tunnel-range"
    tunnelRange.EntityData.BundleName = "cisco_ios_xr"
    tunnelRange.EntityData.ParentYangName = "p2mp-auto-tunnel"
    tunnelRange.EntityData.SegmentPath = "tunnel-range"
    tunnelRange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelRange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelRange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelRange.EntityData.Children = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs = types.NewOrderedMap()
    tunnelRange.EntityData.Leafs.Append("min-tunnel-id", types.YLeaf{"MinTunnelId", tunnelRange.MinTunnelId})
    tunnelRange.EntityData.Leafs.Append("max-tunnel-id", types.YLeaf{"MaxTunnelId", tunnelRange.MaxTunnelId})

    tunnelRange.EntityData.YListKeys = []string {}

    return &(tunnelRange.EntityData)
}

// MplsTe_GlobalAttributes_HardwareOutOfResource
// Configure HW OOR processing in MPLS-TE
type MplsTe_GlobalAttributes_HardwareOutOfResource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for HW OOR Red State.
    OorRedState MplsTe_GlobalAttributes_HardwareOutOfResource_OorRedState

    // Configuration for HW OOR Yellow State.
    OorYellowState MplsTe_GlobalAttributes_HardwareOutOfResource_OorYellowState

    // Configuration for HW OOR Green State.
    OorGreenState MplsTe_GlobalAttributes_HardwareOutOfResource_OorGreenState
}

func (hardwareOutOfResource *MplsTe_GlobalAttributes_HardwareOutOfResource) GetEntityData() *types.CommonEntityData {
    hardwareOutOfResource.EntityData.YFilter = hardwareOutOfResource.YFilter
    hardwareOutOfResource.EntityData.YangName = "hardware-out-of-resource"
    hardwareOutOfResource.EntityData.BundleName = "cisco_ios_xr"
    hardwareOutOfResource.EntityData.ParentYangName = "global-attributes"
    hardwareOutOfResource.EntityData.SegmentPath = "hardware-out-of-resource"
    hardwareOutOfResource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hardwareOutOfResource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hardwareOutOfResource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hardwareOutOfResource.EntityData.Children = types.NewOrderedMap()
    hardwareOutOfResource.EntityData.Children.Append("oor-red-state", types.YChild{"OorRedState", &hardwareOutOfResource.OorRedState})
    hardwareOutOfResource.EntityData.Children.Append("oor-yellow-state", types.YChild{"OorYellowState", &hardwareOutOfResource.OorYellowState})
    hardwareOutOfResource.EntityData.Children.Append("oor-green-state", types.YChild{"OorGreenState", &hardwareOutOfResource.OorGreenState})
    hardwareOutOfResource.EntityData.Leafs = types.NewOrderedMap()

    hardwareOutOfResource.EntityData.YListKeys = []string {}

    return &(hardwareOutOfResource.EntityData)
}

// MplsTe_GlobalAttributes_HardwareOutOfResource_OorRedState
// Configuration for HW OOR Red State
type MplsTe_GlobalAttributes_HardwareOutOfResource_OorRedState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable FRR node-protection when the link is in this OOR State. The type is
    // interface{}.
    OorNodeProtectionDisable interface{}

    // Flood a specific percentage of the available bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage. The default value is
    // 100.
    OorAvailableBandwidthPercentage interface{}

    // Only accept LSPs with at least the specified bandwidth (in kbps). The type
    // is interface{} with range: 0..4294967295. Units are kbit/s. The default
    // value is 0.
    OorAcceptLspMinBandwidth interface{}

    // Allow the setup of reoptimized LSPs over the link in this OOR State. The
    // type is interface{}.
    OorAcceptReoptLsp interface{}

    // Penalty applied to the TE metric of a link in OOR state. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    OorMetricTePenalty interface{}
}

func (oorRedState *MplsTe_GlobalAttributes_HardwareOutOfResource_OorRedState) GetEntityData() *types.CommonEntityData {
    oorRedState.EntityData.YFilter = oorRedState.YFilter
    oorRedState.EntityData.YangName = "oor-red-state"
    oorRedState.EntityData.BundleName = "cisco_ios_xr"
    oorRedState.EntityData.ParentYangName = "hardware-out-of-resource"
    oorRedState.EntityData.SegmentPath = "oor-red-state"
    oorRedState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorRedState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorRedState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorRedState.EntityData.Children = types.NewOrderedMap()
    oorRedState.EntityData.Leafs = types.NewOrderedMap()
    oorRedState.EntityData.Leafs.Append("oor-node-protection-disable", types.YLeaf{"OorNodeProtectionDisable", oorRedState.OorNodeProtectionDisable})
    oorRedState.EntityData.Leafs.Append("oor-available-bandwidth-percentage", types.YLeaf{"OorAvailableBandwidthPercentage", oorRedState.OorAvailableBandwidthPercentage})
    oorRedState.EntityData.Leafs.Append("oor-accept-lsp-min-bandwidth", types.YLeaf{"OorAcceptLspMinBandwidth", oorRedState.OorAcceptLspMinBandwidth})
    oorRedState.EntityData.Leafs.Append("oor-accept-reopt-lsp", types.YLeaf{"OorAcceptReoptLsp", oorRedState.OorAcceptReoptLsp})
    oorRedState.EntityData.Leafs.Append("oor-metric-te-penalty", types.YLeaf{"OorMetricTePenalty", oorRedState.OorMetricTePenalty})

    oorRedState.EntityData.YListKeys = []string {}

    return &(oorRedState.EntityData)
}

// MplsTe_GlobalAttributes_HardwareOutOfResource_OorYellowState
// Configuration for HW OOR Yellow State
type MplsTe_GlobalAttributes_HardwareOutOfResource_OorYellowState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable FRR node-protection when the link is in this OOR State. The type is
    // interface{}.
    OorNodeProtectionDisable interface{}

    // Flood a specific percentage of the available bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage. The default value is
    // 100.
    OorAvailableBandwidthPercentage interface{}

    // Only accept LSPs with at least the specified bandwidth (in kbps). The type
    // is interface{} with range: 0..4294967295. Units are kbit/s. The default
    // value is 0.
    OorAcceptLspMinBandwidth interface{}

    // Allow the setup of reoptimized LSPs over the link in this OOR State. The
    // type is interface{}.
    OorAcceptReoptLsp interface{}

    // Penalty applied to the TE metric of a link in OOR state. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    OorMetricTePenalty interface{}
}

func (oorYellowState *MplsTe_GlobalAttributes_HardwareOutOfResource_OorYellowState) GetEntityData() *types.CommonEntityData {
    oorYellowState.EntityData.YFilter = oorYellowState.YFilter
    oorYellowState.EntityData.YangName = "oor-yellow-state"
    oorYellowState.EntityData.BundleName = "cisco_ios_xr"
    oorYellowState.EntityData.ParentYangName = "hardware-out-of-resource"
    oorYellowState.EntityData.SegmentPath = "oor-yellow-state"
    oorYellowState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorYellowState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorYellowState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorYellowState.EntityData.Children = types.NewOrderedMap()
    oorYellowState.EntityData.Leafs = types.NewOrderedMap()
    oorYellowState.EntityData.Leafs.Append("oor-node-protection-disable", types.YLeaf{"OorNodeProtectionDisable", oorYellowState.OorNodeProtectionDisable})
    oorYellowState.EntityData.Leafs.Append("oor-available-bandwidth-percentage", types.YLeaf{"OorAvailableBandwidthPercentage", oorYellowState.OorAvailableBandwidthPercentage})
    oorYellowState.EntityData.Leafs.Append("oor-accept-lsp-min-bandwidth", types.YLeaf{"OorAcceptLspMinBandwidth", oorYellowState.OorAcceptLspMinBandwidth})
    oorYellowState.EntityData.Leafs.Append("oor-accept-reopt-lsp", types.YLeaf{"OorAcceptReoptLsp", oorYellowState.OorAcceptReoptLsp})
    oorYellowState.EntityData.Leafs.Append("oor-metric-te-penalty", types.YLeaf{"OorMetricTePenalty", oorYellowState.OorMetricTePenalty})

    oorYellowState.EntityData.YListKeys = []string {}

    return &(oorYellowState.EntityData)
}

// MplsTe_GlobalAttributes_HardwareOutOfResource_OorGreenState
// Configuration for HW OOR Green State
type MplsTe_GlobalAttributes_HardwareOutOfResource_OorGreenState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Period of time (minutes) during which the action in Green state are
    // applied. After this period, the processing in TE goes back to normal state.
    // The type is interface{} with range: 0..10080. Units are minute. The default
    // value is 0.
    OorRecoveryDuration interface{}

    // Disable FRR node-protection when the link is in this OOR State. The type is
    // interface{}.
    OorNodeProtectionDisable interface{}

    // Flood a specific percentage of the available bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage. The default value is
    // 100.
    OorAvailableBandwidthPercentage interface{}

    // Only accept LSPs with at least the specified bandwidth (in kbps). The type
    // is interface{} with range: 0..4294967295. Units are kbit/s. The default
    // value is 0.
    OorAcceptLspMinBandwidth interface{}

    // Allow the setup of reoptimized LSPs over the link in this OOR State. The
    // type is interface{}.
    OorAcceptReoptLsp interface{}

    // Penalty applied to the TE metric of a link in OOR state. The type is
    // interface{} with range: 0..4294967295. The default value is 0.
    OorMetricTePenalty interface{}
}

func (oorGreenState *MplsTe_GlobalAttributes_HardwareOutOfResource_OorGreenState) GetEntityData() *types.CommonEntityData {
    oorGreenState.EntityData.YFilter = oorGreenState.YFilter
    oorGreenState.EntityData.YangName = "oor-green-state"
    oorGreenState.EntityData.BundleName = "cisco_ios_xr"
    oorGreenState.EntityData.ParentYangName = "hardware-out-of-resource"
    oorGreenState.EntityData.SegmentPath = "oor-green-state"
    oorGreenState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oorGreenState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oorGreenState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oorGreenState.EntityData.Children = types.NewOrderedMap()
    oorGreenState.EntityData.Leafs = types.NewOrderedMap()
    oorGreenState.EntityData.Leafs.Append("oor-recovery-duration", types.YLeaf{"OorRecoveryDuration", oorGreenState.OorRecoveryDuration})
    oorGreenState.EntityData.Leafs.Append("oor-node-protection-disable", types.YLeaf{"OorNodeProtectionDisable", oorGreenState.OorNodeProtectionDisable})
    oorGreenState.EntityData.Leafs.Append("oor-available-bandwidth-percentage", types.YLeaf{"OorAvailableBandwidthPercentage", oorGreenState.OorAvailableBandwidthPercentage})
    oorGreenState.EntityData.Leafs.Append("oor-accept-lsp-min-bandwidth", types.YLeaf{"OorAcceptLspMinBandwidth", oorGreenState.OorAcceptLspMinBandwidth})
    oorGreenState.EntityData.Leafs.Append("oor-accept-reopt-lsp", types.YLeaf{"OorAcceptReoptLsp", oorGreenState.OorAcceptReoptLsp})
    oorGreenState.EntityData.Leafs.Append("oor-metric-te-penalty", types.YLeaf{"OorMetricTePenalty", oorGreenState.OorMetricTePenalty})

    oorGreenState.EntityData.YListKeys = []string {}

    return &(oorGreenState.EntityData)
}

// MplsTe_GlobalAttributes_SecondaryRouterIds
// Configure MPLS TE Secondary Router ID
type MplsTe_GlobalAttributes_SecondaryRouterIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Secondary Router ID. The type is slice of
    // MplsTe_GlobalAttributes_SecondaryRouterIds_SecondaryRouterId.
    SecondaryRouterId []*MplsTe_GlobalAttributes_SecondaryRouterIds_SecondaryRouterId
}

func (secondaryRouterIds *MplsTe_GlobalAttributes_SecondaryRouterIds) GetEntityData() *types.CommonEntityData {
    secondaryRouterIds.EntityData.YFilter = secondaryRouterIds.YFilter
    secondaryRouterIds.EntityData.YangName = "secondary-router-ids"
    secondaryRouterIds.EntityData.BundleName = "cisco_ios_xr"
    secondaryRouterIds.EntityData.ParentYangName = "global-attributes"
    secondaryRouterIds.EntityData.SegmentPath = "secondary-router-ids"
    secondaryRouterIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryRouterIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryRouterIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryRouterIds.EntityData.Children = types.NewOrderedMap()
    secondaryRouterIds.EntityData.Children.Append("secondary-router-id", types.YChild{"SecondaryRouterId", nil})
    for i := range secondaryRouterIds.SecondaryRouterId {
        secondaryRouterIds.EntityData.Children.Append(types.GetSegmentPath(secondaryRouterIds.SecondaryRouterId[i]), types.YChild{"SecondaryRouterId", secondaryRouterIds.SecondaryRouterId[i]})
    }
    secondaryRouterIds.EntityData.Leafs = types.NewOrderedMap()

    secondaryRouterIds.EntityData.YListKeys = []string {}

    return &(secondaryRouterIds.EntityData)
}

// MplsTe_GlobalAttributes_SecondaryRouterIds_SecondaryRouterId
// Secondary Router ID
type MplsTe_GlobalAttributes_SecondaryRouterIds_SecondaryRouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Secondary TE Router ID. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SecondaryRouterIdValue interface{}
}

func (secondaryRouterId *MplsTe_GlobalAttributes_SecondaryRouterIds_SecondaryRouterId) GetEntityData() *types.CommonEntityData {
    secondaryRouterId.EntityData.YFilter = secondaryRouterId.YFilter
    secondaryRouterId.EntityData.YangName = "secondary-router-id"
    secondaryRouterId.EntityData.BundleName = "cisco_ios_xr"
    secondaryRouterId.EntityData.ParentYangName = "secondary-router-ids"
    secondaryRouterId.EntityData.SegmentPath = "secondary-router-id" + types.AddKeyToken(secondaryRouterId.SecondaryRouterIdValue, "secondary-router-id-value")
    secondaryRouterId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryRouterId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryRouterId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryRouterId.EntityData.Children = types.NewOrderedMap()
    secondaryRouterId.EntityData.Leafs = types.NewOrderedMap()
    secondaryRouterId.EntityData.Leafs.Append("secondary-router-id-value", types.YLeaf{"SecondaryRouterIdValue", secondaryRouterId.SecondaryRouterIdValue})

    secondaryRouterId.EntityData.YListKeys = []string {"SecondaryRouterIdValue"}

    return &(secondaryRouterId.EntityData)
}

// MplsTe_GlobalAttributes_Srlg
// Configure SRLG values and MPLS-TE properties
type MplsTe_GlobalAttributes_Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default Admin weight any SRLG value that does not have one. The type is
    // interface{} with range: 0..4294967295. The default value is 1.
    DefaultAdminWeight interface{}

    // Enter SRLG property configuration. The type is interface{}.
    Enable interface{}

    // Configure SRLG identified by names.
    Names MplsTe_GlobalAttributes_Srlg_Names
}

func (srlg *MplsTe_GlobalAttributes_Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "cisco_ios_xr"
    srlg.EntityData.ParentYangName = "global-attributes"
    srlg.EntityData.SegmentPath = "srlg"
    srlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Children.Append("names", types.YChild{"Names", &srlg.Names})
    srlg.EntityData.Leafs = types.NewOrderedMap()
    srlg.EntityData.Leafs.Append("default-admin-weight", types.YLeaf{"DefaultAdminWeight", srlg.DefaultAdminWeight})
    srlg.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", srlg.Enable})

    srlg.EntityData.YListKeys = []string {}

    return &(srlg.EntityData)
}

// MplsTe_GlobalAttributes_Srlg_Names
// Configure SRLG identified by names
type MplsTe_GlobalAttributes_Srlg_Names struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG name and its MPLS-TE properties. The type is slice of
    // MplsTe_GlobalAttributes_Srlg_Names_Name.
    Name []*MplsTe_GlobalAttributes_Srlg_Names_Name
}

func (names *MplsTe_GlobalAttributes_Srlg_Names) GetEntityData() *types.CommonEntityData {
    names.EntityData.YFilter = names.YFilter
    names.EntityData.YangName = "names"
    names.EntityData.BundleName = "cisco_ios_xr"
    names.EntityData.ParentYangName = "srlg"
    names.EntityData.SegmentPath = "names"
    names.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    names.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    names.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    names.EntityData.Children = types.NewOrderedMap()
    names.EntityData.Children.Append("name", types.YChild{"Name", nil})
    for i := range names.Name {
        names.EntityData.Children.Append(types.GetSegmentPath(names.Name[i]), types.YChild{"Name", names.Name[i]})
    }
    names.EntityData.Leafs = types.NewOrderedMap()

    names.EntityData.YListKeys = []string {}

    return &(names.EntityData)
}

// MplsTe_GlobalAttributes_Srlg_Names_Name
// SRLG name and its MPLS-TE properties
type MplsTe_GlobalAttributes_Srlg_Names_Name struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG membership name. The type is string with
    // length: 1..64.
    SrlgName interface{}

    // Administrative weight for the SRLG value. The type is interface{} with
    // range: 0..4294967295.
    AdminWeight interface{}

    // Configure static SRLG members list.
    StaticSrlgMembers MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers
}

func (name *MplsTe_GlobalAttributes_Srlg_Names_Name) GetEntityData() *types.CommonEntityData {
    name.EntityData.YFilter = name.YFilter
    name.EntityData.YangName = "name"
    name.EntityData.BundleName = "cisco_ios_xr"
    name.EntityData.ParentYangName = "names"
    name.EntityData.SegmentPath = "name" + types.AddKeyToken(name.SrlgName, "srlg-name")
    name.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    name.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    name.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    name.EntityData.Children = types.NewOrderedMap()
    name.EntityData.Children.Append("static-srlg-members", types.YChild{"StaticSrlgMembers", &name.StaticSrlgMembers})
    name.EntityData.Leafs = types.NewOrderedMap()
    name.EntityData.Leafs.Append("srlg-name", types.YLeaf{"SrlgName", name.SrlgName})
    name.EntityData.Leafs.Append("admin-weight", types.YLeaf{"AdminWeight", name.AdminWeight})

    name.EntityData.YListKeys = []string {"SrlgName"}

    return &(name.EntityData)
}

// MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers
// Configure static SRLG members list
type MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A mapping of the local static SRLG member. The type is slice of
    // MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers_StaticSrlgMember.
    StaticSrlgMember []*MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers_StaticSrlgMember
}

func (staticSrlgMembers *MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers) GetEntityData() *types.CommonEntityData {
    staticSrlgMembers.EntityData.YFilter = staticSrlgMembers.YFilter
    staticSrlgMembers.EntityData.YangName = "static-srlg-members"
    staticSrlgMembers.EntityData.BundleName = "cisco_ios_xr"
    staticSrlgMembers.EntityData.ParentYangName = "name"
    staticSrlgMembers.EntityData.SegmentPath = "static-srlg-members"
    staticSrlgMembers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticSrlgMembers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticSrlgMembers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticSrlgMembers.EntityData.Children = types.NewOrderedMap()
    staticSrlgMembers.EntityData.Children.Append("static-srlg-member", types.YChild{"StaticSrlgMember", nil})
    for i := range staticSrlgMembers.StaticSrlgMember {
        staticSrlgMembers.EntityData.Children.Append(types.GetSegmentPath(staticSrlgMembers.StaticSrlgMember[i]), types.YChild{"StaticSrlgMember", staticSrlgMembers.StaticSrlgMember[i]})
    }
    staticSrlgMembers.EntityData.Leafs = types.NewOrderedMap()

    staticSrlgMembers.EntityData.YListKeys = []string {}

    return &(staticSrlgMembers.EntityData)
}

// MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers_StaticSrlgMember
// A mapping of the local static SRLG member
type MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers_StaticSrlgMember struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. From address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FromAddress interface{}

    // To Addres. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    ToAddress interface{}
}

func (staticSrlgMember *MplsTe_GlobalAttributes_Srlg_Names_Name_StaticSrlgMembers_StaticSrlgMember) GetEntityData() *types.CommonEntityData {
    staticSrlgMember.EntityData.YFilter = staticSrlgMember.YFilter
    staticSrlgMember.EntityData.YangName = "static-srlg-member"
    staticSrlgMember.EntityData.BundleName = "cisco_ios_xr"
    staticSrlgMember.EntityData.ParentYangName = "static-srlg-members"
    staticSrlgMember.EntityData.SegmentPath = "static-srlg-member" + types.AddKeyToken(staticSrlgMember.FromAddress, "from-address")
    staticSrlgMember.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticSrlgMember.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticSrlgMember.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticSrlgMember.EntityData.Children = types.NewOrderedMap()
    staticSrlgMember.EntityData.Leafs = types.NewOrderedMap()
    staticSrlgMember.EntityData.Leafs.Append("from-address", types.YLeaf{"FromAddress", staticSrlgMember.FromAddress})
    staticSrlgMember.EntityData.Leafs.Append("to-address", types.YLeaf{"ToAddress", staticSrlgMember.ToAddress})

    staticSrlgMember.EntityData.YListKeys = []string {"FromAddress"}

    return &(staticSrlgMember.EntityData)
}

// MplsTe_GlobalAttributes_Queues
// Configure MPLS TE route priority
type MplsTe_GlobalAttributes_Queues struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure route priority queue value. The type is slice of
    // MplsTe_GlobalAttributes_Queues_Queue.
    Queue []*MplsTe_GlobalAttributes_Queues_Queue
}

func (queues *MplsTe_GlobalAttributes_Queues) GetEntityData() *types.CommonEntityData {
    queues.EntityData.YFilter = queues.YFilter
    queues.EntityData.YangName = "queues"
    queues.EntityData.BundleName = "cisco_ios_xr"
    queues.EntityData.ParentYangName = "global-attributes"
    queues.EntityData.SegmentPath = "queues"
    queues.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queues.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queues.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queues.EntityData.Children = types.NewOrderedMap()
    queues.EntityData.Children.Append("queue", types.YChild{"Queue", nil})
    for i := range queues.Queue {
        queues.EntityData.Children.Append(types.GetSegmentPath(queues.Queue[i]), types.YChild{"Queue", queues.Queue[i]})
    }
    queues.EntityData.Leafs = types.NewOrderedMap()

    queues.EntityData.YListKeys = []string {}

    return &(queues.EntityData)
}

// MplsTe_GlobalAttributes_Queues_Queue
// Configure route priority queue value
type MplsTe_GlobalAttributes_Queues_Queue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Route Priority Tunnel Role. The type is
    // RoutePriorityRole.
    Role interface{}

    // Route priority queue value. The type is interface{} with range: 0..12. This
    // attribute is mandatory.
    Value interface{}
}

func (queue *MplsTe_GlobalAttributes_Queues_Queue) GetEntityData() *types.CommonEntityData {
    queue.EntityData.YFilter = queue.YFilter
    queue.EntityData.YangName = "queue"
    queue.EntityData.BundleName = "cisco_ios_xr"
    queue.EntityData.ParentYangName = "queues"
    queue.EntityData.SegmentPath = "queue" + types.AddKeyToken(queue.Role, "role")
    queue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    queue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    queue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    queue.EntityData.Children = types.NewOrderedMap()
    queue.EntityData.Leafs = types.NewOrderedMap()
    queue.EntityData.Leafs.Append("role", types.YLeaf{"Role", queue.Role})
    queue.EntityData.Leafs.Append("value", types.YLeaf{"Value", queue.Value})

    queue.EntityData.YListKeys = []string {"Role"}

    return &(queue.EntityData)
}

// MplsTe_GlobalAttributes_Mib
// MPLS-TE MIB properties
type MplsTe_GlobalAttributes_Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disables mib midpoint LSP traffic stats collection. The type is
    // interface{}.
    MidpointLspStatsCollectionDisable interface{}
}

func (mib *MplsTe_GlobalAttributes_Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "global-attributes"
    mib.EntityData.SegmentPath = "mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Leafs = types.NewOrderedMap()
    mib.EntityData.Leafs.Append("midpoint-lsp-stats-collection-disable", types.YLeaf{"MidpointLspStatsCollectionDisable", mib.MidpointLspStatsCollectionDisable})

    mib.EntityData.YListKeys = []string {}

    return &(mib.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet
// Attribute AttributeSets
type MplsTe_GlobalAttributes_AttributeSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Option Attribute-Set Table.
    PathOptionAttributes MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes

    // P2MP-TE Tunnel AttributeSets Table.
    P2mpteAttributes MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes

    // P2P-TE Tunnel AttributeSets Table.
    P2pTeAttributes MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes

    // Auto-backup Tunnel Attribute Table.
    AutoBackupAttributes MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes

    // OTN Path Protection Attributes table.
    OtnPpAttributes MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes

    // Auto-mesh Tunnel AttributeSets Table.
    AutoMeshAttributes MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes

    // XRO Tunnel Attributes table.
    XroAttributes MplsTe_GlobalAttributes_AttributeSet_XroAttributes
}

func (attributeSet *MplsTe_GlobalAttributes_AttributeSet) GetEntityData() *types.CommonEntityData {
    attributeSet.EntityData.YFilter = attributeSet.YFilter
    attributeSet.EntityData.YangName = "attribute-set"
    attributeSet.EntityData.BundleName = "cisco_ios_xr"
    attributeSet.EntityData.ParentYangName = "global-attributes"
    attributeSet.EntityData.SegmentPath = "attribute-set"
    attributeSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeSet.EntityData.Children = types.NewOrderedMap()
    attributeSet.EntityData.Children.Append("path-option-attributes", types.YChild{"PathOptionAttributes", &attributeSet.PathOptionAttributes})
    attributeSet.EntityData.Children.Append("p2mpte-attributes", types.YChild{"P2mpteAttributes", &attributeSet.P2mpteAttributes})
    attributeSet.EntityData.Children.Append("p2p-te-attributes", types.YChild{"P2pTeAttributes", &attributeSet.P2pTeAttributes})
    attributeSet.EntityData.Children.Append("auto-backup-attributes", types.YChild{"AutoBackupAttributes", &attributeSet.AutoBackupAttributes})
    attributeSet.EntityData.Children.Append("otn-pp-attributes", types.YChild{"OtnPpAttributes", &attributeSet.OtnPpAttributes})
    attributeSet.EntityData.Children.Append("auto-mesh-attributes", types.YChild{"AutoMeshAttributes", &attributeSet.AutoMeshAttributes})
    attributeSet.EntityData.Children.Append("xro-attributes", types.YChild{"XroAttributes", &attributeSet.XroAttributes})
    attributeSet.EntityData.Leafs = types.NewOrderedMap()

    attributeSet.EntityData.YListKeys = []string {}

    return &(attributeSet.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes
// Path Option Attribute-Set Table
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Option Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute.
    PathOptionAttribute []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute
}

func (pathOptionAttributes *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes) GetEntityData() *types.CommonEntityData {
    pathOptionAttributes.EntityData.YFilter = pathOptionAttributes.YFilter
    pathOptionAttributes.EntityData.YangName = "path-option-attributes"
    pathOptionAttributes.EntityData.BundleName = "cisco_ios_xr"
    pathOptionAttributes.EntityData.ParentYangName = "attribute-set"
    pathOptionAttributes.EntityData.SegmentPath = "path-option-attributes"
    pathOptionAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOptionAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOptionAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOptionAttributes.EntityData.Children = types.NewOrderedMap()
    pathOptionAttributes.EntityData.Children.Append("path-option-attribute", types.YChild{"PathOptionAttribute", nil})
    for i := range pathOptionAttributes.PathOptionAttribute {
        pathOptionAttributes.EntityData.Children.Append(types.GetSegmentPath(pathOptionAttributes.PathOptionAttribute[i]), types.YChild{"PathOptionAttribute", pathOptionAttributes.PathOptionAttribute[i]})
    }
    pathOptionAttributes.EntityData.Leafs = types.NewOrderedMap()

    pathOptionAttributes.EntityData.YListKeys = []string {}

    return &(pathOptionAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute
// Path Option Attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Configure BFD reverse path.
    BfdReversePath MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_BfdReversePath

    // Configure path selection properties.
    AttPathOptionPathSelection MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection

    // Configure pce properties.
    Pce MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce

    // Set the affinity flags and mask.
    AffinityMask MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AffinityMask

    // Tunnel bandwidth requirement.
    Bandwidth MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Bandwidth

    // Tunnel new style affinity attributes table.
    NewStyleAffinityAffinityTypes MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes
}

func (pathOptionAttribute *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute) GetEntityData() *types.CommonEntityData {
    pathOptionAttribute.EntityData.YFilter = pathOptionAttribute.YFilter
    pathOptionAttribute.EntityData.YangName = "path-option-attribute"
    pathOptionAttribute.EntityData.BundleName = "cisco_ios_xr"
    pathOptionAttribute.EntityData.ParentYangName = "path-option-attributes"
    pathOptionAttribute.EntityData.SegmentPath = "path-option-attribute" + types.AddKeyToken(pathOptionAttribute.AttributeSetName, "attribute-set-name")
    pathOptionAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOptionAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOptionAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOptionAttribute.EntityData.Children = types.NewOrderedMap()
    pathOptionAttribute.EntityData.Children.Append("bfd-reverse-path", types.YChild{"BfdReversePath", &pathOptionAttribute.BfdReversePath})
    pathOptionAttribute.EntityData.Children.Append("att-path-option-path-selection", types.YChild{"AttPathOptionPathSelection", &pathOptionAttribute.AttPathOptionPathSelection})
    pathOptionAttribute.EntityData.Children.Append("pce", types.YChild{"Pce", &pathOptionAttribute.Pce})
    pathOptionAttribute.EntityData.Children.Append("affinity-mask", types.YChild{"AffinityMask", &pathOptionAttribute.AffinityMask})
    pathOptionAttribute.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &pathOptionAttribute.Bandwidth})
    pathOptionAttribute.EntityData.Children.Append("new-style-affinity-affinity-types", types.YChild{"NewStyleAffinityAffinityTypes", &pathOptionAttribute.NewStyleAffinityAffinityTypes})
    pathOptionAttribute.EntityData.Leafs = types.NewOrderedMap()
    pathOptionAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", pathOptionAttribute.AttributeSetName})
    pathOptionAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathOptionAttribute.Enable})

    pathOptionAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(pathOptionAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_BfdReversePath
// Configure BFD reverse path
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_BfdReversePath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // BFD reverse path type. The type is BfdReversePath.
    BfdReversePathType interface{}

    // BFD reverse path binding label. The type is interface{} with range:
    // 0..1048575. This attribute is mandatory.
    BindingLabel interface{}
}

func (bfdReversePath *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_BfdReversePath) GetEntityData() *types.CommonEntityData {
    bfdReversePath.EntityData.YFilter = bfdReversePath.YFilter
    bfdReversePath.EntityData.YangName = "bfd-reverse-path"
    bfdReversePath.EntityData.BundleName = "cisco_ios_xr"
    bfdReversePath.EntityData.ParentYangName = "path-option-attribute"
    bfdReversePath.EntityData.SegmentPath = "bfd-reverse-path"
    bfdReversePath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdReversePath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdReversePath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdReversePath.EntityData.Children = types.NewOrderedMap()
    bfdReversePath.EntityData.Leafs = types.NewOrderedMap()
    bfdReversePath.EntityData.Leafs.Append("bfd-reverse-path-type", types.YLeaf{"BfdReversePathType", bfdReversePath.BfdReversePathType})
    bfdReversePath.EntityData.Leafs.Append("binding-label", types.YLeaf{"BindingLabel", bfdReversePath.BindingLabel})

    bfdReversePath.EntityData.YListKeys = []string {}

    return &(bfdReversePath.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection exclude list name configuration. The type is string with
    // length: 1..64.
    PathSelectionExcludeList interface{}

    // Enter path selection configuration. The type is interface{}.
    Enable interface{}

    // Path selection cost limit configuration for this specific tunnel. The type
    // is interface{} with range: 1..4294967295.
    PathSelectionCostLimit interface{}

    // Path invalidation configuration for this specific tunnel.
    Invalidation MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection_Invalidation
}

func (attPathOptionPathSelection *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection) GetEntityData() *types.CommonEntityData {
    attPathOptionPathSelection.EntityData.YFilter = attPathOptionPathSelection.YFilter
    attPathOptionPathSelection.EntityData.YangName = "att-path-option-path-selection"
    attPathOptionPathSelection.EntityData.BundleName = "cisco_ios_xr"
    attPathOptionPathSelection.EntityData.ParentYangName = "path-option-attribute"
    attPathOptionPathSelection.EntityData.SegmentPath = "att-path-option-path-selection"
    attPathOptionPathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attPathOptionPathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attPathOptionPathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attPathOptionPathSelection.EntityData.Children = types.NewOrderedMap()
    attPathOptionPathSelection.EntityData.Children.Append("invalidation", types.YChild{"Invalidation", &attPathOptionPathSelection.Invalidation})
    attPathOptionPathSelection.EntityData.Leafs = types.NewOrderedMap()
    attPathOptionPathSelection.EntityData.Leafs.Append("path-selection-exclude-list", types.YLeaf{"PathSelectionExcludeList", attPathOptionPathSelection.PathSelectionExcludeList})
    attPathOptionPathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", attPathOptionPathSelection.Enable})
    attPathOptionPathSelection.EntityData.Leafs.Append("path-selection-cost-limit", types.YLeaf{"PathSelectionCostLimit", attPathOptionPathSelection.PathSelectionCostLimit})

    attPathOptionPathSelection.EntityData.YListKeys = []string {}

    return &(attPathOptionPathSelection.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection_Invalidation
// Path invalidation configuration for this
// specific tunnel
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection_Invalidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Path Invalidation Timeout. The type is interface{} with range: 0..60000.
    PathInvalidationTimeout interface{}

    // Path Invalidation Action. The type is PathInvalidationAction.
    PathInvalidationAction interface{}
}

func (invalidation *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AttPathOptionPathSelection_Invalidation) GetEntityData() *types.CommonEntityData {
    invalidation.EntityData.YFilter = invalidation.YFilter
    invalidation.EntityData.YangName = "invalidation"
    invalidation.EntityData.BundleName = "cisco_ios_xr"
    invalidation.EntityData.ParentYangName = "att-path-option-path-selection"
    invalidation.EntityData.SegmentPath = "invalidation"
    invalidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidation.EntityData.Children = types.NewOrderedMap()
    invalidation.EntityData.Leafs = types.NewOrderedMap()
    invalidation.EntityData.Leafs.Append("path-invalidation-timeout", types.YLeaf{"PathInvalidationTimeout", invalidation.PathInvalidationTimeout})
    invalidation.EntityData.Leafs.Append("path-invalidation-action", types.YLeaf{"PathInvalidationAction", invalidation.PathInvalidationAction})

    invalidation.EntityData.YListKeys = []string {}

    return &(invalidation.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce
// Configure pce properties
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Always set to true. The type is interface{}.
    Enable interface{}

    // Bidirectional parameters.
    Bidirectional MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_Bidirectional

    // Disjoint path parameters.
    DisjointPath MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_DisjointPath
}

func (pce *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce) GetEntityData() *types.CommonEntityData {
    pce.EntityData.YFilter = pce.YFilter
    pce.EntityData.YangName = "pce"
    pce.EntityData.BundleName = "cisco_ios_xr"
    pce.EntityData.ParentYangName = "path-option-attribute"
    pce.EntityData.SegmentPath = "pce"
    pce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pce.EntityData.Children = types.NewOrderedMap()
    pce.EntityData.Children.Append("bidirectional", types.YChild{"Bidirectional", &pce.Bidirectional})
    pce.EntityData.Children.Append("disjoint-path", types.YChild{"DisjointPath", &pce.DisjointPath})
    pce.EntityData.Leafs = types.NewOrderedMap()
    pce.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pce.Enable})

    pce.EntityData.YListKeys = []string {}

    return &(pce.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_Bidirectional
// Bidirectional parameters
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_Bidirectional struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bidirectional Source IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    BdSourceAddress interface{}

    // Bidirectional Group ID. The type is interface{} with range: 1..4294967295.
    // This attribute is mandatory.
    BdGroupId interface{}
}

func (bidirectional *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_Bidirectional) GetEntityData() *types.CommonEntityData {
    bidirectional.EntityData.YFilter = bidirectional.YFilter
    bidirectional.EntityData.YangName = "bidirectional"
    bidirectional.EntityData.BundleName = "cisco_ios_xr"
    bidirectional.EntityData.ParentYangName = "pce"
    bidirectional.EntityData.SegmentPath = "bidirectional"
    bidirectional.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirectional.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirectional.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirectional.EntityData.Children = types.NewOrderedMap()
    bidirectional.EntityData.Leafs = types.NewOrderedMap()
    bidirectional.EntityData.Leafs.Append("bd-source-address", types.YLeaf{"BdSourceAddress", bidirectional.BdSourceAddress})
    bidirectional.EntityData.Leafs.Append("bd-group-id", types.YLeaf{"BdGroupId", bidirectional.BdGroupId})

    bidirectional.EntityData.YListKeys = []string {}

    return &(bidirectional.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_DisjointPath
// Disjoint path parameters
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_DisjointPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Disjoint Path Source IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    DpSourceAddress interface{}

    // Disjoint Path Type. The type is interface{} with range: 1..3. This
    // attribute is mandatory.
    DpType interface{}

    // Disjoint Path Group ID. The type is interface{} with range: 1..4294967295.
    // This attribute is mandatory.
    DpGroupId interface{}
}

func (disjointPath *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Pce_DisjointPath) GetEntityData() *types.CommonEntityData {
    disjointPath.EntityData.YFilter = disjointPath.YFilter
    disjointPath.EntityData.YangName = "disjoint-path"
    disjointPath.EntityData.BundleName = "cisco_ios_xr"
    disjointPath.EntityData.ParentYangName = "pce"
    disjointPath.EntityData.SegmentPath = "disjoint-path"
    disjointPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPath.EntityData.Children = types.NewOrderedMap()
    disjointPath.EntityData.Leafs = types.NewOrderedMap()
    disjointPath.EntityData.Leafs.Append("dp-source-address", types.YLeaf{"DpSourceAddress", disjointPath.DpSourceAddress})
    disjointPath.EntityData.Leafs.Append("dp-type", types.YLeaf{"DpType", disjointPath.DpType})
    disjointPath.EntityData.Leafs.Append("dp-group-id", types.YLeaf{"DpGroupId", disjointPath.DpGroupId})

    disjointPath.EntityData.YListKeys = []string {}

    return &(disjointPath.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AffinityMask
// Set the affinity flags and mask
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AffinityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Mask interface{}
}

func (affinityMask *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_AffinityMask) GetEntityData() *types.CommonEntityData {
    affinityMask.EntityData.YFilter = affinityMask.YFilter
    affinityMask.EntityData.YangName = "affinity-mask"
    affinityMask.EntityData.BundleName = "cisco_ios_xr"
    affinityMask.EntityData.ParentYangName = "path-option-attribute"
    affinityMask.EntityData.SegmentPath = "affinity-mask"
    affinityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMask.EntityData.Children = types.NewOrderedMap()
    affinityMask.EntityData.Leafs = types.NewOrderedMap()
    affinityMask.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", affinityMask.Affinity})
    affinityMask.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", affinityMask.Mask})

    affinityMask.EntityData.YListKeys = []string {}

    return &(affinityMask.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Bandwidth
// Tunnel bandwidth requirement
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DSTE-standard flag. The type is MplsTeBandwidthDste. This attribute is
    // mandatory.
    DsteType interface{}

    // Class type for the bandwidth allocation. The type is interface{} with
    // range: 0..1. This attribute is mandatory.
    ClassOrPoolType interface{}

    // The value of the bandwidth reserved by this tunnel in kbps. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    Bandwidth interface{}
}

func (bandwidth *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "path-option-attribute"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Leafs = types.NewOrderedMap()
    bandwidth.EntityData.Leafs.Append("dste-type", types.YLeaf{"DsteType", bandwidth.DsteType})
    bandwidth.EntityData.Leafs.Append("class-or-pool-type", types.YLeaf{"ClassOrPoolType", bandwidth.ClassOrPoolType})
    bandwidth.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", bandwidth.Bandwidth})

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes
// Tunnel new style affinity attributes table
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType.
    NewStyleAffinityAffinityType []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1.
    NewStyleAffinityAffinityTypeAffinity1 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2.
    NewStyleAffinityAffinityTypeAffinity1Affinity2 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 []*MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
}

func (newStyleAffinityAffinityTypes *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypes.EntityData.YFilter = newStyleAffinityAffinityTypes.YFilter
    newStyleAffinityAffinityTypes.EntityData.YangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypes.EntityData.ParentYangName = "path-option-attribute"
    newStyleAffinityAffinityTypes.EntityData.SegmentPath = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypes.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type", types.YChild{"NewStyleAffinityAffinityType", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]), types.YChild{"NewStyleAffinityAffinityType", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1", types.YChild{"NewStyleAffinityAffinityTypeAffinity1", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Leafs = types.NewOrderedMap()

    newStyleAffinityAffinityTypes.EntityData.YListKeys = []string {}

    return &(newStyleAffinityAffinityTypes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}
}

func (newStyleAffinityAffinityType *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityType.EntityData.YFilter = newStyleAffinityAffinityType.YFilter
    newStyleAffinityAffinityType.EntityData.YangName = "new-style-affinity-affinity-type"
    newStyleAffinityAffinityType.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityType.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityType.EntityData.SegmentPath = "new-style-affinity-affinity-type" + types.AddKeyToken(newStyleAffinityAffinityType.AffinityType, "affinity-type")
    newStyleAffinityAffinityType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityType.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityType.AffinityType})

    newStyleAffinityAffinityType.EntityData.YListKeys = []string {"AffinityType"}

    return &(newStyleAffinityAffinityType.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1.YFilter
    newStyleAffinityAffinityTypeAffinity1.EntityData.YangName = "new-style-affinity-affinity-type-affinity1"
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.Affinity1, "affinity1")
    newStyleAffinityAffinityTypeAffinity1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1.AffinityType})
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1.Affinity1})

    newStyleAffinityAffinityTypeAffinity1.EntityData.YListKeys = []string {"AffinityType", "Affinity1"}

    return &(newStyleAffinityAffinityTypeAffinity1.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2, "affinity2")
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2})

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3, "affinity3")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4, "affinity4")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5, "affinity5")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6, "affinity6")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7, "affinity7")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8, "affinity8")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9, "affinity9")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}

    // This attribute is a key. The name of the tenth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity10 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 *MplsTe_GlobalAttributes_AttributeSet_PathOptionAttributes_PathOptionAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9, "affinity9") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10, "affinity10")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity10", types.YLeaf{"Affinity10", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9", "Affinity10"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes
// P2MP-TE Tunnel AttributeSets Table
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // P2MP-TE Tunnel Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute.
    P2mpteAttribute []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute
}

func (p2mpteAttributes *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes) GetEntityData() *types.CommonEntityData {
    p2mpteAttributes.EntityData.YFilter = p2mpteAttributes.YFilter
    p2mpteAttributes.EntityData.YangName = "p2mpte-attributes"
    p2mpteAttributes.EntityData.BundleName = "cisco_ios_xr"
    p2mpteAttributes.EntityData.ParentYangName = "attribute-set"
    p2mpteAttributes.EntityData.SegmentPath = "p2mpte-attributes"
    p2mpteAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2mpteAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2mpteAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2mpteAttributes.EntityData.Children = types.NewOrderedMap()
    p2mpteAttributes.EntityData.Children.Append("p2mpte-attribute", types.YChild{"P2mpteAttribute", nil})
    for i := range p2mpteAttributes.P2mpteAttribute {
        p2mpteAttributes.EntityData.Children.Append(types.GetSegmentPath(p2mpteAttributes.P2mpteAttribute[i]), types.YChild{"P2mpteAttribute", p2mpteAttributes.P2mpteAttribute[i]})
    }
    p2mpteAttributes.EntityData.Leafs = types.NewOrderedMap()

    p2mpteAttributes.EntityData.YListKeys = []string {}

    return &(p2mpteAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute
// P2MP-TE Tunnel Attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // The bandwidth of the interface in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    InterfaceBandwidth interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Record the route used by the tunnel. The type is interface{}.
    RecordRoute interface{}

    // Tunnel Setup and Hold Priorities.
    Priority MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Priority

    // Set the affinity flags and mask.
    AffinityMask MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_AffinityMask

    // Tunnel bandwidth requirement.
    Bandwidth MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Bandwidth

    // Configure path selection properties.
    PathSelection MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_PathSelection

    // Tunnel new style affinity attributes table.
    NewStyleAffinityAffinityTypes MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes

    // Specify MPLS tunnel can be fast-rerouted.
    FastReroute MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_FastReroute

    // Log tunnel LSP messages.
    Logging MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Logging
}

func (p2mpteAttribute *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute) GetEntityData() *types.CommonEntityData {
    p2mpteAttribute.EntityData.YFilter = p2mpteAttribute.YFilter
    p2mpteAttribute.EntityData.YangName = "p2mpte-attribute"
    p2mpteAttribute.EntityData.BundleName = "cisco_ios_xr"
    p2mpteAttribute.EntityData.ParentYangName = "p2mpte-attributes"
    p2mpteAttribute.EntityData.SegmentPath = "p2mpte-attribute" + types.AddKeyToken(p2mpteAttribute.AttributeSetName, "attribute-set-name")
    p2mpteAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2mpteAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2mpteAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2mpteAttribute.EntityData.Children = types.NewOrderedMap()
    p2mpteAttribute.EntityData.Children.Append("priority", types.YChild{"Priority", &p2mpteAttribute.Priority})
    p2mpteAttribute.EntityData.Children.Append("affinity-mask", types.YChild{"AffinityMask", &p2mpteAttribute.AffinityMask})
    p2mpteAttribute.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &p2mpteAttribute.Bandwidth})
    p2mpteAttribute.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &p2mpteAttribute.PathSelection})
    p2mpteAttribute.EntityData.Children.Append("new-style-affinity-affinity-types", types.YChild{"NewStyleAffinityAffinityTypes", &p2mpteAttribute.NewStyleAffinityAffinityTypes})
    p2mpteAttribute.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &p2mpteAttribute.FastReroute})
    p2mpteAttribute.EntityData.Children.Append("logging", types.YChild{"Logging", &p2mpteAttribute.Logging})
    p2mpteAttribute.EntityData.Leafs = types.NewOrderedMap()
    p2mpteAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", p2mpteAttribute.AttributeSetName})
    p2mpteAttribute.EntityData.Leafs.Append("interface-bandwidth", types.YLeaf{"InterfaceBandwidth", p2mpteAttribute.InterfaceBandwidth})
    p2mpteAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", p2mpteAttribute.Enable})
    p2mpteAttribute.EntityData.Leafs.Append("record-route", types.YLeaf{"RecordRoute", p2mpteAttribute.RecordRoute})

    p2mpteAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(p2mpteAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "p2mpte-attribute"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", priority.SetupPriority})
    priority.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", priority.HoldPriority})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_AffinityMask
// Set the affinity flags and mask
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_AffinityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Mask interface{}
}

func (affinityMask *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_AffinityMask) GetEntityData() *types.CommonEntityData {
    affinityMask.EntityData.YFilter = affinityMask.YFilter
    affinityMask.EntityData.YangName = "affinity-mask"
    affinityMask.EntityData.BundleName = "cisco_ios_xr"
    affinityMask.EntityData.ParentYangName = "p2mpte-attribute"
    affinityMask.EntityData.SegmentPath = "affinity-mask"
    affinityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMask.EntityData.Children = types.NewOrderedMap()
    affinityMask.EntityData.Leafs = types.NewOrderedMap()
    affinityMask.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", affinityMask.Affinity})
    affinityMask.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", affinityMask.Mask})

    affinityMask.EntityData.YListKeys = []string {}

    return &(affinityMask.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Bandwidth
// Tunnel bandwidth requirement
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DSTE-standard flag. The type is MplsTeBandwidthDste. This attribute is
    // mandatory.
    DsteType interface{}

    // Class type for the bandwidth allocation. The type is interface{} with
    // range: 0..1. This attribute is mandatory.
    ClassOrPoolType interface{}

    // The value of the bandwidth reserved by this tunnel in kbps. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    Bandwidth interface{}
}

func (bandwidth *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "p2mpte-attribute"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Leafs = types.NewOrderedMap()
    bandwidth.EntityData.Leafs.Append("dste-type", types.YLeaf{"DsteType", bandwidth.DsteType})
    bandwidth.EntityData.Leafs.Append("class-or-pool-type", types.YLeaf{"ClassOrPoolType", bandwidth.ClassOrPoolType})
    bandwidth.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", bandwidth.Bandwidth})

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_PathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path selection. The type is interface{}.
    Enable interface{}
}

func (pathSelection *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "p2mpte-attribute"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSelection.Enable})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes
// Tunnel new style affinity attributes table
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType.
    NewStyleAffinityAffinityType []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1.
    NewStyleAffinityAffinityTypeAffinity1 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2.
    NewStyleAffinityAffinityTypeAffinity1Affinity2 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 []*MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
}

func (newStyleAffinityAffinityTypes *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypes.EntityData.YFilter = newStyleAffinityAffinityTypes.YFilter
    newStyleAffinityAffinityTypes.EntityData.YangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypes.EntityData.ParentYangName = "p2mpte-attribute"
    newStyleAffinityAffinityTypes.EntityData.SegmentPath = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypes.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type", types.YChild{"NewStyleAffinityAffinityType", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]), types.YChild{"NewStyleAffinityAffinityType", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1", types.YChild{"NewStyleAffinityAffinityTypeAffinity1", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Leafs = types.NewOrderedMap()

    newStyleAffinityAffinityTypes.EntityData.YListKeys = []string {}

    return &(newStyleAffinityAffinityTypes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}
}

func (newStyleAffinityAffinityType *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityType.EntityData.YFilter = newStyleAffinityAffinityType.YFilter
    newStyleAffinityAffinityType.EntityData.YangName = "new-style-affinity-affinity-type"
    newStyleAffinityAffinityType.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityType.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityType.EntityData.SegmentPath = "new-style-affinity-affinity-type" + types.AddKeyToken(newStyleAffinityAffinityType.AffinityType, "affinity-type")
    newStyleAffinityAffinityType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityType.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityType.AffinityType})

    newStyleAffinityAffinityType.EntityData.YListKeys = []string {"AffinityType"}

    return &(newStyleAffinityAffinityType.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1.YFilter
    newStyleAffinityAffinityTypeAffinity1.EntityData.YangName = "new-style-affinity-affinity-type-affinity1"
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.Affinity1, "affinity1")
    newStyleAffinityAffinityTypeAffinity1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1.AffinityType})
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1.Affinity1})

    newStyleAffinityAffinityTypeAffinity1.EntityData.YListKeys = []string {"AffinityType", "Affinity1"}

    return &(newStyleAffinityAffinityTypeAffinity1.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2, "affinity2")
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2})

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3, "affinity3")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4, "affinity4")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5, "affinity5")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6, "affinity6")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7, "affinity7")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8, "affinity8")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9, "affinity9")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}

    // This attribute is a key. The name of the tenth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity10 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9, "affinity9") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10, "affinity10")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity10", types.YLeaf{"Affinity10", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9", "Affinity10"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_FastReroute
// Specify MPLS tunnel can be fast-rerouted
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bandwidth Protection. The type is interface{} with range: 0..1. This
    // attribute is mandatory.
    BandwidthProtection interface{}

    // Node Protection. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    NodeProtection interface{}
}

func (fastReroute *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "p2mpte-attribute"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("bandwidth-protection", types.YLeaf{"BandwidthProtection", fastReroute.BandwidthProtection})
    fastReroute.EntityData.Leafs.Append("node-protection", types.YLeaf{"NodeProtection", fastReroute.NodeProtection})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Logging
// Log tunnel LSP messages
type MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log tunnel messages for insufficient bandwidth. The type is interface{}.
    InsufficientBwMessage interface{}

    // Log tunnel reoptimized messages. The type is interface{}.
    ReoptimizedMessage interface{}

    // Log tunnel bandwidth change messages. The type is interface{}.
    BandwidthChangeMessage interface{}

    // Log all events for a tunnel. The type is interface{}.
    All interface{}

    // Enable logging for path-calculation failures. The type is interface{}.
    PcalcFailureMessage interface{}

    // Log tunnel state messages. The type is interface{}.
    StateMessage interface{}

    // Log tunnel reoptimization attempts messages. The type is interface{}.
    ReoptimizeAttemptsMessage interface{}

    // Log all tunnel sub-LSP state messages. The type is interface{}.
    SubLspStateMessage interface{}

    // Log tunnel rereoute messages. The type is interface{}.
    RerouteMesssage interface{}
}

func (logging *MplsTe_GlobalAttributes_AttributeSet_P2mpteAttributes_P2mpteAttribute_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "p2mpte-attribute"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("insufficient-bw-message", types.YLeaf{"InsufficientBwMessage", logging.InsufficientBwMessage})
    logging.EntityData.Leafs.Append("reoptimized-message", types.YLeaf{"ReoptimizedMessage", logging.ReoptimizedMessage})
    logging.EntityData.Leafs.Append("bandwidth-change-message", types.YLeaf{"BandwidthChangeMessage", logging.BandwidthChangeMessage})
    logging.EntityData.Leafs.Append("all", types.YLeaf{"All", logging.All})
    logging.EntityData.Leafs.Append("pcalc-failure-message", types.YLeaf{"PcalcFailureMessage", logging.PcalcFailureMessage})
    logging.EntityData.Leafs.Append("state-message", types.YLeaf{"StateMessage", logging.StateMessage})
    logging.EntityData.Leafs.Append("reoptimize-attempts-message", types.YLeaf{"ReoptimizeAttemptsMessage", logging.ReoptimizeAttemptsMessage})
    logging.EntityData.Leafs.Append("sub-lsp-state-message", types.YLeaf{"SubLspStateMessage", logging.SubLspStateMessage})
    logging.EntityData.Leafs.Append("reroute-messsage", types.YLeaf{"RerouteMesssage", logging.RerouteMesssage})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes
// P2P-TE Tunnel AttributeSets Table
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // P2P-TE Tunnel Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute.
    P2pTeAttribute []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute
}

func (p2pTeAttributes *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes) GetEntityData() *types.CommonEntityData {
    p2pTeAttributes.EntityData.YFilter = p2pTeAttributes.YFilter
    p2pTeAttributes.EntityData.YangName = "p2p-te-attributes"
    p2pTeAttributes.EntityData.BundleName = "cisco_ios_xr"
    p2pTeAttributes.EntityData.ParentYangName = "attribute-set"
    p2pTeAttributes.EntityData.SegmentPath = "p2p-te-attributes"
    p2pTeAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2pTeAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2pTeAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2pTeAttributes.EntityData.Children = types.NewOrderedMap()
    p2pTeAttributes.EntityData.Children.Append("p2p-te-attribute", types.YChild{"P2pTeAttribute", nil})
    for i := range p2pTeAttributes.P2pTeAttribute {
        p2pTeAttributes.EntityData.Children.Append(types.GetSegmentPath(p2pTeAttributes.P2pTeAttribute[i]), types.YChild{"P2pTeAttribute", p2pTeAttributes.P2pTeAttribute[i]})
    }
    p2pTeAttributes.EntityData.Leafs = types.NewOrderedMap()

    p2pTeAttributes.EntityData.YListKeys = []string {}

    return &(p2pTeAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute
// P2P-TE Tunnel Attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Configure path selection properties.
    PathSelection MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection

    // Configure pce properties.
    Pce MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce

    // Set the affinity flags and mask.
    AffinityMask MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_AffinityMask

    // Log tunnel LSP messages.
    Logging MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Logging

    // Tunnel new style affinity attributes table.
    NewStyleAffinityAffinityTypes MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes
}

func (p2pTeAttribute *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute) GetEntityData() *types.CommonEntityData {
    p2pTeAttribute.EntityData.YFilter = p2pTeAttribute.YFilter
    p2pTeAttribute.EntityData.YangName = "p2p-te-attribute"
    p2pTeAttribute.EntityData.BundleName = "cisco_ios_xr"
    p2pTeAttribute.EntityData.ParentYangName = "p2p-te-attributes"
    p2pTeAttribute.EntityData.SegmentPath = "p2p-te-attribute" + types.AddKeyToken(p2pTeAttribute.AttributeSetName, "attribute-set-name")
    p2pTeAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    p2pTeAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    p2pTeAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    p2pTeAttribute.EntityData.Children = types.NewOrderedMap()
    p2pTeAttribute.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &p2pTeAttribute.PathSelection})
    p2pTeAttribute.EntityData.Children.Append("pce", types.YChild{"Pce", &p2pTeAttribute.Pce})
    p2pTeAttribute.EntityData.Children.Append("affinity-mask", types.YChild{"AffinityMask", &p2pTeAttribute.AffinityMask})
    p2pTeAttribute.EntityData.Children.Append("logging", types.YChild{"Logging", &p2pTeAttribute.Logging})
    p2pTeAttribute.EntityData.Children.Append("new-style-affinity-affinity-types", types.YChild{"NewStyleAffinityAffinityTypes", &p2pTeAttribute.NewStyleAffinityAffinityTypes})
    p2pTeAttribute.EntityData.Leafs = types.NewOrderedMap()
    p2pTeAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", p2pTeAttribute.AttributeSetName})
    p2pTeAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", p2pTeAttribute.Enable})

    p2pTeAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(p2pTeAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection metric to use in path calculation. The type is
    // MplsTePathSelectionMetric.
    PathSelectionMetric interface{}

    // Segment routing adjacency protection type to use in path calculation. The
    // type is MplsTePathSelectionSegmentRoutingAdjacencyProtection.
    PathSelectionSegmentRoutingAdjacencyProtection interface{}

    // Enter path selection configuration. The type is interface{}.
    Enable interface{}

    // Path selection segment routing prepend configuration.
    SegmentRoutingPrepend MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend

    // Path selection invalidation configuration.
    Invalidation MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_Invalidation
}

func (pathSelection *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "p2p-te-attribute"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Children.Append("segment-routing-prepend", types.YChild{"SegmentRoutingPrepend", &pathSelection.SegmentRoutingPrepend})
    pathSelection.EntityData.Children.Append("invalidation", types.YChild{"Invalidation", &pathSelection.Invalidation})
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("path-selection-metric", types.YLeaf{"PathSelectionMetric", pathSelection.PathSelectionMetric})
    pathSelection.EntityData.Leafs.Append("path-selection-segment-routing-adjacency-protection", types.YLeaf{"PathSelectionSegmentRoutingAdjacencyProtection", pathSelection.PathSelectionSegmentRoutingAdjacencyProtection})
    pathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSelection.Enable})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend
// Path selection segment routing prepend
// configuration
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter path selection segment routing prepend submode. The type is
    // interface{}.
    Enable interface{}

    // Segment routing prepend index table.
    Indexes MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes
}

func (segmentRoutingPrepend *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend) GetEntityData() *types.CommonEntityData {
    segmentRoutingPrepend.EntityData.YFilter = segmentRoutingPrepend.YFilter
    segmentRoutingPrepend.EntityData.YangName = "segment-routing-prepend"
    segmentRoutingPrepend.EntityData.BundleName = "cisco_ios_xr"
    segmentRoutingPrepend.EntityData.ParentYangName = "path-selection"
    segmentRoutingPrepend.EntityData.SegmentPath = "segment-routing-prepend"
    segmentRoutingPrepend.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segmentRoutingPrepend.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segmentRoutingPrepend.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segmentRoutingPrepend.EntityData.Children = types.NewOrderedMap()
    segmentRoutingPrepend.EntityData.Children.Append("indexes", types.YChild{"Indexes", &segmentRoutingPrepend.Indexes})
    segmentRoutingPrepend.EntityData.Leafs = types.NewOrderedMap()
    segmentRoutingPrepend.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", segmentRoutingPrepend.Enable})

    segmentRoutingPrepend.EntityData.YListKeys = []string {}

    return &(segmentRoutingPrepend.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes
// Segment routing prepend index table
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prepend index information. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes_Index.
    Index []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes_Index
}

func (indexes *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes) GetEntityData() *types.CommonEntityData {
    indexes.EntityData.YFilter = indexes.YFilter
    indexes.EntityData.YangName = "indexes"
    indexes.EntityData.BundleName = "cisco_ios_xr"
    indexes.EntityData.ParentYangName = "segment-routing-prepend"
    indexes.EntityData.SegmentPath = "indexes"
    indexes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    indexes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    indexes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    indexes.EntityData.Children = types.NewOrderedMap()
    indexes.EntityData.Children.Append("index", types.YChild{"Index", nil})
    for i := range indexes.Index {
        indexes.EntityData.Children.Append(types.GetSegmentPath(indexes.Index[i]), types.YChild{"Index", indexes.Index[i]})
    }
    indexes.EntityData.Leafs = types.NewOrderedMap()

    indexes.EntityData.YListKeys = []string {}

    return &(indexes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes_Index
// Prepend index information
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes_Index struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index number. The type is interface{} with range:
    // 1..10.
    IndexNumber interface{}

    // Prepend type. The type is SrPrepend. The default value is none-type.
    PrependType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295. The default
    // value is 1048577.
    MplsLabel interface{}
}

func (index *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_SegmentRoutingPrepend_Indexes_Index) GetEntityData() *types.CommonEntityData {
    index.EntityData.YFilter = index.YFilter
    index.EntityData.YangName = "index"
    index.EntityData.BundleName = "cisco_ios_xr"
    index.EntityData.ParentYangName = "indexes"
    index.EntityData.SegmentPath = "index" + types.AddKeyToken(index.IndexNumber, "index-number")
    index.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    index.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    index.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    index.EntityData.Children = types.NewOrderedMap()
    index.EntityData.Leafs = types.NewOrderedMap()
    index.EntityData.Leafs.Append("index-number", types.YLeaf{"IndexNumber", index.IndexNumber})
    index.EntityData.Leafs.Append("prepend-type", types.YLeaf{"PrependType", index.PrependType})
    index.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", index.MplsLabel})

    index.EntityData.YListKeys = []string {"IndexNumber"}

    return &(index.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_Invalidation
// Path selection invalidation configuration
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_Invalidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection invalidation timer value (milliseconds). The type is
    // interface{} with range: 0..60000. Units are millisecond.
    InvalidationTimer interface{}

    // Path selection invalidation timer expire type. The type is
    // MplsTePathSelectionInvalidationTimerExpire. The default value is
    // tunnel-action-tear.
    InvalidationTimerExpireType interface{}
}

func (invalidation *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_PathSelection_Invalidation) GetEntityData() *types.CommonEntityData {
    invalidation.EntityData.YFilter = invalidation.YFilter
    invalidation.EntityData.YangName = "invalidation"
    invalidation.EntityData.BundleName = "cisco_ios_xr"
    invalidation.EntityData.ParentYangName = "path-selection"
    invalidation.EntityData.SegmentPath = "invalidation"
    invalidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidation.EntityData.Children = types.NewOrderedMap()
    invalidation.EntityData.Leafs = types.NewOrderedMap()
    invalidation.EntityData.Leafs.Append("invalidation-timer", types.YLeaf{"InvalidationTimer", invalidation.InvalidationTimer})
    invalidation.EntityData.Leafs.Append("invalidation-timer-expire-type", types.YLeaf{"InvalidationTimerExpireType", invalidation.InvalidationTimerExpireType})

    invalidation.EntityData.YListKeys = []string {}

    return &(invalidation.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce
// Configure pce properties
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Always set to true. The type is interface{}.
    Enable interface{}

    // Bidirectional parameters.
    Bidirectional MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_Bidirectional

    // Disjoint path parameters.
    DisjointPath MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_DisjointPath
}

func (pce *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce) GetEntityData() *types.CommonEntityData {
    pce.EntityData.YFilter = pce.YFilter
    pce.EntityData.YangName = "pce"
    pce.EntityData.BundleName = "cisco_ios_xr"
    pce.EntityData.ParentYangName = "p2p-te-attribute"
    pce.EntityData.SegmentPath = "pce"
    pce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pce.EntityData.Children = types.NewOrderedMap()
    pce.EntityData.Children.Append("bidirectional", types.YChild{"Bidirectional", &pce.Bidirectional})
    pce.EntityData.Children.Append("disjoint-path", types.YChild{"DisjointPath", &pce.DisjointPath})
    pce.EntityData.Leafs = types.NewOrderedMap()
    pce.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pce.Enable})

    pce.EntityData.YListKeys = []string {}

    return &(pce.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_Bidirectional
// Bidirectional parameters
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_Bidirectional struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bidirectional Source IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    BdSourceAddress interface{}

    // Bidirectional Group ID. The type is interface{} with range: 1..4294967295.
    // This attribute is mandatory.
    BdGroupId interface{}
}

func (bidirectional *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_Bidirectional) GetEntityData() *types.CommonEntityData {
    bidirectional.EntityData.YFilter = bidirectional.YFilter
    bidirectional.EntityData.YangName = "bidirectional"
    bidirectional.EntityData.BundleName = "cisco_ios_xr"
    bidirectional.EntityData.ParentYangName = "pce"
    bidirectional.EntityData.SegmentPath = "bidirectional"
    bidirectional.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bidirectional.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bidirectional.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bidirectional.EntityData.Children = types.NewOrderedMap()
    bidirectional.EntityData.Leafs = types.NewOrderedMap()
    bidirectional.EntityData.Leafs.Append("bd-source-address", types.YLeaf{"BdSourceAddress", bidirectional.BdSourceAddress})
    bidirectional.EntityData.Leafs.Append("bd-group-id", types.YLeaf{"BdGroupId", bidirectional.BdGroupId})

    bidirectional.EntityData.YListKeys = []string {}

    return &(bidirectional.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_DisjointPath
// Disjoint path parameters
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_DisjointPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Disjoint Path Source IP Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    DpSourceAddress interface{}

    // Disjoint Path Type. The type is interface{} with range: 1..3. This
    // attribute is mandatory.
    DpType interface{}

    // Disjoint Path Group ID. The type is interface{} with range: 1..4294967295.
    // This attribute is mandatory.
    DpGroupId interface{}
}

func (disjointPath *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Pce_DisjointPath) GetEntityData() *types.CommonEntityData {
    disjointPath.EntityData.YFilter = disjointPath.YFilter
    disjointPath.EntityData.YangName = "disjoint-path"
    disjointPath.EntityData.BundleName = "cisco_ios_xr"
    disjointPath.EntityData.ParentYangName = "pce"
    disjointPath.EntityData.SegmentPath = "disjoint-path"
    disjointPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPath.EntityData.Children = types.NewOrderedMap()
    disjointPath.EntityData.Leafs = types.NewOrderedMap()
    disjointPath.EntityData.Leafs.Append("dp-source-address", types.YLeaf{"DpSourceAddress", disjointPath.DpSourceAddress})
    disjointPath.EntityData.Leafs.Append("dp-type", types.YLeaf{"DpType", disjointPath.DpType})
    disjointPath.EntityData.Leafs.Append("dp-group-id", types.YLeaf{"DpGroupId", disjointPath.DpGroupId})

    disjointPath.EntityData.YListKeys = []string {}

    return &(disjointPath.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_AffinityMask
// Set the affinity flags and mask
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_AffinityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Mask interface{}
}

func (affinityMask *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_AffinityMask) GetEntityData() *types.CommonEntityData {
    affinityMask.EntityData.YFilter = affinityMask.YFilter
    affinityMask.EntityData.YangName = "affinity-mask"
    affinityMask.EntityData.BundleName = "cisco_ios_xr"
    affinityMask.EntityData.ParentYangName = "p2p-te-attribute"
    affinityMask.EntityData.SegmentPath = "affinity-mask"
    affinityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMask.EntityData.Children = types.NewOrderedMap()
    affinityMask.EntityData.Leafs = types.NewOrderedMap()
    affinityMask.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", affinityMask.Affinity})
    affinityMask.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", affinityMask.Mask})

    affinityMask.EntityData.YListKeys = []string {}

    return &(affinityMask.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Logging
// Log tunnel LSP messages
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log tunnel messages for bandwidth change. The type is interface{}.
    LspSwitchOverChangeMessage interface{}

    // Log all events for a tunnel. The type is interface{}.
    All interface{}

    // Log tunnel record-route messages. The type is interface{}.
    RecordRouteMesssage interface{}

    // Enable BFD session state change alarm. The type is interface{}.
    BfdStateMessage interface{}

    // Log tunnel messages for bandwidth change. The type is interface{}.
    BandwidthChangeMessage interface{}

    // Log tunnel reoptimization attempts messages. The type is interface{}.
    ReoptimizeAttemptsMessage interface{}

    // Log tunnel rereoute messages. The type is interface{}.
    RerouteMesssage interface{}

    // Log tunnel state messages. The type is interface{}.
    StateMessage interface{}

    // Log tunnel messages for insufficient bandwidth. The type is interface{}.
    InsufficientBwMessage interface{}

    // Log tunnel reoptimized messages. The type is interface{}.
    ReoptimizedMessage interface{}

    // Enable logging for path-calculation failures. The type is interface{}.
    PcalcFailureMessage interface{}
}

func (logging *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "p2p-te-attribute"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("lsp-switch-over-change-message", types.YLeaf{"LspSwitchOverChangeMessage", logging.LspSwitchOverChangeMessage})
    logging.EntityData.Leafs.Append("all", types.YLeaf{"All", logging.All})
    logging.EntityData.Leafs.Append("record-route-messsage", types.YLeaf{"RecordRouteMesssage", logging.RecordRouteMesssage})
    logging.EntityData.Leafs.Append("bfd-state-message", types.YLeaf{"BfdStateMessage", logging.BfdStateMessage})
    logging.EntityData.Leafs.Append("bandwidth-change-message", types.YLeaf{"BandwidthChangeMessage", logging.BandwidthChangeMessage})
    logging.EntityData.Leafs.Append("reoptimize-attempts-message", types.YLeaf{"ReoptimizeAttemptsMessage", logging.ReoptimizeAttemptsMessage})
    logging.EntityData.Leafs.Append("reroute-messsage", types.YLeaf{"RerouteMesssage", logging.RerouteMesssage})
    logging.EntityData.Leafs.Append("state-message", types.YLeaf{"StateMessage", logging.StateMessage})
    logging.EntityData.Leafs.Append("insufficient-bw-message", types.YLeaf{"InsufficientBwMessage", logging.InsufficientBwMessage})
    logging.EntityData.Leafs.Append("reoptimized-message", types.YLeaf{"ReoptimizedMessage", logging.ReoptimizedMessage})
    logging.EntityData.Leafs.Append("pcalc-failure-message", types.YLeaf{"PcalcFailureMessage", logging.PcalcFailureMessage})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes
// Tunnel new style affinity attributes table
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType.
    NewStyleAffinityAffinityType []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1.
    NewStyleAffinityAffinityTypeAffinity1 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2.
    NewStyleAffinityAffinityTypeAffinity1Affinity2 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 []*MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
}

func (newStyleAffinityAffinityTypes *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypes.EntityData.YFilter = newStyleAffinityAffinityTypes.YFilter
    newStyleAffinityAffinityTypes.EntityData.YangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypes.EntityData.ParentYangName = "p2p-te-attribute"
    newStyleAffinityAffinityTypes.EntityData.SegmentPath = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypes.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type", types.YChild{"NewStyleAffinityAffinityType", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]), types.YChild{"NewStyleAffinityAffinityType", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1", types.YChild{"NewStyleAffinityAffinityTypeAffinity1", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Leafs = types.NewOrderedMap()

    newStyleAffinityAffinityTypes.EntityData.YListKeys = []string {}

    return &(newStyleAffinityAffinityTypes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}
}

func (newStyleAffinityAffinityType *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityType.EntityData.YFilter = newStyleAffinityAffinityType.YFilter
    newStyleAffinityAffinityType.EntityData.YangName = "new-style-affinity-affinity-type"
    newStyleAffinityAffinityType.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityType.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityType.EntityData.SegmentPath = "new-style-affinity-affinity-type" + types.AddKeyToken(newStyleAffinityAffinityType.AffinityType, "affinity-type")
    newStyleAffinityAffinityType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityType.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityType.AffinityType})

    newStyleAffinityAffinityType.EntityData.YListKeys = []string {"AffinityType"}

    return &(newStyleAffinityAffinityType.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1.YFilter
    newStyleAffinityAffinityTypeAffinity1.EntityData.YangName = "new-style-affinity-affinity-type-affinity1"
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.Affinity1, "affinity1")
    newStyleAffinityAffinityTypeAffinity1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1.AffinityType})
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1.Affinity1})

    newStyleAffinityAffinityTypeAffinity1.EntityData.YListKeys = []string {"AffinityType", "Affinity1"}

    return &(newStyleAffinityAffinityTypeAffinity1.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2, "affinity2")
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2})

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3, "affinity3")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4, "affinity4")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5, "affinity5")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6, "affinity6")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7, "affinity7")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8, "affinity8")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9, "affinity9")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}

    // This attribute is a key. The name of the tenth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity10 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 *MplsTe_GlobalAttributes_AttributeSet_P2pTeAttributes_P2pTeAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9, "affinity9") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10, "affinity10")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity10", types.YLeaf{"Affinity10", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9", "Affinity10"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes
// Auto-backup Tunnel Attribute Table
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto-backup Tunnel Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute.
    AutoBackupAttribute []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute
}

func (autoBackupAttributes *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes) GetEntityData() *types.CommonEntityData {
    autoBackupAttributes.EntityData.YFilter = autoBackupAttributes.YFilter
    autoBackupAttributes.EntityData.YangName = "auto-backup-attributes"
    autoBackupAttributes.EntityData.BundleName = "cisco_ios_xr"
    autoBackupAttributes.EntityData.ParentYangName = "attribute-set"
    autoBackupAttributes.EntityData.SegmentPath = "auto-backup-attributes"
    autoBackupAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoBackupAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoBackupAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoBackupAttributes.EntityData.Children = types.NewOrderedMap()
    autoBackupAttributes.EntityData.Children.Append("auto-backup-attribute", types.YChild{"AutoBackupAttribute", nil})
    for i := range autoBackupAttributes.AutoBackupAttribute {
        autoBackupAttributes.EntityData.Children.Append(types.GetSegmentPath(autoBackupAttributes.AutoBackupAttribute[i]), types.YChild{"AutoBackupAttribute", autoBackupAttributes.AutoBackupAttribute[i]})
    }
    autoBackupAttributes.EntityData.Leafs = types.NewOrderedMap()

    autoBackupAttributes.EntityData.YListKeys = []string {}

    return &(autoBackupAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute
// Auto-backup Tunnel Attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Record the route used by the tunnel. The type is interface{}.
    RecordRoute interface{}

    // Signalled name.
    SignalledName MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_SignalledName

    // Log tunnel LSP messages.
    AutoBackupLogging MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AutoBackupLogging

    // Tunnel Setup and Hold Priorities.
    Priority MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_Priority

    // Set the affinity flags and mask.
    AffinityMask MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AffinityMask

    // Configure path selection properties.
    PathSelection MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PathSelection

    // Policy classes for PBTS.
    PolicyClasses MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PolicyClasses

    // Tunnel new style affinity attributes table.
    NewStyleAffinityAffinityTypes MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes
}

func (autoBackupAttribute *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute) GetEntityData() *types.CommonEntityData {
    autoBackupAttribute.EntityData.YFilter = autoBackupAttribute.YFilter
    autoBackupAttribute.EntityData.YangName = "auto-backup-attribute"
    autoBackupAttribute.EntityData.BundleName = "cisco_ios_xr"
    autoBackupAttribute.EntityData.ParentYangName = "auto-backup-attributes"
    autoBackupAttribute.EntityData.SegmentPath = "auto-backup-attribute" + types.AddKeyToken(autoBackupAttribute.AttributeSetName, "attribute-set-name")
    autoBackupAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoBackupAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoBackupAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoBackupAttribute.EntityData.Children = types.NewOrderedMap()
    autoBackupAttribute.EntityData.Children.Append("signalled-name", types.YChild{"SignalledName", &autoBackupAttribute.SignalledName})
    autoBackupAttribute.EntityData.Children.Append("auto-backup-logging", types.YChild{"AutoBackupLogging", &autoBackupAttribute.AutoBackupLogging})
    autoBackupAttribute.EntityData.Children.Append("priority", types.YChild{"Priority", &autoBackupAttribute.Priority})
    autoBackupAttribute.EntityData.Children.Append("affinity-mask", types.YChild{"AffinityMask", &autoBackupAttribute.AffinityMask})
    autoBackupAttribute.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &autoBackupAttribute.PathSelection})
    autoBackupAttribute.EntityData.Children.Append("policy-classes", types.YChild{"PolicyClasses", &autoBackupAttribute.PolicyClasses})
    autoBackupAttribute.EntityData.Children.Append("new-style-affinity-affinity-types", types.YChild{"NewStyleAffinityAffinityTypes", &autoBackupAttribute.NewStyleAffinityAffinityTypes})
    autoBackupAttribute.EntityData.Leafs = types.NewOrderedMap()
    autoBackupAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", autoBackupAttribute.AttributeSetName})
    autoBackupAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autoBackupAttribute.Enable})
    autoBackupAttribute.EntityData.Leafs.Append("record-route", types.YLeaf{"RecordRoute", autoBackupAttribute.RecordRoute})

    autoBackupAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(autoBackupAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_SignalledName
// Signalled name
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_SignalledName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Signalled name. The type is string.
    Name interface{}

    // Source address or name. The type is MplsTeSigNameOption.
    SourceType interface{}

    // Protected-interface address or name. The type is MplsTeSigNameOption.
    ProtectedInterfaceType interface{}

    // Set if merge-point address is to be appended. The type is bool.
    MpAddress interface{}
}

func (signalledName *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_SignalledName) GetEntityData() *types.CommonEntityData {
    signalledName.EntityData.YFilter = signalledName.YFilter
    signalledName.EntityData.YangName = "signalled-name"
    signalledName.EntityData.BundleName = "cisco_ios_xr"
    signalledName.EntityData.ParentYangName = "auto-backup-attribute"
    signalledName.EntityData.SegmentPath = "signalled-name"
    signalledName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalledName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalledName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalledName.EntityData.Children = types.NewOrderedMap()
    signalledName.EntityData.Leafs = types.NewOrderedMap()
    signalledName.EntityData.Leafs.Append("name", types.YLeaf{"Name", signalledName.Name})
    signalledName.EntityData.Leafs.Append("source-type", types.YLeaf{"SourceType", signalledName.SourceType})
    signalledName.EntityData.Leafs.Append("protected-interface-type", types.YLeaf{"ProtectedInterfaceType", signalledName.ProtectedInterfaceType})
    signalledName.EntityData.Leafs.Append("mp-address", types.YLeaf{"MpAddress", signalledName.MpAddress})

    signalledName.EntityData.YListKeys = []string {}

    return &(signalledName.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AutoBackupLogging
// Log tunnel LSP messages
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AutoBackupLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log tunnel messages for bandwidth change. The type is interface{}.
    BandwidthChangeMessage interface{}

    // Log tunnel reoptimization attempts messages. The type is interface{}.
    ReoptimizeAttemptsMessage interface{}

    // Log tunnel state messages. The type is interface{}.
    StateMessage interface{}

    // Log tunnel reoptimized messages. The type is interface{}.
    ReoptimizedMessage interface{}
}

func (autoBackupLogging *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AutoBackupLogging) GetEntityData() *types.CommonEntityData {
    autoBackupLogging.EntityData.YFilter = autoBackupLogging.YFilter
    autoBackupLogging.EntityData.YangName = "auto-backup-logging"
    autoBackupLogging.EntityData.BundleName = "cisco_ios_xr"
    autoBackupLogging.EntityData.ParentYangName = "auto-backup-attribute"
    autoBackupLogging.EntityData.SegmentPath = "auto-backup-logging"
    autoBackupLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoBackupLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoBackupLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoBackupLogging.EntityData.Children = types.NewOrderedMap()
    autoBackupLogging.EntityData.Leafs = types.NewOrderedMap()
    autoBackupLogging.EntityData.Leafs.Append("bandwidth-change-message", types.YLeaf{"BandwidthChangeMessage", autoBackupLogging.BandwidthChangeMessage})
    autoBackupLogging.EntityData.Leafs.Append("reoptimize-attempts-message", types.YLeaf{"ReoptimizeAttemptsMessage", autoBackupLogging.ReoptimizeAttemptsMessage})
    autoBackupLogging.EntityData.Leafs.Append("state-message", types.YLeaf{"StateMessage", autoBackupLogging.StateMessage})
    autoBackupLogging.EntityData.Leafs.Append("reoptimized-message", types.YLeaf{"ReoptimizedMessage", autoBackupLogging.ReoptimizedMessage})

    autoBackupLogging.EntityData.YListKeys = []string {}

    return &(autoBackupLogging.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "auto-backup-attribute"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", priority.SetupPriority})
    priority.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", priority.HoldPriority})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AffinityMask
// Set the affinity flags and mask
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AffinityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Mask interface{}
}

func (affinityMask *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_AffinityMask) GetEntityData() *types.CommonEntityData {
    affinityMask.EntityData.YFilter = affinityMask.YFilter
    affinityMask.EntityData.YangName = "affinity-mask"
    affinityMask.EntityData.BundleName = "cisco_ios_xr"
    affinityMask.EntityData.ParentYangName = "auto-backup-attribute"
    affinityMask.EntityData.SegmentPath = "affinity-mask"
    affinityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMask.EntityData.Children = types.NewOrderedMap()
    affinityMask.EntityData.Leafs = types.NewOrderedMap()
    affinityMask.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", affinityMask.Affinity})
    affinityMask.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", affinityMask.Mask})

    affinityMask.EntityData.YListKeys = []string {}

    return &(affinityMask.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path selection. The type is interface{}.
    Enable interface{}
}

func (pathSelection *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "auto-backup-attribute"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSelection.Enable})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PolicyClasses
// Policy classes for PBTS
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PolicyClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of Policy class. The type is slice of interface{} with range: 1..8.
    PolicyClass []interface{}
}

func (policyClasses *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_PolicyClasses) GetEntityData() *types.CommonEntityData {
    policyClasses.EntityData.YFilter = policyClasses.YFilter
    policyClasses.EntityData.YangName = "policy-classes"
    policyClasses.EntityData.BundleName = "cisco_ios_xr"
    policyClasses.EntityData.ParentYangName = "auto-backup-attribute"
    policyClasses.EntityData.SegmentPath = "policy-classes"
    policyClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyClasses.EntityData.Children = types.NewOrderedMap()
    policyClasses.EntityData.Leafs = types.NewOrderedMap()
    policyClasses.EntityData.Leafs.Append("policy-class", types.YLeaf{"PolicyClass", policyClasses.PolicyClass})

    policyClasses.EntityData.YListKeys = []string {}

    return &(policyClasses.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes
// Tunnel new style affinity attributes table
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType.
    NewStyleAffinityAffinityType []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1.
    NewStyleAffinityAffinityTypeAffinity1 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2.
    NewStyleAffinityAffinityTypeAffinity1Affinity2 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 []*MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
}

func (newStyleAffinityAffinityTypes *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypes.EntityData.YFilter = newStyleAffinityAffinityTypes.YFilter
    newStyleAffinityAffinityTypes.EntityData.YangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypes.EntityData.ParentYangName = "auto-backup-attribute"
    newStyleAffinityAffinityTypes.EntityData.SegmentPath = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypes.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type", types.YChild{"NewStyleAffinityAffinityType", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]), types.YChild{"NewStyleAffinityAffinityType", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1", types.YChild{"NewStyleAffinityAffinityTypeAffinity1", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Leafs = types.NewOrderedMap()

    newStyleAffinityAffinityTypes.EntityData.YListKeys = []string {}

    return &(newStyleAffinityAffinityTypes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}
}

func (newStyleAffinityAffinityType *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityType.EntityData.YFilter = newStyleAffinityAffinityType.YFilter
    newStyleAffinityAffinityType.EntityData.YangName = "new-style-affinity-affinity-type"
    newStyleAffinityAffinityType.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityType.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityType.EntityData.SegmentPath = "new-style-affinity-affinity-type" + types.AddKeyToken(newStyleAffinityAffinityType.AffinityType, "affinity-type")
    newStyleAffinityAffinityType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityType.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityType.AffinityType})

    newStyleAffinityAffinityType.EntityData.YListKeys = []string {"AffinityType"}

    return &(newStyleAffinityAffinityType.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1.YFilter
    newStyleAffinityAffinityTypeAffinity1.EntityData.YangName = "new-style-affinity-affinity-type-affinity1"
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.Affinity1, "affinity1")
    newStyleAffinityAffinityTypeAffinity1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1.AffinityType})
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1.Affinity1})

    newStyleAffinityAffinityTypeAffinity1.EntityData.YListKeys = []string {"AffinityType", "Affinity1"}

    return &(newStyleAffinityAffinityTypeAffinity1.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2, "affinity2")
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2})

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3, "affinity3")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4, "affinity4")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5, "affinity5")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6, "affinity6")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7, "affinity7")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8, "affinity8")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9, "affinity9")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}

    // This attribute is a key. The name of the tenth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity10 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 *MplsTe_GlobalAttributes_AttributeSet_AutoBackupAttributes_AutoBackupAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9, "affinity9") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10, "affinity10")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity10", types.YLeaf{"Affinity10", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9", "Affinity10"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes
// OTN Path Protection Attributes table
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OTN Path Protection Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute.
    OtnPpAttribute []*MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute
}

func (otnPpAttributes *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes) GetEntityData() *types.CommonEntityData {
    otnPpAttributes.EntityData.YFilter = otnPpAttributes.YFilter
    otnPpAttributes.EntityData.YangName = "otn-pp-attributes"
    otnPpAttributes.EntityData.BundleName = "cisco_ios_xr"
    otnPpAttributes.EntityData.ParentYangName = "attribute-set"
    otnPpAttributes.EntityData.SegmentPath = "otn-pp-attributes"
    otnPpAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnPpAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnPpAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnPpAttributes.EntityData.Children = types.NewOrderedMap()
    otnPpAttributes.EntityData.Children.Append("otn-pp-attribute", types.YChild{"OtnPpAttribute", nil})
    for i := range otnPpAttributes.OtnPpAttribute {
        otnPpAttributes.EntityData.Children.Append(types.GetSegmentPath(otnPpAttributes.OtnPpAttribute[i]), types.YChild{"OtnPpAttribute", otnPpAttributes.OtnPpAttribute[i]})
    }
    otnPpAttributes.EntityData.Leafs = types.NewOrderedMap()

    otnPpAttributes.EntityData.YListKeys = []string {}

    return &(otnPpAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute
// OTN Path Protection Attribute
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // The APS protecion mode. The type is MplsTeOtnApsProtectionMode.
    ApsProtectionMode interface{}

    // The APS restoration style. The type is MplsTeOtnApsRestorationStyle.
    ApsRestorationStyle interface{}

    // The APS protecion type. The type is MplsTeOtnApsProtection.
    ApsProtectionType interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Specify APS revert schedule.
    RevertScheduleNames MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames

    // Sub-network connection mode.
    SubNetworkConnectionMode MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_SubNetworkConnectionMode

    // Timers.
    Timers MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_Timers

    // Configure path selection properties.
    PathSelection MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_PathSelection
}

func (otnPpAttribute *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute) GetEntityData() *types.CommonEntityData {
    otnPpAttribute.EntityData.YFilter = otnPpAttribute.YFilter
    otnPpAttribute.EntityData.YangName = "otn-pp-attribute"
    otnPpAttribute.EntityData.BundleName = "cisco_ios_xr"
    otnPpAttribute.EntityData.ParentYangName = "otn-pp-attributes"
    otnPpAttribute.EntityData.SegmentPath = "otn-pp-attribute" + types.AddKeyToken(otnPpAttribute.AttributeSetName, "attribute-set-name")
    otnPpAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otnPpAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otnPpAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otnPpAttribute.EntityData.Children = types.NewOrderedMap()
    otnPpAttribute.EntityData.Children.Append("revert-schedule-names", types.YChild{"RevertScheduleNames", &otnPpAttribute.RevertScheduleNames})
    otnPpAttribute.EntityData.Children.Append("sub-network-connection-mode", types.YChild{"SubNetworkConnectionMode", &otnPpAttribute.SubNetworkConnectionMode})
    otnPpAttribute.EntityData.Children.Append("timers", types.YChild{"Timers", &otnPpAttribute.Timers})
    otnPpAttribute.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &otnPpAttribute.PathSelection})
    otnPpAttribute.EntityData.Leafs = types.NewOrderedMap()
    otnPpAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", otnPpAttribute.AttributeSetName})
    otnPpAttribute.EntityData.Leafs.Append("aps-protection-mode", types.YLeaf{"ApsProtectionMode", otnPpAttribute.ApsProtectionMode})
    otnPpAttribute.EntityData.Leafs.Append("aps-restoration-style", types.YLeaf{"ApsRestorationStyle", otnPpAttribute.ApsRestorationStyle})
    otnPpAttribute.EntityData.Leafs.Append("aps-protection-type", types.YLeaf{"ApsProtectionType", otnPpAttribute.ApsProtectionType})
    otnPpAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", otnPpAttribute.Enable})

    otnPpAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(otnPpAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames
// Specify APS revert schedule
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name Identifier for revert schedule. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName.
    RevertScheduleName []*MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName
}

func (revertScheduleNames *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames) GetEntityData() *types.CommonEntityData {
    revertScheduleNames.EntityData.YFilter = revertScheduleNames.YFilter
    revertScheduleNames.EntityData.YangName = "revert-schedule-names"
    revertScheduleNames.EntityData.BundleName = "cisco_ios_xr"
    revertScheduleNames.EntityData.ParentYangName = "otn-pp-attribute"
    revertScheduleNames.EntityData.SegmentPath = "revert-schedule-names"
    revertScheduleNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    revertScheduleNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    revertScheduleNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    revertScheduleNames.EntityData.Children = types.NewOrderedMap()
    revertScheduleNames.EntityData.Children.Append("revert-schedule-name", types.YChild{"RevertScheduleName", nil})
    for i := range revertScheduleNames.RevertScheduleName {
        revertScheduleNames.EntityData.Children.Append(types.GetSegmentPath(revertScheduleNames.RevertScheduleName[i]), types.YChild{"RevertScheduleName", revertScheduleNames.RevertScheduleName[i]})
    }
    revertScheduleNames.EntityData.Leafs = types.NewOrderedMap()

    revertScheduleNames.EntityData.YListKeys = []string {}

    return &(revertScheduleNames.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName
// Name Identifier for revert schedule
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Enter 64 characters for revert schedule name. The
    // type is string with length: 1..254.
    ScheduleName interface{}

    // Revert Schedule Max tries. The type is interface{} with range: 1..2016.
    RevertScheduleMaxTries interface{}

    // Schedule name enable object. The type is interface{}.
    SchNameEnable interface{}

    // Frequency set as Once, Daily, Weekly. The type is interface{} with range:
    // 1..3.
    RevertScheduleFrequency interface{}

    // Set duration in format hh:mm.
    ScheduleDuration MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDuration

    // Set date in format hh:mm MMM DD YYYY.
    ScheduleDate MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDate
}

func (revertScheduleName *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName) GetEntityData() *types.CommonEntityData {
    revertScheduleName.EntityData.YFilter = revertScheduleName.YFilter
    revertScheduleName.EntityData.YangName = "revert-schedule-name"
    revertScheduleName.EntityData.BundleName = "cisco_ios_xr"
    revertScheduleName.EntityData.ParentYangName = "revert-schedule-names"
    revertScheduleName.EntityData.SegmentPath = "revert-schedule-name" + types.AddKeyToken(revertScheduleName.ScheduleName, "schedule-name")
    revertScheduleName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    revertScheduleName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    revertScheduleName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    revertScheduleName.EntityData.Children = types.NewOrderedMap()
    revertScheduleName.EntityData.Children.Append("schedule-duration", types.YChild{"ScheduleDuration", &revertScheduleName.ScheduleDuration})
    revertScheduleName.EntityData.Children.Append("schedule-date", types.YChild{"ScheduleDate", &revertScheduleName.ScheduleDate})
    revertScheduleName.EntityData.Leafs = types.NewOrderedMap()
    revertScheduleName.EntityData.Leafs.Append("schedule-name", types.YLeaf{"ScheduleName", revertScheduleName.ScheduleName})
    revertScheduleName.EntityData.Leafs.Append("revert-schedule-max-tries", types.YLeaf{"RevertScheduleMaxTries", revertScheduleName.RevertScheduleMaxTries})
    revertScheduleName.EntityData.Leafs.Append("sch-name-enable", types.YLeaf{"SchNameEnable", revertScheduleName.SchNameEnable})
    revertScheduleName.EntityData.Leafs.Append("revert-schedule-frequency", types.YLeaf{"RevertScheduleFrequency", revertScheduleName.RevertScheduleFrequency})

    revertScheduleName.EntityData.YListKeys = []string {"ScheduleName"}

    return &(revertScheduleName.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDuration
// Set duration in format hh:mm
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Hour of day. The type is interface{} with range: 0..167. This attribute is
    // mandatory.
    Hour interface{}

    // Minute of the hour. The type is interface{} with range: 0..59. This
    // attribute is mandatory.
    Minutes interface{}
}

func (scheduleDuration *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDuration) GetEntityData() *types.CommonEntityData {
    scheduleDuration.EntityData.YFilter = scheduleDuration.YFilter
    scheduleDuration.EntityData.YangName = "schedule-duration"
    scheduleDuration.EntityData.BundleName = "cisco_ios_xr"
    scheduleDuration.EntityData.ParentYangName = "revert-schedule-name"
    scheduleDuration.EntityData.SegmentPath = "schedule-duration"
    scheduleDuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scheduleDuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scheduleDuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scheduleDuration.EntityData.Children = types.NewOrderedMap()
    scheduleDuration.EntityData.Leafs = types.NewOrderedMap()
    scheduleDuration.EntityData.Leafs.Append("hour", types.YLeaf{"Hour", scheduleDuration.Hour})
    scheduleDuration.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", scheduleDuration.Minutes})

    scheduleDuration.EntityData.YListKeys = []string {}

    return &(scheduleDuration.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDate
// Set date in format hh:mm MMM DD YYYY
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Hour of day. The type is interface{} with range: 0..23. This attribute is
    // mandatory.
    Hour interface{}

    // Minute of the hour. The type is interface{} with range: 0..59. This
    // attribute is mandatory.
    Minutes interface{}

    // Month of the year. The type is interface{} with range: 0..11. This
    // attribute is mandatory.
    Month interface{}

    // Day of the month. The type is interface{} with range: 1..31. This attribute
    // is mandatory.
    Day interface{}

    // Year. The type is interface{} with range: 2015..2035. This attribute is
    // mandatory.
    Year interface{}
}

func (scheduleDate *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_RevertScheduleNames_RevertScheduleName_ScheduleDate) GetEntityData() *types.CommonEntityData {
    scheduleDate.EntityData.YFilter = scheduleDate.YFilter
    scheduleDate.EntityData.YangName = "schedule-date"
    scheduleDate.EntityData.BundleName = "cisco_ios_xr"
    scheduleDate.EntityData.ParentYangName = "revert-schedule-name"
    scheduleDate.EntityData.SegmentPath = "schedule-date"
    scheduleDate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    scheduleDate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    scheduleDate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    scheduleDate.EntityData.Children = types.NewOrderedMap()
    scheduleDate.EntityData.Leafs = types.NewOrderedMap()
    scheduleDate.EntityData.Leafs.Append("hour", types.YLeaf{"Hour", scheduleDate.Hour})
    scheduleDate.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", scheduleDate.Minutes})
    scheduleDate.EntityData.Leafs.Append("month", types.YLeaf{"Month", scheduleDate.Month})
    scheduleDate.EntityData.Leafs.Append("day", types.YLeaf{"Day", scheduleDate.Day})
    scheduleDate.EntityData.Leafs.Append("year", types.YLeaf{"Year", scheduleDate.Year})

    scheduleDate.EntityData.YListKeys = []string {}

    return &(scheduleDate.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_SubNetworkConnectionMode
// Sub-network connection mode
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_SubNetworkConnectionMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The sub-network connection mode. The type is MplsTeOtnSncMode.
    ConnectionMode interface{}

    // Tandem Connection Monitoring ID for the interface. The type is interface{}
    // with range: 1..6.
    ConnectionMonitoringMode interface{}
}

func (subNetworkConnectionMode *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_SubNetworkConnectionMode) GetEntityData() *types.CommonEntityData {
    subNetworkConnectionMode.EntityData.YFilter = subNetworkConnectionMode.YFilter
    subNetworkConnectionMode.EntityData.YangName = "sub-network-connection-mode"
    subNetworkConnectionMode.EntityData.BundleName = "cisco_ios_xr"
    subNetworkConnectionMode.EntityData.ParentYangName = "otn-pp-attribute"
    subNetworkConnectionMode.EntityData.SegmentPath = "sub-network-connection-mode"
    subNetworkConnectionMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subNetworkConnectionMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subNetworkConnectionMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subNetworkConnectionMode.EntityData.Children = types.NewOrderedMap()
    subNetworkConnectionMode.EntityData.Leafs = types.NewOrderedMap()
    subNetworkConnectionMode.EntityData.Leafs.Append("connection-mode", types.YLeaf{"ConnectionMode", subNetworkConnectionMode.ConnectionMode})
    subNetworkConnectionMode.EntityData.Leafs.Append("connection-monitoring-mode", types.YLeaf{"ConnectionMonitoringMode", subNetworkConnectionMode.ConnectionMonitoringMode})

    subNetworkConnectionMode.EntityData.YListKeys = []string {}

    return &(subNetworkConnectionMode.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_Timers
// Timers
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // G.709 OTN path protection wait to restore timer in seconds. The type is
    // interface{} with range: 0..720. Units are second.
    ApsWaitToRestore interface{}

    // G.709 OTN path protection hold-off timer in milliseconds. The type is
    // interface{} with range: 100..10000. Units are millisecond.
    ApsHoldOff interface{}
}

func (timers *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "otn-pp-attribute"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Leafs = types.NewOrderedMap()
    timers.EntityData.Leafs.Append("aps-wait-to-restore", types.YLeaf{"ApsWaitToRestore", timers.ApsWaitToRestore})
    timers.EntityData.Leafs.Append("aps-hold-off", types.YLeaf{"ApsHoldOff", timers.ApsHoldOff})

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_PathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path selection. The type is interface{}.
    Enable interface{}
}

func (pathSelection *MplsTe_GlobalAttributes_AttributeSet_OtnPpAttributes_OtnPpAttribute_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "otn-pp-attribute"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSelection.Enable})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes
// Auto-mesh Tunnel AttributeSets Table
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto-mesh Tunnel Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute.
    AutoMeshAttribute []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute
}

func (autoMeshAttributes *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes) GetEntityData() *types.CommonEntityData {
    autoMeshAttributes.EntityData.YFilter = autoMeshAttributes.YFilter
    autoMeshAttributes.EntityData.YangName = "auto-mesh-attributes"
    autoMeshAttributes.EntityData.BundleName = "cisco_ios_xr"
    autoMeshAttributes.EntityData.ParentYangName = "attribute-set"
    autoMeshAttributes.EntityData.SegmentPath = "auto-mesh-attributes"
    autoMeshAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMeshAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMeshAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMeshAttributes.EntityData.Children = types.NewOrderedMap()
    autoMeshAttributes.EntityData.Children.Append("auto-mesh-attribute", types.YChild{"AutoMeshAttribute", nil})
    for i := range autoMeshAttributes.AutoMeshAttribute {
        autoMeshAttributes.EntityData.Children.Append(types.GetSegmentPath(autoMeshAttributes.AutoMeshAttribute[i]), types.YChild{"AutoMeshAttribute", autoMeshAttributes.AutoMeshAttribute[i]})
    }
    autoMeshAttributes.EntityData.Leafs = types.NewOrderedMap()

    autoMeshAttributes.EntityData.YListKeys = []string {}

    return &(autoMeshAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute
// Auto-mesh Tunnel Attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // Enable autoroute announce. The type is interface{}.
    AutorouteAnnounce interface{}

    // The bandwidth of the interface in kbps. The type is interface{} with range:
    // 0..4294967295. Units are kbit/s.
    InterfaceBandwidth interface{}

    // Forward class value. The type is interface{} with range: 0..7.
    ForwardClass interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Record the route used by the tunnel. The type is interface{}.
    RecordRoute interface{}

    // Enable bandwidth collection only, no auto-bw adjustment. The type is
    // interface{}.
    CollectionOnly interface{}

    // Enable the soft-preemption feature on the tunnel. The type is interface{}.
    SoftPreemption interface{}

    // Tunnel loadsharing metric. The type is interface{} with range:
    // 1..4294967295.
    LoadShare interface{}

    // Log tunnel LSP messages.
    AutoMeshLogging MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AutoMeshLogging

    // Tunnel Setup and Hold Priorities.
    Priority MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Priority

    // Set the affinity flags and mask.
    AffinityMask MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AffinityMask

    // Tunnel bandwidth requirement.
    Bandwidth MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Bandwidth

    // Configure path selection properties.
    PathSelection MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PathSelection

    // Policy classes for PBTS.
    PolicyClasses MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PolicyClasses

    // Tunnel new style affinity attributes table.
    NewStyleAffinityAffinityTypes MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes

    // Specify MPLS tunnel can be fast-rerouted.
    FastReroute MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_FastReroute
}

func (autoMeshAttribute *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute) GetEntityData() *types.CommonEntityData {
    autoMeshAttribute.EntityData.YFilter = autoMeshAttribute.YFilter
    autoMeshAttribute.EntityData.YangName = "auto-mesh-attribute"
    autoMeshAttribute.EntityData.BundleName = "cisco_ios_xr"
    autoMeshAttribute.EntityData.ParentYangName = "auto-mesh-attributes"
    autoMeshAttribute.EntityData.SegmentPath = "auto-mesh-attribute" + types.AddKeyToken(autoMeshAttribute.AttributeSetName, "attribute-set-name")
    autoMeshAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMeshAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMeshAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMeshAttribute.EntityData.Children = types.NewOrderedMap()
    autoMeshAttribute.EntityData.Children.Append("auto-mesh-logging", types.YChild{"AutoMeshLogging", &autoMeshAttribute.AutoMeshLogging})
    autoMeshAttribute.EntityData.Children.Append("priority", types.YChild{"Priority", &autoMeshAttribute.Priority})
    autoMeshAttribute.EntityData.Children.Append("affinity-mask", types.YChild{"AffinityMask", &autoMeshAttribute.AffinityMask})
    autoMeshAttribute.EntityData.Children.Append("bandwidth", types.YChild{"Bandwidth", &autoMeshAttribute.Bandwidth})
    autoMeshAttribute.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &autoMeshAttribute.PathSelection})
    autoMeshAttribute.EntityData.Children.Append("policy-classes", types.YChild{"PolicyClasses", &autoMeshAttribute.PolicyClasses})
    autoMeshAttribute.EntityData.Children.Append("new-style-affinity-affinity-types", types.YChild{"NewStyleAffinityAffinityTypes", &autoMeshAttribute.NewStyleAffinityAffinityTypes})
    autoMeshAttribute.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &autoMeshAttribute.FastReroute})
    autoMeshAttribute.EntityData.Leafs = types.NewOrderedMap()
    autoMeshAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", autoMeshAttribute.AttributeSetName})
    autoMeshAttribute.EntityData.Leafs.Append("autoroute-announce", types.YLeaf{"AutorouteAnnounce", autoMeshAttribute.AutorouteAnnounce})
    autoMeshAttribute.EntityData.Leafs.Append("interface-bandwidth", types.YLeaf{"InterfaceBandwidth", autoMeshAttribute.InterfaceBandwidth})
    autoMeshAttribute.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", autoMeshAttribute.ForwardClass})
    autoMeshAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", autoMeshAttribute.Enable})
    autoMeshAttribute.EntityData.Leafs.Append("record-route", types.YLeaf{"RecordRoute", autoMeshAttribute.RecordRoute})
    autoMeshAttribute.EntityData.Leafs.Append("collection-only", types.YLeaf{"CollectionOnly", autoMeshAttribute.CollectionOnly})
    autoMeshAttribute.EntityData.Leafs.Append("soft-preemption", types.YLeaf{"SoftPreemption", autoMeshAttribute.SoftPreemption})
    autoMeshAttribute.EntityData.Leafs.Append("load-share", types.YLeaf{"LoadShare", autoMeshAttribute.LoadShare})

    autoMeshAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(autoMeshAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AutoMeshLogging
// Log tunnel LSP messages
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AutoMeshLogging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log tunnel messages for bandwidth change. The type is interface{}.
    BandwidthChangeMessage interface{}

    // Log tunnel reoptimization attempts messages. The type is interface{}.
    ReoptimizeAttemptsMessage interface{}

    // Log tunnel rereoute messages. The type is interface{}.
    RerouteMesssage interface{}

    // Log tunnel state messages. The type is interface{}.
    StateMessage interface{}

    // Log tunnel messages for insufficient bandwidth. The type is interface{}.
    InsufficientBwMessage interface{}

    // Log tunnel reoptimized messages. The type is interface{}.
    ReoptimizedMessage interface{}

    // Enable logging for path-calculation failures. The type is interface{}.
    PcalcFailureMessage interface{}
}

func (autoMeshLogging *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AutoMeshLogging) GetEntityData() *types.CommonEntityData {
    autoMeshLogging.EntityData.YFilter = autoMeshLogging.YFilter
    autoMeshLogging.EntityData.YangName = "auto-mesh-logging"
    autoMeshLogging.EntityData.BundleName = "cisco_ios_xr"
    autoMeshLogging.EntityData.ParentYangName = "auto-mesh-attribute"
    autoMeshLogging.EntityData.SegmentPath = "auto-mesh-logging"
    autoMeshLogging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoMeshLogging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoMeshLogging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoMeshLogging.EntityData.Children = types.NewOrderedMap()
    autoMeshLogging.EntityData.Leafs = types.NewOrderedMap()
    autoMeshLogging.EntityData.Leafs.Append("bandwidth-change-message", types.YLeaf{"BandwidthChangeMessage", autoMeshLogging.BandwidthChangeMessage})
    autoMeshLogging.EntityData.Leafs.Append("reoptimize-attempts-message", types.YLeaf{"ReoptimizeAttemptsMessage", autoMeshLogging.ReoptimizeAttemptsMessage})
    autoMeshLogging.EntityData.Leafs.Append("reroute-messsage", types.YLeaf{"RerouteMesssage", autoMeshLogging.RerouteMesssage})
    autoMeshLogging.EntityData.Leafs.Append("state-message", types.YLeaf{"StateMessage", autoMeshLogging.StateMessage})
    autoMeshLogging.EntityData.Leafs.Append("insufficient-bw-message", types.YLeaf{"InsufficientBwMessage", autoMeshLogging.InsufficientBwMessage})
    autoMeshLogging.EntityData.Leafs.Append("reoptimized-message", types.YLeaf{"ReoptimizedMessage", autoMeshLogging.ReoptimizedMessage})
    autoMeshLogging.EntityData.Leafs.Append("pcalc-failure-message", types.YLeaf{"PcalcFailureMessage", autoMeshLogging.PcalcFailureMessage})

    autoMeshLogging.EntityData.YListKeys = []string {}

    return &(autoMeshLogging.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Priority
// Tunnel Setup and Hold Priorities
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Priority struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Setup Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..7. This attribute is
    // mandatory.
    HoldPriority interface{}
}

func (priority *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Priority) GetEntityData() *types.CommonEntityData {
    priority.EntityData.YFilter = priority.YFilter
    priority.EntityData.YangName = "priority"
    priority.EntityData.BundleName = "cisco_ios_xr"
    priority.EntityData.ParentYangName = "auto-mesh-attribute"
    priority.EntityData.SegmentPath = "priority"
    priority.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    priority.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    priority.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    priority.EntityData.Children = types.NewOrderedMap()
    priority.EntityData.Leafs = types.NewOrderedMap()
    priority.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", priority.SetupPriority})
    priority.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", priority.HoldPriority})

    priority.EntityData.YListKeys = []string {}

    return &(priority.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AffinityMask
// Set the affinity flags and mask
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AffinityMask struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}. This
    // attribute is mandatory.
    Mask interface{}
}

func (affinityMask *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_AffinityMask) GetEntityData() *types.CommonEntityData {
    affinityMask.EntityData.YFilter = affinityMask.YFilter
    affinityMask.EntityData.YangName = "affinity-mask"
    affinityMask.EntityData.BundleName = "cisco_ios_xr"
    affinityMask.EntityData.ParentYangName = "auto-mesh-attribute"
    affinityMask.EntityData.SegmentPath = "affinity-mask"
    affinityMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMask.EntityData.Children = types.NewOrderedMap()
    affinityMask.EntityData.Leafs = types.NewOrderedMap()
    affinityMask.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", affinityMask.Affinity})
    affinityMask.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", affinityMask.Mask})

    affinityMask.EntityData.YListKeys = []string {}

    return &(affinityMask.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Bandwidth
// Tunnel bandwidth requirement
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Bandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DSTE-standard flag. The type is MplsTeBandwidthDste. This attribute is
    // mandatory.
    DsteType interface{}

    // Class type for the bandwidth allocation. The type is interface{} with
    // range: 0..1. This attribute is mandatory.
    ClassOrPoolType interface{}

    // The value of the bandwidth reserved by this tunnel in kbps. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory. Units
    // are kbit/s.
    Bandwidth interface{}
}

func (bandwidth *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_Bandwidth) GetEntityData() *types.CommonEntityData {
    bandwidth.EntityData.YFilter = bandwidth.YFilter
    bandwidth.EntityData.YangName = "bandwidth"
    bandwidth.EntityData.BundleName = "cisco_ios_xr"
    bandwidth.EntityData.ParentYangName = "auto-mesh-attribute"
    bandwidth.EntityData.SegmentPath = "bandwidth"
    bandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidth.EntityData.Children = types.NewOrderedMap()
    bandwidth.EntityData.Leafs = types.NewOrderedMap()
    bandwidth.EntityData.Leafs.Append("dste-type", types.YLeaf{"DsteType", bandwidth.DsteType})
    bandwidth.EntityData.Leafs.Append("class-or-pool-type", types.YLeaf{"ClassOrPoolType", bandwidth.ClassOrPoolType})
    bandwidth.EntityData.Leafs.Append("bandwidth", types.YLeaf{"Bandwidth", bandwidth.Bandwidth})

    bandwidth.EntityData.YListKeys = []string {}

    return &(bandwidth.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path selection. The type is interface{}.
    Enable interface{}
}

func (pathSelection *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "auto-mesh-attribute"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSelection.Enable})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PolicyClasses
// Policy classes for PBTS
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PolicyClasses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of Policy class. The type is slice of interface{} with range: 1..8.
    PolicyClass []interface{}
}

func (policyClasses *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_PolicyClasses) GetEntityData() *types.CommonEntityData {
    policyClasses.EntityData.YFilter = policyClasses.YFilter
    policyClasses.EntityData.YangName = "policy-classes"
    policyClasses.EntityData.BundleName = "cisco_ios_xr"
    policyClasses.EntityData.ParentYangName = "auto-mesh-attribute"
    policyClasses.EntityData.SegmentPath = "policy-classes"
    policyClasses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyClasses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyClasses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyClasses.EntityData.Children = types.NewOrderedMap()
    policyClasses.EntityData.Leafs = types.NewOrderedMap()
    policyClasses.EntityData.Leafs.Append("policy-class", types.YLeaf{"PolicyClass", policyClasses.PolicyClass})

    policyClasses.EntityData.YListKeys = []string {}

    return &(policyClasses.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes
// Tunnel new style affinity attributes table
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType.
    NewStyleAffinityAffinityType []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1.
    NewStyleAffinityAffinityTypeAffinity1 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2.
    NewStyleAffinityAffinityTypeAffinity1Affinity2 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9

    // Tunnel new style affinity attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.
    NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 []*MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
}

func (newStyleAffinityAffinityTypes *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypes.EntityData.YFilter = newStyleAffinityAffinityTypes.YFilter
    newStyleAffinityAffinityTypes.EntityData.YangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypes.EntityData.ParentYangName = "auto-mesh-attribute"
    newStyleAffinityAffinityTypes.EntityData.SegmentPath = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypes.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type", types.YChild{"NewStyleAffinityAffinityType", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]), types.YChild{"NewStyleAffinityAffinityType", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityType[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1", types.YChild{"NewStyleAffinityAffinityTypeAffinity1", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Children.Append("new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10", types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", nil})
    for i := range newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 {
        newStyleAffinityAffinityTypes.EntityData.Children.Append(types.GetSegmentPath(newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]), types.YChild{"NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10", newStyleAffinityAffinityTypes.NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10[i]})
    }
    newStyleAffinityAffinityTypes.EntityData.Leafs = types.NewOrderedMap()

    newStyleAffinityAffinityTypes.EntityData.YListKeys = []string {}

    return &(newStyleAffinityAffinityTypes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}
}

func (newStyleAffinityAffinityType *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityType) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityType.EntityData.YFilter = newStyleAffinityAffinityType.YFilter
    newStyleAffinityAffinityType.EntityData.YangName = "new-style-affinity-affinity-type"
    newStyleAffinityAffinityType.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityType.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityType.EntityData.SegmentPath = "new-style-affinity-affinity-type" + types.AddKeyToken(newStyleAffinityAffinityType.AffinityType, "affinity-type")
    newStyleAffinityAffinityType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityType.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityType.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityType.AffinityType})

    newStyleAffinityAffinityType.EntityData.YListKeys = []string {"AffinityType"}

    return &(newStyleAffinityAffinityType.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1.YFilter
    newStyleAffinityAffinityTypeAffinity1.EntityData.YangName = "new-style-affinity-affinity-type-affinity1"
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1.Affinity1, "affinity1")
    newStyleAffinityAffinityTypeAffinity1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1.AffinityType})
    newStyleAffinityAffinityTypeAffinity1.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1.Affinity1})

    newStyleAffinityAffinityTypeAffinity1.EntityData.YListKeys = []string {"AffinityType", "Affinity1"}

    return &(newStyleAffinityAffinityTypeAffinity1.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2, "affinity2")
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2.Affinity2})

    newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3, "affinity3")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.Affinity3})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4, "affinity4")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.Affinity4})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5, "affinity5")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.Affinity5})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6, "affinity6")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.Affinity6})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7, "affinity7")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.Affinity7})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8, "affinity8")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.Affinity8})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9, "affinity9")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.Affinity9})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10
// Tunnel new style affinity attribute
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of the affinity entry. The type is
    // MplsTeTunnelAffinity.
    AffinityType interface{}

    // This attribute is a key. The name of the first affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity1 interface{}

    // This attribute is a key. The name of the second affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity2 interface{}

    // This attribute is a key. The name of the third affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity3 interface{}

    // This attribute is a key. The name of the fourth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity4 interface{}

    // This attribute is a key. The name of the fifth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity5 interface{}

    // This attribute is a key. The name of the sixth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity6 interface{}

    // This attribute is a key. The name of the seventh affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity7 interface{}

    // This attribute is a key. The name of the eighth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity8 interface{}

    // This attribute is a key. The name of the nineth affinity. The type is
    // string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity9 interface{}

    // This attribute is a key. The name of the tenth affinity. The type is string
    // with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Affinity10 interface{}
}

func (newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10 *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_NewStyleAffinityAffinityTypes_NewStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10) GetEntityData() *types.CommonEntityData {
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YFilter = newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.YFilter
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YangName = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleName = "cisco_ios_xr"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.ParentYangName = "new-style-affinity-affinity-types"
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.SegmentPath = "new-style-affinity-affinity-type-affinity1-affinity2-affinity3-affinity4-affinity5-affinity6-affinity7-affinity8-affinity9-affinity10" + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType, "affinity-type") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1, "affinity1") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2, "affinity2") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3, "affinity3") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4, "affinity4") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5, "affinity5") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6, "affinity6") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7, "affinity7") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8, "affinity8") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9, "affinity9") + types.AddKeyToken(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10, "affinity10")
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Children = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs = types.NewOrderedMap()
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity-type", types.YLeaf{"AffinityType", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.AffinityType})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity1", types.YLeaf{"Affinity1", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity1})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity2", types.YLeaf{"Affinity2", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity2})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity3", types.YLeaf{"Affinity3", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity3})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity4", types.YLeaf{"Affinity4", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity4})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity5", types.YLeaf{"Affinity5", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity5})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity6", types.YLeaf{"Affinity6", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity6})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity7", types.YLeaf{"Affinity7", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity7})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity8", types.YLeaf{"Affinity8", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity8})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity9", types.YLeaf{"Affinity9", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity9})
    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.Leafs.Append("affinity10", types.YLeaf{"Affinity10", newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.Affinity10})

    newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData.YListKeys = []string {"AffinityType", "Affinity1", "Affinity2", "Affinity3", "Affinity4", "Affinity5", "Affinity6", "Affinity7", "Affinity8", "Affinity9", "Affinity10"}

    return &(newStyleAffinityAffinityTypeAffinity1Affinity2Affinity3Affinity4Affinity5Affinity6Affinity7Affinity8Affinity9Affinity10.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_FastReroute
// Specify MPLS tunnel can be fast-rerouted
// This type is a presence type.
type MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Bandwidth Protection. The type is interface{} with range: 0..1. This
    // attribute is mandatory.
    BandwidthProtection interface{}

    // Node Protection. The type is interface{} with range: 0..1. This attribute
    // is mandatory.
    NodeProtection interface{}
}

func (fastReroute *MplsTe_GlobalAttributes_AttributeSet_AutoMeshAttributes_AutoMeshAttribute_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "auto-mesh-attribute"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("bandwidth-protection", types.YLeaf{"BandwidthProtection", fastReroute.BandwidthProtection})
    fastReroute.EntityData.Leafs.Append("node-protection", types.YLeaf{"NodeProtection", fastReroute.NodeProtection})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes
// XRO Tunnel Attributes table
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XRO Attribute. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute.
    XroAttribute []*MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute
}

func (xroAttributes *MplsTe_GlobalAttributes_AttributeSet_XroAttributes) GetEntityData() *types.CommonEntityData {
    xroAttributes.EntityData.YFilter = xroAttributes.YFilter
    xroAttributes.EntityData.YangName = "xro-attributes"
    xroAttributes.EntityData.BundleName = "cisco_ios_xr"
    xroAttributes.EntityData.ParentYangName = "attribute-set"
    xroAttributes.EntityData.SegmentPath = "xro-attributes"
    xroAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xroAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xroAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xroAttributes.EntityData.Children = types.NewOrderedMap()
    xroAttributes.EntityData.Children.Append("xro-attribute", types.YChild{"XroAttribute", nil})
    for i := range xroAttributes.XroAttribute {
        xroAttributes.EntityData.Children.Append(types.GetSegmentPath(xroAttributes.XroAttribute[i]), types.YChild{"XroAttribute", xroAttributes.XroAttribute[i]})
    }
    xroAttributes.EntityData.Leafs = types.NewOrderedMap()

    xroAttributes.EntityData.YListKeys = []string {}

    return &(xroAttributes.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute
// XRO Attribute
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attribute Set Name. The type is string with
    // length: 1..64.
    AttributeSetName interface{}

    // Attribute-set enable object that controls whether this attribute-set is
    // configured or not .This object must be set before other configuration
    // supplied for this attribute-set. The type is interface{}.
    Enable interface{}

    // Path diversity.
    PathDiversity MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity

    // Configure path selection properties.
    PathSelection MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathSelection
}

func (xroAttribute *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute) GetEntityData() *types.CommonEntityData {
    xroAttribute.EntityData.YFilter = xroAttribute.YFilter
    xroAttribute.EntityData.YangName = "xro-attribute"
    xroAttribute.EntityData.BundleName = "cisco_ios_xr"
    xroAttribute.EntityData.ParentYangName = "xro-attributes"
    xroAttribute.EntityData.SegmentPath = "xro-attribute" + types.AddKeyToken(xroAttribute.AttributeSetName, "attribute-set-name")
    xroAttribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xroAttribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xroAttribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xroAttribute.EntityData.Children = types.NewOrderedMap()
    xroAttribute.EntityData.Children.Append("path-diversity", types.YChild{"PathDiversity", &xroAttribute.PathDiversity})
    xroAttribute.EntityData.Children.Append("path-selection", types.YChild{"PathSelection", &xroAttribute.PathSelection})
    xroAttribute.EntityData.Leafs = types.NewOrderedMap()
    xroAttribute.EntityData.Leafs.Append("attribute-set-name", types.YLeaf{"AttributeSetName", xroAttribute.AttributeSetName})
    xroAttribute.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", xroAttribute.Enable})

    xroAttribute.EntityData.YListKeys = []string {"AttributeSetName"}

    return &(xroAttribute.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity
// Path diversity
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG-based path diversity.
    Srlgs MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs

    // LSP-based path diversity.
    Lsp MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp
}

func (pathDiversity *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity) GetEntityData() *types.CommonEntityData {
    pathDiversity.EntityData.YFilter = pathDiversity.YFilter
    pathDiversity.EntityData.YangName = "path-diversity"
    pathDiversity.EntityData.BundleName = "cisco_ios_xr"
    pathDiversity.EntityData.ParentYangName = "xro-attribute"
    pathDiversity.EntityData.SegmentPath = "path-diversity"
    pathDiversity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathDiversity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathDiversity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathDiversity.EntityData.Children = types.NewOrderedMap()
    pathDiversity.EntityData.Children.Append("srlgs", types.YChild{"Srlgs", &pathDiversity.Srlgs})
    pathDiversity.EntityData.Children.Append("lsp", types.YChild{"Lsp", &pathDiversity.Lsp})
    pathDiversity.EntityData.Leafs = types.NewOrderedMap()

    pathDiversity.EntityData.YListKeys = []string {}

    return &(pathDiversity.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs
// SRLG-based path diversity
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG-based path-diversity element. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs_Srlg.
    Srlg []*MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs_Srlg
}

func (srlgs *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs) GetEntityData() *types.CommonEntityData {
    srlgs.EntityData.YFilter = srlgs.YFilter
    srlgs.EntityData.YangName = "srlgs"
    srlgs.EntityData.BundleName = "cisco_ios_xr"
    srlgs.EntityData.ParentYangName = "path-diversity"
    srlgs.EntityData.SegmentPath = "srlgs"
    srlgs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgs.EntityData.Children = types.NewOrderedMap()
    srlgs.EntityData.Children.Append("srlg", types.YChild{"Srlg", nil})
    for i := range srlgs.Srlg {
        srlgs.EntityData.Children.Append(types.GetSegmentPath(srlgs.Srlg[i]), types.YChild{"Srlg", srlgs.Srlg[i]})
    }
    srlgs.EntityData.Leafs = types.NewOrderedMap()

    srlgs.EntityData.YListKeys = []string {}

    return &(srlgs.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs_Srlg
// SRLG-based path-diversity element
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs_Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG. The type is interface{} with range:
    // 0..4294967295.
    Srlg interface{}

    // The diversity conformance requirements. The type is
    // MplsTePathDiversityConformance. This attribute is mandatory.
    Conformance interface{}
}

func (srlg *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Srlgs_Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "cisco_ios_xr"
    srlg.EntityData.ParentYangName = "srlgs"
    srlg.EntityData.SegmentPath = "srlg" + types.AddKeyToken(srlg.Srlg, "srlg")
    srlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Leafs = types.NewOrderedMap()
    srlg.EntityData.Leafs.Append("srlg", types.YLeaf{"Srlg", srlg.Srlg})
    srlg.EntityData.Leafs.Append("conformance", types.YLeaf{"Conformance", srlg.Conformance})

    srlg.EntityData.YListKeys = []string {"Srlg"}

    return &(srlg.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp
// LSP-based path diversity
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FEC LSP-based path diversity.
    Fecs MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs
}

func (lsp *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp) GetEntityData() *types.CommonEntityData {
    lsp.EntityData.YFilter = lsp.YFilter
    lsp.EntityData.YangName = "lsp"
    lsp.EntityData.BundleName = "cisco_ios_xr"
    lsp.EntityData.ParentYangName = "path-diversity"
    lsp.EntityData.SegmentPath = "lsp"
    lsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsp.EntityData.Children = types.NewOrderedMap()
    lsp.EntityData.Children.Append("fecs", types.YChild{"Fecs", &lsp.Fecs})
    lsp.EntityData.Leafs = types.NewOrderedMap()

    lsp.EntityData.YListKeys = []string {}

    return &(lsp.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs
// FEC LSP-based path diversity
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP-based path-diversity, referenced by FEC. The type is slice of
    // MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs_Fec.
    Fec []*MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs_Fec
}

func (fecs *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs) GetEntityData() *types.CommonEntityData {
    fecs.EntityData.YFilter = fecs.YFilter
    fecs.EntityData.YangName = "fecs"
    fecs.EntityData.BundleName = "cisco_ios_xr"
    fecs.EntityData.ParentYangName = "lsp"
    fecs.EntityData.SegmentPath = "fecs"
    fecs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fecs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fecs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fecs.EntityData.Children = types.NewOrderedMap()
    fecs.EntityData.Children.Append("fec", types.YChild{"Fec", nil})
    for i := range fecs.Fec {
        fecs.EntityData.Children.Append(types.GetSegmentPath(fecs.Fec[i]), types.YChild{"Fec", fecs.Fec[i]})
    }
    fecs.EntityData.Leafs = types.NewOrderedMap()

    fecs.EntityData.YListKeys = []string {}

    return &(fecs.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs_Fec
// LSP-based path-diversity, referenced by
// FEC
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs_Fec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Source interface{}

    // This attribute is a key. Destination address. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Destination interface{}

    // This attribute is a key. Tunnel id. The type is interface{} with range:
    // 0..65535.
    TunnelId interface{}

    // This attribute is a key. Extended tunnel-id. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ExtendedTunnelId interface{}

    // This attribute is a key. LSP id. The type is interface{} with range:
    // 0..65535.
    LspId interface{}

    // The diversity conformance requirements. The type is
    // MplsTePathDiversityConformance. This attribute is mandatory.
    Conformance interface{}
}

func (fec *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathDiversity_Lsp_Fecs_Fec) GetEntityData() *types.CommonEntityData {
    fec.EntityData.YFilter = fec.YFilter
    fec.EntityData.YangName = "fec"
    fec.EntityData.BundleName = "cisco_ios_xr"
    fec.EntityData.ParentYangName = "fecs"
    fec.EntityData.SegmentPath = "fec" + types.AddKeyToken(fec.Source, "source") + types.AddKeyToken(fec.Destination, "destination") + types.AddKeyToken(fec.TunnelId, "tunnel-id") + types.AddKeyToken(fec.ExtendedTunnelId, "extended-tunnel-id") + types.AddKeyToken(fec.LspId, "lsp-id")
    fec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fec.EntityData.Children = types.NewOrderedMap()
    fec.EntityData.Leafs = types.NewOrderedMap()
    fec.EntityData.Leafs.Append("source", types.YLeaf{"Source", fec.Source})
    fec.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", fec.Destination})
    fec.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", fec.TunnelId})
    fec.EntityData.Leafs.Append("extended-tunnel-id", types.YLeaf{"ExtendedTunnelId", fec.ExtendedTunnelId})
    fec.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", fec.LspId})
    fec.EntityData.Leafs.Append("conformance", types.YLeaf{"Conformance", fec.Conformance})

    fec.EntityData.YListKeys = []string {"Source", "Destination", "TunnelId", "ExtendedTunnelId", "LspId"}

    return &(fec.EntityData)
}

// MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathSelection
// Configure path selection properties
type MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable path selection. The type is interface{}.
    Enable interface{}
}

func (pathSelection *MplsTe_GlobalAttributes_AttributeSet_XroAttributes_XroAttribute_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "xro-attribute"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pathSelection.Enable})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_BfdOverLsp
// BFD over MPLS TE Global Configurations
type MplsTe_GlobalAttributes_BfdOverLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BFD over LSP Tail Global Configurations.
    Tail MplsTe_GlobalAttributes_BfdOverLsp_Tail

    // BFD over LSP Head Global Configurations.
    Head MplsTe_GlobalAttributes_BfdOverLsp_Head
}

func (bfdOverLsp *MplsTe_GlobalAttributes_BfdOverLsp) GetEntityData() *types.CommonEntityData {
    bfdOverLsp.EntityData.YFilter = bfdOverLsp.YFilter
    bfdOverLsp.EntityData.YangName = "bfd-over-lsp"
    bfdOverLsp.EntityData.BundleName = "cisco_ios_xr"
    bfdOverLsp.EntityData.ParentYangName = "global-attributes"
    bfdOverLsp.EntityData.SegmentPath = "bfd-over-lsp"
    bfdOverLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfdOverLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfdOverLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfdOverLsp.EntityData.Children = types.NewOrderedMap()
    bfdOverLsp.EntityData.Children.Append("tail", types.YChild{"Tail", &bfdOverLsp.Tail})
    bfdOverLsp.EntityData.Children.Append("head", types.YChild{"Head", &bfdOverLsp.Head})
    bfdOverLsp.EntityData.Leafs = types.NewOrderedMap()

    bfdOverLsp.EntityData.YListKeys = []string {}

    return &(bfdOverLsp.EntityData)
}

// MplsTe_GlobalAttributes_BfdOverLsp_Tail
// BFD over LSP Tail Global Configurations
type MplsTe_GlobalAttributes_BfdOverLsp_Tail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify BFD over LSP tail multiplier. The type is interface{} with range:
    // 3..10.
    Multiplier interface{}

    // Specify BFD over LSP tail minimum interval. The type is interface{} with
    // range: 50..30000.
    MinimumInterval interface{}
}

func (tail *MplsTe_GlobalAttributes_BfdOverLsp_Tail) GetEntityData() *types.CommonEntityData {
    tail.EntityData.YFilter = tail.YFilter
    tail.EntityData.YangName = "tail"
    tail.EntityData.BundleName = "cisco_ios_xr"
    tail.EntityData.ParentYangName = "bfd-over-lsp"
    tail.EntityData.SegmentPath = "tail"
    tail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tail.EntityData.Children = types.NewOrderedMap()
    tail.EntityData.Leafs = types.NewOrderedMap()
    tail.EntityData.Leafs.Append("multiplier", types.YLeaf{"Multiplier", tail.Multiplier})
    tail.EntityData.Leafs.Append("minimum-interval", types.YLeaf{"MinimumInterval", tail.MinimumInterval})

    tail.EntityData.YListKeys = []string {}

    return &(tail.EntityData)
}

// MplsTe_GlobalAttributes_BfdOverLsp_Head
// BFD over LSP Head Global Configurations
type MplsTe_GlobalAttributes_BfdOverLsp_Head struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BFD session down reopt timeout. The type is interface{} with range:
    // 120..4294967295.
    ReoptTimeout interface{}

    // Specify BFD session down action. The type is MplsTeBfdSessionDownAction.
    DownAction interface{}
}

func (head *MplsTe_GlobalAttributes_BfdOverLsp_Head) GetEntityData() *types.CommonEntityData {
    head.EntityData.YFilter = head.YFilter
    head.EntityData.YangName = "head"
    head.EntityData.BundleName = "cisco_ios_xr"
    head.EntityData.ParentYangName = "bfd-over-lsp"
    head.EntityData.SegmentPath = "head"
    head.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    head.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    head.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    head.EntityData.Children = types.NewOrderedMap()
    head.EntityData.Leafs = types.NewOrderedMap()
    head.EntityData.Leafs.Append("reopt-timeout", types.YLeaf{"ReoptTimeout", head.ReoptTimeout})
    head.EntityData.Leafs.Append("down-action", types.YLeaf{"DownAction", head.DownAction})

    head.EntityData.YListKeys = []string {}

    return &(head.EntityData)
}

// MplsTe_GlobalAttributes_BandwidthAccounting
// Bandwidth accounting configuration data
type MplsTe_GlobalAttributes_BandwidthAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object sets the sampling interval in seconds for bandwidth accounting.
    // Default to 60 seconds. The type is interface{} with range: 10..600. Units
    // are second. The default value is 60.
    SamplingInterval interface{}

    // This object sets the percentage adjustment factor for the non RSVP-TE
    // bandwidth accounting.  Default is 100%. The type is interface{} with range:
    // 0..200. Units are percentage. The default value is 100.
    AdjustmentFactor interface{}

    // This object enables the bandwidth accounting RSVP-TE sample collection. The
    // type is bool. The default value is false.
    CollectionTypeRsvpTe interface{}

    // This object controls whether BW accounting is enabled. This object must be
    // set before setting any other objects  under the BandwidthAccounting class.
    // The type is interface{}.
    Enable interface{}

    // Bandwidth accounting application configuration data.
    Application MplsTe_GlobalAttributes_BandwidthAccounting_Application

    // This object sets the flooding threshold as percentage of total link
    // bandwidth for bandwidth accounting. Default to 10%, 10%.
    AccountFloodingThreshold MplsTe_GlobalAttributes_BandwidthAccounting_AccountFloodingThreshold
}

func (bandwidthAccounting *MplsTe_GlobalAttributes_BandwidthAccounting) GetEntityData() *types.CommonEntityData {
    bandwidthAccounting.EntityData.YFilter = bandwidthAccounting.YFilter
    bandwidthAccounting.EntityData.YangName = "bandwidth-accounting"
    bandwidthAccounting.EntityData.BundleName = "cisco_ios_xr"
    bandwidthAccounting.EntityData.ParentYangName = "global-attributes"
    bandwidthAccounting.EntityData.SegmentPath = "bandwidth-accounting"
    bandwidthAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bandwidthAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bandwidthAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bandwidthAccounting.EntityData.Children = types.NewOrderedMap()
    bandwidthAccounting.EntityData.Children.Append("application", types.YChild{"Application", &bandwidthAccounting.Application})
    bandwidthAccounting.EntityData.Children.Append("account-flooding-threshold", types.YChild{"AccountFloodingThreshold", &bandwidthAccounting.AccountFloodingThreshold})
    bandwidthAccounting.EntityData.Leafs = types.NewOrderedMap()
    bandwidthAccounting.EntityData.Leafs.Append("sampling-interval", types.YLeaf{"SamplingInterval", bandwidthAccounting.SamplingInterval})
    bandwidthAccounting.EntityData.Leafs.Append("adjustment-factor", types.YLeaf{"AdjustmentFactor", bandwidthAccounting.AdjustmentFactor})
    bandwidthAccounting.EntityData.Leafs.Append("collection-type-rsvp-te", types.YLeaf{"CollectionTypeRsvpTe", bandwidthAccounting.CollectionTypeRsvpTe})
    bandwidthAccounting.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bandwidthAccounting.Enable})

    bandwidthAccounting.EntityData.YListKeys = []string {}

    return &(bandwidthAccounting.EntityData)
}

// MplsTe_GlobalAttributes_BandwidthAccounting_Application
// Bandwidth accounting application configuration
// data
type MplsTe_GlobalAttributes_BandwidthAccounting_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object enables the application. The type is bool. The default value is
    // false.
    ApplicationEnforced interface{}

    // This object sets the application interval in seconds for bandwidth
    // accounting. Default to 180 seconds. The type is interface{} with range:
    // 30..1800. Units are second. The default value is 180.
    ApplicationInterval interface{}
}

func (application *MplsTe_GlobalAttributes_BandwidthAccounting_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xr"
    application.EntityData.ParentYangName = "bandwidth-accounting"
    application.EntityData.SegmentPath = "application"
    application.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    application.EntityData.Children = types.NewOrderedMap()
    application.EntityData.Leafs = types.NewOrderedMap()
    application.EntityData.Leafs.Append("application-enforced", types.YLeaf{"ApplicationEnforced", application.ApplicationEnforced})
    application.EntityData.Leafs.Append("application-interval", types.YLeaf{"ApplicationInterval", application.ApplicationInterval})

    application.EntityData.YListKeys = []string {}

    return &(application.EntityData)
}

// MplsTe_GlobalAttributes_BandwidthAccounting_AccountFloodingThreshold
// This object sets the flooding threshold as
// percentage of total link bandwidth for
// bandwidth accounting. Default to 10%, 10%
type MplsTe_GlobalAttributes_BandwidthAccounting_AccountFloodingThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Upward flooding Threshold in percentages of total bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage. The default value is
    // 10.
    UpThreshold interface{}

    // Downward flooding Threshold in percentages of total bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage. The default value is
    // 10.
    DownThreshold interface{}
}

func (accountFloodingThreshold *MplsTe_GlobalAttributes_BandwidthAccounting_AccountFloodingThreshold) GetEntityData() *types.CommonEntityData {
    accountFloodingThreshold.EntityData.YFilter = accountFloodingThreshold.YFilter
    accountFloodingThreshold.EntityData.YangName = "account-flooding-threshold"
    accountFloodingThreshold.EntityData.BundleName = "cisco_ios_xr"
    accountFloodingThreshold.EntityData.ParentYangName = "bandwidth-accounting"
    accountFloodingThreshold.EntityData.SegmentPath = "account-flooding-threshold"
    accountFloodingThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountFloodingThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountFloodingThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountFloodingThreshold.EntityData.Children = types.NewOrderedMap()
    accountFloodingThreshold.EntityData.Leafs = types.NewOrderedMap()
    accountFloodingThreshold.EntityData.Leafs.Append("up-threshold", types.YLeaf{"UpThreshold", accountFloodingThreshold.UpThreshold})
    accountFloodingThreshold.EntityData.Leafs.Append("down-threshold", types.YLeaf{"DownThreshold", accountFloodingThreshold.DownThreshold})

    accountFloodingThreshold.EntityData.YListKeys = []string {}

    return &(accountFloodingThreshold.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes
// Configuration MPLS TE PCE attributes
type MplsTe_GlobalAttributes_PceAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request timeout value in seconds. The type is interface{} with range:
    // 5..100. Units are second. The default value is 10.
    RequestTimeout interface{}

    // PCE reoptimization period for PCE-based paths. The type is interface{} with
    // range: 60..604800. Units are second. The default value is 60.
    ReoptimizePeriod interface{}

    // Address of this PCE. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Deadtimer interval in seconds. The type is interface{} with range: 0..255.
    // Units are second. The default value is 120.
    Deadtimer interface{}

    // Keepalive interval in seconds. The type is interface{} with range: 0..255.
    // Units are second. The default value is 30.
    Keepalive interface{}

    // Keepalive interval tolerance in seconds. The type is interface{} with
    // range: 0..255. Units are second. The default value is 10.
    KeepaliveTolerance interface{}

    // PCE Peer Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerSourceAddr interface{}

    // PCE speaker entity identifier. The type is string with length: 1..256.
    SpeakerEntityId interface{}

    // PCE segment routing capability. The type is interface{}.
    SegmentRouting interface{}

    // MD5 password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}

    // Keychain based authentication. The type is string with length: 1..32.
    Keychain interface{}

    // Precedence order. The type is interface{} with range: 0..255.
    Precedence interface{}

    // PCE Stateful.
    PceStateful MplsTe_GlobalAttributes_PceAttributes_PceStateful

    // Configure PCE (Path Computation Element) timers.
    Timer MplsTe_GlobalAttributes_PceAttributes_Timer

    // Configure PCE peers.
    Peers MplsTe_GlobalAttributes_PceAttributes_Peers

    // Configure PCE (Path Computation Element) logging feature.
    Logging MplsTe_GlobalAttributes_PceAttributes_Logging
}

func (pceAttributes *MplsTe_GlobalAttributes_PceAttributes) GetEntityData() *types.CommonEntityData {
    pceAttributes.EntityData.YFilter = pceAttributes.YFilter
    pceAttributes.EntityData.YangName = "pce-attributes"
    pceAttributes.EntityData.BundleName = "cisco_ios_xr"
    pceAttributes.EntityData.ParentYangName = "global-attributes"
    pceAttributes.EntityData.SegmentPath = "pce-attributes"
    pceAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceAttributes.EntityData.Children = types.NewOrderedMap()
    pceAttributes.EntityData.Children.Append("pce-stateful", types.YChild{"PceStateful", &pceAttributes.PceStateful})
    pceAttributes.EntityData.Children.Append("timer", types.YChild{"Timer", &pceAttributes.Timer})
    pceAttributes.EntityData.Children.Append("peers", types.YChild{"Peers", &pceAttributes.Peers})
    pceAttributes.EntityData.Children.Append("logging", types.YChild{"Logging", &pceAttributes.Logging})
    pceAttributes.EntityData.Leafs = types.NewOrderedMap()
    pceAttributes.EntityData.Leafs.Append("request-timeout", types.YLeaf{"RequestTimeout", pceAttributes.RequestTimeout})
    pceAttributes.EntityData.Leafs.Append("reoptimize-period", types.YLeaf{"ReoptimizePeriod", pceAttributes.ReoptimizePeriod})
    pceAttributes.EntityData.Leafs.Append("address", types.YLeaf{"Address", pceAttributes.Address})
    pceAttributes.EntityData.Leafs.Append("deadtimer", types.YLeaf{"Deadtimer", pceAttributes.Deadtimer})
    pceAttributes.EntityData.Leafs.Append("keepalive", types.YLeaf{"Keepalive", pceAttributes.Keepalive})
    pceAttributes.EntityData.Leafs.Append("keepalive-tolerance", types.YLeaf{"KeepaliveTolerance", pceAttributes.KeepaliveTolerance})
    pceAttributes.EntityData.Leafs.Append("peer-source-addr", types.YLeaf{"PeerSourceAddr", pceAttributes.PeerSourceAddr})
    pceAttributes.EntityData.Leafs.Append("speaker-entity-id", types.YLeaf{"SpeakerEntityId", pceAttributes.SpeakerEntityId})
    pceAttributes.EntityData.Leafs.Append("segment-routing", types.YLeaf{"SegmentRouting", pceAttributes.SegmentRouting})
    pceAttributes.EntityData.Leafs.Append("password", types.YLeaf{"Password", pceAttributes.Password})
    pceAttributes.EntityData.Leafs.Append("keychain", types.YLeaf{"Keychain", pceAttributes.Keychain})
    pceAttributes.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", pceAttributes.Precedence})

    pceAttributes.EntityData.YListKeys = []string {}

    return &(pceAttributes.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_PceStateful
// PCE Stateful
type MplsTe_GlobalAttributes_PceAttributes_PceStateful struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable reoptimization by PCC after path failures. The type is interface{}.
    FastRepair interface{}

    // PCE stateful instantiation capability. The type is interface{}.
    Instantiation interface{}

    // Enable processing of PCEP Cisco extension. The type is interface{}.
    CiscoExtension interface{}

    // Delegate all statically configured tunnels. The type is interface{}.
    Delegation interface{}

    // Report all statically configured tunnels. The type is interface{}.
    Report interface{}

    // PCE stateful capability. The type is interface{}.
    Enable interface{}

    // Configure Stateful PCE (Path Computation Element) timers.
    StatefulTimers MplsTe_GlobalAttributes_PceAttributes_PceStateful_StatefulTimers
}

func (pceStateful *MplsTe_GlobalAttributes_PceAttributes_PceStateful) GetEntityData() *types.CommonEntityData {
    pceStateful.EntityData.YFilter = pceStateful.YFilter
    pceStateful.EntityData.YangName = "pce-stateful"
    pceStateful.EntityData.BundleName = "cisco_ios_xr"
    pceStateful.EntityData.ParentYangName = "pce-attributes"
    pceStateful.EntityData.SegmentPath = "pce-stateful"
    pceStateful.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceStateful.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceStateful.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceStateful.EntityData.Children = types.NewOrderedMap()
    pceStateful.EntityData.Children.Append("stateful-timers", types.YChild{"StatefulTimers", &pceStateful.StatefulTimers})
    pceStateful.EntityData.Leafs = types.NewOrderedMap()
    pceStateful.EntityData.Leafs.Append("fast-repair", types.YLeaf{"FastRepair", pceStateful.FastRepair})
    pceStateful.EntityData.Leafs.Append("instantiation", types.YLeaf{"Instantiation", pceStateful.Instantiation})
    pceStateful.EntityData.Leafs.Append("cisco-extension", types.YLeaf{"CiscoExtension", pceStateful.CiscoExtension})
    pceStateful.EntityData.Leafs.Append("delegation", types.YLeaf{"Delegation", pceStateful.Delegation})
    pceStateful.EntityData.Leafs.Append("report", types.YLeaf{"Report", pceStateful.Report})
    pceStateful.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", pceStateful.Enable})

    pceStateful.EntityData.YListKeys = []string {}

    return &(pceStateful.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_PceStateful_StatefulTimers
// Configure Stateful PCE (Path Computation
// Element) timers
type MplsTe_GlobalAttributes_PceAttributes_PceStateful_StatefulTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timer for static tunnel redelegation in seconds, default is 180 seconds.
    // The type is interface{} with range: 0..3600. Units are second. The default
    // value is 180.
    RedelegationTimeout interface{}

    // State timeout for LSPs without delegation in seconds, zero means immediate
    // removal, default is 180 seconds. The type is interface{} with range:
    // 0..3600. Units are second. The default value is 180.
    StateTimeout interface{}
}

func (statefulTimers *MplsTe_GlobalAttributes_PceAttributes_PceStateful_StatefulTimers) GetEntityData() *types.CommonEntityData {
    statefulTimers.EntityData.YFilter = statefulTimers.YFilter
    statefulTimers.EntityData.YangName = "stateful-timers"
    statefulTimers.EntityData.BundleName = "cisco_ios_xr"
    statefulTimers.EntityData.ParentYangName = "pce-stateful"
    statefulTimers.EntityData.SegmentPath = "stateful-timers"
    statefulTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statefulTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statefulTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statefulTimers.EntityData.Children = types.NewOrderedMap()
    statefulTimers.EntityData.Leafs = types.NewOrderedMap()
    statefulTimers.EntityData.Leafs.Append("redelegation-timeout", types.YLeaf{"RedelegationTimeout", statefulTimers.RedelegationTimeout})
    statefulTimers.EntityData.Leafs.Append("state-timeout", types.YLeaf{"StateTimeout", statefulTimers.StateTimeout})

    statefulTimers.EntityData.YListKeys = []string {}

    return &(statefulTimers.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_Timer
// Configure PCE (Path Computation Element)
// timers
type MplsTe_GlobalAttributes_PceAttributes_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
}

func (timer *MplsTe_GlobalAttributes_PceAttributes_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "pce-attributes"
    timer.EntityData.SegmentPath = "timer"
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()

    timer.EntityData.YListKeys = []string {}

    return &(timer.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_Peers
// Configure PCE peers
type MplsTe_GlobalAttributes_PceAttributes_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer. The type is slice of
    // MplsTe_GlobalAttributes_PceAttributes_Peers_Peer.
    Peer []*MplsTe_GlobalAttributes_PceAttributes_Peers_Peer
}

func (peers *MplsTe_GlobalAttributes_PceAttributes_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "pce-attributes"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Children.Append("peer", types.YChild{"Peer", nil})
    for i := range peers.Peer {
        peers.EntityData.Children.Append(types.GetSegmentPath(peers.Peer[i]), types.YChild{"Peer", peers.Peer[i]})
    }
    peers.EntityData.Leafs = types.NewOrderedMap()

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_Peers_Peer
// PCE peer
type MplsTe_GlobalAttributes_PceAttributes_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address of PCE Peer. The type is string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PcePeerAddress interface{}

    // Enabled PCE peer (default source address uses local). The type is
    // interface{}.
    Enable interface{}

    // MD5 password. The type is string with pattern: (!.+)|([^!].+).
    Password interface{}

    // Keychain based authentication. The type is string with length: 1..32.
    Keychain interface{}

    // Precedence order. The type is interface{} with range: 0..255.
    Precedence interface{}
}

func (peer *MplsTe_GlobalAttributes_PceAttributes_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PcePeerAddress, "pce-peer-address")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("pce-peer-address", types.YLeaf{"PcePeerAddress", peer.PcePeerAddress})
    peer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", peer.Enable})
    peer.EntityData.Leafs.Append("password", types.YLeaf{"Password", peer.Password})
    peer.EntityData.Leafs.Append("keychain", types.YLeaf{"Keychain", peer.Keychain})
    peer.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", peer.Precedence})

    peer.EntityData.YListKeys = []string {"PcePeerAddress"}

    return &(peer.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_Logging
// Configure PCE (Path Computation Element)
// logging feature
type MplsTe_GlobalAttributes_PceAttributes_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure logging events.
    Events MplsTe_GlobalAttributes_PceAttributes_Logging_Events
}

func (logging *MplsTe_GlobalAttributes_PceAttributes_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "pce-attributes"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Children.Append("events", types.YChild{"Events", &logging.Events})
    logging.EntityData.Leafs = types.NewOrderedMap()

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsTe_GlobalAttributes_PceAttributes_Logging_Events
// Configure logging events
type MplsTe_GlobalAttributes_PceAttributes_Logging_Events struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer status changes logging. The type is interface{}.
    PeerStatus interface{}
}

func (events *MplsTe_GlobalAttributes_PceAttributes_Logging_Events) GetEntityData() *types.CommonEntityData {
    events.EntityData.YFilter = events.YFilter
    events.EntityData.YangName = "events"
    events.EntityData.BundleName = "cisco_ios_xr"
    events.EntityData.ParentYangName = "logging"
    events.EntityData.SegmentPath = "events"
    events.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    events.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    events.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    events.EntityData.Children = types.NewOrderedMap()
    events.EntityData.Leafs = types.NewOrderedMap()
    events.EntityData.Leafs.Append("peer-status", types.YLeaf{"PeerStatus", events.PeerStatus})

    events.EntityData.YListKeys = []string {}

    return &(events.EntityData)
}

// MplsTe_GlobalAttributes_LspOutOfResource
// Configure LSP OOR attributes in MPLS-TE
type MplsTe_GlobalAttributes_LspOutOfResource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for LSP OOR Red/Major State.
    LspOorRedState MplsTe_GlobalAttributes_LspOutOfResource_LspOorRedState

    // Configuration for LSP OOR Yellow/Minor State.
    LspOorYellowState MplsTe_GlobalAttributes_LspOutOfResource_LspOorYellowState
}

func (lspOutOfResource *MplsTe_GlobalAttributes_LspOutOfResource) GetEntityData() *types.CommonEntityData {
    lspOutOfResource.EntityData.YFilter = lspOutOfResource.YFilter
    lspOutOfResource.EntityData.YangName = "lsp-out-of-resource"
    lspOutOfResource.EntityData.BundleName = "cisco_ios_xr"
    lspOutOfResource.EntityData.ParentYangName = "global-attributes"
    lspOutOfResource.EntityData.SegmentPath = "lsp-out-of-resource"
    lspOutOfResource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspOutOfResource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspOutOfResource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspOutOfResource.EntityData.Children = types.NewOrderedMap()
    lspOutOfResource.EntityData.Children.Append("lsp-oor-red-state", types.YChild{"LspOorRedState", &lspOutOfResource.LspOorRedState})
    lspOutOfResource.EntityData.Children.Append("lsp-oor-yellow-state", types.YChild{"LspOorYellowState", &lspOutOfResource.LspOorYellowState})
    lspOutOfResource.EntityData.Leafs = types.NewOrderedMap()

    lspOutOfResource.EntityData.YListKeys = []string {}

    return &(lspOutOfResource.EntityData)
}

// MplsTe_GlobalAttributes_LspOutOfResource_LspOorRedState
// Configuration for LSP OOR Red/Major State
type MplsTe_GlobalAttributes_LspOutOfResource_LspOorRedState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold for all transit LSPs. The type is interface{} with range:
    // 0..4294967295.
    AllTransitLspThreshold interface{}

    // Threshold for unprotected transit LSPs. The type is interface{} with range:
    // 0..4294967295.
    UnprotectedTransitLspThreshold interface{}
}

func (lspOorRedState *MplsTe_GlobalAttributes_LspOutOfResource_LspOorRedState) GetEntityData() *types.CommonEntityData {
    lspOorRedState.EntityData.YFilter = lspOorRedState.YFilter
    lspOorRedState.EntityData.YangName = "lsp-oor-red-state"
    lspOorRedState.EntityData.BundleName = "cisco_ios_xr"
    lspOorRedState.EntityData.ParentYangName = "lsp-out-of-resource"
    lspOorRedState.EntityData.SegmentPath = "lsp-oor-red-state"
    lspOorRedState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspOorRedState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspOorRedState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspOorRedState.EntityData.Children = types.NewOrderedMap()
    lspOorRedState.EntityData.Leafs = types.NewOrderedMap()
    lspOorRedState.EntityData.Leafs.Append("all-transit-lsp-threshold", types.YLeaf{"AllTransitLspThreshold", lspOorRedState.AllTransitLspThreshold})
    lspOorRedState.EntityData.Leafs.Append("unprotected-transit-lsp-threshold", types.YLeaf{"UnprotectedTransitLspThreshold", lspOorRedState.UnprotectedTransitLspThreshold})

    lspOorRedState.EntityData.YListKeys = []string {}

    return &(lspOorRedState.EntityData)
}

// MplsTe_GlobalAttributes_LspOutOfResource_LspOorYellowState
// Configuration for LSP OOR Yellow/Minor State
type MplsTe_GlobalAttributes_LspOutOfResource_LspOorYellowState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Threshold for all transit LSPs. The type is interface{} with range:
    // 0..4294967295.
    AllTransitLspThreshold interface{}

    // Threshold for unprotected transit LSPs. The type is interface{} with range:
    // 0..4294967295.
    UnprotectedTransitLspThreshold interface{}
}

func (lspOorYellowState *MplsTe_GlobalAttributes_LspOutOfResource_LspOorYellowState) GetEntityData() *types.CommonEntityData {
    lspOorYellowState.EntityData.YFilter = lspOorYellowState.YFilter
    lspOorYellowState.EntityData.YangName = "lsp-oor-yellow-state"
    lspOorYellowState.EntityData.BundleName = "cisco_ios_xr"
    lspOorYellowState.EntityData.ParentYangName = "lsp-out-of-resource"
    lspOorYellowState.EntityData.SegmentPath = "lsp-oor-yellow-state"
    lspOorYellowState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspOorYellowState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspOorYellowState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspOorYellowState.EntityData.Children = types.NewOrderedMap()
    lspOorYellowState.EntityData.Leafs = types.NewOrderedMap()
    lspOorYellowState.EntityData.Leafs.Append("all-transit-lsp-threshold", types.YLeaf{"AllTransitLspThreshold", lspOorYellowState.AllTransitLspThreshold})
    lspOorYellowState.EntityData.Leafs.Append("unprotected-transit-lsp-threshold", types.YLeaf{"UnprotectedTransitLspThreshold", lspOorYellowState.UnprotectedTransitLspThreshold})

    lspOorYellowState.EntityData.YListKeys = []string {}

    return &(lspOorYellowState.EntityData)
}

// MplsTe_GlobalAttributes_SoftPreemption
// Soft preemption configuration data
type MplsTe_GlobalAttributes_SoftPreemption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object sets the timeout in seconds before hard preemption is
    // triggered. The type is interface{} with range: 1..300. Units are second.
    // The default value is 60.
    Timeout interface{}

    // This object controls whether FRR rewrite during soft preemption is enabled.
    // The type is interface{}.
    FrrRewrite interface{}

    // This object controls whether soft preemption is enabled. This object must
    // be set before setting any other objects under the SoftPreemption class. The
    // type is bool.
    Enable interface{}
}

func (softPreemption *MplsTe_GlobalAttributes_SoftPreemption) GetEntityData() *types.CommonEntityData {
    softPreemption.EntityData.YFilter = softPreemption.YFilter
    softPreemption.EntityData.YangName = "soft-preemption"
    softPreemption.EntityData.BundleName = "cisco_ios_xr"
    softPreemption.EntityData.ParentYangName = "global-attributes"
    softPreemption.EntityData.SegmentPath = "soft-preemption"
    softPreemption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    softPreemption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    softPreemption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    softPreemption.EntityData.Children = types.NewOrderedMap()
    softPreemption.EntityData.Leafs = types.NewOrderedMap()
    softPreemption.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", softPreemption.Timeout})
    softPreemption.EntityData.Leafs.Append("frr-rewrite", types.YLeaf{"FrrRewrite", softPreemption.FrrRewrite})
    softPreemption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", softPreemption.Enable})

    softPreemption.EntityData.YListKeys = []string {}

    return &(softPreemption.EntityData)
}

// MplsTe_GlobalAttributes_FastReroute
// Configure fast reroute attributes
type MplsTe_GlobalAttributes_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure fast reroute timers.
    Timers MplsTe_GlobalAttributes_FastReroute_Timers
}

func (fastReroute *MplsTe_GlobalAttributes_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "global-attributes"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("timers", types.YChild{"Timers", &fastReroute.Timers})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// MplsTe_GlobalAttributes_FastReroute_Timers
// Configure fast reroute timers
type MplsTe_GlobalAttributes_FastReroute_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Seconds before backup declared UP (0 disables hold-timer). The type is
    // interface{} with range: 0..604800. Units are second.
    HoldBackup interface{}

    // The value of the promotion timer in seconds. The type is interface{} with
    // range: 0..604800. Units are second.
    Promotion interface{}
}

func (timers *MplsTe_GlobalAttributes_FastReroute_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "fast-reroute"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Leafs = types.NewOrderedMap()
    timers.EntityData.Leafs.Append("hold-backup", types.YLeaf{"HoldBackup", timers.HoldBackup})
    timers.EntityData.Leafs.Append("promotion", types.YLeaf{"Promotion", timers.Promotion})

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection
// Path selection configuration
type MplsTe_GlobalAttributes_PathSelection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection cost limit configuration for all tunnels. The type is
    // interface{} with range: 1..4294967295.
    CostLimit interface{}

    // CSPF tiebreaker to use in path calculation. The type is
    // MplsTePathSelectionTiebreaker.
    Tiebreaker interface{}

    // Metric to use in path calculation. The type is MplsTePathSelectionMetric.
    Metric interface{}

    // Use only the IGP instance of the incoming interface. The type is bool.
    LooseDomainMatch interface{}

    // Path selection Loose ERO Metric Class configuration.
    LooseMetrics MplsTe_GlobalAttributes_PathSelection_LooseMetrics

    // Path invalidation configuration for all tunnels.
    Invalidation MplsTe_GlobalAttributes_PathSelection_Invalidation

    // Path selection to ignore overload node during CSPF.
    IgnoreOverloadRole MplsTe_GlobalAttributes_PathSelection_IgnoreOverloadRole

    // Path selection Loose ERO Affinity Class configuration.
    LooseAffinities MplsTe_GlobalAttributes_PathSelection_LooseAffinities
}

func (pathSelection *MplsTe_GlobalAttributes_PathSelection) GetEntityData() *types.CommonEntityData {
    pathSelection.EntityData.YFilter = pathSelection.YFilter
    pathSelection.EntityData.YangName = "path-selection"
    pathSelection.EntityData.BundleName = "cisco_ios_xr"
    pathSelection.EntityData.ParentYangName = "global-attributes"
    pathSelection.EntityData.SegmentPath = "path-selection"
    pathSelection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathSelection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathSelection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathSelection.EntityData.Children = types.NewOrderedMap()
    pathSelection.EntityData.Children.Append("loose-metrics", types.YChild{"LooseMetrics", &pathSelection.LooseMetrics})
    pathSelection.EntityData.Children.Append("invalidation", types.YChild{"Invalidation", &pathSelection.Invalidation})
    pathSelection.EntityData.Children.Append("ignore-overload-role", types.YChild{"IgnoreOverloadRole", &pathSelection.IgnoreOverloadRole})
    pathSelection.EntityData.Children.Append("loose-affinities", types.YChild{"LooseAffinities", &pathSelection.LooseAffinities})
    pathSelection.EntityData.Leafs = types.NewOrderedMap()
    pathSelection.EntityData.Leafs.Append("cost-limit", types.YLeaf{"CostLimit", pathSelection.CostLimit})
    pathSelection.EntityData.Leafs.Append("tiebreaker", types.YLeaf{"Tiebreaker", pathSelection.Tiebreaker})
    pathSelection.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", pathSelection.Metric})
    pathSelection.EntityData.Leafs.Append("loose-domain-match", types.YLeaf{"LooseDomainMatch", pathSelection.LooseDomainMatch})

    pathSelection.EntityData.YListKeys = []string {}

    return &(pathSelection.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection_LooseMetrics
// Path selection Loose ERO Metric Class
// configuration
type MplsTe_GlobalAttributes_PathSelection_LooseMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection Loose ERO Metric configuration. The type is slice of
    // MplsTe_GlobalAttributes_PathSelection_LooseMetrics_LooseMetric.
    LooseMetric []*MplsTe_GlobalAttributes_PathSelection_LooseMetrics_LooseMetric
}

func (looseMetrics *MplsTe_GlobalAttributes_PathSelection_LooseMetrics) GetEntityData() *types.CommonEntityData {
    looseMetrics.EntityData.YFilter = looseMetrics.YFilter
    looseMetrics.EntityData.YangName = "loose-metrics"
    looseMetrics.EntityData.BundleName = "cisco_ios_xr"
    looseMetrics.EntityData.ParentYangName = "path-selection"
    looseMetrics.EntityData.SegmentPath = "loose-metrics"
    looseMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    looseMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    looseMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    looseMetrics.EntityData.Children = types.NewOrderedMap()
    looseMetrics.EntityData.Children.Append("loose-metric", types.YChild{"LooseMetric", nil})
    for i := range looseMetrics.LooseMetric {
        looseMetrics.EntityData.Children.Append(types.GetSegmentPath(looseMetrics.LooseMetric[i]), types.YChild{"LooseMetric", looseMetrics.LooseMetric[i]})
    }
    looseMetrics.EntityData.Leafs = types.NewOrderedMap()

    looseMetrics.EntityData.YListKeys = []string {}

    return &(looseMetrics.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection_LooseMetrics_LooseMetric
// Path selection Loose ERO Metric configuration
type MplsTe_GlobalAttributes_PathSelection_LooseMetrics_LooseMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path Selection class Type. The type is interface{}
    // with range: 0..7.
    ClassType interface{}

    // Metric to use for ERO Expansion. The type is MplsTePathSelectionMetric.
    // This attribute is mandatory.
    MetricType interface{}
}

func (looseMetric *MplsTe_GlobalAttributes_PathSelection_LooseMetrics_LooseMetric) GetEntityData() *types.CommonEntityData {
    looseMetric.EntityData.YFilter = looseMetric.YFilter
    looseMetric.EntityData.YangName = "loose-metric"
    looseMetric.EntityData.BundleName = "cisco_ios_xr"
    looseMetric.EntityData.ParentYangName = "loose-metrics"
    looseMetric.EntityData.SegmentPath = "loose-metric" + types.AddKeyToken(looseMetric.ClassType, "class-type")
    looseMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    looseMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    looseMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    looseMetric.EntityData.Children = types.NewOrderedMap()
    looseMetric.EntityData.Leafs = types.NewOrderedMap()
    looseMetric.EntityData.Leafs.Append("class-type", types.YLeaf{"ClassType", looseMetric.ClassType})
    looseMetric.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", looseMetric.MetricType})

    looseMetric.EntityData.YListKeys = []string {"ClassType"}

    return &(looseMetric.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection_Invalidation
// Path invalidation configuration for all
// tunnels
type MplsTe_GlobalAttributes_PathSelection_Invalidation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path Invalidation Timeout. The type is interface{} with range: 0..60000.
    PathInvalidationTimeout interface{}

    // Path Invalidation Action. The type is PathInvalidationAction.
    PathInvalidationAction interface{}
}

func (invalidation *MplsTe_GlobalAttributes_PathSelection_Invalidation) GetEntityData() *types.CommonEntityData {
    invalidation.EntityData.YFilter = invalidation.YFilter
    invalidation.EntityData.YangName = "invalidation"
    invalidation.EntityData.BundleName = "cisco_ios_xr"
    invalidation.EntityData.ParentYangName = "path-selection"
    invalidation.EntityData.SegmentPath = "invalidation"
    invalidation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidation.EntityData.Children = types.NewOrderedMap()
    invalidation.EntityData.Leafs = types.NewOrderedMap()
    invalidation.EntityData.Leafs.Append("path-invalidation-timeout", types.YLeaf{"PathInvalidationTimeout", invalidation.PathInvalidationTimeout})
    invalidation.EntityData.Leafs.Append("path-invalidation-action", types.YLeaf{"PathInvalidationAction", invalidation.PathInvalidationAction})

    invalidation.EntityData.YListKeys = []string {}

    return &(invalidation.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection_IgnoreOverloadRole
// Path selection to ignore overload node during
// CSPF
type MplsTe_GlobalAttributes_PathSelection_IgnoreOverloadRole struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set if the OL-bit is to be applied to tunnel heads. The type is bool.
    Head interface{}

    // Set if the OL-bit is to be applied to tunnel midpoints. The type is bool.
    Mid interface{}

    // Set if the OL-bit is to be applied to tunnel tails. The type is bool.
    Tail interface{}
}

func (ignoreOverloadRole *MplsTe_GlobalAttributes_PathSelection_IgnoreOverloadRole) GetEntityData() *types.CommonEntityData {
    ignoreOverloadRole.EntityData.YFilter = ignoreOverloadRole.YFilter
    ignoreOverloadRole.EntityData.YangName = "ignore-overload-role"
    ignoreOverloadRole.EntityData.BundleName = "cisco_ios_xr"
    ignoreOverloadRole.EntityData.ParentYangName = "path-selection"
    ignoreOverloadRole.EntityData.SegmentPath = "ignore-overload-role"
    ignoreOverloadRole.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ignoreOverloadRole.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ignoreOverloadRole.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ignoreOverloadRole.EntityData.Children = types.NewOrderedMap()
    ignoreOverloadRole.EntityData.Leafs = types.NewOrderedMap()
    ignoreOverloadRole.EntityData.Leafs.Append("head", types.YLeaf{"Head", ignoreOverloadRole.Head})
    ignoreOverloadRole.EntityData.Leafs.Append("mid", types.YLeaf{"Mid", ignoreOverloadRole.Mid})
    ignoreOverloadRole.EntityData.Leafs.Append("tail", types.YLeaf{"Tail", ignoreOverloadRole.Tail})

    ignoreOverloadRole.EntityData.YListKeys = []string {}

    return &(ignoreOverloadRole.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection_LooseAffinities
// Path selection Loose ERO Affinity Class
// configuration
type MplsTe_GlobalAttributes_PathSelection_LooseAffinities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection Loose ERO Affinity configuration. The type is slice of
    // MplsTe_GlobalAttributes_PathSelection_LooseAffinities_LooseAffinity.
    LooseAffinity []*MplsTe_GlobalAttributes_PathSelection_LooseAffinities_LooseAffinity
}

func (looseAffinities *MplsTe_GlobalAttributes_PathSelection_LooseAffinities) GetEntityData() *types.CommonEntityData {
    looseAffinities.EntityData.YFilter = looseAffinities.YFilter
    looseAffinities.EntityData.YangName = "loose-affinities"
    looseAffinities.EntityData.BundleName = "cisco_ios_xr"
    looseAffinities.EntityData.ParentYangName = "path-selection"
    looseAffinities.EntityData.SegmentPath = "loose-affinities"
    looseAffinities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    looseAffinities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    looseAffinities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    looseAffinities.EntityData.Children = types.NewOrderedMap()
    looseAffinities.EntityData.Children.Append("loose-affinity", types.YChild{"LooseAffinity", nil})
    for i := range looseAffinities.LooseAffinity {
        looseAffinities.EntityData.Children.Append(types.GetSegmentPath(looseAffinities.LooseAffinity[i]), types.YChild{"LooseAffinity", looseAffinities.LooseAffinity[i]})
    }
    looseAffinities.EntityData.Leafs = types.NewOrderedMap()

    looseAffinities.EntityData.YListKeys = []string {}

    return &(looseAffinities.EntityData)
}

// MplsTe_GlobalAttributes_PathSelection_LooseAffinities_LooseAffinity
// Path selection Loose ERO Affinity
// configuration
type MplsTe_GlobalAttributes_PathSelection_LooseAffinities_LooseAffinity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Path Selection class Type. The type is interface{}
    // with range: 0..7.
    ClassType interface{}

    // Affinity flags. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Affinity interface{}

    // Affinity mask. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Mask interface{}
}

func (looseAffinity *MplsTe_GlobalAttributes_PathSelection_LooseAffinities_LooseAffinity) GetEntityData() *types.CommonEntityData {
    looseAffinity.EntityData.YFilter = looseAffinity.YFilter
    looseAffinity.EntityData.YangName = "loose-affinity"
    looseAffinity.EntityData.BundleName = "cisco_ios_xr"
    looseAffinity.EntityData.ParentYangName = "loose-affinities"
    looseAffinity.EntityData.SegmentPath = "loose-affinity" + types.AddKeyToken(looseAffinity.ClassType, "class-type")
    looseAffinity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    looseAffinity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    looseAffinity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    looseAffinity.EntityData.Children = types.NewOrderedMap()
    looseAffinity.EntityData.Leafs = types.NewOrderedMap()
    looseAffinity.EntityData.Leafs.Append("class-type", types.YLeaf{"ClassType", looseAffinity.ClassType})
    looseAffinity.EntityData.Leafs.Append("affinity", types.YLeaf{"Affinity", looseAffinity.Affinity})
    looseAffinity.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", looseAffinity.Mask})

    looseAffinity.EntityData.YListKeys = []string {"ClassType"}

    return &(looseAffinity.EntityData)
}

// MplsTe_GlobalAttributes_AffinityMappings
// Affinity Mapping Table configuration
type MplsTe_GlobalAttributes_AffinityMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity Mapping configuration. The type is slice of
    // MplsTe_GlobalAttributes_AffinityMappings_AffinityMapping.
    AffinityMapping []*MplsTe_GlobalAttributes_AffinityMappings_AffinityMapping
}

func (affinityMappings *MplsTe_GlobalAttributes_AffinityMappings) GetEntityData() *types.CommonEntityData {
    affinityMappings.EntityData.YFilter = affinityMappings.YFilter
    affinityMappings.EntityData.YangName = "affinity-mappings"
    affinityMappings.EntityData.BundleName = "cisco_ios_xr"
    affinityMappings.EntityData.ParentYangName = "global-attributes"
    affinityMappings.EntityData.SegmentPath = "affinity-mappings"
    affinityMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMappings.EntityData.Children = types.NewOrderedMap()
    affinityMappings.EntityData.Children.Append("affinity-mapping", types.YChild{"AffinityMapping", nil})
    for i := range affinityMappings.AffinityMapping {
        affinityMappings.EntityData.Children.Append(types.GetSegmentPath(affinityMappings.AffinityMapping[i]), types.YChild{"AffinityMapping", affinityMappings.AffinityMapping[i]})
    }
    affinityMappings.EntityData.Leafs = types.NewOrderedMap()

    affinityMappings.EntityData.YListKeys = []string {}

    return &(affinityMappings.EntityData)
}

// MplsTe_GlobalAttributes_AffinityMappings_AffinityMapping
// Affinity Mapping configuration
type MplsTe_GlobalAttributes_AffinityMappings_AffinityMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Affinity Name. The type is string with length:
    // 1..32.
    AffinityName interface{}

    // Affinity value type. The type is MplsTeAffinityValue.
    ValueType interface{}

    // Affinity Value in Hex number or by Bit position. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Value interface{}
}

func (affinityMapping *MplsTe_GlobalAttributes_AffinityMappings_AffinityMapping) GetEntityData() *types.CommonEntityData {
    affinityMapping.EntityData.YFilter = affinityMapping.YFilter
    affinityMapping.EntityData.YangName = "affinity-mapping"
    affinityMapping.EntityData.BundleName = "cisco_ios_xr"
    affinityMapping.EntityData.ParentYangName = "affinity-mappings"
    affinityMapping.EntityData.SegmentPath = "affinity-mapping" + types.AddKeyToken(affinityMapping.AffinityName, "affinity-name")
    affinityMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityMapping.EntityData.Children = types.NewOrderedMap()
    affinityMapping.EntityData.Leafs = types.NewOrderedMap()
    affinityMapping.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", affinityMapping.AffinityName})
    affinityMapping.EntityData.Leafs.Append("value-type", types.YLeaf{"ValueType", affinityMapping.ValueType})
    affinityMapping.EntityData.Leafs.Append("value", types.YLeaf{"Value", affinityMapping.Value})

    affinityMapping.EntityData.YListKeys = []string {"AffinityName"}

    return &(affinityMapping.EntityData)
}

// MplsTe_TransportProfile
// MPLS transport profile configuration data
type MplsTe_TransportProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport profile global identifier. The type is interface{} with range:
    // 1..65535.
    GlobalId interface{}

    // Node identifier in IPv4 address format. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NodeId interface{}

    // Fault management.
    Fault MplsTe_TransportProfile_Fault

    // Alarm management.
    Alarm MplsTe_TransportProfile_Alarm

    // Configure BFD parameters.
    Bfd MplsTe_TransportProfile_Bfd

    // MPLS-TP tunnel mid-point table.
    Midpoints MplsTe_TransportProfile_Midpoints
}

func (transportProfile *MplsTe_TransportProfile) GetEntityData() *types.CommonEntityData {
    transportProfile.EntityData.YFilter = transportProfile.YFilter
    transportProfile.EntityData.YangName = "transport-profile"
    transportProfile.EntityData.BundleName = "cisco_ios_xr"
    transportProfile.EntityData.ParentYangName = "mpls-te"
    transportProfile.EntityData.SegmentPath = "transport-profile"
    transportProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportProfile.EntityData.Children = types.NewOrderedMap()
    transportProfile.EntityData.Children.Append("fault", types.YChild{"Fault", &transportProfile.Fault})
    transportProfile.EntityData.Children.Append("alarm", types.YChild{"Alarm", &transportProfile.Alarm})
    transportProfile.EntityData.Children.Append("bfd", types.YChild{"Bfd", &transportProfile.Bfd})
    transportProfile.EntityData.Children.Append("midpoints", types.YChild{"Midpoints", &transportProfile.Midpoints})
    transportProfile.EntityData.Leafs = types.NewOrderedMap()
    transportProfile.EntityData.Leafs.Append("global-id", types.YLeaf{"GlobalId", transportProfile.GlobalId})
    transportProfile.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", transportProfile.NodeId})

    transportProfile.EntityData.YListKeys = []string {}

    return &(transportProfile.EntityData)
}

// MplsTe_TransportProfile_Fault
// Fault management
type MplsTe_TransportProfile_Fault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Waiting time before restoring working LSP. The type is interface{} with
    // range: 0..2147483647. Units are second. The default value is 0.
    WaitToRestoreInterval interface{}

    // Periodic refresh interval for fault OAM messages. The type is interface{}
    // with range: 1..20. Units are second. The default value is 20.
    RefreshInterval interface{}

    // OAM events that trigger protection switching.
    ProtectionTrigger MplsTe_TransportProfile_Fault_ProtectionTrigger
}

func (fault *MplsTe_TransportProfile_Fault) GetEntityData() *types.CommonEntityData {
    fault.EntityData.YFilter = fault.YFilter
    fault.EntityData.YangName = "fault"
    fault.EntityData.BundleName = "cisco_ios_xr"
    fault.EntityData.ParentYangName = "transport-profile"
    fault.EntityData.SegmentPath = "fault"
    fault.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fault.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fault.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fault.EntityData.Children = types.NewOrderedMap()
    fault.EntityData.Children.Append("protection-trigger", types.YChild{"ProtectionTrigger", &fault.ProtectionTrigger})
    fault.EntityData.Leafs = types.NewOrderedMap()
    fault.EntityData.Leafs.Append("wait-to-restore-interval", types.YLeaf{"WaitToRestoreInterval", fault.WaitToRestoreInterval})
    fault.EntityData.Leafs.Append("refresh-interval", types.YLeaf{"RefreshInterval", fault.RefreshInterval})

    fault.EntityData.YListKeys = []string {}

    return &(fault.EntityData)
}

// MplsTe_TransportProfile_Fault_ProtectionTrigger
// OAM events that trigger protection switching
type MplsTe_TransportProfile_Fault_ProtectionTrigger struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable protection switching due to AIS event. The type is interface{}.
    Ais interface{}

    // Protection switching due to LDI event.
    Ldi MplsTe_TransportProfile_Fault_ProtectionTrigger_Ldi

    // Protection switching due to LKR event.
    Lkr MplsTe_TransportProfile_Fault_ProtectionTrigger_Lkr
}

func (protectionTrigger *MplsTe_TransportProfile_Fault_ProtectionTrigger) GetEntityData() *types.CommonEntityData {
    protectionTrigger.EntityData.YFilter = protectionTrigger.YFilter
    protectionTrigger.EntityData.YangName = "protection-trigger"
    protectionTrigger.EntityData.BundleName = "cisco_ios_xr"
    protectionTrigger.EntityData.ParentYangName = "fault"
    protectionTrigger.EntityData.SegmentPath = "protection-trigger"
    protectionTrigger.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectionTrigger.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectionTrigger.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectionTrigger.EntityData.Children = types.NewOrderedMap()
    protectionTrigger.EntityData.Children.Append("ldi", types.YChild{"Ldi", &protectionTrigger.Ldi})
    protectionTrigger.EntityData.Children.Append("lkr", types.YChild{"Lkr", &protectionTrigger.Lkr})
    protectionTrigger.EntityData.Leafs = types.NewOrderedMap()
    protectionTrigger.EntityData.Leafs.Append("ais", types.YLeaf{"Ais", protectionTrigger.Ais})

    protectionTrigger.EntityData.YListKeys = []string {}

    return &(protectionTrigger.EntityData)
}

// MplsTe_TransportProfile_Fault_ProtectionTrigger_Ldi
// Protection switching due to LDI event
type MplsTe_TransportProfile_Fault_ProtectionTrigger_Ldi struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable protection switching due to LDI event. The type is interface{}.
    Disable interface{}
}

func (ldi *MplsTe_TransportProfile_Fault_ProtectionTrigger_Ldi) GetEntityData() *types.CommonEntityData {
    ldi.EntityData.YFilter = ldi.YFilter
    ldi.EntityData.YangName = "ldi"
    ldi.EntityData.BundleName = "cisco_ios_xr"
    ldi.EntityData.ParentYangName = "protection-trigger"
    ldi.EntityData.SegmentPath = "ldi"
    ldi.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ldi.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ldi.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ldi.EntityData.Children = types.NewOrderedMap()
    ldi.EntityData.Leafs = types.NewOrderedMap()
    ldi.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", ldi.Disable})

    ldi.EntityData.YListKeys = []string {}

    return &(ldi.EntityData)
}

// MplsTe_TransportProfile_Fault_ProtectionTrigger_Lkr
// Protection switching due to LKR event
type MplsTe_TransportProfile_Fault_ProtectionTrigger_Lkr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable protection switching due to LKR event. The type is interface{}.
    Disable interface{}
}

func (lkr *MplsTe_TransportProfile_Fault_ProtectionTrigger_Lkr) GetEntityData() *types.CommonEntityData {
    lkr.EntityData.YFilter = lkr.YFilter
    lkr.EntityData.YangName = "lkr"
    lkr.EntityData.BundleName = "cisco_ios_xr"
    lkr.EntityData.ParentYangName = "protection-trigger"
    lkr.EntityData.SegmentPath = "lkr"
    lkr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lkr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lkr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lkr.EntityData.Children = types.NewOrderedMap()
    lkr.EntityData.Leafs = types.NewOrderedMap()
    lkr.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", lkr.Disable})

    lkr.EntityData.YListKeys = []string {}

    return &(lkr.EntityData)
}

// MplsTe_TransportProfile_Alarm
// Alarm management
type MplsTe_TransportProfile_Alarm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Duration of soaking alarms. The type is interface{} with range: 0..10.
    // Units are second. The default value is 3.
    SoakTime interface{}

    // Enable Transport Profile Alarm. The type is interface{}.
    EnableAlarm interface{}

    // Suppress all tunnel/LSP alarms.
    SuppressEvent MplsTe_TransportProfile_Alarm_SuppressEvent
}

func (alarm *MplsTe_TransportProfile_Alarm) GetEntityData() *types.CommonEntityData {
    alarm.EntityData.YFilter = alarm.YFilter
    alarm.EntityData.YangName = "alarm"
    alarm.EntityData.BundleName = "cisco_ios_xr"
    alarm.EntityData.ParentYangName = "transport-profile"
    alarm.EntityData.SegmentPath = "alarm"
    alarm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    alarm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    alarm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    alarm.EntityData.Children = types.NewOrderedMap()
    alarm.EntityData.Children.Append("suppress-event", types.YChild{"SuppressEvent", &alarm.SuppressEvent})
    alarm.EntityData.Leafs = types.NewOrderedMap()
    alarm.EntityData.Leafs.Append("soak-time", types.YLeaf{"SoakTime", alarm.SoakTime})
    alarm.EntityData.Leafs.Append("enable-alarm", types.YLeaf{"EnableAlarm", alarm.EnableAlarm})

    alarm.EntityData.YListKeys = []string {}

    return &(alarm.EntityData)
}

// MplsTe_TransportProfile_Alarm_SuppressEvent
// Suppress all tunnel/LSP alarms
type MplsTe_TransportProfile_Alarm_SuppressEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable alarm suppression. The type is interface{}.
    Disable interface{}
}

func (suppressEvent *MplsTe_TransportProfile_Alarm_SuppressEvent) GetEntityData() *types.CommonEntityData {
    suppressEvent.EntityData.YFilter = suppressEvent.YFilter
    suppressEvent.EntityData.YangName = "suppress-event"
    suppressEvent.EntityData.BundleName = "cisco_ios_xr"
    suppressEvent.EntityData.ParentYangName = "alarm"
    suppressEvent.EntityData.SegmentPath = "suppress-event"
    suppressEvent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    suppressEvent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    suppressEvent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    suppressEvent.EntityData.Children = types.NewOrderedMap()
    suppressEvent.EntityData.Leafs = types.NewOrderedMap()
    suppressEvent.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", suppressEvent.Disable})

    suppressEvent.EntityData.YListKeys = []string {}

    return &(suppressEvent.EntityData)
}

// MplsTe_TransportProfile_Bfd
// Configure BFD parameters
type MplsTe_TransportProfile_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detect multiplier for standby transport profile LSP. The type is
    // interface{} with range: 2..10.
    DetectionMultiplierStandby interface{}

    // Detect multiplier. The type is interface{} with range: 2..10.
    DetectionMultiplier interface{}

    // Hello interval for standby transport profile LSPs, either in milli-seconds
    // or in micro-seconds.
    MinIntervalStandby MplsTe_TransportProfile_Bfd_MinIntervalStandby

    // Hello interval, either in milli-seconds or in micro-seconds.
    MinInterval MplsTe_TransportProfile_Bfd_MinInterval
}

func (bfd *MplsTe_TransportProfile_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "transport-profile"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Children.Append("min-interval-standby", types.YChild{"MinIntervalStandby", &bfd.MinIntervalStandby})
    bfd.EntityData.Children.Append("min-interval", types.YChild{"MinInterval", &bfd.MinInterval})
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier-standby", types.YLeaf{"DetectionMultiplierStandby", bfd.DetectionMultiplierStandby})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// MplsTe_TransportProfile_Bfd_MinIntervalStandby
// Hello interval for standby transport profile
// LSPs, either in milli-seconds or in
// micro-seconds
type MplsTe_TransportProfile_Bfd_MinIntervalStandby struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..5000. Units are millisecond.
    IntervalStandbyMs interface{}

    // Hello interval in micro-seconds. The type is interface{} with range:
    // 3000..5000000. Units are microsecond.
    IntervalStandbyUs interface{}
}

func (minIntervalStandby *MplsTe_TransportProfile_Bfd_MinIntervalStandby) GetEntityData() *types.CommonEntityData {
    minIntervalStandby.EntityData.YFilter = minIntervalStandby.YFilter
    minIntervalStandby.EntityData.YangName = "min-interval-standby"
    minIntervalStandby.EntityData.BundleName = "cisco_ios_xr"
    minIntervalStandby.EntityData.ParentYangName = "bfd"
    minIntervalStandby.EntityData.SegmentPath = "min-interval-standby"
    minIntervalStandby.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minIntervalStandby.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minIntervalStandby.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minIntervalStandby.EntityData.Children = types.NewOrderedMap()
    minIntervalStandby.EntityData.Leafs = types.NewOrderedMap()
    minIntervalStandby.EntityData.Leafs.Append("interval-standby-ms", types.YLeaf{"IntervalStandbyMs", minIntervalStandby.IntervalStandbyMs})
    minIntervalStandby.EntityData.Leafs.Append("interval-standby-us", types.YLeaf{"IntervalStandbyUs", minIntervalStandby.IntervalStandbyUs})

    minIntervalStandby.EntityData.YListKeys = []string {}

    return &(minIntervalStandby.EntityData)
}

// MplsTe_TransportProfile_Bfd_MinInterval
// Hello interval, either in milli-seconds or in
// micro-seconds
type MplsTe_TransportProfile_Bfd_MinInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..5000. Units are millisecond.
    IntervalMs interface{}

    // Hello interval in micro-seconds. The type is interface{} with range:
    // 3000..5000000. Units are microsecond.
    IntervalUs interface{}
}

func (minInterval *MplsTe_TransportProfile_Bfd_MinInterval) GetEntityData() *types.CommonEntityData {
    minInterval.EntityData.YFilter = minInterval.YFilter
    minInterval.EntityData.YangName = "min-interval"
    minInterval.EntityData.BundleName = "cisco_ios_xr"
    minInterval.EntityData.ParentYangName = "bfd"
    minInterval.EntityData.SegmentPath = "min-interval"
    minInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    minInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    minInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    minInterval.EntityData.Children = types.NewOrderedMap()
    minInterval.EntityData.Leafs = types.NewOrderedMap()
    minInterval.EntityData.Leafs.Append("interval-ms", types.YLeaf{"IntervalMs", minInterval.IntervalMs})
    minInterval.EntityData.Leafs.Append("interval-us", types.YLeaf{"IntervalUs", minInterval.IntervalUs})

    minInterval.EntityData.YListKeys = []string {}

    return &(minInterval.EntityData)
}

// MplsTe_TransportProfile_Midpoints
// MPLS-TP tunnel mid-point table
type MplsTe_TransportProfile_Midpoints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport profile mid-point identifier. The type is slice of
    // MplsTe_TransportProfile_Midpoints_Midpoint.
    Midpoint []*MplsTe_TransportProfile_Midpoints_Midpoint
}

func (midpoints *MplsTe_TransportProfile_Midpoints) GetEntityData() *types.CommonEntityData {
    midpoints.EntityData.YFilter = midpoints.YFilter
    midpoints.EntityData.YangName = "midpoints"
    midpoints.EntityData.BundleName = "cisco_ios_xr"
    midpoints.EntityData.ParentYangName = "transport-profile"
    midpoints.EntityData.SegmentPath = "midpoints"
    midpoints.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    midpoints.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    midpoints.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    midpoints.EntityData.Children = types.NewOrderedMap()
    midpoints.EntityData.Children.Append("midpoint", types.YChild{"Midpoint", nil})
    for i := range midpoints.Midpoint {
        midpoints.EntityData.Children.Append(types.GetSegmentPath(midpoints.Midpoint[i]), types.YChild{"Midpoint", midpoints.Midpoint[i]})
    }
    midpoints.EntityData.Leafs = types.NewOrderedMap()

    midpoints.EntityData.YListKeys = []string {}

    return &(midpoints.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint
// Transport profile mid-point identifier
type MplsTe_TransportProfile_Midpoints_Midpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of mid-point. The type is string with length:
    // 1..64.
    MidpointName interface{}

    // Tunnel Name. The type is string.
    TunnelName interface{}

    // Enable LSP protection. The type is interface{}.
    LspProtect interface{}

    // Numeric identifier. The type is interface{} with range: 0..65535.
    LspId interface{}

    // Node identifier, tunnel identifier and optional global identifier of the
    // source of the LSP.
    Source MplsTe_TransportProfile_Midpoints_Midpoint_Source

    // Node identifier, tunnel identifier and optional global identifier of the
    // destination of the LSP.
    Destination MplsTe_TransportProfile_Midpoints_Midpoint_Destination

    // Forward transport profile LSP.
    ForwardLsp MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp

    // none.
    ReverseLsp MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp
}

func (midpoint *MplsTe_TransportProfile_Midpoints_Midpoint) GetEntityData() *types.CommonEntityData {
    midpoint.EntityData.YFilter = midpoint.YFilter
    midpoint.EntityData.YangName = "midpoint"
    midpoint.EntityData.BundleName = "cisco_ios_xr"
    midpoint.EntityData.ParentYangName = "midpoints"
    midpoint.EntityData.SegmentPath = "midpoint" + types.AddKeyToken(midpoint.MidpointName, "midpoint-name")
    midpoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    midpoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    midpoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    midpoint.EntityData.Children = types.NewOrderedMap()
    midpoint.EntityData.Children.Append("source", types.YChild{"Source", &midpoint.Source})
    midpoint.EntityData.Children.Append("destination", types.YChild{"Destination", &midpoint.Destination})
    midpoint.EntityData.Children.Append("forward-lsp", types.YChild{"ForwardLsp", &midpoint.ForwardLsp})
    midpoint.EntityData.Children.Append("reverse-lsp", types.YChild{"ReverseLsp", &midpoint.ReverseLsp})
    midpoint.EntityData.Leafs = types.NewOrderedMap()
    midpoint.EntityData.Leafs.Append("midpoint-name", types.YLeaf{"MidpointName", midpoint.MidpointName})
    midpoint.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", midpoint.TunnelName})
    midpoint.EntityData.Leafs.Append("lsp-protect", types.YLeaf{"LspProtect", midpoint.LspProtect})
    midpoint.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", midpoint.LspId})

    midpoint.EntityData.YListKeys = []string {"MidpointName"}

    return &(midpoint.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint_Source
// Node identifier, tunnel identifier and
// optional global identifier of the source of
// the LSP
// This type is a presence type.
type MplsTe_TransportProfile_Midpoints_Midpoint_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Node identifier in IPv4 address format. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NodeId interface{}

    // Tunnel identifier in numeric value. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    TunnelId interface{}

    // Global identifier in numeric value. The type is interface{} with range:
    // 1..65535.
    GlobalId interface{}
}

func (source *MplsTe_TransportProfile_Midpoints_Midpoint_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "midpoint"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", source.NodeId})
    source.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", source.TunnelId})
    source.EntityData.Leafs.Append("global-id", types.YLeaf{"GlobalId", source.GlobalId})

    source.EntityData.YListKeys = []string {}

    return &(source.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint_Destination
// Node identifier, tunnel identifier and
// optional global identifier of the destination
// of the LSP
// This type is a presence type.
type MplsTe_TransportProfile_Midpoints_Midpoint_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Node identifier in IPv4 address format. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NodeId interface{}

    // Tunnel identifier in numeric value. The type is interface{} with range:
    // 0..65535. This attribute is mandatory.
    TunnelId interface{}

    // Global identifier in numeric value. The type is interface{} with range:
    // 1..65535.
    GlobalId interface{}
}

func (destination *MplsTe_TransportProfile_Midpoints_Midpoint_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "midpoint"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", destination.NodeId})
    destination.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", destination.TunnelId})
    destination.EntityData.Leafs.Append("global-id", types.YLeaf{"GlobalId", destination.GlobalId})

    destination.EntityData.YListKeys = []string {}

    return &(destination.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp
// Forward transport profile LSP
type MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth of forward transport profile LSP. The type is interface{} with
    // range: 0..4294967295. Units are kbit/s.
    ForwardBandwidth interface{}

    // Label cross-connect of forward transport profile LSP.
    ForwardIoMap MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp_ForwardIoMap
}

func (forwardLsp *MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp) GetEntityData() *types.CommonEntityData {
    forwardLsp.EntityData.YFilter = forwardLsp.YFilter
    forwardLsp.EntityData.YangName = "forward-lsp"
    forwardLsp.EntityData.BundleName = "cisco_ios_xr"
    forwardLsp.EntityData.ParentYangName = "midpoint"
    forwardLsp.EntityData.SegmentPath = "forward-lsp"
    forwardLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardLsp.EntityData.Children = types.NewOrderedMap()
    forwardLsp.EntityData.Children.Append("forward-io-map", types.YChild{"ForwardIoMap", &forwardLsp.ForwardIoMap})
    forwardLsp.EntityData.Leafs = types.NewOrderedMap()
    forwardLsp.EntityData.Leafs.Append("forward-bandwidth", types.YLeaf{"ForwardBandwidth", forwardLsp.ForwardBandwidth})

    forwardLsp.EntityData.YListKeys = []string {}

    return &(forwardLsp.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp_ForwardIoMap
// Label cross-connect of forward transport
// profile LSP
// This type is a presence type.
type MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp_ForwardIoMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // MPLS label. The type is interface{} with range: 16..4015.
    InLabel interface{}

    // Outgoing MPLS label. The type is interface{} with range: 16..1048575. This
    // attribute is mandatory.
    OutLabel interface{}

    // Transport profile identifier of outgoing link. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    OutLink interface{}
}

func (forwardIoMap *MplsTe_TransportProfile_Midpoints_Midpoint_ForwardLsp_ForwardIoMap) GetEntityData() *types.CommonEntityData {
    forwardIoMap.EntityData.YFilter = forwardIoMap.YFilter
    forwardIoMap.EntityData.YangName = "forward-io-map"
    forwardIoMap.EntityData.BundleName = "cisco_ios_xr"
    forwardIoMap.EntityData.ParentYangName = "forward-lsp"
    forwardIoMap.EntityData.SegmentPath = "forward-io-map"
    forwardIoMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwardIoMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwardIoMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwardIoMap.EntityData.Children = types.NewOrderedMap()
    forwardIoMap.EntityData.Leafs = types.NewOrderedMap()
    forwardIoMap.EntityData.Leafs.Append("in-label", types.YLeaf{"InLabel", forwardIoMap.InLabel})
    forwardIoMap.EntityData.Leafs.Append("out-label", types.YLeaf{"OutLabel", forwardIoMap.OutLabel})
    forwardIoMap.EntityData.Leafs.Append("out-link", types.YLeaf{"OutLink", forwardIoMap.OutLink})

    forwardIoMap.EntityData.YListKeys = []string {}

    return &(forwardIoMap.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp
// none
type MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth of reverse transport profile LSP. The type is interface{} with
    // range: 0..4294967295. Units are kbit/s.
    ReverseBandwidth interface{}

    // Label cross-connect of reverse transport profile LSP.
    ReverseIoMap MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp_ReverseIoMap
}

func (reverseLsp *MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp) GetEntityData() *types.CommonEntityData {
    reverseLsp.EntityData.YFilter = reverseLsp.YFilter
    reverseLsp.EntityData.YangName = "reverse-lsp"
    reverseLsp.EntityData.BundleName = "cisco_ios_xr"
    reverseLsp.EntityData.ParentYangName = "midpoint"
    reverseLsp.EntityData.SegmentPath = "reverse-lsp"
    reverseLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reverseLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reverseLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reverseLsp.EntityData.Children = types.NewOrderedMap()
    reverseLsp.EntityData.Children.Append("reverse-io-map", types.YChild{"ReverseIoMap", &reverseLsp.ReverseIoMap})
    reverseLsp.EntityData.Leafs = types.NewOrderedMap()
    reverseLsp.EntityData.Leafs.Append("reverse-bandwidth", types.YLeaf{"ReverseBandwidth", reverseLsp.ReverseBandwidth})

    reverseLsp.EntityData.YListKeys = []string {}

    return &(reverseLsp.EntityData)
}

// MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp_ReverseIoMap
// Label cross-connect of reverse transport
// profile LSP
// This type is a presence type.
type MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp_ReverseIoMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // MPLS label. The type is interface{} with range: 16..4015.
    InLabel interface{}

    // Outgoing MPLS label. The type is interface{} with range: 16..1048575. This
    // attribute is mandatory.
    OutLabel interface{}

    // Transport profile identifier of outgoing link. The type is interface{} with
    // range: 1..65535. This attribute is mandatory.
    OutLink interface{}
}

func (reverseIoMap *MplsTe_TransportProfile_Midpoints_Midpoint_ReverseLsp_ReverseIoMap) GetEntityData() *types.CommonEntityData {
    reverseIoMap.EntityData.YFilter = reverseIoMap.YFilter
    reverseIoMap.EntityData.YangName = "reverse-io-map"
    reverseIoMap.EntityData.BundleName = "cisco_ios_xr"
    reverseIoMap.EntityData.ParentYangName = "reverse-lsp"
    reverseIoMap.EntityData.SegmentPath = "reverse-io-map"
    reverseIoMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reverseIoMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reverseIoMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reverseIoMap.EntityData.Children = types.NewOrderedMap()
    reverseIoMap.EntityData.Leafs = types.NewOrderedMap()
    reverseIoMap.EntityData.Leafs.Append("in-label", types.YLeaf{"InLabel", reverseIoMap.InLabel})
    reverseIoMap.EntityData.Leafs.Append("out-label", types.YLeaf{"OutLabel", reverseIoMap.OutLabel})
    reverseIoMap.EntityData.Leafs.Append("out-link", types.YLeaf{"OutLink", reverseIoMap.OutLink})

    reverseIoMap.EntityData.YListKeys = []string {}

    return &(reverseIoMap.EntityData)
}

// MplsTe_Interfaces
// Configure MPLS TE interfaces
type MplsTe_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure an MPLS TE interface. The type is slice of
    // MplsTe_Interfaces_Interface.
    Interface []*MplsTe_Interfaces_Interface
}

func (interfaces *MplsTe_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "mpls-te"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsTe_Interfaces_Interface
// Configure an MPLS TE interface
type MplsTe_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // MPLS transport profile capable link.
    TransportProfileLink MplsTe_Interfaces_Interface_TransportProfileLink

    // LCAC specific MPLS interface configuration.
    Lcac MplsTe_Interfaces_Interface_Lcac

    // MPLS TE global interface configuration.
    GlobalAttributes MplsTe_Interfaces_Interface_GlobalAttributes
}

func (self *MplsTe_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("transport-profile-link", types.YChild{"TransportProfileLink", &self.TransportProfileLink})
    self.EntityData.Children.Append("lcac", types.YChild{"Lcac", &self.Lcac})
    self.EntityData.Children.Append("global-attributes", types.YChild{"GlobalAttributes", &self.GlobalAttributes})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// MplsTe_Interfaces_Interface_TransportProfileLink
// MPLS transport profile capable link
type MplsTe_Interfaces_Interface_TransportProfileLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport profile link table.
    Links MplsTe_Interfaces_Interface_TransportProfileLink_Links
}

func (transportProfileLink *MplsTe_Interfaces_Interface_TransportProfileLink) GetEntityData() *types.CommonEntityData {
    transportProfileLink.EntityData.YFilter = transportProfileLink.YFilter
    transportProfileLink.EntityData.YangName = "transport-profile-link"
    transportProfileLink.EntityData.BundleName = "cisco_ios_xr"
    transportProfileLink.EntityData.ParentYangName = "interface"
    transportProfileLink.EntityData.SegmentPath = "transport-profile-link"
    transportProfileLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transportProfileLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transportProfileLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transportProfileLink.EntityData.Children = types.NewOrderedMap()
    transportProfileLink.EntityData.Children.Append("links", types.YChild{"Links", &transportProfileLink.Links})
    transportProfileLink.EntityData.Leafs = types.NewOrderedMap()

    transportProfileLink.EntityData.YListKeys = []string {}

    return &(transportProfileLink.EntityData)
}

// MplsTe_Interfaces_Interface_TransportProfileLink_Links
// Transport profile link table
type MplsTe_Interfaces_Interface_TransportProfileLink_Links struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Transport profile link. The type is slice of
    // MplsTe_Interfaces_Interface_TransportProfileLink_Links_Link.
    Link []*MplsTe_Interfaces_Interface_TransportProfileLink_Links_Link
}

func (links *MplsTe_Interfaces_Interface_TransportProfileLink_Links) GetEntityData() *types.CommonEntityData {
    links.EntityData.YFilter = links.YFilter
    links.EntityData.YangName = "links"
    links.EntityData.BundleName = "cisco_ios_xr"
    links.EntityData.ParentYangName = "transport-profile-link"
    links.EntityData.SegmentPath = "links"
    links.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    links.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    links.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    links.EntityData.Children = types.NewOrderedMap()
    links.EntityData.Children.Append("link", types.YChild{"Link", nil})
    for i := range links.Link {
        links.EntityData.Children.Append(types.GetSegmentPath(links.Link[i]), types.YChild{"Link", links.Link[i]})
    }
    links.EntityData.Leafs = types.NewOrderedMap()

    links.EntityData.YListKeys = []string {}

    return &(links.EntityData)
}

// MplsTe_Interfaces_Interface_TransportProfileLink_Links_Link
// Transport profile link
type MplsTe_Interfaces_Interface_TransportProfileLink_Links_Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Numeric link identifier. The type is interface{}
    // with range: 1..65535.
    LinkId interface{}

    // Next hop type. The type is LinkNextHop. The default value is ipv4-address.
    NextHopType interface{}

    // Next-hop address in IPv4 format. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopAddress interface{}
}

func (link *MplsTe_Interfaces_Interface_TransportProfileLink_Links_Link) GetEntityData() *types.CommonEntityData {
    link.EntityData.YFilter = link.YFilter
    link.EntityData.YangName = "link"
    link.EntityData.BundleName = "cisco_ios_xr"
    link.EntityData.ParentYangName = "links"
    link.EntityData.SegmentPath = "link" + types.AddKeyToken(link.LinkId, "link-id")
    link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    link.EntityData.Children = types.NewOrderedMap()
    link.EntityData.Leafs = types.NewOrderedMap()
    link.EntityData.Leafs.Append("link-id", types.YLeaf{"LinkId", link.LinkId})
    link.EntityData.Leafs.Append("next-hop-type", types.YLeaf{"NextHopType", link.NextHopType})
    link.EntityData.Leafs.Append("next-hop-address", types.YLeaf{"NextHopAddress", link.NextHopAddress})

    link.EntityData.YListKeys = []string {"LinkId"}

    return &(link.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac
// LCAC specific MPLS interface configuration
type MplsTe_Interfaces_Interface_Lcac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable use of Bidirectional Forwarding Detection. The type is interface{}.
    Bfd interface{}

    // Lockout protection on the interface for Flex LSP. The type is interface{}.
    FaultOamLockout interface{}

    // Set user defined interface attribute flags. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    AttributeFlags interface{}

    // Enable MPLS-TE on the link. The type is interface{}.
    Enable interface{}

    // Set administrative weight for the interface. The type is interface{} with
    // range: 0..4294967295.
    AdminWeight interface{}

    // Set the te-link switching attributes.
    Switchings MplsTe_Interfaces_Interface_Lcac_Switchings

    // Set the IGP instance into which this interface is to be flooded (GMPLS
    // only).
    FloodArea MplsTe_Interfaces_Interface_Lcac_FloodArea

    // Set the interface attribute names.
    AttributeNameXr MplsTe_Interfaces_Interface_Lcac_AttributeNameXr

    // Attribute name table.
    AttributeNames MplsTe_Interfaces_Interface_Lcac_AttributeNames

    // Configure SRLG membership for the interface.
    Srlgs MplsTe_Interfaces_Interface_Lcac_Srlgs

    // Set thresholds for increased resource availability in %.
    UpThresholds MplsTe_Interfaces_Interface_Lcac_UpThresholds

    // Set thresholds for decreased resource availability in %.
    DownThresholds MplsTe_Interfaces_Interface_Lcac_DownThresholds
}

func (lcac *MplsTe_Interfaces_Interface_Lcac) GetEntityData() *types.CommonEntityData {
    lcac.EntityData.YFilter = lcac.YFilter
    lcac.EntityData.YangName = "lcac"
    lcac.EntityData.BundleName = "cisco_ios_xr"
    lcac.EntityData.ParentYangName = "interface"
    lcac.EntityData.SegmentPath = "lcac"
    lcac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcac.EntityData.Children = types.NewOrderedMap()
    lcac.EntityData.Children.Append("switchings", types.YChild{"Switchings", &lcac.Switchings})
    lcac.EntityData.Children.Append("flood-area", types.YChild{"FloodArea", &lcac.FloodArea})
    lcac.EntityData.Children.Append("attribute-name-xr", types.YChild{"AttributeNameXr", &lcac.AttributeNameXr})
    lcac.EntityData.Children.Append("attribute-names", types.YChild{"AttributeNames", &lcac.AttributeNames})
    lcac.EntityData.Children.Append("srlgs", types.YChild{"Srlgs", &lcac.Srlgs})
    lcac.EntityData.Children.Append("up-thresholds", types.YChild{"UpThresholds", &lcac.UpThresholds})
    lcac.EntityData.Children.Append("down-thresholds", types.YChild{"DownThresholds", &lcac.DownThresholds})
    lcac.EntityData.Leafs = types.NewOrderedMap()
    lcac.EntityData.Leafs.Append("bfd", types.YLeaf{"Bfd", lcac.Bfd})
    lcac.EntityData.Leafs.Append("fault-oam-lockout", types.YLeaf{"FaultOamLockout", lcac.FaultOamLockout})
    lcac.EntityData.Leafs.Append("attribute-flags", types.YLeaf{"AttributeFlags", lcac.AttributeFlags})
    lcac.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", lcac.Enable})
    lcac.EntityData.Leafs.Append("admin-weight", types.YLeaf{"AdminWeight", lcac.AdminWeight})

    lcac.EntityData.YListKeys = []string {}

    return &(lcac.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_Switchings
// Set the te-link switching attributes
type MplsTe_Interfaces_Interface_Lcac_Switchings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The te-link switching attributes. The type is slice of
    // MplsTe_Interfaces_Interface_Lcac_Switchings_Switching.
    Switching []*MplsTe_Interfaces_Interface_Lcac_Switchings_Switching
}

func (switchings *MplsTe_Interfaces_Interface_Lcac_Switchings) GetEntityData() *types.CommonEntityData {
    switchings.EntityData.YFilter = switchings.YFilter
    switchings.EntityData.YangName = "switchings"
    switchings.EntityData.BundleName = "cisco_ios_xr"
    switchings.EntityData.ParentYangName = "lcac"
    switchings.EntityData.SegmentPath = "switchings"
    switchings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switchings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switchings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switchings.EntityData.Children = types.NewOrderedMap()
    switchings.EntityData.Children.Append("switching", types.YChild{"Switching", nil})
    for i := range switchings.Switching {
        switchings.EntityData.Children.Append(types.GetSegmentPath(switchings.Switching[i]), types.YChild{"Switching", switchings.Switching[i]})
    }
    switchings.EntityData.Leafs = types.NewOrderedMap()

    switchings.EntityData.YListKeys = []string {}

    return &(switchings.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_Switchings_Switching
// The te-link switching attributes
type MplsTe_Interfaces_Interface_Lcac_Switchings_Switching struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Switching index. The type is one of the following
    // types: enumeration MplsTeSwitchingIndex, or int with range: 1..255.
    SwitchingId interface{}

    // Set the local encoding type. The type is MplsTeSwitchingEncoding.
    Encoding interface{}

    // Set the local switching capability. The type is MplsTeSwitchingCap.
    Capability interface{}
}

func (switching *MplsTe_Interfaces_Interface_Lcac_Switchings_Switching) GetEntityData() *types.CommonEntityData {
    switching.EntityData.YFilter = switching.YFilter
    switching.EntityData.YangName = "switching"
    switching.EntityData.BundleName = "cisco_ios_xr"
    switching.EntityData.ParentYangName = "switchings"
    switching.EntityData.SegmentPath = "switching" + types.AddKeyToken(switching.SwitchingId, "switching-id")
    switching.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    switching.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    switching.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    switching.EntityData.Children = types.NewOrderedMap()
    switching.EntityData.Leafs = types.NewOrderedMap()
    switching.EntityData.Leafs.Append("switching-id", types.YLeaf{"SwitchingId", switching.SwitchingId})
    switching.EntityData.Leafs.Append("encoding", types.YLeaf{"Encoding", switching.Encoding})
    switching.EntityData.Leafs.Append("capability", types.YLeaf{"Capability", switching.Capability})

    switching.EntityData.YListKeys = []string {"SwitchingId"}

    return &(switching.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_FloodArea
// Set the IGP instance into which this
// interface is to be flooded (GMPLS only)
type MplsTe_Interfaces_Interface_Lcac_FloodArea struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP type. The type is MplsLcacFloodingIgp.
    IgpType interface{}

    // Process name. The type is string with length: 1..32.
    ProcessName interface{}

    // Area ID. The type is interface{} with range: 0..4294967295.
    AreaId interface{}
}

func (floodArea *MplsTe_Interfaces_Interface_Lcac_FloodArea) GetEntityData() *types.CommonEntityData {
    floodArea.EntityData.YFilter = floodArea.YFilter
    floodArea.EntityData.YangName = "flood-area"
    floodArea.EntityData.BundleName = "cisco_ios_xr"
    floodArea.EntityData.ParentYangName = "lcac"
    floodArea.EntityData.SegmentPath = "flood-area"
    floodArea.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    floodArea.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    floodArea.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    floodArea.EntityData.Children = types.NewOrderedMap()
    floodArea.EntityData.Leafs = types.NewOrderedMap()
    floodArea.EntityData.Leafs.Append("igp-type", types.YLeaf{"IgpType", floodArea.IgpType})
    floodArea.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", floodArea.ProcessName})
    floodArea.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", floodArea.AreaId})

    floodArea.EntityData.YListKeys = []string {}

    return &(floodArea.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_AttributeNameXr
// Set the interface attribute names
type MplsTe_Interfaces_Interface_Lcac_AttributeNameXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of Attribute Names. The type is slice of string.
    AttributeName []interface{}
}

func (attributeNameXr *MplsTe_Interfaces_Interface_Lcac_AttributeNameXr) GetEntityData() *types.CommonEntityData {
    attributeNameXr.EntityData.YFilter = attributeNameXr.YFilter
    attributeNameXr.EntityData.YangName = "attribute-name-xr"
    attributeNameXr.EntityData.BundleName = "cisco_ios_xr"
    attributeNameXr.EntityData.ParentYangName = "lcac"
    attributeNameXr.EntityData.SegmentPath = "attribute-name-xr"
    attributeNameXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeNameXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeNameXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeNameXr.EntityData.Children = types.NewOrderedMap()
    attributeNameXr.EntityData.Leafs = types.NewOrderedMap()
    attributeNameXr.EntityData.Leafs.Append("attribute-name", types.YLeaf{"AttributeName", attributeNameXr.AttributeName})

    attributeNameXr.EntityData.YListKeys = []string {}

    return &(attributeNameXr.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_AttributeNames
// Attribute name table
type MplsTe_Interfaces_Interface_Lcac_AttributeNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the interface attribute names. The type is slice of
    // MplsTe_Interfaces_Interface_Lcac_AttributeNames_AttributeName.
    AttributeName []*MplsTe_Interfaces_Interface_Lcac_AttributeNames_AttributeName
}

func (attributeNames *MplsTe_Interfaces_Interface_Lcac_AttributeNames) GetEntityData() *types.CommonEntityData {
    attributeNames.EntityData.YFilter = attributeNames.YFilter
    attributeNames.EntityData.YangName = "attribute-names"
    attributeNames.EntityData.BundleName = "cisco_ios_xr"
    attributeNames.EntityData.ParentYangName = "lcac"
    attributeNames.EntityData.SegmentPath = "attribute-names"
    attributeNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeNames.EntityData.Children = types.NewOrderedMap()
    attributeNames.EntityData.Children.Append("attribute-name", types.YChild{"AttributeName", nil})
    for i := range attributeNames.AttributeName {
        attributeNames.EntityData.Children.Append(types.GetSegmentPath(attributeNames.AttributeName[i]), types.YChild{"AttributeName", attributeNames.AttributeName[i]})
    }
    attributeNames.EntityData.Leafs = types.NewOrderedMap()

    attributeNames.EntityData.YListKeys = []string {}

    return &(attributeNames.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_AttributeNames_AttributeName
// Set the interface attribute names
type MplsTe_Interfaces_Interface_Lcac_AttributeNames_AttributeName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the entry index. The type is interface{}
    // with range: 1..9.
    AffinityIndex interface{}

    // Array of Attribute Names. The type is slice of string.
    Value []interface{}
}

func (attributeName *MplsTe_Interfaces_Interface_Lcac_AttributeNames_AttributeName) GetEntityData() *types.CommonEntityData {
    attributeName.EntityData.YFilter = attributeName.YFilter
    attributeName.EntityData.YangName = "attribute-name"
    attributeName.EntityData.BundleName = "cisco_ios_xr"
    attributeName.EntityData.ParentYangName = "attribute-names"
    attributeName.EntityData.SegmentPath = "attribute-name" + types.AddKeyToken(attributeName.AffinityIndex, "affinity-index")
    attributeName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    attributeName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    attributeName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    attributeName.EntityData.Children = types.NewOrderedMap()
    attributeName.EntityData.Leafs = types.NewOrderedMap()
    attributeName.EntityData.Leafs.Append("affinity-index", types.YLeaf{"AffinityIndex", attributeName.AffinityIndex})
    attributeName.EntityData.Leafs.Append("value", types.YLeaf{"Value", attributeName.Value})

    attributeName.EntityData.YListKeys = []string {"AffinityIndex"}

    return &(attributeName.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_Srlgs
// Configure SRLG membership for the interface
type MplsTe_Interfaces_Interface_Lcac_Srlgs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRLG membership number. The type is slice of
    // MplsTe_Interfaces_Interface_Lcac_Srlgs_Srlg.
    Srlg []*MplsTe_Interfaces_Interface_Lcac_Srlgs_Srlg
}

func (srlgs *MplsTe_Interfaces_Interface_Lcac_Srlgs) GetEntityData() *types.CommonEntityData {
    srlgs.EntityData.YFilter = srlgs.YFilter
    srlgs.EntityData.YangName = "srlgs"
    srlgs.EntityData.BundleName = "cisco_ios_xr"
    srlgs.EntityData.ParentYangName = "lcac"
    srlgs.EntityData.SegmentPath = "srlgs"
    srlgs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlgs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlgs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlgs.EntityData.Children = types.NewOrderedMap()
    srlgs.EntityData.Children.Append("srlg", types.YChild{"Srlg", nil})
    for i := range srlgs.Srlg {
        srlgs.EntityData.Children.Append(types.GetSegmentPath(srlgs.Srlg[i]), types.YChild{"Srlg", srlgs.Srlg[i]})
    }
    srlgs.EntityData.Leafs = types.NewOrderedMap()

    srlgs.EntityData.YListKeys = []string {}

    return &(srlgs.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_Srlgs_Srlg
// SRLG membership number
type MplsTe_Interfaces_Interface_Lcac_Srlgs_Srlg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SRLG membership number. The type is interface{}
    // with range: 0..4294967295.
    SrlgNumber interface{}
}

func (srlg *MplsTe_Interfaces_Interface_Lcac_Srlgs_Srlg) GetEntityData() *types.CommonEntityData {
    srlg.EntityData.YFilter = srlg.YFilter
    srlg.EntityData.YangName = "srlg"
    srlg.EntityData.BundleName = "cisco_ios_xr"
    srlg.EntityData.ParentYangName = "srlgs"
    srlg.EntityData.SegmentPath = "srlg" + types.AddKeyToken(srlg.SrlgNumber, "srlg-number")
    srlg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srlg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srlg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srlg.EntityData.Children = types.NewOrderedMap()
    srlg.EntityData.Leafs = types.NewOrderedMap()
    srlg.EntityData.Leafs.Append("srlg-number", types.YLeaf{"SrlgNumber", srlg.SrlgNumber})

    srlg.EntityData.YListKeys = []string {"SrlgNumber"}

    return &(srlg.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_UpThresholds
// Set thresholds for increased resource
// availability in %
type MplsTe_Interfaces_Interface_Lcac_UpThresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of up threshold percentage. The type is slice of interface{} with
    // range: 0..100. Units are percentage.
    UpThreshold []interface{}
}

func (upThresholds *MplsTe_Interfaces_Interface_Lcac_UpThresholds) GetEntityData() *types.CommonEntityData {
    upThresholds.EntityData.YFilter = upThresholds.YFilter
    upThresholds.EntityData.YangName = "up-thresholds"
    upThresholds.EntityData.BundleName = "cisco_ios_xr"
    upThresholds.EntityData.ParentYangName = "lcac"
    upThresholds.EntityData.SegmentPath = "up-thresholds"
    upThresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    upThresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    upThresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    upThresholds.EntityData.Children = types.NewOrderedMap()
    upThresholds.EntityData.Leafs = types.NewOrderedMap()
    upThresholds.EntityData.Leafs.Append("up-threshold", types.YLeaf{"UpThreshold", upThresholds.UpThreshold})

    upThresholds.EntityData.YListKeys = []string {}

    return &(upThresholds.EntityData)
}

// MplsTe_Interfaces_Interface_Lcac_DownThresholds
// Set thresholds for decreased resource
// availability in %
type MplsTe_Interfaces_Interface_Lcac_DownThresholds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of down threshold percentage. The type is slice of interface{} with
    // range: 0..100. Units are percentage.
    DownThreshold []interface{}
}

func (downThresholds *MplsTe_Interfaces_Interface_Lcac_DownThresholds) GetEntityData() *types.CommonEntityData {
    downThresholds.EntityData.YFilter = downThresholds.YFilter
    downThresholds.EntityData.YangName = "down-thresholds"
    downThresholds.EntityData.BundleName = "cisco_ios_xr"
    downThresholds.EntityData.ParentYangName = "lcac"
    downThresholds.EntityData.SegmentPath = "down-thresholds"
    downThresholds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    downThresholds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    downThresholds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    downThresholds.EntityData.Children = types.NewOrderedMap()
    downThresholds.EntityData.Leafs = types.NewOrderedMap()
    downThresholds.EntityData.Leafs.Append("down-threshold", types.YLeaf{"DownThreshold", downThresholds.DownThreshold})

    downThresholds.EntityData.YListKeys = []string {}

    return &(downThresholds.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes
// MPLS TE global interface configuration
type MplsTe_Interfaces_Interface_GlobalAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure MPLS TE backup tunnels for this interface.
    BackupTunnels MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels

    // Auto tunnel configuration.
    AutoTunnel MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel

    // Configure MPLS TE backup tunnels for this interface.
    BackupPaths MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths
}

func (globalAttributes *MplsTe_Interfaces_Interface_GlobalAttributes) GetEntityData() *types.CommonEntityData {
    globalAttributes.EntityData.YFilter = globalAttributes.YFilter
    globalAttributes.EntityData.YangName = "global-attributes"
    globalAttributes.EntityData.BundleName = "cisco_ios_xr"
    globalAttributes.EntityData.ParentYangName = "interface"
    globalAttributes.EntityData.SegmentPath = "global-attributes"
    globalAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalAttributes.EntityData.Children = types.NewOrderedMap()
    globalAttributes.EntityData.Children.Append("backup-tunnels", types.YChild{"BackupTunnels", &globalAttributes.BackupTunnels})
    globalAttributes.EntityData.Children.Append("auto-tunnel", types.YChild{"AutoTunnel", &globalAttributes.AutoTunnel})
    globalAttributes.EntityData.Children.Append("backup-paths", types.YChild{"BackupPaths", &globalAttributes.BackupPaths})
    globalAttributes.EntityData.Leafs = types.NewOrderedMap()

    globalAttributes.EntityData.YListKeys = []string {}

    return &(globalAttributes.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels
// Configure MPLS TE backup tunnels for this
// interface
type MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel name. The type is slice of
    // MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels_BackupTunnel.
    BackupTunnel []*MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels_BackupTunnel
}

func (backupTunnels *MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels) GetEntityData() *types.CommonEntityData {
    backupTunnels.EntityData.YFilter = backupTunnels.YFilter
    backupTunnels.EntityData.YangName = "backup-tunnels"
    backupTunnels.EntityData.BundleName = "cisco_ios_xr"
    backupTunnels.EntityData.ParentYangName = "global-attributes"
    backupTunnels.EntityData.SegmentPath = "backup-tunnels"
    backupTunnels.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupTunnels.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupTunnels.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupTunnels.EntityData.Children = types.NewOrderedMap()
    backupTunnels.EntityData.Children.Append("backup-tunnel", types.YChild{"BackupTunnel", nil})
    for i := range backupTunnels.BackupTunnel {
        backupTunnels.EntityData.Children.Append(types.GetSegmentPath(backupTunnels.BackupTunnel[i]), types.YChild{"BackupTunnel", backupTunnels.BackupTunnel[i]})
    }
    backupTunnels.EntityData.Leafs = types.NewOrderedMap()

    backupTunnels.EntityData.YListKeys = []string {}

    return &(backupTunnels.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels_BackupTunnel
// Tunnel name
type MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels_BackupTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Tunnel name. The type is string with length:
    // 1..59.
    TunnelName interface{}
}

func (backupTunnel *MplsTe_Interfaces_Interface_GlobalAttributes_BackupTunnels_BackupTunnel) GetEntityData() *types.CommonEntityData {
    backupTunnel.EntityData.YFilter = backupTunnel.YFilter
    backupTunnel.EntityData.YangName = "backup-tunnel"
    backupTunnel.EntityData.BundleName = "cisco_ios_xr"
    backupTunnel.EntityData.ParentYangName = "backup-tunnels"
    backupTunnel.EntityData.SegmentPath = "backup-tunnel" + types.AddKeyToken(backupTunnel.TunnelName, "tunnel-name")
    backupTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupTunnel.EntityData.Children = types.NewOrderedMap()
    backupTunnel.EntityData.Leafs = types.NewOrderedMap()
    backupTunnel.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", backupTunnel.TunnelName})

    backupTunnel.EntityData.YListKeys = []string {"TunnelName"}

    return &(backupTunnel.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel
// Auto tunnel configuration
type MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Auto tunnel backup configuration.
    Backup MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup
}

func (autoTunnel *MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel) GetEntityData() *types.CommonEntityData {
    autoTunnel.EntityData.YFilter = autoTunnel.YFilter
    autoTunnel.EntityData.YangName = "auto-tunnel"
    autoTunnel.EntityData.BundleName = "cisco_ios_xr"
    autoTunnel.EntityData.ParentYangName = "global-attributes"
    autoTunnel.EntityData.SegmentPath = "auto-tunnel"
    autoTunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoTunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoTunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoTunnel.EntityData.Children = types.NewOrderedMap()
    autoTunnel.EntityData.Children.Append("backup", types.YChild{"Backup", &autoTunnel.Backup})
    autoTunnel.EntityData.Leafs = types.NewOrderedMap()

    autoTunnel.EntityData.YListKeys = []string {}

    return &(autoTunnel.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup
// Auto tunnel backup configuration
type MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable auto-tunnel backup on this TE link. The type is interface{}.
    Enable interface{}

    // The name of attribute set to be applied to this auto backup lsp. The type
    // is string with length: 1..64.
    AttributeSet interface{}

    // Enable NHOP-only mode for auto-tunnel backup on this TE link. The type is
    // interface{}.
    NextHopOnly interface{}

    // Auto-tunnel backup exclusion criteria.
    Exclude MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup_Exclude
}

func (backup *MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup) GetEntityData() *types.CommonEntityData {
    backup.EntityData.YFilter = backup.YFilter
    backup.EntityData.YangName = "backup"
    backup.EntityData.BundleName = "cisco_ios_xr"
    backup.EntityData.ParentYangName = "auto-tunnel"
    backup.EntityData.SegmentPath = "backup"
    backup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backup.EntityData.Children = types.NewOrderedMap()
    backup.EntityData.Children.Append("exclude", types.YChild{"Exclude", &backup.Exclude})
    backup.EntityData.Leafs = types.NewOrderedMap()
    backup.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", backup.Enable})
    backup.EntityData.Leafs.Append("attribute-set", types.YLeaf{"AttributeSet", backup.AttributeSet})
    backup.EntityData.Leafs.Append("next-hop-only", types.YLeaf{"NextHopOnly", backup.NextHopOnly})

    backup.EntityData.YListKeys = []string {}

    return &(backup.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup_Exclude
// Auto-tunnel backup exclusion criteria
type MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup_Exclude struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set exclude SRLG mode for auto-tunnel backup on this TE link. The type is
    // MplsTesrlgExclude.
    SrlgMode interface{}
}

func (exclude *MplsTe_Interfaces_Interface_GlobalAttributes_AutoTunnel_Backup_Exclude) GetEntityData() *types.CommonEntityData {
    exclude.EntityData.YFilter = exclude.YFilter
    exclude.EntityData.YangName = "exclude"
    exclude.EntityData.BundleName = "cisco_ios_xr"
    exclude.EntityData.ParentYangName = "backup"
    exclude.EntityData.SegmentPath = "exclude"
    exclude.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    exclude.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    exclude.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    exclude.EntityData.Children = types.NewOrderedMap()
    exclude.EntityData.Leafs = types.NewOrderedMap()
    exclude.EntityData.Leafs.Append("srlg-mode", types.YLeaf{"SrlgMode", exclude.SrlgMode})

    exclude.EntityData.YListKeys = []string {}

    return &(exclude.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths
// Configure MPLS TE backup tunnels for this
// interface
type MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel interface number. The type is slice of
    // MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths_BackupPath.
    BackupPath []*MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths_BackupPath
}

func (backupPaths *MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths) GetEntityData() *types.CommonEntityData {
    backupPaths.EntityData.YFilter = backupPaths.YFilter
    backupPaths.EntityData.YangName = "backup-paths"
    backupPaths.EntityData.BundleName = "cisco_ios_xr"
    backupPaths.EntityData.ParentYangName = "global-attributes"
    backupPaths.EntityData.SegmentPath = "backup-paths"
    backupPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPaths.EntityData.Children = types.NewOrderedMap()
    backupPaths.EntityData.Children.Append("backup-path", types.YChild{"BackupPath", nil})
    for i := range backupPaths.BackupPath {
        backupPaths.EntityData.Children.Append(types.GetSegmentPath(backupPaths.BackupPath[i]), types.YChild{"BackupPath", backupPaths.BackupPath[i]})
    }
    backupPaths.EntityData.Leafs = types.NewOrderedMap()

    backupPaths.EntityData.YListKeys = []string {}

    return &(backupPaths.EntityData)
}

// MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths_BackupPath
// Tunnel interface number
type MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths_BackupPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Tunnel interface number. The type is interface{}
    // with range: 0..65535.
    TunnelNumber interface{}
}

func (backupPath *MplsTe_Interfaces_Interface_GlobalAttributes_BackupPaths_BackupPath) GetEntityData() *types.CommonEntityData {
    backupPath.EntityData.YFilter = backupPath.YFilter
    backupPath.EntityData.YangName = "backup-path"
    backupPath.EntityData.BundleName = "cisco_ios_xr"
    backupPath.EntityData.ParentYangName = "backup-paths"
    backupPath.EntityData.SegmentPath = "backup-path" + types.AddKeyToken(backupPath.TunnelNumber, "tunnel-number")
    backupPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    backupPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    backupPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    backupPath.EntityData.Children = types.NewOrderedMap()
    backupPath.EntityData.Leafs = types.NewOrderedMap()
    backupPath.EntityData.Leafs.Append("tunnel-number", types.YLeaf{"TunnelNumber", backupPath.TunnelNumber})

    backupPath.EntityData.YListKeys = []string {"TunnelNumber"}

    return &(backupPath.EntityData)
}

// MplsTe_GmplsNni
// GMPLS-NNI configuration
type MplsTe_GmplsNni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path selection configuration for all gmpls nni tunnels. The type is
    // MplsTePathSelectionMetric.
    PathSelectionMetric interface{}

    // Enable MPLS Traffic Engineering GMPLS-NNI. The type is interface{}.
    EnableGmplsNni interface{}

    // GMPLS-NNI topology instance table.
    TopologyInstances MplsTe_GmplsNni_TopologyInstances

    // GMPLS-NNI tunnel-head table.
    TunnelHeads MplsTe_GmplsNni_TunnelHeads
}

func (gmplsNni *MplsTe_GmplsNni) GetEntityData() *types.CommonEntityData {
    gmplsNni.EntityData.YFilter = gmplsNni.YFilter
    gmplsNni.EntityData.YangName = "gmpls-nni"
    gmplsNni.EntityData.BundleName = "cisco_ios_xr"
    gmplsNni.EntityData.ParentYangName = "mpls-te"
    gmplsNni.EntityData.SegmentPath = "gmpls-nni"
    gmplsNni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmplsNni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmplsNni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmplsNni.EntityData.Children = types.NewOrderedMap()
    gmplsNni.EntityData.Children.Append("topology-instances", types.YChild{"TopologyInstances", &gmplsNni.TopologyInstances})
    gmplsNni.EntityData.Children.Append("tunnel-heads", types.YChild{"TunnelHeads", &gmplsNni.TunnelHeads})
    gmplsNni.EntityData.Leafs = types.NewOrderedMap()
    gmplsNni.EntityData.Leafs.Append("path-selection-metric", types.YLeaf{"PathSelectionMetric", gmplsNni.PathSelectionMetric})
    gmplsNni.EntityData.Leafs.Append("enable-gmpls-nni", types.YLeaf{"EnableGmplsNni", gmplsNni.EnableGmplsNni})

    gmplsNni.EntityData.YListKeys = []string {}

    return &(gmplsNni.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances
// GMPLS-NNI topology instance table
type MplsTe_GmplsNni_TopologyInstances struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // GMPLS-NNI topology instance configuration. The type is slice of
    // MplsTe_GmplsNni_TopologyInstances_TopologyInstance.
    TopologyInstance []*MplsTe_GmplsNni_TopologyInstances_TopologyInstance
}

func (topologyInstances *MplsTe_GmplsNni_TopologyInstances) GetEntityData() *types.CommonEntityData {
    topologyInstances.EntityData.YFilter = topologyInstances.YFilter
    topologyInstances.EntityData.YangName = "topology-instances"
    topologyInstances.EntityData.BundleName = "cisco_ios_xr"
    topologyInstances.EntityData.ParentYangName = "gmpls-nni"
    topologyInstances.EntityData.SegmentPath = "topology-instances"
    topologyInstances.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyInstances.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyInstances.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyInstances.EntityData.Children = types.NewOrderedMap()
    topologyInstances.EntityData.Children.Append("topology-instance", types.YChild{"TopologyInstance", nil})
    for i := range topologyInstances.TopologyInstance {
        topologyInstances.EntityData.Children.Append(types.GetSegmentPath(topologyInstances.TopologyInstance[i]), types.YChild{"TopologyInstance", topologyInstances.TopologyInstance[i]})
    }
    topologyInstances.EntityData.Leafs = types.NewOrderedMap()

    topologyInstances.EntityData.YListKeys = []string {}

    return &(topologyInstances.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance
// GMPLS-NNI topology instance configuration
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OSPF area format. The type is OspfAreaMode.
    OspfAreaType interface{}

    // This attribute is a key. Name of IGP instance. The type is string with
    // length: 1..40.
    IgpInstanceName interface{}

    // This attribute is a key. IGP type. The type is MplsTeIgpProtocol.
    IgpType interface{}

    // ospf int. The type is slice of
    // MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt.
    OspfInt []*MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt

    // ospfip addr. The type is slice of
    // MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr.
    OspfipAddr []*MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr
}

func (topologyInstance *MplsTe_GmplsNni_TopologyInstances_TopologyInstance) GetEntityData() *types.CommonEntityData {
    topologyInstance.EntityData.YFilter = topologyInstance.YFilter
    topologyInstance.EntityData.YangName = "topology-instance"
    topologyInstance.EntityData.BundleName = "cisco_ios_xr"
    topologyInstance.EntityData.ParentYangName = "topology-instances"
    topologyInstance.EntityData.SegmentPath = "topology-instance" + types.AddKeyToken(topologyInstance.OspfAreaType, "ospf-area-type") + types.AddKeyToken(topologyInstance.IgpInstanceName, "igp-instance-name") + types.AddKeyToken(topologyInstance.IgpType, "igp-type")
    topologyInstance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyInstance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyInstance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyInstance.EntityData.Children = types.NewOrderedMap()
    topologyInstance.EntityData.Children.Append("ospf-int", types.YChild{"OspfInt", nil})
    for i := range topologyInstance.OspfInt {
        topologyInstance.EntityData.Children.Append(types.GetSegmentPath(topologyInstance.OspfInt[i]), types.YChild{"OspfInt", topologyInstance.OspfInt[i]})
    }
    topologyInstance.EntityData.Children.Append("ospfip-addr", types.YChild{"OspfipAddr", nil})
    for i := range topologyInstance.OspfipAddr {
        topologyInstance.EntityData.Children.Append(types.GetSegmentPath(topologyInstance.OspfipAddr[i]), types.YChild{"OspfipAddr", topologyInstance.OspfipAddr[i]})
    }
    topologyInstance.EntityData.Leafs = types.NewOrderedMap()
    topologyInstance.EntityData.Leafs.Append("ospf-area-type", types.YLeaf{"OspfAreaType", topologyInstance.OspfAreaType})
    topologyInstance.EntityData.Leafs.Append("igp-instance-name", types.YLeaf{"IgpInstanceName", topologyInstance.IgpInstanceName})
    topologyInstance.EntityData.Leafs.Append("igp-type", types.YLeaf{"IgpType", topologyInstance.IgpType})

    topologyInstance.EntityData.YListKeys = []string {"OspfAreaType", "IgpInstanceName", "IgpType"}

    return &(topologyInstance.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt
// ospf int
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IGP area. The type is interface{} with range:
    // 0..4294967295.
    IgpArea interface{}

    // GMPLS-NNI controllers.
    Controllers MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers
}

func (ospfInt *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt) GetEntityData() *types.CommonEntityData {
    ospfInt.EntityData.YFilter = ospfInt.YFilter
    ospfInt.EntityData.YangName = "ospf-int"
    ospfInt.EntityData.BundleName = "cisco_ios_xr"
    ospfInt.EntityData.ParentYangName = "topology-instance"
    ospfInt.EntityData.SegmentPath = "ospf-int" + types.AddKeyToken(ospfInt.IgpArea, "igp-area")
    ospfInt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfInt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfInt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfInt.EntityData.Children = types.NewOrderedMap()
    ospfInt.EntityData.Children.Append("controllers", types.YChild{"Controllers", &ospfInt.Controllers})
    ospfInt.EntityData.Leafs = types.NewOrderedMap()
    ospfInt.EntityData.Leafs.Append("igp-area", types.YLeaf{"IgpArea", ospfInt.IgpArea})

    ospfInt.EntityData.YListKeys = []string {"IgpArea"}

    return &(ospfInt.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers
// GMPLS-NNI controllers
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a GMPLS NNI controller. The type is slice of
    // MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller.
    Controller []*MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller
}

func (controllers *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "ospf-int"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller
// Configure a GMPLS NNI controller
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // Set administrative weight for the interface. The type is interface{} with
    // range: 0..65535.
    AdminWeight interface{}

    // Enable GMPLS-NNI on the link. The type is interface{}.
    Enable interface{}

    // Set link delay for the interface. The type is interface{} with range:
    // 1..16777215.
    Delay interface{}

    // Set tandem connection monitoring for the interface.
    TtiMode MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_TtiMode

    // Set the interface affinity names.
    AffinityNames MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_AffinityNames
}

func (controller *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("tti-mode", types.YChild{"TtiMode", &controller.TtiMode})
    controller.EntityData.Children.Append("affinity-names", types.YChild{"AffinityNames", &controller.AffinityNames})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})
    controller.EntityData.Leafs.Append("admin-weight", types.YLeaf{"AdminWeight", controller.AdminWeight})
    controller.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", controller.Enable})
    controller.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", controller.Delay})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_TtiMode
// Set tandem connection monitoring for the
// interface
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_TtiMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of Trail Trace Identifier. The type is GmplsttiMode.
    TtiModeType interface{}

    // Tandem Connection Monitoring ID for the interface. The type is interface{}
    // with range: 1..6.
    Tcmid interface{}
}

func (ttiMode *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_TtiMode) GetEntityData() *types.CommonEntityData {
    ttiMode.EntityData.YFilter = ttiMode.YFilter
    ttiMode.EntityData.YangName = "tti-mode"
    ttiMode.EntityData.BundleName = "cisco_ios_xr"
    ttiMode.EntityData.ParentYangName = "controller"
    ttiMode.EntityData.SegmentPath = "tti-mode"
    ttiMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttiMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttiMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttiMode.EntityData.Children = types.NewOrderedMap()
    ttiMode.EntityData.Leafs = types.NewOrderedMap()
    ttiMode.EntityData.Leafs.Append("tti-mode-type", types.YLeaf{"TtiModeType", ttiMode.TtiModeType})
    ttiMode.EntityData.Leafs.Append("tcmid", types.YLeaf{"Tcmid", ttiMode.Tcmid})

    ttiMode.EntityData.YListKeys = []string {}

    return &(ttiMode.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_AffinityNames
// Set the interface affinity names
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_AffinityNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of Attribute Names. The type is slice of string.
    AffinityName []interface{}
}

func (affinityNames *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfInt_Controllers_Controller_AffinityNames) GetEntityData() *types.CommonEntityData {
    affinityNames.EntityData.YFilter = affinityNames.YFilter
    affinityNames.EntityData.YangName = "affinity-names"
    affinityNames.EntityData.BundleName = "cisco_ios_xr"
    affinityNames.EntityData.ParentYangName = "controller"
    affinityNames.EntityData.SegmentPath = "affinity-names"
    affinityNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityNames.EntityData.Children = types.NewOrderedMap()
    affinityNames.EntityData.Leafs = types.NewOrderedMap()
    affinityNames.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", affinityNames.AffinityName})

    affinityNames.EntityData.YListKeys = []string {}

    return &(affinityNames.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr
// ospfip addr
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Area ID if in IP address format. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // GMPLS-NNI controllers.
    Controllers MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers
}

func (ospfipAddr *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr) GetEntityData() *types.CommonEntityData {
    ospfipAddr.EntityData.YFilter = ospfipAddr.YFilter
    ospfipAddr.EntityData.YangName = "ospfip-addr"
    ospfipAddr.EntityData.BundleName = "cisco_ios_xr"
    ospfipAddr.EntityData.ParentYangName = "topology-instance"
    ospfipAddr.EntityData.SegmentPath = "ospfip-addr" + types.AddKeyToken(ospfipAddr.Address, "address")
    ospfipAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfipAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfipAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfipAddr.EntityData.Children = types.NewOrderedMap()
    ospfipAddr.EntityData.Children.Append("controllers", types.YChild{"Controllers", &ospfipAddr.Controllers})
    ospfipAddr.EntityData.Leafs = types.NewOrderedMap()
    ospfipAddr.EntityData.Leafs.Append("address", types.YLeaf{"Address", ospfipAddr.Address})

    ospfipAddr.EntityData.YListKeys = []string {"Address"}

    return &(ospfipAddr.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers
// GMPLS-NNI controllers
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure a GMPLS NNI controller. The type is slice of
    // MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller.
    Controller []*MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller
}

func (controllers *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers) GetEntityData() *types.CommonEntityData {
    controllers.EntityData.YFilter = controllers.YFilter
    controllers.EntityData.YangName = "controllers"
    controllers.EntityData.BundleName = "cisco_ios_xr"
    controllers.EntityData.ParentYangName = "ospfip-addr"
    controllers.EntityData.SegmentPath = "controllers"
    controllers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controllers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controllers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controllers.EntityData.Children = types.NewOrderedMap()
    controllers.EntityData.Children.Append("controller", types.YChild{"Controller", nil})
    for i := range controllers.Controller {
        controllers.EntityData.Children.Append(types.GetSegmentPath(controllers.Controller[i]), types.YChild{"Controller", controllers.Controller[i]})
    }
    controllers.EntityData.Leafs = types.NewOrderedMap()

    controllers.EntityData.YListKeys = []string {}

    return &(controllers.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller
// Configure a GMPLS NNI controller
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    ControllerName interface{}

    // Set administrative weight for the interface. The type is interface{} with
    // range: 0..65535.
    AdminWeight interface{}

    // Enable GMPLS-NNI on the link. The type is interface{}.
    Enable interface{}

    // Set link delay for the interface. The type is interface{} with range:
    // 1..16777215.
    Delay interface{}

    // Set tandem connection monitoring for the interface.
    TtiMode MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_TtiMode

    // Set the interface affinity names.
    AffinityNames MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_AffinityNames
}

func (controller *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "controllers"
    controller.EntityData.SegmentPath = "controller" + types.AddKeyToken(controller.ControllerName, "controller-name")
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("tti-mode", types.YChild{"TtiMode", &controller.TtiMode})
    controller.EntityData.Children.Append("affinity-names", types.YChild{"AffinityNames", &controller.AffinityNames})
    controller.EntityData.Leafs = types.NewOrderedMap()
    controller.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", controller.ControllerName})
    controller.EntityData.Leafs.Append("admin-weight", types.YLeaf{"AdminWeight", controller.AdminWeight})
    controller.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", controller.Enable})
    controller.EntityData.Leafs.Append("delay", types.YLeaf{"Delay", controller.Delay})

    controller.EntityData.YListKeys = []string {"ControllerName"}

    return &(controller.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_TtiMode
// Set tandem connection monitoring for the
// interface
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_TtiMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of Trail Trace Identifier. The type is GmplsttiMode.
    TtiModeType interface{}

    // Tandem Connection Monitoring ID for the interface. The type is interface{}
    // with range: 1..6.
    Tcmid interface{}
}

func (ttiMode *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_TtiMode) GetEntityData() *types.CommonEntityData {
    ttiMode.EntityData.YFilter = ttiMode.YFilter
    ttiMode.EntityData.YangName = "tti-mode"
    ttiMode.EntityData.BundleName = "cisco_ios_xr"
    ttiMode.EntityData.ParentYangName = "controller"
    ttiMode.EntityData.SegmentPath = "tti-mode"
    ttiMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ttiMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ttiMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ttiMode.EntityData.Children = types.NewOrderedMap()
    ttiMode.EntityData.Leafs = types.NewOrderedMap()
    ttiMode.EntityData.Leafs.Append("tti-mode-type", types.YLeaf{"TtiModeType", ttiMode.TtiModeType})
    ttiMode.EntityData.Leafs.Append("tcmid", types.YLeaf{"Tcmid", ttiMode.Tcmid})

    ttiMode.EntityData.YListKeys = []string {}

    return &(ttiMode.EntityData)
}

// MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_AffinityNames
// Set the interface affinity names
type MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_AffinityNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of Attribute Names. The type is slice of string.
    AffinityName []interface{}
}

func (affinityNames *MplsTe_GmplsNni_TopologyInstances_TopologyInstance_OspfipAddr_Controllers_Controller_AffinityNames) GetEntityData() *types.CommonEntityData {
    affinityNames.EntityData.YFilter = affinityNames.YFilter
    affinityNames.EntityData.YangName = "affinity-names"
    affinityNames.EntityData.BundleName = "cisco_ios_xr"
    affinityNames.EntityData.ParentYangName = "controller"
    affinityNames.EntityData.SegmentPath = "affinity-names"
    affinityNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityNames.EntityData.Children = types.NewOrderedMap()
    affinityNames.EntityData.Leafs = types.NewOrderedMap()
    affinityNames.EntityData.Leafs.Append("affinity-name", types.YLeaf{"AffinityName", affinityNames.AffinityName})

    affinityNames.EntityData.YListKeys = []string {}

    return &(affinityNames.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads
// GMPLS-NNI tunnel-head table
type MplsTe_GmplsNni_TunnelHeads struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The configuration for a GMPLS NNI tunnel head-end. The type is slice of
    // MplsTe_GmplsNni_TunnelHeads_TunnelHead.
    TunnelHead []*MplsTe_GmplsNni_TunnelHeads_TunnelHead
}

func (tunnelHeads *MplsTe_GmplsNni_TunnelHeads) GetEntityData() *types.CommonEntityData {
    tunnelHeads.EntityData.YFilter = tunnelHeads.YFilter
    tunnelHeads.EntityData.YangName = "tunnel-heads"
    tunnelHeads.EntityData.BundleName = "cisco_ios_xr"
    tunnelHeads.EntityData.ParentYangName = "gmpls-nni"
    tunnelHeads.EntityData.SegmentPath = "tunnel-heads"
    tunnelHeads.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelHeads.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelHeads.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelHeads.EntityData.Children = types.NewOrderedMap()
    tunnelHeads.EntityData.Children.Append("tunnel-head", types.YChild{"TunnelHead", nil})
    for i := range tunnelHeads.TunnelHead {
        tunnelHeads.EntityData.Children.Append(types.GetSegmentPath(tunnelHeads.TunnelHead[i]), types.YChild{"TunnelHead", tunnelHeads.TunnelHead[i]})
    }
    tunnelHeads.EntityData.Leafs = types.NewOrderedMap()

    tunnelHeads.EntityData.YListKeys = []string {}

    return &(tunnelHeads.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead
// The configuration for a GMPLS NNI tunnel
// head-end
type MplsTe_GmplsNni_TunnelHeads_TunnelHead struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Tunnel ID. The type is interface{} with range:
    // 0..65535.
    TunnelId interface{}

    // The existence of this configuration indicates the a new GMPLS NNI tunnel
    // has been enabled. The type is interface{}.
    Enable interface{}

    // The existence of this configuration indicates the restore LSP of tunnel is
    // shutdown. The type is interface{}.
    RestoreLspShutdown interface{}

    // The existence of this configuration indicates the current/working LSP of
    // tunnel is shutdown. The type is interface{}.
    CurrentLspShutdown interface{}

    // Path selection configuration for this specific tunnel. The type is
    // MplsTePathSelectionMetric.
    PathSelectionMetric interface{}

    // The existence of this configuration indicates the Payload type have been
    // set for the tunnel. The type is OtnPayload.
    Payload interface{}

    // The existence of this configuration indicates the standby/protect LSP of
    // tunnel is shutdown. The type is interface{}.
    StandbyLspShutdown interface{}

    // The existence of this configuration indicates the tunnel is shutdown. The
    // type is interface{}.
    Shutdown interface{}

    // The name of the path-protection profile to be included in signalling
    // messages. The type is string with length: 1..64.
    PathProtectionAttributeSetProfile interface{}

    // Record the route used by the tunnel. The type is interface{}.
    RecordRoute interface{}

    // The name of the tunnel to be included in signalling messages. The type is
    // string with length: 1..254.
    SignalledName interface{}

    // The existence of this configuration indicates the signalled bandwidth has
    // been set for the tunnel.
    SignalledBandwidth MplsTe_GmplsNni_TunnelHeads_TunnelHead_SignalledBandwidth

    // The existence of this configuration indicates the destination has been set
    // for the tunnel.
    Destination MplsTe_GmplsNni_TunnelHeads_TunnelHead_Destination

    // The configuration for a GMPLS NNI tunnel protection switch.
    ProtectionSwitching MplsTe_GmplsNni_TunnelHeads_TunnelHead_ProtectionSwitching

    // Tunnel event logging.
    Logging MplsTe_GmplsNni_TunnelHeads_TunnelHead_Logging

    // GMPLS NNI path options.
    PathOptions MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions

    // The existence of this configuration indicates the static UNI endpoints have
    // been set for the tunnel.
    StaticUni MplsTe_GmplsNni_TunnelHeads_TunnelHead_StaticUni
}

func (tunnelHead *MplsTe_GmplsNni_TunnelHeads_TunnelHead) GetEntityData() *types.CommonEntityData {
    tunnelHead.EntityData.YFilter = tunnelHead.YFilter
    tunnelHead.EntityData.YangName = "tunnel-head"
    tunnelHead.EntityData.BundleName = "cisco_ios_xr"
    tunnelHead.EntityData.ParentYangName = "tunnel-heads"
    tunnelHead.EntityData.SegmentPath = "tunnel-head" + types.AddKeyToken(tunnelHead.TunnelId, "tunnel-id")
    tunnelHead.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelHead.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelHead.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelHead.EntityData.Children = types.NewOrderedMap()
    tunnelHead.EntityData.Children.Append("signalled-bandwidth", types.YChild{"SignalledBandwidth", &tunnelHead.SignalledBandwidth})
    tunnelHead.EntityData.Children.Append("destination", types.YChild{"Destination", &tunnelHead.Destination})
    tunnelHead.EntityData.Children.Append("protection-switching", types.YChild{"ProtectionSwitching", &tunnelHead.ProtectionSwitching})
    tunnelHead.EntityData.Children.Append("logging", types.YChild{"Logging", &tunnelHead.Logging})
    tunnelHead.EntityData.Children.Append("path-options", types.YChild{"PathOptions", &tunnelHead.PathOptions})
    tunnelHead.EntityData.Children.Append("static-uni", types.YChild{"StaticUni", &tunnelHead.StaticUni})
    tunnelHead.EntityData.Leafs = types.NewOrderedMap()
    tunnelHead.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", tunnelHead.TunnelId})
    tunnelHead.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", tunnelHead.Enable})
    tunnelHead.EntityData.Leafs.Append("restore-lsp-shutdown", types.YLeaf{"RestoreLspShutdown", tunnelHead.RestoreLspShutdown})
    tunnelHead.EntityData.Leafs.Append("current-lsp-shutdown", types.YLeaf{"CurrentLspShutdown", tunnelHead.CurrentLspShutdown})
    tunnelHead.EntityData.Leafs.Append("path-selection-metric", types.YLeaf{"PathSelectionMetric", tunnelHead.PathSelectionMetric})
    tunnelHead.EntityData.Leafs.Append("payload", types.YLeaf{"Payload", tunnelHead.Payload})
    tunnelHead.EntityData.Leafs.Append("standby-lsp-shutdown", types.YLeaf{"StandbyLspShutdown", tunnelHead.StandbyLspShutdown})
    tunnelHead.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", tunnelHead.Shutdown})
    tunnelHead.EntityData.Leafs.Append("path-protection-attribute-set-profile", types.YLeaf{"PathProtectionAttributeSetProfile", tunnelHead.PathProtectionAttributeSetProfile})
    tunnelHead.EntityData.Leafs.Append("record-route", types.YLeaf{"RecordRoute", tunnelHead.RecordRoute})
    tunnelHead.EntityData.Leafs.Append("signalled-name", types.YLeaf{"SignalledName", tunnelHead.SignalledName})

    tunnelHead.EntityData.YListKeys = []string {"TunnelId"}

    return &(tunnelHead.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_SignalledBandwidth
// The existence of this configuration indicates
// the signalled bandwidth has been set for the
// tunnel
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_SignalledBandwidth struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The g.709 signal type requested. The type is OtnSignaledBandwidth.
    SignalledBandwidthType interface{}

    // Bitrate value in Kbps for ODUflex framing type. The type is interface{}
    // with range: 0..4294967295. Units are kbit/s.
    Bitrate interface{}

    // Framing type in case of ODUflex signal type. The type is
    // OtnSignaledBandwidthFlexFraming.
    OdUflexFramingType interface{}
}

func (signalledBandwidth *MplsTe_GmplsNni_TunnelHeads_TunnelHead_SignalledBandwidth) GetEntityData() *types.CommonEntityData {
    signalledBandwidth.EntityData.YFilter = signalledBandwidth.YFilter
    signalledBandwidth.EntityData.YangName = "signalled-bandwidth"
    signalledBandwidth.EntityData.BundleName = "cisco_ios_xr"
    signalledBandwidth.EntityData.ParentYangName = "tunnel-head"
    signalledBandwidth.EntityData.SegmentPath = "signalled-bandwidth"
    signalledBandwidth.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    signalledBandwidth.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    signalledBandwidth.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    signalledBandwidth.EntityData.Children = types.NewOrderedMap()
    signalledBandwidth.EntityData.Leafs = types.NewOrderedMap()
    signalledBandwidth.EntityData.Leafs.Append("signalled-bandwidth-type", types.YLeaf{"SignalledBandwidthType", signalledBandwidth.SignalledBandwidthType})
    signalledBandwidth.EntityData.Leafs.Append("bitrate", types.YLeaf{"Bitrate", signalledBandwidth.Bitrate})
    signalledBandwidth.EntityData.Leafs.Append("od-uflex-framing-type", types.YLeaf{"OdUflexFramingType", signalledBandwidth.OdUflexFramingType})

    signalledBandwidth.EntityData.YListKeys = []string {}

    return &(signalledBandwidth.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_Destination
// The existence of this configuration indicates
// the destination has been set for the tunnel
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV4 tunnel destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Destination interface{}

    // Destination type whether it is unicast or unnumbered. The type is
    // OtnDestination.
    DestinationType interface{}

    // Interface index of port. The type is interface{} with range: 0..4294967295.
    InterfaceIfIndex interface{}
}

func (destination *MplsTe_GmplsNni_TunnelHeads_TunnelHead_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "tunnel-head"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("destination", types.YLeaf{"Destination", destination.Destination})
    destination.EntityData.Leafs.Append("destination-type", types.YLeaf{"DestinationType", destination.DestinationType})
    destination.EntityData.Leafs.Append("interface-if-index", types.YLeaf{"InterfaceIfIndex", destination.InterfaceIfIndex})

    destination.EntityData.YListKeys = []string {}

    return &(destination.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_ProtectionSwitching
// The configuration for a GMPLS NNI tunnel
// protection switch
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_ProtectionSwitching struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The configuration is used to prevent switch over for a particular path type
    // in tunnel. The type is OtnProtectionSwitchLockout.
    Lockout interface{}
}

func (protectionSwitching *MplsTe_GmplsNni_TunnelHeads_TunnelHead_ProtectionSwitching) GetEntityData() *types.CommonEntityData {
    protectionSwitching.EntityData.YFilter = protectionSwitching.YFilter
    protectionSwitching.EntityData.YangName = "protection-switching"
    protectionSwitching.EntityData.BundleName = "cisco_ios_xr"
    protectionSwitching.EntityData.ParentYangName = "tunnel-head"
    protectionSwitching.EntityData.SegmentPath = "protection-switching"
    protectionSwitching.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    protectionSwitching.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    protectionSwitching.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    protectionSwitching.EntityData.Children = types.NewOrderedMap()
    protectionSwitching.EntityData.Leafs = types.NewOrderedMap()
    protectionSwitching.EntityData.Leafs.Append("lockout", types.YLeaf{"Lockout", protectionSwitching.Lockout})

    protectionSwitching.EntityData.YListKeys = []string {}

    return &(protectionSwitching.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_Logging
// Tunnel event logging
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log all tunnel messages for changes in Active LSP. The type is interface{}.
    ActiveLspMessage interface{}

    // Log all messages for changes in state of Homepath of Working LSP. The type
    // is interface{}.
    HomepathStateMessage interface{}

    // Log all tunnel sub-LSP state messages. The type is interface{}.
    SignallingStateMessage interface{}

    // Log all tunnel messages for changes in path-change. The type is
    // interface{}.
    PathChangeMessage interface{}

    // Log all tunnel messages for static cross-connect messages. The type is
    // interface{}.
    StaticCrossConnectMessage interface{}

    // Log all tunnel messages for changes in tunnel-state. The type is
    // interface{}.
    TunnelStateMessage interface{}

    // Log tunnel messages for insufficient bandwidth. The type is interface{}.
    InsufficientBwMessage interface{}
}

func (logging *MplsTe_GmplsNni_TunnelHeads_TunnelHead_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "tunnel-head"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("active-lsp-message", types.YLeaf{"ActiveLspMessage", logging.ActiveLspMessage})
    logging.EntityData.Leafs.Append("homepath-state-message", types.YLeaf{"HomepathStateMessage", logging.HomepathStateMessage})
    logging.EntityData.Leafs.Append("signalling-state-message", types.YLeaf{"SignallingStateMessage", logging.SignallingStateMessage})
    logging.EntityData.Leafs.Append("path-change-message", types.YLeaf{"PathChangeMessage", logging.PathChangeMessage})
    logging.EntityData.Leafs.Append("static-cross-connect-message", types.YLeaf{"StaticCrossConnectMessage", logging.StaticCrossConnectMessage})
    logging.EntityData.Leafs.Append("tunnel-state-message", types.YLeaf{"TunnelStateMessage", logging.TunnelStateMessage})
    logging.EntityData.Leafs.Append("insufficient-bw-message", types.YLeaf{"InsufficientBwMessage", logging.InsufficientBwMessage})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions
// GMPLS NNI path options
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The existence of this configuration indicates the path options have been
    // set for the tunnel. The type is slice of
    // MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions_PathOption.
    PathOption []*MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions_PathOption
}

func (pathOptions *MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions) GetEntityData() *types.CommonEntityData {
    pathOptions.EntityData.YFilter = pathOptions.YFilter
    pathOptions.EntityData.YangName = "path-options"
    pathOptions.EntityData.BundleName = "cisco_ios_xr"
    pathOptions.EntityData.ParentYangName = "tunnel-head"
    pathOptions.EntityData.SegmentPath = "path-options"
    pathOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOptions.EntityData.Children = types.NewOrderedMap()
    pathOptions.EntityData.Children.Append("path-option", types.YChild{"PathOption", nil})
    for i := range pathOptions.PathOption {
        pathOptions.EntityData.Children.Append(types.GetSegmentPath(pathOptions.PathOption[i]), types.YChild{"PathOption", pathOptions.PathOption[i]})
    }
    pathOptions.EntityData.Leafs = types.NewOrderedMap()

    pathOptions.EntityData.YListKeys = []string {}

    return &(pathOptions.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions_PathOption
// The existence of this configuration
// indicates the path options have been set for
// the tunnel
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions_PathOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Preference level for this path option. The type is
    // interface{} with range: 1..1000.
    PreferenceLevel interface{}

    // The type of the path option. The type is MplsTePathOption.
    PathType interface{}

    // The ID of the IP explicit path associated with this option. The type is
    // interface{} with range: 1..65535.
    PathId interface{}

    // The name of the IP explicit path associated with this option. The type is
    // string.
    PathName interface{}

    // Preference level of the protecting explicit path. . The type is interface{}
    // with range: 1..1001.
    ProtectedByPreferenceLevel interface{}

    // Preference level of the restore path. . The type is interface{} with range:
    // 1..1000.
    RestoreByPreferenceLevel interface{}

    // The route-exclusion type. The type is interface{}. This attribute is
    // mandatory.
    XroType interface{}

    // The name of the XRO attribute set to be used for this path-option. The type
    // is string with length: 1..64.
    XroAttributeSetName interface{}

    // The name of the affinity attribute set to be used for this path-option. The
    // type is string with length: 1..64.
    AffinityAttributeSetName interface{}

    // Lockdown properties. The type is MplsTePathOptionProperty.
    Lockdown interface{}
}

func (pathOption *MplsTe_GmplsNni_TunnelHeads_TunnelHead_PathOptions_PathOption) GetEntityData() *types.CommonEntityData {
    pathOption.EntityData.YFilter = pathOption.YFilter
    pathOption.EntityData.YangName = "path-option"
    pathOption.EntityData.BundleName = "cisco_ios_xr"
    pathOption.EntityData.ParentYangName = "path-options"
    pathOption.EntityData.SegmentPath = "path-option" + types.AddKeyToken(pathOption.PreferenceLevel, "preference-level")
    pathOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathOption.EntityData.Children = types.NewOrderedMap()
    pathOption.EntityData.Leafs = types.NewOrderedMap()
    pathOption.EntityData.Leafs.Append("preference-level", types.YLeaf{"PreferenceLevel", pathOption.PreferenceLevel})
    pathOption.EntityData.Leafs.Append("path-type", types.YLeaf{"PathType", pathOption.PathType})
    pathOption.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", pathOption.PathId})
    pathOption.EntityData.Leafs.Append("path-name", types.YLeaf{"PathName", pathOption.PathName})
    pathOption.EntityData.Leafs.Append("protected-by-preference-level", types.YLeaf{"ProtectedByPreferenceLevel", pathOption.ProtectedByPreferenceLevel})
    pathOption.EntityData.Leafs.Append("restore-by-preference-level", types.YLeaf{"RestoreByPreferenceLevel", pathOption.RestoreByPreferenceLevel})
    pathOption.EntityData.Leafs.Append("xro-type", types.YLeaf{"XroType", pathOption.XroType})
    pathOption.EntityData.Leafs.Append("xro-attribute-set-name", types.YLeaf{"XroAttributeSetName", pathOption.XroAttributeSetName})
    pathOption.EntityData.Leafs.Append("affinity-attribute-set-name", types.YLeaf{"AffinityAttributeSetName", pathOption.AffinityAttributeSetName})
    pathOption.EntityData.Leafs.Append("lockdown", types.YLeaf{"Lockdown", pathOption.Lockdown})

    pathOption.EntityData.YListKeys = []string {"PreferenceLevel"}

    return &(pathOption.EntityData)
}

// MplsTe_GmplsNni_TunnelHeads_TunnelHead_StaticUni
// The existence of this configuration indicates
// the static UNI endpoints have been set for
// the tunnel
type MplsTe_GmplsNni_TunnelHeads_TunnelHead_StaticUni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of  ingress controller. The type is string with length: 1..255.
    IngressControllerName interface{}

    // Interface index of Egress controller. The type is interface{} with range:
    // 0..4294967295.
    EgressControllerIfIndex interface{}

    // Ingress type whether it is xconnect or terminated. The type is
    // OtnStaticUni.
    IngressType interface{}

    // Egress type whether it is xconnect or terminated. The type is OtnStaticUni.
    EgressType interface{}
}

func (staticUni *MplsTe_GmplsNni_TunnelHeads_TunnelHead_StaticUni) GetEntityData() *types.CommonEntityData {
    staticUni.EntityData.YFilter = staticUni.YFilter
    staticUni.EntityData.YangName = "static-uni"
    staticUni.EntityData.BundleName = "cisco_ios_xr"
    staticUni.EntityData.ParentYangName = "tunnel-head"
    staticUni.EntityData.SegmentPath = "static-uni"
    staticUni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticUni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticUni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticUni.EntityData.Children = types.NewOrderedMap()
    staticUni.EntityData.Leafs = types.NewOrderedMap()
    staticUni.EntityData.Leafs.Append("ingress-controller-name", types.YLeaf{"IngressControllerName", staticUni.IngressControllerName})
    staticUni.EntityData.Leafs.Append("egress-controller-if-index", types.YLeaf{"EgressControllerIfIndex", staticUni.EgressControllerIfIndex})
    staticUni.EntityData.Leafs.Append("ingress-type", types.YLeaf{"IngressType", staticUni.IngressType})
    staticUni.EntityData.Leafs.Append("egress-type", types.YLeaf{"EgressType", staticUni.EgressType})

    staticUni.EntityData.YListKeys = []string {}

    return &(staticUni.EntityData)
}

// MplsTe_Lcac
// LCAC specific MPLS global configuration
type MplsTe_Lcac struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Bandwidth hold timer value (seconds). The type is interface{} with range:
    // 1..300. Units are second.
    BandwidthHoldTimer interface{}

    // Bundle capacity preemption timer value (seconds). The type is interface{}
    // with range: 0..300. Units are second.
    DelayPreemptBundleCapacityTimer interface{}

    // Periodic flooding value (seconds). The type is interface{} with range:
    // 0..3600. Units are second.
    PeriodicFloodingTimer interface{}

    // BFD configuration.
    Bfd MplsTe_Lcac_Bfd

    // Configure flooding threshold as percentage of total link bandwidth.
    FloodingThreshold MplsTe_Lcac_FloodingThreshold
}

func (lcac *MplsTe_Lcac) GetEntityData() *types.CommonEntityData {
    lcac.EntityData.YFilter = lcac.YFilter
    lcac.EntityData.YangName = "lcac"
    lcac.EntityData.BundleName = "cisco_ios_xr"
    lcac.EntityData.ParentYangName = "mpls-te"
    lcac.EntityData.SegmentPath = "lcac"
    lcac.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcac.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcac.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcac.EntityData.Children = types.NewOrderedMap()
    lcac.EntityData.Children.Append("bfd", types.YChild{"Bfd", &lcac.Bfd})
    lcac.EntityData.Children.Append("flooding-threshold", types.YChild{"FloodingThreshold", &lcac.FloodingThreshold})
    lcac.EntityData.Leafs = types.NewOrderedMap()
    lcac.EntityData.Leafs.Append("bandwidth-hold-timer", types.YLeaf{"BandwidthHoldTimer", lcac.BandwidthHoldTimer})
    lcac.EntityData.Leafs.Append("delay-preempt-bundle-capacity-timer", types.YLeaf{"DelayPreemptBundleCapacityTimer", lcac.DelayPreemptBundleCapacityTimer})
    lcac.EntityData.Leafs.Append("periodic-flooding-timer", types.YLeaf{"PeriodicFloodingTimer", lcac.PeriodicFloodingTimer})

    lcac.EntityData.YListKeys = []string {}

    return &(lcac.EntityData)
}

// MplsTe_Lcac_Bfd
// BFD configuration
type MplsTe_Lcac_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval for BFD sessions created by TE. The type is interface{} with
    // range: 15..200. Units are millisecond.
    Interval interface{}

    // Detection multiplier for BFD sessions created by TE. The type is
    // interface{} with range: 2..10.
    DetectionMultiplier interface{}
}

func (bfd *MplsTe_Lcac_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "lcac"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// MplsTe_Lcac_FloodingThreshold
// Configure flooding threshold as percentage of
// total link bandwidth.
type MplsTe_Lcac_FloodingThreshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Upward flooding Threshold in percentages of total bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage.
    UpStream interface{}

    // Downward flooding Threshold in percentages of total bandwidth. The type is
    // interface{} with range: 0..100. Units are percentage.
    DownStream interface{}
}

func (floodingThreshold *MplsTe_Lcac_FloodingThreshold) GetEntityData() *types.CommonEntityData {
    floodingThreshold.EntityData.YFilter = floodingThreshold.YFilter
    floodingThreshold.EntityData.YangName = "flooding-threshold"
    floodingThreshold.EntityData.BundleName = "cisco_ios_xr"
    floodingThreshold.EntityData.ParentYangName = "lcac"
    floodingThreshold.EntityData.SegmentPath = "flooding-threshold"
    floodingThreshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    floodingThreshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    floodingThreshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    floodingThreshold.EntityData.Children = types.NewOrderedMap()
    floodingThreshold.EntityData.Leafs = types.NewOrderedMap()
    floodingThreshold.EntityData.Leafs.Append("up-stream", types.YLeaf{"UpStream", floodingThreshold.UpStream})
    floodingThreshold.EntityData.Leafs.Append("down-stream", types.YLeaf{"DownStream", floodingThreshold.DownStream})

    floodingThreshold.EntityData.YListKeys = []string {}

    return &(floodingThreshold.EntityData)
}

