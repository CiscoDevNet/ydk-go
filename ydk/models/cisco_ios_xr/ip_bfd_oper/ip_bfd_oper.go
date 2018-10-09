// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-bfd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   bfd: Bidirectional Forwarding Detection(BFD) operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_bfd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_bfd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-bfd-oper bfd}", reflect.TypeOf(Bfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-bfd-oper:bfd", reflect.TypeOf(Bfd{}))
}

// BfdSession represents BFD session type
type BfdSession string

const (
    // Session type is undefined
    BfdSession_undefined BfdSession = "undefined"

    // Session type is bundle member
    BfdSession_bundle_member BfdSession = "bundle-member"

    // Session type is bundle interface
    BfdSession_bundle_interface BfdSession = "bundle-interface"

    // Session type is state inheriting
    BfdSession_state_inheriting BfdSession = "state-inheriting"

    // Session type is bundle vlan
    BfdSession_bundle_vlan BfdSession = "bundle-vlan"

    // Session type is MPLS-TP
    BfdSession_mpls_tp BfdSession = "mpls-tp"

    // Session type is GRE tunnel
    BfdSession_gre BfdSession = "gre"

    // Session type is PW-HE
    BfdSession_pseudowire_headend BfdSession = "pseudowire-headend"

    // Session type is IP single hop
    BfdSession_ip_single_hop BfdSession = "ip-single-hop"
)

// BfdAfId represents Bfd af id
type BfdAfId string

const (
    // No Address
    BfdAfId_bfd_af_id_none BfdAfId = "bfd-af-id-none"

    // IPv4 AFI
    BfdAfId_bfd_af_id_ipv4 BfdAfId = "bfd-af-id-ipv4"

    // IPv6 AFI
    BfdAfId_bfd_af_id_ipv6 BfdAfId = "bfd-af-id-ipv6"
)

// BfdMpDownloadState represents Bfd mp download state
type BfdMpDownloadState string

const (
    // bfd mp download none
    BfdMpDownloadState_bfd_mp_download_none BfdMpDownloadState = "bfd-mp-download-none"

    // bfd mp download no lc
    BfdMpDownloadState_bfd_mp_download_no_lc BfdMpDownloadState = "bfd-mp-download-no-lc"

    // bfd mp download downloaded
    BfdMpDownloadState_bfd_mp_download_downloaded BfdMpDownloadState = "bfd-mp-download-downloaded"

    // bfd mp download ack
    BfdMpDownloadState_bfd_mp_download_ack BfdMpDownloadState = "bfd-mp-download-ack"

    // bfd mp download nack
    BfdMpDownloadState_bfd_mp_download_nack BfdMpDownloadState = "bfd-mp-download-nack"

    // bfd mp download delete
    BfdMpDownloadState_bfd_mp_download_delete BfdMpDownloadState = "bfd-mp-download-delete"
)

// BfdMgmtSessionState represents BFD session states
type BfdMgmtSessionState string

const (
    // Session in Administratively Shutdown State
    BfdMgmtSessionState_bfd_mgmt_session_state_admin_down BfdMgmtSessionState = "bfd-mgmt-session-state-admin-down"

    // Session in Down State
    BfdMgmtSessionState_bfd_mgmt_session_state_down BfdMgmtSessionState = "bfd-mgmt-session-state-down"

    // Session in Initializing State
    BfdMgmtSessionState_bfd_mgmt_session_state_init BfdMgmtSessionState = "bfd-mgmt-session-state-init"

    // Session in Up State
    BfdMgmtSessionState_bfd_mgmt_session_state_up BfdMgmtSessionState = "bfd-mgmt-session-state-up"

    // Session in Failing State
    BfdMgmtSessionState_bfd_mgmt_session_state_failing BfdMgmtSessionState = "bfd-mgmt-session-state-failing"

    // Session in Unknown State
    BfdMgmtSessionState_bfd_mgmt_session_state_unknown BfdMgmtSessionState = "bfd-mgmt-session-state-unknown"
)

// BfdMgmtSessionDiag represents BFD session diagnostic
type BfdMgmtSessionDiag string

const (
    // No diagnostic
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_none BfdMgmtSessionDiag = "bfd-mgmt-session-diag-none"

    // Control Detection Time Expired
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_control_detect_expired BfdMgmtSessionDiag = "bfd-mgmt-session-diag-control-detect-expired"

    // Echo Function Failed
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_echo_function_failed BfdMgmtSessionDiag = "bfd-mgmt-session-diag-echo-function-failed"

    // Neighbor Signaled Session Down
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_nb_or_signaled_down BfdMgmtSessionDiag = "bfd-mgmt-session-diag-nb-or-signaled-down"

    // Forwarding Plane Reset
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_fwding_plane_reset BfdMgmtSessionDiag = "bfd-mgmt-session-diag-fwding-plane-reset"

    // Path Down
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_path_down BfdMgmtSessionDiag = "bfd-mgmt-session-diag-path-down"

    // Concatenated Path Down
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_conc_path_down BfdMgmtSessionDiag = "bfd-mgmt-session-diag-conc-path-down"

    // Administratively Down
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_admin_down BfdMgmtSessionDiag = "bfd-mgmt-session-diag-admin-down"

    // Reverse Concatenated Path Down
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_rev_conc_path_down BfdMgmtSessionDiag = "bfd-mgmt-session-diag-rev-conc-path-down"

    // Unknown diagnostic state
    BfdMgmtSessionDiag_bfd_mgmt_session_diag_num BfdMgmtSessionDiag = "bfd-mgmt-session-diag-num"
)

// BfdMgmtPktDisplay represents BFD session type
type BfdMgmtPktDisplay string

const (
    // None
    BfdMgmtPktDisplay_bfd_mgmt_pkt_display_type_none BfdMgmtPktDisplay = "bfd-mgmt-pkt-display-type-none"

    // Display Type Bundle RTR Member
    BfdMgmtPktDisplay_bfd_mgmt_pkt_display_type_bob_mbr BfdMgmtPktDisplay = "bfd-mgmt-pkt-display-type-bob-mbr"

    // Display Type Enum Max Value
    BfdMgmtPktDisplay_bfd_mgmt_pkt_display_type_max BfdMgmtPktDisplay = "bfd-mgmt-pkt-display-type-max"
)

// Bfd
// Bidirectional Forwarding Detection(BFD)
// operational data
type Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of brief information about all Label BFD sessions in the System.
    LabelSessionBriefs Bfd_LabelSessionBriefs

    // Summary information of IPv4 BFD over MPLS-TE Tail.
    Ipv4bfDoMplsteTailSummary Bfd_Ipv4bfDoMplsteTailSummary

    // IPv6 single hop Counters.
    Ipv6SingleHopCounters Bfd_Ipv6SingleHopCounters

    // IPv4 Counters.
    Counters Bfd_Counters

    // Table of detailed information about BFD clients.
    ClientDetails Bfd_ClientDetails

    // Summary information of BFD IPv4 singlehop sessions.
    Ipv4SingleHopSummary Bfd_Ipv4SingleHopSummary

    // Summary information of BFD IPv6 singlehop sessions.
    Ipv6SingleHopSummary Bfd_Ipv6SingleHopSummary

    // Label multipath.
    LabelMultiPaths Bfd_LabelMultiPaths

    // Table of detailed information about all IPv4 multihop BFD sessions in the
    // System .
    Ipv4MultiHopSessionDetails Bfd_Ipv4MultiHopSessionDetails

    // Table of detailed information about all IPv4 singlehop BFD sessions in the
    // System .
    Ipv4SingleHopSessionDetails Bfd_Ipv4SingleHopSessionDetails

    // Table of brief information about all IPv4 multihop BFD sessions in the
    // System.
    Ipv4MultiHopSessionBriefs Bfd_Ipv4MultiHopSessionBriefs

    // Generic summary information about BFD location.
    GenericSummaries Bfd_GenericSummaries

    // IPv6 single hop multipath.
    Ipv6SingleHopMultiPaths Bfd_Ipv6SingleHopMultiPaths

    // Table of summary information about BFD IPv4 singlehop sessions per
    // location.
    Ipv4SingleHopNodeLocationSummaries Bfd_Ipv4SingleHopNodeLocationSummaries

    // Summary information of Label BFD.
    LabelSummary Bfd_LabelSummary

    // Table of brief information about all IPv4 BFD over MPLS-TE Head sessions in
    // the System.
    Ipv4bfDoMplsteHeadSessionBriefs Bfd_Ipv4bfDoMplsteHeadSessionBriefs

    // Table of detailed information about all IPv4 BFD over MPLS-TE Tail sessions
    // in the System.
    Ipv4bfDoMplsteTailSessionDetails Bfd_Ipv4bfDoMplsteTailSessionDetails

    // Table of summary information about BFD IPv4 multihop sessions per location.
    Ipv4MultiHopNodeLocationSummaries Bfd_Ipv4MultiHopNodeLocationSummaries

    // Table of brief information about all IPv4 BFD over MPLS-TE Tail sessions in
    // the System.
    Ipv4bfDoMplsteTailSessionBriefs Bfd_Ipv4bfDoMplsteTailSessionBriefs

    // Table of summary information about BFD IPv6 multihop sessions per location.
    Ipv6MultiHopNodeLocationSummaries Bfd_Ipv6MultiHopNodeLocationSummaries

    // Summary information of BFD IPv4 multihop sessions.
    Ipv4MultiHopSummary Bfd_Ipv4MultiHopSummary

    // IPv4 single hop Counters.
    Ipv4SingleHopCounters Bfd_Ipv4SingleHopCounters

    // Table of detailed information about all IPv6 multihop BFD sessions in the
    // System .
    Ipv6MultiHopSessionDetails Bfd_Ipv6MultiHopSessionDetails

    // IPv6 multi hop multipath.
    Ipv6MultiHopMultiPaths Bfd_Ipv6MultiHopMultiPaths

    // IPv4 BFD over MPLS-TE Counters.
    Ipv4bfDoMplsteHeadCounters Bfd_Ipv4bfDoMplsteHeadCounters

    // BFD session MIB database.
    SessionMibs Bfd_SessionMibs

    // Summary information of BFD IPv6 multihop sessions.
    Ipv6MultiHopSummary Bfd_Ipv6MultiHopSummary

    // Table of summary about Label BFD sessions for location.
    LabelSummaryNodes Bfd_LabelSummaryNodes

    // Table of brief information about all IPv6 multihop BFD sessions in the
    // System.
    Ipv6MultiHopSessionBriefs Bfd_Ipv6MultiHopSessionBriefs

    // Table of brief information about singlehop IPv4 BFD sessions in the System.
    SessionBriefs Bfd_SessionBriefs

    // Table of summary information about BFD IPv6 singlehop sessions per
    // location.
    Ipv6SingleHopNodeLocationSummaries Bfd_Ipv6SingleHopNodeLocationSummaries

    // Summary information of BFD IPv4 singlehop sessions.
    Summary Bfd_Summary

    // Table of summary about IPv4 TE tail BFD sessions for location.
    Ipv4bfdMplsteTailNodeSummaries Bfd_Ipv4bfdMplsteTailNodeSummaries

    // Table of summary information about IPv4 singlehop BFD sessions for
    // location.
    Ipv4SingleHopLocationSummaries Bfd_Ipv4SingleHopLocationSummaries

    // Table of summary about IPv4 TE head BFD sessions for location.
    Ipv4bfdMplsteHeadSummaryNodes Bfd_Ipv4bfdMplsteHeadSummaryNodes

    // Table of detailed information about all Label BFD sessions in the System .
    LabelSessionDetails Bfd_LabelSessionDetails

    // Table of detailed information about all IPv6 singlehop BFD sessions in the
    // System .
    Ipv6SingleHopSessionDetails Bfd_Ipv6SingleHopSessionDetails

    // IPv4 multiple hop Counters.
    Ipv4MultiHopCounters Bfd_Ipv4MultiHopCounters

    // Table of detailed information about IPv4 singlehop BFD sessions in the
    // System .
    SessionDetails Bfd_SessionDetails

    // IPv4 single hop multipath.
    Ipv4SingleHopMultiPaths Bfd_Ipv4SingleHopMultiPaths

    // Table of brief information about all IPv4 singlehop BFD sessions in the
    // System.
    Ipv4SingleHopSessionBriefs Bfd_Ipv4SingleHopSessionBriefs

    // IPv6 multiple hop Counters.
    Ipv6MultiHopCounters Bfd_Ipv6MultiHopCounters

    // Table of summary information about BFD IPv6 singlehop sessions per
    // location.
    Ipv6SingleHopLocationSummaries Bfd_Ipv6SingleHopLocationSummaries

    // Label Counters.
    LabelCounters Bfd_LabelCounters

    // Table of detailed information about all IPv4 BFD over MPLS-TE Head sessions
    // in the System.
    Ipv4bfDoMplsteHeadSessionDetails Bfd_Ipv4bfDoMplsteHeadSessionDetails

    // Table of brief information about all BFD relations in the System.
    RelationBriefs Bfd_RelationBriefs

    // Table of Brief information about BFD clients.
    ClientBriefs Bfd_ClientBriefs

    // IPv4 BFD over MPLS-TE Head multipath.
    Ipv4bfDoMplsteHeadMultiPaths Bfd_Ipv4bfDoMplsteHeadMultiPaths

    // Table of detail information about all BFD relations in the System.
    RelationDetails Bfd_RelationDetails

    // IPv4 BFD over MPLS-TE Counters.
    Ipv4bfDoMplsteTailCounters Bfd_Ipv4bfDoMplsteTailCounters

    // Table of brief information about all IPv6 singlehop BFD sessions in the
    // System.
    Ipv6SingleHopSessionBriefs Bfd_Ipv6SingleHopSessionBriefs

    // IPv4 BFD over MPLS-TE Tail multipath.
    Ipv4bfDoMplsteTailMultiPaths Bfd_Ipv4bfDoMplsteTailMultiPaths

    // IPv4 multi-hop multipath.
    Ipv4MultiHopMultiPaths Bfd_Ipv4MultiHopMultiPaths

    // Summary information of IPv4 BFD over MPLS-TE Head.
    Ipv4bfDoMplsteHeadSummary Bfd_Ipv4bfDoMplsteHeadSummary
}

func (bfd *Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "Cisco-IOS-XR-ip-bfd-oper"
    bfd.EntityData.SegmentPath = "Cisco-IOS-XR-ip-bfd-oper:bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Children.Append("label-session-briefs", types.YChild{"LabelSessionBriefs", &bfd.LabelSessionBriefs})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-tail-summary", types.YChild{"Ipv4bfDoMplsteTailSummary", &bfd.Ipv4bfDoMplsteTailSummary})
    bfd.EntityData.Children.Append("ipv6-single-hop-counters", types.YChild{"Ipv6SingleHopCounters", &bfd.Ipv6SingleHopCounters})
    bfd.EntityData.Children.Append("counters", types.YChild{"Counters", &bfd.Counters})
    bfd.EntityData.Children.Append("client-details", types.YChild{"ClientDetails", &bfd.ClientDetails})
    bfd.EntityData.Children.Append("ipv4-single-hop-summary", types.YChild{"Ipv4SingleHopSummary", &bfd.Ipv4SingleHopSummary})
    bfd.EntityData.Children.Append("ipv6-single-hop-summary", types.YChild{"Ipv6SingleHopSummary", &bfd.Ipv6SingleHopSummary})
    bfd.EntityData.Children.Append("label-multi-paths", types.YChild{"LabelMultiPaths", &bfd.LabelMultiPaths})
    bfd.EntityData.Children.Append("ipv4-multi-hop-session-details", types.YChild{"Ipv4MultiHopSessionDetails", &bfd.Ipv4MultiHopSessionDetails})
    bfd.EntityData.Children.Append("ipv4-single-hop-session-details", types.YChild{"Ipv4SingleHopSessionDetails", &bfd.Ipv4SingleHopSessionDetails})
    bfd.EntityData.Children.Append("ipv4-multi-hop-session-briefs", types.YChild{"Ipv4MultiHopSessionBriefs", &bfd.Ipv4MultiHopSessionBriefs})
    bfd.EntityData.Children.Append("generic-summaries", types.YChild{"GenericSummaries", &bfd.GenericSummaries})
    bfd.EntityData.Children.Append("ipv6-single-hop-multi-paths", types.YChild{"Ipv6SingleHopMultiPaths", &bfd.Ipv6SingleHopMultiPaths})
    bfd.EntityData.Children.Append("ipv4-single-hop-node-location-summaries", types.YChild{"Ipv4SingleHopNodeLocationSummaries", &bfd.Ipv4SingleHopNodeLocationSummaries})
    bfd.EntityData.Children.Append("label-summary", types.YChild{"LabelSummary", &bfd.LabelSummary})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-head-session-briefs", types.YChild{"Ipv4bfDoMplsteHeadSessionBriefs", &bfd.Ipv4bfDoMplsteHeadSessionBriefs})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-tail-session-details", types.YChild{"Ipv4bfDoMplsteTailSessionDetails", &bfd.Ipv4bfDoMplsteTailSessionDetails})
    bfd.EntityData.Children.Append("ipv4-multi-hop-node-location-summaries", types.YChild{"Ipv4MultiHopNodeLocationSummaries", &bfd.Ipv4MultiHopNodeLocationSummaries})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-tail-session-briefs", types.YChild{"Ipv4bfDoMplsteTailSessionBriefs", &bfd.Ipv4bfDoMplsteTailSessionBriefs})
    bfd.EntityData.Children.Append("ipv6-multi-hop-node-location-summaries", types.YChild{"Ipv6MultiHopNodeLocationSummaries", &bfd.Ipv6MultiHopNodeLocationSummaries})
    bfd.EntityData.Children.Append("ipv4-multi-hop-summary", types.YChild{"Ipv4MultiHopSummary", &bfd.Ipv4MultiHopSummary})
    bfd.EntityData.Children.Append("ipv4-single-hop-counters", types.YChild{"Ipv4SingleHopCounters", &bfd.Ipv4SingleHopCounters})
    bfd.EntityData.Children.Append("ipv6-multi-hop-session-details", types.YChild{"Ipv6MultiHopSessionDetails", &bfd.Ipv6MultiHopSessionDetails})
    bfd.EntityData.Children.Append("ipv6-multi-hop-multi-paths", types.YChild{"Ipv6MultiHopMultiPaths", &bfd.Ipv6MultiHopMultiPaths})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-head-counters", types.YChild{"Ipv4bfDoMplsteHeadCounters", &bfd.Ipv4bfDoMplsteHeadCounters})
    bfd.EntityData.Children.Append("session-mibs", types.YChild{"SessionMibs", &bfd.SessionMibs})
    bfd.EntityData.Children.Append("ipv6-multi-hop-summary", types.YChild{"Ipv6MultiHopSummary", &bfd.Ipv6MultiHopSummary})
    bfd.EntityData.Children.Append("label-summary-nodes", types.YChild{"LabelSummaryNodes", &bfd.LabelSummaryNodes})
    bfd.EntityData.Children.Append("ipv6-multi-hop-session-briefs", types.YChild{"Ipv6MultiHopSessionBriefs", &bfd.Ipv6MultiHopSessionBriefs})
    bfd.EntityData.Children.Append("session-briefs", types.YChild{"SessionBriefs", &bfd.SessionBriefs})
    bfd.EntityData.Children.Append("ipv6-single-hop-node-location-summaries", types.YChild{"Ipv6SingleHopNodeLocationSummaries", &bfd.Ipv6SingleHopNodeLocationSummaries})
    bfd.EntityData.Children.Append("summary", types.YChild{"Summary", &bfd.Summary})
    bfd.EntityData.Children.Append("ipv4bfd-mplste-tail-node-summaries", types.YChild{"Ipv4bfdMplsteTailNodeSummaries", &bfd.Ipv4bfdMplsteTailNodeSummaries})
    bfd.EntityData.Children.Append("ipv4-single-hop-location-summaries", types.YChild{"Ipv4SingleHopLocationSummaries", &bfd.Ipv4SingleHopLocationSummaries})
    bfd.EntityData.Children.Append("ipv4bfd-mplste-head-summary-nodes", types.YChild{"Ipv4bfdMplsteHeadSummaryNodes", &bfd.Ipv4bfdMplsteHeadSummaryNodes})
    bfd.EntityData.Children.Append("label-session-details", types.YChild{"LabelSessionDetails", &bfd.LabelSessionDetails})
    bfd.EntityData.Children.Append("ipv6-single-hop-session-details", types.YChild{"Ipv6SingleHopSessionDetails", &bfd.Ipv6SingleHopSessionDetails})
    bfd.EntityData.Children.Append("ipv4-multi-hop-counters", types.YChild{"Ipv4MultiHopCounters", &bfd.Ipv4MultiHopCounters})
    bfd.EntityData.Children.Append("session-details", types.YChild{"SessionDetails", &bfd.SessionDetails})
    bfd.EntityData.Children.Append("ipv4-single-hop-multi-paths", types.YChild{"Ipv4SingleHopMultiPaths", &bfd.Ipv4SingleHopMultiPaths})
    bfd.EntityData.Children.Append("ipv4-single-hop-session-briefs", types.YChild{"Ipv4SingleHopSessionBriefs", &bfd.Ipv4SingleHopSessionBriefs})
    bfd.EntityData.Children.Append("ipv6-multi-hop-counters", types.YChild{"Ipv6MultiHopCounters", &bfd.Ipv6MultiHopCounters})
    bfd.EntityData.Children.Append("ipv6-single-hop-location-summaries", types.YChild{"Ipv6SingleHopLocationSummaries", &bfd.Ipv6SingleHopLocationSummaries})
    bfd.EntityData.Children.Append("label-counters", types.YChild{"LabelCounters", &bfd.LabelCounters})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-head-session-details", types.YChild{"Ipv4bfDoMplsteHeadSessionDetails", &bfd.Ipv4bfDoMplsteHeadSessionDetails})
    bfd.EntityData.Children.Append("relation-briefs", types.YChild{"RelationBriefs", &bfd.RelationBriefs})
    bfd.EntityData.Children.Append("client-briefs", types.YChild{"ClientBriefs", &bfd.ClientBriefs})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-head-multi-paths", types.YChild{"Ipv4bfDoMplsteHeadMultiPaths", &bfd.Ipv4bfDoMplsteHeadMultiPaths})
    bfd.EntityData.Children.Append("relation-details", types.YChild{"RelationDetails", &bfd.RelationDetails})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-tail-counters", types.YChild{"Ipv4bfDoMplsteTailCounters", &bfd.Ipv4bfDoMplsteTailCounters})
    bfd.EntityData.Children.Append("ipv6-single-hop-session-briefs", types.YChild{"Ipv6SingleHopSessionBriefs", &bfd.Ipv6SingleHopSessionBriefs})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-tail-multi-paths", types.YChild{"Ipv4bfDoMplsteTailMultiPaths", &bfd.Ipv4bfDoMplsteTailMultiPaths})
    bfd.EntityData.Children.Append("ipv4-multi-hop-multi-paths", types.YChild{"Ipv4MultiHopMultiPaths", &bfd.Ipv4MultiHopMultiPaths})
    bfd.EntityData.Children.Append("ipv4bf-do-mplste-head-summary", types.YChild{"Ipv4bfDoMplsteHeadSummary", &bfd.Ipv4bfDoMplsteHeadSummary})
    bfd.EntityData.Leafs = types.NewOrderedMap()

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Bfd_LabelSessionBriefs
// Table of brief information about all Label BFD
// sessions in the System
type Bfd_LabelSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single Label BFD session. The type is slice of
    // Bfd_LabelSessionBriefs_LabelSessionBrief.
    LabelSessionBrief []*Bfd_LabelSessionBriefs_LabelSessionBrief
}

func (labelSessionBriefs *Bfd_LabelSessionBriefs) GetEntityData() *types.CommonEntityData {
    labelSessionBriefs.EntityData.YFilter = labelSessionBriefs.YFilter
    labelSessionBriefs.EntityData.YangName = "label-session-briefs"
    labelSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    labelSessionBriefs.EntityData.ParentYangName = "bfd"
    labelSessionBriefs.EntityData.SegmentPath = "label-session-briefs"
    labelSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSessionBriefs.EntityData.Children = types.NewOrderedMap()
    labelSessionBriefs.EntityData.Children.Append("label-session-brief", types.YChild{"LabelSessionBrief", nil})
    for i := range labelSessionBriefs.LabelSessionBrief {
        labelSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(labelSessionBriefs.LabelSessionBrief[i]), types.YChild{"LabelSessionBrief", labelSessionBriefs.LabelSessionBrief[i]})
    }
    labelSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    labelSessionBriefs.EntityData.YListKeys = []string {}

    return &(labelSessionBriefs.EntityData)
}

// Bfd_LabelSessionBriefs_LabelSessionBrief
// Brief information for a single Label BFD
// session
type Bfd_LabelSessionBriefs_LabelSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation
}

func (labelSessionBrief *Bfd_LabelSessionBriefs_LabelSessionBrief) GetEntityData() *types.CommonEntityData {
    labelSessionBrief.EntityData.YFilter = labelSessionBrief.YFilter
    labelSessionBrief.EntityData.YangName = "label-session-brief"
    labelSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    labelSessionBrief.EntityData.ParentYangName = "label-session-briefs"
    labelSessionBrief.EntityData.SegmentPath = "label-session-brief"
    labelSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSessionBrief.EntityData.Children = types.NewOrderedMap()
    labelSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &labelSessionBrief.StatusBriefInformation})
    labelSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    labelSessionBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", labelSessionBrief.InterfaceName})
    labelSessionBrief.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", labelSessionBrief.IncomingLabel})
    labelSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", labelSessionBrief.Location})
    labelSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", labelSessionBrief.NodeId})
    labelSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", labelSessionBrief.State})
    labelSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", labelSessionBrief.SessionType})
    labelSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", labelSessionBrief.SessionSubtype})
    labelSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", labelSessionBrief.SessionFlags})

    labelSessionBrief.EntityData.YListKeys = []string {}

    return &(labelSessionBrief.EntityData)
}

// Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "label-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_LabelSessionBriefs_LabelSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSummary
// Summary information of IPv4 BFD over MPLS-TE
// Tail
type Bfd_Ipv4bfDoMplsteTailSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4bfDoMplsteTailSummary_SessionState
}

func (ipv4bfDoMplsteTailSummary *Bfd_Ipv4bfDoMplsteTailSummary) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailSummary.EntityData.YFilter = ipv4bfDoMplsteTailSummary.YFilter
    ipv4bfDoMplsteTailSummary.EntityData.YangName = "ipv4bf-do-mplste-tail-summary"
    ipv4bfDoMplsteTailSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailSummary.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteTailSummary.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-summary"
    ipv4bfDoMplsteTailSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailSummary.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4bfDoMplsteTailSummary.SessionState})
    ipv4bfDoMplsteTailSummary.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteTailSummary.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailSummary.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4bfDoMplsteTailSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Ipv4bfDoMplsteTailSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv6SingleHopCounters
// IPv6 single hop Counters
type Bfd_Ipv6SingleHopCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv6 single hop Packet counters.
    Ipv6SingleHopPacketCounters Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters
}

func (ipv6SingleHopCounters *Bfd_Ipv6SingleHopCounters) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopCounters.EntityData.YFilter = ipv6SingleHopCounters.YFilter
    ipv6SingleHopCounters.EntityData.YangName = "ipv6-single-hop-counters"
    ipv6SingleHopCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopCounters.EntityData.ParentYangName = "bfd"
    ipv6SingleHopCounters.EntityData.SegmentPath = "ipv6-single-hop-counters"
    ipv6SingleHopCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopCounters.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopCounters.EntityData.Children.Append("ipv6-single-hop-packet-counters", types.YChild{"Ipv6SingleHopPacketCounters", &ipv6SingleHopCounters.Ipv6SingleHopPacketCounters})
    ipv6SingleHopCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopCounters.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopCounters.EntityData)
}

// Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters
// Table of IPv6 single hop Packet counters
type Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface IPv6 single hop Packet counters. The type is slice of
    // Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters_Ipv6SingleHopPacketCounter.
    Ipv6SingleHopPacketCounter []*Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters_Ipv6SingleHopPacketCounter
}

func (ipv6SingleHopPacketCounters *Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopPacketCounters.EntityData.YFilter = ipv6SingleHopPacketCounters.YFilter
    ipv6SingleHopPacketCounters.EntityData.YangName = "ipv6-single-hop-packet-counters"
    ipv6SingleHopPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopPacketCounters.EntityData.ParentYangName = "ipv6-single-hop-counters"
    ipv6SingleHopPacketCounters.EntityData.SegmentPath = "ipv6-single-hop-packet-counters"
    ipv6SingleHopPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopPacketCounters.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopPacketCounters.EntityData.Children.Append("ipv6-single-hop-packet-counter", types.YChild{"Ipv6SingleHopPacketCounter", nil})
    for i := range ipv6SingleHopPacketCounters.Ipv6SingleHopPacketCounter {
        ipv6SingleHopPacketCounters.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopPacketCounters.Ipv6SingleHopPacketCounter[i]), types.YChild{"Ipv6SingleHopPacketCounter", ipv6SingleHopPacketCounters.Ipv6SingleHopPacketCounter[i]})
    }
    ipv6SingleHopPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopPacketCounters.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopPacketCounters.EntityData)
}

// Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters_Ipv6SingleHopPacketCounter
// Interface IPv6 single hop Packet counters
type Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters_Ipv6SingleHopPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (ipv6SingleHopPacketCounter *Bfd_Ipv6SingleHopCounters_Ipv6SingleHopPacketCounters_Ipv6SingleHopPacketCounter) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopPacketCounter.EntityData.YFilter = ipv6SingleHopPacketCounter.YFilter
    ipv6SingleHopPacketCounter.EntityData.YangName = "ipv6-single-hop-packet-counter"
    ipv6SingleHopPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopPacketCounter.EntityData.ParentYangName = "ipv6-single-hop-packet-counters"
    ipv6SingleHopPacketCounter.EntityData.SegmentPath = "ipv6-single-hop-packet-counter"
    ipv6SingleHopPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopPacketCounter.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv6SingleHopPacketCounter.InterfaceName})
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6SingleHopPacketCounter.Location})
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", ipv6SingleHopPacketCounter.HelloTransmitCount})
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", ipv6SingleHopPacketCounter.HelloReceiveCount})
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", ipv6SingleHopPacketCounter.EchoTransmitCount})
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", ipv6SingleHopPacketCounter.EchoReceiveCount})
    ipv6SingleHopPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", ipv6SingleHopPacketCounter.DisplayType})

    ipv6SingleHopPacketCounter.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopPacketCounter.EntityData)
}

