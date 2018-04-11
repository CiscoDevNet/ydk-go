// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-nve package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-ifmgr-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_nve_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_nve_cfg"))
}

// HostReachProtocol represents Host reach protocol
type HostReachProtocol string

const (
    // Use BGP EVPN for VxLAN tunnel endpoint
    // reachability
    HostReachProtocol_bgp HostReachProtocol = "bgp"
)

// VxlanUdpPortEnum represents Vxlan udp port enum
type VxlanUdpPortEnum string

const (
    // IETF defined UDP port for VxLAN
    VxlanUdpPortEnum_ietf_udp_port VxlanUdpPortEnum = "ietf-udp-port"

    // UDP port for iVxLAN
    VxlanUdpPortEnum_ivx_lan_udp_port VxlanUdpPortEnum = "ivx-lan-udp-port"
)

// OverlayEncapEnum represents Overlay encap enum
type OverlayEncapEnum string

const (
    // VxLAN Encapsulation
    OverlayEncapEnum_vx_lan_encapsulation OverlayEncapEnum = "vx-lan-encapsulation"

    // Soft GRE Encapsulation
    OverlayEncapEnum_soft_gre_encapsulation OverlayEncapEnum = "soft-gre-encapsulation"
)

// UnknownUnicastFloodingEnum represents Unknown unicast flooding enum
type UnknownUnicastFloodingEnum string

const (
    // Suppress unknown unicast flooding
    UnknownUnicastFloodingEnum_suppress_uuf UnknownUnicastFloodingEnum = "suppress-uuf"
)

// LoadBalanceEnum represents Load balance enum
type LoadBalanceEnum string

const (
    // Per evi load balance mode
    LoadBalanceEnum_per_evi LoadBalanceEnum = "per-evi"
)

// IrProtocolEnum represents Ir protocol enum
type IrProtocolEnum string

const (
    // Use BGP Protocol for Ingress-Replication
    IrProtocolEnum_bgp IrProtocolEnum = "bgp"
)

