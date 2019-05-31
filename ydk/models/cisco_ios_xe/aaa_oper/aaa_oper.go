// This module contains a collection of YANG definitions
// for AAA operational data.
// Copyright (c) 2017-2018 by Cisco Systems, Inc.
// All rights reserved.
package aaa_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-aaa-oper aaa-data}", reflect.TypeOf(AaaData{}))
    ydk.RegisterEntity("Cisco-IOS-XE-aaa-oper:aaa-data", reflect.TypeOf(AaaData{}))
}

// AaaSessProtType represents the session is established.
type AaaSessProtType string

const (
    // No Protocol type
    AaaSessProtType_aaa_sess_proto_type_none AaaSessProtType = "aaa-sess-proto-type-none"

    // Invalid Protocol type 
    AaaSessProtType_aaa_sess_proto_type_invalid AaaSessProtType = "aaa-sess-proto-type-invalid"

    // LCP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_lcp AaaSessProtType = "aaa-sess-proto-type-lcp"

    // IP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ip AaaSessProtType = "aaa-sess-proto-type-ip"

    // IPSEC Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ipsec AaaSessProtType = "aaa-sess-proto-type-ipsec"

    // IPX Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ipx AaaSessProtType = "aaa-sess-proto-type-ipx"

    // ATALK Protocol type 
    AaaSessProtType_aaa_sess_proto_type_atalk AaaSessProtType = "aaa-sess-proto-type-atalk"

    // XREMOTE Protocol type 
    AaaSessProtType_aaa_sess_proto_type_xremote AaaSessProtType = "aaa-sess-proto-type-xremote"

    // TN3270 Protocol type 
    AaaSessProtType_aaa_sess_proto_type_tn3270 AaaSessProtType = "aaa-sess-proto-type-tn3270"

    // TELNET Protocol type 
    AaaSessProtType_aaa_sess_proto_type_telnet AaaSessProtType = "aaa-sess-proto-type-telnet"

    // TCP_CLEAR Protocol type 
    AaaSessProtType_aaa_sess_proto_type_tcp_clear AaaSessProtType = "aaa-sess-proto-type-tcp-clear"

    // RLOGIN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_rlogin AaaSessProtType = "aaa-sess-proto-type-rlogin"

    // LAT Protocol type 
    AaaSessProtType_aaa_sess_proto_type_lat AaaSessProtType = "aaa-sess-proto-type-lat"

    // PAD Protocol type 
    AaaSessProtType_aaa_sess_proto_type_pad AaaSessProtType = "aaa-sess-proto-type-pad"

    // OSICP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_osicp AaaSessProtType = "aaa-sess-proto-type-osicp"

    // TAGCP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_tagcp AaaSessProtType = "aaa-sess-proto-type-tagcp"

    // BACP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_bacp AaaSessProtType = "aaa-sess-proto-type-bacp"

    // DECNET Protocol type 
    AaaSessProtType_aaa_sess_proto_type_decnet AaaSessProtType = "aaa-sess-proto-type-decnet"

    // CCP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ccp AaaSessProtType = "aaa-sess-proto-type-ccp"

    // CDP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_cdp AaaSessProtType = "aaa-sess-proto-type-cdp"

    // BRIDGING Protocol type 
    AaaSessProtType_aaa_sess_proto_type_bridging AaaSessProtType = "aaa-sess-proto-type-bridging"

    // NBF Protocol type 
    AaaSessProtType_aaa_sess_proto_type_nbf AaaSessProtType = "aaa-sess-proto-type-nbf"

    // BAP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_bap AaaSessProtType = "aaa-sess-proto-type-bap"

    // MULTILINK Protocol type 
    AaaSessProtType_aaa_sess_proto_type_multilink AaaSessProtType = "aaa-sess-proto-type-multilink"

    // H323 Protocol type 
    AaaSessProtType_aaa_sess_proto_type_h323 AaaSessProtType = "aaa-sess-proto-type-h323"

    // UNKNOWN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_unknown AaaSessProtType = "aaa-sess-proto-type-unknown"

    // CALL ACCEPT Protocol type 
    AaaSessProtType_aaa_sess_proto_type_call_accept AaaSessProtType = "aaa-sess-proto-type-call-accept"

    // VPDN SESSION Protocol type 
    AaaSessProtType_aaa_sess_proto_type_vpdn_session AaaSessProtType = "aaa-sess-proto-type-vpdn-session"

    // RM CALL STATUS Protocol type 
    AaaSessProtType_aaa_sess_proto_type_rm_call_status AaaSessProtType = "aaa-sess-proto-type-rm-call-status"

    // RM NAS STATUS Protocol type 
    AaaSessProtType_aaa_sess_proto_type_rm_nas_status AaaSessProtType = "aaa-sess-proto-type-rm-nas-status"

    // DIAL IN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_dial_in AaaSessProtType = "aaa-sess-proto-type-dial-in"

    // DIAL OUT Protocol type 
    AaaSessProtType_aaa_sess_proto_type_dial_out AaaSessProtType = "aaa-sess-proto-type-dial-out"

    // SS7 Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ss7 AaaSessProtType = "aaa-sess-proto-type-ss7"

    // RMS STOP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_rms_stop AaaSessProtType = "aaa-sess-proto-type-rms-stop"

    // RMS START Protocol type 
    AaaSessProtType_aaa_sess_proto_type_rms_start AaaSessProtType = "aaa-sess-proto-type-rms-start"

    // VPDN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_vpdn AaaSessProtType = "aaa-sess-proto-type-vpdn"

    // SSS Protocol type 
    AaaSessProtType_aaa_sess_proto_type_sss AaaSessProtType = "aaa-sess-proto-type-sss"

    // SUBSCRIBER Protocol type 
    AaaSessProtType_aaa_sess_proto_type_subscriber AaaSessProtType = "aaa-sess-proto-type-subscriber"

    // ATM Protocol type 
    AaaSessProtType_aaa_sess_proto_type_atm AaaSessProtType = "aaa-sess-proto-type-atm"

    // SSH Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ssh AaaSessProtType = "aaa-sess-proto-type-ssh"

    // IPV6 Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ipv6 AaaSessProtType = "aaa-sess-proto-type-ipv6"

    // AIRONET Protocol type 
    AaaSessProtType_aaa_sess_proto_type_aironet AaaSessProtType = "aaa-sess-proto-type-aironet"

    // PPOE Protocol type 
    AaaSessProtType_aaa_sess_proto_type_pppoe AaaSessProtType = "aaa-sess-proto-type-pppoe"

    // ENTITY Protocol type 
    AaaSessProtType_aaa_sess_proto_type_entity AaaSessProtType = "aaa-sess-proto-type-entity"

    // CDMA Protocol type 
    AaaSessProtType_aaa_sess_proto_type_cdma AaaSessProtType = "aaa-sess-proto-type-cdma"

    // CRB Protocol type 
    AaaSessProtType_aaa_sess_proto_type_crb AaaSessProtType = "aaa-sess-proto-type-crb"

    // TEMPLATE Protocol type 
    AaaSessProtType_aaa_sess_proto_type_template AaaSessProtType = "aaa-sess-proto-type-template"

    // AAA Protocol type 
    AaaSessProtType_aaa_sess_proto_type_aaa AaaSessProtType = "aaa-sess-proto-type-aaa"

    // EPD Protocol type 
    AaaSessProtType_aaa_sess_proto_type_epd AaaSessProtType = "aaa-sess-proto-type-epd"

    // MAC Protocol type 
    AaaSessProtType_aaa_sess_proto_type_mac AaaSessProtType = "aaa-sess-proto-type-mac"

    // LEAP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_leap AaaSessProtType = "aaa-sess-proto-type-leap"

    // IGMP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_igmp AaaSessProtType = "aaa-sess-proto-type-igmp"

    // WEBVPN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_webvpn AaaSessProtType = "aaa-sess-proto-type-webvpn"

    // CTS Protocol type 
    AaaSessProtType_aaa_sess_proto_type_cts AaaSessProtType = "aaa-sess-proto-type-cts"

    // RADIUS Protocol type 
    AaaSessProtType_aaa_sess_proto_type_radius AaaSessProtType = "aaa-sess-proto-type-radius"

    // EVC Protocol type 
    AaaSessProtType_aaa_sess_proto_type_evc AaaSessProtType = "aaa-sess-proto-type-evc"

    // ELMI Protocol type 
    AaaSessProtType_aaa_sess_proto_type_elmi AaaSessProtType = "aaa-sess-proto-type-elmi"

    // DOT1X Protocol type 
    AaaSessProtType_aaa_sess_proto_type_dot1x AaaSessProtType = "aaa-sess-proto-type-dot1x"

    // DTP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_dtp AaaSessProtType = "aaa-sess-proto-type-dtp"

    // LACP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_lacp AaaSessProtType = "aaa-sess-proto-type-lacp"

    // PAGP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_pagp AaaSessProtType = "aaa-sess-proto-type-pagp"

    // STP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_stp AaaSessProtType = "aaa-sess-proto-type-stp"

    // VTP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_vtp AaaSessProtType = "aaa-sess-proto-type-vtp"

    // ETHERNET MAC TUNNEL Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ethernet_mac_tunnel AaaSessProtType = "aaa-sess-proto-type-ethernet-mac-tunnel"

    // BRIDGE DOMAIN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_bridge_domain AaaSessProtType = "aaa-sess-proto-type-bridge-domain"

    // ETHERNET CFM Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ethernet_cfm AaaSessProtType = "aaa-sess-proto-type-ethernet-cfm"

    // ETHERNET SERVICE INSTANCE Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ethernet_service_instance AaaSessProtType = "aaa-sess-proto-type-ethernet-service-instance"

    // SERVICE GROUP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_service_group AaaSessProtType = "aaa-sess-proto-type-service-group"

    // IP DHCP SNOOPING Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ip_dhcp_snooping AaaSessProtType = "aaa-sess-proto-type-ip-dhcp-snooping"

    // IP SOURCE GUARD Protocol type 
    AaaSessProtType_aaa_sess_proto_type_ip_source_guard AaaSessProtType = "aaa-sess-proto-type-ip-source-guard"

    // ERROR DISABLE Protocol type 
    AaaSessProtType_aaa_sess_proto_type_error_disable AaaSessProtType = "aaa-sess-proto-type-error-disable"

    // CMAC BRIDGE DOMAIN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_cmac_bridge_domain AaaSessProtType = "aaa-sess-proto-type-cmac-bridge-domain"

    // MAC IN MAC TUNNEL Protocol type 
    AaaSessProtType_aaa_sess_proto_type_mac_in_mac_tunnel AaaSessProtType = "aaa-sess-proto-type-mac-in-mac-tunnel"

    // L2VPN Protocol type 
    AaaSessProtType_aaa_sess_proto_type_l2vpn AaaSessProtType = "aaa-sess-proto-type-l2vpn"

    // SNMP Protocol type 
    AaaSessProtType_aaa_sess_proto_type_snmp AaaSessProtType = "aaa-sess-proto-type-snmp"
)