// Bfd_Counters
// IPv4 Counters
type Bfd_Counters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv4 Packet counters.
    PacketCounters Bfd_Counters_PacketCounters
}

func (counters *Bfd_Counters) GetEntityData() *types.CommonEntityData {
    counters.EntityData.YFilter = counters.YFilter
    counters.EntityData.YangName = "counters"
    counters.EntityData.BundleName = "cisco_ios_xr"
    counters.EntityData.ParentYangName = "bfd"
    counters.EntityData.SegmentPath = "counters"
    counters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    counters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    counters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    counters.EntityData.Children = types.NewOrderedMap()
    counters.EntityData.Children.Append("packet-counters", types.YChild{"PacketCounters", &counters.PacketCounters})
    counters.EntityData.Leafs = types.NewOrderedMap()

    counters.EntityData.YListKeys = []string {}

    return &(counters.EntityData)
}

// Bfd_Counters_PacketCounters
// Table of IPv4 Packet counters
type Bfd_Counters_PacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface IPv4 Packet counters. The type is slice of
    // Bfd_Counters_PacketCounters_PacketCounter.
    PacketCounter []*Bfd_Counters_PacketCounters_PacketCounter
}

func (packetCounters *Bfd_Counters_PacketCounters) GetEntityData() *types.CommonEntityData {
    packetCounters.EntityData.YFilter = packetCounters.YFilter
    packetCounters.EntityData.YangName = "packet-counters"
    packetCounters.EntityData.BundleName = "cisco_ios_xr"
    packetCounters.EntityData.ParentYangName = "counters"
    packetCounters.EntityData.SegmentPath = "packet-counters"
    packetCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetCounters.EntityData.Children = types.NewOrderedMap()
    packetCounters.EntityData.Children.Append("packet-counter", types.YChild{"PacketCounter", nil})
    for i := range packetCounters.PacketCounter {
        packetCounters.EntityData.Children.Append(types.GetSegmentPath(packetCounters.PacketCounter[i]), types.YChild{"PacketCounter", packetCounters.PacketCounter[i]})
    }
    packetCounters.EntityData.Leafs = types.NewOrderedMap()

    packetCounters.EntityData.YListKeys = []string {}

    return &(packetCounters.EntityData)
}

// Bfd_Counters_PacketCounters_PacketCounter
// Interface IPv4 Packet counters
type Bfd_Counters_PacketCounters_PacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (packetCounter *Bfd_Counters_PacketCounters_PacketCounter) GetEntityData() *types.CommonEntityData {
    packetCounter.EntityData.YFilter = packetCounter.YFilter
    packetCounter.EntityData.YangName = "packet-counter"
    packetCounter.EntityData.BundleName = "cisco_ios_xr"
    packetCounter.EntityData.ParentYangName = "packet-counters"
    packetCounter.EntityData.SegmentPath = "packet-counter"
    packetCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetCounter.EntityData.Children = types.NewOrderedMap()
    packetCounter.EntityData.Leafs = types.NewOrderedMap()
    packetCounter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", packetCounter.InterfaceName})
    packetCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", packetCounter.Location})
    packetCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", packetCounter.HelloTransmitCount})
    packetCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", packetCounter.HelloReceiveCount})
    packetCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", packetCounter.EchoTransmitCount})
    packetCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", packetCounter.EchoReceiveCount})
    packetCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", packetCounter.DisplayType})

    packetCounter.EntityData.YListKeys = []string {}

    return &(packetCounter.EntityData)
}

// Bfd_ClientDetails
// Table of detailed information about BFD clients
type Bfd_ClientDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information of client. The type is slice of
    // Bfd_ClientDetails_ClientDetail.
    ClientDetail []*Bfd_ClientDetails_ClientDetail
}

func (clientDetails *Bfd_ClientDetails) GetEntityData() *types.CommonEntityData {
    clientDetails.EntityData.YFilter = clientDetails.YFilter
    clientDetails.EntityData.YangName = "client-details"
    clientDetails.EntityData.BundleName = "cisco_ios_xr"
    clientDetails.EntityData.ParentYangName = "bfd"
    clientDetails.EntityData.SegmentPath = "client-details"
    clientDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientDetails.EntityData.Children = types.NewOrderedMap()
    clientDetails.EntityData.Children.Append("client-detail", types.YChild{"ClientDetail", nil})
    for i := range clientDetails.ClientDetail {
        clientDetails.EntityData.Children.Append(types.GetSegmentPath(clientDetails.ClientDetail[i]), types.YChild{"ClientDetail", clientDetails.ClientDetail[i]})
    }
    clientDetails.EntityData.Leafs = types.NewOrderedMap()

    clientDetails.EntityData.YListKeys = []string {}

    return &(clientDetails.EntityData)
}

// Bfd_ClientDetails_ClientDetail
// Detailed information of client
type Bfd_ClientDetails_ClientDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientName interface{}

    // Recreate Time in Seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    RecreateTime interface{}

    // Brief client information.
    Brief Bfd_ClientDetails_ClientDetail_Brief

    // The BFD Client Flags.
    Flags Bfd_ClientDetails_ClientDetail_Flags
}

func (clientDetail *Bfd_ClientDetails_ClientDetail) GetEntityData() *types.CommonEntityData {
    clientDetail.EntityData.YFilter = clientDetail.YFilter
    clientDetail.EntityData.YangName = "client-detail"
    clientDetail.EntityData.BundleName = "cisco_ios_xr"
    clientDetail.EntityData.ParentYangName = "client-details"
    clientDetail.EntityData.SegmentPath = "client-detail" + types.AddKeyToken(clientDetail.ClientName, "client-name")
    clientDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientDetail.EntityData.Children = types.NewOrderedMap()
    clientDetail.EntityData.Children.Append("brief", types.YChild{"Brief", &clientDetail.Brief})
    clientDetail.EntityData.Children.Append("flags", types.YChild{"Flags", &clientDetail.Flags})
    clientDetail.EntityData.Leafs = types.NewOrderedMap()
    clientDetail.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", clientDetail.ClientName})
    clientDetail.EntityData.Leafs.Append("recreate-time", types.YLeaf{"RecreateTime", clientDetail.RecreateTime})

    clientDetail.EntityData.YListKeys = []string {"ClientName"}

    return &(clientDetail.EntityData)
}

// Bfd_ClientDetails_ClientDetail_Brief
// Brief client information
type Bfd_ClientDetails_ClientDetail_Brief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client process name. The type is string with length: 0..257.
    NameXr interface{}

    // Location where client resides. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Number of sessions created by this client. The type is interface{} with
    // range: 0..4294967295.
    SessionCount interface{}
}

func (brief *Bfd_ClientDetails_ClientDetail_Brief) GetEntityData() *types.CommonEntityData {
    brief.EntityData.YFilter = brief.YFilter
    brief.EntityData.YangName = "brief"
    brief.EntityData.BundleName = "cisco_ios_xr"
    brief.EntityData.ParentYangName = "client-detail"
    brief.EntityData.SegmentPath = "brief"
    brief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    brief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    brief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    brief.EntityData.Children = types.NewOrderedMap()
    brief.EntityData.Leafs = types.NewOrderedMap()
    brief.EntityData.Leafs.Append("name-xr", types.YLeaf{"NameXr", brief.NameXr})
    brief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", brief.NodeId})
    brief.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", brief.SessionCount})

    brief.EntityData.YListKeys = []string {}

    return &(brief.EntityData)
}

// Bfd_ClientDetails_ClientDetail_Flags
// The BFD Client Flags
type Bfd_ClientDetails_ClientDetail_Flags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client is in Zombie State. The type is interface{} with range:
    // -2147483648..2147483647.
    IsZombieState interface{}

    // Client is in Recreate State. The type is interface{} with range:
    // -2147483648..2147483647.
    IsRecreateState interface{}
}

func (flags *Bfd_ClientDetails_ClientDetail_Flags) GetEntityData() *types.CommonEntityData {
    flags.EntityData.YFilter = flags.YFilter
    flags.EntityData.YangName = "flags"
    flags.EntityData.BundleName = "cisco_ios_xr"
    flags.EntityData.ParentYangName = "client-detail"
    flags.EntityData.SegmentPath = "flags"
    flags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flags.EntityData.Children = types.NewOrderedMap()
    flags.EntityData.Leafs = types.NewOrderedMap()
    flags.EntityData.Leafs.Append("is-zombie-state", types.YLeaf{"IsZombieState", flags.IsZombieState})
    flags.EntityData.Leafs.Append("is-recreate-state", types.YLeaf{"IsRecreateState", flags.IsRecreateState})

    flags.EntityData.YListKeys = []string {}

    return &(flags.EntityData)
}

// Bfd_Ipv4SingleHopSummary
// Summary information of BFD IPv4 singlehop
// sessions
type Bfd_Ipv4SingleHopSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4SingleHopSummary_SessionState
}

func (ipv4SingleHopSummary *Bfd_Ipv4SingleHopSummary) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopSummary.EntityData.YFilter = ipv4SingleHopSummary.YFilter
    ipv4SingleHopSummary.EntityData.YangName = "ipv4-single-hop-summary"
    ipv4SingleHopSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopSummary.EntityData.ParentYangName = "bfd"
    ipv4SingleHopSummary.EntityData.SegmentPath = "ipv4-single-hop-summary"
    ipv4SingleHopSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopSummary.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4SingleHopSummary.SessionState})
    ipv4SingleHopSummary.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopSummary.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopSummary.EntityData)
}

// Bfd_Ipv4SingleHopSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4SingleHopSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Ipv4SingleHopSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4-single-hop-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv6SingleHopSummary
// Summary information of BFD IPv6 singlehop
// sessions
type Bfd_Ipv6SingleHopSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Ipv6SingleHopSummary_SessionState
}

func (ipv6SingleHopSummary *Bfd_Ipv6SingleHopSummary) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopSummary.EntityData.YFilter = ipv6SingleHopSummary.YFilter
    ipv6SingleHopSummary.EntityData.YangName = "ipv6-single-hop-summary"
    ipv6SingleHopSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopSummary.EntityData.ParentYangName = "bfd"
    ipv6SingleHopSummary.EntityData.SegmentPath = "ipv6-single-hop-summary"
    ipv6SingleHopSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopSummary.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv6SingleHopSummary.SessionState})
    ipv6SingleHopSummary.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopSummary.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopSummary.EntityData)
}

// Bfd_Ipv6SingleHopSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv6SingleHopSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Ipv6SingleHopSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv6-single-hop-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_LabelMultiPaths
// Label multipath
type Bfd_LabelMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label multipath table. The type is slice of
    // Bfd_LabelMultiPaths_LabelMultiPath.
    LabelMultiPath []*Bfd_LabelMultiPaths_LabelMultiPath
}

func (labelMultiPaths *Bfd_LabelMultiPaths) GetEntityData() *types.CommonEntityData {
    labelMultiPaths.EntityData.YFilter = labelMultiPaths.YFilter
    labelMultiPaths.EntityData.YangName = "label-multi-paths"
    labelMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    labelMultiPaths.EntityData.ParentYangName = "bfd"
    labelMultiPaths.EntityData.SegmentPath = "label-multi-paths"
    labelMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelMultiPaths.EntityData.Children = types.NewOrderedMap()
    labelMultiPaths.EntityData.Children.Append("label-multi-path", types.YChild{"LabelMultiPath", nil})
    for i := range labelMultiPaths.LabelMultiPath {
        labelMultiPaths.EntityData.Children.Append(types.GetSegmentPath(labelMultiPaths.LabelMultiPath[i]), types.YChild{"LabelMultiPath", labelMultiPaths.LabelMultiPath[i]})
    }
    labelMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    labelMultiPaths.EntityData.YListKeys = []string {}

    return &(labelMultiPaths.EntityData)
}

// Bfd_LabelMultiPaths_LabelMultiPath
// Label multipath table
type Bfd_LabelMultiPaths_LabelMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (labelMultiPath *Bfd_LabelMultiPaths_LabelMultiPath) GetEntityData() *types.CommonEntityData {
    labelMultiPath.EntityData.YFilter = labelMultiPath.YFilter
    labelMultiPath.EntityData.YangName = "label-multi-path"
    labelMultiPath.EntityData.BundleName = "cisco_ios_xr"
    labelMultiPath.EntityData.ParentYangName = "label-multi-paths"
    labelMultiPath.EntityData.SegmentPath = "label-multi-path"
    labelMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelMultiPath.EntityData.Children = types.NewOrderedMap()
    labelMultiPath.EntityData.Leafs = types.NewOrderedMap()
    labelMultiPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", labelMultiPath.InterfaceName})
    labelMultiPath.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", labelMultiPath.IncomingLabel})
    labelMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", labelMultiPath.Location})
    labelMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", labelMultiPath.SessionSubtype})
    labelMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", labelMultiPath.State})
    labelMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", labelMultiPath.LocalDiscriminator})
    labelMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", labelMultiPath.NodeId})
    labelMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", labelMultiPath.IncomingLabelXr})
    labelMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", labelMultiPath.SessionInterfaceName})

    labelMultiPath.EntityData.YListKeys = []string {}

    return &(labelMultiPath.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails
// Table of detailed information about all IPv4
// multihop BFD sessions in the System 
type Bfd_Ipv4MultiHopSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv4 multihop BFD session. The type is
    // slice of Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail.
    Ipv4MultiHopSessionDetail []*Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail
}

func (ipv4MultiHopSessionDetails *Bfd_Ipv4MultiHopSessionDetails) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopSessionDetails.EntityData.YFilter = ipv4MultiHopSessionDetails.YFilter
    ipv4MultiHopSessionDetails.EntityData.YangName = "ipv4-multi-hop-session-details"
    ipv4MultiHopSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopSessionDetails.EntityData.ParentYangName = "bfd"
    ipv4MultiHopSessionDetails.EntityData.SegmentPath = "ipv4-multi-hop-session-details"
    ipv4MultiHopSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopSessionDetails.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopSessionDetails.EntityData.Children.Append("ipv4-multi-hop-session-detail", types.YChild{"Ipv4MultiHopSessionDetail", nil})
    for i := range ipv4MultiHopSessionDetails.Ipv4MultiHopSessionDetail {
        ipv4MultiHopSessionDetails.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopSessionDetails.Ipv4MultiHopSessionDetail[i]), types.YChild{"Ipv4MultiHopSessionDetail", ipv4MultiHopSessionDetails.Ipv4MultiHopSessionDetail[i]})
    }
    ipv4MultiHopSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopSessionDetails.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopSessionDetails.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail
// Detailed information for a single IPv4 multihop
// BFD session
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Session status information.
    StatusInformation Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation
}

func (ipv4MultiHopSessionDetail *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopSessionDetail.EntityData.YFilter = ipv4MultiHopSessionDetail.YFilter
    ipv4MultiHopSessionDetail.EntityData.YangName = "ipv4-multi-hop-session-detail"
    ipv4MultiHopSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopSessionDetail.EntityData.ParentYangName = "ipv4-multi-hop-session-details"
    ipv4MultiHopSessionDetail.EntityData.SegmentPath = "ipv4-multi-hop-session-detail"
    ipv4MultiHopSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopSessionDetail.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &ipv4MultiHopSessionDetail.StatusInformation})
    ipv4MultiHopSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &ipv4MultiHopSessionDetail.MpDownloadState})
    ipv4MultiHopSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &ipv4MultiHopSessionDetail.LspPingInfo})
    ipv4MultiHopSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range ipv4MultiHopSessionDetail.OwnerInformation {
        ipv4MultiHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", ipv4MultiHopSessionDetail.OwnerInformation[i]})
    }
    ipv4MultiHopSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range ipv4MultiHopSessionDetail.AssociationInformation {
        ipv4MultiHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", ipv4MultiHopSessionDetail.AssociationInformation[i]})
    }
    ipv4MultiHopSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    ipv4MultiHopSessionDetail.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv4MultiHopSessionDetail.SourceAddress})
    ipv4MultiHopSessionDetail.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4MultiHopSessionDetail.DestinationAddress})
    ipv4MultiHopSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4MultiHopSessionDetail.Location})
    ipv4MultiHopSessionDetail.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4MultiHopSessionDetail.VrfName})

    ipv4MultiHopSessionDetail.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopSessionDetail.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation
// Session status information
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "ipv4-multi-hop-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "ipv4-multi-hop-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "ipv4-multi-hop-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "ipv4-multi-hop-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation
// Association session information
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "ipv4-multi-hop-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4MultiHopSessionDetails_Ipv4MultiHopSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails
// Table of detailed information about all IPv4
// singlehop BFD sessions in the System 
type Bfd_Ipv4SingleHopSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv4 singlehop BFD session. The type is
    // slice of Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail.
    Ipv4SingleHopSessionDetail []*Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail
}

func (ipv4SingleHopSessionDetails *Bfd_Ipv4SingleHopSessionDetails) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopSessionDetails.EntityData.YFilter = ipv4SingleHopSessionDetails.YFilter
    ipv4SingleHopSessionDetails.EntityData.YangName = "ipv4-single-hop-session-details"
    ipv4SingleHopSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopSessionDetails.EntityData.ParentYangName = "bfd"
    ipv4SingleHopSessionDetails.EntityData.SegmentPath = "ipv4-single-hop-session-details"
    ipv4SingleHopSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopSessionDetails.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopSessionDetails.EntityData.Children.Append("ipv4-single-hop-session-detail", types.YChild{"Ipv4SingleHopSessionDetail", nil})
    for i := range ipv4SingleHopSessionDetails.Ipv4SingleHopSessionDetail {
        ipv4SingleHopSessionDetails.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopSessionDetails.Ipv4SingleHopSessionDetail[i]), types.YChild{"Ipv4SingleHopSessionDetail", ipv4SingleHopSessionDetails.Ipv4SingleHopSessionDetail[i]})
    }
    ipv4SingleHopSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopSessionDetails.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopSessionDetails.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail
// Detailed information for a single IPv4
// singlehop BFD session
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session status information.
    StatusInformation Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation
}

func (ipv4SingleHopSessionDetail *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopSessionDetail.EntityData.YFilter = ipv4SingleHopSessionDetail.YFilter
    ipv4SingleHopSessionDetail.EntityData.YangName = "ipv4-single-hop-session-detail"
    ipv4SingleHopSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopSessionDetail.EntityData.ParentYangName = "ipv4-single-hop-session-details"
    ipv4SingleHopSessionDetail.EntityData.SegmentPath = "ipv4-single-hop-session-detail"
    ipv4SingleHopSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopSessionDetail.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &ipv4SingleHopSessionDetail.StatusInformation})
    ipv4SingleHopSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &ipv4SingleHopSessionDetail.MpDownloadState})
    ipv4SingleHopSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &ipv4SingleHopSessionDetail.LspPingInfo})
    ipv4SingleHopSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range ipv4SingleHopSessionDetail.OwnerInformation {
        ipv4SingleHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", ipv4SingleHopSessionDetail.OwnerInformation[i]})
    }
    ipv4SingleHopSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range ipv4SingleHopSessionDetail.AssociationInformation {
        ipv4SingleHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", ipv4SingleHopSessionDetail.AssociationInformation[i]})
    }
    ipv4SingleHopSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    ipv4SingleHopSessionDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4SingleHopSessionDetail.InterfaceName})
    ipv4SingleHopSessionDetail.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4SingleHopSessionDetail.DestinationAddress})
    ipv4SingleHopSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4SingleHopSessionDetail.Location})

    ipv4SingleHopSessionDetail.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopSessionDetail.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation
// Session status information
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "ipv4-single-hop-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "ipv4-single-hop-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "ipv4-single-hop-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "ipv4-single-hop-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation
// Association session information
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "ipv4-single-hop-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4SingleHopSessionDetails_Ipv4SingleHopSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4MultiHopSessionBriefs
// Table of brief information about all IPv4
// multihop BFD sessions in the System
type Bfd_Ipv4MultiHopSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv4 multihop BFD session. The type is slice
    // of Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief.
    Ipv4MultiHopSessionBrief []*Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief
}

func (ipv4MultiHopSessionBriefs *Bfd_Ipv4MultiHopSessionBriefs) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopSessionBriefs.EntityData.YFilter = ipv4MultiHopSessionBriefs.YFilter
    ipv4MultiHopSessionBriefs.EntityData.YangName = "ipv4-multi-hop-session-briefs"
    ipv4MultiHopSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopSessionBriefs.EntityData.ParentYangName = "bfd"
    ipv4MultiHopSessionBriefs.EntityData.SegmentPath = "ipv4-multi-hop-session-briefs"
    ipv4MultiHopSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopSessionBriefs.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopSessionBriefs.EntityData.Children.Append("ipv4-multi-hop-session-brief", types.YChild{"Ipv4MultiHopSessionBrief", nil})
    for i := range ipv4MultiHopSessionBriefs.Ipv4MultiHopSessionBrief {
        ipv4MultiHopSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopSessionBriefs.Ipv4MultiHopSessionBrief[i]), types.YChild{"Ipv4MultiHopSessionBrief", ipv4MultiHopSessionBriefs.Ipv4MultiHopSessionBrief[i]})
    }
    ipv4MultiHopSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopSessionBriefs.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopSessionBriefs.EntityData)
}

// Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief
// Brief information for a single IPv4 multihop
// BFD session
type Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation
}

func (ipv4MultiHopSessionBrief *Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopSessionBrief.EntityData.YFilter = ipv4MultiHopSessionBrief.YFilter
    ipv4MultiHopSessionBrief.EntityData.YangName = "ipv4-multi-hop-session-brief"
    ipv4MultiHopSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopSessionBrief.EntityData.ParentYangName = "ipv4-multi-hop-session-briefs"
    ipv4MultiHopSessionBrief.EntityData.SegmentPath = "ipv4-multi-hop-session-brief"
    ipv4MultiHopSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopSessionBrief.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &ipv4MultiHopSessionBrief.StatusBriefInformation})
    ipv4MultiHopSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv4MultiHopSessionBrief.SourceAddress})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4MultiHopSessionBrief.DestinationAddress})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4MultiHopSessionBrief.Location})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4MultiHopSessionBrief.VrfName})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4MultiHopSessionBrief.NodeId})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4MultiHopSessionBrief.State})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", ipv4MultiHopSessionBrief.SessionType})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4MultiHopSessionBrief.SessionSubtype})
    ipv4MultiHopSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", ipv4MultiHopSessionBrief.SessionFlags})

    ipv4MultiHopSessionBrief.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopSessionBrief.EntityData)
}

// Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "ipv4-multi-hop-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4MultiHopSessionBriefs_Ipv4MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_GenericSummaries
// Generic summary information about BFD location
type Bfd_GenericSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Generic summary information for bfd location table. The type is slice of
    // Bfd_GenericSummaries_GenericSummary.
    GenericSummary []*Bfd_GenericSummaries_GenericSummary
}

func (genericSummaries *Bfd_GenericSummaries) GetEntityData() *types.CommonEntityData {
    genericSummaries.EntityData.YFilter = genericSummaries.YFilter
    genericSummaries.EntityData.YangName = "generic-summaries"
    genericSummaries.EntityData.BundleName = "cisco_ios_xr"
    genericSummaries.EntityData.ParentYangName = "bfd"
    genericSummaries.EntityData.SegmentPath = "generic-summaries"
    genericSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericSummaries.EntityData.Children = types.NewOrderedMap()
    genericSummaries.EntityData.Children.Append("generic-summary", types.YChild{"GenericSummary", nil})
    for i := range genericSummaries.GenericSummary {
        genericSummaries.EntityData.Children.Append(types.GetSegmentPath(genericSummaries.GenericSummary[i]), types.YChild{"GenericSummary", genericSummaries.GenericSummary[i]})
    }
    genericSummaries.EntityData.Leafs = types.NewOrderedMap()

    genericSummaries.EntityData.YListKeys = []string {}

    return &(genericSummaries.EntityData)
}

// Bfd_GenericSummaries_GenericSummary
// Generic summary information for bfd location
// table
type Bfd_GenericSummaries_GenericSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Node ID. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Allocated PPS value. The type is interface{} with range: 0..4294967295.
    PpsAllocatedValue interface{}

    // Allocated MP PPS value. The type is interface{} with range: 0..4294967295.
    PpsmpAllocatedValue interface{}

    // Max PPS value. The type is interface{} with range: 0..4294967295.
    PpsMaxValue interface{}

    // Max MP PPS value. The type is interface{} with range: 0..4294967295.
    PpsmpMaxValue interface{}

    // Total Session Number. The type is interface{} with range: 0..4294967295.
    TotalSessionNumber interface{}

    // MP Session Number. The type is interface{} with range: 0..4294967295.
    MpSessionNumber interface{}

    // Max Session Number. The type is interface{} with range: 0..4294967295.
    MaxSessionNumber interface{}

    // All PPS percentage. The type is interface{} with range: 0..4294967295.
    // Units are percentage.
    PpsAllPercentage interface{}

    // MP PPS percentage. The type is interface{} with range: 0..4294967295. Units
    // are percentage.
    PpsmpPercentage interface{}
}

func (genericSummary *Bfd_GenericSummaries_GenericSummary) GetEntityData() *types.CommonEntityData {
    genericSummary.EntityData.YFilter = genericSummary.YFilter
    genericSummary.EntityData.YangName = "generic-summary"
    genericSummary.EntityData.BundleName = "cisco_ios_xr"
    genericSummary.EntityData.ParentYangName = "generic-summaries"
    genericSummary.EntityData.SegmentPath = "generic-summary" + types.AddKeyToken(genericSummary.Location, "location")
    genericSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    genericSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    genericSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    genericSummary.EntityData.Children = types.NewOrderedMap()
    genericSummary.EntityData.Leafs = types.NewOrderedMap()
    genericSummary.EntityData.Leafs.Append("location", types.YLeaf{"Location", genericSummary.Location})
    genericSummary.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", genericSummary.NodeId})
    genericSummary.EntityData.Leafs.Append("pps-allocated-value", types.YLeaf{"PpsAllocatedValue", genericSummary.PpsAllocatedValue})
    genericSummary.EntityData.Leafs.Append("ppsmp-allocated-value", types.YLeaf{"PpsmpAllocatedValue", genericSummary.PpsmpAllocatedValue})
    genericSummary.EntityData.Leafs.Append("pps-max-value", types.YLeaf{"PpsMaxValue", genericSummary.PpsMaxValue})
    genericSummary.EntityData.Leafs.Append("ppsmp-max-value", types.YLeaf{"PpsmpMaxValue", genericSummary.PpsmpMaxValue})
    genericSummary.EntityData.Leafs.Append("total-session-number", types.YLeaf{"TotalSessionNumber", genericSummary.TotalSessionNumber})
    genericSummary.EntityData.Leafs.Append("mp-session-number", types.YLeaf{"MpSessionNumber", genericSummary.MpSessionNumber})
    genericSummary.EntityData.Leafs.Append("max-session-number", types.YLeaf{"MaxSessionNumber", genericSummary.MaxSessionNumber})
    genericSummary.EntityData.Leafs.Append("pps-all-percentage", types.YLeaf{"PpsAllPercentage", genericSummary.PpsAllPercentage})
    genericSummary.EntityData.Leafs.Append("ppsmp-percentage", types.YLeaf{"PpsmpPercentage", genericSummary.PpsmpPercentage})

    genericSummary.EntityData.YListKeys = []string {"Location"}

    return &(genericSummary.EntityData)
}

// Bfd_Ipv6SingleHopMultiPaths
// IPv6 single hop multipath
type Bfd_Ipv6SingleHopMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 single hop multipath table. The type is slice of
    // Bfd_Ipv6SingleHopMultiPaths_Ipv6SingleHopMultiPath.
    Ipv6SingleHopMultiPath []*Bfd_Ipv6SingleHopMultiPaths_Ipv6SingleHopMultiPath
}

func (ipv6SingleHopMultiPaths *Bfd_Ipv6SingleHopMultiPaths) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopMultiPaths.EntityData.YFilter = ipv6SingleHopMultiPaths.YFilter
    ipv6SingleHopMultiPaths.EntityData.YangName = "ipv6-single-hop-multi-paths"
    ipv6SingleHopMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopMultiPaths.EntityData.ParentYangName = "bfd"
    ipv6SingleHopMultiPaths.EntityData.SegmentPath = "ipv6-single-hop-multi-paths"
    ipv6SingleHopMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopMultiPaths.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopMultiPaths.EntityData.Children.Append("ipv6-single-hop-multi-path", types.YChild{"Ipv6SingleHopMultiPath", nil})
    for i := range ipv6SingleHopMultiPaths.Ipv6SingleHopMultiPath {
        ipv6SingleHopMultiPaths.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopMultiPaths.Ipv6SingleHopMultiPath[i]), types.YChild{"Ipv6SingleHopMultiPath", ipv6SingleHopMultiPaths.Ipv6SingleHopMultiPath[i]})
    }
    ipv6SingleHopMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopMultiPaths.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopMultiPaths.EntityData)
}

// Bfd_Ipv6SingleHopMultiPaths_Ipv6SingleHopMultiPath
// IPv6 single hop multipath table
type Bfd_Ipv6SingleHopMultiPaths_Ipv6SingleHopMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (ipv6SingleHopMultiPath *Bfd_Ipv6SingleHopMultiPaths_Ipv6SingleHopMultiPath) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopMultiPath.EntityData.YFilter = ipv6SingleHopMultiPath.YFilter
    ipv6SingleHopMultiPath.EntityData.YangName = "ipv6-single-hop-multi-path"
    ipv6SingleHopMultiPath.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopMultiPath.EntityData.ParentYangName = "ipv6-single-hop-multi-paths"
    ipv6SingleHopMultiPath.EntityData.SegmentPath = "ipv6-single-hop-multi-path"
    ipv6SingleHopMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopMultiPath.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopMultiPath.EntityData.Leafs = types.NewOrderedMap()
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv6SingleHopMultiPath.InterfaceName})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6SingleHopMultiPath.DestinationAddress})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6SingleHopMultiPath.Location})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv6SingleHopMultiPath.SessionSubtype})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv6SingleHopMultiPath.State})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", ipv6SingleHopMultiPath.LocalDiscriminator})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv6SingleHopMultiPath.NodeId})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", ipv6SingleHopMultiPath.IncomingLabelXr})
    ipv6SingleHopMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", ipv6SingleHopMultiPath.SessionInterfaceName})

    ipv6SingleHopMultiPath.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopMultiPath.EntityData)
}

