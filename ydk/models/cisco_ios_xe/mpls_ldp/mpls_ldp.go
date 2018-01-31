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
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP operational data.
    MplsLdpState MplsLdp_MplsLdpState

    // MPLS LDP Configuration.
    MplsLdpConfig MplsLdp_MplsLdpConfig
}

func (mplsLdp *MplsLdp) GetFilter() yfilter.YFilter { return mplsLdp.YFilter }

func (mplsLdp *MplsLdp) SetFilter(yf yfilter.YFilter) { mplsLdp.YFilter = yf }

func (mplsLdp *MplsLdp) GetGoName(yname string) string {
    if yname == "mpls-ldp-state" { return "MplsLdpState" }
    if yname == "mpls-ldp-config" { return "MplsLdpConfig" }
    return ""
}

func (mplsLdp *MplsLdp) GetSegmentPath() string {
    return "Cisco-IOS-XE-mpls-ldp:mpls-ldp"
}

func (mplsLdp *MplsLdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpls-ldp-state" {
        return &mplsLdp.MplsLdpState
    }
    if childYangName == "mpls-ldp-config" {
        return &mplsLdp.MplsLdpConfig
    }
    return nil
}

func (mplsLdp *MplsLdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpls-ldp-state"] = &mplsLdp.MplsLdpState
    children["mpls-ldp-config"] = &mplsLdp.MplsLdpConfig
    return children
}

func (mplsLdp *MplsLdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsLdp *MplsLdp) GetBundleName() string { return "cisco_ios_xe" }

func (mplsLdp *MplsLdp) GetYangName() string { return "mpls-ldp" }

func (mplsLdp *MplsLdp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsLdp *MplsLdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsLdp *MplsLdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsLdp *MplsLdp) SetParent(parent types.Entity) { mplsLdp.parent = parent }

func (mplsLdp *MplsLdp) GetParent() types.Entity { return mplsLdp.parent }

func (mplsLdp *MplsLdp) GetParentYangName() string { return "Cisco-IOS-XE-mpls-ldp" }

// MplsLdp_MplsLdpState
// MPLS LDP operational data.
type MplsLdp_MplsLdpState struct {
    parent types.Entity
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

func (mplsLdpState *MplsLdp_MplsLdpState) GetFilter() yfilter.YFilter { return mplsLdpState.YFilter }

func (mplsLdpState *MplsLdp_MplsLdpState) SetFilter(yf yfilter.YFilter) { mplsLdpState.YFilter = yf }

func (mplsLdpState *MplsLdp_MplsLdpState) GetGoName(yname string) string {
    if yname == "oper-summary" { return "OperSummary" }
    if yname == "forwarding-summary" { return "ForwardingSummary" }
    if yname == "bindings-summary" { return "BindingsSummary" }
    if yname == "nsr-summary-all" { return "NsrSummaryAll" }
    if yname == "icpm-summary-all" { return "IcpmSummaryAll" }
    if yname == "parameters" { return "Parameters" }
    if yname == "capabilities" { return "Capabilities" }
    if yname == "backoff-parameters" { return "BackoffParameters" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "discovery" { return "Discovery" }
    if yname == "forwarding" { return "Forwarding" }
    if yname == "bindings" { return "Bindings" }
    if yname == "neighbors" { return "Neighbors" }
    if yname == "label-ranges" { return "LabelRanges" }
    return ""
}

func (mplsLdpState *MplsLdp_MplsLdpState) GetSegmentPath() string {
    return "mpls-ldp-state"
}

func (mplsLdpState *MplsLdp_MplsLdpState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "oper-summary" {
        return &mplsLdpState.OperSummary
    }
    if childYangName == "forwarding-summary" {
        return &mplsLdpState.ForwardingSummary
    }
    if childYangName == "bindings-summary" {
        return &mplsLdpState.BindingsSummary
    }
    if childYangName == "nsr-summary-all" {
        return &mplsLdpState.NsrSummaryAll
    }
    if childYangName == "icpm-summary-all" {
        return &mplsLdpState.IcpmSummaryAll
    }
    if childYangName == "parameters" {
        return &mplsLdpState.Parameters
    }
    if childYangName == "capabilities" {
        return &mplsLdpState.Capabilities
    }
    if childYangName == "backoff-parameters" {
        return &mplsLdpState.BackoffParameters
    }
    if childYangName == "graceful-restart" {
        return &mplsLdpState.GracefulRestart
    }
    if childYangName == "vrfs" {
        return &mplsLdpState.Vrfs
    }
    if childYangName == "discovery" {
        return &mplsLdpState.Discovery
    }
    if childYangName == "forwarding" {
        return &mplsLdpState.Forwarding
    }
    if childYangName == "bindings" {
        return &mplsLdpState.Bindings
    }
    if childYangName == "neighbors" {
        return &mplsLdpState.Neighbors
    }
    if childYangName == "label-ranges" {
        return &mplsLdpState.LabelRanges
    }
    return nil
}

func (mplsLdpState *MplsLdp_MplsLdpState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["oper-summary"] = &mplsLdpState.OperSummary
    children["forwarding-summary"] = &mplsLdpState.ForwardingSummary
    children["bindings-summary"] = &mplsLdpState.BindingsSummary
    children["nsr-summary-all"] = &mplsLdpState.NsrSummaryAll
    children["icpm-summary-all"] = &mplsLdpState.IcpmSummaryAll
    children["parameters"] = &mplsLdpState.Parameters
    children["capabilities"] = &mplsLdpState.Capabilities
    children["backoff-parameters"] = &mplsLdpState.BackoffParameters
    children["graceful-restart"] = &mplsLdpState.GracefulRestart
    children["vrfs"] = &mplsLdpState.Vrfs
    children["discovery"] = &mplsLdpState.Discovery
    children["forwarding"] = &mplsLdpState.Forwarding
    children["bindings"] = &mplsLdpState.Bindings
    children["neighbors"] = &mplsLdpState.Neighbors
    children["label-ranges"] = &mplsLdpState.LabelRanges
    return children
}

func (mplsLdpState *MplsLdp_MplsLdpState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsLdpState *MplsLdp_MplsLdpState) GetBundleName() string { return "cisco_ios_xe" }

func (mplsLdpState *MplsLdp_MplsLdpState) GetYangName() string { return "mpls-ldp-state" }

func (mplsLdpState *MplsLdp_MplsLdpState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsLdpState *MplsLdp_MplsLdpState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsLdpState *MplsLdp_MplsLdpState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsLdpState *MplsLdp_MplsLdpState) SetParent(parent types.Entity) { mplsLdpState.parent = parent }

func (mplsLdpState *MplsLdp_MplsLdpState) GetParent() types.Entity { return mplsLdpState.parent }

func (mplsLdpState *MplsLdp_MplsLdpState) GetParentYangName() string { return "mpls-ldp" }

// MplsLdp_MplsLdpState_OperSummary
// LDP operational data summary
type MplsLdp_MplsLdpState_OperSummary struct {
    parent types.Entity
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

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetFilter() yfilter.YFilter { return operSummary.YFilter }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) SetFilter(yf yfilter.YFilter) { operSummary.YFilter = yf }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetGoName(yname string) string {
    if yname == "number-of-vrf" { return "NumberOfVrf" }
    if yname == "number-of-vrf-oper" { return "NumberOfVrfOper" }
    if yname == "number-of-interfaces" { return "NumberOfInterfaces" }
    if yname == "number-of-fwd-ref-interfaces" { return "NumberOfFwdRefInterfaces" }
    if yname == "number-of-autocfg-interfaces" { return "NumberOfAutocfgInterfaces" }
    if yname == "no-of-ipv4-rib-tbl" { return "NoOfIpv4RibTbl" }
    if yname == "no-of-ipv4-rib-tbl-reg" { return "NoOfIpv4RibTblReg" }
    if yname == "common" { return "Common" }
    return ""
}

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetSegmentPath() string {
    return "oper-summary"
}

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "common" {
        return &operSummary.Common
    }
    return nil
}

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["common"] = &operSummary.Common
    return children
}

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["number-of-vrf"] = operSummary.NumberOfVrf
    leafs["number-of-vrf-oper"] = operSummary.NumberOfVrfOper
    leafs["number-of-interfaces"] = operSummary.NumberOfInterfaces
    leafs["number-of-fwd-ref-interfaces"] = operSummary.NumberOfFwdRefInterfaces
    leafs["number-of-autocfg-interfaces"] = operSummary.NumberOfAutocfgInterfaces
    leafs["no-of-ipv4-rib-tbl"] = operSummary.NoOfIpv4RibTbl
    leafs["no-of-ipv4-rib-tbl-reg"] = operSummary.NoOfIpv4RibTblReg
    return leafs
}

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetBundleName() string { return "cisco_ios_xe" }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetYangName() string { return "oper-summary" }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) SetParent(parent types.Entity) { operSummary.parent = parent }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetParent() types.Entity { return operSummary.parent }

func (operSummary *MplsLdp_MplsLdpState_OperSummary) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_OperSummary_Common
// Common Summary information
type MplsLdp_MplsLdpState_OperSummary_Common struct {
    parent types.Entity
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
    NumberOfIpv4LdpInterfaces interface{}
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetFilter() yfilter.YFilter { return common.YFilter }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) SetFilter(yf yfilter.YFilter) { common.YFilter = yf }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetGoName(yname string) string {
    if yname == "address-families" { return "AddressFamilies" }
    if yname == "number-of-neighbors" { return "NumberOfNeighbors" }
    if yname == "number-of-graceful-restart-neighbors" { return "NumberOfGracefulRestartNeighbors" }
    if yname == "number-of-downstream-on-demand-neighbors" { return "NumberOfDownstreamOnDemandNeighbors" }
    if yname == "numberof-ipv4-hello-adj" { return "NumberofIpv4HelloAdj" }
    if yname == "number-of-ipv4-routes" { return "NumberOfIpv4Routes" }
    if yname == "number-of-ipv4-local-addresses" { return "NumberOfIpv4LocalAddresses" }
    if yname == "number-of-ldp-interfaces" { return "NumberOfLdpInterfaces" }
    if yname == "number-of-ipv4ldp-interfaces" { return "NumberOfIpv4LdpInterfaces" }
    return ""
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetSegmentPath() string {
    return "common"
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-families"] = common.AddressFamilies
    leafs["number-of-neighbors"] = common.NumberOfNeighbors
    leafs["number-of-graceful-restart-neighbors"] = common.NumberOfGracefulRestartNeighbors
    leafs["number-of-downstream-on-demand-neighbors"] = common.NumberOfDownstreamOnDemandNeighbors
    leafs["numberof-ipv4-hello-adj"] = common.NumberofIpv4HelloAdj
    leafs["number-of-ipv4-routes"] = common.NumberOfIpv4Routes
    leafs["number-of-ipv4-local-addresses"] = common.NumberOfIpv4LocalAddresses
    leafs["number-of-ldp-interfaces"] = common.NumberOfLdpInterfaces
    leafs["number-of-ipv4ldp-interfaces"] = common.NumberOfIpv4LdpInterfaces
    return leafs
}

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetBundleName() string { return "cisco_ios_xe" }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetYangName() string { return "common" }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) SetParent(parent types.Entity) { common.parent = parent }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetParent() types.Entity { return common.parent }

func (common *MplsLdp_MplsLdpState_OperSummary_Common) GetParentYangName() string { return "oper-summary" }

// MplsLdp_MplsLdpState_ForwardingSummary
// Summary information regarding LDP forwarding
// setup
type MplsLdp_MplsLdpState_ForwardingSummary struct {
    parent types.Entity
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

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetFilter() yfilter.YFilter { return forwardingSummary.YFilter }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) SetFilter(yf yfilter.YFilter) { forwardingSummary.YFilter = yf }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetGoName(yname string) string {
    if yname == "intfs-fwd-count" { return "IntfsFwdCount" }
    if yname == "local-lbls" { return "LocalLbls" }
    if yname == "pfxs" { return "Pfxs" }
    if yname == "nhs" { return "Nhs" }
    return ""
}

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetSegmentPath() string {
    return "forwarding-summary"
}

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pfxs" {
        return &forwardingSummary.Pfxs
    }
    if childYangName == "nhs" {
        return &forwardingSummary.Nhs
    }
    return nil
}

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pfxs"] = &forwardingSummary.Pfxs
    children["nhs"] = &forwardingSummary.Nhs
    return children
}

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["intfs-fwd-count"] = forwardingSummary.IntfsFwdCount
    leafs["local-lbls"] = forwardingSummary.LocalLbls
    return leafs
}

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetBundleName() string { return "cisco_ios_xe" }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetYangName() string { return "forwarding-summary" }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) SetParent(parent types.Entity) { forwardingSummary.parent = parent }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetParent() types.Entity { return forwardingSummary.parent }

func (forwardingSummary *MplsLdp_MplsLdpState_ForwardingSummary) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs
// MPLS LDP forwarding prefix rewrite summary
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs struct {
    parent types.Entity
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

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetFilter() yfilter.YFilter { return pfxs.YFilter }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) SetFilter(yf yfilter.YFilter) { pfxs.YFilter = yf }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetGoName(yname string) string {
    if yname == "total-pfxs" { return "TotalPfxs" }
    if yname == "ecmp-pfxs" { return "EcmpPfxs" }
    if yname == "protected-pfxs" { return "ProtectedPfxs" }
    if yname == "labeled-pfxs-aggr" { return "LabeledPfxsAggr" }
    if yname == "labeled-pfxs-primary" { return "LabeledPfxsPrimary" }
    if yname == "labeled-pfxs-backup" { return "LabeledPfxsBackup" }
    return ""
}

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetSegmentPath() string {
    return "pfxs"
}

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled-pfxs-aggr" {
        return &pfxs.LabeledPfxsAggr
    }
    if childYangName == "labeled-pfxs-primary" {
        return &pfxs.LabeledPfxsPrimary
    }
    if childYangName == "labeled-pfxs-backup" {
        return &pfxs.LabeledPfxsBackup
    }
    return nil
}

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled-pfxs-aggr"] = &pfxs.LabeledPfxsAggr
    children["labeled-pfxs-primary"] = &pfxs.LabeledPfxsPrimary
    children["labeled-pfxs-backup"] = &pfxs.LabeledPfxsBackup
    return children
}

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-pfxs"] = pfxs.TotalPfxs
    leafs["ecmp-pfxs"] = pfxs.EcmpPfxs
    leafs["protected-pfxs"] = pfxs.ProtectedPfxs
    return leafs
}

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetBundleName() string { return "cisco_ios_xe" }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetYangName() string { return "pfxs" }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) SetParent(parent types.Entity) { pfxs.parent = parent }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetParent() types.Entity { return pfxs.parent }

func (pfxs *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs) GetParentYangName() string { return "forwarding-summary" }

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr
// Labeled prefix count for all paths
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr struct {
    parent types.Entity
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

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetFilter() yfilter.YFilter { return labeledPfxsAggr.YFilter }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) SetFilter(yf yfilter.YFilter) { labeledPfxsAggr.YFilter = yf }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetGoName(yname string) string {
    if yname == "labeled-pfxs" { return "LabeledPfxs" }
    if yname == "labeled-pfxs-partial" { return "LabeledPfxsPartial" }
    if yname == "unlabeled-pfxs" { return "UnlabeledPfxs" }
    return ""
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetSegmentPath() string {
    return "labeled-pfxs-aggr"
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["labeled-pfxs"] = labeledPfxsAggr.LabeledPfxs
    leafs["labeled-pfxs-partial"] = labeledPfxsAggr.LabeledPfxsPartial
    leafs["unlabeled-pfxs"] = labeledPfxsAggr.UnlabeledPfxs
    return leafs
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetBundleName() string { return "cisco_ios_xe" }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetYangName() string { return "labeled-pfxs-aggr" }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) SetParent(parent types.Entity) { labeledPfxsAggr.parent = parent }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetParent() types.Entity { return labeledPfxsAggr.parent }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsAggr) GetParentYangName() string { return "pfxs" }

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary
// Labeled prefix count related to primary paths
// only
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary struct {
    parent types.Entity
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

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetFilter() yfilter.YFilter { return labeledPfxsPrimary.YFilter }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) SetFilter(yf yfilter.YFilter) { labeledPfxsPrimary.YFilter = yf }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetGoName(yname string) string {
    if yname == "labeled-pfxs" { return "LabeledPfxs" }
    if yname == "labeled-pfxs-partial" { return "LabeledPfxsPartial" }
    if yname == "unlabeled-pfxs" { return "UnlabeledPfxs" }
    return ""
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetSegmentPath() string {
    return "labeled-pfxs-primary"
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["labeled-pfxs"] = labeledPfxsPrimary.LabeledPfxs
    leafs["labeled-pfxs-partial"] = labeledPfxsPrimary.LabeledPfxsPartial
    leafs["unlabeled-pfxs"] = labeledPfxsPrimary.UnlabeledPfxs
    return leafs
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetBundleName() string { return "cisco_ios_xe" }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetYangName() string { return "labeled-pfxs-primary" }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) SetParent(parent types.Entity) { labeledPfxsPrimary.parent = parent }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetParent() types.Entity { return labeledPfxsPrimary.parent }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsPrimary) GetParentYangName() string { return "pfxs" }

// MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup
// Labeled prefix count related to backup paths
// only
type MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup struct {
    parent types.Entity
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

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetFilter() yfilter.YFilter { return labeledPfxsBackup.YFilter }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) SetFilter(yf yfilter.YFilter) { labeledPfxsBackup.YFilter = yf }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetGoName(yname string) string {
    if yname == "labeled-pfxs" { return "LabeledPfxs" }
    if yname == "labeled-pfxs-partial" { return "LabeledPfxsPartial" }
    if yname == "unlabeled-pfxs" { return "UnlabeledPfxs" }
    return ""
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetSegmentPath() string {
    return "labeled-pfxs-backup"
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["labeled-pfxs"] = labeledPfxsBackup.LabeledPfxs
    leafs["labeled-pfxs-partial"] = labeledPfxsBackup.LabeledPfxsPartial
    leafs["unlabeled-pfxs"] = labeledPfxsBackup.UnlabeledPfxs
    return leafs
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetBundleName() string { return "cisco_ios_xe" }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetYangName() string { return "labeled-pfxs-backup" }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) SetParent(parent types.Entity) { labeledPfxsBackup.parent = parent }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetParent() types.Entity { return labeledPfxsBackup.parent }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_ForwardingSummary_Pfxs_LabeledPfxsBackup) GetParentYangName() string { return "pfxs" }

// MplsLdp_MplsLdpState_ForwardingSummary_Nhs
// MPLS LDP forwarding rewrite next-hop/path summary
type MplsLdp_MplsLdpState_ForwardingSummary_Nhs struct {
    parent types.Entity
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

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetFilter() yfilter.YFilter { return nhs.YFilter }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) SetFilter(yf yfilter.YFilter) { nhs.YFilter = yf }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetGoName(yname string) string {
    if yname == "total-paths" { return "TotalPaths" }
    if yname == "protected-paths" { return "ProtectedPaths" }
    if yname == "backup-paths" { return "BackupPaths" }
    if yname == "remote-backup-paths" { return "RemoteBackupPaths" }
    if yname == "labeled-paths" { return "LabeledPaths" }
    if yname == "labeled-backup-paths" { return "LabeledBackupPaths" }
    return ""
}

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetSegmentPath() string {
    return "nhs"
}

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-paths"] = nhs.TotalPaths
    leafs["protected-paths"] = nhs.ProtectedPaths
    leafs["backup-paths"] = nhs.BackupPaths
    leafs["remote-backup-paths"] = nhs.RemoteBackupPaths
    leafs["labeled-paths"] = nhs.LabeledPaths
    leafs["labeled-backup-paths"] = nhs.LabeledBackupPaths
    return leafs
}

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetBundleName() string { return "cisco_ios_xe" }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetYangName() string { return "nhs" }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) SetParent(parent types.Entity) { nhs.parent = parent }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetParent() types.Entity { return nhs.parent }

func (nhs *MplsLdp_MplsLdpState_ForwardingSummary_Nhs) GetParentYangName() string { return "forwarding-summary" }

// MplsLdp_MplsLdpState_BindingsSummary
// Aggregate counters for the MPLS LDP LIB.
type MplsLdp_MplsLdpState_BindingsSummary struct {
    parent types.Entity
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

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetFilter() yfilter.YFilter { return bindingsSummary.YFilter }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) SetFilter(yf yfilter.YFilter) { bindingsSummary.YFilter = yf }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetGoName(yname string) string {
    if yname == "binding-total" { return "BindingTotal" }
    if yname == "binding-no-route" { return "BindingNoRoute" }
    if yname == "binding-local-no-route" { return "BindingLocalNoRoute" }
    if yname == "binding-local" { return "BindingLocal" }
    if yname == "binding-local-null" { return "BindingLocalNull" }
    if yname == "binding-local-implicit-null" { return "BindingLocalImplicitNull" }
    if yname == "binding-local-explicit-null" { return "BindingLocalExplicitNull" }
    if yname == "binding-local-non-null" { return "BindingLocalNonNull" }
    if yname == "binding-local-oor" { return "BindingLocalOor" }
    if yname == "lowest-allocated-label" { return "LowestAllocatedLabel" }
    if yname == "highest-allocated-label" { return "HighestAllocatedLabel" }
    if yname == "binding-remote" { return "BindingRemote" }
    return ""
}

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetSegmentPath() string {
    return "bindings-summary"
}

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["binding-total"] = bindingsSummary.BindingTotal
    leafs["binding-no-route"] = bindingsSummary.BindingNoRoute
    leafs["binding-local-no-route"] = bindingsSummary.BindingLocalNoRoute
    leafs["binding-local"] = bindingsSummary.BindingLocal
    leafs["binding-local-null"] = bindingsSummary.BindingLocalNull
    leafs["binding-local-implicit-null"] = bindingsSummary.BindingLocalImplicitNull
    leafs["binding-local-explicit-null"] = bindingsSummary.BindingLocalExplicitNull
    leafs["binding-local-non-null"] = bindingsSummary.BindingLocalNonNull
    leafs["binding-local-oor"] = bindingsSummary.BindingLocalOor
    leafs["lowest-allocated-label"] = bindingsSummary.LowestAllocatedLabel
    leafs["highest-allocated-label"] = bindingsSummary.HighestAllocatedLabel
    leafs["binding-remote"] = bindingsSummary.BindingRemote
    return leafs
}

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetBundleName() string { return "cisco_ios_xe" }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetYangName() string { return "bindings-summary" }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) SetParent(parent types.Entity) { bindingsSummary.parent = parent }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetParent() types.Entity { return bindingsSummary.parent }

func (bindingsSummary *MplsLdp_MplsLdpState_BindingsSummary) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_NsrSummaryAll
// This is the LDP NSR summary for the device.
type MplsLdp_MplsLdpState_NsrSummaryAll struct {
    parent types.Entity
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

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetFilter() yfilter.YFilter { return nsrSummaryAll.YFilter }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) SetFilter(yf yfilter.YFilter) { nsrSummaryAll.YFilter = yf }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetGoName(yname string) string {
    if yname == "nsr-sum-in-label-reqs-created" { return "NsrSumInLabelReqsCreated" }
    if yname == "nsr-sum-in-label-reqs-freed" { return "NsrSumInLabelReqsFreed" }
    if yname == "nsr-sum-in-label-withdraw-created" { return "NsrSumInLabelWithdrawCreated" }
    if yname == "nsr-sum-in-label-withdraw-freed" { return "NsrSumInLabelWithdrawFreed" }
    if yname == "nsr-sum-lcl-addr-withdraw-set" { return "NsrSumLclAddrWithdrawSet" }
    if yname == "nsr-sum-lcl-addr-withdraw-cleared" { return "NsrSumLclAddrWithdrawCleared" }
    return ""
}

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetSegmentPath() string {
    return "nsr-summary-all"
}

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nsr-sum-in-label-reqs-created"] = nsrSummaryAll.NsrSumInLabelReqsCreated
    leafs["nsr-sum-in-label-reqs-freed"] = nsrSummaryAll.NsrSumInLabelReqsFreed
    leafs["nsr-sum-in-label-withdraw-created"] = nsrSummaryAll.NsrSumInLabelWithdrawCreated
    leafs["nsr-sum-in-label-withdraw-freed"] = nsrSummaryAll.NsrSumInLabelWithdrawFreed
    leafs["nsr-sum-lcl-addr-withdraw-set"] = nsrSummaryAll.NsrSumLclAddrWithdrawSet
    leafs["nsr-sum-lcl-addr-withdraw-cleared"] = nsrSummaryAll.NsrSumLclAddrWithdrawCleared
    return leafs
}

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetBundleName() string { return "cisco_ios_xe" }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetYangName() string { return "nsr-summary-all" }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) SetParent(parent types.Entity) { nsrSummaryAll.parent = parent }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetParent() types.Entity { return nsrSummaryAll.parent }

func (nsrSummaryAll *MplsLdp_MplsLdpState_NsrSummaryAll) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_IcpmSummaryAll
// Summary info for LDP ICPM/ICCP.
type MplsLdp_MplsLdpState_IcpmSummaryAll struct {
    parent types.Entity
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

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetFilter() yfilter.YFilter { return icpmSummaryAll.YFilter }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) SetFilter(yf yfilter.YFilter) { icpmSummaryAll.YFilter = yf }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetGoName(yname string) string {
    if yname == "iccp-rg-conn-count" { return "IccpRgConnCount" }
    if yname == "iccp-rg-disconn-count" { return "IccpRgDisconnCount" }
    if yname == "iccp-rg-notif-count" { return "IccpRgNotifCount" }
    if yname == "iccp-rg-app-data-count" { return "IccpRgAppDataCount" }
    if yname == "icpm-rgid-table-info" { return "IcpmRgidTableInfo" }
    if yname == "icpm-session-table" { return "IcpmSessionTable" }
    return ""
}

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetSegmentPath() string {
    return "icpm-summary-all"
}

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "icpm-rgid-table-info" {
        return &icpmSummaryAll.IcpmRgidTableInfo
    }
    if childYangName == "icpm-session-table" {
        return &icpmSummaryAll.IcpmSessionTable
    }
    return nil
}

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["icpm-rgid-table-info"] = &icpmSummaryAll.IcpmRgidTableInfo
    children["icpm-session-table"] = &icpmSummaryAll.IcpmSessionTable
    return children
}

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["iccp-rg-conn-count"] = icpmSummaryAll.IccpRgConnCount
    leafs["iccp-rg-disconn-count"] = icpmSummaryAll.IccpRgDisconnCount
    leafs["iccp-rg-notif-count"] = icpmSummaryAll.IccpRgNotifCount
    leafs["iccp-rg-app-data-count"] = icpmSummaryAll.IccpRgAppDataCount
    return leafs
}

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetBundleName() string { return "cisco_ios_xe" }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetYangName() string { return "icpm-summary-all" }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) SetParent(parent types.Entity) { icpmSummaryAll.parent = parent }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetParent() types.Entity { return icpmSummaryAll.parent }

