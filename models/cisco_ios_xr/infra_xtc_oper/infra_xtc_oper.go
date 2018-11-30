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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

    // IPv6 Node SID
    PceSrSid_ipv6_node_sid PceSrSid = "ipv6-node-sid"

    // IPv6 Adjacency SID
    PceSrSid_ipv6_adjacency_sid PceSrSid = "ipv6-adjacency-sid"

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

    pceLspData.EntityData.Children = types.NewOrderedMap()
    pceLspData.EntityData.Children.Append("tunnel-infos", types.YChild{"TunnelInfos", &pceLspData.TunnelInfos})
    pceLspData.EntityData.Children.Append("lsp-summary", types.YChild{"LspSummary", &pceLspData.LspSummary})
    pceLspData.EntityData.Children.Append("tunnel-detail-infos", types.YChild{"TunnelDetailInfos", &pceLspData.TunnelDetailInfos})
    pceLspData.EntityData.Leafs = types.NewOrderedMap()

    pceLspData.EntityData.YListKeys = []string {}

    return &(pceLspData.EntityData)
}

// PceLspData_TunnelInfos
// Tunnel database in XTC
type PceLspData_TunnelInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of PceLspData_TunnelInfos_TunnelInfo.
    TunnelInfo []*PceLspData_TunnelInfos_TunnelInfo
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

    tunnelInfos.EntityData.Children = types.NewOrderedMap()
    tunnelInfos.EntityData.Children.Append("tunnel-info", types.YChild{"TunnelInfo", nil})
    for i := range tunnelInfos.TunnelInfo {
        tunnelInfos.EntityData.Children.Append(types.GetSegmentPath(tunnelInfos.TunnelInfo[i]), types.YChild{"TunnelInfo", tunnelInfos.TunnelInfo[i]})
    }
    tunnelInfos.EntityData.Leafs = types.NewOrderedMap()

    tunnelInfos.EntityData.YListKeys = []string {}

    return &(tunnelInfos.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo
// Tunnel information
type PceLspData_TunnelInfos_TunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // 0..4294967295.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string.
    TunnelName interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // PCC address.
    PccAddress PceLspData_TunnelInfos_TunnelInfo_PccAddress

    // Brief LSP information. The type is slice of
    // PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation.
    BriefLspInformation []*PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation
}

func (tunnelInfo *PceLspData_TunnelInfos_TunnelInfo) GetEntityData() *types.CommonEntityData {
    tunnelInfo.EntityData.YFilter = tunnelInfo.YFilter
    tunnelInfo.EntityData.YangName = "tunnel-info"
    tunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelInfo.EntityData.ParentYangName = "tunnel-infos"
    tunnelInfo.EntityData.SegmentPath = "tunnel-info" + types.AddKeyToken(tunnelInfo.PeerAddress, "peer-address") + types.AddKeyToken(tunnelInfo.PlspId, "plsp-id") + types.AddKeyToken(tunnelInfo.TunnelName, "tunnel-name")
    tunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInfo.EntityData.Children = types.NewOrderedMap()
    tunnelInfo.EntityData.Children.Append("pcc-address", types.YChild{"PccAddress", &tunnelInfo.PccAddress})
    tunnelInfo.EntityData.Children.Append("brief-lsp-information", types.YChild{"BriefLspInformation", nil})
    for i := range tunnelInfo.BriefLspInformation {
        tunnelInfo.EntityData.Children.Append(types.GetSegmentPath(tunnelInfo.BriefLspInformation[i]), types.YChild{"BriefLspInformation", tunnelInfo.BriefLspInformation[i]})
    }
    tunnelInfo.EntityData.Leafs = types.NewOrderedMap()
    tunnelInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", tunnelInfo.PeerAddress})
    tunnelInfo.EntityData.Leafs.Append("plsp-id", types.YLeaf{"PlspId", tunnelInfo.PlspId})
    tunnelInfo.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", tunnelInfo.TunnelName})
    tunnelInfo.EntityData.Leafs.Append("tunnel-name-xr", types.YLeaf{"TunnelNameXr", tunnelInfo.TunnelNameXr})

    tunnelInfo.EntityData.YListKeys = []string {"PeerAddress", "PlspId", "TunnelName"}

    return &(tunnelInfo.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo_PccAddress
// PCC address
type PceLspData_TunnelInfos_TunnelInfo_PccAddress struct {
    EntityData types.CommonEntityData
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

func (pccAddress *PceLspData_TunnelInfos_TunnelInfo_PccAddress) GetEntityData() *types.CommonEntityData {
    pccAddress.EntityData.YFilter = pccAddress.YFilter
    pccAddress.EntityData.YangName = "pcc-address"
    pccAddress.EntityData.BundleName = "cisco_ios_xr"
    pccAddress.EntityData.ParentYangName = "tunnel-info"
    pccAddress.EntityData.SegmentPath = "pcc-address"
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = types.NewOrderedMap()
    pccAddress.EntityData.Leafs = types.NewOrderedMap()
    pccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pccAddress.AfName})
    pccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", pccAddress.Ipv4})
    pccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", pccAddress.Ipv6})

    pccAddress.EntityData.YListKeys = []string {}

    return &(pccAddress.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation
// Brief LSP information
type PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    Msd interface{}

    // Source address.
    SourceAddress PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress

    // Destination address.
    DestinationAddress PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress
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

    briefLspInformation.EntityData.Children = types.NewOrderedMap()
    briefLspInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &briefLspInformation.SourceAddress})
    briefLspInformation.EntityData.Children.Append("destination-address", types.YChild{"DestinationAddress", &briefLspInformation.DestinationAddress})
    briefLspInformation.EntityData.Leafs = types.NewOrderedMap()
    briefLspInformation.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", briefLspInformation.TunnelId})
    briefLspInformation.EntityData.Leafs.Append("lspid", types.YLeaf{"Lspid", briefLspInformation.Lspid})
    briefLspInformation.EntityData.Leafs.Append("binding-sid", types.YLeaf{"BindingSid", briefLspInformation.BindingSid})
    briefLspInformation.EntityData.Leafs.Append("lsp-setup-type", types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType})
    briefLspInformation.EntityData.Leafs.Append("operational-state", types.YLeaf{"OperationalState", briefLspInformation.OperationalState})
    briefLspInformation.EntityData.Leafs.Append("administrative-state", types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState})
    briefLspInformation.EntityData.Leafs.Append("msd", types.YLeaf{"Msd", briefLspInformation.Msd})

    briefLspInformation.EntityData.YListKeys = []string {}

    return &(briefLspInformation.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress
// Source address
type PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress struct {
    EntityData types.CommonEntityData
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

func (sourceAddress *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "brief-lsp-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sourceAddress.AfName})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress
// Destination address
type PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress struct {
    EntityData types.CommonEntityData
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

func (destinationAddress *PceLspData_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress) GetEntityData() *types.CommonEntityData {
    destinationAddress.EntityData.YFilter = destinationAddress.YFilter
    destinationAddress.EntityData.YangName = "destination-address"
    destinationAddress.EntityData.BundleName = "cisco_ios_xr"
    destinationAddress.EntityData.ParentYangName = "brief-lsp-information"
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

// PceLspData_LspSummary
// LSP summary database in XTC
type PceLspData_LspSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summary for all peers.
    AllLsPs PceLspData_LspSummary_AllLsPs

    // Number of LSPs for specific peer. The type is slice of
    // PceLspData_LspSummary_PeerLsPsInfo.
    PeerLsPsInfo []*PceLspData_LspSummary_PeerLsPsInfo
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

    lspSummary.EntityData.Children = types.NewOrderedMap()
    lspSummary.EntityData.Children.Append("all-ls-ps", types.YChild{"AllLsPs", &lspSummary.AllLsPs})
    lspSummary.EntityData.Children.Append("peer-ls-ps-info", types.YChild{"PeerLsPsInfo", nil})
    for i := range lspSummary.PeerLsPsInfo {
        lspSummary.EntityData.Children.Append(types.GetSegmentPath(lspSummary.PeerLsPsInfo[i]), types.YChild{"PeerLsPsInfo", lspSummary.PeerLsPsInfo[i]})
    }
    lspSummary.EntityData.Leafs = types.NewOrderedMap()

    lspSummary.EntityData.YListKeys = []string {}

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

    allLsPs.EntityData.Children = types.NewOrderedMap()
    allLsPs.EntityData.Leafs = types.NewOrderedMap()
    allLsPs.EntityData.Leafs.Append("all-ls-ps", types.YLeaf{"AllLsPs", allLsPs.AllLsPs})
    allLsPs.EntityData.Leafs.Append("up-ls-ps", types.YLeaf{"UpLsPs", allLsPs.UpLsPs})
    allLsPs.EntityData.Leafs.Append("admin-up-ls-ps", types.YLeaf{"AdminUpLsPs", allLsPs.AdminUpLsPs})
    allLsPs.EntityData.Leafs.Append("sr-ls-ps", types.YLeaf{"SrLsPs", allLsPs.SrLsPs})
    allLsPs.EntityData.Leafs.Append("rsvp-ls-ps", types.YLeaf{"RsvpLsPs", allLsPs.RsvpLsPs})

    allLsPs.EntityData.YListKeys = []string {}

    return &(allLsPs.EntityData)
}

// PceLspData_LspSummary_PeerLsPsInfo
// Number of LSPs for specific peer
type PceLspData_LspSummary_PeerLsPsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of LSPs for specific peer.
    LspSummary PceLspData_LspSummary_PeerLsPsInfo_LspSummary

    // Peer address.
    PeerAddress PceLspData_LspSummary_PeerLsPsInfo_PeerAddress
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

    peerLsPsInfo.EntityData.Children = types.NewOrderedMap()
    peerLsPsInfo.EntityData.Children.Append("lsp-summary", types.YChild{"LspSummary", &peerLsPsInfo.LspSummary})
    peerLsPsInfo.EntityData.Children.Append("peer-address", types.YChild{"PeerAddress", &peerLsPsInfo.PeerAddress})
    peerLsPsInfo.EntityData.Leafs = types.NewOrderedMap()

    peerLsPsInfo.EntityData.YListKeys = []string {}

    return &(peerLsPsInfo.EntityData)
}

// PceLspData_LspSummary_PeerLsPsInfo_LspSummary
// Number of LSPs for specific peer
type PceLspData_LspSummary_PeerLsPsInfo_LspSummary struct {
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

func (lspSummary *PceLspData_LspSummary_PeerLsPsInfo_LspSummary) GetEntityData() *types.CommonEntityData {
    lspSummary.EntityData.YFilter = lspSummary.YFilter
    lspSummary.EntityData.YangName = "lsp-summary"
    lspSummary.EntityData.BundleName = "cisco_ios_xr"
    lspSummary.EntityData.ParentYangName = "peer-ls-ps-info"
    lspSummary.EntityData.SegmentPath = "lsp-summary"
    lspSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspSummary.EntityData.Children = types.NewOrderedMap()
    lspSummary.EntityData.Leafs = types.NewOrderedMap()
    lspSummary.EntityData.Leafs.Append("all-ls-ps", types.YLeaf{"AllLsPs", lspSummary.AllLsPs})
    lspSummary.EntityData.Leafs.Append("up-ls-ps", types.YLeaf{"UpLsPs", lspSummary.UpLsPs})
    lspSummary.EntityData.Leafs.Append("admin-up-ls-ps", types.YLeaf{"AdminUpLsPs", lspSummary.AdminUpLsPs})
    lspSummary.EntityData.Leafs.Append("sr-ls-ps", types.YLeaf{"SrLsPs", lspSummary.SrLsPs})
    lspSummary.EntityData.Leafs.Append("rsvp-ls-ps", types.YLeaf{"RsvpLsPs", lspSummary.RsvpLsPs})

    lspSummary.EntityData.YListKeys = []string {}

    return &(lspSummary.EntityData)
}

// PceLspData_LspSummary_PeerLsPsInfo_PeerAddress
// Peer address
type PceLspData_LspSummary_PeerLsPsInfo_PeerAddress struct {
    EntityData types.CommonEntityData
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

func (peerAddress *PceLspData_LspSummary_PeerLsPsInfo_PeerAddress) GetEntityData() *types.CommonEntityData {
    peerAddress.EntityData.YFilter = peerAddress.YFilter
    peerAddress.EntityData.YangName = "peer-address"
    peerAddress.EntityData.BundleName = "cisco_ios_xr"
    peerAddress.EntityData.ParentYangName = "peer-ls-ps-info"
    peerAddress.EntityData.SegmentPath = "peer-address"
    peerAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddress.EntityData.Children = types.NewOrderedMap()
    peerAddress.EntityData.Leafs = types.NewOrderedMap()
    peerAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddress.AfName})
    peerAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerAddress.Ipv4})
    peerAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerAddress.Ipv6})

    peerAddress.EntityData.YListKeys = []string {}

    return &(peerAddress.EntityData)
}

// PceLspData_TunnelDetailInfos
// Detailed tunnel database in XTC
type PceLspData_TunnelDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed tunnel information. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo.
    TunnelDetailInfo []*PceLspData_TunnelDetailInfos_TunnelDetailInfo
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

    tunnelDetailInfos.EntityData.Children = types.NewOrderedMap()
    tunnelDetailInfos.EntityData.Children.Append("tunnel-detail-info", types.YChild{"TunnelDetailInfo", nil})
    for i := range tunnelDetailInfos.TunnelDetailInfo {
        tunnelDetailInfos.EntityData.Children.Append(types.GetSegmentPath(tunnelDetailInfos.TunnelDetailInfo[i]), types.YChild{"TunnelDetailInfo", tunnelDetailInfos.TunnelDetailInfo[i]})
    }
    tunnelDetailInfos.EntityData.Leafs = types.NewOrderedMap()

    tunnelDetailInfos.EntityData.YListKeys = []string {}

    return &(tunnelDetailInfos.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo
// Detailed tunnel information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // 0..4294967295.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string.
    TunnelName interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Allow XTC reoptimizations. The type is bool.
    XtcControlled interface{}

    // Color. The type is interface{} with range: 0..4294967295.
    Color interface{}

    // PCC address.
    PccAddress PceLspData_TunnelDetailInfos_TunnelDetailInfo_PccAddress

    // Private LSP information.
    PrivateLspInformation PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation

    // Detail LSP information. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation.
    DetailLspInformation []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
}

