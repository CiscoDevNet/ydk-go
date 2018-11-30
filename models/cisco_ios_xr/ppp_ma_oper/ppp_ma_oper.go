// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma package operational data.
// 
// This module contains definitions
// for the following management objects:
//   ppp: PPP operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per node PPP operational data.
    Nodes Ppp_Nodes
}

func (ppp *Ppp) GetEntityData() *types.CommonEntityData {
    ppp.EntityData.YFilter = ppp.YFilter
    ppp.EntityData.YangName = "ppp"
    ppp.EntityData.BundleName = "cisco_ios_xr"
    ppp.EntityData.ParentYangName = "Cisco-IOS-XR-ppp-ma-oper"
    ppp.EntityData.SegmentPath = "Cisco-IOS-XR-ppp-ma-oper:ppp"
    ppp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ppp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ppp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ppp.EntityData.Children = types.NewOrderedMap()
    ppp.EntityData.Children.Append("nodes", types.YChild{"Nodes", &ppp.Nodes})
    ppp.EntityData.Leafs = types.NewOrderedMap()

    ppp.EntityData.YListKeys = []string {}

    return &(ppp.EntityData)
}

// Ppp_Nodes
// Per node PPP operational data
type Ppp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The PPP operational data for a particular node. The type is slice of
    // Ppp_Nodes_Node.
    Node []*Ppp_Nodes_Node
}

func (nodes *Ppp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "ppp"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Ppp_Nodes_Node
// The PPP operational data for a particular node
type Ppp_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Ppp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Children.Append("node-interfaces", types.YChild{"NodeInterfaces", &node.NodeInterfaces})
    node.EntityData.Children.Append("sso-alerts", types.YChild{"SsoAlerts", &node.SsoAlerts})
    node.EntityData.Children.Append("node-interface-statistics", types.YChild{"NodeInterfaceStatistics", &node.NodeInterfaceStatistics})
    node.EntityData.Children.Append("sso-summary", types.YChild{"SsoSummary", &node.SsoSummary})
    node.EntityData.Children.Append("sso-groups", types.YChild{"SsoGroups", &node.SsoGroups})
    node.EntityData.Children.Append("summary", types.YChild{"Summary", &node.Summary})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Ppp_Nodes_Node_Statistics
// PPP statistics data for a particular node
type Ppp_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPP LCP Statistics.
    LcpStatistics Ppp_Nodes_Node_Statistics_LcpStatistics

    // PPP Authentication statistics.
    AuthenticationStatistics Ppp_Nodes_Node_Statistics_AuthenticationStatistics

    // Array of PPP NCP Statistics. The type is slice of
    // Ppp_Nodes_Node_Statistics_NcpStatisticsArray.
    NcpStatisticsArray []*Ppp_Nodes_Node_Statistics_NcpStatisticsArray
}

