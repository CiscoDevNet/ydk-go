// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_acl_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_acl_datatypes"))
}

// Ipv6PrefixMatchExactLength represents Ipv6 prefix match exact length
type Ipv6PrefixMatchExactLength string

const (
    // Prefix Length Exact match
    Ipv6PrefixMatchExactLength_match_exact_length Ipv6PrefixMatchExactLength = "match-exact-length"
)

// Ipv6AclGrantEnum represents Ipv6 acl grant enum
type Ipv6AclGrantEnum string

const (
    // Deny
    Ipv6AclGrantEnum_deny Ipv6AclGrantEnum = "deny"

    // Permit
    Ipv6AclGrantEnum_permit Ipv6AclGrantEnum = "permit"
)

// Ipv6AclLoggingEnum represents Ipv6 acl logging enum
type Ipv6AclLoggingEnum string

const (
    // Log matches against this entry
    Ipv6AclLoggingEnum_log Ipv6AclLoggingEnum = "log"

    // Log matches against this entry, including input
    // interface
    Ipv6AclLoggingEnum_log_input Ipv6AclLoggingEnum = "log-input"
)

// Ipv6AclProtocolNumber represents Ipv6 acl protocol number
type Ipv6AclProtocolNumber string

const (
    // Any IP protocol
    Ipv6AclProtocolNumber_ip Ipv6AclProtocolNumber = "ip"

    // Internet Control Message Protocol
    Ipv6AclProtocolNumber_icmp Ipv6AclProtocolNumber = "icmp"

    // Internet Gateway Message Protocol
    Ipv6AclProtocolNumber_igmp Ipv6AclProtocolNumber = "igmp"

    // IP in IP tunneling
    Ipv6AclProtocolNumber_ip_in_ip Ipv6AclProtocolNumber = "ip-in-ip"

    // Transport Control Protocol
    Ipv6AclProtocolNumber_tcp Ipv6AclProtocolNumber = "tcp"

    // Cisco's IGRP Routing Protocol
    Ipv6AclProtocolNumber_igrp Ipv6AclProtocolNumber = "igrp"

    // User Datagram Protocol
    Ipv6AclProtocolNumber_udp Ipv6AclProtocolNumber = "udp"

    // Cisco's GRE tunneling
    Ipv6AclProtocolNumber_gre Ipv6AclProtocolNumber = "gre"

    // Encapsulation Security Protocol
    Ipv6AclProtocolNumber_esp Ipv6AclProtocolNumber = "esp"

    // Authentication Header Protocol
    Ipv6AclProtocolNumber_ahp Ipv6AclProtocolNumber = "ahp"

    // Internet Control Message Protocol
    Ipv6AclProtocolNumber_icmpv6 Ipv6AclProtocolNumber = "icmpv6"

    // Cisco's EIGRP Routing Protocol
    Ipv6AclProtocolNumber_eigrp Ipv6AclProtocolNumber = "eigrp"

    // OSPF Routing Protocol
    Ipv6AclProtocolNumber_ospf Ipv6AclProtocolNumber = "ospf"

    // KA9Q NOS Compatible IP over IP tunneling
    Ipv6AclProtocolNumber_nos Ipv6AclProtocolNumber = "nos"

    // Protocol Independent Multicast
    Ipv6AclProtocolNumber_pim Ipv6AclProtocolNumber = "pim"

    // Payload Compression Protocol
    Ipv6AclProtocolNumber_pcp Ipv6AclProtocolNumber = "pcp"

    // Stream Control Transmission Protocol
    Ipv6AclProtocolNumber_sctp Ipv6AclProtocolNumber = "sctp"
)

// Ipv6AclIcmpTypeCodeEnum represents Ipv6 acl icmp type code enum
type Ipv6AclIcmpTypeCodeEnum string