func (tunnelDetailInfo *PceLspData_TunnelDetailInfos_TunnelDetailInfo) GetEntityData() *types.CommonEntityData {
    tunnelDetailInfo.EntityData.YFilter = tunnelDetailInfo.YFilter
    tunnelDetailInfo.EntityData.YangName = "tunnel-detail-info"
    tunnelDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelDetailInfo.EntityData.ParentYangName = "tunnel-detail-infos"
    tunnelDetailInfo.EntityData.SegmentPath = "tunnel-detail-info" + types.AddKeyToken(tunnelDetailInfo.PeerAddress, "peer-address") + types.AddKeyToken(tunnelDetailInfo.PlspId, "plsp-id") + types.AddKeyToken(tunnelDetailInfo.TunnelName, "tunnel-name")
    tunnelDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDetailInfo.EntityData.Children = types.NewOrderedMap()
    tunnelDetailInfo.EntityData.Children.Append("pcc-address", types.YChild{"PccAddress", &tunnelDetailInfo.PccAddress})
    tunnelDetailInfo.EntityData.Children.Append("private-lsp-information", types.YChild{"PrivateLspInformation", &tunnelDetailInfo.PrivateLspInformation})
    tunnelDetailInfo.EntityData.Children.Append("detail-lsp-information", types.YChild{"DetailLspInformation", nil})
    for i := range tunnelDetailInfo.DetailLspInformation {
        tunnelDetailInfo.EntityData.Children.Append(types.GetSegmentPath(tunnelDetailInfo.DetailLspInformation[i]), types.YChild{"DetailLspInformation", tunnelDetailInfo.DetailLspInformation[i]})
    }
    tunnelDetailInfo.EntityData.Leafs = types.NewOrderedMap()
    tunnelDetailInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", tunnelDetailInfo.PeerAddress})
    tunnelDetailInfo.EntityData.Leafs.Append("plsp-id", types.YLeaf{"PlspId", tunnelDetailInfo.PlspId})
    tunnelDetailInfo.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", tunnelDetailInfo.TunnelName})
    tunnelDetailInfo.EntityData.Leafs.Append("tunnel-name-xr", types.YLeaf{"TunnelNameXr", tunnelDetailInfo.TunnelNameXr})
    tunnelDetailInfo.EntityData.Leafs.Append("xtc-controlled", types.YLeaf{"XtcControlled", tunnelDetailInfo.XtcControlled})
    tunnelDetailInfo.EntityData.Leafs.Append("color", types.YLeaf{"Color", tunnelDetailInfo.Color})

    tunnelDetailInfo.EntityData.YListKeys = []string {"PeerAddress", "PlspId", "TunnelName"}

    return &(tunnelDetailInfo.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PccAddress
// PCC address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PccAddress struct {
    EntityData types.CommonEntityData
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

func (pccAddress *PceLspData_TunnelDetailInfos_TunnelDetailInfo_PccAddress) GetEntityData() *types.CommonEntityData {
    pccAddress.EntityData.YFilter = pccAddress.YFilter
    pccAddress.EntityData.YangName = "pcc-address"
    pccAddress.EntityData.BundleName = "cisco_ios_xr"
    pccAddress.EntityData.ParentYangName = "tunnel-detail-info"
    pccAddress.EntityData.SegmentPath = "pcc-address"
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = types.NewOrderedMap()
    pccAddress.EntityData.Leafs = types.NewOrderedMap()
    pccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pccAddress.AfName})
    pccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", pccAddress.Ipv4})
    pccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", pccAddress.Ipv6})

    pccAddress.EntityData.YListKeys = []string {}

    return &(pccAddress.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation
// Private LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Event buffer. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer.
    EventBuffer []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
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

    privateLspInformation.EntityData.Children = types.NewOrderedMap()
    privateLspInformation.EntityData.Children.Append("event-buffer", types.YChild{"EventBuffer", nil})
    for i := range privateLspInformation.EventBuffer {
        privateLspInformation.EntityData.Children.Append(types.GetSegmentPath(privateLspInformation.EventBuffer[i]), types.YChild{"EventBuffer", privateLspInformation.EventBuffer[i]})
    }
    privateLspInformation.EntityData.Leafs = types.NewOrderedMap()

    privateLspInformation.EntityData.YListKeys = []string {}

    return &(privateLspInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
// LSP Event buffer
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event ID in range 1 - 0xFFFFFFFF. 0 is invalid. The type is interface{}
    // with range: 0..4294967295.
    EventId interface{}

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

    eventBuffer.EntityData.Children = types.NewOrderedMap()
    eventBuffer.EntityData.Leafs = types.NewOrderedMap()
    eventBuffer.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", eventBuffer.EventId})
    eventBuffer.EntityData.Leafs.Append("event-message", types.YLeaf{"EventMessage", eventBuffer.EventMessage})
    eventBuffer.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", eventBuffer.TimeStamp})

    eventBuffer.EntityData.YListKeys = []string {}

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

    // Sub delegated PCE.
    SubDelegatedPce PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce

    // State-sync PCE.
    StateSyncPce PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce

    // Reporting PCC address.
    ReportingPccAddress PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress

    // RRO. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro.
    Rro []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
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

    detailLspInformation.EntityData.Children = types.NewOrderedMap()
    detailLspInformation.EntityData.Children.Append("brief-lsp-information", types.YChild{"BriefLspInformation", &detailLspInformation.BriefLspInformation})
    detailLspInformation.EntityData.Children.Append("er-os", types.YChild{"ErOs", &detailLspInformation.ErOs})
    detailLspInformation.EntityData.Children.Append("lsppcep-information", types.YChild{"LsppcepInformation", &detailLspInformation.LsppcepInformation})
    detailLspInformation.EntityData.Children.Append("lsp-association-info", types.YChild{"LspAssociationInfo", &detailLspInformation.LspAssociationInfo})
    detailLspInformation.EntityData.Children.Append("lsp-attributes", types.YChild{"LspAttributes", &detailLspInformation.LspAttributes})
    detailLspInformation.EntityData.Children.Append("sub-delegated-pce", types.YChild{"SubDelegatedPce", &detailLspInformation.SubDelegatedPce})
    detailLspInformation.EntityData.Children.Append("state-sync-pce", types.YChild{"StateSyncPce", &detailLspInformation.StateSyncPce})
    detailLspInformation.EntityData.Children.Append("reporting-pcc-address", types.YChild{"ReportingPccAddress", &detailLspInformation.ReportingPccAddress})
    detailLspInformation.EntityData.Children.Append("rro", types.YChild{"Rro", nil})
    for i := range detailLspInformation.Rro {
        detailLspInformation.EntityData.Children.Append(types.GetSegmentPath(detailLspInformation.Rro[i]), types.YChild{"Rro", detailLspInformation.Rro[i]})
    }
    detailLspInformation.EntityData.Leafs = types.NewOrderedMap()
    detailLspInformation.EntityData.Leafs.Append("signaled-bandwidth-specified", types.YLeaf{"SignaledBandwidthSpecified", detailLspInformation.SignaledBandwidthSpecified})
    detailLspInformation.EntityData.Leafs.Append("signaled-bandwidth", types.YLeaf{"SignaledBandwidth", detailLspInformation.SignaledBandwidth})
    detailLspInformation.EntityData.Leafs.Append("actual-bandwidth-specified", types.YLeaf{"ActualBandwidthSpecified", detailLspInformation.ActualBandwidthSpecified})
    detailLspInformation.EntityData.Leafs.Append("actual-bandwidth", types.YLeaf{"ActualBandwidth", detailLspInformation.ActualBandwidth})
    detailLspInformation.EntityData.Leafs.Append("lsp-role", types.YLeaf{"LspRole", detailLspInformation.LspRole})
    detailLspInformation.EntityData.Leafs.Append("computing-pce", types.YLeaf{"ComputingPce", detailLspInformation.ComputingPce})
    detailLspInformation.EntityData.Leafs.Append("srlg-info", types.YLeaf{"SrlgInfo", detailLspInformation.SrlgInfo})

    detailLspInformation.EntityData.YListKeys = []string {}

    return &(detailLspInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation
// Brief LSP information
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    Msd interface{}

    // Source address.
    SourceAddress PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress

    // Destination address.
    DestinationAddress PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress
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

    briefLspInformation.EntityData.Children = types.NewOrderedMap()
    briefLspInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &briefLspInformation.SourceAddress})
    briefLspInformation.EntityData.Children.Append("destination-address", types.YChild{"DestinationAddress", &briefLspInformation.DestinationAddress})
    briefLspInformation.EntityData.Leafs = types.NewOrderedMap()
    briefLspInformation.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", briefLspInformation.TunnelId})
    briefLspInformation.EntityData.Leafs.Append("lspid", types.YLeaf{"Lspid", briefLspInformation.Lspid})
    briefLspInformation.EntityData.Leafs.Append("binding-sid", types.YLeaf{"BindingSid", briefLspInformation.BindingSid})
    briefLspInformation.EntityData.Leafs.Append("lsp-setup-type", types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType})
    briefLspInformation.EntityData.Leafs.Append("operational-state", types.YLeaf{"OperationalState", briefLspInformation.OperationalState})
    briefLspInformation.EntityData.Leafs.Append("administrative-state", types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState})
    briefLspInformation.EntityData.Leafs.Append("msd", types.YLeaf{"Msd", briefLspInformation.Msd})

    briefLspInformation.EntityData.YListKeys = []string {}

    return &(briefLspInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress
// Source address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress struct {
    EntityData types.CommonEntityData
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

func (sourceAddress *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "brief-lsp-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sourceAddress.AfName})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress
// Destination address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress struct {
    EntityData types.CommonEntityData
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

func (destinationAddress *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress) GetEntityData() *types.CommonEntityData {
    destinationAddress.EntityData.YFilter = destinationAddress.YFilter
    destinationAddress.EntityData.YangName = "destination-address"
    destinationAddress.EntityData.BundleName = "cisco_ios_xr"
    destinationAddress.EntityData.ParentYangName = "brief-lsp-information"
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
    ReportedRsvpPath []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath

    // Reported SR path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath.
    ReportedSrPath []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath

    // Computed RSVP path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath.
    ComputedRsvpPath []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath

    // Computed SR path. The type is slice of
    // PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath.
    ComputedSrPath []*PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
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

    erOs.EntityData.Children = types.NewOrderedMap()
    erOs.EntityData.Children.Append("reported-rsvp-path", types.YChild{"ReportedRsvpPath", nil})
    for i := range erOs.ReportedRsvpPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ReportedRsvpPath[i]), types.YChild{"ReportedRsvpPath", erOs.ReportedRsvpPath[i]})
    }
    erOs.EntityData.Children.Append("reported-sr-path", types.YChild{"ReportedSrPath", nil})
    for i := range erOs.ReportedSrPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ReportedSrPath[i]), types.YChild{"ReportedSrPath", erOs.ReportedSrPath[i]})
    }
    erOs.EntityData.Children.Append("computed-rsvp-path", types.YChild{"ComputedRsvpPath", nil})
    for i := range erOs.ComputedRsvpPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ComputedRsvpPath[i]), types.YChild{"ComputedRsvpPath", erOs.ComputedRsvpPath[i]})
    }
    erOs.EntityData.Children.Append("computed-sr-path", types.YChild{"ComputedSrPath", nil})
    for i := range erOs.ComputedSrPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ComputedSrPath[i]), types.YChild{"ComputedSrPath", erOs.ComputedSrPath[i]})
    }
    erOs.EntityData.Leafs = types.NewOrderedMap()
    erOs.EntityData.Leafs.Append("reported-metric-type", types.YLeaf{"ReportedMetricType", erOs.ReportedMetricType})
    erOs.EntityData.Leafs.Append("reported-metric-value", types.YLeaf{"ReportedMetricValue", erOs.ReportedMetricValue})
    erOs.EntityData.Leafs.Append("computed-metric-type", types.YLeaf{"ComputedMetricType", erOs.ComputedMetricType})
    erOs.EntityData.Leafs.Append("computed-metric-value", types.YLeaf{"ComputedMetricValue", erOs.ComputedMetricValue})
    erOs.EntityData.Leafs.Append("computed-hop-list-time", types.YLeaf{"ComputedHopListTime", erOs.ComputedHopListTime})

    erOs.EntityData.YListKeys = []string {}

    return &(erOs.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath
// Reported RSVP path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    reportedRsvpPath.EntityData.Children = types.NewOrderedMap()
    reportedRsvpPath.EntityData.Leafs = types.NewOrderedMap()
    reportedRsvpPath.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", reportedRsvpPath.HopAddress})

    reportedRsvpPath.EntityData.YListKeys = []string {}

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

    // Local Address.
    LocalAddr PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr

    // Remote Address.
    RemoteAddr PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr
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

    reportedSrPath.EntityData.Children = types.NewOrderedMap()
    reportedSrPath.EntityData.Children.Append("local-addr", types.YChild{"LocalAddr", &reportedSrPath.LocalAddr})
    reportedSrPath.EntityData.Children.Append("remote-addr", types.YChild{"RemoteAddr", &reportedSrPath.RemoteAddr})
    reportedSrPath.EntityData.Leafs = types.NewOrderedMap()
    reportedSrPath.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", reportedSrPath.SidType})
    reportedSrPath.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", reportedSrPath.MplsLabel})

    reportedSrPath.EntityData.YListKeys = []string {}

    return &(reportedSrPath.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr
// Local Address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr struct {
    EntityData types.CommonEntityData
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

func (localAddr *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr) GetEntityData() *types.CommonEntityData {
    localAddr.EntityData.YFilter = localAddr.YFilter
    localAddr.EntityData.YangName = "local-addr"
    localAddr.EntityData.BundleName = "cisco_ios_xr"
    localAddr.EntityData.ParentYangName = "reported-sr-path"
    localAddr.EntityData.SegmentPath = "local-addr"
    localAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddr.EntityData.Children = types.NewOrderedMap()
    localAddr.EntityData.Leafs = types.NewOrderedMap()
    localAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddr.AfName})
    localAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddr.Ipv4})
    localAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddr.Ipv6})

    localAddr.EntityData.YListKeys = []string {}

    return &(localAddr.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr
// Remote Address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr struct {
    EntityData types.CommonEntityData
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

func (remoteAddr *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr) GetEntityData() *types.CommonEntityData {
    remoteAddr.EntityData.YFilter = remoteAddr.YFilter
    remoteAddr.EntityData.YangName = "remote-addr"
    remoteAddr.EntityData.BundleName = "cisco_ios_xr"
    remoteAddr.EntityData.ParentYangName = "reported-sr-path"
    remoteAddr.EntityData.SegmentPath = "remote-addr"
    remoteAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddr.EntityData.Children = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddr.AfName})
    remoteAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddr.Ipv4})
    remoteAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddr.Ipv6})

    remoteAddr.EntityData.YListKeys = []string {}

    return &(remoteAddr.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath
// Computed RSVP path
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    computedRsvpPath.EntityData.Children = types.NewOrderedMap()
    computedRsvpPath.EntityData.Leafs = types.NewOrderedMap()
    computedRsvpPath.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", computedRsvpPath.HopAddress})

    computedRsvpPath.EntityData.YListKeys = []string {}

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

    // Local Address.
    LocalAddr PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr

    // Remote Address.
    RemoteAddr PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr
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

    computedSrPath.EntityData.Children = types.NewOrderedMap()
    computedSrPath.EntityData.Children.Append("local-addr", types.YChild{"LocalAddr", &computedSrPath.LocalAddr})
    computedSrPath.EntityData.Children.Append("remote-addr", types.YChild{"RemoteAddr", &computedSrPath.RemoteAddr})
    computedSrPath.EntityData.Leafs = types.NewOrderedMap()
    computedSrPath.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", computedSrPath.SidType})
    computedSrPath.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", computedSrPath.MplsLabel})

    computedSrPath.EntityData.YListKeys = []string {}

    return &(computedSrPath.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr
// Local Address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr struct {
    EntityData types.CommonEntityData
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

func (localAddr *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr) GetEntityData() *types.CommonEntityData {
    localAddr.EntityData.YFilter = localAddr.YFilter
    localAddr.EntityData.YangName = "local-addr"
    localAddr.EntityData.BundleName = "cisco_ios_xr"
    localAddr.EntityData.ParentYangName = "computed-sr-path"
    localAddr.EntityData.SegmentPath = "local-addr"
    localAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddr.EntityData.Children = types.NewOrderedMap()
    localAddr.EntityData.Leafs = types.NewOrderedMap()
    localAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddr.AfName})
    localAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddr.Ipv4})
    localAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddr.Ipv6})

    localAddr.EntityData.YListKeys = []string {}

    return &(localAddr.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr
// Remote Address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr struct {
    EntityData types.CommonEntityData
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

func (remoteAddr *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr) GetEntityData() *types.CommonEntityData {
    remoteAddr.EntityData.YFilter = remoteAddr.YFilter
    remoteAddr.EntityData.YangName = "remote-addr"
    remoteAddr.EntityData.BundleName = "cisco_ios_xr"
    remoteAddr.EntityData.ParentYangName = "computed-sr-path"
    remoteAddr.EntityData.SegmentPath = "remote-addr"
    remoteAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddr.EntityData.Children = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddr.AfName})
    remoteAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddr.Ipv4})
    remoteAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddr.Ipv6})

    remoteAddr.EntityData.YListKeys = []string {}

    return &(remoteAddr.EntityData)
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

    lsppcepInformation.EntityData.Children = types.NewOrderedMap()
    lsppcepInformation.EntityData.Children.Append("rsvp-error", types.YChild{"RsvpError", &lsppcepInformation.RsvpError})
    lsppcepInformation.EntityData.Leafs = types.NewOrderedMap()
    lsppcepInformation.EntityData.Leafs.Append("pcepid", types.YLeaf{"Pcepid", lsppcepInformation.Pcepid})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-d", types.YLeaf{"PcepFlagD", lsppcepInformation.PcepFlagD})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-s", types.YLeaf{"PcepFlagS", lsppcepInformation.PcepFlagS})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-r", types.YLeaf{"PcepFlagR", lsppcepInformation.PcepFlagR})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-a", types.YLeaf{"PcepFlagA", lsppcepInformation.PcepFlagA})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-o", types.YLeaf{"PcepFlagO", lsppcepInformation.PcepFlagO})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-c", types.YLeaf{"PcepFlagC", lsppcepInformation.PcepFlagC})

    lsppcepInformation.EntityData.YListKeys = []string {}

    return &(lsppcepInformation.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
// RSVP error info
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError struct {
    EntityData types.CommonEntityData
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

func (rsvpError *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetEntityData() *types.CommonEntityData {
    rsvpError.EntityData.YFilter = rsvpError.YFilter
    rsvpError.EntityData.YangName = "rsvp-error"
    rsvpError.EntityData.BundleName = "cisco_ios_xr"
    rsvpError.EntityData.ParentYangName = "lsppcep-information"
    rsvpError.EntityData.SegmentPath = "rsvp-error"
    rsvpError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvpError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvpError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvpError.EntityData.Children = types.NewOrderedMap()
    rsvpError.EntityData.Leafs = types.NewOrderedMap()
    rsvpError.EntityData.Leafs.Append("node-address", types.YLeaf{"NodeAddress", rsvpError.NodeAddress})
    rsvpError.EntityData.Leafs.Append("error-flags", types.YLeaf{"ErrorFlags", rsvpError.ErrorFlags})
    rsvpError.EntityData.Leafs.Append("error-code", types.YLeaf{"ErrorCode", rsvpError.ErrorCode})
    rsvpError.EntityData.Leafs.Append("error-value", types.YLeaf{"ErrorValue", rsvpError.ErrorValue})

    rsvpError.EntityData.YListKeys = []string {}

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

    // Association Source.
    AssociationSource PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource
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

    lspAssociationInfo.EntityData.Children = types.NewOrderedMap()
    lspAssociationInfo.EntityData.Children.Append("association-source", types.YChild{"AssociationSource", &lspAssociationInfo.AssociationSource})
    lspAssociationInfo.EntityData.Leafs = types.NewOrderedMap()
    lspAssociationInfo.EntityData.Leafs.Append("association-type", types.YLeaf{"AssociationType", lspAssociationInfo.AssociationType})
    lspAssociationInfo.EntityData.Leafs.Append("association-id", types.YLeaf{"AssociationId", lspAssociationInfo.AssociationId})

    lspAssociationInfo.EntityData.YListKeys = []string {}

    return &(lspAssociationInfo.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource
// Association Source
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource struct {
    EntityData types.CommonEntityData
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

func (associationSource *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource) GetEntityData() *types.CommonEntityData {
    associationSource.EntityData.YFilter = associationSource.YFilter
    associationSource.EntityData.YangName = "association-source"
    associationSource.EntityData.BundleName = "cisco_ios_xr"
    associationSource.EntityData.ParentYangName = "lsp-association-info"
    associationSource.EntityData.SegmentPath = "association-source"
    associationSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationSource.EntityData.Children = types.NewOrderedMap()
    associationSource.EntityData.Leafs = types.NewOrderedMap()
    associationSource.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", associationSource.AfName})
    associationSource.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", associationSource.Ipv4})
    associationSource.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", associationSource.Ipv6})

    associationSource.EntityData.YListKeys = []string {}

    return &(associationSource.EntityData)
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

    lspAttributes.EntityData.Children = types.NewOrderedMap()
    lspAttributes.EntityData.Leafs = types.NewOrderedMap()
    lspAttributes.EntityData.Leafs.Append("affinity-exclude-any", types.YLeaf{"AffinityExcludeAny", lspAttributes.AffinityExcludeAny})
    lspAttributes.EntityData.Leafs.Append("affinity-include-any", types.YLeaf{"AffinityIncludeAny", lspAttributes.AffinityIncludeAny})
    lspAttributes.EntityData.Leafs.Append("affinity-include-all", types.YLeaf{"AffinityIncludeAll", lspAttributes.AffinityIncludeAll})
    lspAttributes.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", lspAttributes.SetupPriority})
    lspAttributes.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", lspAttributes.HoldPriority})
    lspAttributes.EntityData.Leafs.Append("local-protection", types.YLeaf{"LocalProtection", lspAttributes.LocalProtection})

    lspAttributes.EntityData.YListKeys = []string {}

    return &(lspAttributes.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce
// Sub delegated PCE
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce struct {
    EntityData types.CommonEntityData
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

func (subDelegatedPce *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce) GetEntityData() *types.CommonEntityData {
    subDelegatedPce.EntityData.YFilter = subDelegatedPce.YFilter
    subDelegatedPce.EntityData.YangName = "sub-delegated-pce"
    subDelegatedPce.EntityData.BundleName = "cisco_ios_xr"
    subDelegatedPce.EntityData.ParentYangName = "detail-lsp-information"
    subDelegatedPce.EntityData.SegmentPath = "sub-delegated-pce"
    subDelegatedPce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subDelegatedPce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subDelegatedPce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subDelegatedPce.EntityData.Children = types.NewOrderedMap()
    subDelegatedPce.EntityData.Leafs = types.NewOrderedMap()
    subDelegatedPce.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", subDelegatedPce.AfName})
    subDelegatedPce.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", subDelegatedPce.Ipv4})
    subDelegatedPce.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", subDelegatedPce.Ipv6})

    subDelegatedPce.EntityData.YListKeys = []string {}

    return &(subDelegatedPce.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce
// State-sync PCE
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce struct {
    EntityData types.CommonEntityData
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

func (stateSyncPce *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce) GetEntityData() *types.CommonEntityData {
    stateSyncPce.EntityData.YFilter = stateSyncPce.YFilter
    stateSyncPce.EntityData.YangName = "state-sync-pce"
    stateSyncPce.EntityData.BundleName = "cisco_ios_xr"
    stateSyncPce.EntityData.ParentYangName = "detail-lsp-information"
    stateSyncPce.EntityData.SegmentPath = "state-sync-pce"
    stateSyncPce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSyncPce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSyncPce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSyncPce.EntityData.Children = types.NewOrderedMap()
    stateSyncPce.EntityData.Leafs = types.NewOrderedMap()
    stateSyncPce.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", stateSyncPce.AfName})
    stateSyncPce.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", stateSyncPce.Ipv4})
    stateSyncPce.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", stateSyncPce.Ipv6})

    stateSyncPce.EntityData.YListKeys = []string {}

    return &(stateSyncPce.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress
// Reporting PCC address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress struct {
    EntityData types.CommonEntityData
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

func (reportingPccAddress *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress) GetEntityData() *types.CommonEntityData {
    reportingPccAddress.EntityData.YFilter = reportingPccAddress.YFilter
    reportingPccAddress.EntityData.YangName = "reporting-pcc-address"
    reportingPccAddress.EntityData.BundleName = "cisco_ios_xr"
    reportingPccAddress.EntityData.ParentYangName = "detail-lsp-information"
    reportingPccAddress.EntityData.SegmentPath = "reporting-pcc-address"
    reportingPccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportingPccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportingPccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportingPccAddress.EntityData.Children = types.NewOrderedMap()
    reportingPccAddress.EntityData.Leafs = types.NewOrderedMap()
    reportingPccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", reportingPccAddress.AfName})
    reportingPccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", reportingPccAddress.Ipv4})
    reportingPccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", reportingPccAddress.Ipv6})

    reportingPccAddress.EntityData.YListKeys = []string {}

    return &(reportingPccAddress.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
// RRO
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro struct {
    EntityData types.CommonEntityData
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

func (rro *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetEntityData() *types.CommonEntityData {
    rro.EntityData.YFilter = rro.YFilter
    rro.EntityData.YangName = "rro"
    rro.EntityData.BundleName = "cisco_ios_xr"
    rro.EntityData.ParentYangName = "detail-lsp-information"
    rro.EntityData.SegmentPath = "rro"
    rro.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rro.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rro.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rro.EntityData.Children = types.NewOrderedMap()
    rro.EntityData.Children.Append("sr-rro", types.YChild{"SrRro", &rro.SrRro})
    rro.EntityData.Leafs = types.NewOrderedMap()
    rro.EntityData.Leafs.Append("rro-type", types.YLeaf{"RroType", rro.RroType})
    rro.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", rro.Ipv4Address})
    rro.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", rro.MplsLabel})
    rro.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", rro.Flags})

    rro.EntityData.YListKeys = []string {}

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

    // Local Address.
    LocalAddr PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr

    // Remote Address.
    RemoteAddr PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr
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

    srRro.EntityData.Children = types.NewOrderedMap()
    srRro.EntityData.Children.Append("local-addr", types.YChild{"LocalAddr", &srRro.LocalAddr})
    srRro.EntityData.Children.Append("remote-addr", types.YChild{"RemoteAddr", &srRro.RemoteAddr})
    srRro.EntityData.Leafs = types.NewOrderedMap()
    srRro.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", srRro.SidType})
    srRro.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", srRro.MplsLabel})

    srRro.EntityData.YListKeys = []string {}

    return &(srRro.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr
// Local Address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr struct {
    EntityData types.CommonEntityData
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

func (localAddr *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr) GetEntityData() *types.CommonEntityData {
    localAddr.EntityData.YFilter = localAddr.YFilter
    localAddr.EntityData.YangName = "local-addr"
    localAddr.EntityData.BundleName = "cisco_ios_xr"
    localAddr.EntityData.ParentYangName = "sr-rro"
    localAddr.EntityData.SegmentPath = "local-addr"
    localAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddr.EntityData.Children = types.NewOrderedMap()
    localAddr.EntityData.Leafs = types.NewOrderedMap()
    localAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddr.AfName})
    localAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddr.Ipv4})
    localAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddr.Ipv6})

    localAddr.EntityData.YListKeys = []string {}

    return &(localAddr.EntityData)
}

// PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr
// Remote Address
type PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr struct {
    EntityData types.CommonEntityData
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

func (remoteAddr *PceLspData_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr) GetEntityData() *types.CommonEntityData {
    remoteAddr.EntityData.YFilter = remoteAddr.YFilter
    remoteAddr.EntityData.YangName = "remote-addr"
    remoteAddr.EntityData.BundleName = "cisco_ios_xr"
    remoteAddr.EntityData.ParentYangName = "sr-rro"
    remoteAddr.EntityData.SegmentPath = "remote-addr"
    remoteAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddr.EntityData.Children = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddr.AfName})
    remoteAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddr.Ipv4})
    remoteAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddr.Ipv6})

    remoteAddr.EntityData.YListKeys = []string {}

    return &(remoteAddr.EntityData)
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

    pcePeer.EntityData.Children = types.NewOrderedMap()
    pcePeer.EntityData.Children.Append("peer-detail-infos", types.YChild{"PeerDetailInfos", &pcePeer.PeerDetailInfos})
    pcePeer.EntityData.Children.Append("peer-infos", types.YChild{"PeerInfos", &pcePeer.PeerInfos})
    pcePeer.EntityData.Leafs = types.NewOrderedMap()

    pcePeer.EntityData.YListKeys = []string {}

    return &(pcePeer.EntityData)
}

// PcePeer_PeerDetailInfos
// Detailed peers database in XTC
type PcePeer_PeerDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed PCE peer information. The type is slice of
    // PcePeer_PeerDetailInfos_PeerDetailInfo.
    PeerDetailInfo []*PcePeer_PeerDetailInfos_PeerDetailInfo
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

    peerDetailInfos.EntityData.Children = types.NewOrderedMap()
    peerDetailInfos.EntityData.Children.Append("peer-detail-info", types.YChild{"PeerDetailInfo", nil})
    for i := range peerDetailInfos.PeerDetailInfo {
        peerDetailInfos.EntityData.Children.Append(types.GetSegmentPath(peerDetailInfos.PeerDetailInfo[i]), types.YChild{"PeerDetailInfo", peerDetailInfos.PeerDetailInfo[i]})
    }
    peerDetailInfos.EntityData.Leafs = types.NewOrderedMap()

    peerDetailInfos.EntityData.YListKeys = []string {}

    return &(peerDetailInfos.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo
// Detailed PCE peer information
type PcePeer_PeerDetailInfos_PeerDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    MaxSidDepth interface{}

    // Peer address.
    PeerAddressXr PcePeer_PeerDetailInfos_PeerDetailInfo_PeerAddressXr

    // Detailed PCE protocol information.
    DetailPcepInformation PcePeer_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
}

func (peerDetailInfo *PcePeer_PeerDetailInfos_PeerDetailInfo) GetEntityData() *types.CommonEntityData {
    peerDetailInfo.EntityData.YFilter = peerDetailInfo.YFilter
    peerDetailInfo.EntityData.YangName = "peer-detail-info"
    peerDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfo.EntityData.ParentYangName = "peer-detail-infos"
    peerDetailInfo.EntityData.SegmentPath = "peer-detail-info" + types.AddKeyToken(peerDetailInfo.PeerAddress, "peer-address")
    peerDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfo.EntityData.Children = types.NewOrderedMap()
    peerDetailInfo.EntityData.Children.Append("peer-address-xr", types.YChild{"PeerAddressXr", &peerDetailInfo.PeerAddressXr})
    peerDetailInfo.EntityData.Children.Append("detail-pcep-information", types.YChild{"DetailPcepInformation", &peerDetailInfo.DetailPcepInformation})
    peerDetailInfo.EntityData.Leafs = types.NewOrderedMap()
    peerDetailInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", peerDetailInfo.PeerAddress})
    peerDetailInfo.EntityData.Leafs.Append("peer-protocol", types.YLeaf{"PeerProtocol", peerDetailInfo.PeerProtocol})
    peerDetailInfo.EntityData.Leafs.Append("max-sid-depth", types.YLeaf{"MaxSidDepth", peerDetailInfo.MaxSidDepth})

    peerDetailInfo.EntityData.YListKeys = []string {"PeerAddress"}

    return &(peerDetailInfo.EntityData)
}

// PcePeer_PeerDetailInfos_PeerDetailInfo_PeerAddressXr
// Peer address
type PcePeer_PeerDetailInfos_PeerDetailInfo_PeerAddressXr struct {
    EntityData types.CommonEntityData
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

func (peerAddressXr *PcePeer_PeerDetailInfos_PeerDetailInfo_PeerAddressXr) GetEntityData() *types.CommonEntityData {
    peerAddressXr.EntityData.YFilter = peerAddressXr.YFilter
    peerAddressXr.EntityData.YangName = "peer-address-xr"
    peerAddressXr.EntityData.BundleName = "cisco_ios_xr"
    peerAddressXr.EntityData.ParentYangName = "peer-detail-info"
    peerAddressXr.EntityData.SegmentPath = "peer-address-xr"
    peerAddressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddressXr.EntityData.Children = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddressXr.AfName})
    peerAddressXr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerAddressXr.Ipv4})
    peerAddressXr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerAddressXr.Ipv6})

    peerAddressXr.EntityData.YListKeys = []string {}

    return &(peerAddressXr.EntityData)
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

    // Maximum number of labels the peer can impose. The type is interface{} with
    // range: 0..255.
    MaxSidDepth interface{}

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

    detailPcepInformation.EntityData.Children = types.NewOrderedMap()
    detailPcepInformation.EntityData.Children.Append("brief-pcep-information", types.YChild{"BriefPcepInformation", &detailPcepInformation.BriefPcepInformation})
    detailPcepInformation.EntityData.Children.Append("last-error-rx", types.YChild{"LastErrorRx", &detailPcepInformation.LastErrorRx})
    detailPcepInformation.EntityData.Children.Append("last-error-tx", types.YChild{"LastErrorTx", &detailPcepInformation.LastErrorTx})
    detailPcepInformation.EntityData.Leafs = types.NewOrderedMap()
    detailPcepInformation.EntityData.Leafs.Append("error", types.YLeaf{"Error", detailPcepInformation.Error})
    detailPcepInformation.EntityData.Leafs.Append("speaker-id", types.YLeaf{"SpeakerId", detailPcepInformation.SpeakerId})
    detailPcepInformation.EntityData.Leafs.Append("pcep-up-time", types.YLeaf{"PcepUpTime", detailPcepInformation.PcepUpTime})
    detailPcepInformation.EntityData.Leafs.Append("keepalives", types.YLeaf{"Keepalives", detailPcepInformation.Keepalives})
    detailPcepInformation.EntityData.Leafs.Append("md5-enabled", types.YLeaf{"Md5Enabled", detailPcepInformation.Md5Enabled})
    detailPcepInformation.EntityData.Leafs.Append("keychain-enabled", types.YLeaf{"KeychainEnabled", detailPcepInformation.KeychainEnabled})
    detailPcepInformation.EntityData.Leafs.Append("negotiated-local-keepalive", types.YLeaf{"NegotiatedLocalKeepalive", detailPcepInformation.NegotiatedLocalKeepalive})
    detailPcepInformation.EntityData.Leafs.Append("negotiated-remote-keepalive", types.YLeaf{"NegotiatedRemoteKeepalive", detailPcepInformation.NegotiatedRemoteKeepalive})
    detailPcepInformation.EntityData.Leafs.Append("negotiated-dead-time", types.YLeaf{"NegotiatedDeadTime", detailPcepInformation.NegotiatedDeadTime})
    detailPcepInformation.EntityData.Leafs.Append("pce-request-rx", types.YLeaf{"PceRequestRx", detailPcepInformation.PceRequestRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-request-tx", types.YLeaf{"PceRequestTx", detailPcepInformation.PceRequestTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-reply-rx", types.YLeaf{"PceReplyRx", detailPcepInformation.PceReplyRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-reply-tx", types.YLeaf{"PceReplyTx", detailPcepInformation.PceReplyTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-error-rx", types.YLeaf{"PceErrorRx", detailPcepInformation.PceErrorRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-error-tx", types.YLeaf{"PceErrorTx", detailPcepInformation.PceErrorTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-open-tx", types.YLeaf{"PceOpenTx", detailPcepInformation.PceOpenTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-open-rx", types.YLeaf{"PceOpenRx", detailPcepInformation.PceOpenRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-report-rx", types.YLeaf{"PceReportRx", detailPcepInformation.PceReportRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-report-tx", types.YLeaf{"PceReportTx", detailPcepInformation.PceReportTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-update-rx", types.YLeaf{"PceUpdateRx", detailPcepInformation.PceUpdateRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-update-tx", types.YLeaf{"PceUpdateTx", detailPcepInformation.PceUpdateTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-initiate-rx", types.YLeaf{"PceInitiateRx", detailPcepInformation.PceInitiateRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-initiate-tx", types.YLeaf{"PceInitiateTx", detailPcepInformation.PceInitiateTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-keepalive-tx", types.YLeaf{"PceKeepaliveTx", detailPcepInformation.PceKeepaliveTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-keepalive-rx", types.YLeaf{"PceKeepaliveRx", detailPcepInformation.PceKeepaliveRx})
    detailPcepInformation.EntityData.Leafs.Append("local-session-id", types.YLeaf{"LocalSessionId", detailPcepInformation.LocalSessionId})
    detailPcepInformation.EntityData.Leafs.Append("remote-session-id", types.YLeaf{"RemoteSessionId", detailPcepInformation.RemoteSessionId})
    detailPcepInformation.EntityData.Leafs.Append("minimum-keepalive-interval", types.YLeaf{"MinimumKeepaliveInterval", detailPcepInformation.MinimumKeepaliveInterval})
    detailPcepInformation.EntityData.Leafs.Append("maximum-dead-interval", types.YLeaf{"MaximumDeadInterval", detailPcepInformation.MaximumDeadInterval})
    detailPcepInformation.EntityData.Leafs.Append("max-sid-depth", types.YLeaf{"MaxSidDepth", detailPcepInformation.MaxSidDepth})

    detailPcepInformation.EntityData.YListKeys = []string {}

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

    briefPcepInformation.EntityData.Children = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs.Append("pcep-state", types.YLeaf{"PcepState", briefPcepInformation.PcepState})
    briefPcepInformation.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", briefPcepInformation.Stateful})
    briefPcepInformation.EntityData.Leafs.Append("capability-update", types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate})
    briefPcepInformation.EntityData.Leafs.Append("capability-instantiate", types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate})
    briefPcepInformation.EntityData.Leafs.Append("capability-segment-routing", types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting})
    briefPcepInformation.EntityData.Leafs.Append("capability-triggered-sync", types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync})
    briefPcepInformation.EntityData.Leafs.Append("capability-db-version", types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion})
    briefPcepInformation.EntityData.Leafs.Append("capability-delta-sync", types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync})

    briefPcepInformation.EntityData.YListKeys = []string {}

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

    lastErrorRx.EntityData.Children = types.NewOrderedMap()
    lastErrorRx.EntityData.Leafs = types.NewOrderedMap()
    lastErrorRx.EntityData.Leafs.Append("pc-error-type", types.YLeaf{"PcErrorType", lastErrorRx.PcErrorType})
    lastErrorRx.EntityData.Leafs.Append("pc-error-value", types.YLeaf{"PcErrorValue", lastErrorRx.PcErrorValue})

    lastErrorRx.EntityData.YListKeys = []string {}

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

    lastErrorTx.EntityData.Children = types.NewOrderedMap()
    lastErrorTx.EntityData.Leafs = types.NewOrderedMap()
    lastErrorTx.EntityData.Leafs.Append("pc-error-type", types.YLeaf{"PcErrorType", lastErrorTx.PcErrorType})
    lastErrorTx.EntityData.Leafs.Append("pc-error-value", types.YLeaf{"PcErrorValue", lastErrorTx.PcErrorValue})

    lastErrorTx.EntityData.YListKeys = []string {}

    return &(lastErrorTx.EntityData)
}

// PcePeer_PeerInfos
// Peers database in XTC
type PcePeer_PeerInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer information. The type is slice of PcePeer_PeerInfos_PeerInfo.
    PeerInfo []*PcePeer_PeerInfos_PeerInfo
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

    peerInfos.EntityData.Children = types.NewOrderedMap()
    peerInfos.EntityData.Children.Append("peer-info", types.YChild{"PeerInfo", nil})
    for i := range peerInfos.PeerInfo {
        peerInfos.EntityData.Children.Append(types.GetSegmentPath(peerInfos.PeerInfo[i]), types.YChild{"PeerInfo", peerInfos.PeerInfo[i]})
    }
    peerInfos.EntityData.Leafs = types.NewOrderedMap()

    peerInfos.EntityData.YListKeys = []string {}

    return &(peerInfos.EntityData)
}

// PcePeer_PeerInfos_PeerInfo
// PCE peer information
type PcePeer_PeerInfos_PeerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Peer address.
    PeerAddressXr PcePeer_PeerInfos_PeerInfo_PeerAddressXr

    // PCE protocol information.
    BriefPcepInformation PcePeer_PeerInfos_PeerInfo_BriefPcepInformation
}

func (peerInfo *PcePeer_PeerInfos_PeerInfo) GetEntityData() *types.CommonEntityData {
    peerInfo.EntityData.YFilter = peerInfo.YFilter
    peerInfo.EntityData.YangName = "peer-info"
    peerInfo.EntityData.BundleName = "cisco_ios_xr"
    peerInfo.EntityData.ParentYangName = "peer-infos"
    peerInfo.EntityData.SegmentPath = "peer-info" + types.AddKeyToken(peerInfo.PeerAddress, "peer-address")
    peerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfo.EntityData.Children = types.NewOrderedMap()
    peerInfo.EntityData.Children.Append("peer-address-xr", types.YChild{"PeerAddressXr", &peerInfo.PeerAddressXr})
    peerInfo.EntityData.Children.Append("brief-pcep-information", types.YChild{"BriefPcepInformation", &peerInfo.BriefPcepInformation})
    peerInfo.EntityData.Leafs = types.NewOrderedMap()
    peerInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", peerInfo.PeerAddress})
    peerInfo.EntityData.Leafs.Append("peer-protocol", types.YLeaf{"PeerProtocol", peerInfo.PeerProtocol})

    peerInfo.EntityData.YListKeys = []string {"PeerAddress"}

    return &(peerInfo.EntityData)
}

// PcePeer_PeerInfos_PeerInfo_PeerAddressXr
// Peer address
type PcePeer_PeerInfos_PeerInfo_PeerAddressXr struct {
    EntityData types.CommonEntityData
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

func (peerAddressXr *PcePeer_PeerInfos_PeerInfo_PeerAddressXr) GetEntityData() *types.CommonEntityData {
    peerAddressXr.EntityData.YFilter = peerAddressXr.YFilter
    peerAddressXr.EntityData.YangName = "peer-address-xr"
    peerAddressXr.EntityData.BundleName = "cisco_ios_xr"
    peerAddressXr.EntityData.ParentYangName = "peer-info"
    peerAddressXr.EntityData.SegmentPath = "peer-address-xr"
    peerAddressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddressXr.EntityData.Children = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddressXr.AfName})
    peerAddressXr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerAddressXr.Ipv4})
    peerAddressXr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerAddressXr.Ipv6})

    peerAddressXr.EntityData.YListKeys = []string {}

    return &(peerAddressXr.EntityData)
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

    briefPcepInformation.EntityData.Children = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs.Append("pcep-state", types.YLeaf{"PcepState", briefPcepInformation.PcepState})
    briefPcepInformation.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", briefPcepInformation.Stateful})
    briefPcepInformation.EntityData.Leafs.Append("capability-update", types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate})
    briefPcepInformation.EntityData.Leafs.Append("capability-instantiate", types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate})
    briefPcepInformation.EntityData.Leafs.Append("capability-segment-routing", types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting})
    briefPcepInformation.EntityData.Leafs.Append("capability-triggered-sync", types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync})
    briefPcepInformation.EntityData.Leafs.Append("capability-db-version", types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion})
    briefPcepInformation.EntityData.Leafs.Append("capability-delta-sync", types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync})

    briefPcepInformation.EntityData.YListKeys = []string {}

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

    pceTopology.EntityData.Children = types.NewOrderedMap()
    pceTopology.EntityData.Children.Append("topology-summary", types.YChild{"TopologySummary", &pceTopology.TopologySummary})
    pceTopology.EntityData.Children.Append("topology-nodes", types.YChild{"TopologyNodes", &pceTopology.TopologyNodes})
    pceTopology.EntityData.Children.Append("prefix-infos", types.YChild{"PrefixInfos", &pceTopology.PrefixInfos})
    pceTopology.EntityData.Leafs = types.NewOrderedMap()

    pceTopology.EntityData.YListKeys = []string {}

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

    topologySummary.EntityData.Children = types.NewOrderedMap()
    topologySummary.EntityData.Children.Append("stats-topology-update", types.YChild{"StatsTopologyUpdate", &topologySummary.StatsTopologyUpdate})
    topologySummary.EntityData.Leafs = types.NewOrderedMap()
    topologySummary.EntityData.Leafs.Append("nodes", types.YLeaf{"Nodes", topologySummary.Nodes})
    topologySummary.EntityData.Leafs.Append("lookup-nodes", types.YLeaf{"LookupNodes", topologySummary.LookupNodes})
    topologySummary.EntityData.Leafs.Append("prefixes", types.YLeaf{"Prefixes", topologySummary.Prefixes})
    topologySummary.EntityData.Leafs.Append("prefix-sids", types.YLeaf{"PrefixSids", topologySummary.PrefixSids})
    topologySummary.EntityData.Leafs.Append("regular-prefix-sids", types.YLeaf{"RegularPrefixSids", topologySummary.RegularPrefixSids})
    topologySummary.EntityData.Leafs.Append("strict-prefix-sids", types.YLeaf{"StrictPrefixSids", topologySummary.StrictPrefixSids})
    topologySummary.EntityData.Leafs.Append("links", types.YLeaf{"Links", topologySummary.Links})
    topologySummary.EntityData.Leafs.Append("epe-links", types.YLeaf{"EpeLinks", topologySummary.EpeLinks})
    topologySummary.EntityData.Leafs.Append("adjacency-sids", types.YLeaf{"AdjacencySids", topologySummary.AdjacencySids})
    topologySummary.EntityData.Leafs.Append("epesids", types.YLeaf{"Epesids", topologySummary.Epesids})
    topologySummary.EntityData.Leafs.Append("protected-adjacency-sids", types.YLeaf{"ProtectedAdjacencySids", topologySummary.ProtectedAdjacencySids})
    topologySummary.EntityData.Leafs.Append("un-protected-adjacency-sids", types.YLeaf{"UnProtectedAdjacencySids", topologySummary.UnProtectedAdjacencySids})
    topologySummary.EntityData.Leafs.Append("topology-consistent", types.YLeaf{"TopologyConsistent", topologySummary.TopologyConsistent})

    topologySummary.EntityData.YListKeys = []string {}

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

    statsTopologyUpdate.EntityData.Children = types.NewOrderedMap()
    statsTopologyUpdate.EntityData.Leafs = types.NewOrderedMap()
    statsTopologyUpdate.EntityData.Leafs.Append("num-nodes-added", types.YLeaf{"NumNodesAdded", statsTopologyUpdate.NumNodesAdded})
    statsTopologyUpdate.EntityData.Leafs.Append("num-nodes-deleted", types.YLeaf{"NumNodesDeleted", statsTopologyUpdate.NumNodesDeleted})
    statsTopologyUpdate.EntityData.Leafs.Append("num-links-added", types.YLeaf{"NumLinksAdded", statsTopologyUpdate.NumLinksAdded})
    statsTopologyUpdate.EntityData.Leafs.Append("num-links-deleted", types.YLeaf{"NumLinksDeleted", statsTopologyUpdate.NumLinksDeleted})
    statsTopologyUpdate.EntityData.Leafs.Append("num-prefixes-added", types.YLeaf{"NumPrefixesAdded", statsTopologyUpdate.NumPrefixesAdded})
    statsTopologyUpdate.EntityData.Leafs.Append("num-prefixes-deleted", types.YLeaf{"NumPrefixesDeleted", statsTopologyUpdate.NumPrefixesDeleted})

    statsTopologyUpdate.EntityData.YListKeys = []string {}

    return &(statsTopologyUpdate.EntityData)
}