func (icpmSummaryAll *MplsLdp_MplsLdpState_IcpmSummaryAll) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo
// This defines the ICPM RGID Table
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is the data for an individual ICPM Rredundandy Group,. The type is
    // slice of MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup.
    RedGroup []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetFilter() yfilter.YFilter { return icpmRgidTableInfo.YFilter }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) SetFilter(yf yfilter.YFilter) { icpmRgidTableInfo.YFilter = yf }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetGoName(yname string) string {
    if yname == "red-group" { return "RedGroup" }
    return ""
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetSegmentPath() string {
    return "icpm-rgid-table-info"
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "red-group" {
        for _, c := range icpmRgidTableInfo.RedGroup {
            if icpmRgidTableInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup{}
        icpmRgidTableInfo.RedGroup = append(icpmRgidTableInfo.RedGroup, child)
        return &icpmRgidTableInfo.RedGroup[len(icpmRgidTableInfo.RedGroup)-1]
    }
    return nil
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icpmRgidTableInfo.RedGroup {
        children[icpmRgidTableInfo.RedGroup[i].GetSegmentPath()] = &icpmRgidTableInfo.RedGroup[i]
    }
    return children
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetBundleName() string { return "cisco_ios_xe" }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetYangName() string { return "icpm-rgid-table-info" }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) SetParent(parent types.Entity) { icpmRgidTableInfo.parent = parent }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetParent() types.Entity { return icpmRgidTableInfo.parent }

func (icpmRgidTableInfo *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo) GetParentYangName() string { return "icpm-summary-all" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup
// This is the data for an individual ICPM Rredundandy
// Group,
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the ICPM RG identifier. The type is
    // interface{} with range: 0..4294967295.
    RgId interface{}

    // This list contains all active icpm protocols. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols.
    IcpmProtocols []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetFilter() yfilter.YFilter { return redGroup.YFilter }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) SetFilter(yf yfilter.YFilter) { redGroup.YFilter = yf }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetGoName(yname string) string {
    if yname == "rg-id" { return "RgId" }
    if yname == "icpm-protocols" { return "IcpmProtocols" }
    return ""
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetSegmentPath() string {
    return "red-group" + "[rg-id='" + fmt.Sprintf("%v", redGroup.RgId) + "']"
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "icpm-protocols" {
        for _, c := range redGroup.IcpmProtocols {
            if redGroup.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols{}
        redGroup.IcpmProtocols = append(redGroup.IcpmProtocols, child)
        return &redGroup.IcpmProtocols[len(redGroup.IcpmProtocols)-1]
    }
    return nil
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redGroup.IcpmProtocols {
        children[redGroup.IcpmProtocols[i].GetSegmentPath()] = &redGroup.IcpmProtocols[i]
    }
    return children
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rg-id"] = redGroup.RgId
    return leafs
}

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetBundleName() string { return "cisco_ios_xe" }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetYangName() string { return "red-group" }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) SetParent(parent types.Entity) { redGroup.parent = parent }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetParent() types.Entity { return redGroup.parent }

func (redGroup *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup) GetParentYangName() string { return "icpm-rgid-table-info" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols
// This list contains all active icpm protocols.
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ICPM Type. The type is one of the following:
    // IcpmTypeIccp.
    IcpmType interface{}

    // List of Redundancy Groups. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups.
    RedunGroups []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetFilter() yfilter.YFilter { return icpmProtocols.YFilter }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) SetFilter(yf yfilter.YFilter) { icpmProtocols.YFilter = yf }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetGoName(yname string) string {
    if yname == "icpm-type" { return "IcpmType" }
    if yname == "redun-groups" { return "RedunGroups" }
    return ""
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetSegmentPath() string {
    return "icpm-protocols" + "[icpm-type='" + fmt.Sprintf("%v", icpmProtocols.IcpmType) + "']"
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redun-groups" {
        for _, c := range icpmProtocols.RedunGroups {
            if icpmProtocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups{}
        icpmProtocols.RedunGroups = append(icpmProtocols.RedunGroups, child)
        return &icpmProtocols.RedunGroups[len(icpmProtocols.RedunGroups)-1]
    }
    return nil
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icpmProtocols.RedunGroups {
        children[icpmProtocols.RedunGroups[i].GetSegmentPath()] = &icpmProtocols.RedunGroups[i]
    }
    return children
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icpm-type"] = icpmProtocols.IcpmType
    return leafs
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetBundleName() string { return "cisco_ios_xe" }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetYangName() string { return "icpm-protocols" }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) SetParent(parent types.Entity) { icpmProtocols.parent = parent }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetParent() types.Entity { return icpmProtocols.parent }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols) GetParentYangName() string { return "red-group" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups
// List of Redundancy Groups
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups struct {
    parent types.Entity
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
    IccpApps []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetFilter() yfilter.YFilter { return redunGroups.YFilter }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) SetFilter(yf yfilter.YFilter) { redunGroups.YFilter = yf }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetGoName(yname string) string {
    if yname == "rg-id" { return "RgId" }
    if yname == "peer-id" { return "PeerId" }
    if yname == "client_id" { return "ClientId" }
    if yname == "state" { return "State" }
    if yname == "iccp-apps" { return "IccpApps" }
    return ""
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetSegmentPath() string {
    return "redun-groups" + "[rg-id='" + fmt.Sprintf("%v", redunGroups.RgId) + "']"
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "iccp-apps" {
        for _, c := range redunGroups.IccpApps {
            if redunGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps{}
        redunGroups.IccpApps = append(redunGroups.IccpApps, child)
        return &redunGroups.IccpApps[len(redunGroups.IccpApps)-1]
    }
    return nil
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redunGroups.IccpApps {
        children[redunGroups.IccpApps[i].GetSegmentPath()] = &redunGroups.IccpApps[i]
    }
    return children
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rg-id"] = redunGroups.RgId
    leafs["peer-id"] = redunGroups.PeerId
    leafs["client_id"] = redunGroups.ClientId
    leafs["state"] = redunGroups.State
    return leafs
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetBundleName() string { return "cisco_ios_xe" }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetYangName() string { return "redun-groups" }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) SetParent(parent types.Entity) { redunGroups.parent = parent }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetParent() types.Entity { return redunGroups.parent }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups) GetParentYangName() string { return "icpm-protocols" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps
// List of apps
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps struct {
    parent types.Entity
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

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetFilter() yfilter.YFilter { return iccpApps.YFilter }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) SetFilter(yf yfilter.YFilter) { iccpApps.YFilter = yf }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetGoName(yname string) string {
    if yname == "iccp-app" { return "IccpApp" }
    if yname == "app-state" { return "AppState" }
    if yname == "ptcl-ver" { return "PtclVer" }
    return ""
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetSegmentPath() string {
    return "iccp-apps" + "[iccp-app='" + fmt.Sprintf("%v", iccpApps.IccpApp) + "']"
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["iccp-app"] = iccpApps.IccpApp
    leafs["app-state"] = iccpApps.AppState
    leafs["ptcl-ver"] = iccpApps.PtclVer
    return leafs
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetBundleName() string { return "cisco_ios_xe" }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetYangName() string { return "iccp-apps" }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) SetParent(parent types.Entity) { iccpApps.parent = parent }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetParent() types.Entity { return iccpApps.parent }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmRgidTableInfo_RedGroup_IcpmProtocols_RedunGroups_IccpApps) GetParentYangName() string { return "redun-groups" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable
// This is a list of ICPM sessions.
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ICPM LDP Session Table. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable.
    SessionTable []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetFilter() yfilter.YFilter { return icpmSessionTable.YFilter }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) SetFilter(yf yfilter.YFilter) { icpmSessionTable.YFilter = yf }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetGoName(yname string) string {
    if yname == "session-table" { return "SessionTable" }
    return ""
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetSegmentPath() string {
    return "icpm-session-table"
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-table" {
        for _, c := range icpmSessionTable.SessionTable {
            if icpmSessionTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable{}
        icpmSessionTable.SessionTable = append(icpmSessionTable.SessionTable, child)
        return &icpmSessionTable.SessionTable[len(icpmSessionTable.SessionTable)-1]
    }
    return nil
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icpmSessionTable.SessionTable {
        children[icpmSessionTable.SessionTable[i].GetSegmentPath()] = &icpmSessionTable.SessionTable[i]
    }
    return children
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetBundleName() string { return "cisco_ios_xe" }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetYangName() string { return "icpm-session-table" }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) SetParent(parent types.Entity) { icpmSessionTable.parent = parent }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetParent() types.Entity { return icpmSessionTable.parent }

func (icpmSessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable) GetParentYangName() string { return "icpm-summary-all" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable
// ICPM LDP Session Table
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This is the ICPM sesion identifier. The type is
    // interface{} with range: 0..4294967295.
    SessionId interface{}

    // This list contains all active icpm protocols. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols.
    IcpmProtocols []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetFilter() yfilter.YFilter { return sessionTable.YFilter }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) SetFilter(yf yfilter.YFilter) { sessionTable.YFilter = yf }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "icpm-protocols" { return "IcpmProtocols" }
    return ""
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetSegmentPath() string {
    return "session-table" + "[session-id='" + fmt.Sprintf("%v", sessionTable.SessionId) + "']"
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "icpm-protocols" {
        for _, c := range sessionTable.IcpmProtocols {
            if sessionTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols{}
        sessionTable.IcpmProtocols = append(sessionTable.IcpmProtocols, child)
        return &sessionTable.IcpmProtocols[len(sessionTable.IcpmProtocols)-1]
    }
    return nil
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sessionTable.IcpmProtocols {
        children[sessionTable.IcpmProtocols[i].GetSegmentPath()] = &sessionTable.IcpmProtocols[i]
    }
    return children
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = sessionTable.SessionId
    return leafs
}

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetBundleName() string { return "cisco_ios_xe" }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetYangName() string { return "session-table" }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) SetParent(parent types.Entity) { sessionTable.parent = parent }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetParent() types.Entity { return sessionTable.parent }

func (sessionTable *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable) GetParentYangName() string { return "icpm-session-table" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols
// This list contains all active icpm protocols.
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ICPM Type. The type is one of the following:
    // IcpmTypeIccp.
    IcpmType interface{}

    // List of Redundancy Groups. The type is slice of
    // MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups.
    RedunGroups []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetFilter() yfilter.YFilter { return icpmProtocols.YFilter }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) SetFilter(yf yfilter.YFilter) { icpmProtocols.YFilter = yf }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetGoName(yname string) string {
    if yname == "icpm-type" { return "IcpmType" }
    if yname == "redun-groups" { return "RedunGroups" }
    return ""
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetSegmentPath() string {
    return "icpm-protocols" + "[icpm-type='" + fmt.Sprintf("%v", icpmProtocols.IcpmType) + "']"
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "redun-groups" {
        for _, c := range icpmProtocols.RedunGroups {
            if icpmProtocols.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups{}
        icpmProtocols.RedunGroups = append(icpmProtocols.RedunGroups, child)
        return &icpmProtocols.RedunGroups[len(icpmProtocols.RedunGroups)-1]
    }
    return nil
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range icpmProtocols.RedunGroups {
        children[icpmProtocols.RedunGroups[i].GetSegmentPath()] = &icpmProtocols.RedunGroups[i]
    }
    return children
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["icpm-type"] = icpmProtocols.IcpmType
    return leafs
}

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetBundleName() string { return "cisco_ios_xe" }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetYangName() string { return "icpm-protocols" }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) SetParent(parent types.Entity) { icpmProtocols.parent = parent }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetParent() types.Entity { return icpmProtocols.parent }

func (icpmProtocols *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols) GetParentYangName() string { return "session-table" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups
// List of Redundancy Groups
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups struct {
    parent types.Entity
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
    IccpApps []MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetFilter() yfilter.YFilter { return redunGroups.YFilter }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) SetFilter(yf yfilter.YFilter) { redunGroups.YFilter = yf }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetGoName(yname string) string {
    if yname == "rg-id" { return "RgId" }
    if yname == "peer-id" { return "PeerId" }
    if yname == "client_id" { return "ClientId" }
    if yname == "state" { return "State" }
    if yname == "iccp-apps" { return "IccpApps" }
    return ""
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetSegmentPath() string {
    return "redun-groups" + "[rg-id='" + fmt.Sprintf("%v", redunGroups.RgId) + "']"
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "iccp-apps" {
        for _, c := range redunGroups.IccpApps {
            if redunGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps{}
        redunGroups.IccpApps = append(redunGroups.IccpApps, child)
        return &redunGroups.IccpApps[len(redunGroups.IccpApps)-1]
    }
    return nil
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range redunGroups.IccpApps {
        children[redunGroups.IccpApps[i].GetSegmentPath()] = &redunGroups.IccpApps[i]
    }
    return children
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rg-id"] = redunGroups.RgId
    leafs["peer-id"] = redunGroups.PeerId
    leafs["client_id"] = redunGroups.ClientId
    leafs["state"] = redunGroups.State
    return leafs
}

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetBundleName() string { return "cisco_ios_xe" }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetYangName() string { return "redun-groups" }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) SetParent(parent types.Entity) { redunGroups.parent = parent }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetParent() types.Entity { return redunGroups.parent }

func (redunGroups *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups) GetParentYangName() string { return "icpm-protocols" }

// MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps
// List of apps
type MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps struct {
    parent types.Entity
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

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetFilter() yfilter.YFilter { return iccpApps.YFilter }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) SetFilter(yf yfilter.YFilter) { iccpApps.YFilter = yf }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetGoName(yname string) string {
    if yname == "iccp-app" { return "IccpApp" }
    if yname == "app-state" { return "AppState" }
    if yname == "ptcl-ver" { return "PtclVer" }
    return ""
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetSegmentPath() string {
    return "iccp-apps" + "[iccp-app='" + fmt.Sprintf("%v", iccpApps.IccpApp) + "']"
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["iccp-app"] = iccpApps.IccpApp
    leafs["app-state"] = iccpApps.AppState
    leafs["ptcl-ver"] = iccpApps.PtclVer
    return leafs
}

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetBundleName() string { return "cisco_ios_xe" }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetYangName() string { return "iccp-apps" }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) SetParent(parent types.Entity) { iccpApps.parent = parent }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetParent() types.Entity { return iccpApps.parent }

func (iccpApps *MplsLdp_MplsLdpState_IcpmSummaryAll_IcpmSessionTable_SessionTable_IcpmProtocols_RedunGroups_IccpApps) GetParentYangName() string { return "redun-groups" }

// MplsLdp_MplsLdpState_Parameters
// MPLS LDP Global Parameters
type MplsLdp_MplsLdpState_Parameters struct {
    parent types.Entity
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
    AddressFamilyParameter []MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetFilter() yfilter.YFilter { return parameters.YFilter }

func (parameters *MplsLdp_MplsLdpState_Parameters) SetFilter(yf yfilter.YFilter) { parameters.YFilter = yf }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetGoName(yname string) string {
    if yname == "global-md5-password-enabled" { return "GlobalMd5PasswordEnabled" }
    if yname == "protocol-version" { return "ProtocolVersion" }
    if yname == "keepalive-interval" { return "KeepaliveInterval" }
    if yname == "session-hold-time" { return "SessionHoldTime" }
    if yname == "le-no-route-timeout" { return "LeNoRouteTimeout" }
    if yname == "af-binding-withdraw-delay" { return "AfBindingWithdrawDelay" }
    if yname == "max-intf-attached" { return "MaxIntfAttached" }
    if yname == "max-intf-te" { return "MaxIntfTe" }
    if yname == "max-peer" { return "MaxPeer" }
    if yname == "out-of-mem-state" { return "OutOfMemState" }
    if yname == "discovery-quick-start-disabled-on-interfaces" { return "DiscoveryQuickStartDisabledOnInterfaces" }
    if yname == "dod-max-hop" { return "DodMaxHop" }
    if yname == "feature" { return "Feature" }
    if yname == "loop-detection" { return "LoopDetection" }
    if yname == "address-family-parameter" { return "AddressFamilyParameter" }
    return ""
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetSegmentPath() string {
    return "parameters"
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "address-family-parameter" {
        for _, c := range parameters.AddressFamilyParameter {
            if parameters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter{}
        parameters.AddressFamilyParameter = append(parameters.AddressFamilyParameter, child)
        return &parameters.AddressFamilyParameter[len(parameters.AddressFamilyParameter)-1]
    }
    return nil
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range parameters.AddressFamilyParameter {
        children[parameters.AddressFamilyParameter[i].GetSegmentPath()] = &parameters.AddressFamilyParameter[i]
    }
    return children
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-md5-password-enabled"] = parameters.GlobalMd5PasswordEnabled
    leafs["protocol-version"] = parameters.ProtocolVersion
    leafs["keepalive-interval"] = parameters.KeepaliveInterval
    leafs["session-hold-time"] = parameters.SessionHoldTime
    leafs["le-no-route-timeout"] = parameters.LeNoRouteTimeout
    leafs["af-binding-withdraw-delay"] = parameters.AfBindingWithdrawDelay
    leafs["max-intf-attached"] = parameters.MaxIntfAttached
    leafs["max-intf-te"] = parameters.MaxIntfTe
    leafs["max-peer"] = parameters.MaxPeer
    leafs["out-of-mem-state"] = parameters.OutOfMemState
    leafs["discovery-quick-start-disabled-on-interfaces"] = parameters.DiscoveryQuickStartDisabledOnInterfaces
    leafs["dod-max-hop"] = parameters.DodMaxHop
    leafs["feature"] = parameters.Feature
    leafs["loop-detection"] = parameters.LoopDetection
    return leafs
}

func (parameters *MplsLdp_MplsLdpState_Parameters) GetBundleName() string { return "cisco_ios_xe" }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetYangName() string { return "parameters" }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (parameters *MplsLdp_MplsLdpState_Parameters) SetParent(parent types.Entity) { parameters.parent = parent }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetParent() types.Entity { return parameters.parent }

func (parameters *MplsLdp_MplsLdpState_Parameters) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter
// Per AF parameters
type MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter struct {
    parent types.Entity
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

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetFilter() yfilter.YFilter { return addressFamilyParameter.YFilter }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) SetFilter(yf yfilter.YFilter) { addressFamilyParameter.YFilter = yf }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "discovery-transport-address" { return "DiscoveryTransportAddress" }
    if yname == "is-accepting-targeted-hellos" { return "IsAcceptingTargetedHellos" }
    if yname == "targeted-hello-filter" { return "TargetedHelloFilter" }
    return ""
}

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetSegmentPath() string {
    return "address-family-parameter" + "[address-family='" + fmt.Sprintf("%v", addressFamilyParameter.AddressFamily) + "']"
}

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = addressFamilyParameter.AddressFamily
    leafs["discovery-transport-address"] = addressFamilyParameter.DiscoveryTransportAddress
    leafs["is-accepting-targeted-hellos"] = addressFamilyParameter.IsAcceptingTargetedHellos
    leafs["targeted-hello-filter"] = addressFamilyParameter.TargetedHelloFilter
    return leafs
}

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetBundleName() string { return "cisco_ios_xe" }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetYangName() string { return "address-family-parameter" }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) SetParent(parent types.Entity) { addressFamilyParameter.parent = parent }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetParent() types.Entity { return addressFamilyParameter.parent }

func (addressFamilyParameter *MplsLdp_MplsLdpState_Parameters_AddressFamilyParameter) GetParentYangName() string { return "parameters" }

// MplsLdp_MplsLdpState_Capabilities
// LDP capability database information
type MplsLdp_MplsLdpState_Capabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information on LDP capability. The type is slice of
    // MplsLdp_MplsLdpState_Capabilities_Capability.
    Capability []MplsLdp_MplsLdpState_Capabilities_Capability
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetFilter() yfilter.YFilter { return capabilities.YFilter }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) SetFilter(yf yfilter.YFilter) { capabilities.YFilter = yf }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetGoName(yname string) string {
    if yname == "capability" { return "Capability" }
    return ""
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetSegmentPath() string {
    return "capabilities"
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "capability" {
        for _, c := range capabilities.Capability {
            if capabilities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Capabilities_Capability{}
        capabilities.Capability = append(capabilities.Capability, child)
        return &capabilities.Capability[len(capabilities.Capability)-1]
    }
    return nil
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range capabilities.Capability {
        children[capabilities.Capability[i].GetSegmentPath()] = &capabilities.Capability[i]
    }
    return children
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetBundleName() string { return "cisco_ios_xe" }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetYangName() string { return "capabilities" }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) SetParent(parent types.Entity) { capabilities.parent = parent }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetParent() types.Entity { return capabilities.parent }

func (capabilities *MplsLdp_MplsLdpState_Capabilities) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Capabilities_Capability
// Information on LDP capability
type MplsLdp_MplsLdpState_Capabilities_Capability struct {
    parent types.Entity
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

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetFilter() yfilter.YFilter { return capability.YFilter }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) SetFilter(yf yfilter.YFilter) { capability.YFilter = yf }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetGoName(yname string) string {
    if yname == "cap-type" { return "CapType" }
    if yname == "capability-owner" { return "CapabilityOwner" }
    if yname == "cap-des" { return "CapDes" }
    if yname == "capability-data-length" { return "CapabilityDataLength" }
    if yname == "capability-data" { return "CapabilityData" }
    return ""
}

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetSegmentPath() string {
    return "capability" + "[cap-type='" + fmt.Sprintf("%v", capability.CapType) + "']"
}

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cap-type"] = capability.CapType
    leafs["capability-owner"] = capability.CapabilityOwner
    leafs["cap-des"] = capability.CapDes
    leafs["capability-data-length"] = capability.CapabilityDataLength
    leafs["capability-data"] = capability.CapabilityData
    return leafs
}

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetBundleName() string { return "cisco_ios_xe" }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetYangName() string { return "capability" }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) SetParent(parent types.Entity) { capability.parent = parent }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetParent() types.Entity { return capability.parent }

func (capability *MplsLdp_MplsLdpState_Capabilities_Capability) GetParentYangName() string { return "capabilities" }

// MplsLdp_MplsLdpState_BackoffParameters
// MPLS LDP Session Backoff Information
type MplsLdp_MplsLdpState_BackoffParameters struct {
    parent types.Entity
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

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetFilter() yfilter.YFilter { return backoffParameters.YFilter }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) SetFilter(yf yfilter.YFilter) { backoffParameters.YFilter = yf }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetGoName(yname string) string {
    if yname == "initial-seconds" { return "InitialSeconds" }
    if yname == "maximum-seconds" { return "MaximumSeconds" }
    if yname == "backoff-seconds" { return "BackoffSeconds" }
    if yname == "waiting-seconds" { return "WaitingSeconds" }
    return ""
}

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetSegmentPath() string {
    return "backoff-parameters"
}

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initial-seconds"] = backoffParameters.InitialSeconds
    leafs["maximum-seconds"] = backoffParameters.MaximumSeconds
    leafs["backoff-seconds"] = backoffParameters.BackoffSeconds
    leafs["waiting-seconds"] = backoffParameters.WaitingSeconds
    return leafs
}

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetBundleName() string { return "cisco_ios_xe" }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetYangName() string { return "backoff-parameters" }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) SetParent(parent types.Entity) { backoffParameters.parent = parent }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetParent() types.Entity { return backoffParameters.parent }

func (backoffParameters *MplsLdp_MplsLdpState_BackoffParameters) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_GracefulRestart
// MPLS LDP Graceful Restart Information
type MplsLdp_MplsLdpState_GracefulRestart struct {
    parent types.Entity
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

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetGoName(yname string) string {
    if yname == "is-graceful-restart-configured" { return "IsGracefulRestartConfigured" }
    if yname == "graceful-restart-reconnect-timeout" { return "GracefulRestartReconnectTimeout" }
    if yname == "graceful-restart-forwarding-state-hold-time" { return "GracefulRestartForwardingStateHoldTime" }
    if yname == "is-forwarding-state-hold-timer-running" { return "IsForwardingStateHoldTimerRunning" }
    if yname == "forwarding-state-hold-timer-remaining-seconds" { return "ForwardingStateHoldTimerRemainingSeconds" }
    return ""
}

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-graceful-restart-configured"] = gracefulRestart.IsGracefulRestartConfigured
    leafs["graceful-restart-reconnect-timeout"] = gracefulRestart.GracefulRestartReconnectTimeout
    leafs["graceful-restart-forwarding-state-hold-time"] = gracefulRestart.GracefulRestartForwardingStateHoldTime
    leafs["is-forwarding-state-hold-timer-running"] = gracefulRestart.IsForwardingStateHoldTimerRunning
    leafs["forwarding-state-hold-timer-remaining-seconds"] = gracefulRestart.ForwardingStateHoldTimerRemainingSeconds
    return leafs
}

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetBundleName() string { return "cisco_ios_xe" }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *MplsLdp_MplsLdpState_GracefulRestart) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Vrfs
// MPLS LDP per-VRF operational data.
type MplsLdp_MplsLdpState_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP Operational data for a given VRF. The type is slice of
    // MplsLdp_MplsLdpState_Vrfs_Vrf.
    Vrf []MplsLdp_MplsLdpState_Vrfs_Vrf
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetBundleName() string { return "cisco_ios_xe" }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *MplsLdp_MplsLdpState_Vrfs) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Vrfs_Vrf
// MPLS LDP Operational data for a given VRF.
type MplsLdp_MplsLdpState_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. This contains the VRF Name, where 'default' is
    // used for the default vrf. The type is string.
    VrfName interface{}

    // MPLS LDP per VRF summarized Information.
    VrfSummary MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary

    // Address Family specific operational data.
    Afs MplsLdp_MplsLdpState_Vrfs_Vrf_Afs
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "vrf-summary" { return "VrfSummary" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-summary" {
        return &vrf.VrfSummary
    }
    if childYangName == "afs" {
        return &vrf.Afs
    }
    return nil
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-summary"] = &vrf.VrfSummary
    children["afs"] = &vrf.Afs
    return children
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xe" }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *MplsLdp_MplsLdpState_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary
// MPLS LDP per VRF summarized Information
type MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary struct {
    parent types.Entity
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
    NumberOfIpv4LdpInterfaces interface{}
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetFilter() yfilter.YFilter { return vrfSummary.YFilter }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) SetFilter(yf yfilter.YFilter) { vrfSummary.YFilter = yf }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetGoName(yname string) string {
    if yname == "address-families" { return "AddressFamilies" }
    if yname == "number-of-neighbors" { return "NumberOfNeighbors" }
    if yname == "number-of-graceful-restart-neighbors" { return "NumberOfGracefulRestartNeighbors" }
    if yname == "number-of-downstream-on-demand-neighbors" { return "NumberOfDownstreamOnDemandNeighbors" }
    if yname == "numberof-ipv4-hello-adj" { return "NumberofIpv4HelloAdj" }
    if yname == "number-of-ipv4-routes" { return "NumberOfIpv4Routes" }
    if yname == "number-of-ipv4-local-addresses" { return "NumberOfIpv4LocalAddresses" }
    if yname == "number-of-ldp-interfaces" { return "NumberOfLdpInterfaces" }
    if yname == "number-of-ipv4ldp-interfaces" { return "NumberOfIpv4LdpInterfaces" }
    return ""
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetSegmentPath() string {
    return "vrf-summary"
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-families"] = vrfSummary.AddressFamilies
    leafs["number-of-neighbors"] = vrfSummary.NumberOfNeighbors
    leafs["number-of-graceful-restart-neighbors"] = vrfSummary.NumberOfGracefulRestartNeighbors
    leafs["number-of-downstream-on-demand-neighbors"] = vrfSummary.NumberOfDownstreamOnDemandNeighbors
    leafs["numberof-ipv4-hello-adj"] = vrfSummary.NumberofIpv4HelloAdj
    leafs["number-of-ipv4-routes"] = vrfSummary.NumberOfIpv4Routes
    leafs["number-of-ipv4-local-addresses"] = vrfSummary.NumberOfIpv4LocalAddresses
    leafs["number-of-ldp-interfaces"] = vrfSummary.NumberOfLdpInterfaces
    leafs["number-of-ipv4ldp-interfaces"] = vrfSummary.NumberOfIpv4LdpInterfaces
    return leafs
}

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetBundleName() string { return "cisco_ios_xe" }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetYangName() string { return "vrf-summary" }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) SetParent(parent types.Entity) { vrfSummary.parent = parent }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetParent() types.Entity { return vrfSummary.parent }

func (vrfSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_VrfSummary) GetParentYangName() string { return "vrf" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs
// Address Family specific operational data
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP Operational data for this Address Family. The type is slice of
    // MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af.
    Af []MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetBundleName() string { return "cisco_ios_xe" }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs) GetParentYangName() string { return "vrf" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af
// MPLS LDP Operational data for this Address Family.
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family name. The type is Af.
    AfName interface{}

    // This container holds a summary of information across all interfaces in this
    // AF,.
    InterfaceSummary MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary

    // LDP IGP Synchronization related information.
    Igp MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "interface-summary" { return "InterfaceSummary" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-summary" {
        return &af.InterfaceSummary
    }
    if childYangName == "igp" {
        return &af.Igp
    }
    return nil
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-summary"] = &af.InterfaceSummary
    children["igp"] = &af.Igp
    return children
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    return leafs
}

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetBundleName() string { return "cisco_ios_xe" }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary
// This container holds a summary of information
// across all interfaces in this AF,
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary struct {
    parent types.Entity
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

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetFilter() yfilter.YFilter { return interfaceSummary.YFilter }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) SetFilter(yf yfilter.YFilter) { interfaceSummary.YFilter = yf }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetGoName(yname string) string {
    if yname == "known-ip-interface-count" { return "KnownIpInterfaceCount" }
    if yname == "known-ip-interface-ldp-enabled" { return "KnownIpInterfaceLdpEnabled" }
    if yname == "configured-attached-interface" { return "ConfiguredAttachedInterface" }
    if yname == "configured-te-interface" { return "ConfiguredTeInterface" }
    if yname == "forward-references" { return "ForwardReferences" }
    if yname == "auto-config-disabled" { return "AutoConfigDisabled" }
    if yname == "auto-config" { return "AutoConfig" }
    if yname == "auto-config-forward-reference-interfaces" { return "AutoConfigForwardReferenceInterfaces" }
    return ""
}

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetSegmentPath() string {
    return "interface-summary"
}

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["known-ip-interface-count"] = interfaceSummary.KnownIpInterfaceCount
    leafs["known-ip-interface-ldp-enabled"] = interfaceSummary.KnownIpInterfaceLdpEnabled
    leafs["configured-attached-interface"] = interfaceSummary.ConfiguredAttachedInterface
    leafs["configured-te-interface"] = interfaceSummary.ConfiguredTeInterface
    leafs["forward-references"] = interfaceSummary.ForwardReferences
    leafs["auto-config-disabled"] = interfaceSummary.AutoConfigDisabled
    leafs["auto-config"] = interfaceSummary.AutoConfig
    leafs["auto-config-forward-reference-interfaces"] = interfaceSummary.AutoConfigForwardReferenceInterfaces
    return leafs
}

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetBundleName() string { return "cisco_ios_xe" }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetYangName() string { return "interface-summary" }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) SetParent(parent types.Entity) { interfaceSummary.parent = parent }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetParent() types.Entity { return interfaceSummary.parent }

func (interfaceSummary *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_InterfaceSummary) GetParentYangName() string { return "af" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp
// LDP IGP Synchronization related information
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP-IGP Synchronization related information for an interface. The type is
    // slice of MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync.
    Sync []MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetGoName(yname string) string {
    if yname == "sync" { return "Sync" }
    return ""
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sync" {
        for _, c := range igp.Sync {
            if igp.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync{}
        igp.Sync = append(igp.Sync, child)
        return &igp.Sync[len(igp.Sync)-1]
    }
    return nil
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range igp.Sync {
        children[igp.Sync[i].GetSegmentPath()] = &igp.Sync[i]
    }
    return children
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetBundleName() string { return "cisco_ios_xe" }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetYangName() string { return "igp" }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetParent() types.Entity { return igp.parent }

func (igp *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp) GetParentYangName() string { return "af" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync
// LDP-IGP Synchronization related information
// for an interface
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync struct {
    parent types.Entity
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
    Peers []MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetFilter() yfilter.YFilter { return sync.YFilter }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) SetFilter(yf yfilter.YFilter) { sync.YFilter = yf }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "igp-sync-state" { return "IgpSyncState" }
    if yname == "is-delay-timer-running" { return "IsDelayTimerRunning" }
    if yname == "delay-timer-remaining" { return "DelayTimerRemaining" }
    if yname == "igp-sync-down-reason" { return "IgpSyncDownReason" }
    if yname == "peers" { return "Peers" }
    return ""
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetSegmentPath() string {
    return "sync" + "[interface='" + fmt.Sprintf("%v", sync.Interface) + "']"
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peers" {
        for _, c := range sync.Peers {
            if sync.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers{}
        sync.Peers = append(sync.Peers, child)
        return &sync.Peers[len(sync.Peers)-1]
    }
    return nil
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range sync.Peers {
        children[sync.Peers[i].GetSegmentPath()] = &sync.Peers[i]
    }
    return children
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = sync.Interface
    leafs["igp-sync-state"] = sync.IgpSyncState
    leafs["is-delay-timer-running"] = sync.IsDelayTimerRunning
    leafs["delay-timer-remaining"] = sync.DelayTimerRemaining
    leafs["igp-sync-down-reason"] = sync.IgpSyncDownReason
    return leafs
}

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetBundleName() string { return "cisco_ios_xe" }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetYangName() string { return "sync" }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) SetParent(parent types.Entity) { sync.parent = parent }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetParent() types.Entity { return sync.parent }

func (sync *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync) GetParentYangName() string { return "igp" }

// MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers
// MPLS LDP IGP Sync Interface Peer Information
type MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer Identifier. The type is string.
    PeerId interface{}

    // Is GR enabled session. The type is bool.
    IsGrEnabled interface{}

    // This is set if this peer was created due to check-pointing. The type is
    // interface{}.
    IsChkptCreated interface{}
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetFilter() yfilter.YFilter { return peers.YFilter }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) SetFilter(yf yfilter.YFilter) { peers.YFilter = yf }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetGoName(yname string) string {
    if yname == "peer-id" { return "PeerId" }
    if yname == "is-gr-enabled" { return "IsGrEnabled" }
    if yname == "is-chkpt-created" { return "IsChkptCreated" }
    return ""
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetSegmentPath() string {
    return "peers"
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-id"] = peers.PeerId
    leafs["is-gr-enabled"] = peers.IsGrEnabled
    leafs["is-chkpt-created"] = peers.IsChkptCreated
    return leafs
}

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetBundleName() string { return "cisco_ios_xe" }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetYangName() string { return "peers" }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) SetParent(parent types.Entity) { peers.parent = parent }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetParent() types.Entity { return peers.parent }

func (peers *MplsLdp_MplsLdpState_Vrfs_Vrf_Afs_Af_Igp_Sync_Peers) GetParentYangName() string { return "sync" }

// MplsLdp_MplsLdpState_Discovery
// The LDP Discovery operational state
type MplsLdp_MplsLdpState_Discovery struct {
    parent types.Entity
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

func (discovery *MplsLdp_MplsLdpState_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_MplsLdpState_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetGoName(yname string) string {
    if yname == "discovery-stats" { return "DiscoveryStats" }
    if yname == "link-hello-state" { return "LinkHelloState" }
    if yname == "targeted-hellos" { return "TargetedHellos" }
    return ""
}

func (discovery *MplsLdp_MplsLdpState_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_MplsLdpState_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "discovery-stats" {
        return &discovery.DiscoveryStats
    }
    if childYangName == "link-hello-state" {
        return &discovery.LinkHelloState
    }
    if childYangName == "targeted-hellos" {
        return &discovery.TargetedHellos
    }
    return nil
}

func (discovery *MplsLdp_MplsLdpState_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["discovery-stats"] = &discovery.DiscoveryStats
    children["link-hello-state"] = &discovery.LinkHelloState
    children["targeted-hellos"] = &discovery.TargetedHellos
    return children
}

func (discovery *MplsLdp_MplsLdpState_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (discovery *MplsLdp_MplsLdpState_Discovery) GetBundleName() string { return "cisco_ios_xe" }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (discovery *MplsLdp_MplsLdpState_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_MplsLdpState_Discovery) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Discovery_DiscoveryStats
// MPLS LDP Discovery Summary Information
type MplsLdp_MplsLdpState_Discovery_DiscoveryStats struct {
    parent types.Entity
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

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetFilter() yfilter.YFilter { return discoveryStats.YFilter }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) SetFilter(yf yfilter.YFilter) { discoveryStats.YFilter = yf }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetGoName(yname string) string {
    if yname == "num-of-ldp-interfaces" { return "NumOfLdpInterfaces" }
    if yname == "num-of-active-ldp-interfaces" { return "NumOfActiveLdpInterfaces" }
    if yname == "num-of-lnk-disc-xmit" { return "NumOfLnkDiscXmit" }
    if yname == "num-of-tgt-disc-xmit" { return "NumOfTgtDiscXmit" }
    if yname == "num-of-lnk-disc-recv" { return "NumOfLnkDiscRecv" }
    if yname == "num-of-tgt-disc-recv" { return "NumOfTgtDiscRecv" }
    return ""
}

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetSegmentPath() string {
    return "discovery-stats"
}

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-of-ldp-interfaces"] = discoveryStats.NumOfLdpInterfaces
    leafs["num-of-active-ldp-interfaces"] = discoveryStats.NumOfActiveLdpInterfaces
    leafs["num-of-lnk-disc-xmit"] = discoveryStats.NumOfLnkDiscXmit
    leafs["num-of-tgt-disc-xmit"] = discoveryStats.NumOfTgtDiscXmit
    leafs["num-of-lnk-disc-recv"] = discoveryStats.NumOfLnkDiscRecv
    leafs["num-of-tgt-disc-recv"] = discoveryStats.NumOfTgtDiscRecv
    return leafs
}

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetBundleName() string { return "cisco_ios_xe" }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetYangName() string { return "discovery-stats" }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) SetParent(parent types.Entity) { discoveryStats.parent = parent }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetParent() types.Entity { return discoveryStats.parent }

func (discoveryStats *MplsLdp_MplsLdpState_Discovery_DiscoveryStats) GetParentYangName() string { return "discovery" }

// MplsLdp_MplsLdpState_Discovery_LinkHelloState
// This container holds information for LDP Discovery
// using non-targeted Hellos. These are interface-based
// hellos which form one or more adjacencies for each
// interface and also form adjacencies on multiple
// intefrfaces. Link Hellos can therefore form multiple
// adjacencies with the same peer.
type MplsLdp_MplsLdpState_Discovery_LinkHelloState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents a single LDP Hello Adjacency. An LDP Session can have
    // one or more Hello Adjacencies. The type is slice of
    // MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos.
    LinkHellos []MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetFilter() yfilter.YFilter { return linkHelloState.YFilter }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) SetFilter(yf yfilter.YFilter) { linkHelloState.YFilter = yf }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetGoName(yname string) string {
    if yname == "link-hellos" { return "LinkHellos" }
    return ""
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetSegmentPath() string {
    return "link-hello-state"
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-hellos" {
        for _, c := range linkHelloState.LinkHellos {
            if linkHelloState.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos{}
        linkHelloState.LinkHellos = append(linkHelloState.LinkHellos, child)
        return &linkHelloState.LinkHellos[len(linkHelloState.LinkHellos)-1]
    }
    return nil
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range linkHelloState.LinkHellos {
        children[linkHelloState.LinkHellos[i].GetSegmentPath()] = &linkHelloState.LinkHellos[i]
    }
    return children
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetBundleName() string { return "cisco_ios_xe" }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetYangName() string { return "link-hello-state" }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) SetParent(parent types.Entity) { linkHelloState.parent = parent }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetParent() types.Entity { return linkHelloState.parent }

func (linkHelloState *MplsLdp_MplsLdpState_Discovery_LinkHelloState) GetParentYangName() string { return "discovery" }

// MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos
// Each entry represents a single LDP Hello Adjacency.
// An LDP Session can have one or more Hello
// Adjacencies.
type MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos struct {
    parent types.Entity
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

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetFilter() yfilter.YFilter { return linkHellos.YFilter }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) SetFilter(yf yfilter.YFilter) { linkHellos.YFilter = yf }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "nbr-transport-addr" { return "NbrTransportAddr" }
    if yname == "hello-interval" { return "HelloInterval" }
    if yname == "local-src-addr" { return "LocalSrcAddr" }
    if yname == "local-transport-addr" { return "LocalTransportAddr" }
    if yname == "nbr-src-addr" { return "NbrSrcAddr" }
    if yname == "nbr-ldp-id" { return "NbrLdpId" }
    if yname == "session-up" { return "SessionUp" }
    if yname == "nbr-hold-time" { return "NbrHoldTime" }
    if yname == "next-hello" { return "NextHello" }
    if yname == "hold-time-remaining" { return "HoldTimeRemaining" }
    return ""
}

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetSegmentPath() string {
    return "link-hellos" + "[interface='" + fmt.Sprintf("%v", linkHellos.Interface) + "']" + "[nbr-transport-addr='" + fmt.Sprintf("%v", linkHellos.NbrTransportAddr) + "']"
}

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = linkHellos.Interface
    leafs["nbr-transport-addr"] = linkHellos.NbrTransportAddr
    leafs["hello-interval"] = linkHellos.HelloInterval
    leafs["local-src-addr"] = linkHellos.LocalSrcAddr
    leafs["local-transport-addr"] = linkHellos.LocalTransportAddr
    leafs["nbr-src-addr"] = linkHellos.NbrSrcAddr
    leafs["nbr-ldp-id"] = linkHellos.NbrLdpId
    leafs["session-up"] = linkHellos.SessionUp
    leafs["nbr-hold-time"] = linkHellos.NbrHoldTime
    leafs["next-hello"] = linkHellos.NextHello
    leafs["hold-time-remaining"] = linkHellos.HoldTimeRemaining
    return leafs
}

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetBundleName() string { return "cisco_ios_xe" }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetYangName() string { return "link-hellos" }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) SetParent(parent types.Entity) { linkHellos.parent = parent }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetParent() types.Entity { return linkHellos.parent }

func (linkHellos *MplsLdp_MplsLdpState_Discovery_LinkHelloState_LinkHellos) GetParentYangName() string { return "link-hello-state" }

// MplsLdp_MplsLdpState_Discovery_TargetedHellos
// The LDP Discovery Targeted Hello state.
type MplsLdp_MplsLdpState_Discovery_TargetedHellos struct {
    parent types.Entity
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
    TargetedHello []MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetFilter() yfilter.YFilter { return targetedHellos.YFilter }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) SetFilter(yf yfilter.YFilter) { targetedHellos.YFilter = yf }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetGoName(yname string) string {
    if yname == "targeted-hello-interval" { return "TargetedHelloInterval" }
    if yname == "targeted-hello-hold-time" { return "TargetedHelloHoldTime" }
    if yname == "targeted-hello" { return "TargetedHello" }
    return ""
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetSegmentPath() string {
    return "targeted-hellos"
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "targeted-hello" {
        for _, c := range targetedHellos.TargetedHello {
            if targetedHellos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello{}
        targetedHellos.TargetedHello = append(targetedHellos.TargetedHello, child)
        return &targetedHellos.TargetedHello[len(targetedHellos.TargetedHello)-1]
    }
    return nil
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range targetedHellos.TargetedHello {
        children[targetedHellos.TargetedHello[i].GetSegmentPath()] = &targetedHellos.TargetedHello[i]
    }
    return children
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["targeted-hello-interval"] = targetedHellos.TargetedHelloInterval
    leafs["targeted-hello-hold-time"] = targetedHellos.TargetedHelloHoldTime
    return leafs
}

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetBundleName() string { return "cisco_ios_xe" }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetYangName() string { return "targeted-hellos" }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) SetParent(parent types.Entity) { targetedHellos.parent = parent }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetParent() types.Entity { return targetedHellos.parent }

func (targetedHellos *MplsLdp_MplsLdpState_Discovery_TargetedHellos) GetParentYangName() string { return "discovery" }

// MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello
// The LDP targeted discovery information for a specific
// target. Targetted discovery creates a single adjacency
// between two addresses and not indiviual adjacencies
// across physical interfaces.
type MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello struct {
    parent types.Entity
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

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetFilter() yfilter.YFilter { return targetedHello.YFilter }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) SetFilter(yf yfilter.YFilter) { targetedHello.YFilter = yf }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "target-address" { return "TargetAddress" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "neighbor-ldp-identifier" { return "NeighborLdpIdentifier" }
    if yname == "state" { return "State" }
    if yname == "nbr-hold-time" { return "NbrHoldTime" }
    if yname == "next-hello" { return "NextHello" }
    if yname == "hold-time-remaining" { return "HoldTimeRemaining" }
    return ""
}

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetSegmentPath() string {
    return "targeted-hello" + "[vrf-name='" + fmt.Sprintf("%v", targetedHello.VrfName) + "']" + "[target-address='" + fmt.Sprintf("%v", targetedHello.TargetAddress) + "']"
}

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = targetedHello.VrfName
    leafs["target-address"] = targetedHello.TargetAddress
    leafs["local-address"] = targetedHello.LocalAddress
    leafs["neighbor-ldp-identifier"] = targetedHello.NeighborLdpIdentifier
    leafs["state"] = targetedHello.State
    leafs["nbr-hold-time"] = targetedHello.NbrHoldTime
    leafs["next-hello"] = targetedHello.NextHello
    leafs["hold-time-remaining"] = targetedHello.HoldTimeRemaining
    return leafs
}

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetBundleName() string { return "cisco_ios_xe" }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetYangName() string { return "targeted-hello" }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) SetParent(parent types.Entity) { targetedHello.parent = parent }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetParent() types.Entity { return targetedHello.parent }

func (targetedHello *MplsLdp_MplsLdpState_Discovery_TargetedHellos_TargetedHello) GetParentYangName() string { return "targeted-hellos" }

// MplsLdp_MplsLdpState_Forwarding
// Summary information regarding LDP forwarding
// setup and detailed LDP Forwarding rewrites
type MplsLdp_MplsLdpState_Forwarding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary of forwarding info for this VRF.
    ForwardingVrfSumms MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms

    // This leaf contain the individual LDP forwarding rewrite for a single
    // prefix. The type is slice of
    // MplsLdp_MplsLdpState_Forwarding_ForwardingDetail.
    ForwardingDetail []MplsLdp_MplsLdpState_Forwarding_ForwardingDetail
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetFilter() yfilter.YFilter { return forwarding.YFilter }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) SetFilter(yf yfilter.YFilter) { forwarding.YFilter = yf }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetGoName(yname string) string {
    if yname == "forwarding-vrf-summs" { return "ForwardingVrfSumms" }
    if yname == "forwarding-detail" { return "ForwardingDetail" }
    return ""
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetSegmentPath() string {
    return "forwarding"
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forwarding-vrf-summs" {
        return &forwarding.ForwardingVrfSumms
    }
    if childYangName == "forwarding-detail" {
        for _, c := range forwarding.ForwardingDetail {
            if forwarding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Forwarding_ForwardingDetail{}
        forwarding.ForwardingDetail = append(forwarding.ForwardingDetail, child)
        return &forwarding.ForwardingDetail[len(forwarding.ForwardingDetail)-1]
    }
    return nil
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["forwarding-vrf-summs"] = &forwarding.ForwardingVrfSumms
    for i := range forwarding.ForwardingDetail {
        children[forwarding.ForwardingDetail[i].GetSegmentPath()] = &forwarding.ForwardingDetail[i]
    }
    return children
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetBundleName() string { return "cisco_ios_xe" }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetYangName() string { return "forwarding" }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) SetParent(parent types.Entity) { forwarding.parent = parent }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetParent() types.Entity { return forwarding.parent }

func (forwarding *MplsLdp_MplsLdpState_Forwarding) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms
// Summary of forwarding info for this VRF.
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary of forwarding info for this VRF. The type is slice of
    // MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm.
    ForwardingVrfSumm []MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetFilter() yfilter.YFilter { return forwardingVrfSumms.YFilter }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) SetFilter(yf yfilter.YFilter) { forwardingVrfSumms.YFilter = yf }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetGoName(yname string) string {
    if yname == "forwarding-vrf-summ" { return "ForwardingVrfSumm" }
    return ""
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetSegmentPath() string {
    return "forwarding-vrf-summs"
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "forwarding-vrf-summ" {
        for _, c := range forwardingVrfSumms.ForwardingVrfSumm {
            if forwardingVrfSumms.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm{}
        forwardingVrfSumms.ForwardingVrfSumm = append(forwardingVrfSumms.ForwardingVrfSumm, child)
        return &forwardingVrfSumms.ForwardingVrfSumm[len(forwardingVrfSumms.ForwardingVrfSumm)-1]
    }
    return nil
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range forwardingVrfSumms.ForwardingVrfSumm {
        children[forwardingVrfSumms.ForwardingVrfSumm[i].GetSegmentPath()] = &forwardingVrfSumms.ForwardingVrfSumm[i]
    }
    return children
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetBundleName() string { return "cisco_ios_xe" }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetYangName() string { return "forwarding-vrf-summs" }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) SetParent(parent types.Entity) { forwardingVrfSumms.parent = parent }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetParent() types.Entity { return forwardingVrfSumms.parent }

func (forwardingVrfSumms *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms) GetParentYangName() string { return "forwarding" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm
// Summary of forwarding info for this VRF.
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm struct {
    parent types.Entity
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

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetFilter() yfilter.YFilter { return forwardingVrfSumm.YFilter }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) SetFilter(yf yfilter.YFilter) { forwardingVrfSumm.YFilter = yf }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "intfs-fwd-count" { return "IntfsFwdCount" }
    if yname == "local-lbls" { return "LocalLbls" }
    if yname == "pfxs" { return "Pfxs" }
    if yname == "nhs" { return "Nhs" }
    return ""
}

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetSegmentPath() string {
    return "forwarding-vrf-summ" + "[vrf-name='" + fmt.Sprintf("%v", forwardingVrfSumm.VrfName) + "']"
}

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pfxs" {
        return &forwardingVrfSumm.Pfxs
    }
    if childYangName == "nhs" {
        return &forwardingVrfSumm.Nhs
    }
    return nil
}

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pfxs"] = &forwardingVrfSumm.Pfxs
    children["nhs"] = &forwardingVrfSumm.Nhs
    return children
}

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = forwardingVrfSumm.VrfName
    leafs["intfs-fwd-count"] = forwardingVrfSumm.IntfsFwdCount
    leafs["local-lbls"] = forwardingVrfSumm.LocalLbls
    return leafs
}

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetBundleName() string { return "cisco_ios_xe" }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetYangName() string { return "forwarding-vrf-summ" }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) SetParent(parent types.Entity) { forwardingVrfSumm.parent = parent }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetParent() types.Entity { return forwardingVrfSumm.parent }

