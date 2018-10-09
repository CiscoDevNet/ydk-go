// This module contains a collection of YANG definitions
// for some common IOS structures
// Copyright (c) 2018 by Cisco Systems, Inc.
// All rights reserved.
package ios_common_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ios_common_oper"))
}

// IosLinktype represents Link Type
type IosLinktype string

const (
    // invalid link type
    IosLinktype_ios_linktype_illegal IosLinktype = "ios-linktype-illegal"

    // IP ARP
    IosLinktype_ios_linktype_arp IosLinktype = "ios-linktype-arp"

    // Reverse ARP
    IosLinktype_ios_linktype_rarp IosLinktype = "ios-linktype-rarp"

    // Xerox ARP
    IosLinktype_ios_linktype_xarp IosLinktype = "ios-linktype-xarp"

    // HP's 802 version of ARP
    IosLinktype_ios_linktype_probe IosLinktype = "ios-linktype-probe"

    // Ethernet loopback
    IosLinktype_ios_linktype_loop IosLinktype = "ios-linktype-loop"

    // cisco Systems Serial ARP
    IosLinktype_ios_linktype_address IosLinktype = "ios-linktype-address"

    // Internet Protocol (IP)
    IosLinktype_ios_linktype_ip IosLinktype = "ios-linktype-ip"

    // Xerox PARC Universal Protocol (PUP)
    IosLinktype_ios_linktype_pup IosLinktype = "ios-linktype-pup"

    // Available Unused Slot
    IosLinktype_ios_linktype_9_unused IosLinktype = "ios-linktype-9-unused"

    // XNS protocol (Xerox original)
    IosLinktype_ios_linktype_xns IosLinktype = "ios-linktype-xns"

    // Novell IPX
    IosLinktype_ios_linktype_novell IosLinktype = "ios-linktype-novell"

    // Apollo Domain
    IosLinktype_ios_linktype_apollo IosLinktype = "ios-linktype-apollo"

    // Banyan Vines IP
    IosLinktype_ios_linktype_vines IosLinktype = "ios-linktype-vines"

    // X.2[89] PAD protocol
    IosLinktype_ios_linktype_pad IosLinktype = "ios-linktype-pad"

    // X.25 for purists
    IosLinktype_ios_linktype_x25 IosLinktype = "ios-linktype-x25"

    // AppleTalk - Extended DDP
    IosLinktype_ios_linktype_appletalk IosLinktype = "ios-linktype-appletalk"

    // Appletalk ARP
    IosLinktype_ios_linktype_aarp IosLinktype = "ios-linktype-aarp"

    // DEC spanning tree
    IosLinktype_ios_linktype_dec_spanning IosLinktype = "ios-linktype-dec-spanning"

    // IEEE spanning tree BPDU
    IosLinktype_ios_linktype_ieee_spanning IosLinktype = "ios-linktype-ieee-spanning"

    // NBF NetBios Frames
    IosLinktype_ios_linktype_nbf IosLinktype = "ios-linktype-nbf"

    // Unused Link Type
    IosLinktype_ios_linktype_21_unused IosLinktype = "ios-linktype-21-unused"

    // DECnet Phase IV node unicast
    IosLinktype_ios_linktype_decnet IosLinktype = "ios-linktype-decnet"

    // (Old) DECnet PhaseIV router multicast
    IosLinktype_ios_linktype_decnet_router_l1 IosLinktype = "ios-linktype-decnet-router-l1"

    // DECnet Phase IV node multicast
    IosLinktype_ios_linktype_decnet_node IosLinktype = "ios-linktype-decnet-node"

    // OSI 8473 ES-IS 9542, etc.
    IosLinktype_ios_linktype_clns IosLinktype = "ios-linktype-clns"

    // ES-IS Multicast - all end systems
    IosLinktype_ios_linktype_clns_all_es IosLinktype = "ios-linktype-clns-all-es"

    // ES-IS Multicast - all intermed.
    IosLinktype_ios_linktype_clns_all_is IosLinktype = "ios-linktype-clns-all-is"

    // CLNS & ES-IS Broadcasts
    IosLinktype_ios_linktype_clns_bcast IosLinktype = "ios-linktype-clns-bcast"

    // X.25 over TCP
    IosLinktype_ios_linktype_xot IosLinktype = "ios-linktype-xot"

    // Protocol PPP
    IosLinktype_ios_linktype_ppp IosLinktype = "ios-linktype-ppp"

    // DEC LAT terminal protocol
    IosLinktype_ios_linktype_lat IosLinktype = "ios-linktype-lat"

    // Banyan Vines Echo Protocol
    IosLinktype_ios_linktype_vines_echo IosLinktype = "ios-linktype-vines-echo"

    // Banyan Vines Loopback Protocol
    IosLinktype_ios_linktype_vines_loop IosLinktype = "ios-linktype-vines-loop"

    // AppleTalk short DDP
    IosLinktype_ios_linktype_atalk_short IosLinktype = "ios-linktype-atalk-short"

    // DEC MOP booting protocol
    IosLinktype_ios_linktype_mop_boot IosLinktype = "ios-linktype-mop-boot"

    // DEC MOP console protocol
    IosLinktype_ios_linktype_mop_console IosLinktype = "ios-linktype-mop-console"

    // RSRB raw
    IosLinktype_ios_linktype_rsrb IosLinktype = "ios-linktype-rsrb"

    // transparent bridge traffic
    IosLinktype_ios_linktype_bridge IosLinktype = "ios-linktype-bridge"

    // Stunnel
    IosLinktype_ios_linktype_stun IosLinktype = "ios-linktype-stun"

    // Frame Relay
    IosLinktype_ios_linktype_fr_arp IosLinktype = "ios-linktype-fr-arp"

    // SMDS own ARP
    IosLinktype_ios_linktype_smds_arp IosLinktype = "ios-linktype-smds-arp"

    // MAC packets
    IosLinktype_ios_linktype_mac IosLinktype = "ios-linktype-mac"

    // IBM Network Management packets
    IosLinktype_ios_linktype_ibmnm IosLinktype = "ios-linktype-ibmnm"

    // Ultranet
    IosLinktype_ios_linktype_uarp IosLinktype = "ios-linktype-uarp"

    // Ultranet
    IosLinktype_ios_linktype_ultra_hello IosLinktype = "ios-linktype-ultra-hello"

    // X.25 services
    IosLinktype_ios_linktype_x25service IosLinktype = "ios-linktype-x25service"

    // Uncompressed TCP/IP
    IosLinktype_ios_linktype_uncompressed_tcp IosLinktype = "ios-linktype-uncompressed-tcp"

    // Compressed TCP/IP
    IosLinktype_ios_linktype_compressed_tcp IosLinktype = "ios-linktype-compressed-tcp"

    // LLC type II
    IosLinktype_ios_linktype_llc2 IosLinktype = "ios-linktype-llc2"

    // CMNS--X.25 over non-serial media
    IosLinktype_ios_linktype_cmns IosLinktype = "ios-linktype-cmns"

    // ISIS - All Level 1 ISes
    IosLinktype_ios_linktype_isis_all_l1_is IosLinktype = "ios-linktype-isis-all-l1-is"

    // ISIS - All Level 2 ISes
    IosLinktype_ios_linktype_isis_all_l2_is IosLinktype = "ios-linktype-isis-all-l2-is"

    // FR switching over IP tunnel
    IosLinktype_ios_linktype_fr_switch IosLinktype = "ios-linktype-fr-switch"

    // Ph-IV Prime all routers
    IosLinktype_ios_linktype_decnet_prime_router IosLinktype = "ios-linktype-decnet-prime-router"

    // source route bridge traffic
    IosLinktype_ios_linktype_srb IosLinktype = "ios-linktype-srb"

    // qllc (SNA over X.25)
    IosLinktype_ios_linktype_qllc IosLinktype = "ios-linktype-qllc"

    // X.25 multiprotocol encapsulation
    IosLinktype_ios_linktype_x25_multi_enc IosLinktype = "ios-linktype-x25-multi-enc"

    // Lan Extension Protocol
    IosLinktype_ios_linktype_lex IosLinktype = "ios-linktype-lex"

    // Lex-NCP + LEX Remote commands
    IosLinktype_ios_linktype_lex_rcmd IosLinktype = "ios-linktype-lex-rcmd"

    // DECnet PhaseIV L2 router multcast
    IosLinktype_ios_linktype_decnet_router_l2 IosLinktype = "ios-linktype-decnet-router-l2"

    // Frames from RSRB vring->CLS vring
    IosLinktype_ios_linktype_cls IosLinktype = "ios-linktype-cls"

    // Start dialing, don't send frame
    IosLinktype_ios_linktype_snapshot IosLinktype = "ios-linktype-snapshot"

    // Data Link Switching traffic
    IosLinktype_ios_linktype_dlsw IosLinktype = "ios-linktype-dlsw"

    // 802.10 Secure Data Exchange
    IosLinktype_ios_linktype_sde IosLinktype = "ios-linktype-sde"

    // Cisco Discovery Protocol
    IosLinktype_ios_linktype_cdp IosLinktype = "ios-linktype-cdp"

    // Compression Control Protocol PPP
    IosLinktype_ios_linktype_ppp_compression IosLinktype = "ios-linktype-ppp-compression"

    // Local LNM packet traffic
    IosLinktype_ios_linktype_nmp IosLinktype = "ios-linktype-nmp"

    // block serial tunnel
    IosLinktype_ios_linktype_bstun IosLinktype = "ios-linktype-bstun"

    // Out-of-band IPC
    IosLinktype_ios_linktype_ipc IosLinktype = "ios-linktype-ipc"

    // CIP statistics
    IosLinktype_ios_linktype_love IosLinktype = "ios-linktype-love"

    // CIP asynchronous cfg acks
    IosLinktype_ios_linktype_cfgack IosLinktype = "ios-linktype-cfgack"

    // APPN HPR ANR traffic
    IosLinktype_ios_linktype_appn_anr IosLinktype = "ios-linktype-appn-anr"

    // Multilink protocol
    IosLinktype_ios_linktype_multilink IosLinktype = "ios-linktype-multilink"

    // Next Hop Resolution Protocol (NHRP)
    IosLinktype_ios_linktype_nhrp IosLinktype = "ios-linktype-nhrp"

    // MAC frames are express buffered
    IosLinktype_ios_linktype_mac_exp IosLinktype = "ios-linktype-mac-exp"

    // Cisco Group Management Protocol
    IosLinktype_ios_linktype_cgmp IosLinktype = "ios-linktype-cgmp"

    // Cisco VLAN Transport Protocol
    IosLinktype_ios_linktype_vtp IosLinktype = "ios-linktype-vtp"

    // Combinet proprietary encaps
    IosLinktype_ios_linktype_cpp IosLinktype = "ios-linktype-cpp"

    // IP version 6
    IosLinktype_ios_linktype_ipv6 IosLinktype = "ios-linktype-ipv6"

    // Asynchronous Transfer Mode (ATM)
    IosLinktype_ios_linktype_atm IosLinktype = "ios-linktype-atm"

    // Cisco Inter Switch Link
    IosLinktype_ios_linktype_isl IosLinktype = "ios-linktype-isl"

    // Bandwidth Allocation Protocol
    IosLinktype_ios_linktype_bap IosLinktype = "ios-linktype-bap"

    // Compressed RTP-UDP-IP
    IosLinktype_ios_linktype_compressed_rtp IosLinktype = "ios-linktype-compressed-rtp"

    // Compressed UDP-IP
    IosLinktype_ios_linktype_compressed_udp IosLinktype = "ios-linktype-compressed-udp"

    // Unompressed RTP-UDP-IP
    IosLinktype_ios_linktype_uncompressed_udp IosLinktype = "ios-linktype-uncompressed-udp"

    // RTP context error
    IosLinktype_ios_linktype_context_status IosLinktype = "ios-linktype-context-status"

    // FR switching over serial line
    // internal packet identification only
    IosLinktype_ios_linktype_fr_switch_serial IosLinktype = "ios-linktype-fr-switch-serial"

    // Cat5k inband ipc
    IosLinktype_ios_linktype_c5_ibipc IosLinktype = "ios-linktype-c5-ibipc"

    // NLSP multicast hello/update
    IosLinktype_ios_linktype_nlsp_multicast IosLinktype = "ios-linktype-nlsp-multicast"

    // Tag switched packet
    IosLinktype_ios_linktype_tag IosLinktype = "ios-linktype-tag"

    // Multicast TAG (MTAG)
    IosLinktype_ios_linktype_mtag IosLinktype = "ios-linktype-mtag"

    // Airline Protocol Support
    IosLinktype_ios_linktype_alps IosLinktype = "ios-linktype-alps"

    // Cisco DRIP
    IosLinktype_ios_linktype_drip IosLinktype = "ios-linktype-drip"

    // Voice over ATM
    IosLinktype_ios_linktype_voice IosLinktype = "ios-linktype-voice"

    // FR-ATM network interworking
    IosLinktype_ios_linktype_fr_atm IosLinktype = "ios-linktype-fr-atm"

    // layer3 switching on earl2
    IosLinktype_ios_linktype_fcp IosLinktype = "ios-linktype-fcp"

    // Non real time voice
    IosLinktype_ios_linktype_voice_no_rt IosLinktype = "ios-linktype-voice-no-rt"

    // VLAN Bridge spanning tree BPDU
    IosLinktype_ios_linktype_vlan_br_spanning IosLinktype = "ios-linktype-vlan-br-spanning"

    // IEEE 802.1Q
    IosLinktype_ios_linktype_dot1q IosLinktype = "ios-linktype-dot1q"

    // cisco ipc msg.
    IosLinktype_ios_linktype_cisco_ipc IosLinktype = "ios-linktype-cisco-ipc"

    // Cisco VSI Protocol
    IosLinktype_ios_linktype_vsi IosLinktype = "ios-linktype-vsi"

    // for manual ISDN call
    IosLinktype_ios_linktype_isdn_manual_call IosLinktype = "ios-linktype-isdn-manual-call"

    // IPHC compression
    IosLinktype_ios_linktype_comp_non_tcp IosLinktype = "ios-linktype-comp-non-tcp"

    // IPHC compression
    IosLinktype_ios_linktype_comp_tcp_nodelta IosLinktype = "ios-linktype-comp-tcp-nodelta"

    // FR-ATM service interworking
    IosLinktype_ios_linktype_fr_atm_srv IosLinktype = "ios-linktype-fr-atm-srv"

    // FR end-to-end keepalive
    IosLinktype_ios_linktype_fr_eek IosLinktype = "ios-linktype-fr-eek"

    // Test type for EHSA
    IosLinktype_ios_linktype_ehsa IosLinktype = "ios-linktype-ehsa"

    // Multicast Shortcut PDU
    IosLinktype_ios_linktype_mscp IosLinktype = "ios-linktype-mscp"

    // Cat6k - linecard control
    IosLinktype_ios_linktype_scp IosLinktype = "ios-linktype-scp"

    // PPP over Ethernet Discovery
    IosLinktype_ios_linktype_pppoe_discovery IosLinktype = "ios-linktype-pppoe-discovery"

    // PPP over Ethernet Session
    IosLinktype_ios_linktype_pppoe_session IosLinktype = "ios-linktype-pppoe-session"

    // GPRS l2rly
    IosLinktype_ios_linktype_l2rly IosLinktype = "ios-linktype-l2rly"

    // GPRS GTP
    IosLinktype_ios_linktype_gtp IosLinktype = "ios-linktype-gtp"

    // Compressed RTP-UDP-IP 16 bit CID
    IosLinktype_ios_linktype_comp_rtp_16 IosLinktype = "ios-linktype-comp-rtp-16"

    // Compressed RTP-UDP-IP 16 bit CID
    IosLinktype_ios_linktype_comp_udp_16 IosLinktype = "ios-linktype-comp-udp-16"

    // ATM PDU switching
    IosLinktype_ios_linktype_atm_switch IosLinktype = "ios-linktype-atm-switch"

    // Port Aggregation Protocol
    IosLinktype_ios_linktype_pagp IosLinktype = "ios-linktype-pagp"

    // Shared Spanning Tree BPDU
    IosLinktype_ios_linktype_sstp IosLinktype = "ios-linktype-sstp"

    // STP Root Link Query Request
    IosLinktype_ios_linktype_rlq_req IosLinktype = "ios-linktype-rlq-req"

    // STP Root Link Query Response
    IosLinktype_ios_linktype_rlq_resp IosLinktype = "ios-linktype-rlq-resp"

    // Dynamic Trunk Protocol (DTP) PDU
    IosLinktype_ios_linktype_dyntrk IosLinktype = "ios-linktype-dyntrk"

    // layer3 switching internal flush
    IosLinktype_ios_linktype_mls_ip IosLinktype = "ios-linktype-mls-ip"

    // DTP PDUs sent w/ SW ISL encap
    IosLinktype_ios_linktype_dyntrk_encap IosLinktype = "ios-linktype-dyntrk-encap"

    // UniDirectional Link Protocol
    IosLinktype_ios_linktype_udld IosLinktype = "ios-linktype-udld"

    // Spatial Reuse Protocol (SRP)
    IosLinktype_ios_linktype_srp IosLinktype = "ios-linktype-srp"

    // Router-port Group Management Protocol (RGMP)
    IosLinktype_ios_linktype_rgmp IosLinktype = "ios-linktype-rgmp"

    // Multi Instance Spanning Tree
    IosLinktype_ios_linktype_mistp IosLinktype = "ios-linktype-mistp"

    // Compressed Frame Relay
    IosLinktype_ios_linktype_compressed_fr IosLinktype = "ios-linktype-compressed-fr"

    // Vito SCP protocol
    IosLinktype_ios_linktype_vito_scp IosLinktype = "ios-linktype-vito-scp"

    // Constellation Forwarding Table Edit Protocol
    IosLinktype_ios_linktype_const_ftep IosLinktype = "ios-linktype-const-ftep"

    // VLAN PDU switching
    IosLinktype_ios_linktype_vlan_switch IosLinktype = "ios-linktype-vlan-switch"

    // ATOM (Process) switching
    IosLinktype_ios_linktype_atom IosLinktype = "ios-linktype-atom"

    // Raw switching
    IosLinktype_ios_linktype_raw IosLinktype = "ios-linktype-raw"

    // Optical Supervisory Channel Protocol
    IosLinktype_ios_linktype_oscp IosLinktype = "ios-linktype-oscp"

    // For Diag lpback via Ethernet chip
    IosLinktype_ios_linktype_diag IosLinktype = "ios-linktype-diag"

    // STP DUF STACK discovery
    IosLinktype_ios_linktype_stack_disc IosLinktype = "ios-linktype-stack-disc"

    // STP DUF Fast Transition
    IosLinktype_ios_linktype_fast_trans IosLinktype = "ios-linktype-fast-trans"

    // 802.1X Port Access Entity Type
    IosLinktype_ios_linktype_dot1x IosLinktype = "ios-linktype-dot1x"

    // FR fragmentation
    IosLinktype_ios_linktype_fr_frag IosLinktype = "ios-linktype-fr-frag"

    // 802.3AD Slow Protocol
    IosLinktype_ios_linktype_slow_proto IosLinktype = "ios-linktype-slow-proto"

    // sync damage please name me
    IosLinktype_ios_linktype_eompls IosLinktype = "ios-linktype-eompls"

    // sync damage please name me
    IosLinktype_ios_linktype_ieee_dot1q IosLinktype = "ios-linktype-ieee-dot1q"

    // sync damage please name me
    IosLinktype_ios_linktype_voice_nrt IosLinktype = "ios-linktype-voice-nrt"

    // sync damage please name me
    IosLinktype_ios_linktype_uti_raw IosLinktype = "ios-linktype-uti-raw"

    // ACE IPsec accelerator control
    IosLinktype_ios_linktype_hdlc_hapi IosLinktype = "ios-linktype-hdlc-hapi"

    // IPSec acceleration for CWAN
    IosLinktype_ios_linktype_cwan_ipsec IosLinktype = "ios-linktype-cwan-ipsec"

    // Shared Spanning Tree BPDU
    IosLinktype_ios_linktype_sstp_spanning IosLinktype = "ios-linktype-sstp-spanning"

    // FR SAA round trip measurement
    IosLinktype_ios_linktype_fr_saa IosLinktype = "ios-linktype-fr-saa"

    // Frame Relay over Ethernet
    IosLinktype_ios_linktype_froe IosLinktype = "ios-linktype-froe"

    // Forwarded VPDN packets
    IosLinktype_ios_linktype_vpdn IosLinktype = "ios-linktype-vpdn"

    // SAA SLM Link Type
    IosLinktype_ios_linktype_slm_saa IosLinktype = "ios-linktype-slm-saa"

    // Redundancy Facility keepalives
    IosLinktype_ios_linktype_rf_kpa IosLinktype = "ios-linktype-rf-kpa"

    // Router blade control protocol
    IosLinktype_ios_linktype_rbcp IosLinktype = "ios-linktype-rbcp"

    // P2IPC LinkType
    IosLinktype_ios_linktype_p2ipc IosLinktype = "ios-linktype-p2ipc"

    // P2IPC LinkType for MPLS controller
    IosLinktype_ios_linktype_p2ipc_mpls IosLinktype = "ios-linktype-p2ipc-mpls"

    // Xconnect (Process switching)
    IosLinktype_ios_linktype_xconnect IosLinktype = "ios-linktype-xconnect"

    // FR SLM Link Type
    IosLinktype_ios_linktype_slm_fr_saa IosLinktype = "ios-linktype-slm-fr-saa"

    // CDMA RP
    IosLinktype_ios_linktype_cdma_rp IosLinktype = "ios-linktype-cdma-rp"

    // Deprecated
    // do not reuse due to potential ISSU conflict
    IosLinktype_ios_linktype_pw_ip_deprecated IosLinktype = "ios-linktype-pw-ip-deprecated"

    // VJ Uncompressed TCP/IP
    IosLinktype_ios_linktype_vj_uncomp_tcp IosLinktype = "ios-linktype-vj-uncomp-tcp"

    // VJ Compressed TCP/IP
    IosLinktype_ios_linktype_vj_comp_tcp IosLinktype = "ios-linktype-vj-comp-tcp"

    // Online Diags
    IosLinktype_ios_linktype_online_diags IosLinktype = "ios-linktype-online-diags"

    // VRF Aware Ipsec Crypto
    IosLinktype_ios_linktype_vrf_ipsec IosLinktype = "ios-linktype-vrf-ipsec"

    // Link Broadcom
    IosLinktype_ios_linktype_broadcom IosLinktype = "ios-linktype-broadcom"

    // HyperText Transfer Protocol (HTTP)
    IosLinktype_ios_linktype_http IosLinktype = "ios-linktype-http"

    // Link Fault Propagation
    IosLinktype_ios_linktype_lfp IosLinktype = "ios-linktype-lfp"

    // Dual RPR Interconnect Protocol
    IosLinktype_ios_linktype_drprip IosLinktype = "ios-linktype-drprip"

    // Ethernet Connectivity Fault Management
    // (Version D1)
    IosLinktype_ios_linktype_ether_cfm IosLinktype = "ios-linktype-ether-cfm"

    // Ethernet OAM (802.3ah) 
    IosLinktype_ios_linktype_ether_oam IosLinktype = "ios-linktype-ether-oam"

    // Mac Move Update
    IosLinktype_ios_linktype_mmu IosLinktype = "ios-linktype-mmu"

    // Virtual Switch Link Protocol
    IosLinktype_ios_linktype_vslp IosLinktype = "ios-linktype-vslp"

    // Ethernet Local Management Interface
    IosLinktype_ios_linktype_ether_lmi IosLinktype = "ios-linktype-ether-lmi"

    // Link Layer Discovery Protocol(IEEE 802.1AB)
    IosLinktype_ios_linktype_lldp IosLinktype = "ios-linktype-lldp"

    // GARP VLAN Registration Protocol (802.1Q GVRP)
    IosLinktype_ios_linktype_gvrp IosLinktype = "ios-linktype-gvrp"

    // GARP Multicast Registration Protocol (802.1Q GMRP)
    IosLinktype_ios_linktype_gmrp IosLinktype = "ios-linktype-gmrp"

    // Access Node Control Protocol
    IosLinktype_ios_linktype_ancp IosLinktype = "ios-linktype-ancp"

    // Resilient Ethernet Protocol
    IosLinktype_ios_linktype_rep IosLinktype = "ios-linktype-rep"

    // Resilient Ethernet Protocol Hardware Flood Layer
    IosLinktype_ios_linktype_rep_hfl IosLinktype = "ios-linktype-rep-hfl"

    // IPe Bootstrap
    IosLinktype_ios_linktype_ipe IosLinktype = "ios-linktype-ipe"

    // Multiple VLAN Registration Protocol
    // MVRP (802.1ak)
    IosLinktype_ios_linktype_mvrp IosLinktype = "ios-linktype-mvrp"

    // Virtual Switch Dual Active Protocol
    IosLinktype_ios_linktype_vsda IosLinktype = "ios-linktype-vsda"

    // Ethernet Connectivity Fault Management
    // (IEEE Version)
    IosLinktype_ios_linktype_ecfm IosLinktype = "ios-linktype-ecfm"

    // Ethernet Mac tunneling Protocol
    // (802.1ah I-TAG Version)
    IosLinktype_ios_linktype_dot1ah IosLinktype = "ios-linktype-dot1ah"

    // Satellite Discovery And Control protocol
    IosLinktype_ios_linktype_sdp IosLinktype = "ios-linktype-sdp"

    // CDP Forwarding (CDPFWD)
    IosLinktype_ios_linktype_cdpfwd IosLinktype = "ios-linktype-cdpfwd"

    // IEC 60870-5 101 Async Link
    IosLinktype_ios_linktype_t101 IosLinktype = "ios-linktype-t101"

    // RSVP Protocol 
    IosLinktype_ios_linktype_rsvp IosLinktype = "ios-linktype-rsvp"

    // ISIS L2 OTV - All Level 1 ISes
    IosLinktype_ios_linktype_isisl2_otv_all_l1_is IosLinktype = "ios-linktype-isisl2-otv-all-l1-is"

    // ISIS L2 OTV - All Level 2 ISes
    IosLinktype_ios_linktype_isisl2_otv_all_l2_is IosLinktype = "ios-linktype-isisl2-otv-all-l2-is"

    // G.8032
    IosLinktype_ios_linktype_g8032 IosLinktype = "ios-linktype-g8032"

    // Satellite Discovery And Control protocol
    IosLinktype_ios_linktype_sdac IosLinktype = "ios-linktype-sdac"

    // PTP Peer Delay protocol
    IosLinktype_ios_linktype_ptppd IosLinktype = "ios-linktype-ptppd"

    // RAWTCP
    IosLinktype_ios_linktype_rawtcp IosLinktype = "ios-linktype-rawtcp"

    // Autonomic Networking
    IosLinktype_ios_linktype_an_ch IosLinktype = "ios-linktype-an-ch"

    // FEX SDP/SRP Satellite Discovey Protocol
    // Satellite Registration Protocols
    IosLinktype_ios_linktype_fex_sdp IosLinktype = "ios-linktype-fex-sdp"

    // Ethernet Synchronous Messaging Channel (ESCM)
    IosLinktype_ios_linktype_esmc IosLinktype = "ios-linktype-esmc"

    // Client Information Signaling Protocol (CISP)
    IosLinktype_ios_linktype_cisp IosLinktype = "ios-linktype-cisp"

    // punt/inject path Async Packet
    IosLinktype_ios_linktype_async IosLinktype = "ios-linktype-async"

    // Multiple Stream Reservation Protocol
    IosLinktype_ios_linktype_msrp IosLinktype = "ios-linktype-msrp"

    // Online Diags
    IosLinktype_ios_linktype_macsec_post_exp IosLinktype = "ios-linktype-macsec-post-exp"

    // Online Diags
    IosLinktype_ios_linktype_macsec_sectag IosLinktype = "ios-linktype-macsec-sectag"
)

