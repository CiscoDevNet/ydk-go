// 
// The YANG-module for Protocol Independent Multicast (PIM). 
// The module defines configuration and operational data 
// for the following features:
// PIM Sparse Mode (PIM-SM)
// PIM Source-Specific Multicast (PIM-SSM)
// Bidirectional PIM (Bidir-PIM)
// Anycast-RP for PIM
// Bootstrap Router (BSR)  for PIM
// PIM Dense Mode (PIM-DM)
// Auto-RP - Cisco-propriatary
package pim

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pim"))
}

type AsmMappingMode struct {
}

func (id AsmMappingMode) String() string {
	return "pim:asm-mapping-mode"
}

type OtherMappingMode struct {
}

func (id OtherMappingMode) String() string {
	return "pim:other-mapping-mode"
}

type SsmMappingMode struct {
}

func (id SsmMappingMode) String() string {
	return "pim:ssm-mapping-mode"
}

type SmMappingMode struct {
}

func (id SmMappingMode) String() string {
	return "pim:sm-mapping-mode"
}

type PimBidirMappingMode struct {
}

func (id PimBidirMappingMode) String() string {
	return "pim:pim-bidir-mapping-mode"
}

type GroupToRpMappingMode struct {
}

func (id GroupToRpMappingMode) String() string {
	return "pim:group-to-rp-mapping-mode"
}

type DmMappingMode struct {
}

func (id DmMappingMode) String() string {
	return "pim:dm-mapping-mode"
}

// Origin represents This type  verify all uses of origin in model describes where a state was learned.
type Origin string

const (
    // The state's origin is none of the available sources or it is unknown.
    Origin_other_origin Origin = "other-origin"

    // PIM-request-states are learned by PIM joins (between PIM-routers).
    Origin_pim_request Origin = "pim-request"

    // SSM-Request-states are learned by SSM channel subscription, e.g., through 
    // IGMP3 (primarily only on last-hop-routers, although routers within the domain can keep 
    // this origin-type
    // ).
    Origin_ssm_request Origin = "ssm-request"

    // Fixed states are created automatically by the router at startup, to 
    // correspond to the well-defined prefixes of link-local and unroutable group addresses. 
    // These states are never destroyed.
    Origin_fixed Origin = "fixed"

    // Embedded states are created by the router to correspond to group 
    // prefixes that are to be treated as being in Embedded-RP format.
    Origin_embedded Origin = "embedded"

    // Static states are created and destroyed as a result of static configuration.
    Origin_static Origin = "static"

    // Config-SSM states are created and destroyed as a result of configuration 
    // of SSM address ranges to the local router.
    Origin_config_ssm Origin = "config-ssm"

    // Auto-RP states are created as a result of running Cisco's Auto-RP mechanism.
    Origin_auto_rp Origin = "auto-rp"

    // BSR states are created as a result of running the PIM Bootstrap Router 
    // (BSR) mechanism.
    Origin_bsr Origin = "bsr"

    // MSDP states are (*, G)-entries learned through a Multicast Source Discovery 
    // Protocol (MSDP) peer. This origin is applicable only for an RP running MSDP. 
    Origin_msdp Origin = "msdp"
)

// PimMode represents PIM mode active on an interface.
type PimMode string

const (
    // PIM sparse mode enabled on interface
    PimMode_sparse PimMode = "sparse"

    // PIM dense mode enable on interface.
    PimMode_dense PimMode = "dense"

    // PIM dense mode enable on interface.
    PimMode_sparse_dense PimMode = "sparse-dense"

    // PIM dense mode enable on interface.
    PimMode_dm_proxy PimMode = "dm-proxy"

    // PIM dense mode enable on interface.
    PimMode_none PimMode = "none"
)

// RouteProtocolType represents protocols need be supported.
type RouteProtocolType string

const (
    RouteProtocolType_other RouteProtocolType = "other"

    RouteProtocolType_local RouteProtocolType = "local"

    RouteProtocolType_netmgmt RouteProtocolType = "netmgmt"

    RouteProtocolType_icmp RouteProtocolType = "icmp"

    RouteProtocolType_egp RouteProtocolType = "egp"

    RouteProtocolType_ggp RouteProtocolType = "ggp"

    RouteProtocolType_hello RouteProtocolType = "hello"

    RouteProtocolType_rip RouteProtocolType = "rip"

    RouteProtocolType_isIs RouteProtocolType = "isIs"

    RouteProtocolType_esIs RouteProtocolType = "esIs"

    RouteProtocolType_ciscoIgrp RouteProtocolType = "ciscoIgrp"

    RouteProtocolType_bbnSpfIgp RouteProtocolType = "bbnSpfIgp"

    RouteProtocolType_ospf RouteProtocolType = "ospf"

    RouteProtocolType_bgp RouteProtocolType = "bgp"

    RouteProtocolType_idpr RouteProtocolType = "idpr"

    RouteProtocolType_ciscoEigrp RouteProtocolType = "ciscoEigrp"

    RouteProtocolType_dvmrp RouteProtocolType = "dvmrp"
)

// MrouteProtocolType represents those protocols need be supported.
type MrouteProtocolType string

const (
    MrouteProtocolType_other MrouteProtocolType = "other"

    MrouteProtocolType_local MrouteProtocolType = "local"

    MrouteProtocolType_netmgmt MrouteProtocolType = "netmgmt"

    MrouteProtocolType_dvmrp MrouteProtocolType = "dvmrp"

    MrouteProtocolType_mospf MrouteProtocolType = "mospf"

    MrouteProtocolType_pimSparseDense MrouteProtocolType = "pimSparseDense"

    MrouteProtocolType_cbt MrouteProtocolType = "cbt"

    MrouteProtocolType_pimSparseMode MrouteProtocolType = "pimSparseMode"

    MrouteProtocolType_pimDenseMode MrouteProtocolType = "pimDenseMode"

    MrouteProtocolType_igmpOnly MrouteProtocolType = "igmpOnly"

    MrouteProtocolType_bgmp MrouteProtocolType = "bgmp"

    MrouteProtocolType_msdp MrouteProtocolType = "msdp"
)

