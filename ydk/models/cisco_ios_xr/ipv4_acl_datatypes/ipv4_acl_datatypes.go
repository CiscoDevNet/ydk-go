// This module contains a collection of generally useful
// derived YANG data types.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_acl_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_acl_datatypes"))
}

// Ipv4AclProtocolNumber represents Ipv4 acl protocol number
type Ipv4AclProtocolNumber string

const (
    // Any IP protocol
    Ipv4AclProtocolNumber_ip Ipv4AclProtocolNumber = "ip"

    // Internet Control Message Protocol
    Ipv4AclProtocolNumber_icmp Ipv4AclProtocolNumber = "icmp"

    // Internet Gateway Message Protocol
    Ipv4AclProtocolNumber_igmp Ipv4AclProtocolNumber = "igmp"

    // IP in IP tunneling
    Ipv4AclProtocolNumber_ip_in_ip Ipv4AclProtocolNumber = "ip-in-ip"

    // Transport Control Protocol
    Ipv4AclProtocolNumber_tcp Ipv4AclProtocolNumber = "tcp"

    // Cisco's IGRP Routing Protocol
    Ipv4AclProtocolNumber_igrp Ipv4AclProtocolNumber = "igrp"

    // User Datagram Protocol
    Ipv4AclProtocolNumber_udp Ipv4AclProtocolNumber = "udp"

    // Cisco's GRE tunneling
    Ipv4AclProtocolNumber_gre Ipv4AclProtocolNumber = "gre"

    // Encapsulation Security Protocol
    Ipv4AclProtocolNumber_esp Ipv4AclProtocolNumber = "esp"

    // Authentication Header Protocol
    Ipv4AclProtocolNumber_ahp Ipv4AclProtocolNumber = "ahp"

    // Cisco's EIGRP Routing Protocol
    Ipv4AclProtocolNumber_eigrp Ipv4AclProtocolNumber = "eigrp"

    // OSPF Routing Protocol
    Ipv4AclProtocolNumber_ospf Ipv4AclProtocolNumber = "ospf"

    // KA9Q NOS Compatible IP over IP tunneling
    Ipv4AclProtocolNumber_nos Ipv4AclProtocolNumber = "nos"

    // Protocol Independent Multicast
    Ipv4AclProtocolNumber_pim Ipv4AclProtocolNumber = "pim"

    // Payload Compression Protocol
    Ipv4AclProtocolNumber_pcp Ipv4AclProtocolNumber = "pcp"

    // Stream Control Transmission Protocol
    Ipv4AclProtocolNumber_sctp Ipv4AclProtocolNumber = "sctp"
)

// Ipv4AclTcpMatchOperatorEnum represents Ipv4 acl tcp match operator enum
type Ipv4AclTcpMatchOperatorEnum string

const (
    // Match only packet with all the given TCP bits
    Ipv4AclTcpMatchOperatorEnum_match_all Ipv4AclTcpMatchOperatorEnum = "match-all"

    // Match only packet with any of the given TCP
    // bits
    Ipv4AclTcpMatchOperatorEnum_match_any Ipv4AclTcpMatchOperatorEnum = "match-any"
)

// Ipv4AclGrantEnum represents Ipv4 acl grant enum
type Ipv4AclGrantEnum string

const (
    // Deny
    Ipv4AclGrantEnum_deny Ipv4AclGrantEnum = "deny"

    // Permit
    Ipv4AclGrantEnum_permit Ipv4AclGrantEnum = "permit"
)

// Ipv4AclOperatorEnum represents Ipv4 acl operator enum
type Ipv4AclOperatorEnum string

const (
    // Match only packets on a given port number
    Ipv4AclOperatorEnum_equal Ipv4AclOperatorEnum = "equal"

    // Match only packet with a greater port number
    Ipv4AclOperatorEnum_greater_than Ipv4AclOperatorEnum = "greater-than"

    // Match only packet with a lower port number
    Ipv4AclOperatorEnum_less_than Ipv4AclOperatorEnum = "less-than"

    // Match only packets not on a given port number
    Ipv4AclOperatorEnum_not_equal Ipv4AclOperatorEnum = "not-equal"

    // Match only packets in the range of port numbers
    Ipv4AclOperatorEnum_range_ Ipv4AclOperatorEnum = "range"
)

// Ipv4AclPrecedenceNumber represents Ipv4 acl precedence number
type Ipv4AclPrecedenceNumber string

