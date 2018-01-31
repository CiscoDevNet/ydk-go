// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ppp: PPP operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ma_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ma_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ppp-ma-oper ppp}", reflect.TypeOf(Ppp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ppp-ma-oper:ppp", reflect.TypeOf(Ppp{}))
}

// NcpIdent represents Ncp ident
type NcpIdent string

const (
    // CDP control protocol
    NcpIdent_cdpcp NcpIdent = "cdpcp"

    // IPv4 control protocol
    NcpIdent_ipcp NcpIdent = "ipcp"

    // IPv4 Interworking control protocol
    NcpIdent_ipcpiw NcpIdent = "ipcpiw"

    // IPv6 control protocol
    NcpIdent_ipv6cp NcpIdent = "ipv6cp"

    // MPLS control protocol
    NcpIdent_mplscp NcpIdent = "mplscp"

    // OSI (CLNS) control protocol
    NcpIdent_osicp NcpIdent = "osicp"
)

// PppSsoFsmState represents Ppp sso fsm state
type PppSsoFsmState string

const (
    // Not ready
    PppSsoFsmState_ppp_sso_state_not_ready_0 PppSsoFsmState = "ppp-sso-state-not-ready-0"

    // S UnNegd
    PppSsoFsmState_ppp_sso_state_standby_unnegd_1 PppSsoFsmState = "ppp-sso-state-standby-unnegd-1"

    // A Down
    PppSsoFsmState_ppp_sso_state_active_down_2 PppSsoFsmState = "ppp-sso-state-active-down-2"

    // Deactivating
    PppSsoFsmState_ppp_sso_state_deactivating_3 PppSsoFsmState = "ppp-sso-state-deactivating-3"

    // A UnNegd
    PppSsoFsmState_ppp_sso_state_active_unnegd_4 PppSsoFsmState = "ppp-sso-state-active-unnegd-4"

    // S Negd
    PppSsoFsmState_ppp_sso_state_standby_negd_5 PppSsoFsmState = "ppp-sso-state-standby-negd-5"

    // Activating
    PppSsoFsmState_ppp_sso_state_activating_6 PppSsoFsmState = "ppp-sso-state-activating-6"

    // A Negd
    PppSsoFsmState_ppp_sso_state_active_negd_7 PppSsoFsmState = "ppp-sso-state-active-negd-7"
)

// PppIphcCompression represents IPHC compression type
type PppIphcCompression string

const (
    // None
    PppIphcCompression_ppp_iphc_compression_fmt_none PppIphcCompression = "ppp-iphc-compression-fmt-none"

    // VJ
    PppIphcCompression_ppp_iphc_compression_fmt_vj PppIphcCompression = "ppp-iphc-compression-fmt-vj"

    // IETF
    PppIphcCompression_ppp_iphc_compression_fmt_ietf PppIphcCompression = "ppp-iphc-compression-fmt-ietf"

    // IPHC
    PppIphcCompression_ppp_iphc_compression_fmt_iphc PppIphcCompression = "ppp-iphc-compression-fmt-iphc"

    // CISCO
    PppIphcCompression_ppp_iphc_compression_fmt_cisco PppIphcCompression = "ppp-iphc-compression-fmt-cisco"
)

// PppFsmState represents Ppp fsm state
type PppFsmState string

const (
    // Connection Idle
    PppFsmState_ppp_fsm_state_initial_0 PppFsmState = "ppp-fsm-state-initial-0"

    // This layer required, but lower layer down
    PppFsmState_ppp_fsm_state_starting_1 PppFsmState = "ppp-fsm-state-starting-1"

    // Lower layer up, but this layer not required
    PppFsmState_ppp_fsm_state_closed_2 PppFsmState = "ppp-fsm-state-closed-2"

    // Listening for a Config Request
    PppFsmState_ppp_fsm_state_stopped_3 PppFsmState = "ppp-fsm-state-stopped-3"

    // Shutting down due to local change
    PppFsmState_ppp_fsm_state_closing_4 PppFsmState = "ppp-fsm-state-closing-4"

    // Shutting down due to peer's actions
    PppFsmState_ppp_fsm_state_stopping_5 PppFsmState = "ppp-fsm-state-stopping-5"

    // Config Request Sent
    PppFsmState_ppp_fsm_state_req_sent_6 PppFsmState = "ppp-fsm-state-req-sent-6"

    // Config Ack Received
    PppFsmState_ppp_fsm_state_ack_rcvd_7 PppFsmState = "ppp-fsm-state-ack-rcvd-7"

    // Config Ack Sent
    PppFsmState_ppp_fsm_state_ack_sent_8 PppFsmState = "ppp-fsm-state-ack-sent-8"

    // Connection Open
    PppFsmState_ppp_fsm_state_opened_9 PppFsmState = "ppp-fsm-state-opened-9"
)

// PppLcpMpMbrState represents MP member states
type PppLcpMpMbrState string

const (
    // Detached member
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_detached PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-detached"

    // LCP has not been negotiated
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_lcp_not_negotiated PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-lcp-not-negotiated"

    // Link Noise detected
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_link_noise PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-link-noise"

    // Multilink Bundle is shutdown
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_bundle_shutdown PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-bundle-shutdown"

    // MRRU has been rejected
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_mrru_rejected PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-mrru-rejected"

    // MRRU mismatch
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_mrru_mismatch PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-mrru-mismatch"

    // ED mismatch
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_ed_mismatch PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-ed-mismatch"

    // Authenticated name mismatch
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_auth_name_mismatch PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-auth-name-mismatch"

    // MCMP option rejected by peer
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_mcmp_rejected PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-mcmp-rejected"

    // MCMP option not negotiated
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_mcmp_not_negotiated PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-mcmp-not-negotiated"

    // Local MCMP class mismatch
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_mcmp_local_mismatch PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-mcmp-local-mismatch"

    // Peer MCMP class mismatch
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_mcmp_peer_mismatch PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-mcmp-peer-mismatch"

    // SSO Standby up
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_standby_up PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-standby-up"

    // Active member
    PppLcpMpMbrState_ppp_lcp_mp_mbr_state_active PppLcpMpMbrState = "ppp-lcp-mp-mbr-state-active"
)

// Ppp
// PPP operational data
type Ppp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per node PPP operational data.
    Nodes Ppp_Nodes
}

func (ppp *Ppp) GetFilter() yfilter.YFilter { return ppp.YFilter }

func (ppp *Ppp) SetFilter(yf yfilter.YFilter) { ppp.YFilter = yf }

func (ppp *Ppp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (ppp *Ppp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ppp-ma-oper:ppp"
}

func (ppp *Ppp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &ppp.Nodes
    }
    return nil
}

func (ppp *Ppp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &ppp.Nodes
    return children
}

func (ppp *Ppp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ppp *Ppp) GetBundleName() string { return "cisco_ios_xr" }

func (ppp *Ppp) GetYangName() string { return "ppp" }

func (ppp *Ppp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ppp *Ppp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ppp *Ppp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ppp *Ppp) SetParent(parent types.Entity) { ppp.parent = parent }

func (ppp *Ppp) GetParent() types.Entity { return ppp.parent }

func (ppp *Ppp) GetParentYangName() string { return "Cisco-IOS-XR-ppp-ma-oper" }

// Ppp_Nodes
// Per node PPP operational data
type Ppp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The PPP operational data for a particular node. The type is slice of
    // Ppp_Nodes_Node.
    Node []Ppp_Nodes_Node
}

func (nodes *Ppp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Ppp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Ppp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Ppp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Ppp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Ppp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Ppp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Ppp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Ppp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Ppp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Ppp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Ppp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Ppp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Ppp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Ppp_Nodes) GetParentYangName() string { return "ppp" }

// Ppp_Nodes_Node
// The PPP operational data for a particular node
type Ppp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the node. The type is string
    // with pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // PPP statistics data for a particular node.
    Statistics Ppp_Nodes_Node_Statistics

    // Per interface PPP operational data.
    NodeInterfaces Ppp_Nodes_Node_NodeInterfaces

    // PPP SSO Alert data for a particular node.
    SsoAlerts Ppp_Nodes_Node_SsoAlerts

    // Per interface PPP operational statistics.
    NodeInterfaceStatistics Ppp_Nodes_Node_NodeInterfaceStatistics

    // Summarized PPP SSO data for a particular node.
    SsoSummary Ppp_Nodes_Node_SsoSummary

    // PPP SSO Group data for a particular node.
    SsoGroups Ppp_Nodes_Node_SsoGroups

    // Summarized PPP data for a particular node.
    Summary Ppp_Nodes_Node_Summary
}

