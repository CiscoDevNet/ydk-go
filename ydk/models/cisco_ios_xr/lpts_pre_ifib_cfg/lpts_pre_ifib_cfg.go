// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-pre-ifib package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-lpts-lib-cfg
//   Cisco-IOS-XR-config-mda-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package lpts_pre_ifib_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_pre_ifib_cfg"))
}

// LptsDynamicFlowConfig represents Lpts dynamic flow config
type LptsDynamicFlowConfig string

const (
    // LPTS Flows Limit
    LptsDynamicFlowConfig_flows_config LptsDynamicFlowConfig = "flows-config"

    // Platform Limit
    LptsDynamicFlowConfig_platform_config LptsDynamicFlowConfig = "platform-config"
)

// LptsPreIFibPrecedenceNumber represents Lpts pre i fib precedence number
type LptsPreIFibPrecedenceNumber string

const (
    // Match packets with critical precedence
    LptsPreIFibPrecedenceNumber_critical LptsPreIFibPrecedenceNumber = "critical"

    // Match packets with flash precedence
    LptsPreIFibPrecedenceNumber_flash LptsPreIFibPrecedenceNumber = "flash"

    // Match packets with flash override precedence
    LptsPreIFibPrecedenceNumber_flash_override LptsPreIFibPrecedenceNumber = "flash-override"

    // Match packets with immediate precedence
    LptsPreIFibPrecedenceNumber_immediate LptsPreIFibPrecedenceNumber = "immediate"

    // Match packets with internetwork control
    // precedence
    LptsPreIFibPrecedenceNumber_internet LptsPreIFibPrecedenceNumber = "internet"

    // Match packets with network control precedence
    LptsPreIFibPrecedenceNumber_network LptsPreIFibPrecedenceNumber = "network"

    // Match packets with priority precedence
    LptsPreIFibPrecedenceNumber_priority LptsPreIFibPrecedenceNumber = "priority"

    // Match packets with routine precedence
    LptsPreIFibPrecedenceNumber_routine LptsPreIFibPrecedenceNumber = "routine"
)

// LptsFlow represents Lpts flow
type LptsFlow string

