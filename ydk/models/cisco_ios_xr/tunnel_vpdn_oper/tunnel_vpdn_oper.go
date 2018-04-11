// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-vpdn package operational data.
// 
// This module contains definitions
// for the following management objects:
//   vpdn: VPDN operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_vpdn_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_vpdn_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-vpdn-oper vpdn}", reflect.TypeOf(Vpdn{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-vpdn-oper:vpdn", reflect.TypeOf(Vpdn{}))
}

// VpdnFailcode represents VPDN failure types
type VpdnFailcode string

const (
    // Unknown
    VpdnFailcode_unknown VpdnFailcode = "unknown"

    // Peer action
    VpdnFailcode_peer_action VpdnFailcode = "peer-action"

    // Authentication
    VpdnFailcode_authentication VpdnFailcode = "authentication"

    // Authentication
    VpdnFailcode_authentication_error VpdnFailcode = "authentication-error"

    // Authentication
    VpdnFailcode_authentication_failed VpdnFailcode = "authentication-failed"

    // Authorization
    VpdnFailcode_authorization VpdnFailcode = "authorization"

    // Authorization
    VpdnFailcode_authorization_failed VpdnFailcode = "authorization-failed"

    // No resources available
    VpdnFailcode_home_gatewayfailure VpdnFailcode = "home-gatewayfailure"

    // Connection termination
    VpdnFailcode_connection_termination VpdnFailcode = "connection-termination"

    // No resources available
    VpdnFailcode_no_resources_available VpdnFailcode = "no-resources-available"

    // Timer expiry
    VpdnFailcode_timer_expiry VpdnFailcode = "timer-expiry"

    // Session limit
    VpdnFailcode_session_mid_exceeded VpdnFailcode = "session-mid-exceeded"

    // Softshut
    VpdnFailcode_soft_shut VpdnFailcode = "soft-shut"

    // Session limit
    VpdnFailcode_session_limit_exceeded VpdnFailcode = "session-limit-exceeded"

    // Administrative
    VpdnFailcode_administrative VpdnFailcode = "administrative"

    // Link failure
    VpdnFailcode_link_failure VpdnFailcode = "link-failure"

    // Security
    VpdnFailcode_security VpdnFailcode = "security"

    // Tunnel in HA resync
    VpdnFailcode_tunnel_in_resync VpdnFailcode = "tunnel-in-resync"

    // Call parameters
    VpdnFailcode_call_prarmeters VpdnFailcode = "call-prarmeters"
)

// VpdnState represents Vpdn states
type VpdnState string

const (
    // Initial state
    VpdnState_initial_state VpdnState = "initial-state"

    // Initial Sync in progress
    VpdnState_init_sync_in_progress VpdnState = "init-sync-in-progress"

    // Initial Sync Done
    VpdnState_steady_state VpdnState = "steady-state"
)

// LsgStatus represents LSG Status
type LsgStatus string

const (
    // Member not initialized.
    LsgStatus_none LsgStatus = "none"

    // Member is active.
    LsgStatus_active LsgStatus = "active"

    // Member is down, cannot create new sessions.
    LsgStatus_down LsgStatus = "down"

    // Member is ready for new sessions.
    LsgStatus_testable LsgStatus = "testable"

    // Member is being tested for new session
    LsgStatus_testing LsgStatus = "testing"
)

// TosMode represents TOS modes
type TosMode string

const (
    // default
    TosMode_default_ TosMode = "default"

    // set
    TosMode_set TosMode = "set"

    // reflect
    TosMode_reflect TosMode = "reflect"
)

// VpdnNasPort represents NAS port types
type VpdnNasPort string

const (
    // None
    VpdnNasPort_none VpdnNasPort = "none"

    // Primary
    VpdnNasPort_primary VpdnNasPort = "primary"

    // BRI
    VpdnNasPort_bri VpdnNasPort = "bri"

    // Serial
    VpdnNasPort_serial VpdnNasPort = "serial"

    // Asynchronous
    VpdnNasPort_asynchronous VpdnNasPort = "asynchronous"

    // VTY
    VpdnNasPort_vty VpdnNasPort = "vty"

    // Asynchronous transfer mode
    VpdnNasPort_atm VpdnNasPort = "atm"

    // Ethernet
    VpdnNasPort_ethernet VpdnNasPort = "ethernet"

    // PPP over ATM
    VpdnNasPort_ppp_atm VpdnNasPort = "ppp-atm"

    // PPPoE over ATM
    VpdnNasPort_pppoe_over_atm VpdnNasPort = "pppoe-over-atm"

    // PPPoE over Ethernet
    VpdnNasPort_pppoe_over_ethernet VpdnNasPort = "pppoe-over-ethernet"

    // PPPoE over VLAN
    VpdnNasPort_pppoe_over_vlan VpdnNasPort = "pppoe-over-vlan"

    // PPPoE over Q-in-Q
    VpdnNasPort_pppoe_over_q_in_q VpdnNasPort = "pppoe-over-q-in-q"

    //  V120
    VpdnNasPort_v120 VpdnNasPort = "v120"

    // V110
    VpdnNasPort_v110 VpdnNasPort = "v110"

    // PIAFS
    VpdnNasPort_piafs VpdnNasPort = "piafs"

    // X.75
    VpdnNasPort_x75 VpdnNasPort = "x75"

    // IPSec
    VpdnNasPort_ip_sec VpdnNasPort = "ip-sec"

    // Other
    VpdnNasPort_other VpdnNasPort = "other"

    // Virtual PPPoE over Ethernet
    VpdnNasPort_virtual_pppoe_over_ethernet VpdnNasPort = "virtual-pppoe-over-ethernet"

    //  Virtual PPPoE over VLAN
    VpdnNasPort_virtual_pppoe_over_vlan VpdnNasPort = "virtual-pppoe-over-vlan"

    // Virtual PPPoE over Q-in-Q
    VpdnNasPort_virtual_pppoe_over_q_in_q VpdnNasPort = "virtual-pppoe-over-q-in-q"

    // IPoE over Ethernet
    VpdnNasPort_ipo_e_over_ethernet VpdnNasPort = "ipo-e-over-ethernet"

    // IPoE over VLAN
    VpdnNasPort_ipo_e_over_vlan VpdnNasPort = "ipo-e-over-vlan"

    // IPoE over Q-in-Q
    VpdnNasPort_ipo_e_over_q_in_q VpdnNasPort = "ipo-e-over-q-in-q"

    // Virtual IPoE over ethernet
    VpdnNasPort_virtual_i_po_e_over_ethernet VpdnNasPort = "virtual-i-po-e-over-ethernet"

    // Virtual IPoE over VLAN
    VpdnNasPort_virtual_i_po_e_over_vlan VpdnNasPort = "virtual-i-po-e-over-vlan"

    // Virtual IPoE over Q-in-Q
    VpdnNasPort_virtual_i_po_e_over_q_in_q VpdnNasPort = "virtual-i-po-e-over-q-in-q"

    // Unknown
    VpdnNasPort_unknown VpdnNasPort = "unknown"
)

