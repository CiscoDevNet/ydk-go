// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-ntp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ntp: NTP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_ntp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_ntp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-ntp-oper ntp}", reflect.TypeOf(Ntp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-ntp-oper:ntp", reflect.TypeOf(Ntp{}))
}

// NtpPeerStatus represents Type of peer status
type NtpPeerStatus string

const (
    //    reject
    NtpPeerStatus_ntp_ctl_pst_sel_reject NtpPeerStatus = "ntp-ctl-pst-sel-reject"

    //  x falsetick
    NtpPeerStatus_ntp_ctl_pst_sel_sane NtpPeerStatus = "ntp-ctl-pst-sel-sane"

    //  . excess 
    NtpPeerStatus_ntp_ctl_pst_sel_correct NtpPeerStatus = "ntp-ctl-pst-sel-correct"

    //  - outlyer
    NtpPeerStatus_ntp_ctl_pst_sel_selcand NtpPeerStatus = "ntp-ctl-pst-sel-selcand"

    //  + candidate
    NtpPeerStatus_ntp_ctl_pst_sel_sync_cand NtpPeerStatus = "ntp-ctl-pst-sel-sync-cand"

    //  # selected
    NtpPeerStatus_ntp_ctl_pst_sel_distsys_peer NtpPeerStatus = "ntp-ctl-pst-sel-distsys-peer"

    //  * sys peer
    NtpPeerStatus_ntp_ctl_pst_sel_sys_peer NtpPeerStatus = "ntp-ctl-pst-sel-sys-peer"

    //  o pps peer
    NtpPeerStatus_ntp_ctl_pst_sel_pps NtpPeerStatus = "ntp-ctl-pst-sel-pps"
)

// NtpMode represents Type of mode
type NtpMode string

const (
    // Unspecified probably old NTP version
    NtpMode_ntp_mode_unspec NtpMode = "ntp-mode-unspec"

    // Symmetric active
    NtpMode_ntp_mode_symetric_active NtpMode = "ntp-mode-symetric-active"

    // Symmetric passive
    NtpMode_ntp_mode_symetric_passive NtpMode = "ntp-mode-symetric-passive"

    // Client mode
    NtpMode_ntp_mode_client NtpMode = "ntp-mode-client"

    // Server mode
    NtpMode_ntp_mode_server NtpMode = "ntp-mode-server"

    // Broadcast mode
    NtpMode_ntp_mode_xcast_server NtpMode = "ntp-mode-xcast-server"

    // Control mode packet
    NtpMode_ntp_mode_control NtpMode = "ntp-mode-control"

    // Implementation defined function
    NtpMode_ntp_mode_private NtpMode = "ntp-mode-private"

    // A broadcast client mode
    NtpMode_ntp_mode_xcast_client NtpMode = "ntp-mode-xcast-client"
)

// ClockUpdateNode represents Mode of Clock Update
type ClockUpdateNode string

const (
    //  clock is never updated
    ClockUpdateNode_clk_never_updated ClockUpdateNode = "clk-never-updated"

    //  clock is updated
    ClockUpdateNode_clk_updated ClockUpdateNode = "clk-updated"

    //  clock has no update info
    ClockUpdateNode_clk_no_update_info ClockUpdateNode = "clk-no-update-info"
)

// NtpLoopFilterState represents Loop filter state
type NtpLoopFilterState string

const (
    //  never set
    NtpLoopFilterState_ntp_loop_flt_n_set NtpLoopFilterState = "ntp-loop-flt-n-set"

    //  drift set from file
    NtpLoopFilterState_ntp_loop_flt_f_set NtpLoopFilterState = "ntp-loop-flt-f-set"

    //  spike
    NtpLoopFilterState_ntp_loop_flt_spik NtpLoopFilterState = "ntp-loop-flt-spik"

    //  drift being measured
    NtpLoopFilterState_ntp_loop_flt_freq NtpLoopFilterState = "ntp-loop-flt-freq"

    //  normal controlled loop
    NtpLoopFilterState_ntp_loop_flt_sync NtpLoopFilterState = "ntp-loop-flt-sync"

    //  unknown
    NtpLoopFilterState_ntp_loop_flt_unkn NtpLoopFilterState = "ntp-loop-flt-unkn"
)

// NtpLeap represents Type of leap
type NtpLeap string

const (
    // Normal, no leap second warning
    NtpLeap_ntp_leap_no_warning NtpLeap = "ntp-leap-no-warning"

    // Last minute of day has 61 seconds
    NtpLeap_ntp_leap_addse_cond NtpLeap = "ntp-leap-addse-cond"

    // Last minute of day has 59 seconds
    NtpLeap_ntp_leap_delse_cond NtpLeap = "ntp-leap-delse-cond"

    // Overload, clock is free running
    NtpLeap_ntp_leap_not_in_sync NtpLeap = "ntp-leap-not-in-sync"
)

// Ntp
// NTP operational data
type Ntp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific NTP operational data.
    Nodes Ntp_Nodes
}

func (ntp *Ntp) GetFilter() yfilter.YFilter { return ntp.YFilter }

func (ntp *Ntp) SetFilter(yf yfilter.YFilter) { ntp.YFilter = yf }

func (ntp *Ntp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ntp *Ntp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-ntp-oper:ntp"
}

func (ntp *Ntp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ntp.Nodes
    }
    return nil
}

func (ntp *Ntp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ntp.Nodes
    return children
}

func (ntp *Ntp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ntp *Ntp) GetBundleName() string { return "cisco_ios_xr" }

func (ntp *Ntp) GetYangName() string { return "ntp" }

func (ntp *Ntp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ntp *Ntp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ntp *Ntp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ntp *Ntp) SetParent(parent types.Entity) { ntp.parent = parent }

