// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-tcp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   tcp-connection: TCP connection operational data
//   tcp: tcp
//   tcp-nsr: tcp nsr
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_tcp_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_tcp_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-tcp-oper tcp-connection}", reflect.TypeOf(TcpConnection{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-tcp-oper:tcp-connection", reflect.TypeOf(TcpConnection{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-tcp-oper tcp}", reflect.TypeOf(Tcp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-tcp-oper:tcp", reflect.TypeOf(Tcp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-tcp-oper tcp-nsr}", reflect.TypeOf(TcpNsr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-tcp-oper:tcp-nsr", reflect.TypeOf(TcpNsr{}))
}

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

// Show represents Show
type Show string

const (
    // To dispay all
    Show_all Show = "all"

    // To display static policy
    Show_static_policy Show = "static-policy"

    // To display interface filter
    Show_interface_filter Show = "interface-filter"

    // To display packet type filter
    Show_packet_filter Show = "packet-filter"
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

// TcpConnState represents TCP Connection State
type TcpConnState string

const (
    // Closed
    TcpConnState_closed TcpConnState = "closed"

    // Listen
    TcpConnState_listen TcpConnState = "listen"

    // Syn sent
    TcpConnState_syn_sent TcpConnState = "syn-sent"

    // Syn received
    TcpConnState_syn_received TcpConnState = "syn-received"

    // Established
    TcpConnState_established TcpConnState = "established"

    // Close wait
    TcpConnState_close_wait TcpConnState = "close-wait"

    // FIN Wait1
    TcpConnState_fin_wait1 TcpConnState = "fin-wait1"

    // Closing
    TcpConnState_closing TcpConnState = "closing"

    // Last ack
    TcpConnState_last_ack TcpConnState = "last-ack"

    // FIN Wait2
    TcpConnState_fin_wait2 TcpConnState = "fin-wait2"

    // Time wait
    TcpConnState_time_wait TcpConnState = "time-wait"
)

// PakPrio represents Packet Priority Types
type PakPrio string

const (
    // Unspecified
    PakPrio_unspecified_packet PakPrio = "unspecified-packet"

    // Normal: all traffic routed via this router,
    // Telnet/FTP traffic generated from within this
    // router
    PakPrio_normal_packet PakPrio = "normal-packet"

    // Medium: Packets with low drop probability e.g.
    // Routing updates & requests
    PakPrio_medium_packet PakPrio = "medium-packet"

    // High: Packets with very low drop probability
    // and normal delivery e.g. L3 Keepalives like
    // OSPF/ISIS Hellos
    PakPrio_high_packet PakPrio = "high-packet"

    // Crucial: Packets with very low drop probability
    // and expedited delivery e.g L2 keepalives, HDLC
    // Keepalives
    PakPrio_crucial_packet PakPrio = "crucial-packet"
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
    // Internetwork: UDP, TCP, etc.
    AddrFamily_internetwork AddrFamily = "internetwork"

    // IP version 6
    AddrFamily_ip_version6 AddrFamily = "ip-version6"
)

// NsrStatus represents NSR Stream Status
type NsrStatus string

const (
    // NSR Stream Down
    NsrStatus_down NsrStatus = "down"

    // NSR Stream Up
    NsrStatus_up NsrStatus = "up"

    // NSR Stream Not applicable
    NsrStatus_na NsrStatus = "na"
)

// TcpAddressFamily represents Address Family Type
type TcpAddressFamily string

const (
    // IPv4
    TcpAddressFamily_ipv4 TcpAddressFamily = "ipv4"

    // IPv6
    TcpAddressFamily_ipv6 TcpAddressFamily = "ipv6"
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

// NsrDownReason represents NSR-Down Reasons
type NsrDownReason string

const (
    // None, i.e. NSR was never up
    NsrDownReason_none NsrDownReason = "none"

    // Initial sync was aborted
    NsrDownReason_init_sync_aborted NsrDownReason = "init-sync-aborted"

    // Disabled by Active APP
    NsrDownReason_client_disabled NsrDownReason = "client-disabled"

    // Standby APP disconnected
    NsrDownReason_client_disconnect NsrDownReason = "client-disconnect"

    // Standby TCP disconnected
    NsrDownReason_tcp_disconnect NsrDownReason = "tcp-disconnect"

    // RP/DRP Failover occurred
    NsrDownReason_failover NsrDownReason = "failover"

    // Clear nsr command
    NsrDownReason_nsr_clear NsrDownReason = "nsr-clear"

    // Internal error occurred
    NsrDownReason_internal_error NsrDownReason = "internal-error"

    // Retransmission threshold exceededprobably
    // becauseS-TCP was not healthy
    NsrDownReason_retransmit_threshold_exceed NsrDownReason = "retransmit-threshold-exceed"

    // Init-sync repeat failures have exceeded
    // threshold
    NsrDownReason_init_sync_failure_thresh_exceeded NsrDownReason = "init-sync-failure-thresh-exceeded"

    // Audit operation timed out
    NsrDownReason_audit_timeout NsrDownReason = "audit-timeout"

    // Audit operation failed
    NsrDownReason_audit_failed NsrDownReason = "audit-failed"

    // Standby SSCB deleted
    NsrDownReason_standby_sscb_deleted NsrDownReason = "standby-sscb-deleted"

    // Session was closed on standby
    NsrDownReason_standby_session_close NsrDownReason = "standby-session-close"

    // RX-Path was frozen on standby
    NsrDownReason_standby_rxpath_frozen NsrDownReason = "standby-rxpath-frozen"

    // Partner was deleted from set
    NsrDownReason_partner_deleted NsrDownReason = "partner-deleted"
)

// TcpTimer represents TCP Timer Type
type TcpTimer string

const (
    // Retransmission timer
    TcpTimer_retransmission_timer TcpTimer = "retransmission-timer"

    // Send Window Probe timer
    TcpTimer_window_probe_timer TcpTimer = "window-probe-timer"

    // TIMEWAIT state timer
    TcpTimer_timewait_state_timer TcpTimer = "timewait-state-timer"

    // ACK Hold timer
    TcpTimer_ack_hold_timer TcpTimer = "ack-hold-timer"

    // Keep Alive timer
    TcpTimer_keep_alive_timer TcpTimer = "keep-alive-timer"

    // PMTU Ager Timer
    TcpTimer_pmtu_ager_timer TcpTimer = "pmtu-ager-timer"

    // Retransmission Giveup timer
    TcpTimer_retransmission_giveup_timer TcpTimer = "retransmission-giveup-timer"

    // Throttle (for PAW/xipc) timer
    TcpTimer_throttle_timer TcpTimer = "throttle-timer"
)

// TcpConnection
// TCP connection operational data
type TcpConnection struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of information about all nodes present on the system.
    Nodes TcpConnection_Nodes
}

func (tcpConnection *TcpConnection) GetFilter() yfilter.YFilter { return tcpConnection.YFilter }

func (tcpConnection *TcpConnection) SetFilter(yf yfilter.YFilter) { tcpConnection.YFilter = yf }

func (tcpConnection *TcpConnection) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (tcpConnection *TcpConnection) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-tcp-oper:tcp-connection"
}

func (tcpConnection *TcpConnection) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &tcpConnection.Nodes
    }
    return nil
}

func (tcpConnection *TcpConnection) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &tcpConnection.Nodes
    return children
}

func (tcpConnection *TcpConnection) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcpConnection *TcpConnection) GetBundleName() string { return "cisco_ios_xr" }

func (tcpConnection *TcpConnection) GetYangName() string { return "tcp-connection" }

func (tcpConnection *TcpConnection) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcpConnection *TcpConnection) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcpConnection *TcpConnection) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcpConnection *TcpConnection) SetParent(parent types.Entity) { tcpConnection.parent = parent }

func (tcpConnection *TcpConnection) GetParent() types.Entity { return tcpConnection.parent }

func (tcpConnection *TcpConnection) GetParentYangName() string { return "Cisco-IOS-XR-ip-tcp-oper" }

// TcpConnection_Nodes
// Table of information about all nodes present on
// the system
type TcpConnection_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single node. The type is slice of
    // TcpConnection_Nodes_Node.
    Node []TcpConnection_Nodes_Node
}

func (nodes *TcpConnection_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *TcpConnection_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *TcpConnection_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *TcpConnection_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *TcpConnection_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *TcpConnection_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *TcpConnection_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *TcpConnection_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *TcpConnection_Nodes) GetYangName() string { return "nodes" }

func (nodes *TcpConnection_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *TcpConnection_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *TcpConnection_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *TcpConnection_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *TcpConnection_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *TcpConnection_Nodes) GetParentYangName() string { return "tcp-connection" }

// TcpConnection_Nodes_Node
// Information about a single node
type TcpConnection_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Describing a location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Statistics of all TCP connections.
    Statistics TcpConnection_Nodes_Node_Statistics

    // Extended Filter related Information.
    ExtendedInformation TcpConnection_Nodes_Node_ExtendedInformation

    // Table listing TCP connections for which detailed information is provided.
    DetailInformations TcpConnection_Nodes_Node_DetailInformations

    // Table listing connections for which brief information is provided.Note that
    // not all connections are listed in the brief table.
    BriefInformations TcpConnection_Nodes_Node_BriefInformations
}

func (node *TcpConnection_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *TcpConnection_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *TcpConnection_Nodes_Node) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "statistics" { return "Statistics" }
    if yname == "extended-information" { return "ExtendedInformation" }
    if yname == "detail-informations" { return "DetailInformations" }
    if yname == "brief-informations" { return "BriefInformations" }
    return ""
}

func (node *TcpConnection_Nodes_Node) GetSegmentPath() string {
    return "node" + "[id='" + fmt.Sprintf("%v", node.Id) + "']"
}

func (node *TcpConnection_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    if childYangName == "extended-information" {
        return &node.ExtendedInformation
    }
    if childYangName == "detail-informations" {
        return &node.DetailInformations
    }
    if childYangName == "brief-informations" {
        return &node.BriefInformations
    }
    return nil
}

func (node *TcpConnection_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    children["extended-information"] = &node.ExtendedInformation
    children["detail-informations"] = &node.DetailInformations
    children["brief-informations"] = &node.BriefInformations
    return children
}

func (node *TcpConnection_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = node.Id
    return leafs
}

func (node *TcpConnection_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *TcpConnection_Nodes_Node) GetYangName() string { return "node" }

func (node *TcpConnection_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *TcpConnection_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *TcpConnection_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *TcpConnection_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *TcpConnection_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *TcpConnection_Nodes_Node) GetParentYangName() string { return "nodes" }

