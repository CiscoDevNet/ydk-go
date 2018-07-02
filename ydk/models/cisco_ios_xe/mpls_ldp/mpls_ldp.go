// This module contains a collection of YANG definitions
// for the Cisco MPLS LDP configuration and operational data.
// 
// 
// The configuration is held in the mpls-ldp-config container
// which is broken into the following sections:
//   1) global-cfg contains configuration applicable to the entire
//        LSR.
//   2) nbr-table contains configuration for specific LDP neighbors
//        or peers.
//   3) passwords contains configuration regarding passwords, both
//        local and those to be used with specific neighbors.
//   4) label-cfg contains the label allocation and advertisement
//        configuration and filters.
//   5) discovery contains the configuration for link-hello and
//        targetted hello protocol parameters including
//        interface-specific settings for transport.
//   6) graceful-restart contains the configuration for the
//        graceful restart feature.
//   7) logging contains the configuration for ldp-specific logs.
//   8) interfaces contains the configuration for each interface,
//        including and routing interactions specific to that
//        interface.
// 
// The operational state is held in the mpls-ldp-state container
// which is broken the following sections:
//   1) oper-summary contains the summarized global state.
//   2) forwarding-summary contains the summarized forwarding
//        state.
//   3) bindings-summary contains the summarized forwarding state.
//   4) vrf provides the detailed state on a per VRF basis.
//   5) bindings-advertise-specs - holds the advertisement
//        specification filters
//   6) discovery provides the LDP Discovery operational state.
//   7) forwarding provides summary information regarding LDP
//        forwarding setup and detailed information on the LDP
//        forwarding rewrites
//   8) bindings provides the detailed LDP Bindings of address to
//        label.
//   9) neighbors
// 
// The vrf-table, provides the detailed state on a per VRF basis.
// If the router only supports LDP in a single VRF then this table
// will have a single entry using the vrf-name 'default'.
// Otherwise this table will have one entry for every VRF where
// LDP is enabled on the device.
// 
// Each vrf includes:
//    A list parameters used by the VRF
//    A capability table containing the capabilities exchanged with
//      each neighbor.
//    A table of backoff parameters used in this VRF.
//    The graceful restart information used between the local
//       device and the neighbors should any of them restart.
//    An AF-table which holds all information for a given Address
//       Family. This is extensive and is described below.
//    The LDP ID used by the device for this vrf.
// 
// The AF-table holds information for a given Address Family
// such as:
//      - per-interface state.
//      - IGP synchronization data.
//      - LDP bindings statistics.
//      - LDP forwarding statistics.
// 
// 
//   Terms and Acronyms
// 
//   FRR - Fast Re-Reroute
// 
//   ICCP - Inter-Chassis Communication Protocol
// 
//   LACP - Link Aggregation Control Protocol
// 
//   LDP - Label Distribution Protocol
// 
//   LER - Label Edge Router
// 
//   LFA - Loop Free Alternative
// 
//   LIB - Label Information Base
// 
//   LSR - Label Switch Router
// 
//   MPLS - Multi-Protocol Label Switching
// 
//   PQ node - A node which is a member of both the extended
//       P-space and the Q-space as defined in
//       draft-ietf-rtgwg-rlfa-node-protection.
// 
//   VRF - Virtual Route Forwarding
// 
// Copyright (c) 2014, 2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_ldp

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_ldp"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mpls-ldp mpls-ldp}", reflect.TypeOf(MplsLdp{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mpls-ldp:mpls-ldp", reflect.TypeOf(MplsLdp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mpls-ldp clear-msg-counters}", reflect.TypeOf(ClearMsgCounters{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mpls-ldp:clear-msg-counters", reflect.TypeOf(ClearMsgCounters{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mpls-ldp restart-neighbor}", reflect.TypeOf(RestartNeighbor{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mpls-ldp:restart-neighbor", reflect.TypeOf(RestartNeighbor{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-mpls-ldp clear-forwarding}", reflect.TypeOf(ClearForwarding{}))
    ydk.RegisterEntity("Cisco-IOS-XE-mpls-ldp:clear-forwarding", reflect.TypeOf(ClearForwarding{}))
}

type DownNbrReasonDiscHello struct {
}

func (id DownNbrReasonDiscHello) String() string {
	return "Cisco-IOS-XE-mpls-ldp:down-nbr-reason-disc-hello"
}

type NsrPeerSyncErrAppFail struct {
}

func (id NsrPeerSyncErrAppFail) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-app-fail"
}

type RoutePathLblOwnerBgp struct {
}

func (id RoutePathLblOwnerBgp) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-lbl-owner-bgp"
}

type NsrSyncNackRsnErrAppNotFound struct {
}

func (id NsrSyncNackRsnErrAppNotFound) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-app-not-found"
}

type LdpNsrPeerSyncStOper struct {
}

func (id LdpNsrPeerSyncStOper) String() string {
	return "Cisco-IOS-XE-mpls-ldp:ldp-nsr-peer-sync-st-oper"
}

type IgpSyncDownReasonNa struct {
}

func (id IgpSyncDownReasonNa) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason-na"
}

type NsrSyncNackRsnNoPEndSock struct {
}

func (id NsrSyncNackRsnNoPEndSock) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-no-p-end-sock"
}

type NsrPeerSyncErrLdpGbl struct {
}

func (id NsrPeerSyncErrLdpGbl) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-ldp-gbl"
}

type DownNbrReasonNbrHold struct {
}

func (id DownNbrReasonNbrHold) String() string {
	return "Cisco-IOS-XE-mpls-ldp:down-nbr-reason-nbr-hold"
}

type RoutePathLblOwner struct {
}

func (id RoutePathLblOwner) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-lbl-owner"
}

type IgpSyncDownReason struct {
}

func (id IgpSyncDownReason) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason"
}

type NsrPeerSyncErrTcpGbl struct {
}

func (id NsrPeerSyncErrTcpGbl) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-tcp-gbl"
}

type IgpSyncDownReasonPeerUpdateNotReceived struct {
}

func (id IgpSyncDownReasonPeerUpdateNotReceived) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason-peer-update-not-received"
}

type IcpmTypeIccp struct {
}

func (id IcpmTypeIccp) String() string {
	return "Cisco-IOS-XE-mpls-ldp:icpm-type-iccp"
}

type NsrStatusReady struct {
}

func (id NsrStatusReady) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-status-ready"
}

type NsrSyncNackRsnErrAdjAdd struct {
}

func (id NsrSyncNackRsnErrAdjAdd) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-adj-add"
}

type NsrSyncNackRsnErrUnexpPeerDown struct {
}

func (id NsrSyncNackRsnErrUnexpPeerDown) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-unexp-peer-down"
}

type NsrSyncNackRsnNone struct {
}

func (id NsrSyncNackRsnNone) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-none"
}

type NsrSyncNackRsnNoCtx struct {
}

func (id NsrSyncNackRsnNoCtx) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-no-ctx"
}

type NsrSyncNackRsnMissingElem struct {
}

func (id NsrSyncNackRsnMissingElem) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-missing-elem"
}

type DownNbrReason struct {
}

func (id DownNbrReason) String() string {
	return "Cisco-IOS-XE-mpls-ldp:down-nbr-reason"
}

type NsrPeerSyncErrLdpPeer struct {
}

func (id NsrPeerSyncErrLdpPeer) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-ldp-peer"
}

type NsrSyncNackRsnErrAppInvalid struct {
}

func (id NsrSyncNackRsnErrAppInvalid) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-app-invalid"
}

type NsrStatusDisabled struct {
}

func (id NsrStatusDisabled) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-status-disabled"
}

type RoutePathIpNoFlag struct {
}

func (id RoutePathIpNoFlag) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-ip-no-flag"
}

type NsrPeerSyncErrTcpPeer struct {
}

func (id NsrPeerSyncErrTcpPeer) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-tcp-peer"
}

type NsrPeerSyncErrNone struct {
}

func (id NsrPeerSyncErrNone) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-none"
}

type LdpNsrPeerSyncStNone struct {
}

func (id LdpNsrPeerSyncStNone) String() string {
	return "Cisco-IOS-XE-mpls-ldp:ldp-nsr-peer-sync-st-none"
}

type NsrSyncNackRsnErrPpCreate struct {
}

func (id NsrSyncNackRsnErrPpCreate) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-pp-create"
}

type NsrSyncNackRsnErrRxUnexpOpen struct {
}

func (id NsrSyncNackRsnErrRxUnexpOpen) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-rx-unexp-open"
}

type LabelTypeUnknown struct {
}

func (id LabelTypeUnknown) String() string {
	return "Cisco-IOS-XE-mpls-ldp:label-type-unknown"
}

type IgpSyncDownReasonNoPeerSess struct {
}

func (id IgpSyncDownReasonNoPeerSess) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason-no-peer-sess"
}

type RoutePathType struct {
}

func (id RoutePathType) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-type"
}

type LdpNsrPeerSyncStPrep struct {
}

func (id LdpNsrPeerSyncStPrep) String() string {
	return "Cisco-IOS-XE-mpls-ldp:ldp-nsr-peer-sync-st-prep"
}

type RoutePathLblOwnerStatic struct {
}

func (id RoutePathLblOwnerStatic) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-lbl-owner-static"
}

type IcpmType struct {
}

func (id IcpmType) String() string {
	return "Cisco-IOS-XE-mpls-ldp:icpm-type"
}

type NsrPeerSyncErrLdpSyncNack struct {
}

func (id NsrPeerSyncErrLdpSyncNack) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-ldp-sync-nack"
}

type NsrPeerSyncErr struct {
}

func (id NsrPeerSyncErr) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err"
}

type NsrSyncNackRsnTblIdMismatch struct {
}

func (id NsrSyncNackRsnTblIdMismatch) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-tbl-id-mismatch"
}

type NsrSyncNackRsnErrAddrBind struct {
}

func (id NsrSyncNackRsnErrAddrBind) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-addr-bind"
}

type RoutePathIpBackup struct {
}

func (id RoutePathIpBackup) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-ip-backup"
}

type LabelTypeMpls struct {
}

func (id LabelTypeMpls) String() string {
	return "Cisco-IOS-XE-mpls-ldp:label-type-mpls"
}

type NsrSyncNackRsnErrDhcAdd struct {
}

func (id NsrSyncNackRsnErrDhcAdd) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-dhc-add"
}

type NsrSyncNackRsnErrRxNotif struct {
}

func (id NsrSyncNackRsnErrRxNotif) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-rx-notif"
}

type RoutePathLblOwnerNone struct {
}

func (id RoutePathLblOwnerNone) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-lbl-owner-none"
}

type NsrPeerSyncState struct {
}

func (id NsrPeerSyncState) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-state"
}

type LdpNsrPeerSyncStWait struct {
}

func (id LdpNsrPeerSyncStWait) String() string {
	return "Cisco-IOS-XE-mpls-ldp:ldp-nsr-peer-sync-st-wait"
}

type IccpType struct {
}

func (id IccpType) String() string {
	return "Cisco-IOS-XE-mpls-ldp:iccp-type"
}

type NsrSyncNackRsnErrTpCreate struct {
}

func (id NsrSyncNackRsnErrTpCreate) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-tp-create"
}

type LabelType struct {
}

func (id LabelType) String() string {
	return "Cisco-IOS-XE-mpls-ldp:label-type"
}

type NsrStatusNotReady struct {
}

func (id NsrStatusNotReady) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-status-not-ready"
}

type NsrSyncNackRsnPEndSockNotSynced struct {
}

func (id NsrSyncNackRsnPEndSockNotSynced) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-p-end-sock-not-synced"
}

type IgpSyncDownReasonInternal struct {
}

func (id IgpSyncDownReasonInternal) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason-internal"
}

type NsrSyncNackRsnPpExists struct {
}

func (id NsrSyncNackRsnPpExists) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-pp-exists"
}

type LdpNsrPeerSyncStReady struct {
}

func (id LdpNsrPeerSyncStReady) String() string {
	return "Cisco-IOS-XE-mpls-ldp:ldp-nsr-peer-sync-st-ready"
}

type IgpSyncDownReasonNoHelloAdj struct {
}

func (id IgpSyncDownReasonNoHelloAdj) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason-no-hello-adj"
}

type LabelTypeUnLabeled struct {
}

func (id LabelTypeUnLabeled) String() string {
	return "Cisco-IOS-XE-mpls-ldp:label-type-un-labeled"
}

type IccpTypeMlacp struct {
}

func (id IccpTypeMlacp) String() string {
	return "Cisco-IOS-XE-mpls-ldp:iccp-type-mlacp"
}

type NsrPeerSyncErrSyncPrep struct {
}

func (id NsrPeerSyncErrSyncPrep) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-peer-sync-err-sync-prep"
}

type LdpNsrPeerSyncStAppWait struct {
}

func (id LdpNsrPeerSyncStAppWait) String() string {
	return "Cisco-IOS-XE-mpls-ldp:ldp-nsr-peer-sync-st-app-wait"
}

type RoutePathIpBgpBackup struct {
}

func (id RoutePathIpBgpBackup) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-ip-bgp-backup"
}

type DownNbrReasonNa struct {
}

func (id DownNbrReasonNa) String() string {
	return "Cisco-IOS-XE-mpls-ldp:down-nbr-reason-na"
}

type NsrStatus struct {
}

func (id NsrStatus) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-status"
}

type NsrSyncNackRsnErrRxBadPie struct {
}

func (id NsrSyncNackRsnErrRxBadPie) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-err-rx-bad-pie"
}

type IgpSyncDownReasonPeerUpdateNotDone struct {
}

func (id IgpSyncDownReasonPeerUpdateNotDone) String() string {
	return "Cisco-IOS-XE-mpls-ldp:igp-sync-down-reason-peer-update-not-done"
}

type NsrSyncNackRsn struct {
}

func (id NsrSyncNackRsn) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn"
}

type NsrSyncNackRsnEnomem struct {
}

func (id NsrSyncNackRsnEnomem) String() string {
	return "Cisco-IOS-XE-mpls-ldp:nsr-sync-nack-rsn-enomem"
}

type RoutePathIpProtected struct {
}

func (id RoutePathIpProtected) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-ip-protected"
}

type RoutePathIpBackupRemote struct {
}

func (id RoutePathIpBackupRemote) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-ip-backup-remote"
}

type RoutePathLblOwnerLdp struct {
}

func (id RoutePathLblOwnerLdp) String() string {
	return "Cisco-IOS-XE-mpls-ldp:route-path-lbl-owner-ldp"
}

// SessionState represents for session negotiation behavior.
type SessionState string

const (
    // LDP session state: nonexistent.
    SessionState_nonexistent SessionState = "nonexistent"

    // LDP session state: initialized.
    SessionState_initialized SessionState = "initialized"

    // LDP session state: openrec.
    SessionState_openrec SessionState = "openrec"

    // LDP session state: opensent.
    SessionState_opensent SessionState = "opensent"

    // LDP session state: operational.
    SessionState_operational SessionState = "operational"
)

// LoopDetectionType represents the LSR or enabled on the LSR.
type LoopDetectionType string

const (
    // Loop Detection is not enabled on this LSR.
    LoopDetectionType_none LoopDetectionType = "none"

    // Loop Detection is enabled but by a method
    // other than those defined.
    LoopDetectionType_other LoopDetectionType = "other"

    // Loop Detection is supported by Hop Count only.
    LoopDetectionType_hop_count LoopDetectionType = "hop-count"

    // Loop Detection is supported by Path Vector only.
    LoopDetectionType_path_vector LoopDetectionType = "path-vector"

    // Loop Detection is supported by both Hop Count
    // and Path Vector.
    LoopDetectionType_hop_count_and_path_vector LoopDetectionType = "hop-count-and-path-vector"
)

// AdvLabelType represents advertise for matching prefixes and peers.
type AdvLabelType string

const (
    // Advertise the label for matching prefixes and peers.
    AdvLabelType_use_lable AdvLabelType = "use-lable"

    // Advertise explicit null for matching prefixes and peers.
    AdvLabelType_use_explicit AdvLabelType = "use-explicit"

    // Advertise imlicit null for matching prefixes and peers.
    AdvLabelType_use_implicit AdvLabelType = "use-implicit"

    // Do not advertise labels for matching prefixes and peers.
    AdvLabelType_none AdvLabelType = "none"
)

// AfId represents LDP AF type
type AfId string

const (
    // No Address Family
    AfId_ldp_af_id_none AfId = "ldp-af-id-none"

    // IPv4 AFI
    AfId_ldp_af_id_ipv4 AfId = "ldp-af-id-ipv4"

    // IPv6 AFI
    AfId_ldp_af_id_ipv6 AfId = "ldp-af-id-ipv6"
)

// IgpSyncState represents This is the IGP Synchronization State.
type IgpSyncState string

const (
    // Achieved
    IgpSyncState_isync_ready IgpSyncState = "isync-ready"

    // Not achieved
    IgpSyncState_isync_not_ready IgpSyncState = "isync-not-ready"

    // Deferred due to interface delay or global
    // restart delay
    IgpSyncState_isync_deferred IgpSyncState = "isync-deferred"
)

// Af represents LDP Address Family
type Af string

const (
    // No Address Family
    Af_ldp_af_none Af = "ldp-af-none"

    // IPv4 AFI
    Af_ldp_af_ipv4 Af = "ldp-af-ipv4"

    // IPv6 AFI
    Af_ldp_af_ipv6 Af = "ldp-af-ipv6"

    // Both IPv4/IPv6 AFIs
    Af_ldp_af_ipv4_ipv6 Af = "ldp-af-ipv4-ipv6"
)

// NbrBgpAdvtState represents Type.
type NbrBgpAdvtState string

const (
    // BGP Label Advertisement is not applicable.
    NbrBgpAdvtState_not_applicable NbrBgpAdvtState = "not-applicable"

    // BGP Label Advertisement is permitted.
    NbrBgpAdvtState_permit NbrBgpAdvtState = "permit"

    // BGP Label Advertisement denied.
    NbrBgpAdvtState_deny NbrBgpAdvtState = "deny"
)

// AdjState represents for LDP adjacency peer.
type AdjState string

const (
    // LDP adjacency state: nonexistent.
    AdjState_nonex AdjState = "nonex"

    // LDP session state: unsolicited open pending.
    AdjState_unsol_op_pdg AdjState = "unsol-op-pdg"

    // LDP session state: deferred.
    AdjState_deferred AdjState = "deferred"

    // LDP session state: established
    AdjState_estab AdjState = "estab"

    // LDP session state: LIB expension wait.
    AdjState_lib_exp_wait AdjState = "lib-exp-wait"

    // LDP session state: destroyed.
    AdjState_destroyed AdjState = "destroyed"
)

// LocalLabelState represents This id the MPLS LDP Local Label State Type.
type LocalLabelState string

const (
    // None
    LocalLabelState_local_label_state_none LocalLabelState = "local-label-state-none"

    // Assigned
    LocalLabelState_local_label_state_assigned LocalLabelState = "local-label-state-assigned"

    // Withdrawn
    LocalLabelState_local_label_state_withdrawn LocalLabelState = "local-label-state-withdrawn"
)

// DhcState represents This is the Directed Hello Control State Type.
type DhcState string

const (
    // There is no current Directed Hello Control State.
    DhcState_none DhcState = "none"

    // The Directed Hello is Active.
    DhcState_dhc_active DhcState = "dhc-active"

    // The Directed Hello is Passive.
    DhcState_dhc_passive DhcState = "dhc-passive"

    // The Directed Hello is both Active and Passive.
    DhcState_dhc_active_passive DhcState = "dhc-active-passive"
)

// IccpState represents IETF in TBD.
type IccpState string

const (
    // This state is the starting point for the state machine.
    // It indicates that no ICCP connection exists and that
    // there's no LDP session established between the PEs.
    IccpState_nonexistent IccpState = "nonexistent"

    // This state indicates that an LDP session exists between
    // the PEs but LDP ICCP Capabilitiy have not yet been
    // exchanged between them.
    IccpState_initialized IccpState = "initialized"

    // This state indicates that an LDP session exists between
    // the PEs and that the local PE has avertized LDP ICCP
    // Capability to its peer.
    IccpState_capsent IccpState = "capsent"

    // This state indicates that an LDP session exists between
    // the PEs and that the local PE has both received and
    // advertized LDP ICCP Capability from/to its peer.
    IccpState_caprec IccpState = "caprec"

    // This state indicates that the local PE has initiated an
    // ICCP connection to its peer, and is awaiting its
    // response.
    IccpState_connecting IccpState = "connecting"

    // This state indicates that the ICCP connection is
    // operational.
    IccpState_operational IccpState = "operational"
)

// MplsLdp
// MPLS LDP configuration and operational data.
type MplsLdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP operational data.
    MplsLdpState MplsLdp_MplsLdpState

    // MPLS LDP Configuration.
    MplsLdpConfig MplsLdp_MplsLdpConfig
}

func (mplsLdp *MplsLdp) GetEntityData() *types.CommonEntityData {
    mplsLdp.EntityData.YFilter = mplsLdp.YFilter
    mplsLdp.EntityData.YangName = "mpls-ldp"
    mplsLdp.EntityData.BundleName = "cisco_ios_xe"
    mplsLdp.EntityData.ParentYangName = "Cisco-IOS-XE-mpls-ldp"
    mplsLdp.EntityData.SegmentPath = "Cisco-IOS-XE-mpls-ldp:mpls-ldp"
    mplsLdp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdp.EntityData.Children = types.NewOrderedMap()
    mplsLdp.EntityData.Children.Append("mpls-ldp-state", types.YChild{"MplsLdpState", &mplsLdp.MplsLdpState})
    mplsLdp.EntityData.Children.Append("mpls-ldp-config", types.YChild{"MplsLdpConfig", &mplsLdp.MplsLdpConfig})
    mplsLdp.EntityData.Leafs = types.NewOrderedMap()

    mplsLdp.EntityData.YListKeys = []string {}

    return &(mplsLdp.EntityData)
}

// MplsLdp_MplsLdpState
// MPLS LDP operational data.
type MplsLdp_MplsLdpState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP operational data summary.
    OperSummary MplsLdp_MplsLdpState_OperSummary

    // Summary information regarding LDP forwarding setup.
    ForwardingSummary MplsLdp_MplsLdpState_ForwardingSummary

    // Aggregate counters for the MPLS LDP LIB.
    BindingsSummary MplsLdp_MplsLdpState_BindingsSummary

    // This is the LDP NSR summary for the device.
    NsrSummaryAll MplsLdp_MplsLdpState_NsrSummaryAll

    // Summary info for LDP ICPM/ICCP.
    IcpmSummaryAll MplsLdp_MplsLdpState_IcpmSummaryAll

    // MPLS LDP Global Parameters.
    Parameters MplsLdp_MplsLdpState_Parameters

    // LDP capability database information.
    Capabilities MplsLdp_MplsLdpState_Capabilities

    // MPLS LDP Session Backoff Information.
    BackoffParameters MplsLdp_MplsLdpState_BackoffParameters

    // MPLS LDP Graceful Restart Information.
    GracefulRestart MplsLdp_MplsLdpState_GracefulRestart

    // MPLS LDP per-VRF operational data.
    Vrfs MplsLdp_MplsLdpState_Vrfs

    // The LDP Discovery operational state.
    Discovery MplsLdp_MplsLdpState_Discovery

    // Summary information regarding LDP forwarding setup and detailed LDP
    // Forwarding rewrites.
    Forwarding MplsLdp_MplsLdpState_Forwarding

    // The detailed LDP Bindings.
    Bindings MplsLdp_MplsLdpState_Bindings

    // The LDP Neighbors Information.
    Neighbors MplsLdp_MplsLdpState_Neighbors

    // This contaions holds all the label ranges in use by this LDP instance.
    LabelRanges MplsLdp_MplsLdpState_LabelRanges
}

func (mplsLdpState *MplsLdp_MplsLdpState) GetEntityData() *types.CommonEntityData {
    mplsLdpState.EntityData.YFilter = mplsLdpState.YFilter
    mplsLdpState.EntityData.YangName = "mpls-ldp-state"
    mplsLdpState.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpState.EntityData.ParentYangName = "mpls-ldp"
    mplsLdpState.EntityData.SegmentPath = "mpls-ldp-state"
    mplsLdpState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpState.EntityData.Children = types.NewOrderedMap()
    mplsLdpState.EntityData.Children.Append("oper-summary", types.YChild{"OperSummary", &mplsLdpState.OperSummary})
    mplsLdpState.EntityData.Children.Append("forwarding-summary", types.YChild{"ForwardingSummary", &mplsLdpState.ForwardingSummary})
    mplsLdpState.EntityData.Children.Append("bindings-summary", types.YChild{"BindingsSummary", &mplsLdpState.BindingsSummary})
    mplsLdpState.EntityData.Children.Append("nsr-summary-all", types.YChild{"NsrSummaryAll", &mplsLdpState.NsrSummaryAll})
    mplsLdpState.EntityData.Children.Append("icpm-summary-all", types.YChild{"IcpmSummaryAll", &mplsLdpState.IcpmSummaryAll})
    mplsLdpState.EntityData.Children.Append("parameters", types.YChild{"Parameters", &mplsLdpState.Parameters})
    mplsLdpState.EntityData.Children.Append("capabilities", types.YChild{"Capabilities", &mplsLdpState.Capabilities})
    mplsLdpState.EntityData.Children.Append("backoff-parameters", types.YChild{"BackoffParameters", &mplsLdpState.BackoffParameters})
    mplsLdpState.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &mplsLdpState.GracefulRestart})
    mplsLdpState.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &mplsLdpState.Vrfs})
    mplsLdpState.EntityData.Children.Append("discovery", types.YChild{"Discovery", &mplsLdpState.Discovery})
    mplsLdpState.EntityData.Children.Append("forwarding", types.YChild{"Forwarding", &mplsLdpState.Forwarding})
    mplsLdpState.EntityData.Children.Append("bindings", types.YChild{"Bindings", &mplsLdpState.Bindings})
    mplsLdpState.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &mplsLdpState.Neighbors})
    mplsLdpState.EntityData.Children.Append("label-ranges", types.YChild{"LabelRanges", &mplsLdpState.LabelRanges})
    mplsLdpState.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpState.EntityData.YListKeys = []string {}

    return &(mplsLdpState.EntityData)
}

// MplsLdp_MplsLdpState_OperSummary
// LDP operational data summary
type MplsLdp_MplsLdpState_OperSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of configured VRFs (including default). The type is interface{} with
    // range: 0..4294967295.
    NumberOfVrf interface{}

    // Number of configured operational VRFs (including default). The type is
    // interface{} with range: 0..4294967295.
    NumberOfVrfOper interface{}

    // Number of known interfaces. The type is interface{} with range:
    // 0..4294967295.
    NumberOfInterfaces interface{}

    // Number of Forward Reference interfaces. The type is interface{} with range:
    // 0..4294967295.
    NumberOfFwdRefInterfaces interface{}

    // Number of auto-configured interfaces. The type is interface{} with range:
    // 0..4294967295.
    NumberOfAutocfgInterfaces interface{}

    // Total number of ipv4 RIB tables. The type is interface{} with range:
    // 0..4294967295.
    NoOfIpv4RibTbl interface{}

    // Number of ipv4 RIB tables registered. The type is interface{} with range:
    // 0..4294967295.
    NoOfIpv4RibTblReg interface{}

    // Common Summary information.
    Common MplsLdp_MplsLdpState_OperSummary_Common
}

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetEntityData() *types.CommonEntityData {
    operSummary.EntityData.YFilter = operSummary.YFilter
    operSummary.EntityData.YangName = "oper-summary"
    operSummary.EntityData.BundleName = "cisco_ios_xe"
    operSummary.EntityData.ParentYangName = "mpls-ldp-state"
    operSummary.EntityData.SegmentPath = "oper-summary"
    operSummary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    operSummary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    operSummary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    operSummary.EntityData.Children = types.NewOrderedMap()
    operSummary.EntityData.Children.Append("common", types.YChild{"Common", &operSummary.Common})
    operSummary.EntityData.Leafs = types.NewOrderedMap()
    operSummary.EntityData.Leafs.Append("number-of-vrf", types.YLeaf{"NumberOfVrf", operSummary.NumberOfVrf})
    operSummary.EntityData.Leafs.Append("number-of-vrf-oper", types.YLeaf{"NumberOfVrfOper", operSummary.NumberOfVrfOper})
    operSummary.EntityData.Leafs.Append("number-of-interfaces", types.YLeaf{"NumberOfInterfaces", operSummary.NumberOfInterfaces})
    operSummary.EntityData.Leafs.Append("number-of-fwd-ref-interfaces", types.YLeaf{"NumberOfFwdRefInterfaces", operSummary.NumberOfFwdRefInterfaces})
    operSummary.EntityData.Leafs.Append("number-of-autocfg-interfaces", types.YLeaf{"NumberOfAutocfgInterfaces", operSummary.NumberOfAutocfgInterfaces})
    operSummary.EntityData.Leafs.Append("no-of-ipv4-rib-tbl", types.YLeaf{"NoOfIpv4RibTbl", operSummary.NoOfIpv4RibTbl})
    operSummary.EntityData.Leafs.Append("no-of-ipv4-rib-tbl-reg", types.YLeaf{"NoOfIpv4RibTblReg", operSummary.NoOfIpv4RibTblReg})

    operSummary.EntityData.YListKeys = []string {}

    return &(operSummary.EntityData)
}

// MplsLdp_MplsLdpState_OperSummary_Common
// Common Summary information
type MplsLdp_MplsLdpState_OperSummary_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Families enabled. The type is Af.
    AddressFamilies interface{}

    // Number of neighbor. The type is interface{} with range: 0..4294967295.
    NumberOfNeighbors interface{}

    // Number of Graceful Restart neighbor. The type is interface{} with range:
    // 0..4294967295.
    NumberOfGracefulRestartNeighbors interface{}

    // Number of Downstream-On-Demand neighbor. The type is interface{} with
    // range: 0..4294967295.
    NumberOfDownstreamOnDemandNeighbors interface{}

    // Number of LDP discovery IPv4 hello adjacencies. The type is interface{}
    // with range: 0..4294967295.
    NumberofIpv4HelloAdj interface{}

    // Number of resolved IPv4 routes. The type is interface{} with range:
    // 0..4294967295.
    NumberOfIpv4Routes interface{}

    // Number of IPv4 local addresses. The type is interface{} with range:
    // 0..4294967295.
    NumberOfIpv4LocalAddresses interface{}

    // Number of LDP configured interfaces. The type is interface{} with range:
    // 0..4294967295.
    NumberOfLdpInterfaces interface{}

    // Number of LDP IPv4 configured interfaces. The type is interface{} with
    // range: 0..4294967295.
    NumberOfIpv4ldpInterfaces interface{}
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xe"
    common.EntityData.ParentYangName = "oper-summary"
    common.EntityData.SegmentPath = "common"
    common.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    common.EntityData.Children = types.NewOrderedMap()
    common.EntityData.Leafs = types.NewOrderedMap()
    common.EntityData.Leafs.Append("address-families", types.YLeaf{"AddressFamilies", common.AddressFamilies})
    common.EntityData.Leafs.Append("number-of-neighbors", types.YLeaf{"NumberOfNeighbors", common.NumberOfNeighbors})
    common.EntityData.Leafs.Append("number-of-graceful-restart-neighbors", types.YLeaf{"NumberOfGracefulRestartNeighbors", common.NumberOfGracefulRestartNeighbors})
    common.EntityData.Leafs.Append("number-of-downstream-on-demand-neighbors", types.YLeaf{"NumberOfDownstreamOnDemandNeighbors", common.NumberOfDownstreamOnDemandNeighbors})
    common.EntityData.Leafs.Append("numberof-ipv4-hello-adj", types.YLeaf{"NumberofIpv4HelloAdj", common.NumberofIpv4HelloAdj})
    common.EntityData.Leafs.Append("number-of-ipv4-routes", types.YLeaf{"NumberOfIpv4Routes", common.NumberOfIpv4Routes})
    common.EntityData.Leafs.Append("number-of-ipv4-local-addresses", types.YLeaf{"NumberOfIpv4LocalAddresses", common.NumberOfIpv4LocalAddresses})
    common.EntityData.Leafs.Append("number-of-ldp-interfaces", types.YLeaf{"NumberOfLdpInterfaces", common.NumberOfLdpInterfaces})
    common.EntityData.Leafs.Append("number-of-ipv4ldp-interfaces", types.YLeaf{"NumberOfIpv4ldpInterfaces", common.NumberOfIpv4ldpInterfaces})

    common.EntityData.YListKeys = []string {}

    return &(common.EntityData)
}