func (forwardingVrfSumm *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm) GetParentYangName() string { return "forwarding-vrf-summs" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs
// MPLS LDP forwarding prefix rewrite summary
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs struct {
    parent types.Entity
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

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetFilter() yfilter.YFilter { return pfxs.YFilter }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) SetFilter(yf yfilter.YFilter) { pfxs.YFilter = yf }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetGoName(yname string) string {
    if yname == "total-pfxs" { return "TotalPfxs" }
    if yname == "ecmp-pfxs" { return "EcmpPfxs" }
    if yname == "protected-pfxs" { return "ProtectedPfxs" }
    if yname == "labeled-pfxs-aggr" { return "LabeledPfxsAggr" }
    if yname == "labeled-pfxs-primary" { return "LabeledPfxsPrimary" }
    if yname == "labeled-pfxs-backup" { return "LabeledPfxsBackup" }
    return ""
}

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetSegmentPath() string {
    return "pfxs"
}

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "labeled-pfxs-aggr" {
        return &pfxs.LabeledPfxsAggr
    }
    if childYangName == "labeled-pfxs-primary" {
        return &pfxs.LabeledPfxsPrimary
    }
    if childYangName == "labeled-pfxs-backup" {
        return &pfxs.LabeledPfxsBackup
    }
    return nil
}

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["labeled-pfxs-aggr"] = &pfxs.LabeledPfxsAggr
    children["labeled-pfxs-primary"] = &pfxs.LabeledPfxsPrimary
    children["labeled-pfxs-backup"] = &pfxs.LabeledPfxsBackup
    return children
}

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-pfxs"] = pfxs.TotalPfxs
    leafs["ecmp-pfxs"] = pfxs.EcmpPfxs
    leafs["protected-pfxs"] = pfxs.ProtectedPfxs
    return leafs
}

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetBundleName() string { return "cisco_ios_xe" }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetYangName() string { return "pfxs" }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) SetParent(parent types.Entity) { pfxs.parent = parent }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetParent() types.Entity { return pfxs.parent }

