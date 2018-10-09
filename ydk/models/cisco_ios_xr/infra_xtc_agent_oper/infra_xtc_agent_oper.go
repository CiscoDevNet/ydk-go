// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc-agent package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pcc: Path-computation client in XTC
//   xtc: xtc
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_xtc_agent_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_xtc_agent_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-agent-oper pcc}", reflect.TypeOf(Pcc{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-agent-oper:pcc", reflect.TypeOf(Pcc{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-agent-oper xtc}", reflect.TypeOf(Xtc{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-agent-oper:xtc", reflect.TypeOf(Xtc{}))
}

// XtcSrSid represents XTC SR SID type
type XtcSrSid string

const (
    // IPv4 Node SID
    XtcSrSid_ipv4_node_sid XtcSrSid = "ipv4-node-sid"

    // IPv4 Adjacency SID
    XtcSrSid_ipv4_adjacency_sid XtcSrSid = "ipv4-adjacency-sid"

    // Unknown SID
    XtcSrSid_unknown_sid XtcSrSid = "unknown-sid"
)

// XtcIgpInfoId represents IGP IDs
type XtcIgpInfoId string

const (
    // ISIS
    XtcIgpInfoId_isis XtcIgpInfoId = "isis"

    // OSPF
    XtcIgpInfoId_ospf XtcIgpInfoId = "ospf"

    // BGP
    XtcIgpInfoId_bgp XtcIgpInfoId = "bgp"
)

// XtcAfId represents Xtc af id
type XtcAfId string

const (
    // None
    XtcAfId_none XtcAfId = "none"

    // IPv4
    XtcAfId_ipv4 XtcAfId = "ipv4"

    // IPv6
    XtcAfId_ipv6 XtcAfId = "ipv6"
)

// XtcBsidMode represents XTC BSID MODE type
type XtcBsidMode string

const (
    // Explicit binding SID
    XtcBsidMode_explicit XtcBsidMode = "explicit"

    // Dynamic binding SID
    XtcBsidMode_dynamic XtcBsidMode = "dynamic"
)

// XtcSid1 represents XTC SID Types
type XtcSid1 string

const (
    // Protected Adjacency SID
    XtcSid1_sr_protected_adj_sid XtcSid1 = "sr-protected-adj-sid"

    // Unprotected Adjacency SID
    XtcSid1_sr_unprotected_adj_sid XtcSid1 = "sr-unprotected-adj-sid"

    // BGP egress peer engineering SID
    XtcSid1_sr_bgp_egress_peer_engineering_sid XtcSid1 = "sr-bgp-egress-peer-engineering-sid"

    // Regular prefix SID
    XtcSid1_sr_reqular_prefix_sid XtcSid1 = "sr-reqular-prefix-sid"

    // Strict prefix SID
    XtcSid1_sr_strict_prefix_sid XtcSid1 = "sr-strict-prefix-sid"
)

// XtcDisjointness represents XTC policy path type
type XtcDisjointness string

const (
    // No Disjointness
    XtcDisjointness_no_disjointness XtcDisjointness = "no-disjointness"

    // Link disjointness
    XtcDisjointness_link_disjointness XtcDisjointness = "link-disjointness"

    // Node disjointness
    XtcDisjointness_node_disjointness XtcDisjointness = "node-disjointness"

    // SRLG disjointness
    XtcDisjointness_srlg_disjointness XtcDisjointness = "srlg-disjointness"

    // SRLG-Node disjointness
    XtcDisjointness_srlg_node_disjointness XtcDisjointness = "srlg-node-disjointness"
)

// XtcBsidError represents Xtc bsid error
type XtcBsidError string

const (
    // No error
    XtcBsidError_none XtcBsidError = "none"

    // Error allocating via LSD
    XtcBsidError_allocating XtcBsidError = "allocating"

    // Explicitly configured BSID already exists
    XtcBsidError_exists XtcBsidError = "exists"

    // Internal error
    XtcBsidError_internal XtcBsidError = "internal"

    // Configured BSID used by another color/end-point
    XtcBsidError_color_endpoint_exists XtcBsidError = "color-endpoint-exists"

    // BSID Forwarding rewrite (label xconnect) failed
    XtcBsidError_forwarding_rewrite_error XtcBsidError = "forwarding-rewrite-error"

    // BSID not valid within SRLB range
    XtcBsidError_srlb_invalid_label XtcBsidError = "srlb-invalid-label"

    // Internal error
    XtcBsidError_internal_error XtcBsidError = "internal-error"
)

// XtcSid represents Xtc sid
type XtcSid string

const (
    // None
    XtcSid_none XtcSid = "none"

    // MPLS
    XtcSid_mpls XtcSid = "mpls"

    // IPv6
    XtcSid_ipv6 XtcSid = "ipv6"
)

// XtcPolicyPath represents XTC policy path type
type XtcPolicyPath string

const (
    // Explicit path
    XtcPolicyPath_explicit XtcPolicyPath = "explicit"

    // Dynamic path
    XtcPolicyPath_dynamic XtcPolicyPath = "dynamic"

    // Dynamic PCE-based path
    XtcPolicyPath_dynamic_pce XtcPolicyPath = "dynamic-pce"
)

// XtcAddressFamily represents Xtc address family
type XtcAddressFamily string

const (
    // IPv4 address family
    XtcAddressFamily_ipv4 XtcAddressFamily = "ipv4"

    // IPv6 address family
    XtcAddressFamily_ipv6 XtcAddressFamily = "ipv6"
)

// Pcc
// Path-computation client in XTC
type Pcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCC PLSP database in XTC.
    Plsps Pcc_Plsps

    // PCC peer database in XTC.
    Peers Pcc_Peers
}

func (pcc *Pcc) GetEntityData() *types.CommonEntityData {
    pcc.EntityData.YFilter = pcc.YFilter
    pcc.EntityData.YangName = "pcc"
    pcc.EntityData.BundleName = "cisco_ios_xr"
    pcc.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-agent-oper"
    pcc.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-agent-oper:pcc"
    pcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcc.EntityData.Children = types.NewOrderedMap()
    pcc.EntityData.Children.Append("plsps", types.YChild{"Plsps", &pcc.Plsps})
    pcc.EntityData.Children.Append("peers", types.YChild{"Peers", &pcc.Peers})
    pcc.EntityData.Leafs = types.NewOrderedMap()

    pcc.EntityData.YListKeys = []string {}

    return &(pcc.EntityData)
}

// Pcc_Plsps
// PCC PLSP database in XTC
type Pcc_Plsps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCC PLSP information. The type is slice of Pcc_Plsps_Plsp.
    Plsp []*Pcc_Plsps_Plsp
}

func (plsps *Pcc_Plsps) GetEntityData() *types.CommonEntityData {
    plsps.EntityData.YFilter = plsps.YFilter
    plsps.EntityData.YangName = "plsps"
    plsps.EntityData.BundleName = "cisco_ios_xr"
    plsps.EntityData.ParentYangName = "pcc"
    plsps.EntityData.SegmentPath = "plsps"
    plsps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    plsps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    plsps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    plsps.EntityData.Children = types.NewOrderedMap()
    plsps.EntityData.Children.Append("plsp", types.YChild{"Plsp", nil})
    for i := range plsps.Plsp {
        plsps.EntityData.Children.Append(types.GetSegmentPath(plsps.Plsp[i]), types.YChild{"Plsp", plsps.Plsp[i]})
    }
    plsps.EntityData.Leafs = types.NewOrderedMap()

    plsps.EntityData.YListKeys = []string {}

    return &(plsps.EntityData)
}

// Pcc_Plsps_Plsp
// PCC PLSP information
type Pcc_Plsps_Plsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. PLSP ID. The type is interface{} with range:
    // 0..4294967295.
    PlspId interface{}

    // PLSP ID. The type is interface{} with range: 0..4294967295.
    PlspIdXr interface{}

    // Symbolic Path Name. The type is string.
    SymPathName interface{}

    // Refcnt. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Refcnt interface{}

    // CONN delegated to. The type is interface{} with range: 0..4294967295.
    ConnDelegatedTo interface{}

    // Stats.
    Stats Pcc_Plsps_Plsp_Stats

    // event history. The type is slice of Pcc_Plsps_Plsp_EventHistory.
    EventHistory []*Pcc_Plsps_Plsp_EventHistory

    // path. The type is slice of Pcc_Plsps_Plsp_Path.
    Path []*Pcc_Plsps_Plsp_Path
}

func (plsp *Pcc_Plsps_Plsp) GetEntityData() *types.CommonEntityData {
    plsp.EntityData.YFilter = plsp.YFilter
    plsp.EntityData.YangName = "plsp"
    plsp.EntityData.BundleName = "cisco_ios_xr"
    plsp.EntityData.ParentYangName = "plsps"
    plsp.EntityData.SegmentPath = "plsp" + types.AddKeyToken(plsp.PlspId, "plsp-id")
    plsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    plsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    plsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    plsp.EntityData.Children = types.NewOrderedMap()
    plsp.EntityData.Children.Append("stats", types.YChild{"Stats", &plsp.Stats})
    plsp.EntityData.Children.Append("event-history", types.YChild{"EventHistory", nil})
    for i := range plsp.EventHistory {
        plsp.EntityData.Children.Append(types.GetSegmentPath(plsp.EventHistory[i]), types.YChild{"EventHistory", plsp.EventHistory[i]})
    }
    plsp.EntityData.Children.Append("path", types.YChild{"Path", nil})
    for i := range plsp.Path {
        plsp.EntityData.Children.Append(types.GetSegmentPath(plsp.Path[i]), types.YChild{"Path", plsp.Path[i]})
    }
    plsp.EntityData.Leafs = types.NewOrderedMap()
    plsp.EntityData.Leafs.Append("plsp-id", types.YLeaf{"PlspId", plsp.PlspId})
    plsp.EntityData.Leafs.Append("plsp-id-xr", types.YLeaf{"PlspIdXr", plsp.PlspIdXr})
    plsp.EntityData.Leafs.Append("sym-path-name", types.YLeaf{"SymPathName", plsp.SymPathName})
    plsp.EntityData.Leafs.Append("refcnt", types.YLeaf{"Refcnt", plsp.Refcnt})
    plsp.EntityData.Leafs.Append("conn-delegated-to", types.YLeaf{"ConnDelegatedTo", plsp.ConnDelegatedTo})

    plsp.EntityData.YListKeys = []string {"PlspId"}

    return &(plsp.EntityData)
}

// Pcc_Plsps_Plsp_Stats
// Stats
type Pcc_Plsps_Plsp_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Paths Created. The type is interface{} with range: 0..18446744073709551615.
    PathsCreated interface{}

    // Paths Destroyed. The type is interface{} with range:
    // 0..18446744073709551615.
    PathsDestroyed interface{}

    // Path create errors. The type is interface{} with range:
    // 0..18446744073709551615.
    PathCreateErrors interface{}

    // Path destroy errors. The type is interface{} with range:
    // 0..18446744073709551615.
    PathDestroyErrors interface{}

    // Requests created. The type is interface{} with range:
    // 0..18446744073709551615.
    RequestsCreated interface{}

    // Requests destroyed. The type is interface{} with range:
    // 0..18446744073709551615.
    RequestsDestroyed interface{}

    // Requests failed. The type is interface{} with range:
    // 0..18446744073709551615.
    RequestsFailed interface{}
}

func (stats *Pcc_Plsps_Plsp_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "plsp"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("paths-created", types.YLeaf{"PathsCreated", stats.PathsCreated})
    stats.EntityData.Leafs.Append("paths-destroyed", types.YLeaf{"PathsDestroyed", stats.PathsDestroyed})
    stats.EntityData.Leafs.Append("path-create-errors", types.YLeaf{"PathCreateErrors", stats.PathCreateErrors})
    stats.EntityData.Leafs.Append("path-destroy-errors", types.YLeaf{"PathDestroyErrors", stats.PathDestroyErrors})
    stats.EntityData.Leafs.Append("requests-created", types.YLeaf{"RequestsCreated", stats.RequestsCreated})
    stats.EntityData.Leafs.Append("requests-destroyed", types.YLeaf{"RequestsDestroyed", stats.RequestsDestroyed})
    stats.EntityData.Leafs.Append("requests-failed", types.YLeaf{"RequestsFailed", stats.RequestsFailed})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// Pcc_Plsps_Plsp_EventHistory
// event history
type Pcc_Plsps_Plsp_EventHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timestamp. The type is interface{} with range: 0..18446744073709551615.
    Ts interface{}

    // Description. The type is string.
    Desc interface{}
}

func (eventHistory *Pcc_Plsps_Plsp_EventHistory) GetEntityData() *types.CommonEntityData {
    eventHistory.EntityData.YFilter = eventHistory.YFilter
    eventHistory.EntityData.YangName = "event-history"
    eventHistory.EntityData.BundleName = "cisco_ios_xr"
    eventHistory.EntityData.ParentYangName = "plsp"
    eventHistory.EntityData.SegmentPath = "event-history"
    eventHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventHistory.EntityData.Children = types.NewOrderedMap()
    eventHistory.EntityData.Leafs = types.NewOrderedMap()
    eventHistory.EntityData.Leafs.Append("ts", types.YLeaf{"Ts", eventHistory.Ts})
    eventHistory.EntityData.Leafs.Append("desc", types.YLeaf{"Desc", eventHistory.Desc})

    eventHistory.EntityData.YListKeys = []string {}

    return &(eventHistory.EntityData)
}

// Pcc_Plsps_Plsp_Path
// path
type Pcc_Plsps_Plsp_Path struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // used bw. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    UsedBw interface{}

    // requested bw. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    RequestedBw interface{}

    // metric value. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    MetricValue interface{}

    // refcnt. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Refcnt interface{}

    // LSP PLSP ID. The type is interface{} with range: 0..4294967295.
    LspPlspId interface{}

    // Binding SID. The type is interface{} with range: 0..4294967295.
    BindingSidValue interface{}

    // Ext Tun ID. The type is interface{} with range: 0..4294967295.
    LspIdTlvExtTunnelId interface{}

    // Tun endpoint address. The type is interface{} with range: 0..4294967295.
    LspIdTlvTunnelEndpointAddress interface{}

    // Tun sender address. The type is interface{} with range: 0..4294967295.
    LspIdTlvTunnelSenderAddress interface{}

    // SRP ID. The type is interface{} with range: 0..4294967295.
    SrpId interface{}

    // LSP ID. The type is interface{} with range: 0..65535.
    LspIdTlvLspId interface{}

    // Tunnel ID. The type is interface{} with range: 0..65535.
    LspIdTlvTunnelId interface{}

    // Application LSP ID. The type is interface{} with range: 0..65535.
    LspId interface{}

    // Binding SID type. The type is interface{} with range: 0..65535.
    BindingSidType interface{}

    // LSP oper flags. The type is interface{} with range: 0..255.
    LspOper interface{}

    // Path setup type. The type is interface{} with range: 0..255.
    PathSetupType interface{}

    // Metric type. The type is interface{} with range: 0..255.
    MetricType interface{}

    // is reported. The type is bool.
    IsReported interface{}

    // LSP A Flag. The type is bool.
    LspAFlag interface{}

    // LSP R Flag. The type is bool.
    LspRFlag interface{}

    // LSP S Flag. The type is bool.
    LspSFlag interface{}

    // LSP D Flag. The type is bool.
    LspDFlag interface{}

    // LSP C Flag. The type is bool.
    LspCFlag interface{}

    // stats.
    Stats Pcc_Plsps_Plsp_Path_Stats

    // ero hop. The type is slice of Pcc_Plsps_Plsp_Path_EroHop.
    EroHop []*Pcc_Plsps_Plsp_Path_EroHop

    // rro hop. The type is slice of Pcc_Plsps_Plsp_Path_RroHop.
    RroHop []*Pcc_Plsps_Plsp_Path_RroHop
}

func (path *Pcc_Plsps_Plsp_Path) GetEntityData() *types.CommonEntityData {
    path.EntityData.YFilter = path.YFilter
    path.EntityData.YangName = "path"
    path.EntityData.BundleName = "cisco_ios_xr"
    path.EntityData.ParentYangName = "plsp"
    path.EntityData.SegmentPath = "path"
    path.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    path.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    path.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    path.EntityData.Children = types.NewOrderedMap()
    path.EntityData.Children.Append("stats", types.YChild{"Stats", &path.Stats})
    path.EntityData.Children.Append("ero-hop", types.YChild{"EroHop", nil})
    for i := range path.EroHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.EroHop[i]), types.YChild{"EroHop", path.EroHop[i]})
    }
    path.EntityData.Children.Append("rro-hop", types.YChild{"RroHop", nil})
    for i := range path.RroHop {
        path.EntityData.Children.Append(types.GetSegmentPath(path.RroHop[i]), types.YChild{"RroHop", path.RroHop[i]})
    }
    path.EntityData.Leafs = types.NewOrderedMap()
    path.EntityData.Leafs.Append("used-bw", types.YLeaf{"UsedBw", path.UsedBw})
    path.EntityData.Leafs.Append("requested-bw", types.YLeaf{"RequestedBw", path.RequestedBw})
    path.EntityData.Leafs.Append("metric-value", types.YLeaf{"MetricValue", path.MetricValue})
    path.EntityData.Leafs.Append("refcnt", types.YLeaf{"Refcnt", path.Refcnt})
    path.EntityData.Leafs.Append("lsp-plsp-id", types.YLeaf{"LspPlspId", path.LspPlspId})
    path.EntityData.Leafs.Append("binding-sid-value", types.YLeaf{"BindingSidValue", path.BindingSidValue})
    path.EntityData.Leafs.Append("lsp-id-tlv-ext-tunnel-id", types.YLeaf{"LspIdTlvExtTunnelId", path.LspIdTlvExtTunnelId})
    path.EntityData.Leafs.Append("lsp-id-tlv-tunnel-endpoint-address", types.YLeaf{"LspIdTlvTunnelEndpointAddress", path.LspIdTlvTunnelEndpointAddress})
    path.EntityData.Leafs.Append("lsp-id-tlv-tunnel-sender-address", types.YLeaf{"LspIdTlvTunnelSenderAddress", path.LspIdTlvTunnelSenderAddress})
    path.EntityData.Leafs.Append("srp-id", types.YLeaf{"SrpId", path.SrpId})
    path.EntityData.Leafs.Append("lsp-id-tlv-lsp-id", types.YLeaf{"LspIdTlvLspId", path.LspIdTlvLspId})
    path.EntityData.Leafs.Append("lsp-id-tlv-tunnel-id", types.YLeaf{"LspIdTlvTunnelId", path.LspIdTlvTunnelId})
    path.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", path.LspId})
    path.EntityData.Leafs.Append("binding-sid-type", types.YLeaf{"BindingSidType", path.BindingSidType})
    path.EntityData.Leafs.Append("lsp-oper", types.YLeaf{"LspOper", path.LspOper})
    path.EntityData.Leafs.Append("path-setup-type", types.YLeaf{"PathSetupType", path.PathSetupType})
    path.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", path.MetricType})
    path.EntityData.Leafs.Append("is-reported", types.YLeaf{"IsReported", path.IsReported})
    path.EntityData.Leafs.Append("lsp-a-flag", types.YLeaf{"LspAFlag", path.LspAFlag})
    path.EntityData.Leafs.Append("lsp-r-flag", types.YLeaf{"LspRFlag", path.LspRFlag})
    path.EntityData.Leafs.Append("lsp-s-flag", types.YLeaf{"LspSFlag", path.LspSFlag})
    path.EntityData.Leafs.Append("lsp-d-flag", types.YLeaf{"LspDFlag", path.LspDFlag})
    path.EntityData.Leafs.Append("lsp-c-flag", types.YLeaf{"LspCFlag", path.LspCFlag})

    path.EntityData.YListKeys = []string {}

    return &(path.EntityData)
}