// IosSnpaType represents SNPA Type
type IosSnpaType string

const (
    // illegal address
    IosSnpaType_ios_snpa_type_illegal IosSnpaType = "ios-snpa-type-illegal"

    // 48bit IEEE 802.X address
    IosSnpaType_ios_snpa_type_ieee48 IosSnpaType = "ios-snpa-type-ieee48"

    // 16bit IEEE 802.X address
    IosSnpaType_ios_snpa_type_ieee16 IosSnpaType = "ios-snpa-type-ieee16"

    // Xerox 3MB experimental ether
    IosSnpaType_ios_snpa_type_xerox IosSnpaType = "ios-snpa-type-xerox"

    // CCITT X.121 address
    IosSnpaType_ios_snpa_type_x121 IosSnpaType = "ios-snpa-type-x121"

    // cisco HDLC framing
    IosSnpaType_ios_snpa_type_cisco_hdlc IosSnpaType = "ios-snpa-type-cisco-hdlc"

    // cisco multi-LAPB framing
    IosSnpaType_ios_snpa_type_cisco_mlapb IosSnpaType = "ios-snpa-type-cisco-mlapb"

    // ISO/CCITT LAPB framing
    IosSnpaType_ios_snpa_type_lapb IosSnpaType = "ios-snpa-type-lapb"

    // SMDS w/ 48 bit addresses
    IosSnpaType_ios_snpa_type_smds48 IosSnpaType = "ios-snpa-type-smds48"

    // cisco multi-SDLC framing
    IosSnpaType_ios_snpa_type_cisco_msdlc IosSnpaType = "ios-snpa-type-cisco-msdlc"

    // Frame Relay with 10-bit DLCI
    IosSnpaType_ios_snpa_type_fr10 IosSnpaType = "ios-snpa-type-fr10"

    // CCCI defined Ultranet
    IosSnpaType_ios_snpa_type_ultra IosSnpaType = "ios-snpa-type-ultra"

    // cisco tunnel and EON encoding
    IosSnpaType_ios_snpa_type_cisco_tunnel IosSnpaType = "ios-snpa-type-cisco-tunnel"

    // CLNP tunnel
    IosSnpaType_ios_snpa_type_cisco_ctunnel IosSnpaType = "ios-snpa-type-cisco-ctunnel"

    // rrr tunnel
    IosSnpaType_ios_snpa_type_rrr_tunnel IosSnpaType = "ios-snpa-type-rrr-tunnel"

    // PPP framing
    IosSnpaType_ios_snpa_type_ppp IosSnpaType = "ios-snpa-type-ppp"

    // SMDS 64-bit addressing
    IosSnpaType_ios_snpa_type_smds64 IosSnpaType = "ios-snpa-type-smds64"

    // AIP VC no.
    IosSnpaType_ios_snpa_type_atmvc IosSnpaType = "ios-snpa-type-atmvc"

    // ATM PVC bundle
    IosSnpaType_ios_snpa_type_atm_bundle IosSnpaType = "ios-snpa-type-atm-bundle"

    // ATM SVC bundle
    IosSnpaType_ios_snpa_type_atm_svc_bundle IosSnpaType = "ios-snpa-type-atm-svc-bundle"

    // ATM NSAP address
    IosSnpaType_ios_snpa_type_atmnsap IosSnpaType = "ios-snpa-type-atmnsap"

    // ATM E164 address
    IosSnpaType_ios_snpa_type_atm_e164 IosSnpaType = "ios-snpa-type-atm-e164"

    // ATM User Specified address
    IosSnpaType_ios_snpa_type_atm_userspecified IosSnpaType = "ios-snpa-type-atm-userspecified"

    // SDLC address
    IosSnpaType_ios_snpa_type_sdlc IosSnpaType = "ios-snpa-type-sdlc"

    // X.25 Lci fpr PVC's
    IosSnpaType_ios_snpa_type_x25pvc IosSnpaType = "ios-snpa-type-x25pvc"

    // LAPD framing
    IosSnpaType_ios_snpa_type_lapd IosSnpaType = "ios-snpa-type-lapd"

    // masked ATM NSAP address
    IosSnpaType_ios_snpa_type_masked_atmnsap IosSnpaType = "ios-snpa-type-masked-atmnsap"

    // ATM ESI address
    IosSnpaType_ios_snpa_type_atmesi IosSnpaType = "ios-snpa-type-atmesi"

    // SLIP Framing
    IosSnpaType_ios_snpa_type_slip IosSnpaType = "ios-snpa-type-slip"

    // route descriptor for TR LANE
    IosSnpaType_ios_snpa_type_routedesc IosSnpaType = "ios-snpa-type-routedesc"

    // 48bit SRP address on outer ring
    IosSnpaType_ios_snpa_type_srp_outer IosSnpaType = "ios-snpa-type-srp-outer"

    // 48bit SRP address on inner ring
    IosSnpaType_ios_snpa_type_srp_inner IosSnpaType = "ios-snpa-type-srp-inner"
)