// SessionState represents Session states
type SessionState string

const (
    // Idle state
    SessionState_idle SessionState = "idle"

    // Connected state
    SessionState_connected SessionState = "connected"

    // Established state
    SessionState_established SessionState = "established"
)

// Vpdn
// VPDN operational data
type Vpdn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN session list.
    Sessions Vpdn_Sessions

    // VPDN tunnel Destinations.
    TunnelDestinations Vpdn_TunnelDestinations

    // VPDN Mirroring Statistics.
    VpdnMirroring Vpdn_VpdnMirroring

    // Show VPDN Redundancy information.
    VpdnRedundancy Vpdn_VpdnRedundancy

    // VPDN history failure list.
    HistoryFailures Vpdn_HistoryFailures
}

func (vpdn *Vpdn) GetEntityData() *types.CommonEntityData {
    vpdn.EntityData.YFilter = vpdn.YFilter
    vpdn.EntityData.YangName = "vpdn"
    vpdn.EntityData.BundleName = "cisco_ios_xr"
    vpdn.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-vpdn-oper"
    vpdn.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-vpdn-oper:vpdn"
    vpdn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdn.EntityData.Children = make(map[string]types.YChild)
    vpdn.EntityData.Children["sessions"] = types.YChild{"Sessions", &vpdn.Sessions}
    vpdn.EntityData.Children["tunnel-destinations"] = types.YChild{"TunnelDestinations", &vpdn.TunnelDestinations}
    vpdn.EntityData.Children["vpdn-mirroring"] = types.YChild{"VpdnMirroring", &vpdn.VpdnMirroring}
    vpdn.EntityData.Children["vpdn-redundancy"] = types.YChild{"VpdnRedundancy", &vpdn.VpdnRedundancy}
    vpdn.EntityData.Children["history-failures"] = types.YChild{"HistoryFailures", &vpdn.HistoryFailures}
    vpdn.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vpdn.EntityData)
}

// Vpdn_Sessions
// VPDN session list
type Vpdn_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN session information. The type is slice of Vpdn_Sessions_Session.
    Session []Vpdn_Sessions_Session
}

func (sessions *Vpdn_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "vpdn"
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

// Vpdn_Sessions_Session
//  VPDN session information
type Vpdn_Sessions_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Session label. The type is string with pattern:
    // b'[0-9a-fA-F]{1,8}'.
    SessionLabel interface{}

    // Time to setup session in sec:msec. The type is interface{} with range:
    // 0..4294967295.
    SetupTime interface{}

    // Parent interface name. The type is string with pattern: b'[a-zA-Z0-9./-]+'.
    ParentInterfaceName interface{}

    // Session data.
    Session Vpdn_Sessions_Session_Session_

    // L2TP data.
    L2Tp Vpdn_Sessions_Session_L2Tp

    // Subscriber data.
    Subscriber Vpdn_Sessions_Session_Subscriber

    // Configuration data.
    Configuration Vpdn_Sessions_Session_Configuration
}

func (session *Vpdn_Sessions_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "sessions"
    session.EntityData.SegmentPath = "session" + "[session-label='" + fmt.Sprintf("%v", session.SessionLabel) + "']"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = make(map[string]types.YChild)
    session.EntityData.Children["session"] = types.YChild{"Session", &session.Session}
    session.EntityData.Children["l2tp"] = types.YChild{"L2Tp", &session.L2Tp}
    session.EntityData.Children["subscriber"] = types.YChild{"Subscriber", &session.Subscriber}
    session.EntityData.Children["configuration"] = types.YChild{"Configuration", &session.Configuration}
    session.EntityData.Leafs = make(map[string]types.YLeaf)
    session.EntityData.Leafs["session-label"] = types.YLeaf{"SessionLabel", session.SessionLabel}
    session.EntityData.Leafs["setup-time"] = types.YLeaf{"SetupTime", session.SetupTime}
    session.EntityData.Leafs["parent-interface-name"] = types.YLeaf{"ParentInterfaceName", session.ParentInterfaceName}
    return &(session.EntityData)
}

// Vpdn_Sessions_Session_Session_
// Session data
type Vpdn_Sessions_Session_Session_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Elapsed time since last change in hh:mm:ss format. The type is string.
    LastChange interface{}

    // Session interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Authentication username. The type is string.
    Username interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // Session state. The type is SessionState.
    State interface{}

    // L2TP session ID. The type is interface{} with range: 0..65535.
    L2TpSessionId interface{}

    // L2TP tunnel ID. The type is interface{} with range: 0..65535.
    L2TpTunnelId interface{}

    // Session SRG Slave. The type is bool.
    SrgSlave interface{}
}

func (session_ *Vpdn_Sessions_Session_Session_) GetEntityData() *types.CommonEntityData {
    session_.EntityData.YFilter = session_.YFilter
    session_.EntityData.YangName = "session"
    session_.EntityData.BundleName = "cisco_ios_xr"
    session_.EntityData.ParentYangName = "session"
    session_.EntityData.SegmentPath = "session"
    session_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session_.EntityData.Children = make(map[string]types.YChild)
    session_.EntityData.Leafs = make(map[string]types.YLeaf)
    session_.EntityData.Leafs["last-change"] = types.YLeaf{"LastChange", session_.LastChange}
    session_.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", session_.InterfaceName}
    session_.EntityData.Leafs["username"] = types.YLeaf{"Username", session_.Username}
    session_.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", session_.DomainName}
    session_.EntityData.Leafs["state"] = types.YLeaf{"State", session_.State}
    session_.EntityData.Leafs["l2tp-session-id"] = types.YLeaf{"L2TpSessionId", session_.L2TpSessionId}
    session_.EntityData.Leafs["l2tp-tunnel-id"] = types.YLeaf{"L2TpTunnelId", session_.L2TpTunnelId}
    session_.EntityData.Leafs["srg-slave"] = types.YLeaf{"SrgSlave", session_.SrgSlave}
    return &(session_.EntityData)
}