// AaaData
// Operational state of AAA
type AaaData struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Radius server statistics. The type is slice of AaaData_AaaRadiusStats.
    AaaRadiusStats []*AaaData_AaaRadiusStats

    // AAA TACACS server statistics. The type is slice of AaaData_AaaTacacsStats.
    AaaTacacsStats []*AaaData_AaaTacacsStats

    // LDAP server counters. The type is slice of AaaData_AaaLdapCounters.
    AaaLdapCounters []*AaaData_AaaLdapCounters

    // List of current users. The type is slice of AaaData_AaaUsers.
    AaaUsers []*AaaData_AaaUsers
}

func (aaaData *AaaData) GetEntityData() *types.CommonEntityData {
    aaaData.EntityData.YFilter = aaaData.YFilter
    aaaData.EntityData.YangName = "aaa-data"
    aaaData.EntityData.BundleName = "cisco_ios_xe"
    aaaData.EntityData.ParentYangName = "Cisco-IOS-XE-aaa-oper"
    aaaData.EntityData.SegmentPath = "Cisco-IOS-XE-aaa-oper:aaa-data"
    aaaData.EntityData.AbsolutePath = aaaData.EntityData.SegmentPath
    aaaData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaData.EntityData.Children = types.NewOrderedMap()
    aaaData.EntityData.Children.Append("aaa-radius-stats", types.YChild{"AaaRadiusStats", nil})
    for i := range aaaData.AaaRadiusStats {
        aaaData.EntityData.Children.Append(types.GetSegmentPath(aaaData.AaaRadiusStats[i]), types.YChild{"AaaRadiusStats", aaaData.AaaRadiusStats[i]})
    }
    aaaData.EntityData.Children.Append("aaa-tacacs-stats", types.YChild{"AaaTacacsStats", nil})
    for i := range aaaData.AaaTacacsStats {
        aaaData.EntityData.Children.Append(types.GetSegmentPath(aaaData.AaaTacacsStats[i]), types.YChild{"AaaTacacsStats", aaaData.AaaTacacsStats[i]})
    }
    aaaData.EntityData.Children.Append("aaa-ldap-counters", types.YChild{"AaaLdapCounters", nil})
    for i := range aaaData.AaaLdapCounters {
        aaaData.EntityData.Children.Append(types.GetSegmentPath(aaaData.AaaLdapCounters[i]), types.YChild{"AaaLdapCounters", aaaData.AaaLdapCounters[i]})
    }
    aaaData.EntityData.Children.Append("aaa-users", types.YChild{"AaaUsers", nil})
    for i := range aaaData.AaaUsers {
        aaaData.EntityData.Children.Append(types.GetSegmentPath(aaaData.AaaUsers[i]), types.YChild{"AaaUsers", aaaData.AaaUsers[i]})
    }
    aaaData.EntityData.Leafs = types.NewOrderedMap()

    aaaData.EntityData.YListKeys = []string {}

    return &(aaaData.EntityData)
}