// PceTopology_TopologyNodes
// Node database in XTC
type PceTopology_TopologyNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode.
    TopologyNode []*PceTopology_TopologyNodes_TopologyNode
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

    topologyNodes.EntityData.Children = types.NewOrderedMap()
    topologyNodes.EntityData.Children.Append("topology-node", types.YChild{"TopologyNode", nil})
    for i := range topologyNodes.TopologyNode {
        topologyNodes.EntityData.Children.Append(types.GetSegmentPath(topologyNodes.TopologyNode[i]), types.YChild{"TopologyNode", topologyNodes.TopologyNode[i]})
    }
    topologyNodes.EntityData.Leafs = types.NewOrderedMap()

    topologyNodes.EntityData.YListKeys = []string {}

    return &(topologyNodes.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode
// Node information
type PceTopology_TopologyNodes_TopologyNode struct {
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
    NodeProtocolIdentifier PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier

    // Prefixes. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Prefixe.
    Prefixe []*PceTopology_TopologyNodes_TopologyNode_Prefixe

    // IPv4 Link information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link.
    Ipv4Link []*PceTopology_TopologyNodes_TopologyNode_Ipv4Link

    // IPv6 Link information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link.
    Ipv6Link []*PceTopology_TopologyNodes_TopologyNode_Ipv6Link
}

func (topologyNode *PceTopology_TopologyNodes_TopologyNode) GetEntityData() *types.CommonEntityData {
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
    topologyNode.EntityData.Children.Append("prefixe", types.YChild{"Prefixe", nil})
    for i := range topologyNode.Prefixe {
        topologyNode.EntityData.Children.Append(types.GetSegmentPath(topologyNode.Prefixe[i]), types.YChild{"Prefixe", topologyNode.Prefixe[i]})
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []*PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
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

    nodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", nodeProtocolIdentifier.IgpInformation[i]})
    }
    nodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", nodeProtocolIdentifier.SrgbInformation[i]})
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Prefixe
// Prefixes
type PceTopology_TopologyNodes_TopologyNode_Prefixe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Prefix SID.
    PfxSid PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId
}

func (prefixe *PceTopology_TopologyNodes_TopologyNode_Prefixe) GetEntityData() *types.CommonEntityData {
    prefixe.EntityData.YFilter = prefixe.YFilter
    prefixe.EntityData.YangName = "prefixe"
    prefixe.EntityData.BundleName = "cisco_ios_xr"
    prefixe.EntityData.ParentYangName = "topology-node"
    prefixe.EntityData.SegmentPath = "prefixe"
    prefixe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixe.EntityData.Children = types.NewOrderedMap()
    prefixe.EntityData.Children.Append("pfx-sid", types.YChild{"PfxSid", &prefixe.PfxSid})
    prefixe.EntityData.Children.Append("node-id", types.YChild{"NodeId", &prefixe.NodeId})
    prefixe.EntityData.Leafs = types.NewOrderedMap()
    prefixe.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", prefixe.DomainIdentifier})

    prefixe.EntityData.YListKeys = []string {}

    return &(prefixe.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid
// Prefix SID
type PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

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
    SidPrefix PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix
}

func (pfxSid *PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid) GetEntityData() *types.CommonEntityData {
    pfxSid.EntityData.YFilter = pfxSid.YFilter
    pfxSid.EntityData.YangName = "pfx-sid"
    pfxSid.EntityData.BundleName = "cisco_ios_xr"
    pfxSid.EntityData.ParentYangName = "prefixe"
    pfxSid.EntityData.SegmentPath = "pfx-sid"
    pfxSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfxSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfxSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfxSid.EntityData.Children = types.NewOrderedMap()
    pfxSid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &pfxSid.SidPrefix})
    pfxSid.EntityData.Leafs = types.NewOrderedMap()
    pfxSid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", pfxSid.SidType})
    pfxSid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", pfxSid.MplsLabel})
    pfxSid.EntityData.Leafs.Append("rflag", types.YLeaf{"Rflag", pfxSid.Rflag})
    pfxSid.EntityData.Leafs.Append("nflag", types.YLeaf{"Nflag", pfxSid.Nflag})
    pfxSid.EntityData.Leafs.Append("pflag", types.YLeaf{"Pflag", pfxSid.Pflag})
    pfxSid.EntityData.Leafs.Append("eflag", types.YLeaf{"Eflag", pfxSid.Eflag})
    pfxSid.EntityData.Leafs.Append("vflag", types.YLeaf{"Vflag", pfxSid.Vflag})
    pfxSid.EntityData.Leafs.Append("lflag", types.YLeaf{"Lflag", pfxSid.Lflag})

    pfxSid.EntityData.YListKeys = []string {}

    return &(pfxSid.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix
// Prefix
type PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix struct {
    EntityData types.CommonEntityData
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

func (sidPrefix *PceTopology_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "pfx-sid"
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

// PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "prefixe"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link struct {
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
    LocalIgpInformation PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier

    // Performance metrics.
    PerformanceMetrics PceTopology_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics

    // Adjacency SIDs. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid.
    AdjacencySid []*PceTopology_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
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

    ipv4Link.EntityData.Children = types.NewOrderedMap()
    ipv4Link.EntityData.Children.Append("local-igp-information", types.YChild{"LocalIgpInformation", &ipv4Link.LocalIgpInformation})
    ipv4Link.EntityData.Children.Append("remote-node-protocol-identifier", types.YChild{"RemoteNodeProtocolIdentifier", &ipv4Link.RemoteNodeProtocolIdentifier})
    ipv4Link.EntityData.Children.Append("performance-metrics", types.YChild{"PerformanceMetrics", &ipv4Link.PerformanceMetrics})
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId
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

    localIgpInformation.EntityData.Children = types.NewOrderedMap()
    localIgpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &localIgpInformation.NodeId})
    localIgpInformation.EntityData.Leafs = types.NewOrderedMap()
    localIgpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier})

    localIgpInformation.EntityData.YListKeys = []string {}

    return &(localIgpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "local-igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []*PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
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

    remoteNodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", remoteNodeProtocolIdentifier.IgpInformation[i]})
    }
    remoteNodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", remoteNodeProtocolIdentifier.SrgbInformation[i]})
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics
// Performance metrics
type PceTopology_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Min delay. The type is interface{} with range: 0..4294967295.
    UnidirectionalMinDelay interface{}
}

