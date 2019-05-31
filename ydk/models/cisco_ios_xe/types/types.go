// Cisco XE Native Common Type Definitions
// Copyright (c) 2016-2018 by Cisco Systems, Inc.
// All rights reserved.
package types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package types"))
}

// AccessListInOutType
type AccessListInOutType string

const (
    AccessListInOutType_in AccessListInOutType = "in"

    AccessListInOutType_out AccessListInOutType = "out"
)

// AclUdpPortType
type AclUdpPortType string

const (
    AclUdpPortType_biff AclUdpPortType = "biff"

    AclUdpPortType_bootpc AclUdpPortType = "bootpc"

    AclUdpPortType_bootps AclUdpPortType = "bootps"

    AclUdpPortType_discard AclUdpPortType = "discard"

    AclUdpPortType_dnsix AclUdpPortType = "dnsix"

    AclUdpPortType_domain AclUdpPortType = "domain"

    AclUdpPortType_echo AclUdpPortType = "echo"

    AclUdpPortType_isakmp AclUdpPortType = "isakmp"

    AclUdpPortType_mobile_ip AclUdpPortType = "mobile-ip"

    AclUdpPortType_nameserver AclUdpPortType = "nameserver"

    AclUdpPortType_netbios_dgm AclUdpPortType = "netbios-dgm"

    AclUdpPortType_netbios_ns AclUdpPortType = "netbios-ns"

    AclUdpPortType_netbios_ss AclUdpPortType = "netbios-ss"

    AclUdpPortType_non500_isakmp AclUdpPortType = "non500-isakmp"

    AclUdpPortType_ntp AclUdpPortType = "ntp"

    AclUdpPortType_pim_auto_rp AclUdpPortType = "pim-auto-rp"

    AclUdpPortType_rip AclUdpPortType = "rip"

    AclUdpPortType_ripv6 AclUdpPortType = "ripv6"

    AclUdpPortType_snmp AclUdpPortType = "snmp"

    AclUdpPortType_snmptrap AclUdpPortType = "snmptrap"

    AclUdpPortType_sunrpc AclUdpPortType = "sunrpc"

    AclUdpPortType_syslog AclUdpPortType = "syslog"

    AclUdpPortType_tacacs AclUdpPortType = "tacacs"

    AclUdpPortType_talk AclUdpPortType = "talk"

    AclUdpPortType_tftp AclUdpPortType = "tftp"

    AclUdpPortType_time AclUdpPortType = "time"

    AclUdpPortType_who AclUdpPortType = "who"

    AclUdpPortType_xdmcp AclUdpPortType = "xdmcp"
)

// AclTcpPortType
type AclTcpPortType string

const (
    AclTcpPortType_bgp AclTcpPortType = "bgp"

    AclTcpPortType_chargen AclTcpPortType = "chargen"

    AclTcpPortType_cmd AclTcpPortType = "cmd"

    AclTcpPortType_daytime AclTcpPortType = "daytime"

    AclTcpPortType_discard AclTcpPortType = "discard"

    AclTcpPortType_domain AclTcpPortType = "domain"

    AclTcpPortType_echo AclTcpPortType = "echo"

    AclTcpPortType_exec AclTcpPortType = "exec"

    AclTcpPortType_finger AclTcpPortType = "finger"

    AclTcpPortType_ftp AclTcpPortType = "ftp"

    AclTcpPortType_ftp_data AclTcpPortType = "ftp-data"

    AclTcpPortType_gopher AclTcpPortType = "gopher"

    AclTcpPortType_hostname AclTcpPortType = "hostname"

    AclTcpPortType_ident AclTcpPortType = "ident"

    AclTcpPortType_irc AclTcpPortType = "irc"

    AclTcpPortType_klogin AclTcpPortType = "klogin"

    AclTcpPortType_kshell AclTcpPortType = "kshell"

    AclTcpPortType_login AclTcpPortType = "login"

    AclTcpPortType_lpd AclTcpPortType = "lpd"

    AclTcpPortType_msrpc AclTcpPortType = "msrpc"

    AclTcpPortType_nntp AclTcpPortType = "nntp"

    AclTcpPortType_pim_auto_rp AclTcpPortType = "pim-auto-rp"

    AclTcpPortType_pop2 AclTcpPortType = "pop2"

    AclTcpPortType_pop3 AclTcpPortType = "pop3"

    AclTcpPortType_smtp AclTcpPortType = "smtp"

    AclTcpPortType_sunrpc AclTcpPortType = "sunrpc"

    AclTcpPortType_tacacs AclTcpPortType = "tacacs"

    AclTcpPortType_talk AclTcpPortType = "talk"

    AclTcpPortType_telnet AclTcpPortType = "telnet"

    AclTcpPortType_time AclTcpPortType = "time"

    AclTcpPortType_uucp AclTcpPortType = "uucp"

    AclTcpPortType_whois AclTcpPortType = "whois"

    AclTcpPortType_www AclTcpPortType = "www"
)