func (ntp *Ntp) GetParent() types.Entity { return ntp.parent }

func (ntp *Ntp) GetParentYangName() string { return "Cisco-IOS-XR-ip-ntp-oper" }

// Ntp_Nodes
// Node-specific NTP operational data
type Ntp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NTP operational data for a particular node. The type is slice of
    // Ntp_Nodes_Node.
    Node []Ntp_Nodes_Node
}

func (nodes *Ntp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ntp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ntp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ntp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ntp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ntp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ntp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ntp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ntp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ntp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ntp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ntp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ntp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ntp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ntp_Nodes) GetParentYangName() string { return "ntp" }

// Ntp_Nodes_Node
// NTP operational data for a particular node
type Ntp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The node identifier. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Node interface{}

    // NTP Associations Detail information.
    AssociationsDetail Ntp_Nodes_Node_AssociationsDetail

    // Status of NTP peer(s).
    Status Ntp_Nodes_Node_Status

    // NTP Associations information.
    Associations Ntp_Nodes_Node_Associations
}

func (node *Ntp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ntp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ntp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "associations-detail" { return "AssociationsDetail" }
    if yname == "status" { return "Status" }
    if yname == "associations" { return "Associations" }
    return ""
}

func (node *Ntp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
}

func (node *Ntp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "associations-detail" {
        return &node.AssociationsDetail
    }
    if childYangName == "status" {
        return &node.Status
    }
    if childYangName == "associations" {
        return &node.Associations
    }
    return nil
}

func (node *Ntp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["associations-detail"] = &node.AssociationsDetail
    children["status"] = &node.Status
    children["associations"] = &node.Associations
    return children
}

func (node *Ntp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = node.Node
    return leafs
}

func (node *Ntp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ntp_Nodes_Node) GetYangName() string { return "node" }

func (node *Ntp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ntp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ntp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ntp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ntp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ntp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ntp_Nodes_Node_AssociationsDetail
// NTP Associations Detail information
type Ntp_Nodes_Node_AssociationsDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is NTP enabled. The type is bool.
    IsNtpEnabled interface{}

    // Leap. The type is NtpLeap.
    SysLeap interface{}

    // Peer info. The type is slice of
    // Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo.
    PeerDetailInfo []Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo
}

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetFilter() yfilter.YFilter { return associationsDetail.YFilter }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) SetFilter(yf yfilter.YFilter) { associationsDetail.YFilter = yf }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetGoName(yname string) string {
    if yname == "is-ntp-enabled" { return "IsNtpEnabled" }
    if yname == "sys-leap" { return "SysLeap" }
    if yname == "peer-detail-info" { return "PeerDetailInfo" }
    return ""
}

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetSegmentPath() string {
    return "associations-detail"
}

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-detail-info" {
        for _, c := range associationsDetail.PeerDetailInfo {
            if associationsDetail.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo{}
        associationsDetail.PeerDetailInfo = append(associationsDetail.PeerDetailInfo, child)
        return &associationsDetail.PeerDetailInfo[len(associationsDetail.PeerDetailInfo)-1]
    }
    return nil
}

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range associationsDetail.PeerDetailInfo {
        children[associationsDetail.PeerDetailInfo[i].GetSegmentPath()] = &associationsDetail.PeerDetailInfo[i]
    }
    return children
}

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-ntp-enabled"] = associationsDetail.IsNtpEnabled
    leafs["sys-leap"] = associationsDetail.SysLeap
    return leafs
}

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetBundleName() string { return "cisco_ios_xr" }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetYangName() string { return "associations-detail" }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) SetParent(parent types.Entity) { associationsDetail.parent = parent }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetParent() types.Entity { return associationsDetail.parent }

