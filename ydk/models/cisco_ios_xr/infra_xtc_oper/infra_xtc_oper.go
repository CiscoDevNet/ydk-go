// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-xtc package operational data.
// 
// This module contains definitions
// for the following management objects:
//   pce-lsp-data: PCE LSP's data
//   pce-peer: pce peer
//   pce-topology: pce topology
//   pce: pce
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package infra_xtc_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_xtc_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-oper pce-lsp-data}", reflect.TypeOf(PceLspData{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-oper:pce-lsp-data", reflect.TypeOf(PceLspData{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-oper pce-peer}", reflect.TypeOf(PcePeer{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-oper:pce-peer", reflect.TypeOf(PcePeer{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-oper pce-topology}", reflect.TypeOf(PceTopology{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-oper:pce-topology", reflect.TypeOf(PceTopology{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-xtc-oper pce}", reflect.TypeOf(Pce{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-xtc-oper:pce", reflect.TypeOf(Pce{}))
}

// LspState represents LSP setup type
type LspState string

const (
    // LSP is down
    LspState_lsp_down LspState = "lsp-down"

    // LSP is up
    LspState_lsp_up LspState = "lsp-up"
)

// PceRro represents PCE RRO type
type PceRro string

const (
    // IPv4 Address
    PceRro_rro_type_ipv4_address PceRro = "rro-type-ipv4-address"

    // MPLS Label
    PceRro_rro_type_mpls_label PceRro = "rro-type-mpls-label"

    // Segment Routing IPv4 Node SID
    PceRro_rro_type_sripv4_node_sid PceRro = "rro-type-sripv4-node-sid"

    // Segment Routing IPv4 Adjacency SID
    PceRro_rro_type_sripv4_adjacency_sid PceRro = "rro-type-sripv4-adjacency-sid"

    // Segment Routing with NAI null
    PceRro_rro_type_sr_nai_null PceRro = "rro-type-sr-nai-null"
)

// PceCspfRc represents PCE CSPF Result Code
type PceCspfRc string

const (
    // Not set
    PceCspfRc_pce_cspf_not_set PceCspfRc = "pce-cspf-not-set"

    // Source not found
    PceCspfRc_pce_cspf_src_not_found PceCspfRc = "pce-cspf-src-not-found"

    // Destination not found
    PceCspfRc_pce_cspf_dst_not_found PceCspfRc = "pce-cspf-dst-not-found"

    // Second source not found
    PceCspfRc_pce_cspf_second_src_not_found PceCspfRc = "pce-cspf-second-src-not-found"

    // Second destination not found
    PceCspfRc_pce_cspf_second_dst_not_found PceCspfRc = "pce-cspf-second-dst-not-found"

    // No memory
    PceCspfRc_pce_cspf_no_mem PceCspfRc = "pce-cspf-no-mem"

    // Second path not resolved
    PceCspfRc_pce_cspf_ex_path_not_resolved PceCspfRc = "pce-cspf-ex-path-not-resolved"

    // No path
    PceCspfRc_pce_cspf_no_path PceCspfRc = "pce-cspf-no-path"

    // Shortest path success
    PceCspfRc_pce_cspf_sp_success PceCspfRc = "pce-cspf-sp-success"

    // Error
    PceCspfRc_pce_cspf_error PceCspfRc = "pce-cspf-error"

    // Fallback from SRLG-NODE to NODE
    PceCspfRc_pce_cspf_fallback_srlg_node_node PceCspfRc = "pce-cspf-fallback-srlg-node-node"

    // Fallback from SRLG-NODE to LINK
    PceCspfRc_pce_cspf_fallback_srlg_node_link PceCspfRc = "pce-cspf-fallback-srlg-node-link"

    // Fallback from SRLG-NODE to SP
    PceCspfRc_pce_cspf_fallback_srlg_node_sp PceCspfRc = "pce-cspf-fallback-srlg-node-sp"

    // Fallback from NODE to LINK
    PceCspfRc_pce_cspf_fallback_node_link PceCspfRc = "pce-cspf-fallback-node-link"

    // Fallback from LINK to SP
    PceCspfRc_pce_cspf_fallback_link_sp PceCspfRc = "pce-cspf-fallback-link-sp"

    // Fallback from NODE to SP
    PceCspfRc_pce_cspf_fallback_node_sp PceCspfRc = "pce-cspf-fallback-node-sp"

    // Fallback from SRLG to LINK
    PceCspfRc_pce_cspf_fallback_srlg_link PceCspfRc = "pce-cspf-fallback-srlg-link"

    // Fallback from SRLG to SP
    PceCspfRc_pce_cspf_fallback_srlg_sp PceCspfRc = "pce-cspf-fallback-srlg-sp"

    // Disjoint path success
    PceCspfRc_pce_cspf_dp_success PceCspfRc = "pce-cspf-dp-success"
)

// PceSrSid represents PCE SR SID type
type PceSrSid string

const (
    // IPv4 Node SID
    PceSrSid_ipv4_node_sid PceSrSid = "ipv4-node-sid"

    // IPv4 Adjacency SID
    PceSrSid_ipv4_adjacency_sid PceSrSid = "ipv4-adjacency-sid"

    // Unknown SID
    PceSrSid_unknown_sid PceSrSid = "unknown-sid"
)

// PceHeadendSwap represents PCE Headends Swap Code
type PceHeadendSwap string

const (
    // Headends not swapped
    PceHeadendSwap_pcehs_none PceHeadendSwap = "pcehs-none"

    // Headends swapped
    PceHeadendSwap_pcehs_plain PceHeadendSwap = "pcehs-plain"

    // Headends swapped with increment
    PceHeadendSwap_pcehs_rwi PceHeadendSwap = "pcehs-rwi"
)

// PcepLspState represents PCEP operation protocol
type PcepLspState string

const (
    // LSP is down
    PcepLspState_lsp_down PcepLspState = "lsp-down"

    // LSP is up
    PcepLspState_lsp_up PcepLspState = "lsp-up"

    // LSP is active (carrying traffic)
    PcepLspState_lsp_active PcepLspState = "lsp-active"

    // LSP is going down
    PcepLspState_lsp_going_down PcepLspState = "lsp-going-down"

    // LSP is being signaled
    PcepLspState_lsp_being_signaled PcepLspState = "lsp-being-signaled"
)

// PceIgpInfoId represents IGP IDs
type PceIgpInfoId string

const (
    // ISIS
    PceIgpInfoId_isis PceIgpInfoId = "isis"

    // OSPF
    PceIgpInfoId_ospf PceIgpInfoId = "ospf"

    // BGP
    PceIgpInfoId_bgp PceIgpInfoId = "bgp"
)

// PceProto represents PCE peer protocol
type PceProto string

const (
    // PCE protocol
    PceProto_pcep PceProto = "pcep"

    // Netconf protocol
    PceProto_netconf PceProto = "netconf"
)

// PcepState represents PCEP State
type PcepState string

const (
    // TCP close
    PcepState_tcp_close PcepState = "tcp-close"

    // TCP listen
    PcepState_tcp_listen PcepState = "tcp-listen"

    // TCP connect
    PcepState_tcp_connect PcepState = "tcp-connect"

    // PCEP closed
    PcepState_pcep_closed PcepState = "pcep-closed"

    // PCEP opening
    PcepState_pcep_opening PcepState = "pcep-opening"

    // PCEP open
    PcepState_pcep_open PcepState = "pcep-open"
)

// Sid represents SID Types
type Sid string

const (
    // Protected Adjacency SID
    Sid_sr_protected_adj_sid Sid = "sr-protected-adj-sid"

    // Unprotected Adjacency SID
    Sid_sr_unprotected_adj_sid Sid = "sr-unprotected-adj-sid"

    // BGP egress peer engineering SID
    Sid_sr_bgp_egress_peer_engineering_sid Sid = "sr-bgp-egress-peer-engineering-sid"

    // Regular prefix SID
    Sid_sr_reqular_prefix_sid Sid = "sr-reqular-prefix-sid"

    // Strict prefix SID
    Sid_sr_strict_prefix_sid Sid = "sr-strict-prefix-sid"
)

// LspSetup represents LSP setup type
type LspSetup string

const (
    // LSP is established using RSVP-TE
    LspSetup_setup_rsvp LspSetup = "setup-rsvp"

    // LSP is established using SR-TE
    LspSetup_setup_sr LspSetup = "setup-sr"

    // Unknown LSP establishment method
    LspSetup_setup_unknown LspSetup = "setup-unknown"
)

// PceAsso represents Pce asso
type PceAsso string

const (
    // Unknown type
    PceAsso_unknown PceAsso = "unknown"

    // LINK
    PceAsso_link PceAsso = "link"

    // NODE
    PceAsso_node PceAsso = "node"

    // SRLG
    PceAsso_srlg PceAsso = "srlg"
)

// PceAfId represents Pce af id
type PceAfId string

const (
    // None
    PceAfId_none PceAfId = "none"

    // IPv4
    PceAfId_ipv4 PceAfId = "ipv4"

    // IPv6
    PceAfId_ipv6 PceAfId = "ipv6"
)

// PceLspData
// PCE LSP's data
type PceLspData struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel database in XTC.
    TunnelInfos PceLspData_TunnelInfos

    // LSP summary database in XTC.
    LspSummary PceLspData_LspSummary

    // Detailed tunnel database in XTC.
    TunnelDetailInfos PceLspData_TunnelDetailInfos
}

func (pceLspData *PceLspData) GetFilter() yfilter.YFilter { return pceLspData.YFilter }

func (pceLspData *PceLspData) SetFilter(yf yfilter.YFilter) { pceLspData.YFilter = yf }

func (pceLspData *PceLspData) GetGoName(yname string) string {
    if yname == "tunnel-infos" { return "TunnelInfos" }
    if yname == "lsp-summary" { return "LspSummary" }
    if yname == "tunnel-detail-infos" { return "TunnelDetailInfos" }
    return ""
}

func (pceLspData *PceLspData) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-xtc-oper:pce-lsp-data"
}

func (pceLspData *PceLspData) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-infos" {
        return &pceLspData.TunnelInfos
    }
    if childYangName == "lsp-summary" {
        return &pceLspData.LspSummary
    }
    if childYangName == "tunnel-detail-infos" {
        return &pceLspData.TunnelDetailInfos
    }
    return nil
}

func (pceLspData *PceLspData) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tunnel-infos"] = &pceLspData.TunnelInfos
    children["lsp-summary"] = &pceLspData.LspSummary
    children["tunnel-detail-infos"] = &pceLspData.TunnelDetailInfos
    return children
}

func (pceLspData *PceLspData) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pceLspData *PceLspData) GetBundleName() string { return "cisco_ios_xr" }

func (pceLspData *PceLspData) GetYangName() string { return "pce-lsp-data" }

func (pceLspData *PceLspData) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pceLspData *PceLspData) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pceLspData *PceLspData) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pceLspData *PceLspData) SetParent(parent types.Entity) { pceLspData.parent = parent }

func (pceLspData *PceLspData) GetParent() types.Entity { return pceLspData.parent }

func (pceLspData *PceLspData) GetParentYangName() string { return "Cisco-IOS-XR-infra-xtc-oper" }

// PceLspData_TunnelInfos
// Tunnel database in XTC
type PceLspData_TunnelInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of PceLspData_TunnelInfos_TunnelInfo.
    TunnelInfo []PceLspData_TunnelInfos_TunnelInfo
}

func (tunnelInfos *PceLspData_TunnelInfos) GetFilter() yfilter.YFilter { return tunnelInfos.YFilter }

func (tunnelInfos *PceLspData_TunnelInfos) SetFilter(yf yfilter.YFilter) { tunnelInfos.YFilter = yf }

func (tunnelInfos *PceLspData_TunnelInfos) GetGoName(yname string) string {
    if yname == "tunnel-info" { return "TunnelInfo" }
    return ""
}

func (tunnelInfos *PceLspData_TunnelInfos) GetSegmentPath() string {
    return "tunnel-infos"
}

func (tunnelInfos *PceLspData_TunnelInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-info" {
        for _, c := range tunnelInfos.TunnelInfo {
            if tunnelInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelInfos_TunnelInfo{}
        tunnelInfos.TunnelInfo = append(tunnelInfos.TunnelInfo, child)
        return &tunnelInfos.TunnelInfo[len(tunnelInfos.TunnelInfo)-1]
    }
    return nil
}

func (tunnelInfos *PceLspData_TunnelInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelInfos.TunnelInfo {
        children[tunnelInfos.TunnelInfo[i].GetSegmentPath()] = &tunnelInfos.TunnelInfo[i]
    }
    return children
}

func (tunnelInfos *PceLspData_TunnelInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelInfos *PceLspData_TunnelInfos) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelInfos *PceLspData_TunnelInfos) GetYangName() string { return "tunnel-infos" }

func (tunnelInfos *PceLspData_TunnelInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelInfos *PceLspData_TunnelInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelInfos *PceLspData_TunnelInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelInfos *PceLspData_TunnelInfos) SetParent(parent types.Entity) { tunnelInfos.parent = parent }

func (tunnelInfos *PceLspData_TunnelInfos) GetParent() types.Entity { return tunnelInfos.parent }

func (tunnelInfos *PceLspData_TunnelInfos) GetParentYangName() string { return "pce-lsp-data" }

// PceLspData_TunnelInfos_TunnelInfo
// Tunnel information
type PceLspData_TunnelInfos_TunnelInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Brief LSP information. The type is slice of
    // PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation.
    BriefLspInformation []PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetFilter() yfilter.YFilter { return tunnelInfo.YFilter }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) SetFilter(yf yfilter.YFilter) { tunnelInfo.YFilter = yf }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "plsp-id" { return "PlspId" }
    if yname == "tunnel-name" { return "TunnelName" }
    if yname == "pcc-address" { return "PccAddress" }
    if yname == "tunnel-name-xr" { return "TunnelNameXr" }
    if yname == "brief-lsp-information" { return "BriefLspInformation" }
    return ""
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetSegmentPath() string {
    return "tunnel-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelInfo.TunnelName) + "']"
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-lsp-information" {
        for _, c := range tunnelInfo.BriefLspInformation {
            if tunnelInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation{}
        tunnelInfo.BriefLspInformation = append(tunnelInfo.BriefLspInformation, child)
        return &tunnelInfo.BriefLspInformation[len(tunnelInfo.BriefLspInformation)-1]
    }
    return nil
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelInfo.BriefLspInformation {
        children[tunnelInfo.BriefLspInformation[i].GetSegmentPath()] = &tunnelInfo.BriefLspInformation[i]
    }
    return children
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = tunnelInfo.PeerAddress
    leafs["plsp-id"] = tunnelInfo.PlspId
    leafs["tunnel-name"] = tunnelInfo.TunnelName
    leafs["pcc-address"] = tunnelInfo.PccAddress
    leafs["tunnel-name-xr"] = tunnelInfo.TunnelNameXr
    return leafs
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetYangName() string { return "tunnel-info" }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) SetParent(parent types.Entity) { tunnelInfo.parent = parent }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetParent() types.Entity { return tunnelInfo.parent }

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetParentYangName() string { return "tunnel-infos" }

// PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation
// Brief LSP information
type PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // LSP ID. The type is interface{} with range: 0..4294967295.
    Lspid interface{}

    // Binding SID. The type is interface{} with range: 0..4294967295.
    BindingSid interface{}

    // LSP Setup Type. The type is LspSetup.
    LspSetupType interface{}

    // Operational state. The type is PcepLspState.
    OperationalState interface{}

    // Admin state. The type is LspState.
    AdministrativeState interface{}
}

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetFilter() yfilter.YFilter { return briefLspInformation.YFilter }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) SetFilter(yf yfilter.YFilter) { briefLspInformation.YFilter = yf }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "lspid" { return "Lspid" }
    if yname == "binding-sid" { return "BindingSid" }
    if yname == "lsp-setup-type" { return "LspSetupType" }
    if yname == "operational-state" { return "OperationalState" }
    if yname == "administrative-state" { return "AdministrativeState" }
    return ""
}

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetSegmentPath() string {
    return "brief-lsp-information"
}

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = briefLspInformation.SourceAddress
    leafs["destination-address"] = briefLspInformation.DestinationAddress
    leafs["tunnel-id"] = briefLspInformation.TunnelId
    leafs["lspid"] = briefLspInformation.Lspid
    leafs["binding-sid"] = briefLspInformation.BindingSid
    leafs["lsp-setup-type"] = briefLspInformation.LspSetupType
    leafs["operational-state"] = briefLspInformation.OperationalState
    leafs["administrative-state"] = briefLspInformation.AdministrativeState
    return leafs
}

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetYangName() string { return "brief-lsp-information" }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) SetParent(parent types.Entity) { briefLspInformation.parent = parent }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetParent() types.Entity { return briefLspInformation.parent }

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetParentYangName() string { return "tunnel-info" }

// PceLspData_LspSummary
// LSP summary database in XTC
type PceLspData_LspSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary for all peers.
    AllLsPs PceLspData_LspSummary_AllLsPs

    // Number of LSPs for specific peer. The type is slice of
    // PceLspData_LspSummary_PeerLsPsInfo.
    PeerLsPsInfo []PceLspData_LspSummary_PeerLsPsInfo
}

func (lspSummary *PceLspData_LspSummary) GetFilter() yfilter.YFilter { return lspSummary.YFilter }

func (lspSummary *PceLspData_LspSummary) SetFilter(yf yfilter.YFilter) { lspSummary.YFilter = yf }

func (lspSummary *PceLspData_LspSummary) GetGoName(yname string) string {
    if yname == "all-ls-ps" { return "AllLsPs" }
    if yname == "peer-ls-ps-info" { return "PeerLsPsInfo" }
    return ""
}

func (lspSummary *PceLspData_LspSummary) GetSegmentPath() string {
    return "lsp-summary"
}

func (lspSummary *PceLspData_LspSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-ls-ps" {
        return &lspSummary.AllLsPs
    }
    if childYangName == "peer-ls-ps-info" {
        for _, c := range lspSummary.PeerLsPsInfo {
            if lspSummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_LspSummary_PeerLsPsInfo{}
        lspSummary.PeerLsPsInfo = append(lspSummary.PeerLsPsInfo, child)
        return &lspSummary.PeerLsPsInfo[len(lspSummary.PeerLsPsInfo)-1]
    }
    return nil
}

func (lspSummary *PceLspData_LspSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["all-ls-ps"] = &lspSummary.AllLsPs
    for i := range lspSummary.PeerLsPsInfo {
        children[lspSummary.PeerLsPsInfo[i].GetSegmentPath()] = &lspSummary.PeerLsPsInfo[i]
    }
    return children
}

func (lspSummary *PceLspData_LspSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspSummary *PceLspData_LspSummary) GetBundleName() string { return "cisco_ios_xr" }

func (lspSummary *PceLspData_LspSummary) GetYangName() string { return "lsp-summary" }

func (lspSummary *PceLspData_LspSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspSummary *PceLspData_LspSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspSummary *PceLspData_LspSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspSummary *PceLspData_LspSummary) SetParent(parent types.Entity) { lspSummary.parent = parent }

func (lspSummary *PceLspData_LspSummary) GetParent() types.Entity { return lspSummary.parent }

func (lspSummary *PceLspData_LspSummary) GetParentYangName() string { return "pce-lsp-data" }

// PceLspData_LspSummary_AllLsPs
// Summary for all peers
type PceLspData_LspSummary_AllLsPs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of all LSPs. The type is interface{} with range: 0..4294967295.
    AllLsPs interface{}

    // Number of operational LSPs. The type is interface{} with range:
    // 0..4294967295.
    UpLsPs interface{}

    // Number of administratively up LSPs. The type is interface{} with range:
    // 0..4294967295.
    AdminUpLsPs interface{}

    // Number of LSPs with Segment routing setup type. The type is interface{}
    // with range: 0..4294967295.
    SrLsPs interface{}

    // Number of LSPs with RSVP setup type. The type is interface{} with range:
    // 0..4294967295.
    RsvpLsPs interface{}
}

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetFilter() yfilter.YFilter { return allLsPs.YFilter }

func (allLsPs *PceLspData_LspSummary_AllLsPs) SetFilter(yf yfilter.YFilter) { allLsPs.YFilter = yf }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetGoName(yname string) string {
    if yname == "all-ls-ps" { return "AllLsPs" }
    if yname == "up-ls-ps" { return "UpLsPs" }
    if yname == "admin-up-ls-ps" { return "AdminUpLsPs" }
    if yname == "sr-ls-ps" { return "SrLsPs" }
    if yname == "rsvp-ls-ps" { return "RsvpLsPs" }
    return ""
}

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetSegmentPath() string {
    return "all-ls-ps"
}

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all-ls-ps"] = allLsPs.AllLsPs
    leafs["up-ls-ps"] = allLsPs.UpLsPs
    leafs["admin-up-ls-ps"] = allLsPs.AdminUpLsPs
    leafs["sr-ls-ps"] = allLsPs.SrLsPs
    leafs["rsvp-ls-ps"] = allLsPs.RsvpLsPs
    return leafs
}

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetBundleName() string { return "cisco_ios_xr" }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetYangName() string { return "all-ls-ps" }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allLsPs *PceLspData_LspSummary_AllLsPs) SetParent(parent types.Entity) { allLsPs.parent = parent }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetParent() types.Entity { return allLsPs.parent }

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetParentYangName() string { return "lsp-summary" }

// PceLspData_LspSummary_PeerLsPsInfo
// Number of LSPs for specific peer
type PceLspData_LspSummary_PeerLsPsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Number of LSPs for specific peer.
    LspSummary PceLspData_LspSummary_PeerLsPsInfo_LspSummary
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetFilter() yfilter.YFilter { return peerLsPsInfo.YFilter }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) SetFilter(yf yfilter.YFilter) { peerLsPsInfo.YFilter = yf }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "lsp-summary" { return "LspSummary" }
    return ""
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetSegmentPath() string {
    return "peer-ls-ps-info"
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-summary" {
        return &peerLsPsInfo.LspSummary
    }
    return nil
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsp-summary"] = &peerLsPsInfo.LspSummary
    return children
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peerLsPsInfo.PeerAddress
    return leafs
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetYangName() string { return "peer-ls-ps-info" }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) SetParent(parent types.Entity) { peerLsPsInfo.parent = parent }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetParent() types.Entity { return peerLsPsInfo.parent }

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetParentYangName() string { return "lsp-summary" }

// PceLspData_LspSummary_PeerLsPsInfo_LspSummary
// Number of LSPs for specific peer
type PceLspData_LspSummary_PeerLsPsInfo_LspSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of all LSPs. The type is interface{} with range: 0..4294967295.
    AllLsPs interface{}

    // Number of operational LSPs. The type is interface{} with range:
    // 0..4294967295.
    UpLsPs interface{}

    // Number of administratively up LSPs. The type is interface{} with range:
    // 0..4294967295.
    AdminUpLsPs interface{}

    // Number of LSPs with Segment routing setup type. The type is interface{}
    // with range: 0..4294967295.
    SrLsPs interface{}

    // Number of LSPs with RSVP setup type. The type is interface{} with range:
    // 0..4294967295.
    RsvpLsPs interface{}
}

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetFilter() yfilter.YFilter { return lspSummary.YFilter }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) SetFilter(yf yfilter.YFilter) { lspSummary.YFilter = yf }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetGoName(yname string) string {
    if yname == "all-ls-ps" { return "AllLsPs" }
    if yname == "up-ls-ps" { return "UpLsPs" }
    if yname == "admin-up-ls-ps" { return "AdminUpLsPs" }
    if yname == "sr-ls-ps" { return "SrLsPs" }
    if yname == "rsvp-ls-ps" { return "RsvpLsPs" }
    return ""
}

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetSegmentPath() string {
    return "lsp-summary"
}

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all-ls-ps"] = lspSummary.AllLsPs
    leafs["up-ls-ps"] = lspSummary.UpLsPs
    leafs["admin-up-ls-ps"] = lspSummary.AdminUpLsPs
    leafs["sr-ls-ps"] = lspSummary.SrLsPs
    leafs["rsvp-ls-ps"] = lspSummary.RsvpLsPs
    return leafs
}

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetBundleName() string { return "cisco_ios_xr" }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetYangName() string { return "lsp-summary" }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) SetParent(parent types.Entity) { lspSummary.parent = parent }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetParent() types.Entity { return lspSummary.parent }

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetParentYangName() string { return "peer-ls-ps-info" }