// Pcc_Plsps_Plsp_Path_Stats
// stats
type Pcc_Plsps_Plsp_Path_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Reports requested. The type is interface{} with range:
    // 0..18446744073709551615.
    ReportsRequested interface{}

    // Reports sent. The type is interface{} with range: 0..18446744073709551615.
    ReportsSent interface{}

    // Reports failed. The type is interface{} with range:
    // 0..18446744073709551615.
    ReportsFailedToSend interface{}
}

func (stats *Pcc_Plsps_Plsp_Path_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "path"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("reports-requested", types.YLeaf{"ReportsRequested", stats.ReportsRequested})
    stats.EntityData.Leafs.Append("reports-sent", types.YLeaf{"ReportsSent", stats.ReportsSent})
    stats.EntityData.Leafs.Append("reports-failed-to-send", types.YLeaf{"ReportsFailedToSend", stats.ReportsFailedToSend})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// Pcc_Plsps_Plsp_Path_EroHop
// ero hop
type Pcc_Plsps_Plsp_Path_EroHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // is loose hop. The type is bool.
    Loose interface{}

    // data.
    Data Pcc_Plsps_Plsp_Path_EroHop_Data
}

func (eroHop *Pcc_Plsps_Plsp_Path_EroHop) GetEntityData() *types.CommonEntityData {
    eroHop.EntityData.YFilter = eroHop.YFilter
    eroHop.EntityData.YangName = "ero-hop"
    eroHop.EntityData.BundleName = "cisco_ios_xr"
    eroHop.EntityData.ParentYangName = "path"
    eroHop.EntityData.SegmentPath = "ero-hop"
    eroHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eroHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eroHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eroHop.EntityData.Children = types.NewOrderedMap()
    eroHop.EntityData.Children.Append("data", types.YChild{"Data", &eroHop.Data})
    eroHop.EntityData.Leafs = types.NewOrderedMap()
    eroHop.EntityData.Leafs.Append("loose", types.YLeaf{"Loose", eroHop.Loose})

    eroHop.EntityData.YListKeys = []string {}

    return &(eroHop.EntityData)
}

// Pcc_Plsps_Plsp_Path_EroHop_Data
// data
type Pcc_Plsps_Plsp_Path_EroHop_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HopType. The type is interface{} with range: 0..255.
    HopType interface{}

    // IPv4 hop info.
    Ipv4 Pcc_Plsps_Plsp_Path_EroHop_Data_Ipv4

    // SR IPv4 hop info.
    SrV4 Pcc_Plsps_Plsp_Path_EroHop_Data_SrV4
}

func (data *Pcc_Plsps_Plsp_Path_EroHop_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "ero-hop"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = types.NewOrderedMap()
    data.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &data.Ipv4})
    data.EntityData.Children.Append("sr-v4", types.YChild{"SrV4", &data.SrV4})
    data.EntityData.Leafs = types.NewOrderedMap()
    data.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", data.HopType})

    data.EntityData.YListKeys = []string {}

    return &(data.EntityData)
}

// Pcc_Plsps_Plsp_Path_EroHop_Data_Ipv4
// IPv4 hop info
type Pcc_Plsps_Plsp_Path_EroHop_Data_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 prefix. The type is interface{} with range: 0..4294967295.
    V4Addr interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLen interface{}
}

func (ipv4 *Pcc_Plsps_Plsp_Path_EroHop_Data_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "data"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("v4-addr", types.YLeaf{"V4Addr", ipv4.V4Addr})
    ipv4.EntityData.Leafs.Append("prefix-len", types.YLeaf{"PrefixLen", ipv4.PrefixLen})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Pcc_Plsps_Plsp_Path_EroHop_Data_SrV4
// SR IPv4 hop info
type Pcc_Plsps_Plsp_Path_EroHop_Data_SrV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is interface{} with range: 0..255.
    Type interface{}

    // C flag. The type is bool.
    Cflag interface{}

    // SID. The type is interface{} with range: 0..4294967295.
    Sid interface{}

    // Remote address. The type is interface{} with range: 0..4294967295.
    RemoteAddr interface{}

    // Local address. The type is interface{} with range: 0..4294967295.
    LocalAddr interface{}
}

func (srV4 *Pcc_Plsps_Plsp_Path_EroHop_Data_SrV4) GetEntityData() *types.CommonEntityData {
    srV4.EntityData.YFilter = srV4.YFilter
    srV4.EntityData.YangName = "sr-v4"
    srV4.EntityData.BundleName = "cisco_ios_xr"
    srV4.EntityData.ParentYangName = "data"
    srV4.EntityData.SegmentPath = "sr-v4"
    srV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srV4.EntityData.Children = types.NewOrderedMap()
    srV4.EntityData.Leafs = types.NewOrderedMap()
    srV4.EntityData.Leafs.Append("type", types.YLeaf{"Type", srV4.Type})
    srV4.EntityData.Leafs.Append("cflag", types.YLeaf{"Cflag", srV4.Cflag})
    srV4.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", srV4.Sid})
    srV4.EntityData.Leafs.Append("remote-addr", types.YLeaf{"RemoteAddr", srV4.RemoteAddr})
    srV4.EntityData.Leafs.Append("local-addr", types.YLeaf{"LocalAddr", srV4.LocalAddr})

    srV4.EntityData.YListKeys = []string {}

    return &(srV4.EntityData)
}

// Pcc_Plsps_Plsp_Path_RroHop
// rro hop
type Pcc_Plsps_Plsp_Path_RroHop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // is loose hop. The type is bool.
    Loose interface{}

    // data.
    Data Pcc_Plsps_Plsp_Path_RroHop_Data
}

func (rroHop *Pcc_Plsps_Plsp_Path_RroHop) GetEntityData() *types.CommonEntityData {
    rroHop.EntityData.YFilter = rroHop.YFilter
    rroHop.EntityData.YangName = "rro-hop"
    rroHop.EntityData.BundleName = "cisco_ios_xr"
    rroHop.EntityData.ParentYangName = "path"
    rroHop.EntityData.SegmentPath = "rro-hop"
    rroHop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rroHop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rroHop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rroHop.EntityData.Children = types.NewOrderedMap()
    rroHop.EntityData.Children.Append("data", types.YChild{"Data", &rroHop.Data})
    rroHop.EntityData.Leafs = types.NewOrderedMap()
    rroHop.EntityData.Leafs.Append("loose", types.YLeaf{"Loose", rroHop.Loose})

    rroHop.EntityData.YListKeys = []string {}

    return &(rroHop.EntityData)
}

// Pcc_Plsps_Plsp_Path_RroHop_Data
// data
type Pcc_Plsps_Plsp_Path_RroHop_Data struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HopType. The type is interface{} with range: 0..255.
    HopType interface{}

    // IPv4 hop info.
    Ipv4 Pcc_Plsps_Plsp_Path_RroHop_Data_Ipv4

    // SR IPv4 hop info.
    SrV4 Pcc_Plsps_Plsp_Path_RroHop_Data_SrV4
}

func (data *Pcc_Plsps_Plsp_Path_RroHop_Data) GetEntityData() *types.CommonEntityData {
    data.EntityData.YFilter = data.YFilter
    data.EntityData.YangName = "data"
    data.EntityData.BundleName = "cisco_ios_xr"
    data.EntityData.ParentYangName = "rro-hop"
    data.EntityData.SegmentPath = "data"
    data.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    data.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    data.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    data.EntityData.Children = types.NewOrderedMap()
    data.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &data.Ipv4})
    data.EntityData.Children.Append("sr-v4", types.YChild{"SrV4", &data.SrV4})
    data.EntityData.Leafs = types.NewOrderedMap()
    data.EntityData.Leafs.Append("hop-type", types.YLeaf{"HopType", data.HopType})

    data.EntityData.YListKeys = []string {}

    return &(data.EntityData)
}

// Pcc_Plsps_Plsp_Path_RroHop_Data_Ipv4
// IPv4 hop info
type Pcc_Plsps_Plsp_Path_RroHop_Data_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 prefix. The type is interface{} with range: 0..4294967295.
    V4Addr interface{}

    // Prefix length. The type is interface{} with range: 0..255.
    PrefixLen interface{}
}

func (ipv4 *Pcc_Plsps_Plsp_Path_RroHop_Data_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "data"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("v4-addr", types.YLeaf{"V4Addr", ipv4.V4Addr})
    ipv4.EntityData.Leafs.Append("prefix-len", types.YLeaf{"PrefixLen", ipv4.PrefixLen})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Pcc_Plsps_Plsp_Path_RroHop_Data_SrV4
// SR IPv4 hop info
type Pcc_Plsps_Plsp_Path_RroHop_Data_SrV4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is interface{} with range: 0..255.
    Type interface{}

    // C flag. The type is bool.
    Cflag interface{}

    // SID. The type is interface{} with range: 0..4294967295.
    Sid interface{}

    // Remote address. The type is interface{} with range: 0..4294967295.
    RemoteAddr interface{}

    // Local address. The type is interface{} with range: 0..4294967295.
    LocalAddr interface{}
}

func (srV4 *Pcc_Plsps_Plsp_Path_RroHop_Data_SrV4) GetEntityData() *types.CommonEntityData {
    srV4.EntityData.YFilter = srV4.YFilter
    srV4.EntityData.YangName = "sr-v4"
    srV4.EntityData.BundleName = "cisco_ios_xr"
    srV4.EntityData.ParentYangName = "data"
    srV4.EntityData.SegmentPath = "sr-v4"
    srV4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srV4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srV4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srV4.EntityData.Children = types.NewOrderedMap()
    srV4.EntityData.Leafs = types.NewOrderedMap()
    srV4.EntityData.Leafs.Append("type", types.YLeaf{"Type", srV4.Type})
    srV4.EntityData.Leafs.Append("cflag", types.YLeaf{"Cflag", srV4.Cflag})
    srV4.EntityData.Leafs.Append("sid", types.YLeaf{"Sid", srV4.Sid})
    srV4.EntityData.Leafs.Append("remote-addr", types.YLeaf{"RemoteAddr", srV4.RemoteAddr})
    srV4.EntityData.Leafs.Append("local-addr", types.YLeaf{"LocalAddr", srV4.LocalAddr})

    srV4.EntityData.YListKeys = []string {}

    return &(srV4.EntityData)
}

// Pcc_Peers
// PCC peer database in XTC
type Pcc_Peers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCC peer information. The type is slice of Pcc_Peers_Peer.
    Peer []*Pcc_Peers_Peer
}

func (peers *Pcc_Peers) GetEntityData() *types.CommonEntityData {
    peers.EntityData.YFilter = peers.YFilter
    peers.EntityData.YangName = "peers"
    peers.EntityData.BundleName = "cisco_ios_xr"
    peers.EntityData.ParentYangName = "pcc"
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

// Pcc_Peers_Peer
// PCC peer information
type Pcc_Peers_Peer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddr interface{}

    // internal handle. The type is interface{} with range: 0..4294967295.
    Handle interface{}

    // connection state. The type is string.
    StateStr interface{}

    // local accepted. The type is bool.
    LocalOk interface{}

    // remote accepted. The type is bool.
    RemoteOk interface{}

    // open retry count. The type is interface{} with range: 0..4294967295.
    OpenRetry interface{}

    // ref count. The type is interface{} with range: 0..4294967295.
    RefCnt interface{}

    // socket state. The type is string.
    RxStateStr interface{}

    // holddown counter. The type is interface{} with range: 0..65535.
    HolddownCounter interface{}

    // PCEP up timestamp. The type is interface{} with range:
    // 0..18446744073709551615.
    PcepUpTs interface{}

    // Precedence. The type is interface{} with range: 0..255.
    Precedence interface{}

    // KA interval local. The type is interface{} with range: 0..4294967295.
    KaIntervalLocal interface{}

    // KA interval remote. The type is interface{} with range: 0..4294967295.
    KaIntervalRemote interface{}

    // Dead interval local. The type is interface{} with range: 0..4294967295.
    DeadIntervalLocal interface{}

    // Dead interval remote. The type is interface{} with range: 0..4294967295.
    DeadIntervalRemote interface{}

    // PCEP session ID local. The type is interface{} with range: 0..4294967295.
    PcepSessionIdLocal interface{}

    // PCEP session ID remote. The type is interface{} with range: 0..4294967295.
    PcepSessionIdRemote interface{}

    // PCEP server Ipv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PcepServerIpv4Addr interface{}

    // PCEP client Ipv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PcepClientIpv4Addr interface{}

    // is stateful local. The type is bool.
    IsStatefulLocal interface{}

    // is stateful remote. The type is bool.
    IsStatefulRemote interface{}

    // is stateful with U flag local. The type is bool.
    IsStatefulUFlagLocal interface{}

    // is stateful with U flag remote. The type is bool.
    IsStatefulUFlagRemote interface{}

    // is segment routing local. The type is bool.
    IsSegmentRoutingLocal interface{}

    // is segment routing remote. The type is bool.
    IsSegmentRoutingRemote interface{}

    // local initiate capability. The type is bool.
    IsInitiateLocal interface{}

    // remote initiate capability. The type is bool.
    IsInitiateRemote interface{}

    // is this the best PCE to delegate to. The type is bool.
    IsBestPce interface{}

    // SR MSD local. The type is interface{} with range: 0..255.
    SrMsdLocal interface{}

    // SR MSD remote. The type is interface{} with range: 0..255.
    SrMsdRemote interface{}

    // socket info.
    SocketInfo Pcc_Peers_Peer_SocketInfo

    // stats.
    Stats Pcc_Peers_Peer_Stats
}

func (peer *Pcc_Peers_Peer) GetEntityData() *types.CommonEntityData {
    peer.EntityData.YFilter = peer.YFilter
    peer.EntityData.YangName = "peer"
    peer.EntityData.BundleName = "cisco_ios_xr"
    peer.EntityData.ParentYangName = "peers"
    peer.EntityData.SegmentPath = "peer" + types.AddKeyToken(peer.PeerAddr, "peer-addr")
    peer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peer.EntityData.Children = types.NewOrderedMap()
    peer.EntityData.Children.Append("socket-info", types.YChild{"SocketInfo", &peer.SocketInfo})
    peer.EntityData.Children.Append("stats", types.YChild{"Stats", &peer.Stats})
    peer.EntityData.Leafs = types.NewOrderedMap()
    peer.EntityData.Leafs.Append("peer-addr", types.YLeaf{"PeerAddr", peer.PeerAddr})
    peer.EntityData.Leafs.Append("handle", types.YLeaf{"Handle", peer.Handle})
    peer.EntityData.Leafs.Append("state-str", types.YLeaf{"StateStr", peer.StateStr})
    peer.EntityData.Leafs.Append("local-ok", types.YLeaf{"LocalOk", peer.LocalOk})
    peer.EntityData.Leafs.Append("remote-ok", types.YLeaf{"RemoteOk", peer.RemoteOk})
    peer.EntityData.Leafs.Append("open-retry", types.YLeaf{"OpenRetry", peer.OpenRetry})
    peer.EntityData.Leafs.Append("ref-cnt", types.YLeaf{"RefCnt", peer.RefCnt})
    peer.EntityData.Leafs.Append("rx-state-str", types.YLeaf{"RxStateStr", peer.RxStateStr})
    peer.EntityData.Leafs.Append("holddown-counter", types.YLeaf{"HolddownCounter", peer.HolddownCounter})
    peer.EntityData.Leafs.Append("pcep-up-ts", types.YLeaf{"PcepUpTs", peer.PcepUpTs})
    peer.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", peer.Precedence})
    peer.EntityData.Leafs.Append("ka-interval-local", types.YLeaf{"KaIntervalLocal", peer.KaIntervalLocal})
    peer.EntityData.Leafs.Append("ka-interval-remote", types.YLeaf{"KaIntervalRemote", peer.KaIntervalRemote})
    peer.EntityData.Leafs.Append("dead-interval-local", types.YLeaf{"DeadIntervalLocal", peer.DeadIntervalLocal})
    peer.EntityData.Leafs.Append("dead-interval-remote", types.YLeaf{"DeadIntervalRemote", peer.DeadIntervalRemote})
    peer.EntityData.Leafs.Append("pcep-session-id-local", types.YLeaf{"PcepSessionIdLocal", peer.PcepSessionIdLocal})
    peer.EntityData.Leafs.Append("pcep-session-id-remote", types.YLeaf{"PcepSessionIdRemote", peer.PcepSessionIdRemote})
    peer.EntityData.Leafs.Append("pcep-server-ipv4-addr", types.YLeaf{"PcepServerIpv4Addr", peer.PcepServerIpv4Addr})
    peer.EntityData.Leafs.Append("pcep-client-ipv4-addr", types.YLeaf{"PcepClientIpv4Addr", peer.PcepClientIpv4Addr})
    peer.EntityData.Leafs.Append("is-stateful-local", types.YLeaf{"IsStatefulLocal", peer.IsStatefulLocal})
    peer.EntityData.Leafs.Append("is-stateful-remote", types.YLeaf{"IsStatefulRemote", peer.IsStatefulRemote})
    peer.EntityData.Leafs.Append("is-stateful-u-flag-local", types.YLeaf{"IsStatefulUFlagLocal", peer.IsStatefulUFlagLocal})
    peer.EntityData.Leafs.Append("is-stateful-u-flag-remote", types.YLeaf{"IsStatefulUFlagRemote", peer.IsStatefulUFlagRemote})
    peer.EntityData.Leafs.Append("is-segment-routing-local", types.YLeaf{"IsSegmentRoutingLocal", peer.IsSegmentRoutingLocal})
    peer.EntityData.Leafs.Append("is-segment-routing-remote", types.YLeaf{"IsSegmentRoutingRemote", peer.IsSegmentRoutingRemote})
    peer.EntityData.Leafs.Append("is-initiate-local", types.YLeaf{"IsInitiateLocal", peer.IsInitiateLocal})
    peer.EntityData.Leafs.Append("is-initiate-remote", types.YLeaf{"IsInitiateRemote", peer.IsInitiateRemote})
    peer.EntityData.Leafs.Append("is-best-pce", types.YLeaf{"IsBestPce", peer.IsBestPce})
    peer.EntityData.Leafs.Append("sr-msd-local", types.YLeaf{"SrMsdLocal", peer.SrMsdLocal})
    peer.EntityData.Leafs.Append("sr-msd-remote", types.YLeaf{"SrMsdRemote", peer.SrMsdRemote})

    peer.EntityData.YListKeys = []string {"PeerAddr"}

    return &(peer.EntityData)
}