// TcpConnection_Nodes_Node_Statistics
// Statistics of all TCP connections
type TcpConnection_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table listing clients.
    Clients TcpConnection_Nodes_Node_Statistics_Clients

    // Table listing the TCP connections for which statistics are provided.
    Pcbs TcpConnection_Nodes_Node_Statistics_Pcbs

    // Summary statistics across all TCP connections.
    Summary TcpConnection_Nodes_Node_Statistics_Summary
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *TcpConnection_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "clients" { return "Clients" }
    if yname == "pcbs" { return "Pcbs" }
    if yname == "summary" { return "Summary" }
    return ""
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "clients" {
        return &statistics.Clients
    }
    if childYangName == "pcbs" {
        return &statistics.Pcbs
    }
    if childYangName == "summary" {
        return &statistics.Summary
    }
    return nil
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["clients"] = &statistics.Clients
    children["pcbs"] = &statistics.Pcbs
    children["summary"] = &statistics.Summary
    return children
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *TcpConnection_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *TcpConnection_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// TcpConnection_Nodes_Node_Statistics_Clients
// Table listing clients
type TcpConnection_Nodes_Node_Statistics_Clients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing Client ID. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Clients_Client.
    Client []TcpConnection_Nodes_Node_Statistics_Clients_Client
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetFilter() yfilter.YFilter { return clients.YFilter }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) SetFilter(yf yfilter.YFilter) { clients.YFilter = yf }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetGoName(yname string) string {
    if yname == "client" { return "Client" }
    return ""
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetSegmentPath() string {
    return "clients"
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "client" {
        for _, c := range clients.Client {
            if clients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_Statistics_Clients_Client{}
        clients.Client = append(clients.Client, child)
        return &clients.Client[len(clients.Client)-1]
    }
    return nil
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range clients.Client {
        children[clients.Client[i].GetSegmentPath()] = &clients.Client[i]
    }
    return children
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetBundleName() string { return "cisco_ios_xr" }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetYangName() string { return "clients" }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) SetParent(parent types.Entity) { clients.parent = parent }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetParent() types.Entity { return clients.parent }

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetParentYangName() string { return "statistics" }

// TcpConnection_Nodes_Node_Statistics_Clients_Client
// Describing Client ID
type TcpConnection_Nodes_Node_Statistics_Clients_Client struct {
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

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetGoName(yname string) string {
    if yname == "client-id" { return "ClientId" }
    if yname == "client-jid" { return "ClientJid" }
    if yname == "client-name" { return "ClientName" }
    if yname == "ipv4-received-packets" { return "Ipv4ReceivedPackets" }
    if yname == "ipv4-sent-packets" { return "Ipv4SentPackets" }
    if yname == "ipv6-received-packets" { return "Ipv6ReceivedPackets" }
    if yname == "ipv6-sent-packets" { return "Ipv6SentPackets" }
    return ""
}

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetSegmentPath() string {
    return "client" + "[client-id='" + fmt.Sprintf("%v", client.ClientId) + "']"
}

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetLeafs() map[string]interface{} {
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

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetYangName() string { return "client" }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetParent() types.Entity { return client.parent }

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetParentYangName() string { return "clients" }

// TcpConnection_Nodes_Node_Statistics_Pcbs
// Table listing the TCP connections for which
// statistics are provided
type TcpConnection_Nodes_Node_Statistics_Pcbs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol Control Block ID. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb.
    Pcb []TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetFilter() yfilter.YFilter { return pcbs.YFilter }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) SetFilter(yf yfilter.YFilter) { pcbs.YFilter = yf }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetGoName(yname string) string {
    if yname == "pcb" { return "Pcb" }
    return ""
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetSegmentPath() string {
    return "pcbs"
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pcb" {
        for _, c := range pcbs.Pcb {
            if pcbs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb{}
        pcbs.Pcb = append(pcbs.Pcb, child)
        return &pcbs.Pcb[len(pcbs.Pcb)-1]
    }
    return nil
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range pcbs.Pcb {
        children[pcbs.Pcb[i].GetSegmentPath()] = &pcbs.Pcb[i]
    }
    return children
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetBundleName() string { return "cisco_ios_xr" }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetYangName() string { return "pcbs" }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) SetParent(parent types.Entity) { pcbs.parent = parent }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetParent() types.Entity { return pcbs.parent }

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetParentYangName() string { return "statistics" }

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb
// Protocol Control Block ID
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Displaying statistics associated with a particular
    // PCB. The type is interface{} with range: 0..4294967295.
    Id interface{}

    // PCB Address. The type is interface{} with range: 0..18446744073709551615.
    Pcb interface{}

    // VRF Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Packets received from application. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsSent interface{}

    // XIPC pulses received from application. The type is interface{} with range:
    // 0..18446744073709551615.
    XipcPulseReceived interface{}

    // Segment Instruction received from partner node. The type is interface{}
    // with range: 0..4294967295.
    SegmentInstructionReceived interface{}

    // Packets queued to v4/v6 IO. The type is interface{} with range:
    // 0..18446744073709551615.
    SendPacketsQueued interface{}

    // Packets queued to NetIO. The type is interface{} with range:
    // 0..18446744073709551615.
    SendPacketsQueuedNetIo interface{}

    // Packets failed to be queued to v4/v6 IO. The type is interface{} with
    // range: 0..4294967295.
    SendQueueFailed interface{}

    // Packets failed to be queued to NetIO. The type is interface{} with range:
    // 0..4294967295.
    SendQueueNetIoFailed interface{}

    // Packets received from network. The type is interface{} with range:
    // 0..18446744073709551615.
    PacketsReceived interface{}

    // Received packets failed to be queued to application. The type is
    // interface{} with range: 0..4294967295.
    ReceiveQueueFailed interface{}

    // Received packets queued to application. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedPacketsQueued interface{}

    // No. of times send window shrinkage by peer was ignored. The type is
    // interface{} with range: 0..4294967295.
    SendWindowShrinkIgnored interface{}

    // PAW or non-PAW socket?. The type is bool.
    IsPawSocket interface{}

    // Time at which receive buffer was last read from. The type is interface{}
    // with range: 0..4294967295.
    ReadIoTime interface{}

    // Time at which send buffer was last written to. The type is interface{} with
    // range: 0..4294967295.
    WriteIoTime interface{}

    // Read  I/O counts.
    ReadIoCounts TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts

    // Write I/O counts.
    WriteIoCounts TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts

    // Statistics of Async TCP Sessions.
    AsyncSessionStats TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats
}

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetFilter() yfilter.YFilter { return pcb.YFilter }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) SetFilter(yf yfilter.YFilter) { pcb.YFilter = yf }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "pcb" { return "Pcb" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "packets-sent" { return "PacketsSent" }
    if yname == "xipc-pulse-received" { return "XipcPulseReceived" }
    if yname == "segment-instruction-received" { return "SegmentInstructionReceived" }
    if yname == "send-packets-queued" { return "SendPacketsQueued" }
    if yname == "send-packets-queued-net-io" { return "SendPacketsQueuedNetIo" }
    if yname == "send-queue-failed" { return "SendQueueFailed" }
    if yname == "send-queue-net-io-failed" { return "SendQueueNetIoFailed" }
    if yname == "packets-received" { return "PacketsReceived" }
    if yname == "receive-queue-failed" { return "ReceiveQueueFailed" }
    if yname == "received-packets-queued" { return "ReceivedPacketsQueued" }
    if yname == "send-window-shrink-ignored" { return "SendWindowShrinkIgnored" }
    if yname == "is-paw-socket" { return "IsPawSocket" }
    if yname == "read-io-time" { return "ReadIoTime" }
    if yname == "write-io-time" { return "WriteIoTime" }
    if yname == "read-io-counts" { return "ReadIoCounts" }
    if yname == "write-io-counts" { return "WriteIoCounts" }
    if yname == "async-session-stats" { return "AsyncSessionStats" }
    return ""
}

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetSegmentPath() string {
    return "pcb" + "[id='" + fmt.Sprintf("%v", pcb.Id) + "']"
}

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "read-io-counts" {
        return &pcb.ReadIoCounts
    }
    if childYangName == "write-io-counts" {
        return &pcb.WriteIoCounts
    }
    if childYangName == "async-session-stats" {
        return &pcb.AsyncSessionStats
    }
    return nil
}

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["read-io-counts"] = &pcb.ReadIoCounts
    children["write-io-counts"] = &pcb.WriteIoCounts
    children["async-session-stats"] = &pcb.AsyncSessionStats
    return children
}

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = pcb.Id
    leafs["pcb"] = pcb.Pcb
    leafs["vrf-id"] = pcb.VrfId
    leafs["packets-sent"] = pcb.PacketsSent
    leafs["xipc-pulse-received"] = pcb.XipcPulseReceived
    leafs["segment-instruction-received"] = pcb.SegmentInstructionReceived
    leafs["send-packets-queued"] = pcb.SendPacketsQueued
    leafs["send-packets-queued-net-io"] = pcb.SendPacketsQueuedNetIo
    leafs["send-queue-failed"] = pcb.SendQueueFailed
    leafs["send-queue-net-io-failed"] = pcb.SendQueueNetIoFailed
    leafs["packets-received"] = pcb.PacketsReceived
    leafs["receive-queue-failed"] = pcb.ReceiveQueueFailed
    leafs["received-packets-queued"] = pcb.ReceivedPacketsQueued
    leafs["send-window-shrink-ignored"] = pcb.SendWindowShrinkIgnored
    leafs["is-paw-socket"] = pcb.IsPawSocket
    leafs["read-io-time"] = pcb.ReadIoTime
    leafs["write-io-time"] = pcb.WriteIoTime
    return leafs
}

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetBundleName() string { return "cisco_ios_xr" }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetYangName() string { return "pcb" }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) SetParent(parent types.Entity) { pcb.parent = parent }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetParent() types.Entity { return pcb.parent }

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetParentYangName() string { return "pcbs" }

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts
// Read  I/O counts
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of I/O operations done by application. The type is interface{} with
    // range: 0..4294967295.
    IoCount interface{}

    // How many times socket was armed by application. The type is interface{}
    // with range: 0..4294967295.
    ArmCount interface{}

    // How many times socket was unarmed by application. The type is interface{}
    // with range: 0..4294967295.
    UnarmCount interface{}

    // How many times socket was auto-armed by TCP. The type is interface{} with
    // range: 0..4294967295.
    AutoarmCount interface{}
}

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetFilter() yfilter.YFilter { return readIoCounts.YFilter }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) SetFilter(yf yfilter.YFilter) { readIoCounts.YFilter = yf }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetGoName(yname string) string {
    if yname == "io-count" { return "IoCount" }
    if yname == "arm-count" { return "ArmCount" }
    if yname == "unarm-count" { return "UnarmCount" }
    if yname == "autoarm-count" { return "AutoarmCount" }
    return ""
}

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetSegmentPath() string {
    return "read-io-counts"
}

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["io-count"] = readIoCounts.IoCount
    leafs["arm-count"] = readIoCounts.ArmCount
    leafs["unarm-count"] = readIoCounts.UnarmCount
    leafs["autoarm-count"] = readIoCounts.AutoarmCount
    return leafs
}

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetBundleName() string { return "cisco_ios_xr" }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetYangName() string { return "read-io-counts" }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) SetParent(parent types.Entity) { readIoCounts.parent = parent }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetParent() types.Entity { return readIoCounts.parent }

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetParentYangName() string { return "pcb" }

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts
// Write I/O counts
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of I/O operations done by application. The type is interface{} with
    // range: 0..4294967295.
    IoCount interface{}

    // How many times socket was armed by application. The type is interface{}
    // with range: 0..4294967295.
    ArmCount interface{}

    // How many times socket was unarmed by application. The type is interface{}
    // with range: 0..4294967295.
    UnarmCount interface{}

    // How many times socket was auto-armed by TCP. The type is interface{} with
    // range: 0..4294967295.
    AutoarmCount interface{}
}

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetFilter() yfilter.YFilter { return writeIoCounts.YFilter }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) SetFilter(yf yfilter.YFilter) { writeIoCounts.YFilter = yf }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetGoName(yname string) string {
    if yname == "io-count" { return "IoCount" }
    if yname == "arm-count" { return "ArmCount" }
    if yname == "unarm-count" { return "UnarmCount" }
    if yname == "autoarm-count" { return "AutoarmCount" }
    return ""
}

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetSegmentPath() string {
    return "write-io-counts"
}

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["io-count"] = writeIoCounts.IoCount
    leafs["arm-count"] = writeIoCounts.ArmCount
    leafs["unarm-count"] = writeIoCounts.UnarmCount
    leafs["autoarm-count"] = writeIoCounts.AutoarmCount
    return leafs
}

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetBundleName() string { return "cisco_ios_xr" }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetYangName() string { return "write-io-counts" }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) SetParent(parent types.Entity) { writeIoCounts.parent = parent }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetParent() types.Entity { return writeIoCounts.parent }

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetParentYangName() string { return "pcb" }

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats
// Statistics of Async TCP Sessions
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Flag of async session. The type is bool.
    AsyncSession interface{}

    // Number of successful data write to XIPC. The type is slice of interface{}
    // with range: 0..4294967295.
    DataWriteSuccessNum []interface{}

    // Number of successful data read from XIPC. The type is slice of interface{}
    // with range: 0..4294967295.
    DataReadSuccessNum []interface{}

    // Number of failed data write to XIPC. The type is slice of interface{} with
    // range: 0..4294967295.
    DataWriteErrorNum []interface{}

    // Number of failed data read from XIPC. The type is slice of interface{} with
    // range: 0..4294967295.
    DataReadErrorNum []interface{}

    // Number of successful control write to XIPC. The type is slice of
    // interface{} with range: 0..4294967295.
    ControlWriteSuccessNum []interface{}

    // Number of successful control read to XIPC. The type is slice of interface{}
    // with range: 0..4294967295.
    ControlReadSuccessNum []interface{}

    // Number of failed control write to XIPC. The type is slice of interface{}
    // with range: 0..4294967295.
    ControlWriteErrorNum []interface{}

    // Number of failed control read from XIPC. The type is slice of interface{}
    // with range: 0..4294967295.
    ControlReadErrorNum []interface{}

    // Number of bytes data has been written. The type is slice of interface{}
    // with range: 0..18446744073709551615. Units are byte.
    DataWriteByte []interface{}

    // Number of bytes data has been read. The type is slice of interface{} with
    // range: 0..18446744073709551615. Units are byte.
    DataReadByte []interface{}
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetFilter() yfilter.YFilter { return asyncSessionStats.YFilter }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) SetFilter(yf yfilter.YFilter) { asyncSessionStats.YFilter = yf }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetGoName(yname string) string {
    if yname == "async-session" { return "AsyncSession" }
    if yname == "data-write-success-num" { return "DataWriteSuccessNum" }
    if yname == "data-read-success-num" { return "DataReadSuccessNum" }
    if yname == "data-write-error-num" { return "DataWriteErrorNum" }
    if yname == "data-read-error-num" { return "DataReadErrorNum" }
    if yname == "control-write-success-num" { return "ControlWriteSuccessNum" }
    if yname == "control-read-success-num" { return "ControlReadSuccessNum" }
    if yname == "control-write-error-num" { return "ControlWriteErrorNum" }
    if yname == "control-read-error-num" { return "ControlReadErrorNum" }
    if yname == "data-write-byte" { return "DataWriteByte" }
    if yname == "data-read-byte" { return "DataReadByte" }
    return ""
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetSegmentPath() string {
    return "async-session-stats"
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["async-session"] = asyncSessionStats.AsyncSession
    leafs["data-write-success-num"] = asyncSessionStats.DataWriteSuccessNum
    leafs["data-read-success-num"] = asyncSessionStats.DataReadSuccessNum
    leafs["data-write-error-num"] = asyncSessionStats.DataWriteErrorNum
    leafs["data-read-error-num"] = asyncSessionStats.DataReadErrorNum
    leafs["control-write-success-num"] = asyncSessionStats.ControlWriteSuccessNum
    leafs["control-read-success-num"] = asyncSessionStats.ControlReadSuccessNum
    leafs["control-write-error-num"] = asyncSessionStats.ControlWriteErrorNum
    leafs["control-read-error-num"] = asyncSessionStats.ControlReadErrorNum
    leafs["data-write-byte"] = asyncSessionStats.DataWriteByte
    leafs["data-read-byte"] = asyncSessionStats.DataReadByte
    return leafs
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetBundleName() string { return "cisco_ios_xr" }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetYangName() string { return "async-session-stats" }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) SetParent(parent types.Entity) { asyncSessionStats.parent = parent }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetParent() types.Entity { return asyncSessionStats.parent }

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetParentYangName() string { return "pcb" }

// TcpConnection_Nodes_Node_Statistics_Summary
// Summary statistics across all TCP connections
type TcpConnection_Nodes_Node_Statistics_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Current number of SYN cache entries. The type is interface{} with range:
    // 0..4294967295.
    SynCacheCount interface{}

    // Number of Open sockets. The type is interface{} with range: 0..4294967295.
    NumOpenSockets interface{}

    // Total packets sent. The type is interface{} with range: 0..4294967295.
    TotalPaketsSent interface{}

    // Total transmit packets dropped due to general failures. The type is
    // interface{} with range: 0..4294967295.
    SendPacketsDropped interface{}

    // Total transmit packets dropped due to authentication failures. The type is
    // interface{} with range: 0..4294967295.
    SendAuthPacketsDropped interface{}

    // Data packets sent. The type is interface{} with range: 0..4294967295.
    DataPaketsSent interface{}

    // Data bytes sent. The type is interface{} with range: 0..4294967295. Units
    // are byte.
    DataBytesSent interface{}

    // Data packets retransmitted. The type is interface{} with range:
    // 0..4294967295.
    PacketsRetransmitted interface{}

    // Data bytes retransmitted. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    BytesRetransmitted interface{}

    // Ack only packets sent (incl. delay). The type is interface{} with range:
    // 0..4294967295.
    AckOnlyPacketsSent interface{}

    // Delay ack packets sent. The type is interface{} with range: 0..4294967295.
    DelayAckPacketsSent interface{}

    // Urgent only packets sent. The type is interface{} with range:
    // 0..4294967295.
    UrgentOnlyPacketsSent interface{}

    // Window probe packets sent. The type is interface{} with range:
    // 0..4294967295.
    WindowProbePacketsSent interface{}

    // Window update packets sent. The type is interface{} with range:
    // 0..4294967295.
    WindowUpdatePacketsSent interface{}

    // Control (SYN|FIN|RST) packets sent. The type is interface{} with range:
    // 0..4294967295.
    ControlPacketsSent interface{}

    // RST packets sent. The type is interface{} with range: 0..4294967295.
    RstPacketsSent interface{}

    // Total packets received. The type is interface{} with range: 0..4294967295.
    TotalPacketsReceived interface{}

    // Received packets dropped due to general failures. The type is interface{}
    // with range: 0..4294967295.
    ReceivedPacketsDropped interface{}

    // Received packets dropped due to ACL DENY on SYN pkts. The type is
    // interface{} with range: 0..4294967295.
    SynaclMatchPktsDropped interface{}

    // Received packets dropped due to stale cached header. The type is
    // interface{} with range: 0..4294967295.
    ReceivedPacketsDroppedStaleCHdr interface{}

    // Received packets dropped due to authentication failures. The type is
    // interface{} with range: 0..4294967295.
    ReceivedAuthPacketsDropped interface{}

    // Ack packets received. The type is interface{} with range: 0..4294967295.
    AckPacketsReceived interface{}

    // Bytes acked by ack packets. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    AckbytesReceived interface{}

    // Duplicate ack packets. The type is interface{} with range: 0..4294967295.
    DuplicatedAckPacketsReceived interface{}

    // Ack packets for unsent data. The type is interface{} with range:
    // 0..4294967295.
    AckPacketsForUnsentReceived interface{}

    // Data packets received in sequence. The type is interface{} with range:
    // 0..4294967295.
    DataPacketsReceivedInSequence interface{}

    // Data bytes received in sequence. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    DataBytesReceivedInSequence interface{}

    // Duplicate packets received. The type is interface{} with range:
    // 0..4294967295.
    DuplicatePacketsReceived interface{}

    // Duplicate bytes received. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    DuplicateBytesReceived interface{}

    // Packets with partial dup data. The type is interface{} with range:
    // 0..4294967295.
    PartialDuplicateAckReceived interface{}

    // Bytes with partial dup data. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    PartialDuplicateBytesReceived interface{}

    // Out-of-order packets received. The type is interface{} with range:
    // 0..4294967295.
    OutOfOrderPacketsReceived interface{}

    // Out-of-order bytes received. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    OutOfOrderBytesReceived interface{}

    // After-window packets received. The type is interface{} with range:
    // 0..4294967295.
    AfterWindowPacketsReceived interface{}

    // After-window bytes received. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    AfterWindowBytesReceived interface{}

    // Window probe packets received. The type is interface{} with range:
    // 0..4294967295.
    WindowProbePacketsReceived interface{}

    // Window update packets received. The type is interface{} with range:
    // 0..4294967295.
    WindowUpdatePacketsReceived interface{}

    // Packets received after close. The type is interface{} with range:
    // 0..4294967295.
    PacketsReceivedAfterClosePacket interface{}

    // Packets received with bad checksum. The type is interface{} with range:
    // 0..4294967295.
    BadChecksumPacketsReceived interface{}

    // Packets received with too short size. The type is interface{} with range:
    // 0..4294967295.
    TooShortPacketsReceived interface{}

    // Packets received with malformed header. The type is interface{} with range:
    // 0..4294967295.
    MalformedPacketsReceived interface{}

    // Packets rcceived with no wild listener. The type is interface{} with range:
    // 0..4294967295.
    NoPortPacketsReceived interface{}

    // Connection requests sent. The type is interface{} with range:
    // 0..4294967295.
    ConnectionsRequested interface{}

    // Connection requests accepted. The type is interface{} with range:
    // 0..4294967295.
    ConnectionsAccepted interface{}

    // Connections established. The type is interface{} with range: 0..4294967295.
    ConnectionsEstablished interface{}

    // Connections forcibly closed. The type is interface{} with range:
    // 0..4294967295.
    ConnectionsForciblyClosed interface{}

    // connections closed (incl. drops). The type is interface{} with range:
    // 0..4294967295.
    ConnectionsClosed interface{}

    // connections dropped. The type is interface{} with range: 0..4294967295.
    ConnectionsDropped interface{}

    // Embryonic connections dropped. The type is interface{} with range:
    // 0..4294967295.
    EmbryonicConnectionDropped interface{}

    // Connections failed. The type is interface{} with range: 0..4294967295.
    ConnectionsFailed interface{}

    // Established connections reset. The type is interface{} with range:
    // 0..4294967295.
    EstablishedConnectionsReset interface{}

    // Retransmit timeouts (incl. data packets). The type is interface{} with
    // range: 0..4294967295.
    RetransmitTimeouts interface{}

    // Connection drops during retransmit timeouts. The type is interface{} with
    // range: 0..4294967295.
    RetransmitDropped interface{}

    // Keepalive timeouts. The type is interface{} with range: 0..4294967295.
    KeepAliveTimeouts interface{}

    // Connection drops due to keepalive timeouts. The type is interface{} with
    // range: 0..4294967295.
    KeepAliveDropped interface{}

    // Keepalive probes sent. The type is interface{} with range: 0..4294967295.
    KeepAliveProbes interface{}

    // Segments dropped due to PAWS. The type is interface{} with range:
    // 0..4294967295.
    PawsDropped interface{}

    // Segments dropped due to window probe. The type is interface{} with range:
    // 0..4294967295.
    PersistDropped interface{}

    // Segments dropped due to trylock fail. The type is interface{} with range:
    // 0..4294967295.
    TryLockDropped interface{}

    // Connections rate-limited. The type is interface{} with range:
    // 0..4294967295.
    ConnectionRateLimited interface{}

    // SYN Cache entries added. The type is interface{} with range: 0..4294967295.
    SynCacheAdded interface{}

    // SYN Cache connections completed. The type is interface{} with range:
    // 0..4294967295.
    SynCacheCompleted interface{}

    // SYN Cache entries timed out. The type is interface{} with range:
    // 0..4294967295.
    SynCacheTimedOut interface{}

    // SYN Cache entries dropped due to overflow. The type is interface{} with
    // range: 0..4294967295.
    SynCacheOverflow interface{}

    // SYN Cache entries dropped due to RST. The type is interface{} with range:
    // 0..4294967295.
    SynCacheReset interface{}

    // SYN Cache entries dropped due to ICMP unreach. The type is interface{} with
    // range: 0..4294967295.
    SynCacheUnreach interface{}

    // SYN Cache entries dropped due to bucket overflow. The type is interface{}
    // with range: 0..4294967295.
    SynCacheBucketOflow interface{}

    // SYN Cache entries aborted (no mem). The type is interface{} with range:
    // 0..4294967295.
    SynCacheAborted interface{}

    // SYN Cache duplicate SYNs received. The type is interface{} with range:
    // 0..4294967295.
    SynCacheDuplicateSyNs interface{}

    // SYN Cache entries dropped (no route/mem). The type is interface{} with
    // range: 0..4294967295.
    SynCacheDropped interface{}

    // Punt (down to ip) failures. The type is interface{} with range:
    // 0..4294967295.
    PulseErrors interface{}

    // Packets owned by the socket layer. The type is interface{} with range:
    // 0..4294967295.
    SocketLayerPackets interface{}

    // Packets owned by TCP reassembly. The type is interface{} with range:
    // 0..4294967295.
    ReassemblyPackets interface{}

    // Packets freed after starvation. The type is interface{} with range:
    // 0..4294967295.
    RecoveredPackets interface{}

    // Packet allocation errors. The type is interface{} with range:
    // 0..4294967295.
    PacketFailures interface{}

    // Number of times MSS was increased. The type is interface{} with range:
    // 0..4294967295.
    MssUp interface{}

    // Number of times MSS was decreased. The type is interface{} with range:
    // 0..4294967295.
    MssDown interface{}

    // Segments truncated due to insufficient Write I/O vectors. The type is
    // interface{} with range: 0..4294967295.
    TruncatedWriteIov interface{}

    // Number of times throttle mode was off. The type is interface{} with range:
    // 0..4294967295.
    NoThrottle interface{}

    // Number of times low water mark throttle was on. The type is interface{}
    // with range: 0..4294967295.
    LowWaterMarkThrottle interface{}

    // Number of times high water mark throttle was on. The type is interface{}
    // with range: 0..4294967295.
    HighWaterMarkThrottle interface{}

    // Number of times a stalled tcp timer was tickled. The type is interface{}
    // with range: 0..4294967295.
    StalledTimerTickleCount interface{}

    // Last timestamp when a stalled tcp timer was tickled. The type is
    // interface{} with range: 0..4294967295.
    StalledTimerTickleTime interface{}

    // Number of write attempts from socket-lib into an IQ. The type is
    // interface{} with range: 0..4294967295.
    IqSockWrites interface{}

    // Number of retried write attempts. The type is interface{} with range:
    // 0..4294967295.
    IqSockRetries interface{}

    // Number of aborted socket-lib writes. The type is interface{} with range:
    // 0..4294967295.
    IqSockAborts interface{}

    // Number of total ingress dropped packets. The type is interface{} with
    // range: 0..4294967295.
    IqIngressDrops interface{}

    // Number of TCP internal queues in use. The type is interface{} with range:
    // 0..4294967295.
    TotalIQs interface{}
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetGoName(yname string) string {
    if yname == "syn-cache-count" { return "SynCacheCount" }
    if yname == "num-open-sockets" { return "NumOpenSockets" }
    if yname == "total-pakets-sent" { return "TotalPaketsSent" }
    if yname == "send-packets-dropped" { return "SendPacketsDropped" }
    if yname == "send-auth-packets-dropped" { return "SendAuthPacketsDropped" }
    if yname == "data-pakets-sent" { return "DataPaketsSent" }
    if yname == "data-bytes-sent" { return "DataBytesSent" }
    if yname == "packets-retransmitted" { return "PacketsRetransmitted" }
    if yname == "bytes-retransmitted" { return "BytesRetransmitted" }
    if yname == "ack-only-packets-sent" { return "AckOnlyPacketsSent" }
    if yname == "delay-ack-packets-sent" { return "DelayAckPacketsSent" }
    if yname == "urgent-only-packets-sent" { return "UrgentOnlyPacketsSent" }
    if yname == "window-probe-packets-sent" { return "WindowProbePacketsSent" }
    if yname == "window-update-packets-sent" { return "WindowUpdatePacketsSent" }
    if yname == "control-packets-sent" { return "ControlPacketsSent" }
    if yname == "rst-packets-sent" { return "RstPacketsSent" }
    if yname == "total-packets-received" { return "TotalPacketsReceived" }
    if yname == "received-packets-dropped" { return "ReceivedPacketsDropped" }
    if yname == "synacl-match-pkts-dropped" { return "SynaclMatchPktsDropped" }
    if yname == "received-packets-dropped-stale-c-hdr" { return "ReceivedPacketsDroppedStaleCHdr" }
    if yname == "received-auth-packets-dropped" { return "ReceivedAuthPacketsDropped" }
    if yname == "ack-packets-received" { return "AckPacketsReceived" }
    if yname == "ackbytes-received" { return "AckbytesReceived" }
    if yname == "duplicated-ack-packets-received" { return "DuplicatedAckPacketsReceived" }
    if yname == "ack-packets-for-unsent-received" { return "AckPacketsForUnsentReceived" }
    if yname == "data-packets-received-in-sequence" { return "DataPacketsReceivedInSequence" }
    if yname == "data-bytes-received-in-sequence" { return "DataBytesReceivedInSequence" }
    if yname == "duplicate-packets-received" { return "DuplicatePacketsReceived" }
    if yname == "duplicate-bytes-received" { return "DuplicateBytesReceived" }
    if yname == "partial-duplicate-ack-received" { return "PartialDuplicateAckReceived" }
    if yname == "partial-duplicate-bytes-received" { return "PartialDuplicateBytesReceived" }
    if yname == "out-of-order-packets-received" { return "OutOfOrderPacketsReceived" }
    if yname == "out-of-order-bytes-received" { return "OutOfOrderBytesReceived" }
    if yname == "after-window-packets-received" { return "AfterWindowPacketsReceived" }
    if yname == "after-window-bytes-received" { return "AfterWindowBytesReceived" }
    if yname == "window-probe-packets-received" { return "WindowProbePacketsReceived" }
    if yname == "window-update-packets-received" { return "WindowUpdatePacketsReceived" }
    if yname == "packets-received-after-close-packet" { return "PacketsReceivedAfterClosePacket" }
    if yname == "bad-checksum-packets-received" { return "BadChecksumPacketsReceived" }
    if yname == "too-short-packets-received" { return "TooShortPacketsReceived" }
    if yname == "malformed-packets-received" { return "MalformedPacketsReceived" }
    if yname == "no-port-packets-received" { return "NoPortPacketsReceived" }
    if yname == "connections-requested" { return "ConnectionsRequested" }
    if yname == "connections-accepted" { return "ConnectionsAccepted" }
    if yname == "connections-established" { return "ConnectionsEstablished" }
    if yname == "connections-forcibly-closed" { return "ConnectionsForciblyClosed" }
    if yname == "connections-closed" { return "ConnectionsClosed" }
    if yname == "connections-dropped" { return "ConnectionsDropped" }
    if yname == "embryonic-connection-dropped" { return "EmbryonicConnectionDropped" }
    if yname == "connections-failed" { return "ConnectionsFailed" }
    if yname == "established-connections-reset" { return "EstablishedConnectionsReset" }
    if yname == "retransmit-timeouts" { return "RetransmitTimeouts" }
    if yname == "retransmit-dropped" { return "RetransmitDropped" }
    if yname == "keep-alive-timeouts" { return "KeepAliveTimeouts" }
    if yname == "keep-alive-dropped" { return "KeepAliveDropped" }
    if yname == "keep-alive-probes" { return "KeepAliveProbes" }
    if yname == "paws-dropped" { return "PawsDropped" }
    if yname == "persist-dropped" { return "PersistDropped" }
    if yname == "try-lock-dropped" { return "TryLockDropped" }
    if yname == "connection-rate-limited" { return "ConnectionRateLimited" }
    if yname == "syn-cache-added" { return "SynCacheAdded" }
    if yname == "syn-cache-completed" { return "SynCacheCompleted" }
    if yname == "syn-cache-timed-out" { return "SynCacheTimedOut" }
    if yname == "syn-cache-overflow" { return "SynCacheOverflow" }
    if yname == "syn-cache-reset" { return "SynCacheReset" }
    if yname == "syn-cache-unreach" { return "SynCacheUnreach" }
    if yname == "syn-cache-bucket-oflow" { return "SynCacheBucketOflow" }
    if yname == "syn-cache-aborted" { return "SynCacheAborted" }
    if yname == "syn-cache-duplicate-sy-ns" { return "SynCacheDuplicateSyNs" }
    if yname == "syn-cache-dropped" { return "SynCacheDropped" }
    if yname == "pulse-errors" { return "PulseErrors" }
    if yname == "socket-layer-packets" { return "SocketLayerPackets" }
    if yname == "reassembly-packets" { return "ReassemblyPackets" }
    if yname == "recovered-packets" { return "RecoveredPackets" }
    if yname == "packet-failures" { return "PacketFailures" }
    if yname == "mss-up" { return "MssUp" }
    if yname == "mss-down" { return "MssDown" }
    if yname == "truncated-write-iov" { return "TruncatedWriteIov" }
    if yname == "no-throttle" { return "NoThrottle" }
    if yname == "low-water-mark-throttle" { return "LowWaterMarkThrottle" }
    if yname == "high-water-mark-throttle" { return "HighWaterMarkThrottle" }
    if yname == "stalled-timer-tickle-count" { return "StalledTimerTickleCount" }
    if yname == "stalled-timer-tickle-time" { return "StalledTimerTickleTime" }
    if yname == "iq-sock-writes" { return "IqSockWrites" }
    if yname == "iq-sock-retries" { return "IqSockRetries" }
    if yname == "iq-sock-aborts" { return "IqSockAborts" }
    if yname == "iq-ingress-drops" { return "IqIngressDrops" }
    if yname == "total-i-qs" { return "TotalIQs" }
    return ""
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["syn-cache-count"] = summary.SynCacheCount
    leafs["num-open-sockets"] = summary.NumOpenSockets
    leafs["total-pakets-sent"] = summary.TotalPaketsSent
    leafs["send-packets-dropped"] = summary.SendPacketsDropped
    leafs["send-auth-packets-dropped"] = summary.SendAuthPacketsDropped
    leafs["data-pakets-sent"] = summary.DataPaketsSent
    leafs["data-bytes-sent"] = summary.DataBytesSent
    leafs["packets-retransmitted"] = summary.PacketsRetransmitted
    leafs["bytes-retransmitted"] = summary.BytesRetransmitted
    leafs["ack-only-packets-sent"] = summary.AckOnlyPacketsSent
    leafs["delay-ack-packets-sent"] = summary.DelayAckPacketsSent
    leafs["urgent-only-packets-sent"] = summary.UrgentOnlyPacketsSent
    leafs["window-probe-packets-sent"] = summary.WindowProbePacketsSent
    leafs["window-update-packets-sent"] = summary.WindowUpdatePacketsSent
    leafs["control-packets-sent"] = summary.ControlPacketsSent
    leafs["rst-packets-sent"] = summary.RstPacketsSent
    leafs["total-packets-received"] = summary.TotalPacketsReceived
    leafs["received-packets-dropped"] = summary.ReceivedPacketsDropped
    leafs["synacl-match-pkts-dropped"] = summary.SynaclMatchPktsDropped
    leafs["received-packets-dropped-stale-c-hdr"] = summary.ReceivedPacketsDroppedStaleCHdr
    leafs["received-auth-packets-dropped"] = summary.ReceivedAuthPacketsDropped
    leafs["ack-packets-received"] = summary.AckPacketsReceived
    leafs["ackbytes-received"] = summary.AckbytesReceived
    leafs["duplicated-ack-packets-received"] = summary.DuplicatedAckPacketsReceived
    leafs["ack-packets-for-unsent-received"] = summary.AckPacketsForUnsentReceived
    leafs["data-packets-received-in-sequence"] = summary.DataPacketsReceivedInSequence
    leafs["data-bytes-received-in-sequence"] = summary.DataBytesReceivedInSequence
    leafs["duplicate-packets-received"] = summary.DuplicatePacketsReceived
    leafs["duplicate-bytes-received"] = summary.DuplicateBytesReceived
    leafs["partial-duplicate-ack-received"] = summary.PartialDuplicateAckReceived
    leafs["partial-duplicate-bytes-received"] = summary.PartialDuplicateBytesReceived
    leafs["out-of-order-packets-received"] = summary.OutOfOrderPacketsReceived
    leafs["out-of-order-bytes-received"] = summary.OutOfOrderBytesReceived
    leafs["after-window-packets-received"] = summary.AfterWindowPacketsReceived
    leafs["after-window-bytes-received"] = summary.AfterWindowBytesReceived
    leafs["window-probe-packets-received"] = summary.WindowProbePacketsReceived
    leafs["window-update-packets-received"] = summary.WindowUpdatePacketsReceived
    leafs["packets-received-after-close-packet"] = summary.PacketsReceivedAfterClosePacket
    leafs["bad-checksum-packets-received"] = summary.BadChecksumPacketsReceived
    leafs["too-short-packets-received"] = summary.TooShortPacketsReceived
    leafs["malformed-packets-received"] = summary.MalformedPacketsReceived
    leafs["no-port-packets-received"] = summary.NoPortPacketsReceived
    leafs["connections-requested"] = summary.ConnectionsRequested
    leafs["connections-accepted"] = summary.ConnectionsAccepted
    leafs["connections-established"] = summary.ConnectionsEstablished
    leafs["connections-forcibly-closed"] = summary.ConnectionsForciblyClosed
    leafs["connections-closed"] = summary.ConnectionsClosed
    leafs["connections-dropped"] = summary.ConnectionsDropped
    leafs["embryonic-connection-dropped"] = summary.EmbryonicConnectionDropped
    leafs["connections-failed"] = summary.ConnectionsFailed
    leafs["established-connections-reset"] = summary.EstablishedConnectionsReset
    leafs["retransmit-timeouts"] = summary.RetransmitTimeouts
    leafs["retransmit-dropped"] = summary.RetransmitDropped
    leafs["keep-alive-timeouts"] = summary.KeepAliveTimeouts
    leafs["keep-alive-dropped"] = summary.KeepAliveDropped
    leafs["keep-alive-probes"] = summary.KeepAliveProbes
    leafs["paws-dropped"] = summary.PawsDropped
    leafs["persist-dropped"] = summary.PersistDropped
    leafs["try-lock-dropped"] = summary.TryLockDropped
    leafs["connection-rate-limited"] = summary.ConnectionRateLimited
    leafs["syn-cache-added"] = summary.SynCacheAdded
    leafs["syn-cache-completed"] = summary.SynCacheCompleted
    leafs["syn-cache-timed-out"] = summary.SynCacheTimedOut
    leafs["syn-cache-overflow"] = summary.SynCacheOverflow
    leafs["syn-cache-reset"] = summary.SynCacheReset
    leafs["syn-cache-unreach"] = summary.SynCacheUnreach
    leafs["syn-cache-bucket-oflow"] = summary.SynCacheBucketOflow
    leafs["syn-cache-aborted"] = summary.SynCacheAborted
    leafs["syn-cache-duplicate-sy-ns"] = summary.SynCacheDuplicateSyNs
    leafs["syn-cache-dropped"] = summary.SynCacheDropped
    leafs["pulse-errors"] = summary.PulseErrors
    leafs["socket-layer-packets"] = summary.SocketLayerPackets
    leafs["reassembly-packets"] = summary.ReassemblyPackets
    leafs["recovered-packets"] = summary.RecoveredPackets
    leafs["packet-failures"] = summary.PacketFailures
    leafs["mss-up"] = summary.MssUp
    leafs["mss-down"] = summary.MssDown
    leafs["truncated-write-iov"] = summary.TruncatedWriteIov
    leafs["no-throttle"] = summary.NoThrottle
    leafs["low-water-mark-throttle"] = summary.LowWaterMarkThrottle
    leafs["high-water-mark-throttle"] = summary.HighWaterMarkThrottle
    leafs["stalled-timer-tickle-count"] = summary.StalledTimerTickleCount
    leafs["stalled-timer-tickle-time"] = summary.StalledTimerTickleTime
    leafs["iq-sock-writes"] = summary.IqSockWrites
    leafs["iq-sock-retries"] = summary.IqSockRetries
    leafs["iq-sock-aborts"] = summary.IqSockAborts
    leafs["iq-ingress-drops"] = summary.IqIngressDrops
    leafs["total-i-qs"] = summary.TotalIQs
    return leafs
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetYangName() string { return "summary" }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetParent() types.Entity { return summary.parent }

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetParentYangName() string { return "statistics" }

// TcpConnection_Nodes_Node_ExtendedInformation
// Extended Filter related Information
type TcpConnection_Nodes_Node_ExtendedInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table listing display types.
    DisplayTypes TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetFilter() yfilter.YFilter { return extendedInformation.YFilter }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) SetFilter(yf yfilter.YFilter) { extendedInformation.YFilter = yf }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetGoName(yname string) string {
    if yname == "display-types" { return "DisplayTypes" }
    return ""
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetSegmentPath() string {
    return "extended-information"
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "display-types" {
        return &extendedInformation.DisplayTypes
    }
    return nil
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["display-types"] = &extendedInformation.DisplayTypes
    return children
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetBundleName() string { return "cisco_ios_xr" }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetYangName() string { return "extended-information" }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) SetParent(parent types.Entity) { extendedInformation.parent = parent }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetParent() types.Entity { return extendedInformation.parent }

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetParentYangName() string { return "node" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes
// Table listing display types
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Describing particular display type. The type is slice of
    // TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType.
    DisplayType []TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetFilter() yfilter.YFilter { return displayTypes.YFilter }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) SetFilter(yf yfilter.YFilter) { displayTypes.YFilter = yf }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetGoName(yname string) string {
    if yname == "display-type" { return "DisplayType" }
    return ""
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetSegmentPath() string {
    return "display-types"
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "display-type" {
        for _, c := range displayTypes.DisplayType {
            if displayTypes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType{}
        displayTypes.DisplayType = append(displayTypes.DisplayType, child)
        return &displayTypes.DisplayType[len(displayTypes.DisplayType)-1]
    }
    return nil
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range displayTypes.DisplayType {
        children[displayTypes.DisplayType[i].GetSegmentPath()] = &displayTypes.DisplayType[i]
    }
    return children
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetBundleName() string { return "cisco_ios_xr" }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetYangName() string { return "display-types" }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) SetParent(parent types.Entity) { displayTypes.parent = parent }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetParent() types.Entity { return displayTypes.parent }

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetParentYangName() string { return "extended-information" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType
// Describing particular display type
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Specifying display type. The type is Show.
    DispType interface{}

    // Describing connection ID. The type is slice of
    // TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId.
    ConnectionId []TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetFilter() yfilter.YFilter { return displayType.YFilter }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) SetFilter(yf yfilter.YFilter) { displayType.YFilter = yf }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetGoName(yname string) string {
    if yname == "disp-type" { return "DispType" }
    if yname == "connection-id" { return "ConnectionId" }
    return ""
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetSegmentPath() string {
    return "display-type" + "[disp-type='" + fmt.Sprintf("%v", displayType.DispType) + "']"
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "connection-id" {
        for _, c := range displayType.ConnectionId {
            if displayType.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId{}
        displayType.ConnectionId = append(displayType.ConnectionId, child)
        return &displayType.ConnectionId[len(displayType.ConnectionId)-1]
    }
    return nil
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range displayType.ConnectionId {
        children[displayType.ConnectionId[i].GetSegmentPath()] = &displayType.ConnectionId[i]
    }
    return children
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disp-type"] = displayType.DispType
    return leafs
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetBundleName() string { return "cisco_ios_xr" }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetYangName() string { return "display-type" }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) SetParent(parent types.Entity) { displayType.parent = parent }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetParent() types.Entity { return displayType.parent }

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetParentYangName() string { return "display-types" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId
// Describing connection ID
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Displaying inforamtion based on selected display
    // type associatedwith a particular PCB. The type is interface{} with range:
    // 0..4294967295.
    PcbId interface{}

    // Layer 4 protocol. The type is interface{} with range: 0..4294967295.
    L4Protocol interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Local IP address.
    LocalAddress TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress

    // Remote IP address.
    ForeignAddress TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress

    // Common PCB information.
    Common TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common
}

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetFilter() yfilter.YFilter { return connectionId.YFilter }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) SetFilter(yf yfilter.YFilter) { connectionId.YFilter = yf }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetGoName(yname string) string {
    if yname == "pcb-id" { return "PcbId" }
    if yname == "l4-protocol" { return "L4Protocol" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    if yname == "common" { return "Common" }
    return ""
}

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetSegmentPath() string {
    return "connection-id" + "[pcb-id='" + fmt.Sprintf("%v", connectionId.PcbId) + "']"
}

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &connectionId.LocalAddress
    }
    if childYangName == "foreign-address" {
        return &connectionId.ForeignAddress
    }
    if childYangName == "common" {
        return &connectionId.Common
    }
    return nil
}

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &connectionId.LocalAddress
    children["foreign-address"] = &connectionId.ForeignAddress
    children["common"] = &connectionId.Common
    return children
}

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-id"] = connectionId.PcbId
    leafs["l4-protocol"] = connectionId.L4Protocol
    leafs["local-port"] = connectionId.LocalPort
    leafs["foreign-port"] = connectionId.ForeignPort
    return leafs
}

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetBundleName() string { return "cisco_ios_xr" }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetYangName() string { return "connection-id" }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) SetParent(parent types.Entity) { connectionId.parent = parent }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetParent() types.Entity { return connectionId.parent }

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetParentYangName() string { return "display-type" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress
// Local IP address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress struct {
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

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetParentYangName() string { return "connection-id" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress
// Remote IP address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress struct {
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

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetFilter() yfilter.YFilter { return foreignAddress.YFilter }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) SetFilter(yf yfilter.YFilter) { foreignAddress.YFilter = yf }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetSegmentPath() string {
    return "foreign-address"
}

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = foreignAddress.AfName
    leafs["ipv4-address"] = foreignAddress.Ipv4Address
    leafs["ipv6-address"] = foreignAddress.Ipv6Address
    return leafs
}

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetBundleName() string { return "cisco_ios_xr" }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetYangName() string { return "foreign-address" }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) SetParent(parent types.Entity) { foreignAddress.parent = parent }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetParent() types.Entity { return foreignAddress.parent }

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetParentYangName() string { return "connection-id" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common
// Common PCB information
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address Family. The type is AddrFamily.
    AfName interface{}

    // LPTS PCB information.
    LptsPcb TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetFilter() yfilter.YFilter { return common.YFilter }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) SetFilter(yf yfilter.YFilter) { common.YFilter = yf }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "lpts-pcb" { return "LptsPcb" }
    return ""
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetSegmentPath() string {
    return "common"
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lpts-pcb" {
        return &common.LptsPcb
    }
    return nil
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lpts-pcb"] = &common.LptsPcb
    return children
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = common.AfName
    return leafs
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetBundleName() string { return "cisco_ios_xr" }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetYangName() string { return "common" }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) SetParent(parent types.Entity) { common.parent = parent }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetParent() types.Entity { return common.parent }

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetParentYangName() string { return "connection-id" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb
// LPTS PCB information
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum TTL. The type is interface{} with range: 0..255.
    Ttl interface{}

    // flow information. The type is interface{} with range: 0..4294967295.
    FlowTypesInfo interface{}

    // Receive options.
    Options TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options

    // LPTS flags.
    LptsFlags TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags

    // AcceptMask.
    AcceptMask TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask

    // Interface Filters. The type is slice of
    // TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter.
    Filter []TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetFilter() yfilter.YFilter { return lptsPcb.YFilter }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) SetFilter(yf yfilter.YFilter) { lptsPcb.YFilter = yf }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetGoName(yname string) string {
    if yname == "ttl" { return "Ttl" }
    if yname == "flow-types-info" { return "FlowTypesInfo" }
    if yname == "options" { return "Options" }
    if yname == "lpts-flags" { return "LptsFlags" }
    if yname == "accept-mask" { return "AcceptMask" }
    if yname == "filter" { return "Filter" }
    return ""
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetSegmentPath() string {
    return "lpts-pcb"
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetChildByName(childYangName string, segmentPath string) types.Entity {
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
        child := TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter{}
        lptsPcb.Filter = append(lptsPcb.Filter, child)
        return &lptsPcb.Filter[len(lptsPcb.Filter)-1]
    }
    return nil
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["options"] = &lptsPcb.Options
    children["lpts-flags"] = &lptsPcb.LptsFlags
    children["accept-mask"] = &lptsPcb.AcceptMask
    for i := range lptsPcb.Filter {
        children[lptsPcb.Filter[i].GetSegmentPath()] = &lptsPcb.Filter[i]
    }
    return children
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ttl"] = lptsPcb.Ttl
    leafs["flow-types-info"] = lptsPcb.FlowTypesInfo
    return leafs
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetBundleName() string { return "cisco_ios_xr" }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetYangName() string { return "lpts-pcb" }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) SetParent(parent types.Entity) { lptsPcb.parent = parent }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetParent() types.Entity { return lptsPcb.parent }

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetParentYangName() string { return "common" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options
// Receive options
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Receive filter enabled. The type is bool.
    IsReceiveFilter interface{}

    // IP SLA. The type is bool.
    IsIpSla interface{}
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetGoName(yname string) string {
    if yname == "is-receive-filter" { return "IsReceiveFilter" }
    if yname == "is-ip-sla" { return "IsIpSla" }
    return ""
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetSegmentPath() string {
    return "options"
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-receive-filter"] = options.IsReceiveFilter
    leafs["is-ip-sla"] = options.IsIpSla
    return leafs
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetYangName() string { return "options" }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetParent() types.Entity { return options.parent }

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetParentYangName() string { return "lpts-pcb" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags
// LPTS flags
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // PCB bound. The type is bool.
    IsPcbBound interface{}

    // Sent drop packets. The type is bool.
    IsLocalAddressIgnore interface{}

    // Ignore VRF Filter. The type is bool.
    IsIgnoreVrfFilter interface{}
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetFilter() yfilter.YFilter { return lptsFlags.YFilter }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) SetFilter(yf yfilter.YFilter) { lptsFlags.YFilter = yf }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetGoName(yname string) string {
    if yname == "is-pcb-bound" { return "IsPcbBound" }
    if yname == "is-local-address-ignore" { return "IsLocalAddressIgnore" }
    if yname == "is-ignore-vrf-filter" { return "IsIgnoreVrfFilter" }
    return ""
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetSegmentPath() string {
    return "lpts-flags"
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-pcb-bound"] = lptsFlags.IsPcbBound
    leafs["is-local-address-ignore"] = lptsFlags.IsLocalAddressIgnore
    leafs["is-ignore-vrf-filter"] = lptsFlags.IsIgnoreVrfFilter
    return leafs
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetBundleName() string { return "cisco_ios_xr" }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetYangName() string { return "lpts-flags" }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) SetParent(parent types.Entity) { lptsFlags.parent = parent }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetParent() types.Entity { return lptsFlags.parent }

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetParentYangName() string { return "lpts-pcb" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask
// AcceptMask
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask struct {
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

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetFilter() yfilter.YFilter { return acceptMask.YFilter }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) SetFilter(yf yfilter.YFilter) { acceptMask.YFilter = yf }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetGoName(yname string) string {
    if yname == "is-interface" { return "IsInterface" }
    if yname == "is-packet-type" { return "IsPacketType" }
    if yname == "is-remote-address" { return "IsRemoteAddress" }
    if yname == "is-remote-port" { return "IsRemotePort" }
    if yname == "is-local-address" { return "IsLocalAddress" }
    if yname == "is-local-port" { return "IsLocalPort" }
    return ""
}

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetSegmentPath() string {
    return "accept-mask"
}

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["is-interface"] = acceptMask.IsInterface
    leafs["is-packet-type"] = acceptMask.IsPacketType
    leafs["is-remote-address"] = acceptMask.IsRemoteAddress
    leafs["is-remote-port"] = acceptMask.IsRemotePort
    leafs["is-local-address"] = acceptMask.IsLocalAddress
    leafs["is-local-port"] = acceptMask.IsLocalPort
    return leafs
}

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetBundleName() string { return "cisco_ios_xr" }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetYangName() string { return "accept-mask" }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) SetParent(parent types.Entity) { acceptMask.parent = parent }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetParent() types.Entity { return acceptMask.parent }

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetParentYangName() string { return "lpts-pcb" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter
// Interface Filters
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter struct {
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
    PacketType TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType

    // Remote address.
    RemoteAddress TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress

    // Local address.
    LocalAddress TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress
}

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetFilter() yfilter.YFilter { return filter.YFilter }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) SetFilter(yf yfilter.YFilter) { filter.YFilter = yf }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetGoName(yname string) string {
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

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetSegmentPath() string {
    return "filter"
}

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetChildByName(childYangName string, segmentPath string) types.Entity {
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

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["packet-type"] = &filter.PacketType
    children["remote-address"] = &filter.RemoteAddress
    children["local-address"] = &filter.LocalAddress
    return children
}

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetLeafs() map[string]interface{} {
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

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetBundleName() string { return "cisco_ios_xr" }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetYangName() string { return "filter" }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) SetParent(parent types.Entity) { filter.parent = parent }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetParent() types.Entity { return filter.parent }

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetParentYangName() string { return "lpts-pcb" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType
// Protocol-specific packet type
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType struct {
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

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetFilter() yfilter.YFilter { return packetType.YFilter }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) SetFilter(yf yfilter.YFilter) { packetType.YFilter = yf }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "icmp-message-type" { return "IcmpMessageType" }
    if yname == "icm-pv6-message-type" { return "IcmPv6MessageType" }
    if yname == "igmp-message-type" { return "IgmpMessageType" }
    if yname == "message-id" { return "MessageId" }
    return ""
}

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetSegmentPath() string {
    return "packet-type"
}

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = packetType.Type
    leafs["icmp-message-type"] = packetType.IcmpMessageType
    leafs["icm-pv6-message-type"] = packetType.IcmPv6MessageType
    leafs["igmp-message-type"] = packetType.IgmpMessageType
    leafs["message-id"] = packetType.MessageId
    return leafs
}

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetBundleName() string { return "cisco_ios_xr" }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetYangName() string { return "packet-type" }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) SetParent(parent types.Entity) { packetType.parent = parent }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetParent() types.Entity { return packetType.parent }

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetParentYangName() string { return "filter" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress
// Remote address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress struct {
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

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetFilter() yfilter.YFilter { return remoteAddress.YFilter }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) SetFilter(yf yfilter.YFilter) { remoteAddress.YFilter = yf }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetSegmentPath() string {
    return "remote-address"
}

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = remoteAddress.AfName
    leafs["ipv4-address"] = remoteAddress.Ipv4Address
    leafs["ipv6-address"] = remoteAddress.Ipv6Address
    return leafs
}

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetBundleName() string { return "cisco_ios_xr" }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetYangName() string { return "remote-address" }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) SetParent(parent types.Entity) { remoteAddress.parent = parent }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetParent() types.Entity { return remoteAddress.parent }

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetParentYangName() string { return "filter" }

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress
// Local address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress struct {
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

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetParentYangName() string { return "filter" }

// TcpConnection_Nodes_Node_DetailInformations
// Table listing TCP connections for which
// detailed information is provided
type TcpConnection_Nodes_Node_DetailInformations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Protocol Control Block ID. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation.
    DetailInformation []TcpConnection_Nodes_Node_DetailInformations_DetailInformation
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetFilter() yfilter.YFilter { return detailInformations.YFilter }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) SetFilter(yf yfilter.YFilter) { detailInformations.YFilter = yf }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetGoName(yname string) string {
    if yname == "detail-information" { return "DetailInformation" }
    return ""
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetSegmentPath() string {
    return "detail-informations"
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-information" {
        for _, c := range detailInformations.DetailInformation {
            if detailInformations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_DetailInformations_DetailInformation{}
        detailInformations.DetailInformation = append(detailInformations.DetailInformation, child)
        return &detailInformations.DetailInformation[len(detailInformations.DetailInformation)-1]
    }
    return nil
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detailInformations.DetailInformation {
        children[detailInformations.DetailInformation[i].GetSegmentPath()] = &detailInformations.DetailInformation[i]
    }
    return children
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetBundleName() string { return "cisco_ios_xr" }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetYangName() string { return "detail-informations" }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) SetParent(parent types.Entity) { detailInformations.parent = parent }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetParent() types.Entity { return detailInformations.parent }

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetParentYangName() string { return "node" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation
// Protocol Control Block ID
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Detail information about TCP connection, put null
    // for all. The type is interface{} with range: 0..4294967295.
    PcbId interface{}

    // Address Family. The type is TcpAddressFamily.
    AddressFamily interface{}

    // PCB Address. The type is interface{} with range: 0..18446744073709551615.
    Pcb interface{}

    // Socket Address. The type is interface{} with range:
    // 0..18446744073709551615.
    So interface{}

    // TCPCB Address. The type is interface{} with range: 0..18446744073709551615.
    Tcpcb interface{}

    // VRF Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Connection state. The type is TcpConnState.
    ConnectionState interface{}

    // Time at which connection is established. The type is interface{} with
    // range: 0..4294967295.
    EstablishedTime interface{}

    // Id of the local process. The type is interface{} with range: 0..4294967295.
    LocalPid interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Priority given to packets on this socket. The type is PakPrio.
    PacketPriority interface{}

    // Type of Service value to be applied to transmistted packets. The type is
    // interface{} with range: 0..65535.
    PacketTos interface{}

    // TTL to be applied to transmited packets. The type is interface{} with
    // range: 0..65535.
    PacketTtl interface{}

    // Index of the Hash Bucket. The type is interface{} with range:
    // 0..4294967295.
    HashIndex interface{}

    // Current receive queue size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    CurrentReceiveQueueSize interface{}

    // Max receive queue size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    MaxReceiveQueueSize interface{}

    // Current send queue size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    CurrentSendQueueSize interface{}

    // Max send queue size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    MaxSendQueueSize interface{}

    // Current receive queue size in packets. The type is interface{} with range:
    // 0..4294967295.
    CurrentReceiveQueuePacketSize interface{}

    // Max receive queue size in packets. The type is interface{} with range:
    // 0..4294967295.
    MaxReceiveQueuePacketSize interface{}

    // Save queue (out-of seq data) size in bytes. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    SaveQueueSize interface{}

    // Initial send sequence number. The type is interface{} with range:
    // 0..4294967295.
    SendInitialSequenceNum interface{}

    // Sequence number of unacked data. The type is interface{} with range:
    // 0..4294967295.
    SendUnackSequenceNum interface{}

    // Sequence number of next data to be sent. The type is interface{} with
    // range: 0..4294967295.
    SendNextSequenceNum interface{}

    // Highest sequence number sent. The type is interface{} with range:
    // 0..4294967295.
    SendMaxSequenceNum interface{}

    // Send window size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    SendWindowSize interface{}

    // Send congestion window size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    SendCongestionWindowSize interface{}

    // Initial receive sequence number. The type is interface{} with range:
    // 0..4294967295.
    ReceiveInitialSequenceNum interface{}

    // Next sequence number expected. The type is interface{} with range:
    // 0..4294967295.
    ReceiveNextSequenceNum interface{}

    // Receive advertised window size in bytes. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    ReceiveAdvWindowSize interface{}

    // Receive window size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    ReceiveWindowSize interface{}

    // Max segment size calculated in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    Mss interface{}

    // Max segment size offered by the peer in bytes. The type is interface{} with
    // range: 0..4294967295. Units are byte.
    PeerMss interface{}

    // Smoothed round trip time * 8 (msec). The type is interface{} with range:
    // 0..4294967295.
    Srtt interface{}

    // Round trip timeout (msec). The type is interface{} with range:
    // 0..4294967295.
    Rtto interface{}

    // Round trip time (karn algorithm) (msec). The type is interface{} with
    // range: 0..4294967295.
    Krtt interface{}

    // Smoothed round trip time variance * 4 (msec). The type is interface{} with
    // range: 0..4294967295.
    Srtv interface{}

    // Min RTT (msec). The type is interface{} with range: 0..4294967295.
    MinRtt interface{}

    // Max RTT (msec). The type is interface{} with range: 0..4294967295.
    MaxRtt interface{}

    // Number of retries. The type is interface{} with range: 0..4294967295.
    Retries interface{}

    // ACK hold time (msec). The type is interface{} with range: 0..4294967295.
    AckHoldTime interface{}

    // Giveup time (msec). The type is interface{} with range: 0..4294967295.
    GiveupTime interface{}

    // Keepalive time (msec). The type is interface{} with range: 0..4294967295.
    KeepAliveTime interface{}

    // SYN wait time (msec). The type is interface{} with range: 0..4294967295.
    SynWaitTime interface{}

    // RX Syn acl name. The type is string with length: 0..64.
    RxsyNaclname interface{}

    // Error code from ICMP Notify. The type is interface{} with range:
    // -2147483648..2147483647.
    SoftError interface{}

    // Socket error code. The type is interface{} with range:
    // -2147483648..2147483647.
    SockError interface{}

    // Retransimit forever?. The type is bool.
    IsRetransForever interface{}

    // Lowest MSS ever used. The type is interface{} with range: 0..4294967295.
    MinMss interface{}

    // Highest MSS ever used. The type is interface{} with range: 0..4294967295.
    MaxMss interface{}

    // Number of times connect will be retried?. The type is interface{} with
    // range: 0..65535.
    ConnectRetries interface{}

    // Connect retry interval in seconds. The type is interface{} with range:
    // 0..65535. Units are second.
    ConnectRetryInterval interface{}

    // Window scaling for receive window. The type is interface{} with range:
    // 0..4294967295.
    ReceiveWindowScale interface{}

    // Window scaling for send window. The type is interface{} with range:
    // 0..4294967295.
    SendWindowScale interface{}

    // Requested receive window scale. The type is interface{} with range:
    // 0..4294967295.
    RequestReceiveWindowScale interface{}

    // Requested send window scale. The type is interface{} with range:
    // 0..4294967295.
    RqstSendWndScale interface{}

    // Timestamp from remote host. The type is interface{} with range:
    // 0..4294967295.
    TimeStampRecent interface{}

    // Timestamp when last updated. The type is interface{} with range:
    // 0..4294967295.
    TimeStampRecentAge interface{}

    // ACK number of a sent segment. The type is interface{} with range:
    // 0..4294967295.
    LastAckSent interface{}

    // Send buffer's low water mark. The type is interface{} with range:
    // 0..4294967295.
    SendbufLowwat interface{}

    // Receive buffer's low water mark. The type is interface{} with range:
    // 0..4294967295.
    RecvbufLowwat interface{}

    // Send buffer's high water mark. The type is interface{} with range:
    // 0..4294967295.
    SendbufHiwat interface{}

    // Receive buffer's high water mark. The type is interface{} with range:
    // 0..4294967295.
    RecvbufHiwat interface{}

    // Send buffer's notify threshold. The type is interface{} with range:
    // 0..4294967295.
    SendbufNotifyThresh interface{}

    // Receive buffer's data size. The type is interface{} with range:
    // 0..4294967295.
    RecvbufDatasize interface{}

    // Incoming connection queue size. The type is interface{} with range:
    // 0..4294967295.
    QueueLength interface{}

    // Incoming half-connection queue size. The type is interface{} with range:
    // 0..4294967295.
    QueueZeroLength interface{}

    // Incoming connection queue limit. The type is interface{} with range:
    // 0..4294967295.
    QueueLimit interface{}

    // Socket error status. The type is interface{} with range: 0..4294967295.
    SocketError interface{}

    // Socket auto rearm state. The type is interface{} with range: 0..4294967295.
    AutoRearm interface{}

    // # of PDU's in Send Buffer. The type is interface{} with range:
    // 0..4294967295.
    SendPduCount interface{}

    // Cached Outgoing interface  handle. The type is interface{} with range:
    // 0..4294967295.
    OutputIfhandle interface{}

    // Cached fib pd context size. The type is interface{} with range:
    // 0..4294967295.
    FibPdCtxSize interface{}

    // Number of labels returned by fib lookup. The type is interface{} with
    // range: 0..4294967295.
    NumLabels interface{}

    // Instance number of the local process. The type is interface{} with range:
    // 0..4294967295.
    LocalAppInstance interface{}

    // Cached fib pd context. The type is slice of interface{} with range:
    // 0..4294967295.
    FibPdCtx []interface{}

    // Cached Label stack. The type is slice of interface{} with range:
    // 0..4294967295.
    FibLabelOutput []interface{}

    // Local address.
    LocalAddress TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress

    // Foreign address.
    ForeignAddress TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress

    // Socket option flags.
    SocketOptionFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags

    // Socket state flags.
    SocketStateFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags

    // Connection feature flags.
    FeatureFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags

    // Connection state flags.
    StateFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags

    // Connection request flags.
    RequestFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags

    // Receive buffer state flags.
    ReceiveBufStateFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags

    // Send buffer state flags.
    SendBufStateFlags TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags

    // Timers. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer.
    Timer []TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer

    // Seq nos. of sack blocks. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk.
    SackBlk []TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk

    // Sorted list of sack holes. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole.
    SendSackHole []TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetFilter() yfilter.YFilter { return detailInformation.YFilter }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) SetFilter(yf yfilter.YFilter) { detailInformation.YFilter = yf }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetGoName(yname string) string {
    if yname == "pcb-id" { return "PcbId" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "pcb" { return "Pcb" }
    if yname == "so" { return "So" }
    if yname == "tcpcb" { return "Tcpcb" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "connection-state" { return "ConnectionState" }
    if yname == "established-time" { return "EstablishedTime" }
    if yname == "local-pid" { return "LocalPid" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "packet-priority" { return "PacketPriority" }
    if yname == "packet-tos" { return "PacketTos" }
    if yname == "packet-ttl" { return "PacketTtl" }
    if yname == "hash-index" { return "HashIndex" }
    if yname == "current-receive-queue-size" { return "CurrentReceiveQueueSize" }
    if yname == "max-receive-queue-size" { return "MaxReceiveQueueSize" }
    if yname == "current-send-queue-size" { return "CurrentSendQueueSize" }
    if yname == "max-send-queue-size" { return "MaxSendQueueSize" }
    if yname == "current-receive-queue-packet-size" { return "CurrentReceiveQueuePacketSize" }
    if yname == "max-receive-queue-packet-size" { return "MaxReceiveQueuePacketSize" }
    if yname == "save-queue-size" { return "SaveQueueSize" }
    if yname == "send-initial-sequence-num" { return "SendInitialSequenceNum" }
    if yname == "send-unack-sequence-num" { return "SendUnackSequenceNum" }
    if yname == "send-next-sequence-num" { return "SendNextSequenceNum" }
    if yname == "send-max-sequence-num" { return "SendMaxSequenceNum" }
    if yname == "send-window-size" { return "SendWindowSize" }
    if yname == "send-congestion-window-size" { return "SendCongestionWindowSize" }
    if yname == "receive-initial-sequence-num" { return "ReceiveInitialSequenceNum" }
    if yname == "receive-next-sequence-num" { return "ReceiveNextSequenceNum" }
    if yname == "receive-adv-window-size" { return "ReceiveAdvWindowSize" }
    if yname == "receive-window-size" { return "ReceiveWindowSize" }
    if yname == "mss" { return "Mss" }
    if yname == "peer-mss" { return "PeerMss" }
    if yname == "srtt" { return "Srtt" }
    if yname == "rtto" { return "Rtto" }
    if yname == "krtt" { return "Krtt" }
    if yname == "srtv" { return "Srtv" }
    if yname == "min-rtt" { return "MinRtt" }
    if yname == "max-rtt" { return "MaxRtt" }
    if yname == "retries" { return "Retries" }
    if yname == "ack-hold-time" { return "AckHoldTime" }
    if yname == "giveup-time" { return "GiveupTime" }
    if yname == "keep-alive-time" { return "KeepAliveTime" }
    if yname == "syn-wait-time" { return "SynWaitTime" }
    if yname == "rxsy-naclname" { return "RxsyNaclname" }
    if yname == "soft-error" { return "SoftError" }
    if yname == "sock-error" { return "SockError" }
    if yname == "is-retrans-forever" { return "IsRetransForever" }
    if yname == "min-mss" { return "MinMss" }
    if yname == "max-mss" { return "MaxMss" }
    if yname == "connect-retries" { return "ConnectRetries" }
    if yname == "connect-retry-interval" { return "ConnectRetryInterval" }
    if yname == "receive-window-scale" { return "ReceiveWindowScale" }
    if yname == "send-window-scale" { return "SendWindowScale" }
    if yname == "request-receive-window-scale" { return "RequestReceiveWindowScale" }
    if yname == "rqst-send-wnd-scale" { return "RqstSendWndScale" }
    if yname == "time-stamp-recent" { return "TimeStampRecent" }
    if yname == "time-stamp-recent-age" { return "TimeStampRecentAge" }
    if yname == "last-ack-sent" { return "LastAckSent" }
    if yname == "sendbuf-lowwat" { return "SendbufLowwat" }
    if yname == "recvbuf-lowwat" { return "RecvbufLowwat" }
    if yname == "sendbuf-hiwat" { return "SendbufHiwat" }
    if yname == "recvbuf-hiwat" { return "RecvbufHiwat" }
    if yname == "sendbuf-notify-thresh" { return "SendbufNotifyThresh" }
    if yname == "recvbuf-datasize" { return "RecvbufDatasize" }
    if yname == "queue-length" { return "QueueLength" }
    if yname == "queue-zero-length" { return "QueueZeroLength" }
    if yname == "queue-limit" { return "QueueLimit" }
    if yname == "socket-error" { return "SocketError" }
    if yname == "auto-rearm" { return "AutoRearm" }
    if yname == "send-pdu-count" { return "SendPduCount" }
    if yname == "output-ifhandle" { return "OutputIfhandle" }
    if yname == "fib-pd-ctx-size" { return "FibPdCtxSize" }
    if yname == "num-labels" { return "NumLabels" }
    if yname == "local-app-instance" { return "LocalAppInstance" }
    if yname == "fib-pd-ctx" { return "FibPdCtx" }
    if yname == "fib-label-output" { return "FibLabelOutput" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    if yname == "socket-option-flags" { return "SocketOptionFlags" }
    if yname == "socket-state-flags" { return "SocketStateFlags" }
    if yname == "feature-flags" { return "FeatureFlags" }
    if yname == "state-flags" { return "StateFlags" }
    if yname == "request-flags" { return "RequestFlags" }
    if yname == "receive-buf-state-flags" { return "ReceiveBufStateFlags" }
    if yname == "send-buf-state-flags" { return "SendBufStateFlags" }
    if yname == "timer" { return "Timer" }
    if yname == "sack-blk" { return "SackBlk" }
    if yname == "send-sack-hole" { return "SendSackHole" }
    return ""
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetSegmentPath() string {
    return "detail-information" + "[pcb-id='" + fmt.Sprintf("%v", detailInformation.PcbId) + "']"
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &detailInformation.LocalAddress
    }
    if childYangName == "foreign-address" {
        return &detailInformation.ForeignAddress
    }
    if childYangName == "socket-option-flags" {
        return &detailInformation.SocketOptionFlags
    }
    if childYangName == "socket-state-flags" {
        return &detailInformation.SocketStateFlags
    }
    if childYangName == "feature-flags" {
        return &detailInformation.FeatureFlags
    }
    if childYangName == "state-flags" {
        return &detailInformation.StateFlags
    }
    if childYangName == "request-flags" {
        return &detailInformation.RequestFlags
    }
    if childYangName == "receive-buf-state-flags" {
        return &detailInformation.ReceiveBufStateFlags
    }
    if childYangName == "send-buf-state-flags" {
        return &detailInformation.SendBufStateFlags
    }
    if childYangName == "timer" {
        for _, c := range detailInformation.Timer {
            if detailInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer{}
        detailInformation.Timer = append(detailInformation.Timer, child)
        return &detailInformation.Timer[len(detailInformation.Timer)-1]
    }
    if childYangName == "sack-blk" {
        for _, c := range detailInformation.SackBlk {
            if detailInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk{}
        detailInformation.SackBlk = append(detailInformation.SackBlk, child)
        return &detailInformation.SackBlk[len(detailInformation.SackBlk)-1]
    }
    if childYangName == "send-sack-hole" {
        for _, c := range detailInformation.SendSackHole {
            if detailInformation.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole{}
        detailInformation.SendSackHole = append(detailInformation.SendSackHole, child)
        return &detailInformation.SendSackHole[len(detailInformation.SendSackHole)-1]
    }
    return nil
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &detailInformation.LocalAddress
    children["foreign-address"] = &detailInformation.ForeignAddress
    children["socket-option-flags"] = &detailInformation.SocketOptionFlags
    children["socket-state-flags"] = &detailInformation.SocketStateFlags
    children["feature-flags"] = &detailInformation.FeatureFlags
    children["state-flags"] = &detailInformation.StateFlags
    children["request-flags"] = &detailInformation.RequestFlags
    children["receive-buf-state-flags"] = &detailInformation.ReceiveBufStateFlags
    children["send-buf-state-flags"] = &detailInformation.SendBufStateFlags
    for i := range detailInformation.Timer {
        children[detailInformation.Timer[i].GetSegmentPath()] = &detailInformation.Timer[i]
    }
    for i := range detailInformation.SackBlk {
        children[detailInformation.SackBlk[i].GetSegmentPath()] = &detailInformation.SackBlk[i]
    }
    for i := range detailInformation.SendSackHole {
        children[detailInformation.SendSackHole[i].GetSegmentPath()] = &detailInformation.SendSackHole[i]
    }
    return children
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-id"] = detailInformation.PcbId
    leafs["address-family"] = detailInformation.AddressFamily
    leafs["pcb"] = detailInformation.Pcb
    leafs["so"] = detailInformation.So
    leafs["tcpcb"] = detailInformation.Tcpcb
    leafs["vrf-id"] = detailInformation.VrfId
    leafs["connection-state"] = detailInformation.ConnectionState
    leafs["established-time"] = detailInformation.EstablishedTime
    leafs["local-pid"] = detailInformation.LocalPid
    leafs["local-port"] = detailInformation.LocalPort
    leafs["foreign-port"] = detailInformation.ForeignPort
    leafs["packet-priority"] = detailInformation.PacketPriority
    leafs["packet-tos"] = detailInformation.PacketTos
    leafs["packet-ttl"] = detailInformation.PacketTtl
    leafs["hash-index"] = detailInformation.HashIndex
    leafs["current-receive-queue-size"] = detailInformation.CurrentReceiveQueueSize
    leafs["max-receive-queue-size"] = detailInformation.MaxReceiveQueueSize
    leafs["current-send-queue-size"] = detailInformation.CurrentSendQueueSize
    leafs["max-send-queue-size"] = detailInformation.MaxSendQueueSize
    leafs["current-receive-queue-packet-size"] = detailInformation.CurrentReceiveQueuePacketSize
    leafs["max-receive-queue-packet-size"] = detailInformation.MaxReceiveQueuePacketSize
    leafs["save-queue-size"] = detailInformation.SaveQueueSize
    leafs["send-initial-sequence-num"] = detailInformation.SendInitialSequenceNum
    leafs["send-unack-sequence-num"] = detailInformation.SendUnackSequenceNum
    leafs["send-next-sequence-num"] = detailInformation.SendNextSequenceNum
    leafs["send-max-sequence-num"] = detailInformation.SendMaxSequenceNum
    leafs["send-window-size"] = detailInformation.SendWindowSize
    leafs["send-congestion-window-size"] = detailInformation.SendCongestionWindowSize
    leafs["receive-initial-sequence-num"] = detailInformation.ReceiveInitialSequenceNum
    leafs["receive-next-sequence-num"] = detailInformation.ReceiveNextSequenceNum
    leafs["receive-adv-window-size"] = detailInformation.ReceiveAdvWindowSize
    leafs["receive-window-size"] = detailInformation.ReceiveWindowSize
    leafs["mss"] = detailInformation.Mss
    leafs["peer-mss"] = detailInformation.PeerMss
    leafs["srtt"] = detailInformation.Srtt
    leafs["rtto"] = detailInformation.Rtto
    leafs["krtt"] = detailInformation.Krtt
    leafs["srtv"] = detailInformation.Srtv
    leafs["min-rtt"] = detailInformation.MinRtt
    leafs["max-rtt"] = detailInformation.MaxRtt
    leafs["retries"] = detailInformation.Retries
    leafs["ack-hold-time"] = detailInformation.AckHoldTime
    leafs["giveup-time"] = detailInformation.GiveupTime
    leafs["keep-alive-time"] = detailInformation.KeepAliveTime
    leafs["syn-wait-time"] = detailInformation.SynWaitTime
    leafs["rxsy-naclname"] = detailInformation.RxsyNaclname
    leafs["soft-error"] = detailInformation.SoftError
    leafs["sock-error"] = detailInformation.SockError
    leafs["is-retrans-forever"] = detailInformation.IsRetransForever
    leafs["min-mss"] = detailInformation.MinMss
    leafs["max-mss"] = detailInformation.MaxMss
    leafs["connect-retries"] = detailInformation.ConnectRetries
    leafs["connect-retry-interval"] = detailInformation.ConnectRetryInterval
    leafs["receive-window-scale"] = detailInformation.ReceiveWindowScale
    leafs["send-window-scale"] = detailInformation.SendWindowScale
    leafs["request-receive-window-scale"] = detailInformation.RequestReceiveWindowScale
    leafs["rqst-send-wnd-scale"] = detailInformation.RqstSendWndScale
    leafs["time-stamp-recent"] = detailInformation.TimeStampRecent
    leafs["time-stamp-recent-age"] = detailInformation.TimeStampRecentAge
    leafs["last-ack-sent"] = detailInformation.LastAckSent
    leafs["sendbuf-lowwat"] = detailInformation.SendbufLowwat
    leafs["recvbuf-lowwat"] = detailInformation.RecvbufLowwat
    leafs["sendbuf-hiwat"] = detailInformation.SendbufHiwat
    leafs["recvbuf-hiwat"] = detailInformation.RecvbufHiwat
    leafs["sendbuf-notify-thresh"] = detailInformation.SendbufNotifyThresh
    leafs["recvbuf-datasize"] = detailInformation.RecvbufDatasize
    leafs["queue-length"] = detailInformation.QueueLength
    leafs["queue-zero-length"] = detailInformation.QueueZeroLength
    leafs["queue-limit"] = detailInformation.QueueLimit
    leafs["socket-error"] = detailInformation.SocketError
    leafs["auto-rearm"] = detailInformation.AutoRearm
    leafs["send-pdu-count"] = detailInformation.SendPduCount
    leafs["output-ifhandle"] = detailInformation.OutputIfhandle
    leafs["fib-pd-ctx-size"] = detailInformation.FibPdCtxSize
    leafs["num-labels"] = detailInformation.NumLabels
    leafs["local-app-instance"] = detailInformation.LocalAppInstance
    leafs["fib-pd-ctx"] = detailInformation.FibPdCtx
    leafs["fib-label-output"] = detailInformation.FibLabelOutput
    return leafs
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetBundleName() string { return "cisco_ios_xr" }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetYangName() string { return "detail-information" }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) SetParent(parent types.Entity) { detailInformation.parent = parent }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetParent() types.Entity { return detailInformation.parent }

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetParentYangName() string { return "detail-informations" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress
// Local address
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is TcpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress
// Foreign address
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is TcpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetFilter() yfilter.YFilter { return foreignAddress.YFilter }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) SetFilter(yf yfilter.YFilter) { foreignAddress.YFilter = yf }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetSegmentPath() string {
    return "foreign-address"
}

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = foreignAddress.AfName
    leafs["ipv4-address"] = foreignAddress.Ipv4Address
    leafs["ipv6-address"] = foreignAddress.Ipv6Address
    return leafs
}

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetBundleName() string { return "cisco_ios_xr" }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetYangName() string { return "foreign-address" }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) SetParent(parent types.Entity) { foreignAddress.parent = parent }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetParent() types.Entity { return foreignAddress.parent }

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags
// Socket option flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Turn on debugging info recording. The type is bool.
    Debug interface{}

    // Socket has had listen(). The type is bool.
    AcceptConnection interface{}

    // Allow local address reuse. The type is bool.
    ReuseAddress interface{}

    // Keep connections alive. The type is bool.
    KeepAlive interface{}

    // Just use interface addresses. The type is bool.
    DontRoute interface{}

    // Permit sending of broadcast msgs. The type is bool.
    Broadcast interface{}

    // Bypass hardware when possible. The type is bool.
    UseLoopback interface{}

    // Linger on close if data present. The type is bool.
    Linger interface{}

    // Leave received Out-of-band data inline. The type is bool.
    OutOfBandInline interface{}

    // Allow local address & port reuse. The type is bool.
    ReusePort interface{}

    // Nonblocking socket I/O operation. The type is bool.
    NonblockingIo interface{}
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetFilter() yfilter.YFilter { return socketOptionFlags.YFilter }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) SetFilter(yf yfilter.YFilter) { socketOptionFlags.YFilter = yf }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetGoName(yname string) string {
    if yname == "debug" { return "Debug" }
    if yname == "accept-connection" { return "AcceptConnection" }
    if yname == "reuse-address" { return "ReuseAddress" }
    if yname == "keep-alive" { return "KeepAlive" }
    if yname == "dont-route" { return "DontRoute" }
    if yname == "broadcast" { return "Broadcast" }
    if yname == "use-loopback" { return "UseLoopback" }
    if yname == "linger" { return "Linger" }
    if yname == "out-of-band-inline" { return "OutOfBandInline" }
    if yname == "reuse-port" { return "ReusePort" }
    if yname == "nonblocking-io" { return "NonblockingIo" }
    return ""
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetSegmentPath() string {
    return "socket-option-flags"
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["debug"] = socketOptionFlags.Debug
    leafs["accept-connection"] = socketOptionFlags.AcceptConnection
    leafs["reuse-address"] = socketOptionFlags.ReuseAddress
    leafs["keep-alive"] = socketOptionFlags.KeepAlive
    leafs["dont-route"] = socketOptionFlags.DontRoute
    leafs["broadcast"] = socketOptionFlags.Broadcast
    leafs["use-loopback"] = socketOptionFlags.UseLoopback
    leafs["linger"] = socketOptionFlags.Linger
    leafs["out-of-band-inline"] = socketOptionFlags.OutOfBandInline
    leafs["reuse-port"] = socketOptionFlags.ReusePort
    leafs["nonblocking-io"] = socketOptionFlags.NonblockingIo
    return leafs
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetBundleName() string { return "cisco_ios_xr" }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetYangName() string { return "socket-option-flags" }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) SetParent(parent types.Entity) { socketOptionFlags.parent = parent }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetParent() types.Entity { return socketOptionFlags.parent }

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags
// Socket state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // No file descriptor ref. The type is bool.
    NoFileDescriptorReference interface{}

    // Socket is connected to peer. The type is bool.
    IsConnected interface{}

    // Connecting in progress. The type is bool.
    IsConnecting interface{}

    // Disconnecting in progress. The type is bool.
    IsDisconnecting interface{}

    // Can't send more data to peer. The type is bool.
    CantSendMore interface{}

    // Can't recv more data from peer. The type is bool.
    CantReceiveMore interface{}

    // At mark on input. The type is bool.
    ReceivedAtMark interface{}

    // Privileged for broadcast, raw... The type is bool.
    Privileged interface{}

    // Close is blocked (i.e. socket is a replicated socket on a standby node. The
    // type is bool.
    BlockClose interface{}

    // Async i/o notify. The type is bool.
    AsyncIoNotify interface{}

    // Deciding to accept connection req. The type is bool.
    IsConfirming interface{}

    // Mutex acquired by solock(). The type is bool.
    IsSolock interface{}

    // PCB and socket are detached. The type is bool.
    IsDetached interface{}

    // Socket is blocked for receive - while going through SSO initial sync. The
    // type is bool.
    BlockReceive interface{}

    // Socket is blocked for send (if it is a replicated socket on a standby
    // node). The type is bool.
    BlockSend interface{}
}

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetFilter() yfilter.YFilter { return socketStateFlags.YFilter }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) SetFilter(yf yfilter.YFilter) { socketStateFlags.YFilter = yf }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetGoName(yname string) string {
    if yname == "no-file-descriptor-reference" { return "NoFileDescriptorReference" }
    if yname == "is-connected" { return "IsConnected" }
    if yname == "is-connecting" { return "IsConnecting" }
    if yname == "is-disconnecting" { return "IsDisconnecting" }
    if yname == "cant-send-more" { return "CantSendMore" }
    if yname == "cant-receive-more" { return "CantReceiveMore" }
    if yname == "received-at-mark" { return "ReceivedAtMark" }
    if yname == "privileged" { return "Privileged" }
    if yname == "block-close" { return "BlockClose" }
    if yname == "async-io-notify" { return "AsyncIoNotify" }
    if yname == "is-confirming" { return "IsConfirming" }
    if yname == "is-solock" { return "IsSolock" }
    if yname == "is-detached" { return "IsDetached" }
    if yname == "block-receive" { return "BlockReceive" }
    if yname == "block-send" { return "BlockSend" }
    return ""
}

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetSegmentPath() string {
    return "socket-state-flags"
}

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["no-file-descriptor-reference"] = socketStateFlags.NoFileDescriptorReference
    leafs["is-connected"] = socketStateFlags.IsConnected
    leafs["is-connecting"] = socketStateFlags.IsConnecting
    leafs["is-disconnecting"] = socketStateFlags.IsDisconnecting
    leafs["cant-send-more"] = socketStateFlags.CantSendMore
    leafs["cant-receive-more"] = socketStateFlags.CantReceiveMore
    leafs["received-at-mark"] = socketStateFlags.ReceivedAtMark
    leafs["privileged"] = socketStateFlags.Privileged
    leafs["block-close"] = socketStateFlags.BlockClose
    leafs["async-io-notify"] = socketStateFlags.AsyncIoNotify
    leafs["is-confirming"] = socketStateFlags.IsConfirming
    leafs["is-solock"] = socketStateFlags.IsSolock
    leafs["is-detached"] = socketStateFlags.IsDetached
    leafs["block-receive"] = socketStateFlags.BlockReceive
    leafs["block-send"] = socketStateFlags.BlockSend
    return leafs
}

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetYangName() string { return "socket-state-flags" }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) SetParent(parent types.Entity) { socketStateFlags.parent = parent }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetParent() types.Entity { return socketStateFlags.parent }

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags
// Connection feature flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selective ack on?. The type is bool.
    SelectiveAck interface{}

    // MD5 option on?. The type is bool.
    Md5 interface{}

    // Timestamps on?. The type is bool.
    Timestamps interface{}

    // Window-scaling on?. The type is bool.
    WindowScaling interface{}

    // Nagle algorithm on?. The type is bool.
    Nagle interface{}

    // Giveup timer is on?. The type is bool.
    GiveupTimer interface{}

    // Keepalive timer is on?. The type is bool.
    ConnectionKeepAliveTimer interface{}

    // Path MTU Discovery feature is on?. The type is bool.
    PathMtuDiscovery interface{}

    // tcp mss feature is on?. The type is bool.
    MssCisco interface{}
}

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetFilter() yfilter.YFilter { return featureFlags.YFilter }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) SetFilter(yf yfilter.YFilter) { featureFlags.YFilter = yf }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetGoName(yname string) string {
    if yname == "selective-ack" { return "SelectiveAck" }
    if yname == "md5" { return "Md5" }
    if yname == "timestamps" { return "Timestamps" }
    if yname == "window-scaling" { return "WindowScaling" }
    if yname == "nagle" { return "Nagle" }
    if yname == "giveup-timer" { return "GiveupTimer" }
    if yname == "connection-keep-alive-timer" { return "ConnectionKeepAliveTimer" }
    if yname == "path-mtu-discovery" { return "PathMtuDiscovery" }
    if yname == "mss-cisco" { return "MssCisco" }
    return ""
}

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetSegmentPath() string {
    return "feature-flags"
}

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selective-ack"] = featureFlags.SelectiveAck
    leafs["md5"] = featureFlags.Md5
    leafs["timestamps"] = featureFlags.Timestamps
    leafs["window-scaling"] = featureFlags.WindowScaling
    leafs["nagle"] = featureFlags.Nagle
    leafs["giveup-timer"] = featureFlags.GiveupTimer
    leafs["connection-keep-alive-timer"] = featureFlags.ConnectionKeepAliveTimer
    leafs["path-mtu-discovery"] = featureFlags.PathMtuDiscovery
    leafs["mss-cisco"] = featureFlags.MssCisco
    return leafs
}

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetBundleName() string { return "cisco_ios_xr" }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetYangName() string { return "feature-flags" }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) SetParent(parent types.Entity) { featureFlags.parent = parent }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetParent() types.Entity { return featureFlags.parent }

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags
// Connection state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Nagle has delayed output. The type is bool.
    NagleWait interface{}

    // Send an ACK. The type is bool.
    AckNeeded interface{}

    // FIN has been sent. The type is bool.
    FinSent interface{}

    // Probing a closed window. The type is bool.
    Probing interface{}

    // Need to push data out. The type is bool.
    NeedPush interface{}

    // A segment is pushed due to MSG_PUSH. The type is bool.
    Pushed interface{}

    // Connection is in SYN cache. The type is bool.
    InSynCache interface{}

    // Path MTU aging timer is running. The type is bool.
    PathMtuAger interface{}
}

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetFilter() yfilter.YFilter { return stateFlags.YFilter }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) SetFilter(yf yfilter.YFilter) { stateFlags.YFilter = yf }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetGoName(yname string) string {
    if yname == "nagle-wait" { return "NagleWait" }
    if yname == "ack-needed" { return "AckNeeded" }
    if yname == "fin-sent" { return "FinSent" }
    if yname == "probing" { return "Probing" }
    if yname == "need-push" { return "NeedPush" }
    if yname == "pushed" { return "Pushed" }
    if yname == "in-syn-cache" { return "InSynCache" }
    if yname == "path-mtu-ager" { return "PathMtuAger" }
    return ""
}

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetSegmentPath() string {
    return "state-flags"
}

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["nagle-wait"] = stateFlags.NagleWait
    leafs["ack-needed"] = stateFlags.AckNeeded
    leafs["fin-sent"] = stateFlags.FinSent
    leafs["probing"] = stateFlags.Probing
    leafs["need-push"] = stateFlags.NeedPush
    leafs["pushed"] = stateFlags.Pushed
    leafs["in-syn-cache"] = stateFlags.InSynCache
    leafs["path-mtu-ager"] = stateFlags.PathMtuAger
    return leafs
}

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetYangName() string { return "state-flags" }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) SetParent(parent types.Entity) { stateFlags.parent = parent }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetParent() types.Entity { return stateFlags.parent }

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags
// Connection request flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Selective ack on?. The type is bool.
    SelectiveAck interface{}

    // MD5 option on?. The type is bool.
    Md5 interface{}

    // Timestamps on?. The type is bool.
    Timestamps interface{}

    // Window-scaling on?. The type is bool.
    WindowScaling interface{}

    // Nagle algorithm on?. The type is bool.
    Nagle interface{}

    // Giveup timer is on?. The type is bool.
    GiveupTimer interface{}

    // Keepalive timer is on?. The type is bool.
    ConnectionKeepAliveTimer interface{}

    // Path MTU Discovery feature is on?. The type is bool.
    PathMtuDiscovery interface{}

    // tcp mss feature is on?. The type is bool.
    MssCisco interface{}
}

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetFilter() yfilter.YFilter { return requestFlags.YFilter }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) SetFilter(yf yfilter.YFilter) { requestFlags.YFilter = yf }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetGoName(yname string) string {
    if yname == "selective-ack" { return "SelectiveAck" }
    if yname == "md5" { return "Md5" }
    if yname == "timestamps" { return "Timestamps" }
    if yname == "window-scaling" { return "WindowScaling" }
    if yname == "nagle" { return "Nagle" }
    if yname == "giveup-timer" { return "GiveupTimer" }
    if yname == "connection-keep-alive-timer" { return "ConnectionKeepAliveTimer" }
    if yname == "path-mtu-discovery" { return "PathMtuDiscovery" }
    if yname == "mss-cisco" { return "MssCisco" }
    return ""
}

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetSegmentPath() string {
    return "request-flags"
}

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["selective-ack"] = requestFlags.SelectiveAck
    leafs["md5"] = requestFlags.Md5
    leafs["timestamps"] = requestFlags.Timestamps
    leafs["window-scaling"] = requestFlags.WindowScaling
    leafs["nagle"] = requestFlags.Nagle
    leafs["giveup-timer"] = requestFlags.GiveupTimer
    leafs["connection-keep-alive-timer"] = requestFlags.ConnectionKeepAliveTimer
    leafs["path-mtu-discovery"] = requestFlags.PathMtuDiscovery
    leafs["mss-cisco"] = requestFlags.MssCisco
    return leafs
}

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetBundleName() string { return "cisco_ios_xr" }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetYangName() string { return "request-flags" }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) SetParent(parent types.Entity) { requestFlags.parent = parent }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetParent() types.Entity { return requestFlags.parent }

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags
// Receive buffer state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Lock on data queue (so_rcv only). The type is bool.
    Locked interface{}

    // Someone is waiting to lock. The type is bool.
    WaitingForLock interface{}

    // Someone is waiting for data/space. The type is bool.
    WaitingForData interface{}

    // Buffer is selected for INPUT. The type is bool.
    InputSelect interface{}

    // Async I/O. The type is bool.
    AsyncIo interface{}

    // Not interruptible. The type is bool.
    NotInterruptible interface{}

    // Read/write timer set. The type is bool.
    IoTimerSet interface{}

    // Want delayed wakeups. The type is bool.
    DelayedWakeup interface{}

    // Read/write wakeup pending. The type is bool.
    Wakeup interface{}

    // Connect wakeup pending. The type is bool.
    ConnectWakeup interface{}

    // Buffer is selected for OUTPUT. The type is bool.
    OutputSelect interface{}

    // Buffer is selected for OBAND. The type is bool.
    OutOfBandSelect interface{}
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetFilter() yfilter.YFilter { return receiveBufStateFlags.YFilter }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) SetFilter(yf yfilter.YFilter) { receiveBufStateFlags.YFilter = yf }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetGoName(yname string) string {
    if yname == "locked" { return "Locked" }
    if yname == "waiting-for-lock" { return "WaitingForLock" }
    if yname == "waiting-for-data" { return "WaitingForData" }
    if yname == "input-select" { return "InputSelect" }
    if yname == "async-io" { return "AsyncIo" }
    if yname == "not-interruptible" { return "NotInterruptible" }
    if yname == "io-timer-set" { return "IoTimerSet" }
    if yname == "delayed-wakeup" { return "DelayedWakeup" }
    if yname == "wakeup" { return "Wakeup" }
    if yname == "connect-wakeup" { return "ConnectWakeup" }
    if yname == "output-select" { return "OutputSelect" }
    if yname == "out-of-band-select" { return "OutOfBandSelect" }
    return ""
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetSegmentPath() string {
    return "receive-buf-state-flags"
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["locked"] = receiveBufStateFlags.Locked
    leafs["waiting-for-lock"] = receiveBufStateFlags.WaitingForLock
    leafs["waiting-for-data"] = receiveBufStateFlags.WaitingForData
    leafs["input-select"] = receiveBufStateFlags.InputSelect
    leafs["async-io"] = receiveBufStateFlags.AsyncIo
    leafs["not-interruptible"] = receiveBufStateFlags.NotInterruptible
    leafs["io-timer-set"] = receiveBufStateFlags.IoTimerSet
    leafs["delayed-wakeup"] = receiveBufStateFlags.DelayedWakeup
    leafs["wakeup"] = receiveBufStateFlags.Wakeup
    leafs["connect-wakeup"] = receiveBufStateFlags.ConnectWakeup
    leafs["output-select"] = receiveBufStateFlags.OutputSelect
    leafs["out-of-band-select"] = receiveBufStateFlags.OutOfBandSelect
    return leafs
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetYangName() string { return "receive-buf-state-flags" }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) SetParent(parent types.Entity) { receiveBufStateFlags.parent = parent }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetParent() types.Entity { return receiveBufStateFlags.parent }

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags
// Send buffer state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Lock on data queue (so_rcv only). The type is bool.
    Locked interface{}

    // Someone is waiting to lock. The type is bool.
    WaitingForLock interface{}

    // Someone is waiting for data/space. The type is bool.
    WaitingForData interface{}

    // Buffer is selected for INPUT. The type is bool.
    InputSelect interface{}

    // Async I/O. The type is bool.
    AsyncIo interface{}

    // Not interruptible. The type is bool.
    NotInterruptible interface{}

    // Read/write timer set. The type is bool.
    IoTimerSet interface{}

    // Want delayed wakeups. The type is bool.
    DelayedWakeup interface{}

    // Read/write wakeup pending. The type is bool.
    Wakeup interface{}

    // Connect wakeup pending. The type is bool.
    ConnectWakeup interface{}

    // Buffer is selected for OUTPUT. The type is bool.
    OutputSelect interface{}

    // Buffer is selected for OBAND. The type is bool.
    OutOfBandSelect interface{}
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetFilter() yfilter.YFilter { return sendBufStateFlags.YFilter }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) SetFilter(yf yfilter.YFilter) { sendBufStateFlags.YFilter = yf }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetGoName(yname string) string {
    if yname == "locked" { return "Locked" }
    if yname == "waiting-for-lock" { return "WaitingForLock" }
    if yname == "waiting-for-data" { return "WaitingForData" }
    if yname == "input-select" { return "InputSelect" }
    if yname == "async-io" { return "AsyncIo" }
    if yname == "not-interruptible" { return "NotInterruptible" }
    if yname == "io-timer-set" { return "IoTimerSet" }
    if yname == "delayed-wakeup" { return "DelayedWakeup" }
    if yname == "wakeup" { return "Wakeup" }
    if yname == "connect-wakeup" { return "ConnectWakeup" }
    if yname == "output-select" { return "OutputSelect" }
    if yname == "out-of-band-select" { return "OutOfBandSelect" }
    return ""
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetSegmentPath() string {
    return "send-buf-state-flags"
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["locked"] = sendBufStateFlags.Locked
    leafs["waiting-for-lock"] = sendBufStateFlags.WaitingForLock
    leafs["waiting-for-data"] = sendBufStateFlags.WaitingForData
    leafs["input-select"] = sendBufStateFlags.InputSelect
    leafs["async-io"] = sendBufStateFlags.AsyncIo
    leafs["not-interruptible"] = sendBufStateFlags.NotInterruptible
    leafs["io-timer-set"] = sendBufStateFlags.IoTimerSet
    leafs["delayed-wakeup"] = sendBufStateFlags.DelayedWakeup
    leafs["wakeup"] = sendBufStateFlags.Wakeup
    leafs["connect-wakeup"] = sendBufStateFlags.ConnectWakeup
    leafs["output-select"] = sendBufStateFlags.OutputSelect
    leafs["out-of-band-select"] = sendBufStateFlags.OutOfBandSelect
    return leafs
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetBundleName() string { return "cisco_ios_xr" }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetYangName() string { return "send-buf-state-flags" }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) SetParent(parent types.Entity) { sendBufStateFlags.parent = parent }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetParent() types.Entity { return sendBufStateFlags.parent }

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer
// Timers
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timer Type. The type is TcpTimer.
    TimerType interface{}

    // Count of timer activations. The type is interface{} with range:
    // 0..4294967295.
    TimerActivations interface{}

    // Count of timer expirations. The type is interface{} with range:
    // 0..4294967295.
    TimerExpirations interface{}

    // Timer next activation (msec). The type is interface{} with range:
    // 0..4294967295.
    TimerNextActivation interface{}
}

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetFilter() yfilter.YFilter { return timer.YFilter }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) SetFilter(yf yfilter.YFilter) { timer.YFilter = yf }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetGoName(yname string) string {
    if yname == "timer-type" { return "TimerType" }
    if yname == "timer-activations" { return "TimerActivations" }
    if yname == "timer-expirations" { return "TimerExpirations" }
    if yname == "timer-next-activation" { return "TimerNextActivation" }
    return ""
}

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetSegmentPath() string {
    return "timer"
}

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timer-type"] = timer.TimerType
    leafs["timer-activations"] = timer.TimerActivations
    leafs["timer-expirations"] = timer.TimerExpirations
    leafs["timer-next-activation"] = timer.TimerNextActivation
    return leafs
}

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetBundleName() string { return "cisco_ios_xr" }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetYangName() string { return "timer" }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) SetParent(parent types.Entity) { timer.parent = parent }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetParent() types.Entity { return timer.parent }

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk
// Seq nos. of sack blocks
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start seq no. of sack block. The type is interface{} with range:
    // 0..4294967295.
    Start interface{}

    // End   seq no. of sack block. The type is interface{} with range:
    // 0..4294967295.
    End interface{}
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetFilter() yfilter.YFilter { return sackBlk.YFilter }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) SetFilter(yf yfilter.YFilter) { sackBlk.YFilter = yf }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    return ""
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetSegmentPath() string {
    return "sack-blk"
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = sackBlk.Start
    leafs["end"] = sackBlk.End
    return leafs
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetBundleName() string { return "cisco_ios_xr" }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetYangName() string { return "sack-blk" }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) SetParent(parent types.Entity) { sackBlk.parent = parent }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetParent() types.Entity { return sackBlk.parent }

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole
// Sorted list of sack holes
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Start seq no. of hole. The type is interface{} with range: 0..4294967295.
    Start interface{}

    // End   seq no. of hole. The type is interface{} with range: 0..4294967295.
    End interface{}

    // Number of dup (s)acks for this hole. The type is interface{} with range:
    // 0..4294967295.
    DuplicatedAck interface{}

    // Next seq. no in hole to be retransmitted. The type is interface{} with
    // range: 0..4294967295.
    Retransmitted interface{}
}

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetFilter() yfilter.YFilter { return sendSackHole.YFilter }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) SetFilter(yf yfilter.YFilter) { sendSackHole.YFilter = yf }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetGoName(yname string) string {
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "duplicated-ack" { return "DuplicatedAck" }
    if yname == "retransmitted" { return "Retransmitted" }
    return ""
}

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetSegmentPath() string {
    return "send-sack-hole"
}

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["start"] = sendSackHole.Start
    leafs["end"] = sendSackHole.End
    leafs["duplicated-ack"] = sendSackHole.DuplicatedAck
    leafs["retransmitted"] = sendSackHole.Retransmitted
    return leafs
}

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetBundleName() string { return "cisco_ios_xr" }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetYangName() string { return "send-sack-hole" }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) SetParent(parent types.Entity) { sendSackHole.parent = parent }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetParent() types.Entity { return sendSackHole.parent }

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetParentYangName() string { return "detail-information" }

