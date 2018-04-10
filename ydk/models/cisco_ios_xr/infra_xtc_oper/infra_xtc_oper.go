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

// PceProto represents PCE peer protocol
type PceProto string

const (
    // PCE protocol
    PceProto_pcep PceProto = "pcep"

    // Netconf protocol
    PceProto_netconf PceProto = "netconf"
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

// LspState represents LSP setup type
type LspState string

const (
    // LSP is down
    LspState_lsp_down LspState = "lsp-down"

    // LSP is up
    LspState_lsp_up LspState = "lsp-up"
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

// PceLspData
// PCE LSP's data
type PceLspData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel database in XTC.
    TunnelInfos PceLspData_TunnelInfos

    // LSP summary database in XTC.
    LspSummary PceLspData_LspSummary

    // Detailed tunnel database in XTC.
    TunnelDetailInfos PceLspData_TunnelDetailInfos
}

func (pceLspData *PceLspData) GetEntityData() *types.CommonEntityData {
    pceLspData.EntityData.YFilter = pceLspData.YFilter
    pceLspData.EntityData.YangName = "pce-lsp-data"
    pceLspData.EntityData.BundleName = "cisco_ios_xr"
    pceLspData.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-oper"
    pceLspData.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-oper:pce-lsp-data"
    pceLspData.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceLspData.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceLspData.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceLspData.EntityData.Children = make(map[string]types.YChild)
    pceLspData.EntityData.Children["tunnel-infos"] = types.YChild{"TunnelInfos", &pceLspData.TunnelInfos}
    pceLspData.EntityData.Children["lsp-summary"] = types.YChild{"LspSummary", &pceLspData.LspSummary}
    pceLspData.EntityData.Children["tunnel-detail-infos"] = types.YChild{"TunnelDetailInfos", &pceLspData.TunnelDetailInfos}
    pceLspData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pceLspData.EntityData)
}

// PceLspData_TunnelInfos
// Tunnel database in XTC
type PceLspData_TunnelInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of PceLspData_TunnelInfos_TunnelInfo.
    TunnelInfo []PceLspData_TunnelInfos_TunnelInfo
}

func (tunnelInfos *PceLspData_TunnelInfos) GetEntityData() *types.CommonEntityData {
    tunnelInfos.EntityData.YFilter = tunnelInfos.YFilter
    tunnelInfos.EntityData.YangName = "tunnel-infos"
    tunnelInfos.EntityData.BundleName = "cisco_ios_xr"
    tunnelInfos.EntityData.ParentYangName = "pce-lsp-data"
    tunnelInfos.EntityData.SegmentPath = "tunnel-infos"
    tunnelInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInfos.EntityData.Children = make(map[string]types.YChild)
    tunnelInfos.EntityData.Children["tunnel-info"] = types.YChild{"TunnelInfo", nil}
    for i := range tunnelInfos.TunnelInfo {
        tunnelInfos.EntityData.Children[types.GetSegmentPath(&tunnelInfos.TunnelInfo[i])] = types.YChild{"TunnelInfo", &tunnelInfos.TunnelInfo[i]}
    }
    tunnelInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelInfos.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo
// Tunnel information
type PceLspData_TunnelInfos_TunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Brief LSP information. The type is slice of
    // PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation.
    BriefLspInformation []PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetEntityData() *types.CommonEntityData {
    tunnelInfo.EntityData.YFilter = tunnelInfo.YFilter
    tunnelInfo.EntityData.YangName = "tunnel-info"
    tunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelInfo.EntityData.ParentYangName = "tunnel-infos"
    tunnelInfo.EntityData.SegmentPath = "tunnel-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelInfo.TunnelName) + "']"
    tunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInfo.EntityData.Children = make(map[string]types.YChild)
    tunnelInfo.EntityData.Children["brief-lsp-information"] = types.YChild{"BriefLspInformation", nil}
    for i := range tunnelInfo.BriefLspInformation {
        tunnelInfo.EntityData.Children[types.GetSegmentPath(&tunnelInfo.BriefLspInformation[i])] = types.YChild{"BriefLspInformation", &tunnelInfo.BriefLspInformation[i]}
    }
    tunnelInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", tunnelInfo.PeerAddress}
    tunnelInfo.EntityData.Leafs["plsp-id"] = types.YLeaf{"PlspId", tunnelInfo.PlspId}
    tunnelInfo.EntityData.Leafs["tunnel-name"] = types.YLeaf{"TunnelName", tunnelInfo.TunnelName}
    tunnelInfo.EntityData.Leafs["pcc-address"] = types.YLeaf{"PccAddress", tunnelInfo.PccAddress}
    tunnelInfo.EntityData.Leafs["tunnel-name-xr"] = types.YLeaf{"TunnelNameXr", tunnelInfo.TunnelNameXr}
    return &(tunnelInfo.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation
// Brief LSP information
type PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (briefLspInformation *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation) GetEntityData() *types.CommonEntityData {
    briefLspInformation.EntityData.YFilter = briefLspInformation.YFilter
    briefLspInformation.EntityData.YangName = "brief-lsp-information"
    briefLspInformation.EntityData.BundleName = "cisco_ios_xr"
    briefLspInformation.EntityData.ParentYangName = "tunnel-info"
    briefLspInformation.EntityData.SegmentPath = "brief-lsp-information"
    briefLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLspInformation.EntityData.Children = make(map[string]types.YChild)
    briefLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefLspInformation.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", briefLspInformation.SourceAddress}
    briefLspInformation.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", briefLspInformation.DestinationAddress}
    briefLspInformation.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", briefLspInformation.TunnelId}
    briefLspInformation.EntityData.Leafs["lspid"] = types.YLeaf{"Lspid", briefLspInformation.Lspid}
    briefLspInformation.EntityData.Leafs["binding-sid"] = types.YLeaf{"BindingSid", briefLspInformation.BindingSid}
    briefLspInformation.EntityData.Leafs["lsp-setup-type"] = types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType}
    briefLspInformation.EntityData.Leafs["operational-state"] = types.YLeaf{"OperationalState", briefLspInformation.OperationalState}
    briefLspInformation.EntityData.Leafs["administrative-state"] = types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState}
    return &(briefLspInformation.EntityData)
}

// PceLspData_LspSummary
// LSP summary database in XTC
type PceLspData_LspSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary for all peers.
    AllLsPs PceLspData_LspSummary_AllLsPs

    // Number of LSPs for specific peer. The type is slice of
    // PceLspData_LspSummary_PeerLsPsInfo.
    PeerLsPsInfo []PceLspData_LspSummary_PeerLsPsInfo
}

func (lspSummary *PceLspData_LspSummary) GetEntityData() *types.CommonEntityData {
    lspSummary.EntityData.YFilter = lspSummary.YFilter
    lspSummary.EntityData.YangName = "lsp-summary"
    lspSummary.EntityData.BundleName = "cisco_ios_xr"
    lspSummary.EntityData.ParentYangName = "pce-lsp-data"
    lspSummary.EntityData.SegmentPath = "lsp-summary"
    lspSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspSummary.EntityData.Children = make(map[string]types.YChild)
    lspSummary.EntityData.Children["all-ls-ps"] = types.YChild{"AllLsPs", &lspSummary.AllLsPs}
    lspSummary.EntityData.Children["peer-ls-ps-info"] = types.YChild{"PeerLsPsInfo", nil}
    for i := range lspSummary.PeerLsPsInfo {
        lspSummary.EntityData.Children[types.GetSegmentPath(&lspSummary.PeerLsPsInfo[i])] = types.YChild{"PeerLsPsInfo", &lspSummary.PeerLsPsInfo[i]}
    }
    lspSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspSummary.EntityData)
}

// PceLspData_LspSummary_AllLsPs
// Summary for all peers
type PceLspData_LspSummary_AllLsPs struct {
    EntityData types.CommonEntityData
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

func (allLsPs *PceLspData_LspSummary_AllLsPs) GetEntityData() *types.CommonEntityData {
    allLsPs.EntityData.YFilter = allLsPs.YFilter
    allLsPs.EntityData.YangName = "all-ls-ps"
    allLsPs.EntityData.BundleName = "cisco_ios_xr"
    allLsPs.EntityData.ParentYangName = "lsp-summary"
    allLsPs.EntityData.SegmentPath = "all-ls-ps"
    allLsPs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLsPs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLsPs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLsPs.EntityData.Children = make(map[string]types.YChild)
    allLsPs.EntityData.Leafs = make(map[string]types.YLeaf)
    allLsPs.EntityData.Leafs["all-ls-ps"] = types.YLeaf{"AllLsPs", allLsPs.AllLsPs}
    allLsPs.EntityData.Leafs["up-ls-ps"] = types.YLeaf{"UpLsPs", allLsPs.UpLsPs}
    allLsPs.EntityData.Leafs["admin-up-ls-ps"] = types.YLeaf{"AdminUpLsPs", allLsPs.AdminUpLsPs}
    allLsPs.EntityData.Leafs["sr-ls-ps"] = types.YLeaf{"SrLsPs", allLsPs.SrLsPs}
    allLsPs.EntityData.Leafs["rsvp-ls-ps"] = types.YLeaf{"RsvpLsPs", allLsPs.RsvpLsPs}
    return &(allLsPs.EntityData)
}

// PceLspData_LspSummary_PeerLsPsInfo
// Number of LSPs for specific peer
type PceLspData_LspSummary_PeerLsPsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Number of LSPs for specific peer.
    LspSummary PceLspData_LspSummary_PeerLsPsInfo_LspSummary_
}

func (peerLsPsInfo *PceLspData_LspSummary_PeerLsPsInfo) GetEntityData() *types.CommonEntityData {
    peerLsPsInfo.EntityData.YFilter = peerLsPsInfo.YFilter
    peerLsPsInfo.EntityData.YangName = "peer-ls-ps-info"
    peerLsPsInfo.EntityData.BundleName = "cisco_ios_xr"
    peerLsPsInfo.EntityData.ParentYangName = "lsp-summary"
    peerLsPsInfo.EntityData.SegmentPath = "peer-ls-ps-info"
    peerLsPsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerLsPsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerLsPsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerLsPsInfo.EntityData.Children = make(map[string]types.YChild)
    peerLsPsInfo.EntityData.Children["lsp-summary"] = types.YChild{"LspSummary", &peerLsPsInfo.LspSummary}
    peerLsPsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerLsPsInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", peerLsPsInfo.PeerAddress}
    return &(peerLsPsInfo.EntityData)
}

// PceLspData_LspSummary_PeerLsPsInfo_LspSummary_
// Number of LSPs for specific peer
type PceLspData_LspSummary_PeerLsPsInfo_LspSummary_ struct {
    EntityData types.CommonEntityData
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

func (lspSummary_ *PceLspData_LspSummary_PeerLsPsInfo_LspSummary_) GetEntityData() *types.CommonEntityData {
    lspSummary_.EntityData.YFilter = lspSummary_.YFilter
    lspSummary_.EntityData.YangName = "lsp-summary"
    lspSummary_.EntityData.BundleName = "cisco_ios_xr"
    lspSummary_.EntityData.ParentYangName = "peer-ls-ps-info"
    lspSummary_.EntityData.SegmentPath = "lsp-summary"
    lspSummary_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspSummary_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspSummary_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspSummary_.EntityData.Children = make(map[string]types.YChild)
    lspSummary_.EntityData.Leafs = make(map[string]types.YLeaf)
    lspSummary_.EntityData.Leafs["all-ls-ps"] = types.YLeaf{"AllLsPs", lspSummary_.AllLsPs}
    lspSummary_.EntityData.Leafs["up-ls-ps"] = types.YLeaf{"UpLsPs", lspSummary_.UpLsPs}
    lspSummary_.EntityData.Leafs["admin-up-ls-ps"] = types.YLeaf{"AdminUpLsPs", lspSummary_.AdminUpLsPs}
    lspSummary_.EntityData.Leafs["sr-ls-ps"] = types.YLeaf{"SrLsPs", lspSummary_.SrLsPs}
    lspSummary_.EntityData.Leafs["rsvp-ls-ps"] = types.YLeaf{"RsvpLsPs", lspSummary_.RsvpLsPs}
    return &(lspSummary_.EntityData)
}

// PceLspData_TunnelDetailInfos
// Detailed tunnel database in XTC
type PceLspData_TunnelDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed tunnel information. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo.
    TunnelDetailInfo []PceLspData_TunnelDetailInfos_TunnelDetailInfo
}

func (tunnelDetailInfos *PceLspData_TunnelDetailInfos) GetEntityData() *types.CommonEntityData {
    tunnelDetailInfos.EntityData.YFilter = tunnelDetailInfos.YFilter
    tunnelDetailInfos.EntityData.YangName = "tunnel-detail-infos"
    tunnelDetailInfos.EntityData.BundleName = "cisco_ios_xr"
    tunnelDetailInfos.EntityData.ParentYangName = "pce-lsp-data"
    tunnelDetailInfos.EntityData.SegmentPath = "tunnel-detail-infos"
    tunnelDetailInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDetailInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDetailInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDetailInfos.EntityData.Children = make(map[string]types.YChild)
    tunnelDetailInfos.EntityData.Children["tunnel-detail-info"] = types.YChild{"TunnelDetailInfo", nil}
    for i := range tunnelDetailInfos.TunnelDetailInfo {
        tunnelDetailInfos.EntityData.Children[types.GetSegmentPath(&tunnelDetailInfos.TunnelDetailInfo[i])] = types.YChild{"TunnelDetailInfo", &tunnelDetailInfos.TunnelDetailInfo[i]}
    }
    tunnelDetailInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelDetailInfos.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo
// Detailed tunnel information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Private LSP information.
    PrivateLspInformation PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation

    // Detail LSP information. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation.
    DetailLspInformation []PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetEntityData() *types.CommonEntityData {
    tunnelDetailInfo.EntityData.YFilter = tunnelDetailInfo.YFilter
    tunnelDetailInfo.EntityData.YangName = "tunnel-detail-info"
    tunnelDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelDetailInfo.EntityData.ParentYangName = "tunnel-detail-infos"
    tunnelDetailInfo.EntityData.SegmentPath = "tunnel-detail-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelDetailInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelDetailInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelDetailInfo.TunnelName) + "']"
    tunnelDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDetailInfo.EntityData.Children = make(map[string]types.YChild)
    tunnelDetailInfo.EntityData.Children["private-lsp-information"] = types.YChild{"PrivateLspInformation", &tunnelDetailInfo.PrivateLspInformation}
    tunnelDetailInfo.EntityData.Children["detail-lsp-information"] = types.YChild{"DetailLspInformation", nil}
    for i := range tunnelDetailInfo.DetailLspInformation {
        tunnelDetailInfo.EntityData.Children[types.GetSegmentPath(&tunnelDetailInfo.DetailLspInformation[i])] = types.YChild{"DetailLspInformation", &tunnelDetailInfo.DetailLspInformation[i]}
    }
    tunnelDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelDetailInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", tunnelDetailInfo.PeerAddress}
    tunnelDetailInfo.EntityData.Leafs["plsp-id"] = types.YLeaf{"PlspId", tunnelDetailInfo.PlspId}
    tunnelDetailInfo.EntityData.Leafs["tunnel-name"] = types.YLeaf{"TunnelName", tunnelDetailInfo.TunnelName}
    tunnelDetailInfo.EntityData.Leafs["pcc-address"] = types.YLeaf{"PccAddress", tunnelDetailInfo.PccAddress}
    tunnelDetailInfo.EntityData.Leafs["tunnel-name-xr"] = types.YLeaf{"TunnelNameXr", tunnelDetailInfo.TunnelNameXr}
    return &(tunnelDetailInfo.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation
// Private LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Event buffer. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer.
    EventBuffer []PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
}

func (privateLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetEntityData() *types.CommonEntityData {
    privateLspInformation.EntityData.YFilter = privateLspInformation.YFilter
    privateLspInformation.EntityData.YangName = "private-lsp-information"
    privateLspInformation.EntityData.BundleName = "cisco_ios_xr"
    privateLspInformation.EntityData.ParentYangName = "tunnel-detail-info"
    privateLspInformation.EntityData.SegmentPath = "private-lsp-information"
    privateLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateLspInformation.EntityData.Children = make(map[string]types.YChild)
    privateLspInformation.EntityData.Children["event-buffer"] = types.YChild{"EventBuffer", nil}
    for i := range privateLspInformation.EventBuffer {
        privateLspInformation.EntityData.Children[types.GetSegmentPath(&privateLspInformation.EventBuffer[i])] = types.YChild{"EventBuffer", &privateLspInformation.EventBuffer[i]}
    }
    privateLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(privateLspInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
// LSP Event buffer
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event message. The type is string.
    EventMessage interface{}

    // Event time, relative to Jan 1, 1970. The type is interface{} with range:
    // 0..4294967295.
    TimeStamp interface{}
}

func (eventBuffer *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetEntityData() *types.CommonEntityData {
    eventBuffer.EntityData.YFilter = eventBuffer.YFilter
    eventBuffer.EntityData.YangName = "event-buffer"
    eventBuffer.EntityData.BundleName = "cisco_ios_xr"
    eventBuffer.EntityData.ParentYangName = "private-lsp-information"
    eventBuffer.EntityData.SegmentPath = "event-buffer"
    eventBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventBuffer.EntityData.Children = make(map[string]types.YChild)
    eventBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eventBuffer.EntityData.Leafs["event-message"] = types.YLeaf{"EventMessage", eventBuffer.EventMessage}
    eventBuffer.EntityData.Leafs["time-stamp"] = types.YLeaf{"TimeStamp", eventBuffer.TimeStamp}
    return &(eventBuffer.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
// Detail LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SubDelegatedPce interface{}

    // State-sync PCE. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    StateSyncPce interface{}

    // Reporting PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (detailLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetEntityData() *types.CommonEntityData {
    detailLspInformation.EntityData.YFilter = detailLspInformation.YFilter
    detailLspInformation.EntityData.YangName = "detail-lsp-information"
    detailLspInformation.EntityData.BundleName = "cisco_ios_xr"
    detailLspInformation.EntityData.ParentYangName = "tunnel-detail-info"
    detailLspInformation.EntityData.SegmentPath = "detail-lsp-information"
    detailLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailLspInformation.EntityData.Children = make(map[string]types.YChild)
    detailLspInformation.EntityData.Children["brief-lsp-information"] = types.YChild{"BriefLspInformation", &detailLspInformation.BriefLspInformation}
    detailLspInformation.EntityData.Children["er-os"] = types.YChild{"ErOs", &detailLspInformation.ErOs}
    detailLspInformation.EntityData.Children["lsppcep-information"] = types.YChild{"LsppcepInformation", &detailLspInformation.LsppcepInformation}
    detailLspInformation.EntityData.Children["lsp-association-info"] = types.YChild{"LspAssociationInfo", &detailLspInformation.LspAssociationInfo}
    detailLspInformation.EntityData.Children["lsp-attributes"] = types.YChild{"LspAttributes", &detailLspInformation.LspAttributes}
    detailLspInformation.EntityData.Children["rro"] = types.YChild{"Rro", nil}
    for i := range detailLspInformation.Rro {
        detailLspInformation.EntityData.Children[types.GetSegmentPath(&detailLspInformation.Rro[i])] = types.YChild{"Rro", &detailLspInformation.Rro[i]}
    }
    detailLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    detailLspInformation.EntityData.Leafs["signaled-bandwidth-specified"] = types.YLeaf{"SignaledBandwidthSpecified", detailLspInformation.SignaledBandwidthSpecified}
    detailLspInformation.EntityData.Leafs["signaled-bandwidth"] = types.YLeaf{"SignaledBandwidth", detailLspInformation.SignaledBandwidth}
    detailLspInformation.EntityData.Leafs["actual-bandwidth-specified"] = types.YLeaf{"ActualBandwidthSpecified", detailLspInformation.ActualBandwidthSpecified}
    detailLspInformation.EntityData.Leafs["actual-bandwidth"] = types.YLeaf{"ActualBandwidth", detailLspInformation.ActualBandwidth}
    detailLspInformation.EntityData.Leafs["lsp-role"] = types.YLeaf{"LspRole", detailLspInformation.LspRole}
    detailLspInformation.EntityData.Leafs["computing-pce"] = types.YLeaf{"ComputingPce", detailLspInformation.ComputingPce}
    detailLspInformation.EntityData.Leafs["sub-delegated-pce"] = types.YLeaf{"SubDelegatedPce", detailLspInformation.SubDelegatedPce}
    detailLspInformation.EntityData.Leafs["state-sync-pce"] = types.YLeaf{"StateSyncPce", detailLspInformation.StateSyncPce}
    detailLspInformation.EntityData.Leafs["reporting-pcc-address"] = types.YLeaf{"ReportingPccAddress", detailLspInformation.ReportingPccAddress}
    detailLspInformation.EntityData.Leafs["srlg-info"] = types.YLeaf{"SrlgInfo", detailLspInformation.SrlgInfo}
    return &(detailLspInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation
// Brief LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (briefLspInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetEntityData() *types.CommonEntityData {
    briefLspInformation.EntityData.YFilter = briefLspInformation.YFilter
    briefLspInformation.EntityData.YangName = "brief-lsp-information"
    briefLspInformation.EntityData.BundleName = "cisco_ios_xr"
    briefLspInformation.EntityData.ParentYangName = "detail-lsp-information"
    briefLspInformation.EntityData.SegmentPath = "brief-lsp-information"
    briefLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLspInformation.EntityData.Children = make(map[string]types.YChild)
    briefLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefLspInformation.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", briefLspInformation.SourceAddress}
    briefLspInformation.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", briefLspInformation.DestinationAddress}
    briefLspInformation.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", briefLspInformation.TunnelId}
    briefLspInformation.EntityData.Leafs["lspid"] = types.YLeaf{"Lspid", briefLspInformation.Lspid}
    briefLspInformation.EntityData.Leafs["binding-sid"] = types.YLeaf{"BindingSid", briefLspInformation.BindingSid}
    briefLspInformation.EntityData.Leafs["lsp-setup-type"] = types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType}
    briefLspInformation.EntityData.Leafs["operational-state"] = types.YLeaf{"OperationalState", briefLspInformation.OperationalState}
    briefLspInformation.EntityData.Leafs["administrative-state"] = types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState}
    return &(briefLspInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs
// Paths
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs struct {
    EntityData types.CommonEntityData
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

func (erOs *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetEntityData() *types.CommonEntityData {
    erOs.EntityData.YFilter = erOs.YFilter
    erOs.EntityData.YangName = "er-os"
    erOs.EntityData.BundleName = "cisco_ios_xr"
    erOs.EntityData.ParentYangName = "detail-lsp-information"
    erOs.EntityData.SegmentPath = "er-os"
    erOs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erOs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erOs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erOs.EntityData.Children = make(map[string]types.YChild)
    erOs.EntityData.Children["reported-rsvp-path"] = types.YChild{"ReportedRsvpPath", nil}
    for i := range erOs.ReportedRsvpPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ReportedRsvpPath[i])] = types.YChild{"ReportedRsvpPath", &erOs.ReportedRsvpPath[i]}
    }
    erOs.EntityData.Children["reported-sr-path"] = types.YChild{"ReportedSrPath", nil}
    for i := range erOs.ReportedSrPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ReportedSrPath[i])] = types.YChild{"ReportedSrPath", &erOs.ReportedSrPath[i]}
    }
    erOs.EntityData.Children["computed-rsvp-path"] = types.YChild{"ComputedRsvpPath", nil}
    for i := range erOs.ComputedRsvpPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ComputedRsvpPath[i])] = types.YChild{"ComputedRsvpPath", &erOs.ComputedRsvpPath[i]}
    }
    erOs.EntityData.Children["computed-sr-path"] = types.YChild{"ComputedSrPath", nil}
    for i := range erOs.ComputedSrPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ComputedSrPath[i])] = types.YChild{"ComputedSrPath", &erOs.ComputedSrPath[i]}
    }
    erOs.EntityData.Leafs = make(map[string]types.YLeaf)
    erOs.EntityData.Leafs["reported-metric-type"] = types.YLeaf{"ReportedMetricType", erOs.ReportedMetricType}
    erOs.EntityData.Leafs["reported-metric-value"] = types.YLeaf{"ReportedMetricValue", erOs.ReportedMetricValue}
    erOs.EntityData.Leafs["computed-metric-type"] = types.YLeaf{"ComputedMetricType", erOs.ComputedMetricType}
    erOs.EntityData.Leafs["computed-metric-value"] = types.YLeaf{"ComputedMetricValue", erOs.ComputedMetricValue}
    erOs.EntityData.Leafs["computed-hop-list-time"] = types.YLeaf{"ComputedHopListTime", erOs.ComputedHopListTime}
    return &(erOs.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath
// Reported RSVP path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}
}