// Bfd_Ipv4SingleHopNodeLocationSummaries
// Table of summary information about BFD IPv4
// singlehop sessions per location
type Bfd_Ipv4SingleHopNodeLocationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for BFD IPv4 singlehop sessions for location. The type
    // is slice of
    // Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary.
    Ipv4SingleHopNodeLocationSummary []*Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary
}

func (ipv4SingleHopNodeLocationSummaries *Bfd_Ipv4SingleHopNodeLocationSummaries) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopNodeLocationSummaries.EntityData.YFilter = ipv4SingleHopNodeLocationSummaries.YFilter
    ipv4SingleHopNodeLocationSummaries.EntityData.YangName = "ipv4-single-hop-node-location-summaries"
    ipv4SingleHopNodeLocationSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopNodeLocationSummaries.EntityData.ParentYangName = "bfd"
    ipv4SingleHopNodeLocationSummaries.EntityData.SegmentPath = "ipv4-single-hop-node-location-summaries"
    ipv4SingleHopNodeLocationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopNodeLocationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopNodeLocationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopNodeLocationSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopNodeLocationSummaries.EntityData.Children.Append("ipv4-single-hop-node-location-summary", types.YChild{"Ipv4SingleHopNodeLocationSummary", nil})
    for i := range ipv4SingleHopNodeLocationSummaries.Ipv4SingleHopNodeLocationSummary {
        ipv4SingleHopNodeLocationSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopNodeLocationSummaries.Ipv4SingleHopNodeLocationSummary[i]), types.YChild{"Ipv4SingleHopNodeLocationSummary", ipv4SingleHopNodeLocationSummaries.Ipv4SingleHopNodeLocationSummary[i]})
    }
    ipv4SingleHopNodeLocationSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopNodeLocationSummaries.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopNodeLocationSummaries.EntityData)
}

// Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary
// Summary information for BFD IPv4 singlehop
// sessions for location
type Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary_SessionState
}

func (ipv4SingleHopNodeLocationSummary *Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopNodeLocationSummary.EntityData.YFilter = ipv4SingleHopNodeLocationSummary.YFilter
    ipv4SingleHopNodeLocationSummary.EntityData.YangName = "ipv4-single-hop-node-location-summary"
    ipv4SingleHopNodeLocationSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopNodeLocationSummary.EntityData.ParentYangName = "ipv4-single-hop-node-location-summaries"
    ipv4SingleHopNodeLocationSummary.EntityData.SegmentPath = "ipv4-single-hop-node-location-summary" + types.AddKeyToken(ipv4SingleHopNodeLocationSummary.Location, "location")
    ipv4SingleHopNodeLocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopNodeLocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopNodeLocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopNodeLocationSummary.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopNodeLocationSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4SingleHopNodeLocationSummary.SessionState})
    ipv4SingleHopNodeLocationSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4SingleHopNodeLocationSummary.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4SingleHopNodeLocationSummary.Location})

    ipv4SingleHopNodeLocationSummary.EntityData.YListKeys = []string {"Location"}

    return &(ipv4SingleHopNodeLocationSummary.EntityData)
}

// Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv4SingleHopNodeLocationSummaries_Ipv4SingleHopNodeLocationSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4-single-hop-node-location-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_LabelSummary
// Summary information of Label BFD
type Bfd_LabelSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_LabelSummary_SessionState
}

func (labelSummary *Bfd_LabelSummary) GetEntityData() *types.CommonEntityData {
    labelSummary.EntityData.YFilter = labelSummary.YFilter
    labelSummary.EntityData.YangName = "label-summary"
    labelSummary.EntityData.BundleName = "cisco_ios_xr"
    labelSummary.EntityData.ParentYangName = "bfd"
    labelSummary.EntityData.SegmentPath = "label-summary"
    labelSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSummary.EntityData.Children = types.NewOrderedMap()
    labelSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &labelSummary.SessionState})
    labelSummary.EntityData.Leafs = types.NewOrderedMap()

    labelSummary.EntityData.YListKeys = []string {}

    return &(labelSummary.EntityData)
}

// Bfd_LabelSummary_SessionState
// Statistics of states for sessions
type Bfd_LabelSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_LabelSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "label-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionBriefs
// Table of brief information about all IPv4 BFD
// over MPLS-TE Head sessions in the System
type Bfd_Ipv4bfDoMplsteHeadSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv4 BFD over MPLS-TE Head session. The type
    // is slice of
    // Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief.
    Ipv4bfDoMplsteHeadSessionBrief []*Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief
}

func (ipv4bfDoMplsteHeadSessionBriefs *Bfd_Ipv4bfDoMplsteHeadSessionBriefs) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.YFilter = ipv4bfDoMplsteHeadSessionBriefs.YFilter
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.YangName = "ipv4bf-do-mplste-head-session-briefs"
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.SegmentPath = "ipv4bf-do-mplste-head-session-briefs"
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadSessionBriefs.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.Children.Append("ipv4bf-do-mplste-head-session-brief", types.YChild{"Ipv4bfDoMplsteHeadSessionBrief", nil})
    for i := range ipv4bfDoMplsteHeadSessionBriefs.Ipv4bfDoMplsteHeadSessionBrief {
        ipv4bfDoMplsteHeadSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteHeadSessionBriefs.Ipv4bfDoMplsteHeadSessionBrief[i]), types.YChild{"Ipv4bfDoMplsteHeadSessionBrief", ipv4bfDoMplsteHeadSessionBriefs.Ipv4bfDoMplsteHeadSessionBrief[i]})
    }
    ipv4bfDoMplsteHeadSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteHeadSessionBriefs.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadSessionBriefs.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief
// Brief information for a single IPv4 BFD over
// MPLS-TE Head session
type Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation
}

func (ipv4bfDoMplsteHeadSessionBrief *Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadSessionBrief.EntityData.YFilter = ipv4bfDoMplsteHeadSessionBrief.YFilter
    ipv4bfDoMplsteHeadSessionBrief.EntityData.YangName = "ipv4bf-do-mplste-head-session-brief"
    ipv4bfDoMplsteHeadSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadSessionBrief.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-briefs"
    ipv4bfDoMplsteHeadSessionBrief.EntityData.SegmentPath = "ipv4bf-do-mplste-head-session-brief"
    ipv4bfDoMplsteHeadSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadSessionBrief.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &ipv4bfDoMplsteHeadSessionBrief.StatusBriefInformation})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4bfDoMplsteHeadSessionBrief.InterfaceName})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteHeadSessionBrief.VrfName})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteHeadSessionBrief.IncomingLabel})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteHeadSessionBrief.FeCtype})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteHeadSessionBrief.FecSubgroupId})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteHeadSessionBrief.Feclspid})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteHeadSessionBrief.FecTunnelId})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteHeadSessionBrief.FecExtendedTunnelId})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteHeadSessionBrief.FecSource})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteHeadSessionBrief.FecDestination})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteHeadSessionBrief.Fecp2mpid})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteHeadSessionBrief.FecSubgroupOriginator})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteHeadSessionBrief.FecCtype})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteHeadSessionBrief.Location})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4bfDoMplsteHeadSessionBrief.NodeId})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4bfDoMplsteHeadSessionBrief.State})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", ipv4bfDoMplsteHeadSessionBrief.SessionType})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4bfDoMplsteHeadSessionBrief.SessionSubtype})
    ipv4bfDoMplsteHeadSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", ipv4bfDoMplsteHeadSessionBrief.SessionFlags})

    ipv4bfDoMplsteHeadSessionBrief.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadSessionBrief.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4bfDoMplsteHeadSessionBriefs_Ipv4bfDoMplsteHeadSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails
// Table of detailed information about all IPv4 BFD
// over MPLS-TE Tail sessions in the System
type Bfd_Ipv4bfDoMplsteTailSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv4 BFD over MPLS-TE Tail session. The
    // type is slice of
    // Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail.
    Ipv4bfDoMplsteTailSessionDetail []*Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail
}

func (ipv4bfDoMplsteTailSessionDetails *Bfd_Ipv4bfDoMplsteTailSessionDetails) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailSessionDetails.EntityData.YFilter = ipv4bfDoMplsteTailSessionDetails.YFilter
    ipv4bfDoMplsteTailSessionDetails.EntityData.YangName = "ipv4bf-do-mplste-tail-session-details"
    ipv4bfDoMplsteTailSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailSessionDetails.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteTailSessionDetails.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-session-details"
    ipv4bfDoMplsteTailSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailSessionDetails.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailSessionDetails.EntityData.Children.Append("ipv4bf-do-mplste-tail-session-detail", types.YChild{"Ipv4bfDoMplsteTailSessionDetail", nil})
    for i := range ipv4bfDoMplsteTailSessionDetails.Ipv4bfDoMplsteTailSessionDetail {
        ipv4bfDoMplsteTailSessionDetails.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteTailSessionDetails.Ipv4bfDoMplsteTailSessionDetail[i]), types.YChild{"Ipv4bfDoMplsteTailSessionDetail", ipv4bfDoMplsteTailSessionDetails.Ipv4bfDoMplsteTailSessionDetail[i]})
    }
    ipv4bfDoMplsteTailSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteTailSessionDetails.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailSessionDetails.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail
// Detailed information for a single IPv4 BFD over
// MPLS-TE Tail session
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session status information.
    StatusInformation Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation
}

func (ipv4bfDoMplsteTailSessionDetail *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailSessionDetail.EntityData.YFilter = ipv4bfDoMplsteTailSessionDetail.YFilter
    ipv4bfDoMplsteTailSessionDetail.EntityData.YangName = "ipv4bf-do-mplste-tail-session-detail"
    ipv4bfDoMplsteTailSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailSessionDetail.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-details"
    ipv4bfDoMplsteTailSessionDetail.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-session-detail"
    ipv4bfDoMplsteTailSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailSessionDetail.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &ipv4bfDoMplsteTailSessionDetail.StatusInformation})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &ipv4bfDoMplsteTailSessionDetail.MpDownloadState})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &ipv4bfDoMplsteTailSessionDetail.LspPingInfo})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range ipv4bfDoMplsteTailSessionDetail.OwnerInformation {
        ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteTailSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", ipv4bfDoMplsteTailSessionDetail.OwnerInformation[i]})
    }
    ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range ipv4bfDoMplsteTailSessionDetail.AssociationInformation {
        ipv4bfDoMplsteTailSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteTailSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", ipv4bfDoMplsteTailSessionDetail.AssociationInformation[i]})
    }
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteTailSessionDetail.VrfName})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteTailSessionDetail.IncomingLabel})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteTailSessionDetail.FeCtype})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteTailSessionDetail.FecSubgroupId})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteTailSessionDetail.Feclspid})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteTailSessionDetail.FecTunnelId})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteTailSessionDetail.FecExtendedTunnelId})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteTailSessionDetail.FecSource})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteTailSessionDetail.FecDestination})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteTailSessionDetail.Fecp2mpid})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteTailSessionDetail.FecSubgroupOriginator})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteTailSessionDetail.FecCtype})
    ipv4bfDoMplsteTailSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteTailSessionDetail.Location})

    ipv4bfDoMplsteTailSessionDetail.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailSessionDetail.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation
// Session status information
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation
// Association session information
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4bfDoMplsteTailSessionDetails_Ipv4bfDoMplsteTailSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4MultiHopNodeLocationSummaries
// Table of summary information about BFD IPv4
// multihop sessions per location
type Bfd_Ipv4MultiHopNodeLocationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for BFD IPv4 multihop sessions for location. The type
    // is slice of
    // Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary.
    Ipv4MultiHopNodeLocationSummary []*Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary
}

func (ipv4MultiHopNodeLocationSummaries *Bfd_Ipv4MultiHopNodeLocationSummaries) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopNodeLocationSummaries.EntityData.YFilter = ipv4MultiHopNodeLocationSummaries.YFilter
    ipv4MultiHopNodeLocationSummaries.EntityData.YangName = "ipv4-multi-hop-node-location-summaries"
    ipv4MultiHopNodeLocationSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopNodeLocationSummaries.EntityData.ParentYangName = "bfd"
    ipv4MultiHopNodeLocationSummaries.EntityData.SegmentPath = "ipv4-multi-hop-node-location-summaries"
    ipv4MultiHopNodeLocationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopNodeLocationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopNodeLocationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopNodeLocationSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopNodeLocationSummaries.EntityData.Children.Append("ipv4-multi-hop-node-location-summary", types.YChild{"Ipv4MultiHopNodeLocationSummary", nil})
    for i := range ipv4MultiHopNodeLocationSummaries.Ipv4MultiHopNodeLocationSummary {
        ipv4MultiHopNodeLocationSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopNodeLocationSummaries.Ipv4MultiHopNodeLocationSummary[i]), types.YChild{"Ipv4MultiHopNodeLocationSummary", ipv4MultiHopNodeLocationSummaries.Ipv4MultiHopNodeLocationSummary[i]})
    }
    ipv4MultiHopNodeLocationSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopNodeLocationSummaries.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopNodeLocationSummaries.EntityData)
}

// Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary
// Summary information for BFD IPv4 multihop
// sessions for location
type Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary_SessionState
}

func (ipv4MultiHopNodeLocationSummary *Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopNodeLocationSummary.EntityData.YFilter = ipv4MultiHopNodeLocationSummary.YFilter
    ipv4MultiHopNodeLocationSummary.EntityData.YangName = "ipv4-multi-hop-node-location-summary"
    ipv4MultiHopNodeLocationSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopNodeLocationSummary.EntityData.ParentYangName = "ipv4-multi-hop-node-location-summaries"
    ipv4MultiHopNodeLocationSummary.EntityData.SegmentPath = "ipv4-multi-hop-node-location-summary" + types.AddKeyToken(ipv4MultiHopNodeLocationSummary.Location, "location")
    ipv4MultiHopNodeLocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopNodeLocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopNodeLocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopNodeLocationSummary.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopNodeLocationSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4MultiHopNodeLocationSummary.SessionState})
    ipv4MultiHopNodeLocationSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4MultiHopNodeLocationSummary.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4MultiHopNodeLocationSummary.Location})

    ipv4MultiHopNodeLocationSummary.EntityData.YListKeys = []string {"Location"}

    return &(ipv4MultiHopNodeLocationSummary.EntityData)
}

// Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv4MultiHopNodeLocationSummaries_Ipv4MultiHopNodeLocationSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4-multi-hop-node-location-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionBriefs
// Table of brief information about all IPv4 BFD
// over MPLS-TE Tail sessions in the System
type Bfd_Ipv4bfDoMplsteTailSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv4 BFD over MPLS-TE session. The type is
    // slice of
    // Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief.
    Ipv4bfDoMplsteTailSessionBrief []*Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief
}

func (ipv4bfDoMplsteTailSessionBriefs *Bfd_Ipv4bfDoMplsteTailSessionBriefs) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailSessionBriefs.EntityData.YFilter = ipv4bfDoMplsteTailSessionBriefs.YFilter
    ipv4bfDoMplsteTailSessionBriefs.EntityData.YangName = "ipv4bf-do-mplste-tail-session-briefs"
    ipv4bfDoMplsteTailSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailSessionBriefs.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteTailSessionBriefs.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-session-briefs"
    ipv4bfDoMplsteTailSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailSessionBriefs.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailSessionBriefs.EntityData.Children.Append("ipv4bf-do-mplste-tail-session-brief", types.YChild{"Ipv4bfDoMplsteTailSessionBrief", nil})
    for i := range ipv4bfDoMplsteTailSessionBriefs.Ipv4bfDoMplsteTailSessionBrief {
        ipv4bfDoMplsteTailSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteTailSessionBriefs.Ipv4bfDoMplsteTailSessionBrief[i]), types.YChild{"Ipv4bfDoMplsteTailSessionBrief", ipv4bfDoMplsteTailSessionBriefs.Ipv4bfDoMplsteTailSessionBrief[i]})
    }
    ipv4bfDoMplsteTailSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteTailSessionBriefs.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailSessionBriefs.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief
// Brief information for a single IPv4 BFD over
// MPLS-TE session
type Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation
}

func (ipv4bfDoMplsteTailSessionBrief *Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailSessionBrief.EntityData.YFilter = ipv4bfDoMplsteTailSessionBrief.YFilter
    ipv4bfDoMplsteTailSessionBrief.EntityData.YangName = "ipv4bf-do-mplste-tail-session-brief"
    ipv4bfDoMplsteTailSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailSessionBrief.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-briefs"
    ipv4bfDoMplsteTailSessionBrief.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-session-brief"
    ipv4bfDoMplsteTailSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailSessionBrief.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &ipv4bfDoMplsteTailSessionBrief.StatusBriefInformation})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteTailSessionBrief.VrfName})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteTailSessionBrief.IncomingLabel})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteTailSessionBrief.FeCtype})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteTailSessionBrief.FecSubgroupId})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteTailSessionBrief.Feclspid})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteTailSessionBrief.FecTunnelId})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteTailSessionBrief.FecExtendedTunnelId})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteTailSessionBrief.FecSource})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteTailSessionBrief.FecDestination})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteTailSessionBrief.Fecp2mpid})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteTailSessionBrief.FecSubgroupOriginator})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteTailSessionBrief.FecCtype})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteTailSessionBrief.Location})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4bfDoMplsteTailSessionBrief.NodeId})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4bfDoMplsteTailSessionBrief.State})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", ipv4bfDoMplsteTailSessionBrief.SessionType})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4bfDoMplsteTailSessionBrief.SessionSubtype})
    ipv4bfDoMplsteTailSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", ipv4bfDoMplsteTailSessionBrief.SessionFlags})

    ipv4bfDoMplsteTailSessionBrief.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailSessionBrief.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4bfDoMplsteTailSessionBriefs_Ipv4bfDoMplsteTailSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv6MultiHopNodeLocationSummaries
// Table of summary information about BFD IPv6
// multihop sessions per location
type Bfd_Ipv6MultiHopNodeLocationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for BFD IPv6 multihop sessions for location. The type
    // is slice of
    // Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary.
    Ipv6MultiHopNodeLocationSummary []*Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary
}

func (ipv6MultiHopNodeLocationSummaries *Bfd_Ipv6MultiHopNodeLocationSummaries) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopNodeLocationSummaries.EntityData.YFilter = ipv6MultiHopNodeLocationSummaries.YFilter
    ipv6MultiHopNodeLocationSummaries.EntityData.YangName = "ipv6-multi-hop-node-location-summaries"
    ipv6MultiHopNodeLocationSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopNodeLocationSummaries.EntityData.ParentYangName = "bfd"
    ipv6MultiHopNodeLocationSummaries.EntityData.SegmentPath = "ipv6-multi-hop-node-location-summaries"
    ipv6MultiHopNodeLocationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopNodeLocationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopNodeLocationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopNodeLocationSummaries.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopNodeLocationSummaries.EntityData.Children.Append("ipv6-multi-hop-node-location-summary", types.YChild{"Ipv6MultiHopNodeLocationSummary", nil})
    for i := range ipv6MultiHopNodeLocationSummaries.Ipv6MultiHopNodeLocationSummary {
        ipv6MultiHopNodeLocationSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopNodeLocationSummaries.Ipv6MultiHopNodeLocationSummary[i]), types.YChild{"Ipv6MultiHopNodeLocationSummary", ipv6MultiHopNodeLocationSummaries.Ipv6MultiHopNodeLocationSummary[i]})
    }
    ipv6MultiHopNodeLocationSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopNodeLocationSummaries.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopNodeLocationSummaries.EntityData)
}

// Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary
// Summary information for BFD IPv6 multihop
// sessions for location
type Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary_SessionState
}

func (ipv6MultiHopNodeLocationSummary *Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopNodeLocationSummary.EntityData.YFilter = ipv6MultiHopNodeLocationSummary.YFilter
    ipv6MultiHopNodeLocationSummary.EntityData.YangName = "ipv6-multi-hop-node-location-summary"
    ipv6MultiHopNodeLocationSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopNodeLocationSummary.EntityData.ParentYangName = "ipv6-multi-hop-node-location-summaries"
    ipv6MultiHopNodeLocationSummary.EntityData.SegmentPath = "ipv6-multi-hop-node-location-summary" + types.AddKeyToken(ipv6MultiHopNodeLocationSummary.Location, "location")
    ipv6MultiHopNodeLocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopNodeLocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopNodeLocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopNodeLocationSummary.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopNodeLocationSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv6MultiHopNodeLocationSummary.SessionState})
    ipv6MultiHopNodeLocationSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv6MultiHopNodeLocationSummary.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6MultiHopNodeLocationSummary.Location})

    ipv6MultiHopNodeLocationSummary.EntityData.YListKeys = []string {"Location"}

    return &(ipv6MultiHopNodeLocationSummary.EntityData)
}

// Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv6MultiHopNodeLocationSummaries_Ipv6MultiHopNodeLocationSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv6-multi-hop-node-location-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4MultiHopSummary
// Summary information of BFD IPv4 multihop
// sessions
type Bfd_Ipv4MultiHopSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4MultiHopSummary_SessionState
}

func (ipv4MultiHopSummary *Bfd_Ipv4MultiHopSummary) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopSummary.EntityData.YFilter = ipv4MultiHopSummary.YFilter
    ipv4MultiHopSummary.EntityData.YangName = "ipv4-multi-hop-summary"
    ipv4MultiHopSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopSummary.EntityData.ParentYangName = "bfd"
    ipv4MultiHopSummary.EntityData.SegmentPath = "ipv4-multi-hop-summary"
    ipv4MultiHopSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopSummary.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4MultiHopSummary.SessionState})
    ipv4MultiHopSummary.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopSummary.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopSummary.EntityData)
}

// Bfd_Ipv4MultiHopSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4MultiHopSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Ipv4MultiHopSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4-multi-hop-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4SingleHopCounters
// IPv4 single hop Counters
type Bfd_Ipv4SingleHopCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv4 single hop Packet counters.
    Ipv4SingleHopPacketCounters Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters
}

func (ipv4SingleHopCounters *Bfd_Ipv4SingleHopCounters) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopCounters.EntityData.YFilter = ipv4SingleHopCounters.YFilter
    ipv4SingleHopCounters.EntityData.YangName = "ipv4-single-hop-counters"
    ipv4SingleHopCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopCounters.EntityData.ParentYangName = "bfd"
    ipv4SingleHopCounters.EntityData.SegmentPath = "ipv4-single-hop-counters"
    ipv4SingleHopCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopCounters.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopCounters.EntityData.Children.Append("ipv4-single-hop-packet-counters", types.YChild{"Ipv4SingleHopPacketCounters", &ipv4SingleHopCounters.Ipv4SingleHopPacketCounters})
    ipv4SingleHopCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopCounters.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopCounters.EntityData)
}

// Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters
// Table of IPv4 single hop Packet counters
type Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface IPv4 single hop Packet counters. The type is slice of
    // Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters_Ipv4SingleHopPacketCounter.
    Ipv4SingleHopPacketCounter []*Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters_Ipv4SingleHopPacketCounter
}

func (ipv4SingleHopPacketCounters *Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopPacketCounters.EntityData.YFilter = ipv4SingleHopPacketCounters.YFilter
    ipv4SingleHopPacketCounters.EntityData.YangName = "ipv4-single-hop-packet-counters"
    ipv4SingleHopPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopPacketCounters.EntityData.ParentYangName = "ipv4-single-hop-counters"
    ipv4SingleHopPacketCounters.EntityData.SegmentPath = "ipv4-single-hop-packet-counters"
    ipv4SingleHopPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopPacketCounters.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopPacketCounters.EntityData.Children.Append("ipv4-single-hop-packet-counter", types.YChild{"Ipv4SingleHopPacketCounter", nil})
    for i := range ipv4SingleHopPacketCounters.Ipv4SingleHopPacketCounter {
        ipv4SingleHopPacketCounters.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopPacketCounters.Ipv4SingleHopPacketCounter[i]), types.YChild{"Ipv4SingleHopPacketCounter", ipv4SingleHopPacketCounters.Ipv4SingleHopPacketCounter[i]})
    }
    ipv4SingleHopPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopPacketCounters.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopPacketCounters.EntityData)
}

// Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters_Ipv4SingleHopPacketCounter
// Interface IPv4 single hop Packet counters
type Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters_Ipv4SingleHopPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (ipv4SingleHopPacketCounter *Bfd_Ipv4SingleHopCounters_Ipv4SingleHopPacketCounters_Ipv4SingleHopPacketCounter) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopPacketCounter.EntityData.YFilter = ipv4SingleHopPacketCounter.YFilter
    ipv4SingleHopPacketCounter.EntityData.YangName = "ipv4-single-hop-packet-counter"
    ipv4SingleHopPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopPacketCounter.EntityData.ParentYangName = "ipv4-single-hop-packet-counters"
    ipv4SingleHopPacketCounter.EntityData.SegmentPath = "ipv4-single-hop-packet-counter"
    ipv4SingleHopPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopPacketCounter.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4SingleHopPacketCounter.InterfaceName})
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4SingleHopPacketCounter.Location})
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", ipv4SingleHopPacketCounter.HelloTransmitCount})
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", ipv4SingleHopPacketCounter.HelloReceiveCount})
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", ipv4SingleHopPacketCounter.EchoTransmitCount})
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", ipv4SingleHopPacketCounter.EchoReceiveCount})
    ipv4SingleHopPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", ipv4SingleHopPacketCounter.DisplayType})

    ipv4SingleHopPacketCounter.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopPacketCounter.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails
// Table of detailed information about all IPv6
// multihop BFD sessions in the System 
type Bfd_Ipv6MultiHopSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv6 multihop BFD session. The type is
    // slice of Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail.
    Ipv6MultiHopSessionDetail []*Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail
}

func (ipv6MultiHopSessionDetails *Bfd_Ipv6MultiHopSessionDetails) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopSessionDetails.EntityData.YFilter = ipv6MultiHopSessionDetails.YFilter
    ipv6MultiHopSessionDetails.EntityData.YangName = "ipv6-multi-hop-session-details"
    ipv6MultiHopSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopSessionDetails.EntityData.ParentYangName = "bfd"
    ipv6MultiHopSessionDetails.EntityData.SegmentPath = "ipv6-multi-hop-session-details"
    ipv6MultiHopSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopSessionDetails.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopSessionDetails.EntityData.Children.Append("ipv6-multi-hop-session-detail", types.YChild{"Ipv6MultiHopSessionDetail", nil})
    for i := range ipv6MultiHopSessionDetails.Ipv6MultiHopSessionDetail {
        ipv6MultiHopSessionDetails.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopSessionDetails.Ipv6MultiHopSessionDetail[i]), types.YChild{"Ipv6MultiHopSessionDetail", ipv6MultiHopSessionDetails.Ipv6MultiHopSessionDetail[i]})
    }
    ipv6MultiHopSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopSessionDetails.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopSessionDetails.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail
// Detailed information for a single IPv6 multihop
// BFD session
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Session status information.
    StatusInformation Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation
}

func (ipv6MultiHopSessionDetail *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopSessionDetail.EntityData.YFilter = ipv6MultiHopSessionDetail.YFilter
    ipv6MultiHopSessionDetail.EntityData.YangName = "ipv6-multi-hop-session-detail"
    ipv6MultiHopSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopSessionDetail.EntityData.ParentYangName = "ipv6-multi-hop-session-details"
    ipv6MultiHopSessionDetail.EntityData.SegmentPath = "ipv6-multi-hop-session-detail"
    ipv6MultiHopSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopSessionDetail.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &ipv6MultiHopSessionDetail.StatusInformation})
    ipv6MultiHopSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &ipv6MultiHopSessionDetail.MpDownloadState})
    ipv6MultiHopSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &ipv6MultiHopSessionDetail.LspPingInfo})
    ipv6MultiHopSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range ipv6MultiHopSessionDetail.OwnerInformation {
        ipv6MultiHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", ipv6MultiHopSessionDetail.OwnerInformation[i]})
    }
    ipv6MultiHopSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range ipv6MultiHopSessionDetail.AssociationInformation {
        ipv6MultiHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", ipv6MultiHopSessionDetail.AssociationInformation[i]})
    }
    ipv6MultiHopSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    ipv6MultiHopSessionDetail.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv6MultiHopSessionDetail.SourceAddress})
    ipv6MultiHopSessionDetail.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6MultiHopSessionDetail.DestinationAddress})
    ipv6MultiHopSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6MultiHopSessionDetail.Location})
    ipv6MultiHopSessionDetail.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6MultiHopSessionDetail.VrfName})

    ipv6MultiHopSessionDetail.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopSessionDetail.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation
// Session status information
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "ipv6-multi-hop-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "ipv6-multi-hop-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "ipv6-multi-hop-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "ipv6-multi-hop-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation
// Association session information
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "ipv6-multi-hop-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv6MultiHopSessionDetails_Ipv6MultiHopSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv6MultiHopMultiPaths
// IPv6 multi hop multipath
type Bfd_Ipv6MultiHopMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 multihop multipath table. The type is slice of
    // Bfd_Ipv6MultiHopMultiPaths_Ipv6MultiHopMultiPath.
    Ipv6MultiHopMultiPath []*Bfd_Ipv6MultiHopMultiPaths_Ipv6MultiHopMultiPath
}