const (
    // No route to destination
    Ipv6AclIcmpTypeCodeEnum_no_route_to_destination Ipv6AclIcmpTypeCodeEnum = "no-route-to-destination"

    // Administratively prohibited
    Ipv6AclIcmpTypeCodeEnum_administratively_prohibited Ipv6AclIcmpTypeCodeEnum = "administratively-prohibited"

    // Unreachable beyond scope of address
    Ipv6AclIcmpTypeCodeEnum_beyond_scope_of_source_address Ipv6AclIcmpTypeCodeEnum = "beyond-scope-of-source-address"

    // Host unreachable
    Ipv6AclIcmpTypeCodeEnum_host_unreachable Ipv6AclIcmpTypeCodeEnum = "host-unreachable"

    // Port unreachable
    Ipv6AclIcmpTypeCodeEnum_port_unreachable Ipv6AclIcmpTypeCodeEnum = "port-unreachable"

    // All unreachables
    Ipv6AclIcmpTypeCodeEnum_unreachable Ipv6AclIcmpTypeCodeEnum = "unreachable"

    // packet too big
    Ipv6AclIcmpTypeCodeEnum_packet_too_big Ipv6AclIcmpTypeCodeEnum = "packet-too-big"

    // TTL exceeded
    Ipv6AclIcmpTypeCodeEnum_ttl_exceeded Ipv6AclIcmpTypeCodeEnum = "ttl-exceeded"

    // Reassembly timeout
    Ipv6AclIcmpTypeCodeEnum_reassembly_timeout Ipv6AclIcmpTypeCodeEnum = "reassembly-timeout"

    // All time exceeds
    Ipv6AclIcmpTypeCodeEnum_time_exceeded Ipv6AclIcmpTypeCodeEnum = "time-exceeded"

    // Erroneous header field
    Ipv6AclIcmpTypeCodeEnum_erronenous_header_field Ipv6AclIcmpTypeCodeEnum = "erronenous-header-field"

    // Parameter required but not present
    Ipv6AclIcmpTypeCodeEnum_option_missing Ipv6AclIcmpTypeCodeEnum = "option-missing"

    // Parameter required but no room
    Ipv6AclIcmpTypeCodeEnum_no_room_for_option Ipv6AclIcmpTypeCodeEnum = "no-room-for-option"

    // All parameter problems
    Ipv6AclIcmpTypeCodeEnum_parameter_problem Ipv6AclIcmpTypeCodeEnum = "parameter-problem"

    // Echo ping
    Ipv6AclIcmpTypeCodeEnum_echo Ipv6AclIcmpTypeCodeEnum = "echo"

    // Echo reply
    Ipv6AclIcmpTypeCodeEnum_echo_reply Ipv6AclIcmpTypeCodeEnum = "echo-reply"

    // Multicast listener query
    Ipv6AclIcmpTypeCodeEnum_group_membership_query Ipv6AclIcmpTypeCodeEnum = "group-membership-query"

    // Multicast listener report
    Ipv6AclIcmpTypeCodeEnum_group_membership_report Ipv6AclIcmpTypeCodeEnum = "group-membership-report"

    // Multicast listener done
    Ipv6AclIcmpTypeCodeEnum_group_membership_reduction Ipv6AclIcmpTypeCodeEnum = "group-membership-reduction"

    // Router discovery solicitations
    Ipv6AclIcmpTypeCodeEnum_router_solicitation Ipv6AclIcmpTypeCodeEnum = "router-solicitation"

    // Router discovery advertisements
    Ipv6AclIcmpTypeCodeEnum_router_advertisement Ipv6AclIcmpTypeCodeEnum = "router-advertisement"

    // Neighbor discovery neighbor solicitations
    Ipv6AclIcmpTypeCodeEnum_neighbor_solicitation Ipv6AclIcmpTypeCodeEnum = "neighbor-solicitation"

    // Neighbor discovery neighbor advertisements
    Ipv6AclIcmpTypeCodeEnum_neighbor_advertisement Ipv6AclIcmpTypeCodeEnum = "neighbor-advertisement"

    // All redirects
    Ipv6AclIcmpTypeCodeEnum_redirect Ipv6AclIcmpTypeCodeEnum = "redirect"

    // Router renumbering command
    Ipv6AclIcmpTypeCodeEnum_rr_command Ipv6AclIcmpTypeCodeEnum = "rr-command"

    // Router renumbering result
    Ipv6AclIcmpTypeCodeEnum_rr_result Ipv6AclIcmpTypeCodeEnum = "rr-result"

    // Router renumbering seqnum
    Ipv6AclIcmpTypeCodeEnum_rr_seqnum_reset Ipv6AclIcmpTypeCodeEnum = "rr-seqnum-reset"

    // Router renumbering
    Ipv6AclIcmpTypeCodeEnum_router_renumbering Ipv6AclIcmpTypeCodeEnum = "router-renumbering"

    // Query subject is ipv6 address
    Ipv6AclIcmpTypeCodeEnum_query_subject_is_ipv6_address Ipv6AclIcmpTypeCodeEnum = "query-subject-is-ipv6-address"

    // Query subject is domain name
    Ipv6AclIcmpTypeCodeEnum_query_subject_is_domain_name Ipv6AclIcmpTypeCodeEnum = "query-subject-is-domain-name"

    // Query subject is ipv4 address
    Ipv6AclIcmpTypeCodeEnum_query_subject_is_ipv4_address Ipv6AclIcmpTypeCodeEnum = "query-subject-is-ipv4-address"

    // Who are you request
    Ipv6AclIcmpTypeCodeEnum_who_are_you_request Ipv6AclIcmpTypeCodeEnum = "who-are-you-request"

    // Node information successful reply
    Ipv6AclIcmpTypeCodeEnum_node_information_successful_reply Ipv6AclIcmpTypeCodeEnum = "node-information-successful-reply"

    // Node information reply rejected
    Ipv6AclIcmpTypeCodeEnum_node_information_request_is_refused Ipv6AclIcmpTypeCodeEnum = "node-information-request-is-refused"

    // Unknown query type
    Ipv6AclIcmpTypeCodeEnum_unknown_query_type Ipv6AclIcmpTypeCodeEnum = "unknown-query-type"

    // Who are you reply
    Ipv6AclIcmpTypeCodeEnum_who_are_you_reply Ipv6AclIcmpTypeCodeEnum = "who-are-you-reply"
)

