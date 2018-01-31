// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-udp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   udp: IP UDP Operational Data
//   udp-connection: udp connection
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_udp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_udp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-udp-oper udp}", reflect.TypeOf(Udp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-udp-oper:udp", reflect.TypeOf(Udp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-udp-oper udp-connection}", reflect.TypeOf(UdpConnection{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-udp-oper:udp-connection", reflect.TypeOf(UdpConnection{}))
}

// LptsPcbQuery represents Lpts pcb query
type LptsPcbQuery string

const (
    // No filter
    LptsPcbQuery_all LptsPcbQuery = "all"

    // Static policy filter
    LptsPcbQuery_static_policy LptsPcbQuery = "static-policy"

    // Interface filter
    LptsPcbQuery_interface_ LptsPcbQuery = "interface"

    // Packet type filter
    LptsPcbQuery_packet LptsPcbQuery = "packet"
)

// MessageTypeIcmpv6 represents LPTS ICMPv6 message types
type MessageTypeIcmpv6 string

const (
    // ICMPv6 Packet type: Destination unreachable
    MessageTypeIcmpv6_destination_unreachable MessageTypeIcmpv6 = "destination-unreachable"

    // ICMPv6 Packet type: packet too big
    MessageTypeIcmpv6_packet_too_big MessageTypeIcmpv6 = "packet-too-big"

    // ICMPv6 Packet type: Time exceeded
    MessageTypeIcmpv6_time_exceeded MessageTypeIcmpv6 = "time-exceeded"

    // ICMPv6 Packet type: Parameter problem
    MessageTypeIcmpv6_parameter_problem MessageTypeIcmpv6 = "parameter-problem"

    // ICMPv6 Packet type: Echo request
    MessageTypeIcmpv6_echo_request MessageTypeIcmpv6 = "echo-request"

    // ICMPv6 Packet type: Echo reply
    MessageTypeIcmpv6_echo_reply MessageTypeIcmpv6 = "echo-reply"

    // ICMPv6 Packet type: Multicast listener query
    MessageTypeIcmpv6_multicast_listener_query MessageTypeIcmpv6 = "multicast-listener-query"

    // ICMPv6 Packet type: Multicast listener report
    MessageTypeIcmpv6_multicast_listener_report MessageTypeIcmpv6 = "multicast-listener-report"

    // ICMPv6 Packet type: Multicast listener done
    MessageTypeIcmpv6_multicast_listener_done MessageTypeIcmpv6 = "multicast-listener-done"

    // ICMPv6 Packet type: Router solicitation
    MessageTypeIcmpv6_router_solicitation MessageTypeIcmpv6 = "router-solicitation"

    // ICMPv6 Packet type: Router advertisement
    MessageTypeIcmpv6_router_advertisement MessageTypeIcmpv6 = "router-advertisement"

    // ICMPv6 Packet type: Neighbor solicitation
    MessageTypeIcmpv6_neighbor_solicitation MessageTypeIcmpv6 = "neighbor-solicitation"

    // ICMPv6 Packet type: Neighbor advertisement
    MessageTypeIcmpv6_neighbor_advertisement MessageTypeIcmpv6 = "neighbor-advertisement"

    // ICMPv6 Packet type: Redirect message
    MessageTypeIcmpv6_redirect_message MessageTypeIcmpv6 = "redirect-message"

    // ICMPv6 Packet type: Router renumbering
    MessageTypeIcmpv6_router_renumbering MessageTypeIcmpv6 = "router-renumbering"

    // ICMPv6 Packet type: Node information query
    MessageTypeIcmpv6_node_information_query MessageTypeIcmpv6 = "node-information-query"

    // ICMPv6 Packet type: Node information reply
    MessageTypeIcmpv6_node_information_reply MessageTypeIcmpv6 = "node-information-reply"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // solicitation message
    MessageTypeIcmpv6_inverse_neighbor_discovery_solicitaion MessageTypeIcmpv6 = "inverse-neighbor-discovery-solicitaion"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // advertisement message
    MessageTypeIcmpv6_inverse_neighbor_discover_advertisement MessageTypeIcmpv6 = "inverse-neighbor-discover-advertisement"

    // ICMPv6 Packet type: Version 2 multicast
    // listener report
    MessageTypeIcmpv6_v2_multicast_listener_report MessageTypeIcmpv6 = "v2-multicast-listener-report"

    // ICMPv6 Packet type: Home agent address
    // discovery request message
    MessageTypeIcmpv6_home_agent_address_discovery_request MessageTypeIcmpv6 = "home-agent-address-discovery-request"

    // ICMPv6 Packet type: Home agent address
    // discovery reply message
    MessageTypeIcmpv6_home_agent_address_discovery_reply MessageTypeIcmpv6 = "home-agent-address-discovery-reply"

    // ICMPv6 Packet type: Mobile prefix solicitation
    MessageTypeIcmpv6_mobile_prefix_solicitation MessageTypeIcmpv6 = "mobile-prefix-solicitation"

    // ICMPv6 Packet type: Mobile prefix advertisement
    MessageTypeIcmpv6_mobile_prefix_advertisement MessageTypeIcmpv6 = "mobile-prefix-advertisement"

    // ICMPv6 Packet type: Certification path
    // solicitation message
    MessageTypeIcmpv6_certification_path_solicitation_message MessageTypeIcmpv6 = "certification-path-solicitation-message"

    // ICMPv6 Packet type: Certification path
    // advertisement message
    MessageTypeIcmpv6_certification_path_advertisement_message MessageTypeIcmpv6 = "certification-path-advertisement-message"

    // ICMPv6 Packet type: ICMP messages utilized by
    // experimental mobility  protocols such as
    // seamoby
    MessageTypeIcmpv6_experimental_mobility_protocols MessageTypeIcmpv6 = "experimental-mobility-protocols"

    // ICMPv6 Packet type: Multicast router
    // advertisement
    MessageTypeIcmpv6_multicast_router_advertisement MessageTypeIcmpv6 = "multicast-router-advertisement"

    // ICMPv6 Packet type: Multicast router
    // solicitation
    MessageTypeIcmpv6_multicast_router_solicitation MessageTypeIcmpv6 = "multicast-router-solicitation"

    // ICMPv6 Packet type: Multicast router
    // termination
    MessageTypeIcmpv6_multicast_router_termination MessageTypeIcmpv6 = "multicast-router-termination"

    // ICMPv6 Packet type: FMIPv6 messages
    MessageTypeIcmpv6_fmipv6_messages MessageTypeIcmpv6 = "fmipv6-messages"
)

// MessageTypeIcmp represents LPTS ICMP message types
type MessageTypeIcmp string

const (
    // ICMP Packet type: Echo reply
    MessageTypeIcmp_echo_reply MessageTypeIcmp = "echo-reply"

    // ICMP Packet type: Destination unreachable
    MessageTypeIcmp_destination_unreachable MessageTypeIcmp = "destination-unreachable"

    // ICMP Packet type: Source quench
    MessageTypeIcmp_source_quench MessageTypeIcmp = "source-quench"

    // ICMP Packet type: Redirect
    MessageTypeIcmp_redirect MessageTypeIcmp = "redirect"

    // ICMP Packet type: Alternate host address
    MessageTypeIcmp_alternate_host_address MessageTypeIcmp = "alternate-host-address"

    // ICMP Packet type: Echo
    MessageTypeIcmp_echo MessageTypeIcmp = "echo"

    // ICMP Packet type: Router advertisement
    MessageTypeIcmp_router_advertisement MessageTypeIcmp = "router-advertisement"

    // ICMP Packet type: Router selection
    MessageTypeIcmp_router_selection MessageTypeIcmp = "router-selection"

    // ICMP Packet type: Time exceeded
    MessageTypeIcmp_time_exceeded MessageTypeIcmp = "time-exceeded"

    // ICMP Packet type: Parameter problem
    MessageTypeIcmp_parameter_problem MessageTypeIcmp = "parameter-problem"

    // ICMP Packet type: Time stamp
    MessageTypeIcmp_time_stamp MessageTypeIcmp = "time-stamp"

    // ICMP Packet type: Time stamp reply
    MessageTypeIcmp_time_stamp_reply MessageTypeIcmp = "time-stamp-reply"

    // ICMP Packet type: Information request
    MessageTypeIcmp_information_request MessageTypeIcmp = "information-request"

    // ICMP Packet type: Information reply
    MessageTypeIcmp_information_reply MessageTypeIcmp = "information-reply"

    // ICMP Packet type: Address mask request
    MessageTypeIcmp_address_mask_request MessageTypeIcmp = "address-mask-request"

    // ICMP Packet type: Address mask reply
    MessageTypeIcmp_address_mask_reply MessageTypeIcmp = "address-mask-reply"

    // ICMP Packet type: Trace route
    MessageTypeIcmp_trace_route MessageTypeIcmp = "trace-route"

    // ICMP Packet type: Datagram Conversion error
    MessageTypeIcmp_datagram_conversion_error MessageTypeIcmp = "datagram-conversion-error"

    // ICMP Packet type: Mobile host redirect
    MessageTypeIcmp_mobile_host_redirect MessageTypeIcmp = "mobile-host-redirect"

    // ICMP Packet type: IPv6 where-are-you
    MessageTypeIcmp_where_are_you MessageTypeIcmp = "where-are-you"

    // ICMP Packet type: IPv6 i-am-here
    MessageTypeIcmp_iam_here MessageTypeIcmp = "iam-here"

    // ICMP Packet type: Mobile registration request
    MessageTypeIcmp_mobile_registration_request MessageTypeIcmp = "mobile-registration-request"

    // ICMP Packet type: Mobile registration reply
    MessageTypeIcmp_mobile_registration_reply MessageTypeIcmp = "mobile-registration-reply"

    // ICMP Packet type: Domain name request
    MessageTypeIcmp_domain_name_request MessageTypeIcmp = "domain-name-request"
)

// MessageTypeIgmp represents LPTS IGMP message types
type MessageTypeIgmp string