func (associationsDetail *Ntp_Nodes_Node_AssociationsDetail) GetParentYangName() string { return "node" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo
// Peer info
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Leap. The type is NtpLeap.
    Leap interface{}

    // Peer's association mode. The type is NtpMode.
    PeerMode interface{}

    // Peer poll interval. The type is interface{} with range: 0..255.
    PollInterval interface{}

    // Is refclock. The type is bool.
    IsRefClock interface{}

    // Is authenticated. The type is bool.
    IsAuthenticated interface{}

    // Root delay. The type is string.
    RootDelay interface{}

    // Root dispersion. The type is string.
    RootDispersion interface{}

    // Synch distance. The type is string.
    SynchDistance interface{}

    // Precision. The type is interface{} with range: -128..127.
    Precision interface{}

    // NTP version. The type is interface{} with range: 0..255.
    Version interface{}

    // Index into filter shift register. The type is interface{} with range:
    // 0..4294967295.
    FilterIndex interface{}

    // Common peer info.
    PeerInfoCommon Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon

    // Reference time.
    RefTime Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime

    // Originate timestamp.
    OriginateTime Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime

    // Receive timestamp.
    ReceiveTime Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime

    // Transmit timestamp.
    TransmitTime Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime

    // Filter Details. The type is slice of
    // Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail.
    FilterDetail []Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail
}

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetFilter() yfilter.YFilter { return peerDetailInfo.YFilter }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) SetFilter(yf yfilter.YFilter) { peerDetailInfo.YFilter = yf }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetGoName(yname string) string {
    if yname == "leap" { return "Leap" }
    if yname == "peer-mode" { return "PeerMode" }
    if yname == "poll-interval" { return "PollInterval" }
    if yname == "is-ref-clock" { return "IsRefClock" }
    if yname == "is-authenticated" { return "IsAuthenticated" }
    if yname == "root-delay" { return "RootDelay" }
    if yname == "root-dispersion" { return "RootDispersion" }
    if yname == "synch-distance" { return "SynchDistance" }
    if yname == "precision" { return "Precision" }
    if yname == "version" { return "Version" }
    if yname == "filter-index" { return "FilterIndex" }
    if yname == "peer-info-common" { return "PeerInfoCommon" }
    if yname == "ref-time" { return "RefTime" }
    if yname == "originate-time" { return "OriginateTime" }
    if yname == "receive-time" { return "ReceiveTime" }
    if yname == "transmit-time" { return "TransmitTime" }
    if yname == "filter-detail" { return "FilterDetail" }
    return ""
}

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetSegmentPath() string {
    return "peer-detail-info"
}

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-info-common" {
        return &peerDetailInfo.PeerInfoCommon
    }
    if childYangName == "ref-time" {
        return &peerDetailInfo.RefTime
    }
    if childYangName == "originate-time" {
        return &peerDetailInfo.OriginateTime
    }
    if childYangName == "receive-time" {
        return &peerDetailInfo.ReceiveTime
    }
    if childYangName == "transmit-time" {
        return &peerDetailInfo.TransmitTime
    }
    if childYangName == "filter-detail" {
        for _, c := range peerDetailInfo.FilterDetail {
            if peerDetailInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail{}
        peerDetailInfo.FilterDetail = append(peerDetailInfo.FilterDetail, child)
        return &peerDetailInfo.FilterDetail[len(peerDetailInfo.FilterDetail)-1]
    }
    return nil
}

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-info-common"] = &peerDetailInfo.PeerInfoCommon
    children["ref-time"] = &peerDetailInfo.RefTime
    children["originate-time"] = &peerDetailInfo.OriginateTime
    children["receive-time"] = &peerDetailInfo.ReceiveTime
    children["transmit-time"] = &peerDetailInfo.TransmitTime
    for i := range peerDetailInfo.FilterDetail {
        children[peerDetailInfo.FilterDetail[i].GetSegmentPath()] = &peerDetailInfo.FilterDetail[i]
    }
    return children
}

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["leap"] = peerDetailInfo.Leap
    leafs["peer-mode"] = peerDetailInfo.PeerMode
    leafs["poll-interval"] = peerDetailInfo.PollInterval
    leafs["is-ref-clock"] = peerDetailInfo.IsRefClock
    leafs["is-authenticated"] = peerDetailInfo.IsAuthenticated
    leafs["root-delay"] = peerDetailInfo.RootDelay
    leafs["root-dispersion"] = peerDetailInfo.RootDispersion
    leafs["synch-distance"] = peerDetailInfo.SynchDistance
    leafs["precision"] = peerDetailInfo.Precision
    leafs["version"] = peerDetailInfo.Version
    leafs["filter-index"] = peerDetailInfo.FilterIndex
    return leafs
}

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetYangName() string { return "peer-detail-info" }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) SetParent(parent types.Entity) { peerDetailInfo.parent = parent }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetParent() types.Entity { return peerDetailInfo.parent }

func (peerDetailInfo *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo) GetParentYangName() string { return "associations-detail" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon
// Common peer info
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association mode with this peer. The type is NtpMode.
    HostMode interface{}

    // Is configured. The type is bool.
    IsConfigured interface{}

    // Peer Address. The type is string.
    Address interface{}

    // Peer reference ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferenceId interface{}

    // Host poll. The type is interface{} with range: 0..255.
    HostPoll interface{}

    // Reachability. The type is interface{} with range: 0..255.
    Reachability interface{}

    // Peer stratum. The type is interface{} with range: 0..255.
    Stratum interface{}

    // Peer status. The type is NtpPeerStatus.
    Status interface{}

    // Peer delay. The type is string.
    Delay interface{}

    // Peer offset. The type is string.
    Offset interface{}

    // Peer dispersion. The type is string.
    Dispersion interface{}

    // Indicates whether this is syspeer. The type is bool.
    IsSysPeer interface{}
}

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetFilter() yfilter.YFilter { return peerInfoCommon.YFilter }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) SetFilter(yf yfilter.YFilter) { peerInfoCommon.YFilter = yf }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetGoName(yname string) string {
    if yname == "host-mode" { return "HostMode" }
    if yname == "is-configured" { return "IsConfigured" }
    if yname == "address" { return "Address" }
    if yname == "reference-id" { return "ReferenceId" }
    if yname == "host-poll" { return "HostPoll" }
    if yname == "reachability" { return "Reachability" }
    if yname == "stratum" { return "Stratum" }
    if yname == "status" { return "Status" }
    if yname == "delay" { return "Delay" }
    if yname == "offset" { return "Offset" }
    if yname == "dispersion" { return "Dispersion" }
    if yname == "is-sys-peer" { return "IsSysPeer" }
    return ""
}

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetSegmentPath() string {
    return "peer-info-common"
}

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-mode"] = peerInfoCommon.HostMode
    leafs["is-configured"] = peerInfoCommon.IsConfigured
    leafs["address"] = peerInfoCommon.Address
    leafs["reference-id"] = peerInfoCommon.ReferenceId
    leafs["host-poll"] = peerInfoCommon.HostPoll
    leafs["reachability"] = peerInfoCommon.Reachability
    leafs["stratum"] = peerInfoCommon.Stratum
    leafs["status"] = peerInfoCommon.Status
    leafs["delay"] = peerInfoCommon.Delay
    leafs["offset"] = peerInfoCommon.Offset
    leafs["dispersion"] = peerInfoCommon.Dispersion
    leafs["is-sys-peer"] = peerInfoCommon.IsSysPeer
    return leafs
}

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetBundleName() string { return "cisco_ios_xr" }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetYangName() string { return "peer-info-common" }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) SetParent(parent types.Entity) { peerInfoCommon.parent = parent }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetParent() types.Entity { return peerInfoCommon.parent }