func (reportedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetEntityData() *types.CommonEntityData {
    reportedRsvpPath.EntityData.YFilter = reportedRsvpPath.YFilter
    reportedRsvpPath.EntityData.YangName = "reported-rsvp-path"
    reportedRsvpPath.EntityData.BundleName = "cisco_ios_xr"
    reportedRsvpPath.EntityData.ParentYangName = "er-os"
    reportedRsvpPath.EntityData.SegmentPath = "reported-rsvp-path"
    reportedRsvpPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportedRsvpPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportedRsvpPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportedRsvpPath.EntityData.Children = make(map[string]types.YChild)
    reportedRsvpPath.EntityData.Leafs = make(map[string]types.YLeaf)
    reportedRsvpPath.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", reportedRsvpPath.HopAddress}
    return &(reportedRsvpPath.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath
// Reported SR path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddr interface{}
}

func (reportedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetEntityData() *types.CommonEntityData {
    reportedSrPath.EntityData.YFilter = reportedSrPath.YFilter
    reportedSrPath.EntityData.YangName = "reported-sr-path"
    reportedSrPath.EntityData.BundleName = "cisco_ios_xr"
    reportedSrPath.EntityData.ParentYangName = "er-os"
    reportedSrPath.EntityData.SegmentPath = "reported-sr-path"
    reportedSrPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportedSrPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportedSrPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportedSrPath.EntityData.Children = make(map[string]types.YChild)
    reportedSrPath.EntityData.Leafs = make(map[string]types.YLeaf)
    reportedSrPath.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", reportedSrPath.SidType}
    reportedSrPath.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", reportedSrPath.MplsLabel}
    reportedSrPath.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", reportedSrPath.LocalAddr}
    reportedSrPath.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", reportedSrPath.RemoteAddr}
    return &(reportedSrPath.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath
// Computed RSVP path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}
}

func (computedRsvpPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetEntityData() *types.CommonEntityData {
    computedRsvpPath.EntityData.YFilter = computedRsvpPath.YFilter
    computedRsvpPath.EntityData.YangName = "computed-rsvp-path"
    computedRsvpPath.EntityData.BundleName = "cisco_ios_xr"
    computedRsvpPath.EntityData.ParentYangName = "er-os"
    computedRsvpPath.EntityData.SegmentPath = "computed-rsvp-path"
    computedRsvpPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    computedRsvpPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    computedRsvpPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    computedRsvpPath.EntityData.Children = make(map[string]types.YChild)
    computedRsvpPath.EntityData.Leafs = make(map[string]types.YLeaf)
    computedRsvpPath.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", computedRsvpPath.HopAddress}
    return &(computedRsvpPath.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
// Computed SR path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddr interface{}
}

func (computedSrPath *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetEntityData() *types.CommonEntityData {
    computedSrPath.EntityData.YFilter = computedSrPath.YFilter
    computedSrPath.EntityData.YangName = "computed-sr-path"
    computedSrPath.EntityData.BundleName = "cisco_ios_xr"
    computedSrPath.EntityData.ParentYangName = "er-os"
    computedSrPath.EntityData.SegmentPath = "computed-sr-path"
    computedSrPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    computedSrPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    computedSrPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    computedSrPath.EntityData.Children = make(map[string]types.YChild)
    computedSrPath.EntityData.Leafs = make(map[string]types.YLeaf)
    computedSrPath.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", computedSrPath.SidType}
    computedSrPath.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", computedSrPath.MplsLabel}
    computedSrPath.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", computedSrPath.LocalAddr}
    computedSrPath.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", computedSrPath.RemoteAddr}
    return &(computedSrPath.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation
// PCEP related LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation struct {
    EntityData types.CommonEntityData
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

    // PCEP LSP initiated flag. The type is interface{} with range: 0..255.
    PcepFlagC interface{}

    // RSVP error info.
    RsvpError PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
}

func (lsppcepInformation *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetEntityData() *types.CommonEntityData {
    lsppcepInformation.EntityData.YFilter = lsppcepInformation.YFilter
    lsppcepInformation.EntityData.YangName = "lsppcep-information"
    lsppcepInformation.EntityData.BundleName = "cisco_ios_xr"
    lsppcepInformation.EntityData.ParentYangName = "detail-lsp-information"
    lsppcepInformation.EntityData.SegmentPath = "lsppcep-information"
    lsppcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsppcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsppcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsppcepInformation.EntityData.Children = make(map[string]types.YChild)
    lsppcepInformation.EntityData.Children["rsvp-error"] = types.YChild{"RsvpError", &lsppcepInformation.RsvpError}
    lsppcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    lsppcepInformation.EntityData.Leafs["pcepid"] = types.YLeaf{"Pcepid", lsppcepInformation.Pcepid}
    lsppcepInformation.EntityData.Leafs["pcep-flag-d"] = types.YLeaf{"PcepFlagD", lsppcepInformation.PcepFlagD}
    lsppcepInformation.EntityData.Leafs["pcep-flag-s"] = types.YLeaf{"PcepFlagS", lsppcepInformation.PcepFlagS}
    lsppcepInformation.EntityData.Leafs["pcep-flag-r"] = types.YLeaf{"PcepFlagR", lsppcepInformation.PcepFlagR}
    lsppcepInformation.EntityData.Leafs["pcep-flag-a"] = types.YLeaf{"PcepFlagA", lsppcepInformation.PcepFlagA}
    lsppcepInformation.EntityData.Leafs["pcep-flag-o"] = types.YLeaf{"PcepFlagO", lsppcepInformation.PcepFlagO}
    lsppcepInformation.EntityData.Leafs["pcep-flag-c"] = types.YLeaf{"PcepFlagC", lsppcepInformation.PcepFlagC}
    return &(lsppcepInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
// RSVP error info
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP error node address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NodeAddress interface{}

    // RSVP error flags. The type is interface{} with range: 0..255.
    ErrorFlags interface{}

    // RSVP error code. The type is interface{} with range: 0..255.
    ErrorCode interface{}

    // RSVP error value. The type is interface{} with range: 0..65535.
    ErrorValue interface{}
}

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetEntityData() *types.CommonEntityData {
    rsvpError.EntityData.YFilter = rsvpError.YFilter
    rsvpError.EntityData.YangName = "rsvp-error"
    rsvpError.EntityData.BundleName = "cisco_ios_xr"
    rsvpError.EntityData.ParentYangName = "lsppcep-information"
    rsvpError.EntityData.SegmentPath = "rsvp-error"
    rsvpError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvpError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvpError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvpError.EntityData.Children = make(map[string]types.YChild)
    rsvpError.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpError.EntityData.Leafs["node-address"] = types.YLeaf{"NodeAddress", rsvpError.NodeAddress}
    rsvpError.EntityData.Leafs["error-flags"] = types.YLeaf{"ErrorFlags", rsvpError.ErrorFlags}
    rsvpError.EntityData.Leafs["error-code"] = types.YLeaf{"ErrorCode", rsvpError.ErrorCode}
    rsvpError.EntityData.Leafs["error-value"] = types.YLeaf{"ErrorValue", rsvpError.ErrorValue}
    return &(rsvpError.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo
// LSP association information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AssociationSource interface{}
}

func (lspAssociationInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetEntityData() *types.CommonEntityData {
    lspAssociationInfo.EntityData.YFilter = lspAssociationInfo.YFilter
    lspAssociationInfo.EntityData.YangName = "lsp-association-info"
    lspAssociationInfo.EntityData.BundleName = "cisco_ios_xr"
    lspAssociationInfo.EntityData.ParentYangName = "detail-lsp-information"
    lspAssociationInfo.EntityData.SegmentPath = "lsp-association-info"
    lspAssociationInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspAssociationInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspAssociationInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspAssociationInfo.EntityData.Children = make(map[string]types.YChild)
    lspAssociationInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    lspAssociationInfo.EntityData.Leafs["association-type"] = types.YLeaf{"AssociationType", lspAssociationInfo.AssociationType}
    lspAssociationInfo.EntityData.Leafs["association-id"] = types.YLeaf{"AssociationId", lspAssociationInfo.AssociationId}
    lspAssociationInfo.EntityData.Leafs["association-source"] = types.YLeaf{"AssociationSource", lspAssociationInfo.AssociationSource}
    return &(lspAssociationInfo.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes
// LSP attributes
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes struct {
    EntityData types.CommonEntityData
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

func (lspAttributes *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetEntityData() *types.CommonEntityData {
    lspAttributes.EntityData.YFilter = lspAttributes.YFilter
    lspAttributes.EntityData.YangName = "lsp-attributes"
    lspAttributes.EntityData.BundleName = "cisco_ios_xr"
    lspAttributes.EntityData.ParentYangName = "detail-lsp-information"
    lspAttributes.EntityData.SegmentPath = "lsp-attributes"
    lspAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspAttributes.EntityData.Children = make(map[string]types.YChild)
    lspAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    lspAttributes.EntityData.Leafs["affinity-exclude-any"] = types.YLeaf{"AffinityExcludeAny", lspAttributes.AffinityExcludeAny}
    lspAttributes.EntityData.Leafs["affinity-include-any"] = types.YLeaf{"AffinityIncludeAny", lspAttributes.AffinityIncludeAny}
    lspAttributes.EntityData.Leafs["affinity-include-all"] = types.YLeaf{"AffinityIncludeAll", lspAttributes.AffinityIncludeAll}
    lspAttributes.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", lspAttributes.SetupPriority}
    lspAttributes.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", lspAttributes.HoldPriority}
    lspAttributes.EntityData.Leafs["local-protection"] = types.YLeaf{"LocalProtection", lspAttributes.LocalProtection}
    return &(lspAttributes.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
// RRO
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RRO Type. The type is PceRro.
    RroType interface{}

    // IPv4 address of RRO. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // MPLS label of RRO. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // RRO Flags. The type is interface{} with range: 0..255.
    Flags interface{}

    // Segment Routing RRO info.
    SrRro PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
}

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetEntityData() *types.CommonEntityData {
    rro.EntityData.YFilter = rro.YFilter
    rro.EntityData.YangName = "rro"
    rro.EntityData.BundleName = "cisco_ios_xr"
    rro.EntityData.ParentYangName = "detail-lsp-information"
    rro.EntityData.SegmentPath = "rro"
    rro.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rro.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rro.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rro.EntityData.Children = make(map[string]types.YChild)
    rro.EntityData.Children["sr-rro"] = types.YChild{"SrRro", &rro.SrRro}
    rro.EntityData.Leafs = make(map[string]types.YLeaf)
    rro.EntityData.Leafs["rro-type"] = types.YLeaf{"RroType", rro.RroType}
    rro.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", rro.Ipv4Address}
    rro.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", rro.MplsLabel}
    rro.EntityData.Leafs["flags"] = types.YLeaf{"Flags", rro.Flags}
    return &(rro.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
// Segment Routing RRO info
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddr interface{}
}

func (srRro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetEntityData() *types.CommonEntityData {
    srRro.EntityData.YFilter = srRro.YFilter
    srRro.EntityData.YangName = "sr-rro"
    srRro.EntityData.BundleName = "cisco_ios_xr"
    srRro.EntityData.ParentYangName = "rro"
    srRro.EntityData.SegmentPath = "sr-rro"
    srRro.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srRro.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srRro.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srRro.EntityData.Children = make(map[string]types.YChild)
    srRro.EntityData.Leafs = make(map[string]types.YLeaf)
    srRro.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", srRro.SidType}
    srRro.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", srRro.MplsLabel}
    srRro.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", srRro.LocalAddr}
    srRro.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", srRro.RemoteAddr}
    return &(srRro.EntityData)
}

// PcePeer
// pce peer
type PcePeer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed peers database in XTC.
    PeerDetailInfos PcePeer_PeerDetailInfos

    // Peers database in XTC.
    PeerInfos PcePeer_PeerInfos
}

func (pcePeer *PcePeer) GetEntityData() *types.CommonEntityData {
    pcePeer.EntityData.YFilter = pcePeer.YFilter
    pcePeer.EntityData.YangName = "pce-peer"
    pcePeer.EntityData.BundleName = "cisco_ios_xr"
    pcePeer.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-oper"
    pcePeer.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-oper:pce-peer"
    pcePeer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcePeer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcePeer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcePeer.EntityData.Children = make(map[string]types.YChild)
    pcePeer.EntityData.Children["peer-detail-infos"] = types.YChild{"PeerDetailInfos", &pcePeer.PeerDetailInfos}
    pcePeer.EntityData.Children["peer-infos"] = types.YChild{"PeerInfos", &pcePeer.PeerInfos}
    pcePeer.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pcePeer.EntityData)
}

// PcePeer_PeerDetailInfos
// Detailed peers database in XTC
type PcePeer_PeerDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed PCE peer information. The type is slice of
    // PcePeer_PeerDetailInfos_PeerDetailInfo.
    PeerDetailInfo []PcePeer_PeerDetailInfos_PeerDetailInfo
}

func (peerDetailInfos *PcePeer_PeerDetailInfos) GetEntityData() *types.CommonEntityData {
    peerDetailInfos.EntityData.YFilter = peerDetailInfos.YFilter
    peerDetailInfos.EntityData.YangName = "peer-detail-infos"
    peerDetailInfos.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfos.EntityData.ParentYangName = "pce-peer"
    peerDetailInfos.EntityData.SegmentPath = "peer-detail-infos"
    peerDetailInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfos.EntityData.Children = make(map[string]types.YChild)
    peerDetailInfos.EntityData.Children["peer-detail-info"] = types.YChild{"PeerDetailInfo", nil}
    for i := range peerDetailInfos.PeerDetailInfo {
        peerDetailInfos.EntityData.Children[types.GetSegmentPath(&peerDetailInfos.PeerDetailInfo[i])] = types.YChild{"PeerDetailInfo", &peerDetailInfos.PeerDetailInfo[i]}
    }
    peerDetailInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerDetailInfos.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo
// Detailed PCE peer information
type PcePeer_PeerDetailInfos_PeerDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Detailed PCE protocol information.
    DetailPcepInformation PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetEntityData() *types.CommonEntityData {
    peerDetailInfo.EntityData.YFilter = peerDetailInfo.YFilter
    peerDetailInfo.EntityData.YangName = "peer-detail-info"
    peerDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfo.EntityData.ParentYangName = "peer-detail-infos"
    peerDetailInfo.EntityData.SegmentPath = "peer-detail-info" + "[peer-address='" + fmt.Sprintf("%v", peerDetailInfo.PeerAddress) + "']"
    peerDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfo.EntityData.Children = make(map[string]types.YChild)
    peerDetailInfo.EntityData.Children["detail-pcep-information"] = types.YChild{"DetailPcepInformation", &peerDetailInfo.DetailPcepInformation}
    peerDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerDetailInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", peerDetailInfo.PeerAddress}
    peerDetailInfo.EntityData.Leafs["peer-address-xr"] = types.YLeaf{"PeerAddressXr", peerDetailInfo.PeerAddressXr}
    peerDetailInfo.EntityData.Leafs["peer-protocol"] = types.YLeaf{"PeerProtocol", peerDetailInfo.PeerProtocol}
    return &(peerDetailInfo.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
// Detailed PCE protocol information
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation struct {
    EntityData types.CommonEntityData
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

func (detailPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetEntityData() *types.CommonEntityData {
    detailPcepInformation.EntityData.YFilter = detailPcepInformation.YFilter
    detailPcepInformation.EntityData.YangName = "detail-pcep-information"
    detailPcepInformation.EntityData.BundleName = "cisco_ios_xr"
    detailPcepInformation.EntityData.ParentYangName = "peer-detail-info"
    detailPcepInformation.EntityData.SegmentPath = "detail-pcep-information"
    detailPcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailPcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailPcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailPcepInformation.EntityData.Children = make(map[string]types.YChild)
    detailPcepInformation.EntityData.Children["brief-pcep-information"] = types.YChild{"BriefPcepInformation", &detailPcepInformation.BriefPcepInformation}
    detailPcepInformation.EntityData.Children["last-error-rx"] = types.YChild{"LastErrorRx", &detailPcepInformation.LastErrorRx}
    detailPcepInformation.EntityData.Children["last-error-tx"] = types.YChild{"LastErrorTx", &detailPcepInformation.LastErrorTx}
    detailPcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    detailPcepInformation.EntityData.Leafs["error"] = types.YLeaf{"Error", detailPcepInformation.Error}
    detailPcepInformation.EntityData.Leafs["speaker-id"] = types.YLeaf{"SpeakerId", detailPcepInformation.SpeakerId}
    detailPcepInformation.EntityData.Leafs["pcep-up-time"] = types.YLeaf{"PcepUpTime", detailPcepInformation.PcepUpTime}
    detailPcepInformation.EntityData.Leafs["keepalives"] = types.YLeaf{"Keepalives", detailPcepInformation.Keepalives}
    detailPcepInformation.EntityData.Leafs["md5-enabled"] = types.YLeaf{"Md5Enabled", detailPcepInformation.Md5Enabled}
    detailPcepInformation.EntityData.Leafs["keychain-enabled"] = types.YLeaf{"KeychainEnabled", detailPcepInformation.KeychainEnabled}
    detailPcepInformation.EntityData.Leafs["negotiated-local-keepalive"] = types.YLeaf{"NegotiatedLocalKeepalive", detailPcepInformation.NegotiatedLocalKeepalive}
    detailPcepInformation.EntityData.Leafs["negotiated-remote-keepalive"] = types.YLeaf{"NegotiatedRemoteKeepalive", detailPcepInformation.NegotiatedRemoteKeepalive}
    detailPcepInformation.EntityData.Leafs["negotiated-dead-time"] = types.YLeaf{"NegotiatedDeadTime", detailPcepInformation.NegotiatedDeadTime}
    detailPcepInformation.EntityData.Leafs["pce-request-rx"] = types.YLeaf{"PceRequestRx", detailPcepInformation.PceRequestRx}
    detailPcepInformation.EntityData.Leafs["pce-request-tx"] = types.YLeaf{"PceRequestTx", detailPcepInformation.PceRequestTx}
    detailPcepInformation.EntityData.Leafs["pce-reply-rx"] = types.YLeaf{"PceReplyRx", detailPcepInformation.PceReplyRx}
    detailPcepInformation.EntityData.Leafs["pce-reply-tx"] = types.YLeaf{"PceReplyTx", detailPcepInformation.PceReplyTx}
    detailPcepInformation.EntityData.Leafs["pce-error-rx"] = types.YLeaf{"PceErrorRx", detailPcepInformation.PceErrorRx}
    detailPcepInformation.EntityData.Leafs["pce-error-tx"] = types.YLeaf{"PceErrorTx", detailPcepInformation.PceErrorTx}
    detailPcepInformation.EntityData.Leafs["pce-open-tx"] = types.YLeaf{"PceOpenTx", detailPcepInformation.PceOpenTx}
    detailPcepInformation.EntityData.Leafs["pce-open-rx"] = types.YLeaf{"PceOpenRx", detailPcepInformation.PceOpenRx}
    detailPcepInformation.EntityData.Leafs["pce-report-rx"] = types.YLeaf{"PceReportRx", detailPcepInformation.PceReportRx}
    detailPcepInformation.EntityData.Leafs["pce-report-tx"] = types.YLeaf{"PceReportTx", detailPcepInformation.PceReportTx}
    detailPcepInformation.EntityData.Leafs["pce-update-rx"] = types.YLeaf{"PceUpdateRx", detailPcepInformation.PceUpdateRx}
    detailPcepInformation.EntityData.Leafs["pce-update-tx"] = types.YLeaf{"PceUpdateTx", detailPcepInformation.PceUpdateTx}
    detailPcepInformation.EntityData.Leafs["pce-initiate-rx"] = types.YLeaf{"PceInitiateRx", detailPcepInformation.PceInitiateRx}
    detailPcepInformation.EntityData.Leafs["pce-initiate-tx"] = types.YLeaf{"PceInitiateTx", detailPcepInformation.PceInitiateTx}
    detailPcepInformation.EntityData.Leafs["pce-keepalive-tx"] = types.YLeaf{"PceKeepaliveTx", detailPcepInformation.PceKeepaliveTx}
    detailPcepInformation.EntityData.Leafs["pce-keepalive-rx"] = types.YLeaf{"PceKeepaliveRx", detailPcepInformation.PceKeepaliveRx}
    detailPcepInformation.EntityData.Leafs["local-session-id"] = types.YLeaf{"LocalSessionId", detailPcepInformation.LocalSessionId}
    detailPcepInformation.EntityData.Leafs["remote-session-id"] = types.YLeaf{"RemoteSessionId", detailPcepInformation.RemoteSessionId}
    detailPcepInformation.EntityData.Leafs["minimum-keepalive-interval"] = types.YLeaf{"MinimumKeepaliveInterval", detailPcepInformation.MinimumKeepaliveInterval}
    detailPcepInformation.EntityData.Leafs["maximum-dead-interval"] = types.YLeaf{"MaximumDeadInterval", detailPcepInformation.MaximumDeadInterval}
    return &(detailPcepInformation.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation
// Brief PCE protocol information
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation struct {
    EntityData types.CommonEntityData
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

func (briefPcepInformation *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetEntityData() *types.CommonEntityData {
    briefPcepInformation.EntityData.YFilter = briefPcepInformation.YFilter
    briefPcepInformation.EntityData.YangName = "brief-pcep-information"
    briefPcepInformation.EntityData.BundleName = "cisco_ios_xr"
    briefPcepInformation.EntityData.ParentYangName = "detail-pcep-information"
    briefPcepInformation.EntityData.SegmentPath = "brief-pcep-information"
    briefPcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefPcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefPcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefPcepInformation.EntityData.Children = make(map[string]types.YChild)
    briefPcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefPcepInformation.EntityData.Leafs["pcep-state"] = types.YLeaf{"PcepState", briefPcepInformation.PcepState}
    briefPcepInformation.EntityData.Leafs["stateful"] = types.YLeaf{"Stateful", briefPcepInformation.Stateful}
    briefPcepInformation.EntityData.Leafs["capability-update"] = types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate}
    briefPcepInformation.EntityData.Leafs["capability-instantiate"] = types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate}
    briefPcepInformation.EntityData.Leafs["capability-segment-routing"] = types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting}
    briefPcepInformation.EntityData.Leafs["capability-triggered-sync"] = types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync}
    briefPcepInformation.EntityData.Leafs["capability-db-version"] = types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion}
    briefPcepInformation.EntityData.Leafs["capability-delta-sync"] = types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync}
    return &(briefPcepInformation.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx
// Last PCError received
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorRx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetEntityData() *types.CommonEntityData {
    lastErrorRx.EntityData.YFilter = lastErrorRx.YFilter
    lastErrorRx.EntityData.YangName = "last-error-rx"
    lastErrorRx.EntityData.BundleName = "cisco_ios_xr"
    lastErrorRx.EntityData.ParentYangName = "detail-pcep-information"
    lastErrorRx.EntityData.SegmentPath = "last-error-rx"
    lastErrorRx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErrorRx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErrorRx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErrorRx.EntityData.Children = make(map[string]types.YChild)
    lastErrorRx.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErrorRx.EntityData.Leafs["pc-error-type"] = types.YLeaf{"PcErrorType", lastErrorRx.PcErrorType}
    lastErrorRx.EntityData.Leafs["pc-error-value"] = types.YLeaf{"PcErrorValue", lastErrorRx.PcErrorValue}
    return &(lastErrorRx.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx
// Last PCError sent
type PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorTx *PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetEntityData() *types.CommonEntityData {
    lastErrorTx.EntityData.YFilter = lastErrorTx.YFilter
    lastErrorTx.EntityData.YangName = "last-error-tx"
    lastErrorTx.EntityData.BundleName = "cisco_ios_xr"
    lastErrorTx.EntityData.ParentYangName = "detail-pcep-information"
    lastErrorTx.EntityData.SegmentPath = "last-error-tx"
    lastErrorTx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErrorTx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErrorTx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErrorTx.EntityData.Children = make(map[string]types.YChild)
    lastErrorTx.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErrorTx.EntityData.Leafs["pc-error-type"] = types.YLeaf{"PcErrorType", lastErrorTx.PcErrorType}
    lastErrorTx.EntityData.Leafs["pc-error-value"] = types.YLeaf{"PcErrorValue", lastErrorTx.PcErrorValue}
    return &(lastErrorTx.EntityData)
}

// PcePeer_PeerInfos
// Peers database in XTC
type PcePeer_PeerInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer information. The type is slice of PcePeer_PeerInfos_PeerInfo.
    PeerInfo []PcePeer_PeerInfos_PeerInfo
}

func (peerInfos *PcePeer_PeerInfos) GetEntityData() *types.CommonEntityData {
    peerInfos.EntityData.YFilter = peerInfos.YFilter
    peerInfos.EntityData.YangName = "peer-infos"
    peerInfos.EntityData.BundleName = "cisco_ios_xr"
    peerInfos.EntityData.ParentYangName = "pce-peer"
    peerInfos.EntityData.SegmentPath = "peer-infos"
    peerInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfos.EntityData.Children = make(map[string]types.YChild)
    peerInfos.EntityData.Children["peer-info"] = types.YChild{"PeerInfo", nil}
    for i := range peerInfos.PeerInfo {
        peerInfos.EntityData.Children[types.GetSegmentPath(&peerInfos.PeerInfo[i])] = types.YChild{"PeerInfo", &peerInfos.PeerInfo[i]}
    }
    peerInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerInfos.EntityData)
}

// PcePeer_PeerInfos_PeerInfo
// PCE peer information
type PcePeer_PeerInfos_PeerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // PCE protocol information.
    BriefPcepInformation PcePeer_PeerInfos_PeerInfo_BriefPcepInformation
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetEntityData() *types.CommonEntityData {
    peerInfo.EntityData.YFilter = peerInfo.YFilter
    peerInfo.EntityData.YangName = "peer-info"
    peerInfo.EntityData.BundleName = "cisco_ios_xr"
    peerInfo.EntityData.ParentYangName = "peer-infos"
    peerInfo.EntityData.SegmentPath = "peer-info" + "[peer-address='" + fmt.Sprintf("%v", peerInfo.PeerAddress) + "']"
    peerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfo.EntityData.Children = make(map[string]types.YChild)
    peerInfo.EntityData.Children["brief-pcep-information"] = types.YChild{"BriefPcepInformation", &peerInfo.BriefPcepInformation}
    peerInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", peerInfo.PeerAddress}
    peerInfo.EntityData.Leafs["peer-address-xr"] = types.YLeaf{"PeerAddressXr", peerInfo.PeerAddressXr}
    peerInfo.EntityData.Leafs["peer-protocol"] = types.YLeaf{"PeerProtocol", peerInfo.PeerProtocol}
    return &(peerInfo.EntityData)
}

// PcePeer_PeerInfos_PeerInfo_BriefPcepInformation
// PCE protocol information
type PcePeer_PeerInfos_PeerInfo_BriefPcepInformation struct {
    EntityData types.CommonEntityData
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

func (briefPcepInformation *PcePeer_PeerInfos_PeerInfo_BriefPcepInformation) GetEntityData() *types.CommonEntityData {
    briefPcepInformation.EntityData.YFilter = briefPcepInformation.YFilter
    briefPcepInformation.EntityData.YangName = "brief-pcep-information"
    briefPcepInformation.EntityData.BundleName = "cisco_ios_xr"
    briefPcepInformation.EntityData.ParentYangName = "peer-info"
    briefPcepInformation.EntityData.SegmentPath = "brief-pcep-information"
    briefPcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefPcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefPcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefPcepInformation.EntityData.Children = make(map[string]types.YChild)
    briefPcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefPcepInformation.EntityData.Leafs["pcep-state"] = types.YLeaf{"PcepState", briefPcepInformation.PcepState}
    briefPcepInformation.EntityData.Leafs["stateful"] = types.YLeaf{"Stateful", briefPcepInformation.Stateful}
    briefPcepInformation.EntityData.Leafs["capability-update"] = types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate}
    briefPcepInformation.EntityData.Leafs["capability-instantiate"] = types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate}
    briefPcepInformation.EntityData.Leafs["capability-segment-routing"] = types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting}
    briefPcepInformation.EntityData.Leafs["capability-triggered-sync"] = types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync}
    briefPcepInformation.EntityData.Leafs["capability-db-version"] = types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion}
    briefPcepInformation.EntityData.Leafs["capability-delta-sync"] = types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync}
    return &(briefPcepInformation.EntityData)
}

// PceTopology
// pce topology
type PceTopology struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node summary database in XTC.
    TopologySummary PceTopology_TopologySummary

    // Node database in XTC.
    TopologyNodes PceTopology_TopologyNodes

    // Prefixes database in XTC.
    PrefixInfos PceTopology_PrefixInfos
}

func (pceTopology *PceTopology) GetEntityData() *types.CommonEntityData {
    pceTopology.EntityData.YFilter = pceTopology.YFilter
    pceTopology.EntityData.YangName = "pce-topology"
    pceTopology.EntityData.BundleName = "cisco_ios_xr"
    pceTopology.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-oper"
    pceTopology.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-oper:pce-topology"
    pceTopology.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pceTopology.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pceTopology.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pceTopology.EntityData.Children = make(map[string]types.YChild)
    pceTopology.EntityData.Children["topology-summary"] = types.YChild{"TopologySummary", &pceTopology.TopologySummary}
    pceTopology.EntityData.Children["topology-nodes"] = types.YChild{"TopologyNodes", &pceTopology.TopologyNodes}
    pceTopology.EntityData.Children["prefix-infos"] = types.YChild{"PrefixInfos", &pceTopology.PrefixInfos}
    pceTopology.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pceTopology.EntityData)
}

// PceTopology_TopologySummary
// Node summary database in XTC
type PceTopology_TopologySummary struct {
    EntityData types.CommonEntityData
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

func (topologySummary *PceTopology_TopologySummary) GetEntityData() *types.CommonEntityData {
    topologySummary.EntityData.YFilter = topologySummary.YFilter
    topologySummary.EntityData.YangName = "topology-summary"
    topologySummary.EntityData.BundleName = "cisco_ios_xr"
    topologySummary.EntityData.ParentYangName = "pce-topology"
    topologySummary.EntityData.SegmentPath = "topology-summary"
    topologySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologySummary.EntityData.Children = make(map[string]types.YChild)
    topologySummary.EntityData.Children["stats-topology-update"] = types.YChild{"StatsTopologyUpdate", &topologySummary.StatsTopologyUpdate}
    topologySummary.EntityData.Leafs = make(map[string]types.YLeaf)
    topologySummary.EntityData.Leafs["nodes"] = types.YLeaf{"Nodes", topologySummary.Nodes}
    topologySummary.EntityData.Leafs["lookup-nodes"] = types.YLeaf{"LookupNodes", topologySummary.LookupNodes}
    topologySummary.EntityData.Leafs["prefixes"] = types.YLeaf{"Prefixes", topologySummary.Prefixes}
    topologySummary.EntityData.Leafs["prefix-sids"] = types.YLeaf{"PrefixSids", topologySummary.PrefixSids}
    topologySummary.EntityData.Leafs["regular-prefix-sids"] = types.YLeaf{"RegularPrefixSids", topologySummary.RegularPrefixSids}
    topologySummary.EntityData.Leafs["strict-prefix-sids"] = types.YLeaf{"StrictPrefixSids", topologySummary.StrictPrefixSids}
    topologySummary.EntityData.Leafs["links"] = types.YLeaf{"Links", topologySummary.Links}
    topologySummary.EntityData.Leafs["epe-links"] = types.YLeaf{"EpeLinks", topologySummary.EpeLinks}
    topologySummary.EntityData.Leafs["adjacency-sids"] = types.YLeaf{"AdjacencySids", topologySummary.AdjacencySids}
    topologySummary.EntityData.Leafs["epesids"] = types.YLeaf{"Epesids", topologySummary.Epesids}
    topologySummary.EntityData.Leafs["protected-adjacency-sids"] = types.YLeaf{"ProtectedAdjacencySids", topologySummary.ProtectedAdjacencySids}
    topologySummary.EntityData.Leafs["un-protected-adjacency-sids"] = types.YLeaf{"UnProtectedAdjacencySids", topologySummary.UnProtectedAdjacencySids}
    topologySummary.EntityData.Leafs["topology-consistent"] = types.YLeaf{"TopologyConsistent", topologySummary.TopologyConsistent}
    return &(topologySummary.EntityData)
}

// PceTopology_TopologySummary_StatsTopologyUpdate
// Statistics on topology update
type PceTopology_TopologySummary_StatsTopologyUpdate struct {
    EntityData types.CommonEntityData
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

func (statsTopologyUpdate *PceTopology_TopologySummary_StatsTopologyUpdate) GetEntityData() *types.CommonEntityData {
    statsTopologyUpdate.EntityData.YFilter = statsTopologyUpdate.YFilter
    statsTopologyUpdate.EntityData.YangName = "stats-topology-update"
    statsTopologyUpdate.EntityData.BundleName = "cisco_ios_xr"
    statsTopologyUpdate.EntityData.ParentYangName = "topology-summary"
    statsTopologyUpdate.EntityData.SegmentPath = "stats-topology-update"
    statsTopologyUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsTopologyUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsTopologyUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsTopologyUpdate.EntityData.Children = make(map[string]types.YChild)
    statsTopologyUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    statsTopologyUpdate.EntityData.Leafs["num-nodes-added"] = types.YLeaf{"NumNodesAdded", statsTopologyUpdate.NumNodesAdded}
    statsTopologyUpdate.EntityData.Leafs["num-nodes-deleted"] = types.YLeaf{"NumNodesDeleted", statsTopologyUpdate.NumNodesDeleted}
    statsTopologyUpdate.EntityData.Leafs["num-links-added"] = types.YLeaf{"NumLinksAdded", statsTopologyUpdate.NumLinksAdded}
    statsTopologyUpdate.EntityData.Leafs["num-links-deleted"] = types.YLeaf{"NumLinksDeleted", statsTopologyUpdate.NumLinksDeleted}
    statsTopologyUpdate.EntityData.Leafs["num-prefixes-added"] = types.YLeaf{"NumPrefixesAdded", statsTopologyUpdate.NumPrefixesAdded}
    statsTopologyUpdate.EntityData.Leafs["num-prefixes-deleted"] = types.YLeaf{"NumPrefixesDeleted", statsTopologyUpdate.NumPrefixesDeleted}
    return &(statsTopologyUpdate.EntityData)
}

// PceTopology_TopologyNodes
// Node database in XTC
type PceTopology_TopologyNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode.
    TopologyNode []PceTopology_TopologyNodes_TopologyNode
}

func (topologyNodes *PceTopology_TopologyNodes) GetEntityData() *types.CommonEntityData {
    topologyNodes.EntityData.YFilter = topologyNodes.YFilter
    topologyNodes.EntityData.YangName = "topology-nodes"
    topologyNodes.EntityData.BundleName = "cisco_ios_xr"
    topologyNodes.EntityData.ParentYangName = "pce-topology"
    topologyNodes.EntityData.SegmentPath = "topology-nodes"
    topologyNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNodes.EntityData.Children = make(map[string]types.YChild)
    topologyNodes.EntityData.Children["topology-node"] = types.YChild{"TopologyNode", nil}
    for i := range topologyNodes.TopologyNode {
        topologyNodes.EntityData.Children[types.GetSegmentPath(&topologyNodes.TopologyNode[i])] = types.YChild{"TopologyNode", &topologyNodes.TopologyNode[i]}
    }
    topologyNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(topologyNodes.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode
// Node information
type PceTopology_TopologyNodes_TopologyNode struct {
    EntityData types.CommonEntityData
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

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetEntityData() *types.CommonEntityData {
    topologyNode.EntityData.YFilter = topologyNode.YFilter
    topologyNode.EntityData.YangName = "topology-node"
    topologyNode.EntityData.BundleName = "cisco_ios_xr"
    topologyNode.EntityData.ParentYangName = "topology-nodes"
    topologyNode.EntityData.SegmentPath = "topology-node" + "[node-identifier='" + fmt.Sprintf("%v", topologyNode.NodeIdentifier) + "']"
    topologyNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNode.EntityData.Children = make(map[string]types.YChild)
    topologyNode.EntityData.Children["node-protocol-identifier"] = types.YChild{"NodeProtocolIdentifier", &topologyNode.NodeProtocolIdentifier}
    topologyNode.EntityData.Children["prefix-sid"] = types.YChild{"PrefixSid", nil}
    for i := range topologyNode.PrefixSid {
        topologyNode.EntityData.Children[types.GetSegmentPath(&topologyNode.PrefixSid[i])] = types.YChild{"PrefixSid", &topologyNode.PrefixSid[i]}
    }
    topologyNode.EntityData.Children["ipv4-link"] = types.YChild{"Ipv4Link", nil}
    for i := range topologyNode.Ipv4Link {
        topologyNode.EntityData.Children[types.GetSegmentPath(&topologyNode.Ipv4Link[i])] = types.YChild{"Ipv4Link", &topologyNode.Ipv4Link[i]}
    }
    topologyNode.EntityData.Children["ipv6-link"] = types.YChild{"Ipv6Link", nil}
    for i := range topologyNode.Ipv6Link {
        topologyNode.EntityData.Children[types.GetSegmentPath(&topologyNode.Ipv6Link[i])] = types.YChild{"Ipv6Link", &topologyNode.Ipv6Link[i]}
    }
    topologyNode.EntityData.Leafs = make(map[string]types.YLeaf)
    topologyNode.EntityData.Leafs["node-identifier"] = types.YLeaf{"NodeIdentifier", topologyNode.NodeIdentifier}
    topologyNode.EntityData.Leafs["node-identifier-xr"] = types.YLeaf{"NodeIdentifierXr", topologyNode.NodeIdentifierXr}
    topologyNode.EntityData.Leafs["overload"] = types.YLeaf{"Overload", topologyNode.Overload}
    return &(topologyNode.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier
// Node protocol identifier
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    nodeProtocolIdentifier.EntityData.YFilter = nodeProtocolIdentifier.YFilter
    nodeProtocolIdentifier.EntityData.YangName = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    nodeProtocolIdentifier.EntityData.ParentYangName = "topology-node"
    nodeProtocolIdentifier.EntityData.SegmentPath = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    nodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &nodeProtocolIdentifier.IgpInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &nodeProtocolIdentifier.SrgbInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeProtocolIdentifier.NodeName}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", nodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", nodeProtocolIdentifier.Ipv4BgpRouterId}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", nodeProtocolIdentifier.Ipv4TeRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", nodeProtocolIdentifier.Ipv4TeRouterId}
    return &(nodeProtocolIdentifier.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_PrefixSid
// Prefix SIDs
type PceTopology_TopologyNodes_TopologyNode_PrefixSid struct {
    EntityData types.CommonEntityData
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

func (prefixSid *PceTopology_TopologyNodes_TopologyNode_PrefixSid) GetEntityData() *types.CommonEntityData {
    prefixSid.EntityData.YFilter = prefixSid.YFilter
    prefixSid.EntityData.YangName = "prefix-sid"
    prefixSid.EntityData.BundleName = "cisco_ios_xr"
    prefixSid.EntityData.ParentYangName = "topology-node"
    prefixSid.EntityData.SegmentPath = "prefix-sid"
    prefixSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSid.EntityData.Children = make(map[string]types.YChild)
    prefixSid.EntityData.Children["sid-prefix"] = types.YChild{"SidPrefix", &prefixSid.SidPrefix}
    prefixSid.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", prefixSid.SidType}
    prefixSid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", prefixSid.MplsLabel}
    prefixSid.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", prefixSid.DomainIdentifier}
    prefixSid.EntityData.Leafs["rflag"] = types.YLeaf{"Rflag", prefixSid.Rflag}
    prefixSid.EntityData.Leafs["nflag"] = types.YLeaf{"Nflag", prefixSid.Nflag}
    prefixSid.EntityData.Leafs["pflag"] = types.YLeaf{"Pflag", prefixSid.Pflag}
    prefixSid.EntityData.Leafs["eflag"] = types.YLeaf{"Eflag", prefixSid.Eflag}
    prefixSid.EntityData.Leafs["vflag"] = types.YLeaf{"Vflag", prefixSid.Vflag}
    prefixSid.EntityData.Leafs["lflag"] = types.YLeaf{"Lflag", prefixSid.Lflag}
    return &(prefixSid.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "prefix-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = make(map[string]types.YChild)
    sidPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    sidPrefix.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", sidPrefix.AfName}
    sidPrefix.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", sidPrefix.Ipv4}
    sidPrefix.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", sidPrefix.Ipv6}
    return &(sidPrefix.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIpv4Address interface{}

    // Remote IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipv4Link *PceTopology_TopologyNodes_TopologyNode_Ipv4Link) GetEntityData() *types.CommonEntityData {
    ipv4Link.EntityData.YFilter = ipv4Link.YFilter
    ipv4Link.EntityData.YangName = "ipv4-link"
    ipv4Link.EntityData.BundleName = "cisco_ios_xr"
    ipv4Link.EntityData.ParentYangName = "topology-node"
    ipv4Link.EntityData.SegmentPath = "ipv4-link"
    ipv4Link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Link.EntityData.Children = make(map[string]types.YChild)
    ipv4Link.EntityData.Children["local-igp-information"] = types.YChild{"LocalIgpInformation", &ipv4Link.LocalIgpInformation}
    ipv4Link.EntityData.Children["remote-node-protocol-identifier"] = types.YChild{"RemoteNodeProtocolIdentifier", &ipv4Link.RemoteNodeProtocolIdentifier}
    ipv4Link.EntityData.Children["adjacency-sid"] = types.YChild{"AdjacencySid", nil}
    for i := range ipv4Link.AdjacencySid {
        ipv4Link.EntityData.Children[types.GetSegmentPath(&ipv4Link.AdjacencySid[i])] = types.YChild{"AdjacencySid", &ipv4Link.AdjacencySid[i]}
    }
    ipv4Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Link.EntityData.Leafs["local-ipv4-address"] = types.YLeaf{"LocalIpv4Address", ipv4Link.LocalIpv4Address}
    ipv4Link.EntityData.Leafs["remote-ipv4-address"] = types.YLeaf{"RemoteIpv4Address", ipv4Link.RemoteIpv4Address}
    ipv4Link.EntityData.Leafs["igp-metric"] = types.YLeaf{"IgpMetric", ipv4Link.IgpMetric}
    ipv4Link.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", ipv4Link.TeMetric}
    ipv4Link.EntityData.Leafs["maximum-link-bandwidth"] = types.YLeaf{"MaximumLinkBandwidth", ipv4Link.MaximumLinkBandwidth}
    ipv4Link.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", ipv4Link.MaxReservableBandwidth}
    ipv4Link.EntityData.Leafs["srlgs"] = types.YLeaf{"Srlgs", ipv4Link.Srlgs}
    return &(ipv4Link.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
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

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetEntityData() *types.CommonEntityData {
    localIgpInformation.EntityData.YFilter = localIgpInformation.YFilter
    localIgpInformation.EntityData.YangName = "local-igp-information"
    localIgpInformation.EntityData.BundleName = "cisco_ios_xr"
    localIgpInformation.EntityData.ParentYangName = "ipv4-link"
    localIgpInformation.EntityData.SegmentPath = "local-igp-information"
    localIgpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIgpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIgpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIgpInformation.EntityData.Children = make(map[string]types.YChild)
    localIgpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &localIgpInformation.Igp}
    localIgpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    localIgpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier}
    localIgpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", localIgpInformation.AutonomousSystemNumber}
    return &(localIgpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "local-igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    remoteNodeProtocolIdentifier.EntityData.YFilter = remoteNodeProtocolIdentifier.YFilter
    remoteNodeProtocolIdentifier.EntityData.YangName = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteNodeProtocolIdentifier.EntityData.ParentYangName = "ipv4-link"
    remoteNodeProtocolIdentifier.EntityData.SegmentPath = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteNodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &remoteNodeProtocolIdentifier.IgpInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &remoteNodeProtocolIdentifier.SrgbInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteNodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", remoteNodeProtocolIdentifier.NodeName}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", remoteNodeProtocolIdentifier.Ipv4BgpRouterId}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", remoteNodeProtocolIdentifier.Ipv4TeRouterId}
    return &(remoteNodeProtocolIdentifier.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
// Adjacency SIDs
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid struct {
    EntityData types.CommonEntityData
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

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "ipv4-link"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = make(map[string]types.YChild)
    adjacencySid.EntityData.Children["sid-prefix"] = types.YChild{"SidPrefix", &adjacencySid.SidPrefix}
    adjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    adjacencySid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", adjacencySid.SidType}
    adjacencySid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", adjacencySid.MplsLabel}
    adjacencySid.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", adjacencySid.DomainIdentifier}
    adjacencySid.EntityData.Leafs["rflag"] = types.YLeaf{"Rflag", adjacencySid.Rflag}
    adjacencySid.EntityData.Leafs["nflag"] = types.YLeaf{"Nflag", adjacencySid.Nflag}
    adjacencySid.EntityData.Leafs["pflag"] = types.YLeaf{"Pflag", adjacencySid.Pflag}
    adjacencySid.EntityData.Leafs["eflag"] = types.YLeaf{"Eflag", adjacencySid.Eflag}
    adjacencySid.EntityData.Leafs["vflag"] = types.YLeaf{"Vflag", adjacencySid.Vflag}
    adjacencySid.EntityData.Leafs["lflag"] = types.YLeaf{"Lflag", adjacencySid.Lflag}
    return &(adjacencySid.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "adjacency-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = make(map[string]types.YChild)
    sidPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    sidPrefix.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", sidPrefix.AfName}
    sidPrefix.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", sidPrefix.Ipv4}
    sidPrefix.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", sidPrefix.Ipv6}
    return &(sidPrefix.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalIpv6Address interface{}

    // Remote IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (ipv6Link *PceTopology_TopologyNodes_TopologyNode_Ipv6Link) GetEntityData() *types.CommonEntityData {
    ipv6Link.EntityData.YFilter = ipv6Link.YFilter
    ipv6Link.EntityData.YangName = "ipv6-link"
    ipv6Link.EntityData.BundleName = "cisco_ios_xr"
    ipv6Link.EntityData.ParentYangName = "topology-node"
    ipv6Link.EntityData.SegmentPath = "ipv6-link"
    ipv6Link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Link.EntityData.Children = make(map[string]types.YChild)
    ipv6Link.EntityData.Children["local-igp-information"] = types.YChild{"LocalIgpInformation", &ipv6Link.LocalIgpInformation}
    ipv6Link.EntityData.Children["remote-node-protocol-identifier"] = types.YChild{"RemoteNodeProtocolIdentifier", &ipv6Link.RemoteNodeProtocolIdentifier}
    ipv6Link.EntityData.Children["adjacency-sid"] = types.YChild{"AdjacencySid", nil}
    for i := range ipv6Link.AdjacencySid {
        ipv6Link.EntityData.Children[types.GetSegmentPath(&ipv6Link.AdjacencySid[i])] = types.YChild{"AdjacencySid", &ipv6Link.AdjacencySid[i]}
    }
    ipv6Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Link.EntityData.Leafs["local-ipv6-address"] = types.YLeaf{"LocalIpv6Address", ipv6Link.LocalIpv6Address}
    ipv6Link.EntityData.Leafs["remote-ipv6-address"] = types.YLeaf{"RemoteIpv6Address", ipv6Link.RemoteIpv6Address}
    ipv6Link.EntityData.Leafs["igp-metric"] = types.YLeaf{"IgpMetric", ipv6Link.IgpMetric}
    ipv6Link.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", ipv6Link.TeMetric}
    ipv6Link.EntityData.Leafs["maximum-link-bandwidth"] = types.YLeaf{"MaximumLinkBandwidth", ipv6Link.MaximumLinkBandwidth}
    ipv6Link.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", ipv6Link.MaxReservableBandwidth}
    return &(ipv6Link.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
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

func (localIgpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetEntityData() *types.CommonEntityData {
    localIgpInformation.EntityData.YFilter = localIgpInformation.YFilter
    localIgpInformation.EntityData.YangName = "local-igp-information"
    localIgpInformation.EntityData.BundleName = "cisco_ios_xr"
    localIgpInformation.EntityData.ParentYangName = "ipv6-link"
    localIgpInformation.EntityData.SegmentPath = "local-igp-information"
    localIgpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIgpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIgpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIgpInformation.EntityData.Children = make(map[string]types.YChild)
    localIgpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &localIgpInformation.Igp}
    localIgpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    localIgpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier}
    localIgpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", localIgpInformation.AutonomousSystemNumber}
    return &(localIgpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "local-igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    remoteNodeProtocolIdentifier.EntityData.YFilter = remoteNodeProtocolIdentifier.YFilter
    remoteNodeProtocolIdentifier.EntityData.YangName = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteNodeProtocolIdentifier.EntityData.ParentYangName = "ipv6-link"
    remoteNodeProtocolIdentifier.EntityData.SegmentPath = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteNodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &remoteNodeProtocolIdentifier.IgpInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &remoteNodeProtocolIdentifier.SrgbInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteNodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", remoteNodeProtocolIdentifier.NodeName}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", remoteNodeProtocolIdentifier.Ipv4BgpRouterId}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", remoteNodeProtocolIdentifier.Ipv4TeRouterId}
    return &(remoteNodeProtocolIdentifier.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
// Adjacency SIDs
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid struct {
    EntityData types.CommonEntityData
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

func (adjacencySid *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "ipv6-link"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = make(map[string]types.YChild)
    adjacencySid.EntityData.Children["sid-prefix"] = types.YChild{"SidPrefix", &adjacencySid.SidPrefix}
    adjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    adjacencySid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", adjacencySid.SidType}
    adjacencySid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", adjacencySid.MplsLabel}
    adjacencySid.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", adjacencySid.DomainIdentifier}
    adjacencySid.EntityData.Leafs["rflag"] = types.YLeaf{"Rflag", adjacencySid.Rflag}
    adjacencySid.EntityData.Leafs["nflag"] = types.YLeaf{"Nflag", adjacencySid.Nflag}
    adjacencySid.EntityData.Leafs["pflag"] = types.YLeaf{"Pflag", adjacencySid.Pflag}
    adjacencySid.EntityData.Leafs["eflag"] = types.YLeaf{"Eflag", adjacencySid.Eflag}
    adjacencySid.EntityData.Leafs["vflag"] = types.YLeaf{"Vflag", adjacencySid.Vflag}
    adjacencySid.EntityData.Leafs["lflag"] = types.YLeaf{"Lflag", adjacencySid.Lflag}
    return &(adjacencySid.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "adjacency-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = make(map[string]types.YChild)
    sidPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    sidPrefix.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", sidPrefix.AfName}
    sidPrefix.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", sidPrefix.Ipv4}
    sidPrefix.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", sidPrefix.Ipv6}
    return &(sidPrefix.EntityData)
}

// PceTopology_PrefixInfos
// Prefixes database in XTC
type PceTopology_PrefixInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE prefix information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo.
    PrefixInfo []PceTopology_PrefixInfos_PrefixInfo
}

func (prefixInfos *PceTopology_PrefixInfos) GetEntityData() *types.CommonEntityData {
    prefixInfos.EntityData.YFilter = prefixInfos.YFilter
    prefixInfos.EntityData.YangName = "prefix-infos"
    prefixInfos.EntityData.BundleName = "cisco_ios_xr"
    prefixInfos.EntityData.ParentYangName = "pce-topology"
    prefixInfos.EntityData.SegmentPath = "prefix-infos"
    prefixInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixInfos.EntityData.Children = make(map[string]types.YChild)
    prefixInfos.EntityData.Children["prefix-info"] = types.YChild{"PrefixInfo", nil}
    for i := range prefixInfos.PrefixInfo {
        prefixInfos.EntityData.Children[types.GetSegmentPath(&prefixInfos.PrefixInfo[i])] = types.YChild{"PrefixInfo", &prefixInfos.PrefixInfo[i]}
    }
    prefixInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixInfos.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo
// PCE prefix information
type PceTopology_PrefixInfos_PrefixInfo struct {
    EntityData types.CommonEntityData
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

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetEntityData() *types.CommonEntityData {
    prefixInfo.EntityData.YFilter = prefixInfo.YFilter
    prefixInfo.EntityData.YangName = "prefix-info"
    prefixInfo.EntityData.BundleName = "cisco_ios_xr"
    prefixInfo.EntityData.ParentYangName = "prefix-infos"
    prefixInfo.EntityData.SegmentPath = "prefix-info" + "[node-identifier='" + fmt.Sprintf("%v", prefixInfo.NodeIdentifier) + "']"
    prefixInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixInfo.EntityData.Children = make(map[string]types.YChild)
    prefixInfo.EntityData.Children["node-protocol-identifier"] = types.YChild{"NodeProtocolIdentifier", &prefixInfo.NodeProtocolIdentifier}
    prefixInfo.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range prefixInfo.Address {
        prefixInfo.EntityData.Children[types.GetSegmentPath(&prefixInfo.Address[i])] = types.YChild{"Address", &prefixInfo.Address[i]}
    }
    prefixInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixInfo.EntityData.Leafs["node-identifier"] = types.YLeaf{"NodeIdentifier", prefixInfo.NodeIdentifier}
    prefixInfo.EntityData.Leafs["node-identifier-xr"] = types.YLeaf{"NodeIdentifierXr", prefixInfo.NodeIdentifierXr}
    return &(prefixInfo.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier
// Node protocol identifier
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    nodeProtocolIdentifier.EntityData.YFilter = nodeProtocolIdentifier.YFilter
    nodeProtocolIdentifier.EntityData.YangName = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    nodeProtocolIdentifier.EntityData.ParentYangName = "prefix-info"
    nodeProtocolIdentifier.EntityData.SegmentPath = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    nodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &nodeProtocolIdentifier.IgpInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &nodeProtocolIdentifier.SrgbInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeProtocolIdentifier.NodeName}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", nodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", nodeProtocolIdentifier.Ipv4BgpRouterId}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", nodeProtocolIdentifier.Ipv4TeRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", nodeProtocolIdentifier.Ipv4TeRouterId}
    return &(nodeProtocolIdentifier.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_Address
// Prefix address
type PceTopology_PrefixInfos_PrefixInfo_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix IP address.
    Ip PceTopology_PrefixInfos_PrefixInfo_Address_Ip
}

func (address *PceTopology_PrefixInfos_PrefixInfo_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "prefix-info"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["ip"] = types.YChild{"Ip", &address.Ip}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(address.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_Address_Ip
// Prefix IP address
type PceTopology_PrefixInfos_PrefixInfo_Address_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (ip *PceTopology_PrefixInfos_PrefixInfo_Address_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "address"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    ip.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", ip.AfName}
    ip.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", ip.Ipv4}
    ip.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", ip.Ipv6}
    return &(ip.EntityData)
}

// Pce
// pce
type Pce struct {
    EntityData types.CommonEntityData
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

func (pce *Pce) GetEntityData() *types.CommonEntityData {
    pce.EntityData.YFilter = pce.YFilter
    pce.EntityData.YangName = "pce"
    pce.EntityData.BundleName = "cisco_ios_xr"
    pce.EntityData.ParentYangName = "Cisco-IOS-XR-infra-xtc-oper"
    pce.EntityData.SegmentPath = "Cisco-IOS-XR-infra-xtc-oper:pce"
    pce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pce.EntityData.Children = make(map[string]types.YChild)
    pce.EntityData.Children["association-infos"] = types.YChild{"AssociationInfos", &pce.AssociationInfos}
    pce.EntityData.Children["cspf"] = types.YChild{"Cspf", &pce.Cspf}
    pce.EntityData.Children["topology-summary"] = types.YChild{"TopologySummary", &pce.TopologySummary}
    pce.EntityData.Children["tunnel-infos"] = types.YChild{"TunnelInfos", &pce.TunnelInfos}
    pce.EntityData.Children["peer-detail-infos"] = types.YChild{"PeerDetailInfos", &pce.PeerDetailInfos}
    pce.EntityData.Children["topology-nodes"] = types.YChild{"TopologyNodes", &pce.TopologyNodes}
    pce.EntityData.Children["prefix-infos"] = types.YChild{"PrefixInfos", &pce.PrefixInfos}
    pce.EntityData.Children["lsp-summary"] = types.YChild{"LspSummary", &pce.LspSummary}
    pce.EntityData.Children["peer-infos"] = types.YChild{"PeerInfos", &pce.PeerInfos}
    pce.EntityData.Children["tunnel-detail-infos"] = types.YChild{"TunnelDetailInfos", &pce.TunnelDetailInfos}
    pce.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(pce.EntityData)
}

// Pce_AssociationInfos
// Associaition database in XTC
type Pce_AssociationInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE Association information. The type is slice of
    // Pce_AssociationInfos_AssociationInfo.
    AssociationInfo []Pce_AssociationInfos_AssociationInfo
}

func (associationInfos *Pce_AssociationInfos) GetEntityData() *types.CommonEntityData {
    associationInfos.EntityData.YFilter = associationInfos.YFilter
    associationInfos.EntityData.YangName = "association-infos"
    associationInfos.EntityData.BundleName = "cisco_ios_xr"
    associationInfos.EntityData.ParentYangName = "pce"
    associationInfos.EntityData.SegmentPath = "association-infos"
    associationInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInfos.EntityData.Children = make(map[string]types.YChild)
    associationInfos.EntityData.Children["association-info"] = types.YChild{"AssociationInfo", nil}
    for i := range associationInfos.AssociationInfo {
        associationInfos.EntityData.Children[types.GetSegmentPath(&associationInfos.AssociationInfo[i])] = types.YChild{"AssociationInfo", &associationInfos.AssociationInfo[i]}
    }
    associationInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(associationInfos.EntityData)
}

// Pce_AssociationInfos_AssociationInfo
// PCE Association information
type Pce_AssociationInfos_AssociationInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // -2147483648..2147483647.
    GroupId interface{}

    // Type. The type is PceAsso.
    Type_ interface{}

    // Sub ID. The type is interface{} with range: -2147483648..2147483647.
    SubId interface{}

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetEntityData() *types.CommonEntityData {
    associationInfo.EntityData.YFilter = associationInfo.YFilter
    associationInfo.EntityData.YangName = "association-info"
    associationInfo.EntityData.BundleName = "cisco_ios_xr"
    associationInfo.EntityData.ParentYangName = "association-infos"
    associationInfo.EntityData.SegmentPath = "association-info" + "[group-id='" + fmt.Sprintf("%v", associationInfo.GroupId) + "']"
    associationInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInfo.EntityData.Children = make(map[string]types.YChild)
    associationInfo.EntityData.Children["association-lsp"] = types.YChild{"AssociationLsp", nil}
    for i := range associationInfo.AssociationLsp {
        associationInfo.EntityData.Children[types.GetSegmentPath(&associationInfo.AssociationLsp[i])] = types.YChild{"AssociationLsp", &associationInfo.AssociationLsp[i]}
    }
    associationInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    associationInfo.EntityData.Leafs["group-id"] = types.YLeaf{"GroupId", associationInfo.GroupId}
    associationInfo.EntityData.Leafs["type"] = types.YLeaf{"Type_", associationInfo.Type_}
    associationInfo.EntityData.Leafs["sub-id"] = types.YLeaf{"SubId", associationInfo.SubId}
    associationInfo.EntityData.Leafs["association-type"] = types.YLeaf{"AssociationType", associationInfo.AssociationType}
    associationInfo.EntityData.Leafs["association-id"] = types.YLeaf{"AssociationId", associationInfo.AssociationId}
    associationInfo.EntityData.Leafs["association-source"] = types.YLeaf{"AssociationSource", associationInfo.AssociationSource}
    associationInfo.EntityData.Leafs["strict"] = types.YLeaf{"Strict", associationInfo.Strict}
    associationInfo.EntityData.Leafs["status"] = types.YLeaf{"Status", associationInfo.Status}
    associationInfo.EntityData.Leafs["headends-swapped"] = types.YLeaf{"HeadendsSwapped", associationInfo.HeadendsSwapped}
    return &(associationInfo.EntityData)
}

// Pce_AssociationInfos_AssociationInfo_AssociationLsp
// Association LSP Info
type Pce_AssociationInfos_AssociationInfo_AssociationLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (associationLsp *Pce_AssociationInfos_AssociationInfo_AssociationLsp) GetEntityData() *types.CommonEntityData {
    associationLsp.EntityData.YFilter = associationLsp.YFilter
    associationLsp.EntityData.YangName = "association-lsp"
    associationLsp.EntityData.BundleName = "cisco_ios_xr"
    associationLsp.EntityData.ParentYangName = "association-info"
    associationLsp.EntityData.SegmentPath = "association-lsp"
    associationLsp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationLsp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationLsp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationLsp.EntityData.Children = make(map[string]types.YChild)
    associationLsp.EntityData.Leafs = make(map[string]types.YLeaf)
    associationLsp.EntityData.Leafs["pcc-address"] = types.YLeaf{"PccAddress", associationLsp.PccAddress}
    associationLsp.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", associationLsp.TunnelId}
    associationLsp.EntityData.Leafs["lspid"] = types.YLeaf{"Lspid", associationLsp.Lspid}
    associationLsp.EntityData.Leafs["tunnel-name"] = types.YLeaf{"TunnelName", associationLsp.TunnelName}
    associationLsp.EntityData.Leafs["pce-based"] = types.YLeaf{"PceBased", associationLsp.PceBased}
    associationLsp.EntityData.Leafs["plsp-id"] = types.YLeaf{"PlspId", associationLsp.PlspId}
    return &(associationLsp.EntityData)
}

// Pce_Cspf
// CSPF path info
type Pce_Cspf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This table models the path calculation capabilities in XTC.A GET operation
    // for the complete table will return no entries.
    CspfPaths Pce_Cspf_CspfPaths
}

func (cspf *Pce_Cspf) GetEntityData() *types.CommonEntityData {
    cspf.EntityData.YFilter = cspf.YFilter
    cspf.EntityData.YangName = "cspf"
    cspf.EntityData.BundleName = "cisco_ios_xr"
    cspf.EntityData.ParentYangName = "pce"
    cspf.EntityData.SegmentPath = "cspf"
    cspf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cspf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cspf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cspf.EntityData.Children = make(map[string]types.YChild)
    cspf.EntityData.Children["cspf-paths"] = types.YChild{"CspfPaths", &cspf.CspfPaths}
    cspf.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cspf.EntityData)
}

// Pce_Cspf_CspfPaths
// This table models the path calculation
// capabilities in XTC.A GET operation for the
// complete table will return no entries.
type Pce_Cspf_CspfPaths struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A GET operation on this class returns the path . The type is slice of
    // Pce_Cspf_CspfPaths_CspfPath.
    CspfPath []Pce_Cspf_CspfPaths_CspfPath
}

func (cspfPaths *Pce_Cspf_CspfPaths) GetEntityData() *types.CommonEntityData {
    cspfPaths.EntityData.YFilter = cspfPaths.YFilter
    cspfPaths.EntityData.YangName = "cspf-paths"
    cspfPaths.EntityData.BundleName = "cisco_ios_xr"
    cspfPaths.EntityData.ParentYangName = "cspf"
    cspfPaths.EntityData.SegmentPath = "cspf-paths"
    cspfPaths.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cspfPaths.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cspfPaths.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cspfPaths.EntityData.Children = make(map[string]types.YChild)
    cspfPaths.EntityData.Children["cspf-path"] = types.YChild{"CspfPath", nil}
    for i := range cspfPaths.CspfPath {
        cspfPaths.EntityData.Children[types.GetSegmentPath(&cspfPaths.CspfPath[i])] = types.YChild{"CspfPath", &cspfPaths.CspfPath[i]}
    }
    cspfPaths.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cspfPaths.EntityData)
}

// Pce_Cspf_CspfPaths_CspfPath
// A GET operation on this class returns the path
// .
type Pce_Cspf_CspfPaths_CspfPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is interface{} with
    // range: -2147483648..2147483647.
    Af interface{}

    // This attribute is a key. Source of path 1. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Source1 interface{}

    // This attribute is a key. Destination of path 1. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Destination1 interface{}

    // This attribute is a key. Metric type. The type is interface{} with range:
    // -2147483648..2147483647.
    MetricType interface{}

    // This attribute is a key. Source of path 2. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Source2 interface{}

    // This attribute is a key. Destination of path 2. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetEntityData() *types.CommonEntityData {
    cspfPath.EntityData.YFilter = cspfPath.YFilter
    cspfPath.EntityData.YangName = "cspf-path"
    cspfPath.EntityData.BundleName = "cisco_ios_xr"
    cspfPath.EntityData.ParentYangName = "cspf-paths"
    cspfPath.EntityData.SegmentPath = "cspf-path" + "[af='" + fmt.Sprintf("%v", cspfPath.Af) + "']" + "[source1='" + fmt.Sprintf("%v", cspfPath.Source1) + "']" + "[destination1='" + fmt.Sprintf("%v", cspfPath.Destination1) + "']" + "[metric-type='" + fmt.Sprintf("%v", cspfPath.MetricType) + "']" + "[source2='" + fmt.Sprintf("%v", cspfPath.Source2) + "']" + "[destination2='" + fmt.Sprintf("%v", cspfPath.Destination2) + "']" + "[disjoint-level='" + fmt.Sprintf("%v", cspfPath.DisjointLevel) + "']" + "[disjoint-strict='" + fmt.Sprintf("%v", cspfPath.DisjointStrict) + "']" + "[shortest-path='" + fmt.Sprintf("%v", cspfPath.ShortestPath) + "']"
    cspfPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cspfPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cspfPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cspfPath.EntityData.Children = make(map[string]types.YChild)
    cspfPath.EntityData.Children["output-path"] = types.YChild{"OutputPath", nil}
    for i := range cspfPath.OutputPath {
        cspfPath.EntityData.Children[types.GetSegmentPath(&cspfPath.OutputPath[i])] = types.YChild{"OutputPath", &cspfPath.OutputPath[i]}
    }
    cspfPath.EntityData.Leafs = make(map[string]types.YLeaf)
    cspfPath.EntityData.Leafs["af"] = types.YLeaf{"Af", cspfPath.Af}
    cspfPath.EntityData.Leafs["source1"] = types.YLeaf{"Source1", cspfPath.Source1}
    cspfPath.EntityData.Leafs["destination1"] = types.YLeaf{"Destination1", cspfPath.Destination1}
    cspfPath.EntityData.Leafs["metric-type"] = types.YLeaf{"MetricType", cspfPath.MetricType}
    cspfPath.EntityData.Leafs["source2"] = types.YLeaf{"Source2", cspfPath.Source2}
    cspfPath.EntityData.Leafs["destination2"] = types.YLeaf{"Destination2", cspfPath.Destination2}
    cspfPath.EntityData.Leafs["disjoint-level"] = types.YLeaf{"DisjointLevel", cspfPath.DisjointLevel}
    cspfPath.EntityData.Leafs["disjoint-strict"] = types.YLeaf{"DisjointStrict", cspfPath.DisjointStrict}
    cspfPath.EntityData.Leafs["shortest-path"] = types.YLeaf{"ShortestPath", cspfPath.ShortestPath}
    cspfPath.EntityData.Leafs["headends-swapped"] = types.YLeaf{"HeadendsSwapped", cspfPath.HeadendsSwapped}
    cspfPath.EntityData.Leafs["cspf-result"] = types.YLeaf{"CspfResult", cspfPath.CspfResult}
    return &(cspfPath.EntityData)
}

// Pce_Cspf_CspfPaths_CspfPath_OutputPath
// Output PCE paths
type Pce_Cspf_CspfPaths_CspfPath_OutputPath struct {
    EntityData types.CommonEntityData
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

func (outputPath *Pce_Cspf_CspfPaths_CspfPath_OutputPath) GetEntityData() *types.CommonEntityData {
    outputPath.EntityData.YFilter = outputPath.YFilter
    outputPath.EntityData.YangName = "output-path"
    outputPath.EntityData.BundleName = "cisco_ios_xr"
    outputPath.EntityData.ParentYangName = "cspf-path"
    outputPath.EntityData.SegmentPath = "output-path"
    outputPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    outputPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    outputPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    outputPath.EntityData.Children = make(map[string]types.YChild)
    outputPath.EntityData.Children["source"] = types.YChild{"Source", &outputPath.Source}
    outputPath.EntityData.Children["destination"] = types.YChild{"Destination", &outputPath.Destination}
    outputPath.EntityData.Children["hops"] = types.YChild{"Hops", nil}
    for i := range outputPath.Hops {
        outputPath.EntityData.Children[types.GetSegmentPath(&outputPath.Hops[i])] = types.YChild{"Hops", &outputPath.Hops[i]}
    }
    outputPath.EntityData.Leafs = make(map[string]types.YLeaf)
    outputPath.EntityData.Leafs["cost"] = types.YLeaf{"Cost", outputPath.Cost}
    return &(outputPath.EntityData)
}

// Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source
// Source of path
type Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (source *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Source) GetEntityData() *types.CommonEntityData {
    source.EntityData.YFilter = source.YFilter
    source.EntityData.YangName = "source"
    source.EntityData.BundleName = "cisco_ios_xr"
    source.EntityData.ParentYangName = "output-path"
    source.EntityData.SegmentPath = "source"
    source.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    source.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    source.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    source.EntityData.Children = make(map[string]types.YChild)
    source.EntityData.Leafs = make(map[string]types.YLeaf)
    source.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", source.AfName}
    source.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", source.Ipv4}
    source.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", source.Ipv6}
    return &(source.EntityData)
}

// Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination
// Destination of path
type Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (destination *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Destination) GetEntityData() *types.CommonEntityData {
    destination.EntityData.YFilter = destination.YFilter
    destination.EntityData.YangName = "destination"
    destination.EntityData.BundleName = "cisco_ios_xr"
    destination.EntityData.ParentYangName = "output-path"
    destination.EntityData.SegmentPath = "destination"
    destination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    destination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    destination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    destination.EntityData.Children = make(map[string]types.YChild)
    destination.EntityData.Leafs = make(map[string]types.YLeaf)
    destination.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", destination.AfName}
    destination.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", destination.Ipv4}
    destination.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", destination.Ipv6}
    return &(destination.EntityData)
}

// Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops
// Hop addresses
type Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family. The type is interface{} with range: 0..255.
    AddressFamily interface{}

    // IPv4 prefix. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Prefix interface{}

    // IPv6 prefix. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6Prefix interface{}
}

func (hops *Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops) GetEntityData() *types.CommonEntityData {
    hops.EntityData.YFilter = hops.YFilter
    hops.EntityData.YangName = "hops"
    hops.EntityData.BundleName = "cisco_ios_xr"
    hops.EntityData.ParentYangName = "output-path"
    hops.EntityData.SegmentPath = "hops"
    hops.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hops.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hops.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hops.EntityData.Children = make(map[string]types.YChild)
    hops.EntityData.Leafs = make(map[string]types.YLeaf)
    hops.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", hops.AddressFamily}
    hops.EntityData.Leafs["ipv4-prefix"] = types.YLeaf{"Ipv4Prefix", hops.Ipv4Prefix}
    hops.EntityData.Leafs["ipv6-prefix"] = types.YLeaf{"Ipv6Prefix", hops.Ipv6Prefix}
    return &(hops.EntityData)
}

// Pce_TopologySummary
// Node summary database in XTC
type Pce_TopologySummary struct {
    EntityData types.CommonEntityData
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

func (topologySummary *Pce_TopologySummary) GetEntityData() *types.CommonEntityData {
    topologySummary.EntityData.YFilter = topologySummary.YFilter
    topologySummary.EntityData.YangName = "topology-summary"
    topologySummary.EntityData.BundleName = "cisco_ios_xr"
    topologySummary.EntityData.ParentYangName = "pce"
    topologySummary.EntityData.SegmentPath = "topology-summary"
    topologySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologySummary.EntityData.Children = make(map[string]types.YChild)
    topologySummary.EntityData.Children["stats-topology-update"] = types.YChild{"StatsTopologyUpdate", &topologySummary.StatsTopologyUpdate}
    topologySummary.EntityData.Leafs = make(map[string]types.YLeaf)
    topologySummary.EntityData.Leafs["nodes"] = types.YLeaf{"Nodes", topologySummary.Nodes}
    topologySummary.EntityData.Leafs["lookup-nodes"] = types.YLeaf{"LookupNodes", topologySummary.LookupNodes}
    topologySummary.EntityData.Leafs["prefixes"] = types.YLeaf{"Prefixes", topologySummary.Prefixes}
    topologySummary.EntityData.Leafs["prefix-sids"] = types.YLeaf{"PrefixSids", topologySummary.PrefixSids}
    topologySummary.EntityData.Leafs["regular-prefix-sids"] = types.YLeaf{"RegularPrefixSids", topologySummary.RegularPrefixSids}
    topologySummary.EntityData.Leafs["strict-prefix-sids"] = types.YLeaf{"StrictPrefixSids", topologySummary.StrictPrefixSids}
    topologySummary.EntityData.Leafs["links"] = types.YLeaf{"Links", topologySummary.Links}
    topologySummary.EntityData.Leafs["epe-links"] = types.YLeaf{"EpeLinks", topologySummary.EpeLinks}
    topologySummary.EntityData.Leafs["adjacency-sids"] = types.YLeaf{"AdjacencySids", topologySummary.AdjacencySids}
    topologySummary.EntityData.Leafs["epesids"] = types.YLeaf{"Epesids", topologySummary.Epesids}
    topologySummary.EntityData.Leafs["protected-adjacency-sids"] = types.YLeaf{"ProtectedAdjacencySids", topologySummary.ProtectedAdjacencySids}
    topologySummary.EntityData.Leafs["un-protected-adjacency-sids"] = types.YLeaf{"UnProtectedAdjacencySids", topologySummary.UnProtectedAdjacencySids}
    topologySummary.EntityData.Leafs["topology-consistent"] = types.YLeaf{"TopologyConsistent", topologySummary.TopologyConsistent}
    return &(topologySummary.EntityData)
}

// Pce_TopologySummary_StatsTopologyUpdate
// Statistics on topology update
type Pce_TopologySummary_StatsTopologyUpdate struct {
    EntityData types.CommonEntityData
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

func (statsTopologyUpdate *Pce_TopologySummary_StatsTopologyUpdate) GetEntityData() *types.CommonEntityData {
    statsTopologyUpdate.EntityData.YFilter = statsTopologyUpdate.YFilter
    statsTopologyUpdate.EntityData.YangName = "stats-topology-update"
    statsTopologyUpdate.EntityData.BundleName = "cisco_ios_xr"
    statsTopologyUpdate.EntityData.ParentYangName = "topology-summary"
    statsTopologyUpdate.EntityData.SegmentPath = "stats-topology-update"
    statsTopologyUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statsTopologyUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statsTopologyUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statsTopologyUpdate.EntityData.Children = make(map[string]types.YChild)
    statsTopologyUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    statsTopologyUpdate.EntityData.Leafs["num-nodes-added"] = types.YLeaf{"NumNodesAdded", statsTopologyUpdate.NumNodesAdded}
    statsTopologyUpdate.EntityData.Leafs["num-nodes-deleted"] = types.YLeaf{"NumNodesDeleted", statsTopologyUpdate.NumNodesDeleted}
    statsTopologyUpdate.EntityData.Leafs["num-links-added"] = types.YLeaf{"NumLinksAdded", statsTopologyUpdate.NumLinksAdded}
    statsTopologyUpdate.EntityData.Leafs["num-links-deleted"] = types.YLeaf{"NumLinksDeleted", statsTopologyUpdate.NumLinksDeleted}
    statsTopologyUpdate.EntityData.Leafs["num-prefixes-added"] = types.YLeaf{"NumPrefixesAdded", statsTopologyUpdate.NumPrefixesAdded}
    statsTopologyUpdate.EntityData.Leafs["num-prefixes-deleted"] = types.YLeaf{"NumPrefixesDeleted", statsTopologyUpdate.NumPrefixesDeleted}
    return &(statsTopologyUpdate.EntityData)
}

// Pce_TunnelInfos
// Tunnel database in XTC
type Pce_TunnelInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of Pce_TunnelInfos_TunnelInfo.
    TunnelInfo []Pce_TunnelInfos_TunnelInfo
}

func (tunnelInfos *Pce_TunnelInfos) GetEntityData() *types.CommonEntityData {
    tunnelInfos.EntityData.YFilter = tunnelInfos.YFilter
    tunnelInfos.EntityData.YangName = "tunnel-infos"
    tunnelInfos.EntityData.BundleName = "cisco_ios_xr"
    tunnelInfos.EntityData.ParentYangName = "pce"
    tunnelInfos.EntityData.SegmentPath = "tunnel-infos"
    tunnelInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInfos.EntityData.Children = make(map[string]types.YChild)
    tunnelInfos.EntityData.Children["tunnel-info"] = types.YChild{"TunnelInfo", nil}
    for i := range tunnelInfos.TunnelInfo {
        tunnelInfos.EntityData.Children[types.GetSegmentPath(&tunnelInfos.TunnelInfo[i])] = types.YChild{"TunnelInfo", &tunnelInfos.TunnelInfo[i]}
    }
    tunnelInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelInfos.EntityData)
}

// Pce_TunnelInfos_TunnelInfo
// Tunnel information
type Pce_TunnelInfos_TunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Brief LSP information. The type is slice of
    // Pce_TunnelInfos_TunnelInfo_BriefLspInformation.
    BriefLspInformation []Pce_TunnelInfos_TunnelInfo_BriefLspInformation
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetEntityData() *types.CommonEntityData {
    tunnelInfo.EntityData.YFilter = tunnelInfo.YFilter
    tunnelInfo.EntityData.YangName = "tunnel-info"
    tunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelInfo.EntityData.ParentYangName = "tunnel-infos"
    tunnelInfo.EntityData.SegmentPath = "tunnel-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelInfo.TunnelName) + "']"
    tunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInfo.EntityData.Children = make(map[string]types.YChild)
    tunnelInfo.EntityData.Children["brief-lsp-information"] = types.YChild{"BriefLspInformation", nil}
    for i := range tunnelInfo.BriefLspInformation {
        tunnelInfo.EntityData.Children[types.GetSegmentPath(&tunnelInfo.BriefLspInformation[i])] = types.YChild{"BriefLspInformation", &tunnelInfo.BriefLspInformation[i]}
    }
    tunnelInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", tunnelInfo.PeerAddress}
    tunnelInfo.EntityData.Leafs["plsp-id"] = types.YLeaf{"PlspId", tunnelInfo.PlspId}
    tunnelInfo.EntityData.Leafs["tunnel-name"] = types.YLeaf{"TunnelName", tunnelInfo.TunnelName}
    tunnelInfo.EntityData.Leafs["pcc-address"] = types.YLeaf{"PccAddress", tunnelInfo.PccAddress}
    tunnelInfo.EntityData.Leafs["tunnel-name-xr"] = types.YLeaf{"TunnelNameXr", tunnelInfo.TunnelNameXr}
    return &(tunnelInfo.EntityData)
}

// Pce_TunnelInfos_TunnelInfo_BriefLspInformation
// Brief LSP information
type Pce_TunnelInfos_TunnelInfo_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (briefLspInformation *Pce_TunnelInfos_TunnelInfo_BriefLspInformation) GetEntityData() *types.CommonEntityData {
    briefLspInformation.EntityData.YFilter = briefLspInformation.YFilter
    briefLspInformation.EntityData.YangName = "brief-lsp-information"
    briefLspInformation.EntityData.BundleName = "cisco_ios_xr"
    briefLspInformation.EntityData.ParentYangName = "tunnel-info"
    briefLspInformation.EntityData.SegmentPath = "brief-lsp-information"
    briefLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLspInformation.EntityData.Children = make(map[string]types.YChild)
    briefLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefLspInformation.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", briefLspInformation.SourceAddress}
    briefLspInformation.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", briefLspInformation.DestinationAddress}
    briefLspInformation.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", briefLspInformation.TunnelId}
    briefLspInformation.EntityData.Leafs["lspid"] = types.YLeaf{"Lspid", briefLspInformation.Lspid}
    briefLspInformation.EntityData.Leafs["binding-sid"] = types.YLeaf{"BindingSid", briefLspInformation.BindingSid}
    briefLspInformation.EntityData.Leafs["lsp-setup-type"] = types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType}
    briefLspInformation.EntityData.Leafs["operational-state"] = types.YLeaf{"OperationalState", briefLspInformation.OperationalState}
    briefLspInformation.EntityData.Leafs["administrative-state"] = types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState}
    return &(briefLspInformation.EntityData)
}

// Pce_PeerDetailInfos
// Detailed peers database in XTC
type Pce_PeerDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed PCE peer information. The type is slice of
    // Pce_PeerDetailInfos_PeerDetailInfo.
    PeerDetailInfo []Pce_PeerDetailInfos_PeerDetailInfo
}

func (peerDetailInfos *Pce_PeerDetailInfos) GetEntityData() *types.CommonEntityData {
    peerDetailInfos.EntityData.YFilter = peerDetailInfos.YFilter
    peerDetailInfos.EntityData.YangName = "peer-detail-infos"
    peerDetailInfos.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfos.EntityData.ParentYangName = "pce"
    peerDetailInfos.EntityData.SegmentPath = "peer-detail-infos"
    peerDetailInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfos.EntityData.Children = make(map[string]types.YChild)
    peerDetailInfos.EntityData.Children["peer-detail-info"] = types.YChild{"PeerDetailInfo", nil}
    for i := range peerDetailInfos.PeerDetailInfo {
        peerDetailInfos.EntityData.Children[types.GetSegmentPath(&peerDetailInfos.PeerDetailInfo[i])] = types.YChild{"PeerDetailInfo", &peerDetailInfos.PeerDetailInfo[i]}
    }
    peerDetailInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerDetailInfos.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo
// Detailed PCE peer information
type Pce_PeerDetailInfos_PeerDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Detailed PCE protocol information.
    DetailPcepInformation Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetEntityData() *types.CommonEntityData {
    peerDetailInfo.EntityData.YFilter = peerDetailInfo.YFilter
    peerDetailInfo.EntityData.YangName = "peer-detail-info"
    peerDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfo.EntityData.ParentYangName = "peer-detail-infos"
    peerDetailInfo.EntityData.SegmentPath = "peer-detail-info" + "[peer-address='" + fmt.Sprintf("%v", peerDetailInfo.PeerAddress) + "']"
    peerDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfo.EntityData.Children = make(map[string]types.YChild)
    peerDetailInfo.EntityData.Children["detail-pcep-information"] = types.YChild{"DetailPcepInformation", &peerDetailInfo.DetailPcepInformation}
    peerDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerDetailInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", peerDetailInfo.PeerAddress}
    peerDetailInfo.EntityData.Leafs["peer-address-xr"] = types.YLeaf{"PeerAddressXr", peerDetailInfo.PeerAddressXr}
    peerDetailInfo.EntityData.Leafs["peer-protocol"] = types.YLeaf{"PeerProtocol", peerDetailInfo.PeerProtocol}
    return &(peerDetailInfo.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
// Detailed PCE protocol information
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation struct {
    EntityData types.CommonEntityData
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

func (detailPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation) GetEntityData() *types.CommonEntityData {
    detailPcepInformation.EntityData.YFilter = detailPcepInformation.YFilter
    detailPcepInformation.EntityData.YangName = "detail-pcep-information"
    detailPcepInformation.EntityData.BundleName = "cisco_ios_xr"
    detailPcepInformation.EntityData.ParentYangName = "peer-detail-info"
    detailPcepInformation.EntityData.SegmentPath = "detail-pcep-information"
    detailPcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailPcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailPcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailPcepInformation.EntityData.Children = make(map[string]types.YChild)
    detailPcepInformation.EntityData.Children["brief-pcep-information"] = types.YChild{"BriefPcepInformation", &detailPcepInformation.BriefPcepInformation}
    detailPcepInformation.EntityData.Children["last-error-rx"] = types.YChild{"LastErrorRx", &detailPcepInformation.LastErrorRx}
    detailPcepInformation.EntityData.Children["last-error-tx"] = types.YChild{"LastErrorTx", &detailPcepInformation.LastErrorTx}
    detailPcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    detailPcepInformation.EntityData.Leafs["error"] = types.YLeaf{"Error", detailPcepInformation.Error}
    detailPcepInformation.EntityData.Leafs["speaker-id"] = types.YLeaf{"SpeakerId", detailPcepInformation.SpeakerId}
    detailPcepInformation.EntityData.Leafs["pcep-up-time"] = types.YLeaf{"PcepUpTime", detailPcepInformation.PcepUpTime}
    detailPcepInformation.EntityData.Leafs["keepalives"] = types.YLeaf{"Keepalives", detailPcepInformation.Keepalives}
    detailPcepInformation.EntityData.Leafs["md5-enabled"] = types.YLeaf{"Md5Enabled", detailPcepInformation.Md5Enabled}
    detailPcepInformation.EntityData.Leafs["keychain-enabled"] = types.YLeaf{"KeychainEnabled", detailPcepInformation.KeychainEnabled}
    detailPcepInformation.EntityData.Leafs["negotiated-local-keepalive"] = types.YLeaf{"NegotiatedLocalKeepalive", detailPcepInformation.NegotiatedLocalKeepalive}
    detailPcepInformation.EntityData.Leafs["negotiated-remote-keepalive"] = types.YLeaf{"NegotiatedRemoteKeepalive", detailPcepInformation.NegotiatedRemoteKeepalive}
    detailPcepInformation.EntityData.Leafs["negotiated-dead-time"] = types.YLeaf{"NegotiatedDeadTime", detailPcepInformation.NegotiatedDeadTime}
    detailPcepInformation.EntityData.Leafs["pce-request-rx"] = types.YLeaf{"PceRequestRx", detailPcepInformation.PceRequestRx}
    detailPcepInformation.EntityData.Leafs["pce-request-tx"] = types.YLeaf{"PceRequestTx", detailPcepInformation.PceRequestTx}
    detailPcepInformation.EntityData.Leafs["pce-reply-rx"] = types.YLeaf{"PceReplyRx", detailPcepInformation.PceReplyRx}
    detailPcepInformation.EntityData.Leafs["pce-reply-tx"] = types.YLeaf{"PceReplyTx", detailPcepInformation.PceReplyTx}
    detailPcepInformation.EntityData.Leafs["pce-error-rx"] = types.YLeaf{"PceErrorRx", detailPcepInformation.PceErrorRx}
    detailPcepInformation.EntityData.Leafs["pce-error-tx"] = types.YLeaf{"PceErrorTx", detailPcepInformation.PceErrorTx}
    detailPcepInformation.EntityData.Leafs["pce-open-tx"] = types.YLeaf{"PceOpenTx", detailPcepInformation.PceOpenTx}
    detailPcepInformation.EntityData.Leafs["pce-open-rx"] = types.YLeaf{"PceOpenRx", detailPcepInformation.PceOpenRx}
    detailPcepInformation.EntityData.Leafs["pce-report-rx"] = types.YLeaf{"PceReportRx", detailPcepInformation.PceReportRx}
    detailPcepInformation.EntityData.Leafs["pce-report-tx"] = types.YLeaf{"PceReportTx", detailPcepInformation.PceReportTx}
    detailPcepInformation.EntityData.Leafs["pce-update-rx"] = types.YLeaf{"PceUpdateRx", detailPcepInformation.PceUpdateRx}
    detailPcepInformation.EntityData.Leafs["pce-update-tx"] = types.YLeaf{"PceUpdateTx", detailPcepInformation.PceUpdateTx}
    detailPcepInformation.EntityData.Leafs["pce-initiate-rx"] = types.YLeaf{"PceInitiateRx", detailPcepInformation.PceInitiateRx}
    detailPcepInformation.EntityData.Leafs["pce-initiate-tx"] = types.YLeaf{"PceInitiateTx", detailPcepInformation.PceInitiateTx}
    detailPcepInformation.EntityData.Leafs["pce-keepalive-tx"] = types.YLeaf{"PceKeepaliveTx", detailPcepInformation.PceKeepaliveTx}
    detailPcepInformation.EntityData.Leafs["pce-keepalive-rx"] = types.YLeaf{"PceKeepaliveRx", detailPcepInformation.PceKeepaliveRx}
    detailPcepInformation.EntityData.Leafs["local-session-id"] = types.YLeaf{"LocalSessionId", detailPcepInformation.LocalSessionId}
    detailPcepInformation.EntityData.Leafs["remote-session-id"] = types.YLeaf{"RemoteSessionId", detailPcepInformation.RemoteSessionId}
    detailPcepInformation.EntityData.Leafs["minimum-keepalive-interval"] = types.YLeaf{"MinimumKeepaliveInterval", detailPcepInformation.MinimumKeepaliveInterval}
    detailPcepInformation.EntityData.Leafs["maximum-dead-interval"] = types.YLeaf{"MaximumDeadInterval", detailPcepInformation.MaximumDeadInterval}
    return &(detailPcepInformation.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation
// Brief PCE protocol information
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation struct {
    EntityData types.CommonEntityData
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

func (briefPcepInformation *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_BriefPcepInformation) GetEntityData() *types.CommonEntityData {
    briefPcepInformation.EntityData.YFilter = briefPcepInformation.YFilter
    briefPcepInformation.EntityData.YangName = "brief-pcep-information"
    briefPcepInformation.EntityData.BundleName = "cisco_ios_xr"
    briefPcepInformation.EntityData.ParentYangName = "detail-pcep-information"
    briefPcepInformation.EntityData.SegmentPath = "brief-pcep-information"
    briefPcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefPcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefPcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefPcepInformation.EntityData.Children = make(map[string]types.YChild)
    briefPcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefPcepInformation.EntityData.Leafs["pcep-state"] = types.YLeaf{"PcepState", briefPcepInformation.PcepState}
    briefPcepInformation.EntityData.Leafs["stateful"] = types.YLeaf{"Stateful", briefPcepInformation.Stateful}
    briefPcepInformation.EntityData.Leafs["capability-update"] = types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate}
    briefPcepInformation.EntityData.Leafs["capability-instantiate"] = types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate}
    briefPcepInformation.EntityData.Leafs["capability-segment-routing"] = types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting}
    briefPcepInformation.EntityData.Leafs["capability-triggered-sync"] = types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync}
    briefPcepInformation.EntityData.Leafs["capability-db-version"] = types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion}
    briefPcepInformation.EntityData.Leafs["capability-delta-sync"] = types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync}
    return &(briefPcepInformation.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx
// Last PCError received
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorRx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorRx) GetEntityData() *types.CommonEntityData {
    lastErrorRx.EntityData.YFilter = lastErrorRx.YFilter
    lastErrorRx.EntityData.YangName = "last-error-rx"
    lastErrorRx.EntityData.BundleName = "cisco_ios_xr"
    lastErrorRx.EntityData.ParentYangName = "detail-pcep-information"
    lastErrorRx.EntityData.SegmentPath = "last-error-rx"
    lastErrorRx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErrorRx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErrorRx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErrorRx.EntityData.Children = make(map[string]types.YChild)
    lastErrorRx.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErrorRx.EntityData.Leafs["pc-error-type"] = types.YLeaf{"PcErrorType", lastErrorRx.PcErrorType}
    lastErrorRx.EntityData.Leafs["pc-error-value"] = types.YLeaf{"PcErrorValue", lastErrorRx.PcErrorValue}
    return &(lastErrorRx.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx
// Last PCError sent
type Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCEP Error type. The type is interface{} with range: 0..255.
    PcErrorType interface{}

    // PCEP Error Value. The type is interface{} with range: 0..255.
    PcErrorValue interface{}
}

func (lastErrorTx *Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation_LastErrorTx) GetEntityData() *types.CommonEntityData {
    lastErrorTx.EntityData.YFilter = lastErrorTx.YFilter
    lastErrorTx.EntityData.YangName = "last-error-tx"
    lastErrorTx.EntityData.BundleName = "cisco_ios_xr"
    lastErrorTx.EntityData.ParentYangName = "detail-pcep-information"
    lastErrorTx.EntityData.SegmentPath = "last-error-tx"
    lastErrorTx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lastErrorTx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lastErrorTx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lastErrorTx.EntityData.Children = make(map[string]types.YChild)
    lastErrorTx.EntityData.Leafs = make(map[string]types.YLeaf)
    lastErrorTx.EntityData.Leafs["pc-error-type"] = types.YLeaf{"PcErrorType", lastErrorTx.PcErrorType}
    lastErrorTx.EntityData.Leafs["pc-error-value"] = types.YLeaf{"PcErrorValue", lastErrorTx.PcErrorValue}
    return &(lastErrorTx.EntityData)
}

// Pce_TopologyNodes
// Node database in XTC
type Pce_TopologyNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node information. The type is slice of Pce_TopologyNodes_TopologyNode.
    TopologyNode []Pce_TopologyNodes_TopologyNode
}

func (topologyNodes *Pce_TopologyNodes) GetEntityData() *types.CommonEntityData {
    topologyNodes.EntityData.YFilter = topologyNodes.YFilter
    topologyNodes.EntityData.YangName = "topology-nodes"
    topologyNodes.EntityData.BundleName = "cisco_ios_xr"
    topologyNodes.EntityData.ParentYangName = "pce"
    topologyNodes.EntityData.SegmentPath = "topology-nodes"
    topologyNodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNodes.EntityData.Children = make(map[string]types.YChild)
    topologyNodes.EntityData.Children["topology-node"] = types.YChild{"TopologyNode", nil}
    for i := range topologyNodes.TopologyNode {
        topologyNodes.EntityData.Children[types.GetSegmentPath(&topologyNodes.TopologyNode[i])] = types.YChild{"TopologyNode", &topologyNodes.TopologyNode[i]}
    }
    topologyNodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(topologyNodes.EntityData)
}

// Pce_TopologyNodes_TopologyNode
// Node information
type Pce_TopologyNodes_TopologyNode struct {
    EntityData types.CommonEntityData
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

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetEntityData() *types.CommonEntityData {
    topologyNode.EntityData.YFilter = topologyNode.YFilter
    topologyNode.EntityData.YangName = "topology-node"
    topologyNode.EntityData.BundleName = "cisco_ios_xr"
    topologyNode.EntityData.ParentYangName = "topology-nodes"
    topologyNode.EntityData.SegmentPath = "topology-node" + "[node-identifier='" + fmt.Sprintf("%v", topologyNode.NodeIdentifier) + "']"
    topologyNode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    topologyNode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    topologyNode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    topologyNode.EntityData.Children = make(map[string]types.YChild)
    topologyNode.EntityData.Children["node-protocol-identifier"] = types.YChild{"NodeProtocolIdentifier", &topologyNode.NodeProtocolIdentifier}
    topologyNode.EntityData.Children["prefix-sid"] = types.YChild{"PrefixSid", nil}
    for i := range topologyNode.PrefixSid {
        topologyNode.EntityData.Children[types.GetSegmentPath(&topologyNode.PrefixSid[i])] = types.YChild{"PrefixSid", &topologyNode.PrefixSid[i]}
    }
    topologyNode.EntityData.Children["ipv4-link"] = types.YChild{"Ipv4Link", nil}
    for i := range topologyNode.Ipv4Link {
        topologyNode.EntityData.Children[types.GetSegmentPath(&topologyNode.Ipv4Link[i])] = types.YChild{"Ipv4Link", &topologyNode.Ipv4Link[i]}
    }
    topologyNode.EntityData.Children["ipv6-link"] = types.YChild{"Ipv6Link", nil}
    for i := range topologyNode.Ipv6Link {
        topologyNode.EntityData.Children[types.GetSegmentPath(&topologyNode.Ipv6Link[i])] = types.YChild{"Ipv6Link", &topologyNode.Ipv6Link[i]}
    }
    topologyNode.EntityData.Leafs = make(map[string]types.YLeaf)
    topologyNode.EntityData.Leafs["node-identifier"] = types.YLeaf{"NodeIdentifier", topologyNode.NodeIdentifier}
    topologyNode.EntityData.Leafs["node-identifier-xr"] = types.YLeaf{"NodeIdentifierXr", topologyNode.NodeIdentifierXr}
    topologyNode.EntityData.Leafs["overload"] = types.YLeaf{"Overload", topologyNode.Overload}
    return &(topologyNode.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier
// Node protocol identifier
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    nodeProtocolIdentifier.EntityData.YFilter = nodeProtocolIdentifier.YFilter
    nodeProtocolIdentifier.EntityData.YangName = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    nodeProtocolIdentifier.EntityData.ParentYangName = "topology-node"
    nodeProtocolIdentifier.EntityData.SegmentPath = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    nodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &nodeProtocolIdentifier.IgpInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &nodeProtocolIdentifier.SrgbInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeProtocolIdentifier.NodeName}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", nodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", nodeProtocolIdentifier.Ipv4BgpRouterId}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", nodeProtocolIdentifier.Ipv4TeRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", nodeProtocolIdentifier.Ipv4TeRouterId}
    return &(nodeProtocolIdentifier.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_PrefixSid
// Prefix SIDs
type Pce_TopologyNodes_TopologyNode_PrefixSid struct {
    EntityData types.CommonEntityData
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

func (prefixSid *Pce_TopologyNodes_TopologyNode_PrefixSid) GetEntityData() *types.CommonEntityData {
    prefixSid.EntityData.YFilter = prefixSid.YFilter
    prefixSid.EntityData.YangName = "prefix-sid"
    prefixSid.EntityData.BundleName = "cisco_ios_xr"
    prefixSid.EntityData.ParentYangName = "topology-node"
    prefixSid.EntityData.SegmentPath = "prefix-sid"
    prefixSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixSid.EntityData.Children = make(map[string]types.YChild)
    prefixSid.EntityData.Children["sid-prefix"] = types.YChild{"SidPrefix", &prefixSid.SidPrefix}
    prefixSid.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixSid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", prefixSid.SidType}
    prefixSid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", prefixSid.MplsLabel}
    prefixSid.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", prefixSid.DomainIdentifier}
    prefixSid.EntityData.Leafs["rflag"] = types.YLeaf{"Rflag", prefixSid.Rflag}
    prefixSid.EntityData.Leafs["nflag"] = types.YLeaf{"Nflag", prefixSid.Nflag}
    prefixSid.EntityData.Leafs["pflag"] = types.YLeaf{"Pflag", prefixSid.Pflag}
    prefixSid.EntityData.Leafs["eflag"] = types.YLeaf{"Eflag", prefixSid.Eflag}
    prefixSid.EntityData.Leafs["vflag"] = types.YLeaf{"Vflag", prefixSid.Vflag}
    prefixSid.EntityData.Leafs["lflag"] = types.YLeaf{"Lflag", prefixSid.Lflag}
    return &(prefixSid.EntityData)
}

// Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_PrefixSid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "prefix-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = make(map[string]types.YChild)
    sidPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    sidPrefix.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", sidPrefix.AfName}
    sidPrefix.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", sidPrefix.Ipv4}
    sidPrefix.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", sidPrefix.Ipv6}
    return &(sidPrefix.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type Pce_TopologyNodes_TopologyNode_Ipv4Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalIpv4Address interface{}

    // Remote IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ipv4Link *Pce_TopologyNodes_TopologyNode_Ipv4Link) GetEntityData() *types.CommonEntityData {
    ipv4Link.EntityData.YFilter = ipv4Link.YFilter
    ipv4Link.EntityData.YangName = "ipv4-link"
    ipv4Link.EntityData.BundleName = "cisco_ios_xr"
    ipv4Link.EntityData.ParentYangName = "topology-node"
    ipv4Link.EntityData.SegmentPath = "ipv4-link"
    ipv4Link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Link.EntityData.Children = make(map[string]types.YChild)
    ipv4Link.EntityData.Children["local-igp-information"] = types.YChild{"LocalIgpInformation", &ipv4Link.LocalIgpInformation}
    ipv4Link.EntityData.Children["remote-node-protocol-identifier"] = types.YChild{"RemoteNodeProtocolIdentifier", &ipv4Link.RemoteNodeProtocolIdentifier}
    ipv4Link.EntityData.Children["adjacency-sid"] = types.YChild{"AdjacencySid", nil}
    for i := range ipv4Link.AdjacencySid {
        ipv4Link.EntityData.Children[types.GetSegmentPath(&ipv4Link.AdjacencySid[i])] = types.YChild{"AdjacencySid", &ipv4Link.AdjacencySid[i]}
    }
    ipv4Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4Link.EntityData.Leafs["local-ipv4-address"] = types.YLeaf{"LocalIpv4Address", ipv4Link.LocalIpv4Address}
    ipv4Link.EntityData.Leafs["remote-ipv4-address"] = types.YLeaf{"RemoteIpv4Address", ipv4Link.RemoteIpv4Address}
    ipv4Link.EntityData.Leafs["igp-metric"] = types.YLeaf{"IgpMetric", ipv4Link.IgpMetric}
    ipv4Link.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", ipv4Link.TeMetric}
    ipv4Link.EntityData.Leafs["maximum-link-bandwidth"] = types.YLeaf{"MaximumLinkBandwidth", ipv4Link.MaximumLinkBandwidth}
    ipv4Link.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", ipv4Link.MaxReservableBandwidth}
    ipv4Link.EntityData.Leafs["srlgs"] = types.YLeaf{"Srlgs", ipv4Link.Srlgs}
    return &(ipv4Link.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
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

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation) GetEntityData() *types.CommonEntityData {
    localIgpInformation.EntityData.YFilter = localIgpInformation.YFilter
    localIgpInformation.EntityData.YangName = "local-igp-information"
    localIgpInformation.EntityData.BundleName = "cisco_ios_xr"
    localIgpInformation.EntityData.ParentYangName = "ipv4-link"
    localIgpInformation.EntityData.SegmentPath = "local-igp-information"
    localIgpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIgpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIgpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIgpInformation.EntityData.Children = make(map[string]types.YChild)
    localIgpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &localIgpInformation.Igp}
    localIgpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    localIgpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier}
    localIgpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", localIgpInformation.AutonomousSystemNumber}
    return &(localIgpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "local-igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    remoteNodeProtocolIdentifier.EntityData.YFilter = remoteNodeProtocolIdentifier.YFilter
    remoteNodeProtocolIdentifier.EntityData.YangName = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteNodeProtocolIdentifier.EntityData.ParentYangName = "ipv4-link"
    remoteNodeProtocolIdentifier.EntityData.SegmentPath = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteNodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &remoteNodeProtocolIdentifier.IgpInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &remoteNodeProtocolIdentifier.SrgbInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteNodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", remoteNodeProtocolIdentifier.NodeName}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", remoteNodeProtocolIdentifier.Ipv4BgpRouterId}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", remoteNodeProtocolIdentifier.Ipv4TeRouterId}
    return &(remoteNodeProtocolIdentifier.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
// Adjacency SIDs
type Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid struct {
    EntityData types.CommonEntityData
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

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "ipv4-link"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = make(map[string]types.YChild)
    adjacencySid.EntityData.Children["sid-prefix"] = types.YChild{"SidPrefix", &adjacencySid.SidPrefix}
    adjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    adjacencySid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", adjacencySid.SidType}
    adjacencySid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", adjacencySid.MplsLabel}
    adjacencySid.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", adjacencySid.DomainIdentifier}
    adjacencySid.EntityData.Leafs["rflag"] = types.YLeaf{"Rflag", adjacencySid.Rflag}
    adjacencySid.EntityData.Leafs["nflag"] = types.YLeaf{"Nflag", adjacencySid.Nflag}
    adjacencySid.EntityData.Leafs["pflag"] = types.YLeaf{"Pflag", adjacencySid.Pflag}
    adjacencySid.EntityData.Leafs["eflag"] = types.YLeaf{"Eflag", adjacencySid.Eflag}
    adjacencySid.EntityData.Leafs["vflag"] = types.YLeaf{"Vflag", adjacencySid.Vflag}
    adjacencySid.EntityData.Leafs["lflag"] = types.YLeaf{"Lflag", adjacencySid.Lflag}
    return &(adjacencySid.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "adjacency-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = make(map[string]types.YChild)
    sidPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    sidPrefix.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", sidPrefix.AfName}
    sidPrefix.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", sidPrefix.Ipv4}
    sidPrefix.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", sidPrefix.Ipv6}
    return &(sidPrefix.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type Pce_TopologyNodes_TopologyNode_Ipv6Link struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LocalIpv6Address interface{}

    // Remote IPv6 address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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

func (ipv6Link *Pce_TopologyNodes_TopologyNode_Ipv6Link) GetEntityData() *types.CommonEntityData {
    ipv6Link.EntityData.YFilter = ipv6Link.YFilter
    ipv6Link.EntityData.YangName = "ipv6-link"
    ipv6Link.EntityData.BundleName = "cisco_ios_xr"
    ipv6Link.EntityData.ParentYangName = "topology-node"
    ipv6Link.EntityData.SegmentPath = "ipv6-link"
    ipv6Link.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Link.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Link.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Link.EntityData.Children = make(map[string]types.YChild)
    ipv6Link.EntityData.Children["local-igp-information"] = types.YChild{"LocalIgpInformation", &ipv6Link.LocalIgpInformation}
    ipv6Link.EntityData.Children["remote-node-protocol-identifier"] = types.YChild{"RemoteNodeProtocolIdentifier", &ipv6Link.RemoteNodeProtocolIdentifier}
    ipv6Link.EntityData.Children["adjacency-sid"] = types.YChild{"AdjacencySid", nil}
    for i := range ipv6Link.AdjacencySid {
        ipv6Link.EntityData.Children[types.GetSegmentPath(&ipv6Link.AdjacencySid[i])] = types.YChild{"AdjacencySid", &ipv6Link.AdjacencySid[i]}
    }
    ipv6Link.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Link.EntityData.Leafs["local-ipv6-address"] = types.YLeaf{"LocalIpv6Address", ipv6Link.LocalIpv6Address}
    ipv6Link.EntityData.Leafs["remote-ipv6-address"] = types.YLeaf{"RemoteIpv6Address", ipv6Link.RemoteIpv6Address}
    ipv6Link.EntityData.Leafs["igp-metric"] = types.YLeaf{"IgpMetric", ipv6Link.IgpMetric}
    ipv6Link.EntityData.Leafs["te-metric"] = types.YLeaf{"TeMetric", ipv6Link.TeMetric}
    ipv6Link.EntityData.Leafs["maximum-link-bandwidth"] = types.YLeaf{"MaximumLinkBandwidth", ipv6Link.MaximumLinkBandwidth}
    ipv6Link.EntityData.Leafs["max-reservable-bandwidth"] = types.YLeaf{"MaxReservableBandwidth", ipv6Link.MaxReservableBandwidth}
    return &(ipv6Link.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
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

func (localIgpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation) GetEntityData() *types.CommonEntityData {
    localIgpInformation.EntityData.YFilter = localIgpInformation.YFilter
    localIgpInformation.EntityData.YangName = "local-igp-information"
    localIgpInformation.EntityData.BundleName = "cisco_ios_xr"
    localIgpInformation.EntityData.ParentYangName = "ipv6-link"
    localIgpInformation.EntityData.SegmentPath = "local-igp-information"
    localIgpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localIgpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localIgpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localIgpInformation.EntityData.Children = make(map[string]types.YChild)
    localIgpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &localIgpInformation.Igp}
    localIgpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    localIgpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier}
    localIgpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", localIgpInformation.AutonomousSystemNumber}
    return &(localIgpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "local-igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier
// Remote node protocol identifier
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
}

func (remoteNodeProtocolIdentifier *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    remoteNodeProtocolIdentifier.EntityData.YFilter = remoteNodeProtocolIdentifier.YFilter
    remoteNodeProtocolIdentifier.EntityData.YangName = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    remoteNodeProtocolIdentifier.EntityData.ParentYangName = "ipv6-link"
    remoteNodeProtocolIdentifier.EntityData.SegmentPath = "remote-node-protocol-identifier"
    remoteNodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteNodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteNodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteNodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    remoteNodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &remoteNodeProtocolIdentifier.IgpInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&remoteNodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &remoteNodeProtocolIdentifier.SrgbInformation[i]}
    }
    remoteNodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteNodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", remoteNodeProtocolIdentifier.NodeName}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", remoteNodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", remoteNodeProtocolIdentifier.Ipv4BgpRouterId}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", remoteNodeProtocolIdentifier.Ipv4TeRouterIdSet}
    remoteNodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", remoteNodeProtocolIdentifier.Ipv4TeRouterId}
    return &(remoteNodeProtocolIdentifier.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "remote-node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
// Adjacency SIDs
type Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid struct {
    EntityData types.CommonEntityData
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

func (adjacencySid *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid) GetEntityData() *types.CommonEntityData {
    adjacencySid.EntityData.YFilter = adjacencySid.YFilter
    adjacencySid.EntityData.YangName = "adjacency-sid"
    adjacencySid.EntityData.BundleName = "cisco_ios_xr"
    adjacencySid.EntityData.ParentYangName = "ipv6-link"
    adjacencySid.EntityData.SegmentPath = "adjacency-sid"
    adjacencySid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    adjacencySid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    adjacencySid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    adjacencySid.EntityData.Children = make(map[string]types.YChild)
    adjacencySid.EntityData.Children["sid-prefix"] = types.YChild{"SidPrefix", &adjacencySid.SidPrefix}
    adjacencySid.EntityData.Leafs = make(map[string]types.YLeaf)
    adjacencySid.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", adjacencySid.SidType}
    adjacencySid.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", adjacencySid.MplsLabel}
    adjacencySid.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", adjacencySid.DomainIdentifier}
    adjacencySid.EntityData.Leafs["rflag"] = types.YLeaf{"Rflag", adjacencySid.Rflag}
    adjacencySid.EntityData.Leafs["nflag"] = types.YLeaf{"Nflag", adjacencySid.Nflag}
    adjacencySid.EntityData.Leafs["pflag"] = types.YLeaf{"Pflag", adjacencySid.Pflag}
    adjacencySid.EntityData.Leafs["eflag"] = types.YLeaf{"Eflag", adjacencySid.Eflag}
    adjacencySid.EntityData.Leafs["vflag"] = types.YLeaf{"Vflag", adjacencySid.Vflag}
    adjacencySid.EntityData.Leafs["lflag"] = types.YLeaf{"Lflag", adjacencySid.Lflag}
    return &(adjacencySid.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "adjacency-sid"
    sidPrefix.EntityData.SegmentPath = "sid-prefix"
    sidPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sidPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sidPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sidPrefix.EntityData.Children = make(map[string]types.YChild)
    sidPrefix.EntityData.Leafs = make(map[string]types.YLeaf)
    sidPrefix.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", sidPrefix.AfName}
    sidPrefix.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", sidPrefix.Ipv4}
    sidPrefix.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", sidPrefix.Ipv6}
    return &(sidPrefix.EntityData)
}

// Pce_PrefixInfos
// Prefixes database in XTC
type Pce_PrefixInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE prefix information. The type is slice of Pce_PrefixInfos_PrefixInfo.
    PrefixInfo []Pce_PrefixInfos_PrefixInfo
}

func (prefixInfos *Pce_PrefixInfos) GetEntityData() *types.CommonEntityData {
    prefixInfos.EntityData.YFilter = prefixInfos.YFilter
    prefixInfos.EntityData.YangName = "prefix-infos"
    prefixInfos.EntityData.BundleName = "cisco_ios_xr"
    prefixInfos.EntityData.ParentYangName = "pce"
    prefixInfos.EntityData.SegmentPath = "prefix-infos"
    prefixInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixInfos.EntityData.Children = make(map[string]types.YChild)
    prefixInfos.EntityData.Children["prefix-info"] = types.YChild{"PrefixInfo", nil}
    for i := range prefixInfos.PrefixInfo {
        prefixInfos.EntityData.Children[types.GetSegmentPath(&prefixInfos.PrefixInfo[i])] = types.YChild{"PrefixInfo", &prefixInfos.PrefixInfo[i]}
    }
    prefixInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prefixInfos.EntityData)
}

// Pce_PrefixInfos_PrefixInfo
// PCE prefix information
type Pce_PrefixInfos_PrefixInfo struct {
    EntityData types.CommonEntityData
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

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetEntityData() *types.CommonEntityData {
    prefixInfo.EntityData.YFilter = prefixInfo.YFilter
    prefixInfo.EntityData.YangName = "prefix-info"
    prefixInfo.EntityData.BundleName = "cisco_ios_xr"
    prefixInfo.EntityData.ParentYangName = "prefix-infos"
    prefixInfo.EntityData.SegmentPath = "prefix-info" + "[node-identifier='" + fmt.Sprintf("%v", prefixInfo.NodeIdentifier) + "']"
    prefixInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixInfo.EntityData.Children = make(map[string]types.YChild)
    prefixInfo.EntityData.Children["node-protocol-identifier"] = types.YChild{"NodeProtocolIdentifier", &prefixInfo.NodeProtocolIdentifier}
    prefixInfo.EntityData.Children["address"] = types.YChild{"Address", nil}
    for i := range prefixInfo.Address {
        prefixInfo.EntityData.Children[types.GetSegmentPath(&prefixInfo.Address[i])] = types.YChild{"Address", &prefixInfo.Address[i]}
    }
    prefixInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    prefixInfo.EntityData.Leafs["node-identifier"] = types.YLeaf{"NodeIdentifier", prefixInfo.NodeIdentifier}
    prefixInfo.EntityData.Leafs["node-identifier-xr"] = types.YLeaf{"NodeIdentifierXr", prefixInfo.NodeIdentifierXr}
    return &(prefixInfo.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier
// Node protocol identifier
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node Name. The type is string.
    NodeName interface{}

    // True if IPv4 BGP router ID is set. The type is bool.
    Ipv4BgpRouterIdSet interface{}

    // IPv4 TE router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4TeRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4TeRouterId interface{}

    // IGP information. The type is slice of
    // Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
}

func (nodeProtocolIdentifier *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier) GetEntityData() *types.CommonEntityData {
    nodeProtocolIdentifier.EntityData.YFilter = nodeProtocolIdentifier.YFilter
    nodeProtocolIdentifier.EntityData.YangName = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.BundleName = "cisco_ios_xr"
    nodeProtocolIdentifier.EntityData.ParentYangName = "prefix-info"
    nodeProtocolIdentifier.EntityData.SegmentPath = "node-protocol-identifier"
    nodeProtocolIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeProtocolIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeProtocolIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeProtocolIdentifier.EntityData.Children = make(map[string]types.YChild)
    nodeProtocolIdentifier.EntityData.Children["igp-information"] = types.YChild{"IgpInformation", nil}
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.IgpInformation[i])] = types.YChild{"IgpInformation", &nodeProtocolIdentifier.IgpInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Children["srgb-information"] = types.YChild{"SrgbInformation", nil}
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children[types.GetSegmentPath(&nodeProtocolIdentifier.SrgbInformation[i])] = types.YChild{"SrgbInformation", &nodeProtocolIdentifier.SrgbInformation[i]}
    }
    nodeProtocolIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    nodeProtocolIdentifier.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", nodeProtocolIdentifier.NodeName}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id-set"] = types.YLeaf{"Ipv4BgpRouterIdSet", nodeProtocolIdentifier.Ipv4BgpRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4-bgp-router-id"] = types.YLeaf{"Ipv4BgpRouterId", nodeProtocolIdentifier.Ipv4BgpRouterId}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id-set"] = types.YLeaf{"Ipv4TeRouterIdSet", nodeProtocolIdentifier.Ipv4TeRouterIdSet}
    nodeProtocolIdentifier.EntityData.Leafs["ipv4te-router-id"] = types.YLeaf{"Ipv4TeRouterId", nodeProtocolIdentifier.Ipv4TeRouterId}
    return &(nodeProtocolIdentifier.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
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

func (igpInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation) GetEntityData() *types.CommonEntityData {
    igpInformation.EntityData.YFilter = igpInformation.YFilter
    igpInformation.EntityData.YangName = "igp-information"
    igpInformation.EntityData.BundleName = "cisco_ios_xr"
    igpInformation.EntityData.ParentYangName = "node-protocol-identifier"
    igpInformation.EntityData.SegmentPath = "igp-information"
    igpInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpInformation.EntityData.Children = make(map[string]types.YChild)
    igpInformation.EntityData.Children["igp"] = types.YChild{"Igp", &igpInformation.Igp}
    igpInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    igpInformation.EntityData.Leafs["domain-identifier"] = types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier}
    igpInformation.EntityData.Leafs["autonomous-system-number"] = types.YLeaf{"AutonomousSystemNumber", igpInformation.AutonomousSystemNumber}
    return &(igpInformation.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp
// IGP-specific information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp struct {
    EntityData types.CommonEntityData
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

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "igp-information"
    igp.EntityData.SegmentPath = "igp"
    igp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igp.EntityData.Children = make(map[string]types.YChild)
    igp.EntityData.Children["isis"] = types.YChild{"Isis", &igp.Isis}
    igp.EntityData.Children["ospf"] = types.YChild{"Ospf", &igp.Ospf}
    igp.EntityData.Children["bgp"] = types.YChild{"Bgp", &igp.Bgp}
    igp.EntityData.Leafs = make(map[string]types.YLeaf)
    igp.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igp.IgpId}
    return &(igp.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis
// ISIS information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf
// OSPF information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp
// BGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_Igp_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
// SRGB information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SRGB start. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // SRGB size. The type is interface{} with range: 0..4294967295.
    Size interface{}

    // IGP-specific information.
    IgpSrgb Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
}

func (srgbInformation *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation) GetEntityData() *types.CommonEntityData {
    srgbInformation.EntityData.YFilter = srgbInformation.YFilter
    srgbInformation.EntityData.YangName = "srgb-information"
    srgbInformation.EntityData.BundleName = "cisco_ios_xr"
    srgbInformation.EntityData.ParentYangName = "node-protocol-identifier"
    srgbInformation.EntityData.SegmentPath = "srgb-information"
    srgbInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srgbInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srgbInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srgbInformation.EntityData.Children = make(map[string]types.YChild)
    srgbInformation.EntityData.Children["igp-srgb"] = types.YChild{"IgpSrgb", &srgbInformation.IgpSrgb}
    srgbInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    srgbInformation.EntityData.Leafs["start"] = types.YLeaf{"Start", srgbInformation.Start}
    srgbInformation.EntityData.Leafs["size"] = types.YLeaf{"Size", srgbInformation.Size}
    return &(srgbInformation.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb
// IGP-specific information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb struct {
    EntityData types.CommonEntityData
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

func (igpSrgb *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb) GetEntityData() *types.CommonEntityData {
    igpSrgb.EntityData.YFilter = igpSrgb.YFilter
    igpSrgb.EntityData.YangName = "igp-srgb"
    igpSrgb.EntityData.BundleName = "cisco_ios_xr"
    igpSrgb.EntityData.ParentYangName = "srgb-information"
    igpSrgb.EntityData.SegmentPath = "igp-srgb"
    igpSrgb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    igpSrgb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    igpSrgb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    igpSrgb.EntityData.Children = make(map[string]types.YChild)
    igpSrgb.EntityData.Children["isis"] = types.YChild{"Isis", &igpSrgb.Isis}
    igpSrgb.EntityData.Children["ospf"] = types.YChild{"Ospf", &igpSrgb.Ospf}
    igpSrgb.EntityData.Children["bgp"] = types.YChild{"Bgp", &igpSrgb.Bgp}
    igpSrgb.EntityData.Leafs = make(map[string]types.YLeaf)
    igpSrgb.EntityData.Leafs["igp-id"] = types.YLeaf{"IgpId", igpSrgb.IgpId}
    return &(igpSrgb.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis
// ISIS information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "igp-srgb"
    isis.EntityData.SegmentPath = "isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["system-id"] = types.YLeaf{"SystemId", isis.SystemId}
    isis.EntityData.Leafs["level"] = types.YLeaf{"Level", isis.Level}
    return &(isis.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf
// OSPF information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "igp-srgb"
    ospf.EntityData.SegmentPath = "ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
    ospf.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", ospf.RouterId}
    ospf.EntityData.Leafs["area"] = types.YLeaf{"Area", ospf.Area}
    return &(ospf.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp
// BGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_IgpSrgb_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "igp-srgb"
    bgp.EntityData.SegmentPath = "bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp.EntityData.Leafs["router-id"] = types.YLeaf{"RouterId", bgp.RouterId}
    bgp.EntityData.Leafs["confed-asn"] = types.YLeaf{"ConfedAsn", bgp.ConfedAsn}
    return &(bgp.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_Address
// Prefix address
type Pce_PrefixInfos_PrefixInfo_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Prefix IP address.
    Ip Pce_PrefixInfos_PrefixInfo_Address_Ip
}

func (address *Pce_PrefixInfos_PrefixInfo_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "prefix-info"
    address.EntityData.SegmentPath = "address"
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = make(map[string]types.YChild)
    address.EntityData.Children["ip"] = types.YChild{"Ip", &address.Ip}
    address.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(address.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_Address_Ip
// Prefix IP address
type Pce_PrefixInfos_PrefixInfo_Address_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFName. The type is PceAfId.
    AfName interface{}

    // IPv4 address type. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (ip *Pce_PrefixInfos_PrefixInfo_Address_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "address"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    ip.EntityData.Leafs["af-name"] = types.YLeaf{"AfName", ip.AfName}
    ip.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", ip.Ipv4}
    ip.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", ip.Ipv6}
    return &(ip.EntityData)
}

// Pce_LspSummary
// LSP summary database in XTC
type Pce_LspSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary for all peers.
    AllLsPs Pce_LspSummary_AllLsPs

    // Number of LSPs for specific peer. The type is slice of
    // Pce_LspSummary_PeerLsPsInfo.
    PeerLsPsInfo []Pce_LspSummary_PeerLsPsInfo
}

func (lspSummary *Pce_LspSummary) GetEntityData() *types.CommonEntityData {
    lspSummary.EntityData.YFilter = lspSummary.YFilter
    lspSummary.EntityData.YangName = "lsp-summary"
    lspSummary.EntityData.BundleName = "cisco_ios_xr"
    lspSummary.EntityData.ParentYangName = "pce"
    lspSummary.EntityData.SegmentPath = "lsp-summary"
    lspSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspSummary.EntityData.Children = make(map[string]types.YChild)
    lspSummary.EntityData.Children["all-ls-ps"] = types.YChild{"AllLsPs", &lspSummary.AllLsPs}
    lspSummary.EntityData.Children["peer-ls-ps-info"] = types.YChild{"PeerLsPsInfo", nil}
    for i := range lspSummary.PeerLsPsInfo {
        lspSummary.EntityData.Children[types.GetSegmentPath(&lspSummary.PeerLsPsInfo[i])] = types.YChild{"PeerLsPsInfo", &lspSummary.PeerLsPsInfo[i]}
    }
    lspSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(lspSummary.EntityData)
}

// Pce_LspSummary_AllLsPs
// Summary for all peers
type Pce_LspSummary_AllLsPs struct {
    EntityData types.CommonEntityData
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

func (allLsPs *Pce_LspSummary_AllLsPs) GetEntityData() *types.CommonEntityData {
    allLsPs.EntityData.YFilter = allLsPs.YFilter
    allLsPs.EntityData.YangName = "all-ls-ps"
    allLsPs.EntityData.BundleName = "cisco_ios_xr"
    allLsPs.EntityData.ParentYangName = "lsp-summary"
    allLsPs.EntityData.SegmentPath = "all-ls-ps"
    allLsPs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allLsPs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allLsPs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allLsPs.EntityData.Children = make(map[string]types.YChild)
    allLsPs.EntityData.Leafs = make(map[string]types.YLeaf)
    allLsPs.EntityData.Leafs["all-ls-ps"] = types.YLeaf{"AllLsPs", allLsPs.AllLsPs}
    allLsPs.EntityData.Leafs["up-ls-ps"] = types.YLeaf{"UpLsPs", allLsPs.UpLsPs}
    allLsPs.EntityData.Leafs["admin-up-ls-ps"] = types.YLeaf{"AdminUpLsPs", allLsPs.AdminUpLsPs}
    allLsPs.EntityData.Leafs["sr-ls-ps"] = types.YLeaf{"SrLsPs", allLsPs.SrLsPs}
    allLsPs.EntityData.Leafs["rsvp-ls-ps"] = types.YLeaf{"RsvpLsPs", allLsPs.RsvpLsPs}
    return &(allLsPs.EntityData)
}

// Pce_LspSummary_PeerLsPsInfo
// Number of LSPs for specific peer
type Pce_LspSummary_PeerLsPsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Peer IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Number of LSPs for specific peer.
    LspSummary Pce_LspSummary_PeerLsPsInfo_LspSummary_
}

func (peerLsPsInfo *Pce_LspSummary_PeerLsPsInfo) GetEntityData() *types.CommonEntityData {
    peerLsPsInfo.EntityData.YFilter = peerLsPsInfo.YFilter
    peerLsPsInfo.EntityData.YangName = "peer-ls-ps-info"
    peerLsPsInfo.EntityData.BundleName = "cisco_ios_xr"
    peerLsPsInfo.EntityData.ParentYangName = "lsp-summary"
    peerLsPsInfo.EntityData.SegmentPath = "peer-ls-ps-info"
    peerLsPsInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerLsPsInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerLsPsInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerLsPsInfo.EntityData.Children = make(map[string]types.YChild)
    peerLsPsInfo.EntityData.Children["lsp-summary"] = types.YChild{"LspSummary", &peerLsPsInfo.LspSummary}
    peerLsPsInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerLsPsInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", peerLsPsInfo.PeerAddress}
    return &(peerLsPsInfo.EntityData)
}

// Pce_LspSummary_PeerLsPsInfo_LspSummary_
// Number of LSPs for specific peer
type Pce_LspSummary_PeerLsPsInfo_LspSummary_ struct {
    EntityData types.CommonEntityData
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

func (lspSummary_ *Pce_LspSummary_PeerLsPsInfo_LspSummary_) GetEntityData() *types.CommonEntityData {
    lspSummary_.EntityData.YFilter = lspSummary_.YFilter
    lspSummary_.EntityData.YangName = "lsp-summary"
    lspSummary_.EntityData.BundleName = "cisco_ios_xr"
    lspSummary_.EntityData.ParentYangName = "peer-ls-ps-info"
    lspSummary_.EntityData.SegmentPath = "lsp-summary"
    lspSummary_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspSummary_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspSummary_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspSummary_.EntityData.Children = make(map[string]types.YChild)
    lspSummary_.EntityData.Leafs = make(map[string]types.YLeaf)
    lspSummary_.EntityData.Leafs["all-ls-ps"] = types.YLeaf{"AllLsPs", lspSummary_.AllLsPs}
    lspSummary_.EntityData.Leafs["up-ls-ps"] = types.YLeaf{"UpLsPs", lspSummary_.UpLsPs}
    lspSummary_.EntityData.Leafs["admin-up-ls-ps"] = types.YLeaf{"AdminUpLsPs", lspSummary_.AdminUpLsPs}
    lspSummary_.EntityData.Leafs["sr-ls-ps"] = types.YLeaf{"SrLsPs", lspSummary_.SrLsPs}
    lspSummary_.EntityData.Leafs["rsvp-ls-ps"] = types.YLeaf{"RsvpLsPs", lspSummary_.RsvpLsPs}
    return &(lspSummary_.EntityData)
}

// Pce_PeerInfos
// Peers database in XTC
type Pce_PeerInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer information. The type is slice of Pce_PeerInfos_PeerInfo.
    PeerInfo []Pce_PeerInfos_PeerInfo
}

func (peerInfos *Pce_PeerInfos) GetEntityData() *types.CommonEntityData {
    peerInfos.EntityData.YFilter = peerInfos.YFilter
    peerInfos.EntityData.YangName = "peer-infos"
    peerInfos.EntityData.BundleName = "cisco_ios_xr"
    peerInfos.EntityData.ParentYangName = "pce"
    peerInfos.EntityData.SegmentPath = "peer-infos"
    peerInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfos.EntityData.Children = make(map[string]types.YChild)
    peerInfos.EntityData.Children["peer-info"] = types.YChild{"PeerInfo", nil}
    for i := range peerInfos.PeerInfo {
        peerInfos.EntityData.Children[types.GetSegmentPath(&peerInfos.PeerInfo[i])] = types.YChild{"PeerInfo", &peerInfos.PeerInfo[i]}
    }
    peerInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(peerInfos.EntityData)
}

// Pce_PeerInfos_PeerInfo
// PCE peer information
type Pce_PeerInfos_PeerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // Peer address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PeerAddressXr interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // PCE protocol information.
    BriefPcepInformation Pce_PeerInfos_PeerInfo_BriefPcepInformation
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetEntityData() *types.CommonEntityData {
    peerInfo.EntityData.YFilter = peerInfo.YFilter
    peerInfo.EntityData.YangName = "peer-info"
    peerInfo.EntityData.BundleName = "cisco_ios_xr"
    peerInfo.EntityData.ParentYangName = "peer-infos"
    peerInfo.EntityData.SegmentPath = "peer-info" + "[peer-address='" + fmt.Sprintf("%v", peerInfo.PeerAddress) + "']"
    peerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfo.EntityData.Children = make(map[string]types.YChild)
    peerInfo.EntityData.Children["brief-pcep-information"] = types.YChild{"BriefPcepInformation", &peerInfo.BriefPcepInformation}
    peerInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    peerInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", peerInfo.PeerAddress}
    peerInfo.EntityData.Leafs["peer-address-xr"] = types.YLeaf{"PeerAddressXr", peerInfo.PeerAddressXr}
    peerInfo.EntityData.Leafs["peer-protocol"] = types.YLeaf{"PeerProtocol", peerInfo.PeerProtocol}
    return &(peerInfo.EntityData)
}

// Pce_PeerInfos_PeerInfo_BriefPcepInformation
// PCE protocol information
type Pce_PeerInfos_PeerInfo_BriefPcepInformation struct {
    EntityData types.CommonEntityData
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

func (briefPcepInformation *Pce_PeerInfos_PeerInfo_BriefPcepInformation) GetEntityData() *types.CommonEntityData {
    briefPcepInformation.EntityData.YFilter = briefPcepInformation.YFilter
    briefPcepInformation.EntityData.YangName = "brief-pcep-information"
    briefPcepInformation.EntityData.BundleName = "cisco_ios_xr"
    briefPcepInformation.EntityData.ParentYangName = "peer-info"
    briefPcepInformation.EntityData.SegmentPath = "brief-pcep-information"
    briefPcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefPcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefPcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefPcepInformation.EntityData.Children = make(map[string]types.YChild)
    briefPcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefPcepInformation.EntityData.Leafs["pcep-state"] = types.YLeaf{"PcepState", briefPcepInformation.PcepState}
    briefPcepInformation.EntityData.Leafs["stateful"] = types.YLeaf{"Stateful", briefPcepInformation.Stateful}
    briefPcepInformation.EntityData.Leafs["capability-update"] = types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate}
    briefPcepInformation.EntityData.Leafs["capability-instantiate"] = types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate}
    briefPcepInformation.EntityData.Leafs["capability-segment-routing"] = types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting}
    briefPcepInformation.EntityData.Leafs["capability-triggered-sync"] = types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync}
    briefPcepInformation.EntityData.Leafs["capability-db-version"] = types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion}
    briefPcepInformation.EntityData.Leafs["capability-delta-sync"] = types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync}
    return &(briefPcepInformation.EntityData)
}

// Pce_TunnelDetailInfos
// Detailed tunnel database in XTC
type Pce_TunnelDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed tunnel information. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo.
    TunnelDetailInfo []Pce_TunnelDetailInfos_TunnelDetailInfo
}

func (tunnelDetailInfos *Pce_TunnelDetailInfos) GetEntityData() *types.CommonEntityData {
    tunnelDetailInfos.EntityData.YFilter = tunnelDetailInfos.YFilter
    tunnelDetailInfos.EntityData.YangName = "tunnel-detail-infos"
    tunnelDetailInfos.EntityData.BundleName = "cisco_ios_xr"
    tunnelDetailInfos.EntityData.ParentYangName = "pce"
    tunnelDetailInfos.EntityData.SegmentPath = "tunnel-detail-infos"
    tunnelDetailInfos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDetailInfos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDetailInfos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDetailInfos.EntityData.Children = make(map[string]types.YChild)
    tunnelDetailInfos.EntityData.Children["tunnel-detail-info"] = types.YChild{"TunnelDetailInfo", nil}
    for i := range tunnelDetailInfos.TunnelDetailInfo {
        tunnelDetailInfos.EntityData.Children[types.GetSegmentPath(&tunnelDetailInfos.TunnelDetailInfo[i])] = types.YChild{"TunnelDetailInfo", &tunnelDetailInfos.TunnelDetailInfo[i]}
    }
    tunnelDetailInfos.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelDetailInfos.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo
// Detailed tunnel information
type Pce_TunnelDetailInfos_TunnelDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // -2147483648..2147483647.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TunnelName interface{}

    // PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PccAddress interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Private LSP information.
    PrivateLspInformation Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation

    // Detail LSP information. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation.
    DetailLspInformation []Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetEntityData() *types.CommonEntityData {
    tunnelDetailInfo.EntityData.YFilter = tunnelDetailInfo.YFilter
    tunnelDetailInfo.EntityData.YangName = "tunnel-detail-info"
    tunnelDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelDetailInfo.EntityData.ParentYangName = "tunnel-detail-infos"
    tunnelDetailInfo.EntityData.SegmentPath = "tunnel-detail-info" + "[peer-address='" + fmt.Sprintf("%v", tunnelDetailInfo.PeerAddress) + "']" + "[plsp-id='" + fmt.Sprintf("%v", tunnelDetailInfo.PlspId) + "']" + "[tunnel-name='" + fmt.Sprintf("%v", tunnelDetailInfo.TunnelName) + "']"
    tunnelDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDetailInfo.EntityData.Children = make(map[string]types.YChild)
    tunnelDetailInfo.EntityData.Children["private-lsp-information"] = types.YChild{"PrivateLspInformation", &tunnelDetailInfo.PrivateLspInformation}
    tunnelDetailInfo.EntityData.Children["detail-lsp-information"] = types.YChild{"DetailLspInformation", nil}
    for i := range tunnelDetailInfo.DetailLspInformation {
        tunnelDetailInfo.EntityData.Children[types.GetSegmentPath(&tunnelDetailInfo.DetailLspInformation[i])] = types.YChild{"DetailLspInformation", &tunnelDetailInfo.DetailLspInformation[i]}
    }
    tunnelDetailInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelDetailInfo.EntityData.Leafs["peer-address"] = types.YLeaf{"PeerAddress", tunnelDetailInfo.PeerAddress}
    tunnelDetailInfo.EntityData.Leafs["plsp-id"] = types.YLeaf{"PlspId", tunnelDetailInfo.PlspId}
    tunnelDetailInfo.EntityData.Leafs["tunnel-name"] = types.YLeaf{"TunnelName", tunnelDetailInfo.TunnelName}
    tunnelDetailInfo.EntityData.Leafs["pcc-address"] = types.YLeaf{"PccAddress", tunnelDetailInfo.PccAddress}
    tunnelDetailInfo.EntityData.Leafs["tunnel-name-xr"] = types.YLeaf{"TunnelNameXr", tunnelDetailInfo.TunnelNameXr}
    return &(tunnelDetailInfo.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation
// Private LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Event buffer. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer.
    EventBuffer []Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
}

func (privateLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation) GetEntityData() *types.CommonEntityData {
    privateLspInformation.EntityData.YFilter = privateLspInformation.YFilter
    privateLspInformation.EntityData.YangName = "private-lsp-information"
    privateLspInformation.EntityData.BundleName = "cisco_ios_xr"
    privateLspInformation.EntityData.ParentYangName = "tunnel-detail-info"
    privateLspInformation.EntityData.SegmentPath = "private-lsp-information"
    privateLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    privateLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    privateLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    privateLspInformation.EntityData.Children = make(map[string]types.YChild)
    privateLspInformation.EntityData.Children["event-buffer"] = types.YChild{"EventBuffer", nil}
    for i := range privateLspInformation.EventBuffer {
        privateLspInformation.EntityData.Children[types.GetSegmentPath(&privateLspInformation.EventBuffer[i])] = types.YChild{"EventBuffer", &privateLspInformation.EventBuffer[i]}
    }
    privateLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(privateLspInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
// LSP Event buffer
type Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event message. The type is string.
    EventMessage interface{}

    // Event time, relative to Jan 1, 1970. The type is interface{} with range:
    // 0..4294967295.
    TimeStamp interface{}
}

func (eventBuffer *Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer) GetEntityData() *types.CommonEntityData {
    eventBuffer.EntityData.YFilter = eventBuffer.YFilter
    eventBuffer.EntityData.YangName = "event-buffer"
    eventBuffer.EntityData.BundleName = "cisco_ios_xr"
    eventBuffer.EntityData.ParentYangName = "private-lsp-information"
    eventBuffer.EntityData.SegmentPath = "event-buffer"
    eventBuffer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eventBuffer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eventBuffer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eventBuffer.EntityData.Children = make(map[string]types.YChild)
    eventBuffer.EntityData.Leafs = make(map[string]types.YLeaf)
    eventBuffer.EntityData.Leafs["event-message"] = types.YLeaf{"EventMessage", eventBuffer.EventMessage}
    eventBuffer.EntityData.Leafs["time-stamp"] = types.YLeaf{"TimeStamp", eventBuffer.TimeStamp}
    return &(eventBuffer.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
// Detail LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation struct {
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SubDelegatedPce interface{}

    // State-sync PCE. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    StateSyncPce interface{}

    // Reporting PCC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (detailLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation) GetEntityData() *types.CommonEntityData {
    detailLspInformation.EntityData.YFilter = detailLspInformation.YFilter
    detailLspInformation.EntityData.YangName = "detail-lsp-information"
    detailLspInformation.EntityData.BundleName = "cisco_ios_xr"
    detailLspInformation.EntityData.ParentYangName = "tunnel-detail-info"
    detailLspInformation.EntityData.SegmentPath = "detail-lsp-information"
    detailLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailLspInformation.EntityData.Children = make(map[string]types.YChild)
    detailLspInformation.EntityData.Children["brief-lsp-information"] = types.YChild{"BriefLspInformation", &detailLspInformation.BriefLspInformation}
    detailLspInformation.EntityData.Children["er-os"] = types.YChild{"ErOs", &detailLspInformation.ErOs}
    detailLspInformation.EntityData.Children["lsppcep-information"] = types.YChild{"LsppcepInformation", &detailLspInformation.LsppcepInformation}
    detailLspInformation.EntityData.Children["lsp-association-info"] = types.YChild{"LspAssociationInfo", &detailLspInformation.LspAssociationInfo}
    detailLspInformation.EntityData.Children["lsp-attributes"] = types.YChild{"LspAttributes", &detailLspInformation.LspAttributes}
    detailLspInformation.EntityData.Children["rro"] = types.YChild{"Rro", nil}
    for i := range detailLspInformation.Rro {
        detailLspInformation.EntityData.Children[types.GetSegmentPath(&detailLspInformation.Rro[i])] = types.YChild{"Rro", &detailLspInformation.Rro[i]}
    }
    detailLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    detailLspInformation.EntityData.Leafs["signaled-bandwidth-specified"] = types.YLeaf{"SignaledBandwidthSpecified", detailLspInformation.SignaledBandwidthSpecified}
    detailLspInformation.EntityData.Leafs["signaled-bandwidth"] = types.YLeaf{"SignaledBandwidth", detailLspInformation.SignaledBandwidth}
    detailLspInformation.EntityData.Leafs["actual-bandwidth-specified"] = types.YLeaf{"ActualBandwidthSpecified", detailLspInformation.ActualBandwidthSpecified}
    detailLspInformation.EntityData.Leafs["actual-bandwidth"] = types.YLeaf{"ActualBandwidth", detailLspInformation.ActualBandwidth}
    detailLspInformation.EntityData.Leafs["lsp-role"] = types.YLeaf{"LspRole", detailLspInformation.LspRole}
    detailLspInformation.EntityData.Leafs["computing-pce"] = types.YLeaf{"ComputingPce", detailLspInformation.ComputingPce}
    detailLspInformation.EntityData.Leafs["sub-delegated-pce"] = types.YLeaf{"SubDelegatedPce", detailLspInformation.SubDelegatedPce}
    detailLspInformation.EntityData.Leafs["state-sync-pce"] = types.YLeaf{"StateSyncPce", detailLspInformation.StateSyncPce}
    detailLspInformation.EntityData.Leafs["reporting-pcc-address"] = types.YLeaf{"ReportingPccAddress", detailLspInformation.ReportingPccAddress}
    detailLspInformation.EntityData.Leafs["srlg-info"] = types.YLeaf{"SrlgInfo", detailLspInformation.SrlgInfo}
    return &(detailLspInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation
// Brief LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Source address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Destination address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (briefLspInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation) GetEntityData() *types.CommonEntityData {
    briefLspInformation.EntityData.YFilter = briefLspInformation.YFilter
    briefLspInformation.EntityData.YangName = "brief-lsp-information"
    briefLspInformation.EntityData.BundleName = "cisco_ios_xr"
    briefLspInformation.EntityData.ParentYangName = "detail-lsp-information"
    briefLspInformation.EntityData.SegmentPath = "brief-lsp-information"
    briefLspInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefLspInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefLspInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefLspInformation.EntityData.Children = make(map[string]types.YChild)
    briefLspInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    briefLspInformation.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", briefLspInformation.SourceAddress}
    briefLspInformation.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", briefLspInformation.DestinationAddress}
    briefLspInformation.EntityData.Leafs["tunnel-id"] = types.YLeaf{"TunnelId", briefLspInformation.TunnelId}
    briefLspInformation.EntityData.Leafs["lspid"] = types.YLeaf{"Lspid", briefLspInformation.Lspid}
    briefLspInformation.EntityData.Leafs["binding-sid"] = types.YLeaf{"BindingSid", briefLspInformation.BindingSid}
    briefLspInformation.EntityData.Leafs["lsp-setup-type"] = types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType}
    briefLspInformation.EntityData.Leafs["operational-state"] = types.YLeaf{"OperationalState", briefLspInformation.OperationalState}
    briefLspInformation.EntityData.Leafs["administrative-state"] = types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState}
    return &(briefLspInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs
// Paths
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs struct {
    EntityData types.CommonEntityData
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

func (erOs *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs) GetEntityData() *types.CommonEntityData {
    erOs.EntityData.YFilter = erOs.YFilter
    erOs.EntityData.YangName = "er-os"
    erOs.EntityData.BundleName = "cisco_ios_xr"
    erOs.EntityData.ParentYangName = "detail-lsp-information"
    erOs.EntityData.SegmentPath = "er-os"
    erOs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    erOs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    erOs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    erOs.EntityData.Children = make(map[string]types.YChild)
    erOs.EntityData.Children["reported-rsvp-path"] = types.YChild{"ReportedRsvpPath", nil}
    for i := range erOs.ReportedRsvpPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ReportedRsvpPath[i])] = types.YChild{"ReportedRsvpPath", &erOs.ReportedRsvpPath[i]}
    }
    erOs.EntityData.Children["reported-sr-path"] = types.YChild{"ReportedSrPath", nil}
    for i := range erOs.ReportedSrPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ReportedSrPath[i])] = types.YChild{"ReportedSrPath", &erOs.ReportedSrPath[i]}
    }
    erOs.EntityData.Children["computed-rsvp-path"] = types.YChild{"ComputedRsvpPath", nil}
    for i := range erOs.ComputedRsvpPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ComputedRsvpPath[i])] = types.YChild{"ComputedRsvpPath", &erOs.ComputedRsvpPath[i]}
    }
    erOs.EntityData.Children["computed-sr-path"] = types.YChild{"ComputedSrPath", nil}
    for i := range erOs.ComputedSrPath {
        erOs.EntityData.Children[types.GetSegmentPath(&erOs.ComputedSrPath[i])] = types.YChild{"ComputedSrPath", &erOs.ComputedSrPath[i]}
    }
    erOs.EntityData.Leafs = make(map[string]types.YLeaf)
    erOs.EntityData.Leafs["reported-metric-type"] = types.YLeaf{"ReportedMetricType", erOs.ReportedMetricType}
    erOs.EntityData.Leafs["reported-metric-value"] = types.YLeaf{"ReportedMetricValue", erOs.ReportedMetricValue}
    erOs.EntityData.Leafs["computed-metric-type"] = types.YLeaf{"ComputedMetricType", erOs.ComputedMetricType}
    erOs.EntityData.Leafs["computed-metric-value"] = types.YLeaf{"ComputedMetricValue", erOs.ComputedMetricValue}
    erOs.EntityData.Leafs["computed-hop-list-time"] = types.YLeaf{"ComputedHopListTime", erOs.ComputedHopListTime}
    return &(erOs.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath
// Reported RSVP path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}
}

func (reportedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath) GetEntityData() *types.CommonEntityData {
    reportedRsvpPath.EntityData.YFilter = reportedRsvpPath.YFilter
    reportedRsvpPath.EntityData.YangName = "reported-rsvp-path"
    reportedRsvpPath.EntityData.BundleName = "cisco_ios_xr"
    reportedRsvpPath.EntityData.ParentYangName = "er-os"
    reportedRsvpPath.EntityData.SegmentPath = "reported-rsvp-path"
    reportedRsvpPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportedRsvpPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportedRsvpPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportedRsvpPath.EntityData.Children = make(map[string]types.YChild)
    reportedRsvpPath.EntityData.Leafs = make(map[string]types.YLeaf)
    reportedRsvpPath.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", reportedRsvpPath.HopAddress}
    return &(reportedRsvpPath.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath
// Reported SR path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddr interface{}
}

func (reportedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath) GetEntityData() *types.CommonEntityData {
    reportedSrPath.EntityData.YFilter = reportedSrPath.YFilter
    reportedSrPath.EntityData.YangName = "reported-sr-path"
    reportedSrPath.EntityData.BundleName = "cisco_ios_xr"
    reportedSrPath.EntityData.ParentYangName = "er-os"
    reportedSrPath.EntityData.SegmentPath = "reported-sr-path"
    reportedSrPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportedSrPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportedSrPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportedSrPath.EntityData.Children = make(map[string]types.YChild)
    reportedSrPath.EntityData.Leafs = make(map[string]types.YLeaf)
    reportedSrPath.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", reportedSrPath.SidType}
    reportedSrPath.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", reportedSrPath.MplsLabel}
    reportedSrPath.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", reportedSrPath.LocalAddr}
    reportedSrPath.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", reportedSrPath.RemoteAddr}
    return &(reportedSrPath.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath
// Computed RSVP path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    HopAddress interface{}
}

func (computedRsvpPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath) GetEntityData() *types.CommonEntityData {
    computedRsvpPath.EntityData.YFilter = computedRsvpPath.YFilter
    computedRsvpPath.EntityData.YangName = "computed-rsvp-path"
    computedRsvpPath.EntityData.BundleName = "cisco_ios_xr"
    computedRsvpPath.EntityData.ParentYangName = "er-os"
    computedRsvpPath.EntityData.SegmentPath = "computed-rsvp-path"
    computedRsvpPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    computedRsvpPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    computedRsvpPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    computedRsvpPath.EntityData.Children = make(map[string]types.YChild)
    computedRsvpPath.EntityData.Leafs = make(map[string]types.YLeaf)
    computedRsvpPath.EntityData.Leafs["hop-address"] = types.YLeaf{"HopAddress", computedRsvpPath.HopAddress}
    return &(computedRsvpPath.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
// Computed SR path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddr interface{}
}

func (computedSrPath *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath) GetEntityData() *types.CommonEntityData {
    computedSrPath.EntityData.YFilter = computedSrPath.YFilter
    computedSrPath.EntityData.YangName = "computed-sr-path"
    computedSrPath.EntityData.BundleName = "cisco_ios_xr"
    computedSrPath.EntityData.ParentYangName = "er-os"
    computedSrPath.EntityData.SegmentPath = "computed-sr-path"
    computedSrPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    computedSrPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    computedSrPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    computedSrPath.EntityData.Children = make(map[string]types.YChild)
    computedSrPath.EntityData.Leafs = make(map[string]types.YLeaf)
    computedSrPath.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", computedSrPath.SidType}
    computedSrPath.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", computedSrPath.MplsLabel}
    computedSrPath.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", computedSrPath.LocalAddr}
    computedSrPath.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", computedSrPath.RemoteAddr}
    return &(computedSrPath.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation
// PCEP related LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation struct {
    EntityData types.CommonEntityData
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

    // PCEP LSP initiated flag. The type is interface{} with range: 0..255.
    PcepFlagC interface{}

    // RSVP error info.
    RsvpError Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
}

func (lsppcepInformation *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation) GetEntityData() *types.CommonEntityData {
    lsppcepInformation.EntityData.YFilter = lsppcepInformation.YFilter
    lsppcepInformation.EntityData.YangName = "lsppcep-information"
    lsppcepInformation.EntityData.BundleName = "cisco_ios_xr"
    lsppcepInformation.EntityData.ParentYangName = "detail-lsp-information"
    lsppcepInformation.EntityData.SegmentPath = "lsppcep-information"
    lsppcepInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsppcepInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsppcepInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsppcepInformation.EntityData.Children = make(map[string]types.YChild)
    lsppcepInformation.EntityData.Children["rsvp-error"] = types.YChild{"RsvpError", &lsppcepInformation.RsvpError}
    lsppcepInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    lsppcepInformation.EntityData.Leafs["pcepid"] = types.YLeaf{"Pcepid", lsppcepInformation.Pcepid}
    lsppcepInformation.EntityData.Leafs["pcep-flag-d"] = types.YLeaf{"PcepFlagD", lsppcepInformation.PcepFlagD}
    lsppcepInformation.EntityData.Leafs["pcep-flag-s"] = types.YLeaf{"PcepFlagS", lsppcepInformation.PcepFlagS}
    lsppcepInformation.EntityData.Leafs["pcep-flag-r"] = types.YLeaf{"PcepFlagR", lsppcepInformation.PcepFlagR}
    lsppcepInformation.EntityData.Leafs["pcep-flag-a"] = types.YLeaf{"PcepFlagA", lsppcepInformation.PcepFlagA}
    lsppcepInformation.EntityData.Leafs["pcep-flag-o"] = types.YLeaf{"PcepFlagO", lsppcepInformation.PcepFlagO}
    lsppcepInformation.EntityData.Leafs["pcep-flag-c"] = types.YLeaf{"PcepFlagC", lsppcepInformation.PcepFlagC}
    return &(lsppcepInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
// RSVP error info
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP error node address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    NodeAddress interface{}

    // RSVP error flags. The type is interface{} with range: 0..255.
    ErrorFlags interface{}

    // RSVP error code. The type is interface{} with range: 0..255.
    ErrorCode interface{}

    // RSVP error value. The type is interface{} with range: 0..65535.
    ErrorValue interface{}
}

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetEntityData() *types.CommonEntityData {
    rsvpError.EntityData.YFilter = rsvpError.YFilter
    rsvpError.EntityData.YangName = "rsvp-error"
    rsvpError.EntityData.BundleName = "cisco_ios_xr"
    rsvpError.EntityData.ParentYangName = "lsppcep-information"
    rsvpError.EntityData.SegmentPath = "rsvp-error"
    rsvpError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvpError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvpError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvpError.EntityData.Children = make(map[string]types.YChild)
    rsvpError.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvpError.EntityData.Leafs["node-address"] = types.YLeaf{"NodeAddress", rsvpError.NodeAddress}
    rsvpError.EntityData.Leafs["error-flags"] = types.YLeaf{"ErrorFlags", rsvpError.ErrorFlags}
    rsvpError.EntityData.Leafs["error-code"] = types.YLeaf{"ErrorCode", rsvpError.ErrorCode}
    rsvpError.EntityData.Leafs["error-value"] = types.YLeaf{"ErrorValue", rsvpError.ErrorValue}
    return &(rsvpError.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo
// LSP association information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Source. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    AssociationSource interface{}
}

func (lspAssociationInfo *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo) GetEntityData() *types.CommonEntityData {
    lspAssociationInfo.EntityData.YFilter = lspAssociationInfo.YFilter
    lspAssociationInfo.EntityData.YangName = "lsp-association-info"
    lspAssociationInfo.EntityData.BundleName = "cisco_ios_xr"
    lspAssociationInfo.EntityData.ParentYangName = "detail-lsp-information"
    lspAssociationInfo.EntityData.SegmentPath = "lsp-association-info"
    lspAssociationInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspAssociationInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspAssociationInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspAssociationInfo.EntityData.Children = make(map[string]types.YChild)
    lspAssociationInfo.EntityData.Leafs = make(map[string]types.YLeaf)
    lspAssociationInfo.EntityData.Leafs["association-type"] = types.YLeaf{"AssociationType", lspAssociationInfo.AssociationType}
    lspAssociationInfo.EntityData.Leafs["association-id"] = types.YLeaf{"AssociationId", lspAssociationInfo.AssociationId}
    lspAssociationInfo.EntityData.Leafs["association-source"] = types.YLeaf{"AssociationSource", lspAssociationInfo.AssociationSource}
    return &(lspAssociationInfo.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes
// LSP attributes
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes struct {
    EntityData types.CommonEntityData
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

func (lspAttributes *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAttributes) GetEntityData() *types.CommonEntityData {
    lspAttributes.EntityData.YFilter = lspAttributes.YFilter
    lspAttributes.EntityData.YangName = "lsp-attributes"
    lspAttributes.EntityData.BundleName = "cisco_ios_xr"
    lspAttributes.EntityData.ParentYangName = "detail-lsp-information"
    lspAttributes.EntityData.SegmentPath = "lsp-attributes"
    lspAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspAttributes.EntityData.Children = make(map[string]types.YChild)
    lspAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    lspAttributes.EntityData.Leafs["affinity-exclude-any"] = types.YLeaf{"AffinityExcludeAny", lspAttributes.AffinityExcludeAny}
    lspAttributes.EntityData.Leafs["affinity-include-any"] = types.YLeaf{"AffinityIncludeAny", lspAttributes.AffinityIncludeAny}
    lspAttributes.EntityData.Leafs["affinity-include-all"] = types.YLeaf{"AffinityIncludeAll", lspAttributes.AffinityIncludeAll}
    lspAttributes.EntityData.Leafs["setup-priority"] = types.YLeaf{"SetupPriority", lspAttributes.SetupPriority}
    lspAttributes.EntityData.Leafs["hold-priority"] = types.YLeaf{"HoldPriority", lspAttributes.HoldPriority}
    lspAttributes.EntityData.Leafs["local-protection"] = types.YLeaf{"LocalProtection", lspAttributes.LocalProtection}
    return &(lspAttributes.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
// RRO
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RRO Type. The type is PceRro.
    RroType interface{}

    // IPv4 address of RRO. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}

    // MPLS label of RRO. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // RRO Flags. The type is interface{} with range: 0..255.
    Flags interface{}

    // Segment Routing RRO info.
    SrRro Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
}

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetEntityData() *types.CommonEntityData {
    rro.EntityData.YFilter = rro.YFilter
    rro.EntityData.YangName = "rro"
    rro.EntityData.BundleName = "cisco_ios_xr"
    rro.EntityData.ParentYangName = "detail-lsp-information"
    rro.EntityData.SegmentPath = "rro"
    rro.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rro.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rro.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rro.EntityData.Children = make(map[string]types.YChild)
    rro.EntityData.Children["sr-rro"] = types.YChild{"SrRro", &rro.SrRro}
    rro.EntityData.Leafs = make(map[string]types.YLeaf)
    rro.EntityData.Leafs["rro-type"] = types.YLeaf{"RroType", rro.RroType}
    rro.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", rro.Ipv4Address}
    rro.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", rro.MplsLabel}
    rro.EntityData.Leafs["flags"] = types.YLeaf{"Flags", rro.Flags}
    return &(rro.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro
// Segment Routing RRO info
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID type. The type is PceSrSid.
    SidType interface{}

    // Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

    // Local Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalAddr interface{}

    // Remote Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteAddr interface{}
}

func (srRro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro) GetEntityData() *types.CommonEntityData {
    srRro.EntityData.YFilter = srRro.YFilter
    srRro.EntityData.YangName = "sr-rro"
    srRro.EntityData.BundleName = "cisco_ios_xr"
    srRro.EntityData.ParentYangName = "rro"
    srRro.EntityData.SegmentPath = "sr-rro"
    srRro.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srRro.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srRro.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srRro.EntityData.Children = make(map[string]types.YChild)
    srRro.EntityData.Leafs = make(map[string]types.YLeaf)
    srRro.EntityData.Leafs["sid-type"] = types.YLeaf{"SidType", srRro.SidType}
    srRro.EntityData.Leafs["mpls-label"] = types.YLeaf{"MplsLabel", srRro.MplsLabel}
    srRro.EntityData.Leafs["local-addr"] = types.YLeaf{"LocalAddr", srRro.LocalAddr}
    srRro.EntityData.Leafs["remote-addr"] = types.YLeaf{"RemoteAddr", srRro.RemoteAddr}
    return &(srRro.EntityData)
}