const (
    // IGMP Packet type: Membership query
    MessageTypeIgmp_membership_query MessageTypeIgmp = "membership-query"

    // IGMP Packet type: V1 membership report
    MessageTypeIgmp_v1_membership_report MessageTypeIgmp = "v1-membership-report"

    // IGMP Packet type: DVMRP
    MessageTypeIgmp_dvmrp MessageTypeIgmp = "dvmrp"

    // IGMP Packet type: PIM version 1
    MessageTypeIgmp_pi_mv1 MessageTypeIgmp = "pi-mv1"

    // IGMP Packet type: Cisco Trace Messages
    MessageTypeIgmp_cisco_trace_messages MessageTypeIgmp = "cisco-trace-messages"

    // IGMP Packet type: V2 membership report
    MessageTypeIgmp_v2_membership_report MessageTypeIgmp = "v2-membership-report"

    // IGMP Packet type: V2 leave group
    MessageTypeIgmp_v2_leave_group MessageTypeIgmp = "v2-leave-group"

    // IGMP Packet type: Multicast traceroute response
    MessageTypeIgmp_multicast_traceroute_response MessageTypeIgmp = "multicast-traceroute-response"

    // IGMP Packet type: MulticastTraceroute
    MessageTypeIgmp_multicast_traceroute MessageTypeIgmp = "multicast-traceroute"

    // IGMP Packet type: V3 membership report
    MessageTypeIgmp_v3_membership_report MessageTypeIgmp = "v3-membership-report"

    // IGMP Packet type: Multicast router
    // advertisement
    MessageTypeIgmp_multicast_router_advertisement MessageTypeIgmp = "multicast-router-advertisement"

    // IGMP Packet type: Multicast router solicitation
    MessageTypeIgmp_multicast_router_solicitation MessageTypeIgmp = "multicast-router-solicitation"

    // IGMP Packet type: Multicast router termination
    MessageTypeIgmp_multicast_router_termination MessageTypeIgmp = "multicast-router-termination"
)

// MessageTypeIgmp represents LPTS IGMP message types
type MessageTypeIgmp string

const (
    // IGMP Packet type: Membership query
    MessageTypeIgmp_membership_query MessageTypeIgmp = "membership-query"

    // IGMP Packet type: V1 membership report
    MessageTypeIgmp_v1_membership_report MessageTypeIgmp = "v1-membership-report"

    // IGMP Packet type: DVMRP
    MessageTypeIgmp_dvmrp MessageTypeIgmp = "dvmrp"

    // IGMP Packet type: PIM version 1
    MessageTypeIgmp_pi_mv1 MessageTypeIgmp = "pi-mv1"

    // IGMP Packet type: Cisco Trace Messages
    MessageTypeIgmp_cisco_trace_messages MessageTypeIgmp = "cisco-trace-messages"

    // IGMP Packet type: V2 membership report
    MessageTypeIgmp_v2_membership_report MessageTypeIgmp = "v2-membership-report"

    // IGMP Packet type: V2 leave group
    MessageTypeIgmp_v2_leave_group MessageTypeIgmp = "v2-leave-group"

    // IGMP Packet type: Multicast traceroute response
    MessageTypeIgmp_multicast_traceroute_response MessageTypeIgmp = "multicast-traceroute-response"

    // IGMP Packet type: MulticastTraceroute
    MessageTypeIgmp_multicast_traceroute MessageTypeIgmp = "multicast-traceroute"

    // IGMP Packet type: V3 membership report
    MessageTypeIgmp_v3_membership_report MessageTypeIgmp = "v3-membership-report"

    // IGMP Packet type: Multicast router
    // advertisement
    MessageTypeIgmp_multicast_router_advertisement MessageTypeIgmp = "multicast-router-advertisement"

    // IGMP Packet type: Multicast router solicitation
    MessageTypeIgmp_multicast_router_solicitation MessageTypeIgmp = "multicast-router-solicitation"

    // IGMP Packet type: Multicast router termination
    MessageTypeIgmp_multicast_router_termination MessageTypeIgmp = "multicast-router-termination"
)

// Packet represents Packet type
type Packet string

const (
    // ICMP packet type
    Packet_icmp Packet = "icmp"

    // ICMPv6 packet type
    Packet_icm_pv6 Packet = "icm-pv6"

    // IGMP packet type
    Packet_igmp Packet = "igmp"

    // Packet type unknown
    Packet_unknown Packet = "unknown"
)

// MessageTypeIcmp represents LPTS ICMP message types
type MessageTypeIcmp string

const (
    // ICMP Packet type: Echo reply
    MessageTypeIcmp_echo_reply MessageTypeIcmp = "echo-reply"

    // ICMP Packet type: Destination unreachable
    MessageTypeIcmp_destination_unreachable MessageTypeIcmp = "destination-unreachable"

    // ICMP Packet type: Source quench
    MessageTypeIcmp_source_quench MessageTypeIcmp = "source-quench"

    // ICMP Packet type: Redirect
    MessageTypeIcmp_redirect MessageTypeIcmp = "redirect"

    // ICMP Packet type: Alternate host address
    MessageTypeIcmp_alternate_host_address MessageTypeIcmp = "alternate-host-address"

    // ICMP Packet type: Echo
    MessageTypeIcmp_echo MessageTypeIcmp = "echo"

    // ICMP Packet type: Router advertisement
    MessageTypeIcmp_router_advertisement MessageTypeIcmp = "router-advertisement"

    // ICMP Packet type: Router selection
    MessageTypeIcmp_router_selection MessageTypeIcmp = "router-selection"

    // ICMP Packet type: Time exceeded
    MessageTypeIcmp_time_exceeded MessageTypeIcmp = "time-exceeded"

    // ICMP Packet type: Parameter problem
    MessageTypeIcmp_parameter_problem MessageTypeIcmp = "parameter-problem"

    // ICMP Packet type: Time stamp
    MessageTypeIcmp_time_stamp MessageTypeIcmp = "time-stamp"

    // ICMP Packet type: Time stamp reply
    MessageTypeIcmp_time_stamp_reply MessageTypeIcmp = "time-stamp-reply"

    // ICMP Packet type: Information request
    MessageTypeIcmp_information_request MessageTypeIcmp = "information-request"

    // ICMP Packet type: Information reply
    MessageTypeIcmp_information_reply MessageTypeIcmp = "information-reply"

    // ICMP Packet type: Address mask request
    MessageTypeIcmp_address_mask_request MessageTypeIcmp = "address-mask-request"

    // ICMP Packet type: Address mask reply
    MessageTypeIcmp_address_mask_reply MessageTypeIcmp = "address-mask-reply"

    // ICMP Packet type: Trace route
    MessageTypeIcmp_trace_route MessageTypeIcmp = "trace-route"

    // ICMP Packet type: Datagram Conversion error
    MessageTypeIcmp_datagram_conversion_error MessageTypeIcmp = "datagram-conversion-error"

    // ICMP Packet type: Mobile host redirect
    MessageTypeIcmp_mobile_host_redirect MessageTypeIcmp = "mobile-host-redirect"

    // ICMP Packet type: IPv6 where-are-you
    MessageTypeIcmp_where_are_you MessageTypeIcmp = "where-are-you"

    // ICMP Packet type: IPv6 i-am-here
    MessageTypeIcmp_iam_here MessageTypeIcmp = "iam-here"

    // ICMP Packet type: Mobile registration request
    MessageTypeIcmp_mobile_registration_request MessageTypeIcmp = "mobile-registration-request"

    // ICMP Packet type: Mobile registration reply
    MessageTypeIcmp_mobile_registration_reply MessageTypeIcmp = "mobile-registration-reply"

    // ICMP Packet type: Domain name request
    MessageTypeIcmp_domain_name_request MessageTypeIcmp = "domain-name-request"
)

// AddrFamily represents Address Family Types
type AddrFamily string

const (
    // Unspecified
    AddrFamily_unspecified AddrFamily = "unspecified"

    // Local to host (pipes, portals)
    AddrFamily_local AddrFamily = "local"

    // Internetwork: UDP, TCP, etc.
    AddrFamily_inet AddrFamily = "inet"

    // arpanet imp addresses
    AddrFamily_implink AddrFamily = "implink"

    // Pup protocols: e.g. BSP
    AddrFamily_pup AddrFamily = "pup"

    // mit CHAOS protocols
    AddrFamily_chaos AddrFamily = "chaos"

    // XEROX NS protocols
    AddrFamily_ns AddrFamily = "ns"

    // ISO protocols
    AddrFamily_iso AddrFamily = "iso"

    // European computer manufacturers
    AddrFamily_ecma AddrFamily = "ecma"

    // Datakit protocols
    AddrFamily_data_kit AddrFamily = "data-kit"

    // CCITT protocols, X.25 etc
    AddrFamily_ccitt AddrFamily = "ccitt"

    // IBM SNA
    AddrFamily_sna AddrFamily = "sna"

    // DECnet
    AddrFamily_de_cnet AddrFamily = "de-cnet"

    // DEC Direct data link interface
    AddrFamily_dli AddrFamily = "dli"

    // LAT
    AddrFamily_lat AddrFamily = "lat"

    // NSC Hyperchannel
    AddrFamily_hylink AddrFamily = "hylink"

    // Apple Talk
    AddrFamily_appletalk AddrFamily = "appletalk"

    // Internal Routing Protocol
    AddrFamily_route AddrFamily = "route"

    // Link layer interface
    AddrFamily_link AddrFamily = "link"

    // eXpress Transfer Protocol (no AF)
    AddrFamily_pseudo_xtp AddrFamily = "pseudo-xtp"

    // Connection-oriented IP, aka ST II
    AddrFamily_coip AddrFamily = "coip"

    // Computer Network Technology
    AddrFamily_cnt AddrFamily = "cnt"

    // Help Identify RTIP packets
    AddrFamily_pseudo_rtip AddrFamily = "pseudo-rtip"

    // Novell Internet Protocol
    AddrFamily_ipx AddrFamily = "ipx"

    // Simple Internet Protocol
    AddrFamily_sip AddrFamily = "sip"

    // Help Identify PIP packets
    AddrFamily_pseudo_pip AddrFamily = "pseudo-pip"

    // IP version 6
    AddrFamily_inet6 AddrFamily = "inet6"

    // 802.2 SNAP sockets
    AddrFamily_snap AddrFamily = "snap"

    // SAP_CLNS + nlpid encaps
    AddrFamily_clnl AddrFamily = "clnl"

    // cisco HDLC on serial
    AddrFamily_chdlc AddrFamily = "chdlc"

    // PPP sockets
    AddrFamily_ppp AddrFamily = "ppp"

    // Host-based CAS signaling
    AddrFamily_host_cas AddrFamily = "host-cas"

    // DSP messaging
    AddrFamily_dsp AddrFamily = "dsp"

    // SAP Sockets
    AddrFamily_sap AddrFamily = "sap"

    // ATM Sockets
    AddrFamily_atm AddrFamily = "atm"

    // Frame Relay sockets
    AddrFamily_fr AddrFamily = "fr"

    // Voice Media Stream Sockets
    AddrFamily_mso AddrFamily = "mso"

    // ISDN D Channel Sockets
    AddrFamily_dchan AddrFamily = "dchan"

    // Trunk Framer media IF Sockets
    AddrFamily_cas AddrFamily = "cas"

    // Network Address Translation Sockets
    AddrFamily_nat AddrFamily = "nat"

    // Generic Ethernet Sockets
    AddrFamily_ether AddrFamily = "ether"

    // Spatial Reuse Protocol Sockets
    AddrFamily_srp AddrFamily = "srp"
)