// RedistOspfExternalType
type RedistOspfExternalType string

const (
    RedistOspfExternalType_Y_1 RedistOspfExternalType = "1"

    RedistOspfExternalType_Y_2 RedistOspfExternalType = "2"
)

// CosValueType
type CosValueType string

const (
    CosValueType_cos CosValueType = "cos"

    CosValueType_dscp CosValueType = "dscp"

    CosValueType_exp CosValueType = "exp"

    CosValueType_precedence CosValueType = "precedence"
)

// DscpType
type DscpType string

const (
    DscpType_af11 DscpType = "af11"

    DscpType_af12 DscpType = "af12"

    DscpType_af13 DscpType = "af13"

    DscpType_af21 DscpType = "af21"

    DscpType_af22 DscpType = "af22"

    DscpType_af23 DscpType = "af23"

    DscpType_af31 DscpType = "af31"

    DscpType_af32 DscpType = "af32"

    DscpType_af33 DscpType = "af33"

    DscpType_af41 DscpType = "af41"

    DscpType_af42 DscpType = "af42"

    DscpType_af43 DscpType = "af43"

    DscpType_cs1 DscpType = "cs1"

    DscpType_cs2 DscpType = "cs2"

    DscpType_cs3 DscpType = "cs3"

    DscpType_cs4 DscpType = "cs4"

    DscpType_cs5 DscpType = "cs5"

    DscpType_cs6 DscpType = "cs6"

    DscpType_cs7 DscpType = "cs7"

    DscpType_default_ DscpType = "default"

    DscpType_dscp DscpType = "dscp"

    DscpType_ef DscpType = "ef"

    DscpType_precedence DscpType = "precedence"
)

// ExpValueType
type ExpValueType string

const (
    ExpValueType_cos ExpValueType = "cos"

    ExpValueType_dscp ExpValueType = "dscp"

    ExpValueType_exp ExpValueType = "exp"

    ExpValueType_precedence ExpValueType = "precedence"
)

// InterfaceType
type InterfaceType string

const (
    InterfaceType_BDI InterfaceType = "BDI"

    InterfaceType_FastEthernet InterfaceType = "FastEthernet"

    InterfaceType_GigabitEthernet InterfaceType = "GigabitEthernet"

    InterfaceType_Loopback InterfaceType = "Loopback"

    InterfaceType_Port_channel InterfaceType = "Port-channel"

    InterfaceType_Serial InterfaceType = "Serial"

    InterfaceType_TenGigabitEthernet InterfaceType = "TenGigabitEthernet"

    InterfaceType_Vlan InterfaceType = "Vlan"

    InterfaceType_FiveGigabitEthernet InterfaceType = "FiveGigabitEthernet"

    InterfaceType_TwentyFiveGigE InterfaceType = "TwentyFiveGigE"

    InterfaceType_TwoGigabitEthernet InterfaceType = "TwoGigabitEthernet"
)

// MobilityType
type MobilityType string

const (
    MobilityType_bind_acknowledgement MobilityType = "bind-acknowledgement"

    MobilityType_bind_error MobilityType = "bind-error"

    MobilityType_bind_refresh MobilityType = "bind-refresh"

    MobilityType_bind_update MobilityType = "bind-update"

    MobilityType_cot MobilityType = "cot"

    MobilityType_coti MobilityType = "coti"

    MobilityType_hot MobilityType = "hot"

    MobilityType_hoti MobilityType = "hoti"
)

// PrecValueType
type PrecValueType string

const (
    PrecValueType_cos PrecValueType = "cos"

    PrecValueType_dscp PrecValueType = "dscp"

    PrecValueType_exp PrecValueType = "exp"

    PrecValueType_precedence PrecValueType = "precedence"
)

