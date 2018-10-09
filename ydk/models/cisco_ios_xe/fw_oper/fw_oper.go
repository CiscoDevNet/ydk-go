// This module contains a collection of YANG definitions
// for ZBFW operational data.
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package fw_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package fw_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-fw-oper fw-oper-data}", reflect.TypeOf(FwOperData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-fw-oper:fw-oper-data", reflect.TypeOf(FwOperData{}))
}

// FwOperData
// Operational state of ZBFW
type FwOperData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Firewall Drop Statistcis.
    FwDropStats FwOperData_FwDropStats

    // Firewall Zonepair list entries. The type is slice of
    // FwOperData_FwZonepairEntries.
    FwZonepairEntries []*FwOperData_FwZonepairEntries
}

func (fwOperData *FwOperData) GetEntityData() *types.CommonEntityData {
    fwOperData.EntityData.YFilter = fwOperData.YFilter
    fwOperData.EntityData.YangName = "fw-oper-data"
    fwOperData.EntityData.BundleName = "cisco_ios_xe"
    fwOperData.EntityData.ParentYangName = "Cisco-IOS-XE-fw-oper"
    fwOperData.EntityData.SegmentPath = "Cisco-IOS-XE-fw-oper:fw-oper-data"
    fwOperData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fwOperData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fwOperData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fwOperData.EntityData.Children = types.NewOrderedMap()
    fwOperData.EntityData.Children.Append("fw-drop-stats", types.YChild{"FwDropStats", &fwOperData.FwDropStats})
    fwOperData.EntityData.Children.Append("fw-zonepair-entries", types.YChild{"FwZonepairEntries", nil})
    for i := range fwOperData.FwZonepairEntries {
        fwOperData.EntityData.Children.Append(types.GetSegmentPath(fwOperData.FwZonepairEntries[i]), types.YChild{"FwZonepairEntries", fwOperData.FwZonepairEntries[i]})
    }
    fwOperData.EntityData.Leafs = types.NewOrderedMap()

    fwOperData.EntityData.YListKeys = []string {}

    return &(fwOperData.EntityData)
}