// PceLspData_TunnelDetailInfos
// Detailed tunnel database in XTC
type PceLspData_TunnelDetailInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed tunnel information. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo.
    TunnelDetailInfo []PceLspData_TunnelDetailInfos_TunnelDetailInfo
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetFilter() yfilter.YFilter { return tunnelDetailInfos.YFilter }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) SetFilter(yf yfilter.YFilter) { tunnelDetailInfos.YFilter = yf }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetGoName(yname string) string {
    if yname == "tunnel-detail-info" { return "TunnelDetailInfo" }
    return ""
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetSegmentPath() string {
    return "tunnel-detail-infos"
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-detail-info" {
        for _, c := range tunnelDetailInfos.TunnelDetailInfo {
            if tunnelDetailInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo{}
        tunnelDetailInfos.TunnelDetailInfo = append(tunnelDetailInfos.TunnelDetailInfo, child)
        return &tunnelDetailInfos.TunnelDetailInfo[len(tunnelDetailInfos.TunnelDetailInfo)-1]
    }
    return nil
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelDetailInfos.TunnelDetailInfo {
        children[tunnelDetailInfos.TunnelDetailInfo[i].GetSegmentPath()] = &tunnelDetailInfos.TunnelDetailInfo[i]
    }
    return children
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetYangName() string { return "tunnel-detail-infos" }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) SetParent(parent types.Entity) { tunnelDetailInfos.parent = parent }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetParent() types.Entity { return tunnelDetailInfos.parent }

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetParentYangName() string { return "pce-lsp-data" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo
// Detailed tunnel information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Private LSP information.
    PrivateLspInformation PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation

    // Detail LSP information. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation.
    DetailLspInformation []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetFilter() yfilter.YFilter { return tunnelDetailInfo.YFilter }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) SetFilter(yf yfilter.YFilter) { tunnelDetailInfo.YFilter = yf }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "plsp-id" { return "PlspId" }
    if yname == "tunnel-name" { return "TunnelName" }
    if yname == "pcc-address" { return "PccAddress" }
    if yname == "tunnel-name-xr" { return "TunnelNameXr" }
    if yname == "private-lsp-information" { return "PrivateLspInformation" }
    if yname == "detail-lsp-information" { return "DetailLspInformation" }
    return ""
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetSegmentPath() string {
    return "tunnel-detail-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelDetailInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelDetailInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelDetailInfo.TunnelName) + "']"
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "private-lsp-information" {
        return &tunnelDetailInfo.PrivateLspInformation
    }
    if childYangName == "detail-lsp-information" {
        for _, c := range tunnelDetailInfo.DetailLspInformation {
            if tunnelDetailInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation{}
        tunnelDetailInfo.DetailLspInformation = append(tunnelDetailInfo.DetailLspInformation, child)
        return &tunnelDetailInfo.DetailLspInformation[len(tunnelDetailInfo.DetailLspInformation)-1]
    }
    return nil
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["private-lsp-information"] = &tunnelDetailInfo.PrivateLspInformation
    for i := range tunnelDetailInfo.DetailLspInformation {
        children[tunnelDetailInfo.DetailLspInformation[i].GetSegmentPath()] = &tunnelDetailInfo.DetailLspInformation[i]
    }
    return children
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = tunnelDetailInfo.PeerAddress
    leafs["plsp-id"] = tunnelDetailInfo.PlspId
    leafs["tunnel-name"] = tunnelDetailInfo.TunnelName
    leafs["pcc-address"] = tunnelDetailInfo.PccAddress
    leafs["tunnel-name-xr"] = tunnelDetailInfo.TunnelNameXr
    return leafs
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetYangName() string { return "tunnel-detail-info" }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) SetParent(parent types.Entity) { tunnelDetailInfo.parent = parent }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetParent() types.Entity { return tunnelDetailInfo.parent }

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetParentYangName() string { return "tunnel-detail-infos" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation
// Private LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP Event buffer. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer.
    EventBuffer []PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetFilter() yfilter.YFilter { return privateLspInformation.YFilter }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) SetFilter(yf yfilter.YFilter) { privateLspInformation.YFilter = yf }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetGoName(yname string) string {
    if yname == "event-buffer" { return "EventBuffer" }
    return ""
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetSegmentPath() string {
    return "private-lsp-information"
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "event-buffer" {
        for _, c := range privateLspInformation.EventBuffer {
            if privateLspInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer{}
        privateLspInformation.EventBuffer = append(privateLspInformation.EventBuffer, child)
        return &privateLspInformation.EventBuffer[len(privateLspInformation.EventBuffer)-1]
    }
    return nil
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range privateLspInformation.EventBuffer {
        children[privateLspInformation.EventBuffer[i].GetSegmentPath()] = &privateLspInformation.EventBuffer[i]
    }
    return children
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetYangName() string { return "private-lsp-information" }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) SetParent(parent types.Entity) { privateLspInformation.parent = parent }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetParent() types.Entity { return privateLspInformation.parent }

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetParentYangName() string { return "tunnel-detail-info" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
// LSP Event buffer
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Event message. The type is string.
    EventMessage interface{}

    // Event time, relative to Jan 1, 1970. The type is interface{} with range:
    // 0..4294967295.
    TimeStamp interface{}
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetFilter() yfilter.YFilter { return eventBuffer.YFilter }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) SetFilter(yf yfilter.YFilter) { eventBuffer.YFilter = yf }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetGoName(yname string) string {
    if yname == "event-message" { return "EventMessage" }
    if yname == "time-stamp" { return "TimeStamp" }
    return ""
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetSegmentPath() string {
    return "event-buffer"
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-message"] = eventBuffer.EventMessage
    leafs["time-stamp"] = eventBuffer.TimeStamp
    return leafs
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetYangName() string { return "event-buffer" }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) SetParent(parent types.Entity) { eventBuffer.parent = parent }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetParent() types.Entity { return eventBuffer.parent }

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetParentYangName() string { return "private-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
// Detail LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // True if router notifies signal bandwidth. The type is bool.
    SignaledBandwidthSpecified interface{}

    // Signaled Bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    SignaledBandwidth interface{}

    // True if router notifies actual bandwidth. The type is bool.
    ActualBandwidthSpecified interface{}

    // Actual bandwidth utilized in the data-plane. The type is interface{} with
    // range: 0..18446744073709551615.
    ActualBandwidth interface{}

    // LSP Role. The type is interface{} with range: 0..4294967295.
    LspRole interface{}

    // Computing PCE. The type is interface{} with range: 0..4294967295.
    ComputingPce interface{}

    // Sub delegated PCE. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubDelegatedPce interface{}

    // State-sync PCE. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StateSyncPce interface{}

    // Reporting PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReportingPccAddress interface{}

    // List of SLRGs used by LSP. The type is slice of interface{} with range:
    // 0..4294967295.
    SrlgInfo []interface{}

    // Brief LSP information.
    BriefLspInformation PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation

    // Paths.
    ErOs PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs

    // PCEP related LSP information.
    LsppcepInformation PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation

    // LSP association information.
    LspAssociationInfo PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo

    // LSP attributes.
    LspAttributes PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes

    // RRO. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro.
    Rro []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
}

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetFilter() yfilter.YFilter { return detailLspInformation.YFilter }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) SetFilter(yf yfilter.YFilter) { detailLspInformation.YFilter = yf }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetGoName(yname string) string {
    if yname == "signaled-bandwidth-specified" { return "SignaledBandwidthSpecified" }
    if yname == "signaled-bandwidth" { return "SignaledBandwidth" }
    if yname == "actual-bandwidth-specified" { return "ActualBandwidthSpecified" }
    if yname == "actual-bandwidth" { return "ActualBandwidth" }
    if yname == "lsp-role" { return "LspRole" }
    if yname == "computing-pce" { return "ComputingPce" }
    if yname == "sub-delegated-pce" { return "SubDelegatedPce" }
    if yname == "state-sync-pce" { return "StateSyncPce" }
    if yname == "reporting-pcc-address" { return "ReportingPccAddress" }
    if yname == "srlg-info" { return "SrlgInfo" }
    if yname == "brief-lsp-information" { return "BriefLspInformation" }
    if yname == "er-os" { return "ErOs" }
    if yname == "lsppcep-information" { return "LsppcepInformation" }
    if yname == "lsp-association-info" { return "LspAssociationInfo" }
    if yname == "lsp-attributes" { return "LspAttributes" }
    if yname == "rro" { return "Rro" }
    return ""
}

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetSegmentPath() string {
    return "detail-lsp-information"
}

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-lsp-information" {
        return &detailLspInformation.BriefLspInformation
    }
    if childYangName == "er-os" {
        return &detailLspInformation.ErOs
    }
    if childYangName == "lsppcep-information" {
        return &detailLspInformation.LsppcepInformation
    }
    if childYangName == "lsp-association-info" {
        return &detailLspInformation.LspAssociationInfo
    }
    if childYangName == "lsp-attributes" {
        return &detailLspInformation.LspAttributes
    }
    if childYangName == "rro" {
        for _, c := range detailLspInformation.Rro {
            if detailLspInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro{}
        detailLspInformation.Rro = append(detailLspInformation.Rro, child)
        return &detailLspInformation.Rro[len(detailLspInformation.Rro)-1]
    }
    return nil
}

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-lsp-information"] = &detailLspInformation.BriefLspInformation
    children["er-os"] = &detailLspInformation.ErOs
    children["lsppcep-information"] = &detailLspInformation.LsppcepInformation
    children["lsp-association-info"] = &detailLspInformation.LspAssociationInfo
    children["lsp-attributes"] = &detailLspInformation.LspAttributes
    for i := range detailLspInformation.Rro {
        children[detailLspInformation.Rro[i].GetSegmentPath()] = &detailLspInformation.Rro[i]
    }
    return children
}

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["signaled-bandwidth-specified"] = detailLspInformation.SignaledBandwidthSpecified
    leafs["signaled-bandwidth"] = detailLspInformation.SignaledBandwidth
    leafs["actual-bandwidth-specified"] = detailLspInformation.ActualBandwidthSpecified
    leafs["actual-bandwidth"] = detailLspInformation.ActualBandwidth
    leafs["lsp-role"] = detailLspInformation.LspRole
    leafs["computing-pce"] = detailLspInformation.ComputingPce
    leafs["sub-delegated-pce"] = detailLspInformation.SubDelegatedPce
    leafs["state-sync-pce"] = detailLspInformation.StateSyncPce
    leafs["reporting-pcc-address"] = detailLspInformation.ReportingPccAddress
    leafs["srlg-info"] = detailLspInformation.SrlgInfo
    return leafs
}

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetYangName() string { return "detail-lsp-information" }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) SetParent(parent types.Entity) { detailLspInformation.parent = parent }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetParent() types.Entity { return detailLspInformation.parent }

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetParentYangName() string { return "tunnel-detail-info" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation
// Brief LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // LSP ID. The type is interface{} with range: 0..4294967295.
    Lspid interface{}

    // Binding SID. The type is interface{} with range: 0..4294967295.
    BindingSid interface{}

    // LSP Setup Type. The type is LspSetup.
    LspSetupType interface{}

    // Operational state. The type is PcepLspState.
    OperationalState interface{}

    // Admin state. The type is LspState.
    AdministrativeState interface{}
}

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetFilter() yfilter.YFilter { return briefLspInformation.YFilter }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) SetFilter(yf yfilter.YFilter) { briefLspInformation.YFilter = yf }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "lspid" { return "Lspid" }
    if yname == "binding-sid" { return "BindingSid" }
    if yname == "lsp-setup-type" { return "LspSetupType" }
    if yname == "operational-state" { return "OperationalState" }
    if yname == "administrative-state" { return "AdministrativeState" }
    return ""
}

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetSegmentPath() string {
    return "brief-lsp-information"
}

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = briefLspInformation.SourceAddress
    leafs["destination-address"] = briefLspInformation.DestinationAddress
    leafs["tunnel-id"] = briefLspInformation.TunnelId
    leafs["lspid"] = briefLspInformation.Lspid
    leafs["binding-sid"] = briefLspInformation.BindingSid
    leafs["lsp-setup-type"] = briefLspInformation.LspSetupType
    leafs["operational-state"] = briefLspInformation.OperationalState
    leafs["administrative-state"] = briefLspInformation.AdministrativeState
    return leafs
}

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetYangName() string { return "brief-lsp-information" }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) SetParent(parent types.Entity) { briefLspInformation.parent = parent }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetParent() types.Entity { return briefLspInformation.parent }

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetParentYangName() string { return "detail-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs
// Paths
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reported Metric Type. The type is interface{} with range: 0..4294967295.
    ReportedMetricType interface{}

    // Reported Metric Value. The type is interface{} with range: 0..4294967295.
    ReportedMetricValue interface{}

    // Computed Metric Type. The type is interface{} with range: 0..4294967295.
    ComputedMetricType interface{}

    // Computed Metric Value. The type is interface{} with range: 0..4294967295.
    ComputedMetricValue interface{}

    // Computed Hop List Time. The type is interface{} with range: 0..4294967295.
    ComputedHopListTime interface{}

    // Reported RSVP path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath.
    ReportedRsvpPath []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath

    // Reported SR path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath.
    ReportedSrPath []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath

    // Computed RSVP path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath.
    ComputedRsvpPath []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath

    // Computed SR path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath.
    ComputedSrPath []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
}

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetFilter() yfilter.YFilter { return erOs.YFilter }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) SetFilter(yf yfilter.YFilter) { erOs.YFilter = yf }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetGoName(yname string) string {
    if yname == "reported-metric-type" { return "ReportedMetricType" }
    if yname == "reported-metric-value" { return "ReportedMetricValue" }
    if yname == "computed-metric-type" { return "ComputedMetricType" }
    if yname == "computed-metric-value" { return "ComputedMetricValue" }
    if yname == "computed-hop-list-time" { return "ComputedHopListTime" }
    if yname == "reported-rsvp-path" { return "ReportedRsvpPath" }
    if yname == "reported-sr-path" { return "ReportedSrPath" }
    if yname == "computed-rsvp-path" { return "ComputedRsvpPath" }
    if yname == "computed-sr-path" { return "ComputedSrPath" }
    return ""
}

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetSegmentPath() string {
    return "er-os"
}

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reported-rsvp-path" {
        for _, c := range erOs.ReportedRsvpPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath{}
        erOs.ReportedRsvpPath = append(erOs.ReportedRsvpPath, child)
        return &erOs.ReportedRsvpPath[len(erOs.ReportedRsvpPath)-1]
    }
    if childYangName == "reported-sr-path" {
        for _, c := range erOs.ReportedSrPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath{}
        erOs.ReportedSrPath = append(erOs.ReportedSrPath, child)
        return &erOs.ReportedSrPath[len(erOs.ReportedSrPath)-1]
    }
    if childYangName == "computed-rsvp-path" {
        for _, c := range erOs.ComputedRsvpPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath{}
        erOs.ComputedRsvpPath = append(erOs.ComputedRsvpPath, child)
        return &erOs.ComputedRsvpPath[len(erOs.ComputedRsvpPath)-1]
    }
    if childYangName == "computed-sr-path" {
        for _, c := range erOs.ComputedSrPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath{}
        erOs.ComputedSrPath = append(erOs.ComputedSrPath, child)
        return &erOs.ComputedSrPath[len(erOs.ComputedSrPath)-1]
    }
    return nil
}

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range erOs.ReportedRsvpPath {
        children[erOs.ReportedRsvpPath[i].GetSegmentPath()] = &erOs.ReportedRsvpPath[i]
    }
    for i := range erOs.ReportedSrPath {
        children[erOs.ReportedSrPath[i].GetSegmentPath()] = &erOs.ReportedSrPath[i]
    }
    for i := range erOs.ComputedRsvpPath {
        children[erOs.ComputedRsvpPath[i].GetSegmentPath()] = &erOs.ComputedRsvpPath[i]
    }
    for i := range erOs.ComputedSrPath {
        children[erOs.ComputedSrPath[i].GetSegmentPath()] = &erOs.ComputedSrPath[i]
    }
    return children
}

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reported-metric-type"] = erOs.ReportedMetricType
    leafs["reported-metric-value"] = erOs.ReportedMetricValue
    leafs["computed-metric-type"] = erOs.ComputedMetricType
    leafs["computed-metric-value"] = erOs.ComputedMetricValue
    leafs["computed-hop-list-time"] = erOs.ComputedHopListTime
    return leafs
}

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetBundleName() string { return "cisco_ios_xr" }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetYangName() string { return "er-os" }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) SetParent(parent types.Entity) { erOs.parent = parent }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetParent() types.Entity { return erOs.parent }

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetParentYangName() string { return "detail-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath
// Reported RSVP path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HopAddress interface{}
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetFilter() yfilter.YFilter { return reportedRsvpPath.YFilter }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) SetFilter(yf yfilter.YFilter) { reportedRsvpPath.YFilter = yf }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetGoName(yname string) string {
    if yname == "hop-address" { return "HopAddress" }
    return ""
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetSegmentPath() string {
    return "reported-rsvp-path"
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-address"] = reportedRsvpPath.HopAddress
    return leafs
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetBundleName() string { return "cisco_ios_xr" }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetYangName() string { return "reported-rsvp-path" }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) SetParent(parent types.Entity) { reportedRsvpPath.parent = parent }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetParent() types.Entity { return reportedRsvpPath.parent }

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetParentYangName() string { return "er-os" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath
// Reported SR path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddr interface{}
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetFilter() yfilter.YFilter { return reportedSrPath.YFilter }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) SetFilter(yf yfilter.YFilter) { reportedSrPath.YFilter = yf }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-addr" { return "RemoteAddr" }
    return ""
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetSegmentPath() string {
    return "reported-sr-path"
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = reportedSrPath.SidType
    leafs["mpls-label"] = reportedSrPath.MplsLabel
    leafs["local-addr"] = reportedSrPath.LocalAddr
    leafs["remote-addr"] = reportedSrPath.RemoteAddr
    return leafs
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetBundleName() string { return "cisco_ios_xr" }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetYangName() string { return "reported-sr-path" }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) SetParent(parent types.Entity) { reportedSrPath.parent = parent }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetParent() types.Entity { return reportedSrPath.parent }

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetParentYangName() string { return "er-os" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath
// Computed RSVP path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HopAddress interface{}
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetFilter() yfilter.YFilter { return computedRsvpPath.YFilter }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) SetFilter(yf yfilter.YFilter) { computedRsvpPath.YFilter = yf }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetGoName(yname string) string {
    if yname == "hop-address" { return "HopAddress" }
    return ""
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetSegmentPath() string {
    return "computed-rsvp-path"
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-address"] = computedRsvpPath.HopAddress
    return leafs
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetBundleName() string { return "cisco_ios_xr" }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetYangName() string { return "computed-rsvp-path" }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) SetParent(parent types.Entity) { computedRsvpPath.parent = parent }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetParent() types.Entity { return computedRsvpPath.parent }

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetParentYangName() string { return "er-os" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
// Computed SR path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddr interface{}
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetFilter() yfilter.YFilter { return computedSrPath.YFilter }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) SetFilter(yf yfilter.YFilter) { computedSrPath.YFilter = yf }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-addr" { return "RemoteAddr" }
    return ""
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetSegmentPath() string {
    return "computed-sr-path"
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = computedSrPath.SidType
    leafs["mpls-label"] = computedSrPath.MplsLabel
    leafs["local-addr"] = computedSrPath.LocalAddr
    leafs["remote-addr"] = computedSrPath.RemoteAddr
    return leafs
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetBundleName() string { return "cisco_ios_xr" }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetYangName() string { return "computed-sr-path" }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) SetParent(parent types.Entity) { computedSrPath.parent = parent }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetParent() types.Entity { return computedSrPath.parent }

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetParentYangName() string { return "er-os" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation
// PCEP related LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE protocol identifier. The type is interface{} with range: 0..4294967295.
    Pcepid interface{}

    // PCEP LSP delegation flag. The type is bool.
    PcepFlagD interface{}

    // PCEP LSP state-sync flag. The type is bool.
    PcepFlagS interface{}

    // PCEP LSP remove flag. The type is bool.
    PcepFlagR interface{}

    // PCEP LSP admin flag. The type is bool.
    PcepFlagA interface{}

    // PCEP LSP operation flag. The type is interface{} with range: 0..255.
    PcepFlagO interface{}

    // RSVP error info.
    RsvpError PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetFilter() yfilter.YFilter { return lsppcepInformation.YFilter }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) SetFilter(yf yfilter.YFilter) { lsppcepInformation.YFilter = yf }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetGoName(yname string) string {
    if yname == "pcepid" { return "Pcepid" }
    if yname == "pcep-flag-d" { return "PcepFlagD" }
    if yname == "pcep-flag-s" { return "PcepFlagS" }
    if yname == "pcep-flag-r" { return "PcepFlagR" }
    if yname == "pcep-flag-a" { return "PcepFlagA" }
    if yname == "pcep-flag-o" { return "PcepFlagO" }
    if yname == "rsvp-error" { return "RsvpError" }
    return ""
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetSegmentPath() string {
    return "lsppcep-information"
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvp-error" {
        return &lsppcepInformation.RsvpError
    }
    return nil
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rsvp-error"] = &lsppcepInformation.RsvpError
    return children
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcepid"] = lsppcepInformation.Pcepid
    leafs["pcep-flag-d"] = lsppcepInformation.PcepFlagD
    leafs["pcep-flag-s"] = lsppcepInformation.PcepFlagS
    leafs["pcep-flag-r"] = lsppcepInformation.PcepFlagR
    leafs["pcep-flag-a"] = lsppcepInformation.PcepFlagA
    leafs["pcep-flag-o"] = lsppcepInformation.PcepFlagO
    return leafs
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetYangName() string { return "lsppcep-information" }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) SetParent(parent types.Entity) { lsppcepInformation.parent = parent }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetParent() types.Entity { return lsppcepInformation.parent }

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetParentYangName() string { return "detail-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
// RSVP error info
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP error node address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NodeAddress interface{}

    // RSVP error flags. The type is interface{} with range: 0..255.
    ErrorFlags interface{}

    // RSVP error code. The type is interface{} with range: 0..255.
    ErrorCode interface{}

    // RSVP error value. The type is interface{} with range: 0..65535.
    ErrorValue interface{}
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetFilter() yfilter.YFilter { return rsvpError.YFilter }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) SetFilter(yf yfilter.YFilter) { rsvpError.YFilter = yf }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetGoName(yname string) string {
    if yname == "node-address" { return "NodeAddress" }
    if yname == "error-flags" { return "ErrorFlags" }
    if yname == "error-code" { return "ErrorCode" }
    if yname == "error-value" { return "ErrorValue" }
    return ""
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetSegmentPath() string {
    return "rsvp-error"
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-address"] = rsvpError.NodeAddress
    leafs["error-flags"] = rsvpError.ErrorFlags
    leafs["error-code"] = rsvpError.ErrorCode
    leafs["error-value"] = rsvpError.ErrorValue
    return leafs
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetBundleName() string { return "cisco_ios_xr" }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetYangName() string { return "rsvp-error" }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) SetParent(parent types.Entity) { rsvpError.parent = parent }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetParent() types.Entity { return rsvpError.parent }

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetParentYangName() string { return "lsppcep-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo
// LSP association information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AssociationSource interface{}
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetFilter() yfilter.YFilter { return lspAssociationInfo.YFilter }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) SetFilter(yf yfilter.YFilter) { lspAssociationInfo.YFilter = yf }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetGoName(yname string) string {
    if yname == "association-type" { return "AssociationType" }
    if yname == "association-id" { return "AssociationId" }
    if yname == "association-source" { return "AssociationSource" }
    return ""
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetSegmentPath() string {
    return "lsp-association-info"
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["association-type"] = lspAssociationInfo.AssociationType
    leafs["association-id"] = lspAssociationInfo.AssociationId
    leafs["association-source"] = lspAssociationInfo.AssociationSource
    return leafs
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetBundleName() string { return "cisco_ios_xr" }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetYangName() string { return "lsp-association-info" }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) SetParent(parent types.Entity) { lspAssociationInfo.parent = parent }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetParent() types.Entity { return lspAssociationInfo.parent }

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetParentYangName() string { return "detail-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes
// LSP attributes
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Affinity exclude any. The type is interface{} with range: 0..4294967295.
    AffinityExcludeAny interface{}

    // Affinity include any. The type is interface{} with range: 0..4294967295.
    AffinityIncludeAny interface{}

    // Affinity include all. The type is interface{} with range: 0..4294967295.
    AffinityIncludeAll interface{}

    // Setup Priority. The type is interface{} with range: 0..255.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..255.
    HoldPriority interface{}

    // True, if local protection is desired. The type is bool.
    LocalProtection interface{}
}

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetFilter() yfilter.YFilter { return lspAttributes.YFilter }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) SetFilter(yf yfilter.YFilter) { lspAttributes.YFilter = yf }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetGoName(yname string) string {
    if yname == "affinity-exclude-any" { return "AffinityExcludeAny" }
    if yname == "affinity-include-any" { return "AffinityIncludeAny" }
    if yname == "affinity-include-all" { return "AffinityIncludeAll" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "local-protection" { return "LocalProtection" }
    return ""
}

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetSegmentPath() string {
    return "lsp-attributes"
}

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["affinity-exclude-any"] = lspAttributes.AffinityExcludeAny
    leafs["affinity-include-any"] = lspAttributes.AffinityIncludeAny
    leafs["affinity-include-all"] = lspAttributes.AffinityIncludeAll
    leafs["setup-priority"] = lspAttributes.SetupPriority
    leafs["hold-priority"] = lspAttributes.HoldPriority
    leafs["local-protection"] = lspAttributes.LocalProtection
    return leafs
}

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetYangName() string { return "lsp-attributes" }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) SetParent(parent types.Entity) { lspAttributes.parent = parent }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetParent() types.Entity { return lspAttributes.parent }

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetParentYangName() string { return "detail-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
// RRO
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RRO Type. The type is PceRro.
    RroType interface{}

    // IPv4 address of RRO. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // MPLS label of RRO. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // RRO Flags. The type is interface{} with range: 0..255.
    Flags interface{}

    // Segment Routing RRO info.
    SrRro PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetFilter() yfilter.YFilter { return rro.YFilter }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) SetFilter(yf yfilter.YFilter) { rro.YFilter = yf }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetGoName(yname string) string {
    if yname == "rro-type" { return "RroType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "flags" { return "Flags" }
    if yname == "sr-rro" { return "SrRro" }
    return ""
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetSegmentPath() string {
    return "rro"
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sr-rro" {
        return &rro.SrRro
    }
    return nil
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sr-rro"] = &rro.SrRro
    return children
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rro-type"] = rro.RroType
    leafs["ipv4-address"] = rro.Ipv4Address
    leafs["mpls-label"] = rro.MplsLabel
    leafs["flags"] = rro.Flags
    return leafs
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetBundleName() string { return "cisco_ios_xr" }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetYangName() string { return "rro" }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) SetParent(parent types.Entity) { rro.parent = parent }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetParent() types.Entity { return rro.parent }

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetParentYangName() string { return "detail-lsp-information" }

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
// Segment Routing RRO info
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddr interface{}
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetFilter() yfilter.YFilter { return srRro.YFilter }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) SetFilter(yf yfilter.YFilter) { srRro.YFilter = yf }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-addr" { return "RemoteAddr" }
    return ""
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetSegmentPath() string {
    return "sr-rro"
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = srRro.SidType
    leafs["mpls-label"] = srRro.MplsLabel
    leafs["local-addr"] = srRro.LocalAddr
    leafs["remote-addr"] = srRro.RemoteAddr
    return leafs
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetBundleName() string { return "cisco_ios_xr" }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetYangName() string { return "sr-rro" }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) SetParent(parent types.Entity) { srRro.parent = parent }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetParent() types.Entity { return srRro.parent }

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetParentYangName() string { return "rro" }

// PcePeer
// pce peer
type PcePeer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed peers database in XTC.
    PeerDetailInfos PcePeer_PeerDetailInfos

    // Peers database in XTC.
    PeerInfos PcePeer_PeerInfos
}

func (pcePeer *PcePeer) GetFilter() yfilter.YFilter { return pcePeer.YFilter }

func (pcePeer *PcePeer) SetFilter(yf yfilter.YFilter) { pcePeer.YFilter = yf }

func (pcePeer *PcePeer) GetGoName(yname string) string {
    if yname == "peer-detail-infos" { return "PeerDetailInfos" }
    if yname == "peer-infos" { return "PeerInfos" }
    return ""
}

func (pcePeer *PcePeer) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-xtc-oper:pce-peer"
}

func (pcePeer *PcePeer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-detail-infos" {
        return &pcePeer.PeerDetailInfos
    }
    if childYangName == "peer-infos" {
        return &pcePeer.PeerInfos
    }
    return nil
}

func (pcePeer *PcePeer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["peer-detail-infos"] = &pcePeer.PeerDetailInfos
    children["peer-infos"] = &pcePeer.PeerInfos
    return children
}

func (pcePeer *PcePeer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pcePeer *PcePeer) GetBundleName() string { return "cisco_ios_xr" }

func (pcePeer *PcePeer) GetYangName() string { return "pce-peer" }

func (pcePeer *PcePeer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcePeer *PcePeer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcePeer *PcePeer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcePeer *PcePeer) SetParent(parent types.Entity) { pcePeer.parent = parent }

func (pcePeer *PcePeer) GetParent() types.Entity { return pcePeer.parent }

func (pcePeer *PcePeer) GetParentYangName() string { return "Cisco-IOS-XR-infra-xtc-oper" }

// PcePeer_PeerDetailInfos
// Detailed peers database in XTC
type PcePeer_PeerDetailInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed PCE peer information. The type is slice of
    // PcePeer_PeerDetailInfos_PeerDetailInfo.
    PeerDetailInfo []PcePeer_PeerDetailInfos_PeerDetailInfo
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetFilter() yfilter.YFilter { return peerDetailInfos.YFilter }

