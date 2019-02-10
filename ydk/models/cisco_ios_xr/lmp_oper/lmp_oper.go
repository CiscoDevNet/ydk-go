// This module contains a collection of YANG definitions
// for Cisco IOS-XR lmp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   lmp: Main common UCP/OLM operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package lmp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lmp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-lmp-oper lmp}", reflect.TypeOf(Lmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-lmp-oper:lmp", reflect.TypeOf(Lmp{}))
}

// OlmCompLinkImState represents The OLM Component link IM state
type OlmCompLinkImState string

const (
    // OIR removed
    OlmCompLinkImState_comp_link_im_state_oir OlmCompLinkImState = "comp-link-im-state-oir"

    // Down
    OlmCompLinkImState_comp_link_im_state_down OlmCompLinkImState = "comp-link-im-state-down"

    // Admin Down
    OlmCompLinkImState_comp_link_im_state_admin_down OlmCompLinkImState = "comp-link-im-state-admin-down"

    // Up
    OlmCompLinkImState_comp_link_im_state_up OlmCompLinkImState = "comp-link-im-state-up"

    // Unknown
    OlmCompLinkImState_comp_link_im_state_unknown OlmCompLinkImState = "comp-link-im-state-unknown"
)

// OlmLinkEncoding represents LMP link encoding type as defined in [RFC3471]
type OlmLinkEncoding string

const (
    // None
    OlmLinkEncoding_none OlmLinkEncoding = "none"

    // Packet
    OlmLinkEncoding_packet OlmLinkEncoding = "packet"

    // Ethernet
    OlmLinkEncoding_ethernet OlmLinkEncoding = "ethernet"

    // ANSI/ETSI PDH
    OlmLinkEncoding_ansi_etsi_pdh OlmLinkEncoding = "ansi-etsi-pdh"

    // Reserved
    OlmLinkEncoding_reserved1 OlmLinkEncoding = "reserved1"

    // SDH ITU-T G.707 / SONET ANSI T1.105
    OlmLinkEncoding_sdh_sonet OlmLinkEncoding = "sdh-sonet"

    // Reserved
    OlmLinkEncoding_reserved2 OlmLinkEncoding = "reserved2"

    // Digital Wrapper
    OlmLinkEncoding_digital_wrapper OlmLinkEncoding = "digital-wrapper"

    // Lambda (photonic)
    OlmLinkEncoding_lambda OlmLinkEncoding = "lambda"

    // Fiber
    OlmLinkEncoding_fiber OlmLinkEncoding = "fiber"

    // Reserved
    OlmLinkEncoding_reserved3 OlmLinkEncoding = "reserved3"

    // FiberChannel
    OlmLinkEncoding_fiber_channel OlmLinkEncoding = "fiber-channel"

    // Unknown
    OlmLinkEncoding_lencode_unknown OlmLinkEncoding = "lencode-unknown"
)

// OlmMuxCap represents Multiplexing capability
type OlmMuxCap string

const (
    // PSC 1
    OlmMuxCap_psc1 OlmMuxCap = "psc1"

    // PSC 2
    OlmMuxCap_psc2 OlmMuxCap = "psc2"

    // PSC 3
    OlmMuxCap_psc3 OlmMuxCap = "psc3"

    // PSC 4
    OlmMuxCap_psc4 OlmMuxCap = "psc4"

    // L2SC
    OlmMuxCap_l2sc OlmMuxCap = "l2sc"

    // TDM
    OlmMuxCap_tdm OlmMuxCap = "tdm"

    // LSC
    OlmMuxCap_lsc OlmMuxCap = "lsc"

    // FSC
    OlmMuxCap_fsc OlmMuxCap = "fsc"

    // Unknown Mux Cap
    OlmMuxCap_unknown_mux_cap OlmMuxCap = "unknown-mux-cap"
)

// OlmCompLinkLmpStatus represents The component link LMP status
type OlmCompLinkLmpStatus string

const (
    // Component link IF ID mismatch
    OlmCompLinkLmpStatus_comp_link_lmp_status_if_id_mismatch OlmCompLinkLmpStatus = "comp-link-lmp-status-if-id-mismatch"

    // Component link switching capability ID mismatch
    OlmCompLinkLmpStatus_comp_link_lmp_status_switch_cap_mismatch OlmCompLinkLmpStatus = "comp-link-lmp-status-switch-cap-mismatch"
)

// OlmAddrTypeId represents OLM Link Address Type
type OlmAddrTypeId string

const (
    // Unknown address type
    OlmAddrTypeId_unknown_address OlmAddrTypeId = "unknown-address"

    // IPv4 address type
    OlmAddrTypeId_ipv4 OlmAddrTypeId = "ipv4"

    // IPv6 address type
    OlmAddrTypeId_ipv6 OlmAddrTypeId = "ipv6"

    // Unnumberedaddress type
    OlmAddrTypeId_unnumbered OlmAddrTypeId = "unnumbered"
)

// OlmCompLinkLmpState represents The OLM Component link LMP state
type OlmCompLinkLmpState string

const (
    // Down
    OlmCompLinkLmpState_comp_link_lmp_state_down OlmCompLinkLmpState = "comp-link-lmp-state-down"

    // Test
    OlmCompLinkLmpState_comp_link_lmp_state_test OlmCompLinkLmpState = "comp-link-lmp-state-test"

    // Pasv Test
    OlmCompLinkLmpState_comp_link_lmp_state_passive_test OlmCompLinkLmpState = "comp-link-lmp-state-passive-test"

    // Up Free
    OlmCompLinkLmpState_comp_link_lmp_state_up_free OlmCompLinkLmpState = "comp-link-lmp-state-up-free"

    // Up Allocated
    OlmCompLinkLmpState_comp_link_lmp_state_up_allocated OlmCompLinkLmpState = "comp-link-lmp-state-up-allocated"

    // Unknown/Invalid
    OlmCompLinkLmpState_comp_link_lmp_state_unknown OlmCompLinkLmpState = "comp-link-lmp-state-unknown"
)

// OlmObjectOwner represents The OLM object owner
type OlmObjectOwner string

const (
    // Unknown owner
    OlmObjectOwner_unknown OlmObjectOwner = "unknown"

    // OIF OUNI
    OlmObjectOwner_ouni OlmObjectOwner = "ouni"

    // GMPLS NNI
    OlmObjectOwner_gmpls_nni OlmObjectOwner = "gmpls-nni"

    // GMPLS UNI
    OlmObjectOwner_gmpls_uni OlmObjectOwner = "gmpls-uni"
)

// OlmteLinkLmpState represents The OLM TE link LMP state
type OlmteLinkLmpState string

const (
    // Down
    OlmteLinkLmpState_te_link_lmp_state_down OlmteLinkLmpState = "te-link-lmp-state-down"

    // Init
    OlmteLinkLmpState_te_link_lmp_state_init OlmteLinkLmpState = "te-link-lmp-state-init"

    // Up
    OlmteLinkLmpState_te_link_lmp_state_up OlmteLinkLmpState = "te-link-lmp-state-up"

    // Degraded
    OlmteLinkLmpState_te_link_lmp_state_degraded OlmteLinkLmpState = "te-link-lmp-state-degraded"

    // Unknown/Invalid
    OlmteLinkLmpState_te_link_lmp_state_unknown OlmteLinkLmpState = "te-link-lmp-state-unknown"
)

// Olmipcc represents The OLM IPCC type
type Olmipcc string

const (
    // Global routed IPCC
    Olmipcc_ipcc_type_global_routed Olmipcc = "ipcc-type-global-routed"

    // Global I/F bound IPCC
    Olmipcc_ipcc_type_global_if_bound Olmipcc = "ipcc-type-global-if-bound"

    // SDCC/LDCC in fiber in band type IPCC
    Olmipcc_ipcc_type_ldcc_sdcc Olmipcc = "ipcc-type-ldcc-sdcc"

    // Unknown IPCC type
    Olmipcc_ipcc_type_unknown Olmipcc = "ipcc-type-unknown"
)

// OlmipccState represents The OLM IPCC state
type OlmipccState string

const (
    // OIR Removed
    OlmipccState_ipcc_state_oir_removed OlmipccState = "ipcc-state-oir-removed"

    // OOS
    OlmipccState_ipcc_state_admin_down OlmipccState = "ipcc-state-admin-down"

    // Down
    OlmipccState_ipcc_state_down OlmipccState = "ipcc-state-down"

    // ConfigSend
    OlmipccState_ipcc_state_cfg_send OlmipccState = "ipcc-state-cfg-send"

    // ConfigReceive
    OlmipccState_ipcc_state_cfg_rcv OlmipccState = "ipcc-state-cfg-rcv"

    // Active
    OlmipccState_ipcc_state_active OlmipccState = "ipcc-state-active"

    // Up
    OlmipccState_ipcc_state_up OlmipccState = "ipcc-state-up"

    // Going Down
    OlmipccState_ipcc_state_going_down OlmipccState = "ipcc-state-going-down"

    // Unknown/Invalid
    OlmipccState_ipcc_state_unknown OlmipccState = "ipcc-state-unknown"
)