func (ipv6MultiHopMultiPaths *Bfd_Ipv6MultiHopMultiPaths) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopMultiPaths.EntityData.YFilter = ipv6MultiHopMultiPaths.YFilter
    ipv6MultiHopMultiPaths.EntityData.YangName = "ipv6-multi-hop-multi-paths"
    ipv6MultiHopMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopMultiPaths.EntityData.ParentYangName = "bfd"
    ipv6MultiHopMultiPaths.EntityData.SegmentPath = "ipv6-multi-hop-multi-paths"
    ipv6MultiHopMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopMultiPaths.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopMultiPaths.EntityData.Children.Append("ipv6-multi-hop-multi-path", types.YChild{"Ipv6MultiHopMultiPath", nil})
    for i := range ipv6MultiHopMultiPaths.Ipv6MultiHopMultiPath {
        ipv6MultiHopMultiPaths.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopMultiPaths.Ipv6MultiHopMultiPath[i]), types.YChild{"Ipv6MultiHopMultiPath", ipv6MultiHopMultiPaths.Ipv6MultiHopMultiPath[i]})
    }
    ipv6MultiHopMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopMultiPaths.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopMultiPaths.EntityData)
}

// Bfd_Ipv6MultiHopMultiPaths_Ipv6MultiHopMultiPath
// IPv6 multihop multipath table
type Bfd_Ipv6MultiHopMultiPaths_Ipv6MultiHopMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (ipv6MultiHopMultiPath *Bfd_Ipv6MultiHopMultiPaths_Ipv6MultiHopMultiPath) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopMultiPath.EntityData.YFilter = ipv6MultiHopMultiPath.YFilter
    ipv6MultiHopMultiPath.EntityData.YangName = "ipv6-multi-hop-multi-path"
    ipv6MultiHopMultiPath.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopMultiPath.EntityData.ParentYangName = "ipv6-multi-hop-multi-paths"
    ipv6MultiHopMultiPath.EntityData.SegmentPath = "ipv6-multi-hop-multi-path"
    ipv6MultiHopMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopMultiPath.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopMultiPath.EntityData.Leafs = types.NewOrderedMap()
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv6MultiHopMultiPath.SourceAddress})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6MultiHopMultiPath.DestinationAddress})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6MultiHopMultiPath.Location})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6MultiHopMultiPath.VrfName})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv6MultiHopMultiPath.SessionSubtype})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv6MultiHopMultiPath.State})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", ipv6MultiHopMultiPath.LocalDiscriminator})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv6MultiHopMultiPath.NodeId})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", ipv6MultiHopMultiPath.IncomingLabelXr})
    ipv6MultiHopMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", ipv6MultiHopMultiPath.SessionInterfaceName})

    ipv6MultiHopMultiPath.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopMultiPath.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadCounters
// IPv4 BFD over MPLS-TE Counters
type Bfd_Ipv4bfDoMplsteHeadCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv4 BFD over MPLS-TE Packet counters.
    Ipv4bfDoMplsteHeadPacketCounters Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters
}

func (ipv4bfDoMplsteHeadCounters *Bfd_Ipv4bfDoMplsteHeadCounters) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadCounters.EntityData.YFilter = ipv4bfDoMplsteHeadCounters.YFilter
    ipv4bfDoMplsteHeadCounters.EntityData.YangName = "ipv4bf-do-mplste-head-counters"
    ipv4bfDoMplsteHeadCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadCounters.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteHeadCounters.EntityData.SegmentPath = "ipv4bf-do-mplste-head-counters"
    ipv4bfDoMplsteHeadCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadCounters.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadCounters.EntityData.Children.Append("ipv4bf-do-mplste-head-packet-counters", types.YChild{"Ipv4bfDoMplsteHeadPacketCounters", &ipv4bfDoMplsteHeadCounters.Ipv4bfDoMplsteHeadPacketCounters})
    ipv4bfDoMplsteHeadCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteHeadCounters.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadCounters.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters
// Table of IPv4 BFD over MPLS-TE Packet counters
type Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface  IPv4 BFD over MPLS-TE Packet counters. The type is slice of
    // Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters_Ipv4bfDoMplsteHeadPacketCounter.
    Ipv4bfDoMplsteHeadPacketCounter []*Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters_Ipv4bfDoMplsteHeadPacketCounter
}

func (ipv4bfDoMplsteHeadPacketCounters *Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadPacketCounters.EntityData.YFilter = ipv4bfDoMplsteHeadPacketCounters.YFilter
    ipv4bfDoMplsteHeadPacketCounters.EntityData.YangName = "ipv4bf-do-mplste-head-packet-counters"
    ipv4bfDoMplsteHeadPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadPacketCounters.EntityData.ParentYangName = "ipv4bf-do-mplste-head-counters"
    ipv4bfDoMplsteHeadPacketCounters.EntityData.SegmentPath = "ipv4bf-do-mplste-head-packet-counters"
    ipv4bfDoMplsteHeadPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadPacketCounters.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadPacketCounters.EntityData.Children.Append("ipv4bf-do-mplste-head-packet-counter", types.YChild{"Ipv4bfDoMplsteHeadPacketCounter", nil})
    for i := range ipv4bfDoMplsteHeadPacketCounters.Ipv4bfDoMplsteHeadPacketCounter {
        ipv4bfDoMplsteHeadPacketCounters.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteHeadPacketCounters.Ipv4bfDoMplsteHeadPacketCounter[i]), types.YChild{"Ipv4bfDoMplsteHeadPacketCounter", ipv4bfDoMplsteHeadPacketCounters.Ipv4bfDoMplsteHeadPacketCounter[i]})
    }
    ipv4bfDoMplsteHeadPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteHeadPacketCounters.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadPacketCounters.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters_Ipv4bfDoMplsteHeadPacketCounter
// Interface  IPv4 BFD over MPLS-TE Packet
// counters
type Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters_Ipv4bfDoMplsteHeadPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (ipv4bfDoMplsteHeadPacketCounter *Bfd_Ipv4bfDoMplsteHeadCounters_Ipv4bfDoMplsteHeadPacketCounters_Ipv4bfDoMplsteHeadPacketCounter) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadPacketCounter.EntityData.YFilter = ipv4bfDoMplsteHeadPacketCounter.YFilter
    ipv4bfDoMplsteHeadPacketCounter.EntityData.YangName = "ipv4bf-do-mplste-head-packet-counter"
    ipv4bfDoMplsteHeadPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadPacketCounter.EntityData.ParentYangName = "ipv4bf-do-mplste-head-packet-counters"
    ipv4bfDoMplsteHeadPacketCounter.EntityData.SegmentPath = "ipv4bf-do-mplste-head-packet-counter"
    ipv4bfDoMplsteHeadPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadPacketCounter.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4bfDoMplsteHeadPacketCounter.InterfaceName})
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteHeadPacketCounter.Location})
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", ipv4bfDoMplsteHeadPacketCounter.HelloTransmitCount})
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", ipv4bfDoMplsteHeadPacketCounter.HelloReceiveCount})
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", ipv4bfDoMplsteHeadPacketCounter.EchoTransmitCount})
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", ipv4bfDoMplsteHeadPacketCounter.EchoReceiveCount})
    ipv4bfDoMplsteHeadPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", ipv4bfDoMplsteHeadPacketCounter.DisplayType})

    ipv4bfDoMplsteHeadPacketCounter.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadPacketCounter.EntityData)
}

// Bfd_SessionMibs
// BFD session MIB database
type Bfd_SessionMibs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for BFD session MIB. The type is slice of
    // Bfd_SessionMibs_SessionMib.
    SessionMib []*Bfd_SessionMibs_SessionMib
}

func (sessionMibs *Bfd_SessionMibs) GetEntityData() *types.CommonEntityData {
    sessionMibs.EntityData.YFilter = sessionMibs.YFilter
    sessionMibs.EntityData.YangName = "session-mibs"
    sessionMibs.EntityData.BundleName = "cisco_ios_xr"
    sessionMibs.EntityData.ParentYangName = "bfd"
    sessionMibs.EntityData.SegmentPath = "session-mibs"
    sessionMibs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionMibs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionMibs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionMibs.EntityData.Children = types.NewOrderedMap()
    sessionMibs.EntityData.Children.Append("session-mib", types.YChild{"SessionMib", nil})
    for i := range sessionMibs.SessionMib {
        sessionMibs.EntityData.Children.Append(types.GetSegmentPath(sessionMibs.SessionMib[i]), types.YChild{"SessionMib", sessionMibs.SessionMib[i]})
    }
    sessionMibs.EntityData.Leafs = types.NewOrderedMap()

    sessionMibs.EntityData.YListKeys = []string {}

    return &(sessionMibs.EntityData)
}

// Bfd_SessionMibs_SessionMib
// Brief information for BFD session MIB
type Bfd_SessionMibs_SessionMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session Discr. The type is interface{} with range:
    // 1..4294967295.
    Discriminator interface{}

    // Sessions' Local Discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Sessions' Remote Discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Session BFD Version. The type is interface{} with range: 0..4294967295.
    Sessionversion interface{}

    // Session State. The type is interface{} with range: 0..4294967295.
    SessionState interface{}

    // Trap Generator Bitmap. The type is interface{} with range: 0..4294967295.
    TrapBitmap interface{}

    // Packet In Counter. The type is interface{} with range:
    // 0..18446744073709551615.
    PktIn interface{}

    // Packet Out Counter. The type is interface{} with range:
    // 0..18446744073709551615.
    PktOut interface{}

    // Last Session Up Time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    LastUpTimeSec interface{}

    // Last Session Up Time (nanoseconds). The type is interface{} with range:
    // 0..4294967295. Units are nanosecond.
    LastUpTimeNsec interface{}

    // Last Session Down Time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    LastDownTimeSec interface{}

    // Last Session Down Time (nanoseconds). The type is interface{} with range:
    // 0..4294967295. Units are nanosecond.
    LastDownTimeNsec interface{}

    // Last IO EVM Schd Time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    LastIoEvmSchdTimeSec interface{}

    // Last IO Evm Schd Time (nanoseconds). The type is interface{} with range:
    // 0..4294967295. Units are nanosecond.
    LastIoEvmSchdTimeNsec interface{}

    // Last IO EVM Schd Comp Time (seconds). The type is interface{} with range:
    // 0..18446744073709551615. Units are second.
    LastIoEvmSchdCompTimeSec interface{}

    // Last IO Evm Schd Comp Time (nanoseconds). The type is interface{} with
    // range: 0..4294967295. Units are nanosecond.
    LastIoEvmSchdCompTimeNsec interface{}

    // Last Session Down Diag. The type is BfdMgmtSessionDiag.
    LastDownDiag interface{}

    // Last Rx Session Down Diag. The type is BfdMgmtSessionDiag.
    LastRxDownDiag interface{}

    // Up Count. The type is interface{} with range: 0..4294967295.
    UpCounter interface{}

    // Last Time Session Info Queried. The type is interface{} with range:
    // 0..18446744073709551615.
    LastTimeCached interface{}

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session Interface Handle. The type is interface{} with range:
    // 0..4294967295.
    IntHandle interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Desired Min TX Interval. The type is interface{} with range: 0..4294967295.
    DesiredMinTxInterval interface{}

    // Required Min RX Interval. The type is interface{} with range:
    // 0..4294967295.
    RequiredMinRxInterval interface{}

    // Required Min RX Echo Interval. The type is interface{} with range:
    // 0..4294967295.
    RequiredMinRxEchoInterval interface{}

    // Session Destination address.
    DestAddress Bfd_SessionMibs_SessionMib_DestAddress
}

func (sessionMib *Bfd_SessionMibs_SessionMib) GetEntityData() *types.CommonEntityData {
    sessionMib.EntityData.YFilter = sessionMib.YFilter
    sessionMib.EntityData.YangName = "session-mib"
    sessionMib.EntityData.BundleName = "cisco_ios_xr"
    sessionMib.EntityData.ParentYangName = "session-mibs"
    sessionMib.EntityData.SegmentPath = "session-mib" + types.AddKeyToken(sessionMib.Discriminator, "discriminator")
    sessionMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionMib.EntityData.Children = types.NewOrderedMap()
    sessionMib.EntityData.Children.Append("dest-address", types.YChild{"DestAddress", &sessionMib.DestAddress})
    sessionMib.EntityData.Leafs = types.NewOrderedMap()
    sessionMib.EntityData.Leafs.Append("discriminator", types.YLeaf{"Discriminator", sessionMib.Discriminator})
    sessionMib.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", sessionMib.LocalDiscriminator})
    sessionMib.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", sessionMib.RemoteDiscriminator})
    sessionMib.EntityData.Leafs.Append("sessionversion", types.YLeaf{"Sessionversion", sessionMib.Sessionversion})
    sessionMib.EntityData.Leafs.Append("session-state", types.YLeaf{"SessionState", sessionMib.SessionState})
    sessionMib.EntityData.Leafs.Append("trap-bitmap", types.YLeaf{"TrapBitmap", sessionMib.TrapBitmap})
    sessionMib.EntityData.Leafs.Append("pkt-in", types.YLeaf{"PktIn", sessionMib.PktIn})
    sessionMib.EntityData.Leafs.Append("pkt-out", types.YLeaf{"PktOut", sessionMib.PktOut})
    sessionMib.EntityData.Leafs.Append("last-up-time-sec", types.YLeaf{"LastUpTimeSec", sessionMib.LastUpTimeSec})
    sessionMib.EntityData.Leafs.Append("last-up-time-nsec", types.YLeaf{"LastUpTimeNsec", sessionMib.LastUpTimeNsec})
    sessionMib.EntityData.Leafs.Append("last-down-time-sec", types.YLeaf{"LastDownTimeSec", sessionMib.LastDownTimeSec})
    sessionMib.EntityData.Leafs.Append("last-down-time-nsec", types.YLeaf{"LastDownTimeNsec", sessionMib.LastDownTimeNsec})
    sessionMib.EntityData.Leafs.Append("last-io-evm-schd-time-sec", types.YLeaf{"LastIoEvmSchdTimeSec", sessionMib.LastIoEvmSchdTimeSec})
    sessionMib.EntityData.Leafs.Append("last-io-evm-schd-time-nsec", types.YLeaf{"LastIoEvmSchdTimeNsec", sessionMib.LastIoEvmSchdTimeNsec})
    sessionMib.EntityData.Leafs.Append("last-io-evm-schd-comp-time-sec", types.YLeaf{"LastIoEvmSchdCompTimeSec", sessionMib.LastIoEvmSchdCompTimeSec})
    sessionMib.EntityData.Leafs.Append("last-io-evm-schd-comp-time-nsec", types.YLeaf{"LastIoEvmSchdCompTimeNsec", sessionMib.LastIoEvmSchdCompTimeNsec})
    sessionMib.EntityData.Leafs.Append("last-down-diag", types.YLeaf{"LastDownDiag", sessionMib.LastDownDiag})
    sessionMib.EntityData.Leafs.Append("last-rx-down-diag", types.YLeaf{"LastRxDownDiag", sessionMib.LastRxDownDiag})
    sessionMib.EntityData.Leafs.Append("up-counter", types.YLeaf{"UpCounter", sessionMib.UpCounter})
    sessionMib.EntityData.Leafs.Append("last-time-cached", types.YLeaf{"LastTimeCached", sessionMib.LastTimeCached})
    sessionMib.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", sessionMib.InterfaceName})
    sessionMib.EntityData.Leafs.Append("int-handle", types.YLeaf{"IntHandle", sessionMib.IntHandle})
    sessionMib.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", sessionMib.DetectionMultiplier})
    sessionMib.EntityData.Leafs.Append("desired-min-tx-interval", types.YLeaf{"DesiredMinTxInterval", sessionMib.DesiredMinTxInterval})
    sessionMib.EntityData.Leafs.Append("required-min-rx-interval", types.YLeaf{"RequiredMinRxInterval", sessionMib.RequiredMinRxInterval})
    sessionMib.EntityData.Leafs.Append("required-min-rx-echo-interval", types.YLeaf{"RequiredMinRxEchoInterval", sessionMib.RequiredMinRxEchoInterval})

    sessionMib.EntityData.YListKeys = []string {"Discriminator"}

    return &(sessionMib.EntityData)
}

// Bfd_SessionMibs_SessionMib_DestAddress
// Session Destination address
type Bfd_SessionMibs_SessionMib_DestAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (destAddress *Bfd_SessionMibs_SessionMib_DestAddress) GetEntityData() *types.CommonEntityData {
    destAddress.EntityData.YFilter = destAddress.YFilter
    destAddress.EntityData.YangName = "dest-address"
    destAddress.EntityData.BundleName = "cisco_ios_xr"
    destAddress.EntityData.ParentYangName = "session-mib"
    destAddress.EntityData.SegmentPath = "dest-address"
    destAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destAddress.EntityData.Children = types.NewOrderedMap()
    destAddress.EntityData.Leafs = types.NewOrderedMap()
    destAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", destAddress.Afi})
    destAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", destAddress.Dummy})
    destAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", destAddress.Ipv4})
    destAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", destAddress.Ipv6})

    destAddress.EntityData.YListKeys = []string {}

    return &(destAddress.EntityData)
}

// Bfd_Ipv6MultiHopSummary
// Summary information of BFD IPv6 multihop
// sessions
type Bfd_Ipv6MultiHopSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Ipv6MultiHopSummary_SessionState
}

func (ipv6MultiHopSummary *Bfd_Ipv6MultiHopSummary) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopSummary.EntityData.YFilter = ipv6MultiHopSummary.YFilter
    ipv6MultiHopSummary.EntityData.YangName = "ipv6-multi-hop-summary"
    ipv6MultiHopSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopSummary.EntityData.ParentYangName = "bfd"
    ipv6MultiHopSummary.EntityData.SegmentPath = "ipv6-multi-hop-summary"
    ipv6MultiHopSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopSummary.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv6MultiHopSummary.SessionState})
    ipv6MultiHopSummary.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopSummary.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopSummary.EntityData)
}

// Bfd_Ipv6MultiHopSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv6MultiHopSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Ipv6MultiHopSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv6-multi-hop-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_LabelSummaryNodes
// Table of summary about Label BFD sessions for
// location
type Bfd_LabelSummaryNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of Label BFD . The type is slice of
    // Bfd_LabelSummaryNodes_LabelSummaryNode.
    LabelSummaryNode []*Bfd_LabelSummaryNodes_LabelSummaryNode
}

func (labelSummaryNodes *Bfd_LabelSummaryNodes) GetEntityData() *types.CommonEntityData {
    labelSummaryNodes.EntityData.YFilter = labelSummaryNodes.YFilter
    labelSummaryNodes.EntityData.YangName = "label-summary-nodes"
    labelSummaryNodes.EntityData.BundleName = "cisco_ios_xr"
    labelSummaryNodes.EntityData.ParentYangName = "bfd"
    labelSummaryNodes.EntityData.SegmentPath = "label-summary-nodes"
    labelSummaryNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSummaryNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSummaryNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSummaryNodes.EntityData.Children = types.NewOrderedMap()
    labelSummaryNodes.EntityData.Children.Append("label-summary-node", types.YChild{"LabelSummaryNode", nil})
    for i := range labelSummaryNodes.LabelSummaryNode {
        labelSummaryNodes.EntityData.Children.Append(types.GetSegmentPath(labelSummaryNodes.LabelSummaryNode[i]), types.YChild{"LabelSummaryNode", labelSummaryNodes.LabelSummaryNode[i]})
    }
    labelSummaryNodes.EntityData.Leafs = types.NewOrderedMap()

    labelSummaryNodes.EntityData.YListKeys = []string {}

    return &(labelSummaryNodes.EntityData)
}

// Bfd_LabelSummaryNodes_LabelSummaryNode
// Summary of Label BFD 
type Bfd_LabelSummaryNodes_LabelSummaryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocationName interface{}

    // Statistics of states for sessions.
    SessionState Bfd_LabelSummaryNodes_LabelSummaryNode_SessionState
}

func (labelSummaryNode *Bfd_LabelSummaryNodes_LabelSummaryNode) GetEntityData() *types.CommonEntityData {
    labelSummaryNode.EntityData.YFilter = labelSummaryNode.YFilter
    labelSummaryNode.EntityData.YangName = "label-summary-node"
    labelSummaryNode.EntityData.BundleName = "cisco_ios_xr"
    labelSummaryNode.EntityData.ParentYangName = "label-summary-nodes"
    labelSummaryNode.EntityData.SegmentPath = "label-summary-node" + types.AddKeyToken(labelSummaryNode.LocationName, "location-name")
    labelSummaryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSummaryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSummaryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSummaryNode.EntityData.Children = types.NewOrderedMap()
    labelSummaryNode.EntityData.Children.Append("session-state", types.YChild{"SessionState", &labelSummaryNode.SessionState})
    labelSummaryNode.EntityData.Leafs = types.NewOrderedMap()
    labelSummaryNode.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", labelSummaryNode.LocationName})

    labelSummaryNode.EntityData.YListKeys = []string {"LocationName"}

    return &(labelSummaryNode.EntityData)
}

// Bfd_LabelSummaryNodes_LabelSummaryNode_SessionState
// Statistics of states for sessions
type Bfd_LabelSummaryNodes_LabelSummaryNode_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_LabelSummaryNodes_LabelSummaryNode_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "label-summary-node"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv6MultiHopSessionBriefs
// Table of brief information about all IPv6
// multihop BFD sessions in the System
type Bfd_Ipv6MultiHopSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv6 multihop BFD session. The type is slice
    // of Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief.
    Ipv6MultiHopSessionBrief []*Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief
}

func (ipv6MultiHopSessionBriefs *Bfd_Ipv6MultiHopSessionBriefs) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopSessionBriefs.EntityData.YFilter = ipv6MultiHopSessionBriefs.YFilter
    ipv6MultiHopSessionBriefs.EntityData.YangName = "ipv6-multi-hop-session-briefs"
    ipv6MultiHopSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopSessionBriefs.EntityData.ParentYangName = "bfd"
    ipv6MultiHopSessionBriefs.EntityData.SegmentPath = "ipv6-multi-hop-session-briefs"
    ipv6MultiHopSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopSessionBriefs.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopSessionBriefs.EntityData.Children.Append("ipv6-multi-hop-session-brief", types.YChild{"Ipv6MultiHopSessionBrief", nil})
    for i := range ipv6MultiHopSessionBriefs.Ipv6MultiHopSessionBrief {
        ipv6MultiHopSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopSessionBriefs.Ipv6MultiHopSessionBrief[i]), types.YChild{"Ipv6MultiHopSessionBrief", ipv6MultiHopSessionBriefs.Ipv6MultiHopSessionBrief[i]})
    }
    ipv6MultiHopSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopSessionBriefs.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopSessionBriefs.EntityData)
}

// Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief
// Brief information for a single IPv6 multihop
// BFD session
type Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation
}

func (ipv6MultiHopSessionBrief *Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopSessionBrief.EntityData.YFilter = ipv6MultiHopSessionBrief.YFilter
    ipv6MultiHopSessionBrief.EntityData.YangName = "ipv6-multi-hop-session-brief"
    ipv6MultiHopSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopSessionBrief.EntityData.ParentYangName = "ipv6-multi-hop-session-briefs"
    ipv6MultiHopSessionBrief.EntityData.SegmentPath = "ipv6-multi-hop-session-brief"
    ipv6MultiHopSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopSessionBrief.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &ipv6MultiHopSessionBrief.StatusBriefInformation})
    ipv6MultiHopSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv6MultiHopSessionBrief.SourceAddress})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6MultiHopSessionBrief.DestinationAddress})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6MultiHopSessionBrief.Location})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6MultiHopSessionBrief.VrfName})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv6MultiHopSessionBrief.NodeId})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv6MultiHopSessionBrief.State})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", ipv6MultiHopSessionBrief.SessionType})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv6MultiHopSessionBrief.SessionSubtype})
    ipv6MultiHopSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", ipv6MultiHopSessionBrief.SessionFlags})

    ipv6MultiHopSessionBrief.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopSessionBrief.EntityData)
}

// Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "ipv6-multi-hop-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv6MultiHopSessionBriefs_Ipv6MultiHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_SessionBriefs
// Table of brief information about singlehop IPv4
// BFD sessions in the System
type Bfd_SessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv4 singlehop BFD session. The type is
    // slice of Bfd_SessionBriefs_SessionBrief.
    SessionBrief []*Bfd_SessionBriefs_SessionBrief
}

func (sessionBriefs *Bfd_SessionBriefs) GetEntityData() *types.CommonEntityData {
    sessionBriefs.EntityData.YFilter = sessionBriefs.YFilter
    sessionBriefs.EntityData.YangName = "session-briefs"
    sessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    sessionBriefs.EntityData.ParentYangName = "bfd"
    sessionBriefs.EntityData.SegmentPath = "session-briefs"
    sessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionBriefs.EntityData.Children = types.NewOrderedMap()
    sessionBriefs.EntityData.Children.Append("session-brief", types.YChild{"SessionBrief", nil})
    for i := range sessionBriefs.SessionBrief {
        sessionBriefs.EntityData.Children.Append(types.GetSegmentPath(sessionBriefs.SessionBrief[i]), types.YChild{"SessionBrief", sessionBriefs.SessionBrief[i]})
    }
    sessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    sessionBriefs.EntityData.YListKeys = []string {}

    return &(sessionBriefs.EntityData)
}

// Bfd_SessionBriefs_SessionBrief
// Brief information for a single IPv4 singlehop
// BFD session
type Bfd_SessionBriefs_SessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_SessionBriefs_SessionBrief_StatusBriefInformation
}

func (sessionBrief *Bfd_SessionBriefs_SessionBrief) GetEntityData() *types.CommonEntityData {
    sessionBrief.EntityData.YFilter = sessionBrief.YFilter
    sessionBrief.EntityData.YangName = "session-brief"
    sessionBrief.EntityData.BundleName = "cisco_ios_xr"
    sessionBrief.EntityData.ParentYangName = "session-briefs"
    sessionBrief.EntityData.SegmentPath = "session-brief"
    sessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionBrief.EntityData.Children = types.NewOrderedMap()
    sessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &sessionBrief.StatusBriefInformation})
    sessionBrief.EntityData.Leafs = types.NewOrderedMap()
    sessionBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", sessionBrief.InterfaceName})
    sessionBrief.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", sessionBrief.DestinationAddress})
    sessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", sessionBrief.Location})
    sessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", sessionBrief.NodeId})
    sessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", sessionBrief.State})
    sessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", sessionBrief.SessionType})
    sessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", sessionBrief.SessionSubtype})
    sessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", sessionBrief.SessionFlags})

    sessionBrief.EntityData.YListKeys = []string {}

    return &(sessionBrief.EntityData)
}

// Bfd_SessionBriefs_SessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_SessionBriefs_SessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_SessionBriefs_SessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_SessionBriefs_SessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv6SingleHopNodeLocationSummaries
// Table of summary information about BFD IPv6
// singlehop sessions per location
type Bfd_Ipv6SingleHopNodeLocationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for BFD IPv6 singlehop sessions for location. The type
    // is slice of
    // Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary.
    Ipv6SingleHopNodeLocationSummary []*Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary
}

func (ipv6SingleHopNodeLocationSummaries *Bfd_Ipv6SingleHopNodeLocationSummaries) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopNodeLocationSummaries.EntityData.YFilter = ipv6SingleHopNodeLocationSummaries.YFilter
    ipv6SingleHopNodeLocationSummaries.EntityData.YangName = "ipv6-single-hop-node-location-summaries"
    ipv6SingleHopNodeLocationSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopNodeLocationSummaries.EntityData.ParentYangName = "bfd"
    ipv6SingleHopNodeLocationSummaries.EntityData.SegmentPath = "ipv6-single-hop-node-location-summaries"
    ipv6SingleHopNodeLocationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopNodeLocationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopNodeLocationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopNodeLocationSummaries.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopNodeLocationSummaries.EntityData.Children.Append("ipv6-single-hop-node-location-summary", types.YChild{"Ipv6SingleHopNodeLocationSummary", nil})
    for i := range ipv6SingleHopNodeLocationSummaries.Ipv6SingleHopNodeLocationSummary {
        ipv6SingleHopNodeLocationSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopNodeLocationSummaries.Ipv6SingleHopNodeLocationSummary[i]), types.YChild{"Ipv6SingleHopNodeLocationSummary", ipv6SingleHopNodeLocationSummaries.Ipv6SingleHopNodeLocationSummary[i]})
    }
    ipv6SingleHopNodeLocationSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopNodeLocationSummaries.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopNodeLocationSummaries.EntityData)
}

// Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary
// Summary information for BFD IPv6 singlehop
// sessions for location
type Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary_SessionState
}

func (ipv6SingleHopNodeLocationSummary *Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopNodeLocationSummary.EntityData.YFilter = ipv6SingleHopNodeLocationSummary.YFilter
    ipv6SingleHopNodeLocationSummary.EntityData.YangName = "ipv6-single-hop-node-location-summary"
    ipv6SingleHopNodeLocationSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopNodeLocationSummary.EntityData.ParentYangName = "ipv6-single-hop-node-location-summaries"
    ipv6SingleHopNodeLocationSummary.EntityData.SegmentPath = "ipv6-single-hop-node-location-summary" + types.AddKeyToken(ipv6SingleHopNodeLocationSummary.Location, "location")
    ipv6SingleHopNodeLocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopNodeLocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopNodeLocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopNodeLocationSummary.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopNodeLocationSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv6SingleHopNodeLocationSummary.SessionState})
    ipv6SingleHopNodeLocationSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv6SingleHopNodeLocationSummary.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6SingleHopNodeLocationSummary.Location})

    ipv6SingleHopNodeLocationSummary.EntityData.YListKeys = []string {"Location"}

    return &(ipv6SingleHopNodeLocationSummary.EntityData)
}

// Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv6SingleHopNodeLocationSummaries_Ipv6SingleHopNodeLocationSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv6-single-hop-node-location-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Summary
// Summary information of BFD IPv4 singlehop
// sessions
type Bfd_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Summary_SessionState
}

func (summary *Bfd_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "bfd"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &summary.SessionState})
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Bfd_Summary_SessionState
// Statistics of states for sessions
type Bfd_Summary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Summary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4bfdMplsteTailNodeSummaries
// Table of summary about IPv4 TE tail BFD sessions
// for location
type Bfd_Ipv4bfdMplsteTailNodeSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of IPv4 BFD over MPLS-TE tail. The type is slice of
    // Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary.
    Ipv4bfdMplsteTailNodeSummary []*Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary
}