// FwOperData_FwDropStats
// Firewall Drop Statistcis
// This type is a presence type.
type FwOperData_FwDropStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Total packet drops seen since bringup. The type is interface{} with range:
    // 0..18446744073709551615.
    CatchAll interface{}

    // Packet drops due to maximum L4 half-open sessions reached. The type is
    // interface{} with range: 0..18446744073709551615.
    L4MaxHalfsession interface{}

    // Packet drops due to exceeding the maximum number of inspectable packets
    // allowed per flow. The type is interface{} with range:
    // 0..18446744073709551615.
    L4TooManyPkts interface{}

    // Packet drops for session initiators after exceeding maximum session limit.
    // The type is interface{} with range: 0..18446744073709551615.
    L4SessionLimit interface{}

    // Packet drops due to invalid header/packet length. The type is interface{}
    // with range: 0..18446744073709551615.
    L4InvalidHdr interface{}

    // Packet drops due to a failure in determining direction. The type is
    // interface{} with range: 0..18446744073709551615.
    L4InternalErrUndefinedDir interface{}

    // Packet drops due to session in internal close state. The type is
    // interface{} with range: 0..18446744073709551615.
    L4ScbClose interface{}

    // Packet drops due to invalid TCP ACK flags. The type is interface{} with
    // range: 0..18446744073709551615.
    L4TcpInvalidAckFlag interface{}

    // Packet drops due to invalid ACK number. The type is interface{} with range:
    // 0..18446744073709551615.
    L4TcpInvalidAckNum interface{}

    // Packet drops due to non-SYN packets received without valid session. The
    // type is interface{} with range: 0..18446744073709551615.
    L4TcpInvalidTcpInitiator interface{}

    // Packet drops due to SYN packets having data. The type is interface{} with
    // range: 0..18446744073709551615.
    L4TcpSynWithData interface{}

    // Packet drops due to invalid TCP window scale option. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpInvalidWinScaleOption interface{}

    // Packet drops due to invalid packets received in SYNSENT state. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpInvalidSegSynsentState interface{}

    // Packet drops due to invalid packets received in SYNRCVD state. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpInvalidSegSynrcvdState interface{}

    // Packet drops due to packets being too old/out of window. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpInvalidSegPktTooOld interface{}

    // Packet drops due to receiver window overflow (except when vTCP is enabled).
    // The type is interface{} with range: 0..18446744073709551615.
    L4TcpInvalidSegPktWinOverflow interface{}

    // Packet drops due to payload received after FIN is sent. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpInvalidSegPyldAfterFinSend interface{}

    // Packet drops due to invalid/unexpected TCP flags. The type is interface{}
    // with range: 0..18446744073709551615.
    L4TcpInvalidFlags interface{}

    // Packet drops due to invalid sequence number. The type is interface{} with
    // range: 0..18446744073709551615.
    L4TcpInvalidSeq interface{}

    // Packet drops due to invalid flags in TCP retransmitted packet. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpRetransInvalidFlags interface{}

    // Packet drops due to L7 not accepting out-of-order TCP segments. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpL7OooSeg interface{}

    // Packet drops during SYN flood attack. The type is interface{} with range:
    // 0..18446744073709551615.
    L4TcpSynFloodDrop interface{}

    // Packet drops due to failure of hostdb allocation during a SYN flood attack.
    // The type is interface{} with range: 0..18446744073709551615.
    L4TcpInternalErrSynfloodAllocHostdbFail interface{}

    // Packet drops due to blackout drop time when exceeding configured half-open
    // connections. The type is interface{} with range: 0..18446744073709551615.
    L4TcpSynfloodBlackoutDrop interface{}

    // Packet drops due to receiving TCP packet with payload when a response is
    // expected for SYN. The type is interface{} with range:
    // 0..18446744073709551615.
    L4TcpUnexpectTcpPayload interface{}

    // Packet drops due to receiving SYN in an established connection. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpSynInWin interface{}

    // Packet drops due to receiving RST in an established connection. The type is
    // interface{} with range: 0..18446744073709551615.
    L4TcpRstInWin interface{}

    // Packet drops due to unexpected/stray TCP segments. The type is interface{}
    // with range: 0..18446744073709551615.
    L4TcpStraySeg interface{}

    // RST sent to responder in SYNSENT state when ACK sequence is invalid. The
    // type is interface{} with range: 0..18446744073709551615.
    L4TcpRstToResp interface{}

    // Packet drops when policy exists in zone-pair but PAM lookup fails. The type
    // is interface{} with range: 0..18446744073709551615.
    InspPamLookupFail interface{}

    // Packet drops due to failure to get statistics block. The type is
    // interface{} with range: 0..18446744073709551615.
    InspInternalErrGetStatBlkFail interface{}

    // Packet drops due to destination address lookup failure. The type is
    // interface{} with range: 0..18446744073709551615.
    InspDstaddrLookupFail interface{}

    // Packet drops due to inspection policy not present in zone-pair. The type is
    // interface{} with range: 0..18446744073709551615.
    InspPolicyNotPresent interface{}

    // Packet drops due to session lookup failure and no matching policy present.
    // The type is interface{} with range: 0..18446744073709551615.
    InspSessMissPolicyNotPresent interface{}

    // Packet drops due to protocol classification failure. The type is
    // interface{} with range: 0..18446744073709551615.
    InspClassificationFail interface{}

    // Packet drops due to a drop classification action. The type is interface{}
    // with range: 0..18446744073709551615.
    InspClassActionDrop interface{}

    // Packet drops due to failed classification because of misconfigured security
    // policy. The type is interface{} with range: 0..18446744073709551615.
    InspPolicyMisconfigure interface{}

    // Packet drops after exceeding the maximum number of ICMP error packets
    // allowed per flow. The type is interface{} with range:
    // 0..18446744073709551615.
    L4IcmpTooManyErrPkts interface{}

    // Packet drops when ICMP is NATed without internal NAT info. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpInternalErrNoNat interface{}

    // Packet drops when ICMP failed to get error packets. The type is interface{}
    // with range: 0..18446744073709551615.
    L4IcmpInternalErrAllocFail interface{}

    // Packet drops due to a failure to get statistics block. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpInternalErrGetStatBlkFail interface{}

    // Packet drops due to unidentified ICMP packet direction. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpInternalErrDirNotIdentified interface{}

    // Packet drops due to receiving ICMP packets when session is in internal
    // close state. The type is interface{} with range: 0..18446744073709551615.
    L4IcmpScbClose interface{}

    // Packet drops due to missing IP header in ICMP packets. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpPktNoIpHdr interface{}

    // Packet drops due to ICMP error where packets are too short. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpPktTooShort interface{}

    // Packet drops due to packets not identified as IP or ICMP. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpErrNoIpNoIcmp interface{}

    // Packet drops due to ICMP error where packet bursts exceed a limit of 10.
    // The type is interface{} with range: 0..18446744073709551615.
    L4IcmpErrPktsBurst interface{}

    // Packet drops due to receiving multiple unreachable packets; only 1 is
    // allowed. The type is interface{} with range: 0..18446744073709551615.
    L4IcmpErrMultipleUnreach interface{}

    // Packet drops when inner TCP sequence number of packet doesn't match that of
    // packet originating the ICMP error. The type is interface{} with range:
    // 0..18446744073709551615.
    L4IcmpErrL4InvalidSeq interface{}

    // Packet drops due to inner TCP header invalid ACK. The type is interface{}
    // with range: 0..18446744073709551615.
    L4IcmpErrL4InvalidAck interface{}

    // Packet drops due to missing policy on zone-pair for ICMP. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpErrPolicyNotPresent interface{}

    // Packet drops due to a miss when doing reverse path flow check. The type is
    // interface{} with range: 0..18446744073709551615.
    L4IcmpErrClassificationFail interface{}

    // SYNcookie Packet drops when we've reached maximum number of SYN
    // destinations per zone. The type is interface{} with range:
    // 0..18446744073709551615.
    SyncookieMaxDst interface{}

    // SYNcookie Packet drops due to a failure in allocating memory in the
    // destination table. The type is interface{} with range:
    // 0..18446744073709551615.
    SyncookieInternalErrAllocFail interface{}

    // Packet drops due to a SYNcookie trigger. The type is interface{} with
    // range: 0..18446744073709551615.
    SyncookieTrigger interface{}

    // Packet drops due to dropping fragmented packet when first fragment drops.
    // The type is interface{} with range: 0..18446744073709551615.
    PolicyFragmentDrop interface{}

    // Packet drops when policy action is drop. The type is interface{} with
    // range: 0..18446744073709551615.
    PolicyActionDrop interface{}

    // Packet drops when policy action for the ICMP packet is to drop. The type is
    // interface{} with range: 0..18446744073709551615.
    PolicyIcmpActionDrop interface{}

    // Packet drops when L7 inspection returns drop as the action. The type is
    // interface{} with range: 0..18446744073709551615.
    L7TypeDrop interface{}

    // Packet drops due to receiving segmented packets when ALG doesn't honor
    // them. The type is interface{} with range: 0..18446744073709551615.
    L7NoSeg interface{}

    // Packet drops due to receiving fragmented packets when ALG doesn't honor
    // them. The type is interface{} with range: 0..18446744073709551615.
    L7NoFrag interface{}

    // Packet drops due to unrecognized L7 protocol type. The type is interface{}
    // with range: 0..18446744073709551615.
    L7UnknownProto interface{}

    // Packet drops due to L7 (ALG) deciding to drop the packet. The type is
    // interface{} with range: 0..18446744073709551615.
    L7AlgRetDrop interface{}

    // Packet drops due to L7 sub-channel promotion failure due to no zone pair
    // configured for the sub-channel. The type is interface{} with range:
    // 0..18446744073709551615.
    L7PromoteFailNoZonePair interface{}

    // Packet drops due to L7 sub-channel promotion failure due to no policy
    // configured for the sub-channel. The type is interface{} with range:
    // 0..18446744073709551615.
    L7PromoteFailNoPolicy interface{}

    // Packet drops due to session creation failure. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSession interface{}

    // Packet drops due to internal state not allowing new session creation. The
    // type is interface{} with range: 0..18446744073709551615.
    NoNewSession interface{}

    // Packet drops due to receiving a non-initiator packet for a session. The
    // type is interface{} with range: 0..18446744073709551615.
    NotInitiator interface{}

    // Packet drops due to a zone not configured for interface. The type is
    // interface{} with range: 0..18446744073709551615.
    InvalidZone interface{}

    // Packet drops due to asymmetric routing not configured and box not in active
    // state. The type is interface{} with range: 0..18446744073709551615.
    HaArStandby interface{}

    // Packet drops when Firewall is uninitialized. The type is interface{} with
    // range: 0..18446744073709551615.
    NoForwardingZone interface{}

    // Packet drops due to backpressure by log mechanism. The type is interface{}
    // with range: 0..18446744073709551615.
    Backpressure interface{}

    // Packet drops due to zone mismatch. The type is interface{} with range:
    // 0..18446744073709551615.
    ZoneMismatch interface{}

    // Packet drops due to a failure to register flow with flow database. The type
    // is interface{} with range: 0..18446744073709551615.
    FdbErr interface{}

    // Packet drops due to LISP header restoration failure. The type is
    // interface{} with range: 0..18446744073709551615.
    LispHeaderRestoreFail interface{}

    // Packet drops due to LISP inner packet sanity check failure. The type is
    // interface{} with range: 0..18446744073709551615.
    LispInnerPktInsane interface{}

    // Packet drops due to LISP inner packet IPV4 sanity check failure. The type
    // is interface{} with range: 0..18446744073709551615.
    LispInnerIpv4Insane interface{}

    // Packet drops due to LISP inner packet IPV6 sanity check failure. The type
    // is interface{} with range: 0..18446744073709551615.
    LispInnerIpv6Insane interface{}
}

