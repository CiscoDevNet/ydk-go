// This module contains a collection of YANG definitions
// for Cisco IOS-XR iedge4710 package operational data.
// 
// This module contains definitions
// for the following management objects:
//   subscriber: Subscriber operational data
//   iedge-license-manager: iedge license manager
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package iedge4710_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package iedge4710_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-oper subscriber}", reflect.TypeOf(Subscriber{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-oper:subscriber", reflect.TypeOf(Subscriber{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-iedge4710-oper iedge-license-manager}", reflect.TypeOf(IedgeLicenseManager{}))
    ydk.RegisterEntity("Cisco-IOS-XR-iedge4710-oper:iedge-license-manager", reflect.TypeOf(IedgeLicenseManager{}))
}

// SubscriberAuthorStateFilterFlag represents Subscriber author state filter flag
type SubscriberAuthorStateFilterFlag string

const (
    // UnAuthorized
    SubscriberAuthorStateFilterFlag_un_authorized SubscriberAuthorStateFilterFlag = "un-authorized"

    // Authorized
    SubscriberAuthorStateFilterFlag_authorized SubscriberAuthorStateFilterFlag = "authorized"
)

// SubscriberAddressFamilyFilterFlag represents Subscriber address family filter flag
type SubscriberAddressFamilyFilterFlag string

const (
    // IPv4 only
    SubscriberAddressFamilyFilterFlag_ipv4_only SubscriberAddressFamilyFilterFlag = "ipv4-only"

    // IPv6 only
    SubscriberAddressFamilyFilterFlag_ipv6_only SubscriberAddressFamilyFilterFlag = "ipv6-only"

    // IPv4 all
    SubscriberAddressFamilyFilterFlag_ipv4_all SubscriberAddressFamilyFilterFlag = "ipv4-all"

    // IPv6 all
    SubscriberAddressFamilyFilterFlag_ipv6_all SubscriberAddressFamilyFilterFlag = "ipv6-all"

    // Dual all
    SubscriberAddressFamilyFilterFlag_dual_all SubscriberAddressFamilyFilterFlag = "dual-all"

    // Dual part up
    SubscriberAddressFamilyFilterFlag_dual_part_up SubscriberAddressFamilyFilterFlag = "dual-part-up"

    // Dual up
    SubscriberAddressFamilyFilterFlag_dual_up SubscriberAddressFamilyFilterFlag = "dual-up"

    // LAC
    SubscriberAddressFamilyFilterFlag_lac SubscriberAddressFamilyFilterFlag = "lac"
)

// SubscriberStateFilterFlag represents Subscriber state filter flag
type SubscriberStateFilterFlag string

const (
    // Initializing
    SubscriberStateFilterFlag_initializing SubscriberStateFilterFlag = "initializing"

    // Connecting
    SubscriberStateFilterFlag_connecting SubscriberStateFilterFlag = "connecting"

    // Connected
    SubscriberStateFilterFlag_connected SubscriberStateFilterFlag = "connected"

    // Activated
    SubscriberStateFilterFlag_activated SubscriberStateFilterFlag = "activated"

    // Idle
    SubscriberStateFilterFlag_idle SubscriberStateFilterFlag = "idle"

    // Disconnecting
    SubscriberStateFilterFlag_disconnecting SubscriberStateFilterFlag = "disconnecting"

    // End
    SubscriberStateFilterFlag_end SubscriberStateFilterFlag = "end"
)

// SubscriberAuthenStateFilterFlag represents Subscriber authen state filter flag
type SubscriberAuthenStateFilterFlag string

const (
    // UnAuthenticated
    SubscriberAuthenStateFilterFlag_un_authenticated SubscriberAuthenStateFilterFlag = "un-authenticated"

    // Authenticated
    SubscriberAuthenStateFilterFlag_authenticated SubscriberAuthenStateFilterFlag = "authenticated"
)

// IedgeOperSessionAfState represents Subscriber session address family state
type IedgeOperSessionAfState string

const (
    // Not started
    IedgeOperSessionAfState_not_started IedgeOperSessionAfState = "not-started"

    // Down
    IedgeOperSessionAfState_down IedgeOperSessionAfState = "down"

    // Up Pending
    IedgeOperSessionAfState_up_pending IedgeOperSessionAfState = "up-pending"

    // Up
    IedgeOperSessionAfState_up IedgeOperSessionAfState = "up"
)

// AaaTerminateCause represents AAA terminate cause types
type AaaTerminateCause string

const (
    // None
    AaaTerminateCause_none AaaTerminateCause = "none"

    // User request
    AaaTerminateCause_user_request AaaTerminateCause = "user-request"

    // Lost carrier
    AaaTerminateCause_lost_carrier AaaTerminateCause = "lost-carrier"

    // Lost service
    AaaTerminateCause_lost_service AaaTerminateCause = "lost-service"

    // Idle timeout
    AaaTerminateCause_idle_timeout AaaTerminateCause = "idle-timeout"

    // Session timeout
    AaaTerminateCause_session_timeout AaaTerminateCause = "session-timeout"

    // Admin reset
    AaaTerminateCause_admin_reset AaaTerminateCause = "admin-reset"

    // Admin reboot
    AaaTerminateCause_admin_reboot AaaTerminateCause = "admin-reboot"

    // Port error
    AaaTerminateCause_port_error AaaTerminateCause = "port-error"

    // NAS error
    AaaTerminateCause_nas_error AaaTerminateCause = "nas-error"

    // NAS request
    AaaTerminateCause_nas_request AaaTerminateCause = "nas-request"

    // NAS reboot
    AaaTerminateCause_nas_reboot AaaTerminateCause = "nas-reboot"

    // Port unneeded
    AaaTerminateCause_port_unneeded AaaTerminateCause = "port-unneeded"

    // Port preempted
    AaaTerminateCause_port_preempted AaaTerminateCause = "port-preempted"

    // Port suspended
    AaaTerminateCause_port_suspended AaaTerminateCause = "port-suspended"

    // Service unavailable
    AaaTerminateCause_service_unavailable AaaTerminateCause = "service-unavailable"

    // Callback
    AaaTerminateCause_callback AaaTerminateCause = "callback"

    // User error
    AaaTerminateCause_user_error AaaTerminateCause = "user-error"

    // Host request
    AaaTerminateCause_host_request AaaTerminateCause = "host-request"

    // Supplicant restart
    AaaTerminateCause_supplicant_restart AaaTerminateCause = "supplicant-restart"

    // Reauthorization failure
    AaaTerminateCause_reauthorization_failure AaaTerminateCause = "reauthorization-failure"

    // Port reinitialized
    AaaTerminateCause_port_reinitialized AaaTerminateCause = "port-reinitialized"

    // Admin disabled
    AaaTerminateCause_admin_disabled AaaTerminateCause = "admin-disabled"
)

// AaaInterface represents AAA interface types
type AaaInterface string

const (
    // None
    AaaInterface_none AaaInterface = "none"

    // Primary rate
    AaaInterface_primary_rate AaaInterface = "primary-rate"

    // Basic rate
    AaaInterface_basic_rate AaaInterface = "basic-rate"

    // Serial
    AaaInterface_serial AaaInterface = "serial"

    // Asynchronous
    AaaInterface_asynchronous AaaInterface = "asynchronous"

    // VTY
    AaaInterface_vty AaaInterface = "vty"

    // ATM
    AaaInterface_atm AaaInterface = "atm"

    // Ethernet
    AaaInterface_ethernet AaaInterface = "ethernet"

    // PPP over ATM
    AaaInterface_ppp_over_atm AaaInterface = "ppp-over-atm"

    // PPPoE over ATM
    AaaInterface_pppoe_over_atm AaaInterface = "pppoe-over-atm"

    // PPPoE over ethernet
    AaaInterface_pppoe_over_ethernet AaaInterface = "pppoe-over-ethernet"

    // PPP over VLAN
    AaaInterface_ppp_over_vlan AaaInterface = "ppp-over-vlan"

    // PPP over Q in Q
    AaaInterface_ppp_over_qinq AaaInterface = "ppp-over-qinq"

    // V120
    AaaInterface_v120 AaaInterface = "v120"

    // V110
    AaaInterface_v110 AaaInterface = "v110"

    // PHS internet access forum standard
    AaaInterface_piafs AaaInterface = "piafs"

    // X75
    AaaInterface_x75 AaaInterface = "x75"

    // IP sec
    AaaInterface_ip_sec AaaInterface = "ip-sec"

    // Other
    AaaInterface_other AaaInterface = "other"

    // Virtual PPPoE over ethernet
    AaaInterface_virtual_pppoe_over_ethernet AaaInterface = "virtual-pppoe-over-ethernet"

    // Virtual PPPoE over VLAN
    AaaInterface_virtual_pppoe_over_vlan AaaInterface = "virtual-pppoe-over-vlan"

    // Virtual PPPoE over Q in Q
    AaaInterface_virtual_pppoe_over_qinq AaaInterface = "virtual-pppoe-over-qinq"

    // IPoE over ethernet
    AaaInterface_ipo_e_over_ethernet AaaInterface = "ipo-e-over-ethernet"

    // IPoE over VLAN
    AaaInterface_ipo_e_over_vlan AaaInterface = "ipo-e-over-vlan"

    // IPoE over Q in Q
    AaaInterface_ipo_e_over_qinq AaaInterface = "ipo-e-over-qinq"

    // Virtual IPoE over ethernet
    AaaInterface_virtual_i_po_e_over_ethernet AaaInterface = "virtual-i-po-e-over-ethernet"

    // Virtual IPoE over VLAN
    AaaInterface_virtual_i_po_e_over_vlan AaaInterface = "virtual-i-po-e-over-vlan"

    // Virtual IPoE over Q in Q
    AaaInterface_virtual_i_po_e_over_qinq AaaInterface = "virtual-i-po-e-over-qinq"
)

// AaaTunnelProto represents Tunnel protocol types
type AaaTunnelProto string

const (
    // None
    AaaTunnelProto_none AaaTunnelProto = "none"

    // Point-to-point tunneling protocol
    AaaTunnelProto_pptp AaaTunnelProto = "pptp"

    // Layer 2 forwarding
    AaaTunnelProto_l2f AaaTunnelProto = "l2f"

    // Layer 2 tunneling protocol
    AaaTunnelProto_l2tp AaaTunnelProto = "l2tp"

    // Ascend tunnel management protocol
    AaaTunnelProto_atmp AaaTunnelProto = "atmp"

    // VLAN trunk protocol
    AaaTunnelProto_vtp AaaTunnelProto = "vtp"

    // Authentication header
    AaaTunnelProto_ah AaaTunnelProto = "ah"

    // IP over IP
    AaaTunnelProto_ip_over_ip AaaTunnelProto = "ip-over-ip"

    // Minimum IP over IP
    AaaTunnelProto_minimum_ip_over_ip AaaTunnelProto = "minimum-ip-over-ip"

    // Encapsulating security payload
    AaaTunnelProto_esp AaaTunnelProto = "esp"

    // Generic routing encapsulation
    AaaTunnelProto_gre AaaTunnelProto = "gre"

    // Bay dial virtual services
    AaaTunnelProto_bay_dvs AaaTunnelProto = "bay-dvs"

    // IP in IP
    AaaTunnelProto_ip_in_ip AaaTunnelProto = "ip-in-ip"

    // VLAN
    AaaTunnelProto_vlan AaaTunnelProto = "vlan"
)

// AaaTunnelMedium represents Tunnel medium types
type AaaTunnelMedium string

const (
    // None
    AaaTunnelMedium_none AaaTunnelMedium = "none"

    // IPv4
    AaaTunnelMedium_ipv4 AaaTunnelMedium = "ipv4"

    // IPv6
    AaaTunnelMedium_ipv6 AaaTunnelMedium = "ipv6"

    // NSAP
    AaaTunnelMedium_nsap AaaTunnelMedium = "nsap"

    // HDLC
    AaaTunnelMedium_hdlc AaaTunnelMedium = "hdlc"

    // BBN
    AaaTunnelMedium_bbn AaaTunnelMedium = "bbn"

    // ALL 802
    AaaTunnelMedium_all802 AaaTunnelMedium = "all802"
)

// AaaAuthService represents AAA authorization service types
type AaaAuthService string

const (
    // None
    AaaAuthService_none AaaAuthService = "none"

    // Login
    AaaAuthService_login AaaAuthService = "login"

    // Framed
    AaaAuthService_framed AaaAuthService = "framed"

    // Callback login
    AaaAuthService_callback_login AaaAuthService = "callback-login"

    // Callback framed
    AaaAuthService_callback_framed AaaAuthService = "callback-framed"

    // Outbound
    AaaAuthService_outbound AaaAuthService = "outbound"

    // Administrator
    AaaAuthService_administrator AaaAuthService = "administrator"

    // Prompt
    AaaAuthService_prompt AaaAuthService = "prompt"

    // Authentication only
    AaaAuthService_authentication_only AaaAuthService = "authentication-only"

    // Callback NAS prompt
    AaaAuthService_callback_nas_prompt AaaAuthService = "callback-nas-prompt"

    // Call check
    AaaAuthService_call_check AaaAuthService = "call-check"

    // Callback administrator
    AaaAuthService_callback_administrator AaaAuthService = "callback-administrator"

    // Voice
    AaaAuthService_voice AaaAuthService = "voice"

    // Fax
    AaaAuthService_fax AaaAuthService = "fax"

    // Modem relay
    AaaAuthService_modem_relay AaaAuthService = "modem-relay"

    // EAP over UDP
    AaaAuthService_eap_over_udp AaaAuthService = "eap-over-udp"
)

// IedgeOperSessionState represents Subscriber session states
type IedgeOperSessionState string

const (
    // Initialize
    IedgeOperSessionState_initialize IedgeOperSessionState = "initialize"

    // Connecting
    IedgeOperSessionState_connecting IedgeOperSessionState = "connecting"

    // Connected
    IedgeOperSessionState_connected IedgeOperSessionState = "connected"

    // Activated
    IedgeOperSessionState_activated IedgeOperSessionState = "activated"

    // Idle
    IedgeOperSessionState_idle IedgeOperSessionState = "idle"

    // Disconnecting
    IedgeOperSessionState_disconnecting IedgeOperSessionState = "disconnecting"

    // End
    IedgeOperSessionState_end IedgeOperSessionState = "end"
)

// IedgePppSub represents PPPoE sub types
type IedgePppSub string

const (
    // PPP termination and aggregation
    IedgePppSub_pta IedgePppSub = "pta"

    // L2TP access controller
    IedgePppSub_lac IedgePppSub = "lac"
)

// IedgeOperSession represents Subscriber session types
type IedgeOperSession string

const (
    // Unknown
    IedgeOperSession_unknown IedgeOperSession = "unknown"

    // PPPoE/PPP client
    IedgeOperSession_pppoe IedgeOperSession = "pppoe"

    // PPP serial client
    IedgeOperSession_ppp IedgeOperSession = "ppp"

    // IP subscriber - packet trigger
    IedgeOperSession_ip_packet_trigger IedgeOperSession = "ip-packet-trigger"

    // IP subscriber - DHCP trigger
    IedgeOperSession_ip_packet_dhcp_trigger IedgeOperSession = "ip-packet-dhcp-trigger"
)

// Subscriber
// Subscriber operational data
type Subscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager operational data.
    Manager Subscriber_Manager

    // Subscriber session operational data.
    Session Subscriber_Session
}

func (subscriber *Subscriber) GetEntityData() *types.CommonEntityData {
    subscriber.EntityData.YFilter = subscriber.YFilter
    subscriber.EntityData.YangName = "subscriber"
    subscriber.EntityData.BundleName = "cisco_ios_xr"
    subscriber.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-oper"
    subscriber.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-oper:subscriber"
    subscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriber.EntityData.Children = make(map[string]types.YChild)
    subscriber.EntityData.Children["manager"] = types.YChild{"Manager", &subscriber.Manager}
    subscriber.EntityData.Children["session"] = types.YChild{"Session", &subscriber.Session}
    subscriber.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(subscriber.EntityData)
}

// Subscriber_Manager
// Subscriber manager operational data
type Subscriber_Manager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager list of nodes.
    Nodes Subscriber_Manager_Nodes
}

func (manager *Subscriber_Manager) GetEntityData() *types.CommonEntityData {
    manager.EntityData.YFilter = manager.YFilter
    manager.EntityData.YangName = "manager"
    manager.EntityData.BundleName = "cisco_ios_xr"
    manager.EntityData.ParentYangName = "subscriber"
    manager.EntityData.SegmentPath = "manager"
    manager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    manager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    manager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    manager.EntityData.Children = make(map[string]types.YChild)
    manager.EntityData.Children["nodes"] = types.YChild{"Nodes", &manager.Nodes}
    manager.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(manager.EntityData)
}

// Subscriber_Manager_Nodes
// Subscriber manager list of nodes
type Subscriber_Manager_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber manager operational data for a particular node. The type is
    // slice of Subscriber_Manager_Nodes_Node.
    Node []Subscriber_Manager_Nodes_Node
}

func (nodes *Subscriber_Manager_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "manager"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Subscriber_Manager_Nodes_Node
// Subscriber manager operational data for a
// particular node
type Subscriber_Manager_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Subscriber manager statistics.
    Statistics Subscriber_Manager_Nodes_Node_Statistics
}

func (node *Subscriber_Manager_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["statistics"] = types.YChild{"Statistics", &node.Statistics}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics
// Subscriber manager statistics
type Subscriber_Manager_Nodes_Node_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AAA statistics.
    Aaa Subscriber_Manager_Nodes_Node_Statistics_Aaa

    // Aggregate summary of statistics.
    AggregateSummary Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary

    // Geo Redundancy statistics.
    Srg Subscriber_Manager_Nodes_Node_Statistics_Srg
}

func (statistics *Subscriber_Manager_Nodes_Node_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "node"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = make(map[string]types.YChild)
    statistics.EntityData.Children["aaa"] = types.YChild{"Aaa", &statistics.Aaa}
    statistics.EntityData.Children["aggregate-summary"] = types.YChild{"AggregateSummary", &statistics.AggregateSummary}
    statistics.EntityData.Children["srg"] = types.YChild{"Srg", &statistics.Srg}
    statistics.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(statistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa
// AAA statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Aggregate accounting statistics.
    AggregateAccounting Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting

    // Authentication statistics.
    Authentication Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication

    // Aggregate mobility statistics.
    AggregateMobility Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility

    // Aggregate authentication statistics.
    AggregateAuthentication Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication

    // Display all subscriber management statistics.
    AccountingStatsAll Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll

    // Change of authorization (COA) statistics.
    ChangeOfAuthorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization

    // Authorization statistics.
    Authorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization

    // Aggregate authorization statistics.
    AggregateAuthorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization

    // Display all subscriber management total statistics.
    AggregateAccountingStatsAll Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll

    // Accounting statistics.
    Accounting Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting

    // Mobility statistics.
    Mobility Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility

    // Aggregate change of authorization (COA) statistics.
    AggregateChangeOfAuthorization Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization
}

func (aaa *Subscriber_Manager_Nodes_Node_Statistics_Aaa) GetEntityData() *types.CommonEntityData {
    aaa.EntityData.YFilter = aaa.YFilter
    aaa.EntityData.YangName = "aaa"
    aaa.EntityData.BundleName = "cisco_ios_xr"
    aaa.EntityData.ParentYangName = "statistics"
    aaa.EntityData.SegmentPath = "aaa"
    aaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaa.EntityData.Children = make(map[string]types.YChild)
    aaa.EntityData.Children["aggregate-accounting"] = types.YChild{"AggregateAccounting", &aaa.AggregateAccounting}
    aaa.EntityData.Children["authentication"] = types.YChild{"Authentication", &aaa.Authentication}
    aaa.EntityData.Children["aggregate-mobility"] = types.YChild{"AggregateMobility", &aaa.AggregateMobility}
    aaa.EntityData.Children["aggregate-authentication"] = types.YChild{"AggregateAuthentication", &aaa.AggregateAuthentication}
    aaa.EntityData.Children["accounting-stats-all"] = types.YChild{"AccountingStatsAll", &aaa.AccountingStatsAll}
    aaa.EntityData.Children["change-of-authorization"] = types.YChild{"ChangeOfAuthorization", &aaa.ChangeOfAuthorization}
    aaa.EntityData.Children["authorization"] = types.YChild{"Authorization", &aaa.Authorization}
    aaa.EntityData.Children["aggregate-authorization"] = types.YChild{"AggregateAuthorization", &aaa.AggregateAuthorization}
    aaa.EntityData.Children["aggregate-accounting-stats-all"] = types.YChild{"AggregateAccountingStatsAll", &aaa.AggregateAccountingStatsAll}
    aaa.EntityData.Children["accounting"] = types.YChild{"Accounting", &aaa.Accounting}
    aaa.EntityData.Children["mobility"] = types.YChild{"Mobility", &aaa.Mobility}
    aaa.EntityData.Children["aggregate-change-of-authorization"] = types.YChild{"AggregateChangeOfAuthorization", &aaa.AggregateChangeOfAuthorization}
    aaa.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aaa.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting
// Aggregate accounting statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight
}

func (aggregateAccounting *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting) GetEntityData() *types.CommonEntityData {
    aggregateAccounting.EntityData.YFilter = aggregateAccounting.YFilter
    aggregateAccounting.EntityData.YangName = "aggregate-accounting"
    aggregateAccounting.EntityData.BundleName = "cisco_ios_xr"
    aggregateAccounting.EntityData.ParentYangName = "aaa"
    aggregateAccounting.EntityData.SegmentPath = "aggregate-accounting"
    aggregateAccounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAccounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAccounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAccounting.EntityData.Children = make(map[string]types.YChild)
    aggregateAccounting.EntityData.Children["start"] = types.YChild{"Start", &aggregateAccounting.Start}
    aggregateAccounting.EntityData.Children["stop"] = types.YChild{"Stop", &aggregateAccounting.Stop}
    aggregateAccounting.EntityData.Children["interim"] = types.YChild{"Interim", &aggregateAccounting.Interim}
    aggregateAccounting.EntityData.Children["pass-through"] = types.YChild{"PassThrough", &aggregateAccounting.PassThrough}
    aggregateAccounting.EntityData.Children["update"] = types.YChild{"Update", &aggregateAccounting.Update}
    aggregateAccounting.EntityData.Children["interim-inflight"] = types.YChild{"InterimInflight", &aggregateAccounting.InterimInflight}
    aggregateAccounting.EntityData.Leafs = make(map[string]types.YLeaf)
    aggregateAccounting.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", aggregateAccounting.ActiveSessions}
    aggregateAccounting.EntityData.Leafs["started-sessions"] = types.YLeaf{"StartedSessions", aggregateAccounting.StartedSessions}
    aggregateAccounting.EntityData.Leafs["stopped-sessions"] = types.YLeaf{"StoppedSessions", aggregateAccounting.StoppedSessions}
    aggregateAccounting.EntityData.Leafs["policy-plane-errored-requests"] = types.YLeaf{"PolicyPlaneErroredRequests", aggregateAccounting.PolicyPlaneErroredRequests}
    aggregateAccounting.EntityData.Leafs["policy-plane-unknown-requests"] = types.YLeaf{"PolicyPlaneUnknownRequests", aggregateAccounting.PolicyPlaneUnknownRequests}
    return &(aggregateAccounting.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "aggregate-accounting"
    start.EntityData.SegmentPath = "start"
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = make(map[string]types.YChild)
    start.EntityData.Leafs = make(map[string]types.YLeaf)
    start.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", start.ReceivedRequests}
    start.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", start.ErroredRequests}
    start.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests}
    start.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", start.AaaSentRequests}
    start.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses}
    start.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses}
    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "aggregate-accounting"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    stop.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", stop.ReceivedRequests}
    stop.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", stop.ErroredRequests}
    stop.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests}
    stop.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", stop.AaaSentRequests}
    stop.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses}
    stop.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses}
    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "aggregate-accounting"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = make(map[string]types.YChild)
    interim.EntityData.Leafs = make(map[string]types.YLeaf)
    interim.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", interim.ReceivedRequests}
    interim.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", interim.ErroredRequests}
    interim.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests}
    interim.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", interim.AaaSentRequests}
    interim.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses}
    interim.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses}
    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "aggregate-accounting"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = make(map[string]types.YChild)
    passThrough.EntityData.Leafs = make(map[string]types.YLeaf)
    passThrough.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests}
    passThrough.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", passThrough.ErroredRequests}
    passThrough.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests}
    passThrough.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests}
    passThrough.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses}
    passThrough.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses}
    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "aggregate-accounting"
    update.EntityData.SegmentPath = "update"
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = make(map[string]types.YChild)
    update.EntityData.Leafs = make(map[string]types.YLeaf)
    update.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", update.ReceivedRequests}
    update.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", update.ErroredRequests}
    update.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests}
    update.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", update.AaaSentRequests}
    update.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses}
    update.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses}
    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccounting_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "aggregate-accounting"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = make(map[string]types.YChild)
    interimInflight.EntityData.Leafs = make(map[string]types.YLeaf)
    interimInflight.EntityData.Leafs["quota-exhausts"] = types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts}
    interimInflight.EntityData.Leafs["denied-requests"] = types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests}
    interimInflight.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests}
    interimInflight.EntityData.Leafs["total-quota-of-requests"] = types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests}
    interimInflight.EntityData.Leafs["remaining-quota-of-requests"] = types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests}
    interimInflight.EntityData.Leafs["low-water-mark-quota-of-requests"] = types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests}
    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication
// Authentication statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authentication *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "aaa"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = make(map[string]types.YChild)
    authentication.EntityData.Leafs = make(map[string]types.YLeaf)
    authentication.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", authentication.SentRequests}
    authentication.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", authentication.AcceptedRequests}
    authentication.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", authentication.SuccessfulRequests}
    authentication.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", authentication.RejectedRequests}
    authentication.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", authentication.UnreachableRequests}
    authentication.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", authentication.ErroredRequests}
    authentication.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", authentication.IncompleteRequests}
    authentication.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", authentication.TerminatedRequests}
    return &(authentication.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility
// Aggregate mobility statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (aggregateMobility *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateMobility) GetEntityData() *types.CommonEntityData {
    aggregateMobility.EntityData.YFilter = aggregateMobility.YFilter
    aggregateMobility.EntityData.YangName = "aggregate-mobility"
    aggregateMobility.EntityData.BundleName = "cisco_ios_xr"
    aggregateMobility.EntityData.ParentYangName = "aaa"
    aggregateMobility.EntityData.SegmentPath = "aggregate-mobility"
    aggregateMobility.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateMobility.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateMobility.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateMobility.EntityData.Children = make(map[string]types.YChild)
    aggregateMobility.EntityData.Leafs = make(map[string]types.YLeaf)
    aggregateMobility.EntityData.Leafs["send-request-successes"] = types.YLeaf{"SendRequestSuccesses", aggregateMobility.SendRequestSuccesses}
    aggregateMobility.EntityData.Leafs["send-request-failures"] = types.YLeaf{"SendRequestFailures", aggregateMobility.SendRequestFailures}
    aggregateMobility.EntityData.Leafs["receive-response-successes"] = types.YLeaf{"ReceiveResponseSuccesses", aggregateMobility.ReceiveResponseSuccesses}
    aggregateMobility.EntityData.Leafs["receive-response-failures"] = types.YLeaf{"ReceiveResponseFailures", aggregateMobility.ReceiveResponseFailures}
    return &(aggregateMobility.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication
// Aggregate authentication statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (aggregateAuthentication *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthentication) GetEntityData() *types.CommonEntityData {
    aggregateAuthentication.EntityData.YFilter = aggregateAuthentication.YFilter
    aggregateAuthentication.EntityData.YangName = "aggregate-authentication"
    aggregateAuthentication.EntityData.BundleName = "cisco_ios_xr"
    aggregateAuthentication.EntityData.ParentYangName = "aaa"
    aggregateAuthentication.EntityData.SegmentPath = "aggregate-authentication"
    aggregateAuthentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAuthentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAuthentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAuthentication.EntityData.Children = make(map[string]types.YChild)
    aggregateAuthentication.EntityData.Leafs = make(map[string]types.YLeaf)
    aggregateAuthentication.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", aggregateAuthentication.SentRequests}
    aggregateAuthentication.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", aggregateAuthentication.AcceptedRequests}
    aggregateAuthentication.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", aggregateAuthentication.SuccessfulRequests}
    aggregateAuthentication.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", aggregateAuthentication.RejectedRequests}
    aggregateAuthentication.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", aggregateAuthentication.UnreachableRequests}
    aggregateAuthentication.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", aggregateAuthentication.ErroredRequests}
    aggregateAuthentication.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", aggregateAuthentication.IncompleteRequests}
    aggregateAuthentication.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", aggregateAuthentication.TerminatedRequests}
    return &(aggregateAuthentication.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll
// Display all subscriber management
// statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of stats for accounting.
    AccountingStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics

    // List of stats for authentication.
    AuthenticationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics

    // List of stats for authorization.
    AuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics

    // List of stats for COA.
    ChangeOfAuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics

    // List of stats for Mobility.
    MobilityStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics
}

func (accountingStatsAll *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll) GetEntityData() *types.CommonEntityData {
    accountingStatsAll.EntityData.YFilter = accountingStatsAll.YFilter
    accountingStatsAll.EntityData.YangName = "accounting-stats-all"
    accountingStatsAll.EntityData.BundleName = "cisco_ios_xr"
    accountingStatsAll.EntityData.ParentYangName = "aaa"
    accountingStatsAll.EntityData.SegmentPath = "accounting-stats-all"
    accountingStatsAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatsAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatsAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatsAll.EntityData.Children = make(map[string]types.YChild)
    accountingStatsAll.EntityData.Children["accounting-statistics"] = types.YChild{"AccountingStatistics", &accountingStatsAll.AccountingStatistics}
    accountingStatsAll.EntityData.Children["authentication-statistics"] = types.YChild{"AuthenticationStatistics", &accountingStatsAll.AuthenticationStatistics}
    accountingStatsAll.EntityData.Children["authorization-statistics"] = types.YChild{"AuthorizationStatistics", &accountingStatsAll.AuthorizationStatistics}
    accountingStatsAll.EntityData.Children["change-of-authorization-statistics"] = types.YChild{"ChangeOfAuthorizationStatistics", &accountingStatsAll.ChangeOfAuthorizationStatistics}
    accountingStatsAll.EntityData.Children["mobility-statistics"] = types.YChild{"MobilityStatistics", &accountingStatsAll.MobilityStatistics}
    accountingStatsAll.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accountingStatsAll.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics
// List of stats for accounting
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight
}

func (accountingStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics) GetEntityData() *types.CommonEntityData {
    accountingStatistics.EntityData.YFilter = accountingStatistics.YFilter
    accountingStatistics.EntityData.YangName = "accounting-statistics"
    accountingStatistics.EntityData.BundleName = "cisco_ios_xr"
    accountingStatistics.EntityData.ParentYangName = "accounting-stats-all"
    accountingStatistics.EntityData.SegmentPath = "accounting-statistics"
    accountingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatistics.EntityData.Children = make(map[string]types.YChild)
    accountingStatistics.EntityData.Children["start"] = types.YChild{"Start", &accountingStatistics.Start}
    accountingStatistics.EntityData.Children["stop"] = types.YChild{"Stop", &accountingStatistics.Stop}
    accountingStatistics.EntityData.Children["interim"] = types.YChild{"Interim", &accountingStatistics.Interim}
    accountingStatistics.EntityData.Children["pass-through"] = types.YChild{"PassThrough", &accountingStatistics.PassThrough}
    accountingStatistics.EntityData.Children["update"] = types.YChild{"Update", &accountingStatistics.Update}
    accountingStatistics.EntityData.Children["interim-inflight"] = types.YChild{"InterimInflight", &accountingStatistics.InterimInflight}
    accountingStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    accountingStatistics.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", accountingStatistics.ActiveSessions}
    accountingStatistics.EntityData.Leafs["started-sessions"] = types.YLeaf{"StartedSessions", accountingStatistics.StartedSessions}
    accountingStatistics.EntityData.Leafs["stopped-sessions"] = types.YLeaf{"StoppedSessions", accountingStatistics.StoppedSessions}
    accountingStatistics.EntityData.Leafs["policy-plane-errored-requests"] = types.YLeaf{"PolicyPlaneErroredRequests", accountingStatistics.PolicyPlaneErroredRequests}
    accountingStatistics.EntityData.Leafs["policy-plane-unknown-requests"] = types.YLeaf{"PolicyPlaneUnknownRequests", accountingStatistics.PolicyPlaneUnknownRequests}
    return &(accountingStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "accounting-statistics"
    start.EntityData.SegmentPath = "start"
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = make(map[string]types.YChild)
    start.EntityData.Leafs = make(map[string]types.YLeaf)
    start.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", start.ReceivedRequests}
    start.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", start.ErroredRequests}
    start.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests}
    start.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", start.AaaSentRequests}
    start.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses}
    start.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses}
    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "accounting-statistics"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    stop.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", stop.ReceivedRequests}
    stop.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", stop.ErroredRequests}
    stop.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests}
    stop.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", stop.AaaSentRequests}
    stop.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses}
    stop.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses}
    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting-statistics"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = make(map[string]types.YChild)
    interim.EntityData.Leafs = make(map[string]types.YLeaf)
    interim.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", interim.ReceivedRequests}
    interim.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", interim.ErroredRequests}
    interim.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests}
    interim.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", interim.AaaSentRequests}
    interim.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses}
    interim.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses}
    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "accounting-statistics"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = make(map[string]types.YChild)
    passThrough.EntityData.Leafs = make(map[string]types.YLeaf)
    passThrough.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests}
    passThrough.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", passThrough.ErroredRequests}
    passThrough.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests}
    passThrough.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests}
    passThrough.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses}
    passThrough.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses}
    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "accounting-statistics"
    update.EntityData.SegmentPath = "update"
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = make(map[string]types.YChild)
    update.EntityData.Leafs = make(map[string]types.YLeaf)
    update.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", update.ReceivedRequests}
    update.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", update.ErroredRequests}
    update.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests}
    update.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", update.AaaSentRequests}
    update.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses}
    update.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses}
    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AccountingStatistics_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "accounting-statistics"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = make(map[string]types.YChild)
    interimInflight.EntityData.Leafs = make(map[string]types.YLeaf)
    interimInflight.EntityData.Leafs["quota-exhausts"] = types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts}
    interimInflight.EntityData.Leafs["denied-requests"] = types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests}
    interimInflight.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests}
    interimInflight.EntityData.Leafs["total-quota-of-requests"] = types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests}
    interimInflight.EntityData.Leafs["remaining-quota-of-requests"] = types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests}
    interimInflight.EntityData.Leafs["low-water-mark-quota-of-requests"] = types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests}
    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics
// List of stats for authentication
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authenticationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthenticationStatistics) GetEntityData() *types.CommonEntityData {
    authenticationStatistics.EntityData.YFilter = authenticationStatistics.YFilter
    authenticationStatistics.EntityData.YangName = "authentication-statistics"
    authenticationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authenticationStatistics.EntityData.ParentYangName = "accounting-stats-all"
    authenticationStatistics.EntityData.SegmentPath = "authentication-statistics"
    authenticationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationStatistics.EntityData.Children = make(map[string]types.YChild)
    authenticationStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    authenticationStatistics.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", authenticationStatistics.SentRequests}
    authenticationStatistics.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", authenticationStatistics.AcceptedRequests}
    authenticationStatistics.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", authenticationStatistics.SuccessfulRequests}
    authenticationStatistics.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", authenticationStatistics.RejectedRequests}
    authenticationStatistics.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", authenticationStatistics.UnreachableRequests}
    authenticationStatistics.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", authenticationStatistics.ErroredRequests}
    authenticationStatistics.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", authenticationStatistics.IncompleteRequests}
    authenticationStatistics.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", authenticationStatistics.TerminatedRequests}
    return &(authenticationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics
// List of stats for authorization
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_AuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    authorizationStatistics.EntityData.YFilter = authorizationStatistics.YFilter
    authorizationStatistics.EntityData.YangName = "authorization-statistics"
    authorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authorizationStatistics.EntityData.ParentYangName = "accounting-stats-all"
    authorizationStatistics.EntityData.SegmentPath = "authorization-statistics"
    authorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizationStatistics.EntityData.Children = make(map[string]types.YChild)
    authorizationStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    authorizationStatistics.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", authorizationStatistics.SentRequests}
    authorizationStatistics.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", authorizationStatistics.AcceptedRequests}
    authorizationStatistics.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", authorizationStatistics.SuccessfulRequests}
    authorizationStatistics.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", authorizationStatistics.RejectedRequests}
    authorizationStatistics.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", authorizationStatistics.UnreachableRequests}
    authorizationStatistics.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", authorizationStatistics.ErroredRequests}
    authorizationStatistics.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", authorizationStatistics.IncompleteRequests}
    authorizationStatistics.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", authorizationStatistics.TerminatedRequests}
    return &(authorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics
// List of stats for COA
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
}

func (changeOfAuthorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    changeOfAuthorizationStatistics.EntityData.YFilter = changeOfAuthorizationStatistics.YFilter
    changeOfAuthorizationStatistics.EntityData.YangName = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    changeOfAuthorizationStatistics.EntityData.ParentYangName = "accounting-stats-all"
    changeOfAuthorizationStatistics.EntityData.SegmentPath = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeOfAuthorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeOfAuthorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeOfAuthorizationStatistics.EntityData.Children = make(map[string]types.YChild)
    changeOfAuthorizationStatistics.EntityData.Children["account-logon"] = types.YChild{"AccountLogon", &changeOfAuthorizationStatistics.AccountLogon}
    changeOfAuthorizationStatistics.EntityData.Children["account-logoff"] = types.YChild{"AccountLogoff", &changeOfAuthorizationStatistics.AccountLogoff}
    changeOfAuthorizationStatistics.EntityData.Children["account-update"] = types.YChild{"AccountUpdate", &changeOfAuthorizationStatistics.AccountUpdate}
    changeOfAuthorizationStatistics.EntityData.Children["session-disconnect"] = types.YChild{"SessionDisconnect", &changeOfAuthorizationStatistics.SessionDisconnect}
    changeOfAuthorizationStatistics.EntityData.Children["single-service-logon"] = types.YChild{"SingleServiceLogon", &changeOfAuthorizationStatistics.SingleServiceLogon}
    changeOfAuthorizationStatistics.EntityData.Children["single-service-logoff"] = types.YChild{"SingleServiceLogoff", &changeOfAuthorizationStatistics.SingleServiceLogoff}
    changeOfAuthorizationStatistics.EntityData.Children["single-service-modify"] = types.YChild{"SingleServiceModify", &changeOfAuthorizationStatistics.SingleServiceModify}
    changeOfAuthorizationStatistics.EntityData.Children["service-multi"] = types.YChild{"ServiceMulti", &changeOfAuthorizationStatistics.ServiceMulti}
    changeOfAuthorizationStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    changeOfAuthorizationStatistics.EntityData.Leafs["unknown-account-cmd-resps"] = types.YLeaf{"UnknownAccountCmdResps", changeOfAuthorizationStatistics.UnknownAccountCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["unknown-service-cmd-resps"] = types.YLeaf{"UnknownServiceCmdResps", changeOfAuthorizationStatistics.UnknownServiceCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["unknown-cmd-resps"] = types.YLeaf{"UnknownCmdResps", changeOfAuthorizationStatistics.UnknownCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["attr-list-retrieve-failure-resps"] = types.YLeaf{"AttrListRetrieveFailureResps", changeOfAuthorizationStatistics.AttrListRetrieveFailureResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["resp-send-failure"] = types.YLeaf{"RespSendFailure", changeOfAuthorizationStatistics.RespSendFailure}
    changeOfAuthorizationStatistics.EntityData.Leafs["internal-err-resps"] = types.YLeaf{"InternalErrResps", changeOfAuthorizationStatistics.InternalErrResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["service-profile-push-failure-resps"] = types.YLeaf{"ServiceProfilePushFailureResps", changeOfAuthorizationStatistics.ServiceProfilePushFailureResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["no-cmd-resps"] = types.YLeaf{"NoCmdResps", changeOfAuthorizationStatistics.NoCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["no-session-found-resps"] = types.YLeaf{"NoSessionFoundResps", changeOfAuthorizationStatistics.NoSessionFoundResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["no-session-peer-resps"] = types.YLeaf{"NoSessionPeerResps", changeOfAuthorizationStatistics.NoSessionPeerResps}
    return &(changeOfAuthorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = make(map[string]types.YChild)
    accountLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests}
    accountLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests}
    accountLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests}
    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = make(map[string]types.YChild)
    accountLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests}
    accountLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests}
    accountLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests}
    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = make(map[string]types.YChild)
    accountUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    accountUpdate.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests}
    accountUpdate.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests}
    accountUpdate.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests}
    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "change-of-authorization-statistics"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = make(map[string]types.YChild)
    sessionDisconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDisconnect.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests}
    sessionDisconnect.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests}
    sessionDisconnect.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests}
    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests}
    singleServiceLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests}
    singleServiceLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests}
    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests}
    singleServiceLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests}
    singleServiceLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests}
    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = make(map[string]types.YChild)
    singleServiceModify.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceModify.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests}
    singleServiceModify.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests}
    singleServiceModify.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests}
    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "change-of-authorization-statistics"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = make(map[string]types.YChild)
    serviceMulti.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceMulti.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests}
    serviceMulti.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests}
    serviceMulti.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests}
    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics
// List of stats for Mobility
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (mobilityStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AccountingStatsAll_MobilityStatistics) GetEntityData() *types.CommonEntityData {
    mobilityStatistics.EntityData.YFilter = mobilityStatistics.YFilter
    mobilityStatistics.EntityData.YangName = "mobility-statistics"
    mobilityStatistics.EntityData.BundleName = "cisco_ios_xr"
    mobilityStatistics.EntityData.ParentYangName = "accounting-stats-all"
    mobilityStatistics.EntityData.SegmentPath = "mobility-statistics"
    mobilityStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobilityStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobilityStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobilityStatistics.EntityData.Children = make(map[string]types.YChild)
    mobilityStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    mobilityStatistics.EntityData.Leafs["send-request-successes"] = types.YLeaf{"SendRequestSuccesses", mobilityStatistics.SendRequestSuccesses}
    mobilityStatistics.EntityData.Leafs["send-request-failures"] = types.YLeaf{"SendRequestFailures", mobilityStatistics.SendRequestFailures}
    mobilityStatistics.EntityData.Leafs["receive-response-successes"] = types.YLeaf{"ReceiveResponseSuccesses", mobilityStatistics.ReceiveResponseSuccesses}
    mobilityStatistics.EntityData.Leafs["receive-response-failures"] = types.YLeaf{"ReceiveResponseFailures", mobilityStatistics.ReceiveResponseFailures}
    return &(mobilityStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization
// Change of authorization (COA) statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti
}

func (changeOfAuthorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    changeOfAuthorization.EntityData.YFilter = changeOfAuthorization.YFilter
    changeOfAuthorization.EntityData.YangName = "change-of-authorization"
    changeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    changeOfAuthorization.EntityData.ParentYangName = "aaa"
    changeOfAuthorization.EntityData.SegmentPath = "change-of-authorization"
    changeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeOfAuthorization.EntityData.Children = make(map[string]types.YChild)
    changeOfAuthorization.EntityData.Children["account-logon"] = types.YChild{"AccountLogon", &changeOfAuthorization.AccountLogon}
    changeOfAuthorization.EntityData.Children["account-logoff"] = types.YChild{"AccountLogoff", &changeOfAuthorization.AccountLogoff}
    changeOfAuthorization.EntityData.Children["account-update"] = types.YChild{"AccountUpdate", &changeOfAuthorization.AccountUpdate}
    changeOfAuthorization.EntityData.Children["session-disconnect"] = types.YChild{"SessionDisconnect", &changeOfAuthorization.SessionDisconnect}
    changeOfAuthorization.EntityData.Children["single-service-logon"] = types.YChild{"SingleServiceLogon", &changeOfAuthorization.SingleServiceLogon}
    changeOfAuthorization.EntityData.Children["single-service-logoff"] = types.YChild{"SingleServiceLogoff", &changeOfAuthorization.SingleServiceLogoff}
    changeOfAuthorization.EntityData.Children["single-service-modify"] = types.YChild{"SingleServiceModify", &changeOfAuthorization.SingleServiceModify}
    changeOfAuthorization.EntityData.Children["service-multi"] = types.YChild{"ServiceMulti", &changeOfAuthorization.ServiceMulti}
    changeOfAuthorization.EntityData.Leafs = make(map[string]types.YLeaf)
    changeOfAuthorization.EntityData.Leafs["unknown-account-cmd-resps"] = types.YLeaf{"UnknownAccountCmdResps", changeOfAuthorization.UnknownAccountCmdResps}
    changeOfAuthorization.EntityData.Leafs["unknown-service-cmd-resps"] = types.YLeaf{"UnknownServiceCmdResps", changeOfAuthorization.UnknownServiceCmdResps}
    changeOfAuthorization.EntityData.Leafs["unknown-cmd-resps"] = types.YLeaf{"UnknownCmdResps", changeOfAuthorization.UnknownCmdResps}
    changeOfAuthorization.EntityData.Leafs["attr-list-retrieve-failure-resps"] = types.YLeaf{"AttrListRetrieveFailureResps", changeOfAuthorization.AttrListRetrieveFailureResps}
    changeOfAuthorization.EntityData.Leafs["resp-send-failure"] = types.YLeaf{"RespSendFailure", changeOfAuthorization.RespSendFailure}
    changeOfAuthorization.EntityData.Leafs["internal-err-resps"] = types.YLeaf{"InternalErrResps", changeOfAuthorization.InternalErrResps}
    changeOfAuthorization.EntityData.Leafs["service-profile-push-failure-resps"] = types.YLeaf{"ServiceProfilePushFailureResps", changeOfAuthorization.ServiceProfilePushFailureResps}
    changeOfAuthorization.EntityData.Leafs["no-cmd-resps"] = types.YLeaf{"NoCmdResps", changeOfAuthorization.NoCmdResps}
    changeOfAuthorization.EntityData.Leafs["no-session-found-resps"] = types.YLeaf{"NoSessionFoundResps", changeOfAuthorization.NoSessionFoundResps}
    changeOfAuthorization.EntityData.Leafs["no-session-peer-resps"] = types.YLeaf{"NoSessionPeerResps", changeOfAuthorization.NoSessionPeerResps}
    return &(changeOfAuthorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "change-of-authorization"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = make(map[string]types.YChild)
    accountLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests}
    accountLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests}
    accountLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests}
    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "change-of-authorization"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = make(map[string]types.YChild)
    accountLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests}
    accountLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests}
    accountLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests}
    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "change-of-authorization"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = make(map[string]types.YChild)
    accountUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    accountUpdate.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests}
    accountUpdate.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests}
    accountUpdate.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests}
    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "change-of-authorization"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = make(map[string]types.YChild)
    sessionDisconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDisconnect.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests}
    sessionDisconnect.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests}
    sessionDisconnect.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests}
    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "change-of-authorization"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests}
    singleServiceLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests}
    singleServiceLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests}
    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "change-of-authorization"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests}
    singleServiceLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests}
    singleServiceLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests}
    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "change-of-authorization"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = make(map[string]types.YChild)
    singleServiceModify.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceModify.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests}
    singleServiceModify.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests}
    singleServiceModify.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests}
    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_ChangeOfAuthorization_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "change-of-authorization"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = make(map[string]types.YChild)
    serviceMulti.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceMulti.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests}
    serviceMulti.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests}
    serviceMulti.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests}
    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization
// Authorization statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Authorization) GetEntityData() *types.CommonEntityData {
    authorization.EntityData.YFilter = authorization.YFilter
    authorization.EntityData.YangName = "authorization"
    authorization.EntityData.BundleName = "cisco_ios_xr"
    authorization.EntityData.ParentYangName = "aaa"
    authorization.EntityData.SegmentPath = "authorization"
    authorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorization.EntityData.Children = make(map[string]types.YChild)
    authorization.EntityData.Leafs = make(map[string]types.YLeaf)
    authorization.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", authorization.SentRequests}
    authorization.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", authorization.AcceptedRequests}
    authorization.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", authorization.SuccessfulRequests}
    authorization.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", authorization.RejectedRequests}
    authorization.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", authorization.UnreachableRequests}
    authorization.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", authorization.ErroredRequests}
    authorization.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", authorization.IncompleteRequests}
    authorization.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", authorization.TerminatedRequests}
    return &(authorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization
// Aggregate authorization statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (aggregateAuthorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAuthorization) GetEntityData() *types.CommonEntityData {
    aggregateAuthorization.EntityData.YFilter = aggregateAuthorization.YFilter
    aggregateAuthorization.EntityData.YangName = "aggregate-authorization"
    aggregateAuthorization.EntityData.BundleName = "cisco_ios_xr"
    aggregateAuthorization.EntityData.ParentYangName = "aaa"
    aggregateAuthorization.EntityData.SegmentPath = "aggregate-authorization"
    aggregateAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAuthorization.EntityData.Children = make(map[string]types.YChild)
    aggregateAuthorization.EntityData.Leafs = make(map[string]types.YLeaf)
    aggregateAuthorization.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", aggregateAuthorization.SentRequests}
    aggregateAuthorization.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", aggregateAuthorization.AcceptedRequests}
    aggregateAuthorization.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", aggregateAuthorization.SuccessfulRequests}
    aggregateAuthorization.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", aggregateAuthorization.RejectedRequests}
    aggregateAuthorization.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", aggregateAuthorization.UnreachableRequests}
    aggregateAuthorization.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", aggregateAuthorization.ErroredRequests}
    aggregateAuthorization.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", aggregateAuthorization.IncompleteRequests}
    aggregateAuthorization.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", aggregateAuthorization.TerminatedRequests}
    return &(aggregateAuthorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll
// Display all subscriber management total
// statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of stats for accounting.
    AccountingStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics

    // List of stats for authentication.
    AuthenticationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics

    // List of stats for authorization.
    AuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics

    // List of stats for COA.
    ChangeOfAuthorizationStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics

    // List of stats for Mobility.
    MobilityStatistics Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics
}

func (aggregateAccountingStatsAll *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll) GetEntityData() *types.CommonEntityData {
    aggregateAccountingStatsAll.EntityData.YFilter = aggregateAccountingStatsAll.YFilter
    aggregateAccountingStatsAll.EntityData.YangName = "aggregate-accounting-stats-all"
    aggregateAccountingStatsAll.EntityData.BundleName = "cisco_ios_xr"
    aggregateAccountingStatsAll.EntityData.ParentYangName = "aaa"
    aggregateAccountingStatsAll.EntityData.SegmentPath = "aggregate-accounting-stats-all"
    aggregateAccountingStatsAll.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateAccountingStatsAll.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateAccountingStatsAll.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateAccountingStatsAll.EntityData.Children = make(map[string]types.YChild)
    aggregateAccountingStatsAll.EntityData.Children["accounting-statistics"] = types.YChild{"AccountingStatistics", &aggregateAccountingStatsAll.AccountingStatistics}
    aggregateAccountingStatsAll.EntityData.Children["authentication-statistics"] = types.YChild{"AuthenticationStatistics", &aggregateAccountingStatsAll.AuthenticationStatistics}
    aggregateAccountingStatsAll.EntityData.Children["authorization-statistics"] = types.YChild{"AuthorizationStatistics", &aggregateAccountingStatsAll.AuthorizationStatistics}
    aggregateAccountingStatsAll.EntityData.Children["change-of-authorization-statistics"] = types.YChild{"ChangeOfAuthorizationStatistics", &aggregateAccountingStatsAll.ChangeOfAuthorizationStatistics}
    aggregateAccountingStatsAll.EntityData.Children["mobility-statistics"] = types.YChild{"MobilityStatistics", &aggregateAccountingStatsAll.MobilityStatistics}
    aggregateAccountingStatsAll.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(aggregateAccountingStatsAll.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics
// List of stats for accounting
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight
}

func (accountingStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics) GetEntityData() *types.CommonEntityData {
    accountingStatistics.EntityData.YFilter = accountingStatistics.YFilter
    accountingStatistics.EntityData.YangName = "accounting-statistics"
    accountingStatistics.EntityData.BundleName = "cisco_ios_xr"
    accountingStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    accountingStatistics.EntityData.SegmentPath = "accounting-statistics"
    accountingStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingStatistics.EntityData.Children = make(map[string]types.YChild)
    accountingStatistics.EntityData.Children["start"] = types.YChild{"Start", &accountingStatistics.Start}
    accountingStatistics.EntityData.Children["stop"] = types.YChild{"Stop", &accountingStatistics.Stop}
    accountingStatistics.EntityData.Children["interim"] = types.YChild{"Interim", &accountingStatistics.Interim}
    accountingStatistics.EntityData.Children["pass-through"] = types.YChild{"PassThrough", &accountingStatistics.PassThrough}
    accountingStatistics.EntityData.Children["update"] = types.YChild{"Update", &accountingStatistics.Update}
    accountingStatistics.EntityData.Children["interim-inflight"] = types.YChild{"InterimInflight", &accountingStatistics.InterimInflight}
    accountingStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    accountingStatistics.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", accountingStatistics.ActiveSessions}
    accountingStatistics.EntityData.Leafs["started-sessions"] = types.YLeaf{"StartedSessions", accountingStatistics.StartedSessions}
    accountingStatistics.EntityData.Leafs["stopped-sessions"] = types.YLeaf{"StoppedSessions", accountingStatistics.StoppedSessions}
    accountingStatistics.EntityData.Leafs["policy-plane-errored-requests"] = types.YLeaf{"PolicyPlaneErroredRequests", accountingStatistics.PolicyPlaneErroredRequests}
    accountingStatistics.EntityData.Leafs["policy-plane-unknown-requests"] = types.YLeaf{"PolicyPlaneUnknownRequests", accountingStatistics.PolicyPlaneUnknownRequests}
    return &(accountingStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "accounting-statistics"
    start.EntityData.SegmentPath = "start"
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = make(map[string]types.YChild)
    start.EntityData.Leafs = make(map[string]types.YLeaf)
    start.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", start.ReceivedRequests}
    start.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", start.ErroredRequests}
    start.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests}
    start.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", start.AaaSentRequests}
    start.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses}
    start.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses}
    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "accounting-statistics"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    stop.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", stop.ReceivedRequests}
    stop.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", stop.ErroredRequests}
    stop.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests}
    stop.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", stop.AaaSentRequests}
    stop.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses}
    stop.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses}
    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting-statistics"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = make(map[string]types.YChild)
    interim.EntityData.Leafs = make(map[string]types.YLeaf)
    interim.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", interim.ReceivedRequests}
    interim.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", interim.ErroredRequests}
    interim.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests}
    interim.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", interim.AaaSentRequests}
    interim.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses}
    interim.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses}
    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "accounting-statistics"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = make(map[string]types.YChild)
    passThrough.EntityData.Leafs = make(map[string]types.YLeaf)
    passThrough.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests}
    passThrough.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", passThrough.ErroredRequests}
    passThrough.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests}
    passThrough.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests}
    passThrough.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses}
    passThrough.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses}
    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "accounting-statistics"
    update.EntityData.SegmentPath = "update"
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = make(map[string]types.YChild)
    update.EntityData.Leafs = make(map[string]types.YLeaf)
    update.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", update.ReceivedRequests}
    update.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", update.ErroredRequests}
    update.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests}
    update.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", update.AaaSentRequests}
    update.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses}
    update.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses}
    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AccountingStatistics_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "accounting-statistics"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = make(map[string]types.YChild)
    interimInflight.EntityData.Leafs = make(map[string]types.YLeaf)
    interimInflight.EntityData.Leafs["quota-exhausts"] = types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts}
    interimInflight.EntityData.Leafs["denied-requests"] = types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests}
    interimInflight.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests}
    interimInflight.EntityData.Leafs["total-quota-of-requests"] = types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests}
    interimInflight.EntityData.Leafs["remaining-quota-of-requests"] = types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests}
    interimInflight.EntityData.Leafs["low-water-mark-quota-of-requests"] = types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests}
    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics
// List of stats for authentication
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authenticationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthenticationStatistics) GetEntityData() *types.CommonEntityData {
    authenticationStatistics.EntityData.YFilter = authenticationStatistics.YFilter
    authenticationStatistics.EntityData.YangName = "authentication-statistics"
    authenticationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authenticationStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    authenticationStatistics.EntityData.SegmentPath = "authentication-statistics"
    authenticationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationStatistics.EntityData.Children = make(map[string]types.YChild)
    authenticationStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    authenticationStatistics.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", authenticationStatistics.SentRequests}
    authenticationStatistics.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", authenticationStatistics.AcceptedRequests}
    authenticationStatistics.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", authenticationStatistics.SuccessfulRequests}
    authenticationStatistics.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", authenticationStatistics.RejectedRequests}
    authenticationStatistics.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", authenticationStatistics.UnreachableRequests}
    authenticationStatistics.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", authenticationStatistics.ErroredRequests}
    authenticationStatistics.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", authenticationStatistics.IncompleteRequests}
    authenticationStatistics.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", authenticationStatistics.TerminatedRequests}
    return &(authenticationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics
// List of stats for authorization
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Requests sent to radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    SentRequests interface{}

    // Request accepted by Radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    AcceptedRequests interface{}

    // Requests which are successful. The type is interface{} with range:
    // 0..18446744073709551615.
    SuccessfulRequests interface{}

    // Requests rejected by radius server. The type is interface{} with range:
    // 0..18446744073709551615.
    RejectedRequests interface{}

    // Radius server not available. The type is interface{} with range:
    // 0..18446744073709551615.
    UnreachableRequests interface{}

    // Unexpected errors. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // Incomplete requests - missing attributes. The type is interface{} with
    // range: 0..18446744073709551615.
    IncompleteRequests interface{}

    // Requests terminated by disconnect. The type is interface{} with range:
    // 0..18446744073709551615.
    TerminatedRequests interface{}
}

func (authorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_AuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    authorizationStatistics.EntityData.YFilter = authorizationStatistics.YFilter
    authorizationStatistics.EntityData.YangName = "authorization-statistics"
    authorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    authorizationStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    authorizationStatistics.EntityData.SegmentPath = "authorization-statistics"
    authorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorizationStatistics.EntityData.Children = make(map[string]types.YChild)
    authorizationStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    authorizationStatistics.EntityData.Leafs["sent-requests"] = types.YLeaf{"SentRequests", authorizationStatistics.SentRequests}
    authorizationStatistics.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", authorizationStatistics.AcceptedRequests}
    authorizationStatistics.EntityData.Leafs["successful-requests"] = types.YLeaf{"SuccessfulRequests", authorizationStatistics.SuccessfulRequests}
    authorizationStatistics.EntityData.Leafs["rejected-requests"] = types.YLeaf{"RejectedRequests", authorizationStatistics.RejectedRequests}
    authorizationStatistics.EntityData.Leafs["unreachable-requests"] = types.YLeaf{"UnreachableRequests", authorizationStatistics.UnreachableRequests}
    authorizationStatistics.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", authorizationStatistics.ErroredRequests}
    authorizationStatistics.EntityData.Leafs["incomplete-requests"] = types.YLeaf{"IncompleteRequests", authorizationStatistics.IncompleteRequests}
    authorizationStatistics.EntityData.Leafs["terminated-requests"] = types.YLeaf{"TerminatedRequests", authorizationStatistics.TerminatedRequests}
    return &(authorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics
// List of stats for COA
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
}

func (changeOfAuthorizationStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics) GetEntityData() *types.CommonEntityData {
    changeOfAuthorizationStatistics.EntityData.YFilter = changeOfAuthorizationStatistics.YFilter
    changeOfAuthorizationStatistics.EntityData.YangName = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.BundleName = "cisco_ios_xr"
    changeOfAuthorizationStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    changeOfAuthorizationStatistics.EntityData.SegmentPath = "change-of-authorization-statistics"
    changeOfAuthorizationStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    changeOfAuthorizationStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    changeOfAuthorizationStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    changeOfAuthorizationStatistics.EntityData.Children = make(map[string]types.YChild)
    changeOfAuthorizationStatistics.EntityData.Children["account-logon"] = types.YChild{"AccountLogon", &changeOfAuthorizationStatistics.AccountLogon}
    changeOfAuthorizationStatistics.EntityData.Children["account-logoff"] = types.YChild{"AccountLogoff", &changeOfAuthorizationStatistics.AccountLogoff}
    changeOfAuthorizationStatistics.EntityData.Children["account-update"] = types.YChild{"AccountUpdate", &changeOfAuthorizationStatistics.AccountUpdate}
    changeOfAuthorizationStatistics.EntityData.Children["session-disconnect"] = types.YChild{"SessionDisconnect", &changeOfAuthorizationStatistics.SessionDisconnect}
    changeOfAuthorizationStatistics.EntityData.Children["single-service-logon"] = types.YChild{"SingleServiceLogon", &changeOfAuthorizationStatistics.SingleServiceLogon}
    changeOfAuthorizationStatistics.EntityData.Children["single-service-logoff"] = types.YChild{"SingleServiceLogoff", &changeOfAuthorizationStatistics.SingleServiceLogoff}
    changeOfAuthorizationStatistics.EntityData.Children["single-service-modify"] = types.YChild{"SingleServiceModify", &changeOfAuthorizationStatistics.SingleServiceModify}
    changeOfAuthorizationStatistics.EntityData.Children["service-multi"] = types.YChild{"ServiceMulti", &changeOfAuthorizationStatistics.ServiceMulti}
    changeOfAuthorizationStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    changeOfAuthorizationStatistics.EntityData.Leafs["unknown-account-cmd-resps"] = types.YLeaf{"UnknownAccountCmdResps", changeOfAuthorizationStatistics.UnknownAccountCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["unknown-service-cmd-resps"] = types.YLeaf{"UnknownServiceCmdResps", changeOfAuthorizationStatistics.UnknownServiceCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["unknown-cmd-resps"] = types.YLeaf{"UnknownCmdResps", changeOfAuthorizationStatistics.UnknownCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["attr-list-retrieve-failure-resps"] = types.YLeaf{"AttrListRetrieveFailureResps", changeOfAuthorizationStatistics.AttrListRetrieveFailureResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["resp-send-failure"] = types.YLeaf{"RespSendFailure", changeOfAuthorizationStatistics.RespSendFailure}
    changeOfAuthorizationStatistics.EntityData.Leafs["internal-err-resps"] = types.YLeaf{"InternalErrResps", changeOfAuthorizationStatistics.InternalErrResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["service-profile-push-failure-resps"] = types.YLeaf{"ServiceProfilePushFailureResps", changeOfAuthorizationStatistics.ServiceProfilePushFailureResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["no-cmd-resps"] = types.YLeaf{"NoCmdResps", changeOfAuthorizationStatistics.NoCmdResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["no-session-found-resps"] = types.YLeaf{"NoSessionFoundResps", changeOfAuthorizationStatistics.NoSessionFoundResps}
    changeOfAuthorizationStatistics.EntityData.Leafs["no-session-peer-resps"] = types.YLeaf{"NoSessionPeerResps", changeOfAuthorizationStatistics.NoSessionPeerResps}
    return &(changeOfAuthorizationStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = make(map[string]types.YChild)
    accountLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests}
    accountLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests}
    accountLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests}
    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = make(map[string]types.YChild)
    accountLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests}
    accountLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests}
    accountLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests}
    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "change-of-authorization-statistics"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = make(map[string]types.YChild)
    accountUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    accountUpdate.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests}
    accountUpdate.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests}
    accountUpdate.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests}
    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "change-of-authorization-statistics"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = make(map[string]types.YChild)
    sessionDisconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDisconnect.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests}
    sessionDisconnect.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests}
    sessionDisconnect.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests}
    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests}
    singleServiceLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests}
    singleServiceLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests}
    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests}
    singleServiceLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests}
    singleServiceLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests}
    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "change-of-authorization-statistics"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = make(map[string]types.YChild)
    singleServiceModify.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceModify.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests}
    singleServiceModify.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests}
    singleServiceModify.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests}
    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_ChangeOfAuthorizationStatistics_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "change-of-authorization-statistics"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = make(map[string]types.YChild)
    serviceMulti.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceMulti.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests}
    serviceMulti.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests}
    serviceMulti.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests}
    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics
// List of stats for Mobility
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (mobilityStatistics *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateAccountingStatsAll_MobilityStatistics) GetEntityData() *types.CommonEntityData {
    mobilityStatistics.EntityData.YFilter = mobilityStatistics.YFilter
    mobilityStatistics.EntityData.YangName = "mobility-statistics"
    mobilityStatistics.EntityData.BundleName = "cisco_ios_xr"
    mobilityStatistics.EntityData.ParentYangName = "aggregate-accounting-stats-all"
    mobilityStatistics.EntityData.SegmentPath = "mobility-statistics"
    mobilityStatistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobilityStatistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobilityStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobilityStatistics.EntityData.Children = make(map[string]types.YChild)
    mobilityStatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    mobilityStatistics.EntityData.Leafs["send-request-successes"] = types.YLeaf{"SendRequestSuccesses", mobilityStatistics.SendRequestSuccesses}
    mobilityStatistics.EntityData.Leafs["send-request-failures"] = types.YLeaf{"SendRequestFailures", mobilityStatistics.SendRequestFailures}
    mobilityStatistics.EntityData.Leafs["receive-response-successes"] = types.YLeaf{"ReceiveResponseSuccesses", mobilityStatistics.ReceiveResponseSuccesses}
    mobilityStatistics.EntityData.Leafs["receive-response-failures"] = types.YLeaf{"ReceiveResponseFailures", mobilityStatistics.ReceiveResponseFailures}
    return &(mobilityStatistics.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting
// Accounting statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Active sessions. The type is interface{} with range: 0..4294967295.
    ActiveSessions interface{}

    // Started sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StartedSessions interface{}

    // Stopped sessions. The type is interface{} with range:
    // 0..18446744073709551615.
    StoppedSessions interface{}

    // Policy plane errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneErroredRequests interface{}

    // Policy plane unknown requests. The type is interface{} with range:
    // 0..18446744073709551615.
    PolicyPlaneUnknownRequests interface{}

    // Start statistics.
    Start Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start

    // Stop statistics.
    Stop Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop

    // Interim statistics.
    Interim Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim

    // Pass-through statistics.
    PassThrough Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough

    // Update statistics.
    Update Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update

    // Interim inflight details.
    InterimInflight Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight
}

func (accounting *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "aaa"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = make(map[string]types.YChild)
    accounting.EntityData.Children["start"] = types.YChild{"Start", &accounting.Start}
    accounting.EntityData.Children["stop"] = types.YChild{"Stop", &accounting.Stop}
    accounting.EntityData.Children["interim"] = types.YChild{"Interim", &accounting.Interim}
    accounting.EntityData.Children["pass-through"] = types.YChild{"PassThrough", &accounting.PassThrough}
    accounting.EntityData.Children["update"] = types.YChild{"Update", &accounting.Update}
    accounting.EntityData.Children["interim-inflight"] = types.YChild{"InterimInflight", &accounting.InterimInflight}
    accounting.EntityData.Leafs = make(map[string]types.YLeaf)
    accounting.EntityData.Leafs["active-sessions"] = types.YLeaf{"ActiveSessions", accounting.ActiveSessions}
    accounting.EntityData.Leafs["started-sessions"] = types.YLeaf{"StartedSessions", accounting.StartedSessions}
    accounting.EntityData.Leafs["stopped-sessions"] = types.YLeaf{"StoppedSessions", accounting.StoppedSessions}
    accounting.EntityData.Leafs["policy-plane-errored-requests"] = types.YLeaf{"PolicyPlaneErroredRequests", accounting.PolicyPlaneErroredRequests}
    accounting.EntityData.Leafs["policy-plane-unknown-requests"] = types.YLeaf{"PolicyPlaneUnknownRequests", accounting.PolicyPlaneUnknownRequests}
    return &(accounting.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start
// Start statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (start *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Start) GetEntityData() *types.CommonEntityData {
    start.EntityData.YFilter = start.YFilter
    start.EntityData.YangName = "start"
    start.EntityData.BundleName = "cisco_ios_xr"
    start.EntityData.ParentYangName = "accounting"
    start.EntityData.SegmentPath = "start"
    start.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    start.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    start.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    start.EntityData.Children = make(map[string]types.YChild)
    start.EntityData.Leafs = make(map[string]types.YLeaf)
    start.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", start.ReceivedRequests}
    start.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", start.ErroredRequests}
    start.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", start.AaaErroredRequests}
    start.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", start.AaaSentRequests}
    start.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", start.AaaSucceededResponses}
    start.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", start.AaaFailedResponses}
    return &(start.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop
// Stop statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (stop *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Stop) GetEntityData() *types.CommonEntityData {
    stop.EntityData.YFilter = stop.YFilter
    stop.EntityData.YangName = "stop"
    stop.EntityData.BundleName = "cisco_ios_xr"
    stop.EntityData.ParentYangName = "accounting"
    stop.EntityData.SegmentPath = "stop"
    stop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stop.EntityData.Children = make(map[string]types.YChild)
    stop.EntityData.Leafs = make(map[string]types.YLeaf)
    stop.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", stop.ReceivedRequests}
    stop.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", stop.ErroredRequests}
    stop.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", stop.AaaErroredRequests}
    stop.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", stop.AaaSentRequests}
    stop.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", stop.AaaSucceededResponses}
    stop.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", stop.AaaFailedResponses}
    return &(stop.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim
// Interim statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (interim *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Interim) GetEntityData() *types.CommonEntityData {
    interim.EntityData.YFilter = interim.YFilter
    interim.EntityData.YangName = "interim"
    interim.EntityData.BundleName = "cisco_ios_xr"
    interim.EntityData.ParentYangName = "accounting"
    interim.EntityData.SegmentPath = "interim"
    interim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interim.EntityData.Children = make(map[string]types.YChild)
    interim.EntityData.Leafs = make(map[string]types.YLeaf)
    interim.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", interim.ReceivedRequests}
    interim.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", interim.ErroredRequests}
    interim.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", interim.AaaErroredRequests}
    interim.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", interim.AaaSentRequests}
    interim.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", interim.AaaSucceededResponses}
    interim.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", interim.AaaFailedResponses}
    return &(interim.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough
// Pass-through statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (passThrough *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_PassThrough) GetEntityData() *types.CommonEntityData {
    passThrough.EntityData.YFilter = passThrough.YFilter
    passThrough.EntityData.YangName = "pass-through"
    passThrough.EntityData.BundleName = "cisco_ios_xr"
    passThrough.EntityData.ParentYangName = "accounting"
    passThrough.EntityData.SegmentPath = "pass-through"
    passThrough.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    passThrough.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    passThrough.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    passThrough.EntityData.Children = make(map[string]types.YChild)
    passThrough.EntityData.Leafs = make(map[string]types.YLeaf)
    passThrough.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", passThrough.ReceivedRequests}
    passThrough.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", passThrough.ErroredRequests}
    passThrough.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", passThrough.AaaErroredRequests}
    passThrough.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", passThrough.AaaSentRequests}
    passThrough.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", passThrough.AaaSucceededResponses}
    passThrough.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", passThrough.AaaFailedResponses}
    return &(passThrough.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update
// Update statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ErroredRequests interface{}

    // AAA errored requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaErroredRequests interface{}

    // AAA requests sent. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSentRequests interface{}

    // AAA succeeded responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaSucceededResponses interface{}

    // AAA failed responses. The type is interface{} with range:
    // 0..18446744073709551615.
    AaaFailedResponses interface{}
}

func (update *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_Update) GetEntityData() *types.CommonEntityData {
    update.EntityData.YFilter = update.YFilter
    update.EntityData.YangName = "update"
    update.EntityData.BundleName = "cisco_ios_xr"
    update.EntityData.ParentYangName = "accounting"
    update.EntityData.SegmentPath = "update"
    update.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    update.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    update.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    update.EntityData.Children = make(map[string]types.YChild)
    update.EntityData.Leafs = make(map[string]types.YLeaf)
    update.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", update.ReceivedRequests}
    update.EntityData.Leafs["errored-requests"] = types.YLeaf{"ErroredRequests", update.ErroredRequests}
    update.EntityData.Leafs["aaa-errored-requests"] = types.YLeaf{"AaaErroredRequests", update.AaaErroredRequests}
    update.EntityData.Leafs["aaa-sent-requests"] = types.YLeaf{"AaaSentRequests", update.AaaSentRequests}
    update.EntityData.Leafs["aaa-succeeded-responses"] = types.YLeaf{"AaaSucceededResponses", update.AaaSucceededResponses}
    update.EntityData.Leafs["aaa-failed-responses"] = types.YLeaf{"AaaFailedResponses", update.AaaFailedResponses}
    return &(update.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight
// Interim inflight details
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Quota exhausts. The type is interface{} with range: 0..4294967295.
    QuotaExhausts interface{}

    // Denied requests. The type is interface{} with range: 0..4294967295.
    DeniedRequests interface{}

    // Accepted requests. The type is interface{} with range: 0..4294967295.
    AcceptedRequests interface{}

    // Total quota of requests. The type is interface{} with range: 0..4294967295.
    TotalQuotaOfRequests interface{}

    // Remaining quota of requests. The type is interface{} with range:
    // 0..4294967295.
    RemainingQuotaOfRequests interface{}

    // Low water mark quota of requests. The type is interface{} with range:
    // 0..4294967295.
    LowWaterMarkQuotaOfRequests interface{}
}

func (interimInflight *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Accounting_InterimInflight) GetEntityData() *types.CommonEntityData {
    interimInflight.EntityData.YFilter = interimInflight.YFilter
    interimInflight.EntityData.YangName = "interim-inflight"
    interimInflight.EntityData.BundleName = "cisco_ios_xr"
    interimInflight.EntityData.ParentYangName = "accounting"
    interimInflight.EntityData.SegmentPath = "interim-inflight"
    interimInflight.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interimInflight.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interimInflight.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interimInflight.EntityData.Children = make(map[string]types.YChild)
    interimInflight.EntityData.Leafs = make(map[string]types.YLeaf)
    interimInflight.EntityData.Leafs["quota-exhausts"] = types.YLeaf{"QuotaExhausts", interimInflight.QuotaExhausts}
    interimInflight.EntityData.Leafs["denied-requests"] = types.YLeaf{"DeniedRequests", interimInflight.DeniedRequests}
    interimInflight.EntityData.Leafs["accepted-requests"] = types.YLeaf{"AcceptedRequests", interimInflight.AcceptedRequests}
    interimInflight.EntityData.Leafs["total-quota-of-requests"] = types.YLeaf{"TotalQuotaOfRequests", interimInflight.TotalQuotaOfRequests}
    interimInflight.EntityData.Leafs["remaining-quota-of-requests"] = types.YLeaf{"RemainingQuotaOfRequests", interimInflight.RemainingQuotaOfRequests}
    interimInflight.EntityData.Leafs["low-water-mark-quota-of-requests"] = types.YLeaf{"LowWaterMarkQuotaOfRequests", interimInflight.LowWaterMarkQuotaOfRequests}
    return &(interimInflight.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility
// Mobility statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Request send success. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestSuccesses interface{}

    // Request send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    SendRequestFailures interface{}

    // Response receive success. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseSuccesses interface{}

    // Response receive failures. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceiveResponseFailures interface{}
}

func (mobility *Subscriber_Manager_Nodes_Node_Statistics_Aaa_Mobility) GetEntityData() *types.CommonEntityData {
    mobility.EntityData.YFilter = mobility.YFilter
    mobility.EntityData.YangName = "mobility"
    mobility.EntityData.BundleName = "cisco_ios_xr"
    mobility.EntityData.ParentYangName = "aaa"
    mobility.EntityData.SegmentPath = "mobility"
    mobility.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobility.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobility.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobility.EntityData.Children = make(map[string]types.YChild)
    mobility.EntityData.Leafs = make(map[string]types.YLeaf)
    mobility.EntityData.Leafs["send-request-successes"] = types.YLeaf{"SendRequestSuccesses", mobility.SendRequestSuccesses}
    mobility.EntityData.Leafs["send-request-failures"] = types.YLeaf{"SendRequestFailures", mobility.SendRequestFailures}
    mobility.EntityData.Leafs["receive-response-successes"] = types.YLeaf{"ReceiveResponseSuccesses", mobility.ReceiveResponseSuccesses}
    mobility.EntityData.Leafs["receive-response-failures"] = types.YLeaf{"ReceiveResponseFailures", mobility.ReceiveResponseFailures}
    return &(mobility.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization
// Aggregate change of authorization (COA)
// statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Responses to unknown account command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownAccountCmdResps interface{}

    // Responses to unknown service command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownServiceCmdResps interface{}

    // Responses to unknown command. The type is interface{} with range:
    // 0..18446744073709551615.
    UnknownCmdResps interface{}

    // Responses to attribute list failure errors. The type is interface{} with
    // range: 0..18446744073709551615.
    AttrListRetrieveFailureResps interface{}

    // Response send failures. The type is interface{} with range:
    // 0..18446744073709551615.
    RespSendFailure interface{}

    // Responses to internal error. The type is interface{} with range:
    // 0..18446744073709551615.
    InternalErrResps interface{}

    // Responses to service profile push failures. The type is interface{} with
    // range: 0..18446744073709551615.
    ServiceProfilePushFailureResps interface{}

    // Responses empty (no command) COA request. The type is interface{} with
    // range: 0..18446744073709551615.
    NoCmdResps interface{}

    // Responses to COA with unknown session identifier. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSessionFoundResps interface{}

    // Responses to session peer not found error. The type is interface{} with
    // range: 0..18446744073709551615.
    NoSessionPeerResps interface{}

    // Account logon request statistics.
    AccountLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon

    // Account logoff request statistics.
    AccountLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff

    // Account update request statistics.
    AccountUpdate Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate

    // Session disconnect request statistics.
    SessionDisconnect Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect

    // Service logon request statistics.
    SingleServiceLogon Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon

    // Single Service logoff request statistics.
    SingleServiceLogoff Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff

    // Single Service Modify request statistics.
    SingleServiceModify Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify

    // MA-CoA Service request statistics.
    ServiceMulti Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti
}

func (aggregateChangeOfAuthorization *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    aggregateChangeOfAuthorization.EntityData.YFilter = aggregateChangeOfAuthorization.YFilter
    aggregateChangeOfAuthorization.EntityData.YangName = "aggregate-change-of-authorization"
    aggregateChangeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    aggregateChangeOfAuthorization.EntityData.ParentYangName = "aaa"
    aggregateChangeOfAuthorization.EntityData.SegmentPath = "aggregate-change-of-authorization"
    aggregateChangeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateChangeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateChangeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateChangeOfAuthorization.EntityData.Children = make(map[string]types.YChild)
    aggregateChangeOfAuthorization.EntityData.Children["account-logon"] = types.YChild{"AccountLogon", &aggregateChangeOfAuthorization.AccountLogon}
    aggregateChangeOfAuthorization.EntityData.Children["account-logoff"] = types.YChild{"AccountLogoff", &aggregateChangeOfAuthorization.AccountLogoff}
    aggregateChangeOfAuthorization.EntityData.Children["account-update"] = types.YChild{"AccountUpdate", &aggregateChangeOfAuthorization.AccountUpdate}
    aggregateChangeOfAuthorization.EntityData.Children["session-disconnect"] = types.YChild{"SessionDisconnect", &aggregateChangeOfAuthorization.SessionDisconnect}
    aggregateChangeOfAuthorization.EntityData.Children["single-service-logon"] = types.YChild{"SingleServiceLogon", &aggregateChangeOfAuthorization.SingleServiceLogon}
    aggregateChangeOfAuthorization.EntityData.Children["single-service-logoff"] = types.YChild{"SingleServiceLogoff", &aggregateChangeOfAuthorization.SingleServiceLogoff}
    aggregateChangeOfAuthorization.EntityData.Children["single-service-modify"] = types.YChild{"SingleServiceModify", &aggregateChangeOfAuthorization.SingleServiceModify}
    aggregateChangeOfAuthorization.EntityData.Children["service-multi"] = types.YChild{"ServiceMulti", &aggregateChangeOfAuthorization.ServiceMulti}
    aggregateChangeOfAuthorization.EntityData.Leafs = make(map[string]types.YLeaf)
    aggregateChangeOfAuthorization.EntityData.Leafs["unknown-account-cmd-resps"] = types.YLeaf{"UnknownAccountCmdResps", aggregateChangeOfAuthorization.UnknownAccountCmdResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["unknown-service-cmd-resps"] = types.YLeaf{"UnknownServiceCmdResps", aggregateChangeOfAuthorization.UnknownServiceCmdResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["unknown-cmd-resps"] = types.YLeaf{"UnknownCmdResps", aggregateChangeOfAuthorization.UnknownCmdResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["attr-list-retrieve-failure-resps"] = types.YLeaf{"AttrListRetrieveFailureResps", aggregateChangeOfAuthorization.AttrListRetrieveFailureResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["resp-send-failure"] = types.YLeaf{"RespSendFailure", aggregateChangeOfAuthorization.RespSendFailure}
    aggregateChangeOfAuthorization.EntityData.Leafs["internal-err-resps"] = types.YLeaf{"InternalErrResps", aggregateChangeOfAuthorization.InternalErrResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["service-profile-push-failure-resps"] = types.YLeaf{"ServiceProfilePushFailureResps", aggregateChangeOfAuthorization.ServiceProfilePushFailureResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["no-cmd-resps"] = types.YLeaf{"NoCmdResps", aggregateChangeOfAuthorization.NoCmdResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["no-session-found-resps"] = types.YLeaf{"NoSessionFoundResps", aggregateChangeOfAuthorization.NoSessionFoundResps}
    aggregateChangeOfAuthorization.EntityData.Leafs["no-session-peer-resps"] = types.YLeaf{"NoSessionPeerResps", aggregateChangeOfAuthorization.NoSessionPeerResps}
    return &(aggregateChangeOfAuthorization.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon
// Account logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogon) GetEntityData() *types.CommonEntityData {
    accountLogon.EntityData.YFilter = accountLogon.YFilter
    accountLogon.EntityData.YangName = "account-logon"
    accountLogon.EntityData.BundleName = "cisco_ios_xr"
    accountLogon.EntityData.ParentYangName = "aggregate-change-of-authorization"
    accountLogon.EntityData.SegmentPath = "account-logon"
    accountLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogon.EntityData.Children = make(map[string]types.YChild)
    accountLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogon.ReceivedRequests}
    accountLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogon.AcknowledgedRequests}
    accountLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogon.NonAcknowledgedRequests}
    return &(accountLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff
// Account logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountLogoff) GetEntityData() *types.CommonEntityData {
    accountLogoff.EntityData.YFilter = accountLogoff.YFilter
    accountLogoff.EntityData.YangName = "account-logoff"
    accountLogoff.EntityData.BundleName = "cisco_ios_xr"
    accountLogoff.EntityData.ParentYangName = "aggregate-change-of-authorization"
    accountLogoff.EntityData.SegmentPath = "account-logoff"
    accountLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountLogoff.EntityData.Children = make(map[string]types.YChild)
    accountLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    accountLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountLogoff.ReceivedRequests}
    accountLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountLogoff.AcknowledgedRequests}
    accountLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountLogoff.NonAcknowledgedRequests}
    return &(accountLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate
// Account update request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (accountUpdate *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_AccountUpdate) GetEntityData() *types.CommonEntityData {
    accountUpdate.EntityData.YFilter = accountUpdate.YFilter
    accountUpdate.EntityData.YangName = "account-update"
    accountUpdate.EntityData.BundleName = "cisco_ios_xr"
    accountUpdate.EntityData.ParentYangName = "aggregate-change-of-authorization"
    accountUpdate.EntityData.SegmentPath = "account-update"
    accountUpdate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountUpdate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountUpdate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountUpdate.EntityData.Children = make(map[string]types.YChild)
    accountUpdate.EntityData.Leafs = make(map[string]types.YLeaf)
    accountUpdate.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", accountUpdate.ReceivedRequests}
    accountUpdate.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", accountUpdate.AcknowledgedRequests}
    accountUpdate.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", accountUpdate.NonAcknowledgedRequests}
    return &(accountUpdate.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect
// Session disconnect request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (sessionDisconnect *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SessionDisconnect) GetEntityData() *types.CommonEntityData {
    sessionDisconnect.EntityData.YFilter = sessionDisconnect.YFilter
    sessionDisconnect.EntityData.YangName = "session-disconnect"
    sessionDisconnect.EntityData.BundleName = "cisco_ios_xr"
    sessionDisconnect.EntityData.ParentYangName = "aggregate-change-of-authorization"
    sessionDisconnect.EntityData.SegmentPath = "session-disconnect"
    sessionDisconnect.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionDisconnect.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionDisconnect.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionDisconnect.EntityData.Children = make(map[string]types.YChild)
    sessionDisconnect.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionDisconnect.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", sessionDisconnect.ReceivedRequests}
    sessionDisconnect.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", sessionDisconnect.AcknowledgedRequests}
    sessionDisconnect.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", sessionDisconnect.NonAcknowledgedRequests}
    return &(sessionDisconnect.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon
// Service logon request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogon *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogon) GetEntityData() *types.CommonEntityData {
    singleServiceLogon.EntityData.YFilter = singleServiceLogon.YFilter
    singleServiceLogon.EntityData.YangName = "single-service-logon"
    singleServiceLogon.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogon.EntityData.ParentYangName = "aggregate-change-of-authorization"
    singleServiceLogon.EntityData.SegmentPath = "single-service-logon"
    singleServiceLogon.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogon.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogon.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogon.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogon.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogon.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogon.ReceivedRequests}
    singleServiceLogon.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogon.AcknowledgedRequests}
    singleServiceLogon.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogon.NonAcknowledgedRequests}
    return &(singleServiceLogon.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff
// Single Service logoff request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceLogoff *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceLogoff) GetEntityData() *types.CommonEntityData {
    singleServiceLogoff.EntityData.YFilter = singleServiceLogoff.YFilter
    singleServiceLogoff.EntityData.YangName = "single-service-logoff"
    singleServiceLogoff.EntityData.BundleName = "cisco_ios_xr"
    singleServiceLogoff.EntityData.ParentYangName = "aggregate-change-of-authorization"
    singleServiceLogoff.EntityData.SegmentPath = "single-service-logoff"
    singleServiceLogoff.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceLogoff.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceLogoff.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceLogoff.EntityData.Children = make(map[string]types.YChild)
    singleServiceLogoff.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceLogoff.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceLogoff.ReceivedRequests}
    singleServiceLogoff.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceLogoff.AcknowledgedRequests}
    singleServiceLogoff.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceLogoff.NonAcknowledgedRequests}
    return &(singleServiceLogoff.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify
// Single Service Modify request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (singleServiceModify *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_SingleServiceModify) GetEntityData() *types.CommonEntityData {
    singleServiceModify.EntityData.YFilter = singleServiceModify.YFilter
    singleServiceModify.EntityData.YangName = "single-service-modify"
    singleServiceModify.EntityData.BundleName = "cisco_ios_xr"
    singleServiceModify.EntityData.ParentYangName = "aggregate-change-of-authorization"
    singleServiceModify.EntityData.SegmentPath = "single-service-modify"
    singleServiceModify.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    singleServiceModify.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    singleServiceModify.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    singleServiceModify.EntityData.Children = make(map[string]types.YChild)
    singleServiceModify.EntityData.Leafs = make(map[string]types.YLeaf)
    singleServiceModify.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", singleServiceModify.ReceivedRequests}
    singleServiceModify.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", singleServiceModify.AcknowledgedRequests}
    singleServiceModify.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", singleServiceModify.NonAcknowledgedRequests}
    return &(singleServiceModify.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti
// MA-CoA Service request statistics
type Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Received requests. The type is interface{} with range:
    // 0..18446744073709551615.
    ReceivedRequests interface{}

    // Acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    AcknowledgedRequests interface{}

    // Non acknowledged requests. The type is interface{} with range:
    // 0..18446744073709551615.
    NonAcknowledgedRequests interface{}
}

func (serviceMulti *Subscriber_Manager_Nodes_Node_Statistics_Aaa_AggregateChangeOfAuthorization_ServiceMulti) GetEntityData() *types.CommonEntityData {
    serviceMulti.EntityData.YFilter = serviceMulti.YFilter
    serviceMulti.EntityData.YangName = "service-multi"
    serviceMulti.EntityData.BundleName = "cisco_ios_xr"
    serviceMulti.EntityData.ParentYangName = "aggregate-change-of-authorization"
    serviceMulti.EntityData.SegmentPath = "service-multi"
    serviceMulti.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serviceMulti.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serviceMulti.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serviceMulti.EntityData.Children = make(map[string]types.YChild)
    serviceMulti.EntityData.Leafs = make(map[string]types.YLeaf)
    serviceMulti.EntityData.Leafs["received-requests"] = types.YLeaf{"ReceivedRequests", serviceMulti.ReceivedRequests}
    serviceMulti.EntityData.Leafs["acknowledged-requests"] = types.YLeaf{"AcknowledgedRequests", serviceMulti.AcknowledgedRequests}
    serviceMulti.EntityData.Leafs["non-acknowledged-requests"] = types.YLeaf{"NonAcknowledgedRequests", serviceMulti.NonAcknowledgedRequests}
    return &(serviceMulti.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary
// Aggregate summary of statistics
type Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber control policy not applied on interface. The type is interface{}
    // with range: 0..18446744073709551615.
    NoSubscriberControlPolicyOnInterface interface{}

    // No control policy class match during subscriber start. The type is
    // interface{} with range: 0..18446744073709551615.
    NoClassMatchInStartRequest interface{}

    // NAS port attribute format warnings. The type is interface{} with range:
    // 0..18446744073709551615.
    NasPortAttributeFormatWarnings interface{}

    // NAS port ID attribute format warnings. The type is interface{} with range:
    // 0..18446744073709551615.
    NasPortIdAttributeFormatWarnings interface{}

    // Destination station ID attribute format warnings. The type is interface{}
    // with range: 0..18446744073709551615.
    DestinationStationIdAttributeFormatWarnings interface{}

    // Calling station ID attribute format warnings. The type is interface{} with
    // range: 0..18446744073709551615.
    CallingStationIdAttributeFormatWarnings interface{}

    // Username attribute format warnings. The type is interface{} with range:
    // 0..18446744073709551615.
    UsernameAttributeFormatWarnings interface{}

    // User profiles installed. The type is interface{} with range:
    // 0..18446744073709551615.
    InstallUserProfiles interface{}

    // User profile install errors. The type is interface{} with range:
    // 0..18446744073709551615.
    UserProfileInstallErrors interface{}

    // User profile removals. The type is interface{} with range:
    // 0..18446744073709551615.
    UserProfileRemovals interface{}

    // User profile errors. The type is interface{} with range:
    // 0..18446744073709551615.
    UserProfileErrors interface{}

    // Session Disconnect Quota Exhausts. The type is interface{} with range:
    // 0..18446744073709551615.
    SessDiscQuotaExhausts interface{}

    // Session Disconnect Request Queued, no quota. The type is interface{} with
    // range: 0..18446744073709551615.
    SessDiscNoQuota interface{}

    // Session Disconnect Request Accepted, quota available. The type is
    // interface{} with range: 0..18446744073709551615.
    SessDiscQuotaAvail interface{}

    // Session Disconnect Requests not Dequeued, reconciliation in progress. The
    // type is interface{} with range: 0..18446744073709551615.
    SessDiscReconIp interface{}

    // Session Disconnect Requests not Dequeued, no quota. The type is interface{}
    // with range: 0..18446744073709551615.
    SessDiscNoneStarted interface{}

    // Session Disconnect Quota. The type is interface{} with range:
    // 0..4294967295.
    SessDiscQuota interface{}

    // Session Disconnect Quota Remaining. The type is interface{} with range:
    // 0..4294967295.
    SessDiscQuotaRemaining interface{}

    // Session Disconnect Requests Queued. The type is interface{} with range:
    // 0..4294967295.
    SessDiscQCount interface{}
}

func (aggregateSummary *Subscriber_Manager_Nodes_Node_Statistics_AggregateSummary) GetEntityData() *types.CommonEntityData {
    aggregateSummary.EntityData.YFilter = aggregateSummary.YFilter
    aggregateSummary.EntityData.YangName = "aggregate-summary"
    aggregateSummary.EntityData.BundleName = "cisco_ios_xr"
    aggregateSummary.EntityData.ParentYangName = "statistics"
    aggregateSummary.EntityData.SegmentPath = "aggregate-summary"
    aggregateSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aggregateSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aggregateSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aggregateSummary.EntityData.Children = make(map[string]types.YChild)
    aggregateSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    aggregateSummary.EntityData.Leafs["no-subscriber-control-policy-on-interface"] = types.YLeaf{"NoSubscriberControlPolicyOnInterface", aggregateSummary.NoSubscriberControlPolicyOnInterface}
    aggregateSummary.EntityData.Leafs["no-class-match-in-start-request"] = types.YLeaf{"NoClassMatchInStartRequest", aggregateSummary.NoClassMatchInStartRequest}
    aggregateSummary.EntityData.Leafs["nas-port-attribute-format-warnings"] = types.YLeaf{"NasPortAttributeFormatWarnings", aggregateSummary.NasPortAttributeFormatWarnings}
    aggregateSummary.EntityData.Leafs["nas-port-id-attribute-format-warnings"] = types.YLeaf{"NasPortIdAttributeFormatWarnings", aggregateSummary.NasPortIdAttributeFormatWarnings}
    aggregateSummary.EntityData.Leafs["destination-station-id-attribute-format-warnings"] = types.YLeaf{"DestinationStationIdAttributeFormatWarnings", aggregateSummary.DestinationStationIdAttributeFormatWarnings}
    aggregateSummary.EntityData.Leafs["calling-station-id-attribute-format-warnings"] = types.YLeaf{"CallingStationIdAttributeFormatWarnings", aggregateSummary.CallingStationIdAttributeFormatWarnings}
    aggregateSummary.EntityData.Leafs["username-attribute-format-warnings"] = types.YLeaf{"UsernameAttributeFormatWarnings", aggregateSummary.UsernameAttributeFormatWarnings}
    aggregateSummary.EntityData.Leafs["install-user-profiles"] = types.YLeaf{"InstallUserProfiles", aggregateSummary.InstallUserProfiles}
    aggregateSummary.EntityData.Leafs["user-profile-install-errors"] = types.YLeaf{"UserProfileInstallErrors", aggregateSummary.UserProfileInstallErrors}
    aggregateSummary.EntityData.Leafs["user-profile-removals"] = types.YLeaf{"UserProfileRemovals", aggregateSummary.UserProfileRemovals}
    aggregateSummary.EntityData.Leafs["user-profile-errors"] = types.YLeaf{"UserProfileErrors", aggregateSummary.UserProfileErrors}
    aggregateSummary.EntityData.Leafs["sess-disc-quota-exhausts"] = types.YLeaf{"SessDiscQuotaExhausts", aggregateSummary.SessDiscQuotaExhausts}
    aggregateSummary.EntityData.Leafs["sess-disc-no-quota"] = types.YLeaf{"SessDiscNoQuota", aggregateSummary.SessDiscNoQuota}
    aggregateSummary.EntityData.Leafs["sess-disc-quota-avail"] = types.YLeaf{"SessDiscQuotaAvail", aggregateSummary.SessDiscQuotaAvail}
    aggregateSummary.EntityData.Leafs["sess-disc-recon-ip"] = types.YLeaf{"SessDiscReconIp", aggregateSummary.SessDiscReconIp}
    aggregateSummary.EntityData.Leafs["sess-disc-none-started"] = types.YLeaf{"SessDiscNoneStarted", aggregateSummary.SessDiscNoneStarted}
    aggregateSummary.EntityData.Leafs["sess-disc-quota"] = types.YLeaf{"SessDiscQuota", aggregateSummary.SessDiscQuota}
    aggregateSummary.EntityData.Leafs["sess-disc-quota-remaining"] = types.YLeaf{"SessDiscQuotaRemaining", aggregateSummary.SessDiscQuotaRemaining}
    aggregateSummary.EntityData.Leafs["sess-disc-q-count"] = types.YLeaf{"SessDiscQCount", aggregateSummary.SessDiscQCount}
    return &(aggregateSummary.EntityData)
}

// Subscriber_Manager_Nodes_Node_Statistics_Srg
// Geo Redundancy statistics
type Subscriber_Manager_Nodes_Node_Statistics_Srg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Txlist Send Triggered. The type is interface{} with range: 0..4294967295.
    TxlistSendTriggered interface{}

    // Txlist Send Failed. The type is interface{} with range: 0..4294967295.
    TxlistSendFailed interface{}

    // Txlist send failed due to not active. The type is interface{} with range:
    // 0..4294967295.
    TxlistSendFailedNotactive interface{}

    // Txlist Send Success. The type is interface{} with range: 0..4294967295.
    ActualTxlistSent interface{}

    // Element already in Txlist. The type is interface{} with range:
    // 0..4294967295.
    AlreadyinTxlist interface{}

    // Txlist Encode. The type is interface{} with range: 0..4294967295.
    TxlistEncode interface{}

    // Txlist encode Failed. The type is interface{} with range: 0..4294967295.
    TxlistEncodeFail interface{}

    // Txlist Create Update Encode. The type is interface{} with range:
    // 0..4294967295.
    CreateUpdateEncode interface{}

    // Txlist Delete Encode. The type is interface{} with range: 0..4294967295.
    DeleteEncode interface{}

    // Txlist Create/update clean callback. The type is interface{} with range:
    // 0..4294967295.
    CreateUpdCleanCallback interface{}

    // Txlist Delete clean callback. The type is interface{} with range:
    // 0..4294967295.
    DeleteCleanCallback interface{}

    // Slave Recieved Sync. The type is interface{} with range: 0..4294967295.
    SlaveRecvEntry interface{}

    // Decode failed on Slave. The type is interface{} with range: 0..4294967295.
    SlaveDecodeFail interface{}

    // Create Update received on slave. The type is interface{} with range:
    // 0..4294967295.
    SlaveCreateUpdate interface{}

    // Delete received on slave. The type is interface{} with range:
    // 0..4294967295.
    SlaveDelete interface{}

    // SRG context allocated. The type is interface{} with range: 0..4294967295.
    SrgContextMalloc interface{}

    // SRG context freed. The type is interface{} with range: 0..4294967295.
    SrgContextFree interface{}

    // Number of SODs Received. The type is interface{} with range: 0..4294967295.
    SodCount interface{}

    // Number of EODs Received. The type is interface{} with range: 0..4294967295.
    EodCount interface{}

    // Number of Replay Requests Within SOD EOD Window. The type is interface{}
    // with range: 0..4294967295.
    SodEodReplayReqCount interface{}

    // Number of Sessions Marked as Invalid Within SOD EOD Window. The type is
    // interface{} with range: 0..4294967295.
    SodEodDirtyMarkCount interface{}

    // Number of Sessions Invalid Deletes Within SOD EOD Window. The type is
    // interface{} with range: 0..4294967295.
    SodEodDirtyDeleteCount interface{}

    // Number of ACKs sent to Srg. The type is interface{} with range:
    // 0..4294967295.
    AckToSrg interface{}

    // Number of NACKs sent to Srg. The type is interface{} with range:
    // 0..4294967295.
    NackToSrg interface{}

    // Number of NACKs Failed to send to Srg. The type is interface{} with range:
    // 0..4294967295.
    NackToSrgFailCnt interface{}

    // Number of Txlist remove all calls. The type is interface{} with range:
    // 0..4294967295.
    TxlistRemoveAll interface{}

    // Number for Txlist delete for sync msg. The type is interface{} with range:
    // 0..4294967295.
    TxlistDelSync interface{}

    // Number of Txlist delete for sync which are not linked. The type is
    // interface{} with range: 0..4294967295.
    TxlistDelSyncNotlinked interface{}

    // Number of Txlist delete for App msg. The type is interface{} with range:
    // 0..4294967295.
    TxlistDelApp interface{}

    // Number of Txlist delete for App which are not linked. The type is
    // interface{} with range: 0..4294967295.
    TxlistDelAppNotlinked interface{}

    // Number of Txlist Cleanup called on Invalid subscriber srg state. The type
    // is interface{} with range: 0..4294967295.
    TxlistCleanInvalidState interface{}

    // Number of Internal errors upon Master Txlist remove all call. The type is
    // interface{} with range: 0..4294967295.
    TxlistRemoveAllInternalError interface{}

    // Flag indicating SRG Flow control enabled or not. The type is bool.
    IsSrgFlowControlEnabled interface{}

    // Maximum no.of inflight sessions allowed. The type is interface{} with
    // range: 0..4294967295.
    MaxInflightSessoinCount interface{}

    // Threshold Limit to resume the flow control. The type is interface{} with
    // range: 0..4294967295.
    FlowControlResumeThreshold interface{}

    // No.of Sessions inflight at given time. The type is interface{} with range:
    // 0..4294967295.
    InflightSessionCount interface{}

    // No.of inflight sessions added. The type is interface{} with range:
    // 0..4294967295.
    InflightAddCount interface{}

    // Inflight Underrun Counter. The type is interface{} with range:
    // 0..4294967295.
    InflightUnderRunCount interface{}

    // Memory Alloc Failures for Inflight Entry. The type is interface{} with
    // range: 0..4294967295.
    InflightAllocFails interface{}

    // Inflight Entry Insert Failures. The type is interface{} with range:
    // 0..4294967295.
    InflightInsertFailures interface{}

    // Inflight Deletes Count. The type is interface{} with range: 0..4294967295.
    InflightDeletes interface{}

    // Inflight Entries not found during delete. The type is interface{} with
    // range: 0..4294967295.
    InflightNotFound interface{}

    // Inflight Entry Delete Failures. The type is interface{} with range:
    // 0..4294967295.
    InflightDeleteFailures interface{}

    // Total No.of times Pause is Enabled. The type is interface{} with range:
    // 0..4294967295.
    TotalPauseCount interface{}

    // Total No.of times Resume is triggered. The type is interface{} with range:
    // 0..4294967295.
    TotalResumeCount interface{}

    // Total No of times Dont send to Txlist. The type is interface{} with range:
    // 0..4294967295.
    TotalDontSendToTxlist interface{}

    // Total No of times SRG Not Master. The type is interface{} with range:
    // 0..4294967295.
    TotalSrgNotMaster interface{}

    // Total No of times Master EOMS Pending. The type is interface{} with range:
    // 0..4294967295.
    TotalMasterEomsPending interface{}

    // Amount of time paused during last flow control window. The type is
    // interface{} with range: 0..18446744073709551615.
    LastPausePeriod interface{}

    // Total Amount of time paused during all flow control windows. The type is
    // interface{} with range: 0..18446744073709551615.
    TotalPauseTime interface{}

    // Timestamp of recent Pause Event. The type is interface{} with range:
    // 0..18446744073709551615.
    LastPauseTime interface{}

    // Timestamp of recent Resume Event. The type is interface{} with range:
    // 0..18446744073709551615.
    LastResumeTime interface{}
}

func (srg *Subscriber_Manager_Nodes_Node_Statistics_Srg) GetEntityData() *types.CommonEntityData {
    srg.EntityData.YFilter = srg.YFilter
    srg.EntityData.YangName = "srg"
    srg.EntityData.BundleName = "cisco_ios_xr"
    srg.EntityData.ParentYangName = "statistics"
    srg.EntityData.SegmentPath = "srg"
    srg.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    srg.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    srg.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    srg.EntityData.Children = make(map[string]types.YChild)
    srg.EntityData.Leafs = make(map[string]types.YLeaf)
    srg.EntityData.Leafs["txlist-send-triggered"] = types.YLeaf{"TxlistSendTriggered", srg.TxlistSendTriggered}
    srg.EntityData.Leafs["txlist-send-failed"] = types.YLeaf{"TxlistSendFailed", srg.TxlistSendFailed}
    srg.EntityData.Leafs["txlist-send-failed-notactive"] = types.YLeaf{"TxlistSendFailedNotactive", srg.TxlistSendFailedNotactive}
    srg.EntityData.Leafs["actual-txlist-sent"] = types.YLeaf{"ActualTxlistSent", srg.ActualTxlistSent}
    srg.EntityData.Leafs["alreadyin-txlist"] = types.YLeaf{"AlreadyinTxlist", srg.AlreadyinTxlist}
    srg.EntityData.Leafs["txlist-encode"] = types.YLeaf{"TxlistEncode", srg.TxlistEncode}
    srg.EntityData.Leafs["txlist-encode-fail"] = types.YLeaf{"TxlistEncodeFail", srg.TxlistEncodeFail}
    srg.EntityData.Leafs["create-update-encode"] = types.YLeaf{"CreateUpdateEncode", srg.CreateUpdateEncode}
    srg.EntityData.Leafs["delete-encode"] = types.YLeaf{"DeleteEncode", srg.DeleteEncode}
    srg.EntityData.Leafs["create-upd-clean-callback"] = types.YLeaf{"CreateUpdCleanCallback", srg.CreateUpdCleanCallback}
    srg.EntityData.Leafs["delete-clean-callback"] = types.YLeaf{"DeleteCleanCallback", srg.DeleteCleanCallback}
    srg.EntityData.Leafs["slave-recv-entry"] = types.YLeaf{"SlaveRecvEntry", srg.SlaveRecvEntry}
    srg.EntityData.Leafs["slave-decode-fail"] = types.YLeaf{"SlaveDecodeFail", srg.SlaveDecodeFail}
    srg.EntityData.Leafs["slave-create-update"] = types.YLeaf{"SlaveCreateUpdate", srg.SlaveCreateUpdate}
    srg.EntityData.Leafs["slave-delete"] = types.YLeaf{"SlaveDelete", srg.SlaveDelete}
    srg.EntityData.Leafs["srg-context-malloc"] = types.YLeaf{"SrgContextMalloc", srg.SrgContextMalloc}
    srg.EntityData.Leafs["srg-context-free"] = types.YLeaf{"SrgContextFree", srg.SrgContextFree}
    srg.EntityData.Leafs["sod-count"] = types.YLeaf{"SodCount", srg.SodCount}
    srg.EntityData.Leafs["eod-count"] = types.YLeaf{"EodCount", srg.EodCount}
    srg.EntityData.Leafs["sod-eod-replay-req-count"] = types.YLeaf{"SodEodReplayReqCount", srg.SodEodReplayReqCount}
    srg.EntityData.Leafs["sod-eod-dirty-mark-count"] = types.YLeaf{"SodEodDirtyMarkCount", srg.SodEodDirtyMarkCount}
    srg.EntityData.Leafs["sod-eod-dirty-delete-count"] = types.YLeaf{"SodEodDirtyDeleteCount", srg.SodEodDirtyDeleteCount}
    srg.EntityData.Leafs["ack-to-srg"] = types.YLeaf{"AckToSrg", srg.AckToSrg}
    srg.EntityData.Leafs["nack-to-srg"] = types.YLeaf{"NackToSrg", srg.NackToSrg}
    srg.EntityData.Leafs["nack-to-srg-fail-cnt"] = types.YLeaf{"NackToSrgFailCnt", srg.NackToSrgFailCnt}
    srg.EntityData.Leafs["txlist-remove-all"] = types.YLeaf{"TxlistRemoveAll", srg.TxlistRemoveAll}
    srg.EntityData.Leafs["txlist-del-sync"] = types.YLeaf{"TxlistDelSync", srg.TxlistDelSync}
    srg.EntityData.Leafs["txlist-del-sync-notlinked"] = types.YLeaf{"TxlistDelSyncNotlinked", srg.TxlistDelSyncNotlinked}
    srg.EntityData.Leafs["txlist-del-app"] = types.YLeaf{"TxlistDelApp", srg.TxlistDelApp}
    srg.EntityData.Leafs["txlist-del-app-notlinked"] = types.YLeaf{"TxlistDelAppNotlinked", srg.TxlistDelAppNotlinked}
    srg.EntityData.Leafs["txlist-clean-invalid-state"] = types.YLeaf{"TxlistCleanInvalidState", srg.TxlistCleanInvalidState}
    srg.EntityData.Leafs["txlist-remove-all-internal-error"] = types.YLeaf{"TxlistRemoveAllInternalError", srg.TxlistRemoveAllInternalError}
    srg.EntityData.Leafs["is-srg-flow-control-enabled"] = types.YLeaf{"IsSrgFlowControlEnabled", srg.IsSrgFlowControlEnabled}
    srg.EntityData.Leafs["max-inflight-sessoin-count"] = types.YLeaf{"MaxInflightSessoinCount", srg.MaxInflightSessoinCount}
    srg.EntityData.Leafs["flow-control-resume-threshold"] = types.YLeaf{"FlowControlResumeThreshold", srg.FlowControlResumeThreshold}
    srg.EntityData.Leafs["inflight-session-count"] = types.YLeaf{"InflightSessionCount", srg.InflightSessionCount}
    srg.EntityData.Leafs["inflight-add-count"] = types.YLeaf{"InflightAddCount", srg.InflightAddCount}
    srg.EntityData.Leafs["inflight-under-run-count"] = types.YLeaf{"InflightUnderRunCount", srg.InflightUnderRunCount}
    srg.EntityData.Leafs["inflight-alloc-fails"] = types.YLeaf{"InflightAllocFails", srg.InflightAllocFails}
    srg.EntityData.Leafs["inflight-insert-failures"] = types.YLeaf{"InflightInsertFailures", srg.InflightInsertFailures}
    srg.EntityData.Leafs["inflight-deletes"] = types.YLeaf{"InflightDeletes", srg.InflightDeletes}
    srg.EntityData.Leafs["inflight-not-found"] = types.YLeaf{"InflightNotFound", srg.InflightNotFound}
    srg.EntityData.Leafs["inflight-delete-failures"] = types.YLeaf{"InflightDeleteFailures", srg.InflightDeleteFailures}
    srg.EntityData.Leafs["total-pause-count"] = types.YLeaf{"TotalPauseCount", srg.TotalPauseCount}
    srg.EntityData.Leafs["total-resume-count"] = types.YLeaf{"TotalResumeCount", srg.TotalResumeCount}
    srg.EntityData.Leafs["total-dont-send-to-txlist"] = types.YLeaf{"TotalDontSendToTxlist", srg.TotalDontSendToTxlist}
    srg.EntityData.Leafs["total-srg-not-master"] = types.YLeaf{"TotalSrgNotMaster", srg.TotalSrgNotMaster}
    srg.EntityData.Leafs["total-master-eoms-pending"] = types.YLeaf{"TotalMasterEomsPending", srg.TotalMasterEomsPending}
    srg.EntityData.Leafs["last-pause-period"] = types.YLeaf{"LastPausePeriod", srg.LastPausePeriod}
    srg.EntityData.Leafs["total-pause-time"] = types.YLeaf{"TotalPauseTime", srg.TotalPauseTime}
    srg.EntityData.Leafs["last-pause-time"] = types.YLeaf{"LastPauseTime", srg.LastPauseTime}
    srg.EntityData.Leafs["last-resume-time"] = types.YLeaf{"LastResumeTime", srg.LastResumeTime}
    return &(srg.EntityData)
}

// Subscriber_Session
// Subscriber session operational data
type Subscriber_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of subscriber session supported nodes.
    Nodes Subscriber_Session_Nodes
}

func (session *Subscriber_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "subscriber"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["nodes"] = types.YChild{"Nodes", &session.Nodes}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(session.EntityData)
}

// Subscriber_Session_Nodes
// List of subscriber session supported nodes
type Subscriber_Session_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber session operational data for a particular node. The type is
    // slice of Subscriber_Session_Nodes_Node.
    Node []Subscriber_Session_Nodes_Node
}

func (nodes *Subscriber_Session_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "session"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Subscriber_Session_Nodes_Node
// Subscriber session operational data for a
// particular node
type Subscriber_Session_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node name. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    NodeName interface{}

    // Summary information filtered by authorization state.
    AuthorSummaries Subscriber_Session_Nodes_Node_AuthorSummaries

    // Subscriber session summary information.
    Summary Subscriber_Session_Nodes_Node_Summary

    // Summary information filtered by MAC address.
    MacSummaries Subscriber_Session_Nodes_Node_MacSummaries

    // Summary information filtered by interface.
    InterfaceSummaries Subscriber_Session_Nodes_Node_InterfaceSummaries

    // Summary information filtered by authentication state.
    AuthenticationSummaries Subscriber_Session_Nodes_Node_AuthenticationSummaries

    // Summary information filtered by session state.
    StateSummaries Subscriber_Session_Nodes_Node_StateSummaries

    // Summary information filtered by IPv4 address and VRF.
    Ipv4AddressVrfSummaries Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries

    // Summary information filtered by address family.
    AddressFamilySummaries Subscriber_Session_Nodes_Node_AddressFamilySummaries

    // Summary information filtered by username.
    UsernameSummaries Subscriber_Session_Nodes_Node_UsernameSummaries

    // Summary information filtered by access interface.
    AccessInterfaceSummaries Subscriber_Session_Nodes_Node_AccessInterfaceSummaries

    // Summary information filtered by subscriber IPv4 address.
    Ipv4AddressSummaries Subscriber_Session_Nodes_Node_Ipv4AddressSummaries

    // Summary information filtered by VRF.
    VrfSummaries Subscriber_Session_Nodes_Node_VrfSummaries

    // IP subscriber sessions.
    Sessions Subscriber_Session_Nodes_Node_Sessions
}

func (node *Subscriber_Session_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["author-summaries"] = types.YChild{"AuthorSummaries", &node.AuthorSummaries}
    node.EntityData.Children["summary"] = types.YChild{"Summary", &node.Summary}
    node.EntityData.Children["mac-summaries"] = types.YChild{"MacSummaries", &node.MacSummaries}
    node.EntityData.Children["interface-summaries"] = types.YChild{"InterfaceSummaries", &node.InterfaceSummaries}
    node.EntityData.Children["authentication-summaries"] = types.YChild{"AuthenticationSummaries", &node.AuthenticationSummaries}
    node.EntityData.Children["state-summaries"] = types.YChild{"StateSummaries", &node.StateSummaries}
    node.EntityData.Children["ipv4-address-vrf-summaries"] = types.YChild{"Ipv4AddressVrfSummaries", &node.Ipv4AddressVrfSummaries}
    node.EntityData.Children["address-family-summaries"] = types.YChild{"AddressFamilySummaries", &node.AddressFamilySummaries}
    node.EntityData.Children["username-summaries"] = types.YChild{"UsernameSummaries", &node.UsernameSummaries}
    node.EntityData.Children["access-interface-summaries"] = types.YChild{"AccessInterfaceSummaries", &node.AccessInterfaceSummaries}
    node.EntityData.Children["ipv4-address-summaries"] = types.YChild{"Ipv4AddressSummaries", &node.Ipv4AddressSummaries}
    node.EntityData.Children["vrf-summaries"] = types.YChild{"VrfSummaries", &node.VrfSummaries}
    node.EntityData.Children["sessions"] = types.YChild{"Sessions", &node.Sessions}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    return &(node.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries
// Summary information filtered by authorization
// state
type Subscriber_Session_Nodes_Node_AuthorSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // authorization summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary.
    AuthorSummary []Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary
}

func (authorSummaries *Subscriber_Session_Nodes_Node_AuthorSummaries) GetEntityData() *types.CommonEntityData {
    authorSummaries.EntityData.YFilter = authorSummaries.YFilter
    authorSummaries.EntityData.YangName = "author-summaries"
    authorSummaries.EntityData.BundleName = "cisco_ios_xr"
    authorSummaries.EntityData.ParentYangName = "node"
    authorSummaries.EntityData.SegmentPath = "author-summaries"
    authorSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorSummaries.EntityData.Children = make(map[string]types.YChild)
    authorSummaries.EntityData.Children["author-summary"] = types.YChild{"AuthorSummary", nil}
    for i := range authorSummaries.AuthorSummary {
        authorSummaries.EntityData.Children[types.GetSegmentPath(&authorSummaries.AuthorSummary[i])] = types.YChild{"AuthorSummary", &authorSummaries.AuthorSummary[i]}
    }
    authorSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authorSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary
// authorization summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Authorization state. The type is
    // SubscriberAuthorStateFilterFlag.
    AuthorState interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr
}

func (authorSummary *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary) GetEntityData() *types.CommonEntityData {
    authorSummary.EntityData.YFilter = authorSummary.YFilter
    authorSummary.EntityData.YangName = "author-summary"
    authorSummary.EntityData.BundleName = "cisco_ios_xr"
    authorSummary.EntityData.ParentYangName = "author-summaries"
    authorSummary.EntityData.SegmentPath = "author-summary" + "[author-state='" + fmt.Sprintf("%v", authorSummary.AuthorState) + "']"
    authorSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authorSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authorSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authorSummary.EntityData.Children = make(map[string]types.YChild)
    authorSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &authorSummary.StateXr}
    authorSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &authorSummary.AddressFamilyXr}
    authorSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    authorSummary.EntityData.Leafs["author-state"] = types.YLeaf{"AuthorState", authorSummary.AuthorState}
    return &(authorSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "author-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "author-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthorSummaries_AuthorSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary
// Subscriber session summary information
type Subscriber_Session_Nodes_Node_Summary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_Summary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr
}

func (summary *Subscriber_Session_Nodes_Node_Summary) GetEntityData() *types.CommonEntityData {
    summary.EntityData.YFilter = summary.YFilter
    summary.EntityData.YangName = "summary"
    summary.EntityData.BundleName = "cisco_ios_xr"
    summary.EntityData.ParentYangName = "node"
    summary.EntityData.SegmentPath = "summary"
    summary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summary.EntityData.Children = make(map[string]types.YChild)
    summary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &summary.StateXr}
    summary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &summary.AddressFamilyXr}
    summary.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(summary.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_Summary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_Summary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Summary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Summary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Summary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries
// Summary information filtered by MAC address
type Subscriber_Session_Nodes_Node_MacSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // MAC address summary. The type is slice of
    // Subscriber_Session_Nodes_Node_MacSummaries_MacSummary.
    MacSummary []Subscriber_Session_Nodes_Node_MacSummaries_MacSummary
}

func (macSummaries *Subscriber_Session_Nodes_Node_MacSummaries) GetEntityData() *types.CommonEntityData {
    macSummaries.EntityData.YFilter = macSummaries.YFilter
    macSummaries.EntityData.YangName = "mac-summaries"
    macSummaries.EntityData.BundleName = "cisco_ios_xr"
    macSummaries.EntityData.ParentYangName = "node"
    macSummaries.EntityData.SegmentPath = "mac-summaries"
    macSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSummaries.EntityData.Children = make(map[string]types.YChild)
    macSummaries.EntityData.Children["mac-summary"] = types.YChild{"MacSummary", nil}
    for i := range macSummaries.MacSummary {
        macSummaries.EntityData.Children[types.GetSegmentPath(&macSummaries.MacSummary[i])] = types.YChild{"MacSummary", &macSummaries.MacSummary[i]}
    }
    macSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(macSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary
// MAC address summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subscriber MAC address. The type is string with
    // pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr
}

func (macSummary *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary) GetEntityData() *types.CommonEntityData {
    macSummary.EntityData.YFilter = macSummary.YFilter
    macSummary.EntityData.YangName = "mac-summary"
    macSummary.EntityData.BundleName = "cisco_ios_xr"
    macSummary.EntityData.ParentYangName = "mac-summaries"
    macSummary.EntityData.SegmentPath = "mac-summary" + "[mac-address='" + fmt.Sprintf("%v", macSummary.MacAddress) + "']"
    macSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macSummary.EntityData.Children = make(map[string]types.YChild)
    macSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &macSummary.StateXr}
    macSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &macSummary.AddressFamilyXr}
    macSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    macSummary.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", macSummary.MacAddress}
    return &(macSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "mac-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "mac-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_MacSummaries_MacSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries
// Summary information filtered by interface
type Subscriber_Session_Nodes_Node_InterfaceSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface summary. The type is slice of
    // Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary.
    InterfaceSummary []Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary
}

func (interfaceSummaries *Subscriber_Session_Nodes_Node_InterfaceSummaries) GetEntityData() *types.CommonEntityData {
    interfaceSummaries.EntityData.YFilter = interfaceSummaries.YFilter
    interfaceSummaries.EntityData.YangName = "interface-summaries"
    interfaceSummaries.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummaries.EntityData.ParentYangName = "node"
    interfaceSummaries.EntityData.SegmentPath = "interface-summaries"
    interfaceSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummaries.EntityData.Children = make(map[string]types.YChild)
    interfaceSummaries.EntityData.Children["interface-summary"] = types.YChild{"InterfaceSummary", nil}
    for i := range interfaceSummaries.InterfaceSummary {
        interfaceSummaries.EntityData.Children[types.GetSegmentPath(&interfaceSummaries.InterfaceSummary[i])] = types.YChild{"InterfaceSummary", &interfaceSummaries.InterfaceSummary[i]}
    }
    interfaceSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaceSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary
// Interface summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr
}

func (interfaceSummary *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary) GetEntityData() *types.CommonEntityData {
    interfaceSummary.EntityData.YFilter = interfaceSummary.YFilter
    interfaceSummary.EntityData.YangName = "interface-summary"
    interfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    interfaceSummary.EntityData.ParentYangName = "interface-summaries"
    interfaceSummary.EntityData.SegmentPath = "interface-summary" + "[interface-name='" + fmt.Sprintf("%v", interfaceSummary.InterfaceName) + "']"
    interfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceSummary.EntityData.Children = make(map[string]types.YChild)
    interfaceSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &interfaceSummary.StateXr}
    interfaceSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &interfaceSummary.AddressFamilyXr}
    interfaceSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceSummary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", interfaceSummary.InterfaceName}
    return &(interfaceSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "interface-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "interface-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_InterfaceSummaries_InterfaceSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries
// Summary information filtered by
// authentication state
type Subscriber_Session_Nodes_Node_AuthenticationSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // authentication summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary.
    AuthenticationSummary []Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary
}

func (authenticationSummaries *Subscriber_Session_Nodes_Node_AuthenticationSummaries) GetEntityData() *types.CommonEntityData {
    authenticationSummaries.EntityData.YFilter = authenticationSummaries.YFilter
    authenticationSummaries.EntityData.YangName = "authentication-summaries"
    authenticationSummaries.EntityData.BundleName = "cisco_ios_xr"
    authenticationSummaries.EntityData.ParentYangName = "node"
    authenticationSummaries.EntityData.SegmentPath = "authentication-summaries"
    authenticationSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationSummaries.EntityData.Children = make(map[string]types.YChild)
    authenticationSummaries.EntityData.Children["authentication-summary"] = types.YChild{"AuthenticationSummary", nil}
    for i := range authenticationSummaries.AuthenticationSummary {
        authenticationSummaries.EntityData.Children[types.GetSegmentPath(&authenticationSummaries.AuthenticationSummary[i])] = types.YChild{"AuthenticationSummary", &authenticationSummaries.AuthenticationSummary[i]}
    }
    authenticationSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(authenticationSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary
// authentication summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Authentication state. The type is
    // SubscriberAuthenStateFilterFlag.
    AuthenticationState interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr
}

func (authenticationSummary *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary) GetEntityData() *types.CommonEntityData {
    authenticationSummary.EntityData.YFilter = authenticationSummary.YFilter
    authenticationSummary.EntityData.YangName = "authentication-summary"
    authenticationSummary.EntityData.BundleName = "cisco_ios_xr"
    authenticationSummary.EntityData.ParentYangName = "authentication-summaries"
    authenticationSummary.EntityData.SegmentPath = "authentication-summary" + "[authentication-state='" + fmt.Sprintf("%v", authenticationSummary.AuthenticationState) + "']"
    authenticationSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authenticationSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authenticationSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authenticationSummary.EntityData.Children = make(map[string]types.YChild)
    authenticationSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &authenticationSummary.StateXr}
    authenticationSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &authenticationSummary.AddressFamilyXr}
    authenticationSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    authenticationSummary.EntityData.Leafs["authentication-state"] = types.YLeaf{"AuthenticationState", authenticationSummary.AuthenticationState}
    return &(authenticationSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "authentication-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "authentication-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AuthenticationSummaries_AuthenticationSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries
// Summary information filtered by session state
type Subscriber_Session_Nodes_Node_StateSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State summary. The type is slice of
    // Subscriber_Session_Nodes_Node_StateSummaries_StateSummary.
    StateSummary []Subscriber_Session_Nodes_Node_StateSummaries_StateSummary
}

func (stateSummaries *Subscriber_Session_Nodes_Node_StateSummaries) GetEntityData() *types.CommonEntityData {
    stateSummaries.EntityData.YFilter = stateSummaries.YFilter
    stateSummaries.EntityData.YangName = "state-summaries"
    stateSummaries.EntityData.BundleName = "cisco_ios_xr"
    stateSummaries.EntityData.ParentYangName = "node"
    stateSummaries.EntityData.SegmentPath = "state-summaries"
    stateSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSummaries.EntityData.Children = make(map[string]types.YChild)
    stateSummaries.EntityData.Children["state-summary"] = types.YChild{"StateSummary", nil}
    for i := range stateSummaries.StateSummary {
        stateSummaries.EntityData.Children[types.GetSegmentPath(&stateSummaries.StateSummary[i])] = types.YChild{"StateSummary", &stateSummaries.StateSummary[i]}
    }
    stateSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary
// State summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subscriber state. The type is
    // SubscriberStateFilterFlag.
    State interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr
}

func (stateSummary *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary) GetEntityData() *types.CommonEntityData {
    stateSummary.EntityData.YFilter = stateSummary.YFilter
    stateSummary.EntityData.YangName = "state-summary"
    stateSummary.EntityData.BundleName = "cisco_ios_xr"
    stateSummary.EntityData.ParentYangName = "state-summaries"
    stateSummary.EntityData.SegmentPath = "state-summary" + "[state='" + fmt.Sprintf("%v", stateSummary.State) + "']"
    stateSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateSummary.EntityData.Children = make(map[string]types.YChild)
    stateSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &stateSummary.StateXr}
    stateSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &stateSummary.AddressFamilyXr}
    stateSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    stateSummary.EntityData.Leafs["state"] = types.YLeaf{"State", stateSummary.State}
    return &(stateSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "state-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "state-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_StateSummaries_StateSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries
// Summary information filtered by IPv4 address
// and VRF
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address and VRF summary. The type is slice of
    // Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary.
    Ipv4AddressVrfSummary []Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary
}

func (ipv4AddressVrfSummaries *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries) GetEntityData() *types.CommonEntityData {
    ipv4AddressVrfSummaries.EntityData.YFilter = ipv4AddressVrfSummaries.YFilter
    ipv4AddressVrfSummaries.EntityData.YangName = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressVrfSummaries.EntityData.ParentYangName = "node"
    ipv4AddressVrfSummaries.EntityData.SegmentPath = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressVrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressVrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressVrfSummaries.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressVrfSummaries.EntityData.Children["ipv4-address-vrf-summary"] = types.YChild{"Ipv4AddressVrfSummary", nil}
    for i := range ipv4AddressVrfSummaries.Ipv4AddressVrfSummary {
        ipv4AddressVrfSummaries.EntityData.Children[types.GetSegmentPath(&ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i])] = types.YChild{"Ipv4AddressVrfSummary", &ipv4AddressVrfSummaries.Ipv4AddressVrfSummary[i]}
    }
    ipv4AddressVrfSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4AddressVrfSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary
// IPv4 address and VRF summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Subscriber IPv4 address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr
}

func (ipv4AddressVrfSummary *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary) GetEntityData() *types.CommonEntityData {
    ipv4AddressVrfSummary.EntityData.YFilter = ipv4AddressVrfSummary.YFilter
    ipv4AddressVrfSummary.EntityData.YangName = "ipv4-address-vrf-summary"
    ipv4AddressVrfSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressVrfSummary.EntityData.ParentYangName = "ipv4-address-vrf-summaries"
    ipv4AddressVrfSummary.EntityData.SegmentPath = "ipv4-address-vrf-summary"
    ipv4AddressVrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressVrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressVrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressVrfSummary.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressVrfSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &ipv4AddressVrfSummary.StateXr}
    ipv4AddressVrfSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &ipv4AddressVrfSummary.AddressFamilyXr}
    ipv4AddressVrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressVrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", ipv4AddressVrfSummary.VrfName}
    ipv4AddressVrfSummary.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4AddressVrfSummary.Address}
    return &(ipv4AddressVrfSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "ipv4-address-vrf-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "ipv4-address-vrf-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressVrfSummaries_Ipv4AddressVrfSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries
// Summary information filtered by address
// family
type Subscriber_Session_Nodes_Node_AddressFamilySummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Address family summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary.
    AddressFamilySummary []Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary
}

func (addressFamilySummaries *Subscriber_Session_Nodes_Node_AddressFamilySummaries) GetEntityData() *types.CommonEntityData {
    addressFamilySummaries.EntityData.YFilter = addressFamilySummaries.YFilter
    addressFamilySummaries.EntityData.YangName = "address-family-summaries"
    addressFamilySummaries.EntityData.BundleName = "cisco_ios_xr"
    addressFamilySummaries.EntityData.ParentYangName = "node"
    addressFamilySummaries.EntityData.SegmentPath = "address-family-summaries"
    addressFamilySummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilySummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilySummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilySummaries.EntityData.Children = make(map[string]types.YChild)
    addressFamilySummaries.EntityData.Children["address-family-summary"] = types.YChild{"AddressFamilySummary", nil}
    for i := range addressFamilySummaries.AddressFamilySummary {
        addressFamilySummaries.EntityData.Children[types.GetSegmentPath(&addressFamilySummaries.AddressFamilySummary[i])] = types.YChild{"AddressFamilySummary", &addressFamilySummaries.AddressFamilySummary[i]}
    }
    addressFamilySummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilySummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary
// Address family summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Address family. The type is
    // SubscriberAddressFamilyFilterFlag.
    AddressFamily interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr
}

func (addressFamilySummary *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary) GetEntityData() *types.CommonEntityData {
    addressFamilySummary.EntityData.YFilter = addressFamilySummary.YFilter
    addressFamilySummary.EntityData.YangName = "address-family-summary"
    addressFamilySummary.EntityData.BundleName = "cisco_ios_xr"
    addressFamilySummary.EntityData.ParentYangName = "address-family-summaries"
    addressFamilySummary.EntityData.SegmentPath = "address-family-summary" + "[address-family='" + fmt.Sprintf("%v", addressFamilySummary.AddressFamily) + "']"
    addressFamilySummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilySummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilySummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilySummary.EntityData.Children = make(map[string]types.YChild)
    addressFamilySummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &addressFamilySummary.StateXr}
    addressFamilySummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &addressFamilySummary.AddressFamilyXr}
    addressFamilySummary.EntityData.Leafs = make(map[string]types.YLeaf)
    addressFamilySummary.EntityData.Leafs["address-family"] = types.YLeaf{"AddressFamily", addressFamilySummary.AddressFamily}
    return &(addressFamilySummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "address-family-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "address-family-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AddressFamilySummaries_AddressFamilySummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries
// Summary information filtered by username
type Subscriber_Session_Nodes_Node_UsernameSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Username summary. The type is slice of
    // Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary.
    UsernameSummary []Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary
}

func (usernameSummaries *Subscriber_Session_Nodes_Node_UsernameSummaries) GetEntityData() *types.CommonEntityData {
    usernameSummaries.EntityData.YFilter = usernameSummaries.YFilter
    usernameSummaries.EntityData.YangName = "username-summaries"
    usernameSummaries.EntityData.BundleName = "cisco_ios_xr"
    usernameSummaries.EntityData.ParentYangName = "node"
    usernameSummaries.EntityData.SegmentPath = "username-summaries"
    usernameSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernameSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernameSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernameSummaries.EntityData.Children = make(map[string]types.YChild)
    usernameSummaries.EntityData.Children["username-summary"] = types.YChild{"UsernameSummary", nil}
    for i := range usernameSummaries.UsernameSummary {
        usernameSummaries.EntityData.Children[types.GetSegmentPath(&usernameSummaries.UsernameSummary[i])] = types.YChild{"UsernameSummary", &usernameSummaries.UsernameSummary[i]}
    }
    usernameSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(usernameSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary
// Username summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subscriber username. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Username interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr
}

func (usernameSummary *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary) GetEntityData() *types.CommonEntityData {
    usernameSummary.EntityData.YFilter = usernameSummary.YFilter
    usernameSummary.EntityData.YangName = "username-summary"
    usernameSummary.EntityData.BundleName = "cisco_ios_xr"
    usernameSummary.EntityData.ParentYangName = "username-summaries"
    usernameSummary.EntityData.SegmentPath = "username-summary" + "[username='" + fmt.Sprintf("%v", usernameSummary.Username) + "']"
    usernameSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    usernameSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    usernameSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    usernameSummary.EntityData.Children = make(map[string]types.YChild)
    usernameSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &usernameSummary.StateXr}
    usernameSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &usernameSummary.AddressFamilyXr}
    usernameSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    usernameSummary.EntityData.Leafs["username"] = types.YLeaf{"Username", usernameSummary.Username}
    return &(usernameSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "username-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "username-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_UsernameSummaries_UsernameSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries
// Summary information filtered by access
// interface
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface summary. The type is slice of
    // Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary.
    AccessInterfaceSummary []Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary
}

func (accessInterfaceSummaries *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummaries.EntityData.YFilter = accessInterfaceSummaries.YFilter
    accessInterfaceSummaries.EntityData.YangName = "access-interface-summaries"
    accessInterfaceSummaries.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummaries.EntityData.ParentYangName = "node"
    accessInterfaceSummaries.EntityData.SegmentPath = "access-interface-summaries"
    accessInterfaceSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummaries.EntityData.Children = make(map[string]types.YChild)
    accessInterfaceSummaries.EntityData.Children["access-interface-summary"] = types.YChild{"AccessInterfaceSummary", nil}
    for i := range accessInterfaceSummaries.AccessInterfaceSummary {
        accessInterfaceSummaries.EntityData.Children[types.GetSegmentPath(&accessInterfaceSummaries.AccessInterfaceSummary[i])] = types.YChild{"AccessInterfaceSummary", &accessInterfaceSummaries.AccessInterfaceSummary[i]}
    }
    accessInterfaceSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessInterfaceSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary
// Access interface summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr
}

func (accessInterfaceSummary *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary) GetEntityData() *types.CommonEntityData {
    accessInterfaceSummary.EntityData.YFilter = accessInterfaceSummary.YFilter
    accessInterfaceSummary.EntityData.YangName = "access-interface-summary"
    accessInterfaceSummary.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSummary.EntityData.ParentYangName = "access-interface-summaries"
    accessInterfaceSummary.EntityData.SegmentPath = "access-interface-summary" + "[interface-name='" + fmt.Sprintf("%v", accessInterfaceSummary.InterfaceName) + "']"
    accessInterfaceSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSummary.EntityData.Children = make(map[string]types.YChild)
    accessInterfaceSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &accessInterfaceSummary.StateXr}
    accessInterfaceSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &accessInterfaceSummary.AddressFamilyXr}
    accessInterfaceSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    accessInterfaceSummary.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", accessInterfaceSummary.InterfaceName}
    return &(accessInterfaceSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "access-interface-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "access-interface-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_AccessInterfaceSummaries_AccessInterfaceSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries
// Summary information filtered by subscriber
// IPv4 address
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 address summary. The type is slice of
    // Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary.
    Ipv4AddressSummary []Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary
}

func (ipv4AddressSummaries *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries) GetEntityData() *types.CommonEntityData {
    ipv4AddressSummaries.EntityData.YFilter = ipv4AddressSummaries.YFilter
    ipv4AddressSummaries.EntityData.YangName = "ipv4-address-summaries"
    ipv4AddressSummaries.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressSummaries.EntityData.ParentYangName = "node"
    ipv4AddressSummaries.EntityData.SegmentPath = "ipv4-address-summaries"
    ipv4AddressSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressSummaries.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressSummaries.EntityData.Children["ipv4-address-summary"] = types.YChild{"Ipv4AddressSummary", nil}
    for i := range ipv4AddressSummaries.Ipv4AddressSummary {
        ipv4AddressSummaries.EntityData.Children[types.GetSegmentPath(&ipv4AddressSummaries.Ipv4AddressSummary[i])] = types.YChild{"Ipv4AddressSummary", &ipv4AddressSummaries.Ipv4AddressSummary[i]}
    }
    ipv4AddressSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4AddressSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary
// IPv4 address summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subscriber IPv4 address. The type is string with
    // pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr
}

func (ipv4AddressSummary *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary) GetEntityData() *types.CommonEntityData {
    ipv4AddressSummary.EntityData.YFilter = ipv4AddressSummary.YFilter
    ipv4AddressSummary.EntityData.YangName = "ipv4-address-summary"
    ipv4AddressSummary.EntityData.BundleName = "cisco_ios_xr"
    ipv4AddressSummary.EntityData.ParentYangName = "ipv4-address-summaries"
    ipv4AddressSummary.EntityData.SegmentPath = "ipv4-address-summary" + "[address='" + fmt.Sprintf("%v", ipv4AddressSummary.Address) + "']"
    ipv4AddressSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4AddressSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4AddressSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4AddressSummary.EntityData.Children = make(map[string]types.YChild)
    ipv4AddressSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &ipv4AddressSummary.StateXr}
    ipv4AddressSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &ipv4AddressSummary.AddressFamilyXr}
    ipv4AddressSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4AddressSummary.EntityData.Leafs["address"] = types.YLeaf{"Address", ipv4AddressSummary.Address}
    return &(ipv4AddressSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "ipv4-address-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "ipv4-address-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_Ipv4AddressSummaries_Ipv4AddressSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries
// Summary information filtered by VRF
type Subscriber_Session_Nodes_Node_VrfSummaries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF summary. The type is slice of
    // Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary.
    VrfSummary []Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary
}

func (vrfSummaries *Subscriber_Session_Nodes_Node_VrfSummaries) GetEntityData() *types.CommonEntityData {
    vrfSummaries.EntityData.YFilter = vrfSummaries.YFilter
    vrfSummaries.EntityData.YangName = "vrf-summaries"
    vrfSummaries.EntityData.BundleName = "cisco_ios_xr"
    vrfSummaries.EntityData.ParentYangName = "node"
    vrfSummaries.EntityData.SegmentPath = "vrf-summaries"
    vrfSummaries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummaries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummaries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummaries.EntityData.Children = make(map[string]types.YChild)
    vrfSummaries.EntityData.Children["vrf-summary"] = types.YChild{"VrfSummary", nil}
    for i := range vrfSummaries.VrfSummary {
        vrfSummaries.EntityData.Children[types.GetSegmentPath(&vrfSummaries.VrfSummary[i])] = types.YChild{"VrfSummary", &vrfSummaries.VrfSummary[i]}
    }
    vrfSummaries.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfSummaries.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary
// VRF summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // State summary.
    StateXr Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr

    // Address family summary.
    AddressFamilyXr Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr
}

func (vrfSummary *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary) GetEntityData() *types.CommonEntityData {
    vrfSummary.EntityData.YFilter = vrfSummary.YFilter
    vrfSummary.EntityData.YangName = "vrf-summary"
    vrfSummary.EntityData.BundleName = "cisco_ios_xr"
    vrfSummary.EntityData.ParentYangName = "vrf-summaries"
    vrfSummary.EntityData.SegmentPath = "vrf-summary" + "[vrf-name='" + fmt.Sprintf("%v", vrfSummary.VrfName) + "']"
    vrfSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfSummary.EntityData.Children = make(map[string]types.YChild)
    vrfSummary.EntityData.Children["state-xr"] = types.YChild{"StateXr", &vrfSummary.StateXr}
    vrfSummary.EntityData.Children["address-family-xr"] = types.YChild{"AddressFamilyXr", &vrfSummary.AddressFamilyXr}
    vrfSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfSummary.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrfSummary.VrfName}
    return &(vrfSummary.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr
// State summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket
}

func (stateXr *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr) GetEntityData() *types.CommonEntityData {
    stateXr.EntityData.YFilter = stateXr.YFilter
    stateXr.EntityData.YangName = "state-xr"
    stateXr.EntityData.BundleName = "cisco_ios_xr"
    stateXr.EntityData.ParentYangName = "vrf-summary"
    stateXr.EntityData.SegmentPath = "state-xr"
    stateXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateXr.EntityData.Children = make(map[string]types.YChild)
    stateXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &stateXr.Pppoe}
    stateXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &stateXr.IpSubscriberDhcp}
    stateXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &stateXr.IpSubscriberPacket}
    stateXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(stateXr.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "state-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", pppoe.InitializedSessions}
    pppoe.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", pppoe.ConnectingSessions}
    pppoe.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", pppoe.ConnectedSessions}
    pppoe.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", pppoe.ActivatedSessions}
    pppoe.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", pppoe.IdleSessions}
    pppoe.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", pppoe.DisconnectingSessions}
    pppoe.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", pppoe.EndSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "state-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberDhcp.InitializedSessions}
    ipSubscriberDhcp.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberDhcp.ConnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberDhcp.ConnectedSessions}
    ipSubscriberDhcp.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberDhcp.ActivatedSessions}
    ipSubscriberDhcp.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberDhcp.IdleSessions}
    ipSubscriberDhcp.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberDhcp.DisconnectingSessions}
    ipSubscriberDhcp.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberDhcp.EndSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions in initialized state. The type is interface{} with range:
    // 0..4294967295.
    InitializedSessions interface{}

    // Sessions in connecting state. The type is interface{} with range:
    // 0..4294967295.
    ConnectingSessions interface{}

    // Sessions in connected state. The type is interface{} with range:
    // 0..4294967295.
    ConnectedSessions interface{}

    // Sessions in activated state. The type is interface{} with range:
    // 0..4294967295.
    ActivatedSessions interface{}

    // Sessions in idle state. The type is interface{} with range: 0..4294967295.
    IdleSessions interface{}

    // Sessions in disconnecting state. The type is interface{} with range:
    // 0..4294967295.
    DisconnectingSessions interface{}

    // Sessions in end state. The type is interface{} with range: 0..4294967295.
    EndSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_StateXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "state-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["initialized-sessions"] = types.YLeaf{"InitializedSessions", ipSubscriberPacket.InitializedSessions}
    ipSubscriberPacket.EntityData.Leafs["connecting-sessions"] = types.YLeaf{"ConnectingSessions", ipSubscriberPacket.ConnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["connected-sessions"] = types.YLeaf{"ConnectedSessions", ipSubscriberPacket.ConnectedSessions}
    ipSubscriberPacket.EntityData.Leafs["activated-sessions"] = types.YLeaf{"ActivatedSessions", ipSubscriberPacket.ActivatedSessions}
    ipSubscriberPacket.EntityData.Leafs["idle-sessions"] = types.YLeaf{"IdleSessions", ipSubscriberPacket.IdleSessions}
    ipSubscriberPacket.EntityData.Leafs["disconnecting-sessions"] = types.YLeaf{"DisconnectingSessions", ipSubscriberPacket.DisconnectingSessions}
    ipSubscriberPacket.EntityData.Leafs["end-sessions"] = types.YLeaf{"EndSessions", ipSubscriberPacket.EndSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr
// Address family summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PPPoE summary.
    Pppoe Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe

    // IP subscriber DHCP summary.
    IpSubscriberDhcp Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp

    // IP subscriber packet summary.
    IpSubscriberPacket Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket
}

func (addressFamilyXr *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr) GetEntityData() *types.CommonEntityData {
    addressFamilyXr.EntityData.YFilter = addressFamilyXr.YFilter
    addressFamilyXr.EntityData.YangName = "address-family-xr"
    addressFamilyXr.EntityData.BundleName = "cisco_ios_xr"
    addressFamilyXr.EntityData.ParentYangName = "vrf-summary"
    addressFamilyXr.EntityData.SegmentPath = "address-family-xr"
    addressFamilyXr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamilyXr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamilyXr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamilyXr.EntityData.Children = make(map[string]types.YChild)
    addressFamilyXr.EntityData.Children["pppoe"] = types.YChild{"Pppoe", &addressFamilyXr.Pppoe}
    addressFamilyXr.EntityData.Children["ip-subscriber-dhcp"] = types.YChild{"IpSubscriberDhcp", &addressFamilyXr.IpSubscriberDhcp}
    addressFamilyXr.EntityData.Children["ip-subscriber-packet"] = types.YChild{"IpSubscriberPacket", &addressFamilyXr.IpSubscriberPacket}
    addressFamilyXr.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(addressFamilyXr.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe
// PPPoE summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (pppoe *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_Pppoe) GetEntityData() *types.CommonEntityData {
    pppoe.EntityData.YFilter = pppoe.YFilter
    pppoe.EntityData.YangName = "pppoe"
    pppoe.EntityData.BundleName = "cisco_ios_xr"
    pppoe.EntityData.ParentYangName = "address-family-xr"
    pppoe.EntityData.SegmentPath = "pppoe"
    pppoe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pppoe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pppoe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pppoe.EntityData.Children = make(map[string]types.YChild)
    pppoe.EntityData.Leafs = make(map[string]types.YLeaf)
    pppoe.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", pppoe.InProgressSessions}
    pppoe.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", pppoe.Ipv4OnlySessions}
    pppoe.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", pppoe.Ipv6OnlySessions}
    pppoe.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", pppoe.DualPartUpSessions}
    pppoe.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", pppoe.DualUpSessions}
    pppoe.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", pppoe.LacSessions}
    return &(pppoe.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp
// IP subscriber DHCP summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberDhcp *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberDhcp) GetEntityData() *types.CommonEntityData {
    ipSubscriberDhcp.EntityData.YFilter = ipSubscriberDhcp.YFilter
    ipSubscriberDhcp.EntityData.YangName = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberDhcp.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberDhcp.EntityData.SegmentPath = "ip-subscriber-dhcp"
    ipSubscriberDhcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberDhcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberDhcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberDhcp.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberDhcp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberDhcp.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberDhcp.InProgressSessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberDhcp.Ipv4OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberDhcp.Ipv6OnlySessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberDhcp.DualPartUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberDhcp.DualUpSessions}
    ipSubscriberDhcp.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberDhcp.LacSessions}
    return &(ipSubscriberDhcp.EntityData)
}

// Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket
// IP subscriber packet summary
type Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Sessions with undecided address family. The type is interface{} with range:
    // 0..4294967295.
    InProgressSessions interface{}

    // IPv4 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv4OnlySessions interface{}

    // IPv6 only sessions . The type is interface{} with range: 0..4294967295.
    Ipv6OnlySessions interface{}

    // Dual stack partially up sessions. The type is interface{} with range:
    // 0..4294967295.
    DualPartUpSessions interface{}

    // Dual stack up sessions. The type is interface{} with range: 0..4294967295.
    DualUpSessions interface{}

    // LAC sessions. The type is interface{} with range: 0..4294967295.
    LacSessions interface{}
}