func (ipv4bfdMplsteTailNodeSummaries *Bfd_Ipv4bfdMplsteTailNodeSummaries) GetEntityData() *types.CommonEntityData {
    ipv4bfdMplsteTailNodeSummaries.EntityData.YFilter = ipv4bfdMplsteTailNodeSummaries.YFilter
    ipv4bfdMplsteTailNodeSummaries.EntityData.YangName = "ipv4bfd-mplste-tail-node-summaries"
    ipv4bfdMplsteTailNodeSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfdMplsteTailNodeSummaries.EntityData.ParentYangName = "bfd"
    ipv4bfdMplsteTailNodeSummaries.EntityData.SegmentPath = "ipv4bfd-mplste-tail-node-summaries"
    ipv4bfdMplsteTailNodeSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfdMplsteTailNodeSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfdMplsteTailNodeSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfdMplsteTailNodeSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4bfdMplsteTailNodeSummaries.EntityData.Children.Append("ipv4bfd-mplste-tail-node-summary", types.YChild{"Ipv4bfdMplsteTailNodeSummary", nil})
    for i := range ipv4bfdMplsteTailNodeSummaries.Ipv4bfdMplsteTailNodeSummary {
        ipv4bfdMplsteTailNodeSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4bfdMplsteTailNodeSummaries.Ipv4bfdMplsteTailNodeSummary[i]), types.YChild{"Ipv4bfdMplsteTailNodeSummary", ipv4bfdMplsteTailNodeSummaries.Ipv4bfdMplsteTailNodeSummary[i]})
    }
    ipv4bfdMplsteTailNodeSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfdMplsteTailNodeSummaries.EntityData.YListKeys = []string {}

    return &(ipv4bfdMplsteTailNodeSummaries.EntityData)
}

// Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary
// Summary of IPv4 BFD over MPLS-TE tail
type Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocationName interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary_SessionState
}

func (ipv4bfdMplsteTailNodeSummary *Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary) GetEntityData() *types.CommonEntityData {
    ipv4bfdMplsteTailNodeSummary.EntityData.YFilter = ipv4bfdMplsteTailNodeSummary.YFilter
    ipv4bfdMplsteTailNodeSummary.EntityData.YangName = "ipv4bfd-mplste-tail-node-summary"
    ipv4bfdMplsteTailNodeSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfdMplsteTailNodeSummary.EntityData.ParentYangName = "ipv4bfd-mplste-tail-node-summaries"
    ipv4bfdMplsteTailNodeSummary.EntityData.SegmentPath = "ipv4bfd-mplste-tail-node-summary" + types.AddKeyToken(ipv4bfdMplsteTailNodeSummary.LocationName, "location-name")
    ipv4bfdMplsteTailNodeSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfdMplsteTailNodeSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfdMplsteTailNodeSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfdMplsteTailNodeSummary.EntityData.Children = types.NewOrderedMap()
    ipv4bfdMplsteTailNodeSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4bfdMplsteTailNodeSummary.SessionState})
    ipv4bfdMplsteTailNodeSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfdMplsteTailNodeSummary.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", ipv4bfdMplsteTailNodeSummary.LocationName})

    ipv4bfdMplsteTailNodeSummary.EntityData.YListKeys = []string {"LocationName"}

    return &(ipv4bfdMplsteTailNodeSummary.EntityData)
}

// Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv4bfdMplsteTailNodeSummaries_Ipv4bfdMplsteTailNodeSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4bfd-mplste-tail-node-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4SingleHopLocationSummaries
// Table of summary information about IPv4
// singlehop BFD sessions for location
type Bfd_Ipv4SingleHopLocationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for BFD IPv4 singlehop sessions for location. The type
    // is slice of
    // Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary.
    Ipv4SingleHopLocationSummary []*Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary
}

func (ipv4SingleHopLocationSummaries *Bfd_Ipv4SingleHopLocationSummaries) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopLocationSummaries.EntityData.YFilter = ipv4SingleHopLocationSummaries.YFilter
    ipv4SingleHopLocationSummaries.EntityData.YangName = "ipv4-single-hop-location-summaries"
    ipv4SingleHopLocationSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopLocationSummaries.EntityData.ParentYangName = "bfd"
    ipv4SingleHopLocationSummaries.EntityData.SegmentPath = "ipv4-single-hop-location-summaries"
    ipv4SingleHopLocationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopLocationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopLocationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopLocationSummaries.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopLocationSummaries.EntityData.Children.Append("ipv4-single-hop-location-summary", types.YChild{"Ipv4SingleHopLocationSummary", nil})
    for i := range ipv4SingleHopLocationSummaries.Ipv4SingleHopLocationSummary {
        ipv4SingleHopLocationSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopLocationSummaries.Ipv4SingleHopLocationSummary[i]), types.YChild{"Ipv4SingleHopLocationSummary", ipv4SingleHopLocationSummaries.Ipv4SingleHopLocationSummary[i]})
    }
    ipv4SingleHopLocationSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopLocationSummaries.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopLocationSummaries.EntityData)
}

// Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary
// Summary information for BFD IPv4 singlehop
// sessions for location
type Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocationName interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary_SessionState
}

func (ipv4SingleHopLocationSummary *Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopLocationSummary.EntityData.YFilter = ipv4SingleHopLocationSummary.YFilter
    ipv4SingleHopLocationSummary.EntityData.YangName = "ipv4-single-hop-location-summary"
    ipv4SingleHopLocationSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopLocationSummary.EntityData.ParentYangName = "ipv4-single-hop-location-summaries"
    ipv4SingleHopLocationSummary.EntityData.SegmentPath = "ipv4-single-hop-location-summary" + types.AddKeyToken(ipv4SingleHopLocationSummary.LocationName, "location-name")
    ipv4SingleHopLocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopLocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopLocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopLocationSummary.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopLocationSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4SingleHopLocationSummary.SessionState})
    ipv4SingleHopLocationSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv4SingleHopLocationSummary.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", ipv4SingleHopLocationSummary.LocationName})

    ipv4SingleHopLocationSummary.EntityData.YListKeys = []string {"LocationName"}

    return &(ipv4SingleHopLocationSummary.EntityData)
}

// Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv4SingleHopLocationSummaries_Ipv4SingleHopLocationSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4-single-hop-location-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_Ipv4bfdMplsteHeadSummaryNodes
// Table of summary about IPv4 TE head BFD sessions
// for location
type Bfd_Ipv4bfdMplsteHeadSummaryNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of IPv4 BFD over MPLS-TE head. The type is slice of
    // Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode.
    Ipv4bfdMplsteHeadSummaryNode []*Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode
}

func (ipv4bfdMplsteHeadSummaryNodes *Bfd_Ipv4bfdMplsteHeadSummaryNodes) GetEntityData() *types.CommonEntityData {
    ipv4bfdMplsteHeadSummaryNodes.EntityData.YFilter = ipv4bfdMplsteHeadSummaryNodes.YFilter
    ipv4bfdMplsteHeadSummaryNodes.EntityData.YangName = "ipv4bfd-mplste-head-summary-nodes"
    ipv4bfdMplsteHeadSummaryNodes.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfdMplsteHeadSummaryNodes.EntityData.ParentYangName = "bfd"
    ipv4bfdMplsteHeadSummaryNodes.EntityData.SegmentPath = "ipv4bfd-mplste-head-summary-nodes"
    ipv4bfdMplsteHeadSummaryNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfdMplsteHeadSummaryNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfdMplsteHeadSummaryNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfdMplsteHeadSummaryNodes.EntityData.Children = types.NewOrderedMap()
    ipv4bfdMplsteHeadSummaryNodes.EntityData.Children.Append("ipv4bfd-mplste-head-summary-node", types.YChild{"Ipv4bfdMplsteHeadSummaryNode", nil})
    for i := range ipv4bfdMplsteHeadSummaryNodes.Ipv4bfdMplsteHeadSummaryNode {
        ipv4bfdMplsteHeadSummaryNodes.EntityData.Children.Append(types.GetSegmentPath(ipv4bfdMplsteHeadSummaryNodes.Ipv4bfdMplsteHeadSummaryNode[i]), types.YChild{"Ipv4bfdMplsteHeadSummaryNode", ipv4bfdMplsteHeadSummaryNodes.Ipv4bfdMplsteHeadSummaryNode[i]})
    }
    ipv4bfdMplsteHeadSummaryNodes.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfdMplsteHeadSummaryNodes.EntityData.YListKeys = []string {}

    return &(ipv4bfdMplsteHeadSummaryNodes.EntityData)
}

// Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode
// Summary of IPv4 BFD over MPLS-TE head
type Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocationName interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode_SessionState
}

func (ipv4bfdMplsteHeadSummaryNode *Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode) GetEntityData() *types.CommonEntityData {
    ipv4bfdMplsteHeadSummaryNode.EntityData.YFilter = ipv4bfdMplsteHeadSummaryNode.YFilter
    ipv4bfdMplsteHeadSummaryNode.EntityData.YangName = "ipv4bfd-mplste-head-summary-node"
    ipv4bfdMplsteHeadSummaryNode.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfdMplsteHeadSummaryNode.EntityData.ParentYangName = "ipv4bfd-mplste-head-summary-nodes"
    ipv4bfdMplsteHeadSummaryNode.EntityData.SegmentPath = "ipv4bfd-mplste-head-summary-node" + types.AddKeyToken(ipv4bfdMplsteHeadSummaryNode.LocationName, "location-name")
    ipv4bfdMplsteHeadSummaryNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfdMplsteHeadSummaryNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfdMplsteHeadSummaryNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfdMplsteHeadSummaryNode.EntityData.Children = types.NewOrderedMap()
    ipv4bfdMplsteHeadSummaryNode.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4bfdMplsteHeadSummaryNode.SessionState})
    ipv4bfdMplsteHeadSummaryNode.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfdMplsteHeadSummaryNode.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", ipv4bfdMplsteHeadSummaryNode.LocationName})

    ipv4bfdMplsteHeadSummaryNode.EntityData.YListKeys = []string {"LocationName"}

    return &(ipv4bfdMplsteHeadSummaryNode.EntityData)
}

// Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode_SessionState
// Statistics of states for sessions
type Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv4bfdMplsteHeadSummaryNodes_Ipv4bfdMplsteHeadSummaryNode_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4bfd-mplste-head-summary-node"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_LabelSessionDetails
// Table of detailed information about all Label
// BFD sessions in the System 
type Bfd_LabelSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single BFD session. The type is slice of
    // Bfd_LabelSessionDetails_LabelSessionDetail.
    LabelSessionDetail []*Bfd_LabelSessionDetails_LabelSessionDetail
}

func (labelSessionDetails *Bfd_LabelSessionDetails) GetEntityData() *types.CommonEntityData {
    labelSessionDetails.EntityData.YFilter = labelSessionDetails.YFilter
    labelSessionDetails.EntityData.YangName = "label-session-details"
    labelSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    labelSessionDetails.EntityData.ParentYangName = "bfd"
    labelSessionDetails.EntityData.SegmentPath = "label-session-details"
    labelSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSessionDetails.EntityData.Children = types.NewOrderedMap()
    labelSessionDetails.EntityData.Children.Append("label-session-detail", types.YChild{"LabelSessionDetail", nil})
    for i := range labelSessionDetails.LabelSessionDetail {
        labelSessionDetails.EntityData.Children.Append(types.GetSegmentPath(labelSessionDetails.LabelSessionDetail[i]), types.YChild{"LabelSessionDetail", labelSessionDetails.LabelSessionDetail[i]})
    }
    labelSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    labelSessionDetails.EntityData.YListKeys = []string {}

    return &(labelSessionDetails.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail
// Detailed information for a single BFD session
type Bfd_LabelSessionDetails_LabelSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session status information.
    StatusInformation Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_LabelSessionDetails_LabelSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_LabelSessionDetails_LabelSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation
}

func (labelSessionDetail *Bfd_LabelSessionDetails_LabelSessionDetail) GetEntityData() *types.CommonEntityData {
    labelSessionDetail.EntityData.YFilter = labelSessionDetail.YFilter
    labelSessionDetail.EntityData.YangName = "label-session-detail"
    labelSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    labelSessionDetail.EntityData.ParentYangName = "label-session-details"
    labelSessionDetail.EntityData.SegmentPath = "label-session-detail"
    labelSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelSessionDetail.EntityData.Children = types.NewOrderedMap()
    labelSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &labelSessionDetail.StatusInformation})
    labelSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &labelSessionDetail.MpDownloadState})
    labelSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &labelSessionDetail.LspPingInfo})
    labelSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range labelSessionDetail.OwnerInformation {
        labelSessionDetail.EntityData.Children.Append(types.GetSegmentPath(labelSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", labelSessionDetail.OwnerInformation[i]})
    }
    labelSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range labelSessionDetail.AssociationInformation {
        labelSessionDetail.EntityData.Children.Append(types.GetSegmentPath(labelSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", labelSessionDetail.AssociationInformation[i]})
    }
    labelSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    labelSessionDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", labelSessionDetail.InterfaceName})
    labelSessionDetail.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", labelSessionDetail.IncomingLabel})
    labelSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", labelSessionDetail.Location})

    labelSessionDetail.EntityData.YListKeys = []string {}

    return &(labelSessionDetail.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation
// Session status information
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "label-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_LabelSessionDetails_LabelSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "label-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_LabelSessionDetails_LabelSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "label-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_LabelSessionDetails_LabelSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_LabelSessionDetails_LabelSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_LabelSessionDetails_LabelSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "label-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation
// Association session information
type Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "label-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_LabelSessionDetails_LabelSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails
// Table of detailed information about all IPv6
// singlehop BFD sessions in the System 
type Bfd_Ipv6SingleHopSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv6 singlehop BFD session. The type is
    // slice of Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail.
    Ipv6SingleHopSessionDetail []*Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail
}

func (ipv6SingleHopSessionDetails *Bfd_Ipv6SingleHopSessionDetails) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopSessionDetails.EntityData.YFilter = ipv6SingleHopSessionDetails.YFilter
    ipv6SingleHopSessionDetails.EntityData.YangName = "ipv6-single-hop-session-details"
    ipv6SingleHopSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopSessionDetails.EntityData.ParentYangName = "bfd"
    ipv6SingleHopSessionDetails.EntityData.SegmentPath = "ipv6-single-hop-session-details"
    ipv6SingleHopSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopSessionDetails.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopSessionDetails.EntityData.Children.Append("ipv6-single-hop-session-detail", types.YChild{"Ipv6SingleHopSessionDetail", nil})
    for i := range ipv6SingleHopSessionDetails.Ipv6SingleHopSessionDetail {
        ipv6SingleHopSessionDetails.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopSessionDetails.Ipv6SingleHopSessionDetail[i]), types.YChild{"Ipv6SingleHopSessionDetail", ipv6SingleHopSessionDetails.Ipv6SingleHopSessionDetail[i]})
    }
    ipv6SingleHopSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopSessionDetails.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopSessionDetails.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail
// Detailed information for a single IPv6
// singlehop BFD session
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session status information.
    StatusInformation Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation
}

func (ipv6SingleHopSessionDetail *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopSessionDetail.EntityData.YFilter = ipv6SingleHopSessionDetail.YFilter
    ipv6SingleHopSessionDetail.EntityData.YangName = "ipv6-single-hop-session-detail"
    ipv6SingleHopSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopSessionDetail.EntityData.ParentYangName = "ipv6-single-hop-session-details"
    ipv6SingleHopSessionDetail.EntityData.SegmentPath = "ipv6-single-hop-session-detail"
    ipv6SingleHopSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopSessionDetail.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &ipv6SingleHopSessionDetail.StatusInformation})
    ipv6SingleHopSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &ipv6SingleHopSessionDetail.MpDownloadState})
    ipv6SingleHopSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &ipv6SingleHopSessionDetail.LspPingInfo})
    ipv6SingleHopSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range ipv6SingleHopSessionDetail.OwnerInformation {
        ipv6SingleHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", ipv6SingleHopSessionDetail.OwnerInformation[i]})
    }
    ipv6SingleHopSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range ipv6SingleHopSessionDetail.AssociationInformation {
        ipv6SingleHopSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", ipv6SingleHopSessionDetail.AssociationInformation[i]})
    }
    ipv6SingleHopSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    ipv6SingleHopSessionDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv6SingleHopSessionDetail.InterfaceName})
    ipv6SingleHopSessionDetail.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6SingleHopSessionDetail.DestinationAddress})
    ipv6SingleHopSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6SingleHopSessionDetail.Location})

    ipv6SingleHopSessionDetail.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopSessionDetail.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation
// Session status information
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "ipv6-single-hop-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "ipv6-single-hop-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "ipv6-single-hop-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "ipv6-single-hop-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation
// Association session information
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "ipv6-single-hop-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv6SingleHopSessionDetails_Ipv6SingleHopSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4MultiHopCounters
// IPv4 multiple hop Counters
type Bfd_Ipv4MultiHopCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv4 multiple hop Packet counters.
    Ipv4MultiHopPacketCounters Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters
}

func (ipv4MultiHopCounters *Bfd_Ipv4MultiHopCounters) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopCounters.EntityData.YFilter = ipv4MultiHopCounters.YFilter
    ipv4MultiHopCounters.EntityData.YangName = "ipv4-multi-hop-counters"
    ipv4MultiHopCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopCounters.EntityData.ParentYangName = "bfd"
    ipv4MultiHopCounters.EntityData.SegmentPath = "ipv4-multi-hop-counters"
    ipv4MultiHopCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopCounters.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopCounters.EntityData.Children.Append("ipv4-multi-hop-packet-counters", types.YChild{"Ipv4MultiHopPacketCounters", &ipv4MultiHopCounters.Ipv4MultiHopPacketCounters})
    ipv4MultiHopCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopCounters.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopCounters.EntityData)
}

// Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters
// Table of IPv4 multiple hop Packet counters
type Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 multiple hop Packet counters. The type is slice of
    // Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters_Ipv4MultiHopPacketCounter.
    Ipv4MultiHopPacketCounter []*Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters_Ipv4MultiHopPacketCounter
}

func (ipv4MultiHopPacketCounters *Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopPacketCounters.EntityData.YFilter = ipv4MultiHopPacketCounters.YFilter
    ipv4MultiHopPacketCounters.EntityData.YangName = "ipv4-multi-hop-packet-counters"
    ipv4MultiHopPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopPacketCounters.EntityData.ParentYangName = "ipv4-multi-hop-counters"
    ipv4MultiHopPacketCounters.EntityData.SegmentPath = "ipv4-multi-hop-packet-counters"
    ipv4MultiHopPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopPacketCounters.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopPacketCounters.EntityData.Children.Append("ipv4-multi-hop-packet-counter", types.YChild{"Ipv4MultiHopPacketCounter", nil})
    for i := range ipv4MultiHopPacketCounters.Ipv4MultiHopPacketCounter {
        ipv4MultiHopPacketCounters.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopPacketCounters.Ipv4MultiHopPacketCounter[i]), types.YChild{"Ipv4MultiHopPacketCounter", ipv4MultiHopPacketCounters.Ipv4MultiHopPacketCounter[i]})
    }
    ipv4MultiHopPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopPacketCounters.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopPacketCounters.EntityData)
}

// Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters_Ipv4MultiHopPacketCounter
// IPv4 multiple hop Packet counters
type Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters_Ipv4MultiHopPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (ipv4MultiHopPacketCounter *Bfd_Ipv4MultiHopCounters_Ipv4MultiHopPacketCounters_Ipv4MultiHopPacketCounter) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopPacketCounter.EntityData.YFilter = ipv4MultiHopPacketCounter.YFilter
    ipv4MultiHopPacketCounter.EntityData.YangName = "ipv4-multi-hop-packet-counter"
    ipv4MultiHopPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopPacketCounter.EntityData.ParentYangName = "ipv4-multi-hop-packet-counters"
    ipv4MultiHopPacketCounter.EntityData.SegmentPath = "ipv4-multi-hop-packet-counter"
    ipv4MultiHopPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopPacketCounter.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv4MultiHopPacketCounter.SourceAddress})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4MultiHopPacketCounter.DestinationAddress})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4MultiHopPacketCounter.Location})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4MultiHopPacketCounter.VrfName})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", ipv4MultiHopPacketCounter.HelloTransmitCount})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", ipv4MultiHopPacketCounter.HelloReceiveCount})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", ipv4MultiHopPacketCounter.EchoTransmitCount})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", ipv4MultiHopPacketCounter.EchoReceiveCount})
    ipv4MultiHopPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", ipv4MultiHopPacketCounter.DisplayType})

    ipv4MultiHopPacketCounter.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopPacketCounter.EntityData)
}

// Bfd_SessionDetails
// Table of detailed information about IPv4
// singlehop BFD sessions in the System 
type Bfd_SessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv4 singlehop BFD session. The type is
    // slice of Bfd_SessionDetails_SessionDetail.
    SessionDetail []*Bfd_SessionDetails_SessionDetail
}

func (sessionDetails *Bfd_SessionDetails) GetEntityData() *types.CommonEntityData {
    sessionDetails.EntityData.YFilter = sessionDetails.YFilter
    sessionDetails.EntityData.YangName = "session-details"
    sessionDetails.EntityData.BundleName = "cisco_ios_xr"
    sessionDetails.EntityData.ParentYangName = "bfd"
    sessionDetails.EntityData.SegmentPath = "session-details"
    sessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetails.EntityData.Children = types.NewOrderedMap()
    sessionDetails.EntityData.Children.Append("session-detail", types.YChild{"SessionDetail", nil})
    for i := range sessionDetails.SessionDetail {
        sessionDetails.EntityData.Children.Append(types.GetSegmentPath(sessionDetails.SessionDetail[i]), types.YChild{"SessionDetail", sessionDetails.SessionDetail[i]})
    }
    sessionDetails.EntityData.Leafs = types.NewOrderedMap()

    sessionDetails.EntityData.YListKeys = []string {}

    return &(sessionDetails.EntityData)
}

// Bfd_SessionDetails_SessionDetail
// Detailed information for a single IPv4
// singlehop BFD session
type Bfd_SessionDetails_SessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session status information.
    StatusInformation Bfd_SessionDetails_SessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_SessionDetails_SessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_SessionDetails_SessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_SessionDetails_SessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_SessionDetails_SessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_SessionDetails_SessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_SessionDetails_SessionDetail_AssociationInformation
}

func (sessionDetail *Bfd_SessionDetails_SessionDetail) GetEntityData() *types.CommonEntityData {
    sessionDetail.EntityData.YFilter = sessionDetail.YFilter
    sessionDetail.EntityData.YangName = "session-detail"
    sessionDetail.EntityData.BundleName = "cisco_ios_xr"
    sessionDetail.EntityData.ParentYangName = "session-details"
    sessionDetail.EntityData.SegmentPath = "session-detail"
    sessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDetail.EntityData.Children = types.NewOrderedMap()
    sessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &sessionDetail.StatusInformation})
    sessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &sessionDetail.MpDownloadState})
    sessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &sessionDetail.LspPingInfo})
    sessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range sessionDetail.OwnerInformation {
        sessionDetail.EntityData.Children.Append(types.GetSegmentPath(sessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", sessionDetail.OwnerInformation[i]})
    }
    sessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range sessionDetail.AssociationInformation {
        sessionDetail.EntityData.Children.Append(types.GetSegmentPath(sessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", sessionDetail.AssociationInformation[i]})
    }
    sessionDetail.EntityData.Leafs = types.NewOrderedMap()
    sessionDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", sessionDetail.InterfaceName})
    sessionDetail.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", sessionDetail.DestinationAddress})
    sessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", sessionDetail.Location})

    sessionDetail.EntityData.YListKeys = []string {}

    return &(sessionDetail.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation
// Session status information
type Bfd_SessionDetails_SessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_SessionDetails_SessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_SessionDetails_SessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_SessionDetails_SessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_SessionDetails_SessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_SessionDetails_SessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_SessionDetails_SessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_SessionDetails_SessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_SessionDetails_SessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_SessionDetails_SessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_SessionDetails_SessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_SessionDetails_SessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_SessionDetails_SessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_SessionDetails_SessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_SessionDetails_SessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_SessionDetails_SessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_SessionDetails_SessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_SessionDetails_SessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_SessionDetails_SessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_SessionDetails_SessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_SessionDetails_SessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_SessionDetails_SessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_SessionDetails_SessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_SessionDetails_SessionDetail_MpDownloadState
// MP Dowload State
type Bfd_SessionDetails_SessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_SessionDetails_SessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_SessionDetails_SessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_SessionDetails_SessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_SessionDetails_SessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_SessionDetails_SessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_SessionDetails_SessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_SessionDetails_SessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_SessionDetails_SessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_SessionDetails_SessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_SessionDetails_SessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_SessionDetails_SessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_SessionDetails_SessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_SessionDetails_SessionDetail_AssociationInformation
// Association session information
type Bfd_SessionDetails_SessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_SessionDetails_SessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_SessionDetails_SessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_SessionDetails_SessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_SessionDetails_SessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_SessionDetails_SessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_SessionDetails_SessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_SessionDetails_SessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_SessionDetails_SessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_SessionDetails_SessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_SessionDetails_SessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4SingleHopMultiPaths
// IPv4 single hop multipath
type Bfd_Ipv4SingleHopMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 single hop multipath table. The type is slice of
    // Bfd_Ipv4SingleHopMultiPaths_Ipv4SingleHopMultiPath.
    Ipv4SingleHopMultiPath []*Bfd_Ipv4SingleHopMultiPaths_Ipv4SingleHopMultiPath
}

func (ipv4SingleHopMultiPaths *Bfd_Ipv4SingleHopMultiPaths) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopMultiPaths.EntityData.YFilter = ipv4SingleHopMultiPaths.YFilter
    ipv4SingleHopMultiPaths.EntityData.YangName = "ipv4-single-hop-multi-paths"
    ipv4SingleHopMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopMultiPaths.EntityData.ParentYangName = "bfd"
    ipv4SingleHopMultiPaths.EntityData.SegmentPath = "ipv4-single-hop-multi-paths"
    ipv4SingleHopMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopMultiPaths.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopMultiPaths.EntityData.Children.Append("ipv4-single-hop-multi-path", types.YChild{"Ipv4SingleHopMultiPath", nil})
    for i := range ipv4SingleHopMultiPaths.Ipv4SingleHopMultiPath {
        ipv4SingleHopMultiPaths.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopMultiPaths.Ipv4SingleHopMultiPath[i]), types.YChild{"Ipv4SingleHopMultiPath", ipv4SingleHopMultiPaths.Ipv4SingleHopMultiPath[i]})
    }
    ipv4SingleHopMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopMultiPaths.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopMultiPaths.EntityData)
}

// Bfd_Ipv4SingleHopMultiPaths_Ipv4SingleHopMultiPath
// IPv4 single hop multipath table
type Bfd_Ipv4SingleHopMultiPaths_Ipv4SingleHopMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (ipv4SingleHopMultiPath *Bfd_Ipv4SingleHopMultiPaths_Ipv4SingleHopMultiPath) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopMultiPath.EntityData.YFilter = ipv4SingleHopMultiPath.YFilter
    ipv4SingleHopMultiPath.EntityData.YangName = "ipv4-single-hop-multi-path"
    ipv4SingleHopMultiPath.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopMultiPath.EntityData.ParentYangName = "ipv4-single-hop-multi-paths"
    ipv4SingleHopMultiPath.EntityData.SegmentPath = "ipv4-single-hop-multi-path"
    ipv4SingleHopMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopMultiPath.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopMultiPath.EntityData.Leafs = types.NewOrderedMap()
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4SingleHopMultiPath.InterfaceName})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4SingleHopMultiPath.DestinationAddress})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4SingleHopMultiPath.Location})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4SingleHopMultiPath.SessionSubtype})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4SingleHopMultiPath.State})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", ipv4SingleHopMultiPath.LocalDiscriminator})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4SingleHopMultiPath.NodeId})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", ipv4SingleHopMultiPath.IncomingLabelXr})
    ipv4SingleHopMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", ipv4SingleHopMultiPath.SessionInterfaceName})

    ipv4SingleHopMultiPath.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopMultiPath.EntityData)
}

// Bfd_Ipv4SingleHopSessionBriefs
// Table of brief information about all IPv4
// singlehop BFD sessions in the System
type Bfd_Ipv4SingleHopSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv4 singlehop BFD session. The type is
    // slice of Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief.
    Ipv4SingleHopSessionBrief []*Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief
}

func (ipv4SingleHopSessionBriefs *Bfd_Ipv4SingleHopSessionBriefs) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopSessionBriefs.EntityData.YFilter = ipv4SingleHopSessionBriefs.YFilter
    ipv4SingleHopSessionBriefs.EntityData.YangName = "ipv4-single-hop-session-briefs"
    ipv4SingleHopSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopSessionBriefs.EntityData.ParentYangName = "bfd"
    ipv4SingleHopSessionBriefs.EntityData.SegmentPath = "ipv4-single-hop-session-briefs"
    ipv4SingleHopSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopSessionBriefs.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopSessionBriefs.EntityData.Children.Append("ipv4-single-hop-session-brief", types.YChild{"Ipv4SingleHopSessionBrief", nil})
    for i := range ipv4SingleHopSessionBriefs.Ipv4SingleHopSessionBrief {
        ipv4SingleHopSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(ipv4SingleHopSessionBriefs.Ipv4SingleHopSessionBrief[i]), types.YChild{"Ipv4SingleHopSessionBrief", ipv4SingleHopSessionBriefs.Ipv4SingleHopSessionBrief[i]})
    }
    ipv4SingleHopSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    ipv4SingleHopSessionBriefs.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopSessionBriefs.EntityData)
}

// Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief
// Brief information for a single IPv4 singlehop
// BFD session
type Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation
}