func (fwDropStats *FwOperData_FwDropStats) GetEntityData() *types.CommonEntityData {
    fwDropStats.EntityData.YFilter = fwDropStats.YFilter
    fwDropStats.EntityData.YangName = "fw-drop-stats"
    fwDropStats.EntityData.BundleName = "cisco_ios_xe"
    fwDropStats.EntityData.ParentYangName = "fw-oper-data"
    fwDropStats.EntityData.SegmentPath = "fw-drop-stats"
    fwDropStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fwDropStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fwDropStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fwDropStats.EntityData.Children = types.NewOrderedMap()
    fwDropStats.EntityData.Leafs = types.NewOrderedMap()
    fwDropStats.EntityData.Leafs.Append("catch-all", types.YLeaf{"CatchAll", fwDropStats.CatchAll})
    fwDropStats.EntityData.Leafs.Append("l4-max-halfsession", types.YLeaf{"L4MaxHalfsession", fwDropStats.L4MaxHalfsession})
    fwDropStats.EntityData.Leafs.Append("l4-too-many-pkts", types.YLeaf{"L4TooManyPkts", fwDropStats.L4TooManyPkts})
    fwDropStats.EntityData.Leafs.Append("l4-session-limit", types.YLeaf{"L4SessionLimit", fwDropStats.L4SessionLimit})
    fwDropStats.EntityData.Leafs.Append("l4-invalid-hdr", types.YLeaf{"L4InvalidHdr", fwDropStats.L4InvalidHdr})
    fwDropStats.EntityData.Leafs.Append("l4-internal-err-undefined-dir", types.YLeaf{"L4InternalErrUndefinedDir", fwDropStats.L4InternalErrUndefinedDir})
    fwDropStats.EntityData.Leafs.Append("l4-scb-close", types.YLeaf{"L4ScbClose", fwDropStats.L4ScbClose})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-ack-flag", types.YLeaf{"L4TcpInvalidAckFlag", fwDropStats.L4TcpInvalidAckFlag})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-ack-num", types.YLeaf{"L4TcpInvalidAckNum", fwDropStats.L4TcpInvalidAckNum})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-tcp-initiator", types.YLeaf{"L4TcpInvalidTcpInitiator", fwDropStats.L4TcpInvalidTcpInitiator})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-syn-with-data", types.YLeaf{"L4TcpSynWithData", fwDropStats.L4TcpSynWithData})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-win-scale-option", types.YLeaf{"L4TcpInvalidWinScaleOption", fwDropStats.L4TcpInvalidWinScaleOption})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-seg-synsent-state", types.YLeaf{"L4TcpInvalidSegSynsentState", fwDropStats.L4TcpInvalidSegSynsentState})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-seg-synrcvd-state", types.YLeaf{"L4TcpInvalidSegSynrcvdState", fwDropStats.L4TcpInvalidSegSynrcvdState})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-seg-pkt-too-old", types.YLeaf{"L4TcpInvalidSegPktTooOld", fwDropStats.L4TcpInvalidSegPktTooOld})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-seg-pkt-win-overflow", types.YLeaf{"L4TcpInvalidSegPktWinOverflow", fwDropStats.L4TcpInvalidSegPktWinOverflow})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-seg-pyld-after-fin-send", types.YLeaf{"L4TcpInvalidSegPyldAfterFinSend", fwDropStats.L4TcpInvalidSegPyldAfterFinSend})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-flags", types.YLeaf{"L4TcpInvalidFlags", fwDropStats.L4TcpInvalidFlags})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-invalid-seq", types.YLeaf{"L4TcpInvalidSeq", fwDropStats.L4TcpInvalidSeq})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-retrans-invalid-flags", types.YLeaf{"L4TcpRetransInvalidFlags", fwDropStats.L4TcpRetransInvalidFlags})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-l7-ooo-seg", types.YLeaf{"L4TcpL7OooSeg", fwDropStats.L4TcpL7OooSeg})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-syn-flood-drop", types.YLeaf{"L4TcpSynFloodDrop", fwDropStats.L4TcpSynFloodDrop})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-internal-err-synflood-alloc-hostdb-fail", types.YLeaf{"L4TcpInternalErrSynfloodAllocHostdbFail", fwDropStats.L4TcpInternalErrSynfloodAllocHostdbFail})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-synflood-blackout-drop", types.YLeaf{"L4TcpSynfloodBlackoutDrop", fwDropStats.L4TcpSynfloodBlackoutDrop})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-unexpect-tcp-payload", types.YLeaf{"L4TcpUnexpectTcpPayload", fwDropStats.L4TcpUnexpectTcpPayload})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-syn-in-win", types.YLeaf{"L4TcpSynInWin", fwDropStats.L4TcpSynInWin})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-rst-in-win", types.YLeaf{"L4TcpRstInWin", fwDropStats.L4TcpRstInWin})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-stray-seg", types.YLeaf{"L4TcpStraySeg", fwDropStats.L4TcpStraySeg})
    fwDropStats.EntityData.Leafs.Append("l4-tcp-rst-to-resp", types.YLeaf{"L4TcpRstToResp", fwDropStats.L4TcpRstToResp})
    fwDropStats.EntityData.Leafs.Append("insp-pam-lookup-fail", types.YLeaf{"InspPamLookupFail", fwDropStats.InspPamLookupFail})
    fwDropStats.EntityData.Leafs.Append("insp-internal-err-get-stat-blk-fail", types.YLeaf{"InspInternalErrGetStatBlkFail", fwDropStats.InspInternalErrGetStatBlkFail})
    fwDropStats.EntityData.Leafs.Append("insp-dstaddr-lookup-fail", types.YLeaf{"InspDstaddrLookupFail", fwDropStats.InspDstaddrLookupFail})
    fwDropStats.EntityData.Leafs.Append("insp-policy-not-present", types.YLeaf{"InspPolicyNotPresent", fwDropStats.InspPolicyNotPresent})
    fwDropStats.EntityData.Leafs.Append("insp-sess-miss-policy-not-present", types.YLeaf{"InspSessMissPolicyNotPresent", fwDropStats.InspSessMissPolicyNotPresent})
    fwDropStats.EntityData.Leafs.Append("insp-classification-fail", types.YLeaf{"InspClassificationFail", fwDropStats.InspClassificationFail})
    fwDropStats.EntityData.Leafs.Append("insp-class-action-drop", types.YLeaf{"InspClassActionDrop", fwDropStats.InspClassActionDrop})
    fwDropStats.EntityData.Leafs.Append("insp-policy-misconfigure", types.YLeaf{"InspPolicyMisconfigure", fwDropStats.InspPolicyMisconfigure})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-too-many-err-pkts", types.YLeaf{"L4IcmpTooManyErrPkts", fwDropStats.L4IcmpTooManyErrPkts})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-internal-err-no-nat", types.YLeaf{"L4IcmpInternalErrNoNat", fwDropStats.L4IcmpInternalErrNoNat})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-internal-err-alloc-fail", types.YLeaf{"L4IcmpInternalErrAllocFail", fwDropStats.L4IcmpInternalErrAllocFail})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-internal-err-get-stat-blk-fail", types.YLeaf{"L4IcmpInternalErrGetStatBlkFail", fwDropStats.L4IcmpInternalErrGetStatBlkFail})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-internal-err-dir-not-identified", types.YLeaf{"L4IcmpInternalErrDirNotIdentified", fwDropStats.L4IcmpInternalErrDirNotIdentified})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-scb-close", types.YLeaf{"L4IcmpScbClose", fwDropStats.L4IcmpScbClose})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-pkt-no-ip-hdr", types.YLeaf{"L4IcmpPktNoIpHdr", fwDropStats.L4IcmpPktNoIpHdr})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-pkt-too-short", types.YLeaf{"L4IcmpPktTooShort", fwDropStats.L4IcmpPktTooShort})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-no-ip-no-icmp", types.YLeaf{"L4IcmpErrNoIpNoIcmp", fwDropStats.L4IcmpErrNoIpNoIcmp})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-pkts-burst", types.YLeaf{"L4IcmpErrPktsBurst", fwDropStats.L4IcmpErrPktsBurst})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-multiple-unreach", types.YLeaf{"L4IcmpErrMultipleUnreach", fwDropStats.L4IcmpErrMultipleUnreach})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-l4-invalid-seq", types.YLeaf{"L4IcmpErrL4InvalidSeq", fwDropStats.L4IcmpErrL4InvalidSeq})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-l4-invalid-ack", types.YLeaf{"L4IcmpErrL4InvalidAck", fwDropStats.L4IcmpErrL4InvalidAck})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-policy-not-present", types.YLeaf{"L4IcmpErrPolicyNotPresent", fwDropStats.L4IcmpErrPolicyNotPresent})
    fwDropStats.EntityData.Leafs.Append("l4-icmp-err-classification-fail", types.YLeaf{"L4IcmpErrClassificationFail", fwDropStats.L4IcmpErrClassificationFail})
    fwDropStats.EntityData.Leafs.Append("syncookie-max-dst", types.YLeaf{"SyncookieMaxDst", fwDropStats.SyncookieMaxDst})
    fwDropStats.EntityData.Leafs.Append("syncookie-internal-err-alloc-fail", types.YLeaf{"SyncookieInternalErrAllocFail", fwDropStats.SyncookieInternalErrAllocFail})
    fwDropStats.EntityData.Leafs.Append("syncookie-trigger", types.YLeaf{"SyncookieTrigger", fwDropStats.SyncookieTrigger})
    fwDropStats.EntityData.Leafs.Append("policy-fragment-drop", types.YLeaf{"PolicyFragmentDrop", fwDropStats.PolicyFragmentDrop})
    fwDropStats.EntityData.Leafs.Append("policy-action-drop", types.YLeaf{"PolicyActionDrop", fwDropStats.PolicyActionDrop})
    fwDropStats.EntityData.Leafs.Append("policy-icmp-action-drop", types.YLeaf{"PolicyIcmpActionDrop", fwDropStats.PolicyIcmpActionDrop})
    fwDropStats.EntityData.Leafs.Append("l7-type-drop", types.YLeaf{"L7TypeDrop", fwDropStats.L7TypeDrop})
    fwDropStats.EntityData.Leafs.Append("l7-no-seg", types.YLeaf{"L7NoSeg", fwDropStats.L7NoSeg})
    fwDropStats.EntityData.Leafs.Append("l7-no-frag", types.YLeaf{"L7NoFrag", fwDropStats.L7NoFrag})
    fwDropStats.EntityData.Leafs.Append("l7-unknown-proto", types.YLeaf{"L7UnknownProto", fwDropStats.L7UnknownProto})
    fwDropStats.EntityData.Leafs.Append("l7-alg-ret-drop", types.YLeaf{"L7AlgRetDrop", fwDropStats.L7AlgRetDrop})
    fwDropStats.EntityData.Leafs.Append("l7-promote-fail-no-zone-pair", types.YLeaf{"L7PromoteFailNoZonePair", fwDropStats.L7PromoteFailNoZonePair})
    fwDropStats.EntityData.Leafs.Append("l7-promote-fail-no-policy", types.YLeaf{"L7PromoteFailNoPolicy", fwDropStats.L7PromoteFailNoPolicy})
    fwDropStats.EntityData.Leafs.Append("no-session", types.YLeaf{"NoSession", fwDropStats.NoSession})
    fwDropStats.EntityData.Leafs.Append("no-new-session", types.YLeaf{"NoNewSession", fwDropStats.NoNewSession})
    fwDropStats.EntityData.Leafs.Append("not-initiator", types.YLeaf{"NotInitiator", fwDropStats.NotInitiator})
    fwDropStats.EntityData.Leafs.Append("invalid-zone", types.YLeaf{"InvalidZone", fwDropStats.InvalidZone})
    fwDropStats.EntityData.Leafs.Append("ha-ar-standby", types.YLeaf{"HaArStandby", fwDropStats.HaArStandby})
    fwDropStats.EntityData.Leafs.Append("no-forwarding-zone", types.YLeaf{"NoForwardingZone", fwDropStats.NoForwardingZone})
    fwDropStats.EntityData.Leafs.Append("backpressure", types.YLeaf{"Backpressure", fwDropStats.Backpressure})
    fwDropStats.EntityData.Leafs.Append("zone-mismatch", types.YLeaf{"ZoneMismatch", fwDropStats.ZoneMismatch})
    fwDropStats.EntityData.Leafs.Append("fdb-err", types.YLeaf{"FdbErr", fwDropStats.FdbErr})
    fwDropStats.EntityData.Leafs.Append("lisp-header-restore-fail", types.YLeaf{"LispHeaderRestoreFail", fwDropStats.LispHeaderRestoreFail})
    fwDropStats.EntityData.Leafs.Append("lisp-inner-pkt-insane", types.YLeaf{"LispInnerPktInsane", fwDropStats.LispInnerPktInsane})
    fwDropStats.EntityData.Leafs.Append("lisp-inner-ipv4-insane", types.YLeaf{"LispInnerIpv4Insane", fwDropStats.LispInnerIpv4Insane})
    fwDropStats.EntityData.Leafs.Append("lisp-inner-ipv6-insane", types.YLeaf{"LispInnerIpv6Insane", fwDropStats.LispInnerIpv6Insane})

    fwDropStats.EntityData.YListKeys = []string {}

    return &(fwDropStats.EntityData)
}