// IosEncapsType represents Encaps Type
type IosEncapsType string

const (
    // undefined -- error
    IosEncapsType_ios_encaps_type_null IosEncapsType = "ios-encaps-type-null"

    // Ethernet - DDN style
    IosEncapsType_ios_encaps_type_arpa IosEncapsType = "ios-encaps-type-arpa"

    // Ethernet
    IosEncapsType_ios_encaps_type_sap IosEncapsType = "ios-encaps-type-sap"

    // 802.2 SNAP types
    IosEncapsType_ios_encaps_type_snap IosEncapsType = "ios-encaps-type-snap"

    // DDN - 1822 (obsolete)
    IosEncapsType_ios_encaps_type_1822 IosEncapsType = "ios-encaps-type-1822"

    // Serial - raw HDLC
    IosEncapsType_ios_encaps_type_hdlc IosEncapsType = "ios-encaps-type-hdlc"

    // Unused Placeholder (3MB)
    IosEncapsType_ios_encaps_type_unused2 IosEncapsType = "ios-encaps-type-unused2"

    // Unused Placeholder (HDH)
    IosEncapsType_ios_encaps_type_unused1 IosEncapsType = "ios-encaps-type-unused1"

    // Serial - LAPB
    IosEncapsType_ios_encaps_type_lapb IosEncapsType = "ios-encaps-type-lapb"

    // Serial - X.25
    IosEncapsType_ios_encaps_type_x25 IosEncapsType = "ios-encaps-type-x25"

    // HUB fiber optic
    IosEncapsType_ios_encaps_type_hub IosEncapsType = "ios-encaps-type-hub"

    // Novell style XNS on Ethernet
    IosEncapsType_ios_encaps_type_novell_ether IosEncapsType = "ios-encaps-type-novell-ether"

    // Unsupported protocols
    IosEncapsType_ios_encaps_type_unsupported IosEncapsType = "ios-encaps-type-unsupported"

    // 3Com XNS over TR 802.2 0x80
    IosEncapsType_ios_encaps_type_3com_tr IosEncapsType = "ios-encaps-type-3com-tr"

    // Ungermann-Bass XNS over TR SNAP
    IosEncapsType_ios_encaps_type_ub_tr IosEncapsType = "ios-encaps-type-ub-tr"

    // Apollo domain packets
    IosEncapsType_ios_encaps_type_apollo IosEncapsType = "ios-encaps-type-apollo"

    // Serial PPP
    IosEncapsType_ios_encaps_type_ppp IosEncapsType = "ios-encaps-type-ppp"

    // Unused Placeholder (ISO3)
    IosEncapsType_ios_encaps_type_unused4 IosEncapsType = "ios-encaps-type-unused4"

    // ISO1 with Vines demux byte
    IosEncapsType_ios_encaps_type_vines_tr IosEncapsType = "ios-encaps-type-vines-tr"

    // AppleTalk phase 1 on ethernet
    IosEncapsType_ios_encaps_type_ethertalk IosEncapsType = "ios-encaps-type-ethertalk"

    // Frame Relay
    IosEncapsType_ios_encaps_type_frame_relay IosEncapsType = "ios-encaps-type-frame-relay"

    // Switched Multimegabit Data Service (SMDS)
    IosEncapsType_ios_encaps_type_smds IosEncapsType = "ios-encaps-type-smds"

    // MAC level packets
    IosEncapsType_ios_encaps_type_mac IosEncapsType = "ios-encaps-type-mac"

    // Ultranet
    IosEncapsType_ios_encaps_type_ultra_iso2 IosEncapsType = "ios-encaps-type-ultra-iso2"

    // Ultranet-Hello
    IosEncapsType_ios_encaps_type_ultra_iso1 IosEncapsType = "ios-encaps-type-ultra-iso1"

    // Serial - serial tunnelling
    IosEncapsType_ios_encaps_type_stun IosEncapsType = "ios-encaps-type-stun"

    // Packet in bridge encapsulation
    IosEncapsType_ios_encaps_type_bridge IosEncapsType = "ios-encaps-type-bridge"

    // LLC 2
    IosEncapsType_ios_encaps_type_llc2 IosEncapsType = "ios-encaps-type-llc2"

    // Serial - SDLC (primary)
    IosEncapsType_ios_encaps_type_sdlcp IosEncapsType = "ios-encaps-type-sdlcp"

    // Serial - SDLC (secondary)
    IosEncapsType_ios_encaps_type_sdlcs IosEncapsType = "ios-encaps-type-sdlcs"

    // Async SLIP encapsulation
    IosEncapsType_ios_encaps_type_slip IosEncapsType = "ios-encaps-type-slip"

    // Standard Tunnel Interface
    IosEncapsType_ios_encaps_type_tunnel IosEncapsType = "ios-encaps-type-tunnel"

    // Bridge encap on local gen packs
    IosEncapsType_ios_encaps_type_bridge_encaps IosEncapsType = "ios-encaps-type-bridge-encaps"

    // ATM interface encaps
    IosEncapsType_ios_encaps_type_atm IosEncapsType = "ios-encaps-type-atm"

    // ATM DXI implementation
    IosEncapsType_ios_encaps_type_atm_dxi IosEncapsType = "ios-encaps-type-atm-dxi"

    // Frame Relay with IETF encaps
    IosEncapsType_ios_encaps_type_fr_ietf IosEncapsType = "ios-encaps-type-fr-ietf"

    // SMDS DXI implementation
    IosEncapsType_ios_encaps_type_smds_dxi IosEncapsType = "ios-encaps-type-smds-dxi"

    // atm-dxi with IETF encaps
    IosEncapsType_ios_encaps_type_atm_dxi_ietf IosEncapsType = "ios-encaps-type-atm-dxi-ietf"

    // IBM Channel
    IosEncapsType_ios_encaps_type_channel IosEncapsType = "ios-encaps-type-channel"

    // CLSI compliant SDLC
    IosEncapsType_ios_encaps_type_sdlc IosEncapsType = "ios-encaps-type-sdlc"

    // 802.10 Secure Data Exchange
    IosEncapsType_ios_encaps_type_sde IosEncapsType = "ios-encaps-type-sde"

    // block serial tunnel
    IosEncapsType_ios_encaps_type_bstun IosEncapsType = "ios-encaps-type-bstun"

    // Dialer encapsulation
    IosEncapsType_ios_encaps_type_dialer IosEncapsType = "ios-encaps-type-dialer"

    // Novell style XNS on FDDI
    IosEncapsType_ios_encaps_type_novell_fddi IosEncapsType = "ios-encaps-type-novell-fddi"

    // V120 ISDN->ASYNC encaps
    IosEncapsType_ios_encaps_type_v120 IosEncapsType = "ios-encaps-type-v120"

    // Inter Switch Link - vLANs on FEIP
    IosEncapsType_ios_encaps_type_isl IosEncapsType = "ios-encaps-type-isl"

    // loopback interface
    IosEncapsType_ios_encaps_type_loop IosEncapsType = "ios-encaps-type-loop"

    // IBM Channel internal LAN
    IosEncapsType_ios_encaps_type_channel_ilan IosEncapsType = "ios-encaps-type-channel-ilan"

    // Autodetected for Serials
    IosEncapsType_ios_encaps_type_serial_autodetect IosEncapsType = "ios-encaps-type-serial-autodetect"

    // Combinet proprietary protocol
    IosEncapsType_ios_encaps_type_cpp IosEncapsType = "ios-encaps-type-cpp"

    // NCIA DLC
    IosEncapsType_ios_encaps_type_ndlc IosEncapsType = "ios-encaps-type-ndlc"

    // ISDN Q.921
    IosEncapsType_ios_encaps_type_lapd IosEncapsType = "ios-encaps-type-lapd"

    // StrataCom IPX/Fastpad protocol
    IosEncapsType_ios_encaps_type_ftc_trunk IosEncapsType = "ios-encaps-type-ftc-trunk"

    // ATM T1 Circuit Emulation.
    IosEncapsType_ios_encaps_type_atmces_t1 IosEncapsType = "ios-encaps-type-atmces-t1"

    // ATM E1 Circuit Emulation.
    IosEncapsType_ios_encaps_type_atmces_e1 IosEncapsType = "ios-encaps-type-atmces-e1"

    // Packetized Voice
    IosEncapsType_ios_encaps_type_voice IosEncapsType = "ios-encaps-type-voice"

    // ALC (P1024B) sync protocol
    IosEncapsType_ios_encaps_type_alc IosEncapsType = "ios-encaps-type-alc"

    // UTS (P1024C) sync protocol
    IosEncapsType_ios_encaps_type_uts IosEncapsType = "ios-encaps-type-uts"

    // Token Ring Inter-Switch Link (TRISL)
    IosEncapsType_ios_encaps_type_trisl IosEncapsType = "ios-encaps-type-trisl"

    // Cable - MCNS
    IosEncapsType_ios_encaps_type_mcns IosEncapsType = "ios-encaps-type-mcns"

    // ATM circuit emulation service
    IosEncapsType_ios_encaps_type_atmces IosEncapsType = "ios-encaps-type-atmces"

    // for transparent mode
    IosEncapsType_ios_encaps_type_trans IosEncapsType = "ios-encaps-type-trans"

    // TDM clear channel
    IosEncapsType_ios_encaps_type_clear_channel IosEncapsType = "ios-encaps-type-clear-channel"

    // Tag Controlled ATM interface
    IosEncapsType_ios_encaps_type_tc_atm IosEncapsType = "ios-encaps-type-tc-atm"

    // PPP over Frame Relay
    IosEncapsType_ios_encaps_type_fr_ppp IosEncapsType = "ios-encaps-type-fr-ppp"

    // IEEE 802.1Q
    IosEncapsType_ios_encaps_type_dot1q IosEncapsType = "ios-encaps-type-dot1q"

    // Frame based user network interface
    IosEncapsType_ios_encaps_type_funi IosEncapsType = "ios-encaps-type-funi"

    // LAPB terminal adapter
    IosEncapsType_ios_encaps_type_lapbta IosEncapsType = "ios-encaps-type-lapbta"

    // Cable Modem
    IosEncapsType_ios_encaps_type_docsis IosEncapsType = "ios-encaps-type-docsis"

    // NextPort DSP modem in-band message
    IosEncapsType_ios_encaps_type_np_inband IosEncapsType = "ios-encaps-type-np-inband"

    // SS7 MTP-2 over serial/TDM interface
    IosEncapsType_ios_encaps_type_ss7_mtp2 IosEncapsType = "ios-encaps-type-ss7-mtp2"

    // General Packet Radio Service (GPRS)
    IosEncapsType_ios_encaps_type_gtp IosEncapsType = "ios-encaps-type-gtp"

    // ISL SW encap pkt sent on swch prt
    IosEncapsType_ios_encaps_type_isl_switchport IosEncapsType = "ios-encaps-type-isl-switchport"

    // Spatial Reuse Protocol
    IosEncapsType_ios_encaps_type_srp IosEncapsType = "ios-encaps-type-srp"

    // PPP hdr w/Ethertype protocol type
    IosEncapsType_ios_encaps_type_ppp_etype IosEncapsType = "ios-encaps-type-ppp-etype"

    // PPP over ATM
    IosEncapsType_ios_encaps_type_atm_ppp IosEncapsType = "ios-encaps-type-atm-ppp"

    // Multilink Frame Relay
    IosEncapsType_ios_encaps_type_mfr IosEncapsType = "ios-encaps-type-mfr"

    // SRP version 2
    IosEncapsType_ios_encaps_type_srp2 IosEncapsType = "ios-encaps-type-srp2"

    // Frame Relay extended addressing
    IosEncapsType_ios_encaps_type_fr_extended IosEncapsType = "ios-encaps-type-fr-extended"

    // PPP Over Ethernet
    IosEncapsType_ios_encaps_type_pppoe IosEncapsType = "ios-encaps-type-pppoe"

    // OIF UNI SDCC per OIF2000.125
    IosEncapsType_ios_encaps_type_ouni_sdcc IosEncapsType = "ios-encaps-type-ouni-sdcc"

    // encap type for multiservice i/f
    IosEncapsType_ios_encaps_type_multiservice IosEncapsType = "ios-encaps-type-multiservice"

    // Container
    IosEncapsType_ios_encaps_type_container IosEncapsType = "ios-encaps-type-container"

    // Dot1ad ether type
    IosEncapsType_ios_encaps_type_dot1ad IosEncapsType = "ios-encaps-type-dot1ad"

    // Circuit Emulation Type
    IosEncapsType_ios_encaps_type_cem IosEncapsType = "ios-encaps-type-cem"

    // IEEE 802.1ah Ether type
    IosEncapsType_ios_encaps_type_dot1ah IosEncapsType = "ios-encaps-type-dot1ah"

    // PTP type
    IosEncapsType_ios_encaps_type_ptp IosEncapsType = "ios-encaps-type-ptp"

    // For SSLVPN
    IosEncapsType_ios_encaps_type_ssl IosEncapsType = "ios-encaps-type-ssl"

    // Locator/ID Separation Protocol (LISP)
    IosEncapsType_ios_encaps_type_lisp IosEncapsType = "ios-encaps-type-lisp"

    // DSP-based SPA
    IosEncapsType_ios_encaps_type_dsp_spa IosEncapsType = "ios-encaps-type-dsp-spa"

    // IEC-60870-5-101 for Async Serial interface
    IosEncapsType_ios_encaps_type_t101 IosEncapsType = "ios-encaps-type-t101"

    // Crypto
    IosEncapsType_ios_encaps_type_crypto IosEncapsType = "ios-encaps-type-crypto"

    // Raw socket TCP for Async line interface
    IosEncapsType_ios_encaps_type_rawtcp IosEncapsType = "ios-encaps-type-rawtcp"

    // Async line raw data
    IosEncapsType_ios_encaps_type_async IosEncapsType = "ios-encaps-type-async"

    // Cisco Meta Data (CMD) EtherType=0x8909 encap
    IosEncapsType_ios_encaps_type_cmd IosEncapsType = "ios-encaps-type-cmd"

    // SCADA for Async Serial Interface
    IosEncapsType_ios_encaps_type_scada IosEncapsType = "ios-encaps-type-scada"

    // Raw socket UDP for Async line interface
    IosEncapsType_ios_encaps_type_rawudp IosEncapsType = "ios-encaps-type-rawudp"

    // Line relay for Async line interface
    IosEncapsType_ios_encaps_type_relay_line IosEncapsType = "ios-encaps-type-relay-line"

    // Raw socket over MPLS for Async line interface
    // Market want to call it trans. It's a little conflict
    // TRANSPARENT
    IosEncapsType_ios_encaps_type_rawmpls IosEncapsType = "ios-encaps-type-rawmpls"

    // if NID added after ARPA we add a new type
    IosEncapsType_ios_encaps_type_arpa_nid IosEncapsType = "ios-encaps-type-arpa-nid"
)