const (
    // Invalid flow type used for fallback
    // configuration
    LptsFlow_config_default LptsFlow = "config-default"

    // L2TPv2 Fragments
    LptsFlow_l2tpv2_fragment LptsFlow = "l2tpv2-fragment"

    // Fragments
    LptsFlow_fragment LptsFlow = "fragment"

    // OSPF multicast packets on configured interfaces
    LptsFlow_ospf_multicast_known LptsFlow = "ospf-multicast-known"

    // OSPF multicast packets on unconfigured (or
    // newly-configured) interfaces
    LptsFlow_ospf_multicast_default LptsFlow = "ospf-multicast-default"

    // OSPF unicast packets
    LptsFlow_ospf_unicast_known LptsFlow = "ospf-unicast-known"

    // OSPF unicast packets
    LptsFlow_ospf_unicast_default LptsFlow = "ospf-unicast-default"

    // IS-IS packets on configured interfaces
    LptsFlow_isis_known LptsFlow = "isis-known"

    // IS-IS packets on unconfigured (or
    // newly-configured) interfaces
    LptsFlow_isis_default LptsFlow = "isis-default"

    // BFD packets on configured interfaces
    LptsFlow_bfd_known LptsFlow = "bfd-known"

    // BFD packets on unconfigured (or
    // newly-configured) interfaces
    LptsFlow_bfd_default LptsFlow = "bfd-default"

    // BFD multipath packets on configured interfaces
    LptsFlow_bfd_multipath_known LptsFlow = "bfd-multipath-known"

    // BFD multipath packets on multiple configured
    // interfaces
    LptsFlow_bfd_multipath0 LptsFlow = "bfd-multipath0"

    // BFD packets over Logical Bundle on configured
    // interfaces
    LptsFlow_bfd_blb_known LptsFlow = "bfd-blb-known"

    // BFD packets over Logical Bundle 0
    LptsFlow_bfd_blb0 LptsFlow = "bfd-blb0"

    // BFD packets over Single Path 0
    LptsFlow_bfd_sp0 LptsFlow = "bfd-sp0"

    // Packets from established BGP peering sessions
    LptsFlow_bgp_known LptsFlow = "bgp-known"

    // Packets from a configured BGP peer (SYNs or
    // newly-established sessions)
    LptsFlow_bgp_config_peer LptsFlow = "bgp-config-peer"

    // Packets from unconfigured, newly-configured, or
    // wild-card BGP peer
    LptsFlow_bgp_default LptsFlow = "bgp-default"

    // PIM multicast packets on configured interfaces
    LptsFlow_pim_multicast_default LptsFlow = "pim-multicast-default"

    // PIM multicast packets on unconfigured (or
    // newly-configured) interfaces
    LptsFlow_pim_multicast_known LptsFlow = "pim-multicast-known"

    // PIM unicast packets
    LptsFlow_pim_unicast LptsFlow = "pim-unicast"

    // IGMP packets
    LptsFlow_igmp LptsFlow = "igmp"

    // ICMP or ICMPv6 packets with local interest
    LptsFlow_icmp_local LptsFlow = "icmp-local"

    // ICMP or ICMPv6 packets of interest to
    // applications
    LptsFlow_icmp_app LptsFlow = "icmp-app"

    // ICMP or ICMPv6 packets that are used for
    // control/signalling purpose
    LptsFlow_icmp_control LptsFlow = "icmp-control"

    // Other ICMP or ICMPv6 packets (may be of recent
    // interest to applications)
    LptsFlow_icmp_default LptsFlow = "icmp-default"

    // ICMP or ICMPv6 echo reply packets (when
    // specific entry not present)
    LptsFlow_icmp_app_default LptsFlow = "icmp-app-default"

    // Packets from an established LDP TCP peering
    // session
    LptsFlow_ldp_tcp_known LptsFlow = "ldp-tcp-known"

    // Packets from a configured LDP TCP peer (SYNs or
    // newly-established sessions)
    LptsFlow_ldp_tcp_config_peer LptsFlow = "ldp-tcp-config-peer"

    // Packets from an unconfigured, newly-configured
    // or wild-card LDP TCP peer
    LptsFlow_ldp_tcp_default LptsFlow = "ldp-tcp-default"

    // Unicast LDP UDP packets
    LptsFlow_ldp_udp LptsFlow = "ldp-udp"

    // Packets sent to the all-routers multicast
    // address (includes LDP UDP multicast)
    LptsFlow_all_routers LptsFlow = "all-routers"

    // Packets from an established LMP TCP peering
    // session
    LptsFlow_lmp_tcp_known LptsFlow = "lmp-tcp-known"

    // Packets from a configured LMP TCP peer (SYNs or
    // newly-established sessions)
    LptsFlow_lmp_tcp_config_peer LptsFlow = "lmp-tcp-config-peer"

    // Packets from an unconfigured, newly-configured
    // or wild-card LMP TCP peer
    LptsFlow_lmp_tcp_default LptsFlow = "lmp-tcp-default"

    // Unicast LMP UDP packets
    LptsFlow_lmp_udp LptsFlow = "lmp-udp"

    // RSVP-over-UDP packets
    LptsFlow_rsvp_udp LptsFlow = "rsvp-udp"

    // RSVP (IP protocol 46) packets
    LptsFlow_rsvp_default LptsFlow = "rsvp-default"

    // RSVP (IP protocol 46) packets
    LptsFlow_rsvp_known LptsFlow = "rsvp-known"

    // IKE packets
    LptsFlow_ike LptsFlow = "ike"

    // AH or ESP packets with known SPIs
    LptsFlow_ipsec_known LptsFlow = "ipsec-known"

    // AH or ESP packets with unknown or
    // newly-configured SPIs
    LptsFlow_ipsec_default LptsFlow = "ipsec-default"

    // AH or ESP fragmented packets
    LptsFlow_ipsec_fragment LptsFlow = "ipsec-fragment"

    // Packets from an established MSDP session
    LptsFlow_msdp_known LptsFlow = "msdp-known"

    // Packets from a configured MSDP peer
    LptsFlow_msdp_config_peer LptsFlow = "msdp-config-peer"

    // Packets from an uncofigured, newly-configured
    // or wild-card MSDP peer
    LptsFlow_msdp_default LptsFlow = "msdp-default"

    // SNMP packets
    LptsFlow_snmp LptsFlow = "snmp"

    // Packets from an established SSH session
    LptsFlow_ssh_known LptsFlow = "ssh-known"

    // Packets from a new or newly-established SSH
    // session
    LptsFlow_ssh_default LptsFlow = "ssh-default"

    // Packets from an established HTTP session
    LptsFlow_http_known LptsFlow = "http-known"

    // Packets from a new or newly-established HTTP
    // session
    LptsFlow_http_default LptsFlow = "http-default"

    // Packets from an established SHTTP session
    LptsFlow_shttp_known LptsFlow = "shttp-known"

    // Packets from a new or newly-established SSHTP
    // session
    LptsFlow_shttp_default LptsFlow = "shttp-default"

    // Packets from an established TELNET session
    LptsFlow_telnet_known LptsFlow = "telnet-known"

    // Packets from a new or newly-established TELNET
    // session
    LptsFlow_telnet_default LptsFlow = "telnet-default"

    // Packets from an established CSS session
    LptsFlow_css_known LptsFlow = "css-known"

    // Packets from a new or newly-established CSS
    // session
    LptsFlow_css_default LptsFlow = "css-default"

    // Packets from an established rsh session
    LptsFlow_rsh_known LptsFlow = "rsh-known"

    // Packets from a new or newly-established rsh
    // session
    LptsFlow_rsh_default LptsFlow = "rsh-default"

    // Packets for established UDP sessions
    LptsFlow_udp_known LptsFlow = "udp-known"

    // Packets for configured UDP services
    LptsFlow_udp_listen LptsFlow = "udp-listen"

    // Packets for configured UDP-based protocol
    // sessions
    LptsFlow_udp_config_peer LptsFlow = "udp-config-peer"

    // Packets for unconfigured or newly-configured
    // UDP services
    LptsFlow_udp_default LptsFlow = "udp-default"

    // Packets for established TCP sessions
    LptsFlow_tcp_known LptsFlow = "tcp-known"

    // Packets for configured TCP services
    LptsFlow_tcp_listen LptsFlow = "tcp-listen"

    // Packets for configured TCP peers
    LptsFlow_tcp_config_peer LptsFlow = "tcp-config-peer"

    // Packets for unconfigured or newly-configured
    // TCP services
    LptsFlow_tcp_default LptsFlow = "tcp-default"

    // Packets for configured multicast groups
    LptsFlow_multicast_known LptsFlow = "multicast-known"

    // Packets for unconfigured or newly-configured
    // multicast groups
    LptsFlow_multicast_default LptsFlow = "multicast-default"

    // Packets for configured IP protocols
    LptsFlow_raw_listen LptsFlow = "raw-listen"

    // Packets for unconfigured or newly-configured
    // IPv4 or IPv6 protocols
    LptsFlow_raw_default LptsFlow = "raw-default"

    // IP SLA packets destined to squid Q #4 for
    // timestamping by squid driver
    LptsFlow_ipsla LptsFlow = "ipsla"

    // EIGRP packets.
    LptsFlow_eigrp LptsFlow = "eigrp"

    // RIP packets.
    LptsFlow_rip LptsFlow = "rip"

    // L2TPv3 packets.
    LptsFlow_l2tpv3 LptsFlow = "l2tpv3"

    // PCEP packets.
    LptsFlow_pcep_tcp_default LptsFlow = "pcep-tcp-default"

    // GRE packets.
    LptsFlow_gre LptsFlow = "gre"

    // VRRP Packets.
    LptsFlow_vrrp LptsFlow = "vrrp"

    // HSRP Packets.
    LptsFlow_hsrp LptsFlow = "hsrp"

    // MPLS ping packet coming or arriving from 3503
    // port
    LptsFlow_mpls_ping LptsFlow = "mpls-ping"

    // L2TPv2 default packets.
    LptsFlow_l2tpv2_default LptsFlow = "l2tpv2-default"

    // L2TPv2 known packets.
    LptsFlow_l2tpv2_known LptsFlow = "l2tpv2-known"

    // DNS packets.
    LptsFlow_dns LptsFlow = "dns"

    // RADIUS packets.
    LptsFlow_radius LptsFlow = "radius"

    // TACACS packets.
    LptsFlow_tacacs LptsFlow = "tacacs"

    // NTP packets received at 123 port number any
    // address.
    LptsFlow_ntp_default LptsFlow = "ntp-default"

    // NTP packets received at 123 port number known
    // address.
    LptsFlow_ntp_known LptsFlow = "ntp-known"

    // Mobile IPV6 packets.
    LptsFlow_mobile_ipv6 LptsFlow = "mobile-ipv6"

    // AMT packets received at UDP port number 2268.
    LptsFlow_amt LptsFlow = "amt"

    // SDAC TCP packets.
    LptsFlow_sdac_tcp LptsFlow = "sdac-tcp"

    // RADIUS Change of Authorization packets.
    LptsFlow_radius_coa LptsFlow = "radius-coa"

    // REL UDP packets.
    LptsFlow_rel_udp LptsFlow = "rel-udp"

    // DHCP IPV4 packets.
    LptsFlow_dhcp4 LptsFlow = "dhcp4"

    // DHCP IPV6 packets.
    LptsFlow_dhcp6 LptsFlow = "dhcp6"

    // ONEPK packets.
    LptsFlow_onepk LptsFlow = "onepk"

    // EXR packets.
    LptsFlow_exr LptsFlow = "exr"

    // IETF BFD packets over Logical Bundle.
    LptsFlow_bob_ietf LptsFlow = "bob-ietf"

    // XIPC Throttle Flow.
    LptsFlow_xipc_throt LptsFlow = "xipc-throt"

    // Platform Limit.
    LptsFlow_platform_limit LptsFlow = "platform-limit"
)

// Lptsafi represents Lptsafi
type Lptsafi string

const (
    // IPv4 type
    Lptsafi_ipv4 Lptsafi = "ipv4"

    // IPv6 type
    Lptsafi_ipv6 Lptsafi = "ipv6"
)