func (pfxs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs) GetParentYangName() string { return "forwarding-vrf-summ" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr
// Labeled prefix count for all paths
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr struct {
    parent types.Entity
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

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetFilter() yfilter.YFilter { return labeledPfxsAggr.YFilter }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) SetFilter(yf yfilter.YFilter) { labeledPfxsAggr.YFilter = yf }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetGoName(yname string) string {
    if yname == "labeled-pfxs" { return "LabeledPfxs" }
    if yname == "labeled-pfxs-partial" { return "LabeledPfxsPartial" }
    if yname == "unlabeled-pfxs" { return "UnlabeledPfxs" }
    return ""
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetSegmentPath() string {
    return "labeled-pfxs-aggr"
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["labeled-pfxs"] = labeledPfxsAggr.LabeledPfxs
    leafs["labeled-pfxs-partial"] = labeledPfxsAggr.LabeledPfxsPartial
    leafs["unlabeled-pfxs"] = labeledPfxsAggr.UnlabeledPfxs
    return leafs
}

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetBundleName() string { return "cisco_ios_xe" }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetYangName() string { return "labeled-pfxs-aggr" }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) SetParent(parent types.Entity) { labeledPfxsAggr.parent = parent }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetParent() types.Entity { return labeledPfxsAggr.parent }

func (labeledPfxsAggr *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsAggr) GetParentYangName() string { return "pfxs" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary
// Labeled prefix count related to primary paths
// only
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary struct {
    parent types.Entity
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

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetFilter() yfilter.YFilter { return labeledPfxsPrimary.YFilter }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) SetFilter(yf yfilter.YFilter) { labeledPfxsPrimary.YFilter = yf }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetGoName(yname string) string {
    if yname == "labeled-pfxs" { return "LabeledPfxs" }
    if yname == "labeled-pfxs-partial" { return "LabeledPfxsPartial" }
    if yname == "unlabeled-pfxs" { return "UnlabeledPfxs" }
    return ""
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetSegmentPath() string {
    return "labeled-pfxs-primary"
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["labeled-pfxs"] = labeledPfxsPrimary.LabeledPfxs
    leafs["labeled-pfxs-partial"] = labeledPfxsPrimary.LabeledPfxsPartial
    leafs["unlabeled-pfxs"] = labeledPfxsPrimary.UnlabeledPfxs
    return leafs
}

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetBundleName() string { return "cisco_ios_xe" }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetYangName() string { return "labeled-pfxs-primary" }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) SetParent(parent types.Entity) { labeledPfxsPrimary.parent = parent }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetParent() types.Entity { return labeledPfxsPrimary.parent }

func (labeledPfxsPrimary *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsPrimary) GetParentYangName() string { return "pfxs" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup
// Labeled prefix count related to backup paths
// only
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup struct {
    parent types.Entity
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

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetFilter() yfilter.YFilter { return labeledPfxsBackup.YFilter }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) SetFilter(yf yfilter.YFilter) { labeledPfxsBackup.YFilter = yf }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetGoName(yname string) string {
    if yname == "labeled-pfxs" { return "LabeledPfxs" }
    if yname == "labeled-pfxs-partial" { return "LabeledPfxsPartial" }
    if yname == "unlabeled-pfxs" { return "UnlabeledPfxs" }
    return ""
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetSegmentPath() string {
    return "labeled-pfxs-backup"
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["labeled-pfxs"] = labeledPfxsBackup.LabeledPfxs
    leafs["labeled-pfxs-partial"] = labeledPfxsBackup.LabeledPfxsPartial
    leafs["unlabeled-pfxs"] = labeledPfxsBackup.UnlabeledPfxs
    return leafs
}

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetBundleName() string { return "cisco_ios_xe" }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetYangName() string { return "labeled-pfxs-backup" }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) SetParent(parent types.Entity) { labeledPfxsBackup.parent = parent }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetParent() types.Entity { return labeledPfxsBackup.parent }

func (labeledPfxsBackup *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Pfxs_LabeledPfxsBackup) GetParentYangName() string { return "pfxs" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs
// MPLS LDP forwarding rewrite next-hop/path summary
type MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs struct {
    parent types.Entity
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

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetFilter() yfilter.YFilter { return nhs.YFilter }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) SetFilter(yf yfilter.YFilter) { nhs.YFilter = yf }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetGoName(yname string) string {
    if yname == "total-paths" { return "TotalPaths" }
    if yname == "protected-paths" { return "ProtectedPaths" }
    if yname == "backup-paths" { return "BackupPaths" }
    if yname == "remote-backup-paths" { return "RemoteBackupPaths" }
    if yname == "labeled-paths" { return "LabeledPaths" }
    if yname == "labeled-backup-paths" { return "LabeledBackupPaths" }
    return ""
}

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetSegmentPath() string {
    return "nhs"
}

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-paths"] = nhs.TotalPaths
    leafs["protected-paths"] = nhs.ProtectedPaths
    leafs["backup-paths"] = nhs.BackupPaths
    leafs["remote-backup-paths"] = nhs.RemoteBackupPaths
    leafs["labeled-paths"] = nhs.LabeledPaths
    leafs["labeled-backup-paths"] = nhs.LabeledBackupPaths
    return leafs
}

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetBundleName() string { return "cisco_ios_xe" }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetYangName() string { return "nhs" }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) SetParent(parent types.Entity) { nhs.parent = parent }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetParent() types.Entity { return nhs.parent }

func (nhs *MplsLdp_MplsLdpState_Forwarding_ForwardingVrfSumms_ForwardingVrfSumm_Nhs) GetParentYangName() string { return "forwarding-vrf-summ" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail
// This leaf contain the individual LDP forwarding rewrite
// for a single prefix.
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail struct {
    parent types.Entity
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
    Paths []MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetFilter() yfilter.YFilter { return forwardingDetail.YFilter }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) SetFilter(yf yfilter.YFilter) { forwardingDetail.YFilter = yf }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "prefix" { return "Prefix" }
    if yname == "fwd-prefix" { return "FwdPrefix" }
    if yname == "table-id" { return "TableId" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "route" { return "Route" }
    if yname == "paths" { return "Paths" }
    return ""
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetSegmentPath() string {
    return "forwarding-detail" + "[vrf-name='" + fmt.Sprintf("%v", forwardingDetail.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", forwardingDetail.Prefix) + "']"
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "route" {
        return &forwardingDetail.Route
    }
    if childYangName == "paths" {
        for _, c := range forwardingDetail.Paths {
            if forwardingDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths{}
        forwardingDetail.Paths = append(forwardingDetail.Paths, child)
        return &forwardingDetail.Paths[len(forwardingDetail.Paths)-1]
    }
    return nil
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["route"] = &forwardingDetail.Route
    for i := range forwardingDetail.Paths {
        children[forwardingDetail.Paths[i].GetSegmentPath()] = &forwardingDetail.Paths[i]
    }
    return children
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = forwardingDetail.VrfName
    leafs["prefix"] = forwardingDetail.Prefix
    leafs["fwd-prefix"] = forwardingDetail.FwdPrefix
    leafs["table-id"] = forwardingDetail.TableId
    leafs["prefix-length"] = forwardingDetail.PrefixLength
    return leafs
}

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetBundleName() string { return "cisco_ios_xe" }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetYangName() string { return "forwarding-detail" }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) SetParent(parent types.Entity) { forwardingDetail.parent = parent }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetParent() types.Entity { return forwardingDetail.parent }

func (forwardingDetail *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail) GetParentYangName() string { return "forwarding" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route
// MPLS LDP Forwarding Route information
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route struct {
    parent types.Entity
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

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetFilter() yfilter.YFilter { return route.YFilter }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) SetFilter(yf yfilter.YFilter) { route.YFilter = yf }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetGoName(yname string) string {
    if yname == "version" { return "Version" }
    if yname == "priority" { return "Priority" }
    if yname == "source" { return "Source" }
    if yname == "type" { return "Type" }
    if yname == "metric" { return "Metric" }
    if yname == "is-local-vrf-leaked" { return "IsLocalVrfLeaked" }
    if yname == "routing-update-count" { return "RoutingUpdateCount" }
    if yname == "routing-update-timestamp" { return "RoutingUpdateTimestamp" }
    if yname == "routing-update-age" { return "RoutingUpdateAge" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "forwarding-update-count" { return "ForwardingUpdateCount" }
    if yname == "forwarding-update-timestamp" { return "ForwardingUpdateTimestamp" }
    if yname == "forwarding-update-age" { return "ForwardingUpdateAge" }
    return ""
}

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetSegmentPath() string {
    return "route"
}

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["version"] = route.Version
    leafs["priority"] = route.Priority
    leafs["source"] = route.Source
    leafs["type"] = route.Type
    leafs["metric"] = route.Metric
    leafs["is-local-vrf-leaked"] = route.IsLocalVrfLeaked
    leafs["routing-update-count"] = route.RoutingUpdateCount
    leafs["routing-update-timestamp"] = route.RoutingUpdateTimestamp
    leafs["routing-update-age"] = route.RoutingUpdateAge
    leafs["local-label"] = route.LocalLabel
    leafs["forwarding-update-count"] = route.ForwardingUpdateCount
    leafs["forwarding-update-timestamp"] = route.ForwardingUpdateTimestamp
    leafs["forwarding-update-age"] = route.ForwardingUpdateAge
    return leafs
}

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetBundleName() string { return "cisco_ios_xe" }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetYangName() string { return "route" }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) SetParent(parent types.Entity) { route.parent = parent }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetParent() types.Entity { return route.parent }

func (route *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Route) GetParentYangName() string { return "forwarding-detail" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths
// MPLS LDP Forwarding Path info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP Forwarding Path IP Routing information.
    Routing MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing

    // MPLS LDP Forwarding Path MPLS information.
    Mpls MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetFilter() yfilter.YFilter { return paths.YFilter }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) SetFilter(yf yfilter.YFilter) { paths.YFilter = yf }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetGoName(yname string) string {
    if yname == "routing" { return "Routing" }
    if yname == "mpls" { return "Mpls" }
    return ""
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetSegmentPath() string {
    return "paths"
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routing" {
        return &paths.Routing
    }
    if childYangName == "mpls" {
        return &paths.Mpls
    }
    return nil
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["routing"] = &paths.Routing
    children["mpls"] = &paths.Mpls
    return children
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetBundleName() string { return "cisco_ios_xe" }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetYangName() string { return "paths" }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) SetParent(parent types.Entity) { paths.parent = parent }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetParent() types.Entity { return paths.parent }

func (paths *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths) GetParentYangName() string { return "forwarding-detail" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing
// MPLS LDP Forwarding Path IP Routing information
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing struct {
    parent types.Entity
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

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetFilter() yfilter.YFilter { return routing.YFilter }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) SetFilter(yf yfilter.YFilter) { routing.YFilter = yf }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetGoName(yname string) string {
    if yname == "next-hop" { return "NextHop" }
    if yname == "remote-node-id" { return "RemoteNodeId" }
    if yname == "has-remote-lfa-bkup" { return "HasRemoteLfaBkup" }
    if yname == "interface" { return "Interface" }
    if yname == "nh-is-overriden" { return "NhIsOverriden" }
    if yname == "nexthop-id" { return "NexthopId" }
    if yname == "next-hop-table-id" { return "NextHopTableId" }
    if yname == "load-metric" { return "LoadMetric" }
    if yname == "path-id" { return "PathId" }
    if yname == "bkup-path-id" { return "BkupPathId" }
    if yname == "path-type" { return "PathType" }
    return ""
}

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetSegmentPath() string {
    return "routing"
}

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["next-hop"] = routing.NextHop
    leafs["remote-node-id"] = routing.RemoteNodeId
    leafs["has-remote-lfa-bkup"] = routing.HasRemoteLfaBkup
    leafs["interface"] = routing.Interface
    leafs["nh-is-overriden"] = routing.NhIsOverriden
    leafs["nexthop-id"] = routing.NexthopId
    leafs["next-hop-table-id"] = routing.NextHopTableId
    leafs["load-metric"] = routing.LoadMetric
    leafs["path-id"] = routing.PathId
    leafs["bkup-path-id"] = routing.BkupPathId
    leafs["path-type"] = routing.PathType
    return leafs
}

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetBundleName() string { return "cisco_ios_xe" }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetYangName() string { return "routing" }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) SetParent(parent types.Entity) { routing.parent = parent }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetParent() types.Entity { return routing.parent }

func (routing *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Routing) GetParentYangName() string { return "paths" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls
// MPLS LDP Forwarding Path MPLS information
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS nexthop info.
    MplsOutgoingInfo MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo

    // MPLS LDP Forwarding Path Remote LFA-FRR backup MPLS info.
    RemoteLfa MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetFilter() yfilter.YFilter { return mpls.YFilter }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) SetFilter(yf yfilter.YFilter) { mpls.YFilter = yf }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetGoName(yname string) string {
    if yname == "mpls-outgoing-info" { return "MplsOutgoingInfo" }
    if yname == "remote-lfa" { return "RemoteLfa" }
    return ""
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetSegmentPath() string {
    return "mpls"
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpls-outgoing-info" {
        return &mpls.MplsOutgoingInfo
    }
    if childYangName == "remote-lfa" {
        return &mpls.RemoteLfa
    }
    return nil
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpls-outgoing-info"] = &mpls.MplsOutgoingInfo
    children["remote-lfa"] = &mpls.RemoteLfa
    return children
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetBundleName() string { return "cisco_ios_xe" }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetYangName() string { return "mpls" }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) SetParent(parent types.Entity) { mpls.parent = parent }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetParent() types.Entity { return mpls.parent }

func (mpls *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls) GetParentYangName() string { return "paths" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo
// MPLS nexthop info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo struct {
    parent types.Entity
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

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetFilter() yfilter.YFilter { return mplsOutgoingInfo.YFilter }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) SetFilter(yf yfilter.YFilter) { mplsOutgoingInfo.YFilter = yf }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetGoName(yname string) string {
    if yname == "out-label" { return "OutLabel" }
    if yname == "out-label-type" { return "OutLabelType" }
    if yname == "out-label-owner" { return "OutLabelOwner" }
    if yname == "is-from-graceful-restartable-neighbor" { return "IsFromGracefulRestartableNeighbor" }
    if yname == "is-stale" { return "IsStale" }
    if yname == "nexthop-peer-ldp-ident" { return "NexthopPeerLdpIdent" }
    return ""
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetSegmentPath() string {
    return "mpls-outgoing-info"
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop-peer-ldp-ident" {
        return &mplsOutgoingInfo.NexthopPeerLdpIdent
    }
    return nil
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop-peer-ldp-ident"] = &mplsOutgoingInfo.NexthopPeerLdpIdent
    return children
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-label"] = mplsOutgoingInfo.OutLabel
    leafs["out-label-type"] = mplsOutgoingInfo.OutLabelType
    leafs["out-label-owner"] = mplsOutgoingInfo.OutLabelOwner
    leafs["is-from-graceful-restartable-neighbor"] = mplsOutgoingInfo.IsFromGracefulRestartableNeighbor
    leafs["is-stale"] = mplsOutgoingInfo.IsStale
    return leafs
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetBundleName() string { return "cisco_ios_xe" }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetYangName() string { return "mpls-outgoing-info" }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) SetParent(parent types.Entity) { mplsOutgoingInfo.parent = parent }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetParent() types.Entity { return mplsOutgoingInfo.parent }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo) GetParentYangName() string { return "mpls" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent
// Nexthop LDP peer
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent struct {
    parent types.Entity
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

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetFilter() yfilter.YFilter { return nexthopPeerLdpIdent.YFilter }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) SetFilter(yf yfilter.YFilter) { nexthopPeerLdpIdent.YFilter = yf }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    return ""
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetSegmentPath() string {
    return "nexthop-peer-ldp-ident"
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = nexthopPeerLdpIdent.LsrId
    leafs["label-space-id"] = nexthopPeerLdpIdent.LabelSpaceId
    return leafs
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetBundleName() string { return "cisco_ios_xe" }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetYangName() string { return "nexthop-peer-ldp-ident" }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) SetParent(parent types.Entity) { nexthopPeerLdpIdent.parent = parent }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetParent() types.Entity { return nexthopPeerLdpIdent.parent }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_MplsOutgoingInfo_NexthopPeerLdpIdent) GetParentYangName() string { return "mpls-outgoing-info" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa
// MPLS LDP Forwarding Path Remote LFA-FRR backup
// MPLS info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Whether path has remote LFA backup. The type is bool.
    HasRemoteLfaBkup interface{}

    // Remote LFA MPLS nexthop info.
    MplsOutgoingInfo MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetFilter() yfilter.YFilter { return remoteLfa.YFilter }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) SetFilter(yf yfilter.YFilter) { remoteLfa.YFilter = yf }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetGoName(yname string) string {
    if yname == "has-remote-lfa-bkup" { return "HasRemoteLfaBkup" }
    if yname == "mpls-outgoing-info" { return "MplsOutgoingInfo" }
    return ""
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetSegmentPath() string {
    return "remote-lfa"
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mpls-outgoing-info" {
        return &remoteLfa.MplsOutgoingInfo
    }
    return nil
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mpls-outgoing-info"] = &remoteLfa.MplsOutgoingInfo
    return children
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["has-remote-lfa-bkup"] = remoteLfa.HasRemoteLfaBkup
    return leafs
}

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetBundleName() string { return "cisco_ios_xe" }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetYangName() string { return "remote-lfa" }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) SetParent(parent types.Entity) { remoteLfa.parent = parent }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetParent() types.Entity { return remoteLfa.parent }

func (remoteLfa *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa) GetParentYangName() string { return "mpls" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo
// Remote LFA MPLS nexthop info
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo struct {
    parent types.Entity
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

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetFilter() yfilter.YFilter { return mplsOutgoingInfo.YFilter }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) SetFilter(yf yfilter.YFilter) { mplsOutgoingInfo.YFilter = yf }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetGoName(yname string) string {
    if yname == "out-label" { return "OutLabel" }
    if yname == "out-label-type" { return "OutLabelType" }
    if yname == "out-label-owner" { return "OutLabelOwner" }
    if yname == "is-from-graceful-restartable-neighbor" { return "IsFromGracefulRestartableNeighbor" }
    if yname == "is-stale" { return "IsStale" }
    if yname == "nexthop-peer-ldp-ident" { return "NexthopPeerLdpIdent" }
    return ""
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetSegmentPath() string {
    return "mpls-outgoing-info"
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nexthop-peer-ldp-ident" {
        return &mplsOutgoingInfo.NexthopPeerLdpIdent
    }
    return nil
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nexthop-peer-ldp-ident"] = &mplsOutgoingInfo.NexthopPeerLdpIdent
    return children
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-label"] = mplsOutgoingInfo.OutLabel
    leafs["out-label-type"] = mplsOutgoingInfo.OutLabelType
    leafs["out-label-owner"] = mplsOutgoingInfo.OutLabelOwner
    leafs["is-from-graceful-restartable-neighbor"] = mplsOutgoingInfo.IsFromGracefulRestartableNeighbor
    leafs["is-stale"] = mplsOutgoingInfo.IsStale
    return leafs
}

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetBundleName() string { return "cisco_ios_xe" }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetYangName() string { return "mpls-outgoing-info" }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) SetParent(parent types.Entity) { mplsOutgoingInfo.parent = parent }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetParent() types.Entity { return mplsOutgoingInfo.parent }

func (mplsOutgoingInfo *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo) GetParentYangName() string { return "remote-lfa" }

// MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent
// Nexthop LDP peer
type MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent struct {
    parent types.Entity
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

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetFilter() yfilter.YFilter { return nexthopPeerLdpIdent.YFilter }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) SetFilter(yf yfilter.YFilter) { nexthopPeerLdpIdent.YFilter = yf }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    return ""
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetSegmentPath() string {
    return "nexthop-peer-ldp-ident"
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = nexthopPeerLdpIdent.LsrId
    leafs["label-space-id"] = nexthopPeerLdpIdent.LabelSpaceId
    return leafs
}

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetBundleName() string { return "cisco_ios_xe" }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetYangName() string { return "nexthop-peer-ldp-ident" }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) SetParent(parent types.Entity) { nexthopPeerLdpIdent.parent = parent }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetParent() types.Entity { return nexthopPeerLdpIdent.parent }

func (nexthopPeerLdpIdent *MplsLdp_MplsLdpState_Forwarding_ForwardingDetail_Paths_Mpls_RemoteLfa_MplsOutgoingInfo_NexthopPeerLdpIdent) GetParentYangName() string { return "mpls-outgoing-info" }

// MplsLdp_MplsLdpState_Bindings
// The detailed LDP Bindings.
type MplsLdp_MplsLdpState_Bindings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This container holds the bindings specific to this VRF and AF.
    BindingsSumAfs MplsLdp_MplsLdpState_Bindings_BindingsSumAfs

    // This list contains the MPLS LDP Label Bindings for each IP Prefix. Label
    // bindings provide the local MPLS Label, a list of remote labels, any filters
    // affecting advertisment of that filter, and a list of neighbors to which the
    // label has been advertised. The type is slice of
    // MplsLdp_MplsLdpState_Bindings_Binding.
    Binding []MplsLdp_MplsLdpState_Bindings_Binding
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetFilter() yfilter.YFilter { return bindings.YFilter }

func (bindings *MplsLdp_MplsLdpState_Bindings) SetFilter(yf yfilter.YFilter) { bindings.YFilter = yf }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetGoName(yname string) string {
    if yname == "bindings-sum-afs" { return "BindingsSumAfs" }
    if yname == "binding" { return "Binding" }
    return ""
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetSegmentPath() string {
    return "bindings"
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bindings-sum-afs" {
        return &bindings.BindingsSumAfs
    }
    if childYangName == "binding" {
        for _, c := range bindings.Binding {
            if bindings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Bindings_Binding{}
        bindings.Binding = append(bindings.Binding, child)
        return &bindings.Binding[len(bindings.Binding)-1]
    }
    return nil
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bindings-sum-afs"] = &bindings.BindingsSumAfs
    for i := range bindings.Binding {
        children[bindings.Binding[i].GetSegmentPath()] = &bindings.Binding[i]
    }
    return children
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bindings *MplsLdp_MplsLdpState_Bindings) GetBundleName() string { return "cisco_ios_xe" }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetYangName() string { return "bindings" }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bindings *MplsLdp_MplsLdpState_Bindings) SetParent(parent types.Entity) { bindings.parent = parent }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetParent() types.Entity { return bindings.parent }

func (bindings *MplsLdp_MplsLdpState_Bindings) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Bindings_BindingsSumAfs
// This container holds the bindings specific to this VRF
// and AF.
type MplsLdp_MplsLdpState_Bindings_BindingsSumAfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Counters for the LDP Label Information Base for this VRF/AF. The type is
    // slice of MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf.
    BindingSumAf []MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetFilter() yfilter.YFilter { return bindingsSumAfs.YFilter }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) SetFilter(yf yfilter.YFilter) { bindingsSumAfs.YFilter = yf }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetGoName(yname string) string {
    if yname == "binding-sum-af" { return "BindingSumAf" }
    return ""
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetSegmentPath() string {
    return "bindings-sum-afs"
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "binding-sum-af" {
        for _, c := range bindingsSumAfs.BindingSumAf {
            if bindingsSumAfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf{}
        bindingsSumAfs.BindingSumAf = append(bindingsSumAfs.BindingSumAf, child)
        return &bindingsSumAfs.BindingSumAf[len(bindingsSumAfs.BindingSumAf)-1]
    }
    return nil
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range bindingsSumAfs.BindingSumAf {
        children[bindingsSumAfs.BindingSumAf[i].GetSegmentPath()] = &bindingsSumAfs.BindingSumAf[i]
    }
    return children
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetBundleName() string { return "cisco_ios_xe" }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetYangName() string { return "bindings-sum-afs" }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) SetParent(parent types.Entity) { bindingsSumAfs.parent = parent }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetParent() types.Entity { return bindingsSumAfs.parent }

func (bindingsSumAfs *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs) GetParentYangName() string { return "bindings" }

// MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf
// Counters for the LDP Label Information Base for this
// VRF/AF.
type MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf struct {
    parent types.Entity
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

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetFilter() yfilter.YFilter { return bindingSumAf.YFilter }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) SetFilter(yf yfilter.YFilter) { bindingSumAf.YFilter = yf }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "af-name" { return "AfName" }
    if yname == "binding-total" { return "BindingTotal" }
    if yname == "binding-no-route" { return "BindingNoRoute" }
    if yname == "binding-local-no-route" { return "BindingLocalNoRoute" }
    if yname == "binding-local" { return "BindingLocal" }
    if yname == "binding-local-null" { return "BindingLocalNull" }
    if yname == "binding-local-implicit-null" { return "BindingLocalImplicitNull" }
    if yname == "binding-local-explicit-null" { return "BindingLocalExplicitNull" }
    if yname == "binding-local-non-null" { return "BindingLocalNonNull" }
    if yname == "binding-local-oor" { return "BindingLocalOor" }
    if yname == "lowest-allocated-label" { return "LowestAllocatedLabel" }
    if yname == "highest-allocated-label" { return "HighestAllocatedLabel" }
    if yname == "binding-remote" { return "BindingRemote" }
    return ""
}

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetSegmentPath() string {
    return "binding-sum-af" + "[vrf-name='" + fmt.Sprintf("%v", bindingSumAf.VrfName) + "']" + "[af-name='" + fmt.Sprintf("%v", bindingSumAf.AfName) + "']"
}

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = bindingSumAf.VrfName
    leafs["af-name"] = bindingSumAf.AfName
    leafs["binding-total"] = bindingSumAf.BindingTotal
    leafs["binding-no-route"] = bindingSumAf.BindingNoRoute
    leafs["binding-local-no-route"] = bindingSumAf.BindingLocalNoRoute
    leafs["binding-local"] = bindingSumAf.BindingLocal
    leafs["binding-local-null"] = bindingSumAf.BindingLocalNull
    leafs["binding-local-implicit-null"] = bindingSumAf.BindingLocalImplicitNull
    leafs["binding-local-explicit-null"] = bindingSumAf.BindingLocalExplicitNull
    leafs["binding-local-non-null"] = bindingSumAf.BindingLocalNonNull
    leafs["binding-local-oor"] = bindingSumAf.BindingLocalOor
    leafs["lowest-allocated-label"] = bindingSumAf.LowestAllocatedLabel
    leafs["highest-allocated-label"] = bindingSumAf.HighestAllocatedLabel
    leafs["binding-remote"] = bindingSumAf.BindingRemote
    return leafs
}

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetBundleName() string { return "cisco_ios_xe" }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetYangName() string { return "binding-sum-af" }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) SetParent(parent types.Entity) { bindingSumAf.parent = parent }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetParent() types.Entity { return bindingSumAf.parent }

func (bindingSumAf *MplsLdp_MplsLdpState_Bindings_BindingsSumAfs_BindingSumAf) GetParentYangName() string { return "bindings-sum-afs" }