const (
    // Match packets with critical precedence
    Ipv4AclPrecedenceNumber_critical Ipv4AclPrecedenceNumber = "critical"

    // Match packets with flash precedence
    Ipv4AclPrecedenceNumber_flash Ipv4AclPrecedenceNumber = "flash"

    // Match packets with flash override precedence
    Ipv4AclPrecedenceNumber_flash_override Ipv4AclPrecedenceNumber = "flash-override"

    // Match packets with immediate precedence
    Ipv4AclPrecedenceNumber_immediate Ipv4AclPrecedenceNumber = "immediate"

    // Match packets with internetwork control
    // precedence
    Ipv4AclPrecedenceNumber_internet Ipv4AclPrecedenceNumber = "internet"

    // Match packets with network control precedence
    Ipv4AclPrecedenceNumber_network Ipv4AclPrecedenceNumber = "network"

    // Match packets with priority precedence
    Ipv4AclPrecedenceNumber_priority Ipv4AclPrecedenceNumber = "priority"

    // Match packets with routine precedence
    Ipv4AclPrecedenceNumber_routine Ipv4AclPrecedenceNumber = "routine"
)

// Ipv4AclPortNumber represents Ipv4 acl port number
type Ipv4AclPortNumber string

const (
    // Match on the 'echo' port number
    Ipv4AclPortNumber_echo Ipv4AclPortNumber = "echo"

    // Match on the 'discard' port number
    Ipv4AclPortNumber_discard Ipv4AclPortNumber = "discard"

    // Match on the 'daytime' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_daytime Ipv4AclPortNumber = "daytime"

    // Match on the 'chargen' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_char_gen Ipv4AclPortNumber = "char-gen"

    // Match on the FTP data connections port number
    // (TCP/SCTP only)
    Ipv4AclPortNumber_ftp_data Ipv4AclPortNumber = "ftp-data"

    // Match on the 'ftp' port number (TCP/SCTP only)
    Ipv4AclPortNumber_ftp Ipv4AclPortNumber = "ftp"

    // Match on the 'telnet' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_telnet Ipv4AclPortNumber = "telnet"

    // Match on the 'smtp' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_smtp Ipv4AclPortNumber = "smtp"

    // Match on the 'time' port number
    Ipv4AclPortNumber_time Ipv4AclPortNumber = "time"

    // Match on the IEN116 name service port number
    // (UDP only)
    Ipv4AclPortNumber_name_server Ipv4AclPortNumber = "name-server"

    // Match on the 'nicname' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_who_is Ipv4AclPortNumber = "who-is"

    // Match on the 'tacacs' port number
    Ipv4AclPortNumber_tacacs Ipv4AclPortNumber = "tacacs"

    // Match on the 'dns' port number
    Ipv4AclPortNumber_dns Ipv4AclPortNumber = "dns"

    // Match on the Bootstrap Protocol server port
    // number (UDP only)
    Ipv4AclPortNumber_boot_ps Ipv4AclPortNumber = "boot-ps"

    // Match on the Bootstrap Protocol client port
    // number (UDP only)
    Ipv4AclPortNumber_boot_pc Ipv4AclPortNumber = "boot-pc"

    // Match on the 'tftp' port number (UDP only)
    Ipv4AclPortNumber_tftp Ipv4AclPortNumber = "tftp"

    // Match on the 'gopher' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_gopher Ipv4AclPortNumber = "gopher"

    // Match on the 'finger' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_finger Ipv4AclPortNumber = "finger"

    // Match on the 'http' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_www Ipv4AclPortNumber = "www"

    // Match on the NIC hostname server port number
    // (TCP/SCTP only)
    Ipv4AclPortNumber_host_name Ipv4AclPortNumber = "host-name"

    // Match on the 'pop2' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_pop2 Ipv4AclPortNumber = "pop2"

    // Match on the 'pop3' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_pop3 Ipv4AclPortNumber = "pop3"

    // Match on the Sun RPC port number
    Ipv4AclPortNumber_sun_rpc Ipv4AclPortNumber = "sun-rpc"

    // Match on the 'ident' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_ident Ipv4AclPortNumber = "ident"

    // Match on the 'nntp' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_nntp Ipv4AclPortNumber = "nntp"

    // Match on the 'ntp' port number (UDP only)
    Ipv4AclPortNumber_ntp Ipv4AclPortNumber = "ntp"

    // Match on the NetBIOS name service port number
    // (UDP only)
    Ipv4AclPortNumber_net_bios_ns Ipv4AclPortNumber = "net-bios-ns"

    // Match on the NetBIOS datagram service port
    // number (UDP only)
    Ipv4AclPortNumber_net_bios_dgs Ipv4AclPortNumber = "net-bios-dgs"

    // Match on the NetBIOS session service port
    // number (UDP only)
    Ipv4AclPortNumber_net_bios_ss Ipv4AclPortNumber = "net-bios-ss"

    // Match on the 'snmp' port number (UDP only)
    Ipv4AclPortNumber_snmp Ipv4AclPortNumber = "snmp"

    // Match on the SNMP traps port number (UDP only)
    Ipv4AclPortNumber_snmp_trap Ipv4AclPortNumber = "snmp-trap"

    // Match on the 'xdmcp' port number (UDP only)
    Ipv4AclPortNumber_xdmcp Ipv4AclPortNumber = "xdmcp"

    // Match on the 'bgp' port number (TCP/SCTP only)
    Ipv4AclPortNumber_bgp Ipv4AclPortNumber = "bgp"

    // Match on the 'irc' port number (TCP/SCTP only)
    Ipv4AclPortNumber_irc Ipv4AclPortNumber = "irc"

    // Match on the DNSIX security protocol auditing
    // port number (UDP only)
    Ipv4AclPortNumber_dnsix Ipv4AclPortNumber = "dnsix"

    // Match on the mobile IP registration port
    // number (UDP only)
    Ipv4AclPortNumber_mobile_ip Ipv4AclPortNumber = "mobile-ip"

    // Match on the PIM Auto-RP port number
    Ipv4AclPortNumber_pim_auto_rp Ipv4AclPortNumber = "pim-auto-rp"

    // Match on the 'isakmp' port number (UDP only)
    Ipv4AclPortNumber_isakmp Ipv4AclPortNumber = "isakmp"

    // Match on the port used by TCP/SCTP for 'exec'
    // and by UDP for 'biff'
    Ipv4AclPortNumber_exec_or_biff Ipv4AclPortNumber = "exec-or-biff"

    // Match on the port used by TCP/SCTP for 'login'
    // and by UDP for 'rwho'
    Ipv4AclPortNumber_login_or_who Ipv4AclPortNumber = "login-or-who"

    // Match on the port used by TCP/SCTP for 'rcmd'
    // and by UDP for 'syslog'
    Ipv4AclPortNumber_cmd_or_syslog Ipv4AclPortNumber = "cmd-or-syslog"

    // Match on the 'lpd' port number (TCP/SCTP only)
    Ipv4AclPortNumber_lpd Ipv4AclPortNumber = "lpd"

    // Match on the 'talk' port number
    Ipv4AclPortNumber_talk Ipv4AclPortNumber = "talk"

    // Match on the 'rip' port number (UDP only)
    Ipv4AclPortNumber_rip Ipv4AclPortNumber = "rip"

    // Match on the 'uucp' port number (TCP/SCTP
    // only)
    Ipv4AclPortNumber_uucp Ipv4AclPortNumber = "uucp"

    // Match on the Kerberos login port number
    // (TCP/SCTP only)
    Ipv4AclPortNumber_klogin Ipv4AclPortNumber = "klogin"

    // Match on the Kerberos shell port number
    // (TCP/SCTP only)
    Ipv4AclPortNumber_kshell Ipv4AclPortNumber = "kshell"

    // Match on the LDP port
    Ipv4AclPortNumber_ldp Ipv4AclPortNumber = "ldp"
)