func (node *Ppp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Ppp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Ppp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statistics" { return "Statistics" }
    if yname == "node-interfaces" { return "NodeInterfaces" }
    if yname == "sso-alerts" { return "SsoAlerts" }
    if yname == "node-interface-statistics" { return "NodeInterfaceStatistics" }
    if yname == "sso-summary" { return "SsoSummary" }
    if yname == "sso-groups" { return "SsoGroups" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (node *Ppp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Ppp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    if childYangName == "node-interfaces" {
        return &node.NodeInterfaces
    }
    if childYangName == "sso-alerts" {
        return &node.SsoAlerts
    }
    if childYangName == "node-interface-statistics" {
        return &node.NodeInterfaceStatistics
    }
    if childYangName == "sso-summary" {
        return &node.SsoSummary
    }
    if childYangName == "sso-groups" {
        return &node.SsoGroups
    }
    if childYangName == "summary" {
        return &node.Summary
    }
    return nil
}

func (node *Ppp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    children["node-interfaces"] = &node.NodeInterfaces
    children["sso-alerts"] = &node.SsoAlerts
    children["node-interface-statistics"] = &node.NodeInterfaceStatistics
    children["sso-summary"] = &node.SsoSummary
    children["sso-groups"] = &node.SsoGroups
    children["summary"] = &node.Summary
    return children
}

func (node *Ppp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Ppp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Ppp_Nodes_Node) GetYangName() string { return "node" }

func (node *Ppp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Ppp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Ppp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Ppp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Ppp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Ppp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Ppp_Nodes_Node_Statistics
// PPP statistics data for a particular node
type Ppp_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPP LCP Statistics.
    LcpStatistics Ppp_Nodes_Node_Statistics_LcpStatistics

    // PPP Authentication statistics.
    AuthenticationStatistics Ppp_Nodes_Node_Statistics_AuthenticationStatistics

    // Array of PPP NCP Statistics. The type is slice of
    // Ppp_Nodes_Node_Statistics_NcpStatisticsArray.
    NcpStatisticsArray []Ppp_Nodes_Node_Statistics_NcpStatisticsArray
}

func (statistics *Ppp_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Ppp_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Ppp_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "lcp-statistics" { return "LcpStatistics" }
    if yname == "authentication-statistics" { return "AuthenticationStatistics" }
    if yname == "ncp-statistics-array" { return "NcpStatisticsArray" }
    return ""
}

func (statistics *Ppp_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Ppp_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lcp-statistics" {
        return &statistics.LcpStatistics
    }
    if childYangName == "authentication-statistics" {
        return &statistics.AuthenticationStatistics
    }
    if childYangName == "ncp-statistics-array" {
        for _, c := range statistics.NcpStatisticsArray {
            if statistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_Statistics_NcpStatisticsArray{}
        statistics.NcpStatisticsArray = append(statistics.NcpStatisticsArray, child)
        return &statistics.NcpStatisticsArray[len(statistics.NcpStatisticsArray)-1]
    }
    return nil
}

func (statistics *Ppp_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lcp-statistics"] = &statistics.LcpStatistics
    children["authentication-statistics"] = &statistics.AuthenticationStatistics
    for i := range statistics.NcpStatisticsArray {
        children[statistics.NcpStatisticsArray[i].GetSegmentPath()] = &statistics.NcpStatisticsArray[i]
    }
    return children
}

func (statistics *Ppp_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Ppp_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Ppp_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Ppp_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Ppp_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Ppp_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Ppp_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Ppp_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Ppp_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_Statistics_LcpStatistics
// PPP LCP Statistics
type Ppp_Nodes_Node_Statistics_LcpStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Conf Req Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfReqSent interface{}

    // Conf Req Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfReqRcvd interface{}

    // Conf Ack Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfAckSent interface{}

    // Conf Ack Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfAckRcvd interface{}

    // Conf Nak Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfNakSent interface{}

    // Conf Nak Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfNakRcvd interface{}

    // Conf Rej Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfRejSent interface{}

    // Conf Rej Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfRejRcvd interface{}

    // Term Req Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TermReqSent interface{}

    // Term Req Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    TermReqRcvd interface{}

    // Term Ack Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TermAckSent interface{}

    // Term Ack Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    TermAckRcvd interface{}

    // Code Rej Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    CodeRejSent interface{}

    // Code Rej Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    CodeRejRcvd interface{}

    // Proto Rej Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ProtoRejSent interface{}

    // Proto Rej Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ProtoRejRcvd interface{}

    // Echo Req Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    EchoReqSent interface{}

    // Echo Req Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    EchoReqRcvd interface{}

    // Echo Rep Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    EchoRepSent interface{}

    // Echo Rep Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    EchoRepRcvd interface{}

    // Disc Req Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    DiscReqSent interface{}

    // Disc Req Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    DiscReqRcvd interface{}

    // Line Protocol Up count. The type is interface{} with range:
    // 0..18446744073709551615.
    LinkUp interface{}

    // Keepalive link failure count. The type is interface{} with range:
    // 0..18446744073709551615.
    LinkError interface{}
}

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetFilter() yfilter.YFilter { return lcpStatistics.YFilter }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) SetFilter(yf yfilter.YFilter) { lcpStatistics.YFilter = yf }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetGoName(yname string) string {
    if yname == "conf-req-sent" { return "ConfReqSent" }
    if yname == "conf-req-rcvd" { return "ConfReqRcvd" }
    if yname == "conf-ack-sent" { return "ConfAckSent" }
    if yname == "conf-ack-rcvd" { return "ConfAckRcvd" }
    if yname == "conf-nak-sent" { return "ConfNakSent" }
    if yname == "conf-nak-rcvd" { return "ConfNakRcvd" }
    if yname == "conf-rej-sent" { return "ConfRejSent" }
    if yname == "conf-rej-rcvd" { return "ConfRejRcvd" }
    if yname == "term-req-sent" { return "TermReqSent" }
    if yname == "term-req-rcvd" { return "TermReqRcvd" }
    if yname == "term-ack-sent" { return "TermAckSent" }
    if yname == "term-ack-rcvd" { return "TermAckRcvd" }
    if yname == "code-rej-sent" { return "CodeRejSent" }
    if yname == "code-rej-rcvd" { return "CodeRejRcvd" }
    if yname == "proto-rej-sent" { return "ProtoRejSent" }
    if yname == "proto-rej-rcvd" { return "ProtoRejRcvd" }
    if yname == "echo-req-sent" { return "EchoReqSent" }
    if yname == "echo-req-rcvd" { return "EchoReqRcvd" }
    if yname == "echo-rep-sent" { return "EchoRepSent" }
    if yname == "echo-rep-rcvd" { return "EchoRepRcvd" }
    if yname == "disc-req-sent" { return "DiscReqSent" }
    if yname == "disc-req-rcvd" { return "DiscReqRcvd" }
    if yname == "link-up" { return "LinkUp" }
    if yname == "link-error" { return "LinkError" }
    return ""
}

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetSegmentPath() string {
    return "lcp-statistics"
}

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["conf-req-sent"] = lcpStatistics.ConfReqSent
    leafs["conf-req-rcvd"] = lcpStatistics.ConfReqRcvd
    leafs["conf-ack-sent"] = lcpStatistics.ConfAckSent
    leafs["conf-ack-rcvd"] = lcpStatistics.ConfAckRcvd
    leafs["conf-nak-sent"] = lcpStatistics.ConfNakSent
    leafs["conf-nak-rcvd"] = lcpStatistics.ConfNakRcvd
    leafs["conf-rej-sent"] = lcpStatistics.ConfRejSent
    leafs["conf-rej-rcvd"] = lcpStatistics.ConfRejRcvd
    leafs["term-req-sent"] = lcpStatistics.TermReqSent
    leafs["term-req-rcvd"] = lcpStatistics.TermReqRcvd
    leafs["term-ack-sent"] = lcpStatistics.TermAckSent
    leafs["term-ack-rcvd"] = lcpStatistics.TermAckRcvd
    leafs["code-rej-sent"] = lcpStatistics.CodeRejSent
    leafs["code-rej-rcvd"] = lcpStatistics.CodeRejRcvd
    leafs["proto-rej-sent"] = lcpStatistics.ProtoRejSent
    leafs["proto-rej-rcvd"] = lcpStatistics.ProtoRejRcvd
    leafs["echo-req-sent"] = lcpStatistics.EchoReqSent
    leafs["echo-req-rcvd"] = lcpStatistics.EchoReqRcvd
    leafs["echo-rep-sent"] = lcpStatistics.EchoRepSent
    leafs["echo-rep-rcvd"] = lcpStatistics.EchoRepRcvd
    leafs["disc-req-sent"] = lcpStatistics.DiscReqSent
    leafs["disc-req-rcvd"] = lcpStatistics.DiscReqRcvd
    leafs["link-up"] = lcpStatistics.LinkUp
    leafs["link-error"] = lcpStatistics.LinkError
    return leafs
}

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetYangName() string { return "lcp-statistics" }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) SetParent(parent types.Entity) { lcpStatistics.parent = parent }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetParent() types.Entity { return lcpStatistics.parent }

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetParentYangName() string { return "statistics" }

// Ppp_Nodes_Node_Statistics_AuthenticationStatistics
// PPP Authentication statistics
type Ppp_Nodes_Node_Statistics_AuthenticationStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PAP Request packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PapReqSent interface{}

    // PAP Request packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PapReqRcvd interface{}

    // PAP Ack packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PapAckSent interface{}

    // PAP Ack packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PapAckRcvd interface{}

    // PAP Nak packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    PapNakSent interface{}

    // PAP Nak packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    PapNakRcvd interface{}

    // CHAP challenge packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapChallSent interface{}

    // CHAP challenge packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapChallRcvd interface{}

    // CHAP response packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapRespSent interface{}

    // CHAP response packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapRespRcvd interface{}

    // CHAP reply success packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapRepSuccSent interface{}

    // CHAP reply success packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapRepSuccRcvd interface{}

    // CHAP reply failure packets sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapRepFailSent interface{}

    // CHAP reply failure packets received. The type is interface{} with range:
    // 0..18446744073709551615.
    ChapRepFailRcvd interface{}

    // Authentication timeout count. The type is interface{} with range:
    // 0..18446744073709551615.
    AuthTimeoutCount interface{}
}

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetFilter() yfilter.YFilter { return authenticationStatistics.YFilter }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) SetFilter(yf yfilter.YFilter) { authenticationStatistics.YFilter = yf }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetGoName(yname string) string {
    if yname == "pap-req-sent" { return "PapReqSent" }
    if yname == "pap-req-rcvd" { return "PapReqRcvd" }
    if yname == "pap-ack-sent" { return "PapAckSent" }
    if yname == "pap-ack-rcvd" { return "PapAckRcvd" }
    if yname == "pap-nak-sent" { return "PapNakSent" }
    if yname == "pap-nak-rcvd" { return "PapNakRcvd" }
    if yname == "chap-chall-sent" { return "ChapChallSent" }
    if yname == "chap-chall-rcvd" { return "ChapChallRcvd" }
    if yname == "chap-resp-sent" { return "ChapRespSent" }
    if yname == "chap-resp-rcvd" { return "ChapRespRcvd" }
    if yname == "chap-rep-succ-sent" { return "ChapRepSuccSent" }
    if yname == "chap-rep-succ-rcvd" { return "ChapRepSuccRcvd" }
    if yname == "chap-rep-fail-sent" { return "ChapRepFailSent" }
    if yname == "chap-rep-fail-rcvd" { return "ChapRepFailRcvd" }
    if yname == "auth-timeout-count" { return "AuthTimeoutCount" }
    return ""
}

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetSegmentPath() string {
    return "authentication-statistics"
}

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pap-req-sent"] = authenticationStatistics.PapReqSent
    leafs["pap-req-rcvd"] = authenticationStatistics.PapReqRcvd
    leafs["pap-ack-sent"] = authenticationStatistics.PapAckSent
    leafs["pap-ack-rcvd"] = authenticationStatistics.PapAckRcvd
    leafs["pap-nak-sent"] = authenticationStatistics.PapNakSent
    leafs["pap-nak-rcvd"] = authenticationStatistics.PapNakRcvd
    leafs["chap-chall-sent"] = authenticationStatistics.ChapChallSent
    leafs["chap-chall-rcvd"] = authenticationStatistics.ChapChallRcvd
    leafs["chap-resp-sent"] = authenticationStatistics.ChapRespSent
    leafs["chap-resp-rcvd"] = authenticationStatistics.ChapRespRcvd
    leafs["chap-rep-succ-sent"] = authenticationStatistics.ChapRepSuccSent
    leafs["chap-rep-succ-rcvd"] = authenticationStatistics.ChapRepSuccRcvd
    leafs["chap-rep-fail-sent"] = authenticationStatistics.ChapRepFailSent
    leafs["chap-rep-fail-rcvd"] = authenticationStatistics.ChapRepFailRcvd
    leafs["auth-timeout-count"] = authenticationStatistics.AuthTimeoutCount
    return leafs
}

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetYangName() string { return "authentication-statistics" }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) SetParent(parent types.Entity) { authenticationStatistics.parent = parent }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetParent() types.Entity { return authenticationStatistics.parent }

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetParentYangName() string { return "statistics" }

// Ppp_Nodes_Node_Statistics_NcpStatisticsArray
// Array of PPP NCP Statistics
type Ppp_Nodes_Node_Statistics_NcpStatisticsArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NCP identifier. The type is NcpIdent.
    NcpIdentifier interface{}

    // Conf Req Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfReqSent interface{}

    // Conf Req Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfReqRcvd interface{}

    // Conf Ack Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfAckSent interface{}

    // Conf Ack Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfAckRcvd interface{}

    // Conf Nak Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfNakSent interface{}

    // Conf Nak Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfNakRcvd interface{}

    // Conf Rej Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfRejSent interface{}

    // Conf Rej Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ConfRejRcvd interface{}

    // Term Req Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TermReqSent interface{}

    // Term Req Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    TermReqRcvd interface{}

    // Term Ack Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    TermAckSent interface{}

    // Term Ack Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    TermAckRcvd interface{}

    // Proto Rej Packets Sent. The type is interface{} with range:
    // 0..18446744073709551615.
    ProtoRejSent interface{}

    // Proto Rej Packets Received. The type is interface{} with range:
    // 0..18446744073709551615.
    ProtoRejRcvd interface{}
}

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetFilter() yfilter.YFilter { return ncpStatisticsArray.YFilter }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) SetFilter(yf yfilter.YFilter) { ncpStatisticsArray.YFilter = yf }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetGoName(yname string) string {
    if yname == "ncp-identifier" { return "NcpIdentifier" }
    if yname == "conf-req-sent" { return "ConfReqSent" }
    if yname == "conf-req-rcvd" { return "ConfReqRcvd" }
    if yname == "conf-ack-sent" { return "ConfAckSent" }
    if yname == "conf-ack-rcvd" { return "ConfAckRcvd" }
    if yname == "conf-nak-sent" { return "ConfNakSent" }
    if yname == "conf-nak-rcvd" { return "ConfNakRcvd" }
    if yname == "conf-rej-sent" { return "ConfRejSent" }
    if yname == "conf-rej-rcvd" { return "ConfRejRcvd" }
    if yname == "term-req-sent" { return "TermReqSent" }
    if yname == "term-req-rcvd" { return "TermReqRcvd" }
    if yname == "term-ack-sent" { return "TermAckSent" }
    if yname == "term-ack-rcvd" { return "TermAckRcvd" }
    if yname == "proto-rej-sent" { return "ProtoRejSent" }
    if yname == "proto-rej-rcvd" { return "ProtoRejRcvd" }
    return ""
}

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetSegmentPath() string {
    return "ncp-statistics-array"
}

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ncp-identifier"] = ncpStatisticsArray.NcpIdentifier
    leafs["conf-req-sent"] = ncpStatisticsArray.ConfReqSent
    leafs["conf-req-rcvd"] = ncpStatisticsArray.ConfReqRcvd
    leafs["conf-ack-sent"] = ncpStatisticsArray.ConfAckSent
    leafs["conf-ack-rcvd"] = ncpStatisticsArray.ConfAckRcvd
    leafs["conf-nak-sent"] = ncpStatisticsArray.ConfNakSent
    leafs["conf-nak-rcvd"] = ncpStatisticsArray.ConfNakRcvd
    leafs["conf-rej-sent"] = ncpStatisticsArray.ConfRejSent
    leafs["conf-rej-rcvd"] = ncpStatisticsArray.ConfRejRcvd
    leafs["term-req-sent"] = ncpStatisticsArray.TermReqSent
    leafs["term-req-rcvd"] = ncpStatisticsArray.TermReqRcvd
    leafs["term-ack-sent"] = ncpStatisticsArray.TermAckSent
    leafs["term-ack-rcvd"] = ncpStatisticsArray.TermAckRcvd
    leafs["proto-rej-sent"] = ncpStatisticsArray.ProtoRejSent
    leafs["proto-rej-rcvd"] = ncpStatisticsArray.ProtoRejRcvd
    return leafs
}

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetBundleName() string { return "cisco_ios_xr" }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetYangName() string { return "ncp-statistics-array" }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) SetParent(parent types.Entity) { ncpStatisticsArray.parent = parent }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetParent() types.Entity { return ncpStatisticsArray.parent }

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetParentYangName() string { return "statistics" }

