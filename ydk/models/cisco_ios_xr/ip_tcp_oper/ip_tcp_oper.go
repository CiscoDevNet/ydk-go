// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-tcp package operational data.
// 
// This module contains definitions
// for the following management objects:
//   tcp-connection: TCP connection operational data
//   tcp: tcp
//   tcp-nsr: tcp nsr
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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

// MessageTypeIgmp_ represents LPTS IGMP message types
type MessageTypeIgmp_ string

const (
    // IGMP Packet type: Membership query
    MessageTypeIgmp__membership_query MessageTypeIgmp_ = "membership-query"

    // IGMP Packet type: V1 membership report
    MessageTypeIgmp__v1_membership_report MessageTypeIgmp_ = "v1-membership-report"

    // IGMP Packet type: DVMRP
    MessageTypeIgmp__dvmrp MessageTypeIgmp_ = "dvmrp"

    // IGMP Packet type: PIM version 1
    MessageTypeIgmp__pi_mv1 MessageTypeIgmp_ = "pi-mv1"

    // IGMP Packet type: Cisco Trace Messages
    MessageTypeIgmp__cisco_trace_messages MessageTypeIgmp_ = "cisco-trace-messages"

    // IGMP Packet type: V2 membership report
    MessageTypeIgmp__v2_membership_report MessageTypeIgmp_ = "v2-membership-report"

    // IGMP Packet type: V2 leave group
    MessageTypeIgmp__v2_leave_group MessageTypeIgmp_ = "v2-leave-group"

    // IGMP Packet type: Multicast traceroute response
    MessageTypeIgmp__multicast_traceroute_response MessageTypeIgmp_ = "multicast-traceroute-response"

    // IGMP Packet type: MulticastTraceroute
    MessageTypeIgmp__multicast_traceroute MessageTypeIgmp_ = "multicast-traceroute"

    // IGMP Packet type: V3 membership report
    MessageTypeIgmp__v3_membership_report MessageTypeIgmp_ = "v3-membership-report"

    // IGMP Packet type: Multicast router
    // advertisement
    MessageTypeIgmp__multicast_router_advertisement MessageTypeIgmp_ = "multicast-router-advertisement"

    // IGMP Packet type: Multicast router solicitation
    MessageTypeIgmp__multicast_router_solicitation MessageTypeIgmp_ = "multicast-router-solicitation"

    // IGMP Packet type: Multicast router termination
    MessageTypeIgmp__multicast_router_termination MessageTypeIgmp_ = "multicast-router-termination"
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

// MessageTypeIcmp_ represents LPTS ICMP message types
type MessageTypeIcmp_ string

const (
    // ICMP Packet type: Echo reply
    MessageTypeIcmp__echo_reply MessageTypeIcmp_ = "echo-reply"

    // ICMP Packet type: Destination unreachable
    MessageTypeIcmp__destination_unreachable MessageTypeIcmp_ = "destination-unreachable"

    // ICMP Packet type: Source quench
    MessageTypeIcmp__source_quench MessageTypeIcmp_ = "source-quench"

    // ICMP Packet type: Redirect
    MessageTypeIcmp__redirect MessageTypeIcmp_ = "redirect"

    // ICMP Packet type: Alternate host address
    MessageTypeIcmp__alternate_host_address MessageTypeIcmp_ = "alternate-host-address"

    // ICMP Packet type: Echo
    MessageTypeIcmp__echo MessageTypeIcmp_ = "echo"

    // ICMP Packet type: Router advertisement
    MessageTypeIcmp__router_advertisement MessageTypeIcmp_ = "router-advertisement"

    // ICMP Packet type: Router selection
    MessageTypeIcmp__router_selection MessageTypeIcmp_ = "router-selection"

    // ICMP Packet type: Time exceeded
    MessageTypeIcmp__time_exceeded MessageTypeIcmp_ = "time-exceeded"

    // ICMP Packet type: Parameter problem
    MessageTypeIcmp__parameter_problem MessageTypeIcmp_ = "parameter-problem"

    // ICMP Packet type: Time stamp
    MessageTypeIcmp__time_stamp MessageTypeIcmp_ = "time-stamp"

    // ICMP Packet type: Time stamp reply
    MessageTypeIcmp__time_stamp_reply MessageTypeIcmp_ = "time-stamp-reply"

    // ICMP Packet type: Information request
    MessageTypeIcmp__information_request MessageTypeIcmp_ = "information-request"

    // ICMP Packet type: Information reply
    MessageTypeIcmp__information_reply MessageTypeIcmp_ = "information-reply"

    // ICMP Packet type: Address mask request
    MessageTypeIcmp__address_mask_request MessageTypeIcmp_ = "address-mask-request"

    // ICMP Packet type: Address mask reply
    MessageTypeIcmp__address_mask_reply MessageTypeIcmp_ = "address-mask-reply"

    // ICMP Packet type: Trace route
    MessageTypeIcmp__trace_route MessageTypeIcmp_ = "trace-route"

    // ICMP Packet type: Datagram Conversion error
    MessageTypeIcmp__datagram_conversion_error MessageTypeIcmp_ = "datagram-conversion-error"

    // ICMP Packet type: Mobile host redirect
    MessageTypeIcmp__mobile_host_redirect MessageTypeIcmp_ = "mobile-host-redirect"

    // ICMP Packet type: IPv6 where-are-you
    MessageTypeIcmp__where_are_you MessageTypeIcmp_ = "where-are-you"

    // ICMP Packet type: IPv6 i-am-here
    MessageTypeIcmp__iam_here MessageTypeIcmp_ = "iam-here"

    // ICMP Packet type: Mobile registration request
    MessageTypeIcmp__mobile_registration_request MessageTypeIcmp_ = "mobile-registration-request"

    // ICMP Packet type: Mobile registration reply
    MessageTypeIcmp__mobile_registration_reply MessageTypeIcmp_ = "mobile-registration-reply"

    // ICMP Packet type: Domain name request
    MessageTypeIcmp__domain_name_request MessageTypeIcmp_ = "domain-name-request"
)

// TcpKeyInvalidReason represents TCP AO key state invalid reason
type TcpKeyInvalidReason string

const (
    // No reason
    TcpKeyInvalidReason_none TcpKeyInvalidReason = "none"

    // Incomplete
    TcpKeyInvalidReason_incomplete TcpKeyInvalidReason = "incomplete"

    // Send and accept lifetime are not same
    TcpKeyInvalidReason_lifetime_not_same TcpKeyInvalidReason = "lifetime-not-same"

    // Send ID is invalid
    TcpKeyInvalidReason_send_id_invalid TcpKeyInvalidReason = "send-id-invalid"

    // Receive ID is invalid
    TcpKeyInvalidReason_recv_id_invalid TcpKeyInvalidReason = "recv-id-invalid"
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

// TcpMacAlgo represents TCP AO MAC algorithm type
type TcpMacAlgo string

const (
    // Not configured
    TcpMacAlgo_not_configured TcpMacAlgo = "not-configured"

    // CMAC 96
    TcpMacAlgo_aes_128_cmac_96 TcpMacAlgo = "aes-128-cmac-96"

    // HMAC SHA1 12
    TcpMacAlgo_hmac_sha1_12 TcpMacAlgo = "hmac-sha1-12"

    // MD5 16
    TcpMacAlgo_md5_16 TcpMacAlgo = "md5-16"

    // SHA1 20
    TcpMacAlgo_sha1_20 TcpMacAlgo = "sha1-20"

    // HMAC MD5 16
    TcpMacAlgo_hmac_md5_16 TcpMacAlgo = "hmac-md5-16"

    // HMAC SHA1 20
    TcpMacAlgo_hmac_sha1_20 TcpMacAlgo = "hmac-sha1-20"

    // AES 128 CMAC
    TcpMacAlgo_aes_128_cmac TcpMacAlgo = "aes-128-cmac"

    // AES 256 CMAC
    TcpMacAlgo_aes_256_cmac TcpMacAlgo = "aes-256-cmac"

    // HMAC SHA1 96
    TcpMacAlgo_hmac_sha1_96 TcpMacAlgo = "hmac-sha1-96"

    // HMAC SHA1 256
    TcpMacAlgo_hmac_sha_256 TcpMacAlgo = "hmac-sha-256"
)

// TcpAddressFamily represents Address Family Type
type TcpAddressFamily string

const (
    // IPv4
    TcpAddressFamily_ipv4 TcpAddressFamily = "ipv4"

    // IPv6
    TcpAddressFamily_ipv6 TcpAddressFamily = "ipv6"
)

// MessageTypeIcmpv6_ represents LPTS ICMPv6 message types
type MessageTypeIcmpv6_ string

const (
    // ICMPv6 Packet type: Destination unreachable
    MessageTypeIcmpv6__destination_unreachable MessageTypeIcmpv6_ = "destination-unreachable"

    // ICMPv6 Packet type: packet too big
    MessageTypeIcmpv6__packet_too_big MessageTypeIcmpv6_ = "packet-too-big"

    // ICMPv6 Packet type: Time exceeded
    MessageTypeIcmpv6__time_exceeded MessageTypeIcmpv6_ = "time-exceeded"

    // ICMPv6 Packet type: Parameter problem
    MessageTypeIcmpv6__parameter_problem MessageTypeIcmpv6_ = "parameter-problem"

    // ICMPv6 Packet type: Echo request
    MessageTypeIcmpv6__echo_request MessageTypeIcmpv6_ = "echo-request"

    // ICMPv6 Packet type: Echo reply
    MessageTypeIcmpv6__echo_reply MessageTypeIcmpv6_ = "echo-reply"

    // ICMPv6 Packet type: Multicast listener query
    MessageTypeIcmpv6__multicast_listener_query MessageTypeIcmpv6_ = "multicast-listener-query"

    // ICMPv6 Packet type: Multicast listener report
    MessageTypeIcmpv6__multicast_listener_report MessageTypeIcmpv6_ = "multicast-listener-report"

    // ICMPv6 Packet type: Multicast listener done
    MessageTypeIcmpv6__multicast_listener_done MessageTypeIcmpv6_ = "multicast-listener-done"

    // ICMPv6 Packet type: Router solicitation
    MessageTypeIcmpv6__router_solicitation MessageTypeIcmpv6_ = "router-solicitation"

    // ICMPv6 Packet type: Router advertisement
    MessageTypeIcmpv6__router_advertisement MessageTypeIcmpv6_ = "router-advertisement"

    // ICMPv6 Packet type: Neighbor solicitation
    MessageTypeIcmpv6__neighbor_solicitation MessageTypeIcmpv6_ = "neighbor-solicitation"

    // ICMPv6 Packet type: Neighbor advertisement
    MessageTypeIcmpv6__neighbor_advertisement MessageTypeIcmpv6_ = "neighbor-advertisement"

    // ICMPv6 Packet type: Redirect message
    MessageTypeIcmpv6__redirect_message MessageTypeIcmpv6_ = "redirect-message"

    // ICMPv6 Packet type: Router renumbering
    MessageTypeIcmpv6__router_renumbering MessageTypeIcmpv6_ = "router-renumbering"

    // ICMPv6 Packet type: Node information query
    MessageTypeIcmpv6__node_information_query MessageTypeIcmpv6_ = "node-information-query"

    // ICMPv6 Packet type: Node information reply
    MessageTypeIcmpv6__node_information_reply MessageTypeIcmpv6_ = "node-information-reply"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // solicitation message
    MessageTypeIcmpv6__inverse_neighbor_discovery_solicitaion MessageTypeIcmpv6_ = "inverse-neighbor-discovery-solicitaion"

    // ICMPv6 Packet type: Inverse neighbor discovery
    // advertisement message
    MessageTypeIcmpv6__inverse_neighbor_discover_advertisement MessageTypeIcmpv6_ = "inverse-neighbor-discover-advertisement"

    // ICMPv6 Packet type: Version 2 multicast
    // listener report
    MessageTypeIcmpv6__v2_multicast_listener_report MessageTypeIcmpv6_ = "v2-multicast-listener-report"

    // ICMPv6 Packet type: Home agent address
    // discovery request message
    MessageTypeIcmpv6__home_agent_address_discovery_request MessageTypeIcmpv6_ = "home-agent-address-discovery-request"

    // ICMPv6 Packet type: Home agent address
    // discovery reply message
    MessageTypeIcmpv6__home_agent_address_discovery_reply MessageTypeIcmpv6_ = "home-agent-address-discovery-reply"

    // ICMPv6 Packet type: Mobile prefix solicitation
    MessageTypeIcmpv6__mobile_prefix_solicitation MessageTypeIcmpv6_ = "mobile-prefix-solicitation"

    // ICMPv6 Packet type: Mobile prefix advertisement
    MessageTypeIcmpv6__mobile_prefix_advertisement MessageTypeIcmpv6_ = "mobile-prefix-advertisement"

    // ICMPv6 Packet type: Certification path
    // solicitation message
    MessageTypeIcmpv6__certification_path_solicitation_message MessageTypeIcmpv6_ = "certification-path-solicitation-message"

    // ICMPv6 Packet type: Certification path
    // advertisement message
    MessageTypeIcmpv6__certification_path_advertisement_message MessageTypeIcmpv6_ = "certification-path-advertisement-message"

    // ICMPv6 Packet type: ICMP messages utilized by
    // experimental mobility  protocols such as
    // seamoby
    MessageTypeIcmpv6__experimental_mobility_protocols MessageTypeIcmpv6_ = "experimental-mobility-protocols"

    // ICMPv6 Packet type: Multicast router
    // advertisement
    MessageTypeIcmpv6__multicast_router_advertisement MessageTypeIcmpv6_ = "multicast-router-advertisement"

    // ICMPv6 Packet type: Multicast router
    // solicitation
    MessageTypeIcmpv6__multicast_router_solicitation MessageTypeIcmpv6_ = "multicast-router-solicitation"

    // ICMPv6 Packet type: Multicast router
    // termination
    MessageTypeIcmpv6__multicast_router_termination MessageTypeIcmpv6_ = "multicast-router-termination"

    // ICMPv6 Packet type: FMIPv6 messages
    MessageTypeIcmpv6__fmipv6_messages MessageTypeIcmpv6_ = "fmipv6-messages"
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of information about all nodes present on the system.
    Nodes TcpConnection_Nodes
}

func (tcpConnection *TcpConnection) GetEntityData() *types.CommonEntityData {
    tcpConnection.EntityData.YFilter = tcpConnection.YFilter
    tcpConnection.EntityData.YangName = "tcp-connection"
    tcpConnection.EntityData.BundleName = "cisco_ios_xr"
    tcpConnection.EntityData.ParentYangName = "Cisco-IOS-XR-ip-tcp-oper"
    tcpConnection.EntityData.SegmentPath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection"
    tcpConnection.EntityData.AbsolutePath = tcpConnection.EntityData.SegmentPath
    tcpConnection.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcpConnection.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcpConnection.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcpConnection.EntityData.Children = types.NewOrderedMap()
    tcpConnection.EntityData.Children.Append("nodes", types.YChild{"Nodes", &tcpConnection.Nodes})
    tcpConnection.EntityData.Leafs = types.NewOrderedMap()

    tcpConnection.EntityData.YListKeys = []string {}

    return &(tcpConnection.EntityData)
}

// TcpConnection_Nodes
// Table of information about all nodes present on
// the system
type TcpConnection_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single node. The type is slice of
    // TcpConnection_Nodes_Node.
    Node []*TcpConnection_Nodes_Node
}

func (nodes *TcpConnection_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "tcp-connection"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// TcpConnection_Nodes_Node
// Information about a single node
type TcpConnection_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Describing a location. The type is string with
    // pattern: ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Id interface{}

    // Statistics of all TCP connections.
    Statistics TcpConnection_Nodes_Node_Statistics

    // Extended Filter related Information.
    ExtendedInformation TcpConnection_Nodes_Node_ExtendedInformation

    // Table listing TCP connections for which detailed information is provided.
    DetailInformations TcpConnection_Nodes_Node_DetailInformations

    // Table listing keychains configured for TCP-AO.
    Keychains TcpConnection_Nodes_Node_Keychains

    // Table listing connections for which brief information is provided.Note that
    // not all connections are listed in the brief table.
    BriefInformations TcpConnection_Nodes_Node_BriefInformations
}

func (node *TcpConnection_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Id, "id")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Children.Append("extended-information", types.YChild{"ExtendedInformation", &node.ExtendedInformation})
    node.EntityData.Children.Append("detail-informations", types.YChild{"DetailInformations", &node.DetailInformations})
    node.EntityData.Children.Append("keychains", types.YChild{"Keychains", &node.Keychains})
    node.EntityData.Children.Append("brief-informations", types.YChild{"BriefInformations", &node.BriefInformations})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("id", types.YLeaf{"Id", node.Id})

    node.EntityData.YListKeys = []string {"Id"}

    return &(node.EntityData)
}

// TcpConnection_Nodes_Node_Statistics
// Statistics of all TCP connections
type TcpConnection_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table listing clients.
    Clients TcpConnection_Nodes_Node_Statistics_Clients

    // Table listing the TCP connections for which statistics are provided.
    Pcbs TcpConnection_Nodes_Node_Statistics_Pcbs

    // Summary statistics across all TCP connections.
    Summary TcpConnection_Nodes_Node_Statistics_Summary
}

func (statistics *TcpConnection_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("clients", types.YChild{"Clients", &statistics.Clients})
    statistics.EntityData.Children.Append("pcbs", types.YChild{"Pcbs", &statistics.Pcbs})
    statistics.EntityData.Children.Append("summary", types.YChild{"Summary", &statistics.Summary})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Clients
// Table listing clients
type TcpConnection_Nodes_Node_Statistics_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing Client ID. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Clients_Client.
    Client []*TcpConnection_Nodes_Node_Statistics_Clients_Client
}