func (peerInfoCommon *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_PeerInfoCommon) GetParentYangName() string { return "peer-detail-info" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime
// Reference time
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs
}

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetFilter() yfilter.YFilter { return refTime.YFilter }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) SetFilter(yf yfilter.YFilter) { refTime.YFilter = yf }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetGoName(yname string) string {
    if yname == "sec" { return "Sec" }
    if yname == "frac-secs" { return "FracSecs" }
    return ""
}

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetSegmentPath() string {
    return "ref-time"
}

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sec" {
        return &refTime.Sec
    }
    if childYangName == "frac-secs" {
        return &refTime.FracSecs
    }
    return nil
}

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sec"] = &refTime.Sec
    children["frac-secs"] = &refTime.FracSecs
    return children
}

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetBundleName() string { return "cisco_ios_xr" }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetYangName() string { return "ref-time" }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) SetParent(parent types.Entity) { refTime.parent = parent }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetParent() types.Entity { return refTime.parent }

func (refTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime) GetParentYangName() string { return "peer-detail-info" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetFilter() yfilter.YFilter { return sec.YFilter }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) SetFilter(yf yfilter.YFilter) { sec.YFilter = yf }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetGoName(yname string) string {
    if yname == "int" { return "Int" }
    return ""
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetSegmentPath() string {
    return "sec"
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["int"] = sec.Int
    return leafs
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetBundleName() string { return "cisco_ios_xr" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetYangName() string { return "sec" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) SetParent(parent types.Entity) { sec.parent = parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetParent() types.Entity { return sec.parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_Sec) GetParentYangName() string { return "ref-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetFilter() yfilter.YFilter { return fracSecs.YFilter }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) SetFilter(yf yfilter.YFilter) { fracSecs.YFilter = yf }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetGoName(yname string) string {
    if yname == "frac" { return "Frac" }
    return ""
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetSegmentPath() string {
    return "frac-secs"
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frac"] = fracSecs.Frac
    return leafs
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetBundleName() string { return "cisco_ios_xr" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetYangName() string { return "frac-secs" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) SetParent(parent types.Entity) { fracSecs.parent = parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetParent() types.Entity { return fracSecs.parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_RefTime_FracSecs) GetParentYangName() string { return "ref-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime
// Originate timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs
}

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetFilter() yfilter.YFilter { return originateTime.YFilter }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) SetFilter(yf yfilter.YFilter) { originateTime.YFilter = yf }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetGoName(yname string) string {
    if yname == "sec" { return "Sec" }
    if yname == "frac-secs" { return "FracSecs" }
    return ""
}

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetSegmentPath() string {
    return "originate-time"
}

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sec" {
        return &originateTime.Sec
    }
    if childYangName == "frac-secs" {
        return &originateTime.FracSecs
    }
    return nil
}

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sec"] = &originateTime.Sec
    children["frac-secs"] = &originateTime.FracSecs
    return children
}

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetBundleName() string { return "cisco_ios_xr" }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetYangName() string { return "originate-time" }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) SetParent(parent types.Entity) { originateTime.parent = parent }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetParent() types.Entity { return originateTime.parent }

func (originateTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime) GetParentYangName() string { return "peer-detail-info" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetFilter() yfilter.YFilter { return sec.YFilter }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) SetFilter(yf yfilter.YFilter) { sec.YFilter = yf }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetGoName(yname string) string {
    if yname == "int" { return "Int" }
    return ""
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetSegmentPath() string {
    return "sec"
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["int"] = sec.Int
    return leafs
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetBundleName() string { return "cisco_ios_xr" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetYangName() string { return "sec" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) SetParent(parent types.Entity) { sec.parent = parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetParent() types.Entity { return sec.parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_Sec) GetParentYangName() string { return "originate-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetFilter() yfilter.YFilter { return fracSecs.YFilter }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) SetFilter(yf yfilter.YFilter) { fracSecs.YFilter = yf }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetGoName(yname string) string {
    if yname == "frac" { return "Frac" }
    return ""
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetSegmentPath() string {
    return "frac-secs"
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frac"] = fracSecs.Frac
    return leafs
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetBundleName() string { return "cisco_ios_xr" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetYangName() string { return "frac-secs" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) SetParent(parent types.Entity) { fracSecs.parent = parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetParent() types.Entity { return fracSecs.parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_OriginateTime_FracSecs) GetParentYangName() string { return "originate-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime
// Receive timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs
}

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetFilter() yfilter.YFilter { return receiveTime.YFilter }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) SetFilter(yf yfilter.YFilter) { receiveTime.YFilter = yf }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetGoName(yname string) string {
    if yname == "sec" { return "Sec" }
    if yname == "frac-secs" { return "FracSecs" }
    return ""
}

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetSegmentPath() string {
    return "receive-time"
}

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sec" {
        return &receiveTime.Sec
    }
    if childYangName == "frac-secs" {
        return &receiveTime.FracSecs
    }
    return nil
}

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sec"] = &receiveTime.Sec
    children["frac-secs"] = &receiveTime.FracSecs
    return children
}

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetBundleName() string { return "cisco_ios_xr" }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetYangName() string { return "receive-time" }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) SetParent(parent types.Entity) { receiveTime.parent = parent }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetParent() types.Entity { return receiveTime.parent }

func (receiveTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime) GetParentYangName() string { return "peer-detail-info" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetFilter() yfilter.YFilter { return sec.YFilter }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) SetFilter(yf yfilter.YFilter) { sec.YFilter = yf }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetGoName(yname string) string {
    if yname == "int" { return "Int" }
    return ""
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetSegmentPath() string {
    return "sec"
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["int"] = sec.Int
    return leafs
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetBundleName() string { return "cisco_ios_xr" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetYangName() string { return "sec" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) SetParent(parent types.Entity) { sec.parent = parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetParent() types.Entity { return sec.parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_Sec) GetParentYangName() string { return "receive-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetFilter() yfilter.YFilter { return fracSecs.YFilter }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) SetFilter(yf yfilter.YFilter) { fracSecs.YFilter = yf }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetGoName(yname string) string {
    if yname == "frac" { return "Frac" }
    return ""
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetSegmentPath() string {
    return "frac-secs"
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frac"] = fracSecs.Frac
    return leafs
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetBundleName() string { return "cisco_ios_xr" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetYangName() string { return "frac-secs" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) SetParent(parent types.Entity) { fracSecs.parent = parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetParent() types.Entity { return fracSecs.parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_ReceiveTime_FracSecs) GetParentYangName() string { return "receive-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime
// Transmit timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs
}

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetFilter() yfilter.YFilter { return transmitTime.YFilter }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) SetFilter(yf yfilter.YFilter) { transmitTime.YFilter = yf }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetGoName(yname string) string {
    if yname == "sec" { return "Sec" }
    if yname == "frac-secs" { return "FracSecs" }
    return ""
}

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetSegmentPath() string {
    return "transmit-time"
}

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sec" {
        return &transmitTime.Sec
    }
    if childYangName == "frac-secs" {
        return &transmitTime.FracSecs
    }
    return nil
}

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sec"] = &transmitTime.Sec
    children["frac-secs"] = &transmitTime.FracSecs
    return children
}

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetBundleName() string { return "cisco_ios_xr" }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetYangName() string { return "transmit-time" }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) SetParent(parent types.Entity) { transmitTime.parent = parent }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetParent() types.Entity { return transmitTime.parent }

func (transmitTime *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime) GetParentYangName() string { return "peer-detail-info" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetFilter() yfilter.YFilter { return sec.YFilter }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) SetFilter(yf yfilter.YFilter) { sec.YFilter = yf }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetGoName(yname string) string {
    if yname == "int" { return "Int" }
    return ""
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetSegmentPath() string {
    return "sec"
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["int"] = sec.Int
    return leafs
}

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetBundleName() string { return "cisco_ios_xr" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetYangName() string { return "sec" }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) SetParent(parent types.Entity) { sec.parent = parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetParent() types.Entity { return sec.parent }

func (sec *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_Sec) GetParentYangName() string { return "transmit-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetFilter() yfilter.YFilter { return fracSecs.YFilter }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) SetFilter(yf yfilter.YFilter) { fracSecs.YFilter = yf }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetGoName(yname string) string {
    if yname == "frac" { return "Frac" }
    return ""
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetSegmentPath() string {
    return "frac-secs"
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frac"] = fracSecs.Frac
    return leafs
}

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetBundleName() string { return "cisco_ios_xr" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetYangName() string { return "frac-secs" }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) SetParent(parent types.Entity) { fracSecs.parent = parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetParent() types.Entity { return fracSecs.parent }

func (fracSecs *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_TransmitTime_FracSecs) GetParentYangName() string { return "transmit-time" }

// Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail
// Filter Details
type Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // filter delay. The type is string.
    FilterDelay interface{}

    // filter offset. The type is string.
    FilterOffset interface{}

    // filter disp. The type is string.
    FilterDisp interface{}
}

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetFilter() yfilter.YFilter { return filterDetail.YFilter }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) SetFilter(yf yfilter.YFilter) { filterDetail.YFilter = yf }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetGoName(yname string) string {
    if yname == "filter-delay" { return "FilterDelay" }
    if yname == "filter-offset" { return "FilterOffset" }
    if yname == "filter-disp" { return "FilterDisp" }
    return ""
}

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetSegmentPath() string {
    return "filter-detail"
}

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["filter-delay"] = filterDetail.FilterDelay
    leafs["filter-offset"] = filterDetail.FilterOffset
    leafs["filter-disp"] = filterDetail.FilterDisp
    return leafs
}

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetBundleName() string { return "cisco_ios_xr" }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetYangName() string { return "filter-detail" }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) SetParent(parent types.Entity) { filterDetail.parent = parent }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetParent() types.Entity { return filterDetail.parent }

func (filterDetail *Ntp_Nodes_Node_AssociationsDetail_PeerDetailInfo_FilterDetail) GetParentYangName() string { return "peer-detail-info" }

// Ntp_Nodes_Node_Status
// Status of NTP peer(s)
type Ntp_Nodes_Node_Status struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is NTP enabled. The type is bool.
    IsNtpEnabled interface{}

    // Peer dispersion. The type is string.
    SysDispersion interface{}

    // Clock offset. The type is string.
    SysOffset interface{}

    // Clock period in nanosecs. The type is interface{} with range:
    // 0..4294967295. Units are nanosecond.
    ClockPeriod interface{}

    // leap. The type is NtpLeap.
    SysLeap interface{}

    // Precision. The type is interface{} with range: -128..127.
    SysPrecision interface{}

    // Stratum. The type is interface{} with range: 0..255.
    SysStratum interface{}

    // Reference clock ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SysRefId interface{}

    // Root delay. The type is string.
    SysRootDelay interface{}

    // Root dispersion. The type is string.
    SysRootDispersion interface{}

    // Loop Filter State. The type is NtpLoopFilterState.
    LoopFilterState interface{}

    // Peer poll interval. The type is interface{} with range: 0..255.
    PollInterval interface{}

    // Is clock updated. The type is ClockUpdateNode.
    IsUpdated interface{}

    // Last Update. The type is interface{} with range: -2147483648..2147483647.
    LastUpdate interface{}

    // Reference time.
    SysRefTime Ntp_Nodes_Node_Status_SysRefTime

    // System Drift.
    SysDrift Ntp_Nodes_Node_Status_SysDrift
}

func (status *Ntp_Nodes_Node_Status) GetFilter() yfilter.YFilter { return status.YFilter }

func (status *Ntp_Nodes_Node_Status) SetFilter(yf yfilter.YFilter) { status.YFilter = yf }

func (status *Ntp_Nodes_Node_Status) GetGoName(yname string) string {
    if yname == "is-ntp-enabled" { return "IsNtpEnabled" }
    if yname == "sys-dispersion" { return "SysDispersion" }
    if yname == "sys-offset" { return "SysOffset" }
    if yname == "clock-period" { return "ClockPeriod" }
    if yname == "sys-leap" { return "SysLeap" }
    if yname == "sys-precision" { return "SysPrecision" }
    if yname == "sys-stratum" { return "SysStratum" }
    if yname == "sys-ref-id" { return "SysRefId" }
    if yname == "sys-root-delay" { return "SysRootDelay" }
    if yname == "sys-root-dispersion" { return "SysRootDispersion" }
    if yname == "loop-filter-state" { return "LoopFilterState" }
    if yname == "poll-interval" { return "PollInterval" }
    if yname == "is-updated" { return "IsUpdated" }
    if yname == "last-update" { return "LastUpdate" }
    if yname == "sys-ref-time" { return "SysRefTime" }
    if yname == "sys-drift" { return "SysDrift" }
    return ""
}

func (status *Ntp_Nodes_Node_Status) GetSegmentPath() string {
    return "status"
}

func (status *Ntp_Nodes_Node_Status) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sys-ref-time" {
        return &status.SysRefTime
    }
    if childYangName == "sys-drift" {
        return &status.SysDrift
    }
    return nil
}

func (status *Ntp_Nodes_Node_Status) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sys-ref-time"] = &status.SysRefTime
    children["sys-drift"] = &status.SysDrift
    return children
}

func (status *Ntp_Nodes_Node_Status) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-ntp-enabled"] = status.IsNtpEnabled
    leafs["sys-dispersion"] = status.SysDispersion
    leafs["sys-offset"] = status.SysOffset
    leafs["clock-period"] = status.ClockPeriod
    leafs["sys-leap"] = status.SysLeap
    leafs["sys-precision"] = status.SysPrecision
    leafs["sys-stratum"] = status.SysStratum
    leafs["sys-ref-id"] = status.SysRefId
    leafs["sys-root-delay"] = status.SysRootDelay
    leafs["sys-root-dispersion"] = status.SysRootDispersion
    leafs["loop-filter-state"] = status.LoopFilterState
    leafs["poll-interval"] = status.PollInterval
    leafs["is-updated"] = status.IsUpdated
    leafs["last-update"] = status.LastUpdate
    return leafs
}

func (status *Ntp_Nodes_Node_Status) GetBundleName() string { return "cisco_ios_xr" }

func (status *Ntp_Nodes_Node_Status) GetYangName() string { return "status" }

func (status *Ntp_Nodes_Node_Status) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (status *Ntp_Nodes_Node_Status) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (status *Ntp_Nodes_Node_Status) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (status *Ntp_Nodes_Node_Status) SetParent(parent types.Entity) { status.parent = parent }

func (status *Ntp_Nodes_Node_Status) GetParent() types.Entity { return status.parent }

func (status *Ntp_Nodes_Node_Status) GetParentYangName() string { return "node" }

// Ntp_Nodes_Node_Status_SysRefTime
// Reference time
type Ntp_Nodes_Node_Status_SysRefTime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Nodes_Node_Status_SysRefTime_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Nodes_Node_Status_SysRefTime_FracSecs
}

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetFilter() yfilter.YFilter { return sysRefTime.YFilter }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) SetFilter(yf yfilter.YFilter) { sysRefTime.YFilter = yf }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetGoName(yname string) string {
    if yname == "sec" { return "Sec" }
    if yname == "frac-secs" { return "FracSecs" }
    return ""
}

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetSegmentPath() string {
    return "sys-ref-time"
}

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sec" {
        return &sysRefTime.Sec
    }
    if childYangName == "frac-secs" {
        return &sysRefTime.FracSecs
    }
    return nil
}

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sec"] = &sysRefTime.Sec
    children["frac-secs"] = &sysRefTime.FracSecs
    return children
}

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetBundleName() string { return "cisco_ios_xr" }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetYangName() string { return "sys-ref-time" }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) SetParent(parent types.Entity) { sysRefTime.parent = parent }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetParent() types.Entity { return sysRefTime.parent }