// AaaData_AaaRadiusStats
// Radius server statistics
type AaaData_AaaRadiusStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AAA group name in which the server is defined. 
    // For public servers the group name is "PUBLIC GROUP" by default. The type is
    // string.
    GroupName interface{}

    // This attribute is a key. Radius server IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RadiusServerIp interface{}

    // This attribute is a key. Radius server auth-port. The type is interface{}
    // with range: 0..65535.
    AuthPort interface{}

    // This attribute is a key. Radius server acct-port. The type is interface{}
    // with range: 0..65535.
    AcctPort interface{}

    // Authentication retried access requests. The type is interface{} with range:
    // 0..4294967295.
    AuthenRetriedAccessRequests interface{}

    // Authentication access accepts. The type is interface{} with range:
    // 0..4294967295.
    AuthenAccessAccepts interface{}

    // Authentication access rejects. The type is interface{} with range:
    // 0..4294967295.
    AuthenAccessRejects interface{}

    // Authentication Timeout access requests. The type is interface{} with range:
    // 0..4294967295.
    AuthenTimeoutAccessRequests interface{}

    // Authorization retried access requests. The type is interface{} with range:
    // 0..4294967295.
    AuthorRetriedAccessRequests interface{}

    // Authorization access accepts. The type is interface{} with range:
    // 0..4294967295.
    AuthorAccessAccepts interface{}

    // Authorization access rejects. The type is interface{} with range:
    // 0..4294967295.
    AuthorAccessRejects interface{}

    // Authorization Timeout access requests. The type is interface{} with range:
    // 0..4294967295.
    AuthorTimeoutAccessRequests interface{}

    // Number of new connection requests sent to the RADIUS server. The type is
    // interface{} with range: 0..4294967295.
    ConnectionOpens interface{}

    // Number of connection close requests sent to the server. The type is
    // interface{} with range: 0..4294967295.
    ConnectionCloses interface{}

    // Number of connections aborted. These do not include connections that are
    // closed gracefully. The type is interface{} with range: 0..4294967295.
    ConnectionAborts interface{}

    // Number of connection failures to the RADIUS server. The type is interface{}
    // with range: 0..4294967295.
    ConnectionFailures interface{}

    // Number of connection timeouts to the RADIUS server. The type is interface{}
    // with range: 0..4294967295.
    ConnectionTimeouts interface{}

    // Number of authentication messages sent to the RADIUS server. The type is
    // interface{} with range: 0..4294967295.
    AuthenMessagesSent interface{}

    // Number of authorization messages sent to the RADIUS server. The type is
    // interface{} with range: 0..4294967295.
    AuthorMessagesSent interface{}

    // Number of accounting messages sent to the RADIUS server. The type is
    // interface{} with range: 0..4294967295.
    AcctMessagesSent interface{}

    // Number of authentication messages received by the RADIUS server. The type
    // is interface{} with range: 0..4294967295.
    AuthenMessagesReceived interface{}

    // Number of authorization messages received by the RADIUS server. The type is
    // interface{} with range: 0..4294967295.
    AuthorMessagesReceived interface{}

    // Number of authentication error messages received  from the RADIUS server.
    // The type is interface{} with range: 0..4294967295.
    AuthenErrorsReceived interface{}

    // Number of authorization error messages received  from the RADIUS server.
    // The type is interface{} with range: 0..4294967295.
    AuthorErrorsReceived interface{}

    // Number of accounting error messages received  from the RADIUS server. The
    // type is interface{} with range: 0..4294967295.
    AcctErrorsReceived interface{}

    // Time from which the statistics are valid. This field will be updated when a
    // RADIUS server is configured and also when the RADIUS server statistics are
    // cleared. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StatsTime interface{}
}