// Pcc_Peers_Peer_SocketInfo
// socket info
type Pcc_Peers_Peer_SocketInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // file descriptor. The type is interface{} with range:
    // -9223372036854775808..9223372036854775807.
    Fd interface{}

    // write notify. The type is bool.
    Wnotify interface{}

    // read notify. The type is bool.
    Rnotify interface{}

    // ref count. The type is interface{} with range: 0..4294967295.
    Refcnt interface{}

    // selected. The type is bool.
    Selected interface{}

    // owner. The type is interface{} with range: 0..4294967295.
    Owner interface{}

    // client address. The type is string.
    CsockaddrStr interface{}

    // server address. The type is string.
    SsockaddrStr interface{}
}

func (socketInfo *Pcc_Peers_Peer_SocketInfo) GetEntityData() *types.CommonEntityData {
    socketInfo.EntityData.YFilter = socketInfo.YFilter
    socketInfo.EntityData.YangName = "socket-info"
    socketInfo.EntityData.BundleName = "cisco_ios_xr"
    socketInfo.EntityData.ParentYangName = "peer"
    socketInfo.EntityData.SegmentPath = "socket-info"
    socketInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    socketInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    socketInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    socketInfo.EntityData.Children = types.NewOrderedMap()
    socketInfo.EntityData.Leafs = types.NewOrderedMap()
    socketInfo.EntityData.Leafs.Append("fd", types.YLeaf{"Fd", socketInfo.Fd})
    socketInfo.EntityData.Leafs.Append("wnotify", types.YLeaf{"Wnotify", socketInfo.Wnotify})
    socketInfo.EntityData.Leafs.Append("rnotify", types.YLeaf{"Rnotify", socketInfo.Rnotify})
    socketInfo.EntityData.Leafs.Append("refcnt", types.YLeaf{"Refcnt", socketInfo.Refcnt})
    socketInfo.EntityData.Leafs.Append("selected", types.YLeaf{"Selected", socketInfo.Selected})
    socketInfo.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", socketInfo.Owner})
    socketInfo.EntityData.Leafs.Append("csockaddr-str", types.YLeaf{"CsockaddrStr", socketInfo.CsockaddrStr})
    socketInfo.EntityData.Leafs.Append("ssockaddr-str", types.YLeaf{"SsockaddrStr", socketInfo.SsockaddrStr})

    socketInfo.EntityData.YListKeys = []string {}

    return &(socketInfo.EntityData)
}

// Pcc_Peers_Peer_Stats
// stats
type Pcc_Peers_Peer_Stats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // KA messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    KaMsgRx interface{}

    // KA messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    KaMsgFailRx interface{}

    // KA messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    KaMsgTx interface{}

    // KA messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    KaMsgFailTx interface{}

    // PCREQ messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcreqMsgRx interface{}

    // PCREQ messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcreqMsgFailRx interface{}

    // PCREQ messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcreqMsgTx interface{}

    // PCREQ messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcreqMsgFailTx interface{}

    // PCREP messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrepMsgRx interface{}

    // PCREP messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrepMsgFailRx interface{}

    // PCREP messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrepMsgTx interface{}

    // PCREP messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrepMsgFailTx interface{}

    // PCRPT messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrptMsgRx interface{}

    // PCRPT messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrptMsgFailRx interface{}

    // PCRPT messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrptMsgTx interface{}

    // PCRPT messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcrptMsgFailTx interface{}

    // PCUPD messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcupdMsgRx interface{}

    // PCUPD messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcupdMsgFailRx interface{}

    // PCUPD messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcupdMsgTx interface{}

    // PCUPD messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcupdMsgFailTx interface{}

    // OPEN messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    OpenMsgRx interface{}

    // OPEN messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    OpenMsgFailRx interface{}

    // OPEN messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    OpenMsgTx interface{}

    // OPEN messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    OpenMsgFailTx interface{}

    // PCERR messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcerrMsgRx interface{}

    // PCERR messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcerrMsgFailRx interface{}

    // PCERR messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcerrMsgTx interface{}

    // PCERR messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcerrMsgFailTx interface{}

    // PCNTF messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcntfMsgRx interface{}

    // PCNTF messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcntfMsgFailRx interface{}

    // PCNTF messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcntfMsgTx interface{}

    // PCNTF messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcntfMsgFailTx interface{}

    // PCE EOS messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PceEosMsgTx interface{}

    // PCE EOS messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    PceEosMsgFailTx interface{}

    // CLOSE messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    CloseMsgRx interface{}

    // CLOSE messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    CloseMsgFailRx interface{}

    // CLOSE messages txed. The type is interface{} with range:
    // 0..18446744073709551615.
    CloseMsgTx interface{}

    // CLOSE messages fail txed. The type is interface{} with range:
    // 0..18446744073709551615.
    CloseMsgFailTx interface{}

    // Unexpected messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    UnexpectedMsgRx interface{}

    // Corrupted messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    CorruptedMsgRx interface{}

    // index into recorded reply time. The type is interface{} with range:
    // 0..4294967295.
    ReplyTimeIndex interface{}

    // min reply time. The type is interface{} with range:
    // 0..18446744073709551615.
    MinimumReplyTime interface{}

    // max reply time. The type is interface{} with range:
    // 0..18446744073709551615.
    MaximumReplyTime interface{}

    // requests timed out. The type is interface{} with range:
    // 0..18446744073709551615.
    RequestsTimedOut interface{}

    // last PCERR type received. The type is interface{} with range: 0..255.
    LastPcerrTypeRx interface{}

    // last PCERR value received. The type is interface{} with range: 0..255.
    LastPcerrValRx interface{}

    // last time when PCERR was received. The type is interface{} with range:
    // 0..18446744073709551615.
    LastPcerrRxTs interface{}

    // last PCERR type transmitted. The type is interface{} with range: 0..255.
    LastPcerrTypeTx interface{}

    // last PCERR value transmitted. The type is interface{} with range: 0..255.
    LastPcerrValTx interface{}

    // last time when PCERR was transmitted. The type is interface{} with range:
    // 0..18446744073709551615.
    LastPcerrTxTs interface{}

    // PCINITIATE messages rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcinitiateMsgRx interface{}

    // PCINITIATE messages fail rxed. The type is interface{} with range:
    // 0..18446744073709551615.
    PcinitiateMsgFailRx interface{}

    // Recorded reply time. The type is slice of interface{} with range:
    // 0..18446744073709551615.
    RecordedReplyTime []interface{}
}

func (stats *Pcc_Peers_Peer_Stats) GetEntityData() *types.CommonEntityData {
    stats.EntityData.YFilter = stats.YFilter
    stats.EntityData.YangName = "stats"
    stats.EntityData.BundleName = "cisco_ios_xr"
    stats.EntityData.ParentYangName = "peer"
    stats.EntityData.SegmentPath = "stats"
    stats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stats.EntityData.Children = types.NewOrderedMap()
    stats.EntityData.Leafs = types.NewOrderedMap()
    stats.EntityData.Leafs.Append("ka-msg-rx", types.YLeaf{"KaMsgRx", stats.KaMsgRx})
    stats.EntityData.Leafs.Append("ka-msg-fail-rx", types.YLeaf{"KaMsgFailRx", stats.KaMsgFailRx})
    stats.EntityData.Leafs.Append("ka-msg-tx", types.YLeaf{"KaMsgTx", stats.KaMsgTx})
    stats.EntityData.Leafs.Append("ka-msg-fail-tx", types.YLeaf{"KaMsgFailTx", stats.KaMsgFailTx})
    stats.EntityData.Leafs.Append("pcreq-msg-rx", types.YLeaf{"PcreqMsgRx", stats.PcreqMsgRx})
    stats.EntityData.Leafs.Append("pcreq-msg-fail-rx", types.YLeaf{"PcreqMsgFailRx", stats.PcreqMsgFailRx})
    stats.EntityData.Leafs.Append("pcreq-msg-tx", types.YLeaf{"PcreqMsgTx", stats.PcreqMsgTx})
    stats.EntityData.Leafs.Append("pcreq-msg-fail-tx", types.YLeaf{"PcreqMsgFailTx", stats.PcreqMsgFailTx})
    stats.EntityData.Leafs.Append("pcrep-msg-rx", types.YLeaf{"PcrepMsgRx", stats.PcrepMsgRx})
    stats.EntityData.Leafs.Append("pcrep-msg-fail-rx", types.YLeaf{"PcrepMsgFailRx", stats.PcrepMsgFailRx})
    stats.EntityData.Leafs.Append("pcrep-msg-tx", types.YLeaf{"PcrepMsgTx", stats.PcrepMsgTx})
    stats.EntityData.Leafs.Append("pcrep-msg-fail-tx", types.YLeaf{"PcrepMsgFailTx", stats.PcrepMsgFailTx})
    stats.EntityData.Leafs.Append("pcrpt-msg-rx", types.YLeaf{"PcrptMsgRx", stats.PcrptMsgRx})
    stats.EntityData.Leafs.Append("pcrpt-msg-fail-rx", types.YLeaf{"PcrptMsgFailRx", stats.PcrptMsgFailRx})
    stats.EntityData.Leafs.Append("pcrpt-msg-tx", types.YLeaf{"PcrptMsgTx", stats.PcrptMsgTx})
    stats.EntityData.Leafs.Append("pcrpt-msg-fail-tx", types.YLeaf{"PcrptMsgFailTx", stats.PcrptMsgFailTx})
    stats.EntityData.Leafs.Append("pcupd-msg-rx", types.YLeaf{"PcupdMsgRx", stats.PcupdMsgRx})
    stats.EntityData.Leafs.Append("pcupd-msg-fail-rx", types.YLeaf{"PcupdMsgFailRx", stats.PcupdMsgFailRx})
    stats.EntityData.Leafs.Append("pcupd-msg-tx", types.YLeaf{"PcupdMsgTx", stats.PcupdMsgTx})
    stats.EntityData.Leafs.Append("pcupd-msg-fail-tx", types.YLeaf{"PcupdMsgFailTx", stats.PcupdMsgFailTx})
    stats.EntityData.Leafs.Append("open-msg-rx", types.YLeaf{"OpenMsgRx", stats.OpenMsgRx})
    stats.EntityData.Leafs.Append("open-msg-fail-rx", types.YLeaf{"OpenMsgFailRx", stats.OpenMsgFailRx})
    stats.EntityData.Leafs.Append("open-msg-tx", types.YLeaf{"OpenMsgTx", stats.OpenMsgTx})
    stats.EntityData.Leafs.Append("open-msg-fail-tx", types.YLeaf{"OpenMsgFailTx", stats.OpenMsgFailTx})
    stats.EntityData.Leafs.Append("pcerr-msg-rx", types.YLeaf{"PcerrMsgRx", stats.PcerrMsgRx})
    stats.EntityData.Leafs.Append("pcerr-msg-fail-rx", types.YLeaf{"PcerrMsgFailRx", stats.PcerrMsgFailRx})
    stats.EntityData.Leafs.Append("pcerr-msg-tx", types.YLeaf{"PcerrMsgTx", stats.PcerrMsgTx})
    stats.EntityData.Leafs.Append("pcerr-msg-fail-tx", types.YLeaf{"PcerrMsgFailTx", stats.PcerrMsgFailTx})
    stats.EntityData.Leafs.Append("pcntf-msg-rx", types.YLeaf{"PcntfMsgRx", stats.PcntfMsgRx})
    stats.EntityData.Leafs.Append("pcntf-msg-fail-rx", types.YLeaf{"PcntfMsgFailRx", stats.PcntfMsgFailRx})
    stats.EntityData.Leafs.Append("pcntf-msg-tx", types.YLeaf{"PcntfMsgTx", stats.PcntfMsgTx})
    stats.EntityData.Leafs.Append("pcntf-msg-fail-tx", types.YLeaf{"PcntfMsgFailTx", stats.PcntfMsgFailTx})
    stats.EntityData.Leafs.Append("pce-eos-msg-tx", types.YLeaf{"PceEosMsgTx", stats.PceEosMsgTx})
    stats.EntityData.Leafs.Append("pce-eos-msg-fail-tx", types.YLeaf{"PceEosMsgFailTx", stats.PceEosMsgFailTx})
    stats.EntityData.Leafs.Append("close-msg-rx", types.YLeaf{"CloseMsgRx", stats.CloseMsgRx})
    stats.EntityData.Leafs.Append("close-msg-fail-rx", types.YLeaf{"CloseMsgFailRx", stats.CloseMsgFailRx})
    stats.EntityData.Leafs.Append("close-msg-tx", types.YLeaf{"CloseMsgTx", stats.CloseMsgTx})
    stats.EntityData.Leafs.Append("close-msg-fail-tx", types.YLeaf{"CloseMsgFailTx", stats.CloseMsgFailTx})
    stats.EntityData.Leafs.Append("unexpected-msg-rx", types.YLeaf{"UnexpectedMsgRx", stats.UnexpectedMsgRx})
    stats.EntityData.Leafs.Append("corrupted-msg-rx", types.YLeaf{"CorruptedMsgRx", stats.CorruptedMsgRx})
    stats.EntityData.Leafs.Append("reply-time-index", types.YLeaf{"ReplyTimeIndex", stats.ReplyTimeIndex})
    stats.EntityData.Leafs.Append("minimum-reply-time", types.YLeaf{"MinimumReplyTime", stats.MinimumReplyTime})
    stats.EntityData.Leafs.Append("maximum-reply-time", types.YLeaf{"MaximumReplyTime", stats.MaximumReplyTime})
    stats.EntityData.Leafs.Append("requests-timed-out", types.YLeaf{"RequestsTimedOut", stats.RequestsTimedOut})
    stats.EntityData.Leafs.Append("last-pcerr-type-rx", types.YLeaf{"LastPcerrTypeRx", stats.LastPcerrTypeRx})
    stats.EntityData.Leafs.Append("last-pcerr-val-rx", types.YLeaf{"LastPcerrValRx", stats.LastPcerrValRx})
    stats.EntityData.Leafs.Append("last-pcerr-rx-ts", types.YLeaf{"LastPcerrRxTs", stats.LastPcerrRxTs})
    stats.EntityData.Leafs.Append("last-pcerr-type-tx", types.YLeaf{"LastPcerrTypeTx", stats.LastPcerrTypeTx})
    stats.EntityData.Leafs.Append("last-pcerr-val-tx", types.YLeaf{"LastPcerrValTx", stats.LastPcerrValTx})
    stats.EntityData.Leafs.Append("last-pcerr-tx-ts", types.YLeaf{"LastPcerrTxTs", stats.LastPcerrTxTs})
    stats.EntityData.Leafs.Append("pcinitiate-msg-rx", types.YLeaf{"PcinitiateMsgRx", stats.PcinitiateMsgRx})
    stats.EntityData.Leafs.Append("pcinitiate-msg-fail-rx", types.YLeaf{"PcinitiateMsgFailRx", stats.PcinitiateMsgFailRx})
    stats.EntityData.Leafs.Append("recorded-reply-time", types.YLeaf{"RecordedReplyTime", stats.RecordedReplyTime})

    stats.EntityData.YListKeys = []string {}

    return &(stats.EntityData)
}

// Xtc
// xtc
type Xtc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy database in XTC Agent.
    Policies Xtc_Policies

    // Summary of all policies.
    PolicySummary Xtc_PolicySummary

    // On-Demand Color database in XTC Agent.
    OnDemandColors Xtc_OnDemandColors

    // Forwarding information.
    Forwarding Xtc_Forwarding

    // Controller information.
    Controller Xtc_Controller

    // Node summary database.
    TopologySummary Xtc_TopologySummary

    // Node database in XTC Agent.
    TopologyNodes Xtc_TopologyNodes

    // Prefixes database in XTC Agent.
    PrefixInfos Xtc_PrefixInfos
}

func (xtc *Xtc) GetEntityData() *types.CommonEntityData {
    xtc.EntityData.YFilter = xtc.YFilter
    xtc.EntityData.YangName = "xtc"
    xtc.EntityData.BundleName = "cisco_ios_xr"
    xtc.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-agent-oper"
    xtc.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-agent-oper:xtc"
    xtc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xtc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xtc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xtc.EntityData.Children = types.NewOrderedMap()
    xtc.EntityData.Children.Append("policies", types.YChild{"Policies", &xtc.Policies})
    xtc.EntityData.Children.Append("policy-summary", types.YChild{"PolicySummary", &xtc.PolicySummary})
    xtc.EntityData.Children.Append("on-demand-colors", types.YChild{"OnDemandColors", &xtc.OnDemandColors})
    xtc.EntityData.Children.Append("forwarding", types.YChild{"Forwarding", &xtc.Forwarding})
    xtc.EntityData.Children.Append("controller", types.YChild{"Controller", &xtc.Controller})
    xtc.EntityData.Children.Append("topology-summary", types.YChild{"TopologySummary", &xtc.TopologySummary})
    xtc.EntityData.Children.Append("topology-nodes", types.YChild{"TopologyNodes", &xtc.TopologyNodes})
    xtc.EntityData.Children.Append("prefix-infos", types.YChild{"PrefixInfos", &xtc.PrefixInfos})
    xtc.EntityData.Leafs = types.NewOrderedMap()

    xtc.EntityData.YListKeys = []string {}

    return &(xtc.EntityData)
}

// Xtc_Policies
// Policy database in XTC Agent
type Xtc_Policies struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy information. The type is slice of Xtc_Policies_Policy.
    Policy []*Xtc_Policies_Policy
}

func (policies *Xtc_Policies) GetEntityData() *types.CommonEntityData {
    policies.EntityData.YFilter = policies.YFilter
    policies.EntityData.YangName = "policies"
    policies.EntityData.BundleName = "cisco_ios_xr"
    policies.EntityData.ParentYangName = "xtc"
    policies.EntityData.SegmentPath = "policies"
    policies.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policies.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policies.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policies.EntityData.Children = types.NewOrderedMap()
    policies.EntityData.Children.Append("policy", types.YChild{"Policy", nil})
    for i := range policies.Policy {
        policies.EntityData.Children.Append(types.GetSegmentPath(policies.Policy[i]), types.YChild{"Policy", policies.Policy[i]})
    }
    policies.EntityData.Leafs = types.NewOrderedMap()

    policies.EntityData.YListKeys = []string {}

    return &(policies.EntityData)
}