// Ipv6AclPrecedenceNumber represents Ipv6 acl precedence number
type Ipv6AclPrecedenceNumber string

const (
    // Match packets with critical precedence
    Ipv6AclPrecedenceNumber_critical Ipv6AclPrecedenceNumber = "critical"

    // Match packets with flash precedence
    Ipv6AclPrecedenceNumber_flash Ipv6AclPrecedenceNumber = "flash"

    // Match packets with flash override precedence
    Ipv6AclPrecedenceNumber_flash_override Ipv6AclPrecedenceNumber = "flash-override"

    // Match packets with immediate precedence
    Ipv6AclPrecedenceNumber_immediate Ipv6AclPrecedenceNumber = "immediate"

    // Match packets with internetwork control
    // precedence
    Ipv6AclPrecedenceNumber_internet Ipv6AclPrecedenceNumber = "internet"

    // Match packets with network control precedence
    Ipv6AclPrecedenceNumber_network Ipv6AclPrecedenceNumber = "network"

    // Match packets with priority precedence
    Ipv6AclPrecedenceNumber_priority Ipv6AclPrecedenceNumber = "priority"

    // Match packets with routine precedence
    Ipv6AclPrecedenceNumber_routine Ipv6AclPrecedenceNumber = "routine"
)

// Ipv6AclTypeEnum represents Ipv6 acl type enum
type Ipv6AclTypeEnum string

const (
    // ACL
    Ipv6AclTypeEnum_acl Ipv6AclTypeEnum = "acl"

    // Prefix List
    Ipv6AclTypeEnum_prefix_list Ipv6AclTypeEnum = "prefix-list"
)

// Ipv6PrefixMatchMinLength represents Ipv6 prefix match min length
type Ipv6PrefixMatchMinLength string

const (
    // Enable matching of Prefix Lengths greater than
    // MinPrefixLength
    Ipv6PrefixMatchMinLength_match_min_length Ipv6PrefixMatchMinLength = "match-min-length"
)

// Ipv6AclDscpNumber represents Ipv6 acl dscp number
type Ipv6AclDscpNumber string