// Vpdn_Sessions_Session_L2Tp
// L2TP data
type Vpdn_Sessions_Session_L2Tp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local endpoint IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    LocalEndpoint interface{}

    // Remote endpoint IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RemoteEndpoint interface{}

    // Call serial number. The type is interface{} with range: 0..4294967295.
    CallSerialNumber interface{}

    // True if L2TP class attribute mask is set. The type is bool.
    IsL2TpClassAttributeMaskSet interface{}

    // Local tunnel ID. The type is interface{} with range: 0..65535.
    LocalTunnelId interface{}

    // Remote tunnel ID. The type is interface{} with range: 0..65535.
    RemoteTunnelId interface{}

    // Local session ID. The type is interface{} with range: 0..65535.
    LocalSessionId interface{}

    // Remote session ID. The type is interface{} with range: 0..65535.
    RemoteSessionId interface{}

    // Remote port. The type is interface{} with range: 0..65535.
    RemotePort interface{}

    // Tunnel client authentication ID. The type is string.
    TunnelClientAuthenticationId interface{}

    // Tunnel server authentication ID. The type is string.
    TunnelServerAuthenticationId interface{}

    // Tunnel assignment ID. The type is string.
    TunnelAssignmentId interface{}

    // True if tunnel authentication is enabled. The type is bool.
    IsTunnelAuthenticationEnabled interface{}
}

func (l2Tp *Vpdn_Sessions_Session_L2Tp) GetEntityData() *types.CommonEntityData {
    l2Tp.EntityData.YFilter = l2Tp.YFilter
    l2Tp.EntityData.YangName = "l2tp"
    l2Tp.EntityData.BundleName = "cisco_ios_xr"
    l2Tp.EntityData.ParentYangName = "session"
    l2Tp.EntityData.SegmentPath = "l2tp"
    l2Tp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Tp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Tp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Tp.EntityData.Children = make(map[string]types.YChild)
    l2Tp.EntityData.Leafs = make(map[string]types.YLeaf)
    l2Tp.EntityData.Leafs["local-endpoint"] = types.YLeaf{"LocalEndpoint", l2Tp.LocalEndpoint}
    l2Tp.EntityData.Leafs["remote-endpoint"] = types.YLeaf{"RemoteEndpoint", l2Tp.RemoteEndpoint}
    l2Tp.EntityData.Leafs["call-serial-number"] = types.YLeaf{"CallSerialNumber", l2Tp.CallSerialNumber}
    l2Tp.EntityData.Leafs["is-l2tp-class-attribute-mask-set"] = types.YLeaf{"IsL2TpClassAttributeMaskSet", l2Tp.IsL2TpClassAttributeMaskSet}
    l2Tp.EntityData.Leafs["local-tunnel-id"] = types.YLeaf{"LocalTunnelId", l2Tp.LocalTunnelId}
    l2Tp.EntityData.Leafs["remote-tunnel-id"] = types.YLeaf{"RemoteTunnelId", l2Tp.RemoteTunnelId}
    l2Tp.EntityData.Leafs["local-session-id"] = types.YLeaf{"LocalSessionId", l2Tp.LocalSessionId}
    l2Tp.EntityData.Leafs["remote-session-id"] = types.YLeaf{"RemoteSessionId", l2Tp.RemoteSessionId}
    l2Tp.EntityData.Leafs["remote-port"] = types.YLeaf{"RemotePort", l2Tp.RemotePort}
    l2Tp.EntityData.Leafs["tunnel-client-authentication-id"] = types.YLeaf{"TunnelClientAuthenticationId", l2Tp.TunnelClientAuthenticationId}
    l2Tp.EntityData.Leafs["tunnel-server-authentication-id"] = types.YLeaf{"TunnelServerAuthenticationId", l2Tp.TunnelServerAuthenticationId}
    l2Tp.EntityData.Leafs["tunnel-assignment-id"] = types.YLeaf{"TunnelAssignmentId", l2Tp.TunnelAssignmentId}
    l2Tp.EntityData.Leafs["is-tunnel-authentication-enabled"] = types.YLeaf{"IsTunnelAuthenticationEnabled", l2Tp.IsTunnelAuthenticationEnabled}
    return &(l2Tp.EntityData)
}

// Vpdn_Sessions_Session_Subscriber
// Subscriber data
type Vpdn_Sessions_Session_Subscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NAS port type. The type is VpdnNasPort.
    NasPortType interface{}

    // Physical channel ID. The type is interface{} with range: 0..4294967295.
    PhysicalChannelId interface{}

    // Receive connect speed in nanoseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    ReceiveConnectSpeed interface{}

    // Transmit connect speed in nanoseconds. The type is interface{} with range:
    // 0..18446744073709551615. Units are nanosecond.
    TransmitConnectSpeed interface{}

    // NAS port ID. The type is slice of interface{} with range: 0..255.
    NasPort []interface{}
}

func (subscriber *Vpdn_Sessions_Session_Subscriber) GetEntityData() *types.CommonEntityData {
    subscriber.EntityData.YFilter = subscriber.YFilter
    subscriber.EntityData.YangName = "subscriber"
    subscriber.EntityData.BundleName = "cisco_ios_xr"
    subscriber.EntityData.ParentYangName = "session"
    subscriber.EntityData.SegmentPath = "subscriber"
    subscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriber.EntityData.Children = make(map[string]types.YChild)
    subscriber.EntityData.Leafs = make(map[string]types.YLeaf)
    subscriber.EntityData.Leafs["nas-port-type"] = types.YLeaf{"NasPortType", subscriber.NasPortType}
    subscriber.EntityData.Leafs["physical-channel-id"] = types.YLeaf{"PhysicalChannelId", subscriber.PhysicalChannelId}
    subscriber.EntityData.Leafs["receive-connect-speed"] = types.YLeaf{"ReceiveConnectSpeed", subscriber.ReceiveConnectSpeed}
    subscriber.EntityData.Leafs["transmit-connect-speed"] = types.YLeaf{"TransmitConnectSpeed", subscriber.TransmitConnectSpeed}
    subscriber.EntityData.Leafs["nas-port"] = types.YLeaf{"NasPort", subscriber.NasPort}
    return &(subscriber.EntityData)
}

// Vpdn_Sessions_Session_Configuration
// Configuration data
type Vpdn_Sessions_Session_Configuration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Template name. The type is string.
    TemplateName interface{}

    // VRF name. The type is string.
    VrfName interface{}

    // L2TP busy timeout in seconds. The type is interface{} with range: 0..65535.
    // Units are second.
    L2TpBusyTimeout interface{}

    // TOS mode. The type is TosMode.
    TosMode interface{}

    // TOS setting value. The type is interface{} with range: 0..255.
    Tos interface{}

    // True if DSL line info forwarding is enabled. The type is bool.
    DslLineForwarding interface{}

    // VPN ID.
    VpnId Vpdn_Sessions_Session_Configuration_VpnId
}