// UdpAddressFamily represents Address Family Type
type UdpAddressFamily string

const (
    // IPv4
    UdpAddressFamily_ipv4 UdpAddressFamily = "ipv4"

    // IPv6
    UdpAddressFamily_ipv6 UdpAddressFamily = "ipv6"
)

// MessageTypeIcmpv6 represents LPTS ICMPv6 message types
type MessageTypeIcmpv6 string

const (
    // ICMPv6 Packet type: Destination unreachable
    MessageTypeIcmpv6_destination_unreachable MessageTypeIcmpv6 = "destination-unreachable"

    // ICMPv6 Packet type: packet too big
    MessageTypeIcmpv6_packet_too_big MessageTypeIcmpv6 = "packet-too-big"

    // ICMPv6 Packet type: Time exceeded
    MessageTypeIcmpv6_time_exceeded MessageTypeIcmpv6 = "time-exceeded"

    // ICMPv6 Packet type: Parameter problem
    MessageTypeIcmpv6_parameter_problem MessageTypeIcmpv6 = "parameter-problem"

    // ICMPv6 Packet type: Echo request
    MessageTypeIcmpv6_echo_request MessageTypeIcmpv6 = "echo-request"

    // ICMPv6 Packet type: Echo reply
    MessageTypeIcmpv6_echo_reply MessageTypeIcmpv6 = "echo-reply"

    // ICMPv6 Packet type: Multicast listener query
    MessageTypeIcmpv6_multicast_listener_query MessageTypeIcmpv6 = "multicast-listener-query"

    // ICMPv6 Packet type: Multicast listener report
    MessageTypeIcmpv6_multicast_listener_report MessageTypeIcmpv6 = "multicast-listener-report"

    // ICMPv6 Packet type: Multicast listener done
    MessageTypeIcmpv6_multicast_listener_done MessageTypeIcmpv6 = "multicast-listener-done"

    // ICMPv6 Packet type: Router solicitation
    MessageTypeIcmpv6_router_solicitation MessageTypeIcmpv6 = "router-solicitation"

    // ICMPv6 Packet type: Router advertisement
    MessageTypeIcmpv6_router_advertisement MessageTypeIcmpv6 = "router-advertisement"

    // ICMPv6 Packet type: Neighbor solicitation
    MessageTypeIcmpv6_neighbor_solicitation MessageTypeIcmpv6 = "neighbor-solicitation"

    // ICMPv6 Packet type: Neighbor advertisement
    MessageTypeIcmpv6_neighbor_advertisement MessageTypeIcmpv6 = "neighbor-advertisement"

    // ICMPv6 Packet type: Redirect message
    MessageTypeIcmpv6_redirect_message MessageTypeIcmpv6 = "redirect-message"

    // ICMPv6 Packet type: Router renumbering
    MessageTypeIcmpv6_router_renumbering MessageTypeIcmpv6 = "router-renumbering"

    // ICMPv6 Packet type: Node information query
    MessageTypeIcmpv6_node_information_query MessageTypeIcmpv6 = "node-information-query"

    // ICMPv6 Packet type: Node information reply
    MessageTypeIcmpv6_node_information_reply MessageTypeIcmpv6 = "node-information-reply"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // solicitation message
    MessageTypeIcmpv6_inverse_neighbor_discovery_solicitaion MessageTypeIcmpv6 = "inverse-neighbor-discovery-solicitaion"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // advertisement message
    MessageTypeIcmpv6_inverse_neighbor_discover_advertisement MessageTypeIcmpv6 = "inverse-neighbor-discover-advertisement"

    // ICMPv6 Packet type: Version 2 multicast
    // listener report
    MessageTypeIcmpv6_v2_multicast_listener_report MessageTypeIcmpv6 = "v2-multicast-listener-report"

    // ICMPv6 Packet type: Home agent address
    // discovery request message
    MessageTypeIcmpv6_home_agent_address_discovery_request MessageTypeIcmpv6 = "home-agent-address-discovery-request"

    // ICMPv6 Packet type: Home agent address
    // discovery reply message
    MessageTypeIcmpv6_home_agent_address_discovery_reply MessageTypeIcmpv6 = "home-agent-address-discovery-reply"

    // ICMPv6 Packet type: Mobile prefix solicitation
    MessageTypeIcmpv6_mobile_prefix_solicitation MessageTypeIcmpv6 = "mobile-prefix-solicitation"

    // ICMPv6 Packet type: Mobile prefix advertisement
    MessageTypeIcmpv6_mobile_prefix_advertisement MessageTypeIcmpv6 = "mobile-prefix-advertisement"

    // ICMPv6 Packet type: Certification path
    // solicitation message
    MessageTypeIcmpv6_certification_path_solicitation_message MessageTypeIcmpv6 = "certification-path-solicitation-message"

    // ICMPv6 Packet type: Certification path
    // advertisement message
    MessageTypeIcmpv6_certification_path_advertisement_message MessageTypeIcmpv6 = "certification-path-advertisement-message"

    // ICMPv6 Packet type: ICMP messages utilized by
    // experimental mobility  protocols such as
    // seamoby
    MessageTypeIcmpv6_experimental_mobility_protocols MessageTypeIcmpv6 = "experimental-mobility-protocols"

    // ICMPv6 Packet type: Multicast router
    // advertisement
    MessageTypeIcmpv6_multicast_router_advertisement MessageTypeIcmpv6 = "multicast-router-advertisement"

    // ICMPv6 Packet type: Multicast router
    // solicitation
    MessageTypeIcmpv6_multicast_router_solicitation MessageTypeIcmpv6 = "multicast-router-solicitation"

    // ICMPv6 Packet type: Multicast router
    // termination
    MessageTypeIcmpv6_multicast_router_termination MessageTypeIcmpv6 = "multicast-router-termination"

    // ICMPv6 Packet type: FMIPv6 messages
    MessageTypeIcmpv6_fmipv6_messages MessageTypeIcmpv6 = "fmipv6-messages"
)

// Udp
// IP UDP Operational Data
type Udp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific UDP operational data.
    Nodes Udp_Nodes
}

func (udp *Udp) GetFilter() yfilter.YFilter { return udp.YFilter }

func (udp *Udp) SetFilter(yf yfilter.YFilter) { udp.YFilter = yf }

func (udp *Udp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (udp *Udp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-udp-oper:udp"
}

func (udp *Udp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &udp.Nodes
    }
    return nil
}

func (udp *Udp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &udp.Nodes
    return children
}

func (udp *Udp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udp *Udp) GetBundleName() string { return "cisco_ios_xr" }

func (udp *Udp) GetYangName() string { return "udp" }

func (udp *Udp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udp *Udp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udp *Udp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udp *Udp) SetParent(parent types.Entity) { udp.parent = parent }

func (udp *Udp) GetParent() types.Entity { return udp.parent }

func (udp *Udp) GetParentYangName() string { return "Cisco-IOS-XR-ip-udp-oper" }

// Udp_Nodes
// Node-specific UDP operational data
type Udp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP operational data for a particular node. The type is slice of
    // Udp_Nodes_Node.
    Node []Udp_Nodes_Node
}

func (nodes *Udp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Udp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Udp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Udp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Udp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Udp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Udp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Udp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Udp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Udp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Udp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Udp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Udp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Udp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Udp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Udp_Nodes) GetParentYangName() string { return "udp" }

// Udp_Nodes_Node
// UDP operational data for a particular node
type Udp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Statistical UDP operational data for a node.
    Statistics Udp_Nodes_Node_Statistics
}

func (node *Udp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Udp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Udp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *Udp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Udp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *Udp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    return children
}

func (node *Udp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Udp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Udp_Nodes_Node) GetYangName() string { return "node" }

func (node *Udp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Udp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Udp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Udp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Udp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Udp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Udp_Nodes_Node_Statistics
// Statistical UDP operational data for a node
type Udp_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP Traffic statistics for IPv4.
    Ipv4Traffic Udp_Nodes_Node_Statistics_Ipv4Traffic

    // UDP Traffic statistics for IPv6.
    Ipv6Traffic Udp_Nodes_Node_Statistics_Ipv6Traffic
}

func (statistics *Udp_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Udp_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Udp_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-traffic" { return "Ipv4Traffic" }
    if yname == "ipv6-traffic" { return "Ipv6Traffic" }
    return ""
}

func (statistics *Udp_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Udp_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-traffic" {
        return &statistics.Ipv4Traffic
    }
    if childYangName == "ipv6-traffic" {
        return &statistics.Ipv6Traffic
    }
    return nil
}

func (statistics *Udp_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-traffic"] = &statistics.Ipv4Traffic
    children["ipv6-traffic"] = &statistics.Ipv6Traffic
    return children
}