func (performanceMetrics *PceTopology_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics) GetEntityData() *types.CommonEntityData {
    performanceMetrics.EntityData.YFilter = performanceMetrics.YFilter
    performanceMetrics.EntityData.YangName = "performance-metrics"
    performanceMetrics.EntityData.BundleName = "cisco_ios_xr"
    performanceMetrics.EntityData.ParentYangName = "ipv4-link"
    performanceMetrics.EntityData.SegmentPath = "performance-metrics"
    performanceMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    performanceMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    performanceMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    performanceMetrics.EntityData.Children = types.NewOrderedMap()
    performanceMetrics.EntityData.Leafs = types.NewOrderedMap()
    performanceMetrics.EntityData.Leafs.Append("unidirectional-min-delay", types.YLeaf{"UnidirectionalMinDelay", performanceMetrics.UnidirectionalMinDelay})

    performanceMetrics.EntityData.YListKeys = []string {}

    return &(performanceMetrics.EntityData)
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

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &adjacencySid.SidPrefix})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()
    adjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", adjacencySid.SidType})
    adjacencySid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", adjacencySid.MplsLabel})
    adjacencySid.EntityData.Leafs.Append("rflag", types.YLeaf{"Rflag", adjacencySid.Rflag})
    adjacencySid.EntityData.Leafs.Append("nflag", types.YLeaf{"Nflag", adjacencySid.Nflag})
    adjacencySid.EntityData.Leafs.Append("pflag", types.YLeaf{"Pflag", adjacencySid.Pflag})
    adjacencySid.EntityData.Leafs.Append("eflag", types.YLeaf{"Eflag", adjacencySid.Eflag})
    adjacencySid.EntityData.Leafs.Append("vflag", types.YLeaf{"Vflag", adjacencySid.Vflag})
    adjacencySid.EntityData.Leafs.Append("lflag", types.YLeaf{"Lflag", adjacencySid.Lflag})

    adjacencySid.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link struct {
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
    LocalIgpInformation PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid.
    AdjacencySid []*PceTopology_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId
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

    localIgpInformation.EntityData.Children = types.NewOrderedMap()
    localIgpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &localIgpInformation.NodeId})
    localIgpInformation.EntityData.Leafs = types.NewOrderedMap()
    localIgpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier})

    localIgpInformation.EntityData.YListKeys = []string {}

    return &(localIgpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "local-igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []*PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
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

    remoteNodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", remoteNodeProtocolIdentifier.IgpInformation[i]})
    }
    remoteNodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", remoteNodeProtocolIdentifier.SrgbInformation[i]})
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &adjacencySid.SidPrefix})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()
    adjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", adjacencySid.SidType})
    adjacencySid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", adjacencySid.MplsLabel})
    adjacencySid.EntityData.Leafs.Append("rflag", types.YLeaf{"Rflag", adjacencySid.Rflag})
    adjacencySid.EntityData.Leafs.Append("nflag", types.YLeaf{"Nflag", adjacencySid.Nflag})
    adjacencySid.EntityData.Leafs.Append("pflag", types.YLeaf{"Pflag", adjacencySid.Pflag})
    adjacencySid.EntityData.Leafs.Append("eflag", types.YLeaf{"Eflag", adjacencySid.Eflag})
    adjacencySid.EntityData.Leafs.Append("vflag", types.YLeaf{"Vflag", adjacencySid.Vflag})
    adjacencySid.EntityData.Leafs.Append("lflag", types.YLeaf{"Lflag", adjacencySid.Lflag})

    adjacencySid.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// PceTopology_PrefixInfos
// Prefixes database in XTC
type PceTopology_PrefixInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE prefix information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo.
    PrefixInfo []*PceTopology_PrefixInfos_PrefixInfo
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

    prefixInfos.EntityData.Children = types.NewOrderedMap()
    prefixInfos.EntityData.Children.Append("prefix-info", types.YChild{"PrefixInfo", nil})
    for i := range prefixInfos.PrefixInfo {
        prefixInfos.EntityData.Children.Append(types.GetSegmentPath(prefixInfos.PrefixInfo[i]), types.YChild{"PrefixInfo", prefixInfos.PrefixInfo[i]})
    }
    prefixInfos.EntityData.Leafs = types.NewOrderedMap()

    prefixInfos.EntityData.YListKeys = []string {}

    return &(prefixInfos.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo
// PCE prefix information
type PceTopology_PrefixInfos_PrefixInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is interface{} with range:
    // 0..4294967295.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier

    // Prefix address. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_Address.
    Address []*PceTopology_PrefixInfos_PrefixInfo_Address
}

func (prefixInfo *PceTopology_PrefixInfos_PrefixInfo) GetEntityData() *types.CommonEntityData {
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []*PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
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

    nodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", nodeProtocolIdentifier.IgpInformation[i]})
    }
    nodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", nodeProtocolIdentifier.SrgbInformation[i]})
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *PceTopology_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("ip", types.YChild{"Ip", &address.Ip})
    address.EntityData.Leafs = types.NewOrderedMap()

    address.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", ip.AfName})
    ip.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ip.Ipv4})
    ip.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ip.Ipv6})

    ip.EntityData.YListKeys = []string {}

    return &(ip.EntityData)
}

// Pce
// pce
type Pce struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE Verification events in XTC.
    VerificationEvents Pce_VerificationEvents

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

    pce.EntityData.Children = types.NewOrderedMap()
    pce.EntityData.Children.Append("verification-events", types.YChild{"VerificationEvents", &pce.VerificationEvents})
    pce.EntityData.Children.Append("association-infos", types.YChild{"AssociationInfos", &pce.AssociationInfos})
    pce.EntityData.Children.Append("cspf", types.YChild{"Cspf", &pce.Cspf})
    pce.EntityData.Children.Append("topology-summary", types.YChild{"TopologySummary", &pce.TopologySummary})
    pce.EntityData.Children.Append("tunnel-infos", types.YChild{"TunnelInfos", &pce.TunnelInfos})
    pce.EntityData.Children.Append("peer-detail-infos", types.YChild{"PeerDetailInfos", &pce.PeerDetailInfos})
    pce.EntityData.Children.Append("topology-nodes", types.YChild{"TopologyNodes", &pce.TopologyNodes})
    pce.EntityData.Children.Append("prefix-infos", types.YChild{"PrefixInfos", &pce.PrefixInfos})
    pce.EntityData.Children.Append("lsp-summary", types.YChild{"LspSummary", &pce.LspSummary})
    pce.EntityData.Children.Append("peer-infos", types.YChild{"PeerInfos", &pce.PeerInfos})
    pce.EntityData.Children.Append("tunnel-detail-infos", types.YChild{"TunnelDetailInfos", &pce.TunnelDetailInfos})
    pce.EntityData.Leafs = types.NewOrderedMap()

    pce.EntityData.YListKeys = []string {}

    return &(pce.EntityData)
}

// Pce_VerificationEvents
// PCE Verification events in XTC
type Pce_VerificationEvents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE single verification event. The type is slice of
    // Pce_VerificationEvents_VerificationEvent.
    VerificationEvent []*Pce_VerificationEvents_VerificationEvent
}

func (verificationEvents *Pce_VerificationEvents) GetEntityData() *types.CommonEntityData {
    verificationEvents.EntityData.YFilter = verificationEvents.YFilter
    verificationEvents.EntityData.YangName = "verification-events"
    verificationEvents.EntityData.BundleName = "cisco_ios_xr"
    verificationEvents.EntityData.ParentYangName = "pce"
    verificationEvents.EntityData.SegmentPath = "verification-events"
    verificationEvents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verificationEvents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verificationEvents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verificationEvents.EntityData.Children = types.NewOrderedMap()
    verificationEvents.EntityData.Children.Append("verification-event", types.YChild{"VerificationEvent", nil})
    for i := range verificationEvents.VerificationEvent {
        verificationEvents.EntityData.Children.Append(types.GetSegmentPath(verificationEvents.VerificationEvent[i]), types.YChild{"VerificationEvent", verificationEvents.VerificationEvent[i]})
    }
    verificationEvents.EntityData.Leafs = types.NewOrderedMap()

    verificationEvents.EntityData.YListKeys = []string {}

    return &(verificationEvents.EntityData)
}

// Pce_VerificationEvents_VerificationEvent
// PCE single verification event
type Pce_VerificationEvents_VerificationEvent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Index of an event. The type is interface{} with
    // range: 0..4294967295.
    EventIdx interface{}

    // Event ID in range 1 - 0xFFFFFFFF. 0 is invalid. The type is interface{}
    // with range: 0..4294967295.
    EventId interface{}

    // Event message. The type is string.
    EventMessage interface{}

    // Event time, relative to Jan 1, 1970. The type is interface{} with range:
    // 0..4294967295.
    TimeStamp interface{}
}

func (verificationEvent *Pce_VerificationEvents_VerificationEvent) GetEntityData() *types.CommonEntityData {
    verificationEvent.EntityData.YFilter = verificationEvent.YFilter
    verificationEvent.EntityData.YangName = "verification-event"
    verificationEvent.EntityData.BundleName = "cisco_ios_xr"
    verificationEvent.EntityData.ParentYangName = "verification-events"
    verificationEvent.EntityData.SegmentPath = "verification-event" + types.AddKeyToken(verificationEvent.EventIdx, "event-idx")
    verificationEvent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    verificationEvent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    verificationEvent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    verificationEvent.EntityData.Children = types.NewOrderedMap()
    verificationEvent.EntityData.Leafs = types.NewOrderedMap()
    verificationEvent.EntityData.Leafs.Append("event-idx", types.YLeaf{"EventIdx", verificationEvent.EventIdx})
    verificationEvent.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", verificationEvent.EventId})
    verificationEvent.EntityData.Leafs.Append("event-message", types.YLeaf{"EventMessage", verificationEvent.EventMessage})
    verificationEvent.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", verificationEvent.TimeStamp})

    verificationEvent.EntityData.YListKeys = []string {"EventIdx"}

    return &(verificationEvent.EntityData)
}

// Pce_AssociationInfos
// Associaition database in XTC
type Pce_AssociationInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE Association information. The type is slice of
    // Pce_AssociationInfos_AssociationInfo.
    AssociationInfo []*Pce_AssociationInfos_AssociationInfo
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

    associationInfos.EntityData.Children = types.NewOrderedMap()
    associationInfos.EntityData.Children.Append("association-info", types.YChild{"AssociationInfo", nil})
    for i := range associationInfos.AssociationInfo {
        associationInfos.EntityData.Children.Append(types.GetSegmentPath(associationInfos.AssociationInfo[i]), types.YChild{"AssociationInfo", associationInfos.AssociationInfo[i]})
    }
    associationInfos.EntityData.Leafs = types.NewOrderedMap()

    associationInfos.EntityData.YListKeys = []string {}

    return &(associationInfos.EntityData)
}

// Pce_AssociationInfos_AssociationInfo
// PCE Association information
type Pce_AssociationInfos_AssociationInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Group ID. The type is interface{} with range:
    // 0..4294967295.
    GroupId interface{}

    // Type. The type is PceAsso.
    Type interface{}

    // Sub ID. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    SubId interface{}

    // Association Type. The type is interface{} with range: 0..4294967295.
    AssociationType interface{}

    // Association ID. The type is interface{} with range: 0..4294967295.
    AssociationId interface{}

    // Association Strict Mode. The type is bool.
    Strict interface{}

    // Association Status. The type is interface{} with range: 0..4294967295.
    Status interface{}

    // Headends Swapped. The type is interface{} with range: 0..4294967295.
    HeadendsSwapped interface{}

    // Association Source.
    AssociationSource Pce_AssociationInfos_AssociationInfo_AssociationSource

    // Association LSP Info. The type is slice of
    // Pce_AssociationInfos_AssociationInfo_AssociationLsp.
    AssociationLsp []*Pce_AssociationInfos_AssociationInfo_AssociationLsp
}

func (associationInfo *Pce_AssociationInfos_AssociationInfo) GetEntityData() *types.CommonEntityData {
    associationInfo.EntityData.YFilter = associationInfo.YFilter
    associationInfo.EntityData.YangName = "association-info"
    associationInfo.EntityData.BundleName = "cisco_ios_xr"
    associationInfo.EntityData.ParentYangName = "association-infos"
    associationInfo.EntityData.SegmentPath = "association-info" + types.AddKeyToken(associationInfo.GroupId, "group-id")
    associationInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationInfo.EntityData.Children = types.NewOrderedMap()
    associationInfo.EntityData.Children.Append("association-source", types.YChild{"AssociationSource", &associationInfo.AssociationSource})
    associationInfo.EntityData.Children.Append("association-lsp", types.YChild{"AssociationLsp", nil})
    for i := range associationInfo.AssociationLsp {
        associationInfo.EntityData.Children.Append(types.GetSegmentPath(associationInfo.AssociationLsp[i]), types.YChild{"AssociationLsp", associationInfo.AssociationLsp[i]})
    }
    associationInfo.EntityData.Leafs = types.NewOrderedMap()
    associationInfo.EntityData.Leafs.Append("group-id", types.YLeaf{"GroupId", associationInfo.GroupId})
    associationInfo.EntityData.Leafs.Append("type", types.YLeaf{"Type", associationInfo.Type})
    associationInfo.EntityData.Leafs.Append("sub-id", types.YLeaf{"SubId", associationInfo.SubId})
    associationInfo.EntityData.Leafs.Append("association-type", types.YLeaf{"AssociationType", associationInfo.AssociationType})
    associationInfo.EntityData.Leafs.Append("association-id", types.YLeaf{"AssociationId", associationInfo.AssociationId})
    associationInfo.EntityData.Leafs.Append("strict", types.YLeaf{"Strict", associationInfo.Strict})
    associationInfo.EntityData.Leafs.Append("status", types.YLeaf{"Status", associationInfo.Status})
    associationInfo.EntityData.Leafs.Append("headends-swapped", types.YLeaf{"HeadendsSwapped", associationInfo.HeadendsSwapped})

    associationInfo.EntityData.YListKeys = []string {"GroupId"}

    return &(associationInfo.EntityData)
}

// Pce_AssociationInfos_AssociationInfo_AssociationSource
// Association Source
type Pce_AssociationInfos_AssociationInfo_AssociationSource struct {
    EntityData types.CommonEntityData
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

func (associationSource *Pce_AssociationInfos_AssociationInfo_AssociationSource) GetEntityData() *types.CommonEntityData {
    associationSource.EntityData.YFilter = associationSource.YFilter
    associationSource.EntityData.YangName = "association-source"
    associationSource.EntityData.BundleName = "cisco_ios_xr"
    associationSource.EntityData.ParentYangName = "association-info"
    associationSource.EntityData.SegmentPath = "association-source"
    associationSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationSource.EntityData.Children = types.NewOrderedMap()
    associationSource.EntityData.Leafs = types.NewOrderedMap()
    associationSource.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", associationSource.AfName})
    associationSource.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", associationSource.Ipv4})
    associationSource.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", associationSource.Ipv6})

    associationSource.EntityData.YListKeys = []string {}

    return &(associationSource.EntityData)
}

// Pce_AssociationInfos_AssociationInfo_AssociationLsp
// Association LSP Info
type Pce_AssociationInfos_AssociationInfo_AssociationLsp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // PCC address.
    PccAddress Pce_AssociationInfos_AssociationInfo_AssociationLsp_PccAddress
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

    associationLsp.EntityData.Children = types.NewOrderedMap()
    associationLsp.EntityData.Children.Append("pcc-address", types.YChild{"PccAddress", &associationLsp.PccAddress})
    associationLsp.EntityData.Leafs = types.NewOrderedMap()
    associationLsp.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", associationLsp.TunnelId})
    associationLsp.EntityData.Leafs.Append("lspid", types.YLeaf{"Lspid", associationLsp.Lspid})
    associationLsp.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", associationLsp.TunnelName})
    associationLsp.EntityData.Leafs.Append("pce-based", types.YLeaf{"PceBased", associationLsp.PceBased})
    associationLsp.EntityData.Leafs.Append("plsp-id", types.YLeaf{"PlspId", associationLsp.PlspId})

    associationLsp.EntityData.YListKeys = []string {}

    return &(associationLsp.EntityData)
}

// Pce_AssociationInfos_AssociationInfo_AssociationLsp_PccAddress
// PCC address
type Pce_AssociationInfos_AssociationInfo_AssociationLsp_PccAddress struct {
    EntityData types.CommonEntityData
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

func (pccAddress *Pce_AssociationInfos_AssociationInfo_AssociationLsp_PccAddress) GetEntityData() *types.CommonEntityData {
    pccAddress.EntityData.YFilter = pccAddress.YFilter
    pccAddress.EntityData.YangName = "pcc-address"
    pccAddress.EntityData.BundleName = "cisco_ios_xr"
    pccAddress.EntityData.ParentYangName = "association-lsp"
    pccAddress.EntityData.SegmentPath = "pcc-address"
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = types.NewOrderedMap()
    pccAddress.EntityData.Leafs = types.NewOrderedMap()
    pccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pccAddress.AfName})
    pccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", pccAddress.Ipv4})
    pccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", pccAddress.Ipv6})

    pccAddress.EntityData.YListKeys = []string {}

    return &(pccAddress.EntityData)
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

    cspf.EntityData.Children = types.NewOrderedMap()
    cspf.EntityData.Children.Append("cspf-paths", types.YChild{"CspfPaths", &cspf.CspfPaths})
    cspf.EntityData.Leafs = types.NewOrderedMap()

    cspf.EntityData.YListKeys = []string {}

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
    CspfPath []*Pce_Cspf_CspfPaths_CspfPath
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

    cspfPaths.EntityData.Children = types.NewOrderedMap()
    cspfPaths.EntityData.Children.Append("cspf-path", types.YChild{"CspfPath", nil})
    for i := range cspfPaths.CspfPath {
        cspfPaths.EntityData.Children.Append(types.GetSegmentPath(cspfPaths.CspfPath[i]), types.YChild{"CspfPath", cspfPaths.CspfPath[i]})
    }
    cspfPaths.EntityData.Leafs = types.NewOrderedMap()

    cspfPaths.EntityData.YListKeys = []string {}

    return &(cspfPaths.EntityData)
}

// Pce_Cspf_CspfPaths_CspfPath
// A GET operation on this class returns the path
// .
type Pce_Cspf_CspfPaths_CspfPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address Family. The type is interface{} with
    // range: 0..4294967295.
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
    // 0..4294967295.
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
    // range: 0..4294967295.
    DisjointLevel interface{}

    // This attribute is a key. Strict disjointness required. The type is
    // interface{} with range: 0..4294967295.
    DisjointStrict interface{}

    // This attribute is a key. Whether path 1 or 2 should be shortest. The type
    // is interface{} with range: 0..4294967295.
    ShortestPath interface{}

    // Headends swapped. The type is PceHeadendSwap.
    HeadendsSwapped interface{}

    // CSPF Result. The type is PceCspfRc.
    CspfResult interface{}

    // Iterations of the Suurballe-Tarjan algorithm. The type is interface{} with
    // range: 0..4294967295.
    IterationsDone interface{}

    // Output PCE paths. The type is slice of
    // Pce_Cspf_CspfPaths_CspfPath_OutputPath.
    OutputPath []*Pce_Cspf_CspfPaths_CspfPath_OutputPath
}

