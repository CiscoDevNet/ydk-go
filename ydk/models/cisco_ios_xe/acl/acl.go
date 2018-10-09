// Cisco XE Native Access Control List (ACL) Yang model.
// Copyright (c) 2016, 2018 by Cisco Systems, Inc.
// All rights reserved.
package acl

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package acl"))
}

// AclPortType
type AclPortType string

const (
    AclPortType_bgp AclPortType = "bgp"

    AclPortType_chargen AclPortType = "chargen"

    AclPortType_daytime AclPortType = "daytime"

    AclPortType_discard AclPortType = "discard"

    AclPortType_domain AclPortType = "domain"

    AclPortType_echo AclPortType = "echo"

    AclPortType_finger AclPortType = "finger"

    AclPortType_ftp AclPortType = "ftp"

    AclPortType_ftp_data AclPortType = "ftp-data"

    AclPortType_gopher AclPortType = "gopher"

    AclPortType_hostname AclPortType = "hostname"

    AclPortType_ident AclPortType = "ident"

    AclPortType_irc AclPortType = "irc"

    AclPortType_klogin AclPortType = "klogin"

    AclPortType_kshell AclPortType = "kshell"

    AclPortType_lpd AclPortType = "lpd"

    AclPortType_msrpc AclPortType = "msrpc"

    AclPortType_nntp AclPortType = "nntp"

    AclPortType_pim_auto_rp AclPortType = "pim-auto-rp"

    AclPortType_pop2 AclPortType = "pop2"

    AclPortType_pop3 AclPortType = "pop3"

    AclPortType_smtp AclPortType = "smtp"

    AclPortType_sunrpc AclPortType = "sunrpc"

    AclPortType_tacacs AclPortType = "tacacs"

    AclPortType_talk AclPortType = "talk"

    AclPortType_telnet AclPortType = "telnet"

    AclPortType_time AclPortType = "time"

    AclPortType_uucp AclPortType = "uucp"

    AclPortType_whois AclPortType = "whois"

    AclPortType_www AclPortType = "www"

    AclPortType_biff AclPortType = "biff"

    AclPortType_bootpc AclPortType = "bootpc"

    AclPortType_bootps AclPortType = "bootps"

    AclPortType_dnsix AclPortType = "dnsix"

    AclPortType_isakmp AclPortType = "isakmp"

    AclPortType_mobile_ip AclPortType = "mobile-ip"

    AclPortType_nameserver AclPortType = "nameserver"

    AclPortType_netbios_dgm AclPortType = "netbios-dgm"

    AclPortType_netbios_ns AclPortType = "netbios-ns"

    AclPortType_netbios_ss AclPortType = "netbios-ss"

    AclPortType_non500_isakmp AclPortType = "non500-isakmp"

    AclPortType_ntp AclPortType = "ntp"

    AclPortType_rip AclPortType = "rip"

    AclPortType_ripv6 AclPortType = "ripv6"

    AclPortType_snmp AclPortType = "snmp"

    AclPortType_snmptrap AclPortType = "snmptrap"

    AclPortType_syslog AclPortType = "syslog"

    AclPortType_tftp AclPortType = "tftp"

    AclPortType_who AclPortType = "who"

    AclPortType_xdmcp AclPortType = "xdmcp"
)