// Ppp_Nodes_Node_NodeInterfaces
// Per interface PPP operational data
type Ppp_Nodes_Node_NodeInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LCP and summarized NCP data for an interface running PPP. The type is slice
    // of Ppp_Nodes_Node_NodeInterfaces_NodeInterface.
    NodeInterface []Ppp_Nodes_Node_NodeInterfaces_NodeInterface
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetFilter() yfilter.YFilter { return nodeInterfaces.YFilter }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) SetFilter(yf yfilter.YFilter) { nodeInterfaces.YFilter = yf }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetGoName(yname string) string {
    if yname == "node-interface" { return "NodeInterface" }
    return ""
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetSegmentPath() string {
    return "node-interfaces"
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-interface" {
        for _, c := range nodeInterfaces.NodeInterface {
            if nodeInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_NodeInterfaces_NodeInterface{}
        nodeInterfaces.NodeInterface = append(nodeInterfaces.NodeInterface, child)
        return &nodeInterfaces.NodeInterface[len(nodeInterfaces.NodeInterface)-1]
    }
    return nil
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaces.NodeInterface {
        children[nodeInterfaces.NodeInterface[i].GetSegmentPath()] = &nodeInterfaces.NodeInterface[i]
    }
    return children
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetYangName() string { return "node-interfaces" }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) SetParent(parent types.Entity) { nodeInterfaces.parent = parent }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetParent() types.Entity { return nodeInterfaces.parent }

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface
// LCP and summarized NCP data for an interface
// running PPP
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface running PPP. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Parent state. The type is interface{} with range: 0..4294967295.
    ParentState interface{}

    // Line state. The type is interface{} with range: 0..4294967295.
    LineState interface{}

    // Loopback detected. The type is bool.
    IsLoopbackDetected interface{}

    // Caps IDB SRG role. The type is interface{} with range: 0..4294967295.
    CapsIdbSrgRole interface{}

    // Session SRG role. The type is interface{} with range: 0..4294967295.
    SessionSrgRole interface{}

    // Keepalive value. The type is interface{} with range: 0..4294967295.
    KeepalivePeriod interface{}

    // Keepalive retry count. The type is interface{} with range: 0..4294967295.
    KeepaliveRetryCount interface{}

    // Is SSRP configured. The type is bool.
    IsSsrpConfigured interface{}

    // Is L2 AC. The type is bool.
    IsL2Ac interface{}

    // Provisioned. The type is bool.
    Provisioned interface{}

    // IP Interworking Enabled. The type is bool.
    IpInterworkingEnabled interface{}

    // XConnect ID. The type is interface{} with range: 0..4294967295.
    XconnectId interface{}

    // Is tunneled session. The type is bool.
    IsTunneledSession interface{}

    // SSRP Peer ID. The type is string.
    SsrpPeerId interface{}

    // PPP/LCP state value. The type is PppFsmState.
    LcpState interface{}

    // LCP SSO state. The type is PppSsoFsmState.
    LcpssoState interface{}

    // Is LCP Delayed. The type is bool.
    IsLcpDelayed interface{}

    // Local MRU. The type is interface{} with range: 0..65535.
    LocalMru interface{}

    // Peer MRU. The type is interface{} with range: 0..65535.
    PeerMru interface{}

    // Local MRRU. The type is interface{} with range: 0..65535.
    LocalMrru interface{}

    // Peer MRRU. The type is interface{} with range: 0..65535.
    PeerMrru interface{}

    // Local Endpt Discriminator. The type is string with length: 0..41.
    LocalEd interface{}

    // Peer Endpt Discriminator. The type is string with length: 0..41.
    PeerEd interface{}

    // Is MCMP enabled. The type is bool.
    IsMcmpEnabled interface{}

    // Local MCMP classes. The type is interface{} with range: 0..255.
    LocalMcmpClasses interface{}

    // Peer MCMP classes. The type is interface{} with range: 0..255.
    PeerMcmpClasses interface{}

    // Session expiry time in seconds since 00:00:00 on January 1, 1970, UTC. The
    // type is interface{} with range: 0..4294967295. Units are second.
    SessionExpires interface{}

    // MP information.
    MpInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo

    // Configured timeout.
    ConfiguredTimeout Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout

    // Authentication information.
    AuthInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo

    // Array of per-NCP data. The type is slice of
    // Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray.
    NcpInfoArray []Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetFilter() yfilter.YFilter { return nodeInterface.YFilter }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) SetFilter(yf yfilter.YFilter) { nodeInterface.YFilter = yf }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "parent-state" { return "ParentState" }
    if yname == "line-state" { return "LineState" }
    if yname == "is-loopback-detected" { return "IsLoopbackDetected" }
    if yname == "caps-idb-srg-role" { return "CapsIdbSrgRole" }
    if yname == "session-srg-role" { return "SessionSrgRole" }
    if yname == "keepalive-period" { return "KeepalivePeriod" }
    if yname == "keepalive-retry-count" { return "KeepaliveRetryCount" }
    if yname == "is-ssrp-configured" { return "IsSsrpConfigured" }
    if yname == "is-l2ac" { return "IsL2Ac" }
    if yname == "provisioned" { return "Provisioned" }
    if yname == "ip-interworking-enabled" { return "IpInterworkingEnabled" }
    if yname == "xconnect-id" { return "XconnectId" }
    if yname == "is-tunneled-session" { return "IsTunneledSession" }
    if yname == "ssrp-peer-id" { return "SsrpPeerId" }
    if yname == "lcp-state" { return "LcpState" }
    if yname == "lcpsso-state" { return "LcpssoState" }
    if yname == "is-lcp-delayed" { return "IsLcpDelayed" }
    if yname == "local-mru" { return "LocalMru" }
    if yname == "peer-mru" { return "PeerMru" }
    if yname == "local-mrru" { return "LocalMrru" }
    if yname == "peer-mrru" { return "PeerMrru" }
    if yname == "local-ed" { return "LocalEd" }
    if yname == "peer-ed" { return "PeerEd" }
    if yname == "is-mcmp-enabled" { return "IsMcmpEnabled" }
    if yname == "local-mcmp-classes" { return "LocalMcmpClasses" }
    if yname == "peer-mcmp-classes" { return "PeerMcmpClasses" }
    if yname == "session-expires" { return "SessionExpires" }
    if yname == "mp-info" { return "MpInfo" }
    if yname == "configured-timeout" { return "ConfiguredTimeout" }
    if yname == "auth-info" { return "AuthInfo" }
    if yname == "ncp-info-array" { return "NcpInfoArray" }
    return ""
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetSegmentPath() string {
    return "node-interface" + "[interface='" + fmt.Sprintf("%v", nodeInterface.Interface) + "']"
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mp-info" {
        return &nodeInterface.MpInfo
    }
    if childYangName == "configured-timeout" {
        return &nodeInterface.ConfiguredTimeout
    }
    if childYangName == "auth-info" {
        return &nodeInterface.AuthInfo
    }
    if childYangName == "ncp-info-array" {
        for _, c := range nodeInterface.NcpInfoArray {
            if nodeInterface.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray{}
        nodeInterface.NcpInfoArray = append(nodeInterface.NcpInfoArray, child)
        return &nodeInterface.NcpInfoArray[len(nodeInterface.NcpInfoArray)-1]
    }
    return nil
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mp-info"] = &nodeInterface.MpInfo
    children["configured-timeout"] = &nodeInterface.ConfiguredTimeout
    children["auth-info"] = &nodeInterface.AuthInfo
    for i := range nodeInterface.NcpInfoArray {
        children[nodeInterface.NcpInfoArray[i].GetSegmentPath()] = &nodeInterface.NcpInfoArray[i]
    }
    return children
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = nodeInterface.Interface
    leafs["parent-state"] = nodeInterface.ParentState
    leafs["line-state"] = nodeInterface.LineState
    leafs["is-loopback-detected"] = nodeInterface.IsLoopbackDetected
    leafs["caps-idb-srg-role"] = nodeInterface.CapsIdbSrgRole
    leafs["session-srg-role"] = nodeInterface.SessionSrgRole
    leafs["keepalive-period"] = nodeInterface.KeepalivePeriod
    leafs["keepalive-retry-count"] = nodeInterface.KeepaliveRetryCount
    leafs["is-ssrp-configured"] = nodeInterface.IsSsrpConfigured
    leafs["is-l2ac"] = nodeInterface.IsL2Ac
    leafs["provisioned"] = nodeInterface.Provisioned
    leafs["ip-interworking-enabled"] = nodeInterface.IpInterworkingEnabled
    leafs["xconnect-id"] = nodeInterface.XconnectId
    leafs["is-tunneled-session"] = nodeInterface.IsTunneledSession
    leafs["ssrp-peer-id"] = nodeInterface.SsrpPeerId
    leafs["lcp-state"] = nodeInterface.LcpState
    leafs["lcpsso-state"] = nodeInterface.LcpssoState
    leafs["is-lcp-delayed"] = nodeInterface.IsLcpDelayed
    leafs["local-mru"] = nodeInterface.LocalMru
    leafs["peer-mru"] = nodeInterface.PeerMru
    leafs["local-mrru"] = nodeInterface.LocalMrru
    leafs["peer-mrru"] = nodeInterface.PeerMrru
    leafs["local-ed"] = nodeInterface.LocalEd
    leafs["peer-ed"] = nodeInterface.PeerEd
    leafs["is-mcmp-enabled"] = nodeInterface.IsMcmpEnabled
    leafs["local-mcmp-classes"] = nodeInterface.LocalMcmpClasses
    leafs["peer-mcmp-classes"] = nodeInterface.PeerMcmpClasses
    leafs["session-expires"] = nodeInterface.SessionExpires
    return leafs
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetYangName() string { return "node-interface" }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) SetParent(parent types.Entity) { nodeInterface.parent = parent }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetParent() types.Entity { return nodeInterface.parent }

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetParentYangName() string { return "node-interfaces" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo
// MP information
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is an MP bundle. The type is bool.
    IsMpBundle interface{}

    // MP Bundle Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    MpBundleInterface interface{}

    // MP Bundle Member. The type is bool.
    IsMpBundleMember interface{}

    // MP Group. The type is interface{} with range: 0..4294967295.
    MpGroup interface{}

    // Number of active links. The type is interface{} with range: 0..65535.
    ActiveLinks interface{}

    // Number of inactive links. The type is interface{} with range: 0..65535.
    InactiveLinks interface{}

    // Minimum active links required for the MPbundle to come up. The type is
    // interface{} with range: 0..65535.
    MinimumActiveLinks interface{}

    // Member State. The type is PppLcpMpMbrState.
    MpState interface{}

    // Array of MP members. The type is slice of
    // Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray.
    MpMemberInfoArray []Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetFilter() yfilter.YFilter { return mpInfo.YFilter }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) SetFilter(yf yfilter.YFilter) { mpInfo.YFilter = yf }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetGoName(yname string) string {
    if yname == "is-mp-bundle" { return "IsMpBundle" }
    if yname == "mp-bundle-interface" { return "MpBundleInterface" }
    if yname == "is-mp-bundle-member" { return "IsMpBundleMember" }
    if yname == "mp-group" { return "MpGroup" }
    if yname == "active-links" { return "ActiveLinks" }
    if yname == "inactive-links" { return "InactiveLinks" }
    if yname == "minimum-active-links" { return "MinimumActiveLinks" }
    if yname == "mp-state" { return "MpState" }
    if yname == "mp-member-info-array" { return "MpMemberInfoArray" }
    return ""
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetSegmentPath() string {
    return "mp-info"
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mp-member-info-array" {
        for _, c := range mpInfo.MpMemberInfoArray {
            if mpInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray{}
        mpInfo.MpMemberInfoArray = append(mpInfo.MpMemberInfoArray, child)
        return &mpInfo.MpMemberInfoArray[len(mpInfo.MpMemberInfoArray)-1]
    }
    return nil
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range mpInfo.MpMemberInfoArray {
        children[mpInfo.MpMemberInfoArray[i].GetSegmentPath()] = &mpInfo.MpMemberInfoArray[i]
    }
    return children
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-mp-bundle"] = mpInfo.IsMpBundle
    leafs["mp-bundle-interface"] = mpInfo.MpBundleInterface
    leafs["is-mp-bundle-member"] = mpInfo.IsMpBundleMember
    leafs["mp-group"] = mpInfo.MpGroup
    leafs["active-links"] = mpInfo.ActiveLinks
    leafs["inactive-links"] = mpInfo.InactiveLinks
    leafs["minimum-active-links"] = mpInfo.MinimumActiveLinks
    leafs["mp-state"] = mpInfo.MpState
    return leafs
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetYangName() string { return "mp-info" }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) SetParent(parent types.Entity) { mpInfo.parent = parent }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetParent() types.Entity { return mpInfo.parent }

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetParentYangName() string { return "node-interface" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray
// Array of MP members
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Member Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // Member State. The type is PppLcpMpMbrState.
    State interface{}
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetFilter() yfilter.YFilter { return mpMemberInfoArray.YFilter }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) SetFilter(yf yfilter.YFilter) { mpMemberInfoArray.YFilter = yf }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "state" { return "State" }
    return ""
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetSegmentPath() string {
    return "mp-member-info-array"
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = mpMemberInfoArray.Interface
    leafs["state"] = mpMemberInfoArray.State
    return leafs
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetBundleName() string { return "cisco_ios_xr" }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetYangName() string { return "mp-member-info-array" }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) SetParent(parent types.Entity) { mpMemberInfoArray.parent = parent }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetParent() types.Entity { return mpMemberInfoArray.parent }

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetParentYangName() string { return "mp-info" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout
// Configured timeout
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minutes. The type is interface{} with range: 0..4294967295. Units are
    // minute.
    Minutes interface{}

    // Seconds. The type is interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetFilter() yfilter.YFilter { return configuredTimeout.YFilter }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) SetFilter(yf yfilter.YFilter) { configuredTimeout.YFilter = yf }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetGoName(yname string) string {
    if yname == "minutes" { return "Minutes" }
    if yname == "seconds" { return "Seconds" }
    return ""
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetSegmentPath() string {
    return "configured-timeout"
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minutes"] = configuredTimeout.Minutes
    leafs["seconds"] = configuredTimeout.Seconds
    return leafs
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetBundleName() string { return "cisco_ios_xr" }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetYangName() string { return "configured-timeout" }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) SetParent(parent types.Entity) { configuredTimeout.parent = parent }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetParent() types.Entity { return configuredTimeout.parent }

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetParentYangName() string { return "node-interface" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo
// Authentication information
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is authenticated. The type is bool.
    IsAuthenticated interface{}

    // Is SSO authenticated. The type is bool.
    IsSsoAuthenticated interface{}

    // Of Us authentication type. The type is interface{} with range: 0..255.
    OfUsAuth interface{}

    // Of Peer authentication type. The type is interface{} with range: 0..255.
    OfPeerAuth interface{}

    // Local authenticated name. The type is string.
    OfUsName interface{}

    // Peer's authenticated name. The type is string.
    OfPeerName interface{}

    // Of Us auth SSO FSM State. The type is PppSsoFsmState.
    OfUsSsoState interface{}

    // Of Peer auth SSO FSM State. The type is PppSsoFsmState.
    OfPeerSsoState interface{}
}

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetFilter() yfilter.YFilter { return authInfo.YFilter }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) SetFilter(yf yfilter.YFilter) { authInfo.YFilter = yf }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetGoName(yname string) string {
    if yname == "is-authenticated" { return "IsAuthenticated" }
    if yname == "is-sso-authenticated" { return "IsSsoAuthenticated" }
    if yname == "of-us-auth" { return "OfUsAuth" }
    if yname == "of-peer-auth" { return "OfPeerAuth" }
    if yname == "of-us-name" { return "OfUsName" }
    if yname == "of-peer-name" { return "OfPeerName" }
    if yname == "of-us-sso-state" { return "OfUsSsoState" }
    if yname == "of-peer-sso-state" { return "OfPeerSsoState" }
    return ""
}

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetSegmentPath() string {
    return "auth-info"
}

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-authenticated"] = authInfo.IsAuthenticated
    leafs["is-sso-authenticated"] = authInfo.IsSsoAuthenticated
    leafs["of-us-auth"] = authInfo.OfUsAuth
    leafs["of-peer-auth"] = authInfo.OfPeerAuth
    leafs["of-us-name"] = authInfo.OfUsName
    leafs["of-peer-name"] = authInfo.OfPeerName
    leafs["of-us-sso-state"] = authInfo.OfUsSsoState
    leafs["of-peer-sso-state"] = authInfo.OfPeerSsoState
    return leafs
}

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetBundleName() string { return "cisco_ios_xr" }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetYangName() string { return "auth-info" }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) SetParent(parent types.Entity) { authInfo.parent = parent }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetParent() types.Entity { return authInfo.parent }

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetParentYangName() string { return "node-interface" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray
// Array of per-NCP data
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NCP state value. The type is PppFsmState.
    NcpState interface{}

    // NCP SSO State. The type is PppSsoFsmState.
    NcpssoState interface{}

    // Is Passive. The type is bool.
    IsPassive interface{}

    // NCP state identifier. The type is NcpIdent.
    NcpIdentifier interface{}

    // Specific NCP info.
    NcpInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo
}

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetFilter() yfilter.YFilter { return ncpInfoArray.YFilter }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) SetFilter(yf yfilter.YFilter) { ncpInfoArray.YFilter = yf }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetGoName(yname string) string {
    if yname == "ncp-state" { return "NcpState" }
    if yname == "ncpsso-state" { return "NcpssoState" }
    if yname == "is-passive" { return "IsPassive" }
    if yname == "ncp-identifier" { return "NcpIdentifier" }
    if yname == "ncp-info" { return "NcpInfo" }
    return ""
}

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetSegmentPath() string {
    return "ncp-info-array"
}

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ncp-info" {
        return &ncpInfoArray.NcpInfo
    }
    return nil
}

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ncp-info"] = &ncpInfoArray.NcpInfo
    return children
}

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ncp-state"] = ncpInfoArray.NcpState
    leafs["ncpsso-state"] = ncpInfoArray.NcpssoState
    leafs["is-passive"] = ncpInfoArray.IsPassive
    leafs["ncp-identifier"] = ncpInfoArray.NcpIdentifier
    return leafs
}

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetBundleName() string { return "cisco_ios_xr" }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetYangName() string { return "ncp-info-array" }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) SetParent(parent types.Entity) { ncpInfoArray.parent = parent }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetParent() types.Entity { return ncpInfoArray.parent }

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetParentYangName() string { return "node-interface" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo
// Specific NCP info
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is NcpIdent.
    Type interface{}

    // Info for IPCP.
    IpcpInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo

    // Info for IPCPIW.
    IpcpiwInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo

    // Info for IPv6CP.
    Ipv6CpInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetFilter() yfilter.YFilter { return ncpInfo.YFilter }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) SetFilter(yf yfilter.YFilter) { ncpInfo.YFilter = yf }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "ipcp-info" { return "IpcpInfo" }
    if yname == "ipcpiw-info" { return "IpcpiwInfo" }
    if yname == "ipv6cp-info" { return "Ipv6CpInfo" }
    return ""
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetSegmentPath() string {
    return "ncp-info"
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipcp-info" {
        return &ncpInfo.IpcpInfo
    }
    if childYangName == "ipcpiw-info" {
        return &ncpInfo.IpcpiwInfo
    }
    if childYangName == "ipv6cp-info" {
        return &ncpInfo.Ipv6CpInfo
    }
    return nil
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipcp-info"] = &ncpInfo.IpcpInfo
    children["ipcpiw-info"] = &ncpInfo.IpcpiwInfo
    children["ipv6cp-info"] = &ncpInfo.Ipv6CpInfo
    return children
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = ncpInfo.Type
    return leafs
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetYangName() string { return "ncp-info" }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) SetParent(parent types.Entity) { ncpInfo.parent = parent }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetParent() types.Entity { return ncpInfo.parent }

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetParentYangName() string { return "ncp-info-array" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo
// Info for IPCP
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Peer IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Peer IPv4 netmask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerNetmask interface{}

    // Peer DNS Primary. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DnsPrimary interface{}

    // Peer DNS Secondary. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DnsSecondary interface{}

    // Peer WINS Primary. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    WinsPrimary interface{}

    // Peer WINS Secondary. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    WinsSecondary interface{}

    // Is IPHC Configured. The type is bool.
    IsIphcConfigured interface{}

    // Local IPHC options.
    LocalIphcOptions Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions

    // Peer IPHC options.
    PeerIphcOptions Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions
}

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetFilter() yfilter.YFilter { return ipcpInfo.YFilter }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) SetFilter(yf yfilter.YFilter) { ipcpInfo.YFilter = yf }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetGoName(yname string) string {
    if yname == "local-address" { return "LocalAddress" }
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-netmask" { return "PeerNetmask" }
    if yname == "dns-primary" { return "DnsPrimary" }
    if yname == "dns-secondary" { return "DnsSecondary" }
    if yname == "wins-primary" { return "WinsPrimary" }
    if yname == "wins-secondary" { return "WinsSecondary" }
    if yname == "is-iphc-configured" { return "IsIphcConfigured" }
    if yname == "local-iphc-options" { return "LocalIphcOptions" }
    if yname == "peer-iphc-options" { return "PeerIphcOptions" }
    return ""
}

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetSegmentPath() string {
    return "ipcp-info"
}

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-iphc-options" {
        return &ipcpInfo.LocalIphcOptions
    }
    if childYangName == "peer-iphc-options" {
        return &ipcpInfo.PeerIphcOptions
    }
    return nil
}

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-iphc-options"] = &ipcpInfo.LocalIphcOptions
    children["peer-iphc-options"] = &ipcpInfo.PeerIphcOptions
    return children
}

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-address"] = ipcpInfo.LocalAddress
    leafs["peer-address"] = ipcpInfo.PeerAddress
    leafs["peer-netmask"] = ipcpInfo.PeerNetmask
    leafs["dns-primary"] = ipcpInfo.DnsPrimary
    leafs["dns-secondary"] = ipcpInfo.DnsSecondary
    leafs["wins-primary"] = ipcpInfo.WinsPrimary
    leafs["wins-secondary"] = ipcpInfo.WinsSecondary
    leafs["is-iphc-configured"] = ipcpInfo.IsIphcConfigured
    return leafs
}

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetYangName() string { return "ipcp-info" }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) SetParent(parent types.Entity) { ipcpInfo.parent = parent }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetParent() types.Entity { return ipcpInfo.parent }

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetParentYangName() string { return "ncp-info" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions
// Local IPHC options
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Compression type. The type is PppIphcCompression.
    CompressionType interface{}

    // TCP space. The type is interface{} with range: 0..65535.
    TcpSpace interface{}

    // Non-TCP space. The type is interface{} with range: 0..65535.
    NonTcpSpace interface{}

    // Max period. The type is interface{} with range: 0..65535.
    MaxPeriod interface{}

    // Max time. The type is interface{} with range: 0..65535.
    MaxTime interface{}

    // Max header. The type is interface{} with range: 0..65535.
    MaxHeader interface{}

    // RTP compression. The type is bool.
    RtpCompression interface{}

    // EcRTP compression. The type is bool.
    EcRtpCompression interface{}
}

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetFilter() yfilter.YFilter { return localIphcOptions.YFilter }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) SetFilter(yf yfilter.YFilter) { localIphcOptions.YFilter = yf }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetGoName(yname string) string {
    if yname == "compression-type" { return "CompressionType" }
    if yname == "tcp-space" { return "TcpSpace" }
    if yname == "non-tcp-space" { return "NonTcpSpace" }
    if yname == "max-period" { return "MaxPeriod" }
    if yname == "max-time" { return "MaxTime" }
    if yname == "max-header" { return "MaxHeader" }
    if yname == "rtp-compression" { return "RtpCompression" }
    if yname == "ec-rtp-compression" { return "EcRtpCompression" }
    return ""
}

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetSegmentPath() string {
    return "local-iphc-options"
}

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["compression-type"] = localIphcOptions.CompressionType
    leafs["tcp-space"] = localIphcOptions.TcpSpace
    leafs["non-tcp-space"] = localIphcOptions.NonTcpSpace
    leafs["max-period"] = localIphcOptions.MaxPeriod
    leafs["max-time"] = localIphcOptions.MaxTime
    leafs["max-header"] = localIphcOptions.MaxHeader
    leafs["rtp-compression"] = localIphcOptions.RtpCompression
    leafs["ec-rtp-compression"] = localIphcOptions.EcRtpCompression
    return leafs
}

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetBundleName() string { return "cisco_ios_xr" }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetYangName() string { return "local-iphc-options" }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) SetParent(parent types.Entity) { localIphcOptions.parent = parent }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetParent() types.Entity { return localIphcOptions.parent }

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetParentYangName() string { return "ipcp-info" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions
// Peer IPHC options
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Compression type. The type is PppIphcCompression.
    CompressionType interface{}

    // TCP space. The type is interface{} with range: 0..65535.
    TcpSpace interface{}

    // Non-TCP space. The type is interface{} with range: 0..65535.
    NonTcpSpace interface{}

    // Max period. The type is interface{} with range: 0..65535.
    MaxPeriod interface{}

    // Max time. The type is interface{} with range: 0..65535.
    MaxTime interface{}

    // Max header. The type is interface{} with range: 0..65535.
    MaxHeader interface{}

    // RTP compression. The type is bool.
    RtpCompression interface{}

    // EcRTP compression. The type is bool.
    EcRtpCompression interface{}
}

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetFilter() yfilter.YFilter { return peerIphcOptions.YFilter }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) SetFilter(yf yfilter.YFilter) { peerIphcOptions.YFilter = yf }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetGoName(yname string) string {
    if yname == "compression-type" { return "CompressionType" }
    if yname == "tcp-space" { return "TcpSpace" }
    if yname == "non-tcp-space" { return "NonTcpSpace" }
    if yname == "max-period" { return "MaxPeriod" }
    if yname == "max-time" { return "MaxTime" }
    if yname == "max-header" { return "MaxHeader" }
    if yname == "rtp-compression" { return "RtpCompression" }
    if yname == "ec-rtp-compression" { return "EcRtpCompression" }
    return ""
}

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetSegmentPath() string {
    return "peer-iphc-options"
}

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["compression-type"] = peerIphcOptions.CompressionType
    leafs["tcp-space"] = peerIphcOptions.TcpSpace
    leafs["non-tcp-space"] = peerIphcOptions.NonTcpSpace
    leafs["max-period"] = peerIphcOptions.MaxPeriod
    leafs["max-time"] = peerIphcOptions.MaxTime
    leafs["max-header"] = peerIphcOptions.MaxHeader
    leafs["rtp-compression"] = peerIphcOptions.RtpCompression
    leafs["ec-rtp-compression"] = peerIphcOptions.EcRtpCompression
    return leafs
}

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetBundleName() string { return "cisco_ios_xr" }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetYangName() string { return "peer-iphc-options" }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) SetParent(parent types.Entity) { peerIphcOptions.parent = parent }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetParent() types.Entity { return peerIphcOptions.parent }

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetParentYangName() string { return "ipcp-info" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo
// Info for IPCPIW
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Peer IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetFilter() yfilter.YFilter { return ipcpiwInfo.YFilter }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) SetFilter(yf yfilter.YFilter) { ipcpiwInfo.YFilter = yf }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetGoName(yname string) string {
    if yname == "local-address" { return "LocalAddress" }
    if yname == "peer-address" { return "PeerAddress" }
    return ""
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetSegmentPath() string {
    return "ipcpiw-info"
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-address"] = ipcpiwInfo.LocalAddress
    leafs["peer-address"] = ipcpiwInfo.PeerAddress
    return leafs
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetYangName() string { return "ipcpiw-info" }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) SetParent(parent types.Entity) { ipcpiwInfo.parent = parent }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetParent() types.Entity { return ipcpiwInfo.parent }

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetParentYangName() string { return "ncp-info" }

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo
// Info for IPv6CP
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Peer IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}
}

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetFilter() yfilter.YFilter { return ipv6CpInfo.YFilter }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) SetFilter(yf yfilter.YFilter) { ipv6CpInfo.YFilter = yf }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetGoName(yname string) string {
    if yname == "local-address" { return "LocalAddress" }
    if yname == "peer-address" { return "PeerAddress" }
    return ""
}

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetSegmentPath() string {
    return "ipv6cp-info"
}

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-address"] = ipv6CpInfo.LocalAddress
    leafs["peer-address"] = ipv6CpInfo.PeerAddress
    return leafs
}

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetYangName() string { return "ipv6cp-info" }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) SetParent(parent types.Entity) { ipv6CpInfo.parent = parent }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetParent() types.Entity { return ipv6CpInfo.parent }