func (cspfPath *Pce_Cspf_CspfPaths_CspfPath) GetEntityData() *types.CommonEntityData {
    cspfPath.EntityData.YFilter = cspfPath.YFilter
    cspfPath.EntityData.YangName = "cspf-path"
    cspfPath.EntityData.BundleName = "cisco_ios_xr"
    cspfPath.EntityData.ParentYangName = "cspf-paths"
    cspfPath.EntityData.SegmentPath = "cspf-path" + types.AddKeyToken(cspfPath.Af, "af") + types.AddKeyToken(cspfPath.Source1, "source1") + types.AddKeyToken(cspfPath.Destination1, "destination1") + types.AddKeyToken(cspfPath.MetricType, "metric-type") + types.AddKeyToken(cspfPath.Source2, "source2") + types.AddKeyToken(cspfPath.Destination2, "destination2") + types.AddKeyToken(cspfPath.DisjointLevel, "disjoint-level") + types.AddKeyToken(cspfPath.DisjointStrict, "disjoint-strict") + types.AddKeyToken(cspfPath.ShortestPath, "shortest-path")
    cspfPath.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cspfPath.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cspfPath.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cspfPath.EntityData.Children = types.NewOrderedMap()
    cspfPath.EntityData.Children.Append("output-path", types.YChild{"OutputPath", nil})
    for i := range cspfPath.OutputPath {
        cspfPath.EntityData.Children.Append(types.GetSegmentPath(cspfPath.OutputPath[i]), types.YChild{"OutputPath", cspfPath.OutputPath[i]})
    }
    cspfPath.EntityData.Leafs = types.NewOrderedMap()
    cspfPath.EntityData.Leafs.Append("af", types.YLeaf{"Af", cspfPath.Af})
    cspfPath.EntityData.Leafs.Append("source1", types.YLeaf{"Source1", cspfPath.Source1})
    cspfPath.EntityData.Leafs.Append("destination1", types.YLeaf{"Destination1", cspfPath.Destination1})
    cspfPath.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", cspfPath.MetricType})
    cspfPath.EntityData.Leafs.Append("source2", types.YLeaf{"Source2", cspfPath.Source2})
    cspfPath.EntityData.Leafs.Append("destination2", types.YLeaf{"Destination2", cspfPath.Destination2})
    cspfPath.EntityData.Leafs.Append("disjoint-level", types.YLeaf{"DisjointLevel", cspfPath.DisjointLevel})
    cspfPath.EntityData.Leafs.Append("disjoint-strict", types.YLeaf{"DisjointStrict", cspfPath.DisjointStrict})
    cspfPath.EntityData.Leafs.Append("shortest-path", types.YLeaf{"ShortestPath", cspfPath.ShortestPath})
    cspfPath.EntityData.Leafs.Append("headends-swapped", types.YLeaf{"HeadendsSwapped", cspfPath.HeadendsSwapped})
    cspfPath.EntityData.Leafs.Append("cspf-result", types.YLeaf{"CspfResult", cspfPath.CspfResult})
    cspfPath.EntityData.Leafs.Append("iterations-done", types.YLeaf{"IterationsDone", cspfPath.IterationsDone})

    cspfPath.EntityData.YListKeys = []string {"Af", "Source1", "Destination1", "MetricType", "Source2", "Destination2", "DisjointLevel", "DisjointStrict", "ShortestPath"}

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
    Hops []*Pce_Cspf_CspfPaths_CspfPath_OutputPath_Hops
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

    outputPath.EntityData.Children = types.NewOrderedMap()
    outputPath.EntityData.Children.Append("source", types.YChild{"Source", &outputPath.Source})
    outputPath.EntityData.Children.Append("destination", types.YChild{"Destination", &outputPath.Destination})
    outputPath.EntityData.Children.Append("hops", types.YChild{"Hops", nil})
    for i := range outputPath.Hops {
        outputPath.EntityData.Children.Append(types.GetSegmentPath(outputPath.Hops[i]), types.YChild{"Hops", outputPath.Hops[i]})
    }
    outputPath.EntityData.Leafs = types.NewOrderedMap()
    outputPath.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", outputPath.Cost})

    outputPath.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    source.EntityData.Children = types.NewOrderedMap()
    source.EntityData.Leafs = types.NewOrderedMap()
    source.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", source.AfName})
    source.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", source.Ipv4})
    source.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", source.Ipv6})

    source.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    destination.EntityData.Children = types.NewOrderedMap()
    destination.EntityData.Leafs = types.NewOrderedMap()
    destination.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", destination.AfName})
    destination.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", destination.Ipv4})
    destination.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", destination.Ipv6})

    destination.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Prefix interface{}

    // IPv6 prefix. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    hops.EntityData.Children = types.NewOrderedMap()
    hops.EntityData.Leafs = types.NewOrderedMap()
    hops.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", hops.AddressFamily})
    hops.EntityData.Leafs.Append("ipv4-prefix", types.YLeaf{"Ipv4Prefix", hops.Ipv4Prefix})
    hops.EntityData.Leafs.Append("ipv6-prefix", types.YLeaf{"Ipv6Prefix", hops.Ipv6Prefix})

    hops.EntityData.YListKeys = []string {}

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

    topologySummary.EntityData.Children = types.NewOrderedMap()
    topologySummary.EntityData.Children.Append("stats-topology-update", types.YChild{"StatsTopologyUpdate", &topologySummary.StatsTopologyUpdate})
    topologySummary.EntityData.Leafs = types.NewOrderedMap()
    topologySummary.EntityData.Leafs.Append("nodes", types.YLeaf{"Nodes", topologySummary.Nodes})
    topologySummary.EntityData.Leafs.Append("lookup-nodes", types.YLeaf{"LookupNodes", topologySummary.LookupNodes})
    topologySummary.EntityData.Leafs.Append("prefixes", types.YLeaf{"Prefixes", topologySummary.Prefixes})
    topologySummary.EntityData.Leafs.Append("prefix-sids", types.YLeaf{"PrefixSids", topologySummary.PrefixSids})
    topologySummary.EntityData.Leafs.Append("regular-prefix-sids", types.YLeaf{"RegularPrefixSids", topologySummary.RegularPrefixSids})
    topologySummary.EntityData.Leafs.Append("strict-prefix-sids", types.YLeaf{"StrictPrefixSids", topologySummary.StrictPrefixSids})
    topologySummary.EntityData.Leafs.Append("links", types.YLeaf{"Links", topologySummary.Links})
    topologySummary.EntityData.Leafs.Append("epe-links", types.YLeaf{"EpeLinks", topologySummary.EpeLinks})
    topologySummary.EntityData.Leafs.Append("adjacency-sids", types.YLeaf{"AdjacencySids", topologySummary.AdjacencySids})
    topologySummary.EntityData.Leafs.Append("epesids", types.YLeaf{"Epesids", topologySummary.Epesids})
    topologySummary.EntityData.Leafs.Append("protected-adjacency-sids", types.YLeaf{"ProtectedAdjacencySids", topologySummary.ProtectedAdjacencySids})
    topologySummary.EntityData.Leafs.Append("un-protected-adjacency-sids", types.YLeaf{"UnProtectedAdjacencySids", topologySummary.UnProtectedAdjacencySids})
    topologySummary.EntityData.Leafs.Append("topology-consistent", types.YLeaf{"TopologyConsistent", topologySummary.TopologyConsistent})

    topologySummary.EntityData.YListKeys = []string {}

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

    statsTopologyUpdate.EntityData.Children = types.NewOrderedMap()
    statsTopologyUpdate.EntityData.Leafs = types.NewOrderedMap()
    statsTopologyUpdate.EntityData.Leafs.Append("num-nodes-added", types.YLeaf{"NumNodesAdded", statsTopologyUpdate.NumNodesAdded})
    statsTopologyUpdate.EntityData.Leafs.Append("num-nodes-deleted", types.YLeaf{"NumNodesDeleted", statsTopologyUpdate.NumNodesDeleted})
    statsTopologyUpdate.EntityData.Leafs.Append("num-links-added", types.YLeaf{"NumLinksAdded", statsTopologyUpdate.NumLinksAdded})
    statsTopologyUpdate.EntityData.Leafs.Append("num-links-deleted", types.YLeaf{"NumLinksDeleted", statsTopologyUpdate.NumLinksDeleted})
    statsTopologyUpdate.EntityData.Leafs.Append("num-prefixes-added", types.YLeaf{"NumPrefixesAdded", statsTopologyUpdate.NumPrefixesAdded})
    statsTopologyUpdate.EntityData.Leafs.Append("num-prefixes-deleted", types.YLeaf{"NumPrefixesDeleted", statsTopologyUpdate.NumPrefixesDeleted})

    statsTopologyUpdate.EntityData.YListKeys = []string {}

    return &(statsTopologyUpdate.EntityData)
}

// Pce_TunnelInfos
// Tunnel database in XTC
type Pce_TunnelInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel information. The type is slice of Pce_TunnelInfos_TunnelInfo.
    TunnelInfo []*Pce_TunnelInfos_TunnelInfo
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

    tunnelInfos.EntityData.Children = types.NewOrderedMap()
    tunnelInfos.EntityData.Children.Append("tunnel-info", types.YChild{"TunnelInfo", nil})
    for i := range tunnelInfos.TunnelInfo {
        tunnelInfos.EntityData.Children.Append(types.GetSegmentPath(tunnelInfos.TunnelInfo[i]), types.YChild{"TunnelInfo", tunnelInfos.TunnelInfo[i]})
    }
    tunnelInfos.EntityData.Leafs = types.NewOrderedMap()

    tunnelInfos.EntityData.YListKeys = []string {}

    return &(tunnelInfos.EntityData)
}

// Pce_TunnelInfos_TunnelInfo
// Tunnel information
type Pce_TunnelInfos_TunnelInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // 0..4294967295.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string.
    TunnelName interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // PCC address.
    PccAddress Pce_TunnelInfos_TunnelInfo_PccAddress

    // Brief LSP information. The type is slice of
    // Pce_TunnelInfos_TunnelInfo_BriefLspInformation.
    BriefLspInformation []*Pce_TunnelInfos_TunnelInfo_BriefLspInformation
}

func (tunnelInfo *Pce_TunnelInfos_TunnelInfo) GetEntityData() *types.CommonEntityData {
    tunnelInfo.EntityData.YFilter = tunnelInfo.YFilter
    tunnelInfo.EntityData.YangName = "tunnel-info"
    tunnelInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelInfo.EntityData.ParentYangName = "tunnel-infos"
    tunnelInfo.EntityData.SegmentPath = "tunnel-info" + types.AddKeyToken(tunnelInfo.PeerAddress, "peer-address") + types.AddKeyToken(tunnelInfo.PlspId, "plsp-id") + types.AddKeyToken(tunnelInfo.TunnelName, "tunnel-name")
    tunnelInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelInfo.EntityData.Children = types.NewOrderedMap()
    tunnelInfo.EntityData.Children.Append("pcc-address", types.YChild{"PccAddress", &tunnelInfo.PccAddress})
    tunnelInfo.EntityData.Children.Append("brief-lsp-information", types.YChild{"BriefLspInformation", nil})
    for i := range tunnelInfo.BriefLspInformation {
        tunnelInfo.EntityData.Children.Append(types.GetSegmentPath(tunnelInfo.BriefLspInformation[i]), types.YChild{"BriefLspInformation", tunnelInfo.BriefLspInformation[i]})
    }
    tunnelInfo.EntityData.Leafs = types.NewOrderedMap()
    tunnelInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", tunnelInfo.PeerAddress})
    tunnelInfo.EntityData.Leafs.Append("plsp-id", types.YLeaf{"PlspId", tunnelInfo.PlspId})
    tunnelInfo.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", tunnelInfo.TunnelName})
    tunnelInfo.EntityData.Leafs.Append("tunnel-name-xr", types.YLeaf{"TunnelNameXr", tunnelInfo.TunnelNameXr})

    tunnelInfo.EntityData.YListKeys = []string {"PeerAddress", "PlspId", "TunnelName"}

    return &(tunnelInfo.EntityData)
}

// Pce_TunnelInfos_TunnelInfo_PccAddress
// PCC address
type Pce_TunnelInfos_TunnelInfo_PccAddress struct {
    EntityData types.CommonEntityData
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

func (pccAddress *Pce_TunnelInfos_TunnelInfo_PccAddress) GetEntityData() *types.CommonEntityData {
    pccAddress.EntityData.YFilter = pccAddress.YFilter
    pccAddress.EntityData.YangName = "pcc-address"
    pccAddress.EntityData.BundleName = "cisco_ios_xr"
    pccAddress.EntityData.ParentYangName = "tunnel-info"
    pccAddress.EntityData.SegmentPath = "pcc-address"
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = types.NewOrderedMap()
    pccAddress.EntityData.Leafs = types.NewOrderedMap()
    pccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pccAddress.AfName})
    pccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", pccAddress.Ipv4})
    pccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", pccAddress.Ipv6})

    pccAddress.EntityData.YListKeys = []string {}

    return &(pccAddress.EntityData)
}

// Pce_TunnelInfos_TunnelInfo_BriefLspInformation
// Brief LSP information
type Pce_TunnelInfos_TunnelInfo_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    Msd interface{}

    // Source address.
    SourceAddress Pce_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress

    // Destination address.
    DestinationAddress Pce_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress
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

    briefLspInformation.EntityData.Children = types.NewOrderedMap()
    briefLspInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &briefLspInformation.SourceAddress})
    briefLspInformation.EntityData.Children.Append("destination-address", types.YChild{"DestinationAddress", &briefLspInformation.DestinationAddress})
    briefLspInformation.EntityData.Leafs = types.NewOrderedMap()
    briefLspInformation.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", briefLspInformation.TunnelId})
    briefLspInformation.EntityData.Leafs.Append("lspid", types.YLeaf{"Lspid", briefLspInformation.Lspid})
    briefLspInformation.EntityData.Leafs.Append("binding-sid", types.YLeaf{"BindingSid", briefLspInformation.BindingSid})
    briefLspInformation.EntityData.Leafs.Append("lsp-setup-type", types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType})
    briefLspInformation.EntityData.Leafs.Append("operational-state", types.YLeaf{"OperationalState", briefLspInformation.OperationalState})
    briefLspInformation.EntityData.Leafs.Append("administrative-state", types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState})
    briefLspInformation.EntityData.Leafs.Append("msd", types.YLeaf{"Msd", briefLspInformation.Msd})

    briefLspInformation.EntityData.YListKeys = []string {}

    return &(briefLspInformation.EntityData)
}

// Pce_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress
// Source address
type Pce_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress struct {
    EntityData types.CommonEntityData
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

func (sourceAddress *Pce_TunnelInfos_TunnelInfo_BriefLspInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "brief-lsp-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sourceAddress.AfName})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Pce_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress
// Destination address
type Pce_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress struct {
    EntityData types.CommonEntityData
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

func (destinationAddress *Pce_TunnelInfos_TunnelInfo_BriefLspInformation_DestinationAddress) GetEntityData() *types.CommonEntityData {
    destinationAddress.EntityData.YFilter = destinationAddress.YFilter
    destinationAddress.EntityData.YangName = "destination-address"
    destinationAddress.EntityData.BundleName = "cisco_ios_xr"
    destinationAddress.EntityData.ParentYangName = "brief-lsp-information"
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

// Pce_PeerDetailInfos
// Detailed peers database in XTC
type Pce_PeerDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed PCE peer information. The type is slice of
    // Pce_PeerDetailInfos_PeerDetailInfo.
    PeerDetailInfo []*Pce_PeerDetailInfos_PeerDetailInfo
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

    peerDetailInfos.EntityData.Children = types.NewOrderedMap()
    peerDetailInfos.EntityData.Children.Append("peer-detail-info", types.YChild{"PeerDetailInfo", nil})
    for i := range peerDetailInfos.PeerDetailInfo {
        peerDetailInfos.EntityData.Children.Append(types.GetSegmentPath(peerDetailInfos.PeerDetailInfo[i]), types.YChild{"PeerDetailInfo", peerDetailInfos.PeerDetailInfo[i]})
    }
    peerDetailInfos.EntityData.Leafs = types.NewOrderedMap()

    peerDetailInfos.EntityData.YListKeys = []string {}

    return &(peerDetailInfos.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo
// Detailed PCE peer information
type Pce_PeerDetailInfos_PeerDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    MaxSidDepth interface{}

    // Peer address.
    PeerAddressXr Pce_PeerDetailInfos_PeerDetailInfo_PeerAddressXr

    // Detailed PCE protocol information.
    DetailPcepInformation Pce_PeerDetailInfos_PeerDetailInfo_DetailPcepInformation
}

func (peerDetailInfo *Pce_PeerDetailInfos_PeerDetailInfo) GetEntityData() *types.CommonEntityData {
    peerDetailInfo.EntityData.YFilter = peerDetailInfo.YFilter
    peerDetailInfo.EntityData.YangName = "peer-detail-info"
    peerDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    peerDetailInfo.EntityData.ParentYangName = "peer-detail-infos"
    peerDetailInfo.EntityData.SegmentPath = "peer-detail-info" + types.AddKeyToken(peerDetailInfo.PeerAddress, "peer-address")
    peerDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerDetailInfo.EntityData.Children = types.NewOrderedMap()
    peerDetailInfo.EntityData.Children.Append("peer-address-xr", types.YChild{"PeerAddressXr", &peerDetailInfo.PeerAddressXr})
    peerDetailInfo.EntityData.Children.Append("detail-pcep-information", types.YChild{"DetailPcepInformation", &peerDetailInfo.DetailPcepInformation})
    peerDetailInfo.EntityData.Leafs = types.NewOrderedMap()
    peerDetailInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", peerDetailInfo.PeerAddress})
    peerDetailInfo.EntityData.Leafs.Append("peer-protocol", types.YLeaf{"PeerProtocol", peerDetailInfo.PeerProtocol})
    peerDetailInfo.EntityData.Leafs.Append("max-sid-depth", types.YLeaf{"MaxSidDepth", peerDetailInfo.MaxSidDepth})

    peerDetailInfo.EntityData.YListKeys = []string {"PeerAddress"}

    return &(peerDetailInfo.EntityData)
}

// Pce_PeerDetailInfos_PeerDetailInfo_PeerAddressXr
// Peer address
type Pce_PeerDetailInfos_PeerDetailInfo_PeerAddressXr struct {
    EntityData types.CommonEntityData
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

func (peerAddressXr *Pce_PeerDetailInfos_PeerDetailInfo_PeerAddressXr) GetEntityData() *types.CommonEntityData {
    peerAddressXr.EntityData.YFilter = peerAddressXr.YFilter
    peerAddressXr.EntityData.YangName = "peer-address-xr"
    peerAddressXr.EntityData.BundleName = "cisco_ios_xr"
    peerAddressXr.EntityData.ParentYangName = "peer-detail-info"
    peerAddressXr.EntityData.SegmentPath = "peer-address-xr"
    peerAddressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddressXr.EntityData.Children = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddressXr.AfName})
    peerAddressXr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerAddressXr.Ipv4})
    peerAddressXr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerAddressXr.Ipv6})

    peerAddressXr.EntityData.YListKeys = []string {}

    return &(peerAddressXr.EntityData)
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

    // Maximum number of labels the peer can impose. The type is interface{} with
    // range: 0..255.
    MaxSidDepth interface{}

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

    detailPcepInformation.EntityData.Children = types.NewOrderedMap()
    detailPcepInformation.EntityData.Children.Append("brief-pcep-information", types.YChild{"BriefPcepInformation", &detailPcepInformation.BriefPcepInformation})
    detailPcepInformation.EntityData.Children.Append("last-error-rx", types.YChild{"LastErrorRx", &detailPcepInformation.LastErrorRx})
    detailPcepInformation.EntityData.Children.Append("last-error-tx", types.YChild{"LastErrorTx", &detailPcepInformation.LastErrorTx})
    detailPcepInformation.EntityData.Leafs = types.NewOrderedMap()
    detailPcepInformation.EntityData.Leafs.Append("error", types.YLeaf{"Error", detailPcepInformation.Error})
    detailPcepInformation.EntityData.Leafs.Append("speaker-id", types.YLeaf{"SpeakerId", detailPcepInformation.SpeakerId})
    detailPcepInformation.EntityData.Leafs.Append("pcep-up-time", types.YLeaf{"PcepUpTime", detailPcepInformation.PcepUpTime})
    detailPcepInformation.EntityData.Leafs.Append("keepalives", types.YLeaf{"Keepalives", detailPcepInformation.Keepalives})
    detailPcepInformation.EntityData.Leafs.Append("md5-enabled", types.YLeaf{"Md5Enabled", detailPcepInformation.Md5Enabled})
    detailPcepInformation.EntityData.Leafs.Append("keychain-enabled", types.YLeaf{"KeychainEnabled", detailPcepInformation.KeychainEnabled})
    detailPcepInformation.EntityData.Leafs.Append("negotiated-local-keepalive", types.YLeaf{"NegotiatedLocalKeepalive", detailPcepInformation.NegotiatedLocalKeepalive})
    detailPcepInformation.EntityData.Leafs.Append("negotiated-remote-keepalive", types.YLeaf{"NegotiatedRemoteKeepalive", detailPcepInformation.NegotiatedRemoteKeepalive})
    detailPcepInformation.EntityData.Leafs.Append("negotiated-dead-time", types.YLeaf{"NegotiatedDeadTime", detailPcepInformation.NegotiatedDeadTime})
    detailPcepInformation.EntityData.Leafs.Append("pce-request-rx", types.YLeaf{"PceRequestRx", detailPcepInformation.PceRequestRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-request-tx", types.YLeaf{"PceRequestTx", detailPcepInformation.PceRequestTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-reply-rx", types.YLeaf{"PceReplyRx", detailPcepInformation.PceReplyRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-reply-tx", types.YLeaf{"PceReplyTx", detailPcepInformation.PceReplyTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-error-rx", types.YLeaf{"PceErrorRx", detailPcepInformation.PceErrorRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-error-tx", types.YLeaf{"PceErrorTx", detailPcepInformation.PceErrorTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-open-tx", types.YLeaf{"PceOpenTx", detailPcepInformation.PceOpenTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-open-rx", types.YLeaf{"PceOpenRx", detailPcepInformation.PceOpenRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-report-rx", types.YLeaf{"PceReportRx", detailPcepInformation.PceReportRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-report-tx", types.YLeaf{"PceReportTx", detailPcepInformation.PceReportTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-update-rx", types.YLeaf{"PceUpdateRx", detailPcepInformation.PceUpdateRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-update-tx", types.YLeaf{"PceUpdateTx", detailPcepInformation.PceUpdateTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-initiate-rx", types.YLeaf{"PceInitiateRx", detailPcepInformation.PceInitiateRx})
    detailPcepInformation.EntityData.Leafs.Append("pce-initiate-tx", types.YLeaf{"PceInitiateTx", detailPcepInformation.PceInitiateTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-keepalive-tx", types.YLeaf{"PceKeepaliveTx", detailPcepInformation.PceKeepaliveTx})
    detailPcepInformation.EntityData.Leafs.Append("pce-keepalive-rx", types.YLeaf{"PceKeepaliveRx", detailPcepInformation.PceKeepaliveRx})
    detailPcepInformation.EntityData.Leafs.Append("local-session-id", types.YLeaf{"LocalSessionId", detailPcepInformation.LocalSessionId})
    detailPcepInformation.EntityData.Leafs.Append("remote-session-id", types.YLeaf{"RemoteSessionId", detailPcepInformation.RemoteSessionId})
    detailPcepInformation.EntityData.Leafs.Append("minimum-keepalive-interval", types.YLeaf{"MinimumKeepaliveInterval", detailPcepInformation.MinimumKeepaliveInterval})
    detailPcepInformation.EntityData.Leafs.Append("maximum-dead-interval", types.YLeaf{"MaximumDeadInterval", detailPcepInformation.MaximumDeadInterval})
    detailPcepInformation.EntityData.Leafs.Append("max-sid-depth", types.YLeaf{"MaxSidDepth", detailPcepInformation.MaxSidDepth})

    detailPcepInformation.EntityData.YListKeys = []string {}

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

    briefPcepInformation.EntityData.Children = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs.Append("pcep-state", types.YLeaf{"PcepState", briefPcepInformation.PcepState})
    briefPcepInformation.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", briefPcepInformation.Stateful})
    briefPcepInformation.EntityData.Leafs.Append("capability-update", types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate})
    briefPcepInformation.EntityData.Leafs.Append("capability-instantiate", types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate})
    briefPcepInformation.EntityData.Leafs.Append("capability-segment-routing", types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting})
    briefPcepInformation.EntityData.Leafs.Append("capability-triggered-sync", types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync})
    briefPcepInformation.EntityData.Leafs.Append("capability-db-version", types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion})
    briefPcepInformation.EntityData.Leafs.Append("capability-delta-sync", types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync})

    briefPcepInformation.EntityData.YListKeys = []string {}

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

    lastErrorRx.EntityData.Children = types.NewOrderedMap()
    lastErrorRx.EntityData.Leafs = types.NewOrderedMap()
    lastErrorRx.EntityData.Leafs.Append("pc-error-type", types.YLeaf{"PcErrorType", lastErrorRx.PcErrorType})
    lastErrorRx.EntityData.Leafs.Append("pc-error-value", types.YLeaf{"PcErrorValue", lastErrorRx.PcErrorValue})

    lastErrorRx.EntityData.YListKeys = []string {}

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

    lastErrorTx.EntityData.Children = types.NewOrderedMap()
    lastErrorTx.EntityData.Leafs = types.NewOrderedMap()
    lastErrorTx.EntityData.Leafs.Append("pc-error-type", types.YLeaf{"PcErrorType", lastErrorTx.PcErrorType})
    lastErrorTx.EntityData.Leafs.Append("pc-error-value", types.YLeaf{"PcErrorValue", lastErrorTx.PcErrorValue})

    lastErrorTx.EntityData.YListKeys = []string {}

    return &(lastErrorTx.EntityData)
}

// Pce_TopologyNodes
// Node database in XTC
type Pce_TopologyNodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node information. The type is slice of Pce_TopologyNodes_TopologyNode.
    TopologyNode []*Pce_TopologyNodes_TopologyNode
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

    topologyNodes.EntityData.Children = types.NewOrderedMap()
    topologyNodes.EntityData.Children.Append("topology-node", types.YChild{"TopologyNode", nil})
    for i := range topologyNodes.TopologyNode {
        topologyNodes.EntityData.Children.Append(types.GetSegmentPath(topologyNodes.TopologyNode[i]), types.YChild{"TopologyNode", topologyNodes.TopologyNode[i]})
    }
    topologyNodes.EntityData.Leafs = types.NewOrderedMap()

    topologyNodes.EntityData.YListKeys = []string {}

    return &(topologyNodes.EntityData)
}