// TcpConnection_Nodes_Node_BriefInformations
// Table listing connections for which brief
// information is provided.Note that not all
// connections are listed in the brief table.
type TcpConnection_Nodes_Node_BriefInformations struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about a TCP connection. The type is slice of
    // TcpConnection_Nodes_Node_BriefInformations_BriefInformation.
    BriefInformation []TcpConnection_Nodes_Node_BriefInformations_BriefInformation
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetFilter() yfilter.YFilter { return briefInformations.YFilter }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) SetFilter(yf yfilter.YFilter) { briefInformations.YFilter = yf }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetGoName(yname string) string {
    if yname == "brief-information" { return "BriefInformation" }
    return ""
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetSegmentPath() string {
    return "brief-informations"
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-information" {
        for _, c := range briefInformations.BriefInformation {
            if briefInformations.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpConnection_Nodes_Node_BriefInformations_BriefInformation{}
        briefInformations.BriefInformation = append(briefInformations.BriefInformation, child)
        return &briefInformations.BriefInformation[len(briefInformations.BriefInformation)-1]
    }
    return nil
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefInformations.BriefInformation {
        children[briefInformations.BriefInformation[i].GetSegmentPath()] = &briefInformations.BriefInformation[i]
    }
    return children
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetBundleName() string { return "cisco_ios_xr" }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetYangName() string { return "brief-informations" }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) SetParent(parent types.Entity) { briefInformations.parent = parent }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetParent() types.Entity { return briefInformations.parent }

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetParentYangName() string { return "node" }

// TcpConnection_Nodes_Node_BriefInformations_BriefInformation
// Brief information about a TCP connection
type TcpConnection_Nodes_Node_BriefInformations_BriefInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Protocol Control Block ID. The type is interface{}
    // with range: 0..4294967295.
    PcbId interface{}

    // Address family. The type is TcpAddressFamily.
    AfName interface{}

    // PCB Address. The type is interface{} with range: 0..18446744073709551615.
    Pcb interface{}

    // Connection state. The type is TcpConnState.
    ConnectionState interface{}

    // Id of the local process. The type is interface{} with range: 0..4294967295.
    LocalPid interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // Current receive queue size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    CurrentReceiveQueueSize interface{}

    // Current send queue size in bytes. The type is interface{} with range:
    // 0..4294967295. Units are byte.
    CurrentSendQueueSize interface{}

    // VRF ID. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Local address.
    LocalAddress TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress

    // Foreign address.
    ForeignAddress TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress
}

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetFilter() yfilter.YFilter { return briefInformation.YFilter }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) SetFilter(yf yfilter.YFilter) { briefInformation.YFilter = yf }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetGoName(yname string) string {
    if yname == "pcb-id" { return "PcbId" }
    if yname == "af-name" { return "AfName" }
    if yname == "pcb" { return "Pcb" }
    if yname == "connection-state" { return "ConnectionState" }
    if yname == "local-pid" { return "LocalPid" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "current-receive-queue-size" { return "CurrentReceiveQueueSize" }
    if yname == "current-send-queue-size" { return "CurrentSendQueueSize" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    return ""
}

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetSegmentPath() string {
    return "brief-information" + "[pcb-id='" + fmt.Sprintf("%v", briefInformation.PcbId) + "']"
}

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-address" {
        return &briefInformation.LocalAddress
    }
    if childYangName == "foreign-address" {
        return &briefInformation.ForeignAddress
    }
    return nil
}

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["local-address"] = &briefInformation.LocalAddress
    children["foreign-address"] = &briefInformation.ForeignAddress
    return children
}

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["pcb-id"] = briefInformation.PcbId
    leafs["af-name"] = briefInformation.AfName
    leafs["pcb"] = briefInformation.Pcb
    leafs["connection-state"] = briefInformation.ConnectionState
    leafs["local-pid"] = briefInformation.LocalPid
    leafs["local-port"] = briefInformation.LocalPort
    leafs["foreign-port"] = briefInformation.ForeignPort
    leafs["current-receive-queue-size"] = briefInformation.CurrentReceiveQueueSize
    leafs["current-send-queue-size"] = briefInformation.CurrentSendQueueSize
    leafs["vrf-id"] = briefInformation.VrfId
    return leafs
}

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetBundleName() string { return "cisco_ios_xr" }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetYangName() string { return "brief-information" }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) SetParent(parent types.Entity) { briefInformation.parent = parent }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetParent() types.Entity { return briefInformation.parent }

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetParentYangName() string { return "brief-informations" }

// TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress
// Local address
type TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is TcpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetFilter() yfilter.YFilter { return localAddress.YFilter }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) SetFilter(yf yfilter.YFilter) { localAddress.YFilter = yf }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetSegmentPath() string {
    return "local-address"
}

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = localAddress.AfName
    leafs["ipv4-address"] = localAddress.Ipv4Address
    leafs["ipv6-address"] = localAddress.Ipv6Address
    return leafs
}

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetBundleName() string { return "cisco_ios_xr" }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetYangName() string { return "local-address" }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) SetParent(parent types.Entity) { localAddress.parent = parent }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetParent() types.Entity { return localAddress.parent }

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetParentYangName() string { return "brief-information" }

// TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress
// Foreign address
type TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFName. The type is TcpAddressFamily.
    AfName interface{}

    // IPv4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4Address interface{}

    // IPv6 Address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6Address interface{}
}

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetFilter() yfilter.YFilter { return foreignAddress.YFilter }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) SetFilter(yf yfilter.YFilter) { foreignAddress.YFilter = yf }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetGoName(yname string) string {
    if yname == "af-name" { return "AfName" }
    if yname == "ipv4-address" { return "Ipv4Address" }
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetSegmentPath() string {
    return "foreign-address"
}

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["af-name"] = foreignAddress.AfName
    leafs["ipv4-address"] = foreignAddress.Ipv4Address
    leafs["ipv6-address"] = foreignAddress.Ipv6Address
    return leafs
}

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetBundleName() string { return "cisco_ios_xr" }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetYangName() string { return "foreign-address" }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) SetParent(parent types.Entity) { foreignAddress.parent = parent }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetParent() types.Entity { return foreignAddress.parent }

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetParentYangName() string { return "brief-information" }

// Tcp
// tcp
type Tcp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Node-specific TCP operational data.
    Nodes Tcp_Nodes
}

func (tcp *Tcp) GetFilter() yfilter.YFilter { return tcp.YFilter }

func (tcp *Tcp) SetFilter(yf yfilter.YFilter) { tcp.YFilter = yf }

func (tcp *Tcp) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (tcp *Tcp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-tcp-oper:tcp"
}

func (tcp *Tcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &tcp.Nodes
    }
    return nil
}

func (tcp *Tcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &tcp.Nodes
    return children
}

func (tcp *Tcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcp *Tcp) GetBundleName() string { return "cisco_ios_xr" }

func (tcp *Tcp) GetYangName() string { return "tcp" }

func (tcp *Tcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcp *Tcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcp *Tcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcp *Tcp) SetParent(parent types.Entity) { tcp.parent = parent }

func (tcp *Tcp) GetParent() types.Entity { return tcp.parent }

func (tcp *Tcp) GetParentYangName() string { return "Cisco-IOS-XR-ip-tcp-oper" }

// Tcp_Nodes
// Node-specific TCP operational data
type Tcp_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP operational data for a particular node. The type is slice of
    // Tcp_Nodes_Node.
    Node []Tcp_Nodes_Node
}

func (nodes *Tcp_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Tcp_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Tcp_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Tcp_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Tcp_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tcp_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Tcp_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Tcp_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Tcp_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Tcp_Nodes) GetYangName() string { return "nodes" }

func (nodes *Tcp_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Tcp_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Tcp_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Tcp_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Tcp_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Tcp_Nodes) GetParentYangName() string { return "tcp" }

// Tcp_Nodes_Node
// TCP operational data for a particular node
type Tcp_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Statistical TCP operational data for a node.
    Statistics Tcp_Nodes_Node_Statistics
}

func (node *Tcp_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Tcp_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Tcp_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *Tcp_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Tcp_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *Tcp_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statistics"] = &node.Statistics
    return children
}

func (node *Tcp_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    return leafs
}

func (node *Tcp_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Tcp_Nodes_Node) GetYangName() string { return "node" }

func (node *Tcp_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Tcp_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Tcp_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Tcp_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Tcp_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Tcp_Nodes_Node) GetParentYangName() string { return "nodes" }

// Tcp_Nodes_Node_Statistics
// Statistical TCP operational data for a node
type Tcp_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP Traffic statistics for IPv4.
    Ipv4Traffic Tcp_Nodes_Node_Statistics_Ipv4Traffic

    // TCP Traffic statistics for IPv6.
    Ipv6Traffic Tcp_Nodes_Node_Statistics_Ipv6Traffic
}

func (statistics *Tcp_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *Tcp_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *Tcp_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "ipv4-traffic" { return "Ipv4Traffic" }
    if yname == "ipv6-traffic" { return "Ipv6Traffic" }
    return ""
}

func (statistics *Tcp_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *Tcp_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-traffic" {
        return &statistics.Ipv4Traffic
    }
    if childYangName == "ipv6-traffic" {
        return &statistics.Ipv6Traffic
    }
    return nil
}

func (statistics *Tcp_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-traffic"] = &statistics.Ipv4Traffic
    children["ipv6-traffic"] = &statistics.Ipv6Traffic
    return children
}

func (statistics *Tcp_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *Tcp_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *Tcp_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *Tcp_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *Tcp_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *Tcp_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *Tcp_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *Tcp_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *Tcp_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// Tcp_Nodes_Node_Statistics_Ipv4Traffic
// TCP Traffic statistics for IPv4
type Tcp_Nodes_Node_Statistics_Ipv4Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP packets received. The type is interface{} with range: 0..4294967295.
    TcpInputPackets interface{}

    // TCP packets with checksum errors. The type is interface{} with range:
    // 0..4294967295.
    TcpChecksumErrorPackets interface{}

    // TCP packets dropped (no port). The type is interface{} with range:
    // 0..4294967295.
    TcpDroppedPackets interface{}

    // TCP packets transmitted. The type is interface{} with range: 0..4294967295.
    TcpOutputPackets interface{}

    // TCP packets retransmitted. The type is interface{} with range:
    // 0..4294967295.
    TcpRetransmittedPackets interface{}
}

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetFilter() yfilter.YFilter { return ipv4Traffic.YFilter }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) SetFilter(yf yfilter.YFilter) { ipv4Traffic.YFilter = yf }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetGoName(yname string) string {
    if yname == "tcp-input-packets" { return "TcpInputPackets" }
    if yname == "tcp-checksum-error-packets" { return "TcpChecksumErrorPackets" }
    if yname == "tcp-dropped-packets" { return "TcpDroppedPackets" }
    if yname == "tcp-output-packets" { return "TcpOutputPackets" }
    if yname == "tcp-retransmitted-packets" { return "TcpRetransmittedPackets" }
    return ""
}

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetSegmentPath() string {
    return "ipv4-traffic"
}

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-input-packets"] = ipv4Traffic.TcpInputPackets
    leafs["tcp-checksum-error-packets"] = ipv4Traffic.TcpChecksumErrorPackets
    leafs["tcp-dropped-packets"] = ipv4Traffic.TcpDroppedPackets
    leafs["tcp-output-packets"] = ipv4Traffic.TcpOutputPackets
    leafs["tcp-retransmitted-packets"] = ipv4Traffic.TcpRetransmittedPackets
    return leafs
}

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetYangName() string { return "ipv4-traffic" }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) SetParent(parent types.Entity) { ipv4Traffic.parent = parent }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetParent() types.Entity { return ipv4Traffic.parent }

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetParentYangName() string { return "statistics" }