func (aaaRadiusStats *AaaData_AaaRadiusStats) GetEntityData() *types.CommonEntityData {
    aaaRadiusStats.EntityData.YFilter = aaaRadiusStats.YFilter
    aaaRadiusStats.EntityData.YangName = "aaa-radius-stats"
    aaaRadiusStats.EntityData.BundleName = "cisco_ios_xe"
    aaaRadiusStats.EntityData.ParentYangName = "aaa-data"
    aaaRadiusStats.EntityData.SegmentPath = "aaa-radius-stats" + types.AddKeyToken(aaaRadiusStats.GroupName, "group-name") + types.AddKeyToken(aaaRadiusStats.RadiusServerIp, "radius-server-ip") + types.AddKeyToken(aaaRadiusStats.AuthPort, "auth-port") + types.AddKeyToken(aaaRadiusStats.AcctPort, "acct-port")
    aaaRadiusStats.EntityData.AbsolutePath = "Cisco-IOS-XE-aaa-oper:aaa-data/" + aaaRadiusStats.EntityData.SegmentPath
    aaaRadiusStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaRadiusStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaRadiusStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaRadiusStats.EntityData.Children = types.NewOrderedMap()
    aaaRadiusStats.EntityData.Leafs = types.NewOrderedMap()
    aaaRadiusStats.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", aaaRadiusStats.GroupName})
    aaaRadiusStats.EntityData.Leafs.Append("radius-server-ip", types.YLeaf{"RadiusServerIp", aaaRadiusStats.RadiusServerIp})
    aaaRadiusStats.EntityData.Leafs.Append("auth-port", types.YLeaf{"AuthPort", aaaRadiusStats.AuthPort})
    aaaRadiusStats.EntityData.Leafs.Append("acct-port", types.YLeaf{"AcctPort", aaaRadiusStats.AcctPort})
    aaaRadiusStats.EntityData.Leafs.Append("authen-retried-access-requests", types.YLeaf{"AuthenRetriedAccessRequests", aaaRadiusStats.AuthenRetriedAccessRequests})
    aaaRadiusStats.EntityData.Leafs.Append("authen-access-accepts", types.YLeaf{"AuthenAccessAccepts", aaaRadiusStats.AuthenAccessAccepts})
    aaaRadiusStats.EntityData.Leafs.Append("authen-access-rejects", types.YLeaf{"AuthenAccessRejects", aaaRadiusStats.AuthenAccessRejects})
    aaaRadiusStats.EntityData.Leafs.Append("authen-timeout-access-requests", types.YLeaf{"AuthenTimeoutAccessRequests", aaaRadiusStats.AuthenTimeoutAccessRequests})
    aaaRadiusStats.EntityData.Leafs.Append("author-retried-access-requests", types.YLeaf{"AuthorRetriedAccessRequests", aaaRadiusStats.AuthorRetriedAccessRequests})
    aaaRadiusStats.EntityData.Leafs.Append("author-access-accepts", types.YLeaf{"AuthorAccessAccepts", aaaRadiusStats.AuthorAccessAccepts})
    aaaRadiusStats.EntityData.Leafs.Append("author-access-rejects", types.YLeaf{"AuthorAccessRejects", aaaRadiusStats.AuthorAccessRejects})
    aaaRadiusStats.EntityData.Leafs.Append("author-timeout-access-requests", types.YLeaf{"AuthorTimeoutAccessRequests", aaaRadiusStats.AuthorTimeoutAccessRequests})
    aaaRadiusStats.EntityData.Leafs.Append("connection-opens", types.YLeaf{"ConnectionOpens", aaaRadiusStats.ConnectionOpens})
    aaaRadiusStats.EntityData.Leafs.Append("connection-closes", types.YLeaf{"ConnectionCloses", aaaRadiusStats.ConnectionCloses})
    aaaRadiusStats.EntityData.Leafs.Append("connection-aborts", types.YLeaf{"ConnectionAborts", aaaRadiusStats.ConnectionAborts})
    aaaRadiusStats.EntityData.Leafs.Append("connection-failures", types.YLeaf{"ConnectionFailures", aaaRadiusStats.ConnectionFailures})
    aaaRadiusStats.EntityData.Leafs.Append("connection-timeouts", types.YLeaf{"ConnectionTimeouts", aaaRadiusStats.ConnectionTimeouts})
    aaaRadiusStats.EntityData.Leafs.Append("authen-messages-sent", types.YLeaf{"AuthenMessagesSent", aaaRadiusStats.AuthenMessagesSent})
    aaaRadiusStats.EntityData.Leafs.Append("author-messages-sent", types.YLeaf{"AuthorMessagesSent", aaaRadiusStats.AuthorMessagesSent})
    aaaRadiusStats.EntityData.Leafs.Append("acct-messages-sent", types.YLeaf{"AcctMessagesSent", aaaRadiusStats.AcctMessagesSent})
    aaaRadiusStats.EntityData.Leafs.Append("authen-messages-received", types.YLeaf{"AuthenMessagesReceived", aaaRadiusStats.AuthenMessagesReceived})
    aaaRadiusStats.EntityData.Leafs.Append("author-messages-received", types.YLeaf{"AuthorMessagesReceived", aaaRadiusStats.AuthorMessagesReceived})
    aaaRadiusStats.EntityData.Leafs.Append("authen-errors-received", types.YLeaf{"AuthenErrorsReceived", aaaRadiusStats.AuthenErrorsReceived})
    aaaRadiusStats.EntityData.Leafs.Append("author-errors-received", types.YLeaf{"AuthorErrorsReceived", aaaRadiusStats.AuthorErrorsReceived})
    aaaRadiusStats.EntityData.Leafs.Append("acct-errors-received", types.YLeaf{"AcctErrorsReceived", aaaRadiusStats.AcctErrorsReceived})
    aaaRadiusStats.EntityData.Leafs.Append("stats-time", types.YLeaf{"StatsTime", aaaRadiusStats.StatsTime})

    aaaRadiusStats.EntityData.YListKeys = []string {"GroupName", "RadiusServerIp", "AuthPort", "AcctPort"}

    return &(aaaRadiusStats.EntityData)
}