// PrecedenceType
type PrecedenceType string

const (
    PrecedenceType_critical PrecedenceType = "critical"

    PrecedenceType_flash PrecedenceType = "flash"

    PrecedenceType_flash_override PrecedenceType = "flash-override"

    PrecedenceType_immediate PrecedenceType = "immediate"

    PrecedenceType_internet PrecedenceType = "internet"

    PrecedenceType_network PrecedenceType = "network"

    PrecedenceType_priority PrecedenceType = "priority"

    PrecedenceType_routine PrecedenceType = "routine"
)

// LimitDcNonDcType
type LimitDcNonDcType string

const (
    LimitDcNonDcType_disable LimitDcNonDcType = "disable"
)

// QosValueType
type QosValueType string

const (
    QosValueType_cos QosValueType = "cos"

    QosValueType_dscp QosValueType = "dscp"

    QosValueType_exp QosValueType = "exp"

    QosValueType_precedence QosValueType = "precedence"
)

// WeekdayType
type WeekdayType string

const (
    WeekdayType_Mon WeekdayType = "Mon"

    WeekdayType_Tue WeekdayType = "Tue"

    WeekdayType_Wed WeekdayType = "Wed"

    WeekdayType_Thu WeekdayType = "Thu"

    WeekdayType_Fri WeekdayType = "Fri"

    WeekdayType_Sat WeekdayType = "Sat"

    WeekdayType_Sun WeekdayType = "Sun"
)

// BgpIpv4AfType
type BgpIpv4AfType string

const (
    BgpIpv4AfType_unicast BgpIpv4AfType = "unicast"

    BgpIpv4AfType_multicast BgpIpv4AfType = "multicast"

    BgpIpv4AfType_mdt BgpIpv4AfType = "mdt"

    BgpIpv4AfType_tunnel BgpIpv4AfType = "tunnel"

    BgpIpv4AfType_labeled_unicast BgpIpv4AfType = "labeled-unicast"

    BgpIpv4AfType_flowspec BgpIpv4AfType = "flowspec"

    BgpIpv4AfType_mvpn BgpIpv4AfType = "mvpn"
)

// BgpIpv6AfType
type BgpIpv6AfType string

const (
    BgpIpv6AfType_unicast BgpIpv6AfType = "unicast"

    BgpIpv6AfType_multicast BgpIpv6AfType = "multicast"

    BgpIpv6AfType_mdt BgpIpv6AfType = "mdt"

    BgpIpv6AfType_flowspec BgpIpv6AfType = "flowspec"

    BgpIpv6AfType_mvpn BgpIpv6AfType = "mvpn"
)

// CommunityWellKnownType
type CommunityWellKnownType string

const (
    CommunityWellKnownType_gshut CommunityWellKnownType = "gshut"

    CommunityWellKnownType_internet CommunityWellKnownType = "internet"

    CommunityWellKnownType_local_AS CommunityWellKnownType = "local-AS"

    CommunityWellKnownType_no_advertise CommunityWellKnownType = "no-advertise"

    CommunityWellKnownType_no_export CommunityWellKnownType = "no-export"
)

// CommunityWellKnownAddType
type CommunityWellKnownAddType string

const (
    CommunityWellKnownAddType_gshut CommunityWellKnownAddType = "gshut"

    CommunityWellKnownAddType_internet CommunityWellKnownAddType = "internet"

    CommunityWellKnownAddType_local_AS CommunityWellKnownAddType = "local-AS"

    CommunityWellKnownAddType_no_advertise CommunityWellKnownAddType = "no-advertise"

    CommunityWellKnownAddType_no_export CommunityWellKnownAddType = "no-export"

    CommunityWellKnownAddType_additive CommunityWellKnownAddType = "additive"
)

// MonthType
type MonthType string

const (
    MonthType_Jan MonthType = "Jan"

    MonthType_Feb MonthType = "Feb"

    MonthType_Mar MonthType = "Mar"

    MonthType_Apr MonthType = "Apr"

    MonthType_May MonthType = "May"

    MonthType_Jun MonthType = "Jun"

    MonthType_Jul MonthType = "Jul"

    MonthType_Aug MonthType = "Aug"

    MonthType_Sep MonthType = "Sep"

    MonthType_Oct MonthType = "Oct"

    MonthType_Nov MonthType = "Nov"

    MonthType_Dec MonthType = "Dec"
)