// Tcp_Nodes_Node_Statistics_Ipv6Traffic
// TCP Traffic statistics for IPv6
type Tcp_Nodes_Node_Statistics_Ipv6Traffic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TCP packets received. The type is interface{} with range: 0..4294967295.
    TcpInputPackets interface{}

    // TCP packets with checksum errors. The type is interface{} with range:
    // 0..4294967295.
    TcpChecksumErrorPackets interface{}

    // TCP packets dropped (no port). The type is interface{} with range:
    // 0..4294967295.
    TcpDroppedPackets interface{}

    // TCP packets transmitted. The type is interface{} with range: 0..4294967295.
    TcpOutputPackets interface{}

    // TCP packets retransmitted. The type is interface{} with range:
    // 0..4294967295.
    TcpRetransmittedPackets interface{}
}

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetFilter() yfilter.YFilter { return ipv6Traffic.YFilter }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) SetFilter(yf yfilter.YFilter) { ipv6Traffic.YFilter = yf }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetGoName(yname string) string {
    if yname == "tcp-input-packets" { return "TcpInputPackets" }
    if yname == "tcp-checksum-error-packets" { return "TcpChecksumErrorPackets" }
    if yname == "tcp-dropped-packets" { return "TcpDroppedPackets" }
    if yname == "tcp-output-packets" { return "TcpOutputPackets" }
    if yname == "tcp-retransmitted-packets" { return "TcpRetransmittedPackets" }
    return ""
}

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetSegmentPath() string {
    return "ipv6-traffic"
}

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tcp-input-packets"] = ipv6Traffic.TcpInputPackets
    leafs["tcp-checksum-error-packets"] = ipv6Traffic.TcpChecksumErrorPackets
    leafs["tcp-dropped-packets"] = ipv6Traffic.TcpDroppedPackets
    leafs["tcp-output-packets"] = ipv6Traffic.TcpOutputPackets
    leafs["tcp-retransmitted-packets"] = ipv6Traffic.TcpRetransmittedPackets
    return leafs
}

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetYangName() string { return "ipv6-traffic" }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) SetParent(parent types.Entity) { ipv6Traffic.parent = parent }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetParent() types.Entity { return ipv6Traffic.parent }

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetParentYangName() string { return "statistics" }

// TcpNsr
// tcp nsr
type TcpNsr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of information about all nodes present on the system.
    Nodes TcpNsr_Nodes
}

func (tcpNsr *TcpNsr) GetFilter() yfilter.YFilter { return tcpNsr.YFilter }

func (tcpNsr *TcpNsr) SetFilter(yf yfilter.YFilter) { tcpNsr.YFilter = yf }

func (tcpNsr *TcpNsr) GetGoName(yname string) string {
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (tcpNsr *TcpNsr) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr"
}

func (tcpNsr *TcpNsr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "nodes" {
        return &tcpNsr.Nodes
    }
    return nil
}

func (tcpNsr *TcpNsr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["nodes"] = &tcpNsr.Nodes
    return children
}

func (tcpNsr *TcpNsr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tcpNsr *TcpNsr) GetBundleName() string { return "cisco_ios_xr" }

func (tcpNsr *TcpNsr) GetYangName() string { return "tcp-nsr" }

func (tcpNsr *TcpNsr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tcpNsr *TcpNsr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tcpNsr *TcpNsr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tcpNsr *TcpNsr) SetParent(parent types.Entity) { tcpNsr.parent = parent }

func (tcpNsr *TcpNsr) GetParent() types.Entity { return tcpNsr.parent }

func (tcpNsr *TcpNsr) GetParentYangName() string { return "Cisco-IOS-XR-ip-tcp-oper" }

// TcpNsr_Nodes
// Table of information about all nodes present on
// the system
type TcpNsr_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about a single node. The type is slice of TcpNsr_Nodes_Node.
    Node []TcpNsr_Nodes_Node
}

func (nodes *TcpNsr_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *TcpNsr_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *TcpNsr_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *TcpNsr_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *TcpNsr_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *TcpNsr_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *TcpNsr_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *TcpNsr_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *TcpNsr_Nodes) GetYangName() string { return "nodes" }

func (nodes *TcpNsr_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *TcpNsr_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *TcpNsr_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *TcpNsr_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *TcpNsr_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *TcpNsr_Nodes) GetParentYangName() string { return "tcp-nsr" }

// TcpNsr_Nodes_Node
// Information about a single node
type TcpNsr_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Describing a location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Information about TCP NSR Sessions.
    Session TcpNsr_Nodes_Node_Session

    // Information about TCP NSR Client.
    Client TcpNsr_Nodes_Node_Client

    // Information about TCP NSR Session Sets.
    SessionSet TcpNsr_Nodes_Node_SessionSet

    // Statis Information about TCP NSR connections.
    Statistics TcpNsr_Nodes_Node_Statistics
}

func (node *TcpNsr_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *TcpNsr_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *TcpNsr_Nodes_Node) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "session" { return "Session" }
    if yname == "client" { return "Client" }
    if yname == "session-set" { return "SessionSet" }
    if yname == "statistics" { return "Statistics" }
    return ""
}

func (node *TcpNsr_Nodes_Node) GetSegmentPath() string {
    return "node" + "[id='" + fmt.Sprintf("%v", node.Id) + "']"
}

func (node *TcpNsr_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session" {
        return &node.Session
    }
    if childYangName == "client" {
        return &node.Client
    }
    if childYangName == "session-set" {
        return &node.SessionSet
    }
    if childYangName == "statistics" {
        return &node.Statistics
    }
    return nil
}

func (node *TcpNsr_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session"] = &node.Session
    children["client"] = &node.Client
    children["session-set"] = &node.SessionSet
    children["statistics"] = &node.Statistics
    return children
}

func (node *TcpNsr_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = node.Id
    return leafs
}

func (node *TcpNsr_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *TcpNsr_Nodes_Node) GetYangName() string { return "node" }

func (node *TcpNsr_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *TcpNsr_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *TcpNsr_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *TcpNsr_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *TcpNsr_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *TcpNsr_Nodes_Node) GetParentYangName() string { return "nodes" }

// TcpNsr_Nodes_Node_Session
// Information about TCP NSR Sessions
type TcpNsr_Nodes_Node_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Information about TCP NSR Sessions.
    BriefSessions TcpNsr_Nodes_Node_Session_BriefSessions

    // Table about TCP NSR Sessions details.
    DetailSessions TcpNsr_Nodes_Node_Session_DetailSessions
}

func (session *TcpNsr_Nodes_Node_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *TcpNsr_Nodes_Node_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *TcpNsr_Nodes_Node_Session) GetGoName(yname string) string {
    if yname == "brief-sessions" { return "BriefSessions" }
    if yname == "detail-sessions" { return "DetailSessions" }
    return ""
}

func (session *TcpNsr_Nodes_Node_Session) GetSegmentPath() string {
    return "session"
}

func (session *TcpNsr_Nodes_Node_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-sessions" {
        return &session.BriefSessions
    }
    if childYangName == "detail-sessions" {
        return &session.DetailSessions
    }
    return nil
}

func (session *TcpNsr_Nodes_Node_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["brief-sessions"] = &session.BriefSessions
    children["detail-sessions"] = &session.DetailSessions
    return children
}

func (session *TcpNsr_Nodes_Node_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *TcpNsr_Nodes_Node_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *TcpNsr_Nodes_Node_Session) GetYangName() string { return "session" }

func (session *TcpNsr_Nodes_Node_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *TcpNsr_Nodes_Node_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *TcpNsr_Nodes_Node_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *TcpNsr_Nodes_Node_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *TcpNsr_Nodes_Node_Session) GetParent() types.Entity { return session.parent }

func (session *TcpNsr_Nodes_Node_Session) GetParentYangName() string { return "node" }

// TcpNsr_Nodes_Node_Session_BriefSessions
// Information about TCP NSR Sessions
type TcpNsr_Nodes_Node_Session_BriefSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about NSR Sessions. The type is slice of
    // TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession.
    BriefSession []TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetFilter() yfilter.YFilter { return briefSessions.YFilter }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) SetFilter(yf yfilter.YFilter) { briefSessions.YFilter = yf }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetGoName(yname string) string {
    if yname == "brief-session" { return "BriefSession" }
    return ""
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetSegmentPath() string {
    return "brief-sessions"
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-session" {
        for _, c := range briefSessions.BriefSession {
            if briefSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession{}
        briefSessions.BriefSession = append(briefSessions.BriefSession, child)
        return &briefSessions.BriefSession[len(briefSessions.BriefSession)-1]
    }
    return nil
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefSessions.BriefSession {
        children[briefSessions.BriefSession[i].GetSegmentPath()] = &briefSessions.BriefSession[i]
    }
    return children
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetBundleName() string { return "cisco_ios_xr" }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetYangName() string { return "brief-sessions" }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) SetParent(parent types.Entity) { briefSessions.parent = parent }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetParent() types.Entity { return briefSessions.parent }

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetParentYangName() string { return "session" }

// TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession
// Brief information about NSR Sessions
type TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR Sesison. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address family. The type is AddrFamily.
    AddressFamily interface{}

    // PCB Address. The type is interface{} with range: 0..18446744073709551615.
    Pcb interface{}

    // SSCB Address. The type is interface{} with range: 0..18446744073709551615.
    Sscb interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // VRF Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Is NSR administratively configured?. The type is bool.
    IsAdminConfiguredUp interface{}

    // Is Upstream NSR operational?. The type is NsrStatus.
    IsUsOperationalUp interface{}

    // Is Downstream NSR operational?. The type is NsrStatus.
    IsDsOperationalUp interface{}

    // Is replication limited to receive-path only. The type is bool.
    IsOnlyReceivePathReplication interface{}

    // Local address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress []interface{}

    // Foreign address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ForeignAddress []interface{}
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetFilter() yfilter.YFilter { return briefSession.YFilter }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) SetFilter(yf yfilter.YFilter) { briefSession.YFilter = yf }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "pcb" { return "Pcb" }
    if yname == "sscb" { return "Sscb" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "is-admin-configured-up" { return "IsAdminConfiguredUp" }
    if yname == "is-us-operational-up" { return "IsUsOperationalUp" }
    if yname == "is-ds-operational-up" { return "IsDsOperationalUp" }
    if yname == "is-only-receive-path-replication" { return "IsOnlyReceivePathReplication" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    return ""
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetSegmentPath() string {
    return "brief-session" + "[id='" + fmt.Sprintf("%v", briefSession.Id) + "']"
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = briefSession.Id
    leafs["address-family"] = briefSession.AddressFamily
    leafs["pcb"] = briefSession.Pcb
    leafs["sscb"] = briefSession.Sscb
    leafs["local-port"] = briefSession.LocalPort
    leafs["foreign-port"] = briefSession.ForeignPort
    leafs["vrf-id"] = briefSession.VrfId
    leafs["is-admin-configured-up"] = briefSession.IsAdminConfiguredUp
    leafs["is-us-operational-up"] = briefSession.IsUsOperationalUp
    leafs["is-ds-operational-up"] = briefSession.IsDsOperationalUp
    leafs["is-only-receive-path-replication"] = briefSession.IsOnlyReceivePathReplication
    leafs["local-address"] = briefSession.LocalAddress
    leafs["foreign-address"] = briefSession.ForeignAddress
    return leafs
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetBundleName() string { return "cisco_ios_xr" }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetYangName() string { return "brief-session" }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) SetParent(parent types.Entity) { briefSession.parent = parent }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetParent() types.Entity { return briefSession.parent }

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetParentYangName() string { return "brief-sessions" }

// TcpNsr_Nodes_Node_Session_DetailSessions
// Table about TCP NSR Sessions details
type TcpNsr_Nodes_Node_Session_DetailSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // showing detailed information of NSR Sessions. The type is slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession.
    DetailSession []TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetFilter() yfilter.YFilter { return detailSessions.YFilter }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) SetFilter(yf yfilter.YFilter) { detailSessions.YFilter = yf }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetGoName(yname string) string {
    if yname == "detail-session" { return "DetailSession" }
    return ""
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetSegmentPath() string {
    return "detail-sessions"
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-session" {
        for _, c := range detailSessions.DetailSession {
            if detailSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession{}
        detailSessions.DetailSession = append(detailSessions.DetailSession, child)
        return &detailSessions.DetailSession[len(detailSessions.DetailSession)-1]
    }
    return nil
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detailSessions.DetailSession {
        children[detailSessions.DetailSession[i].GetSegmentPath()] = &detailSessions.DetailSession[i]
    }
    return children
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetBundleName() string { return "cisco_ios_xr" }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetYangName() string { return "detail-sessions" }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) SetParent(parent types.Entity) { detailSessions.parent = parent }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetParent() types.Entity { return detailSessions.parent }

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetParentYangName() string { return "session" }

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession
// showing detailed information of NSR Sessions
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR Sesison. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address family. The type is AddrFamily.
    AddressFamily interface{}

    // PCB Address. The type is interface{} with range: 0..18446744073709551615.
    Pcb interface{}

    // SSCB Address. The type is interface{} with range: 0..18446744073709551615.
    Sscb interface{}

    // Local port. The type is interface{} with range: 0..65535.
    LocalPort interface{}

    // Foreign port. The type is interface{} with range: 0..65535.
    ForeignPort interface{}

    // VRF Id. The type is interface{} with range: 0..4294967295.
    VrfId interface{}

    // Is NSR administratively configured?. The type is bool.
    IsAdminConfiguredUp interface{}

    // Is Upstream NSR operational?. The type is NsrStatus.
    IsUsOperationalUp interface{}

    // Is Downstream NSR operational?. The type is NsrStatus.
    IsDsOperationalUp interface{}

    // Is replication limited to receive-path only. The type is bool.
    IsOnlyReceivePathReplication interface{}

    // Cookie provided by active APP. The type is interface{} with range:
    // 0..18446744073709551615.
    Cookie interface{}

    // Has the session been replicated to standby?. The type is bool.
    IsSessionReplicated interface{}

    // Has the session completed initial-sync?. The type is bool.
    IsSessionSynced interface{}

    // If initial sync is completed, then the FSSN - First Standby Sequence
    // Number. The type is interface{} with range: 0..4294967295.
    FistStandbySequenceNumber interface{}

    // Offset of FSSN in input stream. The type is interface{} with range:
    // 0..4294967295.
    FssnOffset interface{}

    // If NSR is not up, the reason for it. The type is NsrDownReason.
    NsrDownReason interface{}

    // Time at which NSR went down. The type is interface{} with range:
    // 0..4294967295.
    NsrDownTime interface{}

    // ID of the Initial sync operation. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumberOfInitSync interface{}

    // Is initial-sync currently in progress?. The type is bool.
    IsInitSyncInProgress interface{}

    // Is initial sync in the second phase?. The type is bool.
    IsInitSyncSecondPhase interface{}

    // Initial sync failure reason, if any. The type is string.
    InitSyncError interface{}

    // Initial sync failed due to a local error or remote stack. The type is bool.
    IsInitSyncErrorLocal interface{}

    // Time at which the initial sync operation was started (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncStartTime interface{}

    // Time at which the initial sync operation was ended (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncEndTime interface{}

    // Init Sync flags for the session. The type is interface{} with range:
    // 0..4294967295.
    InitSyncFlags interface{}

    // ID of the Initial sync operation. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumberOfInitSyncUpStream interface{}

    // Peer NCD endp handle. The type is interface{} with range:
    // 0..18446744073709551615.
    PeerEndpHdlUpStream interface{}

    // Time at which the initial sync operation was started (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncStartTimeUpStream interface{}

    // Time at which the initial sync operation was ended (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncEndTimeUpStream interface{}

    // FSSN for the upstream partner. The type is interface{} with range:
    // 0..4294967295.
    FistStandbySequenceNumberUpStream interface{}

    // The reason NSR is not up towards the upstream partner. The type is
    // NsrDownReason.
    NsrDownReasonUpStream interface{}

    // Time at which NSR went down. The type is interface{} with range:
    // 0..4294967295.
    NsrDownTimeUpStream interface{}

    // ID of the Initial sync operation. The type is interface{} with range:
    // 0..4294967295.
    SequenceNumberOfInitSyncDownStream interface{}

    // Peer NCD endp handle. The type is interface{} with range:
    // 0..18446744073709551615.
    PeerEndpHdlDownStream interface{}

    // Time at which the initial sync operation was started (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncStartTimeDownStream interface{}

    // Time at which the initial sync operation was ended (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncEndTimeDownStream interface{}

    // FSSN for the upstream partner. The type is interface{} with range:
    // 0..4294967295.
    FistStandbySequenceNumberDownStream interface{}

    // The reason NSR is not up towards the upstream partner. The type is
    // NsrDownReason.
    NsrDownReasonDownStream interface{}

    // Time at which NSR went down. The type is interface{} with range:
    // 0..4294967295.
    NsrDownTimeDownStream interface{}

    // Max number of incoming packets have been held. The type is interface{} with
    // range: -2147483648..2147483647.
    MaxNumberOfHeldPacket interface{}

    // Max number of held incoming packets reaches at. The type is interface{}
    // with range: 0..4294967295.
    MaxNumberOfHeldPacketReachTime interface{}

    // Max number of internal acks have been held. The type is interface{} with
    // range: -2147483648..2147483647.
    MaxNumberOfHeldInternalAck interface{}

    // Max number of held internal acks reaches at. The type is interface{} with
    // range: 0..4294967295.
    MaxNumberOfHeldInternalAckReachTime interface{}

    // Local address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    LocalAddress []interface{}

    // Foreign address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ForeignAddress []interface{}

    // Sesson-set information.
    SetInformation TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation

    // Sequence Number and datalength of each node in hold_pakqueue. The type is
    // slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue.
    PacketHoldQueue []TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue

    // Sequence Number and datalength of each node in hold_iackqueue. The type is
    // slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue.
    InternalAckHoldQueue []TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetFilter() yfilter.YFilter { return detailSession.YFilter }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) SetFilter(yf yfilter.YFilter) { detailSession.YFilter = yf }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "pcb" { return "Pcb" }
    if yname == "sscb" { return "Sscb" }
    if yname == "local-port" { return "LocalPort" }
    if yname == "foreign-port" { return "ForeignPort" }
    if yname == "vrf-id" { return "VrfId" }
    if yname == "is-admin-configured-up" { return "IsAdminConfiguredUp" }
    if yname == "is-us-operational-up" { return "IsUsOperationalUp" }
    if yname == "is-ds-operational-up" { return "IsDsOperationalUp" }
    if yname == "is-only-receive-path-replication" { return "IsOnlyReceivePathReplication" }
    if yname == "cookie" { return "Cookie" }
    if yname == "is-session-replicated" { return "IsSessionReplicated" }
    if yname == "is-session-synced" { return "IsSessionSynced" }
    if yname == "fist-standby-sequence-number" { return "FistStandbySequenceNumber" }
    if yname == "fssn-offset" { return "FssnOffset" }
    if yname == "nsr-down-reason" { return "NsrDownReason" }
    if yname == "nsr-down-time" { return "NsrDownTime" }
    if yname == "sequence-number-of-init-sync" { return "SequenceNumberOfInitSync" }
    if yname == "is-init-sync-in-progress" { return "IsInitSyncInProgress" }
    if yname == "is-init-sync-second-phase" { return "IsInitSyncSecondPhase" }
    if yname == "init-sync-error" { return "InitSyncError" }
    if yname == "is-init-sync-error-local" { return "IsInitSyncErrorLocal" }
    if yname == "init-sync-start-time" { return "InitSyncStartTime" }
    if yname == "init-sync-end-time" { return "InitSyncEndTime" }
    if yname == "init-sync-flags" { return "InitSyncFlags" }
    if yname == "sequence-number-of-init-sync-up-stream" { return "SequenceNumberOfInitSyncUpStream" }
    if yname == "peer-endp-hdl-up-stream" { return "PeerEndpHdlUpStream" }
    if yname == "init-sync-start-time-up-stream" { return "InitSyncStartTimeUpStream" }
    if yname == "init-sync-end-time-up-stream" { return "InitSyncEndTimeUpStream" }
    if yname == "fist-standby-sequence-number-up-stream" { return "FistStandbySequenceNumberUpStream" }
    if yname == "nsr-down-reason-up-stream" { return "NsrDownReasonUpStream" }
    if yname == "nsr-down-time-up-stream" { return "NsrDownTimeUpStream" }
    if yname == "sequence-number-of-init-sync-down-stream" { return "SequenceNumberOfInitSyncDownStream" }
    if yname == "peer-endp-hdl-down-stream" { return "PeerEndpHdlDownStream" }
    if yname == "init-sync-start-time-down-stream" { return "InitSyncStartTimeDownStream" }
    if yname == "init-sync-end-time-down-stream" { return "InitSyncEndTimeDownStream" }
    if yname == "fist-standby-sequence-number-down-stream" { return "FistStandbySequenceNumberDownStream" }
    if yname == "nsr-down-reason-down-stream" { return "NsrDownReasonDownStream" }
    if yname == "nsr-down-time-down-stream" { return "NsrDownTimeDownStream" }
    if yname == "max-number-of-held-packet" { return "MaxNumberOfHeldPacket" }
    if yname == "max-number-of-held-packet-reach-time" { return "MaxNumberOfHeldPacketReachTime" }
    if yname == "max-number-of-held-internal-ack" { return "MaxNumberOfHeldInternalAck" }
    if yname == "max-number-of-held-internal-ack-reach-time" { return "MaxNumberOfHeldInternalAckReachTime" }
    if yname == "local-address" { return "LocalAddress" }
    if yname == "foreign-address" { return "ForeignAddress" }
    if yname == "set-information" { return "SetInformation" }
    if yname == "packet-hold-queue" { return "PacketHoldQueue" }
    if yname == "internal-ack-hold-queue" { return "InternalAckHoldQueue" }
    return ""
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetSegmentPath() string {
    return "detail-session" + "[id='" + fmt.Sprintf("%v", detailSession.Id) + "']"
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "set-information" {
        return &detailSession.SetInformation
    }
    if childYangName == "packet-hold-queue" {
        for _, c := range detailSession.PacketHoldQueue {
            if detailSession.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue{}
        detailSession.PacketHoldQueue = append(detailSession.PacketHoldQueue, child)
        return &detailSession.PacketHoldQueue[len(detailSession.PacketHoldQueue)-1]
    }
    if childYangName == "internal-ack-hold-queue" {
        for _, c := range detailSession.InternalAckHoldQueue {
            if detailSession.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue{}
        detailSession.InternalAckHoldQueue = append(detailSession.InternalAckHoldQueue, child)
        return &detailSession.InternalAckHoldQueue[len(detailSession.InternalAckHoldQueue)-1]
    }
    return nil
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["set-information"] = &detailSession.SetInformation
    for i := range detailSession.PacketHoldQueue {
        children[detailSession.PacketHoldQueue[i].GetSegmentPath()] = &detailSession.PacketHoldQueue[i]
    }
    for i := range detailSession.InternalAckHoldQueue {
        children[detailSession.InternalAckHoldQueue[i].GetSegmentPath()] = &detailSession.InternalAckHoldQueue[i]
    }
    return children
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = detailSession.Id
    leafs["address-family"] = detailSession.AddressFamily
    leafs["pcb"] = detailSession.Pcb
    leafs["sscb"] = detailSession.Sscb
    leafs["local-port"] = detailSession.LocalPort
    leafs["foreign-port"] = detailSession.ForeignPort
    leafs["vrf-id"] = detailSession.VrfId
    leafs["is-admin-configured-up"] = detailSession.IsAdminConfiguredUp
    leafs["is-us-operational-up"] = detailSession.IsUsOperationalUp
    leafs["is-ds-operational-up"] = detailSession.IsDsOperationalUp
    leafs["is-only-receive-path-replication"] = detailSession.IsOnlyReceivePathReplication
    leafs["cookie"] = detailSession.Cookie
    leafs["is-session-replicated"] = detailSession.IsSessionReplicated
    leafs["is-session-synced"] = detailSession.IsSessionSynced
    leafs["fist-standby-sequence-number"] = detailSession.FistStandbySequenceNumber
    leafs["fssn-offset"] = detailSession.FssnOffset
    leafs["nsr-down-reason"] = detailSession.NsrDownReason
    leafs["nsr-down-time"] = detailSession.NsrDownTime
    leafs["sequence-number-of-init-sync"] = detailSession.SequenceNumberOfInitSync
    leafs["is-init-sync-in-progress"] = detailSession.IsInitSyncInProgress
    leafs["is-init-sync-second-phase"] = detailSession.IsInitSyncSecondPhase
    leafs["init-sync-error"] = detailSession.InitSyncError
    leafs["is-init-sync-error-local"] = detailSession.IsInitSyncErrorLocal
    leafs["init-sync-start-time"] = detailSession.InitSyncStartTime
    leafs["init-sync-end-time"] = detailSession.InitSyncEndTime
    leafs["init-sync-flags"] = detailSession.InitSyncFlags
    leafs["sequence-number-of-init-sync-up-stream"] = detailSession.SequenceNumberOfInitSyncUpStream
    leafs["peer-endp-hdl-up-stream"] = detailSession.PeerEndpHdlUpStream
    leafs["init-sync-start-time-up-stream"] = detailSession.InitSyncStartTimeUpStream
    leafs["init-sync-end-time-up-stream"] = detailSession.InitSyncEndTimeUpStream
    leafs["fist-standby-sequence-number-up-stream"] = detailSession.FistStandbySequenceNumberUpStream
    leafs["nsr-down-reason-up-stream"] = detailSession.NsrDownReasonUpStream
    leafs["nsr-down-time-up-stream"] = detailSession.NsrDownTimeUpStream
    leafs["sequence-number-of-init-sync-down-stream"] = detailSession.SequenceNumberOfInitSyncDownStream
    leafs["peer-endp-hdl-down-stream"] = detailSession.PeerEndpHdlDownStream
    leafs["init-sync-start-time-down-stream"] = detailSession.InitSyncStartTimeDownStream
    leafs["init-sync-end-time-down-stream"] = detailSession.InitSyncEndTimeDownStream
    leafs["fist-standby-sequence-number-down-stream"] = detailSession.FistStandbySequenceNumberDownStream
    leafs["nsr-down-reason-down-stream"] = detailSession.NsrDownReasonDownStream
    leafs["nsr-down-time-down-stream"] = detailSession.NsrDownTimeDownStream
    leafs["max-number-of-held-packet"] = detailSession.MaxNumberOfHeldPacket
    leafs["max-number-of-held-packet-reach-time"] = detailSession.MaxNumberOfHeldPacketReachTime
    leafs["max-number-of-held-internal-ack"] = detailSession.MaxNumberOfHeldInternalAck
    leafs["max-number-of-held-internal-ack-reach-time"] = detailSession.MaxNumberOfHeldInternalAckReachTime
    leafs["local-address"] = detailSession.LocalAddress
    leafs["foreign-address"] = detailSession.ForeignAddress
    return leafs
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetBundleName() string { return "cisco_ios_xr" }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetYangName() string { return "detail-session" }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) SetParent(parent types.Entity) { detailSession.parent = parent }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetParent() types.Entity { return detailSession.parent }

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetParentYangName() string { return "detail-sessions" }

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation
// Sesson-set information
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address of the Session Set Control Block. The type is interface{} with
    // range: 0..18446744073709551615.
    Sscb interface{}

    // PID of the Client that owns this Session-set. The type is interface{} with
    // range: 0..4294967295.
    Pid interface{}

    // the name of Clinet that owns this Session-set. The type is string.
    ClientName interface{}

    // Instance of the Client that owns this Session-set. The type is interface{}
    // with range: 0..4294967295.
    ClientInstance interface{}

    // ID of this Session-set. The type is interface{} with range: 0..4294967295.
    SetId interface{}

    // TCP role for this set?. The type is interface{} with range: 0..4294967295.
    SsoRole interface{}

    // Session-set mode. The type is interface{} with range: 0..4294967295.
    Mode interface{}

    // Address Family of the sessions in this set. The type is AddrFamily.
    AddressFamily interface{}

    // Well Known Port of the client. The type is interface{} with range:
    // 0..65535.
    WellKnownPort interface{}

    // Local node of this set. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    LocalNode interface{}

    // Instance of the client application on the local node. The type is
    // interface{} with range: 0..4294967295.
    LocalInstance interface{}

    // The node protecting this set. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    ProtectNode interface{}

    // Instance of the client application on the protection node. The type is
    // interface{} with range: 0..4294967295.
    ProtectInstance interface{}

    // Number of Sessions in the set. The type is interface{} with range:
    // 0..4294967295.
    NumberOfSessions interface{}

    // How many sessions are synced with upstream partner. The type is interface{}
    // with range: 0..4294967295.
    NumberOfSyncedSessionsUpStream interface{}

    // How many sessions are synced with downstream partner. The type is
    // interface{} with range: 0..4294967295.
    NumberOfSyncedSessionsDownStream interface{}

    // Is an initial sync in progress currently?. The type is bool.
    IsInitSyncInProgress interface{}

    // Is the SSCB ready for another initial sync?. The type is bool.
    IsSscbInitSyncReady interface{}
}

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetFilter() yfilter.YFilter { return setInformation.YFilter }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) SetFilter(yf yfilter.YFilter) { setInformation.YFilter = yf }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetGoName(yname string) string {
    if yname == "sscb" { return "Sscb" }
    if yname == "pid" { return "Pid" }
    if yname == "client-name" { return "ClientName" }
    if yname == "client-instance" { return "ClientInstance" }
    if yname == "set-id" { return "SetId" }
    if yname == "sso-role" { return "SsoRole" }
    if yname == "mode" { return "Mode" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "well-known-port" { return "WellKnownPort" }
    if yname == "local-node" { return "LocalNode" }
    if yname == "local-instance" { return "LocalInstance" }
    if yname == "protect-node" { return "ProtectNode" }
    if yname == "protect-instance" { return "ProtectInstance" }
    if yname == "number-of-sessions" { return "NumberOfSessions" }
    if yname == "number-of-synced-sessions-up-stream" { return "NumberOfSyncedSessionsUpStream" }
    if yname == "number-of-synced-sessions-down-stream" { return "NumberOfSyncedSessionsDownStream" }
    if yname == "is-init-sync-in-progress" { return "IsInitSyncInProgress" }
    if yname == "is-sscb-init-sync-ready" { return "IsSscbInitSyncReady" }
    return ""
}

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetSegmentPath() string {
    return "set-information"
}

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sscb"] = setInformation.Sscb
    leafs["pid"] = setInformation.Pid
    leafs["client-name"] = setInformation.ClientName
    leafs["client-instance"] = setInformation.ClientInstance
    leafs["set-id"] = setInformation.SetId
    leafs["sso-role"] = setInformation.SsoRole
    leafs["mode"] = setInformation.Mode
    leafs["address-family"] = setInformation.AddressFamily
    leafs["well-known-port"] = setInformation.WellKnownPort
    leafs["local-node"] = setInformation.LocalNode
    leafs["local-instance"] = setInformation.LocalInstance
    leafs["protect-node"] = setInformation.ProtectNode
    leafs["protect-instance"] = setInformation.ProtectInstance
    leafs["number-of-sessions"] = setInformation.NumberOfSessions
    leafs["number-of-synced-sessions-up-stream"] = setInformation.NumberOfSyncedSessionsUpStream
    leafs["number-of-synced-sessions-down-stream"] = setInformation.NumberOfSyncedSessionsDownStream
    leafs["is-init-sync-in-progress"] = setInformation.IsInitSyncInProgress
    leafs["is-sscb-init-sync-ready"] = setInformation.IsSscbInitSyncReady
    return leafs
}

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetBundleName() string { return "cisco_ios_xr" }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetYangName() string { return "set-information" }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) SetParent(parent types.Entity) { setInformation.parent = parent }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetParent() types.Entity { return setInformation.parent }

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetParentYangName() string { return "detail-session" }

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue
// Sequence Number and datalength of each node in
// hold_pakqueue
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Data Length. The type is interface{} with range: 0..4294967295.
    DataLength interface{}

    // Ack Number. The type is interface{} with range: 0..4294967295.
    AcknoledgementNumber interface{}
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetFilter() yfilter.YFilter { return packetHoldQueue.YFilter }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) SetFilter(yf yfilter.YFilter) { packetHoldQueue.YFilter = yf }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "data-length" { return "DataLength" }
    if yname == "acknoledgement-number" { return "AcknoledgementNumber" }
    return ""
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetSegmentPath() string {
    return "packet-hold-queue"
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = packetHoldQueue.SequenceNumber
    leafs["data-length"] = packetHoldQueue.DataLength
    leafs["acknoledgement-number"] = packetHoldQueue.AcknoledgementNumber
    return leafs
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetBundleName() string { return "cisco_ios_xr" }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetYangName() string { return "packet-hold-queue" }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) SetParent(parent types.Entity) { packetHoldQueue.parent = parent }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetParent() types.Entity { return packetHoldQueue.parent }

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetParentYangName() string { return "detail-session" }

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue
// Sequence Number and datalength of each node in
// hold_iackqueue
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Data Length. The type is interface{} with range: 0..4294967295.
    DataLength interface{}

    // Ack Number. The type is interface{} with range: 0..4294967295.
    AcknoledgementNumber interface{}
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetFilter() yfilter.YFilter { return internalAckHoldQueue.YFilter }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) SetFilter(yf yfilter.YFilter) { internalAckHoldQueue.YFilter = yf }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetGoName(yname string) string {
    if yname == "sequence-number" { return "SequenceNumber" }
    if yname == "data-length" { return "DataLength" }
    if yname == "acknoledgement-number" { return "AcknoledgementNumber" }
    return ""
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetSegmentPath() string {
    return "internal-ack-hold-queue"
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sequence-number"] = internalAckHoldQueue.SequenceNumber
    leafs["data-length"] = internalAckHoldQueue.DataLength
    leafs["acknoledgement-number"] = internalAckHoldQueue.AcknoledgementNumber
    return leafs
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetBundleName() string { return "cisco_ios_xr" }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetYangName() string { return "internal-ack-hold-queue" }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) SetParent(parent types.Entity) { internalAckHoldQueue.parent = parent }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetParent() types.Entity { return internalAckHoldQueue.parent }

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetParentYangName() string { return "detail-session" }

// TcpNsr_Nodes_Node_Client
// Information about TCP NSR Client
type TcpNsr_Nodes_Node_Client struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table about TCP NSR Client details.
    DetailClients TcpNsr_Nodes_Node_Client_DetailClients

    // Information about TCP NSR Client.
    BriefClients TcpNsr_Nodes_Node_Client_BriefClients
}