// MplsLdp_MplsLdpState_ForwardingSummary
// Summary information regarding LDP forwarding
// setup
type MplsLdp_MplsLdpState_ForwardingSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS forwarding enabled interface count. The type is interface{} with
    // range: 0..65535.
    IntfsFwdCount interface{}

    // Local label allocated count. The type is interface{} with range: 0..65535.
    LocalLbls interface{}

    // MPLS LDP forwarding prefix rewrite summary.
    Pfxs MplsLdp_MplsLdpState_ForwardingSummary_Pfxs

    // MPLS LDP forwarding rewrite next-hop/path summary.
    Nhs MplsLdp_MplsLdpState_ForwardingSummary_Nhs
}

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetEntityData() *types.CommonEntityData {
    forwardingSummary.EntityData.YFilter = forwardingSummary.YFilter
    forwardingSummary.EntityData.YangName = "forwarding-summary"
    forwardingSummary.EntityData.BundleName = "cisco_ios_xe"
    forwardingSummary.EntityData.ParentYangName = "mpls-ldp-state"
    forwardingSummary.EntityData.SegmentPath = "forwarding-summary"
    forwardingSummary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwardingSummary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwardingSummary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwardingSummary.EntityData.Children = types.NewOrderedMap()
    forwardingSummary.EntityData.Children.Append("pfxs", types.YChild{"Pfxs", &forwardingSummary.Pfxs})
    forwardingSummary.EntityData.Children.Append("nhs", types.YChild{"Nhs", &forwardingSummary.Nhs})
    forwardingSummary.EntityData.Leafs = types.NewOrderedMap()
    forwardingSummary.EntityData.Leafs.Append("intfs-fwd-count", types.YLeaf{"IntfsFwdCount", forwardingSummary.IntfsFwdCount})
    forwardingSummary.EntityData.Leafs.Append("local-lbls", types.YLeaf{"LocalLbls", forwardingSummary.LocalLbls})

    forwardingSummary.EntityData.YListKeys = []string {}

    return &(forwardingSummary.EntityData)
}

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs
// MPLS LDP forwarding prefix rewrite summary
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Prefix count. The type is interface{} with range: 0..65535.
    TotalPfxs interface{}

    // Count of prefixes with ECMP. The type is interface{} with range: 0..65535.
    EcmpPfxs interface{}

    // Count of FRR protected prefixes. The type is interface{} with range:
    // 0..65535.
    ProtectedPfxs interface{}

    // Labeled prefix count for all paths.
    LabeledPfxsAggr MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr

    // Labeled prefix count related to primary paths only.
    LabeledPfxsPrimary MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary

    // Labeled prefix count related to backup paths only.
    LabeledPfxsBackup MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup
}

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetEntityData() *types.CommonEntityData {
    pfxs.EntityData.YFilter = pfxs.YFilter
    pfxs.EntityData.YangName = "pfxs"
    pfxs.EntityData.BundleName = "cisco_ios_xe"
    pfxs.EntityData.ParentYangName = "forwarding-summary"
    pfxs.EntityData.SegmentPath = "pfxs"
    pfxs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pfxs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pfxs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pfxs.EntityData.Children = types.NewOrderedMap()
    pfxs.EntityData.Children.Append("labeled-pfxs-aggr", types.YChild{"LabeledPfxsAggr", &pfxs.LabeledPfxsAggr})
    pfxs.EntityData.Children.Append("labeled-pfxs-primary", types.YChild{"LabeledPfxsPrimary", &pfxs.LabeledPfxsPrimary})
    pfxs.EntityData.Children.Append("labeled-pfxs-backup", types.YChild{"LabeledPfxsBackup", &pfxs.LabeledPfxsBackup})
    pfxs.EntityData.Leafs = types.NewOrderedMap()
    pfxs.EntityData.Leafs.Append("total-pfxs", types.YLeaf{"TotalPfxs", pfxs.TotalPfxs})
    pfxs.EntityData.Leafs.Append("ecmp-pfxs", types.YLeaf{"EcmpPfxs", pfxs.EcmpPfxs})
    pfxs.EntityData.Leafs.Append("protected-pfxs", types.YLeaf{"ProtectedPfxs", pfxs.ProtectedPfxs})

    pfxs.EntityData.YListKeys = []string {}

    return &(pfxs.EntityData)
}

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr
// Labeled prefix count for all paths
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of labeled prefixes with 1 or more paths labeled. The type is
    // interface{} with range: 0..65535.
    LabeledPfxs interface{}

    // Count of labeled prefixes with some (but not ALL) paths labeled. The type
    // is interface{} with range: 0..65535.
    LabeledPfxsPartial interface{}

    // Count of labeled prefixes with ALL paths unlabeled. The type is interface{}
    // with range: 0..65535.
    UnlabeledPfxs interface{}
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetEntityData() *types.CommonEntityData {
    labeledPfxsAggr.EntityData.YFilter = labeledPfxsAggr.YFilter
    labeledPfxsAggr.EntityData.YangName = "labeled-pfxs-aggr"
    labeledPfxsAggr.EntityData.BundleName = "cisco_ios_xe"
    labeledPfxsAggr.EntityData.ParentYangName = "pfxs"
    labeledPfxsAggr.EntityData.SegmentPath = "labeled-pfxs-aggr"
    labeledPfxsAggr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labeledPfxsAggr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labeledPfxsAggr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labeledPfxsAggr.EntityData.Children = types.NewOrderedMap()
    labeledPfxsAggr.EntityData.Leafs = types.NewOrderedMap()
    labeledPfxsAggr.EntityData.Leafs.Append("labeled-pfxs", types.YLeaf{"LabeledPfxs", labeledPfxsAggr.LabeledPfxs})
    labeledPfxsAggr.EntityData.Leafs.Append("labeled-pfxs-partial", types.YLeaf{"LabeledPfxsPartial", labeledPfxsAggr.LabeledPfxsPartial})
    labeledPfxsAggr.EntityData.Leafs.Append("unlabeled-pfxs", types.YLeaf{"UnlabeledPfxs", labeledPfxsAggr.UnlabeledPfxs})

    labeledPfxsAggr.EntityData.YListKeys = []string {}

    return &(labeledPfxsAggr.EntityData)
}

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary
// Labeled prefix count related to primary paths
// only
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of labeled prefixes with 1 or more paths labeled. The type is
    // interface{} with range: 0..65535.
    LabeledPfxs interface{}

    // Count of labeled prefixes with some (but not ALL) paths labeled. The type
    // is interface{} with range: 0..65535.
    LabeledPfxsPartial interface{}

    // Count of labeled prefixes with ALL paths unlabeled. The type is interface{}
    // with range: 0..65535.
    UnlabeledPfxs interface{}
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetEntityData() *types.CommonEntityData {
    labeledPfxsPrimary.EntityData.YFilter = labeledPfxsPrimary.YFilter
    labeledPfxsPrimary.EntityData.YangName = "labeled-pfxs-primary"
    labeledPfxsPrimary.EntityData.BundleName = "cisco_ios_xe"
    labeledPfxsPrimary.EntityData.ParentYangName = "pfxs"
    labeledPfxsPrimary.EntityData.SegmentPath = "labeled-pfxs-primary"
    labeledPfxsPrimary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labeledPfxsPrimary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labeledPfxsPrimary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labeledPfxsPrimary.EntityData.Children = types.NewOrderedMap()
    labeledPfxsPrimary.EntityData.Leafs = types.NewOrderedMap()
    labeledPfxsPrimary.EntityData.Leafs.Append("labeled-pfxs", types.YLeaf{"LabeledPfxs", labeledPfxsPrimary.LabeledPfxs})
    labeledPfxsPrimary.EntityData.Leafs.Append("labeled-pfxs-partial", types.YLeaf{"LabeledPfxsPartial", labeledPfxsPrimary.LabeledPfxsPartial})
    labeledPfxsPrimary.EntityData.Leafs.Append("unlabeled-pfxs", types.YLeaf{"UnlabeledPfxs", labeledPfxsPrimary.UnlabeledPfxs})

    labeledPfxsPrimary.EntityData.YListKeys = []string {}

    return &(labeledPfxsPrimary.EntityData)
}

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup
// Labeled prefix count related to backup paths
// only
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of labeled prefixes with 1 or more paths labeled. The type is
    // interface{} with range: 0..65535.
    LabeledPfxs interface{}

    // Count of labeled prefixes with some (but not ALL) paths labeled. The type
    // is interface{} with range: 0..65535.
    LabeledPfxsPartial interface{}

    // Count of labeled prefixes with ALL paths unlabeled. The type is interface{}
    // with range: 0..65535.
    UnlabeledPfxs interface{}
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetEntityData() *types.CommonEntityData {
    labeledPfxsBackup.EntityData.YFilter = labeledPfxsBackup.YFilter
    labeledPfxsBackup.EntityData.YangName = "labeled-pfxs-backup"
    labeledPfxsBackup.EntityData.BundleName = "cisco_ios_xe"
    labeledPfxsBackup.EntityData.ParentYangName = "pfxs"
    labeledPfxsBackup.EntityData.SegmentPath = "labeled-pfxs-backup"
    labeledPfxsBackup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labeledPfxsBackup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labeledPfxsBackup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labeledPfxsBackup.EntityData.Children = types.NewOrderedMap()
    labeledPfxsBackup.EntityData.Leafs = types.NewOrderedMap()
    labeledPfxsBackup.EntityData.Leafs.Append("labeled-pfxs", types.YLeaf{"LabeledPfxs", labeledPfxsBackup.LabeledPfxs})
    labeledPfxsBackup.EntityData.Leafs.Append("labeled-pfxs-partial", types.YLeaf{"LabeledPfxsPartial", labeledPfxsBackup.LabeledPfxsPartial})
    labeledPfxsBackup.EntityData.Leafs.Append("unlabeled-pfxs", types.YLeaf{"UnlabeledPfxs", labeledPfxsBackup.UnlabeledPfxs})

    labeledPfxsBackup.EntityData.YListKeys = []string {}

    return &(labeledPfxsBackup.EntityData)
}

// MplsLdp_MplsLdpState_ForwardingSummary_Nhs
// MPLS LDP forwarding rewrite next-hop/path summary
type MplsLdp_MplsLdpState_ForwardingSummary_Nhs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total path count. The type is interface{} with range: 0..4294967295.
    TotalPaths interface{}

    // Count of FRR protected paths. The type is interface{} with range:
    // 0..4294967295.
    ProtectedPaths interface{}

    // Count of non-primary backup paths. The type is interface{} with range:
    // 0..4294967295.
    BackupPaths interface{}

    // Count of non-primary remote backup paths. The type is interface{} with
    // range: 0..4294967295.
    RemoteBackupPaths interface{}

    // Count of all labeled paths. The type is interface{} with range:
    // 0..4294967295.
    LabeledPaths interface{}

    // Count of labeled backup paths. The type is interface{} with range:
    // 0..4294967295.
    LabeledBackupPaths interface{}
}

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetEntityData() *types.CommonEntityData {
    nhs.EntityData.YFilter = nhs.YFilter
    nhs.EntityData.YangName = "nhs"
    nhs.EntityData.BundleName = "cisco_ios_xe"
    nhs.EntityData.ParentYangName = "forwarding-summary"
    nhs.EntityData.SegmentPath = "nhs"
    nhs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhs.EntityData.Children = types.NewOrderedMap()
    nhs.EntityData.Leafs = types.NewOrderedMap()
    nhs.EntityData.Leafs.Append("total-paths", types.YLeaf{"TotalPaths", nhs.TotalPaths})
    nhs.EntityData.Leafs.Append("protected-paths", types.YLeaf{"ProtectedPaths", nhs.ProtectedPaths})
    nhs.EntityData.Leafs.Append("backup-paths", types.YLeaf{"BackupPaths", nhs.BackupPaths})
    nhs.EntityData.Leafs.Append("remote-backup-paths", types.YLeaf{"RemoteBackupPaths", nhs.RemoteBackupPaths})
    nhs.EntityData.Leafs.Append("labeled-paths", types.YLeaf{"LabeledPaths", nhs.LabeledPaths})
    nhs.EntityData.Leafs.Append("labeled-backup-paths", types.YLeaf{"LabeledBackupPaths", nhs.LabeledBackupPaths})

    nhs.EntityData.YListKeys = []string {}

    return &(nhs.EntityData)
}

// MplsLdp_MplsLdpState_BindingsSummary
// Aggregate counters for the MPLS LDP LIB.
type MplsLdp_MplsLdpState_BindingsSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total bindings. The type is interface{} with range: 0..4294967295.
    BindingTotal interface{}

    // Bindings with no route. The type is interface{} with range: 0..4294967295.
    BindingNoRoute interface{}

    // Local bindings with no route. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalNoRoute interface{}

    // Number of local bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocal interface{}

    // Number of local null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalNull interface{}

    // Number of local implicit null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalImplicitNull interface{}

    // Number of local explicit null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalExplicitNull interface{}

    // Number of local non-null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalNonNull interface{}

    // This is the number of local bindings needing label but which hit the
    // Out-Of-Resource condition. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalOor interface{}

    // Lowest allocated label. The type is interface{} with range: 0..4294967295.
    LowestAllocatedLabel interface{}

    // Highest allocated label. The type is interface{} with range: 0..4294967295.
    HighestAllocatedLabel interface{}

    // Number of remote bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingRemote interface{}
}

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetEntityData() *types.CommonEntityData {
    bindingsSummary.EntityData.YFilter = bindingsSummary.YFilter
    bindingsSummary.EntityData.YangName = "bindings-summary"
    bindingsSummary.EntityData.BundleName = "cisco_ios_xe"
    bindingsSummary.EntityData.ParentYangName = "mpls-ldp-state"
    bindingsSummary.EntityData.SegmentPath = "bindings-summary"
    bindingsSummary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bindingsSummary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bindingsSummary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bindingsSummary.EntityData.Children = types.NewOrderedMap()
    bindingsSummary.EntityData.Leafs = types.NewOrderedMap()
    bindingsSummary.EntityData.Leafs.Append("binding-total", types.YLeaf{"BindingTotal", bindingsSummary.BindingTotal})
    bindingsSummary.EntityData.Leafs.Append("binding-no-route", types.YLeaf{"BindingNoRoute", bindingsSummary.BindingNoRoute})
    bindingsSummary.EntityData.Leafs.Append("binding-local-no-route", types.YLeaf{"BindingLocalNoRoute", bindingsSummary.BindingLocalNoRoute})
    bindingsSummary.EntityData.Leafs.Append("binding-local", types.YLeaf{"BindingLocal", bindingsSummary.BindingLocal})
    bindingsSummary.EntityData.Leafs.Append("binding-local-null", types.YLeaf{"BindingLocalNull", bindingsSummary.BindingLocalNull})
    bindingsSummary.EntityData.Leafs.Append("binding-local-implicit-null", types.YLeaf{"BindingLocalImplicitNull", bindingsSummary.BindingLocalImplicitNull})
    bindingsSummary.EntityData.Leafs.Append("binding-local-explicit-null", types.YLeaf{"BindingLocalExplicitNull", bindingsSummary.BindingLocalExplicitNull})
    bindingsSummary.EntityData.Leafs.Append("binding-local-non-null", types.YLeaf{"BindingLocalNonNull", bindingsSummary.BindingLocalNonNull})
    bindingsSummary.EntityData.Leafs.Append("binding-local-oor", types.YLeaf{"BindingLocalOor", bindingsSummary.BindingLocalOor})
    bindingsSummary.EntityData.Leafs.Append("lowest-allocated-label", types.YLeaf{"LowestAllocatedLabel", bindingsSummary.LowestAllocatedLabel})
    bindingsSummary.EntityData.Leafs.Append("highest-allocated-label", types.YLeaf{"HighestAllocatedLabel", bindingsSummary.HighestAllocatedLabel})
    bindingsSummary.EntityData.Leafs.Append("binding-remote", types.YLeaf{"BindingRemote", bindingsSummary.BindingRemote})

    bindingsSummary.EntityData.YListKeys = []string {}

    return &(bindingsSummary.EntityData)
}