func (statistics *Udp_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Udp_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Udp_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Udp_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Udp_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Udp_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Udp_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Udp_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Udp_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Udp_Nodes_Node_Statistics_Ipv4Traffic
// UDP Traffic statistics for IPv4
type Udp_Nodes_Node_Statistics_Ipv4Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP Received. The type is interface{} with range: 0..4294967295.
    UdpInputPackets interface{}

    // UDP Checksum Errors. The type is interface{} with range: 0..4294967295.
    UdpChecksumErrorPackets interface{}

    // UDP No Port. The type is interface{} with range: 0..4294967295.
    UdpNoPortPackets interface{}

    // UDP bad length. The type is interface{} with range: 0..4294967295.
    UdpBadLengthPackets interface{}

    // UDP Transmitted. The type is interface{} with range: 0..4294967295.
    UdpOutputPackets interface{}

    // UDP drop for other reason. The type is interface{} with range:
    // 0..4294967295.
    UdpDroppedPackets interface{}
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetFilter() yfilter.YFilter { return ipv4Traffic.YFilter }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) SetFilter(yf yfilter.YFilter) { ipv4Traffic.YFilter = yf }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetGoName(yname string) string {
    if yname == "udp-input-packets" { return "UdpInputPackets" }
    if yname == "udp-checksum-error-packets" { return "UdpChecksumErrorPackets" }
    if yname == "udp-no-port-packets" { return "UdpNoPortPackets" }
    if yname == "udp-bad-length-packets" { return "UdpBadLengthPackets" }
    if yname == "udp-output-packets" { return "UdpOutputPackets" }
    if yname == "udp-dropped-packets" { return "UdpDroppedPackets" }
    return ""
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetSegmentPath() string {
    return "ipv4-traffic"
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udp-input-packets"] = ipv4Traffic.UdpInputPackets
    leafs["udp-checksum-error-packets"] = ipv4Traffic.UdpChecksumErrorPackets
    leafs["udp-no-port-packets"] = ipv4Traffic.UdpNoPortPackets
    leafs["udp-bad-length-packets"] = ipv4Traffic.UdpBadLengthPackets
    leafs["udp-output-packets"] = ipv4Traffic.UdpOutputPackets
    leafs["udp-dropped-packets"] = ipv4Traffic.UdpDroppedPackets
    return leafs
}

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetYangName() string { return "ipv4-traffic" }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) SetParent(parent types.Entity) { ipv4Traffic.parent = parent }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetParent() types.Entity { return ipv4Traffic.parent }

func (ipv4Traffic *Udp_Nodes_Node_Statistics_Ipv4Traffic) GetParentYangName() string { return "statistics" }

// Udp_Nodes_Node_Statistics_Ipv6Traffic
// UDP Traffic statistics for IPv6
type Udp_Nodes_Node_Statistics_Ipv6Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // UDP Received. The type is interface{} with range: 0..4294967295.
    UdpInputPackets interface{}

    // UDP Checksum Errors. The type is interface{} with range: 0..4294967295.
    UdpChecksumErrorPackets interface{}

    // UDP No Port. The type is interface{} with range: 0..4294967295.
    UdpNoPortPackets interface{}

    // UDP bad length. The type is interface{} with range: 0..4294967295.
    UdpBadLengthPackets interface{}

    // UDP Transmitted. The type is interface{} with range: 0..4294967295.
    UdpOutputPackets interface{}

    // UDP drop for other reason. The type is interface{} with range:
    // 0..4294967295.
    UdpDroppedPackets interface{}
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetFilter() yfilter.YFilter { return ipv6Traffic.YFilter }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) SetFilter(yf yfilter.YFilter) { ipv6Traffic.YFilter = yf }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetGoName(yname string) string {
    if yname == "udp-input-packets" { return "UdpInputPackets" }
    if yname == "udp-checksum-error-packets" { return "UdpChecksumErrorPackets" }
    if yname == "udp-no-port-packets" { return "UdpNoPortPackets" }
    if yname == "udp-bad-length-packets" { return "UdpBadLengthPackets" }
    if yname == "udp-output-packets" { return "UdpOutputPackets" }
    if yname == "udp-dropped-packets" { return "UdpDroppedPackets" }
    return ""
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetSegmentPath() string {
    return "ipv6-traffic"
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["udp-input-packets"] = ipv6Traffic.UdpInputPackets
    leafs["udp-checksum-error-packets"] = ipv6Traffic.UdpChecksumErrorPackets
    leafs["udp-no-port-packets"] = ipv6Traffic.UdpNoPortPackets
    leafs["udp-bad-length-packets"] = ipv6Traffic.UdpBadLengthPackets
    leafs["udp-output-packets"] = ipv6Traffic.UdpOutputPackets
    leafs["udp-dropped-packets"] = ipv6Traffic.UdpDroppedPackets
    return leafs
}

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetYangName() string { return "ipv6-traffic" }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) SetParent(parent types.Entity) { ipv6Traffic.parent = parent }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetParent() types.Entity { return ipv6Traffic.parent }

func (ipv6Traffic *Udp_Nodes_Node_Statistics_Ipv6Traffic) GetParentYangName() string { return "statistics" }

// UdpConnection
// udp connection
type UdpConnection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of UDP connections nodes.
    Nodes UdpConnection_Nodes
}

func (udpConnection *UdpConnection) GetFilter() yfilter.YFilter { return udpConnection.YFilter }

func (udpConnection *UdpConnection) SetFilter(yf yfilter.YFilter) { udpConnection.YFilter = yf }

func (udpConnection *UdpConnection) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (udpConnection *UdpConnection) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-udp-oper:udp-connection"
}

func (udpConnection *UdpConnection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &udpConnection.Nodes
    }
    return nil
}

func (udpConnection *UdpConnection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &udpConnection.Nodes
    return children
}

func (udpConnection *UdpConnection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (udpConnection *UdpConnection) GetBundleName() string { return "cisco_ios_xr" }

func (udpConnection *UdpConnection) GetYangName() string { return "udp-connection" }

func (udpConnection *UdpConnection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (udpConnection *UdpConnection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (udpConnection *UdpConnection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (udpConnection *UdpConnection) SetParent(parent types.Entity) { udpConnection.parent = parent }

func (udpConnection *UdpConnection) GetParent() types.Entity { return udpConnection.parent }

func (udpConnection *UdpConnection) GetParentYangName() string { return "Cisco-IOS-XR-ip-udp-oper" }

// UdpConnection_Nodes
// List of UDP connections nodes
type UdpConnection_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a particular node. The type is slice of
    // UdpConnection_Nodes_Node.
    Node []UdpConnection_Nodes_Node
}

func (nodes *UdpConnection_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *UdpConnection_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *UdpConnection_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *UdpConnection_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *UdpConnection_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *UdpConnection_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *UdpConnection_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *UdpConnection_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *UdpConnection_Nodes) GetYangName() string { return "nodes" }

func (nodes *UdpConnection_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *UdpConnection_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *UdpConnection_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *UdpConnection_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *UdpConnection_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *UdpConnection_Nodes) GetParentYangName() string { return "udp-connection" }

// UdpConnection_Nodes_Node
// Information about a particular node
type UdpConnection_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Statistics of UDP connections.
    Statistics UdpConnection_Nodes_Node_Statistics

    // LPTS statistical data.
    Lpts UdpConnection_Nodes_Node_Lpts

    // Detail information for list of UDP connections .
    PcbDetails UdpConnection_Nodes_Node_PcbDetails

    // Brief information for list of UDP connections.
    PcbBriefs UdpConnection_Nodes_Node_PcbBriefs
}

func (node *UdpConnection_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *UdpConnection_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *UdpConnection_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statistics" { return "Statistics" }
    if yname == "lpts" { return "Lpts" }
    if yname == "pcb-details" { return "PcbDetails" }
    if yname == "pcb-briefs" { return "PcbBriefs" }
    return ""
}

func (node *UdpConnection_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *UdpConnection_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    if childYangName == "lpts" {
        return &node.Lpts
    }
    if childYangName == "pcb-details" {
        return &node.PcbDetails
    }
    if childYangName == "pcb-briefs" {
        return &node.PcbBriefs
    }
    return nil
}

func (node *UdpConnection_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    children["lpts"] = &node.Lpts
    children["pcb-details"] = &node.PcbDetails
    children["pcb-briefs"] = &node.PcbBriefs
    return children
}

func (node *UdpConnection_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *UdpConnection_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *UdpConnection_Nodes_Node) GetYangName() string { return "node" }

func (node *UdpConnection_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *UdpConnection_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *UdpConnection_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *UdpConnection_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *UdpConnection_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *UdpConnection_Nodes_Node) GetParentYangName() string { return "nodes" }

// UdpConnection_Nodes_Node_Statistics
// Statistics of UDP connections
type UdpConnection_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table listing clients.
    Clients UdpConnection_Nodes_Node_Statistics_Clients

    // Summary statistics across all UDP connections.
    Summary UdpConnection_Nodes_Node_Statistics_Summary

    // Table listing the UDP connections for which statistics are provided.
    PcbStatistics UdpConnection_Nodes_Node_Statistics_PcbStatistics
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *UdpConnection_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "summary" { return "Summary" }
    if yname == "pcb-statistics" { return "PcbStatistics" }
    return ""
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clients" {
        return &statistics.Clients
    }
    if childYangName == "summary" {
        return &statistics.Summary
    }
    if childYangName == "pcb-statistics" {
        return &statistics.PcbStatistics
    }
    return nil
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clients"] = &statistics.Clients
    children["summary"] = &statistics.Summary
    children["pcb-statistics"] = &statistics.PcbStatistics
    return children
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *UdpConnection_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *UdpConnection_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *UdpConnection_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// UdpConnection_Nodes_Node_Statistics_Clients
// Table listing clients
type UdpConnection_Nodes_Node_Statistics_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing Client ID. The type is slice of
    // UdpConnection_Nodes_Node_Statistics_Clients_Client.
    Client []UdpConnection_Nodes_Node_Statistics_Clients_Client
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_Statistics_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetYangName() string { return "clients" }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetParent() types.Entity { return clients.parent }

func (clients *UdpConnection_Nodes_Node_Statistics_Clients) GetParentYangName() string { return "statistics" }

// UdpConnection_Nodes_Node_Statistics_Clients_Client
// Describing Client ID
type UdpConnection_Nodes_Node_Statistics_Clients_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Displaying client's aggregated statistics. The
    // type is interface{} with range: 0..4294967295.
    ClientId interface{}

    // Job ID of the transport client. The type is interface{} with range:
    // -2147483648..2147483647.
    ClientJid interface{}

    // Transport client name. The type is string with length: 0..21.
    ClientName interface{}

    // Total IPv4 packets received from client. The type is interface{} with
    // range: 0..4294967295.
    Ipv4ReceivedPackets interface{}

    // Total IPv4 packets sent to client. The type is interface{} with range:
    // 0..4294967295.
    Ipv4SentPackets interface{}

    // Total IPv6 packets received from app. The type is interface{} with range:
    // 0..4294967295.
    Ipv6ReceivedPackets interface{}

    // Total IPv6 packets sent to app. The type is interface{} with range:
    // 0..4294967295.
    Ipv6SentPackets interface{}
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "client-jid" { return "ClientJid" }
    if yname == "client-name" { return "ClientName" }
    if yname == "ipv4-received-packets" { return "Ipv4ReceivedPackets" }
    if yname == "ipv4-sent-packets" { return "Ipv4SentPackets" }
    if yname == "ipv6-received-packets" { return "Ipv6ReceivedPackets" }
    if yname == "ipv6-sent-packets" { return "Ipv6SentPackets" }
    return ""
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-id"] = client.ClientId
    leafs["client-jid"] = client.ClientJid
    leafs["client-name"] = client.ClientName
    leafs["ipv4-received-packets"] = client.Ipv4ReceivedPackets
    leafs["ipv4-sent-packets"] = client.Ipv4SentPackets
    leafs["ipv6-received-packets"] = client.Ipv6ReceivedPackets
    leafs["ipv6-sent-packets"] = client.Ipv6SentPackets
    return leafs
}

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetYangName() string { return "client" }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *UdpConnection_Nodes_Node_Statistics_Clients_Client) GetParentYangName() string { return "clients" }