func (client *TcpNsr_Nodes_Node_Client) GetFilter() yfilter.YFilter { return client.YFilter }

func (client *TcpNsr_Nodes_Node_Client) SetFilter(yf yfilter.YFilter) { client.YFilter = yf }

func (client *TcpNsr_Nodes_Node_Client) GetGoName(yname string) string {
    if yname == "detail-clients" { return "DetailClients" }
    if yname == "brief-clients" { return "BriefClients" }
    return ""
}

func (client *TcpNsr_Nodes_Node_Client) GetSegmentPath() string {
    return "client"
}

func (client *TcpNsr_Nodes_Node_Client) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-clients" {
        return &client.DetailClients
    }
    if childYangName == "brief-clients" {
        return &client.BriefClients
    }
    return nil
}

func (client *TcpNsr_Nodes_Node_Client) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail-clients"] = &client.DetailClients
    children["brief-clients"] = &client.BriefClients
    return children
}

func (client *TcpNsr_Nodes_Node_Client) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (client *TcpNsr_Nodes_Node_Client) GetBundleName() string { return "cisco_ios_xr" }

func (client *TcpNsr_Nodes_Node_Client) GetYangName() string { return "client" }

func (client *TcpNsr_Nodes_Node_Client) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (client *TcpNsr_Nodes_Node_Client) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (client *TcpNsr_Nodes_Node_Client) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (client *TcpNsr_Nodes_Node_Client) SetParent(parent types.Entity) { client.parent = parent }

func (client *TcpNsr_Nodes_Node_Client) GetParent() types.Entity { return client.parent }

func (client *TcpNsr_Nodes_Node_Client) GetParentYangName() string { return "node" }

// TcpNsr_Nodes_Node_Client_DetailClients
// Table about TCP NSR Client details
type TcpNsr_Nodes_Node_Client_DetailClients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // showing detailed information of NSR Clients. The type is slice of
    // TcpNsr_Nodes_Node_Client_DetailClients_DetailClient.
    DetailClient []TcpNsr_Nodes_Node_Client_DetailClients_DetailClient
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetFilter() yfilter.YFilter { return detailClients.YFilter }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) SetFilter(yf yfilter.YFilter) { detailClients.YFilter = yf }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetGoName(yname string) string {
    if yname == "detail-client" { return "DetailClient" }
    return ""
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetSegmentPath() string {
    return "detail-clients"
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-client" {
        for _, c := range detailClients.DetailClient {
            if detailClients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Client_DetailClients_DetailClient{}
        detailClients.DetailClient = append(detailClients.DetailClient, child)
        return &detailClients.DetailClient[len(detailClients.DetailClient)-1]
    }
    return nil
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detailClients.DetailClient {
        children[detailClients.DetailClient[i].GetSegmentPath()] = &detailClients.DetailClient[i]
    }
    return children
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetBundleName() string { return "cisco_ios_xr" }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetYangName() string { return "detail-clients" }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) SetParent(parent types.Entity) { detailClients.parent = parent }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetParent() types.Entity { return detailClients.parent }

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetParentYangName() string { return "client" }

// TcpNsr_Nodes_Node_Client_DetailClients_DetailClient
// showing detailed information of NSR Clients
type TcpNsr_Nodes_Node_Client_DetailClients_DetailClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR client. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address of the Client Control Block. The type is interface{} with range:
    // 0..18446744073709551615.
    Ccb interface{}

    // PID of the Client. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Proc name of Clinet. The type is string.
    ProcessName interface{}

    // JOb ID of Client. The type is interface{} with range:
    // -2147483648..2147483647.
    JobId interface{}

    // Instance of the Client. The type is interface{} with range: 0..4294967295.
    Instance interface{}

    // Number of Sets owned by this client . The type is interface{} with range:
    // 0..4294967295.
    NumberofSets interface{}

    // Number of sessions owned by this client . The type is interface{} with
    // range: 0..4294967295.
    NumberOfSessions interface{}

    // Number of sessions with NSR up. The type is interface{} with range:
    // 0..4294967295.
    NumberOfUpSessions interface{}

    // Time of connect (in seconds since 1st Jan 1970 00:00:00). The type is
    // interface{} with range: 0..4294967295. Units are second.
    ConnectedAt interface{}

    // Registered with TCP for notifications?. The type is bool.
    IsNotificationRegistered interface{}
}

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetFilter() yfilter.YFilter { return detailClient.YFilter }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) SetFilter(yf yfilter.YFilter) { detailClient.YFilter = yf }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "ccb" { return "Ccb" }
    if yname == "pid" { return "Pid" }
    if yname == "process-name" { return "ProcessName" }
    if yname == "job-id" { return "JobId" }
    if yname == "instance" { return "Instance" }
    if yname == "numberof-sets" { return "NumberofSets" }
    if yname == "number-of-sessions" { return "NumberOfSessions" }
    if yname == "number-of-up-sessions" { return "NumberOfUpSessions" }
    if yname == "connected-at" { return "ConnectedAt" }
    if yname == "is-notification-registered" { return "IsNotificationRegistered" }
    return ""
}

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetSegmentPath() string {
    return "detail-client" + "[id='" + fmt.Sprintf("%v", detailClient.Id) + "']"
}

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = detailClient.Id
    leafs["ccb"] = detailClient.Ccb
    leafs["pid"] = detailClient.Pid
    leafs["process-name"] = detailClient.ProcessName
    leafs["job-id"] = detailClient.JobId
    leafs["instance"] = detailClient.Instance
    leafs["numberof-sets"] = detailClient.NumberofSets
    leafs["number-of-sessions"] = detailClient.NumberOfSessions
    leafs["number-of-up-sessions"] = detailClient.NumberOfUpSessions
    leafs["connected-at"] = detailClient.ConnectedAt
    leafs["is-notification-registered"] = detailClient.IsNotificationRegistered
    return leafs
}

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetBundleName() string { return "cisco_ios_xr" }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetYangName() string { return "detail-client" }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) SetParent(parent types.Entity) { detailClient.parent = parent }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetParent() types.Entity { return detailClient.parent }

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetParentYangName() string { return "detail-clients" }

// TcpNsr_Nodes_Node_Client_BriefClients
// Information about TCP NSR Client
type TcpNsr_Nodes_Node_Client_BriefClients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about NSR Client. The type is slice of
    // TcpNsr_Nodes_Node_Client_BriefClients_BriefClient.
    BriefClient []TcpNsr_Nodes_Node_Client_BriefClients_BriefClient
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetFilter() yfilter.YFilter { return briefClients.YFilter }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) SetFilter(yf yfilter.YFilter) { briefClients.YFilter = yf }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetGoName(yname string) string {
    if yname == "brief-client" { return "BriefClient" }
    return ""
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetSegmentPath() string {
    return "brief-clients"
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-client" {
        for _, c := range briefClients.BriefClient {
            if briefClients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Client_BriefClients_BriefClient{}
        briefClients.BriefClient = append(briefClients.BriefClient, child)
        return &briefClients.BriefClient[len(briefClients.BriefClient)-1]
    }
    return nil
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefClients.BriefClient {
        children[briefClients.BriefClient[i].GetSegmentPath()] = &briefClients.BriefClient[i]
    }
    return children
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetBundleName() string { return "cisco_ios_xr" }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetYangName() string { return "brief-clients" }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) SetParent(parent types.Entity) { briefClients.parent = parent }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetParent() types.Entity { return briefClients.parent }

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetParentYangName() string { return "client" }

// TcpNsr_Nodes_Node_Client_BriefClients_BriefClient
// Brief information about NSR Client
type TcpNsr_Nodes_Node_Client_BriefClients_BriefClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR client. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address of the Client Control Block. The type is interface{} with range:
    // 0..18446744073709551615.
    Ccb interface{}

    // PID of the Client. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Proc name of Clinet. The type is string.
    ProcessName interface{}

    // JOb ID of Client. The type is interface{} with range:
    // -2147483648..2147483647.
    JobId interface{}

    // Instance of the Client. The type is interface{} with range: 0..4294967295.
    Instance interface{}

    // Number of Sets owned by this client . The type is interface{} with range:
    // 0..4294967295.
    NumberofSets interface{}

    // Number of sessions owned by this client . The type is interface{} with
    // range: 0..4294967295.
    NumberOfSessions interface{}

    // Number of sessions with NSR up . The type is interface{} with range:
    // 0..4294967295.
    NumberOfUpSessions interface{}
}

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetFilter() yfilter.YFilter { return briefClient.YFilter }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) SetFilter(yf yfilter.YFilter) { briefClient.YFilter = yf }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "ccb" { return "Ccb" }
    if yname == "pid" { return "Pid" }
    if yname == "process-name" { return "ProcessName" }
    if yname == "job-id" { return "JobId" }
    if yname == "instance" { return "Instance" }
    if yname == "numberof-sets" { return "NumberofSets" }
    if yname == "number-of-sessions" { return "NumberOfSessions" }
    if yname == "number-of-up-sessions" { return "NumberOfUpSessions" }
    return ""
}

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetSegmentPath() string {
    return "brief-client" + "[id='" + fmt.Sprintf("%v", briefClient.Id) + "']"
}

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = briefClient.Id
    leafs["ccb"] = briefClient.Ccb
    leafs["pid"] = briefClient.Pid
    leafs["process-name"] = briefClient.ProcessName
    leafs["job-id"] = briefClient.JobId
    leafs["instance"] = briefClient.Instance
    leafs["numberof-sets"] = briefClient.NumberofSets
    leafs["number-of-sessions"] = briefClient.NumberOfSessions
    leafs["number-of-up-sessions"] = briefClient.NumberOfUpSessions
    return leafs
}

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetBundleName() string { return "cisco_ios_xr" }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetYangName() string { return "brief-client" }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) SetParent(parent types.Entity) { briefClient.parent = parent }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetParent() types.Entity { return briefClient.parent }

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetParentYangName() string { return "brief-clients" }

// TcpNsr_Nodes_Node_SessionSet
// Information about TCP NSR Session Sets
type TcpNsr_Nodes_Node_SessionSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table about TCP NSR Session Sets details.
    DetailSets TcpNsr_Nodes_Node_SessionSet_DetailSets

    // Information about TCP NSR Session Sets.
    BriefSets TcpNsr_Nodes_Node_SessionSet_BriefSets
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetFilter() yfilter.YFilter { return sessionSet.YFilter }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) SetFilter(yf yfilter.YFilter) { sessionSet.YFilter = yf }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetGoName(yname string) string {
    if yname == "detail-sets" { return "DetailSets" }
    if yname == "brief-sets" { return "BriefSets" }
    return ""
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetSegmentPath() string {
    return "session-set"
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-sets" {
        return &sessionSet.DetailSets
    }
    if childYangName == "brief-sets" {
        return &sessionSet.BriefSets
    }
    return nil
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detail-sets"] = &sessionSet.DetailSets
    children["brief-sets"] = &sessionSet.BriefSets
    return children
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetBundleName() string { return "cisco_ios_xr" }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetYangName() string { return "session-set" }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) SetParent(parent types.Entity) { sessionSet.parent = parent }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetParent() types.Entity { return sessionSet.parent }

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetParentYangName() string { return "node" }

// TcpNsr_Nodes_Node_SessionSet_DetailSets
// Table about TCP NSR Session Sets details
type TcpNsr_Nodes_Node_SessionSet_DetailSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // showing detailed information of NSR Session Sets. The type is slice of
    // TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet.
    DetailSet []TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetFilter() yfilter.YFilter { return detailSets.YFilter }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) SetFilter(yf yfilter.YFilter) { detailSets.YFilter = yf }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetGoName(yname string) string {
    if yname == "detail-set" { return "DetailSet" }
    return ""
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetSegmentPath() string {
    return "detail-sets"
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detail-set" {
        for _, c := range detailSets.DetailSet {
            if detailSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet{}
        detailSets.DetailSet = append(detailSets.DetailSet, child)
        return &detailSets.DetailSet[len(detailSets.DetailSet)-1]
    }
    return nil
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range detailSets.DetailSet {
        children[detailSets.DetailSet[i].GetSegmentPath()] = &detailSets.DetailSet[i]
    }
    return children
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetBundleName() string { return "cisco_ios_xr" }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetYangName() string { return "detail-sets" }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) SetParent(parent types.Entity) { detailSets.parent = parent }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetParent() types.Entity { return detailSets.parent }

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetParentYangName() string { return "session-set" }

// TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet
// showing detailed information of NSR Session
// Sets
type TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR Sesison Set. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address of the Session Set Control Block. The type is interface{} with
    // range: 0..18446744073709551615.
    Sscb interface{}

    // PID of the Client that owns this Session-set. The type is interface{} with
    // range: 0..4294967295.
    Pid interface{}

    // ID of this Session-set. The type is interface{} with range: 0..4294967295.
    SetId interface{}

    // TCP role for this set?. The type is interface{} with range: 0..4294967295.
    SsoRole interface{}

    // Session-set mode. The type is interface{} with range: 0..4294967295.
    Mode interface{}

    // Address Family of the sessions in this set. The type is AddrFamily.
    AddressFamily interface{}

    // Well Known Port of the client. The type is interface{} with range:
    // 0..65535.
    WellKnownPort interface{}

    // Local node of this set. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    LocalNode interface{}

    // Instance of the client application on the local node. The type is
    // interface{} with range: 0..4294967295.
    LocalInstance interface{}

    // The node protecting this set. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    ProtectNode interface{}

    // Instance of the client application on the protection node. The type is
    // interface{} with range: 0..4294967295.
    ProtectInstance interface{}

    // Number of Sessions in the set. The type is interface{} with range:
    // 0..4294967295.
    NumberOfSessions interface{}

    // How many sessions are synced with upstream partner. The type is interface{}
    // with range: 0..4294967295.
    NumberOfSyncedSessionsUpStream interface{}

    // How many sessions are synced with downstream partner. The type is
    // interface{} with range: 0..4294967295.
    NumberOfSyncedSessionsDownStream interface{}

    // Is an initial sync in progress currently?. The type is bool.
    IsInitSyncInProgress interface{}

    // Is initial sync in the second phase?. The type is bool.
    IsInitSyncSecondPhase interface{}

    // ID of the current or the last initial sync operation. The type is
    // interface{} with range: 0..4294967295.
    SequenceNumberOfInitSync interface{}

    // Time left on the initial sync timer. The type is interface{} with range:
    // 0..4294967295.
    InitSyncTimer interface{}

    // Number of sessions being synced as part of the current initial sync
    // operation. The type is interface{} with range: 0..4294967295.
    TotalNumberOfInitSyncSessions interface{}

    // Number of sessions that are synced as part of the current initial sync
    // operation. The type is interface{} with range: 0..4294967295.
    NumberOfInitSyncedSessions interface{}

    // Number of sessions that failed to sync as part of the current initial sync
    // operation. The type is interface{} with range: 0..4294967295.
    NumberOfSessionsInitSyncFailed interface{}

    // Initial sync failure reason, if any. The type is string.
    InitSyncError interface{}

    // Initial sync failed due to a local error or remote stack. The type is bool.
    IsInitSyncErrorLocal interface{}

    // Time at which last or current initial sync operation was started (in
    // seconds since 1st Jan 1970 00:00:00). The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitSyncStartTime interface{}

    // Time at which the last initial sync operation was ended (in seconds since
    // 1st Jan 1970 00:00 :00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncEndTime interface{}

    // Is the SSCB ready for another initial sync?. The type is bool.
    IsSscbInitSyncReady interface{}

    // Time at which the session was ready for initial sync last (in seconds since
    // 1st Jan 1970 00:00 :00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    InitSyncReadyStartTime interface{}

    // Time at which the session set last went not-ready for initial sync (in
    // seconds since 1st Jan 1970 00:00:00). The type is interface{} with range:
    // 0..4294967295. Units are second.
    InitSyncReadyEndTime interface{}

    // Time at which NSR was last reset on the session set (in seconds since 1st
    // Jan 1970 00:00:00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    NsrResetTime interface{}

    // Is an audit in progress currently?. The type is bool.
    IsAuditInProgress interface{}

    // ID of the current or the last audit operation. The type is interface{} with
    // range: 0..4294967295.
    AuditSeqNumber interface{}

    // Time at which last or current audit operation was started (in seconds since
    // 1st Jan 1970 00:00 :00). The type is interface{} with range: 0..4294967295.
    // Units are second.
    AuditStartTime interface{}

    // Time at which the last audit operation was ended (in seconds since 1st Jan
    // 1970 00:00:00). The type is interface{} with range: 0..4294967295. Units
    // are second.
    AuditEndTime interface{}
}

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetFilter() yfilter.YFilter { return detailSet.YFilter }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) SetFilter(yf yfilter.YFilter) { detailSet.YFilter = yf }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "sscb" { return "Sscb" }
    if yname == "pid" { return "Pid" }
    if yname == "set-id" { return "SetId" }
    if yname == "sso-role" { return "SsoRole" }
    if yname == "mode" { return "Mode" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "well-known-port" { return "WellKnownPort" }
    if yname == "local-node" { return "LocalNode" }
    if yname == "local-instance" { return "LocalInstance" }
    if yname == "protect-node" { return "ProtectNode" }
    if yname == "protect-instance" { return "ProtectInstance" }
    if yname == "number-of-sessions" { return "NumberOfSessions" }
    if yname == "number-of-synced-sessions-up-stream" { return "NumberOfSyncedSessionsUpStream" }
    if yname == "number-of-synced-sessions-down-stream" { return "NumberOfSyncedSessionsDownStream" }
    if yname == "is-init-sync-in-progress" { return "IsInitSyncInProgress" }
    if yname == "is-init-sync-second-phase" { return "IsInitSyncSecondPhase" }
    if yname == "sequence-number-of-init-sync" { return "SequenceNumberOfInitSync" }
    if yname == "init-sync-timer" { return "InitSyncTimer" }
    if yname == "total-number-of-init-sync-sessions" { return "TotalNumberOfInitSyncSessions" }
    if yname == "number-of-init-synced-sessions" { return "NumberOfInitSyncedSessions" }
    if yname == "number-of-sessions-init-sync-failed" { return "NumberOfSessionsInitSyncFailed" }
    if yname == "init-sync-error" { return "InitSyncError" }
    if yname == "is-init-sync-error-local" { return "IsInitSyncErrorLocal" }
    if yname == "init-sync-start-time" { return "InitSyncStartTime" }
    if yname == "init-sync-end-time" { return "InitSyncEndTime" }
    if yname == "is-sscb-init-sync-ready" { return "IsSscbInitSyncReady" }
    if yname == "init-sync-ready-start-time" { return "InitSyncReadyStartTime" }
    if yname == "init-sync-ready-end-time" { return "InitSyncReadyEndTime" }
    if yname == "nsr-reset-time" { return "NsrResetTime" }
    if yname == "is-audit-in-progress" { return "IsAuditInProgress" }
    if yname == "audit-seq-number" { return "AuditSeqNumber" }
    if yname == "audit-start-time" { return "AuditStartTime" }
    if yname == "audit-end-time" { return "AuditEndTime" }
    return ""
}

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetSegmentPath() string {
    return "detail-set" + "[id='" + fmt.Sprintf("%v", detailSet.Id) + "']"
}

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = detailSet.Id
    leafs["sscb"] = detailSet.Sscb
    leafs["pid"] = detailSet.Pid
    leafs["set-id"] = detailSet.SetId
    leafs["sso-role"] = detailSet.SsoRole
    leafs["mode"] = detailSet.Mode
    leafs["address-family"] = detailSet.AddressFamily
    leafs["well-known-port"] = detailSet.WellKnownPort
    leafs["local-node"] = detailSet.LocalNode
    leafs["local-instance"] = detailSet.LocalInstance
    leafs["protect-node"] = detailSet.ProtectNode
    leafs["protect-instance"] = detailSet.ProtectInstance
    leafs["number-of-sessions"] = detailSet.NumberOfSessions
    leafs["number-of-synced-sessions-up-stream"] = detailSet.NumberOfSyncedSessionsUpStream
    leafs["number-of-synced-sessions-down-stream"] = detailSet.NumberOfSyncedSessionsDownStream
    leafs["is-init-sync-in-progress"] = detailSet.IsInitSyncInProgress
    leafs["is-init-sync-second-phase"] = detailSet.IsInitSyncSecondPhase
    leafs["sequence-number-of-init-sync"] = detailSet.SequenceNumberOfInitSync
    leafs["init-sync-timer"] = detailSet.InitSyncTimer
    leafs["total-number-of-init-sync-sessions"] = detailSet.TotalNumberOfInitSyncSessions
    leafs["number-of-init-synced-sessions"] = detailSet.NumberOfInitSyncedSessions
    leafs["number-of-sessions-init-sync-failed"] = detailSet.NumberOfSessionsInitSyncFailed
    leafs["init-sync-error"] = detailSet.InitSyncError
    leafs["is-init-sync-error-local"] = detailSet.IsInitSyncErrorLocal
    leafs["init-sync-start-time"] = detailSet.InitSyncStartTime
    leafs["init-sync-end-time"] = detailSet.InitSyncEndTime
    leafs["is-sscb-init-sync-ready"] = detailSet.IsSscbInitSyncReady
    leafs["init-sync-ready-start-time"] = detailSet.InitSyncReadyStartTime
    leafs["init-sync-ready-end-time"] = detailSet.InitSyncReadyEndTime
    leafs["nsr-reset-time"] = detailSet.NsrResetTime
    leafs["is-audit-in-progress"] = detailSet.IsAuditInProgress
    leafs["audit-seq-number"] = detailSet.AuditSeqNumber
    leafs["audit-start-time"] = detailSet.AuditStartTime
    leafs["audit-end-time"] = detailSet.AuditEndTime
    return leafs
}

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetBundleName() string { return "cisco_ios_xr" }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetYangName() string { return "detail-set" }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) SetParent(parent types.Entity) { detailSet.parent = parent }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetParent() types.Entity { return detailSet.parent }

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetParentYangName() string { return "detail-sets" }

// TcpNsr_Nodes_Node_SessionSet_BriefSets
// Information about TCP NSR Session Sets
type TcpNsr_Nodes_Node_SessionSet_BriefSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Brief information about NSR Session Sets. The type is slice of
    // TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet.
    BriefSet []TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetFilter() yfilter.YFilter { return briefSets.YFilter }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) SetFilter(yf yfilter.YFilter) { briefSets.YFilter = yf }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetGoName(yname string) string {
    if yname == "brief-set" { return "BriefSet" }
    return ""
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetSegmentPath() string {
    return "brief-sets"
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "brief-set" {
        for _, c := range briefSets.BriefSet {
            if briefSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet{}
        briefSets.BriefSet = append(briefSets.BriefSet, child)
        return &briefSets.BriefSet[len(briefSets.BriefSet)-1]
    }
    return nil
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range briefSets.BriefSet {
        children[briefSets.BriefSet[i].GetSegmentPath()] = &briefSets.BriefSet[i]
    }
    return children
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetBundleName() string { return "cisco_ios_xr" }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetYangName() string { return "brief-sets" }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) SetParent(parent types.Entity) { briefSets.parent = parent }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetParent() types.Entity { return briefSets.parent }

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetParentYangName() string { return "session-set" }

// TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet
// Brief information about NSR Session Sets
type TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR Session Set. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address of the Session Set Control Block. The type is interface{} with
    // range: 0..18446744073709551615.
    Sscb interface{}

    // PID of the Client that owns this Session-set. The type is interface{} with
    // range: 0..4294967295.
    Pid interface{}

    // the name of Clinet that owns this Session-set. The type is string.
    ClientName interface{}

    // Instance of the Client that owns this Session-set. The type is interface{}
    // with range: 0..4294967295.
    ClientInstance interface{}

    // ID of this Session-set. The type is interface{} with range: 0..4294967295.
    SetId interface{}

    // TCP role for this set?. The type is interface{} with range: 0..4294967295.
    SsoRole interface{}

    // Session-set mode. The type is interface{} with range: 0..4294967295.
    Mode interface{}

    // Address Family of the sessions in this set. The type is AddrFamily.
    AddressFamily interface{}

    // Well Known Port of the client. The type is interface{} with range:
    // 0..65535.
    WellKnownPort interface{}

    // Local node of this set. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    LocalNode interface{}

    // Instance of the client application on the local node. The type is
    // interface{} with range: 0..4294967295.
    LocalInstance interface{}

    // The node protecting this set. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    ProtectNode interface{}

    // Instance of the client application on the protection node. The type is
    // interface{} with range: 0..4294967295.
    ProtectInstance interface{}

    // Number of Sessions in the set. The type is interface{} with range:
    // 0..4294967295.
    NumberOfSessions interface{}

    // How many sessions are synced with upstream partner. The type is interface{}
    // with range: 0..4294967295.
    NumberOfSyncedSessionsUpStream interface{}

    // How many sessions are synced with downstream partner. The type is
    // interface{} with range: 0..4294967295.
    NumberOfSyncedSessionsDownStream interface{}

    // Is an initial sync in progress currently?. The type is bool.
    IsInitSyncInProgress interface{}

    // Is the SSCB ready for another initial sync?. The type is bool.
    IsSscbInitSyncReady interface{}
}

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetFilter() yfilter.YFilter { return briefSet.YFilter }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) SetFilter(yf yfilter.YFilter) { briefSet.YFilter = yf }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "sscb" { return "Sscb" }
    if yname == "pid" { return "Pid" }
    if yname == "client-name" { return "ClientName" }
    if yname == "client-instance" { return "ClientInstance" }
    if yname == "set-id" { return "SetId" }
    if yname == "sso-role" { return "SsoRole" }
    if yname == "mode" { return "Mode" }
    if yname == "address-family" { return "AddressFamily" }
    if yname == "well-known-port" { return "WellKnownPort" }
    if yname == "local-node" { return "LocalNode" }
    if yname == "local-instance" { return "LocalInstance" }
    if yname == "protect-node" { return "ProtectNode" }
    if yname == "protect-instance" { return "ProtectInstance" }
    if yname == "number-of-sessions" { return "NumberOfSessions" }
    if yname == "number-of-synced-sessions-up-stream" { return "NumberOfSyncedSessionsUpStream" }
    if yname == "number-of-synced-sessions-down-stream" { return "NumberOfSyncedSessionsDownStream" }
    if yname == "is-init-sync-in-progress" { return "IsInitSyncInProgress" }
    if yname == "is-sscb-init-sync-ready" { return "IsSscbInitSyncReady" }
    return ""
}

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetSegmentPath() string {
    return "brief-set" + "[id='" + fmt.Sprintf("%v", briefSet.Id) + "']"
}

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = briefSet.Id
    leafs["sscb"] = briefSet.Sscb
    leafs["pid"] = briefSet.Pid
    leafs["client-name"] = briefSet.ClientName
    leafs["client-instance"] = briefSet.ClientInstance
    leafs["set-id"] = briefSet.SetId
    leafs["sso-role"] = briefSet.SsoRole
    leafs["mode"] = briefSet.Mode
    leafs["address-family"] = briefSet.AddressFamily
    leafs["well-known-port"] = briefSet.WellKnownPort
    leafs["local-node"] = briefSet.LocalNode
    leafs["local-instance"] = briefSet.LocalInstance
    leafs["protect-node"] = briefSet.ProtectNode
    leafs["protect-instance"] = briefSet.ProtectInstance
    leafs["number-of-sessions"] = briefSet.NumberOfSessions
    leafs["number-of-synced-sessions-up-stream"] = briefSet.NumberOfSyncedSessionsUpStream
    leafs["number-of-synced-sessions-down-stream"] = briefSet.NumberOfSyncedSessionsDownStream
    leafs["is-init-sync-in-progress"] = briefSet.IsInitSyncInProgress
    leafs["is-sscb-init-sync-ready"] = briefSet.IsSscbInitSyncReady
    return leafs
}

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetBundleName() string { return "cisco_ios_xr" }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetYangName() string { return "brief-set" }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) SetParent(parent types.Entity) { briefSet.parent = parent }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetParent() types.Entity { return briefSet.parent }

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetParentYangName() string { return "brief-sets" }

// TcpNsr_Nodes_Node_Statistics
// Statis Information about TCP NSR connections
type TcpNsr_Nodes_Node_Statistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Summary statistics across all NSR connections.
    Summary TcpNsr_Nodes_Node_Statistics_Summary

    // Table listing NSR connections for which statistic information is provided.
    StatisticClients TcpNsr_Nodes_Node_Statistics_StatisticClients

    // Table listing NSR connections for which statistic information is provided.
    StatisticSets TcpNsr_Nodes_Node_Statistics_StatisticSets

    // Table listing NSR connections for which statistic information is provided.
    StatisticSessions TcpNsr_Nodes_Node_Statistics_StatisticSessions
}

func (statistics *TcpNsr_Nodes_Node_Statistics) GetFilter() yfilter.YFilter { return statistics.YFilter }

func (statistics *TcpNsr_Nodes_Node_Statistics) SetFilter(yf yfilter.YFilter) { statistics.YFilter = yf }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetGoName(yname string) string {
    if yname == "summary" { return "Summary" }
    if yname == "statistic-clients" { return "StatisticClients" }
    if yname == "statistic-sets" { return "StatisticSets" }
    if yname == "statistic-sessions" { return "StatisticSessions" }
    return ""
}

func (statistics *TcpNsr_Nodes_Node_Statistics) GetSegmentPath() string {
    return "statistics"
}

func (statistics *TcpNsr_Nodes_Node_Statistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "summary" {
        return &statistics.Summary
    }
    if childYangName == "statistic-clients" {
        return &statistics.StatisticClients
    }
    if childYangName == "statistic-sets" {
        return &statistics.StatisticSets
    }
    if childYangName == "statistic-sessions" {
        return &statistics.StatisticSessions
    }
    return nil
}