// OlmRouterId represents The OLM router ID type
type OlmRouterId string

const (
    // No router ID configured
    OlmRouterId_not_configured OlmRouterId = "not-configured"

    // Global router ID
    OlmRouterId_global OlmRouterId = "global"

    // Protocol based CLIrouter ID configured
    OlmRouterId_protocol_based_address OlmRouterId = "protocol-based-address"

    // Protocol based CLI I/Frouter ID configured
    OlmRouterId_interface_ OlmRouterId = "interface"

    // Protocol based CLI I/F routerID configured on
    // I/F that is not known to the system
    OlmRouterId_network_element OlmRouterId = "network-element"

    // Unknown
    OlmRouterId_unknown_type OlmRouterId = "unknown-type"
)

// Lmp
// Main common UCP/OLM operational data
type Lmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Global OLM process information.
    GlobalStatus Lmp_GlobalStatus

    // UCP OLM clients container class.
    Clients Lmp_Clients

    // GMPLS UNI specific OLM/LMP configuration.
    GmplsUni Lmp_GmplsUni

    // UCP OLM component link ID container class.
    ComponentLinkIds Lmp_ComponentLinkIds
}

func (lmp *Lmp) GetEntityData() *types.CommonEntityData {
    lmp.EntityData.YFilter = lmp.YFilter
    lmp.EntityData.YangName = "lmp"
    lmp.EntityData.BundleName = "cisco_ios_xr"
    lmp.EntityData.ParentYangName = "Cisco-IOS-XR-lmp-oper"
    lmp.EntityData.SegmentPath = "Cisco-IOS-XR-lmp-oper:lmp"
    lmp.EntityData.AbsolutePath = lmp.EntityData.SegmentPath
    lmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lmp.EntityData.Children = types.NewOrderedMap()
    lmp.EntityData.Children.Append("global-status", types.YChild{"GlobalStatus", &lmp.GlobalStatus})
    lmp.EntityData.Children.Append("clients", types.YChild{"Clients", &lmp.Clients})
    lmp.EntityData.Children.Append("gmpls-uni", types.YChild{"GmplsUni", &lmp.GmplsUni})
    lmp.EntityData.Children.Append("component-link-ids", types.YChild{"ComponentLinkIds", &lmp.ComponentLinkIds})
    lmp.EntityData.Leafs = types.NewOrderedMap()

    lmp.EntityData.YListKeys = []string {}

    return &(lmp.EntityData)
}

// Lmp_GlobalStatus
// Global OLM process information
type Lmp_GlobalStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local OUNI LMP Node ID I/F. The type is string.
    LocalOuniLmpNodeIdInterface interface{}

    // Local OUNI LMP Node ID type. The type is OlmRouterId.
    LocalOuniLmpNodeIdType interface{}

    // TRUE if any OLM OUNI config exists. The type is bool.
    IsOuniConfigExist interface{}

    // TRUE if any OLM/LNP GMPLS NNI config exists. The type is bool.
    IsGmplsNniConfigExist interface{}

    // TRUE if any OLM/LMP GMPLS UNI config exists. The type is bool.
    IsGmplsUniConfigExist interface{}

    // Local OUNI LMP Node ID.
    LocalOuniLmpNodeId Lmp_GlobalStatus_LocalOuniLmpNodeId

    // MPLS TE LMP Node ID.
    LocalMplsTeLmpNodeId Lmp_GlobalStatus_LocalMplsTeLmpNodeId

    // GMPLS UNI LMP Node ID.
    LocalGmplsUniLmpNodeId Lmp_GlobalStatus_LocalGmplsUniLmpNodeId
}

func (globalStatus *Lmp_GlobalStatus) GetEntityData() *types.CommonEntityData {
    globalStatus.EntityData.YFilter = globalStatus.YFilter
    globalStatus.EntityData.YangName = "global-status"
    globalStatus.EntityData.BundleName = "cisco_ios_xr"
    globalStatus.EntityData.ParentYangName = "lmp"
    globalStatus.EntityData.SegmentPath = "global-status"
    globalStatus.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/" + globalStatus.EntityData.SegmentPath
    globalStatus.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalStatus.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalStatus.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalStatus.EntityData.Children = types.NewOrderedMap()
    globalStatus.EntityData.Children.Append("local-ouni-lmp-node-id", types.YChild{"LocalOuniLmpNodeId", &globalStatus.LocalOuniLmpNodeId})
    globalStatus.EntityData.Children.Append("local-mpls-te-lmp-node-id", types.YChild{"LocalMplsTeLmpNodeId", &globalStatus.LocalMplsTeLmpNodeId})
    globalStatus.EntityData.Children.Append("local-gmpls-uni-lmp-node-id", types.YChild{"LocalGmplsUniLmpNodeId", &globalStatus.LocalGmplsUniLmpNodeId})
    globalStatus.EntityData.Leafs = types.NewOrderedMap()
    globalStatus.EntityData.Leafs.Append("local-ouni-lmp-node-id-interface", types.YLeaf{"LocalOuniLmpNodeIdInterface", globalStatus.LocalOuniLmpNodeIdInterface})
    globalStatus.EntityData.Leafs.Append("local-ouni-lmp-node-id-type", types.YLeaf{"LocalOuniLmpNodeIdType", globalStatus.LocalOuniLmpNodeIdType})
    globalStatus.EntityData.Leafs.Append("is-ouni-config-exist", types.YLeaf{"IsOuniConfigExist", globalStatus.IsOuniConfigExist})
    globalStatus.EntityData.Leafs.Append("is-gmpls-nni-config-exist", types.YLeaf{"IsGmplsNniConfigExist", globalStatus.IsGmplsNniConfigExist})
    globalStatus.EntityData.Leafs.Append("is-gmpls-uni-config-exist", types.YLeaf{"IsGmplsUniConfigExist", globalStatus.IsGmplsUniConfigExist})

    globalStatus.EntityData.YListKeys = []string {}

    return &(globalStatus.EntityData)
}

// Lmp_GlobalStatus_LocalOuniLmpNodeId
// Local OUNI LMP Node ID
type Lmp_GlobalStatus_LocalOuniLmpNodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GlobalStatus_LocalOuniLmpNodeId_Address
}

func (localOuniLmpNodeId *Lmp_GlobalStatus_LocalOuniLmpNodeId) GetEntityData() *types.CommonEntityData {
    localOuniLmpNodeId.EntityData.YFilter = localOuniLmpNodeId.YFilter
    localOuniLmpNodeId.EntityData.YangName = "local-ouni-lmp-node-id"
    localOuniLmpNodeId.EntityData.BundleName = "cisco_ios_xr"
    localOuniLmpNodeId.EntityData.ParentYangName = "global-status"
    localOuniLmpNodeId.EntityData.SegmentPath = "local-ouni-lmp-node-id"
    localOuniLmpNodeId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/global-status/" + localOuniLmpNodeId.EntityData.SegmentPath
    localOuniLmpNodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localOuniLmpNodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localOuniLmpNodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localOuniLmpNodeId.EntityData.Children = types.NewOrderedMap()
    localOuniLmpNodeId.EntityData.Children.Append("address", types.YChild{"Address", &localOuniLmpNodeId.Address})
    localOuniLmpNodeId.EntityData.Leafs = types.NewOrderedMap()

    localOuniLmpNodeId.EntityData.YListKeys = []string {}

    return &(localOuniLmpNodeId.EntityData)
}

// Lmp_GlobalStatus_LocalOuniLmpNodeId_Address
// Address Union
type Lmp_GlobalStatus_LocalOuniLmpNodeId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GlobalStatus_LocalOuniLmpNodeId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-ouni-lmp-node-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/global-status/local-ouni-lmp-node-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GlobalStatus_LocalMplsTeLmpNodeId
// MPLS TE LMP Node ID
type Lmp_GlobalStatus_LocalMplsTeLmpNodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GlobalStatus_LocalMplsTeLmpNodeId_Address
}

func (localMplsTeLmpNodeId *Lmp_GlobalStatus_LocalMplsTeLmpNodeId) GetEntityData() *types.CommonEntityData {
    localMplsTeLmpNodeId.EntityData.YFilter = localMplsTeLmpNodeId.YFilter
    localMplsTeLmpNodeId.EntityData.YangName = "local-mpls-te-lmp-node-id"
    localMplsTeLmpNodeId.EntityData.BundleName = "cisco_ios_xr"
    localMplsTeLmpNodeId.EntityData.ParentYangName = "global-status"
    localMplsTeLmpNodeId.EntityData.SegmentPath = "local-mpls-te-lmp-node-id"
    localMplsTeLmpNodeId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/global-status/" + localMplsTeLmpNodeId.EntityData.SegmentPath
    localMplsTeLmpNodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localMplsTeLmpNodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localMplsTeLmpNodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localMplsTeLmpNodeId.EntityData.Children = types.NewOrderedMap()
    localMplsTeLmpNodeId.EntityData.Children.Append("address", types.YChild{"Address", &localMplsTeLmpNodeId.Address})
    localMplsTeLmpNodeId.EntityData.Leafs = types.NewOrderedMap()

    localMplsTeLmpNodeId.EntityData.YListKeys = []string {}

    return &(localMplsTeLmpNodeId.EntityData)
}