// Xtc_Policies_Policy
// Policy information
type Xtc_Policies_Policy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Policy ID. The type is interface{} with range:
    // 0..4294967295.
    Id interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // Admin up. The type is interface{} with range: 0..4294967295.
    AdministrativeUp interface{}

    // Operational up. The type is interface{} with range: 0..4294967295.
    OperationalUp interface{}

    // Color. The type is interface{} with range: 0..4294967295.
    Color interface{}

    // Whether policy was automatically created or configured. The type is bool.
    IsAutoPolicy interface{}

    // Indicates number of up/down transitions. The type is interface{} with
    // range: 0..4294967295.
    TransitionCount interface{}

    // Forward class of the policy. The type is interface{} with range:
    // 0..4294967295.
    ForwardClass interface{}

    // Policy up time in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    UpTime interface{}

    // Policy up age (since) in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    UpAge interface{}

    // Policy down time in nano seconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    DownTime interface{}

    // Policy down age (since) in nano seconds. The type is interface{} with
    // range: 0..18446744073709551615. Units are nanosecond.
    DownAge interface{}

    // LSP ID. The type is interface{} with range: 0..4294967295.
    LspId interface{}

    // Whether steering to BGP client is disabled. The type is bool.
    SteeringBgpDisabled interface{}

    // Interface handle. The type is interface{} with range: 0..4294967295.
    InterfaceHandle interface{}

    // Policy group identifier. The type is interface{} with range: 0..65535.
    PolicyGroupIdentifier interface{}

    // Local label identifier. The type is interface{} with range: 0..65535.
    LocalLabelIdentifier interface{}

    // Local label. The type is interface{} with range: 0..4294967295.
    LocalLabel interface{}

    // Profile ID. The type is interface{} with range: 0..65535.
    ProfileId interface{}

    // IPv6 caps enabled. The type is bool.
    Ipv6CapsEnabled interface{}

    // Destination address.
    DestinationAddress Xtc_Policies_Policy_DestinationAddress

    // Binding SID information.
    BindingSid Xtc_Policies_Policy_BindingSid

    // Autopolicy information.
    AutoPolicyInfo Xtc_Policies_Policy_AutoPolicyInfo

    // Path options. The type is slice of Xtc_Policies_Policy_Paths.
    Paths []*Xtc_Policies_Policy_Paths
}

func (policy *Xtc_Policies_Policy) GetEntityData() *types.CommonEntityData {
    policy.EntityData.YFilter = policy.YFilter
    policy.EntityData.YangName = "policy"
    policy.EntityData.BundleName = "cisco_ios_xr"
    policy.EntityData.ParentYangName = "policies"
    policy.EntityData.SegmentPath = "policy" + types.AddKeyToken(policy.Id, "id")
    policy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policy.EntityData.Children = types.NewOrderedMap()
    policy.EntityData.Children.Append("destination-address", types.YChild{"DestinationAddress", &policy.DestinationAddress})
    policy.EntityData.Children.Append("binding-sid", types.YChild{"BindingSid", &policy.BindingSid})
    policy.EntityData.Children.Append("auto-policy-info", types.YChild{"AutoPolicyInfo", &policy.AutoPolicyInfo})
    policy.EntityData.Children.Append("paths", types.YChild{"Paths", nil})
    for i := range policy.Paths {
        policy.EntityData.Children.Append(types.GetSegmentPath(policy.Paths[i]), types.YChild{"Paths", policy.Paths[i]})
    }
    policy.EntityData.Leafs = types.NewOrderedMap()
    policy.EntityData.Leafs.Append("id", types.YLeaf{"Id", policy.Id})
    policy.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policy.PolicyName})
    policy.EntityData.Leafs.Append("administrative-up", types.YLeaf{"AdministrativeUp", policy.AdministrativeUp})
    policy.EntityData.Leafs.Append("operational-up", types.YLeaf{"OperationalUp", policy.OperationalUp})
    policy.EntityData.Leafs.Append("color", types.YLeaf{"Color", policy.Color})
    policy.EntityData.Leafs.Append("is-auto-policy", types.YLeaf{"IsAutoPolicy", policy.IsAutoPolicy})
    policy.EntityData.Leafs.Append("transition-count", types.YLeaf{"TransitionCount", policy.TransitionCount})
    policy.EntityData.Leafs.Append("forward-class", types.YLeaf{"ForwardClass", policy.ForwardClass})
    policy.EntityData.Leafs.Append("up-time", types.YLeaf{"UpTime", policy.UpTime})
    policy.EntityData.Leafs.Append("up-age", types.YLeaf{"UpAge", policy.UpAge})
    policy.EntityData.Leafs.Append("down-time", types.YLeaf{"DownTime", policy.DownTime})
    policy.EntityData.Leafs.Append("down-age", types.YLeaf{"DownAge", policy.DownAge})
    policy.EntityData.Leafs.Append("lsp-id", types.YLeaf{"LspId", policy.LspId})
    policy.EntityData.Leafs.Append("steering-bgp-disabled", types.YLeaf{"SteeringBgpDisabled", policy.SteeringBgpDisabled})
    policy.EntityData.Leafs.Append("interface-handle", types.YLeaf{"InterfaceHandle", policy.InterfaceHandle})
    policy.EntityData.Leafs.Append("policy-group-identifier", types.YLeaf{"PolicyGroupIdentifier", policy.PolicyGroupIdentifier})
    policy.EntityData.Leafs.Append("local-label-identifier", types.YLeaf{"LocalLabelIdentifier", policy.LocalLabelIdentifier})
    policy.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", policy.LocalLabel})
    policy.EntityData.Leafs.Append("profile-id", types.YLeaf{"ProfileId", policy.ProfileId})
    policy.EntityData.Leafs.Append("ipv6-caps-enabled", types.YLeaf{"Ipv6CapsEnabled", policy.Ipv6CapsEnabled})

    policy.EntityData.YListKeys = []string {"Id"}

    return &(policy.EntityData)
}

// Xtc_Policies_Policy_DestinationAddress
// Destination address
type Xtc_Policies_Policy_DestinationAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (destinationAddress *Xtc_Policies_Policy_DestinationAddress) GetEntityData() *types.CommonEntityData {
    destinationAddress.EntityData.YFilter = destinationAddress.YFilter
    destinationAddress.EntityData.YangName = "destination-address"
    destinationAddress.EntityData.BundleName = "cisco_ios_xr"
    destinationAddress.EntityData.ParentYangName = "policy"
    destinationAddress.EntityData.SegmentPath = "destination-address"
    destinationAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destinationAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destinationAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destinationAddress.EntityData.Children = types.NewOrderedMap()
    destinationAddress.EntityData.Leafs = types.NewOrderedMap()
    destinationAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", destinationAddress.AfName})
    destinationAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", destinationAddress.Ipv4})
    destinationAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", destinationAddress.Ipv6})

    destinationAddress.EntityData.YListKeys = []string {}

    return &(destinationAddress.EntityData)
}

// Xtc_Policies_Policy_BindingSid
// Binding SID information
type Xtc_Policies_Policy_BindingSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Binding SID Mode. The type is XtcBsidMode.
    BsidMode interface{}

    // Binding SID error, if any. The type is XtcBsidError.
    Error interface{}

    // State of the binding SID. The type is string.
    State interface{}

    // Whether the binding SID is explicit-based. The type is bool.
    ExplicitBased interface{}

    // Whether the policy is selected for forwarding on this BSID. The type is
    // bool.
    PolicySelected interface{}

    // Binding SID value.
    Value Xtc_Policies_Policy_BindingSid_Value
}

func (bindingSid *Xtc_Policies_Policy_BindingSid) GetEntityData() *types.CommonEntityData {
    bindingSid.EntityData.YFilter = bindingSid.YFilter
    bindingSid.EntityData.YangName = "binding-sid"
    bindingSid.EntityData.BundleName = "cisco_ios_xr"
    bindingSid.EntityData.ParentYangName = "policy"
    bindingSid.EntityData.SegmentPath = "binding-sid"
    bindingSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bindingSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bindingSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bindingSid.EntityData.Children = types.NewOrderedMap()
    bindingSid.EntityData.Children.Append("value", types.YChild{"Value", &bindingSid.Value})
    bindingSid.EntityData.Leafs = types.NewOrderedMap()
    bindingSid.EntityData.Leafs.Append("bsid-mode", types.YLeaf{"BsidMode", bindingSid.BsidMode})
    bindingSid.EntityData.Leafs.Append("error", types.YLeaf{"Error", bindingSid.Error})
    bindingSid.EntityData.Leafs.Append("state", types.YLeaf{"State", bindingSid.State})
    bindingSid.EntityData.Leafs.Append("explicit-based", types.YLeaf{"ExplicitBased", bindingSid.ExplicitBased})
    bindingSid.EntityData.Leafs.Append("policy-selected", types.YLeaf{"PolicySelected", bindingSid.PolicySelected})

    bindingSid.EntityData.YListKeys = []string {}

    return &(bindingSid.EntityData)
}

// Xtc_Policies_Policy_BindingSid_Value
// Binding SID value
type Xtc_Policies_Policy_BindingSid_Value struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDType. The type is XtcSid.
    SidType interface{}

    // MPLS label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (value *Xtc_Policies_Policy_BindingSid_Value) GetEntityData() *types.CommonEntityData {
    value.EntityData.YFilter = value.YFilter
    value.EntityData.YangName = "value"
    value.EntityData.BundleName = "cisco_ios_xr"
    value.EntityData.ParentYangName = "binding-sid"
    value.EntityData.SegmentPath = "value"
    value.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    value.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    value.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    value.EntityData.Children = types.NewOrderedMap()
    value.EntityData.Leafs = types.NewOrderedMap()
    value.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", value.SidType})
    value.EntityData.Leafs.Append("label", types.YLeaf{"Label", value.Label})
    value.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", value.Ipv6})

    value.EntityData.YListKeys = []string {}

    return &(value.EntityData)
}

// Xtc_Policies_Policy_AutoPolicyInfo
// Autopolicy information
type Xtc_Policies_Policy_AutoPolicyInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of client who created policy. The type is string.
    CreatorName interface{}

    // Distinguisher. The type is interface{} with range: 0..4294967295.
    Distinguisher interface{}

    // Preference of the policy. The type is interface{} with range:
    // 0..4294967295.
    Preference interface{}
}

func (autoPolicyInfo *Xtc_Policies_Policy_AutoPolicyInfo) GetEntityData() *types.CommonEntityData {
    autoPolicyInfo.EntityData.YFilter = autoPolicyInfo.YFilter
    autoPolicyInfo.EntityData.YangName = "auto-policy-info"
    autoPolicyInfo.EntityData.BundleName = "cisco_ios_xr"
    autoPolicyInfo.EntityData.ParentYangName = "policy"
    autoPolicyInfo.EntityData.SegmentPath = "auto-policy-info"
    autoPolicyInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoPolicyInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoPolicyInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoPolicyInfo.EntityData.Children = types.NewOrderedMap()
    autoPolicyInfo.EntityData.Leafs = types.NewOrderedMap()
    autoPolicyInfo.EntityData.Leafs.Append("creator-name", types.YLeaf{"CreatorName", autoPolicyInfo.CreatorName})
    autoPolicyInfo.EntityData.Leafs.Append("distinguisher", types.YLeaf{"Distinguisher", autoPolicyInfo.Distinguisher})
    autoPolicyInfo.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", autoPolicyInfo.Preference})

    autoPolicyInfo.EntityData.YListKeys = []string {}

    return &(autoPolicyInfo.EntityData)
}

// Xtc_Policies_Policy_Paths
// Path options
type Xtc_Policies_Policy_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index number. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // Path option type. The type is XtcPolicyPath.
    Type interface{}

    // Explicit path option name. The type is string.
    Name interface{}

    // Whether the path is active (used). The type is bool.
    Active interface{}

    // Configured weight of the path-option. The type is interface{} with range:
    // 0..4294967295.
    Weight interface{}

    // Configured path metric type. The type is interface{} with range: 0..255.
    MetricType interface{}

    // Path metric value. The type is interface{} with range: 0..4294967295.
    MetricValue interface{}

    // True if path is valid. The type is bool.
    IsValid interface{}

    // True if the path is to be computed by PCE. The type is bool.
    PceBasedPath interface{}

    // Address of the PCE computed the path. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PceAddress interface{}

    // Error (for display only). The type is string.
    Error interface{}

    // SR path constraints.
    SrPathConstraints Xtc_Policies_Policy_Paths_SrPathConstraints

    // SR hop list. The type is slice of Xtc_Policies_Policy_Paths_Hops.
    Hops []*Xtc_Policies_Policy_Paths_Hops
}

func (paths *Xtc_Policies_Policy_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "policy"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Children.Append("sr-path-constraints", types.YChild{"SrPathConstraints", &paths.SrPathConstraints})
    paths.EntityData.Children.Append("hops", types.YChild{"Hops", nil})
    for i := range paths.Hops {
        paths.EntityData.Children.Append(types.GetSegmentPath(paths.Hops[i]), types.YChild{"Hops", paths.Hops[i]})
    }
    paths.EntityData.Leafs = types.NewOrderedMap()
    paths.EntityData.Leafs.Append("index", types.YLeaf{"Index", paths.Index})
    paths.EntityData.Leafs.Append("type", types.YLeaf{"Type", paths.Type})
    paths.EntityData.Leafs.Append("name", types.YLeaf{"Name", paths.Name})
    paths.EntityData.Leafs.Append("active", types.YLeaf{"Active", paths.Active})
    paths.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", paths.Weight})
    paths.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", paths.MetricType})
    paths.EntityData.Leafs.Append("metric-value", types.YLeaf{"MetricValue", paths.MetricValue})
    paths.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", paths.IsValid})
    paths.EntityData.Leafs.Append("pce-based-path", types.YLeaf{"PceBasedPath", paths.PceBasedPath})
    paths.EntityData.Leafs.Append("pce-address", types.YLeaf{"PceAddress", paths.PceAddress})
    paths.EntityData.Leafs.Append("error", types.YLeaf{"Error", paths.Error})

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// Xtc_Policies_Policy_Paths_SrPathConstraints
// SR path constraints
type Xtc_Policies_Policy_Paths_SrPathConstraints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path metrics.
    PathMetrics Xtc_Policies_Policy_Paths_SrPathConstraints_PathMetrics

    // Segments constraints.
    Segments Xtc_Policies_Policy_Paths_SrPathConstraints_Segments

    // Affinity constraints list. The type is slice of
    // Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint.
    AffinityConstraint []*Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint
}

func (srPathConstraints *Xtc_Policies_Policy_Paths_SrPathConstraints) GetEntityData() *types.CommonEntityData {
    srPathConstraints.EntityData.YFilter = srPathConstraints.YFilter
    srPathConstraints.EntityData.YangName = "sr-path-constraints"
    srPathConstraints.EntityData.BundleName = "cisco_ios_xr"
    srPathConstraints.EntityData.ParentYangName = "paths"
    srPathConstraints.EntityData.SegmentPath = "sr-path-constraints"
    srPathConstraints.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srPathConstraints.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srPathConstraints.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srPathConstraints.EntityData.Children = types.NewOrderedMap()
    srPathConstraints.EntityData.Children.Append("path-metrics", types.YChild{"PathMetrics", &srPathConstraints.PathMetrics})
    srPathConstraints.EntityData.Children.Append("segments", types.YChild{"Segments", &srPathConstraints.Segments})
    srPathConstraints.EntityData.Children.Append("affinity-constraint", types.YChild{"AffinityConstraint", nil})
    for i := range srPathConstraints.AffinityConstraint {
        srPathConstraints.EntityData.Children.Append(types.GetSegmentPath(srPathConstraints.AffinityConstraint[i]), types.YChild{"AffinityConstraint", srPathConstraints.AffinityConstraint[i]})
    }
    srPathConstraints.EntityData.Leafs = types.NewOrderedMap()

    srPathConstraints.EntityData.YListKeys = []string {}

    return &(srPathConstraints.EntityData)
}

// Xtc_Policies_Policy_Paths_SrPathConstraints_PathMetrics
// Path metrics
type Xtc_Policies_Policy_Paths_SrPathConstraints_PathMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Margin Relative. The type is interface{} with range: 0..255.
    MarginRelative interface{}

    // Margin Absolute. The type is interface{} with range: 0..255.
    MarginAbsolute interface{}

    // Maximum number of segments. The type is interface{} with range: 0..65535.
    MaximumSegments interface{}

    // Accumulative TE metric. The type is interface{} with range: 0..4294967295.
    AccumulativeTeMetric interface{}

    // Accumulative IGP metric. The type is interface{} with range: 0..4294967295.
    AccumulativeIgpMetric interface{}

    // Accumulative delay. The type is interface{} with range: 0..4294967295.
    AccumulativeDelay interface{}
}

func (pathMetrics *Xtc_Policies_Policy_Paths_SrPathConstraints_PathMetrics) GetEntityData() *types.CommonEntityData {
    pathMetrics.EntityData.YFilter = pathMetrics.YFilter
    pathMetrics.EntityData.YangName = "path-metrics"
    pathMetrics.EntityData.BundleName = "cisco_ios_xr"
    pathMetrics.EntityData.ParentYangName = "sr-path-constraints"
    pathMetrics.EntityData.SegmentPath = "path-metrics"
    pathMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathMetrics.EntityData.Children = types.NewOrderedMap()
    pathMetrics.EntityData.Leafs = types.NewOrderedMap()
    pathMetrics.EntityData.Leafs.Append("margin-relative", types.YLeaf{"MarginRelative", pathMetrics.MarginRelative})
    pathMetrics.EntityData.Leafs.Append("margin-absolute", types.YLeaf{"MarginAbsolute", pathMetrics.MarginAbsolute})
    pathMetrics.EntityData.Leafs.Append("maximum-segments", types.YLeaf{"MaximumSegments", pathMetrics.MaximumSegments})
    pathMetrics.EntityData.Leafs.Append("accumulative-te-metric", types.YLeaf{"AccumulativeTeMetric", pathMetrics.AccumulativeTeMetric})
    pathMetrics.EntityData.Leafs.Append("accumulative-igp-metric", types.YLeaf{"AccumulativeIgpMetric", pathMetrics.AccumulativeIgpMetric})
    pathMetrics.EntityData.Leafs.Append("accumulative-delay", types.YLeaf{"AccumulativeDelay", pathMetrics.AccumulativeDelay})

    pathMetrics.EntityData.YListKeys = []string {}

    return &(pathMetrics.EntityData)
}