func (peerDetailInfos *PcePeer_PeerDetailInfos) SetFilter(yf yfilter.YFilter) { peerDetailInfos.YFilter = yf }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetGoName(yname string) string {
    if yname == "peer-detail-info" { return "PeerDetailInfo" }
    return ""
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetSegmentPath() string {
    return "peer-detail-infos"
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-detail-info" {
        for _, c := range peerDetailInfos.PeerDetailInfo {
            if peerDetailInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PcePeer_PeerDetailInfos_PeerDetailInfo{}
        peerDetailInfos.PeerDetailInfo = append(peerDetailInfos.PeerDetailInfo, child)
        return &peerDetailInfos.PeerDetailInfo[len(peerDetailInfos.PeerDetailInfo)-1]
    }
    return nil
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerDetailInfos.PeerDetailInfo {
        children[peerDetailInfos.PeerDetailInfo[i].GetSegmentPath()] = &peerDetailInfos.PeerDetailInfo[i]
    }
    return children
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetBundleName() string { return "cisco_ios_xr" }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetYangName() string { return "peer-detail-infos" }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerDetailInfos *PcePeer_PeerDetailInfos) SetParent(parent types.Entity) { peerDetailInfos.parent = parent }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetParent() types.Entity { return peerDetailInfos.parent }

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetParentYangName() string { return "pce-peer" }

// PcePeer_PeerDetailInfos_PeerDetailInfo
// Detailed PCE peer information
type PcePeer_PeerDetailInfos_PeerDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Detailed PCE protocol information.
    DetailPcepInformation PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetFilter() yfilter.YFilter { return peerDetailInfo.YFilter }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) SetFilter(yf yfilter.YFilter) { peerDetailInfo.YFilter = yf }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-address-xr" { return "PeerAddressXr" }
    if yname == "peer-protocol" { return "PeerProtocol" }
    if yname == "detail-pcep-information" { return "DetailPcepInformation" }
    return ""
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetSegmentPath() string {
    return "peer-detail-info" + "[peer-address='" + fmt.Sprintf("%v", peerDetailInfo.PeerAddress) + "']"
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-pcep-information" {
        return &peerDetailInfo.DetailPcepInformation
    }
    return nil
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail-pcep-information"] = &peerDetailInfo.DetailPcepInformation
    return children
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peerDetailInfo.PeerAddress
    leafs["peer-address-xr"] = peerDetailInfo.PeerAddressXr
    leafs["peer-protocol"] = peerDetailInfo.PeerProtocol
    return leafs
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetYangName() string { return "peer-detail-info" }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) SetParent(parent types.Entity) { peerDetailInfo.parent = parent }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetParent() types.Entity { return peerDetailInfo.parent }

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetParentYangName() string { return "peer-detail-infos" }

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
// Detailed PCE protocol information
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error (for display only). The type is string.
    Error interface{}

    // Speaker Entity ID. The type is string.
    SpeakerId interface{}

    // PCEP Up Time. The type is interface{} with range: 0..4294967295.
    PcepUpTime interface{}

    // Keepalive count. The type is interface{} with range: 0..4294967295.
    Keepalives interface{}

    // MD5 Authentication Enabled. The type is bool.
    Md5Enabled interface{}

    // Keychain based Authentication Enabled. The type is bool.
    KeychainEnabled interface{}

    // Negotiated KA. The type is interface{} with range: 0..4294967295.
    NegotiatedLocalKeepalive interface{}

    // Negotiated KA. The type is interface{} with range: 0..4294967295.
    NegotiatedRemoteKeepalive interface{}

    // Negotiated DT. The type is interface{} with range: 0..4294967295.
    NegotiatedDeadTime interface{}

    // PCEReq Rx. The type is interface{} with range: 0..4294967295.
    PceRequestRx interface{}

    // PCEReq Tx. The type is interface{} with range: 0..4294967295.
    PceRequestTx interface{}

    // PCERep Rx. The type is interface{} with range: 0..4294967295.
    PceReplyRx interface{}

    // PCERep Tx. The type is interface{} with range: 0..4294967295.
    PceReplyTx interface{}

    // PCEErr Rx. The type is interface{} with range: 0..4294967295.
    PceErrorRx interface{}

    // PCEErr Tx. The type is interface{} with range: 0..4294967295.
    PceErrorTx interface{}

    // PCEOpen Tx. The type is interface{} with range: 0..4294967295.
    PceOpenTx interface{}

    // PCEOpen Rx. The type is interface{} with range: 0..4294967295.
    PceOpenRx interface{}

    // PCERpt Rx. The type is interface{} with range: 0..4294967295.
    PceReportRx interface{}

    // PCERpt Tx. The type is interface{} with range: 0..4294967295.
    PceReportTx interface{}

    // PCEUpd Rx. The type is interface{} with range: 0..4294967295.
    PceUpdateRx interface{}

    // PCEUpd Tx. The type is interface{} with range: 0..4294967295.
    PceUpdateTx interface{}

    // PCEInit Rx. The type is interface{} with range: 0..4294967295.
    PceInitiateRx interface{}

    // PCEInit Tx. The type is interface{} with range: 0..4294967295.
    PceInitiateTx interface{}

    // PCE Keepalive Tx. The type is interface{} with range:
    // 0..18446744073709551615.
    PceKeepaliveTx interface{}

    // PCE Keepalive Rx. The type is interface{} with range:
    // 0..18446744073709551615.
    PceKeepaliveRx interface{}

    // Local PCEP session ID. The type is interface{} with range: 0..255.
    LocalSessionId interface{}

    // Remote PCEP session ID. The type is interface{} with range: 0..255.
    RemoteSessionId interface{}

    // Minimum keepalive interval for the peer. The type is interface{} with
    // range: 0..255.
    MinimumKeepaliveInterval interface{}

    // Maximum dead interval for the peer. The type is interface{} with range:
    // 0..255.
    MaximumDeadInterval interface{}

    // Brief PCE protocol information.
    BriefPcepInformation PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation

    // Last PCError received.
    LastErrorRx PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx

    // Last PCError sent.
    LastErrorTx PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx
}

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetFilter() yfilter.YFilter { return detailPcepInformation.YFilter }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) SetFilter(yf yfilter.YFilter) { detailPcepInformation.YFilter = yf }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    if yname == "speaker-id" { return "SpeakerId" }
    if yname == "pcep-up-time" { return "PcepUpTime" }
    if yname == "keepalives" { return "Keepalives" }
    if yname == "md5-enabled" { return "Md5Enabled" }
    if yname == "keychain-enabled" { return "KeychainEnabled" }
    if yname == "negotiated-local-keepalive" { return "NegotiatedLocalKeepalive" }
    if yname == "negotiated-remote-keepalive" { return "NegotiatedRemoteKeepalive" }
    if yname == "negotiated-dead-time" { return "NegotiatedDeadTime" }
    if yname == "pce-request-rx" { return "PceRequestRx" }
    if yname == "pce-request-tx" { return "PceRequestTx" }
    if yname == "pce-reply-rx" { return "PceReplyRx" }
    if yname == "pce-reply-tx" { return "PceReplyTx" }
    if yname == "pce-error-rx" { return "PceErrorRx" }
    if yname == "pce-error-tx" { return "PceErrorTx" }
    if yname == "pce-open-tx" { return "PceOpenTx" }
    if yname == "pce-open-rx" { return "PceOpenRx" }
    if yname == "pce-report-rx" { return "PceReportRx" }
    if yname == "pce-report-tx" { return "PceReportTx" }
    if yname == "pce-update-rx" { return "PceUpdateRx" }
    if yname == "pce-update-tx" { return "PceUpdateTx" }
    if yname == "pce-initiate-rx" { return "PceInitiateRx" }
    if yname == "pce-initiate-tx" { return "PceInitiateTx" }
    if yname == "pce-keepalive-tx" { return "PceKeepaliveTx" }
    if yname == "pce-keepalive-rx" { return "PceKeepaliveRx" }
    if yname == "local-session-id" { return "LocalSessionId" }
    if yname == "remote-session-id" { return "RemoteSessionId" }
    if yname == "minimum-keepalive-interval" { return "MinimumKeepaliveInterval" }
    if yname == "maximum-dead-interval" { return "MaximumDeadInterval" }
    if yname == "brief-pcep-information" { return "BriefPcepInformation" }
    if yname == "last-error-rx" { return "LastErrorRx" }
    if yname == "last-error-tx" { return "LastErrorTx" }
    return ""
}

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetSegmentPath() string {
    return "detail-pcep-information"
}

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-pcep-information" {
        return &detailPcepInformation.BriefPcepInformation
    }
    if childYangName == "last-error-rx" {
        return &detailPcepInformation.LastErrorRx
    }
    if childYangName == "last-error-tx" {
        return &detailPcepInformation.LastErrorTx
    }
    return nil
}

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-pcep-information"] = &detailPcepInformation.BriefPcepInformation
    children["last-error-rx"] = &detailPcepInformation.LastErrorRx
    children["last-error-tx"] = &detailPcepInformation.LastErrorTx
    return children
}

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["error"] = detailPcepInformation.Error
    leafs["speaker-id"] = detailPcepInformation.SpeakerId
    leafs["pcep-up-time"] = detailPcepInformation.PcepUpTime
    leafs["keepalives"] = detailPcepInformation.Keepalives
    leafs["md5-enabled"] = detailPcepInformation.Md5Enabled
    leafs["keychain-enabled"] = detailPcepInformation.KeychainEnabled
    leafs["negotiated-local-keepalive"] = detailPcepInformation.NegotiatedLocalKeepalive
    leafs["negotiated-remote-keepalive"] = detailPcepInformation.NegotiatedRemoteKeepalive
    leafs["negotiated-dead-time"] = detailPcepInformation.NegotiatedDeadTime
    leafs["pce-request-rx"] = detailPcepInformation.PceRequestRx
    leafs["pce-request-tx"] = detailPcepInformation.PceRequestTx
    leafs["pce-reply-rx"] = detailPcepInformation.PceReplyRx
    leafs["pce-reply-tx"] = detailPcepInformation.PceReplyTx
    leafs["pce-error-rx"] = detailPcepInformation.PceErrorRx
    leafs["pce-error-tx"] = detailPcepInformation.PceErrorTx
    leafs["pce-open-tx"] = detailPcepInformation.PceOpenTx
    leafs["pce-open-rx"] = detailPcepInformation.PceOpenRx
    leafs["pce-report-rx"] = detailPcepInformation.PceReportRx
    leafs["pce-report-tx"] = detailPcepInformation.PceReportTx
    leafs["pce-update-rx"] = detailPcepInformation.PceUpdateRx
    leafs["pce-update-tx"] = detailPcepInformation.PceUpdateTx
    leafs["pce-initiate-rx"] = detailPcepInformation.PceInitiateRx
    leafs["pce-initiate-tx"] = detailPcepInformation.PceInitiateTx
    leafs["pce-keepalive-tx"] = detailPcepInformation.PceKeepaliveTx
    leafs["pce-keepalive-rx"] = detailPcepInformation.PceKeepaliveRx
    leafs["local-session-id"] = detailPcepInformation.LocalSessionId
    leafs["remote-session-id"] = detailPcepInformation.RemoteSessionId
    leafs["minimum-keepalive-interval"] = detailPcepInformation.MinimumKeepaliveInterval
    leafs["maximum-dead-interval"] = detailPcepInformation.MaximumDeadInterval
    return leafs
}

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetYangName() string { return "detail-pcep-information" }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) SetParent(parent types.Entity) { detailPcepInformation.parent = parent }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetParent() types.Entity { return detailPcepInformation.parent }

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetParentYangName() string { return "peer-detail-info" }

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation
// Brief PCE protocol information
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP State. The type is PcepState.
    PcepState interface{}

    // Stateful. The type is bool.
    Stateful interface{}

    // Update capability. The type is bool.
    CapabilityUpdate interface{}

    // Instantiation capability. The type is bool.
    CapabilityInstantiate interface{}

    // Segment Routing capability. The type is bool.
    CapabilitySegmentRouting interface{}

    // Triggered Synchronization capability. The type is bool.
    CapabilityTriggeredSync interface{}

    // DB version capability. The type is bool.
    CapabilityDbVersion interface{}

    // Delta Synchronization capability. The type is bool.
    CapabilityDeltaSync interface{}
}

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetFilter() yfilter.YFilter { return briefPcepInformation.YFilter }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) SetFilter(yf yfilter.YFilter) { briefPcepInformation.YFilter = yf }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetGoName(yname string) string {
    if yname == "pcep-state" { return "PcepState" }
    if yname == "stateful" { return "Stateful" }
    if yname == "capability-update" { return "CapabilityUpdate" }
    if yname == "capability-instantiate" { return "CapabilityInstantiate" }
    if yname == "capability-segment-routing" { return "CapabilitySegmentRouting" }
    if yname == "capability-triggered-sync" { return "CapabilityTriggeredSync" }
    if yname == "capability-db-version" { return "CapabilityDbVersion" }
    if yname == "capability-delta-sync" { return "CapabilityDeltaSync" }
    return ""
}

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetSegmentPath() string {
    return "brief-pcep-information"
}

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcep-state"] = briefPcepInformation.PcepState
    leafs["stateful"] = briefPcepInformation.Stateful
    leafs["capability-update"] = briefPcepInformation.CapabilityUpdate
    leafs["capability-instantiate"] = briefPcepInformation.CapabilityInstantiate
    leafs["capability-segment-routing"] = briefPcepInformation.CapabilitySegmentRouting
    leafs["capability-triggered-sync"] = briefPcepInformation.CapabilityTriggeredSync
    leafs["capability-db-version"] = briefPcepInformation.CapabilityDbVersion
    leafs["capability-delta-sync"] = briefPcepInformation.CapabilityDeltaSync
    return leafs
}

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetYangName() string { return "brief-pcep-information" }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) SetParent(parent types.Entity) { briefPcepInformation.parent = parent }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetParent() types.Entity { return briefPcepInformation.parent }

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetParentYangName() string { return "detail-pcep-information" }

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx
// Last PCError received
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetFilter() yfilter.YFilter { return lastErrorRx.YFilter }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) SetFilter(yf yfilter.YFilter) { lastErrorRx.YFilter = yf }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetGoName(yname string) string {
    if yname == "pc-error-type" { return "PcErrorType" }
    if yname == "pc-error-value" { return "PcErrorValue" }
    return ""
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetSegmentPath() string {
    return "last-error-rx"
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pc-error-type"] = lastErrorRx.PcErrorType
    leafs["pc-error-value"] = lastErrorRx.PcErrorValue
    return leafs
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetBundleName() string { return "cisco_ios_xr" }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetYangName() string { return "last-error-rx" }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) SetParent(parent types.Entity) { lastErrorRx.parent = parent }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetParent() types.Entity { return lastErrorRx.parent }

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetParentYangName() string { return "detail-pcep-information" }

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx
// Last PCError sent
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetFilter() yfilter.YFilter { return lastErrorTx.YFilter }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) SetFilter(yf yfilter.YFilter) { lastErrorTx.YFilter = yf }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetGoName(yname string) string {
    if yname == "pc-error-type" { return "PcErrorType" }
    if yname == "pc-error-value" { return "PcErrorValue" }
    return ""
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetSegmentPath() string {
    return "last-error-tx"
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pc-error-type"] = lastErrorTx.PcErrorType
    leafs["pc-error-value"] = lastErrorTx.PcErrorValue
    return leafs
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetBundleName() string { return "cisco_ios_xr" }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetYangName() string { return "last-error-tx" }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) SetParent(parent types.Entity) { lastErrorTx.parent = parent }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetParent() types.Entity { return lastErrorTx.parent }

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetParentYangName() string { return "detail-pcep-information" }

// PcePeer_PeerInfos
// Peers database in XTC
type PcePeer_PeerInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE peer information. The type is slice of PcePeer_PeerInfos_PeerInfo.
    PeerInfo []PcePeer_PeerInfos_PeerInfo
}

func (peerInfos *PcePeer_PeerInfos) GetFilter() yfilter.YFilter { return peerInfos.YFilter }

func (peerInfos *PcePeer_PeerInfos) SetFilter(yf yfilter.YFilter) { peerInfos.YFilter = yf }

func (peerInfos *PcePeer_PeerInfos) GetGoName(yname string) string {
    if yname == "peer-info" { return "PeerInfo" }
    return ""
}

func (peerInfos *PcePeer_PeerInfos) GetSegmentPath() string {
    return "peer-infos"
}

func (peerInfos *PcePeer_PeerInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-info" {
        for _, c := range peerInfos.PeerInfo {
            if peerInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PcePeer_PeerInfos_PeerInfo{}
        peerInfos.PeerInfo = append(peerInfos.PeerInfo, child)
        return &peerInfos.PeerInfo[len(peerInfos.PeerInfo)-1]
    }
    return nil
}

func (peerInfos *PcePeer_PeerInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerInfos.PeerInfo {
        children[peerInfos.PeerInfo[i].GetSegmentPath()] = &peerInfos.PeerInfo[i]
    }
    return children
}

func (peerInfos *PcePeer_PeerInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerInfos *PcePeer_PeerInfos) GetBundleName() string { return "cisco_ios_xr" }

func (peerInfos *PcePeer_PeerInfos) GetYangName() string { return "peer-infos" }

func (peerInfos *PcePeer_PeerInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerInfos *PcePeer_PeerInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerInfos *PcePeer_PeerInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerInfos *PcePeer_PeerInfos) SetParent(parent types.Entity) { peerInfos.parent = parent }

func (peerInfos *PcePeer_PeerInfos) GetParent() types.Entity { return peerInfos.parent }

func (peerInfos *PcePeer_PeerInfos) GetParentYangName() string { return "pce-peer" }

// PcePeer_PeerInfos_PeerInfo
// PCE peer information
type PcePeer_PeerInfos_PeerInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // PCE protocol information.
    BriefPcepInformation PcePeer_PeerInfos_PeerInfo_BriefPcepInformation
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetFilter() yfilter.YFilter { return peerInfo.YFilter }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) SetFilter(yf yfilter.YFilter) { peerInfo.YFilter = yf }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-address-xr" { return "PeerAddressXr" }
    if yname == "peer-protocol" { return "PeerProtocol" }
    if yname == "brief-pcep-information" { return "BriefPcepInformation" }
    return ""
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetSegmentPath() string {
    return "peer-info" + "[peer-address='" + fmt.Sprintf("%v", peerInfo.PeerAddress) + "']"
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-pcep-information" {
        return &peerInfo.BriefPcepInformation
    }
    return nil
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-pcep-information"] = &peerInfo.BriefPcepInformation
    return children
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peerInfo.PeerAddress
    leafs["peer-address-xr"] = peerInfo.PeerAddressXr
    leafs["peer-protocol"] = peerInfo.PeerProtocol
    return leafs
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetYangName() string { return "peer-info" }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) SetParent(parent types.Entity) { peerInfo.parent = parent }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetParent() types.Entity { return peerInfo.parent }

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetParentYangName() string { return "peer-infos" }

// PcePeer_PeerInfos_PeerInfo_BriefPcepInformation
// PCE protocol information
type PcePeer_PeerInfos_PeerInfo_BriefPcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP State. The type is PcepState.
    PcepState interface{}

    // Stateful. The type is bool.
    Stateful interface{}

    // Update capability. The type is bool.
    CapabilityUpdate interface{}

    // Instantiation capability. The type is bool.
    CapabilityInstantiate interface{}

    // Segment Routing capability. The type is bool.
    CapabilitySegmentRouting interface{}

    // Triggered Synchronization capability. The type is bool.
    CapabilityTriggeredSync interface{}

    // DB version capability. The type is bool.
    CapabilityDbVersion interface{}

    // Delta Synchronization capability. The type is bool.
    CapabilityDeltaSync interface{}
}

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetFilter() yfilter.YFilter { return briefPcepInformation.YFilter }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) SetFilter(yf yfilter.YFilter) { briefPcepInformation.YFilter = yf }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetGoName(yname string) string {
    if yname == "pcep-state" { return "PcepState" }
    if yname == "stateful" { return "Stateful" }
    if yname == "capability-update" { return "CapabilityUpdate" }
    if yname == "capability-instantiate" { return "CapabilityInstantiate" }
    if yname == "capability-segment-routing" { return "CapabilitySegmentRouting" }
    if yname == "capability-triggered-sync" { return "CapabilityTriggeredSync" }
    if yname == "capability-db-version" { return "CapabilityDbVersion" }
    if yname == "capability-delta-sync" { return "CapabilityDeltaSync" }
    return ""
}

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetSegmentPath() string {
    return "brief-pcep-information"
}

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcep-state"] = briefPcepInformation.PcepState
    leafs["stateful"] = briefPcepInformation.Stateful
    leafs["capability-update"] = briefPcepInformation.CapabilityUpdate
    leafs["capability-instantiate"] = briefPcepInformation.CapabilityInstantiate
    leafs["capability-segment-routing"] = briefPcepInformation.CapabilitySegmentRouting
    leafs["capability-triggered-sync"] = briefPcepInformation.CapabilityTriggeredSync
    leafs["capability-db-version"] = briefPcepInformation.CapabilityDbVersion
    leafs["capability-delta-sync"] = briefPcepInformation.CapabilityDeltaSync
    return leafs
}

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetYangName() string { return "brief-pcep-information" }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) SetParent(parent types.Entity) { briefPcepInformation.parent = parent }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetParent() types.Entity { return briefPcepInformation.parent }

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetParentYangName() string { return "peer-info" }

// PceTopology
// pce topology
type PceTopology struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node summary database in XTC.
    TopologySummary PceTopology_TopologySummary

    // Node database in XTC.
    TopologyNodes PceTopology_TopologyNodes

    // Prefixes database in XTC.
    PrefixInfos PceTopology_PrefixInfos
}

func (pceTopology *PceTopology) GetFilter() yfilter.YFilter { return pceTopology.YFilter }

func (pceTopology *PceTopology) SetFilter(yf yfilter.YFilter) { pceTopology.YFilter = yf }

func (pceTopology *PceTopology) GetGoName(yname string) string {
    if yname == "topology-summary" { return "TopologySummary" }
    if yname == "topology-nodes" { return "TopologyNodes" }
    if yname == "prefix-infos" { return "PrefixInfos" }
    return ""
}

func (pceTopology *PceTopology) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-xtc-oper:pce-topology"
}

func (pceTopology *PceTopology) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology-summary" {
        return &pceTopology.TopologySummary
    }
    if childYangName == "topology-nodes" {
        return &pceTopology.TopologyNodes
    }
    if childYangName == "prefix-infos" {
        return &pceTopology.PrefixInfos
    }
    return nil
}

func (pceTopology *PceTopology) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["topology-summary"] = &pceTopology.TopologySummary
    children["topology-nodes"] = &pceTopology.TopologyNodes
    children["prefix-infos"] = &pceTopology.PrefixInfos
    return children
}

func (pceTopology *PceTopology) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pceTopology *PceTopology) GetBundleName() string { return "cisco_ios_xr" }

func (pceTopology *PceTopology) GetYangName() string { return "pce-topology" }

func (pceTopology *PceTopology) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pceTopology *PceTopology) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pceTopology *PceTopology) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pceTopology *PceTopology) SetParent(parent types.Entity) { pceTopology.parent = parent }

func (pceTopology *PceTopology) GetParent() types.Entity { return pceTopology.parent }

func (pceTopology *PceTopology) GetParentYangName() string { return "Cisco-IOS-XR-infra-xtc-oper" }

// PceTopology_TopologySummary
// Node summary database in XTC
type PceTopology_TopologySummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of PCE nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // Number of lookup nodes. The type is interface{} with range: 0..4294967295.
    LookupNodes interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    Prefixes interface{}

    // Number of total prefix SIDs. The type is interface{} with range:
    // 0..4294967295.
    PrefixSids interface{}

    // Number of reguar prefix SIDs. The type is interface{} with range:
    // 0..4294967295.
    RegularPrefixSids interface{}

    // Number of strict prefix SIDs. The type is interface{} with range:
    // 0..4294967295.
    StrictPrefixSids interface{}

    // Number of links. The type is interface{} with range: 0..4294967295.
    Links interface{}

    // Number of EPE links. The type is interface{} with range: 0..4294967295.
    EpeLinks interface{}

    // Number of total adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    AdjacencySids interface{}

    // Number of total EPE SIDs. The type is interface{} with range:
    // 0..4294967295.
    Epesids interface{}

    // Number of protected adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    ProtectedAdjacencySids interface{}

    // Number of unprotected adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    UnProtectedAdjacencySids interface{}

    // True if topology is consistent. The type is bool.
    TopologyConsistent interface{}

    // Statistics on topology update.
    StatsTopologyUpdate PceTopology_TopologySummary_StatsTopologyUpdate
}

func (topologySummary *PceTopology_TopologySummary) GetFilter() yfilter.YFilter { return topologySummary.YFilter }

func (topologySummary *PceTopology_TopologySummary) SetFilter(yf yfilter.YFilter) { topologySummary.YFilter = yf }

func (topologySummary *PceTopology_TopologySummary) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "lookup-nodes" { return "LookupNodes" }
    if yname == "prefixes" { return "Prefixes" }
    if yname == "prefix-sids" { return "PrefixSids" }
    if yname == "regular-prefix-sids" { return "RegularPrefixSids" }
    if yname == "strict-prefix-sids" { return "StrictPrefixSids" }
    if yname == "links" { return "Links" }
    if yname == "epe-links" { return "EpeLinks" }
    if yname == "adjacency-sids" { return "AdjacencySids" }
    if yname == "epesids" { return "Epesids" }
    if yname == "protected-adjacency-sids" { return "ProtectedAdjacencySids" }
    if yname == "un-protected-adjacency-sids" { return "UnProtectedAdjacencySids" }
    if yname == "topology-consistent" { return "TopologyConsistent" }
    if yname == "stats-topology-update" { return "StatsTopologyUpdate" }
    return ""
}

func (topologySummary *PceTopology_TopologySummary) GetSegmentPath() string {
    return "topology-summary"
}

func (topologySummary *PceTopology_TopologySummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats-topology-update" {
        return &topologySummary.StatsTopologyUpdate
    }
    return nil
}

func (topologySummary *PceTopology_TopologySummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats-topology-update"] = &topologySummary.StatsTopologyUpdate
    return children
}

func (topologySummary *PceTopology_TopologySummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nodes"] = topologySummary.Nodes
    leafs["lookup-nodes"] = topologySummary.LookupNodes
    leafs["prefixes"] = topologySummary.Prefixes
    leafs["prefix-sids"] = topologySummary.PrefixSids
    leafs["regular-prefix-sids"] = topologySummary.RegularPrefixSids
    leafs["strict-prefix-sids"] = topologySummary.StrictPrefixSids
    leafs["links"] = topologySummary.Links
    leafs["epe-links"] = topologySummary.EpeLinks
    leafs["adjacency-sids"] = topologySummary.AdjacencySids
    leafs["epesids"] = topologySummary.Epesids
    leafs["protected-adjacency-sids"] = topologySummary.ProtectedAdjacencySids
    leafs["un-protected-adjacency-sids"] = topologySummary.UnProtectedAdjacencySids
    leafs["topology-consistent"] = topologySummary.TopologyConsistent
    return leafs
}

func (topologySummary *PceTopology_TopologySummary) GetBundleName() string { return "cisco_ios_xr" }

func (topologySummary *PceTopology_TopologySummary) GetYangName() string { return "topology-summary" }

func (topologySummary *PceTopology_TopologySummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologySummary *PceTopology_TopologySummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologySummary *PceTopology_TopologySummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologySummary *PceTopology_TopologySummary) SetParent(parent types.Entity) { topologySummary.parent = parent }

func (topologySummary *PceTopology_TopologySummary) GetParent() types.Entity { return topologySummary.parent }

func (topologySummary *PceTopology_TopologySummary) GetParentYangName() string { return "pce-topology" }

// PceTopology_TopologySummary_StatsTopologyUpdate
// Statistics on topology update
type PceTopology_TopologySummary_StatsTopologyUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of nodes added. The type is interface{} with range: 0..4294967295.
    NumNodesAdded interface{}

    // Number of nodes deleted. The type is interface{} with range: 0..4294967295.
    NumNodesDeleted interface{}

    // Number of links added. The type is interface{} with range: 0..4294967295.
    NumLinksAdded interface{}

    // Number of links deleted. The type is interface{} with range: 0..4294967295.
    NumLinksDeleted interface{}

    // Number of prefixes added. The type is interface{} with range:
    // 0..4294967295.
    NumPrefixesAdded interface{}

    // Number of prefixes deleted. The type is interface{} with range:
    // 0..4294967295.
    NumPrefixesDeleted interface{}
}

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetFilter() yfilter.YFilter { return statsTopologyUpdate.YFilter }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) SetFilter(yf yfilter.YFilter) { statsTopologyUpdate.YFilter = yf }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetGoName(yname string) string {
    if yname == "num-nodes-added" { return "NumNodesAdded" }
    if yname == "num-nodes-deleted" { return "NumNodesDeleted" }
    if yname == "num-links-added" { return "NumLinksAdded" }
    if yname == "num-links-deleted" { return "NumLinksDeleted" }
    if yname == "num-prefixes-added" { return "NumPrefixesAdded" }
    if yname == "num-prefixes-deleted" { return "NumPrefixesDeleted" }
    return ""
}

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetSegmentPath() string {
    return "stats-topology-update"
}

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-nodes-added"] = statsTopologyUpdate.NumNodesAdded
    leafs["num-nodes-deleted"] = statsTopologyUpdate.NumNodesDeleted
    leafs["num-links-added"] = statsTopologyUpdate.NumLinksAdded
    leafs["num-links-deleted"] = statsTopologyUpdate.NumLinksDeleted
    leafs["num-prefixes-added"] = statsTopologyUpdate.NumPrefixesAdded
    leafs["num-prefixes-deleted"] = statsTopologyUpdate.NumPrefixesDeleted
    return leafs
}

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetYangName() string { return "stats-topology-update" }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) SetParent(parent types.Entity) { statsTopologyUpdate.parent = parent }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetParent() types.Entity { return statsTopologyUpdate.parent }

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetParentYangName() string { return "topology-summary" }

// PceTopology_TopologyNodes
// Node database in XTC
type PceTopology_TopologyNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode.
    TopologyNode []PceTopology_TopologyNodes_TopologyNode
}

func (topologyNodes *PceTopology_TopologyNodes) GetFilter() yfilter.YFilter { return topologyNodes.YFilter }

func (topologyNodes *PceTopology_TopologyNodes) SetFilter(yf yfilter.YFilter) { topologyNodes.YFilter = yf }

func (topologyNodes *PceTopology_TopologyNodes) GetGoName(yname string) string {
    if yname == "topology-node" { return "TopologyNode" }
    return ""
}

func (topologyNodes *PceTopology_TopologyNodes) GetSegmentPath() string {
    return "topology-nodes"
}

func (topologyNodes *PceTopology_TopologyNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology-node" {
        for _, c := range topologyNodes.TopologyNode {
            if topologyNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode{}
        topologyNodes.TopologyNode = append(topologyNodes.TopologyNode, child)
        return &topologyNodes.TopologyNode[len(topologyNodes.TopologyNode)-1]
    }
    return nil
}

func (topologyNodes *PceTopology_TopologyNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range topologyNodes.TopologyNode {
        children[topologyNodes.TopologyNode[i].GetSegmentPath()] = &topologyNodes.TopologyNode[i]
    }
    return children
}