// AaaData_AaaTacacsStats
// AAA TACACS server statistics
type AaaData_AaaTacacsStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AAA group name in which the server is defined. For
    // public servers the group name is "PUBLIC GROUP" by default. The type is
    // string.
    GroupName interface{}

    // This attribute is a key. TACACS server IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    TacacsServerAddress interface{}

    // This attribute is a key. TACACS server port. The type is interface{} with
    // range: 0..65535.
    Port interface{}

    // Number of new connection requests sent to the server. The type is
    // interface{} with range: 0..4294967295.
    ConnectionOpens interface{}

    // Number of connection close requests sent to the server. The type is
    // interface{} with range: 0..4294967295.
    ConnectionCloses interface{}

    // Number of aborted connections to the server. These do not include
    // connections that are close gracefully. The type is interface{} with range:
    // 0..4294967295.
    ConnectionAborts interface{}

    // Number of connection failures to the server. The type is interface{} with
    // range: 0..4294967295.
    ConnectionFailures interface{}

    // Number of connection timeouts to the server. The type is interface{} with
    // range: 0..4294967295.
    ConnectionTimeouts interface{}

    // Number of messages sent to the server. The type is interface{} with range:
    // 0..4294967295.
    MessagesSent interface{}

    // Number of messages received by the server. The type is interface{} with
    // range: 0..4294967295.
    MessagesReceived interface{}

    // Number of error messages received from the server. The type is interface{}
    // with range: 0..4294967295.
    ErrorsReceived interface{}

    // This attribute contains stats collection start time. Stats collection
    // starts when the TACACS server is configured. The type is string with
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StatsStartTime interface{}
}