func (ipv6CpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6CpInfo) GetParentYangName() string { return "ncp-info" }

// Ppp_Nodes_Node_SsoAlerts
// PPP SSO Alert data for a particular node
type Ppp_Nodes_Node_SsoAlerts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPP SSO Alert data for a particular interface. The type is slice of
    // Ppp_Nodes_Node_SsoAlerts_SsoAlert.
    SsoAlert []Ppp_Nodes_Node_SsoAlerts_SsoAlert
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetFilter() yfilter.YFilter { return ssoAlerts.YFilter }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) SetFilter(yf yfilter.YFilter) { ssoAlerts.YFilter = yf }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetGoName(yname string) string {
    if yname == "sso-alert" { return "SsoAlert" }
    return ""
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetSegmentPath() string {
    return "sso-alerts"
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sso-alert" {
        for _, c := range ssoAlerts.SsoAlert {
            if ssoAlerts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_SsoAlerts_SsoAlert{}
        ssoAlerts.SsoAlert = append(ssoAlerts.SsoAlert, child)
        return &ssoAlerts.SsoAlert[len(ssoAlerts.SsoAlert)-1]
    }
    return nil
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssoAlerts.SsoAlert {
        children[ssoAlerts.SsoAlert[i].GetSegmentPath()] = &ssoAlerts.SsoAlert[i]
    }
    return children
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetBundleName() string { return "cisco_ios_xr" }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetYangName() string { return "sso-alerts" }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) SetParent(parent types.Entity) { ssoAlerts.parent = parent }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetParent() types.Entity { return ssoAlerts.parent }

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_SsoAlerts_SsoAlert
// PPP SSO Alert data for a particular interface
type Ppp_Nodes_Node_SsoAlerts_SsoAlert struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface with SSO Alert. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // LCP SSO Error.
    LcpError Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError

    // Of-us Authentication SSO Error.
    OfUsAuthError Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError

    // Of-peer Authentication SSO Error.
    OfPeerAuthError Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError

    // IPCP SSO Error.
    IpcpError Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError
}

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetFilter() yfilter.YFilter { return ssoAlert.YFilter }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) SetFilter(yf yfilter.YFilter) { ssoAlert.YFilter = yf }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "lcp-error" { return "LcpError" }
    if yname == "of-us-auth-error" { return "OfUsAuthError" }
    if yname == "of-peer-auth-error" { return "OfPeerAuthError" }
    if yname == "ipcp-error" { return "IpcpError" }
    return ""
}

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetSegmentPath() string {
    return "sso-alert" + "[interface='" + fmt.Sprintf("%v", ssoAlert.Interface) + "']"
}

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lcp-error" {
        return &ssoAlert.LcpError
    }
    if childYangName == "of-us-auth-error" {
        return &ssoAlert.OfUsAuthError
    }
    if childYangName == "of-peer-auth-error" {
        return &ssoAlert.OfPeerAuthError
    }
    if childYangName == "ipcp-error" {
        return &ssoAlert.IpcpError
    }
    return nil
}

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lcp-error"] = &ssoAlert.LcpError
    children["of-us-auth-error"] = &ssoAlert.OfUsAuthError
    children["of-peer-auth-error"] = &ssoAlert.OfPeerAuthError
    children["ipcp-error"] = &ssoAlert.IpcpError
    return children
}

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = ssoAlert.Interface
    return leafs
}

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetBundleName() string { return "cisco_ios_xr" }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetYangName() string { return "sso-alert" }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) SetParent(parent types.Entity) { ssoAlert.parent = parent }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetParent() types.Entity { return ssoAlert.parent }

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetParentYangName() string { return "sso-alerts" }

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError
// LCP SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetFilter() yfilter.YFilter { return lcpError.YFilter }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) SetFilter(yf yfilter.YFilter) { lcpError.YFilter = yf }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetGoName(yname string) string {
    if yname == "is-error" { return "IsError" }
    if yname == "error" { return "Error" }
    if yname == "context" { return "Context" }
    return ""
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetSegmentPath() string {
    return "lcp-error"
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-error"] = lcpError.IsError
    leafs["error"] = lcpError.Error
    leafs["context"] = lcpError.Context
    return leafs
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetBundleName() string { return "cisco_ios_xr" }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetYangName() string { return "lcp-error" }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) SetParent(parent types.Entity) { lcpError.parent = parent }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetParent() types.Entity { return lcpError.parent }

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetParentYangName() string { return "sso-alert" }

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError
// Of-us Authentication SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetFilter() yfilter.YFilter { return ofUsAuthError.YFilter }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) SetFilter(yf yfilter.YFilter) { ofUsAuthError.YFilter = yf }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetGoName(yname string) string {
    if yname == "is-error" { return "IsError" }
    if yname == "error" { return "Error" }
    if yname == "context" { return "Context" }
    return ""
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetSegmentPath() string {
    return "of-us-auth-error"
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-error"] = ofUsAuthError.IsError
    leafs["error"] = ofUsAuthError.Error
    leafs["context"] = ofUsAuthError.Context
    return leafs
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetBundleName() string { return "cisco_ios_xr" }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetYangName() string { return "of-us-auth-error" }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) SetParent(parent types.Entity) { ofUsAuthError.parent = parent }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetParent() types.Entity { return ofUsAuthError.parent }

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetParentYangName() string { return "sso-alert" }

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError
// Of-peer Authentication SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetFilter() yfilter.YFilter { return ofPeerAuthError.YFilter }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) SetFilter(yf yfilter.YFilter) { ofPeerAuthError.YFilter = yf }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetGoName(yname string) string {
    if yname == "is-error" { return "IsError" }
    if yname == "error" { return "Error" }
    if yname == "context" { return "Context" }
    return ""
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetSegmentPath() string {
    return "of-peer-auth-error"
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-error"] = ofPeerAuthError.IsError
    leafs["error"] = ofPeerAuthError.Error
    leafs["context"] = ofPeerAuthError.Context
    return leafs
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetBundleName() string { return "cisco_ios_xr" }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetYangName() string { return "of-peer-auth-error" }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) SetParent(parent types.Entity) { ofPeerAuthError.parent = parent }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetParent() types.Entity { return ofPeerAuthError.parent }

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetParentYangName() string { return "sso-alert" }

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError
// IPCP SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetFilter() yfilter.YFilter { return ipcpError.YFilter }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) SetFilter(yf yfilter.YFilter) { ipcpError.YFilter = yf }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetGoName(yname string) string {
    if yname == "is-error" { return "IsError" }
    if yname == "error" { return "Error" }
    if yname == "context" { return "Context" }
    return ""
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetSegmentPath() string {
    return "ipcp-error"
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-error"] = ipcpError.IsError
    leafs["error"] = ipcpError.Error
    leafs["context"] = ipcpError.Context
    return leafs
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetBundleName() string { return "cisco_ios_xr" }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetYangName() string { return "ipcp-error" }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) SetParent(parent types.Entity) { ipcpError.parent = parent }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetParent() types.Entity { return ipcpError.parent }

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetParentYangName() string { return "sso-alert" }