// Ipv4AclIgmpNumber represents Ipv4 acl igmp number
type Ipv4AclIgmpNumber string

const (
    // Host query
    Ipv4AclIgmpNumber_host_query Ipv4AclIgmpNumber = "host-query"

    // Host report
    Ipv4AclIgmpNumber_host_report Ipv4AclIgmpNumber = "host-report"

    // Distance Vector Multicast Routing Protocol
    Ipv4AclIgmpNumber_dvmrp Ipv4AclIgmpNumber = "dvmrp"

    // Portocol Independent Multicast
    Ipv4AclIgmpNumber_pim Ipv4AclIgmpNumber = "pim"

    // Multicast Trace
    Ipv4AclIgmpNumber_trace Ipv4AclIgmpNumber = "trace"

    // Version 2 report
    Ipv4AclIgmpNumber_v2_report Ipv4AclIgmpNumber = "v2-report"

    // Version 2 leave
    Ipv4AclIgmpNumber_v2_leave Ipv4AclIgmpNumber = "v2-leave"

    // MTrace response
    Ipv4AclIgmpNumber_mtrace_response Ipv4AclIgmpNumber = "mtrace-response"

    // MTrace
    Ipv4AclIgmpNumber_mtrace Ipv4AclIgmpNumber = "mtrace"

    // Version 3 report
    Ipv4AclIgmpNumber_v3_report Ipv4AclIgmpNumber = "v3-report"
)