func (sysRefTime *Ntp_Nodes_Node_Status_SysRefTime) GetParentYangName() string { return "status" }

// Ntp_Nodes_Node_Status_SysRefTime_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Nodes_Node_Status_SysRefTime_Sec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetFilter() yfilter.YFilter { return sec.YFilter }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) SetFilter(yf yfilter.YFilter) { sec.YFilter = yf }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetGoName(yname string) string {
    if yname == "int" { return "Int" }
    return ""
}

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetSegmentPath() string {
    return "sec"
}

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["int"] = sec.Int
    return leafs
}

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetBundleName() string { return "cisco_ios_xr" }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetYangName() string { return "sec" }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) SetParent(parent types.Entity) { sec.parent = parent }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetParent() types.Entity { return sec.parent }

func (sec *Ntp_Nodes_Node_Status_SysRefTime_Sec) GetParentYangName() string { return "sys-ref-time" }

// Ntp_Nodes_Node_Status_SysRefTime_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Nodes_Node_Status_SysRefTime_FracSecs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetFilter() yfilter.YFilter { return fracSecs.YFilter }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) SetFilter(yf yfilter.YFilter) { fracSecs.YFilter = yf }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetGoName(yname string) string {
    if yname == "frac" { return "Frac" }
    return ""
}

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetSegmentPath() string {
    return "frac-secs"
}

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frac"] = fracSecs.Frac
    return leafs
}

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetBundleName() string { return "cisco_ios_xr" }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetYangName() string { return "frac-secs" }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) SetParent(parent types.Entity) { fracSecs.parent = parent }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetParent() types.Entity { return fracSecs.parent }