// Ppp_Nodes_Node_NodeInterfaceStatistics
// Per interface PPP operational statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LCP and NCP statistics for an interface running PPP. The type is slice of
    // Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic.
    NodeInterfaceStatistic []Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetFilter() yfilter.YFilter { return nodeInterfaceStatistics.YFilter }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) SetFilter(yf yfilter.YFilter) { nodeInterfaceStatistics.YFilter = yf }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetGoName(yname string) string {
    if yname == "node-interface-statistic" { return "NodeInterfaceStatistic" }
    return ""
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetSegmentPath() string {
    return "node-interface-statistics"
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-interface-statistic" {
        for _, c := range nodeInterfaceStatistics.NodeInterfaceStatistic {
            if nodeInterfaceStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic{}
        nodeInterfaceStatistics.NodeInterfaceStatistic = append(nodeInterfaceStatistics.NodeInterfaceStatistic, child)
        return &nodeInterfaceStatistics.NodeInterfaceStatistic[len(nodeInterfaceStatistics.NodeInterfaceStatistic)-1]
    }
    return nil
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeInterfaceStatistics.NodeInterfaceStatistic {
        children[nodeInterfaceStatistics.NodeInterfaceStatistic[i].GetSegmentPath()] = &nodeInterfaceStatistics.NodeInterfaceStatistic[i]
    }
    return children
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetYangName() string { return "node-interface-statistics" }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) SetParent(parent types.Entity) { nodeInterfaceStatistics.parent = parent }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetParent() types.Entity { return nodeInterfaceStatistics.parent }

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic
// LCP and NCP statistics for an interface
// running PPP
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface running PPP. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // PPP LCP Statistics.
    LcpStatistics Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics

    // PPP Authentication statistics.
    AuthenticationStatistics Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics

    // Array of PPP NCP Statistics. The type is slice of
    // Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray.
    NcpStatisticsArray []Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetFilter() yfilter.YFilter { return nodeInterfaceStatistic.YFilter }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) SetFilter(yf yfilter.YFilter) { nodeInterfaceStatistic.YFilter = yf }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "lcp-statistics" { return "LcpStatistics" }
    if yname == "authentication-statistics" { return "AuthenticationStatistics" }
    if yname == "ncp-statistics-array" { return "NcpStatisticsArray" }
    return ""
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetSegmentPath() string {
    return "node-interface-statistic" + "[interface-name='" + fmt.Sprintf("%v", nodeInterfaceStatistic.InterfaceName) + "']"
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lcp-statistics" {
        return &nodeInterfaceStatistic.LcpStatistics
    }
    if childYangName == "authentication-statistics" {
        return &nodeInterfaceStatistic.AuthenticationStatistics
    }
    if childYangName == "ncp-statistics-array" {
        for _, c := range nodeInterfaceStatistic.NcpStatisticsArray {
            if nodeInterfaceStatistic.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray{}
        nodeInterfaceStatistic.NcpStatisticsArray = append(nodeInterfaceStatistic.NcpStatisticsArray, child)
        return &nodeInterfaceStatistic.NcpStatisticsArray[len(nodeInterfaceStatistic.NcpStatisticsArray)-1]
    }
    return nil
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lcp-statistics"] = &nodeInterfaceStatistic.LcpStatistics
    children["authentication-statistics"] = &nodeInterfaceStatistic.AuthenticationStatistics
    for i := range nodeInterfaceStatistic.NcpStatisticsArray {
        children[nodeInterfaceStatistic.NcpStatisticsArray[i].GetSegmentPath()] = &nodeInterfaceStatistic.NcpStatisticsArray[i]
    }
    return children
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = nodeInterfaceStatistic.InterfaceName
    return leafs
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetYangName() string { return "node-interface-statistic" }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) SetParent(parent types.Entity) { nodeInterfaceStatistic.parent = parent }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetParent() types.Entity { return nodeInterfaceStatistic.parent }

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetParentYangName() string { return "node-interface-statistics" }

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics
// PPP LCP Statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Conf Req Packets Sent. The type is interface{} with range: 0..65535.
    ConfReqSent interface{}

    // Conf Req Packets Received. The type is interface{} with range: 0..65535.
    ConfReqRcvd interface{}

    // Conf Ack Packets Sent. The type is interface{} with range: 0..65535.
    ConfAckSent interface{}

    // Conf Ack Packets Received. The type is interface{} with range: 0..65535.
    ConfAckRcvd interface{}

    // Conf Nak Packets Sent. The type is interface{} with range: 0..65535.
    ConfNakSent interface{}

    // Conf Nak Packets Received. The type is interface{} with range: 0..65535.
    ConfNakRcvd interface{}

    // Conf Rej Packets Sent. The type is interface{} with range: 0..65535.
    ConfRejSent interface{}

    // Conf Rej Packets Received. The type is interface{} with range: 0..65535.
    ConfRejRcvd interface{}

    // Echo Req Packets Sent. The type is interface{} with range: 0..65535.
    EchoReqSent interface{}

    // Echo Req Packets Received. The type is interface{} with range: 0..65535.
    EchoReqRcvd interface{}

    // Echo Rep Packets Sent. The type is interface{} with range: 0..65535.
    EchoRepSent interface{}

    // Echo Rep Packets Received. The type is interface{} with range: 0..65535.
    EchoRepRcvd interface{}

    // Disc Req Packets Sent. The type is interface{} with range: 0..65535.
    DiscReqSent interface{}

    // Disc Req Packets Received. The type is interface{} with range: 0..65535.
    DiscReqRcvd interface{}

    // Line Protocol Up count. The type is interface{} with range: 0..65535.
    LinkUp interface{}

    // Keepalive link failure count. The type is interface{} with range: 0..65535.
    LinkError interface{}
}

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetFilter() yfilter.YFilter { return lcpStatistics.YFilter }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) SetFilter(yf yfilter.YFilter) { lcpStatistics.YFilter = yf }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetGoName(yname string) string {
    if yname == "conf-req-sent" { return "ConfReqSent" }
    if yname == "conf-req-rcvd" { return "ConfReqRcvd" }
    if yname == "conf-ack-sent" { return "ConfAckSent" }
    if yname == "conf-ack-rcvd" { return "ConfAckRcvd" }
    if yname == "conf-nak-sent" { return "ConfNakSent" }
    if yname == "conf-nak-rcvd" { return "ConfNakRcvd" }
    if yname == "conf-rej-sent" { return "ConfRejSent" }
    if yname == "conf-rej-rcvd" { return "ConfRejRcvd" }
    if yname == "echo-req-sent" { return "EchoReqSent" }
    if yname == "echo-req-rcvd" { return "EchoReqRcvd" }
    if yname == "echo-rep-sent" { return "EchoRepSent" }
    if yname == "echo-rep-rcvd" { return "EchoRepRcvd" }
    if yname == "disc-req-sent" { return "DiscReqSent" }
    if yname == "disc-req-rcvd" { return "DiscReqRcvd" }
    if yname == "link-up" { return "LinkUp" }
    if yname == "link-error" { return "LinkError" }
    return ""
}

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetSegmentPath() string {
    return "lcp-statistics"
}

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["conf-req-sent"] = lcpStatistics.ConfReqSent
    leafs["conf-req-rcvd"] = lcpStatistics.ConfReqRcvd
    leafs["conf-ack-sent"] = lcpStatistics.ConfAckSent
    leafs["conf-ack-rcvd"] = lcpStatistics.ConfAckRcvd
    leafs["conf-nak-sent"] = lcpStatistics.ConfNakSent
    leafs["conf-nak-rcvd"] = lcpStatistics.ConfNakRcvd
    leafs["conf-rej-sent"] = lcpStatistics.ConfRejSent
    leafs["conf-rej-rcvd"] = lcpStatistics.ConfRejRcvd
    leafs["echo-req-sent"] = lcpStatistics.EchoReqSent
    leafs["echo-req-rcvd"] = lcpStatistics.EchoReqRcvd
    leafs["echo-rep-sent"] = lcpStatistics.EchoRepSent
    leafs["echo-rep-rcvd"] = lcpStatistics.EchoRepRcvd
    leafs["disc-req-sent"] = lcpStatistics.DiscReqSent
    leafs["disc-req-rcvd"] = lcpStatistics.DiscReqRcvd
    leafs["link-up"] = lcpStatistics.LinkUp
    leafs["link-error"] = lcpStatistics.LinkError
    return leafs
}

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetYangName() string { return "lcp-statistics" }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) SetParent(parent types.Entity) { lcpStatistics.parent = parent }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetParent() types.Entity { return lcpStatistics.parent }

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetParentYangName() string { return "node-interface-statistic" }

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics
// PPP Authentication statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PAP Request packets sent. The type is interface{} with range: 0..65535.
    PapReqSent interface{}

    // PAP Request packets received. The type is interface{} with range: 0..65535.
    PapReqRcvd interface{}

    // PAP Ack packets sent. The type is interface{} with range: 0..65535.
    PapAckSent interface{}

    // PAP Ack packets received. The type is interface{} with range: 0..65535.
    PapAckRcvd interface{}

    // PAP Nak packets sent. The type is interface{} with range: 0..65535.
    PapNakSent interface{}

    // PAP Nak packets received. The type is interface{} with range: 0..65535.
    PapNakRcvd interface{}

    // CHAP challenge packets sent. The type is interface{} with range: 0..65535.
    ChapChallSent interface{}

    // CHAP challenge packets received. The type is interface{} with range:
    // 0..65535.
    ChapChallRcvd interface{}

    // CHAP response packets sent. The type is interface{} with range: 0..65535.
    ChapRespSent interface{}

    // CHAP response packets received. The type is interface{} with range:
    // 0..65535.
    ChapRespRcvd interface{}

    // CHAP reply success packets sent. The type is interface{} with range:
    // 0..65535.
    ChapRepSuccSent interface{}

    // CHAP reply success packets received. The type is interface{} with range:
    // 0..65535.
    ChapRepSuccRcvd interface{}

    // CHAP reply failure packets sent. The type is interface{} with range:
    // 0..65535.
    ChapRepFailSent interface{}

    // CHAP reply failure packets received. The type is interface{} with range:
    // 0..65535.
    ChapRepFailRcvd interface{}

    // Authentication timeout count. The type is interface{} with range: 0..65535.
    AuthTimeoutCount interface{}
}

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetFilter() yfilter.YFilter { return authenticationStatistics.YFilter }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) SetFilter(yf yfilter.YFilter) { authenticationStatistics.YFilter = yf }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetGoName(yname string) string {
    if yname == "pap-req-sent" { return "PapReqSent" }
    if yname == "pap-req-rcvd" { return "PapReqRcvd" }
    if yname == "pap-ack-sent" { return "PapAckSent" }
    if yname == "pap-ack-rcvd" { return "PapAckRcvd" }
    if yname == "pap-nak-sent" { return "PapNakSent" }
    if yname == "pap-nak-rcvd" { return "PapNakRcvd" }
    if yname == "chap-chall-sent" { return "ChapChallSent" }
    if yname == "chap-chall-rcvd" { return "ChapChallRcvd" }
    if yname == "chap-resp-sent" { return "ChapRespSent" }
    if yname == "chap-resp-rcvd" { return "ChapRespRcvd" }
    if yname == "chap-rep-succ-sent" { return "ChapRepSuccSent" }
    if yname == "chap-rep-succ-rcvd" { return "ChapRepSuccRcvd" }
    if yname == "chap-rep-fail-sent" { return "ChapRepFailSent" }
    if yname == "chap-rep-fail-rcvd" { return "ChapRepFailRcvd" }
    if yname == "auth-timeout-count" { return "AuthTimeoutCount" }
    return ""
}

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetSegmentPath() string {
    return "authentication-statistics"
}

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pap-req-sent"] = authenticationStatistics.PapReqSent
    leafs["pap-req-rcvd"] = authenticationStatistics.PapReqRcvd
    leafs["pap-ack-sent"] = authenticationStatistics.PapAckSent
    leafs["pap-ack-rcvd"] = authenticationStatistics.PapAckRcvd
    leafs["pap-nak-sent"] = authenticationStatistics.PapNakSent
    leafs["pap-nak-rcvd"] = authenticationStatistics.PapNakRcvd
    leafs["chap-chall-sent"] = authenticationStatistics.ChapChallSent
    leafs["chap-chall-rcvd"] = authenticationStatistics.ChapChallRcvd
    leafs["chap-resp-sent"] = authenticationStatistics.ChapRespSent
    leafs["chap-resp-rcvd"] = authenticationStatistics.ChapRespRcvd
    leafs["chap-rep-succ-sent"] = authenticationStatistics.ChapRepSuccSent
    leafs["chap-rep-succ-rcvd"] = authenticationStatistics.ChapRepSuccRcvd
    leafs["chap-rep-fail-sent"] = authenticationStatistics.ChapRepFailSent
    leafs["chap-rep-fail-rcvd"] = authenticationStatistics.ChapRepFailRcvd
    leafs["auth-timeout-count"] = authenticationStatistics.AuthTimeoutCount
    return leafs
}

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetYangName() string { return "authentication-statistics" }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) SetParent(parent types.Entity) { authenticationStatistics.parent = parent }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetParent() types.Entity { return authenticationStatistics.parent }

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetParentYangName() string { return "node-interface-statistic" }

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray
// Array of PPP NCP Statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NCP identifier. The type is NcpIdent.
    NcpIdentifier interface{}

    // Conf Req Packets Sent. The type is interface{} with range: 0..65535.
    ConfReqSent interface{}

    // Conf Req Packets Received. The type is interface{} with range: 0..65535.
    ConfReqRcvd interface{}

    // Conf Ack Packets Sent. The type is interface{} with range: 0..65535.
    ConfAckSent interface{}

    // Conf Ack Packets Received. The type is interface{} with range: 0..65535.
    ConfAckRcvd interface{}

    // Conf Nak Packets Sent. The type is interface{} with range: 0..65535.
    ConfNakSent interface{}

    // Conf Nak Packets Received. The type is interface{} with range: 0..65535.
    ConfNakRcvd interface{}

    // Conf Rej Packets Sent. The type is interface{} with range: 0..65535.
    ConfRejSent interface{}

    // Conf Rej Packets Received. The type is interface{} with range: 0..65535.
    ConfRejRcvd interface{}
}

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetFilter() yfilter.YFilter { return ncpStatisticsArray.YFilter }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) SetFilter(yf yfilter.YFilter) { ncpStatisticsArray.YFilter = yf }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetGoName(yname string) string {
    if yname == "ncp-identifier" { return "NcpIdentifier" }
    if yname == "conf-req-sent" { return "ConfReqSent" }
    if yname == "conf-req-rcvd" { return "ConfReqRcvd" }
    if yname == "conf-ack-sent" { return "ConfAckSent" }
    if yname == "conf-ack-rcvd" { return "ConfAckRcvd" }
    if yname == "conf-nak-sent" { return "ConfNakSent" }
    if yname == "conf-nak-rcvd" { return "ConfNakRcvd" }
    if yname == "conf-rej-sent" { return "ConfRejSent" }
    if yname == "conf-rej-rcvd" { return "ConfRejRcvd" }
    return ""
}

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetSegmentPath() string {
    return "ncp-statistics-array"
}

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ncp-identifier"] = ncpStatisticsArray.NcpIdentifier
    leafs["conf-req-sent"] = ncpStatisticsArray.ConfReqSent
    leafs["conf-req-rcvd"] = ncpStatisticsArray.ConfReqRcvd
    leafs["conf-ack-sent"] = ncpStatisticsArray.ConfAckSent
    leafs["conf-ack-rcvd"] = ncpStatisticsArray.ConfAckRcvd
    leafs["conf-nak-sent"] = ncpStatisticsArray.ConfNakSent
    leafs["conf-nak-rcvd"] = ncpStatisticsArray.ConfNakRcvd
    leafs["conf-rej-sent"] = ncpStatisticsArray.ConfRejSent
    leafs["conf-rej-rcvd"] = ncpStatisticsArray.ConfRejRcvd
    return leafs
}

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetBundleName() string { return "cisco_ios_xr" }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetYangName() string { return "ncp-statistics-array" }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) SetParent(parent types.Entity) { ncpStatisticsArray.parent = parent }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetParent() types.Entity { return ncpStatisticsArray.parent }

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetParentYangName() string { return "node-interface-statistic" }