// Ipv4AclDscpNumber represents Ipv4 acl dscp number
type Ipv4AclDscpNumber string

const (
    // Default DSCP
    Ipv4AclDscpNumber_default_ Ipv4AclDscpNumber = "default"

    // Match packets with AF11 DSCP
    Ipv4AclDscpNumber_af11 Ipv4AclDscpNumber = "af11"

    // Match packets with AF12 DSCP
    Ipv4AclDscpNumber_af12 Ipv4AclDscpNumber = "af12"

    // Match packets with AF13 DSCP
    Ipv4AclDscpNumber_af13 Ipv4AclDscpNumber = "af13"

    // Match packets with AF21 DSCP
    Ipv4AclDscpNumber_af21 Ipv4AclDscpNumber = "af21"

    // Match packets with AF22 DSCP
    Ipv4AclDscpNumber_af22 Ipv4AclDscpNumber = "af22"

    // Match packets with AF23 DSCP
    Ipv4AclDscpNumber_af23 Ipv4AclDscpNumber = "af23"

    // Match packets with AF31 DSCP
    Ipv4AclDscpNumber_af31 Ipv4AclDscpNumber = "af31"

    // Match packets with AF32 DSCP
    Ipv4AclDscpNumber_af32 Ipv4AclDscpNumber = "af32"

    // Match packets with AF33 DSCP
    Ipv4AclDscpNumber_af33 Ipv4AclDscpNumber = "af33"

    // Match packets with AF41 DSCP
    Ipv4AclDscpNumber_af41 Ipv4AclDscpNumber = "af41"

    // Match packets with AF42 DSCP
    Ipv4AclDscpNumber_af42 Ipv4AclDscpNumber = "af42"

    // Match packets with AF43 DSCP
    Ipv4AclDscpNumber_af43 Ipv4AclDscpNumber = "af43"

    // Match packets with CS1 (precedence 1) DSCP
    Ipv4AclDscpNumber_cs1 Ipv4AclDscpNumber = "cs1"

    // Match packets with CS2 (precedence 2) DSCP
    Ipv4AclDscpNumber_cs2 Ipv4AclDscpNumber = "cs2"

    // Match packets with CS3 (precedence 3) DSCP
    Ipv4AclDscpNumber_cs3 Ipv4AclDscpNumber = "cs3"

    // Match packets with CS4 (precedence 4) DSCP
    Ipv4AclDscpNumber_cs4 Ipv4AclDscpNumber = "cs4"

    // Match packets with CS5 (precedence 5) DSCP
    Ipv4AclDscpNumber_cs5 Ipv4AclDscpNumber = "cs5"

    // Match packets with CS6 (precedence 6) DSCP
    Ipv4AclDscpNumber_cs6 Ipv4AclDscpNumber = "cs6"

    // Match packets with CS7 (precedence 7) DSCP
    Ipv4AclDscpNumber_cs7 Ipv4AclDscpNumber = "cs7"

    // Match packets with EF DSCP
    Ipv4AclDscpNumber_ef Ipv4AclDscpNumber = "ef"
)

// Ipv4AclStatusEnum represents Ipv4 acl status enum
type Ipv4AclStatusEnum string

const (
    // Disabled
    Ipv4AclStatusEnum_disabled Ipv4AclStatusEnum = "disabled"

    // Enabled
    Ipv4AclStatusEnum_enabled Ipv4AclStatusEnum = "enabled"
)

// Ipv4AclIcmpTypeCodeEnum represents Ipv4 acl icmp type code enum
type Ipv4AclIcmpTypeCodeEnum string