func (topologyNodes *PceTopology_TopologyNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (topologyNodes *PceTopology_TopologyNodes) GetBundleName() string { return "cisco_ios_xr" }

func (topologyNodes *PceTopology_TopologyNodes) GetYangName() string { return "topology-nodes" }

func (topologyNodes *PceTopology_TopologyNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyNodes *PceTopology_TopologyNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyNodes *PceTopology_TopologyNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyNodes *PceTopology_TopologyNodes) SetParent(parent types.Entity) { topologyNodes.parent = parent }

func (topologyNodes *PceTopology_TopologyNodes) GetParent() types.Entity { return topologyNodes.parent }

func (topologyNodes *PceTopology_TopologyNodes) GetParentYangName() string { return "pce-topology" }

// PceTopology_TopologyNodes_TopologyNode
// Node information
type PceTopology_TopologyNodes_TopologyNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node Identifier. The type is interface{} with
    // range: -2147483648..2147483647.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node Overload Bit. The type is bool.
    Overload interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier

    // Prefix SIDs. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_PrefixSid.
    PrefixSid []PceTopology_TopologyNodes_TopologyNode_PrefixSid

    // IPv4 Link information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link.
    Ipv4Link []PceTopology_TopologyNodes_TopologyNode_Ipv4Link

    // IPv6 Link information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link.
    Ipv6Link []PceTopology_TopologyNodes_TopologyNode_Ipv6Link
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetFilter() yfilter.YFilter { return topologyNode.YFilter }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) SetFilter(yf yfilter.YFilter) { topologyNode.YFilter = yf }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetGoName(yname string) string {
    if yname == "node-identifier" { return "NodeIdentifier" }
    if yname == "node-identifier-xr" { return "NodeIdentifierXr" }
    if yname == "overload" { return "Overload" }
    if yname == "node-protocol-identifier" { return "NodeProtocolIdentifier" }
    if yname == "prefix-sid" { return "PrefixSid" }
    if yname == "ipv4-link" { return "Ipv4Link" }
    if yname == "ipv6-link" { return "Ipv6Link" }
    return ""
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetSegmentPath() string {
    return "topology-node" + "[node-identifier='" + fmt.Sprintf("%v", topologyNode.NodeIdentifier) + "']"
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-protocol-identifier" {
        return &topologyNode.NodeProtocolIdentifier
    }
    if childYangName == "prefix-sid" {
        for _, c := range topologyNode.PrefixSid {
            if topologyNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_PrefixSid{}
        topologyNode.PrefixSid = append(topologyNode.PrefixSid, child)
        return &topologyNode.PrefixSid[len(topologyNode.PrefixSid)-1]
    }
    if childYangName == "ipv4-link" {
        for _, c := range topologyNode.Ipv4Link {
            if topologyNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv4Link{}
        topologyNode.Ipv4Link = append(topologyNode.Ipv4Link, child)
        return &topologyNode.Ipv4Link[len(topologyNode.Ipv4Link)-1]
    }
    if childYangName == "ipv6-link" {
        for _, c := range topologyNode.Ipv6Link {
            if topologyNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv6Link{}
        topologyNode.Ipv6Link = append(topologyNode.Ipv6Link, child)
        return &topologyNode.Ipv6Link[len(topologyNode.Ipv6Link)-1]
    }
    return nil
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-protocol-identifier"] = &topologyNode.NodeProtocolIdentifier
    for i := range topologyNode.PrefixSid {
        children[topologyNode.PrefixSid[i].GetSegmentPath()] = &topologyNode.PrefixSid[i]
    }
    for i := range topologyNode.Ipv4Link {
        children[topologyNode.Ipv4Link[i].GetSegmentPath()] = &topologyNode.Ipv4Link[i]
    }
    for i := range topologyNode.Ipv6Link {
        children[topologyNode.Ipv6Link[i].GetSegmentPath()] = &topologyNode.Ipv6Link[i]
    }
    return children
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-identifier"] = topologyNode.NodeIdentifier
    leafs["node-identifier-xr"] = topologyNode.NodeIdentifierXr
    leafs["overload"] = topologyNode.Overload
    return leafs
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetBundleName() string { return "cisco_ios_xr" }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetYangName() string { return "topology-node" }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) SetParent(parent types.Entity) { topologyNode.parent = parent }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetParent() types.Entity { return topologyNode.parent }

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetParentYangName() string { return "topology-nodes" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier
// Node protocol identifier
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetFilter() yfilter.YFilter { return nodeProtocolIdentifier.YFilter }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { nodeProtocolIdentifier.YFilter = yf }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetSegmentPath() string {
    return "node-protocol-identifier"
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range nodeProtocolIdentifier.IgpInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation{}
        nodeProtocolIdentifier.IgpInformation = append(nodeProtocolIdentifier.IgpInformation, child)
        return &nodeProtocolIdentifier.IgpInformation[len(nodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range nodeProtocolIdentifier.SrgbInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation{}
        nodeProtocolIdentifier.SrgbInformation = append(nodeProtocolIdentifier.SrgbInformation, child)
        return &nodeProtocolIdentifier.SrgbInformation[len(nodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeProtocolIdentifier.IgpInformation {
        children[nodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range nodeProtocolIdentifier.SrgbInformation {
        children[nodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = nodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = nodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = nodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = nodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetYangName() string { return "node-protocol-identifier" }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) SetParent(parent types.Entity) { nodeProtocolIdentifier.parent = parent }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetParent() types.Entity { return nodeProtocolIdentifier.parent }

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetParentYangName() string { return "topology-node" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "node-protocol-identifier" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "node-protocol-identifier" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_PrefixSid
// Prefix SIDs
type PceTopology_TopologyNodes_TopologyNode_PrefixSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // R Flag. The type is bool.
    Rflag interface{}

    // N Flag. The type is bool.
    Nflag interface{}

    // P Flag. The type is bool.
    Pflag interface{}

    // E Flag. The type is bool.
    Eflag interface{}

    // V Flag. The type is bool.
    Vflag interface{}

    // L Flag. The type is bool.
    Lflag interface{}

    // Prefix.
    SidPrefix PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
}

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetFilter() yfilter.YFilter { return prefixSid.YFilter }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) SetFilter(yf yfilter.YFilter) { prefixSid.YFilter = yf }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "rflag" { return "Rflag" }
    if yname == "nflag" { return "Nflag" }
    if yname == "pflag" { return "Pflag" }
    if yname == "eflag" { return "Eflag" }
    if yname == "vflag" { return "Vflag" }
    if yname == "lflag" { return "Lflag" }
    if yname == "sid-prefix" { return "SidPrefix" }
    return ""
}

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetSegmentPath() string {
    return "prefix-sid"
}

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sid-prefix" {
        return &prefixSid.SidPrefix
    }
    return nil
}

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sid-prefix"] = &prefixSid.SidPrefix
    return children
}

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = prefixSid.SidType
    leafs["mpls-label"] = prefixSid.MplsLabel
    leafs["domain-identifier"] = prefixSid.DomainIdentifier
    leafs["rflag"] = prefixSid.Rflag
    leafs["nflag"] = prefixSid.Nflag
    leafs["pflag"] = prefixSid.Pflag
    leafs["eflag"] = prefixSid.Eflag
    leafs["vflag"] = prefixSid.Vflag
    leafs["lflag"] = prefixSid.Lflag
    return leafs
}

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetYangName() string { return "prefix-sid" }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) SetParent(parent types.Entity) { prefixSid.parent = parent }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetParent() types.Entity { return prefixSid.parent }

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetParentYangName() string { return "topology-node" }

// PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetFilter() yfilter.YFilter { return sidPrefix.YFilter }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) SetFilter(yf yfilter.YFilter) { sidPrefix.YFilter = yf }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetSegmentPath() string {
    return "sid-prefix"
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = sidPrefix.AfName
    leafs["ipv4"] = sidPrefix.Ipv4
    leafs["ipv6"] = sidPrefix.Ipv6
    return leafs
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetYangName() string { return "sid-prefix" }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) SetParent(parent types.Entity) { sidPrefix.parent = parent }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetParent() types.Entity { return sidPrefix.parent }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetParentYangName() string { return "prefix-sid" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link struct {
    parent types.Entity
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

    // SRLG Values. The type is slice of interface{} with range: 0..4294967295.
    Srlgs []interface{}

    // Local node IGP information.
    LocalIgpInformation PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid.
    AdjacencySid []PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
}

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetFilter() yfilter.YFilter { return ipv4Link.YFilter }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) SetFilter(yf yfilter.YFilter) { ipv4Link.YFilter = yf }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetGoName(yname string) string {
    if yname == "local-ipv4-address" { return "LocalIpv4Address" }
    if yname == "remote-ipv4-address" { return "RemoteIpv4Address" }
    if yname == "igp-metric" { return "IgpMetric" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "maximum-link-bandwidth" { return "MaximumLinkBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "srlgs" { return "Srlgs" }
    if yname == "local-igp-information" { return "LocalIgpInformation" }
    if yname == "remote-node-protocol-identifier" { return "RemoteNodeProtocolIdentifier" }
    if yname == "adjacency-sid" { return "AdjacencySid" }
    return ""
}

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetSegmentPath() string {
    return "ipv4-link"
}

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-igp-information" {
        return &ipv4Link.LocalIgpInformation
    }
    if childYangName == "remote-node-protocol-identifier" {
        return &ipv4Link.RemoteNodeProtocolIdentifier
    }
    if childYangName == "adjacency-sid" {
        for _, c := range ipv4Link.AdjacencySid {
            if ipv4Link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid{}
        ipv4Link.AdjacencySid = append(ipv4Link.AdjacencySid, child)
        return &ipv4Link.AdjacencySid[len(ipv4Link.AdjacencySid)-1]
    }
    return nil
}

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-igp-information"] = &ipv4Link.LocalIgpInformation
    children["remote-node-protocol-identifier"] = &ipv4Link.RemoteNodeProtocolIdentifier
    for i := range ipv4Link.AdjacencySid {
        children[ipv4Link.AdjacencySid[i].GetSegmentPath()] = &ipv4Link.AdjacencySid[i]
    }
    return children
}

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-ipv4-address"] = ipv4Link.LocalIpv4Address
    leafs["remote-ipv4-address"] = ipv4Link.RemoteIpv4Address
    leafs["igp-metric"] = ipv4Link.IgpMetric
    leafs["te-metric"] = ipv4Link.TeMetric
    leafs["maximum-link-bandwidth"] = ipv4Link.MaximumLinkBandwidth
    leafs["max-reservable-bandwidth"] = ipv4Link.MaxReservableBandwidth
    leafs["srlgs"] = ipv4Link.Srlgs
    return leafs
}

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetYangName() string { return "ipv4-link" }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) SetParent(parent types.Entity) { ipv4Link.parent = parent }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetParent() types.Entity { return ipv4Link.parent }

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetParentYangName() string { return "topology-node" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetFilter() yfilter.YFilter { return localIgpInformation.YFilter }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) SetFilter(yf yfilter.YFilter) { localIgpInformation.YFilter = yf }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetSegmentPath() string {
    return "local-igp-information"
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &localIgpInformation.Igp
    }
    return nil
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &localIgpInformation.Igp
    return children
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = localIgpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = localIgpInformation.AutonomousSystemNumber
    return leafs
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetYangName() string { return "local-igp-information" }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) SetParent(parent types.Entity) { localIgpInformation.parent = parent }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetParent() types.Entity { return localIgpInformation.parent }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetParentYangName() string { return "ipv4-link" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetParentYangName() string { return "local-igp-information" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetFilter() yfilter.YFilter { return remoteNodeProtocolIdentifier.YFilter }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { remoteNodeProtocolIdentifier.YFilter = yf }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetSegmentPath() string {
    return "remote-node-protocol-identifier"
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range remoteNodeProtocolIdentifier.IgpInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation{}
        remoteNodeProtocolIdentifier.IgpInformation = append(remoteNodeProtocolIdentifier.IgpInformation, child)
        return &remoteNodeProtocolIdentifier.IgpInformation[len(remoteNodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range remoteNodeProtocolIdentifier.SrgbInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation{}
        remoteNodeProtocolIdentifier.SrgbInformation = append(remoteNodeProtocolIdentifier.SrgbInformation, child)
        return &remoteNodeProtocolIdentifier.SrgbInformation[len(remoteNodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        children[remoteNodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        children[remoteNodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = remoteNodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = remoteNodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetYangName() string { return "remote-node-protocol-identifier" }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) SetParent(parent types.Entity) { remoteNodeProtocolIdentifier.parent = parent }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetParent() types.Entity { return remoteNodeProtocolIdentifier.parent }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetParentYangName() string { return "ipv4-link" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
// Adjacency SIDs
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // R Flag. The type is bool.
    Rflag interface{}

    // N Flag. The type is bool.
    Nflag interface{}

    // P Flag. The type is bool.
    Pflag interface{}

    // E Flag. The type is bool.
    Eflag interface{}

    // V Flag. The type is bool.
    Vflag interface{}

    // L Flag. The type is bool.
    Lflag interface{}

    // Prefix.
    SidPrefix PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetFilter() yfilter.YFilter { return adjacencySid.YFilter }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) SetFilter(yf yfilter.YFilter) { adjacencySid.YFilter = yf }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "rflag" { return "Rflag" }
    if yname == "nflag" { return "Nflag" }
    if yname == "pflag" { return "Pflag" }
    if yname == "eflag" { return "Eflag" }
    if yname == "vflag" { return "Vflag" }
    if yname == "lflag" { return "Lflag" }
    if yname == "sid-prefix" { return "SidPrefix" }
    return ""
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetSegmentPath() string {
    return "adjacency-sid"
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sid-prefix" {
        return &adjacencySid.SidPrefix
    }
    return nil
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sid-prefix"] = &adjacencySid.SidPrefix
    return children
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = adjacencySid.SidType
    leafs["mpls-label"] = adjacencySid.MplsLabel
    leafs["domain-identifier"] = adjacencySid.DomainIdentifier
    leafs["rflag"] = adjacencySid.Rflag
    leafs["nflag"] = adjacencySid.Nflag
    leafs["pflag"] = adjacencySid.Pflag
    leafs["eflag"] = adjacencySid.Eflag
    leafs["vflag"] = adjacencySid.Vflag
    leafs["lflag"] = adjacencySid.Lflag
    return leafs
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetBundleName() string { return "cisco_ios_xr" }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetYangName() string { return "adjacency-sid" }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) SetParent(parent types.Entity) { adjacencySid.parent = parent }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetParent() types.Entity { return adjacencySid.parent }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetParentYangName() string { return "ipv4-link" }

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetFilter() yfilter.YFilter { return sidPrefix.YFilter }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) SetFilter(yf yfilter.YFilter) { sidPrefix.YFilter = yf }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetSegmentPath() string {
    return "sid-prefix"
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = sidPrefix.AfName
    leafs["ipv4"] = sidPrefix.Ipv4
    leafs["ipv6"] = sidPrefix.Ipv6
    return leafs
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetYangName() string { return "sid-prefix" }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) SetParent(parent types.Entity) { sidPrefix.parent = parent }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetParent() types.Entity { return sidPrefix.parent }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetParentYangName() string { return "adjacency-sid" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link struct {
    parent types.Entity
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
    LocalIgpInformation PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid.
    AdjacencySid []PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
}

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetFilter() yfilter.YFilter { return ipv6Link.YFilter }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) SetFilter(yf yfilter.YFilter) { ipv6Link.YFilter = yf }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetGoName(yname string) string {
    if yname == "local-ipv6-address" { return "LocalIpv6Address" }
    if yname == "remote-ipv6-address" { return "RemoteIpv6Address" }
    if yname == "igp-metric" { return "IgpMetric" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "maximum-link-bandwidth" { return "MaximumLinkBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "local-igp-information" { return "LocalIgpInformation" }
    if yname == "remote-node-protocol-identifier" { return "RemoteNodeProtocolIdentifier" }
    if yname == "adjacency-sid" { return "AdjacencySid" }
    return ""
}

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetSegmentPath() string {
    return "ipv6-link"
}

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-igp-information" {
        return &ipv6Link.LocalIgpInformation
    }
    if childYangName == "remote-node-protocol-identifier" {
        return &ipv6Link.RemoteNodeProtocolIdentifier
    }
    if childYangName == "adjacency-sid" {
        for _, c := range ipv6Link.AdjacencySid {
            if ipv6Link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid{}
        ipv6Link.AdjacencySid = append(ipv6Link.AdjacencySid, child)
        return &ipv6Link.AdjacencySid[len(ipv6Link.AdjacencySid)-1]
    }
    return nil
}

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-igp-information"] = &ipv6Link.LocalIgpInformation
    children["remote-node-protocol-identifier"] = &ipv6Link.RemoteNodeProtocolIdentifier
    for i := range ipv6Link.AdjacencySid {
        children[ipv6Link.AdjacencySid[i].GetSegmentPath()] = &ipv6Link.AdjacencySid[i]
    }
    return children
}

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-ipv6-address"] = ipv6Link.LocalIpv6Address
    leafs["remote-ipv6-address"] = ipv6Link.RemoteIpv6Address
    leafs["igp-metric"] = ipv6Link.IgpMetric
    leafs["te-metric"] = ipv6Link.TeMetric
    leafs["maximum-link-bandwidth"] = ipv6Link.MaximumLinkBandwidth
    leafs["max-reservable-bandwidth"] = ipv6Link.MaxReservableBandwidth
    return leafs
}

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetYangName() string { return "ipv6-link" }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) SetParent(parent types.Entity) { ipv6Link.parent = parent }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetParent() types.Entity { return ipv6Link.parent }

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetParentYangName() string { return "topology-node" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetFilter() yfilter.YFilter { return localIgpInformation.YFilter }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) SetFilter(yf yfilter.YFilter) { localIgpInformation.YFilter = yf }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetSegmentPath() string {
    return "local-igp-information"
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &localIgpInformation.Igp
    }
    return nil
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &localIgpInformation.Igp
    return children
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = localIgpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = localIgpInformation.AutonomousSystemNumber
    return leafs
}

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetYangName() string { return "local-igp-information" }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) SetParent(parent types.Entity) { localIgpInformation.parent = parent }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetParent() types.Entity { return localIgpInformation.parent }

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetParentYangName() string { return "ipv6-link" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetParentYangName() string { return "local-igp-information" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetFilter() yfilter.YFilter { return remoteNodeProtocolIdentifier.YFilter }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { remoteNodeProtocolIdentifier.YFilter = yf }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetSegmentPath() string {
    return "remote-node-protocol-identifier"
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range remoteNodeProtocolIdentifier.IgpInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation{}
        remoteNodeProtocolIdentifier.IgpInformation = append(remoteNodeProtocolIdentifier.IgpInformation, child)
        return &remoteNodeProtocolIdentifier.IgpInformation[len(remoteNodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range remoteNodeProtocolIdentifier.SrgbInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation{}
        remoteNodeProtocolIdentifier.SrgbInformation = append(remoteNodeProtocolIdentifier.SrgbInformation, child)
        return &remoteNodeProtocolIdentifier.SrgbInformation[len(remoteNodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        children[remoteNodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        children[remoteNodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = remoteNodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = remoteNodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetYangName() string { return "remote-node-protocol-identifier" }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) SetParent(parent types.Entity) { remoteNodeProtocolIdentifier.parent = parent }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetParent() types.Entity { return remoteNodeProtocolIdentifier.parent }

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetParentYangName() string { return "ipv6-link" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
// Adjacency SIDs
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // R Flag. The type is bool.
    Rflag interface{}

    // N Flag. The type is bool.
    Nflag interface{}

    // P Flag. The type is bool.
    Pflag interface{}

    // E Flag. The type is bool.
    Eflag interface{}

    // V Flag. The type is bool.
    Vflag interface{}

    // L Flag. The type is bool.
    Lflag interface{}

    // Prefix.
    SidPrefix PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetFilter() yfilter.YFilter { return adjacencySid.YFilter }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) SetFilter(yf yfilter.YFilter) { adjacencySid.YFilter = yf }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "rflag" { return "Rflag" }
    if yname == "nflag" { return "Nflag" }
    if yname == "pflag" { return "Pflag" }
    if yname == "eflag" { return "Eflag" }
    if yname == "vflag" { return "Vflag" }
    if yname == "lflag" { return "Lflag" }
    if yname == "sid-prefix" { return "SidPrefix" }
    return ""
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetSegmentPath() string {
    return "adjacency-sid"
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sid-prefix" {
        return &adjacencySid.SidPrefix
    }
    return nil
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sid-prefix"] = &adjacencySid.SidPrefix
    return children
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = adjacencySid.SidType
    leafs["mpls-label"] = adjacencySid.MplsLabel
    leafs["domain-identifier"] = adjacencySid.DomainIdentifier
    leafs["rflag"] = adjacencySid.Rflag
    leafs["nflag"] = adjacencySid.Nflag
    leafs["pflag"] = adjacencySid.Pflag
    leafs["eflag"] = adjacencySid.Eflag
    leafs["vflag"] = adjacencySid.Vflag
    leafs["lflag"] = adjacencySid.Lflag
    return leafs
}

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetBundleName() string { return "cisco_ios_xr" }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetYangName() string { return "adjacency-sid" }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) SetParent(parent types.Entity) { adjacencySid.parent = parent }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetParent() types.Entity { return adjacencySid.parent }

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetParentYangName() string { return "ipv6-link" }

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetFilter() yfilter.YFilter { return sidPrefix.YFilter }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) SetFilter(yf yfilter.YFilter) { sidPrefix.YFilter = yf }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetSegmentPath() string {
    return "sid-prefix"
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = sidPrefix.AfName
    leafs["ipv4"] = sidPrefix.Ipv4
    leafs["ipv6"] = sidPrefix.Ipv6
    return leafs
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetYangName() string { return "sid-prefix" }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) SetParent(parent types.Entity) { sidPrefix.parent = parent }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetParent() types.Entity { return sidPrefix.parent }

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetParentYangName() string { return "adjacency-sid" }

// PceTopology_PrefixInfos
// Prefixes database in XTC
type PceTopology_PrefixInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE prefix information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo.
    PrefixInfo []PceTopology_PrefixInfos_PrefixInfo
}

func (prefixInfos *PceTopology_PrefixInfos) GetFilter() yfilter.YFilter { return prefixInfos.YFilter }

func (prefixInfos *PceTopology_PrefixInfos) SetFilter(yf yfilter.YFilter) { prefixInfos.YFilter = yf }

func (prefixInfos *PceTopology_PrefixInfos) GetGoName(yname string) string {
    if yname == "prefix-info" { return "PrefixInfo" }
    return ""
}

func (prefixInfos *PceTopology_PrefixInfos) GetSegmentPath() string {
    return "prefix-infos"
}

func (prefixInfos *PceTopology_PrefixInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-info" {
        for _, c := range prefixInfos.PrefixInfo {
            if prefixInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_PrefixInfos_PrefixInfo{}
        prefixInfos.PrefixInfo = append(prefixInfos.PrefixInfo, child)
        return &prefixInfos.PrefixInfo[len(prefixInfos.PrefixInfo)-1]
    }
    return nil
}

func (prefixInfos *PceTopology_PrefixInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixInfos.PrefixInfo {
        children[prefixInfos.PrefixInfo[i].GetSegmentPath()] = &prefixInfos.PrefixInfo[i]
    }
    return children
}

func (prefixInfos *PceTopology_PrefixInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixInfos *PceTopology_PrefixInfos) GetBundleName() string { return "cisco_ios_xr" }

func (prefixInfos *PceTopology_PrefixInfos) GetYangName() string { return "prefix-infos" }

func (prefixInfos *PceTopology_PrefixInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixInfos *PceTopology_PrefixInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixInfos *PceTopology_PrefixInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixInfos *PceTopology_PrefixInfos) SetParent(parent types.Entity) { prefixInfos.parent = parent }

func (prefixInfos *PceTopology_PrefixInfos) GetParent() types.Entity { return prefixInfos.parent }

func (prefixInfos *PceTopology_PrefixInfos) GetParentYangName() string { return "pce-topology" }

// PceTopology_PrefixInfos_PrefixInfo
// PCE prefix information
type PceTopology_PrefixInfos_PrefixInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is interface{} with range:
    // -2147483648..2147483647.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier

    // Prefix address. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_Address.
    Address []PceTopology_PrefixInfos_PrefixInfo_Address
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetFilter() yfilter.YFilter { return prefixInfo.YFilter }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) SetFilter(yf yfilter.YFilter) { prefixInfo.YFilter = yf }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetGoName(yname string) string {
    if yname == "node-identifier" { return "NodeIdentifier" }
    if yname == "node-identifier-xr" { return "NodeIdentifierXr" }
    if yname == "node-protocol-identifier" { return "NodeProtocolIdentifier" }
    if yname == "address" { return "Address" }
    return ""
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetSegmentPath() string {
    return "prefix-info" + "[node-identifier='" + fmt.Sprintf("%v", prefixInfo.NodeIdentifier) + "']"
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-protocol-identifier" {
        return &prefixInfo.NodeProtocolIdentifier
    }
    if childYangName == "address" {
        for _, c := range prefixInfo.Address {
            if prefixInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_PrefixInfos_PrefixInfo_Address{}
        prefixInfo.Address = append(prefixInfo.Address, child)
        return &prefixInfo.Address[len(prefixInfo.Address)-1]
    }
    return nil
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-protocol-identifier"] = &prefixInfo.NodeProtocolIdentifier
    for i := range prefixInfo.Address {
        children[prefixInfo.Address[i].GetSegmentPath()] = &prefixInfo.Address[i]
    }
    return children
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-identifier"] = prefixInfo.NodeIdentifier
    leafs["node-identifier-xr"] = prefixInfo.NodeIdentifierXr
    return leafs
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetBundleName() string { return "cisco_ios_xr" }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetYangName() string { return "prefix-info" }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) SetParent(parent types.Entity) { prefixInfo.parent = parent }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetParent() types.Entity { return prefixInfo.parent }

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetParentYangName() string { return "prefix-infos" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier
// Node protocol identifier
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetFilter() yfilter.YFilter { return nodeProtocolIdentifier.YFilter }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { nodeProtocolIdentifier.YFilter = yf }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetSegmentPath() string {
    return "node-protocol-identifier"
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range nodeProtocolIdentifier.IgpInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation{}
        nodeProtocolIdentifier.IgpInformation = append(nodeProtocolIdentifier.IgpInformation, child)
        return &nodeProtocolIdentifier.IgpInformation[len(nodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range nodeProtocolIdentifier.SrgbInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation{}
        nodeProtocolIdentifier.SrgbInformation = append(nodeProtocolIdentifier.SrgbInformation, child)
        return &nodeProtocolIdentifier.SrgbInformation[len(nodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeProtocolIdentifier.IgpInformation {
        children[nodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range nodeProtocolIdentifier.SrgbInformation {
        children[nodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = nodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = nodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = nodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = nodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetYangName() string { return "node-protocol-identifier" }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) SetParent(parent types.Entity) { nodeProtocolIdentifier.parent = parent }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetParent() types.Entity { return nodeProtocolIdentifier.parent }

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetParentYangName() string { return "prefix-info" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "node-protocol-identifier" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "node-protocol-identifier" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// PceTopology_PrefixInfos_PrefixInfo_Address
// Prefix address
type PceTopology_PrefixInfos_PrefixInfo_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetSegmentPath() string {
    return "address"
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4"] = address.Ipv4
    leafs["ipv6"] = address.Ipv6
    return leafs
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetYangName() string { return "address" }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetParent() types.Entity { return address.parent }

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetParentYangName() string { return "prefix-info" }

// Pce
// pce
type Pce struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Associaition database in XTC.
    AssociationInfos Pce_AssociationInfos

    // CSPF path info.
    Cspf Pce_Cspf

    // Node summary database in XTC.
    TopologySummary Pce_TopologySummary

    // Tunnel database in XTC.
    TunnelInfos Pce_TunnelInfos

    // Detailed peers database in XTC.
    PeerDetailInfos Pce_PeerDetailInfos

    // Node database in XTC.
    TopologyNodes Pce_TopologyNodes

    // Prefixes database in XTC.
    PrefixInfos Pce_PrefixInfos

    // LSP summary database in XTC.
    LspSummary Pce_LspSummary

    // Peers database in XTC.
    PeerInfos Pce_PeerInfos

    // Detailed tunnel database in XTC.
    TunnelDetailInfos Pce_TunnelDetailInfos
}

func (pce *Pce) GetFilter() yfilter.YFilter { return pce.YFilter }

func (pce *Pce) SetFilter(yf yfilter.YFilter) { pce.YFilter = yf }

func (pce *Pce) GetGoName(yname string) string {
    if yname == "association-infos" { return "AssociationInfos" }
    if yname == "cspf" { return "Cspf" }
    if yname == "topology-summary" { return "TopologySummary" }
    if yname == "tunnel-infos" { return "TunnelInfos" }
    if yname == "peer-detail-infos" { return "PeerDetailInfos" }
    if yname == "topology-nodes" { return "TopologyNodes" }
    if yname == "prefix-infos" { return "PrefixInfos" }
    if yname == "lsp-summary" { return "LspSummary" }
    if yname == "peer-infos" { return "PeerInfos" }
    if yname == "tunnel-detail-infos" { return "TunnelDetailInfos" }
    return ""
}

func (pce *Pce) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-xtc-oper:pce"
}

func (pce *Pce) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "association-infos" {
        return &pce.AssociationInfos
    }
    if childYangName == "cspf" {
        return &pce.Cspf
    }
    if childYangName == "topology-summary" {
        return &pce.TopologySummary
    }
    if childYangName == "tunnel-infos" {
        return &pce.TunnelInfos
    }
    if childYangName == "peer-detail-infos" {
        return &pce.PeerDetailInfos
    }
    if childYangName == "topology-nodes" {
        return &pce.TopologyNodes
    }
    if childYangName == "prefix-infos" {
        return &pce.PrefixInfos
    }
    if childYangName == "lsp-summary" {
        return &pce.LspSummary
    }
    if childYangName == "peer-infos" {
        return &pce.PeerInfos
    }
    if childYangName == "tunnel-detail-infos" {
        return &pce.TunnelDetailInfos
    }
    return nil
}

func (pce *Pce) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["association-infos"] = &pce.AssociationInfos
    children["cspf"] = &pce.Cspf
    children["topology-summary"] = &pce.TopologySummary
    children["tunnel-infos"] = &pce.TunnelInfos
    children["peer-detail-infos"] = &pce.PeerDetailInfos
    children["topology-nodes"] = &pce.TopologyNodes
    children["prefix-infos"] = &pce.PrefixInfos
    children["lsp-summary"] = &pce.LspSummary
    children["peer-infos"] = &pce.PeerInfos
    children["tunnel-detail-infos"] = &pce.TunnelDetailInfos
    return children
}

func (pce *Pce) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pce *Pce) GetBundleName() string { return "cisco_ios_xr" }

func (pce *Pce) GetYangName() string { return "pce" }

func (pce *Pce) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pce *Pce) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pce *Pce) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pce *Pce) SetParent(parent types.Entity) { pce.parent = parent }

func (pce *Pce) GetParent() types.Entity { return pce.parent }

func (pce *Pce) GetParentYangName() string { return "Cisco-IOS-XR-infra-xtc-oper" }

// Pce_AssociationInfos
// Associaition database in XTC
type Pce_AssociationInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE Association information. The type is slice of
    // Pce_AssociationInfos_AssociationInfo.
    AssociationInfo []Pce_AssociationInfos_AssociationInfo
}

func (associationInfos *Pce_AssociationInfos) GetFilter() yfilter.YFilter { return associationInfos.YFilter }

func (associationInfos *Pce_AssociationInfos) SetFilter(yf yfilter.YFilter) { associationInfos.YFilter = yf }

func (associationInfos *Pce_AssociationInfos) GetGoName(yname string) string {
    if yname == "association-info" { return "AssociationInfo" }
    return ""
}

func (associationInfos *Pce_AssociationInfos) GetSegmentPath() string {
    return "association-infos"
}

func (associationInfos *Pce_AssociationInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "association-info" {
        for _, c := range associationInfos.AssociationInfo {
            if associationInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_AssociationInfos_AssociationInfo{}
        associationInfos.AssociationInfo = append(associationInfos.AssociationInfo, child)
        return &associationInfos.AssociationInfo[len(associationInfos.AssociationInfo)-1]
    }
    return nil
}

func (associationInfos *Pce_AssociationInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range associationInfos.AssociationInfo {
        children[associationInfos.AssociationInfo[i].GetSegmentPath()] = &associationInfos.AssociationInfo[i]
    }
    return children
}

func (associationInfos *Pce_AssociationInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (associationInfos *Pce_AssociationInfos) GetBundleName() string { return "cisco_ios_xr" }

func (associationInfos *Pce_AssociationInfos) GetYangName() string { return "association-infos" }

func (associationInfos *Pce_AssociationInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (associationInfos *Pce_AssociationInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (associationInfos *Pce_AssociationInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (associationInfos *Pce_AssociationInfos) SetParent(parent types.Entity) { associationInfos.parent = parent }

func (associationInfos *Pce_AssociationInfos) GetParent() types.Entity { return associationInfos.parent }

func (associationInfos *Pce_AssociationInfos) GetParentYangName() string { return "pce" }

// Pce_AssociationInfos_AssociationInfo
// PCE Association information
type Pce_AssociationInfos_AssociationInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // -2147483648..2147483647.
    GroupId interface{}

    // Type. The type is PceAsso.
    Type interface{}

    // Sub ID. The type is interface{} with range: -2147483648..2147483647.
    SubId interface{}

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AssociationSource interface{}

    // Association Strict Mode. The type is bool.
    Strict interface{}

    // Association Status. The type is interface{} with range: 0..4294967295.
    Status interface{}

    // Headends Swapped. The type is interface{} with range: 0..4294967295.
    HeadendsSwapped interface{}

    // Association LSP Info. The type is slice of
    // Pce_AssociationInfos_AssociationInfo_AssociationLsp.
    AssociationLsp []Pce_AssociationInfos_AssociationInfo_AssociationLsp
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetFilter() yfilter.YFilter { return associationInfo.YFilter }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) SetFilter(yf yfilter.YFilter) { associationInfo.YFilter = yf }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetGoName(yname string) string {
    if yname == "group-id" { return "GroupId" }
    if yname == "type" { return "Type" }
    if yname == "sub-id" { return "SubId" }
    if yname == "association-type" { return "AssociationType" }
    if yname == "association-id" { return "AssociationId" }
    if yname == "association-source" { return "AssociationSource" }
    if yname == "strict" { return "Strict" }
    if yname == "status" { return "Status" }
    if yname == "headends-swapped" { return "HeadendsSwapped" }
    if yname == "association-lsp" { return "AssociationLsp" }
    return ""
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetSegmentPath() string {
    return "association-info" + "[group-id='" + fmt.Sprintf("%v", associationInfo.GroupId) + "']"
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "association-lsp" {
        for _, c := range associationInfo.AssociationLsp {
            if associationInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_AssociationInfos_AssociationInfo_AssociationLsp{}
        associationInfo.AssociationLsp = append(associationInfo.AssociationLsp, child)
        return &associationInfo.AssociationLsp[len(associationInfo.AssociationLsp)-1]
    }
    return nil
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range associationInfo.AssociationLsp {
        children[associationInfo.AssociationLsp[i].GetSegmentPath()] = &associationInfo.AssociationLsp[i]
    }
    return children
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-id"] = associationInfo.GroupId
    leafs["type"] = associationInfo.Type
    leafs["sub-id"] = associationInfo.SubId
    leafs["association-type"] = associationInfo.AssociationType
    leafs["association-id"] = associationInfo.AssociationId
    leafs["association-source"] = associationInfo.AssociationSource
    leafs["strict"] = associationInfo.Strict
    leafs["status"] = associationInfo.Status
    leafs["headends-swapped"] = associationInfo.HeadendsSwapped
    return leafs
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetBundleName() string { return "cisco_ios_xr" }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetYangName() string { return "association-info" }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) SetParent(parent types.Entity) { associationInfo.parent = parent }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetParent() types.Entity { return associationInfo.parent }

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetParentYangName() string { return "association-infos" }

// Pce_AssociationInfos_AssociationInfo_AssociationLsp
// Association LSP Info
type Pce_AssociationInfos_AssociationInfo_AssociationLsp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PccAddress interface{}

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // LSP ID. The type is interface{} with range: 0..4294967295.
    Lspid interface{}

    // Tunnel Name. The type is string.
    TunnelName interface{}

    // PCE Based. The type is bool.
    PceBased interface{}

    // PLSP ID. The type is interface{} with range: 0..4294967295.
    PlspId interface{}
}

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetFilter() yfilter.YFilter { return associationLsp.YFilter }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) SetFilter(yf yfilter.YFilter) { associationLsp.YFilter = yf }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetGoName(yname string) string {
    if yname == "pcc-address" { return "PccAddress" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "lspid" { return "Lspid" }
    if yname == "tunnel-name" { return "TunnelName" }
    if yname == "pce-based" { return "PceBased" }
    if yname == "plsp-id" { return "PlspId" }
    return ""
}

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetSegmentPath() string {
    return "association-lsp"
}

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcc-address"] = associationLsp.PccAddress
    leafs["tunnel-id"] = associationLsp.TunnelId
    leafs["lspid"] = associationLsp.Lspid
    leafs["tunnel-name"] = associationLsp.TunnelName
    leafs["pce-based"] = associationLsp.PceBased
    leafs["plsp-id"] = associationLsp.PlspId
    return leafs
}

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetBundleName() string { return "cisco_ios_xr" }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetYangName() string { return "association-lsp" }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) SetParent(parent types.Entity) { associationLsp.parent = parent }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetParent() types.Entity { return associationLsp.parent }

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetParentYangName() string { return "association-info" }

// Pce_Cspf
// CSPF path info
type Pce_Cspf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This table models the path calculation capabilities in XTC.A GET operation
    // for the complete table will return no entries.
    CspfPaths Pce_Cspf_CspfPaths
}

func (cspf *Pce_Cspf) GetFilter() yfilter.YFilter { return cspf.YFilter }

func (cspf *Pce_Cspf) SetFilter(yf yfilter.YFilter) { cspf.YFilter = yf }

func (cspf *Pce_Cspf) GetGoName(yname string) string {
    if yname == "cspf-paths" { return "CspfPaths" }
    return ""
}

func (cspf *Pce_Cspf) GetSegmentPath() string {
    return "cspf"
}

func (cspf *Pce_Cspf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cspf-paths" {
        return &cspf.CspfPaths
    }
    return nil
}

func (cspf *Pce_Cspf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cspf-paths"] = &cspf.CspfPaths
    return children
}

func (cspf *Pce_Cspf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cspf *Pce_Cspf) GetBundleName() string { return "cisco_ios_xr" }

func (cspf *Pce_Cspf) GetYangName() string { return "cspf" }

func (cspf *Pce_Cspf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cspf *Pce_Cspf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cspf *Pce_Cspf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cspf *Pce_Cspf) SetParent(parent types.Entity) { cspf.parent = parent }

func (cspf *Pce_Cspf) GetParent() types.Entity { return cspf.parent }

func (cspf *Pce_Cspf) GetParentYangName() string { return "pce" }

// Pce_Cspf_CspfPaths
// This table models the path calculation
// capabilities in XTC.A GET operation for the
// complete table will return no entries.
type Pce_Cspf_CspfPaths struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A GET operation on this class returns the path . The type is slice of
    // Pce_Cspf_CspfPaths_CspfPath.
    CspfPath []Pce_Cspf_CspfPaths_CspfPath
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetFilter() yfilter.YFilter { return cspfPaths.YFilter }

func (cspfPaths *Pce_Cspf_CspfPaths) SetFilter(yf yfilter.YFilter) { cspfPaths.YFilter = yf }

func (cspfPaths *Pce_Cspf_CspfPaths) GetGoName(yname string) string {
    if yname == "cspf-path" { return "CspfPath" }
    return ""
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetSegmentPath() string {
    return "cspf-paths"
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cspf-path" {
        for _, c := range cspfPaths.CspfPath {
            if cspfPaths.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_Cspf_CspfPaths_CspfPath{}
        cspfPaths.CspfPath = append(cspfPaths.CspfPath, child)
        return &cspfPaths.CspfPath[len(cspfPaths.CspfPath)-1]
    }
    return nil
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cspfPaths.CspfPath {
        children[cspfPaths.CspfPath[i].GetSegmentPath()] = &cspfPaths.CspfPath[i]
    }
    return children
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetBundleName() string { return "cisco_ios_xr" }

func (cspfPaths *Pce_Cspf_CspfPaths) GetYangName() string { return "cspf-paths" }

func (cspfPaths *Pce_Cspf_CspfPaths) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cspfPaths *Pce_Cspf_CspfPaths) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cspfPaths *Pce_Cspf_CspfPaths) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cspfPaths *Pce_Cspf_CspfPaths) SetParent(parent types.Entity) { cspfPaths.parent = parent }

func (cspfPaths *Pce_Cspf_CspfPaths) GetParent() types.Entity { return cspfPaths.parent }

func (cspfPaths *Pce_Cspf_CspfPaths) GetParentYangName() string { return "cspf" }

// Pce_Cspf_CspfPaths_CspfPath
// A GET operation on this class returns the path
// .
type Pce_Cspf_CspfPaths_CspfPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is interface{} with
    // range: -2147483648..2147483647.
    Af interface{}

    // This attribute is a key. Source of path 1. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Source1 interface{}

    // This attribute is a key. Destination of path 1. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Destination1 interface{}

    // This attribute is a key. Metric type. The type is interface{} with range:
    // -2147483648..2147483647.
    MetricType interface{}

    // This attribute is a key. Source of path 2. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Source2 interface{}

    // This attribute is a key. Destination of path 2. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Destination2 interface{}

    // This attribute is a key. Disjointness level. The type is interface{} with
    // range: -2147483648..2147483647.
    DisjointLevel interface{}

    // This attribute is a key. Strict disjointness required. The type is
    // interface{} with range: -2147483648..2147483647.
    DisjointStrict interface{}

    // This attribute is a key. Whether path 1 or 2 should be shortest. The type
    // is interface{} with range: -2147483648..2147483647.
    ShortestPath interface{}

    // Headends swapped. The type is PceHeadendSwap.
    HeadendsSwapped interface{}

    // CSPF Result. The type is PceCspfRc.
    CspfResult interface{}

    // Output PCE paths. The type is slice of
    // Pce_Cspf_CspfPaths_CspfPath_OutputPath.
    OutputPath []Pce_Cspf_CspfPaths_CspfPath_OutputPath
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetFilter() yfilter.YFilter { return cspfPath.YFilter }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) SetFilter(yf yfilter.YFilter) { cspfPath.YFilter = yf }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetGoName(yname string) string {
    if yname == "af" { return "Af" }
    if yname == "source1" { return "Source1" }
    if yname == "destination1" { return "Destination1" }
    if yname == "metric-type" { return "MetricType" }
    if yname == "source2" { return "Source2" }
    if yname == "destination2" { return "Destination2" }
    if yname == "disjoint-level" { return "DisjointLevel" }
    if yname == "disjoint-strict" { return "DisjointStrict" }
    if yname == "shortest-path" { return "ShortestPath" }
    if yname == "headends-swapped" { return "HeadendsSwapped" }
    if yname == "cspf-result" { return "CspfResult" }
    if yname == "output-path" { return "OutputPath" }
    return ""
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetSegmentPath() string {
    return "cspf-path" + "[af='" + fmt.Sprintf("%v", cspfPath.Af) + "']" + "[source1='" + fmt.Sprintf("%v", cspfPath.Source1) + "']" + "[destination1='" + fmt.Sprintf("%v", cspfPath.Destination1) + "']" + "[metric-type='" + fmt.Sprintf("%v", cspfPath.MetricType) + "']" + "[source2='" + fmt.Sprintf("%v", cspfPath.Source2) + "']" + "[destination2='" + fmt.Sprintf("%v", cspfPath.Destination2) + "']" + "[disjoint-level='" + fmt.Sprintf("%v", cspfPath.DisjointLevel) + "']" + "[disjoint-strict='" + fmt.Sprintf("%v", cspfPath.DisjointStrict) + "']" + "[shortest-path='" + fmt.Sprintf("%v", cspfPath.ShortestPath) + "']"
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "output-path" {
        for _, c := range cspfPath.OutputPath {
            if cspfPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_Cspf_CspfPaths_CspfPath_OutputPath{}
        cspfPath.OutputPath = append(cspfPath.OutputPath, child)
        return &cspfPath.OutputPath[len(cspfPath.OutputPath)-1]
    }
    return nil
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cspfPath.OutputPath {
        children[cspfPath.OutputPath[i].GetSegmentPath()] = &cspfPath.OutputPath[i]
    }
    return children
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af"] = cspfPath.Af
    leafs["source1"] = cspfPath.Source1
    leafs["destination1"] = cspfPath.Destination1
    leafs["metric-type"] = cspfPath.MetricType
    leafs["source2"] = cspfPath.Source2
    leafs["destination2"] = cspfPath.Destination2
    leafs["disjoint-level"] = cspfPath.DisjointLevel
    leafs["disjoint-strict"] = cspfPath.DisjointStrict
    leafs["shortest-path"] = cspfPath.ShortestPath
    leafs["headends-swapped"] = cspfPath.HeadendsSwapped
    leafs["cspf-result"] = cspfPath.CspfResult
    return leafs
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetBundleName() string { return "cisco_ios_xr" }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetYangName() string { return "cspf-path" }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) SetParent(parent types.Entity) { cspfPath.parent = parent }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetParent() types.Entity { return cspfPath.parent }

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetParentYangName() string { return "cspf-paths" }

// Pce_Cspf_CspfPaths_CspfPath_OutputPath
// Output PCE paths
type Pce_Cspf_CspfPaths_CspfPath_OutputPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Cost. The type is interface{} with range: 0..18446744073709551615.
    Cost interface{}

    // Source of path.
    Source Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source

    // Destination of path.
    Destination Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination

    // Hop addresses. The type is slice of
    // Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops.
    Hops []Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops
}

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetFilter() yfilter.YFilter { return outputPath.YFilter }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) SetFilter(yf yfilter.YFilter) { outputPath.YFilter = yf }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetGoName(yname string) string {
    if yname == "cost" { return "Cost" }
    if yname == "source" { return "Source" }
    if yname == "destination" { return "Destination" }
    if yname == "hops" { return "Hops" }
    return ""
}

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetSegmentPath() string {
    return "output-path"
}

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "source" {
        return &outputPath.Source
    }
    if childYangName == "destination" {
        return &outputPath.Destination
    }
    if childYangName == "hops" {
        for _, c := range outputPath.Hops {
            if outputPath.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops{}
        outputPath.Hops = append(outputPath.Hops, child)
        return &outputPath.Hops[len(outputPath.Hops)-1]
    }
    return nil
}

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["source"] = &outputPath.Source
    children["destination"] = &outputPath.Destination
    for i := range outputPath.Hops {
        children[outputPath.Hops[i].GetSegmentPath()] = &outputPath.Hops[i]
    }
    return children
}

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cost"] = outputPath.Cost
    return leafs
}

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetBundleName() string { return "cisco_ios_xr" }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetYangName() string { return "output-path" }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) SetParent(parent types.Entity) { outputPath.parent = parent }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetParent() types.Entity { return outputPath.parent }

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetParentYangName() string { return "cspf-path" }

// Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source
// Source of path
type Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetFilter() yfilter.YFilter { return source.YFilter }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) SetFilter(yf yfilter.YFilter) { source.YFilter = yf }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetSegmentPath() string {
    return "source"
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = source.AfName
    leafs["ipv4"] = source.Ipv4
    leafs["ipv6"] = source.Ipv6
    return leafs
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetBundleName() string { return "cisco_ios_xr" }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetYangName() string { return "source" }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) SetParent(parent types.Entity) { source.parent = parent }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetParent() types.Entity { return source.parent }

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetParentYangName() string { return "output-path" }

// Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination
// Destination of path
type Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetFilter() yfilter.YFilter { return destination.YFilter }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) SetFilter(yf yfilter.YFilter) { destination.YFilter = yf }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetSegmentPath() string {
    return "destination"
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = destination.AfName
    leafs["ipv4"] = destination.Ipv4
    leafs["ipv6"] = destination.Ipv6
    return leafs
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetBundleName() string { return "cisco_ios_xr" }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetYangName() string { return "destination" }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) SetParent(parent types.Entity) { destination.parent = parent }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetParent() types.Entity { return destination.parent }

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetParentYangName() string { return "output-path" }

// Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops
// Hop addresses
type Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family. The type is interface{} with range: 0..255.
    AddressFamily interface{}

    // IPv4 prefix. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Prefix interface{}
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetFilter() yfilter.YFilter { return hops.YFilter }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) SetFilter(yf yfilter.YFilter) { hops.YFilter = yf }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetGoName(yname string) string {
    if yname == "address-family" { return "AddressFamily" }
    if yname == "ipv4-prefix" { return "Ipv4Prefix" }
    if yname == "ipv6-prefix" { return "Ipv6Prefix" }
    return ""
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetSegmentPath() string {
    return "hops"
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-family"] = hops.AddressFamily
    leafs["ipv4-prefix"] = hops.Ipv4Prefix
    leafs["ipv6-prefix"] = hops.Ipv6Prefix
    return leafs
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetBundleName() string { return "cisco_ios_xr" }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetYangName() string { return "hops" }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) SetParent(parent types.Entity) { hops.parent = parent }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetParent() types.Entity { return hops.parent }

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetParentYangName() string { return "output-path" }

// Pce_TopologySummary
// Node summary database in XTC
type Pce_TopologySummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of PCE nodes. The type is interface{} with range: 0..4294967295.
    Nodes interface{}

    // Number of lookup nodes. The type is interface{} with range: 0..4294967295.
    LookupNodes interface{}

    // Number of prefixes. The type is interface{} with range: 0..4294967295.
    Prefixes interface{}

    // Number of total prefix SIDs. The type is interface{} with range:
    // 0..4294967295.
    PrefixSids interface{}

    // Number of reguar prefix SIDs. The type is interface{} with range:
    // 0..4294967295.
    RegularPrefixSids interface{}

    // Number of strict prefix SIDs. The type is interface{} with range:
    // 0..4294967295.
    StrictPrefixSids interface{}

    // Number of links. The type is interface{} with range: 0..4294967295.
    Links interface{}

    // Number of EPE links. The type is interface{} with range: 0..4294967295.
    EpeLinks interface{}

    // Number of total adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    AdjacencySids interface{}

    // Number of total EPE SIDs. The type is interface{} with range:
    // 0..4294967295.
    Epesids interface{}

    // Number of protected adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    ProtectedAdjacencySids interface{}

    // Number of unprotected adjacency SIDs. The type is interface{} with range:
    // 0..4294967295.
    UnProtectedAdjacencySids interface{}

    // True if topology is consistent. The type is bool.
    TopologyConsistent interface{}

    // Statistics on topology update.
    StatsTopologyUpdate Pce_TopologySummary_StatsTopologyUpdate
}

func (topologySummary *Pce_TopologySummary) GetFilter() yfilter.YFilter { return topologySummary.YFilter }

func (topologySummary *Pce_TopologySummary) SetFilter(yf yfilter.YFilter) { topologySummary.YFilter = yf }

func (topologySummary *Pce_TopologySummary) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    if yname == "lookup-nodes" { return "LookupNodes" }
    if yname == "prefixes" { return "Prefixes" }
    if yname == "prefix-sids" { return "PrefixSids" }
    if yname == "regular-prefix-sids" { return "RegularPrefixSids" }
    if yname == "strict-prefix-sids" { return "StrictPrefixSids" }
    if yname == "links" { return "Links" }
    if yname == "epe-links" { return "EpeLinks" }
    if yname == "adjacency-sids" { return "AdjacencySids" }
    if yname == "epesids" { return "Epesids" }
    if yname == "protected-adjacency-sids" { return "ProtectedAdjacencySids" }
    if yname == "un-protected-adjacency-sids" { return "UnProtectedAdjacencySids" }
    if yname == "topology-consistent" { return "TopologyConsistent" }
    if yname == "stats-topology-update" { return "StatsTopologyUpdate" }
    return ""
}

func (topologySummary *Pce_TopologySummary) GetSegmentPath() string {
    return "topology-summary"
}

func (topologySummary *Pce_TopologySummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "stats-topology-update" {
        return &topologySummary.StatsTopologyUpdate
    }
    return nil
}

func (topologySummary *Pce_TopologySummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["stats-topology-update"] = &topologySummary.StatsTopologyUpdate
    return children
}

func (topologySummary *Pce_TopologySummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nodes"] = topologySummary.Nodes
    leafs["lookup-nodes"] = topologySummary.LookupNodes
    leafs["prefixes"] = topologySummary.Prefixes
    leafs["prefix-sids"] = topologySummary.PrefixSids
    leafs["regular-prefix-sids"] = topologySummary.RegularPrefixSids
    leafs["strict-prefix-sids"] = topologySummary.StrictPrefixSids
    leafs["links"] = topologySummary.Links
    leafs["epe-links"] = topologySummary.EpeLinks
    leafs["adjacency-sids"] = topologySummary.AdjacencySids
    leafs["epesids"] = topologySummary.Epesids
    leafs["protected-adjacency-sids"] = topologySummary.ProtectedAdjacencySids
    leafs["un-protected-adjacency-sids"] = topologySummary.UnProtectedAdjacencySids
    leafs["topology-consistent"] = topologySummary.TopologyConsistent
    return leafs
}

func (topologySummary *Pce_TopologySummary) GetBundleName() string { return "cisco_ios_xr" }

func (topologySummary *Pce_TopologySummary) GetYangName() string { return "topology-summary" }

func (topologySummary *Pce_TopologySummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologySummary *Pce_TopologySummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologySummary *Pce_TopologySummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologySummary *Pce_TopologySummary) SetParent(parent types.Entity) { topologySummary.parent = parent }

func (topologySummary *Pce_TopologySummary) GetParent() types.Entity { return topologySummary.parent }

func (topologySummary *Pce_TopologySummary) GetParentYangName() string { return "pce" }

// Pce_TopologySummary_StatsTopologyUpdate
// Statistics on topology update
type Pce_TopologySummary_StatsTopologyUpdate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of nodes added. The type is interface{} with range: 0..4294967295.
    NumNodesAdded interface{}

    // Number of nodes deleted. The type is interface{} with range: 0..4294967295.
    NumNodesDeleted interface{}

    // Number of links added. The type is interface{} with range: 0..4294967295.
    NumLinksAdded interface{}

    // Number of links deleted. The type is interface{} with range: 0..4294967295.
    NumLinksDeleted interface{}

    // Number of prefixes added. The type is interface{} with range:
    // 0..4294967295.
    NumPrefixesAdded interface{}

    // Number of prefixes deleted. The type is interface{} with range:
    // 0..4294967295.
    NumPrefixesDeleted interface{}
}

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetFilter() yfilter.YFilter { return statsTopologyUpdate.YFilter }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) SetFilter(yf yfilter.YFilter) { statsTopologyUpdate.YFilter = yf }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetGoName(yname string) string {
    if yname == "num-nodes-added" { return "NumNodesAdded" }
    if yname == "num-nodes-deleted" { return "NumNodesDeleted" }
    if yname == "num-links-added" { return "NumLinksAdded" }
    if yname == "num-links-deleted" { return "NumLinksDeleted" }
    if yname == "num-prefixes-added" { return "NumPrefixesAdded" }
    if yname == "num-prefixes-deleted" { return "NumPrefixesDeleted" }
    return ""
}

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetSegmentPath() string {
    return "stats-topology-update"
}

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-nodes-added"] = statsTopologyUpdate.NumNodesAdded
    leafs["num-nodes-deleted"] = statsTopologyUpdate.NumNodesDeleted
    leafs["num-links-added"] = statsTopologyUpdate.NumLinksAdded
    leafs["num-links-deleted"] = statsTopologyUpdate.NumLinksDeleted
    leafs["num-prefixes-added"] = statsTopologyUpdate.NumPrefixesAdded
    leafs["num-prefixes-deleted"] = statsTopologyUpdate.NumPrefixesDeleted
    return leafs
}

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetBundleName() string { return "cisco_ios_xr" }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetYangName() string { return "stats-topology-update" }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) SetParent(parent types.Entity) { statsTopologyUpdate.parent = parent }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetParent() types.Entity { return statsTopologyUpdate.parent }

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetParentYangName() string { return "topology-summary" }