// Ppp_Nodes_Node_SsoSummary
// Summarized PPP SSO data for a particular node
type Ppp_Nodes_Node_SsoSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LCP SSO FSM States.
    LcpStates Ppp_Nodes_Node_SsoSummary_LcpStates

    // Of-us Authentication SSO FSM States.
    OfUsAuthStates Ppp_Nodes_Node_SsoSummary_OfUsAuthStates

    // Of-peer Authentication SSO FSM States.
    OfPeerAuthStates Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates

    // IPCP SSO FSM States.
    IpcpStates Ppp_Nodes_Node_SsoSummary_IpcpStates
}

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetFilter() yfilter.YFilter { return ssoSummary.YFilter }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) SetFilter(yf yfilter.YFilter) { ssoSummary.YFilter = yf }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetGoName(yname string) string {
    if yname == "lcp-states" { return "LcpStates" }
    if yname == "of-us-auth-states" { return "OfUsAuthStates" }
    if yname == "of-peer-auth-states" { return "OfPeerAuthStates" }
    if yname == "ipcp-states" { return "IpcpStates" }
    return ""
}

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetSegmentPath() string {
    return "sso-summary"
}

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lcp-states" {
        return &ssoSummary.LcpStates
    }
    if childYangName == "of-us-auth-states" {
        return &ssoSummary.OfUsAuthStates
    }
    if childYangName == "of-peer-auth-states" {
        return &ssoSummary.OfPeerAuthStates
    }
    if childYangName == "ipcp-states" {
        return &ssoSummary.IpcpStates
    }
    return nil
}

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lcp-states"] = &ssoSummary.LcpStates
    children["of-us-auth-states"] = &ssoSummary.OfUsAuthStates
    children["of-peer-auth-states"] = &ssoSummary.OfPeerAuthStates
    children["ipcp-states"] = &ssoSummary.IpcpStates
    return children
}

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetBundleName() string { return "cisco_ios_xr" }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetYangName() string { return "sso-summary" }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) SetParent(parent types.Entity) { ssoSummary.parent = parent }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetParent() types.Entity { return ssoSummary.parent }

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_SsoSummary_LcpStates
// LCP SSO FSM States
type Ppp_Nodes_Node_SsoSummary_LcpStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetFilter() yfilter.YFilter { return lcpStates.YFilter }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) SetFilter(yf yfilter.YFilter) { lcpStates.YFilter = yf }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "count" { return "Count" }
    return ""
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetSegmentPath() string {
    return "lcp-states"
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = lcpStates.Total
    leafs["count"] = lcpStates.Count
    return leafs
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetBundleName() string { return "cisco_ios_xr" }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetYangName() string { return "lcp-states" }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) SetParent(parent types.Entity) { lcpStates.parent = parent }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetParent() types.Entity { return lcpStates.parent }

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetParentYangName() string { return "sso-summary" }