// Pce_TopologyNodes_TopologyNode
// Node information
type Pce_TopologyNodes_TopologyNode struct {
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
    NodeProtocolIdentifier Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier

    // Prefixes. The type is slice of Pce_TopologyNodes_TopologyNode_Prefixe.
    Prefixe []*Pce_TopologyNodes_TopologyNode_Prefixe

    // IPv4 Link information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link.
    Ipv4Link []*Pce_TopologyNodes_TopologyNode_Ipv4Link

    // IPv6 Link information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link.
    Ipv6Link []*Pce_TopologyNodes_TopologyNode_Ipv6Link
}

func (topologyNode *Pce_TopologyNodes_TopologyNode) GetEntityData() *types.CommonEntityData {
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
    topologyNode.EntityData.Children.Append("prefixe", types.YChild{"Prefixe", nil})
    for i := range topologyNode.Prefixe {
        topologyNode.EntityData.Children.Append(types.GetSegmentPath(topologyNode.Prefixe[i]), types.YChild{"Prefixe", topologyNode.Prefixe[i]})
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation
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

    nodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", nodeProtocolIdentifier.IgpInformation[i]})
    }
    nodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", nodeProtocolIdentifier.SrgbInformation[i]})
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Prefixe
// Prefixes
type Pce_TopologyNodes_TopologyNode_Prefixe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Prefix SID.
    PfxSid Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Prefixe_NodeId
}

func (prefixe *Pce_TopologyNodes_TopologyNode_Prefixe) GetEntityData() *types.CommonEntityData {
    prefixe.EntityData.YFilter = prefixe.YFilter
    prefixe.EntityData.YangName = "prefixe"
    prefixe.EntityData.BundleName = "cisco_ios_xr"
    prefixe.EntityData.ParentYangName = "topology-node"
    prefixe.EntityData.SegmentPath = "prefixe"
    prefixe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prefixe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prefixe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prefixe.EntityData.Children = types.NewOrderedMap()
    prefixe.EntityData.Children.Append("pfx-sid", types.YChild{"PfxSid", &prefixe.PfxSid})
    prefixe.EntityData.Children.Append("node-id", types.YChild{"NodeId", &prefixe.NodeId})
    prefixe.EntityData.Leafs = types.NewOrderedMap()
    prefixe.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", prefixe.DomainIdentifier})

    prefixe.EntityData.YListKeys = []string {}

    return &(prefixe.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid
// Prefix SID
type Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SID Type. The type is Sid.
    SidType interface{}

    // MPLS Label. The type is interface{} with range: 0..4294967295.
    MplsLabel interface{}

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
    SidPrefix Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix
}

func (pfxSid *Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid) GetEntityData() *types.CommonEntityData {
    pfxSid.EntityData.YFilter = pfxSid.YFilter
    pfxSid.EntityData.YangName = "pfx-sid"
    pfxSid.EntityData.BundleName = "cisco_ios_xr"
    pfxSid.EntityData.ParentYangName = "prefixe"
    pfxSid.EntityData.SegmentPath = "pfx-sid"
    pfxSid.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pfxSid.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pfxSid.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pfxSid.EntityData.Children = types.NewOrderedMap()
    pfxSid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &pfxSid.SidPrefix})
    pfxSid.EntityData.Leafs = types.NewOrderedMap()
    pfxSid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", pfxSid.SidType})
    pfxSid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", pfxSid.MplsLabel})
    pfxSid.EntityData.Leafs.Append("rflag", types.YLeaf{"Rflag", pfxSid.Rflag})
    pfxSid.EntityData.Leafs.Append("nflag", types.YLeaf{"Nflag", pfxSid.Nflag})
    pfxSid.EntityData.Leafs.Append("pflag", types.YLeaf{"Pflag", pfxSid.Pflag})
    pfxSid.EntityData.Leafs.Append("eflag", types.YLeaf{"Eflag", pfxSid.Eflag})
    pfxSid.EntityData.Leafs.Append("vflag", types.YLeaf{"Vflag", pfxSid.Vflag})
    pfxSid.EntityData.Leafs.Append("lflag", types.YLeaf{"Lflag", pfxSid.Lflag})

    pfxSid.EntityData.YListKeys = []string {}

    return &(pfxSid.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix
// Prefix
type Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix struct {
    EntityData types.CommonEntityData
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

func (sidPrefix *Pce_TopologyNodes_TopologyNode_Prefixe_PfxSid_SidPrefix) GetEntityData() *types.CommonEntityData {
    sidPrefix.EntityData.YFilter = sidPrefix.YFilter
    sidPrefix.EntityData.YangName = "sid-prefix"
    sidPrefix.EntityData.BundleName = "cisco_ios_xr"
    sidPrefix.EntityData.ParentYangName = "pfx-sid"
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

// Pce_TopologyNodes_TopologyNode_Prefixe_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Prefixe_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Prefixe_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "prefixe"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Prefixe_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link
// IPv4 Link information
type Pce_TopologyNodes_TopologyNode_Ipv4Link struct {
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
    LocalIgpInformation Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier

    // Performance metrics.
    PerformanceMetrics Pce_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics

    // Adjacency SIDs. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid.
    AdjacencySid []*Pce_TopologyNodes_TopologyNode_Ipv4Link_AdjacencySid
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

    ipv4Link.EntityData.Children = types.NewOrderedMap()
    ipv4Link.EntityData.Children.Append("local-igp-information", types.YChild{"LocalIgpInformation", &ipv4Link.LocalIgpInformation})
    ipv4Link.EntityData.Children.Append("remote-node-protocol-identifier", types.YChild{"RemoteNodeProtocolIdentifier", &ipv4Link.RemoteNodeProtocolIdentifier})
    ipv4Link.EntityData.Children.Append("performance-metrics", types.YChild{"PerformanceMetrics", &ipv4Link.PerformanceMetrics})
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation
// Local node IGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId
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

    localIgpInformation.EntityData.Children = types.NewOrderedMap()
    localIgpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &localIgpInformation.NodeId})
    localIgpInformation.EntityData.Leafs = types.NewOrderedMap()
    localIgpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier})

    localIgpInformation.EntityData.YListKeys = []string {}

    return &(localIgpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "local-igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_LocalIgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation
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

    remoteNodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", remoteNodeProtocolIdentifier.IgpInformation[i]})
    }
    remoteNodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", remoteNodeProtocolIdentifier.SrgbInformation[i]})
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv4Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics
// Performance metrics
type Pce_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Min delay. The type is interface{} with range: 0..4294967295.
    UnidirectionalMinDelay interface{}
}

func (performanceMetrics *Pce_TopologyNodes_TopologyNode_Ipv4Link_PerformanceMetrics) GetEntityData() *types.CommonEntityData {
    performanceMetrics.EntityData.YFilter = performanceMetrics.YFilter
    performanceMetrics.EntityData.YangName = "performance-metrics"
    performanceMetrics.EntityData.BundleName = "cisco_ios_xr"
    performanceMetrics.EntityData.ParentYangName = "ipv4-link"
    performanceMetrics.EntityData.SegmentPath = "performance-metrics"
    performanceMetrics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    performanceMetrics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    performanceMetrics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    performanceMetrics.EntityData.Children = types.NewOrderedMap()
    performanceMetrics.EntityData.Leafs = types.NewOrderedMap()
    performanceMetrics.EntityData.Leafs.Append("unidirectional-min-delay", types.YLeaf{"UnidirectionalMinDelay", performanceMetrics.UnidirectionalMinDelay})

    performanceMetrics.EntityData.YListKeys = []string {}

    return &(performanceMetrics.EntityData)
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

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &adjacencySid.SidPrefix})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()
    adjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", adjacencySid.SidType})
    adjacencySid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", adjacencySid.MplsLabel})
    adjacencySid.EntityData.Leafs.Append("rflag", types.YLeaf{"Rflag", adjacencySid.Rflag})
    adjacencySid.EntityData.Leafs.Append("nflag", types.YLeaf{"Nflag", adjacencySid.Nflag})
    adjacencySid.EntityData.Leafs.Append("pflag", types.YLeaf{"Pflag", adjacencySid.Pflag})
    adjacencySid.EntityData.Leafs.Append("eflag", types.YLeaf{"Eflag", adjacencySid.Eflag})
    adjacencySid.EntityData.Leafs.Append("vflag", types.YLeaf{"Vflag", adjacencySid.Vflag})
    adjacencySid.EntityData.Leafs.Append("lflag", types.YLeaf{"Lflag", adjacencySid.Lflag})

    adjacencySid.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link
// IPv6 Link information
type Pce_TopologyNodes_TopologyNode_Ipv6Link struct {
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
    LocalIgpInformation Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation

    // Remote node protocol identifier.
    RemoteNodeProtocolIdentifier Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier

    // Adjacency SIDs. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid.
    AdjacencySid []*Pce_TopologyNodes_TopologyNode_Ipv6Link_AdjacencySid
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation
// Local node IGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId
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

    localIgpInformation.EntityData.Children = types.NewOrderedMap()
    localIgpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &localIgpInformation.NodeId})
    localIgpInformation.EntityData.Leafs = types.NewOrderedMap()
    localIgpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", localIgpInformation.DomainIdentifier})

    localIgpInformation.EntityData.YListKeys = []string {}

    return &(localIgpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "local-igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_LocalIgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation
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

    remoteNodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    remoteNodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range remoteNodeProtocolIdentifier.IgpInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", remoteNodeProtocolIdentifier.IgpInformation[i]})
    }
    remoteNodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range remoteNodeProtocolIdentifier.SrgbInformation {
        remoteNodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(remoteNodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", remoteNodeProtocolIdentifier.SrgbInformation[i]})
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_TopologyNodes_TopologyNode_Ipv6Link_RemoteNodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    adjacencySid.EntityData.Children = types.NewOrderedMap()
    adjacencySid.EntityData.Children.Append("sid-prefix", types.YChild{"SidPrefix", &adjacencySid.SidPrefix})
    adjacencySid.EntityData.Leafs = types.NewOrderedMap()
    adjacencySid.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", adjacencySid.SidType})
    adjacencySid.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", adjacencySid.MplsLabel})
    adjacencySid.EntityData.Leafs.Append("rflag", types.YLeaf{"Rflag", adjacencySid.Rflag})
    adjacencySid.EntityData.Leafs.Append("nflag", types.YLeaf{"Nflag", adjacencySid.Nflag})
    adjacencySid.EntityData.Leafs.Append("pflag", types.YLeaf{"Pflag", adjacencySid.Pflag})
    adjacencySid.EntityData.Leafs.Append("eflag", types.YLeaf{"Eflag", adjacencySid.Eflag})
    adjacencySid.EntityData.Leafs.Append("vflag", types.YLeaf{"Vflag", adjacencySid.Vflag})
    adjacencySid.EntityData.Leafs.Append("lflag", types.YLeaf{"Lflag", adjacencySid.Lflag})

    adjacencySid.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    sidPrefix.EntityData.Children = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs = types.NewOrderedMap()
    sidPrefix.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sidPrefix.AfName})
    sidPrefix.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sidPrefix.Ipv4})
    sidPrefix.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sidPrefix.Ipv6})

    sidPrefix.EntityData.YListKeys = []string {}

    return &(sidPrefix.EntityData)
}

// Pce_PrefixInfos
// Prefixes database in XTC
type Pce_PrefixInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE prefix information. The type is slice of Pce_PrefixInfos_PrefixInfo.
    PrefixInfo []*Pce_PrefixInfos_PrefixInfo
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

    prefixInfos.EntityData.Children = types.NewOrderedMap()
    prefixInfos.EntityData.Children.Append("prefix-info", types.YChild{"PrefixInfo", nil})
    for i := range prefixInfos.PrefixInfo {
        prefixInfos.EntityData.Children.Append(types.GetSegmentPath(prefixInfos.PrefixInfo[i]), types.YChild{"PrefixInfo", prefixInfos.PrefixInfo[i]})
    }
    prefixInfos.EntityData.Leafs = types.NewOrderedMap()

    prefixInfos.EntityData.YListKeys = []string {}

    return &(prefixInfos.EntityData)
}

// Pce_PrefixInfos_PrefixInfo
// PCE prefix information
type Pce_PrefixInfos_PrefixInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is interface{} with range:
    // 0..4294967295.
    NodeIdentifier interface{}

    // Node identifier. The type is interface{} with range: 0..4294967295.
    NodeIdentifierXr interface{}

    // Node protocol identifier.
    NodeProtocolIdentifier Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier

    // Prefix address. The type is slice of Pce_PrefixInfos_PrefixInfo_Address.
    Address []*Pce_PrefixInfos_PrefixInfo_Address
}

func (prefixInfo *Pce_PrefixInfos_PrefixInfo) GetEntityData() *types.CommonEntityData {
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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4BgpRouterId interface{}

    // True if IPv4 TE router ID is set. The type is bool.
    Ipv4teRouterIdSet interface{}

    // IPv4 BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4teRouterId interface{}

    // IGP information. The type is slice of
    // Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation.
    IgpInformation []*Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation

    // SRGB information. The type is slice of
    // Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation.
    SrgbInformation []*Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation
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

    nodeProtocolIdentifier.EntityData.Children = types.NewOrderedMap()
    nodeProtocolIdentifier.EntityData.Children.Append("igp-information", types.YChild{"IgpInformation", nil})
    for i := range nodeProtocolIdentifier.IgpInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.IgpInformation[i]), types.YChild{"IgpInformation", nodeProtocolIdentifier.IgpInformation[i]})
    }
    nodeProtocolIdentifier.EntityData.Children.Append("srgb-information", types.YChild{"SrgbInformation", nil})
    for i := range nodeProtocolIdentifier.SrgbInformation {
        nodeProtocolIdentifier.EntityData.Children.Append(types.GetSegmentPath(nodeProtocolIdentifier.SrgbInformation[i]), types.YChild{"SrgbInformation", nodeProtocolIdentifier.SrgbInformation[i]})
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation
// IGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId
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

    igpInformation.EntityData.Children = types.NewOrderedMap()
    igpInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &igpInformation.NodeId})
    igpInformation.EntityData.Leafs = types.NewOrderedMap()
    igpInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", igpInformation.DomainIdentifier})

    igpInformation.EntityData.YListKeys = []string {}

    return &(igpInformation.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId
// Link-state node identifier
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
}

func (nodeId *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "igp-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp
// IGP-specific information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis
// ISIS information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp
// BGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_IgpInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    // Domain identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DomainIdentifier interface{}

    // Link-state node identifier.
    NodeId Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId
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

    srgbInformation.EntityData.Children = types.NewOrderedMap()
    srgbInformation.EntityData.Children.Append("node-id", types.YChild{"NodeId", &srgbInformation.NodeId})
    srgbInformation.EntityData.Leafs = types.NewOrderedMap()
    srgbInformation.EntityData.Leafs.Append("start", types.YLeaf{"Start", srgbInformation.Start})
    srgbInformation.EntityData.Leafs.Append("size", types.YLeaf{"Size", srgbInformation.Size})
    srgbInformation.EntityData.Leafs.Append("domain-identifier", types.YLeaf{"DomainIdentifier", srgbInformation.DomainIdentifier})

    srgbInformation.EntityData.YListKeys = []string {}

    return &(srgbInformation.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId
// Link-state node identifier
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Autonomous System Number. The type is interface{} with range:
    // 0..4294967295.
    AutonomousSystemNumber interface{}

    // Link-State identifier. The type is interface{} with range: 0..4294967295.
    LsIdentifier interface{}

    // IGP-specific information.
    Igp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
}

func (nodeId *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId) GetEntityData() *types.CommonEntityData {
    nodeId.EntityData.YFilter = nodeId.YFilter
    nodeId.EntityData.YangName = "node-id"
    nodeId.EntityData.BundleName = "cisco_ios_xr"
    nodeId.EntityData.ParentYangName = "srgb-information"
    nodeId.EntityData.SegmentPath = "node-id"
    nodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodeId.EntityData.Children = types.NewOrderedMap()
    nodeId.EntityData.Children.Append("igp", types.YChild{"Igp", &nodeId.Igp})
    nodeId.EntityData.Leafs = types.NewOrderedMap()
    nodeId.EntityData.Leafs.Append("autonomous-system-number", types.YLeaf{"AutonomousSystemNumber", nodeId.AutonomousSystemNumber})
    nodeId.EntityData.Leafs.Append("ls-identifier", types.YLeaf{"LsIdentifier", nodeId.LsIdentifier})

    nodeId.EntityData.YListKeys = []string {}

    return &(nodeId.EntityData)
}

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp
// IGP-specific information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IGP ID. The type is PceIgpInfoId.
    IgpId interface{}

    // ISIS information.
    Isis Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis

    // OSPF information.
    Ospf Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf

    // BGP information.
    Bgp Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
}

func (igp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp) GetEntityData() *types.CommonEntityData {
    igp.EntityData.YFilter = igp.YFilter
    igp.EntityData.YangName = "igp"
    igp.EntityData.BundleName = "cisco_ios_xr"
    igp.EntityData.ParentYangName = "node-id"
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis
// ISIS information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ISIS system ID. The type is string.
    SystemId interface{}

    // ISIS level. The type is interface{} with range: 0..4294967295.
    Level interface{}
}

func (isis *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Isis) GetEntityData() *types.CommonEntityData {
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf
// OSPF information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // OSPF area. The type is interface{} with range: 0..4294967295.
    Area interface{}
}

func (ospf *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Ospf) GetEntityData() *types.CommonEntityData {
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

// Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp
// BGP information
type Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // BGP router ID. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    RouterId interface{}

    // Confederation ASN. The type is interface{} with range: 0..4294967295.
    ConfedAsn interface{}
}

