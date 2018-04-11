// This MIB module defines textual conventions that
// are commonly used in modeling management information 
// pertaining to configuration, status and activity
// of firewalls.
package cisco_firewall_tc

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_firewall_tc"))
}

// CFWNetworkProtocol represents     Denotes Transmission Control Protocol.
type CFWNetworkProtocol string

const (
    CFWNetworkProtocol_none CFWNetworkProtocol = "none"

    CFWNetworkProtocol_other CFWNetworkProtocol = "other"

    CFWNetworkProtocol_ip CFWNetworkProtocol = "ip"

    CFWNetworkProtocol_icmp CFWNetworkProtocol = "icmp"

    CFWNetworkProtocol_gre CFWNetworkProtocol = "gre"

    CFWNetworkProtocol_udp CFWNetworkProtocol = "udp"

    CFWNetworkProtocol_tcp CFWNetworkProtocol = "tcp"
)

// CFWApplicationProtocol represents     Denotes any protocol not listed.
type CFWApplicationProtocol string

const (
    CFWApplicationProtocol_none CFWApplicationProtocol = "none"

    CFWApplicationProtocol_other CFWApplicationProtocol = "other"

    CFWApplicationProtocol_ftp CFWApplicationProtocol = "ftp"

    CFWApplicationProtocol_telnet CFWApplicationProtocol = "telnet"

    CFWApplicationProtocol_smtp CFWApplicationProtocol = "smtp"

    CFWApplicationProtocol_http CFWApplicationProtocol = "http"

    CFWApplicationProtocol_tacacs CFWApplicationProtocol = "tacacs"

    CFWApplicationProtocol_dns CFWApplicationProtocol = "dns"

    CFWApplicationProtocol_sqlnet CFWApplicationProtocol = "sqlnet"

    CFWApplicationProtocol_https CFWApplicationProtocol = "https"

    CFWApplicationProtocol_tftp CFWApplicationProtocol = "tftp"

    CFWApplicationProtocol_gopher CFWApplicationProtocol = "gopher"

    CFWApplicationProtocol_finger CFWApplicationProtocol = "finger"

    CFWApplicationProtocol_kerberos CFWApplicationProtocol = "kerberos"

    CFWApplicationProtocol_pop2 CFWApplicationProtocol = "pop2"

    CFWApplicationProtocol_pop3 CFWApplicationProtocol = "pop3"

    CFWApplicationProtocol_sunRpc CFWApplicationProtocol = "sunRpc"

    CFWApplicationProtocol_msRpc CFWApplicationProtocol = "msRpc"

    CFWApplicationProtocol_nntp CFWApplicationProtocol = "nntp"

    CFWApplicationProtocol_snmp CFWApplicationProtocol = "snmp"

    CFWApplicationProtocol_imap CFWApplicationProtocol = "imap"

    CFWApplicationProtocol_ldap CFWApplicationProtocol = "ldap"

    CFWApplicationProtocol_exec CFWApplicationProtocol = "exec"

    CFWApplicationProtocol_login CFWApplicationProtocol = "login"

    CFWApplicationProtocol_shell CFWApplicationProtocol = "shell"

    CFWApplicationProtocol_msSql CFWApplicationProtocol = "msSql"

    CFWApplicationProtocol_sybaseSql CFWApplicationProtocol = "sybaseSql"

    CFWApplicationProtocol_nfs CFWApplicationProtocol = "nfs"

    CFWApplicationProtocol_lotusnote CFWApplicationProtocol = "lotusnote"

    CFWApplicationProtocol_h323 CFWApplicationProtocol = "h323"

    CFWApplicationProtocol_cuseeme CFWApplicationProtocol = "cuseeme"

    CFWApplicationProtocol_realmedia CFWApplicationProtocol = "realmedia"

    CFWApplicationProtocol_netshow CFWApplicationProtocol = "netshow"

    CFWApplicationProtocol_streamworks CFWApplicationProtocol = "streamworks"

    CFWApplicationProtocol_vdolive CFWApplicationProtocol = "vdolive"

    CFWApplicationProtocol_sap CFWApplicationProtocol = "sap"

    CFWApplicationProtocol_sip CFWApplicationProtocol = "sip"

    CFWApplicationProtocol_mgcp CFWApplicationProtocol = "mgcp"

    CFWApplicationProtocol_rtsp CFWApplicationProtocol = "rtsp"

    CFWApplicationProtocol_skinny CFWApplicationProtocol = "skinny"

    CFWApplicationProtocol_gtpV0 CFWApplicationProtocol = "gtpV0"

    CFWApplicationProtocol_gtpV1 CFWApplicationProtocol = "gtpV1"

    CFWApplicationProtocol_echo CFWApplicationProtocol = "echo"

    CFWApplicationProtocol_discard CFWApplicationProtocol = "discard"

    CFWApplicationProtocol_daytime CFWApplicationProtocol = "daytime"

    CFWApplicationProtocol_netstat CFWApplicationProtocol = "netstat"

    CFWApplicationProtocol_ssh CFWApplicationProtocol = "ssh"

    CFWApplicationProtocol_time CFWApplicationProtocol = "time"

    CFWApplicationProtocol_tacacsDs CFWApplicationProtocol = "tacacsDs"

    CFWApplicationProtocol_bootps CFWApplicationProtocol = "bootps"

    CFWApplicationProtocol_bootpc CFWApplicationProtocol = "bootpc"

    CFWApplicationProtocol_dnsix CFWApplicationProtocol = "dnsix"

    CFWApplicationProtocol_rtelnet CFWApplicationProtocol = "rtelnet"

    CFWApplicationProtocol_ident CFWApplicationProtocol = "ident"

    CFWApplicationProtocol_sqlServ CFWApplicationProtocol = "sqlServ"

    CFWApplicationProtocol_ntp CFWApplicationProtocol = "ntp"

    CFWApplicationProtocol_pwdgen CFWApplicationProtocol = "pwdgen"

    CFWApplicationProtocol_ciscoFna CFWApplicationProtocol = "ciscoFna"

    CFWApplicationProtocol_ciscoTna CFWApplicationProtocol = "ciscoTna"

    CFWApplicationProtocol_ciscoSys CFWApplicationProtocol = "ciscoSys"

    CFWApplicationProtocol_netbiosNs CFWApplicationProtocol = "netbiosNs"

    CFWApplicationProtocol_netbiosDgm CFWApplicationProtocol = "netbiosDgm"

    CFWApplicationProtocol_netbiosSsn CFWApplicationProtocol = "netbiosSsn"

    CFWApplicationProtocol_sqlSrv CFWApplicationProtocol = "sqlSrv"

    CFWApplicationProtocol_snmpTrap CFWApplicationProtocol = "snmpTrap"

    CFWApplicationProtocol_rsvd CFWApplicationProtocol = "rsvd"

    CFWApplicationProtocol_send CFWApplicationProtocol = "send"

    CFWApplicationProtocol_xdmcp CFWApplicationProtocol = "xdmcp"

    CFWApplicationProtocol_bgp CFWApplicationProtocol = "bgp"

    CFWApplicationProtocol_irc CFWApplicationProtocol = "irc"

    CFWApplicationProtocol_qmtp CFWApplicationProtocol = "qmtp"

    CFWApplicationProtocol_ipx CFWApplicationProtocol = "ipx"

    CFWApplicationProtocol_dbase CFWApplicationProtocol = "dbase"

    CFWApplicationProtocol_imap3 CFWApplicationProtocol = "imap3"

    CFWApplicationProtocol_rsvpTunnel CFWApplicationProtocol = "rsvpTunnel"

    CFWApplicationProtocol_hpCollector CFWApplicationProtocol = "hpCollector"

    CFWApplicationProtocol_hpManagedNode CFWApplicationProtocol = "hpManagedNode"

    CFWApplicationProtocol_hpAlarmMgr CFWApplicationProtocol = "hpAlarmMgr"

    CFWApplicationProtocol_microsoftDs CFWApplicationProtocol = "microsoftDs"

    CFWApplicationProtocol_creativeServer CFWApplicationProtocol = "creativeServer"

    CFWApplicationProtocol_creativePartnr CFWApplicationProtocol = "creativePartnr"

    CFWApplicationProtocol_appleQtc CFWApplicationProtocol = "appleQtc"

    CFWApplicationProtocol_igmpV3Lite CFWApplicationProtocol = "igmpV3Lite"

    CFWApplicationProtocol_isakmp CFWApplicationProtocol = "isakmp"

    CFWApplicationProtocol_biff CFWApplicationProtocol = "biff"

    CFWApplicationProtocol_who CFWApplicationProtocol = "who"

    CFWApplicationProtocol_syslog CFWApplicationProtocol = "syslog"

    CFWApplicationProtocol_router CFWApplicationProtocol = "router"

    CFWApplicationProtocol_ncp CFWApplicationProtocol = "ncp"

    CFWApplicationProtocol_timed CFWApplicationProtocol = "timed"

    CFWApplicationProtocol_ircServ CFWApplicationProtocol = "ircServ"

    CFWApplicationProtocol_uucp CFWApplicationProtocol = "uucp"

    CFWApplicationProtocol_syslogConn CFWApplicationProtocol = "syslogConn"

    CFWApplicationProtocol_sshell CFWApplicationProtocol = "sshell"

    CFWApplicationProtocol_ldaps CFWApplicationProtocol = "ldaps"

    CFWApplicationProtocol_dhcpFailover CFWApplicationProtocol = "dhcpFailover"

    CFWApplicationProtocol_msexchRouting CFWApplicationProtocol = "msexchRouting"

    CFWApplicationProtocol_entrustSvcs CFWApplicationProtocol = "entrustSvcs"

    CFWApplicationProtocol_entrustSvcHandler CFWApplicationProtocol = "entrustSvcHandler"

    CFWApplicationProtocol_ciscoTdp CFWApplicationProtocol = "ciscoTdp"

    CFWApplicationProtocol_webster CFWApplicationProtocol = "webster"

    CFWApplicationProtocol_gdoi CFWApplicationProtocol = "gdoi"

    CFWApplicationProtocol_iscsi CFWApplicationProtocol = "iscsi"

    CFWApplicationProtocol_cddbp CFWApplicationProtocol = "cddbp"

    CFWApplicationProtocol_ftps CFWApplicationProtocol = "ftps"

    CFWApplicationProtocol_telnets CFWApplicationProtocol = "telnets"

    CFWApplicationProtocol_imaps CFWApplicationProtocol = "imaps"

    CFWApplicationProtocol_ircs CFWApplicationProtocol = "ircs"

    CFWApplicationProtocol_pop3s CFWApplicationProtocol = "pop3s"

    CFWApplicationProtocol_socks CFWApplicationProtocol = "socks"

    CFWApplicationProtocol_kazaa CFWApplicationProtocol = "kazaa"

    CFWApplicationProtocol_msSqlM CFWApplicationProtocol = "msSqlM"

    CFWApplicationProtocol_msSna CFWApplicationProtocol = "msSna"

    CFWApplicationProtocol_wins CFWApplicationProtocol = "wins"

    CFWApplicationProtocol_ica CFWApplicationProtocol = "ica"

    CFWApplicationProtocol_orasrv CFWApplicationProtocol = "orasrv"

    CFWApplicationProtocol_rdbDbsDisp CFWApplicationProtocol = "rdbDbsDisp"

    CFWApplicationProtocol_vqp CFWApplicationProtocol = "vqp"

    CFWApplicationProtocol_icabrowser CFWApplicationProtocol = "icabrowser"

    CFWApplicationProtocol_kermit CFWApplicationProtocol = "kermit"

    CFWApplicationProtocol_rsvpEncap CFWApplicationProtocol = "rsvpEncap"

    CFWApplicationProtocol_l2tp CFWApplicationProtocol = "l2tp"

    CFWApplicationProtocol_pptp CFWApplicationProtocol = "pptp"

    CFWApplicationProtocol_h323Gatestat CFWApplicationProtocol = "h323Gatestat"

    CFWApplicationProtocol_rWinsock CFWApplicationProtocol = "rWinsock"

    CFWApplicationProtocol_radius CFWApplicationProtocol = "radius"

    CFWApplicationProtocol_hsrp CFWApplicationProtocol = "hsrp"

    CFWApplicationProtocol_net8Cman CFWApplicationProtocol = "net8Cman"

    CFWApplicationProtocol_oracleEmVp CFWApplicationProtocol = "oracleEmVp"

    CFWApplicationProtocol_oracleNames CFWApplicationProtocol = "oracleNames"

    CFWApplicationProtocol_oracle CFWApplicationProtocol = "oracle"

    CFWApplicationProtocol_ciscoSvcs CFWApplicationProtocol = "ciscoSvcs"

    CFWApplicationProtocol_ciscoNetMgmt CFWApplicationProtocol = "ciscoNetMgmt"

    CFWApplicationProtocol_stun CFWApplicationProtocol = "stun"

    CFWApplicationProtocol_trRsrb CFWApplicationProtocol = "trRsrb"

    CFWApplicationProtocol_ddnsV3 CFWApplicationProtocol = "ddnsV3"

    CFWApplicationProtocol_aceSvr CFWApplicationProtocol = "aceSvr"

    CFWApplicationProtocol_giop CFWApplicationProtocol = "giop"

    CFWApplicationProtocol_ttc CFWApplicationProtocol = "ttc"

    CFWApplicationProtocol_ipass CFWApplicationProtocol = "ipass"

    CFWApplicationProtocol_clp CFWApplicationProtocol = "clp"

    CFWApplicationProtocol_citrixImaClient CFWApplicationProtocol = "citrixImaClient"

    CFWApplicationProtocol_sms CFWApplicationProtocol = "sms"

    CFWApplicationProtocol_citrix CFWApplicationProtocol = "citrix"

    CFWApplicationProtocol_realSecure CFWApplicationProtocol = "realSecure"

    CFWApplicationProtocol_lotusMtap CFWApplicationProtocol = "lotusMtap"

    CFWApplicationProtocol_cifs CFWApplicationProtocol = "cifs"

    CFWApplicationProtocol_msDotnetster CFWApplicationProtocol = "msDotnetster"

    CFWApplicationProtocol_tarantella CFWApplicationProtocol = "tarantella"

    CFWApplicationProtocol_fcipPort CFWApplicationProtocol = "fcipPort"

    CFWApplicationProtocol_ssp CFWApplicationProtocol = "ssp"

    CFWApplicationProtocol_iscsiTarget CFWApplicationProtocol = "iscsiTarget"

    CFWApplicationProtocol_mySql CFWApplicationProtocol = "mySql"

    CFWApplicationProtocol_msClusterNet CFWApplicationProtocol = "msClusterNet"

    CFWApplicationProtocol_ldapAdmin CFWApplicationProtocol = "ldapAdmin"

    CFWApplicationProtocol_ieee80211Iapp CFWApplicationProtocol = "ieee80211Iapp"

    CFWApplicationProtocol_oemAgent CFWApplicationProtocol = "oemAgent"

    CFWApplicationProtocol_rtcPmPort CFWApplicationProtocol = "rtcPmPort"

    CFWApplicationProtocol_dbControlAgent CFWApplicationProtocol = "dbControlAgent"

    CFWApplicationProtocol_ipsecMsft CFWApplicationProtocol = "ipsecMsft"

    CFWApplicationProtocol_sipTls CFWApplicationProtocol = "sipTls"

    CFWApplicationProtocol_aim CFWApplicationProtocol = "aim"

    CFWApplicationProtocol_pcAnyWhereData CFWApplicationProtocol = "pcAnyWhereData"

    CFWApplicationProtocol_pcAnyWhereStat CFWApplicationProtocol = "pcAnyWhereStat"

    CFWApplicationProtocol_x11 CFWApplicationProtocol = "x11"

    CFWApplicationProtocol_ircu CFWApplicationProtocol = "ircu"

    CFWApplicationProtocol_n2h2Server CFWApplicationProtocol = "n2h2Server"

    CFWApplicationProtocol_h323CallSigAlt CFWApplicationProtocol = "h323CallSigAlt"

    CFWApplicationProtocol_yahooMsgr CFWApplicationProtocol = "yahooMsgr"

    CFWApplicationProtocol_msnMsgr CFWApplicationProtocol = "msnMsgr"
)