func (ipv4SingleHopSessionBrief *Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief) GetEntityData() *types.CommonEntityData {
    ipv4SingleHopSessionBrief.EntityData.YFilter = ipv4SingleHopSessionBrief.YFilter
    ipv4SingleHopSessionBrief.EntityData.YangName = "ipv4-single-hop-session-brief"
    ipv4SingleHopSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    ipv4SingleHopSessionBrief.EntityData.ParentYangName = "ipv4-single-hop-session-briefs"
    ipv4SingleHopSessionBrief.EntityData.SegmentPath = "ipv4-single-hop-session-brief"
    ipv4SingleHopSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SingleHopSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SingleHopSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SingleHopSessionBrief.EntityData.Children = types.NewOrderedMap()
    ipv4SingleHopSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &ipv4SingleHopSessionBrief.StatusBriefInformation})
    ipv4SingleHopSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4SingleHopSessionBrief.InterfaceName})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4SingleHopSessionBrief.DestinationAddress})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4SingleHopSessionBrief.Location})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4SingleHopSessionBrief.NodeId})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4SingleHopSessionBrief.State})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", ipv4SingleHopSessionBrief.SessionType})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4SingleHopSessionBrief.SessionSubtype})
    ipv4SingleHopSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", ipv4SingleHopSessionBrief.SessionFlags})

    ipv4SingleHopSessionBrief.EntityData.YListKeys = []string {}

    return &(ipv4SingleHopSessionBrief.EntityData)
}

// Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "ipv4-single-hop-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4SingleHopSessionBriefs_Ipv4SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv6MultiHopCounters
// IPv6 multiple hop Counters
type Bfd_Ipv6MultiHopCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv6 multiple hop Packet counters.
    Ipv6MultiHopPacketCounters Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters
}

func (ipv6MultiHopCounters *Bfd_Ipv6MultiHopCounters) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopCounters.EntityData.YFilter = ipv6MultiHopCounters.YFilter
    ipv6MultiHopCounters.EntityData.YangName = "ipv6-multi-hop-counters"
    ipv6MultiHopCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopCounters.EntityData.ParentYangName = "bfd"
    ipv6MultiHopCounters.EntityData.SegmentPath = "ipv6-multi-hop-counters"
    ipv6MultiHopCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopCounters.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopCounters.EntityData.Children.Append("ipv6-multi-hop-packet-counters", types.YChild{"Ipv6MultiHopPacketCounters", &ipv6MultiHopCounters.Ipv6MultiHopPacketCounters})
    ipv6MultiHopCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopCounters.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopCounters.EntityData)
}

// Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters
// Table of IPv6 multiple hop Packet counters
type Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 multiple hop Packet counters. The type is slice of
    // Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters_Ipv6MultiHopPacketCounter.
    Ipv6MultiHopPacketCounter []*Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters_Ipv6MultiHopPacketCounter
}

func (ipv6MultiHopPacketCounters *Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopPacketCounters.EntityData.YFilter = ipv6MultiHopPacketCounters.YFilter
    ipv6MultiHopPacketCounters.EntityData.YangName = "ipv6-multi-hop-packet-counters"
    ipv6MultiHopPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopPacketCounters.EntityData.ParentYangName = "ipv6-multi-hop-counters"
    ipv6MultiHopPacketCounters.EntityData.SegmentPath = "ipv6-multi-hop-packet-counters"
    ipv6MultiHopPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopPacketCounters.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopPacketCounters.EntityData.Children.Append("ipv6-multi-hop-packet-counter", types.YChild{"Ipv6MultiHopPacketCounter", nil})
    for i := range ipv6MultiHopPacketCounters.Ipv6MultiHopPacketCounter {
        ipv6MultiHopPacketCounters.EntityData.Children.Append(types.GetSegmentPath(ipv6MultiHopPacketCounters.Ipv6MultiHopPacketCounter[i]), types.YChild{"Ipv6MultiHopPacketCounter", ipv6MultiHopPacketCounters.Ipv6MultiHopPacketCounter[i]})
    }
    ipv6MultiHopPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv6MultiHopPacketCounters.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopPacketCounters.EntityData)
}

// Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters_Ipv6MultiHopPacketCounter
// IPv4 multiple hop Packet counters
type Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters_Ipv6MultiHopPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (ipv6MultiHopPacketCounter *Bfd_Ipv6MultiHopCounters_Ipv6MultiHopPacketCounters_Ipv6MultiHopPacketCounter) GetEntityData() *types.CommonEntityData {
    ipv6MultiHopPacketCounter.EntityData.YFilter = ipv6MultiHopPacketCounter.YFilter
    ipv6MultiHopPacketCounter.EntityData.YangName = "ipv6-multi-hop-packet-counter"
    ipv6MultiHopPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    ipv6MultiHopPacketCounter.EntityData.ParentYangName = "ipv6-multi-hop-packet-counters"
    ipv6MultiHopPacketCounter.EntityData.SegmentPath = "ipv6-multi-hop-packet-counter"
    ipv6MultiHopPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6MultiHopPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6MultiHopPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6MultiHopPacketCounter.EntityData.Children = types.NewOrderedMap()
    ipv6MultiHopPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv6MultiHopPacketCounter.SourceAddress})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6MultiHopPacketCounter.DestinationAddress})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6MultiHopPacketCounter.Location})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv6MultiHopPacketCounter.VrfName})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", ipv6MultiHopPacketCounter.HelloTransmitCount})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", ipv6MultiHopPacketCounter.HelloReceiveCount})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", ipv6MultiHopPacketCounter.EchoTransmitCount})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", ipv6MultiHopPacketCounter.EchoReceiveCount})
    ipv6MultiHopPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", ipv6MultiHopPacketCounter.DisplayType})

    ipv6MultiHopPacketCounter.EntityData.YListKeys = []string {}

    return &(ipv6MultiHopPacketCounter.EntityData)
}

// Bfd_Ipv6SingleHopLocationSummaries
// Table of summary information about BFD IPv6
// singlehop sessions per location
type Bfd_Ipv6SingleHopLocationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary information for BFD IPv6 singlehop sessions for location. The type
    // is slice of
    // Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary.
    Ipv6SingleHopLocationSummary []*Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary
}

func (ipv6SingleHopLocationSummaries *Bfd_Ipv6SingleHopLocationSummaries) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopLocationSummaries.EntityData.YFilter = ipv6SingleHopLocationSummaries.YFilter
    ipv6SingleHopLocationSummaries.EntityData.YangName = "ipv6-single-hop-location-summaries"
    ipv6SingleHopLocationSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopLocationSummaries.EntityData.ParentYangName = "bfd"
    ipv6SingleHopLocationSummaries.EntityData.SegmentPath = "ipv6-single-hop-location-summaries"
    ipv6SingleHopLocationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopLocationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopLocationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopLocationSummaries.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopLocationSummaries.EntityData.Children.Append("ipv6-single-hop-location-summary", types.YChild{"Ipv6SingleHopLocationSummary", nil})
    for i := range ipv6SingleHopLocationSummaries.Ipv6SingleHopLocationSummary {
        ipv6SingleHopLocationSummaries.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopLocationSummaries.Ipv6SingleHopLocationSummary[i]), types.YChild{"Ipv6SingleHopLocationSummary", ipv6SingleHopLocationSummaries.Ipv6SingleHopLocationSummary[i]})
    }
    ipv6SingleHopLocationSummaries.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopLocationSummaries.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopLocationSummaries.EntityData)
}

// Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary
// Summary information for BFD IPv6 singlehop
// sessions for location
type Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Location Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    LocationName interface{}

    // Statistics of states for sessions.
    SessionState Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary_SessionState
}

func (ipv6SingleHopLocationSummary *Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopLocationSummary.EntityData.YFilter = ipv6SingleHopLocationSummary.YFilter
    ipv6SingleHopLocationSummary.EntityData.YangName = "ipv6-single-hop-location-summary"
    ipv6SingleHopLocationSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopLocationSummary.EntityData.ParentYangName = "ipv6-single-hop-location-summaries"
    ipv6SingleHopLocationSummary.EntityData.SegmentPath = "ipv6-single-hop-location-summary" + types.AddKeyToken(ipv6SingleHopLocationSummary.LocationName, "location-name")
    ipv6SingleHopLocationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopLocationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopLocationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopLocationSummary.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopLocationSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv6SingleHopLocationSummary.SessionState})
    ipv6SingleHopLocationSummary.EntityData.Leafs = types.NewOrderedMap()
    ipv6SingleHopLocationSummary.EntityData.Leafs.Append("location-name", types.YLeaf{"LocationName", ipv6SingleHopLocationSummary.LocationName})

    ipv6SingleHopLocationSummary.EntityData.YListKeys = []string {"LocationName"}

    return &(ipv6SingleHopLocationSummary.EntityData)
}

// Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}

    // Number of sessions in retry state. The type is interface{} with range:
    // 0..4294967295.
    RetryCount interface{}

    // Number of sessions in standby state. The type is interface{} with range:
    // 0..4294967295.
    StandbyCount interface{}
}

func (sessionState *Bfd_Ipv6SingleHopLocationSummaries_Ipv6SingleHopLocationSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv6-single-hop-location-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})
    sessionState.EntityData.Leafs.Append("retry-count", types.YLeaf{"RetryCount", sessionState.RetryCount})
    sessionState.EntityData.Leafs.Append("standby-count", types.YLeaf{"StandbyCount", sessionState.StandbyCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

// Bfd_LabelCounters
// Label Counters
type Bfd_LabelCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Label Packet counters.
    LabelPacketCounters Bfd_LabelCounters_LabelPacketCounters
}

func (labelCounters *Bfd_LabelCounters) GetEntityData() *types.CommonEntityData {
    labelCounters.EntityData.YFilter = labelCounters.YFilter
    labelCounters.EntityData.YangName = "label-counters"
    labelCounters.EntityData.BundleName = "cisco_ios_xr"
    labelCounters.EntityData.ParentYangName = "bfd"
    labelCounters.EntityData.SegmentPath = "label-counters"
    labelCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelCounters.EntityData.Children = types.NewOrderedMap()
    labelCounters.EntityData.Children.Append("label-packet-counters", types.YChild{"LabelPacketCounters", &labelCounters.LabelPacketCounters})
    labelCounters.EntityData.Leafs = types.NewOrderedMap()

    labelCounters.EntityData.YListKeys = []string {}

    return &(labelCounters.EntityData)
}

// Bfd_LabelCounters_LabelPacketCounters
// Table of Label Packet counters
type Bfd_LabelCounters_LabelPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Label Packet counters. The type is slice of
    // Bfd_LabelCounters_LabelPacketCounters_LabelPacketCounter.
    LabelPacketCounter []*Bfd_LabelCounters_LabelPacketCounters_LabelPacketCounter
}

func (labelPacketCounters *Bfd_LabelCounters_LabelPacketCounters) GetEntityData() *types.CommonEntityData {
    labelPacketCounters.EntityData.YFilter = labelPacketCounters.YFilter
    labelPacketCounters.EntityData.YangName = "label-packet-counters"
    labelPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    labelPacketCounters.EntityData.ParentYangName = "label-counters"
    labelPacketCounters.EntityData.SegmentPath = "label-packet-counters"
    labelPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelPacketCounters.EntityData.Children = types.NewOrderedMap()
    labelPacketCounters.EntityData.Children.Append("label-packet-counter", types.YChild{"LabelPacketCounter", nil})
    for i := range labelPacketCounters.LabelPacketCounter {
        labelPacketCounters.EntityData.Children.Append(types.GetSegmentPath(labelPacketCounters.LabelPacketCounter[i]), types.YChild{"LabelPacketCounter", labelPacketCounters.LabelPacketCounter[i]})
    }
    labelPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    labelPacketCounters.EntityData.YListKeys = []string {}

    return &(labelPacketCounters.EntityData)
}

// Bfd_LabelCounters_LabelPacketCounters_LabelPacketCounter
// Interface Label Packet counters
type Bfd_LabelCounters_LabelPacketCounters_LabelPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (labelPacketCounter *Bfd_LabelCounters_LabelPacketCounters_LabelPacketCounter) GetEntityData() *types.CommonEntityData {
    labelPacketCounter.EntityData.YFilter = labelPacketCounter.YFilter
    labelPacketCounter.EntityData.YangName = "label-packet-counter"
    labelPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    labelPacketCounter.EntityData.ParentYangName = "label-packet-counters"
    labelPacketCounter.EntityData.SegmentPath = "label-packet-counter"
    labelPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    labelPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    labelPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    labelPacketCounter.EntityData.Children = types.NewOrderedMap()
    labelPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    labelPacketCounter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", labelPacketCounter.InterfaceName})
    labelPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", labelPacketCounter.Location})
    labelPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", labelPacketCounter.HelloTransmitCount})
    labelPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", labelPacketCounter.HelloReceiveCount})
    labelPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", labelPacketCounter.EchoTransmitCount})
    labelPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", labelPacketCounter.EchoReceiveCount})
    labelPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", labelPacketCounter.DisplayType})

    labelPacketCounter.EntityData.YListKeys = []string {}

    return &(labelPacketCounter.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails
// Table of detailed information about all IPv4 BFD
// over MPLS-TE Head sessions in the System
type Bfd_Ipv4bfDoMplsteHeadSessionDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed information for a single IPv4 BFD over MPLS-TE head session. The
    // type is slice of
    // Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail.
    Ipv4bfDoMplsteHeadSessionDetail []*Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail
}

func (ipv4bfDoMplsteHeadSessionDetails *Bfd_Ipv4bfDoMplsteHeadSessionDetails) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadSessionDetails.EntityData.YFilter = ipv4bfDoMplsteHeadSessionDetails.YFilter
    ipv4bfDoMplsteHeadSessionDetails.EntityData.YangName = "ipv4bf-do-mplste-head-session-details"
    ipv4bfDoMplsteHeadSessionDetails.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadSessionDetails.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteHeadSessionDetails.EntityData.SegmentPath = "ipv4bf-do-mplste-head-session-details"
    ipv4bfDoMplsteHeadSessionDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadSessionDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadSessionDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadSessionDetails.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSessionDetails.EntityData.Children.Append("ipv4bf-do-mplste-head-session-detail", types.YChild{"Ipv4bfDoMplsteHeadSessionDetail", nil})
    for i := range ipv4bfDoMplsteHeadSessionDetails.Ipv4bfDoMplsteHeadSessionDetail {
        ipv4bfDoMplsteHeadSessionDetails.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteHeadSessionDetails.Ipv4bfDoMplsteHeadSessionDetail[i]), types.YChild{"Ipv4bfDoMplsteHeadSessionDetail", ipv4bfDoMplsteHeadSessionDetails.Ipv4bfDoMplsteHeadSessionDetail[i]})
    }
    ipv4bfDoMplsteHeadSessionDetails.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteHeadSessionDetails.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadSessionDetails.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail
// Detailed information for a single IPv4 BFD over
// MPLS-TE head session
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session status information.
    StatusInformation Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation

    // MP Dowload State.
    MpDownloadState Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState

    // LSP Ping Info.
    LspPingInfo Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_OwnerInformation

    // Association session information. The type is slice of
    // Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation.
    AssociationInformation []*Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation
}

func (ipv4bfDoMplsteHeadSessionDetail *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadSessionDetail.EntityData.YFilter = ipv4bfDoMplsteHeadSessionDetail.YFilter
    ipv4bfDoMplsteHeadSessionDetail.EntityData.YangName = "ipv4bf-do-mplste-head-session-detail"
    ipv4bfDoMplsteHeadSessionDetail.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadSessionDetail.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-details"
    ipv4bfDoMplsteHeadSessionDetail.EntityData.SegmentPath = "ipv4bf-do-mplste-head-session-detail"
    ipv4bfDoMplsteHeadSessionDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadSessionDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadSessionDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadSessionDetail.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append("status-information", types.YChild{"StatusInformation", &ipv4bfDoMplsteHeadSessionDetail.StatusInformation})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append("mp-download-state", types.YChild{"MpDownloadState", &ipv4bfDoMplsteHeadSessionDetail.MpDownloadState})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append("lsp-ping-info", types.YChild{"LspPingInfo", &ipv4bfDoMplsteHeadSessionDetail.LspPingInfo})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range ipv4bfDoMplsteHeadSessionDetail.OwnerInformation {
        ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteHeadSessionDetail.OwnerInformation[i]), types.YChild{"OwnerInformation", ipv4bfDoMplsteHeadSessionDetail.OwnerInformation[i]})
    }
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range ipv4bfDoMplsteHeadSessionDetail.AssociationInformation {
        ipv4bfDoMplsteHeadSessionDetail.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteHeadSessionDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", ipv4bfDoMplsteHeadSessionDetail.AssociationInformation[i]})
    }
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4bfDoMplsteHeadSessionDetail.InterfaceName})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteHeadSessionDetail.VrfName})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteHeadSessionDetail.IncomingLabel})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteHeadSessionDetail.FeCtype})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteHeadSessionDetail.FecSubgroupId})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteHeadSessionDetail.Feclspid})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteHeadSessionDetail.FecTunnelId})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteHeadSessionDetail.FecExtendedTunnelId})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteHeadSessionDetail.FecSource})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteHeadSessionDetail.FecDestination})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteHeadSessionDetail.Fecp2mpid})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteHeadSessionDetail.FecSubgroupOriginator})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteHeadSessionDetail.FecCtype})
    ipv4bfDoMplsteHeadSessionDetail.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteHeadSessionDetail.Location})

    ipv4bfDoMplsteHeadSessionDetail.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadSessionDetail.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation
// Session status information
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Session's Remote discriminator. The type is interface{} with range:
    // 0..4294967295.
    RemoteDiscriminator interface{}

    // Number of times session state went to UP. The type is interface{} with
    // range: 0..4294967295.
    ToUpStateCount interface{}

    // Desired minimum echo transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    DesiredMinimumEchoTransmitInterval interface{}

    // Remote Negotiated Interval in milli-seconds. The type is interface{} with
    // range: 0..4294967295. Units are millisecond.
    RemoteNegotiatedInterval interface{}

    // Number of Latency Samples. Time between Transmit and Receive. The type is
    // interface{} with range: 0..4294967295.
    LatencyNumber interface{}

    // Minimum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMinimum interface{}

    // Maximum value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyMaximum interface{}

    // Average value of Latency (in micro-seconds). The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    LatencyAverage interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Internal Label. The type is interface{} with range: 0..4294967295.
    InternalLabel interface{}

    // Source address.
    SourceAddress Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_SourceAddress

    // Time since last state change.
    LastStateChange Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_LastStateChange

    // Transmit Packet.
    TransmitPacket Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_TransmitPacket

    // Receive Packet.
    ReceivePacket Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_ReceivePacket

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation

    // Statistics of Interval between Async Packets Transmitted (in
    // milli-seconds).
    AsyncTransmitStatistics Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncTransmitStatistics

    // Statistics of Interval between Async Packets Received (in milli-seconds).
    AsyncReceiveStatistics Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncReceiveStatistics

    // Statistics of Interval between Echo Packets Transmitted (in milli-seconds).
    EchoTransmitStatistics Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoTransmitStatistics

    // Statistics of Interval between Echo Packets Received (in milli-seconds).
    EchoReceivedStatistics Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoReceivedStatistics
}

func (statusInformation *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation) GetEntityData() *types.CommonEntityData {
    statusInformation.EntityData.YFilter = statusInformation.YFilter
    statusInformation.EntityData.YangName = "status-information"
    statusInformation.EntityData.BundleName = "cisco_ios_xr"
    statusInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-detail"
    statusInformation.EntityData.SegmentPath = "status-information"
    statusInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusInformation.EntityData.Children = types.NewOrderedMap()
    statusInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &statusInformation.SourceAddress})
    statusInformation.EntityData.Children.Append("last-state-change", types.YChild{"LastStateChange", &statusInformation.LastStateChange})
    statusInformation.EntityData.Children.Append("transmit-packet", types.YChild{"TransmitPacket", &statusInformation.TransmitPacket})
    statusInformation.EntityData.Children.Append("receive-packet", types.YChild{"ReceivePacket", &statusInformation.ReceivePacket})
    statusInformation.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &statusInformation.StatusBriefInformation})
    statusInformation.EntityData.Children.Append("async-transmit-statistics", types.YChild{"AsyncTransmitStatistics", &statusInformation.AsyncTransmitStatistics})
    statusInformation.EntityData.Children.Append("async-receive-statistics", types.YChild{"AsyncReceiveStatistics", &statusInformation.AsyncReceiveStatistics})
    statusInformation.EntityData.Children.Append("echo-transmit-statistics", types.YChild{"EchoTransmitStatistics", &statusInformation.EchoTransmitStatistics})
    statusInformation.EntityData.Children.Append("echo-received-statistics", types.YChild{"EchoReceivedStatistics", &statusInformation.EchoReceivedStatistics})
    statusInformation.EntityData.Leafs = types.NewOrderedMap()
    statusInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", statusInformation.Sessiontype})
    statusInformation.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", statusInformation.SessionSubtype})
    statusInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", statusInformation.State})
    statusInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", statusInformation.LocalDiscriminator})
    statusInformation.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", statusInformation.RemoteDiscriminator})
    statusInformation.EntityData.Leafs.Append("to-up-state-count", types.YLeaf{"ToUpStateCount", statusInformation.ToUpStateCount})
    statusInformation.EntityData.Leafs.Append("desired-minimum-echo-transmit-interval", types.YLeaf{"DesiredMinimumEchoTransmitInterval", statusInformation.DesiredMinimumEchoTransmitInterval})
    statusInformation.EntityData.Leafs.Append("remote-negotiated-interval", types.YLeaf{"RemoteNegotiatedInterval", statusInformation.RemoteNegotiatedInterval})
    statusInformation.EntityData.Leafs.Append("latency-number", types.YLeaf{"LatencyNumber", statusInformation.LatencyNumber})
    statusInformation.EntityData.Leafs.Append("latency-minimum", types.YLeaf{"LatencyMinimum", statusInformation.LatencyMinimum})
    statusInformation.EntityData.Leafs.Append("latency-maximum", types.YLeaf{"LatencyMaximum", statusInformation.LatencyMaximum})
    statusInformation.EntityData.Leafs.Append("latency-average", types.YLeaf{"LatencyAverage", statusInformation.LatencyAverage})
    statusInformation.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", statusInformation.NodeId})
    statusInformation.EntityData.Leafs.Append("internal-label", types.YLeaf{"InternalLabel", statusInformation.InternalLabel})

    statusInformation.EntityData.YListKeys = []string {}

    return &(statusInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_SourceAddress
// Source address
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sourceAddress *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "status-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", sourceAddress.Afi})
    sourceAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", sourceAddress.Dummy})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_LastStateChange
// Time since last state change
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_LastStateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of days since last session state transition. The type is interface{}
    // with range: 0..4294967295. Units are day.
    Days interface{}

    // Number of hours since last session state transition. The type is
    // interface{} with range: 0..255. Units are hour.
    Hours interface{}

    // Number of mins since last session state transition. The type is interface{}
    // with range: 0..255. Units are minute.
    Minutes interface{}

    // Number of seconds since last session state transition. The type is
    // interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (lastStateChange *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_LastStateChange) GetEntityData() *types.CommonEntityData {
    lastStateChange.EntityData.YFilter = lastStateChange.YFilter
    lastStateChange.EntityData.YangName = "last-state-change"
    lastStateChange.EntityData.BundleName = "cisco_ios_xr"
    lastStateChange.EntityData.ParentYangName = "status-information"
    lastStateChange.EntityData.SegmentPath = "last-state-change"
    lastStateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastStateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastStateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastStateChange.EntityData.Children = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs = types.NewOrderedMap()
    lastStateChange.EntityData.Leafs.Append("days", types.YLeaf{"Days", lastStateChange.Days})
    lastStateChange.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lastStateChange.Hours})
    lastStateChange.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lastStateChange.Minutes})
    lastStateChange.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lastStateChange.Seconds})

    lastStateChange.EntityData.YListKeys = []string {}

    return &(lastStateChange.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_TransmitPacket
// Transmit Packet
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_TransmitPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (transmitPacket *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_TransmitPacket) GetEntityData() *types.CommonEntityData {
    transmitPacket.EntityData.YFilter = transmitPacket.YFilter
    transmitPacket.EntityData.YangName = "transmit-packet"
    transmitPacket.EntityData.BundleName = "cisco_ios_xr"
    transmitPacket.EntityData.ParentYangName = "status-information"
    transmitPacket.EntityData.SegmentPath = "transmit-packet"
    transmitPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transmitPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transmitPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transmitPacket.EntityData.Children = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs = types.NewOrderedMap()
    transmitPacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", transmitPacket.Version})
    transmitPacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", transmitPacket.Diagnostic})
    transmitPacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", transmitPacket.IhearYou})
    transmitPacket.EntityData.Leafs.Append("state", types.YLeaf{"State", transmitPacket.State})
    transmitPacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", transmitPacket.Demand})
    transmitPacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", transmitPacket.Poll})
    transmitPacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", transmitPacket.Final})
    transmitPacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", transmitPacket.ControlPlaneIndependent})
    transmitPacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", transmitPacket.AuthenticationPresent})
    transmitPacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", transmitPacket.DetectionMultiplier})
    transmitPacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", transmitPacket.Length})
    transmitPacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", transmitPacket.MyDiscriminator})
    transmitPacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", transmitPacket.YourDiscriminator})
    transmitPacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", transmitPacket.DesiredMinimumTransmitInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", transmitPacket.RequiredMinimumReceiveInterval})
    transmitPacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", transmitPacket.RequiredMinimumEchoReceiveInterval})

    transmitPacket.EntityData.YListKeys = []string {}

    return &(transmitPacket.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_ReceivePacket
// Receive Packet
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_ReceivePacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version. The type is interface{} with range: 0..255.
    Version interface{}

    // Diagnostic. The type is BfdMgmtSessionDiag.
    Diagnostic interface{}

    // I Hear You (v0). The type is interface{} with range:
    // -2147483648..2147483647.
    IhearYou interface{}

    // State (v1). The type is BfdMgmtSessionState.
    State interface{}

    // Demand mode. The type is interface{} with range: -2147483648..2147483647.
    Demand interface{}

    // Poll bit. The type is interface{} with range: -2147483648..2147483647.
    Poll interface{}

    // Final bit. The type is interface{} with range: -2147483648..2147483647.
    Final interface{}

    // BFD implementation does not share fate with its control plane. The type is
    // interface{} with range: -2147483648..2147483647.
    ControlPlaneIndependent interface{}

    // Requesting authentication for the session. The type is interface{} with
    // range: -2147483648..2147483647.
    AuthenticationPresent interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Length. The type is interface{} with range: 0..4294967295.
    Length interface{}

    // My Discriminator. The type is interface{} with range: 0..4294967295.
    MyDiscriminator interface{}

    // Your Discriminator. The type is interface{} with range: 0..4294967295.
    YourDiscriminator interface{}

    // Desired minimum transmit interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    DesiredMinimumTransmitInterval interface{}

    // Required receive interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    RequiredMinimumReceiveInterval interface{}

    // Required echo receive interval in micro-seconds. The type is interface{}
    // with range: 0..4294967295. Units are microsecond.
    RequiredMinimumEchoReceiveInterval interface{}
}

func (receivePacket *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_ReceivePacket) GetEntityData() *types.CommonEntityData {
    receivePacket.EntityData.YFilter = receivePacket.YFilter
    receivePacket.EntityData.YangName = "receive-packet"
    receivePacket.EntityData.BundleName = "cisco_ios_xr"
    receivePacket.EntityData.ParentYangName = "status-information"
    receivePacket.EntityData.SegmentPath = "receive-packet"
    receivePacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receivePacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receivePacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receivePacket.EntityData.Children = types.NewOrderedMap()
    receivePacket.EntityData.Leafs = types.NewOrderedMap()
    receivePacket.EntityData.Leafs.Append("version", types.YLeaf{"Version", receivePacket.Version})
    receivePacket.EntityData.Leafs.Append("diagnostic", types.YLeaf{"Diagnostic", receivePacket.Diagnostic})
    receivePacket.EntityData.Leafs.Append("ihear-you", types.YLeaf{"IhearYou", receivePacket.IhearYou})
    receivePacket.EntityData.Leafs.Append("state", types.YLeaf{"State", receivePacket.State})
    receivePacket.EntityData.Leafs.Append("demand", types.YLeaf{"Demand", receivePacket.Demand})
    receivePacket.EntityData.Leafs.Append("poll", types.YLeaf{"Poll", receivePacket.Poll})
    receivePacket.EntityData.Leafs.Append("final", types.YLeaf{"Final", receivePacket.Final})
    receivePacket.EntityData.Leafs.Append("control-plane-independent", types.YLeaf{"ControlPlaneIndependent", receivePacket.ControlPlaneIndependent})
    receivePacket.EntityData.Leafs.Append("authentication-present", types.YLeaf{"AuthenticationPresent", receivePacket.AuthenticationPresent})
    receivePacket.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", receivePacket.DetectionMultiplier})
    receivePacket.EntityData.Leafs.Append("length", types.YLeaf{"Length", receivePacket.Length})
    receivePacket.EntityData.Leafs.Append("my-discriminator", types.YLeaf{"MyDiscriminator", receivePacket.MyDiscriminator})
    receivePacket.EntityData.Leafs.Append("your-discriminator", types.YLeaf{"YourDiscriminator", receivePacket.YourDiscriminator})
    receivePacket.EntityData.Leafs.Append("desired-minimum-transmit-interval", types.YLeaf{"DesiredMinimumTransmitInterval", receivePacket.DesiredMinimumTransmitInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-receive-interval", types.YLeaf{"RequiredMinimumReceiveInterval", receivePacket.RequiredMinimumReceiveInterval})
    receivePacket.EntityData.Leafs.Append("required-minimum-echo-receive-interval", types.YLeaf{"RequiredMinimumEchoReceiveInterval", receivePacket.RequiredMinimumEchoReceiveInterval})

    receivePacket.EntityData.YListKeys = []string {}

    return &(receivePacket.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "status-information"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncTransmitStatistics
// Statistics of Interval between Async Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncTransmitStatistics *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncTransmitStatistics) GetEntityData() *types.CommonEntityData {
    asyncTransmitStatistics.EntityData.YFilter = asyncTransmitStatistics.YFilter
    asyncTransmitStatistics.EntityData.YangName = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncTransmitStatistics.EntityData.ParentYangName = "status-information"
    asyncTransmitStatistics.EntityData.SegmentPath = "async-transmit-statistics"
    asyncTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncTransmitStatistics.Number})
    asyncTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncTransmitStatistics.Minimum})
    asyncTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncTransmitStatistics.Maximum})
    asyncTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncTransmitStatistics.Average})
    asyncTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncTransmitStatistics.Last})

    asyncTransmitStatistics.EntityData.YListKeys = []string {}

    return &(asyncTransmitStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncReceiveStatistics
// Statistics of Interval between Async Packets
// Received (in milli-seconds)
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncReceiveStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (asyncReceiveStatistics *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_AsyncReceiveStatistics) GetEntityData() *types.CommonEntityData {
    asyncReceiveStatistics.EntityData.YFilter = asyncReceiveStatistics.YFilter
    asyncReceiveStatistics.EntityData.YangName = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.BundleName = "cisco_ios_xr"
    asyncReceiveStatistics.EntityData.ParentYangName = "status-information"
    asyncReceiveStatistics.EntityData.SegmentPath = "async-receive-statistics"
    asyncReceiveStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncReceiveStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncReceiveStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncReceiveStatistics.EntityData.Children = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs = types.NewOrderedMap()
    asyncReceiveStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", asyncReceiveStatistics.Number})
    asyncReceiveStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", asyncReceiveStatistics.Minimum})
    asyncReceiveStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", asyncReceiveStatistics.Maximum})
    asyncReceiveStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", asyncReceiveStatistics.Average})
    asyncReceiveStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", asyncReceiveStatistics.Last})

    asyncReceiveStatistics.EntityData.YListKeys = []string {}

    return &(asyncReceiveStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoTransmitStatistics
// Statistics of Interval between Echo Packets
// Transmitted (in milli-seconds)
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoTransmitStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoTransmitStatistics *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoTransmitStatistics) GetEntityData() *types.CommonEntityData {
    echoTransmitStatistics.EntityData.YFilter = echoTransmitStatistics.YFilter
    echoTransmitStatistics.EntityData.YangName = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoTransmitStatistics.EntityData.ParentYangName = "status-information"
    echoTransmitStatistics.EntityData.SegmentPath = "echo-transmit-statistics"
    echoTransmitStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoTransmitStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoTransmitStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoTransmitStatistics.EntityData.Children = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoTransmitStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoTransmitStatistics.Number})
    echoTransmitStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoTransmitStatistics.Minimum})
    echoTransmitStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoTransmitStatistics.Maximum})
    echoTransmitStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoTransmitStatistics.Average})
    echoTransmitStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoTransmitStatistics.Last})

    echoTransmitStatistics.EntityData.YListKeys = []string {}

    return &(echoTransmitStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoReceivedStatistics
// Statistics of Interval between Echo Packets
// Received (in milli-seconds)
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoReceivedStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Interval Samples between Packets sent/received. The type is
    // interface{} with range: 0..4294967295.
    Number interface{}

    // Minimum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Minimum interface{}

    // Maximum of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Maximum interface{}

    // Average of Transmit/Receive Interval (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Average interface{}

    // Time since last Transmit/Receive (in milli-seconds). The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    Last interface{}
}