// MplsLdp_MplsLdpState_Bindings_Binding
// This list contains the MPLS LDP Label Bindings for each
// IP Prefix. Label bindings provide the local MPLS Label,
// a list of remote labels, any filters affecting
// advertisment of that filter, and a list of neighbors to
// which the label has been advertised.
type MplsLdp_MplsLdpState_Bindings_Binding struct {
    parent types.Entity
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
    RemoteBinding []MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding

    // Peers to which this entry is advertised. The type is slice of
    // MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo.
    PeersAdvertisedTo []MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetFilter() yfilter.YFilter { return binding.YFilter }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) SetFilter(yf yfilter.YFilter) { binding.YFilter = yf }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "prefix" { return "Prefix" }
    if yname == "fwd-prefix" { return "FwdPrefix" }
    if yname == "prefix-length" { return "PrefixLength" }
    if yname == "local-label" { return "LocalLabel" }
    if yname == "le-local-binding-revision" { return "LeLocalBindingRevision" }
    if yname == "le-local-label-state" { return "LeLocalLabelState" }
    if yname == "is-no-route" { return "IsNoRoute" }
    if yname == "label-oor" { return "LabelOor" }
    if yname == "advertise-prefix-filter" { return "AdvertisePrefixFilter" }
    if yname == "advertise-lsr-filter" { return "AdvertiseLsrFilter" }
    if yname == "config-enforced-local-label-value" { return "ConfigEnforcedLocalLabelValue" }
    if yname == "remote-binding" { return "RemoteBinding" }
    if yname == "peers-advertised-to" { return "PeersAdvertisedTo" }
    return ""
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetSegmentPath() string {
    return "binding" + "[vrf-name='" + fmt.Sprintf("%v", binding.VrfName) + "']" + "[prefix='" + fmt.Sprintf("%v", binding.Prefix) + "']"
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-binding" {
        for _, c := range binding.RemoteBinding {
            if binding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding{}
        binding.RemoteBinding = append(binding.RemoteBinding, child)
        return &binding.RemoteBinding[len(binding.RemoteBinding)-1]
    }
    if childYangName == "peers-advertised-to" {
        for _, c := range binding.PeersAdvertisedTo {
            if binding.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo{}
        binding.PeersAdvertisedTo = append(binding.PeersAdvertisedTo, child)
        return &binding.PeersAdvertisedTo[len(binding.PeersAdvertisedTo)-1]
    }
    return nil
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range binding.RemoteBinding {
        children[binding.RemoteBinding[i].GetSegmentPath()] = &binding.RemoteBinding[i]
    }
    for i := range binding.PeersAdvertisedTo {
        children[binding.PeersAdvertisedTo[i].GetSegmentPath()] = &binding.PeersAdvertisedTo[i]
    }
    return children
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = binding.VrfName
    leafs["prefix"] = binding.Prefix
    leafs["fwd-prefix"] = binding.FwdPrefix
    leafs["prefix-length"] = binding.PrefixLength
    leafs["local-label"] = binding.LocalLabel
    leafs["le-local-binding-revision"] = binding.LeLocalBindingRevision
    leafs["le-local-label-state"] = binding.LeLocalLabelState
    leafs["is-no-route"] = binding.IsNoRoute
    leafs["label-oor"] = binding.LabelOor
    leafs["advertise-prefix-filter"] = binding.AdvertisePrefixFilter
    leafs["advertise-lsr-filter"] = binding.AdvertiseLsrFilter
    leafs["config-enforced-local-label-value"] = binding.ConfigEnforcedLocalLabelValue
    return leafs
}

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetBundleName() string { return "cisco_ios_xe" }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetYangName() string { return "binding" }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) SetParent(parent types.Entity) { binding.parent = parent }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetParent() types.Entity { return binding.parent }

func (binding *MplsLdp_MplsLdpState_Bindings_Binding) GetParentYangName() string { return "bindings" }

// MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding
// MPLS LDP Remote Binding Information
type MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is the remote Label. The type is interface{} with range:
    // 0..4294967295.
    RemoteLabel interface{}

    // Is the entry stale. The type is bool.
    IsStale interface{}

    // Assigning peer.
    AssigningPeerLdpIdent MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetFilter() yfilter.YFilter { return remoteBinding.YFilter }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) SetFilter(yf yfilter.YFilter) { remoteBinding.YFilter = yf }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetGoName(yname string) string {
    if yname == "remote-label" { return "RemoteLabel" }
    if yname == "is-stale" { return "IsStale" }
    if yname == "assigning-peer-ldp-ident" { return "AssigningPeerLdpIdent" }
    return ""
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetSegmentPath() string {
    return "remote-binding"
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "assigning-peer-ldp-ident" {
        return &remoteBinding.AssigningPeerLdpIdent
    }
    return nil
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["assigning-peer-ldp-ident"] = &remoteBinding.AssigningPeerLdpIdent
    return children
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-label"] = remoteBinding.RemoteLabel
    leafs["is-stale"] = remoteBinding.IsStale
    return leafs
}

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetBundleName() string { return "cisco_ios_xe" }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetYangName() string { return "remote-binding" }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) SetParent(parent types.Entity) { remoteBinding.parent = parent }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetParent() types.Entity { return remoteBinding.parent }

func (remoteBinding *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding) GetParentYangName() string { return "binding" }

// MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent
// Assigning peer
type MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent struct {
    parent types.Entity
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

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetFilter() yfilter.YFilter { return assigningPeerLdpIdent.YFilter }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) SetFilter(yf yfilter.YFilter) { assigningPeerLdpIdent.YFilter = yf }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    return ""
}

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetSegmentPath() string {
    return "assigning-peer-ldp-ident"
}

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = assigningPeerLdpIdent.LsrId
    leafs["label-space-id"] = assigningPeerLdpIdent.LabelSpaceId
    return leafs
}

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetBundleName() string { return "cisco_ios_xe" }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetYangName() string { return "assigning-peer-ldp-ident" }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) SetParent(parent types.Entity) { assigningPeerLdpIdent.parent = parent }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetParent() types.Entity { return assigningPeerLdpIdent.parent }

func (assigningPeerLdpIdent *MplsLdp_MplsLdpState_Bindings_Binding_RemoteBinding_AssigningPeerLdpIdent) GetParentYangName() string { return "remote-binding" }

// MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo
// Peers to which this entry is advertised.
type MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo struct {
    parent types.Entity
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

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetFilter() yfilter.YFilter { return peersAdvertisedTo.YFilter }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) SetFilter(yf yfilter.YFilter) { peersAdvertisedTo.YFilter = yf }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetGoName(yname string) string {
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    return ""
}

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetSegmentPath() string {
    return "peers-advertised-to"
}

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lsr-id"] = peersAdvertisedTo.LsrId
    leafs["label-space-id"] = peersAdvertisedTo.LabelSpaceId
    return leafs
}

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetBundleName() string { return "cisco_ios_xe" }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetYangName() string { return "peers-advertised-to" }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) SetParent(parent types.Entity) { peersAdvertisedTo.parent = parent }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetParent() types.Entity { return peersAdvertisedTo.parent }

func (peersAdvertisedTo *MplsLdp_MplsLdpState_Bindings_Binding_PeersAdvertisedTo) GetParentYangName() string { return "binding" }

// MplsLdp_MplsLdpState_Neighbors
// The LDP Neighbors Information
type MplsLdp_MplsLdpState_Neighbors struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information on a particular LDP neighbor. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_Neighbor.
    Neighbor []MplsLdp_MplsLdpState_Neighbors_Neighbor

    // For this Neighbor, this is the list of adjacencies between the neighbor and
    // the local node. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_NbrAdjs.
    NbrAdjs []MplsLdp_MplsLdpState_Neighbors_NbrAdjs

    // MPLS LDP Statistics Information.
    StatsInfo MplsLdp_MplsLdpState_Neighbors_StatsInfo

    // LDP Backoff Information.
    Backoffs MplsLdp_MplsLdpState_Neighbors_Backoffs

    // This is the LDP NSR state for this neighbor.
    NsrNbrDetail MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetFilter() yfilter.YFilter { return neighbors.YFilter }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) SetFilter(yf yfilter.YFilter) { neighbors.YFilter = yf }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetGoName(yname string) string {
    if yname == "neighbor" { return "Neighbor" }
    if yname == "nbr-adjs" { return "NbrAdjs" }
    if yname == "stats-info" { return "StatsInfo" }
    if yname == "backoffs" { return "Backoffs" }
    if yname == "nsr-nbr-detail" { return "NsrNbrDetail" }
    return ""
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetSegmentPath() string {
    return "neighbors"
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "neighbor" {
        for _, c := range neighbors.Neighbor {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Neighbors_Neighbor{}
        neighbors.Neighbor = append(neighbors.Neighbor, child)
        return &neighbors.Neighbor[len(neighbors.Neighbor)-1]
    }
    if childYangName == "nbr-adjs" {
        for _, c := range neighbors.NbrAdjs {
            if neighbors.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Neighbors_NbrAdjs{}
        neighbors.NbrAdjs = append(neighbors.NbrAdjs, child)
        return &neighbors.NbrAdjs[len(neighbors.NbrAdjs)-1]
    }
    if childYangName == "stats-info" {
        return &neighbors.StatsInfo
    }
    if childYangName == "backoffs" {
        return &neighbors.Backoffs
    }
    if childYangName == "nsr-nbr-detail" {
        return &neighbors.NsrNbrDetail
    }
    return nil
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range neighbors.Neighbor {
        children[neighbors.Neighbor[i].GetSegmentPath()] = &neighbors.Neighbor[i]
    }
    for i := range neighbors.NbrAdjs {
        children[neighbors.NbrAdjs[i].GetSegmentPath()] = &neighbors.NbrAdjs[i]
    }
    children["stats-info"] = &neighbors.StatsInfo
    children["backoffs"] = &neighbors.Backoffs
    children["nsr-nbr-detail"] = &neighbors.NsrNbrDetail
    return children
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetBundleName() string { return "cisco_ios_xe" }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetYangName() string { return "neighbors" }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) SetParent(parent types.Entity) { neighbors.parent = parent }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetParent() types.Entity { return neighbors.parent }

func (neighbors *MplsLdp_MplsLdpState_Neighbors) GetParentYangName() string { return "mpls-ldp-state" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor
// Information on a particular LDP neighbor
type MplsLdp_MplsLdpState_Neighbors_Neighbor struct {
    parent types.Entity
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

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetFilter() yfilter.YFilter { return neighbor.YFilter }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) SetFilter(yf yfilter.YFilter) { neighbor.YFilter = yf }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "lsr-id" { return "LsrId" }
    if yname == "label-space-id" { return "LabelSpaceId" }
    if yname == "session-role" { return "SessionRole" }
    if yname == "session-prot-ver" { return "SessionProtVer" }
    if yname == "up-time-seconds" { return "UpTimeSeconds" }
    if yname == "nbr-path-vector-limit" { return "NbrPathVectorLimit" }
    if yname == "downstream-on-demand" { return "DownstreamOnDemand" }
    if yname == "peer-hold-time" { return "PeerHoldTime" }
    if yname == "peer-keep-alive-interval" { return "PeerKeepAliveInterval" }
    if yname == "peer-state" { return "PeerState" }
    if yname == "inbound-ipv4" { return "InboundIpv4" }
    if yname == "inbound-ipv6-filter" { return "InboundIpv6Filter" }
    if yname == "outbound-ipv4-filter" { return "OutboundIpv4Filter" }
    if yname == "outbound-ipv6-filter" { return "OutboundIpv6Filter" }
    if yname == "has-sp" { return "HasSp" }
    if yname == "sp-state" { return "SpState" }
    if yname == "sp-filter" { return "SpFilter" }
    if yname == "sp-has-duration" { return "SpHasDuration" }
    if yname == "sp-duration" { return "SpDuration" }
    if yname == "spht-running" { return "SphtRunning" }
    if yname == "spht-remaining" { return "SphtRemaining" }
    if yname == "bgp-advertisement-state" { return "BgpAdvertisementState" }
    if yname == "advertise-bgp-prefixes" { return "AdvertiseBgpPrefixes" }
    if yname == "client" { return "Client" }
    if yname == "duplicate-address" { return "DuplicateAddress" }
    if yname == "nbr-bound-address" { return "NbrBoundAddress" }
    if yname == "nbr-stats" { return "NbrStats" }
    if yname == "graceful-restart-adjacency" { return "GracefulRestartAdjacency" }
    if yname == "tcp-information" { return "TcpInformation" }
    if yname == "capabilities" { return "Capabilities" }
    return ""
}

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetSegmentPath() string {
    return "neighbor" + "[vrf-name='" + fmt.Sprintf("%v", neighbor.VrfName) + "']" + "[lsr-id='" + fmt.Sprintf("%v", neighbor.LsrId) + "']"
}

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nbr-stats" {
        return &neighbor.NbrStats
    }
    if childYangName == "graceful-restart-adjacency" {
        return &neighbor.GracefulRestartAdjacency
    }
    if childYangName == "tcp-information" {
        return &neighbor.TcpInformation
    }
    if childYangName == "capabilities" {
        return &neighbor.Capabilities
    }
    return nil
}

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nbr-stats"] = &neighbor.NbrStats
    children["graceful-restart-adjacency"] = &neighbor.GracefulRestartAdjacency
    children["tcp-information"] = &neighbor.TcpInformation
    children["capabilities"] = &neighbor.Capabilities
    return children
}

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = neighbor.VrfName
    leafs["lsr-id"] = neighbor.LsrId
    leafs["label-space-id"] = neighbor.LabelSpaceId
    leafs["session-role"] = neighbor.SessionRole
    leafs["session-prot-ver"] = neighbor.SessionProtVer
    leafs["up-time-seconds"] = neighbor.UpTimeSeconds
    leafs["nbr-path-vector-limit"] = neighbor.NbrPathVectorLimit
    leafs["downstream-on-demand"] = neighbor.DownstreamOnDemand
    leafs["peer-hold-time"] = neighbor.PeerHoldTime
    leafs["peer-keep-alive-interval"] = neighbor.PeerKeepAliveInterval
    leafs["peer-state"] = neighbor.PeerState
    leafs["inbound-ipv4"] = neighbor.InboundIpv4
    leafs["inbound-ipv6-filter"] = neighbor.InboundIpv6Filter
    leafs["outbound-ipv4-filter"] = neighbor.OutboundIpv4Filter
    leafs["outbound-ipv6-filter"] = neighbor.OutboundIpv6Filter
    leafs["has-sp"] = neighbor.HasSp
    leafs["sp-state"] = neighbor.SpState
    leafs["sp-filter"] = neighbor.SpFilter
    leafs["sp-has-duration"] = neighbor.SpHasDuration
    leafs["sp-duration"] = neighbor.SpDuration
    leafs["spht-running"] = neighbor.SphtRunning
    leafs["spht-remaining"] = neighbor.SphtRemaining
    leafs["bgp-advertisement-state"] = neighbor.BgpAdvertisementState
    leafs["advertise-bgp-prefixes"] = neighbor.AdvertiseBgpPrefixes
    leafs["client"] = neighbor.Client
    leafs["duplicate-address"] = neighbor.DuplicateAddress
    leafs["nbr-bound-address"] = neighbor.NbrBoundAddress
    return leafs
}

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetBundleName() string { return "cisco_ios_xe" }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetYangName() string { return "neighbor" }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) SetParent(parent types.Entity) { neighbor.parent = parent }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetParent() types.Entity { return neighbor.parent }

func (neighbor *MplsLdp_MplsLdpState_Neighbors_Neighbor) GetParentYangName() string { return "neighbors" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats
// Neighbor Statistics.
type MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats struct {
    parent types.Entity
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

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetFilter() yfilter.YFilter { return nbrStats.YFilter }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) SetFilter(yf yfilter.YFilter) { nbrStats.YFilter = yf }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetGoName(yname string) string {
    if yname == "ta-pies-sent" { return "TaPiesSent" }
    if yname == "ta-pies-rcvd" { return "TaPiesRcvd" }
    if yname == "num-of-nbr-ipv4-discovery" { return "NumOfNbrIpv4Discovery" }
    if yname == "num-of-nbr-ipv6-discovery" { return "NumOfNbrIpv6Discovery" }
    if yname == "num-of-nbr-ipv4-addresses" { return "NumOfNbrIpv4Addresses" }
    if yname == "num-of-nbr-ipv6-addresses" { return "NumOfNbrIpv6Addresses" }
    if yname == "num-of-nbr-ipv4-lbl" { return "NumOfNbrIpv4Lbl" }
    if yname == "num-of-nbr-ipv6-lbl" { return "NumOfNbrIpv6Lbl" }
    return ""
}

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetSegmentPath() string {
    return "nbr-stats"
}

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ta-pies-sent"] = nbrStats.TaPiesSent
    leafs["ta-pies-rcvd"] = nbrStats.TaPiesRcvd
    leafs["num-of-nbr-ipv4-discovery"] = nbrStats.NumOfNbrIpv4Discovery
    leafs["num-of-nbr-ipv6-discovery"] = nbrStats.NumOfNbrIpv6Discovery
    leafs["num-of-nbr-ipv4-addresses"] = nbrStats.NumOfNbrIpv4Addresses
    leafs["num-of-nbr-ipv6-addresses"] = nbrStats.NumOfNbrIpv6Addresses
    leafs["num-of-nbr-ipv4-lbl"] = nbrStats.NumOfNbrIpv4Lbl
    leafs["num-of-nbr-ipv6-lbl"] = nbrStats.NumOfNbrIpv6Lbl
    return leafs
}

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetBundleName() string { return "cisco_ios_xe" }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetYangName() string { return "nbr-stats" }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) SetParent(parent types.Entity) { nbrStats.parent = parent }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetParent() types.Entity { return nbrStats.parent }

func (nbrStats *MplsLdp_MplsLdpState_Neighbors_Neighbor_NbrStats) GetParentYangName() string { return "neighbor" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency
// This container holds the graceful restart information
// for this adjacency.
type MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency struct {
    parent types.Entity
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

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetFilter() yfilter.YFilter { return gracefulRestartAdjacency.YFilter }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) SetFilter(yf yfilter.YFilter) { gracefulRestartAdjacency.YFilter = yf }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetGoName(yname string) string {
    if yname == "is-graceful-restartable" { return "IsGracefulRestartable" }
    if yname == "reconnect-timeout" { return "ReconnectTimeout" }
    if yname == "recovery-time" { return "RecoveryTime" }
    if yname == "is-liveness-timer-running" { return "IsLivenessTimerRunning" }
    if yname == "liveness-timer-remaining-seconds" { return "LivenessTimerRemainingSeconds" }
    if yname == "is-recovery-timer-running" { return "IsRecoveryTimerRunning" }
    if yname == "recovery-timer-remaining-seconds" { return "RecoveryTimerRemainingSeconds" }
    if yname == "down-nbr-flap-count" { return "DownNbrFlapCount" }
    if yname == "down-nbr-down-reason" { return "DownNbrDownReason" }
    return ""
}

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetSegmentPath() string {
    return "graceful-restart-adjacency"
}

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-graceful-restartable"] = gracefulRestartAdjacency.IsGracefulRestartable
    leafs["reconnect-timeout"] = gracefulRestartAdjacency.ReconnectTimeout
    leafs["recovery-time"] = gracefulRestartAdjacency.RecoveryTime
    leafs["is-liveness-timer-running"] = gracefulRestartAdjacency.IsLivenessTimerRunning
    leafs["liveness-timer-remaining-seconds"] = gracefulRestartAdjacency.LivenessTimerRemainingSeconds
    leafs["is-recovery-timer-running"] = gracefulRestartAdjacency.IsRecoveryTimerRunning
    leafs["recovery-timer-remaining-seconds"] = gracefulRestartAdjacency.RecoveryTimerRemainingSeconds
    leafs["down-nbr-flap-count"] = gracefulRestartAdjacency.DownNbrFlapCount
    leafs["down-nbr-down-reason"] = gracefulRestartAdjacency.DownNbrDownReason
    return leafs
}

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetBundleName() string { return "cisco_ios_xe" }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetYangName() string { return "graceful-restart-adjacency" }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) SetParent(parent types.Entity) { gracefulRestartAdjacency.parent = parent }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetParent() types.Entity { return gracefulRestartAdjacency.parent }

func (gracefulRestartAdjacency *MplsLdp_MplsLdpState_Neighbors_Neighbor_GracefulRestartAdjacency) GetParentYangName() string { return "neighbor" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation
// MPLS LDP Neighbor TCP Information
type MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation struct {
    parent types.Entity
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

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetFilter() yfilter.YFilter { return tcpInformation.YFilter }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) SetFilter(yf yfilter.YFilter) { tcpInformation.YFilter = yf }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetGoName(yname string) string {
    if yname == "foreign-host" { return "ForeignHost" }
    if yname == "local-host" { return "LocalHost" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "is-md5-on" { return "IsMd5On" }
    if yname == "up-time" { return "UpTime" }
    return ""
}

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetSegmentPath() string {
    return "tcp-information"
}

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["foreign-host"] = tcpInformation.ForeignHost
    leafs["local-host"] = tcpInformation.LocalHost
    leafs["foreign-port"] = tcpInformation.ForeignPort
    leafs["local-port"] = tcpInformation.LocalPort
    leafs["is-md5-on"] = tcpInformation.IsMd5On
    leafs["up-time"] = tcpInformation.UpTime
    return leafs
}

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetBundleName() string { return "cisco_ios_xe" }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetYangName() string { return "tcp-information" }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) SetParent(parent types.Entity) { tcpInformation.parent = parent }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetParent() types.Entity { return tcpInformation.parent }

func (tcpInformation *MplsLdp_MplsLdpState_Neighbors_Neighbor_TcpInformation) GetParentYangName() string { return "neighbor" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities
// Capabilities sent to and received from neighbor
type MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of sent capabilities. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps.
    SentCaps []MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps

    // List of received capabilities. The type is slice of
    // MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps.
    ReceivedCaps []MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetFilter() yfilter.YFilter { return capabilities.YFilter }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) SetFilter(yf yfilter.YFilter) { capabilities.YFilter = yf }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetGoName(yname string) string {
    if yname == "sent-caps" { return "SentCaps" }
    if yname == "received-caps" { return "ReceivedCaps" }
    return ""
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetSegmentPath() string {
    return "capabilities"
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sent-caps" {
        for _, c := range capabilities.SentCaps {
            if capabilities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps{}
        capabilities.SentCaps = append(capabilities.SentCaps, child)
        return &capabilities.SentCaps[len(capabilities.SentCaps)-1]
    }
    if childYangName == "received-caps" {
        for _, c := range capabilities.ReceivedCaps {
            if capabilities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps{}
        capabilities.ReceivedCaps = append(capabilities.ReceivedCaps, child)
        return &capabilities.ReceivedCaps[len(capabilities.ReceivedCaps)-1]
    }
    return nil
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range capabilities.SentCaps {
        children[capabilities.SentCaps[i].GetSegmentPath()] = &capabilities.SentCaps[i]
    }
    for i := range capabilities.ReceivedCaps {
        children[capabilities.ReceivedCaps[i].GetSegmentPath()] = &capabilities.ReceivedCaps[i]
    }
    return children
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetBundleName() string { return "cisco_ios_xe" }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetYangName() string { return "capabilities" }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) SetParent(parent types.Entity) { capabilities.parent = parent }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetParent() types.Entity { return capabilities.parent }

func (capabilities *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities) GetParentYangName() string { return "neighbor" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps
// List of sent capabilities
type MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps struct {
    parent types.Entity
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

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetFilter() yfilter.YFilter { return sentCaps.YFilter }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) SetFilter(yf yfilter.YFilter) { sentCaps.YFilter = yf }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetGoName(yname string) string {
    if yname == "cap-type" { return "CapType" }
    if yname == "cap-des" { return "CapDes" }
    if yname == "capability-data-length" { return "CapabilityDataLength" }
    if yname == "capability-data" { return "CapabilityData" }
    return ""
}

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetSegmentPath() string {
    return "sent-caps" + "[cap-type='" + fmt.Sprintf("%v", sentCaps.CapType) + "']"
}

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cap-type"] = sentCaps.CapType
    leafs["cap-des"] = sentCaps.CapDes
    leafs["capability-data-length"] = sentCaps.CapabilityDataLength
    leafs["capability-data"] = sentCaps.CapabilityData
    return leafs
}

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetBundleName() string { return "cisco_ios_xe" }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetYangName() string { return "sent-caps" }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) SetParent(parent types.Entity) { sentCaps.parent = parent }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetParent() types.Entity { return sentCaps.parent }

func (sentCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_SentCaps) GetParentYangName() string { return "capabilities" }

// MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps
// List of received capabilities
type MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps struct {
    parent types.Entity
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

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetFilter() yfilter.YFilter { return receivedCaps.YFilter }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) SetFilter(yf yfilter.YFilter) { receivedCaps.YFilter = yf }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetGoName(yname string) string {
    if yname == "cap-type" { return "CapType" }
    if yname == "cap-des" { return "CapDes" }
    if yname == "capability-data-length" { return "CapabilityDataLength" }
    if yname == "capability-data" { return "CapabilityData" }
    return ""
}

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetSegmentPath() string {
    return "received-caps" + "[cap-type='" + fmt.Sprintf("%v", receivedCaps.CapType) + "']"
}

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cap-type"] = receivedCaps.CapType
    leafs["cap-des"] = receivedCaps.CapDes
    leafs["capability-data-length"] = receivedCaps.CapabilityDataLength
    leafs["capability-data"] = receivedCaps.CapabilityData
    return leafs
}

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetBundleName() string { return "cisco_ios_xe" }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetYangName() string { return "received-caps" }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) SetParent(parent types.Entity) { receivedCaps.parent = parent }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetParent() types.Entity { return receivedCaps.parent }

func (receivedCaps *MplsLdp_MplsLdpState_Neighbors_Neighbor_Capabilities_ReceivedCaps) GetParentYangName() string { return "capabilities" }

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
    parent types.Entity
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

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetFilter() yfilter.YFilter { return nbrAdjs.YFilter }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) SetFilter(yf yfilter.YFilter) { nbrAdjs.YFilter = yf }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "target-address" { return "TargetAddress" }
    if yname == "target-state" { return "TargetState" }
    return ""
}

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetSegmentPath() string {
    return "nbr-adjs"
}

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = nbrAdjs.Interface
    leafs["local-address"] = nbrAdjs.LocalAddress
    leafs["target-address"] = nbrAdjs.TargetAddress
    leafs["target-state"] = nbrAdjs.TargetState
    return leafs
}

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetBundleName() string { return "cisco_ios_xe" }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetYangName() string { return "nbr-adjs" }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) SetParent(parent types.Entity) { nbrAdjs.parent = parent }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetParent() types.Entity { return nbrAdjs.parent }

func (nbrAdjs *MplsLdp_MplsLdpState_Neighbors_NbrAdjs) GetParentYangName() string { return "neighbors" }

// MplsLdp_MplsLdpState_Neighbors_StatsInfo
// MPLS LDP Statistics Information
type MplsLdp_MplsLdpState_Neighbors_StatsInfo struct {
    parent types.Entity
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

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetFilter() yfilter.YFilter { return statsInfo.YFilter }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) SetFilter(yf yfilter.YFilter) { statsInfo.YFilter = yf }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetGoName(yname string) string {
    if yname == "discon-time" { return "DisconTime" }
    if yname == "session-attempts" { return "SessionAttempts" }
    if yname == "sess-reject-no-hello" { return "SessRejectNoHello" }
    if yname == "sess-rej-ad" { return "SessRejAd" }
    if yname == "sess-rej-max-pdu" { return "SessRejMaxPdu" }
    if yname == "sess-rej-lr" { return "SessRejLr" }
    if yname == "bad-ldpid" { return "BadLdpid" }
    if yname == "bad-pdu-len" { return "BadPduLen" }
    if yname == "bad-msg-len" { return "BadMsgLen" }
    if yname == "bad-tlv-len" { return "BadTlvLen" }
    if yname == "malformed-tlv-val" { return "MalformedTlvVal" }
    if yname == "keep-alive-exp" { return "KeepAliveExp" }
    if yname == "shutdown-notif-rec" { return "ShutdownNotifRec" }
    if yname == "shutdow-notif-sent" { return "ShutdowNotifSent" }
    if yname == "message-out" { return "MessageOut" }
    if yname == "message-in" { return "MessageIn" }
    return ""
}

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetSegmentPath() string {
    return "stats-info"
}

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "message-out" {
        return &statsInfo.MessageOut
    }
    if childYangName == "message-in" {
        return &statsInfo.MessageIn
    }
    return nil
}

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["message-out"] = &statsInfo.MessageOut
    children["message-in"] = &statsInfo.MessageIn
    return children
}

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discon-time"] = statsInfo.DisconTime
    leafs["session-attempts"] = statsInfo.SessionAttempts
    leafs["sess-reject-no-hello"] = statsInfo.SessRejectNoHello
    leafs["sess-rej-ad"] = statsInfo.SessRejAd
    leafs["sess-rej-max-pdu"] = statsInfo.SessRejMaxPdu
    leafs["sess-rej-lr"] = statsInfo.SessRejLr
    leafs["bad-ldpid"] = statsInfo.BadLdpid
    leafs["bad-pdu-len"] = statsInfo.BadPduLen
    leafs["bad-msg-len"] = statsInfo.BadMsgLen
    leafs["bad-tlv-len"] = statsInfo.BadTlvLen
    leafs["malformed-tlv-val"] = statsInfo.MalformedTlvVal
    leafs["keep-alive-exp"] = statsInfo.KeepAliveExp
    leafs["shutdown-notif-rec"] = statsInfo.ShutdownNotifRec
    leafs["shutdow-notif-sent"] = statsInfo.ShutdowNotifSent
    return leafs
}

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetBundleName() string { return "cisco_ios_xe" }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetYangName() string { return "stats-info" }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) SetParent(parent types.Entity) { statsInfo.parent = parent }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetParent() types.Entity { return statsInfo.parent }