// Pce_TunnelInfos
// Tunnel database in XTC
type Pce_TunnelInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of Pce_TunnelInfos_TunnelInfo.
    TunnelInfo []Pce_TunnelInfos_TunnelInfo
}

func (tunnelInfos *Pce_TunnelInfos) GetFilter() yfilter.YFilter { return tunnelInfos.YFilter }

func (tunnelInfos *Pce_TunnelInfos) SetFilter(yf yfilter.YFilter) { tunnelInfos.YFilter = yf }

func (tunnelInfos *Pce_TunnelInfos) GetGoName(yname string) string {
    if yname == "tunnel-info" { return "TunnelInfo" }
    return ""
}

func (tunnelInfos *Pce_TunnelInfos) GetSegmentPath() string {
    return "tunnel-infos"
}

func (tunnelInfos *Pce_TunnelInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-info" {
        for _, c := range tunnelInfos.TunnelInfo {
            if tunnelInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelInfos_TunnelInfo{}
        tunnelInfos.TunnelInfo = append(tunnelInfos.TunnelInfo, child)
        return &tunnelInfos.TunnelInfo[len(tunnelInfos.TunnelInfo)-1]
    }
    return nil
}

func (tunnelInfos *Pce_TunnelInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelInfos.TunnelInfo {
        children[tunnelInfos.TunnelInfo[i].GetSegmentPath()] = &tunnelInfos.TunnelInfo[i]
    }
    return children
}

func (tunnelInfos *Pce_TunnelInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelInfos *Pce_TunnelInfos) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelInfos *Pce_TunnelInfos) GetYangName() string { return "tunnel-infos" }

func (tunnelInfos *Pce_TunnelInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelInfos *Pce_TunnelInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelInfos *Pce_TunnelInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelInfos *Pce_TunnelInfos) SetParent(parent types.Entity) { tunnelInfos.parent = parent }

func (tunnelInfos *Pce_TunnelInfos) GetParent() types.Entity { return tunnelInfos.parent }

func (tunnelInfos *Pce_TunnelInfos) GetParentYangName() string { return "pce" }

// Pce_TunnelInfos_TunnelInfo
// Tunnel information
type Pce_TunnelInfos_TunnelInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Brief LSP information. The type is slice of
    // Pce_TunnelInfos_TunnelInfo_BriefLspInformation.
    BriefLspInformation []Pce_TunnelInfos_TunnelInfo_BriefLspInformation
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetFilter() yfilter.YFilter { return tunnelInfo.YFilter }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) SetFilter(yf yfilter.YFilter) { tunnelInfo.YFilter = yf }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "plsp-id" { return "PlspId" }
    if yname == "tunnel-name" { return "TunnelName" }
    if yname == "pcc-address" { return "PccAddress" }
    if yname == "tunnel-name-xr" { return "TunnelNameXr" }
    if yname == "brief-lsp-information" { return "BriefLspInformation" }
    return ""
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetSegmentPath() string {
    return "tunnel-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelInfo.TunnelName) + "']"
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-lsp-information" {
        for _, c := range tunnelInfo.BriefLspInformation {
            if tunnelInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelInfos_TunnelInfo_BriefLspInformation{}
        tunnelInfo.BriefLspInformation = append(tunnelInfo.BriefLspInformation, child)
        return &tunnelInfo.BriefLspInformation[len(tunnelInfo.BriefLspInformation)-1]
    }
    return nil
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelInfo.BriefLspInformation {
        children[tunnelInfo.BriefLspInformation[i].GetSegmentPath()] = &tunnelInfo.BriefLspInformation[i]
    }
    return children
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = tunnelInfo.PeerAddress
    leafs["plsp-id"] = tunnelInfo.PlspId
    leafs["tunnel-name"] = tunnelInfo.TunnelName
    leafs["pcc-address"] = tunnelInfo.PccAddress
    leafs["tunnel-name-xr"] = tunnelInfo.TunnelNameXr
    return leafs
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetYangName() string { return "tunnel-info" }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) SetParent(parent types.Entity) { tunnelInfo.parent = parent }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetParent() types.Entity { return tunnelInfo.parent }

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetParentYangName() string { return "tunnel-infos" }

// Pce_TunnelInfos_TunnelInfo_BriefLspInformation
// Brief LSP information
type Pce_TunnelInfos_TunnelInfo_BriefLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // LSP ID. The type is interface{} with range: 0..4294967295.
    Lspid interface{}

    // Binding SID. The type is interface{} with range: 0..4294967295.
    BindingSid interface{}

    // LSP Setup Type. The type is LspSetup.
    LspSetupType interface{}

    // Operational state. The type is PcepLspState.
    OperationalState interface{}

    // Admin state. The type is LspState.
    AdministrativeState interface{}
}

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetFilter() yfilter.YFilter { return briefLspInformation.YFilter }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) SetFilter(yf yfilter.YFilter) { briefLspInformation.YFilter = yf }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "lspid" { return "Lspid" }
    if yname == "binding-sid" { return "BindingSid" }
    if yname == "lsp-setup-type" { return "LspSetupType" }
    if yname == "operational-state" { return "OperationalState" }
    if yname == "administrative-state" { return "AdministrativeState" }
    return ""
}

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetSegmentPath() string {
    return "brief-lsp-information"
}

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = briefLspInformation.SourceAddress
    leafs["destination-address"] = briefLspInformation.DestinationAddress
    leafs["tunnel-id"] = briefLspInformation.TunnelId
    leafs["lspid"] = briefLspInformation.Lspid
    leafs["binding-sid"] = briefLspInformation.BindingSid
    leafs["lsp-setup-type"] = briefLspInformation.LspSetupType
    leafs["operational-state"] = briefLspInformation.OperationalState
    leafs["administrative-state"] = briefLspInformation.AdministrativeState
    return leafs
}

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetYangName() string { return "brief-lsp-information" }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) SetParent(parent types.Entity) { briefLspInformation.parent = parent }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetParent() types.Entity { return briefLspInformation.parent }

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetParentYangName() string { return "tunnel-info" }

// Pce_PeerDetailInfos
// Detailed peers database in XTC
type Pce_PeerDetailInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed PCE peer information. The type is slice of
    // Pce_PeerDetailInfos_PeerDetailInfo.
    PeerDetailInfo []Pce_PeerDetailInfos_PeerDetailInfo
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetFilter() yfilter.YFilter { return peerDetailInfos.YFilter }

func (peerDetailInfos *Pce_PeerDetailInfos) SetFilter(yf yfilter.YFilter) { peerDetailInfos.YFilter = yf }