const (
    // Default DSCP
    Ipv6AclDscpNumber_default_ Ipv6AclDscpNumber = "default"

    // Match packets with AF11 DSCP
    Ipv6AclDscpNumber_af11 Ipv6AclDscpNumber = "af11"

    // Match packets with AF12 DSCP
    Ipv6AclDscpNumber_af12 Ipv6AclDscpNumber = "af12"

    // Match packets with AF13 DSCP
    Ipv6AclDscpNumber_af13 Ipv6AclDscpNumber = "af13"

    // Match packets with AF21 DSCP
    Ipv6AclDscpNumber_af21 Ipv6AclDscpNumber = "af21"

    // Match packets with AF22 DSCP
    Ipv6AclDscpNumber_af22 Ipv6AclDscpNumber = "af22"

    // Match packets with AF23 DSCP
    Ipv6AclDscpNumber_af23 Ipv6AclDscpNumber = "af23"

    // Match packets with AF31 DSCP
    Ipv6AclDscpNumber_af31 Ipv6AclDscpNumber = "af31"

    // Match packets with AF32 DSCP
    Ipv6AclDscpNumber_af32 Ipv6AclDscpNumber = "af32"

    // Match packets with AF33 DSCP
    Ipv6AclDscpNumber_af33 Ipv6AclDscpNumber = "af33"

    // Match packets with AF41 DSCP
    Ipv6AclDscpNumber_af41 Ipv6AclDscpNumber = "af41"

    // Match packets with AF42 DSCP
    Ipv6AclDscpNumber_af42 Ipv6AclDscpNumber = "af42"

    // Match packets with AF43 DSCP
    Ipv6AclDscpNumber_af43 Ipv6AclDscpNumber = "af43"

    // Match packets with CS1 (precedence 1) DSCP
    Ipv6AclDscpNumber_cs1 Ipv6AclDscpNumber = "cs1"

    // Match packets with CS2 (precedence 2) DSCP
    Ipv6AclDscpNumber_cs2 Ipv6AclDscpNumber = "cs2"

    // Match packets with CS3 (precedence 3) DSCP
    Ipv6AclDscpNumber_cs3 Ipv6AclDscpNumber = "cs3"

    // Match packets with CS4 (precedence 4) DSCP
    Ipv6AclDscpNumber_cs4 Ipv6AclDscpNumber = "cs4"

    // Match packets with CS5 (precedence 5) DSCP
    Ipv6AclDscpNumber_cs5 Ipv6AclDscpNumber = "cs5"

    // Match packets with CS6 (precedence 6) DSCP
    Ipv6AclDscpNumber_cs6 Ipv6AclDscpNumber = "cs6"

    // Match packets with CS7 (precedence 7) DSCP
    Ipv6AclDscpNumber_cs7 Ipv6AclDscpNumber = "cs7"

    // Match packets with EF DSCP
    Ipv6AclDscpNumber_ef Ipv6AclDscpNumber = "ef"
)

// Ipv6AclTcpBitsNumber represents Ipv6 acl tcp bits number
type Ipv6AclTcpBitsNumber string

const (
    // Match established connections (0x14)
    Ipv6AclTcpBitsNumber_established Ipv6AclTcpBitsNumber = "established"

    // Match on the ACK bit (0x10)
    Ipv6AclTcpBitsNumber_ack Ipv6AclTcpBitsNumber = "ack"

    // Match on the RST bit (0x04)
    Ipv6AclTcpBitsNumber_rst Ipv6AclTcpBitsNumber = "rst"

    // Match on the FIN bit (0x01)
    Ipv6AclTcpBitsNumber_fin Ipv6AclTcpBitsNumber = "fin"

    // Match on the PSH bit (0x08)
    Ipv6AclTcpBitsNumber_psh Ipv6AclTcpBitsNumber = "psh"

    // Match on the SYN bit (0x02)
    Ipv6AclTcpBitsNumber_syn Ipv6AclTcpBitsNumber = "syn"

    // Match on the URG bit (0x20)
    Ipv6AclTcpBitsNumber_urg Ipv6AclTcpBitsNumber = "urg"
)

// Ipv6PrefixMatchMaxLength represents Ipv6 prefix match max length
type Ipv6PrefixMatchMaxLength string

const (
    // Enable matching of Prefix Lengths lesser than
    // MaxPrefixLength
    Ipv6PrefixMatchMaxLength_match_max_length Ipv6PrefixMatchMaxLength = "match-max-length"
)

// Ipv6AclPortNumber represents Ipv6 acl port number
type Ipv6AclPortNumber string