func (echoReceivedStatistics *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_StatusInformation_EchoReceivedStatistics) GetEntityData() *types.CommonEntityData {
    echoReceivedStatistics.EntityData.YFilter = echoReceivedStatistics.YFilter
    echoReceivedStatistics.EntityData.YangName = "echo-received-statistics"
    echoReceivedStatistics.EntityData.BundleName = "cisco_ios_xr"
    echoReceivedStatistics.EntityData.ParentYangName = "status-information"
    echoReceivedStatistics.EntityData.SegmentPath = "echo-received-statistics"
    echoReceivedStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoReceivedStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoReceivedStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoReceivedStatistics.EntityData.Children = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs = types.NewOrderedMap()
    echoReceivedStatistics.EntityData.Leafs.Append("number", types.YLeaf{"Number", echoReceivedStatistics.Number})
    echoReceivedStatistics.EntityData.Leafs.Append("minimum", types.YLeaf{"Minimum", echoReceivedStatistics.Minimum})
    echoReceivedStatistics.EntityData.Leafs.Append("maximum", types.YLeaf{"Maximum", echoReceivedStatistics.Maximum})
    echoReceivedStatistics.EntityData.Leafs.Append("average", types.YLeaf{"Average", echoReceivedStatistics.Average})
    echoReceivedStatistics.EntityData.Leafs.Append("last", types.YLeaf{"Last", echoReceivedStatistics.Last})

    echoReceivedStatistics.EntityData.YListKeys = []string {}

    return &(echoReceivedStatistics.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState
// MP Dowload State
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MP Download State. The type is BfdMpDownloadState.
    MpDownloadState interface{}

    // Change time.
    ChangeTime Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState_ChangeTime
}

func (mpDownloadState *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState) GetEntityData() *types.CommonEntityData {
    mpDownloadState.EntityData.YFilter = mpDownloadState.YFilter
    mpDownloadState.EntityData.YangName = "mp-download-state"
    mpDownloadState.EntityData.BundleName = "cisco_ios_xr"
    mpDownloadState.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-detail"
    mpDownloadState.EntityData.SegmentPath = "mp-download-state"
    mpDownloadState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpDownloadState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpDownloadState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpDownloadState.EntityData.Children = types.NewOrderedMap()
    mpDownloadState.EntityData.Children.Append("change-time", types.YChild{"ChangeTime", &mpDownloadState.ChangeTime})
    mpDownloadState.EntityData.Leafs = types.NewOrderedMap()
    mpDownloadState.EntityData.Leafs.Append("mp-download-state", types.YLeaf{"MpDownloadState", mpDownloadState.MpDownloadState})

    mpDownloadState.EntityData.YListKeys = []string {}

    return &(mpDownloadState.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState_ChangeTime
// Change time
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState_ChangeTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (changeTime *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_MpDownloadState_ChangeTime) GetEntityData() *types.CommonEntityData {
    changeTime.EntityData.YFilter = changeTime.YFilter
    changeTime.EntityData.YangName = "change-time"
    changeTime.EntityData.BundleName = "cisco_ios_xr"
    changeTime.EntityData.ParentYangName = "mp-download-state"
    changeTime.EntityData.SegmentPath = "change-time"
    changeTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeTime.EntityData.Children = types.NewOrderedMap()
    changeTime.EntityData.Leafs = types.NewOrderedMap()
    changeTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", changeTime.Seconds})
    changeTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", changeTime.Nanoseconds})

    changeTime.EntityData.YListKeys = []string {}

    return &(changeTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo
// LSP Ping Info
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Ping Tx count. The type is interface{} with range: 0..4294967295.
    LspPingTxCount interface{}

    // LSP Ping Tx error count. The type is interface{} with range: 0..4294967295.
    LspPingTxErrorCount interface{}

    // LSP Ping Tx last result. The type is string.
    LspPingTxLastRc interface{}

    // LSP Ping Tx last error. The type is string.
    LspPingTxLastErrorRc interface{}

    // LSP Ping Rx last received discriminator. The type is interface{} with
    // range: 0..4294967295.
    LspPingRxLastDiscr interface{}

    // LSP Ping numer of times received. The type is interface{} with range:
    // 0..4294967295.
    LspPingRxCount interface{}

    // LSP Ping Rx Last Code. The type is interface{} with range: 0..255.
    LspPingRxLastCode interface{}

    // LSP Ping Rx Last Subcode. The type is interface{} with range: 0..255.
    LspPingRxLastSubcode interface{}

    // LSP Ping Rx Last Output. The type is string.
    LspPingRxLastOutput interface{}

    // LSP Ping last sent time.
    LspPingTxLastTime Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastTime

    // LSP Ping last error time.
    LspPingTxLastErrorTime Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastErrorTime

    // LSP Ping last received time.
    LspPingRxLastTime Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingRxLastTime
}

func (lspPingInfo *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo) GetEntityData() *types.CommonEntityData {
    lspPingInfo.EntityData.YFilter = lspPingInfo.YFilter
    lspPingInfo.EntityData.YangName = "lsp-ping-info"
    lspPingInfo.EntityData.BundleName = "cisco_ios_xr"
    lspPingInfo.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-detail"
    lspPingInfo.EntityData.SegmentPath = "lsp-ping-info"
    lspPingInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingInfo.EntityData.Children = types.NewOrderedMap()
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-time", types.YChild{"LspPingTxLastTime", &lspPingInfo.LspPingTxLastTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-tx-last-error-time", types.YChild{"LspPingTxLastErrorTime", &lspPingInfo.LspPingTxLastErrorTime})
    lspPingInfo.EntityData.Children.Append("lsp-ping-rx-last-time", types.YChild{"LspPingRxLastTime", &lspPingInfo.LspPingRxLastTime})
    lspPingInfo.EntityData.Leafs = types.NewOrderedMap()
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-count", types.YLeaf{"LspPingTxCount", lspPingInfo.LspPingTxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-error-count", types.YLeaf{"LspPingTxErrorCount", lspPingInfo.LspPingTxErrorCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-rc", types.YLeaf{"LspPingTxLastRc", lspPingInfo.LspPingTxLastRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-tx-last-error-rc", types.YLeaf{"LspPingTxLastErrorRc", lspPingInfo.LspPingTxLastErrorRc})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-discr", types.YLeaf{"LspPingRxLastDiscr", lspPingInfo.LspPingRxLastDiscr})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-count", types.YLeaf{"LspPingRxCount", lspPingInfo.LspPingRxCount})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-code", types.YLeaf{"LspPingRxLastCode", lspPingInfo.LspPingRxLastCode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-subcode", types.YLeaf{"LspPingRxLastSubcode", lspPingInfo.LspPingRxLastSubcode})
    lspPingInfo.EntityData.Leafs.Append("lsp-ping-rx-last-output", types.YLeaf{"LspPingRxLastOutput", lspPingInfo.LspPingRxLastOutput})

    lspPingInfo.EntityData.YListKeys = []string {}

    return &(lspPingInfo.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastTime
// LSP Ping last sent time
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastTime *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastTime.EntityData.YFilter = lspPingTxLastTime.YFilter
    lspPingTxLastTime.EntityData.YangName = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastTime.EntityData.SegmentPath = "lsp-ping-tx-last-time"
    lspPingTxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastTime.Seconds})
    lspPingTxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastTime.Nanoseconds})

    lspPingTxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastErrorTime
// LSP Ping last error time
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastErrorTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingTxLastErrorTime *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingTxLastErrorTime) GetEntityData() *types.CommonEntityData {
    lspPingTxLastErrorTime.EntityData.YFilter = lspPingTxLastErrorTime.YFilter
    lspPingTxLastErrorTime.EntityData.YangName = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingTxLastErrorTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingTxLastErrorTime.EntityData.SegmentPath = "lsp-ping-tx-last-error-time"
    lspPingTxLastErrorTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingTxLastErrorTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingTxLastErrorTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingTxLastErrorTime.EntityData.Children = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingTxLastErrorTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingTxLastErrorTime.Seconds})
    lspPingTxLastErrorTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingTxLastErrorTime.Nanoseconds})

    lspPingTxLastErrorTime.EntityData.YListKeys = []string {}

    return &(lspPingTxLastErrorTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingRxLastTime
// LSP Ping last received time
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingRxLastTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // seconds. The type is interface{} with range: 0..18446744073709551615. Units
    // are second.
    Seconds interface{}

    // nanoseconds. The type is interface{} with range: 0..4294967295. Units are
    // nanosecond.
    Nanoseconds interface{}
}

func (lspPingRxLastTime *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_LspPingInfo_LspPingRxLastTime) GetEntityData() *types.CommonEntityData {
    lspPingRxLastTime.EntityData.YFilter = lspPingRxLastTime.YFilter
    lspPingRxLastTime.EntityData.YangName = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.BundleName = "cisco_ios_xr"
    lspPingRxLastTime.EntityData.ParentYangName = "lsp-ping-info"
    lspPingRxLastTime.EntityData.SegmentPath = "lsp-ping-rx-last-time"
    lspPingRxLastTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspPingRxLastTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspPingRxLastTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspPingRxLastTime.EntityData.Children = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs = types.NewOrderedMap()
    lspPingRxLastTime.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", lspPingRxLastTime.Seconds})
    lspPingRxLastTime.EntityData.Leafs.Append("nanoseconds", types.YLeaf{"Nanoseconds", lspPingRxLastTime.Nanoseconds})

    lspPingRxLastTime.EntityData.YListKeys = []string {}

    return &(lspPingRxLastTime.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-detail"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation
// Association session information
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "ipv4bf-do-mplste-head-session-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_Ipv4bfDoMplsteHeadSessionDetails_Ipv4bfDoMplsteHeadSessionDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_RelationBriefs
// Table of brief information about all BFD
// relations in the System
type Bfd_RelationBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for relation of a single BFD session. The type is slice
    // of Bfd_RelationBriefs_RelationBrief.
    RelationBrief []*Bfd_RelationBriefs_RelationBrief
}

func (relationBriefs *Bfd_RelationBriefs) GetEntityData() *types.CommonEntityData {
    relationBriefs.EntityData.YFilter = relationBriefs.YFilter
    relationBriefs.EntityData.YangName = "relation-briefs"
    relationBriefs.EntityData.BundleName = "cisco_ios_xr"
    relationBriefs.EntityData.ParentYangName = "bfd"
    relationBriefs.EntityData.SegmentPath = "relation-briefs"
    relationBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relationBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relationBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relationBriefs.EntityData.Children = types.NewOrderedMap()
    relationBriefs.EntityData.Children.Append("relation-brief", types.YChild{"RelationBrief", nil})
    for i := range relationBriefs.RelationBrief {
        relationBriefs.EntityData.Children.Append(types.GetSegmentPath(relationBriefs.RelationBrief[i]), types.YChild{"RelationBrief", relationBriefs.RelationBrief[i]})
    }
    relationBriefs.EntityData.Leafs = types.NewOrderedMap()

    relationBriefs.EntityData.YListKeys = []string {}

    return &(relationBriefs.EntityData)
}

// Bfd_RelationBriefs_RelationBrief
// Brief information for relation of a single BFD
// session
type Bfd_RelationBriefs_RelationBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Brief Member Link Information. The type is slice of
    // Bfd_RelationBriefs_RelationBrief_LinkInformation.
    LinkInformation []*Bfd_RelationBriefs_RelationBrief_LinkInformation
}

func (relationBrief *Bfd_RelationBriefs_RelationBrief) GetEntityData() *types.CommonEntityData {
    relationBrief.EntityData.YFilter = relationBrief.YFilter
    relationBrief.EntityData.YangName = "relation-brief"
    relationBrief.EntityData.BundleName = "cisco_ios_xr"
    relationBrief.EntityData.ParentYangName = "relation-briefs"
    relationBrief.EntityData.SegmentPath = "relation-brief"
    relationBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relationBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relationBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relationBrief.EntityData.Children = types.NewOrderedMap()
    relationBrief.EntityData.Children.Append("link-information", types.YChild{"LinkInformation", nil})
    for i := range relationBrief.LinkInformation {
        relationBrief.EntityData.Children.Append(types.GetSegmentPath(relationBrief.LinkInformation[i]), types.YChild{"LinkInformation", relationBrief.LinkInformation[i]})
    }
    relationBrief.EntityData.Leafs = types.NewOrderedMap()
    relationBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", relationBrief.InterfaceName})
    relationBrief.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", relationBrief.DestinationAddress})
    relationBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", relationBrief.State})

    relationBrief.EntityData.YListKeys = []string {}

    return &(relationBrief.EntityData)
}

// Bfd_RelationBriefs_RelationBrief_LinkInformation
// Brief Member Link Information
type Bfd_RelationBriefs_RelationBrief_LinkInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}
}

func (linkInformation *Bfd_RelationBriefs_RelationBrief_LinkInformation) GetEntityData() *types.CommonEntityData {
    linkInformation.EntityData.YFilter = linkInformation.YFilter
    linkInformation.EntityData.YangName = "link-information"
    linkInformation.EntityData.BundleName = "cisco_ios_xr"
    linkInformation.EntityData.ParentYangName = "relation-brief"
    linkInformation.EntityData.SegmentPath = "link-information"
    linkInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkInformation.EntityData.Children = types.NewOrderedMap()
    linkInformation.EntityData.Leafs = types.NewOrderedMap()
    linkInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", linkInformation.State})
    linkInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", linkInformation.InterfaceName})

    linkInformation.EntityData.YListKeys = []string {}

    return &(linkInformation.EntityData)
}

// Bfd_ClientBriefs
// Table of Brief information about BFD clients
type Bfd_ClientBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information of client. The type is slice of
    // Bfd_ClientBriefs_ClientBrief.
    ClientBrief []*Bfd_ClientBriefs_ClientBrief
}

func (clientBriefs *Bfd_ClientBriefs) GetEntityData() *types.CommonEntityData {
    clientBriefs.EntityData.YFilter = clientBriefs.YFilter
    clientBriefs.EntityData.YangName = "client-briefs"
    clientBriefs.EntityData.BundleName = "cisco_ios_xr"
    clientBriefs.EntityData.ParentYangName = "bfd"
    clientBriefs.EntityData.SegmentPath = "client-briefs"
    clientBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientBriefs.EntityData.Children = types.NewOrderedMap()
    clientBriefs.EntityData.Children.Append("client-brief", types.YChild{"ClientBrief", nil})
    for i := range clientBriefs.ClientBrief {
        clientBriefs.EntityData.Children.Append(types.GetSegmentPath(clientBriefs.ClientBrief[i]), types.YChild{"ClientBrief", clientBriefs.ClientBrief[i]})
    }
    clientBriefs.EntityData.Leafs = types.NewOrderedMap()

    clientBriefs.EntityData.YListKeys = []string {}

    return &(clientBriefs.EntityData)
}

// Bfd_ClientBriefs_ClientBrief
// Brief information of client
type Bfd_ClientBriefs_ClientBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Client Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Client process name. The type is string with length: 0..257.
    NameXr interface{}

    // Location where client resides. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Number of sessions created by this client. The type is interface{} with
    // range: 0..4294967295.
    SessionCount interface{}
}

func (clientBrief *Bfd_ClientBriefs_ClientBrief) GetEntityData() *types.CommonEntityData {
    clientBrief.EntityData.YFilter = clientBrief.YFilter
    clientBrief.EntityData.YangName = "client-brief"
    clientBrief.EntityData.BundleName = "cisco_ios_xr"
    clientBrief.EntityData.ParentYangName = "client-briefs"
    clientBrief.EntityData.SegmentPath = "client-brief" + types.AddKeyToken(clientBrief.Name, "name")
    clientBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clientBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clientBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clientBrief.EntityData.Children = types.NewOrderedMap()
    clientBrief.EntityData.Leafs = types.NewOrderedMap()
    clientBrief.EntityData.Leafs.Append("name", types.YLeaf{"Name", clientBrief.Name})
    clientBrief.EntityData.Leafs.Append("name-xr", types.YLeaf{"NameXr", clientBrief.NameXr})
    clientBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", clientBrief.NodeId})
    clientBrief.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", clientBrief.SessionCount})

    clientBrief.EntityData.YListKeys = []string {"Name"}

    return &(clientBrief.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadMultiPaths
// IPv4 BFD over MPLS-TE Head multipath
type Bfd_Ipv4bfDoMplsteHeadMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label multipath table. The type is slice of
    // Bfd_Ipv4bfDoMplsteHeadMultiPaths_Ipv4bfDoMplsteHeadMultiPath.
    Ipv4bfDoMplsteHeadMultiPath []*Bfd_Ipv4bfDoMplsteHeadMultiPaths_Ipv4bfDoMplsteHeadMultiPath
}

func (ipv4bfDoMplsteHeadMultiPaths *Bfd_Ipv4bfDoMplsteHeadMultiPaths) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadMultiPaths.EntityData.YFilter = ipv4bfDoMplsteHeadMultiPaths.YFilter
    ipv4bfDoMplsteHeadMultiPaths.EntityData.YangName = "ipv4bf-do-mplste-head-multi-paths"
    ipv4bfDoMplsteHeadMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadMultiPaths.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteHeadMultiPaths.EntityData.SegmentPath = "ipv4bf-do-mplste-head-multi-paths"
    ipv4bfDoMplsteHeadMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadMultiPaths.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadMultiPaths.EntityData.Children.Append("ipv4bf-do-mplste-head-multi-path", types.YChild{"Ipv4bfDoMplsteHeadMultiPath", nil})
    for i := range ipv4bfDoMplsteHeadMultiPaths.Ipv4bfDoMplsteHeadMultiPath {
        ipv4bfDoMplsteHeadMultiPaths.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteHeadMultiPaths.Ipv4bfDoMplsteHeadMultiPath[i]), types.YChild{"Ipv4bfDoMplsteHeadMultiPath", ipv4bfDoMplsteHeadMultiPaths.Ipv4bfDoMplsteHeadMultiPath[i]})
    }
    ipv4bfDoMplsteHeadMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteHeadMultiPaths.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadMultiPaths.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadMultiPaths_Ipv4bfDoMplsteHeadMultiPath
// Label multipath table
type Bfd_Ipv4bfDoMplsteHeadMultiPaths_Ipv4bfDoMplsteHeadMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (ipv4bfDoMplsteHeadMultiPath *Bfd_Ipv4bfDoMplsteHeadMultiPaths_Ipv4bfDoMplsteHeadMultiPath) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadMultiPath.EntityData.YFilter = ipv4bfDoMplsteHeadMultiPath.YFilter
    ipv4bfDoMplsteHeadMultiPath.EntityData.YangName = "ipv4bf-do-mplste-head-multi-path"
    ipv4bfDoMplsteHeadMultiPath.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadMultiPath.EntityData.ParentYangName = "ipv4bf-do-mplste-head-multi-paths"
    ipv4bfDoMplsteHeadMultiPath.EntityData.SegmentPath = "ipv4bf-do-mplste-head-multi-path"
    ipv4bfDoMplsteHeadMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadMultiPath.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv4bfDoMplsteHeadMultiPath.InterfaceName})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteHeadMultiPath.VrfName})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteHeadMultiPath.IncomingLabel})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteHeadMultiPath.FeCtype})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteHeadMultiPath.FecSubgroupId})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteHeadMultiPath.Feclspid})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteHeadMultiPath.FecTunnelId})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteHeadMultiPath.FecExtendedTunnelId})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteHeadMultiPath.FecSource})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteHeadMultiPath.FecDestination})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteHeadMultiPath.Fecp2mpid})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteHeadMultiPath.FecSubgroupOriginator})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteHeadMultiPath.FecCtype})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteHeadMultiPath.Location})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4bfDoMplsteHeadMultiPath.SessionSubtype})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4bfDoMplsteHeadMultiPath.State})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", ipv4bfDoMplsteHeadMultiPath.LocalDiscriminator})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4bfDoMplsteHeadMultiPath.NodeId})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", ipv4bfDoMplsteHeadMultiPath.IncomingLabelXr})
    ipv4bfDoMplsteHeadMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", ipv4bfDoMplsteHeadMultiPath.SessionInterfaceName})

    ipv4bfDoMplsteHeadMultiPath.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadMultiPath.EntityData)
}

// Bfd_RelationDetails
// Table of detail information about all BFD
// relations in the System
type Bfd_RelationDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detail information for relation of a single BFD session. The type is slice
    // of Bfd_RelationDetails_RelationDetail.
    RelationDetail []*Bfd_RelationDetails_RelationDetail
}

func (relationDetails *Bfd_RelationDetails) GetEntityData() *types.CommonEntityData {
    relationDetails.EntityData.YFilter = relationDetails.YFilter
    relationDetails.EntityData.YangName = "relation-details"
    relationDetails.EntityData.BundleName = "cisco_ios_xr"
    relationDetails.EntityData.ParentYangName = "bfd"
    relationDetails.EntityData.SegmentPath = "relation-details"
    relationDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relationDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relationDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relationDetails.EntityData.Children = types.NewOrderedMap()
    relationDetails.EntityData.Children.Append("relation-detail", types.YChild{"RelationDetail", nil})
    for i := range relationDetails.RelationDetail {
        relationDetails.EntityData.Children.Append(types.GetSegmentPath(relationDetails.RelationDetail[i]), types.YChild{"RelationDetail", relationDetails.RelationDetail[i]})
    }
    relationDetails.EntityData.Leafs = types.NewOrderedMap()

    relationDetails.EntityData.YListKeys = []string {}

    return &(relationDetails.EntityData)
}

// Bfd_RelationDetails_RelationDetail
// Detail information for relation of a single BFD
// session
type Bfd_RelationDetails_RelationDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Detail Member Link Information. The type is slice of
    // Bfd_RelationDetails_RelationDetail_LinkInformation.
    LinkInformation []*Bfd_RelationDetails_RelationDetail_LinkInformation

    // Association session information. The type is slice of
    // Bfd_RelationDetails_RelationDetail_AssociationInformation.
    AssociationInformation []*Bfd_RelationDetails_RelationDetail_AssociationInformation
}

func (relationDetail *Bfd_RelationDetails_RelationDetail) GetEntityData() *types.CommonEntityData {
    relationDetail.EntityData.YFilter = relationDetail.YFilter
    relationDetail.EntityData.YangName = "relation-detail"
    relationDetail.EntityData.BundleName = "cisco_ios_xr"
    relationDetail.EntityData.ParentYangName = "relation-details"
    relationDetail.EntityData.SegmentPath = "relation-detail"
    relationDetail.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relationDetail.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relationDetail.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relationDetail.EntityData.Children = types.NewOrderedMap()
    relationDetail.EntityData.Children.Append("link-information", types.YChild{"LinkInformation", nil})
    for i := range relationDetail.LinkInformation {
        relationDetail.EntityData.Children.Append(types.GetSegmentPath(relationDetail.LinkInformation[i]), types.YChild{"LinkInformation", relationDetail.LinkInformation[i]})
    }
    relationDetail.EntityData.Children.Append("association-information", types.YChild{"AssociationInformation", nil})
    for i := range relationDetail.AssociationInformation {
        relationDetail.EntityData.Children.Append(types.GetSegmentPath(relationDetail.AssociationInformation[i]), types.YChild{"AssociationInformation", relationDetail.AssociationInformation[i]})
    }
    relationDetail.EntityData.Leafs = types.NewOrderedMap()
    relationDetail.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", relationDetail.InterfaceName})
    relationDetail.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", relationDetail.DestinationAddress})
    relationDetail.EntityData.Leafs.Append("state", types.YLeaf{"State", relationDetail.State})
    relationDetail.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", relationDetail.LocalDiscriminator})

    relationDetail.EntityData.YListKeys = []string {}

    return &(relationDetail.EntityData)
}

// Bfd_RelationDetails_RelationDetail_LinkInformation
// Detail Member Link Information
type Bfd_RelationDetails_RelationDetail_LinkInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}
}

func (linkInformation *Bfd_RelationDetails_RelationDetail_LinkInformation) GetEntityData() *types.CommonEntityData {
    linkInformation.EntityData.YFilter = linkInformation.YFilter
    linkInformation.EntityData.YangName = "link-information"
    linkInformation.EntityData.BundleName = "cisco_ios_xr"
    linkInformation.EntityData.ParentYangName = "relation-detail"
    linkInformation.EntityData.SegmentPath = "link-information"
    linkInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkInformation.EntityData.Children = types.NewOrderedMap()
    linkInformation.EntityData.Leafs = types.NewOrderedMap()
    linkInformation.EntityData.Leafs.Append("state", types.YLeaf{"State", linkInformation.State})
    linkInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", linkInformation.InterfaceName})
    linkInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", linkInformation.LocalDiscriminator})

    linkInformation.EntityData.YListKeys = []string {}

    return &(linkInformation.EntityData)
}

// Bfd_RelationDetails_RelationDetail_AssociationInformation
// Association session information
type Bfd_RelationDetails_RelationDetail_AssociationInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session Interface Name. The type is string with length: 0..64.
    InterfaceName interface{}

    // Session type. The type is BfdSession.
    Sessiontype interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // IPv4/v6 dest address.
    IpDestinationAddress Bfd_RelationDetails_RelationDetail_AssociationInformation_IpDestinationAddress

    // Client applications owning the session. The type is slice of
    // Bfd_RelationDetails_RelationDetail_AssociationInformation_OwnerInformation.
    OwnerInformation []*Bfd_RelationDetails_RelationDetail_AssociationInformation_OwnerInformation
}

func (associationInformation *Bfd_RelationDetails_RelationDetail_AssociationInformation) GetEntityData() *types.CommonEntityData {
    associationInformation.EntityData.YFilter = associationInformation.YFilter
    associationInformation.EntityData.YangName = "association-information"
    associationInformation.EntityData.BundleName = "cisco_ios_xr"
    associationInformation.EntityData.ParentYangName = "relation-detail"
    associationInformation.EntityData.SegmentPath = "association-information"
    associationInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInformation.EntityData.Children = types.NewOrderedMap()
    associationInformation.EntityData.Children.Append("ip-destination-address", types.YChild{"IpDestinationAddress", &associationInformation.IpDestinationAddress})
    associationInformation.EntityData.Children.Append("owner-information", types.YChild{"OwnerInformation", nil})
    for i := range associationInformation.OwnerInformation {
        associationInformation.EntityData.Children.Append(types.GetSegmentPath(associationInformation.OwnerInformation[i]), types.YChild{"OwnerInformation", associationInformation.OwnerInformation[i]})
    }
    associationInformation.EntityData.Leafs = types.NewOrderedMap()
    associationInformation.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", associationInformation.InterfaceName})
    associationInformation.EntityData.Leafs.Append("sessiontype", types.YLeaf{"Sessiontype", associationInformation.Sessiontype})
    associationInformation.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", associationInformation.LocalDiscriminator})

    associationInformation.EntityData.YListKeys = []string {}

    return &(associationInformation.EntityData)
}