func (configuration *Vpdn_Sessions_Session_Configuration) GetEntityData() *types.CommonEntityData {
    configuration.EntityData.YFilter = configuration.YFilter
    configuration.EntityData.YangName = "configuration"
    configuration.EntityData.BundleName = "cisco_ios_xr"
    configuration.EntityData.ParentYangName = "session"
    configuration.EntityData.SegmentPath = "configuration"
    configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configuration.EntityData.Children = make(map[string]types.YChild)
    configuration.EntityData.Children["vpn-id"] = types.YChild{"VpnId", &configuration.VpnId}
    configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    configuration.EntityData.Leafs["template-name"] = types.YLeaf{"TemplateName", configuration.TemplateName}
    configuration.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", configuration.VrfName}
    configuration.EntityData.Leafs["l2tp-busy-timeout"] = types.YLeaf{"L2TpBusyTimeout", configuration.L2TpBusyTimeout}
    configuration.EntityData.Leafs["tos-mode"] = types.YLeaf{"TosMode", configuration.TosMode}
    configuration.EntityData.Leafs["tos"] = types.YLeaf{"Tos", configuration.Tos}
    configuration.EntityData.Leafs["dsl-line-forwarding"] = types.YLeaf{"DslLineForwarding", configuration.DslLineForwarding}
    return &(configuration.EntityData)
}

// Vpdn_Sessions_Session_Configuration_VpnId
// VPN ID
type Vpdn_Sessions_Session_Configuration_VpnId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OUI. The type is interface{} with range: 0..4294967295.
    Oui interface{}

    // Index. The type is interface{} with range: 0..4294967295.
    Index interface{}
}

func (vpnId *Vpdn_Sessions_Session_Configuration_VpnId) GetEntityData() *types.CommonEntityData {
    vpnId.EntityData.YFilter = vpnId.YFilter
    vpnId.EntityData.YangName = "vpn-id"
    vpnId.EntityData.BundleName = "cisco_ios_xr"
    vpnId.EntityData.ParentYangName = "configuration"
    vpnId.EntityData.SegmentPath = "vpn-id"
    vpnId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpnId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpnId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpnId.EntityData.Children = make(map[string]types.YChild)
    vpnId.EntityData.Leafs = make(map[string]types.YLeaf)
    vpnId.EntityData.Leafs["oui"] = types.YLeaf{"Oui", vpnId.Oui}
    vpnId.EntityData.Leafs["index"] = types.YLeaf{"Index", vpnId.Index}
    return &(vpnId.EntityData)
}

// Vpdn_TunnelDestinations
// VPDN tunnel Destinations
type Vpdn_TunnelDestinations struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN tunnel destination information. The type is slice of
    // Vpdn_TunnelDestinations_TunnelDestination.
    TunnelDestination []Vpdn_TunnelDestinations_TunnelDestination
}

func (tunnelDestinations *Vpdn_TunnelDestinations) GetEntityData() *types.CommonEntityData {
    tunnelDestinations.EntityData.YFilter = tunnelDestinations.YFilter
    tunnelDestinations.EntityData.YangName = "tunnel-destinations"
    tunnelDestinations.EntityData.BundleName = "cisco_ios_xr"
    tunnelDestinations.EntityData.ParentYangName = "vpdn"
    tunnelDestinations.EntityData.SegmentPath = "tunnel-destinations"
    tunnelDestinations.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDestinations.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDestinations.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDestinations.EntityData.Children = make(map[string]types.YChild)
    tunnelDestinations.EntityData.Children["tunnel-destination"] = types.YChild{"TunnelDestination", nil}
    for i := range tunnelDestinations.TunnelDestination {
        tunnelDestinations.EntityData.Children[types.GetSegmentPath(&tunnelDestinations.TunnelDestination[i])] = types.YChild{"TunnelDestination", &tunnelDestinations.TunnelDestination[i]}
    }
    tunnelDestinations.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(tunnelDestinations.EntityData)
}

// Vpdn_TunnelDestinations_TunnelDestination
// VPDN tunnel destination information
type Vpdn_TunnelDestinations_TunnelDestination struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // VRF name. The type is string.
    VrfNameXr interface{}

    // Current load on LNS. The type is interface{} with range: 0..4294967295.
    Load interface{}

    // Status of LNS. The type is LsgStatus.
    Status interface{}

    // Total count of tunnels connected. The type is interface{} with range:
    // 0..4294967295.
    Connects interface{}

    // Total count of tunnels disconnected. The type is interface{} with range:
    // 0..4294967295.
    Disconnects interface{}

    // Retries. The type is interface{} with range: 0..4294967295.
    Retry interface{}

    // Seconds since last status change. The type is interface{} with range:
    // 0..4294967295. Units are second.
    StatusChangeTime interface{}
}

func (tunnelDestination *Vpdn_TunnelDestinations_TunnelDestination) GetEntityData() *types.CommonEntityData {
    tunnelDestination.EntityData.YFilter = tunnelDestination.YFilter
    tunnelDestination.EntityData.YangName = "tunnel-destination"
    tunnelDestination.EntityData.BundleName = "cisco_ios_xr"
    tunnelDestination.EntityData.ParentYangName = "tunnel-destinations"
    tunnelDestination.EntityData.SegmentPath = "tunnel-destination"
    tunnelDestination.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnelDestination.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnelDestination.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnelDestination.EntityData.Children = make(map[string]types.YChild)
    tunnelDestination.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnelDestination.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", tunnelDestination.VrfName}
    tunnelDestination.EntityData.Leafs["address"] = types.YLeaf{"Address", tunnelDestination.Address}
    tunnelDestination.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", tunnelDestination.VrfNameXr}
    tunnelDestination.EntityData.Leafs["load"] = types.YLeaf{"Load", tunnelDestination.Load}
    tunnelDestination.EntityData.Leafs["status"] = types.YLeaf{"Status", tunnelDestination.Status}
    tunnelDestination.EntityData.Leafs["connects"] = types.YLeaf{"Connects", tunnelDestination.Connects}
    tunnelDestination.EntityData.Leafs["disconnects"] = types.YLeaf{"Disconnects", tunnelDestination.Disconnects}
    tunnelDestination.EntityData.Leafs["retry"] = types.YLeaf{"Retry", tunnelDestination.Retry}
    tunnelDestination.EntityData.Leafs["status-change-time"] = types.YLeaf{"StatusChangeTime", tunnelDestination.StatusChangeTime}
    return &(tunnelDestination.EntityData)
}

// Vpdn_VpdnMirroring
// VPDN Mirroring Statistics
type Vpdn_VpdnMirroring struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // sync not conn cnt. The type is interface{} with range: 0..4294967295.
    SyncNotConnCnt interface{}

    // sso err cnt. The type is interface{} with range: 0..4294967295.
    SsoErrCnt interface{}

    // sso batch err cnt. The type is interface{} with range: 0..4294967295.
    SsoBatchErrCnt interface{}

    // alloc err cnt. The type is interface{} with range: 0..4294967295.
    AllocErrCnt interface{}

    // alloc cnt. The type is interface{} with range: 0..4294967295.
    AllocCnt interface{}

    // qad send stats.
    QadSendStats Vpdn_VpdnMirroring_QadSendStats

    // qad recv stats.
    QadRecvStats Vpdn_VpdnMirroring_QadRecvStats

    // qad send stats last clear.
    QadSendStatsLastClear Vpdn_VpdnMirroring_QadSendStatsLastClear

    // qad recv stats last clear.
    QadRecvStatsLastClear Vpdn_VpdnMirroring_QadRecvStatsLastClear
}