func (statistics *Ppp_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("lcp-statistics", types.YChild{"LcpStatistics", &statistics.LcpStatistics})
    statistics.EntityData.Children.Append("authentication-statistics", types.YChild{"AuthenticationStatistics", &statistics.AuthenticationStatistics})
    statistics.EntityData.Children.Append("ncp-statistics-array", types.YChild{"NcpStatisticsArray", nil})
    for i := range statistics.NcpStatisticsArray {
        statistics.EntityData.Children.Append(types.GetSegmentPath(statistics.NcpStatisticsArray[i]), types.YChild{"NcpStatisticsArray", statistics.NcpStatisticsArray[i]})
    }
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Ppp_Nodes_Node_Statistics_LcpStatistics
// PPP LCP Statistics
type Ppp_Nodes_Node_Statistics_LcpStatistics struct {
    EntityData types.CommonEntityData
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

func (lcpStatistics *Ppp_Nodes_Node_Statistics_LcpStatistics) GetEntityData() *types.CommonEntityData {
    lcpStatistics.EntityData.YFilter = lcpStatistics.YFilter
    lcpStatistics.EntityData.YangName = "lcp-statistics"
    lcpStatistics.EntityData.BundleName = "cisco_ios_xr"
    lcpStatistics.EntityData.ParentYangName = "statistics"
    lcpStatistics.EntityData.SegmentPath = "lcp-statistics"
    lcpStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpStatistics.EntityData.Children = types.NewOrderedMap()
    lcpStatistics.EntityData.Leafs = types.NewOrderedMap()
    lcpStatistics.EntityData.Leafs.Append("conf-req-sent", types.YLeaf{"ConfReqSent", lcpStatistics.ConfReqSent})
    lcpStatistics.EntityData.Leafs.Append("conf-req-rcvd", types.YLeaf{"ConfReqRcvd", lcpStatistics.ConfReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("conf-ack-sent", types.YLeaf{"ConfAckSent", lcpStatistics.ConfAckSent})
    lcpStatistics.EntityData.Leafs.Append("conf-ack-rcvd", types.YLeaf{"ConfAckRcvd", lcpStatistics.ConfAckRcvd})
    lcpStatistics.EntityData.Leafs.Append("conf-nak-sent", types.YLeaf{"ConfNakSent", lcpStatistics.ConfNakSent})
    lcpStatistics.EntityData.Leafs.Append("conf-nak-rcvd", types.YLeaf{"ConfNakRcvd", lcpStatistics.ConfNakRcvd})
    lcpStatistics.EntityData.Leafs.Append("conf-rej-sent", types.YLeaf{"ConfRejSent", lcpStatistics.ConfRejSent})
    lcpStatistics.EntityData.Leafs.Append("conf-rej-rcvd", types.YLeaf{"ConfRejRcvd", lcpStatistics.ConfRejRcvd})
    lcpStatistics.EntityData.Leafs.Append("term-req-sent", types.YLeaf{"TermReqSent", lcpStatistics.TermReqSent})
    lcpStatistics.EntityData.Leafs.Append("term-req-rcvd", types.YLeaf{"TermReqRcvd", lcpStatistics.TermReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("term-ack-sent", types.YLeaf{"TermAckSent", lcpStatistics.TermAckSent})
    lcpStatistics.EntityData.Leafs.Append("term-ack-rcvd", types.YLeaf{"TermAckRcvd", lcpStatistics.TermAckRcvd})
    lcpStatistics.EntityData.Leafs.Append("code-rej-sent", types.YLeaf{"CodeRejSent", lcpStatistics.CodeRejSent})
    lcpStatistics.EntityData.Leafs.Append("code-rej-rcvd", types.YLeaf{"CodeRejRcvd", lcpStatistics.CodeRejRcvd})
    lcpStatistics.EntityData.Leafs.Append("proto-rej-sent", types.YLeaf{"ProtoRejSent", lcpStatistics.ProtoRejSent})
    lcpStatistics.EntityData.Leafs.Append("proto-rej-rcvd", types.YLeaf{"ProtoRejRcvd", lcpStatistics.ProtoRejRcvd})
    lcpStatistics.EntityData.Leafs.Append("echo-req-sent", types.YLeaf{"EchoReqSent", lcpStatistics.EchoReqSent})
    lcpStatistics.EntityData.Leafs.Append("echo-req-rcvd", types.YLeaf{"EchoReqRcvd", lcpStatistics.EchoReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("echo-rep-sent", types.YLeaf{"EchoRepSent", lcpStatistics.EchoRepSent})
    lcpStatistics.EntityData.Leafs.Append("echo-rep-rcvd", types.YLeaf{"EchoRepRcvd", lcpStatistics.EchoRepRcvd})
    lcpStatistics.EntityData.Leafs.Append("disc-req-sent", types.YLeaf{"DiscReqSent", lcpStatistics.DiscReqSent})
    lcpStatistics.EntityData.Leafs.Append("disc-req-rcvd", types.YLeaf{"DiscReqRcvd", lcpStatistics.DiscReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("link-up", types.YLeaf{"LinkUp", lcpStatistics.LinkUp})
    lcpStatistics.EntityData.Leafs.Append("link-error", types.YLeaf{"LinkError", lcpStatistics.LinkError})

    lcpStatistics.EntityData.YListKeys = []string {}

    return &(lcpStatistics.EntityData)
}

// Ppp_Nodes_Node_Statistics_AuthenticationStatistics
// PPP Authentication statistics
type Ppp_Nodes_Node_Statistics_AuthenticationStatistics struct {
    EntityData types.CommonEntityData
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

func (authenticationStatistics *Ppp_Nodes_Node_Statistics_AuthenticationStatistics) GetEntityData() *types.CommonEntityData {
    authenticationStatistics.EntityData.YFilter = authenticationStatistics.YFilter
    authenticationStatistics.EntityData.YangName = "authentication-statistics"
    authenticationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authenticationStatistics.EntityData.ParentYangName = "statistics"
    authenticationStatistics.EntityData.SegmentPath = "authentication-statistics"
    authenticationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationStatistics.EntityData.Children = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs.Append("pap-req-sent", types.YLeaf{"PapReqSent", authenticationStatistics.PapReqSent})
    authenticationStatistics.EntityData.Leafs.Append("pap-req-rcvd", types.YLeaf{"PapReqRcvd", authenticationStatistics.PapReqRcvd})
    authenticationStatistics.EntityData.Leafs.Append("pap-ack-sent", types.YLeaf{"PapAckSent", authenticationStatistics.PapAckSent})
    authenticationStatistics.EntityData.Leafs.Append("pap-ack-rcvd", types.YLeaf{"PapAckRcvd", authenticationStatistics.PapAckRcvd})
    authenticationStatistics.EntityData.Leafs.Append("pap-nak-sent", types.YLeaf{"PapNakSent", authenticationStatistics.PapNakSent})
    authenticationStatistics.EntityData.Leafs.Append("pap-nak-rcvd", types.YLeaf{"PapNakRcvd", authenticationStatistics.PapNakRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-chall-sent", types.YLeaf{"ChapChallSent", authenticationStatistics.ChapChallSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-chall-rcvd", types.YLeaf{"ChapChallRcvd", authenticationStatistics.ChapChallRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-resp-sent", types.YLeaf{"ChapRespSent", authenticationStatistics.ChapRespSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-resp-rcvd", types.YLeaf{"ChapRespRcvd", authenticationStatistics.ChapRespRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-succ-sent", types.YLeaf{"ChapRepSuccSent", authenticationStatistics.ChapRepSuccSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-succ-rcvd", types.YLeaf{"ChapRepSuccRcvd", authenticationStatistics.ChapRepSuccRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-fail-sent", types.YLeaf{"ChapRepFailSent", authenticationStatistics.ChapRepFailSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-fail-rcvd", types.YLeaf{"ChapRepFailRcvd", authenticationStatistics.ChapRepFailRcvd})
    authenticationStatistics.EntityData.Leafs.Append("auth-timeout-count", types.YLeaf{"AuthTimeoutCount", authenticationStatistics.AuthTimeoutCount})

    authenticationStatistics.EntityData.YListKeys = []string {}

    return &(authenticationStatistics.EntityData)
}

// Ppp_Nodes_Node_Statistics_NcpStatisticsArray
// Array of PPP NCP Statistics
type Ppp_Nodes_Node_Statistics_NcpStatisticsArray struct {
    EntityData types.CommonEntityData
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

func (ncpStatisticsArray *Ppp_Nodes_Node_Statistics_NcpStatisticsArray) GetEntityData() *types.CommonEntityData {
    ncpStatisticsArray.EntityData.YFilter = ncpStatisticsArray.YFilter
    ncpStatisticsArray.EntityData.YangName = "ncp-statistics-array"
    ncpStatisticsArray.EntityData.BundleName = "cisco_ios_xr"
    ncpStatisticsArray.EntityData.ParentYangName = "statistics"
    ncpStatisticsArray.EntityData.SegmentPath = "ncp-statistics-array"
    ncpStatisticsArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncpStatisticsArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncpStatisticsArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncpStatisticsArray.EntityData.Children = types.NewOrderedMap()
    ncpStatisticsArray.EntityData.Leafs = types.NewOrderedMap()
    ncpStatisticsArray.EntityData.Leafs.Append("ncp-identifier", types.YLeaf{"NcpIdentifier", ncpStatisticsArray.NcpIdentifier})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-req-sent", types.YLeaf{"ConfReqSent", ncpStatisticsArray.ConfReqSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-req-rcvd", types.YLeaf{"ConfReqRcvd", ncpStatisticsArray.ConfReqRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-ack-sent", types.YLeaf{"ConfAckSent", ncpStatisticsArray.ConfAckSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-ack-rcvd", types.YLeaf{"ConfAckRcvd", ncpStatisticsArray.ConfAckRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-nak-sent", types.YLeaf{"ConfNakSent", ncpStatisticsArray.ConfNakSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-nak-rcvd", types.YLeaf{"ConfNakRcvd", ncpStatisticsArray.ConfNakRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-rej-sent", types.YLeaf{"ConfRejSent", ncpStatisticsArray.ConfRejSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-rej-rcvd", types.YLeaf{"ConfRejRcvd", ncpStatisticsArray.ConfRejRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("term-req-sent", types.YLeaf{"TermReqSent", ncpStatisticsArray.TermReqSent})
    ncpStatisticsArray.EntityData.Leafs.Append("term-req-rcvd", types.YLeaf{"TermReqRcvd", ncpStatisticsArray.TermReqRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("term-ack-sent", types.YLeaf{"TermAckSent", ncpStatisticsArray.TermAckSent})
    ncpStatisticsArray.EntityData.Leafs.Append("term-ack-rcvd", types.YLeaf{"TermAckRcvd", ncpStatisticsArray.TermAckRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("proto-rej-sent", types.YLeaf{"ProtoRejSent", ncpStatisticsArray.ProtoRejSent})
    ncpStatisticsArray.EntityData.Leafs.Append("proto-rej-rcvd", types.YLeaf{"ProtoRejRcvd", ncpStatisticsArray.ProtoRejRcvd})

    ncpStatisticsArray.EntityData.YListKeys = []string {}

    return &(ncpStatisticsArray.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces
// Per interface PPP operational data
type Ppp_Nodes_Node_NodeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LCP and summarized NCP data for an interface running PPP. The type is slice
    // of Ppp_Nodes_Node_NodeInterfaces_NodeInterface.
    NodeInterface []*Ppp_Nodes_Node_NodeInterfaces_NodeInterface
}

func (nodeInterfaces *Ppp_Nodes_Node_NodeInterfaces) GetEntityData() *types.CommonEntityData {
    nodeInterfaces.EntityData.YFilter = nodeInterfaces.YFilter
    nodeInterfaces.EntityData.YangName = "node-interfaces"
    nodeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaces.EntityData.ParentYangName = "node"
    nodeInterfaces.EntityData.SegmentPath = "node-interfaces"
    nodeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaces.EntityData.Children = types.NewOrderedMap()
    nodeInterfaces.EntityData.Children.Append("node-interface", types.YChild{"NodeInterface", nil})
    for i := range nodeInterfaces.NodeInterface {
        nodeInterfaces.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaces.NodeInterface[i]), types.YChild{"NodeInterface", nodeInterfaces.NodeInterface[i]})
    }
    nodeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    nodeInterfaces.EntityData.YListKeys = []string {}

    return &(nodeInterfaces.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface
// LCP and summarized NCP data for an interface
// running PPP
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface running PPP. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
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
    IsL2ac interface{}

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
    NcpInfoArray []*Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray
}

func (nodeInterface *Ppp_Nodes_Node_NodeInterfaces_NodeInterface) GetEntityData() *types.CommonEntityData {
    nodeInterface.EntityData.YFilter = nodeInterface.YFilter
    nodeInterface.EntityData.YangName = "node-interface"
    nodeInterface.EntityData.BundleName = "cisco_ios_xr"
    nodeInterface.EntityData.ParentYangName = "node-interfaces"
    nodeInterface.EntityData.SegmentPath = "node-interface" + types.AddKeyToken(nodeInterface.Interface, "interface")
    nodeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterface.EntityData.Children = types.NewOrderedMap()
    nodeInterface.EntityData.Children.Append("mp-info", types.YChild{"MpInfo", &nodeInterface.MpInfo})
    nodeInterface.EntityData.Children.Append("configured-timeout", types.YChild{"ConfiguredTimeout", &nodeInterface.ConfiguredTimeout})
    nodeInterface.EntityData.Children.Append("auth-info", types.YChild{"AuthInfo", &nodeInterface.AuthInfo})
    nodeInterface.EntityData.Children.Append("ncp-info-array", types.YChild{"NcpInfoArray", nil})
    for i := range nodeInterface.NcpInfoArray {
        nodeInterface.EntityData.Children.Append(types.GetSegmentPath(nodeInterface.NcpInfoArray[i]), types.YChild{"NcpInfoArray", nodeInterface.NcpInfoArray[i]})
    }
    nodeInterface.EntityData.Leafs = types.NewOrderedMap()
    nodeInterface.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", nodeInterface.Interface})
    nodeInterface.EntityData.Leafs.Append("parent-state", types.YLeaf{"ParentState", nodeInterface.ParentState})
    nodeInterface.EntityData.Leafs.Append("line-state", types.YLeaf{"LineState", nodeInterface.LineState})
    nodeInterface.EntityData.Leafs.Append("is-loopback-detected", types.YLeaf{"IsLoopbackDetected", nodeInterface.IsLoopbackDetected})
    nodeInterface.EntityData.Leafs.Append("caps-idb-srg-role", types.YLeaf{"CapsIdbSrgRole", nodeInterface.CapsIdbSrgRole})
    nodeInterface.EntityData.Leafs.Append("session-srg-role", types.YLeaf{"SessionSrgRole", nodeInterface.SessionSrgRole})
    nodeInterface.EntityData.Leafs.Append("keepalive-period", types.YLeaf{"KeepalivePeriod", nodeInterface.KeepalivePeriod})
    nodeInterface.EntityData.Leafs.Append("keepalive-retry-count", types.YLeaf{"KeepaliveRetryCount", nodeInterface.KeepaliveRetryCount})
    nodeInterface.EntityData.Leafs.Append("is-ssrp-configured", types.YLeaf{"IsSsrpConfigured", nodeInterface.IsSsrpConfigured})
    nodeInterface.EntityData.Leafs.Append("is-l2ac", types.YLeaf{"IsL2ac", nodeInterface.IsL2ac})
    nodeInterface.EntityData.Leafs.Append("provisioned", types.YLeaf{"Provisioned", nodeInterface.Provisioned})
    nodeInterface.EntityData.Leafs.Append("ip-interworking-enabled", types.YLeaf{"IpInterworkingEnabled", nodeInterface.IpInterworkingEnabled})
    nodeInterface.EntityData.Leafs.Append("xconnect-id", types.YLeaf{"XconnectId", nodeInterface.XconnectId})
    nodeInterface.EntityData.Leafs.Append("is-tunneled-session", types.YLeaf{"IsTunneledSession", nodeInterface.IsTunneledSession})
    nodeInterface.EntityData.Leafs.Append("ssrp-peer-id", types.YLeaf{"SsrpPeerId", nodeInterface.SsrpPeerId})
    nodeInterface.EntityData.Leafs.Append("lcp-state", types.YLeaf{"LcpState", nodeInterface.LcpState})
    nodeInterface.EntityData.Leafs.Append("lcpsso-state", types.YLeaf{"LcpssoState", nodeInterface.LcpssoState})
    nodeInterface.EntityData.Leafs.Append("is-lcp-delayed", types.YLeaf{"IsLcpDelayed", nodeInterface.IsLcpDelayed})
    nodeInterface.EntityData.Leafs.Append("local-mru", types.YLeaf{"LocalMru", nodeInterface.LocalMru})
    nodeInterface.EntityData.Leafs.Append("peer-mru", types.YLeaf{"PeerMru", nodeInterface.PeerMru})
    nodeInterface.EntityData.Leafs.Append("local-mrru", types.YLeaf{"LocalMrru", nodeInterface.LocalMrru})
    nodeInterface.EntityData.Leafs.Append("peer-mrru", types.YLeaf{"PeerMrru", nodeInterface.PeerMrru})
    nodeInterface.EntityData.Leafs.Append("local-ed", types.YLeaf{"LocalEd", nodeInterface.LocalEd})
    nodeInterface.EntityData.Leafs.Append("peer-ed", types.YLeaf{"PeerEd", nodeInterface.PeerEd})
    nodeInterface.EntityData.Leafs.Append("is-mcmp-enabled", types.YLeaf{"IsMcmpEnabled", nodeInterface.IsMcmpEnabled})
    nodeInterface.EntityData.Leafs.Append("local-mcmp-classes", types.YLeaf{"LocalMcmpClasses", nodeInterface.LocalMcmpClasses})
    nodeInterface.EntityData.Leafs.Append("peer-mcmp-classes", types.YLeaf{"PeerMcmpClasses", nodeInterface.PeerMcmpClasses})
    nodeInterface.EntityData.Leafs.Append("session-expires", types.YLeaf{"SessionExpires", nodeInterface.SessionExpires})

    nodeInterface.EntityData.YListKeys = []string {"Interface"}

    return &(nodeInterface.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo
// MP information
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is an MP bundle. The type is bool.
    IsMpBundle interface{}

    // MP Bundle Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
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
    MpMemberInfoArray []*Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray
}

func (mpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo) GetEntityData() *types.CommonEntityData {
    mpInfo.EntityData.YFilter = mpInfo.YFilter
    mpInfo.EntityData.YangName = "mp-info"
    mpInfo.EntityData.BundleName = "cisco_ios_xr"
    mpInfo.EntityData.ParentYangName = "node-interface"
    mpInfo.EntityData.SegmentPath = "mp-info"
    mpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpInfo.EntityData.Children = types.NewOrderedMap()
    mpInfo.EntityData.Children.Append("mp-member-info-array", types.YChild{"MpMemberInfoArray", nil})
    for i := range mpInfo.MpMemberInfoArray {
        mpInfo.EntityData.Children.Append(types.GetSegmentPath(mpInfo.MpMemberInfoArray[i]), types.YChild{"MpMemberInfoArray", mpInfo.MpMemberInfoArray[i]})
    }
    mpInfo.EntityData.Leafs = types.NewOrderedMap()
    mpInfo.EntityData.Leafs.Append("is-mp-bundle", types.YLeaf{"IsMpBundle", mpInfo.IsMpBundle})
    mpInfo.EntityData.Leafs.Append("mp-bundle-interface", types.YLeaf{"MpBundleInterface", mpInfo.MpBundleInterface})
    mpInfo.EntityData.Leafs.Append("is-mp-bundle-member", types.YLeaf{"IsMpBundleMember", mpInfo.IsMpBundleMember})
    mpInfo.EntityData.Leafs.Append("mp-group", types.YLeaf{"MpGroup", mpInfo.MpGroup})
    mpInfo.EntityData.Leafs.Append("active-links", types.YLeaf{"ActiveLinks", mpInfo.ActiveLinks})
    mpInfo.EntityData.Leafs.Append("inactive-links", types.YLeaf{"InactiveLinks", mpInfo.InactiveLinks})
    mpInfo.EntityData.Leafs.Append("minimum-active-links", types.YLeaf{"MinimumActiveLinks", mpInfo.MinimumActiveLinks})
    mpInfo.EntityData.Leafs.Append("mp-state", types.YLeaf{"MpState", mpInfo.MpState})

    mpInfo.EntityData.YListKeys = []string {}

    return &(mpInfo.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray
// Array of MP members
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Member Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
    Interface interface{}

    // Member State. The type is PppLcpMpMbrState.
    State interface{}
}

func (mpMemberInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_MpInfo_MpMemberInfoArray) GetEntityData() *types.CommonEntityData {
    mpMemberInfoArray.EntityData.YFilter = mpMemberInfoArray.YFilter
    mpMemberInfoArray.EntityData.YangName = "mp-member-info-array"
    mpMemberInfoArray.EntityData.BundleName = "cisco_ios_xr"
    mpMemberInfoArray.EntityData.ParentYangName = "mp-info"
    mpMemberInfoArray.EntityData.SegmentPath = "mp-member-info-array"
    mpMemberInfoArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mpMemberInfoArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mpMemberInfoArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mpMemberInfoArray.EntityData.Children = types.NewOrderedMap()
    mpMemberInfoArray.EntityData.Leafs = types.NewOrderedMap()
    mpMemberInfoArray.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", mpMemberInfoArray.Interface})
    mpMemberInfoArray.EntityData.Leafs.Append("state", types.YLeaf{"State", mpMemberInfoArray.State})

    mpMemberInfoArray.EntityData.YListKeys = []string {}

    return &(mpMemberInfoArray.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout
// Configured timeout
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minutes. The type is interface{} with range: 0..4294967295. Units are
    // minute.
    Minutes interface{}

    // Seconds. The type is interface{} with range: 0..255. Units are second.
    Seconds interface{}
}

func (configuredTimeout *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_ConfiguredTimeout) GetEntityData() *types.CommonEntityData {
    configuredTimeout.EntityData.YFilter = configuredTimeout.YFilter
    configuredTimeout.EntityData.YangName = "configured-timeout"
    configuredTimeout.EntityData.BundleName = "cisco_ios_xr"
    configuredTimeout.EntityData.ParentYangName = "node-interface"
    configuredTimeout.EntityData.SegmentPath = "configured-timeout"
    configuredTimeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuredTimeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuredTimeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuredTimeout.EntityData.Children = types.NewOrderedMap()
    configuredTimeout.EntityData.Leafs = types.NewOrderedMap()
    configuredTimeout.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", configuredTimeout.Minutes})
    configuredTimeout.EntityData.Leafs.Append("seconds", types.YLeaf{"Seconds", configuredTimeout.Seconds})

    configuredTimeout.EntityData.YListKeys = []string {}

    return &(configuredTimeout.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo
// Authentication information
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo struct {
    EntityData types.CommonEntityData
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

func (authInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_AuthInfo) GetEntityData() *types.CommonEntityData {
    authInfo.EntityData.YFilter = authInfo.YFilter
    authInfo.EntityData.YangName = "auth-info"
    authInfo.EntityData.BundleName = "cisco_ios_xr"
    authInfo.EntityData.ParentYangName = "node-interface"
    authInfo.EntityData.SegmentPath = "auth-info"
    authInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authInfo.EntityData.Children = types.NewOrderedMap()
    authInfo.EntityData.Leafs = types.NewOrderedMap()
    authInfo.EntityData.Leafs.Append("is-authenticated", types.YLeaf{"IsAuthenticated", authInfo.IsAuthenticated})
    authInfo.EntityData.Leafs.Append("is-sso-authenticated", types.YLeaf{"IsSsoAuthenticated", authInfo.IsSsoAuthenticated})
    authInfo.EntityData.Leafs.Append("of-us-auth", types.YLeaf{"OfUsAuth", authInfo.OfUsAuth})
    authInfo.EntityData.Leafs.Append("of-peer-auth", types.YLeaf{"OfPeerAuth", authInfo.OfPeerAuth})
    authInfo.EntityData.Leafs.Append("of-us-name", types.YLeaf{"OfUsName", authInfo.OfUsName})
    authInfo.EntityData.Leafs.Append("of-peer-name", types.YLeaf{"OfPeerName", authInfo.OfPeerName})
    authInfo.EntityData.Leafs.Append("of-us-sso-state", types.YLeaf{"OfUsSsoState", authInfo.OfUsSsoState})
    authInfo.EntityData.Leafs.Append("of-peer-sso-state", types.YLeaf{"OfPeerSsoState", authInfo.OfPeerSsoState})

    authInfo.EntityData.YListKeys = []string {}

    return &(authInfo.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray
// Array of per-NCP data
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray struct {
    EntityData types.CommonEntityData
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

func (ncpInfoArray *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray) GetEntityData() *types.CommonEntityData {
    ncpInfoArray.EntityData.YFilter = ncpInfoArray.YFilter
    ncpInfoArray.EntityData.YangName = "ncp-info-array"
    ncpInfoArray.EntityData.BundleName = "cisco_ios_xr"
    ncpInfoArray.EntityData.ParentYangName = "node-interface"
    ncpInfoArray.EntityData.SegmentPath = "ncp-info-array"
    ncpInfoArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncpInfoArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncpInfoArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncpInfoArray.EntityData.Children = types.NewOrderedMap()
    ncpInfoArray.EntityData.Children.Append("ncp-info", types.YChild{"NcpInfo", &ncpInfoArray.NcpInfo})
    ncpInfoArray.EntityData.Leafs = types.NewOrderedMap()
    ncpInfoArray.EntityData.Leafs.Append("ncp-state", types.YLeaf{"NcpState", ncpInfoArray.NcpState})
    ncpInfoArray.EntityData.Leafs.Append("ncpsso-state", types.YLeaf{"NcpssoState", ncpInfoArray.NcpssoState})
    ncpInfoArray.EntityData.Leafs.Append("is-passive", types.YLeaf{"IsPassive", ncpInfoArray.IsPassive})
    ncpInfoArray.EntityData.Leafs.Append("ncp-identifier", types.YLeaf{"NcpIdentifier", ncpInfoArray.NcpIdentifier})

    ncpInfoArray.EntityData.YListKeys = []string {}

    return &(ncpInfoArray.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo
// Specific NCP info
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is NcpIdent.
    Type interface{}

    // Info for IPCP.
    IpcpInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo

    // Info for IPCPIW.
    IpcpiwInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo

    // Info for IPv6CP.
    Ipv6cpInfo Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6cpInfo
}

func (ncpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo) GetEntityData() *types.CommonEntityData {
    ncpInfo.EntityData.YFilter = ncpInfo.YFilter
    ncpInfo.EntityData.YangName = "ncp-info"
    ncpInfo.EntityData.BundleName = "cisco_ios_xr"
    ncpInfo.EntityData.ParentYangName = "ncp-info-array"
    ncpInfo.EntityData.SegmentPath = "ncp-info"
    ncpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncpInfo.EntityData.Children = types.NewOrderedMap()
    ncpInfo.EntityData.Children.Append("ipcp-info", types.YChild{"IpcpInfo", &ncpInfo.IpcpInfo})
    ncpInfo.EntityData.Children.Append("ipcpiw-info", types.YChild{"IpcpiwInfo", &ncpInfo.IpcpiwInfo})
    ncpInfo.EntityData.Children.Append("ipv6cp-info", types.YChild{"Ipv6cpInfo", &ncpInfo.Ipv6cpInfo})
    ncpInfo.EntityData.Leafs = types.NewOrderedMap()
    ncpInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", ncpInfo.Type})

    ncpInfo.EntityData.YListKeys = []string {}

    return &(ncpInfo.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo
// Info for IPCP
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo struct {
    EntityData types.CommonEntityData
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

func (ipcpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo) GetEntityData() *types.CommonEntityData {
    ipcpInfo.EntityData.YFilter = ipcpInfo.YFilter
    ipcpInfo.EntityData.YangName = "ipcp-info"
    ipcpInfo.EntityData.BundleName = "cisco_ios_xr"
    ipcpInfo.EntityData.ParentYangName = "ncp-info"
    ipcpInfo.EntityData.SegmentPath = "ipcp-info"
    ipcpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcpInfo.EntityData.Children = types.NewOrderedMap()
    ipcpInfo.EntityData.Children.Append("local-iphc-options", types.YChild{"LocalIphcOptions", &ipcpInfo.LocalIphcOptions})
    ipcpInfo.EntityData.Children.Append("peer-iphc-options", types.YChild{"PeerIphcOptions", &ipcpInfo.PeerIphcOptions})
    ipcpInfo.EntityData.Leafs = types.NewOrderedMap()
    ipcpInfo.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", ipcpInfo.LocalAddress})
    ipcpInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", ipcpInfo.PeerAddress})
    ipcpInfo.EntityData.Leafs.Append("peer-netmask", types.YLeaf{"PeerNetmask", ipcpInfo.PeerNetmask})
    ipcpInfo.EntityData.Leafs.Append("dns-primary", types.YLeaf{"DnsPrimary", ipcpInfo.DnsPrimary})
    ipcpInfo.EntityData.Leafs.Append("dns-secondary", types.YLeaf{"DnsSecondary", ipcpInfo.DnsSecondary})
    ipcpInfo.EntityData.Leafs.Append("wins-primary", types.YLeaf{"WinsPrimary", ipcpInfo.WinsPrimary})
    ipcpInfo.EntityData.Leafs.Append("wins-secondary", types.YLeaf{"WinsSecondary", ipcpInfo.WinsSecondary})
    ipcpInfo.EntityData.Leafs.Append("is-iphc-configured", types.YLeaf{"IsIphcConfigured", ipcpInfo.IsIphcConfigured})

    ipcpInfo.EntityData.YListKeys = []string {}

    return &(ipcpInfo.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions
// Local IPHC options
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions struct {
    EntityData types.CommonEntityData
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

func (localIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_LocalIphcOptions) GetEntityData() *types.CommonEntityData {
    localIphcOptions.EntityData.YFilter = localIphcOptions.YFilter
    localIphcOptions.EntityData.YangName = "local-iphc-options"
    localIphcOptions.EntityData.BundleName = "cisco_ios_xr"
    localIphcOptions.EntityData.ParentYangName = "ipcp-info"
    localIphcOptions.EntityData.SegmentPath = "local-iphc-options"
    localIphcOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIphcOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIphcOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIphcOptions.EntityData.Children = types.NewOrderedMap()
    localIphcOptions.EntityData.Leafs = types.NewOrderedMap()
    localIphcOptions.EntityData.Leafs.Append("compression-type", types.YLeaf{"CompressionType", localIphcOptions.CompressionType})
    localIphcOptions.EntityData.Leafs.Append("tcp-space", types.YLeaf{"TcpSpace", localIphcOptions.TcpSpace})
    localIphcOptions.EntityData.Leafs.Append("non-tcp-space", types.YLeaf{"NonTcpSpace", localIphcOptions.NonTcpSpace})
    localIphcOptions.EntityData.Leafs.Append("max-period", types.YLeaf{"MaxPeriod", localIphcOptions.MaxPeriod})
    localIphcOptions.EntityData.Leafs.Append("max-time", types.YLeaf{"MaxTime", localIphcOptions.MaxTime})
    localIphcOptions.EntityData.Leafs.Append("max-header", types.YLeaf{"MaxHeader", localIphcOptions.MaxHeader})
    localIphcOptions.EntityData.Leafs.Append("rtp-compression", types.YLeaf{"RtpCompression", localIphcOptions.RtpCompression})
    localIphcOptions.EntityData.Leafs.Append("ec-rtp-compression", types.YLeaf{"EcRtpCompression", localIphcOptions.EcRtpCompression})

    localIphcOptions.EntityData.YListKeys = []string {}

    return &(localIphcOptions.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions
// Peer IPHC options
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions struct {
    EntityData types.CommonEntityData
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

func (peerIphcOptions *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpInfo_PeerIphcOptions) GetEntityData() *types.CommonEntityData {
    peerIphcOptions.EntityData.YFilter = peerIphcOptions.YFilter
    peerIphcOptions.EntityData.YangName = "peer-iphc-options"
    peerIphcOptions.EntityData.BundleName = "cisco_ios_xr"
    peerIphcOptions.EntityData.ParentYangName = "ipcp-info"
    peerIphcOptions.EntityData.SegmentPath = "peer-iphc-options"
    peerIphcOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerIphcOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerIphcOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerIphcOptions.EntityData.Children = types.NewOrderedMap()
    peerIphcOptions.EntityData.Leafs = types.NewOrderedMap()
    peerIphcOptions.EntityData.Leafs.Append("compression-type", types.YLeaf{"CompressionType", peerIphcOptions.CompressionType})
    peerIphcOptions.EntityData.Leafs.Append("tcp-space", types.YLeaf{"TcpSpace", peerIphcOptions.TcpSpace})
    peerIphcOptions.EntityData.Leafs.Append("non-tcp-space", types.YLeaf{"NonTcpSpace", peerIphcOptions.NonTcpSpace})
    peerIphcOptions.EntityData.Leafs.Append("max-period", types.YLeaf{"MaxPeriod", peerIphcOptions.MaxPeriod})
    peerIphcOptions.EntityData.Leafs.Append("max-time", types.YLeaf{"MaxTime", peerIphcOptions.MaxTime})
    peerIphcOptions.EntityData.Leafs.Append("max-header", types.YLeaf{"MaxHeader", peerIphcOptions.MaxHeader})
    peerIphcOptions.EntityData.Leafs.Append("rtp-compression", types.YLeaf{"RtpCompression", peerIphcOptions.RtpCompression})
    peerIphcOptions.EntityData.Leafs.Append("ec-rtp-compression", types.YLeaf{"EcRtpCompression", peerIphcOptions.EcRtpCompression})

    peerIphcOptions.EntityData.YListKeys = []string {}

    return &(peerIphcOptions.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo
// Info for IPCPIW
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Peer IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}
}

func (ipcpiwInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_IpcpiwInfo) GetEntityData() *types.CommonEntityData {
    ipcpiwInfo.EntityData.YFilter = ipcpiwInfo.YFilter
    ipcpiwInfo.EntityData.YangName = "ipcpiw-info"
    ipcpiwInfo.EntityData.BundleName = "cisco_ios_xr"
    ipcpiwInfo.EntityData.ParentYangName = "ncp-info"
    ipcpiwInfo.EntityData.SegmentPath = "ipcpiw-info"
    ipcpiwInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcpiwInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcpiwInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcpiwInfo.EntityData.Children = types.NewOrderedMap()
    ipcpiwInfo.EntityData.Leafs = types.NewOrderedMap()
    ipcpiwInfo.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", ipcpiwInfo.LocalAddress})
    ipcpiwInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", ipcpiwInfo.PeerAddress})

    ipcpiwInfo.EntityData.YListKeys = []string {}

    return &(ipcpiwInfo.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6cpInfo
// Info for IPv6CP
type Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6cpInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalAddress interface{}

    // Peer IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}
}

func (ipv6cpInfo *Ppp_Nodes_Node_NodeInterfaces_NodeInterface_NcpInfoArray_NcpInfo_Ipv6cpInfo) GetEntityData() *types.CommonEntityData {
    ipv6cpInfo.EntityData.YFilter = ipv6cpInfo.YFilter
    ipv6cpInfo.EntityData.YangName = "ipv6cp-info"
    ipv6cpInfo.EntityData.BundleName = "cisco_ios_xr"
    ipv6cpInfo.EntityData.ParentYangName = "ncp-info"
    ipv6cpInfo.EntityData.SegmentPath = "ipv6cp-info"
    ipv6cpInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6cpInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6cpInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6cpInfo.EntityData.Children = types.NewOrderedMap()
    ipv6cpInfo.EntityData.Leafs = types.NewOrderedMap()
    ipv6cpInfo.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", ipv6cpInfo.LocalAddress})
    ipv6cpInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", ipv6cpInfo.PeerAddress})

    ipv6cpInfo.EntityData.YListKeys = []string {}

    return &(ipv6cpInfo.EntityData)
}

// Ppp_Nodes_Node_SsoAlerts
// PPP SSO Alert data for a particular node
type Ppp_Nodes_Node_SsoAlerts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPP SSO Alert data for a particular interface. The type is slice of
    // Ppp_Nodes_Node_SsoAlerts_SsoAlert.
    SsoAlert []*Ppp_Nodes_Node_SsoAlerts_SsoAlert
}

func (ssoAlerts *Ppp_Nodes_Node_SsoAlerts) GetEntityData() *types.CommonEntityData {
    ssoAlerts.EntityData.YFilter = ssoAlerts.YFilter
    ssoAlerts.EntityData.YangName = "sso-alerts"
    ssoAlerts.EntityData.BundleName = "cisco_ios_xr"
    ssoAlerts.EntityData.ParentYangName = "node"
    ssoAlerts.EntityData.SegmentPath = "sso-alerts"
    ssoAlerts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoAlerts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoAlerts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoAlerts.EntityData.Children = types.NewOrderedMap()
    ssoAlerts.EntityData.Children.Append("sso-alert", types.YChild{"SsoAlert", nil})
    for i := range ssoAlerts.SsoAlert {
        ssoAlerts.EntityData.Children.Append(types.GetSegmentPath(ssoAlerts.SsoAlert[i]), types.YChild{"SsoAlert", ssoAlerts.SsoAlert[i]})
    }
    ssoAlerts.EntityData.Leafs = types.NewOrderedMap()

    ssoAlerts.EntityData.YListKeys = []string {}

    return &(ssoAlerts.EntityData)
}

// Ppp_Nodes_Node_SsoAlerts_SsoAlert
// PPP SSO Alert data for a particular interface
type Ppp_Nodes_Node_SsoAlerts_SsoAlert struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface with SSO Alert. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
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

func (ssoAlert *Ppp_Nodes_Node_SsoAlerts_SsoAlert) GetEntityData() *types.CommonEntityData {
    ssoAlert.EntityData.YFilter = ssoAlert.YFilter
    ssoAlert.EntityData.YangName = "sso-alert"
    ssoAlert.EntityData.BundleName = "cisco_ios_xr"
    ssoAlert.EntityData.ParentYangName = "sso-alerts"
    ssoAlert.EntityData.SegmentPath = "sso-alert" + types.AddKeyToken(ssoAlert.Interface, "interface")
    ssoAlert.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoAlert.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoAlert.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoAlert.EntityData.Children = types.NewOrderedMap()
    ssoAlert.EntityData.Children.Append("lcp-error", types.YChild{"LcpError", &ssoAlert.LcpError})
    ssoAlert.EntityData.Children.Append("of-us-auth-error", types.YChild{"OfUsAuthError", &ssoAlert.OfUsAuthError})
    ssoAlert.EntityData.Children.Append("of-peer-auth-error", types.YChild{"OfPeerAuthError", &ssoAlert.OfPeerAuthError})
    ssoAlert.EntityData.Children.Append("ipcp-error", types.YChild{"IpcpError", &ssoAlert.IpcpError})
    ssoAlert.EntityData.Leafs = types.NewOrderedMap()
    ssoAlert.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", ssoAlert.Interface})

    ssoAlert.EntityData.YListKeys = []string {"Interface"}

    return &(ssoAlert.EntityData)
}

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError
// LCP SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (lcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_LcpError) GetEntityData() *types.CommonEntityData {
    lcpError.EntityData.YFilter = lcpError.YFilter
    lcpError.EntityData.YangName = "lcp-error"
    lcpError.EntityData.BundleName = "cisco_ios_xr"
    lcpError.EntityData.ParentYangName = "sso-alert"
    lcpError.EntityData.SegmentPath = "lcp-error"
    lcpError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpError.EntityData.Children = types.NewOrderedMap()
    lcpError.EntityData.Leafs = types.NewOrderedMap()
    lcpError.EntityData.Leafs.Append("is-error", types.YLeaf{"IsError", lcpError.IsError})
    lcpError.EntityData.Leafs.Append("error", types.YLeaf{"Error", lcpError.Error})
    lcpError.EntityData.Leafs.Append("context", types.YLeaf{"Context", lcpError.Context})

    lcpError.EntityData.YListKeys = []string {}

    return &(lcpError.EntityData)
}

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError
// Of-us Authentication SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (ofUsAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfUsAuthError) GetEntityData() *types.CommonEntityData {
    ofUsAuthError.EntityData.YFilter = ofUsAuthError.YFilter
    ofUsAuthError.EntityData.YangName = "of-us-auth-error"
    ofUsAuthError.EntityData.BundleName = "cisco_ios_xr"
    ofUsAuthError.EntityData.ParentYangName = "sso-alert"
    ofUsAuthError.EntityData.SegmentPath = "of-us-auth-error"
    ofUsAuthError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ofUsAuthError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ofUsAuthError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ofUsAuthError.EntityData.Children = types.NewOrderedMap()
    ofUsAuthError.EntityData.Leafs = types.NewOrderedMap()
    ofUsAuthError.EntityData.Leafs.Append("is-error", types.YLeaf{"IsError", ofUsAuthError.IsError})
    ofUsAuthError.EntityData.Leafs.Append("error", types.YLeaf{"Error", ofUsAuthError.Error})
    ofUsAuthError.EntityData.Leafs.Append("context", types.YLeaf{"Context", ofUsAuthError.Context})

    ofUsAuthError.EntityData.YListKeys = []string {}

    return &(ofUsAuthError.EntityData)
}

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError
// Of-peer Authentication SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (ofPeerAuthError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_OfPeerAuthError) GetEntityData() *types.CommonEntityData {
    ofPeerAuthError.EntityData.YFilter = ofPeerAuthError.YFilter
    ofPeerAuthError.EntityData.YangName = "of-peer-auth-error"
    ofPeerAuthError.EntityData.BundleName = "cisco_ios_xr"
    ofPeerAuthError.EntityData.ParentYangName = "sso-alert"
    ofPeerAuthError.EntityData.SegmentPath = "of-peer-auth-error"
    ofPeerAuthError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ofPeerAuthError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ofPeerAuthError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ofPeerAuthError.EntityData.Children = types.NewOrderedMap()
    ofPeerAuthError.EntityData.Leafs = types.NewOrderedMap()
    ofPeerAuthError.EntityData.Leafs.Append("is-error", types.YLeaf{"IsError", ofPeerAuthError.IsError})
    ofPeerAuthError.EntityData.Leafs.Append("error", types.YLeaf{"Error", ofPeerAuthError.Error})
    ofPeerAuthError.EntityData.Leafs.Append("context", types.YLeaf{"Context", ofPeerAuthError.Context})

    ofPeerAuthError.EntityData.YListKeys = []string {}

    return &(ofPeerAuthError.EntityData)
}

// Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError
// IPCP SSO Error
type Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO Error. The type is bool.
    IsError interface{}

    // SSO Error. The type is interface{} with range: 0..4294967295.
    Error interface{}

    // Context. The type is interface{} with range: 0..4294967295.
    Context interface{}
}

func (ipcpError *Ppp_Nodes_Node_SsoAlerts_SsoAlert_IpcpError) GetEntityData() *types.CommonEntityData {
    ipcpError.EntityData.YFilter = ipcpError.YFilter
    ipcpError.EntityData.YangName = "ipcp-error"
    ipcpError.EntityData.BundleName = "cisco_ios_xr"
    ipcpError.EntityData.ParentYangName = "sso-alert"
    ipcpError.EntityData.SegmentPath = "ipcp-error"
    ipcpError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcpError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcpError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcpError.EntityData.Children = types.NewOrderedMap()
    ipcpError.EntityData.Leafs = types.NewOrderedMap()
    ipcpError.EntityData.Leafs.Append("is-error", types.YLeaf{"IsError", ipcpError.IsError})
    ipcpError.EntityData.Leafs.Append("error", types.YLeaf{"Error", ipcpError.Error})
    ipcpError.EntityData.Leafs.Append("context", types.YLeaf{"Context", ipcpError.Context})

    ipcpError.EntityData.YListKeys = []string {}

    return &(ipcpError.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaceStatistics
// Per interface PPP operational statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LCP and NCP statistics for an interface running PPP. The type is slice of
    // Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic.
    NodeInterfaceStatistic []*Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic
}

func (nodeInterfaceStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics) GetEntityData() *types.CommonEntityData {
    nodeInterfaceStatistics.EntityData.YFilter = nodeInterfaceStatistics.YFilter
    nodeInterfaceStatistics.EntityData.YangName = "node-interface-statistics"
    nodeInterfaceStatistics.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceStatistics.EntityData.ParentYangName = "node"
    nodeInterfaceStatistics.EntityData.SegmentPath = "node-interface-statistics"
    nodeInterfaceStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceStatistics.EntityData.Children = types.NewOrderedMap()
    nodeInterfaceStatistics.EntityData.Children.Append("node-interface-statistic", types.YChild{"NodeInterfaceStatistic", nil})
    for i := range nodeInterfaceStatistics.NodeInterfaceStatistic {
        nodeInterfaceStatistics.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaceStatistics.NodeInterfaceStatistic[i]), types.YChild{"NodeInterfaceStatistic", nodeInterfaceStatistics.NodeInterfaceStatistic[i]})
    }
    nodeInterfaceStatistics.EntityData.Leafs = types.NewOrderedMap()

    nodeInterfaceStatistics.EntityData.YListKeys = []string {}

    return &(nodeInterfaceStatistics.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic
// LCP and NCP statistics for an interface
// running PPP
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface running PPP. The type is string with
    // pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // PPP LCP Statistics.
    LcpStatistics Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics

    // PPP Authentication statistics.
    AuthenticationStatistics Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics

    // Array of PPP NCP Statistics. The type is slice of
    // Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray.
    NcpStatisticsArray []*Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray
}

func (nodeInterfaceStatistic *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic) GetEntityData() *types.CommonEntityData {
    nodeInterfaceStatistic.EntityData.YFilter = nodeInterfaceStatistic.YFilter
    nodeInterfaceStatistic.EntityData.YangName = "node-interface-statistic"
    nodeInterfaceStatistic.EntityData.BundleName = "cisco_ios_xr"
    nodeInterfaceStatistic.EntityData.ParentYangName = "node-interface-statistics"
    nodeInterfaceStatistic.EntityData.SegmentPath = "node-interface-statistic" + types.AddKeyToken(nodeInterfaceStatistic.InterfaceName, "interface-name")
    nodeInterfaceStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeInterfaceStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeInterfaceStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeInterfaceStatistic.EntityData.Children = types.NewOrderedMap()
    nodeInterfaceStatistic.EntityData.Children.Append("lcp-statistics", types.YChild{"LcpStatistics", &nodeInterfaceStatistic.LcpStatistics})
    nodeInterfaceStatistic.EntityData.Children.Append("authentication-statistics", types.YChild{"AuthenticationStatistics", &nodeInterfaceStatistic.AuthenticationStatistics})
    nodeInterfaceStatistic.EntityData.Children.Append("ncp-statistics-array", types.YChild{"NcpStatisticsArray", nil})
    for i := range nodeInterfaceStatistic.NcpStatisticsArray {
        nodeInterfaceStatistic.EntityData.Children.Append(types.GetSegmentPath(nodeInterfaceStatistic.NcpStatisticsArray[i]), types.YChild{"NcpStatisticsArray", nodeInterfaceStatistic.NcpStatisticsArray[i]})
    }
    nodeInterfaceStatistic.EntityData.Leafs = types.NewOrderedMap()
    nodeInterfaceStatistic.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", nodeInterfaceStatistic.InterfaceName})

    nodeInterfaceStatistic.EntityData.YListKeys = []string {"InterfaceName"}

    return &(nodeInterfaceStatistic.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics
// PPP LCP Statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics struct {
    EntityData types.CommonEntityData
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

func (lcpStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_LcpStatistics) GetEntityData() *types.CommonEntityData {
    lcpStatistics.EntityData.YFilter = lcpStatistics.YFilter
    lcpStatistics.EntityData.YangName = "lcp-statistics"
    lcpStatistics.EntityData.BundleName = "cisco_ios_xr"
    lcpStatistics.EntityData.ParentYangName = "node-interface-statistic"
    lcpStatistics.EntityData.SegmentPath = "lcp-statistics"
    lcpStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpStatistics.EntityData.Children = types.NewOrderedMap()
    lcpStatistics.EntityData.Leafs = types.NewOrderedMap()
    lcpStatistics.EntityData.Leafs.Append("conf-req-sent", types.YLeaf{"ConfReqSent", lcpStatistics.ConfReqSent})
    lcpStatistics.EntityData.Leafs.Append("conf-req-rcvd", types.YLeaf{"ConfReqRcvd", lcpStatistics.ConfReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("conf-ack-sent", types.YLeaf{"ConfAckSent", lcpStatistics.ConfAckSent})
    lcpStatistics.EntityData.Leafs.Append("conf-ack-rcvd", types.YLeaf{"ConfAckRcvd", lcpStatistics.ConfAckRcvd})
    lcpStatistics.EntityData.Leafs.Append("conf-nak-sent", types.YLeaf{"ConfNakSent", lcpStatistics.ConfNakSent})
    lcpStatistics.EntityData.Leafs.Append("conf-nak-rcvd", types.YLeaf{"ConfNakRcvd", lcpStatistics.ConfNakRcvd})
    lcpStatistics.EntityData.Leafs.Append("conf-rej-sent", types.YLeaf{"ConfRejSent", lcpStatistics.ConfRejSent})
    lcpStatistics.EntityData.Leafs.Append("conf-rej-rcvd", types.YLeaf{"ConfRejRcvd", lcpStatistics.ConfRejRcvd})
    lcpStatistics.EntityData.Leafs.Append("echo-req-sent", types.YLeaf{"EchoReqSent", lcpStatistics.EchoReqSent})
    lcpStatistics.EntityData.Leafs.Append("echo-req-rcvd", types.YLeaf{"EchoReqRcvd", lcpStatistics.EchoReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("echo-rep-sent", types.YLeaf{"EchoRepSent", lcpStatistics.EchoRepSent})
    lcpStatistics.EntityData.Leafs.Append("echo-rep-rcvd", types.YLeaf{"EchoRepRcvd", lcpStatistics.EchoRepRcvd})
    lcpStatistics.EntityData.Leafs.Append("disc-req-sent", types.YLeaf{"DiscReqSent", lcpStatistics.DiscReqSent})
    lcpStatistics.EntityData.Leafs.Append("disc-req-rcvd", types.YLeaf{"DiscReqRcvd", lcpStatistics.DiscReqRcvd})
    lcpStatistics.EntityData.Leafs.Append("link-up", types.YLeaf{"LinkUp", lcpStatistics.LinkUp})
    lcpStatistics.EntityData.Leafs.Append("link-error", types.YLeaf{"LinkError", lcpStatistics.LinkError})

    lcpStatistics.EntityData.YListKeys = []string {}

    return &(lcpStatistics.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics
// PPP Authentication statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics struct {
    EntityData types.CommonEntityData
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

func (authenticationStatistics *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_AuthenticationStatistics) GetEntityData() *types.CommonEntityData {
    authenticationStatistics.EntityData.YFilter = authenticationStatistics.YFilter
    authenticationStatistics.EntityData.YangName = "authentication-statistics"
    authenticationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authenticationStatistics.EntityData.ParentYangName = "node-interface-statistic"
    authenticationStatistics.EntityData.SegmentPath = "authentication-statistics"
    authenticationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationStatistics.EntityData.Children = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs = types.NewOrderedMap()
    authenticationStatistics.EntityData.Leafs.Append("pap-req-sent", types.YLeaf{"PapReqSent", authenticationStatistics.PapReqSent})
    authenticationStatistics.EntityData.Leafs.Append("pap-req-rcvd", types.YLeaf{"PapReqRcvd", authenticationStatistics.PapReqRcvd})
    authenticationStatistics.EntityData.Leafs.Append("pap-ack-sent", types.YLeaf{"PapAckSent", authenticationStatistics.PapAckSent})
    authenticationStatistics.EntityData.Leafs.Append("pap-ack-rcvd", types.YLeaf{"PapAckRcvd", authenticationStatistics.PapAckRcvd})
    authenticationStatistics.EntityData.Leafs.Append("pap-nak-sent", types.YLeaf{"PapNakSent", authenticationStatistics.PapNakSent})
    authenticationStatistics.EntityData.Leafs.Append("pap-nak-rcvd", types.YLeaf{"PapNakRcvd", authenticationStatistics.PapNakRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-chall-sent", types.YLeaf{"ChapChallSent", authenticationStatistics.ChapChallSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-chall-rcvd", types.YLeaf{"ChapChallRcvd", authenticationStatistics.ChapChallRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-resp-sent", types.YLeaf{"ChapRespSent", authenticationStatistics.ChapRespSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-resp-rcvd", types.YLeaf{"ChapRespRcvd", authenticationStatistics.ChapRespRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-succ-sent", types.YLeaf{"ChapRepSuccSent", authenticationStatistics.ChapRepSuccSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-succ-rcvd", types.YLeaf{"ChapRepSuccRcvd", authenticationStatistics.ChapRepSuccRcvd})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-fail-sent", types.YLeaf{"ChapRepFailSent", authenticationStatistics.ChapRepFailSent})
    authenticationStatistics.EntityData.Leafs.Append("chap-rep-fail-rcvd", types.YLeaf{"ChapRepFailRcvd", authenticationStatistics.ChapRepFailRcvd})
    authenticationStatistics.EntityData.Leafs.Append("auth-timeout-count", types.YLeaf{"AuthTimeoutCount", authenticationStatistics.AuthTimeoutCount})

    authenticationStatistics.EntityData.YListKeys = []string {}

    return &(authenticationStatistics.EntityData)
}

// Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray
// Array of PPP NCP Statistics
type Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray struct {
    EntityData types.CommonEntityData
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

func (ncpStatisticsArray *Ppp_Nodes_Node_NodeInterfaceStatistics_NodeInterfaceStatistic_NcpStatisticsArray) GetEntityData() *types.CommonEntityData {
    ncpStatisticsArray.EntityData.YFilter = ncpStatisticsArray.YFilter
    ncpStatisticsArray.EntityData.YangName = "ncp-statistics-array"
    ncpStatisticsArray.EntityData.BundleName = "cisco_ios_xr"
    ncpStatisticsArray.EntityData.ParentYangName = "node-interface-statistic"
    ncpStatisticsArray.EntityData.SegmentPath = "ncp-statistics-array"
    ncpStatisticsArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncpStatisticsArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncpStatisticsArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncpStatisticsArray.EntityData.Children = types.NewOrderedMap()
    ncpStatisticsArray.EntityData.Leafs = types.NewOrderedMap()
    ncpStatisticsArray.EntityData.Leafs.Append("ncp-identifier", types.YLeaf{"NcpIdentifier", ncpStatisticsArray.NcpIdentifier})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-req-sent", types.YLeaf{"ConfReqSent", ncpStatisticsArray.ConfReqSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-req-rcvd", types.YLeaf{"ConfReqRcvd", ncpStatisticsArray.ConfReqRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-ack-sent", types.YLeaf{"ConfAckSent", ncpStatisticsArray.ConfAckSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-ack-rcvd", types.YLeaf{"ConfAckRcvd", ncpStatisticsArray.ConfAckRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-nak-sent", types.YLeaf{"ConfNakSent", ncpStatisticsArray.ConfNakSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-nak-rcvd", types.YLeaf{"ConfNakRcvd", ncpStatisticsArray.ConfNakRcvd})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-rej-sent", types.YLeaf{"ConfRejSent", ncpStatisticsArray.ConfRejSent})
    ncpStatisticsArray.EntityData.Leafs.Append("conf-rej-rcvd", types.YLeaf{"ConfRejRcvd", ncpStatisticsArray.ConfRejRcvd})

    ncpStatisticsArray.EntityData.YListKeys = []string {}

    return &(ncpStatisticsArray.EntityData)
}

// Ppp_Nodes_Node_SsoSummary
// Summarized PPP SSO data for a particular node
type Ppp_Nodes_Node_SsoSummary struct {
    EntityData types.CommonEntityData
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

func (ssoSummary *Ppp_Nodes_Node_SsoSummary) GetEntityData() *types.CommonEntityData {
    ssoSummary.EntityData.YFilter = ssoSummary.YFilter
    ssoSummary.EntityData.YangName = "sso-summary"
    ssoSummary.EntityData.BundleName = "cisco_ios_xr"
    ssoSummary.EntityData.ParentYangName = "node"
    ssoSummary.EntityData.SegmentPath = "sso-summary"
    ssoSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoSummary.EntityData.Children = types.NewOrderedMap()
    ssoSummary.EntityData.Children.Append("lcp-states", types.YChild{"LcpStates", &ssoSummary.LcpStates})
    ssoSummary.EntityData.Children.Append("of-us-auth-states", types.YChild{"OfUsAuthStates", &ssoSummary.OfUsAuthStates})
    ssoSummary.EntityData.Children.Append("of-peer-auth-states", types.YChild{"OfPeerAuthStates", &ssoSummary.OfPeerAuthStates})
    ssoSummary.EntityData.Children.Append("ipcp-states", types.YChild{"IpcpStates", &ssoSummary.IpcpStates})
    ssoSummary.EntityData.Leafs = types.NewOrderedMap()

    ssoSummary.EntityData.YListKeys = []string {}

    return &(ssoSummary.EntityData)
}

// Ppp_Nodes_Node_SsoSummary_LcpStates
// LCP SSO FSM States
type Ppp_Nodes_Node_SsoSummary_LcpStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (lcpStates *Ppp_Nodes_Node_SsoSummary_LcpStates) GetEntityData() *types.CommonEntityData {
    lcpStates.EntityData.YFilter = lcpStates.YFilter
    lcpStates.EntityData.YangName = "lcp-states"
    lcpStates.EntityData.BundleName = "cisco_ios_xr"
    lcpStates.EntityData.ParentYangName = "sso-summary"
    lcpStates.EntityData.SegmentPath = "lcp-states"
    lcpStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpStates.EntityData.Children = types.NewOrderedMap()
    lcpStates.EntityData.Leafs = types.NewOrderedMap()
    lcpStates.EntityData.Leafs.Append("total", types.YLeaf{"Total", lcpStates.Total})
    lcpStates.EntityData.Leafs.Append("count", types.YLeaf{"Count", lcpStates.Count})

    lcpStates.EntityData.YListKeys = []string {}

    return &(lcpStates.EntityData)
}

// Ppp_Nodes_Node_SsoSummary_OfUsAuthStates
// Of-us Authentication SSO FSM States
type Ppp_Nodes_Node_SsoSummary_OfUsAuthStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (ofUsAuthStates *Ppp_Nodes_Node_SsoSummary_OfUsAuthStates) GetEntityData() *types.CommonEntityData {
    ofUsAuthStates.EntityData.YFilter = ofUsAuthStates.YFilter
    ofUsAuthStates.EntityData.YangName = "of-us-auth-states"
    ofUsAuthStates.EntityData.BundleName = "cisco_ios_xr"
    ofUsAuthStates.EntityData.ParentYangName = "sso-summary"
    ofUsAuthStates.EntityData.SegmentPath = "of-us-auth-states"
    ofUsAuthStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ofUsAuthStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ofUsAuthStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ofUsAuthStates.EntityData.Children = types.NewOrderedMap()
    ofUsAuthStates.EntityData.Leafs = types.NewOrderedMap()
    ofUsAuthStates.EntityData.Leafs.Append("total", types.YLeaf{"Total", ofUsAuthStates.Total})
    ofUsAuthStates.EntityData.Leafs.Append("count", types.YLeaf{"Count", ofUsAuthStates.Count})

    ofUsAuthStates.EntityData.YListKeys = []string {}

    return &(ofUsAuthStates.EntityData)
}

// Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates
// Of-peer Authentication SSO FSM States
type Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (ofPeerAuthStates *Ppp_Nodes_Node_SsoSummary_OfPeerAuthStates) GetEntityData() *types.CommonEntityData {
    ofPeerAuthStates.EntityData.YFilter = ofPeerAuthStates.YFilter
    ofPeerAuthStates.EntityData.YangName = "of-peer-auth-states"
    ofPeerAuthStates.EntityData.BundleName = "cisco_ios_xr"
    ofPeerAuthStates.EntityData.ParentYangName = "sso-summary"
    ofPeerAuthStates.EntityData.SegmentPath = "of-peer-auth-states"
    ofPeerAuthStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ofPeerAuthStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ofPeerAuthStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ofPeerAuthStates.EntityData.Children = types.NewOrderedMap()
    ofPeerAuthStates.EntityData.Leafs = types.NewOrderedMap()
    ofPeerAuthStates.EntityData.Leafs.Append("total", types.YLeaf{"Total", ofPeerAuthStates.Total})
    ofPeerAuthStates.EntityData.Leafs.Append("count", types.YLeaf{"Count", ofPeerAuthStates.Count})

    ofPeerAuthStates.EntityData.YListKeys = []string {}

    return &(ofPeerAuthStates.EntityData)
}

// Ppp_Nodes_Node_SsoSummary_IpcpStates
// IPCP SSO FSM States
type Ppp_Nodes_Node_SsoSummary_IpcpStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of SSO FSMs running. The type is interface{} with range:
    // 0..65535.
    Total interface{}

    // Number of SSO FSMs in each State. The type is slice of interface{} with
    // range: 0..65535.
    Count []interface{}
}

func (ipcpStates *Ppp_Nodes_Node_SsoSummary_IpcpStates) GetEntityData() *types.CommonEntityData {
    ipcpStates.EntityData.YFilter = ipcpStates.YFilter
    ipcpStates.EntityData.YangName = "ipcp-states"
    ipcpStates.EntityData.BundleName = "cisco_ios_xr"
    ipcpStates.EntityData.ParentYangName = "sso-summary"
    ipcpStates.EntityData.SegmentPath = "ipcp-states"
    ipcpStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcpStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcpStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcpStates.EntityData.Children = types.NewOrderedMap()
    ipcpStates.EntityData.Leafs = types.NewOrderedMap()
    ipcpStates.EntityData.Leafs.Append("total", types.YLeaf{"Total", ipcpStates.Total})
    ipcpStates.EntityData.Leafs.Append("count", types.YLeaf{"Count", ipcpStates.Count})

    ipcpStates.EntityData.YListKeys = []string {}

    return &(ipcpStates.EntityData)
}

// Ppp_Nodes_Node_SsoGroups
// PPP SSO Group data for a particular node
type Ppp_Nodes_Node_SsoGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPP SSO state data for a particular group. The type is slice of
    // Ppp_Nodes_Node_SsoGroups_SsoGroup.
    SsoGroup []*Ppp_Nodes_Node_SsoGroups_SsoGroup
}

func (ssoGroups *Ppp_Nodes_Node_SsoGroups) GetEntityData() *types.CommonEntityData {
    ssoGroups.EntityData.YFilter = ssoGroups.YFilter
    ssoGroups.EntityData.YangName = "sso-groups"
    ssoGroups.EntityData.BundleName = "cisco_ios_xr"
    ssoGroups.EntityData.ParentYangName = "node"
    ssoGroups.EntityData.SegmentPath = "sso-groups"
    ssoGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoGroups.EntityData.Children = types.NewOrderedMap()
    ssoGroups.EntityData.Children.Append("sso-group", types.YChild{"SsoGroup", nil})
    for i := range ssoGroups.SsoGroup {
        ssoGroups.EntityData.Children.Append(types.GetSegmentPath(ssoGroups.SsoGroup[i]), types.YChild{"SsoGroup", ssoGroups.SsoGroup[i]})
    }
    ssoGroups.EntityData.Leafs = types.NewOrderedMap()

    ssoGroups.EntityData.YListKeys = []string {}

    return &(ssoGroups.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup
// PPP SSO state data for a particular group
type Ppp_Nodes_Node_SsoGroups_SsoGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The identifier for the group. The type is
    // interface{} with range: 1..65535.
    GroupId interface{}

    // PPP SSO State data for a particular group.
    SsoStates Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates
}

func (ssoGroup *Ppp_Nodes_Node_SsoGroups_SsoGroup) GetEntityData() *types.CommonEntityData {
    ssoGroup.EntityData.YFilter = ssoGroup.YFilter
    ssoGroup.EntityData.YangName = "sso-group"
    ssoGroup.EntityData.BundleName = "cisco_ios_xr"
    ssoGroup.EntityData.ParentYangName = "sso-groups"
    ssoGroup.EntityData.SegmentPath = "sso-group" + types.AddKeyToken(ssoGroup.GroupId, "group-id")
    ssoGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoGroup.EntityData.Children = types.NewOrderedMap()
    ssoGroup.EntityData.Children.Append("sso-states", types.YChild{"SsoStates", &ssoGroup.SsoStates})
    ssoGroup.EntityData.Leafs = types.NewOrderedMap()
    ssoGroup.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", ssoGroup.GroupId})

    ssoGroup.EntityData.YListKeys = []string {"GroupId"}

    return &(ssoGroup.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates
// PPP SSO State data for a particular group
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPP SSO State data for a particular interface. The type is slice of
    // Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState.
    SsoState []*Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState
}

func (ssoStates *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates) GetEntityData() *types.CommonEntityData {
    ssoStates.EntityData.YFilter = ssoStates.YFilter
    ssoStates.EntityData.YangName = "sso-states"
    ssoStates.EntityData.BundleName = "cisco_ios_xr"
    ssoStates.EntityData.ParentYangName = "sso-group"
    ssoStates.EntityData.SegmentPath = "sso-states"
    ssoStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoStates.EntityData.Children = types.NewOrderedMap()
    ssoStates.EntityData.Children.Append("sso-state", types.YChild{"SsoState", nil})
    for i := range ssoStates.SsoState {
        ssoStates.EntityData.Children.Append(types.GetSegmentPath(ssoStates.SsoState[i]), types.YChild{"SsoState", ssoStates.SsoState[i]})
    }
    ssoStates.EntityData.Leafs = types.NewOrderedMap()

    ssoStates.EntityData.YListKeys = []string {}

    return &(ssoStates.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState
// PPP SSO State data for a particular
// interface
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session ID for the interface with SSO State. The
    // type is interface{} with range: 1..4294967295.
    SessionId interface{}

    // SSRP Session ID. The type is interface{} with range: 0..4294967295.
    SessionIdXr interface{}

    // Interface. The type is string with pattern: [a-zA-Z0-9._/-]+.
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

func (ssoState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState) GetEntityData() *types.CommonEntityData {
    ssoState.EntityData.YFilter = ssoState.YFilter
    ssoState.EntityData.YangName = "sso-state"
    ssoState.EntityData.BundleName = "cisco_ios_xr"
    ssoState.EntityData.ParentYangName = "sso-states"
    ssoState.EntityData.SegmentPath = "sso-state" + types.AddKeyToken(ssoState.SessionId, "session-id")
    ssoState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssoState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssoState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssoState.EntityData.Children = types.NewOrderedMap()
    ssoState.EntityData.Children.Append("lcp-state", types.YChild{"LcpState", &ssoState.LcpState})
    ssoState.EntityData.Children.Append("of-us-auth-state", types.YChild{"OfUsAuthState", &ssoState.OfUsAuthState})
    ssoState.EntityData.Children.Append("of-peer-auth-state", types.YChild{"OfPeerAuthState", &ssoState.OfPeerAuthState})
    ssoState.EntityData.Children.Append("ipcp-state", types.YChild{"IpcpState", &ssoState.IpcpState})
    ssoState.EntityData.Leafs = types.NewOrderedMap()
    ssoState.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", ssoState.SessionId})
    ssoState.EntityData.Leafs.Append("session-id-xr", types.YLeaf{"SessionIdXr", ssoState.SessionIdXr})
    ssoState.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", ssoState.Interface})

    ssoState.EntityData.YListKeys = []string {"SessionId"}

    return &(ssoState.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState
// LCP SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (lcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_LcpState) GetEntityData() *types.CommonEntityData {
    lcpState.EntityData.YFilter = lcpState.YFilter
    lcpState.EntityData.YangName = "lcp-state"
    lcpState.EntityData.BundleName = "cisco_ios_xr"
    lcpState.EntityData.ParentYangName = "sso-state"
    lcpState.EntityData.SegmentPath = "lcp-state"
    lcpState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpState.EntityData.Children = types.NewOrderedMap()
    lcpState.EntityData.Leafs = types.NewOrderedMap()
    lcpState.EntityData.Leafs.Append("is-running", types.YLeaf{"IsRunning", lcpState.IsRunning})
    lcpState.EntityData.Leafs.Append("state", types.YLeaf{"State", lcpState.State})

    lcpState.EntityData.YListKeys = []string {}

    return &(lcpState.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState
// Of-us Authentication SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (ofUsAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfUsAuthState) GetEntityData() *types.CommonEntityData {
    ofUsAuthState.EntityData.YFilter = ofUsAuthState.YFilter
    ofUsAuthState.EntityData.YangName = "of-us-auth-state"
    ofUsAuthState.EntityData.BundleName = "cisco_ios_xr"
    ofUsAuthState.EntityData.ParentYangName = "sso-state"
    ofUsAuthState.EntityData.SegmentPath = "of-us-auth-state"
    ofUsAuthState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ofUsAuthState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ofUsAuthState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ofUsAuthState.EntityData.Children = types.NewOrderedMap()
    ofUsAuthState.EntityData.Leafs = types.NewOrderedMap()
    ofUsAuthState.EntityData.Leafs.Append("is-running", types.YLeaf{"IsRunning", ofUsAuthState.IsRunning})
    ofUsAuthState.EntityData.Leafs.Append("state", types.YLeaf{"State", ofUsAuthState.State})

    ofUsAuthState.EntityData.YListKeys = []string {}

    return &(ofUsAuthState.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState
// Of-peer Authentication SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (ofPeerAuthState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_OfPeerAuthState) GetEntityData() *types.CommonEntityData {
    ofPeerAuthState.EntityData.YFilter = ofPeerAuthState.YFilter
    ofPeerAuthState.EntityData.YangName = "of-peer-auth-state"
    ofPeerAuthState.EntityData.BundleName = "cisco_ios_xr"
    ofPeerAuthState.EntityData.ParentYangName = "sso-state"
    ofPeerAuthState.EntityData.SegmentPath = "of-peer-auth-state"
    ofPeerAuthState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ofPeerAuthState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ofPeerAuthState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ofPeerAuthState.EntityData.Children = types.NewOrderedMap()
    ofPeerAuthState.EntityData.Leafs = types.NewOrderedMap()
    ofPeerAuthState.EntityData.Leafs.Append("is-running", types.YLeaf{"IsRunning", ofPeerAuthState.IsRunning})
    ofPeerAuthState.EntityData.Leafs.Append("state", types.YLeaf{"State", ofPeerAuthState.State})

    ofPeerAuthState.EntityData.YListKeys = []string {}

    return &(ofPeerAuthState.EntityData)
}

// Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState
// IPCP SSO State
type Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Is SSO FSM Running. The type is bool.
    IsRunning interface{}

    // SSO FSM State. The type is PppSsoFsmState.
    State interface{}
}

func (ipcpState *Ppp_Nodes_Node_SsoGroups_SsoGroup_SsoStates_SsoState_IpcpState) GetEntityData() *types.CommonEntityData {
    ipcpState.EntityData.YFilter = ipcpState.YFilter
    ipcpState.EntityData.YangName = "ipcp-state"
    ipcpState.EntityData.BundleName = "cisco_ios_xr"
    ipcpState.EntityData.ParentYangName = "sso-state"
    ipcpState.EntityData.SegmentPath = "ipcp-state"
    ipcpState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcpState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcpState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcpState.EntityData.Children = types.NewOrderedMap()
    ipcpState.EntityData.Leafs = types.NewOrderedMap()
    ipcpState.EntityData.Leafs.Append("is-running", types.YLeaf{"IsRunning", ipcpState.IsRunning})
    ipcpState.EntityData.Leafs.Append("state", types.YLeaf{"State", ipcpState.State})

    ipcpState.EntityData.YListKeys = []string {}

    return &(ipcpState.EntityData)
}

// Ppp_Nodes_Node_Summary
// Summarized PPP data for a particular node
type Ppp_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interfaces running PPP.
    Intfs Ppp_Nodes_Node_Summary_Intfs

    // FSM States.
    FsmStates Ppp_Nodes_Node_Summary_FsmStates

    // LCP/Auth Phases.
    LcpAuthPhases Ppp_Nodes_Node_Summary_LcpAuthPhases
}

func (summary *Ppp_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("intfs", types.YChild{"Intfs", &summary.Intfs})
    summary.EntityData.Children.Append("fsm-states", types.YChild{"FsmStates", &summary.FsmStates})
    summary.EntityData.Children.Append("lcp-auth-phases", types.YChild{"LcpAuthPhases", &summary.LcpAuthPhases})
    summary.EntityData.Leafs = types.NewOrderedMap()

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// Ppp_Nodes_Node_Summary_Intfs
// Interfaces running PPP
type Ppp_Nodes_Node_Summary_Intfs struct {
    EntityData types.CommonEntityData
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

func (intfs *Ppp_Nodes_Node_Summary_Intfs) GetEntityData() *types.CommonEntityData {
    intfs.EntityData.YFilter = intfs.YFilter
    intfs.EntityData.YangName = "intfs"
    intfs.EntityData.BundleName = "cisco_ios_xr"
    intfs.EntityData.ParentYangName = "summary"
    intfs.EntityData.SegmentPath = "intfs"
    intfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfs.EntityData.Children = types.NewOrderedMap()
    intfs.EntityData.Leafs = types.NewOrderedMap()
    intfs.EntityData.Leafs.Append("pos-count", types.YLeaf{"PosCount", intfs.PosCount})
    intfs.EntityData.Leafs.Append("serial-count", types.YLeaf{"SerialCount", intfs.SerialCount})
    intfs.EntityData.Leafs.Append("pppoe-count", types.YLeaf{"PppoeCount", intfs.PppoeCount})
    intfs.EntityData.Leafs.Append("multilink-bundle-count", types.YLeaf{"MultilinkBundleCount", intfs.MultilinkBundleCount})
    intfs.EntityData.Leafs.Append("gcc0-count", types.YLeaf{"Gcc0Count", intfs.Gcc0Count})
    intfs.EntityData.Leafs.Append("gcc1-count", types.YLeaf{"Gcc1Count", intfs.Gcc1Count})
    intfs.EntityData.Leafs.Append("total", types.YLeaf{"Total", intfs.Total})

    intfs.EntityData.YListKeys = []string {}

    return &(intfs.EntityData)
}

// Ppp_Nodes_Node_Summary_FsmStates
// FSM States
type Ppp_Nodes_Node_Summary_FsmStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Array of per-LCP FSM States.
    LcpfsmStates Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates

    // Array of per-NCP FSM States. The type is slice of
    // Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray.
    NcpfsmStatesArray []*Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray
}

func (fsmStates *Ppp_Nodes_Node_Summary_FsmStates) GetEntityData() *types.CommonEntityData {
    fsmStates.EntityData.YFilter = fsmStates.YFilter
    fsmStates.EntityData.YangName = "fsm-states"
    fsmStates.EntityData.BundleName = "cisco_ios_xr"
    fsmStates.EntityData.ParentYangName = "summary"
    fsmStates.EntityData.SegmentPath = "fsm-states"
    fsmStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fsmStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fsmStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fsmStates.EntityData.Children = types.NewOrderedMap()
    fsmStates.EntityData.Children.Append("lcpfsm-states", types.YChild{"LcpfsmStates", &fsmStates.LcpfsmStates})
    fsmStates.EntityData.Children.Append("ncpfsm-states-array", types.YChild{"NcpfsmStatesArray", nil})
    for i := range fsmStates.NcpfsmStatesArray {
        fsmStates.EntityData.Children.Append(types.GetSegmentPath(fsmStates.NcpfsmStatesArray[i]), types.YChild{"NcpfsmStatesArray", fsmStates.NcpfsmStatesArray[i]})
    }
    fsmStates.EntityData.Leafs = types.NewOrderedMap()

    fsmStates.EntityData.YListKeys = []string {}

    return &(fsmStates.EntityData)
}

// Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates
// Array of per-LCP FSM States
type Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of LCP FSMs running. The type is interface{} with range:
    // 0..4294967295.
    Total interface{}

    // Number of FSMs in each State. The type is slice of interface{} with range:
    // 0..4294967295.
    Count []interface{}
}

func (lcpfsmStates *Ppp_Nodes_Node_Summary_FsmStates_LcpfsmStates) GetEntityData() *types.CommonEntityData {
    lcpfsmStates.EntityData.YFilter = lcpfsmStates.YFilter
    lcpfsmStates.EntityData.YangName = "lcpfsm-states"
    lcpfsmStates.EntityData.BundleName = "cisco_ios_xr"
    lcpfsmStates.EntityData.ParentYangName = "fsm-states"
    lcpfsmStates.EntityData.SegmentPath = "lcpfsm-states"
    lcpfsmStates.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpfsmStates.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpfsmStates.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpfsmStates.EntityData.Children = types.NewOrderedMap()
    lcpfsmStates.EntityData.Leafs = types.NewOrderedMap()
    lcpfsmStates.EntityData.Leafs.Append("total", types.YLeaf{"Total", lcpfsmStates.Total})
    lcpfsmStates.EntityData.Leafs.Append("count", types.YLeaf{"Count", lcpfsmStates.Count})

    lcpfsmStates.EntityData.YListKeys = []string {}

    return &(lcpfsmStates.EntityData)
}

// Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray
// Array of per-NCP FSM States
type Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray struct {
    EntityData types.CommonEntityData
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

func (ncpfsmStatesArray *Ppp_Nodes_Node_Summary_FsmStates_NcpfsmStatesArray) GetEntityData() *types.CommonEntityData {
    ncpfsmStatesArray.EntityData.YFilter = ncpfsmStatesArray.YFilter
    ncpfsmStatesArray.EntityData.YangName = "ncpfsm-states-array"
    ncpfsmStatesArray.EntityData.BundleName = "cisco_ios_xr"
    ncpfsmStatesArray.EntityData.ParentYangName = "fsm-states"
    ncpfsmStatesArray.EntityData.SegmentPath = "ncpfsm-states-array"
    ncpfsmStatesArray.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ncpfsmStatesArray.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ncpfsmStatesArray.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ncpfsmStatesArray.EntityData.Children = types.NewOrderedMap()
    ncpfsmStatesArray.EntityData.Leafs = types.NewOrderedMap()
    ncpfsmStatesArray.EntityData.Leafs.Append("ncp-identifier", types.YLeaf{"NcpIdentifier", ncpfsmStatesArray.NcpIdentifier})
    ncpfsmStatesArray.EntityData.Leafs.Append("total", types.YLeaf{"Total", ncpfsmStatesArray.Total})
    ncpfsmStatesArray.EntityData.Leafs.Append("count", types.YLeaf{"Count", ncpfsmStatesArray.Count})

    ncpfsmStatesArray.EntityData.YListKeys = []string {}

    return &(ncpfsmStatesArray.EntityData)
}

// Ppp_Nodes_Node_Summary_LcpAuthPhases
// LCP/Auth Phases
type Ppp_Nodes_Node_Summary_LcpAuthPhases struct {
    EntityData types.CommonEntityData
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

func (lcpAuthPhases *Ppp_Nodes_Node_Summary_LcpAuthPhases) GetEntityData() *types.CommonEntityData {
    lcpAuthPhases.EntityData.YFilter = lcpAuthPhases.YFilter
    lcpAuthPhases.EntityData.YangName = "lcp-auth-phases"
    lcpAuthPhases.EntityData.BundleName = "cisco_ios_xr"
    lcpAuthPhases.EntityData.ParentYangName = "summary"
    lcpAuthPhases.EntityData.SegmentPath = "lcp-auth-phases"
    lcpAuthPhases.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lcpAuthPhases.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lcpAuthPhases.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lcpAuthPhases.EntityData.Children = types.NewOrderedMap()
    lcpAuthPhases.EntityData.Leafs = types.NewOrderedMap()
    lcpAuthPhases.EntityData.Leafs.Append("lcp-not-negotiated", types.YLeaf{"LcpNotNegotiated", lcpAuthPhases.LcpNotNegotiated})
    lcpAuthPhases.EntityData.Leafs.Append("authenticating", types.YLeaf{"Authenticating", lcpAuthPhases.Authenticating})
    lcpAuthPhases.EntityData.Leafs.Append("line-held-down", types.YLeaf{"LineHeldDown", lcpAuthPhases.LineHeldDown})
    lcpAuthPhases.EntityData.Leafs.Append("up-local-term", types.YLeaf{"UpLocalTerm", lcpAuthPhases.UpLocalTerm})
    lcpAuthPhases.EntityData.Leafs.Append("up-l2-fwded", types.YLeaf{"UpL2Fwded", lcpAuthPhases.UpL2Fwded})
    lcpAuthPhases.EntityData.Leafs.Append("up-tunneled", types.YLeaf{"UpTunneled", lcpAuthPhases.UpTunneled})

    lcpAuthPhases.EntityData.YListKeys = []string {}

    return &(lcpAuthPhases.EntityData)
}