const (
    // Match on the 'echo' port number
    Ipv6AclPortNumber_echo Ipv6AclPortNumber = "echo"

    // Match on the 'discard' port number
    Ipv6AclPortNumber_discard Ipv6AclPortNumber = "discard"

    // Match on the 'daytime' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_daytime Ipv6AclPortNumber = "daytime"

    // Match on the 'chargen' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_char_gen Ipv6AclPortNumber = "char-gen"

    // Match on the FTP data connections port number
    // (TCP/SCTP only)
    Ipv6AclPortNumber_ftp_data Ipv6AclPortNumber = "ftp-data"

    // Match on the 'ftp' port number (TCP/SCTP only)
    Ipv6AclPortNumber_ftp Ipv6AclPortNumber = "ftp"

    // Match on the 'telnet' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_telnet Ipv6AclPortNumber = "telnet"

    // Match on the 'smtp' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_smtp Ipv6AclPortNumber = "smtp"

    // Match on the 'time' port number
    Ipv6AclPortNumber_time Ipv6AclPortNumber = "time"

    // Match on the IEN116 name service port number
    // (UDP only)
    Ipv6AclPortNumber_name_server Ipv6AclPortNumber = "name-server"

    // Match on the 'nicname' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_who_is Ipv6AclPortNumber = "who-is"

    // Match on the 'tacacs' port number
    Ipv6AclPortNumber_tacacs Ipv6AclPortNumber = "tacacs"

    // Match on the 'dns' port number
    Ipv6AclPortNumber_dns Ipv6AclPortNumber = "dns"

    // Match on the Bootstrap Protocol server port
    // number (UDP only)
    Ipv6AclPortNumber_boot_ps Ipv6AclPortNumber = "boot-ps"

    // Match on the Bootstrap Protocol client port
    // number (UDP only)
    Ipv6AclPortNumber_boot_pc Ipv6AclPortNumber = "boot-pc"

    // Match on the 'tftp' port number (UDP only)
    Ipv6AclPortNumber_tftp Ipv6AclPortNumber = "tftp"

    // Match on the 'gopher' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_gopher Ipv6AclPortNumber = "gopher"

    // Match on the 'finger' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_finger Ipv6AclPortNumber = "finger"

    // Match on the 'http' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_www Ipv6AclPortNumber = "www"

    // Match on the NIC hostname server port number
    // (TCP/SCTP only)
    Ipv6AclPortNumber_host_name Ipv6AclPortNumber = "host-name"

    // Match on the 'pop2' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_pop2 Ipv6AclPortNumber = "pop2"

    // Match on the 'pop3' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_pop3 Ipv6AclPortNumber = "pop3"

    // Match on the Sun RPC port number
    Ipv6AclPortNumber_sun_rpc Ipv6AclPortNumber = "sun-rpc"

    // Match on the 'ident' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_ident Ipv6AclPortNumber = "ident"

    // Match on the 'nntp' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_nntp Ipv6AclPortNumber = "nntp"

    // Match on the 'ntp' port number (UDP only)
    Ipv6AclPortNumber_ntp Ipv6AclPortNumber = "ntp"

    // Match on the NetBIOS name service port number
    // (UDP only)
    Ipv6AclPortNumber_net_bios_ns Ipv6AclPortNumber = "net-bios-ns"

    // Match on the NetBIOS datagram service port
    // number (UDP only)
    Ipv6AclPortNumber_net_bios_dgs Ipv6AclPortNumber = "net-bios-dgs"

    // Match on the NetBIOS session service port
    // number (UDP only)
    Ipv6AclPortNumber_net_bios_ss Ipv6AclPortNumber = "net-bios-ss"

    // Match on the 'snmp' port number (UDP only)
    Ipv6AclPortNumber_snmp Ipv6AclPortNumber = "snmp"

    // Match on the SNMP traps port number (UDP only)
    Ipv6AclPortNumber_snmp_trap Ipv6AclPortNumber = "snmp-trap"

    // Match on the 'xdmcp' port number (UDP only)
    Ipv6AclPortNumber_xdmcp Ipv6AclPortNumber = "xdmcp"

    // Match on the 'bgp' port number (TCP/SCTP only)
    Ipv6AclPortNumber_bgp Ipv6AclPortNumber = "bgp"

    // Match on the 'irc' port number (TCP/SCTP only)
    Ipv6AclPortNumber_irc Ipv6AclPortNumber = "irc"

    // Match on the DNSIX security protocol auditing
    // port number (UDP only)
    Ipv6AclPortNumber_dnsix Ipv6AclPortNumber = "dnsix"

    // Match on the mobile IP registration port
    // number (UDP only)
    Ipv6AclPortNumber_mobile_ip Ipv6AclPortNumber = "mobile-ip"

    // Match on the PIM Auto-RP port number
    Ipv6AclPortNumber_pim_auto_rp Ipv6AclPortNumber = "pim-auto-rp"

    // Match on the 'isakmp' port number (UDP only)
    Ipv6AclPortNumber_isakmp Ipv6AclPortNumber = "isakmp"

    // Match on the port used by TCP/SCTP for 'exec'
    // and by UDP for 'biff'
    Ipv6AclPortNumber_exec_or_biff Ipv6AclPortNumber = "exec-or-biff"

    // Match on the port used by TCP/SCTP for 'login'
    // and by UDP for 'rwho'
    Ipv6AclPortNumber_login_or_who Ipv6AclPortNumber = "login-or-who"

    // Match on the port used by TCP/SCTP for 'rcmd'
    // and by UDP for 'syslog'
    Ipv6AclPortNumber_cmd_or_syslog Ipv6AclPortNumber = "cmd-or-syslog"

    // Match on the 'lpd' port number (TCP/SCTP only)
    Ipv6AclPortNumber_lpd Ipv6AclPortNumber = "lpd"

    // Match on the 'talk' port number
    Ipv6AclPortNumber_talk Ipv6AclPortNumber = "talk"

    // Match on the 'rip' port number (UDP only)
    Ipv6AclPortNumber_rip Ipv6AclPortNumber = "rip"

    // Match on the 'uucp' port number (TCP/SCTP
    // only)
    Ipv6AclPortNumber_uucp Ipv6AclPortNumber = "uucp"

    // Match on the Kerberos login port number
    // (TCP/SCTP only)
    Ipv6AclPortNumber_klogin Ipv6AclPortNumber = "klogin"

    // Match on the Kerberos shell port number
    // (TCP/SCTP only)
    Ipv6AclPortNumber_kshell Ipv6AclPortNumber = "kshell"

    // Match on the LDP port
    Ipv6AclPortNumber_ldp Ipv6AclPortNumber = "ldp"
)