// UdpConnection_Nodes_Node_Statistics_Summary
// Summary statistics across all UDP connections
type UdpConnection_Nodes_Node_Statistics_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Total packets received. The type is interface{} with range: 0..4294967295.
    ReceivedTotalPackets interface{}

    // Packets received when no wild listener. The type is interface{} with range:
    // 0..4294967295.
    ReceivedNoPortPackets interface{}

    // Packets received has bad checksum. The type is interface{} with range:
    // 0..4294967295.
    ReceivedBadChecksumPackets interface{}

    // Packets received is too short. The type is interface{} with range:
    // 0..4294967295.
    ReceivedTooShortPackets interface{}

    // Packets dropped for other reasons. The type is interface{} with range:
    // 0..4294967295.
    ReceivedDropPackets interface{}

    // Total packets sent. The type is interface{} with range: 0..4294967295.
    SentTotalPackets interface{}

    // Total send erorr packets. The type is interface{} with range:
    // 0..4294967295.
    SentErrorPackets interface{}

    // Total forwarding broadcast packets. The type is interface{} with range:
    // 0..4294967295.
    ForwardBroadcastPackets interface{}

    // Total cloned packets. The type is interface{} with range: 0..4294967295.
    ClonedPackets interface{}

    // Total failed cloned packets. The type is interface{} with range:
    // 0..4294967295.
    FailedClonePackets interface{}
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetGoName(yname string) string {
    if yname == "received-total-packets" { return "ReceivedTotalPackets" }
    if yname == "received-no-port-packets" { return "ReceivedNoPortPackets" }
    if yname == "received-bad-checksum-packets" { return "ReceivedBadChecksumPackets" }
    if yname == "received-too-short-packets" { return "ReceivedTooShortPackets" }
    if yname == "received-drop-packets" { return "ReceivedDropPackets" }
    if yname == "sent-total-packets" { return "SentTotalPackets" }
    if yname == "sent-error-packets" { return "SentErrorPackets" }
    if yname == "forward-broadcast-packets" { return "ForwardBroadcastPackets" }
    if yname == "cloned-packets" { return "ClonedPackets" }
    if yname == "failed-clone-packets" { return "FailedClonePackets" }
    return ""
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-total-packets"] = summary.ReceivedTotalPackets
    leafs["received-no-port-packets"] = summary.ReceivedNoPortPackets
    leafs["received-bad-checksum-packets"] = summary.ReceivedBadChecksumPackets
    leafs["received-too-short-packets"] = summary.ReceivedTooShortPackets
    leafs["received-drop-packets"] = summary.ReceivedDropPackets
    leafs["sent-total-packets"] = summary.SentTotalPackets
    leafs["sent-error-packets"] = summary.SentErrorPackets
    leafs["forward-broadcast-packets"] = summary.ForwardBroadcastPackets
    leafs["cloned-packets"] = summary.ClonedPackets
    leafs["failed-clone-packets"] = summary.FailedClonePackets
    return leafs
}

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetYangName() string { return "summary" }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetParent() types.Entity { return summary.parent }

func (summary *UdpConnection_Nodes_Node_Statistics_Summary) GetParentYangName() string { return "statistics" }

// UdpConnection_Nodes_Node_Statistics_PcbStatistics
// Table listing the UDP connections for which
// statistics are provided
type UdpConnection_Nodes_Node_Statistics_PcbStatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Satistics associated with a particular PCB. The type is slice of
    // UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic.
    PcbStatistic []UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetFilter() yfilter.YFilter { return pcbStatistics.YFilter }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) SetFilter(yf yfilter.YFilter) { pcbStatistics.YFilter = yf }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetGoName(yname string) string {
    if yname == "pcb-statistic" { return "PcbStatistic" }
    return ""
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetSegmentPath() string {
    return "pcb-statistics"
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcb-statistic" {
        for _, c := range pcbStatistics.PcbStatistic {
            if pcbStatistics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic{}
        pcbStatistics.PcbStatistic = append(pcbStatistics.PcbStatistic, child)
        return &pcbStatistics.PcbStatistic[len(pcbStatistics.PcbStatistic)-1]
    }
    return nil
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pcbStatistics.PcbStatistic {
        children[pcbStatistics.PcbStatistic[i].GetSegmentPath()] = &pcbStatistics.PcbStatistic[i]
    }
    return children
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetBundleName() string { return "cisco_ios_xr" }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetYangName() string { return "pcb-statistics" }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) SetParent(parent types.Entity) { pcbStatistics.parent = parent }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetParent() types.Entity { return pcbStatistics.parent }

func (pcbStatistics *UdpConnection_Nodes_Node_Statistics_PcbStatistics) GetParentYangName() string { return "statistics" }

// UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic
// Satistics associated with a particular PCB
type UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block address. The type is
    // interface{} with range: 0..4294967295.
    PcbAddress interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // True if paw socket. The type is bool.
    IsPawSocket interface{}

    // UDP send statistics.
    Send UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send

    // UDP receive statistics.
    Receive UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetFilter() yfilter.YFilter { return pcbStatistic.YFilter }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) SetFilter(yf yfilter.YFilter) { pcbStatistic.YFilter = yf }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetGoName(yname string) string {
    if yname == "pcb-address" { return "PcbAddress" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "is-paw-socket" { return "IsPawSocket" }
    if yname == "send" { return "Send" }
    if yname == "receive" { return "Receive" }
    return ""
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetSegmentPath() string {
    return "pcb-statistic" + "[pcb-address='" + fmt.Sprintf("%v", pcbStatistic.PcbAddress) + "']"
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "send" {
        return &pcbStatistic.Send
    }
    if childYangName == "receive" {
        return &pcbStatistic.Receive
    }
    return nil
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["send"] = &pcbStatistic.Send
    children["receive"] = &pcbStatistic.Receive
    return children
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-address"] = pcbStatistic.PcbAddress
    leafs["vrf-id"] = pcbStatistic.VrfId
    leafs["is-paw-socket"] = pcbStatistic.IsPawSocket
    return leafs
}

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetYangName() string { return "pcb-statistic" }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) SetParent(parent types.Entity) { pcbStatistic.parent = parent }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetParent() types.Entity { return pcbStatistic.parent }

func (pcbStatistic *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic) GetParentYangName() string { return "pcb-statistics" }

// UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send
// UDP send statistics
type UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Bytes received from application. The type is interface{} with range:
    // 0..18446744073709551615. Units are byte.
    ReceivedApplicationBytes interface{}

    // XIPC pulses received from application. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedXipcPulses interface{}

    // Packets sent to network (v4/v6 IO). The type is interface{} with range:
    // 0..18446744073709551615.
    SentNetworkPackets interface{}

    // Packets sent to network (NetIO). The type is interface{} with range:
    // 0..18446744073709551615.
    SentNetIoPackets interface{}

    // Packets failed getting queued to network (v4/v6 IO). The type is
    // interface{} with range: 0..4294967295.
    FailedQueuedNetworkPackets interface{}

    // Packets failed getting queued to network (NetIO). The type is interface{}
    // with range: 0..4294967295.
    FailedQueuedNetIoPackets interface{}
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetFilter() yfilter.YFilter { return send.YFilter }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) SetFilter(yf yfilter.YFilter) { send.YFilter = yf }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetGoName(yname string) string {
    if yname == "received-application-bytes" { return "ReceivedApplicationBytes" }
    if yname == "received-xipc-pulses" { return "ReceivedXipcPulses" }
    if yname == "sent-network-packets" { return "SentNetworkPackets" }
    if yname == "sent-net-io-packets" { return "SentNetIoPackets" }
    if yname == "failed-queued-network-packets" { return "FailedQueuedNetworkPackets" }
    if yname == "failed-queued-net-io-packets" { return "FailedQueuedNetIoPackets" }
    return ""
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetSegmentPath() string {
    return "send"
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-application-bytes"] = send.ReceivedApplicationBytes
    leafs["received-xipc-pulses"] = send.ReceivedXipcPulses
    leafs["sent-network-packets"] = send.SentNetworkPackets
    leafs["sent-net-io-packets"] = send.SentNetIoPackets
    leafs["failed-queued-network-packets"] = send.FailedQueuedNetworkPackets
    leafs["failed-queued-net-io-packets"] = send.FailedQueuedNetIoPackets
    return leafs
}

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetBundleName() string { return "cisco_ios_xr" }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetYangName() string { return "send" }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) SetParent(parent types.Entity) { send.parent = parent }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetParent() types.Entity { return send.parent }

func (send *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Send) GetParentYangName() string { return "pcb-statistic" }

// UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive
// UDP receive statistics
type UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Packets received from network. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedNetworkPackets interface{}

    // Packets failed queued to application. The type is interface{} with range:
    // 0..4294967295.
    FailedQueuedApplicationPackets interface{}

    // Packets queued to application. The type is interface{} with range:
    // 0..18446744073709551615.
    QueuedApplicationPackets interface{}

    // Packet that couldn't be queued to application.on socket. The type is
    // interface{} with range: 0..4294967295.
    FailedQueuedApplicationSocketPackets interface{}

    // Packets queued to application on socket. The type is interface{} with
    // range: 0..18446744073709551615.
    QueuedApplicationSocketPackets interface{}
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetFilter() yfilter.YFilter { return receive.YFilter }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) SetFilter(yf yfilter.YFilter) { receive.YFilter = yf }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetGoName(yname string) string {
    if yname == "received-network-packets" { return "ReceivedNetworkPackets" }
    if yname == "failed-queued-application-packets" { return "FailedQueuedApplicationPackets" }
    if yname == "queued-application-packets" { return "QueuedApplicationPackets" }
    if yname == "failed-queued-application-socket-packets" { return "FailedQueuedApplicationSocketPackets" }
    if yname == "queued-application-socket-packets" { return "QueuedApplicationSocketPackets" }
    return ""
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetSegmentPath() string {
    return "receive"
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["received-network-packets"] = receive.ReceivedNetworkPackets
    leafs["failed-queued-application-packets"] = receive.FailedQueuedApplicationPackets
    leafs["queued-application-packets"] = receive.QueuedApplicationPackets
    leafs["failed-queued-application-socket-packets"] = receive.FailedQueuedApplicationSocketPackets
    leafs["queued-application-socket-packets"] = receive.QueuedApplicationSocketPackets
    return leafs
}

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetBundleName() string { return "cisco_ios_xr" }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetYangName() string { return "receive" }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) SetParent(parent types.Entity) { receive.parent = parent }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetParent() types.Entity { return receive.parent }

func (receive *UdpConnection_Nodes_Node_Statistics_PcbStatistics_PcbStatistic_Receive) GetParentYangName() string { return "pcb-statistic" }

// UdpConnection_Nodes_Node_Lpts
// LPTS statistical data
type UdpConnection_Nodes_Node_Lpts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of query options.
    Queries UdpConnection_Nodes_Node_Lpts_Queries
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetFilter() yfilter.YFilter { return lpts.YFilter }

func (lpts *UdpConnection_Nodes_Node_Lpts) SetFilter(yf yfilter.YFilter) { lpts.YFilter = yf }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetGoName(yname string) string {
    if yname == "queries" { return "Queries" }
    return ""
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetSegmentPath() string {
    return "lpts"
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "queries" {
        return &lpts.Queries
    }
    return nil
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["queries"] = &lpts.Queries
    return children
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lpts *UdpConnection_Nodes_Node_Lpts) GetBundleName() string { return "cisco_ios_xr" }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetYangName() string { return "lpts" }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lpts *UdpConnection_Nodes_Node_Lpts) SetParent(parent types.Entity) { lpts.parent = parent }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetParent() types.Entity { return lpts.parent }

func (lpts *UdpConnection_Nodes_Node_Lpts) GetParentYangName() string { return "node" }

// UdpConnection_Nodes_Node_Lpts_Queries
// List of query options
type UdpConnection_Nodes_Node_Lpts_Queries struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Query option. The type is slice of
    // UdpConnection_Nodes_Node_Lpts_Queries_Query.
    Query []UdpConnection_Nodes_Node_Lpts_Queries_Query
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetFilter() yfilter.YFilter { return queries.YFilter }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) SetFilter(yf yfilter.YFilter) { queries.YFilter = yf }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetGoName(yname string) string {
    if yname == "query" { return "Query" }
    return ""
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetSegmentPath() string {
    return "queries"
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "query" {
        for _, c := range queries.Query {
            if queries.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_Lpts_Queries_Query{}
        queries.Query = append(queries.Query, child)
        return &queries.Query[len(queries.Query)-1]
    }
    return nil
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range queries.Query {
        children[queries.Query[i].GetSegmentPath()] = &queries.Query[i]
    }
    return children
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetBundleName() string { return "cisco_ios_xr" }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetYangName() string { return "queries" }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) SetParent(parent types.Entity) { queries.parent = parent }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetParent() types.Entity { return queries.parent }

func (queries *UdpConnection_Nodes_Node_Lpts_Queries) GetParentYangName() string { return "lpts" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query
// Query option
type UdpConnection_Nodes_Node_Lpts_Queries_Query struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Query option. The type is LptsPcbQuery.
    QueryName interface{}

    // List of PCBs.
    Pcbs UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetFilter() yfilter.YFilter { return query.YFilter }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) SetFilter(yf yfilter.YFilter) { query.YFilter = yf }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetGoName(yname string) string {
    if yname == "query-name" { return "QueryName" }
    if yname == "pcbs" { return "Pcbs" }
    return ""
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetSegmentPath() string {
    return "query" + "[query-name='" + fmt.Sprintf("%v", query.QueryName) + "']"
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcbs" {
        return &query.Pcbs
    }
    return nil
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pcbs"] = &query.Pcbs
    return children
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["query-name"] = query.QueryName
    return leafs
}

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetBundleName() string { return "cisco_ios_xr" }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetYangName() string { return "query" }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) SetParent(parent types.Entity) { query.parent = parent }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetParent() types.Entity { return query.parent }

func (query *UdpConnection_Nodes_Node_Lpts_Queries_Query) GetParentYangName() string { return "queries" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs
// List of PCBs
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A PCB information. The type is slice of
    // UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb.
    Pcb []UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetFilter() yfilter.YFilter { return pcbs.YFilter }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) SetFilter(yf yfilter.YFilter) { pcbs.YFilter = yf }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetGoName(yname string) string {
    if yname == "pcb" { return "Pcb" }
    return ""
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetSegmentPath() string {
    return "pcbs"
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcb" {
        for _, c := range pcbs.Pcb {
            if pcbs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb{}
        pcbs.Pcb = append(pcbs.Pcb, child)
        return &pcbs.Pcb[len(pcbs.Pcb)-1]
    }
    return nil
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pcbs.Pcb {
        children[pcbs.Pcb[i].GetSegmentPath()] = &pcbs.Pcb[i]
    }
    return children
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetBundleName() string { return "cisco_ios_xr" }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetYangName() string { return "pcbs" }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) SetParent(parent types.Entity) { pcbs.parent = parent }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetParent() types.Entity { return pcbs.parent }

func (pcbs *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs) GetParentYangName() string { return "query" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb
// A PCB information
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. PCB address. The type is interface{} with range:
    // 0..4294967295.
    PcbAddress interface{}

    // Layer 4 protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Local IP address.
    LocalAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress

    // Remote IP address.
    ForeignAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress

    // Common PCB information.
    Common UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetFilter() yfilter.YFilter { return pcb.YFilter }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) SetFilter(yf yfilter.YFilter) { pcb.YFilter = yf }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetGoName(yname string) string {
    if yname == "pcb-address" { return "PcbAddress" }
    if yname == "l4-protocol" { return "L4Protocol" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    if yname == "common" { return "Common" }
    return ""
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetSegmentPath() string {
    return "pcb" + "[pcb-address='" + fmt.Sprintf("%v", pcb.PcbAddress) + "']"
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &pcb.LocalAddress
    }
    if childYangName == "foreign-address" {
        return &pcb.ForeignAddress
    }
    if childYangName == "common" {
        return &pcb.Common
    }
    return nil
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &pcb.LocalAddress
    children["foreign-address"] = &pcb.ForeignAddress
    children["common"] = &pcb.Common
    return children
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-address"] = pcb.PcbAddress
    leafs["l4-protocol"] = pcb.L4Protocol
    leafs["local-port"] = pcb.LocalPort
    leafs["foreign-port"] = pcb.ForeignPort
    return leafs
}

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetBundleName() string { return "cisco_ios_xr" }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetYangName() string { return "pcb" }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) SetParent(parent types.Entity) { pcb.parent = parent }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetParent() types.Entity { return pcb.parent }

func (pcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb) GetParentYangName() string { return "pcbs" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress
// Local IP address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_LocalAddress) GetParentYangName() string { return "pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress
// Remote IP address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetFilter() yfilter.YFilter { return foreignAddress.YFilter }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) SetFilter(yf yfilter.YFilter) { foreignAddress.YFilter = yf }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetSegmentPath() string {
    return "foreign-address"
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = foreignAddress.AfName
    leafs["ipv4-address"] = foreignAddress.Ipv4Address
    leafs["ipv6-address"] = foreignAddress.Ipv6Address
    return leafs
}

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetBundleName() string { return "cisco_ios_xr" }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetYangName() string { return "foreign-address" }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) SetParent(parent types.Entity) { foreignAddress.parent = parent }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetParent() types.Entity { return foreignAddress.parent }

func (foreignAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_ForeignAddress) GetParentYangName() string { return "pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common
// Common PCB information
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family. The type is AddrFamily.
    AfName interface{}

    // LPTS PCB information.
    LptsPcb UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetFilter() yfilter.YFilter { return common.YFilter }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) SetFilter(yf yfilter.YFilter) { common.YFilter = yf }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "lpts-pcb" { return "LptsPcb" }
    return ""
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetSegmentPath() string {
    return "common"
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lpts-pcb" {
        return &common.LptsPcb
    }
    return nil
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lpts-pcb"] = &common.LptsPcb
    return children
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = common.AfName
    return leafs
}

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetBundleName() string { return "cisco_ios_xr" }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetYangName() string { return "common" }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) SetParent(parent types.Entity) { common.parent = parent }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetParent() types.Entity { return common.parent }