func (statsInfo *MplsLdp_MplsLdpState_Neighbors_StatsInfo) GetParentYangName() string { return "neighbors" }

// MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut
// MPLS LDP message sent counters to this neighbor.
type MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut struct {
    parent types.Entity
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

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetFilter() yfilter.YFilter { return messageOut.YFilter }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) SetFilter(yf yfilter.YFilter) { messageOut.YFilter = yf }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetGoName(yname string) string {
    if yname == "total-count" { return "TotalCount" }
    if yname == "init-count" { return "InitCount" }
    if yname == "address-count" { return "AddressCount" }
    if yname == "address-withdraw-count" { return "AddressWithdrawCount" }
    if yname == "label-map-count" { return "LabelMapCount" }
    if yname == "label-withdraw-count" { return "LabelWithdrawCount" }
    if yname == "label-release-count" { return "LabelReleaseCount" }
    if yname == "label-request-count" { return "LabelRequestCount" }
    if yname == "label-abort-request-count" { return "LabelAbortRequestCount" }
    if yname == "notification-count" { return "NotificationCount" }
    if yname == "keep-alive-count" { return "KeepAliveCount" }
    if yname == "iccp-rg-conn-count" { return "IccpRgConnCount" }
    if yname == "iccp-rg-disconn-count" { return "IccpRgDisconnCount" }
    if yname == "iccp-rg-notif-count" { return "IccpRgNotifCount" }
    if yname == "iccp-rg-app-data-count" { return "IccpRgAppDataCount" }
    return ""
}

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetSegmentPath() string {
    return "message-out"
}

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-count"] = messageOut.TotalCount
    leafs["init-count"] = messageOut.InitCount
    leafs["address-count"] = messageOut.AddressCount
    leafs["address-withdraw-count"] = messageOut.AddressWithdrawCount
    leafs["label-map-count"] = messageOut.LabelMapCount
    leafs["label-withdraw-count"] = messageOut.LabelWithdrawCount
    leafs["label-release-count"] = messageOut.LabelReleaseCount
    leafs["label-request-count"] = messageOut.LabelRequestCount
    leafs["label-abort-request-count"] = messageOut.LabelAbortRequestCount
    leafs["notification-count"] = messageOut.NotificationCount
    leafs["keep-alive-count"] = messageOut.KeepAliveCount
    leafs["iccp-rg-conn-count"] = messageOut.IccpRgConnCount
    leafs["iccp-rg-disconn-count"] = messageOut.IccpRgDisconnCount
    leafs["iccp-rg-notif-count"] = messageOut.IccpRgNotifCount
    leafs["iccp-rg-app-data-count"] = messageOut.IccpRgAppDataCount
    return leafs
}

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetBundleName() string { return "cisco_ios_xe" }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetYangName() string { return "message-out" }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) SetParent(parent types.Entity) { messageOut.parent = parent }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetParent() types.Entity { return messageOut.parent }

func (messageOut *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageOut) GetParentYangName() string { return "stats-info" }

// MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn
// MPLS LDP message received counters from this
// neighbor.
type MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn struct {
    parent types.Entity
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

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetFilter() yfilter.YFilter { return messageIn.YFilter }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) SetFilter(yf yfilter.YFilter) { messageIn.YFilter = yf }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetGoName(yname string) string {
    if yname == "total-count" { return "TotalCount" }
    if yname == "init-count" { return "InitCount" }
    if yname == "address-count" { return "AddressCount" }
    if yname == "address-withdraw-count" { return "AddressWithdrawCount" }
    if yname == "label-map-count" { return "LabelMapCount" }
    if yname == "label-withdraw-count" { return "LabelWithdrawCount" }
    if yname == "label-release-count" { return "LabelReleaseCount" }
    if yname == "label-request-count" { return "LabelRequestCount" }
    if yname == "label-abort-request-count" { return "LabelAbortRequestCount" }
    if yname == "notification-count" { return "NotificationCount" }
    if yname == "keep-alive-count" { return "KeepAliveCount" }
    if yname == "iccp-rg-conn-count" { return "IccpRgConnCount" }
    if yname == "iccp-rg-disconn-count" { return "IccpRgDisconnCount" }
    if yname == "iccp-rg-notif-count" { return "IccpRgNotifCount" }
    if yname == "iccp-rg-app-data-count" { return "IccpRgAppDataCount" }
    return ""
}

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetSegmentPath() string {
    return "message-in"
}

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total-count"] = messageIn.TotalCount
    leafs["init-count"] = messageIn.InitCount
    leafs["address-count"] = messageIn.AddressCount
    leafs["address-withdraw-count"] = messageIn.AddressWithdrawCount
    leafs["label-map-count"] = messageIn.LabelMapCount
    leafs["label-withdraw-count"] = messageIn.LabelWithdrawCount
    leafs["label-release-count"] = messageIn.LabelReleaseCount
    leafs["label-request-count"] = messageIn.LabelRequestCount
    leafs["label-abort-request-count"] = messageIn.LabelAbortRequestCount
    leafs["notification-count"] = messageIn.NotificationCount
    leafs["keep-alive-count"] = messageIn.KeepAliveCount
    leafs["iccp-rg-conn-count"] = messageIn.IccpRgConnCount
    leafs["iccp-rg-disconn-count"] = messageIn.IccpRgDisconnCount
    leafs["iccp-rg-notif-count"] = messageIn.IccpRgNotifCount
    leafs["iccp-rg-app-data-count"] = messageIn.IccpRgAppDataCount
    return leafs
}

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetBundleName() string { return "cisco_ios_xe" }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetYangName() string { return "message-in" }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) SetParent(parent types.Entity) { messageIn.parent = parent }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetParent() types.Entity { return messageIn.parent }

func (messageIn *MplsLdp_MplsLdpState_Neighbors_StatsInfo_MessageIn) GetParentYangName() string { return "stats-info" }

// MplsLdp_MplsLdpState_Neighbors_Backoffs
// LDP Backoff Information
type MplsLdp_MplsLdpState_Neighbors_Backoffs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current neighbor backoff count in seconds. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    BackoffSeconds interface{}

    // Current neighbor backoff waiting count in seconds. The type is interface{}
    // with range: 0..4294967295. Units are seconds.
    WaitingSeconds interface{}
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetFilter() yfilter.YFilter { return backoffs.YFilter }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) SetFilter(yf yfilter.YFilter) { backoffs.YFilter = yf }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetGoName(yname string) string {
    if yname == "backoff-seconds" { return "BackoffSeconds" }
    if yname == "waiting-seconds" { return "WaitingSeconds" }
    return ""
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetSegmentPath() string {
    return "backoffs"
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["backoff-seconds"] = backoffs.BackoffSeconds
    leafs["waiting-seconds"] = backoffs.WaitingSeconds
    return leafs
}

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetBundleName() string { return "cisco_ios_xe" }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetYangName() string { return "backoffs" }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) SetParent(parent types.Entity) { backoffs.parent = parent }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetParent() types.Entity { return backoffs.parent }

func (backoffs *MplsLdp_MplsLdpState_Neighbors_Backoffs) GetParentYangName() string { return "neighbors" }

// MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail
// This is the LDP NSR state for this neighbor.
type MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail struct {
    parent types.Entity
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

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetFilter() yfilter.YFilter { return nsrNbrDetail.YFilter }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) SetFilter(yf yfilter.YFilter) { nsrNbrDetail.YFilter = yf }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetGoName(yname string) string {
    if yname == "nsr-state" { return "NsrState" }
    if yname == "nsr-nbr-sync-state" { return "NsrNbrSyncState" }
    if yname == "nsr-nbr-last-sync-error" { return "NsrNbrLastSyncError" }
    if yname == "nsr-nbr-last-sync-nack-reason" { return "NsrNbrLastSyncNackReason" }
    if yname == "nsr-nbr-pend-label-req-resps" { return "NsrNbrPendLabelReqResps" }
    if yname == "nsr-nbr-pend-label-withdraw-resps" { return "NsrNbrPendLabelWithdrawResps" }
    if yname == "nsr-nbr-pend-lcl-addr-withdraw-acks" { return "NsrNbrPendLclAddrWithdrawAcks" }
    if yname == "nsr-nbr-in-label-reqs-created" { return "NsrNbrInLabelReqsCreated" }
    if yname == "nsr-nbr-in-label-reqs-freed" { return "NsrNbrInLabelReqsFreed" }
    if yname == "nsr-nbr-in-label-withdraw-created" { return "NsrNbrInLabelWithdrawCreated" }
    if yname == "nsr-nbr-in-label-withdraw-freed" { return "NsrNbrInLabelWithdrawFreed" }
    if yname == "nsr-nbr-lcl-addr-withdraw-set" { return "NsrNbrLclAddrWithdrawSet" }
    if yname == "nsr-nbr-lcl-addr-withdraw-cleared" { return "NsrNbrLclAddrWithdrawCleared" }
    if yname == "nsr-nbr-xmit-ctxt-enq" { return "NsrNbrXmitCtxtEnq" }
    if yname == "nsr-nbr-xmit-ctxt-deq" { return "NsrNbrXmitCtxtDeq" }
    if yname == "path-vector-limit" { return "PathVectorLimit" }
    if yname == "nbr-sess" { return "NbrSess" }
    return ""
}

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetSegmentPath() string {
    return "nsr-nbr-detail"
}

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nbr-sess" {
        return &nsrNbrDetail.NbrSess
    }
    return nil
}

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nbr-sess"] = &nsrNbrDetail.NbrSess
    return children
}

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nsr-state"] = nsrNbrDetail.NsrState
    leafs["nsr-nbr-sync-state"] = nsrNbrDetail.NsrNbrSyncState
    leafs["nsr-nbr-last-sync-error"] = nsrNbrDetail.NsrNbrLastSyncError
    leafs["nsr-nbr-last-sync-nack-reason"] = nsrNbrDetail.NsrNbrLastSyncNackReason
    leafs["nsr-nbr-pend-label-req-resps"] = nsrNbrDetail.NsrNbrPendLabelReqResps
    leafs["nsr-nbr-pend-label-withdraw-resps"] = nsrNbrDetail.NsrNbrPendLabelWithdrawResps
    leafs["nsr-nbr-pend-lcl-addr-withdraw-acks"] = nsrNbrDetail.NsrNbrPendLclAddrWithdrawAcks
    leafs["nsr-nbr-in-label-reqs-created"] = nsrNbrDetail.NsrNbrInLabelReqsCreated
    leafs["nsr-nbr-in-label-reqs-freed"] = nsrNbrDetail.NsrNbrInLabelReqsFreed
    leafs["nsr-nbr-in-label-withdraw-created"] = nsrNbrDetail.NsrNbrInLabelWithdrawCreated
    leafs["nsr-nbr-in-label-withdraw-freed"] = nsrNbrDetail.NsrNbrInLabelWithdrawFreed
    leafs["nsr-nbr-lcl-addr-withdraw-set"] = nsrNbrDetail.NsrNbrLclAddrWithdrawSet
    leafs["nsr-nbr-lcl-addr-withdraw-cleared"] = nsrNbrDetail.NsrNbrLclAddrWithdrawCleared
    leafs["nsr-nbr-xmit-ctxt-enq"] = nsrNbrDetail.NsrNbrXmitCtxtEnq
    leafs["nsr-nbr-xmit-ctxt-deq"] = nsrNbrDetail.NsrNbrXmitCtxtDeq
    leafs["path-vector-limit"] = nsrNbrDetail.PathVectorLimit
    return leafs
}

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetBundleName() string { return "cisco_ios_xe" }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetYangName() string { return "nsr-nbr-detail" }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) SetParent(parent types.Entity) { nsrNbrDetail.parent = parent }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetParent() types.Entity { return nsrNbrDetail.parent }

func (nsrNbrDetail *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail) GetParentYangName() string { return "neighbors" }

// MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess
// This container holds session information about the
// sessions between these two neighbors.
type MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess struct {
    parent types.Entity
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

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetFilter() yfilter.YFilter { return nbrSess.YFilter }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) SetFilter(yf yfilter.YFilter) { nbrSess.YFilter = yf }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetGoName(yname string) string {
    if yname == "last-stat-change" { return "LastStatChange" }
    if yname == "state" { return "State" }
    if yname == "keep-alive-remain" { return "KeepAliveRemain" }
    if yname == "keep-alive-time" { return "KeepAliveTime" }
    if yname == "max-pdu" { return "MaxPdu" }
    if yname == "discon-time" { return "DisconTime" }
    if yname == "unknown-mess-err" { return "UnknownMessErr" }
    if yname == "unknown-tlv" { return "UnknownTlv" }
    return ""
}

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetSegmentPath() string {
    return "nbr-sess"
}

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["last-stat-change"] = nbrSess.LastStatChange
    leafs["state"] = nbrSess.State
    leafs["keep-alive-remain"] = nbrSess.KeepAliveRemain
    leafs["keep-alive-time"] = nbrSess.KeepAliveTime
    leafs["max-pdu"] = nbrSess.MaxPdu
    leafs["discon-time"] = nbrSess.DisconTime
    leafs["unknown-mess-err"] = nbrSess.UnknownMessErr
    leafs["unknown-tlv"] = nbrSess.UnknownTlv
    return leafs
}

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetBundleName() string { return "cisco_ios_xe" }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetYangName() string { return "nbr-sess" }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) SetParent(parent types.Entity) { nbrSess.parent = parent }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetParent() types.Entity { return nbrSess.parent }

func (nbrSess *MplsLdp_MplsLdpState_Neighbors_NsrNbrDetail_NbrSess) GetParentYangName() string { return "nsr-nbr-detail" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry contains a single range of labels represented by the configured
    // Upper and Lower Bounds pairs.  NOTE: there is NO corresponding LDP message
    // which relates to the information in this table, however, this table does
    // provide a way for a user to 'reserve' a generic label range.  NOTE:  The
    // ranges for a specific LDP Entity are UNIQUE and non-overlapping. The type
    // is slice of MplsLdp_MplsLdpState_LabelRanges_LabelRange.
    LabelRange []MplsLdp_MplsLdpState_LabelRanges_LabelRange
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetFilter() yfilter.YFilter { return labelRanges.YFilter }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) SetFilter(yf yfilter.YFilter) { labelRanges.YFilter = yf }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetGoName(yname string) string {
    if yname == "label-range" { return "LabelRange" }
    return ""
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetSegmentPath() string {
    return "label-ranges"
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-range" {
        for _, c := range labelRanges.LabelRange {
            if labelRanges.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpState_LabelRanges_LabelRange{}
        labelRanges.LabelRange = append(labelRanges.LabelRange, child)
        return &labelRanges.LabelRange[len(labelRanges.LabelRange)-1]
    }
    return nil
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelRanges.LabelRange {
        children[labelRanges.LabelRange[i].GetSegmentPath()] = &labelRanges.LabelRange[i]
    }
    return children
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetBundleName() string { return "cisco_ios_xe" }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetYangName() string { return "label-ranges" }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) SetParent(parent types.Entity) { labelRanges.parent = parent }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetParent() types.Entity { return labelRanges.parent }

func (labelRanges *MplsLdp_MplsLdpState_LabelRanges) GetParentYangName() string { return "mpls-ldp-state" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The minimum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    LrMin interface{}

    // This attribute is a key. The maximum label configured for this range. The
    // type is interface{} with range: 0..1048575.
    LrMax interface{}
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetFilter() yfilter.YFilter { return labelRange.YFilter }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) SetFilter(yf yfilter.YFilter) { labelRange.YFilter = yf }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetGoName(yname string) string {
    if yname == "lr-min" { return "LrMin" }
    if yname == "lr-max" { return "LrMax" }
    return ""
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetSegmentPath() string {
    return "label-range" + "[lr-min='" + fmt.Sprintf("%v", labelRange.LrMin) + "']" + "[lr-max='" + fmt.Sprintf("%v", labelRange.LrMax) + "']"
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lr-min"] = labelRange.LrMin
    leafs["lr-max"] = labelRange.LrMax
    return leafs
}

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetBundleName() string { return "cisco_ios_xe" }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetYangName() string { return "label-range" }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) SetParent(parent types.Entity) { labelRange.parent = parent }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetParent() types.Entity { return labelRange.parent }

func (labelRange *MplsLdp_MplsLdpState_LabelRanges_LabelRange) GetParentYangName() string { return "label-ranges" }

// MplsLdp_MplsLdpConfig
// MPLS LDP Configuration.
type MplsLdp_MplsLdpConfig struct {
    parent types.Entity
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

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetFilter() yfilter.YFilter { return mplsLdpConfig.YFilter }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) SetFilter(yf yfilter.YFilter) { mplsLdpConfig.YFilter = yf }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetGoName(yname string) string {
    if yname == "global-cfg" { return "GlobalCfg" }
    if yname == "nbr-table" { return "NbrTable" }
    if yname == "passwords" { return "Passwords" }
    if yname == "session" { return "Session" }
    if yname == "label-cfg" { return "LabelCfg" }
    if yname == "discovery" { return "Discovery" }
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "logging" { return "Logging" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "routing" { return "Routing" }
    if yname == "dual-stack" { return "DualStack" }
    return ""
}

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetSegmentPath() string {
    return "mpls-ldp-config"
}

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-cfg" {
        return &mplsLdpConfig.GlobalCfg
    }
    if childYangName == "nbr-table" {
        return &mplsLdpConfig.NbrTable
    }
    if childYangName == "passwords" {
        return &mplsLdpConfig.Passwords
    }
    if childYangName == "session" {
        return &mplsLdpConfig.Session
    }
    if childYangName == "label-cfg" {
        return &mplsLdpConfig.LabelCfg
    }
    if childYangName == "discovery" {
        return &mplsLdpConfig.Discovery
    }
    if childYangName == "graceful-restart" {
        return &mplsLdpConfig.GracefulRestart
    }
    if childYangName == "logging" {
        return &mplsLdpConfig.Logging
    }
    if childYangName == "interfaces" {
        return &mplsLdpConfig.Interfaces
    }
    if childYangName == "routing" {
        return &mplsLdpConfig.Routing
    }
    if childYangName == "dual-stack" {
        return &mplsLdpConfig.DualStack
    }
    return nil
}

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global-cfg"] = &mplsLdpConfig.GlobalCfg
    children["nbr-table"] = &mplsLdpConfig.NbrTable
    children["passwords"] = &mplsLdpConfig.Passwords
    children["session"] = &mplsLdpConfig.Session
    children["label-cfg"] = &mplsLdpConfig.LabelCfg
    children["discovery"] = &mplsLdpConfig.Discovery
    children["graceful-restart"] = &mplsLdpConfig.GracefulRestart
    children["logging"] = &mplsLdpConfig.Logging
    children["interfaces"] = &mplsLdpConfig.Interfaces
    children["routing"] = &mplsLdpConfig.Routing
    children["dual-stack"] = &mplsLdpConfig.DualStack
    return children
}

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetBundleName() string { return "cisco_ios_xe" }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetYangName() string { return "mpls-ldp-config" }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) SetParent(parent types.Entity) { mplsLdpConfig.parent = parent }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetParent() types.Entity { return mplsLdpConfig.parent }

func (mplsLdpConfig *MplsLdp_MplsLdpConfig) GetParentYangName() string { return "mpls-ldp" }

// MplsLdp_MplsLdpConfig_GlobalCfg
// This contains hold all MPLS LDP Configuration with Global
// scope. These values affect the entire LSR unless
// overiddden by a parameter with a more localized scope.
type MplsLdp_MplsLdpConfig_GlobalCfg struct {
    parent types.Entity
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
    RouterId []MplsLdp_MplsLdpConfig_GlobalCfg_RouterId

    // Configure session parameters. Session parameters effect the session between
    // LDP peers once the session has been established.
    Session MplsLdp_MplsLdpConfig_GlobalCfg_Session

    // This container holds the global per address family configuration.
    PerAf MplsLdp_MplsLdpConfig_GlobalCfg_PerAf
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetFilter() yfilter.YFilter { return globalCfg.YFilter }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) SetFilter(yf yfilter.YFilter) { globalCfg.YFilter = yf }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetGoName(yname string) string {
    if yname == "shutdown" { return "Shutdown" }
    if yname == "enable-nsr" { return "EnableNsr" }
    if yname == "disable-quick-start" { return "DisableQuickStart" }
    if yname == "loop-detection" { return "LoopDetection" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "dcsp-val" { return "DcspVal" }
    if yname == "high-priority" { return "HighPriority" }
    if yname == "seconds" { return "Seconds" }
    if yname == "disable-delay" { return "DisableDelay" }
    if yname == "seconds-delay-proc" { return "SecondsDelayProc" }
    if yname == "disable-delay-proc" { return "DisableDelayProc" }
    if yname == "protocol" { return "Protocol" }
    if yname == "init-sess-thresh" { return "InitSessThresh" }
    if yname == "router-id" { return "RouterId" }
    if yname == "session" { return "Session" }
    if yname == "per-af" { return "PerAf" }
    return ""
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetSegmentPath() string {
    return "global-cfg"
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "router-id" {
        for _, c := range globalCfg.RouterId {
            if globalCfg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_GlobalCfg_RouterId{}
        globalCfg.RouterId = append(globalCfg.RouterId, child)
        return &globalCfg.RouterId[len(globalCfg.RouterId)-1]
    }
    if childYangName == "session" {
        return &globalCfg.Session
    }
    if childYangName == "per-af" {
        return &globalCfg.PerAf
    }
    return nil
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalCfg.RouterId {
        children[globalCfg.RouterId[i].GetSegmentPath()] = &globalCfg.RouterId[i]
    }
    children["session"] = &globalCfg.Session
    children["per-af"] = &globalCfg.PerAf
    return children
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["shutdown"] = globalCfg.Shutdown
    leafs["enable-nsr"] = globalCfg.EnableNsr
    leafs["disable-quick-start"] = globalCfg.DisableQuickStart
    leafs["loop-detection"] = globalCfg.LoopDetection
    leafs["admin-status"] = globalCfg.AdminStatus
    leafs["dcsp-val"] = globalCfg.DcspVal
    leafs["high-priority"] = globalCfg.HighPriority
    leafs["seconds"] = globalCfg.Seconds
    leafs["disable-delay"] = globalCfg.DisableDelay
    leafs["seconds-delay-proc"] = globalCfg.SecondsDelayProc
    leafs["disable-delay-proc"] = globalCfg.DisableDelayProc
    leafs["protocol"] = globalCfg.Protocol
    leafs["init-sess-thresh"] = globalCfg.InitSessThresh
    return leafs
}

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetBundleName() string { return "cisco_ios_xe" }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetYangName() string { return "global-cfg" }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) SetParent(parent types.Entity) { globalCfg.parent = parent }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetParent() types.Entity { return globalCfg.parent }

func (globalCfg *MplsLdp_MplsLdpConfig_GlobalCfg) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_GlobalCfg_RouterId
// Configuration for LDP Router ID (LDP ID)
type MplsLdp_MplsLdpConfig_GlobalCfg_RouterId struct {
    parent types.Entity
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

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetFilter() yfilter.YFilter { return routerId.YFilter }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) SetFilter(yf yfilter.YFilter) { routerId.YFilter = yf }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "lsr-id-if" { return "LsrIdIf" }
    if yname == "lsr-id-ip" { return "LsrIdIp" }
    if yname == "force" { return "Force" }
    return ""
}

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetSegmentPath() string {
    return "router-id" + "[vrf-name='" + fmt.Sprintf("%v", routerId.VrfName) + "']"
}

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = routerId.VrfName
    leafs["lsr-id-if"] = routerId.LsrIdIf
    leafs["lsr-id-ip"] = routerId.LsrIdIp
    leafs["force"] = routerId.Force
    return leafs
}

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetBundleName() string { return "cisco_ios_xe" }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetYangName() string { return "router-id" }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) SetParent(parent types.Entity) { routerId.parent = parent }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetParent() types.Entity { return routerId.parent }

func (routerId *MplsLdp_MplsLdpConfig_GlobalCfg_RouterId) GetParentYangName() string { return "global-cfg" }

// MplsLdp_MplsLdpConfig_GlobalCfg_Session
// Configure session parameters. Session parameters effect
// the session between LDP peers once the session has been
// established.
type MplsLdp_MplsLdpConfig_GlobalCfg_Session struct {
    parent types.Entity
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
    DownstreamOnDemand []MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand

    // Configure Session Protection parameters.
    Protection MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetGoName(yname string) string {
    if yname == "backoff-init" { return "BackoffInit" }
    if yname == "backoff-max" { return "BackoffMax" }
    if yname == "seconds" { return "Seconds" }
    if yname == "infinite" { return "Infinite" }
    if yname == "downstream-on-demand" { return "DownstreamOnDemand" }
    if yname == "protection" { return "Protection" }
    return ""
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetSegmentPath() string {
    return "session"
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "downstream-on-demand" {
        for _, c := range session.DownstreamOnDemand {
            if session.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand{}
        session.DownstreamOnDemand = append(session.DownstreamOnDemand, child)
        return &session.DownstreamOnDemand[len(session.DownstreamOnDemand)-1]
    }
    if childYangName == "protection" {
        return &session.Protection
    }
    return nil
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range session.DownstreamOnDemand {
        children[session.DownstreamOnDemand[i].GetSegmentPath()] = &session.DownstreamOnDemand[i]
    }
    children["protection"] = &session.Protection
    return children
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["backoff-init"] = session.BackoffInit
    leafs["backoff-max"] = session.BackoffMax
    leafs["seconds"] = session.Seconds
    leafs["infinite"] = session.Infinite
    return leafs
}

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetBundleName() string { return "cisco_ios_xe" }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetYangName() string { return "session" }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetParent() types.Entity { return session.parent }

func (session *MplsLdp_MplsLdpConfig_GlobalCfg_Session) GetParentYangName() string { return "global-cfg" }

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
    parent types.Entity
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

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetFilter() yfilter.YFilter { return downstreamOnDemand.YFilter }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) SetFilter(yf yfilter.YFilter) { downstreamOnDemand.YFilter = yf }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "enabled" { return "Enabled" }
    if yname == "filter" { return "Filter" }
    return ""
}

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetSegmentPath() string {
    return "downstream-on-demand" + "[vrf-name='" + fmt.Sprintf("%v", downstreamOnDemand.VrfName) + "']"
}

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = downstreamOnDemand.VrfName
    leafs["enabled"] = downstreamOnDemand.Enabled
    leafs["filter"] = downstreamOnDemand.Filter
    return leafs
}

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetBundleName() string { return "cisco_ios_xe" }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetYangName() string { return "downstream-on-demand" }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) SetParent(parent types.Entity) { downstreamOnDemand.parent = parent }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetParent() types.Entity { return downstreamOnDemand.parent }