// Xtc_Policies_Policy_Paths_SrPathConstraints_Segments
// Segments constraints
type Xtc_Policies_Policy_Paths_SrPathConstraints_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Algorithm. The type is interface{} with range: 0..255.
    SegmentAlgorithm interface{}
}

func (segments *Xtc_Policies_Policy_Paths_SrPathConstraints_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "sr-path-constraints"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Leafs = types.NewOrderedMap()
    segments.EntityData.Leafs.Append("segment-algorithm", types.YLeaf{"SegmentAlgorithm", segments.SegmentAlgorithm})

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint
// Affinity constraints list
type Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity type. The type is interface{} with range: 0..255.
    Type interface{}

    // Affinity value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Colors. The type is slice of
    // Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint_Color.
    Color []*Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint_Color
}

func (affinityConstraint *Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint) GetEntityData() *types.CommonEntityData {
    affinityConstraint.EntityData.YFilter = affinityConstraint.YFilter
    affinityConstraint.EntityData.YangName = "affinity-constraint"
    affinityConstraint.EntityData.BundleName = "cisco_ios_xr"
    affinityConstraint.EntityData.ParentYangName = "sr-path-constraints"
    affinityConstraint.EntityData.SegmentPath = "affinity-constraint"
    affinityConstraint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityConstraint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityConstraint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityConstraint.EntityData.Children = types.NewOrderedMap()
    affinityConstraint.EntityData.Children.Append("color", types.YChild{"Color", nil})
    for i := range affinityConstraint.Color {
        affinityConstraint.EntityData.Children.Append(types.GetSegmentPath(affinityConstraint.Color[i]), types.YChild{"Color", affinityConstraint.Color[i]})
    }
    affinityConstraint.EntityData.Leafs = types.NewOrderedMap()
    affinityConstraint.EntityData.Leafs.Append("type", types.YLeaf{"Type", affinityConstraint.Type})
    affinityConstraint.EntityData.Leafs.Append("value", types.YLeaf{"Value", affinityConstraint.Value})

    affinityConstraint.EntityData.YListKeys = []string {}

    return &(affinityConstraint.EntityData)
}

// Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint_Color
// Colors
type Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint_Color struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An affinity color. The type is string.
    Color interface{}
}

func (color *Xtc_Policies_Policy_Paths_SrPathConstraints_AffinityConstraint_Color) GetEntityData() *types.CommonEntityData {
    color.EntityData.YFilter = color.YFilter
    color.EntityData.YangName = "color"
    color.EntityData.BundleName = "cisco_ios_xr"
    color.EntityData.ParentYangName = "affinity-constraint"
    color.EntityData.SegmentPath = "color"
    color.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    color.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    color.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    color.EntityData.Children = types.NewOrderedMap()
    color.EntityData.Leafs = types.NewOrderedMap()
    color.EntityData.Leafs.Append("color", types.YLeaf{"Color", color.Color})

    color.EntityData.YListKeys = []string {}

    return &(color.EntityData)
}

// Xtc_Policies_Policy_Paths_Hops
// SR hop list
type Xtc_Policies_Policy_Paths_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is XtcSrSid.
    SidType interface{}

    // Algorithim. The type is interface{} with range: 0..255.
    Algorithm interface{}

    // SID value.
    Sid Xtc_Policies_Policy_Paths_Hops_Sid

    // Local address.
    LocalAddress Xtc_Policies_Policy_Paths_Hops_LocalAddress

    // Remote address.
    RemoteAddress Xtc_Policies_Policy_Paths_Hops_RemoteAddress
}

func (hops *Xtc_Policies_Policy_Paths_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "paths"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = types.NewOrderedMap()
    hops.EntityData.Children.Append("sid", types.YChild{"Sid", &hops.Sid})
    hops.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &hops.LocalAddress})
    hops.EntityData.Children.Append("remote-address", types.YChild{"RemoteAddress", &hops.RemoteAddress})
    hops.EntityData.Leafs = types.NewOrderedMap()
    hops.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", hops.SidType})
    hops.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", hops.Algorithm})

    hops.EntityData.YListKeys = []string {}

    return &(hops.EntityData)
}

// Xtc_Policies_Policy_Paths_Hops_Sid
// SID value
type Xtc_Policies_Policy_Paths_Hops_Sid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDType. The type is XtcSid.
    SidType interface{}

    // MPLS label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sid *Xtc_Policies_Policy_Paths_Hops_Sid) GetEntityData() *types.CommonEntityData {
    sid.EntityData.YFilter = sid.YFilter
    sid.EntityData.YangName = "sid"
    sid.EntityData.BundleName = "cisco_ios_xr"
    sid.EntityData.ParentYangName = "hops"
    sid.EntityData.SegmentPath = "sid"
    sid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sid.EntityData.Children = types.NewOrderedMap()
    sid.EntityData.Leafs = types.NewOrderedMap()
    sid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", sid.SidType})
    sid.EntityData.Leafs.Append("label", types.YLeaf{"Label", sid.Label})
    sid.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sid.Ipv6})

    sid.EntityData.YListKeys = []string {}

    return &(sid.EntityData)
}

// Xtc_Policies_Policy_Paths_Hops_LocalAddress
// Local address
type Xtc_Policies_Policy_Paths_Hops_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (localAddress *Xtc_Policies_Policy_Paths_Hops_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "hops"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddress.AfName})
    localAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddress.Ipv4})
    localAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddress.Ipv6})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// Xtc_Policies_Policy_Paths_Hops_RemoteAddress
// Remote address
type Xtc_Policies_Policy_Paths_Hops_RemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (remoteAddress *Xtc_Policies_Policy_Paths_Hops_RemoteAddress) GetEntityData() *types.CommonEntityData {
    remoteAddress.EntityData.YFilter = remoteAddress.YFilter
    remoteAddress.EntityData.YangName = "remote-address"
    remoteAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteAddress.EntityData.ParentYangName = "hops"
    remoteAddress.EntityData.SegmentPath = "remote-address"
    remoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddress.EntityData.Children = types.NewOrderedMap()
    remoteAddress.EntityData.Leafs = types.NewOrderedMap()
    remoteAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddress.AfName})
    remoteAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddress.Ipv4})
    remoteAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddress.Ipv6})

    remoteAddress.EntityData.YListKeys = []string {}

    return &(remoteAddress.EntityData)
}

// Xtc_PolicySummary
// Summary of all policies
type Xtc_PolicySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Total number of configured policies. The type is interface{} with range:
    // 0..4294967295.
    ConfiguredTotalPolicyCount interface{}

    // Total number of configured policies that are operationally up. The type is
    // interface{} with range: 0..4294967295.
    ConfiguredUpPolicyCount interface{}

    // Total number of configured policies that are operationally down. The type
    // is interface{} with range: 0..4294967295.
    ConfiguredDownPolicyCount interface{}

    // IPv4 address used for IPv4 policies.
    Ipv4SourceAddress Xtc_PolicySummary_Ipv4SourceAddress
}

func (policySummary *Xtc_PolicySummary) GetEntityData() *types.CommonEntityData {
    policySummary.EntityData.YFilter = policySummary.YFilter
    policySummary.EntityData.YangName = "policy-summary"
    policySummary.EntityData.BundleName = "cisco_ios_xr"
    policySummary.EntityData.ParentYangName = "xtc"
    policySummary.EntityData.SegmentPath = "policy-summary"
    policySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policySummary.EntityData.Children = types.NewOrderedMap()
    policySummary.EntityData.Children.Append("ipv4-source-address", types.YChild{"Ipv4SourceAddress", &policySummary.Ipv4SourceAddress})
    policySummary.EntityData.Leafs = types.NewOrderedMap()
    policySummary.EntityData.Leafs.Append("configured-total-policy-count", types.YLeaf{"ConfiguredTotalPolicyCount", policySummary.ConfiguredTotalPolicyCount})
    policySummary.EntityData.Leafs.Append("configured-up-policy-count", types.YLeaf{"ConfiguredUpPolicyCount", policySummary.ConfiguredUpPolicyCount})
    policySummary.EntityData.Leafs.Append("configured-down-policy-count", types.YLeaf{"ConfiguredDownPolicyCount", policySummary.ConfiguredDownPolicyCount})

    policySummary.EntityData.YListKeys = []string {}

    return &(policySummary.EntityData)
}

// Xtc_PolicySummary_Ipv4SourceAddress
// IPv4 address used for IPv4 policies
type Xtc_PolicySummary_Ipv4SourceAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipv4SourceAddress *Xtc_PolicySummary_Ipv4SourceAddress) GetEntityData() *types.CommonEntityData {
    ipv4SourceAddress.EntityData.YFilter = ipv4SourceAddress.YFilter
    ipv4SourceAddress.EntityData.YangName = "ipv4-source-address"
    ipv4SourceAddress.EntityData.BundleName = "cisco_ios_xr"
    ipv4SourceAddress.EntityData.ParentYangName = "policy-summary"
    ipv4SourceAddress.EntityData.SegmentPath = "ipv4-source-address"
    ipv4SourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4SourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4SourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4SourceAddress.EntityData.Children = types.NewOrderedMap()
    ipv4SourceAddress.EntityData.Leafs = types.NewOrderedMap()
    ipv4SourceAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", ipv4SourceAddress.AfName})
    ipv4SourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipv4SourceAddress.Ipv4})
    ipv4SourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipv4SourceAddress.Ipv6})

    ipv4SourceAddress.EntityData.YListKeys = []string {}

    return &(ipv4SourceAddress.EntityData)
}

// Xtc_OnDemandColors
// On-Demand Color database in XTC Agent
type Xtc_OnDemandColors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // On Demand Color information. The type is slice of
    // Xtc_OnDemandColors_OnDemandColor.
    OnDemandColor []*Xtc_OnDemandColors_OnDemandColor
}

func (onDemandColors *Xtc_OnDemandColors) GetEntityData() *types.CommonEntityData {
    onDemandColors.EntityData.YFilter = onDemandColors.YFilter
    onDemandColors.EntityData.YangName = "on-demand-colors"
    onDemandColors.EntityData.BundleName = "cisco_ios_xr"
    onDemandColors.EntityData.ParentYangName = "xtc"
    onDemandColors.EntityData.SegmentPath = "on-demand-colors"
    onDemandColors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColors.EntityData.Children = types.NewOrderedMap()
    onDemandColors.EntityData.Children.Append("on-demand-color", types.YChild{"OnDemandColor", nil})
    for i := range onDemandColors.OnDemandColor {
        onDemandColors.EntityData.Children.Append(types.GetSegmentPath(onDemandColors.OnDemandColor[i]), types.YChild{"OnDemandColor", onDemandColors.OnDemandColor[i]})
    }
    onDemandColors.EntityData.Leafs = types.NewOrderedMap()

    onDemandColors.EntityData.YListKeys = []string {}

    return &(onDemandColors.EntityData)
}

// Xtc_OnDemandColors_OnDemandColor
// On Demand Color information
type Xtc_OnDemandColors_OnDemandColor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Color. The type is interface{} with range:
    // 0..4294967295.
    Color interface{}

    // Color. The type is interface{} with range: 0..4294967295.
    ColorXr interface{}

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    MaximumSidDepth interface{}

    // Disjoint path information.
    DisjointPathInfo Xtc_OnDemandColors_OnDemandColor_DisjointPathInfo
}

func (onDemandColor *Xtc_OnDemandColors_OnDemandColor) GetEntityData() *types.CommonEntityData {
    onDemandColor.EntityData.YFilter = onDemandColor.YFilter
    onDemandColor.EntityData.YangName = "on-demand-color"
    onDemandColor.EntityData.BundleName = "cisco_ios_xr"
    onDemandColor.EntityData.ParentYangName = "on-demand-colors"
    onDemandColor.EntityData.SegmentPath = "on-demand-color" + types.AddKeyToken(onDemandColor.Color, "color")
    onDemandColor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onDemandColor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onDemandColor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onDemandColor.EntityData.Children = types.NewOrderedMap()
    onDemandColor.EntityData.Children.Append("disjoint-path-info", types.YChild{"DisjointPathInfo", &onDemandColor.DisjointPathInfo})
    onDemandColor.EntityData.Leafs = types.NewOrderedMap()
    onDemandColor.EntityData.Leafs.Append("color", types.YLeaf{"Color", onDemandColor.Color})
    onDemandColor.EntityData.Leafs.Append("color-xr", types.YLeaf{"ColorXr", onDemandColor.ColorXr})
    onDemandColor.EntityData.Leafs.Append("maximum-sid-depth", types.YLeaf{"MaximumSidDepth", onDemandColor.MaximumSidDepth})

    onDemandColor.EntityData.YListKeys = []string {"Color"}

    return &(onDemandColor.EntityData)
}

// Xtc_OnDemandColors_OnDemandColor_DisjointPathInfo
// Disjoint path information
type Xtc_OnDemandColors_OnDemandColor_DisjointPathInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disjointness type. The type is XtcDisjointness.
    DisjointnessType interface{}

    // Group ID. The type is interface{} with range: 0..4294967295.
    GroupId interface{}

    // Sub ID. The type is interface{} with range: 0..4294967295.
    SubId interface{}
}

func (disjointPathInfo *Xtc_OnDemandColors_OnDemandColor_DisjointPathInfo) GetEntityData() *types.CommonEntityData {
    disjointPathInfo.EntityData.YFilter = disjointPathInfo.YFilter
    disjointPathInfo.EntityData.YangName = "disjoint-path-info"
    disjointPathInfo.EntityData.BundleName = "cisco_ios_xr"
    disjointPathInfo.EntityData.ParentYangName = "on-demand-color"
    disjointPathInfo.EntityData.SegmentPath = "disjoint-path-info"
    disjointPathInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    disjointPathInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    disjointPathInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    disjointPathInfo.EntityData.Children = types.NewOrderedMap()
    disjointPathInfo.EntityData.Leafs = types.NewOrderedMap()
    disjointPathInfo.EntityData.Leafs.Append("disjointness-type", types.YLeaf{"DisjointnessType", disjointPathInfo.DisjointnessType})
    disjointPathInfo.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", disjointPathInfo.GroupId})
    disjointPathInfo.EntityData.Leafs.Append("sub-id", types.YLeaf{"SubId", disjointPathInfo.SubId})

    disjointPathInfo.EntityData.YListKeys = []string {}

    return &(disjointPathInfo.EntityData)
}

// Xtc_Forwarding
// Forwarding information
type Xtc_Forwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding information for policies.
    PolicyForwardings Xtc_Forwarding_PolicyForwardings
}

func (forwarding *Xtc_Forwarding) GetEntityData() *types.CommonEntityData {
    forwarding.EntityData.YFilter = forwarding.YFilter
    forwarding.EntityData.YangName = "forwarding"
    forwarding.EntityData.BundleName = "cisco_ios_xr"
    forwarding.EntityData.ParentYangName = "xtc"
    forwarding.EntityData.SegmentPath = "forwarding"
    forwarding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    forwarding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    forwarding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    forwarding.EntityData.Children = types.NewOrderedMap()
    forwarding.EntityData.Children.Append("policy-forwardings", types.YChild{"PolicyForwardings", &forwarding.PolicyForwardings})
    forwarding.EntityData.Leafs = types.NewOrderedMap()

    forwarding.EntityData.YListKeys = []string {}

    return &(forwarding.EntityData)
}

// Xtc_Forwarding_PolicyForwardings
// Forwarding information for policies
type Xtc_Forwarding_PolicyForwardings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Forwarding information for the policy. The type is slice of
    // Xtc_Forwarding_PolicyForwardings_PolicyForwarding.
    PolicyForwarding []*Xtc_Forwarding_PolicyForwardings_PolicyForwarding
}

func (policyForwardings *Xtc_Forwarding_PolicyForwardings) GetEntityData() *types.CommonEntityData {
    policyForwardings.EntityData.YFilter = policyForwardings.YFilter
    policyForwardings.EntityData.YangName = "policy-forwardings"
    policyForwardings.EntityData.BundleName = "cisco_ios_xr"
    policyForwardings.EntityData.ParentYangName = "forwarding"
    policyForwardings.EntityData.SegmentPath = "policy-forwardings"
    policyForwardings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyForwardings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyForwardings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyForwardings.EntityData.Children = types.NewOrderedMap()
    policyForwardings.EntityData.Children.Append("policy-forwarding", types.YChild{"PolicyForwarding", nil})
    for i := range policyForwardings.PolicyForwarding {
        policyForwardings.EntityData.Children.Append(types.GetSegmentPath(policyForwardings.PolicyForwarding[i]), types.YChild{"PolicyForwarding", policyForwardings.PolicyForwarding[i]})
    }
    policyForwardings.EntityData.Leafs = types.NewOrderedMap()

    policyForwardings.EntityData.YListKeys = []string {}

    return &(policyForwardings.EntityData)
}

// Xtc_Forwarding_PolicyForwardings_PolicyForwarding
// Forwarding information for the policy
type Xtc_Forwarding_PolicyForwardings_PolicyForwarding struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Policy Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Policy name. The type is string.
    PolicyName interface{}

    // Is local label valid and allocated?. The type is bool.
    IsLocalLabelValid interface{}

    // Local label for SR MPLS policy. The type is interface{} with range:
    // 0..4294967295.
    LocalLabel interface{}

    // Are policy stats valid?. The type is bool.
    AreStatsValid interface{}

    // Number of packets forwarded. The type is interface{} with range:
    // 0..18446744073709551615.
    ForwardingStatsPkts interface{}

    // Number of bytes forwarded. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ForwardingStatsBytes interface{}

    // Forwarding paths. The type is slice of
    // Xtc_Forwarding_PolicyForwardings_PolicyForwarding_Paths.
    Paths []*Xtc_Forwarding_PolicyForwardings_PolicyForwarding_Paths
}