func (peerDetailInfos *Pce_PeerDetailInfos) GetGoName(yname string) string {
    if yname == "peer-detail-info" { return "PeerDetailInfo" }
    return ""
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetSegmentPath() string {
    return "peer-detail-infos"
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-detail-info" {
        for _, c := range peerDetailInfos.PeerDetailInfo {
            if peerDetailInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PeerDetailInfos_PeerDetailInfo{}
        peerDetailInfos.PeerDetailInfo = append(peerDetailInfos.PeerDetailInfo, child)
        return &peerDetailInfos.PeerDetailInfo[len(peerDetailInfos.PeerDetailInfo)-1]
    }
    return nil
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerDetailInfos.PeerDetailInfo {
        children[peerDetailInfos.PeerDetailInfo[i].GetSegmentPath()] = &peerDetailInfos.PeerDetailInfo[i]
    }
    return children
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetBundleName() string { return "cisco_ios_xr" }

func (peerDetailInfos *Pce_PeerDetailInfos) GetYangName() string { return "peer-detail-infos" }

func (peerDetailInfos *Pce_PeerDetailInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerDetailInfos *Pce_PeerDetailInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerDetailInfos *Pce_PeerDetailInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerDetailInfos *Pce_PeerDetailInfos) SetParent(parent types.Entity) { peerDetailInfos.parent = parent }

func (peerDetailInfos *Pce_PeerDetailInfos) GetParent() types.Entity { return peerDetailInfos.parent }

func (peerDetailInfos *Pce_PeerDetailInfos) GetParentYangName() string { return "pce" }

// Pce_PeerDetailInfos_PeerDetailInfo
// Detailed PCE peer information
type Pce_PeerDetailInfos_PeerDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Detailed PCE protocol information.
    DetailPcepInformation Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetFilter() yfilter.YFilter { return peerDetailInfo.YFilter }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) SetFilter(yf yfilter.YFilter) { peerDetailInfo.YFilter = yf }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-address-xr" { return "PeerAddressXr" }
    if yname == "peer-protocol" { return "PeerProtocol" }
    if yname == "detail-pcep-information" { return "DetailPcepInformation" }
    return ""
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetSegmentPath() string {
    return "peer-detail-info" + "[peer-address='" + fmt.Sprintf("%v", peerDetailInfo.PeerAddress) + "']"
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-pcep-information" {
        return &peerDetailInfo.DetailPcepInformation
    }
    return nil
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail-pcep-information"] = &peerDetailInfo.DetailPcepInformation
    return children
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peerDetailInfo.PeerAddress
    leafs["peer-address-xr"] = peerDetailInfo.PeerAddressXr
    leafs["peer-protocol"] = peerDetailInfo.PeerProtocol
    return leafs
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetYangName() string { return "peer-detail-info" }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) SetParent(parent types.Entity) { peerDetailInfo.parent = parent }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetParent() types.Entity { return peerDetailInfo.parent }

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetParentYangName() string { return "peer-detail-infos" }

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
// Detailed PCE protocol information
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Error (for display only). The type is string.
    Error interface{}

    // Speaker Entity ID. The type is string.
    SpeakerId interface{}

    // PCEP Up Time. The type is interface{} with range: 0..4294967295.
    PcepUpTime interface{}

    // Keepalive count. The type is interface{} with range: 0..4294967295.
    Keepalives interface{}

    // MD5 Authentication Enabled. The type is bool.
    Md5Enabled interface{}

    // Keychain based Authentication Enabled. The type is bool.
    KeychainEnabled interface{}

    // Negotiated KA. The type is interface{} with range: 0..4294967295.
    NegotiatedLocalKeepalive interface{}

    // Negotiated KA. The type is interface{} with range: 0..4294967295.
    NegotiatedRemoteKeepalive interface{}

    // Negotiated DT. The type is interface{} with range: 0..4294967295.
    NegotiatedDeadTime interface{}

    // PCEReq Rx. The type is interface{} with range: 0..4294967295.
    PceRequestRx interface{}

    // PCEReq Tx. The type is interface{} with range: 0..4294967295.
    PceRequestTx interface{}

    // PCERep Rx. The type is interface{} with range: 0..4294967295.
    PceReplyRx interface{}

    // PCERep Tx. The type is interface{} with range: 0..4294967295.
    PceReplyTx interface{}

    // PCEErr Rx. The type is interface{} with range: 0..4294967295.
    PceErrorRx interface{}

    // PCEErr Tx. The type is interface{} with range: 0..4294967295.
    PceErrorTx interface{}

    // PCEOpen Tx. The type is interface{} with range: 0..4294967295.
    PceOpenTx interface{}

    // PCEOpen Rx. The type is interface{} with range: 0..4294967295.
    PceOpenRx interface{}

    // PCERpt Rx. The type is interface{} with range: 0..4294967295.
    PceReportRx interface{}

    // PCERpt Tx. The type is interface{} with range: 0..4294967295.
    PceReportTx interface{}

    // PCEUpd Rx. The type is interface{} with range: 0..4294967295.
    PceUpdateRx interface{}

    // PCEUpd Tx. The type is interface{} with range: 0..4294967295.
    PceUpdateTx interface{}

    // PCEInit Rx. The type is interface{} with range: 0..4294967295.
    PceInitiateRx interface{}

    // PCEInit Tx. The type is interface{} with range: 0..4294967295.
    PceInitiateTx interface{}

    // PCE Keepalive Tx. The type is interface{} with range:
    // 0..18446744073709551615.
    PceKeepaliveTx interface{}

    // PCE Keepalive Rx. The type is interface{} with range:
    // 0..18446744073709551615.
    PceKeepaliveRx interface{}

    // Local PCEP session ID. The type is interface{} with range: 0..255.
    LocalSessionId interface{}

    // Remote PCEP session ID. The type is interface{} with range: 0..255.
    RemoteSessionId interface{}

    // Minimum keepalive interval for the peer. The type is interface{} with
    // range: 0..255.
    MinimumKeepaliveInterval interface{}

    // Maximum dead interval for the peer. The type is interface{} with range:
    // 0..255.
    MaximumDeadInterval interface{}

    // Brief PCE protocol information.
    BriefPcepInformation Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation

    // Last PCError received.
    LastErrorRx Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx

    // Last PCError sent.
    LastErrorTx Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx
}

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetFilter() yfilter.YFilter { return detailPcepInformation.YFilter }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) SetFilter(yf yfilter.YFilter) { detailPcepInformation.YFilter = yf }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    if yname == "speaker-id" { return "SpeakerId" }
    if yname == "pcep-up-time" { return "PcepUpTime" }
    if yname == "keepalives" { return "Keepalives" }
    if yname == "md5-enabled" { return "Md5Enabled" }
    if yname == "keychain-enabled" { return "KeychainEnabled" }
    if yname == "negotiated-local-keepalive" { return "NegotiatedLocalKeepalive" }
    if yname == "negotiated-remote-keepalive" { return "NegotiatedRemoteKeepalive" }
    if yname == "negotiated-dead-time" { return "NegotiatedDeadTime" }
    if yname == "pce-request-rx" { return "PceRequestRx" }
    if yname == "pce-request-tx" { return "PceRequestTx" }
    if yname == "pce-reply-rx" { return "PceReplyRx" }
    if yname == "pce-reply-tx" { return "PceReplyTx" }
    if yname == "pce-error-rx" { return "PceErrorRx" }
    if yname == "pce-error-tx" { return "PceErrorTx" }
    if yname == "pce-open-tx" { return "PceOpenTx" }
    if yname == "pce-open-rx" { return "PceOpenRx" }
    if yname == "pce-report-rx" { return "PceReportRx" }
    if yname == "pce-report-tx" { return "PceReportTx" }
    if yname == "pce-update-rx" { return "PceUpdateRx" }
    if yname == "pce-update-tx" { return "PceUpdateTx" }
    if yname == "pce-initiate-rx" { return "PceInitiateRx" }
    if yname == "pce-initiate-tx" { return "PceInitiateTx" }
    if yname == "pce-keepalive-tx" { return "PceKeepaliveTx" }
    if yname == "pce-keepalive-rx" { return "PceKeepaliveRx" }
    if yname == "local-session-id" { return "LocalSessionId" }
    if yname == "remote-session-id" { return "RemoteSessionId" }
    if yname == "minimum-keepalive-interval" { return "MinimumKeepaliveInterval" }
    if yname == "maximum-dead-interval" { return "MaximumDeadInterval" }
    if yname == "brief-pcep-information" { return "BriefPcepInformation" }
    if yname == "last-error-rx" { return "LastErrorRx" }
    if yname == "last-error-tx" { return "LastErrorTx" }
    return ""
}

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetSegmentPath() string {
    return "detail-pcep-information"
}

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-pcep-information" {
        return &detailPcepInformation.BriefPcepInformation
    }
    if childYangName == "last-error-rx" {
        return &detailPcepInformation.LastErrorRx
    }
    if childYangName == "last-error-tx" {
        return &detailPcepInformation.LastErrorTx
    }
    return nil
}

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-pcep-information"] = &detailPcepInformation.BriefPcepInformation
    children["last-error-rx"] = &detailPcepInformation.LastErrorRx
    children["last-error-tx"] = &detailPcepInformation.LastErrorTx
    return children
}

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["error"] = detailPcepInformation.Error
    leafs["speaker-id"] = detailPcepInformation.SpeakerId
    leafs["pcep-up-time"] = detailPcepInformation.PcepUpTime
    leafs["keepalives"] = detailPcepInformation.Keepalives
    leafs["md5-enabled"] = detailPcepInformation.Md5Enabled
    leafs["keychain-enabled"] = detailPcepInformation.KeychainEnabled
    leafs["negotiated-local-keepalive"] = detailPcepInformation.NegotiatedLocalKeepalive
    leafs["negotiated-remote-keepalive"] = detailPcepInformation.NegotiatedRemoteKeepalive
    leafs["negotiated-dead-time"] = detailPcepInformation.NegotiatedDeadTime
    leafs["pce-request-rx"] = detailPcepInformation.PceRequestRx
    leafs["pce-request-tx"] = detailPcepInformation.PceRequestTx
    leafs["pce-reply-rx"] = detailPcepInformation.PceReplyRx
    leafs["pce-reply-tx"] = detailPcepInformation.PceReplyTx
    leafs["pce-error-rx"] = detailPcepInformation.PceErrorRx
    leafs["pce-error-tx"] = detailPcepInformation.PceErrorTx
    leafs["pce-open-tx"] = detailPcepInformation.PceOpenTx
    leafs["pce-open-rx"] = detailPcepInformation.PceOpenRx
    leafs["pce-report-rx"] = detailPcepInformation.PceReportRx
    leafs["pce-report-tx"] = detailPcepInformation.PceReportTx
    leafs["pce-update-rx"] = detailPcepInformation.PceUpdateRx
    leafs["pce-update-tx"] = detailPcepInformation.PceUpdateTx
    leafs["pce-initiate-rx"] = detailPcepInformation.PceInitiateRx
    leafs["pce-initiate-tx"] = detailPcepInformation.PceInitiateTx
    leafs["pce-keepalive-tx"] = detailPcepInformation.PceKeepaliveTx
    leafs["pce-keepalive-rx"] = detailPcepInformation.PceKeepaliveRx
    leafs["local-session-id"] = detailPcepInformation.LocalSessionId
    leafs["remote-session-id"] = detailPcepInformation.RemoteSessionId
    leafs["minimum-keepalive-interval"] = detailPcepInformation.MinimumKeepaliveInterval
    leafs["maximum-dead-interval"] = detailPcepInformation.MaximumDeadInterval
    return leafs
}

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetYangName() string { return "detail-pcep-information" }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) SetParent(parent types.Entity) { detailPcepInformation.parent = parent }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetParent() types.Entity { return detailPcepInformation.parent }

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetParentYangName() string { return "peer-detail-info" }

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation
// Brief PCE protocol information
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP State. The type is PcepState.
    PcepState interface{}

    // Stateful. The type is bool.
    Stateful interface{}

    // Update capability. The type is bool.
    CapabilityUpdate interface{}

    // Instantiation capability. The type is bool.
    CapabilityInstantiate interface{}

    // Segment Routing capability. The type is bool.
    CapabilitySegmentRouting interface{}

    // Triggered Synchronization capability. The type is bool.
    CapabilityTriggeredSync interface{}

    // DB version capability. The type is bool.
    CapabilityDbVersion interface{}

    // Delta Synchronization capability. The type is bool.
    CapabilityDeltaSync interface{}
}

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetFilter() yfilter.YFilter { return briefPcepInformation.YFilter }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) SetFilter(yf yfilter.YFilter) { briefPcepInformation.YFilter = yf }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetGoName(yname string) string {
    if yname == "pcep-state" { return "PcepState" }
    if yname == "stateful" { return "Stateful" }
    if yname == "capability-update" { return "CapabilityUpdate" }
    if yname == "capability-instantiate" { return "CapabilityInstantiate" }
    if yname == "capability-segment-routing" { return "CapabilitySegmentRouting" }
    if yname == "capability-triggered-sync" { return "CapabilityTriggeredSync" }
    if yname == "capability-db-version" { return "CapabilityDbVersion" }
    if yname == "capability-delta-sync" { return "CapabilityDeltaSync" }
    return ""
}

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetSegmentPath() string {
    return "brief-pcep-information"
}

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcep-state"] = briefPcepInformation.PcepState
    leafs["stateful"] = briefPcepInformation.Stateful
    leafs["capability-update"] = briefPcepInformation.CapabilityUpdate
    leafs["capability-instantiate"] = briefPcepInformation.CapabilityInstantiate
    leafs["capability-segment-routing"] = briefPcepInformation.CapabilitySegmentRouting
    leafs["capability-triggered-sync"] = briefPcepInformation.CapabilityTriggeredSync
    leafs["capability-db-version"] = briefPcepInformation.CapabilityDbVersion
    leafs["capability-delta-sync"] = briefPcepInformation.CapabilityDeltaSync
    return leafs
}

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetYangName() string { return "brief-pcep-information" }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) SetParent(parent types.Entity) { briefPcepInformation.parent = parent }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetParent() types.Entity { return briefPcepInformation.parent }

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetParentYangName() string { return "detail-pcep-information" }

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx
// Last PCError received
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetFilter() yfilter.YFilter { return lastErrorRx.YFilter }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) SetFilter(yf yfilter.YFilter) { lastErrorRx.YFilter = yf }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetGoName(yname string) string {
    if yname == "pc-error-type" { return "PcErrorType" }
    if yname == "pc-error-value" { return "PcErrorValue" }
    return ""
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetSegmentPath() string {
    return "last-error-rx"
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pc-error-type"] = lastErrorRx.PcErrorType
    leafs["pc-error-value"] = lastErrorRx.PcErrorValue
    return leafs
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetBundleName() string { return "cisco_ios_xr" }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetYangName() string { return "last-error-rx" }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) SetParent(parent types.Entity) { lastErrorRx.parent = parent }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetParent() types.Entity { return lastErrorRx.parent }

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetParentYangName() string { return "detail-pcep-information" }

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx
// Last PCError sent
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetFilter() yfilter.YFilter { return lastErrorTx.YFilter }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) SetFilter(yf yfilter.YFilter) { lastErrorTx.YFilter = yf }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetGoName(yname string) string {
    if yname == "pc-error-type" { return "PcErrorType" }
    if yname == "pc-error-value" { return "PcErrorValue" }
    return ""
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetSegmentPath() string {
    return "last-error-tx"
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pc-error-type"] = lastErrorTx.PcErrorType
    leafs["pc-error-value"] = lastErrorTx.PcErrorValue
    return leafs
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetBundleName() string { return "cisco_ios_xr" }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetYangName() string { return "last-error-tx" }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) SetParent(parent types.Entity) { lastErrorTx.parent = parent }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetParent() types.Entity { return lastErrorTx.parent }

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetParentYangName() string { return "detail-pcep-information" }

// Pce_TopologyNodes
// Node database in XTC
type Pce_TopologyNodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node information. The type is slice of Pce_TopologyNodes_TopologyNode.
    TopologyNode []Pce_TopologyNodes_TopologyNode
}

func (topologyNodes *Pce_TopologyNodes) GetFilter() yfilter.YFilter { return topologyNodes.YFilter }

func (topologyNodes *Pce_TopologyNodes) SetFilter(yf yfilter.YFilter) { topologyNodes.YFilter = yf }

func (topologyNodes *Pce_TopologyNodes) GetGoName(yname string) string {
    if yname == "topology-node" { return "TopologyNode" }
    return ""
}

func (topologyNodes *Pce_TopologyNodes) GetSegmentPath() string {
    return "topology-nodes"
}