// Lmp_GlobalStatus_LocalMplsTeLmpNodeId_Address
// Address Union
type Lmp_GlobalStatus_LocalMplsTeLmpNodeId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GlobalStatus_LocalMplsTeLmpNodeId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-mpls-te-lmp-node-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/global-status/local-mpls-te-lmp-node-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GlobalStatus_LocalGmplsUniLmpNodeId
// GMPLS UNI LMP Node ID
type Lmp_GlobalStatus_LocalGmplsUniLmpNodeId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GlobalStatus_LocalGmplsUniLmpNodeId_Address
}

func (localGmplsUniLmpNodeId *Lmp_GlobalStatus_LocalGmplsUniLmpNodeId) GetEntityData() *types.CommonEntityData {
    localGmplsUniLmpNodeId.EntityData.YFilter = localGmplsUniLmpNodeId.YFilter
    localGmplsUniLmpNodeId.EntityData.YangName = "local-gmpls-uni-lmp-node-id"
    localGmplsUniLmpNodeId.EntityData.BundleName = "cisco_ios_xr"
    localGmplsUniLmpNodeId.EntityData.ParentYangName = "global-status"
    localGmplsUniLmpNodeId.EntityData.SegmentPath = "local-gmpls-uni-lmp-node-id"
    localGmplsUniLmpNodeId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/global-status/" + localGmplsUniLmpNodeId.EntityData.SegmentPath
    localGmplsUniLmpNodeId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localGmplsUniLmpNodeId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localGmplsUniLmpNodeId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localGmplsUniLmpNodeId.EntityData.Children = types.NewOrderedMap()
    localGmplsUniLmpNodeId.EntityData.Children.Append("address", types.YChild{"Address", &localGmplsUniLmpNodeId.Address})
    localGmplsUniLmpNodeId.EntityData.Leafs = types.NewOrderedMap()

    localGmplsUniLmpNodeId.EntityData.YListKeys = []string {}

    return &(localGmplsUniLmpNodeId.EntityData)
}

// Lmp_GlobalStatus_LocalGmplsUniLmpNodeId_Address
// Address Union
type Lmp_GlobalStatus_LocalGmplsUniLmpNodeId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GlobalStatus_LocalGmplsUniLmpNodeId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-gmpls-uni-lmp-node-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/global-status/local-gmpls-uni-lmp-node-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_Clients
// UCP OLM clients container class
type Lmp_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on a particular OLM API client. The type is slice of
    // Lmp_Clients_Client.
    Client []*Lmp_Clients_Client
}

func (clients *Lmp_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "lmp"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/" + clients.EntityData.SegmentPath
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = types.NewOrderedMap()
    clients.EntityData.Children.Append("client", types.YChild{"Client", nil})
    for i := range clients.Client {
        clients.EntityData.Children.Append(types.GetSegmentPath(clients.Client[i]), types.YChild{"Client", clients.Client[i]})
    }
    clients.EntityData.Leafs = types.NewOrderedMap()

    clients.EntityData.YListKeys = []string {}

    return &(clients.EntityData)
}

// Lmp_Clients_Client
// Information on a particular OLM API client
type Lmp_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Client name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClientName interface{}

    // The RP name that the clientprocess is running on. The type is string.
    NodeName interface{}

    // The time the clientconnected in sec.
    ConnectedTime Lmp_Clients_Client_ConnectedTime
}

func (client *Lmp_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientName, "client-name")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("connected-time", types.YChild{"ConnectedTime", &client.ConnectedTime})
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", client.ClientName})
    client.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", client.NodeName})

    client.EntityData.YListKeys = []string {"ClientName"}

    return &(client.EntityData)
}

// Lmp_Clients_Client_ConnectedTime
// The time the clientconnected in sec
type Lmp_Clients_Client_ConnectedTime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The time the clientconnected in sec. The type is interface{} with range:
    // 0..4294967295.
    TimeConnected interface{}
}

func (connectedTime *Lmp_Clients_Client_ConnectedTime) GetEntityData() *types.CommonEntityData {
    connectedTime.EntityData.YFilter = connectedTime.YFilter
    connectedTime.EntityData.YangName = "connected-time"
    connectedTime.EntityData.BundleName = "cisco_ios_xr"
    connectedTime.EntityData.ParentYangName = "client"
    connectedTime.EntityData.SegmentPath = "connected-time"
    connectedTime.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/clients/client/" + connectedTime.EntityData.SegmentPath
    connectedTime.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectedTime.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectedTime.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectedTime.EntityData.Children = types.NewOrderedMap()
    connectedTime.EntityData.Leafs = types.NewOrderedMap()
    connectedTime.EntityData.Leafs.Append("time-connected", types.YLeaf{"TimeConnected", connectedTime.TimeConnected})

    connectedTime.EntityData.YListKeys = []string {}

    return &(connectedTime.EntityData)
}

// Lmp_GmplsUni
// GMPLS UNI specific OLM/LMP configuration
type Lmp_GmplsUni struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // UCP OLM TE Links container class.
    TeLinks Lmp_GmplsUni_TeLinks

    // UCP OLM neighbors container class.
    Neighbors Lmp_GmplsUni_Neighbors
}

func (gmplsUni *Lmp_GmplsUni) GetEntityData() *types.CommonEntityData {
    gmplsUni.EntityData.YFilter = gmplsUni.YFilter
    gmplsUni.EntityData.YangName = "gmpls-uni"
    gmplsUni.EntityData.BundleName = "cisco_ios_xr"
    gmplsUni.EntityData.ParentYangName = "lmp"
    gmplsUni.EntityData.SegmentPath = "gmpls-uni"
    gmplsUni.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/" + gmplsUni.EntityData.SegmentPath
    gmplsUni.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gmplsUni.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gmplsUni.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gmplsUni.EntityData.Children = types.NewOrderedMap()
    gmplsUni.EntityData.Children.Append("te-links", types.YChild{"TeLinks", &gmplsUni.TeLinks})
    gmplsUni.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &gmplsUni.Neighbors})
    gmplsUni.EntityData.Leafs = types.NewOrderedMap()

    gmplsUni.EntityData.YListKeys = []string {}

    return &(gmplsUni.EntityData)
}

// Lmp_GmplsUni_TeLinks
// UCP OLM TE Links container class
type Lmp_GmplsUni_TeLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on a particular OLM TE Link. The type is slice of
    // Lmp_GmplsUni_TeLinks_TeLink.
    TeLink []*Lmp_GmplsUni_TeLinks_TeLink
}