const (
    // Echo reply
    Ipv4AclIcmpTypeCodeEnum_echo_reply Ipv4AclIcmpTypeCodeEnum = "echo-reply"

    // Network unreachable
    Ipv4AclIcmpTypeCodeEnum_network_unreachable Ipv4AclIcmpTypeCodeEnum = "network-unreachable"

    // Host unreachable
    Ipv4AclIcmpTypeCodeEnum_host_unreachable Ipv4AclIcmpTypeCodeEnum = "host-unreachable"

    // Protocol unreachable
    Ipv4AclIcmpTypeCodeEnum_protocol_unreachable Ipv4AclIcmpTypeCodeEnum = "protocol-unreachable"

    // Port unreachable
    Ipv4AclIcmpTypeCodeEnum_port_unreachable Ipv4AclIcmpTypeCodeEnum = "port-unreachable"

    // Fragmentation needed and DF set
    Ipv4AclIcmpTypeCodeEnum_packet_too_big Ipv4AclIcmpTypeCodeEnum = "packet-too-big"

    // Source route failed
    Ipv4AclIcmpTypeCodeEnum_source_route_failed Ipv4AclIcmpTypeCodeEnum = "source-route-failed"

    // Network unknown
    Ipv4AclIcmpTypeCodeEnum_network_unknown Ipv4AclIcmpTypeCodeEnum = "network-unknown"

    // Host unknown
    Ipv4AclIcmpTypeCodeEnum_host_unknown Ipv4AclIcmpTypeCodeEnum = "host-unknown"

    // Host isolated
    Ipv4AclIcmpTypeCodeEnum_host_isolated Ipv4AclIcmpTypeCodeEnum = "host-isolated"

    // Network prohibited
    Ipv4AclIcmpTypeCodeEnum_dod_net_prohibited Ipv4AclIcmpTypeCodeEnum = "dod-net-prohibited"

    // Host prohibited
    Ipv4AclIcmpTypeCodeEnum_dod_host_prohibited Ipv4AclIcmpTypeCodeEnum = "dod-host-prohibited"

    // Host unreachable for TOS
    Ipv4AclIcmpTypeCodeEnum_host_tos_unreachable Ipv4AclIcmpTypeCodeEnum = "host-tos-unreachable"

    // Network unreachable for TOS
    Ipv4AclIcmpTypeCodeEnum_net_tos_unreachable Ipv4AclIcmpTypeCodeEnum = "net-tos-unreachable"

    // Administratively prohibited
    Ipv4AclIcmpTypeCodeEnum_administratively_prohibited Ipv4AclIcmpTypeCodeEnum = "administratively-prohibited"

    // Host unreachable for precedence
    Ipv4AclIcmpTypeCodeEnum_host_precedence_unreachable Ipv4AclIcmpTypeCodeEnum = "host-precedence-unreachable"

    // Precedence cutoff
    Ipv4AclIcmpTypeCodeEnum_precedence_unreachable Ipv4AclIcmpTypeCodeEnum = "precedence-unreachable"

    // All unreachables
    Ipv4AclIcmpTypeCodeEnum_unreachable Ipv4AclIcmpTypeCodeEnum = "unreachable"

    // Source quenches
    Ipv4AclIcmpTypeCodeEnum_source_quench Ipv4AclIcmpTypeCodeEnum = "source-quench"

    // Network redirect
    Ipv4AclIcmpTypeCodeEnum_network_redirect Ipv4AclIcmpTypeCodeEnum = "network-redirect"

    // Host redirect
    Ipv4AclIcmpTypeCodeEnum_host_redirect Ipv4AclIcmpTypeCodeEnum = "host-redirect"

    // Network redirect for TOS
    Ipv4AclIcmpTypeCodeEnum_net_tos_redirect Ipv4AclIcmpTypeCodeEnum = "net-tos-redirect"

    // Host redirect for TOS
    Ipv4AclIcmpTypeCodeEnum_host_tos_redirect Ipv4AclIcmpTypeCodeEnum = "host-tos-redirect"

    // All redirects
    Ipv4AclIcmpTypeCodeEnum_redirect Ipv4AclIcmpTypeCodeEnum = "redirect"

    // Alternate address
    Ipv4AclIcmpTypeCodeEnum_alternate_address Ipv4AclIcmpTypeCodeEnum = "alternate-address"

    // Echo (ping)
    Ipv4AclIcmpTypeCodeEnum_echo Ipv4AclIcmpTypeCodeEnum = "echo"

    // Router discovery advertisements
    Ipv4AclIcmpTypeCodeEnum_router_advertisement Ipv4AclIcmpTypeCodeEnum = "router-advertisement"

    // Router discovery solicitations
    Ipv4AclIcmpTypeCodeEnum_router_solicitation Ipv4AclIcmpTypeCodeEnum = "router-solicitation"

    // TTL exceeded
    Ipv4AclIcmpTypeCodeEnum_ttl_exceeded Ipv4AclIcmpTypeCodeEnum = "ttl-exceeded"

    // Reassembly timeout
    Ipv4AclIcmpTypeCodeEnum_reassembly_timeout Ipv4AclIcmpTypeCodeEnum = "reassembly-timeout"

    // All time exceeds
    Ipv4AclIcmpTypeCodeEnum_time_exceeded Ipv4AclIcmpTypeCodeEnum = "time-exceeded"

    // Parameter problem
    Ipv4AclIcmpTypeCodeEnum_general_parameter_problem Ipv4AclIcmpTypeCodeEnum = "general-parameter-problem"

    // Parameter required but not present
    Ipv4AclIcmpTypeCodeEnum_option_missing Ipv4AclIcmpTypeCodeEnum = "option-missing"

    // Parameter required but no room
    Ipv4AclIcmpTypeCodeEnum_no_room_for_option Ipv4AclIcmpTypeCodeEnum = "no-room-for-option"

    // All parameter problems
    Ipv4AclIcmpTypeCodeEnum_parameter_problem Ipv4AclIcmpTypeCodeEnum = "parameter-problem"

    // Timestamp requests
    Ipv4AclIcmpTypeCodeEnum_timestamp_request Ipv4AclIcmpTypeCodeEnum = "timestamp-request"

    // Timestamp replies
    Ipv4AclIcmpTypeCodeEnum_timestamp_reply Ipv4AclIcmpTypeCodeEnum = "timestamp-reply"

    // Information request
    Ipv4AclIcmpTypeCodeEnum_information_request Ipv4AclIcmpTypeCodeEnum = "information-request"

    // Information replies
    Ipv4AclIcmpTypeCodeEnum_information_reply Ipv4AclIcmpTypeCodeEnum = "information-reply"

    // Mask requests
    Ipv4AclIcmpTypeCodeEnum_mask_request Ipv4AclIcmpTypeCodeEnum = "mask-request"

    // Mask replies
    Ipv4AclIcmpTypeCodeEnum_mask_reply Ipv4AclIcmpTypeCodeEnum = "mask-reply"

    // Traceroute
    Ipv4AclIcmpTypeCodeEnum_traceroute Ipv4AclIcmpTypeCodeEnum = "traceroute"

    // Datagram conversion
    Ipv4AclIcmpTypeCodeEnum_conversion_error Ipv4AclIcmpTypeCodeEnum = "conversion-error"

    // Mobile host redirect
    Ipv4AclIcmpTypeCodeEnum_mobile_redirect Ipv4AclIcmpTypeCodeEnum = "mobile-redirect"
)