func (fracSecs *Ntp_Nodes_Node_Status_SysRefTime_FracSecs) GetParentYangName() string { return "sys-ref-time" }

// Ntp_Nodes_Node_Status_SysDrift
// System Drift
type Ntp_Nodes_Node_Status_SysDrift struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Second part in 64-bit NTP timestamp.
    Sec Ntp_Nodes_Node_Status_SysDrift_Sec

    // Fractional part in 64-bit NTP timestamp.
    FracSecs Ntp_Nodes_Node_Status_SysDrift_FracSecs
}

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetFilter() yfilter.YFilter { return sysDrift.YFilter }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) SetFilter(yf yfilter.YFilter) { sysDrift.YFilter = yf }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetGoName(yname string) string {
    if yname == "sec" { return "Sec" }
    if yname == "frac-secs" { return "FracSecs" }
    return ""
}

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetSegmentPath() string {
    return "sys-drift"
}

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sec" {
        return &sysDrift.Sec
    }
    if childYangName == "frac-secs" {
        return &sysDrift.FracSecs
    }
    return nil
}

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sec"] = &sysDrift.Sec
    children["frac-secs"] = &sysDrift.FracSecs
    return children
}

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetBundleName() string { return "cisco_ios_xr" }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetYangName() string { return "sys-drift" }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) SetParent(parent types.Entity) { sysDrift.parent = parent }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetParent() types.Entity { return sysDrift.parent }