func (clients *TcpConnection_Nodes_Node_Statistics_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "statistics"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/" + clients.EntityData.SegmentPath
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

// TcpConnection_Nodes_Node_Statistics_Clients_Client
// Describing Client ID
type TcpConnection_Nodes_Node_Statistics_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (client *TcpConnection_Nodes_Node_Statistics_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + types.AddKeyToken(client.ClientId, "client-id")
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/clients/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Leafs = types.NewOrderedMap()
    client.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", client.ClientId})
    client.EntityData.Leafs.Append("client-jid", types.YLeaf{"ClientJid", client.ClientJid})
    client.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", client.ClientName})
    client.EntityData.Leafs.Append("ipv4-received-packets", types.YLeaf{"Ipv4ReceivedPackets", client.Ipv4ReceivedPackets})
    client.EntityData.Leafs.Append("ipv4-sent-packets", types.YLeaf{"Ipv4SentPackets", client.Ipv4SentPackets})
    client.EntityData.Leafs.Append("ipv6-received-packets", types.YLeaf{"Ipv6ReceivedPackets", client.Ipv6ReceivedPackets})
    client.EntityData.Leafs.Append("ipv6-sent-packets", types.YLeaf{"Ipv6SentPackets", client.Ipv6SentPackets})

    client.EntityData.YListKeys = []string {"ClientId"}

    return &(client.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs
// Table listing the TCP connections for which
// statistics are provided
type TcpConnection_Nodes_Node_Statistics_Pcbs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Control Block ID. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb.
    Pcb []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb
}

func (pcbs *TcpConnection_Nodes_Node_Statistics_Pcbs) GetEntityData() *types.CommonEntityData {
    pcbs.EntityData.YFilter = pcbs.YFilter
    pcbs.EntityData.YangName = "pcbs"
    pcbs.EntityData.BundleName = "cisco_ios_xr"
    pcbs.EntityData.ParentYangName = "statistics"
    pcbs.EntityData.SegmentPath = "pcbs"
    pcbs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/" + pcbs.EntityData.SegmentPath
    pcbs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcbs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcbs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcbs.EntityData.Children = types.NewOrderedMap()
    pcbs.EntityData.Children.Append("pcb", types.YChild{"Pcb", nil})
    for i := range pcbs.Pcb {
        pcbs.EntityData.Children.Append(types.GetSegmentPath(pcbs.Pcb[i]), types.YChild{"Pcb", pcbs.Pcb[i]})
    }
    pcbs.EntityData.Leafs = types.NewOrderedMap()

    pcbs.EntityData.YListKeys = []string {}

    return &(pcbs.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb
// Protocol Control Block ID
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Displaying statistics associated with a particular
    // PCB. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (pcb *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb) GetEntityData() *types.CommonEntityData {
    pcb.EntityData.YFilter = pcb.YFilter
    pcb.EntityData.YangName = "pcb"
    pcb.EntityData.BundleName = "cisco_ios_xr"
    pcb.EntityData.ParentYangName = "pcbs"
    pcb.EntityData.SegmentPath = "pcb" + types.AddKeyToken(pcb.Id, "id")
    pcb.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/" + pcb.EntityData.SegmentPath
    pcb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pcb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pcb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pcb.EntityData.Children = types.NewOrderedMap()
    pcb.EntityData.Children.Append("read-io-counts", types.YChild{"ReadIoCounts", &pcb.ReadIoCounts})
    pcb.EntityData.Children.Append("write-io-counts", types.YChild{"WriteIoCounts", &pcb.WriteIoCounts})
    pcb.EntityData.Children.Append("async-session-stats", types.YChild{"AsyncSessionStats", &pcb.AsyncSessionStats})
    pcb.EntityData.Leafs = types.NewOrderedMap()
    pcb.EntityData.Leafs.Append("id", types.YLeaf{"Id", pcb.Id})
    pcb.EntityData.Leafs.Append("pcb", types.YLeaf{"Pcb", pcb.Pcb})
    pcb.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", pcb.VrfId})
    pcb.EntityData.Leafs.Append("packets-sent", types.YLeaf{"PacketsSent", pcb.PacketsSent})
    pcb.EntityData.Leafs.Append("xipc-pulse-received", types.YLeaf{"XipcPulseReceived", pcb.XipcPulseReceived})
    pcb.EntityData.Leafs.Append("segment-instruction-received", types.YLeaf{"SegmentInstructionReceived", pcb.SegmentInstructionReceived})
    pcb.EntityData.Leafs.Append("send-packets-queued", types.YLeaf{"SendPacketsQueued", pcb.SendPacketsQueued})
    pcb.EntityData.Leafs.Append("send-packets-queued-net-io", types.YLeaf{"SendPacketsQueuedNetIo", pcb.SendPacketsQueuedNetIo})
    pcb.EntityData.Leafs.Append("send-queue-failed", types.YLeaf{"SendQueueFailed", pcb.SendQueueFailed})
    pcb.EntityData.Leafs.Append("send-queue-net-io-failed", types.YLeaf{"SendQueueNetIoFailed", pcb.SendQueueNetIoFailed})
    pcb.EntityData.Leafs.Append("packets-received", types.YLeaf{"PacketsReceived", pcb.PacketsReceived})
    pcb.EntityData.Leafs.Append("receive-queue-failed", types.YLeaf{"ReceiveQueueFailed", pcb.ReceiveQueueFailed})
    pcb.EntityData.Leafs.Append("received-packets-queued", types.YLeaf{"ReceivedPacketsQueued", pcb.ReceivedPacketsQueued})
    pcb.EntityData.Leafs.Append("send-window-shrink-ignored", types.YLeaf{"SendWindowShrinkIgnored", pcb.SendWindowShrinkIgnored})
    pcb.EntityData.Leafs.Append("is-paw-socket", types.YLeaf{"IsPawSocket", pcb.IsPawSocket})
    pcb.EntityData.Leafs.Append("read-io-time", types.YLeaf{"ReadIoTime", pcb.ReadIoTime})
    pcb.EntityData.Leafs.Append("write-io-time", types.YLeaf{"WriteIoTime", pcb.WriteIoTime})

    pcb.EntityData.YListKeys = []string {"Id"}

    return &(pcb.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts
// Read  I/O counts
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts struct {
    EntityData types.CommonEntityData
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

func (readIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_ReadIoCounts) GetEntityData() *types.CommonEntityData {
    readIoCounts.EntityData.YFilter = readIoCounts.YFilter
    readIoCounts.EntityData.YangName = "read-io-counts"
    readIoCounts.EntityData.BundleName = "cisco_ios_xr"
    readIoCounts.EntityData.ParentYangName = "pcb"
    readIoCounts.EntityData.SegmentPath = "read-io-counts"
    readIoCounts.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/" + readIoCounts.EntityData.SegmentPath
    readIoCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    readIoCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    readIoCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    readIoCounts.EntityData.Children = types.NewOrderedMap()
    readIoCounts.EntityData.Leafs = types.NewOrderedMap()
    readIoCounts.EntityData.Leafs.Append("io-count", types.YLeaf{"IoCount", readIoCounts.IoCount})
    readIoCounts.EntityData.Leafs.Append("arm-count", types.YLeaf{"ArmCount", readIoCounts.ArmCount})
    readIoCounts.EntityData.Leafs.Append("unarm-count", types.YLeaf{"UnarmCount", readIoCounts.UnarmCount})
    readIoCounts.EntityData.Leafs.Append("autoarm-count", types.YLeaf{"AutoarmCount", readIoCounts.AutoarmCount})

    readIoCounts.EntityData.YListKeys = []string {}

    return &(readIoCounts.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts
// Write I/O counts
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts struct {
    EntityData types.CommonEntityData
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

func (writeIoCounts *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_WriteIoCounts) GetEntityData() *types.CommonEntityData {
    writeIoCounts.EntityData.YFilter = writeIoCounts.YFilter
    writeIoCounts.EntityData.YangName = "write-io-counts"
    writeIoCounts.EntityData.BundleName = "cisco_ios_xr"
    writeIoCounts.EntityData.ParentYangName = "pcb"
    writeIoCounts.EntityData.SegmentPath = "write-io-counts"
    writeIoCounts.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/" + writeIoCounts.EntityData.SegmentPath
    writeIoCounts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    writeIoCounts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    writeIoCounts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    writeIoCounts.EntityData.Children = types.NewOrderedMap()
    writeIoCounts.EntityData.Leafs = types.NewOrderedMap()
    writeIoCounts.EntityData.Leafs.Append("io-count", types.YLeaf{"IoCount", writeIoCounts.IoCount})
    writeIoCounts.EntityData.Leafs.Append("arm-count", types.YLeaf{"ArmCount", writeIoCounts.ArmCount})
    writeIoCounts.EntityData.Leafs.Append("unarm-count", types.YLeaf{"UnarmCount", writeIoCounts.UnarmCount})
    writeIoCounts.EntityData.Leafs.Append("autoarm-count", types.YLeaf{"AutoarmCount", writeIoCounts.AutoarmCount})

    writeIoCounts.EntityData.YListKeys = []string {}

    return &(writeIoCounts.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats
// Statistics of Async TCP Sessions
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Flag of async session. The type is bool.
    AsyncSession interface{}

    // Number of successful data write to XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteSuccessNum.
    DataWriteSuccessNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteSuccessNum

    // Number of successful data read from XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadSuccessNum.
    DataReadSuccessNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadSuccessNum

    // Number of failed data write to XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteErrorNum.
    DataWriteErrorNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteErrorNum

    // Number of failed data read from XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadErrorNum.
    DataReadErrorNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadErrorNum

    // Number of successful control write to XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteSuccessNum.
    ControlWriteSuccessNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteSuccessNum

    // Number of successful control read to XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadSuccessNum.
    ControlReadSuccessNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadSuccessNum

    // Number of failed control write to XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteErrorNum.
    ControlWriteErrorNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteErrorNum

    // Number of failed control read from XIPC. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadErrorNum.
    ControlReadErrorNum []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadErrorNum

    // Number of bytes data has been written. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteByte.
    DataWriteByte []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteByte

    // Number of bytes data has been read. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadByte.
    DataReadByte []*TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadByte
}

func (asyncSessionStats *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats) GetEntityData() *types.CommonEntityData {
    asyncSessionStats.EntityData.YFilter = asyncSessionStats.YFilter
    asyncSessionStats.EntityData.YangName = "async-session-stats"
    asyncSessionStats.EntityData.BundleName = "cisco_ios_xr"
    asyncSessionStats.EntityData.ParentYangName = "pcb"
    asyncSessionStats.EntityData.SegmentPath = "async-session-stats"
    asyncSessionStats.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/" + asyncSessionStats.EntityData.SegmentPath
    asyncSessionStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    asyncSessionStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    asyncSessionStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    asyncSessionStats.EntityData.Children = types.NewOrderedMap()
    asyncSessionStats.EntityData.Children.Append("data-write-success-num", types.YChild{"DataWriteSuccessNum", nil})
    for i := range asyncSessionStats.DataWriteSuccessNum {
        types.SetYListKey(asyncSessionStats.DataWriteSuccessNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.DataWriteSuccessNum[i]), types.YChild{"DataWriteSuccessNum", asyncSessionStats.DataWriteSuccessNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("data-read-success-num", types.YChild{"DataReadSuccessNum", nil})
    for i := range asyncSessionStats.DataReadSuccessNum {
        types.SetYListKey(asyncSessionStats.DataReadSuccessNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.DataReadSuccessNum[i]), types.YChild{"DataReadSuccessNum", asyncSessionStats.DataReadSuccessNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("data-write-error-num", types.YChild{"DataWriteErrorNum", nil})
    for i := range asyncSessionStats.DataWriteErrorNum {
        types.SetYListKey(asyncSessionStats.DataWriteErrorNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.DataWriteErrorNum[i]), types.YChild{"DataWriteErrorNum", asyncSessionStats.DataWriteErrorNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("data-read-error-num", types.YChild{"DataReadErrorNum", nil})
    for i := range asyncSessionStats.DataReadErrorNum {
        types.SetYListKey(asyncSessionStats.DataReadErrorNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.DataReadErrorNum[i]), types.YChild{"DataReadErrorNum", asyncSessionStats.DataReadErrorNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("control-write-success-num", types.YChild{"ControlWriteSuccessNum", nil})
    for i := range asyncSessionStats.ControlWriteSuccessNum {
        types.SetYListKey(asyncSessionStats.ControlWriteSuccessNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.ControlWriteSuccessNum[i]), types.YChild{"ControlWriteSuccessNum", asyncSessionStats.ControlWriteSuccessNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("control-read-success-num", types.YChild{"ControlReadSuccessNum", nil})
    for i := range asyncSessionStats.ControlReadSuccessNum {
        types.SetYListKey(asyncSessionStats.ControlReadSuccessNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.ControlReadSuccessNum[i]), types.YChild{"ControlReadSuccessNum", asyncSessionStats.ControlReadSuccessNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("control-write-error-num", types.YChild{"ControlWriteErrorNum", nil})
    for i := range asyncSessionStats.ControlWriteErrorNum {
        types.SetYListKey(asyncSessionStats.ControlWriteErrorNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.ControlWriteErrorNum[i]), types.YChild{"ControlWriteErrorNum", asyncSessionStats.ControlWriteErrorNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("control-read-error-num", types.YChild{"ControlReadErrorNum", nil})
    for i := range asyncSessionStats.ControlReadErrorNum {
        types.SetYListKey(asyncSessionStats.ControlReadErrorNum[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.ControlReadErrorNum[i]), types.YChild{"ControlReadErrorNum", asyncSessionStats.ControlReadErrorNum[i]})
    }
    asyncSessionStats.EntityData.Children.Append("data-write-byte", types.YChild{"DataWriteByte", nil})
    for i := range asyncSessionStats.DataWriteByte {
        types.SetYListKey(asyncSessionStats.DataWriteByte[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.DataWriteByte[i]), types.YChild{"DataWriteByte", asyncSessionStats.DataWriteByte[i]})
    }
    asyncSessionStats.EntityData.Children.Append("data-read-byte", types.YChild{"DataReadByte", nil})
    for i := range asyncSessionStats.DataReadByte {
        types.SetYListKey(asyncSessionStats.DataReadByte[i], i)
        asyncSessionStats.EntityData.Children.Append(types.GetSegmentPath(asyncSessionStats.DataReadByte[i]), types.YChild{"DataReadByte", asyncSessionStats.DataReadByte[i]})
    }
    asyncSessionStats.EntityData.Leafs = types.NewOrderedMap()
    asyncSessionStats.EntityData.Leafs.Append("async-session", types.YLeaf{"AsyncSession", asyncSessionStats.AsyncSession})

    asyncSessionStats.EntityData.YListKeys = []string {}

    return &(asyncSessionStats.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteSuccessNum
// Number of successful data write to XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteSuccessNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (dataWriteSuccessNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteSuccessNum) GetEntityData() *types.CommonEntityData {
    dataWriteSuccessNum.EntityData.YFilter = dataWriteSuccessNum.YFilter
    dataWriteSuccessNum.EntityData.YangName = "data-write-success-num"
    dataWriteSuccessNum.EntityData.BundleName = "cisco_ios_xr"
    dataWriteSuccessNum.EntityData.ParentYangName = "async-session-stats"
    dataWriteSuccessNum.EntityData.SegmentPath = "data-write-success-num" + types.AddNoKeyToken(dataWriteSuccessNum)
    dataWriteSuccessNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + dataWriteSuccessNum.EntityData.SegmentPath
    dataWriteSuccessNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataWriteSuccessNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataWriteSuccessNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataWriteSuccessNum.EntityData.Children = types.NewOrderedMap()
    dataWriteSuccessNum.EntityData.Leafs = types.NewOrderedMap()
    dataWriteSuccessNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dataWriteSuccessNum.Entry})

    dataWriteSuccessNum.EntityData.YListKeys = []string {}

    return &(dataWriteSuccessNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadSuccessNum
// Number of successful data read from XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadSuccessNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (dataReadSuccessNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadSuccessNum) GetEntityData() *types.CommonEntityData {
    dataReadSuccessNum.EntityData.YFilter = dataReadSuccessNum.YFilter
    dataReadSuccessNum.EntityData.YangName = "data-read-success-num"
    dataReadSuccessNum.EntityData.BundleName = "cisco_ios_xr"
    dataReadSuccessNum.EntityData.ParentYangName = "async-session-stats"
    dataReadSuccessNum.EntityData.SegmentPath = "data-read-success-num" + types.AddNoKeyToken(dataReadSuccessNum)
    dataReadSuccessNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + dataReadSuccessNum.EntityData.SegmentPath
    dataReadSuccessNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataReadSuccessNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataReadSuccessNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataReadSuccessNum.EntityData.Children = types.NewOrderedMap()
    dataReadSuccessNum.EntityData.Leafs = types.NewOrderedMap()
    dataReadSuccessNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dataReadSuccessNum.Entry})

    dataReadSuccessNum.EntityData.YListKeys = []string {}

    return &(dataReadSuccessNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteErrorNum
// Number of failed data write to XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteErrorNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (dataWriteErrorNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteErrorNum) GetEntityData() *types.CommonEntityData {
    dataWriteErrorNum.EntityData.YFilter = dataWriteErrorNum.YFilter
    dataWriteErrorNum.EntityData.YangName = "data-write-error-num"
    dataWriteErrorNum.EntityData.BundleName = "cisco_ios_xr"
    dataWriteErrorNum.EntityData.ParentYangName = "async-session-stats"
    dataWriteErrorNum.EntityData.SegmentPath = "data-write-error-num" + types.AddNoKeyToken(dataWriteErrorNum)
    dataWriteErrorNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + dataWriteErrorNum.EntityData.SegmentPath
    dataWriteErrorNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataWriteErrorNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataWriteErrorNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataWriteErrorNum.EntityData.Children = types.NewOrderedMap()
    dataWriteErrorNum.EntityData.Leafs = types.NewOrderedMap()
    dataWriteErrorNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dataWriteErrorNum.Entry})

    dataWriteErrorNum.EntityData.YListKeys = []string {}

    return &(dataWriteErrorNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadErrorNum
// Number of failed data read from XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadErrorNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (dataReadErrorNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadErrorNum) GetEntityData() *types.CommonEntityData {
    dataReadErrorNum.EntityData.YFilter = dataReadErrorNum.YFilter
    dataReadErrorNum.EntityData.YangName = "data-read-error-num"
    dataReadErrorNum.EntityData.BundleName = "cisco_ios_xr"
    dataReadErrorNum.EntityData.ParentYangName = "async-session-stats"
    dataReadErrorNum.EntityData.SegmentPath = "data-read-error-num" + types.AddNoKeyToken(dataReadErrorNum)
    dataReadErrorNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + dataReadErrorNum.EntityData.SegmentPath
    dataReadErrorNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataReadErrorNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataReadErrorNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataReadErrorNum.EntityData.Children = types.NewOrderedMap()
    dataReadErrorNum.EntityData.Leafs = types.NewOrderedMap()
    dataReadErrorNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dataReadErrorNum.Entry})

    dataReadErrorNum.EntityData.YListKeys = []string {}

    return &(dataReadErrorNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteSuccessNum
// Number of successful control write to XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteSuccessNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (controlWriteSuccessNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteSuccessNum) GetEntityData() *types.CommonEntityData {
    controlWriteSuccessNum.EntityData.YFilter = controlWriteSuccessNum.YFilter
    controlWriteSuccessNum.EntityData.YangName = "control-write-success-num"
    controlWriteSuccessNum.EntityData.BundleName = "cisco_ios_xr"
    controlWriteSuccessNum.EntityData.ParentYangName = "async-session-stats"
    controlWriteSuccessNum.EntityData.SegmentPath = "control-write-success-num" + types.AddNoKeyToken(controlWriteSuccessNum)
    controlWriteSuccessNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + controlWriteSuccessNum.EntityData.SegmentPath
    controlWriteSuccessNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlWriteSuccessNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlWriteSuccessNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlWriteSuccessNum.EntityData.Children = types.NewOrderedMap()
    controlWriteSuccessNum.EntityData.Leafs = types.NewOrderedMap()
    controlWriteSuccessNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", controlWriteSuccessNum.Entry})

    controlWriteSuccessNum.EntityData.YListKeys = []string {}

    return &(controlWriteSuccessNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadSuccessNum
// Number of successful control read to XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadSuccessNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (controlReadSuccessNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadSuccessNum) GetEntityData() *types.CommonEntityData {
    controlReadSuccessNum.EntityData.YFilter = controlReadSuccessNum.YFilter
    controlReadSuccessNum.EntityData.YangName = "control-read-success-num"
    controlReadSuccessNum.EntityData.BundleName = "cisco_ios_xr"
    controlReadSuccessNum.EntityData.ParentYangName = "async-session-stats"
    controlReadSuccessNum.EntityData.SegmentPath = "control-read-success-num" + types.AddNoKeyToken(controlReadSuccessNum)
    controlReadSuccessNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + controlReadSuccessNum.EntityData.SegmentPath
    controlReadSuccessNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlReadSuccessNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlReadSuccessNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlReadSuccessNum.EntityData.Children = types.NewOrderedMap()
    controlReadSuccessNum.EntityData.Leafs = types.NewOrderedMap()
    controlReadSuccessNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", controlReadSuccessNum.Entry})

    controlReadSuccessNum.EntityData.YListKeys = []string {}

    return &(controlReadSuccessNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteErrorNum
// Number of failed control write to XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteErrorNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (controlWriteErrorNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlWriteErrorNum) GetEntityData() *types.CommonEntityData {
    controlWriteErrorNum.EntityData.YFilter = controlWriteErrorNum.YFilter
    controlWriteErrorNum.EntityData.YangName = "control-write-error-num"
    controlWriteErrorNum.EntityData.BundleName = "cisco_ios_xr"
    controlWriteErrorNum.EntityData.ParentYangName = "async-session-stats"
    controlWriteErrorNum.EntityData.SegmentPath = "control-write-error-num" + types.AddNoKeyToken(controlWriteErrorNum)
    controlWriteErrorNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + controlWriteErrorNum.EntityData.SegmentPath
    controlWriteErrorNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlWriteErrorNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlWriteErrorNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlWriteErrorNum.EntityData.Children = types.NewOrderedMap()
    controlWriteErrorNum.EntityData.Leafs = types.NewOrderedMap()
    controlWriteErrorNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", controlWriteErrorNum.Entry})

    controlWriteErrorNum.EntityData.YListKeys = []string {}

    return &(controlWriteErrorNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadErrorNum
// Number of failed control read from XIPC
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadErrorNum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (controlReadErrorNum *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_ControlReadErrorNum) GetEntityData() *types.CommonEntityData {
    controlReadErrorNum.EntityData.YFilter = controlReadErrorNum.YFilter
    controlReadErrorNum.EntityData.YangName = "control-read-error-num"
    controlReadErrorNum.EntityData.BundleName = "cisco_ios_xr"
    controlReadErrorNum.EntityData.ParentYangName = "async-session-stats"
    controlReadErrorNum.EntityData.SegmentPath = "control-read-error-num" + types.AddNoKeyToken(controlReadErrorNum)
    controlReadErrorNum.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + controlReadErrorNum.EntityData.SegmentPath
    controlReadErrorNum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlReadErrorNum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlReadErrorNum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlReadErrorNum.EntityData.Children = types.NewOrderedMap()
    controlReadErrorNum.EntityData.Leafs = types.NewOrderedMap()
    controlReadErrorNum.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", controlReadErrorNum.Entry})

    controlReadErrorNum.EntityData.YListKeys = []string {}

    return &(controlReadErrorNum.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteByte
// Number of bytes data has been written
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteByte struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..18446744073709551615. Units are
    // byte.
    Entry interface{}
}

func (dataWriteByte *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataWriteByte) GetEntityData() *types.CommonEntityData {
    dataWriteByte.EntityData.YFilter = dataWriteByte.YFilter
    dataWriteByte.EntityData.YangName = "data-write-byte"
    dataWriteByte.EntityData.BundleName = "cisco_ios_xr"
    dataWriteByte.EntityData.ParentYangName = "async-session-stats"
    dataWriteByte.EntityData.SegmentPath = "data-write-byte" + types.AddNoKeyToken(dataWriteByte)
    dataWriteByte.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + dataWriteByte.EntityData.SegmentPath
    dataWriteByte.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataWriteByte.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataWriteByte.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataWriteByte.EntityData.Children = types.NewOrderedMap()
    dataWriteByte.EntityData.Leafs = types.NewOrderedMap()
    dataWriteByte.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dataWriteByte.Entry})

    dataWriteByte.EntityData.YListKeys = []string {}

    return &(dataWriteByte.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadByte
// Number of bytes data has been read
type TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadByte struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..18446744073709551615. Units are
    // byte.
    Entry interface{}
}

func (dataReadByte *TcpConnection_Nodes_Node_Statistics_Pcbs_Pcb_AsyncSessionStats_DataReadByte) GetEntityData() *types.CommonEntityData {
    dataReadByte.EntityData.YFilter = dataReadByte.YFilter
    dataReadByte.EntityData.YangName = "data-read-byte"
    dataReadByte.EntityData.BundleName = "cisco_ios_xr"
    dataReadByte.EntityData.ParentYangName = "async-session-stats"
    dataReadByte.EntityData.SegmentPath = "data-read-byte" + types.AddNoKeyToken(dataReadByte)
    dataReadByte.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/pcbs/pcb/async-session-stats/" + dataReadByte.EntityData.SegmentPath
    dataReadByte.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dataReadByte.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dataReadByte.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dataReadByte.EntityData.Children = types.NewOrderedMap()
    dataReadByte.EntityData.Leafs = types.NewOrderedMap()
    dataReadByte.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", dataReadByte.Entry})

    dataReadByte.EntityData.YListKeys = []string {}

    return &(dataReadByte.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Summary
// Summary statistics across all TCP connections
type TcpConnection_Nodes_Node_Statistics_Summary struct {
    EntityData types.CommonEntityData
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

    // Current number of packets extended for scaled sockets. The type is
    // interface{} with range: 0..4294967295.
    SockbufPakResCur interface{}

    // Maximum number of packets extended for scaled sockets. The type is
    // interface{} with range: 0..4294967295.
    SockbufPakResMax interface{}

    // Total Number of Ingress packets on TCP iqs. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalIngpacket.
    IqsTotalIngpacket []*TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalIngpacket

    // Total Number of Egress packets on TCP iqs. The type is slice of
    // TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalEgpacket.
    IqsTotalEgpacket []*TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalEgpacket
}

func (summary *TcpConnection_Nodes_Node_Statistics_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "statistics"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("iqs-total-ingpacket", types.YChild{"IqsTotalIngpacket", nil})
    for i := range summary.IqsTotalIngpacket {
        types.SetYListKey(summary.IqsTotalIngpacket[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.IqsTotalIngpacket[i]), types.YChild{"IqsTotalIngpacket", summary.IqsTotalIngpacket[i]})
    }
    summary.EntityData.Children.Append("iqs-total-egpacket", types.YChild{"IqsTotalEgpacket", nil})
    for i := range summary.IqsTotalEgpacket {
        types.SetYListKey(summary.IqsTotalEgpacket[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.IqsTotalEgpacket[i]), types.YChild{"IqsTotalEgpacket", summary.IqsTotalEgpacket[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("syn-cache-count", types.YLeaf{"SynCacheCount", summary.SynCacheCount})
    summary.EntityData.Leafs.Append("num-open-sockets", types.YLeaf{"NumOpenSockets", summary.NumOpenSockets})
    summary.EntityData.Leafs.Append("total-pakets-sent", types.YLeaf{"TotalPaketsSent", summary.TotalPaketsSent})
    summary.EntityData.Leafs.Append("send-packets-dropped", types.YLeaf{"SendPacketsDropped", summary.SendPacketsDropped})
    summary.EntityData.Leafs.Append("send-auth-packets-dropped", types.YLeaf{"SendAuthPacketsDropped", summary.SendAuthPacketsDropped})
    summary.EntityData.Leafs.Append("data-pakets-sent", types.YLeaf{"DataPaketsSent", summary.DataPaketsSent})
    summary.EntityData.Leafs.Append("data-bytes-sent", types.YLeaf{"DataBytesSent", summary.DataBytesSent})
    summary.EntityData.Leafs.Append("packets-retransmitted", types.YLeaf{"PacketsRetransmitted", summary.PacketsRetransmitted})
    summary.EntityData.Leafs.Append("bytes-retransmitted", types.YLeaf{"BytesRetransmitted", summary.BytesRetransmitted})
    summary.EntityData.Leafs.Append("ack-only-packets-sent", types.YLeaf{"AckOnlyPacketsSent", summary.AckOnlyPacketsSent})
    summary.EntityData.Leafs.Append("delay-ack-packets-sent", types.YLeaf{"DelayAckPacketsSent", summary.DelayAckPacketsSent})
    summary.EntityData.Leafs.Append("urgent-only-packets-sent", types.YLeaf{"UrgentOnlyPacketsSent", summary.UrgentOnlyPacketsSent})
    summary.EntityData.Leafs.Append("window-probe-packets-sent", types.YLeaf{"WindowProbePacketsSent", summary.WindowProbePacketsSent})
    summary.EntityData.Leafs.Append("window-update-packets-sent", types.YLeaf{"WindowUpdatePacketsSent", summary.WindowUpdatePacketsSent})
    summary.EntityData.Leafs.Append("control-packets-sent", types.YLeaf{"ControlPacketsSent", summary.ControlPacketsSent})
    summary.EntityData.Leafs.Append("rst-packets-sent", types.YLeaf{"RstPacketsSent", summary.RstPacketsSent})
    summary.EntityData.Leafs.Append("total-packets-received", types.YLeaf{"TotalPacketsReceived", summary.TotalPacketsReceived})
    summary.EntityData.Leafs.Append("received-packets-dropped", types.YLeaf{"ReceivedPacketsDropped", summary.ReceivedPacketsDropped})
    summary.EntityData.Leafs.Append("synacl-match-pkts-dropped", types.YLeaf{"SynaclMatchPktsDropped", summary.SynaclMatchPktsDropped})
    summary.EntityData.Leafs.Append("received-packets-dropped-stale-c-hdr", types.YLeaf{"ReceivedPacketsDroppedStaleCHdr", summary.ReceivedPacketsDroppedStaleCHdr})
    summary.EntityData.Leafs.Append("received-auth-packets-dropped", types.YLeaf{"ReceivedAuthPacketsDropped", summary.ReceivedAuthPacketsDropped})
    summary.EntityData.Leafs.Append("ack-packets-received", types.YLeaf{"AckPacketsReceived", summary.AckPacketsReceived})
    summary.EntityData.Leafs.Append("ackbytes-received", types.YLeaf{"AckbytesReceived", summary.AckbytesReceived})
    summary.EntityData.Leafs.Append("duplicated-ack-packets-received", types.YLeaf{"DuplicatedAckPacketsReceived", summary.DuplicatedAckPacketsReceived})
    summary.EntityData.Leafs.Append("ack-packets-for-unsent-received", types.YLeaf{"AckPacketsForUnsentReceived", summary.AckPacketsForUnsentReceived})
    summary.EntityData.Leafs.Append("data-packets-received-in-sequence", types.YLeaf{"DataPacketsReceivedInSequence", summary.DataPacketsReceivedInSequence})
    summary.EntityData.Leafs.Append("data-bytes-received-in-sequence", types.YLeaf{"DataBytesReceivedInSequence", summary.DataBytesReceivedInSequence})
    summary.EntityData.Leafs.Append("duplicate-packets-received", types.YLeaf{"DuplicatePacketsReceived", summary.DuplicatePacketsReceived})
    summary.EntityData.Leafs.Append("duplicate-bytes-received", types.YLeaf{"DuplicateBytesReceived", summary.DuplicateBytesReceived})
    summary.EntityData.Leafs.Append("partial-duplicate-ack-received", types.YLeaf{"PartialDuplicateAckReceived", summary.PartialDuplicateAckReceived})
    summary.EntityData.Leafs.Append("partial-duplicate-bytes-received", types.YLeaf{"PartialDuplicateBytesReceived", summary.PartialDuplicateBytesReceived})
    summary.EntityData.Leafs.Append("out-of-order-packets-received", types.YLeaf{"OutOfOrderPacketsReceived", summary.OutOfOrderPacketsReceived})
    summary.EntityData.Leafs.Append("out-of-order-bytes-received", types.YLeaf{"OutOfOrderBytesReceived", summary.OutOfOrderBytesReceived})
    summary.EntityData.Leafs.Append("after-window-packets-received", types.YLeaf{"AfterWindowPacketsReceived", summary.AfterWindowPacketsReceived})
    summary.EntityData.Leafs.Append("after-window-bytes-received", types.YLeaf{"AfterWindowBytesReceived", summary.AfterWindowBytesReceived})
    summary.EntityData.Leafs.Append("window-probe-packets-received", types.YLeaf{"WindowProbePacketsReceived", summary.WindowProbePacketsReceived})
    summary.EntityData.Leafs.Append("window-update-packets-received", types.YLeaf{"WindowUpdatePacketsReceived", summary.WindowUpdatePacketsReceived})
    summary.EntityData.Leafs.Append("packets-received-after-close-packet", types.YLeaf{"PacketsReceivedAfterClosePacket", summary.PacketsReceivedAfterClosePacket})
    summary.EntityData.Leafs.Append("bad-checksum-packets-received", types.YLeaf{"BadChecksumPacketsReceived", summary.BadChecksumPacketsReceived})
    summary.EntityData.Leafs.Append("too-short-packets-received", types.YLeaf{"TooShortPacketsReceived", summary.TooShortPacketsReceived})
    summary.EntityData.Leafs.Append("malformed-packets-received", types.YLeaf{"MalformedPacketsReceived", summary.MalformedPacketsReceived})
    summary.EntityData.Leafs.Append("no-port-packets-received", types.YLeaf{"NoPortPacketsReceived", summary.NoPortPacketsReceived})
    summary.EntityData.Leafs.Append("connections-requested", types.YLeaf{"ConnectionsRequested", summary.ConnectionsRequested})
    summary.EntityData.Leafs.Append("connections-accepted", types.YLeaf{"ConnectionsAccepted", summary.ConnectionsAccepted})
    summary.EntityData.Leafs.Append("connections-established", types.YLeaf{"ConnectionsEstablished", summary.ConnectionsEstablished})
    summary.EntityData.Leafs.Append("connections-forcibly-closed", types.YLeaf{"ConnectionsForciblyClosed", summary.ConnectionsForciblyClosed})
    summary.EntityData.Leafs.Append("connections-closed", types.YLeaf{"ConnectionsClosed", summary.ConnectionsClosed})
    summary.EntityData.Leafs.Append("connections-dropped", types.YLeaf{"ConnectionsDropped", summary.ConnectionsDropped})
    summary.EntityData.Leafs.Append("embryonic-connection-dropped", types.YLeaf{"EmbryonicConnectionDropped", summary.EmbryonicConnectionDropped})
    summary.EntityData.Leafs.Append("connections-failed", types.YLeaf{"ConnectionsFailed", summary.ConnectionsFailed})
    summary.EntityData.Leafs.Append("established-connections-reset", types.YLeaf{"EstablishedConnectionsReset", summary.EstablishedConnectionsReset})
    summary.EntityData.Leafs.Append("retransmit-timeouts", types.YLeaf{"RetransmitTimeouts", summary.RetransmitTimeouts})
    summary.EntityData.Leafs.Append("retransmit-dropped", types.YLeaf{"RetransmitDropped", summary.RetransmitDropped})
    summary.EntityData.Leafs.Append("keep-alive-timeouts", types.YLeaf{"KeepAliveTimeouts", summary.KeepAliveTimeouts})
    summary.EntityData.Leafs.Append("keep-alive-dropped", types.YLeaf{"KeepAliveDropped", summary.KeepAliveDropped})
    summary.EntityData.Leafs.Append("keep-alive-probes", types.YLeaf{"KeepAliveProbes", summary.KeepAliveProbes})
    summary.EntityData.Leafs.Append("paws-dropped", types.YLeaf{"PawsDropped", summary.PawsDropped})
    summary.EntityData.Leafs.Append("persist-dropped", types.YLeaf{"PersistDropped", summary.PersistDropped})
    summary.EntityData.Leafs.Append("try-lock-dropped", types.YLeaf{"TryLockDropped", summary.TryLockDropped})
    summary.EntityData.Leafs.Append("connection-rate-limited", types.YLeaf{"ConnectionRateLimited", summary.ConnectionRateLimited})
    summary.EntityData.Leafs.Append("syn-cache-added", types.YLeaf{"SynCacheAdded", summary.SynCacheAdded})
    summary.EntityData.Leafs.Append("syn-cache-completed", types.YLeaf{"SynCacheCompleted", summary.SynCacheCompleted})
    summary.EntityData.Leafs.Append("syn-cache-timed-out", types.YLeaf{"SynCacheTimedOut", summary.SynCacheTimedOut})
    summary.EntityData.Leafs.Append("syn-cache-overflow", types.YLeaf{"SynCacheOverflow", summary.SynCacheOverflow})
    summary.EntityData.Leafs.Append("syn-cache-reset", types.YLeaf{"SynCacheReset", summary.SynCacheReset})
    summary.EntityData.Leafs.Append("syn-cache-unreach", types.YLeaf{"SynCacheUnreach", summary.SynCacheUnreach})
    summary.EntityData.Leafs.Append("syn-cache-bucket-oflow", types.YLeaf{"SynCacheBucketOflow", summary.SynCacheBucketOflow})
    summary.EntityData.Leafs.Append("syn-cache-aborted", types.YLeaf{"SynCacheAborted", summary.SynCacheAborted})
    summary.EntityData.Leafs.Append("syn-cache-duplicate-sy-ns", types.YLeaf{"SynCacheDuplicateSyNs", summary.SynCacheDuplicateSyNs})
    summary.EntityData.Leafs.Append("syn-cache-dropped", types.YLeaf{"SynCacheDropped", summary.SynCacheDropped})
    summary.EntityData.Leafs.Append("pulse-errors", types.YLeaf{"PulseErrors", summary.PulseErrors})
    summary.EntityData.Leafs.Append("socket-layer-packets", types.YLeaf{"SocketLayerPackets", summary.SocketLayerPackets})
    summary.EntityData.Leafs.Append("reassembly-packets", types.YLeaf{"ReassemblyPackets", summary.ReassemblyPackets})
    summary.EntityData.Leafs.Append("recovered-packets", types.YLeaf{"RecoveredPackets", summary.RecoveredPackets})
    summary.EntityData.Leafs.Append("packet-failures", types.YLeaf{"PacketFailures", summary.PacketFailures})
    summary.EntityData.Leafs.Append("mss-up", types.YLeaf{"MssUp", summary.MssUp})
    summary.EntityData.Leafs.Append("mss-down", types.YLeaf{"MssDown", summary.MssDown})
    summary.EntityData.Leafs.Append("truncated-write-iov", types.YLeaf{"TruncatedWriteIov", summary.TruncatedWriteIov})
    summary.EntityData.Leafs.Append("no-throttle", types.YLeaf{"NoThrottle", summary.NoThrottle})
    summary.EntityData.Leafs.Append("low-water-mark-throttle", types.YLeaf{"LowWaterMarkThrottle", summary.LowWaterMarkThrottle})
    summary.EntityData.Leafs.Append("high-water-mark-throttle", types.YLeaf{"HighWaterMarkThrottle", summary.HighWaterMarkThrottle})
    summary.EntityData.Leafs.Append("stalled-timer-tickle-count", types.YLeaf{"StalledTimerTickleCount", summary.StalledTimerTickleCount})
    summary.EntityData.Leafs.Append("stalled-timer-tickle-time", types.YLeaf{"StalledTimerTickleTime", summary.StalledTimerTickleTime})
    summary.EntityData.Leafs.Append("iq-sock-writes", types.YLeaf{"IqSockWrites", summary.IqSockWrites})
    summary.EntityData.Leafs.Append("iq-sock-retries", types.YLeaf{"IqSockRetries", summary.IqSockRetries})
    summary.EntityData.Leafs.Append("iq-sock-aborts", types.YLeaf{"IqSockAborts", summary.IqSockAborts})
    summary.EntityData.Leafs.Append("iq-ingress-drops", types.YLeaf{"IqIngressDrops", summary.IqIngressDrops})
    summary.EntityData.Leafs.Append("total-i-qs", types.YLeaf{"TotalIQs", summary.TotalIQs})
    summary.EntityData.Leafs.Append("sockbuf-pak-res-cur", types.YLeaf{"SockbufPakResCur", summary.SockbufPakResCur})
    summary.EntityData.Leafs.Append("sockbuf-pak-res-max", types.YLeaf{"SockbufPakResMax", summary.SockbufPakResMax})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalIngpacket
// Total Number of Ingress packets on TCP iqs
type TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalIngpacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (iqsTotalIngpacket *TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalIngpacket) GetEntityData() *types.CommonEntityData {
    iqsTotalIngpacket.EntityData.YFilter = iqsTotalIngpacket.YFilter
    iqsTotalIngpacket.EntityData.YangName = "iqs-total-ingpacket"
    iqsTotalIngpacket.EntityData.BundleName = "cisco_ios_xr"
    iqsTotalIngpacket.EntityData.ParentYangName = "summary"
    iqsTotalIngpacket.EntityData.SegmentPath = "iqs-total-ingpacket" + types.AddNoKeyToken(iqsTotalIngpacket)
    iqsTotalIngpacket.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/summary/" + iqsTotalIngpacket.EntityData.SegmentPath
    iqsTotalIngpacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iqsTotalIngpacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iqsTotalIngpacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iqsTotalIngpacket.EntityData.Children = types.NewOrderedMap()
    iqsTotalIngpacket.EntityData.Leafs = types.NewOrderedMap()
    iqsTotalIngpacket.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", iqsTotalIngpacket.Entry})

    iqsTotalIngpacket.EntityData.YListKeys = []string {}

    return &(iqsTotalIngpacket.EntityData)
}

// TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalEgpacket
// Total Number of Egress packets on TCP iqs
type TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalEgpacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (iqsTotalEgpacket *TcpConnection_Nodes_Node_Statistics_Summary_IqsTotalEgpacket) GetEntityData() *types.CommonEntityData {
    iqsTotalEgpacket.EntityData.YFilter = iqsTotalEgpacket.YFilter
    iqsTotalEgpacket.EntityData.YangName = "iqs-total-egpacket"
    iqsTotalEgpacket.EntityData.BundleName = "cisco_ios_xr"
    iqsTotalEgpacket.EntityData.ParentYangName = "summary"
    iqsTotalEgpacket.EntityData.SegmentPath = "iqs-total-egpacket" + types.AddNoKeyToken(iqsTotalEgpacket)
    iqsTotalEgpacket.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/statistics/summary/" + iqsTotalEgpacket.EntityData.SegmentPath
    iqsTotalEgpacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iqsTotalEgpacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iqsTotalEgpacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iqsTotalEgpacket.EntityData.Children = types.NewOrderedMap()
    iqsTotalEgpacket.EntityData.Leafs = types.NewOrderedMap()
    iqsTotalEgpacket.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", iqsTotalEgpacket.Entry})

    iqsTotalEgpacket.EntityData.YListKeys = []string {}

    return &(iqsTotalEgpacket.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation
// Extended Filter related Information
type TcpConnection_Nodes_Node_ExtendedInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table listing display types.
    DisplayTypes TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes
}

func (extendedInformation *TcpConnection_Nodes_Node_ExtendedInformation) GetEntityData() *types.CommonEntityData {
    extendedInformation.EntityData.YFilter = extendedInformation.YFilter
    extendedInformation.EntityData.YangName = "extended-information"
    extendedInformation.EntityData.BundleName = "cisco_ios_xr"
    extendedInformation.EntityData.ParentYangName = "node"
    extendedInformation.EntityData.SegmentPath = "extended-information"
    extendedInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/" + extendedInformation.EntityData.SegmentPath
    extendedInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    extendedInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    extendedInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    extendedInformation.EntityData.Children = types.NewOrderedMap()
    extendedInformation.EntityData.Children.Append("display-types", types.YChild{"DisplayTypes", &extendedInformation.DisplayTypes})
    extendedInformation.EntityData.Leafs = types.NewOrderedMap()

    extendedInformation.EntityData.YListKeys = []string {}

    return &(extendedInformation.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes
// Table listing display types
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Describing particular display type. The type is slice of
    // TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType.
    DisplayType []*TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType
}

func (displayTypes *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes) GetEntityData() *types.CommonEntityData {
    displayTypes.EntityData.YFilter = displayTypes.YFilter
    displayTypes.EntityData.YangName = "display-types"
    displayTypes.EntityData.BundleName = "cisco_ios_xr"
    displayTypes.EntityData.ParentYangName = "extended-information"
    displayTypes.EntityData.SegmentPath = "display-types"
    displayTypes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/" + displayTypes.EntityData.SegmentPath
    displayTypes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayTypes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayTypes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayTypes.EntityData.Children = types.NewOrderedMap()
    displayTypes.EntityData.Children.Append("display-type", types.YChild{"DisplayType", nil})
    for i := range displayTypes.DisplayType {
        displayTypes.EntityData.Children.Append(types.GetSegmentPath(displayTypes.DisplayType[i]), types.YChild{"DisplayType", displayTypes.DisplayType[i]})
    }
    displayTypes.EntityData.Leafs = types.NewOrderedMap()

    displayTypes.EntityData.YListKeys = []string {}

    return &(displayTypes.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType
// Describing particular display type
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Specifying display type. The type is Show.
    DispType interface{}

    // Describing connection ID. The type is slice of
    // TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId.
    ConnectionId []*TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId
}

func (displayType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType) GetEntityData() *types.CommonEntityData {
    displayType.EntityData.YFilter = displayType.YFilter
    displayType.EntityData.YangName = "display-type"
    displayType.EntityData.BundleName = "cisco_ios_xr"
    displayType.EntityData.ParentYangName = "display-types"
    displayType.EntityData.SegmentPath = "display-type" + types.AddKeyToken(displayType.DispType, "disp-type")
    displayType.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/" + displayType.EntityData.SegmentPath
    displayType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    displayType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    displayType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    displayType.EntityData.Children = types.NewOrderedMap()
    displayType.EntityData.Children.Append("connection-id", types.YChild{"ConnectionId", nil})
    for i := range displayType.ConnectionId {
        displayType.EntityData.Children.Append(types.GetSegmentPath(displayType.ConnectionId[i]), types.YChild{"ConnectionId", displayType.ConnectionId[i]})
    }
    displayType.EntityData.Leafs = types.NewOrderedMap()
    displayType.EntityData.Leafs.Append("disp-type", types.YLeaf{"DispType", displayType.DispType})

    displayType.EntityData.YListKeys = []string {"DispType"}

    return &(displayType.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId
// Describing connection ID
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Displaying inforamtion based on selected display
    // type associatedwith a particular PCB. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
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

func (connectionId *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId) GetEntityData() *types.CommonEntityData {
    connectionId.EntityData.YFilter = connectionId.YFilter
    connectionId.EntityData.YangName = "connection-id"
    connectionId.EntityData.BundleName = "cisco_ios_xr"
    connectionId.EntityData.ParentYangName = "display-type"
    connectionId.EntityData.SegmentPath = "connection-id" + types.AddKeyToken(connectionId.PcbId, "pcb-id")
    connectionId.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/" + connectionId.EntityData.SegmentPath
    connectionId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectionId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectionId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectionId.EntityData.Children = types.NewOrderedMap()
    connectionId.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &connectionId.LocalAddress})
    connectionId.EntityData.Children.Append("foreign-address", types.YChild{"ForeignAddress", &connectionId.ForeignAddress})
    connectionId.EntityData.Children.Append("common", types.YChild{"Common", &connectionId.Common})
    connectionId.EntityData.Leafs = types.NewOrderedMap()
    connectionId.EntityData.Leafs.Append("pcb-id", types.YLeaf{"PcbId", connectionId.PcbId})
    connectionId.EntityData.Leafs.Append("l4-protocol", types.YLeaf{"L4Protocol", connectionId.L4Protocol})
    connectionId.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", connectionId.LocalPort})
    connectionId.EntityData.Leafs.Append("foreign-port", types.YLeaf{"ForeignPort", connectionId.ForeignPort})

    connectionId.EntityData.YListKeys = []string {"PcbId"}

    return &(connectionId.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress
// Local IP address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress struct {
    EntityData types.CommonEntityData
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

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "connection-id"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddress.AfName})
    localAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", localAddress.Ipv4Address})
    localAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", localAddress.Ipv6Address})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress
// Remote IP address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress struct {
    EntityData types.CommonEntityData
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

func (foreignAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "connection-id"
    foreignAddress.EntityData.SegmentPath = "foreign-address"
    foreignAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/" + foreignAddress.EntityData.SegmentPath
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", foreignAddress.AfName})
    foreignAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", foreignAddress.Ipv4Address})
    foreignAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", foreignAddress.Ipv6Address})

    foreignAddress.EntityData.YListKeys = []string {}

    return &(foreignAddress.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common
// Common PCB information
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address Family. The type is AddrFamily.
    AfName interface{}

    // LPTS PCB information.
    LptsPcb TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb
}

func (common *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "connection-id"
    common.EntityData.SegmentPath = "common"
    common.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/" + common.EntityData.SegmentPath
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = types.NewOrderedMap()
    common.EntityData.Children.Append("lpts-pcb", types.YChild{"LptsPcb", &common.LptsPcb})
    common.EntityData.Leafs = types.NewOrderedMap()
    common.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", common.AfName})

    common.EntityData.YListKeys = []string {}

    return &(common.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb
// LPTS PCB information
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb struct {
    EntityData types.CommonEntityData
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
    Filter []*TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter
}

func (lptsPcb *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb) GetEntityData() *types.CommonEntityData {
    lptsPcb.EntityData.YFilter = lptsPcb.YFilter
    lptsPcb.EntityData.YangName = "lpts-pcb"
    lptsPcb.EntityData.BundleName = "cisco_ios_xr"
    lptsPcb.EntityData.ParentYangName = "common"
    lptsPcb.EntityData.SegmentPath = "lpts-pcb"
    lptsPcb.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/" + lptsPcb.EntityData.SegmentPath
    lptsPcb.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsPcb.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsPcb.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsPcb.EntityData.Children = types.NewOrderedMap()
    lptsPcb.EntityData.Children.Append("options", types.YChild{"Options", &lptsPcb.Options})
    lptsPcb.EntityData.Children.Append("lpts-flags", types.YChild{"LptsFlags", &lptsPcb.LptsFlags})
    lptsPcb.EntityData.Children.Append("accept-mask", types.YChild{"AcceptMask", &lptsPcb.AcceptMask})
    lptsPcb.EntityData.Children.Append("filter", types.YChild{"Filter", nil})
    for i := range lptsPcb.Filter {
        types.SetYListKey(lptsPcb.Filter[i], i)
        lptsPcb.EntityData.Children.Append(types.GetSegmentPath(lptsPcb.Filter[i]), types.YChild{"Filter", lptsPcb.Filter[i]})
    }
    lptsPcb.EntityData.Leafs = types.NewOrderedMap()
    lptsPcb.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", lptsPcb.Ttl})
    lptsPcb.EntityData.Leafs.Append("flow-types-info", types.YLeaf{"FlowTypesInfo", lptsPcb.FlowTypesInfo})

    lptsPcb.EntityData.YListKeys = []string {}

    return &(lptsPcb.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options
// Receive options
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Receive filter enabled. The type is bool.
    IsReceiveFilter interface{}

    // IP SLA. The type is bool.
    IsIpSla interface{}
}

func (options *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "lpts-pcb"
    options.EntityData.SegmentPath = "options"
    options.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/" + options.EntityData.SegmentPath
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = types.NewOrderedMap()
    options.EntityData.Leafs = types.NewOrderedMap()
    options.EntityData.Leafs.Append("is-receive-filter", types.YLeaf{"IsReceiveFilter", options.IsReceiveFilter})
    options.EntityData.Leafs.Append("is-ip-sla", types.YLeaf{"IsIpSla", options.IsIpSla})

    options.EntityData.YListKeys = []string {}

    return &(options.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags
// LPTS flags
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PCB bound. The type is bool.
    IsPcbBound interface{}

    // Sent drop packets. The type is bool.
    IsLocalAddressIgnore interface{}

    // Ignore VRF Filter. The type is bool.
    IsIgnoreVrfFilter interface{}
}

func (lptsFlags *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_LptsFlags) GetEntityData() *types.CommonEntityData {
    lptsFlags.EntityData.YFilter = lptsFlags.YFilter
    lptsFlags.EntityData.YangName = "lpts-flags"
    lptsFlags.EntityData.BundleName = "cisco_ios_xr"
    lptsFlags.EntityData.ParentYangName = "lpts-pcb"
    lptsFlags.EntityData.SegmentPath = "lpts-flags"
    lptsFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/" + lptsFlags.EntityData.SegmentPath
    lptsFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lptsFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lptsFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lptsFlags.EntityData.Children = types.NewOrderedMap()
    lptsFlags.EntityData.Leafs = types.NewOrderedMap()
    lptsFlags.EntityData.Leafs.Append("is-pcb-bound", types.YLeaf{"IsPcbBound", lptsFlags.IsPcbBound})
    lptsFlags.EntityData.Leafs.Append("is-local-address-ignore", types.YLeaf{"IsLocalAddressIgnore", lptsFlags.IsLocalAddressIgnore})
    lptsFlags.EntityData.Leafs.Append("is-ignore-vrf-filter", types.YLeaf{"IsIgnoreVrfFilter", lptsFlags.IsIgnoreVrfFilter})

    lptsFlags.EntityData.YListKeys = []string {}

    return &(lptsFlags.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask
// AcceptMask
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask struct {
    EntityData types.CommonEntityData
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

func (acceptMask *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_AcceptMask) GetEntityData() *types.CommonEntityData {
    acceptMask.EntityData.YFilter = acceptMask.YFilter
    acceptMask.EntityData.YangName = "accept-mask"
    acceptMask.EntityData.BundleName = "cisco_ios_xr"
    acceptMask.EntityData.ParentYangName = "lpts-pcb"
    acceptMask.EntityData.SegmentPath = "accept-mask"
    acceptMask.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/" + acceptMask.EntityData.SegmentPath
    acceptMask.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    acceptMask.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    acceptMask.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    acceptMask.EntityData.Children = types.NewOrderedMap()
    acceptMask.EntityData.Leafs = types.NewOrderedMap()
    acceptMask.EntityData.Leafs.Append("is-interface", types.YLeaf{"IsInterface", acceptMask.IsInterface})
    acceptMask.EntityData.Leafs.Append("is-packet-type", types.YLeaf{"IsPacketType", acceptMask.IsPacketType})
    acceptMask.EntityData.Leafs.Append("is-remote-address", types.YLeaf{"IsRemoteAddress", acceptMask.IsRemoteAddress})
    acceptMask.EntityData.Leafs.Append("is-remote-port", types.YLeaf{"IsRemotePort", acceptMask.IsRemotePort})
    acceptMask.EntityData.Leafs.Append("is-local-address", types.YLeaf{"IsLocalAddress", acceptMask.IsLocalAddress})
    acceptMask.EntityData.Leafs.Append("is-local-port", types.YLeaf{"IsLocalPort", acceptMask.IsLocalPort})

    acceptMask.EntityData.YListKeys = []string {}

    return &(acceptMask.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter
// Interface Filters
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Interface name. The type is string with pattern: [a-zA-Z0-9._/-]+.
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

func (filter *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter) GetEntityData() *types.CommonEntityData {
    filter.EntityData.YFilter = filter.YFilter
    filter.EntityData.YangName = "filter"
    filter.EntityData.BundleName = "cisco_ios_xr"
    filter.EntityData.ParentYangName = "lpts-pcb"
    filter.EntityData.SegmentPath = "filter" + types.AddNoKeyToken(filter)
    filter.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/" + filter.EntityData.SegmentPath
    filter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    filter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    filter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    filter.EntityData.Children = types.NewOrderedMap()
    filter.EntityData.Children.Append("packet-type", types.YChild{"PacketType", &filter.PacketType})
    filter.EntityData.Children.Append("remote-address", types.YChild{"RemoteAddress", &filter.RemoteAddress})
    filter.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &filter.LocalAddress})
    filter.EntityData.Leafs = types.NewOrderedMap()
    filter.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", filter.InterfaceName})
    filter.EntityData.Leafs.Append("remote-length", types.YLeaf{"RemoteLength", filter.RemoteLength})
    filter.EntityData.Leafs.Append("local-length", types.YLeaf{"LocalLength", filter.LocalLength})
    filter.EntityData.Leafs.Append("receive-remote-port", types.YLeaf{"ReceiveRemotePort", filter.ReceiveRemotePort})
    filter.EntityData.Leafs.Append("receive-local-port", types.YLeaf{"ReceiveLocalPort", filter.ReceiveLocalPort})
    filter.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", filter.Priority})
    filter.EntityData.Leafs.Append("ttl", types.YLeaf{"Ttl", filter.Ttl})
    filter.EntityData.Leafs.Append("flow-types-info", types.YLeaf{"FlowTypesInfo", filter.FlowTypesInfo})

    filter.EntityData.YListKeys = []string {}

    return &(filter.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType
// Protocol-specific packet type
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type. The type is Packet.
    Type interface{}

    // ICMP message type. The type is MessageTypeIcmp_.
    IcmpMessageType interface{}

    // ICMPv6 message type. The type is MessageTypeIcmpv6_.
    IcmPv6MessageType interface{}

    // IGMP message type. The type is MessageTypeIgmp_.
    IgmpMessageType interface{}

    // Message type in number. The type is interface{} with range: 0..4294967295.
    MessageId interface{}
}

func (packetType *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_PacketType) GetEntityData() *types.CommonEntityData {
    packetType.EntityData.YFilter = packetType.YFilter
    packetType.EntityData.YangName = "packet-type"
    packetType.EntityData.BundleName = "cisco_ios_xr"
    packetType.EntityData.ParentYangName = "filter"
    packetType.EntityData.SegmentPath = "packet-type"
    packetType.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/filter/" + packetType.EntityData.SegmentPath
    packetType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetType.EntityData.Children = types.NewOrderedMap()
    packetType.EntityData.Leafs = types.NewOrderedMap()
    packetType.EntityData.Leafs.Append("type", types.YLeaf{"Type", packetType.Type})
    packetType.EntityData.Leafs.Append("icmp-message-type", types.YLeaf{"IcmpMessageType", packetType.IcmpMessageType})
    packetType.EntityData.Leafs.Append("icm-pv6-message-type", types.YLeaf{"IcmPv6MessageType", packetType.IcmPv6MessageType})
    packetType.EntityData.Leafs.Append("igmp-message-type", types.YLeaf{"IgmpMessageType", packetType.IgmpMessageType})
    packetType.EntityData.Leafs.Append("message-id", types.YLeaf{"MessageId", packetType.MessageId})

    packetType.EntityData.YListKeys = []string {}

    return &(packetType.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress
// Remote address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress struct {
    EntityData types.CommonEntityData
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

func (remoteAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_RemoteAddress) GetEntityData() *types.CommonEntityData {
    remoteAddress.EntityData.YFilter = remoteAddress.YFilter
    remoteAddress.EntityData.YangName = "remote-address"
    remoteAddress.EntityData.BundleName = "cisco_ios_xr"
    remoteAddress.EntityData.ParentYangName = "filter"
    remoteAddress.EntityData.SegmentPath = "remote-address"
    remoteAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/filter/" + remoteAddress.EntityData.SegmentPath
    remoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteAddress.EntityData.Children = types.NewOrderedMap()
    remoteAddress.EntityData.Leafs = types.NewOrderedMap()
    remoteAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", remoteAddress.AfName})
    remoteAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", remoteAddress.Ipv4Address})
    remoteAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", remoteAddress.Ipv6Address})

    remoteAddress.EntityData.YListKeys = []string {}

    return &(remoteAddress.EntityData)
}

// TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress
// Local address
type TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress struct {
    EntityData types.CommonEntityData
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

func (localAddress *TcpConnection_Nodes_Node_ExtendedInformation_DisplayTypes_DisplayType_ConnectionId_Common_LptsPcb_Filter_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "filter"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/extended-information/display-types/display-type/connection-id/common/lpts-pcb/filter/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddress.AfName})
    localAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", localAddress.Ipv4Address})
    localAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", localAddress.Ipv6Address})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations
// Table listing TCP connections for which
// detailed information is provided
type TcpConnection_Nodes_Node_DetailInformations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Protocol Control Block ID. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation.
    DetailInformation []*TcpConnection_Nodes_Node_DetailInformations_DetailInformation
}

func (detailInformations *TcpConnection_Nodes_Node_DetailInformations) GetEntityData() *types.CommonEntityData {
    detailInformations.EntityData.YFilter = detailInformations.YFilter
    detailInformations.EntityData.YangName = "detail-informations"
    detailInformations.EntityData.BundleName = "cisco_ios_xr"
    detailInformations.EntityData.ParentYangName = "node"
    detailInformations.EntityData.SegmentPath = "detail-informations"
    detailInformations.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/" + detailInformations.EntityData.SegmentPath
    detailInformations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailInformations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailInformations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailInformations.EntityData.Children = types.NewOrderedMap()
    detailInformations.EntityData.Children.Append("detail-information", types.YChild{"DetailInformation", nil})
    for i := range detailInformations.DetailInformation {
        detailInformations.EntityData.Children.Append(types.GetSegmentPath(detailInformations.DetailInformation[i]), types.YChild{"DetailInformation", detailInformations.DetailInformation[i]})
    }
    detailInformations.EntityData.Leafs = types.NewOrderedMap()

    detailInformations.EntityData.YListKeys = []string {}

    return &(detailInformations.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation
// Protocol Control Block ID
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Detail information about TCP connection, put null
    // for all. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

    // Cached fib pd context. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibPdCtx.
    FibPdCtx []*TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibPdCtx

    // Cached Label stack. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibLabelOutput.
    FibLabelOutput []*TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibLabelOutput

    // Timers. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer.
    Timer []*TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer

    // Seq nos. of sack blocks. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk.
    SackBlk []*TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk

    // Sorted list of sack holes. The type is slice of
    // TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole.
    SendSackHole []*TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole
}

func (detailInformation *TcpConnection_Nodes_Node_DetailInformations_DetailInformation) GetEntityData() *types.CommonEntityData {
    detailInformation.EntityData.YFilter = detailInformation.YFilter
    detailInformation.EntityData.YangName = "detail-information"
    detailInformation.EntityData.BundleName = "cisco_ios_xr"
    detailInformation.EntityData.ParentYangName = "detail-informations"
    detailInformation.EntityData.SegmentPath = "detail-information" + types.AddKeyToken(detailInformation.PcbId, "pcb-id")
    detailInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/" + detailInformation.EntityData.SegmentPath
    detailInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailInformation.EntityData.Children = types.NewOrderedMap()
    detailInformation.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &detailInformation.LocalAddress})
    detailInformation.EntityData.Children.Append("foreign-address", types.YChild{"ForeignAddress", &detailInformation.ForeignAddress})
    detailInformation.EntityData.Children.Append("socket-option-flags", types.YChild{"SocketOptionFlags", &detailInformation.SocketOptionFlags})
    detailInformation.EntityData.Children.Append("socket-state-flags", types.YChild{"SocketStateFlags", &detailInformation.SocketStateFlags})
    detailInformation.EntityData.Children.Append("feature-flags", types.YChild{"FeatureFlags", &detailInformation.FeatureFlags})
    detailInformation.EntityData.Children.Append("state-flags", types.YChild{"StateFlags", &detailInformation.StateFlags})
    detailInformation.EntityData.Children.Append("request-flags", types.YChild{"RequestFlags", &detailInformation.RequestFlags})
    detailInformation.EntityData.Children.Append("receive-buf-state-flags", types.YChild{"ReceiveBufStateFlags", &detailInformation.ReceiveBufStateFlags})
    detailInformation.EntityData.Children.Append("send-buf-state-flags", types.YChild{"SendBufStateFlags", &detailInformation.SendBufStateFlags})
    detailInformation.EntityData.Children.Append("fib-pd-ctx", types.YChild{"FibPdCtx", nil})
    for i := range detailInformation.FibPdCtx {
        types.SetYListKey(detailInformation.FibPdCtx[i], i)
        detailInformation.EntityData.Children.Append(types.GetSegmentPath(detailInformation.FibPdCtx[i]), types.YChild{"FibPdCtx", detailInformation.FibPdCtx[i]})
    }
    detailInformation.EntityData.Children.Append("fib-label-output", types.YChild{"FibLabelOutput", nil})
    for i := range detailInformation.FibLabelOutput {
        types.SetYListKey(detailInformation.FibLabelOutput[i], i)
        detailInformation.EntityData.Children.Append(types.GetSegmentPath(detailInformation.FibLabelOutput[i]), types.YChild{"FibLabelOutput", detailInformation.FibLabelOutput[i]})
    }
    detailInformation.EntityData.Children.Append("timer", types.YChild{"Timer", nil})
    for i := range detailInformation.Timer {
        types.SetYListKey(detailInformation.Timer[i], i)
        detailInformation.EntityData.Children.Append(types.GetSegmentPath(detailInformation.Timer[i]), types.YChild{"Timer", detailInformation.Timer[i]})
    }
    detailInformation.EntityData.Children.Append("sack-blk", types.YChild{"SackBlk", nil})
    for i := range detailInformation.SackBlk {
        types.SetYListKey(detailInformation.SackBlk[i], i)
        detailInformation.EntityData.Children.Append(types.GetSegmentPath(detailInformation.SackBlk[i]), types.YChild{"SackBlk", detailInformation.SackBlk[i]})
    }
    detailInformation.EntityData.Children.Append("send-sack-hole", types.YChild{"SendSackHole", nil})
    for i := range detailInformation.SendSackHole {
        types.SetYListKey(detailInformation.SendSackHole[i], i)
        detailInformation.EntityData.Children.Append(types.GetSegmentPath(detailInformation.SendSackHole[i]), types.YChild{"SendSackHole", detailInformation.SendSackHole[i]})
    }
    detailInformation.EntityData.Leafs = types.NewOrderedMap()
    detailInformation.EntityData.Leafs.Append("pcb-id", types.YLeaf{"PcbId", detailInformation.PcbId})
    detailInformation.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", detailInformation.AddressFamily})
    detailInformation.EntityData.Leafs.Append("pcb", types.YLeaf{"Pcb", detailInformation.Pcb})
    detailInformation.EntityData.Leafs.Append("so", types.YLeaf{"So", detailInformation.So})
    detailInformation.EntityData.Leafs.Append("tcpcb", types.YLeaf{"Tcpcb", detailInformation.Tcpcb})
    detailInformation.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", detailInformation.VrfId})
    detailInformation.EntityData.Leafs.Append("connection-state", types.YLeaf{"ConnectionState", detailInformation.ConnectionState})
    detailInformation.EntityData.Leafs.Append("established-time", types.YLeaf{"EstablishedTime", detailInformation.EstablishedTime})
    detailInformation.EntityData.Leafs.Append("local-pid", types.YLeaf{"LocalPid", detailInformation.LocalPid})
    detailInformation.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", detailInformation.LocalPort})
    detailInformation.EntityData.Leafs.Append("foreign-port", types.YLeaf{"ForeignPort", detailInformation.ForeignPort})
    detailInformation.EntityData.Leafs.Append("packet-priority", types.YLeaf{"PacketPriority", detailInformation.PacketPriority})
    detailInformation.EntityData.Leafs.Append("packet-tos", types.YLeaf{"PacketTos", detailInformation.PacketTos})
    detailInformation.EntityData.Leafs.Append("packet-ttl", types.YLeaf{"PacketTtl", detailInformation.PacketTtl})
    detailInformation.EntityData.Leafs.Append("hash-index", types.YLeaf{"HashIndex", detailInformation.HashIndex})
    detailInformation.EntityData.Leafs.Append("current-receive-queue-size", types.YLeaf{"CurrentReceiveQueueSize", detailInformation.CurrentReceiveQueueSize})
    detailInformation.EntityData.Leafs.Append("max-receive-queue-size", types.YLeaf{"MaxReceiveQueueSize", detailInformation.MaxReceiveQueueSize})
    detailInformation.EntityData.Leafs.Append("current-send-queue-size", types.YLeaf{"CurrentSendQueueSize", detailInformation.CurrentSendQueueSize})
    detailInformation.EntityData.Leafs.Append("max-send-queue-size", types.YLeaf{"MaxSendQueueSize", detailInformation.MaxSendQueueSize})
    detailInformation.EntityData.Leafs.Append("current-receive-queue-packet-size", types.YLeaf{"CurrentReceiveQueuePacketSize", detailInformation.CurrentReceiveQueuePacketSize})
    detailInformation.EntityData.Leafs.Append("max-receive-queue-packet-size", types.YLeaf{"MaxReceiveQueuePacketSize", detailInformation.MaxReceiveQueuePacketSize})
    detailInformation.EntityData.Leafs.Append("save-queue-size", types.YLeaf{"SaveQueueSize", detailInformation.SaveQueueSize})
    detailInformation.EntityData.Leafs.Append("send-initial-sequence-num", types.YLeaf{"SendInitialSequenceNum", detailInformation.SendInitialSequenceNum})
    detailInformation.EntityData.Leafs.Append("send-unack-sequence-num", types.YLeaf{"SendUnackSequenceNum", detailInformation.SendUnackSequenceNum})
    detailInformation.EntityData.Leafs.Append("send-next-sequence-num", types.YLeaf{"SendNextSequenceNum", detailInformation.SendNextSequenceNum})
    detailInformation.EntityData.Leafs.Append("send-max-sequence-num", types.YLeaf{"SendMaxSequenceNum", detailInformation.SendMaxSequenceNum})
    detailInformation.EntityData.Leafs.Append("send-window-size", types.YLeaf{"SendWindowSize", detailInformation.SendWindowSize})
    detailInformation.EntityData.Leafs.Append("send-congestion-window-size", types.YLeaf{"SendCongestionWindowSize", detailInformation.SendCongestionWindowSize})
    detailInformation.EntityData.Leafs.Append("receive-initial-sequence-num", types.YLeaf{"ReceiveInitialSequenceNum", detailInformation.ReceiveInitialSequenceNum})
    detailInformation.EntityData.Leafs.Append("receive-next-sequence-num", types.YLeaf{"ReceiveNextSequenceNum", detailInformation.ReceiveNextSequenceNum})
    detailInformation.EntityData.Leafs.Append("receive-adv-window-size", types.YLeaf{"ReceiveAdvWindowSize", detailInformation.ReceiveAdvWindowSize})
    detailInformation.EntityData.Leafs.Append("receive-window-size", types.YLeaf{"ReceiveWindowSize", detailInformation.ReceiveWindowSize})
    detailInformation.EntityData.Leafs.Append("mss", types.YLeaf{"Mss", detailInformation.Mss})
    detailInformation.EntityData.Leafs.Append("peer-mss", types.YLeaf{"PeerMss", detailInformation.PeerMss})
    detailInformation.EntityData.Leafs.Append("srtt", types.YLeaf{"Srtt", detailInformation.Srtt})
    detailInformation.EntityData.Leafs.Append("rtto", types.YLeaf{"Rtto", detailInformation.Rtto})
    detailInformation.EntityData.Leafs.Append("krtt", types.YLeaf{"Krtt", detailInformation.Krtt})
    detailInformation.EntityData.Leafs.Append("srtv", types.YLeaf{"Srtv", detailInformation.Srtv})
    detailInformation.EntityData.Leafs.Append("min-rtt", types.YLeaf{"MinRtt", detailInformation.MinRtt})
    detailInformation.EntityData.Leafs.Append("max-rtt", types.YLeaf{"MaxRtt", detailInformation.MaxRtt})
    detailInformation.EntityData.Leafs.Append("retries", types.YLeaf{"Retries", detailInformation.Retries})
    detailInformation.EntityData.Leafs.Append("ack-hold-time", types.YLeaf{"AckHoldTime", detailInformation.AckHoldTime})
    detailInformation.EntityData.Leafs.Append("giveup-time", types.YLeaf{"GiveupTime", detailInformation.GiveupTime})
    detailInformation.EntityData.Leafs.Append("keep-alive-time", types.YLeaf{"KeepAliveTime", detailInformation.KeepAliveTime})
    detailInformation.EntityData.Leafs.Append("syn-wait-time", types.YLeaf{"SynWaitTime", detailInformation.SynWaitTime})
    detailInformation.EntityData.Leafs.Append("rxsy-naclname", types.YLeaf{"RxsyNaclname", detailInformation.RxsyNaclname})
    detailInformation.EntityData.Leafs.Append("soft-error", types.YLeaf{"SoftError", detailInformation.SoftError})
    detailInformation.EntityData.Leafs.Append("sock-error", types.YLeaf{"SockError", detailInformation.SockError})
    detailInformation.EntityData.Leafs.Append("is-retrans-forever", types.YLeaf{"IsRetransForever", detailInformation.IsRetransForever})
    detailInformation.EntityData.Leafs.Append("min-mss", types.YLeaf{"MinMss", detailInformation.MinMss})
    detailInformation.EntityData.Leafs.Append("max-mss", types.YLeaf{"MaxMss", detailInformation.MaxMss})
    detailInformation.EntityData.Leafs.Append("connect-retries", types.YLeaf{"ConnectRetries", detailInformation.ConnectRetries})
    detailInformation.EntityData.Leafs.Append("connect-retry-interval", types.YLeaf{"ConnectRetryInterval", detailInformation.ConnectRetryInterval})
    detailInformation.EntityData.Leafs.Append("receive-window-scale", types.YLeaf{"ReceiveWindowScale", detailInformation.ReceiveWindowScale})
    detailInformation.EntityData.Leafs.Append("send-window-scale", types.YLeaf{"SendWindowScale", detailInformation.SendWindowScale})
    detailInformation.EntityData.Leafs.Append("request-receive-window-scale", types.YLeaf{"RequestReceiveWindowScale", detailInformation.RequestReceiveWindowScale})
    detailInformation.EntityData.Leafs.Append("rqst-send-wnd-scale", types.YLeaf{"RqstSendWndScale", detailInformation.RqstSendWndScale})
    detailInformation.EntityData.Leafs.Append("time-stamp-recent", types.YLeaf{"TimeStampRecent", detailInformation.TimeStampRecent})
    detailInformation.EntityData.Leafs.Append("time-stamp-recent-age", types.YLeaf{"TimeStampRecentAge", detailInformation.TimeStampRecentAge})
    detailInformation.EntityData.Leafs.Append("last-ack-sent", types.YLeaf{"LastAckSent", detailInformation.LastAckSent})
    detailInformation.EntityData.Leafs.Append("sendbuf-lowwat", types.YLeaf{"SendbufLowwat", detailInformation.SendbufLowwat})
    detailInformation.EntityData.Leafs.Append("recvbuf-lowwat", types.YLeaf{"RecvbufLowwat", detailInformation.RecvbufLowwat})
    detailInformation.EntityData.Leafs.Append("sendbuf-hiwat", types.YLeaf{"SendbufHiwat", detailInformation.SendbufHiwat})
    detailInformation.EntityData.Leafs.Append("recvbuf-hiwat", types.YLeaf{"RecvbufHiwat", detailInformation.RecvbufHiwat})
    detailInformation.EntityData.Leafs.Append("sendbuf-notify-thresh", types.YLeaf{"SendbufNotifyThresh", detailInformation.SendbufNotifyThresh})
    detailInformation.EntityData.Leafs.Append("recvbuf-datasize", types.YLeaf{"RecvbufDatasize", detailInformation.RecvbufDatasize})
    detailInformation.EntityData.Leafs.Append("queue-length", types.YLeaf{"QueueLength", detailInformation.QueueLength})
    detailInformation.EntityData.Leafs.Append("queue-zero-length", types.YLeaf{"QueueZeroLength", detailInformation.QueueZeroLength})
    detailInformation.EntityData.Leafs.Append("queue-limit", types.YLeaf{"QueueLimit", detailInformation.QueueLimit})
    detailInformation.EntityData.Leafs.Append("socket-error", types.YLeaf{"SocketError", detailInformation.SocketError})
    detailInformation.EntityData.Leafs.Append("auto-rearm", types.YLeaf{"AutoRearm", detailInformation.AutoRearm})
    detailInformation.EntityData.Leafs.Append("send-pdu-count", types.YLeaf{"SendPduCount", detailInformation.SendPduCount})
    detailInformation.EntityData.Leafs.Append("output-ifhandle", types.YLeaf{"OutputIfhandle", detailInformation.OutputIfhandle})
    detailInformation.EntityData.Leafs.Append("fib-pd-ctx-size", types.YLeaf{"FibPdCtxSize", detailInformation.FibPdCtxSize})
    detailInformation.EntityData.Leafs.Append("num-labels", types.YLeaf{"NumLabels", detailInformation.NumLabels})
    detailInformation.EntityData.Leafs.Append("local-app-instance", types.YLeaf{"LocalAppInstance", detailInformation.LocalAppInstance})

    detailInformation.EntityData.YListKeys = []string {"PcbId"}

    return &(detailInformation.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress
// Local address
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress struct {
    EntityData types.CommonEntityData
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

func (localAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "detail-information"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddress.AfName})
    localAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", localAddress.Ipv4Address})
    localAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", localAddress.Ipv6Address})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress
// Foreign address
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress struct {
    EntityData types.CommonEntityData
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

func (foreignAddress *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "detail-information"
    foreignAddress.EntityData.SegmentPath = "foreign-address"
    foreignAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + foreignAddress.EntityData.SegmentPath
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", foreignAddress.AfName})
    foreignAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", foreignAddress.Ipv4Address})
    foreignAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", foreignAddress.Ipv6Address})

    foreignAddress.EntityData.YListKeys = []string {}

    return &(foreignAddress.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags
// Socket option flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags struct {
    EntityData types.CommonEntityData
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

    // Send buffer scaled. The type is bool.
    SndBufScaled interface{}

    // Receive buffer scaled. The type is bool.
    RcvBufScaled interface{}
}

func (socketOptionFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketOptionFlags) GetEntityData() *types.CommonEntityData {
    socketOptionFlags.EntityData.YFilter = socketOptionFlags.YFilter
    socketOptionFlags.EntityData.YangName = "socket-option-flags"
    socketOptionFlags.EntityData.BundleName = "cisco_ios_xr"
    socketOptionFlags.EntityData.ParentYangName = "detail-information"
    socketOptionFlags.EntityData.SegmentPath = "socket-option-flags"
    socketOptionFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + socketOptionFlags.EntityData.SegmentPath
    socketOptionFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    socketOptionFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    socketOptionFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    socketOptionFlags.EntityData.Children = types.NewOrderedMap()
    socketOptionFlags.EntityData.Leafs = types.NewOrderedMap()
    socketOptionFlags.EntityData.Leafs.Append("debug", types.YLeaf{"Debug", socketOptionFlags.Debug})
    socketOptionFlags.EntityData.Leafs.Append("accept-connection", types.YLeaf{"AcceptConnection", socketOptionFlags.AcceptConnection})
    socketOptionFlags.EntityData.Leafs.Append("reuse-address", types.YLeaf{"ReuseAddress", socketOptionFlags.ReuseAddress})
    socketOptionFlags.EntityData.Leafs.Append("keep-alive", types.YLeaf{"KeepAlive", socketOptionFlags.KeepAlive})
    socketOptionFlags.EntityData.Leafs.Append("dont-route", types.YLeaf{"DontRoute", socketOptionFlags.DontRoute})
    socketOptionFlags.EntityData.Leafs.Append("broadcast", types.YLeaf{"Broadcast", socketOptionFlags.Broadcast})
    socketOptionFlags.EntityData.Leafs.Append("use-loopback", types.YLeaf{"UseLoopback", socketOptionFlags.UseLoopback})
    socketOptionFlags.EntityData.Leafs.Append("linger", types.YLeaf{"Linger", socketOptionFlags.Linger})
    socketOptionFlags.EntityData.Leafs.Append("out-of-band-inline", types.YLeaf{"OutOfBandInline", socketOptionFlags.OutOfBandInline})
    socketOptionFlags.EntityData.Leafs.Append("reuse-port", types.YLeaf{"ReusePort", socketOptionFlags.ReusePort})
    socketOptionFlags.EntityData.Leafs.Append("nonblocking-io", types.YLeaf{"NonblockingIo", socketOptionFlags.NonblockingIo})
    socketOptionFlags.EntityData.Leafs.Append("snd-buf-scaled", types.YLeaf{"SndBufScaled", socketOptionFlags.SndBufScaled})
    socketOptionFlags.EntityData.Leafs.Append("rcv-buf-scaled", types.YLeaf{"RcvBufScaled", socketOptionFlags.RcvBufScaled})

    socketOptionFlags.EntityData.YListKeys = []string {}

    return &(socketOptionFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags
// Socket state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags struct {
    EntityData types.CommonEntityData
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

func (socketStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SocketStateFlags) GetEntityData() *types.CommonEntityData {
    socketStateFlags.EntityData.YFilter = socketStateFlags.YFilter
    socketStateFlags.EntityData.YangName = "socket-state-flags"
    socketStateFlags.EntityData.BundleName = "cisco_ios_xr"
    socketStateFlags.EntityData.ParentYangName = "detail-information"
    socketStateFlags.EntityData.SegmentPath = "socket-state-flags"
    socketStateFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + socketStateFlags.EntityData.SegmentPath
    socketStateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    socketStateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    socketStateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    socketStateFlags.EntityData.Children = types.NewOrderedMap()
    socketStateFlags.EntityData.Leafs = types.NewOrderedMap()
    socketStateFlags.EntityData.Leafs.Append("no-file-descriptor-reference", types.YLeaf{"NoFileDescriptorReference", socketStateFlags.NoFileDescriptorReference})
    socketStateFlags.EntityData.Leafs.Append("is-connected", types.YLeaf{"IsConnected", socketStateFlags.IsConnected})
    socketStateFlags.EntityData.Leafs.Append("is-connecting", types.YLeaf{"IsConnecting", socketStateFlags.IsConnecting})
    socketStateFlags.EntityData.Leafs.Append("is-disconnecting", types.YLeaf{"IsDisconnecting", socketStateFlags.IsDisconnecting})
    socketStateFlags.EntityData.Leafs.Append("cant-send-more", types.YLeaf{"CantSendMore", socketStateFlags.CantSendMore})
    socketStateFlags.EntityData.Leafs.Append("cant-receive-more", types.YLeaf{"CantReceiveMore", socketStateFlags.CantReceiveMore})
    socketStateFlags.EntityData.Leafs.Append("received-at-mark", types.YLeaf{"ReceivedAtMark", socketStateFlags.ReceivedAtMark})
    socketStateFlags.EntityData.Leafs.Append("privileged", types.YLeaf{"Privileged", socketStateFlags.Privileged})
    socketStateFlags.EntityData.Leafs.Append("block-close", types.YLeaf{"BlockClose", socketStateFlags.BlockClose})
    socketStateFlags.EntityData.Leafs.Append("async-io-notify", types.YLeaf{"AsyncIoNotify", socketStateFlags.AsyncIoNotify})
    socketStateFlags.EntityData.Leafs.Append("is-confirming", types.YLeaf{"IsConfirming", socketStateFlags.IsConfirming})
    socketStateFlags.EntityData.Leafs.Append("is-solock", types.YLeaf{"IsSolock", socketStateFlags.IsSolock})
    socketStateFlags.EntityData.Leafs.Append("is-detached", types.YLeaf{"IsDetached", socketStateFlags.IsDetached})
    socketStateFlags.EntityData.Leafs.Append("block-receive", types.YLeaf{"BlockReceive", socketStateFlags.BlockReceive})
    socketStateFlags.EntityData.Leafs.Append("block-send", types.YLeaf{"BlockSend", socketStateFlags.BlockSend})

    socketStateFlags.EntityData.YListKeys = []string {}

    return &(socketStateFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags
// Connection feature flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags struct {
    EntityData types.CommonEntityData
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

func (featureFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FeatureFlags) GetEntityData() *types.CommonEntityData {
    featureFlags.EntityData.YFilter = featureFlags.YFilter
    featureFlags.EntityData.YangName = "feature-flags"
    featureFlags.EntityData.BundleName = "cisco_ios_xr"
    featureFlags.EntityData.ParentYangName = "detail-information"
    featureFlags.EntityData.SegmentPath = "feature-flags"
    featureFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + featureFlags.EntityData.SegmentPath
    featureFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    featureFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    featureFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    featureFlags.EntityData.Children = types.NewOrderedMap()
    featureFlags.EntityData.Leafs = types.NewOrderedMap()
    featureFlags.EntityData.Leafs.Append("selective-ack", types.YLeaf{"SelectiveAck", featureFlags.SelectiveAck})
    featureFlags.EntityData.Leafs.Append("md5", types.YLeaf{"Md5", featureFlags.Md5})
    featureFlags.EntityData.Leafs.Append("timestamps", types.YLeaf{"Timestamps", featureFlags.Timestamps})
    featureFlags.EntityData.Leafs.Append("window-scaling", types.YLeaf{"WindowScaling", featureFlags.WindowScaling})
    featureFlags.EntityData.Leafs.Append("nagle", types.YLeaf{"Nagle", featureFlags.Nagle})
    featureFlags.EntityData.Leafs.Append("giveup-timer", types.YLeaf{"GiveupTimer", featureFlags.GiveupTimer})
    featureFlags.EntityData.Leafs.Append("connection-keep-alive-timer", types.YLeaf{"ConnectionKeepAliveTimer", featureFlags.ConnectionKeepAliveTimer})
    featureFlags.EntityData.Leafs.Append("path-mtu-discovery", types.YLeaf{"PathMtuDiscovery", featureFlags.PathMtuDiscovery})
    featureFlags.EntityData.Leafs.Append("mss-cisco", types.YLeaf{"MssCisco", featureFlags.MssCisco})

    featureFlags.EntityData.YListKeys = []string {}

    return &(featureFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags
// Connection state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags struct {
    EntityData types.CommonEntityData
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

func (stateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_StateFlags) GetEntityData() *types.CommonEntityData {
    stateFlags.EntityData.YFilter = stateFlags.YFilter
    stateFlags.EntityData.YangName = "state-flags"
    stateFlags.EntityData.BundleName = "cisco_ios_xr"
    stateFlags.EntityData.ParentYangName = "detail-information"
    stateFlags.EntityData.SegmentPath = "state-flags"
    stateFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + stateFlags.EntityData.SegmentPath
    stateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateFlags.EntityData.Children = types.NewOrderedMap()
    stateFlags.EntityData.Leafs = types.NewOrderedMap()
    stateFlags.EntityData.Leafs.Append("nagle-wait", types.YLeaf{"NagleWait", stateFlags.NagleWait})
    stateFlags.EntityData.Leafs.Append("ack-needed", types.YLeaf{"AckNeeded", stateFlags.AckNeeded})
    stateFlags.EntityData.Leafs.Append("fin-sent", types.YLeaf{"FinSent", stateFlags.FinSent})
    stateFlags.EntityData.Leafs.Append("probing", types.YLeaf{"Probing", stateFlags.Probing})
    stateFlags.EntityData.Leafs.Append("need-push", types.YLeaf{"NeedPush", stateFlags.NeedPush})
    stateFlags.EntityData.Leafs.Append("pushed", types.YLeaf{"Pushed", stateFlags.Pushed})
    stateFlags.EntityData.Leafs.Append("in-syn-cache", types.YLeaf{"InSynCache", stateFlags.InSynCache})
    stateFlags.EntityData.Leafs.Append("path-mtu-ager", types.YLeaf{"PathMtuAger", stateFlags.PathMtuAger})

    stateFlags.EntityData.YListKeys = []string {}

    return &(stateFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags
// Connection request flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags struct {
    EntityData types.CommonEntityData
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

func (requestFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_RequestFlags) GetEntityData() *types.CommonEntityData {
    requestFlags.EntityData.YFilter = requestFlags.YFilter
    requestFlags.EntityData.YangName = "request-flags"
    requestFlags.EntityData.BundleName = "cisco_ios_xr"
    requestFlags.EntityData.ParentYangName = "detail-information"
    requestFlags.EntityData.SegmentPath = "request-flags"
    requestFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + requestFlags.EntityData.SegmentPath
    requestFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestFlags.EntityData.Children = types.NewOrderedMap()
    requestFlags.EntityData.Leafs = types.NewOrderedMap()
    requestFlags.EntityData.Leafs.Append("selective-ack", types.YLeaf{"SelectiveAck", requestFlags.SelectiveAck})
    requestFlags.EntityData.Leafs.Append("md5", types.YLeaf{"Md5", requestFlags.Md5})
    requestFlags.EntityData.Leafs.Append("timestamps", types.YLeaf{"Timestamps", requestFlags.Timestamps})
    requestFlags.EntityData.Leafs.Append("window-scaling", types.YLeaf{"WindowScaling", requestFlags.WindowScaling})
    requestFlags.EntityData.Leafs.Append("nagle", types.YLeaf{"Nagle", requestFlags.Nagle})
    requestFlags.EntityData.Leafs.Append("giveup-timer", types.YLeaf{"GiveupTimer", requestFlags.GiveupTimer})
    requestFlags.EntityData.Leafs.Append("connection-keep-alive-timer", types.YLeaf{"ConnectionKeepAliveTimer", requestFlags.ConnectionKeepAliveTimer})
    requestFlags.EntityData.Leafs.Append("path-mtu-discovery", types.YLeaf{"PathMtuDiscovery", requestFlags.PathMtuDiscovery})
    requestFlags.EntityData.Leafs.Append("mss-cisco", types.YLeaf{"MssCisco", requestFlags.MssCisco})

    requestFlags.EntityData.YListKeys = []string {}

    return &(requestFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags
// Receive buffer state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags struct {
    EntityData types.CommonEntityData
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

    // Packet Buffer is extended. The type is bool.
    PacketExtended interface{}
}

func (receiveBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_ReceiveBufStateFlags) GetEntityData() *types.CommonEntityData {
    receiveBufStateFlags.EntityData.YFilter = receiveBufStateFlags.YFilter
    receiveBufStateFlags.EntityData.YangName = "receive-buf-state-flags"
    receiveBufStateFlags.EntityData.BundleName = "cisco_ios_xr"
    receiveBufStateFlags.EntityData.ParentYangName = "detail-information"
    receiveBufStateFlags.EntityData.SegmentPath = "receive-buf-state-flags"
    receiveBufStateFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + receiveBufStateFlags.EntityData.SegmentPath
    receiveBufStateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiveBufStateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiveBufStateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiveBufStateFlags.EntityData.Children = types.NewOrderedMap()
    receiveBufStateFlags.EntityData.Leafs = types.NewOrderedMap()
    receiveBufStateFlags.EntityData.Leafs.Append("locked", types.YLeaf{"Locked", receiveBufStateFlags.Locked})
    receiveBufStateFlags.EntityData.Leafs.Append("waiting-for-lock", types.YLeaf{"WaitingForLock", receiveBufStateFlags.WaitingForLock})
    receiveBufStateFlags.EntityData.Leafs.Append("waiting-for-data", types.YLeaf{"WaitingForData", receiveBufStateFlags.WaitingForData})
    receiveBufStateFlags.EntityData.Leafs.Append("input-select", types.YLeaf{"InputSelect", receiveBufStateFlags.InputSelect})
    receiveBufStateFlags.EntityData.Leafs.Append("async-io", types.YLeaf{"AsyncIo", receiveBufStateFlags.AsyncIo})
    receiveBufStateFlags.EntityData.Leafs.Append("not-interruptible", types.YLeaf{"NotInterruptible", receiveBufStateFlags.NotInterruptible})
    receiveBufStateFlags.EntityData.Leafs.Append("io-timer-set", types.YLeaf{"IoTimerSet", receiveBufStateFlags.IoTimerSet})
    receiveBufStateFlags.EntityData.Leafs.Append("delayed-wakeup", types.YLeaf{"DelayedWakeup", receiveBufStateFlags.DelayedWakeup})
    receiveBufStateFlags.EntityData.Leafs.Append("wakeup", types.YLeaf{"Wakeup", receiveBufStateFlags.Wakeup})
    receiveBufStateFlags.EntityData.Leafs.Append("connect-wakeup", types.YLeaf{"ConnectWakeup", receiveBufStateFlags.ConnectWakeup})
    receiveBufStateFlags.EntityData.Leafs.Append("output-select", types.YLeaf{"OutputSelect", receiveBufStateFlags.OutputSelect})
    receiveBufStateFlags.EntityData.Leafs.Append("out-of-band-select", types.YLeaf{"OutOfBandSelect", receiveBufStateFlags.OutOfBandSelect})
    receiveBufStateFlags.EntityData.Leafs.Append("packet-extended", types.YLeaf{"PacketExtended", receiveBufStateFlags.PacketExtended})

    receiveBufStateFlags.EntityData.YListKeys = []string {}

    return &(receiveBufStateFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags
// Send buffer state flags
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags struct {
    EntityData types.CommonEntityData
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

    // Packet Buffer is extended. The type is bool.
    PacketExtended interface{}
}

func (sendBufStateFlags *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendBufStateFlags) GetEntityData() *types.CommonEntityData {
    sendBufStateFlags.EntityData.YFilter = sendBufStateFlags.YFilter
    sendBufStateFlags.EntityData.YangName = "send-buf-state-flags"
    sendBufStateFlags.EntityData.BundleName = "cisco_ios_xr"
    sendBufStateFlags.EntityData.ParentYangName = "detail-information"
    sendBufStateFlags.EntityData.SegmentPath = "send-buf-state-flags"
    sendBufStateFlags.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + sendBufStateFlags.EntityData.SegmentPath
    sendBufStateFlags.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendBufStateFlags.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendBufStateFlags.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendBufStateFlags.EntityData.Children = types.NewOrderedMap()
    sendBufStateFlags.EntityData.Leafs = types.NewOrderedMap()
    sendBufStateFlags.EntityData.Leafs.Append("locked", types.YLeaf{"Locked", sendBufStateFlags.Locked})
    sendBufStateFlags.EntityData.Leafs.Append("waiting-for-lock", types.YLeaf{"WaitingForLock", sendBufStateFlags.WaitingForLock})
    sendBufStateFlags.EntityData.Leafs.Append("waiting-for-data", types.YLeaf{"WaitingForData", sendBufStateFlags.WaitingForData})
    sendBufStateFlags.EntityData.Leafs.Append("input-select", types.YLeaf{"InputSelect", sendBufStateFlags.InputSelect})
    sendBufStateFlags.EntityData.Leafs.Append("async-io", types.YLeaf{"AsyncIo", sendBufStateFlags.AsyncIo})
    sendBufStateFlags.EntityData.Leafs.Append("not-interruptible", types.YLeaf{"NotInterruptible", sendBufStateFlags.NotInterruptible})
    sendBufStateFlags.EntityData.Leafs.Append("io-timer-set", types.YLeaf{"IoTimerSet", sendBufStateFlags.IoTimerSet})
    sendBufStateFlags.EntityData.Leafs.Append("delayed-wakeup", types.YLeaf{"DelayedWakeup", sendBufStateFlags.DelayedWakeup})
    sendBufStateFlags.EntityData.Leafs.Append("wakeup", types.YLeaf{"Wakeup", sendBufStateFlags.Wakeup})
    sendBufStateFlags.EntityData.Leafs.Append("connect-wakeup", types.YLeaf{"ConnectWakeup", sendBufStateFlags.ConnectWakeup})
    sendBufStateFlags.EntityData.Leafs.Append("output-select", types.YLeaf{"OutputSelect", sendBufStateFlags.OutputSelect})
    sendBufStateFlags.EntityData.Leafs.Append("out-of-band-select", types.YLeaf{"OutOfBandSelect", sendBufStateFlags.OutOfBandSelect})
    sendBufStateFlags.EntityData.Leafs.Append("packet-extended", types.YLeaf{"PacketExtended", sendBufStateFlags.PacketExtended})

    sendBufStateFlags.EntityData.YListKeys = []string {}

    return &(sendBufStateFlags.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibPdCtx
// Cached fib pd context
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibPdCtx struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (fibPdCtx *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibPdCtx) GetEntityData() *types.CommonEntityData {
    fibPdCtx.EntityData.YFilter = fibPdCtx.YFilter
    fibPdCtx.EntityData.YangName = "fib-pd-ctx"
    fibPdCtx.EntityData.BundleName = "cisco_ios_xr"
    fibPdCtx.EntityData.ParentYangName = "detail-information"
    fibPdCtx.EntityData.SegmentPath = "fib-pd-ctx" + types.AddNoKeyToken(fibPdCtx)
    fibPdCtx.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + fibPdCtx.EntityData.SegmentPath
    fibPdCtx.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibPdCtx.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibPdCtx.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibPdCtx.EntityData.Children = types.NewOrderedMap()
    fibPdCtx.EntityData.Leafs = types.NewOrderedMap()
    fibPdCtx.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", fibPdCtx.Entry})

    fibPdCtx.EntityData.YListKeys = []string {}

    return &(fibPdCtx.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibLabelOutput
// Cached Label stack
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibLabelOutput struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is interface{} with range: 0..4294967295.
    Entry interface{}
}

func (fibLabelOutput *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_FibLabelOutput) GetEntityData() *types.CommonEntityData {
    fibLabelOutput.EntityData.YFilter = fibLabelOutput.YFilter
    fibLabelOutput.EntityData.YangName = "fib-label-output"
    fibLabelOutput.EntityData.BundleName = "cisco_ios_xr"
    fibLabelOutput.EntityData.ParentYangName = "detail-information"
    fibLabelOutput.EntityData.SegmentPath = "fib-label-output" + types.AddNoKeyToken(fibLabelOutput)
    fibLabelOutput.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + fibLabelOutput.EntityData.SegmentPath
    fibLabelOutput.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fibLabelOutput.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fibLabelOutput.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fibLabelOutput.EntityData.Children = types.NewOrderedMap()
    fibLabelOutput.EntityData.Leafs = types.NewOrderedMap()
    fibLabelOutput.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", fibLabelOutput.Entry})

    fibLabelOutput.EntityData.YListKeys = []string {}

    return &(fibLabelOutput.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer
// Timers
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (timer *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "detail-information"
    timer.EntityData.SegmentPath = "timer" + types.AddNoKeyToken(timer)
    timer.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + timer.EntityData.SegmentPath
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()
    timer.EntityData.Leafs.Append("timer-type", types.YLeaf{"TimerType", timer.TimerType})
    timer.EntityData.Leafs.Append("timer-activations", types.YLeaf{"TimerActivations", timer.TimerActivations})
    timer.EntityData.Leafs.Append("timer-expirations", types.YLeaf{"TimerExpirations", timer.TimerExpirations})
    timer.EntityData.Leafs.Append("timer-next-activation", types.YLeaf{"TimerNextActivation", timer.TimerNextActivation})

    timer.EntityData.YListKeys = []string {}

    return &(timer.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk
// Seq nos. of sack blocks
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Start seq no. of sack block. The type is interface{} with range:
    // 0..4294967295.
    Start interface{}

    // End   seq no. of sack block. The type is interface{} with range:
    // 0..4294967295.
    End interface{}
}

func (sackBlk *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SackBlk) GetEntityData() *types.CommonEntityData {
    sackBlk.EntityData.YFilter = sackBlk.YFilter
    sackBlk.EntityData.YangName = "sack-blk"
    sackBlk.EntityData.BundleName = "cisco_ios_xr"
    sackBlk.EntityData.ParentYangName = "detail-information"
    sackBlk.EntityData.SegmentPath = "sack-blk" + types.AddNoKeyToken(sackBlk)
    sackBlk.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + sackBlk.EntityData.SegmentPath
    sackBlk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sackBlk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sackBlk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sackBlk.EntityData.Children = types.NewOrderedMap()
    sackBlk.EntityData.Leafs = types.NewOrderedMap()
    sackBlk.EntityData.Leafs.Append("start", types.YLeaf{"Start", sackBlk.Start})
    sackBlk.EntityData.Leafs.Append("end", types.YLeaf{"End", sackBlk.End})

    sackBlk.EntityData.YListKeys = []string {}

    return &(sackBlk.EntityData)
}

// TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole
// Sorted list of sack holes
type TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (sendSackHole *TcpConnection_Nodes_Node_DetailInformations_DetailInformation_SendSackHole) GetEntityData() *types.CommonEntityData {
    sendSackHole.EntityData.YFilter = sendSackHole.YFilter
    sendSackHole.EntityData.YangName = "send-sack-hole"
    sendSackHole.EntityData.BundleName = "cisco_ios_xr"
    sendSackHole.EntityData.ParentYangName = "detail-information"
    sendSackHole.EntityData.SegmentPath = "send-sack-hole" + types.AddNoKeyToken(sendSackHole)
    sendSackHole.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/detail-informations/detail-information/" + sendSackHole.EntityData.SegmentPath
    sendSackHole.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendSackHole.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendSackHole.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendSackHole.EntityData.Children = types.NewOrderedMap()
    sendSackHole.EntityData.Leafs = types.NewOrderedMap()
    sendSackHole.EntityData.Leafs.Append("start", types.YLeaf{"Start", sendSackHole.Start})
    sendSackHole.EntityData.Leafs.Append("end", types.YLeaf{"End", sendSackHole.End})
    sendSackHole.EntityData.Leafs.Append("duplicated-ack", types.YLeaf{"DuplicatedAck", sendSackHole.DuplicatedAck})
    sendSackHole.EntityData.Leafs.Append("retransmitted", types.YLeaf{"Retransmitted", sendSackHole.Retransmitted})

    sendSackHole.EntityData.YListKeys = []string {}

    return &(sendSackHole.EntityData)
}

// TcpConnection_Nodes_Node_Keychains
// Table listing keychains configured for TCP-AO.
type TcpConnection_Nodes_Node_Keychains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Details of a keychain. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain.
    Keychain []*TcpConnection_Nodes_Node_Keychains_Keychain
}

func (keychains *TcpConnection_Nodes_Node_Keychains) GetEntityData() *types.CommonEntityData {
    keychains.EntityData.YFilter = keychains.YFilter
    keychains.EntityData.YangName = "keychains"
    keychains.EntityData.BundleName = "cisco_ios_xr"
    keychains.EntityData.ParentYangName = "node"
    keychains.EntityData.SegmentPath = "keychains"
    keychains.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/" + keychains.EntityData.SegmentPath
    keychains.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychains.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychains.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychains.EntityData.Children = types.NewOrderedMap()
    keychains.EntityData.Children.Append("keychain", types.YChild{"Keychain", nil})
    for i := range keychains.Keychain {
        keychains.EntityData.Children.Append(types.GetSegmentPath(keychains.Keychain[i]), types.YChild{"Keychain", keychains.Keychain[i]})
    }
    keychains.EntityData.Leafs = types.NewOrderedMap()

    keychains.EntityData.YListKeys = []string {}

    return &(keychains.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain
// Details of a keychain
type TcpConnection_Nodes_Node_Keychains_Keychain struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Keychain name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    KeychainName interface{}

    // Keychain name. The type is string.
    ChainName interface{}

    // Is keychain configured?. The type is bool.
    IsConfigured interface{}

    // Is desired key available?. The type is bool.
    DesiredKeyAvailable interface{}

    // Desired key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    DesiredKeyId interface{}

    // Keys under this keychain. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_Keys.
    Keys []*TcpConnection_Nodes_Node_Keychains_Keychain_Keys

    // List of active keys. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_ActiveKey.
    ActiveKey []*TcpConnection_Nodes_Node_Keychains_Keychain_ActiveKey

    // Send IDs under this keychain. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_SendId.
    SendId []*TcpConnection_Nodes_Node_Keychains_Keychain_SendId

    // Receive IDs under this keychain. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId.
    ReceiveId []*TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId
}

func (keychain *TcpConnection_Nodes_Node_Keychains_Keychain) GetEntityData() *types.CommonEntityData {
    keychain.EntityData.YFilter = keychain.YFilter
    keychain.EntityData.YangName = "keychain"
    keychain.EntityData.BundleName = "cisco_ios_xr"
    keychain.EntityData.ParentYangName = "keychains"
    keychain.EntityData.SegmentPath = "keychain" + types.AddKeyToken(keychain.KeychainName, "keychain-name")
    keychain.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/" + keychain.EntityData.SegmentPath
    keychain.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keychain.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keychain.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keychain.EntityData.Children = types.NewOrderedMap()
    keychain.EntityData.Children.Append("keys", types.YChild{"Keys", nil})
    for i := range keychain.Keys {
        types.SetYListKey(keychain.Keys[i], i)
        keychain.EntityData.Children.Append(types.GetSegmentPath(keychain.Keys[i]), types.YChild{"Keys", keychain.Keys[i]})
    }
    keychain.EntityData.Children.Append("active-key", types.YChild{"ActiveKey", nil})
    for i := range keychain.ActiveKey {
        types.SetYListKey(keychain.ActiveKey[i], i)
        keychain.EntityData.Children.Append(types.GetSegmentPath(keychain.ActiveKey[i]), types.YChild{"ActiveKey", keychain.ActiveKey[i]})
    }
    keychain.EntityData.Children.Append("send-id", types.YChild{"SendId", nil})
    for i := range keychain.SendId {
        types.SetYListKey(keychain.SendId[i], i)
        keychain.EntityData.Children.Append(types.GetSegmentPath(keychain.SendId[i]), types.YChild{"SendId", keychain.SendId[i]})
    }
    keychain.EntityData.Children.Append("receive-id", types.YChild{"ReceiveId", nil})
    for i := range keychain.ReceiveId {
        types.SetYListKey(keychain.ReceiveId[i], i)
        keychain.EntityData.Children.Append(types.GetSegmentPath(keychain.ReceiveId[i]), types.YChild{"ReceiveId", keychain.ReceiveId[i]})
    }
    keychain.EntityData.Leafs = types.NewOrderedMap()
    keychain.EntityData.Leafs.Append("keychain-name", types.YLeaf{"KeychainName", keychain.KeychainName})
    keychain.EntityData.Leafs.Append("chain-name", types.YLeaf{"ChainName", keychain.ChainName})
    keychain.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", keychain.IsConfigured})
    keychain.EntityData.Leafs.Append("desired-key-available", types.YLeaf{"DesiredKeyAvailable", keychain.DesiredKeyAvailable})
    keychain.EntityData.Leafs.Append("desired-key-id", types.YLeaf{"DesiredKeyId", keychain.DesiredKeyId})

    keychain.EntityData.YListKeys = []string {"KeychainName"}

    return &(keychain.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_Keys
// Keys under this keychain
type TcpConnection_Nodes_Node_Keychains_Keychain_Keys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    KeyId interface{}

    // Is key active. The type is bool.
    IsActive interface{}

    // Is key expired. The type is bool.
    IsExpired interface{}

    // Is key valid. The type is bool.
    IsValid interface{}

    // Key invalid reason. The type is TcpKeyInvalidReason.
    Reason interface{}

    // Send ID. The type is interface{} with range: 0..255.
    SendId interface{}

    // Receive ID. The type is interface{} with range: 0..255.
    RecvId interface{}

    // Cryptography algorithm associated with the key. The type is TcpMacAlgo.
    CryptAlgo interface{}

    // Is key configured?. The type is bool.
    IsConfigured interface{}

    // Is overlapping key available?. The type is bool.
    OverlappingKeyAvailable interface{}

    // Overlapping key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    OverlappingKey interface{}

    // List of keys invalidated. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_Keys_InvalidatedKey.
    InvalidatedKey []*TcpConnection_Nodes_Node_Keychains_Keychain_Keys_InvalidatedKey
}

func (keys *TcpConnection_Nodes_Node_Keychains_Keychain_Keys) GetEntityData() *types.CommonEntityData {
    keys.EntityData.YFilter = keys.YFilter
    keys.EntityData.YangName = "keys"
    keys.EntityData.BundleName = "cisco_ios_xr"
    keys.EntityData.ParentYangName = "keychain"
    keys.EntityData.SegmentPath = "keys" + types.AddNoKeyToken(keys)
    keys.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/" + keys.EntityData.SegmentPath
    keys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keys.EntityData.Children = types.NewOrderedMap()
    keys.EntityData.Children.Append("invalidated-key", types.YChild{"InvalidatedKey", nil})
    for i := range keys.InvalidatedKey {
        types.SetYListKey(keys.InvalidatedKey[i], i)
        keys.EntityData.Children.Append(types.GetSegmentPath(keys.InvalidatedKey[i]), types.YChild{"InvalidatedKey", keys.InvalidatedKey[i]})
    }
    keys.EntityData.Leafs = types.NewOrderedMap()
    keys.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", keys.KeyId})
    keys.EntityData.Leafs.Append("is-active", types.YLeaf{"IsActive", keys.IsActive})
    keys.EntityData.Leafs.Append("is-expired", types.YLeaf{"IsExpired", keys.IsExpired})
    keys.EntityData.Leafs.Append("is-valid", types.YLeaf{"IsValid", keys.IsValid})
    keys.EntityData.Leafs.Append("reason", types.YLeaf{"Reason", keys.Reason})
    keys.EntityData.Leafs.Append("send-id", types.YLeaf{"SendId", keys.SendId})
    keys.EntityData.Leafs.Append("recv-id", types.YLeaf{"RecvId", keys.RecvId})
    keys.EntityData.Leafs.Append("crypt-algo", types.YLeaf{"CryptAlgo", keys.CryptAlgo})
    keys.EntityData.Leafs.Append("is-configured", types.YLeaf{"IsConfigured", keys.IsConfigured})
    keys.EntityData.Leafs.Append("overlapping-key-available", types.YLeaf{"OverlappingKeyAvailable", keys.OverlappingKeyAvailable})
    keys.EntityData.Leafs.Append("overlapping-key", types.YLeaf{"OverlappingKey", keys.OverlappingKey})

    keys.EntityData.YListKeys = []string {}

    return &(keys.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_Keys_InvalidatedKey
// List of keys invalidated
type TcpConnection_Nodes_Node_Keychains_Keychain_Keys_InvalidatedKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    KeyId interface{}
}

func (invalidatedKey *TcpConnection_Nodes_Node_Keychains_Keychain_Keys_InvalidatedKey) GetEntityData() *types.CommonEntityData {
    invalidatedKey.EntityData.YFilter = invalidatedKey.YFilter
    invalidatedKey.EntityData.YangName = "invalidated-key"
    invalidatedKey.EntityData.BundleName = "cisco_ios_xr"
    invalidatedKey.EntityData.ParentYangName = "keys"
    invalidatedKey.EntityData.SegmentPath = "invalidated-key" + types.AddNoKeyToken(invalidatedKey)
    invalidatedKey.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/keys/" + invalidatedKey.EntityData.SegmentPath
    invalidatedKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    invalidatedKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    invalidatedKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    invalidatedKey.EntityData.Children = types.NewOrderedMap()
    invalidatedKey.EntityData.Leafs = types.NewOrderedMap()
    invalidatedKey.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", invalidatedKey.KeyId})

    invalidatedKey.EntityData.YListKeys = []string {}

    return &(invalidatedKey.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_ActiveKey
// List of active keys
type TcpConnection_Nodes_Node_Keychains_Keychain_ActiveKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    KeyId interface{}
}

func (activeKey *TcpConnection_Nodes_Node_Keychains_Keychain_ActiveKey) GetEntityData() *types.CommonEntityData {
    activeKey.EntityData.YFilter = activeKey.YFilter
    activeKey.EntityData.YangName = "active-key"
    activeKey.EntityData.BundleName = "cisco_ios_xr"
    activeKey.EntityData.ParentYangName = "keychain"
    activeKey.EntityData.SegmentPath = "active-key" + types.AddNoKeyToken(activeKey)
    activeKey.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/" + activeKey.EntityData.SegmentPath
    activeKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activeKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activeKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activeKey.EntityData.Children = types.NewOrderedMap()
    activeKey.EntityData.Leafs = types.NewOrderedMap()
    activeKey.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", activeKey.KeyId})

    activeKey.EntityData.YListKeys = []string {}

    return &(activeKey.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_SendId
// Send IDs under this keychain
type TcpConnection_Nodes_Node_Keychains_Keychain_SendId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Identifier. The type is interface{} with range: 0..255.
    Id interface{}

    // List of keys having this id. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_SendId_Keys.
    Keys []*TcpConnection_Nodes_Node_Keychains_Keychain_SendId_Keys
}

func (sendId *TcpConnection_Nodes_Node_Keychains_Keychain_SendId) GetEntityData() *types.CommonEntityData {
    sendId.EntityData.YFilter = sendId.YFilter
    sendId.EntityData.YangName = "send-id"
    sendId.EntityData.BundleName = "cisco_ios_xr"
    sendId.EntityData.ParentYangName = "keychain"
    sendId.EntityData.SegmentPath = "send-id" + types.AddNoKeyToken(sendId)
    sendId.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/" + sendId.EntityData.SegmentPath
    sendId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sendId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sendId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sendId.EntityData.Children = types.NewOrderedMap()
    sendId.EntityData.Children.Append("keys", types.YChild{"Keys", nil})
    for i := range sendId.Keys {
        types.SetYListKey(sendId.Keys[i], i)
        sendId.EntityData.Children.Append(types.GetSegmentPath(sendId.Keys[i]), types.YChild{"Keys", sendId.Keys[i]})
    }
    sendId.EntityData.Leafs = types.NewOrderedMap()
    sendId.EntityData.Leafs.Append("id", types.YLeaf{"Id", sendId.Id})

    sendId.EntityData.YListKeys = []string {}

    return &(sendId.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_SendId_Keys
// List of keys having this id
type TcpConnection_Nodes_Node_Keychains_Keychain_SendId_Keys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    KeyId interface{}
}

func (keys *TcpConnection_Nodes_Node_Keychains_Keychain_SendId_Keys) GetEntityData() *types.CommonEntityData {
    keys.EntityData.YFilter = keys.YFilter
    keys.EntityData.YangName = "keys"
    keys.EntityData.BundleName = "cisco_ios_xr"
    keys.EntityData.ParentYangName = "send-id"
    keys.EntityData.SegmentPath = "keys" + types.AddNoKeyToken(keys)
    keys.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/send-id/" + keys.EntityData.SegmentPath
    keys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keys.EntityData.Children = types.NewOrderedMap()
    keys.EntityData.Leafs = types.NewOrderedMap()
    keys.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", keys.KeyId})

    keys.EntityData.YListKeys = []string {}

    return &(keys.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId
// Receive IDs under this keychain
type TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Identifier. The type is interface{} with range: 0..255.
    Id interface{}

    // List of keys having this id. The type is slice of
    // TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId_Keys.
    Keys []*TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId_Keys
}

func (receiveId *TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId) GetEntityData() *types.CommonEntityData {
    receiveId.EntityData.YFilter = receiveId.YFilter
    receiveId.EntityData.YangName = "receive-id"
    receiveId.EntityData.BundleName = "cisco_ios_xr"
    receiveId.EntityData.ParentYangName = "keychain"
    receiveId.EntityData.SegmentPath = "receive-id" + types.AddNoKeyToken(receiveId)
    receiveId.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/" + receiveId.EntityData.SegmentPath
    receiveId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    receiveId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    receiveId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    receiveId.EntityData.Children = types.NewOrderedMap()
    receiveId.EntityData.Children.Append("keys", types.YChild{"Keys", nil})
    for i := range receiveId.Keys {
        types.SetYListKey(receiveId.Keys[i], i)
        receiveId.EntityData.Children.Append(types.GetSegmentPath(receiveId.Keys[i]), types.YChild{"Keys", receiveId.Keys[i]})
    }
    receiveId.EntityData.Leafs = types.NewOrderedMap()
    receiveId.EntityData.Leafs.Append("id", types.YLeaf{"Id", receiveId.Id})

    receiveId.EntityData.YListKeys = []string {}

    return &(receiveId.EntityData)
}

// TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId_Keys
// List of keys having this id
type TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId_Keys struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Key identifier. The type is interface{} with range:
    // 0..18446744073709551615.
    KeyId interface{}
}

func (keys *TcpConnection_Nodes_Node_Keychains_Keychain_ReceiveId_Keys) GetEntityData() *types.CommonEntityData {
    keys.EntityData.YFilter = keys.YFilter
    keys.EntityData.YangName = "keys"
    keys.EntityData.BundleName = "cisco_ios_xr"
    keys.EntityData.ParentYangName = "receive-id"
    keys.EntityData.SegmentPath = "keys" + types.AddNoKeyToken(keys)
    keys.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/keychains/keychain/receive-id/" + keys.EntityData.SegmentPath
    keys.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keys.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keys.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keys.EntityData.Children = types.NewOrderedMap()
    keys.EntityData.Leafs = types.NewOrderedMap()
    keys.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", keys.KeyId})

    keys.EntityData.YListKeys = []string {}

    return &(keys.EntityData)
}

// TcpConnection_Nodes_Node_BriefInformations
// Table listing connections for which brief
// information is provided.Note that not all
// connections are listed in the brief table.
type TcpConnection_Nodes_Node_BriefInformations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about a TCP connection. The type is slice of
    // TcpConnection_Nodes_Node_BriefInformations_BriefInformation.
    BriefInformation []*TcpConnection_Nodes_Node_BriefInformations_BriefInformation
}

func (briefInformations *TcpConnection_Nodes_Node_BriefInformations) GetEntityData() *types.CommonEntityData {
    briefInformations.EntityData.YFilter = briefInformations.YFilter
    briefInformations.EntityData.YangName = "brief-informations"
    briefInformations.EntityData.BundleName = "cisco_ios_xr"
    briefInformations.EntityData.ParentYangName = "node"
    briefInformations.EntityData.SegmentPath = "brief-informations"
    briefInformations.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/" + briefInformations.EntityData.SegmentPath
    briefInformations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefInformations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefInformations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefInformations.EntityData.Children = types.NewOrderedMap()
    briefInformations.EntityData.Children.Append("brief-information", types.YChild{"BriefInformation", nil})
    for i := range briefInformations.BriefInformation {
        briefInformations.EntityData.Children.Append(types.GetSegmentPath(briefInformations.BriefInformation[i]), types.YChild{"BriefInformation", briefInformations.BriefInformation[i]})
    }
    briefInformations.EntityData.Leafs = types.NewOrderedMap()

    briefInformations.EntityData.YListKeys = []string {}

    return &(briefInformations.EntityData)
}

// TcpConnection_Nodes_Node_BriefInformations_BriefInformation
// Brief information about a TCP connection
type TcpConnection_Nodes_Node_BriefInformations_BriefInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Protocol Control Block ID. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
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

func (briefInformation *TcpConnection_Nodes_Node_BriefInformations_BriefInformation) GetEntityData() *types.CommonEntityData {
    briefInformation.EntityData.YFilter = briefInformation.YFilter
    briefInformation.EntityData.YangName = "brief-information"
    briefInformation.EntityData.BundleName = "cisco_ios_xr"
    briefInformation.EntityData.ParentYangName = "brief-informations"
    briefInformation.EntityData.SegmentPath = "brief-information" + types.AddKeyToken(briefInformation.PcbId, "pcb-id")
    briefInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/brief-informations/" + briefInformation.EntityData.SegmentPath
    briefInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefInformation.EntityData.Children = types.NewOrderedMap()
    briefInformation.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", &briefInformation.LocalAddress})
    briefInformation.EntityData.Children.Append("foreign-address", types.YChild{"ForeignAddress", &briefInformation.ForeignAddress})
    briefInformation.EntityData.Leafs = types.NewOrderedMap()
    briefInformation.EntityData.Leafs.Append("pcb-id", types.YLeaf{"PcbId", briefInformation.PcbId})
    briefInformation.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", briefInformation.AfName})
    briefInformation.EntityData.Leafs.Append("pcb", types.YLeaf{"Pcb", briefInformation.Pcb})
    briefInformation.EntityData.Leafs.Append("connection-state", types.YLeaf{"ConnectionState", briefInformation.ConnectionState})
    briefInformation.EntityData.Leafs.Append("local-pid", types.YLeaf{"LocalPid", briefInformation.LocalPid})
    briefInformation.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", briefInformation.LocalPort})
    briefInformation.EntityData.Leafs.Append("foreign-port", types.YLeaf{"ForeignPort", briefInformation.ForeignPort})
    briefInformation.EntityData.Leafs.Append("current-receive-queue-size", types.YLeaf{"CurrentReceiveQueueSize", briefInformation.CurrentReceiveQueueSize})
    briefInformation.EntityData.Leafs.Append("current-send-queue-size", types.YLeaf{"CurrentSendQueueSize", briefInformation.CurrentSendQueueSize})
    briefInformation.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", briefInformation.VrfId})

    briefInformation.EntityData.YListKeys = []string {"PcbId"}

    return &(briefInformation.EntityData)
}

// TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress
// Local address
type TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress struct {
    EntityData types.CommonEntityData
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

func (localAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "brief-information"
    localAddress.EntityData.SegmentPath = "local-address"
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/brief-informations/brief-information/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", localAddress.AfName})
    localAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", localAddress.Ipv4Address})
    localAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", localAddress.Ipv6Address})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress
// Foreign address
type TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress struct {
    EntityData types.CommonEntityData
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

func (foreignAddress *TcpConnection_Nodes_Node_BriefInformations_BriefInformation_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "brief-information"
    foreignAddress.EntityData.SegmentPath = "foreign-address"
    foreignAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-connection/nodes/node/brief-informations/brief-information/" + foreignAddress.EntityData.SegmentPath
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", foreignAddress.AfName})
    foreignAddress.EntityData.Leafs.Append("ipv4-address", types.YLeaf{"Ipv4Address", foreignAddress.Ipv4Address})
    foreignAddress.EntityData.Leafs.Append("ipv6-address", types.YLeaf{"Ipv6Address", foreignAddress.Ipv6Address})

    foreignAddress.EntityData.YListKeys = []string {}

    return &(foreignAddress.EntityData)
}

// Tcp
// tcp
type Tcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Node-specific TCP operational data.
    Nodes Tcp_Nodes
}

func (tcp *Tcp) GetEntityData() *types.CommonEntityData {
    tcp.EntityData.YFilter = tcp.YFilter
    tcp.EntityData.YangName = "tcp"
    tcp.EntityData.BundleName = "cisco_ios_xr"
    tcp.EntityData.ParentYangName = "Cisco-IOS-XR-ip-tcp-oper"
    tcp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-tcp-oper:tcp"
    tcp.EntityData.AbsolutePath = tcp.EntityData.SegmentPath
    tcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcp.EntityData.Children = types.NewOrderedMap()
    tcp.EntityData.Children.Append("nodes", types.YChild{"Nodes", &tcp.Nodes})
    tcp.EntityData.Leafs = types.NewOrderedMap()

    tcp.EntityData.YListKeys = []string {}

    return &(tcp.EntityData)
}

// Tcp_Nodes
// Node-specific TCP operational data
type Tcp_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP operational data for a particular node. The type is slice of
    // Tcp_Nodes_Node.
    Node []*Tcp_Nodes_Node
}

func (nodes *Tcp_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "tcp"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Tcp_Nodes_Node
// TCP operational data for a particular node
type Tcp_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Node name. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Statistical TCP operational data for a node.
    Statistics Tcp_Nodes_Node_Statistics
}

func (node *Tcp_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Tcp_Nodes_Node_Statistics
// Statistical TCP operational data for a node
type Tcp_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TCP Traffic statistics for IPv4.
    Ipv4Traffic Tcp_Nodes_Node_Statistics_Ipv4Traffic

    // TCP Traffic statistics for IPv6.
    Ipv6Traffic Tcp_Nodes_Node_Statistics_Ipv6Traffic
}

func (statistics *Tcp_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp/nodes/node/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("ipv4-traffic", types.YChild{"Ipv4Traffic", &statistics.Ipv4Traffic})
    statistics.EntityData.Children.Append("ipv6-traffic", types.YChild{"Ipv6Traffic", &statistics.Ipv6Traffic})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// Tcp_Nodes_Node_Statistics_Ipv4Traffic
// TCP Traffic statistics for IPv4
type Tcp_Nodes_Node_Statistics_Ipv4Traffic struct {
    EntityData types.CommonEntityData
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

func (ipv4Traffic *Tcp_Nodes_Node_Statistics_Ipv4Traffic) GetEntityData() *types.CommonEntityData {
    ipv4Traffic.EntityData.YFilter = ipv4Traffic.YFilter
    ipv4Traffic.EntityData.YangName = "ipv4-traffic"
    ipv4Traffic.EntityData.BundleName = "cisco_ios_xr"
    ipv4Traffic.EntityData.ParentYangName = "statistics"
    ipv4Traffic.EntityData.SegmentPath = "ipv4-traffic"
    ipv4Traffic.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp/nodes/node/statistics/" + ipv4Traffic.EntityData.SegmentPath
    ipv4Traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Traffic.EntityData.Children = types.NewOrderedMap()
    ipv4Traffic.EntityData.Leafs = types.NewOrderedMap()
    ipv4Traffic.EntityData.Leafs.Append("tcp-input-packets", types.YLeaf{"TcpInputPackets", ipv4Traffic.TcpInputPackets})
    ipv4Traffic.EntityData.Leafs.Append("tcp-checksum-error-packets", types.YLeaf{"TcpChecksumErrorPackets", ipv4Traffic.TcpChecksumErrorPackets})
    ipv4Traffic.EntityData.Leafs.Append("tcp-dropped-packets", types.YLeaf{"TcpDroppedPackets", ipv4Traffic.TcpDroppedPackets})
    ipv4Traffic.EntityData.Leafs.Append("tcp-output-packets", types.YLeaf{"TcpOutputPackets", ipv4Traffic.TcpOutputPackets})
    ipv4Traffic.EntityData.Leafs.Append("tcp-retransmitted-packets", types.YLeaf{"TcpRetransmittedPackets", ipv4Traffic.TcpRetransmittedPackets})

    ipv4Traffic.EntityData.YListKeys = []string {}

    return &(ipv4Traffic.EntityData)
}

// Tcp_Nodes_Node_Statistics_Ipv6Traffic
// TCP Traffic statistics for IPv6
type Tcp_Nodes_Node_Statistics_Ipv6Traffic struct {
    EntityData types.CommonEntityData
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

func (ipv6Traffic *Tcp_Nodes_Node_Statistics_Ipv6Traffic) GetEntityData() *types.CommonEntityData {
    ipv6Traffic.EntityData.YFilter = ipv6Traffic.YFilter
    ipv6Traffic.EntityData.YangName = "ipv6-traffic"
    ipv6Traffic.EntityData.BundleName = "cisco_ios_xr"
    ipv6Traffic.EntityData.ParentYangName = "statistics"
    ipv6Traffic.EntityData.SegmentPath = "ipv6-traffic"
    ipv6Traffic.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp/nodes/node/statistics/" + ipv6Traffic.EntityData.SegmentPath
    ipv6Traffic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Traffic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Traffic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Traffic.EntityData.Children = types.NewOrderedMap()
    ipv6Traffic.EntityData.Leafs = types.NewOrderedMap()
    ipv6Traffic.EntityData.Leafs.Append("tcp-input-packets", types.YLeaf{"TcpInputPackets", ipv6Traffic.TcpInputPackets})
    ipv6Traffic.EntityData.Leafs.Append("tcp-checksum-error-packets", types.YLeaf{"TcpChecksumErrorPackets", ipv6Traffic.TcpChecksumErrorPackets})
    ipv6Traffic.EntityData.Leafs.Append("tcp-dropped-packets", types.YLeaf{"TcpDroppedPackets", ipv6Traffic.TcpDroppedPackets})
    ipv6Traffic.EntityData.Leafs.Append("tcp-output-packets", types.YLeaf{"TcpOutputPackets", ipv6Traffic.TcpOutputPackets})
    ipv6Traffic.EntityData.Leafs.Append("tcp-retransmitted-packets", types.YLeaf{"TcpRetransmittedPackets", ipv6Traffic.TcpRetransmittedPackets})

    ipv6Traffic.EntityData.YListKeys = []string {}

    return &(ipv6Traffic.EntityData)
}

// TcpNsr
// tcp nsr
type TcpNsr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of information about all nodes present on the system.
    Nodes TcpNsr_Nodes
}

func (tcpNsr *TcpNsr) GetEntityData() *types.CommonEntityData {
    tcpNsr.EntityData.YFilter = tcpNsr.YFilter
    tcpNsr.EntityData.YangName = "tcp-nsr"
    tcpNsr.EntityData.BundleName = "cisco_ios_xr"
    tcpNsr.EntityData.ParentYangName = "Cisco-IOS-XR-ip-tcp-oper"
    tcpNsr.EntityData.SegmentPath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr"
    tcpNsr.EntityData.AbsolutePath = tcpNsr.EntityData.SegmentPath
    tcpNsr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tcpNsr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tcpNsr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tcpNsr.EntityData.Children = types.NewOrderedMap()
    tcpNsr.EntityData.Children.Append("nodes", types.YChild{"Nodes", &tcpNsr.Nodes})
    tcpNsr.EntityData.Leafs = types.NewOrderedMap()

    tcpNsr.EntityData.YListKeys = []string {}

    return &(tcpNsr.EntityData)
}

// TcpNsr_Nodes
// Table of information about all nodes present on
// the system
type TcpNsr_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about a single node. The type is slice of TcpNsr_Nodes_Node.
    Node []*TcpNsr_Nodes_Node
}

func (nodes *TcpNsr_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "tcp-nsr"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/" + nodes.EntityData.SegmentPath
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// TcpNsr_Nodes_Node
// Information about a single node
type TcpNsr_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (node *TcpNsr_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.Id, "id")
    node.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/" + node.EntityData.SegmentPath
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Children.Append("session", types.YChild{"Session", &node.Session})
    node.EntityData.Children.Append("client", types.YChild{"Client", &node.Client})
    node.EntityData.Children.Append("session-set", types.YChild{"SessionSet", &node.SessionSet})
    node.EntityData.Children.Append("statistics", types.YChild{"Statistics", &node.Statistics})
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("id", types.YLeaf{"Id", node.Id})

    node.EntityData.YListKeys = []string {"Id"}

    return &(node.EntityData)
}

// TcpNsr_Nodes_Node_Session
// Information about TCP NSR Sessions
type TcpNsr_Nodes_Node_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Information about TCP NSR Sessions.
    BriefSessions TcpNsr_Nodes_Node_Session_BriefSessions

    // Table about TCP NSR Sessions details.
    DetailSessions TcpNsr_Nodes_Node_Session_DetailSessions
}

func (session *TcpNsr_Nodes_Node_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "node"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("brief-sessions", types.YChild{"BriefSessions", &session.BriefSessions})
    session.EntityData.Children.Append("detail-sessions", types.YChild{"DetailSessions", &session.DetailSessions})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// TcpNsr_Nodes_Node_Session_BriefSessions
// Information about TCP NSR Sessions
type TcpNsr_Nodes_Node_Session_BriefSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about NSR Sessions. The type is slice of
    // TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession.
    BriefSession []*TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession
}

func (briefSessions *TcpNsr_Nodes_Node_Session_BriefSessions) GetEntityData() *types.CommonEntityData {
    briefSessions.EntityData.YFilter = briefSessions.YFilter
    briefSessions.EntityData.YangName = "brief-sessions"
    briefSessions.EntityData.BundleName = "cisco_ios_xr"
    briefSessions.EntityData.ParentYangName = "session"
    briefSessions.EntityData.SegmentPath = "brief-sessions"
    briefSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/" + briefSessions.EntityData.SegmentPath
    briefSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefSessions.EntityData.Children = types.NewOrderedMap()
    briefSessions.EntityData.Children.Append("brief-session", types.YChild{"BriefSession", nil})
    for i := range briefSessions.BriefSession {
        briefSessions.EntityData.Children.Append(types.GetSegmentPath(briefSessions.BriefSession[i]), types.YChild{"BriefSession", briefSessions.BriefSession[i]})
    }
    briefSessions.EntityData.Leafs = types.NewOrderedMap()

    briefSessions.EntityData.YListKeys = []string {}

    return &(briefSessions.EntityData)
}

// TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession
// Brief information about NSR Sessions
type TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Local address. The type is slice of
    // TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_LocalAddress.
    LocalAddress []*TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_LocalAddress

    // Foreign address. The type is slice of
    // TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_ForeignAddress.
    ForeignAddress []*TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_ForeignAddress
}

func (briefSession *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession) GetEntityData() *types.CommonEntityData {
    briefSession.EntityData.YFilter = briefSession.YFilter
    briefSession.EntityData.YangName = "brief-session"
    briefSession.EntityData.BundleName = "cisco_ios_xr"
    briefSession.EntityData.ParentYangName = "brief-sessions"
    briefSession.EntityData.SegmentPath = "brief-session" + types.AddKeyToken(briefSession.Id, "id")
    briefSession.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/brief-sessions/" + briefSession.EntityData.SegmentPath
    briefSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefSession.EntityData.Children = types.NewOrderedMap()
    briefSession.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", nil})
    for i := range briefSession.LocalAddress {
        types.SetYListKey(briefSession.LocalAddress[i], i)
        briefSession.EntityData.Children.Append(types.GetSegmentPath(briefSession.LocalAddress[i]), types.YChild{"LocalAddress", briefSession.LocalAddress[i]})
    }
    briefSession.EntityData.Children.Append("foreign-address", types.YChild{"ForeignAddress", nil})
    for i := range briefSession.ForeignAddress {
        types.SetYListKey(briefSession.ForeignAddress[i], i)
        briefSession.EntityData.Children.Append(types.GetSegmentPath(briefSession.ForeignAddress[i]), types.YChild{"ForeignAddress", briefSession.ForeignAddress[i]})
    }
    briefSession.EntityData.Leafs = types.NewOrderedMap()
    briefSession.EntityData.Leafs.Append("id", types.YLeaf{"Id", briefSession.Id})
    briefSession.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", briefSession.AddressFamily})
    briefSession.EntityData.Leafs.Append("pcb", types.YLeaf{"Pcb", briefSession.Pcb})
    briefSession.EntityData.Leafs.Append("sscb", types.YLeaf{"Sscb", briefSession.Sscb})
    briefSession.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", briefSession.LocalPort})
    briefSession.EntityData.Leafs.Append("foreign-port", types.YLeaf{"ForeignPort", briefSession.ForeignPort})
    briefSession.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", briefSession.VrfId})
    briefSession.EntityData.Leafs.Append("is-admin-configured-up", types.YLeaf{"IsAdminConfiguredUp", briefSession.IsAdminConfiguredUp})
    briefSession.EntityData.Leafs.Append("is-us-operational-up", types.YLeaf{"IsUsOperationalUp", briefSession.IsUsOperationalUp})
    briefSession.EntityData.Leafs.Append("is-ds-operational-up", types.YLeaf{"IsDsOperationalUp", briefSession.IsDsOperationalUp})
    briefSession.EntityData.Leafs.Append("is-only-receive-path-replication", types.YLeaf{"IsOnlyReceivePathReplication", briefSession.IsOnlyReceivePathReplication})

    briefSession.EntityData.YListKeys = []string {"Id"}

    return &(briefSession.EntityData)
}

// TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_LocalAddress
// Local address
type TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Entry interface{}
}

func (localAddress *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "brief-session"
    localAddress.EntityData.SegmentPath = "local-address" + types.AddNoKeyToken(localAddress)
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/brief-sessions/brief-session/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", localAddress.Entry})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_ForeignAddress
// Foreign address
type TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_ForeignAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Entry interface{}
}

func (foreignAddress *TcpNsr_Nodes_Node_Session_BriefSessions_BriefSession_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "brief-session"
    foreignAddress.EntityData.SegmentPath = "foreign-address" + types.AddNoKeyToken(foreignAddress)
    foreignAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/brief-sessions/brief-session/" + foreignAddress.EntityData.SegmentPath
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", foreignAddress.Entry})

    foreignAddress.EntityData.YListKeys = []string {}

    return &(foreignAddress.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions
// Table about TCP NSR Sessions details
type TcpNsr_Nodes_Node_Session_DetailSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // showing detailed information of NSR Sessions. The type is slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession.
    DetailSession []*TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession
}

func (detailSessions *TcpNsr_Nodes_Node_Session_DetailSessions) GetEntityData() *types.CommonEntityData {
    detailSessions.EntityData.YFilter = detailSessions.YFilter
    detailSessions.EntityData.YangName = "detail-sessions"
    detailSessions.EntityData.BundleName = "cisco_ios_xr"
    detailSessions.EntityData.ParentYangName = "session"
    detailSessions.EntityData.SegmentPath = "detail-sessions"
    detailSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/" + detailSessions.EntityData.SegmentPath
    detailSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailSessions.EntityData.Children = types.NewOrderedMap()
    detailSessions.EntityData.Children.Append("detail-session", types.YChild{"DetailSession", nil})
    for i := range detailSessions.DetailSession {
        detailSessions.EntityData.Children.Append(types.GetSegmentPath(detailSessions.DetailSession[i]), types.YChild{"DetailSession", detailSessions.DetailSession[i]})
    }
    detailSessions.EntityData.Leafs = types.NewOrderedMap()

    detailSessions.EntityData.YListKeys = []string {}

    return &(detailSessions.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession
// showing detailed information of NSR Sessions
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

    // Sesson-set information.
    SetInformation TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation

    // Local address. The type is slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_LocalAddress.
    LocalAddress []*TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_LocalAddress

    // Foreign address. The type is slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_ForeignAddress.
    ForeignAddress []*TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_ForeignAddress

    // Sequence Number and datalength of each node in hold_pakqueue. The type is
    // slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue.
    PacketHoldQueue []*TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue

    // Sequence Number and datalength of each node in hold_iackqueue. The type is
    // slice of
    // TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue.
    InternalAckHoldQueue []*TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue
}

func (detailSession *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession) GetEntityData() *types.CommonEntityData {
    detailSession.EntityData.YFilter = detailSession.YFilter
    detailSession.EntityData.YangName = "detail-session"
    detailSession.EntityData.BundleName = "cisco_ios_xr"
    detailSession.EntityData.ParentYangName = "detail-sessions"
    detailSession.EntityData.SegmentPath = "detail-session" + types.AddKeyToken(detailSession.Id, "id")
    detailSession.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/detail-sessions/" + detailSession.EntityData.SegmentPath
    detailSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailSession.EntityData.Children = types.NewOrderedMap()
    detailSession.EntityData.Children.Append("set-information", types.YChild{"SetInformation", &detailSession.SetInformation})
    detailSession.EntityData.Children.Append("local-address", types.YChild{"LocalAddress", nil})
    for i := range detailSession.LocalAddress {
        types.SetYListKey(detailSession.LocalAddress[i], i)
        detailSession.EntityData.Children.Append(types.GetSegmentPath(detailSession.LocalAddress[i]), types.YChild{"LocalAddress", detailSession.LocalAddress[i]})
    }
    detailSession.EntityData.Children.Append("foreign-address", types.YChild{"ForeignAddress", nil})
    for i := range detailSession.ForeignAddress {
        types.SetYListKey(detailSession.ForeignAddress[i], i)
        detailSession.EntityData.Children.Append(types.GetSegmentPath(detailSession.ForeignAddress[i]), types.YChild{"ForeignAddress", detailSession.ForeignAddress[i]})
    }
    detailSession.EntityData.Children.Append("packet-hold-queue", types.YChild{"PacketHoldQueue", nil})
    for i := range detailSession.PacketHoldQueue {
        types.SetYListKey(detailSession.PacketHoldQueue[i], i)
        detailSession.EntityData.Children.Append(types.GetSegmentPath(detailSession.PacketHoldQueue[i]), types.YChild{"PacketHoldQueue", detailSession.PacketHoldQueue[i]})
    }
    detailSession.EntityData.Children.Append("internal-ack-hold-queue", types.YChild{"InternalAckHoldQueue", nil})
    for i := range detailSession.InternalAckHoldQueue {
        types.SetYListKey(detailSession.InternalAckHoldQueue[i], i)
        detailSession.EntityData.Children.Append(types.GetSegmentPath(detailSession.InternalAckHoldQueue[i]), types.YChild{"InternalAckHoldQueue", detailSession.InternalAckHoldQueue[i]})
    }
    detailSession.EntityData.Leafs = types.NewOrderedMap()
    detailSession.EntityData.Leafs.Append("id", types.YLeaf{"Id", detailSession.Id})
    detailSession.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", detailSession.AddressFamily})
    detailSession.EntityData.Leafs.Append("pcb", types.YLeaf{"Pcb", detailSession.Pcb})
    detailSession.EntityData.Leafs.Append("sscb", types.YLeaf{"Sscb", detailSession.Sscb})
    detailSession.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", detailSession.LocalPort})
    detailSession.EntityData.Leafs.Append("foreign-port", types.YLeaf{"ForeignPort", detailSession.ForeignPort})
    detailSession.EntityData.Leafs.Append("vrf-id", types.YLeaf{"VrfId", detailSession.VrfId})
    detailSession.EntityData.Leafs.Append("is-admin-configured-up", types.YLeaf{"IsAdminConfiguredUp", detailSession.IsAdminConfiguredUp})
    detailSession.EntityData.Leafs.Append("is-us-operational-up", types.YLeaf{"IsUsOperationalUp", detailSession.IsUsOperationalUp})
    detailSession.EntityData.Leafs.Append("is-ds-operational-up", types.YLeaf{"IsDsOperationalUp", detailSession.IsDsOperationalUp})
    detailSession.EntityData.Leafs.Append("is-only-receive-path-replication", types.YLeaf{"IsOnlyReceivePathReplication", detailSession.IsOnlyReceivePathReplication})
    detailSession.EntityData.Leafs.Append("cookie", types.YLeaf{"Cookie", detailSession.Cookie})
    detailSession.EntityData.Leafs.Append("is-session-replicated", types.YLeaf{"IsSessionReplicated", detailSession.IsSessionReplicated})
    detailSession.EntityData.Leafs.Append("is-session-synced", types.YLeaf{"IsSessionSynced", detailSession.IsSessionSynced})
    detailSession.EntityData.Leafs.Append("fist-standby-sequence-number", types.YLeaf{"FistStandbySequenceNumber", detailSession.FistStandbySequenceNumber})
    detailSession.EntityData.Leafs.Append("fssn-offset", types.YLeaf{"FssnOffset", detailSession.FssnOffset})
    detailSession.EntityData.Leafs.Append("nsr-down-reason", types.YLeaf{"NsrDownReason", detailSession.NsrDownReason})
    detailSession.EntityData.Leafs.Append("nsr-down-time", types.YLeaf{"NsrDownTime", detailSession.NsrDownTime})
    detailSession.EntityData.Leafs.Append("sequence-number-of-init-sync", types.YLeaf{"SequenceNumberOfInitSync", detailSession.SequenceNumberOfInitSync})
    detailSession.EntityData.Leafs.Append("is-init-sync-in-progress", types.YLeaf{"IsInitSyncInProgress", detailSession.IsInitSyncInProgress})
    detailSession.EntityData.Leafs.Append("is-init-sync-second-phase", types.YLeaf{"IsInitSyncSecondPhase", detailSession.IsInitSyncSecondPhase})
    detailSession.EntityData.Leafs.Append("init-sync-error", types.YLeaf{"InitSyncError", detailSession.InitSyncError})
    detailSession.EntityData.Leafs.Append("is-init-sync-error-local", types.YLeaf{"IsInitSyncErrorLocal", detailSession.IsInitSyncErrorLocal})
    detailSession.EntityData.Leafs.Append("init-sync-start-time", types.YLeaf{"InitSyncStartTime", detailSession.InitSyncStartTime})
    detailSession.EntityData.Leafs.Append("init-sync-end-time", types.YLeaf{"InitSyncEndTime", detailSession.InitSyncEndTime})
    detailSession.EntityData.Leafs.Append("init-sync-flags", types.YLeaf{"InitSyncFlags", detailSession.InitSyncFlags})
    detailSession.EntityData.Leafs.Append("sequence-number-of-init-sync-up-stream", types.YLeaf{"SequenceNumberOfInitSyncUpStream", detailSession.SequenceNumberOfInitSyncUpStream})
    detailSession.EntityData.Leafs.Append("peer-endp-hdl-up-stream", types.YLeaf{"PeerEndpHdlUpStream", detailSession.PeerEndpHdlUpStream})
    detailSession.EntityData.Leafs.Append("init-sync-start-time-up-stream", types.YLeaf{"InitSyncStartTimeUpStream", detailSession.InitSyncStartTimeUpStream})
    detailSession.EntityData.Leafs.Append("init-sync-end-time-up-stream", types.YLeaf{"InitSyncEndTimeUpStream", detailSession.InitSyncEndTimeUpStream})
    detailSession.EntityData.Leafs.Append("fist-standby-sequence-number-up-stream", types.YLeaf{"FistStandbySequenceNumberUpStream", detailSession.FistStandbySequenceNumberUpStream})
    detailSession.EntityData.Leafs.Append("nsr-down-reason-up-stream", types.YLeaf{"NsrDownReasonUpStream", detailSession.NsrDownReasonUpStream})
    detailSession.EntityData.Leafs.Append("nsr-down-time-up-stream", types.YLeaf{"NsrDownTimeUpStream", detailSession.NsrDownTimeUpStream})
    detailSession.EntityData.Leafs.Append("sequence-number-of-init-sync-down-stream", types.YLeaf{"SequenceNumberOfInitSyncDownStream", detailSession.SequenceNumberOfInitSyncDownStream})
    detailSession.EntityData.Leafs.Append("peer-endp-hdl-down-stream", types.YLeaf{"PeerEndpHdlDownStream", detailSession.PeerEndpHdlDownStream})
    detailSession.EntityData.Leafs.Append("init-sync-start-time-down-stream", types.YLeaf{"InitSyncStartTimeDownStream", detailSession.InitSyncStartTimeDownStream})
    detailSession.EntityData.Leafs.Append("init-sync-end-time-down-stream", types.YLeaf{"InitSyncEndTimeDownStream", detailSession.InitSyncEndTimeDownStream})
    detailSession.EntityData.Leafs.Append("fist-standby-sequence-number-down-stream", types.YLeaf{"FistStandbySequenceNumberDownStream", detailSession.FistStandbySequenceNumberDownStream})
    detailSession.EntityData.Leafs.Append("nsr-down-reason-down-stream", types.YLeaf{"NsrDownReasonDownStream", detailSession.NsrDownReasonDownStream})
    detailSession.EntityData.Leafs.Append("nsr-down-time-down-stream", types.YLeaf{"NsrDownTimeDownStream", detailSession.NsrDownTimeDownStream})
    detailSession.EntityData.Leafs.Append("max-number-of-held-packet", types.YLeaf{"MaxNumberOfHeldPacket", detailSession.MaxNumberOfHeldPacket})
    detailSession.EntityData.Leafs.Append("max-number-of-held-packet-reach-time", types.YLeaf{"MaxNumberOfHeldPacketReachTime", detailSession.MaxNumberOfHeldPacketReachTime})
    detailSession.EntityData.Leafs.Append("max-number-of-held-internal-ack", types.YLeaf{"MaxNumberOfHeldInternalAck", detailSession.MaxNumberOfHeldInternalAck})
    detailSession.EntityData.Leafs.Append("max-number-of-held-internal-ack-reach-time", types.YLeaf{"MaxNumberOfHeldInternalAckReachTime", detailSession.MaxNumberOfHeldInternalAckReachTime})

    detailSession.EntityData.YListKeys = []string {"Id"}

    return &(detailSession.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation
// Sesson-set information
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation struct {
    EntityData types.CommonEntityData
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

func (setInformation *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_SetInformation) GetEntityData() *types.CommonEntityData {
    setInformation.EntityData.YFilter = setInformation.YFilter
    setInformation.EntityData.YangName = "set-information"
    setInformation.EntityData.BundleName = "cisco_ios_xr"
    setInformation.EntityData.ParentYangName = "detail-session"
    setInformation.EntityData.SegmentPath = "set-information"
    setInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/detail-sessions/detail-session/" + setInformation.EntityData.SegmentPath
    setInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    setInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    setInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    setInformation.EntityData.Children = types.NewOrderedMap()
    setInformation.EntityData.Leafs = types.NewOrderedMap()
    setInformation.EntityData.Leafs.Append("sscb", types.YLeaf{"Sscb", setInformation.Sscb})
    setInformation.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", setInformation.Pid})
    setInformation.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", setInformation.ClientName})
    setInformation.EntityData.Leafs.Append("client-instance", types.YLeaf{"ClientInstance", setInformation.ClientInstance})
    setInformation.EntityData.Leafs.Append("set-id", types.YLeaf{"SetId", setInformation.SetId})
    setInformation.EntityData.Leafs.Append("sso-role", types.YLeaf{"SsoRole", setInformation.SsoRole})
    setInformation.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", setInformation.Mode})
    setInformation.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", setInformation.AddressFamily})
    setInformation.EntityData.Leafs.Append("well-known-port", types.YLeaf{"WellKnownPort", setInformation.WellKnownPort})
    setInformation.EntityData.Leafs.Append("local-node", types.YLeaf{"LocalNode", setInformation.LocalNode})
    setInformation.EntityData.Leafs.Append("local-instance", types.YLeaf{"LocalInstance", setInformation.LocalInstance})
    setInformation.EntityData.Leafs.Append("protect-node", types.YLeaf{"ProtectNode", setInformation.ProtectNode})
    setInformation.EntityData.Leafs.Append("protect-instance", types.YLeaf{"ProtectInstance", setInformation.ProtectInstance})
    setInformation.EntityData.Leafs.Append("number-of-sessions", types.YLeaf{"NumberOfSessions", setInformation.NumberOfSessions})
    setInformation.EntityData.Leafs.Append("number-of-synced-sessions-up-stream", types.YLeaf{"NumberOfSyncedSessionsUpStream", setInformation.NumberOfSyncedSessionsUpStream})
    setInformation.EntityData.Leafs.Append("number-of-synced-sessions-down-stream", types.YLeaf{"NumberOfSyncedSessionsDownStream", setInformation.NumberOfSyncedSessionsDownStream})
    setInformation.EntityData.Leafs.Append("is-init-sync-in-progress", types.YLeaf{"IsInitSyncInProgress", setInformation.IsInitSyncInProgress})
    setInformation.EntityData.Leafs.Append("is-sscb-init-sync-ready", types.YLeaf{"IsSscbInitSyncReady", setInformation.IsSscbInitSyncReady})

    setInformation.EntityData.YListKeys = []string {}

    return &(setInformation.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_LocalAddress
// Local address
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_LocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Entry interface{}
}

func (localAddress *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_LocalAddress) GetEntityData() *types.CommonEntityData {
    localAddress.EntityData.YFilter = localAddress.YFilter
    localAddress.EntityData.YangName = "local-address"
    localAddress.EntityData.BundleName = "cisco_ios_xr"
    localAddress.EntityData.ParentYangName = "detail-session"
    localAddress.EntityData.SegmentPath = "local-address" + types.AddNoKeyToken(localAddress)
    localAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/detail-sessions/detail-session/" + localAddress.EntityData.SegmentPath
    localAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localAddress.EntityData.Children = types.NewOrderedMap()
    localAddress.EntityData.Leafs = types.NewOrderedMap()
    localAddress.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", localAddress.Entry})

    localAddress.EntityData.YListKeys = []string {}

    return &(localAddress.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_ForeignAddress
// Foreign address
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_ForeignAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Entry interface{}
}

func (foreignAddress *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_ForeignAddress) GetEntityData() *types.CommonEntityData {
    foreignAddress.EntityData.YFilter = foreignAddress.YFilter
    foreignAddress.EntityData.YangName = "foreign-address"
    foreignAddress.EntityData.BundleName = "cisco_ios_xr"
    foreignAddress.EntityData.ParentYangName = "detail-session"
    foreignAddress.EntityData.SegmentPath = "foreign-address" + types.AddNoKeyToken(foreignAddress)
    foreignAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/detail-sessions/detail-session/" + foreignAddress.EntityData.SegmentPath
    foreignAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    foreignAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    foreignAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    foreignAddress.EntityData.Children = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs = types.NewOrderedMap()
    foreignAddress.EntityData.Leafs.Append("entry", types.YLeaf{"Entry", foreignAddress.Entry})

    foreignAddress.EntityData.YListKeys = []string {}

    return &(foreignAddress.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue
// Sequence Number and datalength of each node in
// hold_pakqueue
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Data Length. The type is interface{} with range: 0..4294967295.
    DataLength interface{}

    // Ack Number. The type is interface{} with range: 0..4294967295.
    AcknoledgementNumber interface{}
}

func (packetHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_PacketHoldQueue) GetEntityData() *types.CommonEntityData {
    packetHoldQueue.EntityData.YFilter = packetHoldQueue.YFilter
    packetHoldQueue.EntityData.YangName = "packet-hold-queue"
    packetHoldQueue.EntityData.BundleName = "cisco_ios_xr"
    packetHoldQueue.EntityData.ParentYangName = "detail-session"
    packetHoldQueue.EntityData.SegmentPath = "packet-hold-queue" + types.AddNoKeyToken(packetHoldQueue)
    packetHoldQueue.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/detail-sessions/detail-session/" + packetHoldQueue.EntityData.SegmentPath
    packetHoldQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    packetHoldQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    packetHoldQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    packetHoldQueue.EntityData.Children = types.NewOrderedMap()
    packetHoldQueue.EntityData.Leafs = types.NewOrderedMap()
    packetHoldQueue.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", packetHoldQueue.SequenceNumber})
    packetHoldQueue.EntityData.Leafs.Append("data-length", types.YLeaf{"DataLength", packetHoldQueue.DataLength})
    packetHoldQueue.EntityData.Leafs.Append("acknoledgement-number", types.YLeaf{"AcknoledgementNumber", packetHoldQueue.AcknoledgementNumber})

    packetHoldQueue.EntityData.YListKeys = []string {}

    return &(packetHoldQueue.EntityData)
}

// TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue
// Sequence Number and datalength of each node in
// hold_iackqueue
type TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Sequence Number. The type is interface{} with range: 0..4294967295.
    SequenceNumber interface{}

    // Data Length. The type is interface{} with range: 0..4294967295.
    DataLength interface{}

    // Ack Number. The type is interface{} with range: 0..4294967295.
    AcknoledgementNumber interface{}
}

func (internalAckHoldQueue *TcpNsr_Nodes_Node_Session_DetailSessions_DetailSession_InternalAckHoldQueue) GetEntityData() *types.CommonEntityData {
    internalAckHoldQueue.EntityData.YFilter = internalAckHoldQueue.YFilter
    internalAckHoldQueue.EntityData.YangName = "internal-ack-hold-queue"
    internalAckHoldQueue.EntityData.BundleName = "cisco_ios_xr"
    internalAckHoldQueue.EntityData.ParentYangName = "detail-session"
    internalAckHoldQueue.EntityData.SegmentPath = "internal-ack-hold-queue" + types.AddNoKeyToken(internalAckHoldQueue)
    internalAckHoldQueue.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session/detail-sessions/detail-session/" + internalAckHoldQueue.EntityData.SegmentPath
    internalAckHoldQueue.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalAckHoldQueue.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalAckHoldQueue.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalAckHoldQueue.EntityData.Children = types.NewOrderedMap()
    internalAckHoldQueue.EntityData.Leafs = types.NewOrderedMap()
    internalAckHoldQueue.EntityData.Leafs.Append("sequence-number", types.YLeaf{"SequenceNumber", internalAckHoldQueue.SequenceNumber})
    internalAckHoldQueue.EntityData.Leafs.Append("data-length", types.YLeaf{"DataLength", internalAckHoldQueue.DataLength})
    internalAckHoldQueue.EntityData.Leafs.Append("acknoledgement-number", types.YLeaf{"AcknoledgementNumber", internalAckHoldQueue.AcknoledgementNumber})

    internalAckHoldQueue.EntityData.YListKeys = []string {}

    return &(internalAckHoldQueue.EntityData)
}

// TcpNsr_Nodes_Node_Client
// Information about TCP NSR Client
type TcpNsr_Nodes_Node_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table about TCP NSR Client details.
    DetailClients TcpNsr_Nodes_Node_Client_DetailClients

    // Information about TCP NSR Client.
    BriefClients TcpNsr_Nodes_Node_Client_BriefClients
}

func (client *TcpNsr_Nodes_Node_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "node"
    client.EntityData.SegmentPath = "client"
    client.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/" + client.EntityData.SegmentPath
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = types.NewOrderedMap()
    client.EntityData.Children.Append("detail-clients", types.YChild{"DetailClients", &client.DetailClients})
    client.EntityData.Children.Append("brief-clients", types.YChild{"BriefClients", &client.BriefClients})
    client.EntityData.Leafs = types.NewOrderedMap()

    client.EntityData.YListKeys = []string {}

    return &(client.EntityData)
}

// TcpNsr_Nodes_Node_Client_DetailClients
// Table about TCP NSR Client details
type TcpNsr_Nodes_Node_Client_DetailClients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // showing detailed information of NSR Clients. The type is slice of
    // TcpNsr_Nodes_Node_Client_DetailClients_DetailClient.
    DetailClient []*TcpNsr_Nodes_Node_Client_DetailClients_DetailClient
}

func (detailClients *TcpNsr_Nodes_Node_Client_DetailClients) GetEntityData() *types.CommonEntityData {
    detailClients.EntityData.YFilter = detailClients.YFilter
    detailClients.EntityData.YangName = "detail-clients"
    detailClients.EntityData.BundleName = "cisco_ios_xr"
    detailClients.EntityData.ParentYangName = "client"
    detailClients.EntityData.SegmentPath = "detail-clients"
    detailClients.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/client/" + detailClients.EntityData.SegmentPath
    detailClients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailClients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailClients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailClients.EntityData.Children = types.NewOrderedMap()
    detailClients.EntityData.Children.Append("detail-client", types.YChild{"DetailClient", nil})
    for i := range detailClients.DetailClient {
        detailClients.EntityData.Children.Append(types.GetSegmentPath(detailClients.DetailClient[i]), types.YChild{"DetailClient", detailClients.DetailClient[i]})
    }
    detailClients.EntityData.Leafs = types.NewOrderedMap()

    detailClients.EntityData.YListKeys = []string {}

    return &(detailClients.EntityData)
}

// TcpNsr_Nodes_Node_Client_DetailClients_DetailClient
// showing detailed information of NSR Clients
type TcpNsr_Nodes_Node_Client_DetailClients_DetailClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (detailClient *TcpNsr_Nodes_Node_Client_DetailClients_DetailClient) GetEntityData() *types.CommonEntityData {
    detailClient.EntityData.YFilter = detailClient.YFilter
    detailClient.EntityData.YangName = "detail-client"
    detailClient.EntityData.BundleName = "cisco_ios_xr"
    detailClient.EntityData.ParentYangName = "detail-clients"
    detailClient.EntityData.SegmentPath = "detail-client" + types.AddKeyToken(detailClient.Id, "id")
    detailClient.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/client/detail-clients/" + detailClient.EntityData.SegmentPath
    detailClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailClient.EntityData.Children = types.NewOrderedMap()
    detailClient.EntityData.Leafs = types.NewOrderedMap()
    detailClient.EntityData.Leafs.Append("id", types.YLeaf{"Id", detailClient.Id})
    detailClient.EntityData.Leafs.Append("ccb", types.YLeaf{"Ccb", detailClient.Ccb})
    detailClient.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", detailClient.Pid})
    detailClient.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", detailClient.ProcessName})
    detailClient.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", detailClient.JobId})
    detailClient.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", detailClient.Instance})
    detailClient.EntityData.Leafs.Append("numberof-sets", types.YLeaf{"NumberofSets", detailClient.NumberofSets})
    detailClient.EntityData.Leafs.Append("number-of-sessions", types.YLeaf{"NumberOfSessions", detailClient.NumberOfSessions})
    detailClient.EntityData.Leafs.Append("number-of-up-sessions", types.YLeaf{"NumberOfUpSessions", detailClient.NumberOfUpSessions})
    detailClient.EntityData.Leafs.Append("connected-at", types.YLeaf{"ConnectedAt", detailClient.ConnectedAt})
    detailClient.EntityData.Leafs.Append("is-notification-registered", types.YLeaf{"IsNotificationRegistered", detailClient.IsNotificationRegistered})

    detailClient.EntityData.YListKeys = []string {"Id"}

    return &(detailClient.EntityData)
}

// TcpNsr_Nodes_Node_Client_BriefClients
// Information about TCP NSR Client
type TcpNsr_Nodes_Node_Client_BriefClients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about NSR Client. The type is slice of
    // TcpNsr_Nodes_Node_Client_BriefClients_BriefClient.
    BriefClient []*TcpNsr_Nodes_Node_Client_BriefClients_BriefClient
}

func (briefClients *TcpNsr_Nodes_Node_Client_BriefClients) GetEntityData() *types.CommonEntityData {
    briefClients.EntityData.YFilter = briefClients.YFilter
    briefClients.EntityData.YangName = "brief-clients"
    briefClients.EntityData.BundleName = "cisco_ios_xr"
    briefClients.EntityData.ParentYangName = "client"
    briefClients.EntityData.SegmentPath = "brief-clients"
    briefClients.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/client/" + briefClients.EntityData.SegmentPath
    briefClients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefClients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefClients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefClients.EntityData.Children = types.NewOrderedMap()
    briefClients.EntityData.Children.Append("brief-client", types.YChild{"BriefClient", nil})
    for i := range briefClients.BriefClient {
        briefClients.EntityData.Children.Append(types.GetSegmentPath(briefClients.BriefClient[i]), types.YChild{"BriefClient", briefClients.BriefClient[i]})
    }
    briefClients.EntityData.Leafs = types.NewOrderedMap()

    briefClients.EntityData.YListKeys = []string {}

    return &(briefClients.EntityData)
}

// TcpNsr_Nodes_Node_Client_BriefClients_BriefClient
// Brief information about NSR Client
type TcpNsr_Nodes_Node_Client_BriefClients_BriefClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (briefClient *TcpNsr_Nodes_Node_Client_BriefClients_BriefClient) GetEntityData() *types.CommonEntityData {
    briefClient.EntityData.YFilter = briefClient.YFilter
    briefClient.EntityData.YangName = "brief-client"
    briefClient.EntityData.BundleName = "cisco_ios_xr"
    briefClient.EntityData.ParentYangName = "brief-clients"
    briefClient.EntityData.SegmentPath = "brief-client" + types.AddKeyToken(briefClient.Id, "id")
    briefClient.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/client/brief-clients/" + briefClient.EntityData.SegmentPath
    briefClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefClient.EntityData.Children = types.NewOrderedMap()
    briefClient.EntityData.Leafs = types.NewOrderedMap()
    briefClient.EntityData.Leafs.Append("id", types.YLeaf{"Id", briefClient.Id})
    briefClient.EntityData.Leafs.Append("ccb", types.YLeaf{"Ccb", briefClient.Ccb})
    briefClient.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", briefClient.Pid})
    briefClient.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", briefClient.ProcessName})
    briefClient.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", briefClient.JobId})
    briefClient.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", briefClient.Instance})
    briefClient.EntityData.Leafs.Append("numberof-sets", types.YLeaf{"NumberofSets", briefClient.NumberofSets})
    briefClient.EntityData.Leafs.Append("number-of-sessions", types.YLeaf{"NumberOfSessions", briefClient.NumberOfSessions})
    briefClient.EntityData.Leafs.Append("number-of-up-sessions", types.YLeaf{"NumberOfUpSessions", briefClient.NumberOfUpSessions})

    briefClient.EntityData.YListKeys = []string {"Id"}

    return &(briefClient.EntityData)
}

// TcpNsr_Nodes_Node_SessionSet
// Information about TCP NSR Session Sets
type TcpNsr_Nodes_Node_SessionSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table about TCP NSR Session Sets details.
    DetailSets TcpNsr_Nodes_Node_SessionSet_DetailSets

    // Information about TCP NSR Session Sets.
    BriefSets TcpNsr_Nodes_Node_SessionSet_BriefSets
}

func (sessionSet *TcpNsr_Nodes_Node_SessionSet) GetEntityData() *types.CommonEntityData {
    sessionSet.EntityData.YFilter = sessionSet.YFilter
    sessionSet.EntityData.YangName = "session-set"
    sessionSet.EntityData.BundleName = "cisco_ios_xr"
    sessionSet.EntityData.ParentYangName = "node"
    sessionSet.EntityData.SegmentPath = "session-set"
    sessionSet.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/" + sessionSet.EntityData.SegmentPath
    sessionSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionSet.EntityData.Children = types.NewOrderedMap()
    sessionSet.EntityData.Children.Append("detail-sets", types.YChild{"DetailSets", &sessionSet.DetailSets})
    sessionSet.EntityData.Children.Append("brief-sets", types.YChild{"BriefSets", &sessionSet.BriefSets})
    sessionSet.EntityData.Leafs = types.NewOrderedMap()

    sessionSet.EntityData.YListKeys = []string {}

    return &(sessionSet.EntityData)
}

// TcpNsr_Nodes_Node_SessionSet_DetailSets
// Table about TCP NSR Session Sets details
type TcpNsr_Nodes_Node_SessionSet_DetailSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // showing detailed information of NSR Session Sets. The type is slice of
    // TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet.
    DetailSet []*TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet
}

func (detailSets *TcpNsr_Nodes_Node_SessionSet_DetailSets) GetEntityData() *types.CommonEntityData {
    detailSets.EntityData.YFilter = detailSets.YFilter
    detailSets.EntityData.YangName = "detail-sets"
    detailSets.EntityData.BundleName = "cisco_ios_xr"
    detailSets.EntityData.ParentYangName = "session-set"
    detailSets.EntityData.SegmentPath = "detail-sets"
    detailSets.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session-set/" + detailSets.EntityData.SegmentPath
    detailSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailSets.EntityData.Children = types.NewOrderedMap()
    detailSets.EntityData.Children.Append("detail-set", types.YChild{"DetailSet", nil})
    for i := range detailSets.DetailSet {
        detailSets.EntityData.Children.Append(types.GetSegmentPath(detailSets.DetailSet[i]), types.YChild{"DetailSet", detailSets.DetailSet[i]})
    }
    detailSets.EntityData.Leafs = types.NewOrderedMap()

    detailSets.EntityData.YListKeys = []string {}

    return &(detailSets.EntityData)
}

// TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet
// showing detailed information of NSR Session
// Sets
type TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (detailSet *TcpNsr_Nodes_Node_SessionSet_DetailSets_DetailSet) GetEntityData() *types.CommonEntityData {
    detailSet.EntityData.YFilter = detailSet.YFilter
    detailSet.EntityData.YangName = "detail-set"
    detailSet.EntityData.BundleName = "cisco_ios_xr"
    detailSet.EntityData.ParentYangName = "detail-sets"
    detailSet.EntityData.SegmentPath = "detail-set" + types.AddKeyToken(detailSet.Id, "id")
    detailSet.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session-set/detail-sets/" + detailSet.EntityData.SegmentPath
    detailSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    detailSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    detailSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    detailSet.EntityData.Children = types.NewOrderedMap()
    detailSet.EntityData.Leafs = types.NewOrderedMap()
    detailSet.EntityData.Leafs.Append("id", types.YLeaf{"Id", detailSet.Id})
    detailSet.EntityData.Leafs.Append("sscb", types.YLeaf{"Sscb", detailSet.Sscb})
    detailSet.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", detailSet.Pid})
    detailSet.EntityData.Leafs.Append("set-id", types.YLeaf{"SetId", detailSet.SetId})
    detailSet.EntityData.Leafs.Append("sso-role", types.YLeaf{"SsoRole", detailSet.SsoRole})
    detailSet.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", detailSet.Mode})
    detailSet.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", detailSet.AddressFamily})
    detailSet.EntityData.Leafs.Append("well-known-port", types.YLeaf{"WellKnownPort", detailSet.WellKnownPort})
    detailSet.EntityData.Leafs.Append("local-node", types.YLeaf{"LocalNode", detailSet.LocalNode})
    detailSet.EntityData.Leafs.Append("local-instance", types.YLeaf{"LocalInstance", detailSet.LocalInstance})
    detailSet.EntityData.Leafs.Append("protect-node", types.YLeaf{"ProtectNode", detailSet.ProtectNode})
    detailSet.EntityData.Leafs.Append("protect-instance", types.YLeaf{"ProtectInstance", detailSet.ProtectInstance})
    detailSet.EntityData.Leafs.Append("number-of-sessions", types.YLeaf{"NumberOfSessions", detailSet.NumberOfSessions})
    detailSet.EntityData.Leafs.Append("number-of-synced-sessions-up-stream", types.YLeaf{"NumberOfSyncedSessionsUpStream", detailSet.NumberOfSyncedSessionsUpStream})
    detailSet.EntityData.Leafs.Append("number-of-synced-sessions-down-stream", types.YLeaf{"NumberOfSyncedSessionsDownStream", detailSet.NumberOfSyncedSessionsDownStream})
    detailSet.EntityData.Leafs.Append("is-init-sync-in-progress", types.YLeaf{"IsInitSyncInProgress", detailSet.IsInitSyncInProgress})
    detailSet.EntityData.Leafs.Append("is-init-sync-second-phase", types.YLeaf{"IsInitSyncSecondPhase", detailSet.IsInitSyncSecondPhase})
    detailSet.EntityData.Leafs.Append("sequence-number-of-init-sync", types.YLeaf{"SequenceNumberOfInitSync", detailSet.SequenceNumberOfInitSync})
    detailSet.EntityData.Leafs.Append("init-sync-timer", types.YLeaf{"InitSyncTimer", detailSet.InitSyncTimer})
    detailSet.EntityData.Leafs.Append("total-number-of-init-sync-sessions", types.YLeaf{"TotalNumberOfInitSyncSessions", detailSet.TotalNumberOfInitSyncSessions})
    detailSet.EntityData.Leafs.Append("number-of-init-synced-sessions", types.YLeaf{"NumberOfInitSyncedSessions", detailSet.NumberOfInitSyncedSessions})
    detailSet.EntityData.Leafs.Append("number-of-sessions-init-sync-failed", types.YLeaf{"NumberOfSessionsInitSyncFailed", detailSet.NumberOfSessionsInitSyncFailed})
    detailSet.EntityData.Leafs.Append("init-sync-error", types.YLeaf{"InitSyncError", detailSet.InitSyncError})
    detailSet.EntityData.Leafs.Append("is-init-sync-error-local", types.YLeaf{"IsInitSyncErrorLocal", detailSet.IsInitSyncErrorLocal})
    detailSet.EntityData.Leafs.Append("init-sync-start-time", types.YLeaf{"InitSyncStartTime", detailSet.InitSyncStartTime})
    detailSet.EntityData.Leafs.Append("init-sync-end-time", types.YLeaf{"InitSyncEndTime", detailSet.InitSyncEndTime})
    detailSet.EntityData.Leafs.Append("is-sscb-init-sync-ready", types.YLeaf{"IsSscbInitSyncReady", detailSet.IsSscbInitSyncReady})
    detailSet.EntityData.Leafs.Append("init-sync-ready-start-time", types.YLeaf{"InitSyncReadyStartTime", detailSet.InitSyncReadyStartTime})
    detailSet.EntityData.Leafs.Append("init-sync-ready-end-time", types.YLeaf{"InitSyncReadyEndTime", detailSet.InitSyncReadyEndTime})
    detailSet.EntityData.Leafs.Append("nsr-reset-time", types.YLeaf{"NsrResetTime", detailSet.NsrResetTime})
    detailSet.EntityData.Leafs.Append("is-audit-in-progress", types.YLeaf{"IsAuditInProgress", detailSet.IsAuditInProgress})
    detailSet.EntityData.Leafs.Append("audit-seq-number", types.YLeaf{"AuditSeqNumber", detailSet.AuditSeqNumber})
    detailSet.EntityData.Leafs.Append("audit-start-time", types.YLeaf{"AuditStartTime", detailSet.AuditStartTime})
    detailSet.EntityData.Leafs.Append("audit-end-time", types.YLeaf{"AuditEndTime", detailSet.AuditEndTime})

    detailSet.EntityData.YListKeys = []string {"Id"}

    return &(detailSet.EntityData)
}

// TcpNsr_Nodes_Node_SessionSet_BriefSets
// Information about TCP NSR Session Sets
type TcpNsr_Nodes_Node_SessionSet_BriefSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Brief information about NSR Session Sets. The type is slice of
    // TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet.
    BriefSet []*TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet
}

func (briefSets *TcpNsr_Nodes_Node_SessionSet_BriefSets) GetEntityData() *types.CommonEntityData {
    briefSets.EntityData.YFilter = briefSets.YFilter
    briefSets.EntityData.YangName = "brief-sets"
    briefSets.EntityData.BundleName = "cisco_ios_xr"
    briefSets.EntityData.ParentYangName = "session-set"
    briefSets.EntityData.SegmentPath = "brief-sets"
    briefSets.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session-set/" + briefSets.EntityData.SegmentPath
    briefSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefSets.EntityData.Children = types.NewOrderedMap()
    briefSets.EntityData.Children.Append("brief-set", types.YChild{"BriefSet", nil})
    for i := range briefSets.BriefSet {
        briefSets.EntityData.Children.Append(types.GetSegmentPath(briefSets.BriefSet[i]), types.YChild{"BriefSet", briefSets.BriefSet[i]})
    }
    briefSets.EntityData.Leafs = types.NewOrderedMap()

    briefSets.EntityData.YListKeys = []string {}

    return &(briefSets.EntityData)
}

// TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet
// Brief information about NSR Session Sets
type TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (briefSet *TcpNsr_Nodes_Node_SessionSet_BriefSets_BriefSet) GetEntityData() *types.CommonEntityData {
    briefSet.EntityData.YFilter = briefSet.YFilter
    briefSet.EntityData.YangName = "brief-set"
    briefSet.EntityData.BundleName = "cisco_ios_xr"
    briefSet.EntityData.ParentYangName = "brief-sets"
    briefSet.EntityData.SegmentPath = "brief-set" + types.AddKeyToken(briefSet.Id, "id")
    briefSet.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/session-set/brief-sets/" + briefSet.EntityData.SegmentPath
    briefSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    briefSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    briefSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    briefSet.EntityData.Children = types.NewOrderedMap()
    briefSet.EntityData.Leafs = types.NewOrderedMap()
    briefSet.EntityData.Leafs.Append("id", types.YLeaf{"Id", briefSet.Id})
    briefSet.EntityData.Leafs.Append("sscb", types.YLeaf{"Sscb", briefSet.Sscb})
    briefSet.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", briefSet.Pid})
    briefSet.EntityData.Leafs.Append("client-name", types.YLeaf{"ClientName", briefSet.ClientName})
    briefSet.EntityData.Leafs.Append("client-instance", types.YLeaf{"ClientInstance", briefSet.ClientInstance})
    briefSet.EntityData.Leafs.Append("set-id", types.YLeaf{"SetId", briefSet.SetId})
    briefSet.EntityData.Leafs.Append("sso-role", types.YLeaf{"SsoRole", briefSet.SsoRole})
    briefSet.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", briefSet.Mode})
    briefSet.EntityData.Leafs.Append("address-family", types.YLeaf{"AddressFamily", briefSet.AddressFamily})
    briefSet.EntityData.Leafs.Append("well-known-port", types.YLeaf{"WellKnownPort", briefSet.WellKnownPort})
    briefSet.EntityData.Leafs.Append("local-node", types.YLeaf{"LocalNode", briefSet.LocalNode})
    briefSet.EntityData.Leafs.Append("local-instance", types.YLeaf{"LocalInstance", briefSet.LocalInstance})
    briefSet.EntityData.Leafs.Append("protect-node", types.YLeaf{"ProtectNode", briefSet.ProtectNode})
    briefSet.EntityData.Leafs.Append("protect-instance", types.YLeaf{"ProtectInstance", briefSet.ProtectInstance})
    briefSet.EntityData.Leafs.Append("number-of-sessions", types.YLeaf{"NumberOfSessions", briefSet.NumberOfSessions})
    briefSet.EntityData.Leafs.Append("number-of-synced-sessions-up-stream", types.YLeaf{"NumberOfSyncedSessionsUpStream", briefSet.NumberOfSyncedSessionsUpStream})
    briefSet.EntityData.Leafs.Append("number-of-synced-sessions-down-stream", types.YLeaf{"NumberOfSyncedSessionsDownStream", briefSet.NumberOfSyncedSessionsDownStream})
    briefSet.EntityData.Leafs.Append("is-init-sync-in-progress", types.YLeaf{"IsInitSyncInProgress", briefSet.IsInitSyncInProgress})
    briefSet.EntityData.Leafs.Append("is-sscb-init-sync-ready", types.YLeaf{"IsSscbInitSyncReady", briefSet.IsSscbInitSyncReady})

    briefSet.EntityData.YListKeys = []string {"Id"}

    return &(briefSet.EntityData)
}

// TcpNsr_Nodes_Node_Statistics
// Statis Information about TCP NSR connections
type TcpNsr_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
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

func (statistics *TcpNsr_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Children.Append("summary", types.YChild{"Summary", &statistics.Summary})
    statistics.EntityData.Children.Append("statistic-clients", types.YChild{"StatisticClients", &statistics.StatisticClients})
    statistics.EntityData.Children.Append("statistic-sets", types.YChild{"StatisticSets", &statistics.StatisticSets})
    statistics.EntityData.Children.Append("statistic-sessions", types.YChild{"StatisticSessions", &statistics.StatisticSessions})
    statistics.EntityData.Leafs = types.NewOrderedMap()

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary
// Summary statistics across all NSR connections
type TcpNsr_Nodes_Node_Statistics_Summary struct {
    EntityData types.CommonEntityData
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
    NotificationStatistic []*TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic
}

func (summary *TcpNsr_Nodes_Node_Statistics_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "statistics"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/" + summary.EntityData.SegmentPath
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = types.NewOrderedMap()
    summary.EntityData.Children.Append("snd-counters", types.YChild{"SndCounters", &summary.SndCounters})
    summary.EntityData.Children.Append("audit-counters", types.YChild{"AuditCounters", &summary.AuditCounters})
    summary.EntityData.Children.Append("notification-statistic", types.YChild{"NotificationStatistic", nil})
    for i := range summary.NotificationStatistic {
        types.SetYListKey(summary.NotificationStatistic[i], i)
        summary.EntityData.Children.Append(types.GetSegmentPath(summary.NotificationStatistic[i]), types.YChild{"NotificationStatistic", summary.NotificationStatistic[i]})
    }
    summary.EntityData.Leafs = types.NewOrderedMap()
    summary.EntityData.Leafs.Append("last-cleared-time", types.YLeaf{"LastClearedTime", summary.LastClearedTime})
    summary.EntityData.Leafs.Append("number-of-connected-clients", types.YLeaf{"NumberOfConnectedClients", summary.NumberOfConnectedClients})
    summary.EntityData.Leafs.Append("number-of-disconnected-clients", types.YLeaf{"NumberOfDisconnectedClients", summary.NumberOfDisconnectedClients})
    summary.EntityData.Leafs.Append("number-of-current-clients", types.YLeaf{"NumberOfCurrentClients", summary.NumberOfCurrentClients})
    summary.EntityData.Leafs.Append("number-of-created-session-sets", types.YLeaf{"NumberOfCreatedSessionSets", summary.NumberOfCreatedSessionSets})
    summary.EntityData.Leafs.Append("number-of-destroyed-session-sets", types.YLeaf{"NumberOfDestroyedSessionSets", summary.NumberOfDestroyedSessionSets})
    summary.EntityData.Leafs.Append("number-of-current-session-sets", types.YLeaf{"NumberOfCurrentSessionSets", summary.NumberOfCurrentSessionSets})
    summary.EntityData.Leafs.Append("number-of-added-sessions", types.YLeaf{"NumberOfAddedSessions", summary.NumberOfAddedSessions})
    summary.EntityData.Leafs.Append("number-of-deleted-sessions", types.YLeaf{"NumberOfDeletedSessions", summary.NumberOfDeletedSessions})
    summary.EntityData.Leafs.Append("number-of-current-sessions", types.YLeaf{"NumberOfCurrentSessions", summary.NumberOfCurrentSessions})
    summary.EntityData.Leafs.Append("number-of-partner-node", types.YLeaf{"NumberOfPartnerNode", summary.NumberOfPartnerNode})
    summary.EntityData.Leafs.Append("number-of-attempted-init-sync", types.YLeaf{"NumberOfAttemptedInitSync", summary.NumberOfAttemptedInitSync})
    summary.EntityData.Leafs.Append("number-of-succeeded-init-sync", types.YLeaf{"NumberOfSucceededInitSync", summary.NumberOfSucceededInitSync})
    summary.EntityData.Leafs.Append("number-of-failed-init-sync", types.YLeaf{"NumberOfFailedInitSync", summary.NumberOfFailedInitSync})
    summary.EntityData.Leafs.Append("number-of-held-packets", types.YLeaf{"NumberOfHeldPackets", summary.NumberOfHeldPackets})
    summary.EntityData.Leafs.Append("number-of-held-but-dropped-packets", types.YLeaf{"NumberOfHeldButDroppedPackets", summary.NumberOfHeldButDroppedPackets})
    summary.EntityData.Leafs.Append("number-of-held-internal-acks", types.YLeaf{"NumberOfHeldInternalAcks", summary.NumberOfHeldInternalAcks})
    summary.EntityData.Leafs.Append("number-of-held-but-dropped-internal-acks", types.YLeaf{"NumberOfHeldButDroppedInternalAcks", summary.NumberOfHeldButDroppedInternalAcks})
    summary.EntityData.Leafs.Append("number-of-sent-internal-acks", types.YLeaf{"NumberOfSentInternalAcks", summary.NumberOfSentInternalAcks})
    summary.EntityData.Leafs.Append("number-of-received-internal-acks", types.YLeaf{"NumberOfReceivedInternalAcks", summary.NumberOfReceivedInternalAcks})
    summary.EntityData.Leafs.Append("number-of-qad-receive-messages-drops", types.YLeaf{"NumberOfQadReceiveMessagesDrops", summary.NumberOfQadReceiveMessagesDrops})
    summary.EntityData.Leafs.Append("number-of-qad-receive-messages-unknowns", types.YLeaf{"NumberOfQadReceiveMessagesUnknowns", summary.NumberOfQadReceiveMessagesUnknowns})
    summary.EntityData.Leafs.Append("number-of-qad-receive-messages-accepts", types.YLeaf{"NumberOfQadReceiveMessagesAccepts", summary.NumberOfQadReceiveMessagesAccepts})
    summary.EntityData.Leafs.Append("number-of-qad-stale-receive-messages-drops", types.YLeaf{"NumberOfQadStaleReceiveMessagesDrops", summary.NumberOfQadStaleReceiveMessagesDrops})
    summary.EntityData.Leafs.Append("number-of-qad-transfer-message-sent", types.YLeaf{"NumberOfQadTransferMessageSent", summary.NumberOfQadTransferMessageSent})
    summary.EntityData.Leafs.Append("number-of-qad-transfer-message-drops", types.YLeaf{"NumberOfQadTransferMessageDrops", summary.NumberOfQadTransferMessageDrops})
    summary.EntityData.Leafs.Append("number-of-internal-ack-drops-no-pcb", types.YLeaf{"NumberOfInternalAckDropsNoPcb", summary.NumberOfInternalAckDropsNoPcb})
    summary.EntityData.Leafs.Append("number-of-internal-ack-drops-no-scbdp", types.YLeaf{"NumberOfInternalAckDropsNoScbdp", summary.NumberOfInternalAckDropsNoScbdp})
    summary.EntityData.Leafs.Append("internal-ack-drops-not-replicated", types.YLeaf{"InternalAckDropsNotReplicated", summary.InternalAckDropsNotReplicated})
    summary.EntityData.Leafs.Append("internal-ack-drops-initsync-first-phase", types.YLeaf{"InternalAckDropsInitsyncFirstPhase", summary.InternalAckDropsInitsyncFirstPhase})
    summary.EntityData.Leafs.Append("internal-ack-drops-stale", types.YLeaf{"InternalAckDropsStale", summary.InternalAckDropsStale})
    summary.EntityData.Leafs.Append("internal-ack-drops-immediate-match", types.YLeaf{"InternalAckDropsImmediateMatch", summary.InternalAckDropsImmediateMatch})
    summary.EntityData.Leafs.Append("held-packet-drops", types.YLeaf{"HeldPacketDrops", summary.HeldPacketDrops})

    summary.EntityData.YListKeys = []string {}

    return &(summary.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_SndCounters
// Aggregate Send path counters
type TcpNsr_Nodes_Node_Statistics_Summary_SndCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Common send path counters.
    Common TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common

    // Aggregate only send path counters.
    AggrOnly TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly
}

func (sndCounters *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters) GetEntityData() *types.CommonEntityData {
    sndCounters.EntityData.YFilter = sndCounters.YFilter
    sndCounters.EntityData.YangName = "snd-counters"
    sndCounters.EntityData.BundleName = "cisco_ios_xr"
    sndCounters.EntityData.ParentYangName = "summary"
    sndCounters.EntityData.SegmentPath = "snd-counters"
    sndCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/" + sndCounters.EntityData.SegmentPath
    sndCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sndCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sndCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sndCounters.EntityData.Children = types.NewOrderedMap()
    sndCounters.EntityData.Children.Append("common", types.YChild{"Common", &sndCounters.Common})
    sndCounters.EntityData.Children.Append("aggr-only", types.YChild{"AggrOnly", &sndCounters.AggrOnly})
    sndCounters.EntityData.Leafs = types.NewOrderedMap()

    sndCounters.EntityData.YListKeys = []string {}

    return &(sndCounters.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common
// Common send path counters
type TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common struct {
    EntityData types.CommonEntityData
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

func (common *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "snd-counters"
    common.EntityData.SegmentPath = "common"
    common.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/snd-counters/" + common.EntityData.SegmentPath
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = types.NewOrderedMap()
    common.EntityData.Leafs = types.NewOrderedMap()
    common.EntityData.Leafs.Append("data-xfer-send", types.YLeaf{"DataXferSend", common.DataXferSend})
    common.EntityData.Leafs.Append("data-xfer-send-total", types.YLeaf{"DataXferSendTotal", common.DataXferSendTotal})
    common.EntityData.Leafs.Append("data-xfer-send-drop", types.YLeaf{"DataXferSendDrop", common.DataXferSendDrop})
    common.EntityData.Leafs.Append("data-xfer-send-iov-alloc", types.YLeaf{"DataXferSendIovAlloc", common.DataXferSendIovAlloc})
    common.EntityData.Leafs.Append("data-xfer-rcv", types.YLeaf{"DataXferRcv", common.DataXferRcv})
    common.EntityData.Leafs.Append("data-xfer-rcv-success", types.YLeaf{"DataXferRcvSuccess", common.DataXferRcvSuccess})
    common.EntityData.Leafs.Append("data-xfer-rcv-fail-buffer-trim", types.YLeaf{"DataXferRcvFailBufferTrim", common.DataXferRcvFailBufferTrim})
    common.EntityData.Leafs.Append("data-xfer-rcv-fail-snd-una-out-of-sync", types.YLeaf{"DataXferRcvFailSndUnaOutOfSync", common.DataXferRcvFailSndUnaOutOfSync})
    common.EntityData.Leafs.Append("seg-instr-send", types.YLeaf{"SegInstrSend", common.SegInstrSend})
    common.EntityData.Leafs.Append("seg-instr-send-units", types.YLeaf{"SegInstrSendUnits", common.SegInstrSendUnits})
    common.EntityData.Leafs.Append("seg-instr-send-drop", types.YLeaf{"SegInstrSendDrop", common.SegInstrSendDrop})
    common.EntityData.Leafs.Append("seg-instr-rcv", types.YLeaf{"SegInstrRcv", common.SegInstrRcv})
    common.EntityData.Leafs.Append("seg-instr-rcv-success", types.YLeaf{"SegInstrRcvSuccess", common.SegInstrRcvSuccess})
    common.EntityData.Leafs.Append("seg-instr-rcv-fail-buffer-trim", types.YLeaf{"SegInstrRcvFailBufferTrim", common.SegInstrRcvFailBufferTrim})
    common.EntityData.Leafs.Append("seg-instr-rcv-fail-tcp-process", types.YLeaf{"SegInstrRcvFailTcpProcess", common.SegInstrRcvFailTcpProcess})
    common.EntityData.Leafs.Append("nack-send", types.YLeaf{"NackSend", common.NackSend})
    common.EntityData.Leafs.Append("nack-send-drop", types.YLeaf{"NackSendDrop", common.NackSendDrop})
    common.EntityData.Leafs.Append("nack-rcv", types.YLeaf{"NackRcv", common.NackRcv})
    common.EntityData.Leafs.Append("nack-rcv-success", types.YLeaf{"NackRcvSuccess", common.NackRcvSuccess})
    common.EntityData.Leafs.Append("nack-rcv-fail-data-send", types.YLeaf{"NackRcvFailDataSend", common.NackRcvFailDataSend})
    common.EntityData.Leafs.Append("cleanup-send", types.YLeaf{"CleanupSend", common.CleanupSend})
    common.EntityData.Leafs.Append("cleanup-send-drop", types.YLeaf{"CleanupSendDrop", common.CleanupSendDrop})
    common.EntityData.Leafs.Append("cleanup-rcv", types.YLeaf{"CleanupRcv", common.CleanupRcv})
    common.EntityData.Leafs.Append("cleanup-rcv-success", types.YLeaf{"CleanupRcvSuccess", common.CleanupRcvSuccess})
    common.EntityData.Leafs.Append("cleanup-rcv-fail-buffer-trim", types.YLeaf{"CleanupRcvFailBufferTrim", common.CleanupRcvFailBufferTrim})

    common.EntityData.YListKeys = []string {}

    return &(common.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly
// Aggregate only send path counters
type TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly struct {
    EntityData types.CommonEntityData
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

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_SndCounters_AggrOnly) GetEntityData() *types.CommonEntityData {
    aggrOnly.EntityData.YFilter = aggrOnly.YFilter
    aggrOnly.EntityData.YangName = "aggr-only"
    aggrOnly.EntityData.BundleName = "cisco_ios_xr"
    aggrOnly.EntityData.ParentYangName = "snd-counters"
    aggrOnly.EntityData.SegmentPath = "aggr-only"
    aggrOnly.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/snd-counters/" + aggrOnly.EntityData.SegmentPath
    aggrOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggrOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggrOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggrOnly.EntityData.Children = types.NewOrderedMap()
    aggrOnly.EntityData.Leafs = types.NewOrderedMap()
    aggrOnly.EntityData.Leafs.Append("data-xfer-rcv-drop-no-pcb", types.YLeaf{"DataXferRcvDropNoPcb", aggrOnly.DataXferRcvDropNoPcb})
    aggrOnly.EntityData.Leafs.Append("data-xfer-rcv-drop-no-scb-dp", types.YLeaf{"DataXferRcvDropNoScbDp", aggrOnly.DataXferRcvDropNoScbDp})
    aggrOnly.EntityData.Leafs.Append("seg-instr-rcv-drop-no-pcb", types.YLeaf{"SegInstrRcvDropNoPcb", aggrOnly.SegInstrRcvDropNoPcb})
    aggrOnly.EntityData.Leafs.Append("seg-instr-rcv-drop-no-scb-dp", types.YLeaf{"SegInstrRcvDropNoScbDp", aggrOnly.SegInstrRcvDropNoScbDp})
    aggrOnly.EntityData.Leafs.Append("nack-rcv-drop-no-pcb", types.YLeaf{"NackRcvDropNoPcb", aggrOnly.NackRcvDropNoPcb})
    aggrOnly.EntityData.Leafs.Append("nack-rcv-drop-no-scb-dp", types.YLeaf{"NackRcvDropNoScbDp", aggrOnly.NackRcvDropNoScbDp})
    aggrOnly.EntityData.Leafs.Append("cleanup-rcv-drop-no-pcb", types.YLeaf{"CleanupRcvDropNoPcb", aggrOnly.CleanupRcvDropNoPcb})
    aggrOnly.EntityData.Leafs.Append("cleanup-rcv-drop-no-scb-dp", types.YLeaf{"CleanupRcvDropNoScbDp", aggrOnly.CleanupRcvDropNoScbDp})

    aggrOnly.EntityData.YListKeys = []string {}

    return &(aggrOnly.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters
// Aggregate Audit counters
type TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Common audit counters.
    Common TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common

    // Aggregate only audit counters.
    AggrOnly TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly
}

func (auditCounters *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters) GetEntityData() *types.CommonEntityData {
    auditCounters.EntityData.YFilter = auditCounters.YFilter
    auditCounters.EntityData.YangName = "audit-counters"
    auditCounters.EntityData.BundleName = "cisco_ios_xr"
    auditCounters.EntityData.ParentYangName = "summary"
    auditCounters.EntityData.SegmentPath = "audit-counters"
    auditCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/" + auditCounters.EntityData.SegmentPath
    auditCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    auditCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    auditCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    auditCounters.EntityData.Children = types.NewOrderedMap()
    auditCounters.EntityData.Children.Append("common", types.YChild{"Common", &auditCounters.Common})
    auditCounters.EntityData.Children.Append("aggr-only", types.YChild{"AggrOnly", &auditCounters.AggrOnly})
    auditCounters.EntityData.Leafs = types.NewOrderedMap()

    auditCounters.EntityData.YListKeys = []string {}

    return &(auditCounters.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common
// Common audit counters
type TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common struct {
    EntityData types.CommonEntityData
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

func (common *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_Common) GetEntityData() *types.CommonEntityData {
    common.EntityData.YFilter = common.YFilter
    common.EntityData.YangName = "common"
    common.EntityData.BundleName = "cisco_ios_xr"
    common.EntityData.ParentYangName = "audit-counters"
    common.EntityData.SegmentPath = "common"
    common.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/audit-counters/" + common.EntityData.SegmentPath
    common.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    common.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    common.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    common.EntityData.Children = types.NewOrderedMap()
    common.EntityData.Leafs = types.NewOrderedMap()
    common.EntityData.Leafs.Append("mark-session-set-send", types.YLeaf{"MarkSessionSetSend", common.MarkSessionSetSend})
    common.EntityData.Leafs.Append("mark-session-set-send-drop", types.YLeaf{"MarkSessionSetSendDrop", common.MarkSessionSetSendDrop})
    common.EntityData.Leafs.Append("mark-session-set-rcv", types.YLeaf{"MarkSessionSetRcv", common.MarkSessionSetRcv})
    common.EntityData.Leafs.Append("mark-session-set-rcv-drop", types.YLeaf{"MarkSessionSetRcvDrop", common.MarkSessionSetRcvDrop})
    common.EntityData.Leafs.Append("session-send", types.YLeaf{"SessionSend", common.SessionSend})
    common.EntityData.Leafs.Append("session-send-drop", types.YLeaf{"SessionSendDrop", common.SessionSendDrop})
    common.EntityData.Leafs.Append("session-rcv", types.YLeaf{"SessionRcv", common.SessionRcv})
    common.EntityData.Leafs.Append("session-rcv-drop", types.YLeaf{"SessionRcvDrop", common.SessionRcvDrop})
    common.EntityData.Leafs.Append("sweep-session-set-send", types.YLeaf{"SweepSessionSetSend", common.SweepSessionSetSend})
    common.EntityData.Leafs.Append("sweep-session-set-send-drop", types.YLeaf{"SweepSessionSetSendDrop", common.SweepSessionSetSendDrop})
    common.EntityData.Leafs.Append("sweep-session-set-rcv", types.YLeaf{"SweepSessionSetRcv", common.SweepSessionSetRcv})
    common.EntityData.Leafs.Append("sweep-session-set-rcv-drop", types.YLeaf{"SweepSessionSetRcvDrop", common.SweepSessionSetRcvDrop})
    common.EntityData.Leafs.Append("session-set-response-send", types.YLeaf{"SessionSetResponseSend", common.SessionSetResponseSend})
    common.EntityData.Leafs.Append("session-set-response-send-drop", types.YLeaf{"SessionSetResponseSendDrop", common.SessionSetResponseSendDrop})
    common.EntityData.Leafs.Append("session-set-response-rcv", types.YLeaf{"SessionSetResponseRcv", common.SessionSetResponseRcv})
    common.EntityData.Leafs.Append("session-set-response-rcv-drop", types.YLeaf{"SessionSetResponseRcvDrop", common.SessionSetResponseRcvDrop})
    common.EntityData.Leafs.Append("mark-session-set-ack-send", types.YLeaf{"MarkSessionSetAckSend", common.MarkSessionSetAckSend})
    common.EntityData.Leafs.Append("mark-session-set-ack-send-drop", types.YLeaf{"MarkSessionSetAckSendDrop", common.MarkSessionSetAckSendDrop})
    common.EntityData.Leafs.Append("mark-session-set-ack-rcv", types.YLeaf{"MarkSessionSetAckRcv", common.MarkSessionSetAckRcv})
    common.EntityData.Leafs.Append("mark-session-set-ack-rcv-drop", types.YLeaf{"MarkSessionSetAckRcvDrop", common.MarkSessionSetAckRcvDrop})
    common.EntityData.Leafs.Append("mark-session-set-nack-send", types.YLeaf{"MarkSessionSetNackSend", common.MarkSessionSetNackSend})
    common.EntityData.Leafs.Append("mark-session-set-nack-send-drop", types.YLeaf{"MarkSessionSetNackSendDrop", common.MarkSessionSetNackSendDrop})
    common.EntityData.Leafs.Append("mark-session-set-nack-rcv", types.YLeaf{"MarkSessionSetNackRcv", common.MarkSessionSetNackRcv})
    common.EntityData.Leafs.Append("mark-session-set-nack-rcv-drop", types.YLeaf{"MarkSessionSetNackRcvDrop", common.MarkSessionSetNackRcvDrop})
    common.EntityData.Leafs.Append("abort", types.YLeaf{"Abort", common.Abort})

    common.EntityData.YListKeys = []string {}

    return &(common.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly
// Aggregate only audit counters
type TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly struct {
    EntityData types.CommonEntityData
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

func (aggrOnly *TcpNsr_Nodes_Node_Statistics_Summary_AuditCounters_AggrOnly) GetEntityData() *types.CommonEntityData {
    aggrOnly.EntityData.YFilter = aggrOnly.YFilter
    aggrOnly.EntityData.YangName = "aggr-only"
    aggrOnly.EntityData.BundleName = "cisco_ios_xr"
    aggrOnly.EntityData.ParentYangName = "audit-counters"
    aggrOnly.EntityData.SegmentPath = "aggr-only"
    aggrOnly.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/audit-counters/" + aggrOnly.EntityData.SegmentPath
    aggrOnly.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggrOnly.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggrOnly.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggrOnly.EntityData.Children = types.NewOrderedMap()
    aggrOnly.EntityData.Leafs = types.NewOrderedMap()
    aggrOnly.EntityData.Leafs.Append("mark-session-set-rcv-drop-aggr", types.YLeaf{"MarkSessionSetRcvDropAggr", aggrOnly.MarkSessionSetRcvDropAggr})
    aggrOnly.EntityData.Leafs.Append("session-rcv-drop-aggr", types.YLeaf{"SessionRcvDropAggr", aggrOnly.SessionRcvDropAggr})
    aggrOnly.EntityData.Leafs.Append("sweep-session-set-rcv-drop-aggr", types.YLeaf{"SweepSessionSetRcvDropAggr", aggrOnly.SweepSessionSetRcvDropAggr})
    aggrOnly.EntityData.Leafs.Append("session-set-response-rcv-drop-aggr", types.YLeaf{"SessionSetResponseRcvDropAggr", aggrOnly.SessionSetResponseRcvDropAggr})
    aggrOnly.EntityData.Leafs.Append("mark-session-set-ack-rcv-drop-aggr", types.YLeaf{"MarkSessionSetAckRcvDropAggr", aggrOnly.MarkSessionSetAckRcvDropAggr})
    aggrOnly.EntityData.Leafs.Append("mark-session-set-nack-rcv-drop-aggr", types.YLeaf{"MarkSessionSetNackRcvDropAggr", aggrOnly.MarkSessionSetNackRcvDropAggr})

    aggrOnly.EntityData.YListKeys = []string {}

    return &(aggrOnly.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic
// Various types of notification stats
type TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_Summary_NotificationStatistic) GetEntityData() *types.CommonEntityData {
    notificationStatistic.EntityData.YFilter = notificationStatistic.YFilter
    notificationStatistic.EntityData.YangName = "notification-statistic"
    notificationStatistic.EntityData.BundleName = "cisco_ios_xr"
    notificationStatistic.EntityData.ParentYangName = "summary"
    notificationStatistic.EntityData.SegmentPath = "notification-statistic" + types.AddNoKeyToken(notificationStatistic)
    notificationStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/summary/" + notificationStatistic.EntityData.SegmentPath
    notificationStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationStatistic.EntityData.Children = types.NewOrderedMap()
    notificationStatistic.EntityData.Leafs = types.NewOrderedMap()
    notificationStatistic.EntityData.Leafs.Append("queued-count", types.YLeaf{"QueuedCount", notificationStatistic.QueuedCount})
    notificationStatistic.EntityData.Leafs.Append("failed-count", types.YLeaf{"FailedCount", notificationStatistic.FailedCount})
    notificationStatistic.EntityData.Leafs.Append("delivered-count", types.YLeaf{"DeliveredCount", notificationStatistic.DeliveredCount})
    notificationStatistic.EntityData.Leafs.Append("dropped-count", types.YLeaf{"DroppedCount", notificationStatistic.DroppedCount})

    notificationStatistic.EntityData.YListKeys = []string {}

    return &(notificationStatistic.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticClients
// Table listing NSR connections for which
// statistic information is provided
type TcpNsr_Nodes_Node_Statistics_StatisticClients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // showing statistic information of NSR Clients. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient.
    StatisticClient []*TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient
}

func (statisticClients *TcpNsr_Nodes_Node_Statistics_StatisticClients) GetEntityData() *types.CommonEntityData {
    statisticClients.EntityData.YFilter = statisticClients.YFilter
    statisticClients.EntityData.YangName = "statistic-clients"
    statisticClients.EntityData.BundleName = "cisco_ios_xr"
    statisticClients.EntityData.ParentYangName = "statistics"
    statisticClients.EntityData.SegmentPath = "statistic-clients"
    statisticClients.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/" + statisticClients.EntityData.SegmentPath
    statisticClients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticClients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticClients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticClients.EntityData.Children = types.NewOrderedMap()
    statisticClients.EntityData.Children.Append("statistic-client", types.YChild{"StatisticClient", nil})
    for i := range statisticClients.StatisticClient {
        statisticClients.EntityData.Children.Append(types.GetSegmentPath(statisticClients.StatisticClient[i]), types.YChild{"StatisticClient", statisticClients.StatisticClient[i]})
    }
    statisticClients.EntityData.Leafs = types.NewOrderedMap()

    statisticClients.EntityData.YListKeys = []string {}

    return &(statisticClients.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient
// showing statistic information of NSR Clients
type TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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
    NotificationStatistic []*TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic
}

func (statisticClient *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient) GetEntityData() *types.CommonEntityData {
    statisticClient.EntityData.YFilter = statisticClient.YFilter
    statisticClient.EntityData.YangName = "statistic-client"
    statisticClient.EntityData.BundleName = "cisco_ios_xr"
    statisticClient.EntityData.ParentYangName = "statistic-clients"
    statisticClient.EntityData.SegmentPath = "statistic-client" + types.AddKeyToken(statisticClient.Id, "id")
    statisticClient.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/statistic-clients/" + statisticClient.EntityData.SegmentPath
    statisticClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticClient.EntityData.Children = types.NewOrderedMap()
    statisticClient.EntityData.Children.Append("notification-statistic", types.YChild{"NotificationStatistic", nil})
    for i := range statisticClient.NotificationStatistic {
        types.SetYListKey(statisticClient.NotificationStatistic[i], i)
        statisticClient.EntityData.Children.Append(types.GetSegmentPath(statisticClient.NotificationStatistic[i]), types.YChild{"NotificationStatistic", statisticClient.NotificationStatistic[i]})
    }
    statisticClient.EntityData.Leafs = types.NewOrderedMap()
    statisticClient.EntityData.Leafs.Append("id", types.YLeaf{"Id", statisticClient.Id})
    statisticClient.EntityData.Leafs.Append("ccb", types.YLeaf{"Ccb", statisticClient.Ccb})
    statisticClient.EntityData.Leafs.Append("pid", types.YLeaf{"Pid", statisticClient.Pid})
    statisticClient.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", statisticClient.ProcessName})
    statisticClient.EntityData.Leafs.Append("job-id", types.YLeaf{"JobId", statisticClient.JobId})
    statisticClient.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", statisticClient.Instance})
    statisticClient.EntityData.Leafs.Append("connected-at", types.YLeaf{"ConnectedAt", statisticClient.ConnectedAt})
    statisticClient.EntityData.Leafs.Append("number-of-created-sscb", types.YLeaf{"NumberOfCreatedSscb", statisticClient.NumberOfCreatedSscb})
    statisticClient.EntityData.Leafs.Append("number-of-deleted-sscb", types.YLeaf{"NumberOfDeletedSscb", statisticClient.NumberOfDeletedSscb})
    statisticClient.EntityData.Leafs.Append("last-cleared-time", types.YLeaf{"LastClearedTime", statisticClient.LastClearedTime})

    statisticClient.EntityData.YListKeys = []string {"Id"}

    return &(statisticClient.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic
// Various types of notification stats
type TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (notificationStatistic *TcpNsr_Nodes_Node_Statistics_StatisticClients_StatisticClient_NotificationStatistic) GetEntityData() *types.CommonEntityData {
    notificationStatistic.EntityData.YFilter = notificationStatistic.YFilter
    notificationStatistic.EntityData.YangName = "notification-statistic"
    notificationStatistic.EntityData.BundleName = "cisco_ios_xr"
    notificationStatistic.EntityData.ParentYangName = "statistic-client"
    notificationStatistic.EntityData.SegmentPath = "notification-statistic" + types.AddNoKeyToken(notificationStatistic)
    notificationStatistic.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/statistic-clients/statistic-client/" + notificationStatistic.EntityData.SegmentPath
    notificationStatistic.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notificationStatistic.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notificationStatistic.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notificationStatistic.EntityData.Children = types.NewOrderedMap()
    notificationStatistic.EntityData.Leafs = types.NewOrderedMap()
    notificationStatistic.EntityData.Leafs.Append("queued-count", types.YLeaf{"QueuedCount", notificationStatistic.QueuedCount})
    notificationStatistic.EntityData.Leafs.Append("failed-count", types.YLeaf{"FailedCount", notificationStatistic.FailedCount})
    notificationStatistic.EntityData.Leafs.Append("delivered-count", types.YLeaf{"DeliveredCount", notificationStatistic.DeliveredCount})
    notificationStatistic.EntityData.Leafs.Append("dropped-count", types.YLeaf{"DroppedCount", notificationStatistic.DroppedCount})

    notificationStatistic.EntityData.YListKeys = []string {}

    return &(notificationStatistic.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticSets
// Table listing NSR connections for which
// statistic information is provided
type TcpNsr_Nodes_Node_Statistics_StatisticSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // showing statistic information of NSR Session Set. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet.
    StatisticSet []*TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet
}

func (statisticSets *TcpNsr_Nodes_Node_Statistics_StatisticSets) GetEntityData() *types.CommonEntityData {
    statisticSets.EntityData.YFilter = statisticSets.YFilter
    statisticSets.EntityData.YangName = "statistic-sets"
    statisticSets.EntityData.BundleName = "cisco_ios_xr"
    statisticSets.EntityData.ParentYangName = "statistics"
    statisticSets.EntityData.SegmentPath = "statistic-sets"
    statisticSets.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/" + statisticSets.EntityData.SegmentPath
    statisticSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticSets.EntityData.Children = types.NewOrderedMap()
    statisticSets.EntityData.Children.Append("statistic-set", types.YChild{"StatisticSet", nil})
    for i := range statisticSets.StatisticSet {
        statisticSets.EntityData.Children.Append(types.GetSegmentPath(statisticSets.StatisticSet[i]), types.YChild{"StatisticSet", statisticSets.StatisticSet[i]})
    }
    statisticSets.EntityData.Leafs = types.NewOrderedMap()

    statisticSets.EntityData.YListKeys = []string {}

    return &(statisticSets.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet
// showing statistic information of NSR Session
// Set
type TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (statisticSet *TcpNsr_Nodes_Node_Statistics_StatisticSets_StatisticSet) GetEntityData() *types.CommonEntityData {
    statisticSet.EntityData.YFilter = statisticSet.YFilter
    statisticSet.EntityData.YangName = "statistic-set"
    statisticSet.EntityData.BundleName = "cisco_ios_xr"
    statisticSet.EntityData.ParentYangName = "statistic-sets"
    statisticSet.EntityData.SegmentPath = "statistic-set" + types.AddKeyToken(statisticSet.Id, "id")
    statisticSet.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/statistic-sets/" + statisticSet.EntityData.SegmentPath
    statisticSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticSet.EntityData.Children = types.NewOrderedMap()
    statisticSet.EntityData.Leafs = types.NewOrderedMap()
    statisticSet.EntityData.Leafs.Append("id", types.YLeaf{"Id", statisticSet.Id})
    statisticSet.EntityData.Leafs.Append("sscb", types.YLeaf{"Sscb", statisticSet.Sscb})
    statisticSet.EntityData.Leafs.Append("set-id", types.YLeaf{"SetId", statisticSet.SetId})
    statisticSet.EntityData.Leafs.Append("number-of-attempted-init-sync", types.YLeaf{"NumberOfAttemptedInitSync", statisticSet.NumberOfAttemptedInitSync})
    statisticSet.EntityData.Leafs.Append("number-of-succeeded-init-sync", types.YLeaf{"NumberOfSucceededInitSync", statisticSet.NumberOfSucceededInitSync})
    statisticSet.EntityData.Leafs.Append("number-of-failed-init-sync", types.YLeaf{"NumberOfFailedInitSync", statisticSet.NumberOfFailedInitSync})
    statisticSet.EntityData.Leafs.Append("number-of-failover", types.YLeaf{"NumberOfFailover", statisticSet.NumberOfFailover})
    statisticSet.EntityData.Leafs.Append("number-of-nsr-resets", types.YLeaf{"NumberOfNsrResets", statisticSet.NumberOfNsrResets})
    statisticSet.EntityData.Leafs.Append("last-cleared-time", types.YLeaf{"LastClearedTime", statisticSet.LastClearedTime})

    statisticSet.EntityData.YListKeys = []string {"Id"}

    return &(statisticSet.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticSessions
// Table listing NSR connections for which
// statistic information is provided
type TcpNsr_Nodes_Node_Statistics_StatisticSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // showing statistic information of TCP connections. The type is slice of
    // TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession.
    StatisticSession []*TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession
}

func (statisticSessions *TcpNsr_Nodes_Node_Statistics_StatisticSessions) GetEntityData() *types.CommonEntityData {
    statisticSessions.EntityData.YFilter = statisticSessions.YFilter
    statisticSessions.EntityData.YangName = "statistic-sessions"
    statisticSessions.EntityData.BundleName = "cisco_ios_xr"
    statisticSessions.EntityData.ParentYangName = "statistics"
    statisticSessions.EntityData.SegmentPath = "statistic-sessions"
    statisticSessions.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/" + statisticSessions.EntityData.SegmentPath
    statisticSessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticSessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticSessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticSessions.EntityData.Children = types.NewOrderedMap()
    statisticSessions.EntityData.Children.Append("statistic-session", types.YChild{"StatisticSession", nil})
    for i := range statisticSessions.StatisticSession {
        statisticSessions.EntityData.Children.Append(types.GetSegmentPath(statisticSessions.StatisticSession[i]), types.YChild{"StatisticSession", statisticSessions.StatisticSession[i]})
    }
    statisticSessions.EntityData.Leafs = types.NewOrderedMap()

    statisticSessions.EntityData.YListKeys = []string {}

    return &(statisticSessions.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession
// showing statistic information of TCP
// connections
type TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (statisticSession *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession) GetEntityData() *types.CommonEntityData {
    statisticSession.EntityData.YFilter = statisticSession.YFilter
    statisticSession.EntityData.YangName = "statistic-session"
    statisticSession.EntityData.BundleName = "cisco_ios_xr"
    statisticSession.EntityData.ParentYangName = "statistic-sessions"
    statisticSession.EntityData.SegmentPath = "statistic-session" + types.AddKeyToken(statisticSession.Id, "id")
    statisticSession.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/statistic-sessions/" + statisticSession.EntityData.SegmentPath
    statisticSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statisticSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statisticSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statisticSession.EntityData.Children = types.NewOrderedMap()
    statisticSession.EntityData.Children.Append("snd-counters", types.YChild{"SndCounters", &statisticSession.SndCounters})
    statisticSession.EntityData.Leafs = types.NewOrderedMap()
    statisticSession.EntityData.Leafs.Append("id", types.YLeaf{"Id", statisticSession.Id})
    statisticSession.EntityData.Leafs.Append("pcb", types.YLeaf{"Pcb", statisticSession.Pcb})
    statisticSession.EntityData.Leafs.Append("number-of-times-nsr-up", types.YLeaf{"NumberOfTimesNsrUp", statisticSession.NumberOfTimesNsrUp})
    statisticSession.EntityData.Leafs.Append("number-of-timers-nsr-down", types.YLeaf{"NumberOfTimersNsrDown", statisticSession.NumberOfTimersNsrDown})
    statisticSession.EntityData.Leafs.Append("number-of-times-nsr-disabled", types.YLeaf{"NumberOfTimesNsrDisabled", statisticSession.NumberOfTimesNsrDisabled})
    statisticSession.EntityData.Leafs.Append("number-of-times-nsr-fail-over", types.YLeaf{"NumberOfTimesNsrFailOver", statisticSession.NumberOfTimesNsrFailOver})
    statisticSession.EntityData.Leafs.Append("internal-ack-drops-not-replicated", types.YLeaf{"InternalAckDropsNotReplicated", statisticSession.InternalAckDropsNotReplicated})
    statisticSession.EntityData.Leafs.Append("internal-ack-drops-initsync-first-phase", types.YLeaf{"InternalAckDropsInitsyncFirstPhase", statisticSession.InternalAckDropsInitsyncFirstPhase})
    statisticSession.EntityData.Leafs.Append("internal-ack-drops-stale", types.YLeaf{"InternalAckDropsStale", statisticSession.InternalAckDropsStale})
    statisticSession.EntityData.Leafs.Append("internal-ack-drops-immediate-match", types.YLeaf{"InternalAckDropsImmediateMatch", statisticSession.InternalAckDropsImmediateMatch})
    statisticSession.EntityData.Leafs.Append("last-cleared-time", types.YLeaf{"LastClearedTime", statisticSession.LastClearedTime})

    statisticSession.EntityData.YListKeys = []string {"Id"}

    return &(statisticSession.EntityData)
}

// TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters
// Send path counters for the PCB
type TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters struct {
    EntityData types.CommonEntityData
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

func (sndCounters *TcpNsr_Nodes_Node_Statistics_StatisticSessions_StatisticSession_SndCounters) GetEntityData() *types.CommonEntityData {
    sndCounters.EntityData.YFilter = sndCounters.YFilter
    sndCounters.EntityData.YangName = "snd-counters"
    sndCounters.EntityData.BundleName = "cisco_ios_xr"
    sndCounters.EntityData.ParentYangName = "statistic-session"
    sndCounters.EntityData.SegmentPath = "snd-counters"
    sndCounters.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-tcp-oper:tcp-nsr/nodes/node/statistics/statistic-sessions/statistic-session/" + sndCounters.EntityData.SegmentPath
    sndCounters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sndCounters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sndCounters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sndCounters.EntityData.Children = types.NewOrderedMap()
    sndCounters.EntityData.Leafs = types.NewOrderedMap()
    sndCounters.EntityData.Leafs.Append("data-xfer-send", types.YLeaf{"DataXferSend", sndCounters.DataXferSend})
    sndCounters.EntityData.Leafs.Append("data-xfer-send-total", types.YLeaf{"DataXferSendTotal", sndCounters.DataXferSendTotal})
    sndCounters.EntityData.Leafs.Append("data-xfer-send-drop", types.YLeaf{"DataXferSendDrop", sndCounters.DataXferSendDrop})
    sndCounters.EntityData.Leafs.Append("data-xfer-send-iov-alloc", types.YLeaf{"DataXferSendIovAlloc", sndCounters.DataXferSendIovAlloc})
    sndCounters.EntityData.Leafs.Append("data-xfer-rcv", types.YLeaf{"DataXferRcv", sndCounters.DataXferRcv})
    sndCounters.EntityData.Leafs.Append("data-xfer-rcv-success", types.YLeaf{"DataXferRcvSuccess", sndCounters.DataXferRcvSuccess})
    sndCounters.EntityData.Leafs.Append("data-xfer-rcv-fail-buffer-trim", types.YLeaf{"DataXferRcvFailBufferTrim", sndCounters.DataXferRcvFailBufferTrim})
    sndCounters.EntityData.Leafs.Append("data-xfer-rcv-fail-snd-una-out-of-sync", types.YLeaf{"DataXferRcvFailSndUnaOutOfSync", sndCounters.DataXferRcvFailSndUnaOutOfSync})
    sndCounters.EntityData.Leafs.Append("seg-instr-send", types.YLeaf{"SegInstrSend", sndCounters.SegInstrSend})
    sndCounters.EntityData.Leafs.Append("seg-instr-send-units", types.YLeaf{"SegInstrSendUnits", sndCounters.SegInstrSendUnits})
    sndCounters.EntityData.Leafs.Append("seg-instr-send-drop", types.YLeaf{"SegInstrSendDrop", sndCounters.SegInstrSendDrop})
    sndCounters.EntityData.Leafs.Append("seg-instr-rcv", types.YLeaf{"SegInstrRcv", sndCounters.SegInstrRcv})
    sndCounters.EntityData.Leafs.Append("seg-instr-rcv-success", types.YLeaf{"SegInstrRcvSuccess", sndCounters.SegInstrRcvSuccess})
    sndCounters.EntityData.Leafs.Append("seg-instr-rcv-fail-buffer-trim", types.YLeaf{"SegInstrRcvFailBufferTrim", sndCounters.SegInstrRcvFailBufferTrim})
    sndCounters.EntityData.Leafs.Append("seg-instr-rcv-fail-tcp-process", types.YLeaf{"SegInstrRcvFailTcpProcess", sndCounters.SegInstrRcvFailTcpProcess})
    sndCounters.EntityData.Leafs.Append("nack-send", types.YLeaf{"NackSend", sndCounters.NackSend})
    sndCounters.EntityData.Leafs.Append("nack-send-drop", types.YLeaf{"NackSendDrop", sndCounters.NackSendDrop})
    sndCounters.EntityData.Leafs.Append("nack-rcv", types.YLeaf{"NackRcv", sndCounters.NackRcv})
    sndCounters.EntityData.Leafs.Append("nack-rcv-success", types.YLeaf{"NackRcvSuccess", sndCounters.NackRcvSuccess})
    sndCounters.EntityData.Leafs.Append("nack-rcv-fail-data-send", types.YLeaf{"NackRcvFailDataSend", sndCounters.NackRcvFailDataSend})
    sndCounters.EntityData.Leafs.Append("cleanup-send", types.YLeaf{"CleanupSend", sndCounters.CleanupSend})
    sndCounters.EntityData.Leafs.Append("cleanup-send-drop", types.YLeaf{"CleanupSendDrop", sndCounters.CleanupSendDrop})
    sndCounters.EntityData.Leafs.Append("cleanup-rcv", types.YLeaf{"CleanupRcv", sndCounters.CleanupRcv})
    sndCounters.EntityData.Leafs.Append("cleanup-rcv-success", types.YLeaf{"CleanupRcvSuccess", sndCounters.CleanupRcvSuccess})
    sndCounters.EntityData.Leafs.Append("cleanup-rcv-fail-buffer-trim", types.YLeaf{"CleanupRcvFailBufferTrim", sndCounters.CleanupRcvFailBufferTrim})

    sndCounters.EntityData.YListKeys = []string {}

    return &(sndCounters.EntityData)
}