func (topologyNodes *Pce_TopologyNodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "topology-node" {
        for _, c := range topologyNodes.TopologyNode {
            if topologyNodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode{}
        topologyNodes.TopologyNode = append(topologyNodes.TopologyNode, child)
        return &topologyNodes.TopologyNode[len(topologyNodes.TopologyNode)-1]
    }
    return nil
}

func (topologyNodes *Pce_TopologyNodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range topologyNodes.TopologyNode {
        children[topologyNodes.TopologyNode[i].GetSegmentPath()] = &topologyNodes.TopologyNode[i]
    }
    return children
}

func (topologyNodes *Pce_TopologyNodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (topologyNodes *Pce_TopologyNodes) GetBundleName() string { return "cisco_ios_xr" }

func (topologyNodes *Pce_TopologyNodes) GetYangName() string { return "topology-nodes" }

func (topologyNodes *Pce_TopologyNodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyNodes *Pce_TopologyNodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyNodes *Pce_TopologyNodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyNodes *Pce_TopologyNodes) SetParent(parent types.Entity) { topologyNodes.parent = parent }

func (topologyNodes *Pce_TopologyNodes) GetParent() types.Entity { return topologyNodes.parent }

func (topologyNodes *Pce_TopologyNodes) GetParentYangName() string { return "pce" }

// Pce_TopologyNodes_TopologyNode
// Node information
type Pce_TopologyNodes_TopologyNode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node Identifier. The type is interface{} with
    // range: -2147483648..2147483647.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node Overload Bit. The type is bool.
    Overload interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier

    // Prefix SIDs. The type is slice of Pce_TopologyNodes_TopologyNode_PrefixSid.
    PrefixSid []Pce_TopologyNodes_TopologyNode_PrefixSid

    // IPv4 Link information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link.
    Ipv4Link []Pce_TopologyNodes_TopologyNode_Ipv4Link

    // IPv6 Link information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link.
    Ipv6Link []Pce_TopologyNodes_TopologyNode_Ipv6Link
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetFilter() yfilter.YFilter { return topologyNode.YFilter }

func (topologyNode *Pce_TopologyNodes_TopologyNode) SetFilter(yf yfilter.YFilter) { topologyNode.YFilter = yf }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetGoName(yname string) string {
    if yname == "node-identifier" { return "NodeIdentifier" }
    if yname == "node-identifier-xr" { return "NodeIdentifierXr" }
    if yname == "overload" { return "Overload" }
    if yname == "node-protocol-identifier" { return "NodeProtocolIdentifier" }
    if yname == "prefix-sid" { return "PrefixSid" }
    if yname == "ipv4-link" { return "Ipv4Link" }
    if yname == "ipv6-link" { return "Ipv6Link" }
    return ""
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetSegmentPath() string {
    return "topology-node" + "[node-identifier='" + fmt.Sprintf("%v", topologyNode.NodeIdentifier) + "']"
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-protocol-identifier" {
        return &topologyNode.NodeProtocolIdentifier
    }
    if childYangName == "prefix-sid" {
        for _, c := range topologyNode.PrefixSid {
            if topologyNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_PrefixSid{}
        topologyNode.PrefixSid = append(topologyNode.PrefixSid, child)
        return &topologyNode.PrefixSid[len(topologyNode.PrefixSid)-1]
    }
    if childYangName == "ipv4-link" {
        for _, c := range topologyNode.Ipv4Link {
            if topologyNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv4Link{}
        topologyNode.Ipv4Link = append(topologyNode.Ipv4Link, child)
        return &topologyNode.Ipv4Link[len(topologyNode.Ipv4Link)-1]
    }
    if childYangName == "ipv6-link" {
        for _, c := range topologyNode.Ipv6Link {
            if topologyNode.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv6Link{}
        topologyNode.Ipv6Link = append(topologyNode.Ipv6Link, child)
        return &topologyNode.Ipv6Link[len(topologyNode.Ipv6Link)-1]
    }
    return nil
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-protocol-identifier"] = &topologyNode.NodeProtocolIdentifier
    for i := range topologyNode.PrefixSid {
        children[topologyNode.PrefixSid[i].GetSegmentPath()] = &topologyNode.PrefixSid[i]
    }
    for i := range topologyNode.Ipv4Link {
        children[topologyNode.Ipv4Link[i].GetSegmentPath()] = &topologyNode.Ipv4Link[i]
    }
    for i := range topologyNode.Ipv6Link {
        children[topologyNode.Ipv6Link[i].GetSegmentPath()] = &topologyNode.Ipv6Link[i]
    }
    return children
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-identifier"] = topologyNode.NodeIdentifier
    leafs["node-identifier-xr"] = topologyNode.NodeIdentifierXr
    leafs["overload"] = topologyNode.Overload
    return leafs
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetBundleName() string { return "cisco_ios_xr" }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetYangName() string { return "topology-node" }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (topologyNode *Pce_TopologyNodes_TopologyNode) SetParent(parent types.Entity) { topologyNode.parent = parent }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetParent() types.Entity { return topologyNode.parent }

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetParentYangName() string { return "topology-nodes" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier
// Node protocol identifier
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetFilter() yfilter.YFilter { return nodeProtocolIdentifier.YFilter }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { nodeProtocolIdentifier.YFilter = yf }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetSegmentPath() string {
    return "node-protocol-identifier"
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range nodeProtocolIdentifier.IgpInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation{}
        nodeProtocolIdentifier.IgpInformation = append(nodeProtocolIdentifier.IgpInformation, child)
        return &nodeProtocolIdentifier.IgpInformation[len(nodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range nodeProtocolIdentifier.SrgbInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation{}
        nodeProtocolIdentifier.SrgbInformation = append(nodeProtocolIdentifier.SrgbInformation, child)
        return &nodeProtocolIdentifier.SrgbInformation[len(nodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeProtocolIdentifier.IgpInformation {
        children[nodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range nodeProtocolIdentifier.SrgbInformation {
        children[nodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = nodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = nodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = nodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = nodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetYangName() string { return "node-protocol-identifier" }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) SetParent(parent types.Entity) { nodeProtocolIdentifier.parent = parent }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetParent() types.Entity { return nodeProtocolIdentifier.parent }

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetParentYangName() string { return "topology-node" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "node-protocol-identifier" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "node-protocol-identifier" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_PrefixSid
// Prefix SIDs
type Pce_TopologyNodes_TopologyNode_PrefixSid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // R Flag. The type is bool.
    Rflag interface{}

    // N Flag. The type is bool.
    Nflag interface{}

    // P Flag. The type is bool.
    Pflag interface{}

    // E Flag. The type is bool.
    Eflag interface{}

    // V Flag. The type is bool.
    Vflag interface{}

    // L Flag. The type is bool.
    Lflag interface{}

    // Prefix.
    SidPrefix Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
}

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetFilter() yfilter.YFilter { return prefixSid.YFilter }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) SetFilter(yf yfilter.YFilter) { prefixSid.YFilter = yf }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "rflag" { return "Rflag" }
    if yname == "nflag" { return "Nflag" }
    if yname == "pflag" { return "Pflag" }
    if yname == "eflag" { return "Eflag" }
    if yname == "vflag" { return "Vflag" }
    if yname == "lflag" { return "Lflag" }
    if yname == "sid-prefix" { return "SidPrefix" }
    return ""
}

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetSegmentPath() string {
    return "prefix-sid"
}

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sid-prefix" {
        return &prefixSid.SidPrefix
    }
    return nil
}

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sid-prefix"] = &prefixSid.SidPrefix
    return children
}

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = prefixSid.SidType
    leafs["mpls-label"] = prefixSid.MplsLabel
    leafs["domain-identifier"] = prefixSid.DomainIdentifier
    leafs["rflag"] = prefixSid.Rflag
    leafs["nflag"] = prefixSid.Nflag
    leafs["pflag"] = prefixSid.Pflag
    leafs["eflag"] = prefixSid.Eflag
    leafs["vflag"] = prefixSid.Vflag
    leafs["lflag"] = prefixSid.Lflag
    return leafs
}

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetBundleName() string { return "cisco_ios_xr" }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetYangName() string { return "prefix-sid" }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) SetParent(parent types.Entity) { prefixSid.parent = parent }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetParent() types.Entity { return prefixSid.parent }

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetParentYangName() string { return "topology-node" }

// Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetFilter() yfilter.YFilter { return sidPrefix.YFilter }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) SetFilter(yf yfilter.YFilter) { sidPrefix.YFilter = yf }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetSegmentPath() string {
    return "sid-prefix"
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = sidPrefix.AfName
    leafs["ipv4"] = sidPrefix.Ipv4
    leafs["ipv6"] = sidPrefix.Ipv6
    return leafs
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetYangName() string { return "sid-prefix" }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) SetParent(parent types.Entity) { sidPrefix.parent = parent }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetParent() types.Entity { return sidPrefix.parent }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetParentYangName() string { return "prefix-sid" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type Pce_TopologyNodes_TopologyNode_Ipv4Link struct {
    parent types.Entity
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

    // SRLG Values. The type is slice of interface{} with range: 0..4294967295.
    Srlgs []interface{}

    // Local node IGP information.
    LocalIgpInformation Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid.
    AdjacencySid []Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
}

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetFilter() yfilter.YFilter { return ipv4Link.YFilter }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) SetFilter(yf yfilter.YFilter) { ipv4Link.YFilter = yf }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetGoName(yname string) string {
    if yname == "local-ipv4-address" { return "LocalIpv4Address" }
    if yname == "remote-ipv4-address" { return "RemoteIpv4Address" }
    if yname == "igp-metric" { return "IgpMetric" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "maximum-link-bandwidth" { return "MaximumLinkBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "srlgs" { return "Srlgs" }
    if yname == "local-igp-information" { return "LocalIgpInformation" }
    if yname == "remote-node-protocol-identifier" { return "RemoteNodeProtocolIdentifier" }
    if yname == "adjacency-sid" { return "AdjacencySid" }
    return ""
}

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetSegmentPath() string {
    return "ipv4-link"
}

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-igp-information" {
        return &ipv4Link.LocalIgpInformation
    }
    if childYangName == "remote-node-protocol-identifier" {
        return &ipv4Link.RemoteNodeProtocolIdentifier
    }
    if childYangName == "adjacency-sid" {
        for _, c := range ipv4Link.AdjacencySid {
            if ipv4Link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid{}
        ipv4Link.AdjacencySid = append(ipv4Link.AdjacencySid, child)
        return &ipv4Link.AdjacencySid[len(ipv4Link.AdjacencySid)-1]
    }
    return nil
}

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-igp-information"] = &ipv4Link.LocalIgpInformation
    children["remote-node-protocol-identifier"] = &ipv4Link.RemoteNodeProtocolIdentifier
    for i := range ipv4Link.AdjacencySid {
        children[ipv4Link.AdjacencySid[i].GetSegmentPath()] = &ipv4Link.AdjacencySid[i]
    }
    return children
}

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-ipv4-address"] = ipv4Link.LocalIpv4Address
    leafs["remote-ipv4-address"] = ipv4Link.RemoteIpv4Address
    leafs["igp-metric"] = ipv4Link.IgpMetric
    leafs["te-metric"] = ipv4Link.TeMetric
    leafs["maximum-link-bandwidth"] = ipv4Link.MaximumLinkBandwidth
    leafs["max-reservable-bandwidth"] = ipv4Link.MaxReservableBandwidth
    leafs["srlgs"] = ipv4Link.Srlgs
    return leafs
}

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetYangName() string { return "ipv4-link" }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) SetParent(parent types.Entity) { ipv4Link.parent = parent }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetParent() types.Entity { return ipv4Link.parent }

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetParentYangName() string { return "topology-node" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetFilter() yfilter.YFilter { return localIgpInformation.YFilter }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) SetFilter(yf yfilter.YFilter) { localIgpInformation.YFilter = yf }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetSegmentPath() string {
    return "local-igp-information"
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &localIgpInformation.Igp
    }
    return nil
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &localIgpInformation.Igp
    return children
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = localIgpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = localIgpInformation.AutonomousSystemNumber
    return leafs
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetYangName() string { return "local-igp-information" }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) SetParent(parent types.Entity) { localIgpInformation.parent = parent }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetParent() types.Entity { return localIgpInformation.parent }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetParentYangName() string { return "ipv4-link" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetParentYangName() string { return "local-igp-information" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetFilter() yfilter.YFilter { return remoteNodeProtocolIdentifier.YFilter }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { remoteNodeProtocolIdentifier.YFilter = yf }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetSegmentPath() string {
    return "remote-node-protocol-identifier"
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range remoteNodeProtocolIdentifier.IgpInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation{}
        remoteNodeProtocolIdentifier.IgpInformation = append(remoteNodeProtocolIdentifier.IgpInformation, child)
        return &remoteNodeProtocolIdentifier.IgpInformation[len(remoteNodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range remoteNodeProtocolIdentifier.SrgbInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation{}
        remoteNodeProtocolIdentifier.SrgbInformation = append(remoteNodeProtocolIdentifier.SrgbInformation, child)
        return &remoteNodeProtocolIdentifier.SrgbInformation[len(remoteNodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        children[remoteNodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        children[remoteNodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = remoteNodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = remoteNodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetYangName() string { return "remote-node-protocol-identifier" }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) SetParent(parent types.Entity) { remoteNodeProtocolIdentifier.parent = parent }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetParent() types.Entity { return remoteNodeProtocolIdentifier.parent }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetParentYangName() string { return "ipv4-link" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
// Adjacency SIDs
type Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // R Flag. The type is bool.
    Rflag interface{}

    // N Flag. The type is bool.
    Nflag interface{}

    // P Flag. The type is bool.
    Pflag interface{}

    // E Flag. The type is bool.
    Eflag interface{}

    // V Flag. The type is bool.
    Vflag interface{}

    // L Flag. The type is bool.
    Lflag interface{}

    // Prefix.
    SidPrefix Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetFilter() yfilter.YFilter { return adjacencySid.YFilter }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) SetFilter(yf yfilter.YFilter) { adjacencySid.YFilter = yf }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "rflag" { return "Rflag" }
    if yname == "nflag" { return "Nflag" }
    if yname == "pflag" { return "Pflag" }
    if yname == "eflag" { return "Eflag" }
    if yname == "vflag" { return "Vflag" }
    if yname == "lflag" { return "Lflag" }
    if yname == "sid-prefix" { return "SidPrefix" }
    return ""
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetSegmentPath() string {
    return "adjacency-sid"
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sid-prefix" {
        return &adjacencySid.SidPrefix
    }
    return nil
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sid-prefix"] = &adjacencySid.SidPrefix
    return children
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = adjacencySid.SidType
    leafs["mpls-label"] = adjacencySid.MplsLabel
    leafs["domain-identifier"] = adjacencySid.DomainIdentifier
    leafs["rflag"] = adjacencySid.Rflag
    leafs["nflag"] = adjacencySid.Nflag
    leafs["pflag"] = adjacencySid.Pflag
    leafs["eflag"] = adjacencySid.Eflag
    leafs["vflag"] = adjacencySid.Vflag
    leafs["lflag"] = adjacencySid.Lflag
    return leafs
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetBundleName() string { return "cisco_ios_xr" }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetYangName() string { return "adjacency-sid" }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) SetParent(parent types.Entity) { adjacencySid.parent = parent }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetParent() types.Entity { return adjacencySid.parent }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetParentYangName() string { return "ipv4-link" }

// Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetFilter() yfilter.YFilter { return sidPrefix.YFilter }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) SetFilter(yf yfilter.YFilter) { sidPrefix.YFilter = yf }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetSegmentPath() string {
    return "sid-prefix"
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = sidPrefix.AfName
    leafs["ipv4"] = sidPrefix.Ipv4
    leafs["ipv6"] = sidPrefix.Ipv6
    return leafs
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetYangName() string { return "sid-prefix" }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) SetParent(parent types.Entity) { sidPrefix.parent = parent }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetParent() types.Entity { return sidPrefix.parent }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetParentYangName() string { return "adjacency-sid" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type Pce_TopologyNodes_TopologyNode_Ipv6Link struct {
    parent types.Entity
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
    LocalIgpInformation Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid.
    AdjacencySid []Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
}

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetFilter() yfilter.YFilter { return ipv6Link.YFilter }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) SetFilter(yf yfilter.YFilter) { ipv6Link.YFilter = yf }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetGoName(yname string) string {
    if yname == "local-ipv6-address" { return "LocalIpv6Address" }
    if yname == "remote-ipv6-address" { return "RemoteIpv6Address" }
    if yname == "igp-metric" { return "IgpMetric" }
    if yname == "te-metric" { return "TeMetric" }
    if yname == "maximum-link-bandwidth" { return "MaximumLinkBandwidth" }
    if yname == "max-reservable-bandwidth" { return "MaxReservableBandwidth" }
    if yname == "local-igp-information" { return "LocalIgpInformation" }
    if yname == "remote-node-protocol-identifier" { return "RemoteNodeProtocolIdentifier" }
    if yname == "adjacency-sid" { return "AdjacencySid" }
    return ""
}

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetSegmentPath() string {
    return "ipv6-link"
}

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-igp-information" {
        return &ipv6Link.LocalIgpInformation
    }
    if childYangName == "remote-node-protocol-identifier" {
        return &ipv6Link.RemoteNodeProtocolIdentifier
    }
    if childYangName == "adjacency-sid" {
        for _, c := range ipv6Link.AdjacencySid {
            if ipv6Link.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid{}
        ipv6Link.AdjacencySid = append(ipv6Link.AdjacencySid, child)
        return &ipv6Link.AdjacencySid[len(ipv6Link.AdjacencySid)-1]
    }
    return nil
}

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-igp-information"] = &ipv6Link.LocalIgpInformation
    children["remote-node-protocol-identifier"] = &ipv6Link.RemoteNodeProtocolIdentifier
    for i := range ipv6Link.AdjacencySid {
        children[ipv6Link.AdjacencySid[i].GetSegmentPath()] = &ipv6Link.AdjacencySid[i]
    }
    return children
}

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-ipv6-address"] = ipv6Link.LocalIpv6Address
    leafs["remote-ipv6-address"] = ipv6Link.RemoteIpv6Address
    leafs["igp-metric"] = ipv6Link.IgpMetric
    leafs["te-metric"] = ipv6Link.TeMetric
    leafs["maximum-link-bandwidth"] = ipv6Link.MaximumLinkBandwidth
    leafs["max-reservable-bandwidth"] = ipv6Link.MaxReservableBandwidth
    return leafs
}

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetYangName() string { return "ipv6-link" }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) SetParent(parent types.Entity) { ipv6Link.parent = parent }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetParent() types.Entity { return ipv6Link.parent }

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetParentYangName() string { return "topology-node" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetFilter() yfilter.YFilter { return localIgpInformation.YFilter }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) SetFilter(yf yfilter.YFilter) { localIgpInformation.YFilter = yf }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetSegmentPath() string {
    return "local-igp-information"
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &localIgpInformation.Igp
    }
    return nil
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &localIgpInformation.Igp
    return children
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = localIgpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = localIgpInformation.AutonomousSystemNumber
    return leafs
}

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetYangName() string { return "local-igp-information" }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) SetParent(parent types.Entity) { localIgpInformation.parent = parent }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetParent() types.Entity { return localIgpInformation.parent }

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetParentYangName() string { return "ipv6-link" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetParentYangName() string { return "local-igp-information" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetFilter() yfilter.YFilter { return remoteNodeProtocolIdentifier.YFilter }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { remoteNodeProtocolIdentifier.YFilter = yf }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetSegmentPath() string {
    return "remote-node-protocol-identifier"
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range remoteNodeProtocolIdentifier.IgpInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation{}
        remoteNodeProtocolIdentifier.IgpInformation = append(remoteNodeProtocolIdentifier.IgpInformation, child)
        return &remoteNodeProtocolIdentifier.IgpInformation[len(remoteNodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range remoteNodeProtocolIdentifier.SrgbInformation {
            if remoteNodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation{}
        remoteNodeProtocolIdentifier.SrgbInformation = append(remoteNodeProtocolIdentifier.SrgbInformation, child)
        return &remoteNodeProtocolIdentifier.SrgbInformation[len(remoteNodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        children[remoteNodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        children[remoteNodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &remoteNodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = remoteNodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = remoteNodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = remoteNodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetYangName() string { return "remote-node-protocol-identifier" }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) SetParent(parent types.Entity) { remoteNodeProtocolIdentifier.parent = parent }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetParent() types.Entity { return remoteNodeProtocolIdentifier.parent }

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetParentYangName() string { return "ipv6-link" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "remote-node-protocol-identifier" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
// Adjacency SIDs
type Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // R Flag. The type is bool.
    Rflag interface{}

    // N Flag. The type is bool.
    Nflag interface{}

    // P Flag. The type is bool.
    Pflag interface{}

    // E Flag. The type is bool.
    Eflag interface{}

    // V Flag. The type is bool.
    Vflag interface{}

    // L Flag. The type is bool.
    Lflag interface{}

    // Prefix.
    SidPrefix Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetFilter() yfilter.YFilter { return adjacencySid.YFilter }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) SetFilter(yf yfilter.YFilter) { adjacencySid.YFilter = yf }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "rflag" { return "Rflag" }
    if yname == "nflag" { return "Nflag" }
    if yname == "pflag" { return "Pflag" }
    if yname == "eflag" { return "Eflag" }
    if yname == "vflag" { return "Vflag" }
    if yname == "lflag" { return "Lflag" }
    if yname == "sid-prefix" { return "SidPrefix" }
    return ""
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetSegmentPath() string {
    return "adjacency-sid"
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sid-prefix" {
        return &adjacencySid.SidPrefix
    }
    return nil
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sid-prefix"] = &adjacencySid.SidPrefix
    return children
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = adjacencySid.SidType
    leafs["mpls-label"] = adjacencySid.MplsLabel
    leafs["domain-identifier"] = adjacencySid.DomainIdentifier
    leafs["rflag"] = adjacencySid.Rflag
    leafs["nflag"] = adjacencySid.Nflag
    leafs["pflag"] = adjacencySid.Pflag
    leafs["eflag"] = adjacencySid.Eflag
    leafs["vflag"] = adjacencySid.Vflag
    leafs["lflag"] = adjacencySid.Lflag
    return leafs
}

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetBundleName() string { return "cisco_ios_xr" }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetYangName() string { return "adjacency-sid" }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) SetParent(parent types.Entity) { adjacencySid.parent = parent }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetParent() types.Entity { return adjacencySid.parent }

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetParentYangName() string { return "ipv6-link" }

// Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetFilter() yfilter.YFilter { return sidPrefix.YFilter }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) SetFilter(yf yfilter.YFilter) { sidPrefix.YFilter = yf }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetSegmentPath() string {
    return "sid-prefix"
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = sidPrefix.AfName
    leafs["ipv4"] = sidPrefix.Ipv4
    leafs["ipv6"] = sidPrefix.Ipv6
    return leafs
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetBundleName() string { return "cisco_ios_xr" }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetYangName() string { return "sid-prefix" }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) SetParent(parent types.Entity) { sidPrefix.parent = parent }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetParent() types.Entity { return sidPrefix.parent }

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetParentYangName() string { return "adjacency-sid" }

// Pce_PrefixInfos
// Prefixes database in XTC
type Pce_PrefixInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE prefix information. The type is slice of Pce_PrefixInfos_PrefixInfo.
    PrefixInfo []Pce_PrefixInfos_PrefixInfo
}

func (prefixInfos *Pce_PrefixInfos) GetFilter() yfilter.YFilter { return prefixInfos.YFilter }

func (prefixInfos *Pce_PrefixInfos) SetFilter(yf yfilter.YFilter) { prefixInfos.YFilter = yf }

func (prefixInfos *Pce_PrefixInfos) GetGoName(yname string) string {
    if yname == "prefix-info" { return "PrefixInfo" }
    return ""
}

func (prefixInfos *Pce_PrefixInfos) GetSegmentPath() string {
    return "prefix-infos"
}

func (prefixInfos *Pce_PrefixInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "prefix-info" {
        for _, c := range prefixInfos.PrefixInfo {
            if prefixInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PrefixInfos_PrefixInfo{}
        prefixInfos.PrefixInfo = append(prefixInfos.PrefixInfo, child)
        return &prefixInfos.PrefixInfo[len(prefixInfos.PrefixInfo)-1]
    }
    return nil
}

func (prefixInfos *Pce_PrefixInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range prefixInfos.PrefixInfo {
        children[prefixInfos.PrefixInfo[i].GetSegmentPath()] = &prefixInfos.PrefixInfo[i]
    }
    return children
}

func (prefixInfos *Pce_PrefixInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (prefixInfos *Pce_PrefixInfos) GetBundleName() string { return "cisco_ios_xr" }

func (prefixInfos *Pce_PrefixInfos) GetYangName() string { return "prefix-infos" }

func (prefixInfos *Pce_PrefixInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixInfos *Pce_PrefixInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixInfos *Pce_PrefixInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixInfos *Pce_PrefixInfos) SetParent(parent types.Entity) { prefixInfos.parent = parent }

func (prefixInfos *Pce_PrefixInfos) GetParent() types.Entity { return prefixInfos.parent }

func (prefixInfos *Pce_PrefixInfos) GetParentYangName() string { return "pce" }

// Pce_PrefixInfos_PrefixInfo
// PCE prefix information
type Pce_PrefixInfos_PrefixInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is interface{} with range:
    // -2147483648..2147483647.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier

    // Prefix address. The type is slice of Pce_PrefixInfos_PrefixInfo_Address.
    Address []Pce_PrefixInfos_PrefixInfo_Address
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetFilter() yfilter.YFilter { return prefixInfo.YFilter }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) SetFilter(yf yfilter.YFilter) { prefixInfo.YFilter = yf }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetGoName(yname string) string {
    if yname == "node-identifier" { return "NodeIdentifier" }
    if yname == "node-identifier-xr" { return "NodeIdentifierXr" }
    if yname == "node-protocol-identifier" { return "NodeProtocolIdentifier" }
    if yname == "address" { return "Address" }
    return ""
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetSegmentPath() string {
    return "prefix-info" + "[node-identifier='" + fmt.Sprintf("%v", prefixInfo.NodeIdentifier) + "']"
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node-protocol-identifier" {
        return &prefixInfo.NodeProtocolIdentifier
    }
    if childYangName == "address" {
        for _, c := range prefixInfo.Address {
            if prefixInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PrefixInfos_PrefixInfo_Address{}
        prefixInfo.Address = append(prefixInfo.Address, child)
        return &prefixInfo.Address[len(prefixInfo.Address)-1]
    }
    return nil
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["node-protocol-identifier"] = &prefixInfo.NodeProtocolIdentifier
    for i := range prefixInfo.Address {
        children[prefixInfo.Address[i].GetSegmentPath()] = &prefixInfo.Address[i]
    }
    return children
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-identifier"] = prefixInfo.NodeIdentifier
    leafs["node-identifier-xr"] = prefixInfo.NodeIdentifierXr
    return leafs
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetBundleName() string { return "cisco_ios_xr" }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetYangName() string { return "prefix-info" }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) SetParent(parent types.Entity) { prefixInfo.parent = parent }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetParent() types.Entity { return prefixInfo.parent }

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetParentYangName() string { return "prefix-infos" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier
// Node protocol identifier
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetFilter() yfilter.YFilter { return nodeProtocolIdentifier.YFilter }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) SetFilter(yf yfilter.YFilter) { nodeProtocolIdentifier.YFilter = yf }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "ipv4-bgp-router-id-set" { return "Ipv4BgpRouterIdSet" }
    if yname == "ipv4-bgp-router-id" { return "Ipv4BgpRouterId" }
    if yname == "ipv4te-router-id-set" { return "Ipv4TeRouterIdSet" }
    if yname == "ipv4te-router-id" { return "Ipv4TeRouterId" }
    if yname == "igp-information" { return "IgpInformation" }
    if yname == "srgb-information" { return "SrgbInformation" }
    return ""
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetSegmentPath() string {
    return "node-protocol-identifier"
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-information" {
        for _, c := range nodeProtocolIdentifier.IgpInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation{}
        nodeProtocolIdentifier.IgpInformation = append(nodeProtocolIdentifier.IgpInformation, child)
        return &nodeProtocolIdentifier.IgpInformation[len(nodeProtocolIdentifier.IgpInformation)-1]
    }
    if childYangName == "srgb-information" {
        for _, c := range nodeProtocolIdentifier.SrgbInformation {
            if nodeProtocolIdentifier.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation{}
        nodeProtocolIdentifier.SrgbInformation = append(nodeProtocolIdentifier.SrgbInformation, child)
        return &nodeProtocolIdentifier.SrgbInformation[len(nodeProtocolIdentifier.SrgbInformation)-1]
    }
    return nil
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodeProtocolIdentifier.IgpInformation {
        children[nodeProtocolIdentifier.IgpInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.IgpInformation[i]
    }
    for i := range nodeProtocolIdentifier.SrgbInformation {
        children[nodeProtocolIdentifier.SrgbInformation[i].GetSegmentPath()] = &nodeProtocolIdentifier.SrgbInformation[i]
    }
    return children
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = nodeProtocolIdentifier.NodeName
    leafs["ipv4-bgp-router-id-set"] = nodeProtocolIdentifier.Ipv4BgpRouterIdSet
    leafs["ipv4-bgp-router-id"] = nodeProtocolIdentifier.Ipv4BgpRouterId
    leafs["ipv4te-router-id-set"] = nodeProtocolIdentifier.Ipv4TeRouterIdSet
    leafs["ipv4te-router-id"] = nodeProtocolIdentifier.Ipv4TeRouterId
    return leafs
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetYangName() string { return "node-protocol-identifier" }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) SetParent(parent types.Entity) { nodeProtocolIdentifier.parent = parent }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetParent() types.Entity { return nodeProtocolIdentifier.parent }

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetParentYangName() string { return "prefix-info" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // IGP-specific information.
    Igp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
}

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetFilter() yfilter.YFilter { return igpInformation.YFilter }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) SetFilter(yf yfilter.YFilter) { igpInformation.YFilter = yf }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetGoName(yname string) string {
    if yname == "domain-identifier" { return "DomainIdentifier" }
    if yname == "autonomous-system-number" { return "AutonomousSystemNumber" }
    if yname == "igp" { return "Igp" }
    return ""
}

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetSegmentPath() string {
    return "igp-information"
}

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp" {
        return &igpInformation.Igp
    }
    return nil
}

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp"] = &igpInformation.Igp
    return children
}

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["domain-identifier"] = igpInformation.DomainIdentifier
    leafs["autonomous-system-number"] = igpInformation.AutonomousSystemNumber
    return leafs
}

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetBundleName() string { return "cisco_ios_xr" }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetYangName() string { return "igp-information" }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) SetParent(parent types.Entity) { igpInformation.parent = parent }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetParent() types.Entity { return igpInformation.parent }

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetParentYangName() string { return "node-protocol-identifier" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis

    // OSPF information.
    Ospf Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf

    // BGP information.
    Bgp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetFilter() yfilter.YFilter { return igp.YFilter }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) SetFilter(yf yfilter.YFilter) { igp.YFilter = yf }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetSegmentPath() string {
    return "igp"
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igp.Isis
    }
    if childYangName == "ospf" {
        return &igp.Ospf
    }
    if childYangName == "bgp" {
        return &igp.Bgp
    }
    return nil
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igp.Isis
    children["ospf"] = &igp.Ospf
    children["bgp"] = &igp.Bgp
    return children
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igp.IgpId
    return leafs
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleName() string { return "cisco_ios_xr" }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetYangName() string { return "igp" }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) SetParent(parent types.Entity) { igp.parent = parent }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetParent() types.Entity { return igp.parent }

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetParentYangName() string { return "igp-information" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetYangName() string { return "isis" }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetParentYangName() string { return "igp" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetParentYangName() string { return "igp" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetParentYangName() string { return "igp" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetFilter() yfilter.YFilter { return srgbInformation.YFilter }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) SetFilter(yf yfilter.YFilter) { srgbInformation.YFilter = yf }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "size" { return "Size" }
    if yname == "igp-srgb" { return "IgpSrgb" }
    return ""
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetSegmentPath() string {
    return "srgb-information"
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "igp-srgb" {
        return &srgbInformation.IgpSrgb
    }
    return nil
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["igp-srgb"] = &srgbInformation.IgpSrgb
    return children
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = srgbInformation.Start
    leafs["size"] = srgbInformation.Size
    return leafs
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetBundleName() string { return "cisco_ios_xr" }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetYangName() string { return "srgb-information" }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) SetParent(parent types.Entity) { srgbInformation.parent = parent }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetParent() types.Entity { return srgbInformation.parent }

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetParentYangName() string { return "node-protocol-identifier" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis

    // OSPF information.
    Ospf Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf

    // BGP information.
    Bgp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
}

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetFilter() yfilter.YFilter { return igpSrgb.YFilter }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetFilter(yf yfilter.YFilter) { igpSrgb.YFilter = yf }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetGoName(yname string) string {
    if yname == "igp-id" { return "IgpId" }
    if yname == "isis" { return "Isis" }
    if yname == "ospf" { return "Ospf" }
    if yname == "bgp" { return "Bgp" }
    return ""
}

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetSegmentPath() string {
    return "igp-srgb"
}

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "isis" {
        return &igpSrgb.Isis
    }
    if childYangName == "ospf" {
        return &igpSrgb.Ospf
    }
    if childYangName == "bgp" {
        return &igpSrgb.Bgp
    }
    return nil
}

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["isis"] = &igpSrgb.Isis
    children["ospf"] = &igpSrgb.Ospf
    children["bgp"] = &igpSrgb.Bgp
    return children
}

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["igp-id"] = igpSrgb.IgpId
    return leafs
}

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleName() string { return "cisco_ios_xr" }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetYangName() string { return "igp-srgb" }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) SetParent(parent types.Entity) { igpSrgb.parent = parent }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParent() types.Entity { return igpSrgb.parent }

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetParentYangName() string { return "srgb-information" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetGoName(yname string) string {
    if yname == "system-id" { return "SystemId" }
    if yname == "level" { return "Level" }
    return ""
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetSegmentPath() string {
    return "isis"
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["system-id"] = isis.SystemId
    leafs["level"] = isis.Level
    return leafs
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetYangName() string { return "isis" }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetParentYangName() string { return "igp-srgb" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "area" { return "Area" }
    return ""
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetSegmentPath() string {
    return "ospf"
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = ospf.RouterId
    leafs["area"] = ospf.Area
    return leafs
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetYangName() string { return "ospf" }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetParentYangName() string { return "igp-srgb" }

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetGoName(yname string) string {
    if yname == "router-id" { return "RouterId" }
    if yname == "confed-asn" { return "ConfedAsn" }
    return ""
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetSegmentPath() string {
    return "bgp"
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["router-id"] = bgp.RouterId
    leafs["confed-asn"] = bgp.ConfedAsn
    return leafs
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetYangName() string { return "bgp" }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetParentYangName() string { return "igp-srgb" }

// Pce_PrefixInfos_PrefixInfo_Address
// Prefix address
type Pce_PrefixInfos_PrefixInfo_Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetFilter() yfilter.YFilter { return address.YFilter }

func (address *Pce_PrefixInfos_PrefixInfo_Address) SetFilter(yf yfilter.YFilter) { address.YFilter = yf }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetSegmentPath() string {
    return "address"
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = address.AfName
    leafs["ipv4"] = address.Ipv4
    leafs["ipv6"] = address.Ipv6
    return leafs
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetBundleName() string { return "cisco_ios_xr" }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetYangName() string { return "address" }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (address *Pce_PrefixInfos_PrefixInfo_Address) SetParent(parent types.Entity) { address.parent = parent }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetParent() types.Entity { return address.parent }

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetParentYangName() string { return "prefix-info" }

// Pce_LspSummary
// LSP summary database in XTC
type Pce_LspSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary for all peers.
    AllLsPs Pce_LspSummary_AllLsPs

    // Number of LSPs for specific peer. The type is slice of
    // Pce_LspSummary_PeerLsPsInfo.
    PeerLsPsInfo []Pce_LspSummary_PeerLsPsInfo
}

func (lspSummary *Pce_LspSummary) GetFilter() yfilter.YFilter { return lspSummary.YFilter }

func (lspSummary *Pce_LspSummary) SetFilter(yf yfilter.YFilter) { lspSummary.YFilter = yf }

func (lspSummary *Pce_LspSummary) GetGoName(yname string) string {
    if yname == "all-ls-ps" { return "AllLsPs" }
    if yname == "peer-ls-ps-info" { return "PeerLsPsInfo" }
    return ""
}

func (lspSummary *Pce_LspSummary) GetSegmentPath() string {
    return "lsp-summary"
}

func (lspSummary *Pce_LspSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "all-ls-ps" {
        return &lspSummary.AllLsPs
    }
    if childYangName == "peer-ls-ps-info" {
        for _, c := range lspSummary.PeerLsPsInfo {
            if lspSummary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_LspSummary_PeerLsPsInfo{}
        lspSummary.PeerLsPsInfo = append(lspSummary.PeerLsPsInfo, child)
        return &lspSummary.PeerLsPsInfo[len(lspSummary.PeerLsPsInfo)-1]
    }
    return nil
}

func (lspSummary *Pce_LspSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["all-ls-ps"] = &lspSummary.AllLsPs
    for i := range lspSummary.PeerLsPsInfo {
        children[lspSummary.PeerLsPsInfo[i].GetSegmentPath()] = &lspSummary.PeerLsPsInfo[i]
    }
    return children
}

func (lspSummary *Pce_LspSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lspSummary *Pce_LspSummary) GetBundleName() string { return "cisco_ios_xr" }

func (lspSummary *Pce_LspSummary) GetYangName() string { return "lsp-summary" }

func (lspSummary *Pce_LspSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspSummary *Pce_LspSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspSummary *Pce_LspSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspSummary *Pce_LspSummary) SetParent(parent types.Entity) { lspSummary.parent = parent }

func (lspSummary *Pce_LspSummary) GetParent() types.Entity { return lspSummary.parent }

func (lspSummary *Pce_LspSummary) GetParentYangName() string { return "pce" }

// Pce_LspSummary_AllLsPs
// Summary for all peers
type Pce_LspSummary_AllLsPs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of all LSPs. The type is interface{} with range: 0..4294967295.
    AllLsPs interface{}

    // Number of operational LSPs. The type is interface{} with range:
    // 0..4294967295.
    UpLsPs interface{}

    // Number of administratively up LSPs. The type is interface{} with range:
    // 0..4294967295.
    AdminUpLsPs interface{}

    // Number of LSPs with Segment routing setup type. The type is interface{}
    // with range: 0..4294967295.
    SrLsPs interface{}

    // Number of LSPs with RSVP setup type. The type is interface{} with range:
    // 0..4294967295.
    RsvpLsPs interface{}
}

func (allLsPs *Pce_LspSummary_AllLsPs) GetFilter() yfilter.YFilter { return allLsPs.YFilter }

func (allLsPs *Pce_LspSummary_AllLsPs) SetFilter(yf yfilter.YFilter) { allLsPs.YFilter = yf }

func (allLsPs *Pce_LspSummary_AllLsPs) GetGoName(yname string) string {
    if yname == "all-ls-ps" { return "AllLsPs" }
    if yname == "up-ls-ps" { return "UpLsPs" }
    if yname == "admin-up-ls-ps" { return "AdminUpLsPs" }
    if yname == "sr-ls-ps" { return "SrLsPs" }
    if yname == "rsvp-ls-ps" { return "RsvpLsPs" }
    return ""
}

func (allLsPs *Pce_LspSummary_AllLsPs) GetSegmentPath() string {
    return "all-ls-ps"
}

func (allLsPs *Pce_LspSummary_AllLsPs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (allLsPs *Pce_LspSummary_AllLsPs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (allLsPs *Pce_LspSummary_AllLsPs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all-ls-ps"] = allLsPs.AllLsPs
    leafs["up-ls-ps"] = allLsPs.UpLsPs
    leafs["admin-up-ls-ps"] = allLsPs.AdminUpLsPs
    leafs["sr-ls-ps"] = allLsPs.SrLsPs
    leafs["rsvp-ls-ps"] = allLsPs.RsvpLsPs
    return leafs
}

func (allLsPs *Pce_LspSummary_AllLsPs) GetBundleName() string { return "cisco_ios_xr" }

func (allLsPs *Pce_LspSummary_AllLsPs) GetYangName() string { return "all-ls-ps" }

func (allLsPs *Pce_LspSummary_AllLsPs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (allLsPs *Pce_LspSummary_AllLsPs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (allLsPs *Pce_LspSummary_AllLsPs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (allLsPs *Pce_LspSummary_AllLsPs) SetParent(parent types.Entity) { allLsPs.parent = parent }

func (allLsPs *Pce_LspSummary_AllLsPs) GetParent() types.Entity { return allLsPs.parent }

func (allLsPs *Pce_LspSummary_AllLsPs) GetParentYangName() string { return "lsp-summary" }

// Pce_LspSummary_PeerLsPsInfo
// Number of LSPs for specific peer
type Pce_LspSummary_PeerLsPsInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Peer IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Number of LSPs for specific peer.
    LspSummary Pce_LspSummary_PeerLsPsInfo_LspSummary
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetFilter() yfilter.YFilter { return peerLsPsInfo.YFilter }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) SetFilter(yf yfilter.YFilter) { peerLsPsInfo.YFilter = yf }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "lsp-summary" { return "LspSummary" }
    return ""
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetSegmentPath() string {
    return "peer-ls-ps-info"
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsp-summary" {
        return &peerLsPsInfo.LspSummary
    }
    return nil
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsp-summary"] = &peerLsPsInfo.LspSummary
    return children
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peerLsPsInfo.PeerAddress
    return leafs
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetYangName() string { return "peer-ls-ps-info" }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) SetParent(parent types.Entity) { peerLsPsInfo.parent = parent }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetParent() types.Entity { return peerLsPsInfo.parent }

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetParentYangName() string { return "lsp-summary" }

// Pce_LspSummary_PeerLsPsInfo_LspSummary
// Number of LSPs for specific peer
type Pce_LspSummary_PeerLsPsInfo_LspSummary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of all LSPs. The type is interface{} with range: 0..4294967295.
    AllLsPs interface{}

    // Number of operational LSPs. The type is interface{} with range:
    // 0..4294967295.
    UpLsPs interface{}

    // Number of administratively up LSPs. The type is interface{} with range:
    // 0..4294967295.
    AdminUpLsPs interface{}

    // Number of LSPs with Segment routing setup type. The type is interface{}
    // with range: 0..4294967295.
    SrLsPs interface{}

    // Number of LSPs with RSVP setup type. The type is interface{} with range:
    // 0..4294967295.
    RsvpLsPs interface{}
}

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetFilter() yfilter.YFilter { return lspSummary.YFilter }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) SetFilter(yf yfilter.YFilter) { lspSummary.YFilter = yf }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetGoName(yname string) string {
    if yname == "all-ls-ps" { return "AllLsPs" }
    if yname == "up-ls-ps" { return "UpLsPs" }
    if yname == "admin-up-ls-ps" { return "AdminUpLsPs" }
    if yname == "sr-ls-ps" { return "SrLsPs" }
    if yname == "rsvp-ls-ps" { return "RsvpLsPs" }
    return ""
}

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetSegmentPath() string {
    return "lsp-summary"
}

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all-ls-ps"] = lspSummary.AllLsPs
    leafs["up-ls-ps"] = lspSummary.UpLsPs
    leafs["admin-up-ls-ps"] = lspSummary.AdminUpLsPs
    leafs["sr-ls-ps"] = lspSummary.SrLsPs
    leafs["rsvp-ls-ps"] = lspSummary.RsvpLsPs
    return leafs
}

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetBundleName() string { return "cisco_ios_xr" }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetYangName() string { return "lsp-summary" }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) SetParent(parent types.Entity) { lspSummary.parent = parent }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetParent() types.Entity { return lspSummary.parent }

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetParentYangName() string { return "peer-ls-ps-info" }

// Pce_PeerInfos
// Peers database in XTC
type Pce_PeerInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE peer information. The type is slice of Pce_PeerInfos_PeerInfo.
    PeerInfo []Pce_PeerInfos_PeerInfo
}

func (peerInfos *Pce_PeerInfos) GetFilter() yfilter.YFilter { return peerInfos.YFilter }

func (peerInfos *Pce_PeerInfos) SetFilter(yf yfilter.YFilter) { peerInfos.YFilter = yf }

func (peerInfos *Pce_PeerInfos) GetGoName(yname string) string {
    if yname == "peer-info" { return "PeerInfo" }
    return ""
}

func (peerInfos *Pce_PeerInfos) GetSegmentPath() string {
    return "peer-infos"
}

func (peerInfos *Pce_PeerInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "peer-info" {
        for _, c := range peerInfos.PeerInfo {
            if peerInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_PeerInfos_PeerInfo{}
        peerInfos.PeerInfo = append(peerInfos.PeerInfo, child)
        return &peerInfos.PeerInfo[len(peerInfos.PeerInfo)-1]
    }
    return nil
}

func (peerInfos *Pce_PeerInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range peerInfos.PeerInfo {
        children[peerInfos.PeerInfo[i].GetSegmentPath()] = &peerInfos.PeerInfo[i]
    }
    return children
}

func (peerInfos *Pce_PeerInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (peerInfos *Pce_PeerInfos) GetBundleName() string { return "cisco_ios_xr" }

func (peerInfos *Pce_PeerInfos) GetYangName() string { return "peer-infos" }

func (peerInfos *Pce_PeerInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerInfos *Pce_PeerInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerInfos *Pce_PeerInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerInfos *Pce_PeerInfos) SetParent(parent types.Entity) { peerInfos.parent = parent }

func (peerInfos *Pce_PeerInfos) GetParent() types.Entity { return peerInfos.parent }

func (peerInfos *Pce_PeerInfos) GetParentYangName() string { return "pce" }

// Pce_PeerInfos_PeerInfo
// PCE peer information
type Pce_PeerInfos_PeerInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // PCE protocol information.
    BriefPcepInformation Pce_PeerInfos_PeerInfo_BriefPcepInformation
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetFilter() yfilter.YFilter { return peerInfo.YFilter }

func (peerInfo *Pce_PeerInfos_PeerInfo) SetFilter(yf yfilter.YFilter) { peerInfo.YFilter = yf }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "peer-address-xr" { return "PeerAddressXr" }
    if yname == "peer-protocol" { return "PeerProtocol" }
    if yname == "brief-pcep-information" { return "BriefPcepInformation" }
    return ""
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetSegmentPath() string {
    return "peer-info" + "[peer-address='" + fmt.Sprintf("%v", peerInfo.PeerAddress) + "']"
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-pcep-information" {
        return &peerInfo.BriefPcepInformation
    }
    return nil
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-pcep-information"] = &peerInfo.BriefPcepInformation
    return children
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = peerInfo.PeerAddress
    leafs["peer-address-xr"] = peerInfo.PeerAddressXr
    leafs["peer-protocol"] = peerInfo.PeerProtocol
    return leafs
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetBundleName() string { return "cisco_ios_xr" }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetYangName() string { return "peer-info" }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (peerInfo *Pce_PeerInfos_PeerInfo) SetParent(parent types.Entity) { peerInfo.parent = parent }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetParent() types.Entity { return peerInfo.parent }

func (peerInfo *Pce_PeerInfos_PeerInfo) GetParentYangName() string { return "peer-infos" }

// Pce_PeerInfos_PeerInfo_BriefPcepInformation
// PCE protocol information
type Pce_PeerInfos_PeerInfo_BriefPcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCEP State. The type is PcepState.
    PcepState interface{}

    // Stateful. The type is bool.
    Stateful interface{}

    // Update capability. The type is bool.
    CapabilityUpdate interface{}

    // Instantiation capability. The type is bool.
    CapabilityInstantiate interface{}

    // Segment Routing capability. The type is bool.
    CapabilitySegmentRouting interface{}

    // Triggered Synchronization capability. The type is bool.
    CapabilityTriggeredSync interface{}

    // DB version capability. The type is bool.
    CapabilityDbVersion interface{}

    // Delta Synchronization capability. The type is bool.
    CapabilityDeltaSync interface{}
}

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetFilter() yfilter.YFilter { return briefPcepInformation.YFilter }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) SetFilter(yf yfilter.YFilter) { briefPcepInformation.YFilter = yf }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetGoName(yname string) string {
    if yname == "pcep-state" { return "PcepState" }
    if yname == "stateful" { return "Stateful" }
    if yname == "capability-update" { return "CapabilityUpdate" }
    if yname == "capability-instantiate" { return "CapabilityInstantiate" }
    if yname == "capability-segment-routing" { return "CapabilitySegmentRouting" }
    if yname == "capability-triggered-sync" { return "CapabilityTriggeredSync" }
    if yname == "capability-db-version" { return "CapabilityDbVersion" }
    if yname == "capability-delta-sync" { return "CapabilityDeltaSync" }
    return ""
}

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetSegmentPath() string {
    return "brief-pcep-information"
}

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcep-state"] = briefPcepInformation.PcepState
    leafs["stateful"] = briefPcepInformation.Stateful
    leafs["capability-update"] = briefPcepInformation.CapabilityUpdate
    leafs["capability-instantiate"] = briefPcepInformation.CapabilityInstantiate
    leafs["capability-segment-routing"] = briefPcepInformation.CapabilitySegmentRouting
    leafs["capability-triggered-sync"] = briefPcepInformation.CapabilityTriggeredSync
    leafs["capability-db-version"] = briefPcepInformation.CapabilityDbVersion
    leafs["capability-delta-sync"] = briefPcepInformation.CapabilityDeltaSync
    return leafs
}

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetYangName() string { return "brief-pcep-information" }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) SetParent(parent types.Entity) { briefPcepInformation.parent = parent }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetParent() types.Entity { return briefPcepInformation.parent }

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetParentYangName() string { return "peer-info" }

// Pce_TunnelDetailInfos
// Detailed tunnel database in XTC
type Pce_TunnelDetailInfos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detailed tunnel information. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo.
    TunnelDetailInfo []Pce_TunnelDetailInfos_TunnelDetailInfo
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetFilter() yfilter.YFilter { return tunnelDetailInfos.YFilter }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) SetFilter(yf yfilter.YFilter) { tunnelDetailInfos.YFilter = yf }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetGoName(yname string) string {
    if yname == "tunnel-detail-info" { return "TunnelDetailInfo" }
    return ""
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetSegmentPath() string {
    return "tunnel-detail-infos"
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tunnel-detail-info" {
        for _, c := range tunnelDetailInfos.TunnelDetailInfo {
            if tunnelDetailInfos.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo{}
        tunnelDetailInfos.TunnelDetailInfo = append(tunnelDetailInfos.TunnelDetailInfo, child)
        return &tunnelDetailInfos.TunnelDetailInfo[len(tunnelDetailInfos.TunnelDetailInfo)-1]
    }
    return nil
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tunnelDetailInfos.TunnelDetailInfo {
        children[tunnelDetailInfos.TunnelDetailInfo[i].GetSegmentPath()] = &tunnelDetailInfos.TunnelDetailInfo[i]
    }
    return children
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetYangName() string { return "tunnel-detail-infos" }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) SetParent(parent types.Entity) { tunnelDetailInfos.parent = parent }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetParent() types.Entity { return tunnelDetailInfos.parent }

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetParentYangName() string { return "pce" }

// Pce_TunnelDetailInfos_TunnelDetailInfo
// Detailed tunnel information
type Pce_TunnelDetailInfos_TunnelDetailInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Private LSP information.
    PrivateLspInformation Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation

    // Detail LSP information. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation.
    DetailLspInformation []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetFilter() yfilter.YFilter { return tunnelDetailInfo.YFilter }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) SetFilter(yf yfilter.YFilter) { tunnelDetailInfo.YFilter = yf }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetGoName(yname string) string {
    if yname == "peer-address" { return "PeerAddress" }
    if yname == "plsp-id" { return "PlspId" }
    if yname == "tunnel-name" { return "TunnelName" }
    if yname == "pcc-address" { return "PccAddress" }
    if yname == "tunnel-name-xr" { return "TunnelNameXr" }
    if yname == "private-lsp-information" { return "PrivateLspInformation" }
    if yname == "detail-lsp-information" { return "DetailLspInformation" }
    return ""
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetSegmentPath() string {
    return "tunnel-detail-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelDetailInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelDetailInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelDetailInfo.TunnelName) + "']"
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "private-lsp-information" {
        return &tunnelDetailInfo.PrivateLspInformation
    }
    if childYangName == "detail-lsp-information" {
        for _, c := range tunnelDetailInfo.DetailLspInformation {
            if tunnelDetailInfo.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation{}
        tunnelDetailInfo.DetailLspInformation = append(tunnelDetailInfo.DetailLspInformation, child)
        return &tunnelDetailInfo.DetailLspInformation[len(tunnelDetailInfo.DetailLspInformation)-1]
    }
    return nil
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["private-lsp-information"] = &tunnelDetailInfo.PrivateLspInformation
    for i := range tunnelDetailInfo.DetailLspInformation {
        children[tunnelDetailInfo.DetailLspInformation[i].GetSegmentPath()] = &tunnelDetailInfo.DetailLspInformation[i]
    }
    return children
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["peer-address"] = tunnelDetailInfo.PeerAddress
    leafs["plsp-id"] = tunnelDetailInfo.PlspId
    leafs["tunnel-name"] = tunnelDetailInfo.TunnelName
    leafs["pcc-address"] = tunnelDetailInfo.PccAddress
    leafs["tunnel-name-xr"] = tunnelDetailInfo.TunnelNameXr
    return leafs
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetBundleName() string { return "cisco_ios_xr" }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetYangName() string { return "tunnel-detail-info" }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) SetParent(parent types.Entity) { tunnelDetailInfo.parent = parent }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetParent() types.Entity { return tunnelDetailInfo.parent }

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetParentYangName() string { return "tunnel-detail-infos" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation
// Private LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // LSP Event buffer. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer.
    EventBuffer []Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetFilter() yfilter.YFilter { return privateLspInformation.YFilter }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) SetFilter(yf yfilter.YFilter) { privateLspInformation.YFilter = yf }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetGoName(yname string) string {
    if yname == "event-buffer" { return "EventBuffer" }
    return ""
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetSegmentPath() string {
    return "private-lsp-information"
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "event-buffer" {
        for _, c := range privateLspInformation.EventBuffer {
            if privateLspInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer{}
        privateLspInformation.EventBuffer = append(privateLspInformation.EventBuffer, child)
        return &privateLspInformation.EventBuffer[len(privateLspInformation.EventBuffer)-1]
    }
    return nil
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range privateLspInformation.EventBuffer {
        children[privateLspInformation.EventBuffer[i].GetSegmentPath()] = &privateLspInformation.EventBuffer[i]
    }
    return children
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetYangName() string { return "private-lsp-information" }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) SetParent(parent types.Entity) { privateLspInformation.parent = parent }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetParent() types.Entity { return privateLspInformation.parent }

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetParentYangName() string { return "tunnel-detail-info" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
// LSP Event buffer
type Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Event message. The type is string.
    EventMessage interface{}

    // Event time, relative to Jan 1, 1970. The type is interface{} with range:
    // 0..4294967295.
    TimeStamp interface{}
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetFilter() yfilter.YFilter { return eventBuffer.YFilter }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) SetFilter(yf yfilter.YFilter) { eventBuffer.YFilter = yf }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetGoName(yname string) string {
    if yname == "event-message" { return "EventMessage" }
    if yname == "time-stamp" { return "TimeStamp" }
    return ""
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetSegmentPath() string {
    return "event-buffer"
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["event-message"] = eventBuffer.EventMessage
    leafs["time-stamp"] = eventBuffer.TimeStamp
    return leafs
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetBundleName() string { return "cisco_ios_xr" }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetYangName() string { return "event-buffer" }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) SetParent(parent types.Entity) { eventBuffer.parent = parent }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetParent() types.Entity { return eventBuffer.parent }

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetParentYangName() string { return "private-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
// Detail LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // True if router notifies signal bandwidth. The type is bool.
    SignaledBandwidthSpecified interface{}

    // Signaled Bandwidth. The type is interface{} with range:
    // 0..18446744073709551615.
    SignaledBandwidth interface{}

    // True if router notifies actual bandwidth. The type is bool.
    ActualBandwidthSpecified interface{}

    // Actual bandwidth utilized in the data-plane. The type is interface{} with
    // range: 0..18446744073709551615.
    ActualBandwidth interface{}

    // LSP Role. The type is interface{} with range: 0..4294967295.
    LspRole interface{}

    // Computing PCE. The type is interface{} with range: 0..4294967295.
    ComputingPce interface{}

    // Sub delegated PCE. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubDelegatedPce interface{}

    // State-sync PCE. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    StateSyncPce interface{}

    // Reporting PCC address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ReportingPccAddress interface{}

    // List of SLRGs used by LSP. The type is slice of interface{} with range:
    // 0..4294967295.
    SrlgInfo []interface{}

    // Brief LSP information.
    BriefLspInformation Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation

    // Paths.
    ErOs Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs

    // PCEP related LSP information.
    LsppcepInformation Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation

    // LSP association information.
    LspAssociationInfo Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo

    // LSP attributes.
    LspAttributes Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes

    // RRO. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro.
    Rro []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
}

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetFilter() yfilter.YFilter { return detailLspInformation.YFilter }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) SetFilter(yf yfilter.YFilter) { detailLspInformation.YFilter = yf }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetGoName(yname string) string {
    if yname == "signaled-bandwidth-specified" { return "SignaledBandwidthSpecified" }
    if yname == "signaled-bandwidth" { return "SignaledBandwidth" }
    if yname == "actual-bandwidth-specified" { return "ActualBandwidthSpecified" }
    if yname == "actual-bandwidth" { return "ActualBandwidth" }
    if yname == "lsp-role" { return "LspRole" }
    if yname == "computing-pce" { return "ComputingPce" }
    if yname == "sub-delegated-pce" { return "SubDelegatedPce" }
    if yname == "state-sync-pce" { return "StateSyncPce" }
    if yname == "reporting-pcc-address" { return "ReportingPccAddress" }
    if yname == "srlg-info" { return "SrlgInfo" }
    if yname == "brief-lsp-information" { return "BriefLspInformation" }
    if yname == "er-os" { return "ErOs" }
    if yname == "lsppcep-information" { return "LsppcepInformation" }
    if yname == "lsp-association-info" { return "LspAssociationInfo" }
    if yname == "lsp-attributes" { return "LspAttributes" }
    if yname == "rro" { return "Rro" }
    return ""
}

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetSegmentPath() string {
    return "detail-lsp-information"
}

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-lsp-information" {
        return &detailLspInformation.BriefLspInformation
    }
    if childYangName == "er-os" {
        return &detailLspInformation.ErOs
    }
    if childYangName == "lsppcep-information" {
        return &detailLspInformation.LsppcepInformation
    }
    if childYangName == "lsp-association-info" {
        return &detailLspInformation.LspAssociationInfo
    }
    if childYangName == "lsp-attributes" {
        return &detailLspInformation.LspAttributes
    }
    if childYangName == "rro" {
        for _, c := range detailLspInformation.Rro {
            if detailLspInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro{}
        detailLspInformation.Rro = append(detailLspInformation.Rro, child)
        return &detailLspInformation.Rro[len(detailLspInformation.Rro)-1]
    }
    return nil
}

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-lsp-information"] = &detailLspInformation.BriefLspInformation
    children["er-os"] = &detailLspInformation.ErOs
    children["lsppcep-information"] = &detailLspInformation.LsppcepInformation
    children["lsp-association-info"] = &detailLspInformation.LspAssociationInfo
    children["lsp-attributes"] = &detailLspInformation.LspAttributes
    for i := range detailLspInformation.Rro {
        children[detailLspInformation.Rro[i].GetSegmentPath()] = &detailLspInformation.Rro[i]
    }
    return children
}

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["signaled-bandwidth-specified"] = detailLspInformation.SignaledBandwidthSpecified
    leafs["signaled-bandwidth"] = detailLspInformation.SignaledBandwidth
    leafs["actual-bandwidth-specified"] = detailLspInformation.ActualBandwidthSpecified
    leafs["actual-bandwidth"] = detailLspInformation.ActualBandwidth
    leafs["lsp-role"] = detailLspInformation.LspRole
    leafs["computing-pce"] = detailLspInformation.ComputingPce
    leafs["sub-delegated-pce"] = detailLspInformation.SubDelegatedPce
    leafs["state-sync-pce"] = detailLspInformation.StateSyncPce
    leafs["reporting-pcc-address"] = detailLspInformation.ReportingPccAddress
    leafs["srlg-info"] = detailLspInformation.SrlgInfo
    return leafs
}

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetYangName() string { return "detail-lsp-information" }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) SetParent(parent types.Entity) { detailLspInformation.parent = parent }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetParent() types.Entity { return detailLspInformation.parent }

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetParentYangName() string { return "tunnel-detail-info" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation
// Brief LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DestinationAddress interface{}

    // Tunnel ID. The type is interface{} with range: 0..4294967295.
    TunnelId interface{}

    // LSP ID. The type is interface{} with range: 0..4294967295.
    Lspid interface{}

    // Binding SID. The type is interface{} with range: 0..4294967295.
    BindingSid interface{}

    // LSP Setup Type. The type is LspSetup.
    LspSetupType interface{}

    // Operational state. The type is PcepLspState.
    OperationalState interface{}

    // Admin state. The type is LspState.
    AdministrativeState interface{}
}

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetFilter() yfilter.YFilter { return briefLspInformation.YFilter }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) SetFilter(yf yfilter.YFilter) { briefLspInformation.YFilter = yf }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetGoName(yname string) string {
    if yname == "source-address" { return "SourceAddress" }
    if yname == "destination-address" { return "DestinationAddress" }
    if yname == "tunnel-id" { return "TunnelId" }
    if yname == "lspid" { return "Lspid" }
    if yname == "binding-sid" { return "BindingSid" }
    if yname == "lsp-setup-type" { return "LspSetupType" }
    if yname == "operational-state" { return "OperationalState" }
    if yname == "administrative-state" { return "AdministrativeState" }
    return ""
}

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetSegmentPath() string {
    return "brief-lsp-information"
}

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["source-address"] = briefLspInformation.SourceAddress
    leafs["destination-address"] = briefLspInformation.DestinationAddress
    leafs["tunnel-id"] = briefLspInformation.TunnelId
    leafs["lspid"] = briefLspInformation.Lspid
    leafs["binding-sid"] = briefLspInformation.BindingSid
    leafs["lsp-setup-type"] = briefLspInformation.LspSetupType
    leafs["operational-state"] = briefLspInformation.OperationalState
    leafs["administrative-state"] = briefLspInformation.AdministrativeState
    return leafs
}

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetYangName() string { return "brief-lsp-information" }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) SetParent(parent types.Entity) { briefLspInformation.parent = parent }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetParent() types.Entity { return briefLspInformation.parent }

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetParentYangName() string { return "detail-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs
// Paths
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Reported Metric Type. The type is interface{} with range: 0..4294967295.
    ReportedMetricType interface{}

    // Reported Metric Value. The type is interface{} with range: 0..4294967295.
    ReportedMetricValue interface{}

    // Computed Metric Type. The type is interface{} with range: 0..4294967295.
    ComputedMetricType interface{}

    // Computed Metric Value. The type is interface{} with range: 0..4294967295.
    ComputedMetricValue interface{}

    // Computed Hop List Time. The type is interface{} with range: 0..4294967295.
    ComputedHopListTime interface{}

    // Reported RSVP path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath.
    ReportedRsvpPath []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath

    // Reported SR path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath.
    ReportedSrPath []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath

    // Computed RSVP path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath.
    ComputedRsvpPath []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath

    // Computed SR path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath.
    ComputedSrPath []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
}

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetFilter() yfilter.YFilter { return erOs.YFilter }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) SetFilter(yf yfilter.YFilter) { erOs.YFilter = yf }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetGoName(yname string) string {
    if yname == "reported-metric-type" { return "ReportedMetricType" }
    if yname == "reported-metric-value" { return "ReportedMetricValue" }
    if yname == "computed-metric-type" { return "ComputedMetricType" }
    if yname == "computed-metric-value" { return "ComputedMetricValue" }
    if yname == "computed-hop-list-time" { return "ComputedHopListTime" }
    if yname == "reported-rsvp-path" { return "ReportedRsvpPath" }
    if yname == "reported-sr-path" { return "ReportedSrPath" }
    if yname == "computed-rsvp-path" { return "ComputedRsvpPath" }
    if yname == "computed-sr-path" { return "ComputedSrPath" }
    return ""
}

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetSegmentPath() string {
    return "er-os"
}

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reported-rsvp-path" {
        for _, c := range erOs.ReportedRsvpPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath{}
        erOs.ReportedRsvpPath = append(erOs.ReportedRsvpPath, child)
        return &erOs.ReportedRsvpPath[len(erOs.ReportedRsvpPath)-1]
    }
    if childYangName == "reported-sr-path" {
        for _, c := range erOs.ReportedSrPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath{}
        erOs.ReportedSrPath = append(erOs.ReportedSrPath, child)
        return &erOs.ReportedSrPath[len(erOs.ReportedSrPath)-1]
    }
    if childYangName == "computed-rsvp-path" {
        for _, c := range erOs.ComputedRsvpPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath{}
        erOs.ComputedRsvpPath = append(erOs.ComputedRsvpPath, child)
        return &erOs.ComputedRsvpPath[len(erOs.ComputedRsvpPath)-1]
    }
    if childYangName == "computed-sr-path" {
        for _, c := range erOs.ComputedSrPath {
            if erOs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath{}
        erOs.ComputedSrPath = append(erOs.ComputedSrPath, child)
        return &erOs.ComputedSrPath[len(erOs.ComputedSrPath)-1]
    }
    return nil
}

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range erOs.ReportedRsvpPath {
        children[erOs.ReportedRsvpPath[i].GetSegmentPath()] = &erOs.ReportedRsvpPath[i]
    }
    for i := range erOs.ReportedSrPath {
        children[erOs.ReportedSrPath[i].GetSegmentPath()] = &erOs.ReportedSrPath[i]
    }
    for i := range erOs.ComputedRsvpPath {
        children[erOs.ComputedRsvpPath[i].GetSegmentPath()] = &erOs.ComputedRsvpPath[i]
    }
    for i := range erOs.ComputedSrPath {
        children[erOs.ComputedSrPath[i].GetSegmentPath()] = &erOs.ComputedSrPath[i]
    }
    return children
}

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["reported-metric-type"] = erOs.ReportedMetricType
    leafs["reported-metric-value"] = erOs.ReportedMetricValue
    leafs["computed-metric-type"] = erOs.ComputedMetricType
    leafs["computed-metric-value"] = erOs.ComputedMetricValue
    leafs["computed-hop-list-time"] = erOs.ComputedHopListTime
    return leafs
}

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetBundleName() string { return "cisco_ios_xr" }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetYangName() string { return "er-os" }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) SetParent(parent types.Entity) { erOs.parent = parent }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetParent() types.Entity { return erOs.parent }

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetParentYangName() string { return "detail-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath
// Reported RSVP path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HopAddress interface{}
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetFilter() yfilter.YFilter { return reportedRsvpPath.YFilter }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) SetFilter(yf yfilter.YFilter) { reportedRsvpPath.YFilter = yf }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetGoName(yname string) string {
    if yname == "hop-address" { return "HopAddress" }
    return ""
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetSegmentPath() string {
    return "reported-rsvp-path"
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-address"] = reportedRsvpPath.HopAddress
    return leafs
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetBundleName() string { return "cisco_ios_xr" }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetYangName() string { return "reported-rsvp-path" }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) SetParent(parent types.Entity) { reportedRsvpPath.parent = parent }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetParent() types.Entity { return reportedRsvpPath.parent }

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetParentYangName() string { return "er-os" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath
// Reported SR path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddr interface{}
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetFilter() yfilter.YFilter { return reportedSrPath.YFilter }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) SetFilter(yf yfilter.YFilter) { reportedSrPath.YFilter = yf }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-addr" { return "RemoteAddr" }
    return ""
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetSegmentPath() string {
    return "reported-sr-path"
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = reportedSrPath.SidType
    leafs["mpls-label"] = reportedSrPath.MplsLabel
    leafs["local-addr"] = reportedSrPath.LocalAddr
    leafs["remote-addr"] = reportedSrPath.RemoteAddr
    return leafs
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetBundleName() string { return "cisco_ios_xr" }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetYangName() string { return "reported-sr-path" }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) SetParent(parent types.Entity) { reportedSrPath.parent = parent }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetParent() types.Entity { return reportedSrPath.parent }

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetParentYangName() string { return "er-os" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath
// Computed RSVP path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    HopAddress interface{}
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetFilter() yfilter.YFilter { return computedRsvpPath.YFilter }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) SetFilter(yf yfilter.YFilter) { computedRsvpPath.YFilter = yf }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetGoName(yname string) string {
    if yname == "hop-address" { return "HopAddress" }
    return ""
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetSegmentPath() string {
    return "computed-rsvp-path"
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hop-address"] = computedRsvpPath.HopAddress
    return leafs
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetBundleName() string { return "cisco_ios_xr" }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetYangName() string { return "computed-rsvp-path" }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) SetParent(parent types.Entity) { computedRsvpPath.parent = parent }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetParent() types.Entity { return computedRsvpPath.parent }

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetParentYangName() string { return "er-os" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
// Computed SR path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddr interface{}
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetFilter() yfilter.YFilter { return computedSrPath.YFilter }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) SetFilter(yf yfilter.YFilter) { computedSrPath.YFilter = yf }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-addr" { return "RemoteAddr" }
    return ""
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetSegmentPath() string {
    return "computed-sr-path"
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = computedSrPath.SidType
    leafs["mpls-label"] = computedSrPath.MplsLabel
    leafs["local-addr"] = computedSrPath.LocalAddr
    leafs["remote-addr"] = computedSrPath.RemoteAddr
    return leafs
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetBundleName() string { return "cisco_ios_xr" }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetYangName() string { return "computed-sr-path" }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) SetParent(parent types.Entity) { computedSrPath.parent = parent }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetParent() types.Entity { return computedSrPath.parent }

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetParentYangName() string { return "er-os" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation
// PCEP related LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCE protocol identifier. The type is interface{} with range: 0..4294967295.
    Pcepid interface{}

    // PCEP LSP delegation flag. The type is bool.
    PcepFlagD interface{}

    // PCEP LSP state-sync flag. The type is bool.
    PcepFlagS interface{}

    // PCEP LSP remove flag. The type is bool.
    PcepFlagR interface{}

    // PCEP LSP admin flag. The type is bool.
    PcepFlagA interface{}

    // PCEP LSP operation flag. The type is interface{} with range: 0..255.
    PcepFlagO interface{}

    // RSVP error info.
    RsvpError Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetFilter() yfilter.YFilter { return lsppcepInformation.YFilter }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) SetFilter(yf yfilter.YFilter) { lsppcepInformation.YFilter = yf }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetGoName(yname string) string {
    if yname == "pcepid" { return "Pcepid" }
    if yname == "pcep-flag-d" { return "PcepFlagD" }
    if yname == "pcep-flag-s" { return "PcepFlagS" }
    if yname == "pcep-flag-r" { return "PcepFlagR" }
    if yname == "pcep-flag-a" { return "PcepFlagA" }
    if yname == "pcep-flag-o" { return "PcepFlagO" }
    if yname == "rsvp-error" { return "RsvpError" }
    return ""
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetSegmentPath() string {
    return "lsppcep-information"
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rsvp-error" {
        return &lsppcepInformation.RsvpError
    }
    return nil
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rsvp-error"] = &lsppcepInformation.RsvpError
    return children
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcepid"] = lsppcepInformation.Pcepid
    leafs["pcep-flag-d"] = lsppcepInformation.PcepFlagD
    leafs["pcep-flag-s"] = lsppcepInformation.PcepFlagS
    leafs["pcep-flag-r"] = lsppcepInformation.PcepFlagR
    leafs["pcep-flag-a"] = lsppcepInformation.PcepFlagA
    leafs["pcep-flag-o"] = lsppcepInformation.PcepFlagO
    return leafs
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetBundleName() string { return "cisco_ios_xr" }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetYangName() string { return "lsppcep-information" }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) SetParent(parent types.Entity) { lsppcepInformation.parent = parent }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetParent() types.Entity { return lsppcepInformation.parent }

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetParentYangName() string { return "detail-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
// RSVP error info
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RSVP error node address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NodeAddress interface{}

    // RSVP error flags. The type is interface{} with range: 0..255.
    ErrorFlags interface{}

    // RSVP error code. The type is interface{} with range: 0..255.
    ErrorCode interface{}

    // RSVP error value. The type is interface{} with range: 0..65535.
    ErrorValue interface{}
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetFilter() yfilter.YFilter { return rsvpError.YFilter }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) SetFilter(yf yfilter.YFilter) { rsvpError.YFilter = yf }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetGoName(yname string) string {
    if yname == "node-address" { return "NodeAddress" }
    if yname == "error-flags" { return "ErrorFlags" }
    if yname == "error-code" { return "ErrorCode" }
    if yname == "error-value" { return "ErrorValue" }
    return ""
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetSegmentPath() string {
    return "rsvp-error"
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-address"] = rsvpError.NodeAddress
    leafs["error-flags"] = rsvpError.ErrorFlags
    leafs["error-code"] = rsvpError.ErrorCode
    leafs["error-value"] = rsvpError.ErrorValue
    return leafs
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetBundleName() string { return "cisco_ios_xr" }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetYangName() string { return "rsvp-error" }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) SetParent(parent types.Entity) { rsvpError.parent = parent }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetParent() types.Entity { return rsvpError.parent }

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetParentYangName() string { return "lsppcep-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo
// LSP association information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Source. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    AssociationSource interface{}
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetFilter() yfilter.YFilter { return lspAssociationInfo.YFilter }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) SetFilter(yf yfilter.YFilter) { lspAssociationInfo.YFilter = yf }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetGoName(yname string) string {
    if yname == "association-type" { return "AssociationType" }
    if yname == "association-id" { return "AssociationId" }
    if yname == "association-source" { return "AssociationSource" }
    return ""
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetSegmentPath() string {
    return "lsp-association-info"
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["association-type"] = lspAssociationInfo.AssociationType
    leafs["association-id"] = lspAssociationInfo.AssociationId
    leafs["association-source"] = lspAssociationInfo.AssociationSource
    return leafs
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetBundleName() string { return "cisco_ios_xr" }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetYangName() string { return "lsp-association-info" }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) SetParent(parent types.Entity) { lspAssociationInfo.parent = parent }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetParent() types.Entity { return lspAssociationInfo.parent }

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetParentYangName() string { return "detail-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes
// LSP attributes
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Affinity exclude any. The type is interface{} with range: 0..4294967295.
    AffinityExcludeAny interface{}

    // Affinity include any. The type is interface{} with range: 0..4294967295.
    AffinityIncludeAny interface{}

    // Affinity include all. The type is interface{} with range: 0..4294967295.
    AffinityIncludeAll interface{}

    // Setup Priority. The type is interface{} with range: 0..255.
    SetupPriority interface{}

    // Hold Priority. The type is interface{} with range: 0..255.
    HoldPriority interface{}

    // True, if local protection is desired. The type is bool.
    LocalProtection interface{}
}

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetFilter() yfilter.YFilter { return lspAttributes.YFilter }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) SetFilter(yf yfilter.YFilter) { lspAttributes.YFilter = yf }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetGoName(yname string) string {
    if yname == "affinity-exclude-any" { return "AffinityExcludeAny" }
    if yname == "affinity-include-any" { return "AffinityIncludeAny" }
    if yname == "affinity-include-all" { return "AffinityIncludeAll" }
    if yname == "setup-priority" { return "SetupPriority" }
    if yname == "hold-priority" { return "HoldPriority" }
    if yname == "local-protection" { return "LocalProtection" }
    return ""
}

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetSegmentPath() string {
    return "lsp-attributes"
}

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["affinity-exclude-any"] = lspAttributes.AffinityExcludeAny
    leafs["affinity-include-any"] = lspAttributes.AffinityIncludeAny
    leafs["affinity-include-all"] = lspAttributes.AffinityIncludeAll
    leafs["setup-priority"] = lspAttributes.SetupPriority
    leafs["hold-priority"] = lspAttributes.HoldPriority
    leafs["local-protection"] = lspAttributes.LocalProtection
    return leafs
}

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetBundleName() string { return "cisco_ios_xr" }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetYangName() string { return "lsp-attributes" }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) SetParent(parent types.Entity) { lspAttributes.parent = parent }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetParent() types.Entity { return lspAttributes.parent }

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetParentYangName() string { return "detail-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
// RRO
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RRO Type. The type is PceRro.
    RroType interface{}

    // IPv4 address of RRO. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // MPLS label of RRO. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // RRO Flags. The type is interface{} with range: 0..255.
    Flags interface{}

    // Segment Routing RRO info.
    SrRro Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetFilter() yfilter.YFilter { return rro.YFilter }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) SetFilter(yf yfilter.YFilter) { rro.YFilter = yf }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetGoName(yname string) string {
    if yname == "rro-type" { return "RroType" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "flags" { return "Flags" }
    if yname == "sr-rro" { return "SrRro" }
    return ""
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetSegmentPath() string {
    return "rro"
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sr-rro" {
        return &rro.SrRro
    }
    return nil
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sr-rro"] = &rro.SrRro
    return children
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rro-type"] = rro.RroType
    leafs["ipv4-address"] = rro.Ipv4Address
    leafs["mpls-label"] = rro.MplsLabel
    leafs["flags"] = rro.Flags
    return leafs
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetBundleName() string { return "cisco_ios_xr" }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetYangName() string { return "rro" }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) SetParent(parent types.Entity) { rro.parent = parent }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetParent() types.Entity { return rro.parent }

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetParentYangName() string { return "detail-lsp-information" }

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
// Segment Routing RRO info
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RemoteAddr interface{}
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetFilter() yfilter.YFilter { return srRro.YFilter }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) SetFilter(yf yfilter.YFilter) { srRro.YFilter = yf }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetGoName(yname string) string {
    if yname == "sid-type" { return "SidType" }
    if yname == "mpls-label" { return "MplsLabel" }
    if yname == "local-addr" { return "LocalAddr" }
    if yname == "remote-addr" { return "RemoteAddr" }
    return ""
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetSegmentPath() string {
    return "sr-rro"
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sid-type"] = srRro.SidType
    leafs["mpls-label"] = srRro.MplsLabel
    leafs["local-addr"] = srRro.LocalAddr
    leafs["remote-addr"] = srRro.RemoteAddr
    return leafs
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetBundleName() string { return "cisco_ios_xr" }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetYangName() string { return "sr-rro" }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) SetParent(parent types.Entity) { srRro.parent = parent }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetParent() types.Entity { return srRro.parent }

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetParentYangName() string { return "rro" }