// FwOperData_FwZonepairEntries
// Firewall Zonepair list entries
type FwOperData_FwZonepairEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the zone pair. The type is string.
    ZonepairName interface{}

    // Name of the source zone. The type is string.
    SrcZoneName interface{}

    // Name of the destination zone. The type is string.
    DstZoneName interface{}

    // Name of the policy applied for this zone pair. The type is string.
    PolicyName interface{}

    // Firewall Traffic class list entries. The type is slice of
    // FwOperData_FwZonepairEntries_FwTrafficClassEntry.
    FwTrafficClassEntry []*FwOperData_FwZonepairEntries_FwTrafficClassEntry
}

func (fwZonepairEntries *FwOperData_FwZonepairEntries) GetEntityData() *types.CommonEntityData {
    fwZonepairEntries.EntityData.YFilter = fwZonepairEntries.YFilter
    fwZonepairEntries.EntityData.YangName = "fw-zonepair-entries"
    fwZonepairEntries.EntityData.BundleName = "cisco_ios_xe"
    fwZonepairEntries.EntityData.ParentYangName = "fw-oper-data"
    fwZonepairEntries.EntityData.SegmentPath = "fw-zonepair-entries" + types.AddKeyToken(fwZonepairEntries.ZonepairName, "zonepair-name")
    fwZonepairEntries.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fwZonepairEntries.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fwZonepairEntries.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fwZonepairEntries.EntityData.Children = types.NewOrderedMap()
    fwZonepairEntries.EntityData.Children.Append("fw-traffic-class-entry", types.YChild{"FwTrafficClassEntry", nil})
    for i := range fwZonepairEntries.FwTrafficClassEntry {
        fwZonepairEntries.EntityData.Children.Append(types.GetSegmentPath(fwZonepairEntries.FwTrafficClassEntry[i]), types.YChild{"FwTrafficClassEntry", fwZonepairEntries.FwTrafficClassEntry[i]})
    }
    fwZonepairEntries.EntityData.Leafs = types.NewOrderedMap()
    fwZonepairEntries.EntityData.Leafs.Append("zonepair-name", types.YLeaf{"ZonepairName", fwZonepairEntries.ZonepairName})
    fwZonepairEntries.EntityData.Leafs.Append("src-zone-name", types.YLeaf{"SrcZoneName", fwZonepairEntries.SrcZoneName})
    fwZonepairEntries.EntityData.Leafs.Append("dst-zone-name", types.YLeaf{"DstZoneName", fwZonepairEntries.DstZoneName})
    fwZonepairEntries.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", fwZonepairEntries.PolicyName})

    fwZonepairEntries.EntityData.YListKeys = []string {"ZonepairName"}

    return &(fwZonepairEntries.EntityData)
}