func (downstreamOnDemand *MplsLdp_MplsLdpConfig_GlobalCfg_Session_DownstreamOnDemand) GetParentYangName() string { return "session" }

// MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection
// Configure Session Protection parameters
type MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection struct {
    parent types.Entity
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

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetFilter() yfilter.YFilter { return protection.YFilter }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) SetFilter(yf yfilter.YFilter) { protection.YFilter = yf }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetGoName(yname string) string {
    if yname == "enable-prot" { return "EnableProt" }
    if yname == "peer-filter" { return "PeerFilter" }
    if yname == "seconds" { return "Seconds" }
    if yname == "inf" { return "Inf" }
    return ""
}

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetSegmentPath() string {
    return "protection"
}

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-prot"] = protection.EnableProt
    leafs["peer-filter"] = protection.PeerFilter
    leafs["seconds"] = protection.Seconds
    leafs["inf"] = protection.Inf
    return leafs
}

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetBundleName() string { return "cisco_ios_xe" }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetYangName() string { return "protection" }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) SetParent(parent types.Entity) { protection.parent = parent }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetParent() types.Entity { return protection.parent }

func (protection *MplsLdp_MplsLdpConfig_GlobalCfg_Session_Protection) GetParentYangName() string { return "session" }

// MplsLdp_MplsLdpConfig_GlobalCfg_PerAf
// This container holds the global per address family
// configuration.
type MplsLdp_MplsLdpConfig_GlobalCfg_PerAf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This container holds the global per address family configuration. The type
    // is slice of MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg.
    AfCfg []MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetFilter() yfilter.YFilter { return perAf.YFilter }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) SetFilter(yf yfilter.YFilter) { perAf.YFilter = yf }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetGoName(yname string) string {
    if yname == "af-cfg" { return "AfCfg" }
    return ""
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetSegmentPath() string {
    return "per-af"
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af-cfg" {
        for _, c := range perAf.AfCfg {
            if perAf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg{}
        perAf.AfCfg = append(perAf.AfCfg, child)
        return &perAf.AfCfg[len(perAf.AfCfg)-1]
    }
    return nil
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range perAf.AfCfg {
        children[perAf.AfCfg[i].GetSegmentPath()] = &perAf.AfCfg[i]
    }
    return children
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetBundleName() string { return "cisco_ios_xe" }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetYangName() string { return "per-af" }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) SetParent(parent types.Entity) { perAf.parent = parent }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetParent() types.Entity { return perAf.parent }

func (perAf *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf) GetParentYangName() string { return "global-cfg" }

// MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg
// This container holds the global per address family
// configuration.
type MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg struct {
    parent types.Entity
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

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetFilter() yfilter.YFilter { return afCfg.YFilter }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) SetFilter(yf yfilter.YFilter) { afCfg.YFilter = yf }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "af-name" { return "AfName" }
    if yname == "default-route" { return "DefaultRoute" }
    if yname == "ipaddr" { return "Ipaddr" }
    if yname == "interface" { return "Interface" }
    if yname == "implicit" { return "Implicit" }
    return ""
}

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetSegmentPath() string {
    return "af-cfg" + "[vrf-name='" + fmt.Sprintf("%v", afCfg.VrfName) + "']" + "[af-name='" + fmt.Sprintf("%v", afCfg.AfName) + "']"
}

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = afCfg.VrfName
    leafs["af-name"] = afCfg.AfName
    leafs["default-route"] = afCfg.DefaultRoute
    leafs["ipaddr"] = afCfg.Ipaddr
    leafs["interface"] = afCfg.Interface
    leafs["implicit"] = afCfg.Implicit
    return leafs
}

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetBundleName() string { return "cisco_ios_xe" }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetYangName() string { return "af-cfg" }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) SetParent(parent types.Entity) { afCfg.parent = parent }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetParent() types.Entity { return afCfg.parent }

func (afCfg *MplsLdp_MplsLdpConfig_GlobalCfg_PerAf_AfCfg) GetParentYangName() string { return "per-af" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry holds the configuration of a single neighbor identified by the
    // IP address of that neighbor. The type is slice of
    // MplsLdp_MplsLdpConfig_NbrTable_NbrCfg.
    NbrCfg []MplsLdp_MplsLdpConfig_NbrTable_NbrCfg
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetFilter() yfilter.YFilter { return nbrTable.YFilter }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) SetFilter(yf yfilter.YFilter) { nbrTable.YFilter = yf }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetGoName(yname string) string {
    if yname == "nbr-cfg" { return "NbrCfg" }
    return ""
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetSegmentPath() string {
    return "nbr-table"
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nbr-cfg" {
        for _, c := range nbrTable.NbrCfg {
            if nbrTable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_NbrTable_NbrCfg{}
        nbrTable.NbrCfg = append(nbrTable.NbrCfg, child)
        return &nbrTable.NbrCfg[len(nbrTable.NbrCfg)-1]
    }
    return nil
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nbrTable.NbrCfg {
        children[nbrTable.NbrCfg[i].GetSegmentPath()] = &nbrTable.NbrCfg[i]
    }
    return children
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetBundleName() string { return "cisco_ios_xe" }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetYangName() string { return "nbr-table" }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) SetParent(parent types.Entity) { nbrTable.parent = parent }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetParent() types.Entity { return nbrTable.parent }

func (nbrTable *MplsLdp_MplsLdpConfig_NbrTable) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_NbrTable_NbrCfg
// This entry holds the configuration of a single neighbor
// identified by the IP address of that neighbor.
type MplsLdp_MplsLdpConfig_NbrTable_NbrCfg struct {
    parent types.Entity
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

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetFilter() yfilter.YFilter { return nbrCfg.YFilter }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) SetFilter(yf yfilter.YFilter) { nbrCfg.YFilter = yf }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetGoName(yname string) string {
    if yname == "nbr-vrf" { return "NbrVrf" }
    if yname == "nbr-ip" { return "NbrIp" }
    if yname == "admin-status" { return "AdminStatus" }
    if yname == "implicit-withdraw" { return "ImplicitWithdraw" }
    if yname == "targeted" { return "Targeted" }
    if yname == "label-protocol" { return "LabelProtocol" }
    if yname == "label-binding-filter" { return "LabelBindingFilter" }
    if yname == "password" { return "Password" }
    return ""
}

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetSegmentPath() string {
    return "nbr-cfg" + "[nbr-vrf='" + fmt.Sprintf("%v", nbrCfg.NbrVrf) + "']" + "[nbr-ip='" + fmt.Sprintf("%v", nbrCfg.NbrIp) + "']"
}

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nbr-vrf"] = nbrCfg.NbrVrf
    leafs["nbr-ip"] = nbrCfg.NbrIp
    leafs["admin-status"] = nbrCfg.AdminStatus
    leafs["implicit-withdraw"] = nbrCfg.ImplicitWithdraw
    leafs["targeted"] = nbrCfg.Targeted
    leafs["label-protocol"] = nbrCfg.LabelProtocol
    leafs["label-binding-filter"] = nbrCfg.LabelBindingFilter
    leafs["password"] = nbrCfg.Password
    return leafs
}

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetBundleName() string { return "cisco_ios_xe" }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetYangName() string { return "nbr-cfg" }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) SetParent(parent types.Entity) { nbrCfg.parent = parent }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetParent() types.Entity { return nbrCfg.parent }

func (nbrCfg *MplsLdp_MplsLdpConfig_NbrTable_NbrCfg) GetParentYangName() string { return "nbr-table" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This holds the MPLS LDP password configuration for use with a single LDP
    // neighbor or group of LDP neighbors. The type is slice of
    // MplsLdp_MplsLdpConfig_Passwords_Password.
    Password []MplsLdp_MplsLdpConfig_Passwords_Password
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetFilter() yfilter.YFilter { return passwords.YFilter }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) SetFilter(yf yfilter.YFilter) { passwords.YFilter = yf }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetGoName(yname string) string {
    if yname == "password" { return "Password" }
    return ""
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetSegmentPath() string {
    return "passwords"
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "password" {
        for _, c := range passwords.Password {
            if passwords.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_Passwords_Password{}
        passwords.Password = append(passwords.Password, child)
        return &passwords.Password[len(passwords.Password)-1]
    }
    return nil
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range passwords.Password {
        children[passwords.Password[i].GetSegmentPath()] = &passwords.Password[i]
    }
    return children
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetBundleName() string { return "cisco_ios_xe" }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetYangName() string { return "passwords" }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) SetParent(parent types.Entity) { passwords.parent = parent }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetParent() types.Entity { return passwords.parent }

func (passwords *MplsLdp_MplsLdpConfig_Passwords) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_Passwords_Password
// This holds the MPLS LDP password configuration for use
// with a single LDP neighbor or group of LDP neighbors.
type MplsLdp_MplsLdpConfig_Passwords_Password struct {
    parent types.Entity
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

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetFilter() yfilter.YFilter { return password.YFilter }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) SetFilter(yf yfilter.YFilter) { password.YFilter = yf }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetGoName(yname string) string {
    if yname == "nbr-vrf" { return "NbrVrf" }
    if yname == "nbr-id" { return "NbrId" }
    if yname == "password-num" { return "PasswordNum" }
    if yname == "pass-required" { return "PassRequired" }
    if yname == "clear-pass" { return "ClearPass" }
    if yname == "encrypt-pass" { return "EncryptPass" }
    if yname == "keychain-pass" { return "KeychainPass" }
    return ""
}

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetSegmentPath() string {
    return "password" + "[nbr-vrf='" + fmt.Sprintf("%v", password.NbrVrf) + "']" + "[nbr-id='" + fmt.Sprintf("%v", password.NbrId) + "']" + "[password-num='" + fmt.Sprintf("%v", password.PasswordNum) + "']"
}

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nbr-vrf"] = password.NbrVrf
    leafs["nbr-id"] = password.NbrId
    leafs["password-num"] = password.PasswordNum
    leafs["pass-required"] = password.PassRequired
    leafs["clear-pass"] = password.ClearPass
    leafs["encrypt-pass"] = password.EncryptPass
    leafs["keychain-pass"] = password.KeychainPass
    return leafs
}

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetBundleName() string { return "cisco_ios_xe" }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetYangName() string { return "password" }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) SetParent(parent types.Entity) { password.parent = parent }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetParent() types.Entity { return password.parent }

func (password *MplsLdp_MplsLdpConfig_Passwords_Password) GetParentYangName() string { return "passwords" }

// MplsLdp_MplsLdpConfig_Session
// Configure session parameters
type MplsLdp_MplsLdpConfig_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initial session backoff time (seconds). The type is interface{} with range:
    // 0..4294967295. The default value is 15.
    Backoff interface{}

    // Session holdtime in seconds. The type is interface{} with range: 0..65535.
    Seconds interface{}

    // Ignore LDP session holdtime. The type is interface{}.
    Infinite interface{}
}

func (session *MplsLdp_MplsLdpConfig_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *MplsLdp_MplsLdpConfig_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *MplsLdp_MplsLdpConfig_Session) GetGoName(yname string) string {
    if yname == "backoff" { return "Backoff" }
    if yname == "seconds" { return "Seconds" }
    if yname == "infinite" { return "Infinite" }
    return ""
}

func (session *MplsLdp_MplsLdpConfig_Session) GetSegmentPath() string {
    return "session"
}

func (session *MplsLdp_MplsLdpConfig_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (session *MplsLdp_MplsLdpConfig_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (session *MplsLdp_MplsLdpConfig_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["backoff"] = session.Backoff
    leafs["seconds"] = session.Seconds
    leafs["infinite"] = session.Infinite
    return leafs
}

func (session *MplsLdp_MplsLdpConfig_Session) GetBundleName() string { return "cisco_ios_xe" }

func (session *MplsLdp_MplsLdpConfig_Session) GetYangName() string { return "session" }

func (session *MplsLdp_MplsLdpConfig_Session) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (session *MplsLdp_MplsLdpConfig_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (session *MplsLdp_MplsLdpConfig_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (session *MplsLdp_MplsLdpConfig_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *MplsLdp_MplsLdpConfig_Session) GetParent() types.Entity { return session.parent }

func (session *MplsLdp_MplsLdpConfig_Session) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_LabelCfg
// This container holds the label allocation and
// advertisement configuration for the LDP Label Information
// Base. These control what prefixes may be allocated and
// advertised to peers.
type MplsLdp_MplsLdpConfig_LabelCfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This is an allocation filter and advertisement filters for LDP labels in
    // this address family. The type is slice of
    // MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg.
    LabelAfCfg []MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetFilter() yfilter.YFilter { return labelCfg.YFilter }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) SetFilter(yf yfilter.YFilter) { labelCfg.YFilter = yf }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetGoName(yname string) string {
    if yname == "label-af-cfg" { return "LabelAfCfg" }
    return ""
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetSegmentPath() string {
    return "label-cfg"
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "label-af-cfg" {
        for _, c := range labelCfg.LabelAfCfg {
            if labelCfg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg{}
        labelCfg.LabelAfCfg = append(labelCfg.LabelAfCfg, child)
        return &labelCfg.LabelAfCfg[len(labelCfg.LabelAfCfg)-1]
    }
    return nil
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelCfg.LabelAfCfg {
        children[labelCfg.LabelAfCfg[i].GetSegmentPath()] = &labelCfg.LabelAfCfg[i]
    }
    return children
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetBundleName() string { return "cisco_ios_xe" }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetYangName() string { return "label-cfg" }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) SetParent(parent types.Entity) { labelCfg.parent = parent }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetParent() types.Entity { return labelCfg.parent }

func (labelCfg *MplsLdp_MplsLdpConfig_LabelCfg) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg
// This is an allocation filter and advertisement filters
// for LDP labels in this address family.
type MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg struct {
    parent types.Entity
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
    AdvtFilter []MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetFilter() yfilter.YFilter { return labelAfCfg.YFilter }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) SetFilter(yf yfilter.YFilter) { labelAfCfg.YFilter = yf }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "af-name" { return "AfName" }
    if yname == "prefix-filter" { return "PrefixFilter" }
    if yname == "host-route-enable" { return "HostRouteEnable" }
    if yname == "advt-filter" { return "AdvtFilter" }
    return ""
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetSegmentPath() string {
    return "label-af-cfg" + "[vrf-name='" + fmt.Sprintf("%v", labelAfCfg.VrfName) + "']" + "[af-name='" + fmt.Sprintf("%v", labelAfCfg.AfName) + "']"
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "advt-filter" {
        for _, c := range labelAfCfg.AdvtFilter {
            if labelAfCfg.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter{}
        labelAfCfg.AdvtFilter = append(labelAfCfg.AdvtFilter, child)
        return &labelAfCfg.AdvtFilter[len(labelAfCfg.AdvtFilter)-1]
    }
    return nil
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range labelAfCfg.AdvtFilter {
        children[labelAfCfg.AdvtFilter[i].GetSegmentPath()] = &labelAfCfg.AdvtFilter[i]
    }
    return children
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = labelAfCfg.VrfName
    leafs["af-name"] = labelAfCfg.AfName
    leafs["prefix-filter"] = labelAfCfg.PrefixFilter
    leafs["host-route-enable"] = labelAfCfg.HostRouteEnable
    return leafs
}

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetBundleName() string { return "cisco_ios_xe" }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetYangName() string { return "label-af-cfg" }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) SetParent(parent types.Entity) { labelAfCfg.parent = parent }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetParent() types.Entity { return labelAfCfg.parent }

func (labelAfCfg *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg) GetParentYangName() string { return "label-cfg" }

// MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter
// MPLS LDP Label advertisement filter restrictions.
type MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter struct {
    parent types.Entity
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

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetFilter() yfilter.YFilter { return advtFilter.YFilter }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) SetFilter(yf yfilter.YFilter) { advtFilter.YFilter = yf }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetGoName(yname string) string {
    if yname == "prefix-filter" { return "PrefixFilter" }
    if yname == "peer-filter" { return "PeerFilter" }
    if yname == "interface" { return "Interface" }
    if yname == "adv-label-cfg" { return "AdvLabelCfg" }
    return ""
}

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetSegmentPath() string {
    return "advt-filter" + "[prefix-filter='" + fmt.Sprintf("%v", advtFilter.PrefixFilter) + "']" + "[peer-filter='" + fmt.Sprintf("%v", advtFilter.PeerFilter) + "']" + "[interface='" + fmt.Sprintf("%v", advtFilter.Interface) + "']"
}

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["prefix-filter"] = advtFilter.PrefixFilter
    leafs["peer-filter"] = advtFilter.PeerFilter
    leafs["interface"] = advtFilter.Interface
    leafs["adv-label-cfg"] = advtFilter.AdvLabelCfg
    return leafs
}

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetBundleName() string { return "cisco_ios_xe" }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetYangName() string { return "advt-filter" }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) SetParent(parent types.Entity) { advtFilter.parent = parent }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetParent() types.Entity { return advtFilter.parent }

func (advtFilter *MplsLdp_MplsLdpConfig_LabelCfg_LabelAfCfg_AdvtFilter) GetParentYangName() string { return "label-af-cfg" }

// MplsLdp_MplsLdpConfig_Discovery
// LDP discovery
type MplsLdp_MplsLdpConfig_Discovery struct {
    parent types.Entity
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

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetFilter() yfilter.YFilter { return discovery.YFilter }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) SetFilter(yf yfilter.YFilter) { discovery.YFilter = yf }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetGoName(yname string) string {
    if yname == "instance-tlv" { return "InstanceTlv" }
    if yname == "link-hello" { return "LinkHello" }
    if yname == "targeted-hello" { return "TargetedHello" }
    if yname == "int-trans-addrs" { return "IntTransAddrs" }
    return ""
}

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetSegmentPath() string {
    return "discovery"
}

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-hello" {
        return &discovery.LinkHello
    }
    if childYangName == "targeted-hello" {
        return &discovery.TargetedHello
    }
    if childYangName == "int-trans-addrs" {
        return &discovery.IntTransAddrs
    }
    return nil
}

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-hello"] = &discovery.LinkHello
    children["targeted-hello"] = &discovery.TargetedHello
    children["int-trans-addrs"] = &discovery.IntTransAddrs
    return children
}

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["instance-tlv"] = discovery.InstanceTlv
    return leafs
}

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetBundleName() string { return "cisco_ios_xe" }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetYangName() string { return "discovery" }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) SetParent(parent types.Entity) { discovery.parent = parent }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetParent() types.Entity { return discovery.parent }

func (discovery *MplsLdp_MplsLdpConfig_Discovery) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_Discovery_LinkHello
// This container holds the parameters for the non-targeted
// link hello.
type MplsLdp_MplsLdpConfig_Discovery_LinkHello struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LDP discovery link hello holdtime in seconds. The type is interface{} with
    // range: 0..65535.
    Holdtime interface{}

    // LDP discovery link hello interval in seconds. The type is interface{} with
    // range: 0..65535.
    Interval interface{}
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetFilter() yfilter.YFilter { return linkHello.YFilter }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) SetFilter(yf yfilter.YFilter) { linkHello.YFilter = yf }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetGoName(yname string) string {
    if yname == "holdtime" { return "Holdtime" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetSegmentPath() string {
    return "link-hello"
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["holdtime"] = linkHello.Holdtime
    leafs["interval"] = linkHello.Interval
    return leafs
}

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetBundleName() string { return "cisco_ios_xe" }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetYangName() string { return "link-hello" }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) SetParent(parent types.Entity) { linkHello.parent = parent }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetParent() types.Entity { return linkHello.parent }

func (linkHello *MplsLdp_MplsLdpConfig_Discovery_LinkHello) GetParentYangName() string { return "discovery" }

// MplsLdp_MplsLdpConfig_Discovery_TargetedHello
// This container holds the parameters for the targeted
// link hello.
type MplsLdp_MplsLdpConfig_Discovery_TargetedHello struct {
    parent types.Entity
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

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetFilter() yfilter.YFilter { return targetedHello.YFilter }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) SetFilter(yf yfilter.YFilter) { targetedHello.YFilter = yf }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetGoName(yname string) string {
    if yname == "holdtime" { return "Holdtime" }
    if yname == "interval" { return "Interval" }
    if yname == "enable" { return "Enable" }
    if yname == "accept" { return "Accept" }
    return ""
}

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetSegmentPath() string {
    return "targeted-hello"
}

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accept" {
        return &targetedHello.Accept
    }
    return nil
}

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accept"] = &targetedHello.Accept
    return children
}

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["holdtime"] = targetedHello.Holdtime
    leafs["interval"] = targetedHello.Interval
    leafs["enable"] = targetedHello.Enable
    return leafs
}

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetBundleName() string { return "cisco_ios_xe" }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetYangName() string { return "targeted-hello" }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) SetParent(parent types.Entity) { targetedHello.parent = parent }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetParent() types.Entity { return targetedHello.parent }

func (targetedHello *MplsLdp_MplsLdpConfig_Discovery_TargetedHello) GetParentYangName() string { return "discovery" }

// MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept
// Enables router to respond to requests for targeted
// hello messages
type MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set to true if targeted hello messages may be accepted. The type is bool.
    Enable interface{}

    // Only respond to requests for targeted hello messages from sources matching
    // this filter. The filter type is device specific and could be an ACL, a
    // prefix list, or other mechanism. The type is string.
    SrcFilter interface{}
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetFilter() yfilter.YFilter { return accept.YFilter }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) SetFilter(yf yfilter.YFilter) { accept.YFilter = yf }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "src-filter" { return "SrcFilter" }
    return ""
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetSegmentPath() string {
    return "accept"
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = accept.Enable
    leafs["src-filter"] = accept.SrcFilter
    return leafs
}

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetBundleName() string { return "cisco_ios_xe" }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetYangName() string { return "accept" }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) SetParent(parent types.Entity) { accept.parent = parent }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetParent() types.Entity { return accept.parent }

func (accept *MplsLdp_MplsLdpConfig_Discovery_TargetedHello_Accept) GetParentYangName() string { return "targeted-hello" }

// MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs
// This list contains the per-interface transport
// addresses, which overide the global and default
// values.
type MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry contains the per-interface transport addresses, which overide
    // the global and default values. The type is slice of
    // MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr.
    IntTransAddr []MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetFilter() yfilter.YFilter { return intTransAddrs.YFilter }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) SetFilter(yf yfilter.YFilter) { intTransAddrs.YFilter = yf }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetGoName(yname string) string {
    if yname == "int-trans-addr" { return "IntTransAddr" }
    return ""
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetSegmentPath() string {
    return "int-trans-addrs"
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "int-trans-addr" {
        for _, c := range intTransAddrs.IntTransAddr {
            if intTransAddrs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr{}
        intTransAddrs.IntTransAddr = append(intTransAddrs.IntTransAddr, child)
        return &intTransAddrs.IntTransAddr[len(intTransAddrs.IntTransAddr)-1]
    }
    return nil
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intTransAddrs.IntTransAddr {
        children[intTransAddrs.IntTransAddr[i].GetSegmentPath()] = &intTransAddrs.IntTransAddr[i]
    }
    return children
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetBundleName() string { return "cisco_ios_xe" }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetYangName() string { return "int-trans-addrs" }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) SetParent(parent types.Entity) { intTransAddrs.parent = parent }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetParent() types.Entity { return intTransAddrs.parent }

func (intTransAddrs *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs) GetParentYangName() string { return "discovery" }

// MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr
// This entry contains the per-interface transport
// addresses, which overide the global and default
// values.
type MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr struct {
    parent types.Entity
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

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetFilter() yfilter.YFilter { return intTransAddr.YFilter }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) SetFilter(yf yfilter.YFilter) { intTransAddr.YFilter = yf }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "int-name" { return "IntName" }
    if yname == "trans-ip" { return "TransIp" }
    if yname == "trans-int" { return "TransInt" }
    return ""
}

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetSegmentPath() string {
    return "int-trans-addr" + "[af-name='" + fmt.Sprintf("%v", intTransAddr.AfName) + "']" + "[int-name='" + fmt.Sprintf("%v", intTransAddr.IntName) + "']"
}

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = intTransAddr.AfName
    leafs["int-name"] = intTransAddr.IntName
    leafs["trans-ip"] = intTransAddr.TransIp
    leafs["trans-int"] = intTransAddr.TransInt
    return leafs
}

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetBundleName() string { return "cisco_ios_xe" }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetYangName() string { return "int-trans-addr" }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) SetParent(parent types.Entity) { intTransAddr.parent = parent }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetParent() types.Entity { return intTransAddr.parent }

func (intTransAddr *MplsLdp_MplsLdpConfig_Discovery_IntTransAddrs_IntTransAddr) GetParentYangName() string { return "int-trans-addrs" }

// MplsLdp_MplsLdpConfig_GracefulRestart
// Configure LDP Graceful Restart
type MplsLdp_MplsLdpConfig_GracefulRestart struct {
    parent types.Entity
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
    Helper []MplsLdp_MplsLdpConfig_GracefulRestart_Helper
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetFilter() yfilter.YFilter { return gracefulRestart.YFilter }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) SetFilter(yf yfilter.YFilter) { gracefulRestart.YFilter = yf }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetGoName(yname string) string {
    if yname == "is-graceful-restartable" { return "IsGracefulRestartable" }
    if yname == "forwarding-holding" { return "ForwardingHolding" }
    if yname == "max-recovery" { return "MaxRecovery" }
    if yname == "nbr-liveness" { return "NbrLiveness" }
    if yname == "helper" { return "Helper" }
    return ""
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetSegmentPath() string {
    return "graceful-restart"
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper" {
        for _, c := range gracefulRestart.Helper {
            if gracefulRestart.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_GracefulRestart_Helper{}
        gracefulRestart.Helper = append(gracefulRestart.Helper, child)
        return &gracefulRestart.Helper[len(gracefulRestart.Helper)-1]
    }
    return nil
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range gracefulRestart.Helper {
        children[gracefulRestart.Helper[i].GetSegmentPath()] = &gracefulRestart.Helper[i]
    }
    return children
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-graceful-restartable"] = gracefulRestart.IsGracefulRestartable
    leafs["forwarding-holding"] = gracefulRestart.ForwardingHolding
    leafs["max-recovery"] = gracefulRestart.MaxRecovery
    leafs["nbr-liveness"] = gracefulRestart.NbrLiveness
    return leafs
}

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetBundleName() string { return "cisco_ios_xe" }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetYangName() string { return "graceful-restart" }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) SetParent(parent types.Entity) { gracefulRestart.parent = parent }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetParent() types.Entity { return gracefulRestart.parent }