func (policyForwarding *Xtc_Forwarding_PolicyForwardings_PolicyForwarding) GetEntityData() *types.CommonEntityData {
    policyForwarding.EntityData.YFilter = policyForwarding.YFilter
    policyForwarding.EntityData.YangName = "policy-forwarding"
    policyForwarding.EntityData.BundleName = "cisco_ios_xr"
    policyForwarding.EntityData.ParentYangName = "policy-forwardings"
    policyForwarding.EntityData.SegmentPath = "policy-forwarding" + types.AddKeyToken(policyForwarding.Name, "name")
    policyForwarding.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyForwarding.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyForwarding.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyForwarding.EntityData.Children = types.NewOrderedMap()
    policyForwarding.EntityData.Children.Append("paths", types.YChild{"Paths", nil})
    for i := range policyForwarding.Paths {
        policyForwarding.EntityData.Children.Append(types.GetSegmentPath(policyForwarding.Paths[i]), types.YChild{"Paths", policyForwarding.Paths[i]})
    }
    policyForwarding.EntityData.Leafs = types.NewOrderedMap()
    policyForwarding.EntityData.Leafs.Append("name", types.YLeaf{"Name", policyForwarding.Name})
    policyForwarding.EntityData.Leafs.Append("policy-name", types.YLeaf{"PolicyName", policyForwarding.PolicyName})
    policyForwarding.EntityData.Leafs.Append("is-local-label-valid", types.YLeaf{"IsLocalLabelValid", policyForwarding.IsLocalLabelValid})
    policyForwarding.EntityData.Leafs.Append("local-label", types.YLeaf{"LocalLabel", policyForwarding.LocalLabel})
    policyForwarding.EntityData.Leafs.Append("are-stats-valid", types.YLeaf{"AreStatsValid", policyForwarding.AreStatsValid})
    policyForwarding.EntityData.Leafs.Append("forwarding-stats-pkts", types.YLeaf{"ForwardingStatsPkts", policyForwarding.ForwardingStatsPkts})
    policyForwarding.EntityData.Leafs.Append("forwarding-stats-bytes", types.YLeaf{"ForwardingStatsBytes", policyForwarding.ForwardingStatsBytes})

    policyForwarding.EntityData.YListKeys = []string {"Name"}

    return &(policyForwarding.EntityData)
}

// Xtc_Forwarding_PolicyForwardings_PolicyForwarding_Paths
// Forwarding paths
type Xtc_Forwarding_PolicyForwardings_PolicyForwarding_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Outgoing interface handle. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    OutgoingInterface interface{}

    // IPv4 Next Hop. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextHopIpv4 interface{}

    // Table ID for nexthop address. The type is interface{} with range:
    // 0..4294967295.
    NextHopTableId interface{}

    // Is this path protected ?. The type is bool.
    IsProtected interface{}

    // Is this path a pure backup ?. The type is bool.
    IsPureBkup interface{}

    // Path's load metric for load balancing. The type is interface{} with range:
    // 0..4294967295.
    LoadMetric interface{}

    // path Id. The type is interface{} with range: 0..255.
    PathId interface{}

    // Backup path Id. The type is interface{} with range: 0..255.
    BkupPathId interface{}

    // Associated segment-list. The type is string.
    SegmentListName interface{}

    // Are per path stats valid?. The type is bool.
    AreStatsValid interface{}

    // Number of packets forwarded on this path. The type is interface{} with
    // range: 0..18446744073709551615.
    ForwardingStatsPkts interface{}

    // Number of bytes forwarded on this path. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ForwardingStatsBytes interface{}

    // Path outgoing labels. The type is slice of interface{} with range:
    // 0..4294967295.
    LabelStack []interface{}
}

func (paths *Xtc_Forwarding_PolicyForwardings_PolicyForwarding_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "policy-forwarding"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Leafs = types.NewOrderedMap()
    paths.EntityData.Leafs.Append("outgoing-interface", types.YLeaf{"OutgoingInterface", paths.OutgoingInterface})
    paths.EntityData.Leafs.Append("next-hop-ipv4", types.YLeaf{"NextHopIpv4", paths.NextHopIpv4})
    paths.EntityData.Leafs.Append("next-hop-table-id", types.YLeaf{"NextHopTableId", paths.NextHopTableId})
    paths.EntityData.Leafs.Append("is-protected", types.YLeaf{"IsProtected", paths.IsProtected})
    paths.EntityData.Leafs.Append("is-pure-bkup", types.YLeaf{"IsPureBkup", paths.IsPureBkup})
    paths.EntityData.Leafs.Append("load-metric", types.YLeaf{"LoadMetric", paths.LoadMetric})
    paths.EntityData.Leafs.Append("path-id", types.YLeaf{"PathId", paths.PathId})
    paths.EntityData.Leafs.Append("bkup-path-id", types.YLeaf{"BkupPathId", paths.BkupPathId})
    paths.EntityData.Leafs.Append("segment-list-name", types.YLeaf{"SegmentListName", paths.SegmentListName})
    paths.EntityData.Leafs.Append("are-stats-valid", types.YLeaf{"AreStatsValid", paths.AreStatsValid})
    paths.EntityData.Leafs.Append("forwarding-stats-pkts", types.YLeaf{"ForwardingStatsPkts", paths.ForwardingStatsPkts})
    paths.EntityData.Leafs.Append("forwarding-stats-bytes", types.YLeaf{"ForwardingStatsBytes", paths.ForwardingStatsBytes})
    paths.EntityData.Leafs.Append("label-stack", types.YLeaf{"LabelStack", paths.LabelStack})

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// Xtc_Controller
// Controller information
type Xtc_Controller struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table containing policy requests.
    PolicyRequests Xtc_Controller_PolicyRequests
}

func (controller *Xtc_Controller) GetEntityData() *types.CommonEntityData {
    controller.EntityData.YFilter = controller.YFilter
    controller.EntityData.YangName = "controller"
    controller.EntityData.BundleName = "cisco_ios_xr"
    controller.EntityData.ParentYangName = "xtc"
    controller.EntityData.SegmentPath = "controller"
    controller.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controller.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controller.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controller.EntityData.Children = types.NewOrderedMap()
    controller.EntityData.Children.Append("policy-requests", types.YChild{"PolicyRequests", &controller.PolicyRequests})
    controller.EntityData.Leafs = types.NewOrderedMap()

    controller.EntityData.YListKeys = []string {}

    return &(controller.EntityData)
}

// Xtc_Controller_PolicyRequests
// Table containing policy requests
type Xtc_Controller_PolicyRequests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Policy request information. The type is slice of
    // Xtc_Controller_PolicyRequests_PolicyRequest.
    PolicyRequest []*Xtc_Controller_PolicyRequests_PolicyRequest
}

func (policyRequests *Xtc_Controller_PolicyRequests) GetEntityData() *types.CommonEntityData {
    policyRequests.EntityData.YFilter = policyRequests.YFilter
    policyRequests.EntityData.YangName = "policy-requests"
    policyRequests.EntityData.BundleName = "cisco_ios_xr"
    policyRequests.EntityData.ParentYangName = "controller"
    policyRequests.EntityData.SegmentPath = "policy-requests"
    policyRequests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyRequests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyRequests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyRequests.EntityData.Children = types.NewOrderedMap()
    policyRequests.EntityData.Children.Append("policy-request", types.YChild{"PolicyRequest", nil})
    for i := range policyRequests.PolicyRequest {
        policyRequests.EntityData.Children.Append(types.GetSegmentPath(policyRequests.PolicyRequest[i]), types.YChild{"PolicyRequest", policyRequests.PolicyRequest[i]})
    }
    policyRequests.EntityData.Leafs = types.NewOrderedMap()

    policyRequests.EntityData.YListKeys = []string {}

    return &(policyRequests.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest
// Policy request information
type Xtc_Controller_PolicyRequests_PolicyRequest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Source Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // This attribute is a key. Endpoint Address Type. The type is
    // XtcAddressFamily.
    EndPointType interface{}

    // This attribute is a key. Endpoint Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    EndPointAddress interface{}

    // This attribute is a key. Color. The type is interface{} with range:
    // 1..4294967295.
    Color interface{}

    // This attribute is a key. Route Distinguisher. The type is interface{} with
    // range: 0..4294967295.
    RouteDistinguisher interface{}

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddressXr interface{}

    // Binding SID. The type is interface{} with range: 0..4294967295.
    BindingSid interface{}

    // Preference. The type is interface{} with range: 0..4294967295.
    Preference interface{}

    // Color. The type is interface{} with range: 0..4294967295.
    ColorXr interface{}

    // Route distinguisher. The type is interface{} with range: 0..4294967295.
    RouteDistinguisherXr interface{}

    // Creation time of the request in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    CreationTime interface{}

    // Last updated time of the request in seconds. The type is interface{} with
    // range: 0..4294967295. Units are second.
    LastUpdatedTime interface{}

    // End point.
    EndPoint Xtc_Controller_PolicyRequests_PolicyRequest_EndPoint

    // Path options. The type is slice of
    // Xtc_Controller_PolicyRequests_PolicyRequest_Paths.
    Paths []*Xtc_Controller_PolicyRequests_PolicyRequest_Paths
}

func (policyRequest *Xtc_Controller_PolicyRequests_PolicyRequest) GetEntityData() *types.CommonEntityData {
    policyRequest.EntityData.YFilter = policyRequest.YFilter
    policyRequest.EntityData.YangName = "policy-request"
    policyRequest.EntityData.BundleName = "cisco_ios_xr"
    policyRequest.EntityData.ParentYangName = "policy-requests"
    policyRequest.EntityData.SegmentPath = "policy-request" + types.AddKeyToken(policyRequest.SourceAddress, "source-address") + types.AddKeyToken(policyRequest.EndPointType, "end-point-type") + types.AddKeyToken(policyRequest.EndPointAddress, "end-point-address") + types.AddKeyToken(policyRequest.Color, "color") + types.AddKeyToken(policyRequest.RouteDistinguisher, "route-distinguisher")
    policyRequest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyRequest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyRequest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyRequest.EntityData.Children = types.NewOrderedMap()
    policyRequest.EntityData.Children.Append("end-point", types.YChild{"EndPoint", &policyRequest.EndPoint})
    policyRequest.EntityData.Children.Append("paths", types.YChild{"Paths", nil})
    for i := range policyRequest.Paths {
        policyRequest.EntityData.Children.Append(types.GetSegmentPath(policyRequest.Paths[i]), types.YChild{"Paths", policyRequest.Paths[i]})
    }
    policyRequest.EntityData.Leafs = types.NewOrderedMap()
    policyRequest.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", policyRequest.SourceAddress})
    policyRequest.EntityData.Leafs.Append("end-point-type", types.YLeaf{"EndPointType", policyRequest.EndPointType})
    policyRequest.EntityData.Leafs.Append("end-point-address", types.YLeaf{"EndPointAddress", policyRequest.EndPointAddress})
    policyRequest.EntityData.Leafs.Append("color", types.YLeaf{"Color", policyRequest.Color})
    policyRequest.EntityData.Leafs.Append("route-distinguisher", types.YLeaf{"RouteDistinguisher", policyRequest.RouteDistinguisher})
    policyRequest.EntityData.Leafs.Append("source-address-xr", types.YLeaf{"SourceAddressXr", policyRequest.SourceAddressXr})
    policyRequest.EntityData.Leafs.Append("binding-sid", types.YLeaf{"BindingSid", policyRequest.BindingSid})
    policyRequest.EntityData.Leafs.Append("preference", types.YLeaf{"Preference", policyRequest.Preference})
    policyRequest.EntityData.Leafs.Append("color-xr", types.YLeaf{"ColorXr", policyRequest.ColorXr})
    policyRequest.EntityData.Leafs.Append("route-distinguisher-xr", types.YLeaf{"RouteDistinguisherXr", policyRequest.RouteDistinguisherXr})
    policyRequest.EntityData.Leafs.Append("creation-time", types.YLeaf{"CreationTime", policyRequest.CreationTime})
    policyRequest.EntityData.Leafs.Append("last-updated-time", types.YLeaf{"LastUpdatedTime", policyRequest.LastUpdatedTime})

    policyRequest.EntityData.YListKeys = []string {"SourceAddress", "EndPointType", "EndPointAddress", "Color", "RouteDistinguisher"}

    return &(policyRequest.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_EndPoint
// End point
type Xtc_Controller_PolicyRequests_PolicyRequest_EndPoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (endPoint *Xtc_Controller_PolicyRequests_PolicyRequest_EndPoint) GetEntityData() *types.CommonEntityData {
    endPoint.EntityData.YFilter = endPoint.YFilter
    endPoint.EntityData.YangName = "end-point"
    endPoint.EntityData.BundleName = "cisco_ios_xr"
    endPoint.EntityData.ParentYangName = "policy-request"
    endPoint.EntityData.SegmentPath = "end-point"
    endPoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    endPoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    endPoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    endPoint.EntityData.Children = types.NewOrderedMap()
    endPoint.EntityData.Leafs = types.NewOrderedMap()
    endPoint.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", endPoint.AfName})
    endPoint.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", endPoint.Ipv4})
    endPoint.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", endPoint.Ipv6})

    endPoint.EntityData.YListKeys = []string {}

    return &(endPoint.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths
// Path options
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Index number. The type is interface{} with range: 0..4294967295.
    Index interface{}

    // Path option type. The type is XtcPolicyPath.
    Type interface{}

    // Explicit path option name. The type is string.
    Name interface{}

    // Whether the path is active (used). The type is bool.
    Active interface{}

    // Configured weight of the path-option. The type is interface{} with range:
    // 0..4294967295.
    Weight interface{}

    // Configured path metric type. The type is interface{} with range: 0..255.
    MetricType interface{}

    // Path metric value. The type is interface{} with range: 0..4294967295.
    MetricValue interface{}

    // True if path is valid. The type is bool.
    IsValid interface{}

    // True if the path is to be computed by PCE. The type is bool.
    PceBasedPath interface{}

    // Address of the PCE computed the path. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PceAddress interface{}

    // Error (for display only). The type is string.
    Error interface{}

    // SR path constraints.
    SrPathConstraints Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints

    // SR hop list. The type is slice of
    // Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops.
    Hops []*Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops
}

func (paths *Xtc_Controller_PolicyRequests_PolicyRequest_Paths) GetEntityData() *types.CommonEntityData {
    paths.EntityData.YFilter = paths.YFilter
    paths.EntityData.YangName = "paths"
    paths.EntityData.BundleName = "cisco_ios_xr"
    paths.EntityData.ParentYangName = "policy-request"
    paths.EntityData.SegmentPath = "paths"
    paths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    paths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    paths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    paths.EntityData.Children = types.NewOrderedMap()
    paths.EntityData.Children.Append("sr-path-constraints", types.YChild{"SrPathConstraints", &paths.SrPathConstraints})
    paths.EntityData.Children.Append("hops", types.YChild{"Hops", nil})
    for i := range paths.Hops {
        paths.EntityData.Children.Append(types.GetSegmentPath(paths.Hops[i]), types.YChild{"Hops", paths.Hops[i]})
    }
    paths.EntityData.Leafs = types.NewOrderedMap()
    paths.EntityData.Leafs.Append("index", types.YLeaf{"Index", paths.Index})
    paths.EntityData.Leafs.Append("type", types.YLeaf{"Type", paths.Type})
    paths.EntityData.Leafs.Append("name", types.YLeaf{"Name", paths.Name})
    paths.EntityData.Leafs.Append("active", types.YLeaf{"Active", paths.Active})
    paths.EntityData.Leafs.Append("weight", types.YLeaf{"Weight", paths.Weight})
    paths.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", paths.MetricType})
    paths.EntityData.Leafs.Append("metric-value", types.YLeaf{"MetricValue", paths.MetricValue})
    paths.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", paths.IsValid})
    paths.EntityData.Leafs.Append("pce-based-path", types.YLeaf{"PceBasedPath", paths.PceBasedPath})
    paths.EntityData.Leafs.Append("pce-address", types.YLeaf{"PceAddress", paths.PceAddress})
    paths.EntityData.Leafs.Append("error", types.YLeaf{"Error", paths.Error})

    paths.EntityData.YListKeys = []string {}

    return &(paths.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints
// SR path constraints
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Path metrics.
    PathMetrics Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_PathMetrics

    // Segments constraints.
    Segments Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_Segments

    // Affinity constraints list. The type is slice of
    // Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint.
    AffinityConstraint []*Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint
}

func (srPathConstraints *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints) GetEntityData() *types.CommonEntityData {
    srPathConstraints.EntityData.YFilter = srPathConstraints.YFilter
    srPathConstraints.EntityData.YangName = "sr-path-constraints"
    srPathConstraints.EntityData.BundleName = "cisco_ios_xr"
    srPathConstraints.EntityData.ParentYangName = "paths"
    srPathConstraints.EntityData.SegmentPath = "sr-path-constraints"
    srPathConstraints.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srPathConstraints.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srPathConstraints.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srPathConstraints.EntityData.Children = types.NewOrderedMap()
    srPathConstraints.EntityData.Children.Append("path-metrics", types.YChild{"PathMetrics", &srPathConstraints.PathMetrics})
    srPathConstraints.EntityData.Children.Append("segments", types.YChild{"Segments", &srPathConstraints.Segments})
    srPathConstraints.EntityData.Children.Append("affinity-constraint", types.YChild{"AffinityConstraint", nil})
    for i := range srPathConstraints.AffinityConstraint {
        srPathConstraints.EntityData.Children.Append(types.GetSegmentPath(srPathConstraints.AffinityConstraint[i]), types.YChild{"AffinityConstraint", srPathConstraints.AffinityConstraint[i]})
    }
    srPathConstraints.EntityData.Leafs = types.NewOrderedMap()

    srPathConstraints.EntityData.YListKeys = []string {}

    return &(srPathConstraints.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_PathMetrics
// Path metrics
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_PathMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Margin Relative. The type is interface{} with range: 0..255.
    MarginRelative interface{}

    // Margin Absolute. The type is interface{} with range: 0..255.
    MarginAbsolute interface{}

    // Maximum number of segments. The type is interface{} with range: 0..65535.
    MaximumSegments interface{}

    // Accumulative TE metric. The type is interface{} with range: 0..4294967295.
    AccumulativeTeMetric interface{}

    // Accumulative IGP metric. The type is interface{} with range: 0..4294967295.
    AccumulativeIgpMetric interface{}

    // Accumulative delay. The type is interface{} with range: 0..4294967295.
    AccumulativeDelay interface{}
}

func (pathMetrics *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_PathMetrics) GetEntityData() *types.CommonEntityData {
    pathMetrics.EntityData.YFilter = pathMetrics.YFilter
    pathMetrics.EntityData.YangName = "path-metrics"
    pathMetrics.EntityData.BundleName = "cisco_ios_xr"
    pathMetrics.EntityData.ParentYangName = "sr-path-constraints"
    pathMetrics.EntityData.SegmentPath = "path-metrics"
    pathMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pathMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pathMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pathMetrics.EntityData.Children = types.NewOrderedMap()
    pathMetrics.EntityData.Leafs = types.NewOrderedMap()
    pathMetrics.EntityData.Leafs.Append("margin-relative", types.YLeaf{"MarginRelative", pathMetrics.MarginRelative})
    pathMetrics.EntityData.Leafs.Append("margin-absolute", types.YLeaf{"MarginAbsolute", pathMetrics.MarginAbsolute})
    pathMetrics.EntityData.Leafs.Append("maximum-segments", types.YLeaf{"MaximumSegments", pathMetrics.MaximumSegments})
    pathMetrics.EntityData.Leafs.Append("accumulative-te-metric", types.YLeaf{"AccumulativeTeMetric", pathMetrics.AccumulativeTeMetric})
    pathMetrics.EntityData.Leafs.Append("accumulative-igp-metric", types.YLeaf{"AccumulativeIgpMetric", pathMetrics.AccumulativeIgpMetric})
    pathMetrics.EntityData.Leafs.Append("accumulative-delay", types.YLeaf{"AccumulativeDelay", pathMetrics.AccumulativeDelay})

    pathMetrics.EntityData.YListKeys = []string {}

    return &(pathMetrics.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_Segments
// Segments constraints
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_Segments struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Segment Algorithm. The type is interface{} with range: 0..255.
    SegmentAlgorithm interface{}
}

func (segments *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_Segments) GetEntityData() *types.CommonEntityData {
    segments.EntityData.YFilter = segments.YFilter
    segments.EntityData.YangName = "segments"
    segments.EntityData.BundleName = "cisco_ios_xr"
    segments.EntityData.ParentYangName = "sr-path-constraints"
    segments.EntityData.SegmentPath = "segments"
    segments.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    segments.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    segments.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    segments.EntityData.Children = types.NewOrderedMap()
    segments.EntityData.Leafs = types.NewOrderedMap()
    segments.EntityData.Leafs.Append("segment-algorithm", types.YLeaf{"SegmentAlgorithm", segments.SegmentAlgorithm})

    segments.EntityData.YListKeys = []string {}

    return &(segments.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint
// Affinity constraints list
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Affinity type. The type is interface{} with range: 0..255.
    Type interface{}

    // Affinity value. The type is interface{} with range: 0..4294967295.
    Value interface{}

    // Colors. The type is slice of
    // Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint_Color.
    Color []*Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint_Color
}

func (affinityConstraint *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint) GetEntityData() *types.CommonEntityData {
    affinityConstraint.EntityData.YFilter = affinityConstraint.YFilter
    affinityConstraint.EntityData.YangName = "affinity-constraint"
    affinityConstraint.EntityData.BundleName = "cisco_ios_xr"
    affinityConstraint.EntityData.ParentYangName = "sr-path-constraints"
    affinityConstraint.EntityData.SegmentPath = "affinity-constraint"
    affinityConstraint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    affinityConstraint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    affinityConstraint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    affinityConstraint.EntityData.Children = types.NewOrderedMap()
    affinityConstraint.EntityData.Children.Append("color", types.YChild{"Color", nil})
    for i := range affinityConstraint.Color {
        affinityConstraint.EntityData.Children.Append(types.GetSegmentPath(affinityConstraint.Color[i]), types.YChild{"Color", affinityConstraint.Color[i]})
    }
    affinityConstraint.EntityData.Leafs = types.NewOrderedMap()
    affinityConstraint.EntityData.Leafs.Append("type", types.YLeaf{"Type", affinityConstraint.Type})
    affinityConstraint.EntityData.Leafs.Append("value", types.YLeaf{"Value", affinityConstraint.Value})

    affinityConstraint.EntityData.YListKeys = []string {}

    return &(affinityConstraint.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint_Color
// Colors
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint_Color struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An affinity color. The type is string.
    Color interface{}
}

func (color *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_SrPathConstraints_AffinityConstraint_Color) GetEntityData() *types.CommonEntityData {
    color.EntityData.YFilter = color.YFilter
    color.EntityData.YangName = "color"
    color.EntityData.BundleName = "cisco_ios_xr"
    color.EntityData.ParentYangName = "affinity-constraint"
    color.EntityData.SegmentPath = "color"
    color.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    color.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    color.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    color.EntityData.Children = types.NewOrderedMap()
    color.EntityData.Leafs = types.NewOrderedMap()
    color.EntityData.Leafs.Append("color", types.YLeaf{"Color", color.Color})

    color.EntityData.YListKeys = []string {}

    return &(color.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops
// SR hop list
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is XtcSrSid.
    SidType interface{}

    // Algorithim. The type is interface{} with range: 0..255.
    Algorithm interface{}

    // SID value.
    Sid Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_Sid

    // Local address.
    LocalAddress Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_LocalAddress

    // Remote address.
    RemoteAddress Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_RemoteAddress
}

func (hops *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "paths"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = types.NewOrderedMap()
    hops.EntityData.Children.Append("sid", types.YChild{"Sid", &hops.Sid})
    hops.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &hops.LocalAddress})
    hops.EntityData.Children.Append("remote-address", types.YChild{"RemoteAddress", &hops.RemoteAddress})
    hops.EntityData.Leafs = types.NewOrderedMap()
    hops.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", hops.SidType})
    hops.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", hops.Algorithm})

    hops.EntityData.YListKeys = []string {}

    return &(hops.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_Sid
// SID value
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_Sid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SIDType. The type is XtcSid.
    SidType interface{}

    // MPLS label. The type is interface{} with range: 0..4294967295.
    Label interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sid *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_Sid) GetEntityData() *types.CommonEntityData {
    sid.EntityData.YFilter = sid.YFilter
    sid.EntityData.YangName = "sid"
    sid.EntityData.BundleName = "cisco_ios_xr"
    sid.EntityData.ParentYangName = "hops"
    sid.EntityData.SegmentPath = "sid"
    sid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sid.EntityData.Children = types.NewOrderedMap()
    sid.EntityData.Leafs = types.NewOrderedMap()
    sid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", sid.SidType})
    sid.EntityData.Leafs.Append("label", types.YLeaf{"Label", sid.Label})
    sid.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sid.Ipv6})

    sid.EntityData.YListKeys = []string {}

    return &(sid.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_LocalAddress
// Local address
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (localAddress *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "hops"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddress.AfName})
    localAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddress.Ipv4})
    localAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddress.Ipv6})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_RemoteAddress
// Remote address
type Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_RemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (remoteAddress *Xtc_Controller_PolicyRequests_PolicyRequest_Paths_Hops_RemoteAddress) GetEntityData() *types.CommonEntityData {
    remoteAddress.EntityData.YFilter = remoteAddress.YFilter
    remoteAddress.EntityData.YangName = "remote-address"
    remoteAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteAddress.EntityData.ParentYangName = "hops"
    remoteAddress.EntityData.SegmentPath = "remote-address"
    remoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddress.EntityData.Children = types.NewOrderedMap()
    remoteAddress.EntityData.Leafs = types.NewOrderedMap()
    remoteAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddress.AfName})
    remoteAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddress.Ipv4})
    remoteAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddress.Ipv6})

    remoteAddress.EntityData.YListKeys = []string {}

    return &(remoteAddress.EntityData)
}