func (bgp *Pce_PrefixInfos_PrefixInfo_NodeProtocolIdentifier_SrgbInformation_NodeId_Igp_Bgp) GetEntityData() *types.CommonEntityData {
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
    bgp.EntityData.Leafs.Append("confed-asn", types.YLeaf{"ConfedAsn", bgp.ConfedAsn})

    bgp.EntityData.YListKeys = []string {}

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

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Children.Append("ip", types.YChild{"Ip", &address.Ip})
    address.EntityData.Leafs = types.NewOrderedMap()

    address.EntityData.YListKeys = []string {}

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
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
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

    ip.EntityData.Children = types.NewOrderedMap()
    ip.EntityData.Leafs = types.NewOrderedMap()
    ip.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", ip.AfName})
    ip.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ip.Ipv4})
    ip.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ip.Ipv6})

    ip.EntityData.YListKeys = []string {}

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
    PeerLsPsInfo []*Pce_LspSummary_PeerLsPsInfo
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

    lspSummary.EntityData.Children = types.NewOrderedMap()
    lspSummary.EntityData.Children.Append("all-ls-ps", types.YChild{"AllLsPs", &lspSummary.AllLsPs})
    lspSummary.EntityData.Children.Append("peer-ls-ps-info", types.YChild{"PeerLsPsInfo", nil})
    for i := range lspSummary.PeerLsPsInfo {
        lspSummary.EntityData.Children.Append(types.GetSegmentPath(lspSummary.PeerLsPsInfo[i]), types.YChild{"PeerLsPsInfo", lspSummary.PeerLsPsInfo[i]})
    }
    lspSummary.EntityData.Leafs = types.NewOrderedMap()

    lspSummary.EntityData.YListKeys = []string {}

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

    allLsPs.EntityData.Children = types.NewOrderedMap()
    allLsPs.EntityData.Leafs = types.NewOrderedMap()
    allLsPs.EntityData.Leafs.Append("all-ls-ps", types.YLeaf{"AllLsPs", allLsPs.AllLsPs})
    allLsPs.EntityData.Leafs.Append("up-ls-ps", types.YLeaf{"UpLsPs", allLsPs.UpLsPs})
    allLsPs.EntityData.Leafs.Append("admin-up-ls-ps", types.YLeaf{"AdminUpLsPs", allLsPs.AdminUpLsPs})
    allLsPs.EntityData.Leafs.Append("sr-ls-ps", types.YLeaf{"SrLsPs", allLsPs.SrLsPs})
    allLsPs.EntityData.Leafs.Append("rsvp-ls-ps", types.YLeaf{"RsvpLsPs", allLsPs.RsvpLsPs})

    allLsPs.EntityData.YListKeys = []string {}

    return &(allLsPs.EntityData)
}

// Pce_LspSummary_PeerLsPsInfo
// Number of LSPs for specific peer
type Pce_LspSummary_PeerLsPsInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of LSPs for specific peer.
    LspSummary Pce_LspSummary_PeerLsPsInfo_LspSummary

    // Peer address.
    PeerAddress Pce_LspSummary_PeerLsPsInfo_PeerAddress
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

    peerLsPsInfo.EntityData.Children = types.NewOrderedMap()
    peerLsPsInfo.EntityData.Children.Append("lsp-summary", types.YChild{"LspSummary", &peerLsPsInfo.LspSummary})
    peerLsPsInfo.EntityData.Children.Append("peer-address", types.YChild{"PeerAddress", &peerLsPsInfo.PeerAddress})
    peerLsPsInfo.EntityData.Leafs = types.NewOrderedMap()

    peerLsPsInfo.EntityData.YListKeys = []string {}

    return &(peerLsPsInfo.EntityData)
}

// Pce_LspSummary_PeerLsPsInfo_LspSummary
// Number of LSPs for specific peer
type Pce_LspSummary_PeerLsPsInfo_LspSummary struct {
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

func (lspSummary *Pce_LspSummary_PeerLsPsInfo_LspSummary) GetEntityData() *types.CommonEntityData {
    lspSummary.EntityData.YFilter = lspSummary.YFilter
    lspSummary.EntityData.YangName = "lsp-summary"
    lspSummary.EntityData.BundleName = "cisco_ios_xr"
    lspSummary.EntityData.ParentYangName = "peer-ls-ps-info"
    lspSummary.EntityData.SegmentPath = "lsp-summary"
    lspSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lspSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lspSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lspSummary.EntityData.Children = types.NewOrderedMap()
    lspSummary.EntityData.Leafs = types.NewOrderedMap()
    lspSummary.EntityData.Leafs.Append("all-ls-ps", types.YLeaf{"AllLsPs", lspSummary.AllLsPs})
    lspSummary.EntityData.Leafs.Append("up-ls-ps", types.YLeaf{"UpLsPs", lspSummary.UpLsPs})
    lspSummary.EntityData.Leafs.Append("admin-up-ls-ps", types.YLeaf{"AdminUpLsPs", lspSummary.AdminUpLsPs})
    lspSummary.EntityData.Leafs.Append("sr-ls-ps", types.YLeaf{"SrLsPs", lspSummary.SrLsPs})
    lspSummary.EntityData.Leafs.Append("rsvp-ls-ps", types.YLeaf{"RsvpLsPs", lspSummary.RsvpLsPs})

    lspSummary.EntityData.YListKeys = []string {}

    return &(lspSummary.EntityData)
}

// Pce_LspSummary_PeerLsPsInfo_PeerAddress
// Peer address
type Pce_LspSummary_PeerLsPsInfo_PeerAddress struct {
    EntityData types.CommonEntityData
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

func (peerAddress *Pce_LspSummary_PeerLsPsInfo_PeerAddress) GetEntityData() *types.CommonEntityData {
    peerAddress.EntityData.YFilter = peerAddress.YFilter
    peerAddress.EntityData.YangName = "peer-address"
    peerAddress.EntityData.BundleName = "cisco_ios_xr"
    peerAddress.EntityData.ParentYangName = "peer-ls-ps-info"
    peerAddress.EntityData.SegmentPath = "peer-address"
    peerAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddress.EntityData.Children = types.NewOrderedMap()
    peerAddress.EntityData.Leafs = types.NewOrderedMap()
    peerAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddress.AfName})
    peerAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerAddress.Ipv4})
    peerAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerAddress.Ipv6})

    peerAddress.EntityData.YListKeys = []string {}

    return &(peerAddress.EntityData)
}

// Pce_PeerInfos
// Peers database in XTC
type Pce_PeerInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCE peer information. The type is slice of Pce_PeerInfos_PeerInfo.
    PeerInfo []*Pce_PeerInfos_PeerInfo
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

    peerInfos.EntityData.Children = types.NewOrderedMap()
    peerInfos.EntityData.Children.Append("peer-info", types.YChild{"PeerInfo", nil})
    for i := range peerInfos.PeerInfo {
        peerInfos.EntityData.Children.Append(types.GetSegmentPath(peerInfos.PeerInfo[i]), types.YChild{"PeerInfo", peerInfos.PeerInfo[i]})
    }
    peerInfos.EntityData.Leafs = types.NewOrderedMap()

    peerInfos.EntityData.YListKeys = []string {}

    return &(peerInfos.EntityData)
}

// Pce_PeerInfos_PeerInfo
// PCE peer information
type Pce_PeerInfos_PeerInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // Protocol between PCE and peer. The type is PceProto.
    PeerProtocol interface{}

    // Peer address.
    PeerAddressXr Pce_PeerInfos_PeerInfo_PeerAddressXr

    // PCE protocol information.
    BriefPcepInformation Pce_PeerInfos_PeerInfo_BriefPcepInformation
}

func (peerInfo *Pce_PeerInfos_PeerInfo) GetEntityData() *types.CommonEntityData {
    peerInfo.EntityData.YFilter = peerInfo.YFilter
    peerInfo.EntityData.YangName = "peer-info"
    peerInfo.EntityData.BundleName = "cisco_ios_xr"
    peerInfo.EntityData.ParentYangName = "peer-infos"
    peerInfo.EntityData.SegmentPath = "peer-info" + types.AddKeyToken(peerInfo.PeerAddress, "peer-address")
    peerInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerInfo.EntityData.Children = types.NewOrderedMap()
    peerInfo.EntityData.Children.Append("peer-address-xr", types.YChild{"PeerAddressXr", &peerInfo.PeerAddressXr})
    peerInfo.EntityData.Children.Append("brief-pcep-information", types.YChild{"BriefPcepInformation", &peerInfo.BriefPcepInformation})
    peerInfo.EntityData.Leafs = types.NewOrderedMap()
    peerInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", peerInfo.PeerAddress})
    peerInfo.EntityData.Leafs.Append("peer-protocol", types.YLeaf{"PeerProtocol", peerInfo.PeerProtocol})

    peerInfo.EntityData.YListKeys = []string {"PeerAddress"}

    return &(peerInfo.EntityData)
}