func (common *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common) GetParentYangName() string { return "pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb
// LPTS PCB information
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // flow information. The type is interface{} with range: 0..4294967295.
    FlowTypesInfo interface{}

    // Receive options.
    Options UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options

    // LPTS flags.
    LptsFlags UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags

    // AcceptMask.
    AcceptMask UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask

    // Interface Filters. The type is slice of
    // UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter.
    Filter []UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetFilter() yfilter.YFilter { return lptsPcb.YFilter }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) SetFilter(yf yfilter.YFilter) { lptsPcb.YFilter = yf }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetGoName(yname string) string {
    if yname == "ttl" { return "Ttl" }
    if yname == "flow-types-info" { return "FlowTypesInfo" }
    if yname == "options" { return "Options" }
    if yname == "lpts-flags" { return "LptsFlags" }
    if yname == "accept-mask" { return "AcceptMask" }
    if yname == "filter" { return "Filter" }
    return ""
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetSegmentPath() string {
    return "lpts-pcb"
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "options" {
        return &lptsPcb.Options
    }
    if childYangName == "lpts-flags" {
        return &lptsPcb.LptsFlags
    }
    if childYangName == "accept-mask" {
        return &lptsPcb.AcceptMask
    }
    if childYangName == "filter" {
        for _, c := range lptsPcb.Filter {
            if lptsPcb.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter{}
        lptsPcb.Filter = append(lptsPcb.Filter, child)
        return &lptsPcb.Filter[len(lptsPcb.Filter)-1]
    }
    return nil
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["options"] = &lptsPcb.Options
    children["lpts-flags"] = &lptsPcb.LptsFlags
    children["accept-mask"] = &lptsPcb.AcceptMask
    for i := range lptsPcb.Filter {
        children[lptsPcb.Filter[i].GetSegmentPath()] = &lptsPcb.Filter[i]
    }
    return children
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ttl"] = lptsPcb.Ttl
    leafs["flow-types-info"] = lptsPcb.FlowTypesInfo
    return leafs
}

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetBundleName() string { return "cisco_ios_xr" }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetYangName() string { return "lpts-pcb" }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) SetParent(parent types.Entity) { lptsPcb.parent = parent }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetParent() types.Entity { return lptsPcb.parent }

func (lptsPcb *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb) GetParentYangName() string { return "common" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options
// Receive options
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Receive filter enabled. The type is bool.
    IsReceiveFilter interface{}

    // IP SLA. The type is bool.
    IsIpSla interface{}
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetGoName(yname string) string {
    if yname == "is-receive-filter" { return "IsReceiveFilter" }
    if yname == "is-ip-sla" { return "IsIpSla" }
    return ""
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetSegmentPath() string {
    return "options"
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-receive-filter"] = options.IsReceiveFilter
    leafs["is-ip-sla"] = options.IsIpSla
    return leafs
}

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetYangName() string { return "options" }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetParent() types.Entity { return options.parent }

func (options *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Options) GetParentYangName() string { return "lpts-pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags
// LPTS flags
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCB bound. The type is bool.
    IsPcbBound interface{}

    // Sent drop packets. The type is bool.
    IsLocalAddressIgnore interface{}

    // Ignore VRF Filter. The type is bool.
    IsIgnoreVrfFilter interface{}
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetFilter() yfilter.YFilter { return lptsFlags.YFilter }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) SetFilter(yf yfilter.YFilter) { lptsFlags.YFilter = yf }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetGoName(yname string) string {
    if yname == "is-pcb-bound" { return "IsPcbBound" }
    if yname == "is-local-address-ignore" { return "IsLocalAddressIgnore" }
    if yname == "is-ignore-vrf-filter" { return "IsIgnoreVrfFilter" }
    return ""
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetSegmentPath() string {
    return "lpts-flags"
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-pcb-bound"] = lptsFlags.IsPcbBound
    leafs["is-local-address-ignore"] = lptsFlags.IsLocalAddressIgnore
    leafs["is-ignore-vrf-filter"] = lptsFlags.IsIgnoreVrfFilter
    return leafs
}

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetBundleName() string { return "cisco_ios_xr" }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetYangName() string { return "lpts-flags" }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) SetParent(parent types.Entity) { lptsFlags.parent = parent }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetParent() types.Entity { return lptsFlags.parent }

func (lptsFlags *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_LptsFlags) GetParentYangName() string { return "lpts-pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask
// AcceptMask
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set interface. The type is bool.
    IsInterface interface{}

    // Set packet type. The type is bool.
    IsPacketType interface{}

    // Set Remote address. The type is bool.
    IsRemoteAddress interface{}

    // Set Remote Port. The type is bool.
    IsRemotePort interface{}

    // Set Local Address. The type is bool.
    IsLocalAddress interface{}

    // Set Local Port. The type is bool.
    IsLocalPort interface{}
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetFilter() yfilter.YFilter { return acceptMask.YFilter }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) SetFilter(yf yfilter.YFilter) { acceptMask.YFilter = yf }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetGoName(yname string) string {
    if yname == "is-interface" { return "IsInterface" }
    if yname == "is-packet-type" { return "IsPacketType" }
    if yname == "is-remote-address" { return "IsRemoteAddress" }
    if yname == "is-remote-port" { return "IsRemotePort" }
    if yname == "is-local-address" { return "IsLocalAddress" }
    if yname == "is-local-port" { return "IsLocalPort" }
    return ""
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetSegmentPath() string {
    return "accept-mask"
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-interface"] = acceptMask.IsInterface
    leafs["is-packet-type"] = acceptMask.IsPacketType
    leafs["is-remote-address"] = acceptMask.IsRemoteAddress
    leafs["is-remote-port"] = acceptMask.IsRemotePort
    leafs["is-local-address"] = acceptMask.IsLocalAddress
    leafs["is-local-port"] = acceptMask.IsLocalPort
    return leafs
}

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetBundleName() string { return "cisco_ios_xr" }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetYangName() string { return "accept-mask" }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) SetParent(parent types.Entity) { acceptMask.parent = parent }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetParent() types.Entity { return acceptMask.parent }

func (acceptMask *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_AcceptMask) GetParentYangName() string { return "lpts-pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter
// Interface Filters
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface name. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Remote address length. The type is interface{} with range: 0..65535.
    RemoteLength interface{}

    // Local address length. The type is interface{} with range: 0..65535.
    LocalLength interface{}

    // Receive Remote port. The type is interface{} with range: 0..65535.
    ReceiveRemotePort interface{}

    // Receive Local port. The type is interface{} with range: 0..65535.
    ReceiveLocalPort interface{}

    // Priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Minimum TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // flow information. The type is interface{} with range: 0..4294967295.
    FlowTypesInfo interface{}

    // Protocol-specific packet type.
    PacketType UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType

    // Remote address.
    RemoteAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress

    // Local address.
    LocalAddress UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetFilter() yfilter.YFilter { return filter.YFilter }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) SetFilter(yf yfilter.YFilter) { filter.YFilter = yf }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "remote-length" { return "RemoteLength" }
    if yname == "local-length" { return "LocalLength" }
    if yname == "receive-remote-port" { return "ReceiveRemotePort" }
    if yname == "receive-local-port" { return "ReceiveLocalPort" }
    if yname == "priority" { return "Priority" }
    if yname == "ttl" { return "Ttl" }
    if yname == "flow-types-info" { return "FlowTypesInfo" }
    if yname == "packet-type" { return "PacketType" }
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "local-address" { return "LocalAddress" }
    return ""
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetSegmentPath() string {
    return "filter"
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "packet-type" {
        return &filter.PacketType
    }
    if childYangName == "remote-address" {
        return &filter.RemoteAddress
    }
    if childYangName == "local-address" {
        return &filter.LocalAddress
    }
    return nil
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-type"] = &filter.PacketType
    children["remote-address"] = &filter.RemoteAddress
    children["local-address"] = &filter.LocalAddress
    return children
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = filter.InterfaceName
    leafs["remote-length"] = filter.RemoteLength
    leafs["local-length"] = filter.LocalLength
    leafs["receive-remote-port"] = filter.ReceiveRemotePort
    leafs["receive-local-port"] = filter.ReceiveLocalPort
    leafs["priority"] = filter.Priority
    leafs["ttl"] = filter.Ttl
    leafs["flow-types-info"] = filter.FlowTypesInfo
    return leafs
}

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetBundleName() string { return "cisco_ios_xr" }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetYangName() string { return "filter" }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) SetParent(parent types.Entity) { filter.parent = parent }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetParent() types.Entity { return filter.parent }

func (filter *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter) GetParentYangName() string { return "lpts-pcb" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType
// Protocol-specific packet type
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type. The type is Packet.
    Type interface{}

    // ICMP message type. The type is MessageTypeIcmp.
    IcmpMessageType interface{}

    // ICMPv6 message type. The type is MessageTypeIcmpv6.
    IcmPv6MessageType interface{}

    // IGMP message type. The type is MessageTypeIgmp.
    IgmpMessageType interface{}

    // Message type in number. The type is interface{} with range: 0..4294967295.
    MessageId interface{}
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetFilter() yfilter.YFilter { return packetType.YFilter }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) SetFilter(yf yfilter.YFilter) { packetType.YFilter = yf }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "icmp-message-type" { return "IcmpMessageType" }
    if yname == "icm-pv6-message-type" { return "IcmPv6MessageType" }
    if yname == "igmp-message-type" { return "IgmpMessageType" }
    if yname == "message-id" { return "MessageId" }
    return ""
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetSegmentPath() string {
    return "packet-type"
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = packetType.Type
    leafs["icmp-message-type"] = packetType.IcmpMessageType
    leafs["icm-pv6-message-type"] = packetType.IcmPv6MessageType
    leafs["igmp-message-type"] = packetType.IgmpMessageType
    leafs["message-id"] = packetType.MessageId
    return leafs
}

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetBundleName() string { return "cisco_ios_xr" }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetYangName() string { return "packet-type" }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) SetParent(parent types.Entity) { packetType.parent = parent }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetParent() types.Entity { return packetType.parent }

func (packetType *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_PacketType) GetParentYangName() string { return "filter" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress
// Remote address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetFilter() yfilter.YFilter { return remoteAddress.YFilter }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) SetFilter(yf yfilter.YFilter) { remoteAddress.YFilter = yf }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetSegmentPath() string {
    return "remote-address"
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = remoteAddress.AfName
    leafs["ipv4-address"] = remoteAddress.Ipv4Address
    leafs["ipv6-address"] = remoteAddress.Ipv6Address
    return leafs
}

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetBundleName() string { return "cisco_ios_xr" }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetYangName() string { return "remote-address" }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) SetParent(parent types.Entity) { remoteAddress.parent = parent }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetParent() types.Entity { return remoteAddress.parent }