// Xtc_TopologySummary
// Node summary database
type Xtc_TopologySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    Prefixes interface{}

    // Number of prefix SIDs. The type is interface{} with range: 0..4294967295.
    PrefixSids interface{}

    // Number of links. The type is interface{} with range: 0..4294967295.
    Links interface{}

    // Number of adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    AdjacencySids interface{}
}

func (topologySummary *Xtc_TopologySummary) GetEntityData() *types.CommonEntityData {
    topologySummary.EntityData.YFilter = topologySummary.YFilter
    topologySummary.EntityData.YangName = "topology-summary"
    topologySummary.EntityData.BundleName = "cisco_ios_xr"
    topologySummary.EntityData.ParentYangName = "xtc"
    topologySummary.EntityData.SegmentPath = "topology-summary"
    topologySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologySummary.EntityData.Children = types.NewOrderedMap()
    topologySummary.EntityData.Leafs = types.NewOrderedMap()
    topologySummary.EntityData.Leafs.Append("nodes", types.YLeaf{"Nodes", topologySummary.Nodes})
    topologySummary.EntityData.Leafs.Append("prefixes", types.YLeaf{"Prefixes", topologySummary.Prefixes})
    topologySummary.EntityData.Leafs.Append("prefix-sids", types.YLeaf{"PrefixSids", topologySummary.PrefixSids})
    topologySummary.EntityData.Leafs.Append("links", types.YLeaf{"Links", topologySummary.Links})
    topologySummary.EntityData.Leafs.Append("adjacency-sids", types.YLeaf{"AdjacencySids", topologySummary.AdjacencySids})

    topologySummary.EntityData.YListKeys = []string {}

    return &(topologySummary.EntityData)
}

// Xtc_TopologyNodes
// Node database in XTC Agent
type Xtc_TopologyNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node information. The type is slice of Xtc_TopologyNodes_TopologyNode.
    TopologyNode []*Xtc_TopologyNodes_TopologyNode
}

func (topologyNodes *Xtc_TopologyNodes) GetEntityData() *types.CommonEntityData {
    topologyNodes.EntityData.YFilter = topologyNodes.YFilter
    topologyNodes.EntityData.YangName = "topology-nodes"
    topologyNodes.EntityData.BundleName = "cisco_ios_xr"
    topologyNodes.EntityData.ParentYangName = "xtc"
    topologyNodes.EntityData.SegmentPath = "topology-nodes"
    topologyNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNodes.EntityData.Children = types.NewOrderedMap()
    topologyNodes.EntityData.Children.Append("topology-node", types.YChild{"TopologyNode", nil})
    for i := range topologyNodes.TopologyNode {
        topologyNodes.EntityData.Children.Append(types.GetSegmentPath(topologyNodes.TopologyNode[i]), types.YChild{"TopologyNode", topologyNodes.TopologyNode[i]})
    }
    topologyNodes.EntityData.Leafs = types.NewOrderedMap()

    topologyNodes.EntityData.YListKeys = []string {}

    return &(topologyNodes.EntityData)
}

// Xtc_TopologyNodes_TopologyNode
// Node information
type Xtc_TopologyNodes_TopologyNode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node Identifier. The type is interface{} with
    // range: 0..4294967295.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node Overload Bit. The type is bool.
    Overload interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier

    // Prefix SIDs. The type is slice of Xtc_TopologyNodes_TopologyNode_PrefixSid.
    PrefixSid []*Xtc_TopologyNodes_TopologyNode_PrefixSid

    // IPv4 Link information. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_Ipv4Link.
    Ipv4Link []*Xtc_TopologyNodes_TopologyNode_Ipv4Link

    // IPv6 Link information. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_Ipv6Link.
    Ipv6Link []*Xtc_TopologyNodes_TopologyNode_Ipv6Link
}

func (topologyNode *Xtc_TopologyNodes_TopologyNode) GetEntityData() *types.CommonEntityData {
    topologyNode.EntityData.YFilter = topologyNode.YFilter
    topologyNode.EntityData.YangName = "topology-node"
    topologyNode.EntityData.BundleName = "cisco_ios_xr"
    topologyNode.EntityData.ParentYangName = "topology-nodes"
    topologyNode.EntityData.SegmentPath = "topology-node" + types.AddKeyToken(topologyNode.NodeIdentifier, "node-identifier")
    topologyNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNode.EntityData.Children = types.NewOrderedMap()
    topologyNode.EntityData.Children.Append("node-protocol-identifier", types.YChild{"NodeProtocolIdentifier", &topologyNode.NodeProtocolIdentifier})
    topologyNode.EntityData.Children.Append("prefix-sid", types.YChild{"PrefixSid", nil})
    for i := range topologyNode.PrefixSid {
        topologyNode.EntityData.Children.Append(types.GetSegmentPath(topologyNode.PrefixSid[i]), types.YChild{"PrefixSid", topologyNode.PrefixSid[i]})
    }
    topologyNode.EntityData.Children.Append("ipv4-link", types.YChild{"Ipv4Link", nil})
    for i := range topologyNode.Ipv4Link {
        topologyNode.EntityData.Children.Append(types.GetSegmentPath(topologyNode.Ipv4Link[i]), types.YChild{"Ipv4Link", topologyNode.Ipv4Link[i]})
    }
    topologyNode.EntityData.Children.Append("ipv6-link", types.YChild{"Ipv6Link", nil})
    for i := range topologyNode.Ipv6Link {
        topologyNode.EntityData.Children.Append(types.GetSegmentPath(topologyNode.Ipv6Link[i]), types.YChild{"Ipv6Link", topologyNode.Ipv6Link[i]})
    }
    topologyNode.EntityData.Leafs = types.NewOrderedMap()
    topologyNode.EntityData.Leafs.Append("node-identifier", types.YLeaf{"NodeIdentifier", topologyNode.NodeIdentifier})
    topologyNode.EntityData.Leafs.Append("node-identifier-xr", types.YLeaf{"NodeIdentifierXr", topologyNode.NodeIdentifierXr})
    topologyNode.EntityData.Leafs.Append("overload", types.YLeaf{"Overload", topologyNode.Overload})

    topologyNode.EntityData.YListKeys = []string {"NodeIdentifier"}

    return &(topologyNode.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier
// Node protocol identifier
type Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
}

func (nodeProtocolIdentifier *Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    nodeProtocolIdentifier.EntityData.YFilter = nodeProtocolIdentifier.YFilter
    nodeProtocolIdentifier.EntityData.YangName = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    nodeProtocolIdentifier.EntityData.ParentYangName = "topology-node"
    nodeProtocolIdentifier.EntityData.SegmentPath = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", nodeProtocolIdentifier.IgpInformation[i]})
    }
    nodeProtocolIdentifier.EntityData.Leafs = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodeProtocolIdentifier.NodeName})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id-set", types.YLeaf{"Ipv4BgpRouterIdSet", nodeProtocolIdentifier.Ipv4BgpRouterIdSet})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id", types.YLeaf{"Ipv4BgpRouterId", nodeProtocolIdentifier.Ipv4BgpRouterId})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id-set", types.YLeaf{"Ipv4teRouterIdSet", nodeProtocolIdentifier.Ipv4teRouterIdSet})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id", types.YLeaf{"Ipv4teRouterId", nodeProtocolIdentifier.Ipv4teRouterId})

    nodeProtocolIdentifier.EntityData.YListKeys = []string {}

    return &(nodeProtocolIdentifier.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // IGP-specific information.
    Igp Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("igp", types.YChild{"Igp", &igpInformation.Igp})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is XtcIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("isis", types.YChild{"Isis", &igp.Isis})
    igp.EntityData.Children.Append("ospf", types.YChild{"Ospf", &igp.Ospf})
    igp.EntityData.Children.Append("bgp", types.YChild{"Bgp", &igp.Bgp})
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("igp-id", types.YLeaf{"IgpId", igp.IgpId})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", isis.SystemId})
    isis.EntityData.Leafs.Append("level", types.YLeaf{"Level", isis.Level})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospf.RouterId})
    ospf.EntityData.Leafs.Append("area", types.YLeaf{"Area", ospf.Area})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (bgp *Xtc_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", bgp.RouterId})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_PrefixSid
// Prefix SIDs
type Xtc_TopologyNodes_TopologyNode_PrefixSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Type. The type is XtcSid1.
    SidType interface{}

    // Prefix-SID algorithm number. The type is interface{} with range:
    // 0..4294967295.
    Algorithm interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Prefix.
    SidPrefix Xtc_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
}

func (prefixSid *Xtc_TopologyNodes_TopologyNode_PrefixSid) GetEntityData() *types.CommonEntityData {
    prefixSid.EntityData.YFilter = prefixSid.YFilter
    prefixSid.EntityData.YangName = "prefix-sid"
    prefixSid.EntityData.BundleName = "cisco_ios_xr"
    prefixSid.EntityData.ParentYangName = "topology-node"
    prefixSid.EntityData.SegmentPath = "prefix-sid"
    prefixSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSid.EntityData.Children = types.NewOrderedMap()
    prefixSid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &prefixSid.SidPrefix})
    prefixSid.EntityData.Leafs = types.NewOrderedMap()
    prefixSid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", prefixSid.SidType})
    prefixSid.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", prefixSid.Algorithm})
    prefixSid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", prefixSid.MplsLabel})

    prefixSid.EntityData.YListKeys = []string {}

    return &(prefixSid.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
// Prefix
type Xtc_TopologyNodes_TopologyNode_PrefixSid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *Xtc_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "prefix-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalIpv4Address interface{}

    // Remote IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteIpv4Address interface{}

    // IGP Metric. The type is interface{} with range: 0..4294967295.
    IgpMetric interface{}

    // TE Metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Max link bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    MaximumLinkBandwidth interface{}

    // Max Reservable bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxReservableBandwidth interface{}

    // Link admin-groups. The type is interface{} with range: 0..4294967295.
    AdministrativeGroups interface{}

    // SRLG Values. The type is slice of interface{} with range: 0..4294967295.
    Srlgs []interface{}

    // Local node IGP information.
    LocalIgpInformation Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid.
    AdjacencySid []*Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
}

func (ipv4Link *Xtc_TopologyNodes_TopologyNode_Ipv4Link) GetEntityData() *types.CommonEntityData {
    ipv4Link.EntityData.YFilter = ipv4Link.YFilter
    ipv4Link.EntityData.YangName = "ipv4-link"
    ipv4Link.EntityData.BundleName = "cisco_ios_xr"
    ipv4Link.EntityData.ParentYangName = "topology-node"
    ipv4Link.EntityData.SegmentPath = "ipv4-link"
    ipv4Link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Link.EntityData.Children = types.NewOrderedMap()
    ipv4Link.EntityData.Children.Append("local-igp-information", types.YChild{"LocalIgpInformation", &ipv4Link.LocalIgpInformation})
    ipv4Link.EntityData.Children.Append("remote-node-protocol-identifier", types.YChild{"RemoteNodeProtocolIdentifier", &ipv4Link.RemoteNodeProtocolIdentifier})
    ipv4Link.EntityData.Children.Append("adjacency-sid", types.YChild{"AdjacencySid", nil})
    for i := range ipv4Link.AdjacencySid {
        ipv4Link.EntityData.Children.Append(types.GetSegmentPath(ipv4Link.AdjacencySid[i]), types.YChild{"AdjacencySid", ipv4Link.AdjacencySid[i]})
    }
    ipv4Link.EntityData.Leafs = types.NewOrderedMap()
    ipv4Link.EntityData.Leafs.Append("local-ipv4-address", types.YLeaf{"LocalIpv4Address", ipv4Link.LocalIpv4Address})
    ipv4Link.EntityData.Leafs.Append("remote-ipv4-address", types.YLeaf{"RemoteIpv4Address", ipv4Link.RemoteIpv4Address})
    ipv4Link.EntityData.Leafs.Append("igp-metric", types.YLeaf{"IgpMetric", ipv4Link.IgpMetric})
    ipv4Link.EntityData.Leafs.Append("te-metric", types.YLeaf{"TeMetric", ipv4Link.TeMetric})
    ipv4Link.EntityData.Leafs.Append("maximum-link-bandwidth", types.YLeaf{"MaximumLinkBandwidth", ipv4Link.MaximumLinkBandwidth})
    ipv4Link.EntityData.Leafs.Append("max-reservable-bandwidth", types.YLeaf{"MaxReservableBandwidth", ipv4Link.MaxReservableBandwidth})
    ipv4Link.EntityData.Leafs.Append("administrative-groups", types.YLeaf{"AdministrativeGroups", ipv4Link.AdministrativeGroups})
    ipv4Link.EntityData.Leafs.Append("srlgs", types.YLeaf{"Srlgs", ipv4Link.Srlgs})

    ipv4Link.EntityData.YListKeys = []string {}

    return &(ipv4Link.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // IGP-specific information.
    Igp Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
}

func (localIgpInformation *Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetEntityData() *types.CommonEntityData {
    localIgpInformation.EntityData.YFilter = localIgpInformation.YFilter
    localIgpInformation.EntityData.YangName = "local-igp-information"
    localIgpInformation.EntityData.BundleName = "cisco_ios_xr"
    localIgpInformation.EntityData.ParentYangName = "ipv4-link"
    localIgpInformation.EntityData.SegmentPath = "local-igp-information"
    localIgpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIgpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIgpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIgpInformation.EntityData.Children = types.NewOrderedMap()
    localIgpInformation.EntityData.Children.Append("igp", types.YChild{"Igp", &localIgpInformation.Igp})
    localIgpInformation.EntityData.Leafs = types.NewOrderedMap()
    localIgpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier})

    localIgpInformation.EntityData.YListKeys = []string {}

    return &(localIgpInformation.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
// IGP-specific information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is XtcIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis

    // OSPF information.
    Ospf Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf

    // BGP information.
    Bgp Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
}

func (igp *Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "local-igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("isis", types.YChild{"Isis", &igp.Isis})
    igp.EntityData.Children.Append("ospf", types.YChild{"Ospf", &igp.Ospf})
    igp.EntityData.Children.Append("bgp", types.YChild{"Bgp", &igp.Bgp})
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("igp-id", types.YLeaf{"IgpId", igp.IgpId})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis
// ISIS information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", isis.SystemId})
    isis.EntityData.Leafs.Append("level", types.YLeaf{"Level", isis.Level})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospf.RouterId})
    ospf.EntityData.Leafs.Append("area", types.YLeaf{"Area", ospf.Area})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