// FwOperData_FwZonepairEntries_FwTrafficClassEntry
// Firewall Traffic class list entries
type FwOperData_FwZonepairEntries_FwTrafficClassEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the traffic class. The type is string.
    ClassName interface{}

    // Zonepair name. The type is string.
    ZonepairName interface{}

    // Action for the traffic class. The type is string.
    ClassAction interface{}

    // Total Packets. The type is interface{} with range: 0..18446744073709551615.
    PktsCounter interface{}

    // Total bytes. The type is interface{} with range: 0..18446744073709551615.
    BytesCounter interface{}

    // Total number for the attempted connections matching this traffic class. The
    // type is interface{} with range: 0..18446744073709551615.
    AttemptedConn interface{}

    // Current number of active connections matching this traffic class. The type
    // is interface{} with range: 0..18446744073709551615.
    CurrentActiveConn interface{}

    // Maximum number of active connections seen for this traffic class. The type
    // is interface{} with range: 0..18446744073709551615.
    MaxActiveConn interface{}

    // Current number of half-open connections seen for this traffic class. The
    // type is interface{} with range: 0..18446744073709551615.
    CurrentHalfopenConn interface{}

    // Maximum number of half-open connections seen for this traffic class. The
    // type is interface{} with range: 0..18446744073709551615.
    MaxHalfopenConn interface{}

    // Current number of terminating connections seen for this traffic class. The
    // type is interface{} with range: 0..18446744073709551615.
    CurrentTerminatingConn interface{}

    // Maximum number of terminating connections seen for this traffic class. The
    // type is interface{} with range: 0..18446744073709551615.
    MaxTerminatingConn interface{}

    // Seconds since last session creation. The type is interface{} with range:
    // 0..18446744073709551615.
    TimeSinceLastSessionCreate interface{}

    // List of match conditions. The type is slice of
    // FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcMatchEntry.
    FwTcMatchEntry []*FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcMatchEntry

    // Firewall Traffic class protocol list entries. The type is slice of
    // FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcProtoEntry.
    FwTcProtoEntry []*FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcProtoEntry
}