func (vpdnMirroring *Vpdn_VpdnMirroring) GetEntityData() *types.CommonEntityData {
    vpdnMirroring.EntityData.YFilter = vpdnMirroring.YFilter
    vpdnMirroring.EntityData.YangName = "vpdn-mirroring"
    vpdnMirroring.EntityData.BundleName = "cisco_ios_xr"
    vpdnMirroring.EntityData.ParentYangName = "vpdn"
    vpdnMirroring.EntityData.SegmentPath = "vpdn-mirroring"
    vpdnMirroring.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdnMirroring.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdnMirroring.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdnMirroring.EntityData.Children = make(map[string]types.YChild)
    vpdnMirroring.EntityData.Children["qad-send-stats"] = types.YChild{"QadSendStats", &vpdnMirroring.QadSendStats}
    vpdnMirroring.EntityData.Children["qad-recv-stats"] = types.YChild{"QadRecvStats", &vpdnMirroring.QadRecvStats}
    vpdnMirroring.EntityData.Children["qad-send-stats-last-clear"] = types.YChild{"QadSendStatsLastClear", &vpdnMirroring.QadSendStatsLastClear}
    vpdnMirroring.EntityData.Children["qad-recv-stats-last-clear"] = types.YChild{"QadRecvStatsLastClear", &vpdnMirroring.QadRecvStatsLastClear}
    vpdnMirroring.EntityData.Leafs = make(map[string]types.YLeaf)
    vpdnMirroring.EntityData.Leafs["sync-not-conn-cnt"] = types.YLeaf{"SyncNotConnCnt", vpdnMirroring.SyncNotConnCnt}
    vpdnMirroring.EntityData.Leafs["sso-err-cnt"] = types.YLeaf{"SsoErrCnt", vpdnMirroring.SsoErrCnt}
    vpdnMirroring.EntityData.Leafs["sso-batch-err-cnt"] = types.YLeaf{"SsoBatchErrCnt", vpdnMirroring.SsoBatchErrCnt}
    vpdnMirroring.EntityData.Leafs["alloc-err-cnt"] = types.YLeaf{"AllocErrCnt", vpdnMirroring.AllocErrCnt}
    vpdnMirroring.EntityData.Leafs["alloc-cnt"] = types.YLeaf{"AllocCnt", vpdnMirroring.AllocCnt}
    return &(vpdnMirroring.EntityData)
}

// Vpdn_VpdnMirroring_QadSendStats
// qad send stats
type Vpdn_VpdnMirroring_QadSendStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // msgs sent. The type is interface{} with range: 0..4294967295.
    MsgsSent interface{}

    // acks sent. The type is interface{} with range: 0..4294967295.
    AcksSent interface{}

    // no partner. The type is interface{} with range: 0..4294967295.
    NoPartner interface{}

    // sends failed. The type is interface{} with range: 0..4294967295.
    SendsFailed interface{}

    // acks failed. The type is interface{} with range: 0..4294967295.
    AcksFailed interface{}

    // pending acks. The type is interface{} with range: 0..4294967295.
    PendingAcks interface{}

    // timeouts. The type is interface{} with range: 0..4294967295.
    Timeouts interface{}

    // suspends. The type is interface{} with range: 0..4294967295.
    Suspends interface{}

    // resumes. The type is interface{} with range: 0..4294967295.
    Resumes interface{}

    // sends fragment. The type is interface{} with range: 0..4294967295.
    SendsFragment interface{}

    // qad last seq number. The type is interface{} with range: 0..4294967295.
    QadLastSeqNumber interface{}

    // qad frag count. The type is interface{} with range: 0..4294967295.
    QadFragCount interface{}

    // qad ack count. The type is interface{} with range: 0..4294967295.
    QadAckCount interface{}

    // qad unknown acks. The type is interface{} with range: 0..4294967295.
    QadUnknownAcks interface{}

    // qad timeouts. The type is interface{} with range: 0..4294967295.
    QadTimeouts interface{}

    // qad rx count. The type is interface{} with range: 0..4294967295.
    QadRxCount interface{}

    // qad rx list count. The type is interface{} with range: 0..4294967295.
    QadRxListCount interface{}

    // qad rx list q size. The type is interface{} with range: 0..4294967295.
    QadRxListQSize interface{}

    // qad rx first seq number. The type is interface{} with range: 0..4294967295.
    QadRxFirstSeqNumber interface{}
}

func (qadSendStats *Vpdn_VpdnMirroring_QadSendStats) GetEntityData() *types.CommonEntityData {
    qadSendStats.EntityData.YFilter = qadSendStats.YFilter
    qadSendStats.EntityData.YangName = "qad-send-stats"
    qadSendStats.EntityData.BundleName = "cisco_ios_xr"
    qadSendStats.EntityData.ParentYangName = "vpdn-mirroring"
    qadSendStats.EntityData.SegmentPath = "qad-send-stats"
    qadSendStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadSendStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadSendStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadSendStats.EntityData.Children = make(map[string]types.YChild)
    qadSendStats.EntityData.Leafs = make(map[string]types.YLeaf)
    qadSendStats.EntityData.Leafs["msgs-sent"] = types.YLeaf{"MsgsSent", qadSendStats.MsgsSent}
    qadSendStats.EntityData.Leafs["acks-sent"] = types.YLeaf{"AcksSent", qadSendStats.AcksSent}
    qadSendStats.EntityData.Leafs["no-partner"] = types.YLeaf{"NoPartner", qadSendStats.NoPartner}
    qadSendStats.EntityData.Leafs["sends-failed"] = types.YLeaf{"SendsFailed", qadSendStats.SendsFailed}
    qadSendStats.EntityData.Leafs["acks-failed"] = types.YLeaf{"AcksFailed", qadSendStats.AcksFailed}
    qadSendStats.EntityData.Leafs["pending-acks"] = types.YLeaf{"PendingAcks", qadSendStats.PendingAcks}
    qadSendStats.EntityData.Leafs["timeouts"] = types.YLeaf{"Timeouts", qadSendStats.Timeouts}
    qadSendStats.EntityData.Leafs["suspends"] = types.YLeaf{"Suspends", qadSendStats.Suspends}
    qadSendStats.EntityData.Leafs["resumes"] = types.YLeaf{"Resumes", qadSendStats.Resumes}
    qadSendStats.EntityData.Leafs["sends-fragment"] = types.YLeaf{"SendsFragment", qadSendStats.SendsFragment}
    qadSendStats.EntityData.Leafs["qad-last-seq-number"] = types.YLeaf{"QadLastSeqNumber", qadSendStats.QadLastSeqNumber}
    qadSendStats.EntityData.Leafs["qad-frag-count"] = types.YLeaf{"QadFragCount", qadSendStats.QadFragCount}
    qadSendStats.EntityData.Leafs["qad-ack-count"] = types.YLeaf{"QadAckCount", qadSendStats.QadAckCount}
    qadSendStats.EntityData.Leafs["qad-unknown-acks"] = types.YLeaf{"QadUnknownAcks", qadSendStats.QadUnknownAcks}
    qadSendStats.EntityData.Leafs["qad-timeouts"] = types.YLeaf{"QadTimeouts", qadSendStats.QadTimeouts}
    qadSendStats.EntityData.Leafs["qad-rx-count"] = types.YLeaf{"QadRxCount", qadSendStats.QadRxCount}
    qadSendStats.EntityData.Leafs["qad-rx-list-count"] = types.YLeaf{"QadRxListCount", qadSendStats.QadRxListCount}
    qadSendStats.EntityData.Leafs["qad-rx-list-q-size"] = types.YLeaf{"QadRxListQSize", qadSendStats.QadRxListQSize}
    qadSendStats.EntityData.Leafs["qad-rx-first-seq-number"] = types.YLeaf{"QadRxFirstSeqNumber", qadSendStats.QadRxFirstSeqNumber}
    return &(qadSendStats.EntityData)
}