func (statistics *TcpNsr_Nodes_Node_Statistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["summary"] = &statistics.Summary
    children["statistic-clients"] = &statistics.StatisticClients
    children["statistic-sets"] = &statistics.StatisticSets
    children["statistic-sessions"] = &statistics.StatisticSessions
    return children
}

func (statistics *TcpNsr_Nodes_Node_Statistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statistics *TcpNsr_Nodes_Node_Statistics) GetBundleName() string { return "cisco_ios_xr" }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetYangName() string { return "statistics" }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statistics *TcpNsr_Nodes_Node_Statistics) SetParent(parent types.Entity) { statistics.parent = parent }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetParent() types.Entity { return statistics.parent }

func (statistics *TcpNsr_Nodes_Node_Statistics) GetParentYangName() string { return "node" }

// TcpNsr_Nodes_Node_Statistics_Summary
// Summary statistics across all NSR connections
type TcpNsr_Nodes_Node_Statistics_Summary struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time of last clear (in seconds since 1st Jan 1970 00:00:00). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastClearedTime interface{}

    // Number of disconnected clients. The type is interface{} with range:
    // 0..4294967295.
    NumberOfConnectedClients interface{}

    // Number of disconnected clients. The type is interface{} with range:
    // 0..4294967295.
    NumberOfDisconnectedClients interface{}

    // Number of current  clients. The type is interface{} with range:
    // 0..4294967295.
    NumberOfCurrentClients interface{}

    // Number of created session sets. The type is interface{} with range:
    // 0..4294967295.
    NumberOfCreatedSessionSets interface{}

    // Number of destroyed session sets. The type is interface{} with range:
    // 0..4294967295.
    NumberOfDestroyedSessionSets interface{}

    // Number of current session sets. The type is interface{} with range:
    // 0..4294967295.
    NumberOfCurrentSessionSets interface{}

    // Number of added sessions. The type is interface{} with range:
    // 0..4294967295.
    NumberOfAddedSessions interface{}

    // Number of deleted sessions. The type is interface{} with range:
    // 0..4294967295.
    NumberOfDeletedSessions interface{}

    // Number of current sessions. The type is interface{} with range:
    // 0..4294967295.
    NumberOfCurrentSessions interface{}

    // Number of Parner Nodes. The type is interface{} with range: 0..4294967295.
    NumberOfPartnerNode interface{}

    // no. of initial-sync attempts. The type is interface{} with range:
    // 0..4294967295.
    NumberOfAttemptedInitSync interface{}

    // no. of initial-sync successes. The type is interface{} with range:
    // 0..4294967295.
    NumberOfSucceededInitSync interface{}

    // no. of initial-sync fails. The type is interface{} with range:
    // 0..4294967295.
    NumberOfFailedInitSync interface{}

    // Number of Packets held by Active TCP. The type is interface{} with range:
    // 0..4294967295.
    NumberOfHeldPackets interface{}

    // Number of held packets dropped by Active TCP. The type is interface{} with
    // range: 0..4294967295.
    NumberOfHeldButDroppedPackets interface{}

    // Number of Internal Acks held by Active TCP. The type is interface{} with
    // range: 0..4294967295.
    NumberOfHeldInternalAcks interface{}

    // Number of held Internal Acks dropped by Active TCP. The type is interface{}
    // with range: 0..4294967295.
    NumberOfHeldButDroppedInternalAcks interface{}

    // Number of Internal Acks sent to Active TCP by Standby TCP. The type is
    // interface{} with range: 0..4294967295.
    NumberOfSentInternalAcks interface{}

    // Number of Internal Acks received by Active TCP. The type is interface{}
    // with range: 0..4294967295.
    NumberOfReceivedInternalAcks interface{}

    // Number of dropped messages from partner TCP stack(s). The type is
    // interface{} with range: 0..4294967295.
    NumberOfQadReceiveMessagesDrops interface{}

    // Number of unknown messages from partner TCP stack(s). The type is
    // interface{} with range: 0..4294967295.
    NumberOfQadReceiveMessagesUnknowns interface{}

    // Number of messages accepted from partner TCP stack(s). The type is
    // interface{} with range: 0..4294967295.
    NumberOfQadReceiveMessagesAccepts interface{}

    // Number of dropped messages from partner TCP stack(s) because they were
    // out-of-order. The type is interface{} with range: 0..4294967295.
    NumberOfQadStaleReceiveMessagesDrops interface{}

    // Number of messages sent to partner TCP stack(s). The type is interface{}
    // with range: 0..4294967295.
    NumberOfQadTransferMessageSent interface{}

    // Number of messages failed to be sent to partner TCP stack(s). The type is
    // interface{} with range: 0..4294967295.
    NumberOfQadTransferMessageDrops interface{}

    // Number of iACKs dropped because there is no PCB. The type is interface{}
    // with range: 0..4294967295.
    NumberOfInternalAckDropsNoPcb interface{}

    // Number of iACKs dropped because there is no datapath SCB. The type is
    // interface{} with range: 0..4294967295.
    NumberOfInternalAckDropsNoScbdp interface{}

    // Number of iACKs dropped because session is not replicated. The type is
    // interface{} with range: 0..4294967295.
    InternalAckDropsNotReplicated interface{}

    // Number of iACKs dropped because init-sync is in 1st phase. The type is
    // interface{} with range: 0..4294967295.
    InternalAckDropsInitsyncFirstPhase interface{}

    // Number of stale iACKs dropped. The type is interface{} with range:
    // 0..4294967295.
    InternalAckDropsStale interface{}

    // Number of iACKs not held because of an immediate match. The type is
    // interface{} with range: 0..4294967295.
    InternalAckDropsImmediateMatch interface{}

    // Number of held packets dropped because of errors. The type is interface{}
    // with range: 0..4294967295.
    HeldPacketDrops interface{}

    // Aggregate Send path counters.
    SndCounters TcpNsr_Nodes_Node_Statistics_Summary_SndCounters

    // Aggregate Audit counters.
    AuditCounters TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters

    // Various types of notification stats. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic.
    NotificationStatistic []TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetFilter() yfilter.YFilter { return summary.YFilter }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) SetFilter(yf yfilter.YFilter) { summary.YFilter = yf }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetGoName(yname string) string {
    if yname == "last-cleared-time" { return "LastClearedTime" }
    if yname == "number-of-connected-clients" { return "NumberOfConnectedClients" }
    if yname == "number-of-disconnected-clients" { return "NumberOfDisconnectedClients" }
    if yname == "number-of-current-clients" { return "NumberOfCurrentClients" }
    if yname == "number-of-created-session-sets" { return "NumberOfCreatedSessionSets" }
    if yname == "number-of-destroyed-session-sets" { return "NumberOfDestroyedSessionSets" }
    if yname == "number-of-current-session-sets" { return "NumberOfCurrentSessionSets" }
    if yname == "number-of-added-sessions" { return "NumberOfAddedSessions" }
    if yname == "number-of-deleted-sessions" { return "NumberOfDeletedSessions" }
    if yname == "number-of-current-sessions" { return "NumberOfCurrentSessions" }
    if yname == "number-of-partner-node" { return "NumberOfPartnerNode" }
    if yname == "number-of-attempted-init-sync" { return "NumberOfAttemptedInitSync" }
    if yname == "number-of-succeeded-init-sync" { return "NumberOfSucceededInitSync" }
    if yname == "number-of-failed-init-sync" { return "NumberOfFailedInitSync" }
    if yname == "number-of-held-packets" { return "NumberOfHeldPackets" }
    if yname == "number-of-held-but-dropped-packets" { return "NumberOfHeldButDroppedPackets" }
    if yname == "number-of-held-internal-acks" { return "NumberOfHeldInternalAcks" }
    if yname == "number-of-held-but-dropped-internal-acks" { return "NumberOfHeldButDroppedInternalAcks" }
    if yname == "number-of-sent-internal-acks" { return "NumberOfSentInternalAcks" }
    if yname == "number-of-received-internal-acks" { return "NumberOfReceivedInternalAcks" }
    if yname == "number-of-qad-receive-messages-drops" { return "NumberOfQadReceiveMessagesDrops" }
    if yname == "number-of-qad-receive-messages-unknowns" { return "NumberOfQadReceiveMessagesUnknowns" }
    if yname == "number-of-qad-receive-messages-accepts" { return "NumberOfQadReceiveMessagesAccepts" }
    if yname == "number-of-qad-stale-receive-messages-drops" { return "NumberOfQadStaleReceiveMessagesDrops" }
    if yname == "number-of-qad-transfer-message-sent" { return "NumberOfQadTransferMessageSent" }
    if yname == "number-of-qad-transfer-message-drops" { return "NumberOfQadTransferMessageDrops" }
    if yname == "number-of-internal-ack-drops-no-pcb" { return "NumberOfInternalAckDropsNoPcb" }
    if yname == "number-of-internal-ack-drops-no-scbdp" { return "NumberOfInternalAckDropsNoScbdp" }
    if yname == "internal-ack-drops-not-replicated" { return "InternalAckDropsNotReplicated" }
    if yname == "internal-ack-drops-initsync-first-phase" { return "InternalAckDropsInitsyncFirstPhase" }
    if yname == "internal-ack-drops-stale" { return "InternalAckDropsStale" }
    if yname == "internal-ack-drops-immediate-match" { return "InternalAckDropsImmediateMatch" }
    if yname == "held-packet-drops" { return "HeldPacketDrops" }
    if yname == "snd-counters" { return "SndCounters" }
    if yname == "audit-counters" { return "AuditCounters" }
    if yname == "notification-statistic" { return "NotificationStatistic" }
    return ""
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetSegmentPath() string {
    return "summary"
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snd-counters" {
        return &summary.SndCounters
    }
    if childYangName == "audit-counters" {
        return &summary.AuditCounters
    }
    if childYangName == "notification-statistic" {
        for _, c := range summary.NotificationStatistic {
            if summary.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic{}
        summary.NotificationStatistic = append(summary.NotificationStatistic, child)
        return &summary.NotificationStatistic[len(summary.NotificationStatistic)-1]
    }
    return nil
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snd-counters"] = &summary.SndCounters
    children["audit-counters"] = &summary.AuditCounters
    for i := range summary.NotificationStatistic {
        children[summary.NotificationStatistic[i].GetSegmentPath()] = &summary.NotificationStatistic[i]
    }
    return children
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["last-cleared-time"] = summary.LastClearedTime
    leafs["number-of-connected-clients"] = summary.NumberOfConnectedClients
    leafs["number-of-disconnected-clients"] = summary.NumberOfDisconnectedClients
    leafs["number-of-current-clients"] = summary.NumberOfCurrentClients
    leafs["number-of-created-session-sets"] = summary.NumberOfCreatedSessionSets
    leafs["number-of-destroyed-session-sets"] = summary.NumberOfDestroyedSessionSets
    leafs["number-of-current-session-sets"] = summary.NumberOfCurrentSessionSets
    leafs["number-of-added-sessions"] = summary.NumberOfAddedSessions
    leafs["number-of-deleted-sessions"] = summary.NumberOfDeletedSessions
    leafs["number-of-current-sessions"] = summary.NumberOfCurrentSessions
    leafs["number-of-partner-node"] = summary.NumberOfPartnerNode
    leafs["number-of-attempted-init-sync"] = summary.NumberOfAttemptedInitSync
    leafs["number-of-succeeded-init-sync"] = summary.NumberOfSucceededInitSync
    leafs["number-of-failed-init-sync"] = summary.NumberOfFailedInitSync
    leafs["number-of-held-packets"] = summary.NumberOfHeldPackets
    leafs["number-of-held-but-dropped-packets"] = summary.NumberOfHeldButDroppedPackets
    leafs["number-of-held-internal-acks"] = summary.NumberOfHeldInternalAcks
    leafs["number-of-held-but-dropped-internal-acks"] = summary.NumberOfHeldButDroppedInternalAcks
    leafs["number-of-sent-internal-acks"] = summary.NumberOfSentInternalAcks
    leafs["number-of-received-internal-acks"] = summary.NumberOfReceivedInternalAcks
    leafs["number-of-qad-receive-messages-drops"] = summary.NumberOfQadReceiveMessagesDrops
    leafs["number-of-qad-receive-messages-unknowns"] = summary.NumberOfQadReceiveMessagesUnknowns
    leafs["number-of-qad-receive-messages-accepts"] = summary.NumberOfQadReceiveMessagesAccepts
    leafs["number-of-qad-stale-receive-messages-drops"] = summary.NumberOfQadStaleReceiveMessagesDrops
    leafs["number-of-qad-transfer-message-sent"] = summary.NumberOfQadTransferMessageSent
    leafs["number-of-qad-transfer-message-drops"] = summary.NumberOfQadTransferMessageDrops
    leafs["number-of-internal-ack-drops-no-pcb"] = summary.NumberOfInternalAckDropsNoPcb
    leafs["number-of-internal-ack-drops-no-scbdp"] = summary.NumberOfInternalAckDropsNoScbdp
    leafs["internal-ack-drops-not-replicated"] = summary.InternalAckDropsNotReplicated
    leafs["internal-ack-drops-initsync-first-phase"] = summary.InternalAckDropsInitsyncFirstPhase
    leafs["internal-ack-drops-stale"] = summary.InternalAckDropsStale
    leafs["internal-ack-drops-immediate-match"] = summary.InternalAckDropsImmediateMatch
    leafs["held-packet-drops"] = summary.HeldPacketDrops
    return leafs
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetBundleName() string { return "cisco_ios_xr" }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetYangName() string { return "summary" }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) SetParent(parent types.Entity) { summary.parent = parent }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetParent() types.Entity { return summary.parent }

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetParentYangName() string { return "statistics" }

// TcpNsr_Nodes_Node_Statistics_Summary_SndCounters
// Aggregate Send path counters
type TcpNsr_Nodes_Node_Statistics_Summary_SndCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Common send path counters.
    Common TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common

    // Aggregate only send path counters.
    AggrOnly TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetFilter() yfilter.YFilter { return sndCounters.YFilter }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) SetFilter(yf yfilter.YFilter) { sndCounters.YFilter = yf }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetGoName(yname string) string {
    if yname == "common" { return "Common" }
    if yname == "aggr-only" { return "AggrOnly" }
    return ""
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetSegmentPath() string {
    return "snd-counters"
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "common" {
        return &sndCounters.Common
    }
    if childYangName == "aggr-only" {
        return &sndCounters.AggrOnly
    }
    return nil
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["common"] = &sndCounters.Common
    children["aggr-only"] = &sndCounters.AggrOnly
    return children
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetBundleName() string { return "cisco_ios_xr" }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetYangName() string { return "snd-counters" }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) SetParent(parent types.Entity) { sndCounters.parent = parent }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetParent() types.Entity { return sndCounters.parent }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetParentYangName() string { return "summary" }

// TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common
// Common send path counters
type TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of successful DATA transfers. The type is interface{} with range:
    // 0..4294967295.
    DataXferSend interface{}

    // Amount of data transferred. The type is interface{} with range:
    // 0..18446744073709551615.
    DataXferSendTotal interface{}

    // Number of failed DATA transfers. The type is interface{} with range:
    // 0..4294967295.
    DataXferSendDrop interface{}

    // Number of data transfer msgs., that required new IOV's to be allocated. The
    // type is interface{} with range: 0..4294967295.
    DataXferSendIovAlloc interface{}

    // Number of received DATA transfers. The type is interface{} with range:
    // 0..4294967295.
    DataXferRcv interface{}

    // Number of successfully received DATA transfers. The type is interface{}
    // with range: 0..4294967295.
    DataXferRcvSuccess interface{}

    // Number of received DATA transfers that had buffer trim failures. The type
    // is interface{} with range: 0..4294967295.
    DataXferRcvFailBufferTrim interface{}

    // Number of received DATA transfers that had failures because the send path
    // was out of sync. The type is interface{} with range: 0..4294967295.
    DataXferRcvFailSndUnaOutOfSync interface{}

    // Number of successful Segmentation instruction messages. The type is
    // interface{} with range: 0..4294967295.
    SegInstrSend interface{}

    // Number of segement units transferred via the successful Segmentation
    // instruction messages. The type is interface{} with range: 0..4294967295.
    SegInstrSendUnits interface{}

    // Number of failed Segmentation instruction messages. The type is interface{}
    // with range: 0..4294967295.
    SegInstrSendDrop interface{}

    // Number of received Segmentation instruction messages. The type is
    // interface{} with range: 0..4294967295.
    SegInstrRcv interface{}

    // Number of successfully received Segmentation instruction messages. The type
    // is interface{} with range: 0..4294967295.
    SegInstrRcvSuccess interface{}

    // Number of received Segmentation instructions that had buffer trim failures.
    // The type is interface{} with range: 0..4294967295.
    SegInstrRcvFailBufferTrim interface{}

    // Number of received Segmentation instructions that had failures during TCP
    // processing. The type is interface{} with range: 0..4294967295.
    SegInstrRcvFailTcpProcess interface{}

    // Number of successful NACK messages. The type is interface{} with range:
    // 0..4294967295.
    NackSend interface{}

    // Number of failed NACK messages. The type is interface{} with range:
    // 0..4294967295.
    NackSendDrop interface{}

    // Number of received NACK messages. The type is interface{} with range:
    // 0..4294967295.
    NackRcv interface{}

    // Number of successfully received NACK messages. The type is interface{} with
    // range: 0..4294967295.
    NackRcvSuccess interface{}

    // Number of received NACK messages that had failures when sending data in
    // response to the NACK. The type is interface{} with range: 0..4294967295.
    NackRcvFailDataSend interface{}

    // Number of successful Cleanup messages. The type is interface{} with range:
    // 0..4294967295.
    CleanupSend interface{}

    // Number of failed Cleanup messages. The type is interface{} with range:
    // 0..4294967295.
    CleanupSendDrop interface{}

    // Number of received Cleanup messages. The type is interface{} with range:
    // 0..4294967295.
    CleanupRcv interface{}

    // Number of successfully received Cleanup messages. The type is interface{}
    // with range: 0..4294967295.
    CleanupRcvSuccess interface{}

    // Number of Cleanup messages that had trim failures. The type is interface{}
    // with range: 0..4294967295.
    CleanupRcvFailBufferTrim interface{}
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetFilter() yfilter.YFilter { return common.YFilter }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) SetFilter(yf yfilter.YFilter) { common.YFilter = yf }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetGoName(yname string) string {
    if yname == "data-xfer-send" { return "DataXferSend" }
    if yname == "data-xfer-send-total" { return "DataXferSendTotal" }
    if yname == "data-xfer-send-drop" { return "DataXferSendDrop" }
    if yname == "data-xfer-send-iov-alloc" { return "DataXferSendIovAlloc" }
    if yname == "data-xfer-rcv" { return "DataXferRcv" }
    if yname == "data-xfer-rcv-success" { return "DataXferRcvSuccess" }
    if yname == "data-xfer-rcv-fail-buffer-trim" { return "DataXferRcvFailBufferTrim" }
    if yname == "data-xfer-rcv-fail-snd-una-out-of-sync" { return "DataXferRcvFailSndUnaOutOfSync" }
    if yname == "seg-instr-send" { return "SegInstrSend" }
    if yname == "seg-instr-send-units" { return "SegInstrSendUnits" }
    if yname == "seg-instr-send-drop" { return "SegInstrSendDrop" }
    if yname == "seg-instr-rcv" { return "SegInstrRcv" }
    if yname == "seg-instr-rcv-success" { return "SegInstrRcvSuccess" }
    if yname == "seg-instr-rcv-fail-buffer-trim" { return "SegInstrRcvFailBufferTrim" }
    if yname == "seg-instr-rcv-fail-tcp-process" { return "SegInstrRcvFailTcpProcess" }
    if yname == "nack-send" { return "NackSend" }
    if yname == "nack-send-drop" { return "NackSendDrop" }
    if yname == "nack-rcv" { return "NackRcv" }
    if yname == "nack-rcv-success" { return "NackRcvSuccess" }
    if yname == "nack-rcv-fail-data-send" { return "NackRcvFailDataSend" }
    if yname == "cleanup-send" { return "CleanupSend" }
    if yname == "cleanup-send-drop" { return "CleanupSendDrop" }
    if yname == "cleanup-rcv" { return "CleanupRcv" }
    if yname == "cleanup-rcv-success" { return "CleanupRcvSuccess" }
    if yname == "cleanup-rcv-fail-buffer-trim" { return "CleanupRcvFailBufferTrim" }
    return ""
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetSegmentPath() string {
    return "common"
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-xfer-send"] = common.DataXferSend
    leafs["data-xfer-send-total"] = common.DataXferSendTotal
    leafs["data-xfer-send-drop"] = common.DataXferSendDrop
    leafs["data-xfer-send-iov-alloc"] = common.DataXferSendIovAlloc
    leafs["data-xfer-rcv"] = common.DataXferRcv
    leafs["data-xfer-rcv-success"] = common.DataXferRcvSuccess
    leafs["data-xfer-rcv-fail-buffer-trim"] = common.DataXferRcvFailBufferTrim
    leafs["data-xfer-rcv-fail-snd-una-out-of-sync"] = common.DataXferRcvFailSndUnaOutOfSync
    leafs["seg-instr-send"] = common.SegInstrSend
    leafs["seg-instr-send-units"] = common.SegInstrSendUnits
    leafs["seg-instr-send-drop"] = common.SegInstrSendDrop
    leafs["seg-instr-rcv"] = common.SegInstrRcv
    leafs["seg-instr-rcv-success"] = common.SegInstrRcvSuccess
    leafs["seg-instr-rcv-fail-buffer-trim"] = common.SegInstrRcvFailBufferTrim
    leafs["seg-instr-rcv-fail-tcp-process"] = common.SegInstrRcvFailTcpProcess
    leafs["nack-send"] = common.NackSend
    leafs["nack-send-drop"] = common.NackSendDrop
    leafs["nack-rcv"] = common.NackRcv
    leafs["nack-rcv-success"] = common.NackRcvSuccess
    leafs["nack-rcv-fail-data-send"] = common.NackRcvFailDataSend
    leafs["cleanup-send"] = common.CleanupSend
    leafs["cleanup-send-drop"] = common.CleanupSendDrop
    leafs["cleanup-rcv"] = common.CleanupRcv
    leafs["cleanup-rcv-success"] = common.CleanupRcvSuccess
    leafs["cleanup-rcv-fail-buffer-trim"] = common.CleanupRcvFailBufferTrim
    return leafs
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetBundleName() string { return "cisco_ios_xr" }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetYangName() string { return "common" }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) SetParent(parent types.Entity) { common.parent = parent }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetParent() types.Entity { return common.parent }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetParentYangName() string { return "snd-counters" }

// TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly
// Aggregate only send path counters
type TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of Data transfer messages dropped because PCB wasn't found. The type
    // is interface{} with range: 0..4294967295.
    DataXferRcvDropNoPcb interface{}

    // Number of Data transfer messages dropped because SCB DP wasn't found. The
    // type is interface{} with range: 0..4294967295.
    DataXferRcvDropNoScbDp interface{}

    // Number of Segmentation instruction messages dropped because PCB wasn't
    // found. The type is interface{} with range: 0..4294967295.
    SegInstrRcvDropNoPcb interface{}

    // Number of Segmentation instruction messages dropped because SCB DP wasn't
    // found. The type is interface{} with range: 0..4294967295.
    SegInstrRcvDropNoScbDp interface{}

    // Number of NACK messages dropped because PCB wasn't found. The type is
    // interface{} with range: 0..4294967295.
    NackRcvDropNoPcb interface{}

    // Number of NACK messages dropped because SCB DP wasn't found. The type is
    // interface{} with range: 0..4294967295.
    NackRcvDropNoScbDp interface{}

    // Number of Cleanup messages dropped because PCB wasn't found. The type is
    // interface{} with range: 0..4294967295.
    CleanupRcvDropNoPcb interface{}

    // Number of Cleanup messages dropped because SCB DP wasn't found. The type is
    // interface{} with range: 0..4294967295.
    CleanupRcvDropNoScbDp interface{}
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetFilter() yfilter.YFilter { return aggrOnly.YFilter }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) SetFilter(yf yfilter.YFilter) { aggrOnly.YFilter = yf }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetGoName(yname string) string {
    if yname == "data-xfer-rcv-drop-no-pcb" { return "DataXferRcvDropNoPcb" }
    if yname == "data-xfer-rcv-drop-no-scb-dp" { return "DataXferRcvDropNoScbDp" }
    if yname == "seg-instr-rcv-drop-no-pcb" { return "SegInstrRcvDropNoPcb" }
    if yname == "seg-instr-rcv-drop-no-scb-dp" { return "SegInstrRcvDropNoScbDp" }
    if yname == "nack-rcv-drop-no-pcb" { return "NackRcvDropNoPcb" }
    if yname == "nack-rcv-drop-no-scb-dp" { return "NackRcvDropNoScbDp" }
    if yname == "cleanup-rcv-drop-no-pcb" { return "CleanupRcvDropNoPcb" }
    if yname == "cleanup-rcv-drop-no-scb-dp" { return "CleanupRcvDropNoScbDp" }
    return ""
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetSegmentPath() string {
    return "aggr-only"
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-xfer-rcv-drop-no-pcb"] = aggrOnly.DataXferRcvDropNoPcb
    leafs["data-xfer-rcv-drop-no-scb-dp"] = aggrOnly.DataXferRcvDropNoScbDp
    leafs["seg-instr-rcv-drop-no-pcb"] = aggrOnly.SegInstrRcvDropNoPcb
    leafs["seg-instr-rcv-drop-no-scb-dp"] = aggrOnly.SegInstrRcvDropNoScbDp
    leafs["nack-rcv-drop-no-pcb"] = aggrOnly.NackRcvDropNoPcb
    leafs["nack-rcv-drop-no-scb-dp"] = aggrOnly.NackRcvDropNoScbDp
    leafs["cleanup-rcv-drop-no-pcb"] = aggrOnly.CleanupRcvDropNoPcb
    leafs["cleanup-rcv-drop-no-scb-dp"] = aggrOnly.CleanupRcvDropNoScbDp
    return leafs
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetBundleName() string { return "cisco_ios_xr" }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetYangName() string { return "aggr-only" }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) SetParent(parent types.Entity) { aggrOnly.parent = parent }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetParent() types.Entity { return aggrOnly.parent }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetParentYangName() string { return "snd-counters" }

// TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters
// Aggregate Audit counters
type TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Common audit counters.
    Common TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common

    // Aggregate only audit counters.
    AggrOnly TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetFilter() yfilter.YFilter { return auditCounters.YFilter }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) SetFilter(yf yfilter.YFilter) { auditCounters.YFilter = yf }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetGoName(yname string) string {
    if yname == "common" { return "Common" }
    if yname == "aggr-only" { return "AggrOnly" }
    return ""
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetSegmentPath() string {
    return "audit-counters"
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "common" {
        return &auditCounters.Common
    }
    if childYangName == "aggr-only" {
        return &auditCounters.AggrOnly
    }
    return nil
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["common"] = &auditCounters.Common
    children["aggr-only"] = &auditCounters.AggrOnly
    return children
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetBundleName() string { return "cisco_ios_xr" }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetYangName() string { return "audit-counters" }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) SetParent(parent types.Entity) { auditCounters.parent = parent }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetParent() types.Entity { return auditCounters.parent }

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetParentYangName() string { return "summary" }

// TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common
// Common audit counters
type TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of successful session-set Mark's sent by active. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetSend interface{}

    // Number of failed session-set Mark's. The type is interface{} with range:
    // 0..4294967295.
    MarkSessionSetSendDrop interface{}

    // Number of successful session-set Mark's received by standby. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetRcv interface{}

    // Number of session-set Mark's dropped by standby. The type is interface{}
    // with range: 0..4294967295.
    MarkSessionSetRcvDrop interface{}

    // Number of successful session audits sent by active. The type is interface{}
    // with range: 0..4294967295.
    SessionSend interface{}

    // Number of session audits that couldn't be sent by active. The type is
    // interface{} with range: 0..4294967295.
    SessionSendDrop interface{}

    // Number of session audits received by standby. The type is interface{} with
    // range: 0..4294967295.
    SessionRcv interface{}

    // Number of session audits dropped by standby. The type is interface{} with
    // range: 0..4294967295.
    SessionRcvDrop interface{}

    // Number of successful session-set Sweep's sent by active. The type is
    // interface{} with range: 0..4294967295.
    SweepSessionSetSend interface{}

    // Number of failed session-set Sweep's. The type is interface{} with range:
    // 0..4294967295.
    SweepSessionSetSendDrop interface{}

    // Number of successful session-set Sweep's received by standby. The type is
    // interface{} with range: 0..4294967295.
    SweepSessionSetRcv interface{}

    // Number of session-set Sweep's dropped by standby. The type is interface{}
    // with range: 0..4294967295.
    SweepSessionSetRcvDrop interface{}

    // Number of successful audit responses sent by standby. The type is
    // interface{} with range: 0..4294967295.
    SessionSetResponseSend interface{}

    // Number of audit responses that couldn't be sent by standby. The type is
    // interface{} with range: 0..4294967295.
    SessionSetResponseSendDrop interface{}

    // Number of audit responses received by active. The type is interface{} with
    // range: 0..4294967295.
    SessionSetResponseRcv interface{}

    // Number of audit responses dropped by active. The type is interface{} with
    // range: 0..4294967295.
    SessionSetResponseRcvDrop interface{}

    // Number of successful audit mark acks sent by standby. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetAckSend interface{}

    // Number of audit mark acks that couldn't be sent by standby. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetAckSendDrop interface{}

    // Number of audit mark acks received by active. The type is interface{} with
    // range: 0..4294967295.
    MarkSessionSetAckRcv interface{}

    // Number of audit mark acks dropped by active. The type is interface{} with
    // range: 0..4294967295.
    MarkSessionSetAckRcvDrop interface{}

    // Number of successful audit mark nacks sent by standby. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetNackSend interface{}

    // Number of audit mark nacks that couldn't be sent by standby. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetNackSendDrop interface{}

    // Number of audit mark nacks received by active. The type is interface{} with
    // range: 0..4294967295.
    MarkSessionSetNackRcv interface{}

    // Number of audit mark nacks dropped by active. The type is interface{} with
    // range: 0..4294967295.
    MarkSessionSetNackRcvDrop interface{}

    // Number of times the active aborted an audit session. The type is
    // interface{} with range: 0..4294967295.
    Abort interface{}
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetFilter() yfilter.YFilter { return common.YFilter }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) SetFilter(yf yfilter.YFilter) { common.YFilter = yf }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetGoName(yname string) string {
    if yname == "mark-session-set-send" { return "MarkSessionSetSend" }
    if yname == "mark-session-set-send-drop" { return "MarkSessionSetSendDrop" }
    if yname == "mark-session-set-rcv" { return "MarkSessionSetRcv" }
    if yname == "mark-session-set-rcv-drop" { return "MarkSessionSetRcvDrop" }
    if yname == "session-send" { return "SessionSend" }
    if yname == "session-send-drop" { return "SessionSendDrop" }
    if yname == "session-rcv" { return "SessionRcv" }
    if yname == "session-rcv-drop" { return "SessionRcvDrop" }
    if yname == "sweep-session-set-send" { return "SweepSessionSetSend" }
    if yname == "sweep-session-set-send-drop" { return "SweepSessionSetSendDrop" }
    if yname == "sweep-session-set-rcv" { return "SweepSessionSetRcv" }
    if yname == "sweep-session-set-rcv-drop" { return "SweepSessionSetRcvDrop" }
    if yname == "session-set-response-send" { return "SessionSetResponseSend" }
    if yname == "session-set-response-send-drop" { return "SessionSetResponseSendDrop" }
    if yname == "session-set-response-rcv" { return "SessionSetResponseRcv" }
    if yname == "session-set-response-rcv-drop" { return "SessionSetResponseRcvDrop" }
    if yname == "mark-session-set-ack-send" { return "MarkSessionSetAckSend" }
    if yname == "mark-session-set-ack-send-drop" { return "MarkSessionSetAckSendDrop" }
    if yname == "mark-session-set-ack-rcv" { return "MarkSessionSetAckRcv" }
    if yname == "mark-session-set-ack-rcv-drop" { return "MarkSessionSetAckRcvDrop" }
    if yname == "mark-session-set-nack-send" { return "MarkSessionSetNackSend" }
    if yname == "mark-session-set-nack-send-drop" { return "MarkSessionSetNackSendDrop" }
    if yname == "mark-session-set-nack-rcv" { return "MarkSessionSetNackRcv" }
    if yname == "mark-session-set-nack-rcv-drop" { return "MarkSessionSetNackRcvDrop" }
    if yname == "abort" { return "Abort" }
    return ""
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetSegmentPath() string {
    return "common"
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-session-set-send"] = common.MarkSessionSetSend
    leafs["mark-session-set-send-drop"] = common.MarkSessionSetSendDrop
    leafs["mark-session-set-rcv"] = common.MarkSessionSetRcv
    leafs["mark-session-set-rcv-drop"] = common.MarkSessionSetRcvDrop
    leafs["session-send"] = common.SessionSend
    leafs["session-send-drop"] = common.SessionSendDrop
    leafs["session-rcv"] = common.SessionRcv
    leafs["session-rcv-drop"] = common.SessionRcvDrop
    leafs["sweep-session-set-send"] = common.SweepSessionSetSend
    leafs["sweep-session-set-send-drop"] = common.SweepSessionSetSendDrop
    leafs["sweep-session-set-rcv"] = common.SweepSessionSetRcv
    leafs["sweep-session-set-rcv-drop"] = common.SweepSessionSetRcvDrop
    leafs["session-set-response-send"] = common.SessionSetResponseSend
    leafs["session-set-response-send-drop"] = common.SessionSetResponseSendDrop
    leafs["session-set-response-rcv"] = common.SessionSetResponseRcv
    leafs["session-set-response-rcv-drop"] = common.SessionSetResponseRcvDrop
    leafs["mark-session-set-ack-send"] = common.MarkSessionSetAckSend
    leafs["mark-session-set-ack-send-drop"] = common.MarkSessionSetAckSendDrop
    leafs["mark-session-set-ack-rcv"] = common.MarkSessionSetAckRcv
    leafs["mark-session-set-ack-rcv-drop"] = common.MarkSessionSetAckRcvDrop
    leafs["mark-session-set-nack-send"] = common.MarkSessionSetNackSend
    leafs["mark-session-set-nack-send-drop"] = common.MarkSessionSetNackSendDrop
    leafs["mark-session-set-nack-rcv"] = common.MarkSessionSetNackRcv
    leafs["mark-session-set-nack-rcv-drop"] = common.MarkSessionSetNackRcvDrop
    leafs["abort"] = common.Abort
    return leafs
}

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetBundleName() string { return "cisco_ios_xr" }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetYangName() string { return "common" }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) SetParent(parent types.Entity) { common.parent = parent }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetParent() types.Entity { return common.parent }

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetParentYangName() string { return "audit-counters" }

// TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly
// Aggregate only audit counters
type TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of session-set Mark messages dropped by standby. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetRcvDropAggr interface{}

    // Number of session audit messages dropped by standby. The type is
    // interface{} with range: 0..4294967295.
    SessionRcvDropAggr interface{}

    // Number of session-set Sweep messages dropped by standby. The type is
    // interface{} with range: 0..4294967295.
    SweepSessionSetRcvDropAggr interface{}

    // Number of session-set response messages dropped by active. The type is
    // interface{} with range: 0..4294967295.
    SessionSetResponseRcvDropAggr interface{}

    // Number of session-set mark ack messages dropped by active. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetAckRcvDropAggr interface{}

    // Number of session-set mark nack messages dropped by active. The type is
    // interface{} with range: 0..4294967295.
    MarkSessionSetNackRcvDropAggr interface{}
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetFilter() yfilter.YFilter { return aggrOnly.YFilter }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) SetFilter(yf yfilter.YFilter) { aggrOnly.YFilter = yf }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetGoName(yname string) string {
    if yname == "mark-session-set-rcv-drop-aggr" { return "MarkSessionSetRcvDropAggr" }
    if yname == "session-rcv-drop-aggr" { return "SessionRcvDropAggr" }
    if yname == "sweep-session-set-rcv-drop-aggr" { return "SweepSessionSetRcvDropAggr" }
    if yname == "session-set-response-rcv-drop-aggr" { return "SessionSetResponseRcvDropAggr" }
    if yname == "mark-session-set-ack-rcv-drop-aggr" { return "MarkSessionSetAckRcvDropAggr" }
    if yname == "mark-session-set-nack-rcv-drop-aggr" { return "MarkSessionSetNackRcvDropAggr" }
    return ""
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetSegmentPath() string {
    return "aggr-only"
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mark-session-set-rcv-drop-aggr"] = aggrOnly.MarkSessionSetRcvDropAggr
    leafs["session-rcv-drop-aggr"] = aggrOnly.SessionRcvDropAggr
    leafs["sweep-session-set-rcv-drop-aggr"] = aggrOnly.SweepSessionSetRcvDropAggr
    leafs["session-set-response-rcv-drop-aggr"] = aggrOnly.SessionSetResponseRcvDropAggr
    leafs["mark-session-set-ack-rcv-drop-aggr"] = aggrOnly.MarkSessionSetAckRcvDropAggr
    leafs["mark-session-set-nack-rcv-drop-aggr"] = aggrOnly.MarkSessionSetNackRcvDropAggr
    return leafs
}

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetBundleName() string { return "cisco_ios_xr" }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetYangName() string { return "aggr-only" }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) SetParent(parent types.Entity) { aggrOnly.parent = parent }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetParent() types.Entity { return aggrOnly.parent }

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetParentYangName() string { return "audit-counters" }

// TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic
// Various types of notification stats
type TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // how many were queued. The type is interface{} with range: 0..4294967295.
    QueuedCount interface{}

    // Errors while queuing the notifs. The type is interface{} with range:
    // 0..4294967295.
    FailedCount interface{}

    // How many were picked up by app?. The type is interface{} with range:
    // 0..4294967295.
    DeliveredCount interface{}

    // How many were dropped because of timeout. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetFilter() yfilter.YFilter { return notificationStatistic.YFilter }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) SetFilter(yf yfilter.YFilter) { notificationStatistic.YFilter = yf }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetGoName(yname string) string {
    if yname == "queued-count" { return "QueuedCount" }
    if yname == "failed-count" { return "FailedCount" }
    if yname == "delivered-count" { return "DeliveredCount" }
    if yname == "dropped-count" { return "DroppedCount" }
    return ""
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetSegmentPath() string {
    return "notification-statistic"
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queued-count"] = notificationStatistic.QueuedCount
    leafs["failed-count"] = notificationStatistic.FailedCount
    leafs["delivered-count"] = notificationStatistic.DeliveredCount
    leafs["dropped-count"] = notificationStatistic.DroppedCount
    return leafs
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetYangName() string { return "notification-statistic" }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) SetParent(parent types.Entity) { notificationStatistic.parent = parent }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetParent() types.Entity { return notificationStatistic.parent }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetParentYangName() string { return "summary" }

// TcpNsr_Nodes_Node_Statistics_StatisticClients
// Table listing NSR connections for which
// statistic information is provided
type TcpNsr_Nodes_Node_Statistics_StatisticClients struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // showing statistic information of NSR Clients. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient.
    StatisticClient []TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetFilter() yfilter.YFilter { return statisticClients.YFilter }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) SetFilter(yf yfilter.YFilter) { statisticClients.YFilter = yf }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetGoName(yname string) string {
    if yname == "statistic-client" { return "StatisticClient" }
    return ""
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetSegmentPath() string {
    return "statistic-clients"
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic-client" {
        for _, c := range statisticClients.StatisticClient {
            if statisticClients.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient{}
        statisticClients.StatisticClient = append(statisticClients.StatisticClient, child)
        return &statisticClients.StatisticClient[len(statisticClients.StatisticClient)-1]
    }
    return nil
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticClients.StatisticClient {
        children[statisticClients.StatisticClient[i].GetSegmentPath()] = &statisticClients.StatisticClient[i]
    }
    return children
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetBundleName() string { return "cisco_ios_xr" }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetYangName() string { return "statistic-clients" }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) SetParent(parent types.Entity) { statisticClients.parent = parent }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetParent() types.Entity { return statisticClients.parent }

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetParentYangName() string { return "statistics" }

// TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient
// showing statistic information of NSR Clients
type TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR Client. The type is string with pattern:
    // [0-9a-fA-F]{1,8}.
    Id interface{}

    // Address of the Client Control Block. The type is interface{} with range:
    // 0..18446744073709551615.
    Ccb interface{}

    // PID of the Client. The type is interface{} with range: 0..4294967295.
    Pid interface{}

    // Proc name of Clinet. The type is string.
    ProcessName interface{}

    // JOb ID of Client. The type is interface{} with range:
    // -2147483648..2147483647.
    JobId interface{}

    // Instance of the Client. The type is interface{} with range: 0..4294967295.
    Instance interface{}

    // Time of connect (in seconds since 1st Jan 1970 00:00:00). The type is
    // interface{} with range: 0..4294967295. Units are second.
    ConnectedAt interface{}

    // Num of created session sets. The type is interface{} with range:
    // 0..4294967295.
    NumberOfCreatedSscb interface{}

    // Num of deleted session sets. The type is interface{} with range:
    // 0..4294967295.
    NumberOfDeletedSscb interface{}

    // Time of last clear (in seconds since 1st Jan 1970 00:00:00). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastClearedTime interface{}

    // Various types of notification stats. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic.
    NotificationStatistic []TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetFilter() yfilter.YFilter { return statisticClient.YFilter }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) SetFilter(yf yfilter.YFilter) { statisticClient.YFilter = yf }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "ccb" { return "Ccb" }
    if yname == "pid" { return "Pid" }
    if yname == "process-name" { return "ProcessName" }
    if yname == "job-id" { return "JobId" }
    if yname == "instance" { return "Instance" }
    if yname == "connected-at" { return "ConnectedAt" }
    if yname == "number-of-created-sscb" { return "NumberOfCreatedSscb" }
    if yname == "number-of-deleted-sscb" { return "NumberOfDeletedSscb" }
    if yname == "last-cleared-time" { return "LastClearedTime" }
    if yname == "notification-statistic" { return "NotificationStatistic" }
    return ""
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetSegmentPath() string {
    return "statistic-client" + "[id='" + fmt.Sprintf("%v", statisticClient.Id) + "']"
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "notification-statistic" {
        for _, c := range statisticClient.NotificationStatistic {
            if statisticClient.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic{}
        statisticClient.NotificationStatistic = append(statisticClient.NotificationStatistic, child)
        return &statisticClient.NotificationStatistic[len(statisticClient.NotificationStatistic)-1]
    }
    return nil
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticClient.NotificationStatistic {
        children[statisticClient.NotificationStatistic[i].GetSegmentPath()] = &statisticClient.NotificationStatistic[i]
    }
    return children
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = statisticClient.Id
    leafs["ccb"] = statisticClient.Ccb
    leafs["pid"] = statisticClient.Pid
    leafs["process-name"] = statisticClient.ProcessName
    leafs["job-id"] = statisticClient.JobId
    leafs["instance"] = statisticClient.Instance
    leafs["connected-at"] = statisticClient.ConnectedAt
    leafs["number-of-created-sscb"] = statisticClient.NumberOfCreatedSscb
    leafs["number-of-deleted-sscb"] = statisticClient.NumberOfDeletedSscb
    leafs["last-cleared-time"] = statisticClient.LastClearedTime
    return leafs
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetBundleName() string { return "cisco_ios_xr" }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetYangName() string { return "statistic-client" }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) SetParent(parent types.Entity) { statisticClient.parent = parent }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetParent() types.Entity { return statisticClient.parent }

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetParentYangName() string { return "statistic-clients" }

// TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic
// Various types of notification stats
type TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // how many were queued. The type is interface{} with range: 0..4294967295.
    QueuedCount interface{}

    // Errors while queuing the notifs. The type is interface{} with range:
    // 0..4294967295.
    FailedCount interface{}

    // How many were picked up by app?. The type is interface{} with range:
    // 0..4294967295.
    DeliveredCount interface{}

    // How many were dropped because of timeout. The type is interface{} with
    // range: 0..4294967295.
    DroppedCount interface{}
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetFilter() yfilter.YFilter { return notificationStatistic.YFilter }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) SetFilter(yf yfilter.YFilter) { notificationStatistic.YFilter = yf }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetGoName(yname string) string {
    if yname == "queued-count" { return "QueuedCount" }
    if yname == "failed-count" { return "FailedCount" }
    if yname == "delivered-count" { return "DeliveredCount" }
    if yname == "dropped-count" { return "DroppedCount" }
    return ""
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetSegmentPath() string {
    return "notification-statistic"
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["queued-count"] = notificationStatistic.QueuedCount
    leafs["failed-count"] = notificationStatistic.FailedCount
    leafs["delivered-count"] = notificationStatistic.DeliveredCount
    leafs["dropped-count"] = notificationStatistic.DroppedCount
    return leafs
}

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetBundleName() string { return "cisco_ios_xr" }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetYangName() string { return "notification-statistic" }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) SetParent(parent types.Entity) { notificationStatistic.parent = parent }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetParent() types.Entity { return notificationStatistic.parent }

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetParentYangName() string { return "statistic-client" }

// TcpNsr_Nodes_Node_Statistics_StatisticSets
// Table listing NSR connections for which
// statistic information is provided
type TcpNsr_Nodes_Node_Statistics_StatisticSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // showing statistic information of NSR Session Set. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet.
    StatisticSet []TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetFilter() yfilter.YFilter { return statisticSets.YFilter }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) SetFilter(yf yfilter.YFilter) { statisticSets.YFilter = yf }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetGoName(yname string) string {
    if yname == "statistic-set" { return "StatisticSet" }
    return ""
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetSegmentPath() string {
    return "statistic-sets"
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic-set" {
        for _, c := range statisticSets.StatisticSet {
            if statisticSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet{}
        statisticSets.StatisticSet = append(statisticSets.StatisticSet, child)
        return &statisticSets.StatisticSet[len(statisticSets.StatisticSet)-1]
    }
    return nil
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticSets.StatisticSet {
        children[statisticSets.StatisticSet[i].GetSegmentPath()] = &statisticSets.StatisticSet[i]
    }
    return children
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetBundleName() string { return "cisco_ios_xr" }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetYangName() string { return "statistic-sets" }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) SetParent(parent types.Entity) { statisticSets.parent = parent }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetParent() types.Entity { return statisticSets.parent }

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetParentYangName() string { return "statistics" }

// TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet
// showing statistic information of NSR Session
// Set
type TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of NSR Session Set. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Id interface{}

    // SSCB Address. The type is interface{} with range: 0..18446744073709551615.
    Sscb interface{}

    // ID of this Session-set. The type is interface{} with range: 0..4294967295.
    SetId interface{}

    // no. of initial-sync attempts. The type is interface{} with range:
    // 0..4294967295.
    NumberOfAttemptedInitSync interface{}

    // no. of initial-sync successes. The type is interface{} with range:
    // 0..4294967295.
    NumberOfSucceededInitSync interface{}

    // no. of initial-sync failures. The type is interface{} with range:
    // 0..4294967295.
    NumberOfFailedInitSync interface{}

    // Number of Switch-overs. The type is interface{} with range: 0..4294967295.
    NumberOfFailover interface{}

    // Number of times NSR was reset for the session. The type is interface{} with
    // range: 0..4294967295.
    NumberOfNsrResets interface{}

    // Time of last clear (in seconds since 1st Jan 1970 00:00:00). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastClearedTime interface{}
}

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetFilter() yfilter.YFilter { return statisticSet.YFilter }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) SetFilter(yf yfilter.YFilter) { statisticSet.YFilter = yf }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "sscb" { return "Sscb" }
    if yname == "set-id" { return "SetId" }
    if yname == "number-of-attempted-init-sync" { return "NumberOfAttemptedInitSync" }
    if yname == "number-of-succeeded-init-sync" { return "NumberOfSucceededInitSync" }
    if yname == "number-of-failed-init-sync" { return "NumberOfFailedInitSync" }
    if yname == "number-of-failover" { return "NumberOfFailover" }
    if yname == "number-of-nsr-resets" { return "NumberOfNsrResets" }
    if yname == "last-cleared-time" { return "LastClearedTime" }
    return ""
}

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetSegmentPath() string {
    return "statistic-set" + "[id='" + fmt.Sprintf("%v", statisticSet.Id) + "']"
}

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = statisticSet.Id
    leafs["sscb"] = statisticSet.Sscb
    leafs["set-id"] = statisticSet.SetId
    leafs["number-of-attempted-init-sync"] = statisticSet.NumberOfAttemptedInitSync
    leafs["number-of-succeeded-init-sync"] = statisticSet.NumberOfSucceededInitSync
    leafs["number-of-failed-init-sync"] = statisticSet.NumberOfFailedInitSync
    leafs["number-of-failover"] = statisticSet.NumberOfFailover
    leafs["number-of-nsr-resets"] = statisticSet.NumberOfNsrResets
    leafs["last-cleared-time"] = statisticSet.LastClearedTime
    return leafs
}

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetBundleName() string { return "cisco_ios_xr" }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetYangName() string { return "statistic-set" }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) SetParent(parent types.Entity) { statisticSet.parent = parent }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetParent() types.Entity { return statisticSet.parent }

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetParentYangName() string { return "statistic-sets" }

// TcpNsr_Nodes_Node_Statistics_StatisticSessions
// Table listing NSR connections for which
// statistic information is provided
type TcpNsr_Nodes_Node_Statistics_StatisticSessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // showing statistic information of TCP connections. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession.
    StatisticSession []TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetFilter() yfilter.YFilter { return statisticSessions.YFilter }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) SetFilter(yf yfilter.YFilter) { statisticSessions.YFilter = yf }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetGoName(yname string) string {
    if yname == "statistic-session" { return "StatisticSession" }
    return ""
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetSegmentPath() string {
    return "statistic-sessions"
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statistic-session" {
        for _, c := range statisticSessions.StatisticSession {
            if statisticSessions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession{}
        statisticSessions.StatisticSession = append(statisticSessions.StatisticSession, child)
        return &statisticSessions.StatisticSession[len(statisticSessions.StatisticSession)-1]
    }
    return nil
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statisticSessions.StatisticSession {
        children[statisticSessions.StatisticSession[i].GetSegmentPath()] = &statisticSessions.StatisticSession[i]
    }
    return children
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetBundleName() string { return "cisco_ios_xr" }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetYangName() string { return "statistic-sessions" }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) SetParent(parent types.Entity) { statisticSessions.parent = parent }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetParent() types.Entity { return statisticSessions.parent }

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetParentYangName() string { return "statistics" }

// TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession
// showing statistic information of TCP
// connections
type TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. ID of TCP connection. The type is string with
    // pattern: [0-9a-fA-F]{1,8}.
    Id interface{}

    // PCB Address. The type is interface{} with range: 0..18446744073709551615.
    Pcb interface{}

    // no. of times nsr went up. The type is interface{} with range:
    // 0..4294967295.
    NumberOfTimesNsrUp interface{}

    // no. of times nsr went down. The type is interface{} with range:
    // 0..4294967295.
    NumberOfTimersNsrDown interface{}

    // no. of times nsr was disabled. The type is interface{} with range:
    // 0..4294967295.
    NumberOfTimesNsrDisabled interface{}

    // no. of times fail-over occured. The type is interface{} with range:
    // 0..4294967295.
    NumberOfTimesNsrFailOver interface{}

    // Number of iACKs dropped because session is not replicated. The type is
    // interface{} with range: 0..18446744073709551615.
    InternalAckDropsNotReplicated interface{}

    // Number of iACKs dropped because 1st phase of init-sync is in progress. The
    // type is interface{} with range: 0..18446744073709551615.
    InternalAckDropsInitsyncFirstPhase interface{}

    // Number of stale iACKs dropped. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalAckDropsStale interface{}

    // Number of iACKs not held because of an immediate match. The type is
    // interface{} with range: 0..18446744073709551615.
    InternalAckDropsImmediateMatch interface{}

    // Time of last clear (in seconds since 1st Jan 1970 00:00:00). The type is
    // interface{} with range: 0..4294967295. Units are second.
    LastClearedTime interface{}

    // Send path counters for the PCB.
    SndCounters TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters
}

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetFilter() yfilter.YFilter { return statisticSession.YFilter }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) SetFilter(yf yfilter.YFilter) { statisticSession.YFilter = yf }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetGoName(yname string) string {
    if yname == "id" { return "Id" }
    if yname == "pcb" { return "Pcb" }
    if yname == "number-of-times-nsr-up" { return "NumberOfTimesNsrUp" }
    if yname == "number-of-timers-nsr-down" { return "NumberOfTimersNsrDown" }
    if yname == "number-of-times-nsr-disabled" { return "NumberOfTimesNsrDisabled" }
    if yname == "number-of-times-nsr-fail-over" { return "NumberOfTimesNsrFailOver" }
    if yname == "internal-ack-drops-not-replicated" { return "InternalAckDropsNotReplicated" }
    if yname == "internal-ack-drops-initsync-first-phase" { return "InternalAckDropsInitsyncFirstPhase" }
    if yname == "internal-ack-drops-stale" { return "InternalAckDropsStale" }
    if yname == "internal-ack-drops-immediate-match" { return "InternalAckDropsImmediateMatch" }
    if yname == "last-cleared-time" { return "LastClearedTime" }
    if yname == "snd-counters" { return "SndCounters" }
    return ""
}

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetSegmentPath() string {
    return "statistic-session" + "[id='" + fmt.Sprintf("%v", statisticSession.Id) + "']"
}

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snd-counters" {
        return &statisticSession.SndCounters
    }
    return nil
}

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snd-counters"] = &statisticSession.SndCounters
    return children
}

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["id"] = statisticSession.Id
    leafs["pcb"] = statisticSession.Pcb
    leafs["number-of-times-nsr-up"] = statisticSession.NumberOfTimesNsrUp
    leafs["number-of-timers-nsr-down"] = statisticSession.NumberOfTimersNsrDown
    leafs["number-of-times-nsr-disabled"] = statisticSession.NumberOfTimesNsrDisabled
    leafs["number-of-times-nsr-fail-over"] = statisticSession.NumberOfTimesNsrFailOver
    leafs["internal-ack-drops-not-replicated"] = statisticSession.InternalAckDropsNotReplicated
    leafs["internal-ack-drops-initsync-first-phase"] = statisticSession.InternalAckDropsInitsyncFirstPhase
    leafs["internal-ack-drops-stale"] = statisticSession.InternalAckDropsStale
    leafs["internal-ack-drops-immediate-match"] = statisticSession.InternalAckDropsImmediateMatch
    leafs["last-cleared-time"] = statisticSession.LastClearedTime
    return leafs
}

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetBundleName() string { return "cisco_ios_xr" }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetYangName() string { return "statistic-session" }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) SetParent(parent types.Entity) { statisticSession.parent = parent }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetParent() types.Entity { return statisticSession.parent }

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetParentYangName() string { return "statistic-sessions" }

// TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters
// Send path counters for the PCB
type TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of successful DATA transfers. The type is interface{} with range:
    // 0..4294967295.
    DataXferSend interface{}

    // Amount of data transferred. The type is interface{} with range:
    // 0..18446744073709551615.
    DataXferSendTotal interface{}

    // Number of failed DATA transfers. The type is interface{} with range:
    // 0..4294967295.
    DataXferSendDrop interface{}

    // Number of data transfer msgs., that required new IOV's to be allocated. The
    // type is interface{} with range: 0..4294967295.
    DataXferSendIovAlloc interface{}

    // Number of received DATA transfers. The type is interface{} with range:
    // 0..4294967295.
    DataXferRcv interface{}

    // Number of successfully received DATA transfers. The type is interface{}
    // with range: 0..4294967295.
    DataXferRcvSuccess interface{}

    // Number of received DATA transfers that had buffer trim failures. The type
    // is interface{} with range: 0..4294967295.
    DataXferRcvFailBufferTrim interface{}

    // Number of received DATA transfers that had failures because the send path
    // was out of sync. The type is interface{} with range: 0..4294967295.
    DataXferRcvFailSndUnaOutOfSync interface{}

    // Number of successful Segmentation instruction messages. The type is
    // interface{} with range: 0..4294967295.
    SegInstrSend interface{}

    // Number of segement units transferred via the successful Segmentation
    // instruction messages. The type is interface{} with range: 0..4294967295.
    SegInstrSendUnits interface{}

    // Number of failed Segmentation instruction messages. The type is interface{}
    // with range: 0..4294967295.
    SegInstrSendDrop interface{}

    // Number of received Segmentation instruction messages. The type is
    // interface{} with range: 0..4294967295.
    SegInstrRcv interface{}

    // Number of successfully received Segmentation instruction messages. The type
    // is interface{} with range: 0..4294967295.
    SegInstrRcvSuccess interface{}

    // Number of received Segmentation instructions that had buffer trim failures.
    // The type is interface{} with range: 0..4294967295.
    SegInstrRcvFailBufferTrim interface{}

    // Number of received Segmentation instructions that had failures during TCP
    // processing. The type is interface{} with range: 0..4294967295.
    SegInstrRcvFailTcpProcess interface{}

    // Number of successful NACK messages. The type is interface{} with range:
    // 0..4294967295.
    NackSend interface{}

    // Number of failed NACK messages. The type is interface{} with range:
    // 0..4294967295.
    NackSendDrop interface{}

    // Number of received NACK messages. The type is interface{} with range:
    // 0..4294967295.
    NackRcv interface{}

    // Number of successfully received NACK messages. The type is interface{} with
    // range: 0..4294967295.
    NackRcvSuccess interface{}

    // Number of received NACK messages that had failures when sending data in
    // response to the NACK. The type is interface{} with range: 0..4294967295.
    NackRcvFailDataSend interface{}

    // Number of successful Cleanup messages. The type is interface{} with range:
    // 0..4294967295.
    CleanupSend interface{}

    // Number of failed Cleanup messages. The type is interface{} with range:
    // 0..4294967295.
    CleanupSendDrop interface{}

    // Number of received Cleanup messages. The type is interface{} with range:
    // 0..4294967295.
    CleanupRcv interface{}

    // Number of successfully received Cleanup messages. The type is interface{}
    // with range: 0..4294967295.
    CleanupRcvSuccess interface{}

    // Number of Cleanup messages that had trim failures. The type is interface{}
    // with range: 0..4294967295.
    CleanupRcvFailBufferTrim interface{}
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetFilter() yfilter.YFilter { return sndCounters.YFilter }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) SetFilter(yf yfilter.YFilter) { sndCounters.YFilter = yf }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetGoName(yname string) string {
    if yname == "data-xfer-send" { return "DataXferSend" }
    if yname == "data-xfer-send-total" { return "DataXferSendTotal" }
    if yname == "data-xfer-send-drop" { return "DataXferSendDrop" }
    if yname == "data-xfer-send-iov-alloc" { return "DataXferSendIovAlloc" }
    if yname == "data-xfer-rcv" { return "DataXferRcv" }
    if yname == "data-xfer-rcv-success" { return "DataXferRcvSuccess" }
    if yname == "data-xfer-rcv-fail-buffer-trim" { return "DataXferRcvFailBufferTrim" }
    if yname == "data-xfer-rcv-fail-snd-una-out-of-sync" { return "DataXferRcvFailSndUnaOutOfSync" }
    if yname == "seg-instr-send" { return "SegInstrSend" }
    if yname == "seg-instr-send-units" { return "SegInstrSendUnits" }
    if yname == "seg-instr-send-drop" { return "SegInstrSendDrop" }
    if yname == "seg-instr-rcv" { return "SegInstrRcv" }
    if yname == "seg-instr-rcv-success" { return "SegInstrRcvSuccess" }
    if yname == "seg-instr-rcv-fail-buffer-trim" { return "SegInstrRcvFailBufferTrim" }
    if yname == "seg-instr-rcv-fail-tcp-process" { return "SegInstrRcvFailTcpProcess" }
    if yname == "nack-send" { return "NackSend" }
    if yname == "nack-send-drop" { return "NackSendDrop" }
    if yname == "nack-rcv" { return "NackRcv" }
    if yname == "nack-rcv-success" { return "NackRcvSuccess" }
    if yname == "nack-rcv-fail-data-send" { return "NackRcvFailDataSend" }
    if yname == "cleanup-send" { return "CleanupSend" }
    if yname == "cleanup-send-drop" { return "CleanupSendDrop" }
    if yname == "cleanup-rcv" { return "CleanupRcv" }
    if yname == "cleanup-rcv-success" { return "CleanupRcvSuccess" }
    if yname == "cleanup-rcv-fail-buffer-trim" { return "CleanupRcvFailBufferTrim" }
    return ""
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetSegmentPath() string {
    return "snd-counters"
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["data-xfer-send"] = sndCounters.DataXferSend
    leafs["data-xfer-send-total"] = sndCounters.DataXferSendTotal
    leafs["data-xfer-send-drop"] = sndCounters.DataXferSendDrop
    leafs["data-xfer-send-iov-alloc"] = sndCounters.DataXferSendIovAlloc
    leafs["data-xfer-rcv"] = sndCounters.DataXferRcv
    leafs["data-xfer-rcv-success"] = sndCounters.DataXferRcvSuccess
    leafs["data-xfer-rcv-fail-buffer-trim"] = sndCounters.DataXferRcvFailBufferTrim
    leafs["data-xfer-rcv-fail-snd-una-out-of-sync"] = sndCounters.DataXferRcvFailSndUnaOutOfSync
    leafs["seg-instr-send"] = sndCounters.SegInstrSend
    leafs["seg-instr-send-units"] = sndCounters.SegInstrSendUnits
    leafs["seg-instr-send-drop"] = sndCounters.SegInstrSendDrop
    leafs["seg-instr-rcv"] = sndCounters.SegInstrRcv
    leafs["seg-instr-rcv-success"] = sndCounters.SegInstrRcvSuccess
    leafs["seg-instr-rcv-fail-buffer-trim"] = sndCounters.SegInstrRcvFailBufferTrim
    leafs["seg-instr-rcv-fail-tcp-process"] = sndCounters.SegInstrRcvFailTcpProcess
    leafs["nack-send"] = sndCounters.NackSend
    leafs["nack-send-drop"] = sndCounters.NackSendDrop
    leafs["nack-rcv"] = sndCounters.NackRcv
    leafs["nack-rcv-success"] = sndCounters.NackRcvSuccess
    leafs["nack-rcv-fail-data-send"] = sndCounters.NackRcvFailDataSend
    leafs["cleanup-send"] = sndCounters.CleanupSend
    leafs["cleanup-send-drop"] = sndCounters.CleanupSendDrop
    leafs["cleanup-rcv"] = sndCounters.CleanupRcv
    leafs["cleanup-rcv-success"] = sndCounters.CleanupRcvSuccess
    leafs["cleanup-rcv-fail-buffer-trim"] = sndCounters.CleanupRcvFailBufferTrim
    return leafs
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetBundleName() string { return "cisco_ios_xr" }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetYangName() string { return "snd-counters" }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) SetParent(parent types.Entity) { sndCounters.parent = parent }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetParent() types.Entity { return sndCounters.parent }

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetParentYangName() string { return "statistic-session" }