func (fwTrafficClassEntry *FwOperData_FwZonepairEntries_FwTrafficClassEntry) GetEntityData() *types.CommonEntityData {
    fwTrafficClassEntry.EntityData.YFilter = fwTrafficClassEntry.YFilter
    fwTrafficClassEntry.EntityData.YangName = "fw-traffic-class-entry"
    fwTrafficClassEntry.EntityData.BundleName = "cisco_ios_xe"
    fwTrafficClassEntry.EntityData.ParentYangName = "fw-zonepair-entries"
    fwTrafficClassEntry.EntityData.SegmentPath = "fw-traffic-class-entry" + types.AddKeyToken(fwTrafficClassEntry.ClassName, "class-name")
    fwTrafficClassEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fwTrafficClassEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fwTrafficClassEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fwTrafficClassEntry.EntityData.Children = types.NewOrderedMap()
    fwTrafficClassEntry.EntityData.Children.Append("fw-tc-match-entry", types.YChild{"FwTcMatchEntry", nil})
    for i := range fwTrafficClassEntry.FwTcMatchEntry {
        fwTrafficClassEntry.EntityData.Children.Append(types.GetSegmentPath(fwTrafficClassEntry.FwTcMatchEntry[i]), types.YChild{"FwTcMatchEntry", fwTrafficClassEntry.FwTcMatchEntry[i]})
    }
    fwTrafficClassEntry.EntityData.Children.Append("fw-tc-proto-entry", types.YChild{"FwTcProtoEntry", nil})
    for i := range fwTrafficClassEntry.FwTcProtoEntry {
        fwTrafficClassEntry.EntityData.Children.Append(types.GetSegmentPath(fwTrafficClassEntry.FwTcProtoEntry[i]), types.YChild{"FwTcProtoEntry", fwTrafficClassEntry.FwTcProtoEntry[i]})
    }
    fwTrafficClassEntry.EntityData.Leafs = types.NewOrderedMap()
    fwTrafficClassEntry.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", fwTrafficClassEntry.ClassName})
    fwTrafficClassEntry.EntityData.Leafs.Append("zonepair-name", types.YLeaf{"ZonepairName", fwTrafficClassEntry.ZonepairName})
    fwTrafficClassEntry.EntityData.Leafs.Append("class-action", types.YLeaf{"ClassAction", fwTrafficClassEntry.ClassAction})
    fwTrafficClassEntry.EntityData.Leafs.Append("pkts-counter", types.YLeaf{"PktsCounter", fwTrafficClassEntry.PktsCounter})
    fwTrafficClassEntry.EntityData.Leafs.Append("bytes-counter", types.YLeaf{"BytesCounter", fwTrafficClassEntry.BytesCounter})
    fwTrafficClassEntry.EntityData.Leafs.Append("attempted-conn", types.YLeaf{"AttemptedConn", fwTrafficClassEntry.AttemptedConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("current-active-conn", types.YLeaf{"CurrentActiveConn", fwTrafficClassEntry.CurrentActiveConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("max-active-conn", types.YLeaf{"MaxActiveConn", fwTrafficClassEntry.MaxActiveConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("current-halfopen-conn", types.YLeaf{"CurrentHalfopenConn", fwTrafficClassEntry.CurrentHalfopenConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("max-halfopen-conn", types.YLeaf{"MaxHalfopenConn", fwTrafficClassEntry.MaxHalfopenConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("current-terminating-conn", types.YLeaf{"CurrentTerminatingConn", fwTrafficClassEntry.CurrentTerminatingConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("max-terminating-conn", types.YLeaf{"MaxTerminatingConn", fwTrafficClassEntry.MaxTerminatingConn})
    fwTrafficClassEntry.EntityData.Leafs.Append("time-since-last-session-create", types.YLeaf{"TimeSinceLastSessionCreate", fwTrafficClassEntry.TimeSinceLastSessionCreate})

    fwTrafficClassEntry.EntityData.YListKeys = []string {"ClassName"}

    return &(fwTrafficClassEntry.EntityData)
}

// FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcMatchEntry
// List of match conditions
type FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcMatchEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Match Name. The type is string.
    MatchName interface{}

    // This attribute is a key. Match Type Identifier. The type is interface{}
    // with range: 0..255.
    MatchTypeId interface{}

    // Match Type. The type is string.
    MatchType interface{}
}

func (fwTcMatchEntry *FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcMatchEntry) GetEntityData() *types.CommonEntityData {
    fwTcMatchEntry.EntityData.YFilter = fwTcMatchEntry.YFilter
    fwTcMatchEntry.EntityData.YangName = "fw-tc-match-entry"
    fwTcMatchEntry.EntityData.BundleName = "cisco_ios_xe"
    fwTcMatchEntry.EntityData.ParentYangName = "fw-traffic-class-entry"
    fwTcMatchEntry.EntityData.SegmentPath = "fw-tc-match-entry" + types.AddKeyToken(fwTcMatchEntry.MatchName, "match-name") + types.AddKeyToken(fwTcMatchEntry.MatchTypeId, "match-type-id")
    fwTcMatchEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fwTcMatchEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fwTcMatchEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fwTcMatchEntry.EntityData.Children = types.NewOrderedMap()
    fwTcMatchEntry.EntityData.Leafs = types.NewOrderedMap()
    fwTcMatchEntry.EntityData.Leafs.Append("match-name", types.YLeaf{"MatchName", fwTcMatchEntry.MatchName})
    fwTcMatchEntry.EntityData.Leafs.Append("match-type-id", types.YLeaf{"MatchTypeId", fwTcMatchEntry.MatchTypeId})
    fwTcMatchEntry.EntityData.Leafs.Append("match-type", types.YLeaf{"MatchType", fwTcMatchEntry.MatchType})

    fwTcMatchEntry.EntityData.YListKeys = []string {"MatchName", "MatchTypeId"}

    return &(fwTcMatchEntry.EntityData)
}

// FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcProtoEntry
// Firewall Traffic class protocol list entries
type FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcProtoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol ID. The type is interface{} with range:
    // 0..65535.
    ProtoId interface{}

    // Protocol Name. The type is string.
    ProtocolName interface{}

    // Number of bytes matching this protocol. The type is interface{} with range:
    // 0..18446744073709551615.
    ByteCounters interface{}

    // Number of packets matching this protocol. The type is interface{} with
    // range: 0..18446744073709551615.
    PktCounters interface{}
}

func (fwTcProtoEntry *FwOperData_FwZonepairEntries_FwTrafficClassEntry_FwTcProtoEntry) GetEntityData() *types.CommonEntityData {
    fwTcProtoEntry.EntityData.YFilter = fwTcProtoEntry.YFilter
    fwTcProtoEntry.EntityData.YangName = "fw-tc-proto-entry"
    fwTcProtoEntry.EntityData.BundleName = "cisco_ios_xe"
    fwTcProtoEntry.EntityData.ParentYangName = "fw-traffic-class-entry"
    fwTcProtoEntry.EntityData.SegmentPath = "fw-tc-proto-entry" + types.AddKeyToken(fwTcProtoEntry.ProtoId, "proto-id")
    fwTcProtoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    fwTcProtoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    fwTcProtoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    fwTcProtoEntry.EntityData.Children = types.NewOrderedMap()
    fwTcProtoEntry.EntityData.Leafs = types.NewOrderedMap()
    fwTcProtoEntry.EntityData.Leafs.Append("proto-id", types.YLeaf{"ProtoId", fwTcProtoEntry.ProtoId})
    fwTcProtoEntry.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", fwTcProtoEntry.ProtocolName})
    fwTcProtoEntry.EntityData.Leafs.Append("byte-counters", types.YLeaf{"ByteCounters", fwTcProtoEntry.ByteCounters})
    fwTcProtoEntry.EntityData.Leafs.Append("pkt-counters", types.YLeaf{"PktCounters", fwTcProtoEntry.PktCounters})

    fwTcProtoEntry.EntityData.YListKeys = []string {"ProtoId"}

    return &(fwTcProtoEntry.EntityData)
}