func (ipSubscriberPacket *Subscriber_Session_Nodes_Node_VrfSummaries_VrfSummary_AddressFamilyXr_IpSubscriberPacket) GetEntityData() *types.CommonEntityData {
    ipSubscriberPacket.EntityData.YFilter = ipSubscriberPacket.YFilter
    ipSubscriberPacket.EntityData.YangName = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.BundleName = "cisco_ios_xr"
    ipSubscriberPacket.EntityData.ParentYangName = "address-family-xr"
    ipSubscriberPacket.EntityData.SegmentPath = "ip-subscriber-packet"
    ipSubscriberPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSubscriberPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSubscriberPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSubscriberPacket.EntityData.Children = make(map[string]types.YChild)
    ipSubscriberPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    ipSubscriberPacket.EntityData.Leafs["in-progress-sessions"] = types.YLeaf{"InProgressSessions", ipSubscriberPacket.InProgressSessions}
    ipSubscriberPacket.EntityData.Leafs["ipv4-only-sessions"] = types.YLeaf{"Ipv4OnlySessions", ipSubscriberPacket.Ipv4OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["ipv6-only-sessions"] = types.YLeaf{"Ipv6OnlySessions", ipSubscriberPacket.Ipv6OnlySessions}
    ipSubscriberPacket.EntityData.Leafs["dual-part-up-sessions"] = types.YLeaf{"DualPartUpSessions", ipSubscriberPacket.DualPartUpSessions}
    ipSubscriberPacket.EntityData.Leafs["dual-up-sessions"] = types.YLeaf{"DualUpSessions", ipSubscriberPacket.DualUpSessions}
    ipSubscriberPacket.EntityData.Leafs["lac-sessions"] = types.YLeaf{"LacSessions", ipSubscriberPacket.LacSessions}
    return &(ipSubscriberPacket.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions
// IP subscriber sessions
type Subscriber_Session_Nodes_Node_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber session information. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session.
    Session []Subscriber_Session_Nodes_Node_Sessions_Session_
}

func (sessions *Subscriber_Session_Nodes_Node_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "node"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = make(map[string]types.YChild)
    sessions.EntityData.Children["session"] = types.YChild{"Session", nil}
    for i := range sessions.Session {
        sessions.EntityData.Children[types.GetSegmentPath(&sessions.Session[i])] = types.YChild{"Session", &sessions.Session[i]}
    }
    sessions.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(sessions.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session_
// Subscriber session information
type Subscriber_Session_Nodes_Node_Sessions_Session_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session ID. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    SessionId interface{}

    // Subscriber session type. The type is IedgeOperSession.
    SessionType interface{}

    // PPPoE sub type. The type is IedgePppSub.
    PppoeSubType interface{}

    // Interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // Circuit ID. The type is string.
    CircuitId interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // PPPoE LNS address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LnsAddress interface{}

    // PPPoE LAC address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LacAddress interface{}

    // PPPoE LAC tunnel client authentication ID. The type is string.
    TunnelClientAuthenticationId interface{}

    // PPPoE LAC tunnel server authentication ID. The type is string.
    TunnelServerAuthenticationId interface{}

    // Session ip address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SessionIpAddress interface{}

    // Session IPv6 address. The type is string.
    SessionIpv6Address interface{}

    // Session IPv6 prefix. The type is string.
    SessionIpv6Prefix interface{}

    // Session delegated IPv6 prefix. The type is string.
    DelegatedIpv6Prefix interface{}

    // IPv6 Interface ID. The type is string with pattern:
    // b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    Ipv6InterfaceId interface{}

    // MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Accounting session ID. The type is string.
    AccountSessionId interface{}

    // NAS port. The type is string.
    NasPort interface{}

    // Username. The type is string.
    Username interface{}

    // Client Username. The type is string.
    Clientname interface{}

    // Formatted Username. The type is string.
    Formattedname interface{}

    // If true, session is authentic. The type is bool.
    IsSessionAuthentic interface{}

    // If true, session is authorized. The type is bool.
    IsSessionAuthor interface{}

    // Session state. The type is IedgeOperSessionState.
    SessionState interface{}

    // Session creation time in DDD MMM DD HH:MM:SS YYYY format eg: Tue Apr 11
    // 21:30:47 2011. The type is string.
    SessionCreationTime interface{}

    // Time when idle state change occurred in DDD MMM DD HH:MM:SS YYYY format eg:
    // Tue Apr 11 21:30:47 2011. The type is string.
    IdleStateChangeTime interface{}

    // Total session idle time (in seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    TotalSessionIdleTime interface{}

    // Access interface name associated with the session. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    AccessInterfaceName interface{}

    // Active pending callbacks bitmask. The type is interface{} with range:
    // 0..18446744073709551615.
    PendingCallbacks interface{}

    // AF status per Subscriber Session. The type is interface{} with range:
    // 0..4294967295.
    AfUpStatus interface{}

    // Session IPv4 state. The type is IedgeOperSessionAfState.
    SessionIpv4State interface{}

    // Session IPv6 state. The type is IedgeOperSessionAfState.
    SessionIpv6State interface{}

    // Accounting information.
    Accounting Subscriber_Session_Nodes_Node_Sessions_Session__Accounting

    // List of user profile attributes collected for subscriber session.
    UserProfileAttributes Subscriber_Session_Nodes_Node_Sessions_Session__UserProfileAttributes

    // List of mobility attributes collected for subscriber session.
    MobilityAttributes Subscriber_Session_Nodes_Node_Sessions_Session__MobilityAttributes

    // Subscriber change of authorization information. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session__SessionChangeOfAuthorization.
    SessionChangeOfAuthorization []Subscriber_Session_Nodes_Node_Sessions_Session__SessionChangeOfAuthorization
}

func (session_ *Subscriber_Session_Nodes_Node_Sessions_Session_) GetEntityData() *types.CommonEntityData {
    session_.EntityData.YFilter = session_.YFilter
    session_.EntityData.YangName = "session"
    session_.EntityData.BundleName = "cisco_ios_xr"
    session_.EntityData.ParentYangName = "sessions"
    session_.EntityData.SegmentPath = "session" + "[session-id='" + fmt.Sprintf("%v", session_.SessionId) + "']"
    session_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session_.EntityData.Children = make(map[string]types.YChild)
    session_.EntityData.Children["accounting"] = types.YChild{"Accounting", &session_.Accounting}
    session_.EntityData.Children["user-profile-attributes"] = types.YChild{"UserProfileAttributes", &session_.UserProfileAttributes}
    session_.EntityData.Children["mobility-attributes"] = types.YChild{"MobilityAttributes", &session_.MobilityAttributes}
    session_.EntityData.Children["session-change-of-authorization"] = types.YChild{"SessionChangeOfAuthorization", nil}
    for i := range session_.SessionChangeOfAuthorization {
        session_.EntityData.Children[types.GetSegmentPath(&session_.SessionChangeOfAuthorization[i])] = types.YChild{"SessionChangeOfAuthorization", &session_.SessionChangeOfAuthorization[i]}
    }
    session_.EntityData.Leafs = make(map[string]types.YLeaf)
    session_.EntityData.Leafs["session-id"] = types.YLeaf{"SessionId", session_.SessionId}
    session_.EntityData.Leafs["session-type"] = types.YLeaf{"SessionType", session_.SessionType}
    session_.EntityData.Leafs["pppoe-sub-type"] = types.YLeaf{"PppoeSubType", session_.PppoeSubType}
    session_.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", session_.InterfaceName}
    session_.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", session_.VrfName}
    session_.EntityData.Leafs["circuit-id"] = types.YLeaf{"CircuitId", session_.CircuitId}
    session_.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", session_.RemoteId}
    session_.EntityData.Leafs["lns-address"] = types.YLeaf{"LnsAddress", session_.LnsAddress}
    session_.EntityData.Leafs["lac-address"] = types.YLeaf{"LacAddress", session_.LacAddress}
    session_.EntityData.Leafs["tunnel-client-authentication-id"] = types.YLeaf{"TunnelClientAuthenticationId", session_.TunnelClientAuthenticationId}
    session_.EntityData.Leafs["tunnel-server-authentication-id"] = types.YLeaf{"TunnelServerAuthenticationId", session_.TunnelServerAuthenticationId}
    session_.EntityData.Leafs["session-ip-address"] = types.YLeaf{"SessionIpAddress", session_.SessionIpAddress}
    session_.EntityData.Leafs["session-ipv6-address"] = types.YLeaf{"SessionIpv6Address", session_.SessionIpv6Address}
    session_.EntityData.Leafs["session-ipv6-prefix"] = types.YLeaf{"SessionIpv6Prefix", session_.SessionIpv6Prefix}
    session_.EntityData.Leafs["delegated-ipv6-prefix"] = types.YLeaf{"DelegatedIpv6Prefix", session_.DelegatedIpv6Prefix}
    session_.EntityData.Leafs["ipv6-interface-id"] = types.YLeaf{"Ipv6InterfaceId", session_.Ipv6InterfaceId}
    session_.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", session_.MacAddress}
    session_.EntityData.Leafs["account-session-id"] = types.YLeaf{"AccountSessionId", session_.AccountSessionId}
    session_.EntityData.Leafs["nas-port"] = types.YLeaf{"NasPort", session_.NasPort}
    session_.EntityData.Leafs["username"] = types.YLeaf{"Username", session_.Username}
    session_.EntityData.Leafs["clientname"] = types.YLeaf{"Clientname", session_.Clientname}
    session_.EntityData.Leafs["formattedname"] = types.YLeaf{"Formattedname", session_.Formattedname}
    session_.EntityData.Leafs["is-session-authentic"] = types.YLeaf{"IsSessionAuthentic", session_.IsSessionAuthentic}
    session_.EntityData.Leafs["is-session-author"] = types.YLeaf{"IsSessionAuthor", session_.IsSessionAuthor}
    session_.EntityData.Leafs["session-state"] = types.YLeaf{"SessionState", session_.SessionState}
    session_.EntityData.Leafs["session-creation-time"] = types.YLeaf{"SessionCreationTime", session_.SessionCreationTime}
    session_.EntityData.Leafs["idle-state-change-time"] = types.YLeaf{"IdleStateChangeTime", session_.IdleStateChangeTime}
    session_.EntityData.Leafs["total-session-idle-time"] = types.YLeaf{"TotalSessionIdleTime", session_.TotalSessionIdleTime}
    session_.EntityData.Leafs["access-interface-name"] = types.YLeaf{"AccessInterfaceName", session_.AccessInterfaceName}
    session_.EntityData.Leafs["pending-callbacks"] = types.YLeaf{"PendingCallbacks", session_.PendingCallbacks}
    session_.EntityData.Leafs["af-up-status"] = types.YLeaf{"AfUpStatus", session_.AfUpStatus}
    session_.EntityData.Leafs["session-ipv4-state"] = types.YLeaf{"SessionIpv4State", session_.SessionIpv4State}
    session_.EntityData.Leafs["session-ipv6-state"] = types.YLeaf{"SessionIpv6State", session_.SessionIpv6State}
    return &(session_.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session__Accounting
// Accounting information
type Subscriber_Session_Nodes_Node_Sessions_Session__Accounting struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting information. The type is slice of
    // Subscriber_Session_Nodes_Node_Sessions_Session__Accounting_AccountingSession.
    AccountingSession []Subscriber_Session_Nodes_Node_Sessions_Session__Accounting_AccountingSession
}

func (accounting *Subscriber_Session_Nodes_Node_Sessions_Session__Accounting) GetEntityData() *types.CommonEntityData {
    accounting.EntityData.YFilter = accounting.YFilter
    accounting.EntityData.YangName = "accounting"
    accounting.EntityData.BundleName = "cisco_ios_xr"
    accounting.EntityData.ParentYangName = "session"
    accounting.EntityData.SegmentPath = "accounting"
    accounting.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accounting.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accounting.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accounting.EntityData.Children = make(map[string]types.YChild)
    accounting.EntityData.Children["accounting-session"] = types.YChild{"AccountingSession", nil}
    for i := range accounting.AccountingSession {
        accounting.EntityData.Children[types.GetSegmentPath(&accounting.AccountingSession[i])] = types.YChild{"AccountingSession", &accounting.AccountingSession[i]}
    }
    accounting.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accounting.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session__Accounting_AccountingSession
// Accounting information
type Subscriber_Session_Nodes_Node_Sessions_Session__Accounting_AccountingSession struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Accounting State Error Code for Accounting Session. The type is interface{}
    // with range: 0..4294967295.
    AccountingStateRc interface{}

    // Accounting Stop State. The type is interface{} with range: 0..4294967295.
    AccountingStopState interface{}

    // Accounting record context name. The type is string.
    RecordContextName interface{}

    // AAA method list name used to perform accounting. The type is string.
    MethodListName interface{}

    // Accounting session ID. The type is string.
    AccountSessionId interface{}

    // Accounting start time in DDD MMM DD HH:MM:SS YYYY format eg: Tue Feb 15
    // 15:12:49 2011. The type is string.
    AccountingStartTime interface{}

    // True if interim accounting is enabled. The type is bool.
    IsInterimAccountingEnabled interface{}

    // Interim accounting interval (in minutes). The type is interface{} with
    // range: 0..4294967295. Units are minute.
    InterimInterval interface{}

    // Time of last successful interim update in DDD MMM DD HH:MM:SS YYYY format
    // eg: Tue Apr 11 21:30 :47 2011. The type is string.
    LastSuccessfulInterimUpdateTime interface{}

    // Time of next interim update attempt (in seconds). The type is interface{}
    // with range: 0..4294967295. Units are second.
    NextInterimUpdateAttemptTime interface{}

    // Time of last interim update attempt in DDD MMM DD HH:MM:SS YYYY format eg:
    // Tue Apr 11 21:30:47 2011. The type is string.
    LastInterimUpdateAttemptTime interface{}

    // Number of interim updates sent. The type is interface{} with range:
    // 0..4294967295.
    SentInterimUpdates interface{}

    // Number of interim updates accepted. The type is interface{} with range:
    // 0..4294967295.
    AcceptedInterimUpdates interface{}

    // Number of interim updates rejected. The type is interface{} with range:
    // 0..4294967295.
    RejectedInterimUpdates interface{}

    // Number of interim update send failures. The type is interface{} with range:
    // 0..4294967295.
    SentInterimUpdateFailures interface{}
}

func (accountingSession *Subscriber_Session_Nodes_Node_Sessions_Session__Accounting_AccountingSession) GetEntityData() *types.CommonEntityData {
    accountingSession.EntityData.YFilter = accountingSession.YFilter
    accountingSession.EntityData.YangName = "accounting-session"
    accountingSession.EntityData.BundleName = "cisco_ios_xr"
    accountingSession.EntityData.ParentYangName = "accounting"
    accountingSession.EntityData.SegmentPath = "accounting-session"
    accountingSession.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accountingSession.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accountingSession.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accountingSession.EntityData.Children = make(map[string]types.YChild)
    accountingSession.EntityData.Leafs = make(map[string]types.YLeaf)
    accountingSession.EntityData.Leafs["accounting-state-rc"] = types.YLeaf{"AccountingStateRc", accountingSession.AccountingStateRc}
    accountingSession.EntityData.Leafs["accounting-stop-state"] = types.YLeaf{"AccountingStopState", accountingSession.AccountingStopState}
    accountingSession.EntityData.Leafs["record-context-name"] = types.YLeaf{"RecordContextName", accountingSession.RecordContextName}
    accountingSession.EntityData.Leafs["method-list-name"] = types.YLeaf{"MethodListName", accountingSession.MethodListName}
    accountingSession.EntityData.Leafs["account-session-id"] = types.YLeaf{"AccountSessionId", accountingSession.AccountSessionId}
    accountingSession.EntityData.Leafs["accounting-start-time"] = types.YLeaf{"AccountingStartTime", accountingSession.AccountingStartTime}
    accountingSession.EntityData.Leafs["is-interim-accounting-enabled"] = types.YLeaf{"IsInterimAccountingEnabled", accountingSession.IsInterimAccountingEnabled}
    accountingSession.EntityData.Leafs["interim-interval"] = types.YLeaf{"InterimInterval", accountingSession.InterimInterval}
    accountingSession.EntityData.Leafs["last-successful-interim-update-time"] = types.YLeaf{"LastSuccessfulInterimUpdateTime", accountingSession.LastSuccessfulInterimUpdateTime}
    accountingSession.EntityData.Leafs["next-interim-update-attempt-time"] = types.YLeaf{"NextInterimUpdateAttemptTime", accountingSession.NextInterimUpdateAttemptTime}
    accountingSession.EntityData.Leafs["last-interim-update-attempt-time"] = types.YLeaf{"LastInterimUpdateAttemptTime", accountingSession.LastInterimUpdateAttemptTime}
    accountingSession.EntityData.Leafs["sent-interim-updates"] = types.YLeaf{"SentInterimUpdates", accountingSession.SentInterimUpdates}
    accountingSession.EntityData.Leafs["accepted-interim-updates"] = types.YLeaf{"AcceptedInterimUpdates", accountingSession.AcceptedInterimUpdates}
    accountingSession.EntityData.Leafs["rejected-interim-updates"] = types.YLeaf{"RejectedInterimUpdates", accountingSession.RejectedInterimUpdates}
    accountingSession.EntityData.Leafs["sent-interim-update-failures"] = types.YLeaf{"SentInterimUpdateFailures", accountingSession.SentInterimUpdateFailures}
    return &(accountingSession.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session__UserProfileAttributes
// List of user profile attributes collected for
// subscriber session
type Subscriber_Session_Nodes_Node_Sessions_Session__UserProfileAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv4 maximum transmission unit. The type is interface{} with range:
    // 0..4294967295.
    Ipv4Mtu interface{}

    // IPv4 unnumbered. The type is string.
    Ipv4Unnumbered interface{}

    // Authorization service type. The type is AaaAuthService.
    AuthorizationServiceType interface{}

    // Tunnel client endpoint. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TunnelClientEndpoint interface{}

    // Tunnel server endpoint. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    TunnelServerEndpoint interface{}

    // Tunnel TOS setting. The type is interface{} with range: 0..4294967295.
    TunnelTosSetting interface{}

    // Tunnel medium. The type is AaaTunnelMedium.
    TunnelMedium interface{}

    // Tunnel preference. The type is interface{} with range: 0..4294967295.
    TunnelPreference interface{}

    // Tunnel client authentication ID. The type is string.
    TunnelClientAuthenticationId interface{}

    // Tunnel protocol. The type is AaaTunnelProto.
    TunnelProtocol interface{}

    // Actual data rate upstream (in Mbps). The type is interface{} with range:
    // 0..4294967295. Units are Mbit/s.
    ActualDataRateUpstream interface{}

    // Actual data rate downstream (in Mbps). The type is interface{} with range:
    // 0..4294967295. Units are Mbit/s.
    ActualDataRateDownstream interface{}

    // Attainable data rate upstream (in Mbps). The type is interface{} with
    // range: 0..4294967295. Units are Mbit/s.
    AttainableDataRateUpstream interface{}

    // Attainable data rate downstream (in Mbps). The type is interface{} with
    // range: 0..4294967295. Units are Mbit/s.
    AttainableDataRateDownstream interface{}

    // IP address pool. The type is string.
    PoolAddress interface{}

    // Circuit ID. The type is string.
    CircuitId interface{}

    // Connection receive speed. The type is interface{} with range:
    // 0..4294967295.
    ConnectionReceiveSpeed interface{}

    // Connection transmission speed. The type is interface{} with range:
    // 0..4294967295.
    ConnectionTransmissionSpeed interface{}

    // Destination station ID. The type is string.
    DestinationStationId interface{}

    // Primary DNS server address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryDnsServerAddress interface{}

    // Secondary DNS server address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SecondaryDnsServerAddress interface{}

    // Formatted calling station id. The type is string.
    FormattedCallingStationId interface{}

    // Interface name. The type is string.
    InterfaceName interface{}

    // Interface type. The type is AaaInterface.
    InterfaceType interface{}

    // Interim accounting interval. The type is interface{} with range:
    // 0..4294967295.
    InterimAccountingInterval interface{}

    // Ingress access list. The type is string.
    IngressAccessList interface{}

    // Egress access list. The type is string.
    EgressAccessList interface{}

    // IP netmask for the user. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    IpNetmask interface{}

    // True, if interworking functionality. The type is bool.
    IsInterworkingFunctionality interface{}

    // Maximum interleaving delay downstream (in Mbps). The type is interface{}
    // with range: 0..4294967295. Units are Mbit/s.
    MaxInterleavingDelayDownstream interface{}

    // Maximum interleaving delay upstream (in Mbps). The type is interface{} with
    // range: 0..4294967295. Units are Mbit/s.
    MaxInterleavingDelayUpstream interface{}

    // Maximum data rate upstream (in Mbps). The type is interface{} with range:
    // 0..4294967295. Units are Mbit/s.
    MaxDataRateUpstream interface{}

    // Maximum data rate downstream (in Mbps). The type is interface{} with range:
    // 0..4294967295. Units are Mbit/s.
    MaxDataRateDownstream interface{}

    // Minimum data rate downstream (in Mbps). The type is interface{} with range:
    // 0..4294967295. Units are Mbit/s.
    MinDataRateDownstream interface{}

    // Minimum data rate downstream low power (in Mbps). The type is interface{}
    // with range: 0..4294967295. Units are Mbit/s.
    MinDataRateDownstreamLowPower interface{}

    // Minimum data rate upstream low power (in Mbps). The type is interface{}
    // with range: 0..4294967295. Units are Mbit/s.
    MinDataRateUpstreamLowPower interface{}

    // Primary net bios server address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryNetBiosServerAddress interface{}

    // Secondary net bios server address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SecondaryNetBiosServerAddress interface{}

    // Parent interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    ParentInterfaceName interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Route information for a user session. The type is string.
    Route interface{}

    // Session timeout (in seconds). The type is interface{} with range:
    // 0..4294967295. Units are second.
    SessionTimeout interface{}

    // Strict RPF packets. The type is interface{} with range: 0..4294967295.
    StrictRpfPackets interface{}

    // Accounting session ID. The type is string.
    AccountingSessionId interface{}

    // Upstream parameterized QoS policy to be applied on the subscriber side. The
    // type is string.
    UpstreamParameterizedQosPolicy interface{}

    // Downstream parameterized QoS policy to be applied on the subscriber side.
    // The type is string.
    DownstreamParameterizedQosPolicy interface{}

    // Upstream QoS policy to be applied on the subscriber side. The type is
    // string.
    UpstreamQosPolicy interface{}

    // Downstream QoS policy to be applied on the subscriber side. The type is
    // string.
    DownstreamQosPolicy interface{}

    // Session termination cause. The type is AaaTerminateCause.
    SessionTerminationCause interface{}
}

func (userProfileAttributes *Subscriber_Session_Nodes_Node_Sessions_Session__UserProfileAttributes) GetEntityData() *types.CommonEntityData {
    userProfileAttributes.EntityData.YFilter = userProfileAttributes.YFilter
    userProfileAttributes.EntityData.YangName = "user-profile-attributes"
    userProfileAttributes.EntityData.BundleName = "cisco_ios_xr"
    userProfileAttributes.EntityData.ParentYangName = "session"
    userProfileAttributes.EntityData.SegmentPath = "user-profile-attributes"
    userProfileAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    userProfileAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    userProfileAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    userProfileAttributes.EntityData.Children = make(map[string]types.YChild)
    userProfileAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    userProfileAttributes.EntityData.Leafs["ipv4mtu"] = types.YLeaf{"Ipv4Mtu", userProfileAttributes.Ipv4Mtu}
    userProfileAttributes.EntityData.Leafs["ipv4-unnumbered"] = types.YLeaf{"Ipv4Unnumbered", userProfileAttributes.Ipv4Unnumbered}
    userProfileAttributes.EntityData.Leafs["authorization-service-type"] = types.YLeaf{"AuthorizationServiceType", userProfileAttributes.AuthorizationServiceType}
    userProfileAttributes.EntityData.Leafs["tunnel-client-endpoint"] = types.YLeaf{"TunnelClientEndpoint", userProfileAttributes.TunnelClientEndpoint}
    userProfileAttributes.EntityData.Leafs["tunnel-server-endpoint"] = types.YLeaf{"TunnelServerEndpoint", userProfileAttributes.TunnelServerEndpoint}
    userProfileAttributes.EntityData.Leafs["tunnel-tos-setting"] = types.YLeaf{"TunnelTosSetting", userProfileAttributes.TunnelTosSetting}
    userProfileAttributes.EntityData.Leafs["tunnel-medium"] = types.YLeaf{"TunnelMedium", userProfileAttributes.TunnelMedium}
    userProfileAttributes.EntityData.Leafs["tunnel-preference"] = types.YLeaf{"TunnelPreference", userProfileAttributes.TunnelPreference}
    userProfileAttributes.EntityData.Leafs["tunnel-client-authentication-id"] = types.YLeaf{"TunnelClientAuthenticationId", userProfileAttributes.TunnelClientAuthenticationId}
    userProfileAttributes.EntityData.Leafs["tunnel-protocol"] = types.YLeaf{"TunnelProtocol", userProfileAttributes.TunnelProtocol}
    userProfileAttributes.EntityData.Leafs["actual-data-rate-upstream"] = types.YLeaf{"ActualDataRateUpstream", userProfileAttributes.ActualDataRateUpstream}
    userProfileAttributes.EntityData.Leafs["actual-data-rate-downstream"] = types.YLeaf{"ActualDataRateDownstream", userProfileAttributes.ActualDataRateDownstream}
    userProfileAttributes.EntityData.Leafs["attainable-data-rate-upstream"] = types.YLeaf{"AttainableDataRateUpstream", userProfileAttributes.AttainableDataRateUpstream}
    userProfileAttributes.EntityData.Leafs["attainable-data-rate-downstream"] = types.YLeaf{"AttainableDataRateDownstream", userProfileAttributes.AttainableDataRateDownstream}
    userProfileAttributes.EntityData.Leafs["pool-address"] = types.YLeaf{"PoolAddress", userProfileAttributes.PoolAddress}
    userProfileAttributes.EntityData.Leafs["circuit-id"] = types.YLeaf{"CircuitId", userProfileAttributes.CircuitId}
    userProfileAttributes.EntityData.Leafs["connection-receive-speed"] = types.YLeaf{"ConnectionReceiveSpeed", userProfileAttributes.ConnectionReceiveSpeed}
    userProfileAttributes.EntityData.Leafs["connection-transmission-speed"] = types.YLeaf{"ConnectionTransmissionSpeed", userProfileAttributes.ConnectionTransmissionSpeed}
    userProfileAttributes.EntityData.Leafs["destination-station-id"] = types.YLeaf{"DestinationStationId", userProfileAttributes.DestinationStationId}
    userProfileAttributes.EntityData.Leafs["primary-dns-server-address"] = types.YLeaf{"PrimaryDnsServerAddress", userProfileAttributes.PrimaryDnsServerAddress}
    userProfileAttributes.EntityData.Leafs["secondary-dns-server-address"] = types.YLeaf{"SecondaryDnsServerAddress", userProfileAttributes.SecondaryDnsServerAddress}
    userProfileAttributes.EntityData.Leafs["formatted-calling-station-id"] = types.YLeaf{"FormattedCallingStationId", userProfileAttributes.FormattedCallingStationId}
    userProfileAttributes.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", userProfileAttributes.InterfaceName}
    userProfileAttributes.EntityData.Leafs["interface-type"] = types.YLeaf{"InterfaceType", userProfileAttributes.InterfaceType}
    userProfileAttributes.EntityData.Leafs["interim-accounting-interval"] = types.YLeaf{"InterimAccountingInterval", userProfileAttributes.InterimAccountingInterval}
    userProfileAttributes.EntityData.Leafs["ingress-access-list"] = types.YLeaf{"IngressAccessList", userProfileAttributes.IngressAccessList}
    userProfileAttributes.EntityData.Leafs["egress-access-list"] = types.YLeaf{"EgressAccessList", userProfileAttributes.EgressAccessList}
    userProfileAttributes.EntityData.Leafs["ip-netmask"] = types.YLeaf{"IpNetmask", userProfileAttributes.IpNetmask}
    userProfileAttributes.EntityData.Leafs["is-interworking-functionality"] = types.YLeaf{"IsInterworkingFunctionality", userProfileAttributes.IsInterworkingFunctionality}
    userProfileAttributes.EntityData.Leafs["max-interleaving-delay-downstream"] = types.YLeaf{"MaxInterleavingDelayDownstream", userProfileAttributes.MaxInterleavingDelayDownstream}
    userProfileAttributes.EntityData.Leafs["max-interleaving-delay-upstream"] = types.YLeaf{"MaxInterleavingDelayUpstream", userProfileAttributes.MaxInterleavingDelayUpstream}
    userProfileAttributes.EntityData.Leafs["max-data-rate-upstream"] = types.YLeaf{"MaxDataRateUpstream", userProfileAttributes.MaxDataRateUpstream}
    userProfileAttributes.EntityData.Leafs["max-data-rate-downstream"] = types.YLeaf{"MaxDataRateDownstream", userProfileAttributes.MaxDataRateDownstream}
    userProfileAttributes.EntityData.Leafs["min-data-rate-downstream"] = types.YLeaf{"MinDataRateDownstream", userProfileAttributes.MinDataRateDownstream}
    userProfileAttributes.EntityData.Leafs["min-data-rate-downstream-low-power"] = types.YLeaf{"MinDataRateDownstreamLowPower", userProfileAttributes.MinDataRateDownstreamLowPower}
    userProfileAttributes.EntityData.Leafs["min-data-rate-upstream-low-power"] = types.YLeaf{"MinDataRateUpstreamLowPower", userProfileAttributes.MinDataRateUpstreamLowPower}
    userProfileAttributes.EntityData.Leafs["primary-net-bios-server-address"] = types.YLeaf{"PrimaryNetBiosServerAddress", userProfileAttributes.PrimaryNetBiosServerAddress}
    userProfileAttributes.EntityData.Leafs["secondary-net-bios-server-address"] = types.YLeaf{"SecondaryNetBiosServerAddress", userProfileAttributes.SecondaryNetBiosServerAddress}
    userProfileAttributes.EntityData.Leafs["parent-interface-name"] = types.YLeaf{"ParentInterfaceName", userProfileAttributes.ParentInterfaceName}
    userProfileAttributes.EntityData.Leafs["remote-id"] = types.YLeaf{"RemoteId", userProfileAttributes.RemoteId}
    userProfileAttributes.EntityData.Leafs["route"] = types.YLeaf{"Route", userProfileAttributes.Route}
    userProfileAttributes.EntityData.Leafs["session-timeout"] = types.YLeaf{"SessionTimeout", userProfileAttributes.SessionTimeout}
    userProfileAttributes.EntityData.Leafs["strict-rpf-packets"] = types.YLeaf{"StrictRpfPackets", userProfileAttributes.StrictRpfPackets}
    userProfileAttributes.EntityData.Leafs["accounting-session-id"] = types.YLeaf{"AccountingSessionId", userProfileAttributes.AccountingSessionId}
    userProfileAttributes.EntityData.Leafs["upstream-parameterized-qos-policy"] = types.YLeaf{"UpstreamParameterizedQosPolicy", userProfileAttributes.UpstreamParameterizedQosPolicy}
    userProfileAttributes.EntityData.Leafs["downstream-parameterized-qos-policy"] = types.YLeaf{"DownstreamParameterizedQosPolicy", userProfileAttributes.DownstreamParameterizedQosPolicy}
    userProfileAttributes.EntityData.Leafs["upstream-qos-policy"] = types.YLeaf{"UpstreamQosPolicy", userProfileAttributes.UpstreamQosPolicy}
    userProfileAttributes.EntityData.Leafs["downstream-qos-policy"] = types.YLeaf{"DownstreamQosPolicy", userProfileAttributes.DownstreamQosPolicy}
    userProfileAttributes.EntityData.Leafs["session-termination-cause"] = types.YLeaf{"SessionTerminationCause", userProfileAttributes.SessionTerminationCause}
    return &(userProfileAttributes.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session__MobilityAttributes
// List of mobility attributes collected for
// subscriber session
type Subscriber_Session_Nodes_Node_Sessions_Session__MobilityAttributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Cisco MPC Protocol. The type is bool.
    MpcProtocol interface{}

    // IPv4 address of Mobility Node. The type is string.
    MobilityIpv4Address interface{}

    // Default Gateway IPv4 Address. The type is string.
    MobilityDefaultIpv4Gateway interface{}

    // DNS Server Primary. The type is string.
    MobilityDnsServer interface{}

    // DHCP Server. The type is string.
    MobilityDhcpServer interface{}

    // IPv4 Netmask. The type is string.
    MobilityIpv4Netmask interface{}

    // Domain Name. The type is string.
    DomainName interface{}

    // Uplink GRE Key. The type is string.
    UplinkGreKey interface{}

    // Downlink GRE Key. The type is string.
    DownlinkGreKey interface{}

    // Duration of lease in seconds. The type is string. Units are second.
    LeaseTime interface{}
}

func (mobilityAttributes *Subscriber_Session_Nodes_Node_Sessions_Session__MobilityAttributes) GetEntityData() *types.CommonEntityData {
    mobilityAttributes.EntityData.YFilter = mobilityAttributes.YFilter
    mobilityAttributes.EntityData.YangName = "mobility-attributes"
    mobilityAttributes.EntityData.BundleName = "cisco_ios_xr"
    mobilityAttributes.EntityData.ParentYangName = "session"
    mobilityAttributes.EntityData.SegmentPath = "mobility-attributes"
    mobilityAttributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mobilityAttributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mobilityAttributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mobilityAttributes.EntityData.Children = make(map[string]types.YChild)
    mobilityAttributes.EntityData.Leafs = make(map[string]types.YLeaf)
    mobilityAttributes.EntityData.Leafs["mpc-protocol"] = types.YLeaf{"MpcProtocol", mobilityAttributes.MpcProtocol}
    mobilityAttributes.EntityData.Leafs["mobility-ipv4-address"] = types.YLeaf{"MobilityIpv4Address", mobilityAttributes.MobilityIpv4Address}
    mobilityAttributes.EntityData.Leafs["mobility-default-ipv4-gateway"] = types.YLeaf{"MobilityDefaultIpv4Gateway", mobilityAttributes.MobilityDefaultIpv4Gateway}
    mobilityAttributes.EntityData.Leafs["mobility-dns-server"] = types.YLeaf{"MobilityDnsServer", mobilityAttributes.MobilityDnsServer}
    mobilityAttributes.EntityData.Leafs["mobility-dhcp-server"] = types.YLeaf{"MobilityDhcpServer", mobilityAttributes.MobilityDhcpServer}
    mobilityAttributes.EntityData.Leafs["mobility-ipv4-netmask"] = types.YLeaf{"MobilityIpv4Netmask", mobilityAttributes.MobilityIpv4Netmask}
    mobilityAttributes.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", mobilityAttributes.DomainName}
    mobilityAttributes.EntityData.Leafs["uplink-gre-key"] = types.YLeaf{"UplinkGreKey", mobilityAttributes.UplinkGreKey}
    mobilityAttributes.EntityData.Leafs["downlink-gre-key"] = types.YLeaf{"DownlinkGreKey", mobilityAttributes.DownlinkGreKey}
    mobilityAttributes.EntityData.Leafs["lease-time"] = types.YLeaf{"LeaseTime", mobilityAttributes.LeaseTime}
    return &(mobilityAttributes.EntityData)
}

// Subscriber_Session_Nodes_Node_Sessions_Session__SessionChangeOfAuthorization
// Subscriber change of authorization information
type Subscriber_Session_Nodes_Node_Sessions_Session__SessionChangeOfAuthorization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Coa Request Acked. The type is bool.
    RequestAcked interface{}

    // Request time in DDD MMM DD HH:MM:SS YYYY format eg: Tue Apr 11 21:30:47
    // 2011. The type is string.
    RequestTime interface{}

    // List of Request Attributes collected in COA response. The type is string
    // with pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    CoaRequestAttributes interface{}

    // Reply time in DDD MMM DD HH:MM:SS YYYY format eg : Tue Apr 11 21:30:47
    // 2011. The type is string.
    ReplyTime interface{}

    // List of Reply Attributes collected in COA response. The type is string with
    // pattern: b'([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?'.
    CoaReplyAttributes interface{}
}

func (sessionChangeOfAuthorization *Subscriber_Session_Nodes_Node_Sessions_Session__SessionChangeOfAuthorization) GetEntityData() *types.CommonEntityData {
    sessionChangeOfAuthorization.EntityData.YFilter = sessionChangeOfAuthorization.YFilter
    sessionChangeOfAuthorization.EntityData.YangName = "session-change-of-authorization"
    sessionChangeOfAuthorization.EntityData.BundleName = "cisco_ios_xr"
    sessionChangeOfAuthorization.EntityData.ParentYangName = "session"
    sessionChangeOfAuthorization.EntityData.SegmentPath = "session-change-of-authorization"
    sessionChangeOfAuthorization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionChangeOfAuthorization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionChangeOfAuthorization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionChangeOfAuthorization.EntityData.Children = make(map[string]types.YChild)
    sessionChangeOfAuthorization.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionChangeOfAuthorization.EntityData.Leafs["request-acked"] = types.YLeaf{"RequestAcked", sessionChangeOfAuthorization.RequestAcked}
    sessionChangeOfAuthorization.EntityData.Leafs["request-time"] = types.YLeaf{"RequestTime", sessionChangeOfAuthorization.RequestTime}
    sessionChangeOfAuthorization.EntityData.Leafs["coa-request-attributes"] = types.YLeaf{"CoaRequestAttributes", sessionChangeOfAuthorization.CoaRequestAttributes}
    sessionChangeOfAuthorization.EntityData.Leafs["reply-time"] = types.YLeaf{"ReplyTime", sessionChangeOfAuthorization.ReplyTime}
    sessionChangeOfAuthorization.EntityData.Leafs["coa-reply-attributes"] = types.YLeaf{"CoaReplyAttributes", sessionChangeOfAuthorization.CoaReplyAttributes}
    return &(sessionChangeOfAuthorization.EntityData)
}

// IedgeLicenseManager
// iedge license manager
type IedgeLicenseManager struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session License Manager operational data for a location.
    Nodes IedgeLicenseManager_Nodes
}

func (iedgeLicenseManager *IedgeLicenseManager) GetEntityData() *types.CommonEntityData {
    iedgeLicenseManager.EntityData.YFilter = iedgeLicenseManager.YFilter
    iedgeLicenseManager.EntityData.YangName = "iedge-license-manager"
    iedgeLicenseManager.EntityData.BundleName = "cisco_ios_xr"
    iedgeLicenseManager.EntityData.ParentYangName = "Cisco-IOS-XR-iedge4710-oper"
    iedgeLicenseManager.EntityData.SegmentPath = "Cisco-IOS-XR-iedge4710-oper:iedge-license-manager"
    iedgeLicenseManager.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iedgeLicenseManager.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iedgeLicenseManager.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iedgeLicenseManager.EntityData.Children = make(map[string]types.YChild)
    iedgeLicenseManager.EntityData.Children["nodes"] = types.YChild{"Nodes", &iedgeLicenseManager.Nodes}
    iedgeLicenseManager.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(iedgeLicenseManager.EntityData)
}

// IedgeLicenseManager_Nodes
// Session License Manager operational data for a
// location
type IedgeLicenseManager_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Location. For example, 0/1/CPU0. The type is slice of
    // IedgeLicenseManager_Nodes_Node.
    Node []IedgeLicenseManager_Nodes_Node
}

func (nodes *IedgeLicenseManager_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "iedge-license-manager"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// IedgeLicenseManager_Nodes_Node
// Location. For example, 0/1/CPU0
type IedgeLicenseManager_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The node id to filter on. For example, 0/1/CPU0.
    // The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
    Nodeid interface{}

    // Display Session License Manager summary data.
    IedgeLicenseManagerSummary IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary
}

func (node *IedgeLicenseManager_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + "[nodeid='" + fmt.Sprintf("%v", node.Nodeid) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Children["iedge-license-manager-summary"] = types.YChild{"IedgeLicenseManagerSummary", &node.IedgeLicenseManagerSummary}
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["nodeid"] = types.YLeaf{"Nodeid", node.Nodeid}
    return &(node.EntityData)
}

// IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary
// Display Session License Manager summary data
type IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // configured session limit. The type is interface{} with range:
    // 0..4294967295.
    SessionLimit interface{}

    // configured session threshold. The type is interface{} with range:
    // 0..4294967295.
    SessionThreshold interface{}

    // number of license. The type is interface{} with range: 0..4294967295.
    SessionLicenseCount interface{}

    // number of sessions. The type is interface{} with range: 0..4294967295.
    SessionCount interface{}
}

func (iedgeLicenseManagerSummary *IedgeLicenseManager_Nodes_Node_IedgeLicenseManagerSummary) GetEntityData() *types.CommonEntityData {
    iedgeLicenseManagerSummary.EntityData.YFilter = iedgeLicenseManagerSummary.YFilter
    iedgeLicenseManagerSummary.EntityData.YangName = "iedge-license-manager-summary"
    iedgeLicenseManagerSummary.EntityData.BundleName = "cisco_ios_xr"
    iedgeLicenseManagerSummary.EntityData.ParentYangName = "node"
    iedgeLicenseManagerSummary.EntityData.SegmentPath = "iedge-license-manager-summary"
    iedgeLicenseManagerSummary.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    iedgeLicenseManagerSummary.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    iedgeLicenseManagerSummary.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    iedgeLicenseManagerSummary.EntityData.Children = make(map[string]types.YChild)
    iedgeLicenseManagerSummary.EntityData.Leafs = make(map[string]types.YLeaf)
    iedgeLicenseManagerSummary.EntityData.Leafs["session-limit"] = types.YLeaf{"SessionLimit", iedgeLicenseManagerSummary.SessionLimit}
    iedgeLicenseManagerSummary.EntityData.Leafs["session-threshold"] = types.YLeaf{"SessionThreshold", iedgeLicenseManagerSummary.SessionThreshold}
    iedgeLicenseManagerSummary.EntityData.Leafs["session-license-count"] = types.YLeaf{"SessionLicenseCount", iedgeLicenseManagerSummary.SessionLicenseCount}
    iedgeLicenseManagerSummary.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", iedgeLicenseManagerSummary.SessionCount}
    return &(iedgeLicenseManagerSummary.EntityData)
}