// Vpdn_VpdnMirroring_QadRecvStats
// qad recv stats
type Vpdn_VpdnMirroring_QadRecvStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // msgs recvd. The type is interface{} with range: 0..4294967295.
    MsgsRecvd interface{}

    // acks recvd. The type is interface{} with range: 0..4294967295.
    AcksRecvd interface{}

    // recvd acks failed. The type is interface{} with range: 0..4294967295.
    RecvdAcksFailed interface{}

    // init drops. The type is interface{} with range: 0..4294967295.
    InitDrops interface{}

    // msg drops. The type is interface{} with range: 0..4294967295.
    MsgDrops interface{}

    // ooo drops. The type is interface{} with range: 0..4294967295.
    OooDrops interface{}

    // stale msgs. The type is interface{} with range: 0..4294967295.
    StaleMsgs interface{}
}

func (qadRecvStats *Vpdn_VpdnMirroring_QadRecvStats) GetEntityData() *types.CommonEntityData {
    qadRecvStats.EntityData.YFilter = qadRecvStats.YFilter
    qadRecvStats.EntityData.YangName = "qad-recv-stats"
    qadRecvStats.EntityData.BundleName = "cisco_ios_xr"
    qadRecvStats.EntityData.ParentYangName = "vpdn-mirroring"
    qadRecvStats.EntityData.SegmentPath = "qad-recv-stats"
    qadRecvStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadRecvStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadRecvStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadRecvStats.EntityData.Children = make(map[string]types.YChild)
    qadRecvStats.EntityData.Leafs = make(map[string]types.YLeaf)
    qadRecvStats.EntityData.Leafs["msgs-recvd"] = types.YLeaf{"MsgsRecvd", qadRecvStats.MsgsRecvd}
    qadRecvStats.EntityData.Leafs["acks-recvd"] = types.YLeaf{"AcksRecvd", qadRecvStats.AcksRecvd}
    qadRecvStats.EntityData.Leafs["recvd-acks-failed"] = types.YLeaf{"RecvdAcksFailed", qadRecvStats.RecvdAcksFailed}
    qadRecvStats.EntityData.Leafs["init-drops"] = types.YLeaf{"InitDrops", qadRecvStats.InitDrops}
    qadRecvStats.EntityData.Leafs["msg-drops"] = types.YLeaf{"MsgDrops", qadRecvStats.MsgDrops}
    qadRecvStats.EntityData.Leafs["ooo-drops"] = types.YLeaf{"OooDrops", qadRecvStats.OooDrops}
    qadRecvStats.EntityData.Leafs["stale-msgs"] = types.YLeaf{"StaleMsgs", qadRecvStats.StaleMsgs}
    return &(qadRecvStats.EntityData)
}

// Vpdn_VpdnMirroring_QadSendStatsLastClear
// qad send stats last clear
type Vpdn_VpdnMirroring_QadSendStatsLastClear struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // msgs sent. The type is interface{} with range: 0..4294967295.
    MsgsSent interface{}

    // acks sent. The type is interface{} with range: 0..4294967295.
    AcksSent interface{}

    // no partner. The type is interface{} with range: 0..4294967295.
    NoPartner interface{}

    // sends failed. The type is interface{} with range: 0..4294967295.
    SendsFailed interface{}

    // acks failed. The type is interface{} with range: 0..4294967295.
    AcksFailed interface{}

    // pending acks. The type is interface{} with range: 0..4294967295.
    PendingAcks interface{}

    // timeouts. The type is interface{} with range: 0..4294967295.
    Timeouts interface{}

    // suspends. The type is interface{} with range: 0..4294967295.
    Suspends interface{}

    // resumes. The type is interface{} with range: 0..4294967295.
    Resumes interface{}

    // sends fragment. The type is interface{} with range: 0..4294967295.
    SendsFragment interface{}

    // qad last seq number. The type is interface{} with range: 0..4294967295.
    QadLastSeqNumber interface{}

    // qad frag count. The type is interface{} with range: 0..4294967295.
    QadFragCount interface{}

    // qad ack count. The type is interface{} with range: 0..4294967295.
    QadAckCount interface{}

    // qad unknown acks. The type is interface{} with range: 0..4294967295.
    QadUnknownAcks interface{}

    // qad timeouts. The type is interface{} with range: 0..4294967295.
    QadTimeouts interface{}

    // qad rx count. The type is interface{} with range: 0..4294967295.
    QadRxCount interface{}

    // qad rx list count. The type is interface{} with range: 0..4294967295.
    QadRxListCount interface{}

    // qad rx list q size. The type is interface{} with range: 0..4294967295.
    QadRxListQSize interface{}

    // qad rx first seq number. The type is interface{} with range: 0..4294967295.
    QadRxFirstSeqNumber interface{}
}