func (aaaTacacsStats *AaaData_AaaTacacsStats) GetEntityData() *types.CommonEntityData {
    aaaTacacsStats.EntityData.YFilter = aaaTacacsStats.YFilter
    aaaTacacsStats.EntityData.YangName = "aaa-tacacs-stats"
    aaaTacacsStats.EntityData.BundleName = "cisco_ios_xe"
    aaaTacacsStats.EntityData.ParentYangName = "aaa-data"
    aaaTacacsStats.EntityData.SegmentPath = "aaa-tacacs-stats" + types.AddKeyToken(aaaTacacsStats.GroupName, "group-name") + types.AddKeyToken(aaaTacacsStats.TacacsServerAddress, "tacacs-server-address") + types.AddKeyToken(aaaTacacsStats.Port, "port")
    aaaTacacsStats.EntityData.AbsolutePath = "Cisco-IOS-XE-aaa-oper:aaa-data/" + aaaTacacsStats.EntityData.SegmentPath
    aaaTacacsStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaTacacsStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaTacacsStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaTacacsStats.EntityData.Children = types.NewOrderedMap()
    aaaTacacsStats.EntityData.Leafs = types.NewOrderedMap()
    aaaTacacsStats.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", aaaTacacsStats.GroupName})
    aaaTacacsStats.EntityData.Leafs.Append("tacacs-server-address", types.YLeaf{"TacacsServerAddress", aaaTacacsStats.TacacsServerAddress})
    aaaTacacsStats.EntityData.Leafs.Append("port", types.YLeaf{"Port", aaaTacacsStats.Port})
    aaaTacacsStats.EntityData.Leafs.Append("connection-opens", types.YLeaf{"ConnectionOpens", aaaTacacsStats.ConnectionOpens})
    aaaTacacsStats.EntityData.Leafs.Append("connection-closes", types.YLeaf{"ConnectionCloses", aaaTacacsStats.ConnectionCloses})
    aaaTacacsStats.EntityData.Leafs.Append("connection-aborts", types.YLeaf{"ConnectionAborts", aaaTacacsStats.ConnectionAborts})
    aaaTacacsStats.EntityData.Leafs.Append("connection-failures", types.YLeaf{"ConnectionFailures", aaaTacacsStats.ConnectionFailures})
    aaaTacacsStats.EntityData.Leafs.Append("connection-timeouts", types.YLeaf{"ConnectionTimeouts", aaaTacacsStats.ConnectionTimeouts})
    aaaTacacsStats.EntityData.Leafs.Append("messages-sent", types.YLeaf{"MessagesSent", aaaTacacsStats.MessagesSent})
    aaaTacacsStats.EntityData.Leafs.Append("messages-received", types.YLeaf{"MessagesReceived", aaaTacacsStats.MessagesReceived})
    aaaTacacsStats.EntityData.Leafs.Append("errors-received", types.YLeaf{"ErrorsReceived", aaaTacacsStats.ErrorsReceived})
    aaaTacacsStats.EntityData.Leafs.Append("stats-start-time", types.YLeaf{"StatsStartTime", aaaTacacsStats.StatsStartTime})

    aaaTacacsStats.EntityData.YListKeys = []string {"GroupName", "TacacsServerAddress", "Port"}

    return &(aaaTacacsStats.EntityData)
}