// Ipv6AclStatusEnum represents Ipv6 acl status enum
type Ipv6AclStatusEnum string

const (
    // Disabled
    Ipv6AclStatusEnum_disabled Ipv6AclStatusEnum = "disabled"

    // Enabled
    Ipv6AclStatusEnum_enabled Ipv6AclStatusEnum = "enabled"
)

// Ipv6AclOperatorEnum represents Ipv6 acl operator enum
type Ipv6AclOperatorEnum string

const (
    // Match only packets on a given port number
    Ipv6AclOperatorEnum_equal Ipv6AclOperatorEnum = "equal"

    // Match only packet with a greater port number
    Ipv6AclOperatorEnum_greater_than Ipv6AclOperatorEnum = "greater-than"

    // Match only packet with a lower port number
    Ipv6AclOperatorEnum_less_than Ipv6AclOperatorEnum = "less-than"

    // Match only packets not on a given port number
    Ipv6AclOperatorEnum_not_equal Ipv6AclOperatorEnum = "not-equal"

    // Match only packets in the range of port numbers
    Ipv6AclOperatorEnum_range_ Ipv6AclOperatorEnum = "range"
)

// Ipv6AclTcpMatchOperatorEnum represents Ipv6 acl tcp match operator enum
type Ipv6AclTcpMatchOperatorEnum string

const (
    // Match only packet with all the given TCP bits
    Ipv6AclTcpMatchOperatorEnum_match_all Ipv6AclTcpMatchOperatorEnum = "match-all"

    // Match only packet with any of the given TCP
    // bits
    Ipv6AclTcpMatchOperatorEnum_match_any Ipv6AclTcpMatchOperatorEnum = "match-any"
)