func (sysDrift *Ntp_Nodes_Node_Status_SysDrift) GetParentYangName() string { return "status" }

// Ntp_Nodes_Node_Status_SysDrift_Sec
// Second part in 64-bit NTP timestamp
type Ntp_Nodes_Node_Status_SysDrift_Sec struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Integer format in NTP reference code. The type is interface{} with range:
    // 0..4294967295.
    Int interface{}
}

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetFilter() yfilter.YFilter { return sec.YFilter }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) SetFilter(yf yfilter.YFilter) { sec.YFilter = yf }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetGoName(yname string) string {
    if yname == "int" { return "Int" }
    return ""
}

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetSegmentPath() string {
    return "sec"
}

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["int"] = sec.Int
    return leafs
}

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetBundleName() string { return "cisco_ios_xr" }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetYangName() string { return "sec" }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) SetParent(parent types.Entity) { sec.parent = parent }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetParent() types.Entity { return sec.parent }

func (sec *Ntp_Nodes_Node_Status_SysDrift_Sec) GetParentYangName() string { return "sys-drift" }

// Ntp_Nodes_Node_Status_SysDrift_FracSecs
// Fractional part in 64-bit NTP timestamp
type Ntp_Nodes_Node_Status_SysDrift_FracSecs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fractional format in NTP reference code. The type is interface{} with
    // range: 0..4294967295.
    Frac interface{}
}

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetFilter() yfilter.YFilter { return fracSecs.YFilter }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) SetFilter(yf yfilter.YFilter) { fracSecs.YFilter = yf }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetGoName(yname string) string {
    if yname == "frac" { return "Frac" }
    return ""
}

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetSegmentPath() string {
    return "frac-secs"
}

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["frac"] = fracSecs.Frac
    return leafs
}

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetBundleName() string { return "cisco_ios_xr" }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetYangName() string { return "frac-secs" }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) SetParent(parent types.Entity) { fracSecs.parent = parent }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetParent() types.Entity { return fracSecs.parent }

func (fracSecs *Ntp_Nodes_Node_Status_SysDrift_FracSecs) GetParentYangName() string { return "sys-drift" }

// Ntp_Nodes_Node_Associations
// NTP Associations information
type Ntp_Nodes_Node_Associations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is NTP enabled. The type is bool.
    IsNtpEnabled interface{}

    // Leap. The type is NtpLeap.
    SysLeap interface{}

    // Peer info. The type is slice of
    // Ntp_Nodes_Node_Associations_PeerSummaryInfo.
    PeerSummaryInfo []Ntp_Nodes_Node_Associations_PeerSummaryInfo
}

func (associations *Ntp_Nodes_Node_Associations) GetFilter() yfilter.YFilter { return associations.YFilter }

func (associations *Ntp_Nodes_Node_Associations) SetFilter(yf yfilter.YFilter) { associations.YFilter = yf }

func (associations *Ntp_Nodes_Node_Associations) GetGoName(yname string) string {
    if yname == "is-ntp-enabled" { return "IsNtpEnabled" }
    if yname == "sys-leap" { return "SysLeap" }
    if yname == "peer-summary-info" { return "PeerSummaryInfo" }
    return ""
}

func (associations *Ntp_Nodes_Node_Associations) GetSegmentPath() string {
    return "associations"
}