// AaaData_AaaLdapCounters
// LDAP server counters
type AaaData_AaaLdapCounters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. LDAP server IP address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    LdapServerAddress interface{}

    // This attribute is a key. LDAP server listening port - TCP. The type is
    // interface{} with range: 0..65535.
    LdapServerPort interface{}

    // Number of new connection requests sent to the LDAP server. The type is
    // interface{} with range: 0..4294967295.
    ConnectionOpens interface{}

    // Number of messages sent to the LDAP server. The type is interface{} with
    // range: 0..4294967295.
    MessagesSent interface{}

    // Number of messages received by the LDAP server. The type is interface{}
    // with range: 0..4294967295.
    MessagesReceived interface{}

    // Number of error messages received from the LDAP server. The type is
    // interface{} with range: 0..4294967295.
    ErrorsReceived interface{}

    // Number of connection close requests sent to the server. The type is
    // interface{} with range: 0..4294967295.
    ConnectionCloses interface{}

    // Number of connections aborted. These do not include connections that are
    // close gracefully. The type is interface{} with range: 0..4294967295.
    ConnectionAborts interface{}

    // Number of connection failures to the LDAP server. The type is interface{}
    // with range: 0..4294967295.
    ConnectionFailures interface{}

    // Number of connection timeouts to the LDAP server. The type is interface{}
    // with range: 0..4294967295.
    ConnectionTimeouts interface{}

    // This attribute contains LDAP counters collection start time. Counters
    // collection starts when a LDAP server is configured. Counters collection
    // will be reset when the LDAP server counters are cleared. The type is string
    // with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    CountersStartTime interface{}
}

func (aaaLdapCounters *AaaData_AaaLdapCounters) GetEntityData() *types.CommonEntityData {
    aaaLdapCounters.EntityData.YFilter = aaaLdapCounters.YFilter
    aaaLdapCounters.EntityData.YangName = "aaa-ldap-counters"
    aaaLdapCounters.EntityData.BundleName = "cisco_ios_xe"
    aaaLdapCounters.EntityData.ParentYangName = "aaa-data"
    aaaLdapCounters.EntityData.SegmentPath = "aaa-ldap-counters" + types.AddKeyToken(aaaLdapCounters.LdapServerAddress, "ldap-server-address") + types.AddKeyToken(aaaLdapCounters.LdapServerPort, "ldap-server-port")
    aaaLdapCounters.EntityData.AbsolutePath = "Cisco-IOS-XE-aaa-oper:aaa-data/" + aaaLdapCounters.EntityData.SegmentPath
    aaaLdapCounters.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaLdapCounters.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaLdapCounters.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaLdapCounters.EntityData.Children = types.NewOrderedMap()
    aaaLdapCounters.EntityData.Leafs = types.NewOrderedMap()
    aaaLdapCounters.EntityData.Leafs.Append("ldap-server-address", types.YLeaf{"LdapServerAddress", aaaLdapCounters.LdapServerAddress})
    aaaLdapCounters.EntityData.Leafs.Append("ldap-server-port", types.YLeaf{"LdapServerPort", aaaLdapCounters.LdapServerPort})
    aaaLdapCounters.EntityData.Leafs.Append("connection-opens", types.YLeaf{"ConnectionOpens", aaaLdapCounters.ConnectionOpens})
    aaaLdapCounters.EntityData.Leafs.Append("messages-sent", types.YLeaf{"MessagesSent", aaaLdapCounters.MessagesSent})
    aaaLdapCounters.EntityData.Leafs.Append("messages-received", types.YLeaf{"MessagesReceived", aaaLdapCounters.MessagesReceived})
    aaaLdapCounters.EntityData.Leafs.Append("errors-received", types.YLeaf{"ErrorsReceived", aaaLdapCounters.ErrorsReceived})
    aaaLdapCounters.EntityData.Leafs.Append("connection-closes", types.YLeaf{"ConnectionCloses", aaaLdapCounters.ConnectionCloses})
    aaaLdapCounters.EntityData.Leafs.Append("connection-aborts", types.YLeaf{"ConnectionAborts", aaaLdapCounters.ConnectionAborts})
    aaaLdapCounters.EntityData.Leafs.Append("connection-failures", types.YLeaf{"ConnectionFailures", aaaLdapCounters.ConnectionFailures})
    aaaLdapCounters.EntityData.Leafs.Append("connection-timeouts", types.YLeaf{"ConnectionTimeouts", aaaLdapCounters.ConnectionTimeouts})
    aaaLdapCounters.EntityData.Leafs.Append("counters-start-time", types.YLeaf{"CountersStartTime", aaaLdapCounters.CountersStartTime})

    aaaLdapCounters.EntityData.YListKeys = []string {"LdapServerAddress", "LdapServerPort"}

    return &(aaaLdapCounters.EntityData)
}