// Ipv4AclFragFlags represents Ipv4 acl frag flags
type Ipv4AclFragFlags string

const (
    // Match don't fragment flag
    Ipv4AclFragFlags_dont_fragment Ipv4AclFragFlags = "dont-fragment"

    // Match is fragment flag
    Ipv4AclFragFlags_is_fragment Ipv4AclFragFlags = "is-fragment"

    // Match first fragment flag
    Ipv4AclFragFlags_first_fragment Ipv4AclFragFlags = "first-fragment"

    // Match last fragment flag
    Ipv4AclFragFlags_last_fragment Ipv4AclFragFlags = "last-fragment"

    // Match don't fragment and is fragment flag
    Ipv4AclFragFlags_dont_fragment_is_fragment Ipv4AclFragFlags = "dont-fragment-is-fragment"

    // Match don't fragment and first fragment flag
    Ipv4AclFragFlags_dont_fragment_first_fragment Ipv4AclFragFlags = "dont-fragment-first-fragment"

    // Match don't fragment and last fragment flag
    Ipv4AclFragFlags_dont_fragment_last_fragment Ipv4AclFragFlags = "dont-fragment-last-fragment"
)

// Ipv4AclLoggingEnum represents Ipv4 acl logging enum
type Ipv4AclLoggingEnum string

const (
    // Log matches against this entry
    Ipv4AclLoggingEnum_log Ipv4AclLoggingEnum = "log"

    // Log matches against this entry, including input
    // interface
    Ipv4AclLoggingEnum_log_input Ipv4AclLoggingEnum = "log-input"
)