func (qadSendStatsLastClear *Vpdn_VpdnMirroring_QadSendStatsLastClear) GetEntityData() *types.CommonEntityData {
    qadSendStatsLastClear.EntityData.YFilter = qadSendStatsLastClear.YFilter
    qadSendStatsLastClear.EntityData.YangName = "qad-send-stats-last-clear"
    qadSendStatsLastClear.EntityData.BundleName = "cisco_ios_xr"
    qadSendStatsLastClear.EntityData.ParentYangName = "vpdn-mirroring"
    qadSendStatsLastClear.EntityData.SegmentPath = "qad-send-stats-last-clear"
    qadSendStatsLastClear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadSendStatsLastClear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadSendStatsLastClear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadSendStatsLastClear.EntityData.Children = make(map[string]types.YChild)
    qadSendStatsLastClear.EntityData.Leafs = make(map[string]types.YLeaf)
    qadSendStatsLastClear.EntityData.Leafs["msgs-sent"] = types.YLeaf{"MsgsSent", qadSendStatsLastClear.MsgsSent}
    qadSendStatsLastClear.EntityData.Leafs["acks-sent"] = types.YLeaf{"AcksSent", qadSendStatsLastClear.AcksSent}
    qadSendStatsLastClear.EntityData.Leafs["no-partner"] = types.YLeaf{"NoPartner", qadSendStatsLastClear.NoPartner}
    qadSendStatsLastClear.EntityData.Leafs["sends-failed"] = types.YLeaf{"SendsFailed", qadSendStatsLastClear.SendsFailed}
    qadSendStatsLastClear.EntityData.Leafs["acks-failed"] = types.YLeaf{"AcksFailed", qadSendStatsLastClear.AcksFailed}
    qadSendStatsLastClear.EntityData.Leafs["pending-acks"] = types.YLeaf{"PendingAcks", qadSendStatsLastClear.PendingAcks}
    qadSendStatsLastClear.EntityData.Leafs["timeouts"] = types.YLeaf{"Timeouts", qadSendStatsLastClear.Timeouts}
    qadSendStatsLastClear.EntityData.Leafs["suspends"] = types.YLeaf{"Suspends", qadSendStatsLastClear.Suspends}
    qadSendStatsLastClear.EntityData.Leafs["resumes"] = types.YLeaf{"Resumes", qadSendStatsLastClear.Resumes}
    qadSendStatsLastClear.EntityData.Leafs["sends-fragment"] = types.YLeaf{"SendsFragment", qadSendStatsLastClear.SendsFragment}
    qadSendStatsLastClear.EntityData.Leafs["qad-last-seq-number"] = types.YLeaf{"QadLastSeqNumber", qadSendStatsLastClear.QadLastSeqNumber}
    qadSendStatsLastClear.EntityData.Leafs["qad-frag-count"] = types.YLeaf{"QadFragCount", qadSendStatsLastClear.QadFragCount}
    qadSendStatsLastClear.EntityData.Leafs["qad-ack-count"] = types.YLeaf{"QadAckCount", qadSendStatsLastClear.QadAckCount}
    qadSendStatsLastClear.EntityData.Leafs["qad-unknown-acks"] = types.YLeaf{"QadUnknownAcks", qadSendStatsLastClear.QadUnknownAcks}
    qadSendStatsLastClear.EntityData.Leafs["qad-timeouts"] = types.YLeaf{"QadTimeouts", qadSendStatsLastClear.QadTimeouts}
    qadSendStatsLastClear.EntityData.Leafs["qad-rx-count"] = types.YLeaf{"QadRxCount", qadSendStatsLastClear.QadRxCount}
    qadSendStatsLastClear.EntityData.Leafs["qad-rx-list-count"] = types.YLeaf{"QadRxListCount", qadSendStatsLastClear.QadRxListCount}
    qadSendStatsLastClear.EntityData.Leafs["qad-rx-list-q-size"] = types.YLeaf{"QadRxListQSize", qadSendStatsLastClear.QadRxListQSize}
    qadSendStatsLastClear.EntityData.Leafs["qad-rx-first-seq-number"] = types.YLeaf{"QadRxFirstSeqNumber", qadSendStatsLastClear.QadRxFirstSeqNumber}
    return &(qadSendStatsLastClear.EntityData)
}

// Vpdn_VpdnMirroring_QadRecvStatsLastClear
// qad recv stats last clear
type Vpdn_VpdnMirroring_QadRecvStatsLastClear struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // msgs recvd. The type is interface{} with range: 0..4294967295.
    MsgsRecvd interface{}

    // acks recvd. The type is interface{} with range: 0..4294967295.
    AcksRecvd interface{}

    // recvd acks failed. The type is interface{} with range: 0..4294967295.
    RecvdAcksFailed interface{}

    // init drops. The type is interface{} with range: 0..4294967295.
    InitDrops interface{}

    // msg drops. The type is interface{} with range: 0..4294967295.
    MsgDrops interface{}

    // ooo drops. The type is interface{} with range: 0..4294967295.
    OooDrops interface{}

    // stale msgs. The type is interface{} with range: 0..4294967295.
    StaleMsgs interface{}
}

func (qadRecvStatsLastClear *Vpdn_VpdnMirroring_QadRecvStatsLastClear) GetEntityData() *types.CommonEntityData {
    qadRecvStatsLastClear.EntityData.YFilter = qadRecvStatsLastClear.YFilter
    qadRecvStatsLastClear.EntityData.YangName = "qad-recv-stats-last-clear"
    qadRecvStatsLastClear.EntityData.BundleName = "cisco_ios_xr"
    qadRecvStatsLastClear.EntityData.ParentYangName = "vpdn-mirroring"
    qadRecvStatsLastClear.EntityData.SegmentPath = "qad-recv-stats-last-clear"
    qadRecvStatsLastClear.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    qadRecvStatsLastClear.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    qadRecvStatsLastClear.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    qadRecvStatsLastClear.EntityData.Children = make(map[string]types.YChild)
    qadRecvStatsLastClear.EntityData.Leafs = make(map[string]types.YLeaf)
    qadRecvStatsLastClear.EntityData.Leafs["msgs-recvd"] = types.YLeaf{"MsgsRecvd", qadRecvStatsLastClear.MsgsRecvd}
    qadRecvStatsLastClear.EntityData.Leafs["acks-recvd"] = types.YLeaf{"AcksRecvd", qadRecvStatsLastClear.AcksRecvd}
    qadRecvStatsLastClear.EntityData.Leafs["recvd-acks-failed"] = types.YLeaf{"RecvdAcksFailed", qadRecvStatsLastClear.RecvdAcksFailed}
    qadRecvStatsLastClear.EntityData.Leafs["init-drops"] = types.YLeaf{"InitDrops", qadRecvStatsLastClear.InitDrops}
    qadRecvStatsLastClear.EntityData.Leafs["msg-drops"] = types.YLeaf{"MsgDrops", qadRecvStatsLastClear.MsgDrops}
    qadRecvStatsLastClear.EntityData.Leafs["ooo-drops"] = types.YLeaf{"OooDrops", qadRecvStatsLastClear.OooDrops}
    qadRecvStatsLastClear.EntityData.Leafs["stale-msgs"] = types.YLeaf{"StaleMsgs", qadRecvStatsLastClear.StaleMsgs}
    return &(qadRecvStatsLastClear.EntityData)
}