// Ppp_Nodes_Node_SsoSummary_OfUsAuthStates
// Of-us Authentication SSO FSM States
type Ppp_Nodes_Node_SsoSummary_OfUsAuthStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetFilter() yfilter.YFilter { return ofUsAuthStates.YFilter }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) SetFilter(yf yfilter.YFilter) { ofUsAuthStates.YFilter = yf }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "count" { return "Count" }
    return ""
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetSegmentPath() string {
    return "of-us-auth-states"
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = ofUsAuthStates.Total
    leafs["count"] = ofUsAuthStates.Count
    return leafs
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetBundleName() string { return "cisco_ios_xr" }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetYangName() string { return "of-us-auth-states" }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) SetParent(parent types.Entity) { ofUsAuthStates.parent = parent }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetParent() types.Entity { return ofUsAuthStates.parent }

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetParentYangName() string { return "sso-summary" }

// Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates
// Of-peer Authentication SSO FSM States
type Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetFilter() yfilter.YFilter { return ofPeerAuthStates.YFilter }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) SetFilter(yf yfilter.YFilter) { ofPeerAuthStates.YFilter = yf }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "count" { return "Count" }
    return ""
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetSegmentPath() string {
    return "of-peer-auth-states"
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = ofPeerAuthStates.Total
    leafs["count"] = ofPeerAuthStates.Count
    return leafs
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetBundleName() string { return "cisco_ios_xr" }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetYangName() string { return "of-peer-auth-states" }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) SetParent(parent types.Entity) { ofPeerAuthStates.parent = parent }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetParent() types.Entity { return ofPeerAuthStates.parent }

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetParentYangName() string { return "sso-summary" }

// Ppp_Nodes_Node_SsoSummary_IpcpStates
// IPCP SSO FSM States
type Ppp_Nodes_Node_SsoSummary_IpcpStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetFilter() yfilter.YFilter { return ipcpStates.YFilter }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) SetFilter(yf yfilter.YFilter) { ipcpStates.YFilter = yf }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "count" { return "Count" }
    return ""
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetSegmentPath() string {
    return "ipcp-states"
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = ipcpStates.Total
    leafs["count"] = ipcpStates.Count
    return leafs
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetBundleName() string { return "cisco_ios_xr" }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetYangName() string { return "ipcp-states" }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) SetParent(parent types.Entity) { ipcpStates.parent = parent }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetParent() types.Entity { return ipcpStates.parent }

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetParentYangName() string { return "sso-summary" }

// Ppp_Nodes_Node_SsoGroups
// PPP SSO Group data for a particular node
type Ppp_Nodes_Node_SsoGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPP SSO state data for a particular group. The type is slice of
    // Ppp_Nodes_Node_SsoGroups_SsoGroup.
    SsoGroup []Ppp_Nodes_Node_SsoGroups_SsoGroup
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetFilter() yfilter.YFilter { return ssoGroups.YFilter }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) SetFilter(yf yfilter.YFilter) { ssoGroups.YFilter = yf }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetGoName(yname string) string {
    if yname == "sso-group" { return "SsoGroup" }
    return ""
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetSegmentPath() string {
    return "sso-groups"
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sso-group" {
        for _, c := range ssoGroups.SsoGroup {
            if ssoGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_SsoGroups_SsoGroup{}
        ssoGroups.SsoGroup = append(ssoGroups.SsoGroup, child)
        return &ssoGroups.SsoGroup[len(ssoGroups.SsoGroup)-1]
    }
    return nil
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssoGroups.SsoGroup {
        children[ssoGroups.SsoGroup[i].GetSegmentPath()] = &ssoGroups.SsoGroup[i]
    }
    return children
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetBundleName() string { return "cisco_ios_xr" }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetYangName() string { return "sso-groups" }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) SetParent(parent types.Entity) { ssoGroups.parent = parent }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetParent() types.Entity { return ssoGroups.parent }

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup
// PPP SSO state data for a particular group
type Ppp_Nodes_Node_SsoGroups_SsoGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the group. The type is
    // interface{} with range: 1..65535.
    GroupId interface{}

    // PPP SSO State data for a particular group.
    SsoStates Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetFilter() yfilter.YFilter { return ssoGroup.YFilter }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) SetFilter(yf yfilter.YFilter) { ssoGroup.YFilter = yf }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "sso-states" { return "SsoStates" }
    return ""
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetSegmentPath() string {
    return "sso-group" + "[group-id='" + fmt.Sprintf("%v", ssoGroup.GroupId) + "']"
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sso-states" {
        return &ssoGroup.SsoStates
    }
    return nil
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sso-states"] = &ssoGroup.SsoStates
    return children
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = ssoGroup.GroupId
    return leafs
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetBundleName() string { return "cisco_ios_xr" }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetYangName() string { return "sso-group" }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) SetParent(parent types.Entity) { ssoGroup.parent = parent }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetParent() types.Entity { return ssoGroup.parent }

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetParentYangName() string { return "sso-groups" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates
// PPP SSO State data for a particular group
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PPP SSO State data for a particular interface. The type is slice of
    // Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState.
    SsoState []Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetFilter() yfilter.YFilter { return ssoStates.YFilter }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) SetFilter(yf yfilter.YFilter) { ssoStates.YFilter = yf }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetGoName(yname string) string {
    if yname == "sso-state" { return "SsoState" }
    return ""
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetSegmentPath() string {
    return "sso-states"
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sso-state" {
        for _, c := range ssoStates.SsoState {
            if ssoStates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState{}
        ssoStates.SsoState = append(ssoStates.SsoState, child)
        return &ssoStates.SsoState[len(ssoStates.SsoState)-1]
    }
    return nil
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ssoStates.SsoState {
        children[ssoStates.SsoState[i].GetSegmentPath()] = &ssoStates.SsoState[i]
    }
    return children
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetBundleName() string { return "cisco_ios_xr" }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetYangName() string { return "sso-states" }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) SetParent(parent types.Entity) { ssoStates.parent = parent }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetParent() types.Entity { return ssoStates.parent }

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetParentYangName() string { return "sso-group" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState
// PPP SSO State data for a particular
// interface
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Session ID for the interface with SSO State. The
    // type is interface{} with range: 1..4294967295.
    SessionId interface{}

    // SSRP Session ID. The type is interface{} with range: 0..4294967295.
    SessionIdXr interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9./-]+.
    Interface interface{}

    // LCP SSO State.
    LcpState Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState

    // Of-us Authentication SSO State.
    OfUsAuthState Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState

    // Of-peer Authentication SSO State.
    OfPeerAuthState Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState

    // IPCP SSO State.
    IpcpState Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState
}

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetFilter() yfilter.YFilter { return ssoState.YFilter }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) SetFilter(yf yfilter.YFilter) { ssoState.YFilter = yf }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetGoName(yname string) string {
    if yname == "session-id" { return "SessionId" }
    if yname == "session-id-xr" { return "SessionIdXr" }
    if yname == "interface" { return "Interface" }
    if yname == "lcp-state" { return "LcpState" }
    if yname == "of-us-auth-state" { return "OfUsAuthState" }
    if yname == "of-peer-auth-state" { return "OfPeerAuthState" }
    if yname == "ipcp-state" { return "IpcpState" }
    return ""
}

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetSegmentPath() string {
    return "sso-state" + "[session-id='" + fmt.Sprintf("%v", ssoState.SessionId) + "']"
}

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lcp-state" {
        return &ssoState.LcpState
    }
    if childYangName == "of-us-auth-state" {
        return &ssoState.OfUsAuthState
    }
    if childYangName == "of-peer-auth-state" {
        return &ssoState.OfPeerAuthState
    }
    if childYangName == "ipcp-state" {
        return &ssoState.IpcpState
    }
    return nil
}

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lcp-state"] = &ssoState.LcpState
    children["of-us-auth-state"] = &ssoState.OfUsAuthState
    children["of-peer-auth-state"] = &ssoState.OfPeerAuthState
    children["ipcp-state"] = &ssoState.IpcpState
    return children
}

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-id"] = ssoState.SessionId
    leafs["session-id-xr"] = ssoState.SessionIdXr
    leafs["interface"] = ssoState.Interface
    return leafs
}

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetBundleName() string { return "cisco_ios_xr" }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetYangName() string { return "sso-state" }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) SetParent(parent types.Entity) { ssoState.parent = parent }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetParent() types.Entity { return ssoState.parent }

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetParentYangName() string { return "sso-states" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState
// LCP SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetFilter() yfilter.YFilter { return lcpState.YFilter }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) SetFilter(yf yfilter.YFilter) { lcpState.YFilter = yf }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetGoName(yname string) string {
    if yname == "is-running" { return "IsRunning" }
    if yname == "state" { return "State" }
    return ""
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetSegmentPath() string {
    return "lcp-state"
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-running"] = lcpState.IsRunning
    leafs["state"] = lcpState.State
    return leafs
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetBundleName() string { return "cisco_ios_xr" }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetYangName() string { return "lcp-state" }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) SetParent(parent types.Entity) { lcpState.parent = parent }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetParent() types.Entity { return lcpState.parent }

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetParentYangName() string { return "sso-state" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState
// Of-us Authentication SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetFilter() yfilter.YFilter { return ofUsAuthState.YFilter }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) SetFilter(yf yfilter.YFilter) { ofUsAuthState.YFilter = yf }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetGoName(yname string) string {
    if yname == "is-running" { return "IsRunning" }
    if yname == "state" { return "State" }
    return ""
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetSegmentPath() string {
    return "of-us-auth-state"
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-running"] = ofUsAuthState.IsRunning
    leafs["state"] = ofUsAuthState.State
    return leafs
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetBundleName() string { return "cisco_ios_xr" }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetYangName() string { return "of-us-auth-state" }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) SetParent(parent types.Entity) { ofUsAuthState.parent = parent }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetParent() types.Entity { return ofUsAuthState.parent }

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetParentYangName() string { return "sso-state" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState
// Of-peer Authentication SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetFilter() yfilter.YFilter { return ofPeerAuthState.YFilter }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) SetFilter(yf yfilter.YFilter) { ofPeerAuthState.YFilter = yf }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetGoName(yname string) string {
    if yname == "is-running" { return "IsRunning" }
    if yname == "state" { return "State" }
    return ""
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetSegmentPath() string {
    return "of-peer-auth-state"
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-running"] = ofPeerAuthState.IsRunning
    leafs["state"] = ofPeerAuthState.State
    return leafs
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetBundleName() string { return "cisco_ios_xr" }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetYangName() string { return "of-peer-auth-state" }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) SetParent(parent types.Entity) { ofPeerAuthState.parent = parent }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetParent() types.Entity { return ofPeerAuthState.parent }

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetParentYangName() string { return "sso-state" }

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState
// IPCP SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetFilter() yfilter.YFilter { return ipcpState.YFilter }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) SetFilter(yf yfilter.YFilter) { ipcpState.YFilter = yf }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetGoName(yname string) string {
    if yname == "is-running" { return "IsRunning" }
    if yname == "state" { return "State" }
    return ""
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetSegmentPath() string {
    return "ipcp-state"
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-running"] = ipcpState.IsRunning
    leafs["state"] = ipcpState.State
    return leafs
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetBundleName() string { return "cisco_ios_xr" }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetYangName() string { return "ipcp-state" }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) SetParent(parent types.Entity) { ipcpState.parent = parent }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetParent() types.Entity { return ipcpState.parent }

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetParentYangName() string { return "sso-state" }

// Ppp_Nodes_Node_Summary
// Summarized PPP data for a particular node
type Ppp_Nodes_Node_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interfaces running PPP.
    Intfs Ppp_Nodes_Node_Summary_Intfs

    // FSM States.
    FsmStates Ppp_Nodes_Node_Summary_FsmStates

    // LCP/Auth Phases.
    LcpAuthPhases Ppp_Nodes_Node_Summary_LcpAuthPhases
}

func (summary *Ppp_Nodes_Node_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *Ppp_Nodes_Node_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *Ppp_Nodes_Node_Summary) GetGoName(yname string) string {
    if yname == "intfs" { return "Intfs" }
    if yname == "fsm-states" { return "FsmStates" }
    if yname == "lcp-auth-phases" { return "LcpAuthPhases" }
    return ""
}

func (summary *Ppp_Nodes_Node_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *Ppp_Nodes_Node_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intfs" {
        return &summary.Intfs
    }
    if childYangName == "fsm-states" {
        return &summary.FsmStates
    }
    if childYangName == "lcp-auth-phases" {
        return &summary.LcpAuthPhases
    }
    return nil
}

func (summary *Ppp_Nodes_Node_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intfs"] = &summary.Intfs
    children["fsm-states"] = &summary.FsmStates
    children["lcp-auth-phases"] = &summary.LcpAuthPhases
    return children
}

func (summary *Ppp_Nodes_Node_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (summary *Ppp_Nodes_Node_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *Ppp_Nodes_Node_Summary) GetYangName() string { return "summary" }

func (summary *Ppp_Nodes_Node_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *Ppp_Nodes_Node_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *Ppp_Nodes_Node_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *Ppp_Nodes_Node_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *Ppp_Nodes_Node_Summary) GetParent() types.Entity { return summary.parent }

func (summary *Ppp_Nodes_Node_Summary) GetParentYangName() string { return "node" }

// Ppp_Nodes_Node_Summary_Intfs
// Interfaces running PPP
type Ppp_Nodes_Node_Summary_Intfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // POS Count. The type is interface{} with range: 0..4294967295.
    PosCount interface{}

    // Serial Count. The type is interface{} with range: 0..4294967295.
    SerialCount interface{}

    // PPPoE Count. The type is interface{} with range: 0..4294967295.
    PppoeCount interface{}

    // Multilink Bundle Count. The type is interface{} with range: 0..4294967295.
    MultilinkBundleCount interface{}

    // GCC0 Count. The type is interface{} with range: 0..4294967295.
    Gcc0Count interface{}

    // GCC1 Count. The type is interface{} with range: 0..4294967295.
    Gcc1Count interface{}

    // Total Count. The type is interface{} with range: 0..4294967295.
    Total interface{}
}

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetFilter() yfilter.YFilter { return intfs.YFilter }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) SetFilter(yf yfilter.YFilter) { intfs.YFilter = yf }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetGoName(yname string) string {
    if yname == "pos-count" { return "PosCount" }
    if yname == "serial-count" { return "SerialCount" }
    if yname == "pppoe-count" { return "PppoeCount" }
    if yname == "multilink-bundle-count" { return "MultilinkBundleCount" }
    if yname == "gcc0-count" { return "Gcc0Count" }
    if yname == "gcc1-count" { return "Gcc1Count" }
    if yname == "total" { return "Total" }
    return ""
}

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetSegmentPath() string {
    return "intfs"
}

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pos-count"] = intfs.PosCount
    leafs["serial-count"] = intfs.SerialCount
    leafs["pppoe-count"] = intfs.PppoeCount
    leafs["multilink-bundle-count"] = intfs.MultilinkBundleCount
    leafs["gcc0-count"] = intfs.Gcc0Count
    leafs["gcc1-count"] = intfs.Gcc1Count
    leafs["total"] = intfs.Total
    return leafs
}

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetBundleName() string { return "cisco_ios_xr" }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetYangName() string { return "intfs" }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) SetParent(parent types.Entity) { intfs.parent = parent }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetParent() types.Entity { return intfs.parent }

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetParentYangName() string { return "summary" }