func (remoteAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_RemoteAddress) GetParentYangName() string { return "filter" }

// UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress
// Local address
type UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is AddrFamily.
    AfName interface{}

    // IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *UdpConnection_Nodes_Node_Lpts_Queries_Query_Pcbs_Pcb_Common_LptsPcb_Filter_LocalAddress) GetParentYangName() string { return "filter" }

// UdpConnection_Nodes_Node_PcbDetails
// Detail information for list of UDP connections
// .
type UdpConnection_Nodes_Node_PcbDetails struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detail information about a UDP connection. The type is slice of
    // UdpConnection_Nodes_Node_PcbDetails_PcbDetail.
    PcbDetail []UdpConnection_Nodes_Node_PcbDetails_PcbDetail
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetFilter() yfilter.YFilter { return pcbDetails.YFilter }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) SetFilter(yf yfilter.YFilter) { pcbDetails.YFilter = yf }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetGoName(yname string) string {
    if yname == "pcb-detail" { return "PcbDetail" }
    return ""
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetSegmentPath() string {
    return "pcb-details"
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcb-detail" {
        for _, c := range pcbDetails.PcbDetail {
            if pcbDetails.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_PcbDetails_PcbDetail{}
        pcbDetails.PcbDetail = append(pcbDetails.PcbDetail, child)
        return &pcbDetails.PcbDetail[len(pcbDetails.PcbDetail)-1]
    }
    return nil
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pcbDetails.PcbDetail {
        children[pcbDetails.PcbDetail[i].GetSegmentPath()] = &pcbDetails.PcbDetail[i]
    }
    return children
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetBundleName() string { return "cisco_ios_xr" }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetYangName() string { return "pcb-details" }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) SetParent(parent types.Entity) { pcbDetails.parent = parent }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetParent() types.Entity { return pcbDetails.parent }

func (pcbDetails *UdpConnection_Nodes_Node_PcbDetails) GetParentYangName() string { return "node" }

// UdpConnection_Nodes_Node_PcbDetails_PcbDetail
// Detail information about a UDP connection
type UdpConnection_Nodes_Node_PcbDetails_PcbDetail struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block address. The type is
    // interface{} with range: 0..4294967295.
    PcbAddress interface{}

    // Address family. The type is UdpAddressFamily.
    AfName interface{}

    // ID of local process. The type is interface{} with range: 0..4294967295.
    LocalProcessId interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Receive queue count. The type is interface{} with range: 0..4294967295.
    ReceiveQueue interface{}

    // Send queue count. The type is interface{} with range: 0..4294967295.
    SendQueue interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Local address.
    LocalAddress UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress

    // Foreign address.
    ForeignAddress UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetFilter() yfilter.YFilter { return pcbDetail.YFilter }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) SetFilter(yf yfilter.YFilter) { pcbDetail.YFilter = yf }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetGoName(yname string) string {
    if yname == "pcb-address" { return "PcbAddress" }
    if yname == "af-name" { return "AfName" }
    if yname == "local-process-id" { return "LocalProcessId" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "receive-queue" { return "ReceiveQueue" }
    if yname == "send-queue" { return "SendQueue" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    return ""
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetSegmentPath() string {
    return "pcb-detail" + "[pcb-address='" + fmt.Sprintf("%v", pcbDetail.PcbAddress) + "']"
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &pcbDetail.LocalAddress
    }
    if childYangName == "foreign-address" {
        return &pcbDetail.ForeignAddress
    }
    return nil
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &pcbDetail.LocalAddress
    children["foreign-address"] = &pcbDetail.ForeignAddress
    return children
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-address"] = pcbDetail.PcbAddress
    leafs["af-name"] = pcbDetail.AfName
    leafs["local-process-id"] = pcbDetail.LocalProcessId
    leafs["local-port"] = pcbDetail.LocalPort
    leafs["foreign-port"] = pcbDetail.ForeignPort
    leafs["receive-queue"] = pcbDetail.ReceiveQueue
    leafs["send-queue"] = pcbDetail.SendQueue
    leafs["vrf-id"] = pcbDetail.VrfId
    return leafs
}

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetBundleName() string { return "cisco_ios_xr" }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetYangName() string { return "pcb-detail" }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) SetParent(parent types.Entity) { pcbDetail.parent = parent }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetParent() types.Entity { return pcbDetail.parent }

func (pcbDetail *UdpConnection_Nodes_Node_PcbDetails_PcbDetail) GetParentYangName() string { return "pcb-details" }

// UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress
// Local address
type UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_LocalAddress) GetParentYangName() string { return "pcb-detail" }

// UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress
// Foreign address
type UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetFilter() yfilter.YFilter { return foreignAddress.YFilter }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) SetFilter(yf yfilter.YFilter) { foreignAddress.YFilter = yf }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetSegmentPath() string {
    return "foreign-address"
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = foreignAddress.AfName
    leafs["ipv4-address"] = foreignAddress.Ipv4Address
    leafs["ipv6-address"] = foreignAddress.Ipv6Address
    return leafs
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetBundleName() string { return "cisco_ios_xr" }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetYangName() string { return "foreign-address" }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) SetParent(parent types.Entity) { foreignAddress.parent = parent }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetParent() types.Entity { return foreignAddress.parent }

func (foreignAddress *UdpConnection_Nodes_Node_PcbDetails_PcbDetail_ForeignAddress) GetParentYangName() string { return "pcb-detail" }

// UdpConnection_Nodes_Node_PcbBriefs
// Brief information for list of UDP connections.
type UdpConnection_Nodes_Node_PcbBriefs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about a UDP connection. The type is slice of
    // UdpConnection_Nodes_Node_PcbBriefs_PcbBrief.
    PcbBrief []UdpConnection_Nodes_Node_PcbBriefs_PcbBrief
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetFilter() yfilter.YFilter { return pcbBriefs.YFilter }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) SetFilter(yf yfilter.YFilter) { pcbBriefs.YFilter = yf }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetGoName(yname string) string {
    if yname == "pcb-brief" { return "PcbBrief" }
    return ""
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetSegmentPath() string {
    return "pcb-briefs"
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcb-brief" {
        for _, c := range pcbBriefs.PcbBrief {
            if pcbBriefs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := UdpConnection_Nodes_Node_PcbBriefs_PcbBrief{}
        pcbBriefs.PcbBrief = append(pcbBriefs.PcbBrief, child)
        return &pcbBriefs.PcbBrief[len(pcbBriefs.PcbBrief)-1]
    }
    return nil
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pcbBriefs.PcbBrief {
        children[pcbBriefs.PcbBrief[i].GetSegmentPath()] = &pcbBriefs.PcbBrief[i]
    }
    return children
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetBundleName() string { return "cisco_ios_xr" }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetYangName() string { return "pcb-briefs" }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) SetParent(parent types.Entity) { pcbBriefs.parent = parent }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetParent() types.Entity { return pcbBriefs.parent }

func (pcbBriefs *UdpConnection_Nodes_Node_PcbBriefs) GetParentYangName() string { return "node" }

// UdpConnection_Nodes_Node_PcbBriefs_PcbBrief
// Brief information about a UDP connection
type UdpConnection_Nodes_Node_PcbBriefs_PcbBrief struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block address. The type is
    // interface{} with range: 0..4294967295.
    PcbAddress interface{}

    // Address family. The type is UdpAddressFamily.
    AfName interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Receive queue count. The type is interface{} with range: 0..4294967295.
    ReceiveQueue interface{}

    // Send queue count. The type is interface{} with range: 0..4294967295.
    SendQueue interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Local address.
    LocalAddress UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress

    // Foreign address.
    ForeignAddress UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetFilter() yfilter.YFilter { return pcbBrief.YFilter }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) SetFilter(yf yfilter.YFilter) { pcbBrief.YFilter = yf }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetGoName(yname string) string {
    if yname == "pcb-address" { return "PcbAddress" }
    if yname == "af-name" { return "AfName" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "receive-queue" { return "ReceiveQueue" }
    if yname == "send-queue" { return "SendQueue" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    return ""
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetSegmentPath() string {
    return "pcb-brief" + "[pcb-address='" + fmt.Sprintf("%v", pcbBrief.PcbAddress) + "']"
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &pcbBrief.LocalAddress
    }
    if childYangName == "foreign-address" {
        return &pcbBrief.ForeignAddress
    }
    return nil
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &pcbBrief.LocalAddress
    children["foreign-address"] = &pcbBrief.ForeignAddress
    return children
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-address"] = pcbBrief.PcbAddress
    leafs["af-name"] = pcbBrief.AfName
    leafs["local-port"] = pcbBrief.LocalPort
    leafs["foreign-port"] = pcbBrief.ForeignPort
    leafs["receive-queue"] = pcbBrief.ReceiveQueue
    leafs["send-queue"] = pcbBrief.SendQueue
    leafs["vrf-id"] = pcbBrief.VrfId
    return leafs
}

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetBundleName() string { return "cisco_ios_xr" }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetYangName() string { return "pcb-brief" }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) SetParent(parent types.Entity) { pcbBrief.parent = parent }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetParent() types.Entity { return pcbBrief.parent }

func (pcbBrief *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief) GetParentYangName() string { return "pcb-briefs" }

// UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress
// Local address
type UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_LocalAddress) GetParentYangName() string { return "pcb-brief" }

// UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress
// Foreign address
type UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is UdpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetFilter() yfilter.YFilter { return foreignAddress.YFilter }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) SetFilter(yf yfilter.YFilter) { foreignAddress.YFilter = yf }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetSegmentPath() string {
    return "foreign-address"
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = foreignAddress.AfName
    leafs["ipv4-address"] = foreignAddress.Ipv4Address
    leafs["ipv6-address"] = foreignAddress.Ipv6Address
    return leafs
}

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetBundleName() string { return "cisco_ios_xr" }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetYangName() string { return "foreign-address" }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) SetParent(parent types.Entity) { foreignAddress.parent = parent }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetParent() types.Entity { return foreignAddress.parent }

func (foreignAddress *UdpConnection_Nodes_Node_PcbBriefs_PcbBrief_ForeignAddress) GetParentYangName() string { return "pcb-brief" }