// AaaData_AaaUsers
// List of current users
type AaaData_AaaUsers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The username used to logged into the device. The
    // type is string.
    Username interface{}

    // Sessions associated with the users. The type is slice of
    // AaaData_AaaUsers_AaaSessions.
    AaaSessions []*AaaData_AaaUsers_AaaSessions
}

func (aaaUsers *AaaData_AaaUsers) GetEntityData() *types.CommonEntityData {
    aaaUsers.EntityData.YFilter = aaaUsers.YFilter
    aaaUsers.EntityData.YangName = "aaa-users"
    aaaUsers.EntityData.BundleName = "cisco_ios_xe"
    aaaUsers.EntityData.ParentYangName = "aaa-data"
    aaaUsers.EntityData.SegmentPath = "aaa-users" + types.AddKeyToken(aaaUsers.Username, "username")
    aaaUsers.EntityData.AbsolutePath = "Cisco-IOS-XE-aaa-oper:aaa-data/" + aaaUsers.EntityData.SegmentPath
    aaaUsers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaUsers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaUsers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaUsers.EntityData.Children = types.NewOrderedMap()
    aaaUsers.EntityData.Children.Append("aaa-sessions", types.YChild{"AaaSessions", nil})
    for i := range aaaUsers.AaaSessions {
        aaaUsers.EntityData.Children.Append(types.GetSegmentPath(aaaUsers.AaaSessions[i]), types.YChild{"AaaSessions", aaaUsers.AaaSessions[i]})
    }
    aaaUsers.EntityData.Leafs = types.NewOrderedMap()
    aaaUsers.EntityData.Leafs.Append("username", types.YLeaf{"Username", aaaUsers.Username})

    aaaUsers.EntityData.YListKeys = []string {"Username"}

    return &(aaaUsers.EntityData)
}

// AaaData_AaaUsers_AaaSessions
// Sessions associated with the users
type AaaData_AaaUsers_AaaSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. AAA Unique ID. The type is interface{} with range:
    // 0..4294967295.
    AaaUid interface{}

    // AAA Session ID. The type is interface{} with range: 0..4294967295.
    SessionId interface{}

    // Source IP address that initiated the session. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddr interface{}

    // AAA protocol type Protocol used in this session. The type is
    // AaaSessProtType.
    Protocol interface{}

    // Login-time for this session present in aaa code. The type is string with
    // pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    LoginTime interface{}
}

func (aaaSessions *AaaData_AaaUsers_AaaSessions) GetEntityData() *types.CommonEntityData {
    aaaSessions.EntityData.YFilter = aaaSessions.YFilter
    aaaSessions.EntityData.YangName = "aaa-sessions"
    aaaSessions.EntityData.BundleName = "cisco_ios_xe"
    aaaSessions.EntityData.ParentYangName = "aaa-users"
    aaaSessions.EntityData.SegmentPath = "aaa-sessions" + types.AddKeyToken(aaaSessions.AaaUid, "aaa-uid")
    aaaSessions.EntityData.AbsolutePath = "Cisco-IOS-XE-aaa-oper:aaa-data/aaa-users/" + aaaSessions.EntityData.SegmentPath
    aaaSessions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaSessions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaSessions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaSessions.EntityData.Children = types.NewOrderedMap()
    aaaSessions.EntityData.Leafs = types.NewOrderedMap()
    aaaSessions.EntityData.Leafs.Append("aaa-uid", types.YLeaf{"AaaUid", aaaSessions.AaaUid})
    aaaSessions.EntityData.Leafs.Append("session-id", types.YLeaf{"SessionId", aaaSessions.SessionId})
    aaaSessions.EntityData.Leafs.Append("ip-addr", types.YLeaf{"IpAddr", aaaSessions.IpAddr})
    aaaSessions.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", aaaSessions.Protocol})
    aaaSessions.EntityData.Leafs.Append("login-time", types.YLeaf{"LoginTime", aaaSessions.LoginTime})

    aaaSessions.EntityData.YListKeys = []string {"AaaUid"}

    return &(aaaSessions.EntityData)
}