// Pce_PeerInfos_PeerInfo_PeerAddressXr
// Peer address
type Pce_PeerInfos_PeerInfo_PeerAddressXr struct {
    EntityData types.CommonEntityData
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

func (peerAddressXr *Pce_PeerInfos_PeerInfo_PeerAddressXr) GetEntityData() *types.CommonEntityData {
    peerAddressXr.EntityData.YFilter = peerAddressXr.YFilter
    peerAddressXr.EntityData.YangName = "peer-address-xr"
    peerAddressXr.EntityData.BundleName = "cisco_ios_xr"
    peerAddressXr.EntityData.ParentYangName = "peer-info"
    peerAddressXr.EntityData.SegmentPath = "peer-address-xr"
    peerAddressXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerAddressXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerAddressXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerAddressXr.EntityData.Children = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs = types.NewOrderedMap()
    peerAddressXr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", peerAddressXr.AfName})
    peerAddressXr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", peerAddressXr.Ipv4})
    peerAddressXr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", peerAddressXr.Ipv6})

    peerAddressXr.EntityData.YListKeys = []string {}

    return &(peerAddressXr.EntityData)
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

    briefPcepInformation.EntityData.Children = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs = types.NewOrderedMap()
    briefPcepInformation.EntityData.Leafs.Append("pcep-state", types.YLeaf{"PcepState", briefPcepInformation.PcepState})
    briefPcepInformation.EntityData.Leafs.Append("stateful", types.YLeaf{"Stateful", briefPcepInformation.Stateful})
    briefPcepInformation.EntityData.Leafs.Append("capability-update", types.YLeaf{"CapabilityUpdate", briefPcepInformation.CapabilityUpdate})
    briefPcepInformation.EntityData.Leafs.Append("capability-instantiate", types.YLeaf{"CapabilityInstantiate", briefPcepInformation.CapabilityInstantiate})
    briefPcepInformation.EntityData.Leafs.Append("capability-segment-routing", types.YLeaf{"CapabilitySegmentRouting", briefPcepInformation.CapabilitySegmentRouting})
    briefPcepInformation.EntityData.Leafs.Append("capability-triggered-sync", types.YLeaf{"CapabilityTriggeredSync", briefPcepInformation.CapabilityTriggeredSync})
    briefPcepInformation.EntityData.Leafs.Append("capability-db-version", types.YLeaf{"CapabilityDbVersion", briefPcepInformation.CapabilityDbVersion})
    briefPcepInformation.EntityData.Leafs.Append("capability-delta-sync", types.YLeaf{"CapabilityDeltaSync", briefPcepInformation.CapabilityDeltaSync})

    briefPcepInformation.EntityData.YListKeys = []string {}

    return &(briefPcepInformation.EntityData)
}

// Pce_TunnelDetailInfos
// Detailed tunnel database in XTC
type Pce_TunnelDetailInfos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detailed tunnel information. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo.
    TunnelDetailInfo []*Pce_TunnelDetailInfos_TunnelDetailInfo
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

    tunnelDetailInfos.EntityData.Children = types.NewOrderedMap()
    tunnelDetailInfos.EntityData.Children.Append("tunnel-detail-info", types.YChild{"TunnelDetailInfo", nil})
    for i := range tunnelDetailInfos.TunnelDetailInfo {
        tunnelDetailInfos.EntityData.Children.Append(types.GetSegmentPath(tunnelDetailInfos.TunnelDetailInfo[i]), types.YChild{"TunnelDetailInfo", tunnelDetailInfos.TunnelDetailInfo[i]})
    }
    tunnelDetailInfos.EntityData.Leafs = types.NewOrderedMap()

    tunnelDetailInfos.EntityData.YListKeys = []string {}

    return &(tunnelDetailInfos.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo
// Detailed tunnel information
type Pce_TunnelDetailInfos_TunnelDetailInfo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Peer Address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    PeerAddress interface{}

    // This attribute is a key. PCEP LSP ID. The type is interface{} with range:
    // 0..4294967295.
    PlspId interface{}

    // This attribute is a key. Tunnel name. The type is string.
    TunnelName interface{}

    // Tunnel Name. The type is string.
    TunnelNameXr interface{}

    // Allow XTC reoptimizations. The type is bool.
    XtcControlled interface{}

    // Color. The type is interface{} with range: 0..4294967295.
    Color interface{}

    // PCC address.
    PccAddress Pce_TunnelDetailInfos_TunnelDetailInfo_PccAddress

    // Private LSP information.
    PrivateLspInformation Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation

    // Detail LSP information. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation.
    DetailLspInformation []*Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation
}

func (tunnelDetailInfo *Pce_TunnelDetailInfos_TunnelDetailInfo) GetEntityData() *types.CommonEntityData {
    tunnelDetailInfo.EntityData.YFilter = tunnelDetailInfo.YFilter
    tunnelDetailInfo.EntityData.YangName = "tunnel-detail-info"
    tunnelDetailInfo.EntityData.BundleName = "cisco_ios_xr"
    tunnelDetailInfo.EntityData.ParentYangName = "tunnel-detail-infos"
    tunnelDetailInfo.EntityData.SegmentPath = "tunnel-detail-info" + types.AddKeyToken(tunnelDetailInfo.PeerAddress, "peer-address") + types.AddKeyToken(tunnelDetailInfo.PlspId, "plsp-id") + types.AddKeyToken(tunnelDetailInfo.TunnelName, "tunnel-name")
    tunnelDetailInfo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDetailInfo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDetailInfo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDetailInfo.EntityData.Children = types.NewOrderedMap()
    tunnelDetailInfo.EntityData.Children.Append("pcc-address", types.YChild{"PccAddress", &tunnelDetailInfo.PccAddress})
    tunnelDetailInfo.EntityData.Children.Append("private-lsp-information", types.YChild{"PrivateLspInformation", &tunnelDetailInfo.PrivateLspInformation})
    tunnelDetailInfo.EntityData.Children.Append("detail-lsp-information", types.YChild{"DetailLspInformation", nil})
    for i := range tunnelDetailInfo.DetailLspInformation {
        tunnelDetailInfo.EntityData.Children.Append(types.GetSegmentPath(tunnelDetailInfo.DetailLspInformation[i]), types.YChild{"DetailLspInformation", tunnelDetailInfo.DetailLspInformation[i]})
    }
    tunnelDetailInfo.EntityData.Leafs = types.NewOrderedMap()
    tunnelDetailInfo.EntityData.Leafs.Append("peer-address", types.YLeaf{"PeerAddress", tunnelDetailInfo.PeerAddress})
    tunnelDetailInfo.EntityData.Leafs.Append("plsp-id", types.YLeaf{"PlspId", tunnelDetailInfo.PlspId})
    tunnelDetailInfo.EntityData.Leafs.Append("tunnel-name", types.YLeaf{"TunnelName", tunnelDetailInfo.TunnelName})
    tunnelDetailInfo.EntityData.Leafs.Append("tunnel-name-xr", types.YLeaf{"TunnelNameXr", tunnelDetailInfo.TunnelNameXr})
    tunnelDetailInfo.EntityData.Leafs.Append("xtc-controlled", types.YLeaf{"XtcControlled", tunnelDetailInfo.XtcControlled})
    tunnelDetailInfo.EntityData.Leafs.Append("color", types.YLeaf{"Color", tunnelDetailInfo.Color})

    tunnelDetailInfo.EntityData.YListKeys = []string {"PeerAddress", "PlspId", "TunnelName"}

    return &(tunnelDetailInfo.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_PccAddress
// PCC address
type Pce_TunnelDetailInfos_TunnelDetailInfo_PccAddress struct {
    EntityData types.CommonEntityData
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

func (pccAddress *Pce_TunnelDetailInfos_TunnelDetailInfo_PccAddress) GetEntityData() *types.CommonEntityData {
    pccAddress.EntityData.YFilter = pccAddress.YFilter
    pccAddress.EntityData.YangName = "pcc-address"
    pccAddress.EntityData.BundleName = "cisco_ios_xr"
    pccAddress.EntityData.ParentYangName = "tunnel-detail-info"
    pccAddress.EntityData.SegmentPath = "pcc-address"
    pccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pccAddress.EntityData.Children = types.NewOrderedMap()
    pccAddress.EntityData.Leafs = types.NewOrderedMap()
    pccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", pccAddress.AfName})
    pccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", pccAddress.Ipv4})
    pccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", pccAddress.Ipv6})

    pccAddress.EntityData.YListKeys = []string {}

    return &(pccAddress.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation
// Private LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSP Event buffer. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer.
    EventBuffer []*Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
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

    privateLspInformation.EntityData.Children = types.NewOrderedMap()
    privateLspInformation.EntityData.Children.Append("event-buffer", types.YChild{"EventBuffer", nil})
    for i := range privateLspInformation.EventBuffer {
        privateLspInformation.EntityData.Children.Append(types.GetSegmentPath(privateLspInformation.EventBuffer[i]), types.YChild{"EventBuffer", privateLspInformation.EventBuffer[i]})
    }
    privateLspInformation.EntityData.Leafs = types.NewOrderedMap()

    privateLspInformation.EntityData.YListKeys = []string {}

    return &(privateLspInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer
// LSP Event buffer
type Pce_TunnelDetailInfos_TunnelDetailInfo_PrivateLspInformation_EventBuffer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Event ID in range 1 - 0xFFFFFFFF. 0 is invalid. The type is interface{}
    // with range: 0..4294967295.
    EventId interface{}

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

    eventBuffer.EntityData.Children = types.NewOrderedMap()
    eventBuffer.EntityData.Leafs = types.NewOrderedMap()
    eventBuffer.EntityData.Leafs.Append("event-id", types.YLeaf{"EventId", eventBuffer.EventId})
    eventBuffer.EntityData.Leafs.Append("event-message", types.YLeaf{"EventMessage", eventBuffer.EventMessage})
    eventBuffer.EntityData.Leafs.Append("time-stamp", types.YLeaf{"TimeStamp", eventBuffer.TimeStamp})

    eventBuffer.EntityData.YListKeys = []string {}

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

    // Sub delegated PCE.
    SubDelegatedPce Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce

    // State-sync PCE.
    StateSyncPce Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce

    // Reporting PCC address.
    ReportingPccAddress Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress

    // RRO. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro.
    Rro []*Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
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

    detailLspInformation.EntityData.Children = types.NewOrderedMap()
    detailLspInformation.EntityData.Children.Append("brief-lsp-information", types.YChild{"BriefLspInformation", &detailLspInformation.BriefLspInformation})
    detailLspInformation.EntityData.Children.Append("er-os", types.YChild{"ErOs", &detailLspInformation.ErOs})
    detailLspInformation.EntityData.Children.Append("lsppcep-information", types.YChild{"LsppcepInformation", &detailLspInformation.LsppcepInformation})
    detailLspInformation.EntityData.Children.Append("lsp-association-info", types.YChild{"LspAssociationInfo", &detailLspInformation.LspAssociationInfo})
    detailLspInformation.EntityData.Children.Append("lsp-attributes", types.YChild{"LspAttributes", &detailLspInformation.LspAttributes})
    detailLspInformation.EntityData.Children.Append("sub-delegated-pce", types.YChild{"SubDelegatedPce", &detailLspInformation.SubDelegatedPce})
    detailLspInformation.EntityData.Children.Append("state-sync-pce", types.YChild{"StateSyncPce", &detailLspInformation.StateSyncPce})
    detailLspInformation.EntityData.Children.Append("reporting-pcc-address", types.YChild{"ReportingPccAddress", &detailLspInformation.ReportingPccAddress})
    detailLspInformation.EntityData.Children.Append("rro", types.YChild{"Rro", nil})
    for i := range detailLspInformation.Rro {
        detailLspInformation.EntityData.Children.Append(types.GetSegmentPath(detailLspInformation.Rro[i]), types.YChild{"Rro", detailLspInformation.Rro[i]})
    }
    detailLspInformation.EntityData.Leafs = types.NewOrderedMap()
    detailLspInformation.EntityData.Leafs.Append("signaled-bandwidth-specified", types.YLeaf{"SignaledBandwidthSpecified", detailLspInformation.SignaledBandwidthSpecified})
    detailLspInformation.EntityData.Leafs.Append("signaled-bandwidth", types.YLeaf{"SignaledBandwidth", detailLspInformation.SignaledBandwidth})
    detailLspInformation.EntityData.Leafs.Append("actual-bandwidth-specified", types.YLeaf{"ActualBandwidthSpecified", detailLspInformation.ActualBandwidthSpecified})
    detailLspInformation.EntityData.Leafs.Append("actual-bandwidth", types.YLeaf{"ActualBandwidth", detailLspInformation.ActualBandwidth})
    detailLspInformation.EntityData.Leafs.Append("lsp-role", types.YLeaf{"LspRole", detailLspInformation.LspRole})
    detailLspInformation.EntityData.Leafs.Append("computing-pce", types.YLeaf{"ComputingPce", detailLspInformation.ComputingPce})
    detailLspInformation.EntityData.Leafs.Append("srlg-info", types.YLeaf{"SrlgInfo", detailLspInformation.SrlgInfo})

    detailLspInformation.EntityData.YListKeys = []string {}

    return &(detailLspInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation
// Brief LSP information
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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

    // Maximum SID Depth. The type is interface{} with range: 0..4294967295.
    Msd interface{}

    // Source address.
    SourceAddress Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress

    // Destination address.
    DestinationAddress Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress
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

    briefLspInformation.EntityData.Children = types.NewOrderedMap()
    briefLspInformation.EntityData.Children.Append("source-address", types.YChild{"SourceAddress", &briefLspInformation.SourceAddress})
    briefLspInformation.EntityData.Children.Append("destination-address", types.YChild{"DestinationAddress", &briefLspInformation.DestinationAddress})
    briefLspInformation.EntityData.Leafs = types.NewOrderedMap()
    briefLspInformation.EntityData.Leafs.Append("tunnel-id", types.YLeaf{"TunnelId", briefLspInformation.TunnelId})
    briefLspInformation.EntityData.Leafs.Append("lspid", types.YLeaf{"Lspid", briefLspInformation.Lspid})
    briefLspInformation.EntityData.Leafs.Append("binding-sid", types.YLeaf{"BindingSid", briefLspInformation.BindingSid})
    briefLspInformation.EntityData.Leafs.Append("lsp-setup-type", types.YLeaf{"LspSetupType", briefLspInformation.LspSetupType})
    briefLspInformation.EntityData.Leafs.Append("operational-state", types.YLeaf{"OperationalState", briefLspInformation.OperationalState})
    briefLspInformation.EntityData.Leafs.Append("administrative-state", types.YLeaf{"AdministrativeState", briefLspInformation.AdministrativeState})
    briefLspInformation.EntityData.Leafs.Append("msd", types.YLeaf{"Msd", briefLspInformation.Msd})

    briefLspInformation.EntityData.YListKeys = []string {}

    return &(briefLspInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress
// Source address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress struct {
    EntityData types.CommonEntityData
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

func (sourceAddress *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_SourceAddress) GetEntityData() *types.CommonEntityData {
    sourceAddress.EntityData.YFilter = sourceAddress.YFilter
    sourceAddress.EntityData.YangName = "source-address"
    sourceAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceAddress.EntityData.ParentYangName = "brief-lsp-information"
    sourceAddress.EntityData.SegmentPath = "source-address"
    sourceAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceAddress.EntityData.Children = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs = types.NewOrderedMap()
    sourceAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", sourceAddress.AfName})
    sourceAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", sourceAddress.Ipv4})
    sourceAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", sourceAddress.Ipv6})

    sourceAddress.EntityData.YListKeys = []string {}

    return &(sourceAddress.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress
// Destination address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress struct {
    EntityData types.CommonEntityData
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

func (destinationAddress *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_BriefLspInformation_DestinationAddress) GetEntityData() *types.CommonEntityData {
    destinationAddress.EntityData.YFilter = destinationAddress.YFilter
    destinationAddress.EntityData.YangName = "destination-address"
    destinationAddress.EntityData.BundleName = "cisco_ios_xr"
    destinationAddress.EntityData.ParentYangName = "brief-lsp-information"
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
    ReportedRsvpPath []*Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath

    // Reported SR path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath.
    ReportedSrPath []*Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath

    // Computed RSVP path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath.
    ComputedRsvpPath []*Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath

    // Computed SR path. The type is slice of
    // Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath.
    ComputedSrPath []*Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath
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

    erOs.EntityData.Children = types.NewOrderedMap()
    erOs.EntityData.Children.Append("reported-rsvp-path", types.YChild{"ReportedRsvpPath", nil})
    for i := range erOs.ReportedRsvpPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ReportedRsvpPath[i]), types.YChild{"ReportedRsvpPath", erOs.ReportedRsvpPath[i]})
    }
    erOs.EntityData.Children.Append("reported-sr-path", types.YChild{"ReportedSrPath", nil})
    for i := range erOs.ReportedSrPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ReportedSrPath[i]), types.YChild{"ReportedSrPath", erOs.ReportedSrPath[i]})
    }
    erOs.EntityData.Children.Append("computed-rsvp-path", types.YChild{"ComputedRsvpPath", nil})
    for i := range erOs.ComputedRsvpPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ComputedRsvpPath[i]), types.YChild{"ComputedRsvpPath", erOs.ComputedRsvpPath[i]})
    }
    erOs.EntityData.Children.Append("computed-sr-path", types.YChild{"ComputedSrPath", nil})
    for i := range erOs.ComputedSrPath {
        erOs.EntityData.Children.Append(types.GetSegmentPath(erOs.ComputedSrPath[i]), types.YChild{"ComputedSrPath", erOs.ComputedSrPath[i]})
    }
    erOs.EntityData.Leafs = types.NewOrderedMap()
    erOs.EntityData.Leafs.Append("reported-metric-type", types.YLeaf{"ReportedMetricType", erOs.ReportedMetricType})
    erOs.EntityData.Leafs.Append("reported-metric-value", types.YLeaf{"ReportedMetricValue", erOs.ReportedMetricValue})
    erOs.EntityData.Leafs.Append("computed-metric-type", types.YLeaf{"ComputedMetricType", erOs.ComputedMetricType})
    erOs.EntityData.Leafs.Append("computed-metric-value", types.YLeaf{"ComputedMetricValue", erOs.ComputedMetricValue})
    erOs.EntityData.Leafs.Append("computed-hop-list-time", types.YLeaf{"ComputedHopListTime", erOs.ComputedHopListTime})

    erOs.EntityData.YListKeys = []string {}

    return &(erOs.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath
// Reported RSVP path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    reportedRsvpPath.EntityData.Children = types.NewOrderedMap()
    reportedRsvpPath.EntityData.Leafs = types.NewOrderedMap()
    reportedRsvpPath.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", reportedRsvpPath.HopAddress})

    reportedRsvpPath.EntityData.YListKeys = []string {}

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

    // Local Address.
    LocalAddr Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr

    // Remote Address.
    RemoteAddr Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr
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

    reportedSrPath.EntityData.Children = types.NewOrderedMap()
    reportedSrPath.EntityData.Children.Append("local-addr", types.YChild{"LocalAddr", &reportedSrPath.LocalAddr})
    reportedSrPath.EntityData.Children.Append("remote-addr", types.YChild{"RemoteAddr", &reportedSrPath.RemoteAddr})
    reportedSrPath.EntityData.Leafs = types.NewOrderedMap()
    reportedSrPath.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", reportedSrPath.SidType})
    reportedSrPath.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", reportedSrPath.MplsLabel})

    reportedSrPath.EntityData.YListKeys = []string {}

    return &(reportedSrPath.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr
// Local Address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr struct {
    EntityData types.CommonEntityData
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

func (localAddr *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_LocalAddr) GetEntityData() *types.CommonEntityData {
    localAddr.EntityData.YFilter = localAddr.YFilter
    localAddr.EntityData.YangName = "local-addr"
    localAddr.EntityData.BundleName = "cisco_ios_xr"
    localAddr.EntityData.ParentYangName = "reported-sr-path"
    localAddr.EntityData.SegmentPath = "local-addr"
    localAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddr.EntityData.Children = types.NewOrderedMap()
    localAddr.EntityData.Leafs = types.NewOrderedMap()
    localAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddr.AfName})
    localAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddr.Ipv4})
    localAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddr.Ipv6})

    localAddr.EntityData.YListKeys = []string {}

    return &(localAddr.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr
// Remote Address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr struct {
    EntityData types.CommonEntityData
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

func (remoteAddr *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ReportedSrPath_RemoteAddr) GetEntityData() *types.CommonEntityData {
    remoteAddr.EntityData.YFilter = remoteAddr.YFilter
    remoteAddr.EntityData.YangName = "remote-addr"
    remoteAddr.EntityData.BundleName = "cisco_ios_xr"
    remoteAddr.EntityData.ParentYangName = "reported-sr-path"
    remoteAddr.EntityData.SegmentPath = "remote-addr"
    remoteAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddr.EntityData.Children = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddr.AfName})
    remoteAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddr.Ipv4})
    remoteAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddr.Ipv6})

    remoteAddr.EntityData.YListKeys = []string {}

    return &(remoteAddr.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath
// Computed RSVP path
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedRsvpPath struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RSVP hop address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
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

    computedRsvpPath.EntityData.Children = types.NewOrderedMap()
    computedRsvpPath.EntityData.Leafs = types.NewOrderedMap()
    computedRsvpPath.EntityData.Leafs.Append("hop-address", types.YLeaf{"HopAddress", computedRsvpPath.HopAddress})

    computedRsvpPath.EntityData.YListKeys = []string {}

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

    // Local Address.
    LocalAddr Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr

    // Remote Address.
    RemoteAddr Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr
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

    computedSrPath.EntityData.Children = types.NewOrderedMap()
    computedSrPath.EntityData.Children.Append("local-addr", types.YChild{"LocalAddr", &computedSrPath.LocalAddr})
    computedSrPath.EntityData.Children.Append("remote-addr", types.YChild{"RemoteAddr", &computedSrPath.RemoteAddr})
    computedSrPath.EntityData.Leafs = types.NewOrderedMap()
    computedSrPath.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", computedSrPath.SidType})
    computedSrPath.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", computedSrPath.MplsLabel})

    computedSrPath.EntityData.YListKeys = []string {}

    return &(computedSrPath.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr
// Local Address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr struct {
    EntityData types.CommonEntityData
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

func (localAddr *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_LocalAddr) GetEntityData() *types.CommonEntityData {
    localAddr.EntityData.YFilter = localAddr.YFilter
    localAddr.EntityData.YangName = "local-addr"
    localAddr.EntityData.BundleName = "cisco_ios_xr"
    localAddr.EntityData.ParentYangName = "computed-sr-path"
    localAddr.EntityData.SegmentPath = "local-addr"
    localAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddr.EntityData.Children = types.NewOrderedMap()
    localAddr.EntityData.Leafs = types.NewOrderedMap()
    localAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddr.AfName})
    localAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddr.Ipv4})
    localAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddr.Ipv6})

    localAddr.EntityData.YListKeys = []string {}

    return &(localAddr.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr
// Remote Address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr struct {
    EntityData types.CommonEntityData
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

func (remoteAddr *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ErOs_ComputedSrPath_RemoteAddr) GetEntityData() *types.CommonEntityData {
    remoteAddr.EntityData.YFilter = remoteAddr.YFilter
    remoteAddr.EntityData.YangName = "remote-addr"
    remoteAddr.EntityData.BundleName = "cisco_ios_xr"
    remoteAddr.EntityData.ParentYangName = "computed-sr-path"
    remoteAddr.EntityData.SegmentPath = "remote-addr"
    remoteAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddr.EntityData.Children = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddr.AfName})
    remoteAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddr.Ipv4})
    remoteAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddr.Ipv6})

    remoteAddr.EntityData.YListKeys = []string {}

    return &(remoteAddr.EntityData)
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

    lsppcepInformation.EntityData.Children = types.NewOrderedMap()
    lsppcepInformation.EntityData.Children.Append("rsvp-error", types.YChild{"RsvpError", &lsppcepInformation.RsvpError})
    lsppcepInformation.EntityData.Leafs = types.NewOrderedMap()
    lsppcepInformation.EntityData.Leafs.Append("pcepid", types.YLeaf{"Pcepid", lsppcepInformation.Pcepid})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-d", types.YLeaf{"PcepFlagD", lsppcepInformation.PcepFlagD})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-s", types.YLeaf{"PcepFlagS", lsppcepInformation.PcepFlagS})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-r", types.YLeaf{"PcepFlagR", lsppcepInformation.PcepFlagR})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-a", types.YLeaf{"PcepFlagA", lsppcepInformation.PcepFlagA})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-o", types.YLeaf{"PcepFlagO", lsppcepInformation.PcepFlagO})
    lsppcepInformation.EntityData.Leafs.Append("pcep-flag-c", types.YLeaf{"PcepFlagC", lsppcepInformation.PcepFlagC})

    lsppcepInformation.EntityData.YListKeys = []string {}

    return &(lsppcepInformation.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError
// RSVP error info
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError struct {
    EntityData types.CommonEntityData
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

func (rsvpError *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LsppcepInformation_RsvpError) GetEntityData() *types.CommonEntityData {
    rsvpError.EntityData.YFilter = rsvpError.YFilter
    rsvpError.EntityData.YangName = "rsvp-error"
    rsvpError.EntityData.BundleName = "cisco_ios_xr"
    rsvpError.EntityData.ParentYangName = "lsppcep-information"
    rsvpError.EntityData.SegmentPath = "rsvp-error"
    rsvpError.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvpError.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvpError.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvpError.EntityData.Children = types.NewOrderedMap()
    rsvpError.EntityData.Leafs = types.NewOrderedMap()
    rsvpError.EntityData.Leafs.Append("node-address", types.YLeaf{"NodeAddress", rsvpError.NodeAddress})
    rsvpError.EntityData.Leafs.Append("error-flags", types.YLeaf{"ErrorFlags", rsvpError.ErrorFlags})
    rsvpError.EntityData.Leafs.Append("error-code", types.YLeaf{"ErrorCode", rsvpError.ErrorCode})
    rsvpError.EntityData.Leafs.Append("error-value", types.YLeaf{"ErrorValue", rsvpError.ErrorValue})

    rsvpError.EntityData.YListKeys = []string {}

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

    // Association Source.
    AssociationSource Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource
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

    lspAssociationInfo.EntityData.Children = types.NewOrderedMap()
    lspAssociationInfo.EntityData.Children.Append("association-source", types.YChild{"AssociationSource", &lspAssociationInfo.AssociationSource})
    lspAssociationInfo.EntityData.Leafs = types.NewOrderedMap()
    lspAssociationInfo.EntityData.Leafs.Append("association-type", types.YLeaf{"AssociationType", lspAssociationInfo.AssociationType})
    lspAssociationInfo.EntityData.Leafs.Append("association-id", types.YLeaf{"AssociationId", lspAssociationInfo.AssociationId})

    lspAssociationInfo.EntityData.YListKeys = []string {}

    return &(lspAssociationInfo.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource
// Association Source
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource struct {
    EntityData types.CommonEntityData
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

func (associationSource *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_LspAssociationInfo_AssociationSource) GetEntityData() *types.CommonEntityData {
    associationSource.EntityData.YFilter = associationSource.YFilter
    associationSource.EntityData.YangName = "association-source"
    associationSource.EntityData.BundleName = "cisco_ios_xr"
    associationSource.EntityData.ParentYangName = "lsp-association-info"
    associationSource.EntityData.SegmentPath = "association-source"
    associationSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    associationSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    associationSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    associationSource.EntityData.Children = types.NewOrderedMap()
    associationSource.EntityData.Leafs = types.NewOrderedMap()
    associationSource.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", associationSource.AfName})
    associationSource.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", associationSource.Ipv4})
    associationSource.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", associationSource.Ipv6})

    associationSource.EntityData.YListKeys = []string {}

    return &(associationSource.EntityData)
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

    lspAttributes.EntityData.Children = types.NewOrderedMap()
    lspAttributes.EntityData.Leafs = types.NewOrderedMap()
    lspAttributes.EntityData.Leafs.Append("affinity-exclude-any", types.YLeaf{"AffinityExcludeAny", lspAttributes.AffinityExcludeAny})
    lspAttributes.EntityData.Leafs.Append("affinity-include-any", types.YLeaf{"AffinityIncludeAny", lspAttributes.AffinityIncludeAny})
    lspAttributes.EntityData.Leafs.Append("affinity-include-all", types.YLeaf{"AffinityIncludeAll", lspAttributes.AffinityIncludeAll})
    lspAttributes.EntityData.Leafs.Append("setup-priority", types.YLeaf{"SetupPriority", lspAttributes.SetupPriority})
    lspAttributes.EntityData.Leafs.Append("hold-priority", types.YLeaf{"HoldPriority", lspAttributes.HoldPriority})
    lspAttributes.EntityData.Leafs.Append("local-protection", types.YLeaf{"LocalProtection", lspAttributes.LocalProtection})

    lspAttributes.EntityData.YListKeys = []string {}

    return &(lspAttributes.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce
// Sub delegated PCE
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce struct {
    EntityData types.CommonEntityData
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

func (subDelegatedPce *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_SubDelegatedPce) GetEntityData() *types.CommonEntityData {
    subDelegatedPce.EntityData.YFilter = subDelegatedPce.YFilter
    subDelegatedPce.EntityData.YangName = "sub-delegated-pce"
    subDelegatedPce.EntityData.BundleName = "cisco_ios_xr"
    subDelegatedPce.EntityData.ParentYangName = "detail-lsp-information"
    subDelegatedPce.EntityData.SegmentPath = "sub-delegated-pce"
    subDelegatedPce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subDelegatedPce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subDelegatedPce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subDelegatedPce.EntityData.Children = types.NewOrderedMap()
    subDelegatedPce.EntityData.Leafs = types.NewOrderedMap()
    subDelegatedPce.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", subDelegatedPce.AfName})
    subDelegatedPce.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", subDelegatedPce.Ipv4})
    subDelegatedPce.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", subDelegatedPce.Ipv6})

    subDelegatedPce.EntityData.YListKeys = []string {}

    return &(subDelegatedPce.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce
// State-sync PCE
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce struct {
    EntityData types.CommonEntityData
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

func (stateSyncPce *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_StateSyncPce) GetEntityData() *types.CommonEntityData {
    stateSyncPce.EntityData.YFilter = stateSyncPce.YFilter
    stateSyncPce.EntityData.YangName = "state-sync-pce"
    stateSyncPce.EntityData.BundleName = "cisco_ios_xr"
    stateSyncPce.EntityData.ParentYangName = "detail-lsp-information"
    stateSyncPce.EntityData.SegmentPath = "state-sync-pce"
    stateSyncPce.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSyncPce.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSyncPce.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSyncPce.EntityData.Children = types.NewOrderedMap()
    stateSyncPce.EntityData.Leafs = types.NewOrderedMap()
    stateSyncPce.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", stateSyncPce.AfName})
    stateSyncPce.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", stateSyncPce.Ipv4})
    stateSyncPce.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", stateSyncPce.Ipv6})

    stateSyncPce.EntityData.YListKeys = []string {}

    return &(stateSyncPce.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress
// Reporting PCC address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress struct {
    EntityData types.CommonEntityData
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

func (reportingPccAddress *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_ReportingPccAddress) GetEntityData() *types.CommonEntityData {
    reportingPccAddress.EntityData.YFilter = reportingPccAddress.YFilter
    reportingPccAddress.EntityData.YangName = "reporting-pcc-address"
    reportingPccAddress.EntityData.BundleName = "cisco_ios_xr"
    reportingPccAddress.EntityData.ParentYangName = "detail-lsp-information"
    reportingPccAddress.EntityData.SegmentPath = "reporting-pcc-address"
    reportingPccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reportingPccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reportingPccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reportingPccAddress.EntityData.Children = types.NewOrderedMap()
    reportingPccAddress.EntityData.Leafs = types.NewOrderedMap()
    reportingPccAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", reportingPccAddress.AfName})
    reportingPccAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", reportingPccAddress.Ipv4})
    reportingPccAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", reportingPccAddress.Ipv6})

    reportingPccAddress.EntityData.YListKeys = []string {}

    return &(reportingPccAddress.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro
// RRO
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro struct {
    EntityData types.CommonEntityData
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

func (rro *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro) GetEntityData() *types.CommonEntityData {
    rro.EntityData.YFilter = rro.YFilter
    rro.EntityData.YangName = "rro"
    rro.EntityData.BundleName = "cisco_ios_xr"
    rro.EntityData.ParentYangName = "detail-lsp-information"
    rro.EntityData.SegmentPath = "rro"
    rro.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rro.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rro.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rro.EntityData.Children = types.NewOrderedMap()
    rro.EntityData.Children.Append("sr-rro", types.YChild{"SrRro", &rro.SrRro})
    rro.EntityData.Leafs = types.NewOrderedMap()
    rro.EntityData.Leafs.Append("rro-type", types.YLeaf{"RroType", rro.RroType})
    rro.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", rro.Ipv4Address})
    rro.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", rro.MplsLabel})
    rro.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", rro.Flags})

    rro.EntityData.YListKeys = []string {}

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

    // Local Address.
    LocalAddr Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr

    // Remote Address.
    RemoteAddr Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr
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

    srRro.EntityData.Children = types.NewOrderedMap()
    srRro.EntityData.Children.Append("local-addr", types.YChild{"LocalAddr", &srRro.LocalAddr})
    srRro.EntityData.Children.Append("remote-addr", types.YChild{"RemoteAddr", &srRro.RemoteAddr})
    srRro.EntityData.Leafs = types.NewOrderedMap()
    srRro.EntityData.Leafs.Append("sid-type", types.YLeaf{"SidType", srRro.SidType})
    srRro.EntityData.Leafs.Append("mpls-label", types.YLeaf{"MplsLabel", srRro.MplsLabel})

    srRro.EntityData.YListKeys = []string {}

    return &(srRro.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr
// Local Address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr struct {
    EntityData types.CommonEntityData
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

func (localAddr *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_LocalAddr) GetEntityData() *types.CommonEntityData {
    localAddr.EntityData.YFilter = localAddr.YFilter
    localAddr.EntityData.YangName = "local-addr"
    localAddr.EntityData.BundleName = "cisco_ios_xr"
    localAddr.EntityData.ParentYangName = "sr-rro"
    localAddr.EntityData.SegmentPath = "local-addr"
    localAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddr.EntityData.Children = types.NewOrderedMap()
    localAddr.EntityData.Leafs = types.NewOrderedMap()
    localAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddr.AfName})
    localAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", localAddr.Ipv4})
    localAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", localAddr.Ipv6})

    localAddr.EntityData.YListKeys = []string {}

    return &(localAddr.EntityData)
}

// Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr
// Remote Address
type Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr struct {
    EntityData types.CommonEntityData
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

func (remoteAddr *Pce_TunnelDetailInfos_TunnelDetailInfo_DetailLspInformation_Rro_SrRro_RemoteAddr) GetEntityData() *types.CommonEntityData {
    remoteAddr.EntityData.YFilter = remoteAddr.YFilter
    remoteAddr.EntityData.YangName = "remote-addr"
    remoteAddr.EntityData.BundleName = "cisco_ios_xr"
    remoteAddr.EntityData.ParentYangName = "sr-rro"
    remoteAddr.EntityData.SegmentPath = "remote-addr"
    remoteAddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddr.EntityData.Children = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs = types.NewOrderedMap()
    remoteAddr.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddr.AfName})
    remoteAddr.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", remoteAddr.Ipv4})
    remoteAddr.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", remoteAddr.Ipv6})

    remoteAddr.EntityData.YListKeys = []string {}

    return &(remoteAddr.EntityData)
}