func (associations *Ntp_Nodes_Node_Associations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-summary-info" {
        for _, c := range associations.PeerSummaryInfo {
            if associations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ntp_Nodes_Node_Associations_PeerSummaryInfo{}
        associations.PeerSummaryInfo = append(associations.PeerSummaryInfo, child)
        return &associations.PeerSummaryInfo[len(associations.PeerSummaryInfo)-1]
    }
    return nil
}

func (associations *Ntp_Nodes_Node_Associations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range associations.PeerSummaryInfo {
        children[associations.PeerSummaryInfo[i].GetSegmentPath()] = &associations.PeerSummaryInfo[i]
    }
    return children
}

func (associations *Ntp_Nodes_Node_Associations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-ntp-enabled"] = associations.IsNtpEnabled
    leafs["sys-leap"] = associations.SysLeap
    return leafs
}

func (associations *Ntp_Nodes_Node_Associations) GetBundleName() string { return "cisco_ios_xr" }

func (associations *Ntp_Nodes_Node_Associations) GetYangName() string { return "associations" }

func (associations *Ntp_Nodes_Node_Associations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (associations *Ntp_Nodes_Node_Associations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (associations *Ntp_Nodes_Node_Associations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (associations *Ntp_Nodes_Node_Associations) SetParent(parent types.Entity) { associations.parent = parent }

func (associations *Ntp_Nodes_Node_Associations) GetParent() types.Entity { return associations.parent }

func (associations *Ntp_Nodes_Node_Associations) GetParentYangName() string { return "node" }

// Ntp_Nodes_Node_Associations_PeerSummaryInfo
// Peer info
type Ntp_Nodes_Node_Associations_PeerSummaryInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time since last frame received (-1=none). The type is interface{} with
    // range: -2147483648..2147483647.
    TimeSince interface{}

    // Common peer info.
    PeerInfoCommon Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon
}

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetFilter() yfilter.YFilter { return peerSummaryInfo.YFilter }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) SetFilter(yf yfilter.YFilter) { peerSummaryInfo.YFilter = yf }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetGoName(yname string) string {
    if yname == "time-since" { return "TimeSince" }
    if yname == "peer-info-common" { return "PeerInfoCommon" }
    return ""
}

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetSegmentPath() string {
    return "peer-summary-info"
}

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-info-common" {
        return &peerSummaryInfo.PeerInfoCommon
    }
    return nil
}

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-info-common"] = &peerSummaryInfo.PeerInfoCommon
    return children
}

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["time-since"] = peerSummaryInfo.TimeSince
    return leafs
}

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetYangName() string { return "peer-summary-info" }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) SetParent(parent types.Entity) { peerSummaryInfo.parent = parent }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetParent() types.Entity { return peerSummaryInfo.parent }

func (peerSummaryInfo *Ntp_Nodes_Node_Associations_PeerSummaryInfo) GetParentYangName() string { return "associations" }

// Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon
// Common peer info
type Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association mode with this peer. The type is NtpMode.
    HostMode interface{}

    // Is configured. The type is bool.
    IsConfigured interface{}

    // Peer Address. The type is string.
    Address interface{}

    // Peer reference ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReferenceId interface{}

    // Host poll. The type is interface{} with range: 0..255.
    HostPoll interface{}

    // Reachability. The type is interface{} with range: 0..255.
    Reachability interface{}

    // Peer stratum. The type is interface{} with range: 0..255.
    Stratum interface{}

    // Peer status. The type is NtpPeerStatus.
    Status interface{}

    // Peer delay. The type is string.
    Delay interface{}

    // Peer offset. The type is string.
    Offset interface{}

    // Peer dispersion. The type is string.
    Dispersion interface{}

    // Indicates whether this is syspeer. The type is bool.
    IsSysPeer interface{}
}

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetFilter() yfilter.YFilter { return peerInfoCommon.YFilter }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) SetFilter(yf yfilter.YFilter) { peerInfoCommon.YFilter = yf }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetGoName(yname string) string {
    if yname == "host-mode" { return "HostMode" }
    if yname == "is-configured" { return "IsConfigured" }
    if yname == "address" { return "Address" }
    if yname == "reference-id" { return "ReferenceId" }
    if yname == "host-poll" { return "HostPoll" }
    if yname == "reachability" { return "Reachability" }
    if yname == "stratum" { return "Stratum" }
    if yname == "status" { return "Status" }
    if yname == "delay" { return "Delay" }
    if yname == "offset" { return "Offset" }
    if yname == "dispersion" { return "Dispersion" }
    if yname == "is-sys-peer" { return "IsSysPeer" }
    return ""
}

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetSegmentPath() string {
    return "peer-info-common"
}

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["host-mode"] = peerInfoCommon.HostMode
    leafs["is-configured"] = peerInfoCommon.IsConfigured
    leafs["address"] = peerInfoCommon.Address
    leafs["reference-id"] = peerInfoCommon.ReferenceId
    leafs["host-poll"] = peerInfoCommon.HostPoll
    leafs["reachability"] = peerInfoCommon.Reachability
    leafs["stratum"] = peerInfoCommon.Stratum
    leafs["status"] = peerInfoCommon.Status
    leafs["delay"] = peerInfoCommon.Delay
    leafs["offset"] = peerInfoCommon.Offset
    leafs["dispersion"] = peerInfoCommon.Dispersion
    leafs["is-sys-peer"] = peerInfoCommon.IsSysPeer
    return leafs
}

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetBundleName() string { return "cisco_ios_xr" }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetYangName() string { return "peer-info-common" }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) SetParent(parent types.Entity) { peerInfoCommon.parent = parent }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetParent() types.Entity { return peerInfoCommon.parent }

func (peerInfoCommon *Ntp_Nodes_Node_Associations_PeerSummaryInfo_PeerInfoCommon) GetParentYangName() string { return "peer-summary-info" }