func (gracefulRestart *MplsLdp_MplsLdpConfig_GracefulRestart) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_GracefulRestart_Helper
// This contains the filter name for peers for which this
// LSR will act as a graceful-restart helper.
type MplsLdp_MplsLdpConfig_GracefulRestart_Helper struct {
    parent types.Entity
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

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetFilter() yfilter.YFilter { return helper.YFilter }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) SetFilter(yf yfilter.YFilter) { helper.YFilter = yf }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetGoName(yname string) string {
    if yname == "helper-vrf" { return "HelperVrf" }
    if yname == "helper-filter" { return "HelperFilter" }
    return ""
}

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetSegmentPath() string {
    return "helper" + "[helper-vrf='" + fmt.Sprintf("%v", helper.HelperVrf) + "']" + "[helper-filter='" + fmt.Sprintf("%v", helper.HelperFilter) + "']"
}

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["helper-vrf"] = helper.HelperVrf
    leafs["helper-filter"] = helper.HelperFilter
    return leafs
}

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetBundleName() string { return "cisco_ios_xe" }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetYangName() string { return "helper" }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) SetParent(parent types.Entity) { helper.parent = parent }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetParent() types.Entity { return helper.parent }

func (helper *MplsLdp_MplsLdpConfig_GracefulRestart_Helper) GetParentYangName() string { return "graceful-restart" }

// MplsLdp_MplsLdpConfig_Logging
// Enable LDP logging
type MplsLdp_MplsLdpConfig_Logging struct {
    parent types.Entity
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

func (logging *MplsLdp_MplsLdpConfig_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *MplsLdp_MplsLdpConfig_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetGoName(yname string) string {
    if yname == "graceful-restart" { return "GracefulRestart" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "nsr" { return "Nsr" }
    if yname == "adjacency" { return "Adjacency" }
    if yname == "session-protection" { return "SessionProtection" }
    if yname == "password" { return "Password" }
    return ""
}

func (logging *MplsLdp_MplsLdpConfig_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *MplsLdp_MplsLdpConfig_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "password" {
        return &logging.Password
    }
    return nil
}

func (logging *MplsLdp_MplsLdpConfig_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["password"] = &logging.Password
    return children
}

func (logging *MplsLdp_MplsLdpConfig_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["graceful-restart"] = logging.GracefulRestart
    leafs["neighbor"] = logging.Neighbor
    leafs["nsr"] = logging.Nsr
    leafs["adjacency"] = logging.Adjacency
    leafs["session-protection"] = logging.SessionProtection
    return leafs
}

func (logging *MplsLdp_MplsLdpConfig_Logging) GetBundleName() string { return "cisco_ios_xe" }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetYangName() string { return "logging" }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (logging *MplsLdp_MplsLdpConfig_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetParent() types.Entity { return logging.parent }

func (logging *MplsLdp_MplsLdpConfig_Logging) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_Logging_Password
// Enable logging of password messages.
type MplsLdp_MplsLdpConfig_Logging_Password struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log MPLS LDP password configuration changes.
    ConfigMsg MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg

    // Log MPLS LDP password rollover messages.
    RolloverMsg MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetFilter() yfilter.YFilter { return password.YFilter }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) SetFilter(yf yfilter.YFilter) { password.YFilter = yf }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetGoName(yname string) string {
    if yname == "config-msg" { return "ConfigMsg" }
    if yname == "rollover-msg" { return "RolloverMsg" }
    return ""
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetSegmentPath() string {
    return "password"
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "config-msg" {
        return &password.ConfigMsg
    }
    if childYangName == "rollover-msg" {
        return &password.RolloverMsg
    }
    return nil
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["config-msg"] = &password.ConfigMsg
    children["rollover-msg"] = &password.RolloverMsg
    return children
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetBundleName() string { return "cisco_ios_xe" }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetYangName() string { return "password" }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) SetParent(parent types.Entity) { password.parent = parent }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetParent() types.Entity { return password.parent }

func (password *MplsLdp_MplsLdpConfig_Logging_Password) GetParentYangName() string { return "logging" }

// MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg
// Log MPLS LDP password configuration changes.
type MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log MPLS LDP password configuration changes. The type is bool.
    Enable interface{}

    // This is the number of messages per minute to limit the logging. A value of
    // 0 indicates no limits on the number of logged messages. The type is
    // interface{} with range: 0..4294967295.
    RateLimit interface{}
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetFilter() yfilter.YFilter { return configMsg.YFilter }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) SetFilter(yf yfilter.YFilter) { configMsg.YFilter = yf }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "rate-limit" { return "RateLimit" }
    return ""
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetSegmentPath() string {
    return "config-msg"
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = configMsg.Enable
    leafs["rate-limit"] = configMsg.RateLimit
    return leafs
}

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetBundleName() string { return "cisco_ios_xe" }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetYangName() string { return "config-msg" }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) SetParent(parent types.Entity) { configMsg.parent = parent }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetParent() types.Entity { return configMsg.parent }

func (configMsg *MplsLdp_MplsLdpConfig_Logging_Password_ConfigMsg) GetParentYangName() string { return "password" }

// MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg
// Log MPLS LDP password rollover messages.
type MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Log MPLS LDP password rollover messages. The type is bool.
    Enable interface{}

    // This is the number of messages per minute to limit the logging. A value of
    // 0 indicates no limits on the number of logged messages. The type is
    // interface{} with range: 0..4294967295.
    RateLimit interface{}
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetFilter() yfilter.YFilter { return rolloverMsg.YFilter }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) SetFilter(yf yfilter.YFilter) { rolloverMsg.YFilter = yf }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "rate-limit" { return "RateLimit" }
    return ""
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetSegmentPath() string {
    return "rollover-msg"
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rolloverMsg.Enable
    leafs["rate-limit"] = rolloverMsg.RateLimit
    return leafs
}

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetBundleName() string { return "cisco_ios_xe" }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetYangName() string { return "rollover-msg" }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) SetParent(parent types.Entity) { rolloverMsg.parent = parent }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetParent() types.Entity { return rolloverMsg.parent }

func (rolloverMsg *MplsLdp_MplsLdpConfig_Logging_Password_RolloverMsg) GetParentYangName() string { return "password" }

// MplsLdp_MplsLdpConfig_Interfaces
// MPLS LDP Interface configuration commands.
type MplsLdp_MplsLdpConfig_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP Interface configuration commands. Where a corresponding global
    // configuration command exists, the interface level command will take
    // precedence when configured. The type is slice of
    // MplsLdp_MplsLdpConfig_Interfaces_Interface.
    Interface []MplsLdp_MplsLdpConfig_Interfaces_Interface
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetBundleName() string { return "cisco_ios_xe" }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *MplsLdp_MplsLdpConfig_Interfaces) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_Interfaces_Interface
// MPLS LDP Interface configuration commands. Where a
// corresponding global configuration command exists, the
// interface level command will take precedence when
// configured.
type MplsLdp_MplsLdpConfig_Interfaces_Interface struct {
    parent types.Entity
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

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "interface" { return "Interface" }
    if yname == "link-hello-int" { return "LinkHelloInt" }
    if yname == "link-hello-hold" { return "LinkHelloHold" }
    if yname == "disable-quick-start-int" { return "DisableQuickStartInt" }
    if yname == "seconds" { return "Seconds" }
    if yname == "disable-delay" { return "DisableDelay" }
    if yname == "afs" { return "Afs" }
    return ""
}

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[vrf='" + fmt.Sprintf("%v", self.Vrf) + "']" + "[interface='" + fmt.Sprintf("%v", self.Interface) + "']"
}

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "afs" {
        return &self.Afs
    }
    return nil
}

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["afs"] = &self.Afs
    return children
}

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = self.Vrf
    leafs["interface"] = self.Interface
    leafs["link-hello-int"] = self.LinkHelloInt
    leafs["link-hello-hold"] = self.LinkHelloHold
    leafs["disable-quick-start-int"] = self.DisableQuickStartInt
    leafs["seconds"] = self.Seconds
    leafs["disable-delay"] = self.DisableDelay
    return leafs
}

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xe" }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *MplsLdp_MplsLdpConfig_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs
// Address Family specific operational data
type MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // MPLS LDP Operational data for this Address Family. The type is slice of
    // MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af.
    Af []MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetFilter() yfilter.YFilter { return afs.YFilter }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) SetFilter(yf yfilter.YFilter) { afs.YFilter = yf }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    return ""
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetSegmentPath() string {
    return "afs"
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "af" {
        for _, c := range afs.Af {
            if afs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af{}
        afs.Af = append(afs.Af, child)
        return &afs.Af[len(afs.Af)-1]
    }
    return nil
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range afs.Af {
        children[afs.Af[i].GetSegmentPath()] = &afs.Af[i]
    }
    return children
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetBundleName() string { return "cisco_ios_xe" }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetYangName() string { return "afs" }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) SetParent(parent types.Entity) { afs.parent = parent }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetParent() types.Entity { return afs.parent }

func (afs *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs) GetParentYangName() string { return "interface" }

// MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af
// MPLS LDP Operational data for this Address Family.
type MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af struct {
    parent types.Entity
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

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetFilter() yfilter.YFilter { return af.YFilter }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) SetFilter(yf yfilter.YFilter) { af.YFilter = yf }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "enable" { return "Enable" }
    if yname == "autoconfig-disable" { return "AutoconfigDisable" }
    if yname == "bgp-redist" { return "BgpRedist" }
    return ""
}

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetSegmentPath() string {
    return "af" + "[af-name='" + fmt.Sprintf("%v", af.AfName) + "']"
}

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp-redist" {
        return &af.BgpRedist
    }
    return nil
}

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp-redist"] = &af.BgpRedist
    return children
}

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = af.AfName
    leafs["enable"] = af.Enable
    leafs["autoconfig-disable"] = af.AutoconfigDisable
    return leafs
}

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetBundleName() string { return "cisco_ios_xe" }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetYangName() string { return "af" }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) SetParent(parent types.Entity) { af.parent = parent }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetParent() types.Entity { return af.parent }

func (af *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af) GetParentYangName() string { return "afs" }

// MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist
// MPLS LDP configuration for protocol
// redistribution. By default, redistribution of BGP
// routes is disabled. It can be enabled for all
// BGP routes or for a specific AS. Also it can be
// redistributed to all LDP peers or to a filtered
// group of peers.
type MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist struct {
    parent types.Entity
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

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetFilter() yfilter.YFilter { return bgpRedist.YFilter }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) SetFilter(yf yfilter.YFilter) { bgpRedist.YFilter = yf }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetGoName(yname string) string {
    if yname == "as-xx" { return "AsXx" }
    if yname == "as-yy" { return "AsYy" }
    if yname == "advertise-to" { return "AdvertiseTo" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetSegmentPath() string {
    return "bgp-redist"
}

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["as-xx"] = bgpRedist.AsXx
    leafs["as-yy"] = bgpRedist.AsYy
    leafs["advertise-to"] = bgpRedist.AdvertiseTo
    leafs["enable"] = bgpRedist.Enable
    return leafs
}

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetBundleName() string { return "cisco_ios_xe" }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetYangName() string { return "bgp-redist" }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) SetParent(parent types.Entity) { bgpRedist.parent = parent }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetParent() types.Entity { return bgpRedist.parent }

func (bgpRedist *MplsLdp_MplsLdpConfig_Interfaces_Interface_Afs_Af_BgpRedist) GetParentYangName() string { return "af" }

// MplsLdp_MplsLdpConfig_Routing
// This containter provides the MPLS LDP config for routing
// protocols from which it can obtain addresses to
// associate with labels.
type MplsLdp_MplsLdpConfig_Routing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This entry provides the MPLS LDP config for this routing instance. The type
    // is slice of MplsLdp_MplsLdpConfig_Routing_RoutingInst.
    RoutingInst []MplsLdp_MplsLdpConfig_Routing_RoutingInst
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetFilter() yfilter.YFilter { return routing.YFilter }

func (routing *MplsLdp_MplsLdpConfig_Routing) SetFilter(yf yfilter.YFilter) { routing.YFilter = yf }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetGoName(yname string) string {
    if yname == "routing-inst" { return "RoutingInst" }
    return ""
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetSegmentPath() string {
    return "routing"
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "routing-inst" {
        for _, c := range routing.RoutingInst {
            if routing.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := MplsLdp_MplsLdpConfig_Routing_RoutingInst{}
        routing.RoutingInst = append(routing.RoutingInst, child)
        return &routing.RoutingInst[len(routing.RoutingInst)-1]
    }
    return nil
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range routing.RoutingInst {
        children[routing.RoutingInst[i].GetSegmentPath()] = &routing.RoutingInst[i]
    }
    return children
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (routing *MplsLdp_MplsLdpConfig_Routing) GetBundleName() string { return "cisco_ios_xe" }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetYangName() string { return "routing" }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (routing *MplsLdp_MplsLdpConfig_Routing) SetParent(parent types.Entity) { routing.parent = parent }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetParent() types.Entity { return routing.parent }

func (routing *MplsLdp_MplsLdpConfig_Routing) GetParentYangName() string { return "mpls-ldp-config" }

// MplsLdp_MplsLdpConfig_Routing_RoutingInst
// This entry provides the MPLS LDP config for this
// routing instance.
type MplsLdp_MplsLdpConfig_Routing_RoutingInst struct {
    parent types.Entity
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

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetFilter() yfilter.YFilter { return routingInst.YFilter }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) SetFilter(yf yfilter.YFilter) { routingInst.YFilter = yf }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetGoName(yname string) string {
    if yname == "routing-inst-name" { return "RoutingInstName" }
    if yname == "autoconfig-enable" { return "AutoconfigEnable" }
    if yname == "area-id" { return "AreaId" }
    if yname == "level-id" { return "LevelId" }
    if yname == "sync" { return "Sync" }
    return ""
}

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetSegmentPath() string {
    return "routing-inst" + "[routing-inst-name='" + fmt.Sprintf("%v", routingInst.RoutingInstName) + "']"
}

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["routing-inst-name"] = routingInst.RoutingInstName
    leafs["autoconfig-enable"] = routingInst.AutoconfigEnable
    leafs["area-id"] = routingInst.AreaId
    leafs["level-id"] = routingInst.LevelId
    leafs["sync"] = routingInst.Sync
    return leafs
}

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetBundleName() string { return "cisco_ios_xe" }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetYangName() string { return "routing-inst" }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) SetParent(parent types.Entity) { routingInst.parent = parent }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetParent() types.Entity { return routingInst.parent }

func (routingInst *MplsLdp_MplsLdpConfig_Routing_RoutingInst) GetParentYangName() string { return "routing" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Wait time in seconds (0 indicates no preference). The type is interface{}
    // with range: 0..60.
    MaxWait interface{}

    // This contains the filter name for peers where IPv4 connections are
    // preferred over IPv6 connections. The filter type is device specific and
    // could be an ACL, a prefix list, or other mechanism. The type is string.
    PreferIpv4Peers interface{}
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetFilter() yfilter.YFilter { return dualStack.YFilter }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) SetFilter(yf yfilter.YFilter) { dualStack.YFilter = yf }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetGoName(yname string) string {
    if yname == "max-wait" { return "MaxWait" }
    if yname == "prefer-ipv4-peers" { return "PreferIpv4Peers" }
    return ""
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetSegmentPath() string {
    return "dual-stack"
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-wait"] = dualStack.MaxWait
    leafs["prefer-ipv4-peers"] = dualStack.PreferIpv4Peers
    return leafs
}

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetBundleName() string { return "cisco_ios_xe" }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetYangName() string { return "dual-stack" }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) SetParent(parent types.Entity) { dualStack.parent = parent }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetParent() types.Entity { return dualStack.parent }

func (dualStack *MplsLdp_MplsLdpConfig_DualStack) GetParentYangName() string { return "mpls-ldp-config" }

// ClearMsgCounters
// This RPC clears the LDP message counters for either a single
// neighbor or for all neighbors.
type ClearMsgCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearMsgCounters_Input

    
    Output ClearMsgCounters_Output
}

func (clearMsgCounters *ClearMsgCounters) GetFilter() yfilter.YFilter { return clearMsgCounters.YFilter }

func (clearMsgCounters *ClearMsgCounters) SetFilter(yf yfilter.YFilter) { clearMsgCounters.YFilter = yf }

func (clearMsgCounters *ClearMsgCounters) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (clearMsgCounters *ClearMsgCounters) GetSegmentPath() string {
    return "Cisco-IOS-XE-mpls-ldp:clear-msg-counters"
}

func (clearMsgCounters *ClearMsgCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearMsgCounters.Input
    }
    if childYangName == "output" {
        return &clearMsgCounters.Output
    }
    return nil
}

func (clearMsgCounters *ClearMsgCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearMsgCounters.Input
    children["output"] = &clearMsgCounters.Output
    return children
}

func (clearMsgCounters *ClearMsgCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearMsgCounters *ClearMsgCounters) GetBundleName() string { return "cisco_ios_xe" }

func (clearMsgCounters *ClearMsgCounters) GetYangName() string { return "clear-msg-counters" }

func (clearMsgCounters *ClearMsgCounters) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (clearMsgCounters *ClearMsgCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (clearMsgCounters *ClearMsgCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (clearMsgCounters *ClearMsgCounters) SetParent(parent types.Entity) { clearMsgCounters.parent = parent }

func (clearMsgCounters *ClearMsgCounters) GetParent() types.Entity { return clearMsgCounters.parent }

func (clearMsgCounters *ClearMsgCounters) GetParentYangName() string { return "Cisco-IOS-XE-mpls-ldp" }

// ClearMsgCounters_Input
type ClearMsgCounters_Input struct {
    parent types.Entity
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

func (input *ClearMsgCounters_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearMsgCounters_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearMsgCounters_Input) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "nbr-ip" { return "NbrIp" }
    if yname == "all" { return "All" }
    return ""
}

func (input *ClearMsgCounters_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearMsgCounters_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *ClearMsgCounters_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *ClearMsgCounters_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = input.VrfName
    leafs["nbr-ip"] = input.NbrIp
    leafs["all"] = input.All
    return leafs
}

func (input *ClearMsgCounters_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *ClearMsgCounters_Input) GetYangName() string { return "input" }

func (input *ClearMsgCounters_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *ClearMsgCounters_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *ClearMsgCounters_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *ClearMsgCounters_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearMsgCounters_Input) GetParent() types.Entity { return input.parent }

func (input *ClearMsgCounters_Input) GetParentYangName() string { return "clear-msg-counters" }

// ClearMsgCounters_Output
type ClearMsgCounters_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Return status will be 'OK' on success or an explanation string on failure.
    // The type is string.
    Status interface{}
}

func (output *ClearMsgCounters_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *ClearMsgCounters_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *ClearMsgCounters_Output) GetGoName(yname string) string {
    if yname == "status" { return "Status" }
    return ""
}

func (output *ClearMsgCounters_Output) GetSegmentPath() string {
    return "output"
}

func (output *ClearMsgCounters_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *ClearMsgCounters_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *ClearMsgCounters_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["status"] = output.Status
    return leafs
}

func (output *ClearMsgCounters_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *ClearMsgCounters_Output) GetYangName() string { return "output" }

func (output *ClearMsgCounters_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *ClearMsgCounters_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *ClearMsgCounters_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *ClearMsgCounters_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *ClearMsgCounters_Output) GetParent() types.Entity { return output.parent }

func (output *ClearMsgCounters_Output) GetParentYangName() string { return "clear-msg-counters" }

// RestartNeighbor
// This RPC restarts a single LDP session or all LDP sessions,
// but does not restart the LDP process itself, if the device
// supports that capability.
type RestartNeighbor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input RestartNeighbor_Input

    
    Output RestartNeighbor_Output
}

func (restartNeighbor *RestartNeighbor) GetFilter() yfilter.YFilter { return restartNeighbor.YFilter }

func (restartNeighbor *RestartNeighbor) SetFilter(yf yfilter.YFilter) { restartNeighbor.YFilter = yf }

func (restartNeighbor *RestartNeighbor) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (restartNeighbor *RestartNeighbor) GetSegmentPath() string {
    return "Cisco-IOS-XE-mpls-ldp:restart-neighbor"
}

func (restartNeighbor *RestartNeighbor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &restartNeighbor.Input
    }
    if childYangName == "output" {
        return &restartNeighbor.Output
    }
    return nil
}

func (restartNeighbor *RestartNeighbor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &restartNeighbor.Input
    children["output"] = &restartNeighbor.Output
    return children
}

func (restartNeighbor *RestartNeighbor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (restartNeighbor *RestartNeighbor) GetBundleName() string { return "cisco_ios_xe" }

func (restartNeighbor *RestartNeighbor) GetYangName() string { return "restart-neighbor" }

func (restartNeighbor *RestartNeighbor) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (restartNeighbor *RestartNeighbor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (restartNeighbor *RestartNeighbor) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (restartNeighbor *RestartNeighbor) SetParent(parent types.Entity) { restartNeighbor.parent = parent }

func (restartNeighbor *RestartNeighbor) GetParent() types.Entity { return restartNeighbor.parent }

func (restartNeighbor *RestartNeighbor) GetParentYangName() string { return "Cisco-IOS-XE-mpls-ldp" }

// RestartNeighbor_Input
type RestartNeighbor_Input struct {
    parent types.Entity
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

func (input *RestartNeighbor_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *RestartNeighbor_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *RestartNeighbor_Input) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "nbr-ip" { return "NbrIp" }
    if yname == "all" { return "All" }
    return ""
}

func (input *RestartNeighbor_Input) GetSegmentPath() string {
    return "input"
}

func (input *RestartNeighbor_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *RestartNeighbor_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *RestartNeighbor_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = input.VrfName
    leafs["nbr-ip"] = input.NbrIp
    leafs["all"] = input.All
    return leafs
}

func (input *RestartNeighbor_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *RestartNeighbor_Input) GetYangName() string { return "input" }

func (input *RestartNeighbor_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *RestartNeighbor_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *RestartNeighbor_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *RestartNeighbor_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *RestartNeighbor_Input) GetParent() types.Entity { return input.parent }

func (input *RestartNeighbor_Input) GetParentYangName() string { return "restart-neighbor" }

// RestartNeighbor_Output
type RestartNeighbor_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Return status will be 'OK' on success or an explanation string on failure.
    // The type is string.
    Status interface{}
}

func (output *RestartNeighbor_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *RestartNeighbor_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *RestartNeighbor_Output) GetGoName(yname string) string {
    if yname == "status" { return "Status" }
    return ""
}

func (output *RestartNeighbor_Output) GetSegmentPath() string {
    return "output"
}

func (output *RestartNeighbor_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *RestartNeighbor_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *RestartNeighbor_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["status"] = output.Status
    return leafs
}

func (output *RestartNeighbor_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *RestartNeighbor_Output) GetYangName() string { return "output" }

func (output *RestartNeighbor_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *RestartNeighbor_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *RestartNeighbor_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *RestartNeighbor_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *RestartNeighbor_Output) GetParent() types.Entity { return output.parent }

func (output *RestartNeighbor_Output) GetParentYangName() string { return "restart-neighbor" }

// ClearForwarding
// This command resets LDP installed forwarding state for all
// prefixes or a given prefix. It is useful when installed 
// LDP forwarding state needs to be reprogrammed in LSD and
// MPLS forwarding.
type ClearForwarding struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Input ClearForwarding_Input

    
    Output ClearForwarding_Output
}

func (clearForwarding *ClearForwarding) GetFilter() yfilter.YFilter { return clearForwarding.YFilter }

func (clearForwarding *ClearForwarding) SetFilter(yf yfilter.YFilter) { clearForwarding.YFilter = yf }

func (clearForwarding *ClearForwarding) GetGoName(yname string) string {
    if yname == "input" { return "Input" }
    if yname == "output" { return "Output" }
    return ""
}

func (clearForwarding *ClearForwarding) GetSegmentPath() string {
    return "Cisco-IOS-XE-mpls-ldp:clear-forwarding"
}

func (clearForwarding *ClearForwarding) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "input" {
        return &clearForwarding.Input
    }
    if childYangName == "output" {
        return &clearForwarding.Output
    }
    return nil
}

func (clearForwarding *ClearForwarding) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["input"] = &clearForwarding.Input
    children["output"] = &clearForwarding.Output
    return children
}

func (clearForwarding *ClearForwarding) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clearForwarding *ClearForwarding) GetBundleName() string { return "cisco_ios_xe" }

func (clearForwarding *ClearForwarding) GetYangName() string { return "clear-forwarding" }

func (clearForwarding *ClearForwarding) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (clearForwarding *ClearForwarding) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (clearForwarding *ClearForwarding) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (clearForwarding *ClearForwarding) SetParent(parent types.Entity) { clearForwarding.parent = parent }

func (clearForwarding *ClearForwarding) GetParent() types.Entity { return clearForwarding.parent }

func (clearForwarding *ClearForwarding) GetParentYangName() string { return "Cisco-IOS-XE-mpls-ldp" }

// ClearForwarding_Input
type ClearForwarding_Input struct {
    parent types.Entity
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

func (input *ClearForwarding_Input) GetFilter() yfilter.YFilter { return input.YFilter }

func (input *ClearForwarding_Input) SetFilter(yf yfilter.YFilter) { input.YFilter = yf }

func (input *ClearForwarding_Input) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "prefix-ip" { return "PrefixIp" }
    if yname == "all" { return "All" }
    return ""
}

func (input *ClearForwarding_Input) GetSegmentPath() string {
    return "input"
}

func (input *ClearForwarding_Input) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (input *ClearForwarding_Input) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (input *ClearForwarding_Input) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = input.VrfName
    leafs["prefix-ip"] = input.PrefixIp
    leafs["all"] = input.All
    return leafs
}

func (input *ClearForwarding_Input) GetBundleName() string { return "cisco_ios_xe" }

func (input *ClearForwarding_Input) GetYangName() string { return "input" }

func (input *ClearForwarding_Input) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (input *ClearForwarding_Input) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (input *ClearForwarding_Input) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (input *ClearForwarding_Input) SetParent(parent types.Entity) { input.parent = parent }

func (input *ClearForwarding_Input) GetParent() types.Entity { return input.parent }

func (input *ClearForwarding_Input) GetParentYangName() string { return "clear-forwarding" }

// ClearForwarding_Output
type ClearForwarding_Output struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Return status will be 'OK' on success or an explanatory string on failure.
    // The type is string.
    Status interface{}
}

func (output *ClearForwarding_Output) GetFilter() yfilter.YFilter { return output.YFilter }

func (output *ClearForwarding_Output) SetFilter(yf yfilter.YFilter) { output.YFilter = yf }

func (output *ClearForwarding_Output) GetGoName(yname string) string {
    if yname == "status" { return "Status" }
    return ""
}

func (output *ClearForwarding_Output) GetSegmentPath() string {
    return "output"
}

func (output *ClearForwarding_Output) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (output *ClearForwarding_Output) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (output *ClearForwarding_Output) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["status"] = output.Status
    return leafs
}

func (output *ClearForwarding_Output) GetBundleName() string { return "cisco_ios_xe" }

func (output *ClearForwarding_Output) GetYangName() string { return "output" }

func (output *ClearForwarding_Output) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (output *ClearForwarding_Output) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (output *ClearForwarding_Output) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (output *ClearForwarding_Output) SetParent(parent types.Entity) { output.parent = parent }

func (output *ClearForwarding_Output) GetParent() types.Entity { return output.parent }

func (output *ClearForwarding_Output) GetParentYangName() string { return "clear-forwarding" }