// CFWPolicyTargetType represents     and Forwarding (VRFs) defined by IOS.
type CFWPolicyTargetType string

const (
    CFWPolicyTargetType_all CFWPolicyTargetType = "all"

    CFWPolicyTargetType_other CFWPolicyTargetType = "other"

    CFWPolicyTargetType_interface_ CFWPolicyTargetType = "interface"

    CFWPolicyTargetType_zone CFWPolicyTargetType = "zone"

    CFWPolicyTargetType_zonepair CFWPolicyTargetType = "zonepair"

    CFWPolicyTargetType_user CFWPolicyTargetType = "user"

    CFWPolicyTargetType_usergroup CFWPolicyTargetType = "usergroup"

    CFWPolicyTargetType_context CFWPolicyTargetType = "context"
)

// CFWUrlfVendorId represents     http://www.n2h2.com
type CFWUrlfVendorId string

const (
    CFWUrlfVendorId_other CFWUrlfVendorId = "other"

    CFWUrlfVendorId_websense CFWUrlfVendorId = "websense"

    CFWUrlfVendorId_n2h2 CFWUrlfVendorId = "n2h2"
)

// CFWUrlServerStatus represents     cannot be determined
type CFWUrlServerStatus string

const (
    CFWUrlServerStatus_online CFWUrlServerStatus = "online"

    CFWUrlServerStatus_offline CFWUrlServerStatus = "offline"

    CFWUrlServerStatus_indeterminate CFWUrlServerStatus = "indeterminate"
)