// Ppp_Nodes_Node_Summary_FsmStates
// FSM States
type Ppp_Nodes_Node_Summary_FsmStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Array of per-LCP FSM States.
    LcpfsmStates Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates

    // Array of per-NCP FSM States. The type is slice of
    // Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray.
    NcpfsmStatesArray []Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetFilter() yfilter.YFilter { return fsmStates.YFilter }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) SetFilter(yf yfilter.YFilter) { fsmStates.YFilter = yf }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetGoName(yname string) string {
    if yname == "lcpfsm-states" { return "LcpfsmStates" }
    if yname == "ncpfsm-states-array" { return "NcpfsmStatesArray" }
    return ""
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetSegmentPath() string {
    return "fsm-states"
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lcpfsm-states" {
        return &fsmStates.LcpfsmStates
    }
    if childYangName == "ncpfsm-states-array" {
        for _, c := range fsmStates.NcpfsmStatesArray {
            if fsmStates.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray{}
        fsmStates.NcpfsmStatesArray = append(fsmStates.NcpfsmStatesArray, child)
        return &fsmStates.NcpfsmStatesArray[len(fsmStates.NcpfsmStatesArray)-1]
    }
    return nil
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lcpfsm-states"] = &fsmStates.LcpfsmStates
    for i := range fsmStates.NcpfsmStatesArray {
        children[fsmStates.NcpfsmStatesArray[i].GetSegmentPath()] = &fsmStates.NcpfsmStatesArray[i]
    }
    return children
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetBundleName() string { return "cisco_ios_xr" }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetYangName() string { return "fsm-states" }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) SetParent(parent types.Entity) { fsmStates.parent = parent }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetParent() types.Entity { return fsmStates.parent }

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetParentYangName() string { return "summary" }

// Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates
// Array of per-LCP FSM States
type Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total number of LCP FSMs running. The type is interface{} with range:
    // 0..4294967295.
    Total interface{}

    // Number of FSMs in each State. The type is slice of interface{} with range:
    // 0..4294967295.
    Count []interface{}
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetFilter() yfilter.YFilter { return lcpfsmStates.YFilter }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) SetFilter(yf yfilter.YFilter) { lcpfsmStates.YFilter = yf }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetGoName(yname string) string {
    if yname == "total" { return "Total" }
    if yname == "count" { return "Count" }
    return ""
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetSegmentPath() string {
    return "lcpfsm-states"
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["total"] = lcpfsmStates.Total
    leafs["count"] = lcpfsmStates.Count
    return leafs
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetBundleName() string { return "cisco_ios_xr" }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetYangName() string { return "lcpfsm-states" }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) SetParent(parent types.Entity) { lcpfsmStates.parent = parent }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetParent() types.Entity { return lcpfsmStates.parent }

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetParentYangName() string { return "fsm-states" }

// Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray
// Array of per-NCP FSM States
type Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NCP Identifier. The type is NcpIdent.
    NcpIdentifier interface{}

    // Total number of FSMs running for this NCP. The type is interface{} with
    // range: 0..4294967295.
    Total interface{}

    // Number of FSMs in each State. The type is slice of interface{} with range:
    // 0..4294967295.
    Count []interface{}
}

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetFilter() yfilter.YFilter { return ncpfsmStatesArray.YFilter }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) SetFilter(yf yfilter.YFilter) { ncpfsmStatesArray.YFilter = yf }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetGoName(yname string) string {
    if yname == "ncp-identifier" { return "NcpIdentifier" }
    if yname == "total" { return "Total" }
    if yname == "count" { return "Count" }
    return ""
}

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetSegmentPath() string {
    return "ncpfsm-states-array"
}

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ncp-identifier"] = ncpfsmStatesArray.NcpIdentifier
    leafs["total"] = ncpfsmStatesArray.Total
    leafs["count"] = ncpfsmStatesArray.Count
    return leafs
}

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetBundleName() string { return "cisco_ios_xr" }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetYangName() string { return "ncpfsm-states-array" }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) SetParent(parent types.Entity) { ncpfsmStatesArray.parent = parent }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetParent() types.Entity { return ncpfsmStatesArray.parent }

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetParentYangName() string { return "fsm-states" }

// Ppp_Nodes_Node_Summary_LcpAuthPhases
// LCP/Auth Phases
type Ppp_Nodes_Node_Summary_LcpAuthPhases struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of sessions with LCP not negotiated. The type is interface{} with
    // range: 0..4294967295.
    LcpNotNegotiated interface{}

    // Number of sessions authenticating. The type is interface{} with range:
    // 0..4294967295.
    Authenticating interface{}

    // Number of sessions negotiated but with the line held down. The type is
    // interface{} with range: 0..4294967295.
    LineHeldDown interface{}

    // Number of locally terminated sessions brought up. The type is interface{}
    // with range: 0..4294967295.
    UpLocalTerm interface{}

    // Number of L2 forwarded sessions brought up. The type is interface{} with
    // range: 0..4294967295.
    UpL2Fwded interface{}

    // Number of VPDN tunneled sessions brought up. The type is interface{} with
    // range: 0..4294967295.
    UpTunneled interface{}
}

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetFilter() yfilter.YFilter { return lcpAuthPhases.YFilter }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) SetFilter(yf yfilter.YFilter) { lcpAuthPhases.YFilter = yf }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetGoName(yname string) string {
    if yname == "lcp-not-negotiated" { return "LcpNotNegotiated" }
    if yname == "authenticating" { return "Authenticating" }
    if yname == "line-held-down" { return "LineHeldDown" }
    if yname == "up-local-term" { return "UpLocalTerm" }
    if yname == "up-l2-fwded" { return "UpL2Fwded" }
    if yname == "up-tunneled" { return "UpTunneled" }
    return ""
}

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetSegmentPath() string {
    return "lcp-auth-phases"
}

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lcp-not-negotiated"] = lcpAuthPhases.LcpNotNegotiated
    leafs["authenticating"] = lcpAuthPhases.Authenticating
    leafs["line-held-down"] = lcpAuthPhases.LineHeldDown
    leafs["up-local-term"] = lcpAuthPhases.UpLocalTerm
    leafs["up-l2-fwded"] = lcpAuthPhases.UpL2Fwded
    leafs["up-tunneled"] = lcpAuthPhases.UpTunneled
    return leafs
}

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetBundleName() string { return "cisco_ios_xr" }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetYangName() string { return "lcp-auth-phases" }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) SetParent(parent types.Entity) { lcpAuthPhases.parent = parent }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetParent() types.Entity { return lcpAuthPhases.parent }

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetParentYangName() string { return "summary" }