// MplsLdp_MplsLdpState_NsrSummaryAll
// This is the LDP NSR summary for the device.
type MplsLdp_MplsLdpState_NsrSummaryAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // In label Request Records created. The type is interface{} with range:
    // 0..4294967295.
    NsrSumInLabelReqsCreated interface{}

    // In label Request Records freed. The type is interface{} with range:
    // 0..4294967295.
    NsrSumInLabelReqsFreed interface{}

    // In label Withdraw Records created. The type is interface{} with range:
    // 0..4294967295.
    NsrSumInLabelWithdrawCreated interface{}

    // In label Withdraw Records freed. The type is interface{} with range:
    // 0..4294967295.
    NsrSumInLabelWithdrawFreed interface{}

    // Local Address Withdraw set. The type is interface{} with range:
    // 0..4294967295.
    NsrSumLclAddrWithdrawSet interface{}

    // Local Address Withdraw cleared. The type is interface{} with range:
    // 0..4294967295.
    NsrSumLclAddrWithdrawCleared interface{}
}

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetEntityData() *types.CommonEntityData {
    nsrSummaryAll.EntityData.YFilter = nsrSummaryAll.YFilter
    nsrSummaryAll.EntityData.YangName = "nsr-summary-all"
    nsrSummaryAll.EntityData.BundleName = "cisco_ios_xe"
    nsrSummaryAll.EntityData.ParentYangName = "mpls-ldp-state"
    nsrSummaryAll.EntityData.SegmentPath = "nsr-summary-all"
    nsrSummaryAll.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nsrSummaryAll.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nsrSummaryAll.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nsrSummaryAll.EntityData.Children = types.NewOrderedMap()
    nsrSummaryAll.EntityData.Leafs = types.NewOrderedMap()
    nsrSummaryAll.EntityData.Leafs.Append("nsr-sum-in-label-reqs-created", types.YLeaf{"NsrSumInLabelReqsCreated", nsrSummaryAll.NsrSumInLabelReqsCreated})
    nsrSummaryAll.EntityData.Leafs.Append("nsr-sum-in-label-reqs-freed", types.YLeaf{"NsrSumInLabelReqsFreed", nsrSummaryAll.NsrSumInLabelReqsFreed})
    nsrSummaryAll.EntityData.Leafs.Append("nsr-sum-in-label-withdraw-created", types.YLeaf{"NsrSumInLabelWithdrawCreated", nsrSummaryAll.NsrSumInLabelWithdrawCreated})
    nsrSummaryAll.EntityData.Leafs.Append("nsr-sum-in-label-withdraw-freed", types.YLeaf{"NsrSumInLabelWithdrawFreed", nsrSummaryAll.NsrSumInLabelWithdrawFreed})
    nsrSummaryAll.EntityData.Leafs.Append("nsr-sum-lcl-addr-withdraw-set", types.YLeaf{"NsrSumLclAddrWithdrawSet", nsrSummaryAll.NsrSumLclAddrWithdrawSet})
    nsrSummaryAll.EntityData.Leafs.Append("nsr-sum-lcl-addr-withdraw-cleared", types.YLeaf{"NsrSumLclAddrWithdrawCleared", nsrSummaryAll.NsrSumLclAddrWithdrawCleared})

    nsrSummaryAll.EntityData.YListKeys = []string {}

    return &(nsrSummaryAll.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll
// Summary info for LDP ICPM/ICCP.
type MplsLdp_MplsLdpState_IcpmSummaryAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICCP RG Connect count. The type is interface{} with range: 0..4294967295.
    IccpRgConnCount interface{}

    // ICCP RG Disconnect count. The type is interface{} with range:
    // 0..4294967295.
    IccpRgDisconnCount interface{}

    // ICCP RG Notif count. The type is interface{} with range: 0..4294967295.
    IccpRgNotifCount interface{}

    // ICCP RG App Data count. The type is interface{} with range: 0..4294967295.
    IccpRgAppDataCount interface{}

    // This defines the ICPM RGID Table.
    IcpmRgidTableInfo MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo

    // This is a list of ICPM sessions.
    IcpmSessionTable MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable
}

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetEntityData() *types.CommonEntityData {
    icpmSummaryAll.EntityData.YFilter = icpmSummaryAll.YFilter
    icpmSummaryAll.EntityData.YangName = "icpm-summary-all"
    icpmSummaryAll.EntityData.BundleName = "cisco_ios_xe"
    icpmSummaryAll.EntityData.ParentYangName = "mpls-ldp-state"
    icpmSummaryAll.EntityData.SegmentPath = "icpm-summary-all"
    icpmSummaryAll.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icpmSummaryAll.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icpmSummaryAll.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icpmSummaryAll.EntityData.Children = types.NewOrderedMap()
    icpmSummaryAll.EntityData.Children.Append("icpm-rgid-table-info", types.YChild{"IcpmRgidTableInfo", &icpmSummaryAll.IcpmRgidTableInfo})
    icpmSummaryAll.EntityData.Children.Append("icpm-session-table", types.YChild{"IcpmSessionTable", &icpmSummaryAll.IcpmSessionTable})
    icpmSummaryAll.EntityData.Leafs = types.NewOrderedMap()
    icpmSummaryAll.EntityData.Leafs.Append("iccp-rg-conn-count", types.YLeaf{"IccpRgConnCount", icpmSummaryAll.IccpRgConnCount})
    icpmSummaryAll.EntityData.Leafs.Append("iccp-rg-disconn-count", types.YLeaf{"IccpRgDisconnCount", icpmSummaryAll.IccpRgDisconnCount})
    icpmSummaryAll.EntityData.Leafs.Append("iccp-rg-notif-count", types.YLeaf{"IccpRgNotifCount", icpmSummaryAll.IccpRgNotifCount})
    icpmSummaryAll.EntityData.Leafs.Append("iccp-rg-app-data-count", types.YLeaf{"IccpRgAppDataCount", icpmSummaryAll.IccpRgAppDataCount})

    icpmSummaryAll.EntityData.YListKeys = []string {}

    return &(icpmSummaryAll.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo
// This defines the ICPM RGID Table
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is the data for an individual ICPM Rredundandy Group,. The type is
    // slice of MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup.
    RedGroup []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetEntityData() *types.CommonEntityData {
    icpmRgidTableInfo.EntityData.YFilter = icpmRgidTableInfo.YFilter
    icpmRgidTableInfo.EntityData.YangName = "icpm-rgid-table-info"
    icpmRgidTableInfo.EntityData.BundleName = "cisco_ios_xe"
    icpmRgidTableInfo.EntityData.ParentYangName = "icpm-summary-all"
    icpmRgidTableInfo.EntityData.SegmentPath = "icpm-rgid-table-info"
    icpmRgidTableInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icpmRgidTableInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icpmRgidTableInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icpmRgidTableInfo.EntityData.Children = types.NewOrderedMap()
    icpmRgidTableInfo.EntityData.Children.Append("red-group", types.YChild{"RedGroup", nil})
    for i := range icpmRgidTableInfo.RedGroup {
        icpmRgidTableInfo.EntityData.Children.Append(types.GetSegmentPath(icpmRgidTableInfo.RedGroup[i]), types.YChild{"RedGroup", icpmRgidTableInfo.RedGroup[i]})
    }
    icpmRgidTableInfo.EntityData.Leafs = types.NewOrderedMap()

    icpmRgidTableInfo.EntityData.YListKeys = []string {}

    return &(icpmRgidTableInfo.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup
// This is the data for an individual ICPM Rredundandy
// Group,
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is the ICPM RG identifier. The type is
    // interface{} with range: 0..4294967295.
    RgId interface{}

    // This list contains all active icpm protocols. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols.
    IcpmProtocols []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetEntityData() *types.CommonEntityData {
    redGroup.EntityData.YFilter = redGroup.YFilter
    redGroup.EntityData.YangName = "red-group"
    redGroup.EntityData.BundleName = "cisco_ios_xe"
    redGroup.EntityData.ParentYangName = "icpm-rgid-table-info"
    redGroup.EntityData.SegmentPath = "red-group" + types.AddKeyToken(redGroup.RgId, "rg-id")
    redGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    redGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    redGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    redGroup.EntityData.Children = types.NewOrderedMap()
    redGroup.EntityData.Children.Append("icpm-protocols", types.YChild{"IcpmProtocols", nil})
    for i := range redGroup.IcpmProtocols {
        redGroup.EntityData.Children.Append(types.GetSegmentPath(redGroup.IcpmProtocols[i]), types.YChild{"IcpmProtocols", redGroup.IcpmProtocols[i]})
    }
    redGroup.EntityData.Leafs = types.NewOrderedMap()
    redGroup.EntityData.Leafs.Append("rg-id", types.YLeaf{"RgId", redGroup.RgId})

    redGroup.EntityData.YListKeys = []string {"RgId"}

    return &(redGroup.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols
// This list contains all active icpm protocols.
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ICPM Type. The type is one of the following:
    // IcpmTypeIccp.
    IcpmType interface{}

    // List of Redundancy Groups. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups.
    RedunGroups []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetEntityData() *types.CommonEntityData {
    icpmProtocols.EntityData.YFilter = icpmProtocols.YFilter
    icpmProtocols.EntityData.YangName = "icpm-protocols"
    icpmProtocols.EntityData.BundleName = "cisco_ios_xe"
    icpmProtocols.EntityData.ParentYangName = "red-group"
    icpmProtocols.EntityData.SegmentPath = "icpm-protocols" + types.AddKeyToken(icpmProtocols.IcpmType, "icpm-type")
    icpmProtocols.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icpmProtocols.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icpmProtocols.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icpmProtocols.EntityData.Children = types.NewOrderedMap()
    icpmProtocols.EntityData.Children.Append("redun-groups", types.YChild{"RedunGroups", nil})
    for i := range icpmProtocols.RedunGroups {
        icpmProtocols.EntityData.Children.Append(types.GetSegmentPath(icpmProtocols.RedunGroups[i]), types.YChild{"RedunGroups", icpmProtocols.RedunGroups[i]})
    }
    icpmProtocols.EntityData.Leafs = types.NewOrderedMap()
    icpmProtocols.EntityData.Leafs.Append("icpm-type", types.YLeaf{"IcpmType", icpmProtocols.IcpmType})

    icpmProtocols.EntityData.YListKeys = []string {"IcpmType"}

    return &(icpmProtocols.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups
// List of Redundancy Groups
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Redundancy Group Identifier. The type is
    // interface{} with range: 0..4294967295.
    RgId interface{}

    // LSR identifier. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerId interface{}

    // Client Identifier. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // ICCP State. The type is string.
    State interface{}

    // List of apps. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps.
    IccpApps []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetEntityData() *types.CommonEntityData {
    redunGroups.EntityData.YFilter = redunGroups.YFilter
    redunGroups.EntityData.YangName = "redun-groups"
    redunGroups.EntityData.BundleName = "cisco_ios_xe"
    redunGroups.EntityData.ParentYangName = "icpm-protocols"
    redunGroups.EntityData.SegmentPath = "redun-groups" + types.AddKeyToken(redunGroups.RgId, "rg-id")
    redunGroups.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    redunGroups.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    redunGroups.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    redunGroups.EntityData.Children = types.NewOrderedMap()
    redunGroups.EntityData.Children.Append("iccp-apps", types.YChild{"IccpApps", nil})
    for i := range redunGroups.IccpApps {
        redunGroups.EntityData.Children.Append(types.GetSegmentPath(redunGroups.IccpApps[i]), types.YChild{"IccpApps", redunGroups.IccpApps[i]})
    }
    redunGroups.EntityData.Leafs = types.NewOrderedMap()
    redunGroups.EntityData.Leafs.Append("rg-id", types.YLeaf{"RgId", redunGroups.RgId})
    redunGroups.EntityData.Leafs.Append("peer-id", types.YLeaf{"PeerId", redunGroups.PeerId})
    redunGroups.EntityData.Leafs.Append("client_id", types.YLeaf{"ClientId", redunGroups.ClientId})
    redunGroups.EntityData.Leafs.Append("state", types.YLeaf{"State", redunGroups.State})

    redunGroups.EntityData.YListKeys = []string {"RgId"}

    return &(redunGroups.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps
// List of apps
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ICCP App Type. The type is one of the following:
    // IccpTypeMlacp.
    IccpApp interface{}

    // App State. The type is IccpState.
    AppState interface{}

    // ICCP App Protocol Version. The type is interface{} with range:
    // 0..4294967295.
    PtclVer interface{}
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetEntityData() *types.CommonEntityData {
    iccpApps.EntityData.YFilter = iccpApps.YFilter
    iccpApps.EntityData.YangName = "iccp-apps"
    iccpApps.EntityData.BundleName = "cisco_ios_xe"
    iccpApps.EntityData.ParentYangName = "redun-groups"
    iccpApps.EntityData.SegmentPath = "iccp-apps" + types.AddKeyToken(iccpApps.IccpApp, "iccp-app")
    iccpApps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iccpApps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iccpApps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iccpApps.EntityData.Children = types.NewOrderedMap()
    iccpApps.EntityData.Leafs = types.NewOrderedMap()
    iccpApps.EntityData.Leafs.Append("iccp-app", types.YLeaf{"IccpApp", iccpApps.IccpApp})
    iccpApps.EntityData.Leafs.Append("app-state", types.YLeaf{"AppState", iccpApps.AppState})
    iccpApps.EntityData.Leafs.Append("ptcl-ver", types.YLeaf{"PtclVer", iccpApps.PtclVer})

    iccpApps.EntityData.YListKeys = []string {"IccpApp"}

    return &(iccpApps.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable
// This is a list of ICPM sessions.
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ICPM LDP Session Table. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable.
    SessionTable []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetEntityData() *types.CommonEntityData {
    icpmSessionTable.EntityData.YFilter = icpmSessionTable.YFilter
    icpmSessionTable.EntityData.YangName = "icpm-session-table"
    icpmSessionTable.EntityData.BundleName = "cisco_ios_xe"
    icpmSessionTable.EntityData.ParentYangName = "icpm-summary-all"
    icpmSessionTable.EntityData.SegmentPath = "icpm-session-table"
    icpmSessionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icpmSessionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icpmSessionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icpmSessionTable.EntityData.Children = types.NewOrderedMap()
    icpmSessionTable.EntityData.Children.Append("session-table", types.YChild{"SessionTable", nil})
    for i := range icpmSessionTable.SessionTable {
        icpmSessionTable.EntityData.Children.Append(types.GetSegmentPath(icpmSessionTable.SessionTable[i]), types.YChild{"SessionTable", icpmSessionTable.SessionTable[i]})
    }
    icpmSessionTable.EntityData.Leafs = types.NewOrderedMap()

    icpmSessionTable.EntityData.YListKeys = []string {}

    return &(icpmSessionTable.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable
// ICPM LDP Session Table
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This is the ICPM sesion identifier. The type is
    // interface{} with range: 0..4294967295.
    SessionId interface{}

    // This list contains all active icpm protocols. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols.
    IcpmProtocols []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetEntityData() *types.CommonEntityData {
    sessionTable.EntityData.YFilter = sessionTable.YFilter
    sessionTable.EntityData.YangName = "session-table"
    sessionTable.EntityData.BundleName = "cisco_ios_xe"
    sessionTable.EntityData.ParentYangName = "icpm-session-table"
    sessionTable.EntityData.SegmentPath = "session-table" + types.AddKeyToken(sessionTable.SessionId, "session-id")
    sessionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sessionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sessionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sessionTable.EntityData.Children = types.NewOrderedMap()
    sessionTable.EntityData.Children.Append("icpm-protocols", types.YChild{"IcpmProtocols", nil})
    for i := range sessionTable.IcpmProtocols {
        sessionTable.EntityData.Children.Append(types.GetSegmentPath(sessionTable.IcpmProtocols[i]), types.YChild{"IcpmProtocols", sessionTable.IcpmProtocols[i]})
    }
    sessionTable.EntityData.Leafs = types.NewOrderedMap()
    sessionTable.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", sessionTable.SessionId})

    sessionTable.EntityData.YListKeys = []string {"SessionId"}

    return &(sessionTable.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols
// This list contains all active icpm protocols.
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ICPM Type. The type is one of the following:
    // IcpmTypeIccp.
    IcpmType interface{}

    // List of Redundancy Groups. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups.
    RedunGroups []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetEntityData() *types.CommonEntityData {
    icpmProtocols.EntityData.YFilter = icpmProtocols.YFilter
    icpmProtocols.EntityData.YangName = "icpm-protocols"
    icpmProtocols.EntityData.BundleName = "cisco_ios_xe"
    icpmProtocols.EntityData.ParentYangName = "session-table"
    icpmProtocols.EntityData.SegmentPath = "icpm-protocols" + types.AddKeyToken(icpmProtocols.IcpmType, "icpm-type")
    icpmProtocols.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    icpmProtocols.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    icpmProtocols.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    icpmProtocols.EntityData.Children = types.NewOrderedMap()
    icpmProtocols.EntityData.Children.Append("redun-groups", types.YChild{"RedunGroups", nil})
    for i := range icpmProtocols.RedunGroups {
        icpmProtocols.EntityData.Children.Append(types.GetSegmentPath(icpmProtocols.RedunGroups[i]), types.YChild{"RedunGroups", icpmProtocols.RedunGroups[i]})
    }
    icpmProtocols.EntityData.Leafs = types.NewOrderedMap()
    icpmProtocols.EntityData.Leafs.Append("icpm-type", types.YLeaf{"IcpmType", icpmProtocols.IcpmType})

    icpmProtocols.EntityData.YListKeys = []string {"IcpmType"}

    return &(icpmProtocols.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups
// List of Redundancy Groups
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Redundancy Group Identifier. The type is
    // interface{} with range: 0..4294967295.
    RgId interface{}

    // LSR identifier. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerId interface{}

    // Client Identifier. The type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // ICCP State. The type is string.
    State interface{}

    // List of apps. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps.
    IccpApps []*MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetEntityData() *types.CommonEntityData {
    redunGroups.EntityData.YFilter = redunGroups.YFilter
    redunGroups.EntityData.YangName = "redun-groups"
    redunGroups.EntityData.BundleName = "cisco_ios_xe"
    redunGroups.EntityData.ParentYangName = "icpm-protocols"
    redunGroups.EntityData.SegmentPath = "redun-groups" + types.AddKeyToken(redunGroups.RgId, "rg-id")
    redunGroups.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    redunGroups.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    redunGroups.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    redunGroups.EntityData.Children = types.NewOrderedMap()
    redunGroups.EntityData.Children.Append("iccp-apps", types.YChild{"IccpApps", nil})
    for i := range redunGroups.IccpApps {
        redunGroups.EntityData.Children.Append(types.GetSegmentPath(redunGroups.IccpApps[i]), types.YChild{"IccpApps", redunGroups.IccpApps[i]})
    }
    redunGroups.EntityData.Leafs = types.NewOrderedMap()
    redunGroups.EntityData.Leafs.Append("rg-id", types.YLeaf{"RgId", redunGroups.RgId})
    redunGroups.EntityData.Leafs.Append("peer-id", types.YLeaf{"PeerId", redunGroups.PeerId})
    redunGroups.EntityData.Leafs.Append("client_id", types.YLeaf{"ClientId", redunGroups.ClientId})
    redunGroups.EntityData.Leafs.Append("state", types.YLeaf{"State", redunGroups.State})

    redunGroups.EntityData.YListKeys = []string {"RgId"}

    return &(redunGroups.EntityData)
}

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps
// List of apps
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. ICCP App Type. The type is one of the following:
    // IccpTypeMlacp.
    IccpApp interface{}

    // App State. The type is IccpState.
    AppState interface{}

    // ICCP App Protocol Version. The type is interface{} with range:
    // 0..4294967295.
    PtclVer interface{}
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetEntityData() *types.CommonEntityData {
    iccpApps.EntityData.YFilter = iccpApps.YFilter
    iccpApps.EntityData.YangName = "iccp-apps"
    iccpApps.EntityData.BundleName = "cisco_ios_xe"
    iccpApps.EntityData.ParentYangName = "redun-groups"
    iccpApps.EntityData.SegmentPath = "iccp-apps" + types.AddKeyToken(iccpApps.IccpApp, "iccp-app")
    iccpApps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    iccpApps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    iccpApps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    iccpApps.EntityData.Children = types.NewOrderedMap()
    iccpApps.EntityData.Leafs = types.NewOrderedMap()
    iccpApps.EntityData.Leafs.Append("iccp-app", types.YLeaf{"IccpApp", iccpApps.IccpApp})
    iccpApps.EntityData.Leafs.Append("app-state", types.YLeaf{"AppState", iccpApps.AppState})
    iccpApps.EntityData.Leafs.Append("ptcl-ver", types.YLeaf{"PtclVer", iccpApps.PtclVer})

    iccpApps.EntityData.YListKeys = []string {"IccpApp"}

    return &(iccpApps.EntityData)
}

// MplsLdp_MplsLdpState_Parameters
// MPLS LDP Global Parameters
type MplsLdp_MplsLdpState_Parameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global MD5 password enabled. The type is bool.
    GlobalMd5PasswordEnabled interface{}

    // Protocol version. The type is interface{} with range: 0..4294967295.
    ProtocolVersion interface{}

    // Keepalive interval in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    KeepaliveInterval interface{}

    // Session hold time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    SessionHoldTime interface{}

    // LIB entry no route timeout in second. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    LeNoRouteTimeout interface{}

    // Delay (sec) in Binding Withdrawal for an Address Family. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    AfBindingWithdrawDelay interface{}

    // Maximum number of LDP enabled attached interfaces. The type is interface{}
    // with range: 0..4294967295.
    MaxIntfAttached interface{}

    // Maximum number of LDP enabled TE interfaces. The type is interface{} with
    // range: 0..4294967295.
    MaxIntfTe interface{}

    // Maximum number of LDP peers. The type is interface{} with range:
    // 0..4294967295.
    MaxPeer interface{}

    // This is a counter of the number of times LDP attempted to create a label or
    // binding and failed due a memory allocation failure. The type is interface{}
    // with range: 0..4294967295.
    OutOfMemState interface{}

    // Discovery quick-start disabled on some enabled interfaces. The type is
    // bool.
    DiscoveryQuickStartDisabledOnInterfaces interface{}

    // Maximum number of hops for Downstream-on-Demand. The type is interface{}
    // with range: 0..4294967295.
    DodMaxHop interface{}

    // This entry describes an LDP feature available on the device. This does not
    // indicate whether the feature is enabled, just the raw ability to support
    // the feature. The features may include, but are not limited to:
    // 'Auto-Configuration', 'Basic', 'ICPM', 'IP-over-MPLS', 'IGP-Sync', 'LLAF',
    // 'TCP-MD5-Rollover', 'TDP', and 'NSR'. The type is slice of string.
    Feature []interface{}

    // A indication of whether this LSR has enabled loop detection. Since Loop
    // Detection is determined during Session Initialization, an individual
    // session may not be running with loop detection.  This object simply gives
    // an indication of whether or not the LSR has the ability enabled to support
    // Loop Detection and which types. The type is LoopDetectionType.
    LoopDetection interface{}

    // Per AF parameters. The type is slice of
    // MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter.
    AddressFamilyParameter []*MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xe"
    parameters.EntityData.ParentYangName = "mpls-ldp-state"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Children.Append("address-family-parameter", types.YChild{"AddressFamilyParameter", nil})
    for i := range parameters.AddressFamilyParameter {
        parameters.EntityData.Children.Append(types.GetSegmentPath(parameters.AddressFamilyParameter[i]), types.YChild{"AddressFamilyParameter", parameters.AddressFamilyParameter[i]})
    }
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("global-md5-password-enabled", types.YLeaf{"GlobalMd5PasswordEnabled", parameters.GlobalMd5PasswordEnabled})
    parameters.EntityData.Leafs.Append("protocol-version", types.YLeaf{"ProtocolVersion", parameters.ProtocolVersion})
    parameters.EntityData.Leafs.Append("keepalive-interval", types.YLeaf{"KeepaliveInterval", parameters.KeepaliveInterval})
    parameters.EntityData.Leafs.Append("session-hold-time", types.YLeaf{"SessionHoldTime", parameters.SessionHoldTime})
    parameters.EntityData.Leafs.Append("le-no-route-timeout", types.YLeaf{"LeNoRouteTimeout", parameters.LeNoRouteTimeout})
    parameters.EntityData.Leafs.Append("af-binding-withdraw-delay", types.YLeaf{"AfBindingWithdrawDelay", parameters.AfBindingWithdrawDelay})
    parameters.EntityData.Leafs.Append("max-intf-attached", types.YLeaf{"MaxIntfAttached", parameters.MaxIntfAttached})
    parameters.EntityData.Leafs.Append("max-intf-te", types.YLeaf{"MaxIntfTe", parameters.MaxIntfTe})
    parameters.EntityData.Leafs.Append("max-peer", types.YLeaf{"MaxPeer", parameters.MaxPeer})
    parameters.EntityData.Leafs.Append("out-of-mem-state", types.YLeaf{"OutOfMemState", parameters.OutOfMemState})
    parameters.EntityData.Leafs.Append("discovery-quick-start-disabled-on-interfaces", types.YLeaf{"DiscoveryQuickStartDisabledOnInterfaces", parameters.DiscoveryQuickStartDisabledOnInterfaces})
    parameters.EntityData.Leafs.Append("dod-max-hop", types.YLeaf{"DodMaxHop", parameters.DodMaxHop})
    parameters.EntityData.Leafs.Append("feature", types.YLeaf{"Feature", parameters.Feature})
    parameters.EntityData.Leafs.Append("loop-detection", types.YLeaf{"LoopDetection", parameters.LoopDetection})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

// MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter
// Per AF parameters
type MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is Af.
    AddressFamily interface{}

    // This is the Discovery transport address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DiscoveryTransportAddress interface{}

    // Accepting targeted Hellos. The type is bool.
    IsAcceptingTargetedHellos interface{}

    // This contains the filter name for targeted hellos. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string.
    TargetedHelloFilter interface{}
}

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetEntityData() *types.CommonEntityData {
    addressFamilyParameter.EntityData.YFilter = addressFamilyParameter.YFilter
    addressFamilyParameter.EntityData.YangName = "address-family-parameter"
    addressFamilyParameter.EntityData.BundleName = "cisco_ios_xe"
    addressFamilyParameter.EntityData.ParentYangName = "parameters"
    addressFamilyParameter.EntityData.SegmentPath = "address-family-parameter" + types.AddKeyToken(addressFamilyParameter.AddressFamily, "address-family")
    addressFamilyParameter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    addressFamilyParameter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    addressFamilyParameter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    addressFamilyParameter.EntityData.Children = types.NewOrderedMap()
    addressFamilyParameter.EntityData.Leafs = types.NewOrderedMap()
    addressFamilyParameter.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", addressFamilyParameter.AddressFamily})
    addressFamilyParameter.EntityData.Leafs.Append("discovery-transport-address", types.YLeaf{"DiscoveryTransportAddress", addressFamilyParameter.DiscoveryTransportAddress})
    addressFamilyParameter.EntityData.Leafs.Append("is-accepting-targeted-hellos", types.YLeaf{"IsAcceptingTargetedHellos", addressFamilyParameter.IsAcceptingTargetedHellos})
    addressFamilyParameter.EntityData.Leafs.Append("targeted-hello-filter", types.YLeaf{"TargetedHelloFilter", addressFamilyParameter.TargetedHelloFilter})

    addressFamilyParameter.EntityData.YListKeys = []string {"AddressFamily"}

    return &(addressFamilyParameter.EntityData)
}

// MplsLdp_MplsLdpState_Capabilities
// LDP capability database information
type MplsLdp_MplsLdpState_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on LDP capability. The type is slice of
    // MplsLdp_MplsLdpState_Capabilities_Capability.
    Capability []*MplsLdp_MplsLdpState_Capabilities_Capability
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "cisco_ios_xe"
    capabilities.EntityData.ParentYangName = "mpls-ldp-state"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capabilities.EntityData.Children = types.NewOrderedMap()
    capabilities.EntityData.Children.Append("capability", types.YChild{"Capability", nil})
    for i := range capabilities.Capability {
        capabilities.EntityData.Children.Append(types.GetSegmentPath(capabilities.Capability[i]), types.YChild{"Capability", capabilities.Capability[i]})
    }
    capabilities.EntityData.Leafs = types.NewOrderedMap()

    capabilities.EntityData.YListKeys = []string {}

    return &(capabilities.EntityData)
}

// MplsLdp_MplsLdpState_Capabilities_Capability
// Information on LDP capability
type MplsLdp_MplsLdpState_Capabilities_Capability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Capability type (IANA assigned). The type is
    // interface{} with range: 0..65535.
    CapType interface{}

    // Capability owner. The type is string.
    CapabilityOwner interface{}

    // Capability description. The type is string with length: 0..80.
    CapDes interface{}

    // Capability data length. The type is interface{} with range: 0..65535.
    CapabilityDataLength interface{}

    // Capability data. The type is string.
    CapabilityData interface{}
}

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetEntityData() *types.CommonEntityData {
    capability.EntityData.YFilter = capability.YFilter
    capability.EntityData.YangName = "capability"
    capability.EntityData.BundleName = "cisco_ios_xe"
    capability.EntityData.ParentYangName = "capabilities"
    capability.EntityData.SegmentPath = "capability" + types.AddKeyToken(capability.CapType, "cap-type")
    capability.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capability.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capability.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capability.EntityData.Children = types.NewOrderedMap()
    capability.EntityData.Leafs = types.NewOrderedMap()
    capability.EntityData.Leafs.Append("cap-type", types.YLeaf{"CapType", capability.CapType})
    capability.EntityData.Leafs.Append("capability-owner", types.YLeaf{"CapabilityOwner", capability.CapabilityOwner})
    capability.EntityData.Leafs.Append("cap-des", types.YLeaf{"CapDes", capability.CapDes})
    capability.EntityData.Leafs.Append("capability-data-length", types.YLeaf{"CapabilityDataLength", capability.CapabilityDataLength})
    capability.EntityData.Leafs.Append("capability-data", types.YLeaf{"CapabilityData", capability.CapabilityData})

    capability.EntityData.YListKeys = []string {"CapType"}

    return &(capability.EntityData)
}

// MplsLdp_MplsLdpState_BackoffParameters
// MPLS LDP Session Backoff Information
type MplsLdp_MplsLdpState_BackoffParameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial backoff value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    InitialSeconds interface{}

    // Maximum backoff value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    MaximumSeconds interface{}

    // Current backoff seconds count. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    BackoffSeconds interface{}

    // Current backoff waiting seconds count. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    WaitingSeconds interface{}
}

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetEntityData() *types.CommonEntityData {
    backoffParameters.EntityData.YFilter = backoffParameters.YFilter
    backoffParameters.EntityData.YangName = "backoff-parameters"
    backoffParameters.EntityData.BundleName = "cisco_ios_xe"
    backoffParameters.EntityData.ParentYangName = "mpls-ldp-state"
    backoffParameters.EntityData.SegmentPath = "backoff-parameters"
    backoffParameters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    backoffParameters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    backoffParameters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    backoffParameters.EntityData.Children = types.NewOrderedMap()
    backoffParameters.EntityData.Leafs = types.NewOrderedMap()
    backoffParameters.EntityData.Leafs.Append("initial-seconds", types.YLeaf{"InitialSeconds", backoffParameters.InitialSeconds})
    backoffParameters.EntityData.Leafs.Append("maximum-seconds", types.YLeaf{"MaximumSeconds", backoffParameters.MaximumSeconds})
    backoffParameters.EntityData.Leafs.Append("backoff-seconds", types.YLeaf{"BackoffSeconds", backoffParameters.BackoffSeconds})
    backoffParameters.EntityData.Leafs.Append("waiting-seconds", types.YLeaf{"WaitingSeconds", backoffParameters.WaitingSeconds})

    backoffParameters.EntityData.YListKeys = []string {}

    return &(backoffParameters.EntityData)
}

// MplsLdp_MplsLdpState_GracefulRestart
// MPLS LDP Graceful Restart Information
type MplsLdp_MplsLdpState_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is graceful restart configured. The type is bool.
    IsGracefulRestartConfigured interface{}

    // Reconnect timeout value in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    GracefulRestartReconnectTimeout interface{}

    // Graceful restart forward state hold time in seconds. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    GracefulRestartForwardingStateHoldTime interface{}

    // Is graceful restart forwarding state hold timer running. The type is
    // interface{}.
    IsForwardingStateHoldTimerRunning interface{}

    // Forwarding state hold timer remaining time in seconds. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    ForwardingStateHoldTimerRemainingSeconds interface{}
}

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xe"
    gracefulRestart.EntityData.ParentYangName = "mpls-ldp-state"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs.Append("is-graceful-restart-configured", types.YLeaf{"IsGracefulRestartConfigured", gracefulRestart.IsGracefulRestartConfigured})
    gracefulRestart.EntityData.Leafs.Append("graceful-restart-reconnect-timeout", types.YLeaf{"GracefulRestartReconnectTimeout", gracefulRestart.GracefulRestartReconnectTimeout})
    gracefulRestart.EntityData.Leafs.Append("graceful-restart-forwarding-state-hold-time", types.YLeaf{"GracefulRestartForwardingStateHoldTime", gracefulRestart.GracefulRestartForwardingStateHoldTime})
    gracefulRestart.EntityData.Leafs.Append("is-forwarding-state-hold-timer-running", types.YLeaf{"IsForwardingStateHoldTimerRunning", gracefulRestart.IsForwardingStateHoldTimerRunning})
    gracefulRestart.EntityData.Leafs.Append("forwarding-state-hold-timer-remaining-seconds", types.YLeaf{"ForwardingStateHoldTimerRemainingSeconds", gracefulRestart.ForwardingStateHoldTimerRemainingSeconds})

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs
// MPLS LDP per-VRF operational data.
type MplsLdp_MplsLdpState_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP Operational data for a given VRF. The type is slice of
    // MplsLdp_MplsLdpState_Vrfs_Vrf.
    Vrf []*MplsLdp_MplsLdpState_Vrfs_Vrf
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xe"
    vrfs.EntityData.ParentYangName = "mpls-ldp-state"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf
// MPLS LDP Operational data for a given VRF.
type MplsLdp_MplsLdpState_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // MPLS LDP per VRF summarized Information.
    VrfSummary MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary

    // Address Family specific operational data.
    Afs MplsLdp_MplsLdpState_Vrfs_Vrf_Afs
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xe"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("vrf-summary", types.YChild{"VrfSummary", &vrf.VrfSummary})
    vrf.EntityData.Children.Append("afs", types.YChild{"Afs", &vrf.Afs})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary
// MPLS LDP per VRF summarized Information
type MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Families enabled. The type is Af.
    AddressFamilies interface{}

    // Number of neighbor. The type is interface{} with range: 0..4294967295.
    NumberOfNeighbors interface{}

    // Number of Graceful Restart neighbor. The type is interface{} with range:
    // 0..4294967295.
    NumberOfGracefulRestartNeighbors interface{}

    // Number of Downstream-On-Demand neighbor. The type is interface{} with
    // range: 0..4294967295.
    NumberOfDownstreamOnDemandNeighbors interface{}

    // Number of LDP discovery IPv4 hello adjacencies. The type is interface{}
    // with range: 0..4294967295.
    NumberofIpv4HelloAdj interface{}

    // Number of resolved IPv4 routes. The type is interface{} with range:
    // 0..4294967295.
    NumberOfIpv4Routes interface{}

    // Number of IPv4 local addresses. The type is interface{} with range:
    // 0..4294967295.
    NumberOfIpv4LocalAddresses interface{}

    // Number of LDP configured interfaces. The type is interface{} with range:
    // 0..4294967295.
    NumberOfLdpInterfaces interface{}

    // Number of LDP IPv4 configured interfaces. The type is interface{} with
    // range: 0..4294967295.
    NumberOfIpv4ldpInterfaces interface{}
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xe"
    vrfSummary.EntityData.ParentYangName = "vrf"
    vrfSummary.EntityData.SegmentPath = "vrf-summary"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    vrfSummary.EntityData.Children = types.NewOrderedMap()
    vrfSummary.EntityData.Leafs = types.NewOrderedMap()
    vrfSummary.EntityData.Leafs.Append("address-families", types.YLeaf{"AddressFamilies", vrfSummary.AddressFamilies})
    vrfSummary.EntityData.Leafs.Append("number-of-neighbors", types.YLeaf{"NumberOfNeighbors", vrfSummary.NumberOfNeighbors})
    vrfSummary.EntityData.Leafs.Append("number-of-graceful-restart-neighbors", types.YLeaf{"NumberOfGracefulRestartNeighbors", vrfSummary.NumberOfGracefulRestartNeighbors})
    vrfSummary.EntityData.Leafs.Append("number-of-downstream-on-demand-neighbors", types.YLeaf{"NumberOfDownstreamOnDemandNeighbors", vrfSummary.NumberOfDownstreamOnDemandNeighbors})
    vrfSummary.EntityData.Leafs.Append("numberof-ipv4-hello-adj", types.YLeaf{"NumberofIpv4HelloAdj", vrfSummary.NumberofIpv4HelloAdj})
    vrfSummary.EntityData.Leafs.Append("number-of-ipv4-routes", types.YLeaf{"NumberOfIpv4Routes", vrfSummary.NumberOfIpv4Routes})
    vrfSummary.EntityData.Leafs.Append("number-of-ipv4-local-addresses", types.YLeaf{"NumberOfIpv4LocalAddresses", vrfSummary.NumberOfIpv4LocalAddresses})
    vrfSummary.EntityData.Leafs.Append("number-of-ldp-interfaces", types.YLeaf{"NumberOfLdpInterfaces", vrfSummary.NumberOfLdpInterfaces})
    vrfSummary.EntityData.Leafs.Append("number-of-ipv4ldp-interfaces", types.YLeaf{"NumberOfIpv4ldpInterfaces", vrfSummary.NumberOfIpv4ldpInterfaces})

    vrfSummary.EntityData.YListKeys = []string {}

    return &(vrfSummary.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs
// Address Family specific operational data
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP Operational data for this Address Family. The type is slice of
    // MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af.
    Af []*MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xe"
    afs.EntityData.ParentYangName = "vrf"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af
// MPLS LDP Operational data for this Address Family.
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // This container holds a summary of information across all interfaces in this
    // AF,.
    InterfaceSummary MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary

    // LDP IGP Synchronization related information.
    Igp MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xe"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("interface-summary", types.YChild{"InterfaceSummary", &af.InterfaceSummary})
    af.EntityData.Children.Append("igp", types.YChild{"Igp", &af.Igp})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary
// This container holds a summary of information
// across all interfaces in this AF,
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of known IP Interfaces. The type is interface{} with range:
    // 0..4294967295.
    KnownIpInterfaceCount interface{}

    // Number of known IP Interfaces with LDP Enabled. The type is interface{}
    // with range: 0..4294967295.
    KnownIpInterfaceLdpEnabled interface{}

    // Number of attached interfaces configured in LDP. The type is interface{}
    // with range: 0..4294967295.
    ConfiguredAttachedInterface interface{}

    // Number of TE tunnel interfaces configured in LDP. The type is interface{}
    // with range: 0..4294967295.
    ConfiguredTeInterface interface{}

    // Number of forward referenced interfaces. The type is interface{} with
    // range: 0..4294967295.
    ForwardReferences interface{}

    // Autoconfigure disabled. The type is interface{} with range: 0..4294967295.
    AutoConfigDisabled interface{}

    // Auto-configured interfaces. The type is interface{} with range:
    // 0..4294967295.
    AutoConfig interface{}

    // Auto-configured forward references. The type is interface{} with range:
    // 0..4294967295.
    AutoConfigForwardReferenceInterfaces interface{}
}

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xe"
    interfaceSummary.EntityData.ParentYangName = "af"
    interfaceSummary.EntityData.SegmentPath = "interface-summary"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaceSummary.EntityData.Children = types.NewOrderedMap()
    interfaceSummary.EntityData.Leafs = types.NewOrderedMap()
    interfaceSummary.EntityData.Leafs.Append("known-ip-interface-count", types.YLeaf{"KnownIpInterfaceCount", interfaceSummary.KnownIpInterfaceCount})
    interfaceSummary.EntityData.Leafs.Append("known-ip-interface-ldp-enabled", types.YLeaf{"KnownIpInterfaceLdpEnabled", interfaceSummary.KnownIpInterfaceLdpEnabled})
    interfaceSummary.EntityData.Leafs.Append("configured-attached-interface", types.YLeaf{"ConfiguredAttachedInterface", interfaceSummary.ConfiguredAttachedInterface})
    interfaceSummary.EntityData.Leafs.Append("configured-te-interface", types.YLeaf{"ConfiguredTeInterface", interfaceSummary.ConfiguredTeInterface})
    interfaceSummary.EntityData.Leafs.Append("forward-references", types.YLeaf{"ForwardReferences", interfaceSummary.ForwardReferences})
    interfaceSummary.EntityData.Leafs.Append("auto-config-disabled", types.YLeaf{"AutoConfigDisabled", interfaceSummary.AutoConfigDisabled})
    interfaceSummary.EntityData.Leafs.Append("auto-config", types.YLeaf{"AutoConfig", interfaceSummary.AutoConfig})
    interfaceSummary.EntityData.Leafs.Append("auto-config-forward-reference-interfaces", types.YLeaf{"AutoConfigForwardReferenceInterfaces", interfaceSummary.AutoConfigForwardReferenceInterfaces})

    interfaceSummary.EntityData.YListKeys = []string {}

    return &(interfaceSummary.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp
// LDP IGP Synchronization related information
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP-IGP Synchronization related information for an interface. The type is
    // slice of MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync.
    Sync []*MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xe"
    igp.EntityData.ParentYangName = "af"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("sync", types.YChild{"Sync", nil})
    for i := range igp.Sync {
        igp.EntityData.Children.Append(types.GetSegmentPath(igp.Sync[i]), types.YChild{"Sync", igp.Sync[i]})
    }
    igp.EntityData.Leafs = types.NewOrderedMap()

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync
// LDP-IGP Synchronization related information
// for an interface
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This leaf contains the interface name for the IGP
    // Synchronization information. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // IGP Sync state. The type is IgpSyncState.
    IgpSyncState interface{}

    // This is set when the sync delay timer running. The type is interface{}.
    IsDelayTimerRunning interface{}

    // Remaining timer (seconds) until expiry of sync delay timer. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    DelayTimerRemaining interface{}

    // Reason IGP Sync Not Achieved. The type is one of the following:
    // IgpSyncDownReasonNaIgpSyncDownReasonPeerUpdateNotReceivedIgpSyncDownReasonNoPeerSessIgpSyncDownReasonInternalIgpSyncDownReasonNoHelloAdjIgpSyncDownReasonPeerUpdateNotDone.
    IgpSyncDownReason interface{}

    // MPLS LDP IGP Sync Interface Peer Information. The type is slice of
    // MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers.
    Peers []*MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetEntityData() *types.CommonEntityData {
    sync.EntityData.YFilter = sync.YFilter
    sync.EntityData.YangName = "sync"
    sync.EntityData.BundleName = "cisco_ios_xe"
    sync.EntityData.ParentYangName = "igp"
    sync.EntityData.SegmentPath = "sync" + types.AddKeyToken(sync.Interface, "interface")
    sync.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sync.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sync.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sync.EntityData.Children = types.NewOrderedMap()
    sync.EntityData.Children.Append("peers", types.YChild{"Peers", nil})
    for i := range sync.Peers {
        sync.EntityData.Children.Append(types.GetSegmentPath(sync.Peers[i]), types.YChild{"Peers", sync.Peers[i]})
    }
    sync.EntityData.Leafs = types.NewOrderedMap()
    sync.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", sync.Interface})
    sync.EntityData.Leafs.Append("igp-sync-state", types.YLeaf{"IgpSyncState", sync.IgpSyncState})
    sync.EntityData.Leafs.Append("is-delay-timer-running", types.YLeaf{"IsDelayTimerRunning", sync.IsDelayTimerRunning})
    sync.EntityData.Leafs.Append("delay-timer-remaining", types.YLeaf{"DelayTimerRemaining", sync.DelayTimerRemaining})
    sync.EntityData.Leafs.Append("igp-sync-down-reason", types.YLeaf{"IgpSyncDownReason", sync.IgpSyncDownReason})

    sync.EntityData.YListKeys = []string {"Interface"}

    return &(sync.EntityData)
}

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers
// MPLS LDP IGP Sync Interface Peer Information
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer Identifier. The type is string.
    PeerId interface{}

    // Is GR enabled session. The type is bool.
    IsGrEnabled interface{}

    // This is set if this peer was created due to check-pointing. The type is
    // interface{}.
    IsChkptCreated interface{}
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xe"
    peers.EntityData.ParentYangName = "sync"
    peers.EntityData.SegmentPath = "peers"
    peers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    peers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    peers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    peers.EntityData.Children = types.NewOrderedMap()
    peers.EntityData.Leafs = types.NewOrderedMap()
    peers.EntityData.Leafs.Append("peer-id", types.YLeaf{"PeerId", peers.PeerId})
    peers.EntityData.Leafs.Append("is-gr-enabled", types.YLeaf{"IsGrEnabled", peers.IsGrEnabled})
    peers.EntityData.Leafs.Append("is-chkpt-created", types.YLeaf{"IsChkptCreated", peers.IsChkptCreated})

    peers.EntityData.YListKeys = []string {}

    return &(peers.EntityData)
}

// MplsLdp_MplsLdpState_Discovery
// The LDP Discovery operational state
type MplsLdp_MplsLdpState_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP Discovery Summary Information.
    DiscoveryStats MplsLdp_MplsLdpState_Discovery_DiscoveryStats

    // This container holds information for LDP Discovery using non-targeted
    // Hellos. These are interface-based hellos which form one or more adjacencies
    // for each interface and also form adjacencies on multiple intefrfaces. Link
    // Hellos can therefore form multiple adjacencies with the same peer.
    LinkHelloState MplsLdp_MplsLdpState_Discovery_LinkHelloState

    // The LDP Discovery Targeted Hello state.
    TargetedHellos MplsLdp_MplsLdpState_Discovery_TargetedHellos
}

func (discovery *MplsLdp_MplsLdpState_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xe"
    discovery.EntityData.ParentYangName = "mpls-ldp-state"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("discovery-stats", types.YChild{"DiscoveryStats", &discovery.DiscoveryStats})
    discovery.EntityData.Children.Append("link-hello-state", types.YChild{"LinkHelloState", &discovery.LinkHelloState})
    discovery.EntityData.Children.Append("targeted-hellos", types.YChild{"TargetedHellos", &discovery.TargetedHellos})
    discovery.EntityData.Leafs = types.NewOrderedMap()

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_MplsLdpState_Discovery_DiscoveryStats
// MPLS LDP Discovery Summary Information
type MplsLdp_MplsLdpState_Discovery_DiscoveryStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Number of LDP configured interfaces. The type is interface{} with
    // range: 0..4294967295.
    NumOfLdpInterfaces interface{}

    // Number of active LDP enabled interfaces. The type is interface{} with
    // range: 0..4294967295.
    NumOfActiveLdpInterfaces interface{}

    // Number of link hello discoveries in xmit state. The type is interface{}
    // with range: 0..4294967295.
    NumOfLnkDiscXmit interface{}

    // Number of targeted hello discoveries in xmit state. The type is interface{}
    // with range: 0..4294967295.
    NumOfTgtDiscXmit interface{}

    // Number of link hello discoveries in recv state. The type is interface{}
    // with range: 0..4294967295.
    NumOfLnkDiscRecv interface{}

    // Number of targeted hello discoveries in recv state. The type is interface{}
    // with range: 0..4294967295.
    NumOfTgtDiscRecv interface{}
}

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetEntityData() *types.CommonEntityData {
    discoveryStats.EntityData.YFilter = discoveryStats.YFilter
    discoveryStats.EntityData.YangName = "discovery-stats"
    discoveryStats.EntityData.BundleName = "cisco_ios_xe"
    discoveryStats.EntityData.ParentYangName = "discovery"
    discoveryStats.EntityData.SegmentPath = "discovery-stats"
    discoveryStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    discoveryStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    discoveryStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    discoveryStats.EntityData.Children = types.NewOrderedMap()
    discoveryStats.EntityData.Leafs = types.NewOrderedMap()
    discoveryStats.EntityData.Leafs.Append("num-of-ldp-interfaces", types.YLeaf{"NumOfLdpInterfaces", discoveryStats.NumOfLdpInterfaces})
    discoveryStats.EntityData.Leafs.Append("num-of-active-ldp-interfaces", types.YLeaf{"NumOfActiveLdpInterfaces", discoveryStats.NumOfActiveLdpInterfaces})
    discoveryStats.EntityData.Leafs.Append("num-of-lnk-disc-xmit", types.YLeaf{"NumOfLnkDiscXmit", discoveryStats.NumOfLnkDiscXmit})
    discoveryStats.EntityData.Leafs.Append("num-of-tgt-disc-xmit", types.YLeaf{"NumOfTgtDiscXmit", discoveryStats.NumOfTgtDiscXmit})
    discoveryStats.EntityData.Leafs.Append("num-of-lnk-disc-recv", types.YLeaf{"NumOfLnkDiscRecv", discoveryStats.NumOfLnkDiscRecv})
    discoveryStats.EntityData.Leafs.Append("num-of-tgt-disc-recv", types.YLeaf{"NumOfTgtDiscRecv", discoveryStats.NumOfTgtDiscRecv})

    discoveryStats.EntityData.YListKeys = []string {}

    return &(discoveryStats.EntityData)
}

// MplsLdp_MplsLdpState_Discovery_LinkHelloState
// This container holds information for LDP Discovery
// using non-targeted Hellos. These are interface-based
// hellos which form one or more adjacencies for each
// interface and also form adjacencies on multiple
// intefrfaces. Link Hellos can therefore form multiple
// adjacencies with the same peer.
type MplsLdp_MplsLdpState_Discovery_LinkHelloState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents a single LDP Hello Adjacency. An LDP Session can have
    // one or more Hello Adjacencies. The type is slice of
    // MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos.
    LinkHellos []*MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetEntityData() *types.CommonEntityData {
    linkHelloState.EntityData.YFilter = linkHelloState.YFilter
    linkHelloState.EntityData.YangName = "link-hello-state"
    linkHelloState.EntityData.BundleName = "cisco_ios_xe"
    linkHelloState.EntityData.ParentYangName = "discovery"
    linkHelloState.EntityData.SegmentPath = "link-hello-state"
    linkHelloState.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkHelloState.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkHelloState.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkHelloState.EntityData.Children = types.NewOrderedMap()
    linkHelloState.EntityData.Children.Append("link-hellos", types.YChild{"LinkHellos", nil})
    for i := range linkHelloState.LinkHellos {
        linkHelloState.EntityData.Children.Append(types.GetSegmentPath(linkHelloState.LinkHellos[i]), types.YChild{"LinkHellos", linkHelloState.LinkHellos[i]})
    }
    linkHelloState.EntityData.Leafs = types.NewOrderedMap()

    linkHelloState.EntityData.YListKeys = []string {}

    return &(linkHelloState.EntityData)
}

// MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos
// Each entry represents a single LDP Hello Adjacency.
// An LDP Session can have one or more Hello
// Adjacencies.
type MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The Discovery Interface. The type is string.
    // Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // This attribute is a key. This is the MPLS LDP Hello Neighbor transport
    // address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NbrTransportAddr interface{}

    // Hello interval in seconds. This is the value used to send hello messages.
    // The type is interface{} with range: 0..4294967295. Units are seconds.
    HelloInterval interface{}

    // MPLS LDP Discovery Local source address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalSrcAddr interface{}

    // MPLS LDP Discovery Local transport address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalTransportAddr interface{}

    // This is the MPLS LDP Hello Neighbor source address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NbrSrcAddr interface{}

    // Neighbor LDP Identifier. The type is string.
    NbrLdpId interface{}

    // Set when the session is up for this adjacency. The type is interface{}.
    SessionUp interface{}

    // The Hello hold time which is negotiated between the Entity and the Peer. 
    // The entity associated with this Hello Adjacency issues a proposed Hello
    // Hold Time value in the EntityHelloHoldTimer object.  The peer also proposes
    // a value and this object represents the negotiated value.  A value of 0
    // means the default, which is 15 seconds for Link Hellos and 45 seconds for
    // Targeted Hellos. A value of 65535 indicates an infinite hold time. The type
    // is interface{} with range: 0..65535.
    NbrHoldTime interface{}

    // Next hello due time in milliseconds. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    NextHello interface{}

    // This is the MPLS LDP Hello Discovery expiry time in seconds. If the value
    // of this object is 65535, this means that the hold time is infinite (i.e.,
    // wait forever).  Otherwise, the time remaining for this Hello Adjacency to
    // receive its next Hello Message.  This interval will change when the 'next'
    // Hello Message which corresponds to this Hello Adjacency is received unless
    // it is infinite. The type is interface{} with range: 0..65535. Units are
    // seconds.
    HoldTimeRemaining interface{}
}

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetEntityData() *types.CommonEntityData {
    linkHellos.EntityData.YFilter = linkHellos.YFilter
    linkHellos.EntityData.YangName = "link-hellos"
    linkHellos.EntityData.BundleName = "cisco_ios_xe"
    linkHellos.EntityData.ParentYangName = "link-hello-state"
    linkHellos.EntityData.SegmentPath = "link-hellos" + types.AddKeyToken(linkHellos.Interface, "interface") + types.AddKeyToken(linkHellos.NbrTransportAddr, "nbr-transport-addr")
    linkHellos.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkHellos.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkHellos.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkHellos.EntityData.Children = types.NewOrderedMap()
    linkHellos.EntityData.Leafs = types.NewOrderedMap()
    linkHellos.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", linkHellos.Interface})
    linkHellos.EntityData.Leafs.Append("nbr-transport-addr", types.YLeaf{"NbrTransportAddr", linkHellos.NbrTransportAddr})
    linkHellos.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", linkHellos.HelloInterval})
    linkHellos.EntityData.Leafs.Append("local-src-addr", types.YLeaf{"LocalSrcAddr", linkHellos.LocalSrcAddr})
    linkHellos.EntityData.Leafs.Append("local-transport-addr", types.YLeaf{"LocalTransportAddr", linkHellos.LocalTransportAddr})
    linkHellos.EntityData.Leafs.Append("nbr-src-addr", types.YLeaf{"NbrSrcAddr", linkHellos.NbrSrcAddr})
    linkHellos.EntityData.Leafs.Append("nbr-ldp-id", types.YLeaf{"NbrLdpId", linkHellos.NbrLdpId})
    linkHellos.EntityData.Leafs.Append("session-up", types.YLeaf{"SessionUp", linkHellos.SessionUp})
    linkHellos.EntityData.Leafs.Append("nbr-hold-time", types.YLeaf{"NbrHoldTime", linkHellos.NbrHoldTime})
    linkHellos.EntityData.Leafs.Append("next-hello", types.YLeaf{"NextHello", linkHellos.NextHello})
    linkHellos.EntityData.Leafs.Append("hold-time-remaining", types.YLeaf{"HoldTimeRemaining", linkHellos.HoldTimeRemaining})

    linkHellos.EntityData.YListKeys = []string {"Interface", "NbrTransportAddr"}

    return &(linkHellos.EntityData)
}

// MplsLdp_MplsLdpState_Discovery_TargetedHellos
// The LDP Discovery Targeted Hello state.
type MplsLdp_MplsLdpState_Discovery_TargetedHellos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local Targeted Hello interval in seconds. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    TargetedHelloInterval interface{}

    // Local Targeted hold time in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds.
    TargetedHelloHoldTime interface{}

    // The LDP targeted discovery information for a specific target. Targetted
    // discovery creates a single adjacency between two addresses and not
    // indiviual adjacencies across physical interfaces. The type is slice of
    // MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello.
    TargetedHello []*MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetEntityData() *types.CommonEntityData {
    targetedHellos.EntityData.YFilter = targetedHellos.YFilter
    targetedHellos.EntityData.YangName = "targeted-hellos"
    targetedHellos.EntityData.BundleName = "cisco_ios_xe"
    targetedHellos.EntityData.ParentYangName = "discovery"
    targetedHellos.EntityData.SegmentPath = "targeted-hellos"
    targetedHellos.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    targetedHellos.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    targetedHellos.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    targetedHellos.EntityData.Children = types.NewOrderedMap()
    targetedHellos.EntityData.Children.Append("targeted-hello", types.YChild{"TargetedHello", nil})
    for i := range targetedHellos.TargetedHello {
        targetedHellos.EntityData.Children.Append(types.GetSegmentPath(targetedHellos.TargetedHello[i]), types.YChild{"TargetedHello", targetedHellos.TargetedHello[i]})
    }
    targetedHellos.EntityData.Leafs = types.NewOrderedMap()
    targetedHellos.EntityData.Leafs.Append("targeted-hello-interval", types.YLeaf{"TargetedHelloInterval", targetedHellos.TargetedHelloInterval})
    targetedHellos.EntityData.Leafs.Append("targeted-hello-hold-time", types.YLeaf{"TargetedHelloHoldTime", targetedHellos.TargetedHelloHoldTime})

    targetedHellos.EntityData.YListKeys = []string {}

    return &(targetedHellos.EntityData)
}

// MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello
// The LDP targeted discovery information for a specific
// target. Targetted discovery creates a single adjacency
// between two addresses and not indiviual adjacencies
// across physical interfaces.
type MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. The target IP Address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    TargetAddress interface{}

    // Local IP Address. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Neighbor LDP Identifier. The type is string.
    NeighborLdpIdentifier interface{}

    // This is the MPLS LDP Targeted Hello state. The type is DhcState.
    State interface{}

    // The Hello hold time which is negotiated between the Entity and the Peer. 
    // The entity associated with this Hello Adjacency issues a proposed Hello
    // Hold Time value in the EntityHelloHoldTimer object.  The peer also proposes
    // a value and this object represents the negotiated value.  A value of 0
    // means the default, which is 15 seconds for Link Hellos and 45 seconds for
    // Targeted Hellos. A value of 65535 indicates an infinite hold time. The type
    // is interface{} with range: 0..65535.
    NbrHoldTime interface{}

    // Next hello due time in milliseconds. The type is interface{} with range:
    // 0..4294967295. Units are milliseconds.
    NextHello interface{}

    // This is the MPLS LDP Hello Discovery expiry time in seconds. If the value
    // of this object is 65535, this means that the hold time is infinite (i.e.,
    // wait forever).  Otherwise, the time remaining for this Hello Adjacency to
    // receive its next Hello Message.  This interval will change when the 'next'
    // Hello Message which corresponds to this Hello Adjacency is received unless
    // it is infinite. The type is interface{} with range: 0..65535. Units are
    // seconds.
    HoldTimeRemaining interface{}
}

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetEntityData() *types.CommonEntityData {
    targetedHello.EntityData.YFilter = targetedHello.YFilter
    targetedHello.EntityData.YangName = "targeted-hello"
    targetedHello.EntityData.BundleName = "cisco_ios_xe"
    targetedHello.EntityData.ParentYangName = "targeted-hellos"
    targetedHello.EntityData.SegmentPath = "targeted-hello" + types.AddKeyToken(targetedHello.VrfName, "vrf-name") + types.AddKeyToken(targetedHello.TargetAddress, "target-address")
    targetedHello.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    targetedHello.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    targetedHello.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    targetedHello.EntityData.Children = types.NewOrderedMap()
    targetedHello.EntityData.Leafs = types.NewOrderedMap()
    targetedHello.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", targetedHello.VrfName})
    targetedHello.EntityData.Leafs.Append("target-address", types.YLeaf{"TargetAddress", targetedHello.TargetAddress})
    targetedHello.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", targetedHello.LocalAddress})
    targetedHello.EntityData.Leafs.Append("neighbor-ldp-identifier", types.YLeaf{"NeighborLdpIdentifier", targetedHello.NeighborLdpIdentifier})
    targetedHello.EntityData.Leafs.Append("state", types.YLeaf{"State", targetedHello.State})
    targetedHello.EntityData.Leafs.Append("nbr-hold-time", types.YLeaf{"NbrHoldTime", targetedHello.NbrHoldTime})
    targetedHello.EntityData.Leafs.Append("next-hello", types.YLeaf{"NextHello", targetedHello.NextHello})
    targetedHello.EntityData.Leafs.Append("hold-time-remaining", types.YLeaf{"HoldTimeRemaining", targetedHello.HoldTimeRemaining})

    targetedHello.EntityData.YListKeys = []string {"VrfName", "TargetAddress"}

    return &(targetedHello.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding
// Summary information regarding LDP forwarding
// setup and detailed LDP Forwarding rewrites
type MplsLdp_MplsLdpState_Forwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of forwarding info for this VRF.
    ForwardingVrfSumms MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms

    // This leaf contain the individual LDP forwarding rewrite for a single
    // prefix. The type is slice of
    // MplsLdp_MplsLdpState_Forwarding_ForwardingDetail.
    ForwardingDetail []*MplsLdp_MplsLdpState_Forwarding_ForwardingDetail
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetEntityData() *types.CommonEntityData {
    forwarding.EntityData.YFilter = forwarding.YFilter
    forwarding.EntityData.YangName = "forwarding"
    forwarding.EntityData.BundleName = "cisco_ios_xe"
    forwarding.EntityData.ParentYangName = "mpls-ldp-state"
    forwarding.EntityData.SegmentPath = "forwarding"
    forwarding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwarding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwarding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwarding.EntityData.Children = types.NewOrderedMap()
    forwarding.EntityData.Children.Append("forwarding-vrf-summs", types.YChild{"ForwardingVrfSumms", &forwarding.ForwardingVrfSumms})
    forwarding.EntityData.Children.Append("forwarding-detail", types.YChild{"ForwardingDetail", nil})
    for i := range forwarding.ForwardingDetail {
        forwarding.EntityData.Children.Append(types.GetSegmentPath(forwarding.ForwardingDetail[i]), types.YChild{"ForwardingDetail", forwarding.ForwardingDetail[i]})
    }
    forwarding.EntityData.Leafs = types.NewOrderedMap()

    forwarding.EntityData.YListKeys = []string {}

    return &(forwarding.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms
// Summary of forwarding info for this VRF.
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary of forwarding info for this VRF. The type is slice of
    // MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm.
    ForwardingVrfSumm []*MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetEntityData() *types.CommonEntityData {
    forwardingVrfSumms.EntityData.YFilter = forwardingVrfSumms.YFilter
    forwardingVrfSumms.EntityData.YangName = "forwarding-vrf-summs"
    forwardingVrfSumms.EntityData.BundleName = "cisco_ios_xe"
    forwardingVrfSumms.EntityData.ParentYangName = "forwarding"
    forwardingVrfSumms.EntityData.SegmentPath = "forwarding-vrf-summs"
    forwardingVrfSumms.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwardingVrfSumms.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwardingVrfSumms.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwardingVrfSumms.EntityData.Children = types.NewOrderedMap()
    forwardingVrfSumms.EntityData.Children.Append("forwarding-vrf-summ", types.YChild{"ForwardingVrfSumm", nil})
    for i := range forwardingVrfSumms.ForwardingVrfSumm {
        forwardingVrfSumms.EntityData.Children.Append(types.GetSegmentPath(forwardingVrfSumms.ForwardingVrfSumm[i]), types.YChild{"ForwardingVrfSumm", forwardingVrfSumms.ForwardingVrfSumm[i]})
    }
    forwardingVrfSumms.EntityData.Leafs = types.NewOrderedMap()

    forwardingVrfSumms.EntityData.YListKeys = []string {}

    return &(forwardingVrfSumms.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm
// Summary of forwarding info for this VRF.
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // MPLS forwarding enabled interface count. The type is interface{} with
    // range: 0..65535.
    IntfsFwdCount interface{}

    // Local label allocated count. The type is interface{} with range: 0..65535.
    LocalLbls interface{}

    // MPLS LDP forwarding prefix rewrite summary.
    Pfxs MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs

    // MPLS LDP forwarding rewrite next-hop/path summary.
    Nhs MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs
}

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetEntityData() *types.CommonEntityData {
    forwardingVrfSumm.EntityData.YFilter = forwardingVrfSumm.YFilter
    forwardingVrfSumm.EntityData.YangName = "forwarding-vrf-summ"
    forwardingVrfSumm.EntityData.BundleName = "cisco_ios_xe"
    forwardingVrfSumm.EntityData.ParentYangName = "forwarding-vrf-summs"
    forwardingVrfSumm.EntityData.SegmentPath = "forwarding-vrf-summ" + types.AddKeyToken(forwardingVrfSumm.VrfName, "vrf-name")
    forwardingVrfSumm.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwardingVrfSumm.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwardingVrfSumm.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwardingVrfSumm.EntityData.Children = types.NewOrderedMap()
    forwardingVrfSumm.EntityData.Children.Append("pfxs", types.YChild{"Pfxs", &forwardingVrfSumm.Pfxs})
    forwardingVrfSumm.EntityData.Children.Append("nhs", types.YChild{"Nhs", &forwardingVrfSumm.Nhs})
    forwardingVrfSumm.EntityData.Leafs = types.NewOrderedMap()
    forwardingVrfSumm.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", forwardingVrfSumm.VrfName})
    forwardingVrfSumm.EntityData.Leafs.Append("intfs-fwd-count", types.YLeaf{"IntfsFwdCount", forwardingVrfSumm.IntfsFwdCount})
    forwardingVrfSumm.EntityData.Leafs.Append("local-lbls", types.YLeaf{"LocalLbls", forwardingVrfSumm.LocalLbls})

    forwardingVrfSumm.EntityData.YListKeys = []string {"VrfName"}

    return &(forwardingVrfSumm.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs
// MPLS LDP forwarding prefix rewrite summary
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total Prefix count. The type is interface{} with range: 0..65535.
    TotalPfxs interface{}

    // Count of prefixes with ECMP. The type is interface{} with range: 0..65535.
    EcmpPfxs interface{}

    // Count of FRR protected prefixes. The type is interface{} with range:
    // 0..65535.
    ProtectedPfxs interface{}

    // Labeled prefix count for all paths.
    LabeledPfxsAggr MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr

    // Labeled prefix count related to primary paths only.
    LabeledPfxsPrimary MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary

    // Labeled prefix count related to backup paths only.
    LabeledPfxsBackup MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup
}

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetEntityData() *types.CommonEntityData {
    pfxs.EntityData.YFilter = pfxs.YFilter
    pfxs.EntityData.YangName = "pfxs"
    pfxs.EntityData.BundleName = "cisco_ios_xe"
    pfxs.EntityData.ParentYangName = "forwarding-vrf-summ"
    pfxs.EntityData.SegmentPath = "pfxs"
    pfxs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    pfxs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    pfxs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    pfxs.EntityData.Children = types.NewOrderedMap()
    pfxs.EntityData.Children.Append("labeled-pfxs-aggr", types.YChild{"LabeledPfxsAggr", &pfxs.LabeledPfxsAggr})
    pfxs.EntityData.Children.Append("labeled-pfxs-primary", types.YChild{"LabeledPfxsPrimary", &pfxs.LabeledPfxsPrimary})
    pfxs.EntityData.Children.Append("labeled-pfxs-backup", types.YChild{"LabeledPfxsBackup", &pfxs.LabeledPfxsBackup})
    pfxs.EntityData.Leafs = types.NewOrderedMap()
    pfxs.EntityData.Leafs.Append("total-pfxs", types.YLeaf{"TotalPfxs", pfxs.TotalPfxs})
    pfxs.EntityData.Leafs.Append("ecmp-pfxs", types.YLeaf{"EcmpPfxs", pfxs.EcmpPfxs})
    pfxs.EntityData.Leafs.Append("protected-pfxs", types.YLeaf{"ProtectedPfxs", pfxs.ProtectedPfxs})

    pfxs.EntityData.YListKeys = []string {}

    return &(pfxs.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr
// Labeled prefix count for all paths
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of labeled prefixes with 1 or more paths labeled. The type is
    // interface{} with range: 0..65535.
    LabeledPfxs interface{}

    // Count of labeled prefixes with some (but not ALL) paths labeled. The type
    // is interface{} with range: 0..65535.
    LabeledPfxsPartial interface{}

    // Count of labeled prefixes with ALL paths unlabeled. The type is interface{}
    // with range: 0..65535.
    UnlabeledPfxs interface{}
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetEntityData() *types.CommonEntityData {
    labeledPfxsAggr.EntityData.YFilter = labeledPfxsAggr.YFilter
    labeledPfxsAggr.EntityData.YangName = "labeled-pfxs-aggr"
    labeledPfxsAggr.EntityData.BundleName = "cisco_ios_xe"
    labeledPfxsAggr.EntityData.ParentYangName = "pfxs"
    labeledPfxsAggr.EntityData.SegmentPath = "labeled-pfxs-aggr"
    labeledPfxsAggr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labeledPfxsAggr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labeledPfxsAggr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labeledPfxsAggr.EntityData.Children = types.NewOrderedMap()
    labeledPfxsAggr.EntityData.Leafs = types.NewOrderedMap()
    labeledPfxsAggr.EntityData.Leafs.Append("labeled-pfxs", types.YLeaf{"LabeledPfxs", labeledPfxsAggr.LabeledPfxs})
    labeledPfxsAggr.EntityData.Leafs.Append("labeled-pfxs-partial", types.YLeaf{"LabeledPfxsPartial", labeledPfxsAggr.LabeledPfxsPartial})
    labeledPfxsAggr.EntityData.Leafs.Append("unlabeled-pfxs", types.YLeaf{"UnlabeledPfxs", labeledPfxsAggr.UnlabeledPfxs})

    labeledPfxsAggr.EntityData.YListKeys = []string {}

    return &(labeledPfxsAggr.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary
// Labeled prefix count related to primary paths
// only
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of labeled prefixes with 1 or more paths labeled. The type is
    // interface{} with range: 0..65535.
    LabeledPfxs interface{}

    // Count of labeled prefixes with some (but not ALL) paths labeled. The type
    // is interface{} with range: 0..65535.
    LabeledPfxsPartial interface{}

    // Count of labeled prefixes with ALL paths unlabeled. The type is interface{}
    // with range: 0..65535.
    UnlabeledPfxs interface{}
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetEntityData() *types.CommonEntityData {
    labeledPfxsPrimary.EntityData.YFilter = labeledPfxsPrimary.YFilter
    labeledPfxsPrimary.EntityData.YangName = "labeled-pfxs-primary"
    labeledPfxsPrimary.EntityData.BundleName = "cisco_ios_xe"
    labeledPfxsPrimary.EntityData.ParentYangName = "pfxs"
    labeledPfxsPrimary.EntityData.SegmentPath = "labeled-pfxs-primary"
    labeledPfxsPrimary.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labeledPfxsPrimary.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labeledPfxsPrimary.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labeledPfxsPrimary.EntityData.Children = types.NewOrderedMap()
    labeledPfxsPrimary.EntityData.Leafs = types.NewOrderedMap()
    labeledPfxsPrimary.EntityData.Leafs.Append("labeled-pfxs", types.YLeaf{"LabeledPfxs", labeledPfxsPrimary.LabeledPfxs})
    labeledPfxsPrimary.EntityData.Leafs.Append("labeled-pfxs-partial", types.YLeaf{"LabeledPfxsPartial", labeledPfxsPrimary.LabeledPfxsPartial})
    labeledPfxsPrimary.EntityData.Leafs.Append("unlabeled-pfxs", types.YLeaf{"UnlabeledPfxs", labeledPfxsPrimary.UnlabeledPfxs})

    labeledPfxsPrimary.EntityData.YListKeys = []string {}

    return &(labeledPfxsPrimary.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup
// Labeled prefix count related to backup paths
// only
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Count of labeled prefixes with 1 or more paths labeled. The type is
    // interface{} with range: 0..65535.
    LabeledPfxs interface{}

    // Count of labeled prefixes with some (but not ALL) paths labeled. The type
    // is interface{} with range: 0..65535.
    LabeledPfxsPartial interface{}

    // Count of labeled prefixes with ALL paths unlabeled. The type is interface{}
    // with range: 0..65535.
    UnlabeledPfxs interface{}
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetEntityData() *types.CommonEntityData {
    labeledPfxsBackup.EntityData.YFilter = labeledPfxsBackup.YFilter
    labeledPfxsBackup.EntityData.YangName = "labeled-pfxs-backup"
    labeledPfxsBackup.EntityData.BundleName = "cisco_ios_xe"
    labeledPfxsBackup.EntityData.ParentYangName = "pfxs"
    labeledPfxsBackup.EntityData.SegmentPath = "labeled-pfxs-backup"
    labeledPfxsBackup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labeledPfxsBackup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labeledPfxsBackup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labeledPfxsBackup.EntityData.Children = types.NewOrderedMap()
    labeledPfxsBackup.EntityData.Leafs = types.NewOrderedMap()
    labeledPfxsBackup.EntityData.Leafs.Append("labeled-pfxs", types.YLeaf{"LabeledPfxs", labeledPfxsBackup.LabeledPfxs})
    labeledPfxsBackup.EntityData.Leafs.Append("labeled-pfxs-partial", types.YLeaf{"LabeledPfxsPartial", labeledPfxsBackup.LabeledPfxsPartial})
    labeledPfxsBackup.EntityData.Leafs.Append("unlabeled-pfxs", types.YLeaf{"UnlabeledPfxs", labeledPfxsBackup.UnlabeledPfxs})

    labeledPfxsBackup.EntityData.YListKeys = []string {}

    return &(labeledPfxsBackup.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs
// MPLS LDP forwarding rewrite next-hop/path summary
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total path count. The type is interface{} with range: 0..4294967295.
    TotalPaths interface{}

    // Count of FRR protected paths. The type is interface{} with range:
    // 0..4294967295.
    ProtectedPaths interface{}

    // Count of non-primary backup paths. The type is interface{} with range:
    // 0..4294967295.
    BackupPaths interface{}

    // Count of non-primary remote backup paths. The type is interface{} with
    // range: 0..4294967295.
    RemoteBackupPaths interface{}

    // Count of all labeled paths. The type is interface{} with range:
    // 0..4294967295.
    LabeledPaths interface{}

    // Count of labeled backup paths. The type is interface{} with range:
    // 0..4294967295.
    LabeledBackupPaths interface{}
}

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetEntityData() *types.CommonEntityData {
    nhs.EntityData.YFilter = nhs.YFilter
    nhs.EntityData.YangName = "nhs"
    nhs.EntityData.BundleName = "cisco_ios_xe"
    nhs.EntityData.ParentYangName = "forwarding-vrf-summ"
    nhs.EntityData.SegmentPath = "nhs"
    nhs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nhs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nhs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nhs.EntityData.Children = types.NewOrderedMap()
    nhs.EntityData.Leafs = types.NewOrderedMap()
    nhs.EntityData.Leafs.Append("total-paths", types.YLeaf{"TotalPaths", nhs.TotalPaths})
    nhs.EntityData.Leafs.Append("protected-paths", types.YLeaf{"ProtectedPaths", nhs.ProtectedPaths})
    nhs.EntityData.Leafs.Append("backup-paths", types.YLeaf{"BackupPaths", nhs.BackupPaths})
    nhs.EntityData.Leafs.Append("remote-backup-paths", types.YLeaf{"RemoteBackupPaths", nhs.RemoteBackupPaths})
    nhs.EntityData.Leafs.Append("labeled-paths", types.YLeaf{"LabeledPaths", nhs.LabeledPaths})
    nhs.EntityData.Leafs.Append("labeled-backup-paths", types.YLeaf{"LabeledBackupPaths", nhs.LabeledBackupPaths})

    nhs.EntityData.YListKeys = []string {}

    return &(nhs.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail
// This leaf contain the individual LDP forwarding rewrite
// for a single prefix.
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. The IP Prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // This is the MPLS LDP Forward IP Prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FwdPrefix interface{}

    // Table ID associated with IP prefix. The type is interface{} with range:
    // 0..4294967295.
    TableId interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLength interface{}

    // MPLS LDP Forwarding Route information.
    Route MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route

    // MPLS LDP Forwarding Path info. The type is slice of
    // MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths.
    Paths []*MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetEntityData() *types.CommonEntityData {
    forwardingDetail.EntityData.YFilter = forwardingDetail.YFilter
    forwardingDetail.EntityData.YangName = "forwarding-detail"
    forwardingDetail.EntityData.BundleName = "cisco_ios_xe"
    forwardingDetail.EntityData.ParentYangName = "forwarding"
    forwardingDetail.EntityData.SegmentPath = "forwarding-detail" + types.AddKeyToken(forwardingDetail.VrfName, "vrf-name") + types.AddKeyToken(forwardingDetail.Prefix, "prefix")
    forwardingDetail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    forwardingDetail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    forwardingDetail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    forwardingDetail.EntityData.Children = types.NewOrderedMap()
    forwardingDetail.EntityData.Children.Append("route", types.YChild{"Route", &forwardingDetail.Route})
    forwardingDetail.EntityData.Children.Append("paths", types.YChild{"Paths", nil})
    for i := range forwardingDetail.Paths {
        forwardingDetail.EntityData.Children.Append(types.GetSegmentPath(forwardingDetail.Paths[i]), types.YChild{"Paths", forwardingDetail.Paths[i]})
    }
    forwardingDetail.EntityData.Leafs = types.NewOrderedMap()
    forwardingDetail.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", forwardingDetail.VrfName})
    forwardingDetail.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", forwardingDetail.Prefix})
    forwardingDetail.EntityData.Leafs.Append("fwd-prefix", types.YLeaf{"FwdPrefix", forwardingDetail.FwdPrefix})
    forwardingDetail.EntityData.Leafs.Append("table-id", types.YLeaf{"TableId", forwardingDetail.TableId})
    forwardingDetail.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", forwardingDetail.PrefixLength})

    forwardingDetail.EntityData.YListKeys = []string {"VrfName", "Prefix"}

    return &(forwardingDetail.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route
// MPLS LDP Forwarding Route information
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Route RIB version. The type is interface{} with range: 0..4294967295.
    Version interface{}

    // Route priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Route source protocol Id. The type is interface{} with range: 0..65535.
    Source interface{}

    // Route type. The type is interface{} with range: 0..65535.
    Type interface{}

    // Route metric. The type is interface{} with range: 0..4294967295.
    Metric interface{}

    // Is this route leaked across local VRFs?. The type is bool.
    IsLocalVrfLeaked interface{}

    // Number of routing updates. The type is interface{} with range:
    // 0..4294967295.
    RoutingUpdateCount interface{}

    // Last Routing update nanosec timestamp. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanoseconds.
    RoutingUpdateTimestamp interface{}

    // Last Routing update nanosec age. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanoseconds.
    RoutingUpdateAge interface{}

    // Local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Number of forwarding updates. The type is interface{} with range:
    // 0..4294967295.
    ForwardingUpdateCount interface{}

    // Last Forwarding update nanosec timestamp. The type is interface{} with
    // range: 0..18446744073709551615. Units are nanoseconds.
    ForwardingUpdateTimestamp interface{}

    // Last Forwarding update nanosec age. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanoseconds.
    ForwardingUpdateAge interface{}
}

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetEntityData() *types.CommonEntityData {
    route.EntityData.YFilter = route.YFilter
    route.EntityData.YangName = "route"
    route.EntityData.BundleName = "cisco_ios_xe"
    route.EntityData.ParentYangName = "forwarding-detail"
    route.EntityData.SegmentPath = "route"
    route.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    route.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    route.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    route.EntityData.Children = types.NewOrderedMap()
    route.EntityData.Leafs = types.NewOrderedMap()
    route.EntityData.Leafs.Append("version", types.YLeaf{"Version", route.Version})
    route.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", route.Priority})
    route.EntityData.Leafs.Append("source", types.YLeaf{"Source", route.Source})
    route.EntityData.Leafs.Append("type", types.YLeaf{"Type", route.Type})
    route.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", route.Metric})
    route.EntityData.Leafs.Append("is-local-vrf-leaked", types.YLeaf{"IsLocalVrfLeaked", route.IsLocalVrfLeaked})
    route.EntityData.Leafs.Append("routing-update-count", types.YLeaf{"RoutingUpdateCount", route.RoutingUpdateCount})
    route.EntityData.Leafs.Append("routing-update-timestamp", types.YLeaf{"RoutingUpdateTimestamp", route.RoutingUpdateTimestamp})
    route.EntityData.Leafs.Append("routing-update-age", types.YLeaf{"RoutingUpdateAge", route.RoutingUpdateAge})
    route.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", route.LocalLabel})
    route.EntityData.Leafs.Append("forwarding-update-count", types.YLeaf{"ForwardingUpdateCount", route.ForwardingUpdateCount})
    route.EntityData.Leafs.Append("forwarding-update-timestamp", types.YLeaf{"ForwardingUpdateTimestamp", route.ForwardingUpdateTimestamp})
    route.EntityData.Leafs.Append("forwarding-update-age", types.YLeaf{"ForwardingUpdateAge", route.ForwardingUpdateAge})

    route.EntityData.YListKeys = []string {}

    return &(route.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths
// MPLS LDP Forwarding Path info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP Forwarding Path IP Routing information.
    Routing MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing

    // MPLS LDP Forwarding Path MPLS information.
    Mpls MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xe"
    paths.EntityData.ParentYangName = "forwarding-detail"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Children.Append("routing", types.YChild{"Routing", &paths.Routing})
    paths.EntityData.Children.Append("mpls", types.YChild{"Mpls", &paths.Mpls})
    paths.EntityData.Leafs = types.NewOrderedMap()

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing
// MPLS LDP Forwarding Path IP Routing information
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is the Next Hop address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NextHop interface{}

    // This is the Remote/PQ node address. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteNodeId interface{}

    // This is true if the path has a remote LFA backup. The type is bool.
    HasRemoteLfaBkup interface{}

    // This is the interface. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // This is set when the nexthop is overriden by LDP. The type is interface{}.
    NhIsOverriden interface{}

    // Nexthop Identifier. The type is interface{} with range: 0..4294967295.
    NexthopId interface{}

    // Table ID for nexthop address. The type is interface{} with range:
    // 0..4294967295.
    NextHopTableId interface{}

    // Path's load metric for load balancing. The type is interface{} with range:
    // 0..4294967295.
    LoadMetric interface{}

    // path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Backup path Id. The type is interface{} with range: 0..255.
    BkupPathId interface{}

    // Routing path type. The type is one of the following:
    // RoutePathIpNoFlagRoutePathIpBackupRoutePathIpBgpBackupRoutePathIpProtectedRoutePathIpBackupRemote.
    PathType interface{}
}

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetEntityData() *types.CommonEntityData {
    routing.EntityData.YFilter = routing.YFilter
    routing.EntityData.YangName = "routing"
    routing.EntityData.BundleName = "cisco_ios_xe"
    routing.EntityData.ParentYangName = "paths"
    routing.EntityData.SegmentPath = "routing"
    routing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routing.EntityData.Children = types.NewOrderedMap()
    routing.EntityData.Leafs = types.NewOrderedMap()
    routing.EntityData.Leafs.Append("next-hop", types.YLeaf{"NextHop", routing.NextHop})
    routing.EntityData.Leafs.Append("remote-node-id", types.YLeaf{"RemoteNodeId", routing.RemoteNodeId})
    routing.EntityData.Leafs.Append("has-remote-lfa-bkup", types.YLeaf{"HasRemoteLfaBkup", routing.HasRemoteLfaBkup})
    routing.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", routing.Interface})
    routing.EntityData.Leafs.Append("nh-is-overriden", types.YLeaf{"NhIsOverriden", routing.NhIsOverriden})
    routing.EntityData.Leafs.Append("nexthop-id", types.YLeaf{"NexthopId", routing.NexthopId})
    routing.EntityData.Leafs.Append("next-hop-table-id", types.YLeaf{"NextHopTableId", routing.NextHopTableId})
    routing.EntityData.Leafs.Append("load-metric", types.YLeaf{"LoadMetric", routing.LoadMetric})
    routing.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", routing.PathId})
    routing.EntityData.Leafs.Append("bkup-path-id", types.YLeaf{"BkupPathId", routing.BkupPathId})
    routing.EntityData.Leafs.Append("path-type", types.YLeaf{"PathType", routing.PathType})

    routing.EntityData.YListKeys = []string {}

    return &(routing.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls
// MPLS LDP Forwarding Path MPLS information
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS nexthop info.
    MplsOutgoingInfo MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo

    // MPLS LDP Forwarding Path Remote LFA-FRR backup MPLS info.
    RemoteLfa MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetEntityData() *types.CommonEntityData {
    mpls.EntityData.YFilter = mpls.YFilter
    mpls.EntityData.YangName = "mpls"
    mpls.EntityData.BundleName = "cisco_ios_xe"
    mpls.EntityData.ParentYangName = "paths"
    mpls.EntityData.SegmentPath = "mpls"
    mpls.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mpls.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mpls.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mpls.EntityData.Children = types.NewOrderedMap()
    mpls.EntityData.Children.Append("mpls-outgoing-info", types.YChild{"MplsOutgoingInfo", &mpls.MplsOutgoingInfo})
    mpls.EntityData.Children.Append("remote-lfa", types.YChild{"RemoteLfa", &mpls.RemoteLfa})
    mpls.EntityData.Leafs = types.NewOrderedMap()

    mpls.EntityData.YListKeys = []string {}

    return &(mpls.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo
// MPLS nexthop info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outgoing label. The type is interface{} with range: 0..4294967295.
    OutLabel interface{}

    // Outgoing Label Type. The type is one of the following:
    // LabelTypeUnknownLabelTypeMplsLabelTypeUnLabeled.
    OutLabelType interface{}

    // Outgoing label owner. The type is one of the following:
    // RoutePathLblOwnerBgpRoutePathLblOwnerStaticRoutePathLblOwnerNoneRoutePathLblOwnerLdp.
    OutLabelOwner interface{}

    // Is from a GR neighbor. The type is bool.
    IsFromGracefulRestartableNeighbor interface{}

    // Is the entry stale? This may happen during a graceful restart. The type is
    // bool.
    IsStale interface{}

    // Nexthop LDP peer.
    NexthopPeerLdpIdent MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetEntityData() *types.CommonEntityData {
    mplsOutgoingInfo.EntityData.YFilter = mplsOutgoingInfo.YFilter
    mplsOutgoingInfo.EntityData.YangName = "mpls-outgoing-info"
    mplsOutgoingInfo.EntityData.BundleName = "cisco_ios_xe"
    mplsOutgoingInfo.EntityData.ParentYangName = "mpls"
    mplsOutgoingInfo.EntityData.SegmentPath = "mpls-outgoing-info"
    mplsOutgoingInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsOutgoingInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsOutgoingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsOutgoingInfo.EntityData.Children = types.NewOrderedMap()
    mplsOutgoingInfo.EntityData.Children.Append("nexthop-peer-ldp-ident", types.YChild{"NexthopPeerLdpIdent", &mplsOutgoingInfo.NexthopPeerLdpIdent})
    mplsOutgoingInfo.EntityData.Leafs = types.NewOrderedMap()
    mplsOutgoingInfo.EntityData.Leafs.Append("out-label", types.YLeaf{"OutLabel", mplsOutgoingInfo.OutLabel})
    mplsOutgoingInfo.EntityData.Leafs.Append("out-label-type", types.YLeaf{"OutLabelType", mplsOutgoingInfo.OutLabelType})
    mplsOutgoingInfo.EntityData.Leafs.Append("out-label-owner", types.YLeaf{"OutLabelOwner", mplsOutgoingInfo.OutLabelOwner})
    mplsOutgoingInfo.EntityData.Leafs.Append("is-from-graceful-restartable-neighbor", types.YLeaf{"IsFromGracefulRestartableNeighbor", mplsOutgoingInfo.IsFromGracefulRestartableNeighbor})
    mplsOutgoingInfo.EntityData.Leafs.Append("is-stale", types.YLeaf{"IsStale", mplsOutgoingInfo.IsStale})

    mplsOutgoingInfo.EntityData.YListKeys = []string {}

    return &(mplsOutgoingInfo.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent
// Nexthop LDP peer
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSR identifier. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Label space identifier. The type is interface{} with range: 0..65535.
    LabelSpaceId interface{}
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetEntityData() *types.CommonEntityData {
    nexthopPeerLdpIdent.EntityData.YFilter = nexthopPeerLdpIdent.YFilter
    nexthopPeerLdpIdent.EntityData.YangName = "nexthop-peer-ldp-ident"
    nexthopPeerLdpIdent.EntityData.BundleName = "cisco_ios_xe"
    nexthopPeerLdpIdent.EntityData.ParentYangName = "mpls-outgoing-info"
    nexthopPeerLdpIdent.EntityData.SegmentPath = "nexthop-peer-ldp-ident"
    nexthopPeerLdpIdent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nexthopPeerLdpIdent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nexthopPeerLdpIdent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nexthopPeerLdpIdent.EntityData.Children = types.NewOrderedMap()
    nexthopPeerLdpIdent.EntityData.Leafs = types.NewOrderedMap()
    nexthopPeerLdpIdent.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", nexthopPeerLdpIdent.LsrId})
    nexthopPeerLdpIdent.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", nexthopPeerLdpIdent.LabelSpaceId})

    nexthopPeerLdpIdent.EntityData.YListKeys = []string {}

    return &(nexthopPeerLdpIdent.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa
// MPLS LDP Forwarding Path Remote LFA-FRR backup
// MPLS info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether path has remote LFA backup. The type is bool.
    HasRemoteLfaBkup interface{}

    // Remote LFA MPLS nexthop info.
    MplsOutgoingInfo MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetEntityData() *types.CommonEntityData {
    remoteLfa.EntityData.YFilter = remoteLfa.YFilter
    remoteLfa.EntityData.YangName = "remote-lfa"
    remoteLfa.EntityData.BundleName = "cisco_ios_xe"
    remoteLfa.EntityData.ParentYangName = "mpls"
    remoteLfa.EntityData.SegmentPath = "remote-lfa"
    remoteLfa.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    remoteLfa.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    remoteLfa.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    remoteLfa.EntityData.Children = types.NewOrderedMap()
    remoteLfa.EntityData.Children.Append("mpls-outgoing-info", types.YChild{"MplsOutgoingInfo", &remoteLfa.MplsOutgoingInfo})
    remoteLfa.EntityData.Leafs = types.NewOrderedMap()
    remoteLfa.EntityData.Leafs.Append("has-remote-lfa-bkup", types.YLeaf{"HasRemoteLfaBkup", remoteLfa.HasRemoteLfaBkup})

    remoteLfa.EntityData.YListKeys = []string {}

    return &(remoteLfa.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo
// Remote LFA MPLS nexthop info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outgoing label. The type is interface{} with range: 0..4294967295.
    OutLabel interface{}

    // Outgoing Label Type. The type is one of the following:
    // LabelTypeUnknownLabelTypeMplsLabelTypeUnLabeled.
    OutLabelType interface{}

    // Outgoing label owner. The type is one of the following:
    // RoutePathLblOwnerBgpRoutePathLblOwnerStaticRoutePathLblOwnerNoneRoutePathLblOwnerLdp.
    OutLabelOwner interface{}

    // Is from a GR neighbor. The type is bool.
    IsFromGracefulRestartableNeighbor interface{}

    // Is the entry stale? This may happen during a graceful restart. The type is
    // bool.
    IsStale interface{}

    // Nexthop LDP peer.
    NexthopPeerLdpIdent MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetEntityData() *types.CommonEntityData {
    mplsOutgoingInfo.EntityData.YFilter = mplsOutgoingInfo.YFilter
    mplsOutgoingInfo.EntityData.YangName = "mpls-outgoing-info"
    mplsOutgoingInfo.EntityData.BundleName = "cisco_ios_xe"
    mplsOutgoingInfo.EntityData.ParentYangName = "remote-lfa"
    mplsOutgoingInfo.EntityData.SegmentPath = "mpls-outgoing-info"
    mplsOutgoingInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsOutgoingInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsOutgoingInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsOutgoingInfo.EntityData.Children = types.NewOrderedMap()
    mplsOutgoingInfo.EntityData.Children.Append("nexthop-peer-ldp-ident", types.YChild{"NexthopPeerLdpIdent", &mplsOutgoingInfo.NexthopPeerLdpIdent})
    mplsOutgoingInfo.EntityData.Leafs = types.NewOrderedMap()
    mplsOutgoingInfo.EntityData.Leafs.Append("out-label", types.YLeaf{"OutLabel", mplsOutgoingInfo.OutLabel})
    mplsOutgoingInfo.EntityData.Leafs.Append("out-label-type", types.YLeaf{"OutLabelType", mplsOutgoingInfo.OutLabelType})
    mplsOutgoingInfo.EntityData.Leafs.Append("out-label-owner", types.YLeaf{"OutLabelOwner", mplsOutgoingInfo.OutLabelOwner})
    mplsOutgoingInfo.EntityData.Leafs.Append("is-from-graceful-restartable-neighbor", types.YLeaf{"IsFromGracefulRestartableNeighbor", mplsOutgoingInfo.IsFromGracefulRestartableNeighbor})
    mplsOutgoingInfo.EntityData.Leafs.Append("is-stale", types.YLeaf{"IsStale", mplsOutgoingInfo.IsStale})

    mplsOutgoingInfo.EntityData.YListKeys = []string {}

    return &(mplsOutgoingInfo.EntityData)
}

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent
// Nexthop LDP peer
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSR identifier. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Label space identifier. The type is interface{} with range: 0..65535.
    LabelSpaceId interface{}
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetEntityData() *types.CommonEntityData {
    nexthopPeerLdpIdent.EntityData.YFilter = nexthopPeerLdpIdent.YFilter
    nexthopPeerLdpIdent.EntityData.YangName = "nexthop-peer-ldp-ident"
    nexthopPeerLdpIdent.EntityData.BundleName = "cisco_ios_xe"
    nexthopPeerLdpIdent.EntityData.ParentYangName = "mpls-outgoing-info"
    nexthopPeerLdpIdent.EntityData.SegmentPath = "nexthop-peer-ldp-ident"
    nexthopPeerLdpIdent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nexthopPeerLdpIdent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nexthopPeerLdpIdent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nexthopPeerLdpIdent.EntityData.Children = types.NewOrderedMap()
    nexthopPeerLdpIdent.EntityData.Leafs = types.NewOrderedMap()
    nexthopPeerLdpIdent.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", nexthopPeerLdpIdent.LsrId})
    nexthopPeerLdpIdent.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", nexthopPeerLdpIdent.LabelSpaceId})

    nexthopPeerLdpIdent.EntityData.YListKeys = []string {}

    return &(nexthopPeerLdpIdent.EntityData)
}

// MplsLdp_MplsLdpState_Bindings
// The detailed LDP Bindings.
type MplsLdp_MplsLdpState_Bindings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This container holds the bindings specific to this VRF and AF.
    BindingsSumAfs MplsLdp_MplsLdpState_Bindings_BindingsSumAfs

    // This list contains the MPLS LDP Label Bindings for each IP Prefix. Label
    // bindings provide the local MPLS Label, a list of remote labels, any filters
    // affecting advertisment of that filter, and a list of neighbors to which the
    // label has been advertised. The type is slice of
    // MplsLdp_MplsLdpState_Bindings_Binding.
    Binding []*MplsLdp_MplsLdpState_Bindings_Binding
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetEntityData() *types.CommonEntityData {
    bindings.EntityData.YFilter = bindings.YFilter
    bindings.EntityData.YangName = "bindings"
    bindings.EntityData.BundleName = "cisco_ios_xe"
    bindings.EntityData.ParentYangName = "mpls-ldp-state"
    bindings.EntityData.SegmentPath = "bindings"
    bindings.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bindings.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bindings.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bindings.EntityData.Children = types.NewOrderedMap()
    bindings.EntityData.Children.Append("bindings-sum-afs", types.YChild{"BindingsSumAfs", &bindings.BindingsSumAfs})
    bindings.EntityData.Children.Append("binding", types.YChild{"Binding", nil})
    for i := range bindings.Binding {
        bindings.EntityData.Children.Append(types.GetSegmentPath(bindings.Binding[i]), types.YChild{"Binding", bindings.Binding[i]})
    }
    bindings.EntityData.Leafs = types.NewOrderedMap()

    bindings.EntityData.YListKeys = []string {}

    return &(bindings.EntityData)
}

// MplsLdp_MplsLdpState_Bindings_BindingsSumAfs
// This container holds the bindings specific to this VRF
// and AF.
type MplsLdp_MplsLdpState_Bindings_BindingsSumAfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Counters for the LDP Label Information Base for this VRF/AF. The type is
    // slice of MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf.
    BindingSumAf []*MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetEntityData() *types.CommonEntityData {
    bindingsSumAfs.EntityData.YFilter = bindingsSumAfs.YFilter
    bindingsSumAfs.EntityData.YangName = "bindings-sum-afs"
    bindingsSumAfs.EntityData.BundleName = "cisco_ios_xe"
    bindingsSumAfs.EntityData.ParentYangName = "bindings"
    bindingsSumAfs.EntityData.SegmentPath = "bindings-sum-afs"
    bindingsSumAfs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bindingsSumAfs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bindingsSumAfs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bindingsSumAfs.EntityData.Children = types.NewOrderedMap()
    bindingsSumAfs.EntityData.Children.Append("binding-sum-af", types.YChild{"BindingSumAf", nil})
    for i := range bindingsSumAfs.BindingSumAf {
        bindingsSumAfs.EntityData.Children.Append(types.GetSegmentPath(bindingsSumAfs.BindingSumAf[i]), types.YChild{"BindingSumAf", bindingsSumAfs.BindingSumAf[i]})
    }
    bindingsSumAfs.EntityData.Leafs = types.NewOrderedMap()

    bindingsSumAfs.EntityData.YListKeys = []string {}

    return &(bindingsSumAfs.EntityData)
}

// MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf
// Counters for the LDP Label Information Base for this
// VRF/AF.
type MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // Total bindings. The type is interface{} with range: 0..4294967295.
    BindingTotal interface{}

    // Bindings with no route. The type is interface{} with range: 0..4294967295.
    BindingNoRoute interface{}

    // Local bindings with no route. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalNoRoute interface{}

    // Number of local bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocal interface{}

    // Number of local null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalNull interface{}

    // Number of local implicit null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalImplicitNull interface{}

    // Number of local explicit null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalExplicitNull interface{}

    // Number of local non-null bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalNonNull interface{}

    // This is the number of local bindings needing label but which hit the
    // Out-Of-Resource condition. The type is interface{} with range:
    // 0..4294967295.
    BindingLocalOor interface{}

    // Lowest allocated label. The type is interface{} with range: 0..4294967295.
    LowestAllocatedLabel interface{}

    // Highest allocated label. The type is interface{} with range: 0..4294967295.
    HighestAllocatedLabel interface{}

    // Number of remote bindings. The type is interface{} with range:
    // 0..4294967295.
    BindingRemote interface{}
}

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetEntityData() *types.CommonEntityData {
    bindingSumAf.EntityData.YFilter = bindingSumAf.YFilter
    bindingSumAf.EntityData.YangName = "binding-sum-af"
    bindingSumAf.EntityData.BundleName = "cisco_ios_xe"
    bindingSumAf.EntityData.ParentYangName = "bindings-sum-afs"
    bindingSumAf.EntityData.SegmentPath = "binding-sum-af" + types.AddKeyToken(bindingSumAf.VrfName, "vrf-name") + types.AddKeyToken(bindingSumAf.AfName, "af-name")
    bindingSumAf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bindingSumAf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bindingSumAf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bindingSumAf.EntityData.Children = types.NewOrderedMap()
    bindingSumAf.EntityData.Leafs = types.NewOrderedMap()
    bindingSumAf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", bindingSumAf.VrfName})
    bindingSumAf.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", bindingSumAf.AfName})
    bindingSumAf.EntityData.Leafs.Append("binding-total", types.YLeaf{"BindingTotal", bindingSumAf.BindingTotal})
    bindingSumAf.EntityData.Leafs.Append("binding-no-route", types.YLeaf{"BindingNoRoute", bindingSumAf.BindingNoRoute})
    bindingSumAf.EntityData.Leafs.Append("binding-local-no-route", types.YLeaf{"BindingLocalNoRoute", bindingSumAf.BindingLocalNoRoute})
    bindingSumAf.EntityData.Leafs.Append("binding-local", types.YLeaf{"BindingLocal", bindingSumAf.BindingLocal})
    bindingSumAf.EntityData.Leafs.Append("binding-local-null", types.YLeaf{"BindingLocalNull", bindingSumAf.BindingLocalNull})
    bindingSumAf.EntityData.Leafs.Append("binding-local-implicit-null", types.YLeaf{"BindingLocalImplicitNull", bindingSumAf.BindingLocalImplicitNull})
    bindingSumAf.EntityData.Leafs.Append("binding-local-explicit-null", types.YLeaf{"BindingLocalExplicitNull", bindingSumAf.BindingLocalExplicitNull})
    bindingSumAf.EntityData.Leafs.Append("binding-local-non-null", types.YLeaf{"BindingLocalNonNull", bindingSumAf.BindingLocalNonNull})
    bindingSumAf.EntityData.Leafs.Append("binding-local-oor", types.YLeaf{"BindingLocalOor", bindingSumAf.BindingLocalOor})
    bindingSumAf.EntityData.Leafs.Append("lowest-allocated-label", types.YLeaf{"LowestAllocatedLabel", bindingSumAf.LowestAllocatedLabel})
    bindingSumAf.EntityData.Leafs.Append("highest-allocated-label", types.YLeaf{"HighestAllocatedLabel", bindingSumAf.HighestAllocatedLabel})
    bindingSumAf.EntityData.Leafs.Append("binding-remote", types.YLeaf{"BindingRemote", bindingSumAf.BindingRemote})

    bindingSumAf.EntityData.YListKeys = []string {"VrfName", "AfName"}

    return &(bindingSumAf.EntityData)
}

// MplsLdp_MplsLdpState_Bindings_Binding
// This list contains the MPLS LDP Label Bindings for each
// IP Prefix. Label bindings provide the local MPLS Label,
// a list of remote labels, any filters affecting
// advertisment of that filter, and a list of neighbors to
// which the label has been advertised.
type MplsLdp_MplsLdpState_Bindings_Binding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. This leaf contains the IP Prefix being bound. The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    Prefix interface{}

    // This is the MPLS LDP Binding IP Prefix. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    FwdPrefix interface{}

    // This is the MPLS LDP Binding Prefix Length. The type is interface{} with
    // range: 0..255.
    PrefixLength interface{}

    // This is the MPLS LDP Binding Local label. The type is interface{} with
    // range: 0..4294967295.
    LocalLabel interface{}

    // This is the MPLS LDP Binding Local Binding revision. The type is
    // interface{} with range: 0..4294967295.
    LeLocalBindingRevision interface{}

    // This is the MPLS LDP Binding Local label state. The type is
    // LocalLabelState.
    LeLocalLabelState interface{}

    // This is true if the MPLS LDP Binding has no route. The type is bool.
    IsNoRoute interface{}

    // This is true if the MPLS LDP Binding Label space is depleted, Out Of
    // Resource. No new labels can be allocated. The type is bool.
    LabelOor interface{}

    // This contains the filter name for this binding's prefix. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string.
    AdvertisePrefixFilter interface{}

    // This contains the filter name for this binding's Advertise LSR. The filter
    // type is device specific and could be an ACL, a prefix list, or other
    // mechanism. The type is string.
    AdvertiseLsrFilter interface{}

    // Config/User enforced local label value. The type is bool.
    ConfigEnforcedLocalLabelValue interface{}

    // MPLS LDP Remote Binding Information. The type is slice of
    // MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding.
    RemoteBinding []*MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding

    // Peers to which this entry is advertised. The type is slice of
    // MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo.
    PeersAdvertisedTo []*MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetEntityData() *types.CommonEntityData {
    binding.EntityData.YFilter = binding.YFilter
    binding.EntityData.YangName = "binding"
    binding.EntityData.BundleName = "cisco_ios_xe"
    binding.EntityData.ParentYangName = "bindings"
    binding.EntityData.SegmentPath = "binding" + types.AddKeyToken(binding.VrfName, "vrf-name") + types.AddKeyToken(binding.Prefix, "prefix")
    binding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    binding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    binding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    binding.EntityData.Children = types.NewOrderedMap()
    binding.EntityData.Children.Append("remote-binding", types.YChild{"RemoteBinding", nil})
    for i := range binding.RemoteBinding {
        binding.EntityData.Children.Append(types.GetSegmentPath(binding.RemoteBinding[i]), types.YChild{"RemoteBinding", binding.RemoteBinding[i]})
    }
    binding.EntityData.Children.Append("peers-advertised-to", types.YChild{"PeersAdvertisedTo", nil})
    for i := range binding.PeersAdvertisedTo {
        binding.EntityData.Children.Append(types.GetSegmentPath(binding.PeersAdvertisedTo[i]), types.YChild{"PeersAdvertisedTo", binding.PeersAdvertisedTo[i]})
    }
    binding.EntityData.Leafs = types.NewOrderedMap()
    binding.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", binding.VrfName})
    binding.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", binding.Prefix})
    binding.EntityData.Leafs.Append("fwd-prefix", types.YLeaf{"FwdPrefix", binding.FwdPrefix})
    binding.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", binding.PrefixLength})
    binding.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", binding.LocalLabel})
    binding.EntityData.Leafs.Append("le-local-binding-revision", types.YLeaf{"LeLocalBindingRevision", binding.LeLocalBindingRevision})
    binding.EntityData.Leafs.Append("le-local-label-state", types.YLeaf{"LeLocalLabelState", binding.LeLocalLabelState})
    binding.EntityData.Leafs.Append("is-no-route", types.YLeaf{"IsNoRoute", binding.IsNoRoute})
    binding.EntityData.Leafs.Append("label-oor", types.YLeaf{"LabelOor", binding.LabelOor})
    binding.EntityData.Leafs.Append("advertise-prefix-filter", types.YLeaf{"AdvertisePrefixFilter", binding.AdvertisePrefixFilter})
    binding.EntityData.Leafs.Append("advertise-lsr-filter", types.YLeaf{"AdvertiseLsrFilter", binding.AdvertiseLsrFilter})
    binding.EntityData.Leafs.Append("config-enforced-local-label-value", types.YLeaf{"ConfigEnforcedLocalLabelValue", binding.ConfigEnforcedLocalLabelValue})

    binding.EntityData.YListKeys = []string {"VrfName", "Prefix"}

    return &(binding.EntityData)
}

// MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding
// MPLS LDP Remote Binding Information
type MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is the remote Label. The type is interface{} with range:
    // 0..4294967295.
    RemoteLabel interface{}

    // Is the entry stale. The type is bool.
    IsStale interface{}

    // Assigning peer.
    AssigningPeerLdpIdent MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetEntityData() *types.CommonEntityData {
    remoteBinding.EntityData.YFilter = remoteBinding.YFilter
    remoteBinding.EntityData.YangName = "remote-binding"
    remoteBinding.EntityData.BundleName = "cisco_ios_xe"
    remoteBinding.EntityData.ParentYangName = "binding"
    remoteBinding.EntityData.SegmentPath = "remote-binding"
    remoteBinding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    remoteBinding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    remoteBinding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    remoteBinding.EntityData.Children = types.NewOrderedMap()
    remoteBinding.EntityData.Children.Append("assigning-peer-ldp-ident", types.YChild{"AssigningPeerLdpIdent", &remoteBinding.AssigningPeerLdpIdent})
    remoteBinding.EntityData.Leafs = types.NewOrderedMap()
    remoteBinding.EntityData.Leafs.Append("remote-label", types.YLeaf{"RemoteLabel", remoteBinding.RemoteLabel})
    remoteBinding.EntityData.Leafs.Append("is-stale", types.YLeaf{"IsStale", remoteBinding.IsStale})

    remoteBinding.EntityData.YListKeys = []string {}

    return &(remoteBinding.EntityData)
}

// MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent
// Assigning peer
type MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSR identifier. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Label space identifier. The type is interface{} with range: 0..65535.
    LabelSpaceId interface{}
}

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetEntityData() *types.CommonEntityData {
    assigningPeerLdpIdent.EntityData.YFilter = assigningPeerLdpIdent.YFilter
    assigningPeerLdpIdent.EntityData.YangName = "assigning-peer-ldp-ident"
    assigningPeerLdpIdent.EntityData.BundleName = "cisco_ios_xe"
    assigningPeerLdpIdent.EntityData.ParentYangName = "remote-binding"
    assigningPeerLdpIdent.EntityData.SegmentPath = "assigning-peer-ldp-ident"
    assigningPeerLdpIdent.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    assigningPeerLdpIdent.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    assigningPeerLdpIdent.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    assigningPeerLdpIdent.EntityData.Children = types.NewOrderedMap()
    assigningPeerLdpIdent.EntityData.Leafs = types.NewOrderedMap()
    assigningPeerLdpIdent.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", assigningPeerLdpIdent.LsrId})
    assigningPeerLdpIdent.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", assigningPeerLdpIdent.LabelSpaceId})

    assigningPeerLdpIdent.EntityData.YListKeys = []string {}

    return &(assigningPeerLdpIdent.EntityData)
}

// MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo
// Peers to which this entry is advertised.
type MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSR identifier. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Label space identifier. The type is interface{} with range: 0..65535.
    LabelSpaceId interface{}
}

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetEntityData() *types.CommonEntityData {
    peersAdvertisedTo.EntityData.YFilter = peersAdvertisedTo.YFilter
    peersAdvertisedTo.EntityData.YangName = "peers-advertised-to"
    peersAdvertisedTo.EntityData.BundleName = "cisco_ios_xe"
    peersAdvertisedTo.EntityData.ParentYangName = "binding"
    peersAdvertisedTo.EntityData.SegmentPath = "peers-advertised-to"
    peersAdvertisedTo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    peersAdvertisedTo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    peersAdvertisedTo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    peersAdvertisedTo.EntityData.Children = types.NewOrderedMap()
    peersAdvertisedTo.EntityData.Leafs = types.NewOrderedMap()
    peersAdvertisedTo.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", peersAdvertisedTo.LsrId})
    peersAdvertisedTo.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", peersAdvertisedTo.LabelSpaceId})

    peersAdvertisedTo.EntityData.YListKeys = []string {}

    return &(peersAdvertisedTo.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors
// The LDP Neighbors Information
type MplsLdp_MplsLdpState_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on a particular LDP neighbor. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_Neighbor.
    Neighbor []*MplsLdp_MplsLdpState_Neighbors_Neighbor

    // For this Neighbor, this is the list of adjacencies between the neighbor and
    // the local node. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_NbrAdjs.
    NbrAdjs []*MplsLdp_MplsLdpState_Neighbors_NbrAdjs

    // MPLS LDP Statistics Information.
    StatsInfo MplsLdp_MplsLdpState_Neighbors_StatsInfo

    // LDP Backoff Information.
    Backoffs MplsLdp_MplsLdpState_Neighbors_Backoffs

    // This is the LDP NSR state for this neighbor.
    NsrNbrDetail MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xe"
    neighbors.EntityData.ParentYangName = "mpls-ldp-state"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Children.Append("nbr-adjs", types.YChild{"NbrAdjs", nil})
    for i := range neighbors.NbrAdjs {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.NbrAdjs[i]), types.YChild{"NbrAdjs", neighbors.NbrAdjs[i]})
    }
    neighbors.EntityData.Children.Append("stats-info", types.YChild{"StatsInfo", &neighbors.StatsInfo})
    neighbors.EntityData.Children.Append("backoffs", types.YChild{"Backoffs", &neighbors.Backoffs})
    neighbors.EntityData.Children.Append("nsr-nbr-detail", types.YChild{"NsrNbrDetail", &neighbors.NsrNbrDetail})
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor
// Information on a particular LDP neighbor
type MplsLdp_MplsLdpState_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. LSR ID of neighbor. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LsrId interface{}

    // Label space ID of neighbor. The type is interface{} with range: 0..65535.
    LabelSpaceId interface{}

    // During session establishment the LSR/LER takes either the active role or
    // the passive role based on address comparisons.  This object indicates
    // whether this LSR/LER was behaving in an active role or passive role during
    // this session's establishment.  The value of unknown(1), indicates that the
    // role is not able to be determined at the present time. The type is
    // SessionRole.
    SessionRole interface{}

    // The version of the LDP Protocol which this session is using.  This is the
    // version of the LDP protocol which has been negotiated during session
    // initialization. The type is interface{} with range: 1..65535.
    SessionProtVer interface{}

    // Up time in seconds. The type is interface{} with range: 0..4294967295.
    // Units are seconds.
    UpTimeSeconds interface{}

    // If the value of this object is 0 (zero) then Loop Dection for Path Vectors
    // for this neighor is disabled.  Otherwise, if this object has a value
    // greater than zero, then Loop Dection for Path  Vectors for this neighbor is
    // enabled and the Path Vector Limit is this value. The type is interface{}
    // with range: 0..255.
    NbrPathVectorLimit interface{}

    // Is Label advertisement mode in Downstream On Demand mode or not?. The type
    // is bool.
    DownstreamOnDemand interface{}

    // Session holdtime value in seconds from the peer. The type is interface{}
    // with range: 0..4294967295. Units are seconds.
    PeerHoldTime interface{}

    // Session keepalive interval in seconds from the peer. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    PeerKeepAliveInterval interface{}

    // LDP adjacency peer state. The type is AdjState.
    PeerState interface{}

    // This contains the IPv4 Inbound accept filter name. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string with length: 0..80.
    InboundIpv4 interface{}

    // This contains the IPv6 Inbound accept filter name. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string with length: 0..80.
    InboundIpv6Filter interface{}

    // This contains the IPv4 Outbound advertise filter name. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string with length: 0..80.
    OutboundIpv4Filter interface{}

    // This contains the IPv6 Outbound advertise filter name. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string with length: 0..80.
    OutboundIpv6Filter interface{}

    // Session Protection enabled. The type is bool.
    HasSp interface{}

    // Session Protection state. The type is string with length: 0..80.
    SpState interface{}

    // This contains the Session Protection filter name. The filter type is device
    // specific and could be an ACL, a prefix list, or other mechanism. The type
    // is string with length: 0..80.
    SpFilter interface{}

    // Session Protection has non-default duration. The type is bool.
    SpHasDuration interface{}

    // Session protection holdup time duration in seconds. The type is interface{}
    // with range: 0..4294967295. Units are seconds.
    SpDuration interface{}

    // Session Protection holdup timer is running. The type is bool.
    SphtRunning interface{}

    // Session Protection holdup time remaining value in seconds. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    SphtRemaining interface{}

    // BGP labeled prefixes advertisement state. The type is NbrBgpAdvtState.
    BgpAdvertisementState interface{}

    // True if BGP labeled prefixes are advertised to the neighbor. The type is
    // bool.
    AdvertiseBgpPrefixes interface{}

    // Targeted Session clients. The type is slice of string.
    Client []interface{}

    // Duplicate IPv4/IPv6 address bound to this peer. The type is one of the
    // following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DuplicateAddress []interface{}

    // This is the MPLS LDP Neighbor Bound IPv4/IPv6 Address. The type is one of
    // the following types: slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NbrBoundAddress []interface{}

    // Neighbor Statistics.
    NbrStats MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats

    // This container holds the graceful restart information for this adjacency.
    GracefulRestartAdjacency MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency

    // MPLS LDP Neighbor TCP Information.
    TcpInformation MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation

    // Capabilities sent to and received from neighbor.
    Capabilities MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities
}

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xe"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.VrfName, "vrf-name") + types.AddKeyToken(neighbor.LsrId, "lsr-id")
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("nbr-stats", types.YChild{"NbrStats", &neighbor.NbrStats})
    neighbor.EntityData.Children.Append("graceful-restart-adjacency", types.YChild{"GracefulRestartAdjacency", &neighbor.GracefulRestartAdjacency})
    neighbor.EntityData.Children.Append("tcp-information", types.YChild{"TcpInformation", &neighbor.TcpInformation})
    neighbor.EntityData.Children.Append("capabilities", types.YChild{"Capabilities", &neighbor.Capabilities})
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", neighbor.VrfName})
    neighbor.EntityData.Leafs.Append("lsr-id", types.YLeaf{"LsrId", neighbor.LsrId})
    neighbor.EntityData.Leafs.Append("label-space-id", types.YLeaf{"LabelSpaceId", neighbor.LabelSpaceId})
    neighbor.EntityData.Leafs.Append("session-role", types.YLeaf{"SessionRole", neighbor.SessionRole})
    neighbor.EntityData.Leafs.Append("session-prot-ver", types.YLeaf{"SessionProtVer", neighbor.SessionProtVer})
    neighbor.EntityData.Leafs.Append("up-time-seconds", types.YLeaf{"UpTimeSeconds", neighbor.UpTimeSeconds})
    neighbor.EntityData.Leafs.Append("nbr-path-vector-limit", types.YLeaf{"NbrPathVectorLimit", neighbor.NbrPathVectorLimit})
    neighbor.EntityData.Leafs.Append("downstream-on-demand", types.YLeaf{"DownstreamOnDemand", neighbor.DownstreamOnDemand})
    neighbor.EntityData.Leafs.Append("peer-hold-time", types.YLeaf{"PeerHoldTime", neighbor.PeerHoldTime})
    neighbor.EntityData.Leafs.Append("peer-keep-alive-interval", types.YLeaf{"PeerKeepAliveInterval", neighbor.PeerKeepAliveInterval})
    neighbor.EntityData.Leafs.Append("peer-state", types.YLeaf{"PeerState", neighbor.PeerState})
    neighbor.EntityData.Leafs.Append("inbound-ipv4", types.YLeaf{"InboundIpv4", neighbor.InboundIpv4})
    neighbor.EntityData.Leafs.Append("inbound-ipv6-filter", types.YLeaf{"InboundIpv6Filter", neighbor.InboundIpv6Filter})
    neighbor.EntityData.Leafs.Append("outbound-ipv4-filter", types.YLeaf{"OutboundIpv4Filter", neighbor.OutboundIpv4Filter})
    neighbor.EntityData.Leafs.Append("outbound-ipv6-filter", types.YLeaf{"OutboundIpv6Filter", neighbor.OutboundIpv6Filter})
    neighbor.EntityData.Leafs.Append("has-sp", types.YLeaf{"HasSp", neighbor.HasSp})
    neighbor.EntityData.Leafs.Append("sp-state", types.YLeaf{"SpState", neighbor.SpState})
    neighbor.EntityData.Leafs.Append("sp-filter", types.YLeaf{"SpFilter", neighbor.SpFilter})
    neighbor.EntityData.Leafs.Append("sp-has-duration", types.YLeaf{"SpHasDuration", neighbor.SpHasDuration})
    neighbor.EntityData.Leafs.Append("sp-duration", types.YLeaf{"SpDuration", neighbor.SpDuration})
    neighbor.EntityData.Leafs.Append("spht-running", types.YLeaf{"SphtRunning", neighbor.SphtRunning})
    neighbor.EntityData.Leafs.Append("spht-remaining", types.YLeaf{"SphtRemaining", neighbor.SphtRemaining})
    neighbor.EntityData.Leafs.Append("bgp-advertisement-state", types.YLeaf{"BgpAdvertisementState", neighbor.BgpAdvertisementState})
    neighbor.EntityData.Leafs.Append("advertise-bgp-prefixes", types.YLeaf{"AdvertiseBgpPrefixes", neighbor.AdvertiseBgpPrefixes})
    neighbor.EntityData.Leafs.Append("client", types.YLeaf{"Client", neighbor.Client})
    neighbor.EntityData.Leafs.Append("duplicate-address", types.YLeaf{"DuplicateAddress", neighbor.DuplicateAddress})
    neighbor.EntityData.Leafs.Append("nbr-bound-address", types.YLeaf{"NbrBoundAddress", neighbor.NbrBoundAddress})

    neighbor.EntityData.YListKeys = []string {"VrfName", "LsrId"}

    return &(neighbor.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats
// Neighbor Statistics.
type MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of MPLS LDP messages sent to this neighbor. The type is interface{}
    // with range: 0..4294967295.
    TaPiesSent interface{}

    // Number of MPLS LDP messages received from this neighbor. The type is
    // interface{} with range: 0..4294967295.
    TaPiesRcvd interface{}

    // Number of neighbor IPv4 discovery sources. The type is interface{} with
    // range: 0..4294967295.
    NumOfNbrIpv4Discovery interface{}

    // Number of neighbor IPv6 discovery sources. The type is interface{} with
    // range: 0..4294967295.
    NumOfNbrIpv6Discovery interface{}

    // Number of IPv4 addresses for which the neighbor is advertising labels. The
    // type is interface{} with range: 0..4294967295.
    NumOfNbrIpv4Addresses interface{}

    // Number of IPv6 addresses for which the neighbor is advertising labels. The
    // type is interface{} with range: 0..4294967295.
    NumOfNbrIpv6Addresses interface{}

    // Number of IPv4 labels the neighbor is advertising. The type is interface{}
    // with range: 0..4294967295.
    NumOfNbrIpv4Lbl interface{}

    // Number of IPv6 labels the neighbor is advertising. The type is interface{}
    // with range: 0..4294967295.
    NumOfNbrIpv6Lbl interface{}
}

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetEntityData() *types.CommonEntityData {
    nbrStats.EntityData.YFilter = nbrStats.YFilter
    nbrStats.EntityData.YangName = "nbr-stats"
    nbrStats.EntityData.BundleName = "cisco_ios_xe"
    nbrStats.EntityData.ParentYangName = "neighbor"
    nbrStats.EntityData.SegmentPath = "nbr-stats"
    nbrStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nbrStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nbrStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nbrStats.EntityData.Children = types.NewOrderedMap()
    nbrStats.EntityData.Leafs = types.NewOrderedMap()
    nbrStats.EntityData.Leafs.Append("ta-pies-sent", types.YLeaf{"TaPiesSent", nbrStats.TaPiesSent})
    nbrStats.EntityData.Leafs.Append("ta-pies-rcvd", types.YLeaf{"TaPiesRcvd", nbrStats.TaPiesRcvd})
    nbrStats.EntityData.Leafs.Append("num-of-nbr-ipv4-discovery", types.YLeaf{"NumOfNbrIpv4Discovery", nbrStats.NumOfNbrIpv4Discovery})
    nbrStats.EntityData.Leafs.Append("num-of-nbr-ipv6-discovery", types.YLeaf{"NumOfNbrIpv6Discovery", nbrStats.NumOfNbrIpv6Discovery})
    nbrStats.EntityData.Leafs.Append("num-of-nbr-ipv4-addresses", types.YLeaf{"NumOfNbrIpv4Addresses", nbrStats.NumOfNbrIpv4Addresses})
    nbrStats.EntityData.Leafs.Append("num-of-nbr-ipv6-addresses", types.YLeaf{"NumOfNbrIpv6Addresses", nbrStats.NumOfNbrIpv6Addresses})
    nbrStats.EntityData.Leafs.Append("num-of-nbr-ipv4-lbl", types.YLeaf{"NumOfNbrIpv4Lbl", nbrStats.NumOfNbrIpv4Lbl})
    nbrStats.EntityData.Leafs.Append("num-of-nbr-ipv6-lbl", types.YLeaf{"NumOfNbrIpv6Lbl", nbrStats.NumOfNbrIpv6Lbl})

    nbrStats.EntityData.YListKeys = []string {}

    return &(nbrStats.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency
// This container holds the graceful restart information
// for this adjacency.
type MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is this neighbor graceful restartable?. The type is bool.
    IsGracefulRestartable interface{}

    // This leaf is the reconnect timeout in microseconds. The type is interface{}
    // with range: 0..4294967295. Units are microseconds.
    ReconnectTimeout interface{}

    // This leaf is the recovery time in microseconds. The type is interface{}
    // with range: 0..4294967295. Units are microseconds.
    RecoveryTime interface{}

    // This is set if the liveness timer is running. The type is interface{}.
    IsLivenessTimerRunning interface{}

    // Remaining time from liveness timer in seconds. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    LivenessTimerRemainingSeconds interface{}

    // This is set if the recovery timer is running. The type is interface{}.
    IsRecoveryTimerRunning interface{}

    // Recovery timer remaining time in seconds. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    RecoveryTimerRemainingSeconds interface{}

    // This is the current count of back-to-back flaps. The type is interface{}
    // with range: 0..4294967295.
    DownNbrFlapCount interface{}

    // This identity provides the reason that the LDP Session with this neighbor
    // is down. The reason does not persist if the session was down but is now
    // recovered. The type is one of the following:
    // DownNbrReasonDiscHelloDownNbrReasonNbrHoldDownNbrReasonNa.
    DownNbrDownReason interface{}
}

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetEntityData() *types.CommonEntityData {
    gracefulRestartAdjacency.EntityData.YFilter = gracefulRestartAdjacency.YFilter
    gracefulRestartAdjacency.EntityData.YangName = "graceful-restart-adjacency"
    gracefulRestartAdjacency.EntityData.BundleName = "cisco_ios_xe"
    gracefulRestartAdjacency.EntityData.ParentYangName = "neighbor"
    gracefulRestartAdjacency.EntityData.SegmentPath = "graceful-restart-adjacency"
    gracefulRestartAdjacency.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    gracefulRestartAdjacency.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    gracefulRestartAdjacency.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    gracefulRestartAdjacency.EntityData.Children = types.NewOrderedMap()
    gracefulRestartAdjacency.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestartAdjacency.EntityData.Leafs.Append("is-graceful-restartable", types.YLeaf{"IsGracefulRestartable", gracefulRestartAdjacency.IsGracefulRestartable})
    gracefulRestartAdjacency.EntityData.Leafs.Append("reconnect-timeout", types.YLeaf{"ReconnectTimeout", gracefulRestartAdjacency.ReconnectTimeout})
    gracefulRestartAdjacency.EntityData.Leafs.Append("recovery-time", types.YLeaf{"RecoveryTime", gracefulRestartAdjacency.RecoveryTime})
    gracefulRestartAdjacency.EntityData.Leafs.Append("is-liveness-timer-running", types.YLeaf{"IsLivenessTimerRunning", gracefulRestartAdjacency.IsLivenessTimerRunning})
    gracefulRestartAdjacency.EntityData.Leafs.Append("liveness-timer-remaining-seconds", types.YLeaf{"LivenessTimerRemainingSeconds", gracefulRestartAdjacency.LivenessTimerRemainingSeconds})
    gracefulRestartAdjacency.EntityData.Leafs.Append("is-recovery-timer-running", types.YLeaf{"IsRecoveryTimerRunning", gracefulRestartAdjacency.IsRecoveryTimerRunning})
    gracefulRestartAdjacency.EntityData.Leafs.Append("recovery-timer-remaining-seconds", types.YLeaf{"RecoveryTimerRemainingSeconds", gracefulRestartAdjacency.RecoveryTimerRemainingSeconds})
    gracefulRestartAdjacency.EntityData.Leafs.Append("down-nbr-flap-count", types.YLeaf{"DownNbrFlapCount", gracefulRestartAdjacency.DownNbrFlapCount})
    gracefulRestartAdjacency.EntityData.Leafs.Append("down-nbr-down-reason", types.YLeaf{"DownNbrDownReason", gracefulRestartAdjacency.DownNbrDownReason})

    gracefulRestartAdjacency.EntityData.YListKeys = []string {}

    return &(gracefulRestartAdjacency.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation
// MPLS LDP Neighbor TCP Information
type MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is the foreign host address used by TCP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    ForeignHost interface{}

    // This is the local host address used by TCP. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalHost interface{}

    // Foreign port number. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Local port number. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Is MD5 Digest on. The type is bool.
    IsMd5On interface{}

    // up time. The type is string.
    UpTime interface{}
}

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetEntityData() *types.CommonEntityData {
    tcpInformation.EntityData.YFilter = tcpInformation.YFilter
    tcpInformation.EntityData.YangName = "tcp-information"
    tcpInformation.EntityData.BundleName = "cisco_ios_xe"
    tcpInformation.EntityData.ParentYangName = "neighbor"
    tcpInformation.EntityData.SegmentPath = "tcp-information"
    tcpInformation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    tcpInformation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    tcpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    tcpInformation.EntityData.Children = types.NewOrderedMap()
    tcpInformation.EntityData.Leafs = types.NewOrderedMap()
    tcpInformation.EntityData.Leafs.Append("foreign-host", types.YLeaf{"ForeignHost", tcpInformation.ForeignHost})
    tcpInformation.EntityData.Leafs.Append("local-host", types.YLeaf{"LocalHost", tcpInformation.LocalHost})
    tcpInformation.EntityData.Leafs.Append("foreign-port", types.YLeaf{"ForeignPort", tcpInformation.ForeignPort})
    tcpInformation.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", tcpInformation.LocalPort})
    tcpInformation.EntityData.Leafs.Append("is-md5-on", types.YLeaf{"IsMd5On", tcpInformation.IsMd5On})
    tcpInformation.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", tcpInformation.UpTime})

    tcpInformation.EntityData.YListKeys = []string {}

    return &(tcpInformation.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities
// Capabilities sent to and received from neighbor
type MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of sent capabilities. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps.
    SentCaps []*MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps

    // List of received capabilities. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps.
    ReceivedCaps []*MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetEntityData() *types.CommonEntityData {
    capabilities.EntityData.YFilter = capabilities.YFilter
    capabilities.EntityData.YangName = "capabilities"
    capabilities.EntityData.BundleName = "cisco_ios_xe"
    capabilities.EntityData.ParentYangName = "neighbor"
    capabilities.EntityData.SegmentPath = "capabilities"
    capabilities.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    capabilities.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    capabilities.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    capabilities.EntityData.Children = types.NewOrderedMap()
    capabilities.EntityData.Children.Append("sent-caps", types.YChild{"SentCaps", nil})
    for i := range capabilities.SentCaps {
        capabilities.EntityData.Children.Append(types.GetSegmentPath(capabilities.SentCaps[i]), types.YChild{"SentCaps", capabilities.SentCaps[i]})
    }
    capabilities.EntityData.Children.Append("received-caps", types.YChild{"ReceivedCaps", nil})
    for i := range capabilities.ReceivedCaps {
        capabilities.EntityData.Children.Append(types.GetSegmentPath(capabilities.ReceivedCaps[i]), types.YChild{"ReceivedCaps", capabilities.ReceivedCaps[i]})
    }
    capabilities.EntityData.Leafs = types.NewOrderedMap()

    capabilities.EntityData.YListKeys = []string {}

    return &(capabilities.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps
// List of sent capabilities
type MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Capability type (IANA assigned). The type is
    // interface{} with range: 0..65535.
    CapType interface{}

    // Capability description. The type is string with length: 0..80.
    CapDes interface{}

    // Capability data length. The type is interface{} with range: 0..65535.
    CapabilityDataLength interface{}

    // Capability data. The type is string.
    CapabilityData interface{}
}

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetEntityData() *types.CommonEntityData {
    sentCaps.EntityData.YFilter = sentCaps.YFilter
    sentCaps.EntityData.YangName = "sent-caps"
    sentCaps.EntityData.BundleName = "cisco_ios_xe"
    sentCaps.EntityData.ParentYangName = "capabilities"
    sentCaps.EntityData.SegmentPath = "sent-caps" + types.AddKeyToken(sentCaps.CapType, "cap-type")
    sentCaps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    sentCaps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    sentCaps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    sentCaps.EntityData.Children = types.NewOrderedMap()
    sentCaps.EntityData.Leafs = types.NewOrderedMap()
    sentCaps.EntityData.Leafs.Append("cap-type", types.YLeaf{"CapType", sentCaps.CapType})
    sentCaps.EntityData.Leafs.Append("cap-des", types.YLeaf{"CapDes", sentCaps.CapDes})
    sentCaps.EntityData.Leafs.Append("capability-data-length", types.YLeaf{"CapabilityDataLength", sentCaps.CapabilityDataLength})
    sentCaps.EntityData.Leafs.Append("capability-data", types.YLeaf{"CapabilityData", sentCaps.CapabilityData})

    sentCaps.EntityData.YListKeys = []string {"CapType"}

    return &(sentCaps.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps
// List of received capabilities
type MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Capability type (IANA assigned). The type is
    // interface{} with range: 0..65535.
    CapType interface{}

    // Capability description. The type is string with length: 0..80.
    CapDes interface{}

    // Capability data length. The type is interface{} with range: 0..65535.
    CapabilityDataLength interface{}

    // Capability data. The type is string.
    CapabilityData interface{}
}

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetEntityData() *types.CommonEntityData {
    receivedCaps.EntityData.YFilter = receivedCaps.YFilter
    receivedCaps.EntityData.YangName = "received-caps"
    receivedCaps.EntityData.BundleName = "cisco_ios_xe"
    receivedCaps.EntityData.ParentYangName = "capabilities"
    receivedCaps.EntityData.SegmentPath = "received-caps" + types.AddKeyToken(receivedCaps.CapType, "cap-type")
    receivedCaps.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    receivedCaps.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    receivedCaps.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    receivedCaps.EntityData.Children = types.NewOrderedMap()
    receivedCaps.EntityData.Leafs = types.NewOrderedMap()
    receivedCaps.EntityData.Leafs.Append("cap-type", types.YLeaf{"CapType", receivedCaps.CapType})
    receivedCaps.EntityData.Leafs.Append("cap-des", types.YLeaf{"CapDes", receivedCaps.CapDes})
    receivedCaps.EntityData.Leafs.Append("capability-data-length", types.YLeaf{"CapabilityDataLength", receivedCaps.CapabilityDataLength})
    receivedCaps.EntityData.Leafs.Append("capability-data", types.YLeaf{"CapabilityData", receivedCaps.CapabilityData})

    receivedCaps.EntityData.YListKeys = []string {"CapType"}

    return &(receivedCaps.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole represents able to be determined at the present time.
type MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole string

const (
    // The role of this LSR in the session is unknown.
    MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole_unknown MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole = "unknown"

    // The role of this LSR in the session is active.
    MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole_active MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole = "active"

    // The role of this LSR in the session is passive.
    MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole_passive MplsLdp_MplsLdpState_Neighbors_Neighbor_SessionRole = "passive"
)

// MplsLdp_MplsLdpState_Neighbors_NbrAdjs
// For this Neighbor, this is the list of adjacencies
// between the neighbor and the local node.
type MplsLdp_MplsLdpState_Neighbors_NbrAdjs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is the interface used by MPLS LDP Link Hello. The type is string.
    // Refers to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // This is the local address used to send the Targeted Hello. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // This is the destination address used to send the Targeted Hello. The type
    // is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    TargetAddress interface{}

    // This is the state of this Targeted Hello instance. The type is DhcState.
    TargetState interface{}
}

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetEntityData() *types.CommonEntityData {
    nbrAdjs.EntityData.YFilter = nbrAdjs.YFilter
    nbrAdjs.EntityData.YangName = "nbr-adjs"
    nbrAdjs.EntityData.BundleName = "cisco_ios_xe"
    nbrAdjs.EntityData.ParentYangName = "neighbors"
    nbrAdjs.EntityData.SegmentPath = "nbr-adjs"
    nbrAdjs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nbrAdjs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nbrAdjs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nbrAdjs.EntityData.Children = types.NewOrderedMap()
    nbrAdjs.EntityData.Leafs = types.NewOrderedMap()
    nbrAdjs.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", nbrAdjs.Interface})
    nbrAdjs.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", nbrAdjs.LocalAddress})
    nbrAdjs.EntityData.Leafs.Append("target-address", types.YLeaf{"TargetAddress", nbrAdjs.TargetAddress})
    nbrAdjs.EntityData.Leafs.Append("target-state", types.YLeaf{"TargetState", nbrAdjs.TargetState})

    nbrAdjs.EntityData.YListKeys = []string {}

    return &(nbrAdjs.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_StatsInfo
// MPLS LDP Statistics Information
type MplsLdp_MplsLdpState_Neighbors_StatsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this entity's counters suffered a discontinuity.  The relevant counters
    // are the specific instances associated with this entity of any counter32
    // object contained in the 'EntityStatsTable'.  If no such discontinuities
    // have occurred since the last re-initialization of the local management
    // subsystem, then this object contains a zero value. The type is interface{}
    // with range: 0..4294967295.
    DisconTime interface{}

    // A count of the Session Initialization messages which were sent or received
    // by this LDP Entity and were NAK'd.   In other words, this counter counts
    // the number of session initializations that failed.  Discontinuities in the
    // value of this counter can occur at re-initialization of the management
    // system, and at other times as indicated by the value of discon-time. The
    // type is interface{} with range: 0..4294967295.
    SessionAttempts interface{}

    // A count of the Session Rejected/No Hello Error Notification Messages sent
    // or received by this LDP Entity.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of discon-time. The type is
    // interface{} with range: 0..4294967295.
    SessRejectNoHello interface{}

    // A count of the Session Rejected/Parameters Advertisement Mode Error
    // Notification Messages sent or received by this LDP Entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    SessRejAd interface{}

    // A count of the Session Rejected/Parameters  Max Pdu Length Error
    // Notification Messages sent or received by this LDP Entity.  Discontinuities
    // in the value of this counter can occur at re-initialization of the
    // management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    SessRejMaxPdu interface{}

    // A count of the Session Rejected/Parameters Label Range Notification
    // Messages sent or received by this LDP Entity.  Discontinuities in the value
    // of this counter can occur at re-initialization of the management system,
    // and at other times as indicated by the value of discon-time. The type is
    // interface{} with range: 0..4294967295.
    SessRejLr interface{}

    // This object counts the number of Bad LDP Identifier Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    BadLdpid interface{}

    // This object counts the number of Bad PDU Length Fatal Errors detected by
    // the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    BadPduLen interface{}

    // This object counts the number of Bad Message Length Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    BadMsgLen interface{}

    // This object counts the number of Bad TLV Length Fatal Errors detected by
    // the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    BadTlvLen interface{}

    // This object counts the number of Malformed TLV Value Fatal Errors detected
    // by the session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    MalformedTlvVal interface{}

    // This object counts the number of Session Keep Alive Timer Expired Errors
    // detected by the session(s) (past and present) associated with this LDP
    // Entity.  Discontinuities in the value of this counter can occur at
    // re-initialization of the management system, and at other times as indicated
    // by the value of discon-time. The type is interface{} with range:
    // 0..4294967295.
    KeepAliveExp interface{}

    // This object counts the number of Shutdown Notifications received related to
    // session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of
    // discon-time. The type is interface{} with range: 0..4294967295.
    ShutdownNotifRec interface{}

    // This object counts the number of Shutdown Notfications sent related to
    // session(s) (past and present) associated with this LDP Entity. 
    // Discontinuities in the value of this counter can occur at re-initialization
    // of the management system, and at other times as indicated by the value of 
    // discon-time. The type is interface{} with range: 0..4294967295.
    ShutdowNotifSent interface{}

    // MPLS LDP message sent counters to this neighbor.
    MessageOut MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut

    // MPLS LDP message received counters from this neighbor.
    MessageIn MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn
}

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetEntityData() *types.CommonEntityData {
    statsInfo.EntityData.YFilter = statsInfo.YFilter
    statsInfo.EntityData.YangName = "stats-info"
    statsInfo.EntityData.BundleName = "cisco_ios_xe"
    statsInfo.EntityData.ParentYangName = "neighbors"
    statsInfo.EntityData.SegmentPath = "stats-info"
    statsInfo.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    statsInfo.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    statsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    statsInfo.EntityData.Children = types.NewOrderedMap()
    statsInfo.EntityData.Children.Append("message-out", types.YChild{"MessageOut", &statsInfo.MessageOut})
    statsInfo.EntityData.Children.Append("message-in", types.YChild{"MessageIn", &statsInfo.MessageIn})
    statsInfo.EntityData.Leafs = types.NewOrderedMap()
    statsInfo.EntityData.Leafs.Append("discon-time", types.YLeaf{"DisconTime", statsInfo.DisconTime})
    statsInfo.EntityData.Leafs.Append("session-attempts", types.YLeaf{"SessionAttempts", statsInfo.SessionAttempts})
    statsInfo.EntityData.Leafs.Append("sess-reject-no-hello", types.YLeaf{"SessRejectNoHello", statsInfo.SessRejectNoHello})
    statsInfo.EntityData.Leafs.Append("sess-rej-ad", types.YLeaf{"SessRejAd", statsInfo.SessRejAd})
    statsInfo.EntityData.Leafs.Append("sess-rej-max-pdu", types.YLeaf{"SessRejMaxPdu", statsInfo.SessRejMaxPdu})
    statsInfo.EntityData.Leafs.Append("sess-rej-lr", types.YLeaf{"SessRejLr", statsInfo.SessRejLr})
    statsInfo.EntityData.Leafs.Append("bad-ldpid", types.YLeaf{"BadLdpid", statsInfo.BadLdpid})
    statsInfo.EntityData.Leafs.Append("bad-pdu-len", types.YLeaf{"BadPduLen", statsInfo.BadPduLen})
    statsInfo.EntityData.Leafs.Append("bad-msg-len", types.YLeaf{"BadMsgLen", statsInfo.BadMsgLen})
    statsInfo.EntityData.Leafs.Append("bad-tlv-len", types.YLeaf{"BadTlvLen", statsInfo.BadTlvLen})
    statsInfo.EntityData.Leafs.Append("malformed-tlv-val", types.YLeaf{"MalformedTlvVal", statsInfo.MalformedTlvVal})
    statsInfo.EntityData.Leafs.Append("keep-alive-exp", types.YLeaf{"KeepAliveExp", statsInfo.KeepAliveExp})
    statsInfo.EntityData.Leafs.Append("shutdown-notif-rec", types.YLeaf{"ShutdownNotifRec", statsInfo.ShutdownNotifRec})
    statsInfo.EntityData.Leafs.Append("shutdow-notif-sent", types.YLeaf{"ShutdowNotifSent", statsInfo.ShutdowNotifSent})

    statsInfo.EntityData.YListKeys = []string {}

    return &(statsInfo.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut
// MPLS LDP message sent counters to this neighbor.
type MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total count of all messages. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Init message count. The type is interface{} with range: 0..4294967295.
    InitCount interface{}

    // Address message count. The type is interface{} with range: 0..4294967295.
    AddressCount interface{}

    // Address withdraw count. The type is interface{} with range: 0..4294967295.
    AddressWithdrawCount interface{}

    // Label map count. The type is interface{} with range: 0..4294967295.
    LabelMapCount interface{}

    // Label withdraw count. The type is interface{} with range: 0..4294967295.
    LabelWithdrawCount interface{}

    // Label release count. The type is interface{} with range: 0..4294967295.
    LabelReleaseCount interface{}

    // Label request count. The type is interface{} with range: 0..4294967295.
    LabelRequestCount interface{}

    // Label abort request count. The type is interface{} with range:
    // 0..4294967295.
    LabelAbortRequestCount interface{}

    // Notification count. The type is interface{} with range: 0..4294967295.
    NotificationCount interface{}

    // Keepalive count. The type is interface{} with range: 0..4294967295.
    KeepAliveCount interface{}

    // ICCP RG Connect count. The type is interface{} with range: 0..4294967295.
    IccpRgConnCount interface{}

    // ICCP RG Disconnect count. The type is interface{} with range:
    // 0..4294967295.
    IccpRgDisconnCount interface{}

    // ICCP RG Notify count. The type is interface{} with range: 0..4294967295.
    IccpRgNotifCount interface{}

    // ICCP RG App Data count. The type is interface{} with range: 0..4294967295.
    IccpRgAppDataCount interface{}
}

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetEntityData() *types.CommonEntityData {
    messageOut.EntityData.YFilter = messageOut.YFilter
    messageOut.EntityData.YangName = "message-out"
    messageOut.EntityData.BundleName = "cisco_ios_xe"
    messageOut.EntityData.ParentYangName = "stats-info"
    messageOut.EntityData.SegmentPath = "message-out"
    messageOut.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    messageOut.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    messageOut.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    messageOut.EntityData.Children = types.NewOrderedMap()
    messageOut.EntityData.Leafs = types.NewOrderedMap()
    messageOut.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", messageOut.TotalCount})
    messageOut.EntityData.Leafs.Append("init-count", types.YLeaf{"InitCount", messageOut.InitCount})
    messageOut.EntityData.Leafs.Append("address-count", types.YLeaf{"AddressCount", messageOut.AddressCount})
    messageOut.EntityData.Leafs.Append("address-withdraw-count", types.YLeaf{"AddressWithdrawCount", messageOut.AddressWithdrawCount})
    messageOut.EntityData.Leafs.Append("label-map-count", types.YLeaf{"LabelMapCount", messageOut.LabelMapCount})
    messageOut.EntityData.Leafs.Append("label-withdraw-count", types.YLeaf{"LabelWithdrawCount", messageOut.LabelWithdrawCount})
    messageOut.EntityData.Leafs.Append("label-release-count", types.YLeaf{"LabelReleaseCount", messageOut.LabelReleaseCount})
    messageOut.EntityData.Leafs.Append("label-request-count", types.YLeaf{"LabelRequestCount", messageOut.LabelRequestCount})
    messageOut.EntityData.Leafs.Append("label-abort-request-count", types.YLeaf{"LabelAbortRequestCount", messageOut.LabelAbortRequestCount})
    messageOut.EntityData.Leafs.Append("notification-count", types.YLeaf{"NotificationCount", messageOut.NotificationCount})
    messageOut.EntityData.Leafs.Append("keep-alive-count", types.YLeaf{"KeepAliveCount", messageOut.KeepAliveCount})
    messageOut.EntityData.Leafs.Append("iccp-rg-conn-count", types.YLeaf{"IccpRgConnCount", messageOut.IccpRgConnCount})
    messageOut.EntityData.Leafs.Append("iccp-rg-disconn-count", types.YLeaf{"IccpRgDisconnCount", messageOut.IccpRgDisconnCount})
    messageOut.EntityData.Leafs.Append("iccp-rg-notif-count", types.YLeaf{"IccpRgNotifCount", messageOut.IccpRgNotifCount})
    messageOut.EntityData.Leafs.Append("iccp-rg-app-data-count", types.YLeaf{"IccpRgAppDataCount", messageOut.IccpRgAppDataCount})

    messageOut.EntityData.YListKeys = []string {}

    return &(messageOut.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn
// MPLS LDP message received counters from this
// neighbor.
type MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total count of all messages. The type is interface{} with range:
    // 0..4294967295.
    TotalCount interface{}

    // Init message count. The type is interface{} with range: 0..4294967295.
    InitCount interface{}

    // Address message count. The type is interface{} with range: 0..4294967295.
    AddressCount interface{}

    // Address withdraw count. The type is interface{} with range: 0..4294967295.
    AddressWithdrawCount interface{}

    // Label map count. The type is interface{} with range: 0..4294967295.
    LabelMapCount interface{}

    // Label withdraw count. The type is interface{} with range: 0..4294967295.
    LabelWithdrawCount interface{}

    // Label release count. The type is interface{} with range: 0..4294967295.
    LabelReleaseCount interface{}

    // Label request count. The type is interface{} with range: 0..4294967295.
    LabelRequestCount interface{}

    // Label abort request count. The type is interface{} with range:
    // 0..4294967295.
    LabelAbortRequestCount interface{}

    // Notification count. The type is interface{} with range: 0..4294967295.
    NotificationCount interface{}

    // Keepalive count. The type is interface{} with range: 0..4294967295.
    KeepAliveCount interface{}

    // ICCP RG Connect count. The type is interface{} with range: 0..4294967295.
    IccpRgConnCount interface{}

    // ICCP RG Disconnect count. The type is interface{} with range:
    // 0..4294967295.
    IccpRgDisconnCount interface{}

    // ICCP RG Notify count. The type is interface{} with range: 0..4294967295.
    IccpRgNotifCount interface{}

    // ICCP RG App Data count. The type is interface{} with range: 0..4294967295.
    IccpRgAppDataCount interface{}
}

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetEntityData() *types.CommonEntityData {
    messageIn.EntityData.YFilter = messageIn.YFilter
    messageIn.EntityData.YangName = "message-in"
    messageIn.EntityData.BundleName = "cisco_ios_xe"
    messageIn.EntityData.ParentYangName = "stats-info"
    messageIn.EntityData.SegmentPath = "message-in"
    messageIn.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    messageIn.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    messageIn.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    messageIn.EntityData.Children = types.NewOrderedMap()
    messageIn.EntityData.Leafs = types.NewOrderedMap()
    messageIn.EntityData.Leafs.Append("total-count", types.YLeaf{"TotalCount", messageIn.TotalCount})
    messageIn.EntityData.Leafs.Append("init-count", types.YLeaf{"InitCount", messageIn.InitCount})
    messageIn.EntityData.Leafs.Append("address-count", types.YLeaf{"AddressCount", messageIn.AddressCount})
    messageIn.EntityData.Leafs.Append("address-withdraw-count", types.YLeaf{"AddressWithdrawCount", messageIn.AddressWithdrawCount})
    messageIn.EntityData.Leafs.Append("label-map-count", types.YLeaf{"LabelMapCount", messageIn.LabelMapCount})
    messageIn.EntityData.Leafs.Append("label-withdraw-count", types.YLeaf{"LabelWithdrawCount", messageIn.LabelWithdrawCount})
    messageIn.EntityData.Leafs.Append("label-release-count", types.YLeaf{"LabelReleaseCount", messageIn.LabelReleaseCount})
    messageIn.EntityData.Leafs.Append("label-request-count", types.YLeaf{"LabelRequestCount", messageIn.LabelRequestCount})
    messageIn.EntityData.Leafs.Append("label-abort-request-count", types.YLeaf{"LabelAbortRequestCount", messageIn.LabelAbortRequestCount})
    messageIn.EntityData.Leafs.Append("notification-count", types.YLeaf{"NotificationCount", messageIn.NotificationCount})
    messageIn.EntityData.Leafs.Append("keep-alive-count", types.YLeaf{"KeepAliveCount", messageIn.KeepAliveCount})
    messageIn.EntityData.Leafs.Append("iccp-rg-conn-count", types.YLeaf{"IccpRgConnCount", messageIn.IccpRgConnCount})
    messageIn.EntityData.Leafs.Append("iccp-rg-disconn-count", types.YLeaf{"IccpRgDisconnCount", messageIn.IccpRgDisconnCount})
    messageIn.EntityData.Leafs.Append("iccp-rg-notif-count", types.YLeaf{"IccpRgNotifCount", messageIn.IccpRgNotifCount})
    messageIn.EntityData.Leafs.Append("iccp-rg-app-data-count", types.YLeaf{"IccpRgAppDataCount", messageIn.IccpRgAppDataCount})

    messageIn.EntityData.YListKeys = []string {}

    return &(messageIn.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_Backoffs
// LDP Backoff Information
type MplsLdp_MplsLdpState_Neighbors_Backoffs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Current neighbor backoff count in seconds. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    BackoffSeconds interface{}

    // Current neighbor backoff waiting count in seconds. The type is interface{}
    // with range: 0..4294967295. Units are seconds.
    WaitingSeconds interface{}
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetEntityData() *types.CommonEntityData {
    backoffs.EntityData.YFilter = backoffs.YFilter
    backoffs.EntityData.YangName = "backoffs"
    backoffs.EntityData.BundleName = "cisco_ios_xe"
    backoffs.EntityData.ParentYangName = "neighbors"
    backoffs.EntityData.SegmentPath = "backoffs"
    backoffs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    backoffs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    backoffs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    backoffs.EntityData.Children = types.NewOrderedMap()
    backoffs.EntityData.Leafs = types.NewOrderedMap()
    backoffs.EntityData.Leafs.Append("backoff-seconds", types.YLeaf{"BackoffSeconds", backoffs.BackoffSeconds})
    backoffs.EntityData.Leafs.Append("waiting-seconds", types.YLeaf{"WaitingSeconds", backoffs.WaitingSeconds})

    backoffs.EntityData.YListKeys = []string {}

    return &(backoffs.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail
// This is the LDP NSR state for this neighbor.
type MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Non-Stop Routing State. The type is one of the following:
    // NsrStatusReadyNsrStatusDisabledNsrStatusNotReady.
    NsrState interface{}

    // NSR Sync State. The type is one of the following:
    // LdpNsrPeerSyncStOperLdpNsrPeerSyncStNoneLdpNsrPeerSyncStPrepLdpNsrPeerSyncStWaitLdpNsrPeerSyncStReadyLdpNsrPeerSyncStAppWait.
    NsrNbrSyncState interface{}

    // This is the last NSR sync error recieved. It indicates the last reason the
    // sync failed even if the sync has now succeeded. This allows this
    // information to be viewed when the state is flapping, even if the
    // syncronization is successful at the time of the query. The type is one of
    // the following:
    // NsrPeerSyncErrAppFailNsrPeerSyncErrLdpGblNsrPeerSyncErrTcpGblNsrPeerSyncErrLdpPeerNsrPeerSyncErrTcpPeerNsrPeerSyncErrNoneNsrPeerSyncErrLdpSyncNackNsrPeerSyncErrSyncPrep.
    NsrNbrLastSyncError interface{}

    // Last NSR sync NACK reason. The type is one of the following:
    // NsrSyncNackRsnErrAppNotFoundNsrSyncNackRsnNoPEndSockNsrSyncNackRsnErrAdjAddNsrSyncNackRsnErrUnexpPeerDownNsrSyncNackRsnNoneNsrSyncNackRsnNoCtxNsrSyncNackRsnMissingElemNsrSyncNackRsnErrAppInvalidNsrSyncNackRsnErrPpCreateNsrSyncNackRsnErrRxUnexpOpenNsrSyncNackRsnTblIdMismatchNsrSyncNackRsnErrAddrBindNsrSyncNackRsnErrDhcAddNsrSyncNackRsnErrRxNotifNsrSyncNackRsnErrTpCreateNsrSyncNackRsnPEndSockNotSyncedNsrSyncNackRsnPpExistsNsrSyncNackRsnErrRxBadPieNsrSyncNackRsnEnomem.
    NsrNbrLastSyncNackReason interface{}

    // Pending Label-Request responses. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrPendLabelReqResps interface{}

    // Pending Label-Withdraw responses. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrPendLabelWithdrawResps interface{}

    // Pending Local Address Withdraw Acks:. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrPendLclAddrWithdrawAcks interface{}

    // In label Request Records created. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrInLabelReqsCreated interface{}

    // In label Request Records freed. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrInLabelReqsFreed interface{}

    // In label Withdraw Records created. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrInLabelWithdrawCreated interface{}

    // In label Withdraw Records freed. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrInLabelWithdrawFreed interface{}

    // Local Address Withdraw set. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrLclAddrWithdrawSet interface{}

    // Local Address Withdraw cleared. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrLclAddrWithdrawCleared interface{}

    // Transmit contexts enqueued. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrXmitCtxtEnq interface{}

    // Transmit contexts dequeued. The type is interface{} with range:
    // 0..4294967295.
    NsrNbrXmitCtxtDeq interface{}

    // If the value of this object is 0 (zero) then Loop Dection for Path Vectors
    // for this Peer is disabled.  Otherwise, if this object has a value greater
    // than zero, then Loop Dection for Path  Vectors for this Peer is enabled and
    // the Path Vector Limit is this value. The type is interface{} with range:
    // 0..255.
    PathVectorLimit interface{}

    // This container holds session information about the sessions between these
    // two neighbors.
    NbrSess MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess
}

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetEntityData() *types.CommonEntityData {
    nsrNbrDetail.EntityData.YFilter = nsrNbrDetail.YFilter
    nsrNbrDetail.EntityData.YangName = "nsr-nbr-detail"
    nsrNbrDetail.EntityData.BundleName = "cisco_ios_xe"
    nsrNbrDetail.EntityData.ParentYangName = "neighbors"
    nsrNbrDetail.EntityData.SegmentPath = "nsr-nbr-detail"
    nsrNbrDetail.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nsrNbrDetail.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nsrNbrDetail.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nsrNbrDetail.EntityData.Children = types.NewOrderedMap()
    nsrNbrDetail.EntityData.Children.Append("nbr-sess", types.YChild{"NbrSess", &nsrNbrDetail.NbrSess})
    nsrNbrDetail.EntityData.Leafs = types.NewOrderedMap()
    nsrNbrDetail.EntityData.Leafs.Append("nsr-state", types.YLeaf{"NsrState", nsrNbrDetail.NsrState})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-sync-state", types.YLeaf{"NsrNbrSyncState", nsrNbrDetail.NsrNbrSyncState})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-last-sync-error", types.YLeaf{"NsrNbrLastSyncError", nsrNbrDetail.NsrNbrLastSyncError})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-last-sync-nack-reason", types.YLeaf{"NsrNbrLastSyncNackReason", nsrNbrDetail.NsrNbrLastSyncNackReason})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-pend-label-req-resps", types.YLeaf{"NsrNbrPendLabelReqResps", nsrNbrDetail.NsrNbrPendLabelReqResps})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-pend-label-withdraw-resps", types.YLeaf{"NsrNbrPendLabelWithdrawResps", nsrNbrDetail.NsrNbrPendLabelWithdrawResps})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-pend-lcl-addr-withdraw-acks", types.YLeaf{"NsrNbrPendLclAddrWithdrawAcks", nsrNbrDetail.NsrNbrPendLclAddrWithdrawAcks})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-in-label-reqs-created", types.YLeaf{"NsrNbrInLabelReqsCreated", nsrNbrDetail.NsrNbrInLabelReqsCreated})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-in-label-reqs-freed", types.YLeaf{"NsrNbrInLabelReqsFreed", nsrNbrDetail.NsrNbrInLabelReqsFreed})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-in-label-withdraw-created", types.YLeaf{"NsrNbrInLabelWithdrawCreated", nsrNbrDetail.NsrNbrInLabelWithdrawCreated})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-in-label-withdraw-freed", types.YLeaf{"NsrNbrInLabelWithdrawFreed", nsrNbrDetail.NsrNbrInLabelWithdrawFreed})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-lcl-addr-withdraw-set", types.YLeaf{"NsrNbrLclAddrWithdrawSet", nsrNbrDetail.NsrNbrLclAddrWithdrawSet})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-lcl-addr-withdraw-cleared", types.YLeaf{"NsrNbrLclAddrWithdrawCleared", nsrNbrDetail.NsrNbrLclAddrWithdrawCleared})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-xmit-ctxt-enq", types.YLeaf{"NsrNbrXmitCtxtEnq", nsrNbrDetail.NsrNbrXmitCtxtEnq})
    nsrNbrDetail.EntityData.Leafs.Append("nsr-nbr-xmit-ctxt-deq", types.YLeaf{"NsrNbrXmitCtxtDeq", nsrNbrDetail.NsrNbrXmitCtxtDeq})
    nsrNbrDetail.EntityData.Leafs.Append("path-vector-limit", types.YLeaf{"PathVectorLimit", nsrNbrDetail.PathVectorLimit})

    nsrNbrDetail.EntityData.YListKeys = []string {}

    return &(nsrNbrDetail.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess
// This container holds session information about the
// sessions between these two neighbors.
type MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of sysUpTime at the time this Session entered its current state
    // as denoted by the SessionState object. The type is interface{} with range:
    // 0..4294967295.
    LastStatChange interface{}

    // The current state of the session, all of the states 1 to 5 are based on the
    // state machine for session negotiation behavior. The type is State.
    State interface{}

    // The keep alive hold time remaining for this session in seconds. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    KeepAliveRemain interface{}

    // The negotiated KeepAlive Time which represents the amount of seconds
    // between keep alive messages.  The EntityKeepAliveHoldTimer related to this
    // Session is the value that was proposed as the KeepAlive Time for this
    // session.  This value is negotiated during session initialization between
    // the entity's proposed value (i.e., the value configured in
    // EntityKeepAliveHoldTimer) and the peer's proposed KeepAlive Hold Timer
    // value. This value is the smaller of the two proposed values. The type is
    // interface{} with range: 1..65535. Units are seconds.
    KeepAliveTime interface{}

    // The value of maximum allowable length for LDP PDUs this session.  This
    // value may have been negotiated for during the Session Initialization.  This
    // object is related to the EntityMaxPduLength object.  The EntityMaxPduLength
    // object specifies the requested LDP PDU length, and this object reflects the
    // negotiated LDP PDU length between the Entity and the Peer. The type is
    // interface{} with range: 1..65535. Units are octets.
    MaxPdu interface{}

    // The value of sysUpTime on the most recent occasion at which any one or more
    // of this session's counters suffered a discontinuity.  The relevant counters
    // are the specific instances associated with this session of any counter32
    // object contained in the session-stats table.  The initial value of this
    // object is the value of sysUpTime when the entry was created in this table. 
    // Also, a command generator can distinguish when a session between a given
    // Entity and Peer goes away and a new session is established.  This value
    // would change and thus indicate to the command generator that this is a
    // different session. The type is interface{} with range: 0..4294967295.
    DisconTime interface{}

    // This object counts the number of Unknown Message Type Errors detected by
    // this LSR/LER during this session.  Discontinuities in the value of this
    // counter can occur at re-initialization of the management system, and at
    // other times as indicated by the value of discon-time. The type is
    // interface{} with range: 0..4294967295.
    UnknownMessErr interface{}

    // This object counts the number of Unknown TLV Errors detected by this
    // LSR/LER during this session.  Discontinuities in the value of this counter
    // can occur at re-initialization of the management system, and at other times
    // as indicated by the value of discon-time. The type is interface{} with
    // range: 0..4294967295.
    UnknownTlv interface{}
}

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetEntityData() *types.CommonEntityData {
    nbrSess.EntityData.YFilter = nbrSess.YFilter
    nbrSess.EntityData.YangName = "nbr-sess"
    nbrSess.EntityData.BundleName = "cisco_ios_xe"
    nbrSess.EntityData.ParentYangName = "nsr-nbr-detail"
    nbrSess.EntityData.SegmentPath = "nbr-sess"
    nbrSess.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nbrSess.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nbrSess.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nbrSess.EntityData.Children = types.NewOrderedMap()
    nbrSess.EntityData.Leafs = types.NewOrderedMap()
    nbrSess.EntityData.Leafs.Append("last-stat-change", types.YLeaf{"LastStatChange", nbrSess.LastStatChange})
    nbrSess.EntityData.Leafs.Append("state", types.YLeaf{"State", nbrSess.State})
    nbrSess.EntityData.Leafs.Append("keep-alive-remain", types.YLeaf{"KeepAliveRemain", nbrSess.KeepAliveRemain})
    nbrSess.EntityData.Leafs.Append("keep-alive-time", types.YLeaf{"KeepAliveTime", nbrSess.KeepAliveTime})
    nbrSess.EntityData.Leafs.Append("max-pdu", types.YLeaf{"MaxPdu", nbrSess.MaxPdu})
    nbrSess.EntityData.Leafs.Append("discon-time", types.YLeaf{"DisconTime", nbrSess.DisconTime})
    nbrSess.EntityData.Leafs.Append("unknown-mess-err", types.YLeaf{"UnknownMessErr", nbrSess.UnknownMessErr})
    nbrSess.EntityData.Leafs.Append("unknown-tlv", types.YLeaf{"UnknownTlv", nbrSess.UnknownTlv})

    nbrSess.EntityData.YListKeys = []string {}

    return &(nbrSess.EntityData)
}

// MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State represents for session negotiation behavior.
type MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State string

const (
    // State: nonexistent.
    MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State_nonexistent MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State = "nonexistent"

    // State: initialized.
    MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State_initialized MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State = "initialized"

    // State: openrec.
    MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State_openrec MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State = "openrec"

    // State: opensent.
    MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State_opensent MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State = "opensent"

    // State: operational.
    MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State_operational MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess_State = "operational"
)

// MplsLdp_MplsLdpState_LabelRanges
// This contaions holds all the label ranges in use
// by this LDP instance.
type MplsLdp_MplsLdpState_LabelRanges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry contains a single range of labels represented by the configured
    // Upper and Lower Bounds pairs.  NOTE: there is NO corresponding LDP message
    // which relates to the information in this table, however, this table does
    // provide a way for a user to 'reserve' a generic label range.  NOTE:  The
    // ranges for a specific LDP Entity are UNIQUE and non-overlapping. The type
    // is slice of MplsLdp_MplsLdpState_LabelRanges_LabelRange.
    LabelRange []*MplsLdp_MplsLdpState_LabelRanges_LabelRange
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetEntityData() *types.CommonEntityData {
    labelRanges.EntityData.YFilter = labelRanges.YFilter
    labelRanges.EntityData.YangName = "label-ranges"
    labelRanges.EntityData.BundleName = "cisco_ios_xe"
    labelRanges.EntityData.ParentYangName = "mpls-ldp-state"
    labelRanges.EntityData.SegmentPath = "label-ranges"
    labelRanges.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelRanges.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelRanges.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelRanges.EntityData.Children = types.NewOrderedMap()
    labelRanges.EntityData.Children.Append("label-range", types.YChild{"LabelRange", nil})
    for i := range labelRanges.LabelRange {
        labelRanges.EntityData.Children.Append(types.GetSegmentPath(labelRanges.LabelRange[i]), types.YChild{"LabelRange", labelRanges.LabelRange[i]})
    }
    labelRanges.EntityData.Leafs = types.NewOrderedMap()

    labelRanges.EntityData.YListKeys = []string {}

    return &(labelRanges.EntityData)
}

// MplsLdp_MplsLdpState_LabelRanges_LabelRange
// This entry contains a single range of labels
// represented by the configured Upper and Lower
// Bounds pairs.  NOTE: there is NO corresponding
// LDP message which relates to the information
// in this table, however, this table does provide
// a way for a user to 'reserve' a generic label
// range.
// 
// NOTE:  The ranges for a specific LDP Entity
// are UNIQUE and non-overlapping.
type MplsLdp_MplsLdpState_LabelRanges_LabelRange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The minimum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    LrMin interface{}

    // This attribute is a key. The maximum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    LrMax interface{}
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetEntityData() *types.CommonEntityData {
    labelRange.EntityData.YFilter = labelRange.YFilter
    labelRange.EntityData.YangName = "label-range"
    labelRange.EntityData.BundleName = "cisco_ios_xe"
    labelRange.EntityData.ParentYangName = "label-ranges"
    labelRange.EntityData.SegmentPath = "label-range" + types.AddKeyToken(labelRange.LrMin, "lr-min") + types.AddKeyToken(labelRange.LrMax, "lr-max")
    labelRange.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelRange.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelRange.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelRange.EntityData.Children = types.NewOrderedMap()
    labelRange.EntityData.Leafs = types.NewOrderedMap()
    labelRange.EntityData.Leafs.Append("lr-min", types.YLeaf{"LrMin", labelRange.LrMin})
    labelRange.EntityData.Leafs.Append("lr-max", types.YLeaf{"LrMax", labelRange.LrMax})

    labelRange.EntityData.YListKeys = []string {"LrMin", "LrMax"}

    return &(labelRange.EntityData)
}

// MplsLdp_MplsLdpConfig
// MPLS LDP Configuration.
type MplsLdp_MplsLdpConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This contains hold all MPLS LDP Configuration with Global scope. These
    // values affect the entire LSR unless overiddden by a parameter with a more
    // localized scope.
    GlobalCfg MplsLdp_MplsLdpConfig_GlobalCfg

    // This container holds the list of neighbor configuration parameters.
    NbrTable MplsLdp_MplsLdpConfig_NbrTable

    // This holds the MPLS LDP password configuration for use with LDP neighbors.
    Passwords MplsLdp_MplsLdpConfig_Passwords

    // Configure session parameters.
    Session MplsLdp_MplsLdpConfig_Session

    // This container holds the label allocation and advertisement configuration
    // for the LDP Label Information Base. These control what prefixes may be
    // allocated and advertised to peers.
    LabelCfg MplsLdp_MplsLdpConfig_LabelCfg

    // LDP discovery.
    Discovery MplsLdp_MplsLdpConfig_Discovery

    // Configure LDP Graceful Restart.
    GracefulRestart MplsLdp_MplsLdpConfig_GracefulRestart

    // Enable LDP logging.
    Logging MplsLdp_MplsLdpConfig_Logging

    // MPLS LDP Interface configuration commands.
    Interfaces MplsLdp_MplsLdpConfig_Interfaces

    // This containter provides the MPLS LDP config for routing protocols from
    // which it can obtain addresses to associate with labels.
    Routing MplsLdp_MplsLdpConfig_Routing

    // This container holds the configuration of dual IPv4 and IPv6 stack peers.
    DualStack MplsLdp_MplsLdpConfig_DualStack
}

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetEntityData() *types.CommonEntityData {
    mplsLdpConfig.EntityData.YFilter = mplsLdpConfig.YFilter
    mplsLdpConfig.EntityData.YangName = "mpls-ldp-config"
    mplsLdpConfig.EntityData.BundleName = "cisco_ios_xe"
    mplsLdpConfig.EntityData.ParentYangName = "mpls-ldp"
    mplsLdpConfig.EntityData.SegmentPath = "mpls-ldp-config"
    mplsLdpConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    mplsLdpConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    mplsLdpConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    mplsLdpConfig.EntityData.Children = types.NewOrderedMap()
    mplsLdpConfig.EntityData.Children.Append("global-cfg", types.YChild{"GlobalCfg", &mplsLdpConfig.GlobalCfg})
    mplsLdpConfig.EntityData.Children.Append("nbr-table", types.YChild{"NbrTable", &mplsLdpConfig.NbrTable})
    mplsLdpConfig.EntityData.Children.Append("passwords", types.YChild{"Passwords", &mplsLdpConfig.Passwords})
    mplsLdpConfig.EntityData.Children.Append("session", types.YChild{"Session", &mplsLdpConfig.Session})
    mplsLdpConfig.EntityData.Children.Append("label-cfg", types.YChild{"LabelCfg", &mplsLdpConfig.LabelCfg})
    mplsLdpConfig.EntityData.Children.Append("discovery", types.YChild{"Discovery", &mplsLdpConfig.Discovery})
    mplsLdpConfig.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &mplsLdpConfig.GracefulRestart})
    mplsLdpConfig.EntityData.Children.Append("logging", types.YChild{"Logging", &mplsLdpConfig.Logging})
    mplsLdpConfig.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &mplsLdpConfig.Interfaces})
    mplsLdpConfig.EntityData.Children.Append("routing", types.YChild{"Routing", &mplsLdpConfig.Routing})
    mplsLdpConfig.EntityData.Children.Append("dual-stack", types.YChild{"DualStack", &mplsLdpConfig.DualStack})
    mplsLdpConfig.EntityData.Leafs = types.NewOrderedMap()

    mplsLdpConfig.EntityData.YListKeys = []string {}

    return &(mplsLdpConfig.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg
// This contains hold all MPLS LDP Configuration with Global
// scope. These values affect the entire LSR unless
// overiddden by a parameter with a more localized scope.
type MplsLdp_MplsLdpConfig_GlobalCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Writing this leaf tears down all LDP sessions, withdraws all outgoing
    // labels from the forwarding plane, and frees all local labels that have been
    // allocated. The type is interface{}.
    Shutdown interface{}

    // This leaf controls whether Non-Stop-Routing should be enabled to include
    // MPLS LDP. The type is bool.
    EnableNsr interface{}

    // When set to true, disable LDP discovery's quick start mode for this LSR.
    // The type is bool.
    DisableQuickStart interface{}

    // This leaf enables or disables Loop Detection globally for the LSR. The type
    // is bool.
    LoopDetection interface{}

    // This leaf controls the administrative status of LDP for this LSR. If set to
    // disable, then all LDP activity will be disabled and all LDP sessions with
    // peers will terminate. The LDP configuration will remain intact.  When the
    // admin status is set back to 'enable', then LDP will resume operations and
    // attempt to establish new sessions with the peers. The type is AdminStatus.
    AdminStatus interface{}

    // This sets the 6-bit Differentiated Services Code Point (DSCP) value in the
    // TCP packets for LDP messages being sent from the LSR. The type is
    // interface{} with range: 0..63.
    DcspVal interface{}

    // This sets the priority within the LSR for TCP packets for LDP messages
    // being sent from the LSR. They are given a higher transmission priorty and
    // will avoid being queued behind lower priority messages. The type is bool.
    HighPriority interface{}

    // Time in seconds to delay IGP sync after session comes up. The type is
    // interface{} with range: 5..300. Units are second.
    Seconds interface{}

    // This choice causes IGP sync up immediately upon session up. The type is
    // interface{}.
    DisableDelay interface{}

    // Time in seconds to delay IGP sync after session comes up. The type is
    // interface{} with range: 5..300. Units are second.
    SecondsDelayProc interface{}

    // This choice causes IGP sync up immediately upon session up. The type is
    // interface{}.
    DisableDelayProc interface{}

    // This leaf defines the protocol to be used. The default is LDP. The type is
    // Protocol.
    Protocol interface{}

    // When attempting to establish a session with a given Peer, the given LDP
    // Entity should send out the YANG notification, 'init-sess-thresh-ex', when
    // the number of Session Initialization messages sent exceeds this threshold. 
    // The notification is used to notify an operator when this Entity and its
    // Peer are possibly engaged in an endless sequence of messages as each NAKs
    // the other's  Initialization messages with Error Notification messages. 
    // Setting this threshold which triggers the notification is one way to notify
    // the operator.  The notification should be generated each time this
    // threshold is exceeded and for every subsequent Initialization message which
    // is NAK'd with an Error Notification message after this threshold is
    // exceeded.  A value of 0 (zero) for this object indicates that the threshold
    // is infinity, thus the YANG notification will never be generated. The type
    // is interface{} with range: 0..100.
    InitSessThresh interface{}

    // Configuration for LDP Router ID (LDP ID). The type is slice of
    // MplsLdp_MplsLdpConfig_GlobalCfg_RouterId.
    RouterId []*MplsLdp_MplsLdpConfig_GlobalCfg_RouterId

    // Configure session parameters. Session parameters effect the session between
    // LDP peers once the session has been established.
    Session MplsLdp_MplsLdpConfig_GlobalCfg_Session

    // This container holds the global per address family configuration.
    PerAf MplsLdp_MplsLdpConfig_GlobalCfg_PerAf
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetEntityData() *types.CommonEntityData {
    globalCfg.EntityData.YFilter = globalCfg.YFilter
    globalCfg.EntityData.YangName = "global-cfg"
    globalCfg.EntityData.BundleName = "cisco_ios_xe"
    globalCfg.EntityData.ParentYangName = "mpls-ldp-config"
    globalCfg.EntityData.SegmentPath = "global-cfg"
    globalCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    globalCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    globalCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    globalCfg.EntityData.Children = types.NewOrderedMap()
    globalCfg.EntityData.Children.Append("router-id", types.YChild{"RouterId", nil})
    for i := range globalCfg.RouterId {
        globalCfg.EntityData.Children.Append(types.GetSegmentPath(globalCfg.RouterId[i]), types.YChild{"RouterId", globalCfg.RouterId[i]})
    }
    globalCfg.EntityData.Children.Append("session", types.YChild{"Session", &globalCfg.Session})
    globalCfg.EntityData.Children.Append("per-af", types.YChild{"PerAf", &globalCfg.PerAf})
    globalCfg.EntityData.Leafs = types.NewOrderedMap()
    globalCfg.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", globalCfg.Shutdown})
    globalCfg.EntityData.Leafs.Append("enable-nsr", types.YLeaf{"EnableNsr", globalCfg.EnableNsr})
    globalCfg.EntityData.Leafs.Append("disable-quick-start", types.YLeaf{"DisableQuickStart", globalCfg.DisableQuickStart})
    globalCfg.EntityData.Leafs.Append("loop-detection", types.YLeaf{"LoopDetection", globalCfg.LoopDetection})
    globalCfg.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", globalCfg.AdminStatus})
    globalCfg.EntityData.Leafs.Append("dcsp-val", types.YLeaf{"DcspVal", globalCfg.DcspVal})
    globalCfg.EntityData.Leafs.Append("high-priority", types.YLeaf{"HighPriority", globalCfg.HighPriority})
    globalCfg.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", globalCfg.Seconds})
    globalCfg.EntityData.Leafs.Append("disable-delay", types.YLeaf{"DisableDelay", globalCfg.DisableDelay})
    globalCfg.EntityData.Leafs.Append("seconds-delay-proc", types.YLeaf{"SecondsDelayProc", globalCfg.SecondsDelayProc})
    globalCfg.EntityData.Leafs.Append("disable-delay-proc", types.YLeaf{"DisableDelayProc", globalCfg.DisableDelayProc})
    globalCfg.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", globalCfg.Protocol})
    globalCfg.EntityData.Leafs.Append("init-sess-thresh", types.YLeaf{"InitSessThresh", globalCfg.InitSessThresh})

    globalCfg.EntityData.YListKeys = []string {}

    return &(globalCfg.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_RouterId
// Configuration for LDP Router ID (LDP ID)
type MplsLdp_MplsLdpConfig_GlobalCfg_RouterId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This defines the interface to use for the LDP LSR identifier address for
    // all sessions. The IP address of this interface will be used as the
    // identifier. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    LsrIdIf interface{}

    // This is the IP address to be used as the LDP LSR ID for all sessions. The
    // type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LsrIdIp interface{}

    // Force the router to use the specified identifier as the router ID more
    // quickly. The type is interface{}.
    Force interface{}
}

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetEntityData() *types.CommonEntityData {
    routerId.EntityData.YFilter = routerId.YFilter
    routerId.EntityData.YangName = "router-id"
    routerId.EntityData.BundleName = "cisco_ios_xe"
    routerId.EntityData.ParentYangName = "global-cfg"
    routerId.EntityData.SegmentPath = "router-id" + types.AddKeyToken(routerId.VrfName, "vrf-name")
    routerId.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routerId.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routerId.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routerId.EntityData.Children = types.NewOrderedMap()
    routerId.EntityData.Leafs = types.NewOrderedMap()
    routerId.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", routerId.VrfName})
    routerId.EntityData.Leafs.Append("lsr-id-if", types.YLeaf{"LsrIdIf", routerId.LsrIdIf})
    routerId.EntityData.Leafs.Append("lsr-id-ip", types.YLeaf{"LsrIdIp", routerId.LsrIdIp})
    routerId.EntityData.Leafs.Append("force", types.YLeaf{"Force", routerId.Force})

    routerId.EntityData.YListKeys = []string {"VrfName"}

    return &(routerId.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_Session
// Configure session parameters. Session parameters effect
// the session between LDP peers once the session has been
// established.
type MplsLdp_MplsLdpConfig_GlobalCfg_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial session backoff time (seconds). The LDP backoff mechanism prevents
    // two incompatibly configured label switch routers (LSRs) from engaging in an
    // unthrottled sequence of session setup failures.  For example, an
    // incompatibility arises when two neighboring routers attempt to perform
    // LC-ATM (label-controlled ATM) but the two are using different ranges of
    // VPI/VCI values for labels.  If a session setup attempt fails due to an
    // incompatibility, each LSR delays its next attempt (that is, backs off),
    // increasing the delay exponentially with each successive failure until the
    // maximum backoff delay is reached.  The default settings correspond to the
    // lowest settings for initial and maximum backoff values defined by the LDP
    // protocol specification. You should change the settings from the default
    // values only if such settings result in undesirable behavior. The type is
    // interface{} with range: 0..4294967295. Units are seconds. The default value
    // is 15.
    BackoffInit interface{}

    // The maximum session backoff time (seconds) The LDP backoff mechanism
    // prevents two incompatibly configured label switch routers (LSRs) from
    // engaging in an unthrottled sequence of session setup failures.  For
    // example, an incompatibility arises when two neighboring routers attempt to
    // perform LC-ATM (label-controlled ATM) but the two are using different
    // ranges of VPI/VCI values for labels.  If a session setup attempt fails due
    // to an incompatibility, each LSR delays its next attempt (that is, backs
    // off), increasing the delay exponentially with each successive failure until
    // the maximum backoff delay is reached.  The default settings correspond to
    // the lowest settings for initial and maximum backoff values defined by the
    // LDP protocol specification. You should change the settings from the default
    // values only if such settings result in undesirable behavior. The type is
    // interface{} with range: 0..4294967295. Units are seconds. The default value
    // is 15.
    BackoffMax interface{}

    // Number from 15 to 2147483, that defines the time, in seconds, an LDP
    // session is maintained in the absence of LDP messages from the session peer.
    // The type is interface{} with range: 0..65535.
    Seconds interface{}

    // If set to true, the session is held indefinitely in the absence of LDP
    // messages from the peer. The type is bool.
    Infinite interface{}

    // This container holds config for Downstream on Demand. For it to be enabled,
    // the Downstream on demand feature has to be configured on both peers of the
    // session. If only one peer in the session has downstream-on-demand feature
    // configured, then the session does not use downstream-on-demand mode. If,
    // after, a label request is sent, and no remote label is received from the
    // peer, the router will periodically resend the label request. After the peer
    // advertises a label after receiving the label request, it will automatically
    // readvertise the label if any label attribute changes subsequently. The type
    // is slice of MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand.
    DownstreamOnDemand []*MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand

    // Configure Session Protection parameters.
    Protection MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xe"
    session.EntityData.ParentYangName = "global-cfg"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("downstream-on-demand", types.YChild{"DownstreamOnDemand", nil})
    for i := range session.DownstreamOnDemand {
        session.EntityData.Children.Append(types.GetSegmentPath(session.DownstreamOnDemand[i]), types.YChild{"DownstreamOnDemand", session.DownstreamOnDemand[i]})
    }
    session.EntityData.Children.Append("protection", types.YChild{"Protection", &session.Protection})
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("backoff-init", types.YLeaf{"BackoffInit", session.BackoffInit})
    session.EntityData.Leafs.Append("backoff-max", types.YLeaf{"BackoffMax", session.BackoffMax})
    session.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", session.Seconds})
    session.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", session.Infinite})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand
// This container holds config for Downstream on Demand.
// For it to be enabled, the Downstream on demand
// feature has to be configured on both peers of the
// session. If only one peer in the session has
// downstream-on-demand feature configured, then the
// session does not use downstream-on-demand mode.
// If, after, a label request is sent, and no remote
// label is received from the peer, the router will
// periodically resend the label request. After the
// peer advertises a label after receiving the label
// request, it will automatically readvertise the label
// if any label attribute changes subsequently.
type MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // Enable Downstream on Demand for this LSR. In this mode a label is not
    // advertised to a peer, unless the peer explicitly requests it. At the same
    // time, since the peer does not automatically advertise labels, the label
    // request is sent whenever the next-hop points out to a peer that no remote
    // label has been assigned. The type is bool.
    Enabled interface{}

    // This filter contains a list of peer IDs that are configured for
    // downstream-on-demand mode. When the filter is changed or configured, the
    // list of established neighbors is traversed. If a session's
    // downstream-on-demand configuration has changed, the session is reset in
    // order that the new down-stream-on-demand mode can be configured. The reason
    // for resetting the session is to ensure that the labels are properly
    // advertised between the peers. When a new session is established, the ACL is
    // verified to determine whether the session should negotiate for
    // downstream-on-demand mode. If the filter string is configured and the
    // corresponding filter does not exist or is empty, then downstream-on-demand
    // mode is not configured for any neighbor. The filter type is device specific
    // and could be an ACL, a prefix list, or other mechanism. The type is string.
    Filter interface{}
}

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetEntityData() *types.CommonEntityData {
    downstreamOnDemand.EntityData.YFilter = downstreamOnDemand.YFilter
    downstreamOnDemand.EntityData.YangName = "downstream-on-demand"
    downstreamOnDemand.EntityData.BundleName = "cisco_ios_xe"
    downstreamOnDemand.EntityData.ParentYangName = "session"
    downstreamOnDemand.EntityData.SegmentPath = "downstream-on-demand" + types.AddKeyToken(downstreamOnDemand.VrfName, "vrf-name")
    downstreamOnDemand.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    downstreamOnDemand.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    downstreamOnDemand.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    downstreamOnDemand.EntityData.Children = types.NewOrderedMap()
    downstreamOnDemand.EntityData.Leafs = types.NewOrderedMap()
    downstreamOnDemand.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", downstreamOnDemand.VrfName})
    downstreamOnDemand.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", downstreamOnDemand.Enabled})
    downstreamOnDemand.EntityData.Leafs.Append("filter", types.YLeaf{"Filter", downstreamOnDemand.Filter})

    downstreamOnDemand.EntityData.YListKeys = []string {"VrfName"}

    return &(downstreamOnDemand.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection
// Configure Session Protection parameters
type MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is set true to enable session protection. The type is bool.
    EnableProt interface{}

    // This is an optional filter to restrict session protection. If the string is
    // null or unconfigured then session protection applied to all peers. The
    // filter type is device specific and could be an ACL, a prefix list, or other
    // mechanism. The type is string.
    PeerFilter interface{}

    // This is the sessiom holdup duration in seconds. The type is interface{}
    // with range: 30..2147483.
    Seconds interface{}

    // This sessiom holdup duration is infinite. The type is interface{}.
    Inf interface{}
}

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetEntityData() *types.CommonEntityData {
    protection.EntityData.YFilter = protection.YFilter
    protection.EntityData.YangName = "protection"
    protection.EntityData.BundleName = "cisco_ios_xe"
    protection.EntityData.ParentYangName = "session"
    protection.EntityData.SegmentPath = "protection"
    protection.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    protection.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    protection.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    protection.EntityData.Children = types.NewOrderedMap()
    protection.EntityData.Leafs = types.NewOrderedMap()
    protection.EntityData.Leafs.Append("enable-prot", types.YLeaf{"EnableProt", protection.EnableProt})
    protection.EntityData.Leafs.Append("peer-filter", types.YLeaf{"PeerFilter", protection.PeerFilter})
    protection.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", protection.Seconds})
    protection.EntityData.Leafs.Append("inf", types.YLeaf{"Inf", protection.Inf})

    protection.EntityData.YListKeys = []string {}

    return &(protection.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_PerAf
// This container holds the global per address family
// configuration.
type MplsLdp_MplsLdpConfig_GlobalCfg_PerAf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This container holds the global per address family configuration. The type
    // is slice of MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg.
    AfCfg []*MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetEntityData() *types.CommonEntityData {
    perAf.EntityData.YFilter = perAf.YFilter
    perAf.EntityData.YangName = "per-af"
    perAf.EntityData.BundleName = "cisco_ios_xe"
    perAf.EntityData.ParentYangName = "global-cfg"
    perAf.EntityData.SegmentPath = "per-af"
    perAf.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    perAf.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    perAf.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    perAf.EntityData.Children = types.NewOrderedMap()
    perAf.EntityData.Children.Append("af-cfg", types.YChild{"AfCfg", nil})
    for i := range perAf.AfCfg {
        perAf.EntityData.Children.Append(types.GetSegmentPath(perAf.AfCfg[i]), types.YChild{"AfCfg", perAf.AfCfg[i]})
    }
    perAf.EntityData.Leafs = types.NewOrderedMap()

    perAf.EntityData.YListKeys = []string {}

    return &(perAf.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg
// This container holds the global per address family
// configuration.
type MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // When set true, this enables MPLS forwarding for the ip default route. The
    // type is bool.
    DefaultRoute interface{}

    // Advertise this address as the explicit address in LDP discovery hello
    // messages and use it for LDP transport. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipaddr interface{}

    // Advertise this interface's address as the explicit address in LDP discovery
    // hello messages and use it for LDP transport. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // Do not advertise an explicit address in LDP discovery hello messages or
    // advertise a default address. Use the default address for LDP transport. The
    // type is interface{}.
    Implicit interface{}
}

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetEntityData() *types.CommonEntityData {
    afCfg.EntityData.YFilter = afCfg.YFilter
    afCfg.EntityData.YangName = "af-cfg"
    afCfg.EntityData.BundleName = "cisco_ios_xe"
    afCfg.EntityData.ParentYangName = "per-af"
    afCfg.EntityData.SegmentPath = "af-cfg" + types.AddKeyToken(afCfg.VrfName, "vrf-name") + types.AddKeyToken(afCfg.AfName, "af-name")
    afCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    afCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    afCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    afCfg.EntityData.Children = types.NewOrderedMap()
    afCfg.EntityData.Leafs = types.NewOrderedMap()
    afCfg.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", afCfg.VrfName})
    afCfg.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", afCfg.AfName})
    afCfg.EntityData.Leafs.Append("default-route", types.YLeaf{"DefaultRoute", afCfg.DefaultRoute})
    afCfg.EntityData.Leafs.Append("ipaddr", types.YLeaf{"Ipaddr", afCfg.Ipaddr})
    afCfg.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", afCfg.Interface})
    afCfg.EntityData.Leafs.Append("implicit", types.YLeaf{"Implicit", afCfg.Implicit})

    afCfg.EntityData.YListKeys = []string {"VrfName", "AfName"}

    return &(afCfg.EntityData)
}

// MplsLdp_MplsLdpConfig_GlobalCfg_AdminStatus represents sessions with the peers.
type MplsLdp_MplsLdpConfig_GlobalCfg_AdminStatus string

const (
    // Enable LDP globablly on this LSR.
    MplsLdp_MplsLdpConfig_GlobalCfg_AdminStatus_enable MplsLdp_MplsLdpConfig_GlobalCfg_AdminStatus = "enable"

    // Disable LDP globablly on this LSR.
    MplsLdp_MplsLdpConfig_GlobalCfg_AdminStatus_disable MplsLdp_MplsLdpConfig_GlobalCfg_AdminStatus = "disable"
)

// MplsLdp_MplsLdpConfig_GlobalCfg_Protocol represents is LDP.
type MplsLdp_MplsLdpConfig_GlobalCfg_Protocol string

const (
    // This LSR should use the LDP tagging protocol.
    MplsLdp_MplsLdpConfig_GlobalCfg_Protocol_ldp MplsLdp_MplsLdpConfig_GlobalCfg_Protocol = "ldp"

    // This LSR should use the TDP tagging protocol.
    MplsLdp_MplsLdpConfig_GlobalCfg_Protocol_tdp MplsLdp_MplsLdpConfig_GlobalCfg_Protocol = "tdp"

    // This LSR should use the both LDP and TDP tagging
    // protocol.
    MplsLdp_MplsLdpConfig_GlobalCfg_Protocol_both MplsLdp_MplsLdpConfig_GlobalCfg_Protocol = "both"
)

// MplsLdp_MplsLdpConfig_NbrTable
// This container holds the list of neighbor configuration
// parameters.
type MplsLdp_MplsLdpConfig_NbrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry holds the configuration of a single neighbor identified by the
    // IP address of that neighbor. The type is slice of
    // MplsLdp_MplsLdpConfig_NbrTable_NbrCfg.
    NbrCfg []*MplsLdp_MplsLdpConfig_NbrTable_NbrCfg
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetEntityData() *types.CommonEntityData {
    nbrTable.EntityData.YFilter = nbrTable.YFilter
    nbrTable.EntityData.YangName = "nbr-table"
    nbrTable.EntityData.BundleName = "cisco_ios_xe"
    nbrTable.EntityData.ParentYangName = "mpls-ldp-config"
    nbrTable.EntityData.SegmentPath = "nbr-table"
    nbrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nbrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nbrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nbrTable.EntityData.Children = types.NewOrderedMap()
    nbrTable.EntityData.Children.Append("nbr-cfg", types.YChild{"NbrCfg", nil})
    for i := range nbrTable.NbrCfg {
        nbrTable.EntityData.Children.Append(types.GetSegmentPath(nbrTable.NbrCfg[i]), types.YChild{"NbrCfg", nbrTable.NbrCfg[i]})
    }
    nbrTable.EntityData.Leafs = types.NewOrderedMap()

    nbrTable.EntityData.YListKeys = []string {}

    return &(nbrTable.EntityData)
}

// MplsLdp_MplsLdpConfig_NbrTable_NbrCfg
// This entry holds the configuration of a single neighbor
// identified by the IP address of that neighbor.
type MplsLdp_MplsLdpConfig_NbrTable_NbrCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    NbrVrf interface{}

    // This attribute is a key. The IP address for the LDP neighbor. This may be
    // IPv4 or IPv6. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NbrIp interface{}

    // The administrative status of this neighbor. If this object is changed from
    // 'enable' to 'disable' and this entity has already attempted to establish
    // contact with a neighbor, a 'tear-down' for that session is issued and the
    // session and all information related to that session ceases to exist).  When
    // the admin status is set back to 'enable', then this Entity will attempt to
    // establish a new session with the neighbor. The type is AdminStatus.
    AdminStatus interface{}

    // Enable LDP implicit withdraw label for this peer. The type is bool.
    ImplicitWithdraw interface{}

    // Establish or delete a targeted session. The type is bool.
    Targeted interface{}

    // This leaf defines the protocol to be used. The default is LDP. The type is
    // LabelProtocol.
    LabelProtocol interface{}

    // Accept only labels matching this filter. The filter type is device specific
    // and could be an ACL, a prefix list, or other mechanism. The type is string.
    LabelBindingFilter interface{}

    // Enables password authentication and stores the password using a
    // cryptographic hash. The type is string.
    Password interface{}
}

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetEntityData() *types.CommonEntityData {
    nbrCfg.EntityData.YFilter = nbrCfg.YFilter
    nbrCfg.EntityData.YangName = "nbr-cfg"
    nbrCfg.EntityData.BundleName = "cisco_ios_xe"
    nbrCfg.EntityData.ParentYangName = "nbr-table"
    nbrCfg.EntityData.SegmentPath = "nbr-cfg" + types.AddKeyToken(nbrCfg.NbrVrf, "nbr-vrf") + types.AddKeyToken(nbrCfg.NbrIp, "nbr-ip")
    nbrCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    nbrCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    nbrCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    nbrCfg.EntityData.Children = types.NewOrderedMap()
    nbrCfg.EntityData.Leafs = types.NewOrderedMap()
    nbrCfg.EntityData.Leafs.Append("nbr-vrf", types.YLeaf{"NbrVrf", nbrCfg.NbrVrf})
    nbrCfg.EntityData.Leafs.Append("nbr-ip", types.YLeaf{"NbrIp", nbrCfg.NbrIp})
    nbrCfg.EntityData.Leafs.Append("admin-status", types.YLeaf{"AdminStatus", nbrCfg.AdminStatus})
    nbrCfg.EntityData.Leafs.Append("implicit-withdraw", types.YLeaf{"ImplicitWithdraw", nbrCfg.ImplicitWithdraw})
    nbrCfg.EntityData.Leafs.Append("targeted", types.YLeaf{"Targeted", nbrCfg.Targeted})
    nbrCfg.EntityData.Leafs.Append("label-protocol", types.YLeaf{"LabelProtocol", nbrCfg.LabelProtocol})
    nbrCfg.EntityData.Leafs.Append("label-binding-filter", types.YLeaf{"LabelBindingFilter", nbrCfg.LabelBindingFilter})
    nbrCfg.EntityData.Leafs.Append("password", types.YLeaf{"Password", nbrCfg.Password})

    nbrCfg.EntityData.YListKeys = []string {"NbrVrf", "NbrIp"}

    return &(nbrCfg.EntityData)
}

// MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_AdminStatus represents with the neighbor.
type MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_AdminStatus string

const (
    // Set the administrative status of this neighbor
    // to enabled.
    MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_AdminStatus_enable MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_AdminStatus = "enable"

    // Set the administrative status of this neighbor
    // to disabled.
    MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_AdminStatus_disable MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_AdminStatus = "disable"
)

// MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_LabelProtocol represents is LDP.
type MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_LabelProtocol string

const (
    // This LSR should use the LDP tagging protocol.
    MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_LabelProtocol_ldp MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_LabelProtocol = "ldp"

    // This LSR should use the TDP tagging protocol.
    MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_LabelProtocol_tdp MplsLdp_MplsLdpConfig_NbrTable_NbrCfg_LabelProtocol = "tdp"
)

// MplsLdp_MplsLdpConfig_Passwords
// This holds the MPLS LDP password configuration for use
// with LDP neighbors.
type MplsLdp_MplsLdpConfig_Passwords struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This holds the MPLS LDP password configuration for use with a single LDP
    // neighbor or group of LDP neighbors. The type is slice of
    // MplsLdp_MplsLdpConfig_Passwords_Password.
    Password []*MplsLdp_MplsLdpConfig_Passwords_Password
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetEntityData() *types.CommonEntityData {
    passwords.EntityData.YFilter = passwords.YFilter
    passwords.EntityData.YangName = "passwords"
    passwords.EntityData.BundleName = "cisco_ios_xe"
    passwords.EntityData.ParentYangName = "mpls-ldp-config"
    passwords.EntityData.SegmentPath = "passwords"
    passwords.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    passwords.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    passwords.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    passwords.EntityData.Children = types.NewOrderedMap()
    passwords.EntityData.Children.Append("password", types.YChild{"Password", nil})
    for i := range passwords.Password {
        passwords.EntityData.Children.Append(types.GetSegmentPath(passwords.Password[i]), types.YChild{"Password", passwords.Password[i]})
    }
    passwords.EntityData.Leafs = types.NewOrderedMap()

    passwords.EntityData.YListKeys = []string {}

    return &(passwords.EntityData)
}

// MplsLdp_MplsLdpConfig_Passwords_Password
// This holds the MPLS LDP password configuration for use
// with a single LDP neighbor or group of LDP neighbors.
type MplsLdp_MplsLdpConfig_Passwords_Password struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    NbrVrf interface{}

    // This attribute is a key. This leaf holds the neighbor id for this password.
    // This id may be an lsr-id, an ip-address, or a filter describing a group of
    // neighbors. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.,
    // or string.
    NbrId interface{}

    // This attribute is a key. This is a user-assigned unique number identifying
    // a password for this neighbor or group of neighbors. Multiple passwords may
    // be assigned to a neighbor. If that is the case, each password is tried
    // starting with the lowest number to the highest until a passsword matches or
    // the list is exhausted. The type is interface{} with range: 0..4294967295.
    PasswordNum interface{}

    // This leaf is set true if the password is required and false if the password
    // is not required. The type is bool.
    PassRequired interface{}

    // This is a clear-text (non-encrypted password to be used with the neighbor.
    // The type is string.
    ClearPass interface{}

    // This is an encrypted password to be used with the neighbor. The type is
    // string.
    EncryptPass interface{}

    // This is a keychain identifier, which identifies an separately configured
    // keychain to be used with the neighbor neighbor. The type is string.
    KeychainPass interface{}
}

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetEntityData() *types.CommonEntityData {
    password.EntityData.YFilter = password.YFilter
    password.EntityData.YangName = "password"
    password.EntityData.BundleName = "cisco_ios_xe"
    password.EntityData.ParentYangName = "passwords"
    password.EntityData.SegmentPath = "password" + types.AddKeyToken(password.NbrVrf, "nbr-vrf") + types.AddKeyToken(password.NbrId, "nbr-id") + types.AddKeyToken(password.PasswordNum, "password-num")
    password.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    password.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    password.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    password.EntityData.Children = types.NewOrderedMap()
    password.EntityData.Leafs = types.NewOrderedMap()
    password.EntityData.Leafs.Append("nbr-vrf", types.YLeaf{"NbrVrf", password.NbrVrf})
    password.EntityData.Leafs.Append("nbr-id", types.YLeaf{"NbrId", password.NbrId})
    password.EntityData.Leafs.Append("password-num", types.YLeaf{"PasswordNum", password.PasswordNum})
    password.EntityData.Leafs.Append("pass-required", types.YLeaf{"PassRequired", password.PassRequired})
    password.EntityData.Leafs.Append("clear-pass", types.YLeaf{"ClearPass", password.ClearPass})
    password.EntityData.Leafs.Append("encrypt-pass", types.YLeaf{"EncryptPass", password.EncryptPass})
    password.EntityData.Leafs.Append("keychain-pass", types.YLeaf{"KeychainPass", password.KeychainPass})

    password.EntityData.YListKeys = []string {"NbrVrf", "NbrId", "PasswordNum"}

    return &(password.EntityData)
}

// MplsLdp_MplsLdpConfig_Session
// Configure session parameters
type MplsLdp_MplsLdpConfig_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial session backoff time (seconds). The type is interface{} with range:
    // 0..4294967295. The default value is 15.
    Backoff interface{}

    // Session holdtime in seconds. The type is interface{} with range: 0..65535.
    Seconds interface{}

    // Ignore LDP session holdtime. The type is interface{}.
    Infinite interface{}
}

func (session *MplsLdp_MplsLdpConfig_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xe"
    session.EntityData.ParentYangName = "mpls-ldp-config"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("backoff", types.YLeaf{"Backoff", session.Backoff})
    session.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", session.Seconds})
    session.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", session.Infinite})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// MplsLdp_MplsLdpConfig_LabelCfg
// This container holds the label allocation and
// advertisement configuration for the LDP Label Information
// Base. These control what prefixes may be allocated and
// advertised to peers.
type MplsLdp_MplsLdpConfig_LabelCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This is an allocation filter and advertisement filters for LDP labels in
    // this address family. The type is slice of
    // MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg.
    LabelAfCfg []*MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetEntityData() *types.CommonEntityData {
    labelCfg.EntityData.YFilter = labelCfg.YFilter
    labelCfg.EntityData.YangName = "label-cfg"
    labelCfg.EntityData.BundleName = "cisco_ios_xe"
    labelCfg.EntityData.ParentYangName = "mpls-ldp-config"
    labelCfg.EntityData.SegmentPath = "label-cfg"
    labelCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelCfg.EntityData.Children = types.NewOrderedMap()
    labelCfg.EntityData.Children.Append("label-af-cfg", types.YChild{"LabelAfCfg", nil})
    for i := range labelCfg.LabelAfCfg {
        labelCfg.EntityData.Children.Append(types.GetSegmentPath(labelCfg.LabelAfCfg[i]), types.YChild{"LabelAfCfg", labelCfg.LabelAfCfg[i]})
    }
    labelCfg.EntityData.Leafs = types.NewOrderedMap()

    labelCfg.EntityData.YListKeys = []string {}

    return &(labelCfg.EntityData)
}

// MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg
// This is an allocation filter and advertisement filters
// for LDP labels in this address family.
type MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // This contains the filter name for this label's prefix. The filter type is
    // device specific and could be an ACL, a prefix list, or other mechanism. The
    // type is string with length: 0..64.
    PrefixFilter interface{}

    // True if this LSR should allocate host-routes only. The type is bool.
    HostRouteEnable interface{}

    // MPLS LDP Label advertisement filter restrictions. The type is slice of
    // MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter.
    AdvtFilter []*MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetEntityData() *types.CommonEntityData {
    labelAfCfg.EntityData.YFilter = labelAfCfg.YFilter
    labelAfCfg.EntityData.YangName = "label-af-cfg"
    labelAfCfg.EntityData.BundleName = "cisco_ios_xe"
    labelAfCfg.EntityData.ParentYangName = "label-cfg"
    labelAfCfg.EntityData.SegmentPath = "label-af-cfg" + types.AddKeyToken(labelAfCfg.VrfName, "vrf-name") + types.AddKeyToken(labelAfCfg.AfName, "af-name")
    labelAfCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    labelAfCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    labelAfCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    labelAfCfg.EntityData.Children = types.NewOrderedMap()
    labelAfCfg.EntityData.Children.Append("advt-filter", types.YChild{"AdvtFilter", nil})
    for i := range labelAfCfg.AdvtFilter {
        labelAfCfg.EntityData.Children.Append(types.GetSegmentPath(labelAfCfg.AdvtFilter[i]), types.YChild{"AdvtFilter", labelAfCfg.AdvtFilter[i]})
    }
    labelAfCfg.EntityData.Leafs = types.NewOrderedMap()
    labelAfCfg.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", labelAfCfg.VrfName})
    labelAfCfg.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", labelAfCfg.AfName})
    labelAfCfg.EntityData.Leafs.Append("prefix-filter", types.YLeaf{"PrefixFilter", labelAfCfg.PrefixFilter})
    labelAfCfg.EntityData.Leafs.Append("host-route-enable", types.YLeaf{"HostRouteEnable", labelAfCfg.HostRouteEnable})

    labelAfCfg.EntityData.YListKeys = []string {"VrfName", "AfName"}

    return &(labelAfCfg.EntityData)
}

// MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter
// MPLS LDP Label advertisement filter restrictions.
type MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the filter name for this label's
    // prefix.  The filter type is device specific and could be an ACL, a prefix
    // list, or other mechanism. The type is string with length: 0..64.
    PrefixFilter interface{}

    // This attribute is a key. This contains the filter name for this label's
    // Peer. The filter type is device specific and could be an ACL, a prefix
    // list, or other mechanism. The type is string with length: 0..64.
    PeerFilter interface{}

    // This attribute is a key. This is an optional interface that may be used to
    // restrict the scope of the label advertisement. The type is string. Refers
    // to ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // This leaf controls what type of label is advertised for matching prefixes
    // to the matching peers. The type is AdvLabelType.
    AdvLabelCfg interface{}
}

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetEntityData() *types.CommonEntityData {
    advtFilter.EntityData.YFilter = advtFilter.YFilter
    advtFilter.EntityData.YangName = "advt-filter"
    advtFilter.EntityData.BundleName = "cisco_ios_xe"
    advtFilter.EntityData.ParentYangName = "label-af-cfg"
    advtFilter.EntityData.SegmentPath = "advt-filter" + types.AddKeyToken(advtFilter.PrefixFilter, "prefix-filter") + types.AddKeyToken(advtFilter.PeerFilter, "peer-filter") + types.AddKeyToken(advtFilter.Interface, "interface")
    advtFilter.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    advtFilter.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    advtFilter.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    advtFilter.EntityData.Children = types.NewOrderedMap()
    advtFilter.EntityData.Leafs = types.NewOrderedMap()
    advtFilter.EntityData.Leafs.Append("prefix-filter", types.YLeaf{"PrefixFilter", advtFilter.PrefixFilter})
    advtFilter.EntityData.Leafs.Append("peer-filter", types.YLeaf{"PeerFilter", advtFilter.PeerFilter})
    advtFilter.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", advtFilter.Interface})
    advtFilter.EntityData.Leafs.Append("adv-label-cfg", types.YLeaf{"AdvLabelCfg", advtFilter.AdvLabelCfg})

    advtFilter.EntityData.YListKeys = []string {"PrefixFilter", "PeerFilter", "Interface"}

    return &(advtFilter.EntityData)
}

// MplsLdp_MplsLdpConfig_Discovery
// LDP discovery
type MplsLdp_MplsLdpConfig_Discovery struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set this leaf to true to disable transmit and receive processing for
    // Type-Length-Value (TLV) in the discovery messages. The type is bool.
    InstanceTlv interface{}

    // This container holds the parameters for the non-targeted link hello.
    LinkHello MplsLdp_MplsLdpConfig_Discovery_LinkHello

    // This container holds the parameters for the targeted link hello.
    TargetedHello MplsLdp_MplsLdpConfig_Discovery_TargetedHello

    // This list contains the per-interface transport addresses, which overide the
    // global and default values.
    IntTransAddrs MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs
}

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetEntityData() *types.CommonEntityData {
    discovery.EntityData.YFilter = discovery.YFilter
    discovery.EntityData.YangName = "discovery"
    discovery.EntityData.BundleName = "cisco_ios_xe"
    discovery.EntityData.ParentYangName = "mpls-ldp-config"
    discovery.EntityData.SegmentPath = "discovery"
    discovery.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    discovery.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    discovery.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    discovery.EntityData.Children = types.NewOrderedMap()
    discovery.EntityData.Children.Append("link-hello", types.YChild{"LinkHello", &discovery.LinkHello})
    discovery.EntityData.Children.Append("targeted-hello", types.YChild{"TargetedHello", &discovery.TargetedHello})
    discovery.EntityData.Children.Append("int-trans-addrs", types.YChild{"IntTransAddrs", &discovery.IntTransAddrs})
    discovery.EntityData.Leafs = types.NewOrderedMap()
    discovery.EntityData.Leafs.Append("instance-tlv", types.YLeaf{"InstanceTlv", discovery.InstanceTlv})

    discovery.EntityData.YListKeys = []string {}

    return &(discovery.EntityData)
}

// MplsLdp_MplsLdpConfig_Discovery_LinkHello
// This container holds the parameters for the non-targeted
// link hello.
type MplsLdp_MplsLdpConfig_Discovery_LinkHello struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP discovery link hello holdtime in seconds. The type is interface{} with
    // range: 0..65535.
    Holdtime interface{}

    // LDP discovery link hello interval in seconds. The type is interface{} with
    // range: 0..65535.
    Interval interface{}
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetEntityData() *types.CommonEntityData {
    linkHello.EntityData.YFilter = linkHello.YFilter
    linkHello.EntityData.YangName = "link-hello"
    linkHello.EntityData.BundleName = "cisco_ios_xe"
    linkHello.EntityData.ParentYangName = "discovery"
    linkHello.EntityData.SegmentPath = "link-hello"
    linkHello.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    linkHello.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    linkHello.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    linkHello.EntityData.Children = types.NewOrderedMap()
    linkHello.EntityData.Leafs = types.NewOrderedMap()
    linkHello.EntityData.Leafs.Append("holdtime", types.YLeaf{"Holdtime", linkHello.Holdtime})
    linkHello.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", linkHello.Interval})

    linkHello.EntityData.YListKeys = []string {}

    return &(linkHello.EntityData)
}

// MplsLdp_MplsLdpConfig_Discovery_TargetedHello
// This container holds the parameters for the targeted
// link hello.
type MplsLdp_MplsLdpConfig_Discovery_TargetedHello struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LDP discovery targeted hello holdtime in seconds. The type is interface{}
    // with range: 0..65535.
    Holdtime interface{}

    // LDP discovery targeted hello interval in seconds. The type is interface{}
    // with range: 0..65535.
    Interval interface{}

    // Set to true if targeted hello messages may be accepted. The type is bool.
    Enable interface{}

    // Enables router to respond to requests for targeted hello messages.
    Accept MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept
}

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetEntityData() *types.CommonEntityData {
    targetedHello.EntityData.YFilter = targetedHello.YFilter
    targetedHello.EntityData.YangName = "targeted-hello"
    targetedHello.EntityData.BundleName = "cisco_ios_xe"
    targetedHello.EntityData.ParentYangName = "discovery"
    targetedHello.EntityData.SegmentPath = "targeted-hello"
    targetedHello.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    targetedHello.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    targetedHello.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    targetedHello.EntityData.Children = types.NewOrderedMap()
    targetedHello.EntityData.Children.Append("accept", types.YChild{"Accept", &targetedHello.Accept})
    targetedHello.EntityData.Leafs = types.NewOrderedMap()
    targetedHello.EntityData.Leafs.Append("holdtime", types.YLeaf{"Holdtime", targetedHello.Holdtime})
    targetedHello.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", targetedHello.Interval})
    targetedHello.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", targetedHello.Enable})

    targetedHello.EntityData.YListKeys = []string {}

    return &(targetedHello.EntityData)
}

// MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept
// Enables router to respond to requests for targeted
// hello messages
type MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set to true if targeted hello messages may be accepted. The type is bool.
    Enable interface{}

    // Only respond to requests for targeted hello messages from sources matching
    // this filter. The filter type is device specific and could be an ACL, a
    // prefix list, or other mechanism. The type is string.
    SrcFilter interface{}
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetEntityData() *types.CommonEntityData {
    accept.EntityData.YFilter = accept.YFilter
    accept.EntityData.YangName = "accept"
    accept.EntityData.BundleName = "cisco_ios_xe"
    accept.EntityData.ParentYangName = "targeted-hello"
    accept.EntityData.SegmentPath = "accept"
    accept.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    accept.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    accept.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    accept.EntityData.Children = types.NewOrderedMap()
    accept.EntityData.Leafs = types.NewOrderedMap()
    accept.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", accept.Enable})
    accept.EntityData.Leafs.Append("src-filter", types.YLeaf{"SrcFilter", accept.SrcFilter})

    accept.EntityData.YListKeys = []string {}

    return &(accept.EntityData)
}

// MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs
// This list contains the per-interface transport
// addresses, which overide the global and default
// values.
type MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry contains the per-interface transport addresses, which overide
    // the global and default values. The type is slice of
    // MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr.
    IntTransAddr []*MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetEntityData() *types.CommonEntityData {
    intTransAddrs.EntityData.YFilter = intTransAddrs.YFilter
    intTransAddrs.EntityData.YangName = "int-trans-addrs"
    intTransAddrs.EntityData.BundleName = "cisco_ios_xe"
    intTransAddrs.EntityData.ParentYangName = "discovery"
    intTransAddrs.EntityData.SegmentPath = "int-trans-addrs"
    intTransAddrs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intTransAddrs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intTransAddrs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intTransAddrs.EntityData.Children = types.NewOrderedMap()
    intTransAddrs.EntityData.Children.Append("int-trans-addr", types.YChild{"IntTransAddr", nil})
    for i := range intTransAddrs.IntTransAddr {
        intTransAddrs.EntityData.Children.Append(types.GetSegmentPath(intTransAddrs.IntTransAddr[i]), types.YChild{"IntTransAddr", intTransAddrs.IntTransAddr[i]})
    }
    intTransAddrs.EntityData.Leafs = types.NewOrderedMap()

    intTransAddrs.EntityData.YListKeys = []string {}

    return &(intTransAddrs.EntityData)
}

// MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr
// This entry contains the per-interface transport
// addresses, which overide the global and default
// values.
type MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // This attribute is a key. The Interface Name. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    IntName interface{}

    // Advertise this address as the address in LDP discovery hello messages and
    // use it for LDP transport. The type is one of the following types: string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    TransIp interface{}

    // Advertise this interface's address as the address in LDP discovery hello
    // messages and use it for LDP transport. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    TransInt interface{}
}

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetEntityData() *types.CommonEntityData {
    intTransAddr.EntityData.YFilter = intTransAddr.YFilter
    intTransAddr.EntityData.YangName = "int-trans-addr"
    intTransAddr.EntityData.BundleName = "cisco_ios_xe"
    intTransAddr.EntityData.ParentYangName = "int-trans-addrs"
    intTransAddr.EntityData.SegmentPath = "int-trans-addr" + types.AddKeyToken(intTransAddr.AfName, "af-name") + types.AddKeyToken(intTransAddr.IntName, "int-name")
    intTransAddr.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    intTransAddr.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    intTransAddr.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    intTransAddr.EntityData.Children = types.NewOrderedMap()
    intTransAddr.EntityData.Leafs = types.NewOrderedMap()
    intTransAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", intTransAddr.AfName})
    intTransAddr.EntityData.Leafs.Append("int-name", types.YLeaf{"IntName", intTransAddr.IntName})
    intTransAddr.EntityData.Leafs.Append("trans-ip", types.YLeaf{"TransIp", intTransAddr.TransIp})
    intTransAddr.EntityData.Leafs.Append("trans-int", types.YLeaf{"TransInt", intTransAddr.TransInt})

    intTransAddr.EntityData.YListKeys = []string {"AfName", "IntName"}

    return &(intTransAddr.EntityData)
}

// MplsLdp_MplsLdpConfig_GracefulRestart
// Configure LDP Graceful Restart
type MplsLdp_MplsLdpConfig_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable graceful restartable. The type is bool.
    IsGracefulRestartable interface{}

    // Specifies the amount of time the MPLS LDP forwarding state must be
    // preserved after the control plane restarts. The type is interface{} with
    // range: 5..300. Units are seconds.
    ForwardingHolding interface{}

    // Amount of time (in seconds) that the router should hold stale label-FEC
    // bindings after an LDP session has been reestablished. The type is
    // interface{} with range: 5..300. Units are seconds.
    MaxRecovery interface{}

    // Amount of time (in seconds) that the router must wait for an LDP session to
    // be reestablished. The type is interface{} with range: 5..300. Units are
    // seconds.
    NbrLiveness interface{}

    // This contains the filter name for peers for which this LSR will act as a
    // graceful-restart helper. The type is slice of
    // MplsLdp_MplsLdpConfig_GracefulRestart_Helper.
    Helper []*MplsLdp_MplsLdpConfig_GracefulRestart_Helper
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xe"
    gracefulRestart.EntityData.ParentYangName = "mpls-ldp-config"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Children.Append("helper", types.YChild{"Helper", nil})
    for i := range gracefulRestart.Helper {
        gracefulRestart.EntityData.Children.Append(types.GetSegmentPath(gracefulRestart.Helper[i]), types.YChild{"Helper", gracefulRestart.Helper[i]})
    }
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs.Append("is-graceful-restartable", types.YLeaf{"IsGracefulRestartable", gracefulRestart.IsGracefulRestartable})
    gracefulRestart.EntityData.Leafs.Append("forwarding-holding", types.YLeaf{"ForwardingHolding", gracefulRestart.ForwardingHolding})
    gracefulRestart.EntityData.Leafs.Append("max-recovery", types.YLeaf{"MaxRecovery", gracefulRestart.MaxRecovery})
    gracefulRestart.EntityData.Leafs.Append("nbr-liveness", types.YLeaf{"NbrLiveness", gracefulRestart.NbrLiveness})

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// MplsLdp_MplsLdpConfig_GracefulRestart_Helper
// This contains the filter name for peers for which this
// LSR will act as a graceful-restart helper.
type MplsLdp_MplsLdpConfig_GracefulRestart_Helper struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    HelperVrf interface{}

    // This attribute is a key. This contains the filter name for peers for which
    // this LSR will act as a graceful-restart helper. The filter type is device
    // specific and could be an ACL, a prefix list, or other mechanism. The type
    // is string with length: 0..64.
    HelperFilter interface{}
}

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetEntityData() *types.CommonEntityData {
    helper.EntityData.YFilter = helper.YFilter
    helper.EntityData.YangName = "helper"
    helper.EntityData.BundleName = "cisco_ios_xe"
    helper.EntityData.ParentYangName = "graceful-restart"
    helper.EntityData.SegmentPath = "helper" + types.AddKeyToken(helper.HelperVrf, "helper-vrf") + types.AddKeyToken(helper.HelperFilter, "helper-filter")
    helper.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    helper.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    helper.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    helper.EntityData.Children = types.NewOrderedMap()
    helper.EntityData.Leafs = types.NewOrderedMap()
    helper.EntityData.Leafs.Append("helper-vrf", types.YLeaf{"HelperVrf", helper.HelperVrf})
    helper.EntityData.Leafs.Append("helper-filter", types.YLeaf{"HelperFilter", helper.HelperFilter})

    helper.EntityData.YListKeys = []string {"HelperVrf", "HelperFilter"}

    return &(helper.EntityData)
}

// MplsLdp_MplsLdpConfig_Logging
// Enable LDP logging
type MplsLdp_MplsLdpConfig_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable logging of graceful-restart messages. The type is bool.
    GracefulRestart interface{}

    // Enable logging of neighbor messages. The type is bool.
    Neighbor interface{}

    // Enable logging of nsr messages. The type is bool.
    Nsr interface{}

    // Enable logging of adjacency messages. The type is bool.
    Adjacency interface{}

    // Enable logging of session-protection messages. The type is bool.
    SessionProtection interface{}

    // Enable logging of password messages.
    Password MplsLdp_MplsLdpConfig_Logging_Password
}

func (logging *MplsLdp_MplsLdpConfig_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xe"
    logging.EntityData.ParentYangName = "mpls-ldp-config"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Children.Append("password", types.YChild{"Password", &logging.Password})
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("graceful-restart", types.YLeaf{"GracefulRestart", logging.GracefulRestart})
    logging.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", logging.Neighbor})
    logging.EntityData.Leafs.Append("nsr", types.YLeaf{"Nsr", logging.Nsr})
    logging.EntityData.Leafs.Append("adjacency", types.YLeaf{"Adjacency", logging.Adjacency})
    logging.EntityData.Leafs.Append("session-protection", types.YLeaf{"SessionProtection", logging.SessionProtection})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// MplsLdp_MplsLdpConfig_Logging_Password
// Enable logging of password messages.
type MplsLdp_MplsLdpConfig_Logging_Password struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log MPLS LDP password configuration changes.
    ConfigMsg MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg

    // Log MPLS LDP password rollover messages.
    RolloverMsg MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetEntityData() *types.CommonEntityData {
    password.EntityData.YFilter = password.YFilter
    password.EntityData.YangName = "password"
    password.EntityData.BundleName = "cisco_ios_xe"
    password.EntityData.ParentYangName = "logging"
    password.EntityData.SegmentPath = "password"
    password.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    password.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    password.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    password.EntityData.Children = types.NewOrderedMap()
    password.EntityData.Children.Append("config-msg", types.YChild{"ConfigMsg", &password.ConfigMsg})
    password.EntityData.Children.Append("rollover-msg", types.YChild{"RolloverMsg", &password.RolloverMsg})
    password.EntityData.Leafs = types.NewOrderedMap()

    password.EntityData.YListKeys = []string {}

    return &(password.EntityData)
}

// MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg
// Log MPLS LDP password configuration changes.
type MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log MPLS LDP password configuration changes. The type is bool.
    Enable interface{}

    // This is the number of messages per minute to limit the logging. A value of
    // 0 indicates no limits on the number of logged messages. The type is
    // interface{} with range: 0..4294967295.
    RateLimit interface{}
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetEntityData() *types.CommonEntityData {
    configMsg.EntityData.YFilter = configMsg.YFilter
    configMsg.EntityData.YangName = "config-msg"
    configMsg.EntityData.BundleName = "cisco_ios_xe"
    configMsg.EntityData.ParentYangName = "password"
    configMsg.EntityData.SegmentPath = "config-msg"
    configMsg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    configMsg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    configMsg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    configMsg.EntityData.Children = types.NewOrderedMap()
    configMsg.EntityData.Leafs = types.NewOrderedMap()
    configMsg.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", configMsg.Enable})
    configMsg.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", configMsg.RateLimit})

    configMsg.EntityData.YListKeys = []string {}

    return &(configMsg.EntityData)
}

// MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg
// Log MPLS LDP password rollover messages.
type MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Log MPLS LDP password rollover messages. The type is bool.
    Enable interface{}

    // This is the number of messages per minute to limit the logging. A value of
    // 0 indicates no limits on the number of logged messages. The type is
    // interface{} with range: 0..4294967295.
    RateLimit interface{}
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetEntityData() *types.CommonEntityData {
    rolloverMsg.EntityData.YFilter = rolloverMsg.YFilter
    rolloverMsg.EntityData.YangName = "rollover-msg"
    rolloverMsg.EntityData.BundleName = "cisco_ios_xe"
    rolloverMsg.EntityData.ParentYangName = "password"
    rolloverMsg.EntityData.SegmentPath = "rollover-msg"
    rolloverMsg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    rolloverMsg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    rolloverMsg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    rolloverMsg.EntityData.Children = types.NewOrderedMap()
    rolloverMsg.EntityData.Leafs = types.NewOrderedMap()
    rolloverMsg.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rolloverMsg.Enable})
    rolloverMsg.EntityData.Leafs.Append("rate-limit", types.YLeaf{"RateLimit", rolloverMsg.RateLimit})

    rolloverMsg.EntityData.YListKeys = []string {}

    return &(rolloverMsg.EntityData)
}

// MplsLdp_MplsLdpConfig_Interfaces
// MPLS LDP Interface configuration commands.
type MplsLdp_MplsLdpConfig_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP Interface configuration commands. Where a corresponding global
    // configuration command exists, the interface level command will take
    // precedence when configured. The type is slice of
    // MplsLdp_MplsLdpConfig_Interfaces_Interface.
    Interface []*MplsLdp_MplsLdpConfig_Interfaces_Interface
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xe"
    interfaces.EntityData.ParentYangName = "mpls-ldp-config"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// MplsLdp_MplsLdpConfig_Interfaces_Interface
// MPLS LDP Interface configuration commands. Where a
// corresponding global configuration command exists, the
// interface level command will take precedence when
// configured.
type MplsLdp_MplsLdpConfig_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    Vrf interface{}

    // This attribute is a key. The Interface Name. The type is string. Refers to
    // ietf_interfaces.Interfaces_Interface_Name
    Interface interface{}

    // LDP discovery link hello interval in seconds for this interface. This value
    // overides the global setting. The type is interface{} with range:
    // 0..4294967295. Units are second. The default value is 5.
    LinkHelloInt interface{}

    // LDP discovery link hello holdtime in seconds for this interface. This value
    // overides the global setting. The type is interface{} with range:
    // 0..4294967295. Units are second. The default value is 15.
    LinkHelloHold interface{}

    // When set to true, disable LDP discovery's quick start mode for this
    // interface. The type is bool.
    DisableQuickStartInt interface{}

    // Time in seconds to delay IGP sync after session comes up. The type is
    // interface{} with range: 5..300. Units are second.
    Seconds interface{}

    // This choice causes IGP sync up immediately upon session up. The type is
    // interface{}.
    DisableDelay interface{}

    // Address Family specific operational data.
    Afs MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs
}

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xe"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.Vrf, "vrf") + types.AddKeyToken(self.Interface, "interface")
    self.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("afs", types.YChild{"Afs", &self.Afs})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", self.Vrf})
    self.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", self.Interface})
    self.EntityData.Leafs.Append("link-hello-int", types.YLeaf{"LinkHelloInt", self.LinkHelloInt})
    self.EntityData.Leafs.Append("link-hello-hold", types.YLeaf{"LinkHelloHold", self.LinkHelloHold})
    self.EntityData.Leafs.Append("disable-quick-start-int", types.YLeaf{"DisableQuickStartInt", self.DisableQuickStartInt})
    self.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", self.Seconds})
    self.EntityData.Leafs.Append("disable-delay", types.YLeaf{"DisableDelay", self.DisableDelay})

    self.EntityData.YListKeys = []string {"Vrf", "Interface"}

    return &(self.EntityData)
}

// MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs
// Address Family specific operational data
type MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MPLS LDP Operational data for this Address Family. The type is slice of
    // MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af.
    Af []*MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetEntityData() *types.CommonEntityData {
    afs.EntityData.YFilter = afs.YFilter
    afs.EntityData.YangName = "afs"
    afs.EntityData.BundleName = "cisco_ios_xe"
    afs.EntityData.ParentYangName = "interface"
    afs.EntityData.SegmentPath = "afs"
    afs.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    afs.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    afs.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    afs.EntityData.Children = types.NewOrderedMap()
    afs.EntityData.Children.Append("af", types.YChild{"Af", nil})
    for i := range afs.Af {
        afs.EntityData.Children.Append(types.GetSegmentPath(afs.Af[i]), types.YChild{"Af", afs.Af[i]})
    }
    afs.EntityData.Leafs = types.NewOrderedMap()

    afs.EntityData.YListKeys = []string {}

    return &(afs.EntityData)
}

// MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af
// MPLS LDP Operational data for this Address Family.
type MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // This is set true to enable LDP on this interface. The type is bool.
    Enable interface{}

    // True if LDP autoconfig is explicitly disabled on this interface. The type
    // is bool.
    AutoconfigDisable interface{}

    // MPLS LDP configuration for protocol redistribution. By default,
    // redistribution of BGP routes is disabled. It can be enabled for all BGP
    // routes or for a specific AS. Also it can be redistributed to all LDP peers
    // or to a filtered group of peers.
    BgpRedist MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist
}

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xe"
    af.EntityData.ParentYangName = "afs"
    af.EntityData.SegmentPath = "af" + types.AddKeyToken(af.AfName, "af-name")
    af.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Children.Append("bgp-redist", types.YChild{"BgpRedist", &af.BgpRedist})
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", af.Enable})
    af.EntityData.Leafs.Append("autoconfig-disable", types.YLeaf{"AutoconfigDisable", af.AutoconfigDisable})

    af.EntityData.YListKeys = []string {"AfName"}

    return &(af.EntityData)
}

// MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist
// MPLS LDP configuration for protocol
// redistribution. By default, redistribution of BGP
// routes is disabled. It can be enabled for all
// BGP routes or for a specific AS. Also it can be
// redistributed to all LDP peers or to a filtered
// group of peers.
type MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // First half of BGP AS number in XX.YY format.  Mandatory Must be a non-zero
    // value if second half is zero. The type is interface{} with range: 0..65535.
    AsXx interface{}

    // Second half of BGP AS number in XX.YY format. Mandatory Must be a non-zero
    // value if first half is zero. The type is interface{} with range:
    // 0..4294967295.
    AsYy interface{}

    // Filter of neighbors to receive BGP route redistributions from LDP. If the
    // list is empty or unset, all LDP neighbors will receive redistributions. The
    // type is string.
    AdvertiseTo interface{}

    // This is set true to allow LDP to redistribute BGP routes. The type is bool.
    Enable interface{}
}

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetEntityData() *types.CommonEntityData {
    bgpRedist.EntityData.YFilter = bgpRedist.YFilter
    bgpRedist.EntityData.YangName = "bgp-redist"
    bgpRedist.EntityData.BundleName = "cisco_ios_xe"
    bgpRedist.EntityData.ParentYangName = "af"
    bgpRedist.EntityData.SegmentPath = "bgp-redist"
    bgpRedist.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    bgpRedist.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    bgpRedist.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    bgpRedist.EntityData.Children = types.NewOrderedMap()
    bgpRedist.EntityData.Leafs = types.NewOrderedMap()
    bgpRedist.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", bgpRedist.AsXx})
    bgpRedist.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", bgpRedist.AsYy})
    bgpRedist.EntityData.Leafs.Append("advertise-to", types.YLeaf{"AdvertiseTo", bgpRedist.AdvertiseTo})
    bgpRedist.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bgpRedist.Enable})

    bgpRedist.EntityData.YListKeys = []string {}

    return &(bgpRedist.EntityData)
}

// MplsLdp_MplsLdpConfig_Routing
// This containter provides the MPLS LDP config for routing
// protocols from which it can obtain addresses to
// associate with labels.
type MplsLdp_MplsLdpConfig_Routing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This entry provides the MPLS LDP config for this routing instance. The type
    // is slice of MplsLdp_MplsLdpConfig_Routing_RoutingInst.
    RoutingInst []*MplsLdp_MplsLdpConfig_Routing_RoutingInst
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetEntityData() *types.CommonEntityData {
    routing.EntityData.YFilter = routing.YFilter
    routing.EntityData.YangName = "routing"
    routing.EntityData.BundleName = "cisco_ios_xe"
    routing.EntityData.ParentYangName = "mpls-ldp-config"
    routing.EntityData.SegmentPath = "routing"
    routing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routing.EntityData.Children = types.NewOrderedMap()
    routing.EntityData.Children.Append("routing-inst", types.YChild{"RoutingInst", nil})
    for i := range routing.RoutingInst {
        routing.EntityData.Children.Append(types.GetSegmentPath(routing.RoutingInst[i]), types.YChild{"RoutingInst", routing.RoutingInst[i]})
    }
    routing.EntityData.Leafs = types.NewOrderedMap()

    routing.EntityData.YListKeys = []string {}

    return &(routing.EntityData)
}

// MplsLdp_MplsLdpConfig_Routing_RoutingInst
// This entry provides the MPLS LDP config for this
// routing instance.
type MplsLdp_MplsLdpConfig_Routing_RoutingInst struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the routing instance for which this MPLS
    // LDP configuration applies. The type is string.
    RoutingInstName interface{}

    // This leaf enables or disables LDP for all interfaces covered by this
    // routing instance subject to the autoconfig-scope. The type is bool.
    AutoconfigEnable interface{}

    // This leaf restricts the LDP Autoconfiguration feature to enable LDP on
    // interfaces belonging to an OSPF process for a specific area. If no area is
    // specified, then this applies to all interfaces associated with the. If an
    // area ID is specified, then only interfaces associated with that OSPF area
    // are automatically enabled with LDP. Any interface-specific ldp
    // configuration will overide this setting for that interface. The type is
    // interface{} with range: 0..4294967295.
    AreaId interface{}

    // This leaf restricts the LDP Autoconfiguration feature to enable LDP on
    // interfaces belonging to an ISIS process for a specific level. If no level
    // is specified, then this applies to all interfaces associated with the. If a
    // level is specified, then only interfaces associated with that ISIS level
    // are automatically enabled with LDP. Any interface-specific ldp
    // configuration will overide this setting for that interface. The type is
    // LevelId.
    LevelId interface{}

    // When set to true this enables LDP IGP synchronization. Without
    // syncrhonization, packet loss can occur because the actions of the IGP and
    // LDP are not synchronized. The type is bool.
    Sync interface{}
}

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetEntityData() *types.CommonEntityData {
    routingInst.EntityData.YFilter = routingInst.YFilter
    routingInst.EntityData.YangName = "routing-inst"
    routingInst.EntityData.BundleName = "cisco_ios_xe"
    routingInst.EntityData.ParentYangName = "routing"
    routingInst.EntityData.SegmentPath = "routing-inst" + types.AddKeyToken(routingInst.RoutingInstName, "routing-inst-name")
    routingInst.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    routingInst.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    routingInst.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    routingInst.EntityData.Children = types.NewOrderedMap()
    routingInst.EntityData.Leafs = types.NewOrderedMap()
    routingInst.EntityData.Leafs.Append("routing-inst-name", types.YLeaf{"RoutingInstName", routingInst.RoutingInstName})
    routingInst.EntityData.Leafs.Append("autoconfig-enable", types.YLeaf{"AutoconfigEnable", routingInst.AutoconfigEnable})
    routingInst.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", routingInst.AreaId})
    routingInst.EntityData.Leafs.Append("level-id", types.YLeaf{"LevelId", routingInst.LevelId})
    routingInst.EntityData.Leafs.Append("sync", types.YLeaf{"Sync", routingInst.Sync})

    routingInst.EntityData.YListKeys = []string {"RoutingInstName"}

    return &(routingInst.EntityData)
}

// MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId represents overide this setting for that interface.
type MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId string

const (
    // This leaf restricts the LDP Autoconfiguration
    // feature to enable LDP on interfaces belonging
    // to an IS-IS process level 1.
    // Any interface-specific ldp configuration will
    // overide this setting for that interface.
    MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId_level_1 MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId = "level-1"

    // This leaf restricts the LDP Autoconfiguration
    // feature to enable LDP on interfaces belonging
    // to an IS-IS process level 1.
    // Any interface-specific ldp configuration will
    // overide this setting for that interface.
    MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId_level_2 MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId = "level-2"

    // This leaf restricts the LDP Autoconfiguration
    // feature to enable LDP on interfaces belonging
    // to an IS-IS process level 2.
    // Any interface-specific ldp configuration will
    // overide this setting for that interface.
    MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId_level_1_2 MplsLdp_MplsLdpConfig_Routing_RoutingInst_LevelId = "level-1-2"
)

// MplsLdp_MplsLdpConfig_DualStack
// This container holds the configuration of dual IPv4 and
// IPv6 stack peers.
type MplsLdp_MplsLdpConfig_DualStack struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait time in seconds (0 indicates no preference). The type is interface{}
    // with range: 0..60.
    MaxWait interface{}

    // This contains the filter name for peers where IPv4 connections are
    // preferred over IPv6 connections. The filter type is device specific and
    // could be an ACL, a prefix list, or other mechanism. The type is string.
    PreferIpv4Peers interface{}
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetEntityData() *types.CommonEntityData {
    dualStack.EntityData.YFilter = dualStack.YFilter
    dualStack.EntityData.YangName = "dual-stack"
    dualStack.EntityData.BundleName = "cisco_ios_xe"
    dualStack.EntityData.ParentYangName = "mpls-ldp-config"
    dualStack.EntityData.SegmentPath = "dual-stack"
    dualStack.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    dualStack.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    dualStack.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    dualStack.EntityData.Children = types.NewOrderedMap()
    dualStack.EntityData.Leafs = types.NewOrderedMap()
    dualStack.EntityData.Leafs.Append("max-wait", types.YLeaf{"MaxWait", dualStack.MaxWait})
    dualStack.EntityData.Leafs.Append("prefer-ipv4-peers", types.YLeaf{"PreferIpv4Peers", dualStack.PreferIpv4Peers})

    dualStack.EntityData.YListKeys = []string {}

    return &(dualStack.EntityData)
}

// ClearMsgCounters
// This RPC clears the LDP message counters for either a single
// neighbor or for all neighbors.
type ClearMsgCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearMsgCounters_Input

    
    Output ClearMsgCounters_Output
}

func (clearMsgCounters *ClearMsgCounters) GetEntityData() *types.CommonEntityData {
    clearMsgCounters.EntityData.YFilter = clearMsgCounters.YFilter
    clearMsgCounters.EntityData.YangName = "clear-msg-counters"
    clearMsgCounters.EntityData.BundleName = "cisco_ios_xe"
    clearMsgCounters.EntityData.ParentYangName = "Cisco-IOS-XE-mpls-ldp"
    clearMsgCounters.EntityData.SegmentPath = "Cisco-IOS-XE-mpls-ldp:clear-msg-counters"
    clearMsgCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clearMsgCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clearMsgCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clearMsgCounters.EntityData.Children = types.NewOrderedMap()
    clearMsgCounters.EntityData.Children.Append("input", types.YChild{"Input", &clearMsgCounters.Input})
    clearMsgCounters.EntityData.Children.Append("output", types.YChild{"Output", &clearMsgCounters.Output})
    clearMsgCounters.EntityData.Leafs = types.NewOrderedMap()

    clearMsgCounters.EntityData.YListKeys = []string {}

    return &(clearMsgCounters.EntityData)
}

// ClearMsgCounters_Input
type ClearMsgCounters_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This contains the VRF Name, where 'default' is used for the default vrf.
    // The type is string.
    VrfName interface{}

    // LSR ID of the neighbor. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NbrIp interface{}

    // Clear information for all neighbors. The type is interface{}.
    All interface{}
}

func (input *ClearMsgCounters_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "clear-msg-counters"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", input.VrfName})
    input.EntityData.Leafs.Append("nbr-ip", types.YLeaf{"NbrIp", input.NbrIp})
    input.EntityData.Leafs.Append("all", types.YLeaf{"All", input.All})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearMsgCounters_Output
type ClearMsgCounters_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Return status will be 'OK' on success or an explanation string on failure.
    // The type is string.
    Status interface{}
}

func (output *ClearMsgCounters_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "clear-msg-counters"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("status", types.YLeaf{"Status", output.Status})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// RestartNeighbor
// This RPC restarts a single LDP session or all LDP sessions,
// but does not restart the LDP process itself, if the device
// supports that capability.
type RestartNeighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input RestartNeighbor_Input

    
    Output RestartNeighbor_Output
}

func (restartNeighbor *RestartNeighbor) GetEntityData() *types.CommonEntityData {
    restartNeighbor.EntityData.YFilter = restartNeighbor.YFilter
    restartNeighbor.EntityData.YangName = "restart-neighbor"
    restartNeighbor.EntityData.BundleName = "cisco_ios_xe"
    restartNeighbor.EntityData.ParentYangName = "Cisco-IOS-XE-mpls-ldp"
    restartNeighbor.EntityData.SegmentPath = "Cisco-IOS-XE-mpls-ldp:restart-neighbor"
    restartNeighbor.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    restartNeighbor.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    restartNeighbor.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    restartNeighbor.EntityData.Children = types.NewOrderedMap()
    restartNeighbor.EntityData.Children.Append("input", types.YChild{"Input", &restartNeighbor.Input})
    restartNeighbor.EntityData.Children.Append("output", types.YChild{"Output", &restartNeighbor.Output})
    restartNeighbor.EntityData.Leafs = types.NewOrderedMap()

    restartNeighbor.EntityData.YListKeys = []string {}

    return &(restartNeighbor.EntityData)
}

// RestartNeighbor_Input
type RestartNeighbor_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This contains the VRF Name, where 'default' is used for the default vrf.
    // The type is string.
    VrfName interface{}

    // LSR ID of the neighbor. The type is one of the following types: string with
    // pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    NbrIp interface{}

    // Restart sessions for all neighbors. The type is interface{}.
    All interface{}
}

func (input *RestartNeighbor_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "restart-neighbor"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", input.VrfName})
    input.EntityData.Leafs.Append("nbr-ip", types.YLeaf{"NbrIp", input.NbrIp})
    input.EntityData.Leafs.Append("all", types.YLeaf{"All", input.All})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// RestartNeighbor_Output
type RestartNeighbor_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Return status will be 'OK' on success or an explanation string on failure.
    // The type is string.
    Status interface{}
}

func (output *RestartNeighbor_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "restart-neighbor"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("status", types.YLeaf{"Status", output.Status})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

// ClearForwarding
// This command resets LDP installed forwarding state for all
// prefixes or a given prefix. It is useful when installed 
// LDP forwarding state needs to be reprogrammed in LSD and
// MPLS forwarding.
type ClearForwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Input ClearForwarding_Input

    
    Output ClearForwarding_Output
}

func (clearForwarding *ClearForwarding) GetEntityData() *types.CommonEntityData {
    clearForwarding.EntityData.YFilter = clearForwarding.YFilter
    clearForwarding.EntityData.YangName = "clear-forwarding"
    clearForwarding.EntityData.BundleName = "cisco_ios_xe"
    clearForwarding.EntityData.ParentYangName = "Cisco-IOS-XE-mpls-ldp"
    clearForwarding.EntityData.SegmentPath = "Cisco-IOS-XE-mpls-ldp:clear-forwarding"
    clearForwarding.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clearForwarding.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clearForwarding.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clearForwarding.EntityData.Children = types.NewOrderedMap()
    clearForwarding.EntityData.Children.Append("input", types.YChild{"Input", &clearForwarding.Input})
    clearForwarding.EntityData.Children.Append("output", types.YChild{"Output", &clearForwarding.Output})
    clearForwarding.EntityData.Leafs = types.NewOrderedMap()

    clearForwarding.EntityData.YListKeys = []string {}

    return &(clearForwarding.EntityData)
}

// ClearForwarding_Input
type ClearForwarding_Input struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This contains the VRF Name, where 'default' is used for the default vrf.
    // The type is string.
    VrfName interface{}

    // This case provides the IP prefix for the forwarding entry whose data should
    // be cleared. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PrefixIp interface{}

    // This case is used to clear the forwarding entries for all prefixes. The
    // type is interface{}.
    All interface{}
}

func (input *ClearForwarding_Input) GetEntityData() *types.CommonEntityData {
    input.EntityData.YFilter = input.YFilter
    input.EntityData.YangName = "input"
    input.EntityData.BundleName = "cisco_ios_xe"
    input.EntityData.ParentYangName = "clear-forwarding"
    input.EntityData.SegmentPath = "input"
    input.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    input.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    input.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    input.EntityData.Children = types.NewOrderedMap()
    input.EntityData.Leafs = types.NewOrderedMap()
    input.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", input.VrfName})
    input.EntityData.Leafs.Append("prefix-ip", types.YLeaf{"PrefixIp", input.PrefixIp})
    input.EntityData.Leafs.Append("all", types.YLeaf{"All", input.All})

    input.EntityData.YListKeys = []string {}

    return &(input.EntityData)
}

// ClearForwarding_Output
type ClearForwarding_Output struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Return status will be 'OK' on success or an explanatory string on failure.
    // The type is string.
    Status interface{}
}

func (output *ClearForwarding_Output) GetEntityData() *types.CommonEntityData {
    output.EntityData.YFilter = output.YFilter
    output.EntityData.YangName = "output"
    output.EntityData.BundleName = "cisco_ios_xe"
    output.EntityData.ParentYangName = "clear-forwarding"
    output.EntityData.SegmentPath = "output"
    output.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    output.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    output.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    output.EntityData.Children = types.NewOrderedMap()
    output.EntityData.Leafs = types.NewOrderedMap()
    output.EntityData.Leafs.Append("status", types.YLeaf{"Status", output.Status})

    output.EntityData.YListKeys = []string {}

    return &(output.EntityData)
}