// Vpdn_VpdnRedundancy
// Show VPDN Redundancy information
type Vpdn_VpdnRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // session total. The type is interface{} with range: 0..4294967295.
    SessionTotal interface{}

    // session synced. The type is interface{} with range: 0..4294967295.
    SessionSynced interface{}

    // state. The type is VpdnState.
    State interface{}

    // start time. The type is interface{} with range: 0..18446744073709551615.
    StartTime interface{}

    // finish time. The type is interface{} with range: 0..18446744073709551615.
    FinishTime interface{}

    // abort time. The type is interface{} with range: 0..18446744073709551615.
    AbortTime interface{}
}

func (vpdnRedundancy *Vpdn_VpdnRedundancy) GetEntityData() *types.CommonEntityData {
    vpdnRedundancy.EntityData.YFilter = vpdnRedundancy.YFilter
    vpdnRedundancy.EntityData.YangName = "vpdn-redundancy"
    vpdnRedundancy.EntityData.BundleName = "cisco_ios_xr"
    vpdnRedundancy.EntityData.ParentYangName = "vpdn"
    vpdnRedundancy.EntityData.SegmentPath = "vpdn-redundancy"
    vpdnRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpdnRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpdnRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpdnRedundancy.EntityData.Children = make(map[string]types.YChild)
    vpdnRedundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    vpdnRedundancy.EntityData.Leafs["session-total"] = types.YLeaf{"SessionTotal", vpdnRedundancy.SessionTotal}
    vpdnRedundancy.EntityData.Leafs["session-synced"] = types.YLeaf{"SessionSynced", vpdnRedundancy.SessionSynced}
    vpdnRedundancy.EntityData.Leafs["state"] = types.YLeaf{"State", vpdnRedundancy.State}
    vpdnRedundancy.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", vpdnRedundancy.StartTime}
    vpdnRedundancy.EntityData.Leafs["finish-time"] = types.YLeaf{"FinishTime", vpdnRedundancy.FinishTime}
    vpdnRedundancy.EntityData.Leafs["abort-time"] = types.YLeaf{"AbortTime", vpdnRedundancy.AbortTime}
    return &(vpdnRedundancy.EntityData)
}

// Vpdn_HistoryFailures
// VPDN history failure list
type Vpdn_HistoryFailures struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPDN history failure information. The type is slice of
    // Vpdn_HistoryFailures_HistoryFailure.
    HistoryFailure []Vpdn_HistoryFailures_HistoryFailure
}

func (historyFailures *Vpdn_HistoryFailures) GetEntityData() *types.CommonEntityData {
    historyFailures.EntityData.YFilter = historyFailures.YFilter
    historyFailures.EntityData.YangName = "history-failures"
    historyFailures.EntityData.BundleName = "cisco_ios_xr"
    historyFailures.EntityData.ParentYangName = "vpdn"
    historyFailures.EntityData.SegmentPath = "history-failures"
    historyFailures.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyFailures.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyFailures.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyFailures.EntityData.Children = make(map[string]types.YChild)
    historyFailures.EntityData.Children["history-failure"] = types.YChild{"HistoryFailure", nil}
    for i := range historyFailures.HistoryFailure {
        historyFailures.EntityData.Children[types.GetSegmentPath(&historyFailures.HistoryFailure[i])] = types.YChild{"HistoryFailure", &historyFailures.HistoryFailure[i]}
    }
    historyFailures.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(historyFailures.EntityData)
}

// Vpdn_HistoryFailures_HistoryFailure
// VPDN history failure information
type Vpdn_HistoryFailures_HistoryFailure struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Username. The type is string.
    Username interface{}

    // Remote name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    RemoteName interface{}

    // Authentication username. The type is string.
    UsernameXr interface{}

    // Domain name. The type is string.
    DomainName interface{}

    // VPDN user session ID. The type is interface{} with range: 0..65535.
    Mid interface{}

    // Network access server. The type is string.
    Nas interface{}

    // NAS IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Remote client ID. The type is interface{} with range: 0..65535.
    RemoteClientId interface{}

    // Home gateway. The type is string.
    HomeGateway interface{}

    // Source IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // Local client ID. The type is interface{} with range: 0..65535.
    LocalClientId interface{}

    // Event logged time in Ex: Wed Aug  3 10:28:30 2011. The type is string.
    EventTime interface{}

    // Error repeat count. The type is interface{} with range: 0..65535.
    ErrorRepeatCount interface{}

    // Failure type. The type is VpdnFailcode.
    FailureType interface{}
}

func (historyFailure *Vpdn_HistoryFailures_HistoryFailure) GetEntityData() *types.CommonEntityData {
    historyFailure.EntityData.YFilter = historyFailure.YFilter
    historyFailure.EntityData.YangName = "history-failure"
    historyFailure.EntityData.BundleName = "cisco_ios_xr"
    historyFailure.EntityData.ParentYangName = "history-failures"
    historyFailure.EntityData.SegmentPath = "history-failure"
    historyFailure.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historyFailure.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historyFailure.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historyFailure.EntityData.Children = make(map[string]types.YChild)
    historyFailure.EntityData.Leafs = make(map[string]types.YLeaf)
    historyFailure.EntityData.Leafs["username"] = types.YLeaf{"Username", historyFailure.Username}
    historyFailure.EntityData.Leafs["remote-name"] = types.YLeaf{"RemoteName", historyFailure.RemoteName}
    historyFailure.EntityData.Leafs["username-xr"] = types.YLeaf{"UsernameXr", historyFailure.UsernameXr}
    historyFailure.EntityData.Leafs["domain-name"] = types.YLeaf{"DomainName", historyFailure.DomainName}
    historyFailure.EntityData.Leafs["mid"] = types.YLeaf{"Mid", historyFailure.Mid}
    historyFailure.EntityData.Leafs["nas"] = types.YLeaf{"Nas", historyFailure.Nas}
    historyFailure.EntityData.Leafs["destination-address"] = types.YLeaf{"DestinationAddress", historyFailure.DestinationAddress}
    historyFailure.EntityData.Leafs["remote-client-id"] = types.YLeaf{"RemoteClientId", historyFailure.RemoteClientId}
    historyFailure.EntityData.Leafs["home-gateway"] = types.YLeaf{"HomeGateway", historyFailure.HomeGateway}
    historyFailure.EntityData.Leafs["source-address"] = types.YLeaf{"SourceAddress", historyFailure.SourceAddress}
    historyFailure.EntityData.Leafs["local-client-id"] = types.YLeaf{"LocalClientId", historyFailure.LocalClientId}
    historyFailure.EntityData.Leafs["event-time"] = types.YLeaf{"EventTime", historyFailure.EventTime}
    historyFailure.EntityData.Leafs["error-repeat-count"] = types.YLeaf{"ErrorRepeatCount", historyFailure.ErrorRepeatCount}
    historyFailure.EntityData.Leafs["failure-type"] = types.YLeaf{"FailureType", historyFailure.FailureType}
    return &(historyFailure.EntityData)
}