func (teLinks *Lmp_GmplsUni_TeLinks) GetEntityData() *types.CommonEntityData {
    teLinks.EntityData.YFilter = teLinks.YFilter
    teLinks.EntityData.YangName = "te-links"
    teLinks.EntityData.BundleName = "cisco_ios_xr"
    teLinks.EntityData.ParentYangName = "gmpls-uni"
    teLinks.EntityData.SegmentPath = "te-links"
    teLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/" + teLinks.EntityData.SegmentPath
    teLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teLinks.EntityData.Children = types.NewOrderedMap()
    teLinks.EntityData.Children.Append("te-link", types.YChild{"TeLink", nil})
    for i := range teLinks.TeLink {
        teLinks.EntityData.Children.Append(types.GetSegmentPath(teLinks.TeLink[i]), types.YChild{"TeLink", teLinks.TeLink[i]})
    }
    teLinks.EntityData.Leafs = types.NewOrderedMap()

    teLinks.EntityData.YListKeys = []string {}

    return &(teLinks.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink
// Information on a particular OLM TE Link
type Lmp_GmplsUni_TeLinks_TeLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    ControllerName interface{}

    // Interface forOLM info. The type is string.
    InterfaceName interface{}

    // Protocol owningthis te-link. The type is OlmObjectOwner.
    ProtocolOwner interface{}

    // The name of the neighbor. The type is string.
    NeighborName interface{}

    // The IPCC ID. The type is interface{} with range: 0..4294967295.
    IpccId interface{}

    // OLM IPCC type. The type is Olmipcc.
    IpcCtype interface{}

    // The name ofthe IPCC associated with the TE Link. The type is string.
    IpccName interface{}

    // The local mux capability. The type is OlmMuxCap.
    LocalMuxCap interface{}

    // The remote mux capability. The type is OlmMuxCap.
    RemoteMuxCap interface{}

    // data link IM state. The type is OlmCompLinkImState.
    ImState interface{}

    // data link LMP state. The type is OlmCompLinkLmpState.
    LmpState interface{}

    // TE LinkLMP state. The type is OlmteLinkLmpState.
    TeLinkLmpState interface{}

    // GMPLS localminimum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkLocalMinimumBandwidth interface{}

    // GMPLS localmaximum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkLocalMaximumBandwidth interface{}

    // GMPLSneighbor minimum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkNeighborMinimumBandwidth interface{}

    // GMPLSneighbor maximum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkNeighborMaximumBandwidth interface{}

    // GMPLS locallink encoding type. The type is OlmLinkEncoding.
    GmplsTeLinkLocalEncodingType interface{}

    // GMPLS neighborlink encoding type. The type is OlmLinkEncoding.
    GmplsTeLinkNeighborEncodingType interface{}

    // Is LMP enabledon this TE link. The type is bool.
    IsLmpEnabled interface{}

    // LMP transmitmessage ID. The type is interface{} with range: 0..4294967295.
    LmpTransmitMsgId interface{}

    // LMP receivemessage ID. The type is interface{} with range: 0..4294967295.
    LmpReceiveMsgId interface{}

    // Component link LMP status indicators. The type is slice of
    // OlmCompLinkLmpStatus.
    LmpCompLinkStatus []interface{}

    // The local datalink ID.
    LocalLinkId Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId

    // The remote datalink ID.
    RemoteLinkId Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId

    // Local TE-Link ID/ TNA address.
    LocalTeLinkId Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId

    // Remote TE-Link ID/ TNA address.
    RemoteTeLinkId Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId

    // The address of the neighbor.
    NeighborAddress Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress

    // The remote node's IPCC address.
    RemoteIpccAddress Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress
}

func (teLink *Lmp_GmplsUni_TeLinks_TeLink) GetEntityData() *types.CommonEntityData {
    teLink.EntityData.YFilter = teLink.YFilter
    teLink.EntityData.YangName = "te-link"
    teLink.EntityData.BundleName = "cisco_ios_xr"
    teLink.EntityData.ParentYangName = "te-links"
    teLink.EntityData.SegmentPath = "te-link" + types.AddKeyToken(teLink.ControllerName, "controller-name")
    teLink.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/" + teLink.EntityData.SegmentPath
    teLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teLink.EntityData.Children = types.NewOrderedMap()
    teLink.EntityData.Children.Append("local-link-id", types.YChild{"LocalLinkId", &teLink.LocalLinkId})
    teLink.EntityData.Children.Append("remote-link-id", types.YChild{"RemoteLinkId", &teLink.RemoteLinkId})
    teLink.EntityData.Children.Append("local-te-link-id", types.YChild{"LocalTeLinkId", &teLink.LocalTeLinkId})
    teLink.EntityData.Children.Append("remote-te-link-id", types.YChild{"RemoteTeLinkId", &teLink.RemoteTeLinkId})
    teLink.EntityData.Children.Append("neighbor-address", types.YChild{"NeighborAddress", &teLink.NeighborAddress})
    teLink.EntityData.Children.Append("remote-ipcc-address", types.YChild{"RemoteIpccAddress", &teLink.RemoteIpccAddress})
    teLink.EntityData.Leafs = types.NewOrderedMap()
    teLink.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", teLink.ControllerName})
    teLink.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", teLink.InterfaceName})
    teLink.EntityData.Leafs.Append("protocol-owner", types.YLeaf{"ProtocolOwner", teLink.ProtocolOwner})
    teLink.EntityData.Leafs.Append("neighbor-name", types.YLeaf{"NeighborName", teLink.NeighborName})
    teLink.EntityData.Leafs.Append("ipcc-id", types.YLeaf{"IpccId", teLink.IpccId})
    teLink.EntityData.Leafs.Append("ipc-ctype", types.YLeaf{"IpcCtype", teLink.IpcCtype})
    teLink.EntityData.Leafs.Append("ipcc-name", types.YLeaf{"IpccName", teLink.IpccName})
    teLink.EntityData.Leafs.Append("local-mux-cap", types.YLeaf{"LocalMuxCap", teLink.LocalMuxCap})
    teLink.EntityData.Leafs.Append("remote-mux-cap", types.YLeaf{"RemoteMuxCap", teLink.RemoteMuxCap})
    teLink.EntityData.Leafs.Append("im-state", types.YLeaf{"ImState", teLink.ImState})
    teLink.EntityData.Leafs.Append("lmp-state", types.YLeaf{"LmpState", teLink.LmpState})
    teLink.EntityData.Leafs.Append("te-link-lmp-state", types.YLeaf{"TeLinkLmpState", teLink.TeLinkLmpState})
    teLink.EntityData.Leafs.Append("gmpls-te-link-local-minimum-bandwidth", types.YLeaf{"GmplsTeLinkLocalMinimumBandwidth", teLink.GmplsTeLinkLocalMinimumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-local-maximum-bandwidth", types.YLeaf{"GmplsTeLinkLocalMaximumBandwidth", teLink.GmplsTeLinkLocalMaximumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-neighbor-minimum-bandwidth", types.YLeaf{"GmplsTeLinkNeighborMinimumBandwidth", teLink.GmplsTeLinkNeighborMinimumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-neighbor-maximum-bandwidth", types.YLeaf{"GmplsTeLinkNeighborMaximumBandwidth", teLink.GmplsTeLinkNeighborMaximumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-local-encoding-type", types.YLeaf{"GmplsTeLinkLocalEncodingType", teLink.GmplsTeLinkLocalEncodingType})
    teLink.EntityData.Leafs.Append("gmpls-te-link-neighbor-encoding-type", types.YLeaf{"GmplsTeLinkNeighborEncodingType", teLink.GmplsTeLinkNeighborEncodingType})
    teLink.EntityData.Leafs.Append("is-lmp-enabled", types.YLeaf{"IsLmpEnabled", teLink.IsLmpEnabled})
    teLink.EntityData.Leafs.Append("lmp-transmit-msg-id", types.YLeaf{"LmpTransmitMsgId", teLink.LmpTransmitMsgId})
    teLink.EntityData.Leafs.Append("lmp-receive-msg-id", types.YLeaf{"LmpReceiveMsgId", teLink.LmpReceiveMsgId})
    teLink.EntityData.Leafs.Append("lmp-comp-link-status", types.YLeaf{"LmpCompLinkStatus", teLink.LmpCompLinkStatus})

    teLink.EntityData.YListKeys = []string {"ControllerName"}

    return &(teLink.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId
// The local datalink ID
type Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId_Address
}

func (localLinkId *Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId) GetEntityData() *types.CommonEntityData {
    localLinkId.EntityData.YFilter = localLinkId.YFilter
    localLinkId.EntityData.YangName = "local-link-id"
    localLinkId.EntityData.BundleName = "cisco_ios_xr"
    localLinkId.EntityData.ParentYangName = "te-link"
    localLinkId.EntityData.SegmentPath = "local-link-id"
    localLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/" + localLinkId.EntityData.SegmentPath
    localLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLinkId.EntityData.Children = types.NewOrderedMap()
    localLinkId.EntityData.Children.Append("address", types.YChild{"Address", &localLinkId.Address})
    localLinkId.EntityData.Leafs = types.NewOrderedMap()

    localLinkId.EntityData.YListKeys = []string {}

    return &(localLinkId.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId_Address
// Address Union
type Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_TeLinks_TeLink_LocalLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/local-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId
// The remote datalink ID
type Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId_Address
}

func (remoteLinkId *Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId) GetEntityData() *types.CommonEntityData {
    remoteLinkId.EntityData.YFilter = remoteLinkId.YFilter
    remoteLinkId.EntityData.YangName = "remote-link-id"
    remoteLinkId.EntityData.BundleName = "cisco_ios_xr"
    remoteLinkId.EntityData.ParentYangName = "te-link"
    remoteLinkId.EntityData.SegmentPath = "remote-link-id"
    remoteLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/" + remoteLinkId.EntityData.SegmentPath
    remoteLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLinkId.EntityData.Children = types.NewOrderedMap()
    remoteLinkId.EntityData.Children.Append("address", types.YChild{"Address", &remoteLinkId.Address})
    remoteLinkId.EntityData.Leafs = types.NewOrderedMap()

    remoteLinkId.EntityData.YListKeys = []string {}

    return &(remoteLinkId.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId_Address
// Address Union
type Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_TeLinks_TeLink_RemoteLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/remote-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId
// Local TE-Link ID/ TNA address
type Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId_Address
}

func (localTeLinkId *Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId) GetEntityData() *types.CommonEntityData {
    localTeLinkId.EntityData.YFilter = localTeLinkId.YFilter
    localTeLinkId.EntityData.YangName = "local-te-link-id"
    localTeLinkId.EntityData.BundleName = "cisco_ios_xr"
    localTeLinkId.EntityData.ParentYangName = "te-link"
    localTeLinkId.EntityData.SegmentPath = "local-te-link-id"
    localTeLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/" + localTeLinkId.EntityData.SegmentPath
    localTeLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTeLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTeLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTeLinkId.EntityData.Children = types.NewOrderedMap()
    localTeLinkId.EntityData.Children.Append("address", types.YChild{"Address", &localTeLinkId.Address})
    localTeLinkId.EntityData.Leafs = types.NewOrderedMap()

    localTeLinkId.EntityData.YListKeys = []string {}

    return &(localTeLinkId.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId_Address
// Address Union
type Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_TeLinks_TeLink_LocalTeLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-te-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/local-te-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId
// Remote TE-Link ID/ TNA address
type Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId_Address
}

func (remoteTeLinkId *Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId) GetEntityData() *types.CommonEntityData {
    remoteTeLinkId.EntityData.YFilter = remoteTeLinkId.YFilter
    remoteTeLinkId.EntityData.YangName = "remote-te-link-id"
    remoteTeLinkId.EntityData.BundleName = "cisco_ios_xr"
    remoteTeLinkId.EntityData.ParentYangName = "te-link"
    remoteTeLinkId.EntityData.SegmentPath = "remote-te-link-id"
    remoteTeLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/" + remoteTeLinkId.EntityData.SegmentPath
    remoteTeLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteTeLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteTeLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteTeLinkId.EntityData.Children = types.NewOrderedMap()
    remoteTeLinkId.EntityData.Children.Append("address", types.YChild{"Address", &remoteTeLinkId.Address})
    remoteTeLinkId.EntityData.Leafs = types.NewOrderedMap()

    remoteTeLinkId.EntityData.YListKeys = []string {}

    return &(remoteTeLinkId.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId_Address
// Address Union
type Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_TeLinks_TeLink_RemoteTeLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-te-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/remote-te-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress
// The address of the neighbor
type Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress_Address
}

func (neighborAddress *Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress) GetEntityData() *types.CommonEntityData {
    neighborAddress.EntityData.YFilter = neighborAddress.YFilter
    neighborAddress.EntityData.YangName = "neighbor-address"
    neighborAddress.EntityData.BundleName = "cisco_ios_xr"
    neighborAddress.EntityData.ParentYangName = "te-link"
    neighborAddress.EntityData.SegmentPath = "neighbor-address"
    neighborAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/" + neighborAddress.EntityData.SegmentPath
    neighborAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborAddress.EntityData.Children = types.NewOrderedMap()
    neighborAddress.EntityData.Children.Append("address", types.YChild{"Address", &neighborAddress.Address})
    neighborAddress.EntityData.Leafs = types.NewOrderedMap()

    neighborAddress.EntityData.YListKeys = []string {}

    return &(neighborAddress.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress_Address
// Address Union
type Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_TeLinks_TeLink_NeighborAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "neighbor-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/neighbor-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress
// The remote node's IPCC address
type Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress_Address
}

func (remoteIpccAddress *Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress) GetEntityData() *types.CommonEntityData {
    remoteIpccAddress.EntityData.YFilter = remoteIpccAddress.YFilter
    remoteIpccAddress.EntityData.YangName = "remote-ipcc-address"
    remoteIpccAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteIpccAddress.EntityData.ParentYangName = "te-link"
    remoteIpccAddress.EntityData.SegmentPath = "remote-ipcc-address"
    remoteIpccAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/" + remoteIpccAddress.EntityData.SegmentPath
    remoteIpccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIpccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIpccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIpccAddress.EntityData.Children = types.NewOrderedMap()
    remoteIpccAddress.EntityData.Children.Append("address", types.YChild{"Address", &remoteIpccAddress.Address})
    remoteIpccAddress.EntityData.Leafs = types.NewOrderedMap()

    remoteIpccAddress.EntityData.YListKeys = []string {}

    return &(remoteIpccAddress.EntityData)
}

// Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress_Address
// Address Union
type Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_TeLinks_TeLink_RemoteIpccAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-ipcc-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/te-links/te-link/remote-ipcc-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors
// UCP OLM neighbors container class
type Lmp_GmplsUni_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information on a particular OLM neighbor. The type is slice of
    // Lmp_GmplsUni_Neighbors_Neighbor.
    Neighbor []*Lmp_GmplsUni_Neighbors_Neighbor
}

func (neighbors *Lmp_GmplsUni_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "gmpls-uni"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor
// Information on a particular OLM neighbor
type Lmp_GmplsUni_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Neighbor name. The type is string with length:
    // 1..64.
    NeighborName interface{}

    // Protocol owningthis neighbor. The type is OlmObjectOwner.
    ProtocolOwner interface{}

    // The global active IPCCfor this neighbor. The type is interface{} with
    // range: 0..4294967295.
    IpccId interface{}

    // Is LMP enabled on this neighbor [DEPRECATED]. The type is bool.
    IsLmpEnabled interface{}

    // Are LMP hellos disabled through configuration for this neighbor
    // [DEPRECATED]. The type is bool.
    IsLmpConfigDisabled interface{}

    // LMP transmit message ID [DEPRECATED]. The type is interface{} with range:
    // 0..4294967295.
    LmpTransmitMsgId interface{}

    // LMP receive message ID [DEPRECATED]. The type is interface{} with range:
    // 0..4294967295.
    LmpReceiveMsgId interface{}

    // LMP link summary transmit packet count [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpLinkSumTransmitPackets interface{}

    // LMP link summary receive packet count[DEPRECATED]. The type is interface{}
    // with range: 0..4294967295.
    LmpLinkSumReceivePackets interface{}

    // The remote node ID of the neighbor.
    NeighborAddress Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress

    // A list of TE Links connected to this neighbor. The type is slice of
    // Lmp_GmplsUni_Neighbors_Neighbor_TeLink.
    TeLink []*Lmp_GmplsUni_Neighbors_Neighbor_TeLink

    // A list of IPCCs connected to this neighbor. The type is slice of
    // Lmp_GmplsUni_Neighbors_Neighbor_Ipcc.
    Ipcc []*Lmp_GmplsUni_Neighbors_Neighbor_Ipcc
}

func (neighbor *Lmp_GmplsUni_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborName, "neighbor-name")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Children.Append("neighbor-address", types.YChild{"NeighborAddress", &neighbor.NeighborAddress})
    neighbor.EntityData.Children.Append("te-link", types.YChild{"TeLink", nil})
    for i := range neighbor.TeLink {
        types.SetYListKey(neighbor.TeLink[i], i)
        neighbor.EntityData.Children.Append(types.GetSegmentPath(neighbor.TeLink[i]), types.YChild{"TeLink", neighbor.TeLink[i]})
    }
    neighbor.EntityData.Children.Append("ipcc", types.YChild{"Ipcc", nil})
    for i := range neighbor.Ipcc {
        types.SetYListKey(neighbor.Ipcc[i], i)
        neighbor.EntityData.Children.Append(types.GetSegmentPath(neighbor.Ipcc[i]), types.YChild{"Ipcc", neighbor.Ipcc[i]})
    }
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-name", types.YLeaf{"NeighborName", neighbor.NeighborName})
    neighbor.EntityData.Leafs.Append("protocol-owner", types.YLeaf{"ProtocolOwner", neighbor.ProtocolOwner})
    neighbor.EntityData.Leafs.Append("ipcc-id", types.YLeaf{"IpccId", neighbor.IpccId})
    neighbor.EntityData.Leafs.Append("is-lmp-enabled", types.YLeaf{"IsLmpEnabled", neighbor.IsLmpEnabled})
    neighbor.EntityData.Leafs.Append("is-lmp-config-disabled", types.YLeaf{"IsLmpConfigDisabled", neighbor.IsLmpConfigDisabled})
    neighbor.EntityData.Leafs.Append("lmp-transmit-msg-id", types.YLeaf{"LmpTransmitMsgId", neighbor.LmpTransmitMsgId})
    neighbor.EntityData.Leafs.Append("lmp-receive-msg-id", types.YLeaf{"LmpReceiveMsgId", neighbor.LmpReceiveMsgId})
    neighbor.EntityData.Leafs.Append("lmp-link-sum-transmit-packets", types.YLeaf{"LmpLinkSumTransmitPackets", neighbor.LmpLinkSumTransmitPackets})
    neighbor.EntityData.Leafs.Append("lmp-link-sum-receive-packets", types.YLeaf{"LmpLinkSumReceivePackets", neighbor.LmpLinkSumReceivePackets})

    neighbor.EntityData.YListKeys = []string {"NeighborName"}

    return &(neighbor.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress
// The remote node ID of the neighbor
type Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress_Address
}

func (neighborAddress *Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress) GetEntityData() *types.CommonEntityData {
    neighborAddress.EntityData.YFilter = neighborAddress.YFilter
    neighborAddress.EntityData.YangName = "neighbor-address"
    neighborAddress.EntityData.BundleName = "cisco_ios_xr"
    neighborAddress.EntityData.ParentYangName = "neighbor"
    neighborAddress.EntityData.SegmentPath = "neighbor-address"
    neighborAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/" + neighborAddress.EntityData.SegmentPath
    neighborAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborAddress.EntityData.Children = types.NewOrderedMap()
    neighborAddress.EntityData.Children.Append("address", types.YChild{"Address", &neighborAddress.Address})
    neighborAddress.EntityData.Leafs = types.NewOrderedMap()

    neighborAddress.EntityData.YListKeys = []string {}

    return &(neighborAddress.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_NeighborAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "neighbor-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/neighbor-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink
// A list of TE Links connected to this neighbor
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface forOLM info. The type is string.
    InterfaceName interface{}

    // Protocol owningthis te-link. The type is OlmObjectOwner.
    ProtocolOwner interface{}

    // The name of the neighbor. The type is string.
    NeighborName interface{}

    // The IPCC ID. The type is interface{} with range: 0..4294967295.
    IpccId interface{}

    // OLM IPCC type. The type is Olmipcc.
    IpcCtype interface{}

    // The name ofthe IPCC associated with the TE Link. The type is string.
    IpccName interface{}

    // The local mux capability. The type is OlmMuxCap.
    LocalMuxCap interface{}

    // The remote mux capability. The type is OlmMuxCap.
    RemoteMuxCap interface{}

    // data link IM state. The type is OlmCompLinkImState.
    ImState interface{}

    // data link LMP state. The type is OlmCompLinkLmpState.
    LmpState interface{}

    // TE LinkLMP state. The type is OlmteLinkLmpState.
    TeLinkLmpState interface{}

    // GMPLS localminimum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkLocalMinimumBandwidth interface{}

    // GMPLS localmaximum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkLocalMaximumBandwidth interface{}

    // GMPLSneighbor minimum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkNeighborMinimumBandwidth interface{}

    // GMPLSneighbor maximum B/W in bytes/sec. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte/s.
    GmplsTeLinkNeighborMaximumBandwidth interface{}

    // GMPLS locallink encoding type. The type is OlmLinkEncoding.
    GmplsTeLinkLocalEncodingType interface{}

    // GMPLS neighborlink encoding type. The type is OlmLinkEncoding.
    GmplsTeLinkNeighborEncodingType interface{}

    // Is LMP enabledon this TE link. The type is bool.
    IsLmpEnabled interface{}

    // LMP transmitmessage ID. The type is interface{} with range: 0..4294967295.
    LmpTransmitMsgId interface{}

    // LMP receivemessage ID. The type is interface{} with range: 0..4294967295.
    LmpReceiveMsgId interface{}

    // Component link LMP status indicators. The type is slice of
    // OlmCompLinkLmpStatus.
    LmpCompLinkStatus []interface{}

    // The local datalink ID.
    LocalLinkId Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId

    // The remote datalink ID.
    RemoteLinkId Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId

    // Local TE-Link ID/ TNA address.
    LocalTeLinkId Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId

    // Remote TE-Link ID/ TNA address.
    RemoteTeLinkId Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId

    // The address of the neighbor.
    NeighborAddress Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress

    // The remote node's IPCC address.
    RemoteIpccAddress Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress
}

func (teLink *Lmp_GmplsUni_Neighbors_Neighbor_TeLink) GetEntityData() *types.CommonEntityData {
    teLink.EntityData.YFilter = teLink.YFilter
    teLink.EntityData.YangName = "te-link"
    teLink.EntityData.BundleName = "cisco_ios_xr"
    teLink.EntityData.ParentYangName = "neighbor"
    teLink.EntityData.SegmentPath = "te-link" + types.AddNoKeyToken(teLink)
    teLink.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/" + teLink.EntityData.SegmentPath
    teLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    teLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    teLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    teLink.EntityData.Children = types.NewOrderedMap()
    teLink.EntityData.Children.Append("local-link-id", types.YChild{"LocalLinkId", &teLink.LocalLinkId})
    teLink.EntityData.Children.Append("remote-link-id", types.YChild{"RemoteLinkId", &teLink.RemoteLinkId})
    teLink.EntityData.Children.Append("local-te-link-id", types.YChild{"LocalTeLinkId", &teLink.LocalTeLinkId})
    teLink.EntityData.Children.Append("remote-te-link-id", types.YChild{"RemoteTeLinkId", &teLink.RemoteTeLinkId})
    teLink.EntityData.Children.Append("neighbor-address", types.YChild{"NeighborAddress", &teLink.NeighborAddress})
    teLink.EntityData.Children.Append("remote-ipcc-address", types.YChild{"RemoteIpccAddress", &teLink.RemoteIpccAddress})
    teLink.EntityData.Leafs = types.NewOrderedMap()
    teLink.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", teLink.InterfaceName})
    teLink.EntityData.Leafs.Append("protocol-owner", types.YLeaf{"ProtocolOwner", teLink.ProtocolOwner})
    teLink.EntityData.Leafs.Append("neighbor-name", types.YLeaf{"NeighborName", teLink.NeighborName})
    teLink.EntityData.Leafs.Append("ipcc-id", types.YLeaf{"IpccId", teLink.IpccId})
    teLink.EntityData.Leafs.Append("ipc-ctype", types.YLeaf{"IpcCtype", teLink.IpcCtype})
    teLink.EntityData.Leafs.Append("ipcc-name", types.YLeaf{"IpccName", teLink.IpccName})
    teLink.EntityData.Leafs.Append("local-mux-cap", types.YLeaf{"LocalMuxCap", teLink.LocalMuxCap})
    teLink.EntityData.Leafs.Append("remote-mux-cap", types.YLeaf{"RemoteMuxCap", teLink.RemoteMuxCap})
    teLink.EntityData.Leafs.Append("im-state", types.YLeaf{"ImState", teLink.ImState})
    teLink.EntityData.Leafs.Append("lmp-state", types.YLeaf{"LmpState", teLink.LmpState})
    teLink.EntityData.Leafs.Append("te-link-lmp-state", types.YLeaf{"TeLinkLmpState", teLink.TeLinkLmpState})
    teLink.EntityData.Leafs.Append("gmpls-te-link-local-minimum-bandwidth", types.YLeaf{"GmplsTeLinkLocalMinimumBandwidth", teLink.GmplsTeLinkLocalMinimumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-local-maximum-bandwidth", types.YLeaf{"GmplsTeLinkLocalMaximumBandwidth", teLink.GmplsTeLinkLocalMaximumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-neighbor-minimum-bandwidth", types.YLeaf{"GmplsTeLinkNeighborMinimumBandwidth", teLink.GmplsTeLinkNeighborMinimumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-neighbor-maximum-bandwidth", types.YLeaf{"GmplsTeLinkNeighborMaximumBandwidth", teLink.GmplsTeLinkNeighborMaximumBandwidth})
    teLink.EntityData.Leafs.Append("gmpls-te-link-local-encoding-type", types.YLeaf{"GmplsTeLinkLocalEncodingType", teLink.GmplsTeLinkLocalEncodingType})
    teLink.EntityData.Leafs.Append("gmpls-te-link-neighbor-encoding-type", types.YLeaf{"GmplsTeLinkNeighborEncodingType", teLink.GmplsTeLinkNeighborEncodingType})
    teLink.EntityData.Leafs.Append("is-lmp-enabled", types.YLeaf{"IsLmpEnabled", teLink.IsLmpEnabled})
    teLink.EntityData.Leafs.Append("lmp-transmit-msg-id", types.YLeaf{"LmpTransmitMsgId", teLink.LmpTransmitMsgId})
    teLink.EntityData.Leafs.Append("lmp-receive-msg-id", types.YLeaf{"LmpReceiveMsgId", teLink.LmpReceiveMsgId})
    teLink.EntityData.Leafs.Append("lmp-comp-link-status", types.YLeaf{"LmpCompLinkStatus", teLink.LmpCompLinkStatus})

    teLink.EntityData.YListKeys = []string {}

    return &(teLink.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId
// The local datalink ID
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId_Address
}

func (localLinkId *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId) GetEntityData() *types.CommonEntityData {
    localLinkId.EntityData.YFilter = localLinkId.YFilter
    localLinkId.EntityData.YangName = "local-link-id"
    localLinkId.EntityData.BundleName = "cisco_ios_xr"
    localLinkId.EntityData.ParentYangName = "te-link"
    localLinkId.EntityData.SegmentPath = "local-link-id"
    localLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/" + localLinkId.EntityData.SegmentPath
    localLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localLinkId.EntityData.Children = types.NewOrderedMap()
    localLinkId.EntityData.Children.Append("address", types.YChild{"Address", &localLinkId.Address})
    localLinkId.EntityData.Leafs = types.NewOrderedMap()

    localLinkId.EntityData.YListKeys = []string {}

    return &(localLinkId.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/local-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId
// The remote datalink ID
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId_Address
}

func (remoteLinkId *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId) GetEntityData() *types.CommonEntityData {
    remoteLinkId.EntityData.YFilter = remoteLinkId.YFilter
    remoteLinkId.EntityData.YangName = "remote-link-id"
    remoteLinkId.EntityData.BundleName = "cisco_ios_xr"
    remoteLinkId.EntityData.ParentYangName = "te-link"
    remoteLinkId.EntityData.SegmentPath = "remote-link-id"
    remoteLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/" + remoteLinkId.EntityData.SegmentPath
    remoteLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteLinkId.EntityData.Children = types.NewOrderedMap()
    remoteLinkId.EntityData.Children.Append("address", types.YChild{"Address", &remoteLinkId.Address})
    remoteLinkId.EntityData.Leafs = types.NewOrderedMap()

    remoteLinkId.EntityData.YListKeys = []string {}

    return &(remoteLinkId.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/remote-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId
// Local TE-Link ID/ TNA address
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId_Address
}

func (localTeLinkId *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId) GetEntityData() *types.CommonEntityData {
    localTeLinkId.EntityData.YFilter = localTeLinkId.YFilter
    localTeLinkId.EntityData.YangName = "local-te-link-id"
    localTeLinkId.EntityData.BundleName = "cisco_ios_xr"
    localTeLinkId.EntityData.ParentYangName = "te-link"
    localTeLinkId.EntityData.SegmentPath = "local-te-link-id"
    localTeLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/" + localTeLinkId.EntityData.SegmentPath
    localTeLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localTeLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localTeLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localTeLinkId.EntityData.Children = types.NewOrderedMap()
    localTeLinkId.EntityData.Children.Append("address", types.YChild{"Address", &localTeLinkId.Address})
    localTeLinkId.EntityData.Leafs = types.NewOrderedMap()

    localTeLinkId.EntityData.YListKeys = []string {}

    return &(localTeLinkId.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_LocalTeLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "local-te-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/local-te-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId
// Remote TE-Link ID/ TNA address
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId_Address
}

func (remoteTeLinkId *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId) GetEntityData() *types.CommonEntityData {
    remoteTeLinkId.EntityData.YFilter = remoteTeLinkId.YFilter
    remoteTeLinkId.EntityData.YangName = "remote-te-link-id"
    remoteTeLinkId.EntityData.BundleName = "cisco_ios_xr"
    remoteTeLinkId.EntityData.ParentYangName = "te-link"
    remoteTeLinkId.EntityData.SegmentPath = "remote-te-link-id"
    remoteTeLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/" + remoteTeLinkId.EntityData.SegmentPath
    remoteTeLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteTeLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteTeLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteTeLinkId.EntityData.Children = types.NewOrderedMap()
    remoteTeLinkId.EntityData.Children.Append("address", types.YChild{"Address", &remoteTeLinkId.Address})
    remoteTeLinkId.EntityData.Leafs = types.NewOrderedMap()

    remoteTeLinkId.EntityData.YListKeys = []string {}

    return &(remoteTeLinkId.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteTeLinkId_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-te-link-id"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/remote-te-link-id/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress
// The address of the neighbor
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress_Address
}

func (neighborAddress *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress) GetEntityData() *types.CommonEntityData {
    neighborAddress.EntityData.YFilter = neighborAddress.YFilter
    neighborAddress.EntityData.YangName = "neighbor-address"
    neighborAddress.EntityData.BundleName = "cisco_ios_xr"
    neighborAddress.EntityData.ParentYangName = "te-link"
    neighborAddress.EntityData.SegmentPath = "neighbor-address"
    neighborAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/" + neighborAddress.EntityData.SegmentPath
    neighborAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighborAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighborAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighborAddress.EntityData.Children = types.NewOrderedMap()
    neighborAddress.EntityData.Children.Append("address", types.YChild{"Address", &neighborAddress.Address})
    neighborAddress.EntityData.Leafs = types.NewOrderedMap()

    neighborAddress.EntityData.YListKeys = []string {}

    return &(neighborAddress.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_NeighborAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "neighbor-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/neighbor-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress
// The remote node's IPCC address
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress_Address
}

func (remoteIpccAddress *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress) GetEntityData() *types.CommonEntityData {
    remoteIpccAddress.EntityData.YFilter = remoteIpccAddress.YFilter
    remoteIpccAddress.EntityData.YangName = "remote-ipcc-address"
    remoteIpccAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteIpccAddress.EntityData.ParentYangName = "te-link"
    remoteIpccAddress.EntityData.SegmentPath = "remote-ipcc-address"
    remoteIpccAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/" + remoteIpccAddress.EntityData.SegmentPath
    remoteIpccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIpccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIpccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIpccAddress.EntityData.Children = types.NewOrderedMap()
    remoteIpccAddress.EntityData.Children.Append("address", types.YChild{"Address", &remoteIpccAddress.Address})
    remoteIpccAddress.EntityData.Leafs = types.NewOrderedMap()

    remoteIpccAddress.EntityData.YListKeys = []string {}

    return &(remoteIpccAddress.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_TeLink_RemoteIpccAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-ipcc-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/te-link/remote-ipcc-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc
// A list of IPCCs connected to this neighbor
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Global active IPCCfor this neighbor. The type is interface{} with range:
    // 0..4294967295.
    IpccId interface{}

    // OLM IPCC type. The type is Olmipcc.
    IpcCtype interface{}

    // The interface name forI/F IPCCs. The type is string.
    InterfaceName interface{}

    // Neighbor name of theIPCCs neighbor. The type is string.
    NeighborName interface{}

    // OLM IPCC master state. The type is OlmipccState.
    IpccState interface{}

    // LMP hello send interval in msec [DEPRECATED]. The type is interface{} with
    // range: 0..4294967295.
    LmpHelloInterval interface{}

    // LMP minimum acceptable hello send interval [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpHelloIntervalMinimum interface{}

    // LMP maximum acceptable hello send interval [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpHelloIntervalMaximum interface{}

    // LMP hello dead interval [DEPRECATED]. The type is interface{} with range:
    // 0..4294967295.
    LmpHelloDeadInterval interface{}

    // LMP minimum acceptable hello dead interval [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpHelloDeadIntervalMinimum interface{}

    // LMP maximum acceptable hello dead interval [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpHelloDeadIntervalMaximum interface{}

    // LMP hello transmit packet count[DEPRECATED]. The type is interface{} with
    // range: 0..4294967295.
    LmpHelloTransmitPackets interface{}

    // LMP hello receive packet count [DEPRECATED]. The type is interface{} with
    // range: 0..4294967295.
    LmpHelloReceivePackets interface{}

    // LMP hello transmit packet sequence number [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpHelloTransmitPacketSequenceNumber interface{}

    // LMP hello receive packet sequence number[DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpHelloReceivePacketSequenceNumber interface{}

    // LMP transmit message ID[DEPRECATED]. The type is interface{} with range:
    // 0..4294967295.
    LmpTransmitMsgId interface{}

    // LMP receive message ID[DEPRECATED]. The type is interface{} with range:
    // 0..4294967295.
    LmpReceiveMsgId interface{}

    // LMP link summary transmit packet count [DEPRECATED]. The type is
    // interface{} with range: 0..4294967295.
    LmpLinkSumTransmitPackets interface{}

    // LMP link summary receive packet count [DEPRECATED]. The type is interface{}
    // with range: 0..4294967295.
    LmpLinkSumReceivePackets interface{}

    // The remote node'sIPCC address.
    RemoteIpccAddress Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress

    // The local IPCC address.
    SourceIpCcAddress Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress
}

func (ipcc *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc) GetEntityData() *types.CommonEntityData {
    ipcc.EntityData.YFilter = ipcc.YFilter
    ipcc.EntityData.YangName = "ipcc"
    ipcc.EntityData.BundleName = "cisco_ios_xr"
    ipcc.EntityData.ParentYangName = "neighbor"
    ipcc.EntityData.SegmentPath = "ipcc" + types.AddNoKeyToken(ipcc)
    ipcc.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/" + ipcc.EntityData.SegmentPath
    ipcc.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipcc.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipcc.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipcc.EntityData.Children = types.NewOrderedMap()
    ipcc.EntityData.Children.Append("remote-ipcc-address", types.YChild{"RemoteIpccAddress", &ipcc.RemoteIpccAddress})
    ipcc.EntityData.Children.Append("source-ip-cc-address", types.YChild{"SourceIpCcAddress", &ipcc.SourceIpCcAddress})
    ipcc.EntityData.Leafs = types.NewOrderedMap()
    ipcc.EntityData.Leafs.Append("ipcc-id", types.YLeaf{"IpccId", ipcc.IpccId})
    ipcc.EntityData.Leafs.Append("ipc-ctype", types.YLeaf{"IpcCtype", ipcc.IpcCtype})
    ipcc.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", ipcc.InterfaceName})
    ipcc.EntityData.Leafs.Append("neighbor-name", types.YLeaf{"NeighborName", ipcc.NeighborName})
    ipcc.EntityData.Leafs.Append("ipcc-state", types.YLeaf{"IpccState", ipcc.IpccState})
    ipcc.EntityData.Leafs.Append("lmp-hello-interval", types.YLeaf{"LmpHelloInterval", ipcc.LmpHelloInterval})
    ipcc.EntityData.Leafs.Append("lmp-hello-interval-minimum", types.YLeaf{"LmpHelloIntervalMinimum", ipcc.LmpHelloIntervalMinimum})
    ipcc.EntityData.Leafs.Append("lmp-hello-interval-maximum", types.YLeaf{"LmpHelloIntervalMaximum", ipcc.LmpHelloIntervalMaximum})
    ipcc.EntityData.Leafs.Append("lmp-hello-dead-interval", types.YLeaf{"LmpHelloDeadInterval", ipcc.LmpHelloDeadInterval})
    ipcc.EntityData.Leafs.Append("lmp-hello-dead-interval-minimum", types.YLeaf{"LmpHelloDeadIntervalMinimum", ipcc.LmpHelloDeadIntervalMinimum})
    ipcc.EntityData.Leafs.Append("lmp-hello-dead-interval-maximum", types.YLeaf{"LmpHelloDeadIntervalMaximum", ipcc.LmpHelloDeadIntervalMaximum})
    ipcc.EntityData.Leafs.Append("lmp-hello-transmit-packets", types.YLeaf{"LmpHelloTransmitPackets", ipcc.LmpHelloTransmitPackets})
    ipcc.EntityData.Leafs.Append("lmp-hello-receive-packets", types.YLeaf{"LmpHelloReceivePackets", ipcc.LmpHelloReceivePackets})
    ipcc.EntityData.Leafs.Append("lmp-hello-transmit-packet-sequence-number", types.YLeaf{"LmpHelloTransmitPacketSequenceNumber", ipcc.LmpHelloTransmitPacketSequenceNumber})
    ipcc.EntityData.Leafs.Append("lmp-hello-receive-packet-sequence-number", types.YLeaf{"LmpHelloReceivePacketSequenceNumber", ipcc.LmpHelloReceivePacketSequenceNumber})
    ipcc.EntityData.Leafs.Append("lmp-transmit-msg-id", types.YLeaf{"LmpTransmitMsgId", ipcc.LmpTransmitMsgId})
    ipcc.EntityData.Leafs.Append("lmp-receive-msg-id", types.YLeaf{"LmpReceiveMsgId", ipcc.LmpReceiveMsgId})
    ipcc.EntityData.Leafs.Append("lmp-link-sum-transmit-packets", types.YLeaf{"LmpLinkSumTransmitPackets", ipcc.LmpLinkSumTransmitPackets})
    ipcc.EntityData.Leafs.Append("lmp-link-sum-receive-packets", types.YLeaf{"LmpLinkSumReceivePackets", ipcc.LmpLinkSumReceivePackets})

    ipcc.EntityData.YListKeys = []string {}

    return &(ipcc.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress
// The remote node'sIPCC address
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress_Address
}

func (remoteIpccAddress *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress) GetEntityData() *types.CommonEntityData {
    remoteIpccAddress.EntityData.YFilter = remoteIpccAddress.YFilter
    remoteIpccAddress.EntityData.YangName = "remote-ipcc-address"
    remoteIpccAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteIpccAddress.EntityData.ParentYangName = "ipcc"
    remoteIpccAddress.EntityData.SegmentPath = "remote-ipcc-address"
    remoteIpccAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/ipcc/" + remoteIpccAddress.EntityData.SegmentPath
    remoteIpccAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteIpccAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteIpccAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteIpccAddress.EntityData.Children = types.NewOrderedMap()
    remoteIpccAddress.EntityData.Children.Append("address", types.YChild{"Address", &remoteIpccAddress.Address})
    remoteIpccAddress.EntityData.Leafs = types.NewOrderedMap()

    remoteIpccAddress.EntityData.YListKeys = []string {}

    return &(remoteIpccAddress.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_RemoteIpccAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "remote-ipcc-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/ipcc/remote-ipcc-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress
// The local IPCC address
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Union.
    Address Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress_Address
}

func (sourceIpCcAddress *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress) GetEntityData() *types.CommonEntityData {
    sourceIpCcAddress.EntityData.YFilter = sourceIpCcAddress.YFilter
    sourceIpCcAddress.EntityData.YangName = "source-ip-cc-address"
    sourceIpCcAddress.EntityData.BundleName = "cisco_ios_xr"
    sourceIpCcAddress.EntityData.ParentYangName = "ipcc"
    sourceIpCcAddress.EntityData.SegmentPath = "source-ip-cc-address"
    sourceIpCcAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/ipcc/" + sourceIpCcAddress.EntityData.SegmentPath
    sourceIpCcAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sourceIpCcAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sourceIpCcAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sourceIpCcAddress.EntityData.Children = types.NewOrderedMap()
    sourceIpCcAddress.EntityData.Children.Append("address", types.YChild{"Address", &sourceIpCcAddress.Address})
    sourceIpCcAddress.EntityData.Leafs = types.NewOrderedMap()

    sourceIpCcAddress.EntityData.YListKeys = []string {}

    return &(sourceIpCcAddress.EntityData)
}

// Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress_Address
// Address Union
type Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AddressType. The type is OlmAddrTypeId.
    AddressType interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}

    // Unnumberedaddress. The type is interface{} with range: 0..4294967295.
    UnnumberedAddress interface{}
}

func (address *Lmp_GmplsUni_Neighbors_Neighbor_Ipcc_SourceIpCcAddress_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "source-ip-cc-address"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/gmpls-uni/neighbors/neighbor/ipcc/source-ip-cc-address/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("address-type", types.YLeaf{"AddressType", address.AddressType})
    address.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", address.Ipv4Address})
    address.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", address.Ipv6Address})
    address.EntityData.Leafs.Append("unnumbered-address", types.YLeaf{"UnnumberedAddress", address.UnnumberedAddress})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Lmp_ComponentLinkIds
// UCP OLM component link ID container class
type Lmp_ComponentLinkIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Retrieve the LMP component link ID for a given controller. The type is
    // slice of Lmp_ComponentLinkIds_ComponentLinkId.
    ComponentLinkId []*Lmp_ComponentLinkIds_ComponentLinkId
}

func (componentLinkIds *Lmp_ComponentLinkIds) GetEntityData() *types.CommonEntityData {
    componentLinkIds.EntityData.YFilter = componentLinkIds.YFilter
    componentLinkIds.EntityData.YangName = "component-link-ids"
    componentLinkIds.EntityData.BundleName = "cisco_ios_xr"
    componentLinkIds.EntityData.ParentYangName = "lmp"
    componentLinkIds.EntityData.SegmentPath = "component-link-ids"
    componentLinkIds.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/" + componentLinkIds.EntityData.SegmentPath
    componentLinkIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    componentLinkIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    componentLinkIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    componentLinkIds.EntityData.Children = types.NewOrderedMap()
    componentLinkIds.EntityData.Children.Append("component-link-id", types.YChild{"ComponentLinkId", nil})
    for i := range componentLinkIds.ComponentLinkId {
        componentLinkIds.EntityData.Children.Append(types.GetSegmentPath(componentLinkIds.ComponentLinkId[i]), types.YChild{"ComponentLinkId", componentLinkIds.ComponentLinkId[i]})
    }
    componentLinkIds.EntityData.Leafs = types.NewOrderedMap()

    componentLinkIds.EntityData.YListKeys = []string {}

    return &(componentLinkIds.EntityData)
}

// Lmp_ComponentLinkIds_ComponentLinkId
// Retrieve the LMP component link ID for a given
// controller
type Lmp_ComponentLinkIds_ComponentLinkId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Controller name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    ControllerName interface{}

    // LMP component link ID for an I/F. The type is interface{} with range:
    // 0..4294967295.
    ComponentInterfaceId interface{}
}

func (componentLinkId *Lmp_ComponentLinkIds_ComponentLinkId) GetEntityData() *types.CommonEntityData {
    componentLinkId.EntityData.YFilter = componentLinkId.YFilter
    componentLinkId.EntityData.YangName = "component-link-id"
    componentLinkId.EntityData.BundleName = "cisco_ios_xr"
    componentLinkId.EntityData.ParentYangName = "component-link-ids"
    componentLinkId.EntityData.SegmentPath = "component-link-id" + types.AddKeyToken(componentLinkId.ControllerName, "controller-name")
    componentLinkId.EntityData.AbsolutePath = "Cisco-IOS-XR-lmp-oper:lmp/component-link-ids/" + componentLinkId.EntityData.SegmentPath
    componentLinkId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    componentLinkId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    componentLinkId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    componentLinkId.EntityData.Children = types.NewOrderedMap()
    componentLinkId.EntityData.Leafs = types.NewOrderedMap()
    componentLinkId.EntityData.Leafs.Append("controller-name", types.YLeaf{"ControllerName", componentLinkId.ControllerName})
    componentLinkId.EntityData.Leafs.Append("component-interface-id", types.YLeaf{"ComponentInterfaceId", componentLinkId.ComponentInterfaceId})

    componentLinkId.EntityData.YListKeys = []string {"ControllerName"}

    return &(componentLinkId.EntityData)
}