// BGP information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (bgp *Xtc_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", bgp.RouterId})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
}

func (remoteNodeProtocolIdentifier *Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    remoteNodeProtocolIdentifier.EntityData.YFilter = remoteNodeProtocolIdentifier.YFilter
    remoteNodeProtocolIdentifier.EntityData.YangName = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteNodeProtocolIdentifier.EntityData.ParentYangName = "ipv4-link"
    remoteNodeProtocolIdentifier.EntityData.SegmentPath = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", remoteNodeProtocolIdentifier.IgpInformation[i]})
    }
    remoteNodeProtocolIdentifier.EntityData.Leafs = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", remoteNodeProtocolIdentifier.NodeName})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id-set", types.YLeaf{"Ipv4BgpRouterIdSet", remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id", types.YLeaf{"Ipv4BgpRouterId", remoteNodeProtocolIdentifier.Ipv4BgpRouterId})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id-set", types.YLeaf{"Ipv4teRouterIdSet", remoteNodeProtocolIdentifier.Ipv4teRouterIdSet})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id", types.YLeaf{"Ipv4teRouterId", remoteNodeProtocolIdentifier.Ipv4teRouterId})

    remoteNodeProtocolIdentifier.EntityData.YListKeys = []string {}

    return &(remoteNodeProtocolIdentifier.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // IGP-specific information.
    Igp Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("igp", types.YChild{"Igp", &igpInformation.Igp})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is XtcIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("isis", types.YChild{"Isis", &igp.Isis})
    igp.EntityData.Children.Append("ospf", types.YChild{"Ospf", &igp.Ospf})
    igp.EntityData.Children.Append("bgp", types.YChild{"Bgp", &igp.Bgp})
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("igp-id", types.YLeaf{"IgpId", igp.IgpId})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", isis.SystemId})
    isis.EntityData.Leafs.Append("level", types.YLeaf{"Level", isis.Level})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospf.RouterId})
    ospf.EntityData.Leafs.Append("area", types.YLeaf{"Area", ospf.Area})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (bgp *Xtc_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", bgp.RouterId})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
// Adjacency SIDs
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Type. The type is XtcSid1.
    SidType interface{}

    // Prefix-SID algorithm number. The type is interface{} with range:
    // 0..4294967295.
    Algorithm interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Prefix.
    SidPrefix Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
}

func (adjacencySid *Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "ipv4-link"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &adjacencySid.SidPrefix})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()
    adjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", adjacencySid.SidType})
    adjacencySid.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", adjacencySid.Algorithm})
    adjacencySid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", adjacencySid.MplsLabel})

    adjacencySid.EntityData.YListKeys = []string {}

    return &(adjacencySid.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
// Prefix
type Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *Xtc_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "adjacency-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LocalIpv6Address interface{}

    // Remote IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteIpv6Address interface{}

    // IGP Metric. The type is interface{} with range: 0..4294967295.
    IgpMetric interface{}

    // TE Metric. The type is interface{} with range: 0..4294967295.
    TeMetric interface{}

    // Max link bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    MaximumLinkBandwidth interface{}

    // Max Reservable bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    MaxReservableBandwidth interface{}

    // Local node IGP information.
    LocalIgpInformation Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid.
    AdjacencySid []*Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
}

func (ipv6Link *Xtc_TopologyNodes_TopologyNode_Ipv6Link) GetEntityData() *types.CommonEntityData {
    ipv6Link.EntityData.YFilter = ipv6Link.YFilter
    ipv6Link.EntityData.YangName = "ipv6-link"
    ipv6Link.EntityData.BundleName = "cisco_ios_xr"
    ipv6Link.EntityData.ParentYangName = "topology-node"
    ipv6Link.EntityData.SegmentPath = "ipv6-link"
    ipv6Link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Link.EntityData.Children = types.NewOrderedMap()
    ipv6Link.EntityData.Children.Append("local-igp-information", types.YChild{"LocalIgpInformation", &ipv6Link.LocalIgpInformation})
    ipv6Link.EntityData.Children.Append("remote-node-protocol-identifier", types.YChild{"RemoteNodeProtocolIdentifier", &ipv6Link.RemoteNodeProtocolIdentifier})
    ipv6Link.EntityData.Children.Append("adjacency-sid", types.YChild{"AdjacencySid", nil})
    for i := range ipv6Link.AdjacencySid {
        ipv6Link.EntityData.Children.Append(types.GetSegmentPath(ipv6Link.AdjacencySid[i]), types.YChild{"AdjacencySid", ipv6Link.AdjacencySid[i]})
    }
    ipv6Link.EntityData.Leafs = types.NewOrderedMap()
    ipv6Link.EntityData.Leafs.Append("local-ipv6-address", types.YLeaf{"LocalIpv6Address", ipv6Link.LocalIpv6Address})
    ipv6Link.EntityData.Leafs.Append("remote-ipv6-address", types.YLeaf{"RemoteIpv6Address", ipv6Link.RemoteIpv6Address})
    ipv6Link.EntityData.Leafs.Append("igp-metric", types.YLeaf{"IgpMetric", ipv6Link.IgpMetric})
    ipv6Link.EntityData.Leafs.Append("te-metric", types.YLeaf{"TeMetric", ipv6Link.TeMetric})
    ipv6Link.EntityData.Leafs.Append("maximum-link-bandwidth", types.YLeaf{"MaximumLinkBandwidth", ipv6Link.MaximumLinkBandwidth})
    ipv6Link.EntityData.Leafs.Append("max-reservable-bandwidth", types.YLeaf{"MaxReservableBandwidth", ipv6Link.MaxReservableBandwidth})

    ipv6Link.EntityData.YListKeys = []string {}

    return &(ipv6Link.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // IGP-specific information.
    Igp Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
}

func (localIgpInformation *Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetEntityData() *types.CommonEntityData {
    localIgpInformation.EntityData.YFilter = localIgpInformation.YFilter
    localIgpInformation.EntityData.YangName = "local-igp-information"
    localIgpInformation.EntityData.BundleName = "cisco_ios_xr"
    localIgpInformation.EntityData.ParentYangName = "ipv6-link"
    localIgpInformation.EntityData.SegmentPath = "local-igp-information"
    localIgpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIgpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIgpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIgpInformation.EntityData.Children = types.NewOrderedMap()
    localIgpInformation.EntityData.Children.Append("igp", types.YChild{"Igp", &localIgpInformation.Igp})
    localIgpInformation.EntityData.Leafs = types.NewOrderedMap()
    localIgpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier})

    localIgpInformation.EntityData.YListKeys = []string {}

    return &(localIgpInformation.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
// IGP-specific information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is XtcIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis

    // OSPF information.
    Ospf Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf

    // BGP information.
    Bgp Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
}

func (igp *Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "local-igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("isis", types.YChild{"Isis", &igp.Isis})
    igp.EntityData.Children.Append("ospf", types.YChild{"Ospf", &igp.Ospf})
    igp.EntityData.Children.Append("bgp", types.YChild{"Bgp", &igp.Bgp})
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("igp-id", types.YLeaf{"IgpId", igp.IgpId})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis
// ISIS information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", isis.SystemId})
    isis.EntityData.Leafs.Append("level", types.YLeaf{"Level", isis.Level})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospf.RouterId})
    ospf.EntityData.Leafs.Append("area", types.YLeaf{"Area", ospf.Area})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
// BGP information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (bgp *Xtc_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", bgp.RouterId})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
}

func (remoteNodeProtocolIdentifier *Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    remoteNodeProtocolIdentifier.EntityData.YFilter = remoteNodeProtocolIdentifier.YFilter
    remoteNodeProtocolIdentifier.EntityData.YangName = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteNodeProtocolIdentifier.EntityData.ParentYangName = "ipv6-link"
    remoteNodeProtocolIdentifier.EntityData.SegmentPath = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", remoteNodeProtocolIdentifier.IgpInformation[i]})
    }
    remoteNodeProtocolIdentifier.EntityData.Leafs = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", remoteNodeProtocolIdentifier.NodeName})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id-set", types.YLeaf{"Ipv4BgpRouterIdSet", remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id", types.YLeaf{"Ipv4BgpRouterId", remoteNodeProtocolIdentifier.Ipv4BgpRouterId})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id-set", types.YLeaf{"Ipv4teRouterIdSet", remoteNodeProtocolIdentifier.Ipv4teRouterIdSet})
    remoteNodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id", types.YLeaf{"Ipv4teRouterId", remoteNodeProtocolIdentifier.Ipv4teRouterId})

    remoteNodeProtocolIdentifier.EntityData.YListKeys = []string {}

    return &(remoteNodeProtocolIdentifier.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // IGP-specific information.
    Igp Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("igp", types.YChild{"Igp", &igpInformation.Igp})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is XtcIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("isis", types.YChild{"Isis", &igp.Isis})
    igp.EntityData.Children.Append("ospf", types.YChild{"Ospf", &igp.Ospf})
    igp.EntityData.Children.Append("bgp", types.YChild{"Bgp", &igp.Bgp})
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("igp-id", types.YLeaf{"IgpId", igp.IgpId})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", isis.SystemId})
    isis.EntityData.Leafs.Append("level", types.YLeaf{"Level", isis.Level})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospf.RouterId})
    ospf.EntityData.Leafs.Append("area", types.YLeaf{"Area", ospf.Area})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (bgp *Xtc_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", bgp.RouterId})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
// Adjacency SIDs
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Type. The type is XtcSid1.
    SidType interface{}

    // Prefix-SID algorithm number. The type is interface{} with range:
    // 0..4294967295.
    Algorithm interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Prefix.
    SidPrefix Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
}

func (adjacencySid *Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "ipv6-link"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &adjacencySid.SidPrefix})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()
    adjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", adjacencySid.SidType})
    adjacencySid.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", adjacencySid.Algorithm})
    adjacencySid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", adjacencySid.MplsLabel})

    adjacencySid.EntityData.YListKeys = []string {}

    return &(adjacencySid.EntityData)
}

// Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
// Prefix
type Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *Xtc_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "adjacency-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// Xtc_PrefixInfos
// Prefixes database in XTC Agent
type Xtc_PrefixInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix information. The type is slice of Xtc_PrefixInfos_PrefixInfo.
    PrefixInfo []*Xtc_PrefixInfos_PrefixInfo
}

func (prefixInfos *Xtc_PrefixInfos) GetEntityData() *types.CommonEntityData {
    prefixInfos.EntityData.YFilter = prefixInfos.YFilter
    prefixInfos.EntityData.YangName = "prefix-infos"
    prefixInfos.EntityData.BundleName = "cisco_ios_xr"
    prefixInfos.EntityData.ParentYangName = "xtc"
    prefixInfos.EntityData.SegmentPath = "prefix-infos"
    prefixInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixInfos.EntityData.Children = types.NewOrderedMap()
    prefixInfos.EntityData.Children.Append("prefix-info", types.YChild{"PrefixInfo", nil})
    for i := range prefixInfos.PrefixInfo {
        prefixInfos.EntityData.Children.Append(types.GetSegmentPath(prefixInfos.PrefixInfo[i]), types.YChild{"PrefixInfo", prefixInfos.PrefixInfo[i]})
    }
    prefixInfos.EntityData.Leafs = types.NewOrderedMap()

    prefixInfos.EntityData.YListKeys = []string {}

    return &(prefixInfos.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo
// Prefix information
type Xtc_PrefixInfos_PrefixInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is interface{} with range:
    // 0..4294967295.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier

    // Prefix address. The type is slice of Xtc_PrefixInfos_PrefixInfo_Address.
    Address []*Xtc_PrefixInfos_PrefixInfo_Address
}

func (prefixInfo *Xtc_PrefixInfos_PrefixInfo) GetEntityData() *types.CommonEntityData {
    prefixInfo.EntityData.YFilter = prefixInfo.YFilter
    prefixInfo.EntityData.YangName = "prefix-info"
    prefixInfo.EntityData.BundleName = "cisco_ios_xr"
    prefixInfo.EntityData.ParentYangName = "prefix-infos"
    prefixInfo.EntityData.SegmentPath = "prefix-info" + types.AddKeyToken(prefixInfo.NodeIdentifier, "node-identifier")
    prefixInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixInfo.EntityData.Children = types.NewOrderedMap()
    prefixInfo.EntityData.Children.Append("node-protocol-identifier", types.YChild{"NodeProtocolIdentifier", &prefixInfo.NodeProtocolIdentifier})
    prefixInfo.EntityData.Children.Append("address", types.YChild{"Address", nil})
    for i := range prefixInfo.Address {
        prefixInfo.EntityData.Children.Append(types.GetSegmentPath(prefixInfo.Address[i]), types.YChild{"Address", prefixInfo.Address[i]})
    }
    prefixInfo.EntityData.Leafs = types.NewOrderedMap()
    prefixInfo.EntityData.Leafs.Append("node-identifier", types.YLeaf{"NodeIdentifier", prefixInfo.NodeIdentifier})
    prefixInfo.EntityData.Leafs.Append("node-identifier-xr", types.YLeaf{"NodeIdentifierXr", prefixInfo.NodeIdentifierXr})

    prefixInfo.EntityData.YListKeys = []string {"NodeIdentifier"}

    return &(prefixInfo.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier
// Node protocol identifier
type Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
}

func (nodeProtocolIdentifier *Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    nodeProtocolIdentifier.EntityData.YFilter = nodeProtocolIdentifier.YFilter
    nodeProtocolIdentifier.EntityData.YangName = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    nodeProtocolIdentifier.EntityData.ParentYangName = "prefix-info"
    nodeProtocolIdentifier.EntityData.SegmentPath = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", nodeProtocolIdentifier.IgpInformation[i]})
    }
    nodeProtocolIdentifier.EntityData.Leafs = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", nodeProtocolIdentifier.NodeName})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id-set", types.YLeaf{"Ipv4BgpRouterIdSet", nodeProtocolIdentifier.Ipv4BgpRouterIdSet})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4-bgp-router-id", types.YLeaf{"Ipv4BgpRouterId", nodeProtocolIdentifier.Ipv4BgpRouterId})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id-set", types.YLeaf{"Ipv4teRouterIdSet", nodeProtocolIdentifier.Ipv4teRouterIdSet})
    nodeProtocolIdentifier.EntityData.Leafs.Append("ipv4te-router-id", types.YLeaf{"Ipv4teRouterId", nodeProtocolIdentifier.Ipv4teRouterId})

    nodeProtocolIdentifier.EntityData.YListKeys = []string {}

    return &(nodeProtocolIdentifier.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // IGP-specific information.
    Igp Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("igp", types.YChild{"Igp", &igpInformation.Igp})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is XtcIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = types.NewOrderedMap()
    igp.EntityData.Children.Append("isis", types.YChild{"Isis", &igp.Isis})
    igp.EntityData.Children.Append("ospf", types.YChild{"Ospf", &igp.Ospf})
    igp.EntityData.Children.Append("bgp", types.YChild{"Bgp", &igp.Bgp})
    igp.EntityData.Leafs = types.NewOrderedMap()
    igp.EntityData.Leafs.Append("igp-id", types.YLeaf{"IgpId", igp.IgpId})

    igp.EntityData.YListKeys = []string {}

    return &(igp.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("system-id", types.YLeaf{"SystemId", isis.SystemId})
    isis.EntityData.Leafs.Append("level", types.YLeaf{"Level", isis.Level})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Leafs = types.NewOrderedMap()
    ospf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", ospf.RouterId})
    ospf.EntityData.Leafs.Append("area", types.YLeaf{"Area", ospf.Area})

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}
}

func (bgp *Xtc_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", bgp.RouterId})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_Address
// Prefix address
type Xtc_PrefixInfos_PrefixInfo_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix IP address.
    IpAddress Xtc_PrefixInfos_PrefixInfo_Address_IpAddress
}

func (address *Xtc_PrefixInfos_PrefixInfo_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "prefix-info"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("ip-address", types.YChild{"IpAddress", &address.IpAddress})
    address.EntityData.Leafs = types.NewOrderedMap()

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Xtc_PrefixInfos_PrefixInfo_Address_IpAddress
// Prefix IP address
type Xtc_PrefixInfos_PrefixInfo_Address_IpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is XtcAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipAddress *Xtc_PrefixInfos_PrefixInfo_Address_IpAddress) GetEntityData() *types.CommonEntityData {
    ipAddress.EntityData.YFilter = ipAddress.YFilter
    ipAddress.EntityData.YangName = "ip-address"
    ipAddress.EntityData.BundleName = "cisco_ios_xr"
    ipAddress.EntityData.ParentYangName = "address"
    ipAddress.EntityData.SegmentPath = "ip-address"
    ipAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipAddress.EntityData.Children = types.NewOrderedMap()
    ipAddress.EntityData.Leafs = types.NewOrderedMap()
    ipAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", ipAddress.AfName})
    ipAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipAddress.Ipv4})
    ipAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipAddress.Ipv6})

    ipAddress.EntityData.YListKeys = []string {}

    return &(ipAddress.EntityData)
}

