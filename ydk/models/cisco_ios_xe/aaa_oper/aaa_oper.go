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

    // List of current users. The type is slice of AaaData_AaaUsers.
    AaaUsers []AaaData_AaaUsers
}

func (aaaData *AaaData) GetEntityData() *types.CommonEntityData {
    aaaData.EntityData.YFilter = aaaData.YFilter
    aaaData.EntityData.YangName = "aaa-data"
    aaaData.EntityData.BundleName = "cisco_ios_xe"
    aaaData.EntityData.ParentYangName = "Cisco-IOS-XE-aaa-oper"
    aaaData.EntityData.SegmentPath = "Cisco-IOS-XE-aaa-oper:aaa-data"
    aaaData.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaData.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaData.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaData.EntityData.Children = make(map[string]types.YChild)
    aaaData.EntityData.Children["aaa-users"] = types.YChild{"AaaUsers", nil}
    for i := range aaaData.AaaUsers {
        aaaData.EntityData.Children[types.GetSegmentPath(&aaaData.AaaUsers[i])] = types.YChild{"AaaUsers", &aaaData.AaaUsers[i]}
    }
    aaaData.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aaaData.EntityData)
}

// AaaData_AaaUsers
// List of current users
type AaaData_AaaUsers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The username used to logged into the device. The
    // type is string.
    Username interface{}

    // Sessions associated with the users. The type is slice of
    // AaaData_AaaUsers_AaaSessions.
    AaaSessions []AaaData_AaaUsers_AaaSessions
}

func (aaaUsers *AaaData_AaaUsers) GetEntityData() *types.CommonEntityData {
    aaaUsers.EntityData.YFilter = aaaUsers.YFilter
    aaaUsers.EntityData.YangName = "aaa-users"
    aaaUsers.EntityData.BundleName = "cisco_ios_xe"
    aaaUsers.EntityData.ParentYangName = "aaa-data"
    aaaUsers.EntityData.SegmentPath = "aaa-users" + "[username='" + fmt.Sprintf("%v", aaaUsers.Username) + "']"
    aaaUsers.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaUsers.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaUsers.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaUsers.EntityData.Children = make(map[string]types.YChild)
    aaaUsers.EntityData.Children["aaa-sessions"] = types.YChild{"AaaSessions", nil}
    for i := range aaaUsers.AaaSessions {
        aaaUsers.EntityData.Children[types.GetSegmentPath(&aaaUsers.AaaSessions[i])] = types.YChild{"AaaSessions", &aaaUsers.AaaSessions[i]}
    }
    aaaUsers.EntityData.Leafs = make(map[string]types.YLeaf)
    aaaUsers.EntityData.Leafs["username"] = types.YLeaf{"Username", aaaUsers.Username}
    return &(aaaUsers.EntityData)
}

// AaaData_AaaUsers_AaaSessions
// Sessions associated with the users
type AaaData_AaaUsers_AaaSessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

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
    aaaSessions.EntityData.SegmentPath = "aaa-sessions" + "[aaa-uid='" + fmt.Sprintf("%v", aaaSessions.AaaUid) + "']"
    aaaSessions.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    aaaSessions.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    aaaSessions.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    aaaSessions.EntityData.Children = make(map[string]types.YChild)
    aaaSessions.EntityData.Leafs = make(map[string]types.YLeaf)
    aaaSessions.EntityData.Leafs["aaa-uid"] = types.YLeaf{"AaaUid", aaaSessions.AaaUid}
    aaaSessions.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", aaaSessions.SessionId}
    aaaSessions.EntityData.Leafs["ip-addr"] = types.YLeaf{"IpAddr", aaaSessions.IpAddr}
    aaaSessions.EntityData.Leafs["protocol"] = types.YLeaf{"Protocol", aaaSessions.Protocol}
    aaaSessions.EntityData.Leafs["login-time"] = types.YLeaf{"LoginTime", aaaSessions.LoginTime}
    return &(aaaSessions.EntityData)
}