// Bfd_RelationDetails_RelationDetail_AssociationInformation_IpDestinationAddress
// IPv4/v6 dest address
type Bfd_RelationDetails_RelationDetail_AssociationInformation_IpDestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipDestinationAddress *Bfd_RelationDetails_RelationDetail_AssociationInformation_IpDestinationAddress) GetEntityData() *types.CommonEntityData {
    ipDestinationAddress.EntityData.YFilter = ipDestinationAddress.YFilter
    ipDestinationAddress.EntityData.YangName = "ip-destination-address"
    ipDestinationAddress.EntityData.BundleName = "cisco_ios_xr"
    ipDestinationAddress.EntityData.ParentYangName = "association-information"
    ipDestinationAddress.EntityData.SegmentPath = "ip-destination-address"
    ipDestinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipDestinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipDestinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipDestinationAddress.EntityData.Children = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs = types.NewOrderedMap()
    ipDestinationAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipDestinationAddress.Afi})
    ipDestinationAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipDestinationAddress.Dummy})
    ipDestinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipDestinationAddress.Ipv4})
    ipDestinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipDestinationAddress.Ipv6})

    ipDestinationAddress.EntityData.YListKeys = []string {}

    return &(ipDestinationAddress.EntityData)
}

// Bfd_RelationDetails_RelationDetail_AssociationInformation_OwnerInformation
// Client applications owning the session
type Bfd_RelationDetails_RelationDetail_AssociationInformation_OwnerInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Client specified minimum transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    Interval interface{}

    // Client specified detection multiplier to compute detection time. The type
    // is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}

    // Adjusted minimum transmit interval in milli-seconds. The type is
    // interface{} with range: 0..4294967295. Units are millisecond.
    AdjustedInterval interface{}

    // Adjusted detection multiplier to compute detection time. The type is
    // interface{} with range: 0..4294967295.
    AdjustedDetectionMultiplier interface{}

    // Client process name. The type is string with length: 0..257.
    Name interface{}
}

func (ownerInformation *Bfd_RelationDetails_RelationDetail_AssociationInformation_OwnerInformation) GetEntityData() *types.CommonEntityData {
    ownerInformation.EntityData.YFilter = ownerInformation.YFilter
    ownerInformation.EntityData.YangName = "owner-information"
    ownerInformation.EntityData.BundleName = "cisco_ios_xr"
    ownerInformation.EntityData.ParentYangName = "association-information"
    ownerInformation.EntityData.SegmentPath = "owner-information"
    ownerInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ownerInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ownerInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ownerInformation.EntityData.Children = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs = types.NewOrderedMap()
    ownerInformation.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", ownerInformation.Interval})
    ownerInformation.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", ownerInformation.DetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("adjusted-interval", types.YLeaf{"AdjustedInterval", ownerInformation.AdjustedInterval})
    ownerInformation.EntityData.Leafs.Append("adjusted-detection-multiplier", types.YLeaf{"AdjustedDetectionMultiplier", ownerInformation.AdjustedDetectionMultiplier})
    ownerInformation.EntityData.Leafs.Append("name", types.YLeaf{"Name", ownerInformation.Name})

    ownerInformation.EntityData.YListKeys = []string {}

    return &(ownerInformation.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailCounters
// IPv4 BFD over MPLS-TE Counters
type Bfd_Ipv4bfDoMplsteTailCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of IPv4 BFD over MPLS-TE Packet counters.
    Ipv4bfDoMplsteTailPacketCounters Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters
}

func (ipv4bfDoMplsteTailCounters *Bfd_Ipv4bfDoMplsteTailCounters) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailCounters.EntityData.YFilter = ipv4bfDoMplsteTailCounters.YFilter
    ipv4bfDoMplsteTailCounters.EntityData.YangName = "ipv4bf-do-mplste-tail-counters"
    ipv4bfDoMplsteTailCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailCounters.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteTailCounters.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-counters"
    ipv4bfDoMplsteTailCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailCounters.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailCounters.EntityData.Children.Append("ipv4bf-do-mplste-tail-packet-counters", types.YChild{"Ipv4bfDoMplsteTailPacketCounters", &ipv4bfDoMplsteTailCounters.Ipv4bfDoMplsteTailPacketCounters})
    ipv4bfDoMplsteTailCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteTailCounters.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailCounters.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters
// Table of IPv4 BFD over MPLS-TE Packet counters
type Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface  IPv4 BFD over MPLS-TE Packet counters. The type is slice of
    // Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters_Ipv4bfDoMplsteTailPacketCounter.
    Ipv4bfDoMplsteTailPacketCounter []*Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters_Ipv4bfDoMplsteTailPacketCounter
}

func (ipv4bfDoMplsteTailPacketCounters *Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailPacketCounters.EntityData.YFilter = ipv4bfDoMplsteTailPacketCounters.YFilter
    ipv4bfDoMplsteTailPacketCounters.EntityData.YangName = "ipv4bf-do-mplste-tail-packet-counters"
    ipv4bfDoMplsteTailPacketCounters.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailPacketCounters.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-counters"
    ipv4bfDoMplsteTailPacketCounters.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-packet-counters"
    ipv4bfDoMplsteTailPacketCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailPacketCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailPacketCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailPacketCounters.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailPacketCounters.EntityData.Children.Append("ipv4bf-do-mplste-tail-packet-counter", types.YChild{"Ipv4bfDoMplsteTailPacketCounter", nil})
    for i := range ipv4bfDoMplsteTailPacketCounters.Ipv4bfDoMplsteTailPacketCounter {
        ipv4bfDoMplsteTailPacketCounters.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteTailPacketCounters.Ipv4bfDoMplsteTailPacketCounter[i]), types.YChild{"Ipv4bfDoMplsteTailPacketCounter", ipv4bfDoMplsteTailPacketCounters.Ipv4bfDoMplsteTailPacketCounter[i]})
    }
    ipv4bfDoMplsteTailPacketCounters.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteTailPacketCounters.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailPacketCounters.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters_Ipv4bfDoMplsteTailPacketCounter
// Interface  IPv4 BFD over MPLS-TE Packet
// counters
type Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters_Ipv4bfDoMplsteTailPacketCounter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Number of Hellos transmitted. The type is interface{} with range:
    // 0..4294967295.
    HelloTransmitCount interface{}

    // Number of Hellos received. The type is interface{} with range:
    // 0..4294967295.
    HelloReceiveCount interface{}

    // Number of echo packets transmitted. The type is interface{} with range:
    // 0..4294967295.
    EchoTransmitCount interface{}

    // Number of echo packets received. The type is interface{} with range:
    // 0..4294967295.
    EchoReceiveCount interface{}

    // Packet Display Type. The type is BfdMgmtPktDisplay.
    DisplayType interface{}
}

func (ipv4bfDoMplsteTailPacketCounter *Bfd_Ipv4bfDoMplsteTailCounters_Ipv4bfDoMplsteTailPacketCounters_Ipv4bfDoMplsteTailPacketCounter) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailPacketCounter.EntityData.YFilter = ipv4bfDoMplsteTailPacketCounter.YFilter
    ipv4bfDoMplsteTailPacketCounter.EntityData.YangName = "ipv4bf-do-mplste-tail-packet-counter"
    ipv4bfDoMplsteTailPacketCounter.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailPacketCounter.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-packet-counters"
    ipv4bfDoMplsteTailPacketCounter.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-packet-counter"
    ipv4bfDoMplsteTailPacketCounter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailPacketCounter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailPacketCounter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailPacketCounter.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteTailPacketCounter.VrfName})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteTailPacketCounter.IncomingLabel})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteTailPacketCounter.FeCtype})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteTailPacketCounter.FecSubgroupId})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteTailPacketCounter.Feclspid})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteTailPacketCounter.FecTunnelId})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteTailPacketCounter.FecExtendedTunnelId})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteTailPacketCounter.FecSource})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteTailPacketCounter.FecDestination})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteTailPacketCounter.Fecp2mpid})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteTailPacketCounter.FecSubgroupOriginator})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteTailPacketCounter.FecCtype})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteTailPacketCounter.Location})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("hello-transmit-count", types.YLeaf{"HelloTransmitCount", ipv4bfDoMplsteTailPacketCounter.HelloTransmitCount})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("hello-receive-count", types.YLeaf{"HelloReceiveCount", ipv4bfDoMplsteTailPacketCounter.HelloReceiveCount})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("echo-transmit-count", types.YLeaf{"EchoTransmitCount", ipv4bfDoMplsteTailPacketCounter.EchoTransmitCount})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("echo-receive-count", types.YLeaf{"EchoReceiveCount", ipv4bfDoMplsteTailPacketCounter.EchoReceiveCount})
    ipv4bfDoMplsteTailPacketCounter.EntityData.Leafs.Append("display-type", types.YLeaf{"DisplayType", ipv4bfDoMplsteTailPacketCounter.DisplayType})

    ipv4bfDoMplsteTailPacketCounter.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailPacketCounter.EntityData)
}

// Bfd_Ipv6SingleHopSessionBriefs
// Table of brief information about all IPv6
// singlehop BFD sessions in the System
type Bfd_Ipv6SingleHopSessionBriefs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information for a single IPv6 singlehop BFD session. The type is
    // slice of Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief.
    Ipv6SingleHopSessionBrief []*Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief
}

func (ipv6SingleHopSessionBriefs *Bfd_Ipv6SingleHopSessionBriefs) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopSessionBriefs.EntityData.YFilter = ipv6SingleHopSessionBriefs.YFilter
    ipv6SingleHopSessionBriefs.EntityData.YangName = "ipv6-single-hop-session-briefs"
    ipv6SingleHopSessionBriefs.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopSessionBriefs.EntityData.ParentYangName = "bfd"
    ipv6SingleHopSessionBriefs.EntityData.SegmentPath = "ipv6-single-hop-session-briefs"
    ipv6SingleHopSessionBriefs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopSessionBriefs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopSessionBriefs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopSessionBriefs.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopSessionBriefs.EntityData.Children.Append("ipv6-single-hop-session-brief", types.YChild{"Ipv6SingleHopSessionBrief", nil})
    for i := range ipv6SingleHopSessionBriefs.Ipv6SingleHopSessionBrief {
        ipv6SingleHopSessionBriefs.EntityData.Children.Append(types.GetSegmentPath(ipv6SingleHopSessionBriefs.Ipv6SingleHopSessionBrief[i]), types.YChild{"Ipv6SingleHopSessionBrief", ipv6SingleHopSessionBriefs.Ipv6SingleHopSessionBrief[i]})
    }
    ipv6SingleHopSessionBriefs.EntityData.Leafs = types.NewOrderedMap()

    ipv6SingleHopSessionBriefs.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopSessionBriefs.EntityData)
}

// Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief
// Brief information for a single IPv6 singlehop
// BFD session
type Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session type. The type is BfdSession.
    SessionType interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // Session Flags. The type is interface{} with range: 0..4294967295.
    SessionFlags interface{}

    // Brief Status Information.
    StatusBriefInformation Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation
}

func (ipv6SingleHopSessionBrief *Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief) GetEntityData() *types.CommonEntityData {
    ipv6SingleHopSessionBrief.EntityData.YFilter = ipv6SingleHopSessionBrief.YFilter
    ipv6SingleHopSessionBrief.EntityData.YangName = "ipv6-single-hop-session-brief"
    ipv6SingleHopSessionBrief.EntityData.BundleName = "cisco_ios_xr"
    ipv6SingleHopSessionBrief.EntityData.ParentYangName = "ipv6-single-hop-session-briefs"
    ipv6SingleHopSessionBrief.EntityData.SegmentPath = "ipv6-single-hop-session-brief"
    ipv6SingleHopSessionBrief.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6SingleHopSessionBrief.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6SingleHopSessionBrief.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6SingleHopSessionBrief.EntityData.Children = types.NewOrderedMap()
    ipv6SingleHopSessionBrief.EntityData.Children.Append("status-brief-information", types.YChild{"StatusBriefInformation", &ipv6SingleHopSessionBrief.StatusBriefInformation})
    ipv6SingleHopSessionBrief.EntityData.Leafs = types.NewOrderedMap()
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipv6SingleHopSessionBrief.InterfaceName})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv6SingleHopSessionBrief.DestinationAddress})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv6SingleHopSessionBrief.Location})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv6SingleHopSessionBrief.NodeId})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv6SingleHopSessionBrief.State})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("session-type", types.YLeaf{"SessionType", ipv6SingleHopSessionBrief.SessionType})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv6SingleHopSessionBrief.SessionSubtype})
    ipv6SingleHopSessionBrief.EntityData.Leafs.Append("session-flags", types.YLeaf{"SessionFlags", ipv6SingleHopSessionBrief.SessionFlags})

    ipv6SingleHopSessionBrief.EntityData.YListKeys = []string {}

    return &(ipv6SingleHopSessionBrief.EntityData)
}

// Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation
// Brief Status Information
type Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Async Interval and Detect Multiplier Information.
    AsyncIntervalMultiplier Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier

    // Echo Interval and Detect Multiplier Information.
    EchoIntervalMultiplier Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
}

func (statusBriefInformation *Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation) GetEntityData() *types.CommonEntityData {
    statusBriefInformation.EntityData.YFilter = statusBriefInformation.YFilter
    statusBriefInformation.EntityData.YangName = "status-brief-information"
    statusBriefInformation.EntityData.BundleName = "cisco_ios_xr"
    statusBriefInformation.EntityData.ParentYangName = "ipv6-single-hop-session-brief"
    statusBriefInformation.EntityData.SegmentPath = "status-brief-information"
    statusBriefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statusBriefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statusBriefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statusBriefInformation.EntityData.Children = types.NewOrderedMap()
    statusBriefInformation.EntityData.Children.Append("async-interval-multiplier", types.YChild{"AsyncIntervalMultiplier", &statusBriefInformation.AsyncIntervalMultiplier})
    statusBriefInformation.EntityData.Children.Append("echo-interval-multiplier", types.YChild{"EchoIntervalMultiplier", &statusBriefInformation.EchoIntervalMultiplier})
    statusBriefInformation.EntityData.Leafs = types.NewOrderedMap()

    statusBriefInformation.EntityData.YListKeys = []string {}

    return &(statusBriefInformation.EntityData)
}

// Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier
// Async Interval and Detect Multiplier Information
type Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated remote transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedRemoteTransmitInterval interface{}

    // Negotiated local transmit interval in micro-seconds. The type is
    // interface{} with range: 0..4294967295. Units are microsecond.
    NegotiatedLocalTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (asyncIntervalMultiplier *Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_AsyncIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    asyncIntervalMultiplier.EntityData.YFilter = asyncIntervalMultiplier.YFilter
    asyncIntervalMultiplier.EntityData.YangName = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    asyncIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    asyncIntervalMultiplier.EntityData.SegmentPath = "async-interval-multiplier"
    asyncIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-remote-transmit-interval", types.YLeaf{"NegotiatedRemoteTransmitInterval", asyncIntervalMultiplier.NegotiatedRemoteTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("negotiated-local-transmit-interval", types.YLeaf{"NegotiatedLocalTransmitInterval", asyncIntervalMultiplier.NegotiatedLocalTransmitInterval})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", asyncIntervalMultiplier.DetectionTime})
    asyncIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", asyncIntervalMultiplier.DetectionMultiplier})

    asyncIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(asyncIntervalMultiplier.EntityData)
}

// Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier
// Echo Interval and Detect Multiplier Information
type Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Negotiated transmit interval in micro-seconds. The type is interface{} with
    // range: 0..4294967295. Units are microsecond.
    NegotiatedTransmitInterval interface{}

    // Detection time in micro-seconds. The type is interface{} with range:
    // 0..4294967295. Units are microsecond.
    DetectionTime interface{}

    // Detection Multiplier. The type is interface{} with range: 0..4294967295.
    DetectionMultiplier interface{}
}

func (echoIntervalMultiplier *Bfd_Ipv6SingleHopSessionBriefs_Ipv6SingleHopSessionBrief_StatusBriefInformation_EchoIntervalMultiplier) GetEntityData() *types.CommonEntityData {
    echoIntervalMultiplier.EntityData.YFilter = echoIntervalMultiplier.YFilter
    echoIntervalMultiplier.EntityData.YangName = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.BundleName = "cisco_ios_xr"
    echoIntervalMultiplier.EntityData.ParentYangName = "status-brief-information"
    echoIntervalMultiplier.EntityData.SegmentPath = "echo-interval-multiplier"
    echoIntervalMultiplier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    echoIntervalMultiplier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    echoIntervalMultiplier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    echoIntervalMultiplier.EntityData.Children = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs = types.NewOrderedMap()
    echoIntervalMultiplier.EntityData.Leafs.Append("negotiated-transmit-interval", types.YLeaf{"NegotiatedTransmitInterval", echoIntervalMultiplier.NegotiatedTransmitInterval})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-time", types.YLeaf{"DetectionTime", echoIntervalMultiplier.DetectionTime})
    echoIntervalMultiplier.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", echoIntervalMultiplier.DetectionMultiplier})

    echoIntervalMultiplier.EntityData.YListKeys = []string {}

    return &(echoIntervalMultiplier.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailMultiPaths
// IPv4 BFD over MPLS-TE Tail multipath
type Bfd_Ipv4bfDoMplsteTailMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Label multipath table. The type is slice of
    // Bfd_Ipv4bfDoMplsteTailMultiPaths_Ipv4bfDoMplsteTailMultiPath.
    Ipv4bfDoMplsteTailMultiPath []*Bfd_Ipv4bfDoMplsteTailMultiPaths_Ipv4bfDoMplsteTailMultiPath
}

func (ipv4bfDoMplsteTailMultiPaths *Bfd_Ipv4bfDoMplsteTailMultiPaths) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailMultiPaths.EntityData.YFilter = ipv4bfDoMplsteTailMultiPaths.YFilter
    ipv4bfDoMplsteTailMultiPaths.EntityData.YangName = "ipv4bf-do-mplste-tail-multi-paths"
    ipv4bfDoMplsteTailMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailMultiPaths.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteTailMultiPaths.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-multi-paths"
    ipv4bfDoMplsteTailMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailMultiPaths.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailMultiPaths.EntityData.Children.Append("ipv4bf-do-mplste-tail-multi-path", types.YChild{"Ipv4bfDoMplsteTailMultiPath", nil})
    for i := range ipv4bfDoMplsteTailMultiPaths.Ipv4bfDoMplsteTailMultiPath {
        ipv4bfDoMplsteTailMultiPaths.EntityData.Children.Append(types.GetSegmentPath(ipv4bfDoMplsteTailMultiPaths.Ipv4bfDoMplsteTailMultiPath[i]), types.YChild{"Ipv4bfDoMplsteTailMultiPath", ipv4bfDoMplsteTailMultiPaths.Ipv4bfDoMplsteTailMultiPath[i]})
    }
    ipv4bfDoMplsteTailMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteTailMultiPaths.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailMultiPaths.EntityData)
}

// Bfd_Ipv4bfDoMplsteTailMultiPaths_Ipv4bfDoMplsteTailMultiPath
// Label multipath table
type Bfd_Ipv4bfDoMplsteTailMultiPaths_Ipv4bfDoMplsteTailMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabel interface{}

    // FEC Type. The type is interface{} with range: 0..4294967295.
    FeCtype interface{}

    // FEC Subgroup ID. The type is interface{} with range: 0..4294967295.
    FecSubgroupId interface{}

    // FEC LSP ID. The type is interface{} with range: 0..4294967295.
    Feclspid interface{}

    // FEC Tunnel ID. The type is interface{} with range: 0..4294967295.
    FecTunnelId interface{}

    // FEC Extended Tunnel ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecExtendedTunnelId interface{}

    // FEC Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSource interface{}

    // FEC Destination. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecDestination interface{}

    // FEC P2MP ID. The type is interface{} with range: 0..4294967295.
    Fecp2mpid interface{}

    // FEC Subgroup originator. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    FecSubgroupOriginator interface{}

    // FEC C Type. The type is interface{} with range: 0..4294967295.
    FecCtype interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (ipv4bfDoMplsteTailMultiPath *Bfd_Ipv4bfDoMplsteTailMultiPaths_Ipv4bfDoMplsteTailMultiPath) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteTailMultiPath.EntityData.YFilter = ipv4bfDoMplsteTailMultiPath.YFilter
    ipv4bfDoMplsteTailMultiPath.EntityData.YangName = "ipv4bf-do-mplste-tail-multi-path"
    ipv4bfDoMplsteTailMultiPath.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteTailMultiPath.EntityData.ParentYangName = "ipv4bf-do-mplste-tail-multi-paths"
    ipv4bfDoMplsteTailMultiPath.EntityData.SegmentPath = "ipv4bf-do-mplste-tail-multi-path"
    ipv4bfDoMplsteTailMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteTailMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteTailMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteTailMultiPath.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs = types.NewOrderedMap()
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4bfDoMplsteTailMultiPath.VrfName})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("incoming-label", types.YLeaf{"IncomingLabel", ipv4bfDoMplsteTailMultiPath.IncomingLabel})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fe-ctype", types.YLeaf{"FeCtype", ipv4bfDoMplsteTailMultiPath.FeCtype})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-subgroup-id", types.YLeaf{"FecSubgroupId", ipv4bfDoMplsteTailMultiPath.FecSubgroupId})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("feclspid", types.YLeaf{"Feclspid", ipv4bfDoMplsteTailMultiPath.Feclspid})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-tunnel-id", types.YLeaf{"FecTunnelId", ipv4bfDoMplsteTailMultiPath.FecTunnelId})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-extended-tunnel-id", types.YLeaf{"FecExtendedTunnelId", ipv4bfDoMplsteTailMultiPath.FecExtendedTunnelId})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-source", types.YLeaf{"FecSource", ipv4bfDoMplsteTailMultiPath.FecSource})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-destination", types.YLeaf{"FecDestination", ipv4bfDoMplsteTailMultiPath.FecDestination})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fecp2mpid", types.YLeaf{"Fecp2mpid", ipv4bfDoMplsteTailMultiPath.Fecp2mpid})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-subgroup-originator", types.YLeaf{"FecSubgroupOriginator", ipv4bfDoMplsteTailMultiPath.FecSubgroupOriginator})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("fec-ctype", types.YLeaf{"FecCtype", ipv4bfDoMplsteTailMultiPath.FecCtype})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4bfDoMplsteTailMultiPath.Location})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4bfDoMplsteTailMultiPath.SessionSubtype})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4bfDoMplsteTailMultiPath.State})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", ipv4bfDoMplsteTailMultiPath.LocalDiscriminator})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4bfDoMplsteTailMultiPath.NodeId})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", ipv4bfDoMplsteTailMultiPath.IncomingLabelXr})
    ipv4bfDoMplsteTailMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", ipv4bfDoMplsteTailMultiPath.SessionInterfaceName})

    ipv4bfDoMplsteTailMultiPath.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteTailMultiPath.EntityData)
}

// Bfd_Ipv4MultiHopMultiPaths
// IPv4 multi-hop multipath
type Bfd_Ipv4MultiHopMultiPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 multi hop multipath table. The type is slice of
    // Bfd_Ipv4MultiHopMultiPaths_Ipv4MultiHopMultiPath.
    Ipv4MultiHopMultiPath []*Bfd_Ipv4MultiHopMultiPaths_Ipv4MultiHopMultiPath
}

func (ipv4MultiHopMultiPaths *Bfd_Ipv4MultiHopMultiPaths) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopMultiPaths.EntityData.YFilter = ipv4MultiHopMultiPaths.YFilter
    ipv4MultiHopMultiPaths.EntityData.YangName = "ipv4-multi-hop-multi-paths"
    ipv4MultiHopMultiPaths.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopMultiPaths.EntityData.ParentYangName = "bfd"
    ipv4MultiHopMultiPaths.EntityData.SegmentPath = "ipv4-multi-hop-multi-paths"
    ipv4MultiHopMultiPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopMultiPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopMultiPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopMultiPaths.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopMultiPaths.EntityData.Children.Append("ipv4-multi-hop-multi-path", types.YChild{"Ipv4MultiHopMultiPath", nil})
    for i := range ipv4MultiHopMultiPaths.Ipv4MultiHopMultiPath {
        ipv4MultiHopMultiPaths.EntityData.Children.Append(types.GetSegmentPath(ipv4MultiHopMultiPaths.Ipv4MultiHopMultiPath[i]), types.YChild{"Ipv4MultiHopMultiPath", ipv4MultiHopMultiPaths.Ipv4MultiHopMultiPath[i]})
    }
    ipv4MultiHopMultiPaths.EntityData.Leafs = types.NewOrderedMap()

    ipv4MultiHopMultiPaths.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopMultiPaths.EntityData)
}

// Bfd_Ipv4MultiHopMultiPaths_Ipv4MultiHopMultiPath
// IPv4 multi hop multipath table
type Bfd_Ipv4MultiHopMultiPaths_Ipv4MultiHopMultiPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}

    // VRF name. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Session subtype. The type is string.
    SessionSubtype interface{}

    // State. The type is BfdMgmtSessionState.
    State interface{}

    // Session's Local discriminator. The type is interface{} with range:
    // 0..4294967295.
    LocalDiscriminator interface{}

    // Location where session is housed. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeId interface{}

    // Incoming Label. The type is interface{} with range: 0..4294967295.
    IncomingLabelXr interface{}

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
    SessionInterfaceName interface{}
}

func (ipv4MultiHopMultiPath *Bfd_Ipv4MultiHopMultiPaths_Ipv4MultiHopMultiPath) GetEntityData() *types.CommonEntityData {
    ipv4MultiHopMultiPath.EntityData.YFilter = ipv4MultiHopMultiPath.YFilter
    ipv4MultiHopMultiPath.EntityData.YangName = "ipv4-multi-hop-multi-path"
    ipv4MultiHopMultiPath.EntityData.BundleName = "cisco_ios_xr"
    ipv4MultiHopMultiPath.EntityData.ParentYangName = "ipv4-multi-hop-multi-paths"
    ipv4MultiHopMultiPath.EntityData.SegmentPath = "ipv4-multi-hop-multi-path"
    ipv4MultiHopMultiPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4MultiHopMultiPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4MultiHopMultiPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4MultiHopMultiPath.EntityData.Children = types.NewOrderedMap()
    ipv4MultiHopMultiPath.EntityData.Leafs = types.NewOrderedMap()
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", ipv4MultiHopMultiPath.SourceAddress})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", ipv4MultiHopMultiPath.DestinationAddress})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("location", types.YLeaf{"Location", ipv4MultiHopMultiPath.Location})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", ipv4MultiHopMultiPath.VrfName})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("session-subtype", types.YLeaf{"SessionSubtype", ipv4MultiHopMultiPath.SessionSubtype})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("state", types.YLeaf{"State", ipv4MultiHopMultiPath.State})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", ipv4MultiHopMultiPath.LocalDiscriminator})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("node-id", types.YLeaf{"NodeId", ipv4MultiHopMultiPath.NodeId})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("incoming-label-xr", types.YLeaf{"IncomingLabelXr", ipv4MultiHopMultiPath.IncomingLabelXr})
    ipv4MultiHopMultiPath.EntityData.Leafs.Append("session-interface-name", types.YLeaf{"SessionInterfaceName", ipv4MultiHopMultiPath.SessionInterfaceName})

    ipv4MultiHopMultiPath.EntityData.YListKeys = []string {}

    return &(ipv4MultiHopMultiPath.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSummary
// Summary information of IPv4 BFD over MPLS-TE
// Head
type Bfd_Ipv4bfDoMplsteHeadSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Statistics of states for sessions.
    SessionState Bfd_Ipv4bfDoMplsteHeadSummary_SessionState
}

func (ipv4bfDoMplsteHeadSummary *Bfd_Ipv4bfDoMplsteHeadSummary) GetEntityData() *types.CommonEntityData {
    ipv4bfDoMplsteHeadSummary.EntityData.YFilter = ipv4bfDoMplsteHeadSummary.YFilter
    ipv4bfDoMplsteHeadSummary.EntityData.YangName = "ipv4bf-do-mplste-head-summary"
    ipv4bfDoMplsteHeadSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4bfDoMplsteHeadSummary.EntityData.ParentYangName = "bfd"
    ipv4bfDoMplsteHeadSummary.EntityData.SegmentPath = "ipv4bf-do-mplste-head-summary"
    ipv4bfDoMplsteHeadSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4bfDoMplsteHeadSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4bfDoMplsteHeadSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4bfDoMplsteHeadSummary.EntityData.Children = types.NewOrderedMap()
    ipv4bfDoMplsteHeadSummary.EntityData.Children.Append("session-state", types.YChild{"SessionState", &ipv4bfDoMplsteHeadSummary.SessionState})
    ipv4bfDoMplsteHeadSummary.EntityData.Leafs = types.NewOrderedMap()

    ipv4bfDoMplsteHeadSummary.EntityData.YListKeys = []string {}

    return &(ipv4bfDoMplsteHeadSummary.EntityData)
}

// Bfd_Ipv4bfDoMplsteHeadSummary_SessionState
// Statistics of states for sessions
type Bfd_Ipv4bfDoMplsteHeadSummary_SessionState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of sessions in database. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Number of sessions in down state. The type is interface{} with range:
    // 0..4294967295.
    DownCount interface{}

    // Number of sessions in up state. The type is interface{} with range:
    // 0..4294967295.
    UpCount interface{}

    // Number of sessions in unknown state. The type is interface{} with range:
    // 0..4294967295.
    UnknownCount interface{}
}

func (sessionState *Bfd_Ipv4bfDoMplsteHeadSummary_SessionState) GetEntityData() *types.CommonEntityData {
    sessionState.EntityData.YFilter = sessionState.YFilter
    sessionState.EntityData.YangName = "session-state"
    sessionState.EntityData.BundleName = "cisco_ios_xr"
    sessionState.EntityData.ParentYangName = "ipv4bf-do-mplste-head-summary"
    sessionState.EntityData.SegmentPath = "session-state"
    sessionState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionState.EntityData.Children = types.NewOrderedMap()
    sessionState.EntityData.Leafs = types.NewOrderedMap()
    sessionState.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", sessionState.TotalCount})
    sessionState.EntityData.Leafs.Append("down-count", types.YLeaf{"DownCount", sessionState.DownCount})
    sessionState.EntityData.Leafs.Append("up-count", types.YLeaf{"UpCount", sessionState.UpCount})
    sessionState.EntityData.Leafs.Append("unknown-count", types.YLeaf{"UnknownCount", sessionState.UnknownCount})

    sessionState.EntityData.YListKeys = []string {}

    return &(sessionState.EntityData)
}

